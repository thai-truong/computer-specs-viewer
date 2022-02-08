package cpugui

import (
	cpuinfo "computer-specs-viewer/src/cpu_info"
	"computer-specs-viewer/utils"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CpuInformationGui cpuinfo.CpuInformation

func GetTitle() string {
	return "CPU"
}

func GetDesc() string {
	return "This page contains information about this computer's CPU"
}

func CreateScreen(_ fyne.Window) fyne.CanvasObject {
	cpuSlice := cpuinfo.GetCpuInfo()
	cpuAccordion := widget.NewAccordion()

	for i, cpu := range cpuSlice {
		cpuGui := CpuInformationGui(cpu)
		cpuOrder := fmt.Sprint(i + 1)

		cpuInfoLabel := utils.GetInfoGuiLabel(cpuGui)
		accordionItem := utils.CreateAccordionItem("CPU", cpuOrder, []fyne.CanvasObject{cpuInfoLabel})
		cpuAccordion.Append(accordionItem)
	}

	return utils.NewScrollVBox(cpuAccordion)
}
