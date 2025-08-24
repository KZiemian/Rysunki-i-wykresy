package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function 3^x.

	pointsOfFunctionPlot := make(plotter.XYs, 1_625)

	pointsOfFunctionPlot[0].X = -10.0
	pointsOfFunctionPlot[0].Y = 0.0

	pointsOfFunctionPlot[1].X = -6.34
	pointsOfFunctionPlot[1].Y = 0.0

	pointsOfFunctionPlot[2].X = -6.33
	pointsOfFunctionPlot[2].Y = 0.001

	pointsOfFunctionPlot[3].X = -6.32
	pointsOfFunctionPlot[3].Y = 0.001

	pointsOfFunctionPlot[4].X = -6.31
	pointsOfFunctionPlot[4].Y = 0.001

	pointsOfFunctionPlot[5].X = -6.3
	pointsOfFunctionPlot[5].Y = 0.001

	pointsOfFunctionPlot[6].X = -6.29
	pointsOfFunctionPlot[6].Y = 0.001

	pointsOfFunctionPlot[7].X = -6.28
	pointsOfFunctionPlot[7].Y = 0.001

	pointsOfFunctionPlot[8].X = -6.27
	pointsOfFunctionPlot[8].Y = 0.001

	pointsOfFunctionPlot[9].X = -6.26
	pointsOfFunctionPlot[9].Y = 0.001

	pointsOfFunctionPlot[10].X = -6.25
	pointsOfFunctionPlot[10].Y = 0.001

	pointsOfFunctionPlot[11].X = -6.24
	pointsOfFunctionPlot[11].Y = 0.001

	pointsOfFunctionPlot[12].X = -6.23
	pointsOfFunctionPlot[12].Y = 0.001

	pointsOfFunctionPlot[13].X = -6.22
	pointsOfFunctionPlot[13].Y = 0.001

	pointsOfFunctionPlot[14].X = -6.21
	pointsOfFunctionPlot[14].Y = 0.001

	pointsOfFunctionPlot[15].X = -6.2
	pointsOfFunctionPlot[15].Y = 0.001

	pointsOfFunctionPlot[16].X = -6.19
	pointsOfFunctionPlot[16].Y = 0.001

	pointsOfFunctionPlot[17].X = -6.18
	pointsOfFunctionPlot[17].Y = 0.001

	pointsOfFunctionPlot[18].X = -6.17
	pointsOfFunctionPlot[18].Y = 0.001

	pointsOfFunctionPlot[19].X = -6.16
	pointsOfFunctionPlot[19].Y = 0.001

	pointsOfFunctionPlot[20].X = -6.15
	pointsOfFunctionPlot[20].Y = 0.001

	pointsOfFunctionPlot[21].X = -6.14
	pointsOfFunctionPlot[21].Y = 0.001

	pointsOfFunctionPlot[22].X = -6.13
	pointsOfFunctionPlot[22].Y = 0.001

	pointsOfFunctionPlot[23].X = -6.12
	pointsOfFunctionPlot[23].Y = 0.001

	pointsOfFunctionPlot[24].X = -6.11
	pointsOfFunctionPlot[24].Y = 0.001

	pointsOfFunctionPlot[25].X = -6.1
	pointsOfFunctionPlot[25].Y = 0.001

	pointsOfFunctionPlot[26].X = -6.1
	pointsOfFunctionPlot[26].Y = 0.001

	pointsOfFunctionPlot[27].X = -6.09
	pointsOfFunctionPlot[27].Y = 0.001

	pointsOfFunctionPlot[28].X = -6.08
	pointsOfFunctionPlot[28].Y = 0.001

	pointsOfFunctionPlot[29].X = -6.07
	pointsOfFunctionPlot[29].Y = 0.001

	pointsOfFunctionPlot[30].X = -6.06
	pointsOfFunctionPlot[30].Y = 0.001

	pointsOfFunctionPlot[31].X = -6.05
	pointsOfFunctionPlot[31].Y = 0.001

	pointsOfFunctionPlot[32].X = -6.04
	pointsOfFunctionPlot[32].Y = 0.001

	pointsOfFunctionPlot[33].X = -6.03
	pointsOfFunctionPlot[33].Y = 0.001

	pointsOfFunctionPlot[34].X = -6.02
	pointsOfFunctionPlot[34].Y = 0.001

	pointsOfFunctionPlot[35].X = -6.01
	pointsOfFunctionPlot[35].Y = 0.001

	pointsOfFunctionPlot[36].X = -6.0
	pointsOfFunctionPlot[36].Y = 0.001

	pointsOfFunctionPlot[37].X = -5.99
	pointsOfFunctionPlot[37].Y = 0.001

	pointsOfFunctionPlot[37].X = -5.98
	pointsOfFunctionPlot[37].Y = 0.001

	pointsOfFunctionPlot[38].X = -5.97
	pointsOfFunctionPlot[38].Y = 0.001

	pointsOfFunctionPlot[39].X = -5.96
	pointsOfFunctionPlot[39].Y = 0.001

	pointsOfFunctionPlot[40].X = -5.95
	pointsOfFunctionPlot[40].Y = 0.001

	pointsOfFunctionPlot[41].X = -5.94
	pointsOfFunctionPlot[41].Y = 0.001

	pointsOfFunctionPlot[42].X = -5.93
	pointsOfFunctionPlot[42].Y = 0.001

	pointsOfFunctionPlot[43].X = -5.92
	pointsOfFunctionPlot[43].Y = 0.001

	pointsOfFunctionPlot[44].X = -5.91
	pointsOfFunctionPlot[44].Y = 0.001

	pointsOfFunctionPlot[45].X = -5.9
	pointsOfFunctionPlot[45].Y = 0.001

	pointsOfFunctionPlot[46].X = -5.89
	pointsOfFunctionPlot[46].Y = 0.001

	pointsOfFunctionPlot[47].X = -5.88
	pointsOfFunctionPlot[47].Y = 0.001

	pointsOfFunctionPlot[48].X = -5.87
	pointsOfFunctionPlot[48].Y = 0.001

	pointsOfFunctionPlot[49].X = -5.86
	pointsOfFunctionPlot[49].Y = 0.001

	pointsOfFunctionPlot[50].X = -5.85
	pointsOfFunctionPlot[50].Y = 0.001

	pointsOfFunctionPlot[51].X = -5.84
	pointsOfFunctionPlot[51].Y = 0.001

	pointsOfFunctionPlot[52].X = -5.82
	pointsOfFunctionPlot[52].Y = 0.001

	pointsOfFunctionPlot[53].X = -5.82
	pointsOfFunctionPlot[53].Y = 0.001

	pointsOfFunctionPlot[54].X = -5.81
	pointsOfFunctionPlot[54].Y = 0.001

	pointsOfFunctionPlot[55].X = -5.8
	pointsOfFunctionPlot[55].Y = 0.001

	pointsOfFunctionPlot[56].X = -5.79
	pointsOfFunctionPlot[56].Y = 0.001

	pointsOfFunctionPlot[57].X = -5.78
	pointsOfFunctionPlot[57].Y = 0.001

	pointsOfFunctionPlot[58].X = -5.77
	pointsOfFunctionPlot[58].Y = 0.001

	pointsOfFunctionPlot[59].X = -5.76
	pointsOfFunctionPlot[59].Y = 0.001

	pointsOfFunctionPlot[60].X = -5.75
	pointsOfFunctionPlot[60].Y = 0.001

	pointsOfFunctionPlot[61].X = -5.74
	pointsOfFunctionPlot[61].Y = 0.001

	pointsOfFunctionPlot[62].X = -5.73
	pointsOfFunctionPlot[62].Y = 0.001

	pointsOfFunctionPlot[63].X = -5.72
	pointsOfFunctionPlot[63].Y = 0.001

	pointsOfFunctionPlot[64].X = -5.71
	pointsOfFunctionPlot[64].Y = 0.001

	pointsOfFunctionPlot[65].X = -5.7
	pointsOfFunctionPlot[65].Y = 0.001

	pointsOfFunctionPlot[66].X = -5.69
	pointsOfFunctionPlot[66].Y = 0.001

	pointsOfFunctionPlot[67].X = -5.68
	pointsOfFunctionPlot[67].Y = 0.001

	pointsOfFunctionPlot[68].X = -5.67
	pointsOfFunctionPlot[68].Y = 0.002

	pointsOfFunctionPlot[69].X = -5.66
	pointsOfFunctionPlot[69].Y = 0.002

	pointsOfFunctionPlot[70].X = -5.65
	pointsOfFunctionPlot[70].Y = 0.002

	pointsOfFunctionPlot[71].X = -5.64
	pointsOfFunctionPlot[71].Y = 0.002

	pointsOfFunctionPlot[72].X = -5.63
	pointsOfFunctionPlot[72].Y = 0.002

	pointsOfFunctionPlot[73].X = -5.62
	pointsOfFunctionPlot[73].Y = 0.002

	pointsOfFunctionPlot[74].X = -5.61
	pointsOfFunctionPlot[74].Y = 0.002

	pointsOfFunctionPlot[75].X = -5.6
	pointsOfFunctionPlot[75].Y = 0.002

	pointsOfFunctionPlot[76].X = -5.59
	pointsOfFunctionPlot[76].Y = 0.002

	pointsOfFunctionPlot[77].X = -5.58
	pointsOfFunctionPlot[77].Y = 0.002

	pointsOfFunctionPlot[78].X = -5.57
	pointsOfFunctionPlot[78].Y = 0.002

	pointsOfFunctionPlot[79].X = -5.56
	pointsOfFunctionPlot[79].Y = 0.002

	pointsOfFunctionPlot[80].X = -5.55
	pointsOfFunctionPlot[80].Y = 0.002

	pointsOfFunctionPlot[81].X = -5.54
	pointsOfFunctionPlot[81].Y = 0.002

	pointsOfFunctionPlot[82].X = -5.53
	pointsOfFunctionPlot[82].Y = 0.002

	pointsOfFunctionPlot[83].X = -5.52
	pointsOfFunctionPlot[83].Y = 0.002

	pointsOfFunctionPlot[84].X = -5.51
	pointsOfFunctionPlot[84].Y = 0.002

	pointsOfFunctionPlot[85].X = -5.5
	pointsOfFunctionPlot[85].Y = 0.002

	pointsOfFunctionPlot[86].X = -5.49
	pointsOfFunctionPlot[86].Y = 0.002

	pointsOfFunctionPlot[87].X = -5.48
	pointsOfFunctionPlot[87].Y = 0.002

	pointsOfFunctionPlot[88].X = -5.47
	pointsOfFunctionPlot[88].Y = 0.002

	pointsOfFunctionPlot[89].X = -5.46
	pointsOfFunctionPlot[89].Y = 0.002

	pointsOfFunctionPlot[90].X = -5.45
	pointsOfFunctionPlot[90].Y = 0.002

	pointsOfFunctionPlot[91].X = -5.44
	pointsOfFunctionPlot[91].Y = 0.002

	pointsOfFunctionPlot[92].X = -5.43
	pointsOfFunctionPlot[92].Y = 0.002

	pointsOfFunctionPlot[93].X = -5.42
	pointsOfFunctionPlot[93].Y = 0.002

	pointsOfFunctionPlot[94].X = -5.41
	pointsOfFunctionPlot[94].Y = 0.002

	pointsOfFunctionPlot[95].X = -5.4
	pointsOfFunctionPlot[95].Y = 0.002

	pointsOfFunctionPlot[96].X = -5.39
	pointsOfFunctionPlot[96].Y = 0.002

	pointsOfFunctionPlot[97].X = -5.38
	pointsOfFunctionPlot[97].Y = 0.002

	pointsOfFunctionPlot[98].X = -5.37
	pointsOfFunctionPlot[98].Y = 0.002

	pointsOfFunctionPlot[99].X = -5.36
	pointsOfFunctionPlot[99].Y = 0.002

	pointsOfFunctionPlot[100].X = -5.35
	pointsOfFunctionPlot[100].Y = 0.002

	pointsOfFunctionPlot[101].X = -5.34
	pointsOfFunctionPlot[101].Y = 0.002

	pointsOfFunctionPlot[102].X = -5.33
	pointsOfFunctionPlot[102].Y = 0.002

	pointsOfFunctionPlot[103].X = -5.32
	pointsOfFunctionPlot[103].Y = 0.002

	pointsOfFunctionPlot[104].X = -5.31
	pointsOfFunctionPlot[104].Y = 0.002

	pointsOfFunctionPlot[105].X = -5.3
	pointsOfFunctionPlot[105].Y = 0.003

	pointsOfFunctionPlot[106].X = -5.29
	pointsOfFunctionPlot[106].Y = 0.003

	pointsOfFunctionPlot[107].X = -5.28
	pointsOfFunctionPlot[107].Y = 0.003

	pointsOfFunctionPlot[108].X = -5.27
	pointsOfFunctionPlot[108].Y = 0.003

	pointsOfFunctionPlot[109].X = -5.26
	pointsOfFunctionPlot[109].Y = 0.003

	pointsOfFunctionPlot[110].X = -5.25
	pointsOfFunctionPlot[110].Y = 0.003

	pointsOfFunctionPlot[111].X = -5.24
	pointsOfFunctionPlot[111].Y = 0.003

	pointsOfFunctionPlot[112].X = -5.23
	pointsOfFunctionPlot[112].Y = 0.003

	pointsOfFunctionPlot[113].X = -5.22
	pointsOfFunctionPlot[113].Y = 0.003

	pointsOfFunctionPlot[114].X = -5.21
	pointsOfFunctionPlot[114].Y = 0.003

	pointsOfFunctionPlot[115].X = -5.2
	pointsOfFunctionPlot[115].Y = 0.003

	pointsOfFunctionPlot[116].X = -5.19
	pointsOfFunctionPlot[116].Y = 0.003

	pointsOfFunctionPlot[117].X = -5.18
	pointsOfFunctionPlot[117].Y = 0.003

	pointsOfFunctionPlot[118].X = -5.17
	pointsOfFunctionPlot[118].Y = 0.003

	pointsOfFunctionPlot[119].X = -5.16
	pointsOfFunctionPlot[119].Y = 0.003

	pointsOfFunctionPlot[120].X = -5.15
	pointsOfFunctionPlot[120].Y = 0.003

	pointsOfFunctionPlot[121].X = -5.14
	pointsOfFunctionPlot[121].Y = 0.003

	pointsOfFunctionPlot[122].X = -5.13
	pointsOfFunctionPlot[122].Y = 0.003

	pointsOfFunctionPlot[123].X = -5.12
	pointsOfFunctionPlot[123].Y = 0.003

	pointsOfFunctionPlot[124].X = -5.11
	pointsOfFunctionPlot[124].Y = 0.003

	pointsOfFunctionPlot[125].X = -5.1
	pointsOfFunctionPlot[125].Y = 0.003

	pointsOfFunctionPlot[126].X = -5.09
	pointsOfFunctionPlot[126].Y = 0.003

	pointsOfFunctionPlot[127].X = -5.08
	pointsOfFunctionPlot[127].Y = 0.003

	pointsOfFunctionPlot[128].X = -5.07
	pointsOfFunctionPlot[128].Y = 0.003

	pointsOfFunctionPlot[129].X = -5.06
	pointsOfFunctionPlot[129].Y = 0.003

	pointsOfFunctionPlot[130].X = -5.05
	pointsOfFunctionPlot[130].Y = 0.003

	pointsOfFunctionPlot[131].X = -5.04
	pointsOfFunctionPlot[131].Y = 0.003

	pointsOfFunctionPlot[132].X = -5.03
	pointsOfFunctionPlot[132].Y = 0.004

	pointsOfFunctionPlot[133].X = -5.02
	pointsOfFunctionPlot[133].Y = 0.004

	pointsOfFunctionPlot[134].X = -5.01
	pointsOfFunctionPlot[134].Y = 0.004

	pointsOfFunctionPlot[135].X = -5.0
	pointsOfFunctionPlot[135].Y = 0.004

	pointsOfFunctionPlot[136].X = -4.99
	pointsOfFunctionPlot[136].Y = 0.004

	pointsOfFunctionPlot[137].X = -4.98
	pointsOfFunctionPlot[137].Y = 0.004

	pointsOfFunctionPlot[138].X = -4.97
	pointsOfFunctionPlot[138].Y = 0.004

	pointsOfFunctionPlot[139].X = -4.96
	pointsOfFunctionPlot[139].Y = 0.004

	pointsOfFunctionPlot[140].X = -4.95
	pointsOfFunctionPlot[140].Y = 0.004

	pointsOfFunctionPlot[141].X = -4.94
	pointsOfFunctionPlot[141].Y = 0.004

	pointsOfFunctionPlot[142].X = -4.93
	pointsOfFunctionPlot[142].Y = 0.004

	pointsOfFunctionPlot[143].X = -4.92
	pointsOfFunctionPlot[143].Y = 0.004

	pointsOfFunctionPlot[144].X = -4.91
	pointsOfFunctionPlot[144].Y = 0.004

	pointsOfFunctionPlot[145].X = -4.9
	pointsOfFunctionPlot[145].Y = 0.004

	pointsOfFunctionPlot[146].X = -4.89
	pointsOfFunctionPlot[146].Y = 0.004

	pointsOfFunctionPlot[147].X = -4.88
	pointsOfFunctionPlot[147].Y = 0.004

	pointsOfFunctionPlot[148].X = -4.87
	pointsOfFunctionPlot[148].Y = 0.004

	pointsOfFunctionPlot[149].X = -4.86
	pointsOfFunctionPlot[149].Y = 0.004

	pointsOfFunctionPlot[150].X = -4.85
	pointsOfFunctionPlot[150].Y = 0.004

	pointsOfFunctionPlot[151].X = -4.84
	pointsOfFunctionPlot[151].Y = 0.004

	pointsOfFunctionPlot[152].X = -4.83
	pointsOfFunctionPlot[152].Y = 0.005

	pointsOfFunctionPlot[153].X = -4.82
	pointsOfFunctionPlot[153].Y = 0.005

	pointsOfFunctionPlot[154].X = -4.81
	pointsOfFunctionPlot[154].Y = 0.005

	pointsOfFunctionPlot[155].X = -4.8
	pointsOfFunctionPlot[155].Y = 0.005

	pointsOfFunctionPlot[156].X = -4.79
	pointsOfFunctionPlot[156].Y = 0.005

	pointsOfFunctionPlot[157].X = -4.78
	pointsOfFunctionPlot[157].Y = 0.005

	pointsOfFunctionPlot[158].X = -4.77
	pointsOfFunctionPlot[158].Y = 0.005

	pointsOfFunctionPlot[159].X = -4.76
	pointsOfFunctionPlot[159].Y = 0.005

	pointsOfFunctionPlot[160].X = -4.75
	pointsOfFunctionPlot[160].Y = 0.005

	pointsOfFunctionPlot[161].X = -4.74
	pointsOfFunctionPlot[161].Y = 0.005

	pointsOfFunctionPlot[162].X = -4.73
	pointsOfFunctionPlot[162].Y = 0.005

	pointsOfFunctionPlot[163].X = -4.72
	pointsOfFunctionPlot[163].Y = 0.005

	pointsOfFunctionPlot[164].X = -4.71
	pointsOfFunctionPlot[164].Y = 0.005

	pointsOfFunctionPlot[165].X = -4.7
	pointsOfFunctionPlot[165].Y = 0.005

	pointsOfFunctionPlot[166].X = -4.69
	pointsOfFunctionPlot[166].Y = 0.005

	pointsOfFunctionPlot[167].X = -4.68
	pointsOfFunctionPlot[167].Y = 0.005

	pointsOfFunctionPlot[168].X = -4.67
	pointsOfFunctionPlot[168].Y = 0.005

	pointsOfFunctionPlot[169].X = -4.66
	pointsOfFunctionPlot[169].Y = 0.006

	pointsOfFunctionPlot[170].X = -4.65
	pointsOfFunctionPlot[170].Y = 0.006

	pointsOfFunctionPlot[171].X = -4.64
	pointsOfFunctionPlot[171].Y = 0.006

	pointsOfFunctionPlot[172].X = -4.63
	pointsOfFunctionPlot[172].Y = 0.006

	pointsOfFunctionPlot[173].X = -4.62
	pointsOfFunctionPlot[173].Y = 0.006

	pointsOfFunctionPlot[174].X = -4.61
	pointsOfFunctionPlot[174].Y = 0.006

	pointsOfFunctionPlot[175].X = -4.6
	pointsOfFunctionPlot[175].Y = 0.006

	pointsOfFunctionPlot[176].X = -4.59
	pointsOfFunctionPlot[176].Y = 0.006

	pointsOfFunctionPlot[177].X = -4.58
	pointsOfFunctionPlot[177].Y = 0.006

	pointsOfFunctionPlot[178].X = -4.57
	pointsOfFunctionPlot[178].Y = 0.006

	pointsOfFunctionPlot[179].X = -4.56
	pointsOfFunctionPlot[179].Y = 0.006

	pointsOfFunctionPlot[180].X = -4.55
	pointsOfFunctionPlot[180].Y = 0.006

	pointsOfFunctionPlot[181].X = -4.54
	pointsOfFunctionPlot[181].Y = 0.006

	pointsOfFunctionPlot[182].X = -4.53
	pointsOfFunctionPlot[182].Y = 0.006

	pointsOfFunctionPlot[183].X = -4.52
	pointsOfFunctionPlot[183].Y = 0.006

	pointsOfFunctionPlot[184].X = -4.51
	pointsOfFunctionPlot[184].Y = 0.006

	pointsOfFunctionPlot[185].X = -4.5
	pointsOfFunctionPlot[185].Y = 0.007

	pointsOfFunctionPlot[186].X = -4.49
	pointsOfFunctionPlot[186].Y = 0.007

	pointsOfFunctionPlot[187].X = -4.48
	pointsOfFunctionPlot[187].Y = 0.007

	pointsOfFunctionPlot[188].X = -4.47
	pointsOfFunctionPlot[188].Y = 0.007

	pointsOfFunctionPlot[189].X = -4.46
	pointsOfFunctionPlot[189].Y = 0.007

	pointsOfFunctionPlot[190].X = -4.45
	pointsOfFunctionPlot[190].Y = 0.007

	pointsOfFunctionPlot[191].X = -4.44
	pointsOfFunctionPlot[191].Y = 0.007

	pointsOfFunctionPlot[192].X = -4.43
	pointsOfFunctionPlot[192].Y = 0.007

	pointsOfFunctionPlot[193].X = -4.42
	pointsOfFunctionPlot[193].Y = 0.007

	pointsOfFunctionPlot[194].X = -4.41
	pointsOfFunctionPlot[194].Y = 0.007

	pointsOfFunctionPlot[195].X = -4.4
	pointsOfFunctionPlot[195].Y = 0.008

	pointsOfFunctionPlot[196].X = -4.39
	pointsOfFunctionPlot[196].Y = 0.008

	pointsOfFunctionPlot[197].X = -4.38
	pointsOfFunctionPlot[197].Y = 0.008

	pointsOfFunctionPlot[198].X = -4.37
	pointsOfFunctionPlot[198].Y = 0.008

	pointsOfFunctionPlot[199].X = -4.36
	pointsOfFunctionPlot[199].Y = 0.008

	pointsOfFunctionPlot[200].X = -4.35
	pointsOfFunctionPlot[200].Y = 0.008

	pointsOfFunctionPlot[201].X = -4.34
	pointsOfFunctionPlot[201].Y = 0.008

	pointsOfFunctionPlot[202].X = -4.33
	pointsOfFunctionPlot[202].Y = 0.008

	pointsOfFunctionPlot[203].X = -4.32
	pointsOfFunctionPlot[203].Y = 0.008

	pointsOfFunctionPlot[204].X = -4.31
	pointsOfFunctionPlot[204].Y = 0.008

	pointsOfFunctionPlot[205].X = -4.3
	pointsOfFunctionPlot[205].Y = 0.008

	pointsOfFunctionPlot[206].X = -4.29
	pointsOfFunctionPlot[206].Y = 0.009

	pointsOfFunctionPlot[207].X = -4.28
	pointsOfFunctionPlot[207].Y = 0.009

	pointsOfFunctionPlot[208].X = -4.27
	pointsOfFunctionPlot[208].Y = 0.009

	pointsOfFunctionPlot[209].X = -4.26
	pointsOfFunctionPlot[209].Y = 0.009

	pointsOfFunctionPlot[210].X = -4.25
	pointsOfFunctionPlot[210].Y = 0.009

	pointsOfFunctionPlot[211].X = -4.24
	pointsOfFunctionPlot[211].Y = 0.009

	pointsOfFunctionPlot[212].X = -4.23
	pointsOfFunctionPlot[212].Y = 0.009

	pointsOfFunctionPlot[213].X = -4.22
	pointsOfFunctionPlot[213].Y = 0.009

	pointsOfFunctionPlot[214].X = -4.21
	pointsOfFunctionPlot[214].Y = 0.009

	pointsOfFunctionPlot[215].X = -4.2
	pointsOfFunctionPlot[215].Y = 0.009

	pointsOfFunctionPlot[216].X = -4.19
	pointsOfFunctionPlot[216].Y = 0.01

	pointsOfFunctionPlot[217].X = -4.18
	pointsOfFunctionPlot[217].Y = 0.01

	pointsOfFunctionPlot[218].X = -4.17
	pointsOfFunctionPlot[218].Y = 0.01

	pointsOfFunctionPlot[219].X = -4.16
	pointsOfFunctionPlot[219].Y = 0.01

	pointsOfFunctionPlot[220].X = -4.15
	pointsOfFunctionPlot[220].Y = 0.01

	pointsOfFunctionPlot[221].X = -4.14
	pointsOfFunctionPlot[221].Y = 0.01

	pointsOfFunctionPlot[222].X = -4.13
	pointsOfFunctionPlot[222].Y = 0.01

	pointsOfFunctionPlot[223].X = -4.12
	pointsOfFunctionPlot[223].Y = 0.01

	pointsOfFunctionPlot[224].X = -4.11
	pointsOfFunctionPlot[224].Y = 0.01

	pointsOfFunctionPlot[225].X = -4.1
	pointsOfFunctionPlot[225].Y = 0.011

	pointsOfFunctionPlot[226].X = -4.09
	pointsOfFunctionPlot[226].Y = 0.011

	pointsOfFunctionPlot[227].X = -4.08
	pointsOfFunctionPlot[227].Y = 0.011

	pointsOfFunctionPlot[228].X = -4.07
	pointsOfFunctionPlot[228].Y = 0.011

	pointsOfFunctionPlot[229].X = -4.06
	pointsOfFunctionPlot[229].Y = 0.011

	pointsOfFunctionPlot[230].X = -4.05
	pointsOfFunctionPlot[230].Y = 0.011

	pointsOfFunctionPlot[231].X = -4.04
	pointsOfFunctionPlot[231].Y = 0.011

	pointsOfFunctionPlot[232].X = -4.03
	pointsOfFunctionPlot[232].Y = 0.011

	pointsOfFunctionPlot[233].X = -4.02
	pointsOfFunctionPlot[233].Y = 0.012

	pointsOfFunctionPlot[234].X = -4.01
	pointsOfFunctionPlot[234].Y = 0.012

	pointsOfFunctionPlot[235].X = -4.0
	pointsOfFunctionPlot[235].Y = 0.012

	pointsOfFunctionPlot[236].X = -3.99
	pointsOfFunctionPlot[236].Y = 0.012

	pointsOfFunctionPlot[237].X = -3.98
	pointsOfFunctionPlot[237].Y = 0.012

	pointsOfFunctionPlot[238].X = -3.97
	pointsOfFunctionPlot[238].Y = 0.012

	pointsOfFunctionPlot[239].X = -3.96
	pointsOfFunctionPlot[239].Y = 0.012

	pointsOfFunctionPlot[240].X = -3.95
	pointsOfFunctionPlot[240].Y = 0.013

	pointsOfFunctionPlot[241].X = -3.94
	pointsOfFunctionPlot[241].Y = 0.013

	pointsOfFunctionPlot[242].X = -3.93
	pointsOfFunctionPlot[242].Y = 0.013

	pointsOfFunctionPlot[243].X = -3.92
	pointsOfFunctionPlot[243].Y = 0.013

	pointsOfFunctionPlot[244].X = -3.91
	pointsOfFunctionPlot[244].Y = 0.013

	pointsOfFunctionPlot[245].X = -3.9
	pointsOfFunctionPlot[245].Y = 0.013

	pointsOfFunctionPlot[246].X = -3.89
	pointsOfFunctionPlot[246].Y = 0.013

	pointsOfFunctionPlot[247].X = -3.88
	pointsOfFunctionPlot[247].Y = 0.014

	pointsOfFunctionPlot[248].X = -3.87
	pointsOfFunctionPlot[248].Y = 0.014

	pointsOfFunctionPlot[249].X = -3.86
	pointsOfFunctionPlot[249].Y = 0.014

	pointsOfFunctionPlot[250].X = -3.85
	pointsOfFunctionPlot[250].Y = 0.014

	pointsOfFunctionPlot[251].X = -3.84
	pointsOfFunctionPlot[251].Y = 0.014

	pointsOfFunctionPlot[252].X = -3.83
	pointsOfFunctionPlot[252].Y = 0.014

	pointsOfFunctionPlot[253].X = -3.82
	pointsOfFunctionPlot[253].Y = 0.015

	pointsOfFunctionPlot[254].X = -3.81
	pointsOfFunctionPlot[254].Y = 0.015

	pointsOfFunctionPlot[255].X = -3.8
	pointsOfFunctionPlot[255].Y = 0.015

	pointsOfFunctionPlot[256].X = -3.79
	pointsOfFunctionPlot[256].Y = 0.015

	pointsOfFunctionPlot[257].X = -3.78
	pointsOfFunctionPlot[257].Y = 0.015

	pointsOfFunctionPlot[258].X = -3.77
	pointsOfFunctionPlot[258].Y = 0.015

	pointsOfFunctionPlot[259].X = -3.76
	pointsOfFunctionPlot[259].Y = 0.016

	pointsOfFunctionPlot[260].X = -3.75
	pointsOfFunctionPlot[260].Y = 0.016

	pointsOfFunctionPlot[261].X = -3.74
	pointsOfFunctionPlot[261].Y = 0.016

	pointsOfFunctionPlot[262].X = -3.73
	pointsOfFunctionPlot[262].Y = 0.016

	pointsOfFunctionPlot[263].X = -3.72
	pointsOfFunctionPlot[263].Y = 0.016

	pointsOfFunctionPlot[264].X = -3.71
	pointsOfFunctionPlot[264].Y = 0.017

	pointsOfFunctionPlot[265].X = -3.7
	pointsOfFunctionPlot[265].Y = 0.017

	pointsOfFunctionPlot[266].X = -3.69
	pointsOfFunctionPlot[266].Y = 0.017

	pointsOfFunctionPlot[267].X = -3.68
	pointsOfFunctionPlot[267].Y = 0.017

	pointsOfFunctionPlot[268].X = -3.67
	pointsOfFunctionPlot[268].Y = 0.017

	pointsOfFunctionPlot[269].X = -3.66
	pointsOfFunctionPlot[269].Y = 0.017

	pointsOfFunctionPlot[270].X = -3.65
	pointsOfFunctionPlot[270].Y = 0.018

	pointsOfFunctionPlot[271].X = -3.64
	pointsOfFunctionPlot[271].Y = 0.018

	pointsOfFunctionPlot[272].X = -3.63
	pointsOfFunctionPlot[272].Y = 0.018

	pointsOfFunctionPlot[273].X = -3.62
	pointsOfFunctionPlot[273].Y = 0.018

	pointsOfFunctionPlot[274].X = -3.61
	pointsOfFunctionPlot[274].Y = 0.018

	pointsOfFunctionPlot[275].X = -3.6
	pointsOfFunctionPlot[275].Y = 0.019

	pointsOfFunctionPlot[276].X = -3.59
	pointsOfFunctionPlot[276].Y = 0.019

	pointsOfFunctionPlot[277].X = -3.58
	pointsOfFunctionPlot[277].Y = 0.019

	pointsOfFunctionPlot[278].X = -3.57
	pointsOfFunctionPlot[278].Y = 0.019

	pointsOfFunctionPlot[279].X = -3.56
	pointsOfFunctionPlot[279].Y = 0.02

	pointsOfFunctionPlot[280].X = -3.55
	pointsOfFunctionPlot[280].Y = 0.02

	pointsOfFunctionPlot[281].X = -3.54
	pointsOfFunctionPlot[281].Y = 0.02

	pointsOfFunctionPlot[282].X = -3.53
	pointsOfFunctionPlot[282].Y = 0.02

	pointsOfFunctionPlot[283].X = -3.52
	pointsOfFunctionPlot[283].Y = 0.02

	pointsOfFunctionPlot[284].X = -3.51
	pointsOfFunctionPlot[284].Y = 0.021

	pointsOfFunctionPlot[285].X = -3.5
	pointsOfFunctionPlot[285].Y = 0.021

	pointsOfFunctionPlot[286].X = -3.49
	pointsOfFunctionPlot[286].Y = 0.021

	pointsOfFunctionPlot[287].X = -3.48
	pointsOfFunctionPlot[287].Y = 0.021

	pointsOfFunctionPlot[288].X = -3.47
	pointsOfFunctionPlot[288].Y = 0.022

	pointsOfFunctionPlot[289].X = -3.46
	pointsOfFunctionPlot[289].Y = 0.022

	pointsOfFunctionPlot[290].X = -3.45
	pointsOfFunctionPlot[290].Y = 0.022

	pointsOfFunctionPlot[291].X = -3.44
	pointsOfFunctionPlot[291].Y = 0.022

	pointsOfFunctionPlot[292].X = -3.43
	pointsOfFunctionPlot[292].Y = 0.023

	pointsOfFunctionPlot[293].X = -3.42
	pointsOfFunctionPlot[293].Y = 0.023

	pointsOfFunctionPlot[294].X = -3.41
	pointsOfFunctionPlot[294].Y = 0.023

	pointsOfFunctionPlot[295].X = -3.4
	pointsOfFunctionPlot[295].Y = 0.023

	pointsOfFunctionPlot[296].X = -3.39
	pointsOfFunctionPlot[296].Y = 0.024

	pointsOfFunctionPlot[297].X = -3.38
	pointsOfFunctionPlot[297].Y = 0.024

	pointsOfFunctionPlot[298].X = -3.37
	pointsOfFunctionPlot[298].Y = 0.024

	pointsOfFunctionPlot[299].X = -3.36
	pointsOfFunctionPlot[299].Y = 0.024

	pointsOfFunctionPlot[300].X = -3.35
	pointsOfFunctionPlot[300].Y = 0.025

	pointsOfFunctionPlot[301].X = -3.34
	pointsOfFunctionPlot[301].Y = 0.025

	pointsOfFunctionPlot[302].X = -3.33
	pointsOfFunctionPlot[302].Y = 0.025

	pointsOfFunctionPlot[303].X = -3.32
	pointsOfFunctionPlot[303].Y = 0.026

	pointsOfFunctionPlot[304].X = -3.31
	pointsOfFunctionPlot[304].Y = 0.026

	pointsOfFunctionPlot[305].X = -3.3
	pointsOfFunctionPlot[305].Y = 0.026

	pointsOfFunctionPlot[306].X = -3.29
	pointsOfFunctionPlot[306].Y = 0.026

	pointsOfFunctionPlot[307].X = -3.28
	pointsOfFunctionPlot[307].Y = 0.027

	pointsOfFunctionPlot[308].X = -3.27
	pointsOfFunctionPlot[308].Y = 0.027

	pointsOfFunctionPlot[309].X = -3.26
	pointsOfFunctionPlot[309].Y = 0.027

	pointsOfFunctionPlot[310].X = -3.25
	pointsOfFunctionPlot[310].Y = 0.028

	pointsOfFunctionPlot[311].X = -3.24
	pointsOfFunctionPlot[311].Y = 0.028

	pointsOfFunctionPlot[312].X = -3.23
	pointsOfFunctionPlot[312].Y = 0.028

	pointsOfFunctionPlot[313].X = -3.22
	pointsOfFunctionPlot[313].Y = 0.029

	pointsOfFunctionPlot[314].X = -3.21
	pointsOfFunctionPlot[314].Y = 0.029

	pointsOfFunctionPlot[315].X = -3.2
	pointsOfFunctionPlot[315].Y = 0.029

	pointsOfFunctionPlot[316].X = -3.19
	pointsOfFunctionPlot[316].Y = 0.03

	pointsOfFunctionPlot[317].X = -3.18
	pointsOfFunctionPlot[317].Y = 0.03

	pointsOfFunctionPlot[318].X = -3.17
	pointsOfFunctionPlot[318].Y = 0.03

	pointsOfFunctionPlot[319].X = -3.16
	pointsOfFunctionPlot[319].Y = 0.031

	pointsOfFunctionPlot[320].X = -3.15
	pointsOfFunctionPlot[320].Y = 0.031

	pointsOfFunctionPlot[321].X = -3.14
	pointsOfFunctionPlot[321].Y = 0.031

	pointsOfFunctionPlot[322].X = -3.13
	pointsOfFunctionPlot[322].Y = 0.032

	pointsOfFunctionPlot[323].X = -3.12
	pointsOfFunctionPlot[323].Y = 0.032

	pointsOfFunctionPlot[324].X = -3.11
	pointsOfFunctionPlot[324].Y = 0.032

	pointsOfFunctionPlot[325].X = -3.1
	pointsOfFunctionPlot[325].Y = 0.033

	pointsOfFunctionPlot[326].X = -3.09
	pointsOfFunctionPlot[326].Y = 0.033

	pointsOfFunctionPlot[327].X = -3.08
	pointsOfFunctionPlot[327].Y = 0.033

	pointsOfFunctionPlot[328].X = -3.07
	pointsOfFunctionPlot[328].Y = 0.034

	pointsOfFunctionPlot[329].X = -3.06
	pointsOfFunctionPlot[329].Y = 0.034

	pointsOfFunctionPlot[330].X = -3.05
	pointsOfFunctionPlot[330].Y = 0.035

	pointsOfFunctionPlot[331].X = -3.04
	pointsOfFunctionPlot[331].Y = 0.035

	pointsOfFunctionPlot[332].X = -3.03
	pointsOfFunctionPlot[332].Y = 0.035

	pointsOfFunctionPlot[333].X = -3.02
	pointsOfFunctionPlot[333].Y = 0.036

	pointsOfFunctionPlot[334].X = -3.01
	pointsOfFunctionPlot[334].Y = 0.036

	pointsOfFunctionPlot[335].X = -3.0
	pointsOfFunctionPlot[335].Y = 0.037

	pointsOfFunctionPlot[336].X = -2.99
	pointsOfFunctionPlot[336].Y = 0.037

	pointsOfFunctionPlot[337].X = -2.98
	pointsOfFunctionPlot[337].Y = 0.037

	pointsOfFunctionPlot[338].X = -2.97
	pointsOfFunctionPlot[338].Y = 0.038

	pointsOfFunctionPlot[339].X = -2.96
	pointsOfFunctionPlot[339].Y = 0.038

	pointsOfFunctionPlot[340].X = -2.95
	pointsOfFunctionPlot[340].Y = 0.039

	pointsOfFunctionPlot[341].X = -2.94
	pointsOfFunctionPlot[341].Y = 0.039

	pointsOfFunctionPlot[342].X = -2.93
	pointsOfFunctionPlot[342].Y = 0.04

	pointsOfFunctionPlot[343].X = -2.92
	pointsOfFunctionPlot[343].Y = 0.04

	pointsOfFunctionPlot[344].X = -2.91
	pointsOfFunctionPlot[344].Y = 0.04

	pointsOfFunctionPlot[345].X = -2.9
	pointsOfFunctionPlot[345].Y = 0.041

	pointsOfFunctionPlot[346].X = -2.89
	pointsOfFunctionPlot[346].Y = 0.041

	pointsOfFunctionPlot[347].X = -2.88
	pointsOfFunctionPlot[347].Y = 0.042

	pointsOfFunctionPlot[348].X = -2.87
	pointsOfFunctionPlot[348].Y = 0.042

	pointsOfFunctionPlot[349].X = -2.86
	pointsOfFunctionPlot[349].Y = 0.043

	pointsOfFunctionPlot[350].X = -2.85
	pointsOfFunctionPlot[350].Y = 0.043

	pointsOfFunctionPlot[351].X = -2.84
	pointsOfFunctionPlot[351].Y = 0.044

	pointsOfFunctionPlot[352].X = -2.83
	pointsOfFunctionPlot[352].Y = 0.044

	pointsOfFunctionPlot[353].X = -2.82
	pointsOfFunctionPlot[353].Y = 0.045

	pointsOfFunctionPlot[354].X = -2.81
	pointsOfFunctionPlot[354].Y = 0.045

	pointsOfFunctionPlot[355].X = -2.8
	pointsOfFunctionPlot[355].Y = 0.046

	pointsOfFunctionPlot[356].X = -2.79
	pointsOfFunctionPlot[356].Y = 0.046

	pointsOfFunctionPlot[357].X = -2.78
	pointsOfFunctionPlot[357].Y = 0.047

	pointsOfFunctionPlot[358].X = -2.77
	pointsOfFunctionPlot[358].Y = 0.047

	pointsOfFunctionPlot[359].X = -2.76
	pointsOfFunctionPlot[359].Y = 0.048

	pointsOfFunctionPlot[360].X = -2.75
	pointsOfFunctionPlot[360].Y = 0.048

	pointsOfFunctionPlot[361].X = -2.74
	pointsOfFunctionPlot[361].Y = 0.049

	pointsOfFunctionPlot[362].X = -2.73
	pointsOfFunctionPlot[362].Y = 0.049

	pointsOfFunctionPlot[363].X = -2.72
	pointsOfFunctionPlot[363].Y = 0.05

	pointsOfFunctionPlot[364].X = -2.71
	pointsOfFunctionPlot[364].Y = 0.05

	pointsOfFunctionPlot[365].X = -2.7
	pointsOfFunctionPlot[365].Y = 0.051

	pointsOfFunctionPlot[366].X = -2.69
	pointsOfFunctionPlot[366].Y = 0.052

	pointsOfFunctionPlot[367].X = -2.68
	pointsOfFunctionPlot[367].Y = 0.052

	pointsOfFunctionPlot[368].X = -2.67
	pointsOfFunctionPlot[368].Y = 0.053

	pointsOfFunctionPlot[369].X = -2.66
	pointsOfFunctionPlot[369].Y = 0.053

	pointsOfFunctionPlot[370].X = -2.65
	pointsOfFunctionPlot[370].Y = 0.054

	pointsOfFunctionPlot[371].X = -2.64
	pointsOfFunctionPlot[371].Y = 0.055

	pointsOfFunctionPlot[372].X = -2.63
	pointsOfFunctionPlot[372].Y = 0.055

	pointsOfFunctionPlot[373].X = -2.62
	pointsOfFunctionPlot[373].Y = 0.056

	pointsOfFunctionPlot[374].X = -2.61
	pointsOfFunctionPlot[374].Y = 0.056

	pointsOfFunctionPlot[375].X = -2.6
	pointsOfFunctionPlot[375].Y = 0.057

	pointsOfFunctionPlot[376].X = -2.59
	pointsOfFunctionPlot[376].Y = 0.058

	pointsOfFunctionPlot[377].X = -2.58
	pointsOfFunctionPlot[377].Y = 0.059

	pointsOfFunctionPlot[378].X = -2.57
	pointsOfFunctionPlot[378].Y = 0.059

	pointsOfFunctionPlot[379].X = -2.56
	pointsOfFunctionPlot[379].Y = 0.06

	pointsOfFunctionPlot[380].X = -2.55
	pointsOfFunctionPlot[380].Y = 0.06

	pointsOfFunctionPlot[381].X = -2.54
	pointsOfFunctionPlot[381].Y = 0.061

	pointsOfFunctionPlot[382].X = -2.53
	pointsOfFunctionPlot[382].Y = 0.062

	pointsOfFunctionPlot[383].X = -2.52
	pointsOfFunctionPlot[383].Y = 0.062

	pointsOfFunctionPlot[384].X = -2.51
	pointsOfFunctionPlot[384].Y = 0.063

	pointsOfFunctionPlot[385].X = -2.5
	pointsOfFunctionPlot[385].Y = 0.064

	pointsOfFunctionPlot[386].X = -2.49
	pointsOfFunctionPlot[386].Y = 0.064

	pointsOfFunctionPlot[387].X = -2.48
	pointsOfFunctionPlot[387].Y = 0.065

	pointsOfFunctionPlot[388].X = -2.47
	pointsOfFunctionPlot[388].Y = 0.066

	pointsOfFunctionPlot[389].X = -2.46
	pointsOfFunctionPlot[389].Y = 0.067

	pointsOfFunctionPlot[390].X = -2.45
	pointsOfFunctionPlot[390].Y = 0.067

	pointsOfFunctionPlot[391].X = -2.44
	pointsOfFunctionPlot[391].Y = 0.068

	pointsOfFunctionPlot[392].X = -2.43
	pointsOfFunctionPlot[392].Y = 0.069

	pointsOfFunctionPlot[393].X = -2.42
	pointsOfFunctionPlot[393].Y = 0.07

	pointsOfFunctionPlot[394].X = -2.41
	pointsOfFunctionPlot[394].Y = 0.07

	pointsOfFunctionPlot[395].X = -2.4
	pointsOfFunctionPlot[395].Y = 0.071

	pointsOfFunctionPlot[396].X = -2.39
	pointsOfFunctionPlot[396].Y = 0.072

	pointsOfFunctionPlot[397].X = -2.38
	pointsOfFunctionPlot[397].Y = 0.073

	pointsOfFunctionPlot[398].X = -2.37
	pointsOfFunctionPlot[398].Y = 0.074

	pointsOfFunctionPlot[399].X = -2.36
	pointsOfFunctionPlot[399].Y = 0.074

	pointsOfFunctionPlot[400].X = -2.35
	pointsOfFunctionPlot[400].Y = 0.075

	pointsOfFunctionPlot[401].X = -2.34
	pointsOfFunctionPlot[401].Y = 0.076

	pointsOfFunctionPlot[402].X = -2.33
	pointsOfFunctionPlot[402].Y = 0.077

	pointsOfFunctionPlot[403].X = -2.32
	pointsOfFunctionPlot[403].Y = 0.078

	pointsOfFunctionPlot[404].X = -2.31
	pointsOfFunctionPlot[404].Y = 0.079

	pointsOfFunctionPlot[405].X = -2.3
	pointsOfFunctionPlot[405].Y = 0.079

	pointsOfFunctionPlot[406].X = -2.29
	pointsOfFunctionPlot[406].Y = 0.08

	pointsOfFunctionPlot[407].X = -2.28
	pointsOfFunctionPlot[407].Y = 0.081

	pointsOfFunctionPlot[408].X = -2.27
	pointsOfFunctionPlot[408].Y = 0.082

	pointsOfFunctionPlot[409].X = -2.26
	pointsOfFunctionPlot[409].Y = 0.083

	pointsOfFunctionPlot[410].X = -2.25
	pointsOfFunctionPlot[410].Y = 0.084

	pointsOfFunctionPlot[411].X = -2.24
	pointsOfFunctionPlot[411].Y = 0.085

	pointsOfFunctionPlot[412].X = -2.23
	pointsOfFunctionPlot[412].Y = 0.086

	pointsOfFunctionPlot[413].X = -2.22
	pointsOfFunctionPlot[413].Y = 0.087

	pointsOfFunctionPlot[414].X = -2.21
	pointsOfFunctionPlot[414].Y = 0.088

	pointsOfFunctionPlot[415].X = -2.2
	pointsOfFunctionPlot[415].Y = 0.089

	pointsOfFunctionPlot[416].X = -2.19
	pointsOfFunctionPlot[416].Y = 0.09

	pointsOfFunctionPlot[417].X = -2.18
	pointsOfFunctionPlot[417].Y = 0.091

	pointsOfFunctionPlot[418].X = -2.17
	pointsOfFunctionPlot[418].Y = 0.092

	pointsOfFunctionPlot[419].X = -2.16
	pointsOfFunctionPlot[419].Y = 0.093

	pointsOfFunctionPlot[420].X = -2.15
	pointsOfFunctionPlot[420].Y = 0.094

	pointsOfFunctionPlot[421].X = -2.14
	pointsOfFunctionPlot[421].Y = 0.095

	pointsOfFunctionPlot[422].X = -2.13
	pointsOfFunctionPlot[422].Y = 0.096

	pointsOfFunctionPlot[423].X = -2.12
	pointsOfFunctionPlot[423].Y = 0.097

	pointsOfFunctionPlot[424].X = -2.11
	pointsOfFunctionPlot[424].Y = 0.098

	pointsOfFunctionPlot[425].X = -2.1
	pointsOfFunctionPlot[425].Y = 0.099

	pointsOfFunctionPlot[426].X = -2.09
	pointsOfFunctionPlot[426].Y = 0.1

	pointsOfFunctionPlot[427].X = -2.08
	pointsOfFunctionPlot[427].Y = 0.101

	pointsOfFunctionPlot[428].X = -2.07
	pointsOfFunctionPlot[428].Y = 0.102

	pointsOfFunctionPlot[429].X = -2.06
	pointsOfFunctionPlot[429].Y = 0.104

	pointsOfFunctionPlot[430].X = -2.05
	pointsOfFunctionPlot[430].Y = 0.105

	pointsOfFunctionPlot[431].X = -2.04
	pointsOfFunctionPlot[431].Y = 0.106

	pointsOfFunctionPlot[432].X = -2.03
	pointsOfFunctionPlot[432].Y = 0.107

	pointsOfFunctionPlot[433].X = -2.02
	pointsOfFunctionPlot[433].Y = 0.108

	pointsOfFunctionPlot[434].X = -2.01
	pointsOfFunctionPlot[434].Y = 0.109

	pointsOfFunctionPlot[435].X = -2.0
	pointsOfFunctionPlot[435].Y = 0.111

	pointsOfFunctionPlot[436].X = -1.99
	pointsOfFunctionPlot[436].Y = 0.112

	pointsOfFunctionPlot[437].X = -1.98
	pointsOfFunctionPlot[437].Y = 0.113

	pointsOfFunctionPlot[438].X = -1.97
	pointsOfFunctionPlot[438].Y = 0.114

	pointsOfFunctionPlot[439].X = -1.96
	pointsOfFunctionPlot[439].Y = 0.116

	pointsOfFunctionPlot[440].X = -1.95
	pointsOfFunctionPlot[440].Y = 0.117

	pointsOfFunctionPlot[441].X = -1.94
	pointsOfFunctionPlot[441].Y = 0.118

	pointsOfFunctionPlot[442].X = -1.93
	pointsOfFunctionPlot[442].Y = 0.12

	pointsOfFunctionPlot[443].X = -1.92
	pointsOfFunctionPlot[443].Y = 0.121

	pointsOfFunctionPlot[444].X = -1.91
	pointsOfFunctionPlot[444].Y = 0.122

	pointsOfFunctionPlot[445].X = -1.9
	pointsOfFunctionPlot[445].Y = 0.124

	pointsOfFunctionPlot[446].X = -1.89
	pointsOfFunctionPlot[446].Y = 0.125

	pointsOfFunctionPlot[447].X = -1.88
	pointsOfFunctionPlot[447].Y = 0.126

	pointsOfFunctionPlot[448].X = -1.87
	pointsOfFunctionPlot[448].Y = 0.128

	pointsOfFunctionPlot[449].X = -1.86
	pointsOfFunctionPlot[449].Y = 0.129

	pointsOfFunctionPlot[450].X = -1.85
	pointsOfFunctionPlot[450].Y = 0.131

	pointsOfFunctionPlot[451].X = -1.84
	pointsOfFunctionPlot[451].Y = 0.132

	pointsOfFunctionPlot[452].X = -1.83
	pointsOfFunctionPlot[452].Y = 0.133

	pointsOfFunctionPlot[453].X = -1.82
	pointsOfFunctionPlot[453].Y = 0.135

	pointsOfFunctionPlot[454].X = -1.81
	pointsOfFunctionPlot[454].Y = 0.136

	pointsOfFunctionPlot[455].X = -1.8
	pointsOfFunctionPlot[455].Y = 0.138

	pointsOfFunctionPlot[456].X = -1.79
	pointsOfFunctionPlot[456].Y = 0.139

	pointsOfFunctionPlot[457].X = -1.78
	pointsOfFunctionPlot[457].Y = 0.141

	pointsOfFunctionPlot[458].X = -1.77
	pointsOfFunctionPlot[458].Y = 0.143

	pointsOfFunctionPlot[459].X = -1.76
	pointsOfFunctionPlot[459].Y = 0.144

	pointsOfFunctionPlot[460].X = -1.75
	pointsOfFunctionPlot[460].Y = 0.146

	pointsOfFunctionPlot[461].X = -1.74
	pointsOfFunctionPlot[461].Y = 0.147

	pointsOfFunctionPlot[462].X = -1.73
	pointsOfFunctionPlot[462].Y = 0.149

	pointsOfFunctionPlot[463].X = -1.72
	pointsOfFunctionPlot[463].Y = 0.151

	pointsOfFunctionPlot[464].X = -1.71
	pointsOfFunctionPlot[464].Y = 0.152

	pointsOfFunctionPlot[465].X = -1.7
	pointsOfFunctionPlot[465].Y = 0.154

	pointsOfFunctionPlot[466].X = -1.69
	pointsOfFunctionPlot[466].Y = 0.156

	pointsOfFunctionPlot[467].X = -1.68
	pointsOfFunctionPlot[467].Y = 0.157

	pointsOfFunctionPlot[468].X = -1.67
	pointsOfFunctionPlot[468].Y = 0.159

	pointsOfFunctionPlot[469].X = -1.66
	pointsOfFunctionPlot[469].Y = 0.161

	pointsOfFunctionPlot[470].X = -1.65
	pointsOfFunctionPlot[470].Y = 0.163

	pointsOfFunctionPlot[471].X = -1.64
	pointsOfFunctionPlot[471].Y = 0.165

	pointsOfFunctionPlot[472].X = -1.63
	pointsOfFunctionPlot[472].Y = 0.166

	pointsOfFunctionPlot[473].X = -1.62
	pointsOfFunctionPlot[473].Y = 0.168

	pointsOfFunctionPlot[474].X = -1.61
	pointsOfFunctionPlot[474].Y = 0.17

	pointsOfFunctionPlot[475].X = -1.6
	pointsOfFunctionPlot[475].Y = 0.172

	pointsOfFunctionPlot[476].X = -1.59
	pointsOfFunctionPlot[476].Y = 0.174

	pointsOfFunctionPlot[477].X = -1.58
	pointsOfFunctionPlot[477].Y = 0.176

	pointsOfFunctionPlot[478].X = -1.57
	pointsOfFunctionPlot[478].Y = 0.178

	pointsOfFunctionPlot[479].X = -1.56
	pointsOfFunctionPlot[479].Y = 0.18

	pointsOfFunctionPlot[480].X = -1.55
	pointsOfFunctionPlot[480].Y = 0.182

	pointsOfFunctionPlot[481].X = -1.54
	pointsOfFunctionPlot[481].Y = 0.184

	pointsOfFunctionPlot[482].X = -1.53
	pointsOfFunctionPlot[482].Y = 0.186

	pointsOfFunctionPlot[483].X = -1.52
	pointsOfFunctionPlot[483].Y = 0.188

	pointsOfFunctionPlot[484].X = -1.51
	pointsOfFunctionPlot[484].Y = 0.19

	pointsOfFunctionPlot[485].X = -1.5
	pointsOfFunctionPlot[485].Y = 0.192

	pointsOfFunctionPlot[486].X = -1.49
	pointsOfFunctionPlot[486].Y = 0.194

	pointsOfFunctionPlot[487].X = -1.48
	pointsOfFunctionPlot[487].Y = 0.196

	pointsOfFunctionPlot[488].X = -1.47
	pointsOfFunctionPlot[488].Y = 0.198

	pointsOfFunctionPlot[489].X = -1.46
	pointsOfFunctionPlot[489].Y = 0.201

	pointsOfFunctionPlot[490].X = -1.45
	pointsOfFunctionPlot[490].Y = 0.203

	pointsOfFunctionPlot[491].X = -1.44
	pointsOfFunctionPlot[491].Y = 0.205

	pointsOfFunctionPlot[492].X = -1.43
	pointsOfFunctionPlot[492].Y = 0.207

	pointsOfFunctionPlot[493].X = -1.42
	pointsOfFunctionPlot[493].Y = 0.21

	pointsOfFunctionPlot[494].X = -1.41
	pointsOfFunctionPlot[494].Y = 0.212

	pointsOfFunctionPlot[495].X = -1.4
	pointsOfFunctionPlot[495].Y = 0.214

	pointsOfFunctionPlot[496].X = -1.39
	pointsOfFunctionPlot[496].Y = 0.217

	pointsOfFunctionPlot[497].X = -1.38
	pointsOfFunctionPlot[497].Y = 0.219

	pointsOfFunctionPlot[498].X = -1.37
	pointsOfFunctionPlot[498].Y = 0.222

	pointsOfFunctionPlot[499].X = -1.36
	pointsOfFunctionPlot[499].Y = 0.224

	pointsOfFunctionPlot[500].X = -1.35
	pointsOfFunctionPlot[500].Y = 0.226

	pointsOfFunctionPlot[501].X = -1.34
	pointsOfFunctionPlot[501].Y = 0.229

	pointsOfFunctionPlot[502].X = -1.33
	pointsOfFunctionPlot[502].Y = 0.232

	pointsOfFunctionPlot[503].X = -1.32
	pointsOfFunctionPlot[503].Y = 0.234

	pointsOfFunctionPlot[504].X = -1.31
	pointsOfFunctionPlot[504].Y = 0.237

	pointsOfFunctionPlot[505].X = -1.3
	pointsOfFunctionPlot[505].Y = 0.239

	pointsOfFunctionPlot[506].X = -1.29
	pointsOfFunctionPlot[506].Y = 0.242

	pointsOfFunctionPlot[507].X = -1.28
	pointsOfFunctionPlot[507].Y = 0.245

	pointsOfFunctionPlot[508].X = -1.27
	pointsOfFunctionPlot[508].Y = 0.247

	pointsOfFunctionPlot[509].X = -1.26
	pointsOfFunctionPlot[509].Y = 0.25

	pointsOfFunctionPlot[510].X = -1.25
	pointsOfFunctionPlot[510].Y = 0.253

	pointsOfFunctionPlot[511].X = -1.24
	pointsOfFunctionPlot[511].Y = 0.256

	pointsOfFunctionPlot[512].X = -1.23
	pointsOfFunctionPlot[512].Y = 0.258

	pointsOfFunctionPlot[513].X = -1.22
	pointsOfFunctionPlot[513].Y = 0.261

	pointsOfFunctionPlot[514].X = -1.21
	pointsOfFunctionPlot[514].Y = 0.264

	pointsOfFunctionPlot[515].X = -1.2
	pointsOfFunctionPlot[515].Y = 0.267

	pointsOfFunctionPlot[516].X = -1.19
	pointsOfFunctionPlot[516].Y = 0.27

	pointsOfFunctionPlot[517].X = -1.18
	pointsOfFunctionPlot[517].Y = 0.273

	pointsOfFunctionPlot[518].X = -1.17
	pointsOfFunctionPlot[518].Y = 0.276

	pointsOfFunctionPlot[519].X = -1.16
	pointsOfFunctionPlot[519].Y = 0.279

	pointsOfFunctionPlot[520].X = -1.15
	pointsOfFunctionPlot[520].Y = 0.282

	pointsOfFunctionPlot[521].X = -1.14
	pointsOfFunctionPlot[521].Y = 0.285

	pointsOfFunctionPlot[522].X = -1.13
	pointsOfFunctionPlot[522].Y = 0.289

	pointsOfFunctionPlot[523].X = -1.12
	pointsOfFunctionPlot[523].Y = 0.292

	pointsOfFunctionPlot[524].X = -1.11
	pointsOfFunctionPlot[524].Y = 0.295

	pointsOfFunctionPlot[525].X = -1.1
	pointsOfFunctionPlot[525].Y = 0.298

	pointsOfFunctionPlot[526].X = -1.09
	pointsOfFunctionPlot[526].Y = 0.302

	pointsOfFunctionPlot[527].X = -1.08
	pointsOfFunctionPlot[527].Y = 0.305

	pointsOfFunctionPlot[528].X = -1.07
	pointsOfFunctionPlot[528].Y = 0.308

	pointsOfFunctionPlot[529].X = -1.06
	pointsOfFunctionPlot[529].Y = 0.312

	pointsOfFunctionPlot[530].X = -1.05
	pointsOfFunctionPlot[530].Y = 0.315

	pointsOfFunctionPlot[531].X = -1.04
	pointsOfFunctionPlot[531].Y = 0.319

	pointsOfFunctionPlot[532].X = -1.03
	pointsOfFunctionPlot[532].Y = 0.322

	pointsOfFunctionPlot[533].X = -1.02
	pointsOfFunctionPlot[533].Y = 0.326

	pointsOfFunctionPlot[534].X = -1.01
	pointsOfFunctionPlot[534].Y = 0.329

	pointsOfFunctionPlot[535].X = -1.0
	pointsOfFunctionPlot[535].Y = 0.333

	pointsOfFunctionPlot[536].X = -0.99
	pointsOfFunctionPlot[536].Y = 0.337

	pointsOfFunctionPlot[537].X = -0.98
	pointsOfFunctionPlot[537].Y = 0.34

	pointsOfFunctionPlot[538].X = -0.97
	pointsOfFunctionPlot[538].Y = 0.344

	pointsOfFunctionPlot[539].X = -0.96
	pointsOfFunctionPlot[539].Y = 0.348

	pointsOfFunctionPlot[540].X = -0.95
	pointsOfFunctionPlot[540].Y = 0.352

	pointsOfFunctionPlot[541].X = -0.94
	pointsOfFunctionPlot[541].Y = 0.356

	pointsOfFunctionPlot[542].X = -0.93
	pointsOfFunctionPlot[542].Y = 0.36

	pointsOfFunctionPlot[543].X = -0.92
	pointsOfFunctionPlot[543].Y = 0.364

	pointsOfFunctionPlot[544].X = -0.91
	pointsOfFunctionPlot[544].Y = 0.368

	pointsOfFunctionPlot[545].X = -0.9
	pointsOfFunctionPlot[545].Y = 0.372

	pointsOfFunctionPlot[546].X = -0.89
	pointsOfFunctionPlot[546].Y = 0.376

	pointsOfFunctionPlot[547].X = -0.88
	pointsOfFunctionPlot[547].Y = 0.38

	pointsOfFunctionPlot[548].X = -0.87
	pointsOfFunctionPlot[548].Y = 0.384

	pointsOfFunctionPlot[549].X = -0.86
	pointsOfFunctionPlot[549].Y = 0.388

	pointsOfFunctionPlot[550].X = -0.85
	pointsOfFunctionPlot[550].Y = 0.393

	pointsOfFunctionPlot[551].X = -0.84
	pointsOfFunctionPlot[551].Y = 0.397

	pointsOfFunctionPlot[552].X = -0.83
	pointsOfFunctionPlot[552].Y = 0.401

	pointsOfFunctionPlot[553].X = -0.82
	pointsOfFunctionPlot[553].Y = 0.406

	pointsOfFunctionPlot[554].X = -0.81
	pointsOfFunctionPlot[554].Y = 0.41

	pointsOfFunctionPlot[555].X = -0.8
	pointsOfFunctionPlot[555].Y = 0.415

	pointsOfFunctionPlot[556].X = -0.79
	pointsOfFunctionPlot[556].Y = 0.419

	pointsOfFunctionPlot[557].X = -0.78
	pointsOfFunctionPlot[557].Y = 0.424

	pointsOfFunctionPlot[558].X = -0.77
	pointsOfFunctionPlot[558].Y = 0.429

	pointsOfFunctionPlot[559].X = -0.76
	pointsOfFunctionPlot[559].Y = 0.433

	pointsOfFunctionPlot[560].X = -0.75
	pointsOfFunctionPlot[560].Y = 0.438

	pointsOfFunctionPlot[561].X = -0.74
	pointsOfFunctionPlot[561].Y = 0.443

	pointsOfFunctionPlot[562].X = -0.73
	pointsOfFunctionPlot[562].Y = 0.448

	pointsOfFunctionPlot[563].X = -0.72
	pointsOfFunctionPlot[563].Y = 0.453

	pointsOfFunctionPlot[564].X = -0.71
	pointsOfFunctionPlot[564].Y = 0.458

	pointsOfFunctionPlot[565].X = -0.7
	pointsOfFunctionPlot[565].Y = 0.463

	pointsOfFunctionPlot[566].X = -0.69
	pointsOfFunctionPlot[566].Y = 0.468

	pointsOfFunctionPlot[567].X = -0.68
	pointsOfFunctionPlot[567].Y = 0.473

	pointsOfFunctionPlot[568].X = -0.67
	pointsOfFunctionPlot[568].Y = 0.479

	pointsOfFunctionPlot[569].X = -0.66
	pointsOfFunctionPlot[569].Y = 0.484

	pointsOfFunctionPlot[570].X = -0.65
	pointsOfFunctionPlot[570].Y = 0.489

	pointsOfFunctionPlot[571].X = -0.64
	pointsOfFunctionPlot[571].Y = 0.495

	pointsOfFunctionPlot[572].X = -0.63
	pointsOfFunctionPlot[572].Y = 0.5

	pointsOfFunctionPlot[573].X = -0.62
	pointsOfFunctionPlot[573].Y = 0.506

	pointsOfFunctionPlot[574].X = -0.61
	pointsOfFunctionPlot[574].Y = 0.511

	pointsOfFunctionPlot[575].X = -0.6
	pointsOfFunctionPlot[575].Y = 0.517

	pointsOfFunctionPlot[574].X = -0.59
	pointsOfFunctionPlot[574].Y = 0.523

	pointsOfFunctionPlot[573].X = -0.58
	pointsOfFunctionPlot[573].Y = 0.528

	pointsOfFunctionPlot[572].X = -0.57
	pointsOfFunctionPlot[572].Y = 0.534

	pointsOfFunctionPlot[571].X = -0.56
	pointsOfFunctionPlot[571].Y = 0.54

	pointsOfFunctionPlot[570].X = -0.55
	pointsOfFunctionPlot[570].Y = 0.546

	pointsOfFunctionPlot[571].X = -0.54
	pointsOfFunctionPlot[571].Y = 0.552

	pointsOfFunctionPlot[572].X = -0.53
	pointsOfFunctionPlot[572].Y = 0.558

	pointsOfFunctionPlot[573].X = -0.52
	pointsOfFunctionPlot[573].Y = 0.564

	pointsOfFunctionPlot[574].X = -0.51
	pointsOfFunctionPlot[574].Y = 0.571

	pointsOfFunctionPlot[575].X = -0.5
	pointsOfFunctionPlot[575].Y = 0.577

	pointsOfFunctionPlot[576].X = -0.49
	pointsOfFunctionPlot[576].Y = 0.583

	pointsOfFunctionPlot[577].X = -0.48
	pointsOfFunctionPlot[577].Y = 0.59

	pointsOfFunctionPlot[578].X = -0.47
	pointsOfFunctionPlot[578].Y = 0.596

	pointsOfFunctionPlot[579].X = -0.46
	pointsOfFunctionPlot[579].Y = 0.603

	pointsOfFunctionPlot[580].X = -0.45
	pointsOfFunctionPlot[580].Y = 0.61

	pointsOfFunctionPlot[581].X = -0.44
	pointsOfFunctionPlot[581].Y = 0.616

	pointsOfFunctionPlot[582].X = -0.43
	pointsOfFunctionPlot[582].Y = 0.623

	pointsOfFunctionPlot[583].X = -0.42
	pointsOfFunctionPlot[583].Y = 0.63

	pointsOfFunctionPlot[584].X = -0.41
	pointsOfFunctionPlot[584].Y = 0.637

	pointsOfFunctionPlot[585].X = -0.4
	pointsOfFunctionPlot[585].Y = 0.644

	pointsOfFunctionPlot[586].X = -0.39
	pointsOfFunctionPlot[586].Y = 0.651

	pointsOfFunctionPlot[587].X = -0.38
	pointsOfFunctionPlot[587].Y = 0.658

	pointsOfFunctionPlot[588].X = -0.37
	pointsOfFunctionPlot[588].Y = 0.666

	pointsOfFunctionPlot[589].X = -0.36
	pointsOfFunctionPlot[589].Y = 0.673

	pointsOfFunctionPlot[590].X = -0.35
	pointsOfFunctionPlot[590].Y = 0.68

	pointsOfFunctionPlot[591].X = -0.34
	pointsOfFunctionPlot[591].Y = 0.688

	pointsOfFunctionPlot[592].X = -0.33
	pointsOfFunctionPlot[592].Y = 0.695

	pointsOfFunctionPlot[593].X = -0.32
	pointsOfFunctionPlot[593].Y = 0.703

	pointsOfFunctionPlot[594].X = -0.31
	pointsOfFunctionPlot[594].Y = 0.711

	pointsOfFunctionPlot[595].X = -0.3
	pointsOfFunctionPlot[595].Y = 0.719

	pointsOfFunctionPlot[596].X = -0.29
	pointsOfFunctionPlot[596].Y = 0.727

	pointsOfFunctionPlot[597].X = -0.28
	pointsOfFunctionPlot[597].Y = 0.735

	pointsOfFunctionPlot[598].X = -0.27
	pointsOfFunctionPlot[598].Y = 0.743

	pointsOfFunctionPlot[599].X = -0.26
	pointsOfFunctionPlot[599].Y = 0.751

	pointsOfFunctionPlot[600].X = -0.25
	pointsOfFunctionPlot[600].Y = 0.759

	pointsOfFunctionPlot[601].X = -0.24
	pointsOfFunctionPlot[601].Y = 0.768

	pointsOfFunctionPlot[602].X = -0.23
	pointsOfFunctionPlot[602].Y = 0.776

	pointsOfFunctionPlot[603].X = -0.22
	pointsOfFunctionPlot[603].Y = 0.785

	pointsOfFunctionPlot[604].X = -0.21
	pointsOfFunctionPlot[604].Y = 0.795

	pointsOfFunctionPlot[605].X = -0.2
	pointsOfFunctionPlot[605].Y = 0.802

	pointsOfFunctionPlot[606].X = -0.19
	pointsOfFunctionPlot[606].Y = 0.811

	pointsOfFunctionPlot[607].X = -0.18
	pointsOfFunctionPlot[607].Y = 0.82

	pointsOfFunctionPlot[608].X = -0.17
	pointsOfFunctionPlot[608].Y = 0.829

	pointsOfFunctionPlot[609].X = -0.16
	pointsOfFunctionPlot[609].Y = 0.838

	pointsOfFunctionPlot[610].X = -0.15
	pointsOfFunctionPlot[610].Y = 0.848

	pointsOfFunctionPlot[611].X = -0.14
	pointsOfFunctionPlot[611].Y = 0.857

	pointsOfFunctionPlot[612].X = -0.13
	pointsOfFunctionPlot[612].Y = 0.866

	pointsOfFunctionPlot[613].X = -0.12
	pointsOfFunctionPlot[613].Y = 0.876

	pointsOfFunctionPlot[614].X = -0.11
	pointsOfFunctionPlot[614].Y = 0.886

	pointsOfFunctionPlot[615].X = -0.1
	pointsOfFunctionPlot[615].Y = 0.896

	pointsOfFunctionPlot[616].X = -0.09
	pointsOfFunctionPlot[616].Y = 0.905

	pointsOfFunctionPlot[617].X = -0.08
	pointsOfFunctionPlot[617].Y = 0.915

	pointsOfFunctionPlot[618].X = -0.07
	pointsOfFunctionPlot[618].Y = 0.926

	pointsOfFunctionPlot[619].X = -0.06
	pointsOfFunctionPlot[619].Y = 0.936

	pointsOfFunctionPlot[620].X = -0.05
	pointsOfFunctionPlot[620].Y = 0.946

	pointsOfFunctionPlot[621].X = -0.04
	pointsOfFunctionPlot[621].Y = 0.957

	pointsOfFunctionPlot[622].X = -0.03
	pointsOfFunctionPlot[622].Y = 0.967

	pointsOfFunctionPlot[623].X = -0.02
	pointsOfFunctionPlot[623].Y = 0.978

	pointsOfFunctionPlot[624].X = -0.01
	pointsOfFunctionPlot[624].Y = 0.989

	pointsOfFunctionPlot[625].X = 0.0
	pointsOfFunctionPlot[625].Y = 1.0

	pointsOfFunctionPlot[626].X = 0.01
	pointsOfFunctionPlot[626].Y = 1.011

	pointsOfFunctionPlot[627].X = 0.02
	pointsOfFunctionPlot[627].Y = 1.022

	pointsOfFunctionPlot[628].X = 0.03
	pointsOfFunctionPlot[628].Y = 1.033

	pointsOfFunctionPlot[629].X = 0.04
	pointsOfFunctionPlot[629].Y = 1.044

	pointsOfFunctionPlot[630].X = 0.05
	pointsOfFunctionPlot[630].Y = 1.056

	pointsOfFunctionPlot[631].X = 0.06
	pointsOfFunctionPlot[631].Y = 1.068

	pointsOfFunctionPlot[632].X = 0.07
	pointsOfFunctionPlot[632].Y = 1.079

	pointsOfFunctionPlot[633].X = 0.08
	pointsOfFunctionPlot[633].Y = 1.091

	pointsOfFunctionPlot[634].X = 0.09
	pointsOfFunctionPlot[634].Y = 1.103

	pointsOfFunctionPlot[635].X = 0.1
	pointsOfFunctionPlot[635].Y = 1.116

	pointsOfFunctionPlot[636].X = 0.11
	pointsOfFunctionPlot[636].Y = 1.128

	pointsOfFunctionPlot[637].X = 0.12
	pointsOfFunctionPlot[637].Y = 1.14

	pointsOfFunctionPlot[638].X = 0.13
	pointsOfFunctionPlot[638].Y = 1.153

	pointsOfFunctionPlot[639].X = 0.14
	pointsOfFunctionPlot[639].Y = 1.166

	pointsOfFunctionPlot[640].X = 0.15
	pointsOfFunctionPlot[640].Y = 1.179

	pointsOfFunctionPlot[641].X = 0.16
	pointsOfFunctionPlot[641].Y = 1.192

	pointsOfFunctionPlot[642].X = 0.17
	pointsOfFunctionPlot[642].Y = 1.205

	pointsOfFunctionPlot[643].X = 0.18
	pointsOfFunctionPlot[643].Y = 1.218

	pointsOfFunctionPlot[644].X = 0.19
	pointsOfFunctionPlot[644].Y = 1.232

	pointsOfFunctionPlot[645].X = 0.2
	pointsOfFunctionPlot[645].Y = 1.245

	pointsOfFunctionPlot[646].X = 0.21
	pointsOfFunctionPlot[646].Y = 1.259

	pointsOfFunctionPlot[647].X = 0.22
	pointsOfFunctionPlot[647].Y = 1.273

	pointsOfFunctionPlot[648].X = 0.23
	pointsOfFunctionPlot[648].Y = 1.287

	pointsOfFunctionPlot[649].X = 0.24
	pointsOfFunctionPlot[649].Y = 1.301

	pointsOfFunctionPlot[650].X = 0.25
	pointsOfFunctionPlot[650].Y = 1.316

	pointsOfFunctionPlot[651].X = 0.26
	pointsOfFunctionPlot[651].Y = 1.33

	pointsOfFunctionPlot[652].X = 0.27
	pointsOfFunctionPlot[652].Y = 1.345

	pointsOfFunctionPlot[653].X = 0.28
	pointsOfFunctionPlot[653].Y = 1.36

	pointsOfFunctionPlot[654].X = 0.29
	pointsOfFunctionPlot[654].Y = 1.375

	pointsOfFunctionPlot[655].X = 0.3
	pointsOfFunctionPlot[655].Y = 1.39

	pointsOfFunctionPlot[656].X = 0.31
	pointsOfFunctionPlot[656].Y = 1.405

	pointsOfFunctionPlot[657].X = 0.32
	pointsOfFunctionPlot[657].Y = 1.421

	pointsOfFunctionPlot[658].X = 0.33
	pointsOfFunctionPlot[658].Y = 1.437

	pointsOfFunctionPlot[659].X = 0.34
	pointsOfFunctionPlot[659].Y = 1.452

	pointsOfFunctionPlot[660].X = 0.35
	pointsOfFunctionPlot[660].Y = 1.468

	pointsOfFunctionPlot[661].X = 0.36
	pointsOfFunctionPlot[661].Y = 1.485

	pointsOfFunctionPlot[662].X = 0.37
	pointsOfFunctionPlot[662].Y = 1.501

	pointsOfFunctionPlot[663].X = 0.38
	pointsOfFunctionPlot[663].Y = 1.518

	pointsOfFunctionPlot[664].X = 0.39
	pointsOfFunctionPlot[664].Y = 1.534

	pointsOfFunctionPlot[665].X = 0.40
	pointsOfFunctionPlot[665].Y = 1.551

	pointsOfFunctionPlot[666].X = 0.41
	pointsOfFunctionPlot[666].Y = 1.569

	pointsOfFunctionPlot[667].X = 0.42
	pointsOfFunctionPlot[667].Y = 1.586

	pointsOfFunctionPlot[668].X = 0.43
	pointsOfFunctionPlot[668].Y = 1.603

	pointsOfFunctionPlot[669].X = 0.44
	pointsOfFunctionPlot[669].Y = 1.621

	pointsOfFunctionPlot[670].X = 0.45
	pointsOfFunctionPlot[670].Y = 1.639

	pointsOfFunctionPlot[671].X = 0.46
	pointsOfFunctionPlot[671].Y = 1.657

	pointsOfFunctionPlot[672].X = 0.47
	pointsOfFunctionPlot[672].Y = 1.675

	pointsOfFunctionPlot[673].X = 0.48
	pointsOfFunctionPlot[673].Y = 1.694

	pointsOfFunctionPlot[674].X = 0.49
	pointsOfFunctionPlot[674].Y = 1.713

	pointsOfFunctionPlot[675].X = 0.5
	pointsOfFunctionPlot[675].Y = 1.732

	pointsOfFunctionPlot[676].X = 0.51
	pointsOfFunctionPlot[676].Y = 1.751

	pointsOfFunctionPlot[677].X = 0.52
	pointsOfFunctionPlot[677].Y = 1.77

	pointsOfFunctionPlot[678].X = 0.53
	pointsOfFunctionPlot[678].Y = 1.79

	pointsOfFunctionPlot[679].X = 0.54
	pointsOfFunctionPlot[679].Y = 1.809

	pointsOfFunctionPlot[680].X = 0.55
	pointsOfFunctionPlot[680].Y = 1.829

	pointsOfFunctionPlot[681].X = 0.56
	pointsOfFunctionPlot[681].Y = 1.85

	pointsOfFunctionPlot[682].X = 0.57
	pointsOfFunctionPlot[682].Y = 1.87

	pointsOfFunctionPlot[683].X = 0.58
	pointsOfFunctionPlot[683].Y = 1.891

	pointsOfFunctionPlot[684].X = 0.59
	pointsOfFunctionPlot[684].Y = 1.912

	pointsOfFunctionPlot[685].X = 0.6
	pointsOfFunctionPlot[685].Y = 1.933

	pointsOfFunctionPlot[686].X = 0.61
	pointsOfFunctionPlot[686].Y = 1.954

	pointsOfFunctionPlot[687].X = 0.62
	pointsOfFunctionPlot[687].Y = 1.976

	pointsOfFunctionPlot[688].X = 0.63
	pointsOfFunctionPlot[688].Y = 1.998

	pointsOfFunctionPlot[689].X = 0.64
	pointsOfFunctionPlot[689].Y = 2.02

	pointsOfFunctionPlot[690].X = 0.65
	pointsOfFunctionPlot[690].Y = 2.042

	pointsOfFunctionPlot[691].X = 0.66
	pointsOfFunctionPlot[691].Y = 2.064

	pointsOfFunctionPlot[692].X = 0.67
	pointsOfFunctionPlot[692].Y = 2.087

	pointsOfFunctionPlot[693].X = 0.68
	pointsOfFunctionPlot[693].Y = 2.11

	pointsOfFunctionPlot[694].X = 0.69
	pointsOfFunctionPlot[694].Y = 2.134

	pointsOfFunctionPlot[695].X = 0.7
	pointsOfFunctionPlot[695].Y = 2.157

	pointsOfFunctionPlot[696].X = 0.71
	pointsOfFunctionPlot[696].Y = 2.181

	pointsOfFunctionPlot[697].X = 0.72
	pointsOfFunctionPlot[697].Y = 2.205

	pointsOfFunctionPlot[698].X = 0.73
	pointsOfFunctionPlot[698].Y = 2.23

	pointsOfFunctionPlot[699].X = 0.74
	pointsOfFunctionPlot[699].Y = 2.254

	pointsOfFunctionPlot[700].X = 0.75
	pointsOfFunctionPlot[700].Y = 2.279

	pointsOfFunctionPlot[701].X = 0.76
	pointsOfFunctionPlot[701].Y = 2.304

	pointsOfFunctionPlot[702].X = 0.77
	pointsOfFunctionPlot[702].Y = 2.33

	pointsOfFunctionPlot[703].X = 0.78
	pointsOfFunctionPlot[703].Y = 2.355

	pointsOfFunctionPlot[704].X = 0.79
	pointsOfFunctionPlot[704].Y = 2.381

	pointsOfFunctionPlot[705].X = 0.8
	pointsOfFunctionPlot[705].Y = 2.408

	pointsOfFunctionPlot[706].X = 0.81
	pointsOfFunctionPlot[706].Y = 2.434

	pointsOfFunctionPlot[707].X = 0.82
	pointsOfFunctionPlot[707].Y = 2.461

	pointsOfFunctionPlot[708].X = 0.83
	pointsOfFunctionPlot[708].Y = 2.488

	pointsOfFunctionPlot[709].X = 0.84
	pointsOfFunctionPlot[709].Y = 2.516

	pointsOfFunctionPlot[710].X = 0.85
	pointsOfFunctionPlot[710].Y = 2.544

	pointsOfFunctionPlot[711].X = 0.86
	pointsOfFunctionPlot[711].Y = 2.572

	pointsOfFunctionPlot[712].X = 0.87
	pointsOfFunctionPlot[712].Y = 2.6

	pointsOfFunctionPlot[713].X = 0.88
	pointsOfFunctionPlot[713].Y = 2.629

	pointsOfFunctionPlot[714].X = 0.89
	pointsOfFunctionPlot[714].Y = 2.658

	pointsOfFunctionPlot[715].X = 0.9
	pointsOfFunctionPlot[715].Y = 2.687

	pointsOfFunctionPlot[716].X = 0.91
	pointsOfFunctionPlot[716].Y = 2.717

	pointsOfFunctionPlot[717].X = 0.92
	pointsOfFunctionPlot[717].Y = 2.747

	pointsOfFunctionPlot[718].X = 0.93
	pointsOfFunctionPlot[718].Y = 2.777

	pointsOfFunctionPlot[719].X = 0.94
	pointsOfFunctionPlot[719].Y = 2.808

	pointsOfFunctionPlot[720].X = 0.95
	pointsOfFunctionPlot[720].Y = 2.839

	pointsOfFunctionPlot[721].X = 0.96
	pointsOfFunctionPlot[721].Y = 2.871

	pointsOfFunctionPlot[722].X = 0.97
	pointsOfFunctionPlot[722].Y = 2.902

	pointsOfFunctionPlot[723].X = 0.98
	pointsOfFunctionPlot[723].Y = 2.934

	pointsOfFunctionPlot[724].X = 0.99
	pointsOfFunctionPlot[724].Y = 2.967

	pointsOfFunctionPlot[725].X = 1.0
	pointsOfFunctionPlot[725].Y = 3.0

	pointsOfFunctionPlot[726].X = 1.01
	pointsOfFunctionPlot[726].Y = 3.033

	pointsOfFunctionPlot[727].X = 1.02
	pointsOfFunctionPlot[727].Y = 3.066

	pointsOfFunctionPlot[728].X = 1.03
	pointsOfFunctionPlot[728].Y = 3.1

	pointsOfFunctionPlot[729].X = 1.04
	pointsOfFunctionPlot[729].Y = 3.134

	pointsOfFunctionPlot[730].X = 1.05
	pointsOfFunctionPlot[730].Y = 3.169

	pointsOfFunctionPlot[731].X = 1.06
	pointsOfFunctionPlot[731].Y = 3.204

	pointsOfFunctionPlot[732].X = 1.07
	pointsOfFunctionPlot[732].Y = 3.239

	pointsOfFunctionPlot[733].X = 1.08
	pointsOfFunctionPlot[733].Y = 3.275

	pointsOfFunctionPlot[734].X = 1.09
	pointsOfFunctionPlot[734].Y = 3.311

	pointsOfFunctionPlot[735].X = 1.1
	pointsOfFunctionPlot[735].Y = 3.348

	pointsOfFunctionPlot[736].X = 1.11
	pointsOfFunctionPlot[736].Y = 3.385

	pointsOfFunctionPlot[737].X = 1.12
	pointsOfFunctionPlot[737].Y = 3.422

	pointsOfFunctionPlot[738].X = 1.13
	pointsOfFunctionPlot[738].Y = 3.46

	pointsOfFunctionPlot[739].X = 1.14
	pointsOfFunctionPlot[739].Y = 3.498

	pointsOfFunctionPlot[740].X = 1.15
	pointsOfFunctionPlot[740].Y = 3.537

	pointsOfFunctionPlot[741].X = 1.16
	pointsOfFunctionPlot[741].Y = 3.576

	pointsOfFunctionPlot[742].X = 1.17
	pointsOfFunctionPlot[742].Y = 3.616

	pointsOfFunctionPlot[743].X = 1.18
	pointsOfFunctionPlot[743].Y = 3.656

	pointsOfFunctionPlot[744].X = 1.19
	pointsOfFunctionPlot[744].Y = 3.696

	pointsOfFunctionPlot[745].X = 1.2
	pointsOfFunctionPlot[745].Y = 3.737

	pointsOfFunctionPlot[746].X = 1.21
	pointsOfFunctionPlot[746].Y = 3.778

	pointsOfFunctionPlot[747].X = 1.22
	pointsOfFunctionPlot[747].Y = 3.82

	pointsOfFunctionPlot[748].X = 1.23
	pointsOfFunctionPlot[748].Y = 3.862

	pointsOfFunctionPlot[749].X = 1.24
	pointsOfFunctionPlot[749].Y = 3.905

	pointsOfFunctionPlot[750].X = 1.25
	pointsOfFunctionPlot[750].Y = 3.948

	pointsOfFunctionPlot[751].X = 1.26
	pointsOfFunctionPlot[751].Y = 3.991

	pointsOfFunctionPlot[752].X = 1.27
	pointsOfFunctionPlot[752].Y = 4.035

	pointsOfFunctionPlot[753].X = 1.28
	pointsOfFunctionPlot[753].Y = 4.08

	pointsOfFunctionPlot[754].X = 1.29
	pointsOfFunctionPlot[754].Y = 4.125

	pointsOfFunctionPlot[755].X = 1.3
	pointsOfFunctionPlot[755].Y = 4.171

	pointsOfFunctionPlot[756].X = 1.31
	pointsOfFunctionPlot[756].Y = 4.217

	pointsOfFunctionPlot[757].X = 1.32
	pointsOfFunctionPlot[757].Y = 4.263

	pointsOfFunctionPlot[758].X = 1.33
	pointsOfFunctionPlot[758].Y = 4.31

	pointsOfFunctionPlot[759].X = 1.34
	pointsOfFunctionPlot[759].Y = 4.358

	pointsOfFunctionPlot[760].X = 1.35
	pointsOfFunctionPlot[760].Y = 4.406

	pointsOfFunctionPlot[761].X = 1.36
	pointsOfFunctionPlot[761].Y = 4.455

	pointsOfFunctionPlot[762].X = 1.37
	pointsOfFunctionPlot[762].Y = 4.504

	pointsOfFunctionPlot[763].X = 1.38
	pointsOfFunctionPlot[763].Y = 4.554

	pointsOfFunctionPlot[764].X = 1.39
	pointsOfFunctionPlot[764].Y = 4.604

	pointsOfFunctionPlot[765].X = 1.40
	pointsOfFunctionPlot[765].Y = 4.655

	pointsOfFunctionPlot[766].X = 1.41
	pointsOfFunctionPlot[766].Y = 4.707

	pointsOfFunctionPlot[767].X = 1.42
	pointsOfFunctionPlot[767].Y = 4.759

	pointsOfFunctionPlot[768].X = 1.43
	pointsOfFunctionPlot[768].Y = 4.811

	pointsOfFunctionPlot[769].X = 1.44
	pointsOfFunctionPlot[769].Y = 4.864

	pointsOfFunctionPlot[770].X = 1.45
	pointsOfFunctionPlot[770].Y = 4.918

	pointsOfFunctionPlot[771].X = 1.46
	pointsOfFunctionPlot[771].Y = 4.972

	pointsOfFunctionPlot[772].X = 1.47
	pointsOfFunctionPlot[772].Y = 5.027

	pointsOfFunctionPlot[773].X = 1.48
	pointsOfFunctionPlot[773].Y = 5.083

	pointsOfFunctionPlot[774].X = 1.49
	pointsOfFunctionPlot[774].Y = 5.139

	pointsOfFunctionPlot[775].X = 1.5
	pointsOfFunctionPlot[775].Y = 5.196

	pointsOfFunctionPlot[776].X = 1.51
	pointsOfFunctionPlot[776].Y = 5.253

	pointsOfFunctionPlot[777].X = 1.52
	pointsOfFunctionPlot[777].Y = 5.311

	pointsOfFunctionPlot[778].X = 1.53
	pointsOfFunctionPlot[778].Y = 5.37

	pointsOfFunctionPlot[779].X = 1.54
	pointsOfFunctionPlot[779].Y = 5.429

	pointsOfFunctionPlot[780].X = 1.55
	pointsOfFunctionPlot[780].Y = 5.489

	pointsOfFunctionPlot[781].X = 1.56
	pointsOfFunctionPlot[781].Y = 5.55

	pointsOfFunctionPlot[782].X = 1.57
	pointsOfFunctionPlot[782].Y = 5.611

	pointsOfFunctionPlot[783].X = 1.58
	pointsOfFunctionPlot[783].Y = 5.673

	pointsOfFunctionPlot[784].X = 1.59
	pointsOfFunctionPlot[784].Y = 5.736

	pointsOfFunctionPlot[785].X = 1.6
	pointsOfFunctionPlot[785].Y = 5.799

	pointsOfFunctionPlot[786].X = 1.61
	pointsOfFunctionPlot[786].Y = 5.863

	pointsOfFunctionPlot[787].X = 1.62
	pointsOfFunctionPlot[787].Y = 5.928

	pointsOfFunctionPlot[788].X = 1.63
	pointsOfFunctionPlot[788].Y = 5.993

	pointsOfFunctionPlot[789].X = 1.64
	pointsOfFunctionPlot[789].Y = 6.06

	pointsOfFunctionPlot[790].X = 1.65
	pointsOfFunctionPlot[790].Y = 6.127

	pointsOfFunctionPlot[791].X = 1.66
	pointsOfFunctionPlot[791].Y = 6.194

	pointsOfFunctionPlot[792].X = 1.67
	pointsOfFunctionPlot[792].Y = 6.263

	pointsOfFunctionPlot[793].X = 1.68
	pointsOfFunctionPlot[793].Y = 6.332

	pointsOfFunctionPlot[794].X = 1.69
	pointsOfFunctionPlot[794].Y = 6.402

	pointsOfFunctionPlot[795].X = 1.7
	pointsOfFunctionPlot[795].Y = 6.473

	pointsOfFunctionPlot[796].X = 1.71
	pointsOfFunctionPlot[796].Y = 6.544

	pointsOfFunctionPlot[797].X = 1.72
	pointsOfFunctionPlot[797].Y = 6.616

	pointsOfFunctionPlot[798].X = 1.73
	pointsOfFunctionPlot[798].Y = 6.689

	pointsOfFunctionPlot[799].X = 1.74
	pointsOfFunctionPlot[799].Y = 6.763

	pointsOfFunctionPlot[800].X = 1.75
	pointsOfFunctionPlot[800].Y = 6.838

	pointsOfFunctionPlot[801].X = 1.76
	pointsOfFunctionPlot[801].Y = 6.914

	pointsOfFunctionPlot[802].X = 1.77
	pointsOfFunctionPlot[802].Y = 6.99

	pointsOfFunctionPlot[803].X = 1.78
	pointsOfFunctionPlot[803].Y = 7.067

	pointsOfFunctionPlot[804].X = 1.79
	pointsOfFunctionPlot[804].Y = 7.145

	pointsOfFunctionPlot[805].X = 1.8
	pointsOfFunctionPlot[805].Y = 7.224

	pointsOfFunctionPlot[806].X = 1.81
	pointsOfFunctionPlot[806].Y = 7.304

	pointsOfFunctionPlot[807].X = 1.82
	pointsOfFunctionPlot[807].Y = 7.385

	pointsOfFunctionPlot[808].X = 1.83
	pointsOfFunctionPlot[808].Y = 7.466

	pointsOfFunctionPlot[809].X = 1.84
	pointsOfFunctionPlot[809].Y = 7.632

	pointsOfFunctionPlot[810].X = 1.85
	pointsOfFunctionPlot[810].Y = 7.632

	pointsOfFunctionPlot[811].X = 1.86
	pointsOfFunctionPlot[811].Y = 7.716

	pointsOfFunctionPlot[812].X = 1.87
	pointsOfFunctionPlot[812].Y = 7.802

	pointsOfFunctionPlot[813].X = 1.88
	pointsOfFunctionPlot[813].Y = 7.888

	pointsOfFunctionPlot[814].X = 1.89
	pointsOfFunctionPlot[814].Y = 7.975

	pointsOfFunctionPlot[815].X = 1.9
	pointsOfFunctionPlot[815].Y = 8.063

	pointsOfFunctionPlot[816].X = 1.91
	pointsOfFunctionPlot[816].Y = 8.152

	pointsOfFunctionPlot[817].X = 1.92
	pointsOfFunctionPlot[817].Y = 8.242

	pointsOfFunctionPlot[818].X = 1.93
	pointsOfFunctionPlot[818].Y = 8.333

	pointsOfFunctionPlot[819].X = 1.94
	pointsOfFunctionPlot[819].Y = 8.425

	pointsOfFunctionPlot[820].X = 1.95
	pointsOfFunctionPlot[820].Y = 8.519

	pointsOfFunctionPlot[821].X = 1.96
	pointsOfFunctionPlot[821].Y = 8.613

	pointsOfFunctionPlot[822].X = 1.97
	pointsOfFunctionPlot[822].Y = 8.708

	pointsOfFunctionPlot[823].X = 1.98
	pointsOfFunctionPlot[823].Y = 8.804

	pointsOfFunctionPlot[824].X = 1.99
	pointsOfFunctionPlot[824].Y = 8.901

	pointsOfFunctionPlot[825].X = 2.0
	pointsOfFunctionPlot[825].Y = 9.0

	pointsOfFunctionPlot[826].X = 2.01
	pointsOfFunctionPlot[826].Y = 9.099

	pointsOfFunctionPlot[827].X = 2.02
	pointsOfFunctionPlot[827].Y = 9.199

	pointsOfFunctionPlot[828].X = 2.03
	pointsOfFunctionPlot[828].Y = 9.301

	pointsOfFunctionPlot[829].X = 2.04
	pointsOfFunctionPlot[829].Y = 9.404

	pointsOfFunctionPlot[830].X = 2.05
	pointsOfFunctionPlot[830].Y = 9.508

	pointsOfFunctionPlot[831].X = 2.06
	pointsOfFunctionPlot[831].Y = 9.613

	pointsOfFunctionPlot[832].X = 2.07
	pointsOfFunctionPlot[832].Y = 9.719

	pointsOfFunctionPlot[833].X = 2.08
	pointsOfFunctionPlot[833].Y = 9.826

	pointsOfFunctionPlot[834].X = 2.09
	pointsOfFunctionPlot[834].Y = 9.935

	pointsOfFunctionPlot[835].X = 2.1
	pointsOfFunctionPlot[835].Y = 10.045

	pointsOfFunctionPlot[836].X = 2.11
	pointsOfFunctionPlot[836].Y = 10.156

	pointsOfFunctionPlot[837].X = 2.12
	pointsOfFunctionPlot[837].Y = 10.268

	pointsOfFunctionPlot[838].X = 2.13
	pointsOfFunctionPlot[838].Y = 10.381

	pointsOfFunctionPlot[839].X = 2.14
	pointsOfFunctionPlot[839].Y = 10.496

	pointsOfFunctionPlot[840].X = 2.15
	pointsOfFunctionPlot[840].Y = 10.612

	pointsOfFunctionPlot[841].X = 2.16
	pointsOfFunctionPlot[841].Y = 10.729

	pointsOfFunctionPlot[842].X = 2.17
	pointsOfFunctionPlot[842].Y = 10.848

	pointsOfFunctionPlot[843].X = 2.18
	pointsOfFunctionPlot[843].Y = 10.967

	pointsOfFunctionPlot[844].X = 2.19
	pointsOfFunctionPlot[844].Y = 11.089

	pointsOfFunctionPlot[845].X = 2.2
	pointsOfFunctionPlot[845].Y = 11.211

	pointsOfFunctionPlot[846].X = 2.21
	pointsOfFunctionPlot[846].Y = 11.335

	pointsOfFunctionPlot[847].X = 2.22
	pointsOfFunctionPlot[847].Y = 11.46

	pointsOfFunctionPlot[848].X = 2.23
	pointsOfFunctionPlot[848].Y = 11.587

	pointsOfFunctionPlot[849].X = 2.24
	pointsOfFunctionPlot[849].Y = 11.715

	pointsOfFunctionPlot[850].X = 2.25
	pointsOfFunctionPlot[850].Y = 11.844

	pointsOfFunctionPlot[851].X = 2.26
	pointsOfFunctionPlot[851].Y = 11.975

	pointsOfFunctionPlot[852].X = 2.27
	pointsOfFunctionPlot[852].Y = 12.107

	pointsOfFunctionPlot[853].X = 2.28
	pointsOfFunctionPlot[853].Y = 12.241

	pointsOfFunctionPlot[854].X = 2.29
	pointsOfFunctionPlot[854].Y = 12.376

	pointsOfFunctionPlot[855].X = 2.3
	pointsOfFunctionPlot[855].Y = 12.513

	pointsOfFunctionPlot[856].X = 2.31
	pointsOfFunctionPlot[856].Y = 12.651

	pointsOfFunctionPlot[857].X = 2.32
	pointsOfFunctionPlot[857].Y = 12.791

	pointsOfFunctionPlot[858].X = 2.33
	pointsOfFunctionPlot[858].Y = 12.932

	pointsOfFunctionPlot[859].X = 2.34
	pointsOfFunctionPlot[859].Y = 13.075

	pointsOfFunctionPlot[860].X = 2.35
	pointsOfFunctionPlot[860].Y = 13.22

	pointsOfFunctionPlot[861].X = 2.36
	pointsOfFunctionPlot[861].Y = 13.366

	pointsOfFunctionPlot[862].X = 2.37
	pointsOfFunctionPlot[862].Y = 13.513

	pointsOfFunctionPlot[863].X = 2.38
	pointsOfFunctionPlot[863].Y = 13.663

	pointsOfFunctionPlot[864].X = 2.39
	pointsOfFunctionPlot[864].Y = 13.814

	pointsOfFunctionPlot[865].X = 2.4
	pointsOfFunctionPlot[865].Y = 13.966

	pointsOfFunctionPlot[866].X = 2.41
	pointsOfFunctionPlot[866].Y = 14.12

	pointsOfFunctionPlot[867].X = 2.42
	pointsOfFunctionPlot[867].Y = 14.276

	pointsOfFunctionPlot[868].X = 2.43
	pointsOfFunctionPlot[868].Y = 14.434

	pointsOfFunctionPlot[869].X = 2.44
	pointsOfFunctionPlot[869].Y = 14.594

	pointsOfFunctionPlot[870].X = 2.45
	pointsOfFunctionPlot[870].Y = 14.755

	pointsOfFunctionPlot[871].X = 2.46
	pointsOfFunctionPlot[871].Y = 14.918

	pointsOfFunctionPlot[872].X = 2.47
	pointsOfFunctionPlot[872].Y = 15.083

	pointsOfFunctionPlot[873].X = 2.48
	pointsOfFunctionPlot[873].Y = 15.249

	pointsOfFunctionPlot[874].X = 2.49
	pointsOfFunctionPlot[874].Y = 15.418

	pointsOfFunctionPlot[875].X = 2.5
	pointsOfFunctionPlot[875].Y = 15.588

	pointsOfFunctionPlot[876].X = 2.51
	pointsOfFunctionPlot[876].Y = 15.76

	pointsOfFunctionPlot[877].X = 2.52
	pointsOfFunctionPlot[877].Y = 15.934

	pointsOfFunctionPlot[878].X = 2.53
	pointsOfFunctionPlot[878].Y = 16.11

	pointsOfFunctionPlot[879].X = 2.54
	pointsOfFunctionPlot[879].Y = 16.288

	pointsOfFunctionPlot[880].X = 2.55
	pointsOfFunctionPlot[880].Y = 16.468

	pointsOfFunctionPlot[881].X = 2.56
	pointsOfFunctionPlot[881].Y = 16.65

	pointsOfFunctionPlot[882].X = 2.57
	pointsOfFunctionPlot[882].Y = 16.834

	pointsOfFunctionPlot[883].X = 2.58
	pointsOfFunctionPlot[883].Y = 17.02

	pointsOfFunctionPlot[884].X = 2.59
	pointsOfFunctionPlot[884].Y = 17.208

	pointsOfFunctionPlot[885].X = 2.6
	pointsOfFunctionPlot[885].Y = 17.398

	pointsOfFunctionPlot[886].X = 2.61
	pointsOfFunctionPlot[886].Y = 17.59

	pointsOfFunctionPlot[887].X = 2.62
	pointsOfFunctionPlot[887].Y = 17.785

	pointsOfFunctionPlot[888].X = 2.63
	pointsOfFunctionPlot[888].Y = 17.981

	pointsOfFunctionPlot[889].X = 2.64
	pointsOfFunctionPlot[889].Y = 18.18

	pointsOfFunctionPlot[890].X = 2.65
	pointsOfFunctionPlot[890].Y = 18.381

	pointsOfFunctionPlot[891].X = 2.66
	pointsOfFunctionPlot[891].Y = 18.584

	pointsOfFunctionPlot[892].X = 2.67
	pointsOfFunctionPlot[892].Y = 18.789

	pointsOfFunctionPlot[893].X = 2.68
	pointsOfFunctionPlot[893].Y = 18.997

	pointsOfFunctionPlot[894].X = 2.69
	pointsOfFunctionPlot[894].Y = 19.206

	pointsOfFunctionPlot[895].X = 2.7
	pointsOfFunctionPlot[895].Y = 19.419

	pointsOfFunctionPlot[896].X = 2.71
	pointsOfFunctionPlot[896].Y = 19.633

	pointsOfFunctionPlot[897].X = 2.72
	pointsOfFunctionPlot[897].Y = 19.85

	pointsOfFunctionPlot[898].X = 2.73
	pointsOfFunctionPlot[898].Y = 20.069

	pointsOfFunctionPlot[899].X = 2.74
	pointsOfFunctionPlot[899].Y = 20.291

	pointsOfFunctionPlot[900].X = 2.75
	pointsOfFunctionPlot[900].Y = 20.515

	pointsOfFunctionPlot[901].X = 2.76
	pointsOfFunctionPlot[901].Y = 20.742

	pointsOfFunctionPlot[902].X = 2.77
	pointsOfFunctionPlot[902].Y = 20.971

	pointsOfFunctionPlot[903].X = 2.78
	pointsOfFunctionPlot[903].Y = 21.203

	pointsOfFunctionPlot[904].X = 2.79
	pointsOfFunctionPlot[904].Y = 21.437

	pointsOfFunctionPlot[905].X = 2.8
	pointsOfFunctionPlot[905].Y = 21.674

	pointsOfFunctionPlot[906].X = 2.81
	pointsOfFunctionPlot[906].Y = 21.913

	pointsOfFunctionPlot[907].X = 2.82
	pointsOfFunctionPlot[907].Y = 22.155

	pointsOfFunctionPlot[908].X = 2.83
	pointsOfFunctionPlot[908].Y = 22.4

	pointsOfFunctionPlot[909].X = 2.84
	pointsOfFunctionPlot[909].Y = 22.647

	pointsOfFunctionPlot[910].X = 2.85
	pointsOfFunctionPlot[910].Y = 22.897

	pointsOfFunctionPlot[911].X = 2.86
	pointsOfFunctionPlot[911].Y = 23.15

	pointsOfFunctionPlot[912].X = 2.87
	pointsOfFunctionPlot[912].Y = 23.406

	pointsOfFunctionPlot[913].X = 2.88
	pointsOfFunctionPlot[913].Y = 23.665

	pointsOfFunctionPlot[914].X = 2.89
	pointsOfFunctionPlot[914].Y = 23.926

	pointsOfFunctionPlot[915].X = 2.9
	pointsOfFunctionPlot[915].Y = 24.19

	pointsOfFunctionPlot[916].X = 2.91
	pointsOfFunctionPlot[916].Y = 24.458

	pointsOfFunctionPlot[917].X = 2.92
	pointsOfFunctionPlot[917].Y = 24.728

	pointsOfFunctionPlot[918].X = 2.93
	pointsOfFunctionPlot[918].Y = 25.001

	pointsOfFunctionPlot[919].X = 2.94
	pointsOfFunctionPlot[919].Y = 25.277

	pointsOfFunctionPlot[920].X = 2.95
	pointsOfFunctionPlot[920].Y = 25.556

	pointsOfFunctionPlot[921].X = 2.96
	pointsOfFunctionPlot[921].Y = 25.839

	pointsOfFunctionPlot[922].X = 2.97
	pointsOfFunctionPlot[922].Y = 26.124

	pointsOfFunctionPlot[922].X = 2.98
	pointsOfFunctionPlot[922].Y = 26.413

	pointsOfFunctionPlot[923].X = 2.99
	pointsOfFunctionPlot[923].Y = 26.705

	pointsOfFunctionPlot[924].X = 3.0
	pointsOfFunctionPlot[924].Y = 27.0

	pointsOfFunctionPlot[925].X = 3.01
	pointsOfFunctionPlot[925].Y = 27.298

	pointsOfFunctionPlot[926].X = 3.02
	pointsOfFunctionPlot[926].Y = 27.599

	pointsOfFunctionPlot[927].X = 3.03
	pointsOfFunctionPlot[927].Y = 27.904

	pointsOfFunctionPlot[928].X = 3.04
	pointsOfFunctionPlot[928].Y = 28.213

	pointsOfFunctionPlot[929].X = 3.05
	pointsOfFunctionPlot[929].Y = 28.524

	pointsOfFunctionPlot[930].X = 3.06
	pointsOfFunctionPlot[930].Y = 28.839

	pointsOfFunctionPlot[931].X = 3.07
	pointsOfFunctionPlot[931].Y = 29.158

	pointsOfFunctionPlot[932].X = 3.08
	pointsOfFunctionPlot[932].Y = 29.48

	pointsOfFunctionPlot[933].X = 3.09
	pointsOfFunctionPlot[933].Y = 29.806

	pointsOfFunctionPlot[934].X = 3.1
	pointsOfFunctionPlot[934].Y = 30.135

	pointsOfFunctionPlot[935].X = 3.11
	pointsOfFunctionPlot[935].Y = 30.468

	pointsOfFunctionPlot[936].X = 3.12
	pointsOfFunctionPlot[936].Y = 30.804

	pointsOfFunctionPlot[937].X = 3.13
	pointsOfFunctionPlot[937].Y = 31.145

	pointsOfFunctionPlot[938].X = 3.14
	pointsOfFunctionPlot[938].Y = 31.489

	pointsOfFunctionPlot[939].X = 3.15
	pointsOfFunctionPlot[939].Y = 31.837

	pointsOfFunctionPlot[940].X = 3.16
	pointsOfFunctionPlot[940].Y = 32.188

	pointsOfFunctionPlot[941].X = 3.17
	pointsOfFunctionPlot[941].Y = 32.544

	pointsOfFunctionPlot[942].X = 3.18
	pointsOfFunctionPlot[942].Y = 32.903

	pointsOfFunctionPlot[943].X = 3.19
	pointsOfFunctionPlot[943].Y = 33.267

	pointsOfFunctionPlot[944].X = 3.2
	pointsOfFunctionPlot[944].Y = 33.634

	pointsOfFunctionPlot[945].X = 3.21
	pointsOfFunctionPlot[945].Y = 34.006

	pointsOfFunctionPlot[946].X = 3.22
	pointsOfFunctionPlot[946].Y = 34.381

	pointsOfFunctionPlot[947].X = 3.23
	pointsOfFunctionPlot[947].Y = 34.761

	pointsOfFunctionPlot[948].X = 3.24
	pointsOfFunctionPlot[948].Y = 35.145

	pointsOfFunctionPlot[949].X = 3.25
	pointsOfFunctionPlot[949].Y = 35.534

	pointsOfFunctionPlot[950].X = 3.26
	pointsOfFunctionPlot[950].Y = 35.926

	pointsOfFunctionPlot[951].X = 3.27
	pointsOfFunctionPlot[951].Y = 36.323

	pointsOfFunctionPlot[952].X = 3.28
	pointsOfFunctionPlot[952].Y = 36.724

	pointsOfFunctionPlot[953].X = 3.29
	pointsOfFunctionPlot[953].Y = 37.13

	pointsOfFunctionPlot[954].X = 3.3
	pointsOfFunctionPlot[954].Y = 37.54

	pointsOfFunctionPlot[955].X = 3.31
	pointsOfFunctionPlot[955].Y = 37.955

	pointsOfFunctionPlot[956].X = 3.32
	pointsOfFunctionPlot[956].Y = 38.374

	pointsOfFunctionPlot[957].X = 3.33
	pointsOfFunctionPlot[957].Y = 38.798

	pointsOfFunctionPlot[958].X = 3.34
	pointsOfFunctionPlot[958].Y = 39.227

	pointsOfFunctionPlot[959].X = 3.35
	pointsOfFunctionPlot[959].Y = 39.66

	pointsOfFunctionPlot[960].X = 3.36
	pointsOfFunctionPlot[960].Y = 40.098

	pointsOfFunctionPlot[961].X = 3.37
	pointsOfFunctionPlot[961].Y = 40.541

	pointsOfFunctionPlot[962].X = 3.38
	pointsOfFunctionPlot[962].Y = 40.989

	pointsOfFunctionPlot[963].X = 3.39
	pointsOfFunctionPlot[963].Y = 41.442

	pointsOfFunctionPlot[964].X = 3.4
	pointsOfFunctionPlot[964].Y = 41.899

	pointsOfFunctionPlot[965].X = 3.41
	pointsOfFunctionPlot[965].Y = 42.362

	pointsOfFunctionPlot[966].X = 3.42
	pointsOfFunctionPlot[966].Y = 42.83

	pointsOfFunctionPlot[967].X = 3.43
	pointsOfFunctionPlot[967].Y = 43.303

	pointsOfFunctionPlot[968].X = 3.44
	pointsOfFunctionPlot[968].Y = 43.782

	pointsOfFunctionPlot[969].X = 3.45
	pointsOfFunctionPlot[969].Y = 44.265

	pointsOfFunctionPlot[970].X = 3.46
	pointsOfFunctionPlot[970].Y = 44.754

	pointsOfFunctionPlot[971].X = 3.47
	pointsOfFunctionPlot[971].Y = 45.249

	pointsOfFunctionPlot[972].X = 3.48
	pointsOfFunctionPlot[972].Y = 45.749

	pointsOfFunctionPlot[973].X = 3.49
	pointsOfFunctionPlot[973].Y = 46.254

	pointsOfFunctionPlot[974].X = 3.5
	pointsOfFunctionPlot[974].Y = 46.765

	pointsOfFunctionPlot[975].X = 3.51
	pointsOfFunctionPlot[975].Y = 47.282

	pointsOfFunctionPlot[976].X = 3.52
	pointsOfFunctionPlot[976].Y = 47.804

	pointsOfFunctionPlot[977].X = 3.53
	pointsOfFunctionPlot[977].Y = 48.332

	pointsOfFunctionPlot[978].X = 3.54
	pointsOfFunctionPlot[978].Y = 48.866

	pointsOfFunctionPlot[979].X = 3.55
	pointsOfFunctionPlot[979].Y = 49.406

	pointsOfFunctionPlot[980].X = 3.56
	pointsOfFunctionPlot[980].Y = 49.951

	pointsOfFunctionPlot[981].X = 3.57
	pointsOfFunctionPlot[981].Y = 50.503

	pointsOfFunctionPlot[982].X = 3.58
	pointsOfFunctionPlot[982].Y = 51.061

	pointsOfFunctionPlot[983].X = 3.59
	pointsOfFunctionPlot[983].Y = 51.625

	pointsOfFunctionPlot[984].X = 3.6
	pointsOfFunctionPlot[984].Y = 52.195

	pointsOfFunctionPlot[985].X = 3.61
	pointsOfFunctionPlot[985].Y = 52.772

	pointsOfFunctionPlot[986].X = 3.62
	pointsOfFunctionPlot[986].Y = 53.355

	pointsOfFunctionPlot[987].X = 3.63
	pointsOfFunctionPlot[987].Y = 53.944

	pointsOfFunctionPlot[988].X = 3.64
	pointsOfFunctionPlot[988].Y = 54.54

	pointsOfFunctionPlot[989].X = 3.65
	pointsOfFunctionPlot[989].Y = 55.143

	pointsOfFunctionPlot[990].X = 3.66
	pointsOfFunctionPlot[990].Y = 55.752

	pointsOfFunctionPlot[991].X = 3.67
	pointsOfFunctionPlot[991].Y = 56.368

	pointsOfFunctionPlot[992].X = 3.68
	pointsOfFunctionPlot[992].Y = 56.991

	pointsOfFunctionPlot[993].X = 3.69
	pointsOfFunctionPlot[993].Y = 57.62

	pointsOfFunctionPlot[994].X = 3.7
	pointsOfFunctionPlot[994].Y = 58.257

	pointsOfFunctionPlot[995].X = 3.71
	pointsOfFunctionPlot[995].Y = 58.9

	pointsOfFunctionPlot[996].X = 3.72
	pointsOfFunctionPlot[996].Y = 59.551

	pointsOfFunctionPlot[997].X = 3.73
	pointsOfFunctionPlot[997].Y = 60.209

	pointsOfFunctionPlot[998].X = 3.74
	pointsOfFunctionPlot[998].Y = 60.874

	pointsOfFunctionPlot[999].X = 3.75
	pointsOfFunctionPlot[999].Y = 61.546

	pointsOfFunctionPlot[1_000].X = 3.76
	pointsOfFunctionPlot[1_000].Y = 62.226

	pointsOfFunctionPlot[1_001].X = 3.77
	pointsOfFunctionPlot[1_001].Y = 62.914

	pointsOfFunctionPlot[1_002].X = 3.78
	pointsOfFunctionPlot[1_002].Y = 63.609

	pointsOfFunctionPlot[1_003].X = 3.79
	pointsOfFunctionPlot[1_003].Y = 64.311

	pointsOfFunctionPlot[1_004].X = 3.8
	pointsOfFunctionPlot[1_004].Y = 65.022

	pointsOfFunctionPlot[1_005].X = 3.81
	pointsOfFunctionPlot[1_005].Y = 65.74

	pointsOfFunctionPlot[1_006].X = 3.82
	pointsOfFunctionPlot[1_006].Y = 66.466

	pointsOfFunctionPlot[1_007].X = 3.83
	pointsOfFunctionPlot[1_007].Y = 67.2

	pointsOfFunctionPlot[1_008].X = 3.84
	pointsOfFunctionPlot[1_008].Y = 67.943

	pointsOfFunctionPlot[1_009].X = 3.85
	pointsOfFunctionPlot[1_009].Y = 68.693

	pointsOfFunctionPlot[1_010].X = 3.86
	pointsOfFunctionPlot[1_010].Y = 69.452

	pointsOfFunctionPlot[1_011].X = 3.87
	pointsOfFunctionPlot[1_011].Y = 70.219

	pointsOfFunctionPlot[1_012].X = 3.88
	pointsOfFunctionPlot[1_012].Y = 70.995

	pointsOfFunctionPlot[1_013].X = 3.89
	pointsOfFunctionPlot[1_013].Y = 71.779

	pointsOfFunctionPlot[1_014].X = 3.9
	pointsOfFunctionPlot[1_014].Y = 72.572

	pointsOfFunctionPlot[1_015].X = 3.91
	pointsOfFunctionPlot[1_015].Y = 73.374

	pointsOfFunctionPlot[1_016].X = 3.92
	pointsOfFunctionPlot[1_016].Y = 74.184

	pointsOfFunctionPlot[1_017].X = 3.93
	pointsOfFunctionPlot[1_017].Y = 75.004

	pointsOfFunctionPlot[1_018].X = 3.94
	pointsOfFunctionPlot[1_018].Y = 75.832

	pointsOfFunctionPlot[1_019].X = 3.95
	pointsOfFunctionPlot[1_019].Y = 76.67

	pointsOfFunctionPlot[1_020].X = 3.96
	pointsOfFunctionPlot[1_020].Y = 77.517

	pointsOfFunctionPlot[1_021].X = 3.97
	pointsOfFunctionPlot[1_021].Y = 78.373

	pointsOfFunctionPlot[1_022].X = 3.98
	pointsOfFunctionPlot[1_022].Y = 79.239

	pointsOfFunctionPlot[1_023].X = 3.99
	pointsOfFunctionPlot[1_023].Y = 80.115

	pointsOfFunctionPlot[1_024].X = 4.0
	pointsOfFunctionPlot[1_024].Y = 81.0

	pointsOfFunctionPlot[1_025].X = 4.01
	pointsOfFunctionPlot[1_025].Y = 81.894

	pointsOfFunctionPlot[1_026].X = 4.02
	pointsOfFunctionPlot[1_026].Y = 82.799

	pointsOfFunctionPlot[1_027].X = 4.03
	pointsOfFunctionPlot[1_027].Y = 83.714

	pointsOfFunctionPlot[1_028].X = 4.04
	pointsOfFunctionPlot[1_028].Y = 84.638

	pointsOfFunctionPlot[1_029].X = 4.05
	pointsOfFunctionPlot[1_029].Y = 85.573

	pointsOfFunctionPlot[1_030].X = 4.06
	pointsOfFunctionPlot[1_030].Y = 86.519

	pointsOfFunctionPlot[1_031].X = 4.07
	pointsOfFunctionPlot[1_031].Y = 87.474

	pointsOfFunctionPlot[1_032].X = 4.08
	pointsOfFunctionPlot[1_032].Y = 88.441

	pointsOfFunctionPlot[1_033].X = 4.09
	pointsOfFunctionPlot[1_033].Y = 89.418

	pointsOfFunctionPlot[1_034].X = 4.1
	pointsOfFunctionPlot[1_034].Y = 90.406

	pointsOfFunctionPlot[1_035].X = 4.11
	pointsOfFunctionPlot[1_035].Y = 91.404

	pointsOfFunctionPlot[1_036].X = 4.12
	pointsOfFunctionPlot[1_036].Y = 92.414

	pointsOfFunctionPlot[1_037].X = 4.13
	pointsOfFunctionPlot[1_037].Y = 93.435

	pointsOfFunctionPlot[1_038].X = 4.14
	pointsOfFunctionPlot[1_038].Y = 94.467

	pointsOfFunctionPlot[1_039].X = 4.15
	pointsOfFunctionPlot[1_039].Y = 95.511

	pointsOfFunctionPlot[1_040].X = 4.16
	pointsOfFunctionPlot[1_040].Y = 96.566

	pointsOfFunctionPlot[1_041].X = 4.17
	pointsOfFunctionPlot[1_041].Y = 97.632

	pointsOfFunctionPlot[1_042].X = 4.18
	pointsOfFunctionPlot[1_042].Y = 98.711

	pointsOfFunctionPlot[1_043].X = 4.19
	pointsOfFunctionPlot[1_043].Y = 99.801

	pointsOfFunctionPlot[1_044].X = 4.2
	pointsOfFunctionPlot[1_044].Y = 100.904

	pointsOfFunctionPlot[1_045].X = 4.21
	pointsOfFunctionPlot[1_045].Y = 102.018

	pointsOfFunctionPlot[1_046].X = 4.22
	pointsOfFunctionPlot[1_046].Y = 103.145

	pointsOfFunctionPlot[1_047].X = 4.23
	pointsOfFunctionPlot[1_047].Y = 104.285

	pointsOfFunctionPlot[1_048].X = 4.24
	pointsOfFunctionPlot[1_048].Y = 105.437

	pointsOfFunctionPlot[1_049].X = 4.25
	pointsOfFunctionPlot[1_049].Y = 106.602

	pointsOfFunctionPlot[1_050].X = 4.26
	pointsOfFunctionPlot[1_050].Y = 107.779

	pointsOfFunctionPlot[1_051].X = 4.27
	pointsOfFunctionPlot[1_051].Y = 108.97

	pointsOfFunctionPlot[1_052].X = 4.28
	pointsOfFunctionPlot[1_052].Y = 110.174

	pointsOfFunctionPlot[1_053].X = 4.29
	pointsOfFunctionPlot[1_053].Y = 111.391

	pointsOfFunctionPlot[1_054].X = 4.3
	pointsOfFunctionPlot[1_054].Y = 112.621

	pointsOfFunctionPlot[1_055].X = 4.31
	pointsOfFunctionPlot[1_055].Y = 113.865

	pointsOfFunctionPlot[1_056].X = 4.32
	pointsOfFunctionPlot[1_056].Y = 115.123

	pointsOfFunctionPlot[1_057].X = 4.33
	pointsOfFunctionPlot[1_057].Y = 116.395

	pointsOfFunctionPlot[1_058].X = 4.34
	pointsOfFunctionPlot[1_058].Y = 117.681

	pointsOfFunctionPlot[1_059].X = 4.35
	pointsOfFunctionPlot[1_059].Y = 118.981

	pointsOfFunctionPlot[1_060].X = 4.36
	pointsOfFunctionPlot[1_060].Y = 120.295

	pointsOfFunctionPlot[1_061].X = 4.37
	pointsOfFunctionPlot[1_061].Y = 121.624

	pointsOfFunctionPlot[1_062].X = 4.38
	pointsOfFunctionPlot[1_062].Y = 122.967

	pointsOfFunctionPlot[1_063].X = 4.39
	pointsOfFunctionPlot[1_063].Y = 124.326

	pointsOfFunctionPlot[1_064].X = 4.4
	pointsOfFunctionPlot[1_064].Y = 125.699

	pointsOfFunctionPlot[1_065].X = 4.41
	pointsOfFunctionPlot[1_065].Y = 127.088

	pointsOfFunctionPlot[1_066].X = 4.42
	pointsOfFunctionPlot[1_066].Y = 128.492

	pointsOfFunctionPlot[1_067].X = 4.43
	pointsOfFunctionPlot[1_067].Y = 129.911

	pointsOfFunctionPlot[1_068].X = 4.44
	pointsOfFunctionPlot[1_068].Y = 131.346

	pointsOfFunctionPlot[1_069].X = 4.45
	pointsOfFunctionPlot[1_069].Y = 132.797

	pointsOfFunctionPlot[1_070].X = 4.46
	pointsOfFunctionPlot[1_070].Y = 134.264

	pointsOfFunctionPlot[1_071].X = 4.47
	pointsOfFunctionPlot[1_071].Y = 135.747

	pointsOfFunctionPlot[1_072].X = 4.48
	pointsOfFunctionPlot[1_072].Y = 137.247

	pointsOfFunctionPlot[1_073].X = 4.49
	pointsOfFunctionPlot[1_073].Y = 138.763

	pointsOfFunctionPlot[1_074].X = 4.5
	pointsOfFunctionPlot[1_074].Y = 140.296

	pointsOfFunctionPlot[1_075].X = 4.51
	pointsOfFunctionPlot[1_075].Y = 141.845

	pointsOfFunctionPlot[1_076].X = 4.52
	pointsOfFunctionPlot[1_076].Y = 143.412

	pointsOfFunctionPlot[1_077].X = 4.53
	pointsOfFunctionPlot[1_077].Y = 144.997

	pointsOfFunctionPlot[1_078].X = 4.54
	pointsOfFunctionPlot[1_078].Y = 146.598

	pointsOfFunctionPlot[1_079].X = 4.55
	pointsOfFunctionPlot[1_079].Y = 148.218

	pointsOfFunctionPlot[1_080].X = 4.56
	pointsOfFunctionPlot[1_080].Y = 149.855

	pointsOfFunctionPlot[1_081].X = 4.57
	pointsOfFunctionPlot[1_081].Y = 151.511

	pointsOfFunctionPlot[1_082].X = 4.58
	pointsOfFunctionPlot[1_082].Y = 153.184

	pointsOfFunctionPlot[1_083].X = 4.59
	pointsOfFunctionPlot[1_083].Y = 154.876

	pointsOfFunctionPlot[1_084].X = 4.6
	pointsOfFunctionPlot[1_084].Y = 156.587

	pointsOfFunctionPlot[1_085].X = 4.61
	pointsOfFunctionPlot[1_085].Y = 158.317

	pointsOfFunctionPlot[1_086].X = 4.62
	pointsOfFunctionPlot[1_086].Y = 160.066

	pointsOfFunctionPlot[1_087].X = 4.63
	pointsOfFunctionPlot[1_087].Y = 161.834

	pointsOfFunctionPlot[1_088].X = 4.64
	pointsOfFunctionPlot[1_088].Y = 163.622

	pointsOfFunctionPlot[1_089].X = 4.65
	pointsOfFunctionPlot[1_089].Y = 165.429

	pointsOfFunctionPlot[1_090].X = 4.66
	pointsOfFunctionPlot[1_090].Y = 167.257

	pointsOfFunctionPlot[1_091].X = 4.67
	pointsOfFunctionPlot[1_091].Y = 169.104

	pointsOfFunctionPlot[1_092].X = 4.68
	pointsOfFunctionPlot[1_092].Y = 170.973

	pointsOfFunctionPlot[1_093].X = 4.69
	pointsOfFunctionPlot[1_093].Y = 172.861

	pointsOfFunctionPlot[1_094].X = 4.7
	pointsOfFunctionPlot[1_094].Y = 174.771

	pointsOfFunctionPlot[1_095].X = 4.71
	pointsOfFunctionPlot[1_095].Y = 176.701

	pointsOfFunctionPlot[1_096].X = 4.72
	pointsOfFunctionPlot[1_096].Y = 178.653

	pointsOfFunctionPlot[1_097].X = 4.73
	pointsOfFunctionPlot[1_097].Y = 180.627

	pointsOfFunctionPlot[1_098].X = 4.74
	pointsOfFunctionPlot[1_098].Y = 182.622

	pointsOfFunctionPlot[1_099].X = 4.75
	pointsOfFunctionPlot[1_099].Y = 184.64

	pointsOfFunctionPlot[1_100].X = 4.76
	pointsOfFunctionPlot[1_100].Y = 186.679

	pointsOfFunctionPlot[1_101].X = 4.77
	pointsOfFunctionPlot[1_101].Y = 188.741

	pointsOfFunctionPlot[1_102].X = 4.78
	pointsOfFunctionPlot[1_102].Y = 190.826

	pointsOfFunctionPlot[1_103].X = 4.79
	pointsOfFunctionPlot[1_103].Y = 192.934

	pointsOfFunctionPlot[1_104].X = 4.8
	pointsOfFunctionPlot[1_104].Y = 195.066

	pointsOfFunctionPlot[1_105].X = 4.81
	pointsOfFunctionPlot[1_105].Y = 197.221

	pointsOfFunctionPlot[1_106].X = 4.82
	pointsOfFunctionPlot[1_106].Y = 199.399

	pointsOfFunctionPlot[1_107].X = 4.83
	pointsOfFunctionPlot[1_107].Y = 201.602

	pointsOfFunctionPlot[1_108].X = 4.84
	pointsOfFunctionPlot[1_108].Y = 203.829

	pointsOfFunctionPlot[1_109].X = 4.85
	pointsOfFunctionPlot[1_109].Y = 206.081

	pointsOfFunctionPlot[1_110].X = 4.86
	pointsOfFunctionPlot[1_110].Y = 208.357

	pointsOfFunctionPlot[1_111].X = 4.87
	pointsOfFunctionPlot[1_111].Y = 210.659

	pointsOfFunctionPlot[1_112].X = 4.88
	pointsOfFunctionPlot[1_112].Y = 212.986

	pointsOfFunctionPlot[1_113].X = 4.89
	pointsOfFunctionPlot[1_113].Y = 215.339

	pointsOfFunctionPlot[1_114].X = 4.9
	pointsOfFunctionPlot[1_114].Y = 217.717

	pointsOfFunctionPlot[1_115].X = 4.91
	pointsOfFunctionPlot[1_115].Y = 220.123

	pointsOfFunctionPlot[1_116].X = 4.92
	pointsOfFunctionPlot[1_116].Y = 222.554

	pointsOfFunctionPlot[1_117].X = 4.93
	pointsOfFunctionPlot[1_117].Y = 225.013

	pointsOfFunctionPlot[1_118].X = 4.94
	pointsOfFunctionPlot[1_118].Y = 227.498

	pointsOfFunctionPlot[1_119].X = 4.95
	pointsOfFunctionPlot[1_119].Y = 230.011

	pointsOfFunctionPlot[1_120].X = 4.96
	pointsOfFunctionPlot[1_120].Y = 232.552

	pointsOfFunctionPlot[1_121].X = 4.97
	pointsOfFunctionPlot[1_121].Y = 235.121

	pointsOfFunctionPlot[1_122].X = 4.98
	pointsOfFunctionPlot[1_122].Y = 237.719

	pointsOfFunctionPlot[1_123].X = 4.99
	pointsOfFunctionPlot[1_123].Y = 240.345

	pointsOfFunctionPlot[1_124].X = 5.0
	pointsOfFunctionPlot[1_124].Y = 243.0

	pointsOfFunctionPlot[1_125].X = 5.01
	pointsOfFunctionPlot[1_125].Y = 245.684

	pointsOfFunctionPlot[1_126].X = 5.02
	pointsOfFunctionPlot[1_126].Y = 248.398

	pointsOfFunctionPlot[1_127].X = 5.03
	pointsOfFunctionPlot[1_127].Y = 251.142

	pointsOfFunctionPlot[1_128].X = 5.04
	pointsOfFunctionPlot[1_128].Y = 253.916

	pointsOfFunctionPlot[1_129].X = 5.05
	pointsOfFunctionPlot[1_129].Y = 256.916

	pointsOfFunctionPlot[1_130].X = 5.06
	pointsOfFunctionPlot[1_130].Y = 259.557

	pointsOfFunctionPlot[1_131].X = 5.07
	pointsOfFunctionPlot[1_131].Y = 262.424

	pointsOfFunctionPlot[1_132].X = 5.08
	pointsOfFunctionPlot[1_132].Y = 265.323

	pointsOfFunctionPlot[1_133].X = 5.09
	pointsOfFunctionPlot[1_133].Y = 268.254

	pointsOfFunctionPlot[1_134].X = 5.1
	pointsOfFunctionPlot[1_134].Y = 271.217

	pointsOfFunctionPlot[1_135].X = 5.11
	pointsOfFunctionPlot[1_135].Y = 274.214

	pointsOfFunctionPlot[1_136].X = 5.12
	pointsOfFunctionPlot[1_136].Y = 277.243

	pointsOfFunctionPlot[1_137].X = 5.13
	pointsOfFunctionPlot[1_137].Y = 280.305

	pointsOfFunctionPlot[1_138].X = 5.14
	pointsOfFunctionPlot[1_138].Y = 283.402

	pointsOfFunctionPlot[1_139].X = 5.15
	pointsOfFunctionPlot[1_139].Y = 286.532

	pointsOfFunctionPlot[1_140].X = 5.16
	pointsOfFunctionPlot[1_140].Y = 289.698

	pointsOfFunctionPlot[1_141].X = 5.17
	pointsOfFunctionPlot[1_141].Y = 292.898

	pointsOfFunctionPlot[1_142].X = 5.18
	pointsOfFunctionPlot[1_142].Y = 296.133

	pointsOfFunctionPlot[1_143].X = 5.19
	pointsOfFunctionPlot[1_143].Y = 299.405

	pointsOfFunctionPlot[1_144].X = 5.2
	pointsOfFunctionPlot[1_144].Y = 302.712

	pointsOfFunctionPlot[1_145].X = 5.21
	pointsOfFunctionPlot[1_145].Y = 306.056

	pointsOfFunctionPlot[1_146].X = 5.22
	pointsOfFunctionPlot[1_146].Y = 309.437

	pointsOfFunctionPlot[1_147].X = 5.23
	pointsOfFunctionPlot[1_147].Y = 312.855

	pointsOfFunctionPlot[1_148].X = 5.24
	pointsOfFunctionPlot[1_148].Y = 316.311

	pointsOfFunctionPlot[1_149].X = 5.25
	pointsOfFunctionPlot[1_149].Y = 319.806

	pointsOfFunctionPlot[1_150].X = 5.26
	pointsOfFunctionPlot[1_150].Y = 323.338

	pointsOfFunctionPlot[1_151].X = 5.27
	pointsOfFunctionPlot[1_151].Y = 326.91

	pointsOfFunctionPlot[1_152].X = 5.28
	pointsOfFunctionPlot[1_152].Y = 330.521

	pointsOfFunctionPlot[1_153].X = 5.29
	pointsOfFunctionPlot[1_153].Y = 334.173

	pointsOfFunctionPlot[1_154].X = 5.3
	pointsOfFunctionPlot[1_154].Y = 337.864

	pointsOfFunctionPlot[1_155].X = 5.31
	pointsOfFunctionPlot[1_155].Y = 341.596

	pointsOfFunctionPlot[1_156].X = 5.32
	pointsOfFunctionPlot[1_156].Y = 345.37

	pointsOfFunctionPlot[1_157].X = 5.33
	pointsOfFunctionPlot[1_157].Y = 349.185

	pointsOfFunctionPlot[1_158].X = 5.34
	pointsOfFunctionPlot[1_158].Y = 353.042

	pointsOfFunctionPlot[1_159].X = 5.35
	pointsOfFunctionPlot[1_159].Y = 356.942

	pointsOfFunctionPlot[1_160].X = 5.36
	pointsOfFunctionPlot[1_160].Y = 360.885

	pointsOfFunctionPlot[1_161].X = 5.37
	pointsOfFunctionPlot[1_161].Y = 364.872

	pointsOfFunctionPlot[1_162].X = 5.38
	pointsOfFunctionPlot[1_162].Y = 368.903

	pointsOfFunctionPlot[1_163].X = 5.39
	pointsOfFunctionPlot[1_163].Y = 372.978

	pointsOfFunctionPlot[1_164].X = 5.4
	pointsOfFunctionPlot[1_164].Y = 377.098

	pointsOfFunctionPlot[1_165].X = 5.41
	pointsOfFunctionPlot[1_165].Y = 381.264

	pointsOfFunctionPlot[1_166].X = 5.42
	pointsOfFunctionPlot[1_166].Y = 385.475

	pointsOfFunctionPlot[1_167].X = 5.43
	pointsOfFunctionPlot[1_167].Y = 389.734

	pointsOfFunctionPlot[1_168].X = 5.44
	pointsOfFunctionPlot[1_168].Y = 394.039

	pointsOfFunctionPlot[1_169].X = 5.45
	pointsOfFunctionPlot[1_169].Y = 398.392

	pointsOfFunctionPlot[1_170].X = 5.46
	pointsOfFunctionPlot[1_170].Y = 402.793

	pointsOfFunctionPlot[1_171].X = 5.47
	pointsOfFunctionPlot[1_171].Y = 407.242

	pointsOfFunctionPlot[1_172].X = 5.48
	pointsOfFunctionPlot[1_172].Y = 411.741

	pointsOfFunctionPlot[1_173].X = 5.49
	pointsOfFunctionPlot[1_173].Y = 416.289

	pointsOfFunctionPlot[1_174].X = 5.5
	pointsOfFunctionPlot[1_174].Y = 420.888

	pointsOfFunctionPlot[1_175].X = 5.51
	pointsOfFunctionPlot[1_175].Y = 425.537

	pointsOfFunctionPlot[1_176].X = 5.52
	pointsOfFunctionPlot[1_176].Y = 430.238

	pointsOfFunctionPlot[1_177].X = 5.53
	pointsOfFunctionPlot[1_177].Y = 434.991

	pointsOfFunctionPlot[1_178].X = 5.54
	pointsOfFunctionPlot[1_178].Y = 439.796

	pointsOfFunctionPlot[1_179].X = 5.55
	pointsOfFunctionPlot[1_179].Y = 444.654

	pointsOfFunctionPlot[1_180].X = 5.56
	pointsOfFunctionPlot[1_180].Y = 449.566

	pointsOfFunctionPlot[1_181].X = 5.57
	pointsOfFunctionPlot[1_181].Y = 454.533

	pointsOfFunctionPlot[1_182].X = 5.58
	pointsOfFunctionPlot[1_182].Y = 459.554

	pointsOfFunctionPlot[1_183].X = 5.59
	pointsOfFunctionPlot[1_183].Y = 464.63

	pointsOfFunctionPlot[1_184].X = 5.6
	pointsOfFunctionPlot[1_184].Y = 469.763

	pointsOfFunctionPlot[1_185].X = 5.61
	pointsOfFunctionPlot[1_185].Y = 474.952

	pointsOfFunctionPlot[1_186].X = 5.62
	pointsOfFunctionPlot[1_186].Y = 480.199

	pointsOfFunctionPlot[1_187].X = 5.63
	pointsOfFunctionPlot[1_187].Y = 485.503

	pointsOfFunctionPlot[1_188].X = 5.64
	pointsOfFunctionPlot[1_188].Y = 490.867

	pointsOfFunctionPlot[1_189].X = 5.65
	pointsOfFunctionPlot[1_189].Y = 496.289

	pointsOfFunctionPlot[1_190].X = 5.66
	pointsOfFunctionPlot[1_190].Y = 501.771

	pointsOfFunctionPlot[1_191].X = 5.67
	pointsOfFunctionPlot[1_191].Y = 507.314

	pointsOfFunctionPlot[1_192].X = 5.68
	pointsOfFunctionPlot[1_192].Y = 512.918

	pointsOfFunctionPlot[1_193].X = 5.69
	pointsOfFunctionPlot[1_193].Y = 518.585

	pointsOfFunctionPlot[1_194].X = 5.7
	pointsOfFunctionPlot[1_194].Y = 524.313

	pointsOfFunctionPlot[1_195].X = 5.71
	pointsOfFunctionPlot[1_195].Y = 530.105

	pointsOfFunctionPlot[1_196].X = 5.72
	pointsOfFunctionPlot[1_196].Y = 535.961

	pointsOfFunctionPlot[1_197].X = 5.73
	pointsOfFunctionPlot[1_197].Y = 541.882

	pointsOfFunctionPlot[1_198].X = 5.74
	pointsOfFunctionPlot[1_198].Y = 547.868

	pointsOfFunctionPlot[1_199].X = 5.75
	pointsOfFunctionPlot[1_199].Y = 553.92

	pointsOfFunctionPlot[1_200].X = 5.76
	pointsOfFunctionPlot[1_200].Y = 560.039

	pointsOfFunctionPlot[1_201].X = 5.77
	pointsOfFunctionPlot[1_201].Y = 566.225

	pointsOfFunctionPlot[1_202].X = 5.78
	pointsOfFunctionPlot[1_202].Y = 572.48

	pointsOfFunctionPlot[1_203].X = 5.79
	pointsOfFunctionPlot[1_203].Y = 578.804

	pointsOfFunctionPlot[1_204].X = 5.8
	pointsOfFunctionPlot[1_204].Y = 585.198

	pointsOfFunctionPlot[1_205].X = 5.81
	pointsOfFunctionPlot[1_205].Y = 591.663

	pointsOfFunctionPlot[1_206].X = 5.82
	pointsOfFunctionPlot[1_206].Y = 598.199

	pointsOfFunctionPlot[1_207].X = 5.83
	pointsOfFunctionPlot[1_207].Y = 604.807

	pointsOfFunctionPlot[1_208].X = 5.84
	pointsOfFunctionPlot[1_208].Y = 611.488

	pointsOfFunctionPlot[1_209].X = 5.85
	pointsOfFunctionPlot[1_209].Y = 618.243

	pointsOfFunctionPlot[1_210].X = 5.86
	pointsOfFunctionPlot[1_210].Y = 625.072

	pointsOfFunctionPlot[1_211].X = 5.87
	pointsOfFunctionPlot[1_211].Y = 631.977

	pointsOfFunctionPlot[1_212].X = 5.88
	pointsOfFunctionPlot[1_212].Y = 638.959

	pointsOfFunctionPlot[1_213].X = 5.89
	pointsOfFunctionPlot[1_213].Y = 646.017

	pointsOfFunctionPlot[1_214].X = 5.9
	pointsOfFunctionPlot[1_214].Y = 653.153

	pointsOfFunctionPlot[1_215].X = 5.91
	pointsOfFunctionPlot[1_215].Y = 660.368

	pointsOfFunctionPlot[1_216].X = 5.92
	pointsOfFunctionPlot[1_216].Y = 667.663

	pointsOfFunctionPlot[1_217].X = 5.93
	pointsOfFunctionPlot[1_217].Y = 675.039

	pointsOfFunctionPlot[1_218].X = 5.94
	pointsOfFunctionPlot[1_218].Y = 682.496

	pointsOfFunctionPlot[1_219].X = 5.95
	pointsOfFunctionPlot[1_219].Y = 690.035

	pointsOfFunctionPlot[1_220].X = 5.96
	pointsOfFunctionPlot[1_220].Y = 697.658

	pointsOfFunctionPlot[1_221].X = 5.97
	pointsOfFunctionPlot[1_221].Y = 705.365

	pointsOfFunctionPlot[1_222].X = 5.98
	pointsOfFunctionPlot[1_222].Y = 713.156

	pointsOfFunctionPlot[1_223].X = 5.99
	pointsOfFunctionPlot[1_223].Y = 721.034

	pointsOfFunctionPlot[1_224].X = 6.0
	pointsOfFunctionPlot[1_224].Y = 729.0

	pointsOfFunctionPlot[1_225].X = 6.01
	pointsOfFunctionPlot[1_225].Y = 737.053

	pointsOfFunctionPlot[1_226].X = 6.02
	pointsOfFunctionPlot[1_226].Y = 745.195

	pointsOfFunctionPlot[1_227].X = 6.03
	pointsOfFunctionPlot[1_227].Y = 753.427

	pointsOfFunctionPlot[1_228].X = 6.04
	pointsOfFunctionPlot[1_228].Y = 761.749

	pointsOfFunctionPlot[1_229].X = 6.05
	pointsOfFunctionPlot[1_229].Y = 770.164

	pointsOfFunctionPlot[1_230].X = 6.06
	pointsOfFunctionPlot[1_230].Y = 778.672

	pointsOfFunctionPlot[1_231].X = 6.07
	pointsOfFunctionPlot[1_231].Y = 787.274

	pointsOfFunctionPlot[1_232].X = 6.08
	pointsOfFunctionPlot[1_232].Y = 795.971

	pointsOfFunctionPlot[1_233].X = 6.09
	pointsOfFunctionPlot[1_233].Y = 804.763

	pointsOfFunctionPlot[1_234].X = 6.1
	pointsOfFunctionPlot[1_234].Y = 813.653

	pointsOfFunctionPlot[1_235].X = 6.11
	pointsOfFunctionPlot[1_235].Y = 822.642

	pointsOfFunctionPlot[1_236].X = 6.12
	pointsOfFunctionPlot[1_236].Y = 831.729

	pointsOfFunctionPlot[1_237].X = 6.13
	pointsOfFunctionPlot[1_237].Y = 840.917

	pointsOfFunctionPlot[1_238].X = 6.14
	pointsOfFunctionPlot[1_238].Y = 850.206

	pointsOfFunctionPlot[1_239].X = 6.15
	pointsOfFunctionPlot[1_239].Y = 859.598

	pointsOfFunctionPlot[1_240].X = 6.16
	pointsOfFunctionPlot[1_240].Y = 869.094

	pointsOfFunctionPlot[1_241].X = 6.17
	pointsOfFunctionPlot[1_241].Y = 878.695

	pointsOfFunctionPlot[1_242].X = 6.18
	pointsOfFunctionPlot[1_242].Y = 888.401

	pointsOfFunctionPlot[1_243].X = 6.19
	pointsOfFunctionPlot[1_243].Y = 898.215

	pointsOfFunctionPlot[1_244].X = 6.2
	pointsOfFunctionPlot[1_244].Y = 909.137

	pointsOfFunctionPlot[1_245].X = 6.21
	pointsOfFunctionPlot[1_245].Y = 918.169

	pointsOfFunctionPlot[1_246].X = 6.22
	pointsOfFunctionPlot[1_246].Y = 928.312

	pointsOfFunctionPlot[1_247].X = 6.23
	pointsOfFunctionPlot[1_247].Y = 938.567

	pointsOfFunctionPlot[1_248].X = 6.24
	pointsOfFunctionPlot[1_248].Y = 948.935

	pointsOfFunctionPlot[1_249].X = 6.25
	pointsOfFunctionPlot[1_249].Y = 959.418

	pointsOfFunctionPlot[1_250].X = 6.26
	pointsOfFunctionPlot[1_250].Y = 970.016

	pointsOfFunctionPlot[1_251].X = 6.27
	pointsOfFunctionPlot[1_251].Y = 980.731

	pointsOfFunctionPlot[1_252].X = 6.28
	pointsOfFunctionPlot[1_252].Y = 991.565

	pointsOfFunctionPlot[1_253].X = 6.29
	pointsOfFunctionPlot[1_253].Y = 1_002.519

	pointsOfFunctionPlot[1_254].X = 6.3
	pointsOfFunctionPlot[1_254].Y = 1_013.593

	pointsOfFunctionPlot[1_255].X = 6.31
	pointsOfFunctionPlot[1_255].Y = 1_024.79

	pointsOfFunctionPlot[1_256].X = 6.32
	pointsOfFunctionPlot[1_256].Y = 1_036.111

	pointsOfFunctionPlot[1_257].X = 6.33
	pointsOfFunctionPlot[1_257].Y = 1_047.556

	pointsOfFunctionPlot[1_258].X = 6.34
	pointsOfFunctionPlot[1_258].Y = 1_059.128

	pointsOfFunctionPlot[1_259].X = 6.35
	pointsOfFunctionPlot[1_259].Y = 1_070.828

	pointsOfFunctionPlot[1_260].X = 6.36
	pointsOfFunctionPlot[1_260].Y = 1_082.657

	pointsOfFunctionPlot[1_261].X = 6.37
	pointsOfFunctionPlot[1_261].Y = 1_094.617

	pointsOfFunctionPlot[1_262].X = 6.38
	pointsOfFunctionPlot[1_262].Y = 1_106.709

	pointsOfFunctionPlot[1_263].X = 6.39
	pointsOfFunctionPlot[1_263].Y = 1_118.934

	pointsOfFunctionPlot[1_264].X = 6.4
	pointsOfFunctionPlot[1_264].Y = 1_131.295

	pointsOfFunctionPlot[1_265].X = 6.41
	pointsOfFunctionPlot[1_265].Y = 1_143.792

	pointsOfFunctionPlot[1_266].X = 6.42
	pointsOfFunctionPlot[1_266].Y = 1_156.427

	pointsOfFunctionPlot[1_267].X = 6.43
	pointsOfFunctionPlot[1_267].Y = 1_169.202

	pointsOfFunctionPlot[1_268].X = 6.44
	pointsOfFunctionPlot[1_268].Y = 1_182.118

	pointsOfFunctionPlot[1_269].X = 6.45
	pointsOfFunctionPlot[1_269].Y = 1_195.176

	pointsOfFunctionPlot[1_270].X = 6.46
	pointsOfFunctionPlot[1_270].Y = 1_208.379

	pointsOfFunctionPlot[1_271].X = 6.47
	pointsOfFunctionPlot[1_271].Y = 1_221.728

	pointsOfFunctionPlot[1_272].X = 6.48
	pointsOfFunctionPlot[1_272].Y = 1_235.224

	pointsOfFunctionPlot[1_273].X = 6.49
	pointsOfFunctionPlot[1_273].Y = 1_248.869

	pointsOfFunctionPlot[1_274].X = 6.5
	pointsOfFunctionPlot[1_274].Y = 1_262.665

	pointsOfFunctionPlot[1_275].X = 6.51
	pointsOfFunctionPlot[1_275].Y = 1_276.613

	pointsOfFunctionPlot[1_276].X = 6.52
	pointsOfFunctionPlot[1_276].Y = 1_290.715

	pointsOfFunctionPlot[1_277].X = 6.53
	pointsOfFunctionPlot[1_277].Y = 1_304.973

	pointsOfFunctionPlot[1_278].X = 6.54
	pointsOfFunctionPlot[1_278].Y = 1_319.389

	pointsOfFunctionPlot[1_279].X = 6.55
	pointsOfFunctionPlot[1_279].Y = 1_333.964

	pointsOfFunctionPlot[1_280].X = 6.56
	pointsOfFunctionPlot[1_280].Y = 1_348.7

	pointsOfFunctionPlot[1_281].X = 6.57
	pointsOfFunctionPlot[1_281].Y = 1_363.598

	pointsOfFunctionPlot[1_282].X = 6.58
	pointsOfFunctionPlot[1_282].Y = 1_378.662

	pointsOfFunctionPlot[1_283].X = 6.59
	pointsOfFunctionPlot[1_283].Y = 1_393.891

	pointsOfFunctionPlot[1_284].X = 6.6
	pointsOfFunctionPlot[1_284].Y = 1_409.289

	pointsOfFunctionPlot[1_285].X = 6.61
	pointsOfFunctionPlot[1_285].Y = 1_424.857

	pointsOfFunctionPlot[1_286].X = 6.62
	pointsOfFunctionPlot[1_286].Y = 1_440.597

	pointsOfFunctionPlot[1_287].X = 6.63
	pointsOfFunctionPlot[1_287].Y = 1_456.511

	pointsOfFunctionPlot[1_288].X = 6.64
	pointsOfFunctionPlot[1_288].Y = 1_472.601

	pointsOfFunctionPlot[1_289].X = 6.65
	pointsOfFunctionPlot[1_289].Y = 1_488.868

	pointsOfFunctionPlot[1_290].X = 6.66
	pointsOfFunctionPlot[1_290].Y = 1_505.315

	pointsOfFunctionPlot[1_291].X = 6.67
	pointsOfFunctionPlot[1_291].Y = 1_521.944

	pointsOfFunctionPlot[1_292].X = 6.68
	pointsOfFunctionPlot[1_292].Y = 1_538.756

	pointsOfFunctionPlot[1_293].X = 6.69
	pointsOfFunctionPlot[1_293].Y = 1_555.755

	pointsOfFunctionPlot[1_294].X = 6.7
	pointsOfFunctionPlot[1_294].Y = 1_572.94

	pointsOfFunctionPlot[1_295].X = 6.71
	pointsOfFunctionPlot[1_295].Y = 1_590.316

	pointsOfFunctionPlot[1_296].X = 6.72
	pointsOfFunctionPlot[1_296].Y = 1_607.884

	pointsOfFunctionPlot[1_297].X = 6.73
	pointsOfFunctionPlot[1_297].Y = 1_625.646

	pointsOfFunctionPlot[1_298].X = 6.74
	pointsOfFunctionPlot[1_298].Y = 1_643.604

	pointsOfFunctionPlot[1_299].X = 6.75
	pointsOfFunctionPlot[1_299].Y = 1_661.76

	pointsOfFunctionPlot[1_300].X = 6.76
	pointsOfFunctionPlot[1_300].Y = 1_680.117

	pointsOfFunctionPlot[1_301].X = 6.77
	pointsOfFunctionPlot[1_301].Y = 1_698.677

	pointsOfFunctionPlot[1_302].X = 6.78
	pointsOfFunctionPlot[1_302].Y = 1_717.442

	pointsOfFunctionPlot[1_303].X = 6.79
	pointsOfFunctionPlot[1_303].Y = 1_736.414

	pointsOfFunctionPlot[1_304].X = 6.8
	pointsOfFunctionPlot[1_304].Y = 1_755.595

	pointsOfFunctionPlot[1_305].X = 6.81
	pointsOfFunctionPlot[1_305].Y = 1_774.989

	pointsOfFunctionPlot[1_306].X = 6.82
	pointsOfFunctionPlot[1_306].Y = 1_794.597

	pointsOfFunctionPlot[1_307].X = 6.83
	pointsOfFunctionPlot[1_307].Y = 1_814.421

	pointsOfFunctionPlot[1_308].X = 6.84
	pointsOfFunctionPlot[1_308].Y = 1_834.464

	pointsOfFunctionPlot[1_309].X = 6.85
	pointsOfFunctionPlot[1_309].Y = 1_854.729

	pointsOfFunctionPlot[1_310].X = 6.86
	pointsOfFunctionPlot[1_310].Y = 1_875.218

	pointsOfFunctionPlot[1_311].X = 6.87
	pointsOfFunctionPlot[1_311].Y = 1_895.933

	pointsOfFunctionPlot[1_312].X = 6.88
	pointsOfFunctionPlot[1_312].Y = 1_916.876

	pointsOfFunctionPlot[1_313].X = 6.89
	pointsOfFunctionPlot[1_313].Y = 1_938.052

	pointsOfFunctionPlot[1_314].X = 6.9
	pointsOfFunctionPlot[1_314].Y = 1_959.461

	pointsOfFunctionPlot[1_315].X = 6.91
	pointsOfFunctionPlot[1_315].Y = 1_981.106

	pointsOfFunctionPlot[1_316].X = 6.92
	pointsOfFunctionPlot[1_316].Y = 2_002.991

	pointsOfFunctionPlot[1_317].X = 6.93
	pointsOfFunctionPlot[1_317].Y = 2_025.117

	pointsOfFunctionPlot[1_318].X = 6.94
	pointsOfFunctionPlot[1_318].Y = 2_047.488

	pointsOfFunctionPlot[1_319].X = 6.95
	pointsOfFunctionPlot[1_319].Y = 2_070.106

	pointsOfFunctionPlot[1_320].X = 6.96
	pointsOfFunctionPlot[1_320].Y = 2_092.974

	pointsOfFunctionPlot[1_321].X = 6.97
	pointsOfFunctionPlot[1_321].Y = 2_116.094

	pointsOfFunctionPlot[1_322].X = 6.98
	pointsOfFunctionPlot[1_322].Y = 2_139.47

	pointsOfFunctionPlot[1_323].X = 6.99
	pointsOfFunctionPlot[1_323].Y = 2_163.104

	pointsOfFunctionPlot[1_324].X = 7.0
	pointsOfFunctionPlot[1_324].Y = 2_187.0

	pointsOfFunctionPlot[1_325].X = 7.01
	pointsOfFunctionPlot[1_325].Y = 2_211.159

	pointsOfFunctionPlot[1_326].X = 7.02
	pointsOfFunctionPlot[1_326].Y = 2_235.585

	pointsOfFunctionPlot[1_327].X = 7.03
	pointsOfFunctionPlot[1_327].Y = 2_260.28

	pointsOfFunctionPlot[1_328].X = 7.04
	pointsOfFunctionPlot[1_328].Y = 2_285.249

	pointsOfFunctionPlot[1_329].X = 7.05
	pointsOfFunctionPlot[1_329].Y = 2_310.494

	pointsOfFunctionPlot[1_330].X = 7.06
	pointsOfFunctionPlot[1_330].Y = 2_336.017

	pointsOfFunctionPlot[1_331].X = 7.07
	pointsOfFunctionPlot[1_331].Y = 2_361.822

	pointsOfFunctionPlot[1_332].X = 7.08
	pointsOfFunctionPlot[1_332].Y = 2_387.912

	pointsOfFunctionPlot[1_333].X = 7.09
	pointsOfFunctionPlot[1_333].Y = 2_414.291

	pointsOfFunctionPlot[1_334].X = 7.1
	pointsOfFunctionPlot[1_334].Y = 2_440.961

	pointsOfFunctionPlot[1_335].X = 7.11
	pointsOfFunctionPlot[1_335].Y = 2_467.925

	pointsOfFunctionPlot[1_336].X = 7.12
	pointsOfFunctionPlot[1_336].Y = 2_495.188

	pointsOfFunctionPlot[1_337].X = 7.13
	pointsOfFunctionPlot[1_337].Y = 2_522.751

	pointsOfFunctionPlot[1_338].X = 7.14
	pointsOfFunctionPlot[1_338].Y = 2_550.62

	pointsOfFunctionPlot[1_339].X = 7.15
	pointsOfFunctionPlot[1_339].Y = 2_578.795

	pointsOfFunctionPlot[1_340].X = 7.16
	pointsOfFunctionPlot[1_340].Y = 2_607.283

	pointsOfFunctionPlot[1_341].X = 7.17
	pointsOfFunctionPlot[1_341].Y = 2_636.084

	pointsOfFunctionPlot[1_342].X = 7.18
	pointsOfFunctionPlot[1_342].Y = 2_665.204

	pointsOfFunctionPlot[1_343].X = 7.19
	pointsOfFunctionPlot[1_343].Y = 2_694.646

	pointsOfFunctionPlot[1_344].X = 7.2
	pointsOfFunctionPlot[1_344].Y = 2_724.413

	pointsOfFunctionPlot[1_345].X = 7.21
	pointsOfFunctionPlot[1_345].Y = 2_754.509

	pointsOfFunctionPlot[1_346].X = 7.22
	pointsOfFunctionPlot[1_346].Y = 2_784.937

	pointsOfFunctionPlot[1_347].X = 7.23
	pointsOfFunctionPlot[1_347].Y = 2_815.701

	pointsOfFunctionPlot[1_348].X = 7.24
	pointsOfFunctionPlot[1_348].Y = 2_846.806

	pointsOfFunctionPlot[1_349].X = 7.25
	pointsOfFunctionPlot[1_349].Y = 2_878.253

	pointsOfFunctionPlot[1_350].X = 7.26
	pointsOfFunctionPlot[1_350].Y = 2_910.049

	pointsOfFunctionPlot[1_351].X = 7.27
	pointsOfFunctionPlot[1_351].Y = 2_942.195

	pointsOfFunctionPlot[1_352].X = 7.28
	pointsOfFunctionPlot[1_352].Y = 2_974.697

	pointsOfFunctionPlot[1_353].X = 7.29
	pointsOfFunctionPlot[1_353].Y = 3_007.557

	pointsOfFunctionPlot[1_354].X = 7.3
	pointsOfFunctionPlot[1_354].Y = 3_040.781

	pointsOfFunctionPlot[1_355].X = 7.31
	pointsOfFunctionPlot[1_355].Y = 3_074.371

	pointsOfFunctionPlot[1_356].X = 7.32
	pointsOfFunctionPlot[1_356].Y = 3_108.333

	pointsOfFunctionPlot[1_357].X = 7.33
	pointsOfFunctionPlot[1_357].Y = 3_142.67

	pointsOfFunctionPlot[1_358].X = 7.34
	pointsOfFunctionPlot[1_358].Y = 3_177.386

	pointsOfFunctionPlot[1_359].X = 7.35
	pointsOfFunctionPlot[1_359].Y = 3_212.485

	pointsOfFunctionPlot[1_360].X = 7.36
	pointsOfFunctionPlot[1_360].Y = 3_247.973

	pointsOfFunctionPlot[1_361].X = 7.37
	pointsOfFunctionPlot[1_361].Y = 3_283.852

	pointsOfFunctionPlot[1_362].X = 7.38
	pointsOfFunctionPlot[1_362].Y = 3_320.128

	pointsOfFunctionPlot[1_363].X = 7.39
	pointsOfFunctionPlot[1_363].Y = 3_356.804

	pointsOfFunctionPlot[1_364].X = 7.4
	pointsOfFunctionPlot[1_364].Y = 3_393.886

	pointsOfFunctionPlot[1_365].X = 7.41
	pointsOfFunctionPlot[1_365].Y = 3_431.377

	pointsOfFunctionPlot[1_366].X = 7.42
	pointsOfFunctionPlot[1_366].Y = 3_469.282

	pointsOfFunctionPlot[1_367].X = 7.43
	pointsOfFunctionPlot[1_367].Y = 3_507.607

	pointsOfFunctionPlot[1_368].X = 7.44
	pointsOfFunctionPlot[1_368].Y = 3_546.354

	pointsOfFunctionPlot[1_369].X = 7.45
	pointsOfFunctionPlot[1_369].Y = 3_585.529

	pointsOfFunctionPlot[1_370].X = 7.46
	pointsOfFunctionPlot[1_370].Y = 3_625.138

	pointsOfFunctionPlot[1_371].X = 7.47
	pointsOfFunctionPlot[1_371].Y = 3_665.183

	pointsOfFunctionPlot[1_372].X = 7.48
	pointsOfFunctionPlot[1_372].Y = 3_705.672

	pointsOfFunctionPlot[1_373].X = 7.49
	pointsOfFunctionPlot[1_373].Y = 3_746.607

	pointsOfFunctionPlot[1_374].X = 7.5
	pointsOfFunctionPlot[1_374].Y = 3_787.995

	pointsOfFunctionPlot[1_375].X = 7.51
	pointsOfFunctionPlot[1_375].Y = 3_829.839

	pointsOfFunctionPlot[1_376].X = 7.52
	pointsOfFunctionPlot[1_376].Y = 3_872.147

	pointsOfFunctionPlot[1_377].X = 7.53
	pointsOfFunctionPlot[1_377].Y = 3_914.921

	pointsOfFunctionPlot[1_378].X = 7.54
	pointsOfFunctionPlot[1_378].Y = 3_958.168

	pointsOfFunctionPlot[1_379].X = 7.55
	pointsOfFunctionPlot[1_379].Y = 4_001.893

	pointsOfFunctionPlot[1_380].X = 7.56
	pointsOfFunctionPlot[1_380].Y = 4_046.1

	pointsOfFunctionPlot[1_381].X = 7.57
	pointsOfFunctionPlot[1_381].Y = 4_090.796

	pointsOfFunctionPlot[1_382].X = 7.58
	pointsOfFunctionPlot[1_382].Y = 4_135.986

	pointsOfFunctionPlot[1_383].X = 7.59
	pointsOfFunctionPlot[1_383].Y = 4_181.675

	pointsOfFunctionPlot[1_384].X = 7.6
	pointsOfFunctionPlot[1_384].Y = 4_227.869

	pointsOfFunctionPlot[1_385].X = 7.61
	pointsOfFunctionPlot[1_385].Y = 4_274.573

	pointsOfFunctionPlot[1_386].X = 7.62
	pointsOfFunctionPlot[1_386].Y = 4_321.793

	pointsOfFunctionPlot[1_387].X = 7.63
	pointsOfFunctionPlot[1_387].Y = 4_369.534

	pointsOfFunctionPlot[1_388].X = 7.64
	pointsOfFunctionPlot[1_388].Y = 4_417.803

	pointsOfFunctionPlot[1_389].X = 7.65
	pointsOfFunctionPlot[1_389].Y = 4_466.605

	pointsOfFunctionPlot[1_390].X = 7.66
	pointsOfFunctionPlot[1_390].Y = 4_515.946

	pointsOfFunctionPlot[1_391].X = 7.67
	pointsOfFunctionPlot[1_391].Y = 4_565.833

	pointsOfFunctionPlot[1_392].X = 7.68
	pointsOfFunctionPlot[1_392].Y = 4_616.27

	pointsOfFunctionPlot[1_393].X = 7.69
	pointsOfFunctionPlot[1_393].Y = 4_667.264

	pointsOfFunctionPlot[1_394].X = 7.7
	pointsOfFunctionPlot[1_394].Y = 4_718.822

	pointsOfFunctionPlot[1_395].X = 7.71
	pointsOfFunctionPlot[1_395].Y = 4_770.95

	pointsOfFunctionPlot[1_396].X = 7.72
	pointsOfFunctionPlot[1_396].Y = 4_823.653

	pointsOfFunctionPlot[1_397].X = 7.73
	pointsOfFunctionPlot[1_397].Y = 4_876.938

	pointsOfFunctionPlot[1_398].X = 7.74
	pointsOfFunctionPlot[1_398].Y = 4_930.812

	pointsOfFunctionPlot[1_399].X = 7.75
	pointsOfFunctionPlot[1_399].Y = 4_985.281

	pointsOfFunctionPlot[1_400].X = 7.76
	pointsOfFunctionPlot[1_400].Y = 5_040.352

	pointsOfFunctionPlot[1_401].X = 7.77
	pointsOfFunctionPlot[1_401].Y = 5_096.032

	pointsOfFunctionPlot[1_402].X = 7.78
	pointsOfFunctionPlot[1_402].Y = 5_152.326

	pointsOfFunctionPlot[1_403].X = 7.79
	pointsOfFunctionPlot[1_403].Y = 5_209.242

	pointsOfFunctionPlot[1_404].X = 7.8
	pointsOfFunctionPlot[1_404].Y = 5_266.787

	pointsOfFunctionPlot[1_405].X = 7.81
	pointsOfFunctionPlot[1_405].Y = 5_324.968

	pointsOfFunctionPlot[1_406].X = 7.82
	pointsOfFunctionPlot[1_406].Y = 5_383.791

	pointsOfFunctionPlot[1_407].X = 7.83
	pointsOfFunctionPlot[1_407].Y = 5_443.264

	pointsOfFunctionPlot[1_408].X = 7.84
	pointsOfFunctionPlot[1_408].Y = 5_503.394

	pointsOfFunctionPlot[1_409].X = 7.85
	pointsOfFunctionPlot[1_409].Y = 5_564.188

	pointsOfFunctionPlot[1_410].X = 7.86
	pointsOfFunctionPlot[1_410].Y = 5_625.654

	pointsOfFunctionPlot[1_411].X = 7.87
	pointsOfFunctionPlot[1_411].Y = 5_687.799

	pointsOfFunctionPlot[1_412].X = 7.88
	pointsOfFunctionPlot[1_412].Y = 5_750.63

	pointsOfFunctionPlot[1_413].X = 7.89
	pointsOfFunctionPlot[1_413].Y = 5_814.156

	pointsOfFunctionPlot[1_414].X = 7.9
	pointsOfFunctionPlot[1_414].Y = 5_878.383

	pointsOfFunctionPlot[1_415].X = 7.91
	pointsOfFunctionPlot[1_415].Y = 5_943.32

	pointsOfFunctionPlot[1_416].X = 7.92
	pointsOfFunctionPlot[1_416].Y = 6_008.974

	pointsOfFunctionPlot[1_417].X = 7.93
	pointsOfFunctionPlot[1_417].Y = 6_075.353

	pointsOfFunctionPlot[1_418].X = 7.94
	pointsOfFunctionPlot[1_418].Y = 6_142.466

	pointsOfFunctionPlot[1_419].X = 7.95
	pointsOfFunctionPlot[1_419].Y = 6_210.319

	pointsOfFunctionPlot[1_420].X = 7.96
	pointsOfFunctionPlot[1_420].Y = 6_278.923

	pointsOfFunctionPlot[1_421].X = 7.97
	pointsOfFunctionPlot[1_421].Y = 6_348.284

	pointsOfFunctionPlot[1_422].X = 7.98
	pointsOfFunctionPlot[1_422].Y = 6_418.412

	pointsOfFunctionPlot[1_423].X = 7.99
	pointsOfFunctionPlot[1_423].Y = 6_489.314

	pointsOfFunctionPlot[1_424].X = 8.0
	pointsOfFunctionPlot[1_424].Y = 6_561.0

	pointsOfFunctionPlot[1_425].X = 8.01
	pointsOfFunctionPlot[1_425].Y = 6_633.477

	pointsOfFunctionPlot[1_426].X = 8.02
	pointsOfFunctionPlot[1_426].Y = 6_706.755

	pointsOfFunctionPlot[1_427].X = 8.03
	pointsOfFunctionPlot[1_427].Y = 6_780.842

	pointsOfFunctionPlot[1_428].X = 8.04
	pointsOfFunctionPlot[1_428].Y = 6_855.748

	pointsOfFunctionPlot[1_429].X = 8.05
	pointsOfFunctionPlot[1_429].Y = 6_931.482

	pointsOfFunctionPlot[1_430].X = 8.06
	pointsOfFunctionPlot[1_430].Y = 7_008.052

	pointsOfFunctionPlot[1_431].X = 8.07
	pointsOfFunctionPlot[1_431].Y = 7_085.467

	pointsOfFunctionPlot[1_432].X = 8.08
	pointsOfFunctionPlot[1_432].Y = 7_163.738

	pointsOfFunctionPlot[1_433].X = 8.09
	pointsOfFunctionPlot[1_433].Y = 7_242.874

	pointsOfFunctionPlot[1_434].X = 8.1
	pointsOfFunctionPlot[1_434].Y = 7_322.884

	pointsOfFunctionPlot[1_435].X = 8.11
	pointsOfFunctionPlot[1_435].Y = 7_403.777

	pointsOfFunctionPlot[1_436].X = 8.12
	pointsOfFunctionPlot[1_436].Y = 7_485.565

	pointsOfFunctionPlot[1_437].X = 8.13
	pointsOfFunctionPlot[1_437].Y = 7_568.255

	pointsOfFunctionPlot[1_438].X = 8.14
	pointsOfFunctionPlot[1_438].Y = 7_651.86

	pointsOfFunctionPlot[1_439].X = 8.15
	pointsOfFunctionPlot[1_439].Y = 7_736.387

	pointsOfFunctionPlot[1_440].X = 8.16
	pointsOfFunctionPlot[1_440].Y = 7_821.849

	pointsOfFunctionPlot[1_441].X = 8.17
	pointsOfFunctionPlot[1_441].Y = 7_908.254

	pointsOfFunctionPlot[1_442].X = 8.18
	pointsOfFunctionPlot[1_442].Y = 7_995.614

	pointsOfFunctionPlot[1_443].X = 8.19
	pointsOfFunctionPlot[1_443].Y = 8_083.939

	pointsOfFunctionPlot[1_444].X = 8.2
	pointsOfFunctionPlot[1_444].Y = 8_173.24

	pointsOfFunctionPlot[1_445].X = 8.21
	pointsOfFunctionPlot[1_445].Y = 8_263.528

	pointsOfFunctionPlot[1_446].X = 8.22
	pointsOfFunctionPlot[1_446].Y = 8_354.812

	pointsOfFunctionPlot[1_447].X = 8.23
	pointsOfFunctionPlot[1_447].Y = 8_447.105

	pointsOfFunctionPlot[1_448].X = 8.24
	pointsOfFunctionPlot[1_448].Y = 8_540.418

	pointsOfFunctionPlot[1_449].X = 8.25
	pointsOfFunctionPlot[1_449].Y = 8_634.761

	pointsOfFunctionPlot[1_450].X = 8.26
	pointsOfFunctionPlot[1_450].Y = 8_730.147

	pointsOfFunctionPlot[1_451].X = 8.27
	pointsOfFunctionPlot[1_451].Y = 8_826.586

	pointsOfFunctionPlot[1_452].X = 8.28
	pointsOfFunctionPlot[1_452].Y = 8_924.091

	pointsOfFunctionPlot[1_453].X = 8.29
	pointsOfFunctionPlot[1_453].Y = 9_022.672

	pointsOfFunctionPlot[1_454].X = 8.3
	pointsOfFunctionPlot[1_454].Y = 9_122.343

	pointsOfFunctionPlot[1_455].X = 8.31
	pointsOfFunctionPlot[1_455].Y = 9_223.115

	pointsOfFunctionPlot[1_456].X = 8.32
	pointsOfFunctionPlot[1_456].Y = 9_325.0

	pointsOfFunctionPlot[1_457].X = 8.33
	pointsOfFunctionPlot[1_457].Y = 9_428.01

	pointsOfFunctionPlot[1_458].X = 8.34
	pointsOfFunctionPlot[1_458].Y = 9_532.158

	pointsOfFunctionPlot[1_459].X = 8.35
	pointsOfFunctionPlot[1_459].Y = 9_637.457

	pointsOfFunctionPlot[1_460].X = 8.36
	pointsOfFunctionPlot[1_460].Y = 9_743.919

	pointsOfFunctionPlot[1_461].X = 8.37
	pointsOfFunctionPlot[1_461].Y = 9_851.557

	pointsOfFunctionPlot[1_462].X = 8.38
	pointsOfFunctionPlot[1_462].Y = 9_960.384

	pointsOfFunctionPlot[1_463].X = 8.39
	pointsOfFunctionPlot[1_463].Y = 10_070.414

	pointsOfFunctionPlot[1_464].X = 8.4
	pointsOfFunctionPlot[1_464].Y = 10_181.658

	pointsOfFunctionPlot[1_465].X = 8.41
	pointsOfFunctionPlot[1_465].Y = 10_294.132

	pointsOfFunctionPlot[1_466].X = 8.42
	pointsOfFunctionPlot[1_466].Y = 10_407.846

	pointsOfFunctionPlot[1_467].X = 8.43
	pointsOfFunctionPlot[1_467].Y = 10_522.82

	pointsOfFunctionPlot[1_468].X = 8.44
	pointsOfFunctionPlot[1_468].Y = 10_639.063

	pointsOfFunctionPlot[1_469].X = 8.45
	pointsOfFunctionPlot[1_469].Y = 10_756.589

	pointsOfFunctionPlot[1_470].X = 8.46
	pointsOfFunctionPlot[1_470].Y = 10_875.414

	pointsOfFunctionPlot[1_471].X = 8.47
	pointsOfFunctionPlot[1_471].Y = 10_995.551

	pointsOfFunctionPlot[1_472].X = 8.48
	pointsOfFunctionPlot[1_472].Y = 11_117.016

	pointsOfFunctionPlot[1_473].X = 8.49
	pointsOfFunctionPlot[1_473].Y = 11_239.822

	pointsOfFunctionPlot[1_474].X = 8.5
	pointsOfFunctionPlot[1_474].Y = 11_363.985

	pointsOfFunctionPlot[1_475].X = 8.51
	pointsOfFunctionPlot[1_475].Y = 11_489.519

	pointsOfFunctionPlot[1_476].X = 8.52
	pointsOfFunctionPlot[1_476].Y = 11_616.441

	pointsOfFunctionPlot[1_477].X = 8.53
	pointsOfFunctionPlot[1_477].Y = 11_744.764

	pointsOfFunctionPlot[1_478].X = 8.54
	pointsOfFunctionPlot[1_478].Y = 11_874.505

	pointsOfFunctionPlot[1_479].X = 8.55
	pointsOfFunctionPlot[1_479].Y = 12_005.679

	pointsOfFunctionPlot[1_480].X = 8.56
	pointsOfFunctionPlot[1_480].Y = 12_138.302

	pointsOfFunctionPlot[1_481].X = 8.57
	pointsOfFunctionPlot[1_481].Y = 12_272.39

	pointsOfFunctionPlot[1_482].X = 8.58
	pointsOfFunctionPlot[1_482].Y = 12_407.959

	pointsOfFunctionPlot[1_483].X = 8.59
	pointsOfFunctionPlot[1_483].Y = 12_545.026

	pointsOfFunctionPlot[1_484].X = 8.6
	pointsOfFunctionPlot[1_484].Y = 12_683.607

	pointsOfFunctionPlot[1_485].X = 8.61
	pointsOfFunctionPlot[1_485].Y = 12_823.719

	pointsOfFunctionPlot[1_486].X = 8.62
	pointsOfFunctionPlot[1_486].Y = 12_965.379

	pointsOfFunctionPlot[1_487].X = 8.63
	pointsOfFunctionPlot[1_487].Y = 13_108.603

	pointsOfFunctionPlot[1_488].X = 8.64
	pointsOfFunctionPlot[1_488].Y = 13_253.41

	pointsOfFunctionPlot[1_489].X = 8.65
	pointsOfFunctionPlot[1_489].Y = 13_399.816

	pointsOfFunctionPlot[1_490].X = 8.66
	pointsOfFunctionPlot[1_490].Y = 13_547.84

	pointsOfFunctionPlot[1_491].X = 8.67
	pointsOfFunctionPlot[1_491].Y = 13_697.499

	pointsOfFunctionPlot[1_492].X = 8.68
	pointsOfFunctionPlot[1_492].Y = 13_848.811

	pointsOfFunctionPlot[1_493].X = 8.69
	pointsOfFunctionPlot[1_493].Y = 14_001.794

	pointsOfFunctionPlot[1_494].X = 8.7
	pointsOfFunctionPlot[1_494].Y = 14_156.468

	pointsOfFunctionPlot[1_495].X = 8.71
	pointsOfFunctionPlot[1_495].Y = 14_156.468

	pointsOfFunctionPlot[1_496].X = 8.72
	pointsOfFunctionPlot[1_496].Y = 14_470.959

	pointsOfFunctionPlot[1_497].X = 8.73
	pointsOfFunctionPlot[1_497].Y = 14_630.816

	pointsOfFunctionPlot[1_498].X = 8.74
	pointsOfFunctionPlot[1_498].Y = 14_792.438

	pointsOfFunctionPlot[1_499].X = 8.75
	pointsOfFunctionPlot[1_499].Y = 14_955.845

	pointsOfFunctionPlot[1_500].X = 8.76
	pointsOfFunctionPlot[1_500].Y = 15_121.058

	pointsOfFunctionPlot[1_501].X = 8.77
	pointsOfFunctionPlot[1_501].Y = 15_288.096

	pointsOfFunctionPlot[1_502].X = 8.78
	pointsOfFunctionPlot[1_502].Y = 15_456.979

	pointsOfFunctionPlot[1_503].X = 8.79
	pointsOfFunctionPlot[1_503].Y = 15_627.727

	pointsOfFunctionPlot[1_504].X = 8.8
	pointsOfFunctionPlot[1_504].Y = 15_800.362

	pointsOfFunctionPlot[1_505].X = 8.81
	pointsOfFunctionPlot[1_505].Y = 15_974.903

	pointsOfFunctionPlot[1_506].X = 8.82
	pointsOfFunctionPlot[1_506].Y = 16_151.373

	pointsOfFunctionPlot[1_507].X = 8.83
	pointsOfFunctionPlot[1_507].Y = 16_329.793

	pointsOfFunctionPlot[1_508].X = 8.84
	pointsOfFunctionPlot[1_508].Y = 16_510.183

	pointsOfFunctionPlot[1_509].X = 8.85
	pointsOfFunctionPlot[1_509].Y = 16_692.566

	pointsOfFunctionPlot[1_510].X = 8.86
	pointsOfFunctionPlot[1_510].Y = 16_876.963

	pointsOfFunctionPlot[1_511].X = 8.87
	pointsOfFunctionPlot[1_511].Y = 17_063.398

	pointsOfFunctionPlot[1_512].X = 8.88
	pointsOfFunctionPlot[1_512].Y = 17_251.892

	pointsOfFunctionPlot[1_513].X = 8.89
	pointsOfFunctionPlot[1_513].Y = 17_442.468

	pointsOfFunctionPlot[1_514].X = 8.9
	pointsOfFunctionPlot[1_514].Y = 17_635.15

	pointsOfFunctionPlot[1_515].X = 8.91
	pointsOfFunctionPlot[1_515].Y = 17_829.96

	pointsOfFunctionPlot[1_516].X = 8.92
	pointsOfFunctionPlot[1_516].Y = 18_026.922

	pointsOfFunctionPlot[1_517].X = 8.93
	pointsOfFunctionPlot[1_517].Y = 18_226.06

	pointsOfFunctionPlot[1_518].X = 8.94
	pointsOfFunctionPlot[1_518].Y = 18_427.398

	pointsOfFunctionPlot[1_519].X = 8.95
	pointsOfFunctionPlot[1_519].Y = 18_630.959

	pointsOfFunctionPlot[1_520].X = 8.96
	pointsOfFunctionPlot[1_520].Y = 18_836.77

	pointsOfFunctionPlot[1_521].X = 8.97
	pointsOfFunctionPlot[1_521].Y = 19_044.854

	pointsOfFunctionPlot[1_522].X = 8.98
	pointsOfFunctionPlot[1_522].Y = 19_255.237

	pointsOfFunctionPlot[1_523].X = 8.99
	pointsOfFunctionPlot[1_523].Y = 19_467.943

	pointsOfFunctionPlot[1_524].X = 9.0
	pointsOfFunctionPlot[1_524].Y = 19_683.0

	pointsOfFunctionPlot[1_525].X = 9.01
	pointsOfFunctionPlot[1_525].Y = 19_900.432

	pointsOfFunctionPlot[1_526].X = 9.02
	pointsOfFunctionPlot[1_526].Y = 20_120.266

	pointsOfFunctionPlot[1_527].X = 9.03
	pointsOfFunctionPlot[1_527].Y = 20_342.528

	pointsOfFunctionPlot[1_528].X = 9.04
	pointsOfFunctionPlot[1_528].Y = 20_567.246

	pointsOfFunctionPlot[1_529].X = 9.05
	pointsOfFunctionPlot[1_529].Y = 20_794.446

	pointsOfFunctionPlot[1_530].X = 9.06
	pointsOfFunctionPlot[1_530].Y = 21_024.155

	pointsOfFunctionPlot[1_531].X = 9.07
	pointsOfFunctionPlot[1_531].Y = 21_256.403

	pointsOfFunctionPlot[1_532].X = 9.08
	pointsOfFunctionPlot[1_532].Y = 21_491.216

	pointsOfFunctionPlot[1_533].X = 9.09
	pointsOfFunctionPlot[1_533].Y = 21_728.623

	pointsOfFunctionPlot[1_534].X = 9.1
	pointsOfFunctionPlot[1_534].Y = 21_968.652

	pointsOfFunctionPlot[1_535].X = 9.11
	pointsOfFunctionPlot[1_535].Y = 22_211.333

	pointsOfFunctionPlot[1_536].X = 9.12
	pointsOfFunctionPlot[1_536].Y = 22_456.695

	pointsOfFunctionPlot[1_537].X = 9.13
	pointsOfFunctionPlot[1_537].Y = 22_704.767

	pointsOfFunctionPlot[1_538].X = 9.14
	pointsOfFunctionPlot[1_538].Y = 22_955.579

	pointsOfFunctionPlot[1_539].X = 9.15
	pointsOfFunctionPlot[1_539].Y = 23_209.163

	pointsOfFunctionPlot[1_540].X = 9.16
	pointsOfFunctionPlot[1_540].Y = 23_465.547

	pointsOfFunctionPlot[1_541].X = 9.17
	pointsOfFunctionPlot[1_541].Y = 23_724.764

	pointsOfFunctionPlot[1_542].X = 9.18
	pointsOfFunctionPlot[1_542].Y = 23_986.844

	pointsOfFunctionPlot[1_543].X = 9.19
	pointsOfFunctionPlot[1_543].Y = 24_251.819

	pointsOfFunctionPlot[1_544].X = 9.2
	pointsOfFunctionPlot[1_544].Y = 24_519.722

	pointsOfFunctionPlot[1_545].X = 9.21
	pointsOfFunctionPlot[1_545].Y = 24_790.583

	pointsOfFunctionPlot[1_546].X = 9.22
	pointsOfFunctionPlot[1_546].Y = 25_064.437

	pointsOfFunctionPlot[1_547].X = 9.23
	pointsOfFunctionPlot[1_547].Y = 25_341.317

	pointsOfFunctionPlot[1_548].X = 9.24
	pointsOfFunctionPlot[1_548].Y = 25_621.254

	pointsOfFunctionPlot[1_549].X = 9.25
	pointsOfFunctionPlot[1_549].Y = 25_904.284

	pointsOfFunctionPlot[1_550].X = 9.26
	pointsOfFunctionPlot[1_550].Y = 26_190.441

	pointsOfFunctionPlot[1_551].X = 9.27
	pointsOfFunctionPlot[1_551].Y = 26_479.759

	pointsOfFunctionPlot[1_552].X = 9.28
	pointsOfFunctionPlot[1_552].Y = 26_772.272

	pointsOfFunctionPlot[1_553].X = 9.29
	pointsOfFunctionPlot[1_553].Y = 27_068.018

	pointsOfFunctionPlot[1_554].X = 9.3
	pointsOfFunctionPlot[1_554].Y = 27_367.03

	pointsOfFunctionPlot[1_555].X = 9.31
	pointsOfFunctionPlot[1_555].Y = 27_367.03

	pointsOfFunctionPlot[1_556].X = 9.32
	pointsOfFunctionPlot[1_556].Y = 27_669.345

	pointsOfFunctionPlot[1_557].X = 9.33
	pointsOfFunctionPlot[1_557].Y = 28_284.031

	pointsOfFunctionPlot[1_558].X = 9.34
	pointsOfFunctionPlot[1_558].Y = 28_596.476

	pointsOfFunctionPlot[1_559].X = 9.35
	pointsOfFunctionPlot[1_559].Y = 28_912.372

	pointsOfFunctionPlot[1_560].X = 9.36
	pointsOfFunctionPlot[1_560].Y = 29_231.758

	pointsOfFunctionPlot[1_561].X = 9.37
	pointsOfFunctionPlot[1_561].Y = 29_554.672

	pointsOfFunctionPlot[1_562].X = 9.38
	pointsOfFunctionPlot[1_562].Y = 29_881.154

	pointsOfFunctionPlot[1_563].X = 9.39
	pointsOfFunctionPlot[1_563].Y = 30_211.242

	pointsOfFunctionPlot[1_564].X = 9.4
	pointsOfFunctionPlot[1_564].Y = 30_544.976

	pointsOfFunctionPlot[1_565].X = 9.41
	pointsOfFunctionPlot[1_565].Y = 30_882.397

	pointsOfFunctionPlot[1_566].X = 9.42
	pointsOfFunctionPlot[1_566].Y = 31_223.545

	pointsOfFunctionPlot[1_567].X = 9.43
	pointsOfFunctionPlot[1_567].Y = 31_568.462

	pointsOfFunctionPlot[1_568].X = 9.44
	pointsOfFunctionPlot[1_568].Y = 31_917.189

	pointsOfFunctionPlot[1_569].X = 9.45
	pointsOfFunctionPlot[1_569].Y = 32_269.769

	pointsOfFunctionPlot[1_570].X = 9.46
	pointsOfFunctionPlot[1_570].Y = 32_262.243

	pointsOfFunctionPlot[1_571].X = 9.47
	pointsOfFunctionPlot[1_571].Y = 32_986.655

	pointsOfFunctionPlot[1_572].X = 9.48
	pointsOfFunctionPlot[1_572].Y = 33_351.048

	pointsOfFunctionPlot[1_573].X = 9.49
	pointsOfFunctionPlot[1_573].Y = 33_719.467

	pointsOfFunctionPlot[1_574].X = 9.5
	pointsOfFunctionPlot[1_574].Y = 34_091.956

	pointsOfFunctionPlot[1_575].X = 9.51
	pointsOfFunctionPlot[1_575].Y = 34_468.559

	pointsOfFunctionPlot[1_576].X = 9.52
	pointsOfFunctionPlot[1_576].Y = 34_849.322

	pointsOfFunctionPlot[1_577].X = 9.53
	pointsOfFunctionPlot[1_577].Y = 35_234.292

	pointsOfFunctionPlot[1_578].X = 9.54
	pointsOfFunctionPlot[1_578].Y = 35_623.515

	pointsOfFunctionPlot[1_579].X = 9.55
	pointsOfFunctionPlot[1_579].Y = 36_017.037

	pointsOfFunctionPlot[1_580].X = 9.56
	pointsOfFunctionPlot[1_580].Y = 36_414.906

	pointsOfFunctionPlot[1_581].X = 9.57
	pointsOfFunctionPlot[1_581].Y = 36_817.17

	pointsOfFunctionPlot[1_582].X = 9.58
	pointsOfFunctionPlot[1_582].Y = 37_223.878

	pointsOfFunctionPlot[1_583].X = 9.59
	pointsOfFunctionPlot[1_583].Y = 37_635.079

	pointsOfFunctionPlot[1_584].X = 9.6
	pointsOfFunctionPlot[1_584].Y = 38_050.822

	pointsOfFunctionPlot[1_585].X = 9.61
	pointsOfFunctionPlot[1_585].Y = 38_471.157

	pointsOfFunctionPlot[1_586].X = 9.62
	pointsOfFunctionPlot[1_586].Y = 38_896.136

	pointsOfFunctionPlot[1_587].X = 9.63
	pointsOfFunctionPlot[1_587].Y = 39_325.81

	pointsOfFunctionPlot[1_588].X = 9.64
	pointsOfFunctionPlot[1_588].Y = 39_760.23

	pointsOfFunctionPlot[1_589].X = 9.65
	pointsOfFunctionPlot[1_589].Y = 40_199.449

	pointsOfFunctionPlot[1_590].X = 9.66
	pointsOfFunctionPlot[1_590].Y = 40_643.52

	pointsOfFunctionPlot[1_591].X = 9.67
	pointsOfFunctionPlot[1_591].Y = 41_092.497

	pointsOfFunctionPlot[1_592].X = 9.68
	pointsOfFunctionPlot[1_592].Y = 41_546.433

	pointsOfFunctionPlot[1_593].X = 9.69
	pointsOfFunctionPlot[1_593].Y = 42_005.383

	pointsOfFunctionPlot[1_594].X = 9.7
	pointsOfFunctionPlot[1_594].Y = 42_469.404

	pointsOfFunctionPlot[1_595].X = 9.71
	pointsOfFunctionPlot[1_595].Y = 42_938.55

	pointsOfFunctionPlot[1_596].X = 9.72
	pointsOfFunctionPlot[1_596].Y = 43_892.448

	pointsOfFunctionPlot[1_597].X = 9.73
	pointsOfFunctionPlot[1_597].Y = 43_892.448

	pointsOfFunctionPlot[1_598].X = 9.74
	pointsOfFunctionPlot[1_598].Y = 44_377.314

	pointsOfFunctionPlot[1_599].X = 9.75
	pointsOfFunctionPlot[1_599].Y = 44_867.537

	pointsOfFunctionPlot[1_600].X = 9.76
	pointsOfFunctionPlot[1_600].Y = 45_363.175

	pointsOfFunctionPlot[1_601].X = 9.77
	pointsOfFunctionPlot[1_601].Y = 45_864.288

	pointsOfFunctionPlot[1_602].X = 9.78
	pointsOfFunctionPlot[1_602].Y = 46_370.937

	pointsOfFunctionPlot[1_603].X = 9.79
	pointsOfFunctionPlot[1_603].Y = 46_883.182

	pointsOfFunctionPlot[1_604].X = 9.8
	pointsOfFunctionPlot[1_604].Y = 47_401.086

	pointsOfFunctionPlot[1_605].X = 9.81
	pointsOfFunctionPlot[1_605].Y = 47_924.711

	pointsOfFunctionPlot[1_606].X = 9.82
	pointsOfFunctionPlot[1_606].Y = 48_454.121

	pointsOfFunctionPlot[1_607].X = 9.83
	pointsOfFunctionPlot[1_607].Y = 48_989.379

	pointsOfFunctionPlot[1_608].X = 9.84
	pointsOfFunctionPlot[1_608].Y = 49_530.549

	pointsOfFunctionPlot[1_609].X = 9.85
	pointsOfFunctionPlot[1_609].Y = 50_077.698

	pointsOfFunctionPlot[1_610].X = 9.86
	pointsOfFunctionPlot[1_610].Y = 50_630.891

	pointsOfFunctionPlot[1_611].X = 9.87
	pointsOfFunctionPlot[1_611].Y = 51_190.195

	pointsOfFunctionPlot[1_612].X = 9.88
	pointsOfFunctionPlot[1_612].Y = 51_755.677

	pointsOfFunctionPlot[1_613].X = 9.89
	pointsOfFunctionPlot[1_613].Y = 52_327.406

	pointsOfFunctionPlot[1_614].X = 9.9
	pointsOfFunctionPlot[1_614].Y = 52_905.451

	pointsOfFunctionPlot[1_615].X = 9.91
	pointsOfFunctionPlot[1_615].Y = 53_489.881

	pointsOfFunctionPlot[1_616].X = 9.92
	pointsOfFunctionPlot[1_616].Y = 54_080.767

	pointsOfFunctionPlot[1_617].X = 9.93
	pointsOfFunctionPlot[1_617].Y = 54_678.181

	pointsOfFunctionPlot[1_618].X = 9.94
	pointsOfFunctionPlot[1_618].Y = 55_282.194

	pointsOfFunctionPlot[1_619].X = 9.95
	pointsOfFunctionPlot[1_619].Y = 55_892.879

	pointsOfFunctionPlot[1_620].X = 9.96
	pointsOfFunctionPlot[1_620].Y = 56_510.31

	pointsOfFunctionPlot[1_621].X = 9.97
	pointsOfFunctionPlot[1_621].Y = 57_134.562

	pointsOfFunctionPlot[1_622].X = 9.98
	pointsOfFunctionPlot[1_622].Y = 57_765.71

	pointsOfFunctionPlot[1_623].X = 9.99
	pointsOfFunctionPlot[1_623].Y = 58_403.83

	pointsOfFunctionPlot[1_624].X = 10.0
	pointsOfFunctionPlot[1_624].Y = 59_049.0










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function 3^x"

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
		"Power-of-3-plot-01.png"); err != nil {

		panic(err)
	}
}
