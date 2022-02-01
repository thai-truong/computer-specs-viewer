package cpugui

import (
	cpuinfo "computer-specs-viewer/src/cpu_info"
	"computer-specs-viewer/utils"
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CpuInformationGui cpuinfo.CpuInformation

func getCpuInfoStrings(cpu CpuInformationGui) []string {
	res := []string{}
	cpuValues := reflect.ValueOf(cpu)
	cpuType := cpuValues.Type()

	for i := 0; i < cpuValues.NumField(); i++ {
		infoFieldName := utils.SpaceOutFieldNames(cpuType.Field(i).Name)
		infoFieldValue := cpuValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func getInfoLabels(cpuGui CpuInformationGui) []fyne.CanvasObject {
	return utils.ConvertStringsToLabels(getCpuInfoStrings(cpuGui))
}

func CreateInfoScreen(_ fyne.Window) fyne.CanvasObject {
	cpuSlice := cpuinfo.GetCpuInfo()
	cpuAccordion := widget.NewAccordion()

	for i, cpu := range cpuSlice {
		cpuGui := CpuInformationGui(cpu)

		cpuOrder := fmt.Sprint(i + 1)
		accordionItem := CreateAccordionItem(cpuOrder, cpuGui)
		cpuAccordion.Append(accordionItem)
	}

	return utils.NewScrollVBox(cpuAccordion)
}

func CreateAccordionItem(order string, cpu CpuInformationGui) *widget.AccordionItem {
	title := utils.GetStrWithOrder("CPU", order)
	infoLabels := getInfoLabels(cpu)

	return widget.NewAccordionItem(title, container.NewVBox(infoLabels...))
}
