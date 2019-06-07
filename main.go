package main

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	defer fmt.Println("Quitting App")

	fmt.Println("starting app")

	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	fmt.Println("show Window")
	w.ShowAndRun()

	fmt.Println("Window running")
}
