package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	mathCosFunPlotData := make(plotter.XYs, 61)

	x := 0.0
	deltaX := math.Pi / 60.0

	mathCosFunPlotData[0].X = x
	mathCosFunPlotData[0].Y = math.Cos(x)

	x += deltaX

	mathCosFunPlotData[1].X = x
	mathCosFunPlotData[1].Y = math.Cos(x)

	x += deltaX

	mathCosFunPlotData[2].X = x
	mathCosFunPlotData[2].Y = math.Cos(x)
















	plotOfFun := plot.New()

	plotOfFun.Title.Text = "Wykres funkcji math.Cos(x)"

	plotOfFun.X.Label.Text = "x"
	plotOfFun.Y.Label.Text = "y"

	l, err := plotter.NewLine(mathCosFunPlotData)

	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFun.Add(l)
	plotOfFun.Legend.Add("math.Cos(x)", l)

	if err := plotOfFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Cos_plot_01.png"); err != nil {

		panic(err)
	}
}
