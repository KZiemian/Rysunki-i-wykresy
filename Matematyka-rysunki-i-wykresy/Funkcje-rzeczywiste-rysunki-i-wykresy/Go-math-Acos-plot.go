package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	mathAcosFunPlotData := make(plotter.XYs, 601)

	x := -1.0
	deltaX := 1.0/100.0

	for i := range mathAcosFunPlotData {
		mathAcosFunPlotData[i].X = x
		mathAcosFunPlotData[i].Y = math.Acos(x)

		x += deltaX

		if x >= 1.0 {
			break
		}
	}

	mathAcosFunPlotData[600].X = 1.0
	mathAcosFunPlotData[600].Y = math.Acos(1.0)


	plotFun := plot.New()

	plotFun.Title.Text = "Wykres funkcji math.Acos(x)"

	plotFun.X.Label.Text = "x"
	plotFun.Y.Label.Text = "y"

	err := plotutil.AddLinePoints(plotFun,"math.Acos",
		mathAcosFunPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Acos_plot.png"); err != nil {

		panic(err)
	}
}
