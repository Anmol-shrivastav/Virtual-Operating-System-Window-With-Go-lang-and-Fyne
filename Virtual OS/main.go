package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var a fyne.App = app.New()
var w fyne.Window = a.NewWindow("Virtual OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget

var desktopBtn fyne.Widget

var panelContent *fyne.Container

func main() {

	a.Settings().SetTheme(theme.DarkTheme())

	img := canvas.NewImageFromFile("pic.jpg")
	//img.FillMode = canvas.ImageFillOriginal
	img.Resize(fyne.NewSize(1280, 720))

	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		go showWeatherApp(a)
	})

	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		go showCalculator(a)
	})

	btn3 = widget.NewButtonWithIcon("Text Editor", theme.HomeIcon(), func() {
		go showTextEditor(a)
	})

	btn4 = widget.NewButtonWithIcon("Gallery App", theme.StorageIcon(), func() {
		go showGalleryAPP(a)
	})

	btn5 = widget.NewButtonWithIcon("Quiz App", theme.ComputerIcon(), func() {
		go QuizApp(a)
	})

	btn6 = widget.NewButtonWithIcon("Tansen Audio Player", theme.MediaMusicIcon(), func() {
		go tansenApp(a)
	})

	desktopBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		w.SetContent(
			container.NewBorder(panelContent, nil, nil, nil, img),
		)
	})

	panelContent = container.NewVBox(
		container.NewGridWithColumns(5,
			desktopBtn, btn1, btn2, btn3, btn4, btn5, btn6,
		),
	)

	w.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)

	w.Resize(fyne.NewSize(1280, 720))
	w.CenterOnScreen()
	w.ShowAndRun()
}
