package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var label = widget.NewLabel(fmt.Sprintf("Charset: %s", s))

func msgElement() *widget.List {
	list := widget.NewList(
		func() int {
			return int(indexTable.Entries)
		},
		func() fyne.CanvasObject {
			c := container.NewVBox(
				widget.NewEntry())

			return c
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*fyne.Container).Objects[0].(*widget.Entry).SetText(strconv.Itoa(int(indexTable.MsgEntries[id].Offset)))
		},
	)

	return list
}
