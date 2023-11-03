package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	mathAcoshFunPlotData := make(plotter.XYs, 20)


	mathAcoshValue := math.Acosh(1.0)

	mathAcoshFunPlotData[0].X = 1.0
	mathAcoshFunPlotData[0].Y = mathAcoshValue

	fmt.Printf("math.Acosh(1.0) = %.3f.\n", mathAcoshValue)

	mathAcoshValue = math.Acosh(1.1)

	mathAcoshFunPlotData[1].X = 1.1
	mathAcoshFunPlotData[1].Y = mathAcoshValue

	fmt.Printf("math.Acosh(1.1) = %.3f.\n", mathAcoshValue)

	mathAcoshFunPlotData[2].X = 1.2
	mathAcoshFunPlotData[2].Y = math.Acosh(1.2)

	mathAcoshFunPlotData[3].X = 1.3
	mathAcoshFunPlotData[3].Y = math.Acosh(1.3)

	mathAcoshFunPlotData[4].X = 1.4
	mathAcoshFunPlotData[4].Y = math.Acosh(1.4)

	mathAcoshFunPlotData[5].X = 1.5
	mathAcoshFunPlotData[5].Y = math.Acosh(1.5)

	mathAcoshFunPlotData[6].X = 1.6
	mathAcoshFunPlotData[6].Y = math.Acosh(1.6)

	mathAcoshFunPlotData[7].X = 1.7
	mathAcoshFunPlotData[7].Y = math.Acosh(1.7)

	mathAcoshFunPlotData[8].X = 1.8
	mathAcoshFunPlotData[8].Y = math.Acosh(1.8)






	plotFun := plot.New()

	plotFun.Title.Text = "Wykres funkcji math.Acosh(x)"

	plotFun.X.Label.Text = "x"
	plotFun.Y.Label.Text = "y"

	err := plotutil.AddLinePoints(plotFun,
		"math.Acosh", mathAcoshFunPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotFun.Save(10*vg.Inch, 10*vg.Inch,
		"Go_math_Acosh_plot.png"); err != nil {
		panic(err)
	}
}
