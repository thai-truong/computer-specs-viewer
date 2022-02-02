package gpugui

import (
	custom_types "computer-specs-viewer/gui/custom-types"
	gpuinfo "computer-specs-viewer/src/gpu_info"
	"computer-specs-viewer/utils"
	"fmt"
	"reflect"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type GpuInformationGui struct {
	Name                 string
	Version              string
	LastModified         time.Time
	Resolution           string
	NumberOfColors       uint64
	CurrentRefreshRate   uint32
	RefreshRateRange     string
	BitsPerPixel         uint32
	Status               string
	Availability         string
	MemorySize           custom_types.Space
	AdapterCompatibility string
	AdapterDacType       string
	IsMonochrome         bool
	Manufacturer         string
	PresentOnSystem      bool
}

func transformInput(gpu gpuinfo.GpuInformation) GpuInformationGui {
	return GpuInformationGui{
		Name:                 gpu.Name,
		Version:              gpu.Version,
		LastModified:         gpu.LastModified,
		Resolution:           gpu.Resolution,
		NumberOfColors:       gpu.NumColors,
		CurrentRefreshRate:   gpu.CurrRefreshRate,
		RefreshRateRange:     gpu.RefreshRateRange,
		BitsPerPixel:         gpu.BitsPerPixel,
		Status:               gpu.Status,
		Availability:         gpu.Availability,
		MemorySize:           custom_types.NumToCustomSpaceType(uint64(gpu.MemorySize)),
		AdapterCompatibility: gpu.AdapterCompatibility,
		AdapterDacType:       gpu.AdapterDACType,
		IsMonochrome:         gpu.IsMonochrome,
		Manufacturer:         gpu.Manufacturer,
		PresentOnSystem:      gpu.PresentOnSystem,
	}
}

func getGpuInfoStrings(gpuGui GpuInformationGui) []string {
	res := []string{}
	diskValues := reflect.ValueOf(gpuGui)
	diskType := diskValues.Type()

	for i := 0; i < diskValues.NumField(); i++ {
		infoFieldName := utils.SpaceOutFieldNames(diskType.Field(i).Name)
		infoFieldValue := diskValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func getGpuInfoLabel(gpuGui GpuInformationGui) fyne.CanvasObject {
	return utils.CreateStrSliceInfoLabel(getGpuInfoStrings(gpuGui))
}

func CreateInfoScreen(_ fyne.Window) fyne.CanvasObject {
	gpuSlice := gpuinfo.GetGPUsInformation()
	gpuAccordion := widget.NewAccordion()

	for i, gpu := range gpuSlice {
		gpuGui := transformInput(gpu)
		gpuOrder := fmt.Sprint(i + 1)

		gpuInfoLabel := getGpuInfoLabel(gpuGui)
		accordionItem := utils.CreateAccordionItem("GPU", gpuOrder, []fyne.CanvasObject{gpuInfoLabel})
		gpuAccordion.Append(accordionItem)
	}

	return utils.NewScrollVBox(gpuAccordion)
}
