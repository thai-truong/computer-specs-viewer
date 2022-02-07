package memgui

import (
	custom_types "computer-specs-viewer/gui/custom-types"
	meminfo "computer-specs-viewer/src/mem_info"
	"computer-specs-viewer/utils"
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type SwapDeviceInfoGui struct {
	Name string
	Used custom_types.Space
	Free custom_types.Space
}

type MemInformationGui struct {
	TotalSpace  custom_types.Space
	UsedSpace   custom_types.Space
	FreeSpace   custom_types.Space
	UsedPercent custom_types.Percent
	FreePercent custom_types.Percent
}

type AllMemInfoGui struct {
	SwapDevices []SwapDeviceInfoGui
	SwapMem     MemInformationGui
	VirtualMem  MemInformationGui
}

func GetTitle() string {
	return "Memory"
}

func GetDesc() string {
	return "This page contains information about this computer's memory information"
}

func transformSwapDevsInput(swapDevs []meminfo.SwapDeviceInfo) []SwapDeviceInfoGui {
	swapDevGuiSlice := []SwapDeviceInfoGui{}

	for _, swapDev := range swapDevs {
		swapDevGui := SwapDeviceInfoGui{
			Name: swapDev.Name,
			Used: custom_types.NumToCustomSpaceType(swapDev.Used).ToGb(),
			Free: custom_types.NumToCustomSpaceType(swapDev.Free).ToGb(),
		}
		swapDevGuiSlice = append(swapDevGuiSlice, swapDevGui)
	}

	return swapDevGuiSlice
}

func transformMemInput(mem meminfo.MemInformation) MemInformationGui {
	return MemInformationGui{
		TotalSpace:  custom_types.NumToCustomSpaceType(mem.TotalSpace).ToGb(),
		UsedSpace:   custom_types.NumToCustomSpaceType(mem.UsedSpace).ToGb(),
		FreeSpace:   custom_types.NumToCustomSpaceType(mem.FreeSpace).ToGb(),
		UsedPercent: custom_types.CreatePercent(mem.UsedPercent),
		FreePercent: custom_types.CreatePercent(mem.FreePercent),
	}
}

func transformInput(mem meminfo.AllMemInfo) AllMemInfoGui {
	return AllMemInfoGui{
		SwapDevices: transformSwapDevsInput(mem.SwapDevices),
		SwapMem:     transformMemInput(mem.SwapMem),
		VirtualMem:  transformMemInput(mem.VirtualMem),
	}
}

func getMemInfoStrings(memGui MemInformationGui) []string {
	res := []string{}
	diskValues := reflect.ValueOf(memGui)
	diskType := diskValues.Type()

	for i := 0; i < diskValues.NumField(); i++ {
		infoFieldName := utils.SpaceOutFieldNames(diskType.Field(i).Name)
		infoFieldValue := diskValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func getMemInfoLabel(memGui MemInformationGui) fyne.CanvasObject {
	return utils.SliceToSingleFyneLabel(getMemInfoStrings(memGui))
}

func getSwapDevInfoStrings(sDevGui SwapDeviceInfoGui) []string {
	res := []string{}
	diskValues := reflect.ValueOf(sDevGui)
	diskType := diskValues.Type()

	for i := 0; i < diskValues.NumField(); i++ {
		infoFieldName := utils.SpaceOutFieldNames(diskType.Field(i).Name)
		infoFieldValue := diskValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func getSwapDevInfoLabel(sDevsGui []SwapDeviceInfoGui) fyne.CanvasObject {
	allSwapDevs := []string{}

	for i, sDevGui := range sDevsGui {
		sDevOrder := fmt.Sprint(i + 1)
		currSwapDevStrs := []string{utils.GetStrWithOrder("Swap Device", sDevOrder) + ":"}

		currSwapDevStrs = append(currSwapDevStrs, getSwapDevInfoStrings(sDevGui)...)
		allSwapDevs = append(allSwapDevs, currSwapDevStrs...)
		allSwapDevs = append(allSwapDevs, "\n")
	}

	return utils.SliceToSingleFyneLabel(allSwapDevs)
}

func CreateInfoScreen(_ fyne.Window) fyne.CanvasObject {
	memInfo := meminfo.GetAllMemInfo()
	memAccordion := widget.NewAccordion()

	memInfoGui := transformInput(memInfo)
	vMemLabel := getMemInfoLabel(memInfoGui.VirtualMem)
	sMemLabel := getMemInfoLabel(memInfoGui.SwapMem)
	swapDevsLabel := getSwapDevInfoLabel(memInfoGui.SwapDevices)

	vMemAccordionItem := utils.CreateAccordionItem("Virtual Memory", "", []fyne.CanvasObject{vMemLabel})
	sMemAccordionItem := utils.CreateAccordionItem("Swap Memory", "", []fyne.CanvasObject{sMemLabel})
	swapDevsAccordionItem := utils.CreateAccordionItem("Swap Devices", "", []fyne.CanvasObject{swapDevsLabel})

	memAccordion.Append(sMemAccordionItem)
	memAccordion.Append(vMemAccordionItem)
	memAccordion.Append(swapDevsAccordionItem)

	return utils.NewScrollVBox(memAccordion)
}
