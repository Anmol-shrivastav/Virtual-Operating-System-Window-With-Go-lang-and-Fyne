package main

// https://app.quicktype.io
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showWeatherApp(a fyne.App) {
	// a := app.New()
	w := a.NewWindow("Weather App")

	image := canvas.NewImageFromFile("pic.png")
	image.FillMode = canvas.ImageFillOriginal
	image.Resize(fyne.NewSize(500, 200))

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Your State Name....")
	stateName := ""

	country := widget.NewLabel("")
	windSpeed := widget.NewLabel("")
	state := widget.NewLabel("")
	weatherDescription := widget.NewLabel("")
	visibility := widget.NewLabel("")
	temperature := widget.NewLabel("")
	humidity := widget.NewLabel("")

	api := "https://api.openweathermap.org/data/2.5/weather?q="
	saveBtn := widget.NewButton("Get Weather", func() {
		api = "https://api.openweathermap.org/data/2.5/weather?q="
		stateName = input.Text
		api += stateName + "&appid=a1d7b9ef6a211e7dbc71a77562c5bbab"

		res, err := http.Get(api) //get response
		if err != nil {
			log.Fatalln("There is Some Error in https request, Can't get API Data")
		}
		defer res.Body.Close()
		//Read data
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln("There is Some Error in Reading API Data")
		}
		weather, err := UnmarshalWelcome(body)
		if err != nil {
			errorMessage := widget.NewLabel("City Name not Found, Please enter a valid City Name")
			w.SetContent(
				container.NewVBox(
					errorMessage,
				),
			)
		} else {
			countrStr := fmt.Sprintf("Country : %s", weather.Sys.Country)
			windSpeedStr := fmt.Sprintf("Wind Speed : %.1f", weather.Wind.Speed)
			stateStr := fmt.Sprintf("Location : %s", weather.Name)
			weatherDescriptionStr := fmt.Sprintf("Weather Description : %s", weather.Weather[0].Description)
			visibilityStr := fmt.Sprintf("Visiblity : %d", weather.Visibility)
			temperatureStr := fmt.Sprintf("Temperature : %.2f", weather.Main.Temp)
			humidityStr := fmt.Sprintf("Humidity : %d", weather.Main.Humidity)

			country.SetText(countrStr)
			windSpeed.SetText(windSpeedStr)
			state.SetText(stateStr)
			weatherDescription.SetText(weatherDescriptionStr)
			visibility.SetText(visibilityStr)
			temperature.SetText(temperatureStr)
			humidity.SetText(humidityStr)
		}
	})

	w.SetContent(container.NewVBox(
		image, input, saveBtn, country, windSpeed, state, weatherDescription, visibility, temperature, humidity,
	))

	//w.Resize(fyne.NewSize(600, 800))
	w.Show()
}

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
