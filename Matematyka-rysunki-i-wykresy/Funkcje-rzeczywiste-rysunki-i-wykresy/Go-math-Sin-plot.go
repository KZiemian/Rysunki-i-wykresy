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
	mathSinFunPlotData := make(plotter.XYs, 601)

	x := 0.0
	deltaX := math.Pi / 300

	for i := range mathSinFunPlotData {
		mathSinFunPlotData[i].X = x
		mathSinFunPlotData[i].Y = math.Sin(x)

		x += deltaX
	}



	plotFun := plot.New()

	plotFun.Title.Text = "Wykres funkcji math.Sin(x)"

	plotFun.X.Label.Text = "x"
	plotFun.Y.Label.Text = "y"

	err := plotutil.AddLinePoints(plotFun, "math.Sin",
		mathSinFunPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Sin_plot.png"); err != nil {

		panic(err)
	}
}
