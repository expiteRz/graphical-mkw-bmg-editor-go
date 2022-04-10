package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/expiteRz/graphical-mkw-bmg-editor-go/utils"
	"strconv"
)

var label = widget.NewLabel(fmt.Sprintf("Charset: %s", utils.S))
var msgListLayout = msgElementAdvanced()

func msgElementAdvanced() *widget.List {
	list := widget.NewList(
		func() int {
			return len(utils.I.MsgEntries)
		},
		func() fyne.CanvasObject {
			msgIdWidget, msgWidget := widget.NewLabel(""), widget.NewEntry()
			msgEscapeWidget := widget.NewSelect([]string{"None", "Test1", "Test2"}, func(s string) {})
			msgFontWidget := widget.NewSelect([]string{"Countdown/Finish strings", "Standard", "Red font", "Blue font"}, func(s string) {})
			msgWidget.SetPlaceHolder("Message")

			msgEscapeWidget.SetSelectedIndex(0)
			msgFontWidget.SetSelectedIndex(1)

			msgIDLayout := container.NewGridWrap(fyne.NewSize(90, 38), msgIdWidget)
			msgEscapeLayout := container.NewGridWrap(fyne.NewSize(160, 38), msgEscapeWidget)
			msgFontLayout := container.NewGridWrap(fyne.NewSize(180, 38), msgFontWidget)
			msgLayout := container.NewGridWrap(fyne.NewSize(appWidth-100, 38), msgWidget)

			return container.NewHBox(msgIDLayout, msgFontLayout, msgEscapeLayout, msgLayout)
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			m := object.(*fyne.Container)
			l1, l2 := m.Objects[0].(*fyne.Container), m.Objects[3].(*fyne.Container)
			l1.Objects[0].(*widget.Label).SetText(fmt.Sprintf("%x", utils.M.Ids[id]))
			l2.Objects[0].(*widget.Entry).SetText(strconv.Itoa(int(utils.I.MsgEntries[id].FontType)))
		},
	)

	return list
}

func msgElementSimple() *widget.List {
	return widget.NewList(func() int {
		return len(utils.I.MsgEntries)
	}, func() fyne.CanvasObject {
		msgIdWidget, msgWidget := widget.NewLabel(""), widget.NewEntry()
		msgWidget.SetPlaceHolder("Message")
		msgIdLayout, msgLayout :=
			container.NewGridWrap(fyne.NewSize(90, 38), msgIdWidget),
			container.NewGridWrap(fyne.NewSize(appWidth-100, 38), msgWidget)

		return container.NewHBox(msgIdLayout, msgLayout)
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		m := object.(*fyne.Container)
		l1, l2 := m.Objects[0].(*fyne.Container), m.Objects[1].(*fyne.Container)
		l1.Objects[0].(*widget.Label).SetText(fmt.Sprintf("%x", utils.I.MsgEntries[id].Offset))
		l2.Objects[0].(*widget.Entry).SetText(strconv.Itoa(int(utils.I.MsgEntries[id].FontType)))
	})
}
