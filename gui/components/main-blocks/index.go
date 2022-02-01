package main_blocks

import (
	index_tree "computer-specs-viewer/gui/components/index-tree"
	"computer-specs-viewer/gui/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	splitOffset   = 0.3
	titleInitText = "Title"
	descInitText  = "Description"
)

func CreateMainViewScreen(window fyne.Window) fyne.CanvasObject {
	infoTitle := widget.NewLabel(titleInitText)
	infoDescription := widget.NewLabel(descInitText)
	infoContent := container.NewMax()

	setSelectedInfo := func(info data.TreeNodeContent) {
		infoTitle.SetText(info.Title)
		infoDescription.SetText(info.Description)

		infoContent.Objects = []fyne.CanvasObject{info.DisplayContent(window)}
		infoContent.Refresh()
	}

	navTree := index_tree.CreateIndexTree(setSelectedInfo)
	mainScreen := container.NewBorder(container.NewVBox(infoTitle, widget.NewSeparator(), infoDescription, widget.NewSeparator()), nil, nil, nil, infoContent)

	splitScreens := container.NewHSplit(navTree, mainScreen)
	splitScreens.Offset = splitOffset

	return splitScreens
}
