package netgui

import (
	netinfo "computer-specs-viewer/src/net_info"
	"computer-specs-viewer/utils"
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type NetworkInterfaceInformationGui struct {
	BaseInfo NetworkInterfaceBaseInformationGui
	IO       NetworkInterfaceIOGui
}

type NetworkInterfaceBaseInformationGui struct {
	OrderNumber  int
	MTU          int
	Name         string
	HardwareAddr string
	Flags        string
	Addresses    string
}

type NetworkInterfaceIOGui struct {
	BytesSent       uint64
	BytesReceived   uint64
	PacketsSent     uint64
	PacketsReceived uint64
	SendErrorCount  uint64
	RecvErrorCount  uint64
	SendDropCount   uint64
	RecvDropCount   uint64
}

func GetTitle() string {
	return "Network"
}

func GetDesc() string {
	return "This page contains information about this computer's network interfaces."
}

func transformNetworkIoInput(io netinfo.NetworkInterfaceIO) NetworkInterfaceIOGui {
	return NetworkInterfaceIOGui{
		BytesSent:       io.BytesSent,
		BytesReceived:   io.BytesRecv,
		PacketsSent:     io.PacketsSent,
		PacketsReceived: io.PacketsRecv,
		SendErrorCount:  io.SendErrCount,
		RecvErrorCount:  io.RecvErrCount,
		SendDropCount:   io.SendDropCount,
		RecvDropCount:   io.RecvDropCount,
	}
}

func transformNetworkBaseInfoInput(nwInterface netinfo.NetworkInterfaceInformation) NetworkInterfaceBaseInformationGui {
	return NetworkInterfaceBaseInformationGui{
		OrderNumber:  nwInterface.OrderNumber,
		MTU:          nwInterface.MTU,
		Name:         nwInterface.Name,
		HardwareAddr: nwInterface.HardwareAddr,
		Flags:        utils.StrListToPrettyStr(nwInterface.Flags),
		Addresses:    utils.StrListToPrettyStr(nwInterface.Addresses),
	}
}

func transformInput(nwInterface netinfo.NetworkInterfaceInformation) NetworkInterfaceInformationGui {
	return NetworkInterfaceInformationGui{
		BaseInfo: transformNetworkBaseInfoInput(nwInterface),
		IO:       transformNetworkIoInput(nwInterface.IOInfo),
	}
}

func getNetworkBaseInfoStrings(basicInfoGui NetworkInterfaceBaseInformationGui) []string {
	res := []string{}
	diskValues := reflect.ValueOf(basicInfoGui)
	diskType := diskValues.Type()

	for i := 0; i < diskValues.NumField(); i++ {
		infoFieldName := utils.SpaceOutFieldNames(diskType.Field(i).Name)
		infoFieldValue := diskValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func getNetworkBaseInfoLabel(basicInfoGui NetworkInterfaceBaseInformationGui) fyne.CanvasObject {
	return utils.SliceToSingleFyneLabel(getNetworkBaseInfoStrings(basicInfoGui))
}

func getNetworkIoInfoStrings(ioGui NetworkInterfaceIOGui) []string {
	res := []string{}
	diskValues := reflect.ValueOf(ioGui)
	diskType := diskValues.Type()

	for i := 0; i < diskValues.NumField(); i++ {
		infoFieldName := utils.SpaceOutFieldNames(diskType.Field(i).Name)
		infoFieldValue := diskValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func getNetworkIoInfoLabel(ioGui NetworkInterfaceIOGui) fyne.CanvasObject {
	return utils.SliceToSingleFyneLabel(getNetworkIoInfoStrings(ioGui))
}

func CreateInfoScreen(_ fyne.Window) fyne.CanvasObject {
	nwInterfaces := netinfo.GetNetworkInterfacesInfo()
	nwInterfaceAccordion := widget.NewAccordion()

	for _, nwInterface := range nwInterfaces {
		nwInterfaceGui := transformInput(nwInterface)

		nwBaseInfoTitle := widget.NewLabel("Base Information:")
		nwBaseInfoTitle.TextStyle.Bold = true
		nwBaseInfoLabel := getNetworkBaseInfoLabel(nwInterfaceGui.BaseInfo)

		nwIoInfoTitle := widget.NewLabel("IO Information:")
		nwIoInfoTitle.TextStyle.Bold = true
		nwIoInfoLabel := getNetworkIoInfoLabel(nwInterfaceGui.IO)

		nwInterfaceOrder := fmt.Sprint(nwInterfaceGui.BaseInfo.OrderNumber)
		nwInterfaceLabels := []fyne.CanvasObject{nwBaseInfoTitle, nwBaseInfoLabel, nwIoInfoTitle, nwIoInfoLabel}
		nwInterfaceAccordionItem := utils.CreateAccordionItem("Network Interface", nwInterfaceOrder, nwInterfaceLabels)
		nwInterfaceAccordion.Append(nwInterfaceAccordionItem)
	}

	return utils.NewScrollVBox(nwInterfaceAccordion)
}
