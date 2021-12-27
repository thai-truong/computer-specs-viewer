package main

import (
	cpuinfo "computer-specs-viewer/src/cpu_info"
	diskinfo "computer-specs-viewer/src/disk_info"
	gpuinfo "computer-specs-viewer/src/gpu_info"
	hostinfo "computer-specs-viewer/src/host_info"
	meminfo "computer-specs-viewer/src/mem_info"
	netinfo "computer-specs-viewer/src/net_info"
)

func main() {
	cpuinfo.PrintCpusInfo()
	diskinfo.PrintAllDiskPartitionsInfo()
	hostinfo.PrintHostInfo()
	meminfo.PrintAllMemInfo()
	netinfo.PrintNetworkInterfacesInfo()
	gpuinfo.PrintGpusInfo()
}
