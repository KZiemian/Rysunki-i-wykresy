package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
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



	plotFun := plot.New()

	plotFun.Title.Text = "Wykres funkcji math.Acosh(x)"

	plotFun.X.Label.Text = "x"
	plotFun.Y.Label.Text = "y"

	err := plotutil.AddLinePoints(plotFun, "math.Acosh",
		mathAcoshFunPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotFun.Save(10*vg.Inch, 10*vg.Inch,
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



	plotFun = plot.New()

	plotFun.Title.Text = "Wykres funkcji math.Acosh(x)"

	plotFun.X.Label.Text = "x"
	plotFun.Y.Label.Text = "y"

	err = plotutil.AddLinePoints(plotFun, "math.Acosh",
		mathAcoshFunPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Acosh_plot_02.png"); err != nil {

		panic(err)
	}
}
