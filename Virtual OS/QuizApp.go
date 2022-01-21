package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var Questions [5]string
var options [20]string
var currOptions [4]string
var corrAns [5]string

var currQue string
var opA string
var opB string
var opC string
var opD string
var correctAns string

var i int
var j int
var score int

var queLabel fyne.Widget
var opABtn fyne.Widget
var opBBtn fyne.Widget
var opCBtn fyne.Widget
var opDBtn fyne.Widget

func nextQuestion(queLabel *widget.Label, selected string, w fyne.Window) {
	fmt.Println("I am in")
	if selected == corrAns[i] {
		score++
	}

	i++
	j += 4
	if i >= len(Questions) {
		fmt.Println("in if block")
		msg := widget.NewLabel("Test is complete.")
		totalQue := "Total Question Was = " + strconv.Itoa(len(Questions))
		totalQueMsg := widget.NewLabel(totalQue)
		res := "Your Score is = " + strconv.Itoa(score)
		resMsg := widget.NewLabel(res)
		var per int
		per = (100 * score) / len(Questions)
		perStr := "Credit Percentage = " + strconv.Itoa(per)
		perMsg := widget.NewLabel(perStr)

		w.SetContent(container.NewVBox(
			msg, totalQueMsg, resMsg, perMsg,
		))

	} else {

		fmt.Println("in else block")

		currQue = Questions[i]
		opA = options[j]
		opB = options[j+1]
		opC = options[j+2]
		opD = options[j+3]

		queLabel.SetText(currQue)
		currOptions[0] = opA
		currOptions[1] = opB
		currOptions[2] = opC
		currOptions[3] = opD
	}
}

func QuizApp(a fyne.App) {
	// a := app.New()
	w := a.NewWindow("Quiz App")

	Questions[0] = "What are the advantages of GO?"
	options[0] = "GO compiles very quickly"
	options[1] = "Go supports concurrency at the language level"
	options[2] = "Functions are firstclass objects in GO"
	options[3] = "All of these"
	corrAns[0] = "All of these"

	Questions[1] = "Which of the following is initial value (zero value) for interfaces, slice, pointers, maps, channels and functions?"
	options[4] = "0"
	options[5] = "''"
	options[6] = "Nil"
	options[7] = "false"
	corrAns[1] = "Nil"

	Questions[2] = "Which of the following is false about Golang?"
	options[8] = "It is designed by Google."
	options[9] = "It is statically typed programming language."
	options[10] = "Golang is syntactically is similar to JAVA."
	options[11] = "'Testbook' is valid declaration."
	corrAns[2] = "Golang is syntactically is similar to JAVA."

	Questions[3] = "What are the several Built-in support in Go?"
	options[12] = "Container"
	options[13] = "Web Server"
	options[14] = "Database"
	options[15] = "All of these"
	corrAns[3] = "All of these"

	Questions[4] = "In Golang which of the following transfers control to the labeled statement"
	options[16] = "enum"
	options[17] = "goto"
	options[18] = "jump"
	options[19] = "return"
	corrAns[4] = "goto"

	//UI part begins
	i = 0
	j = 0
	score = 0
	currQue = Questions[i]
	opA = options[j]
	opB = options[j+1]
	opC = options[j+2]
	opD = options[j+3]

	queLabel := widget.NewLabel(currQue)
	currOptions[0] = opA
	currOptions[1] = opB
	currOptions[2] = opC
	currOptions[3] = opD

	var selected string
	selected = ""
	combo := widget.NewSelect(currOptions[:], func(value string) {
		selected = value
	})

	nextBtn := widget.NewButton("Save and Next", func() {
		combo.Selected = "(Select one)"
		if selected != "" && selected != "(Select one)" {
			nextQuestion(queLabel, selected, w)
		}
	})

	//display
	w.Resize(fyne.NewSize(600, 600))

	w.SetContent(container.NewVBox(
		queLabel, combo, nextBtn,
	),
	)
	w.Show()
}
