package data

import (
	cpugui "computer-specs-viewer/gui/components/screens/system-info/cpu-gui"
	top_level_options "computer-specs-viewer/gui/components/screens/top-level-options"

	"fyne.io/fyne/v2"
)

type TreeNodeContent struct {
	Title          string
	Description    string
	DisplayContent func(w fyne.Window) fyne.CanvasObject
}

var (
	IndexTreeMapping = map[string][]string{
		"":           {"welcome", "systemInfo", "settings"},
		"systemInfo": {"cpu", "disk", "gpu", "host", "mem", "motherboard", "net"},
		"settings":   {"themeMode"},
	}

	TitleToContentMapping = map[string]TreeNodeContent{
		"welcome":     {"Welcome", "", nil},
		"systemInfo":  {"System Information", "", top_level_options.SystemInformationScreen},
		"cpu":         {"CPU", "This page contains information on this computer's CPU.", cpugui.CreateInfoScreen},
		"disk":        {"Disk", "", nil},
		"gpu":         {"GPU", "", nil},
		"host":        {"Host", "", nil},
		"mem":         {"Memory", "", nil},
		"motherboard": {"Motherboard", "", nil},
		"net":         {"Network", "", nil},
		"settings":    {"Settings", "", top_level_options.SettingsScreen},
		"themeMode":   {"Theme Mode", "", nil},
	}
)
