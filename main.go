package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/expiteRz/graphical-mkw-bmg-editor-go/utils"
	"log"
)

const appNameFormat = "Easy BMG Editor - %s"

var a fyne.App
var w fyne.Window

func init() {
	a = app.New()
	w = a.NewWindow("Easy BMG Editor")
	w.Resize(fyne.Size{
		Width:  520,
		Height: 600,
	})
}

func main() {
	box := container.NewVBox(
		label,
		msgElement(),
	)
	w.SetMainMenu(newMenu())
	w.SetContent(box)
	w.ShowAndRun()
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
	dialog.ShowFileOpen(func(closer fyne.URIReadCloser, err error) {
		if closer == nil {
			return
		}

		filepath = closer.URI().Path()

		if err = parseBmg(); err != nil {
			log.Println("Failed to parse bmg")
			return
		}

		s = utils.CharsetString[header.Charset]
		label.SetText(fmt.Sprintf("Charset: %s", s))
		w.SetTitle(fmt.Sprintf(appNameFormat, filepath))
	}, w)
}
