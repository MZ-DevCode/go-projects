package main

import (
	"fmt"
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main(){
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(600, 400))

	label1 := widget.NewLabel("First Number: ")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Second Number: ")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("")
	btn := widget.NewButton("Calculate", func(){
		n1, err := strconv.ParseFloat(entry1.Text, 64)
		n2, err := strconv.ParseFloat(entry2.Text, 64)
		if err != nil{
			answer.SetText("Invalid input")
		} else{

		sum := n1 + n2
		min := n1 - n2
		mul := n1 * n2
		div := n1 / n2

		answer.SetText(fmt.Sprintf("(+) %f\n(-) %f\n(*) %f\n(/) %f", sum, min, mul, div))
	}
	})

	w.SetContent(container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		btn,
		answer,
	))

	w.ShowAndRun()
}
