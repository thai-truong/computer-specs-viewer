package diskinfo

import (
	"computer-specs-viewer/utils"
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

type DiskPartitionInfo struct {
	Name              string
	Mountpoint        string
	Fstype            string
	Opts              string
	TotalSpace        uint64
	FreeSpace         uint64
	UsedSpace         uint64
	FreePercent       float64
	UsedPercent       float64
	TotalInodes       uint64
	FreeInodes        uint64
	UsedInodes        uint64
	FreeInodesPercent float64
	UsedInodesPercent float64
}

func createDiskPartitionInfo(diskPt disk.PartitionStat) DiskPartitionInfo {
	mountPt := diskPt.Mountpoint
	diskInfo := DiskPartitionInfo{
		Name:       diskPt.Device,
		Mountpoint: diskPt.Mountpoint,
		Fstype:     diskPt.Fstype,
		Opts:       diskPt.Opts,
	}

	diskUsage, err := disk.Usage(mountPt)

	if err != nil {
		return diskInfo
	}

	diskInfo.TotalSpace = diskUsage.Total
	diskInfo.FreeSpace = diskUsage.Free
	diskInfo.UsedSpace = diskUsage.Used
	diskInfo.TotalInodes = diskUsage.InodesTotal
	diskInfo.FreeInodes = diskUsage.InodesFree
	diskInfo.UsedInodes = diskUsage.InodesUsed

	freeSpacePerc, usedSpacePerc := utils.GetFreeUsedPercents(diskUsage.Total, diskUsage.Free, diskUsage.Used)
	freeInodesPerc, usedInodesPerc := utils.GetFreeUsedPercents(diskUsage.InodesTotal, diskUsage.InodesFree, diskUsage.InodesUsed)

	diskInfo.FreePercent = freeSpacePerc
	diskInfo.UsedPercent = usedSpacePerc
	diskInfo.FreeInodesPercent = freeInodesPerc
	diskInfo.UsedInodesPercent = usedInodesPerc

	return diskInfo
}

func GetDiskPartitionsInfo(getAll bool) []DiskPartitionInfo {
	partitions, err := disk.Partitions(getAll)
	dpsInfo := []DiskPartitionInfo{}

	if err != nil {
		return []DiskPartitionInfo{}
	}

	for i := 0; i < len(partitions); i++ {
		dpsInfo = append(dpsInfo, createDiskPartitionInfo(partitions[i]))
	}

	return dpsInfo
}

func PrintAllDiskPartitionsInfo() {
	printDiskPartitionsInfo(true)
}

func PrintOnlyPhysDiskPartitionsInfo() {
	printDiskPartitionsInfo(false)
}

func printDiskPartitionsInfo(getAll bool) {
	dpsInfo := GetDiskPartitionsInfo(getAll)

	utils.PrintSectionTitle("(All) Disk Partitions")
	utils.PrintStartBlock()

	for i := 0; i < len(dpsInfo); i++ {
		printSingleDiskPartitionInfo(dpsInfo[i])

		if i < len(dpsInfo)-1 {
			utils.PrintInfoDelim()
		}
	}

	utils.PrintEndBlock()
}

func printSingleDiskPartitionInfo(dpi DiskPartitionInfo) {
	fmt.Printf("Disk partition name: %v\n", dpi.Name)
	fmt.Printf("Mountpoint: %v\n", dpi.Mountpoint)
	fmt.Printf("File system type: %v\n", dpi.Fstype)
	fmt.Printf("Disk operations: %v\n", dpi.Opts)
	fmt.Printf("Total disk space: %v\n", utils.GetSpaceString(dpi.TotalSpace, "GB"))
	fmt.Printf("Free disk space: %v\n", utils.GetSpaceString(dpi.FreeSpace, "GB"))
	fmt.Printf("Free disk space percentage: %v\n", utils.GetPercentString(dpi.FreePercent))
	fmt.Printf("Used disk space: %v\n", utils.GetSpaceString(dpi.UsedSpace, "GB"))
	fmt.Printf("Used disk space percentage: %v\n", utils.GetPercentString(dpi.UsedPercent))
	fmt.Printf("Total disk inodes: %v\n", dpi.TotalInodes)
	fmt.Printf("Free disk inodes: %v\n", dpi.FreeInodes)
	fmt.Printf("Free disk inodes percentage: %v\n", utils.GetPercentString(dpi.FreeInodesPercent))
	fmt.Printf("Used disk inodes: %v\n", dpi.UsedInodes)
	fmt.Printf("Used disk inodes percentage: %v\n", utils.GetPercentString(dpi.UsedInodesPercent))
}
