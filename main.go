package main

import (
	main_blocks "computer-specs-viewer/gui/components/main-blocks"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	viewerApp := app.NewWithID("computer.specs.viewer")
	viewerWindow := viewerApp.NewWindow("Computer Specs Viewer")

	viewerWindow.SetContent(main_blocks.CreateMainViewScreen(viewerWindow))

	viewerWindow.Resize(fyne.NewSize(640, 460))
	viewerWindow.ShowAndRun()
}
