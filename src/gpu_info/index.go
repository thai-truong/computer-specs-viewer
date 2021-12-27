package gpuinfo

import (
	"computer-specs-viewer/src"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/yusufpapurcu/wmi"
)

var Timeout = 3 * time.Second

type Win32_VideoController struct {
	Name                        string
	DriverVersion               string
	DriverDate                  time.Time
	CurrentHorizontalResolution uint32
	CurrentVerticalResolution   uint32
	CurrentNumberOfColors       uint64
	CurrentRefreshRate          uint32
	MaxRefreshRate              uint32
	MinRefreshRate              uint32
	CurrentBitsPerPixel         uint32
	Status                      string
	Availability                uint16
	AdapterRAM                  uint32 // Currently WMI gives the wrong value for this https://docs.microsoft.com/en-us/troubleshoot/windows-client/deployment/msinfo32-report-wrong-display-adapter-ram
	AdapterCompatibility        string
	Monochrome                  bool
	PNPDeviceID                 string
	VideoArchitecture           uint16
}

type Win32_PnPEntity struct {
	Manufacturer string
	Present      bool
}

type GpuInformation struct {
	Index                int
	Name                 string
	Version              string
	LastModified         time.Time
	Resolution           string
	NumColors            uint64
	CurrRefreshRate      uint32
	RefreshRateRange     string
	BitsPerPixel         uint32
	Status               string
	Availability         string
	MemorySize           uint32
	AdapterCompatibility string
	IsMonochrome         bool
	Manufacturer         string
	PresentOnSystem      bool
	VideoArchitecture    string
}

// WMIQueryWithContext - wraps wmi.Query with a timed-out context to avoid hanging
func WMIQueryWithContext(ctx context.Context, query string, dst interface{}, connectServerArgs ...interface{}) error {
	if _, ok := ctx.Deadline(); !ok {
		ctxTimeout, cancel := context.WithTimeout(ctx, Timeout)
		defer cancel()
		ctx = ctxTimeout
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- wmi.Query(query, dst, connectServerArgs...)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func GetGPUsInformation() []GpuInformation {
	gpus, err := GetGPUsInformationWithErr()

	if err != nil {
		return []GpuInformation{}
	}

	return gpus
}

func GetGPUsInformationWithErr() ([]GpuInformation, error) {
	var gpus []GpuInformation
	var dst []Win32_VideoController

	ctx := context.Background()
	query := wmi.CreateQuery(&dst, "")

	if err := WMIQueryWithContext(ctx, query, &dst); err != nil {
		return gpus, err
	}

	for idx, gpu := range dst {
		gpuInfo := GpuInformation{
			Index:                idx,
			Name:                 gpu.Name,
			Version:              gpu.DriverVersion,
			LastModified:         gpu.DriverDate,
			Resolution:           getResolutionStr(gpu.CurrentHorizontalResolution, gpu.CurrentVerticalResolution),
			NumColors:            gpu.CurrentNumberOfColors,
			CurrRefreshRate:      gpu.CurrentRefreshRate,
			RefreshRateRange:     getRefreshRateRangeStr(gpu.MinRefreshRate, gpu.MaxRefreshRate),
			BitsPerPixel:         gpu.CurrentBitsPerPixel,
			Status:               gpu.Status,
			Availability:         AvailabilityMapping[gpu.Availability],
			MemorySize:           gpu.AdapterRAM,
			AdapterCompatibility: gpu.AdapterCompatibility,
			IsMonochrome:         gpu.Monochrome,
			VideoArchitecture:    VideoArchitectureMapping[gpu.VideoArchitecture],
		}

		var pnpDst []Win32_PnPEntity

		escapedDeviceID := strings.Replace(gpu.PNPDeviceID, "\\", "\\\\", -1)
		whereClause := fmt.Sprintf("WHERE DeviceID = '%s'", escapedDeviceID)
		pnpQuery := wmi.CreateQuery(&pnpDst, whereClause)

		pnpErr := WMIQueryWithContext(ctx, pnpQuery, &pnpDst)

		if pnpErr != nil {
			fmt.Printf("Error while looking for corresponding PNPEntity of current GPU: %v", pnpErr)
		} else {
			var manufacturer string
			var present bool

			// Should only run for 1 iteration since DeviceID is unique
			for _, pnpGPU := range pnpDst {
				manufacturer = pnpGPU.Manufacturer
				present = pnpGPU.Present
			}

			gpuInfo.Manufacturer = manufacturer
			gpuInfo.PresentOnSystem = present
		}

		gpus = append(gpus, gpuInfo)
	}

	return gpus, nil
}

func getResolutionStr(hRes uint32, vRes uint32) string {
	return fmt.Sprintf("%v x %v", hRes, vRes)
}

func getRefreshRateRangeStr(minRate uint32, maxRate uint32) string {
	return fmt.Sprintf("%v - %v", minRate, maxRate)
}

func PrintGpusInfo() {
	gpus := GetGPUsInformation()

	src.PrintSectionTitle("GPU Devices")
	src.PrintStartBlock()

	for i := 0; i < len(gpus); i++ {
		printSingleGPUInfo(gpus[i])

		if i < len(gpus)-1 {
			src.PrintInfoDelim()
		}
	}

	src.PrintEndBlock()
}

func printSingleGPUInfo(gpu GpuInformation) {
	src.PrintStrWithOrder("GPU device", gpu.Index)
	fmt.Printf("Name: %v\n", gpu.Name)
	fmt.Printf("Manufacturer: %v\n", gpu.Manufacturer)
	fmt.Printf("Version: %v\n", gpu.Version)
	fmt.Printf("Last modified: %v\n", gpu.LastModified)
	fmt.Printf("Resolution: %v\n", gpu.Resolution)
	fmt.Printf("Number of colors available: %v\n", gpu.NumColors)
	fmt.Printf("Current refresh rate: %v\n", gpu.CurrRefreshRate)
	fmt.Printf("Refresh rate range: %v\n", gpu.RefreshRateRange)
	fmt.Printf("Number of bits used per pixel: %v\n", gpu.BitsPerPixel)
	fmt.Printf("Device status: %v\n", gpu.Status)
	fmt.Printf("Device availability: %v\n", gpu.Availability)
	fmt.Printf("Present on system: %v\n", gpu.PresentOnSystem)
	fmt.Printf("Memory (RAM) size: %v\n", src.GetSpaceString(uint64(gpu.MemorySize), "MB"))
	fmt.Printf("Adapter compatibility: %v compatible\n", gpu.AdapterCompatibility)
	fmt.Printf("Is monochrome: %v\n", gpu.IsMonochrome)
	fmt.Printf("Video architecture: %v\n", gpu.VideoArchitecture)
}
