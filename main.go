package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))
	myWindow.Resize(fyne.NewSize(400, 500))

	// btn := widget.NewButton("Button", func() { fmt.Println("button") })
	btn1 := widget.NewButton("Button", func() { fmt.Println("button_1") })
	// btn2 := widget.NewButton("Button", func() { fmt.Println("button_2") })

	myWindow.SetContent(btn1)

	myWindow.Show()
	myApp.Run()
}
