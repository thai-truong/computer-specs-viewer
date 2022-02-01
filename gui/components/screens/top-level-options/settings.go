package top_level_options

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func SettingsScreen(_ fyne.Window) fyne.CanvasObject {
	content := container.NewVBox()

	return container.NewCenter(content)
}
