package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/expiteRz/graphical-mkw-bmg-editor-go/utils"
	"github.com/sqweek/dialog"
	"log"
)

const appNameFormat = "Easy BMG Editor - %s"
const (
	SimpleEdit   = 0
	AdvancedEdit = 1
)

type EditModeStringSet string

const (
	Simple   EditModeStringSet = "Simple"
	Advanced EditModeStringSet = "Advanced"
)

var a fyne.App
var w fyne.Window
var appWidth float32 = 640
var editMode = SimpleEdit
var editModeString = Simple

func init() {
	a = app.New()
	a.Settings().SetTheme(&MKFTheme{})
	w = a.NewWindow("Easy BMG Editor")
	w.Resize(fyne.NewSize(appWidth, 530))
}

func main() {
	w.SetMainMenu(newMenu())
	w.SetContent(render())
	w.ShowAndRun()
}

func render() *fyne.Container {
	return container.New(layout.NewPaddedLayout(),
		msgListLayout)
}

func newMenu() *fyne.MainMenu {
	return &fyne.MainMenu{Items: []*fyne.Menu{
		{
			Label: "File",
			Items: []*fyne.MenuItem{
				{
					Label:  "Open",
					Action: openFile,
				},
			},
		},
	}}
}

func openFile() {
	// Fyne implemented file dialog is weird
	// Use native dialog instead
	//dialog.ShowFileOpen(func(closer fyne.URIReadCloser, err error) {
	//    if closer == nil {
	//        return
	//    }
	//
	//    utils.Filepath = closer.URI().Path()
	//
	//    if err = utils.ParseBmg(); err != nil {
	//        log.Println("Failed to parse bmg")
	//        return
	//    }
	//
	//    utils.S = utils.CharsetString[utils.H.Charset]
	//    label.SetText(fmt.Sprintf("Charset: %s", utils.S))
	//    w.SetTitle(fmt.Sprintf(appNameFormat, utils.Filepath))
	//}, w)

	load, err := dialog.File().Filter("BMG file (.bmg)", "bmg").Load()
	if err != nil {
		return
	}
	utils.Filepath = load

	if err = utils.ParseBmg(); err != nil {
		log.Println("Failed to parse bmg")
		return
	}

	utils.S = utils.CharsetString[utils.H.Charset]
	w.SetTitle(fmt.Sprintf(appNameFormat, utils.Filepath))
}

func editModeToggle() {
	switch editMode {
	case SimpleEdit:
		editMode = AdvancedEdit
		editModeString = Advanced
	case AdvancedEdit:
		editMode = SimpleEdit
		editModeString = Simple
	default:
		log.Println("Edit Mode accidentally has unsupported mode. Reset to Simple mode.")
		editMode = SimpleEdit
		editModeString = Simple
	}
}
