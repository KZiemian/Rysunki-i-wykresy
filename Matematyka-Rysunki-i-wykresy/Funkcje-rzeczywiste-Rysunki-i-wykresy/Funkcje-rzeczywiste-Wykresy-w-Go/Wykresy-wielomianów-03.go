package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres wielomianu f(x) = x^4 - 20x^3 + 33.75x^2 - 235x + 89.0625

	punktyWykresuWielomianu := make(plotter.XYs, 100)

	punktyWykresuWielomianu[0].X = 0.0
	punktyWykresuWielomianu[0].Y = 89.0625

	punktyWykresuWielomianu[1].X = 0.01
	punktyWykresuWielomianu[1].Y = 86.715

	punktyWykresuWielomianu[2].X = 0.02
	punktyWykresuWielomianu[2].Y = 84.375

	punktyWykresuWielomianu[3].X = 0.03
	punktyWykresuWielomianu[3].Y = 82.0423

	punktyWykresuWielomianu[4].X = 0.04
	punktyWykresuWielomianu[4].Y = 79.715

	punktyWykresuWielomianu[5].X = 0.05
	punktyWykresuWielomianu[5].Y = 77.394

	punktyWykresuWielomianu[6].X = 0.06
	punktyWykresuWielomianu[6].Y = 75.079

	punktyWykresuWielomianu[7].X = 0.07
	punktyWykresuWielomianu[7].Y = 72.771

	punktyWykresuWielomianu[8].X = 0.08
	punktyWykresuWielomianu[8].Y = 70.468

	punktyWykresuWielomianu[9].X = 0.09
	punktyWykresuWielomianu[9].Y = 68.171

	punktyWykresuWielomianu[10].X = 0.1
	punktyWykresuWielomianu[10].Y = 65.88

	punktyWykresuWielomianu[11].X = 0.11
	punktyWykresuWielomianu[11].Y = 63.594

	punktyWykresuWielomianu[12].X = 0.12
	punktyWykresuWielomianu[12].Y = 61.314

	punktyWykresuWielomianu[13].X = 0.13
	punktyWykresuWielomianu[13].Y = 59.039

	punktyWykresuWielomianu[14].X = 0.14
	punktyWykresuWielomianu[14].Y = 56.769

	punktyWykresuWielomianu[15].X = 0.15
	punktyWykresuWielomianu[15].Y = 54.504

	punktyWykresuWielomianu[16].X = 0.16
	punktyWykresuWielomianu[16].Y = 52.245

	punktyWykresuWielomianu[17].X = 0.17
	punktyWykresuWielomianu[17].Y = 49.99

	punktyWykresuWielomianu[18].X = 0.18
	punktyWykresuWielomianu[18].Y = 47.74

	punktyWykresuWielomianu[19].X = 0.19
	punktyWykresuWielomianu[19].Y = 45.495

	punktyWykresuWielomianu[20].X = 0.2
	punktyWykresuWielomianu[20].Y = 43.254

	punktyWykresuWielomianu[21].X = 0.21
	punktyWykresuWielomianu[21].Y = 41.017

	punktyWykresuWielomianu[22].X = 0.22
	punktyWykresuWielomianu[22].Y = 38.785

	punktyWykresuWielomianu[23].X = 0.23
	punktyWykresuWielomianu[23].Y = 36.557

	punktyWykresuWielomianu[24].X = 0.24
	punktyWykresuWielomianu[24].Y = 34.333

	punktyWykresuWielomianu[25].X = 0.25
	punktyWykresuWielomianu[25].Y = 32.113

	punktyWykresuWielomianu[26].X = 0.26
	punktyWykresuWielomianu[26].Y = 29.897

	punktyWykresuWielomianu[27].X = 0.27
	punktyWykresuWielomianu[27].Y = 27.684

	punktyWykresuWielomianu[28].X = 0.28
	punktyWykresuWielomianu[28].Y = 25.475

	punktyWykresuWielomianu[29].X = 0.29
	punktyWykresuWielomianu[29].Y = 23.27

	punktyWykresuWielomianu[30].X = 0.3
	punktyWykresuWielomianu[30].Y = 21.068

	punktyWykresuWielomianu[31].X = 0.31
	punktyWykresuWielomianu[31].Y = 18.869

	punktyWykresuWielomianu[32].X = 0.32
	punktyWykresuWielomianu[32].Y = 16.673

	punktyWykresuWielomianu[33].X = 0.33
	punktyWykresuWielomianu[33].Y = 14.481

	punktyWykresuWielomianu[34].X = 0.34
	punktyWykresuWielomianu[34].Y = 12.291

	punktyWykresuWielomianu[35].X = 0.35
	punktyWykresuWielomianu[35].Y = 10.104

	punktyWykresuWielomianu[36].X = 0.36
	punktyWykresuWielomianu[36].Y = 7.92

	punktyWykresuWielomianu[37].X = 0.37
	punktyWykresuWielomianu[37].Y = 5.738

	punktyWykresuWielomianu[38].X = 0.38
	punktyWykresuWielomianu[38].Y = 3.559

	punktyWykresuWielomianu[39].X = 0.39
	punktyWykresuWielomianu[39].Y = 1.382

	punktyWykresuWielomianu[40].X = 0.4
	punktyWykresuWielomianu[40].Y = -0.791

	// ???
	punktyWykresuWielomianu[41].X = 0.41
	punktyWykresuWielomianu[41].Y = -0.791

	punktyWykresuWielomianu[42].X = 0.42
	punktyWykresuWielomianu[42].Y = -0.791

	punktyWykresuWielomianu[43].X = 0.43
	punktyWykresuWielomianu[43].Y = -0.791

	punktyWykresuWielomianu[44].X = 0.44
	punktyWykresuWielomianu[44].Y = -0.791

	punktyWykresuWielomianu[45].X = 0.45
	punktyWykresuWielomianu[45].Y = -0.791

	punktyWykresuWielomianu[46].X = 0.46
	punktyWykresuWielomianu[46].Y = -0.791

	punktyWykresuWielomianu[47].X = 0.47
	punktyWykresuWielomianu[47].Y = -0.791

	punktyWykresuWielomianu[48].X = 0.48
	punktyWykresuWielomianu[48].Y = -0.791

	punktyWykresuWielomianu[49].X = 0.49
	punktyWykresuWielomianu[49].Y = -0.791

	punktyWykresuWielomianu[50].X = 0.5
	punktyWykresuWielomianu[50].Y = -0.791

	punktyWykresuWielomianu[51].X = 0.51
	punktyWykresuWielomianu[51].Y = -0.791

	punktyWykresuWielomianu[52].X = 0.52
	punktyWykresuWielomianu[52].Y = -0.791

	punktyWykresuWielomianu[53].X = 0.53
	punktyWykresuWielomianu[53].Y = -0.791

	punktyWykresuWielomianu[54].X = 0.54
	punktyWykresuWielomianu[54].Y = -0.791

	punktyWykresuWielomianu[55].X = 0.55
	punktyWykresuWielomianu[55].Y = -0.791

	punktyWykresuWielomianu[56].X = 0.56
	punktyWykresuWielomianu[56].Y = -0.791

	punktyWykresuWielomianu[57].X = 0.57
	punktyWykresuWielomianu[57].Y = -0.791

	punktyWykresuWielomianu[58].X = 0.58
	punktyWykresuWielomianu[58].Y = -0.791

	punktyWykresuWielomianu[59].X = 0.59
	punktyWykresuWielomianu[59].Y = -0.791

	punktyWykresuWielomianu[60].X = 0.6
	punktyWykresuWielomianu[60].Y = -0.791

	punktyWykresuWielomianu[61].X = 0.6
	punktyWykresuWielomianu[61].Y = -0.791

	punktyWykresuWielomianu[62].X = 0.6
	punktyWykresuWielomianu[62].Y = -0.791

	punktyWykresuWielomianu[63].X = 0.6
	punktyWykresuWielomianu[63].Y = -0.791

	punktyWykresuWielomianu[64].X = 0.6
	punktyWykresuWielomianu[64].Y = -0.791

	punktyWykresuWielomianu[65].X = 0.6
	punktyWykresuWielomianu[65].Y = -0.791

	punktyWykresuWielomianu[66].X = 0.6
	punktyWykresuWielomianu[66].Y = -0.791

	punktyWykresuWielomianu[67].X = 0.6
	punktyWykresuWielomianu[67].Y = -0.791

	punktyWykresuWielomianu[68].X = 0.6
	punktyWykresuWielomianu[68].Y = -0.791

	punktyWykresuWielomianu[69].X = 0.6
	punktyWykresuWielomianu[69].Y = -0.791

	punktyWykresuWielomianu[70].X = 0.6
	punktyWykresuWielomianu[70].Y = -0.791

	punktyWykresuWielomianu[71].X = 0.6
	punktyWykresuWielomianu[71].Y = -0.791

	punktyWykresuWielomianu[72].X = 0.6
	punktyWykresuWielomianu[72].Y = -0.791

	punktyWykresuWielomianu[73].X = 0.6
	punktyWykresuWielomianu[73].Y = -0.791

	punktyWykresuWielomianu[74].X = 0.6
	punktyWykresuWielomianu[74].Y = -0.791

	punktyWykresuWielomianu[75].X = 0.6
	punktyWykresuWielomianu[75].Y = -0.791



	wykresWielomianu := plot.New()

	wykresWielomianu.Title.Text = "Wykres funkcji f(x) = x^4 - 20x^3 - 33.75x^2 - 235x + 89.0625"

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
		"Wykres-wielomianu-03.png"); err != nil {

		panic(err)
	}
}
