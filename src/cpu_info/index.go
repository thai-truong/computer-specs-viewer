package cpu_info

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"

	"computer-specs-viewer/src"
)

type CpuInformation struct {
	ModelName     string
	LogicalCores  int
	PhysicalCores int
	Mhz           float64
	PhysicalId    string
	VendorId      string
	Family        string
	OrderNumber   int
}

func GetCpuInfo() []CpuInformation {
	var existingCpus []CpuInformation

	cpuInfo, err := cpu.Info()

	if err != nil {
		return []CpuInformation{}
	}

	for i := 0; i < len(cpuInfo); i++ {
		currCpuInfo := extractSingleCpuInfo(cpuInfo[i])
		existingCpus = append(existingCpus, currCpuInfo)
	}

	return existingCpus
}

func extractSingleCpuInfo(cpuInfo cpu.InfoStat) CpuInformation {
	modelName := cpuInfo.ModelName
	vendorId := cpuInfo.VendorID
	family := cpuInfo.Family
	physId := cpuInfo.PhysicalID
	orderNumber := cpuInfo.CPU

	mhz := cpuInfo.Mhz
	logCoreCount, errLC := cpu.Counts(true)
	physCoreCount, errPC := cpu.Counts(false)

	if errLC != nil {
		logCoreCount = -1
	}

	if errPC != nil {
		physCoreCount = -1
	}

	return CpuInformation{
		ModelName:     modelName,
		VendorId:      vendorId,
		Family:        family,
		PhysicalId:    physId,
		Mhz:           mhz,
		LogicalCores:  logCoreCount,
		PhysicalCores: physCoreCount,
		OrderNumber:   int(orderNumber),
	}
}

func PrintCpusInfo() {
	cpus := GetCpuInfo()

	src.PrintSectionTitle("CPU")
	src.PrintStartBlock()

	for i := 0; i < len(cpus); i++ {
		printSingleCpuInfo(cpus[i])

		if i < len(cpus)-1 {
			src.PrintInfoDelim()
		}
	}

	src.PrintEndBlock()
}

func printSingleCpuInfo(cpu CpuInformation) {
	fmt.Printf("** CPU number %v **\n", cpu.OrderNumber)
	fmt.Printf("Model name: %v\n", cpu.ModelName)
	fmt.Printf("Vendor: %v\n", cpu.VendorId)
	fmt.Printf("Family: %v\n", cpu.Family)
	fmt.Printf("Physical ID: %v\n", cpu.PhysicalId)
	fmt.Printf("CPU speed: %v Mhz\n", cpu.Mhz)
	fmt.Printf("Number of logical cores: %v\n", cpu.LogicalCores)
	fmt.Printf("Number of physical cores: %v\n", cpu.PhysicalCores)
}
