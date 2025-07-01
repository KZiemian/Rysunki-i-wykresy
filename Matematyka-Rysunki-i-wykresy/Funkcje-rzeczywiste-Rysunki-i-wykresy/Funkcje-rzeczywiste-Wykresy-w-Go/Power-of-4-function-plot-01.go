package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function 4^x

	pointsOfFunctionPlot := make(plotter.XYs, 1_504)

	pointsOfFunctionPlot[0].X = -10.0
	pointsOfFunctionPlot[0].Y = 0.0

	pointsOfFunctionPlot[1].X = -5.02
	pointsOfFunctionPlot[1].Y = 0.0

	pointsOfFunctionPlot[2].X = -5.01
	pointsOfFunctionPlot[2].Y = 0.001

	pointsOfFunctionPlot[3].X = -5.0
	pointsOfFunctionPlot[3].Y = 0.001

	pointsOfFunctionPlot[4].X = -4.99
	pointsOfFunctionPlot[4].Y = 0.001

	pointsOfFunctionPlot[5].X = -4.98
	pointsOfFunctionPlot[5].Y = 0.001

	pointsOfFunctionPlot[6].X = -4.97
	pointsOfFunctionPlot[6].Y = 0.001

	pointsOfFunctionPlot[7].X = -4.96
	pointsOfFunctionPlot[7].Y = 0.001

	pointsOfFunctionPlot[8].X = -4.95
	pointsOfFunctionPlot[8].Y = 0.001

	pointsOfFunctionPlot[9].X = -4.94
	pointsOfFunctionPlot[9].Y = 0.001

	pointsOfFunctionPlot[10].X = -4.93
	pointsOfFunctionPlot[10].Y = 0.001

	pointsOfFunctionPlot[11].X = -4.92
	pointsOfFunctionPlot[11].Y = 0.001

	pointsOfFunctionPlot[12].X = -4.91
	pointsOfFunctionPlot[12].Y = 0.001

	pointsOfFunctionPlot[13].X = -4.9
	pointsOfFunctionPlot[13].Y = 0.001

	pointsOfFunctionPlot[14].X = -4.89
	pointsOfFunctionPlot[14].Y = 0.001

	pointsOfFunctionPlot[15].X = -4.88
	pointsOfFunctionPlot[15].Y = 0.001

	pointsOfFunctionPlot[16].X = -4.87
	pointsOfFunctionPlot[16].Y = 0.001

	pointsOfFunctionPlot[17].X = -4.86
	pointsOfFunctionPlot[17].Y = 0.001

	pointsOfFunctionPlot[18].X = -4.85
	pointsOfFunctionPlot[18].Y = 0.001

	pointsOfFunctionPlot[19].X = -4.84
	pointsOfFunctionPlot[19].Y = 0.001

	pointsOfFunctionPlot[20].X = -4.83
	pointsOfFunctionPlot[20].Y = 0.001

	pointsOfFunctionPlot[21].X = -4.82
	pointsOfFunctionPlot[21].Y = 0.001

	pointsOfFunctionPlot[22].X = -4.81
	pointsOfFunctionPlot[22].Y = 0.001

	pointsOfFunctionPlot[23].X = -4.8
	pointsOfFunctionPlot[23].Y = 0.001

	pointsOfFunctionPlot[24].X = -4.79
	pointsOfFunctionPlot[24].Y = 0.001

	pointsOfFunctionPlot[25].X = -4.78
	pointsOfFunctionPlot[25].Y = 0.001

	pointsOfFunctionPlot[26].X = -4.77
	pointsOfFunctionPlot[26].Y = 0.001

	pointsOfFunctionPlot[27].X = -4.76
	pointsOfFunctionPlot[27].Y = 0.001

	pointsOfFunctionPlot[28].X = -4.75
	pointsOfFunctionPlot[28].Y = 0.001

	pointsOfFunctionPlot[29].X = -4.74
	pointsOfFunctionPlot[29].Y = 0.001

	pointsOfFunctionPlot[30].X = -4.73
	pointsOfFunctionPlot[30].Y = 0.001

	pointsOfFunctionPlot[31].X = -4.72
	pointsOfFunctionPlot[31].Y = 0.001

	pointsOfFunctionPlot[32].X = -4.71
	pointsOfFunctionPlot[32].Y = 0.001

	pointsOfFunctionPlot[33].X = -4.7
	pointsOfFunctionPlot[33].Y = 0.001

	pointsOfFunctionPlot[34].X = -4.69
	pointsOfFunctionPlot[34].Y = 0.001

	pointsOfFunctionPlot[35].X = -4.68
	pointsOfFunctionPlot[35].Y = 0.001

	pointsOfFunctionPlot[36].X = -4.67
	pointsOfFunctionPlot[36].Y = 0.001

	pointsOfFunctionPlot[37].X = -4.66
	pointsOfFunctionPlot[37].Y = 0.001

	pointsOfFunctionPlot[38].X = -4.65
	pointsOfFunctionPlot[38].Y = 0.001

	pointsOfFunctionPlot[39].X = -4.64
	pointsOfFunctionPlot[39].Y = 0.001

	pointsOfFunctionPlot[40].X = -4.63
	pointsOfFunctionPlot[40].Y = 0.001

	pointsOfFunctionPlot[41].X = -4.62
	pointsOfFunctionPlot[41].Y = 0.001

	pointsOfFunctionPlot[42].X = -4.61
	pointsOfFunctionPlot[42].Y = 0.001

	pointsOfFunctionPlot[43].X = -4.6
	pointsOfFunctionPlot[43].Y = 0.001

	pointsOfFunctionPlot[44].X = -4.59
	pointsOfFunctionPlot[44].Y = 0.001

	pointsOfFunctionPlot[45].X = -4.58
	pointsOfFunctionPlot[45].Y = 0.001

	pointsOfFunctionPlot[46].X = -4.57
	pointsOfFunctionPlot[46].Y = 0.001

	pointsOfFunctionPlot[47].X = -4.56
	pointsOfFunctionPlot[47].Y = 0.001

	pointsOfFunctionPlot[48].X = -4.55
	pointsOfFunctionPlot[48].Y = 0.001

	pointsOfFunctionPlot[49].X = -4.54
	pointsOfFunctionPlot[49].Y = 0.001

	pointsOfFunctionPlot[50].X = -4.53
	pointsOfFunctionPlot[50].Y = 0.001

	pointsOfFunctionPlot[51].X = -4.52
	pointsOfFunctionPlot[51].Y = 0.001

	pointsOfFunctionPlot[52].X = -4.51
	pointsOfFunctionPlot[52].Y = 0.001

	pointsOfFunctionPlot[53].X = -4.5
	pointsOfFunctionPlot[53].Y = 0.002

	pointsOfFunctionPlot[54].X = -4.49
	pointsOfFunctionPlot[54].Y = 0.002

	pointsOfFunctionPlot[55].X = -4.48
	pointsOfFunctionPlot[55].Y = 0.002

	pointsOfFunctionPlot[56].X = -4.47
	pointsOfFunctionPlot[56].Y = 0.002

	pointsOfFunctionPlot[57].X = -4.46
	pointsOfFunctionPlot[57].Y = 0.002

	pointsOfFunctionPlot[58].X = -4.45
	pointsOfFunctionPlot[58].Y = 0.002

	pointsOfFunctionPlot[59].X = -4.44
	pointsOfFunctionPlot[59].Y = 0.002

	pointsOfFunctionPlot[60].X = -4.43
	pointsOfFunctionPlot[60].Y = 0.002

	pointsOfFunctionPlot[61].X = -4.42
	pointsOfFunctionPlot[61].Y = 0.002

	pointsOfFunctionPlot[62].X = -4.41
	pointsOfFunctionPlot[62].Y = 0.002

	pointsOfFunctionPlot[63].X = -4.4
	pointsOfFunctionPlot[63].Y = 0.002

	pointsOfFunctionPlot[64].X = -4.39
	pointsOfFunctionPlot[64].Y = 0.002

	pointsOfFunctionPlot[65].X = -4.38
	pointsOfFunctionPlot[65].Y = 0.002

	pointsOfFunctionPlot[66].X = -4.37
	pointsOfFunctionPlot[66].Y = 0.002

	pointsOfFunctionPlot[67].X = -4.36
	pointsOfFunctionPlot[67].Y = 0.002

	pointsOfFunctionPlot[68].X = -4.35
	pointsOfFunctionPlot[68].Y = 0.002

	pointsOfFunctionPlot[69].X = -4.34
	pointsOfFunctionPlot[69].Y = 0.002

	pointsOfFunctionPlot[70].X = -4.33
	pointsOfFunctionPlot[70].Y = 0.002

	pointsOfFunctionPlot[71].X = -4.32
	pointsOfFunctionPlot[71].Y = 0.002

	pointsOfFunctionPlot[72].X = -4.31
	pointsOfFunctionPlot[72].Y = 0.002

	pointsOfFunctionPlot[73].X = -4.3
	pointsOfFunctionPlot[73].Y = 0.002

	pointsOfFunctionPlot[74].X = -4.29
	pointsOfFunctionPlot[74].Y = 0.002

	pointsOfFunctionPlot[75].X = -4.28
	pointsOfFunctionPlot[75].Y = 0.002

	pointsOfFunctionPlot[76].X = -4.27
	pointsOfFunctionPlot[76].Y = 0.002

	pointsOfFunctionPlot[77].X = -4.26
	pointsOfFunctionPlot[77].Y = 0.002

	pointsOfFunctionPlot[78].X = -4.25
	pointsOfFunctionPlot[78].Y = 0.002

	pointsOfFunctionPlot[79].X = -4.24
	pointsOfFunctionPlot[79].Y = 0.002

	pointsOfFunctionPlot[80].X = -4.23
	pointsOfFunctionPlot[80].Y = 0.002

	pointsOfFunctionPlot[81].X = -4.22
	pointsOfFunctionPlot[81].Y = 0.002

	pointsOfFunctionPlot[82].X = -4.21
	pointsOfFunctionPlot[82].Y = 0.002

	pointsOfFunctionPlot[83].X = -4.2
	pointsOfFunctionPlot[83].Y = 0.003

	pointsOfFunctionPlot[84].X = -4.19
	pointsOfFunctionPlot[84].Y = 0.003

	pointsOfFunctionPlot[85].X = -4.18
	pointsOfFunctionPlot[85].Y = 0.003

	pointsOfFunctionPlot[86].X = -4.17
	pointsOfFunctionPlot[86].Y = 0.003

	pointsOfFunctionPlot[87].X = -4.16
	pointsOfFunctionPlot[87].Y = 0.003

	pointsOfFunctionPlot[88].X = -4.15
	pointsOfFunctionPlot[88].Y = 0.003

	pointsOfFunctionPlot[89].X = -4.14
	pointsOfFunctionPlot[89].Y = 0.003

	pointsOfFunctionPlot[90].X = -4.13
	pointsOfFunctionPlot[90].Y = 0.003

	pointsOfFunctionPlot[91].X = -4.12
	pointsOfFunctionPlot[91].Y = 0.003

	pointsOfFunctionPlot[92].X = -4.11
	pointsOfFunctionPlot[92].Y = 0.003

	pointsOfFunctionPlot[93].X = -4.1
	pointsOfFunctionPlot[93].Y = 0.003

	pointsOfFunctionPlot[94].X = -4.09
	pointsOfFunctionPlot[94].Y = 0.003

	pointsOfFunctionPlot[95].X = -4.08
	pointsOfFunctionPlot[95].Y = 0.003

	pointsOfFunctionPlot[96].X = -4.07
	pointsOfFunctionPlot[96].Y = 0.003

	pointsOfFunctionPlot[97].X = -4.06
	pointsOfFunctionPlot[97].Y = 0.003

	pointsOfFunctionPlot[98].X = -4.05
	pointsOfFunctionPlot[98].Y = 0.003

	pointsOfFunctionPlot[99].X = -4.04
	pointsOfFunctionPlot[99].Y = 0.003

	pointsOfFunctionPlot[100].X = -4.03
	pointsOfFunctionPlot[100].Y = 0.003

	pointsOfFunctionPlot[101].X = -4.02
	pointsOfFunctionPlot[101].Y = 0.003

	pointsOfFunctionPlot[102].X = -4.01
	pointsOfFunctionPlot[102].Y = 0.003

	pointsOfFunctionPlot[103].X = -4.0
	pointsOfFunctionPlot[103].Y = 0.003

	pointsOfFunctionPlot[104].X = -3.99
	pointsOfFunctionPlot[104].Y = 0.004

	pointsOfFunctionPlot[105].X = -3.98
	pointsOfFunctionPlot[105].Y = 0.004

	pointsOfFunctionPlot[106].X = -3.97
	pointsOfFunctionPlot[106].Y = 0.004

	pointsOfFunctionPlot[107].X = -3.96
	pointsOfFunctionPlot[107].Y = 0.004

	pointsOfFunctionPlot[108].X = -3.95
	pointsOfFunctionPlot[108].Y = 0.004

	pointsOfFunctionPlot[109].X = -3.94
	pointsOfFunctionPlot[109].Y = 0.004

	pointsOfFunctionPlot[110].X = -3.93
	pointsOfFunctionPlot[110].Y = 0.004

	pointsOfFunctionPlot[111].X = -3.92
	pointsOfFunctionPlot[111].Y = 0.004

	pointsOfFunctionPlot[112].X = -3.91
	pointsOfFunctionPlot[112].Y = 0.004

	pointsOfFunctionPlot[113].X = -3.9
	pointsOfFunctionPlot[113].Y = 0.004

	pointsOfFunctionPlot[114].X = -3.89
	pointsOfFunctionPlot[114].Y = 0.004

	pointsOfFunctionPlot[115].X = -3.88
	pointsOfFunctionPlot[115].Y = 0.004

	pointsOfFunctionPlot[116].X = -3.87
	pointsOfFunctionPlot[116].Y = 0.004

	pointsOfFunctionPlot[117].X = -3.86
	pointsOfFunctionPlot[117].Y = 0.004

	pointsOfFunctionPlot[118].X = -3.85
	pointsOfFunctionPlot[118].Y = 0.004

	pointsOfFunctionPlot[119].X = -3.84
	pointsOfFunctionPlot[119].Y = 0.004

	pointsOfFunctionPlot[120].X = -3.83
	pointsOfFunctionPlot[120].Y = 0.004

	pointsOfFunctionPlot[121].X = -3.82
	pointsOfFunctionPlot[121].Y = 0.005

	pointsOfFunctionPlot[122].X = -3.81
	pointsOfFunctionPlot[122].Y = 0.005

	pointsOfFunctionPlot[123].X = -3.8
	pointsOfFunctionPlot[123].Y = 0.005

	pointsOfFunctionPlot[124].X = -3.79
	pointsOfFunctionPlot[124].Y = 0.005

	pointsOfFunctionPlot[125].X = -3.78
	pointsOfFunctionPlot[125].Y = 0.005

	pointsOfFunctionPlot[126].X = -3.77
	pointsOfFunctionPlot[126].Y = 0.005

	pointsOfFunctionPlot[127].X = -3.76
	pointsOfFunctionPlot[127].Y = 0.005

	pointsOfFunctionPlot[128].X = -3.75
	pointsOfFunctionPlot[128].Y = 0.005

	pointsOfFunctionPlot[129].X = -3.74
	pointsOfFunctionPlot[129].Y = 0.005

	pointsOfFunctionPlot[130].X = -3.73
	pointsOfFunctionPlot[130].Y = 0.005

	pointsOfFunctionPlot[131].X = -3.72
	pointsOfFunctionPlot[131].Y = 0.005

	pointsOfFunctionPlot[132].X = -3.71
	pointsOfFunctionPlot[132].Y = 0.005

	pointsOfFunctionPlot[133].X = -3.7
	pointsOfFunctionPlot[133].Y = 0.005

	pointsOfFunctionPlot[134].X = -3.69
	pointsOfFunctionPlot[134].Y = 0.006

	pointsOfFunctionPlot[135].X = -3.68
	pointsOfFunctionPlot[135].Y = 0.006

	pointsOfFunctionPlot[136].X = -3.67
	pointsOfFunctionPlot[136].Y = 0.006

	pointsOfFunctionPlot[137].X = -3.66
	pointsOfFunctionPlot[137].Y = 0.006

	pointsOfFunctionPlot[138].X = -3.65
	pointsOfFunctionPlot[138].Y = 0.006

	pointsOfFunctionPlot[139].X = -3.64
	pointsOfFunctionPlot[139].Y = 0.006

	pointsOfFunctionPlot[140].X = -3.63
	pointsOfFunctionPlot[140].Y = 0.006

	pointsOfFunctionPlot[141].X = -3.62
	pointsOfFunctionPlot[141].Y = 0.006

	pointsOfFunctionPlot[142].X = -3.61
	pointsOfFunctionPlot[142].Y = 0.006

	pointsOfFunctionPlot[143].X = -3.6
	pointsOfFunctionPlot[143].Y = 0.006

	pointsOfFunctionPlot[144].X = -3.59
	pointsOfFunctionPlot[144].Y = 0.006

	pointsOfFunctionPlot[145].X = -3.58
	pointsOfFunctionPlot[145].Y = 0.007

	pointsOfFunctionPlot[146].X = -3.57
	pointsOfFunctionPlot[146].Y = 0.007

	pointsOfFunctionPlot[147].X = -3.56
	pointsOfFunctionPlot[147].Y = 0.007

	pointsOfFunctionPlot[148].X = -3.55
	pointsOfFunctionPlot[148].Y = 0.007

	pointsOfFunctionPlot[149].X = -3.54
	pointsOfFunctionPlot[149].Y = 0.007

	pointsOfFunctionPlot[150].X = -3.53
	pointsOfFunctionPlot[150].Y = 0.007

	pointsOfFunctionPlot[151].X = -3.52
	pointsOfFunctionPlot[151].Y = 0.007

	pointsOfFunctionPlot[152].X = -3.51
	pointsOfFunctionPlot[152].Y = 0.007

	pointsOfFunctionPlot[153].X = -3.5
	pointsOfFunctionPlot[153].Y = 0.007

	pointsOfFunctionPlot[154].X = -3.49
	pointsOfFunctionPlot[154].Y = 0.007

	pointsOfFunctionPlot[155].X = -3.48
	pointsOfFunctionPlot[155].Y = 0.008

	pointsOfFunctionPlot[156].X = -3.47
	pointsOfFunctionPlot[156].Y = 0.008

	pointsOfFunctionPlot[157].X = -3.46
	pointsOfFunctionPlot[157].Y = 0.008

	pointsOfFunctionPlot[158].X = -3.45
	pointsOfFunctionPlot[158].Y = 0.008

	pointsOfFunctionPlot[159].X = -3.44
	pointsOfFunctionPlot[159].Y = 0.008

	pointsOfFunctionPlot[160].X = -3.43
	pointsOfFunctionPlot[160].Y = 0.008

	pointsOfFunctionPlot[161].X = -3.42
	pointsOfFunctionPlot[161].Y = 0.008

	pointsOfFunctionPlot[162].X = -3.41
	pointsOfFunctionPlot[162].Y = 0.008

	pointsOfFunctionPlot[163].X = -3.4
	pointsOfFunctionPlot[163].Y = 0.009

	pointsOfFunctionPlot[164].X = -3.39
	pointsOfFunctionPlot[164].Y = 0.009

	pointsOfFunctionPlot[165].X = -3.38
	pointsOfFunctionPlot[165].Y = 0.009

	pointsOfFunctionPlot[166].X = -3.37
	pointsOfFunctionPlot[166].Y = 0.009

	pointsOfFunctionPlot[167].X = -3.36
	pointsOfFunctionPlot[167].Y = 0.009

	pointsOfFunctionPlot[168].X = -3.35
	pointsOfFunctionPlot[168].Y = 0.009

	pointsOfFunctionPlot[169].X = -3.34
	pointsOfFunctionPlot[169].Y = 0.009

	pointsOfFunctionPlot[170].X = -3.33
	pointsOfFunctionPlot[170].Y = 0.009

	pointsOfFunctionPlot[171].X = -3.32
	pointsOfFunctionPlot[171].Y = 0.01

	pointsOfFunctionPlot[172].X = -3.31
	pointsOfFunctionPlot[172].Y = 0.01

	pointsOfFunctionPlot[173].X = -3.3
	pointsOfFunctionPlot[173].Y = 0.01

	pointsOfFunctionPlot[174].X = -3.29
	pointsOfFunctionPlot[174].Y = 0.01

	pointsOfFunctionPlot[175].X = -3.28
	pointsOfFunctionPlot[175].Y = 0.01

	pointsOfFunctionPlot[176].X = -3.27
	pointsOfFunctionPlot[176].Y = 0.01

	pointsOfFunctionPlot[177].X = -3.26
	pointsOfFunctionPlot[177].Y = 0.01

	pointsOfFunctionPlot[178].X = -3.25
	pointsOfFunctionPlot[178].Y = 0.011

	pointsOfFunctionPlot[179].X = -3.24
	pointsOfFunctionPlot[179].Y = 0.011

	pointsOfFunctionPlot[180].X = -3.23
	pointsOfFunctionPlot[180].Y = 0.011

	pointsOfFunctionPlot[181].X = -3.22
	pointsOfFunctionPlot[181].Y = 0.011

	pointsOfFunctionPlot[182].X = -3.21
	pointsOfFunctionPlot[182].Y = 0.011

	pointsOfFunctionPlot[183].X = -3.2
	pointsOfFunctionPlot[183].Y = 0.011

	pointsOfFunctionPlot[184].X = -3.19
	pointsOfFunctionPlot[184].Y = 0.012

	pointsOfFunctionPlot[185].X = -3.18
	pointsOfFunctionPlot[185].Y = 0.012

	pointsOfFunctionPlot[186].X = -3.17
	pointsOfFunctionPlot[186].Y = 0.012

	pointsOfFunctionPlot[187].X = -3.16
	pointsOfFunctionPlot[187].Y = 0.012

	pointsOfFunctionPlot[188].X = -3.15
	pointsOfFunctionPlot[188].Y = 0.012

	pointsOfFunctionPlot[189].X = -3.14
	pointsOfFunctionPlot[189].Y = 0.012

	pointsOfFunctionPlot[190].X = -3.13
	pointsOfFunctionPlot[190].Y = 0.013

	pointsOfFunctionPlot[191].X = -3.12
	pointsOfFunctionPlot[191].Y = 0.013

	pointsOfFunctionPlot[192].X = -3.11
	pointsOfFunctionPlot[192].Y = 0.013

	pointsOfFunctionPlot[193].X = -3.1
	pointsOfFunctionPlot[193].Y = 0.013

	pointsOfFunctionPlot[194].X = -3.09
	pointsOfFunctionPlot[194].Y = 0.013

	pointsOfFunctionPlot[195].X = -3.08
	pointsOfFunctionPlot[195].Y = 0.014

	pointsOfFunctionPlot[196].X = -3.07
	pointsOfFunctionPlot[196].Y = 0.014

	pointsOfFunctionPlot[197].X = -3.06
	pointsOfFunctionPlot[197].Y = 0.014

	pointsOfFunctionPlot[198].X = -3.05
	pointsOfFunctionPlot[198].Y = 0.014

	pointsOfFunctionPlot[199].X = -3.04
	pointsOfFunctionPlot[199].Y = 0.014

	pointsOfFunctionPlot[200].X = -3.03
	pointsOfFunctionPlot[200].Y = 0.015

	pointsOfFunctionPlot[201].X = -3.02
	pointsOfFunctionPlot[201].Y = 0.015

	pointsOfFunctionPlot[202].X = -3.01
	pointsOfFunctionPlot[202].Y = 0.015

	pointsOfFunctionPlot[203].X = -3.0
	pointsOfFunctionPlot[203].Y = 0.015

	pointsOfFunctionPlot[204].X = -2.99
	pointsOfFunctionPlot[204].Y = 0.015

	pointsOfFunctionPlot[205].X = -2.98
	pointsOfFunctionPlot[205].Y = 0.016

	pointsOfFunctionPlot[206].X = -2.97
	pointsOfFunctionPlot[206].Y = 0.016

	pointsOfFunctionPlot[207].X = -2.96
	pointsOfFunctionPlot[207].Y = 0.016

	pointsOfFunctionPlot[208].X = -2.95
	pointsOfFunctionPlot[208].Y = 0.016

	pointsOfFunctionPlot[209].X = -2.94
	pointsOfFunctionPlot[209].Y = 0.017

	pointsOfFunctionPlot[210].X = -2.93
	pointsOfFunctionPlot[210].Y = 0.017

	pointsOfFunctionPlot[211].X = -2.92
	pointsOfFunctionPlot[211].Y = 0.017

	pointsOfFunctionPlot[212].X = -2.91
	pointsOfFunctionPlot[212].Y = 0.017

	pointsOfFunctionPlot[213].X = -2.9
	pointsOfFunctionPlot[213].Y = 0.017

	pointsOfFunctionPlot[214].X = -2.89
	pointsOfFunctionPlot[214].Y = 0.018

	pointsOfFunctionPlot[215].X = -2.88
	pointsOfFunctionPlot[215].Y = 0.018

	pointsOfFunctionPlot[216].X = -2.87
	pointsOfFunctionPlot[216].Y = 0.018

	pointsOfFunctionPlot[217].X = -2.86
	pointsOfFunctionPlot[217].Y = 0.019

	pointsOfFunctionPlot[218].X = -2.85
	pointsOfFunctionPlot[218].Y = 0.019

	pointsOfFunctionPlot[219].X = -2.84
	pointsOfFunctionPlot[219].Y = 0.019

	pointsOfFunctionPlot[220].X = -2.83
	pointsOfFunctionPlot[220].Y = 0.019

	pointsOfFunctionPlot[221].X = -2.82
	pointsOfFunctionPlot[221].Y = 0.02

	pointsOfFunctionPlot[222].X = -2.81
	pointsOfFunctionPlot[222].Y = 0.02

	pointsOfFunctionPlot[223].X = -2.8
	pointsOfFunctionPlot[223].Y = 0.02

	pointsOfFunctionPlot[224].X = -2.79
	pointsOfFunctionPlot[224].Y = 0.02

	pointsOfFunctionPlot[225].X = -2.78
	pointsOfFunctionPlot[225].Y = 0.021

	pointsOfFunctionPlot[226].X = -2.77
	pointsOfFunctionPlot[226].Y = 0.021

	pointsOfFunctionPlot[227].X = -2.76
	pointsOfFunctionPlot[227].Y = 0.021

	pointsOfFunctionPlot[228].X = -2.75
	pointsOfFunctionPlot[228].Y = 0.022

	pointsOfFunctionPlot[229].X = -2.74
	pointsOfFunctionPlot[229].Y = 0.022

	pointsOfFunctionPlot[230].X = -2.73
	pointsOfFunctionPlot[230].Y = 0.022

	pointsOfFunctionPlot[231].X = -2.72
	pointsOfFunctionPlot[231].Y = 0.023

	pointsOfFunctionPlot[232].X = -2.71
	pointsOfFunctionPlot[232].Y = 0.023

	pointsOfFunctionPlot[233].X = -2.7
	pointsOfFunctionPlot[233].Y = 0.023

	pointsOfFunctionPlot[234].X = -2.69
	pointsOfFunctionPlot[234].Y = 0.024

	pointsOfFunctionPlot[235].X = -2.68
	pointsOfFunctionPlot[235].Y = 0.024

	pointsOfFunctionPlot[236].X = -2.67
	pointsOfFunctionPlot[236].Y = 0.024

	pointsOfFunctionPlot[237].X = -2.66
	pointsOfFunctionPlot[237].Y = 0.025

	pointsOfFunctionPlot[238].X = -2.65
	pointsOfFunctionPlot[238].Y = 0.025

	pointsOfFunctionPlot[239].X = -2.64
	pointsOfFunctionPlot[239].Y = 0.025

	pointsOfFunctionPlot[240].X = -2.63
	pointsOfFunctionPlot[240].Y = 0.026

	pointsOfFunctionPlot[241].X = -2.62
	pointsOfFunctionPlot[241].Y = 0.026

	pointsOfFunctionPlot[242].X = -2.61
	pointsOfFunctionPlot[242].Y = 0.026

	pointsOfFunctionPlot[243].X = -2.6
	pointsOfFunctionPlot[243].Y = 0.027

	pointsOfFunctionPlot[244].X = -2.59
	pointsOfFunctionPlot[244].Y = 0.027

	pointsOfFunctionPlot[245].X = -2.58
	pointsOfFunctionPlot[245].Y = 0.028

	pointsOfFunctionPlot[246].X = -2.57
	pointsOfFunctionPlot[246].Y = 0.028

	pointsOfFunctionPlot[247].X = -2.56
	pointsOfFunctionPlot[247].Y = 0.028

	pointsOfFunctionPlot[248].X = -2.55
	pointsOfFunctionPlot[248].Y = 0.029

	pointsOfFunctionPlot[249].X = -2.54
	pointsOfFunctionPlot[249].Y = 0.029

	pointsOfFunctionPlot[250].X = -2.53
	pointsOfFunctionPlot[250].Y = 0.03

	pointsOfFunctionPlot[251].X = -2.52
	pointsOfFunctionPlot[251].Y = 0.03

	pointsOfFunctionPlot[252].X = -2.51
	pointsOfFunctionPlot[252].Y = 0.03

	pointsOfFunctionPlot[253].X = -2.5
	pointsOfFunctionPlot[253].Y = 0.031

	pointsOfFunctionPlot[254].X = -2.49
	pointsOfFunctionPlot[254].Y = 0.031

	pointsOfFunctionPlot[255].X = -2.48
	pointsOfFunctionPlot[255].Y = 0.031

	pointsOfFunctionPlot[256].X = -2.47
	pointsOfFunctionPlot[256].Y = 0.032

	pointsOfFunctionPlot[257].X = -2.46
	pointsOfFunctionPlot[257].Y = 0.033

	pointsOfFunctionPlot[258].X = -2.45
	pointsOfFunctionPlot[258].Y = 0.033

	pointsOfFunctionPlot[259].X = -2.44
	pointsOfFunctionPlot[259].Y = 0.034

	pointsOfFunctionPlot[260].X = -2.43
	pointsOfFunctionPlot[260].Y = 0.034

	pointsOfFunctionPlot[261].X = -2.42
	pointsOfFunctionPlot[261].Y = 0.034

	pointsOfFunctionPlot[262].X = -2.41
	pointsOfFunctionPlot[262].Y = 0.035

	pointsOfFunctionPlot[263].X = -2.4
	pointsOfFunctionPlot[263].Y = 0.035

	pointsOfFunctionPlot[264].X = -2.39
	pointsOfFunctionPlot[264].Y = 0.036

	pointsOfFunctionPlot[265].X = -2.38
	pointsOfFunctionPlot[265].Y = 0.036

	pointsOfFunctionPlot[266].X = -2.37
	pointsOfFunctionPlot[266].Y = 0.037

	pointsOfFunctionPlot[267].X = -2.36
	pointsOfFunctionPlot[267].Y = 0.037

	pointsOfFunctionPlot[268].X = -2.35
	pointsOfFunctionPlot[268].Y = 0.038

	pointsOfFunctionPlot[269].X = -2.34
	pointsOfFunctionPlot[269].Y = 0.039

	pointsOfFunctionPlot[270].X = -2.33
	pointsOfFunctionPlot[270].Y = 0.039

	pointsOfFunctionPlot[271].X = -2.32
	pointsOfFunctionPlot[271].Y = 0.04

	pointsOfFunctionPlot[272].X = -2.31
	pointsOfFunctionPlot[272].Y = 0.04

	pointsOfFunctionPlot[273].X = -2.3
	pointsOfFunctionPlot[273].Y = 0.041

	pointsOfFunctionPlot[274].X = -2.29
	pointsOfFunctionPlot[274].Y = 0.041

	pointsOfFunctionPlot[275].X = -2.28
	pointsOfFunctionPlot[275].Y = 0.042

	pointsOfFunctionPlot[276].X = -2.27
	pointsOfFunctionPlot[276].Y = 0.043

	pointsOfFunctionPlot[277].X = -2.26
	pointsOfFunctionPlot[277].Y = 0.043

	pointsOfFunctionPlot[278].X = -2.25
	pointsOfFunctionPlot[278].Y = 0.044

	pointsOfFunctionPlot[279].X = -2.24
	pointsOfFunctionPlot[279].Y = 0.044

	pointsOfFunctionPlot[280].X = -2.23
	pointsOfFunctionPlot[280].Y = 0.045

	pointsOfFunctionPlot[281].X = -2.22
	pointsOfFunctionPlot[281].Y = 0.046

	pointsOfFunctionPlot[282].X = -2.21
	pointsOfFunctionPlot[282].Y = 0.046

	pointsOfFunctionPlot[283].X = -2.2
	pointsOfFunctionPlot[283].Y = 0.047

	pointsOfFunctionPlot[284].X = -2.19
	pointsOfFunctionPlot[284].Y = 0.048

	pointsOfFunctionPlot[285].X = -2.18
	pointsOfFunctionPlot[285].Y = 0.048

	pointsOfFunctionPlot[286].X = -2.17
	pointsOfFunctionPlot[286].Y = 0.049

	pointsOfFunctionPlot[287].X = -2.16
	pointsOfFunctionPlot[287].Y = 0.05

	pointsOfFunctionPlot[288].X = -2.15
	pointsOfFunctionPlot[288].Y = 0.05

	pointsOfFunctionPlot[289].X = -2.14
	pointsOfFunctionPlot[289].Y = 0.051

	pointsOfFunctionPlot[290].X = -2.13
	pointsOfFunctionPlot[290].Y = 0.052

	pointsOfFunctionPlot[291].X = -2.12
	pointsOfFunctionPlot[291].Y = 0.052

	pointsOfFunctionPlot[292].X = -2.11
	pointsOfFunctionPlot[292].Y = 0.053

	pointsOfFunctionPlot[293].X = -2.1
	pointsOfFunctionPlot[293].Y = 0.054

	pointsOfFunctionPlot[294].X = -2.09
	pointsOfFunctionPlot[294].Y = 0.055

	pointsOfFunctionPlot[295].X = -2.08
	pointsOfFunctionPlot[295].Y = 0.055

	pointsOfFunctionPlot[296].X = -2.07
	pointsOfFunctionPlot[296].Y = 0.056

	pointsOfFunctionPlot[297].X = -2.06
	pointsOfFunctionPlot[297].Y = 0.057

	pointsOfFunctionPlot[298].X = -2.05
	pointsOfFunctionPlot[298].Y = 0.058

	pointsOfFunctionPlot[299].X = -2.04
	pointsOfFunctionPlot[299].Y = 0.059

	pointsOfFunctionPlot[300].X = -2.03
	pointsOfFunctionPlot[300].Y = 0.06

	pointsOfFunctionPlot[301].X = -2.02
	pointsOfFunctionPlot[301].Y = 0.06

	pointsOfFunctionPlot[302].X = -2.01
	pointsOfFunctionPlot[302].Y = 0.061

	pointsOfFunctionPlot[303].X = -2.0
	pointsOfFunctionPlot[303].Y = 0.062

	pointsOfFunctionPlot[304].X = -1.99
	pointsOfFunctionPlot[304].Y = 0.062

	pointsOfFunctionPlot[305].X = -1.98
	pointsOfFunctionPlot[305].Y = 0.063

	pointsOfFunctionPlot[306].X = -1.97
	pointsOfFunctionPlot[306].Y = 0.064

	pointsOfFunctionPlot[307].X = -1.96
	pointsOfFunctionPlot[307].Y = 0.065

	pointsOfFunctionPlot[308].X = -1.95
	pointsOfFunctionPlot[308].Y = 0.067

	pointsOfFunctionPlot[309].X = -1.94
	pointsOfFunctionPlot[309].Y = 0.067

	pointsOfFunctionPlot[310].X = -1.93
	pointsOfFunctionPlot[310].Y = 0.068

	pointsOfFunctionPlot[311].X = -1.92
	pointsOfFunctionPlot[311].Y = 0.069

	pointsOfFunctionPlot[312].X = -1.91
	pointsOfFunctionPlot[312].Y = 0.07

	pointsOfFunctionPlot[313].X = -1.9
	pointsOfFunctionPlot[313].Y = 0.071

	pointsOfFunctionPlot[314].X = -1.89
	pointsOfFunctionPlot[314].Y = 0.072

	pointsOfFunctionPlot[315].X = -1.88
	pointsOfFunctionPlot[315].Y = 0.073

	pointsOfFunctionPlot[316].X = -1.87
	pointsOfFunctionPlot[316].Y = 0.074

	pointsOfFunctionPlot[317].X = -1.86
	pointsOfFunctionPlot[317].Y = 0.075

	pointsOfFunctionPlot[318].X = -1.85
	pointsOfFunctionPlot[318].Y = 0.076

	pointsOfFunctionPlot[319].X = -1.84
	pointsOfFunctionPlot[319].Y = 0.078

	pointsOfFunctionPlot[320].X = -1.83
	pointsOfFunctionPlot[320].Y = 0.079

	pointsOfFunctionPlot[321].X = -1.82
	pointsOfFunctionPlot[321].Y = 0.08

	pointsOfFunctionPlot[322].X = -1.81
	pointsOfFunctionPlot[322].Y = 0.081

	pointsOfFunctionPlot[323].X = -1.8
	pointsOfFunctionPlot[323].Y = 0.082

	pointsOfFunctionPlot[324].X = -1.79
	pointsOfFunctionPlot[324].Y = 0.083

	pointsOfFunctionPlot[325].X = -1.78
	pointsOfFunctionPlot[325].Y = 0.084

	pointsOfFunctionPlot[326].X = -1.77
	pointsOfFunctionPlot[326].Y = 0.086

	pointsOfFunctionPlot[327].X = -1.76
	pointsOfFunctionPlot[327].Y = 0.087

	pointsOfFunctionPlot[328].X = -1.75
	pointsOfFunctionPlot[328].Y = 0.088

	pointsOfFunctionPlot[329].X = -1.74
	pointsOfFunctionPlot[329].Y = 0.089

	pointsOfFunctionPlot[330].X = -1.73
	pointsOfFunctionPlot[330].Y = 0.09

	pointsOfFunctionPlot[331].X = -1.72
	pointsOfFunctionPlot[331].Y = 0.092

	pointsOfFunctionPlot[332].X = -1.71
	pointsOfFunctionPlot[332].Y = 0.093

	pointsOfFunctionPlot[333].X = -1.7
	pointsOfFunctionPlot[333].Y = 0.094

	pointsOfFunctionPlot[334].X = -1.69
	pointsOfFunctionPlot[334].Y = 0.096

	pointsOfFunctionPlot[335].X = -1.68
	pointsOfFunctionPlot[335].Y = 0.097

	pointsOfFunctionPlot[336].X = -1.67
	pointsOfFunctionPlot[336].Y = 0.098

	pointsOfFunctionPlot[337].X = -1.66
	pointsOfFunctionPlot[337].Y = 0.1

	pointsOfFunctionPlot[338].X = -1.65
	pointsOfFunctionPlot[338].Y = 0.101

	pointsOfFunctionPlot[339].X = -1.64
	pointsOfFunctionPlot[339].Y = 0.102

	pointsOfFunctionPlot[340].X = -1.63
	pointsOfFunctionPlot[340].Y = 0.104

	pointsOfFunctionPlot[341].X = -1.62
	pointsOfFunctionPlot[341].Y = 0.105

	pointsOfFunctionPlot[342].X = -1.61
	pointsOfFunctionPlot[342].Y = 0.107

	pointsOfFunctionPlot[343].X = -1.6
	pointsOfFunctionPlot[343].Y = 0.108

	pointsOfFunctionPlot[344].X = -1.59
	pointsOfFunctionPlot[344].Y = 0.11

	pointsOfFunctionPlot[345].X = -1.58
	pointsOfFunctionPlot[345].Y = 0.111

	pointsOfFunctionPlot[346].X = -1.57
	pointsOfFunctionPlot[346].Y = 0.113

	pointsOfFunctionPlot[347].X = -1.56
	pointsOfFunctionPlot[347].Y = 0.115

	pointsOfFunctionPlot[348].X = -1.55
	pointsOfFunctionPlot[348].Y = 0.116

	pointsOfFunctionPlot[349].X = -1.54
	pointsOfFunctionPlot[349].Y = 0.118

	pointsOfFunctionPlot[350].X = -1.53
	pointsOfFunctionPlot[350].Y = 0.119

	pointsOfFunctionPlot[351].X = -1.52
	pointsOfFunctionPlot[351].Y = 0.121

	pointsOfFunctionPlot[352].X = -1.51
	pointsOfFunctionPlot[352].Y = 0.123

	pointsOfFunctionPlot[353].X = -1.5
	pointsOfFunctionPlot[353].Y = 0.125

	pointsOfFunctionPlot[354].X = -1.49
	pointsOfFunctionPlot[354].Y = 0.126

	pointsOfFunctionPlot[355].X = -1.48
	pointsOfFunctionPlot[355].Y = 0.128

	pointsOfFunctionPlot[356].X = -1.47
	pointsOfFunctionPlot[356].Y = 0.13

	pointsOfFunctionPlot[357].X = -1.46
	pointsOfFunctionPlot[357].Y = 0.132

	pointsOfFunctionPlot[358].X = -1.45
	pointsOfFunctionPlot[358].Y = 0.134

	pointsOfFunctionPlot[359].X = -1.44
	pointsOfFunctionPlot[359].Y = 0.135

	pointsOfFunctionPlot[360].X = -1.43
	pointsOfFunctionPlot[360].Y = 0.137

	pointsOfFunctionPlot[361].X = -1.42
	pointsOfFunctionPlot[361].Y = 0.139

	pointsOfFunctionPlot[362].X = -1.41
	pointsOfFunctionPlot[362].Y = 0.141

	pointsOfFunctionPlot[363].X = -1.4
	pointsOfFunctionPlot[363].Y = 0.143

	pointsOfFunctionPlot[364].X = -1.39
	pointsOfFunctionPlot[364].Y = 0.145

	pointsOfFunctionPlot[365].X = -1.38
	pointsOfFunctionPlot[365].Y = 0.147

	pointsOfFunctionPlot[366].X = -1.37
	pointsOfFunctionPlot[366].Y = 0.149

	pointsOfFunctionPlot[367].X = -1.36
	pointsOfFunctionPlot[367].Y = 0.151

	pointsOfFunctionPlot[368].X = -1.35
	pointsOfFunctionPlot[368].Y = 0.153

	pointsOfFunctionPlot[369].X = -1.34
	pointsOfFunctionPlot[369].Y = 0.156

	pointsOfFunctionPlot[370].X = -1.33
	pointsOfFunctionPlot[370].Y = 0.158

	pointsOfFunctionPlot[371].X = -1.32
	pointsOfFunctionPlot[371].Y = 0.16

	pointsOfFunctionPlot[372].X = -1.31
	pointsOfFunctionPlot[372].Y = 0.162

	pointsOfFunctionPlot[373].X = -1.3
	pointsOfFunctionPlot[373].Y = 0.164

	pointsOfFunctionPlot[374].X = -1.29
	pointsOfFunctionPlot[374].Y = 0.167

	pointsOfFunctionPlot[375].X = -1.28
	pointsOfFunctionPlot[375].Y = 0.169

	pointsOfFunctionPlot[376].X = -1.27
	pointsOfFunctionPlot[376].Y = 0.171

	pointsOfFunctionPlot[377].X = -1.26
	pointsOfFunctionPlot[377].Y = 0.174

	pointsOfFunctionPlot[378].X = -1.25
	pointsOfFunctionPlot[378].Y = 0.176

	pointsOfFunctionPlot[379].X = -1.24
	pointsOfFunctionPlot[379].Y = 0.179

	pointsOfFunctionPlot[380].X = -1.23
	pointsOfFunctionPlot[380].Y = 0.181

	pointsOfFunctionPlot[381].X = -1.22
	pointsOfFunctionPlot[381].Y = 0.184

	pointsOfFunctionPlot[382].X = -1.21
	pointsOfFunctionPlot[382].Y = 0.186

	pointsOfFunctionPlot[383].X = -1.2
	pointsOfFunctionPlot[383].Y = 0.189

	pointsOfFunctionPlot[384].X = -1.19
	pointsOfFunctionPlot[384].Y = 0.192

	pointsOfFunctionPlot[385].X = -1.18
	pointsOfFunctionPlot[385].Y = 0.194

	pointsOfFunctionPlot[386].X = -1.17
	pointsOfFunctionPlot[386].Y = 0.197

	pointsOfFunctionPlot[387].X = -1.16
	pointsOfFunctionPlot[387].Y = 0.2

	pointsOfFunctionPlot[388].X = -1.15
	pointsOfFunctionPlot[388].Y = 0.203

	pointsOfFunctionPlot[389].X = -1.14
	pointsOfFunctionPlot[389].Y = 0.205

	pointsOfFunctionPlot[390].X = -1.13
	pointsOfFunctionPlot[390].Y = 0.208

	pointsOfFunctionPlot[391].X = -1.12
	pointsOfFunctionPlot[391].Y = 0.211

	pointsOfFunctionPlot[392].X = -1.11
	pointsOfFunctionPlot[392].Y = 0.214

	pointsOfFunctionPlot[393].X = -1.1
	pointsOfFunctionPlot[393].Y = 0.217

	pointsOfFunctionPlot[394].X = -1.09
	pointsOfFunctionPlot[394].Y = 0.22

	pointsOfFunctionPlot[395].X = -1.08
	pointsOfFunctionPlot[395].Y = 0.223

	pointsOfFunctionPlot[396].X = -1.07
	pointsOfFunctionPlot[396].Y = 0.226

	pointsOfFunctionPlot[397].X = -1.06
	pointsOfFunctionPlot[397].Y = 0.23

	pointsOfFunctionPlot[398].X = -1.05
	pointsOfFunctionPlot[398].Y = 0.233

	pointsOfFunctionPlot[399].X = -1.04
	pointsOfFunctionPlot[399].Y = 0.236

	pointsOfFunctionPlot[400].X = -1.03
	pointsOfFunctionPlot[400].Y = 0.239

	pointsOfFunctionPlot[401].X = -1.02
	pointsOfFunctionPlot[401].Y = 0.243

	pointsOfFunctionPlot[402].X = -1.01
	pointsOfFunctionPlot[402].Y = 0.246

	pointsOfFunctionPlot[403].X = -1.0
	pointsOfFunctionPlot[403].Y = 0.25

	pointsOfFunctionPlot[404].X = -0.99
	pointsOfFunctionPlot[404].Y = 0.253

	pointsOfFunctionPlot[405].X = -0.98
	pointsOfFunctionPlot[405].Y = 0.257

	pointsOfFunctionPlot[406].X = -0.97
	pointsOfFunctionPlot[406].Y = 0.26

	pointsOfFunctionPlot[407].X = -0.96
	pointsOfFunctionPlot[407].Y = 0.264

	pointsOfFunctionPlot[408].X = -0.95
	pointsOfFunctionPlot[408].Y = 0.267

	pointsOfFunctionPlot[409].X = -0.94
	pointsOfFunctionPlot[409].Y = 0.271

	pointsOfFunctionPlot[410].X = -0.93
	pointsOfFunctionPlot[410].Y = 0.275

	pointsOfFunctionPlot[411].X = -0.92
	pointsOfFunctionPlot[411].Y = 0.279

	pointsOfFunctionPlot[412].X = -0.91
	pointsOfFunctionPlot[412].Y = 0.283

	pointsOfFunctionPlot[413].X = -0.9
	pointsOfFunctionPlot[413].Y = 0.287

	pointsOfFunctionPlot[414].X = -0.89
	pointsOfFunctionPlot[414].Y = 0.291

	pointsOfFunctionPlot[415].X = -0.88
	pointsOfFunctionPlot[415].Y = 0.295

	pointsOfFunctionPlot[416].X = -0.87
	pointsOfFunctionPlot[416].Y = 0.299

	pointsOfFunctionPlot[417].X = -0.86
	pointsOfFunctionPlot[417].Y = 0.303

	pointsOfFunctionPlot[418].X = -0.85
	pointsOfFunctionPlot[418].Y = 0.307

	pointsOfFunctionPlot[419].X = -0.84
	pointsOfFunctionPlot[419].Y = 0.312

	pointsOfFunctionPlot[420].X = -0.83
	pointsOfFunctionPlot[420].Y = 0.316

	pointsOfFunctionPlot[421].X = -0.82
	pointsOfFunctionPlot[421].Y = 0.32

	pointsOfFunctionPlot[422].X = -0.81
	pointsOfFunctionPlot[422].Y = 0.325

	pointsOfFunctionPlot[423].X = -0.8
	pointsOfFunctionPlot[423].Y = 0.329

	pointsOfFunctionPlot[424].X = -0.79
	pointsOfFunctionPlot[424].Y = 0.334

	pointsOfFunctionPlot[425].X = -0.78
	pointsOfFunctionPlot[425].Y = 0.339

	pointsOfFunctionPlot[426].X = -0.77
	pointsOfFunctionPlot[426].Y = 0.343

	pointsOfFunctionPlot[427].X = -0.76
	pointsOfFunctionPlot[427].Y = 0.348

	pointsOfFunctionPlot[428].X = -0.75
	pointsOfFunctionPlot[428].Y = 0.353

	pointsOfFunctionPlot[429].X = -0.74
	pointsOfFunctionPlot[429].Y = 0.358

	pointsOfFunctionPlot[430].X = -0.73
	pointsOfFunctionPlot[430].Y = 0.363

	pointsOfFunctionPlot[431].X = -0.72
	pointsOfFunctionPlot[431].Y = 0.368

	pointsOfFunctionPlot[432].X = -0.71
	pointsOfFunctionPlot[432].Y = 0.373

	pointsOfFunctionPlot[433].X = -0.7
	pointsOfFunctionPlot[433].Y = 0.378

	pointsOfFunctionPlot[434].X = -0.69
	pointsOfFunctionPlot[434].Y = 0.384

	pointsOfFunctionPlot[435].X = -0.68
	pointsOfFunctionPlot[435].Y = 0.389

	pointsOfFunctionPlot[436].X = -0.67
	pointsOfFunctionPlot[436].Y = 0.395

	pointsOfFunctionPlot[437].X = -0.66
	pointsOfFunctionPlot[437].Y = 0.4

	pointsOfFunctionPlot[438].X = -0.65
	pointsOfFunctionPlot[438].Y = 0.406

	pointsOfFunctionPlot[439].X = -0.64
	pointsOfFunctionPlot[439].Y = 0.411

	pointsOfFunctionPlot[440].X = -0.63
	pointsOfFunctionPlot[440].Y = 0.417

	pointsOfFunctionPlot[441].X = -0.62
	pointsOfFunctionPlot[441].Y = 0.423

	pointsOfFunctionPlot[442].X = -0.61
	pointsOfFunctionPlot[442].Y = 0.429

	pointsOfFunctionPlot[443].X = -0.6
	pointsOfFunctionPlot[443].Y = 0.435

	pointsOfFunctionPlot[444].X = -0.59
	pointsOfFunctionPlot[444].Y = 0.441

	pointsOfFunctionPlot[445].X = -0.58
	pointsOfFunctionPlot[445].Y = 0.447

	pointsOfFunctionPlot[446].X = -0.57
	pointsOfFunctionPlot[446].Y = 0.453

	pointsOfFunctionPlot[447].X = -0.56
	pointsOfFunctionPlot[447].Y = 0.46

	pointsOfFunctionPlot[448].X = -0.55
	pointsOfFunctionPlot[448].Y = 0.466

	pointsOfFunctionPlot[449].X = -0.54
	pointsOfFunctionPlot[449].Y = 0.473

	pointsOfFunctionPlot[450].X = -0.53
	pointsOfFunctionPlot[450].Y = 0.479

	pointsOfFunctionPlot[451].X = -0.52
	pointsOfFunctionPlot[451].Y = 0.486

	pointsOfFunctionPlot[452].X = -0.51
	pointsOfFunctionPlot[452].Y = 0.493

	pointsOfFunctionPlot[453].X = -0.5
	pointsOfFunctionPlot[453].Y = 0.5

	pointsOfFunctionPlot[454].X = -0.49
	pointsOfFunctionPlot[454].Y = 0.507

	pointsOfFunctionPlot[455].X = -0.48
	pointsOfFunctionPlot[455].Y = 0.514

	pointsOfFunctionPlot[456].X = -0.47
	pointsOfFunctionPlot[456].Y = 0.521

	pointsOfFunctionPlot[457].X = -0.46
	pointsOfFunctionPlot[457].Y = 0.528

	pointsOfFunctionPlot[458].X = -0.45
	pointsOfFunctionPlot[458].Y = 0.535

	pointsOfFunctionPlot[459].X = -0.44
	pointsOfFunctionPlot[459].Y = 0.543

	pointsOfFunctionPlot[460].X = -0.43
	pointsOfFunctionPlot[460].Y = 0.551

	pointsOfFunctionPlot[461].X = -0.42
	pointsOfFunctionPlot[461].Y = 0.558

	pointsOfFunctionPlot[462].X = -0.41
	pointsOfFunctionPlot[462].Y = 0.566

	pointsOfFunctionPlot[463].X = -0.4
	pointsOfFunctionPlot[463].Y = 0.574

	pointsOfFunctionPlot[464].X = -0.39
	pointsOfFunctionPlot[464].Y = 0.582

	pointsOfFunctionPlot[465].X = -0.38
	pointsOfFunctionPlot[465].Y = 0.59

	pointsOfFunctionPlot[466].X = -0.37
	pointsOfFunctionPlot[466].Y = 0.598

	pointsOfFunctionPlot[467].X = -0.36
	pointsOfFunctionPlot[467].Y = 0.607

	pointsOfFunctionPlot[468].X = -0.35
	pointsOfFunctionPlot[468].Y = 0.615

	pointsOfFunctionPlot[469].X = -0.34
	pointsOfFunctionPlot[469].Y = 0.624

	pointsOfFunctionPlot[470].X = -0.33
	pointsOfFunctionPlot[470].Y = 0.632

	pointsOfFunctionPlot[471].X = -0.32
	pointsOfFunctionPlot[471].Y = 0.641

	pointsOfFunctionPlot[472].X = -0.31
	pointsOfFunctionPlot[472].Y = 0.65

	pointsOfFunctionPlot[473].X = -0.3
	pointsOfFunctionPlot[473].Y = 0.659

	pointsOfFunctionPlot[474].X = -0.29
	pointsOfFunctionPlot[474].Y = 0.669

	pointsOfFunctionPlot[475].X = -0.28
	pointsOfFunctionPlot[475].Y = 0.678

	pointsOfFunctionPlot[476].X = -0.27
	pointsOfFunctionPlot[476].Y = 0.687

	pointsOfFunctionPlot[477].X = -0.26
	pointsOfFunctionPlot[477].Y = 0.697

	pointsOfFunctionPlot[478].X = -0.25
	pointsOfFunctionPlot[478].Y = 0.707

	pointsOfFunctionPlot[479].X = -0.24
	pointsOfFunctionPlot[479].Y = 0.717

	pointsOfFunctionPlot[480].X = -0.23
	pointsOfFunctionPlot[480].Y = 0.727

	pointsOfFunctionPlot[481].X = -0.22
	pointsOfFunctionPlot[481].Y = 0.737

	pointsOfFunctionPlot[482].X = -0.21
	pointsOfFunctionPlot[482].Y = 0.747

	pointsOfFunctionPlot[483].X = -0.2
	pointsOfFunctionPlot[483].Y = 0.757

	pointsOfFunctionPlot[484].X = -0.19
	pointsOfFunctionPlot[484].Y = 0.768

	pointsOfFunctionPlot[485].X = -0.18
	pointsOfFunctionPlot[485].Y = 0.779

	pointsOfFunctionPlot[486].X = -0.17
	pointsOfFunctionPlot[486].Y = 0.79

	pointsOfFunctionPlot[487].X = -0.16
	pointsOfFunctionPlot[487].Y = 0.801

	pointsOfFunctionPlot[488].X = -0.15
	pointsOfFunctionPlot[488].Y = 0.812

	pointsOfFunctionPlot[489].X = -0.14
	pointsOfFunctionPlot[489].Y = 0.823

	pointsOfFunctionPlot[490].X = -0.13
	pointsOfFunctionPlot[490].Y = 0.835

	pointsOfFunctionPlot[491].X = -0.12
	pointsOfFunctionPlot[491].Y = 0.846

	pointsOfFunctionPlot[492].X = -0.11
	pointsOfFunctionPlot[492].Y = 0.858

	pointsOfFunctionPlot[493].X = -0.1
	pointsOfFunctionPlot[493].Y = 0.87

	pointsOfFunctionPlot[494].X = -0.09
	pointsOfFunctionPlot[494].Y = 0.882

	pointsOfFunctionPlot[495].X = -0.08
	pointsOfFunctionPlot[495].Y = 0.895

	pointsOfFunctionPlot[496].X = -0.07
	pointsOfFunctionPlot[496].Y = 0.907

	pointsOfFunctionPlot[497].X = -0.06
	pointsOfFunctionPlot[497].Y = 0.92

	pointsOfFunctionPlot[498].X = -0.05
	pointsOfFunctionPlot[498].Y = 0.933

	pointsOfFunctionPlot[499].X = -0.04
	pointsOfFunctionPlot[499].Y = 0.946

	pointsOfFunctionPlot[500].X = -0.03
	pointsOfFunctionPlot[500].Y = 0.959

	pointsOfFunctionPlot[501].X = -0.02
	pointsOfFunctionPlot[501].Y = 0.972

	pointsOfFunctionPlot[502].X = -0.01
	pointsOfFunctionPlot[502].Y = 0.986

	pointsOfFunctionPlot[503].X = 0.0
	pointsOfFunctionPlot[503].Y = 1.0

	pointsOfFunctionPlot[504].X = 0.01
	pointsOfFunctionPlot[504].Y = 1.014

	pointsOfFunctionPlot[505].X = 0.02
	pointsOfFunctionPlot[505].Y = 1.028

	pointsOfFunctionPlot[506].X = 0.03
	pointsOfFunctionPlot[506].Y = 1.042

	pointsOfFunctionPlot[507].X = 0.04
	pointsOfFunctionPlot[507].Y = 1.057

	pointsOfFunctionPlot[508].X = 0.05
	pointsOfFunctionPlot[508].Y = 1.071

	pointsOfFunctionPlot[509].X = 0.06
	pointsOfFunctionPlot[509].Y = 1.086

	pointsOfFunctionPlot[510].X = 0.07
	pointsOfFunctionPlot[510].Y = 1.101

	pointsOfFunctionPlot[511].X = 0.08
	pointsOfFunctionPlot[511].Y = 1.117

	pointsOfFunctionPlot[512].X = 0.09
	pointsOfFunctionPlot[512].Y = 1.132

	pointsOfFunctionPlot[513].X = 0.1
	pointsOfFunctionPlot[513].Y = 1.148

	pointsOfFunctionPlot[514].X = 0.11
	pointsOfFunctionPlot[514].Y = 1.164

	pointsOfFunctionPlot[515].X = 0.12
	pointsOfFunctionPlot[515].Y = 1.181

	pointsOfFunctionPlot[516].X = 0.13
	pointsOfFunctionPlot[516].Y = 1.197

	pointsOfFunctionPlot[517].X = 0.14
	pointsOfFunctionPlot[517].Y = 1.214

	pointsOfFunctionPlot[518].X = 0.15
	pointsOfFunctionPlot[518].Y = 1.231

	pointsOfFunctionPlot[519].X = 0.16
	pointsOfFunctionPlot[519].Y = 1.248

	pointsOfFunctionPlot[520].X = 0.17
	pointsOfFunctionPlot[520].Y = 1.265

	pointsOfFunctionPlot[521].X = 0.18
	pointsOfFunctionPlot[521].Y = 1.283

	pointsOfFunctionPlot[522].X = 0.19
	pointsOfFunctionPlot[522].Y = 1.301

	pointsOfFunctionPlot[523].X = 0.2
	pointsOfFunctionPlot[523].Y = 1.319

	pointsOfFunctionPlot[524].X = 0.21
	pointsOfFunctionPlot[524].Y = 1.337

	pointsOfFunctionPlot[525].X = 0.22
	pointsOfFunctionPlot[525].Y = 1.356

	pointsOfFunctionPlot[526].X = 0.23
	pointsOfFunctionPlot[526].Y = 1.375

	pointsOfFunctionPlot[527].X = 0.24
	pointsOfFunctionPlot[527].Y = 1.394

	pointsOfFunctionPlot[528].X = 0.25
	pointsOfFunctionPlot[528].Y = 1.414

	pointsOfFunctionPlot[529].X = 0.26
	pointsOfFunctionPlot[529].Y = 1.434

	pointsOfFunctionPlot[530].X = 0.27
	pointsOfFunctionPlot[530].Y = 1.454

	pointsOfFunctionPlot[531].X = 0.28
	pointsOfFunctionPlot[531].Y = 1.474

	pointsOfFunctionPlot[532].X = 0.29
	pointsOfFunctionPlot[532].Y = 1.494

	pointsOfFunctionPlot[533].X = 0.3
	pointsOfFunctionPlot[533].Y = 1.515

	pointsOfFunctionPlot[534].X = 0.31
	pointsOfFunctionPlot[534].Y = 1.536

	pointsOfFunctionPlot[535].X = 0.32
	pointsOfFunctionPlot[535].Y = 1.558

	pointsOfFunctionPlot[536].X = 0.33
	pointsOfFunctionPlot[536].Y = 1.58

	pointsOfFunctionPlot[537].X = 0.34
	pointsOfFunctionPlot[537].Y = 1.602

	pointsOfFunctionPlot[538].X = 0.35
	pointsOfFunctionPlot[538].Y = 1.624

	pointsOfFunctionPlot[539].X = 0.36
	pointsOfFunctionPlot[539].Y = 1.647

	pointsOfFunctionPlot[540].X = 0.37
	pointsOfFunctionPlot[540].Y = 1.67

	pointsOfFunctionPlot[541].X = 0.38
	pointsOfFunctionPlot[541].Y = 1.693

	pointsOfFunctionPlot[542].X = 0.39
	pointsOfFunctionPlot[542].Y = 1.717

	pointsOfFunctionPlot[543].X = 0.40
	pointsOfFunctionPlot[543].Y = 1.741

	pointsOfFunctionPlot[544].X = 0.41
	pointsOfFunctionPlot[544].Y = 1.765

	pointsOfFunctionPlot[545].X = 0.42
	pointsOfFunctionPlot[545].Y = 1.79

	pointsOfFunctionPlot[546].X = 0.43
	pointsOfFunctionPlot[546].Y = 1.815

	pointsOfFunctionPlot[547].X = 0.44
	pointsOfFunctionPlot[547].Y = 1.84

	pointsOfFunctionPlot[548].X = 0.45
	pointsOfFunctionPlot[548].Y = 1.866

	pointsOfFunctionPlot[549].X = 0.46
	pointsOfFunctionPlot[549].Y = 1.892

	pointsOfFunctionPlot[550].X = 0.47
	pointsOfFunctionPlot[550].Y = 1.918

	pointsOfFunctionPlot[551].X = 0.48
	pointsOfFunctionPlot[551].Y = 1.945

	pointsOfFunctionPlot[552].X = 0.49
	pointsOfFunctionPlot[552].Y = 1.972

	pointsOfFunctionPlot[553].X = 0.5
	pointsOfFunctionPlot[553].Y = 2.0

	pointsOfFunctionPlot[554].X = 0.51
	pointsOfFunctionPlot[554].Y = 2.027

	pointsOfFunctionPlot[555].X = 0.52
	pointsOfFunctionPlot[555].Y = 2.056

	pointsOfFunctionPlot[556].X = 0.53
	pointsOfFunctionPlot[556].Y = 2.084

	pointsOfFunctionPlot[557].X = 0.54
	pointsOfFunctionPlot[557].Y = 2.114

	pointsOfFunctionPlot[558].X = 0.55
	pointsOfFunctionPlot[558].Y = 2.143

	pointsOfFunctionPlot[559].X = 0.56
	pointsOfFunctionPlot[559].Y = 2.173

	pointsOfFunctionPlot[560].X = 0.57
	pointsOfFunctionPlot[560].Y = 2.203

	pointsOfFunctionPlot[561].X = 0.58
	pointsOfFunctionPlot[561].Y = 2.234

	pointsOfFunctionPlot[562].X = 0.59
	pointsOfFunctionPlot[562].Y = 2.265

	pointsOfFunctionPlot[563].X = 0.6
	pointsOfFunctionPlot[563].Y = 2.297

	pointsOfFunctionPlot[564].X = 0.61
	pointsOfFunctionPlot[564].Y = 2.297

	pointsOfFunctionPlot[565].X = 0.62
	pointsOfFunctionPlot[565].Y = 2.362

	pointsOfFunctionPlot[566].X = 0.63
	pointsOfFunctionPlot[566].Y = 2.395

	pointsOfFunctionPlot[567].X = 0.64
	pointsOfFunctionPlot[567].Y = 2.428

	pointsOfFunctionPlot[568].X = 0.65
	pointsOfFunctionPlot[568].Y = 2.462

	pointsOfFunctionPlot[569].X = 0.66
	pointsOfFunctionPlot[569].Y = 2.496

	pointsOfFunctionPlot[570].X = 0.67
	pointsOfFunctionPlot[570].Y = 2.531

	pointsOfFunctionPlot[571].X = 0.68
	pointsOfFunctionPlot[571].Y = 2.566

	pointsOfFunctionPlot[572].X = 0.69
	pointsOfFunctionPlot[572].Y = 2.602

	pointsOfFunctionPlot[573].X = 0.7
	pointsOfFunctionPlot[573].Y = 2.639

	pointsOfFunctionPlot[574].X = 0.71
	pointsOfFunctionPlot[574].Y = 2.675

	pointsOfFunctionPlot[575].X = 0.72
	pointsOfFunctionPlot[575].Y = 2.713

	pointsOfFunctionPlot[576].X = 0.73
	pointsOfFunctionPlot[576].Y = 2.751

	pointsOfFunctionPlot[577].X = 0.74
	pointsOfFunctionPlot[577].Y = 2.789

	pointsOfFunctionPlot[578].X = 0.75
	pointsOfFunctionPlot[578].Y = 2.828

	pointsOfFunctionPlot[579].X = 0.76
	pointsOfFunctionPlot[579].Y = 2.867

	pointsOfFunctionPlot[580].X = 0.77
	pointsOfFunctionPlot[580].Y = 2.907

	pointsOfFunctionPlot[581].X = 0.78
	pointsOfFunctionPlot[581].Y = 2.948

	pointsOfFunctionPlot[582].X = 0.79
	pointsOfFunctionPlot[582].Y = 2.989

	pointsOfFunctionPlot[583].X = 0.8
	pointsOfFunctionPlot[583].Y = 3.031

	pointsOfFunctionPlot[584].X = 0.81
	pointsOfFunctionPlot[584].Y = 3.073

	pointsOfFunctionPlot[585].X = 0.82
	pointsOfFunctionPlot[585].Y = 3.116

	pointsOfFunctionPlot[586].X = 0.83
	pointsOfFunctionPlot[586].Y = 3.16

	pointsOfFunctionPlot[587].X = 0.84
	pointsOfFunctionPlot[587].Y = 3.204

	pointsOfFunctionPlot[588].X = 0.85
	pointsOfFunctionPlot[588].Y = 3.249

	pointsOfFunctionPlot[589].X = 0.86
	pointsOfFunctionPlot[589].Y = 3.294

	pointsOfFunctionPlot[590].X = 0.87
	pointsOfFunctionPlot[590].Y = 3.34

	pointsOfFunctionPlot[591].X = 0.88
	pointsOfFunctionPlot[591].Y = 3.387

	pointsOfFunctionPlot[592].X = 0.89
	pointsOfFunctionPlot[592].Y = 3.434

	pointsOfFunctionPlot[593].X = 0.9
	pointsOfFunctionPlot[593].Y = 3.482

	pointsOfFunctionPlot[594].X = 0.91
	pointsOfFunctionPlot[594].Y = 3.53

	pointsOfFunctionPlot[595].X = 0.92
	pointsOfFunctionPlot[595].Y = 3.58

	pointsOfFunctionPlot[596].X = 0.93
	pointsOfFunctionPlot[596].Y = 3.63

	pointsOfFunctionPlot[597].X = 0.94
	pointsOfFunctionPlot[597].Y = 3.68

	pointsOfFunctionPlot[598].X = 0.95
	pointsOfFunctionPlot[598].Y = 3.732

	pointsOfFunctionPlot[599].X = 0.96
	pointsOfFunctionPlot[599].Y = 3.784

	pointsOfFunctionPlot[600].X = 0.97
	pointsOfFunctionPlot[600].Y = 3.837

	pointsOfFunctionPlot[601].X = 0.98
	pointsOfFunctionPlot[601].Y = 3.89

	pointsOfFunctionPlot[602].X = 0.99
	pointsOfFunctionPlot[602].Y = 3.944

	pointsOfFunctionPlot[603].X = 1.0
	pointsOfFunctionPlot[603].Y = 4.0

	pointsOfFunctionPlot[604].X = 1.01
	pointsOfFunctionPlot[604].Y = 4.055

	pointsOfFunctionPlot[605].X = 1.02
	pointsOfFunctionPlot[605].Y = 4.112

	pointsOfFunctionPlot[606].X = 1.03
	pointsOfFunctionPlot[606].Y = 4.169

	pointsOfFunctionPlot[607].X = 1.04
	pointsOfFunctionPlot[607].Y = 4.228

	pointsOfFunctionPlot[608].X = 1.05
	pointsOfFunctionPlot[608].Y = 4.287

	pointsOfFunctionPlot[609].X = 1.06
	pointsOfFunctionPlot[609].Y = 4.346

	pointsOfFunctionPlot[610].X = 1.07
	pointsOfFunctionPlot[610].Y = 4.407

	pointsOfFunctionPlot[611].X = 1.08
	pointsOfFunctionPlot[611].Y = 4.469

	pointsOfFunctionPlot[612].X = 1.09
	pointsOfFunctionPlot[612].Y = 4.531

	pointsOfFunctionPlot[613].X = 1.1
	pointsOfFunctionPlot[613].Y = 4.594

	pointsOfFunctionPlot[614].X = 1.11
	pointsOfFunctionPlot[614].Y = 4.658

	pointsOfFunctionPlot[615].X = 1.12
	pointsOfFunctionPlot[615].Y = 4.742

	pointsOfFunctionPlot[616].X = 1.13
	pointsOfFunctionPlot[616].Y = 4.789

	pointsOfFunctionPlot[617].X = 1.14
	pointsOfFunctionPlot[617].Y = 4.856

	pointsOfFunctionPlot[618].X = 1.15
	pointsOfFunctionPlot[618].Y = 4.924

	pointsOfFunctionPlot[619].X = 1.16
	pointsOfFunctionPlot[619].Y = 4.993

	pointsOfFunctionPlot[620].X = 1.17
	pointsOfFunctionPlot[620].Y = 5.063

	pointsOfFunctionPlot[621].X = 1.18
	pointsOfFunctionPlot[621].Y = 5.133

	pointsOfFunctionPlot[622].X = 1.19
	pointsOfFunctionPlot[622].Y = 5.204

	pointsOfFunctionPlot[623].X = 1.2
	pointsOfFunctionPlot[623].Y = 5.278

	pointsOfFunctionPlot[624].X = 1.21
	pointsOfFunctionPlot[624].Y = 5.351

	pointsOfFunctionPlot[625].X = 1.22
	pointsOfFunctionPlot[625].Y = 5.426

	pointsOfFunctionPlot[626].X = 1.23
	pointsOfFunctionPlot[626].Y = 5.502

	pointsOfFunctionPlot[627].X = 1.24
	pointsOfFunctionPlot[627].Y = 5.579

	pointsOfFunctionPlot[628].X = 1.25
	pointsOfFunctionPlot[628].Y = 5.656

	pointsOfFunctionPlot[629].X = 1.26
	pointsOfFunctionPlot[629].Y = 5.735

	pointsOfFunctionPlot[630].X = 1.27
	pointsOfFunctionPlot[630].Y = 5.815

	pointsOfFunctionPlot[631].X = 1.28
	pointsOfFunctionPlot[631].Y = 5.897

	pointsOfFunctionPlot[632].X = 1.29
	pointsOfFunctionPlot[632].Y = 5.979

	pointsOfFunctionPlot[633].X = 1.3
	pointsOfFunctionPlot[633].Y = 6.062

	pointsOfFunctionPlot[634].X = 1.31
	pointsOfFunctionPlot[634].Y = 6.147

	pointsOfFunctionPlot[635].X = 1.32
	pointsOfFunctionPlot[635].Y = 6.233

	pointsOfFunctionPlot[636].X = 1.33
	pointsOfFunctionPlot[636].Y = 6.32

	pointsOfFunctionPlot[637].X = 1.34
	pointsOfFunctionPlot[637].Y = 6.408

	pointsOfFunctionPlot[638].X = 1.35
	pointsOfFunctionPlot[638].Y = 6.498

	pointsOfFunctionPlot[639].X = 1.36
	pointsOfFunctionPlot[639].Y = 6.588

	pointsOfFunctionPlot[640].X = 1.37
	pointsOfFunctionPlot[640].Y = 6.68

	pointsOfFunctionPlot[641].X = 1.38
	pointsOfFunctionPlot[641].Y = 6.774

	pointsOfFunctionPlot[642].X = 1.39
	pointsOfFunctionPlot[642].Y = 6.868

	pointsOfFunctionPlot[643].X = 1.40
	pointsOfFunctionPlot[643].Y = 6.964

	pointsOfFunctionPlot[644].X = 1.41
	pointsOfFunctionPlot[644].Y = 7.061

	pointsOfFunctionPlot[645].X = 1.42
	pointsOfFunctionPlot[645].Y = 7.16

	pointsOfFunctionPlot[646].X = 1.43
	pointsOfFunctionPlot[646].Y = 7.26

	pointsOfFunctionPlot[647].X = 1.44
	pointsOfFunctionPlot[647].Y = 7.361

	pointsOfFunctionPlot[648].X = 1.45
	pointsOfFunctionPlot[648].Y = 7.464

	pointsOfFunctionPlot[649].X = 1.46
	pointsOfFunctionPlot[649].Y = 7.568

	pointsOfFunctionPlot[650].X = 1.47
	pointsOfFunctionPlot[650].Y = 7.674

	pointsOfFunctionPlot[651].X = 1.48
	pointsOfFunctionPlot[651].Y = 7.781

	pointsOfFunctionPlot[652].X = 1.49
	pointsOfFunctionPlot[652].Y = 7.889

	pointsOfFunctionPlot[653].X = 1.5
	pointsOfFunctionPlot[653].Y = 8.0

	pointsOfFunctionPlot[654].X = 1.51
	pointsOfFunctionPlot[654].Y = 8.111

	pointsOfFunctionPlot[655].X = 1.52
	pointsOfFunctionPlot[655].Y = 8.224

	pointsOfFunctionPlot[656].X = 1.53
	pointsOfFunctionPlot[656].Y = 8.339

	pointsOfFunctionPlot[657].X = 1.54
	pointsOfFunctionPlot[657].Y = 8.465

	pointsOfFunctionPlot[658].X = 1.55
	pointsOfFunctionPlot[658].Y = 8.574

	pointsOfFunctionPlot[659].X = 1.56
	pointsOfFunctionPlot[659].Y = 8.693

	pointsOfFunctionPlot[660].X = 1.57
	pointsOfFunctionPlot[660].Y = 8.815

	pointsOfFunctionPlot[661].X = 1.58
	pointsOfFunctionPlot[661].Y = 8.938

	pointsOfFunctionPlot[662].X = 1.59
	pointsOfFunctionPlot[662].Y = 9.063

	pointsOfFunctionPlot[663].X = 1.6
	pointsOfFunctionPlot[663].Y = 9.189

	pointsOfFunctionPlot[664].X = 1.61
	pointsOfFunctionPlot[664].Y = 9.317

	pointsOfFunctionPlot[665].X = 1.62
	pointsOfFunctionPlot[665].Y = 9.447

	pointsOfFunctionPlot[666].X = 1.63
	pointsOfFunctionPlot[666].Y = 9.579

	pointsOfFunctionPlot[667].X = 1.64
	pointsOfFunctionPlot[667].Y = 9.713

	pointsOfFunctionPlot[668].X = 1.65
	pointsOfFunctionPlot[668].Y = 9.849

	pointsOfFunctionPlot[669].X = 1.66
	pointsOfFunctionPlot[669].Y = 9.986

	pointsOfFunctionPlot[670].X = 1.67
	pointsOfFunctionPlot[670].Y = 10.126

	pointsOfFunctionPlot[671].X = 1.68
	pointsOfFunctionPlot[671].Y = 10.267

	pointsOfFunctionPlot[672].X = 1.69
	pointsOfFunctionPlot[672].Y = 10.41

	pointsOfFunctionPlot[673].X = 1.7
	pointsOfFunctionPlot[673].Y = 10.556

	pointsOfFunctionPlot[674].X = 1.71
	pointsOfFunctionPlot[674].Y = 10.703

	pointsOfFunctionPlot[675].X = 1.72
	pointsOfFunctionPlot[675].Y = 10.852

	pointsOfFunctionPlot[676].X = 1.73
	pointsOfFunctionPlot[676].Y = 11.004

	pointsOfFunctionPlot[677].X = 1.74
	pointsOfFunctionPlot[677].Y = 11.157

	pointsOfFunctionPlot[678].X = 1.75
	pointsOfFunctionPlot[678].Y = 11.313

	pointsOfFunctionPlot[679].X = 1.76
	pointsOfFunctionPlot[679].Y = 11.471

	pointsOfFunctionPlot[680].X = 1.77
	pointsOfFunctionPlot[680].Y = 11.631

	pointsOfFunctionPlot[681].X = 1.78
	pointsOfFunctionPlot[681].Y = 11.794

	pointsOfFunctionPlot[682].X = 1.79
	pointsOfFunctionPlot[682].Y = 11.958

	pointsOfFunctionPlot[683].X = 1.8
	pointsOfFunctionPlot[683].Y = 12.125

	pointsOfFunctionPlot[684].X = 1.81
	pointsOfFunctionPlot[684].Y = 12.295

	pointsOfFunctionPlot[685].X = 1.82
	pointsOfFunctionPlot[685].Y = 12.466

	pointsOfFunctionPlot[686].X = 1.83
	pointsOfFunctionPlot[686].Y = 12.64

	pointsOfFunctionPlot[687].X = 1.84
	pointsOfFunctionPlot[687].Y = 12.817

	pointsOfFunctionPlot[688].X = 1.85
	pointsOfFunctionPlot[688].Y = 12.996

	pointsOfFunctionPlot[689].X = 1.86
	pointsOfFunctionPlot[689].Y = 13.177

	pointsOfFunctionPlot[690].X = 1.87
	pointsOfFunctionPlot[690].Y = 13.361

	pointsOfFunctionPlot[691].X = 1.88
	pointsOfFunctionPlot[691].Y = 13.547

	pointsOfFunctionPlot[692].X = 1.89
	pointsOfFunctionPlot[692].Y = 13.737

	pointsOfFunctionPlot[693].X = 1.9
	pointsOfFunctionPlot[693].Y = 13.928

	pointsOfFunctionPlot[694].X = 1.91
	pointsOfFunctionPlot[694].Y = 14.123

	pointsOfFunctionPlot[695].X = 1.92
	pointsOfFunctionPlot[695].Y = 14.32

	pointsOfFunctionPlot[696].X = 1.93
	pointsOfFunctionPlot[696].Y = 14.52

	pointsOfFunctionPlot[697].X = 1.94
	pointsOfFunctionPlot[697].Y = 14.723

	pointsOfFunctionPlot[698].X = 1.95
	pointsOfFunctionPlot[698].Y = 14.928

	pointsOfFunctionPlot[699].X = 1.96
	pointsOfFunctionPlot[699].Y = 15.136

	pointsOfFunctionPlot[700].X = 1.97
	pointsOfFunctionPlot[700].Y = 15.348

	pointsOfFunctionPlot[701].X = 1.98
	pointsOfFunctionPlot[701].Y = 15.562

	pointsOfFunctionPlot[702].X = 1.99
	pointsOfFunctionPlot[702].Y = 15.779

	pointsOfFunctionPlot[703].X = 2.0
	pointsOfFunctionPlot[703].Y = 16.0

	pointsOfFunctionPlot[704].X = 2.01
	pointsOfFunctionPlot[704].Y = 16.223

	pointsOfFunctionPlot[705].X = 2.02
	pointsOfFunctionPlot[705].Y = 16.449

	pointsOfFunctionPlot[706].X = 2.03
	pointsOfFunctionPlot[706].Y = 16.679

	pointsOfFunctionPlot[707].X = 2.04
	pointsOfFunctionPlot[707].Y = 16.912

	pointsOfFunctionPlot[708].X = 2.05
	pointsOfFunctionPlot[708].Y = 17.148

	pointsOfFunctionPlot[709].X = 2.06
	pointsOfFunctionPlot[709].Y = 17.387

	pointsOfFunctionPlot[710].X = 2.07
	pointsOfFunctionPlot[710].Y = 17.63

	pointsOfFunctionPlot[711].X = 2.08
	pointsOfFunctionPlot[711].Y = 17.876

	pointsOfFunctionPlot[712].X = 2.09
	pointsOfFunctionPlot[712].Y = 18.126

	pointsOfFunctionPlot[713].X = 2.1
	pointsOfFunctionPlot[713].Y = 18.379

	pointsOfFunctionPlot[714].X = 2.11
	pointsOfFunctionPlot[714].Y = 18.635

	pointsOfFunctionPlot[715].X = 2.12
	pointsOfFunctionPlot[715].Y = 18.895

	pointsOfFunctionPlot[716].X = 2.13
	pointsOfFunctionPlot[716].Y = 19.159

	pointsOfFunctionPlot[717].X = 2.14
	pointsOfFunctionPlot[717].Y = 19.427

	pointsOfFunctionPlot[718].X = 2.15
	pointsOfFunctionPlot[718].Y = 19.698

	pointsOfFunctionPlot[719].X = 2.16
	pointsOfFunctionPlot[719].Y = 19.973

	pointsOfFunctionPlot[720].X = 2.17
	pointsOfFunctionPlot[720].Y = 20.252

	pointsOfFunctionPlot[721].X = 2.18
	pointsOfFunctionPlot[721].Y = 20.534

	pointsOfFunctionPlot[722].X = 2.19
	pointsOfFunctionPlot[722].Y = 20.821

	pointsOfFunctionPlot[723].X = 2.2
	pointsOfFunctionPlot[723].Y = 21.112

	pointsOfFunctionPlot[724].X = 2.21
	pointsOfFunctionPlot[724].Y = 21.406

	pointsOfFunctionPlot[725].X = 2.22
	pointsOfFunctionPlot[725].Y = 21.705

	pointsOfFunctionPlot[726].X = 2.23
	pointsOfFunctionPlot[726].Y = 22.008

	pointsOfFunctionPlot[727].X = 2.24
	pointsOfFunctionPlot[727].Y = 22.315

	pointsOfFunctionPlot[728].X = 2.25
	pointsOfFunctionPlot[728].Y = 22.627

	pointsOfFunctionPlot[729].X = 2.26
	pointsOfFunctionPlot[729].Y = 22.943

	pointsOfFunctionPlot[730].X = 2.27
	pointsOfFunctionPlot[730].Y = 23.263

	pointsOfFunctionPlot[731].X = 2.28
	pointsOfFunctionPlot[731].Y = 23.588

	pointsOfFunctionPlot[732].X = 2.29
	pointsOfFunctionPlot[732].Y = 23.917

	pointsOfFunctionPlot[733].X = 2.3
	pointsOfFunctionPlot[733].Y = 24.251

	pointsOfFunctionPlot[734].X = 2.31
	pointsOfFunctionPlot[734].Y = 24.59

	pointsOfFunctionPlot[735].X = 2.32
	pointsOfFunctionPlot[735].Y = 24.933

	pointsOfFunctionPlot[736].X = 2.33
	pointsOfFunctionPlot[736].Y = 25.281

	pointsOfFunctionPlot[737].X = 2.34
	pointsOfFunctionPlot[737].Y = 25.634

	pointsOfFunctionPlot[738].X = 2.35
	pointsOfFunctionPlot[738].Y = 25.992

	pointsOfFunctionPlot[739].X = 2.36
	pointsOfFunctionPlot[739].Y = 26.354

	pointsOfFunctionPlot[740].X = 2.37
	pointsOfFunctionPlot[740].Y = 26.722

	pointsOfFunctionPlot[741].X = 2.38
	pointsOfFunctionPlot[741].Y = 27.095

	pointsOfFunctionPlot[742].X = 2.39
	pointsOfFunctionPlot[742].Y = 27.474

	pointsOfFunctionPlot[743].X = 2.4
	pointsOfFunctionPlot[743].Y = 27.857

	pointsOfFunctionPlot[744].X = 2.41
	pointsOfFunctionPlot[744].Y = 28.246

	pointsOfFunctionPlot[745].X = 2.42
	pointsOfFunctionPlot[745].Y = 28.64

	pointsOfFunctionPlot[746].X = 2.43
	pointsOfFunctionPlot[746].Y = 29.04

	pointsOfFunctionPlot[747].X = 2.44
	pointsOfFunctionPlot[747].Y = 29.446

	pointsOfFunctionPlot[748].X = 2.45
	pointsOfFunctionPlot[748].Y = 29.857

	pointsOfFunctionPlot[749].X = 2.46
	pointsOfFunctionPlot[749].Y = 30.273

	pointsOfFunctionPlot[750].X = 2.47
	pointsOfFunctionPlot[750].Y = 30.696

	pointsOfFunctionPlot[751].X = 2.48
	pointsOfFunctionPlot[751].Y = 31.125

	pointsOfFunctionPlot[752].X = 2.49
	pointsOfFunctionPlot[752].Y = 31.559

	pointsOfFunctionPlot[753].X = 2.5
	pointsOfFunctionPlot[753].Y = 32.0

	pointsOfFunctionPlot[754].X = 2.51
	pointsOfFunctionPlot[754].Y = 32.446

	pointsOfFunctionPlot[755].X = 2.52
	pointsOfFunctionPlot[755].Y = 32.899

	pointsOfFunctionPlot[756].X = 2.53
	pointsOfFunctionPlot[756].Y = 33.358

	pointsOfFunctionPlot[757].X = 2.54
	pointsOfFunctionPlot[757].Y = 33.824

	pointsOfFunctionPlot[758].X = 2.55
	pointsOfFunctionPlot[758].Y = 34.296

	pointsOfFunctionPlot[759].X = 2.56
	pointsOfFunctionPlot[759].Y = 34.775

	pointsOfFunctionPlot[760].X = 2.57
	pointsOfFunctionPlot[760].Y = 35.261

	pointsOfFunctionPlot[761].X = 2.58
	pointsOfFunctionPlot[761].Y = 35.753

	pointsOfFunctionPlot[762].X = 2.59
	pointsOfFunctionPlot[762].Y = 36.252

	pointsOfFunctionPlot[763].X = 2.6
	pointsOfFunctionPlot[763].Y = 36.758

	pointsOfFunctionPlot[764].X = 2.61
	pointsOfFunctionPlot[764].Y = 37.271

	pointsOfFunctionPlot[765].X = 2.62
	pointsOfFunctionPlot[765].Y = 37.791

	pointsOfFunctionPlot[766].X = 2.63
	pointsOfFunctionPlot[766].Y = 38.319

	pointsOfFunctionPlot[767].X = 2.64
	pointsOfFunctionPlot[767].Y = 38.854

	pointsOfFunctionPlot[768].X = 2.65
	pointsOfFunctionPlot[768].Y = 39.396

	pointsOfFunctionPlot[769].X = 2.66
	pointsOfFunctionPlot[769].Y = 39.946

	pointsOfFunctionPlot[770].X = 2.67
	pointsOfFunctionPlot[770].Y = 40.504

	pointsOfFunctionPlot[771].X = 2.68
	pointsOfFunctionPlot[771].Y = 41.069

	pointsOfFunctionPlot[772].X = 2.69
	pointsOfFunctionPlot[772].Y = 41.642

	pointsOfFunctionPlot[773].X = 2.7
	pointsOfFunctionPlot[773].Y = 42.224

	pointsOfFunctionPlot[774].X = 2.71
	pointsOfFunctionPlot[774].Y = 42.813

	pointsOfFunctionPlot[775].X = 2.72
	pointsOfFunctionPlot[775].Y = 43.411

	pointsOfFunctionPlot[776].X = 2.73
	pointsOfFunctionPlot[776].Y = 44.017

	pointsOfFunctionPlot[777].X = 2.74
	pointsOfFunctionPlot[777].Y = 44.631

	pointsOfFunctionPlot[778].X = 2.75
	pointsOfFunctionPlot[778].Y = 45.254

	pointsOfFunctionPlot[779].X = 2.76
	pointsOfFunctionPlot[779].Y = 45.886

	pointsOfFunctionPlot[780].X = 2.77
	pointsOfFunctionPlot[780].Y = 46.527

	pointsOfFunctionPlot[781].X = 2.78
	pointsOfFunctionPlot[781].Y = 47.176

	pointsOfFunctionPlot[782].X = 2.79
	pointsOfFunctionPlot[782].Y = 47.835

	pointsOfFunctionPlot[783].X = 2.8
	pointsOfFunctionPlot[783].Y = 48.502

	pointsOfFunctionPlot[784].X = 2.81
	pointsOfFunctionPlot[784].Y = 48.502

	pointsOfFunctionPlot[785].X = 2.82
	pointsOfFunctionPlot[785].Y = 49.866

	pointsOfFunctionPlot[786].X = 2.83
	pointsOfFunctionPlot[786].Y = 50.562

	pointsOfFunctionPlot[787].X = 2.84
	pointsOfFunctionPlot[787].Y = 51.268

	pointsOfFunctionPlot[788].X = 2.85
	pointsOfFunctionPlot[788].Y = 51.984

	pointsOfFunctionPlot[789].X = 2.86
	pointsOfFunctionPlot[789].Y = 52.709

	pointsOfFunctionPlot[790].X = 2.87
	pointsOfFunctionPlot[790].Y = 53.445

	pointsOfFunctionPlot[791].X = 2.88
	pointsOfFunctionPlot[791].Y = 54.191

	pointsOfFunctionPlot[792].X = 2.89
	pointsOfFunctionPlot[792].Y = 54.948

	pointsOfFunctionPlot[793].X = 2.9
	pointsOfFunctionPlot[793].Y = 55.715

	pointsOfFunctionPlot[794].X = 2.91
	pointsOfFunctionPlot[794].Y = 56.493

	pointsOfFunctionPlot[795].X = 2.92
	pointsOfFunctionPlot[795].Y = 57.281

	pointsOfFunctionPlot[796].X = 2.93
	pointsOfFunctionPlot[796].Y = 58.081

	pointsOfFunctionPlot[797].X = 2.94
	pointsOfFunctionPlot[797].Y = 58.892

	pointsOfFunctionPlot[798].X = 2.95
	pointsOfFunctionPlot[798].Y = 59.714

	pointsOfFunctionPlot[799].X = 2.96
	pointsOfFunctionPlot[799].Y = 60.547

	pointsOfFunctionPlot[800].X = 2.97
	pointsOfFunctionPlot[800].Y = 61.392

	pointsOfFunctionPlot[801].X = 2.98
	pointsOfFunctionPlot[801].Y = 62.249

	pointsOfFunctionPlot[802].X = 2.99
	pointsOfFunctionPlot[802].Y = 63.118

	pointsOfFunctionPlot[803].X = 3.0
	pointsOfFunctionPlot[803].Y = 64.0

	pointsOfFunctionPlot[804].X = 3.01
	pointsOfFunctionPlot[804].Y = 64.893

	pointsOfFunctionPlot[805].X = 3.02
	pointsOfFunctionPlot[805].Y = 65.799

	pointsOfFunctionPlot[806].X = 3.03
	pointsOfFunctionPlot[806].Y = 66.717

	pointsOfFunctionPlot[807].X = 3.04
	pointsOfFunctionPlot[807].Y = 67.649

	pointsOfFunctionPlot[808].X = 3.05
	pointsOfFunctionPlot[808].Y = 68.593

	pointsOfFunctionPlot[809].X = 3.06
	pointsOfFunctionPlot[809].Y = 69.551

	pointsOfFunctionPlot[810].X = 3.07
	pointsOfFunctionPlot[810].Y = 70.521

	pointsOfFunctionPlot[811].X = 3.08
	pointsOfFunctionPlot[811].Y = 71.506

	pointsOfFunctionPlot[812].X = 3.09
	pointsOfFunctionPlot[812].Y = 72.504

	pointsOfFunctionPlot[813].X = 3.1
	pointsOfFunctionPlot[813].Y = 73.516

	pointsOfFunctionPlot[814].X = 3.11
	pointsOfFunctionPlot[814].Y = 74.542

	pointsOfFunctionPlot[815].X = 3.12
	pointsOfFunctionPlot[815].Y = 75.583

	pointsOfFunctionPlot[816].X = 3.13
	pointsOfFunctionPlot[816].Y = 76.638

	pointsOfFunctionPlot[817].X = 3.14
	pointsOfFunctionPlot[817].Y = 77.708

	pointsOfFunctionPlot[818].X = 3.15
	pointsOfFunctionPlot[818].Y = 78.793

	pointsOfFunctionPlot[819].X = 3.16
	pointsOfFunctionPlot[819].Y = 79.893

	pointsOfFunctionPlot[820].X = 3.17
	pointsOfFunctionPlot[820].Y = 81.008

	pointsOfFunctionPlot[821].X = 3.18
	pointsOfFunctionPlot[821].Y = 82.139

	pointsOfFunctionPlot[822].X = 3.19
	pointsOfFunctionPlot[822].Y = 83.285

	pointsOfFunctionPlot[823].X = 3.2
	pointsOfFunctionPlot[823].Y = 84.448

	pointsOfFunctionPlot[824].X = 3.21
	pointsOfFunctionPlot[824].Y = 85.627

	pointsOfFunctionPlot[825].X = 3.22
	pointsOfFunctionPlot[825].Y = 86.822

	pointsOfFunctionPlot[826].X = 3.23
	pointsOfFunctionPlot[826].Y = 88.034

	pointsOfFunctionPlot[827].X = 3.24
	pointsOfFunctionPlot[827].Y = 89.263

	pointsOfFunctionPlot[828].X = 3.25
	pointsOfFunctionPlot[828].Y = 90.509

	pointsOfFunctionPlot[829].X = 3.26
	pointsOfFunctionPlot[829].Y = 91.773

	pointsOfFunctionPlot[830].X = 3.27
	pointsOfFunctionPlot[830].Y = 93.054

	pointsOfFunctionPlot[831].X = 3.28
	pointsOfFunctionPlot[831].Y = 94.353

	pointsOfFunctionPlot[832].X = 3.29
	pointsOfFunctionPlot[832].Y = 95.67

	pointsOfFunctionPlot[833].X = 3.3
	pointsOfFunctionPlot[833].Y = 97.005

	pointsOfFunctionPlot[834].X = 3.31
	pointsOfFunctionPlot[834].Y = 98.36

	pointsOfFunctionPlot[835].X = 3.32
	pointsOfFunctionPlot[835].Y = 99.733

	pointsOfFunctionPlot[836].X = 3.33
	pointsOfFunctionPlot[836].Y = 101.125

	pointsOfFunctionPlot[837].X = 3.34
	pointsOfFunctionPlot[837].Y = 102.536

	pointsOfFunctionPlot[838].X = 3.35
	pointsOfFunctionPlot[838].Y = 103.968

	pointsOfFunctionPlot[839].X = 3.36
	pointsOfFunctionPlot[839].Y = 105.419

	pointsOfFunctionPlot[840].X = 3.37
	pointsOfFunctionPlot[840].Y = 106.891

	pointsOfFunctionPlot[841].X = 3.38
	pointsOfFunctionPlot[841].Y = 108.383

	pointsOfFunctionPlot[842].X = 3.39
	pointsOfFunctionPlot[842].Y = 109.896

	pointsOfFunctionPlot[843].X = 3.4
	pointsOfFunctionPlot[843].Y = 111.43

	pointsOfFunctionPlot[844].X = 3.41
	pointsOfFunctionPlot[844].Y = 112.986

	pointsOfFunctionPlot[845].X = 3.42
	pointsOfFunctionPlot[845].Y = 114.563

	pointsOfFunctionPlot[846].X = 3.43
	pointsOfFunctionPlot[846].Y = 116.162

	pointsOfFunctionPlot[847].X = 3.44
	pointsOfFunctionPlot[847].Y = 117.784

	pointsOfFunctionPlot[848].X = 3.45
	pointsOfFunctionPlot[848].Y = 119.428

	pointsOfFunctionPlot[849].X = 3.46
	pointsOfFunctionPlot[849].Y = 121.095

	pointsOfFunctionPlot[850].X = 3.47
	pointsOfFunctionPlot[850].Y = 122.785

	pointsOfFunctionPlot[851].X = 3.48
	pointsOfFunctionPlot[851].Y = 124.499

	pointsOfFunctionPlot[852].X = 3.49
	pointsOfFunctionPlot[852].Y = 126.237

	pointsOfFunctionPlot[853].X = 3.5
	pointsOfFunctionPlot[853].Y = 128.0

	pointsOfFunctionPlot[854].X = 3.51
	pointsOfFunctionPlot[854].Y = 129.786

	pointsOfFunctionPlot[855].X = 3.52
	pointsOfFunctionPlot[855].Y = 131.598

	pointsOfFunctionPlot[856].X = 3.53
	pointsOfFunctionPlot[856].Y = 133.435

	pointsOfFunctionPlot[857].X = 3.54
	pointsOfFunctionPlot[857].Y = 135.298

	pointsOfFunctionPlot[858].X = 3.55
	pointsOfFunctionPlot[858].Y = 137.187

	pointsOfFunctionPlot[859].X = 3.56
	pointsOfFunctionPlot[859].Y = 139.102

	pointsOfFunctionPlot[860].X = 3.57
	pointsOfFunctionPlot[860].Y = 141.043

	pointsOfFunctionPlot[861].X = 3.58
	pointsOfFunctionPlot[861].Y = 143.012

	pointsOfFunctionPlot[862].X = 3.59
	pointsOfFunctionPlot[862].Y = 145.009

	pointsOfFunctionPlot[863].X = 3.6
	pointsOfFunctionPlot[863].Y = 147.033

	pointsOfFunctionPlot[864].X = 3.61
	pointsOfFunctionPlot[864].Y = 149.085

	pointsOfFunctionPlot[865].X = 3.62
	pointsOfFunctionPlot[865].Y = 151.167

	pointsOfFunctionPlot[866].X = 3.63
	pointsOfFunctionPlot[866].Y = 153.277

	pointsOfFunctionPlot[867].X = 3.64
	pointsOfFunctionPlot[867].Y = 155.416

	pointsOfFunctionPlot[868].X = 3.65
	pointsOfFunctionPlot[868].Y = 157.586

	pointsOfFunctionPlot[869].X = 3.66
	pointsOfFunctionPlot[869].Y = 159.786

	pointsOfFunctionPlot[870].X = 3.67
	pointsOfFunctionPlot[870].Y = 162.016

	pointsOfFunctionPlot[871].X = 3.68
	pointsOfFunctionPlot[871].Y = 164.278

	pointsOfFunctionPlot[872].X = 3.69
	pointsOfFunctionPlot[872].Y = 166.571

	pointsOfFunctionPlot[873].X = 3.7
	pointsOfFunctionPlot[873].Y = 168.897

	pointsOfFunctionPlot[874].X = 3.71
	pointsOfFunctionPlot[874].Y = 171.254

	pointsOfFunctionPlot[875].X = 3.72
	pointsOfFunctionPlot[875].Y = 173.645

	pointsOfFunctionPlot[876].X = 3.73
	pointsOfFunctionPlot[876].Y = 176.069

	pointsOfFunctionPlot[877].X = 3.74
	pointsOfFunctionPlot[877].Y = 178.527

	pointsOfFunctionPlot[878].X = 3.75
	pointsOfFunctionPlot[878].Y = 181.019

	pointsOfFunctionPlot[879].X = 3.76
	pointsOfFunctionPlot[879].Y = 183.546

	pointsOfFunctionPlot[880].X = 3.77
	pointsOfFunctionPlot[880].Y = 186.108

	pointsOfFunctionPlot[881].X = 3.78
	pointsOfFunctionPlot[881].Y = 188.706

	pointsOfFunctionPlot[882].X = 3.79
	pointsOfFunctionPlot[882].Y = 191.34

	pointsOfFunctionPlot[883].X = 3.8
	pointsOfFunctionPlot[883].Y = 194.011

	pointsOfFunctionPlot[884].X = 3.81
	pointsOfFunctionPlot[884].Y = 194.011

	pointsOfFunctionPlot[885].X = 3.82
	pointsOfFunctionPlot[885].Y = 199.466

	pointsOfFunctionPlot[886].X = 3.83
	pointsOfFunctionPlot[886].Y = 202.25

	pointsOfFunctionPlot[887].X = 3.84
	pointsOfFunctionPlot[887].Y = 205.073

	pointsOfFunctionPlot[888].X = 3.85
	pointsOfFunctionPlot[888].Y = 207.936

	pointsOfFunctionPlot[889].X = 3.86
	pointsOfFunctionPlot[889].Y = 210.839

	pointsOfFunctionPlot[890].X = 3.87
	pointsOfFunctionPlot[890].Y = 213.782

	pointsOfFunctionPlot[891].X = 3.88
	pointsOfFunctionPlot[891].Y = 216.766

	pointsOfFunctionPlot[892].X = 3.89
	pointsOfFunctionPlot[892].Y = 219.792

	pointsOfFunctionPlot[893].X = 3.9
	pointsOfFunctionPlot[893].Y = 222.86

	pointsOfFunctionPlot[894].X = 3.91
	pointsOfFunctionPlot[894].Y = 225.972

	pointsOfFunctionPlot[895].X = 3.92
	pointsOfFunctionPlot[895].Y = 229.126

	pointsOfFunctionPlot[896].X = 3.93
	pointsOfFunctionPlot[896].Y = 232.324

	pointsOfFunctionPlot[897].X = 3.94
	pointsOfFunctionPlot[897].Y = 235.568

	pointsOfFunctionPlot[898].X = 3.95
	pointsOfFunctionPlot[898].Y = 238.856

	pointsOfFunctionPlot[899].X = 3.96
	pointsOfFunctionPlot[899].Y = 242.19

	pointsOfFunctionPlot[900].X = 3.97
	pointsOfFunctionPlot[900].Y = 245.571

	pointsOfFunctionPlot[901].X = 3.98
	pointsOfFunctionPlot[901].Y = 248.999

	pointsOfFunctionPlot[902].X = 3.99
	pointsOfFunctionPlot[902].Y = 252.475

	pointsOfFunctionPlot[903].X = 4.0
	pointsOfFunctionPlot[903].Y = 256.0

	pointsOfFunctionPlot[904].X = 4.01
	pointsOfFunctionPlot[904].Y = 259.573

	pointsOfFunctionPlot[905].X = 4.02
	pointsOfFunctionPlot[905].Y = 263.197

	pointsOfFunctionPlot[906].X = 4.03
	pointsOfFunctionPlot[906].Y = 266.871

	pointsOfFunctionPlot[907].X = 4.04
	pointsOfFunctionPlot[907].Y = 270.596

	pointsOfFunctionPlot[908].X = 4.05
	pointsOfFunctionPlot[908].Y = 274.374

	pointsOfFunctionPlot[909].X = 4.06
	pointsOfFunctionPlot[909].Y = 278.204

	pointsOfFunctionPlot[910].X = 4.07
	pointsOfFunctionPlot[910].Y = 282.087

	pointsOfFunctionPlot[911].X = 4.08
	pointsOfFunctionPlot[911].Y = 286.025

	pointsOfFunctionPlot[912].X = 4.09
	pointsOfFunctionPlot[912].Y = 290.018

	pointsOfFunctionPlot[913].X = 4.1
	pointsOfFunctionPlot[913].Y = 294.066

	pointsOfFunctionPlot[914].X = 4.11
	pointsOfFunctionPlot[914].Y = 298.171

	pointsOfFunctionPlot[915].X = 4.12
	pointsOfFunctionPlot[915].Y = 302.334

	pointsOfFunctionPlot[916].X = 4.13
	pointsOfFunctionPlot[916].Y = 306.554

	pointsOfFunctionPlot[917].X = 4.14
	pointsOfFunctionPlot[917].Y = 310.833

	pointsOfFunctionPlot[918].X = 4.15
	pointsOfFunctionPlot[918].Y = 315.173

	pointsOfFunctionPlot[919].X = 4.16
	pointsOfFunctionPlot[919].Y = 319.572

	pointsOfFunctionPlot[920].X = 4.17
	pointsOfFunctionPlot[920].Y = 324.033

	pointsOfFunctionPlot[921].X = 4.18
	pointsOfFunctionPlot[921].Y = 328.557

	pointsOfFunctionPlot[922].X = 4.19
	pointsOfFunctionPlot[922].Y = 333.143

	pointsOfFunctionPlot[923].X = 4.2
	pointsOfFunctionPlot[923].Y = 337.794

	pointsOfFunctionPlot[924].X = 4.21
	pointsOfFunctionPlot[924].Y = 342.509

	pointsOfFunctionPlot[925].X = 4.22
	pointsOfFunctionPlot[925].Y = 347.29

	pointsOfFunctionPlot[926].X = 4.23
	pointsOfFunctionPlot[926].Y = 352.138

	pointsOfFunctionPlot[927].X = 4.24
	pointsOfFunctionPlot[927].Y = 357.054

	pointsOfFunctionPlot[928].X = 4.25
	pointsOfFunctionPlot[928].Y = 362.038

	pointsOfFunctionPlot[929].X = 4.26
	pointsOfFunctionPlot[929].Y = 367.092

	pointsOfFunctionPlot[930].X = 4.27
	pointsOfFunctionPlot[930].Y = 372.217

	pointsOfFunctionPlot[931].X = 4.28
	pointsOfFunctionPlot[931].Y = 377.412

	pointsOfFunctionPlot[932].X = 4.29
	pointsOfFunctionPlot[932].Y = 382.681

	pointsOfFunctionPlot[933].X = 4.3
	pointsOfFunctionPlot[933].Y = 388.023

	pointsOfFunctionPlot[934].X = 4.31
	pointsOfFunctionPlot[934].Y = 393.44

	pointsOfFunctionPlot[935].X = 4.32
	pointsOfFunctionPlot[935].Y = 398.932

	pointsOfFunctionPlot[936].X = 4.33
	pointsOfFunctionPlot[936].Y = 404.501

	pointsOfFunctionPlot[937].X = 4.34
	pointsOfFunctionPlot[937].Y = 410.147

	pointsOfFunctionPlot[938].X = 4.35
	pointsOfFunctionPlot[938].Y = 415.873

	pointsOfFunctionPlot[939].X = 4.36
	pointsOfFunctionPlot[939].Y = 421.678

	pointsOfFunctionPlot[940].X = 4.37
	pointsOfFunctionPlot[940].Y = 427.565

	pointsOfFunctionPlot[941].X = 4.38
	pointsOfFunctionPlot[941].Y = 433.533

	pointsOfFunctionPlot[942].X = 4.39
	pointsOfFunctionPlot[942].Y = 439.585

	pointsOfFunctionPlot[943].X = 4.4
	pointsOfFunctionPlot[943].Y = 445.721

	pointsOfFunctionPlot[944].X = 4.41
	pointsOfFunctionPlot[944].Y = 451.943

	pointsOfFunctionPlot[945].X = 4.42
	pointsOfFunctionPlot[945].Y = 458.252

	pointsOfFunctionPlot[946].X = 4.43
	pointsOfFunctionPlot[946].Y = 464.649

	pointsOfFunctionPlot[947].X = 4.44
	pointsOfFunctionPlot[947].Y = 471.136

	pointsOfFunctionPlot[948].X = 4.45
	pointsOfFunctionPlot[948].Y = 477.712

	pointsOfFunctionPlot[949].X = 4.46
	pointsOfFunctionPlot[949].Y = 484.381

	pointsOfFunctionPlot[950].X = 4.47
	pointsOfFunctionPlot[950].Y = 491.143

	pointsOfFunctionPlot[951].X = 4.48
	pointsOfFunctionPlot[951].Y = 497.999

	pointsOfFunctionPlot[952].X = 4.49
	pointsOfFunctionPlot[952].Y = 504.951

	pointsOfFunctionPlot[953].X = 4.5
	pointsOfFunctionPlot[953].Y = 512.0

	pointsOfFunctionPlot[954].X = 4.51
	pointsOfFunctionPlot[954].Y = 519.147

	pointsOfFunctionPlot[955].X = 4.52
	pointsOfFunctionPlot[955].Y = 526.394

	pointsOfFunctionPlot[956].X = 4.53
	pointsOfFunctionPlot[956].Y = 533.742

	pointsOfFunctionPlot[957].X = 4.54
	pointsOfFunctionPlot[957].Y = 541.193

	pointsOfFunctionPlot[958].X = 4.55
	pointsOfFunctionPlot[958].Y = 548.748

	pointsOfFunctionPlot[959].X = 4.56
	pointsOfFunctionPlot[959].Y = 556.408

	pointsOfFunctionPlot[960].X = 4.57
	pointsOfFunctionPlot[960].Y = 564.175

	pointsOfFunctionPlot[961].X = 4.58
	pointsOfFunctionPlot[961].Y = 572.051

	pointsOfFunctionPlot[962].X = 4.59
	pointsOfFunctionPlot[962].Y = 580.036

	pointsOfFunctionPlot[963].X = 4.6
	pointsOfFunctionPlot[963].Y = 588.133

	pointsOfFunctionPlot[964].X = 4.61
	pointsOfFunctionPlot[964].Y = 596.343

	pointsOfFunctionPlot[965].X = 4.62
	pointsOfFunctionPlot[965].Y = 604.668

	pointsOfFunctionPlot[966].X = 4.63
	pointsOfFunctionPlot[966].Y = 613.109

	pointsOfFunctionPlot[967].X = 4.64
	pointsOfFunctionPlot[967].Y = 621.667

	pointsOfFunctionPlot[968].X = 4.65
	pointsOfFunctionPlot[968].Y = 630.345

	pointsOfFunctionPlot[969].X = 4.66
	pointsOfFunctionPlot[969].Y = 639.145

	pointsOfFunctionPlot[970].X = 4.67
	pointsOfFunctionPlot[970].Y = 648.067

	pointsOfFunctionPlot[971].X = 4.68
	pointsOfFunctionPlot[971].Y = 657.114

	pointsOfFunctionPlot[972].X = 4.69
	pointsOfFunctionPlot[972].Y = 666.287

	pointsOfFunctionPlot[973].X = 4.7
	pointsOfFunctionPlot[973].Y = 675.588

	pointsOfFunctionPlot[974].X = 4.71
	pointsOfFunctionPlot[974].Y = 685.018

	pointsOfFunctionPlot[975].X = 4.72
	pointsOfFunctionPlot[975].Y = 694.581

	pointsOfFunctionPlot[976].X = 4.73
	pointsOfFunctionPlot[976].Y = 704.277

	pointsOfFunctionPlot[977].X = 4.74
	pointsOfFunctionPlot[977].Y = 714.108

	pointsOfFunctionPlot[978].X = 4.75
	pointsOfFunctionPlot[978].Y = 724.077

	pointsOfFunctionPlot[979].X = 4.76
	pointsOfFunctionPlot[979].Y = 734.185

	pointsOfFunctionPlot[980].X = 4.77
	pointsOfFunctionPlot[980].Y = 744.433

	pointsOfFunctionPlot[981].X = 4.78
	pointsOfFunctionPlot[981].Y = 754.825

	pointsOfFunctionPlot[982].X = 4.79
	pointsOfFunctionPlot[982].Y = 765.362

	pointsOfFunctionPlot[983].X = 4.8
	pointsOfFunctionPlot[983].Y = 776.046

	pointsOfFunctionPlot[984].X = 4.81
	pointsOfFunctionPlot[984].Y = 776.046

	pointsOfFunctionPlot[985].X = 4.82
	pointsOfFunctionPlot[985].Y = 786.88

	pointsOfFunctionPlot[986].X = 4.83
	pointsOfFunctionPlot[986].Y = 797.864

	pointsOfFunctionPlot[987].X = 4.84
	pointsOfFunctionPlot[987].Y = 820.295

	pointsOfFunctionPlot[988].X = 4.85
	pointsOfFunctionPlot[988].Y = 831.746

	pointsOfFunctionPlot[989].X = 4.86
	pointsOfFunctionPlot[989].Y = 843.357

	pointsOfFunctionPlot[990].X = 4.87
	pointsOfFunctionPlot[990].Y = 855.13

	pointsOfFunctionPlot[991].X = 4.88
	pointsOfFunctionPlot[991].Y = 867.067

	pointsOfFunctionPlot[992].X = 4.89
	pointsOfFunctionPlot[992].Y = 879.171

	pointsOfFunctionPlot[993].X = 4.9
	pointsOfFunctionPlot[993].Y = 891.443

	pointsOfFunctionPlot[994].X = 4.91
	pointsOfFunctionPlot[994].Y = 903.887

	pointsOfFunctionPlot[995].X = 4.92
	pointsOfFunctionPlot[995].Y = 916.505

	pointsOfFunctionPlot[996].X = 4.93
	pointsOfFunctionPlot[996].Y = 929.299

	pointsOfFunctionPlot[997].X = 4.94
	pointsOfFunctionPlot[997].Y = 942.272

	pointsOfFunctionPlot[998].X = 4.95
	pointsOfFunctionPlot[998].Y = 955.425

	pointsOfFunctionPlot[999].X = 4.96
	pointsOfFunctionPlot[999].Y = 968.763

	pointsOfFunctionPlot[1_000].X = 4.97
	pointsOfFunctionPlot[1_000].Y = 982.286

	pointsOfFunctionPlot[1_001].X = 4.98
	pointsOfFunctionPlot[1_001].Y = 995.998

	pointsOfFunctionPlot[1_002].X = 4.99
	pointsOfFunctionPlot[1_002].Y = 1_009.902

	pointsOfFunctionPlot[1_003].X = 5.0
	pointsOfFunctionPlot[1_003].Y = 1_024.0

	pointsOfFunctionPlot[1_004].X = 5.01
	pointsOfFunctionPlot[1_004].Y = 1_038.294

	pointsOfFunctionPlot[1_005].X = 5.02
	pointsOfFunctionPlot[1_005].Y = 1_052.788

	pointsOfFunctionPlot[1_006].X = 5.03
	pointsOfFunctionPlot[1_006].Y = 1_067.484

	pointsOfFunctionPlot[1_007].X = 5.04
	pointsOfFunctionPlot[1_007].Y = 1_082.386

	pointsOfFunctionPlot[1_008].X = 5.05
	pointsOfFunctionPlot[1_008].Y = 1_097.496

	pointsOfFunctionPlot[1_009].X = 5.06
	pointsOfFunctionPlot[1_009].Y = 1_112.816

	pointsOfFunctionPlot[1_010].X = 5.07
	pointsOfFunctionPlot[1_010].Y = 1_128.35

	pointsOfFunctionPlot[1_011].X = 5.08
	pointsOfFunctionPlot[1_011].Y = 1_144.102

	pointsOfFunctionPlot[1_012].X = 5.09
	pointsOfFunctionPlot[1_012].Y = 1_160.073

	pointsOfFunctionPlot[1_013].X = 5.1
	pointsOfFunctionPlot[1_013].Y = 1_176.267

	pointsOfFunctionPlot[1_014].X = 5.11
	pointsOfFunctionPlot[1_014].Y = 1_192.687

	pointsOfFunctionPlot[1_015].X = 5.12
	pointsOfFunctionPlot[1_015].Y = 1_209.336

	pointsOfFunctionPlot[1_016].X = 5.13
	pointsOfFunctionPlot[1_016].Y = 1_226.218

	pointsOfFunctionPlot[1_017].X = 5.14
	pointsOfFunctionPlot[1_017].Y = 1_243.335

	pointsOfFunctionPlot[1_018].X = 5.15
	pointsOfFunctionPlot[1_018].Y = 1_260.691

	pointsOfFunctionPlot[1_019].X = 5.16
	pointsOfFunctionPlot[1_019].Y = 1_278.29

	pointsOfFunctionPlot[1_020].X = 5.17
	pointsOfFunctionPlot[1_020].Y = 1_296.134

	pointsOfFunctionPlot[1_021].X = 5.18
	pointsOfFunctionPlot[1_021].Y = 1_314.228

	pointsOfFunctionPlot[1_022].X = 5.19
	pointsOfFunctionPlot[1_022].Y = 1_332.574

	pointsOfFunctionPlot[1_023].X = 5.2
	pointsOfFunctionPlot[1_023].Y = 1_351.176

	pointsOfFunctionPlot[1_024].X = 5.21
	pointsOfFunctionPlot[1_024].Y = 1_370.037

	pointsOfFunctionPlot[1_025].X = 5.22
	pointsOfFunctionPlot[1_025].Y = 1_389.162

	pointsOfFunctionPlot[1_026].X = 5.23
	pointsOfFunctionPlot[1_026].Y = 1_408.554

	pointsOfFunctionPlot[1_027].X = 5.24
	pointsOfFunctionPlot[1_027].Y = 1_428.217

	pointsOfFunctionPlot[1_028].X = 5.25
	pointsOfFunctionPlot[1_028].Y = 1_448.154

	pointsOfFunctionPlot[1_029].X = 5.26
	pointsOfFunctionPlot[1_029].Y = 1_468.37

	pointsOfFunctionPlot[1_030].X = 5.27
	pointsOfFunctionPlot[1_030].Y = 1_488.867

	pointsOfFunctionPlot[1_031].X = 5.28
	pointsOfFunctionPlot[1_031].Y = 1_509.651

	pointsOfFunctionPlot[1_032].X = 5.29
	pointsOfFunctionPlot[1_032].Y = 1_530.725

	pointsOfFunctionPlot[1_033].X = 5.3
	pointsOfFunctionPlot[1_033].Y = 1_552.093

	pointsOfFunctionPlot[1_034].X = 5.31
	pointsOfFunctionPlot[1_034].Y = 1_573.76

	pointsOfFunctionPlot[1_035].X = 5.32
	pointsOfFunctionPlot[1_035].Y = 1_595.729

	pointsOfFunctionPlot[1_036].X = 5.33
	pointsOfFunctionPlot[1_036].Y = 1_618.004

	pointsOfFunctionPlot[1_037].X = 5.34
	pointsOfFunctionPlot[1_037].Y = 1_640.591

	pointsOfFunctionPlot[1_038].X = 5.35
	pointsOfFunctionPlot[1_038].Y = 1_663.492

	pointsOfFunctionPlot[1_039].X = 5.36
	pointsOfFunctionPlot[1_039].Y = 1_686.714

	pointsOfFunctionPlot[1_040].X = 5.37
	pointsOfFunctionPlot[1_040].Y = 1_710.26

	pointsOfFunctionPlot[1_041].X = 5.38
	pointsOfFunctionPlot[1_041].Y = 1_734.134

	pointsOfFunctionPlot[1_042].X = 5.39
	pointsOfFunctionPlot[1_042].Y = 1_758.342

	pointsOfFunctionPlot[1_043].X = 5.4
	pointsOfFunctionPlot[1_043].Y = 1_782.887

	pointsOfFunctionPlot[1_044].X = 5.41
	pointsOfFunctionPlot[1_044].Y = 1_807.775

	pointsOfFunctionPlot[1_045].X = 5.42
	pointsOfFunctionPlot[1_045].Y = 1_833.011

	pointsOfFunctionPlot[1_046].X = 5.43
	pointsOfFunctionPlot[1_046].Y = 1_858.599

	pointsOfFunctionPlot[1_047].X = 5.44
	pointsOfFunctionPlot[1_047].Y = 1_884.544

	pointsOfFunctionPlot[1_048].X = 5.45
	pointsOfFunctionPlot[1_048].Y = 1_910.851

	pointsOfFunctionPlot[1_049].X = 5.46
	pointsOfFunctionPlot[1_049].Y = 1_937.526

	pointsOfFunctionPlot[1_050].X = 5.47
	pointsOfFunctionPlot[1_050].Y = 1_964.572

	pointsOfFunctionPlot[1_051].X = 5.48
	pointsOfFunctionPlot[1_051].Y = 1_991.997

	pointsOfFunctionPlot[1_052].X = 5.49
	pointsOfFunctionPlot[1_052].Y = 2_019.804

	pointsOfFunctionPlot[1_053].X = 5.5
	pointsOfFunctionPlot[1_053].Y = 2_048.0

	pointsOfFunctionPlot[1_054].X = 5.51
	pointsOfFunctionPlot[1_054].Y = 2_076.589

	pointsOfFunctionPlot[1_055].X = 5.52
	pointsOfFunctionPlot[1_055].Y = 2_105.577

	pointsOfFunctionPlot[1_056].X = 5.53
	pointsOfFunctionPlot[1_056].Y = 2_134.969

	pointsOfFunctionPlot[1_057].X = 5.54
	pointsOfFunctionPlot[1_057].Y = 2_164.772

	pointsOfFunctionPlot[1_058].X = 5.55
	pointsOfFunctionPlot[1_058].Y = 2_194.992

	pointsOfFunctionPlot[1_059].X = 5.56
	pointsOfFunctionPlot[1_059].Y = 2_225.633

	pointsOfFunctionPlot[1_060].X = 5.57
	pointsOfFunctionPlot[1_060].Y = 2_256.701

	pointsOfFunctionPlot[1_061].X = 5.58
	pointsOfFunctionPlot[1_061].Y = 2_288.204

	pointsOfFunctionPlot[1_062].X = 5.59
	pointsOfFunctionPlot[1_062].Y = 2_320.146

	pointsOfFunctionPlot[1_063].X = 5.6
	pointsOfFunctionPlot[1_063].Y = 2_352.534

	pointsOfFunctionPlot[1_064].X = 5.61
	pointsOfFunctionPlot[1_064].Y = 2_385.374

	pointsOfFunctionPlot[1_065].X = 5.62
	pointsOfFunctionPlot[1_065].Y = 2_418.673

	pointsOfFunctionPlot[1_066].X = 5.63
	pointsOfFunctionPlot[1_066].Y = 2_452.346

	pointsOfFunctionPlot[1_067].X = 5.64
	pointsOfFunctionPlot[1_067].Y = 2_486.671

	pointsOfFunctionPlot[1_068].X = 5.65
	pointsOfFunctionPlot[1_068].Y = 2_521.383

	pointsOfFunctionPlot[1_069].X = 5.66
	pointsOfFunctionPlot[1_069].Y = 2_556.581

	pointsOfFunctionPlot[1_070].X = 5.67
	pointsOfFunctionPlot[1_070].Y = 2_592.269

	pointsOfFunctionPlot[1_071].X = 5.68
	pointsOfFunctionPlot[1_071].Y = 2_628.456

	pointsOfFunctionPlot[1_072].X = 5.69
	pointsOfFunctionPlot[1_072].Y = 2_665.148

	pointsOfFunctionPlot[1_073].X = 5.7
	pointsOfFunctionPlot[1_073].Y = 2_702.352

	pointsOfFunctionPlot[1_074].X = 5.71
	pointsOfFunctionPlot[1_074].Y = 2_740.075

	pointsOfFunctionPlot[1_075].X = 5.72
	pointsOfFunctionPlot[1_075].Y = 2_778.325

	pointsOfFunctionPlot[1_076].X = 5.73
	pointsOfFunctionPlot[1_076].Y = 2_817.109

	pointsOfFunctionPlot[1_077].X = 5.74
	pointsOfFunctionPlot[1_077].Y = 2_856.435

	pointsOfFunctionPlot[1_078].X = 5.75
	pointsOfFunctionPlot[1_078].Y = 2_896.309

	pointsOfFunctionPlot[1_079].X = 5.76
	pointsOfFunctionPlot[1_079].Y = 2_936.74

	pointsOfFunctionPlot[1_080].X = 5.77
	pointsOfFunctionPlot[1_080].Y = 2_977.735

	pointsOfFunctionPlot[1_081].X = 5.78
	pointsOfFunctionPlot[1_081].Y = 3_019.303

	pointsOfFunctionPlot[1_082].X = 5.79
	pointsOfFunctionPlot[1_082].Y = 3_061.451

	pointsOfFunctionPlot[1_083].X = 5.8
	pointsOfFunctionPlot[1_083].Y = 3_104.187

	pointsOfFunctionPlot[1_084].X = 5.81
	pointsOfFunctionPlot[1_084].Y = 3_147.52

	pointsOfFunctionPlot[1_085].X = 5.82
	pointsOfFunctionPlot[1_085].Y = 3_191

	pointsOfFunctionPlot[1_086].X = 5.83
	pointsOfFunctionPlot[1_086].Y = 3_236.009

	pointsOfFunctionPlot[1_087].X = 5.84
	pointsOfFunctionPlot[1_087].Y = 3_281.182

	pointsOfFunctionPlot[1_088].X = 5.85
	pointsOfFunctionPlot[1_088].Y = 3_326.985

	pointsOfFunctionPlot[1_089].X = 5.86
	pointsOfFunctionPlot[1_089].Y = 3_373.428

	pointsOfFunctionPlot[1_090].X = 5.87
	pointsOfFunctionPlot[1_090].Y = 3_420.52

	pointsOfFunctionPlot[1_091].X = 5.88
	pointsOfFunctionPlot[1_091].Y = 3_468.268

	pointsOfFunctionPlot[1_092].X = 5.89
	pointsOfFunctionPlot[1_092].Y = 3_516.684

	pointsOfFunctionPlot[1_093].X = 5.9
	pointsOfFunctionPlot[1_093].Y = 3_565.775

	pointsOfFunctionPlot[1_094].X = 5.91
	pointsOfFunctionPlot[1_094].Y = 3_615.551

	pointsOfFunctionPlot[1_095].X = 5.92
	pointsOfFunctionPlot[1_095].Y = 3_666.022

	pointsOfFunctionPlot[1_096].X = 5.93
	pointsOfFunctionPlot[1_096].Y = 3_717.198

	pointsOfFunctionPlot[1_097].X = 5.94
	pointsOfFunctionPlot[1_097].Y = 3_769.088

	pointsOfFunctionPlot[1_098].X = 5.95
	pointsOfFunctionPlot[1_098].Y = 3_821.703

	pointsOfFunctionPlot[1_099].X = 5.96
	pointsOfFunctionPlot[1_099].Y = 3_875.052

	pointsOfFunctionPlot[1_100].X = 5.97
	pointsOfFunctionPlot[1_100].Y = 3_929.145

	pointsOfFunctionPlot[1_101].X = 5.98
	pointsOfFunctionPlot[1_101].Y = 3_983.994

	pointsOfFunctionPlot[1_102].X = 5.99
	pointsOfFunctionPlot[1_102].Y = 4_039.609

	pointsOfFunctionPlot[1_103].X = 6.0
	pointsOfFunctionPlot[1_103].Y = 4096.0

	pointsOfFunctionPlot[1_104].X = 6.01
	pointsOfFunctionPlot[1_104].Y = 4_153.178

	pointsOfFunctionPlot[1_105].X = 6.02
	pointsOfFunctionPlot[1_105].Y = 4_211.154

	pointsOfFunctionPlot[1_106].X = 6.03
	pointsOfFunctionPlot[1_106].Y = 4_269.939

	pointsOfFunctionPlot[1_107].X = 6.04
	pointsOfFunctionPlot[1_107].Y = 4_329.545

	pointsOfFunctionPlot[1_108].X = 6.05
	pointsOfFunctionPlot[1_108].Y = 4_389.984

	pointsOfFunctionPlot[1_109].X = 6.06
	pointsOfFunctionPlot[1_109].Y = 4_451.266

	pointsOfFunctionPlot[1_110].X = 6.07
	pointsOfFunctionPlot[1_110].Y = 4_513.403

	pointsOfFunctionPlot[1_111].X = 6.08
	pointsOfFunctionPlot[1_111].Y = 4_576.408

	pointsOfFunctionPlot[1_112].X = 6.09
	pointsOfFunctionPlot[1_112].Y = 4_640.292

	pointsOfFunctionPlot[1_113].X = 6.1
	pointsOfFunctionPlot[1_113].Y = 4_705.068

	pointsOfFunctionPlot[1_114].X = 6.11
	pointsOfFunctionPlot[1_114].Y = 4_770.748

	pointsOfFunctionPlot[1_115].X = 6.12
	pointsOfFunctionPlot[1_115].Y = 4_837.345

	pointsOfFunctionPlot[1_116].X = 6.13
	pointsOfFunctionPlot[1_116].Y = 4_904.872

	pointsOfFunctionPlot[1_117].X = 6.14
	pointsOfFunctionPlot[1_117].Y = 4_973.342

	pointsOfFunctionPlot[1_118].X = 6.15
	pointsOfFunctionPlot[1_118].Y = 5_042.767

	pointsOfFunctionPlot[1_119].X = 6.16
	pointsOfFunctionPlot[1_119].Y = 5_113.161

	pointsOfFunctionPlot[1_120].X = 6.17
	pointsOfFunctionPlot[1_120].Y = 5_184.539

	pointsOfFunctionPlot[1_121].X = 6.18
	pointsOfFunctionPlot[1_121].Y = 5_256.912

	pointsOfFunctionPlot[1_122].X = 6.19
	pointsOfFunctionPlot[1_122].Y = 5_330.296

	pointsOfFunctionPlot[1_123].X = 6.2
	pointsOfFunctionPlot[1_123].Y = 5_404.704

	pointsOfFunctionPlot[1_124].X = 6.21
	pointsOfFunctionPlot[1_124].Y = 5_480.151

	pointsOfFunctionPlot[1_125].X = 6.22
	pointsOfFunctionPlot[1_125].Y = 5_556.651

	pointsOfFunctionPlot[1_126].X = 6.23
	pointsOfFunctionPlot[1_126].Y = 5_634.219

	pointsOfFunctionPlot[1_127].X = 6.24
	pointsOfFunctionPlot[1_127].Y = 5_712.87

	pointsOfFunctionPlot[1_128].X = 6.25
	pointsOfFunctionPlot[1_128].Y = 5_792.618

	pointsOfFunctionPlot[1_129].X = 6.26
	pointsOfFunctionPlot[1_129].Y = 5_873.48

	pointsOfFunctionPlot[1_130].X = 6.27
	pointsOfFunctionPlot[1_130].Y = 5_955.471

	pointsOfFunctionPlot[1_131].X = 6.28
	pointsOfFunctionPlot[1_131].Y = 6_038.606

	pointsOfFunctionPlot[1_132].X = 6.29
	pointsOfFunctionPlot[1_132].Y = 6_122.902

	pointsOfFunctionPlot[1_133].X = 6.3
	pointsOfFunctionPlot[1_133].Y = 6_208.375

	pointsOfFunctionPlot[1_134].X = 6.31
	pointsOfFunctionPlot[1_134].Y = 6_208.375

	pointsOfFunctionPlot[1_135].X = 6.32
	pointsOfFunctionPlot[1_135].Y = 6_382.916

	pointsOfFunctionPlot[1_136].X = 6.33
	pointsOfFunctionPlot[1_136].Y = 6_472.018

	pointsOfFunctionPlot[1_137].X = 6.34
	pointsOfFunctionPlot[1_137].Y = 6_562.364

	pointsOfFunctionPlot[1_138].X = 6.35
	pointsOfFunctionPlot[1_138].Y = 6_653.971

	pointsOfFunctionPlot[1_139].X = 6.36
	pointsOfFunctionPlot[1_139].Y = 6_746.857

	pointsOfFunctionPlot[1_140].X = 6.37
	pointsOfFunctionPlot[1_140].Y = 6_841.02

	pointsOfFunctionPlot[1_141].X = 6.38
	pointsOfFunctionPlot[1_141].Y = 6_936.537

	pointsOfFunctionPlot[1_142].X = 6.39
	pointsOfFunctionPlot[1_142].Y = 7_033.368

	pointsOfFunctionPlot[1_143].X = 6.4
	pointsOfFunctionPlot[1_143].Y = 7_131.55

	pointsOfFunctionPlot[1_144].X = 6.41
	pointsOfFunctionPlot[1_144].Y = 7_231.102

	pointsOfFunctionPlot[1_145].X = 6.42
	pointsOfFunctionPlot[1_145].Y = 7_332.045

	pointsOfFunctionPlot[1_146].X = 6.43
	pointsOfFunctionPlot[1_146].Y = 7_434.396

	pointsOfFunctionPlot[1_147].X = 6.44
	pointsOfFunctionPlot[1_147].Y = 7_538.177

	pointsOfFunctionPlot[1_148].X = 6.45
	pointsOfFunctionPlot[1_148].Y = 7_643.406

	pointsOfFunctionPlot[1_149].X = 6.46
	pointsOfFunctionPlot[1_149].Y = 7_750.104

	pointsOfFunctionPlot[1_150].X = 6.47
	pointsOfFunctionPlot[1_150].Y = 7_858.291

	pointsOfFunctionPlot[1_151].X = 6.48
	pointsOfFunctionPlot[1_151].Y = 7_967.989

	pointsOfFunctionPlot[1_152].X = 6.49
	pointsOfFunctionPlot[1_152].Y = 8_079.218

	pointsOfFunctionPlot[1_153].X = 6.5
	pointsOfFunctionPlot[1_153].Y = 8_192.0

	pointsOfFunctionPlot[1_154].X = 6.51
	pointsOfFunctionPlot[1_154].Y = 8_306.356

	pointsOfFunctionPlot[1_155].X = 6.52
	pointsOfFunctionPlot[1_155].Y = 8_422.308

	pointsOfFunctionPlot[1_156].X = 6.53
	pointsOfFunctionPlot[1_156].Y = 8_539.879

	pointsOfFunctionPlot[1_157].X = 6.54
	pointsOfFunctionPlot[1_157].Y = 8_659.091

	pointsOfFunctionPlot[1_158].X = 6.55
	pointsOfFunctionPlot[1_158].Y = 8_779.968

	pointsOfFunctionPlot[1_159].X = 6.56
	pointsOfFunctionPlot[1_159].Y = 8_902.532

	pointsOfFunctionPlot[1_160].X = 6.57
	pointsOfFunctionPlot[1_160].Y = 9_026.806

	pointsOfFunctionPlot[1_161].X = 6.58
	pointsOfFunctionPlot[1_161].Y = 9_152.816

	pointsOfFunctionPlot[1_162].X = 6.59
	pointsOfFunctionPlot[1_162].Y = 9_280.584

	pointsOfFunctionPlot[1_163].X = 6.6
	pointsOfFunctionPlot[1_163].Y = 9_410.136

	pointsOfFunctionPlot[1_164].X = 6.61
	pointsOfFunctionPlot[1_164].Y = 9_541.497

	pointsOfFunctionPlot[1_165].X = 6.62
	pointsOfFunctionPlot[1_165].Y = 9_674.691

	pointsOfFunctionPlot[1_166].X = 6.63
	pointsOfFunctionPlot[1_166].Y = 9_809.745

	pointsOfFunctionPlot[1_167].X = 6.64
	pointsOfFunctionPlot[1_167].Y = 9_946.684

	pointsOfFunctionPlot[1_168].X = 6.65
	pointsOfFunctionPlot[1_168].Y = 10_085.535

	pointsOfFunctionPlot[1_169].X = 6.66
	pointsOfFunctionPlot[1_169].Y = 10_226.323

	pointsOfFunctionPlot[1_170].X = 6.67
	pointsOfFunctionPlot[1_170].Y = 10_369.078

	pointsOfFunctionPlot[1_171].X = 6.68
	pointsOfFunctionPlot[1_171].Y = 10_513.825

	pointsOfFunctionPlot[1_172].X = 6.69
	pointsOfFunctionPlot[1_172].Y = 10_660.592

	pointsOfFunctionPlot[1_173].X = 6.7
	pointsOfFunctionPlot[1_173].Y = 10_809.408

	pointsOfFunctionPlot[1_174].X = 6.71
	pointsOfFunctionPlot[1_174].Y = 10_960.302

	pointsOfFunctionPlot[1_175].X = 6.72
	pointsOfFunctionPlot[1_175].Y = 11_113.302

	pointsOfFunctionPlot[1_176].X = 6.73
	pointsOfFunctionPlot[1_176].Y = 11_268.438

	pointsOfFunctionPlot[1_177].X = 6.74
	pointsOfFunctionPlot[1_177].Y = 11_425.74

	pointsOfFunctionPlot[1_178].X = 6.75
	pointsOfFunctionPlot[1_178].Y = 11_585.237

	pointsOfFunctionPlot[1_179].X = 6.76
	pointsOfFunctionPlot[1_179].Y = 11_746.961

	pointsOfFunctionPlot[1_180].X = 6.77
	pointsOfFunctionPlot[1_180].Y = 11_910.942

	pointsOfFunctionPlot[1_181].X = 6.78
	pointsOfFunctionPlot[1_181].Y = 12_077.213

	pointsOfFunctionPlot[1_182].X = 6.79
	pointsOfFunctionPlot[1_182].Y = 12_245.805

	pointsOfFunctionPlot[1_183].X = 6.8
	pointsOfFunctionPlot[1_183].Y = 12_416.75

	pointsOfFunctionPlot[1_184].X = 6.81
	pointsOfFunctionPlot[1_184].Y = 12_590.081

	pointsOfFunctionPlot[1_185].X = 6.82
	pointsOfFunctionPlot[1_185].Y = 12_765.832

	pointsOfFunctionPlot[1_186].X = 6.83
	pointsOfFunctionPlot[1_186].Y = 12_944.036

	pointsOfFunctionPlot[1_187].X = 6.84
	pointsOfFunctionPlot[1_187].Y = 13_124.728

	pointsOfFunctionPlot[1_188].X = 6.85
	pointsOfFunctionPlot[1_188].Y = 13_307.943

	pointsOfFunctionPlot[1_189].X = 6.86
	pointsOfFunctionPlot[1_189].Y = 13_493.715

	pointsOfFunctionPlot[1_190].X = 6.87
	pointsOfFunctionPlot[1_190].Y = 13_682.08

	pointsOfFunctionPlot[1_191].X = 6.88
	pointsOfFunctionPlot[1_191].Y = 13_873.075

	pointsOfFunctionPlot[1_192].X = 6.89
	pointsOfFunctionPlot[1_192].Y = 14_066.736

	pointsOfFunctionPlot[1_193].X = 6.9
	pointsOfFunctionPlot[1_193].Y = 14_263.1

	pointsOfFunctionPlot[1_194].X = 6.91
	pointsOfFunctionPlot[1_194].Y = 14_462.205

	pointsOfFunctionPlot[1_195].X = 6.92
	pointsOfFunctionPlot[1_195].Y = 14_664.09

	pointsOfFunctionPlot[1_196].X = 6.93
	pointsOfFunctionPlot[1_196].Y = 14_868.793

	pointsOfFunctionPlot[1_197].X = 6.94
	pointsOfFunctionPlot[1_197].Y = 15_076.354

	pointsOfFunctionPlot[1_198].X = 6.95
	pointsOfFunctionPlot[1_198].Y = 15_286.812

	pointsOfFunctionPlot[1_199].X = 6.96
	pointsOfFunctionPlot[1_199].Y = 15_500.208

	pointsOfFunctionPlot[1_200].X = 6.97
	pointsOfFunctionPlot[1_200].Y = 15_716.583

	pointsOfFunctionPlot[1_201].X = 6.98
	pointsOfFunctionPlot[1_201].Y = 15_935.978

	pointsOfFunctionPlot[1_202].X = 6.99
	pointsOfFunctionPlot[1_202].Y = 16_158.436

	pointsOfFunctionPlot[1_203].X = 7.0
	pointsOfFunctionPlot[1_203].Y = 16_384.0

	pointsOfFunctionPlot[1_204].X = 7.01
	pointsOfFunctionPlot[1_204].Y = 16_612.712

	pointsOfFunctionPlot[1_205].X = 7.02
	pointsOfFunctionPlot[1_205].Y = 16_844.616

	pointsOfFunctionPlot[1_206].X = 7.03
	pointsOfFunctionPlot[1_206].Y = 17_079.759

	pointsOfFunctionPlot[1_207].X = 7.04
	pointsOfFunctionPlot[1_207].Y = 17_318.183

	pointsOfFunctionPlot[1_208].X = 7.05
	pointsOfFunctionPlot[1_208].Y = 17_559.936

	pointsOfFunctionPlot[1_209].X = 7.06
	pointsOfFunctionPlot[1_209].Y = 17_805.064

	pointsOfFunctionPlot[1_210].X = 7.07
	pointsOfFunctionPlot[1_210].Y = 18_053.613

	pointsOfFunctionPlot[1_211].X = 7.08
	pointsOfFunctionPlot[1_211].Y = 18_305.632

	pointsOfFunctionPlot[1_212].X = 7.09
	pointsOfFunctionPlot[1_212].Y = 18_561.169

	pointsOfFunctionPlot[1_213].X = 7.1
	pointsOfFunctionPlot[1_213].Y = 18_820.273

	pointsOfFunctionPlot[1_214].X = 7.11
	pointsOfFunctionPlot[1_214].Y = 19_082.995

	pointsOfFunctionPlot[1_215].X = 7.12
	pointsOfFunctionPlot[1_215].Y = 19_349.383

	pointsOfFunctionPlot[1_216].X = 7.13
	pointsOfFunctionPlot[1_216].Y = 19_619.491

	pointsOfFunctionPlot[1_217].X = 7.14
	pointsOfFunctionPlot[1_217].Y = 19_893.369

	pointsOfFunctionPlot[1_218].X = 7.15
	pointsOfFunctionPlot[1_218].Y = 20_171.07

	pointsOfFunctionPlot[1_219].X = 7.16
	pointsOfFunctionPlot[1_219].Y = 20_452.647

	pointsOfFunctionPlot[1_220].X = 7.17
	pointsOfFunctionPlot[1_220].Y = 20_738.156

	pointsOfFunctionPlot[1_221].X = 7.18
	pointsOfFunctionPlot[1_221].Y = 21_027.649

	pointsOfFunctionPlot[1_222].X = 7.19
	pointsOfFunctionPlot[1_222].Y = 21_321.185

	pointsOfFunctionPlot[1_223].X = 7.2
	pointsOfFunctionPlot[1_223].Y = 21_618.817

	pointsOfFunctionPlot[1_224].X = 7.21
	pointsOfFunctionPlot[1_224].Y = 21_920.605

	pointsOfFunctionPlot[1_225].X = 7.22
	pointsOfFunctionPlot[1_225].Y = 22_226.605

	pointsOfFunctionPlot[1_226].X = 7.23
	pointsOfFunctionPlot[1_226].Y = 22_536.877

	pointsOfFunctionPlot[1_227].X = 7.24
	pointsOfFunctionPlot[1_227].Y = 22_851.48

	pointsOfFunctionPlot[1_228].X = 7.25
	pointsOfFunctionPlot[1_228].Y = 23_170.475

	pointsOfFunctionPlot[1_229].X = 7.26
	pointsOfFunctionPlot[1_229].Y = 23_493.922

	pointsOfFunctionPlot[1_230].X = 7.27
	pointsOfFunctionPlot[1_230].Y = 23_821.885

	pointsOfFunctionPlot[1_231].X = 7.28
	pointsOfFunctionPlot[1_231].Y = 24_154.426

	pointsOfFunctionPlot[1_232].X = 7.29
	pointsOfFunctionPlot[1_232].Y = 24_491.61

	pointsOfFunctionPlot[1_233].X = 7.3
	pointsOfFunctionPlot[1_233].Y = 24_833.5

	pointsOfFunctionPlot[1_234].X = 7.31
	pointsOfFunctionPlot[1_234].Y = 25_180.163

	pointsOfFunctionPlot[1_235].X = 7.32
	pointsOfFunctionPlot[1_235].Y = 25_531.664

	pointsOfFunctionPlot[1_236].X = 7.33
	pointsOfFunctionPlot[1_236].Y = 25_888.073

	pointsOfFunctionPlot[1_237].X = 7.34
	pointsOfFunctionPlot[1_237].Y = 26_249.457

	pointsOfFunctionPlot[1_238].X = 7.35
	pointsOfFunctionPlot[1_238].Y = 26_615.886

	pointsOfFunctionPlot[1_239].X = 7.36
	pointsOfFunctionPlot[1_239].Y = 26_987.43

	pointsOfFunctionPlot[1_240].X = 7.37
	pointsOfFunctionPlot[1_240].Y = 27_364.16

	pointsOfFunctionPlot[1_241].X = 7.38
	pointsOfFunctionPlot[1_241].Y = 27_746.15

	pointsOfFunctionPlot[1_242].X = 7.39
	pointsOfFunctionPlot[1_242].Y = 28_133.472

	pointsOfFunctionPlot[1_243].X = 7.4
	pointsOfFunctionPlot[1_243].Y = 28_526.2

	pointsOfFunctionPlot[1_244].X = 7.41
	pointsOfFunctionPlot[1_244].Y = 28_526.2

	pointsOfFunctionPlot[1_245].X = 7.42
	pointsOfFunctionPlot[1_245].Y = 28_924.411

	pointsOfFunctionPlot[1_246].X = 7.43
	pointsOfFunctionPlot[1_246].Y = 29_737.587

	pointsOfFunctionPlot[1_247].X = 7.44
	pointsOfFunctionPlot[1_247].Y = 30_152.708

	pointsOfFunctionPlot[1_248].X = 7.45
	pointsOfFunctionPlot[1_248].Y = 30_573.625

	pointsOfFunctionPlot[1_249].X = 7.46
	pointsOfFunctionPlot[1_249].Y = 31_000.417

	pointsOfFunctionPlot[1_250].X = 7.47
	pointsOfFunctionPlot[1_250].Y = 31_433.166

	pointsOfFunctionPlot[1_251].X = 7.48
	pointsOfFunctionPlot[1_251].Y = 31_871.957

	pointsOfFunctionPlot[1_252].X = 7.49
	pointsOfFunctionPlot[1_252].Y = 32_316.873

	pointsOfFunctionPlot[1_253].X = 7.5
	pointsOfFunctionPlot[1_253].Y = 32_768.0

	pointsOfFunctionPlot[1_254].X = 7.51
	pointsOfFunctionPlot[1_254].Y = 33_225.424

	pointsOfFunctionPlot[1_255].X = 7.52
	pointsOfFunctionPlot[1_255].Y = 33_689.233

	pointsOfFunctionPlot[1_256].X = 7.53
	pointsOfFunctionPlot[1_256].Y = 34_159.518

	pointsOfFunctionPlot[1_257].X = 7.54
	pointsOfFunctionPlot[1_257].Y = 34_636.367

	pointsOfFunctionPlot[1_258].X = 7.55
	pointsOfFunctionPlot[1_258].Y = 35_119.872

	pointsOfFunctionPlot[1_259].X = 7.56
	pointsOfFunctionPlot[1_259].Y = 35_610.128

	pointsOfFunctionPlot[1_260].X = 7.57
	pointsOfFunctionPlot[1_260].Y = 36_107.226

	pointsOfFunctionPlot[1_261].X = 7.58
	pointsOfFunctionPlot[1_261].Y = 36_611.264

	pointsOfFunctionPlot[1_262].X = 7.59
	pointsOfFunctionPlot[1_262].Y = 37_122.339

	pointsOfFunctionPlot[1_263].X = 7.6
	pointsOfFunctionPlot[1_263].Y = 37_640.547

	pointsOfFunctionPlot[1_264].X = 7.61
	pointsOfFunctionPlot[1_264].Y = 38_165.99

	pointsOfFunctionPlot[1_265].X = 7.62
	pointsOfFunctionPlot[1_265].Y = 38_698.767

	pointsOfFunctionPlot[1_266].X = 7.63
	pointsOfFunctionPlot[1_266].Y = 39_238.982

	pointsOfFunctionPlot[1_267].X = 7.64
	pointsOfFunctionPlot[1_267].Y = 39_786.738

	pointsOfFunctionPlot[1_268].X = 7.65
	pointsOfFunctionPlot[1_268].Y = 40_342.14

	pointsOfFunctionPlot[1_269].X = 7.66
	pointsOfFunctionPlot[1_269].Y = 40_905.295

	pointsOfFunctionPlot[1_270].X = 7.67
	pointsOfFunctionPlot[1_270].Y = 41_476.312

	pointsOfFunctionPlot[1_271].X = 7.68
	pointsOfFunctionPlot[1_271].Y = 42_055.299

	pointsOfFunctionPlot[1_272].X = 7.69
	pointsOfFunctionPlot[1_272].Y = 42_642.369

	pointsOfFunctionPlot[1_273].X = 7.7
	pointsOfFunctionPlot[1_273].Y = 43_237.635

	pointsOfFunctionPlot[1_274].X = 7.71
	pointsOfFunctionPlot[1_274].Y = 43_841.21

	pointsOfFunctionPlot[1_275].X = 7.72
	pointsOfFunctionPlot[1_275].Y = 44_453.21

	pointsOfFunctionPlot[1_276].X = 7.73
	pointsOfFunctionPlot[1_276].Y = 45_073.754

	pointsOfFunctionPlot[1_277].X = 7.74
	pointsOfFunctionPlot[1_277].Y = 45_073.754

	pointsOfFunctionPlot[1_278].X = 7.75
	pointsOfFunctionPlot[1_278].Y = 46_340.95

	pointsOfFunctionPlot[1_279].X = 7.76
	pointsOfFunctionPlot[1_279].Y = 46_987.845

	pointsOfFunctionPlot[1_280].X = 7.77
	pointsOfFunctionPlot[1_280].Y = 47_643.771

	pointsOfFunctionPlot[1_281].X = 7.78
	pointsOfFunctionPlot[1_281].Y = 48_308.853

	pointsOfFunctionPlot[1_282].X = 7.79
	pointsOfFunctionPlot[1_282].Y = 48_983.22

	pointsOfFunctionPlot[1_283].X = 7.8
	pointsOfFunctionPlot[1_283].Y = 49_667.0

	pointsOfFunctionPlot[1_284].X = 7.81
	pointsOfFunctionPlot[1_284].Y = 50_360.325

	pointsOfFunctionPlot[1_285].X = 7.82
	pointsOfFunctionPlot[1_285].Y = 51_063.329

	pointsOfFunctionPlot[1_286].X = 7.83
	pointsOfFunctionPlot[1_286].Y = 51_776.147

	pointsOfFunctionPlot[1_287].X = 7.84
	pointsOfFunctionPlot[1_287].Y = 52_498.915

	pointsOfFunctionPlot[1_288].X = 7.85
	pointsOfFunctionPlot[1_288].Y = 53_231.773

	pointsOfFunctionPlot[1_289].X = 7.86
	pointsOfFunctionPlot[1_289].Y = 53_974.86

	pointsOfFunctionPlot[1_290].X = 7.87
	pointsOfFunctionPlot[1_290].Y = 54_728.321

	pointsOfFunctionPlot[1_291].X = 7.88
	pointsOfFunctionPlot[1_291].Y = 55_492.3

	pointsOfFunctionPlot[1_292].X = 7.89
	pointsOfFunctionPlot[1_292].Y = 56_266.944

	pointsOfFunctionPlot[1_293].X = 7.9
	pointsOfFunctionPlot[1_293].Y = 57_052.401

	pointsOfFunctionPlot[1_294].X = 7.91
	pointsOfFunctionPlot[1_294].Y = 57_848.823

	pointsOfFunctionPlot[1_295].X = 7.92
	pointsOfFunctionPlot[1_295].Y = 58_656.363

	pointsOfFunctionPlot[1_296].X = 7.93
	pointsOfFunctionPlot[1_296].Y = 59_475.175

	pointsOfFunctionPlot[1_297].X = 7.94
	pointsOfFunctionPlot[1_297].Y = 60_305.417

	pointsOfFunctionPlot[1_298].X = 7.95
	pointsOfFunctionPlot[1_298].Y = 61_146.25

	pointsOfFunctionPlot[1_299].X = 7.96
	pointsOfFunctionPlot[1_299].Y = 62_000.833

	pointsOfFunctionPlot[1_300].X = 7.97
	pointsOfFunctionPlot[1_300].Y = 62_866.333

	pointsOfFunctionPlot[1_301].X = 7.98
	pointsOfFunctionPlot[1_301].Y = 63_743.914

	pointsOfFunctionPlot[1_302].X = 7.99
	pointsOfFunctionPlot[1_302].Y = 64_633.746

	pointsOfFunctionPlot[1_303].X = 8.0
	pointsOfFunctionPlot[1_303].Y = 65_536.0

	pointsOfFunctionPlot[1_304].X = 8.01
	pointsOfFunctionPlot[1_304].Y = 66_450.848

	pointsOfFunctionPlot[1_305].X = 8.02
	pointsOfFunctionPlot[1_305].Y = 67_378.467

	pointsOfFunctionPlot[1_306].X = 8.03
	pointsOfFunctionPlot[1_306].Y = 68_319.036

	pointsOfFunctionPlot[1_307].X = 8.04
	pointsOfFunctionPlot[1_307].Y = 69_272.734

	pointsOfFunctionPlot[1_308].X = 8.05
	pointsOfFunctionPlot[1_308].Y = 70_239.745

	pointsOfFunctionPlot[1_309].X = 8.06
	pointsOfFunctionPlot[1_309].Y = 71_220.256

	pointsOfFunctionPlot[1_310].X = 8.07
	pointsOfFunctionPlot[1_310].Y = 72_214.453

	pointsOfFunctionPlot[1_311].X = 8.08
	pointsOfFunctionPlot[1_311].Y = 73_222.529

	pointsOfFunctionPlot[1_312].X = 8.09
	pointsOfFunctionPlot[1_312].Y = 74_244.678

	pointsOfFunctionPlot[1_313].X = 8.1
	pointsOfFunctionPlot[1_313].Y = 75_281.095

	pointsOfFunctionPlot[1_314].X = 8.11
	pointsOfFunctionPlot[1_314].Y = 76_331.98

	pointsOfFunctionPlot[1_315].X = 8.12
	pointsOfFunctionPlot[1_315].Y = 77_397.535

	pointsOfFunctionPlot[1_316].X = 8.13
	pointsOfFunctionPlot[1_316].Y = 78_477.964

	pointsOfFunctionPlot[1_317].X = 8.14
	pointsOfFunctionPlot[1_317].Y = 79_573.475

	pointsOfFunctionPlot[1_318].X = 8.15
	pointsOfFunctionPlot[1_318].Y = 80_684.29

	pointsOfFunctionPlot[1_319].X = 8.16
	pointsOfFunctionPlot[1_319].Y = 81_810.59

	pointsOfFunctionPlot[1_320].X = 8.17
	pointsOfFunctionPlot[1_320].Y = 82_952.624

	pointsOfFunctionPlot[1_321].X = 8.18
	pointsOfFunctionPlot[1_321].Y = 84_110.599

	pointsOfFunctionPlot[1_322].X = 8.19
	pointsOfFunctionPlot[1_322].Y = 85_284.739

	pointsOfFunctionPlot[1_323].X = 8.2
	pointsOfFunctionPlot[1_323].Y = 86_475.27

	pointsOfFunctionPlot[1_324].X = 8.21
	pointsOfFunctionPlot[1_324].Y = 87_682.42

	pointsOfFunctionPlot[1_325].X = 8.22
	pointsOfFunctionPlot[1_325].Y = 88_906.421

	pointsOfFunctionPlot[1_326].X = 8.23
	pointsOfFunctionPlot[1_326].Y = 90_147.508

	pointsOfFunctionPlot[1_327].X = 8.24
	pointsOfFunctionPlot[1_327].Y = 91_405.92

	pointsOfFunctionPlot[1_328].X = 8.25
	pointsOfFunctionPlot[1_328].Y = 92_681.9

	pointsOfFunctionPlot[1_329].X = 8.26
	pointsOfFunctionPlot[1_329].Y = 93_975.691

	pointsOfFunctionPlot[1_330].X = 8.27
	pointsOfFunctionPlot[1_330].Y = 95_287.542

	pointsOfFunctionPlot[1_331].X = 8.28
	pointsOfFunctionPlot[1_331].Y = 96_617.707

	pointsOfFunctionPlot[1_332].X = 8.29
	pointsOfFunctionPlot[1_332].Y = 97_966.44

	pointsOfFunctionPlot[1_333].X = 8.3
	pointsOfFunctionPlot[1_333].Y = 99_334.0

	pointsOfFunctionPlot[1_334].X = 8.31
	pointsOfFunctionPlot[1_334].Y = 100_720.651

	pointsOfFunctionPlot[1_335].X = 8.32
	pointsOfFunctionPlot[1_335].Y = 102_126.659

	pointsOfFunctionPlot[1_336].X = 8.33
	pointsOfFunctionPlot[1_336].Y = 103_552.294

	pointsOfFunctionPlot[1_337].X = 8.34
	pointsOfFunctionPlot[1_337].Y = 104_997.831

	pointsOfFunctionPlot[1_338].X = 8.35
	pointsOfFunctionPlot[1_338].Y = 106_463.546

	pointsOfFunctionPlot[1_339].X = 8.36
	pointsOfFunctionPlot[1_339].Y = 107_949.721

	pointsOfFunctionPlot[1_340].X = 8.37
	pointsOfFunctionPlot[1_340].Y = 109_456.643

	pointsOfFunctionPlot[1_341].X = 8.38
	pointsOfFunctionPlot[1_341].Y = 110_984.601

	pointsOfFunctionPlot[1_342].X = 8.39
	pointsOfFunctionPlot[1_342].Y = 112_533.888

	pointsOfFunctionPlot[1_343].X = 8.4
	pointsOfFunctionPlot[1_343].Y = 114_104.803

	pointsOfFunctionPlot[1_344].X = 8.41
	pointsOfFunctionPlot[1_344].Y = 115_697.647

	pointsOfFunctionPlot[1_345].X = 8.42
	pointsOfFunctionPlot[1_345].Y = 117_312.726

	pointsOfFunctionPlot[1_346].X = 8.43
	pointsOfFunctionPlot[1_346].Y = 118_950.35

	pointsOfFunctionPlot[1_347].X = 8.44
	pointsOfFunctionPlot[1_347].Y = 120_610.835

	pointsOfFunctionPlot[1_348].X = 8.45
	pointsOfFunctionPlot[1_348].Y = 122_294.5

	pointsOfFunctionPlot[1_349].X = 8.46
	pointsOfFunctionPlot[1_349].Y = 124_001.667

	pointsOfFunctionPlot[1_350].X = 8.47
	pointsOfFunctionPlot[1_350].Y = 125_732.666

	pointsOfFunctionPlot[1_351].X = 8.48
	pointsOfFunctionPlot[1_351].Y = 127_487.829

	pointsOfFunctionPlot[1_352].X = 8.49
	pointsOfFunctionPlot[1_352].Y = 129_267.493

	pointsOfFunctionPlot[1_353].X = 8.5
	pointsOfFunctionPlot[1_353].Y = 131_072.0

	pointsOfFunctionPlot[1_354].X = 8.51
	pointsOfFunctionPlot[1_354].Y = 132_901.696

	pointsOfFunctionPlot[1_355].X = 8.52
	pointsOfFunctionPlot[1_355].Y = 134_756.935

	pointsOfFunctionPlot[1_356].X = 8.53
	pointsOfFunctionPlot[1_356].Y = 136_638.072

	pointsOfFunctionPlot[1_357].X = 8.54
	pointsOfFunctionPlot[1_357].Y = 138_545.468

	pointsOfFunctionPlot[1_358].X = 8.55
	pointsOfFunctionPlot[1_358].Y = 140_479.491

	pointsOfFunctionPlot[1_359].X = 8.56
	pointsOfFunctionPlot[1_359].Y = 142_440.491

	pointsOfFunctionPlot[1_360].X = 8.57
	pointsOfFunctionPlot[1_360].Y = 144_428.907

	pointsOfFunctionPlot[1_361].X = 8.58
	pointsOfFunctionPlot[1_361].Y = 146_445.059

	pointsOfFunctionPlot[1_362].X = 8.59
	pointsOfFunctionPlot[1_362].Y = 148_489.356

	pointsOfFunctionPlot[1_363].X = 8.6
	pointsOfFunctionPlot[1_363].Y = 150_562.19

	pointsOfFunctionPlot[1_364].X = 8.61
	pointsOfFunctionPlot[1_364].Y = 152_663.96

	pointsOfFunctionPlot[1_365].X = 8.62
	pointsOfFunctionPlot[1_365].Y = 154_795.07

	pointsOfFunctionPlot[1_366].X = 8.63
	pointsOfFunctionPlot[1_366].Y = 156_955.928

	pointsOfFunctionPlot[1_367].X = 8.64
	pointsOfFunctionPlot[1_367].Y = 159_146.951

	pointsOfFunctionPlot[1_368].X = 8.65
	pointsOfFunctionPlot[1_368].Y = 161_368.56

	pointsOfFunctionPlot[1_369].X = 8.66
	pointsOfFunctionPlot[1_369].Y = 163_621.181

	pointsOfFunctionPlot[1_370].X = 8.67
	pointsOfFunctionPlot[1_370].Y = 165_905.248

	pointsOfFunctionPlot[1_371].X = 8.68
	pointsOfFunctionPlot[1_371].Y = 168_221.199

	pointsOfFunctionPlot[1_372].X = 8.69
	pointsOfFunctionPlot[1_372].Y = 170_569.479

	pointsOfFunctionPlot[1_373].X = 8.7
	pointsOfFunctionPlot[1_373].Y = 172_950.54

	pointsOfFunctionPlot[1_374].X = 8.71
	pointsOfFunctionPlot[1_374].Y = 175_364.84

	pointsOfFunctionPlot[1_375].X = 8.72
	pointsOfFunctionPlot[1_375].Y = 177_812.842

	pointsOfFunctionPlot[1_376].X = 8.73
	pointsOfFunctionPlot[1_376].Y = 180_295.017

	pointsOfFunctionPlot[1_377].X = 8.74
	pointsOfFunctionPlot[1_377].Y = 182_811.841

	pointsOfFunctionPlot[1_378].X = 8.75
	pointsOfFunctionPlot[1_378].Y = 185_363.8

	pointsOfFunctionPlot[1_379].X = 8.76
	pointsOfFunctionPlot[1_379].Y = 187_951.382

	pointsOfFunctionPlot[1_380].X = 8.77
	pointsOfFunctionPlot[1_380].Y = 190_575.085

	pointsOfFunctionPlot[1_381].X = 8.78
	pointsOfFunctionPlot[1_381].Y = 193_235.414

	pointsOfFunctionPlot[1_382].X = 8.79
	pointsOfFunctionPlot[1_382].Y = 195_932.88

	pointsOfFunctionPlot[1_383].X = 8.8
	pointsOfFunctionPlot[1_383].Y = 198_668.001

	pointsOfFunctionPlot[1_384].X = 8.81
	pointsOfFunctionPlot[1_384].Y = 201_441.303

	pointsOfFunctionPlot[1_385].X = 8.82
	pointsOfFunctionPlot[1_385].Y = 204_253.319

	pointsOfFunctionPlot[1_386].X = 8.83
	pointsOfFunctionPlot[1_386].Y = 207_104.589

	pointsOfFunctionPlot[1_387].X = 8.84
	pointsOfFunctionPlot[1_387].Y = 209_995.662

	pointsOfFunctionPlot[1_388].X = 8.85
	pointsOfFunctionPlot[1_388].Y = 212_927.092

	pointsOfFunctionPlot[1_389].X = 8.86
	pointsOfFunctionPlot[1_389].Y = 215_899.443

	pointsOfFunctionPlot[1_390].X = 8.87
	pointsOfFunctionPlot[1_390].Y = 218_913.287

	pointsOfFunctionPlot[1_391].X = 8.88
	pointsOfFunctionPlot[1_391].Y = 221_969.203

	pointsOfFunctionPlot[1_392].X = 8.89
	pointsOfFunctionPlot[1_392].Y = 225_067.777

	pointsOfFunctionPlot[1_393].X = 8.9
	pointsOfFunctionPlot[1_393].Y = 228_209.606

	pointsOfFunctionPlot[1_394].X = 8.91
	pointsOfFunctionPlot[1_394].Y = 231_395.294

	pointsOfFunctionPlot[1_395].X = 8.92
	pointsOfFunctionPlot[1_395].Y = 234_625.452

	pointsOfFunctionPlot[1_396].X = 8.93
	pointsOfFunctionPlot[1_396].Y = 237_900.701

	pointsOfFunctionPlot[1_397].X = 8.94
	pointsOfFunctionPlot[1_397].Y = 241_221.671

	pointsOfFunctionPlot[1_398].X = 8.95
	pointsOfFunctionPlot[1_398].Y = 244_589.0

	pointsOfFunctionPlot[1_399].X = 8.96
	pointsOfFunctionPlot[1_399].Y = 248_003.335

	pointsOfFunctionPlot[1_400].X = 8.97
	pointsOfFunctionPlot[1_400].Y = 251_465.333

	pointsOfFunctionPlot[1_401].X = 8.98
	pointsOfFunctionPlot[1_401].Y = 254_975.658

	pointsOfFunctionPlot[1_402].X = 8.99
	pointsOfFunctionPlot[1_402].Y = 258_534.986

	pointsOfFunctionPlot[1_403].X = 9.0
	pointsOfFunctionPlot[1_403].Y = 262_144.0

	pointsOfFunctionPlot[1_404].X = 9.01
	pointsOfFunctionPlot[1_404].Y = 265_803.393

	pointsOfFunctionPlot[1_405].X = 9.02
	pointsOfFunctionPlot[1_405].Y = 269_513.871

	pointsOfFunctionPlot[1_406].X = 9.03
	pointsOfFunctionPlot[1_406].Y = 273_276.144

	pointsOfFunctionPlot[1_407].X = 9.04
	pointsOfFunctionPlot[1_407].Y = 277_090.937

	pointsOfFunctionPlot[1_408].X = 9.05
	pointsOfFunctionPlot[1_408].Y = 280_958.982

	pointsOfFunctionPlot[1_409].X = 9.06
	pointsOfFunctionPlot[1_409].Y = 284_881.023

	pointsOfFunctionPlot[1_410].X = 9.07
	pointsOfFunctionPlot[1_410].Y = 288_857.814

	pointsOfFunctionPlot[1_411].X = 9.08
	pointsOfFunctionPlot[1_411].Y = 292_890.119

	pointsOfFunctionPlot[1_412].X = 9.09
	pointsOfFunctionPlot[1_412].Y = 296_978.713

	pointsOfFunctionPlot[1_413].X = 9.1
	pointsOfFunctionPlot[1_413].Y = 301_124.381

	pointsOfFunctionPlot[1_414].X = 9.11
	pointsOfFunctionPlot[1_414].Y = 305_327.921

	pointsOfFunctionPlot[1_415].X = 9.12
	pointsOfFunctionPlot[1_415].Y = 309_590.14

	pointsOfFunctionPlot[1_416].X = 9.13
	pointsOfFunctionPlot[1_416].Y = 313_911.857

	pointsOfFunctionPlot[1_417].X = 9.14
	pointsOfFunctionPlot[1_417].Y = 318_293.903

	pointsOfFunctionPlot[1_418].X = 9.15
	pointsOfFunctionPlot[1_418].Y = 322_737.121

	pointsOfFunctionPlot[1_419].X = 9.16
	pointsOfFunctionPlot[1_419].Y = 327_342.363

	pointsOfFunctionPlot[1_420].X = 9.17
	pointsOfFunctionPlot[1_420].Y = 331_810.496

	pointsOfFunctionPlot[1_421].X = 9.18
	pointsOfFunctionPlot[1_421].Y = 336_442.398

	pointsOfFunctionPlot[1_422].X = 9.19
	pointsOfFunctionPlot[1_422].Y = 341_138.959

	pointsOfFunctionPlot[1_423].X = 9.2
	pointsOfFunctionPlot[1_423].Y = 345_901.081

	pointsOfFunctionPlot[1_424].X = 9.21
	pointsOfFunctionPlot[1_424].Y = 350_729.68

	pointsOfFunctionPlot[1_425].X = 9.22
	pointsOfFunctionPlot[1_425].Y = 355_625.684

	pointsOfFunctionPlot[1_426].X = 9.23
	pointsOfFunctionPlot[1_426].Y = 360_590.034

	pointsOfFunctionPlot[1_427].X = 9.24
	pointsOfFunctionPlot[1_427].Y = 365_623.683

	pointsOfFunctionPlot[1_428].X = 9.25
	pointsOfFunctionPlot[1_428].Y = 370_727.6

	pointsOfFunctionPlot[1_429].X = 9.26
	pointsOfFunctionPlot[1_429].Y = 375_902.764

	pointsOfFunctionPlot[1_430].X = 9.27
	pointsOfFunctionPlot[1_430].Y = 381_150.171

	pointsOfFunctionPlot[1_431].X = 9.28
	pointsOfFunctionPlot[1_431].Y = 386_480.829

	pointsOfFunctionPlot[1_432].X = 9.29
	pointsOfFunctionPlot[1_432].Y = 391_865.761

	pointsOfFunctionPlot[1_433].X = 9.3
	pointsOfFunctionPlot[1_433].Y = 397_336.003

	pointsOfFunctionPlot[1_434].X = 9.31
	pointsOfFunctionPlot[1_434].Y = 402_882.607

	pointsOfFunctionPlot[1_435].X = 9.32
	pointsOfFunctionPlot[1_435].Y = 408_506.639

	pointsOfFunctionPlot[1_436].X = 9.33
	pointsOfFunctionPlot[1_436].Y = 414_209.179

	pointsOfFunctionPlot[1_437].X = 9.34
	pointsOfFunctionPlot[1_437].Y = 419_991.324

	pointsOfFunctionPlot[1_438].X = 9.35
	pointsOfFunctionPlot[1_438].Y = 425_854.184

	pointsOfFunctionPlot[1_439].X = 9.36
	pointsOfFunctionPlot[1_439].Y = 431_798.887

	pointsOfFunctionPlot[1_440].X = 9.37
	pointsOfFunctionPlot[1_440].Y = 437_826.575

	pointsOfFunctionPlot[1_441].X = 9.38
	pointsOfFunctionPlot[1_441].Y = 443_938.406

	pointsOfFunctionPlot[1_442].X = 9.39
	pointsOfFunctionPlot[1_442].Y = 450_135.555

	pointsOfFunctionPlot[1_443].X = 9.4
	pointsOfFunctionPlot[1_443].Y = 456_419.213

	pointsOfFunctionPlot[1_444].X = 9.41
	pointsOfFunctionPlot[1_444].Y = 462_790.588

	pointsOfFunctionPlot[1_445].X = 9.42
	pointsOfFunctionPlot[1_445].Y = 469_250.904

	pointsOfFunctionPlot[1_446].X = 9.43
	pointsOfFunctionPlot[1_446].Y = 475_801.402

	pointsOfFunctionPlot[1_447].X = 9.44
	pointsOfFunctionPlot[1_447].Y = 482_443.343

	pointsOfFunctionPlot[1_448].X = 9.45
	pointsOfFunctionPlot[1_448].Y = 489_178.001

	pointsOfFunctionPlot[1_449].X = 9.46
	pointsOfFunctionPlot[1_449].Y = 496_006.671

	pointsOfFunctionPlot[1_450].X = 9.47
	pointsOfFunctionPlot[1_450].Y = 502_930.666

	pointsOfFunctionPlot[1_451].X = 9.48
	pointsOfFunctionPlot[1_451].Y = 509_951.317

	pointsOfFunctionPlot[1_452].X = 9.49
	pointsOfFunctionPlot[1_452].Y = 517_069.972

	pointsOfFunctionPlot[1_453].X = 9.5
	pointsOfFunctionPlot[1_453].Y = 524_288.0

	pointsOfFunctionPlot[1_454].X = 9.51
	pointsOfFunctionPlot[1_454].Y = 531_606.787

	pointsOfFunctionPlot[1_455].X = 9.52
	pointsOfFunctionPlot[1_455].Y = 539_027.741

	pointsOfFunctionPlot[1_456].X = 9.53
	pointsOfFunctionPlot[1_456].Y = 546_552.288

	pointsOfFunctionPlot[1_457].X = 9.54
	pointsOfFunctionPlot[1_457].Y = 554_181.874

	pointsOfFunctionPlot[1_458].X = 9.55
	pointsOfFunctionPlot[1_458].Y = 561_917.965

	pointsOfFunctionPlot[1_459].X = 9.56
	pointsOfFunctionPlot[1_459].Y = 569_762.047

	pointsOfFunctionPlot[1_460].X = 9.57
	pointsOfFunctionPlot[1_460].Y = 577_715.629

	pointsOfFunctionPlot[1_461].X = 9.58
	pointsOfFunctionPlot[1_461].Y = 585_780.239

	pointsOfFunctionPlot[1_462].X = 9.59
	pointsOfFunctionPlot[1_462].Y = 593_957.426

	pointsOfFunctionPlot[1_463].X = 9.6
	pointsOfFunctionPlot[1_463].Y = 602_248.763

	pointsOfFunctionPlot[1_464].X = 9.61
	pointsOfFunctionPlot[1_464].Y = 610_655.842

	pointsOfFunctionPlot[1_465].X = 9.62
	pointsOfFunctionPlot[1_465].Y = 619_180.28

	pointsOfFunctionPlot[1_466].X = 9.63
	pointsOfFunctionPlot[1_466].Y = 627_823.715

	pointsOfFunctionPlot[1_467].X = 9.64
	pointsOfFunctionPlot[1_467].Y = 636_587.807

	pointsOfFunctionPlot[1_468].X = 9.65
	pointsOfFunctionPlot[1_468].Y = 645_474.242

	pointsOfFunctionPlot[1_469].X = 9.66
	pointsOfFunctionPlot[1_469].Y = 654_484.726

	pointsOfFunctionPlot[1_470].X = 9.67
	pointsOfFunctionPlot[1_470].Y = 663_620.993

	pointsOfFunctionPlot[1_471].X = 9.68
	pointsOfFunctionPlot[1_471].Y = 672_884.797

	pointsOfFunctionPlot[1_472].X = 9.69
	pointsOfFunctionPlot[1_472].Y = 682_277.918

	pointsOfFunctionPlot[1_473].X = 9.7
	pointsOfFunctionPlot[1_473].Y = 691_802.163

	pointsOfFunctionPlot[1_474].X = 9.71
	pointsOfFunctionPlot[1_474].Y = 701_459.361

	pointsOfFunctionPlot[1_475].X = 9.72
	pointsOfFunctionPlot[1_475].Y = 711_251.369

	pointsOfFunctionPlot[1_476].X = 9.73
	pointsOfFunctionPlot[1_476].Y = 721_180.068

	pointsOfFunctionPlot[1_477].X = 9.74
	pointsOfFunctionPlot[1_477].Y = 731_247.367

	pointsOfFunctionPlot[1_478].X = 9.75
	pointsOfFunctionPlot[1_478].Y = 741_455.2

	pointsOfFunctionPlot[1_479].X = 9.76
	pointsOfFunctionPlot[1_479].Y = 751_805.529

	pointsOfFunctionPlot[1_480].X = 9.77
	pointsOfFunctionPlot[1_480].Y = 762_300.343

	pointsOfFunctionPlot[1_481].X = 9.78
	pointsOfFunctionPlot[1_481].Y = 772_941.659

	pointsOfFunctionPlot[1_482].X = 9.79
	pointsOfFunctionPlot[1_482].Y = 783_731.522

	pointsOfFunctionPlot[1_483].X = 9.8
	pointsOfFunctionPlot[1_483].Y = 794_672.007

	pointsOfFunctionPlot[1_484].X = 9.81
	pointsOfFunctionPlot[1_484].Y = 805_765.215

	pointsOfFunctionPlot[1_485].X = 9.82
	pointsOfFunctionPlot[1_485].Y = 817_013.278

	pointsOfFunctionPlot[1_486].X = 9.83
	pointsOfFunctionPlot[1_486].Y = 828_418.358

	pointsOfFunctionPlot[1_487].X = 9.84
	pointsOfFunctionPlot[1_487].Y = 839_982.648

	pointsOfFunctionPlot[1_488].X = 9.85
	pointsOfFunctionPlot[1_488].Y = 851_708.368

	pointsOfFunctionPlot[1_489].X = 9.86
	pointsOfFunctionPlot[1_489].Y = 863_597.774

	pointsOfFunctionPlot[1_490].X = 9.87
	pointsOfFunctionPlot[1_490].Y = 875_653.15

	pointsOfFunctionPlot[1_491].X = 9.88
	pointsOfFunctionPlot[1_491].Y = 887_876.812

	pointsOfFunctionPlot[1_492].X = 9.89
	pointsOfFunctionPlot[1_492].Y = 900_271.111

	pointsOfFunctionPlot[1_493].X = 9.9
	pointsOfFunctionPlot[1_493].Y = 912_838.437

	pointsOfFunctionPlot[1_494].X = 9.91
	pointsOfFunctionPlot[1_494].Y = 925_581.177

	pointsOfFunctionPlot[1_495].X = 9.92
	pointsOfFunctionPlot[1_495].Y = 938_501.808

	pointsOfFunctionPlot[1_496].X = 9.93
	pointsOfFunctionPlot[1_496].Y = 951_602.805

	pointsOfFunctionPlot[1_497].X = 9.94
	pointsOfFunctionPlot[1_497].Y = 964_886.685

	pointsOfFunctionPlot[1_498].X = 9.95
	pointsOfFunctionPlot[1_498].Y = 978_356.002

	pointsOfFunctionPlot[1_499].X = 9.96
	pointsOfFunctionPlot[1_499].Y = 992_013.343

	pointsOfFunctionPlot[1_500].X = 9.97
	pointsOfFunctionPlot[1_500].Y = 1_005_861.333

	pointsOfFunctionPlot[1_501].X = 9.98
	pointsOfFunctionPlot[1_501].Y = 1_019_902.634

	pointsOfFunctionPlot[1_502].X = 9.99
	pointsOfFunctionPlot[1_502].Y = 1_034_139.944

	pointsOfFunctionPlot[1_503].X = 10.0
	pointsOfFunctionPlot[1_503].Y = 1_048_576.0










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function 4^x"

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
		"Power-of-4-plot-01.png"); err != nil {

		panic(err)
	}
}
