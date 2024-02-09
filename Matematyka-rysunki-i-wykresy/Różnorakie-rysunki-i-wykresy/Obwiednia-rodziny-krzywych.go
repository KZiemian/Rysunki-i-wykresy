package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)


func main() {
	p := plot.New()

	p.Title.Text = "Obwiednia rodziny krzywych"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	firstLine := plotter.NewFunction(func(x float64) float64 {
		return 0.2*x + 0.9797
	})
	firstLine.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	secondLine := plotter.NewFunction(func(x float64) float64 {
		return 0.15*x + 0.9886
	})
	secondLine.Color = color.RGBA{R: 200, G:55, B: 0, A: 255}

	thirdLine := plotter.NewFunction(func(x float64) float64 {
		return 0.1*x + 0.9949
	})
	thirdLine.Color = color.RGBA{R: 200, G: 0, B: 55, A: 255}

	fourthLine := plotter.NewFunction(func(x float64) float64 {
		return 0.05*x + 0.9987
	})
	fourthLine.Color = color.RGBA{R: 180, G: 45, B: 30, A: 255}

	fifthLine := plotter.NewFunction(func(x float64) float64 {
		return 1
	})
	fifthLine.Color = color.RGBA{R: 180, G: 30, B: 45, A: 255}

	sixthLine := plotter.NewFunction(func(x float64) float64 {
		return -0.05*x + 0.9987
	})
	sixthLine.Color = color.RGBA{R: 150, G: 45, B: 60, A: 255}

	seventhLine := plotter.NewFunction(func(x float64) float64 {
		return -0.1*x + 0.9949
	})
	seventhLine.Color = color.RGBA{R: 150, G: 60, B: 45, A: 255}

	eighthLine := plotter.NewFunction(func(x float64) float64 {
		return -0.15*x + 0.9886
	})
	eighthLine.Color = color.RGBA{R: 130, G: 60, B: 65, A: 255}

	ninthLine := plotter.NewFunction(func(x float64) float64 {
		return -0.2*x + 0.9797
	})


	p.Add(firstLine, secondLine, thirdLine, fourthLine, fifthLine,
		sixthLine, seventhLine, eighthLine, ninthLine)

	p.Legend.Add("0.2", firstLine)
	p.Legend.Add("0.15", secondLine)
	p.Legend.Add("0.1", thirdLine)
	p.Legend.Add("0.05", fourthLine)
	p.Legend.Add("0.0", fifthLine)
	p.Legend.Add("-0.05", sixthLine)
	p.Legend.Add("-0.1", seventhLine)
	p.Legend.Add("-0.15", eighthLine)
	p.Legend.Add("-0.2", ninthLine)

	p.X.Min = -1.25
	p.X.Max = 1.25
	p.Y.Min = 0.8
	p.Y.Max = 1.2

	if err := p.Save(10*vg.Inch, 10*vg.Inch,
		"Obwiednia-rodziny-krzywych.png"); err != nil {
		panic(err)
	}
}
