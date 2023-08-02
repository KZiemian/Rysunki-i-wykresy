package main

import (
	"math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	mathAbsFunPlotData := make(plotter.XYs, 3)

	mathAbsFunPlotData[0].X = -1.0
	mathAbsFunPlotData[0].Y = math.Abs(-1.0)

	mathAbsFunPlotData[1].X = 0.0
	mathAbsFunPlotData[1].Y = math.Abs(0.0)

	mathAbsFunPlotData[2].X = 1.0
	mathAbsFunPlotData[2].Y = math.Abs(1.0)



	plotFun := plot.New()

	plotFun.Title.Text = "Wykres funkcji math.Abs(x)"

	plotFun.X.Label.Text = "x"
	plotFun.Y.Label.Text = "y"

	err := plotutil.AddLinePoints(plotFun,
		"math.Abs", mathAbsFunPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Abs_plot.png"); err != nil {
		panic(err)
	}
}
