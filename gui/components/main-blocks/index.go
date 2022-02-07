package main_blocks

import (
	indextree "computer-specs-viewer/gui/components/index-tree"
	"computer-specs-viewer/gui/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var splitOffset = 0.3

func getStartScreenContent(w fyne.Window) (screenTitle *widget.Label, screenDesc *widget.Label, screenContent *fyne.Container) {
	welcomeData := data.GetStartScreenData()

	screenTitle = widget.NewLabel(welcomeData.Title)
	screenDesc = widget.NewLabel(welcomeData.Description)

	screenContent = container.NewMax()
	screenContent.Objects = []fyne.CanvasObject{welcomeData.DisplayContent(w)}

	return screenTitle, screenDesc, screenContent
}

func CreateMainViewScreen(window fyne.Window) fyne.CanvasObject {
	screenTitle, screenDesc, screenContent := getStartScreenContent(window)

	setSelectedInfo := func(info data.TreeNodeContent) {
		screenTitle.SetText(info.Title)
		screenDesc.SetText(info.Description)

		screenContent.Objects = []fyne.CanvasObject{info.DisplayContent(window)}
		screenContent.Refresh()
	}

	navTree := indextree.CreateIndexTree(setSelectedInfo)
	navTree.OpenBranch(data.GetStartScreenName())
	mainScreen := container.NewBorder(container.NewVBox(screenTitle, widget.NewSeparator(), screenDesc, widget.NewSeparator()), nil, nil, nil, screenContent)

	splitScreens := container.NewHSplit(navTree, mainScreen)
	splitScreens.Offset = splitOffset

	return splitScreens
}
