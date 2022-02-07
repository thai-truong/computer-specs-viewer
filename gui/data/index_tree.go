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
	"computer-specs-viewer/gui/components/screens/welcome"

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
		"settings":   {},
	}

	ContentMapping = map[string]TreeNodeContent{
		"welcome":     {welcome.GetTitle(), welcome.GetDesc(), welcome.CreateScreen},
		"systemInfo":  {systeminfo.GetTitle(), systeminfo.GetDesc(), systeminfo.CreateScreen},
		"cpu":         {cpugui.GetTitle(), cpugui.GetDesc(), cpugui.CreateInfoScreen},
		"disk":        {diskgui.GetTitle(), diskgui.GetDesc(), diskgui.CreateInfoScreen},
		"gpu":         {gpugui.GetTitle(), gpugui.GetDesc(), gpugui.CreateInfoScreen},
		"host":        {hostgui.GetTitle(), hostgui.GetDesc(), hostgui.CreateInfoScreen},
		"mem":         {memgui.GetTitle(), memgui.GetDesc(), memgui.CreateInfoScreen},
		"motherboard": {motherboardgui.GetTitle(), motherboardgui.GetDesc(), motherboardgui.CreateInfoScreen},
		"net":         {netgui.GetTitle(), netgui.GetDesc(), netgui.CreateInfoScreen},
		"settings":    {settings.GetTitle(), settings.GetDesc(), settings.CreateScreen},
	}
)

func GetWelcomeScreenData() TreeNodeContent {
	return ContentMapping["welcome"]
}
