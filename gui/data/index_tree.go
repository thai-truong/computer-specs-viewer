package data

import (
	cpugui "computer-specs-viewer/gui/components/screens/system-info/cpu-gui"
	diskgui "computer-specs-viewer/gui/components/screens/system-info/disk-gui"
	gpugui "computer-specs-viewer/gui/components/screens/system-info/gpu-gui"
	hostgui "computer-specs-viewer/gui/components/screens/system-info/host-gui"
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
		"cpu":         {"CPU", "This page contains information about this computer's CPU.", cpugui.CreateInfoScreen},
		"disk":        {"Disk", "This page contains information about this computer's disk partitions.", diskgui.CreateInfoScreen},
		"gpu":         {"GPU", "This page contains information about this computer's GPU.", gpugui.CreateInfoScreen},
		"host":        {"Host", "This page contains information about this computer's current host.", hostgui.CreateInfoScreen},
		"mem":         {"Memory", "", nil},
		"motherboard": {"Motherboard", "", nil},
		"net":         {"Network", "", nil},
		"settings":    {"Settings", "", top_level_options.SettingsScreen},
		"themeMode":   {"Theme Mode", "", nil},
	}
)
