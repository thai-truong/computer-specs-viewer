package utils

import (
	"fmt"

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
