package motherboardgui

import (
	motherboardinfo "computer-specs-viewer/src/motherboard_info"
	"computer-specs-viewer/utils"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MotherboardInformationGui motherboardinfo.MotherboardInformation

func GetTitle() string {
	return "Motherboard"
}

func GetDesc() string {
	return "This page contains information about this computer's motherboards"
}

func CreateScreen(_ fyne.Window) fyne.CanvasObject {
	motherboards := motherboardinfo.GetMotherboardsInfo()
	mBoardAccordion := widget.NewAccordion()

	for i, motherboard := range motherboards {
		mBoardGui := MotherboardInformationGui(motherboard)

		mBoardInfoLabel := utils.GetInfoGuiLabel(mBoardGui)
		mBoardOrder := fmt.Sprint(i + 1)

		mBoardAccordionItem := utils.CreateAccordionItem("Motherboard", mBoardOrder, []fyne.CanvasObject{mBoardInfoLabel})
		mBoardAccordion.Append(mBoardAccordionItem)
	}

	return utils.NewScrollVBox(mBoardAccordion)
}
