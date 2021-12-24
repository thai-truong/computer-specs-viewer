package diskinfo

import (
	"computer-specs-viewer/src"
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

type DiskPartitionInfo struct {
	Name        string
	Mountpoint  string
	Fstype      string
	Opts        string
	TotalSpace  uint64
	FreeSpace   uint64
	UsedSpace   uint64
	TotalInodes uint64
	FreeInodes  uint64
	UsedInodes  uint64
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

func getFreeUsedPercents(total uint64, free uint64, used uint64) (freePercent float64, usedPercent float64) {
	totalF := float64(total)
	freeF := float64(free)
	usedF := float64(used)

	if totalF == 0 {
		return 0, 0
	}

	freePercent = freeF / totalF
	usedPercent = usedF / totalF

	return freePercent, usedPercent
}

func PrintAllDiskPartitionsInfo() {
	printDiskPartitionsInfo(true)
}

func PrintOnlyPhysDiskPartitionsInfo() {
	printDiskPartitionsInfo(false)
}

func printDiskPartitionsInfo(getAll bool) {
	dpsInfo := GetDiskPartitionsInfo(getAll)

	src.PrintSectionTitle("(All) Disk Partitions")
	src.PrintStartBlock()

	for i := 0; i < len(dpsInfo); i++ {
		printSingleDiskPartitionInfo(dpsInfo[i])

		if i < len(dpsInfo)-1 {
			src.PrintInfoDelim()
		}
	}

	src.PrintEndBlock()
}

func printSingleDiskPartitionInfo(dpi DiskPartitionInfo) {
	freePercent, usedPercent := getFreeUsedPercents(dpi.TotalSpace, dpi.FreeSpace, dpi.UsedSpace)
	freeInodesPercent, usedInodesPercent := getFreeUsedPercents(dpi.TotalInodes, dpi.FreeInodes, dpi.UsedInodes)

	fmt.Printf("Disk partition name: %v\n", dpi.Name)
	fmt.Printf("Mountpoint: %v\n", dpi.Mountpoint)
	fmt.Printf("File system type: %v\n", dpi.Fstype)
	fmt.Printf("Disk operations: %v\n", dpi.Opts)
	fmt.Printf("Total disk space: %v\n", src.GetSpaceString(dpi.TotalSpace, "GB"))
	fmt.Printf("Free disk space: %v\n", src.GetSpaceString(dpi.FreeSpace, "GB"))
	fmt.Printf("Free disk space percentage: %v\n", src.GetPercentString(freePercent))
	fmt.Printf("Used disk space: %v\n", src.GetSpaceString(dpi.UsedSpace, "GB"))
	fmt.Printf("Used disk space percentage: %v\n", src.GetPercentString(usedPercent))
	fmt.Printf("Total disk inodes: %v\n", dpi.TotalInodes)
	fmt.Printf("Free disk inodes: %v\n", dpi.FreeInodes)
	fmt.Printf("Free disk inodes percentage: %v\n", src.GetPercentString(freeInodesPercent))
	fmt.Printf("Used disk inodes: %v\n", dpi.UsedInodes)
	fmt.Printf("Used disk inodes percentage: %v\n", src.GetPercentString(usedInodesPercent))
}
