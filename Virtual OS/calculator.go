package main

// https://githubmemory.com/repo/fyne-io/developer.fyne.io/issues/12
import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalculator(a fyne.App) {
	//a := app.New()
	w := a.NewWindow("Calculator")
	show := ""
	showWithHistory := ""
	historyToggle := false
	outputScreen := widget.NewLabel(show)
	var historyAry []string

	//create all components
	historyBtn := widget.NewButton("History", func() {
		if !historyToggle {
			for i := 0; i < len(historyAry); i++ {
				showWithHistory += historyAry[i] + "\n"
			}
			showWithHistory += "\n" + "\n" + show
			outputScreen.SetText(showWithHistory)
			historyToggle = true
		} else {
			outputScreen.SetText(show)
			historyToggle = false
			showWithHistory = ""
		}

	})
	backBtn := widget.NewButton("Delete", func() {
		if len(show) > 0 {
			show = show[:len(show)-1]
			outputScreen.SetText(show)
		} else {
			show = ""
			outputScreen.SetText(show)
		}
	})

	clearBtn := widget.NewButton("All Clear", func() {
		show = ""
		outputScreen.SetText(show)
	})
	openBtn := widget.NewButton("(", func() {
		show += "("
		outputScreen.SetText(show)
	})
	closeBtn := widget.NewButton(")", func() {
		show += ")"
		outputScreen.SetText(show)
	})
	divBtn := widget.NewButton("/", func() {
		show += "/"
		outputScreen.SetText(show)
	})

	sevenBtn := widget.NewButton("7", func() {
		show += "7"
		outputScreen.SetText(show)
	})
	eightBtn := widget.NewButton("8", func() {
		show += "8"
		outputScreen.SetText(show)
	})
	nineBtn := widget.NewButton("9", func() {
		show += "9"
		outputScreen.SetText(show)
	})
	mulBtn := widget.NewButton("x", func() {
		show += "*"
		outputScreen.SetText(show)
	})

	fourBtn := widget.NewButton("4", func() {
		show += "4"
		outputScreen.SetText(show)
	})
	fiveBtn := widget.NewButton("5", func() {
		show += "5"
		outputScreen.SetText(show)
	})
	sixBtn := widget.NewButton("6", func() {
		show += "6"
		outputScreen.SetText(show)
	})
	minusBtn := widget.NewButton("-", func() {
		show += "-"
		outputScreen.SetText(show)
	})

	oneBtn := widget.NewButton("1", func() {
		show += "1"
		outputScreen.SetText(show)
	})
	twoBtn := widget.NewButton("2", func() {
		show += "2"
		outputScreen.SetText(show)
	})
	threeBtn := widget.NewButton("3", func() {
		show += "3"
		outputScreen.SetText(show)
	})
	plusBtn := widget.NewButton("+", func() {
		show += "+"
		outputScreen.SetText(show)
	})

	zeroBtn := widget.NewButton("0", func() {
		show += "0"
		outputScreen.SetText(show)
	})
	dotBtn := widget.NewButton(".", func() {
		show += "."
		outputScreen.SetText(show)
	})

	equalBtn := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(show)
		currEquationWithAns := ""
		currEquationWithAns += show + " = "
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				show = strconv.FormatFloat(result.(float64), 'f', -1, 64)
				currEquationWithAns += show
			} else {
				show = "Error in expression"
				currEquationWithAns += show
			}
		} else {
			show = "Error in expression"
			currEquationWithAns += show
		}
		historyAry = append(historyAry, currEquationWithAns)
		outputScreen.SetText(show)
	})

	calculator := container.NewVBox(
		outputScreen,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn, clearBtn,
			),
			container.NewGridWithColumns(4,
				backBtn, openBtn, closeBtn, divBtn,
			),
			container.NewGridWithColumns(4,
				sevenBtn, eightBtn, nineBtn, mulBtn,
			),
			container.NewGridWithColumns(4,
				fourBtn, fiveBtn, sixBtn, minusBtn,
			),
			container.NewGridWithColumns(4,
				oneBtn, twoBtn, threeBtn, plusBtn,
			),
			container.NewGridWithColumns(3,
				zeroBtn, dotBtn, equalBtn,
			),
		),
	)

	w.SetContent(
		container.NewBorder(calculator, nil, nil, nil),
	)

	w.Resize(fyne.NewSize(500, 0))
	w.Show()
}
