package main

import (
	"io/ioutil"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func showTextEditor(a fyne.App) {
	var count int = 1
	// a := app.New()
	w := a.NewWindow("Notepad")

	heading := widget.NewLabel("Text Editor")

	content := container.NewVBox(
		container.NewHBox(
			heading,
		),
	)

	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File " + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Resize(fyne.NewSize(400, 400))

	saveBtn := widget.NewButton("Save", func() {
		saveDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, e error) {
				textData := []byte(input.Text)

				uc.Write(textData)
			}, w)

		saveDialog.SetFileName("New File " + strconv.Itoa(count))

		saveDialog.Show()
	})
	openBtn := widget.NewButton("Open File", func() {
		openDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, e error) {
				readData, err := ioutil.ReadAll(r)
				if err != nil {
					log.Fatal(err)
				}

				output := fyne.NewStaticResource("New File", readData)

				viewData := widget.NewMultiLineEntry()
				viewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName),
				)

				viewData.SetPlaceHolder("Enter text...")
				viewData.Resize(fyne.NewSize(400, 600))

				btn := widget.NewButton("Save", func() {
					saveDialog := dialog.NewFileSave(
						func(r fyne.URIWriteCloser, e error) {
							textData := []byte(viewData.Text)

							r.Write(textData)
						}, w)

					saveDialog.SetFileName("New File " + strconv.Itoa(count))

					saveDialog.Show()
				})

				//containerBox := container.New(layout.NewVBoxLayout(), text, layout.NewSpacer(), btn)
				//containerBox.Resize(fyne.NewSize(400, 400))

				w.SetContent(container.NewVBox(
					viewData, btn,
				))

				w.Resize(fyne.NewSize(400, 400))

				w.Show()
			}, w)

		openDialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}),
		)

		openDialog.Show()
	})

	w.SetContent(container.NewVBox(
		content, input,
		container.NewHBox(
			saveBtn, openBtn,
		),
	))

	w.Resize(fyne.NewSize(600, 600))
	w.Show()
}
