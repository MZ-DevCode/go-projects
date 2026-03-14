package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Test")

	label1 := widget.NewLabel("Hello World")
	label2 := widget.NewButton("Click Me", func() {
		label1.SetText("Button Clicked")
	})
	w.SetContent(container.NewVBox(
		label1,
		label2,
	))
	w.ShowAndRun()
}
