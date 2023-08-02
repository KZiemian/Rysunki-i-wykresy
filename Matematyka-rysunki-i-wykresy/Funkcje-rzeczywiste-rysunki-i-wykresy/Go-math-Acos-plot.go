package main

import (
	"math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	mathAcosFunPlotData := make(plotter.XYs, 21)

	mathAcosFunPlotData[0].X = -1.0
	mathAcosFunPlotData[0].Y = math.Acos(-1.0)

	mathAcosFunPlotData[1].X = -0.9
	mathAcosFunPlotData[1].Y = math.Acos(-0.9)

	mathAcosFunPlotData[2].X = -0.8
	mathAcosFunPlotData[2].Y = math.Acos(-0.8)

	mathAcosFunPlotData[3].X = -0.7
	mathAcosFunPlotData[3].Y = math.Acos(-0.7)

	mathAcosFunPlotData[4].X = -0.6
	mathAcosFunPlotData[4].Y = math.Acos(-0.6)

	mathAcosFunPlotData[5].X = -0.5
	mathAcosFunPlotData[5].Y = math.Acos(-0.5)

	mathAcosFunPlotData[6].X = -0.4
	mathAcosFunPlotData[6].Y = math.Acos(-0.4)

	mathAcosFunPlotData[7].X = -0.3
	mathAcosFunPlotData[7].Y = math.Acos(-0.3)

	mathAcosFunPlotData[8].X = -0.2
	mathAcosFunPlotData[8].Y = math.Acos(-0.2)

	mathAcosFunPlotData[9].X = -0.1
	mathAcosFunPlotData[9].Y = math.Acos(-0.1)

	mathAcosFunPlotData[10].X = 0.0
	mathAcosFunPlotData[10].Y = math.Acos(0.0)

	mathAcosFunPlotData[11].X = 0.1
	mathAcosFunPlotData[11].Y = math.Acos(0.1)

	mathAcosFunPlotData[12].X = 0.2
	mathAcosFunPlotData[12].Y = math.Acos(0.2)

	mathAcosFunPlotData[13].X = 0.3
	mathAcosFunPlotData[13].Y = math.Acos(0.3)

	mathAcosFunPlotData[14].X = 0.4
	mathAcosFunPlotData[14].Y = math.Acos(0.4)

	mathAcosFunPlotData[15].X = 0.5
	mathAcosFunPlotData[15].Y = math.Acos(0.5)

	mathAcosFunPlotData[16].X = 0.6
	mathAcosFunPlotData[16].Y = math.Acos(0.6)

	mathAcosFunPlotData[17].X = 0.7
	mathAcosFunPlotData[17].Y = math.Acos(0.7)

	mathAcosFunPlotData[18].X = 0.8
	mathAcosFunPlotData[18].Y = math.Acos(0.8)

	mathAcosFunPlotData[19].X = 0.9
	mathAcosFunPlotData[19].Y = math.Acos(0.9)

	mathAcosFunPlotData[20].X = 1.0
	mathAcosFunPlotData[20].Y = math.Acos(1.0)



	plotFun := plot.New()

	plotFun.Title.Text = "Wykres funkcji math.Acos(x)"

	plotFun.X.Label.Text = "x"
	plotFun.Y.Label.Text = "y"

	err := plotutil.AddLinePoints(plotFun,
		"math.Acos", mathAcosFunPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Acos_plot.png"); err != nil {
		panic(err)
	}
}
