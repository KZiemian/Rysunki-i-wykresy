package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
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



	plotOfFun := plot.New()

	plotOfFun.Title.Text = "Wykres funkcji math.Abs(x)"

	plotOfFun.X.Label.Text = "x"
	plotOfFun.Y.Label.Text = "y"

	l, err := plotter.NewLine(mathAbsFunPlotData)

	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFun.Add(l)
	plotOfFun.Legend.Add("math.Abs(x)", l)


	if err := plotOfFun.Save(10*vg.Inch, 5*vg.Inch,
		"Go_math_Abs_plot_01.png"); err != nil {

		panic(err)
	}
}
