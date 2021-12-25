package netinfo

import "github.com/shirou/gopsutil/net"

type NetworkInterfaceInformation struct {
	OrderNumber  int
	MTU          int
	Name         string
	HardwareAddr string
	Flags        []string
	Addresses    []string
	IOInfo       NetworkInterfaceIO
}

type NetworkInterfaceIO struct {
	BytesSent     uint64
	BytesRecv     uint64
	PacketsSent   uint64
	PacketsRecv   uint64
	SendErrCount  uint64
	RecvErrCount  uint64
	SendDropCount uint64
	RecvDropCount uint64
}

func GetNetworkInterfacesInfo() []NetworkInterfaceInformation {
	interfaces, err := net.Interfaces()
	interfacesInfo := []NetworkInterfaceInformation{}

	if err != nil {
		return []NetworkInterfaceInformation{}
	}

	interfaceIOMap := createInterfaceIOMap()

	for i := 0; i < len(interfaces); i++ {
		currInterface := extractSingleNetworkInterfaceInfo(interfaces[i], &interfaceIOMap)
		ioInfo, interfaceFound := interfaceIOMap[interfaces[i].Name]

		if interfaceFound {
			currInterface.IOInfo = ioInfo
		}

		interfacesInfo = append(interfacesInfo, currInterface)
	}

	return interfacesInfo
}

func extractSingleNetworkInterfaceInfo(intf net.InterfaceStat, interfaceIOMap *map[string]NetworkInterfaceIO) NetworkInterfaceInformation {
	var currIOInfo NetworkInterfaceIO

	addrStrings := []string{}

	for i := 0; i < len(intf.Addrs); i++ {
		addrStrings = append(addrStrings, intf.Addrs[i].Addr)
	}

	ioInfo, foundInterface := (*interfaceIOMap)[intf.Name]

	if foundInterface {
		currIOInfo = ioInfo
	}

	return NetworkInterfaceInformation{
		OrderNumber:  intf.Index,
		MTU:          intf.MTU,
		Name:         intf.Name,
		HardwareAddr: intf.HardwareAddr,
		Flags:        intf.Flags,
		Addresses:    addrStrings,
		IOInfo:       currIOInfo,
	}
}

func createInterfaceIOMap() map[string]NetworkInterfaceIO {
	interfacesIO, err := net.IOCounters(true)
	interfaceIOMap := make(map[string]NetworkInterfaceIO)

	if err != nil {
		return make(map[string]NetworkInterfaceIO)
	}

	for i := 0; i < len(interfacesIO); i++ {
		currIO := interfacesIO[i]
		interfaceIOMap[currIO.Name] = extractSingleNetworkInterfaceIO(currIO)
	}

	return interfaceIOMap
}

func extractSingleNetworkInterfaceIO(ioInfo net.IOCountersStat) NetworkInterfaceIO {
	return NetworkInterfaceIO{
		BytesSent:     ioInfo.BytesSent,
		BytesRecv:     ioInfo.BytesRecv,
		PacketsSent:   ioInfo.PacketsSent,
		PacketsRecv:   ioInfo.PacketsRecv,
		SendErrCount:  ioInfo.Errout,
		RecvErrCount:  ioInfo.Errin,
		SendDropCount: ioInfo.Dropout,
		RecvDropCount: ioInfo.Dropin,
	}
}
