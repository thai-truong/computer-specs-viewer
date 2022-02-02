package cpugui

import (
	cpuinfo "computer-specs-viewer/src/cpu_info"
	"computer-specs-viewer/utils"
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CpuInformationGui cpuinfo.CpuInformation

func getCpuInfoStrings(cpuGui CpuInformationGui) []string {
	res := []string{}
	cpuValues := reflect.ValueOf(cpuGui)
	cpuType := cpuValues.Type()

	for i := 0; i < cpuValues.NumField(); i++ {
		infoFieldName := utils.SpaceOutFieldNames(cpuType.Field(i).Name)
		infoFieldValue := cpuValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func getCpuInfoLabel(cpuGui CpuInformationGui) fyne.CanvasObject {
	return utils.CreateStrSliceInfoLabel(getCpuInfoStrings(cpuGui))
}

func CreateInfoScreen(_ fyne.Window) fyne.CanvasObject {
	cpuSlice := cpuinfo.GetCpuInfo()
	cpuAccordion := widget.NewAccordion()

	for i, cpu := range cpuSlice {
		cpuGui := CpuInformationGui(cpu)
		cpuOrder := fmt.Sprint(i + 1)

		cpuInfoLabel := getCpuInfoLabel(cpuGui)
		accordionItem := utils.CreateAccordionItem("CPU", cpuOrder, []fyne.CanvasObject{cpuInfoLabel})
		cpuAccordion.Append(accordionItem)
	}

	return utils.NewScrollVBox(cpuAccordion)
}
