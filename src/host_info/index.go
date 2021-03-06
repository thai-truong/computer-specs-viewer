package hostinfo

import (
	"computer-specs-viewer/utils"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/host"
)

type HostInformation struct {
	Name        string
	Uptime      uint64
	BootTime    time.Time
	NumProcs    uint64
	OS          string
	Platform    string
	PlatformFam string
	KernelVer   string
	KernelArch  string
}

func GetHostInfo() HostInformation {
	origHostInfo, err := host.Info()

	if err != nil {
		return HostInformation{}
	}

	return extractHostInformation(*origHostInfo)
}

func extractHostInformation(origHostInfo host.InfoStat) HostInformation {
	return HostInformation{
		Name:        origHostInfo.Hostname,
		Uptime:      origHostInfo.Uptime,
		BootTime:    time.Unix(int64(origHostInfo.BootTime), 0),
		NumProcs:    origHostInfo.Procs,
		OS:          origHostInfo.OS,
		Platform:    fmt.Sprintf("%v Version %v", origHostInfo.Platform, origHostInfo.PlatformVersion),
		PlatformFam: origHostInfo.PlatformFamily,
		KernelVer:   origHostInfo.KernelVersion,
		KernelArch:  origHostInfo.KernelArch,
	}
}

func PrintHostInfo() {
	hostInfo := GetHostInfo()

	utils.PrintSectionTitle("Host")
	utils.PrintStartBlock()

	printHostInfoDetails(hostInfo)

	utils.PrintEndBlock()
}

func printHostInfoDetails(hostInfo HostInformation) {
	fmt.Printf("Name: %v\n", hostInfo.Name)
	fmt.Printf("Amount of uptime: %v seconds\n", hostInfo.Uptime)
	fmt.Printf("Time of booting: %v\n", hostInfo.BootTime)
	fmt.Printf("Number of processes: %v\n", hostInfo.NumProcs)
	fmt.Printf("OS: %v\n", hostInfo.OS)
	fmt.Printf("Platform: %v\n", hostInfo.Platform)
	fmt.Printf("Platform family: %v\n", hostInfo.PlatformFam)
	fmt.Printf("Kernel Version: %v\n", hostInfo.KernelVer)
	fmt.Printf("Kernel Architecture: %v\n", hostInfo.KernelArch)
}
