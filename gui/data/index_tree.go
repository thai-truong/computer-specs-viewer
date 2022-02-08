package data

import (
	cpugui "computer-specs-viewer/gui/components/screens/system-info/cpu-gui"
	diskgui "computer-specs-viewer/gui/components/screens/system-info/disk-gui"
	gpugui "computer-specs-viewer/gui/components/screens/system-info/gpu-gui"
	hostgui "computer-specs-viewer/gui/components/screens/system-info/host-gui"
	memgui "computer-specs-viewer/gui/components/screens/system-info/mem-gui"
	motherboardgui "computer-specs-viewer/gui/components/screens/system-info/motherboard-gui"
	netgui "computer-specs-viewer/gui/components/screens/system-info/net-gui"
	"computer-specs-viewer/gui/components/screens/top-level-options/settings"
	systeminfo "computer-specs-viewer/gui/components/screens/top-level-options/system-info"

	"fyne.io/fyne/v2"
)

type TreeNodeContent struct {
	Title          string
	Description    string
	DisplayContent func(w fyne.Window) fyne.CanvasObject
}

var (
	IndexTreeMapping = map[string][]string{
		"":           {"systemInfo", "settings"},
		"systemInfo": {"cpu", "disk", "gpu", "host", "mem", "motherboard", "net"},
		"settings":   {},
	}

	ContentMapping = map[string]TreeNodeContent{
		"systemInfo":  {systeminfo.GetTitle(), systeminfo.GetDesc(), systeminfo.CreateScreen},
		"cpu":         {cpugui.GetTitle(), cpugui.GetDesc(), cpugui.CreateScreen},
		"disk":        {diskgui.GetTitle(), diskgui.GetDesc(), diskgui.CreateScreen},
		"gpu":         {gpugui.GetTitle(), gpugui.GetDesc(), gpugui.CreateScreen},
		"host":        {hostgui.GetTitle(), hostgui.GetDesc(), hostgui.CreateScreen},
		"mem":         {memgui.GetTitle(), memgui.GetDesc(), memgui.CreateScreen},
		"motherboard": {motherboardgui.GetTitle(), motherboardgui.GetDesc(), motherboardgui.CreateScreen},
		"net":         {netgui.GetTitle(), netgui.GetDesc(), netgui.CreateScreen},
		"settings":    {settings.GetTitle(), settings.GetDesc(), settings.CreateScreen},
	}
)

func GetStartScreenData() TreeNodeContent {
	return ContentMapping["systemInfo"]
}

func GetStartScreenName() string {
	return "systemInfo"
}
