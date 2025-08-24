package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function 5^x.

	pointsOfFunctionPlot := make(plotter.XYs, 1_435)

	pointsOfFunctionPlot[0].X = -10.0
	pointsOfFunctionPlot[0].Y = 0.0

	pointsOfFunctionPlot[1].X = -4.33
	pointsOfFunctionPlot[1].Y = 0.0

	pointsOfFunctionPlot[2].X = -4.32
	pointsOfFunctionPlot[2].Y = 0.001

	pointsOfFunctionPlot[3].X = -4.31
	pointsOfFunctionPlot[3].Y = 0.001

	pointsOfFunctionPlot[4].X = -4.3
	pointsOfFunctionPlot[4].Y = 0.001

	pointsOfFunctionPlot[5].X = -4.29
	pointsOfFunctionPlot[5].Y = 0.001

	pointsOfFunctionPlot[6].X = -4.28
	pointsOfFunctionPlot[6].Y = 0.001

	pointsOfFunctionPlot[7].X = -4.27
	pointsOfFunctionPlot[7].Y = 0.001

	pointsOfFunctionPlot[8].X = -4.26
	pointsOfFunctionPlot[8].Y = 0.001

	pointsOfFunctionPlot[9].X = -4.25
	pointsOfFunctionPlot[9].Y = 0.001

	pointsOfFunctionPlot[10].X = -4.24
	pointsOfFunctionPlot[10].Y = 0.001

	pointsOfFunctionPlot[11].X = -4.23
	pointsOfFunctionPlot[11].Y = 0.001

	pointsOfFunctionPlot[12].X = -4.22
	pointsOfFunctionPlot[12].Y = 0.001

	pointsOfFunctionPlot[13].X = -4.21
	pointsOfFunctionPlot[13].Y = 0.001

	pointsOfFunctionPlot[14].X = -4.2
	pointsOfFunctionPlot[14].Y = 0.001

	pointsOfFunctionPlot[15].X = -4.19
	pointsOfFunctionPlot[15].Y = 0.001

	pointsOfFunctionPlot[16].X = -4.18
	pointsOfFunctionPlot[16].Y = 0.001

	pointsOfFunctionPlot[17].X = -4.17
	pointsOfFunctionPlot[17].Y = 0.001

	pointsOfFunctionPlot[18].X = -4.16
	pointsOfFunctionPlot[18].Y = 0.001

	pointsOfFunctionPlot[19].X = -4.15
	pointsOfFunctionPlot[19].Y = 0.001

	pointsOfFunctionPlot[20].X = -4.14
	pointsOfFunctionPlot[20].Y = 0.001

	pointsOfFunctionPlot[21].X = -4.13
	pointsOfFunctionPlot[21].Y = 0.001

	pointsOfFunctionPlot[22].X = -4.12
	pointsOfFunctionPlot[22].Y = 0.001

	pointsOfFunctionPlot[23].X = -4.11
	pointsOfFunctionPlot[23].Y = 0.001

	pointsOfFunctionPlot[24].X = -4.1
	pointsOfFunctionPlot[24].Y = 0.001

	pointsOfFunctionPlot[25].X = -4.09
	pointsOfFunctionPlot[25].Y = 0.001

	pointsOfFunctionPlot[26].X = -4.08
	pointsOfFunctionPlot[26].Y = 0.001

	pointsOfFunctionPlot[27].X = -4.07
	pointsOfFunctionPlot[27].Y = 0.001

	pointsOfFunctionPlot[28].X = -4.06
	pointsOfFunctionPlot[28].Y = 0.001

	pointsOfFunctionPlot[29].X = -4.05
	pointsOfFunctionPlot[29].Y = 0.001

	pointsOfFunctionPlot[30].X = -4.04
	pointsOfFunctionPlot[30].Y = 0.001

	pointsOfFunctionPlot[31].X = -4.03
	pointsOfFunctionPlot[31].Y = 0.001

	pointsOfFunctionPlot[32].X = -4.02
	pointsOfFunctionPlot[32].Y = 0.001

	pointsOfFunctionPlot[33].X = -4.01
	pointsOfFunctionPlot[33].Y = 0.001

	pointsOfFunctionPlot[34].X = -4.0
	pointsOfFunctionPlot[34].Y = 0.001

	pointsOfFunctionPlot[35].X = -3.99
	pointsOfFunctionPlot[35].Y = 0.001

	pointsOfFunctionPlot[36].X = -3.98
	pointsOfFunctionPlot[36].Y = 0.001

	pointsOfFunctionPlot[37].X = -3.97
	pointsOfFunctionPlot[37].Y = 0.001

	pointsOfFunctionPlot[38].X = -3.96
	pointsOfFunctionPlot[38].Y = 0.001

	pointsOfFunctionPlot[39].X = -3.95
	pointsOfFunctionPlot[39].Y = 0.001

	pointsOfFunctionPlot[40].X = -3.94
	pointsOfFunctionPlot[40].Y = 0.001

	pointsOfFunctionPlot[41].X = -3.93
	pointsOfFunctionPlot[41].Y = 0.001

	pointsOfFunctionPlot[42].X = -3.92
	pointsOfFunctionPlot[42].Y = 0.001

	pointsOfFunctionPlot[43].X = -3.91
	pointsOfFunctionPlot[43].Y = 0.001

	pointsOfFunctionPlot[44].X = -3.9
	pointsOfFunctionPlot[44].Y = 0.001

	pointsOfFunctionPlot[45].X = -3.89
	pointsOfFunctionPlot[45].Y = 0.001

	pointsOfFunctionPlot[46].X = -3.88
	pointsOfFunctionPlot[46].Y = 0.001

	pointsOfFunctionPlot[47].X = -3.87
	pointsOfFunctionPlot[47].Y = 0.002

	pointsOfFunctionPlot[48].X = -3.86
	pointsOfFunctionPlot[48].Y = 0.002

	pointsOfFunctionPlot[49].X = -3.85
	pointsOfFunctionPlot[49].Y = 0.002

	pointsOfFunctionPlot[50].X = -3.84
	pointsOfFunctionPlot[50].Y = 0.002

	pointsOfFunctionPlot[51].X = -3.83
	pointsOfFunctionPlot[51].Y = 0.002

	pointsOfFunctionPlot[52].X = -3.82
	pointsOfFunctionPlot[52].Y = 0.002

	pointsOfFunctionPlot[53].X = -3.81
	pointsOfFunctionPlot[53].Y = 0.002

	pointsOfFunctionPlot[54].X = -3.8
	pointsOfFunctionPlot[54].Y = 0.002

	pointsOfFunctionPlot[55].X = -3.79
	pointsOfFunctionPlot[55].Y = 0.002

	pointsOfFunctionPlot[56].X = -3.78
	pointsOfFunctionPlot[56].Y = 0.002

	pointsOfFunctionPlot[57].X = -3.77
	pointsOfFunctionPlot[57].Y = 0.002

	pointsOfFunctionPlot[58].X = -3.76
	pointsOfFunctionPlot[58].Y = 0.002

	pointsOfFunctionPlot[59].X = -3.75
	pointsOfFunctionPlot[59].Y = 0.002

	pointsOfFunctionPlot[60].X = -3.74
	pointsOfFunctionPlot[60].Y = 0.002

	pointsOfFunctionPlot[61].X = -3.73
	pointsOfFunctionPlot[61].Y = 0.002

	pointsOfFunctionPlot[62].X = -3.72
	pointsOfFunctionPlot[62].Y = 0.002

	pointsOfFunctionPlot[63].X = -3.71
	pointsOfFunctionPlot[63].Y = 0.002

	pointsOfFunctionPlot[64].X = -3.7
	pointsOfFunctionPlot[64].Y = 0.002

	pointsOfFunctionPlot[65].X = -3.69
	pointsOfFunctionPlot[65].Y = 0.002

	pointsOfFunctionPlot[66].X = -3.68
	pointsOfFunctionPlot[66].Y = 0.002

	pointsOfFunctionPlot[67].X = -3.67
	pointsOfFunctionPlot[67].Y = 0.002

	pointsOfFunctionPlot[68].X = -3.66
	pointsOfFunctionPlot[68].Y = 0.002

	pointsOfFunctionPlot[69].X = -3.65
	pointsOfFunctionPlot[69].Y = 0.002

	pointsOfFunctionPlot[70].X = -3.64
	pointsOfFunctionPlot[70].Y = 0.002

	pointsOfFunctionPlot[71].X = -3.63
	pointsOfFunctionPlot[71].Y = 0.002

	pointsOfFunctionPlot[72].X = -3.62
	pointsOfFunctionPlot[72].Y = 0.002

	pointsOfFunctionPlot[73].X = -3.61
	pointsOfFunctionPlot[73].Y = 0.003

	pointsOfFunctionPlot[74].X = -3.6
	pointsOfFunctionPlot[74].Y = 0.003

	pointsOfFunctionPlot[75].X = -3.59
	pointsOfFunctionPlot[75].Y = 0.003

	pointsOfFunctionPlot[76].X = -3.58
	pointsOfFunctionPlot[76].Y = 0.003

	pointsOfFunctionPlot[77].X = -3.57
	pointsOfFunctionPlot[77].Y = 0.003

	pointsOfFunctionPlot[78].X = -3.56
	pointsOfFunctionPlot[78].Y = 0.003

	pointsOfFunctionPlot[79].X = -3.55
	pointsOfFunctionPlot[79].Y = 0.003

	pointsOfFunctionPlot[80].X = -3.54
	pointsOfFunctionPlot[80].Y = 0.003

	pointsOfFunctionPlot[81].X = -3.53
	pointsOfFunctionPlot[81].Y = 0.003

	pointsOfFunctionPlot[82].X = -3.52
	pointsOfFunctionPlot[82].Y = 0.003

	pointsOfFunctionPlot[83].X = -3.51
	pointsOfFunctionPlot[83].Y = 0.003

	pointsOfFunctionPlot[84].X = -3.5
	pointsOfFunctionPlot[84].Y = 0.003

	pointsOfFunctionPlot[85].X = -3.49
	pointsOfFunctionPlot[85].Y = 0.003

	pointsOfFunctionPlot[86].X = -3.48
	pointsOfFunctionPlot[86].Y = 0.003

	pointsOfFunctionPlot[87].X = -3.47
	pointsOfFunctionPlot[87].Y = 0.003

	pointsOfFunctionPlot[88].X = -3.46
	pointsOfFunctionPlot[88].Y = 0.003

	pointsOfFunctionPlot[89].X = -3.45
	pointsOfFunctionPlot[89].Y = 0.003

	pointsOfFunctionPlot[90].X = -3.44
	pointsOfFunctionPlot[90].Y = 0.003

	pointsOfFunctionPlot[91].X = -3.43
	pointsOfFunctionPlot[91].Y = 0.004

	pointsOfFunctionPlot[92].X = -3.42
	pointsOfFunctionPlot[92].Y = 0.004

	pointsOfFunctionPlot[93].X = -3.41
	pointsOfFunctionPlot[93].Y = 0.004

	pointsOfFunctionPlot[94].X = -3.4
	pointsOfFunctionPlot[94].Y = 0.004

	pointsOfFunctionPlot[95].X = -3.39
	pointsOfFunctionPlot[95].Y = 0.004

	pointsOfFunctionPlot[96].X = -3.38
	pointsOfFunctionPlot[96].Y = 0.004

	pointsOfFunctionPlot[97].X = -3.37
	pointsOfFunctionPlot[97].Y = 0.004

	pointsOfFunctionPlot[98].X = -3.36
	pointsOfFunctionPlot[98].Y = 0.004

	pointsOfFunctionPlot[99].X = -3.35
	pointsOfFunctionPlot[99].Y = 0.004

	pointsOfFunctionPlot[100].X = -3.34
	pointsOfFunctionPlot[100].Y = 0.004

	pointsOfFunctionPlot[101].X = -3.33
	pointsOfFunctionPlot[101].Y = 0.004

	pointsOfFunctionPlot[102].X = -3.32
	pointsOfFunctionPlot[102].Y = 0.004

	pointsOfFunctionPlot[103].X = -3.31
	pointsOfFunctionPlot[103].Y = 0.004

	pointsOfFunctionPlot[104].X = -3.3
	pointsOfFunctionPlot[104].Y = 0.004

	pointsOfFunctionPlot[105].X = -3.29
	pointsOfFunctionPlot[105].Y = 0.005

	pointsOfFunctionPlot[106].X = -3.28
	pointsOfFunctionPlot[106].Y = 0.005

	pointsOfFunctionPlot[107].X = -3.27
	pointsOfFunctionPlot[107].Y = 0.005

	pointsOfFunctionPlot[108].X = -3.26
	pointsOfFunctionPlot[108].Y = 0.005

	pointsOfFunctionPlot[109].X = -3.25
	pointsOfFunctionPlot[109].Y = 0.005

	pointsOfFunctionPlot[110].X = -3.24
	pointsOfFunctionPlot[110].Y = 0.005

	pointsOfFunctionPlot[111].X = -3.23
	pointsOfFunctionPlot[111].Y = 0.005

	pointsOfFunctionPlot[112].X = -3.22
	pointsOfFunctionPlot[112].Y = 0.005

	pointsOfFunctionPlot[113].X = -3.21
	pointsOfFunctionPlot[113].Y = 0.005

	pointsOfFunctionPlot[114].X = -3.2
	pointsOfFunctionPlot[114].Y = 0.005

	pointsOfFunctionPlot[115].X = -3.19
	pointsOfFunctionPlot[115].Y = 0.005

	pointsOfFunctionPlot[116].X = -3.18
	pointsOfFunctionPlot[116].Y = 0.006

	pointsOfFunctionPlot[117].X = -3.17
	pointsOfFunctionPlot[117].Y = 0.006

	pointsOfFunctionPlot[118].X = -3.16
	pointsOfFunctionPlot[118].Y = 0.006

	pointsOfFunctionPlot[119].X = -3.15
	pointsOfFunctionPlot[119].Y = 0.006

	pointsOfFunctionPlot[120].X = -3.14
	pointsOfFunctionPlot[120].Y = 0.006

	pointsOfFunctionPlot[121].X = -3.13
	pointsOfFunctionPlot[121].Y = 0.006

	pointsOfFunctionPlot[122].X = -3.12
	pointsOfFunctionPlot[122].Y = 0.006

	pointsOfFunctionPlot[123].X = -3.11
	pointsOfFunctionPlot[123].Y = 0.006

	pointsOfFunctionPlot[124].X = -3.1
	pointsOfFunctionPlot[124].Y = 0.006

	pointsOfFunctionPlot[125].X = -3.09
	pointsOfFunctionPlot[125].Y = 0.006

	pointsOfFunctionPlot[126].X = -3.08
	pointsOfFunctionPlot[126].Y = 0.007

	pointsOfFunctionPlot[127].X = -3.07
	pointsOfFunctionPlot[127].Y = 0.007

	pointsOfFunctionPlot[128].X = -3.06
	pointsOfFunctionPlot[128].Y = 0.007

	pointsOfFunctionPlot[129].X = -3.05
	pointsOfFunctionPlot[129].Y = 0.007

	pointsOfFunctionPlot[130].X = -3.04
	pointsOfFunctionPlot[130].Y = 0.007

	pointsOfFunctionPlot[131].X = -3.03
	pointsOfFunctionPlot[131].Y = 0.007

	pointsOfFunctionPlot[132].X = -3.02
	pointsOfFunctionPlot[132].Y = 0.007

	pointsOfFunctionPlot[133].X = -3.01
	pointsOfFunctionPlot[133].Y = 0.007

	pointsOfFunctionPlot[134].X = -3.0
	pointsOfFunctionPlot[134].Y = 0.008

	pointsOfFunctionPlot[135].X = -2.99
	pointsOfFunctionPlot[135].Y = 0.008

	pointsOfFunctionPlot[136].X = -2.98
	pointsOfFunctionPlot[136].Y = 0.008

	pointsOfFunctionPlot[137].X = -2.97
	pointsOfFunctionPlot[137].Y = 0.008

	pointsOfFunctionPlot[138].X = -2.96
	pointsOfFunctionPlot[138].Y = 0.008

	pointsOfFunctionPlot[139].X = -2.95
	pointsOfFunctionPlot[139].Y = 0.008

	pointsOfFunctionPlot[140].X = -2.94
	pointsOfFunctionPlot[140].Y = 0.008

	pointsOfFunctionPlot[141].X = -2.93
	pointsOfFunctionPlot[141].Y = 0.009

	pointsOfFunctionPlot[142].X = -2.92
	pointsOfFunctionPlot[142].Y = 0.009

	pointsOfFunctionPlot[143].X = -2.91
	pointsOfFunctionPlot[143].Y = 0.009

	pointsOfFunctionPlot[144].X = -2.9
	pointsOfFunctionPlot[144].Y = 0.009

	pointsOfFunctionPlot[145].X = -2.89
	pointsOfFunctionPlot[145].Y = 0.009

	pointsOfFunctionPlot[146].X = -2.88
	pointsOfFunctionPlot[146].Y = 0.009

	pointsOfFunctionPlot[147].X = -2.87
	pointsOfFunctionPlot[147].Y = 0.009

	pointsOfFunctionPlot[148].X = -2.86
	pointsOfFunctionPlot[148].Y = 0.01

	pointsOfFunctionPlot[149].X = -2.85
	pointsOfFunctionPlot[149].Y = 0.01

	pointsOfFunctionPlot[150].X = -2.84
	pointsOfFunctionPlot[150].Y = 0.01

	pointsOfFunctionPlot[151].X = -2.83
	pointsOfFunctionPlot[151].Y = 0.01

	pointsOfFunctionPlot[152].X = -2.82
	pointsOfFunctionPlot[152].Y = 0.01

	pointsOfFunctionPlot[153].X = -2.81
	pointsOfFunctionPlot[153].Y = 0.01

	pointsOfFunctionPlot[154].X = -2.8
	pointsOfFunctionPlot[154].Y = 0.011

	pointsOfFunctionPlot[155].X = -2.79
	pointsOfFunctionPlot[155].Y = 0.011

	pointsOfFunctionPlot[156].X = -2.78
	pointsOfFunctionPlot[156].Y = 0.011

	pointsOfFunctionPlot[157].X = -2.77
	pointsOfFunctionPlot[157].Y = 0.011

	pointsOfFunctionPlot[158].X = -2.76
	pointsOfFunctionPlot[158].Y = 0.011

	pointsOfFunctionPlot[159].X = -2.75
	pointsOfFunctionPlot[159].Y = 0.012

	pointsOfFunctionPlot[160].X = -2.74
	pointsOfFunctionPlot[160].Y = 0.012

	pointsOfFunctionPlot[161].X = -2.73
	pointsOfFunctionPlot[161].Y = 0.012

	pointsOfFunctionPlot[162].X = -2.72
	pointsOfFunctionPlot[162].Y = 0.012

	pointsOfFunctionPlot[163].X = -2.71
	pointsOfFunctionPlot[163].Y = 0.012

	pointsOfFunctionPlot[164].X = -2.7
	pointsOfFunctionPlot[164].Y = 0.013

	pointsOfFunctionPlot[165].X = -2.69
	pointsOfFunctionPlot[165].Y = 0.013

	pointsOfFunctionPlot[166].X = -2.68
	pointsOfFunctionPlot[166].Y = 0.013

	pointsOfFunctionPlot[167].X = -2.67
	pointsOfFunctionPlot[167].Y = 0.013

	pointsOfFunctionPlot[168].X = -2.66
	pointsOfFunctionPlot[168].Y = 0.013

	pointsOfFunctionPlot[169].X = -2.65
	pointsOfFunctionPlot[169].Y = 0.014

	pointsOfFunctionPlot[170].X = -2.64
	pointsOfFunctionPlot[170].Y = 0.014

	pointsOfFunctionPlot[171].X = -2.63
	pointsOfFunctionPlot[171].Y = 0.014

	pointsOfFunctionPlot[172].X = -2.62
	pointsOfFunctionPlot[172].Y = 0.014

	pointsOfFunctionPlot[173].X = -2.61
	pointsOfFunctionPlot[173].Y = 0.015

	pointsOfFunctionPlot[174].X = -2.6
	pointsOfFunctionPlot[174].Y = 0.015

	pointsOfFunctionPlot[175].X = -2.59
	pointsOfFunctionPlot[175].Y = 0.015

	pointsOfFunctionPlot[176].X = -2.58
	pointsOfFunctionPlot[176].Y = 0.015

	pointsOfFunctionPlot[177].X = -2.57
	pointsOfFunctionPlot[177].Y = 0.016

	pointsOfFunctionPlot[178].X = -2.56
	pointsOfFunctionPlot[178].Y = 0.016

	pointsOfFunctionPlot[179].X = -2.55
	pointsOfFunctionPlot[179].Y = 0.016

	pointsOfFunctionPlot[180].X = -2.54
	pointsOfFunctionPlot[180].Y = 0.016

	pointsOfFunctionPlot[181].X = -2.53
	pointsOfFunctionPlot[181].Y = 0.017

	pointsOfFunctionPlot[182].X = -2.52
	pointsOfFunctionPlot[182].Y = 0.017

	pointsOfFunctionPlot[183].X = -2.51
	pointsOfFunctionPlot[183].Y = 0.017

	pointsOfFunctionPlot[184].X = -2.5
	pointsOfFunctionPlot[184].Y = 0.017

	pointsOfFunctionPlot[185].X = -2.49
	pointsOfFunctionPlot[185].Y = 0.018

	pointsOfFunctionPlot[186].X = -2.48
	pointsOfFunctionPlot[186].Y = 0.018

	pointsOfFunctionPlot[187].X = -2.47
	pointsOfFunctionPlot[187].Y = 0.018

	pointsOfFunctionPlot[188].X = -2.46
	pointsOfFunctionPlot[188].Y = 0.019

	pointsOfFunctionPlot[189].X = -2.45
	pointsOfFunctionPlot[189].Y = 0.019

	pointsOfFunctionPlot[190].X = -2.44
	pointsOfFunctionPlot[190].Y = 0.019

	pointsOfFunctionPlot[191].X = -2.43
	pointsOfFunctionPlot[191].Y = 0.019

	pointsOfFunctionPlot[192].X = -2.42
	pointsOfFunctionPlot[192].Y = 0.02

	pointsOfFunctionPlot[193].X = -2.41
	pointsOfFunctionPlot[193].Y = 0.02

	pointsOfFunctionPlot[194].X = -2.4
	pointsOfFunctionPlot[194].Y = 0.021

	pointsOfFunctionPlot[195].X = -2.39
	pointsOfFunctionPlot[195].Y = 0.021

	pointsOfFunctionPlot[196].X = -2.38
	pointsOfFunctionPlot[196].Y = 0.021

	pointsOfFunctionPlot[197].X = -2.37
	pointsOfFunctionPlot[197].Y = 0.022

	pointsOfFunctionPlot[198].X = -2.36
	pointsOfFunctionPlot[198].Y = 0.022

	pointsOfFunctionPlot[199].X = -2.35
	pointsOfFunctionPlot[199].Y = 0.022

	pointsOfFunctionPlot[200].X = -2.34
	pointsOfFunctionPlot[200].Y = 0.023

	pointsOfFunctionPlot[201].X = -2.33
	pointsOfFunctionPlot[201].Y = 0.023

	pointsOfFunctionPlot[202].X = -2.32
	pointsOfFunctionPlot[202].Y = 0.023

	pointsOfFunctionPlot[203].X = -2.31
	pointsOfFunctionPlot[203].Y = 0.024

	pointsOfFunctionPlot[204].X = -2.3
	pointsOfFunctionPlot[204].Y = 0.024

	pointsOfFunctionPlot[205].X = -2.29
	pointsOfFunctionPlot[205].Y = 0.025

	pointsOfFunctionPlot[206].X = -2.28
	pointsOfFunctionPlot[206].Y = 0.025

	pointsOfFunctionPlot[207].X = -2.27
	pointsOfFunctionPlot[207].Y = 0.025

	pointsOfFunctionPlot[208].X = -2.26
	pointsOfFunctionPlot[208].Y = 0.026

	pointsOfFunctionPlot[209].X = -2.25
	pointsOfFunctionPlot[209].Y = 0.026

	pointsOfFunctionPlot[210].X = -2.24
	pointsOfFunctionPlot[210].Y = 0.027

	pointsOfFunctionPlot[211].X = -2.23
	pointsOfFunctionPlot[211].Y = 0.027

	pointsOfFunctionPlot[212].X = -2.22
	pointsOfFunctionPlot[212].Y = 0.028

	pointsOfFunctionPlot[213].X = -2.21
	pointsOfFunctionPlot[213].Y = 0.028

	pointsOfFunctionPlot[214].X = -2.2
	pointsOfFunctionPlot[214].Y = 0.029

	pointsOfFunctionPlot[215].X = -2.19
	pointsOfFunctionPlot[215].Y = 0.029

	pointsOfFunctionPlot[216].X = -2.18
	pointsOfFunctionPlot[216].Y = 0.029

	pointsOfFunctionPlot[217].X = -2.17
	pointsOfFunctionPlot[217].Y = 0.03

	pointsOfFunctionPlot[218].X = -2.16
	pointsOfFunctionPlot[218].Y = 0.03

	pointsOfFunctionPlot[219].X = -2.15
	pointsOfFunctionPlot[219].Y = 0.031

	pointsOfFunctionPlot[220].X = -2.14
	pointsOfFunctionPlot[220].Y = 0.031

	pointsOfFunctionPlot[221].X = -2.13
	pointsOfFunctionPlot[221].Y = 0.032

	pointsOfFunctionPlot[222].X = -2.12
	pointsOfFunctionPlot[222].Y = 0.033

	pointsOfFunctionPlot[223].X = -2.11
	pointsOfFunctionPlot[223].Y = 0.033

	pointsOfFunctionPlot[224].X = -2.1
	pointsOfFunctionPlot[224].Y = 0.034

	pointsOfFunctionPlot[225].X = -2.09
	pointsOfFunctionPlot[225].Y = 0.034

	pointsOfFunctionPlot[226].X = -2.08
	pointsOfFunctionPlot[226].Y = 0.035

	pointsOfFunctionPlot[227].X = -2.07
	pointsOfFunctionPlot[227].Y = 0.035

	pointsOfFunctionPlot[228].X = -2.06
	pointsOfFunctionPlot[228].Y = 0.036

	pointsOfFunctionPlot[229].X = -2.05
	pointsOfFunctionPlot[229].Y = 0.036

	pointsOfFunctionPlot[230].X = -2.04
	pointsOfFunctionPlot[230].Y = 0.037

	pointsOfFunctionPlot[231].X = -2.03
	pointsOfFunctionPlot[231].Y = 0.038

	pointsOfFunctionPlot[232].X = -2.02
	pointsOfFunctionPlot[232].Y = 0.038

	pointsOfFunctionPlot[233].X = -2.01
	pointsOfFunctionPlot[233].Y = 0.039

	pointsOfFunctionPlot[234].X = -2.0
	pointsOfFunctionPlot[234].Y = 0.04

	pointsOfFunctionPlot[235].X = -1.99
	pointsOfFunctionPlot[235].Y = 0.04

	pointsOfFunctionPlot[236].X = -1.98
	pointsOfFunctionPlot[236].Y = 0.041

	pointsOfFunctionPlot[237].X = -1.97
	pointsOfFunctionPlot[237].Y = 0.042

	pointsOfFunctionPlot[238].X = -1.96
	pointsOfFunctionPlot[238].Y = 0.042

	pointsOfFunctionPlot[239].X = -1.95
	pointsOfFunctionPlot[239].Y = 0.043

	pointsOfFunctionPlot[240].X = -1.94
	pointsOfFunctionPlot[240].Y = 0.044

	pointsOfFunctionPlot[241].X = -1.93
	pointsOfFunctionPlot[241].Y = 0.044

	pointsOfFunctionPlot[242].X = -1.92
	pointsOfFunctionPlot[242].Y = 0.045

	pointsOfFunctionPlot[243].X = -1.91
	pointsOfFunctionPlot[243].Y = 0.046

	pointsOfFunctionPlot[244].X = -1.9
	pointsOfFunctionPlot[244].Y = 0.047

	pointsOfFunctionPlot[245].X = -1.89
	pointsOfFunctionPlot[245].Y = 0.047

	pointsOfFunctionPlot[246].X = -1.88
	pointsOfFunctionPlot[246].Y = 0.048

	pointsOfFunctionPlot[247].X = -1.87
	pointsOfFunctionPlot[247].Y = 0.049

	pointsOfFunctionPlot[248].X = -1.86
	pointsOfFunctionPlot[248].Y = 0.05

	pointsOfFunctionPlot[249].X = -1.85
	pointsOfFunctionPlot[249].Y = 0.05

	pointsOfFunctionPlot[250].X = -1.84
	pointsOfFunctionPlot[250].Y = 0.051

	pointsOfFunctionPlot[251].X = -1.83
	pointsOfFunctionPlot[251].Y = 0.052

	pointsOfFunctionPlot[252].X = -1.82
	pointsOfFunctionPlot[252].Y = 0.053

	pointsOfFunctionPlot[253].X = -1.81
	pointsOfFunctionPlot[253].Y = 0.054

	pointsOfFunctionPlot[254].X = -1.8
	pointsOfFunctionPlot[254].Y = 0.055

	pointsOfFunctionPlot[255].X = -1.79
	pointsOfFunctionPlot[255].Y = 0.056

	pointsOfFunctionPlot[256].X = -1.78
	pointsOfFunctionPlot[256].Y = 0.057

	pointsOfFunctionPlot[257].X = -1.77
	pointsOfFunctionPlot[257].Y = 0.057

	pointsOfFunctionPlot[258].X = -1.76
	pointsOfFunctionPlot[258].Y = 0.058

	pointsOfFunctionPlot[259].X = -1.75
	pointsOfFunctionPlot[259].Y = 0.059

	pointsOfFunctionPlot[260].X = -1.74
	pointsOfFunctionPlot[260].Y = 0.06

	pointsOfFunctionPlot[261].X = -1.73
	pointsOfFunctionPlot[261].Y = 0.061

	pointsOfFunctionPlot[262].X = -1.72
	pointsOfFunctionPlot[262].Y = 0.062

	pointsOfFunctionPlot[263].X = -1.71
	pointsOfFunctionPlot[263].Y = 0.063

	pointsOfFunctionPlot[264].X = -1.7
	pointsOfFunctionPlot[264].Y = 0.064

	pointsOfFunctionPlot[265].X = -1.69
	pointsOfFunctionPlot[265].Y = 0.065

	pointsOfFunctionPlot[266].X = -1.68
	pointsOfFunctionPlot[266].Y = 0.066

	pointsOfFunctionPlot[267].X = -1.67
	pointsOfFunctionPlot[267].Y = 0.068

	pointsOfFunctionPlot[268].X = -1.66
	pointsOfFunctionPlot[268].Y = 0.069

	pointsOfFunctionPlot[269].X = -1.65
	pointsOfFunctionPlot[269].Y = 0.07

	pointsOfFunctionPlot[270].X = -1.64
	pointsOfFunctionPlot[270].Y = 0.071

	pointsOfFunctionPlot[271].X = -1.63
	pointsOfFunctionPlot[271].Y = 0.072

	pointsOfFunctionPlot[272].X = -1.62
	pointsOfFunctionPlot[272].Y = 0.073

	pointsOfFunctionPlot[273].X = -1.61
	pointsOfFunctionPlot[273].Y = 0.074

	pointsOfFunctionPlot[274].X = -1.6
	pointsOfFunctionPlot[274].Y = 0.076

	pointsOfFunctionPlot[275].X = -1.59
	pointsOfFunctionPlot[275].Y = 0.077

	pointsOfFunctionPlot[276].X = -1.58
	pointsOfFunctionPlot[276].Y = 0.078

	pointsOfFunctionPlot[277].X = -1.57
	pointsOfFunctionPlot[277].Y = 0.079

	pointsOfFunctionPlot[278].X = -1.56
	pointsOfFunctionPlot[278].Y = 0.081

	pointsOfFunctionPlot[279].X = -1.55
	pointsOfFunctionPlot[279].Y = 0.082

	pointsOfFunctionPlot[280].X = -1.54
	pointsOfFunctionPlot[280].Y = 0.083

	pointsOfFunctionPlot[281].X = -1.53
	pointsOfFunctionPlot[281].Y = 0.085

	pointsOfFunctionPlot[282].X = -1.52
	pointsOfFunctionPlot[282].Y = 0.086

	pointsOfFunctionPlot[283].X = -1.51
	pointsOfFunctionPlot[283].Y = 0.088

	pointsOfFunctionPlot[284].X = -1.5
	pointsOfFunctionPlot[284].Y = 0.089

	pointsOfFunctionPlot[285].X = -1.49
	pointsOfFunctionPlot[285].Y = 0.09

	pointsOfFunctionPlot[286].X = -1.48
	pointsOfFunctionPlot[286].Y = 0.092

	pointsOfFunctionPlot[287].X = -1.47
	pointsOfFunctionPlot[287].Y = 0.093

	pointsOfFunctionPlot[288].X = -1.46
	pointsOfFunctionPlot[288].Y = 0.095

	pointsOfFunctionPlot[289].X = -1.45
	pointsOfFunctionPlot[289].Y = 0.096

	pointsOfFunctionPlot[290].X = -1.44
	pointsOfFunctionPlot[290].Y = 0.098

	pointsOfFunctionPlot[291].X = -1.43
	pointsOfFunctionPlot[291].Y = 0.1

	pointsOfFunctionPlot[292].X = -1.42
	pointsOfFunctionPlot[292].Y = 0.101

	pointsOfFunctionPlot[293].X = -1.41
	pointsOfFunctionPlot[293].Y = 0.103

	pointsOfFunctionPlot[294].X = -1.4
	pointsOfFunctionPlot[294].Y = 0.105

	pointsOfFunctionPlot[295].X = -1.39
	pointsOfFunctionPlot[295].Y = 0.106

	pointsOfFunctionPlot[296].X = -1.38
	pointsOfFunctionPlot[296].Y = 0.108

	pointsOfFunctionPlot[297].X = -1.37
	pointsOfFunctionPlot[297].Y = 0.11

	pointsOfFunctionPlot[298].X = -1.36
	pointsOfFunctionPlot[298].Y = 0.112

	pointsOfFunctionPlot[299].X = -1.35
	pointsOfFunctionPlot[299].Y = 0.113

	pointsOfFunctionPlot[300].X = -1.34
	pointsOfFunctionPlot[300].Y = 0.115

	pointsOfFunctionPlot[301].X = -1.33
	pointsOfFunctionPlot[301].Y = 0.117

	pointsOfFunctionPlot[302].X = -1.32
	pointsOfFunctionPlot[302].Y = 0.119

	pointsOfFunctionPlot[303].X = -1.31
	pointsOfFunctionPlot[303].Y = 0.121

	pointsOfFunctionPlot[304].X = -1.3
	pointsOfFunctionPlot[304].Y = 0.123

	pointsOfFunctionPlot[305].X = -1.29
	pointsOfFunctionPlot[305].Y = 0.125

	pointsOfFunctionPlot[306].X = -1.28
	pointsOfFunctionPlot[306].Y = 0.127

	pointsOfFunctionPlot[307].X = -1.27
	pointsOfFunctionPlot[307].Y = 0.129

	pointsOfFunctionPlot[308].X = -1.26
	pointsOfFunctionPlot[308].Y = 0.131

	pointsOfFunctionPlot[309].X = -1.25
	pointsOfFunctionPlot[309].Y = 0.133

	pointsOfFunctionPlot[310].X = -1.24
	pointsOfFunctionPlot[310].Y = 0.135

	pointsOfFunctionPlot[311].X = -1.23
	pointsOfFunctionPlot[311].Y = 0.138

	pointsOfFunctionPlot[312].X = -1.22
	pointsOfFunctionPlot[312].Y = 0.14

	pointsOfFunctionPlot[313].X = -1.21
	pointsOfFunctionPlot[313].Y = 0.142

	pointsOfFunctionPlot[314].X = -1.2
	pointsOfFunctionPlot[314].Y = 0.145

	pointsOfFunctionPlot[315].X = -1.19
	pointsOfFunctionPlot[315].Y = 0.147

	pointsOfFunctionPlot[316].X = -1.18
	pointsOfFunctionPlot[316].Y = 0.149

	pointsOfFunctionPlot[317].X = -1.17
	pointsOfFunctionPlot[317].Y = 0.152

	pointsOfFunctionPlot[318].X = -1.16
	pointsOfFunctionPlot[318].Y = 0.154

	pointsOfFunctionPlot[319].X = -1.15
	pointsOfFunctionPlot[319].Y = 0.157

	pointsOfFunctionPlot[320].X = -1.14
	pointsOfFunctionPlot[320].Y = 0.159

	pointsOfFunctionPlot[321].X = -1.13
	pointsOfFunctionPlot[321].Y = 0.162

	pointsOfFunctionPlot[322].X = -1.12
	pointsOfFunctionPlot[322].Y = 0.164

	pointsOfFunctionPlot[323].X = -1.11
	pointsOfFunctionPlot[323].Y = 0.167

	pointsOfFunctionPlot[324].X = -1.1
	pointsOfFunctionPlot[324].Y = 0.17

	pointsOfFunctionPlot[325].X = -1.09
	pointsOfFunctionPlot[325].Y = 0.173

	pointsOfFunctionPlot[326].X = -1.08
	pointsOfFunctionPlot[326].Y = 0.175

	pointsOfFunctionPlot[327].X = -1.07
	pointsOfFunctionPlot[327].Y = 0.178

	pointsOfFunctionPlot[328].X = -1.06
	pointsOfFunctionPlot[328].Y = 0.181

	pointsOfFunctionPlot[329].X = -1.05
	pointsOfFunctionPlot[329].Y = 0.184

	pointsOfFunctionPlot[330].X = -1.04
	pointsOfFunctionPlot[330].Y = 0.187

	pointsOfFunctionPlot[331].X = -1.03
	pointsOfFunctionPlot[331].Y = 0.19

	pointsOfFunctionPlot[332].X = -1.02
	pointsOfFunctionPlot[332].Y = 0.193

	pointsOfFunctionPlot[333].X = -1.01
	pointsOfFunctionPlot[333].Y = 0.196

	pointsOfFunctionPlot[334].X = -1.0
	pointsOfFunctionPlot[334].Y = 0.2

	pointsOfFunctionPlot[335].X = -0.99
	pointsOfFunctionPlot[335].Y = 0.203

	pointsOfFunctionPlot[336].X = -0.98
	pointsOfFunctionPlot[336].Y = 0.206

	pointsOfFunctionPlot[337].X = -0.97
	pointsOfFunctionPlot[337].Y = 0.209

	pointsOfFunctionPlot[338].X = -0.96
	pointsOfFunctionPlot[338].Y = 0.213

	pointsOfFunctionPlot[339].X = -0.95
	pointsOfFunctionPlot[339].Y = 0.216

	pointsOfFunctionPlot[340].X = -0.94
	pointsOfFunctionPlot[340].Y = 0.22

	pointsOfFunctionPlot[341].X = -0.93
	pointsOfFunctionPlot[341].Y = 0.223

	pointsOfFunctionPlot[342].X = -0.92
	pointsOfFunctionPlot[342].Y = 0.227

	pointsOfFunctionPlot[343].X = -0.91
	pointsOfFunctionPlot[343].Y = 0.231

	pointsOfFunctionPlot[344].X = -0.9
	pointsOfFunctionPlot[344].Y = 0.234

	pointsOfFunctionPlot[345].X = -0.89
	pointsOfFunctionPlot[345].Y = 0.238

	pointsOfFunctionPlot[346].X = -0.88
	pointsOfFunctionPlot[346].Y = 0.242

	pointsOfFunctionPlot[347].X = -0.87
	pointsOfFunctionPlot[347].Y = 0.246

	pointsOfFunctionPlot[348].X = -0.86
	pointsOfFunctionPlot[348].Y = 0.25

	pointsOfFunctionPlot[349].X = -0.85
	pointsOfFunctionPlot[349].Y = 0.254

	pointsOfFunctionPlot[350].X = -0.84
	pointsOfFunctionPlot[350].Y = 0.258

	pointsOfFunctionPlot[351].X = -0.83
	pointsOfFunctionPlot[351].Y = 0.262

	pointsOfFunctionPlot[352].X = -0.82
	pointsOfFunctionPlot[352].Y = 0.267

	pointsOfFunctionPlot[353].X = -0.81
	pointsOfFunctionPlot[353].Y = 0.271

	pointsOfFunctionPlot[354].X = -0.8
	pointsOfFunctionPlot[354].Y = 0.275

	pointsOfFunctionPlot[355].X = -0.79
	pointsOfFunctionPlot[355].Y = 0.28

	pointsOfFunctionPlot[356].X = -0.78
	pointsOfFunctionPlot[356].Y = 0.285

	pointsOfFunctionPlot[357].X = -0.77
	pointsOfFunctionPlot[357].Y = 0.289

	pointsOfFunctionPlot[358].X = -0.76
	pointsOfFunctionPlot[358].Y = 0.294

	pointsOfFunctionPlot[359].X = -0.75
	pointsOfFunctionPlot[359].Y = 0.299

	pointsOfFunctionPlot[360].X = -0.74
	pointsOfFunctionPlot[360].Y = 0.303

	pointsOfFunctionPlot[361].X = -0.73
	pointsOfFunctionPlot[361].Y = 0.308

	pointsOfFunctionPlot[362].X = -0.72
	pointsOfFunctionPlot[362].Y = 0.313

	pointsOfFunctionPlot[363].X = -0.71
	pointsOfFunctionPlot[363].Y = 0.319

	pointsOfFunctionPlot[364].X = -0.7
	pointsOfFunctionPlot[364].Y = 0.324

	pointsOfFunctionPlot[365].X = -0.69
	pointsOfFunctionPlot[365].Y = 0.329

	pointsOfFunctionPlot[366].X = -0.68
	pointsOfFunctionPlot[366].Y = 0.334

	pointsOfFunctionPlot[367].X = -0.67
	pointsOfFunctionPlot[367].Y = 0.34

	pointsOfFunctionPlot[368].X = -0.66
	pointsOfFunctionPlot[368].Y = 0.345

	pointsOfFunctionPlot[369].X = -0.65
	pointsOfFunctionPlot[369].Y = 0.351

	pointsOfFunctionPlot[370].X = -0.64
	pointsOfFunctionPlot[370].Y = 0.357

	pointsOfFunctionPlot[371].X = -0.63
	pointsOfFunctionPlot[371].Y = 0.362

	pointsOfFunctionPlot[372].X = -0.62
	pointsOfFunctionPlot[372].Y = 0.368

	pointsOfFunctionPlot[373].X = -0.61
	pointsOfFunctionPlot[373].Y = 0.374

	pointsOfFunctionPlot[374].X = -0.6
	pointsOfFunctionPlot[374].Y = 0.38

	pointsOfFunctionPlot[375].X = -0.59
	pointsOfFunctionPlot[375].Y = 0.386

	pointsOfFunctionPlot[376].X = -0.58
	pointsOfFunctionPlot[376].Y = 0.393

	pointsOfFunctionPlot[377].X = -0.57
	pointsOfFunctionPlot[377].Y = 0.399

	pointsOfFunctionPlot[378].X = -0.56
	pointsOfFunctionPlot[378].Y = 0.406

	pointsOfFunctionPlot[379].X = -0.55
	pointsOfFunctionPlot[379].Y = 0.412

	pointsOfFunctionPlot[380].X = -0.54
	pointsOfFunctionPlot[380].Y = 0.419

	pointsOfFunctionPlot[381].X = -0.53
	pointsOfFunctionPlot[381].Y = 0.426

	pointsOfFunctionPlot[382].X = -0.52
	pointsOfFunctionPlot[382].Y = 0.433

	pointsOfFunctionPlot[383].X = -0.51
	pointsOfFunctionPlot[383].Y = 0.44

	pointsOfFunctionPlot[384].X = -0.5
	pointsOfFunctionPlot[384].Y = 0.447

	pointsOfFunctionPlot[385].X = -0.49
	pointsOfFunctionPlot[385].Y = 0.454

	pointsOfFunctionPlot[386].X = -0.48
	pointsOfFunctionPlot[386].Y = 0.461

	pointsOfFunctionPlot[387].X = -0.47
	pointsOfFunctionPlot[387].Y = 0.469

	pointsOfFunctionPlot[388].X = -0.46
	pointsOfFunctionPlot[388].Y = 0.477

	pointsOfFunctionPlot[389].X = -0.45
	pointsOfFunctionPlot[389].Y = 0.484

	pointsOfFunctionPlot[390].X = -0.44
	pointsOfFunctionPlot[390].Y = 0.492

	pointsOfFunctionPlot[391].X = -0.43
	pointsOfFunctionPlot[391].Y = 0.5

	pointsOfFunctionPlot[392].X = -0.42
	pointsOfFunctionPlot[392].Y = 0.508

	pointsOfFunctionPlot[393].X = -0.41
	pointsOfFunctionPlot[393].Y = 0.516

	pointsOfFunctionPlot[394].X = -0.4
	pointsOfFunctionPlot[394].Y = 0.525

	pointsOfFunctionPlot[395].X = -0.39
	pointsOfFunctionPlot[395].Y = 0.533

	pointsOfFunctionPlot[396].X = -0.38
	pointsOfFunctionPlot[396].Y = 0.542

	pointsOfFunctionPlot[397].X = -0.37
	pointsOfFunctionPlot[397].Y = 0.551

	pointsOfFunctionPlot[398].X = -0.36
	pointsOfFunctionPlot[398].Y = 0.56

	pointsOfFunctionPlot[399].X = -0.35
	pointsOfFunctionPlot[399].Y = 0.569

	pointsOfFunctionPlot[400].X = -0.34
	pointsOfFunctionPlot[400].Y = 0.578

	pointsOfFunctionPlot[401].X = -0.33
	pointsOfFunctionPlot[401].Y = 0.587

	pointsOfFunctionPlot[402].X = -0.32
	pointsOfFunctionPlot[402].Y = 0.597

	pointsOfFunctionPlot[403].X = -0.31
	pointsOfFunctionPlot[403].Y = 0.607

	pointsOfFunctionPlot[404].X = -0.3
	pointsOfFunctionPlot[404].Y = 0.617

	pointsOfFunctionPlot[405].X = -0.29
	pointsOfFunctionPlot[405].Y = 0.627

	pointsOfFunctionPlot[406].X = -0.28
	pointsOfFunctionPlot[406].Y = 0.637

	pointsOfFunctionPlot[407].X = -0.27
	pointsOfFunctionPlot[407].Y = 0.647

	pointsOfFunctionPlot[408].X = -0.26
	pointsOfFunctionPlot[408].Y = 0.658

	pointsOfFunctionPlot[409].X = -0.25
	pointsOfFunctionPlot[409].Y = 0.668

	pointsOfFunctionPlot[410].X = -0.24
	pointsOfFunctionPlot[410].Y = 0.679

	pointsOfFunctionPlot[411].X = -0.23
	pointsOfFunctionPlot[411].Y = 0.69

	pointsOfFunctionPlot[412].X = -0.22
	pointsOfFunctionPlot[412].Y = 0.701

	pointsOfFunctionPlot[413].X = -0.21
	pointsOfFunctionPlot[413].Y = 0.713

	pointsOfFunctionPlot[414].X = -0.2
	pointsOfFunctionPlot[414].Y = 0.724

	pointsOfFunctionPlot[415].X = -0.19
	pointsOfFunctionPlot[415].Y = 0.736

	pointsOfFunctionPlot[416].X = -0.18
	pointsOfFunctionPlot[416].Y = 0.748

	pointsOfFunctionPlot[417].X = -0.17
	pointsOfFunctionPlot[417].Y = 0.76

	pointsOfFunctionPlot[418].X = -0.16
	pointsOfFunctionPlot[418].Y = 0.773

	pointsOfFunctionPlot[419].X = -0.15
	pointsOfFunctionPlot[419].Y = 0.785

	pointsOfFunctionPlot[420].X = -0.14
	pointsOfFunctionPlot[420].Y = 0.798

	pointsOfFunctionPlot[421].X = -0.13
	pointsOfFunctionPlot[421].Y = 0.811

	pointsOfFunctionPlot[422].X = -0.12
	pointsOfFunctionPlot[422].Y = 0.824

	pointsOfFunctionPlot[423].X = -0.11
	pointsOfFunctionPlot[423].Y = 0.837

	pointsOfFunctionPlot[424].X = -0.1
	pointsOfFunctionPlot[424].Y = 0.851

	pointsOfFunctionPlot[425].X = -0.09
	pointsOfFunctionPlot[425].Y = 0.865

	pointsOfFunctionPlot[426].X = -0.08
	pointsOfFunctionPlot[426].Y = 0.879

	pointsOfFunctionPlot[427].X = -0.07
	pointsOfFunctionPlot[427].Y = 0.893

	pointsOfFunctionPlot[428].X = -0.06
	pointsOfFunctionPlot[428].Y = 0.907

	pointsOfFunctionPlot[429].X = -0.05
	pointsOfFunctionPlot[429].Y = 0.922

	pointsOfFunctionPlot[430].X = -0.04
	pointsOfFunctionPlot[430].Y = 0.937

	pointsOfFunctionPlot[431].X = -0.03
	pointsOfFunctionPlot[431].Y = 0.952

	pointsOfFunctionPlot[432].X = -0.02
	pointsOfFunctionPlot[432].Y = 0.968

	pointsOfFunctionPlot[433].X = -0.01
	pointsOfFunctionPlot[433].Y = 0.984

	pointsOfFunctionPlot[434].X = 0.0
	pointsOfFunctionPlot[434].Y = 1.0

	pointsOfFunctionPlot[435].X = 0.01
	pointsOfFunctionPlot[435].Y = 1.016

	pointsOfFunctionPlot[436].X = 0.02
	pointsOfFunctionPlot[436].Y = 1.032

	pointsOfFunctionPlot[437].X = 0.03
	pointsOfFunctionPlot[437].Y = 1.049

	pointsOfFunctionPlot[438].X = 0.04
	pointsOfFunctionPlot[438].Y = 1.066

	pointsOfFunctionPlot[439].X = 0.05
	pointsOfFunctionPlot[439].Y = 1.083

	pointsOfFunctionPlot[440].X = 0.06
	pointsOfFunctionPlot[440].Y = 1.101

	pointsOfFunctionPlot[441].X = 0.07
	pointsOfFunctionPlot[441].Y = 1.119

	pointsOfFunctionPlot[442].X = 0.08
	pointsOfFunctionPlot[442].Y = 1.137

	pointsOfFunctionPlot[443].X = 0.09
	pointsOfFunctionPlot[443].Y = 1.155

	pointsOfFunctionPlot[444].X = 0.1
	pointsOfFunctionPlot[444].Y = 1.174

	pointsOfFunctionPlot[445].X = 0.11
	pointsOfFunctionPlot[445].Y = 1.193

	pointsOfFunctionPlot[446].X = 0.12
	pointsOfFunctionPlot[446].Y = 1.213

	pointsOfFunctionPlot[447].X = 0.13
	pointsOfFunctionPlot[447].Y = 1.232

	pointsOfFunctionPlot[448].X = 0.14
	pointsOfFunctionPlot[448].Y = 1.252

	pointsOfFunctionPlot[449].X = 0.15
	pointsOfFunctionPlot[449].Y = 1.273

	pointsOfFunctionPlot[450].X = 0.16
	pointsOfFunctionPlot[450].Y = 1.293

	pointsOfFunctionPlot[451].X = 0.17
	pointsOfFunctionPlot[451].Y = 1.314

	pointsOfFunctionPlot[452].X = 0.18
	pointsOfFunctionPlot[452].Y = 1.336

	pointsOfFunctionPlot[453].X = 0.19
	pointsOfFunctionPlot[453].Y = 1.357

	pointsOfFunctionPlot[454].X = 0.2
	pointsOfFunctionPlot[454].Y = 1.379

	pointsOfFunctionPlot[455].X = 0.21
	pointsOfFunctionPlot[455].Y = 1.402

	pointsOfFunctionPlot[456].X = 0.22
	pointsOfFunctionPlot[456].Y = 1.424

	pointsOfFunctionPlot[457].X = 0.23
	pointsOfFunctionPlot[457].Y = 1.448

	pointsOfFunctionPlot[458].X = 0.24
	pointsOfFunctionPlot[458].Y = 1.471

	pointsOfFunctionPlot[459].X = 0.25
	pointsOfFunctionPlot[459].Y = 1.495

	pointsOfFunctionPlot[460].X = 0.26
	pointsOfFunctionPlot[460].Y = 1.519

	pointsOfFunctionPlot[461].X = 0.27
	pointsOfFunctionPlot[461].Y = 1.544

	pointsOfFunctionPlot[462].X = 0.28
	pointsOfFunctionPlot[462].Y = 1.569

	pointsOfFunctionPlot[463].X = 0.29
	pointsOfFunctionPlot[463].Y = 1.594

	pointsOfFunctionPlot[464].X = 0.3
	pointsOfFunctionPlot[464].Y = 1.62

	pointsOfFunctionPlot[465].X = 0.31
	pointsOfFunctionPlot[465].Y = 1.647

	pointsOfFunctionPlot[466].X = 0.32
	pointsOfFunctionPlot[466].Y = 1.673

	pointsOfFunctionPlot[467].X = 0.33
	pointsOfFunctionPlot[467].Y = 1.7

	pointsOfFunctionPlot[468].X = 0.34
	pointsOfFunctionPlot[468].Y = 1.728

	pointsOfFunctionPlot[469].X = 0.35
	pointsOfFunctionPlot[469].Y = 1.756

	pointsOfFunctionPlot[470].X = 0.36
	pointsOfFunctionPlot[470].Y = 1.785

	pointsOfFunctionPlot[471].X = 0.37
	pointsOfFunctionPlot[471].Y = 1.813

	pointsOfFunctionPlot[472].X = 0.38
	pointsOfFunctionPlot[472].Y = 1.843

	pointsOfFunctionPlot[473].X = 0.39
	pointsOfFunctionPlot[473].Y = 1.873

	pointsOfFunctionPlot[474].X = 0.40
	pointsOfFunctionPlot[474].Y = 1.903

	pointsOfFunctionPlot[475].X = 0.41
	pointsOfFunctionPlot[475].Y = 1.934

	pointsOfFunctionPlot[476].X = 0.42
	pointsOfFunctionPlot[476].Y = 1.965

	pointsOfFunctionPlot[477].X = 0.43
	pointsOfFunctionPlot[477].Y = 1.997

	pointsOfFunctionPlot[478].X = 0.44
	pointsOfFunctionPlot[478].Y = 2.03

	pointsOfFunctionPlot[479].X = 0.45
	pointsOfFunctionPlot[479].Y = 2.063

	pointsOfFunctionPlot[480].X = 0.46
	pointsOfFunctionPlot[480].Y = 2.096

	pointsOfFunctionPlot[481].X = 0.47
	pointsOfFunctionPlot[481].Y = 2.13

	pointsOfFunctionPlot[482].X = 0.48
	pointsOfFunctionPlot[482].Y = 2.165

	pointsOfFunctionPlot[483].X = 0.49
	pointsOfFunctionPlot[483].Y = 2.2

	pointsOfFunctionPlot[484].X = 0.5
	pointsOfFunctionPlot[484].Y = 2.236

	pointsOfFunctionPlot[485].X = 0.51
	pointsOfFunctionPlot[485].Y = 2.272

	pointsOfFunctionPlot[486].X = 0.52
	pointsOfFunctionPlot[486].Y = 2.309

	pointsOfFunctionPlot[487].X = 0.53
	pointsOfFunctionPlot[487].Y = 2.346

	pointsOfFunctionPlot[488].X = 0.54
	pointsOfFunctionPlot[488].Y = 2.384

	pointsOfFunctionPlot[489].X = 0.55
	pointsOfFunctionPlot[489].Y = 2.423

	pointsOfFunctionPlot[490].X = 0.56
	pointsOfFunctionPlot[490].Y = 2.462

	pointsOfFunctionPlot[491].X = 0.57
	pointsOfFunctionPlot[491].Y = 2.502

	pointsOfFunctionPlot[492].X = 0.58
	pointsOfFunctionPlot[492].Y = 2.543

	pointsOfFunctionPlot[493].X = 0.59
	pointsOfFunctionPlot[493].Y = 2.584

	pointsOfFunctionPlot[494].X = 0.6
	pointsOfFunctionPlot[494].Y = 2.626

	pointsOfFunctionPlot[495].X = 0.61
	pointsOfFunctionPlot[495].Y = 2.669

	pointsOfFunctionPlot[496].X = 0.62
	pointsOfFunctionPlot[496].Y = 2.712

	pointsOfFunctionPlot[497].X = 0.63
	pointsOfFunctionPlot[497].Y = 2.756

	pointsOfFunctionPlot[498].X = 0.64
	pointsOfFunctionPlot[498].Y = 2.801

	pointsOfFunctionPlot[499].X = 0.65
	pointsOfFunctionPlot[499].Y = 2.846

	pointsOfFunctionPlot[500].X = 0.66
	pointsOfFunctionPlot[500].Y = 2.892

	pointsOfFunctionPlot[501].X = 0.67
	pointsOfFunctionPlot[501].Y = 2.939

	pointsOfFunctionPlot[502].X = 0.68
	pointsOfFunctionPlot[502].Y = 2.987

	pointsOfFunctionPlot[503].X = 0.69
	pointsOfFunctionPlot[503].Y = 3.035

	pointsOfFunctionPlot[504].X = 0.7
	pointsOfFunctionPlot[504].Y = 3.085

	pointsOfFunctionPlot[505].X = 0.71
	pointsOfFunctionPlot[505].Y = 3.135

	pointsOfFunctionPlot[506].X = 0.72
	pointsOfFunctionPlot[506].Y = 3.186

	pointsOfFunctionPlot[507].X = 0.73
	pointsOfFunctionPlot[507].Y = 3.237

	pointsOfFunctionPlot[508].X = 0.74
	pointsOfFunctionPlot[508].Y = 3.29

	pointsOfFunctionPlot[509].X = 0.75
	pointsOfFunctionPlot[509].Y = 3.343

	pointsOfFunctionPlot[510].X = 0.76
	pointsOfFunctionPlot[510].Y = 3.398

	pointsOfFunctionPlot[511].X = 0.77
	pointsOfFunctionPlot[511].Y = 3.453

	pointsOfFunctionPlot[512].X = 0.78
	pointsOfFunctionPlot[512].Y = 3.509

	pointsOfFunctionPlot[513].X = 0.79
	pointsOfFunctionPlot[513].Y = 3.566

	pointsOfFunctionPlot[514].X = 0.8
	pointsOfFunctionPlot[514].Y = 3.623

	pointsOfFunctionPlot[515].X = 0.81
	pointsOfFunctionPlot[515].Y = 3.682

	pointsOfFunctionPlot[516].X = 0.82
	pointsOfFunctionPlot[516].Y = 3.742

	pointsOfFunctionPlot[517].X = 0.83
	pointsOfFunctionPlot[517].Y = 3.803

	pointsOfFunctionPlot[518].X = 0.84
	pointsOfFunctionPlot[518].Y = 3.864

	pointsOfFunctionPlot[519].X = 0.85
	pointsOfFunctionPlot[519].Y = 3.927

	pointsOfFunctionPlot[520].X = 0.86
	pointsOfFunctionPlot[520].Y = 3.991

	pointsOfFunctionPlot[521].X = 0.87
	pointsOfFunctionPlot[521].Y = 4.056

	pointsOfFunctionPlot[522].X = 0.88
	pointsOfFunctionPlot[522].Y = 4.121

	pointsOfFunctionPlot[523].X = 0.89
	pointsOfFunctionPlot[523].Y = 4.188

	pointsOfFunctionPlot[524].X = 0.9
	pointsOfFunctionPlot[524].Y = 4.256

	pointsOfFunctionPlot[525].X = 0.91
	pointsOfFunctionPlot[525].Y = 4.325

	pointsOfFunctionPlot[526].X = 0.92
	pointsOfFunctionPlot[526].Y = 4.395

	pointsOfFunctionPlot[527].X = 0.93
	pointsOfFunctionPlot[527].Y = 4.467

	pointsOfFunctionPlot[528].X = 0.94
	pointsOfFunctionPlot[528].Y = 4.539

	pointsOfFunctionPlot[529].X = 0.95
	pointsOfFunctionPlot[529].Y = 4.613

	pointsOfFunctionPlot[530].X = 0.96
	pointsOfFunctionPlot[530].Y = 4.688

	pointsOfFunctionPlot[531].X = 0.97
	pointsOfFunctionPlot[531].Y = 4.764

	pointsOfFunctionPlot[532].X = 0.98
	pointsOfFunctionPlot[532].Y = 4.841

	pointsOfFunctionPlot[533].X = 0.99
	pointsOfFunctionPlot[533].Y = 4.92

	pointsOfFunctionPlot[534].X = 1.0
	pointsOfFunctionPlot[534].Y = 5.0

	pointsOfFunctionPlot[535].X = 1.01
	pointsOfFunctionPlot[535].Y = 5.081

	pointsOfFunctionPlot[536].X = 1.02
	pointsOfFunctionPlot[536].Y = 5.163

	pointsOfFunctionPlot[537].X = 1.03
	pointsOfFunctionPlot[537].Y = 5.247

	pointsOfFunctionPlot[538].X = 1.04
	pointsOfFunctionPlot[538].Y = 5.332

	pointsOfFunctionPlot[539].X = 1.05
	pointsOfFunctionPlot[539].Y = 5.419

	pointsOfFunctionPlot[540].X = 1.06
	pointsOfFunctionPlot[540].Y = 5.506

	pointsOfFunctionPlot[541].X = 1.07
	pointsOfFunctionPlot[541].Y = 5.596

	pointsOfFunctionPlot[542].X = 1.08
	pointsOfFunctionPlot[542].Y = 5.687

	pointsOfFunctionPlot[543].X = 1.09
	pointsOfFunctionPlot[543].Y = 5.779

	pointsOfFunctionPlot[544].X = 1.1
	pointsOfFunctionPlot[544].Y = 5.873

	pointsOfFunctionPlot[545].X = 1.11
	pointsOfFunctionPlot[545].Y = 5.968

	pointsOfFunctionPlot[546].X = 1.12
	pointsOfFunctionPlot[546].Y = 6.065

	pointsOfFunctionPlot[547].X = 1.13
	pointsOfFunctionPlot[547].Y = 6.163

	pointsOfFunctionPlot[548].X = 1.14
	pointsOfFunctionPlot[548].Y = 6.263

	pointsOfFunctionPlot[549].X = 1.15
	pointsOfFunctionPlot[549].Y = 6.365

	pointsOfFunctionPlot[550].X = 1.16
	pointsOfFunctionPlot[550].Y = 6.468

	pointsOfFunctionPlot[551].X = 1.17
	pointsOfFunctionPlot[551].Y = 6.573

	pointsOfFunctionPlot[552].X = 1.18
	pointsOfFunctionPlot[552].Y = 6.68

	pointsOfFunctionPlot[553].X = 1.19
	pointsOfFunctionPlot[553].Y = 6.788

	pointsOfFunctionPlot[554].X = 1.2
	pointsOfFunctionPlot[554].Y = 6.898

	pointsOfFunctionPlot[555].X = 1.21
	pointsOfFunctionPlot[555].Y = 7.01

	pointsOfFunctionPlot[556].X = 1.22
	pointsOfFunctionPlot[556].Y = 7.124

	pointsOfFunctionPlot[557].X = 1.23
	pointsOfFunctionPlot[557].Y = 7.239

	pointsOfFunctionPlot[558].X = 1.24
	pointsOfFunctionPlot[558].Y = 7.357

	pointsOfFunctionPlot[559].X = 1.25
	pointsOfFunctionPlot[559].Y = 7.476

	pointsOfFunctionPlot[560].X = 1.26
	pointsOfFunctionPlot[560].Y = 7.598

	pointsOfFunctionPlot[561].X = 1.27
	pointsOfFunctionPlot[561].Y = 7.721

	pointsOfFunctionPlot[562].X = 1.28
	pointsOfFunctionPlot[562].Y = 7.846

	pointsOfFunctionPlot[563].X = 1.29
	pointsOfFunctionPlot[563].Y = 7.973

	pointsOfFunctionPlot[564].X = 1.3
	pointsOfFunctionPlot[564].Y = 8.103

	pointsOfFunctionPlot[565].X = 1.31
	pointsOfFunctionPlot[565].Y = 8.234

	pointsOfFunctionPlot[566].X = 1.32
	pointsOfFunctionPlot[566].Y = 8.368

	pointsOfFunctionPlot[567].X = 1.33
	pointsOfFunctionPlot[567].Y = 8.504

	pointsOfFunctionPlot[568].X = 1.34
	pointsOfFunctionPlot[568].Y = 8.642

	pointsOfFunctionPlot[569].X = 1.35
	pointsOfFunctionPlot[569].Y = 8.782

	pointsOfFunctionPlot[570].X = 1.36
	pointsOfFunctionPlot[570].Y = 8.924

	pointsOfFunctionPlot[571].X = 1.37
	pointsOfFunctionPlot[571].Y = 9.069

	pointsOfFunctionPlot[572].X = 1.38
	pointsOfFunctionPlot[572].Y = 9.216

	pointsOfFunctionPlot[573].X = 1.39
	pointsOfFunctionPlot[573].Y = 9.366

	pointsOfFunctionPlot[574].X = 1.40
	pointsOfFunctionPlot[574].Y = 9.518

	pointsOfFunctionPlot[575].X = 1.41
	pointsOfFunctionPlot[575].Y = 9.672

	pointsOfFunctionPlot[576].X = 1.42
	pointsOfFunctionPlot[576].Y = 9.829

	pointsOfFunctionPlot[577].X = 1.43
	pointsOfFunctionPlot[577].Y = 9.989

	pointsOfFunctionPlot[578].X = 1.44
	pointsOfFunctionPlot[578].Y = 10.151

	pointsOfFunctionPlot[579].X = 1.45
	pointsOfFunctionPlot[579].Y = 10.315

	pointsOfFunctionPlot[580].X = 1.46
	pointsOfFunctionPlot[580].Y = 10.483

	pointsOfFunctionPlot[581].X = 1.47
	pointsOfFunctionPlot[581].Y = 10.653

	pointsOfFunctionPlot[582].X = 1.48
	pointsOfFunctionPlot[582].Y = 10.826

	pointsOfFunctionPlot[583].X = 1.49
	pointsOfFunctionPlot[583].Y = 11.001

	pointsOfFunctionPlot[584].X = 1.5
	pointsOfFunctionPlot[584].Y = 11.18

	pointsOfFunctionPlot[585].X = 1.51
	pointsOfFunctionPlot[585].Y = 11.361

	pointsOfFunctionPlot[586].X = 1.52
	pointsOfFunctionPlot[586].Y = 11.546

	pointsOfFunctionPlot[587].X = 1.53
	pointsOfFunctionPlot[587].Y = 11.733

	pointsOfFunctionPlot[588].X = 1.54
	pointsOfFunctionPlot[588].Y = 11.923

	pointsOfFunctionPlot[589].X = 1.55
	pointsOfFunctionPlot[589].Y = 12.117

	pointsOfFunctionPlot[590].X = 1.56
	pointsOfFunctionPlot[590].Y = 12.313

	pointsOfFunctionPlot[591].X = 1.57
	pointsOfFunctionPlot[591].Y = 12.513

	pointsOfFunctionPlot[592].X = 1.58
	pointsOfFunctionPlot[592].Y = 12.716

	pointsOfFunctionPlot[593].X = 1.59
	pointsOfFunctionPlot[593].Y = 12.923

	pointsOfFunctionPlot[594].X = 1.6
	pointsOfFunctionPlot[594].Y = 13.132

	pointsOfFunctionPlot[595].X = 1.61
	pointsOfFunctionPlot[595].Y = 13.345

	pointsOfFunctionPlot[596].X = 1.62
	pointsOfFunctionPlot[596].Y = 13.562

	pointsOfFunctionPlot[597].X = 1.63
	pointsOfFunctionPlot[597].Y = 13.782

	pointsOfFunctionPlot[598].X = 1.64
	pointsOfFunctionPlot[598].Y = 14.005

	pointsOfFunctionPlot[599].X = 1.65
	pointsOfFunctionPlot[599].Y = 14.233

	pointsOfFunctionPlot[600].X = 1.66
	pointsOfFunctionPlot[600].Y = 14.464

	pointsOfFunctionPlot[601].X = 1.67
	pointsOfFunctionPlot[601].Y = 14.698

	pointsOfFunctionPlot[602].X = 1.68
	pointsOfFunctionPlot[602].Y = 14.937

	pointsOfFunctionPlot[603].X = 1.69
	pointsOfFunctionPlot[603].Y = 15.179

	pointsOfFunctionPlot[604].X = 1.7
	pointsOfFunctionPlot[604].Y = 15.425

	pointsOfFunctionPlot[605].X = 1.71
	pointsOfFunctionPlot[605].Y = 15.676

	pointsOfFunctionPlot[606].X = 1.72
	pointsOfFunctionPlot[606].Y = 15.93

	pointsOfFunctionPlot[607].X = 1.73
	pointsOfFunctionPlot[607].Y = 16.188

	pointsOfFunctionPlot[608].X = 1.74
	pointsOfFunctionPlot[608].Y = 16.451

	pointsOfFunctionPlot[609].X = 1.75
	pointsOfFunctionPlot[609].Y = 16.718

	pointsOfFunctionPlot[610].X = 1.76
	pointsOfFunctionPlot[610].Y = 16.989

	pointsOfFunctionPlot[611].X = 1.77
	pointsOfFunctionPlot[611].Y = 17.265

	pointsOfFunctionPlot[612].X = 1.78
	pointsOfFunctionPlot[612].Y = 17.545

	pointsOfFunctionPlot[613].X = 1.79
	pointsOfFunctionPlot[613].Y = 17.83

	pointsOfFunctionPlot[614].X = 1.8
	pointsOfFunctionPlot[614].Y = 18.119

	pointsOfFunctionPlot[615].X = 1.81
	pointsOfFunctionPlot[615].Y = 18.413

	pointsOfFunctionPlot[616].X = 1.82
	pointsOfFunctionPlot[616].Y = 18.712

	pointsOfFunctionPlot[617].X = 1.83
	pointsOfFunctionPlot[617].Y = 19.015

	pointsOfFunctionPlot[618].X = 1.84
	pointsOfFunctionPlot[618].Y = 19.324

	pointsOfFunctionPlot[619].X = 1.85
	pointsOfFunctionPlot[619].Y = 19.637

	pointsOfFunctionPlot[620].X = 1.86
	pointsOfFunctionPlot[620].Y = 19.956

	pointsOfFunctionPlot[621].X = 1.87
	pointsOfFunctionPlot[621].Y = 20.28

	pointsOfFunctionPlot[622].X = 1.88
	pointsOfFunctionPlot[622].Y = 20.609

	pointsOfFunctionPlot[623].X = 1.89
	pointsOfFunctionPlot[623].Y = 20.943

	pointsOfFunctionPlot[624].X = 1.9
	pointsOfFunctionPlot[624].Y = 21.283

	pointsOfFunctionPlot[625].X = 1.91
	pointsOfFunctionPlot[625].Y = 21.628

	pointsOfFunctionPlot[626].X = 1.92
	pointsOfFunctionPlot[626].Y = 21.979

	pointsOfFunctionPlot[627].X = 1.93
	pointsOfFunctionPlot[627].Y = 22.336

	pointsOfFunctionPlot[628].X = 1.94
	pointsOfFunctionPlot[628].Y = 22.698

	pointsOfFunctionPlot[629].X = 1.95
	pointsOfFunctionPlot[629].Y = 23.067

	pointsOfFunctionPlot[630].X = 1.96
	pointsOfFunctionPlot[630].Y = 23.441

	pointsOfFunctionPlot[631].X = 1.97
	pointsOfFunctionPlot[631].Y = 23.821

	pointsOfFunctionPlot[632].X = 1.98
	pointsOfFunctionPlot[632].Y = 24.208

	pointsOfFunctionPlot[633].X = 1.99
	pointsOfFunctionPlot[633].Y = 24.6

	pointsOfFunctionPlot[634].X = 2.0
	pointsOfFunctionPlot[634].Y = 25.0

	pointsOfFunctionPlot[635].X = 2.01
	pointsOfFunctionPlot[635].Y = 25.405

	pointsOfFunctionPlot[636].X = 2.02
	pointsOfFunctionPlot[636].Y = 25.817

	pointsOfFunctionPlot[637].X = 2.03
	pointsOfFunctionPlot[637].Y = 26.236

	pointsOfFunctionPlot[638].X = 2.04
	pointsOfFunctionPlot[638].Y = 26.662

	pointsOfFunctionPlot[639].X = 2.05
	pointsOfFunctionPlot[639].Y = 27.095

	pointsOfFunctionPlot[640].X = 2.06
	pointsOfFunctionPlot[640].Y = 27.534

	pointsOfFunctionPlot[641].X = 2.07
	pointsOfFunctionPlot[641].Y = 27.981

	pointsOfFunctionPlot[642].X = 2.08
	pointsOfFunctionPlot[642].Y = 28.435

	pointsOfFunctionPlot[643].X = 2.09
	pointsOfFunctionPlot[643].Y = 28.896

	pointsOfFunctionPlot[644].X = 2.1
	pointsOfFunctionPlot[644].Y = 29.365

	pointsOfFunctionPlot[645].X = 2.11
	pointsOfFunctionPlot[645].Y = 29.841

	pointsOfFunctionPlot[646].X = 2.12
	pointsOfFunctionPlot[646].Y = 30.326

	pointsOfFunctionPlot[647].X = 2.13
	pointsOfFunctionPlot[647].Y = 30.818

	pointsOfFunctionPlot[648].X = 2.14
	pointsOfFunctionPlot[648].Y = 31.318

	pointsOfFunctionPlot[649].X = 2.15
	pointsOfFunctionPlot[649].Y = 31.826

	pointsOfFunctionPlot[650].X = 2.16
	pointsOfFunctionPlot[650].Y = 32.342

	pointsOfFunctionPlot[651].X = 2.17
	pointsOfFunctionPlot[651].Y = 32.867

	pointsOfFunctionPlot[652].X = 2.18
	pointsOfFunctionPlot[652].Y = 33.4

	pointsOfFunctionPlot[653].X = 2.19
	pointsOfFunctionPlot[653].Y = 33.942

	pointsOfFunctionPlot[654].X = 2.2
	pointsOfFunctionPlot[654].Y = 34.493

	pointsOfFunctionPlot[655].X = 2.21
	pointsOfFunctionPlot[655].Y = 35.052

	pointsOfFunctionPlot[656].X = 2.22
	pointsOfFunctionPlot[656].Y = 35.621

	pointsOfFunctionPlot[657].X = 2.23
	pointsOfFunctionPlot[657].Y = 36.199

	pointsOfFunctionPlot[658].X = 2.24
	pointsOfFunctionPlot[658].Y = 36.786

	pointsOfFunctionPlot[659].X = 2.25
	pointsOfFunctionPlot[659].Y = 37.383

	pointsOfFunctionPlot[660].X = 2.26
	pointsOfFunctionPlot[660].Y = 37.99

	pointsOfFunctionPlot[661].X = 2.27
	pointsOfFunctionPlot[661].Y = 38.606

	pointsOfFunctionPlot[662].X = 2.28
	pointsOfFunctionPlot[662].Y = 39.233

	pointsOfFunctionPlot[663].X = 2.29
	pointsOfFunctionPlot[663].Y = 39.869

	pointsOfFunctionPlot[664].X = 2.3
	pointsOfFunctionPlot[664].Y = 40.516

	pointsOfFunctionPlot[665].X = 2.31
	pointsOfFunctionPlot[665].Y = 41.173

	pointsOfFunctionPlot[666].X = 2.32
	pointsOfFunctionPlot[666].Y = 41.841

	pointsOfFunctionPlot[667].X = 2.33
	pointsOfFunctionPlot[667].Y = 42.52

	pointsOfFunctionPlot[668].X = 2.34
	pointsOfFunctionPlot[668].Y = 43.21

	pointsOfFunctionPlot[669].X = 2.35
	pointsOfFunctionPlot[669].Y = 43.911

	pointsOfFunctionPlot[670].X = 2.36
	pointsOfFunctionPlot[670].Y = 44.624

	pointsOfFunctionPlot[671].X = 2.37
	pointsOfFunctionPlot[671].Y = 45.348

	pointsOfFunctionPlot[672].X = 2.38
	pointsOfFunctionPlot[672].Y = 46.083

	pointsOfFunctionPlot[673].X = 2.39
	pointsOfFunctionPlot[673].Y = 46.831

	pointsOfFunctionPlot[674].X = 2.4
	pointsOfFunctionPlot[674].Y = 47.591

	pointsOfFunctionPlot[675].X = 2.41
	pointsOfFunctionPlot[675].Y = 48.363

	pointsOfFunctionPlot[676].X = 2.42
	pointsOfFunctionPlot[676].Y = 49.148

	pointsOfFunctionPlot[677].X = 2.43
	pointsOfFunctionPlot[677].Y = 49.945

	pointsOfFunctionPlot[678].X = 2.44
	pointsOfFunctionPlot[678].Y = 50.755

	pointsOfFunctionPlot[679].X = 2.45
	pointsOfFunctionPlot[679].Y = 51.579

	pointsOfFunctionPlot[680].X = 2.46
	pointsOfFunctionPlot[680].Y = 52.416

	pointsOfFunctionPlot[681].X = 2.47
	pointsOfFunctionPlot[681].Y = 53.266

	pointsOfFunctionPlot[682].X = 2.48
	pointsOfFunctionPlot[682].Y = 54.13

	pointsOfFunctionPlot[683].X = 2.49
	pointsOfFunctionPlot[683].Y = 55.009

	pointsOfFunctionPlot[684].X = 2.5
	pointsOfFunctionPlot[684].Y = 55.901

	pointsOfFunctionPlot[685].X = 2.51
	pointsOfFunctionPlot[685].Y = 56.808

	pointsOfFunctionPlot[686].X = 2.52
	pointsOfFunctionPlot[686].Y = 57.73

	pointsOfFunctionPlot[687].X = 2.53
	pointsOfFunctionPlot[687].Y = 58.667

	pointsOfFunctionPlot[688].X = 2.54
	pointsOfFunctionPlot[688].Y = 59.618

	pointsOfFunctionPlot[689].X = 2.55
	pointsOfFunctionPlot[689].Y = 60.586

	pointsOfFunctionPlot[690].X = 2.56
	pointsOfFunctionPlot[690].Y = 61.569

	pointsOfFunctionPlot[691].X = 2.57
	pointsOfFunctionPlot[691].Y = 62.568

	pointsOfFunctionPlot[692].X = 2.58
	pointsOfFunctionPlot[692].Y = 63.583

	pointsOfFunctionPlot[693].X = 2.59
	pointsOfFunctionPlot[693].Y = 64.614

	pointsOfFunctionPlot[694].X = 2.6
	pointsOfFunctionPlot[694].Y = 65.663

	pointsOfFunctionPlot[695].X = 2.61
	pointsOfFunctionPlot[695].Y = 66.728

	pointsOfFunctionPlot[696].X = 2.62
	pointsOfFunctionPlot[696].Y = 67.811

	pointsOfFunctionPlot[697].X = 2.63
	pointsOfFunctionPlot[697].Y = 69.911

	pointsOfFunctionPlot[698].X = 2.64
	pointsOfFunctionPlot[698].Y = 70.029

	pointsOfFunctionPlot[699].X = 2.65
	pointsOfFunctionPlot[699].Y = 71.165

	pointsOfFunctionPlot[700].X = 2.66
	pointsOfFunctionPlot[700].Y = 72.32

	pointsOfFunctionPlot[701].X = 2.67
	pointsOfFunctionPlot[701].Y = 73.493

	pointsOfFunctionPlot[702].X = 2.68
	pointsOfFunctionPlot[702].Y = 74.686

	pointsOfFunctionPlot[703].X = 2.69
	pointsOfFunctionPlot[703].Y = 75.897

	pointsOfFunctionPlot[704].X = 2.7
	pointsOfFunctionPlot[704].Y = 77.129

	pointsOfFunctionPlot[705].X = 2.71
	pointsOfFunctionPlot[705].Y = 78.38

	pointsOfFunctionPlot[706].X = 2.72
	pointsOfFunctionPlot[706].Y = 79.652

	pointsOfFunctionPlot[707].X = 2.73
	pointsOfFunctionPlot[707].Y = 80.944

	pointsOfFunctionPlot[708].X = 2.74
	pointsOfFunctionPlot[708].Y = 82.257

	pointsOfFunctionPlot[709].X = 2.75
	pointsOfFunctionPlot[709].Y = 83.592

	pointsOfFunctionPlot[710].X = 2.76
	pointsOfFunctionPlot[710].Y = 84.948

	pointsOfFunctionPlot[711].X = 2.77
	pointsOfFunctionPlot[711].Y = 86.327

	pointsOfFunctionPlot[712].X = 2.78
	pointsOfFunctionPlot[712].Y = 87.727

	pointsOfFunctionPlot[713].X = 2.79
	pointsOfFunctionPlot[713].Y = 89.151

	pointsOfFunctionPlot[714].X = 2.8
	pointsOfFunctionPlot[714].Y = 90.597

	pointsOfFunctionPlot[715].X = 2.81
	pointsOfFunctionPlot[715].Y = 92.067

	pointsOfFunctionPlot[716].X = 2.82
	pointsOfFunctionPlot[716].Y = 93.561

	pointsOfFunctionPlot[717].X = 2.83
	pointsOfFunctionPlot[717].Y = 95.079

	pointsOfFunctionPlot[718].X = 2.84
	pointsOfFunctionPlot[718].Y = 96.621

	pointsOfFunctionPlot[719].X = 2.85
	pointsOfFunctionPlot[719].Y = 98.189

	pointsOfFunctionPlot[720].X = 2.86
	pointsOfFunctionPlot[720].Y = 99.782

	pointsOfFunctionPlot[721].X = 2.87
	pointsOfFunctionPlot[721].Y = 101.401

	pointsOfFunctionPlot[722].X = 2.88
	pointsOfFunctionPlot[722].Y = 103.046

	pointsOfFunctionPlot[723].X = 2.89
	pointsOfFunctionPlot[723].Y = 104.718

	pointsOfFunctionPlot[724].X = 2.9
	pointsOfFunctionPlot[724].Y = 106.417

	pointsOfFunctionPlot[725].X = 2.91
	pointsOfFunctionPlot[725].Y = 108.144

	pointsOfFunctionPlot[726].X = 2.92
	pointsOfFunctionPlot[726].Y = 109.898

	pointsOfFunctionPlot[727].X = 2.93
	pointsOfFunctionPlot[727].Y = 111.681

	pointsOfFunctionPlot[728].X = 2.94
	pointsOfFunctionPlot[728].Y = 113.493

	pointsOfFunctionPlot[729].X = 2.95
	pointsOfFunctionPlot[729].Y = 115.335

	pointsOfFunctionPlot[730].X = 2.96
	pointsOfFunctionPlot[730].Y = 117.206

	pointsOfFunctionPlot[731].X = 2.97
	pointsOfFunctionPlot[731].Y = 119.108

	pointsOfFunctionPlot[732].X = 2.98
	pointsOfFunctionPlot[732].Y = 121.04

	pointsOfFunctionPlot[733].X = 2.99
	pointsOfFunctionPlot[733].Y = 123.004

	pointsOfFunctionPlot[734].X = 3.0
	pointsOfFunctionPlot[734].Y = 125.0

	pointsOfFunctionPlot[735].X = 3.01
	pointsOfFunctionPlot[735].Y = 127.028

	pointsOfFunctionPlot[736].X = 3.02
	pointsOfFunctionPlot[736].Y = 129.089

	pointsOfFunctionPlot[737].X = 3.03
	pointsOfFunctionPlot[737].Y = 131.183

	pointsOfFunctionPlot[738].X = 3.04
	pointsOfFunctionPlot[738].Y = 133.311

	pointsOfFunctionPlot[739].X = 3.05
	pointsOfFunctionPlot[739].Y = 135.474

	pointsOfFunctionPlot[740].X = 3.06
	pointsOfFunctionPlot[740].Y = 137.672

	pointsOfFunctionPlot[741].X = 3.07
	pointsOfFunctionPlot[741].Y = 139.906

	pointsOfFunctionPlot[742].X = 3.08
	pointsOfFunctionPlot[742].Y = 142.176

	pointsOfFunctionPlot[743].X = 3.09
	pointsOfFunctionPlot[743].Y = 144.483

	pointsOfFunctionPlot[744].X = 3.1
	pointsOfFunctionPlot[744].Y = 146.827

	pointsOfFunctionPlot[745].X = 3.11
	pointsOfFunctionPlot[745].Y = 149.209

	pointsOfFunctionPlot[746].X = 3.12
	pointsOfFunctionPlot[746].Y = 151.63

	pointsOfFunctionPlot[747].X = 3.13
	pointsOfFunctionPlot[747].Y = 154.09

	pointsOfFunctionPlot[748].X = 3.14
	pointsOfFunctionPlot[748].Y = 156.59

	pointsOfFunctionPlot[749].X = 3.15
	pointsOfFunctionPlot[749].Y = 159.131

	pointsOfFunctionPlot[750].X = 3.16
	pointsOfFunctionPlot[750].Y = 161.713

	pointsOfFunctionPlot[751].X = 3.17
	pointsOfFunctionPlot[751].Y = 164.336

	pointsOfFunctionPlot[752].X = 3.18
	pointsOfFunctionPlot[752].Y = 167.003

	pointsOfFunctionPlot[753].X = 3.19
	pointsOfFunctionPlot[753].Y = 169.712

	pointsOfFunctionPlot[754].X = 3.2
	pointsOfFunctionPlot[754].Y = 172.466

	pointsOfFunctionPlot[755].X = 3.21
	pointsOfFunctionPlot[755].Y = 175.264

	pointsOfFunctionPlot[756].X = 3.22
	pointsOfFunctionPlot[756].Y = 178.108

	pointsOfFunctionPlot[757].X = 3.23
	pointsOfFunctionPlot[757].Y = 180.997

	pointsOfFunctionPlot[758].X = 3.24
	pointsOfFunctionPlot[758].Y = 183.934

	pointsOfFunctionPlot[759].X = 3.25
	pointsOfFunctionPlot[759].Y = 186.918

	pointsOfFunctionPlot[760].X = 3.26
	pointsOfFunctionPlot[760].Y = 189.951

	pointsOfFunctionPlot[761].X = 3.27
	pointsOfFunctionPlot[761].Y = 193.033

	pointsOfFunctionPlot[762].X = 3.28
	pointsOfFunctionPlot[762].Y = 196.165

	pointsOfFunctionPlot[763].X = 3.29
	pointsOfFunctionPlot[763].Y = 199.347

	pointsOfFunctionPlot[764].X = 3.3
	pointsOfFunctionPlot[764].Y = 202.582

	pointsOfFunctionPlot[765].X = 3.31
	pointsOfFunctionPlot[765].Y = 205.868

	pointsOfFunctionPlot[766].X = 3.32
	pointsOfFunctionPlot[766].Y = 209.209

	pointsOfFunctionPlot[767].X = 3.33
	pointsOfFunctionPlot[767].Y = 212.603

	pointsOfFunctionPlot[768].X = 3.34
	pointsOfFunctionPlot[768].Y = 216.052

	pointsOfFunctionPlot[769].X = 3.35
	pointsOfFunctionPlot[769].Y = 219.558

	pointsOfFunctionPlot[770].X = 3.36
	pointsOfFunctionPlot[770].Y = 223.12

	pointsOfFunctionPlot[771].X = 3.37
	pointsOfFunctionPlot[771].Y = 226.74

	pointsOfFunctionPlot[772].X = 3.38
	pointsOfFunctionPlot[772].Y = 230.419

	pointsOfFunctionPlot[773].X = 3.39
	pointsOfFunctionPlot[773].Y = 234.157

	pointsOfFunctionPlot[774].X = 3.4
	pointsOfFunctionPlot[774].Y = 237.956

	pointsOfFunctionPlot[775].X = 3.41
	pointsOfFunctionPlot[775].Y = 241.817

	pointsOfFunctionPlot[776].X = 3.42
	pointsOfFunctionPlot[776].Y = 245.74

	pointsOfFunctionPlot[777].X = 3.43
	pointsOfFunctionPlot[777].Y = 249.727

	pointsOfFunctionPlot[778].X = 3.44
	pointsOfFunctionPlot[778].Y = 253.779

	pointsOfFunctionPlot[779].X = 3.45
	pointsOfFunctionPlot[779].Y = 257.897

	pointsOfFunctionPlot[780].X = 3.46
	pointsOfFunctionPlot[780].Y = 262.081

	pointsOfFunctionPlot[781].X = 3.47
	pointsOfFunctionPlot[781].Y = 266.333

	pointsOfFunctionPlot[782].X = 3.48
	pointsOfFunctionPlot[782].Y = 270.654

	pointsOfFunctionPlot[783].X = 3.49
	pointsOfFunctionPlot[783].Y = 275.046

	pointsOfFunctionPlot[784].X = 3.5
	pointsOfFunctionPlot[784].Y = 279.508

	pointsOfFunctionPlot[785].X = 3.51
	pointsOfFunctionPlot[785].Y = 284.043

	pointsOfFunctionPlot[786].X = 3.52
	pointsOfFunctionPlot[786].Y = 288.651

	pointsOfFunctionPlot[787].X = 3.53
	pointsOfFunctionPlot[787].Y = 293.335

	pointsOfFunctionPlot[788].X = 3.54
	pointsOfFunctionPlot[788].Y = 298.094

	pointsOfFunctionPlot[789].X = 3.55
	pointsOfFunctionPlot[789].Y = 302.93

	pointsOfFunctionPlot[790].X = 3.56
	pointsOfFunctionPlot[790].Y = 307.845

	pointsOfFunctionPlot[791].X = 3.57
	pointsOfFunctionPlot[791].Y = 312.84

	pointsOfFunctionPlot[792].X = 3.58
	pointsOfFunctionPlot[792].Y = 317.916

	pointsOfFunctionPlot[793].X = 3.59
	pointsOfFunctionPlot[793].Y = 323.074

	pointsOfFunctionPlot[794].X = 3.6
	pointsOfFunctionPlot[794].Y = 328.316

	pointsOfFunctionPlot[795].X = 3.61
	pointsOfFunctionPlot[795].Y = 333.642

	pointsOfFunctionPlot[796].X = 3.62
	pointsOfFunctionPlot[796].Y = 339.056

	pointsOfFunctionPlot[797].X = 3.63
	pointsOfFunctionPlot[797].Y = 344.557

	pointsOfFunctionPlot[798].X = 3.64
	pointsOfFunctionPlot[798].Y = 350.147

	pointsOfFunctionPlot[799].X = 3.65
	pointsOfFunctionPlot[799].Y = 355.828

	pointsOfFunctionPlot[800].X = 3.66
	pointsOfFunctionPlot[800].Y = 361.601

	pointsOfFunctionPlot[801].X = 3.67
	pointsOfFunctionPlot[801].Y = 367.468

	pointsOfFunctionPlot[802].X = 3.68
	pointsOfFunctionPlot[802].Y = 373.43

	pointsOfFunctionPlot[803].X = 3.69
	pointsOfFunctionPlot[803].Y = 379.489

	pointsOfFunctionPlot[804].X = 3.7
	pointsOfFunctionPlot[804].Y = 385.646

	pointsOfFunctionPlot[805].X = 3.71
	pointsOfFunctionPlot[805].Y = 391.903

	pointsOfFunctionPlot[806].X = 3.72
	pointsOfFunctionPlot[806].Y = 398.261

	pointsOfFunctionPlot[807].X = 3.73
	pointsOfFunctionPlot[807].Y = 404.723

	pointsOfFunctionPlot[808].X = 3.74
	pointsOfFunctionPlot[808].Y = 411.289

	pointsOfFunctionPlot[809].X = 3.75
	pointsOfFunctionPlot[809].Y = 417.962

	pointsOfFunctionPlot[810].X = 3.76
	pointsOfFunctionPlot[810].Y = 424.744

	pointsOfFunctionPlot[811].X = 3.77
	pointsOfFunctionPlot[811].Y = 431.635

	pointsOfFunctionPlot[812].X = 3.78
	pointsOfFunctionPlot[812].Y = 438.638

	pointsOfFunctionPlot[813].X = 3.79
	pointsOfFunctionPlot[813].Y = 445.755

	pointsOfFunctionPlot[814].X = 3.8
	pointsOfFunctionPlot[814].Y = 452.987

	pointsOfFunctionPlot[815].X = 3.81
	pointsOfFunctionPlot[815].Y = 460.336

	pointsOfFunctionPlot[816].X = 3.82
	pointsOfFunctionPlot[816].Y = 467.805

	pointsOfFunctionPlot[817].X = 3.83
	pointsOfFunctionPlot[817].Y = 475.395

	pointsOfFunctionPlot[818].X = 3.84
	pointsOfFunctionPlot[818].Y = 483.108

	pointsOfFunctionPlot[819].X = 3.85
	pointsOfFunctionPlot[819].Y = 490.946

	pointsOfFunctionPlot[820].X = 3.86
	pointsOfFunctionPlot[820].Y = 498.912

	pointsOfFunctionPlot[821].X = 3.87
	pointsOfFunctionPlot[821].Y = 507.007

	pointsOfFunctionPlot[822].X = 3.88
	pointsOfFunctionPlot[822].Y = 515.232

	pointsOfFunctionPlot[823].X = 3.89
	pointsOfFunctionPlot[823].Y = 523.592

	pointsOfFunctionPlot[824].X = 3.9
	pointsOfFunctionPlot[824].Y = 532.087

	pointsOfFunctionPlot[825].X = 3.91
	pointsOfFunctionPlot[825].Y = 540.72

	pointsOfFunctionPlot[826].X = 3.92
	pointsOfFunctionPlot[826].Y = 549.493

	pointsOfFunctionPlot[827].X = 3.93
	pointsOfFunctionPlot[827].Y = 558.408

	pointsOfFunctionPlot[828].X = 3.94
	pointsOfFunctionPlot[828].Y = 567.468

	pointsOfFunctionPlot[829].X = 3.95
	pointsOfFunctionPlot[829].Y = 576.675

	pointsOfFunctionPlot[830].X = 3.96
	pointsOfFunctionPlot[830].Y = 586.031

	pointsOfFunctionPlot[831].X = 3.97
	pointsOfFunctionPlot[831].Y = 595.54

	pointsOfFunctionPlot[832].X = 3.98
	pointsOfFunctionPlot[832].Y = 605.202

	pointsOfFunctionPlot[833].X = 3.99
	pointsOfFunctionPlot[833].Y = 615.021

	pointsOfFunctionPlot[834].X = 4.0
	pointsOfFunctionPlot[834].Y = 625.0

	pointsOfFunctionPlot[835].X = 4.01
	pointsOfFunctionPlot[835].Y = 635.14

	pointsOfFunctionPlot[836].X = 4.02
	pointsOfFunctionPlot[836].Y = 645.445

	pointsOfFunctionPlot[837].X = 4.03
	pointsOfFunctionPlot[837].Y = 655.917

	pointsOfFunctionPlot[838].X = 4.04
	pointsOfFunctionPlot[838].Y = 666.559

	pointsOfFunctionPlot[839].X = 4.05
	pointsOfFunctionPlot[839].Y = 677.374

	pointsOfFunctionPlot[840].X = 4.06
	pointsOfFunctionPlot[840].Y = 688.364

	pointsOfFunctionPlot[841].X = 4.07
	pointsOfFunctionPlot[841].Y = 699.532

	pointsOfFunctionPlot[842].X = 4.08
	pointsOfFunctionPlot[842].Y = 710.882

	pointsOfFunctionPlot[843].X = 4.09
	pointsOfFunctionPlot[843].Y = 722.415

	pointsOfFunctionPlot[844].X = 4.1
	pointsOfFunctionPlot[844].Y = 734.136

	pointsOfFunctionPlot[845].X = 4.11
	pointsOfFunctionPlot[845].Y = 746.047

	pointsOfFunctionPlot[846].X = 4.12
	pointsOfFunctionPlot[846].Y = 758.152

	pointsOfFunctionPlot[847].X = 4.13
	pointsOfFunctionPlot[847].Y = 770.452

	pointsOfFunctionPlot[848].X = 4.14
	pointsOfFunctionPlot[848].Y = 782.953

	pointsOfFunctionPlot[849].X = 4.15
	pointsOfFunctionPlot[849].Y = 795.656

	pointsOfFunctionPlot[850].X = 4.16
	pointsOfFunctionPlot[850].Y = 808.565

	pointsOfFunctionPlot[851].X = 4.17
	pointsOfFunctionPlot[851].Y = 821.684

	pointsOfFunctionPlot[852].X = 4.18
	pointsOfFunctionPlot[852].Y = 835.015

	pointsOfFunctionPlot[853].X = 4.19
	pointsOfFunctionPlot[853].Y = 848.563

	pointsOfFunctionPlot[854].X = 4.2
	pointsOfFunctionPlot[854].Y = 862.331

	pointsOfFunctionPlot[855].X = 4.21
	pointsOfFunctionPlot[855].Y = 876.322

	pointsOfFunctionPlot[856].X = 4.22
	pointsOfFunctionPlot[856].Y = 890.54

	pointsOfFunctionPlot[857].X = 4.23
	pointsOfFunctionPlot[857].Y = 904.988

	pointsOfFunctionPlot[858].X = 4.24
	pointsOfFunctionPlot[858].Y = 919.671

	pointsOfFunctionPlot[859].X = 4.25
	pointsOfFunctionPlot[859].Y = 934.593

	pointsOfFunctionPlot[860].X = 4.26
	pointsOfFunctionPlot[860].Y = 949.756

	pointsOfFunctionPlot[861].X = 4.27
	pointsOfFunctionPlot[861].Y = 965.165

	pointsOfFunctionPlot[862].X = 4.28
	pointsOfFunctionPlot[862].Y = 980.825

	pointsOfFunctionPlot[863].X = 4.29
	pointsOfFunctionPlot[863].Y = 996.738

	pointsOfFunctionPlot[864].X = 4.3
	pointsOfFunctionPlot[864].Y = 1_012.91

	pointsOfFunctionPlot[865].X = 4.31
	pointsOfFunctionPlot[865].Y = 1_029.344

	pointsOfFunctionPlot[866].X = 4.32
	pointsOfFunctionPlot[866].Y = 1_046.045

	pointsOfFunctionPlot[867].X = 4.33
	pointsOfFunctionPlot[867].Y = 1_063.016

	pointsOfFunctionPlot[868].X = 4.34
	pointsOfFunctionPlot[868].Y = 1_080.263

	pointsOfFunctionPlot[869].X = 4.35
	pointsOfFunctionPlot[869].Y = 1_097.79

	pointsOfFunctionPlot[870].X = 4.36
	pointsOfFunctionPlot[870].Y = 1_115.601

	pointsOfFunctionPlot[871].X = 4.37
	pointsOfFunctionPlot[871].Y = 1_133.702

	pointsOfFunctionPlot[872].X = 4.38
	pointsOfFunctionPlot[872].Y = 1_152.095

	pointsOfFunctionPlot[873].X = 4.39
	pointsOfFunctionPlot[873].Y = 1_170.788

	pointsOfFunctionPlot[874].X = 4.4
	pointsOfFunctionPlot[874].Y = 1_189.783

	pointsOfFunctionPlot[875].X = 4.41
	pointsOfFunctionPlot[875].Y = 1_209.087

	pointsOfFunctionPlot[876].X = 4.42
	pointsOfFunctionPlot[876].Y = 1_228.704

	pointsOfFunctionPlot[877].X = 4.43
	pointsOfFunctionPlot[877].Y = 1_248.639

	pointsOfFunctionPlot[878].X = 4.44
	pointsOfFunctionPlot[878].Y = 1_268.898

	pointsOfFunctionPlot[879].X = 4.45
	pointsOfFunctionPlot[879].Y = 1_289.485

	pointsOfFunctionPlot[880].X = 4.46
	pointsOfFunctionPlot[880].Y = 1_310.407

	pointsOfFunctionPlot[881].X = 4.47
	pointsOfFunctionPlot[881].Y = 1_331.667

	pointsOfFunctionPlot[882].X = 4.48
	pointsOfFunctionPlot[882].Y = 1_353.273

	pointsOfFunctionPlot[883].X = 4.49
	pointsOfFunctionPlot[883].Y = 1_375.229

	pointsOfFunctionPlot[884].X = 4.5
	pointsOfFunctionPlot[884].Y = 1_397.542

	pointsOfFunctionPlot[885].X = 4.51
	pointsOfFunctionPlot[885].Y = 1_420.217

	pointsOfFunctionPlot[886].X = 4.52
	pointsOfFunctionPlot[886].Y = 1_443.259

	pointsOfFunctionPlot[887].X = 4.53
	pointsOfFunctionPlot[887].Y = 1_466.675

	pointsOfFunctionPlot[888].X = 4.54
	pointsOfFunctionPlot[888].Y = 1_490.472

	pointsOfFunctionPlot[889].X = 4.55
	pointsOfFunctionPlot[889].Y = 1_514.654

	pointsOfFunctionPlot[890].X = 4.56
	pointsOfFunctionPlot[890].Y = 1_539.228

	pointsOfFunctionPlot[891].X = 4.57
	pointsOfFunctionPlot[891].Y = 1_564.202

	pointsOfFunctionPlot[892].X = 4.58
	pointsOfFunctionPlot[892].Y = 1_589.58

	pointsOfFunctionPlot[893].X = 4.59
	pointsOfFunctionPlot[893].Y = 1_615.371

	pointsOfFunctionPlot[894].X = 4.6
	pointsOfFunctionPlot[894].Y = 1_641.579

	pointsOfFunctionPlot[895].X = 4.61
	pointsOfFunctionPlot[895].Y = 1_668.213

	pointsOfFunctionPlot[896].X = 4.62
	pointsOfFunctionPlot[896].Y = 1_695.279

	pointsOfFunctionPlot[897].X = 4.63
	pointsOfFunctionPlot[897].Y = 1_722.785

	pointsOfFunctionPlot[898].X = 4.64
	pointsOfFunctionPlot[898].Y = 1_750.736

	pointsOfFunctionPlot[899].X = 4.65
	pointsOfFunctionPlot[899].Y = 1_779.141

	pointsOfFunctionPlot[900].X = 4.66
	pointsOfFunctionPlot[900].Y = 1_808.007

	pointsOfFunctionPlot[901].X = 4.67
	pointsOfFunctionPlot[901].Y = 1_837.341

	pointsOfFunctionPlot[902].X = 4.68
	pointsOfFunctionPlot[902].Y = 1_867.151

	pointsOfFunctionPlot[903].X = 4.69
	pointsOfFunctionPlot[903].Y = 1_897.445

	pointsOfFunctionPlot[904].X = 4.7
	pointsOfFunctionPlot[904].Y = 1_928.23

	pointsOfFunctionPlot[905].X = 4.71
	pointsOfFunctionPlot[905].Y = 1_959.515

	pointsOfFunctionPlot[906].X = 4.72
	pointsOfFunctionPlot[906].Y = 1_991.307

	pointsOfFunctionPlot[907].X = 4.73
	pointsOfFunctionPlot[907].Y = 2_023.616

	pointsOfFunctionPlot[908].X = 4.74
	pointsOfFunctionPlot[908].Y = 2_056.448

	pointsOfFunctionPlot[909].X = 4.75
	pointsOfFunctionPlot[909].Y = 2_089.813

	pointsOfFunctionPlot[910].X = 4.76
	pointsOfFunctionPlot[910].Y = 2_123.719

	pointsOfFunctionPlot[911].X = 4.77
	pointsOfFunctionPlot[911].Y = 2_158.176

	pointsOfFunctionPlot[912].X = 4.78
	pointsOfFunctionPlot[912].Y = 2_193.191

	pointsOfFunctionPlot[913].X = 4.79
	pointsOfFunctionPlot[913].Y = 2_228.775

	pointsOfFunctionPlot[914].X = 4.8
	pointsOfFunctionPlot[914].Y = 2_264.936

	pointsOfFunctionPlot[915].X = 4.81
	pointsOfFunctionPlot[915].Y = 2_301.684

	pointsOfFunctionPlot[916].X = 4.82
	pointsOfFunctionPlot[916].Y = 2_339.028

	pointsOfFunctionPlot[917].X = 4.83
	pointsOfFunctionPlot[917].Y = 2_376.977

	pointsOfFunctionPlot[918].X = 4.84
	pointsOfFunctionPlot[918].Y = 2_415.543

	pointsOfFunctionPlot[919].X = 4.85
	pointsOfFunctionPlot[919].Y = 2_454.734

	pointsOfFunctionPlot[920].X = 4.86
	pointsOfFunctionPlot[920].Y = 2_494.561

	pointsOfFunctionPlot[921].X = 4.87
	pointsOfFunctionPlot[921].Y = 2_535.034

	pointsOfFunctionPlot[922].X = 4.88
	pointsOfFunctionPlot[922].Y = 2_576.164

	pointsOfFunctionPlot[923].X = 4.89
	pointsOfFunctionPlot[923].Y = 2_617.961

	pointsOfFunctionPlot[924].X = 4.9
	pointsOfFunctionPlot[924].Y = 2_660.437

	pointsOfFunctionPlot[925].X = 4.91
	pointsOfFunctionPlot[925].Y = 2_703.601

	pointsOfFunctionPlot[926].X = 4.92
	pointsOfFunctionPlot[926].Y = 2_747.466

	pointsOfFunctionPlot[927].X = 4.93
	pointsOfFunctionPlot[927].Y = 2_792.043

	pointsOfFunctionPlot[928].X = 4.94
	pointsOfFunctionPlot[928].Y = 2_837.342

	pointsOfFunctionPlot[929].X = 4.95
	pointsOfFunctionPlot[929].Y = 2_883.377

	pointsOfFunctionPlot[930].X = 4.96
	pointsOfFunctionPlot[930].Y = 2_930.159

	pointsOfFunctionPlot[931].X = 4.97
	pointsOfFunctionPlot[931].Y = 2_977.699

	pointsOfFunctionPlot[932].X = 4.98
	pointsOfFunctionPlot[932].Y = 3_026.011

	pointsOfFunctionPlot[933].X = 4.99
	pointsOfFunctionPlot[933].Y = 3_075.107

	pointsOfFunctionPlot[934].X = 5.0
	pointsOfFunctionPlot[934].Y = 3_125.0

	pointsOfFunctionPlot[935].X = 5.01
	pointsOfFunctionPlot[935].Y = 3_175.701

	pointsOfFunctionPlot[936].X = 5.02
	pointsOfFunctionPlot[936].Y = 3_227.226

	pointsOfFunctionPlot[937].X = 5.03
	pointsOfFunctionPlot[937].Y = 3_279.586

	pointsOfFunctionPlot[938].X = 5.04
	pointsOfFunctionPlot[938].Y = 3_332.796

	pointsOfFunctionPlot[939].X = 5.05
	pointsOfFunctionPlot[939].Y = 3_386.87

	pointsOfFunctionPlot[940].X = 5.06
	pointsOfFunctionPlot[940].Y = 3_441.82

	pointsOfFunctionPlot[941].X = 5.07
	pointsOfFunctionPlot[941].Y = 3_497.662

	pointsOfFunctionPlot[942].X = 5.08
	pointsOfFunctionPlot[942].Y = 3_554.41

	pointsOfFunctionPlot[943].X = 5.09
	pointsOfFunctionPlot[943].Y = 3_612.079

	pointsOfFunctionPlot[944].X = 5.1
	pointsOfFunctionPlot[944].Y = 3_670.684

	pointsOfFunctionPlot[945].X = 5.11
	pointsOfFunctionPlot[945].Y = 3_730.239

	pointsOfFunctionPlot[946].X = 5.12
	pointsOfFunctionPlot[946].Y = 3_790.761

	pointsOfFunctionPlot[947].X = 5.13
	pointsOfFunctionPlot[947].Y = 3_852.264

	pointsOfFunctionPlot[948].X = 5.14
	pointsOfFunctionPlot[948].Y = 3_914.766

	pointsOfFunctionPlot[949].X = 5.15
	pointsOfFunctionPlot[949].Y = 3_978.281

	pointsOfFunctionPlot[950].X = 5.16
	pointsOfFunctionPlot[950].Y = 4_042.827

	pointsOfFunctionPlot[951].X = 5.17
	pointsOfFunctionPlot[951].Y = 4_108.42

	pointsOfFunctionPlot[952].X = 5.18
	pointsOfFunctionPlot[952].Y = 4_175.078

	pointsOfFunctionPlot[953].X = 5.19
	pointsOfFunctionPlot[953].Y = 4_242.817

	pointsOfFunctionPlot[954].X = 5.2
	pointsOfFunctionPlot[954].Y = 4_311.655

	pointsOfFunctionPlot[955].X = 5.21
	pointsOfFunctionPlot[955].Y = 4_381.61

	pointsOfFunctionPlot[956].X = 5.22
	pointsOfFunctionPlot[956].Y = 4_452.699

	pointsOfFunctionPlot[957].X = 5.23
	pointsOfFunctionPlot[957].Y = 4_524.943

	pointsOfFunctionPlot[958].X = 5.24
	pointsOfFunctionPlot[958].Y = 4_598.358

	pointsOfFunctionPlot[959].X = 5.25
	pointsOfFunctionPlot[959].Y = 4_672.964

	pointsOfFunctionPlot[960].X = 5.26
	pointsOfFunctionPlot[960].Y = 4_748.781

	pointsOfFunctionPlot[961].X = 5.27
	pointsOfFunctionPlot[961].Y = 4_825.828

	pointsOfFunctionPlot[962].X = 5.28
	pointsOfFunctionPlot[962].Y = 4_904.126

	pointsOfFunctionPlot[963].X = 5.29
	pointsOfFunctionPlot[963].Y = 4_983.693

	pointsOfFunctionPlot[964].X = 5.3
	pointsOfFunctionPlot[964].Y = 5_064.551

	pointsOfFunctionPlot[965].X = 5.31
	pointsOfFunctionPlot[965].Y = 5_146.722

	pointsOfFunctionPlot[966].X = 5.32
	pointsOfFunctionPlot[966].Y = 5_230.225

	pointsOfFunctionPlot[967].X = 5.33
	pointsOfFunctionPlot[967].Y = 5_315.083

	pointsOfFunctionPlot[968].X = 5.34
	pointsOfFunctionPlot[968].Y = 5_401.318

	pointsOfFunctionPlot[969].X = 5.35
	pointsOfFunctionPlot[969].Y = 5_488.953

	pointsOfFunctionPlot[970].X = 5.36
	pointsOfFunctionPlot[970].Y = 5_578.009

	pointsOfFunctionPlot[971].X = 5.37
	pointsOfFunctionPlot[971].Y = 5_668.51

	pointsOfFunctionPlot[972].X = 5.38
	pointsOfFunctionPlot[972].Y = 5_760.479

	pointsOfFunctionPlot[973].X = 5.39
	pointsOfFunctionPlot[973].Y = 5_853.94

	pointsOfFunctionPlot[974].X = 5.4
	pointsOfFunctionPlot[974].Y = 5_948.918

	pointsOfFunctionPlot[975].X = 5.41
	pointsOfFunctionPlot[975].Y = 6_045.437

	pointsOfFunctionPlot[976].X = 5.42
	pointsOfFunctionPlot[976].Y = 6_143.522

	pointsOfFunctionPlot[977].X = 5.43
	pointsOfFunctionPlot[977].Y = 6_243.198

	pointsOfFunctionPlot[978].X = 5.44
	pointsOfFunctionPlot[978].Y = 6_344.491

	pointsOfFunctionPlot[979].X = 5.45
	pointsOfFunctionPlot[979].Y = 6_447.428

	pointsOfFunctionPlot[980].X = 5.46
	pointsOfFunctionPlot[980].Y = 6_552.035

	pointsOfFunctionPlot[981].X = 5.47
	pointsOfFunctionPlot[981].Y = 6_658.339

	pointsOfFunctionPlot[982].X = 5.48
	pointsOfFunctionPlot[982].Y = 6_766.368

	pointsOfFunctionPlot[983].X = 5.49
	pointsOfFunctionPlot[983].Y = 6_876.712

	pointsOfFunctionPlot[984].X = 5.5
	pointsOfFunctionPlot[984].Y = 6_987.712

	pointsOfFunctionPlot[985].X = 5.51
	pointsOfFunctionPlot[985].Y = 7_101.085

	pointsOfFunctionPlot[986].X = 5.52
	pointsOfFunctionPlot[986].Y = 7_216.297

	pointsOfFunctionPlot[987].X = 5.53
	pointsOfFunctionPlot[987].Y = 7_333.378

	pointsOfFunctionPlot[988].X = 5.54
	pointsOfFunctionPlot[988].Y = 7_452.36

	pointsOfFunctionPlot[989].X = 5.55
	pointsOfFunctionPlot[989].Y = 7_573.271

	pointsOfFunctionPlot[990].X = 5.56
	pointsOfFunctionPlot[990].Y = 7_696.144

	pointsOfFunctionPlot[991].X = 5.57
	pointsOfFunctionPlot[991].Y = 7_821.011

	pointsOfFunctionPlot[992].X = 5.58
	pointsOfFunctionPlot[992].Y = 7_947.904

	pointsOfFunctionPlot[993].X = 5.59
	pointsOfFunctionPlot[993].Y = 8_076.855

	pointsOfFunctionPlot[994].X = 5.6
	pointsOfFunctionPlot[994].Y = 8_207.899

	pointsOfFunctionPlot[995].X = 5.61
	pointsOfFunctionPlot[995].Y = 8_341.069

	pointsOfFunctionPlot[996].X = 5.62
	pointsOfFunctionPlot[996].Y = 8_476.399

	pointsOfFunctionPlot[997].X = 5.63
	pointsOfFunctionPlot[997].Y = 8_613.925

	pointsOfFunctionPlot[998].X = 5.64
	pointsOfFunctionPlot[998].Y = 8_753.683

	pointsOfFunctionPlot[999].X = 5.65
	pointsOfFunctionPlot[999].Y = 8_895.708

	pointsOfFunctionPlot[1_000].X = 5.66
	pointsOfFunctionPlot[1_000].Y = 9_040.037

	pointsOfFunctionPlot[1_001].X = 5.67
	pointsOfFunctionPlot[1_001].Y = 9_186.708

	pointsOfFunctionPlot[1_002].X = 5.68
	pointsOfFunctionPlot[1_002].Y = 9_335.758

	pointsOfFunctionPlot[1_003].X = 5.69
	pointsOfFunctionPlot[1_003].Y = 9_487.227

	pointsOfFunctionPlot[1_004].X = 5.7
	pointsOfFunctionPlot[1_004].Y = 9_641.154

	pointsOfFunctionPlot[1_005].X = 5.71
	pointsOfFunctionPlot[1_005].Y = 9_797.577

	pointsOfFunctionPlot[1_006].X = 5.72
	pointsOfFunctionPlot[1_006].Y = 9_956.539

	pointsOfFunctionPlot[1_007].X = 5.73
	pointsOfFunctionPlot[1_007].Y = 10_118.08

	pointsOfFunctionPlot[1_008].X = 5.74
	pointsOfFunctionPlot[1_008].Y = 10_282.242

	pointsOfFunctionPlot[1_009].X = 5.75
	pointsOfFunctionPlot[1_009].Y = 10_449.067

	pointsOfFunctionPlot[1_010].X = 5.76
	pointsOfFunctionPlot[1_010].Y = 10_618.599

	pointsOfFunctionPlot[1_011].X = 5.77
	pointsOfFunctionPlot[1_011].Y = 10_790.881

	pointsOfFunctionPlot[1_012].X = 5.78
	pointsOfFunctionPlot[1_012].Y = 10_965.959

	pointsOfFunctionPlot[1_013].X = 5.79
	pointsOfFunctionPlot[1_013].Y = 11_143.877

	pointsOfFunctionPlot[1_014].X = 5.8
	pointsOfFunctionPlot[1_014].Y = 11_324.682

	pointsOfFunctionPlot[1_015].X = 5.81
	pointsOfFunctionPlot[1_015].Y = 11_508.42

	pointsOfFunctionPlot[1_016].X = 5.82
	pointsOfFunctionPlot[1_016].Y = 11_695.14

	pointsOfFunctionPlot[1_017].X = 5.83
	pointsOfFunctionPlot[1_017].Y = 11_884.888

	pointsOfFunctionPlot[1_018].X = 5.84
	pointsOfFunctionPlot[1_018].Y = 12_077.716

	pointsOfFunctionPlot[1_019].X = 5.85
	pointsOfFunctionPlot[1_019].Y = 12_273.672

	pointsOfFunctionPlot[1_020].X = 5.86
	pointsOfFunctionPlot[1_020].Y = 12_472.807

	pointsOfFunctionPlot[1_021].X = 5.87
	pointsOfFunctionPlot[1_021].Y = 12_675.173

	pointsOfFunctionPlot[1_022].X = 5.88
	pointsOfFunctionPlot[1_022].Y = 12_880.823

	pointsOfFunctionPlot[1_023].X = 5.89
	pointsOfFunctionPlot[1_023].Y = 13_089.809

	pointsOfFunctionPlot[1_024].X = 5.9
	pointsOfFunctionPlot[1_024].Y = 13_302.186

	pointsOfFunctionPlot[1_025].X = 5.91
	pointsOfFunctionPlot[1_025].Y = 13_518.008

	pointsOfFunctionPlot[1_026].X = 5.92
	pointsOfFunctionPlot[1_026].Y = 13_737.333

	pointsOfFunctionPlot[1_027].X = 5.93
	pointsOfFunctionPlot[1_027].Y = 13_960.215

	pointsOfFunctionPlot[1_028].X = 5.94
	pointsOfFunctionPlot[1_028].Y = 14_186.714

	pointsOfFunctionPlot[1_029].X = 5.95
	pointsOfFunctionPlot[1_029].Y = 14_416.888

	pointsOfFunctionPlot[1_030].X = 5.96
	pointsOfFunctionPlot[1_030].Y = 14_650.796

	pointsOfFunctionPlot[1_031].X = 5.97
	pointsOfFunctionPlot[1_031].Y = 14_888.499

	pointsOfFunctionPlot[1_032].X = 5.98
	pointsOfFunctionPlot[1_032].Y = 15_130.059

	pointsOfFunctionPlot[1_033].X = 5.99
	pointsOfFunctionPlot[1_033].Y = 15_375.538

	pointsOfFunctionPlot[1_034].X = 6.0
	pointsOfFunctionPlot[1_034].Y = 15_625.0

	pointsOfFunctionPlot[1_035].X = 6.01
	pointsOfFunctionPlot[1_035].Y = 15_878.509

	pointsOfFunctionPlot[1_036].X = 6.02
	pointsOfFunctionPlot[1_036].Y = 16_136.131

	pointsOfFunctionPlot[1_037].X = 6.03
	pointsOfFunctionPlot[1_037].Y = 16_397.933

	pointsOfFunctionPlot[1_038].X = 6.04
	pointsOfFunctionPlot[1_038].Y = 16_663.983

	pointsOfFunctionPlot[1_039].X = 6.05
	pointsOfFunctionPlot[1_039].Y = 16_934.349

	pointsOfFunctionPlot[1_040].X = 6.06
	pointsOfFunctionPlot[1_040].Y = 17_209.102

	pointsOfFunctionPlot[1_041].X = 6.07
	pointsOfFunctionPlot[1_041].Y = 17_488.313

	pointsOfFunctionPlot[1_042].X = 6.08
	pointsOfFunctionPlot[1_042].Y = 17_772.054

	pointsOfFunctionPlot[1_043].X = 6.09
	pointsOfFunctionPlot[1_043].Y = 18_060.398

	pointsOfFunctionPlot[1_044].X = 6.1
	pointsOfFunctionPlot[1_044].Y = 18_353.421

	pointsOfFunctionPlot[1_045].X = 6.11
	pointsOfFunctionPlot[1_045].Y = 18_651.197

	pointsOfFunctionPlot[1_046].X = 6.12
	pointsOfFunctionPlot[1_046].Y = 18_953.805

	pointsOfFunctionPlot[1_047].X = 6.13
	pointsOfFunctionPlot[1_047].Y = 19_261.323

	pointsOfFunctionPlot[1_048].X = 6.14
	pointsOfFunctionPlot[1_048].Y = 19_573.83

	pointsOfFunctionPlot[1_049].X = 6.15
	pointsOfFunctionPlot[1_049].Y = 19_891.408

	pointsOfFunctionPlot[1_050].X = 6.16
	pointsOfFunctionPlot[1_050].Y = 20_214.138

	pointsOfFunctionPlot[1_051].X = 6.17
	pointsOfFunctionPlot[1_051].Y = 20_542.104

	pointsOfFunctionPlot[1_052].X = 6.18
	pointsOfFunctionPlot[1_052].Y = 20_875.391

	pointsOfFunctionPlot[1_053].X = 6.19
	pointsOfFunctionPlot[1_053].Y = 21_214.086

	pointsOfFunctionPlot[1_054].X = 6.2
	pointsOfFunctionPlot[1_054].Y = 21_558.276

	pointsOfFunctionPlot[1_055].X = 6.21
	pointsOfFunctionPlot[1_055].Y = 21_908.05

	pointsOfFunctionPlot[1_056].X = 6.22
	pointsOfFunctionPlot[1_056].Y = 22_263.499

	pointsOfFunctionPlot[1_057].X = 6.23
	pointsOfFunctionPlot[1_057].Y = 22_624.715

	pointsOfFunctionPlot[1_058].X = 6.24
	pointsOfFunctionPlot[1_058].Y = 22_991.792

	pointsOfFunctionPlot[1_059].X = 6.25
	pointsOfFunctionPlot[1_059].Y = 23_364.824

	pointsOfFunctionPlot[1_060].X = 6.26
	pointsOfFunctionPlot[1_060].Y = 23_743.909

	pointsOfFunctionPlot[1_061].X = 6.27
	pointsOfFunctionPlot[1_061].Y = 24_129.144

	pointsOfFunctionPlot[1_062].X = 6.28
	pointsOfFunctionPlot[1_062].Y = 24_520.63

	pointsOfFunctionPlot[1_063].X = 6.29
	pointsOfFunctionPlot[1_063].Y = 24_918.467

	pointsOfFunctionPlot[1_064].X = 6.3
	pointsOfFunctionPlot[1_064].Y = 25_322.759

	pointsOfFunctionPlot[1_065].X = 6.31
	pointsOfFunctionPlot[1_065].Y = 25_733.61

	pointsOfFunctionPlot[1_066].X = 6.32
	pointsOfFunctionPlot[1_066].Y = 26_151.128

	pointsOfFunctionPlot[1_067].X = 6.33
	pointsOfFunctionPlot[1_067].Y = 26_575.419

	pointsOfFunctionPlot[1_068].X = 6.34
	pointsOfFunctionPlot[1_068].Y = 27_006.594

	pointsOfFunctionPlot[1_069].X = 6.35
	pointsOfFunctionPlot[1_069].Y = 27_444.765

	pointsOfFunctionPlot[1_070].X = 6.36
	pointsOfFunctionPlot[1_070].Y = 27_890.045

	pointsOfFunctionPlot[1_071].X = 6.37
	pointsOfFunctionPlot[1_071].Y = 28_342.55

	pointsOfFunctionPlot[1_072].X = 6.38
	pointsOfFunctionPlot[1_072].Y = 28_802.396

	pointsOfFunctionPlot[1_073].X = 6.39
	pointsOfFunctionPlot[1_073].Y = 29_269.703

	pointsOfFunctionPlot[1_074].X = 6.4
	pointsOfFunctionPlot[1_074].Y = 29_744.592

	pointsOfFunctionPlot[1_075].X = 6.41
	pointsOfFunctionPlot[1_075].Y = 30_227.186

	pointsOfFunctionPlot[1_076].X = 6.42
	pointsOfFunctionPlot[1_076].Y = 30_717.61

	pointsOfFunctionPlot[1_077].X = 6.43
	pointsOfFunctionPlot[1_077].Y = 31_215.991

	pointsOfFunctionPlot[1_078].X = 6.44
	pointsOfFunctionPlot[1_078].Y = 31_722.457

	pointsOfFunctionPlot[1_079].X = 6.45
	pointsOfFunctionPlot[1_079].Y = 32_237.141

	pointsOfFunctionPlot[1_080].X = 6.46
	pointsOfFunctionPlot[1_080].Y = 32_760.176

	pointsOfFunctionPlot[1_081].X = 6.47
	pointsOfFunctionPlot[1_081].Y = 33_291.696

	pointsOfFunctionPlot[1_082].X = 6.48
	pointsOfFunctionPlot[1_082].Y = 33_831.84

	pointsOfFunctionPlot[1_083].X = 6.49
	pointsOfFunctionPlot[1_083].Y = 34_380.748

	pointsOfFunctionPlot[1_084].X = 6.5
	pointsOfFunctionPlot[1_084].Y = 34_938.562

	pointsOfFunctionPlot[1_085].X = 6.51
	pointsOfFunctionPlot[1_085].Y = 35_505.426

	pointsOfFunctionPlot[1_086].X = 6.52
	pointsOfFunctionPlot[1_086].Y = 36_081.487

	pointsOfFunctionPlot[1_087].X = 6.53
	pointsOfFunctionPlot[1_087].Y = 36_666.894

	pointsOfFunctionPlot[1_088].X = 6.54
	pointsOfFunctionPlot[1_088].Y = 37_261.799

	pointsOfFunctionPlot[1_089].X = 6.55
	pointsOfFunctionPlot[1_089].Y = 37_866.357

	pointsOfFunctionPlot[1_090].X = 6.56
	pointsOfFunctionPlot[1_090].Y = 38_480.723

	pointsOfFunctionPlot[1_091].X = 6.57
	pointsOfFunctionPlot[1_091].Y = 39_105.057

	pointsOfFunctionPlot[1_092].X = 6.58
	pointsOfFunctionPlot[1_092].Y = 39_739.521

	pointsOfFunctionPlot[1_093].X = 6.59
	pointsOfFunctionPlot[1_093].Y = 40_384.278

	pointsOfFunctionPlot[1_094].X = 6.6
	pointsOfFunctionPlot[1_094].Y = 41_039.496

	pointsOfFunctionPlot[1_095].X = 6.61
	pointsOfFunctionPlot[1_095].Y = 41_705.346

	pointsOfFunctionPlot[1_096].X = 6.62
	pointsOfFunctionPlot[1_096].Y = 42_381.998

	pointsOfFunctionPlot[1_097].X = 6.63
	pointsOfFunctionPlot[1_097].Y = 43_069.628

	pointsOfFunctionPlot[1_098].X = 6.64
	pointsOfFunctionPlot[1_098].Y = 43_768.415

	pointsOfFunctionPlot[1_099].X = 6.65
	pointsOfFunctionPlot[1_099].Y = 44_478.54

	pointsOfFunctionPlot[1_100].X = 6.66
	pointsOfFunctionPlot[1_100].Y = 45_200.186

	pointsOfFunctionPlot[1_101].X = 6.67
	pointsOfFunctionPlot[1_101].Y = 45_933.541

	pointsOfFunctionPlot[1_102].X = 6.68
	pointsOfFunctionPlot[1_102].Y = 46_678.794

	pointsOfFunctionPlot[1_103].X = 6.69
	pointsOfFunctionPlot[1_103].Y = 47_436.138

	pointsOfFunctionPlot[1_104].X = 6.7
	pointsOfFunctionPlot[1_104].Y = 48_205.77

	pointsOfFunctionPlot[1_105].X = 6.71
	pointsOfFunctionPlot[1_105].Y = 48_987.889

	pointsOfFunctionPlot[1_106].X = 6.72
	pointsOfFunctionPlot[1_106].Y = 49_782.697

	pointsOfFunctionPlot[1_107].X = 6.73
	pointsOfFunctionPlot[1_107].Y = 50_590.401

	pointsOfFunctionPlot[1_108].X = 6.74
	pointsOfFunctionPlot[1_108].Y = 51_411.21

	pointsOfFunctionPlot[1_109].X = 6.75
	pointsOfFunctionPlot[1_109].Y = 52_245.336

	pointsOfFunctionPlot[1_110].X = 6.76
	pointsOfFunctionPlot[1_110].Y = 53_092.995

	pointsOfFunctionPlot[1_111].X = 6.77
	pointsOfFunctionPlot[1_111].Y = 53_954.407

	pointsOfFunctionPlot[1_112].X = 6.78
	pointsOfFunctionPlot[1_112].Y = 54_829.795

	pointsOfFunctionPlot[1_113].X = 6.79
	pointsOfFunctionPlot[1_113].Y = 55_719.386

	pointsOfFunctionPlot[1_114].X = 6.8
	pointsOfFunctionPlot[1_114].Y = 56_623.411

	pointsOfFunctionPlot[1_115].X = 6.81
	pointsOfFunctionPlot[1_115].Y = 57_542.102

	pointsOfFunctionPlot[1_116].X = 6.82
	pointsOfFunctionPlot[1_116].Y = 58_475.7

	pointsOfFunctionPlot[1_117].X = 6.83
	pointsOfFunctionPlot[1_117].Y = 59_424.444

	pointsOfFunctionPlot[1_118].X = 6.84
	pointsOfFunctionPlot[1_118].Y = 60_388.581

	pointsOfFunctionPlot[1_119].X = 6.85
	pointsOfFunctionPlot[1_119].Y = 61_368.361

	pointsOfFunctionPlot[1_120].X = 6.86
	pointsOfFunctionPlot[1_120].Y = 62_364.038

	pointsOfFunctionPlot[1_121].X = 6.87
	pointsOfFunctionPlot[1_121].Y = 63_375.869

	pointsOfFunctionPlot[1_122].X = 6.88
	pointsOfFunctionPlot[1_122].Y = 64_404.116

	pointsOfFunctionPlot[1_123].X = 6.89
	pointsOfFunctionPlot[1_123].Y = 65_449.047

	pointsOfFunctionPlot[1_124].X = 6.9
	pointsOfFunctionPlot[1_124].Y = 66_510.931

	pointsOfFunctionPlot[1_125].X = 6.91
	pointsOfFunctionPlot[1_125].Y = 67_590.044

	pointsOfFunctionPlot[1_126].X = 6.92
	pointsOfFunctionPlot[1_126].Y = 68_686.665

	pointsOfFunctionPlot[1_127].X = 6.93
	pointsOfFunctionPlot[1_127].Y = 69_801.078

	pointsOfFunctionPlot[1_128].X = 6.94
	pointsOfFunctionPlot[1_128].Y = 70_933.572

	pointsOfFunctionPlot[1_129].X = 6.95
	pointsOfFunctionPlot[1_129].Y = 72_084.44

	pointsOfFunctionPlot[1_130].X = 6.96
	pointsOfFunctionPlot[1_130].Y = 73_253.98

	pointsOfFunctionPlot[1_131].X = 6.97
	pointsOfFunctionPlot[1_131].Y = 74_442.496

	pointsOfFunctionPlot[1_132].X = 6.98
	pointsOfFunctionPlot[1_132].Y = 75_650.295

	pointsOfFunctionPlot[1_133].X = 6.99
	pointsOfFunctionPlot[1_133].Y = 76_877.69

	pointsOfFunctionPlot[1_134].X = 7.0
	pointsOfFunctionPlot[1_134].Y = 78_125.0

	pointsOfFunctionPlot[1_135].X = 7.01
	pointsOfFunctionPlot[1_135].Y = 79_392.546

	pointsOfFunctionPlot[1_136].X = 7.02
	pointsOfFunctionPlot[1_136].Y = 80_680.657

	pointsOfFunctionPlot[1_137].X = 7.03
	pointsOfFunctionPlot[1_137].Y = 81_989.668

	pointsOfFunctionPlot[1_138].X = 7.04
	pointsOfFunctionPlot[1_138].Y = 83_319.917

	pointsOfFunctionPlot[1_139].X = 7.05
	pointsOfFunctionPlot[1_139].Y = 84_671.749

	pointsOfFunctionPlot[1_140].X = 7.06
	pointsOfFunctionPlot[1_140].Y = 86_045.513

	pointsOfFunctionPlot[1_141].X = 7.07
	pointsOfFunctionPlot[1_141].Y = 87_441.566

	pointsOfFunctionPlot[1_142].X = 7.08
	pointsOfFunctionPlot[1_142].Y = 88_860.27

	pointsOfFunctionPlot[1_143].X = 7.09
	pointsOfFunctionPlot[1_143].Y = 90_301.992

	pointsOfFunctionPlot[1_144].X = 7.1
	pointsOfFunctionPlot[1_144].Y = 91_767.104

	pointsOfFunctionPlot[1_145].X = 7.11
	pointsOfFunctionPlot[1_145].Y = 93_255.988

	pointsOfFunctionPlot[1_146].X = 7.12
	pointsOfFunctionPlot[1_146].Y = 94_769.029

	pointsOfFunctionPlot[1_147].X = 7.13
	pointsOfFunctionPlot[1_147].Y = 96_306.617

	pointsOfFunctionPlot[1_148].X = 7.14
	pointsOfFunctionPlot[1_148].Y = 97_869.153

	pointsOfFunctionPlot[1_149].X = 7.15
	pointsOfFunctionPlot[1_149].Y = 99_457.04

	pointsOfFunctionPlot[1_150].X = 7.16
	pointsOfFunctionPlot[1_150].Y = 101_070.69

	pointsOfFunctionPlot[1_151].X = 7.17
	pointsOfFunctionPlot[1_151].Y = 102_710.52

	pointsOfFunctionPlot[1_152].X = 7.18
	pointsOfFunctionPlot[1_152].Y = 104_376.957

	pointsOfFunctionPlot[1_153].X = 7.19
	pointsOfFunctionPlot[1_153].Y = 106_070.43

	pointsOfFunctionPlot[1_154].X = 7.2
	pointsOfFunctionPlot[1_154].Y = 107_791.379

	pointsOfFunctionPlot[1_155].X = 7.21
	pointsOfFunctionPlot[1_155].Y = 109_540.25

	pointsOfFunctionPlot[1_156].X = 7.22
	pointsOfFunctionPlot[1_156].Y = 111_317.496

	pointsOfFunctionPlot[1_157].X = 7.23
	pointsOfFunctionPlot[1_157].Y = 113_123.577

	pointsOfFunctionPlot[1_158].X = 7.24
	pointsOfFunctionPlot[1_158].Y = 114_958.961

	pointsOfFunctionPlot[1_159].X = 7.25
	pointsOfFunctionPlot[1_159].Y = 116_824.123

	pointsOfFunctionPlot[1_160].X = 7.26
	pointsOfFunctionPlot[1_160].Y = 118_719.547

	pointsOfFunctionPlot[1_161].X = 7.27
	pointsOfFunctionPlot[1_161].Y = 120_645.723

	pointsOfFunctionPlot[1_162].X = 7.28
	pointsOfFunctionPlot[1_162].Y = 122_603.15

	pointsOfFunctionPlot[1_163].X = 7.29
	pointsOfFunctionPlot[1_163].Y = 124_592.336

	pointsOfFunctionPlot[1_164].X = 7.3
	pointsOfFunctionPlot[1_164].Y = 126_613.796

	pointsOfFunctionPlot[1_165].X = 7.31
	pointsOfFunctionPlot[1_165].Y = 128_668.053

	pointsOfFunctionPlot[1_166].X = 7.32
	pointsOfFunctionPlot[1_166].Y = 130_755.64

	pointsOfFunctionPlot[1_167].X = 7.33
	pointsOfFunctionPlot[1_167].Y = 132_877.097

	pointsOfFunctionPlot[1_168].X = 7.34
	pointsOfFunctionPlot[1_168].Y = 135_032.973

	pointsOfFunctionPlot[1_169].X = 7.35
	pointsOfFunctionPlot[1_169].Y = 137_223.828

	pointsOfFunctionPlot[1_170].X = 7.36
	pointsOfFunctionPlot[1_170].Y = 139_450.229

	pointsOfFunctionPlot[1_171].X = 7.37
	pointsOfFunctionPlot[1_171].Y = 141_712.752

	pointsOfFunctionPlot[1_172].X = 7.38
	pointsOfFunctionPlot[1_172].Y = 144_011.983

	pointsOfFunctionPlot[1_173].X = 7.39
	pointsOfFunctionPlot[1_173].Y = 146_348.519

	pointsOfFunctionPlot[1_174].X = 7.4
	pointsOfFunctionPlot[1_174].Y = 148_722.964

	pointsOfFunctionPlot[1_175].X = 7.41
	pointsOfFunctionPlot[1_175].Y = 151_135.933

	pointsOfFunctionPlot[1_176].X = 7.42
	pointsOfFunctionPlot[1_176].Y = 153_588.052

	pointsOfFunctionPlot[1_177].X = 7.43
	pointsOfFunctionPlot[1_177].Y = 156_079.955

	pointsOfFunctionPlot[1_178].X = 7.44
	pointsOfFunctionPlot[1_178].Y = 158_612.288

	pointsOfFunctionPlot[1_179].X = 7.45
	pointsOfFunctionPlot[1_179].Y = 161_185.708

	pointsOfFunctionPlot[1_180].X = 7.46
	pointsOfFunctionPlot[1_180].Y = 163_800.88

	pointsOfFunctionPlot[1_181].X = 7.47
	pointsOfFunctionPlot[1_181].Y = 166_458.483

	pointsOfFunctionPlot[1_182].X = 7.48
	pointsOfFunctionPlot[1_182].Y = 169_159.203

	pointsOfFunctionPlot[1_183].X = 7.49
	pointsOfFunctionPlot[1_183].Y = 171_903.742

	pointsOfFunctionPlot[1_184].X = 7.5
	pointsOfFunctionPlot[1_184].Y = 174_692.81

	pointsOfFunctionPlot[1_185].X = 7.51
	pointsOfFunctionPlot[1_185].Y = 177_527.13

	pointsOfFunctionPlot[1_186].X = 7.52
	pointsOfFunctionPlot[1_186].Y = 180_407.435

	pointsOfFunctionPlot[1_187].X = 7.53
	pointsOfFunctionPlot[1_187].Y = 183_334.472

	pointsOfFunctionPlot[1_188].X = 7.54
	pointsOfFunctionPlot[1_188].Y = 186_308.999

	pointsOfFunctionPlot[1_189].X = 7.55
	pointsOfFunctionPlot[1_189].Y = 189_331.786

	pointsOfFunctionPlot[1_190].X = 7.56
	pointsOfFunctionPlot[1_190].Y = 192_403.617

	pointsOfFunctionPlot[1_191].X = 7.57
	pointsOfFunctionPlot[1_191].Y = 195_525.287

	pointsOfFunctionPlot[1_192].X = 7.58
	pointsOfFunctionPlot[1_192].Y = 198_697.605

	pointsOfFunctionPlot[1_193].X = 7.59
	pointsOfFunctionPlot[1_193].Y = 201_921.392

	pointsOfFunctionPlot[1_194].X = 7.6
	pointsOfFunctionPlot[1_194].Y = 205_197.484

	pointsOfFunctionPlot[1_195].X = 7.61
	pointsOfFunctionPlot[1_195].Y = 208_526.73

	pointsOfFunctionPlot[1_196].X = 7.62
	pointsOfFunctionPlot[1_196].Y = 211_909.991

	pointsOfFunctionPlot[1_197].X = 7.63
	pointsOfFunctionPlot[1_197].Y = 215_348.144

	pointsOfFunctionPlot[1_198].X = 7.64
	pointsOfFunctionPlot[1_198].Y = 218_842.079

	pointsOfFunctionPlot[1_199].X = 7.65
	pointsOfFunctionPlot[1_199].Y = 222_392.702

	pointsOfFunctionPlot[1_200].X = 7.66
	pointsOfFunctionPlot[1_200].Y = 226_000.933

	pointsOfFunctionPlot[1_201].X = 7.67
	pointsOfFunctionPlot[1_201].Y = 229_667.706

	pointsOfFunctionPlot[1_202].X = 7.68
	pointsOfFunctionPlot[1_202].Y = 233_393.971

	pointsOfFunctionPlot[1_203].X = 7.69
	pointsOfFunctionPlot[1_203].Y = 237_180.692

	pointsOfFunctionPlot[1_204].X = 7.7
	pointsOfFunctionPlot[1_204].Y = 241_028.852

	pointsOfFunctionPlot[1_205].X = 7.71
	pointsOfFunctionPlot[1_205].Y = 244_939.447

	pointsOfFunctionPlot[1_206].X = 7.72
	pointsOfFunctionPlot[1_206].Y = 248_913.489

	pointsOfFunctionPlot[1_207].X = 7.73
	pointsOfFunctionPlot[1_207].Y = 252_952.009

	pointsOfFunctionPlot[1_208].X = 7.74
	pointsOfFunctionPlot[1_208].Y = 257_056.052

	pointsOfFunctionPlot[1_209].X = 7.75
	pointsOfFunctionPlot[1_209].Y = 261_226.681

	pointsOfFunctionPlot[1_210].X = 7.76
	pointsOfFunctionPlot[1_210].Y = 265_464.977

	pointsOfFunctionPlot[1_211].X = 7.77
	pointsOfFunctionPlot[1_211].Y = 269_772.038

	pointsOfFunctionPlot[1_212].X = 7.78
	pointsOfFunctionPlot[1_212].Y = 274_148.979

	pointsOfFunctionPlot[1_213].X = 7.79
	pointsOfFunctionPlot[1_213].Y = 278_596.934

	pointsOfFunctionPlot[1_214].X = 7.8
	pointsOfFunctionPlot[1_214].Y = 283_117.056

	pointsOfFunctionPlot[1_215].X = 7.81
	pointsOfFunctionPlot[1_215].Y = 287_710.514

	pointsOfFunctionPlot[1_216].X = 7.82
	pointsOfFunctionPlot[1_216].Y = 292_378.5

	pointsOfFunctionPlot[1_217].X = 7.83
	pointsOfFunctionPlot[1_217].Y = 297_122.221

	pointsOfFunctionPlot[1_218].X = 7.84
	pointsOfFunctionPlot[1_218].Y = 301_942.908

	pointsOfFunctionPlot[1_219].X = 7.85
	pointsOfFunctionPlot[1_219].Y = 306_841.808

	pointsOfFunctionPlot[1_220].X = 7.86
	pointsOfFunctionPlot[1_220].Y = 311_820.191

	pointsOfFunctionPlot[1_221].X = 7.87
	pointsOfFunctionPlot[1_221].Y = 316_879.346

	pointsOfFunctionPlot[1_222].X = 7.88
	pointsOfFunctionPlot[1_222].Y = 322_020.584

	pointsOfFunctionPlot[1_223].X = 7.89
	pointsOfFunctionPlot[1_223].Y = 327_245.237

	pointsOfFunctionPlot[1_224].X = 7.9
	pointsOfFunctionPlot[1_224].Y = 332_554.657

	pointsOfFunctionPlot[1_225].X = 7.91
	pointsOfFunctionPlot[1_225].Y = 337_950.22

	pointsOfFunctionPlot[1_226].X = 7.92
	pointsOfFunctionPlot[1_226].Y = 343_433.324

	pointsOfFunctionPlot[1_227].X = 7.93
	pointsOfFunctionPlot[1_227].Y = 349_005.39

	pointsOfFunctionPlot[1_228].X = 7.94
	pointsOfFunctionPlot[1_228].Y = 354_667.859

	pointsOfFunctionPlot[1_229].X = 7.95
	pointsOfFunctionPlot[1_229].Y = 360_422.201

	pointsOfFunctionPlot[1_230].X = 7.96
	pointsOfFunctionPlot[1_230].Y = 366_269.903

	pointsOfFunctionPlot[1_231].X = 7.97
	pointsOfFunctionPlot[1_231].Y = 372_212.483

	pointsOfFunctionPlot[1_232].X = 7.98
	pointsOfFunctionPlot[1_232].Y = 378_251.478

	pointsOfFunctionPlot[1_233].X = 7.99
	pointsOfFunctionPlot[1_233].Y = 384_388.454

	pointsOfFunctionPlot[1_234].X = 8.0
	pointsOfFunctionPlot[1_234].Y = 390_625.0

	pointsOfFunctionPlot[1_235].X = 8.01
	pointsOfFunctionPlot[1_235].Y = 396_962.731

	pointsOfFunctionPlot[1_236].X = 8.02
	pointsOfFunctionPlot[1_236].Y = 403_403.289

	pointsOfFunctionPlot[1_237].X = 8.03
	pointsOfFunctionPlot[1_237].Y = 409_948.342

	pointsOfFunctionPlot[1_238].X = 8.04
	pointsOfFunctionPlot[1_238].Y = 416_599.586

	pointsOfFunctionPlot[1_239].X = 8.05
	pointsOfFunctionPlot[1_239].Y = 423_358.744

	pointsOfFunctionPlot[1_240].X = 8.06
	pointsOfFunctionPlot[1_240].Y = 430_227.567

	pointsOfFunctionPlot[1_241].X = 8.07
	pointsOfFunctionPlot[1_241].Y = 437_207.833

	pointsOfFunctionPlot[1_242].X = 8.08
	pointsOfFunctionPlot[1_242].Y = 444_301.352

	pointsOfFunctionPlot[1_243].X = 8.09
	pointsOfFunctionPlot[1_243].Y = 451_509.86

	pointsOfFunctionPlot[1_244].X = 8.1
	pointsOfFunctionPlot[1_244].Y = 458_835.524

	pointsOfFunctionPlot[1_245].X = 8.11
	pointsOfFunctionPlot[1_245].Y = 466_279.943

	pointsOfFunctionPlot[1_246].X = 8.12
	pointsOfFunctionPlot[1_246].Y = 473_845.145

	pointsOfFunctionPlot[1_247].X = 8.13
	pointsOfFunctionPlot[1_247].Y = 481_533.088

	pointsOfFunctionPlot[1_248].X = 8.14
	pointsOfFunctionPlot[1_248].Y = 489_345.766

	pointsOfFunctionPlot[1_249].X = 8.15
	pointsOfFunctionPlot[1_249].Y = 497_285.201

	pointsOfFunctionPlot[1_250].X = 8.16
	pointsOfFunctionPlot[1_250].Y = 505_353.45

	pointsOfFunctionPlot[1_251].X = 8.17
	pointsOfFunctionPlot[1_251].Y = 513_552.603

	pointsOfFunctionPlot[1_252].X = 8.18
	pointsOfFunctionPlot[1_252].Y = 521_884.784

	pointsOfFunctionPlot[1_253].X = 8.19
	pointsOfFunctionPlot[1_253].Y = 530_352.152

	pointsOfFunctionPlot[1_254].X = 8.2
	pointsOfFunctionPlot[1_254].Y = 538_956.899

	pointsOfFunctionPlot[1_255].X = 8.21
	pointsOfFunctionPlot[1_255].Y = 547_701.254

	pointsOfFunctionPlot[1_256].X = 8.22
	pointsOfFunctionPlot[1_256].Y = 556_587.483

	pointsOfFunctionPlot[1_257].X = 8.23
	pointsOfFunctionPlot[1_257].Y = 565_617.887

	pointsOfFunctionPlot[1_258].X = 8.24
	pointsOfFunctionPlot[1_258].Y = 574_794.806

	pointsOfFunctionPlot[1_259].X = 8.25
	pointsOfFunctionPlot[1_259].Y = 584_120.617

	pointsOfFunctionPlot[1_260].X = 8.26
	pointsOfFunctionPlot[1_260].Y = 593_597.735

	pointsOfFunctionPlot[1_261].X = 8.27
	pointsOfFunctionPlot[1_261].Y = 603_228.616

	pointsOfFunctionPlot[1_262].X = 8.28
	pointsOfFunctionPlot[1_262].Y = 613_015.754

	pointsOfFunctionPlot[1_263].X = 8.29
	pointsOfFunctionPlot[1_263].Y = 622_961.684

	pointsOfFunctionPlot[1_264].X = 8.3
	pointsOfFunctionPlot[1_264].Y = 633_068.983

	pointsOfFunctionPlot[1_265].X = 8.31
	pointsOfFunctionPlot[1_265].Y = 643_340.268

	pointsOfFunctionPlot[1_266].X = 8.32
	pointsOfFunctionPlot[1_266].Y = 653_778.201

	pointsOfFunctionPlot[1_267].X = 8.33
	pointsOfFunctionPlot[1_267].Y = 664_385.485

	pointsOfFunctionPlot[1_268].X = 8.34
	pointsOfFunctionPlot[1_268].Y = 675_165.868

	pointsOfFunctionPlot[1_269].X = 8.35
	pointsOfFunctionPlot[1_269].Y = 686_119.142

	pointsOfFunctionPlot[1_270].X = 8.36
	pointsOfFunctionPlot[1_270].Y = 697_251.145

	pointsOfFunctionPlot[1_271].X = 8.37
	pointsOfFunctionPlot[1_271].Y = 708_563.76

	pointsOfFunctionPlot[1_272].X = 8.38
	pointsOfFunctionPlot[1_272].Y = 720_059.917

	pointsOfFunctionPlot[1_273].X = 8.39
	pointsOfFunctionPlot[1_273].Y = 731_742.595

	pointsOfFunctionPlot[1_274].X = 8.4
	pointsOfFunctionPlot[1_274].Y = 743_614.819

	pointsOfFunctionPlot[1_275].X = 8.41
	pointsOfFunctionPlot[1_275].Y = 755_679.666

	pointsOfFunctionPlot[1_276].X = 8.42
	pointsOfFunctionPlot[1_276].Y = 767_940.26

	pointsOfFunctionPlot[1_277].X = 8.43
	pointsOfFunctionPlot[1_277].Y = 780_399.776

	pointsOfFunctionPlot[1_278].X = 8.44
	pointsOfFunctionPlot[1_278].Y = 793_061.444

	pointsOfFunctionPlot[1_279].X = 8.45
	pointsOfFunctionPlot[1_279].Y = 805_928.542

	pointsOfFunctionPlot[1_280].X = 8.46
	pointsOfFunctionPlot[1_280].Y = 819_004.403

	pointsOfFunctionPlot[1_281].X = 8.47
	pointsOfFunctionPlot[1_281].Y = 832_292.414

	pointsOfFunctionPlot[1_282].X = 8.48
	pointsOfFunctionPlot[1_282].Y = 845_796.019

	pointsOfFunctionPlot[1_283].X = 8.49
	pointsOfFunctionPlot[1_283].Y = 858_518.713

	pointsOfFunctionPlot[1_284].X = 8.5
	pointsOfFunctionPlot[1_284].Y = 873_464.053

	pointsOfFunctionPlot[1_285].X = 8.51
	pointsOfFunctionPlot[1_285].Y = 887_635.651

	pointsOfFunctionPlot[1_286].X = 8.52
	pointsOfFunctionPlot[1_286].Y = 902_037.176

	pointsOfFunctionPlot[1_287].X = 8.53
	pointsOfFunctionPlot[1_287].Y = 916_672.361

	pointsOfFunctionPlot[1_288].X = 8.54
	pointsOfFunctionPlot[1_288].Y = 931_544.995

	pointsOfFunctionPlot[1_289].X = 8.55
	pointsOfFunctionPlot[1_289].Y = 946_658.932

	pointsOfFunctionPlot[1_290].X = 8.56
	pointsOfFunctionPlot[1_290].Y = 962_018.086

	pointsOfFunctionPlot[1_291].X = 8.57
	pointsOfFunctionPlot[1_291].Y = 977_626.436

	pointsOfFunctionPlot[1_292].X = 8.58
	pointsOfFunctionPlot[1_292].Y = 993_488.026

	pointsOfFunctionPlot[1_293].X = 8.59
	pointsOfFunctionPlot[1_293].Y = 1_009_606.963

	pointsOfFunctionPlot[1_294].X = 8.6
	pointsOfFunctionPlot[1_294].Y = 1_025_987.423

	pointsOfFunctionPlot[1_295].X = 8.61
	pointsOfFunctionPlot[1_295].Y = 1_042_633.65

	pointsOfFunctionPlot[1_296].X = 8.62
	pointsOfFunctionPlot[1_296].Y = 1_059_549.955

	pointsOfFunctionPlot[1_297].X = 8.63
	pointsOfFunctionPlot[1_297].Y = 1_076_740.72

	pointsOfFunctionPlot[1_298].X = 8.64
	pointsOfFunctionPlot[1_298].Y = 1_094_210.398

	pointsOfFunctionPlot[1_299].X = 8.65
	pointsOfFunctionPlot[1_299].Y = 1_111_963.514

	pointsOfFunctionPlot[1_300].X = 8.66
	pointsOfFunctionPlot[1_300].Y = 1_130_004.668

	pointsOfFunctionPlot[1_301].X = 8.67
	pointsOfFunctionPlot[1_301].Y = 1_148_338.531

	pointsOfFunctionPlot[1_302].X = 8.68
	pointsOfFunctionPlot[1_302].Y = 1_166_969.855

	pointsOfFunctionPlot[1_303].X = 8.69
	pointsOfFunctionPlot[1_303].Y = 1_185_903.464

	pointsOfFunctionPlot[1_304].X = 8.7
	pointsOfFunctionPlot[1_304].Y = 1_205_144.263

	pointsOfFunctionPlot[1_305].X = 8.71
	pointsOfFunctionPlot[1_305].Y = 1_224_697.236

	pointsOfFunctionPlot[1_306].X = 8.72
	pointsOfFunctionPlot[1_306].Y = 1_244_567.448

	pointsOfFunctionPlot[1_307].X = 8.73
	pointsOfFunctionPlot[1_307].Y = 1_264_760.046

	pointsOfFunctionPlot[1_308].X = 8.74
	pointsOfFunctionPlot[1_308].Y = 1_285_280.261

	pointsOfFunctionPlot[1_309].X = 8.75
	pointsOfFunctionPlot[1_309].Y = 1_306_133.408

	pointsOfFunctionPlot[1_310].X = 8.76
	pointsOfFunctionPlot[1_310].Y = 1_327_324.888

	pointsOfFunctionPlot[1_311].X = 8.77
	pointsOfFunctionPlot[1_311].Y = 1_348_860.192

	pointsOfFunctionPlot[1_312].X = 8.78
	pointsOfFunctionPlot[1_312].Y = 1_370_744.897

	pointsOfFunctionPlot[1_313].X = 8.79
	pointsOfFunctionPlot[1_313].Y = 1_392_984.673

	pointsOfFunctionPlot[1_314].X = 8.8
	pointsOfFunctionPlot[1_314].Y = 1_415_585.28

	pointsOfFunctionPlot[1_315].X = 8.81
	pointsOfFunctionPlot[1_315].Y = 1_415_585.28

	pointsOfFunctionPlot[1_316].X = 8.82
	pointsOfFunctionPlot[1_316].Y = 1_461_892.5

	pointsOfFunctionPlot[1_317].X = 8.83
	pointsOfFunctionPlot[1_317].Y = 1_485_611.109

	pointsOfFunctionPlot[1_318].X = 8.84
	pointsOfFunctionPlot[1_318].Y = 1_509_714.542

	pointsOfFunctionPlot[1_319].X = 8.85
	pointsOfFunctionPlot[1_319].Y = 1_534_209.043

	pointsOfFunctionPlot[1_320].X = 8.86
	pointsOfFunctionPlot[1_320].Y = 1_559_100.958

	pointsOfFunctionPlot[1_321].X = 8.87
	pointsOfFunctionPlot[1_321].Y = 1_584_396.733

	pointsOfFunctionPlot[1_322].X = 8.88
	pointsOfFunctionPlot[1_322].Y = 1_610_102.923

	pointsOfFunctionPlot[1_323].X = 8.89
	pointsOfFunctionPlot[1_323].Y = 1_636_226.185

	pointsOfFunctionPlot[1_324].X = 8.9
	pointsOfFunctionPlot[1_324].Y = 1_662_773.286

	pointsOfFunctionPlot[1_325].X = 8.91
	pointsOfFunctionPlot[1_325].Y = 1_689_751.103

	pointsOfFunctionPlot[1_326].X = 8.92
	pointsOfFunctionPlot[1_326].Y = 1_717_166.624

	pointsOfFunctionPlot[1_327].X = 8.93
	pointsOfFunctionPlot[1_327].Y = 1_745_026.95

	pointsOfFunctionPlot[1_328].X = 8.94
	pointsOfFunctionPlot[1_328].Y = 1_773_339.299

	pointsOfFunctionPlot[1_329].X = 8.95
	pointsOfFunctionPlot[1_329].Y = 1_802_111.005

	pointsOfFunctionPlot[1_330].X = 8.96
	pointsOfFunctionPlot[1_330].Y = 1_831_349.519

	pointsOfFunctionPlot[1_331].X = 8.97
	pointsOfFunctionPlot[1_331].Y = 1_861_062.417

	pointsOfFunctionPlot[1_332].X = 8.98
	pointsOfFunctionPlot[1_332].Y = 1_891_257.394

	pointsOfFunctionPlot[1_333].X = 8.99
	pointsOfFunctionPlot[1_333].Y = 1_921_942.272

	pointsOfFunctionPlot[1_334].X = 9.0
	pointsOfFunctionPlot[1_334].Y = 1_953_125.0

	pointsOfFunctionPlot[1_335].X = 9.01
	pointsOfFunctionPlot[1_335].Y = 1_984_813.654

	pointsOfFunctionPlot[1_336].X = 9.02
	pointsOfFunctionPlot[1_336].Y = 2_017_016.445

	pointsOfFunctionPlot[1_337].X = 9.03
	pointsOfFunctionPlot[1_337].Y = 2_049_741.712

	pointsOfFunctionPlot[1_338].X = 9.04
	pointsOfFunctionPlot[1_338].Y = 2_082_997.934

	pointsOfFunctionPlot[1_339].X = 9.05
	pointsOfFunctionPlot[1_339].Y = 2_116_793.724

	pointsOfFunctionPlot[1_340].X = 9.06
	pointsOfFunctionPlot[1_340].Y = 2_151_137.837

	pointsOfFunctionPlot[1_341].X = 9.07
	pointsOfFunctionPlot[1_341].Y = 2_186_039.169

	pointsOfFunctionPlot[1_342].X = 9.08
	pointsOfFunctionPlot[1_342].Y = 2_221_506.761

	pointsOfFunctionPlot[1_343].X = 9.09
	pointsOfFunctionPlot[1_343].Y = 2_257_549.8

	pointsOfFunctionPlot[1_344].X = 9.1
	pointsOfFunctionPlot[1_344].Y = 2_294_177.623

	pointsOfFunctionPlot[1_345].X = 9.11
	pointsOfFunctionPlot[1_345].Y = 2_331_399.717

	pointsOfFunctionPlot[1_346].X = 9.12
	pointsOfFunctionPlot[1_346].Y = 2_369_225.724

	pointsOfFunctionPlot[1_347].X = 9.13
	pointsOfFunctionPlot[1_347].Y = 2_407_665.444

	pointsOfFunctionPlot[1_348].X = 9.14
	pointsOfFunctionPlot[1_348].Y = 2_446_728.831

	pointsOfFunctionPlot[1_349].X = 9.15
	pointsOfFunctionPlot[1_349].Y = 2_486_426.006

	pointsOfFunctionPlot[1_350].X = 9.16
	pointsOfFunctionPlot[1_350].Y = 2_526_767.252

	pointsOfFunctionPlot[1_351].X = 9.17
	pointsOfFunctionPlot[1_351].Y = 2_567_763.018

	pointsOfFunctionPlot[1_352].X = 9.18
	pointsOfFunctionPlot[1_352].Y = 2_609_423.924

	pointsOfFunctionPlot[1_353].X = 9.19
	pointsOfFunctionPlot[1_353].Y = 2_651_760.76

	pointsOfFunctionPlot[1_354].X = 9.2
	pointsOfFunctionPlot[1_354].Y = 2_694_784.495

	pointsOfFunctionPlot[1_355].X = 9.21
	pointsOfFunctionPlot[1_355].Y = 2_738_506.272

	pointsOfFunctionPlot[1_356].X = 9.22
	pointsOfFunctionPlot[1_356].Y = 2_782_937.417

	pointsOfFunctionPlot[1_357].X = 9.23
	pointsOfFunctionPlot[1_357].Y = 2_828_089.439

	pointsOfFunctionPlot[1_358].X = 9.24
	pointsOfFunctionPlot[1_358].Y = 2_873_974.034

	pointsOfFunctionPlot[1_359].X = 9.25
	pointsOfFunctionPlot[1_359].Y = 2_920_603.088

	pointsOfFunctionPlot[1_360].X = 9.26
	pointsOfFunctionPlot[1_360].Y = 2_967_988.679

	pointsOfFunctionPlot[1_361].X = 9.27
	pointsOfFunctionPlot[1_361].Y = 3_016_143.082

	pointsOfFunctionPlot[1_362].X = 9.28
	pointsOfFunctionPlot[1_362].Y = 3_065_078.771

	pointsOfFunctionPlot[1_363].X = 9.29
	pointsOfFunctionPlot[1_363].Y = 3_114_808.421

	pointsOfFunctionPlot[1_364].X = 9.3
	pointsOfFunctionPlot[1_364].Y = 3_165_344.915

	pointsOfFunctionPlot[1_365].X = 9.31
	pointsOfFunctionPlot[1_365].Y = 3_216_701.342

	pointsOfFunctionPlot[1_366].X = 9.32
	pointsOfFunctionPlot[1_366].Y = 3_268_891.007

	pointsOfFunctionPlot[1_367].X = 9.33
	pointsOfFunctionPlot[1_367].Y = 3_321_927.427

	pointsOfFunctionPlot[1_368].X = 9.34
	pointsOfFunctionPlot[1_368].Y = 3_375_824.342

	pointsOfFunctionPlot[1_369].X = 9.35
	pointsOfFunctionPlot[1_369].Y = 3_430_595.712

	pointsOfFunctionPlot[1_370].X = 9.36
	pointsOfFunctionPlot[1_370].Y = 3_486_255.726

	pointsOfFunctionPlot[1_371].X = 9.37
	pointsOfFunctionPlot[1_371].Y = 3_542_818.8

	pointsOfFunctionPlot[1_372].X = 9.38
	pointsOfFunctionPlot[1_372].Y = 3_600_299.587

	pointsOfFunctionPlot[1_373].X = 9.39
	pointsOfFunctionPlot[1_373].Y = 3_658_712.976

	pointsOfFunctionPlot[1_374].X = 9.4
	pointsOfFunctionPlot[1_374].Y = 3_718_074.099

	pointsOfFunctionPlot[1_375].X = 9.41
	pointsOfFunctionPlot[1_375].Y = 3_778_398.331

	pointsOfFunctionPlot[1_376].X = 9.42
	pointsOfFunctionPlot[1_376].Y = 3_839_701.3

	pointsOfFunctionPlot[1_377].X = 9.43
	pointsOfFunctionPlot[1_377].Y = 3_901_998.884

	pointsOfFunctionPlot[1_378].X = 9.44
	pointsOfFunctionPlot[1_378].Y = 3_965_307.221

	pointsOfFunctionPlot[1_379].X = 9.45
	pointsOfFunctionPlot[1_379].Y = 4_029_642.71

	pointsOfFunctionPlot[1_380].X = 9.46
	pointsOfFunctionPlot[1_380].Y = 4_095_022.016

	pointsOfFunctionPlot[1_381].X = 9.47
	pointsOfFunctionPlot[1_381].Y = 4_161_462.074

	pointsOfFunctionPlot[1_382].X = 9.48
	pointsOfFunctionPlot[1_382].Y = 4_228_980.095

	pointsOfFunctionPlot[1_383].X = 9.49
	pointsOfFunctionPlot[1_383].Y = 4_297_593.569

	pointsOfFunctionPlot[1_384].X = 9.5
	pointsOfFunctionPlot[1_384].Y = 4_367_320.268

	pointsOfFunctionPlot[1_385].X = 9.51
	pointsOfFunctionPlot[1_385].Y = 4_438_178.254

	pointsOfFunctionPlot[1_386].X = 9.52
	pointsOfFunctionPlot[1_386].Y = 4_510_185.883

	pointsOfFunctionPlot[1_387].X = 9.53
	pointsOfFunctionPlot[1_387].Y = 4_583_361.805

	pointsOfFunctionPlot[1_388].X = 9.54
	pointsOfFunctionPlot[1_388].Y = 4_657_724.977

	pointsOfFunctionPlot[1_389].X = 9.55
	pointsOfFunctionPlot[1_389].Y = 4_733_294.661

	pointsOfFunctionPlot[1_390].X = 9.56
	pointsOfFunctionPlot[1_390].Y = 4_810_090.432

	pointsOfFunctionPlot[1_391].X = 9.57
	pointsOfFunctionPlot[1_391].Y = 4_888_132.183

	pointsOfFunctionPlot[1_392].X = 9.58
	pointsOfFunctionPlot[1_392].Y = 4_967_440.13

	pointsOfFunctionPlot[1_393].X = 9.59
	pointsOfFunctionPlot[1_393].Y = 5_048_034.816

	pointsOfFunctionPlot[1_394].X = 9.6
	pointsOfFunctionPlot[1_394].Y = 5_129_937.118

	pointsOfFunctionPlot[1_395].X = 9.61
	pointsOfFunctionPlot[1_395].Y = 5_213_168.25

	pointsOfFunctionPlot[1_396].X = 9.62
	pointsOfFunctionPlot[1_396].Y = 5_297_749.775

	pointsOfFunctionPlot[1_397].X = 9.63
	pointsOfFunctionPlot[1_397].Y = 5_383_703.599

	pointsOfFunctionPlot[1_398].X = 9.64
	pointsOfFunctionPlot[1_398].Y = 5_471_051.99

	pointsOfFunctionPlot[1_399].X = 9.65
	pointsOfFunctionPlot[1_399].Y = 5_559_817.572

	pointsOfFunctionPlot[1_400].X = 9.66
	pointsOfFunctionPlot[1_400].Y = 5_650_023.34

	pointsOfFunctionPlot[1_401].X = 9.67
	pointsOfFunctionPlot[1_401].Y = 5_741_692.659

	pointsOfFunctionPlot[1_402].X = 9.68
	pointsOfFunctionPlot[1_402].Y = 5_834_849.276

	pointsOfFunctionPlot[1_403].X = 9.69
	pointsOfFunctionPlot[1_403].Y = 5_929_517.32

	pointsOfFunctionPlot[1_404].X = 9.7
	pointsOfFunctionPlot[1_404].Y = 6_025_721.315

	pointsOfFunctionPlot[1_405].X = 9.71
	pointsOfFunctionPlot[1_405].Y = 6_123_486.181

	pointsOfFunctionPlot[1_406].X = 9.72
	pointsOfFunctionPlot[1_406].Y = 6_222_837.241

	pointsOfFunctionPlot[1_407].X = 9.73
	pointsOfFunctionPlot[1_407].Y = 6_323_800.232

	pointsOfFunctionPlot[1_408].X = 9.74
	pointsOfFunctionPlot[1_408].Y = 6_426_401.306

	pointsOfFunctionPlot[1_409].X = 9.75
	pointsOfFunctionPlot[1_409].Y = 6_530_667.04

	pointsOfFunctionPlot[1_410].X = 9.76
	pointsOfFunctionPlot[1_410].Y = 6_636_624.444

	pointsOfFunctionPlot[1_411].X = 9.77
	pointsOfFunctionPlot[1_411].Y = 6_744_300.963

	pointsOfFunctionPlot[1_412].X = 9.78
	pointsOfFunctionPlot[1_412].Y = 6_853_724.489

	pointsOfFunctionPlot[1_413].X = 9.79
	pointsOfFunctionPlot[1_413].Y = 6_964_923.368

	pointsOfFunctionPlot[1_414].X = 9.8
	pointsOfFunctionPlot[1_414].Y = 7_077_926.403

	pointsOfFunctionPlot[1_415].X = 9.81
	pointsOfFunctionPlot[1_415].Y = 7_192_762.866

	pointsOfFunctionPlot[1_416].X = 9.82
	pointsOfFunctionPlot[1_416].Y = 7_309_462.503

	pointsOfFunctionPlot[1_417].X = 9.83
	pointsOfFunctionPlot[1_417].Y = 7_428_055.545

	pointsOfFunctionPlot[1_418].X = 9.84
	pointsOfFunctionPlot[1_418].Y = 7_548_572.71

	pointsOfFunctionPlot[1_419].X = 9.85
	pointsOfFunctionPlot[1_419].Y = 7_671_045.217

	pointsOfFunctionPlot[1_420].X = 9.86
	pointsOfFunctionPlot[1_420].Y = 7_795_504.79

	pointsOfFunctionPlot[1_421].X = 9.87
	pointsOfFunctionPlot[1_421].Y = 7_921_983.669

	pointsOfFunctionPlot[1_422].X = 9.88
	pointsOfFunctionPlot[1_422].Y = 8_050_514.616

	pointsOfFunctionPlot[1_423].X = 9.89
	pointsOfFunctionPlot[1_423].Y = 8_181_130.925

	pointsOfFunctionPlot[1_424].X = 9.9
	pointsOfFunctionPlot[1_424].Y = 8_313_866.43

	pointsOfFunctionPlot[1_425].X = 9.91
	pointsOfFunctionPlot[1_425].Y = 8_448_755.515

	pointsOfFunctionPlot[1_426].X = 9.92
	pointsOfFunctionPlot[1_426].Y = 8_585_833.12

	pointsOfFunctionPlot[1_427].X = 9.93
	pointsOfFunctionPlot[1_427].Y = 8_725_134.753

	pointsOfFunctionPlot[1_428].X = 9.94
	pointsOfFunctionPlot[1_428].Y = 8_866_696.498

	pointsOfFunctionPlot[1_429].X = 9.95
	pointsOfFunctionPlot[1_429].Y = 9_010_555.025

	pointsOfFunctionPlot[1_430].X = 9.96
	pointsOfFunctionPlot[1_430].Y = 9_156_747.597

	pointsOfFunctionPlot[1_431].X = 9.97
	pointsOfFunctionPlot[1_431].Y = 9_305_312.084

	pointsOfFunctionPlot[1_432].X = 9.98
	pointsOfFunctionPlot[1_432].Y = 9_456_286.97

	pointsOfFunctionPlot[1_433].X = 9.99
	pointsOfFunctionPlot[1_433].Y = 9_609_711.361

	pointsOfFunctionPlot[1_434].X = 10.0
	pointsOfFunctionPlot[1_434].Y = 9_765_625.0










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function 5^x"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("f(x)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"Power-of-5-function-plot-01.png"); err != nil {

		panic(err)
	}
}
