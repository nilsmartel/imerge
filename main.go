package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func showHelpWindow(app fyne.App) {
	w := app.NewWindow("Help")
	helpText := "Use IMerge to blend multiple images evenly into one.\nSelect the folder in which all your Images are stored,\nthen press 'Go'.\nAll images in that directory will be merged into one"

	w.SetContent(
		widget.NewVBox(
			widget.NewGroup(
				"Help",
				widget.NewLabel(helpText),
			),
			widget.NewHBox(
				widget.NewButton("Ok", func() {
					w.Close()
				}),
			),
		),
	)
	w.CenterOnScreen()
	w.Show()
}

func main() {
	app := app.New()

	quit := func() { app.Quit() }

	//  TODO merge Image Funciton
	mergeImages := quit

	//  TODO Show Files
	showFiles := quit

	// TODO Select directory funciotn
	selectDirectory := quit

	w := app.NewWindow("IMerge")

	directoryLabel := widget.NewLabel("Directory")

	amountLabel := widget.NewLabel("Amount: 0")

	dimensionLabel := widget.NewLabel("Dimension:\n width: 0px\n height: 0px")

	w.SetContent(
		widget.NewVBox(
			widget.NewLabel("IMerge"),
			// TODO maybe a little space (~16px) here?
			directoryLabel,
			widget.NewHBox(
				widget.NewButton("Select", selectDirectory),
				widget.NewButton("Show", showFiles),
			),

			widget.NewGroup(
				"Info",
				widget.NewVBox(
					amountLabel,
					dimensionLabel,
				),
			),

			widget.NewButton("Go", mergeImages),
			widget.NewHBox(
				widget.NewButton("Quit", quit),
				widget.NewButton("Help", func() { showHelpWindow(app) }),
			),
		))

	w.ShowAndRun()
}
