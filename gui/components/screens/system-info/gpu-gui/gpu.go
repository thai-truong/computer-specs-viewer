package gpugui

import (
	custom_types "computer-specs-viewer/gui/custom-types"
	gpuinfo "computer-specs-viewer/src/gpu_info"
	"computer-specs-viewer/utils"
	"fmt"
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

func GetTitle() string {
	return "GPU"
}

func GetDesc() string {
	return "This page contains information about this computer's GPU"
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

func CreateScreen(_ fyne.Window) fyne.CanvasObject {
	gpuSlice := gpuinfo.GetGPUsInformation()
	gpuAccordion := widget.NewAccordion()

	for i, gpu := range gpuSlice {
		gpuGui := transformInput(gpu)
		gpuOrder := fmt.Sprint(i + 1)

		gpuInfoLabel := utils.GetInfoGuiLabel(gpuGui)
		accordionItem := utils.CreateAccordionItem("GPU", gpuOrder, []fyne.CanvasObject{gpuInfoLabel})
		gpuAccordion.Append(accordionItem)
	}

	return utils.NewScrollVBox(gpuAccordion)
}
