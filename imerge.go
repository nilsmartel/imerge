package main

import (
	_ "image/jpeg"
	_ "image/png"
	"strconv"

	"io/ioutil"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/sqweek/dialog"
)

func main() {
	application := app.New()
	labels := struct {
		directory *widget.Label
		amount    *widget.Label
		dimension *widget.Label
	}{
		widget.NewLabel("Directory"),
		widget.NewLabel("Amount: 0"),
		widget.NewLabel("Dimension:\n width: 0px\n height: 0px"),
	}
	directory := ""
	var images []FileInfo

	setDirectory := func(dir string) {
		{
			directoryFiles, err := ioutil.ReadDir(dir)

			if err != nil {
				showDialog(application, "Error", "Failed to read Directory "+dir)
				return
			}

			directory = dir
			labels.directory.SetText(CutOff(directory, 32))

			images = make([]FileInfo, 0)
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
					images = append(images, FileInfo{file.Name(), directory + "/" + file.Name()})
				}
			}

			labels.amount.SetText("Amount: " + strconv.Itoa(len(images)))

			if len(images) == 0 {
				return
			}

			labels.dimension.SetText("Loading Image Information")

			go func() {
				image := readImage(images[0].Path)

				width := image.Bounds().Max.X
				height := image.Bounds().Max.Y
				labels.dimension.SetText("Dimension:\n width: " + strconv.Itoa(width) + "px\n height: " + strconv.Itoa(height) + "px")
			}()
		}
	}

	quit := application.Quit

	//  TODO merge Image Function
	mergeImages := quit

	showFiles := func() {
		if len(images) == 0 {
			if directory == "" {
				showDialog(application, "No Directory", "No directory selected.\nPlease specify directory first.")
			} else {
				showDialog(application, "No Images", "No Images in selected Directory")
			}
		}

		label := ""
		for _, v := range images {
			label += v.Name + "\n"
		}

		content :=
			widget.NewScrollContainer(
				widget.NewVBox(
					widget.NewLabel(label),
				),
			)
		content.Resize(fyne.NewSize(256, 256))

		w := application.NewWindow("Images to blend")
		w.SetContent(
			content,
		)

		w.CenterOnScreen()
		w.Show()
	}

	selectDirectory := func() {
		directory, err := dialog.Directory().Title("Select Images").Browse()
		// err might be nil in case user clicks 'cancel'. In that case he won't even be surprised, if it's not working
		if err == nil {
			setDirectory(directory)
		}
	}

	w := application.NewWindow("IMerge")

	w.SetContent(
		widget.NewVBox(
			widget.NewLabel("IMerge"),
			// TODO maybe a little space (~16px) here?
			labels.directory,
			widget.NewHBox(
				widget.NewButton("Select", selectDirectory),
				widget.NewButton("Show", showFiles),
			),

			widget.NewGroup(
				"Info",
				widget.NewVBox(
					labels.amount,
					labels.dimension,
				),
			),

			widget.NewButton("Go", mergeImages),
			widget.NewHBox(
				widget.NewButton("Quit", quit),
				widget.NewButton("Help", func() { showHelpWindow(application) }),
			),
		))

	w.ShowAndRun()
}
