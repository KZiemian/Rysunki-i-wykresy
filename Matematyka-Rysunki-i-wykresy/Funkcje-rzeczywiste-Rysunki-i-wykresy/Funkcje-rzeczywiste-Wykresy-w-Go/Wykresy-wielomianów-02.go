package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres wielomianu f(x) = x^4 - x^3 - x^2 + x

	punktyWykresuWielomianu := make(plotter.XYs, 26)

	punktyWykresuWielomianu[0].X = 0.0
	punktyWykresuWielomianu[0].Y = 0.0

	punktyWykresuWielomianu[1].X = 0.05
	punktyWykresuWielomianu[1].Y = -0.052

	punktyWykresuWielomianu[2].X = 0.1
	punktyWykresuWielomianu[2].Y = -0.11

	punktyWykresuWielomianu[3].X = 0.15
	punktyWykresuWielomianu[3].Y = -0.175

	punktyWykresuWielomianu[4].X = 0.15
	punktyWykresuWielomianu[4].Y = -0.175

	punktyWykresuWielomianu[5].X = 0.2
	punktyWykresuWielomianu[5].Y = -0.246

	punktyWykresuWielomianu[6].X = 0.25
	punktyWykresuWielomianu[6].Y = -0.324

	punktyWykresuWielomianu[7].X = 0.3
	punktyWykresuWielomianu[7].Y = -0.408

	punktyWykresuWielomianu[8].X = 0.35
	punktyWykresuWielomianu[8].Y = -0.5

	punktyWykresuWielomianu[9].X = 0.4
	punktyWykresuWielomianu[9].Y = -0.598

	punktyWykresuWielomianu[10].X = 0.45
	punktyWykresuWielomianu[10].Y = -0.702

	punktyWykresuWielomianu[11].X = 0.5
	punktyWykresuWielomianu[11].Y = -0.812

	punktyWykresuWielomianu[12].X = 0.55
	punktyWykresuWielomianu[12].Y = -0.927

	punktyWykresuWielomianu[13].X = 0.6
	punktyWykresuWielomianu[13].Y = -1.046

	punktyWykresuWielomianu[14].X = 0.65
	punktyWykresuWielomianu[14].Y = -1.168

	punktyWykresuWielomianu[15].X = 0.7
	punktyWykresuWielomianu[15].Y = -1.292

	punktyWykresuWielomianu[16].X = 0.75
	punktyWykresuWielomianu[16].Y = -1.418

	punktyWykresuWielomianu[17].X = 0.8
	punktyWykresuWielomianu[17].Y = -1.542

	punktyWykresuWielomianu[18].X = 0.85
	punktyWykresuWielomianu[18].Y = -1.664

	punktyWykresuWielomianu[19].X = 0.9
	punktyWykresuWielomianu[19].Y = -1.782

	punktyWykresuWielomianu[20].X = 0.95
	punktyWykresuWielomianu[20].Y = -1.895

	punktyWykresuWielomianu[21].X = 1.0
	punktyWykresuWielomianu[21].Y = -2.0

	punktyWykresuWielomianu[22].X = 2.0
	punktyWykresuWielomianu[22].Y = 2.0

	punktyWykresuWielomianu[23].X = 3.0
	punktyWykresuWielomianu[23].Y = 42.0

	punktyWykresuWielomianu[24].X = 4.0
	punktyWykresuWielomianu[24].Y = 172.0

	punktyWykresuWielomianu[25].X = 5.0
	punktyWykresuWielomianu[25].Y = 470.0





	wykresWielomianu := plot.New()

	wykresWielomianu.Title.Text = "Wykres funkcji f(x) = x^4 - x^3 - x^2 - x"

	wykresWielomianu.X.Label.Text = "x"
	wykresWielomianu.Y.Label.Text = "y"

	liniaWykresu, err := plotter.NewLine(punktyWykresuWielomianu)

	if err != nil {
		panic(err)
	}

	liniaWykresu.LineStyle.Width = vg.Points(0.1)
	liniaWykresu.Color = color.RGBA{R: 200, G: 100, B: 100}

	wykresWielomianu.Add(liniaWykresu)
	wykresWielomianu.Legend.Add("f(x)", liniaWykresu)

	if err := wykresWielomianu.Save(10*vg.Inch, 10*vg.Inch,
		"Wykres-wielomianu-02.png"); err != nil {

		panic(err)
	}
}
