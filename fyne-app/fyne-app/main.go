package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Test")
	w.Resize(fyne.NewSize(600, 400))

	label1 := widget.NewLabel("Hello World")

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter text...")

	label2 := widget.NewButton("Click Me", func() {
		label1.SetText("Button Clicked")
		data := entry.Text
		fmt.Println(data)
		label1.SetText(data)
	})

	w.SetContent(container.NewVBox(
		label1,
		label2,
		entry,
	))
	w.ShowAndRun()
}
