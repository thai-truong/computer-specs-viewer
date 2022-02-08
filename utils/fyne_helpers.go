package utils

import (
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ConvertStringsToLabels(strs []string) []fyne.CanvasObject {
	labels := []fyne.CanvasObject{}

	for _, str := range strs {
		labels = append(labels, widget.NewLabel(str))
	}

	return labels
}

func SliceToSingleFyneLabel(strs []string) fyne.CanvasObject {
	fyneLabelText := ""

	for _, label := range strs {
		fyneLabelText += label + "\n"
	}

	return widget.NewLabel(fyneLabelText)
}

func GetOrderLabel(name string, order string) *widget.Label {
	return widget.NewLabel(fmt.Sprintf("%v %v\n", name, order))
}

func NewScrollVBox(content ...fyne.CanvasObject) fyne.CanvasObject {
	return container.NewScroll(container.NewVBox(content...))
}

func CreateAccordionItem(labelName string, order string, objs []fyne.CanvasObject) *widget.AccordionItem {
	title := GetStrWithOrder(labelName, order)
	return widget.NewAccordionItem(title, container.NewVBox(objs...))
}

func CreateSection(labelName string, content fyne.CanvasObject) fyne.CanvasObject {
	label := widget.NewLabel(labelName)
	section := container.NewVBox(label, content)

	return section
}

func GetInfoGuiStrings(infoGui interface{}) []string {
	res := []string{}
	cpuValues := reflect.ValueOf(infoGui)
	cpuType := cpuValues.Type()

	for i := 0; i < cpuValues.NumField(); i++ {
		infoFieldName := SpaceOutFieldNames(cpuType.Field(i).Name)
		infoFieldValue := cpuValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func GetInfoGuiLabel(infoGui interface{}) fyne.CanvasObject {
	return SliceToSingleFyneLabel(GetInfoGuiStrings(infoGui))
}
