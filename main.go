package main

import (
	"bufio"
	cpuinfo "computer-specs-viewer/src/cpu_info"
	diskinfo "computer-specs-viewer/src/disk_info"
	gpuinfo "computer-specs-viewer/src/gpu_info"
	hostinfo "computer-specs-viewer/src/host_info"
	meminfo "computer-specs-viewer/src/mem_info"
	motherboardinfo "computer-specs-viewer/src/motherboard_info"
	netinfo "computer-specs-viewer/src/net_info"
	"fmt"
	"os"
	"strings"
)

var printChoiceMapping = map[string](func()){
	"cpu":         cpuinfo.PrintCpusInfo,
	"disk":        diskinfo.PrintAllDiskPartitionsInfo,
	"gpu":         gpuinfo.PrintGpusInfo,
	"host":        hostinfo.PrintHostInfo,
	"memory":      meminfo.PrintAllMemInfo,
	"net":         netinfo.PrintNetworkInterfacesInfo,
	"motherboard": motherboardinfo.PrintMotherboardsInfo,
	"all":         printAllInfo,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printStartingInstructions()

		fmt.Print("\nEnter information listed above that you want to see: ")
		scanner.Scan()
		choice := scanner.Text()

		if foundQuitInput(choice) {
			break
		}

		printFuncToCall, foundChoice := printChoiceMapping[choice]

		if !foundChoice {
			fmt.Println(fmt.Errorf("\n\nError: Your input (%v) does not match any of the provided choices for input", choice))
			continue
		}

		printFuncToCall()
	}
}

func printAllInfo() {
	cpuinfo.PrintCpusInfo()
	diskinfo.PrintAllDiskPartitionsInfo()
	hostinfo.PrintHostInfo()
	meminfo.PrintAllMemInfo()
	netinfo.PrintNetworkInterfacesInfo()
	gpuinfo.PrintGpusInfo()
	motherboardinfo.PrintMotherboardsInfo()
}

func printStartingInstructions() {
	fmt.Println("\n- Type one of these choices to view their information: [cpu, disk, gpu, host, memory, net, motherboard]")
	fmt.Println("- To view all types of information, type \"all\"")
	fmt.Println("- To quit, type \"q\" or \"quit\"!")
}

func foundQuitInput(input string) bool {
	lowercaseInput := strings.ToLower(input)

	return lowercaseInput == "q" || lowercaseInput == "quit"
}
