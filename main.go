package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func newLabel(s string) *widget.Label {
	return widget.NewLabel(s)
}

func main() {
	a := app.New()
	w := a.NewWindow("Calc")

	hello := newLabel("X + Y = Z")
	input := widget.NewEntry()
	w.SetContent(container.NewGridWithRows(
		3,
		input,
		hello,
		newLabel("bye"),
	))

	input.OnChanged = func(input string) {
		vals := strings.Split(input, " ")
		if(len(vals) == 3) {
			for i:=0; i<len(vals[0]); i++ {
				fmt.Println(vals[0] + vals[2])
			}
		} else {
			fmt.Println(len(vals))
		}
	}

	w.ShowAndRun()
}
