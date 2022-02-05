package welcome

import (
	"computer-specs-viewer/utils"

	"fyne.io/fyne/v2"
)

func GetTitle() string {
	return "Welcome Page"
}

func GetDesc() string {
	return "Welcome to the Computer Specs Viewer!"
}

func CreateScreen(_ fyne.Window) fyne.CanvasObject {
	textContent := []string{
		"Computer Specs Viewer (CSV) can be used to see information about different parts of your computer, which includes:",
		"\t - CPU",
		"\t - Disk Partitions",
		"\t - GPU",
		"\t - Host",
		"\t - Memory Devices",
		"\t - Motherboard",
		"\t - Network Interfaces",
		"\nWe only support retrieval of information about these parts as of now.",
		"Moreover, a few kinds of information only have Windows support, which includes GPU and Motherboard.",
	}
	welcomeLabel := utils.SliceToSingleFyneLabel(textContent)

	return utils.NewScrollVBox(welcomeLabel)
}
