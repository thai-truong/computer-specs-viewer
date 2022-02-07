package hostgui

import (
	hostinfo "computer-specs-viewer/src/host_info"
	"computer-specs-viewer/utils"
	"fmt"
	"reflect"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type HostInformationGui struct {
	Name               string
	Uptime             string
	BootTime           time.Time
	NumberOfProcesses  uint64
	OS                 string
	Platform           string
	PlatformFamily     string
	KernelVersion      string
	KernelArchitecture string
}

func GetTitle() string {
	return "Host"
}

func GetDesc() string {
	return "This page contains information about this computer's current host"
}

func transformInput(host hostinfo.HostInformation) HostInformationGui {
	return HostInformationGui{
		Name:               host.Name,
		Uptime:             fmt.Sprintf("%.2f hrs", utils.SecondsAmtToHours(host.Uptime)),
		BootTime:           host.BootTime,
		NumberOfProcesses:  host.NumProcs,
		OS:                 host.OS,
		Platform:           host.Platform,
		PlatformFamily:     host.PlatformFam,
		KernelVersion:      host.KernelVer,
		KernelArchitecture: host.KernelArch,
	}
}

func getHostInfoStrings(hostGui HostInformationGui) []string {
	res := []string{}
	cpuValues := reflect.ValueOf(hostGui)
	cpuType := cpuValues.Type()

	for i := 0; i < cpuValues.NumField(); i++ {
		infoFieldName := utils.SpaceOutFieldNames(cpuType.Field(i).Name)
		infoFieldValue := cpuValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func getHostInfoLabel(hostGui HostInformationGui) fyne.CanvasObject {
	return utils.SliceToSingleFyneLabel(getHostInfoStrings(hostGui))
}

func CreateInfoScreen(_ fyne.Window) fyne.CanvasObject {
	host := hostinfo.GetHostInfo()
	cpuAccordion := widget.NewAccordion()

	hostGui := transformInput(host)
	hostInfoLabel := getHostInfoLabel(hostGui)

	accordionItem := utils.CreateAccordionItem("Host", "", []fyne.CanvasObject{hostInfoLabel})
	cpuAccordion.Append(accordionItem)

	return utils.NewScrollVBox(cpuAccordion)
}
