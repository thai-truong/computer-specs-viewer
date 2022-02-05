package diskgui

import (
	custom_types "computer-specs-viewer/gui/custom-types"
	diskinfo "computer-specs-viewer/src/disk_info"
	"computer-specs-viewer/utils"
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

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

func GetTitle() string {
	return "Disk"
}

func GetDesc() string {
	return "This page contains information about this computer's disk partitions."
}

func transformInput(disk diskinfo.DiskPartitionInfo) DiskPartitionInfoGui {
	return DiskPartitionInfoGui{
		Name:              disk.Name,
		Mountpoint:        disk.Mountpoint,
		Fstype:            disk.Fstype,
		Opts:              disk.Opts,
		TotalSpace:        custom_types.NumToCustomSpaceType(disk.TotalSpace).ToGb(),
		FreeSpace:         custom_types.NumToCustomSpaceType(disk.FreeSpace).ToGb(),
		UsedSpace:         custom_types.NumToCustomSpaceType(disk.UsedSpace).ToGb(),
		FreePercent:       custom_types.CreatePercent(disk.FreePercent),
		UsedPercent:       custom_types.CreatePercent(disk.UsedPercent),
		TotalInodes:       disk.TotalInodes,
		FreeInodes:        disk.FreeInodes,
		UsedInodes:        disk.UsedInodes,
		FreeInodesPercent: custom_types.CreatePercent(disk.FreeInodesPercent),
		UsedInodesPercent: custom_types.CreatePercent(disk.UsedInodesPercent),
	}
}

func getDiskInfoStrings(diskGui DiskPartitionInfoGui) []string {
	res := []string{}
	diskValues := reflect.ValueOf(diskGui)
	diskType := diskValues.Type()

	for i := 0; i < diskValues.NumField(); i++ {
		infoFieldName := utils.SpaceOutFieldNames(diskType.Field(i).Name)
		infoFieldValue := diskValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func getDiskInfoLabel(diskGui DiskPartitionInfoGui) fyne.CanvasObject {
	return utils.SliceToSingleFyneLabel(getDiskInfoStrings(diskGui))
}

func CreateInfoScreen(_ fyne.Window) fyne.CanvasObject {
	diskSlice := diskinfo.GetDiskPartitionsInfo(true)
	diskAccordion := widget.NewAccordion()

	for _, disk := range diskSlice {
		diskGui := transformInput(disk)
		diskOrder := diskGui.Name[:len(diskGui.Name)-1]

		diskInfoLabel := getDiskInfoLabel(diskGui)
		accordionItem := utils.CreateAccordionItem("Disk", diskOrder, []fyne.CanvasObject{diskInfoLabel})
		diskAccordion.Append(accordionItem)
	}

	return utils.NewScrollVBox(diskAccordion)
}
