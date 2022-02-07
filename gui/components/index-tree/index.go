package indextree

import (
	"computer-specs-viewer/gui/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var placeholderText = "Placeholder"

type InfoTreeTab struct {
	Title          string
	Intro          string
	DisplayContent func(w fyne.Window) fyne.CanvasObject
}

func CreateIndexTree(setSelectedInfo func(data.TreeNodeContent)) *widget.Tree {
	indexTree := widget.NewTree(getChildUIDs, isBranchNode, createNewNode, updateNode)
	setOnSelected(indexTree, setSelectedInfo)

	return indexTree
}

func getChildUIDs(uid string) []string {
	return data.IndexTreeMapping[uid]
}

func isBranchNode(uid string) bool {
	children, found := data.IndexTreeMapping[uid]

	return found && len(children) > 0
}

func createNewNode(isBranch bool) fyne.CanvasObject {
	placeholder := widget.NewLabel(placeholderText)

	return placeholder
}

func updateNode(uid string, isBranch bool, nodeContent fyne.CanvasObject) {
	t, ok := data.ContentMapping[uid]

	if !ok {
		fyne.LogError("Invalid system information requested: "+uid, nil)
		return
	}

	nodeContent.(*widget.Label).SetText(t.Title)
}

func setOnSelected(tree *widget.Tree, setSelectedInfo func(data.TreeNodeContent)) {
	tree.OnSelected = func(uid string) {
		info, found := data.ContentMapping[uid]

		if found {
			setSelectedInfo(info)
		}
	}
}
