package settings

import (
	"computer-specs-viewer/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	SliderMin = 0
	SliderMax = 1
)

func GetTitle() string {
	return "Settings"
}

func GetDesc() string {
	return "Settings page for the Computer Specs Viewer"
}

func CreateScreen(_ fyne.Window) fyne.CanvasObject {
	return utils.NewScrollVBox(createSections()...)
}

func createSections() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		utils.CreateSection("Theme:", createThemeToggle()),
	}
}

func createThemeToggle() fyne.CanvasObject {
	app := fyne.CurrentApp()

	darkThemeBtn := widget.NewButton("Dark Theme", func() {
		app.Settings().SetTheme(theme.DarkTheme())
	})
	lightThemeBtn := widget.NewButton("Light Theme", func() {
		app.Settings().SetTheme(theme.LightTheme())
	})

	return container.NewHBox(darkThemeBtn, lightThemeBtn)
}
