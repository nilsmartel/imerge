package main

import (
	"io/ioutil"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/sqweek/dialog"
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

func cutOff(s string, length int) string {
	if len(s) < length {
		return s
	}
	return s[:length] + "..."
}

func main() {
	app := app.New()
	directoryLabel := widget.NewLabel("Directory")
	amountLabel := widget.NewLabel("Amount: 0")
	dimensionLabel := widget.NewLabel("Dimension:\n width: 0px\n height: 0px")

	directory := ""
	var images []string

	setDirectory := func(dir string) {
		{
			directoryFiles, err := ioutil.ReadDir(dir)

			if err != nil {
				showDialog(app, "Error", "Failed to read Directory "+dir)
				return
			}

			directory = dir
			directoryLabel.SetText(cutOff(directory, 32))

			images = make([]string, 0)
			isImage := func(s string) bool {
				switch s {
				case ".png", ".jpg", ".jpeg":
					return true
				default:
					return false
				}
			}

			for _, file := range directoryFiles {
				if !file.IsDir() && isImage(filepath.Ext(file.Name())) {
					images = append(images, file.Name())
				}
			}
		}
	}

	quit := app.Quit

	//  TODO merge Image Funciton
	mergeImages := quit

	showFiles := func() {
		label := ""
		for _, v := range images {
			label = label + "\n" + v
		}

		w := app.NewWindow("Images to blend")
		w.SetContent(
			widget.NewVBox(
				widget.NewLabel(label),
			),
		)

		w.CenterOnScreen()
		w.Show()
	}

	selectDirectory := func() {
		directory, err := dialog.Directory().Title("Select Images").Browse()
		// err might be nil in case uses clicks 'cancel'. In that case he woun't even be surprised, if it's not working
		if err == nil {
			showDialog(app, "Success", "Selected "+directory)
			setDirectory(directory)
		}
	}

	w := app.NewWindow("IMerge")

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
