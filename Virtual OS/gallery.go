package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func showGalleryAPP(a fyne.App) {
	fmt.Println("Yes Iam here you are good to go")
	// a := app.New()
	w := a.NewWindow("Gallery App")
	var picAry []string
	show := "No Image Found!!!"
	output := widget.NewLabel(show)
	path := "C:\\Program Files\\Go\\src\\Go Projects\\Virtual OS//picture"

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err) //it will stop the program excution with printing the error message
	}

	for _, f := range files {
		if !f.IsDir() {
			filename := f.Name()
			fileExtenssion := strings.Split(filename, ".")[1]
			if fileExtenssion == "png" || fileExtenssion == "jpg" {
				picAry = append(picAry, filename)
			}
		}
	}

	if len(picAry) > 0 {
		tabs := container.NewAppTabs(
			container.NewTabItemWithIcon("Home", theme.HomeIcon(), canvas.NewImageFromFile("picture/"+picAry[0])),
		)

		for i := 1; i < len(picAry); i++ {
			filename := strings.Split(picAry[i], ".")[0]
			tabs.Append(container.NewTabItem(filename, canvas.NewImageFromFile("picture/"+picAry[i])))
		}
		w.SetContent(tabs)
	} else {
		w.SetContent(container.NewVBox(
			output,
		),
		)
	}

	w.Resize(fyne.NewSize(800, 600))
	w.Show()
}
