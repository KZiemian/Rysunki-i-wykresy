package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of Gamma function.

	pointsOfFunctionPlot1 := make(plotter.XYs, 101)

	pointsOfFunctionPlot1[0].X = -6.999
	pointsOfFunctionPlot1[0].Y = -0.198

	pointsOfFunctionPlot1[1].X = -6.99
	pointsOfFunctionPlot1[1].Y = -0.02

	pointsOfFunctionPlot1[2].X = -6.98
	pointsOfFunctionPlot1[2].Y = -0.01

	pointsOfFunctionPlot1[3].X = -6.97
	pointsOfFunctionPlot1[3].Y = -0.007

	pointsOfFunctionPlot1[4].X = -6.96
	pointsOfFunctionPlot1[4].Y = -0.005

	pointsOfFunctionPlot1[5].X = -6.95
	pointsOfFunctionPlot1[5].Y = -0.004

	pointsOfFunctionPlot1[6].X = -6.94
	pointsOfFunctionPlot1[6].Y = -0.003

	pointsOfFunctionPlot1[7].X = -6.93
	pointsOfFunctionPlot1[7].Y = -0.003

	pointsOfFunctionPlot1[8].X = -6.92
	pointsOfFunctionPlot1[8].Y = -0.002

	pointsOfFunctionPlot1[9].X = -6.91
	pointsOfFunctionPlot1[9].Y = -0.002

	pointsOfFunctionPlot1[10].X = -6.90
	pointsOfFunctionPlot1[10].Y = -0.002

	pointsOfFunctionPlot1[11].X = -6.89
	pointsOfFunctionPlot1[11].Y = -0.002

	pointsOfFunctionPlot1[12].X = -6.88
	pointsOfFunctionPlot1[12].Y = -0.002

	pointsOfFunctionPlot1[13].X = -6.87
	pointsOfFunctionPlot1[13].Y = -0.002

	pointsOfFunctionPlot1[14].X = -6.86
	pointsOfFunctionPlot1[14].Y = -0.001

	pointsOfFunctionPlot1[15].X = -6.85
	pointsOfFunctionPlot1[15].Y = -0.001

	pointsOfFunctionPlot1[16].X = -6.84
	pointsOfFunctionPlot1[16].Y = -0.001

	pointsOfFunctionPlot1[17].X = -6.83
	pointsOfFunctionPlot1[17].Y = -0.001

	pointsOfFunctionPlot1[18].X = -6.82
	pointsOfFunctionPlot1[18].Y = -0.001

	pointsOfFunctionPlot1[19].X = -6.81
	pointsOfFunctionPlot1[19].Y = -0.001

	pointsOfFunctionPlot1[20].X = -6.80
	pointsOfFunctionPlot1[20].Y = -0.001

	pointsOfFunctionPlot1[21].X = -6.79
	pointsOfFunctionPlot1[21].Y = -0.001

	pointsOfFunctionPlot1[22].X = -6.78
	pointsOfFunctionPlot1[22].Y = -0.001

	pointsOfFunctionPlot1[23].X = -6.77
	pointsOfFunctionPlot1[23].Y = -0.001

	pointsOfFunctionPlot1[24].X = -6.76
	pointsOfFunctionPlot1[24].Y = -0.001

	pointsOfFunctionPlot1[25].X = -6.75
	pointsOfFunctionPlot1[25].Y = -0.001

	pointsOfFunctionPlot1[26].X = -6.74
	pointsOfFunctionPlot1[26].Y = -0.001

	pointsOfFunctionPlot1[27].X = -6.73
	pointsOfFunctionPlot1[27].Y = -0.001

	pointsOfFunctionPlot1[28].X = -6.72
	pointsOfFunctionPlot1[28].Y = -0.001

	pointsOfFunctionPlot1[29].X = -6.71
	pointsOfFunctionPlot1[29].Y = -0.001

	pointsOfFunctionPlot1[30].X = -6.70
	pointsOfFunctionPlot1[30].Y = -0.001

	pointsOfFunctionPlot1[31].X = -6.69
	pointsOfFunctionPlot1[31].Y = -0.001

	pointsOfFunctionPlot1[32].X = -6.68
	pointsOfFunctionPlot1[32].Y = -0.001

	pointsOfFunctionPlot1[33].X = -6.67
	pointsOfFunctionPlot1[33].Y = -0.001

	pointsOfFunctionPlot1[34].X = -6.66
	pointsOfFunctionPlot1[34].Y = -0.001

	pointsOfFunctionPlot1[35].X = -6.65
	pointsOfFunctionPlot1[35].Y = -0.001

	pointsOfFunctionPlot1[36].X = -6.64
	pointsOfFunctionPlot1[36].Y = -0.001

	pointsOfFunctionPlot1[37].X = -6.63
	pointsOfFunctionPlot1[37].Y = -0.001

	pointsOfFunctionPlot1[38].X = -6.62
	pointsOfFunctionPlot1[38].Y = -0.001

	pointsOfFunctionPlot1[39].X = -6.61
	pointsOfFunctionPlot1[39].Y = -0.001

	pointsOfFunctionPlot1[40].X = -6.60
	pointsOfFunctionPlot1[40].Y = -0.001

	pointsOfFunctionPlot1[41].X = -6.59
	pointsOfFunctionPlot1[41].Y = -0.001

	pointsOfFunctionPlot1[42].X = -6.58
	pointsOfFunctionPlot1[42].Y = -0.001

	pointsOfFunctionPlot1[43].X = -6.57
	pointsOfFunctionPlot1[43].Y = -0.001

	pointsOfFunctionPlot1[44].X = -6.56
	pointsOfFunctionPlot1[44].Y = -0.001

	pointsOfFunctionPlot1[45].X = -6.55
	pointsOfFunctionPlot1[45].Y = -0.001

	pointsOfFunctionPlot1[46].X = -6.54
	pointsOfFunctionPlot1[46].Y = -0.001

	pointsOfFunctionPlot1[47].X = -6.53
	pointsOfFunctionPlot1[47].Y = -0.001

	pointsOfFunctionPlot1[48].X = -6.52
	pointsOfFunctionPlot1[48].Y = -0.001

	pointsOfFunctionPlot1[49].X = -6.51
	pointsOfFunctionPlot1[49].Y = -0.001

	pointsOfFunctionPlot1[50].X = -6.50
	pointsOfFunctionPlot1[50].Y = -0.001

	pointsOfFunctionPlot1[51].X = -6.49
	pointsOfFunctionPlot1[51].Y = -0.001

	pointsOfFunctionPlot1[52].X = -6.48
	pointsOfFunctionPlot1[52].Y = -0.001

	pointsOfFunctionPlot1[53].X = -6.47
	pointsOfFunctionPlot1[53].Y = -0.001

	pointsOfFunctionPlot1[54].X = -6.46
	pointsOfFunctionPlot1[54].Y = -0.001

	pointsOfFunctionPlot1[55].X = -6.45
	pointsOfFunctionPlot1[55].Y = -0.001

	pointsOfFunctionPlot1[56].X = -6.44
	pointsOfFunctionPlot1[56].Y = -0.001

	pointsOfFunctionPlot1[57].X = -6.43
	pointsOfFunctionPlot1[57].Y = -0.001

	pointsOfFunctionPlot1[58].X = -6.42
	pointsOfFunctionPlot1[58].Y = -0.002

	pointsOfFunctionPlot1[59].X = -6.41
	pointsOfFunctionPlot1[59].Y = -0.002

	pointsOfFunctionPlot1[60].X = -6.40
	pointsOfFunctionPlot1[60].Y = -0.002

	pointsOfFunctionPlot1[61].X = -6.39
	pointsOfFunctionPlot1[61].Y = -0.002

	pointsOfFunctionPlot1[62].X = -6.38
	pointsOfFunctionPlot1[62].Y = -0.002

	pointsOfFunctionPlot1[63].X = -6.37
	pointsOfFunctionPlot1[63].Y = -0.002

	pointsOfFunctionPlot1[64].X = -6.36
	pointsOfFunctionPlot1[64].Y = -0.002

	pointsOfFunctionPlot1[65].X = -6.35
	pointsOfFunctionPlot1[65].Y = -0.002

	pointsOfFunctionPlot1[66].X = -6.34
	pointsOfFunctionPlot1[66].Y = -0.002

	pointsOfFunctionPlot1[67].X = -6.33
	pointsOfFunctionPlot1[67].Y = -0.002

	pointsOfFunctionPlot1[68].X = -6.32
	pointsOfFunctionPlot1[68].Y = -0.002

	pointsOfFunctionPlot1[69].X = -6.31
	pointsOfFunctionPlot1[69].Y = -0.002

	pointsOfFunctionPlot1[70].X = -6.30
	pointsOfFunctionPlot1[70].Y = -0.003

	pointsOfFunctionPlot1[71].X = -6.29
	pointsOfFunctionPlot1[71].Y = -0.003

	pointsOfFunctionPlot1[72].X = -6.28
	pointsOfFunctionPlot1[72].Y = -0.003

	pointsOfFunctionPlot1[73].X = -6.27
	pointsOfFunctionPlot1[73].Y = -0.003

	pointsOfFunctionPlot1[74].X = -6.26
	pointsOfFunctionPlot1[74].Y = -0.003

	pointsOfFunctionPlot1[75].X = -6.25
	pointsOfFunctionPlot1[75].Y = -0.003

	pointsOfFunctionPlot1[76].X = -6.24
	pointsOfFunctionPlot1[76].Y = -0.004

	pointsOfFunctionPlot1[77].X = -6.23
	pointsOfFunctionPlot1[77].Y = -0.004

	pointsOfFunctionPlot1[78].X = -6.22
	pointsOfFunctionPlot1[78].Y = -0.004

	pointsOfFunctionPlot1[79].X = -6.21
	pointsOfFunctionPlot1[79].Y = -0.004

	pointsOfFunctionPlot1[80].X = -6.20
	pointsOfFunctionPlot1[80].Y = -0.005

	pointsOfFunctionPlot1[81].X = -6.19
	pointsOfFunctionPlot1[81].Y = -0.005

	pointsOfFunctionPlot1[82].X = -6.18
	pointsOfFunctionPlot1[82].Y = -0.005

	pointsOfFunctionPlot1[83].X = -6.17
	pointsOfFunctionPlot1[83].Y = -0.006

	pointsOfFunctionPlot1[84].X = -6.16
	pointsOfFunctionPlot1[84].Y = -0.006

	pointsOfFunctionPlot1[85].X = -6.15
	pointsOfFunctionPlot1[85].Y = -0.007

	pointsOfFunctionPlot1[86].X = -6.14
	pointsOfFunctionPlot1[86].Y = -0.007

	pointsOfFunctionPlot1[87].X = -6.13
	pointsOfFunctionPlot1[87].Y = -0.008

	pointsOfFunctionPlot1[88].X = -6.12
	pointsOfFunctionPlot1[88].Y = -0.009

	pointsOfFunctionPlot1[89].X = -6.11
	pointsOfFunctionPlot1[89].Y = -0.01

	pointsOfFunctionPlot1[90].X = -6.10
	pointsOfFunctionPlot1[90].Y = -0.011

	pointsOfFunctionPlot1[91].X = -6.09
	pointsOfFunctionPlot1[91].Y = -0.013

	pointsOfFunctionPlot1[92].X = -6.08
	pointsOfFunctionPlot1[92].Y = -0.015

	pointsOfFunctionPlot1[93].X = -6.07
	pointsOfFunctionPlot1[93].Y = -0.017

	pointsOfFunctionPlot1[94].X = -6.06
	pointsOfFunctionPlot1[94].Y = -0.02

	pointsOfFunctionPlot1[95].X = -6.05
	pointsOfFunctionPlot1[95].Y = -0.025

	pointsOfFunctionPlot1[96].X = -6.04
	pointsOfFunctionPlot1[96].Y = -0.032

	pointsOfFunctionPlot1[97].X = -6.03
	pointsOfFunctionPlot1[97].Y = -0.043

	pointsOfFunctionPlot1[98].X = -6.02
	pointsOfFunctionPlot1[98].Y = -0.066

	pointsOfFunctionPlot1[99].X = -6.01
	pointsOfFunctionPlot1[99].Y = -0.136

	pointsOfFunctionPlot1[100].X = -6.001
	pointsOfFunctionPlot1[100].Y = -3.127





	pointsOfFunctionPlot2 := make(plotter.XYs, 100)

	pointsOfFunctionPlot2[0].X = -5.999
	pointsOfFunctionPlot2[0].Y = 1.391

	pointsOfFunctionPlot2[1].X = -5.99
	pointsOfFunctionPlot2[1].Y = 0.141

	pointsOfFunctionPlot2[2].X = -5.98
	pointsOfFunctionPlot2[2].Y = 0.072

	pointsOfFunctionPlot2[3].X = -5.97
	pointsOfFunctionPlot2[3].Y = 0.049

	pointsOfFunctionPlot2[4].X = -5.96
	pointsOfFunctionPlot2[4].Y = 0.037

	pointsOfFunctionPlot2[5].X = -5.95
	pointsOfFunctionPlot2[5].Y = 0.03

	pointsOfFunctionPlot2[6].X = -5.94
	pointsOfFunctionPlot2[6].Y = 0.026

	pointsOfFunctionPlot2[7].X = -5.93
	pointsOfFunctionPlot2[7].Y = 0.022

	pointsOfFunctionPlot2[8].X = -5.92
	pointsOfFunctionPlot2[8].Y = 0.02

	pointsOfFunctionPlot2[9].X = -5.91
	pointsOfFunctionPlot2[9].Y = 0.018

	pointsOfFunctionPlot2[10].X = -5.90
	pointsOfFunctionPlot2[10].Y = 0.017

	pointsOfFunctionPlot2[11].X = -5.89
	pointsOfFunctionPlot2[11].Y = 0.015

	pointsOfFunctionPlot2[12].X = -5.88
	pointsOfFunctionPlot2[12].Y = 0.014

	pointsOfFunctionPlot2[13].X = -5.87
	pointsOfFunctionPlot2[13].Y = 0.013

	pointsOfFunctionPlot2[14].X = -5.86
	pointsOfFunctionPlot2[14].Y = 0.013

	pointsOfFunctionPlot2[15].X = -5.85
	pointsOfFunctionPlot2[15].Y = 0.012

	pointsOfFunctionPlot2[16].X = -5.84
	pointsOfFunctionPlot2[16].Y = 0.012

	pointsOfFunctionPlot2[17].X = -5.83
	pointsOfFunctionPlot2[17].Y = 0.011

	pointsOfFunctionPlot2[18].X = -5.82
	pointsOfFunctionPlot2[18].Y = 0.011

	pointsOfFunctionPlot2[19].X = -5.81
	pointsOfFunctionPlot2[19].Y = 0.011

	pointsOfFunctionPlot2[20].X = -5.80
	pointsOfFunctionPlot2[20].Y = 0.01

	pointsOfFunctionPlot2[21].X = -5.79
	pointsOfFunctionPlot2[21].Y = 0.01

	pointsOfFunctionPlot2[22].X = -5.78
	pointsOfFunctionPlot2[22].Y = 0.01

	pointsOfFunctionPlot2[23].X = -5.77
	pointsOfFunctionPlot2[23].Y = 0.01

	pointsOfFunctionPlot2[24].X = -5.76
	pointsOfFunctionPlot2[24].Y = 0.009

	pointsOfFunctionPlot2[25].X = -5.75
	pointsOfFunctionPlot2[25].Y = 0.009

	pointsOfFunctionPlot2[26].X = -5.74
	pointsOfFunctionPlot2[26].Y = 0.009

	pointsOfFunctionPlot2[27].X = -5.73
	pointsOfFunctionPlot2[27].Y = 0.009

	pointsOfFunctionPlot2[28].X = -5.72
	pointsOfFunctionPlot2[28].Y = 0.009

	pointsOfFunctionPlot2[29].X = -5.71
	pointsOfFunctionPlot2[29].Y = 0.009

	pointsOfFunctionPlot2[30].X = -5.70
	pointsOfFunctionPlot2[30].Y = 0.009

	pointsOfFunctionPlot2[31].X = -5.69
	pointsOfFunctionPlot2[31].Y = 0.009

	pointsOfFunctionPlot2[32].X = -5.68
	pointsOfFunctionPlot2[32].Y = 0.009

	pointsOfFunctionPlot2[33].X = -5.67
	pointsOfFunctionPlot2[33].Y = 0.009

	pointsOfFunctionPlot2[34].X = -5.66
	pointsOfFunctionPlot2[34].Y = 0.009

	pointsOfFunctionPlot2[35].X = -5.65
	pointsOfFunctionPlot2[35].Y = 0.009

	pointsOfFunctionPlot2[36].X = -5.64
	pointsOfFunctionPlot2[36].Y = 0.009

	pointsOfFunctionPlot2[37].X = -5.63
	pointsOfFunctionPlot2[37].Y = 0.009

	pointsOfFunctionPlot2[38].X = -5.62
	pointsOfFunctionPlot2[38].Y = 0.009

	pointsOfFunctionPlot2[39].X = -5.61
	pointsOfFunctionPlot2[39].Y = 0.009

	pointsOfFunctionPlot2[40].X = -5.60
	pointsOfFunctionPlot2[40].Y = 0.009

	pointsOfFunctionPlot2[41].X = -5.59
	pointsOfFunctionPlot2[41].Y = 0.009

	pointsOfFunctionPlot2[42].X = -5.58
	pointsOfFunctionPlot2[42].Y = 0.009

	pointsOfFunctionPlot2[43].X = -5.57
	pointsOfFunctionPlot2[43].Y = 0.009

	pointsOfFunctionPlot2[44].X = -5.56
	pointsOfFunctionPlot2[44].Y = 0.009

	pointsOfFunctionPlot2[45].X = -5.55
	pointsOfFunctionPlot2[45].Y = 0.01

	pointsOfFunctionPlot2[46].X = -5.54
	pointsOfFunctionPlot2[46].Y = 0.01

	pointsOfFunctionPlot2[47].X = -5.53
	pointsOfFunctionPlot2[47].Y = 0.01

	pointsOfFunctionPlot2[48].X = -5.52
	pointsOfFunctionPlot2[48].Y = 0.01

	pointsOfFunctionPlot2[49].X = -5.51
	pointsOfFunctionPlot2[49].Y = 0.01

	pointsOfFunctionPlot2[50].X = -5.50
	pointsOfFunctionPlot2[50].Y = 0.01

	pointsOfFunctionPlot2[51].X = -5.49
	pointsOfFunctionPlot2[51].Y = 0.011

	pointsOfFunctionPlot2[52].X = -5.48
	pointsOfFunctionPlot2[52].Y = 0.011

	pointsOfFunctionPlot2[53].X = -5.47
	pointsOfFunctionPlot2[53].Y = 0.011

	pointsOfFunctionPlot2[54].X = -5.46
	pointsOfFunctionPlot2[54].Y = 0.011

	pointsOfFunctionPlot2[55].X = -5.45
	pointsOfFunctionPlot2[55].Y = 0.012

	pointsOfFunctionPlot2[56].X = -5.44
	pointsOfFunctionPlot2[56].Y = 0.012

	pointsOfFunctionPlot2[57].X = -5.43
	pointsOfFunctionPlot2[57].Y = 0.012

	pointsOfFunctionPlot2[58].X = -5.42
	pointsOfFunctionPlot2[58].Y = 0.012

	pointsOfFunctionPlot2[59].X = -5.41
	pointsOfFunctionPlot2[59].Y = 0.013

	pointsOfFunctionPlot2[60].X = -5.40
	pointsOfFunctionPlot2[60].Y = 0.013

	pointsOfFunctionPlot2[61].X = -5.39
	pointsOfFunctionPlot2[61].Y = 0.014

	pointsOfFunctionPlot2[62].X = -5.38
	pointsOfFunctionPlot2[62].Y = 0.014

	pointsOfFunctionPlot2[63].X = -5.37
	pointsOfFunctionPlot2[63].Y = 0.014

	pointsOfFunctionPlot2[64].X = -5.36
	pointsOfFunctionPlot2[64].Y = 0.015

	pointsOfFunctionPlot2[65].X = -5.35
	pointsOfFunctionPlot2[65].Y = 0.015

	pointsOfFunctionPlot2[66].X = -5.34
	pointsOfFunctionPlot2[66].Y = 0.016

	pointsOfFunctionPlot2[67].X = -5.33
	pointsOfFunctionPlot2[67].Y = 0.017

	pointsOfFunctionPlot2[68].X = -5.32
	pointsOfFunctionPlot2[68].Y = 0.017

	pointsOfFunctionPlot2[69].X = -5.31
	pointsOfFunctionPlot2[69].Y = 0.018

	pointsOfFunctionPlot2[70].X = -5.30
	pointsOfFunctionPlot2[70].Y = 0.019

	pointsOfFunctionPlot2[71].X = -5.29
	pointsOfFunctionPlot2[71].Y = 0.02

	pointsOfFunctionPlot2[72].X = -5.28
	pointsOfFunctionPlot2[72].Y = 0.02

	pointsOfFunctionPlot2[73].X = -5.27
	pointsOfFunctionPlot2[73].Y = 0.021

	pointsOfFunctionPlot2[74].X = -5.26
	pointsOfFunctionPlot2[74].Y = 0.022

	pointsOfFunctionPlot2[75].X = -5.25
	pointsOfFunctionPlot2[75].Y = 0.024

	pointsOfFunctionPlot2[76].X = -5.24
	pointsOfFunctionPlot2[76].Y = 0.025

	pointsOfFunctionPlot2[77].X = -5.23
	pointsOfFunctionPlot2[77].Y = 0.026

	pointsOfFunctionPlot2[78].X = -5.22
	pointsOfFunctionPlot2[78].Y = 0.028

	pointsOfFunctionPlot2[79].X = -5.21
	pointsOfFunctionPlot2[79].Y = 0.029

	pointsOfFunctionPlot2[80].X = -5.20
	pointsOfFunctionPlot2[80].Y = 0.031

	pointsOfFunctionPlot2[81].X = -5.19
	pointsOfFunctionPlot2[81].Y = 0.033

	pointsOfFunctionPlot2[82].X = -5.18
	pointsOfFunctionPlot2[82].Y = 0.035

	pointsOfFunctionPlot2[83].X = -5.17
	pointsOfFunctionPlot2[83].Y = 0.038

	pointsOfFunctionPlot2[84].X = -5.16
	pointsOfFunctionPlot2[84].Y = 0.041

	pointsOfFunctionPlot2[85].X = -5.15
	pointsOfFunctionPlot2[85].Y = 0.044

	pointsOfFunctionPlot2[86].X = -5.14
	pointsOfFunctionPlot2[86].Y = 0.048

	pointsOfFunctionPlot2[87].X = -5.13
	pointsOfFunctionPlot2[87].Y = 0.052

	pointsOfFunctionPlot2[88].X = -5.12
	pointsOfFunctionPlot2[88].Y = 0.057

	pointsOfFunctionPlot2[89].X = -5.11
	pointsOfFunctionPlot2[89].Y = 0.063

	pointsOfFunctionPlot2[90].X = -5.10
	pointsOfFunctionPlot2[90].Y = 0.071

	pointsOfFunctionPlot2[91].X = -5.09
	pointsOfFunctionPlot2[91].Y = 0.08

	pointsOfFunctionPlot2[92].X = -5.08
	pointsOfFunctionPlot2[92].Y = 0.091

	pointsOfFunctionPlot2[93].X = -5.07
	pointsOfFunctionPlot2[93].Y = 0.106

	pointsOfFunctionPlot2[94].X = -5.06
	pointsOfFunctionPlot2[94].Y = 0.126

	pointsOfFunctionPlot2[95].X = -5.05
	pointsOfFunctionPlot2[95].Y = 0.153

	pointsOfFunctionPlot2[96].X = -5.04
	pointsOfFunctionPlot2[96].Y = 0.195

	pointsOfFunctionPlot2[97].X = -5.03
	pointsOfFunctionPlot2[97].Y = 0.264

	pointsOfFunctionPlot2[98].X = -5.02
	pointsOfFunctionPlot2[98].Y = 0.402

	pointsOfFunctionPlot2[99].X = -5.01
	pointsOfFunctionPlot2[99].Y = 0.819





	pointsOfFunctionPlot3 := make(plotter.XYs, 100)

	pointsOfFunctionPlot3[0].X = -4.999
	pointsOfFunctionPlot3[0].Y = -8.347

	pointsOfFunctionPlot3[1].X = -4.99
	pointsOfFunctionPlot3[1].Y = -0.847

	pointsOfFunctionPlot3[2].X = -4.98
	pointsOfFunctionPlot3[2].Y = -0.431

	pointsOfFunctionPlot3[3].X = -4.97
	pointsOfFunctionPlot3[3].Y = -0.292

	pointsOfFunctionPlot3[4].X = -4.96
	pointsOfFunctionPlot3[4].Y = -0.223

	pointsOfFunctionPlot3[5].X = -4.95
	pointsOfFunctionPlot3[5].Y = -0.182

	pointsOfFunctionPlot3[6].X = -4.94
	pointsOfFunctionPlot3[6].Y = -0.154

	pointsOfFunctionPlot3[7].X = -4.93
	pointsOfFunctionPlot3[7].Y = -0.135

	pointsOfFunctionPlot3[8].X = -4.92
	pointsOfFunctionPlot3[8].Y = -0.12

	pointsOfFunctionPlot3[9].X = -4.91
	pointsOfFunctionPlot3[9].Y = -0.109

	pointsOfFunctionPlot3[10].X = -4.90
	pointsOfFunctionPlot3[10].Y = -0.1

	pointsOfFunctionPlot3[11].X = -4.89
	pointsOfFunctionPlot3[11].Y = -0.093

	pointsOfFunctionPlot3[12].X = -4.88
	pointsOfFunctionPlot3[12].Y = -0.087

	pointsOfFunctionPlot3[13].X = -4.87
	pointsOfFunctionPlot3[13].Y = -0.082

	pointsOfFunctionPlot3[14].X = -4.86
	pointsOfFunctionPlot3[14].Y = -0.077

	pointsOfFunctionPlot3[15].X = -4.85
	pointsOfFunctionPlot3[15].Y = -0.074

	pointsOfFunctionPlot3[16].X = -4.84
	pointsOfFunctionPlot3[16].Y = -0.071

	pointsOfFunctionPlot3[17].X = -4.83
	pointsOfFunctionPlot3[17].Y = -0.068

	pointsOfFunctionPlot3[18].X = -4.82
	pointsOfFunctionPlot3[18].Y = -0.066

	pointsOfFunctionPlot3[19].X = -4.81
	pointsOfFunctionPlot3[19].Y = -0.064

	pointsOfFunctionPlot3[20].X = -4.80
	pointsOfFunctionPlot3[20].Y = -0.062

	pointsOfFunctionPlot3[21].X = -4.79
	pointsOfFunctionPlot3[21].Y = -0.06

	pointsOfFunctionPlot3[22].X = -4.78
	pointsOfFunctionPlot3[22].Y = -0.059

	pointsOfFunctionPlot3[23].X = -4.77
	pointsOfFunctionPlot3[23].Y = -0.058

	pointsOfFunctionPlot3[24].X = -4.76
	pointsOfFunctionPlot3[24].Y = -0.057

	pointsOfFunctionPlot3[25].X = -4.75
	pointsOfFunctionPlot3[25].Y = -0.056

	pointsOfFunctionPlot3[26].X = -4.74
	pointsOfFunctionPlot3[26].Y = -0.055

	pointsOfFunctionPlot3[27].X = -4.73
	pointsOfFunctionPlot3[27].Y = -0.054

	pointsOfFunctionPlot3[28].X = -4.72
	pointsOfFunctionPlot3[28].Y = -0.054

	pointsOfFunctionPlot3[29].X = -4.71
	pointsOfFunctionPlot3[29].Y = -0.053

	pointsOfFunctionPlot3[30].X = -4.70
	pointsOfFunctionPlot3[30].Y = -0.053

	pointsOfFunctionPlot3[31].X = -4.69
	pointsOfFunctionPlot3[31].Y = -0.053

	pointsOfFunctionPlot3[32].X = -4.68
	pointsOfFunctionPlot3[32].Y = -0.053

	pointsOfFunctionPlot3[33].X = -4.67
	pointsOfFunctionPlot3[33].Y = -0.052

	pointsOfFunctionPlot3[34].X = -4.66
	pointsOfFunctionPlot3[34].Y = -0.052

	pointsOfFunctionPlot3[35].X = -4.65
	pointsOfFunctionPlot3[35].Y = -0.052

	pointsOfFunctionPlot3[36].X = -4.64
	pointsOfFunctionPlot3[36].Y = -0.052

	pointsOfFunctionPlot3[37].X = -4.63
	pointsOfFunctionPlot3[37].Y = -0.052

	pointsOfFunctionPlot3[38].X = -4.62
	pointsOfFunctionPlot3[38].Y = -0.053

	pointsOfFunctionPlot3[39].X = -4.61
	pointsOfFunctionPlot3[39].Y = -0.053

	pointsOfFunctionPlot3[40].X = -4.60
	pointsOfFunctionPlot3[40].Y = -0.053

	pointsOfFunctionPlot3[41].X = -4.59
	pointsOfFunctionPlot3[41].Y = -0.054

	pointsOfFunctionPlot3[42].X = -4.58
	pointsOfFunctionPlot3[42].Y = -0.054

	pointsOfFunctionPlot3[43].X = -4.57
	pointsOfFunctionPlot3[43].Y = -0.054

	pointsOfFunctionPlot3[44].X = -4.56
	pointsOfFunctionPlot3[44].Y = -0.055

	pointsOfFunctionPlot3[45].X = -4.55
	pointsOfFunctionPlot3[45].Y = -0.056

	pointsOfFunctionPlot3[46].X = -4.54
	pointsOfFunctionPlot3[46].Y = -0.056

	pointsOfFunctionPlot3[47].X = -4.53
	pointsOfFunctionPlot3[47].Y = -0.057

	pointsOfFunctionPlot3[48].X = -4.52
	pointsOfFunctionPlot3[48].Y = -0.058

	pointsOfFunctionPlot3[49].X = -4.51
	pointsOfFunctionPlot3[49].Y = -0.059

	pointsOfFunctionPlot3[50].X = -4.50
	pointsOfFunctionPlot3[50].Y = -0.06

	pointsOfFunctionPlot3[51].X = -4.49
	pointsOfFunctionPlot3[51].Y = -0.061

	pointsOfFunctionPlot3[52].X = -4.48
	pointsOfFunctionPlot3[52].Y = -0.062

	pointsOfFunctionPlot3[53].X = -4.47
	pointsOfFunctionPlot3[53].Y = -0.063

	pointsOfFunctionPlot3[54].X = -4.46
	pointsOfFunctionPlot3[54].Y = -0.064

	pointsOfFunctionPlot3[55].X = -4.45
	pointsOfFunctionPlot3[55].Y = -0.065

	pointsOfFunctionPlot3[56].X = -4.44
	pointsOfFunctionPlot3[56].Y = -0.067

	pointsOfFunctionPlot3[57].X = -4.43
	pointsOfFunctionPlot3[57].Y = -0.068

	pointsOfFunctionPlot3[58].X = -4.42
	pointsOfFunctionPlot3[58].Y = -0.07

	pointsOfFunctionPlot3[59].X = -4.41
	pointsOfFunctionPlot3[59].Y = -0.072

	pointsOfFunctionPlot3[60].X = -4.40
	pointsOfFunctionPlot3[60].Y = -0.074

	pointsOfFunctionPlot3[61].X = -4.39
	pointsOfFunctionPlot3[61].Y = -0.076

	pointsOfFunctionPlot3[62].X = -4.38
	pointsOfFunctionPlot3[62].Y = -0.078

	pointsOfFunctionPlot3[63].X = -4.37
	pointsOfFunctionPlot3[63].Y = -0.08

	pointsOfFunctionPlot3[64].X = -4.36
	pointsOfFunctionPlot3[64].Y = -0.082

	pointsOfFunctionPlot3[65].X = -4.35
	pointsOfFunctionPlot3[65].Y = -0.085

	pointsOfFunctionPlot3[66].X = -4.34
	pointsOfFunctionPlot3[66].Y = -0.088

	pointsOfFunctionPlot3[67].X = -4.33
	pointsOfFunctionPlot3[67].Y = -0.091

	pointsOfFunctionPlot3[68].X = -4.32
	pointsOfFunctionPlot3[68].Y = -0.094

	pointsOfFunctionPlot3[69].X = -4.31
	pointsOfFunctionPlot3[69].Y = -0.098

	pointsOfFunctionPlot3[70].X = -4.30
	pointsOfFunctionPlot3[70].Y = -0.101

	pointsOfFunctionPlot3[71].X = -4.29
	pointsOfFunctionPlot3[71].Y = -0.106

	pointsOfFunctionPlot3[72].X = -4.28
	pointsOfFunctionPlot3[72].Y = -0.11

	pointsOfFunctionPlot3[73].X = -4.27
	pointsOfFunctionPlot3[73].Y = -0.115

	pointsOfFunctionPlot3[74].X = -4.26
	pointsOfFunctionPlot3[74].Y = -0.12

	pointsOfFunctionPlot3[75].X = -4.25
	pointsOfFunctionPlot3[75].Y = -0.126

	pointsOfFunctionPlot3[76].X = -4.24
	pointsOfFunctionPlot3[76].Y = -0.132

	pointsOfFunctionPlot3[77].X = -4.23
	pointsOfFunctionPlot3[77].Y = -0.139

	pointsOfFunctionPlot3[78].X = -4.22
	pointsOfFunctionPlot3[78].Y = -0.146

	pointsOfFunctionPlot3[79].X = -4.21
	pointsOfFunctionPlot3[79].Y = -0.154

	pointsOfFunctionPlot3[80].X = -4.20
	pointsOfFunctionPlot3[80].Y = -0.164

	pointsOfFunctionPlot3[81].X = -4.19
	pointsOfFunctionPlot3[81].Y = -0.174

	pointsOfFunctionPlot3[82].X = -4.18
	pointsOfFunctionPlot3[82].Y = -0.185

	pointsOfFunctionPlot3[83].X = -4.17
	pointsOfFunctionPlot3[83].Y = -0.198

	pointsOfFunctionPlot3[84].X = -4.16
	pointsOfFunctionPlot3[84].Y = -0.212

	pointsOfFunctionPlot3[85].X = -4.15
	pointsOfFunctionPlot3[85].Y = -0.229

	pointsOfFunctionPlot3[86].X = -4.14
	pointsOfFunctionPlot3[86].Y = -0.248

	pointsOfFunctionPlot3[87].X = -4.13
	pointsOfFunctionPlot3[87].Y = -0.27

	pointsOfFunctionPlot3[88].X = -4.12
	pointsOfFunctionPlot3[88].Y = -0.296

	pointsOfFunctionPlot3[89].X = -4.11
	pointsOfFunctionPlot3[89].Y = -0.326

	pointsOfFunctionPlot3[90].X = -4.10
	pointsOfFunctionPlot3[90].Y = -0.363

	pointsOfFunctionPlot3[91].X = -4.09
	pointsOfFunctionPlot3[91].Y = -0.409

	pointsOfFunctionPlot3[92].X = -4.08
	pointsOfFunctionPlot3[92].Y = -0.466

	pointsOfFunctionPlot3[93].X = -4.07
	pointsOfFunctionPlot3[93].Y = -0.539

	pointsOfFunctionPlot3[94].X = -4.06
	pointsOfFunctionPlot3[94].Y = -0.637

	pointsOfFunctionPlot3[95].X = -4.05
	pointsOfFunctionPlot3[95].Y = -0.775

	pointsOfFunctionPlot3[96].X = -4.04
	pointsOfFunctionPlot3[96].Y = -0.983

	pointsOfFunctionPlot3[97].X = -4.03
	pointsOfFunctionPlot3[97].Y = -1.329

	pointsOfFunctionPlot3[98].X = -4.02
	pointsOfFunctionPlot3[98].Y = -2.022

	pointsOfFunctionPlot3[99].X = -4.01
	pointsOfFunctionPlot3[99].Y = -4.105





	pointsOfFunctionPlot4 := make(plotter.XYs, 100)

	pointsOfFunctionPlot4[0].X = -3.999
	pointsOfFunctionPlot4[0].Y = 41.729

	pointsOfFunctionPlot4[1].X = -3.99
	pointsOfFunctionPlot4[1].Y = 4.23

	pointsOfFunctionPlot4[2].X = -3.98
	pointsOfFunctionPlot4[2].Y = 2.148

	pointsOfFunctionPlot4[3].X = -3.97
	pointsOfFunctionPlot4[3].Y = 1.455

	pointsOfFunctionPlot4[4].X = -3.96
	pointsOfFunctionPlot4[4].Y = 1.109

	pointsOfFunctionPlot4[5].X = -3.95
	pointsOfFunctionPlot4[5].Y = 0.901

	pointsOfFunctionPlot4[6].X = -3.94
	pointsOfFunctionPlot4[6].Y = 0.764

	pointsOfFunctionPlot4[7].X = -3.93
	pointsOfFunctionPlot4[7].Y = 0.666

	pointsOfFunctionPlot4[8].X = -3.92
	pointsOfFunctionPlot4[8].Y = 0.593

	pointsOfFunctionPlot4[9].X = -3.91
	pointsOfFunctionPlot4[9].Y = 0.536

	pointsOfFunctionPlot4[10].X = -3.90
	pointsOfFunctionPlot4[10].Y = 0.491

	pointsOfFunctionPlot4[11].X = -3.89
	pointsOfFunctionPlot4[11].Y = 0.455

	pointsOfFunctionPlot4[12].X = -3.88
	pointsOfFunctionPlot4[12].Y = 0.425

	pointsOfFunctionPlot4[13].X = -3.87
	pointsOfFunctionPlot4[13].Y = 0.4

	pointsOfFunctionPlot4[14].X = -3.86
	pointsOfFunctionPlot4[14].Y = 0.378

	pointsOfFunctionPlot4[15].X = -3.85
	pointsOfFunctionPlot4[15].Y = 0.36

	pointsOfFunctionPlot4[16].X = -3.84
	pointsOfFunctionPlot4[16].Y = 0.344

	pointsOfFunctionPlot4[17].X = -3.83
	pointsOfFunctionPlot4[17].Y = 0.331

	pointsOfFunctionPlot4[18].X = -3.82
	pointsOfFunctionPlot4[18].Y = 0.319

	pointsOfFunctionPlot4[19].X = -3.81
	pointsOfFunctionPlot4[19].Y = 0.308

	pointsOfFunctionPlot4[20].X = -3.80
	pointsOfFunctionPlot4[20].Y = 0.299

	pointsOfFunctionPlot4[21].X = -3.79
	pointsOfFunctionPlot4[21].Y = 0.291

	pointsOfFunctionPlot4[22].X = -3.78
	pointsOfFunctionPlot4[22].Y = 0.284

	pointsOfFunctionPlot4[23].X = -3.77
	pointsOfFunctionPlot4[23].Y = 0.278

	pointsOfFunctionPlot4[24].X = -3.76
	pointsOfFunctionPlot4[24].Y = 0.272

	pointsOfFunctionPlot4[25].X = -3.75
	pointsOfFunctionPlot4[25].Y = 0.267

	pointsOfFunctionPlot4[26].X = -3.74
	pointsOfFunctionPlot4[26].Y = 0.263

	pointsOfFunctionPlot4[27].X = -3.73
	pointsOfFunctionPlot4[27].Y = 0.259

	pointsOfFunctionPlot4[28].X = -3.72
	pointsOfFunctionPlot4[28].Y = 0.256

	pointsOfFunctionPlot4[29].X = -3.71
	pointsOfFunctionPlot4[29].Y = 0.253

	pointsOfFunctionPlot4[30].X = -3.70
	pointsOfFunctionPlot4[30].Y = 0.251

	pointsOfFunctionPlot4[31].X = -3.69
	pointsOfFunctionPlot4[31].Y = 0.249

	pointsOfFunctionPlot4[32].X = -3.68
	pointsOfFunctionPlot4[32].Y = 0.248

	pointsOfFunctionPlot4[33].X = -3.67
	pointsOfFunctionPlot4[33].Y = 0.246

	pointsOfFunctionPlot4[34].X = -3.66
	pointsOfFunctionPlot4[34].Y = 0.246

	pointsOfFunctionPlot4[35].X = -3.65
	pointsOfFunctionPlot4[35].Y = 0.245

	pointsOfFunctionPlot4[36].X = -3.64
	pointsOfFunctionPlot4[36].Y = 0.245

	pointsOfFunctionPlot4[37].X = -3.63
	pointsOfFunctionPlot4[37].Y = 0.245

	pointsOfFunctionPlot4[38].X = -3.62
	pointsOfFunctionPlot4[38].Y = 0.245

	pointsOfFunctionPlot4[39].X = -3.61
	pointsOfFunctionPlot4[39].Y = 0.246

	pointsOfFunctionPlot4[40].X = -3.60
	pointsOfFunctionPlot4[40].Y = 0.246

	pointsOfFunctionPlot4[41].X = -3.59
	pointsOfFunctionPlot4[41].Y = 0.247

	pointsOfFunctionPlot4[42].X = -3.58
	pointsOfFunctionPlot4[42].Y = 0.249

	pointsOfFunctionPlot4[43].X = -3.57
	pointsOfFunctionPlot4[43].Y = 0.25

	pointsOfFunctionPlot4[44].X = -3.56
	pointsOfFunctionPlot4[44].Y = 0.252

	pointsOfFunctionPlot4[45].X = -3.55
	pointsOfFunctionPlot4[45].Y = 0.255

	pointsOfFunctionPlot4[46].X = -3.54
	pointsOfFunctionPlot4[46].Y = 0.257

	pointsOfFunctionPlot4[47].X = -3.53
	pointsOfFunctionPlot4[47].Y = 0.26

	pointsOfFunctionPlot4[48].X = -3.52
	pointsOfFunctionPlot4[48].Y = 0.263

	pointsOfFunctionPlot4[49].X = -3.51
	pointsOfFunctionPlot4[49].Y = 0.266

	pointsOfFunctionPlot4[50].X = -3.50
	pointsOfFunctionPlot4[50].Y = 0.27

	pointsOfFunctionPlot4[51].X = -3.49
	pointsOfFunctionPlot4[51].Y = 0.273

	pointsOfFunctionPlot4[52].X = -3.48
	pointsOfFunctionPlot4[52].Y = 0.278

	pointsOfFunctionPlot4[53].X = -3.47
	pointsOfFunctionPlot4[53].Y = 0.282

	pointsOfFunctionPlot4[54].X = -3.46
	pointsOfFunctionPlot4[54].Y = 0.287

	pointsOfFunctionPlot4[55].X = -3.45
	pointsOfFunctionPlot4[55].Y = 0.293

	pointsOfFunctionPlot4[56].X = -3.44
	pointsOfFunctionPlot4[56].Y = 0.298

	pointsOfFunctionPlot4[57].X = -3.43
	pointsOfFunctionPlot4[57].Y = 0.304

	pointsOfFunctionPlot4[58].X = -3.42
	pointsOfFunctionPlot4[58].Y = 0.311

	pointsOfFunctionPlot4[59].X = -3.41
	pointsOfFunctionPlot4[59].Y = 0.318

	pointsOfFunctionPlot4[60].X = -3.40
	pointsOfFunctionPlot4[60].Y = 0.325

	pointsOfFunctionPlot4[61].X = -3.39
	pointsOfFunctionPlot4[61].Y = 0.333

	pointsOfFunctionPlot4[62].X = -3.38
	pointsOfFunctionPlot4[62].Y = 0.342

	pointsOfFunctionPlot4[63].X = -3.37
	pointsOfFunctionPlot4[63].Y = 0.351

	pointsOfFunctionPlot4[64].X = -3.36
	pointsOfFunctionPlot4[64].Y = 0.361

	pointsOfFunctionPlot4[65].X = -3.35
	pointsOfFunctionPlot4[65].Y = 0.372

	pointsOfFunctionPlot4[66].X = -3.34
	pointsOfFunctionPlot4[66].Y = 0.383

	pointsOfFunctionPlot4[67].X = -3.33
	pointsOfFunctionPlot4[67].Y = 0.395

	pointsOfFunctionPlot4[68].X = -3.32
	pointsOfFunctionPlot4[68].Y = 0.409

	pointsOfFunctionPlot4[69].X = -3.31
	pointsOfFunctionPlot4[69].Y = 0.423

	pointsOfFunctionPlot4[70].X = -3.30
	pointsOfFunctionPlot4[70].Y = 0.438

	pointsOfFunctionPlot4[71].X = -3.29
	pointsOfFunctionPlot4[71].Y = 0.455

	pointsOfFunctionPlot4[72].X = -3.28
	pointsOfFunctionPlot4[72].Y = 0.472

	pointsOfFunctionPlot4[73].X = -3.27
	pointsOfFunctionPlot4[73].Y = 0.492

	pointsOfFunctionPlot4[74].X = -3.26
	pointsOfFunctionPlot4[74].Y = 0.513

	pointsOfFunctionPlot4[75].X = -3.25
	pointsOfFunctionPlot4[75].Y = 0.536

	pointsOfFunctionPlot4[76].X = -3.24
	pointsOfFunctionPlot4[76].Y = 0.561

	pointsOfFunctionPlot4[77].X = -3.23
	pointsOfFunctionPlot4[77].Y = 0.588

	pointsOfFunctionPlot4[78].X = -3.22
	pointsOfFunctionPlot4[78].Y = 0.618

	pointsOfFunctionPlot4[79].X = -3.21
	pointsOfFunctionPlot4[79].Y = 0.652

	pointsOfFunctionPlot4[80].X = -3.20
	pointsOfFunctionPlot4[80].Y = 0.689

	pointsOfFunctionPlot4[81].X = -3.19
	pointsOfFunctionPlot4[81].Y = 0.73

	pointsOfFunctionPlot4[82].X = -3.18
	pointsOfFunctionPlot4[82].Y = 0.775

	pointsOfFunctionPlot4[83].X = -3.17
	pointsOfFunctionPlot4[83].Y = 0.827

	pointsOfFunctionPlot4[84].X = -3.16
	pointsOfFunctionPlot4[84].Y = 0.885

	pointsOfFunctionPlot4[85].X = -3.15
	pointsOfFunctionPlot4[85].Y = 0.952

	pointsOfFunctionPlot4[86].X = -3.14
	pointsOfFunctionPlot4[86].Y = 1.028

	pointsOfFunctionPlot4[87].X = -3.13
	pointsOfFunctionPlot4[87].Y = 1.117

	pointsOfFunctionPlot4[88].X = -3.12
	pointsOfFunctionPlot4[88].Y = 1.22

	pointsOfFunctionPlot4[89].X = -3.11
	pointsOfFunctionPlot4[89].Y = 1.343

	pointsOfFunctionPlot4[90].X = -3.10
	pointsOfFunctionPlot4[90].Y = 1.492

	pointsOfFunctionPlot4[91].X = -3.09
	pointsOfFunctionPlot4[91].Y = 1.674

	pointsOfFunctionPlot4[92].X = -3.08
	pointsOfFunctionPlot4[92].Y = 1.902

	pointsOfFunctionPlot4[93].X = -3.07
	pointsOfFunctionPlot4[93].Y = 2.196

	pointsOfFunctionPlot4[94].X = -3.06
	pointsOfFunctionPlot4[94].Y = 2.59

	pointsOfFunctionPlot4[95].X = -3.05
	pointsOfFunctionPlot4[95].Y = 3.142

	pointsOfFunctionPlot4[96].X = -3.04
	pointsOfFunctionPlot4[96].Y = 3.972

	pointsOfFunctionPlot4[97].X = -3.03
	pointsOfFunctionPlot4[97].Y = 5.357

	pointsOfFunctionPlot4[98].X = -3.02
	pointsOfFunctionPlot4[98].Y = 8.131

	pointsOfFunctionPlot4[99].X = -3.01
	pointsOfFunctionPlot4[99].Y = 16.461





	pointsOfFunctionPlot5 := make(plotter.XYs, 97)

	pointsOfFunctionPlot5[0].X = -2.99
	pointsOfFunctionPlot5[0].Y = -16.879

	pointsOfFunctionPlot5[1].X = -2.98
	pointsOfFunctionPlot5[1].Y = -8.55

	pointsOfFunctionPlot5[2].X = -2.97
	pointsOfFunctionPlot5[2].Y = -5.776

	pointsOfFunctionPlot5[3].X = -2.96
	pointsOfFunctionPlot5[3].Y = -4.391

	pointsOfFunctionPlot5[4].X = -2.95
	pointsOfFunctionPlot5[4].Y = -3.562

	pointsOfFunctionPlot5[5].X = -2.94
	pointsOfFunctionPlot5[5].Y = -3.011

	pointsOfFunctionPlot5[6].X = -2.93
	pointsOfFunctionPlot5[6].Y = -2.619

	pointsOfFunctionPlot5[7].X = -2.92
	pointsOfFunctionPlot5[7].Y = -2.325

	pointsOfFunctionPlot5[8].X = -2.91
	pointsOfFunctionPlot5[8].Y = -2.098

	pointsOfFunctionPlot5[9].X = -2.90
	pointsOfFunctionPlot5[9].Y = -1.918

	pointsOfFunctionPlot5[10].X = -2.89
	pointsOfFunctionPlot5[10].Y = -1.771

	pointsOfFunctionPlot5[11].X = -2.88
	pointsOfFunctionPlot5[11].Y = -1.65

	pointsOfFunctionPlot5[12].X = -2.87
	pointsOfFunctionPlot5[12].Y = -1.548

	pointsOfFunctionPlot5[13].X = -2.86
	pointsOfFunctionPlot5[13].Y = -1.462

	pointsOfFunctionPlot5[14].X = -2.85
	pointsOfFunctionPlot5[14].Y = -1.387

	pointsOfFunctionPlot5[15].X = -2.84
	pointsOfFunctionPlot5[15].Y = -1.323

	pointsOfFunctionPlot5[16].X = -2.83
	pointsOfFunctionPlot5[16].Y = -1.268

	pointsOfFunctionPlot5[17].X = -2.82
	pointsOfFunctionPlot5[17].Y = -1.219

	pointsOfFunctionPlot5[18].X = -2.81
	pointsOfFunctionPlot5[18].Y = -1.176

	pointsOfFunctionPlot5[19].X = -2.80
	pointsOfFunctionPlot5[19].Y = -1.138

	pointsOfFunctionPlot5[20].X = -2.79
	pointsOfFunctionPlot5[20].Y = -1.105

	pointsOfFunctionPlot5[21].X = -2.78
	pointsOfFunctionPlot5[21].Y = -1.075

	pointsOfFunctionPlot5[22].X = -2.77
	pointsOfFunctionPlot5[22].Y = -1.048

	pointsOfFunctionPlot5[23].X = -2.76
	pointsOfFunctionPlot5[23].Y = -1.025

	pointsOfFunctionPlot5[24].X = -2.75
	pointsOfFunctionPlot5[24].Y = -1.004

	pointsOfFunctionPlot5[25].X = -2.74
	pointsOfFunctionPlot5[25].Y = -0.985

	pointsOfFunctionPlot5[26].X = -2.73
	pointsOfFunctionPlot5[26].Y = -0.969

	pointsOfFunctionPlot5[27].X = -2.72
	pointsOfFunctionPlot5[27].Y = -0.954

	pointsOfFunctionPlot5[28].X = -2.71
	pointsOfFunctionPlot5[28].Y = -0.942

	pointsOfFunctionPlot5[29].X = -2.70
	pointsOfFunctionPlot5[29].Y = -0.931

	pointsOfFunctionPlot5[30].X = -2.69
	pointsOfFunctionPlot5[30].Y = -0.921

	pointsOfFunctionPlot5[31].X = -2.68
	pointsOfFunctionPlot5[31].Y = -0.913

	pointsOfFunctionPlot5[32].X = -2.67
	pointsOfFunctionPlot5[32].Y = -0.906

	pointsOfFunctionPlot5[33].X = -2.66
	pointsOfFunctionPlot5[33].Y = -0.9

	pointsOfFunctionPlot5[34].X = -2.65
	pointsOfFunctionPlot5[34].Y = -0.895

	pointsOfFunctionPlot5[35].X = -2.64
	pointsOfFunctionPlot5[35].Y = -0.892

	pointsOfFunctionPlot5[36].X = -2.63
	pointsOfFunctionPlot5[36].Y = -0.889

	pointsOfFunctionPlot5[37].X = -2.62
	pointsOfFunctionPlot5[37].Y = -0.888

	pointsOfFunctionPlot5[38].X = -2.61
	pointsOfFunctionPlot5[38].Y = -0.888

	pointsOfFunctionPlot5[39].X = -2.60
	pointsOfFunctionPlot5[39].Y = -0.888

	pointsOfFunctionPlot5[40].X = -2.59
	pointsOfFunctionPlot5[40].Y = -0.89

	pointsOfFunctionPlot5[41].X = -2.58
	pointsOfFunctionPlot5[41].Y = -0.892

	pointsOfFunctionPlot5[42].X = -2.57
	pointsOfFunctionPlot5[42].Y = -0.895

	pointsOfFunctionPlot5[43].X = -2.56
	pointsOfFunctionPlot5[43].Y = -0.9

	pointsOfFunctionPlot5[44].X = -2.55
	pointsOfFunctionPlot5[44].Y = -0.905

	pointsOfFunctionPlot5[45].X = -2.54
	pointsOfFunctionPlot5[45].Y = -0.911

	pointsOfFunctionPlot5[46].X = -2.53
	pointsOfFunctionPlot5[46].Y = -0.918

	pointsOfFunctionPlot5[47].X = -2.52
	pointsOfFunctionPlot5[47].Y = -0.926

	pointsOfFunctionPlot5[48].X = -2.51
	pointsOfFunctionPlot5[48].Y = -0.935

	pointsOfFunctionPlot5[49].X = -2.50
	pointsOfFunctionPlot5[49].Y = -0.945

	pointsOfFunctionPlot5[50].X = -2.49
	pointsOfFunctionPlot5[50].Y = -0.956

	pointsOfFunctionPlot5[51].X = -2.48
	pointsOfFunctionPlot5[51].Y = -0.968

	pointsOfFunctionPlot5[52].X = -2.47
	pointsOfFunctionPlot5[52].Y = -0.981

	pointsOfFunctionPlot5[53].X = -2.46
	pointsOfFunctionPlot5[53].Y = -0.995

	pointsOfFunctionPlot5[54].X = -2.45
	pointsOfFunctionPlot5[54].Y = -1.01

	pointsOfFunctionPlot5[55].X = -2.44
	pointsOfFunctionPlot5[55].Y = -1.027

	pointsOfFunctionPlot5[56].X = -2.43
	pointsOfFunctionPlot5[56].Y = -1.045

	pointsOfFunctionPlot5[57].X = -2.42
	pointsOfFunctionPlot5[57].Y = -1.064

	pointsOfFunctionPlot5[58].X = -2.41
	pointsOfFunctionPlot5[58].Y = -1.085

	pointsOfFunctionPlot5[59].X = -2.40
	pointsOfFunctionPlot5[59].Y = -1.108

	pointsOfFunctionPlot5[60].X = -2.39
	pointsOfFunctionPlot5[60].Y = -1.132

	pointsOfFunctionPlot5[61].X = -2.38
	pointsOfFunctionPlot5[61].Y = -1.157

	pointsOfFunctionPlot5[62].X = -2.37
	pointsOfFunctionPlot5[62].Y = -1.185

	pointsOfFunctionPlot5[63].X = -2.36
	pointsOfFunctionPlot5[63].Y = -1.215

	pointsOfFunctionPlot5[64].X = -2.35
	pointsOfFunctionPlot5[64].Y = -1.247

	pointsOfFunctionPlot5[65].X = -2.34
	pointsOfFunctionPlot5[65].Y = -1.281

	pointsOfFunctionPlot5[66].X = -2.33
	pointsOfFunctionPlot5[66].Y = -1.318

	pointsOfFunctionPlot5[67].X = -2.32
	pointsOfFunctionPlot5[67].Y = -1.358

	pointsOfFunctionPlot5[68].X = -2.31
	pointsOfFunctionPlot5[68].Y = -1.4

	pointsOfFunctionPlot5[69].X = -2.30
	pointsOfFunctionPlot5[69].Y = -1.447

	pointsOfFunctionPlot5[70].X = -2.29
	pointsOfFunctionPlot5[70].Y = -1.497

	pointsOfFunctionPlot5[71].X = -2.28
	pointsOfFunctionPlot5[71].Y = -1.551

	pointsOfFunctionPlot5[72].X = -2.27
	pointsOfFunctionPlot5[72].Y = -1.609

	pointsOfFunctionPlot5[73].X = -2.26
	pointsOfFunctionPlot5[73].Y = -1.673

	pointsOfFunctionPlot5[74].X = -2.25
	pointsOfFunctionPlot5[74].Y = -1.742

	pointsOfFunctionPlot5[75].X = -2.24
	pointsOfFunctionPlot5[75].Y = -1.818

	pointsOfFunctionPlot5[76].X = -2.23
	pointsOfFunctionPlot5[76].Y = -1.901

	pointsOfFunctionPlot5[77].X = -2.22
	pointsOfFunctionPlot5[77].Y = -1.992

	pointsOfFunctionPlot5[78].X = -2.21
	pointsOfFunctionPlot5[78].Y = -2.093

	pointsOfFunctionPlot5[79].X = -2.20
	pointsOfFunctionPlot5[79].Y = -2.204

	pointsOfFunctionPlot5[80].X = -2.19
	pointsOfFunctionPlot5[80].Y = -2.328

	pointsOfFunctionPlot5[81].X = -2.18
	pointsOfFunctionPlot5[81].Y = -2.467

	pointsOfFunctionPlot5[82].X = -2.17
	pointsOfFunctionPlot5[82].Y = -2.623

	pointsOfFunctionPlot5[83].X = -2.16
	pointsOfFunctionPlot5[83].Y = -2.799

	pointsOfFunctionPlot5[84].X = -2.15
	pointsOfFunctionPlot5[84].Y = -2.999

	pointsOfFunctionPlot5[85].X = -2.14
	pointsOfFunctionPlot5[85].Y = -3.229

	pointsOfFunctionPlot5[86].X = -2.13
	pointsOfFunctionPlot5[86].Y = -3.496

	pointsOfFunctionPlot5[87].X = -2.12
	pointsOfFunctionPlot5[87].Y = -3.809

	pointsOfFunctionPlot5[88].X = -2.11
	pointsOfFunctionPlot5[88].Y = -4.179

	pointsOfFunctionPlot5[89].X = -2.10
	pointsOfFunctionPlot5[89].Y = -4.626

	pointsOfFunctionPlot5[90].X = -2.09
	pointsOfFunctionPlot5[90].Y = -5.173

	pointsOfFunctionPlot5[91].X = -2.08
	pointsOfFunctionPlot5[91].Y = -5.859

	pointsOfFunctionPlot5[92].X = -2.07
	pointsOfFunctionPlot5[92].Y = -6.743

	pointsOfFunctionPlot5[93].X = -2.06
	pointsOfFunctionPlot5[93].Y = -7.925

	pointsOfFunctionPlot5[94].X = -2.05
	pointsOfFunctionPlot5[94].Y = -9.583

	pointsOfFunctionPlot5[95].X = -2.04
	pointsOfFunctionPlot5[95].Y = -12.074

	pointsOfFunctionPlot5[96].X = -2.03
	pointsOfFunctionPlot5[96].Y = -16.232










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of Gamma function"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot1)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("Gamma(x)", plotLine)



	plotLine, err = plotter.NewLine(pointsOfFunctionPlot2)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)



	plotLine, err = plotter.NewLine(pointsOfFunctionPlot3)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)



	plotLine, err = plotter.NewLine(pointsOfFunctionPlot4)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)



	plotLine, err = plotter.NewLine(pointsOfFunctionPlot5)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)



	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"Gamma-function-plot-03.png"); err != nil {

		panic(err)
	}
}
