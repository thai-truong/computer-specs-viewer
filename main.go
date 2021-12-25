package main

import (
	cpuinfo "computer-specs-viewer/src/cpu_info"
	diskinfo "computer-specs-viewer/src/disk_info"
	hostinfo "computer-specs-viewer/src/host_info"
	meminfo "computer-specs-viewer/src/mem_info"
	"fmt"
	"net"
)

func main() {
	//diskPartitions, _ := disk.Partitions(true)
	//cpuInfo, _ := cpu.Info()
	//hostInfo, _ := host.Info()
	//swapMem, _ := mem.SwapMemory()
	//swapDevices, _ := mem.SwapDevices()
	//virtualMem, _ := mem.VirtualMemory()
	nwInterfaceStats, _ := net.Interfaces()

	//fmt.Println(diskPartitions)
	//fmt.Println(cpuInfo)
	cpuinfo.PrintCpusInfo()
	diskinfo.PrintAllDiskPartitionsInfo()
	hostinfo.PrintHostInfo()
	meminfo.PrintAllMemInfo()
	//fmt.Println(hostInfo)
	//fmt.Println(swapMem)
	//fmt.Println(virtualMem)
	//fmt.Println(swapDevices)
	fmt.Println(nwInterfaceStats)
}
