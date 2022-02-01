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

func convertInfoStringsToLabels(infoStrs []string) []fyne.CanvasObject {
	labels := []fyne.CanvasObject{}

	for _, infoStr := range infoStrs {
		labels = append(labels, widget.NewLabel(infoStr))
	}

	return labels
}

func getInfoLabels(cpuGui CpuInformationGui) []fyne.CanvasObject {
	return convertInfoStringsToLabels(getCpuInfoStrings(cpuGui))
}

func CreateInfoScreen(_ fyne.Window) fyne.CanvasObject {
	cpuList := cpuinfo.GetCpuInfo()
	allCpuInfo := []fyne.CanvasObject{}

	for i, cpu := range cpuList {
		cpuGui := CpuInformationGui(cpu)

		orderLabel := []fyne.CanvasObject{widget.NewLabel(fmt.Sprintf("CPU #%v", (i + 1)))}
		infoLabels := append(orderLabel, getInfoLabels(cpuGui)...)
		addSeparator(i, len(cpuList), infoLabels)

		if i < len(cpuList)-1 {
			infoLabels = append(infoLabels, widget.NewSeparator())
		}

		cpuContent := container.NewVBox(infoLabels...)
		allCpuInfo = append(allCpuInfo, cpuContent)
	}

	return container.NewCenter(allCpuInfo...)
}

func addSeparator(idx int, infoListLen int, infoLabels []fyne.CanvasObject) {
	if idx < infoListLen {
		infoLabels = append(infoLabels, widget.NewSeparator())
	}
}
