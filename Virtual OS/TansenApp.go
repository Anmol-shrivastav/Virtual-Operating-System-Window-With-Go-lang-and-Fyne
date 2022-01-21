package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func tansenApp(a fyne.App) {
	// a := app.New()
	w := a.NewWindow("Tansen Audio Player")

	var format beep.Format
	var streamer beep.StreamSeekCloser
	var pause bool = false

	image := canvas.NewImageFromFile("logo.png")
	image.FillMode = canvas.ImageFillOriginal

	title := widget.NewLabelWithStyle("Lets Play Some Music...", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: true})

	browseContainer := widget.NewButtonWithIcon("Browse", theme.FileIcon(), func() {
		dialogBox := dialog.NewFileOpen(
			func(uc fyne.URIReadCloser, e error) {
				streamer, format, _ = mp3.Decode(uc)
				title.Text = uc.URI().Name()
				title.Refresh()
			}, w)

		dialogBox.Show()
	})

	controlBox := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if pause == false {
				pause = true
				speaker.Lock() //it will pause play music
			} else {
				pause = false
				speaker.Unlock() //it will play music
			}
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()
			title.SetText("Lets Play Some Music...")
			title.Refresh()
		}),
		widget.NewToolbarSpacer(),
	)

	//display
	w.SetContent(container.NewVBox(
		image, title, browseContainer, controlBox,
	))
	w.Resize(fyne.NewSize(300, 600))
	w.Show()
}
