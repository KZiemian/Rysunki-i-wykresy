package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	mathAcoshFunPlotData := make(plotter.XYs, 301)

	x := 1.0
	deltaX := 1.0 / 100.0


	for i := range mathAcoshFunPlotData {
		mathAcoshFunPlotData[i].X = x
		mathAcoshFunPlotData[i].Y = math.Acosh(x)

		x += deltaX
	}



	plotOfFun := plot.New()

	plotOfFun.Title.Text = "Wykres funkcji math.Acosh(x)"

	plotOfFun.X.Label.Text = "x"
	plotOfFun.Y.Label.Text = "y"

	l, err := plotter.NewLine(mathAcoshFunPlotData)

	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFun.Add(l)
	plotOfFun.Legend.Add("math.Acosh", l)

	// pictureWidthInInches := 10.0
	// pictureHeightInInches := (pictureWidthInInches * mathAcoshFunPlotData[300].Y / mathAcoshFunPlotData[300].X)

	if err := plotOfFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Acosh_plot_01.png"); err != nil {

		panic(err)
	}




	mathAcoshFunPlotData = make(plotter.XYs, 1_001)

	x = 1.0

	for i := range mathAcoshFunPlotData {
		mathAcoshFunPlotData[i].X = x
		mathAcoshFunPlotData[i].Y = math.Acosh(x)

		x += deltaX
	}



	plotOfFun = plot.New()

	plotOfFun.Title.Text = "Wykres funkcji math.Acosh(x)"

	plotOfFun.X.Label.Text = "x"
	plotOfFun.Y.Label.Text = "y"

	l, err = plotter.NewLine(mathAcoshFunPlotData)

	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFun.Add(l)
	plotOfFun.Legend.Add("math.Acosh(x)", l)

	if err := plotOfFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Acosh_plot_02.png"); err != nil {

		panic(err)
	}
}
