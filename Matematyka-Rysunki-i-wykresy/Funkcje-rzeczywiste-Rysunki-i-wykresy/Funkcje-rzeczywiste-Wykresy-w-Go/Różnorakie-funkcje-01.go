package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres funkcji f(x) = -(2/3)x^3 + x - 5cos(x) + 5
	// Jest to funkcja pierwotna funkcji g(x) = -2x^2 + 1 + 5sin(x),
	// spełniający warunek f(0) = 0

	punktyWykresuFunkcji := make(plotter.XYs, 100)

	punktyWykresuFunkcji[0].X = 0.0
	punktyWykresuFunkcji[0].Y = 0.0

	punktyWykresuFunkcji[1].X = 0.01
	punktyWykresuFunkcji[1].Y = 0.01

	punktyWykresuFunkcji[2].X = 0.02
	punktyWykresuFunkcji[2].Y = 0.021

	punktyWykresuFunkcji[3].X = 0.03
	punktyWykresuFunkcji[3].Y = 0.032

	punktyWykresuFunkcji[4].X = 0.04
	punktyWykresuFunkcji[4].Y = 0.044

	punktyWykresuFunkcji[5].X = 0.05
	punktyWykresuFunkcji[5].Y = 0.056

	punktyWykresuFunkcji[6].X = 0.06
	punktyWykresuFunkcji[6].Y = 0.068

	punktyWykresuFunkcji[7].X = 0.07
	punktyWykresuFunkcji[7].Y = 0.082

	punktyWykresuFunkcji[8].X = 0.08
	punktyWykresuFunkcji[8].Y = 0.095

	punktyWykresuFunkcji[9].X = 0.09
	punktyWykresuFunkcji[9].Y = 0.109

	punktyWykresuFunkcji[10].X = 0.1
	punktyWykresuFunkcji[10].Y = 0.124

	punktyWykresuFunkcji[11].X = 0.11
	punktyWykresuFunkcji[11].Y = 0.139

	punktyWykresuFunkcji[12].X = 0.12
	punktyWykresuFunkcji[12].Y = 0.154

	punktyWykresuFunkcji[13].X = 0.13
	punktyWykresuFunkcji[13].Y = 0.17

	punktyWykresuFunkcji[14].X = 0.14
	punktyWykresuFunkcji[14].Y = 0.187

	punktyWykresuFunkcji[15].X = 0.15
	punktyWykresuFunkcji[15].Y = 0.203

	punktyWykresuFunkcji[16].X = 0.16
	punktyWykresuFunkcji[16].Y = 0.221

	punktyWykresuFunkcji[17].X = 0.17
	punktyWykresuFunkcji[17].Y = 0.238

	punktyWykresuFunkcji[18].X = 0.18
	punktyWykresuFunkcji[18].Y = 0.256

	punktyWykresuFunkcji[19].X = 0.19
	punktyWykresuFunkcji[19].Y = 0.275

	punktyWykresuFunkcji[20].X = 0.2
	punktyWykresuFunkcji[20].Y = 0.294


















	wykresFunkcji := plot.New()

	wykresFunkcji.Title.Text = "Wykres funkcji f(x) = -(2/3)x^3 + x - 5cos(x) + 5"

	wykresFunkcji.X.Label.Text = "x"
	wykresFunkcji.Y.Label.Text = "y"

	liniaWykresu, err := plotter.NewLine(punktyWykresuFunkcji)

	if err != nil {
		panic(err)
	}

	liniaWykresu.LineStyle.Width = vg.Points(0.1)
	liniaWykresu.Color = color.RGBA{R: 200, G: 100, B: 100}

	wykresFunkcji.Add(liniaWykresu)
	wykresFunkcji.Legend.Add("f(x)", liniaWykresu)

	if err := wykresFunkcji.Save(10*vg.Inch, 10*vg.Inch,
		"Wykresy-funkcji-01.png"); err != nil {

		panic(err)
	}
}
