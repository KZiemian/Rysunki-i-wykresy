package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	mathAcosFunPlotData := make(plotter.XYs, 201)

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

	mathAcosFunPlotData[200].X = 1.0
	mathAcosFunPlotData[200].Y = math.Acos(1.0)


	plotOfFun := plot.New()

	plotOfFun.Title.Text = "Wykres funkcji math.Acos(x)"

	plotOfFun.X.Label.Text = "x"
	plotOfFun.Y.Label.Text = "y"

	l, err := plotter.NewLine(mathAcosFunPlotData)

	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFun.Add(l)
	plotOfFun.Legend.Add("math.Acos(x)", l)

	if err := plotOfFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Acos_plot.png"); err != nil {

		panic(err)
	}
}
