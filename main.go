package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func showHelp(app fyne.App) {
	helpText := "Use IMerge to blend multiple images evenly into one. First, select the folder in which all your Images are stored. Press 'Go' and all images in that directory will be merged into one"
	w := app.NewWindow("Help")

	w.SetContent(widget.NewLabel(helpText))

	w.ShowAndRun()
}

func main() {
	app := app.New()

	w := app.NewWindow("IMerge")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("IMerge"),
		widget.NewHBox(
			widget.NewButton("Quit", func() {
				app.Quit()
			}),
			widget.NewButton("Help", func() {
				showHelp(app)
			}),
		),
	))

	w.ShowAndRun()
}
