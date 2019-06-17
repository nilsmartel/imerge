package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func showHelpWindow(app fyne.App) {
	w := app.NewWindow("Help")
	helpText := `Use IMerge to blend multiple images evenly into one.
Select the folder in which all your Images are stored,
then press 'Go'.
All images in that directory will be merged into one`

	w.SetContent(
		widget.NewVBox(
			widget.NewGroup(
				"Help",
				widget.NewLabel(helpText),
			),
			widget.NewHBox(
				widget.NewButton("Ok", w.Close),
			),
		),
	)
	w.CenterOnScreen()
	w.Show()
}

func showDialog(app fyne.App, title, message string) {
	w := app.NewWindow(title)
	w.SetContent(widget.NewLabel(message))
	w.CenterOnScreen()
	w.Show()
}

