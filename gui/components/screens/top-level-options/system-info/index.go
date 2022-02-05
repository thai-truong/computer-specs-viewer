package systeminfo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func GetTitle() string {
	return "System Information"
}

func GetDesc() string {
	return "This page contains links to information about different parts of your computer (at least the ones we support)"
}

func CreateScreen(_ fyne.Window) fyne.CanvasObject {
	content := container.NewVBox()

	return container.NewCenter(content)
}
