package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	mathSinFunPlotData := make(plotter.XYs, 601)

	x := 0.0
	deltaX := math.Pi / 300.0

	for i := range mathSinFunPlotData {
		mathSinFunPlotData[i].X = x
		mathSinFunPlotData[i].Y = math.Sin(x)

		x += deltaX
	}



	plotOfFun := plot.New()

	plotOfFun.Title.Text = "Wykres funkcji math.Sin(x)"

	plotOfFun.X.Label.Text = "x"
	plotOfFun.Y.Label.Text = "y"

	l, err := plotter.NewLine(mathSinFunPlotData)

	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFun.Add(l)
	plotOfFun.Legend.Add("math.Sin(x)", l)

	if err := plotOfFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Sin_plot.png"); err != nil {

		panic(err)
	}
}
