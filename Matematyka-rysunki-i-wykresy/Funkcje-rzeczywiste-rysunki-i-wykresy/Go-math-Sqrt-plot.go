package main

import (
	// "fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	mathSqrtFunPlotData := make(plotter.XYs, 201)

	x := 0.0
	deltaX := 2.0 / 200

	for i := range mathSqrtFunPlotData {
		mathSqrtFunPlotData[i].X = x
		mathSqrtFunPlotData[i].Y = math.Sqrt(x)

		x += deltaX
	}



	plotFun := plot.New()

	plotFun.Title.Text = "Wykres funkcji math.Sqrt(x)"

	plotFun.X.Label.Text = "x"
	plotFun.Y.Label.Text = "y"

	err := plotutil.AddLinePoints(plotFun, "math.Sqrt",
		mathSqrtFunPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Sqrt_plot_01.png"); err != nil {

		panic(err)
	}





	mathSqrtFunPlotData = make(plotter.XYs, 501)

	x = 0.0
	deltaX = 5.0 / 500

	for i := range mathSqrtFunPlotData {
		mathSqrtFunPlotData[i].X = x
		mathSqrtFunPlotData[i].Y = math.Sqrt(x)

		x += deltaX
	}



	plotFun = plot.New()

	plotFun.Title.Text = "Wykres funkcji math.Sqrt(x)"

	plotFun.X.Label.Text = "x"
	plotFun.Y.Label.Text = "y"

	err = plotutil.AddLinePoints(plotFun, "math.Sqrt",
		mathSqrtFunPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Sqrt_plot_02.png"); err != nil {

		panic(err)
	}





	mathSqrtFunPlotData = make(plotter.XYs, 1_001)

	x = 0.0
	deltaX = 10.0 / 1_001.0

	for i := range mathSqrtFunPlotData {
		mathSqrtFunPlotData[i].X = x
		mathSqrtFunPlotData[i].Y = math.Sqrt(x)

		x += deltaX
	}



	plotFun = plot.New()

	plotFun.Title.Text = "Wykres funkcji math.Sqrt(x)"

	plotFun.X.Label.Text = "x"
	plotFun.Y.Label.Text = "y"

	err = plotutil.AddLinePoints(plotFun, "math.Sqrt",
		mathSqrtFunPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Sqrt_plot_03.png"); err != nil {

		panic(err)
	}
}
