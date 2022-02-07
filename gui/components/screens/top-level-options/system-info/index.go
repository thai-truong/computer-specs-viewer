package systeminfo

import (
	"computer-specs-viewer/utils"

	"fyne.io/fyne/v2"
)

func GetTitle() string {
	return "System Information"
}

func GetDesc() string {
	return "Welcome to the Computer Specs Viewer program!"
}

func CreateScreen(_ fyne.Window) fyne.CanvasObject {
	textContent := []string{
		"Computer Specs Viewer can be used to see information about different parts of your computer, which includes:",
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
