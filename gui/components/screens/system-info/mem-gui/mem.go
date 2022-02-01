package memgui

import custom_types "computer-specs-viewer/gui/custom-types"

type SwapDeviceInfoGUI struct {
	Name string
	Used custom_types.Space
	Free custom_types.Space
}

type MemInformationGUI struct {
	TotalSpace  custom_types.Space
	UsedSpace   custom_types.Space
	FreeSpace   custom_types.Space
	UsedPercent custom_types.Percent
	FreePercent custom_types.Percent
}

type AllMemInfo struct {
	SwapDevices []SwapDeviceInfoGUI
	SwapMem     MemInformationGUI
	VirtualMem  MemInformationGUI
}
