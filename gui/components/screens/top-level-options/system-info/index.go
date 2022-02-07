package systeminfo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func GetTitle() string {
	return "System Information"
}

func GetDesc() string {
	return ""
}

func CreateScreen(w fyne.Window) fyne.CanvasObject {
	return container.NewCenter()
}
