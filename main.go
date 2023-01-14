package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var sortBy string
	var dirPath string

	application := app.New()
	window := application.NewWindow("Organizador de Arquivos")
	window.Resize(fyne.NewSize(600, 450))

	title := canvas.NewText("Como deseja organizar os arquivos?", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 16

	textWidget := widget.NewLabel("Selecione a pasta")
	textWidget.Alignment = fyne.TextAlignCenter
	textWidget.Wrapping = fyne.TextWrapWord

	selectFolderButton := widget.NewButton(
		"Selecionar pasta",
		func() {
			dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
				if err != nil {
					textWidget.Text = err.Error()
				} else {
					if uri != nil {
						fmt.Println(uri.Path())
						textWidget.Text = uri.Path()
						textWidget.Show()
					}
				}

			}, window)
		},
	)

	okButton := widget.NewButton(
		"OK",
		func() {
			err := organizeFiles(sortBy, dirPath)

			if err != nil {
				fmt.Println(err.Error())
			}
		},
	)

	radio := widget.NewRadioGroup(
		[]string{
			"Por nome",
			"Por data",
		},
		func(value string) {
			if value == "Por nome" {
				sortBy = "-n"
			} else {
				sortBy = "-d"
			}
		},
	)

	buttonHBox := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		selectFolderButton,
		okButton,
		layout.NewSpacer(),
	)

	radioHBox := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		radio,
		layout.NewSpacer(),
	)

	VBox := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		title,
		radioHBox,
		buttonHBox,
		textWidget,
		layout.NewSpacer(),
	)

	window.SetContent(VBox)
	window.ShowAndRun()
}
