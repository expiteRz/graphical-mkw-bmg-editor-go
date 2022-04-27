package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	fDialog "fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"github.com/expiteRz/graphical-mkw-bmg-editor-go/utils"
	"github.com/expiteRz/graphical-mkw-bmg-editor-go/utils/reg"
	"github.com/sqweek/dialog"
	"log"
	"os"
)

const (
	appNameFormat = "Easy BMG Editor %s - %s"
	Version       = "Alpha 1.0"
)
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
	w = a.NewWindow(fmt.Sprintf("Easy BMG Editor %s", Version))
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
					Label:  "New",
					Action: utils.InitBmg,
				},
				{
					Label:  "Open",
					Action: openFile,
				},
				{
					IsSeparator: true,
				},
				{
					Label: "Save",
					Action: func() {
						fDialog.ShowInformation("Save", "Save function is not implemented yet.", w)
					},
				},
				{
					Label:  "Save as new",
					Action: saveFile,
				},
			},
		},
		{
			Label: "Help",
			Items: []*fyne.MenuItem{
				{
					Label:  "About",
					Action: showAbout,
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

	load, err := dialog.File().Filter("BMG file (.bmg)", "bmg").SetStartDir(reg.ReadFileDir()).Load()
	if err != nil {
		return
	}
	utils.Filepath = load
	_ = reg.SetFileDir(load)

	if err = utils.ParseBmg(); err != nil {
		log.Println("Failed to parse bmg")
		return
	}

	utils.S = utils.CharsetString[utils.H.Charset]
	w.SetTitle(fmt.Sprintf(appNameFormat, Version, utils.Filepath))
	msgListLayout.Refresh()
}

func saveFile() {
	save, err := dialog.File().Filter("BMG File (*.bmg)", "bmg").SetStartDir(reg.ReadFileDir()).Save()
	if err != nil {
		fDialog.ShowError(err, w)
		return
	}

	if err = reg.SetFileDir(save); err != nil {
		return
	}

	bmg, err := utils.CombineBmg()
	if err != nil {
		log.Printf("Error: %v", err)
		fDialog.ShowError(err, w)
		return
	}

	file, err := os.Create(save)
	if err != nil {
		fDialog.ShowError(err, w)
		return
	}

	if _, err = file.Write(bmg.Bytes()); err != nil {
		fDialog.ShowError(err, w)
		return
	}
}

func showAbout() {
	fDialog.NewInformation("About", fmt.Sprintf("Easy BMG Editor %s\n\nGUI designed with Fyne\nDocumentation from Custom Mario Kart Wiiki\nAny knowledges from Wiimms SZS Tools by Wiimm\n and CTools by Chadderz", Version), w).Show()
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
