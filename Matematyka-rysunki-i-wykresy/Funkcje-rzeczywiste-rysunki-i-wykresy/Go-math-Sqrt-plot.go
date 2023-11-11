package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	mathSqrtFunPlotData := make(plotter.XYs, 201)

	x := 0.0
	deltaX := 1.0 / 100.0

	for i := range mathSqrtFunPlotData {
		mathSqrtFunPlotData[i].X = x
		mathSqrtFunPlotData[i].Y = math.Sqrt(x)

		x += deltaX
	}



	plotOfFun := plot.New()

	plotOfFun.Title.Text = "Wykres funkcji math.Sqrt(x)"

	plotOfFun.X.Label.Text = "x"
	plotOfFun.Y.Label.Text = "y"

	l, err := plotter.NewLine(mathSqrtFunPlotData)

	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFun.Add(l)
	plotOfFun.Legend.Add("math.Sqrt(x)", l)

	if err := plotOfFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Sqrt_plot_01.png"); err != nil {

		panic(err)
	}





	mathSqrtFunPlotData = make(plotter.XYs, 501)

	x = 0.0

	for i := range mathSqrtFunPlotData {
		mathSqrtFunPlotData[i].X = x
		mathSqrtFunPlotData[i].Y = math.Sqrt(x)

		x += deltaX
	}



	plotOfFun = plot.New()

	plotOfFun.Title.Text = "Wykres funkcji math.Sqrt(x)"

	plotOfFun.X.Label.Text = "x"
	plotOfFun.Y.Label.Text = "y"

	l, err = plotter.NewLine(mathSqrtFunPlotData)

	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFun.Add(l)
	plotOfFun.Legend.Add("math.Sqrt(x)", l)

	if err := plotOfFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Sqrt_plot_02.png"); err != nil {

		panic(err)
	}





	mathSqrtFunPlotData = make(plotter.XYs, 1_001)

	x = 0.0

	for i := range mathSqrtFunPlotData {
		mathSqrtFunPlotData[i].X = x
		mathSqrtFunPlotData[i].Y = math.Sqrt(x)

		x += deltaX
	}



	plotOfFun = plot.New()

	plotOfFun.Title.Text = "Wykres funkcji math.Sqrt(x)"

	plotOfFun.X.Label.Text = "x"
	plotOfFun.Y.Label.Text = "y"

	l, err = plotter.NewLine(mathSqrtFunPlotData)

	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFun.Add(l)
	plotOfFun.Legend.Add("math.Sqrt(x)", l)

	if err := plotOfFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Sqrt_plot_03.png"); err != nil {

		panic(err)
	}
}
