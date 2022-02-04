package motherboardgui

import (
	motherboardinfo "computer-specs-viewer/src/motherboard_info"
	"computer-specs-viewer/utils"
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MotherboardInformationGui motherboardinfo.MotherboardInformation

func getMotherboardInfoStrings(mBoardGui MotherboardInformationGui) []string {
	res := []string{}
	cpuValues := reflect.ValueOf(mBoardGui)
	cpuType := cpuValues.Type()

	for i := 0; i < cpuValues.NumField(); i++ {
		infoFieldName := utils.SpaceOutFieldNames(cpuType.Field(i).Name)
		infoFieldValue := cpuValues.Field(i).Interface()

		infoStr := fmt.Sprintf("%s: %v\n", infoFieldName, infoFieldValue)
		res = append(res, infoStr)
	}

	return res
}

func getMotherboardInfoLabel(mBoardGui MotherboardInformationGui) fyne.CanvasObject {
	return utils.SliceToSingleFyneLabel(getMotherboardInfoStrings(mBoardGui))
}

func CreateInfoScreen(_ fyne.Window) fyne.CanvasObject {
	motherboards := motherboardinfo.GetMotherboardsInfo()
	mBoardAccordion := widget.NewAccordion()

	for i, motherboard := range motherboards {
		mBoardGui := MotherboardInformationGui(motherboard)

		mBoardInfoLabel := getMotherboardInfoLabel(mBoardGui)
		mBoardOrder := fmt.Sprint(i + 1)

		mBoardAccordionItem := utils.CreateAccordionItem("Motherboard", mBoardOrder, []fyne.CanvasObject{mBoardInfoLabel})
		mBoardAccordion.Append(mBoardAccordionItem)
	}

	return utils.NewScrollVBox(mBoardAccordion)
}
