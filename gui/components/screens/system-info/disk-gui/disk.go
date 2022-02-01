package diskgui

import custom_types "computer-specs-viewer/gui/custom-types"

type DiskPartitionInfoGui struct {
	Name              string
	Mountpoint        string
	Fstype            string
	Opts              string
	TotalSpace        custom_types.Space
	FreeSpace         custom_types.Space
	UsedSpace         custom_types.Space
	FreePercent       custom_types.Percent
	UsedPercent       custom_types.Percent
	TotalInodes       uint64
	FreeInodes        uint64
	UsedInodes        uint64
	FreeInodesPercent custom_types.Percent
	UsedInodesPercent custom_types.Percent
}

func transformInput() {}
