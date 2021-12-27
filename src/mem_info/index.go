package meminfo

import (
	"computer-specs-viewer/utils"
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type SwapDeviceInfo struct {
	Name string
	Used uint64
	Free uint64
}

type MemInformation struct {
	TotalSpace  uint64
	UsedSpace   uint64
	FreeSpace   uint64
	UsedPercent float64
	FreePercent float64
}

type AllMemInfo struct {
	SwapDevices []SwapDeviceInfo
	SwapMem     MemInformation
	VirtualMem  MemInformation
}

func GetSwapDevicesInfo() []SwapDeviceInfo {
	swapDevices, err := mem.SwapDevices()
	swapDevicesInfo := []SwapDeviceInfo{}

	if err != nil {
		return []SwapDeviceInfo{}
	}

	for i := 0; i < len(swapDevices); i++ {
		swapDevicesInfo = append(swapDevicesInfo, extractSwapDeviceInfo(*swapDevices[i]))
	}

	return swapDevicesInfo
}

func GetAllMemInfo() AllMemInfo {
	return AllMemInfo{
		SwapDevices: GetSwapDevicesInfo(),
		SwapMem:     GetSwapMemInfo(),
		VirtualMem:  GetVirtualMemInfo(),
	}
}

func extractSwapDeviceInfo(device mem.SwapDevice) SwapDeviceInfo {
	return SwapDeviceInfo{
		Name: device.Name,
		Used: device.UsedBytes,
		Free: device.FreeBytes,
	}
}

func GetSwapMemInfo() MemInformation {
	origSwapInfo, err := mem.SwapMemory()

	if err != nil {
		return MemInformation{}
	}

	return extractSwapMemInfo(*origSwapInfo)
}

func extractSwapMemInfo(swapMem mem.SwapMemoryStat) MemInformation {
	freeP, usedP := utils.GetFreeUsedPercents(swapMem.Total, swapMem.Free, swapMem.Used)

	return MemInformation{
		TotalSpace:  swapMem.Total,
		UsedSpace:   swapMem.Used,
		FreeSpace:   swapMem.Free,
		UsedPercent: usedP,
		FreePercent: freeP,
	}
}

func GetVirtualMemInfo() MemInformation {
	origVirtualInfo, err := mem.VirtualMemory()

	if err != nil {
		return MemInformation{}
	}

	return extractVirtualMemInfo(*origVirtualInfo)
}

func extractVirtualMemInfo(vMem mem.VirtualMemoryStat) MemInformation {
	freeP, usedP := utils.GetFreeUsedPercents(vMem.Total, vMem.Available, vMem.Used)

	return MemInformation{
		TotalSpace:  vMem.Total,
		UsedSpace:   vMem.Used,
		FreeSpace:   vMem.Available,
		UsedPercent: usedP,
		FreePercent: freeP,
	}
}

func PrintAllMemInfo() {
	allMemInfo := GetAllMemInfo()

	utils.PrintSectionTitle("Memory (Swap, Virtual)")
	utils.PrintStartBlock()

	printMemInfo(allMemInfo.SwapMem, "Swap")
	utils.PrintInfoDelim()

	printSwapDevicesInfo(allMemInfo.SwapDevices)
	utils.PrintInfoDelim()

	printMemInfo(allMemInfo.VirtualMem, "Virtual")
	utils.PrintEndBlock()
}

func printSwapDevicesInfo(devices []SwapDeviceInfo) {
	for i := 0; i < len(devices); i++ {
		utils.PrintStrWithOrder("Swap device", i)
		printSingleDeviceInfo(devices[i])

		if i < len(devices)-1 {
			utils.PrintItemDelim()
		}
	}
}

func printSingleDeviceInfo(device SwapDeviceInfo) {
	fmt.Printf("Name: %v\n", device.Name)
	fmt.Printf("Free space: %v\n", utils.GetSpaceString(device.Free, "GB"))
	fmt.Printf("Used space: %v\n", utils.GetSpaceString(device.Used, "GB"))
}

func printMemInfo(memInfo MemInformation, memType string) {
	fmt.Printf("%v memory:\n", memType)
	fmt.Printf("Total space: %v\n", utils.GetSpaceString(memInfo.TotalSpace, "GB"))
	fmt.Printf("Free space: %v\n", utils.GetSpaceString(memInfo.FreeSpace, "GB"))
	fmt.Printf("Used space: %v\n", utils.GetSpaceString(memInfo.UsedSpace, "GB"))
	fmt.Printf("Free space percent: %v\n", utils.GetPercentString(memInfo.FreePercent))
	fmt.Printf("Used space percent: %v\n", utils.GetPercentString(memInfo.UsedPercent))
}
