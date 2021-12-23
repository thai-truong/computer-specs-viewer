package main

import (
	"computer-specs-viewer/src/cpu_info"
	"computer-specs-viewer/src/disk_info"
)

func main() {
	//diskPartitions, _ := disk.Partitions(true)
	//cpuInfo, _ := cpu.Info()
	//hostInfo, _ := host.Info()
	//swapMem, _ := mem.SwapMemory()
	//swapDevices, _ := mem.SwapDevices()
	//virtualMem, _ := mem.VirtualMemory()
	//nwInterfaceStats, _ := net.Interfaces()

	//fmt.Println(diskPartitions)
	//fmt.Println(cpuInfo)
	cpu_info.PrintCpusInfo()
	disk_info.PrintAllDiskPartitionsInfo()
	//fmt.Println(hostInfo)
	//fmt.Println(swapMem)
	//fmt.Println(virtualMem)
	//fmt.Println(swapDevices)
	//fmt.Println(nwInterfaceStats)
}
