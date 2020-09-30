package main

import (
	"fmt"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))
	myWindow.Resize(fyne.NewSize(400, 500))

	btn := widget.NewButton("Button", func() { fmt.Println("button") })
	btn1 := widget.NewButton("Button", func() { fmt.Println("button_1") })
	btn2 := widget.NewButton("Button", func() { fmt.Println("button_2") })
	btn3 := widget.NewButtonWithIcon("this", theme.HomeIcon(), func() { log.Println("home tapped!") })
	btn4 := widget.NewButtonWithIcon("this", theme.ComputerIcon(), func() { log.Println("computer tapped!") })

	icon := widget.NewIcon(theme.HelpIcon())

	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("enter text here!")
	entry2 := widget.NewPasswordEntry()
	entry2.SetPlaceHolder("password hidden!")

	myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), btn, btn1, btn2, btn3, btn4, icon, entry1, entry2))

	myWindow.ShowAndRun()
}
