package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function exp(x)

	pointsOfFunctionPlot := make(plotter.XYs, 1_698)

	pointsOfFunctionPlot[0].X = -10.0
	pointsOfFunctionPlot[0].Y = 0.0

	pointsOfFunctionPlot[1].X = -6.97
	pointsOfFunctionPlot[1].Y = 0.0

	pointsOfFunctionPlot[2].X = -6.96
	pointsOfFunctionPlot[2].Y = 0.001

	pointsOfFunctionPlot[3].X = -6.95
	pointsOfFunctionPlot[3].Y = 0.001

	pointsOfFunctionPlot[4].X = -6.94
	pointsOfFunctionPlot[4].Y = 0.001

	pointsOfFunctionPlot[5].X = -6.93
	pointsOfFunctionPlot[5].Y = 0.001

	pointsOfFunctionPlot[5].X = -6.92
	pointsOfFunctionPlot[5].Y = 0.001

	pointsOfFunctionPlot[6].X = -6.91
	pointsOfFunctionPlot[6].Y = 0.001

	pointsOfFunctionPlot[7].X = -6.9
	pointsOfFunctionPlot[7].Y = 0.001

	pointsOfFunctionPlot[8].X = -6.89
	pointsOfFunctionPlot[8].Y = 0.001

	pointsOfFunctionPlot[9].X = -6.88
	pointsOfFunctionPlot[9].Y = 0.001

	pointsOfFunctionPlot[10].X = -6.87
	pointsOfFunctionPlot[10].Y = 0.001

	pointsOfFunctionPlot[11].X = -6.86
	pointsOfFunctionPlot[11].Y = 0.001

	pointsOfFunctionPlot[12].X = -6.85
	pointsOfFunctionPlot[12].Y = 0.001

	pointsOfFunctionPlot[13].X = -6.84
	pointsOfFunctionPlot[13].Y = 0.001

	pointsOfFunctionPlot[14].X = -6.83
	pointsOfFunctionPlot[14].Y = 0.001

	pointsOfFunctionPlot[15].X = -6.82
	pointsOfFunctionPlot[15].Y = 0.001

	pointsOfFunctionPlot[16].X = -6.81
	pointsOfFunctionPlot[16].Y = 0.001

	pointsOfFunctionPlot[17].X = -6.8
	pointsOfFunctionPlot[17].Y = 0.001

	pointsOfFunctionPlot[18].X = -6.79
	pointsOfFunctionPlot[18].Y = 0.001

	pointsOfFunctionPlot[19].X = -6.78
	pointsOfFunctionPlot[19].Y = 0.001

	pointsOfFunctionPlot[20].X = -6.77
	pointsOfFunctionPlot[20].Y = 0.001

	pointsOfFunctionPlot[21].X = -6.76
	pointsOfFunctionPlot[21].Y = 0.001

	pointsOfFunctionPlot[22].X = -6.75
	pointsOfFunctionPlot[22].Y = 0.001

	pointsOfFunctionPlot[23].X = -6.74
	pointsOfFunctionPlot[23].Y = 0.001

	pointsOfFunctionPlot[24].X = -6.73
	pointsOfFunctionPlot[24].Y = 0.001

	pointsOfFunctionPlot[25].X = -6.72
	pointsOfFunctionPlot[25].Y = 0.001

	pointsOfFunctionPlot[26].X = -6.71
	pointsOfFunctionPlot[26].Y = 0.001

	pointsOfFunctionPlot[27].X = -6.7
	pointsOfFunctionPlot[27].Y = 0.001

	pointsOfFunctionPlot[28].X = -6.69
	pointsOfFunctionPlot[28].Y = 0.001

	pointsOfFunctionPlot[29].X = -6.68
	pointsOfFunctionPlot[29].Y = 0.001

	pointsOfFunctionPlot[30].X = -6.67
	pointsOfFunctionPlot[30].Y = 0.001

	pointsOfFunctionPlot[31].X = -6.66
	pointsOfFunctionPlot[31].Y = 0.001

	pointsOfFunctionPlot[32].X = -6.65
	pointsOfFunctionPlot[32].Y = 0.001

	pointsOfFunctionPlot[33].X = -6.64
	pointsOfFunctionPlot[33].Y = 0.001

	pointsOfFunctionPlot[34].X = -6.63
	pointsOfFunctionPlot[34].Y = 0.001

	pointsOfFunctionPlot[35].X = -6.62
	pointsOfFunctionPlot[35].Y = 0.001

	pointsOfFunctionPlot[36].X = -6.61
	pointsOfFunctionPlot[36].Y = 0.001

	pointsOfFunctionPlot[37].X = -6.6
	pointsOfFunctionPlot[37].Y = 0.001

	pointsOfFunctionPlot[38].X = -6.59
	pointsOfFunctionPlot[38].Y = 0.001

	pointsOfFunctionPlot[39].X = -6.58
	pointsOfFunctionPlot[39].Y = 0.001

	pointsOfFunctionPlot[40].X = -6.57
	pointsOfFunctionPlot[40].Y = 0.001

	pointsOfFunctionPlot[41].X = -6.56
	pointsOfFunctionPlot[41].Y = 0.001

	pointsOfFunctionPlot[42].X = -6.55
	pointsOfFunctionPlot[42].Y = 0.001

	pointsOfFunctionPlot[43].X = -6.54
	pointsOfFunctionPlot[43].Y = 0.001

	pointsOfFunctionPlot[44].X = -6.53
	pointsOfFunctionPlot[44].Y = 0.001

	pointsOfFunctionPlot[45].X = -6.52
	pointsOfFunctionPlot[45].Y = 0.001

	pointsOfFunctionPlot[46].X = -6.51
	pointsOfFunctionPlot[46].Y = 0.001

	pointsOfFunctionPlot[47].X = -6.5
	pointsOfFunctionPlot[47].Y = 0.001

	pointsOfFunctionPlot[48].X = -6.49
	pointsOfFunctionPlot[48].Y = 0.001

	pointsOfFunctionPlot[49].X = -6.48
	pointsOfFunctionPlot[49].Y = 0.001

	pointsOfFunctionPlot[50].X = -6.47
	pointsOfFunctionPlot[50].Y = 0.001

	pointsOfFunctionPlot[51].X = -6.46
	pointsOfFunctionPlot[51].Y = 0.001

	pointsOfFunctionPlot[52].X = -6.45
	pointsOfFunctionPlot[52].Y = 0.001

	pointsOfFunctionPlot[53].X = -6.44
	pointsOfFunctionPlot[53].Y = 0.001

	pointsOfFunctionPlot[54].X = -6.43
	pointsOfFunctionPlot[54].Y = 0.001

	pointsOfFunctionPlot[55].X = -6.42
	pointsOfFunctionPlot[55].Y = 0.001

	pointsOfFunctionPlot[56].X = -6.41
	pointsOfFunctionPlot[56].Y = 0.001

	pointsOfFunctionPlot[57].X = -6.4
	pointsOfFunctionPlot[57].Y = 0.001

	pointsOfFunctionPlot[58].X = -6.39
	pointsOfFunctionPlot[58].Y = 0.001

	pointsOfFunctionPlot[59].X = -6.38
	pointsOfFunctionPlot[59].Y = 0.001

	pointsOfFunctionPlot[60].X = -6.37
	pointsOfFunctionPlot[60].Y = 0.001

	pointsOfFunctionPlot[61].X = -6.36
	pointsOfFunctionPlot[61].Y = 0.001

	pointsOfFunctionPlot[62].X = -6.35
	pointsOfFunctionPlot[62].Y = 0.001

	pointsOfFunctionPlot[63].X = -6.34
	pointsOfFunctionPlot[63].Y = 0.001

	pointsOfFunctionPlot[64].X = -6.33
	pointsOfFunctionPlot[64].Y = 0.001

	pointsOfFunctionPlot[65].X = -6.32
	pointsOfFunctionPlot[65].Y = 0.001

	pointsOfFunctionPlot[66].X = -6.31
	pointsOfFunctionPlot[66].Y = 0.001

	pointsOfFunctionPlot[67].X = -6.3
	pointsOfFunctionPlot[67].Y = 0.001

	pointsOfFunctionPlot[68].X = -6.29
	pointsOfFunctionPlot[68].Y = 0.001

	pointsOfFunctionPlot[69].X = -6.28
	pointsOfFunctionPlot[69].Y = 0.001

	pointsOfFunctionPlot[70].X = -6.27
	pointsOfFunctionPlot[70].Y = 0.001

	pointsOfFunctionPlot[71].X = -6.26
	pointsOfFunctionPlot[71].Y = 0.001

	pointsOfFunctionPlot[72].X = -6.25
	pointsOfFunctionPlot[72].Y = 0.001

	pointsOfFunctionPlot[73].X = -6.24
	pointsOfFunctionPlot[73].Y = 0.001

	pointsOfFunctionPlot[74].X = -6.23
	pointsOfFunctionPlot[74].Y = 0.002

	pointsOfFunctionPlot[75].X = -6.22
	pointsOfFunctionPlot[75].Y = 0.002

	pointsOfFunctionPlot[76].X = -6.21
	pointsOfFunctionPlot[76].Y = 0.002

	pointsOfFunctionPlot[77].X = -6.2
	pointsOfFunctionPlot[77].Y = 0.002

	pointsOfFunctionPlot[78].X = -6.19
	pointsOfFunctionPlot[78].Y = 0.002

	pointsOfFunctionPlot[79].X = -6.18
	pointsOfFunctionPlot[79].Y = 0.002

	pointsOfFunctionPlot[80].X = -6.17
	pointsOfFunctionPlot[80].Y = 0.002

	pointsOfFunctionPlot[81].X = -6.16
	pointsOfFunctionPlot[81].Y = 0.002

	pointsOfFunctionPlot[82].X = -6.15
	pointsOfFunctionPlot[82].Y = 0.002

	pointsOfFunctionPlot[83].X = -6.14
	pointsOfFunctionPlot[83].Y = 0.002

	pointsOfFunctionPlot[84].X = -6.13
	pointsOfFunctionPlot[84].Y = 0.002

	pointsOfFunctionPlot[85].X = -6.12
	pointsOfFunctionPlot[85].Y = 0.002

	pointsOfFunctionPlot[86].X = -6.11
	pointsOfFunctionPlot[86].Y = 0.002

	pointsOfFunctionPlot[87].X = -6.1
	pointsOfFunctionPlot[87].Y = 0.002

	pointsOfFunctionPlot[88].X = -6.09
	pointsOfFunctionPlot[88].Y = 0.002

	pointsOfFunctionPlot[89].X = -6.08
	pointsOfFunctionPlot[89].Y = 0.002

	pointsOfFunctionPlot[90].X = -6.07
	pointsOfFunctionPlot[90].Y = 0.002

	pointsOfFunctionPlot[91].X = -6.06
	pointsOfFunctionPlot[91].Y = 0.002

	pointsOfFunctionPlot[92].X = -6.05
	pointsOfFunctionPlot[92].Y = 0.002

	pointsOfFunctionPlot[93].X = -6.04
	pointsOfFunctionPlot[93].Y = 0.002

	pointsOfFunctionPlot[94].X = -6.03
	pointsOfFunctionPlot[94].Y = 0.002

	pointsOfFunctionPlot[95].X = -6.02
	pointsOfFunctionPlot[95].Y = 0.002

	pointsOfFunctionPlot[96].X = -6.01
	pointsOfFunctionPlot[96].Y = 0.002

	pointsOfFunctionPlot[97].X = -6.0
	pointsOfFunctionPlot[97].Y = 0.002

	pointsOfFunctionPlot[98].X = -5.99
	pointsOfFunctionPlot[98].Y = 0.002

	pointsOfFunctionPlot[99].X = -5.98
	pointsOfFunctionPlot[99].Y = 0.002

	pointsOfFunctionPlot[100].X = -5.97
	pointsOfFunctionPlot[100].Y = 0.002

	pointsOfFunctionPlot[101].X = -5.96
	pointsOfFunctionPlot[101].Y = 0.002

	pointsOfFunctionPlot[102].X = -5.95
	pointsOfFunctionPlot[102].Y = 0.002

	pointsOfFunctionPlot[103].X = -5.94
	pointsOfFunctionPlot[103].Y = 0.002

	pointsOfFunctionPlot[104].X = -5.93
	pointsOfFunctionPlot[104].Y = 0.002

	pointsOfFunctionPlot[105].X = -5.92
	pointsOfFunctionPlot[105].Y = 0.002

	pointsOfFunctionPlot[106].X = -5.91
	pointsOfFunctionPlot[106].Y = 0.002

	pointsOfFunctionPlot[107].X = -5.9
	pointsOfFunctionPlot[107].Y = 0.002

	pointsOfFunctionPlot[108].X = -5.89
	pointsOfFunctionPlot[108].Y = 0.002

	pointsOfFunctionPlot[109].X = -5.88
	pointsOfFunctionPlot[109].Y = 0.002

	pointsOfFunctionPlot[110].X = -5.87
	pointsOfFunctionPlot[110].Y = 0.002

	pointsOfFunctionPlot[111].X = -5.86
	pointsOfFunctionPlot[111].Y = 0.002

	pointsOfFunctionPlot[112].X = -5.85
	pointsOfFunctionPlot[112].Y = 0.002

	pointsOfFunctionPlot[113].X = -5.84
	pointsOfFunctionPlot[113].Y = 0.002

	pointsOfFunctionPlot[114].X = -5.82
	pointsOfFunctionPlot[114].Y = 0.002

	pointsOfFunctionPlot[115].X = -5.82
	pointsOfFunctionPlot[115].Y = 0.003

	pointsOfFunctionPlot[116].X = -5.81
	pointsOfFunctionPlot[116].Y = 0.003

	pointsOfFunctionPlot[117].X = -5.8
	pointsOfFunctionPlot[117].Y = 0.003

	pointsOfFunctionPlot[118].X = -5.79
	pointsOfFunctionPlot[118].Y = 0.003

	pointsOfFunctionPlot[119].X = -5.78
	pointsOfFunctionPlot[119].Y = 0.003

	pointsOfFunctionPlot[120].X = -5.77
	pointsOfFunctionPlot[120].Y = 0.003

	pointsOfFunctionPlot[121].X = -5.76
	pointsOfFunctionPlot[121].Y = 0.003

	pointsOfFunctionPlot[122].X = -5.75
	pointsOfFunctionPlot[122].Y = 0.003

	pointsOfFunctionPlot[123].X = -5.74
	pointsOfFunctionPlot[123].Y = 0.003

	pointsOfFunctionPlot[124].X = -5.73
	pointsOfFunctionPlot[124].Y = 0.003

	pointsOfFunctionPlot[125].X = -5.72
	pointsOfFunctionPlot[125].Y = 0.003

	pointsOfFunctionPlot[126].X = -5.71
	pointsOfFunctionPlot[126].Y = 0.003

	pointsOfFunctionPlot[127].X = -5.7
	pointsOfFunctionPlot[127].Y = 0.003

	pointsOfFunctionPlot[128].X = -5.69
	pointsOfFunctionPlot[128].Y = 0.003

	pointsOfFunctionPlot[129].X = -5.68
	pointsOfFunctionPlot[129].Y = 0.003

	pointsOfFunctionPlot[130].X = -5.67
	pointsOfFunctionPlot[130].Y = 0.003

	pointsOfFunctionPlot[131].X = -5.66
	pointsOfFunctionPlot[131].Y = 0.003

	pointsOfFunctionPlot[132].X = -5.65
	pointsOfFunctionPlot[132].Y = 0.003

	pointsOfFunctionPlot[133].X = -5.64
	pointsOfFunctionPlot[133].Y = 0.003

	pointsOfFunctionPlot[134].X = -5.63
	pointsOfFunctionPlot[134].Y = 0.003

	pointsOfFunctionPlot[135].X = -5.62
	pointsOfFunctionPlot[135].Y = 0.003

	pointsOfFunctionPlot[136].X = -5.61
	pointsOfFunctionPlot[136].Y = 0.003

	pointsOfFunctionPlot[137].X = -5.6
	pointsOfFunctionPlot[137].Y = 0.003

	pointsOfFunctionPlot[138].X = -5.59
	pointsOfFunctionPlot[138].Y = 0.003

	pointsOfFunctionPlot[139].X = -5.58
	pointsOfFunctionPlot[139].Y = 0.003

	pointsOfFunctionPlot[140].X = -5.57
	pointsOfFunctionPlot[140].Y = 0.003

	pointsOfFunctionPlot[141].X = -5.56
	pointsOfFunctionPlot[141].Y = 0.003

	pointsOfFunctionPlot[142].X = -5.55
	pointsOfFunctionPlot[142].Y = 0.003

	pointsOfFunctionPlot[143].X = -5.54
	pointsOfFunctionPlot[143].Y = 0.003

	pointsOfFunctionPlot[144].X = -5.53
	pointsOfFunctionPlot[144].Y = 0.004

	pointsOfFunctionPlot[145].X = -5.52
	pointsOfFunctionPlot[145].Y = 0.004

	pointsOfFunctionPlot[146].X = -5.51
	pointsOfFunctionPlot[146].Y = 0.004

	pointsOfFunctionPlot[147].X = -5.5
	pointsOfFunctionPlot[147].Y = 0.004

	pointsOfFunctionPlot[148].X = -5.49
	pointsOfFunctionPlot[148].Y = 0.004

	pointsOfFunctionPlot[149].X = -5.48
	pointsOfFunctionPlot[149].Y = 0.004

	pointsOfFunctionPlot[150].X = -5.47
	pointsOfFunctionPlot[150].Y = 0.004

	pointsOfFunctionPlot[151].X = -5.46
	pointsOfFunctionPlot[151].Y = 0.004

	pointsOfFunctionPlot[152].X = -5.45
	pointsOfFunctionPlot[152].Y = 0.004

	pointsOfFunctionPlot[153].X = -5.44
	pointsOfFunctionPlot[153].Y = 0.004

	pointsOfFunctionPlot[154].X = -5.43
	pointsOfFunctionPlot[154].Y = 0.004

	pointsOfFunctionPlot[155].X = -5.42
	pointsOfFunctionPlot[155].Y = 0.004

	pointsOfFunctionPlot[156].X = -5.41
	pointsOfFunctionPlot[156].Y = 0.004

	pointsOfFunctionPlot[157].X = -5.4
	pointsOfFunctionPlot[157].Y = 0.004

	pointsOfFunctionPlot[158].X = -5.39
	pointsOfFunctionPlot[158].Y = 0.004

	pointsOfFunctionPlot[159].X = -5.38
	pointsOfFunctionPlot[159].Y = 0.004

	pointsOfFunctionPlot[160].X = -5.37
	pointsOfFunctionPlot[160].Y = 0.004

	pointsOfFunctionPlot[161].X = -5.36
	pointsOfFunctionPlot[161].Y = 0.004

	pointsOfFunctionPlot[162].X = -5.35
	pointsOfFunctionPlot[162].Y = 0.004

	pointsOfFunctionPlot[163].X = -5.34
	pointsOfFunctionPlot[163].Y = 0.004

	pointsOfFunctionPlot[164].X = -5.33
	pointsOfFunctionPlot[164].Y = 0.004

	pointsOfFunctionPlot[165].X = -5.32
	pointsOfFunctionPlot[165].Y = 0.004

	pointsOfFunctionPlot[166].X = -5.31
	pointsOfFunctionPlot[166].Y = 0.004

	pointsOfFunctionPlot[167].X = -5.3
	pointsOfFunctionPlot[167].Y = 0.005

	pointsOfFunctionPlot[168].X = -5.29
	pointsOfFunctionPlot[168].Y = 0.005

	pointsOfFunctionPlot[169].X = -5.28
	pointsOfFunctionPlot[169].Y = 0.005

	pointsOfFunctionPlot[170].X = -5.27
	pointsOfFunctionPlot[170].Y = 0.005

	pointsOfFunctionPlot[171].X = -5.26
	pointsOfFunctionPlot[171].Y = 0.005

	pointsOfFunctionPlot[172].X = -5.25
	pointsOfFunctionPlot[172].Y = 0.005

	pointsOfFunctionPlot[173].X = -5.24
	pointsOfFunctionPlot[173].Y = 0.005

	pointsOfFunctionPlot[174].X = -5.23
	pointsOfFunctionPlot[174].Y = 0.005

	pointsOfFunctionPlot[175].X = -5.22
	pointsOfFunctionPlot[175].Y = 0.005

	pointsOfFunctionPlot[176].X = -5.21
	pointsOfFunctionPlot[176].Y = 0.005

	pointsOfFunctionPlot[177].X = -5.2
	pointsOfFunctionPlot[177].Y = 0.005

	pointsOfFunctionPlot[178].X = -5.19
	pointsOfFunctionPlot[178].Y = 0.005

	pointsOfFunctionPlot[179].X = -5.18
	pointsOfFunctionPlot[179].Y = 0.005

	pointsOfFunctionPlot[180].X = -5.17
	pointsOfFunctionPlot[180].Y = 0.005

	pointsOfFunctionPlot[181].X = -5.16
	pointsOfFunctionPlot[181].Y = 0.005

	pointsOfFunctionPlot[182].X = -5.15
	pointsOfFunctionPlot[182].Y = 0.005

	pointsOfFunctionPlot[183].X = -5.14
	pointsOfFunctionPlot[183].Y = 0.005

	pointsOfFunctionPlot[184].X = -5.13
	pointsOfFunctionPlot[184].Y = 0.005

	pointsOfFunctionPlot[185].X = -5.12
	pointsOfFunctionPlot[185].Y = 0.006

	pointsOfFunctionPlot[186].X = -5.11
	pointsOfFunctionPlot[186].Y = 0.006

	pointsOfFunctionPlot[187].X = -5.1
	pointsOfFunctionPlot[187].Y = 0.006

	pointsOfFunctionPlot[188].X = -5.09
	pointsOfFunctionPlot[188].Y = 0.006

	pointsOfFunctionPlot[189].X = -5.08
	pointsOfFunctionPlot[189].Y = 0.006

	pointsOfFunctionPlot[190].X = -5.07
	pointsOfFunctionPlot[190].Y = 0.006

	pointsOfFunctionPlot[191].X = -5.06
	pointsOfFunctionPlot[191].Y = 0.006

	pointsOfFunctionPlot[192].X = -5.05
	pointsOfFunctionPlot[192].Y = 0.006

	pointsOfFunctionPlot[193].X = -5.04
	pointsOfFunctionPlot[193].Y = 0.006

	pointsOfFunctionPlot[194].X = -5.03
	pointsOfFunctionPlot[194].Y = 0.006

	pointsOfFunctionPlot[195].X = -5.02
	pointsOfFunctionPlot[195].Y = 0.006

	pointsOfFunctionPlot[196].X = -5.01
	pointsOfFunctionPlot[196].Y = 0.006

	pointsOfFunctionPlot[197].X = -5.0
	pointsOfFunctionPlot[197].Y = 0.006

	pointsOfFunctionPlot[198].X = -4.99
	pointsOfFunctionPlot[198].Y = 0.006

	pointsOfFunctionPlot[199].X = -4.98
	pointsOfFunctionPlot[199].Y = 0.006

	pointsOfFunctionPlot[200].X = -4.97
	pointsOfFunctionPlot[200].Y = 0.006

	pointsOfFunctionPlot[201].X = -4.96
	pointsOfFunctionPlot[201].Y = 0.007

	pointsOfFunctionPlot[202].X = -4.95
	pointsOfFunctionPlot[202].Y = 0.007

	pointsOfFunctionPlot[203].X = -4.94
	pointsOfFunctionPlot[203].Y = 0.007

	pointsOfFunctionPlot[204].X = -4.93
	pointsOfFunctionPlot[204].Y = 0.007

	pointsOfFunctionPlot[205].X = -4.92
	pointsOfFunctionPlot[205].Y = 0.007

	pointsOfFunctionPlot[206].X = -4.91
	pointsOfFunctionPlot[206].Y = 0.007

	pointsOfFunctionPlot[207].X = -4.9
	pointsOfFunctionPlot[207].Y = 0.007

	pointsOfFunctionPlot[208].X = -4.89
	pointsOfFunctionPlot[208].Y = 0.007

	pointsOfFunctionPlot[209].X = -4.88
	pointsOfFunctionPlot[209].Y = 0.007

	pointsOfFunctionPlot[210].X = -4.87
	pointsOfFunctionPlot[210].Y = 0.007

	pointsOfFunctionPlot[211].X = -4.86
	pointsOfFunctionPlot[211].Y = 0.007

	pointsOfFunctionPlot[212].X = -4.85
	pointsOfFunctionPlot[212].Y = 0.007

	pointsOfFunctionPlot[213].X = -4.84
	pointsOfFunctionPlot[213].Y = 0.007

	pointsOfFunctionPlot[214].X = -4.83
	pointsOfFunctionPlot[214].Y = 0.008

	pointsOfFunctionPlot[215].X = -4.82
	pointsOfFunctionPlot[215].Y = 0.008

	pointsOfFunctionPlot[216].X = -4.81
	pointsOfFunctionPlot[216].Y = 0.008

	pointsOfFunctionPlot[217].X = -4.8
	pointsOfFunctionPlot[217].Y = 0.008

	pointsOfFunctionPlot[218].X = -4.79
	pointsOfFunctionPlot[218].Y = 0.008

	pointsOfFunctionPlot[219].X = -4.78
	pointsOfFunctionPlot[219].Y = 0.008

	pointsOfFunctionPlot[220].X = -4.77
	pointsOfFunctionPlot[220].Y = 0.008

	pointsOfFunctionPlot[221].X = -4.76
	pointsOfFunctionPlot[221].Y = 0.008

	pointsOfFunctionPlot[222].X = -4.75
	pointsOfFunctionPlot[222].Y = 0.008

	pointsOfFunctionPlot[223].X = -4.74
	pointsOfFunctionPlot[223].Y = 0.008

	pointsOfFunctionPlot[224].X = -4.73
	pointsOfFunctionPlot[224].Y = 0.008

	pointsOfFunctionPlot[225].X = -4.72
	pointsOfFunctionPlot[225].Y = 0.008

	pointsOfFunctionPlot[226].X = -4.71
	pointsOfFunctionPlot[226].Y = 0.009

	pointsOfFunctionPlot[227].X = -4.7
	pointsOfFunctionPlot[227].Y = 0.009

	pointsOfFunctionPlot[228].X = -4.69
	pointsOfFunctionPlot[228].Y = 0.009

	pointsOfFunctionPlot[229].X = -4.68
	pointsOfFunctionPlot[229].Y = 0.009

	pointsOfFunctionPlot[230].X = -4.67
	pointsOfFunctionPlot[230].Y = 0.009

	pointsOfFunctionPlot[231].X = -4.66
	pointsOfFunctionPlot[231].Y = 0.009

	pointsOfFunctionPlot[232].X = -4.65
	pointsOfFunctionPlot[232].Y = 0.009

	pointsOfFunctionPlot[233].X = -4.64
	pointsOfFunctionPlot[233].Y = 0.009

	pointsOfFunctionPlot[234].X = -4.63
	pointsOfFunctionPlot[234].Y = 0.009

	pointsOfFunctionPlot[235].X = -4.62
	pointsOfFunctionPlot[235].Y = 0.009

	pointsOfFunctionPlot[236].X = -4.61
	pointsOfFunctionPlot[236].Y = 0.01

	pointsOfFunctionPlot[237].X = -4.6
	pointsOfFunctionPlot[237].Y = 0.01

	pointsOfFunctionPlot[238].X = -4.59
	pointsOfFunctionPlot[238].Y = 0.01

	pointsOfFunctionPlot[239].X = -4.58
	pointsOfFunctionPlot[239].Y = 0.01

	pointsOfFunctionPlot[240].X = -4.57
	pointsOfFunctionPlot[240].Y = 0.01

	pointsOfFunctionPlot[241].X = -4.56
	pointsOfFunctionPlot[241].Y = 0.01

	pointsOfFunctionPlot[242].X = -4.55
	pointsOfFunctionPlot[242].Y = 0.01

	pointsOfFunctionPlot[243].X = -4.54
	pointsOfFunctionPlot[243].Y = 0.01

	pointsOfFunctionPlot[244].X = -4.53
	pointsOfFunctionPlot[244].Y = 0.01

	pointsOfFunctionPlot[245].X = -4.52
	pointsOfFunctionPlot[245].Y = 0.01

	pointsOfFunctionPlot[246].X = -4.51
	pointsOfFunctionPlot[246].Y = 0.011

	pointsOfFunctionPlot[247].X = -4.5
	pointsOfFunctionPlot[247].Y = 0.011

	pointsOfFunctionPlot[248].X = -4.49
	pointsOfFunctionPlot[248].Y = 0.011

	pointsOfFunctionPlot[249].X = -4.48
	pointsOfFunctionPlot[249].Y = 0.011

	pointsOfFunctionPlot[250].X = -4.47
	pointsOfFunctionPlot[250].Y = 0.011

	pointsOfFunctionPlot[251].X = -4.46
	pointsOfFunctionPlot[251].Y = 0.011

	pointsOfFunctionPlot[252].X = -4.45
	pointsOfFunctionPlot[252].Y = 0.011

	pointsOfFunctionPlot[253].X = -4.44
	pointsOfFunctionPlot[253].Y = 0.011

	pointsOfFunctionPlot[254].X = -4.43
	pointsOfFunctionPlot[254].Y = 0.011

	pointsOfFunctionPlot[255].X = -4.42
	pointsOfFunctionPlot[255].Y = 0.012

	pointsOfFunctionPlot[256].X = -4.41
	pointsOfFunctionPlot[256].Y = 0.012

	pointsOfFunctionPlot[257].X = -4.4
	pointsOfFunctionPlot[257].Y = 0.012

	pointsOfFunctionPlot[258].X = -4.39
	pointsOfFunctionPlot[258].Y = 0.012

	pointsOfFunctionPlot[259].X = -4.38
	pointsOfFunctionPlot[259].Y = 0.012

	pointsOfFunctionPlot[260].X = -4.37
	pointsOfFunctionPlot[260].Y = 0.012

	pointsOfFunctionPlot[261].X = -4.36
	pointsOfFunctionPlot[261].Y = 0.012

	pointsOfFunctionPlot[262].X = -4.35
	pointsOfFunctionPlot[262].Y = 0.012

	pointsOfFunctionPlot[263].X = -4.34
	pointsOfFunctionPlot[263].Y = 0.013

	pointsOfFunctionPlot[264].X = -4.33
	pointsOfFunctionPlot[264].Y = 0.013

	pointsOfFunctionPlot[265].X = -4.32
	pointsOfFunctionPlot[265].Y = 0.013

	pointsOfFunctionPlot[266].X = -4.31
	pointsOfFunctionPlot[266].Y = 0.013

	pointsOfFunctionPlot[267].X = -4.3
	pointsOfFunctionPlot[267].Y = 0.013

	pointsOfFunctionPlot[268].X = -4.29
	pointsOfFunctionPlot[268].Y = 0.013

	pointsOfFunctionPlot[269].X = -4.28
	pointsOfFunctionPlot[269].Y = 0.013

	pointsOfFunctionPlot[270].X = -4.27
	pointsOfFunctionPlot[270].Y = 0.014

	pointsOfFunctionPlot[271].X = -4.26
	pointsOfFunctionPlot[271].Y = 0.014

	pointsOfFunctionPlot[272].X = -4.25
	pointsOfFunctionPlot[272].Y = 0.014

	pointsOfFunctionPlot[273].X = -4.24
	pointsOfFunctionPlot[273].Y = 0.014

	pointsOfFunctionPlot[274].X = -4.23
	pointsOfFunctionPlot[274].Y = 0.014

	pointsOfFunctionPlot[275].X = -4.22
	pointsOfFunctionPlot[275].Y = 0.014

	pointsOfFunctionPlot[276].X = -4.21
	pointsOfFunctionPlot[276].Y = 0.014

	pointsOfFunctionPlot[277].X = -4.2
	pointsOfFunctionPlot[277].Y = 0.015

	pointsOfFunctionPlot[278].X = -4.19
	pointsOfFunctionPlot[278].Y = 0.015

	pointsOfFunctionPlot[279].X = -4.18
	pointsOfFunctionPlot[279].Y = 0.015

	pointsOfFunctionPlot[280].X = -4.17
	pointsOfFunctionPlot[280].Y = 0.015

	pointsOfFunctionPlot[281].X = -4.16
	pointsOfFunctionPlot[281].Y = 0.015

	pointsOfFunctionPlot[282].X = -4.15
	pointsOfFunctionPlot[282].Y = 0.015

	pointsOfFunctionPlot[283].X = -4.14
	pointsOfFunctionPlot[283].Y = 0.015

	pointsOfFunctionPlot[284].X = -4.13
	pointsOfFunctionPlot[284].Y = 0.016

	pointsOfFunctionPlot[285].X = -4.12
	pointsOfFunctionPlot[285].Y = 0.016

	pointsOfFunctionPlot[286].X = -4.11
	pointsOfFunctionPlot[286].Y = 0.016

	pointsOfFunctionPlot[287].X = -4.1
	pointsOfFunctionPlot[287].Y = 0.016

	pointsOfFunctionPlot[288].X = -4.09
	pointsOfFunctionPlot[288].Y = 0.016

	pointsOfFunctionPlot[289].X = -4.08
	pointsOfFunctionPlot[289].Y = 0.016

	pointsOfFunctionPlot[290].X = -4.07
	pointsOfFunctionPlot[290].Y = 0.017

	pointsOfFunctionPlot[291].X = -4.06
	pointsOfFunctionPlot[291].Y = 0.017

	pointsOfFunctionPlot[292].X = -4.05
	pointsOfFunctionPlot[292].Y = 0.017

	pointsOfFunctionPlot[293].X = -4.04
	pointsOfFunctionPlot[293].Y = 0.017

	pointsOfFunctionPlot[294].X = -4.03
	pointsOfFunctionPlot[294].Y = 0.017

	pointsOfFunctionPlot[295].X = -4.02
	pointsOfFunctionPlot[295].Y = 0.018

	pointsOfFunctionPlot[296].X = -4.01
	pointsOfFunctionPlot[296].Y = 0.018

	pointsOfFunctionPlot[297].X = -4.0
	pointsOfFunctionPlot[297].Y = 0.018

	pointsOfFunctionPlot[298].X = -3.99
	pointsOfFunctionPlot[298].Y = 0.018

	pointsOfFunctionPlot[299].X = -3.98
	pointsOfFunctionPlot[299].Y = 0.018

	pointsOfFunctionPlot[300].X = -3.97
	pointsOfFunctionPlot[300].Y = 0.018

	pointsOfFunctionPlot[301].X = -3.96
	pointsOfFunctionPlot[301].Y = 0.019

	pointsOfFunctionPlot[302].X = -3.95
	pointsOfFunctionPlot[302].Y = 0.019

	pointsOfFunctionPlot[303].X = -3.94
	pointsOfFunctionPlot[303].Y = 0.019

	pointsOfFunctionPlot[304].X = -3.93
	pointsOfFunctionPlot[304].Y = 0.019

	pointsOfFunctionPlot[305].X = -3.92
	pointsOfFunctionPlot[305].Y = 0.019

	pointsOfFunctionPlot[306].X = -3.91
	pointsOfFunctionPlot[306].Y = 0.02

	pointsOfFunctionPlot[307].X = -3.9
	pointsOfFunctionPlot[307].Y = 0.02

	pointsOfFunctionPlot[308].X = -3.89
	pointsOfFunctionPlot[308].Y = 0.02

	pointsOfFunctionPlot[309].X = -3.88
	pointsOfFunctionPlot[309].Y = 0.02

	pointsOfFunctionPlot[310].X = -3.87
	pointsOfFunctionPlot[310].Y = 0.02

	pointsOfFunctionPlot[311].X = -3.86
	pointsOfFunctionPlot[311].Y = 0.021

	pointsOfFunctionPlot[312].X = -3.85
	pointsOfFunctionPlot[312].Y = 0.021

	pointsOfFunctionPlot[313].X = -3.84
	pointsOfFunctionPlot[313].Y = 0.021

	pointsOfFunctionPlot[314].X = -3.83
	pointsOfFunctionPlot[314].Y = 0.021

	pointsOfFunctionPlot[315].X = -3.82
	pointsOfFunctionPlot[315].Y = 0.021

	pointsOfFunctionPlot[316].X = -3.81
	pointsOfFunctionPlot[316].Y = 0.022

	pointsOfFunctionPlot[317].X = -3.8
	pointsOfFunctionPlot[317].Y = 0.022

	pointsOfFunctionPlot[318].X = -3.79
	pointsOfFunctionPlot[318].Y = 0.022

	pointsOfFunctionPlot[319].X = -3.78
	pointsOfFunctionPlot[319].Y = 0.022

	pointsOfFunctionPlot[320].X = -3.77
	pointsOfFunctionPlot[320].Y = 0.023

	pointsOfFunctionPlot[321].X = -3.76
	pointsOfFunctionPlot[321].Y = 0.023

	pointsOfFunctionPlot[322].X = -3.75
	pointsOfFunctionPlot[322].Y = 0.023

	pointsOfFunctionPlot[323].X = -3.74
	pointsOfFunctionPlot[323].Y = 0.023

	pointsOfFunctionPlot[324].X = -3.73
	pointsOfFunctionPlot[324].Y = 0.024

	pointsOfFunctionPlot[325].X = -3.72
	pointsOfFunctionPlot[325].Y = 0.024

	pointsOfFunctionPlot[326].X = -3.71
	pointsOfFunctionPlot[326].Y = 0.024

	pointsOfFunctionPlot[327].X = -3.7
	pointsOfFunctionPlot[327].Y = 0.024

	pointsOfFunctionPlot[328].X = -3.69
	pointsOfFunctionPlot[328].Y = 0.025

	pointsOfFunctionPlot[329].X = -3.68
	pointsOfFunctionPlot[329].Y = 0.025

	pointsOfFunctionPlot[330].X = -3.67
	pointsOfFunctionPlot[330].Y = 0.025

	pointsOfFunctionPlot[331].X = -3.66
	pointsOfFunctionPlot[331].Y = 0.025

	pointsOfFunctionPlot[332].X = -3.65
	pointsOfFunctionPlot[332].Y = 0.026

	pointsOfFunctionPlot[333].X = -3.64
	pointsOfFunctionPlot[333].Y = 0.026

	pointsOfFunctionPlot[334].X = -3.63
	pointsOfFunctionPlot[334].Y = 0.026

	pointsOfFunctionPlot[335].X = -3.62
	pointsOfFunctionPlot[335].Y = 0.026

	pointsOfFunctionPlot[336].X = -3.61
	pointsOfFunctionPlot[336].Y = 0.027

	pointsOfFunctionPlot[337].X = -3.6
	pointsOfFunctionPlot[337].Y = 0.027

	pointsOfFunctionPlot[338].X = -3.59
	pointsOfFunctionPlot[338].Y = 0.027

	pointsOfFunctionPlot[339].X = -3.58
	pointsOfFunctionPlot[339].Y = 0.027

	pointsOfFunctionPlot[340].X = -3.57
	pointsOfFunctionPlot[340].Y = 0.028

	pointsOfFunctionPlot[341].X = -3.56
	pointsOfFunctionPlot[341].Y = 0.028

	pointsOfFunctionPlot[342].X = -3.55
	pointsOfFunctionPlot[342].Y = 0.028

	pointsOfFunctionPlot[343].X = -3.54
	pointsOfFunctionPlot[343].Y = 0.029

	pointsOfFunctionPlot[344].X = -3.53
	pointsOfFunctionPlot[344].Y = 0.029

	pointsOfFunctionPlot[345].X = -3.52
	pointsOfFunctionPlot[345].Y = 0.029

	pointsOfFunctionPlot[346].X = -3.51
	pointsOfFunctionPlot[346].Y = 0.029

	pointsOfFunctionPlot[347].X = -3.5
	pointsOfFunctionPlot[347].Y = 0.03

	pointsOfFunctionPlot[348].X = -3.49
	pointsOfFunctionPlot[348].Y = 0.03

	pointsOfFunctionPlot[349].X = -3.48
	pointsOfFunctionPlot[349].Y = 0.03

	pointsOfFunctionPlot[350].X = -3.47
	pointsOfFunctionPlot[350].Y = 0.031

	pointsOfFunctionPlot[351].X = -3.46
	pointsOfFunctionPlot[351].Y = 0.031

	pointsOfFunctionPlot[352].X = -3.45
	pointsOfFunctionPlot[352].Y = 0.031

	pointsOfFunctionPlot[353].X = -3.44
	pointsOfFunctionPlot[353].Y = 0.032

	pointsOfFunctionPlot[354].X = -3.43
	pointsOfFunctionPlot[354].Y = 0.032

	pointsOfFunctionPlot[355].X = -3.42
	pointsOfFunctionPlot[355].Y = 0.032

	pointsOfFunctionPlot[356].X = -3.41
	pointsOfFunctionPlot[356].Y = 0.033

	pointsOfFunctionPlot[357].X = -3.4
	pointsOfFunctionPlot[357].Y = 0.033

	pointsOfFunctionPlot[358].X = -3.39
	pointsOfFunctionPlot[358].Y = 0.033

	pointsOfFunctionPlot[359].X = -3.38
	pointsOfFunctionPlot[359].Y = 0.034

	pointsOfFunctionPlot[360].X = -3.37
	pointsOfFunctionPlot[360].Y = 0.034

	pointsOfFunctionPlot[361].X = -3.36
	pointsOfFunctionPlot[361].Y = 0.034

	pointsOfFunctionPlot[362].X = -3.35
	pointsOfFunctionPlot[362].Y = 0.035

	pointsOfFunctionPlot[363].X = -3.34
	pointsOfFunctionPlot[363].Y = 0.035

	pointsOfFunctionPlot[364].X = -3.33
	pointsOfFunctionPlot[364].Y = 0.035

	pointsOfFunctionPlot[365].X = -3.32
	pointsOfFunctionPlot[365].Y = 0.036

	pointsOfFunctionPlot[366].X = -3.31
	pointsOfFunctionPlot[366].Y = 0.036

	pointsOfFunctionPlot[367].X = -3.3
	pointsOfFunctionPlot[367].Y = 0.036

	pointsOfFunctionPlot[368].X = -3.29
	pointsOfFunctionPlot[368].Y = 0.037

	pointsOfFunctionPlot[369].X = -3.28
	pointsOfFunctionPlot[369].Y = 0.037

	pointsOfFunctionPlot[370].X = -3.27
	pointsOfFunctionPlot[370].Y = 0.038

	pointsOfFunctionPlot[371].X = -3.26
	pointsOfFunctionPlot[371].Y = 0.038

	pointsOfFunctionPlot[372].X = -3.25
	pointsOfFunctionPlot[372].Y = 0.038

	pointsOfFunctionPlot[373].X = -3.24
	pointsOfFunctionPlot[373].Y = 0.039

	pointsOfFunctionPlot[374].X = -3.23
	pointsOfFunctionPlot[374].Y = 0.039

	pointsOfFunctionPlot[375].X = -3.22
	pointsOfFunctionPlot[375].Y = 0.04

	pointsOfFunctionPlot[376].X = -3.21
	pointsOfFunctionPlot[376].Y = 0.04

	pointsOfFunctionPlot[377].X = -3.2
	pointsOfFunctionPlot[377].Y = 0.041

	pointsOfFunctionPlot[378].X = -3.19
	pointsOfFunctionPlot[378].Y = 0.041

	pointsOfFunctionPlot[379].X = -3.18
	pointsOfFunctionPlot[379].Y = 0.041

	pointsOfFunctionPlot[380].X = -3.17
	pointsOfFunctionPlot[380].Y = 0.042

	pointsOfFunctionPlot[381].X = -3.16
	pointsOfFunctionPlot[381].Y = 0.042

	pointsOfFunctionPlot[382].X = -3.15
	pointsOfFunctionPlot[382].Y = 0.042

	pointsOfFunctionPlot[383].X = -3.14
	pointsOfFunctionPlot[383].Y = 0.043

	pointsOfFunctionPlot[384].X = -3.13
	pointsOfFunctionPlot[384].Y = 0.043

	pointsOfFunctionPlot[385].X = -3.12
	pointsOfFunctionPlot[385].Y = 0.044

	pointsOfFunctionPlot[386].X = -3.11
	pointsOfFunctionPlot[386].Y = 0.044

	pointsOfFunctionPlot[387].X = -3.1
	pointsOfFunctionPlot[387].Y = 0.045

	pointsOfFunctionPlot[388].X = -3.09
	pointsOfFunctionPlot[388].Y = 0.045

	pointsOfFunctionPlot[389].X = -3.08
	pointsOfFunctionPlot[389].Y = 0.046

	pointsOfFunctionPlot[390].X = -3.07
	pointsOfFunctionPlot[390].Y = 0.046

	pointsOfFunctionPlot[391].X = -3.06
	pointsOfFunctionPlot[391].Y = 0.046

	pointsOfFunctionPlot[392].X = -3.05
	pointsOfFunctionPlot[392].Y = 0.047

	pointsOfFunctionPlot[393].X = -3.04
	pointsOfFunctionPlot[393].Y = 0.047

	pointsOfFunctionPlot[394].X = -3.03
	pointsOfFunctionPlot[394].Y = 0.048

	pointsOfFunctionPlot[395].X = -3.02
	pointsOfFunctionPlot[395].Y = 0.048

	pointsOfFunctionPlot[396].X = -3.01
	pointsOfFunctionPlot[396].Y = 0.049

	pointsOfFunctionPlot[397].X = -3.0
	pointsOfFunctionPlot[397].Y = 0.049

	pointsOfFunctionPlot[398].X = -2.99
	pointsOfFunctionPlot[398].Y = 0.05

	pointsOfFunctionPlot[399].X = -2.98
	pointsOfFunctionPlot[399].Y = 0.05

	pointsOfFunctionPlot[400].X = -2.97
	pointsOfFunctionPlot[400].Y = 0.051

	pointsOfFunctionPlot[401].X = -2.96
	pointsOfFunctionPlot[401].Y = 0.051

	pointsOfFunctionPlot[402].X = -2.95
	pointsOfFunctionPlot[402].Y = 0.052

	pointsOfFunctionPlot[403].X = -2.94
	pointsOfFunctionPlot[403].Y = 0.052

	pointsOfFunctionPlot[404].X = -2.93
	pointsOfFunctionPlot[404].Y = 0.053

	pointsOfFunctionPlot[405].X = -2.92
	pointsOfFunctionPlot[405].Y = 0.053

	pointsOfFunctionPlot[406].X = -2.91
	pointsOfFunctionPlot[406].Y = 0.054

	pointsOfFunctionPlot[407].X = -2.9
	pointsOfFunctionPlot[407].Y = 0.055

	pointsOfFunctionPlot[408].X = -2.89
	pointsOfFunctionPlot[408].Y = 0.055

	pointsOfFunctionPlot[409].X = -2.88
	pointsOfFunctionPlot[409].Y = 0.056

	pointsOfFunctionPlot[410].X = -2.87
	pointsOfFunctionPlot[410].Y = 0.056

	pointsOfFunctionPlot[411].X = -2.86
	pointsOfFunctionPlot[411].Y = 0.057

	pointsOfFunctionPlot[412].X = -2.85
	pointsOfFunctionPlot[412].Y = 0.057

	pointsOfFunctionPlot[413].X = -2.84
	pointsOfFunctionPlot[413].Y = 0.058

	pointsOfFunctionPlot[414].X = -2.83
	pointsOfFunctionPlot[414].Y = 0.059

	pointsOfFunctionPlot[415].X = -2.82
	pointsOfFunctionPlot[415].Y = 0.059

	pointsOfFunctionPlot[416].X = -2.81
	pointsOfFunctionPlot[416].Y = 0.06

	pointsOfFunctionPlot[417].X = -2.8
	pointsOfFunctionPlot[417].Y = 0.06

	pointsOfFunctionPlot[418].X = -2.79
	pointsOfFunctionPlot[418].Y = 0.061

	pointsOfFunctionPlot[419].X = -2.78
	pointsOfFunctionPlot[419].Y = 0.062

	pointsOfFunctionPlot[420].X = -2.77
	pointsOfFunctionPlot[420].Y = 0.062

	pointsOfFunctionPlot[421].X = -2.76
	pointsOfFunctionPlot[421].Y = 0.063

	pointsOfFunctionPlot[422].X = -2.75
	pointsOfFunctionPlot[422].Y = 0.063

	pointsOfFunctionPlot[423].X = -2.74
	pointsOfFunctionPlot[423].Y = 0.064

	pointsOfFunctionPlot[424].X = -2.73
	pointsOfFunctionPlot[424].Y = 0.065

	pointsOfFunctionPlot[425].X = -2.72
	pointsOfFunctionPlot[425].Y = 0.065

	pointsOfFunctionPlot[426].X = -2.71
	pointsOfFunctionPlot[426].Y = 0.066

	pointsOfFunctionPlot[427].X = -2.7
	pointsOfFunctionPlot[427].Y = 0.067

	pointsOfFunctionPlot[428].X = -2.69
	pointsOfFunctionPlot[428].Y = 0.067

	pointsOfFunctionPlot[429].X = -2.68
	pointsOfFunctionPlot[429].Y = 0.068

	pointsOfFunctionPlot[430].X = -2.67
	pointsOfFunctionPlot[430].Y = 0.069

	pointsOfFunctionPlot[431].X = -2.66
	pointsOfFunctionPlot[431].Y = 0.069

	pointsOfFunctionPlot[432].X = -2.65
	pointsOfFunctionPlot[432].Y = 0.07

	pointsOfFunctionPlot[433].X = -2.64
	pointsOfFunctionPlot[433].Y = 0.071

	pointsOfFunctionPlot[434].X = -2.63
	pointsOfFunctionPlot[434].Y = 0.072

	pointsOfFunctionPlot[435].X = -2.62
	pointsOfFunctionPlot[435].Y = 0.072

	pointsOfFunctionPlot[436].X = -2.61
	pointsOfFunctionPlot[436].Y = 0.073

	pointsOfFunctionPlot[437].X = -2.6
	pointsOfFunctionPlot[437].Y = 0.074

	pointsOfFunctionPlot[438].X = -2.59
	pointsOfFunctionPlot[438].Y = 0.075

	pointsOfFunctionPlot[439].X = -2.58
	pointsOfFunctionPlot[439].Y = 0.075

	pointsOfFunctionPlot[440].X = -2.57
	pointsOfFunctionPlot[440].Y = 0.076

	pointsOfFunctionPlot[441].X = -2.56
	pointsOfFunctionPlot[441].Y = 0.077

	pointsOfFunctionPlot[442].X = -2.55
	pointsOfFunctionPlot[442].Y = 0.078

	pointsOfFunctionPlot[443].X = -2.54
	pointsOfFunctionPlot[443].Y = 0.078

	pointsOfFunctionPlot[444].X = -2.53
	pointsOfFunctionPlot[444].Y = 0.079

	pointsOfFunctionPlot[445].X = -2.52
	pointsOfFunctionPlot[445].Y = 0.08

	pointsOfFunctionPlot[446].X = -2.51
	pointsOfFunctionPlot[446].Y = 0.081

	pointsOfFunctionPlot[447].X = -2.5
	pointsOfFunctionPlot[447].Y = 0.082

	pointsOfFunctionPlot[448].X = -2.49
	pointsOfFunctionPlot[448].Y = 0.082

	pointsOfFunctionPlot[449].X = -2.48
	pointsOfFunctionPlot[449].Y = 0.083

	pointsOfFunctionPlot[450].X = -2.47
	pointsOfFunctionPlot[450].Y = 0.084

	pointsOfFunctionPlot[451].X = -2.46
	pointsOfFunctionPlot[451].Y = 0.085

	pointsOfFunctionPlot[452].X = -2.45
	pointsOfFunctionPlot[452].Y = 0.086

	pointsOfFunctionPlot[453].X = -2.44
	pointsOfFunctionPlot[453].Y = 0.087

	pointsOfFunctionPlot[454].X = -2.43
	pointsOfFunctionPlot[454].Y = 0.088

	pointsOfFunctionPlot[455].X = -2.42
	pointsOfFunctionPlot[455].Y = 0.088

	pointsOfFunctionPlot[456].X = -2.41
	pointsOfFunctionPlot[456].Y = 0.089

	pointsOfFunctionPlot[457].X = -2.4
	pointsOfFunctionPlot[457].Y = 0.09

	pointsOfFunctionPlot[458].X = -2.39
	pointsOfFunctionPlot[458].Y = 0.091

	pointsOfFunctionPlot[459].X = -2.38
	pointsOfFunctionPlot[459].Y = 0.092

	pointsOfFunctionPlot[460].X = -2.37
	pointsOfFunctionPlot[460].Y = 0.093

	pointsOfFunctionPlot[461].X = -2.36
	pointsOfFunctionPlot[461].Y = 0.094

	pointsOfFunctionPlot[462].X = -2.35
	pointsOfFunctionPlot[462].Y = 0.095

	pointsOfFunctionPlot[463].X = -2.34
	pointsOfFunctionPlot[463].Y = 0.096

	pointsOfFunctionPlot[464].X = -2.33
	pointsOfFunctionPlot[464].Y = 0.097

	pointsOfFunctionPlot[465].X = -2.32
	pointsOfFunctionPlot[465].Y = 0.098

	pointsOfFunctionPlot[466].X = -2.31
	pointsOfFunctionPlot[466].Y = 0.099

	pointsOfFunctionPlot[467].X = -2.3
	pointsOfFunctionPlot[467].Y = 0.1

	pointsOfFunctionPlot[468].X = -2.29
	pointsOfFunctionPlot[468].Y = 0.101

	pointsOfFunctionPlot[469].X = -2.28
	pointsOfFunctionPlot[469].Y = 0.102

	pointsOfFunctionPlot[470].X = -2.27
	pointsOfFunctionPlot[470].Y = 0.103

	pointsOfFunctionPlot[471].X = -2.26
	pointsOfFunctionPlot[471].Y = 0.104

	pointsOfFunctionPlot[472].X = -2.25
	pointsOfFunctionPlot[472].Y = 0.105

	pointsOfFunctionPlot[473].X = -2.24
	pointsOfFunctionPlot[473].Y = 0.106

	pointsOfFunctionPlot[474].X = -2.23
	pointsOfFunctionPlot[474].Y = 0.107

	pointsOfFunctionPlot[475].X = -2.22
	pointsOfFunctionPlot[475].Y = 0.108

	pointsOfFunctionPlot[476].X = -2.21
	pointsOfFunctionPlot[476].Y = 0.109

	pointsOfFunctionPlot[477].X = -2.2
	pointsOfFunctionPlot[477].Y = 0.11

	pointsOfFunctionPlot[478].X = -2.19
	pointsOfFunctionPlot[478].Y = 0.111

	pointsOfFunctionPlot[479].X = -2.18
	pointsOfFunctionPlot[479].Y = 0.113

	pointsOfFunctionPlot[480].X = -2.17
	pointsOfFunctionPlot[480].Y = 0.114

	pointsOfFunctionPlot[481].X = -2.16
	pointsOfFunctionPlot[481].Y = 0.115

	pointsOfFunctionPlot[482].X = -2.15
	pointsOfFunctionPlot[482].Y = 0.116

	pointsOfFunctionPlot[483].X = -2.14
	pointsOfFunctionPlot[483].Y = 0.117

	pointsOfFunctionPlot[484].X = -2.13
	pointsOfFunctionPlot[484].Y = 0.118

	pointsOfFunctionPlot[485].X = -2.12
	pointsOfFunctionPlot[485].Y = 0.12

	pointsOfFunctionPlot[486].X = -2.11
	pointsOfFunctionPlot[486].Y = 0.121

	pointsOfFunctionPlot[487].X = -2.1
	pointsOfFunctionPlot[487].Y = 0.122

	pointsOfFunctionPlot[488].X = -2.09
	pointsOfFunctionPlot[488].Y = 0.123

	pointsOfFunctionPlot[489].X = -2.08
	pointsOfFunctionPlot[489].Y = 0.124

	pointsOfFunctionPlot[490].X = -2.07
	pointsOfFunctionPlot[490].Y = 0.126

	pointsOfFunctionPlot[491].X = -2.06
	pointsOfFunctionPlot[491].Y = 0.127

	pointsOfFunctionPlot[492].X = -2.05
	pointsOfFunctionPlot[492].Y = 0.128

	pointsOfFunctionPlot[493].X = -2.04
	pointsOfFunctionPlot[493].Y = 0.13

	pointsOfFunctionPlot[494].X = -2.03
	pointsOfFunctionPlot[494].Y = 0.131

	pointsOfFunctionPlot[495].X = -2.02
	pointsOfFunctionPlot[495].Y = 0.132

	pointsOfFunctionPlot[496].X = -2.01
	pointsOfFunctionPlot[496].Y = 0.134

	pointsOfFunctionPlot[497].X = -2.0
	pointsOfFunctionPlot[497].Y = 0.135

	pointsOfFunctionPlot[498].X = -1.99
	pointsOfFunctionPlot[498].Y = 0.136

	pointsOfFunctionPlot[499].X = -1.98
	pointsOfFunctionPlot[499].Y = 0.138

	pointsOfFunctionPlot[500].X = -1.97
	pointsOfFunctionPlot[500].Y = 0.139

	pointsOfFunctionPlot[501].X = -1.96
	pointsOfFunctionPlot[501].Y = 0.14

	pointsOfFunctionPlot[502].X = -1.95
	pointsOfFunctionPlot[502].Y = 0.142

	pointsOfFunctionPlot[503].X = -1.94
	pointsOfFunctionPlot[503].Y = 0.143

	pointsOfFunctionPlot[504].X = -1.93
	pointsOfFunctionPlot[504].Y = 0.145

	pointsOfFunctionPlot[505].X = -1.92
	pointsOfFunctionPlot[505].Y = 0.146

	pointsOfFunctionPlot[506].X = -1.91
	pointsOfFunctionPlot[506].Y = 0.148

	pointsOfFunctionPlot[507].X = -1.9
	pointsOfFunctionPlot[507].Y = 0.149

	pointsOfFunctionPlot[508].X = -1.89
	pointsOfFunctionPlot[508].Y = 0.151

	pointsOfFunctionPlot[509].X = -1.88
	pointsOfFunctionPlot[509].Y = 0.152

	pointsOfFunctionPlot[510].X = -1.87
	pointsOfFunctionPlot[510].Y = 0.154

	pointsOfFunctionPlot[511].X = -1.86
	pointsOfFunctionPlot[511].Y = 0.155

	pointsOfFunctionPlot[512].X = -1.85
	pointsOfFunctionPlot[512].Y = 0.157

	pointsOfFunctionPlot[513].X = -1.84
	pointsOfFunctionPlot[513].Y = 0.158

	pointsOfFunctionPlot[514].X = -1.83
	pointsOfFunctionPlot[514].Y = 0.16

	pointsOfFunctionPlot[515].X = -1.82
	pointsOfFunctionPlot[515].Y = 0.162

	pointsOfFunctionPlot[516].X = -1.81
	pointsOfFunctionPlot[516].Y = 0.163

	pointsOfFunctionPlot[517].X = -1.8
	pointsOfFunctionPlot[517].Y = 0.165

	pointsOfFunctionPlot[518].X = -1.79
	pointsOfFunctionPlot[518].Y = 0.167

	pointsOfFunctionPlot[519].X = -1.78
	pointsOfFunctionPlot[519].Y = 0.168

	pointsOfFunctionPlot[520].X = -1.77
	pointsOfFunctionPlot[520].Y = 0.17

	pointsOfFunctionPlot[521].X = -1.76
	pointsOfFunctionPlot[521].Y = 0.172

	pointsOfFunctionPlot[522].X = -1.75
	pointsOfFunctionPlot[522].Y = 0.173

	pointsOfFunctionPlot[523].X = -1.74
	pointsOfFunctionPlot[523].Y = 0.175

	pointsOfFunctionPlot[524].X = -1.73
	pointsOfFunctionPlot[524].Y = 0.177

	pointsOfFunctionPlot[525].X = -1.72
	pointsOfFunctionPlot[525].Y = 0.179

	pointsOfFunctionPlot[526].X = -1.71
	pointsOfFunctionPlot[526].Y = 0.18

	pointsOfFunctionPlot[527].X = -1.7
	pointsOfFunctionPlot[527].Y = 0.182

	pointsOfFunctionPlot[528].X = -1.69
	pointsOfFunctionPlot[528].Y = 0.184

	pointsOfFunctionPlot[529].X = -1.68
	pointsOfFunctionPlot[529].Y = 0.186

	pointsOfFunctionPlot[530].X = -1.67
	pointsOfFunctionPlot[530].Y = 0.188

	pointsOfFunctionPlot[531].X = -1.66
	pointsOfFunctionPlot[531].Y = 0.19

	pointsOfFunctionPlot[532].X = -1.65
	pointsOfFunctionPlot[532].Y = 0.192

	pointsOfFunctionPlot[533].X = -1.64
	pointsOfFunctionPlot[533].Y = 0.194

	pointsOfFunctionPlot[534].X = -1.63
	pointsOfFunctionPlot[534].Y = 0.195

	pointsOfFunctionPlot[535].X = -1.62
	pointsOfFunctionPlot[535].Y = 0.197

	pointsOfFunctionPlot[536].X = -1.61
	pointsOfFunctionPlot[536].Y = 0.199

	pointsOfFunctionPlot[537].X = -1.6
	pointsOfFunctionPlot[537].Y = 0.201

	pointsOfFunctionPlot[538].X = -1.59
	pointsOfFunctionPlot[538].Y = 0.203

	pointsOfFunctionPlot[539].X = -1.58
	pointsOfFunctionPlot[539].Y = 0.206

	pointsOfFunctionPlot[540].X = -1.57
	pointsOfFunctionPlot[540].Y = 0.208

	pointsOfFunctionPlot[541].X = -1.56
	pointsOfFunctionPlot[541].Y = 0.21

	pointsOfFunctionPlot[542].X = -1.55
	pointsOfFunctionPlot[542].Y = 0.212

	pointsOfFunctionPlot[543].X = -1.54
	pointsOfFunctionPlot[543].Y = 0.214

	pointsOfFunctionPlot[544].X = -1.53
	pointsOfFunctionPlot[544].Y = 0.216

	pointsOfFunctionPlot[545].X = -1.52
	pointsOfFunctionPlot[545].Y = 0.218

	pointsOfFunctionPlot[546].X = -1.51
	pointsOfFunctionPlot[546].Y = 0.22

	pointsOfFunctionPlot[547].X = -1.5
	pointsOfFunctionPlot[547].Y = 0.223

	pointsOfFunctionPlot[548].X = -1.49
	pointsOfFunctionPlot[548].Y = 0.225

	pointsOfFunctionPlot[549].X = -1.48
	pointsOfFunctionPlot[549].Y = 0.227

	pointsOfFunctionPlot[550].X = -1.47
	pointsOfFunctionPlot[550].Y = 0.229

	pointsOfFunctionPlot[551].X = -1.46
	pointsOfFunctionPlot[551].Y = 0.232

	pointsOfFunctionPlot[552].X = -1.45
	pointsOfFunctionPlot[552].Y = 0.234

	pointsOfFunctionPlot[553].X = -1.44
	pointsOfFunctionPlot[553].Y = 0.236

	pointsOfFunctionPlot[554].X = -1.43
	pointsOfFunctionPlot[554].Y = 0.239

	pointsOfFunctionPlot[555].X = -1.42
	pointsOfFunctionPlot[555].Y = 0.241

	pointsOfFunctionPlot[556].X = -1.41
	pointsOfFunctionPlot[556].Y = 0.244

	pointsOfFunctionPlot[557].X = -1.4
	pointsOfFunctionPlot[557].Y = 0.246

	pointsOfFunctionPlot[558].X = -1.39
	pointsOfFunctionPlot[558].Y = 0.249

	pointsOfFunctionPlot[559].X = -1.38
	pointsOfFunctionPlot[559].Y = 0.251

	pointsOfFunctionPlot[560].X = -1.37
	pointsOfFunctionPlot[560].Y = 0.254

	pointsOfFunctionPlot[561].X = -1.36
	pointsOfFunctionPlot[561].Y = 0.256

	pointsOfFunctionPlot[562].X = -1.35
	pointsOfFunctionPlot[562].Y = 0.259

	pointsOfFunctionPlot[563].X = -1.34
	pointsOfFunctionPlot[563].Y = 0.261

	pointsOfFunctionPlot[564].X = -1.33
	pointsOfFunctionPlot[564].Y = 0.264

	pointsOfFunctionPlot[565].X = -1.32
	pointsOfFunctionPlot[565].Y = 0.267

	pointsOfFunctionPlot[566].X = -1.31
	pointsOfFunctionPlot[566].Y = 0.269

	pointsOfFunctionPlot[567].X = -1.3
	pointsOfFunctionPlot[567].Y = 0.272

	pointsOfFunctionPlot[568].X = -1.29
	pointsOfFunctionPlot[568].Y = 0.275

	pointsOfFunctionPlot[569].X = -1.28
	pointsOfFunctionPlot[569].Y = 0.278

	pointsOfFunctionPlot[570].X = -1.27
	pointsOfFunctionPlot[570].Y = 0.28

	pointsOfFunctionPlot[571].X = -1.26
	pointsOfFunctionPlot[571].Y = 0.283

	pointsOfFunctionPlot[572].X = -1.25
	pointsOfFunctionPlot[572].Y = 0.286

	pointsOfFunctionPlot[573].X = -1.24
	pointsOfFunctionPlot[573].Y = 0.289

	pointsOfFunctionPlot[574].X = -1.23
	pointsOfFunctionPlot[574].Y = 0.292

	pointsOfFunctionPlot[575].X = -1.22
	pointsOfFunctionPlot[575].Y = 0.295

	pointsOfFunctionPlot[576].X = -1.21
	pointsOfFunctionPlot[576].Y = 0.298

	pointsOfFunctionPlot[577].X = -1.2
	pointsOfFunctionPlot[577].Y = 0.301

	pointsOfFunctionPlot[578].X = -1.19
	pointsOfFunctionPlot[578].Y = 0.304

	pointsOfFunctionPlot[579].X = -1.18
	pointsOfFunctionPlot[579].Y = 0.307

	pointsOfFunctionPlot[580].X = -1.17
	pointsOfFunctionPlot[580].Y = 0.31

	pointsOfFunctionPlot[581].X = -1.16
	pointsOfFunctionPlot[581].Y = 0.313

	pointsOfFunctionPlot[582].X = -1.15
	pointsOfFunctionPlot[582].Y = 0.316

	pointsOfFunctionPlot[583].X = -1.14
	pointsOfFunctionPlot[583].Y = 0.319

	pointsOfFunctionPlot[584].X = -1.13
	pointsOfFunctionPlot[584].Y = 0.323

	pointsOfFunctionPlot[585].X = -1.12
	pointsOfFunctionPlot[585].Y = 0.326

	pointsOfFunctionPlot[586].X = -1.11
	pointsOfFunctionPlot[586].Y = 0.329

	pointsOfFunctionPlot[587].X = -1.1
	pointsOfFunctionPlot[587].Y = 0.332

	pointsOfFunctionPlot[588].X = -1.09
	pointsOfFunctionPlot[588].Y = 0.336

	pointsOfFunctionPlot[589].X = -1.08
	pointsOfFunctionPlot[589].Y = 0.339

	pointsOfFunctionPlot[590].X = -1.07
	pointsOfFunctionPlot[590].Y = 0.343

	pointsOfFunctionPlot[591].X = -1.06
	pointsOfFunctionPlot[591].Y = 0.346

	pointsOfFunctionPlot[592].X = -1.05
	pointsOfFunctionPlot[592].Y = 0.349

	pointsOfFunctionPlot[593].X = -1.04
	pointsOfFunctionPlot[593].Y = 0.353

	pointsOfFunctionPlot[594].X = -1.03
	pointsOfFunctionPlot[594].Y = 0.357

	pointsOfFunctionPlot[595].X = -1.02
	pointsOfFunctionPlot[595].Y = 0.36

	pointsOfFunctionPlot[596].X = -1.01
	pointsOfFunctionPlot[596].Y = 0.364

	pointsOfFunctionPlot[597].X = -1.0
	pointsOfFunctionPlot[597].Y = 0.367

	pointsOfFunctionPlot[598].X = -0.99
	pointsOfFunctionPlot[598].Y = 0.371

	pointsOfFunctionPlot[599].X = -0.98
	pointsOfFunctionPlot[599].Y = 0.375

	pointsOfFunctionPlot[600].X = -0.97
	pointsOfFunctionPlot[600].Y = 0.379

	pointsOfFunctionPlot[601].X = -0.96
	pointsOfFunctionPlot[601].Y = 0.382

	pointsOfFunctionPlot[602].X = -0.95
	pointsOfFunctionPlot[602].Y = 0.386

	pointsOfFunctionPlot[603].X = -0.94
	pointsOfFunctionPlot[603].Y = 0.39

	pointsOfFunctionPlot[604].X = -0.93
	pointsOfFunctionPlot[604].Y = 0.394

	pointsOfFunctionPlot[605].X = -0.92
	pointsOfFunctionPlot[605].Y = 0.398

	pointsOfFunctionPlot[606].X = -0.91
	pointsOfFunctionPlot[606].Y = 0.402

	pointsOfFunctionPlot[607].X = -0.9
	pointsOfFunctionPlot[607].Y = 0.406

	pointsOfFunctionPlot[608].X = -0.89
	pointsOfFunctionPlot[608].Y = 0.41

	pointsOfFunctionPlot[609].X = -0.88
	pointsOfFunctionPlot[609].Y = 0.414

	pointsOfFunctionPlot[610].X = -0.87
	pointsOfFunctionPlot[610].Y = 0.419

	pointsOfFunctionPlot[611].X = -0.86
	pointsOfFunctionPlot[611].Y = 0.423

	pointsOfFunctionPlot[612].X = -0.85
	pointsOfFunctionPlot[612].Y = 0.427

	pointsOfFunctionPlot[613].X = -0.84
	pointsOfFunctionPlot[613].Y = 0.431

	pointsOfFunctionPlot[614].X = -0.83
	pointsOfFunctionPlot[614].Y = 0.436

	pointsOfFunctionPlot[615].X = -0.82
	pointsOfFunctionPlot[615].Y = 0.44

	pointsOfFunctionPlot[616].X = -0.81
	pointsOfFunctionPlot[616].Y = 0.444

	pointsOfFunctionPlot[617].X = -0.8
	pointsOfFunctionPlot[617].Y = 0.449

	pointsOfFunctionPlot[618].X = -0.79
	pointsOfFunctionPlot[618].Y = 0.453

	pointsOfFunctionPlot[619].X = -0.78
	pointsOfFunctionPlot[619].Y = 0.458

	pointsOfFunctionPlot[620].X = -0.77
	pointsOfFunctionPlot[620].Y = 0.463

	pointsOfFunctionPlot[621].X = -0.76
	pointsOfFunctionPlot[621].Y = 0.467

	pointsOfFunctionPlot[622].X = -0.75
	pointsOfFunctionPlot[622].Y = 0.472

	pointsOfFunctionPlot[623].X = -0.74
	pointsOfFunctionPlot[623].Y = 0.477

	pointsOfFunctionPlot[624].X = -0.73
	pointsOfFunctionPlot[624].Y = 0.481

	pointsOfFunctionPlot[625].X = -0.72
	pointsOfFunctionPlot[625].Y = 0.486

	pointsOfFunctionPlot[626].X = -0.71
	pointsOfFunctionPlot[626].Y = 0.491

	pointsOfFunctionPlot[627].X = -0.7
	pointsOfFunctionPlot[627].Y = 0.496

	pointsOfFunctionPlot[628].X = -0.69
	pointsOfFunctionPlot[628].Y = 0.501

	pointsOfFunctionPlot[629].X = -0.68
	pointsOfFunctionPlot[629].Y = 0.506

	pointsOfFunctionPlot[630].X = -0.67
	pointsOfFunctionPlot[630].Y = 0.511

	pointsOfFunctionPlot[631].X = -0.66
	pointsOfFunctionPlot[631].Y = 0.516

	pointsOfFunctionPlot[632].X = -0.65
	pointsOfFunctionPlot[632].Y = 0.522

	pointsOfFunctionPlot[633].X = -0.64
	pointsOfFunctionPlot[633].Y = 0.527

	pointsOfFunctionPlot[634].X = -0.63
	pointsOfFunctionPlot[634].Y = 0.532

	pointsOfFunctionPlot[635].X = -0.62
	pointsOfFunctionPlot[635].Y = 0.537

	pointsOfFunctionPlot[636].X = -0.61
	pointsOfFunctionPlot[636].Y = 0.543

	pointsOfFunctionPlot[637].X = -0.6
	pointsOfFunctionPlot[637].Y = 0.548

	pointsOfFunctionPlot[638].X = -0.59
	pointsOfFunctionPlot[638].Y = 0.554

	pointsOfFunctionPlot[639].X = -0.58
	pointsOfFunctionPlot[639].Y = 0.559

	pointsOfFunctionPlot[640].X = -0.57
	pointsOfFunctionPlot[640].Y = 0.565

	pointsOfFunctionPlot[641].X = -0.56
	pointsOfFunctionPlot[641].Y = 0.571

	pointsOfFunctionPlot[642].X = -0.55
	pointsOfFunctionPlot[642].Y = 0.576

	pointsOfFunctionPlot[643].X = -0.54
	pointsOfFunctionPlot[643].Y = 0.582

	pointsOfFunctionPlot[644].X = -0.53
	pointsOfFunctionPlot[644].Y = 0.588

	pointsOfFunctionPlot[645].X = -0.52
	pointsOfFunctionPlot[645].Y = 0.594

	pointsOfFunctionPlot[646].X = -0.51
	pointsOfFunctionPlot[646].Y = 0.6

	pointsOfFunctionPlot[647].X = -0.5
	pointsOfFunctionPlot[647].Y = 0.606

	pointsOfFunctionPlot[648].X = -0.49
	pointsOfFunctionPlot[648].Y = 0.612

	pointsOfFunctionPlot[649].X = -0.48
	pointsOfFunctionPlot[649].Y = 0.618

	pointsOfFunctionPlot[650].X = -0.47
	pointsOfFunctionPlot[650].Y = 0.625

	pointsOfFunctionPlot[651].X = -0.46
	pointsOfFunctionPlot[651].Y = 0.631

	pointsOfFunctionPlot[652].X = -0.45
	pointsOfFunctionPlot[652].Y = 0.637

	pointsOfFunctionPlot[653].X = -0.44
	pointsOfFunctionPlot[653].Y = 0.644

	pointsOfFunctionPlot[654].X = -0.43
	pointsOfFunctionPlot[654].Y = 0.65

	pointsOfFunctionPlot[655].X = -0.42
	pointsOfFunctionPlot[655].Y = 0.657

	pointsOfFunctionPlot[656].X = -0.41
	pointsOfFunctionPlot[656].Y = 0.663

	pointsOfFunctionPlot[657].X = -0.4
	pointsOfFunctionPlot[657].Y = 0.67

	pointsOfFunctionPlot[658].X = -0.39
	pointsOfFunctionPlot[658].Y = 0.677

	pointsOfFunctionPlot[659].X = -0.38
	pointsOfFunctionPlot[659].Y = 0.683

	pointsOfFunctionPlot[660].X = -0.37
	pointsOfFunctionPlot[660].Y = 0.69

	pointsOfFunctionPlot[661].X = -0.36
	pointsOfFunctionPlot[661].Y = 0.697

	pointsOfFunctionPlot[662].X = -0.35
	pointsOfFunctionPlot[662].Y = 0.704

	pointsOfFunctionPlot[663].X = -0.34
	pointsOfFunctionPlot[663].Y = 0.711

	pointsOfFunctionPlot[664].X = -0.33
	pointsOfFunctionPlot[664].Y = 0.718

	pointsOfFunctionPlot[665].X = -0.32
	pointsOfFunctionPlot[665].Y = 0.726

	pointsOfFunctionPlot[666].X = -0.31
	pointsOfFunctionPlot[666].Y = 0.733

	pointsOfFunctionPlot[667].X = -0.3
	pointsOfFunctionPlot[667].Y = 0.74

	pointsOfFunctionPlot[668].X = -0.29
	pointsOfFunctionPlot[668].Y = 0.748

	pointsOfFunctionPlot[669].X = -0.28
	pointsOfFunctionPlot[669].Y = 0.755

	pointsOfFunctionPlot[670].X = -0.27
	pointsOfFunctionPlot[670].Y = 0.763

	pointsOfFunctionPlot[671].X = -0.26
	pointsOfFunctionPlot[671].Y = 0.771

	pointsOfFunctionPlot[672].X = -0.25
	pointsOfFunctionPlot[672].Y = 0.778

	pointsOfFunctionPlot[673].X = -0.24
	pointsOfFunctionPlot[673].Y = 0.786

	pointsOfFunctionPlot[674].X = -0.23
	pointsOfFunctionPlot[674].Y = 0.794

	pointsOfFunctionPlot[675].X = -0.22
	pointsOfFunctionPlot[675].Y = 0.802

	pointsOfFunctionPlot[676].X = -0.21
	pointsOfFunctionPlot[676].Y = 0.81

	pointsOfFunctionPlot[677].X = -0.2
	pointsOfFunctionPlot[677].Y = 0.818

	pointsOfFunctionPlot[678].X = -0.19
	pointsOfFunctionPlot[678].Y = 0.827

	pointsOfFunctionPlot[679].X = -0.18
	pointsOfFunctionPlot[679].Y = 0.835

	pointsOfFunctionPlot[680].X = -0.17
	pointsOfFunctionPlot[680].Y = 0.843

	pointsOfFunctionPlot[681].X = -0.16
	pointsOfFunctionPlot[681].Y = 0.852

	pointsOfFunctionPlot[682].X = -0.15
	pointsOfFunctionPlot[682].Y = 0.86

	pointsOfFunctionPlot[683].X = -0.14
	pointsOfFunctionPlot[683].Y = 0.869

	pointsOfFunctionPlot[684].X = -0.13
	pointsOfFunctionPlot[684].Y = 0.878

	pointsOfFunctionPlot[685].X = -0.12
	pointsOfFunctionPlot[685].Y = 0.886

	pointsOfFunctionPlot[686].X = -0.11
	pointsOfFunctionPlot[686].Y = 0.895

	pointsOfFunctionPlot[687].X = -0.1
	pointsOfFunctionPlot[687].Y = 0.904

	pointsOfFunctionPlot[688].X = -0.09
	pointsOfFunctionPlot[688].Y = 0.913

	pointsOfFunctionPlot[689].X = -0.08
	pointsOfFunctionPlot[689].Y = 0.923

	pointsOfFunctionPlot[690].X = -0.07
	pointsOfFunctionPlot[690].Y = 0.932

	pointsOfFunctionPlot[691].X = -0.06
	pointsOfFunctionPlot[691].Y = 0.941

	pointsOfFunctionPlot[692].X = -0.05
	pointsOfFunctionPlot[692].Y = 0.951

	pointsOfFunctionPlot[693].X = -0.04
	pointsOfFunctionPlot[693].Y = 0.96

	pointsOfFunctionPlot[694].X = -0.03
	pointsOfFunctionPlot[694].Y = 0.97

	pointsOfFunctionPlot[695].X = -0.02
	pointsOfFunctionPlot[695].Y = 0.98

	pointsOfFunctionPlot[696].X = -0.01
	pointsOfFunctionPlot[696].Y = 0.99

	pointsOfFunctionPlot[697].X = 0.0
	pointsOfFunctionPlot[697].Y = 1.0

	pointsOfFunctionPlot[698].X = 0.01
	pointsOfFunctionPlot[698].Y = 1.01

	pointsOfFunctionPlot[699].X = 0.02
	pointsOfFunctionPlot[699].Y = 1.02

	pointsOfFunctionPlot[700].X = 0.03
	pointsOfFunctionPlot[700].Y = 1.03

	pointsOfFunctionPlot[701].X = 0.04
	pointsOfFunctionPlot[701].Y = 1.04

	pointsOfFunctionPlot[702].X = 0.05
	pointsOfFunctionPlot[702].Y = 1.051

	pointsOfFunctionPlot[703].X = 0.06
	pointsOfFunctionPlot[703].Y = 1.061

	pointsOfFunctionPlot[704].X = 0.07
	pointsOfFunctionPlot[704].Y = 1.072

	pointsOfFunctionPlot[705].X = 0.08
	pointsOfFunctionPlot[705].Y = 1.083

	pointsOfFunctionPlot[706].X = 0.09
	pointsOfFunctionPlot[706].Y = 1.094

	pointsOfFunctionPlot[707].X = 0.1
	pointsOfFunctionPlot[707].Y = 1.105

	pointsOfFunctionPlot[708].X = 0.11
	pointsOfFunctionPlot[708].Y = 1.116

	pointsOfFunctionPlot[709].X = 0.12
	pointsOfFunctionPlot[709].Y = 1.127

	pointsOfFunctionPlot[710].X = 0.13
	pointsOfFunctionPlot[710].Y = 1.138

	pointsOfFunctionPlot[711].X = 0.14
	pointsOfFunctionPlot[711].Y = 1.15

	pointsOfFunctionPlot[712].X = 0.15
	pointsOfFunctionPlot[712].Y = 1.161

	pointsOfFunctionPlot[713].X = 0.16
	pointsOfFunctionPlot[713].Y = 1.173

	pointsOfFunctionPlot[714].X = 0.17
	pointsOfFunctionPlot[714].Y = 1.185

	pointsOfFunctionPlot[715].X = 0.18
	pointsOfFunctionPlot[715].Y = 1.197

	pointsOfFunctionPlot[716].X = 0.19
	pointsOfFunctionPlot[716].Y = 1.209

	pointsOfFunctionPlot[717].X = 0.2
	pointsOfFunctionPlot[717].Y = 1.221

	pointsOfFunctionPlot[718].X = 0.21
	pointsOfFunctionPlot[718].Y = 1.233

	pointsOfFunctionPlot[719].X = 0.22
	pointsOfFunctionPlot[719].Y = 1.246

	pointsOfFunctionPlot[720].X = 0.23
	pointsOfFunctionPlot[720].Y = 1.258

	pointsOfFunctionPlot[721].X = 0.24
	pointsOfFunctionPlot[721].Y = 1.271

	pointsOfFunctionPlot[722].X = 0.25
	pointsOfFunctionPlot[722].Y = 1.284

	pointsOfFunctionPlot[723].X = 0.26
	pointsOfFunctionPlot[723].Y = 1.296

	pointsOfFunctionPlot[724].X = 0.27
	pointsOfFunctionPlot[724].Y = 1.31

	pointsOfFunctionPlot[725].X = 0.28
	pointsOfFunctionPlot[725].Y = 1.323

	pointsOfFunctionPlot[726].X = 0.29
	pointsOfFunctionPlot[726].Y = 1.336

	pointsOfFunctionPlot[727].X = 0.3
	pointsOfFunctionPlot[727].Y = 1.349

	pointsOfFunctionPlot[728].X = 0.31
	pointsOfFunctionPlot[728].Y = 1.363

	pointsOfFunctionPlot[729].X = 0.32
	pointsOfFunctionPlot[729].Y = 1.377

	pointsOfFunctionPlot[730].X = 0.33
	pointsOfFunctionPlot[730].Y = 1.391

	pointsOfFunctionPlot[731].X = 0.34
	pointsOfFunctionPlot[731].Y = 1.404

	pointsOfFunctionPlot[732].X = 0.35
	pointsOfFunctionPlot[732].Y = 1.419

	pointsOfFunctionPlot[733].X = 0.36
	pointsOfFunctionPlot[733].Y = 1.433

	pointsOfFunctionPlot[734].X = 0.37
	pointsOfFunctionPlot[734].Y = 1.447

	pointsOfFunctionPlot[735].X = 0.38
	pointsOfFunctionPlot[735].Y = 1.462

	pointsOfFunctionPlot[736].X = 0.39
	pointsOfFunctionPlot[736].Y = 1.477

	pointsOfFunctionPlot[737].X = 0.40
	pointsOfFunctionPlot[737].Y = 1.491

	pointsOfFunctionPlot[738].X = 0.41
	pointsOfFunctionPlot[738].Y = 1.506

	pointsOfFunctionPlot[739].X = 0.42
	pointsOfFunctionPlot[739].Y = 1.522

	pointsOfFunctionPlot[740].X = 0.43
	pointsOfFunctionPlot[740].Y = 1.537

	pointsOfFunctionPlot[741].X = 0.44
	pointsOfFunctionPlot[741].Y = 1.552

	pointsOfFunctionPlot[742].X = 0.45
	pointsOfFunctionPlot[742].Y = 1.568

	pointsOfFunctionPlot[743].X = 0.46
	pointsOfFunctionPlot[743].Y = 1.584

	pointsOfFunctionPlot[744].X = 0.47
	pointsOfFunctionPlot[744].Y = 1.6

	pointsOfFunctionPlot[745].X = 0.48
	pointsOfFunctionPlot[745].Y = 1.616

	pointsOfFunctionPlot[746].X = 0.49
	pointsOfFunctionPlot[746].Y = 1.632

	pointsOfFunctionPlot[747].X = 0.5
	pointsOfFunctionPlot[747].Y = 1.648

	pointsOfFunctionPlot[748].X = 0.51
	pointsOfFunctionPlot[748].Y = 1.665

	pointsOfFunctionPlot[749].X = 0.52
	pointsOfFunctionPlot[749].Y = 1.682

	pointsOfFunctionPlot[750].X = 0.53
	pointsOfFunctionPlot[750].Y = 1.698

	pointsOfFunctionPlot[751].X = 0.54
	pointsOfFunctionPlot[751].Y = 1.716

	pointsOfFunctionPlot[752].X = 0.55
	pointsOfFunctionPlot[752].Y = 1.733

	pointsOfFunctionPlot[753].X = 0.56
	pointsOfFunctionPlot[753].Y = 1.75

	pointsOfFunctionPlot[754].X = 0.57
	pointsOfFunctionPlot[754].Y = 1.768

	pointsOfFunctionPlot[755].X = 0.58
	pointsOfFunctionPlot[755].Y = 1.786

	pointsOfFunctionPlot[756].X = 0.59
	pointsOfFunctionPlot[756].Y = 1.804

	pointsOfFunctionPlot[757].X = 0.6
	pointsOfFunctionPlot[757].Y = 1.822

	pointsOfFunctionPlot[758].X = 0.61
	pointsOfFunctionPlot[758].Y = 1.84

	pointsOfFunctionPlot[759].X = 0.62
	pointsOfFunctionPlot[759].Y = 1.858

	pointsOfFunctionPlot[760].X = 0.63
	pointsOfFunctionPlot[760].Y = 1.877

	pointsOfFunctionPlot[761].X = 0.64
	pointsOfFunctionPlot[761].Y = 1.896

	pointsOfFunctionPlot[762].X = 0.65
	pointsOfFunctionPlot[762].Y = 1.915

	pointsOfFunctionPlot[763].X = 0.66
	pointsOfFunctionPlot[763].Y = 1.938

	pointsOfFunctionPlot[764].X = 0.67
	pointsOfFunctionPlot[764].Y = 1.954

	pointsOfFunctionPlot[765].X = 0.68
	pointsOfFunctionPlot[765].Y = 1.973

	pointsOfFunctionPlot[766].X = 0.69
	pointsOfFunctionPlot[766].Y = 1.993

	pointsOfFunctionPlot[767].X = 0.7
	pointsOfFunctionPlot[767].Y = 2.013

	pointsOfFunctionPlot[768].X = 0.71
	pointsOfFunctionPlot[768].Y = 2.034

	pointsOfFunctionPlot[769].X = 0.72
	pointsOfFunctionPlot[769].Y = 2.054

	pointsOfFunctionPlot[770].X = 0.73
	pointsOfFunctionPlot[770].Y = 2.075

	pointsOfFunctionPlot[771].X = 0.74
	pointsOfFunctionPlot[771].Y = 2.095

	pointsOfFunctionPlot[772].X = 0.75
	pointsOfFunctionPlot[772].Y = 2.117

	pointsOfFunctionPlot[773].X = 0.76
	pointsOfFunctionPlot[773].Y = 2.138

	pointsOfFunctionPlot[774].X = 0.77
	pointsOfFunctionPlot[774].Y = 2.159

	pointsOfFunctionPlot[775].X = 0.78
	pointsOfFunctionPlot[775].Y = 2.181

	pointsOfFunctionPlot[776].X = 0.79
	pointsOfFunctionPlot[776].Y = 2.203

	pointsOfFunctionPlot[777].X = 0.8
	pointsOfFunctionPlot[777].Y = 2.225

	pointsOfFunctionPlot[778].X = 0.81
	pointsOfFunctionPlot[778].Y = 2.247

	pointsOfFunctionPlot[779].X = 0.82
	pointsOfFunctionPlot[779].Y = 2.27

	pointsOfFunctionPlot[780].X = 0.83
	pointsOfFunctionPlot[780].Y = 2.293

	pointsOfFunctionPlot[781].X = 0.84
	pointsOfFunctionPlot[781].Y = 2.316

	pointsOfFunctionPlot[782].X = 0.85
	pointsOfFunctionPlot[782].Y = 2.339

	pointsOfFunctionPlot[783].X = 0.86
	pointsOfFunctionPlot[783].Y = 2.363

	pointsOfFunctionPlot[784].X = 0.87
	pointsOfFunctionPlot[784].Y = 2.386

	pointsOfFunctionPlot[785].X = 0.88
	pointsOfFunctionPlot[785].Y = 2.41

	pointsOfFunctionPlot[786].X = 0.89
	pointsOfFunctionPlot[786].Y = 2.435

	pointsOfFunctionPlot[787].X = 0.9
	pointsOfFunctionPlot[787].Y = 2.459

	pointsOfFunctionPlot[788].X = 0.91
	pointsOfFunctionPlot[788].Y = 2.484

	pointsOfFunctionPlot[789].X = 0.92
	pointsOfFunctionPlot[789].Y = 2.509

	pointsOfFunctionPlot[790].X = 0.93
	pointsOfFunctionPlot[790].Y = 2.534

	pointsOfFunctionPlot[791].X = 0.94
	pointsOfFunctionPlot[791].Y = 2.56

	pointsOfFunctionPlot[792].X = 0.95
	pointsOfFunctionPlot[792].Y = 2.585

	pointsOfFunctionPlot[793].X = 0.96
	pointsOfFunctionPlot[793].Y = 2.611

	pointsOfFunctionPlot[794].X = 0.97
	pointsOfFunctionPlot[794].Y = 2.637

	pointsOfFunctionPlot[795].X = 0.98
	pointsOfFunctionPlot[795].Y = 2.664

	pointsOfFunctionPlot[796].X = 0.99
	pointsOfFunctionPlot[796].Y = 2.691

	pointsOfFunctionPlot[797].X = 1.0
	pointsOfFunctionPlot[797].Y = 2.718

	pointsOfFunctionPlot[798].X = 1.01
	pointsOfFunctionPlot[798].Y = 2.745

	pointsOfFunctionPlot[799].X = 1.02
	pointsOfFunctionPlot[799].Y = 2.773

	pointsOfFunctionPlot[800].X = 1.03
	pointsOfFunctionPlot[800].Y = 2.801

	pointsOfFunctionPlot[801].X = 1.04
	pointsOfFunctionPlot[801].Y = 2.829

	pointsOfFunctionPlot[802].X = 1.05
	pointsOfFunctionPlot[802].Y = 2.857

	pointsOfFunctionPlot[803].X = 1.06
	pointsOfFunctionPlot[803].Y = 2.886

	pointsOfFunctionPlot[804].X = 1.07
	pointsOfFunctionPlot[804].Y = 2.915

	pointsOfFunctionPlot[805].X = 1.08
	pointsOfFunctionPlot[805].Y = 2.944

	pointsOfFunctionPlot[806].X = 1.09
	pointsOfFunctionPlot[806].Y = 2.974

	pointsOfFunctionPlot[807].X = 1.1
	pointsOfFunctionPlot[807].Y = 3.004

	pointsOfFunctionPlot[808].X = 1.11
	pointsOfFunctionPlot[808].Y = 3.034

	pointsOfFunctionPlot[809].X = 1.12
	pointsOfFunctionPlot[809].Y = 3.064

	pointsOfFunctionPlot[810].X = 1.13
	pointsOfFunctionPlot[810].Y = 3.095

	pointsOfFunctionPlot[811].X = 1.14
	pointsOfFunctionPlot[811].Y = 3.126

	pointsOfFunctionPlot[812].X = 1.15
	pointsOfFunctionPlot[812].Y = 3.158

	pointsOfFunctionPlot[813].X = 1.16
	pointsOfFunctionPlot[813].Y = 3.189

	pointsOfFunctionPlot[814].X = 1.17
	pointsOfFunctionPlot[814].Y = 3.222

	pointsOfFunctionPlot[815].X = 1.18
	pointsOfFunctionPlot[815].Y = 3.254

	pointsOfFunctionPlot[816].X = 1.19
	pointsOfFunctionPlot[816].Y = 3.287

	pointsOfFunctionPlot[817].X = 1.2
	pointsOfFunctionPlot[817].Y = 3.32

	pointsOfFunctionPlot[818].X = 1.21
	pointsOfFunctionPlot[818].Y = 3.353

	pointsOfFunctionPlot[819].X = 1.22
	pointsOfFunctionPlot[819].Y = 3.387

	pointsOfFunctionPlot[820].X = 1.23
	pointsOfFunctionPlot[820].Y = 3.421

	pointsOfFunctionPlot[821].X = 1.24
	pointsOfFunctionPlot[821].Y = 3.455

	pointsOfFunctionPlot[822].X = 1.25
	pointsOfFunctionPlot[822].Y = 3.49

	pointsOfFunctionPlot[823].X = 1.26
	pointsOfFunctionPlot[823].Y = 3.525

	pointsOfFunctionPlot[824].X = 1.27
	pointsOfFunctionPlot[824].Y = 3.56

	pointsOfFunctionPlot[825].X = 1.28
	pointsOfFunctionPlot[825].Y = 3.596

	pointsOfFunctionPlot[826].X = 1.29
	pointsOfFunctionPlot[826].Y = 3.632

	pointsOfFunctionPlot[827].X = 1.3
	pointsOfFunctionPlot[827].Y = 3.669

	pointsOfFunctionPlot[828].X = 1.31
	pointsOfFunctionPlot[828].Y = 3.706

	pointsOfFunctionPlot[829].X = 1.32
	pointsOfFunctionPlot[829].Y = 3.743

	pointsOfFunctionPlot[830].X = 1.33
	pointsOfFunctionPlot[830].Y = 3.781

	pointsOfFunctionPlot[831].X = 1.34
	pointsOfFunctionPlot[831].Y = 3.819

	pointsOfFunctionPlot[832].X = 1.35
	pointsOfFunctionPlot[832].Y = 3.857

	pointsOfFunctionPlot[833].X = 1.36
	pointsOfFunctionPlot[833].Y = 3.896

	pointsOfFunctionPlot[834].X = 1.37
	pointsOfFunctionPlot[834].Y = 3.935

	pointsOfFunctionPlot[835].X = 1.38
	pointsOfFunctionPlot[835].Y = 3.974

	pointsOfFunctionPlot[836].X = 1.39
	pointsOfFunctionPlot[836].Y = 4.014

	pointsOfFunctionPlot[837].X = 1.40
	pointsOfFunctionPlot[837].Y = 4.055

	pointsOfFunctionPlot[838].X = 1.41
	pointsOfFunctionPlot[838].Y = 4.096

	pointsOfFunctionPlot[839].X = 1.42
	pointsOfFunctionPlot[839].Y = 4.137

	pointsOfFunctionPlot[840].X = 1.43
	pointsOfFunctionPlot[840].Y = 4.178

	pointsOfFunctionPlot[841].X = 1.44
	pointsOfFunctionPlot[841].Y = 4.22

	pointsOfFunctionPlot[842].X = 1.45
	pointsOfFunctionPlot[842].Y = 4.263

	pointsOfFunctionPlot[843].X = 1.46
	pointsOfFunctionPlot[843].Y = 4.306

	pointsOfFunctionPlot[844].X = 1.47
	pointsOfFunctionPlot[844].Y = 4.349

	pointsOfFunctionPlot[845].X = 1.48
	pointsOfFunctionPlot[845].Y = 4.392

	pointsOfFunctionPlot[846].X = 1.49
	pointsOfFunctionPlot[846].Y = 4.437

	pointsOfFunctionPlot[847].X = 1.5
	pointsOfFunctionPlot[847].Y = 4.481

	pointsOfFunctionPlot[848].X = 1.51
	pointsOfFunctionPlot[848].Y = 4.526

	pointsOfFunctionPlot[849].X = 1.52
	pointsOfFunctionPlot[849].Y = 4.572

	pointsOfFunctionPlot[850].X = 1.53
	pointsOfFunctionPlot[850].Y = 4.618

	pointsOfFunctionPlot[851].X = 1.54
	pointsOfFunctionPlot[851].Y = 4.664

	pointsOfFunctionPlot[852].X = 1.55
	pointsOfFunctionPlot[852].Y = 4.711

	pointsOfFunctionPlot[853].X = 1.56
	pointsOfFunctionPlot[853].Y = 4.758

	pointsOfFunctionPlot[854].X = 1.57
	pointsOfFunctionPlot[854].Y = 4.806

	pointsOfFunctionPlot[855].X = 1.58
	pointsOfFunctionPlot[855].Y = 4.855

	pointsOfFunctionPlot[856].X = 1.59
	pointsOfFunctionPlot[856].Y = 4.903

	pointsOfFunctionPlot[857].X = 1.6
	pointsOfFunctionPlot[857].Y = 4.953

	pointsOfFunctionPlot[858].X = 1.61
	pointsOfFunctionPlot[858].Y = 5.002

	pointsOfFunctionPlot[859].X = 1.62
	pointsOfFunctionPlot[859].Y = 5.053

	pointsOfFunctionPlot[860].X = 1.63
	pointsOfFunctionPlot[860].Y = 5.103

	pointsOfFunctionPlot[861].X = 1.64
	pointsOfFunctionPlot[861].Y = 5.155

	pointsOfFunctionPlot[862].X = 1.65
	pointsOfFunctionPlot[862].Y = 5.207

	pointsOfFunctionPlot[863].X = 1.66
	pointsOfFunctionPlot[863].Y = 5.259

	pointsOfFunctionPlot[864].X = 1.67
	pointsOfFunctionPlot[864].Y = 5.312

	pointsOfFunctionPlot[865].X = 1.68
	pointsOfFunctionPlot[865].Y = 5.365

	pointsOfFunctionPlot[866].X = 1.69
	pointsOfFunctionPlot[866].Y = 5.419

	pointsOfFunctionPlot[867].X = 1.7
	pointsOfFunctionPlot[867].Y = 5.473

	pointsOfFunctionPlot[868].X = 1.71
	pointsOfFunctionPlot[868].Y = 5.529

	pointsOfFunctionPlot[869].X = 1.72
	pointsOfFunctionPlot[869].Y = 5.584

	pointsOfFunctionPlot[870].X = 1.73
	pointsOfFunctionPlot[870].Y = 5.64

	pointsOfFunctionPlot[871].X = 1.74
	pointsOfFunctionPlot[871].Y = 5.697

	pointsOfFunctionPlot[872].X = 1.75
	pointsOfFunctionPlot[872].Y = 5.754

	pointsOfFunctionPlot[873].X = 1.76
	pointsOfFunctionPlot[873].Y = 5.812

	pointsOfFunctionPlot[874].X = 1.77
	pointsOfFunctionPlot[874].Y = 5.87

	pointsOfFunctionPlot[875].X = 1.78
	pointsOfFunctionPlot[875].Y = 5.929

	pointsOfFunctionPlot[876].X = 1.79
	pointsOfFunctionPlot[876].Y = 5.989

	pointsOfFunctionPlot[877].X = 1.8
	pointsOfFunctionPlot[877].Y = 6.049

	pointsOfFunctionPlot[878].X = 1.81
	pointsOfFunctionPlot[878].Y = 6.11

	pointsOfFunctionPlot[879].X = 1.82
	pointsOfFunctionPlot[879].Y = 6.171

	pointsOfFunctionPlot[880].X = 1.83
	pointsOfFunctionPlot[880].Y = 6.233

	pointsOfFunctionPlot[881].X = 1.84
	pointsOfFunctionPlot[881].Y = 6.296

	pointsOfFunctionPlot[882].X = 1.85
	pointsOfFunctionPlot[882].Y = 6.359

	pointsOfFunctionPlot[883].X = 1.86
	pointsOfFunctionPlot[883].Y = 6.423

	pointsOfFunctionPlot[884].X = 1.87
	pointsOfFunctionPlot[884].Y = 6.488

	pointsOfFunctionPlot[885].X = 1.88
	pointsOfFunctionPlot[885].Y = 6.553

	pointsOfFunctionPlot[886].X = 1.89
	pointsOfFunctionPlot[886].Y = 6.619

	pointsOfFunctionPlot[887].X = 1.9
	pointsOfFunctionPlot[887].Y = 6.685

	pointsOfFunctionPlot[888].X = 1.91
	pointsOfFunctionPlot[888].Y = 6.753

	pointsOfFunctionPlot[889].X = 1.92
	pointsOfFunctionPlot[889].Y = 6.821

	pointsOfFunctionPlot[890].X = 1.93
	pointsOfFunctionPlot[890].Y = 6.889

	pointsOfFunctionPlot[891].X = 1.94
	pointsOfFunctionPlot[891].Y = 6.958

	pointsOfFunctionPlot[892].X = 1.95
	pointsOfFunctionPlot[892].Y = 7.028

	pointsOfFunctionPlot[893].X = 1.96
	pointsOfFunctionPlot[893].Y = 7.099

	pointsOfFunctionPlot[894].X = 1.97
	pointsOfFunctionPlot[894].Y = 7.17

	pointsOfFunctionPlot[895].X = 1.98
	pointsOfFunctionPlot[895].Y = 7.242

	pointsOfFunctionPlot[896].X = 1.99
	pointsOfFunctionPlot[896].Y = 7.315

	pointsOfFunctionPlot[897].X = 2.0
	pointsOfFunctionPlot[897].Y = 7.389

	pointsOfFunctionPlot[898].X = 2.01
	pointsOfFunctionPlot[898].Y = 7.463

	pointsOfFunctionPlot[899].X = 2.02
	pointsOfFunctionPlot[899].Y = 7.538

	pointsOfFunctionPlot[900].X = 2.03
	pointsOfFunctionPlot[900].Y = 7.614

	pointsOfFunctionPlot[901].X = 2.04
	pointsOfFunctionPlot[901].Y = 7.69

	pointsOfFunctionPlot[902].X = 2.05
	pointsOfFunctionPlot[902].Y = 7.767

	pointsOfFunctionPlot[903].X = 2.06
	pointsOfFunctionPlot[903].Y = 7.846

	pointsOfFunctionPlot[904].X = 2.07
	pointsOfFunctionPlot[904].Y = 7.924

	pointsOfFunctionPlot[905].X = 2.08
	pointsOfFunctionPlot[905].Y = 8.004

	pointsOfFunctionPlot[906].X = 2.09
	pointsOfFunctionPlot[906].Y = 8.084

	pointsOfFunctionPlot[907].X = 2.1
	pointsOfFunctionPlot[907].Y = 8.166

	pointsOfFunctionPlot[908].X = 2.11
	pointsOfFunctionPlot[908].Y = 8.248

	pointsOfFunctionPlot[909].X = 2.12
	pointsOfFunctionPlot[909].Y = 8.331

	pointsOfFunctionPlot[910].X = 2.13
	pointsOfFunctionPlot[910].Y = 8.414

	pointsOfFunctionPlot[911].X = 2.14
	pointsOfFunctionPlot[911].Y = 8.499

	pointsOfFunctionPlot[912].X = 2.15
	pointsOfFunctionPlot[912].Y = 8.584

	pointsOfFunctionPlot[913].X = 2.16
	pointsOfFunctionPlot[913].Y = 8.671

	pointsOfFunctionPlot[914].X = 2.17
	pointsOfFunctionPlot[914].Y = 8.758

	pointsOfFunctionPlot[915].X = 2.18
	pointsOfFunctionPlot[915].Y = 8.846

	pointsOfFunctionPlot[916].X = 2.19
	pointsOfFunctionPlot[916].Y = 8.935

	pointsOfFunctionPlot[917].X = 2.2
	pointsOfFunctionPlot[917].Y = 9.025

	pointsOfFunctionPlot[918].X = 2.21
	pointsOfFunctionPlot[918].Y = 9.115

	pointsOfFunctionPlot[919].X = 2.22
	pointsOfFunctionPlot[919].Y = 9.207

	pointsOfFunctionPlot[920].X = 2.23
	pointsOfFunctionPlot[920].Y = 9.299

	pointsOfFunctionPlot[921].X = 2.24
	pointsOfFunctionPlot[921].Y = 9.393

	pointsOfFunctionPlot[922].X = 2.25
	pointsOfFunctionPlot[922].Y = 9.487

	pointsOfFunctionPlot[923].X = 2.26
	pointsOfFunctionPlot[923].Y = 9.583

	pointsOfFunctionPlot[924].X = 2.27
	pointsOfFunctionPlot[924].Y = 9.679

	pointsOfFunctionPlot[925].X = 2.28
	pointsOfFunctionPlot[925].Y = 9.776

	pointsOfFunctionPlot[926].X = 2.29
	pointsOfFunctionPlot[926].Y = 9.874

	pointsOfFunctionPlot[927].X = 2.3
	pointsOfFunctionPlot[927].Y = 9.974

	pointsOfFunctionPlot[928].X = 2.31
	pointsOfFunctionPlot[928].Y = 10.074

	pointsOfFunctionPlot[929].X = 2.32
	pointsOfFunctionPlot[929].Y = 10.175

	pointsOfFunctionPlot[930].X = 2.33
	pointsOfFunctionPlot[930].Y = 10.277

	pointsOfFunctionPlot[931].X = 2.34
	pointsOfFunctionPlot[931].Y = 10.381

	pointsOfFunctionPlot[932].X = 2.35
	pointsOfFunctionPlot[932].Y = 10.485

	pointsOfFunctionPlot[933].X = 2.36
	pointsOfFunctionPlot[933].Y = 10.591

	pointsOfFunctionPlot[934].X = 2.37
	pointsOfFunctionPlot[934].Y = 10.697

	pointsOfFunctionPlot[935].X = 2.38
	pointsOfFunctionPlot[935].Y = 10.804

	pointsOfFunctionPlot[936].X = 2.39
	pointsOfFunctionPlot[936].Y = 10.913

	pointsOfFunctionPlot[937].X = 2.4
	pointsOfFunctionPlot[937].Y = 11.023

	pointsOfFunctionPlot[938].X = 2.41
	pointsOfFunctionPlot[938].Y = 11.134

	pointsOfFunctionPlot[939].X = 2.42
	pointsOfFunctionPlot[939].Y = 11.245

	pointsOfFunctionPlot[940].X = 2.43
	pointsOfFunctionPlot[940].Y = 11.358

	pointsOfFunctionPlot[941].X = 2.44
	pointsOfFunctionPlot[941].Y = 11.473

	pointsOfFunctionPlot[942].X = 2.45
	pointsOfFunctionPlot[942].Y = 11.588

	pointsOfFunctionPlot[943].X = 2.46
	pointsOfFunctionPlot[943].Y = 11.704

	pointsOfFunctionPlot[944].X = 2.47
	pointsOfFunctionPlot[944].Y = 11.822

	pointsOfFunctionPlot[945].X = 2.48
	pointsOfFunctionPlot[945].Y = 11.941

	pointsOfFunctionPlot[946].X = 2.49
	pointsOfFunctionPlot[946].Y = 12.061

	pointsOfFunctionPlot[947].X = 2.5
	pointsOfFunctionPlot[947].Y = 12.182

	pointsOfFunctionPlot[948].X = 2.51
	pointsOfFunctionPlot[948].Y = 12.304

	pointsOfFunctionPlot[949].X = 2.52
	pointsOfFunctionPlot[949].Y = 12.428

	pointsOfFunctionPlot[950].X = 2.53
	pointsOfFunctionPlot[950].Y = 12.553

	pointsOfFunctionPlot[951].X = 2.54
	pointsOfFunctionPlot[951].Y = 12.679

	pointsOfFunctionPlot[952].X = 2.55
	pointsOfFunctionPlot[952].Y = 12.807

	pointsOfFunctionPlot[953].X = 2.56
	pointsOfFunctionPlot[953].Y = 12.935

	pointsOfFunctionPlot[954].X = 2.57
	pointsOfFunctionPlot[954].Y = 13.065

	pointsOfFunctionPlot[955].X = 2.58
	pointsOfFunctionPlot[955].Y = 13.197

	pointsOfFunctionPlot[956].X = 2.59
	pointsOfFunctionPlot[956].Y = 13.329

	pointsOfFunctionPlot[957].X = 2.6
	pointsOfFunctionPlot[957].Y = 13.463

	pointsOfFunctionPlot[958].X = 2.61
	pointsOfFunctionPlot[958].Y = 13.599

	pointsOfFunctionPlot[959].X = 2.62
	pointsOfFunctionPlot[959].Y = 13.735

	pointsOfFunctionPlot[960].X = 2.63
	pointsOfFunctionPlot[960].Y = 13.873

	pointsOfFunctionPlot[961].X = 2.64
	pointsOfFunctionPlot[961].Y = 14.013

	pointsOfFunctionPlot[962].X = 2.65
	pointsOfFunctionPlot[962].Y = 14.154

	pointsOfFunctionPlot[963].X = 2.66
	pointsOfFunctionPlot[963].Y = 14.296

	pointsOfFunctionPlot[964].X = 2.67
	pointsOfFunctionPlot[964].Y = 14.44

	pointsOfFunctionPlot[965].X = 2.68
	pointsOfFunctionPlot[965].Y = 14.585

	pointsOfFunctionPlot[966].X = 2.69
	pointsOfFunctionPlot[966].Y = 14.731

	pointsOfFunctionPlot[967].X = 2.7
	pointsOfFunctionPlot[967].Y = 14.879

	pointsOfFunctionPlot[968].X = 2.71
	pointsOfFunctionPlot[968].Y = 15.029

	pointsOfFunctionPlot[969].X = 2.72
	pointsOfFunctionPlot[969].Y = 15.18

	pointsOfFunctionPlot[970].X = 2.73
	pointsOfFunctionPlot[970].Y = 15.332

	pointsOfFunctionPlot[971].X = 2.74
	pointsOfFunctionPlot[971].Y = 15.487

	pointsOfFunctionPlot[972].X = 2.75
	pointsOfFunctionPlot[972].Y = 15.642

	pointsOfFunctionPlot[973].X = 2.76
	pointsOfFunctionPlot[973].Y = 15.799

	pointsOfFunctionPlot[974].X = 2.77
	pointsOfFunctionPlot[974].Y = 15.958

	pointsOfFunctionPlot[975].X = 2.78
	pointsOfFunctionPlot[975].Y = 16.119

	pointsOfFunctionPlot[976].X = 2.79
	pointsOfFunctionPlot[976].Y = 16.281

	pointsOfFunctionPlot[977].X = 2.8
	pointsOfFunctionPlot[977].Y = 16.444

	pointsOfFunctionPlot[978].X = 2.81
	pointsOfFunctionPlot[978].Y = 16.606

	pointsOfFunctionPlot[979].X = 2.82
	pointsOfFunctionPlot[979].Y = 16.776

	pointsOfFunctionPlot[980].X = 2.83
	pointsOfFunctionPlot[980].Y = 16.945

	pointsOfFunctionPlot[981].X = 2.84
	pointsOfFunctionPlot[981].Y = 17.115

	pointsOfFunctionPlot[982].X = 2.85
	pointsOfFunctionPlot[982].Y = 17.287

	pointsOfFunctionPlot[983].X = 2.86
	pointsOfFunctionPlot[983].Y = 17.461

	pointsOfFunctionPlot[984].X = 2.87
	pointsOfFunctionPlot[984].Y = 17.637

	pointsOfFunctionPlot[985].X = 2.88
	pointsOfFunctionPlot[985].Y = 17.814

	pointsOfFunctionPlot[986].X = 2.89
	pointsOfFunctionPlot[986].Y = 17.993

	pointsOfFunctionPlot[987].X = 2.9
	pointsOfFunctionPlot[987].Y = 18.174

	pointsOfFunctionPlot[988].X = 2.91
	pointsOfFunctionPlot[988].Y = 18.356

	pointsOfFunctionPlot[989].X = 2.92
	pointsOfFunctionPlot[989].Y = 18.541

	pointsOfFunctionPlot[990].X = 2.93
	pointsOfFunctionPlot[990].Y = 18.727

	pointsOfFunctionPlot[991].X = 2.94
	pointsOfFunctionPlot[991].Y = 18.915

	pointsOfFunctionPlot[992].X = 2.95
	pointsOfFunctionPlot[992].Y = 19.106

	pointsOfFunctionPlot[993].X = 2.96
	pointsOfFunctionPlot[993].Y = 19.298

	pointsOfFunctionPlot[994].X = 2.97
	pointsOfFunctionPlot[994].Y = 19.491

	pointsOfFunctionPlot[995].X = 2.98
	pointsOfFunctionPlot[995].Y = 19.687

	pointsOfFunctionPlot[996].X = 2.99
	pointsOfFunctionPlot[996].Y = 19.885

	pointsOfFunctionPlot[997].X = 3.0
	pointsOfFunctionPlot[997].Y = 20.085

	pointsOfFunctionPlot[998].X = 3.01
	pointsOfFunctionPlot[998].Y = 20.085

	pointsOfFunctionPlot[999].X = 3.02
	pointsOfFunctionPlot[999].Y = 20.287

	pointsOfFunctionPlot[1_000].X = 3.03
	pointsOfFunctionPlot[1_000].Y = 20.697

	pointsOfFunctionPlot[1_001].X = 3.04
	pointsOfFunctionPlot[1_001].Y = 20.905

	pointsOfFunctionPlot[1_002].X = 3.05
	pointsOfFunctionPlot[1_002].Y = 21.115

	pointsOfFunctionPlot[1_003].X = 3.06
	pointsOfFunctionPlot[1_003].Y = 21.327

	pointsOfFunctionPlot[1_004].X = 3.07
	pointsOfFunctionPlot[1_004].Y = 21.541

	pointsOfFunctionPlot[1_005].X = 3.08
	pointsOfFunctionPlot[1_005].Y = 21.758

	pointsOfFunctionPlot[1_006].X = 3.09
	pointsOfFunctionPlot[1_006].Y = 21.977

	pointsOfFunctionPlot[1_007].X = 3.1
	pointsOfFunctionPlot[1_007].Y = 22.198

	pointsOfFunctionPlot[1_008].X = 3.11
	pointsOfFunctionPlot[1_008].Y = 22.421

	pointsOfFunctionPlot[1_009].X = 3.12
	pointsOfFunctionPlot[1_009].Y = 22.646

	pointsOfFunctionPlot[1_010].X = 3.13
	pointsOfFunctionPlot[1_010].Y = 22.874

	pointsOfFunctionPlot[1_011].X = 3.14
	pointsOfFunctionPlot[1_011].Y = 23.103

	pointsOfFunctionPlot[1_012].X = 3.15
	pointsOfFunctionPlot[1_012].Y = 23.336

	pointsOfFunctionPlot[1_013].X = 3.16
	pointsOfFunctionPlot[1_013].Y = 23.57

	pointsOfFunctionPlot[1_014].X = 3.17
	pointsOfFunctionPlot[1_014].Y = 23.807

	pointsOfFunctionPlot[1_015].X = 3.18
	pointsOfFunctionPlot[1_015].Y = 24.046

	pointsOfFunctionPlot[1_016].X = 3.19
	pointsOfFunctionPlot[1_016].Y = 24.288

	pointsOfFunctionPlot[1_017].X = 3.2
	pointsOfFunctionPlot[1_017].Y = 24.532

	pointsOfFunctionPlot[1_018].X = 3.21
	pointsOfFunctionPlot[1_018].Y = 24.779

	pointsOfFunctionPlot[1_019].X = 3.22
	pointsOfFunctionPlot[1_019].Y = 25.028

	pointsOfFunctionPlot[1_020].X = 3.23
	pointsOfFunctionPlot[1_020].Y = 25.279

	pointsOfFunctionPlot[1_021].X = 3.24
	pointsOfFunctionPlot[1_021].Y = 25.533

	pointsOfFunctionPlot[1_022].X = 3.25
	pointsOfFunctionPlot[1_022].Y = 25.79

	pointsOfFunctionPlot[1_023].X = 3.26
	pointsOfFunctionPlot[1_023].Y = 26.049

	pointsOfFunctionPlot[1_024].X = 3.27
	pointsOfFunctionPlot[1_024].Y = 26.311

	pointsOfFunctionPlot[1_025].X = 3.28
	pointsOfFunctionPlot[1_025].Y = 26.575

	pointsOfFunctionPlot[1_026].X = 3.29
	pointsOfFunctionPlot[1_026].Y = 26.842

	pointsOfFunctionPlot[1_027].X = 3.3
	pointsOfFunctionPlot[1_027].Y = 27.112

	pointsOfFunctionPlot[1_028].X = 3.31
	pointsOfFunctionPlot[1_028].Y = 27.385

	pointsOfFunctionPlot[1_029].X = 3.32
	pointsOfFunctionPlot[1_029].Y = 27.66

	pointsOfFunctionPlot[1_030].X = 3.33
	pointsOfFunctionPlot[1_030].Y = 27.938

	pointsOfFunctionPlot[1_031].X = 3.34
	pointsOfFunctionPlot[1_031].Y = 28.219

	pointsOfFunctionPlot[1_032].X = 3.35
	pointsOfFunctionPlot[1_032].Y = 28.502

	pointsOfFunctionPlot[1_033].X = 3.36
	pointsOfFunctionPlot[1_033].Y = 28.789

	pointsOfFunctionPlot[1_034].X = 3.37
	pointsOfFunctionPlot[1_034].Y = 29.078

	pointsOfFunctionPlot[1_035].X = 3.38
	pointsOfFunctionPlot[1_035].Y = 29.37

	pointsOfFunctionPlot[1_036].X = 3.39
	pointsOfFunctionPlot[1_036].Y = 29.666

	pointsOfFunctionPlot[1_037].X = 3.4
	pointsOfFunctionPlot[1_037].Y = 29.964

	pointsOfFunctionPlot[1_038].X = 3.41
	pointsOfFunctionPlot[1_038].Y = 30.265

	pointsOfFunctionPlot[1_039].X = 3.42
	pointsOfFunctionPlot[1_039].Y = 30.569

	pointsOfFunctionPlot[1_040].X = 3.43
	pointsOfFunctionPlot[1_040].Y = 30.876

	pointsOfFunctionPlot[1_041].X = 3.44
	pointsOfFunctionPlot[1_041].Y = 31.187

	pointsOfFunctionPlot[1_042].X = 3.45
	pointsOfFunctionPlot[1_042].Y = 31.5

	pointsOfFunctionPlot[1_043].X = 3.46
	pointsOfFunctionPlot[1_043].Y = 31.817

	pointsOfFunctionPlot[1_044].X = 3.47
	pointsOfFunctionPlot[1_044].Y = 32.136

	pointsOfFunctionPlot[1_045].X = 3.48
	pointsOfFunctionPlot[1_045].Y = 32.459

	pointsOfFunctionPlot[1_046].X = 3.49
	pointsOfFunctionPlot[1_046].Y = 32.785

	pointsOfFunctionPlot[1_047].X = 3.5
	pointsOfFunctionPlot[1_047].Y = 33.115

	pointsOfFunctionPlot[1_048].X = 3.51
	pointsOfFunctionPlot[1_048].Y = 33.448

	pointsOfFunctionPlot[1_049].X = 3.52
	pointsOfFunctionPlot[1_049].Y = 33.784

	pointsOfFunctionPlot[1_050].X = 3.53
	pointsOfFunctionPlot[1_050].Y = 34.124

	pointsOfFunctionPlot[1_051].X = 3.54
	pointsOfFunctionPlot[1_051].Y = 34.466

	pointsOfFunctionPlot[1_052].X = 3.55
	pointsOfFunctionPlot[1_052].Y = 34.813

	pointsOfFunctionPlot[1_053].X = 3.56
	pointsOfFunctionPlot[1_053].Y = 35.163

	pointsOfFunctionPlot[1_054].X = 3.57
	pointsOfFunctionPlot[1_054].Y = 35.516

	pointsOfFunctionPlot[1_055].X = 3.58
	pointsOfFunctionPlot[1_055].Y = 35.873

	pointsOfFunctionPlot[1_056].X = 3.59
	pointsOfFunctionPlot[1_056].Y = 36.234

	pointsOfFunctionPlot[1_057].X = 3.6
	pointsOfFunctionPlot[1_057].Y = 36.598

	pointsOfFunctionPlot[1_058].X = 3.61
	pointsOfFunctionPlot[1_058].Y = 36.966

	pointsOfFunctionPlot[1_059].X = 3.62
	pointsOfFunctionPlot[1_059].Y = 37.337

	pointsOfFunctionPlot[1_060].X = 3.63
	pointsOfFunctionPlot[1_060].Y = 37.712

	pointsOfFunctionPlot[1_061].X = 3.64
	pointsOfFunctionPlot[1_061].Y = 38.091

	pointsOfFunctionPlot[1_062].X = 3.65
	pointsOfFunctionPlot[1_062].Y = 38.474

	pointsOfFunctionPlot[1_063].X = 3.66
	pointsOfFunctionPlot[1_063].Y = 38.861

	pointsOfFunctionPlot[1_064].X = 3.67
	pointsOfFunctionPlot[1_064].Y = 39.251

	pointsOfFunctionPlot[1_065].X = 3.68
	pointsOfFunctionPlot[1_065].Y = 39.646

	pointsOfFunctionPlot[1_066].X = 3.69
	pointsOfFunctionPlot[1_066].Y = 40.044

	pointsOfFunctionPlot[1_067].X = 3.7
	pointsOfFunctionPlot[1_067].Y = 40.447

	pointsOfFunctionPlot[1_068].X = 3.71
	pointsOfFunctionPlot[1_068].Y = 40.853

	pointsOfFunctionPlot[1_069].X = 3.72
	pointsOfFunctionPlot[1_069].Y = 41.264

	pointsOfFunctionPlot[1_070].X = 3.73
	pointsOfFunctionPlot[1_070].Y = 41.679

	pointsOfFunctionPlot[1_071].X = 3.74
	pointsOfFunctionPlot[1_071].Y = 42.098

	pointsOfFunctionPlot[1_072].X = 3.75
	pointsOfFunctionPlot[1_072].Y = 42.521

	pointsOfFunctionPlot[1_073].X = 3.76
	pointsOfFunctionPlot[1_073].Y = 42.948

	pointsOfFunctionPlot[1_074].X = 3.77
	pointsOfFunctionPlot[1_074].Y = 43.38

	pointsOfFunctionPlot[1_075].X = 3.78
	pointsOfFunctionPlot[1_075].Y = 43.816

	pointsOfFunctionPlot[1_076].X = 3.79
	pointsOfFunctionPlot[1_076].Y = 44.256

	pointsOfFunctionPlot[1_077].X = 3.8
	pointsOfFunctionPlot[1_077].Y = 44.701

	pointsOfFunctionPlot[1_078].X = 3.81
	pointsOfFunctionPlot[1_078].Y = 45.15

	pointsOfFunctionPlot[1_079].X = 3.82
	pointsOfFunctionPlot[1_079].Y = 45.604

	pointsOfFunctionPlot[1_080].X = 3.83
	pointsOfFunctionPlot[1_080].Y = 46.062

	pointsOfFunctionPlot[1_081].X = 3.84
	pointsOfFunctionPlot[1_081].Y = 46.525

	pointsOfFunctionPlot[1_082].X = 3.85
	pointsOfFunctionPlot[1_082].Y = 46.993

	pointsOfFunctionPlot[1_083].X = 3.86
	pointsOfFunctionPlot[1_083].Y = 47.465

	pointsOfFunctionPlot[1_084].X = 3.87
	pointsOfFunctionPlot[1_084].Y = 47.942

	pointsOfFunctionPlot[1_085].X = 3.88
	pointsOfFunctionPlot[1_085].Y = 48.424

	pointsOfFunctionPlot[1_086].X = 3.89
	pointsOfFunctionPlot[1_086].Y = 48.91

	pointsOfFunctionPlot[1_087].X = 3.9
	pointsOfFunctionPlot[1_087].Y = 49.402

	pointsOfFunctionPlot[1_088].X = 3.91
	pointsOfFunctionPlot[1_088].Y = 49.899

	pointsOfFunctionPlot[1_089].X = 3.92
	pointsOfFunctionPlot[1_089].Y = 50.4

	pointsOfFunctionPlot[1_090].X = 3.93
	pointsOfFunctionPlot[1_090].Y = 50.907

	pointsOfFunctionPlot[1_091].X = 3.94
	pointsOfFunctionPlot[1_091].Y = 51.418

	pointsOfFunctionPlot[1_092].X = 3.95
	pointsOfFunctionPlot[1_092].Y = 51.935

	pointsOfFunctionPlot[1_093].X = 3.96
	pointsOfFunctionPlot[1_093].Y = 52.457

	pointsOfFunctionPlot[1_094].X = 3.97
	pointsOfFunctionPlot[1_094].Y = 52.984

	pointsOfFunctionPlot[1_095].X = 3.98
	pointsOfFunctionPlot[1_095].Y = 53.517

	pointsOfFunctionPlot[1_096].X = 3.99
	pointsOfFunctionPlot[1_096].Y = 54.054

	pointsOfFunctionPlot[1_097].X = 4.0
	pointsOfFunctionPlot[1_097].Y = 54.598

	pointsOfFunctionPlot[1_098].X = 4.01
	pointsOfFunctionPlot[1_098].Y = 55.146

	pointsOfFunctionPlot[1_099].X = 4.02
	pointsOfFunctionPlot[1_099].Y = 55.701

	pointsOfFunctionPlot[1_100].X = 4.03
	pointsOfFunctionPlot[1_100].Y = 56.26

	pointsOfFunctionPlot[1_101].X = 4.04
	pointsOfFunctionPlot[1_101].Y = 56.826

	pointsOfFunctionPlot[1_102].X = 4.05
	pointsOfFunctionPlot[1_102].Y = 57.397

	pointsOfFunctionPlot[1_103].X = 4.06
	pointsOfFunctionPlot[1_103].Y = 57.974

	pointsOfFunctionPlot[1_104].X = 4.07
	pointsOfFunctionPlot[1_104].Y = 58.557

	pointsOfFunctionPlot[1_105].X = 4.08
	pointsOfFunctionPlot[1_105].Y = 59.145

	pointsOfFunctionPlot[1_106].X = 4.09
	pointsOfFunctionPlot[1_106].Y = 59.739

	pointsOfFunctionPlot[1_107].X = 4.1
	pointsOfFunctionPlot[1_107].Y = 60.34

	pointsOfFunctionPlot[1_108].X = 4.11
	pointsOfFunctionPlot[1_108].Y = 60.946

	pointsOfFunctionPlot[1_109].X = 4.12
	pointsOfFunctionPlot[1_109].Y = 61.559

	pointsOfFunctionPlot[1_110].X = 4.13
	pointsOfFunctionPlot[1_110].Y = 62.177

	pointsOfFunctionPlot[1_111].X = 4.14
	pointsOfFunctionPlot[1_111].Y = 62.802

	pointsOfFunctionPlot[1_112].X = 4.15
	pointsOfFunctionPlot[1_112].Y = 63.434

	pointsOfFunctionPlot[1_113].X = 4.16
	pointsOfFunctionPlot[1_113].Y = 64.071

	pointsOfFunctionPlot[1_114].X = 4.17
	pointsOfFunctionPlot[1_114].Y = 64.715

	pointsOfFunctionPlot[1_115].X = 4.18
	pointsOfFunctionPlot[1_115].Y = 65.365

	pointsOfFunctionPlot[1_116].X = 4.19
	pointsOfFunctionPlot[1_116].Y = 66.022

	pointsOfFunctionPlot[1_117].X = 4.2
	pointsOfFunctionPlot[1_117].Y = 66.686

	pointsOfFunctionPlot[1_118].X = 4.21
	pointsOfFunctionPlot[1_118].Y = 67.356

	pointsOfFunctionPlot[1_119].X = 4.22
	pointsOfFunctionPlot[1_119].Y = 68.033

	pointsOfFunctionPlot[1_120].X = 4.23
	pointsOfFunctionPlot[1_120].Y = 68.717

	pointsOfFunctionPlot[1_121].X = 4.24
	pointsOfFunctionPlot[1_121].Y = 69.407

	pointsOfFunctionPlot[1_122].X = 4.25
	pointsOfFunctionPlot[1_122].Y = 70.105

	pointsOfFunctionPlot[1_123].X = 4.26
	pointsOfFunctionPlot[1_123].Y = 70.81

	pointsOfFunctionPlot[1_124].X = 4.27
	pointsOfFunctionPlot[1_124].Y = 71.521

	pointsOfFunctionPlot[1_125].X = 4.28
	pointsOfFunctionPlot[1_125].Y = 72.24

	pointsOfFunctionPlot[1_126].X = 4.29
	pointsOfFunctionPlot[1_126].Y = 72.966

	pointsOfFunctionPlot[1_127].X = 4.3
	pointsOfFunctionPlot[1_127].Y = 73.699

	pointsOfFunctionPlot[1_128].X = 4.31
	pointsOfFunctionPlot[1_128].Y = 74.44

	pointsOfFunctionPlot[1_129].X = 4.32
	pointsOfFunctionPlot[1_129].Y = 75.188

	pointsOfFunctionPlot[1_130].X = 4.33
	pointsOfFunctionPlot[1_130].Y = 75.944

	pointsOfFunctionPlot[1_131].X = 4.34
	pointsOfFunctionPlot[1_131].Y = 76.707

	pointsOfFunctionPlot[1_132].X = 4.35
	pointsOfFunctionPlot[1_132].Y = 77.478

	pointsOfFunctionPlot[1_133].X = 4.36
	pointsOfFunctionPlot[1_133].Y = 78.257

	pointsOfFunctionPlot[1_134].X = 4.37
	pointsOfFunctionPlot[1_134].Y = 79.043

	pointsOfFunctionPlot[1_135].X = 4.38
	pointsOfFunctionPlot[1_135].Y = 79.838

	pointsOfFunctionPlot[1_136].X = 4.39
	pointsOfFunctionPlot[1_136].Y = 80.64

	pointsOfFunctionPlot[1_137].X = 4.4
	pointsOfFunctionPlot[1_137].Y = 81.45

	pointsOfFunctionPlot[1_138].X = 4.41
	pointsOfFunctionPlot[1_138].Y = 82.269

	pointsOfFunctionPlot[1_139].X = 4.42
	pointsOfFunctionPlot[1_139].Y = 83.096

	pointsOfFunctionPlot[1_140].X = 4.43
	pointsOfFunctionPlot[1_140].Y = 83.931

	pointsOfFunctionPlot[1_141].X = 4.44
	pointsOfFunctionPlot[1_141].Y = 84.774

	pointsOfFunctionPlot[1_142].X = 4.45
	pointsOfFunctionPlot[1_142].Y = 85.626

	pointsOfFunctionPlot[1_143].X = 4.46
	pointsOfFunctionPlot[1_143].Y = 86.487

	pointsOfFunctionPlot[1_144].X = 4.47
	pointsOfFunctionPlot[1_144].Y = 87.356

	pointsOfFunctionPlot[1_145].X = 4.48
	pointsOfFunctionPlot[1_145].Y = 88.234

	pointsOfFunctionPlot[1_146].X = 4.49
	pointsOfFunctionPlot[1_146].Y = 89.121

	pointsOfFunctionPlot[1_147].X = 4.5
	pointsOfFunctionPlot[1_147].Y = 90.017

	pointsOfFunctionPlot[1_148].X = 4.51
	pointsOfFunctionPlot[1_148].Y = 90.921

	pointsOfFunctionPlot[1_149].X = 4.52
	pointsOfFunctionPlot[1_149].Y = 91.835

	pointsOfFunctionPlot[1_150].X = 4.53
	pointsOfFunctionPlot[1_150].Y = 92.758

	pointsOfFunctionPlot[1_151].X = 4.54
	pointsOfFunctionPlot[1_151].Y = 93.69

	pointsOfFunctionPlot[1_152].X = 4.55
	pointsOfFunctionPlot[1_152].Y = 94.632

	pointsOfFunctionPlot[1_153].X = 4.56
	pointsOfFunctionPlot[1_153].Y = 95.583

	pointsOfFunctionPlot[1_154].X = 4.57
	pointsOfFunctionPlot[1_154].Y = 96.544

	pointsOfFunctionPlot[1_155].X = 4.58
	pointsOfFunctionPlot[1_155].Y = 97.514

	pointsOfFunctionPlot[1_156].X = 4.59
	pointsOfFunctionPlot[1_156].Y = 98.494

	pointsOfFunctionPlot[1_157].X = 4.6
	pointsOfFunctionPlot[1_157].Y = 99.484

	pointsOfFunctionPlot[1_158].X = 4.61
	pointsOfFunctionPlot[1_158].Y = 100.484

	pointsOfFunctionPlot[1_159].X = 4.62
	pointsOfFunctionPlot[1_159].Y = 101.494

	pointsOfFunctionPlot[1_160].X = 4.63
	pointsOfFunctionPlot[1_160].Y = 102.514

	pointsOfFunctionPlot[1_161].X = 4.64
	pointsOfFunctionPlot[1_161].Y = 103.544

	pointsOfFunctionPlot[1_162].X = 4.65
	pointsOfFunctionPlot[1_162].Y = 104.585

	pointsOfFunctionPlot[1_163].X = 4.66
	pointsOfFunctionPlot[1_163].Y = 105.636

	pointsOfFunctionPlot[1_164].X = 4.67
	pointsOfFunctionPlot[1_164].Y = 106.697

	pointsOfFunctionPlot[1_165].X = 4.68
	pointsOfFunctionPlot[1_165].Y = 107.77

	pointsOfFunctionPlot[1_166].X = 4.69
	pointsOfFunctionPlot[1_166].Y = 108.853

	pointsOfFunctionPlot[1_167].X = 4.7
	pointsOfFunctionPlot[1_167].Y = 109.947

	pointsOfFunctionPlot[1_168].X = 4.71
	pointsOfFunctionPlot[1_168].Y = 111.052

	pointsOfFunctionPlot[1_169].X = 4.72
	pointsOfFunctionPlot[1_169].Y = 112.168

	pointsOfFunctionPlot[1_170].X = 4.73
	pointsOfFunctionPlot[1_170].Y = 113.295

	pointsOfFunctionPlot[1_171].X = 4.74
	pointsOfFunctionPlot[1_171].Y = 114.434

	pointsOfFunctionPlot[1_172].X = 4.75
	pointsOfFunctionPlot[1_172].Y = 115.584

	pointsOfFunctionPlot[1_173].X = 4.76
	pointsOfFunctionPlot[1_173].Y = 116.745

	pointsOfFunctionPlot[1_174].X = 4.77
	pointsOfFunctionPlot[1_174].Y = 117.919

	pointsOfFunctionPlot[1_175].X = 4.78
	pointsOfFunctionPlot[1_175].Y = 119.104

	pointsOfFunctionPlot[1_176].X = 4.79
	pointsOfFunctionPlot[1_176].Y = 120.301

	pointsOfFunctionPlot[1_177].X = 4.8
	pointsOfFunctionPlot[1_177].Y = 121.51

	pointsOfFunctionPlot[1_178].X = 4.81
	pointsOfFunctionPlot[1_178].Y = 122.731

	pointsOfFunctionPlot[1_179].X = 4.82
	pointsOfFunctionPlot[1_179].Y = 123.965

	pointsOfFunctionPlot[1_180].X = 4.83
	pointsOfFunctionPlot[1_180].Y = 125.211

	pointsOfFunctionPlot[1_181].X = 4.84
	pointsOfFunctionPlot[1_181].Y = 126.469

	pointsOfFunctionPlot[1_182].X = 4.85
	pointsOfFunctionPlot[1_182].Y = 127.74

	pointsOfFunctionPlot[1_183].X = 4.86
	pointsOfFunctionPlot[1_183].Y = 129.024

	pointsOfFunctionPlot[1_184].X = 4.87
	pointsOfFunctionPlot[1_184].Y = 130.32

	pointsOfFunctionPlot[1_185].X = 4.88
	pointsOfFunctionPlot[1_185].Y = 131.63

	pointsOfFunctionPlot[1_186].X = 4.89
	pointsOfFunctionPlot[1_186].Y = 132.953

	pointsOfFunctionPlot[1_187].X = 4.9
	pointsOfFunctionPlot[1_187].Y = 134.289

	pointsOfFunctionPlot[1_188].X = 4.91
	pointsOfFunctionPlot[1_188].Y = 135.639

	pointsOfFunctionPlot[1_189].X = 4.92
	pointsOfFunctionPlot[1_189].Y = 137.002

	pointsOfFunctionPlot[1_190].X = 4.93
	pointsOfFunctionPlot[1_190].Y = 138.379

	pointsOfFunctionPlot[1_191].X = 4.94
	pointsOfFunctionPlot[1_191].Y = 139.77

	pointsOfFunctionPlot[1_192].X = 4.95
	pointsOfFunctionPlot[1_192].Y = 141.175

	pointsOfFunctionPlot[1_193].X = 4.96
	pointsOfFunctionPlot[1_193].Y = 142.593

	pointsOfFunctionPlot[1_194].X = 4.97
	pointsOfFunctionPlot[1_194].Y = 144.026

	pointsOfFunctionPlot[1_195].X = 4.98
	pointsOfFunctionPlot[1_195].Y = 145.474

	pointsOfFunctionPlot[1_196].X = 4.99
	pointsOfFunctionPlot[1_196].Y = 146.936

	pointsOfFunctionPlot[1_197].X = 5.0
	pointsOfFunctionPlot[1_197].Y = 148.413

	pointsOfFunctionPlot[1_198].X = 5.01
	pointsOfFunctionPlot[1_198].Y = 149.904

	pointsOfFunctionPlot[1_199].X = 5.02
	pointsOfFunctionPlot[1_199].Y = 151.411

	pointsOfFunctionPlot[1_200].X = 5.03
	pointsOfFunctionPlot[1_200].Y = 152.933

	pointsOfFunctionPlot[1_201].X = 5.04
	pointsOfFunctionPlot[1_201].Y = 154.47

	pointsOfFunctionPlot[1_202].X = 5.05
	pointsOfFunctionPlot[1_202].Y = 156.022

	pointsOfFunctionPlot[1_203].X = 5.06
	pointsOfFunctionPlot[1_203].Y = 157.59

	pointsOfFunctionPlot[1_204].X = 5.07
	pointsOfFunctionPlot[1_204].Y = 159.174

	pointsOfFunctionPlot[1_205].X = 5.08
	pointsOfFunctionPlot[1_205].Y = 160.774

	pointsOfFunctionPlot[1_206].X = 5.09
	pointsOfFunctionPlot[1_206].Y = 162.389

	pointsOfFunctionPlot[1_207].X = 5.1
	pointsOfFunctionPlot[1_207].Y = 164.021

	pointsOfFunctionPlot[1_208].X = 5.11
	pointsOfFunctionPlot[1_208].Y = 165.67

	pointsOfFunctionPlot[1_209].X = 5.12
	pointsOfFunctionPlot[1_209].Y = 167.335

	pointsOfFunctionPlot[1_210].X = 5.13
	pointsOfFunctionPlot[1_210].Y = 169.017

	pointsOfFunctionPlot[1_211].X = 5.14
	pointsOfFunctionPlot[1_211].Y = 170.715

	pointsOfFunctionPlot[1_212].X = 5.15
	pointsOfFunctionPlot[1_212].Y = 172.431

	pointsOfFunctionPlot[1_213].X = 5.16
	pointsOfFunctionPlot[1_213].Y = 174.164

	pointsOfFunctionPlot[1_214].X = 5.17
	pointsOfFunctionPlot[1_214].Y = 175.914

	pointsOfFunctionPlot[1_215].X = 5.18
	pointsOfFunctionPlot[1_215].Y = 177.682

	pointsOfFunctionPlot[1_216].X = 5.19
	pointsOfFunctionPlot[1_216].Y = 179.468

	pointsOfFunctionPlot[1_217].X = 5.2
	pointsOfFunctionPlot[1_217].Y = 181.272

	pointsOfFunctionPlot[1_218].X = 5.21
	pointsOfFunctionPlot[1_218].Y = 183.094

	pointsOfFunctionPlot[1_219].X = 5.22
	pointsOfFunctionPlot[1_219].Y = 184.934

	pointsOfFunctionPlot[1_220].X = 5.23
	pointsOfFunctionPlot[1_220].Y = 186.792

	pointsOfFunctionPlot[1_221].X = 5.24
	pointsOfFunctionPlot[1_221].Y = 188.67

	pointsOfFunctionPlot[1_222].X = 5.25
	pointsOfFunctionPlot[1_222].Y = 190.566

	pointsOfFunctionPlot[1_223].X = 5.26
	pointsOfFunctionPlot[1_223].Y = 192.481

	pointsOfFunctionPlot[1_224].X = 5.27
	pointsOfFunctionPlot[1_224].Y = 194.416

	pointsOfFunctionPlot[1_225].X = 5.28
	pointsOfFunctionPlot[1_225].Y = 196.369

	pointsOfFunctionPlot[1_226].X = 5.29
	pointsOfFunctionPlot[1_226].Y = 198.343

	pointsOfFunctionPlot[1_227].X = 5.3
	pointsOfFunctionPlot[1_227].Y = 200.336

	pointsOfFunctionPlot[1_228].X = 5.31
	pointsOfFunctionPlot[1_228].Y = 202.35

	pointsOfFunctionPlot[1_229].X = 5.32
	pointsOfFunctionPlot[1_229].Y = 204.383

	pointsOfFunctionPlot[1_230].X = 5.33
	pointsOfFunctionPlot[1_230].Y = 206.438

	pointsOfFunctionPlot[1_231].X = 5.34
	pointsOfFunctionPlot[1_231].Y = 208.512

	pointsOfFunctionPlot[1_232].X = 5.35
	pointsOfFunctionPlot[1_232].Y = 210.608

	pointsOfFunctionPlot[1_233].X = 5.36
	pointsOfFunctionPlot[1_233].Y = 212.724

	pointsOfFunctionPlot[1_234].X = 5.37
	pointsOfFunctionPlot[1_234].Y = 214.862

	pointsOfFunctionPlot[1_235].X = 5.38
	pointsOfFunctionPlot[1_235].Y = 217.022

	pointsOfFunctionPlot[1_236].X = 5.39
	pointsOfFunctionPlot[1_236].Y = 219.203

	pointsOfFunctionPlot[1_237].X = 5.4
	pointsOfFunctionPlot[1_237].Y = 221.406

	pointsOfFunctionPlot[1_238].X = 5.41
	pointsOfFunctionPlot[1_238].Y = 223.631

	pointsOfFunctionPlot[1_239].X = 5.42
	pointsOfFunctionPlot[1_239].Y = 225.879

	pointsOfFunctionPlot[1_240].X = 5.43
	pointsOfFunctionPlot[1_240].Y = 228.149

	pointsOfFunctionPlot[1_241].X = 5.44
	pointsOfFunctionPlot[1_241].Y = 230.442

	pointsOfFunctionPlot[1_242].X = 5.45
	pointsOfFunctionPlot[1_242].Y = 232.758

	pointsOfFunctionPlot[1_243].X = 5.46
	pointsOfFunctionPlot[1_243].Y = 235.097

	pointsOfFunctionPlot[1_244].X = 5.47
	pointsOfFunctionPlot[1_244].Y = 237.46

	pointsOfFunctionPlot[1_245].X = 5.48
	pointsOfFunctionPlot[1_245].Y = 239.846

	pointsOfFunctionPlot[1_246].X = 5.49
	pointsOfFunctionPlot[1_246].Y = 242.257

	pointsOfFunctionPlot[1_247].X = 5.5
	pointsOfFunctionPlot[1_247].Y = 244.691

	pointsOfFunctionPlot[1_248].X = 5.51
	pointsOfFunctionPlot[1_248].Y = 247.151

	pointsOfFunctionPlot[1_249].X = 5.52
	pointsOfFunctionPlot[1_249].Y = 249.635

	pointsOfFunctionPlot[1_250].X = 5.53
	pointsOfFunctionPlot[1_250].Y = 252.143

	pointsOfFunctionPlot[1_251].X = 5.54
	pointsOfFunctionPlot[1_251].Y = 254.678

	pointsOfFunctionPlot[1_252].X = 5.55
	pointsOfFunctionPlot[1_252].Y = 257.237

	pointsOfFunctionPlot[1_253].X = 5.56
	pointsOfFunctionPlot[1_253].Y = 259.822

	pointsOfFunctionPlot[1_254].X = 5.57
	pointsOfFunctionPlot[1_254].Y = 262.434

	pointsOfFunctionPlot[1_255].X = 5.58
	pointsOfFunctionPlot[1_255].Y = 265.071

	pointsOfFunctionPlot[1_256].X = 5.59
	pointsOfFunctionPlot[1_256].Y = 267.735

	pointsOfFunctionPlot[1_257].X = 5.6
	pointsOfFunctionPlot[1_257].Y = 270.426

	pointsOfFunctionPlot[1_258].X = 5.61
	pointsOfFunctionPlot[1_258].Y = 273.144

	pointsOfFunctionPlot[1_259].X = 5.62
	pointsOfFunctionPlot[1_259].Y = 275.889

	pointsOfFunctionPlot[1_260].X = 5.63
	pointsOfFunctionPlot[1_260].Y = 278.662

	pointsOfFunctionPlot[1_261].X = 5.64
	pointsOfFunctionPlot[1_261].Y = 281.462

	pointsOfFunctionPlot[1_262].X = 5.65
	pointsOfFunctionPlot[1_262].Y = 284.291

	pointsOfFunctionPlot[1_263].X = 5.66
	pointsOfFunctionPlot[1_263].Y = 287.148

	pointsOfFunctionPlot[1_264].X = 5.67
	pointsOfFunctionPlot[1_264].Y = 290.034

	pointsOfFunctionPlot[1_265].X = 5.68
	pointsOfFunctionPlot[1_265].Y = 292.949

	pointsOfFunctionPlot[1_266].X = 5.69
	pointsOfFunctionPlot[1_266].Y = 295.893

	pointsOfFunctionPlot[1_267].X = 5.7
	pointsOfFunctionPlot[1_267].Y = 298.867

	pointsOfFunctionPlot[1_268].X = 5.71
	pointsOfFunctionPlot[1_268].Y = 301.871

	pointsOfFunctionPlot[1_269].X = 5.72
	pointsOfFunctionPlot[1_269].Y = 304.904

	pointsOfFunctionPlot[1_270].X = 5.73
	pointsOfFunctionPlot[1_270].Y = 307.969

	pointsOfFunctionPlot[1_271].X = 5.74
	pointsOfFunctionPlot[1_271].Y = 311.064

	pointsOfFunctionPlot[1_272].X = 5.75
	pointsOfFunctionPlot[1_272].Y = 314.19

	pointsOfFunctionPlot[1_273].X = 5.76
	pointsOfFunctionPlot[1_273].Y = 317.348

	pointsOfFunctionPlot[1_274].X = 5.77
	pointsOfFunctionPlot[1_274].Y = 320.537

	pointsOfFunctionPlot[1_275].X = 5.78
	pointsOfFunctionPlot[1_275].Y = 323.759

	pointsOfFunctionPlot[1_276].X = 5.79
	pointsOfFunctionPlot[1_276].Y = 327.013

	pointsOfFunctionPlot[1_277].X = 5.8
	pointsOfFunctionPlot[1_277].Y = 330.299

	pointsOfFunctionPlot[1_278].X = 5.81
	pointsOfFunctionPlot[1_278].Y = 333.619

	pointsOfFunctionPlot[1_279].X = 5.82
	pointsOfFunctionPlot[1_279].Y = 336.972

	pointsOfFunctionPlot[1_280].X = 5.83
	pointsOfFunctionPlot[1_280].Y = 340.358

	pointsOfFunctionPlot[1_281].X = 5.84
	pointsOfFunctionPlot[1_281].Y = 343.779

	pointsOfFunctionPlot[1_282].X = 5.85
	pointsOfFunctionPlot[1_282].Y = 347.234

	pointsOfFunctionPlot[1_283].X = 5.86
	pointsOfFunctionPlot[1_283].Y = 350.724

	pointsOfFunctionPlot[1_284].X = 5.87
	pointsOfFunctionPlot[1_284].Y = 354.249

	pointsOfFunctionPlot[1_285].X = 5.88
	pointsOfFunctionPlot[1_285].Y = 357.809

	pointsOfFunctionPlot[1_286].X = 5.89
	pointsOfFunctionPlot[1_286].Y = 361.405

	pointsOfFunctionPlot[1_287].X = 5.9
	pointsOfFunctionPlot[1_287].Y = 365.037

	pointsOfFunctionPlot[1_288].X = 5.91
	pointsOfFunctionPlot[1_288].Y = 368.706

	pointsOfFunctionPlot[1_289].X = 5.92
	pointsOfFunctionPlot[1_289].Y = 372.411

	pointsOfFunctionPlot[1_290].X = 5.93
	pointsOfFunctionPlot[1_290].Y = 376.154

	pointsOfFunctionPlot[1_291].X = 5.94
	pointsOfFunctionPlot[1_291].Y = 379.934

	pointsOfFunctionPlot[1_292].X = 5.95
	pointsOfFunctionPlot[1_292].Y = 383.753

	pointsOfFunctionPlot[1_293].X = 5.96
	pointsOfFunctionPlot[1_293].Y = 387.61

	pointsOfFunctionPlot[1_294].X = 5.97
	pointsOfFunctionPlot[1_294].Y = 391.505

	pointsOfFunctionPlot[1_295].X = 5.98
	pointsOfFunctionPlot[1_295].Y = 395.44

	pointsOfFunctionPlot[1_296].X = 5.99
	pointsOfFunctionPlot[1_296].Y = 399.414

	pointsOfFunctionPlot[1_297].X = 6.0
	pointsOfFunctionPlot[1_297].Y = 403.428

	pointsOfFunctionPlot[1_298].X = 6.01
	pointsOfFunctionPlot[1_298].Y = 407.483

	pointsOfFunctionPlot[1_299].X = 6.02
	pointsOfFunctionPlot[1_299].Y = 411.578

	pointsOfFunctionPlot[1_300].X = 6.03
	pointsOfFunctionPlot[1_300].Y = 415.715

	pointsOfFunctionPlot[1_301].X = 6.04
	pointsOfFunctionPlot[1_301].Y = 419.893

	pointsOfFunctionPlot[1_302].X = 6.05
	pointsOfFunctionPlot[1_302].Y = 424.113

	pointsOfFunctionPlot[1_303].X = 6.06
	pointsOfFunctionPlot[1_303].Y = 428.375

	pointsOfFunctionPlot[1_304].X = 6.07
	pointsOfFunctionPlot[1_304].Y = 432.68

	pointsOfFunctionPlot[1_305].X = 6.08
	pointsOfFunctionPlot[1_305].Y = 437.029

	pointsOfFunctionPlot[1_306].X = 6.09
	pointsOfFunctionPlot[1_306].Y = 441.421

	pointsOfFunctionPlot[1_307].X = 6.1
	pointsOfFunctionPlot[1_307].Y = 445.857

	pointsOfFunctionPlot[1_308].X = 6.11
	pointsOfFunctionPlot[1_308].Y = 450.338

	pointsOfFunctionPlot[1_309].X = 6.12
	pointsOfFunctionPlot[1_309].Y = 454.864

	pointsOfFunctionPlot[1_310].X = 6.13
	pointsOfFunctionPlot[1_310].Y = 459.436

	pointsOfFunctionPlot[1_311].X = 6.14
	pointsOfFunctionPlot[1_311].Y = 464.053

	pointsOfFunctionPlot[1_312].X = 6.15
	pointsOfFunctionPlot[1_312].Y = 468.717

	pointsOfFunctionPlot[1_313].X = 6.16
	pointsOfFunctionPlot[1_313].Y = 473.428

	pointsOfFunctionPlot[1_314].X = 6.17
	pointsOfFunctionPlot[1_314].Y = 478.186

	pointsOfFunctionPlot[1_315].X = 6.18
	pointsOfFunctionPlot[1_315].Y = 482.992

	pointsOfFunctionPlot[1_316].X = 6.19
	pointsOfFunctionPlot[1_316].Y = 487.846

	pointsOfFunctionPlot[1_317].X = 6.2
	pointsOfFunctionPlot[1_317].Y = 492.749

	pointsOfFunctionPlot[1_318].X = 6.21
	pointsOfFunctionPlot[1_318].Y = 497.701

	pointsOfFunctionPlot[1_319].X = 6.22
	pointsOfFunctionPlot[1_319].Y = 502.703

	pointsOfFunctionPlot[1_320].X = 6.23
	pointsOfFunctionPlot[1_320].Y = 507.755

	pointsOfFunctionPlot[1_321].X = 6.24
	pointsOfFunctionPlot[1_321].Y = 512.858

	pointsOfFunctionPlot[1_322].X = 6.25
	pointsOfFunctionPlot[1_322].Y = 518.012

	pointsOfFunctionPlot[1_323].X = 6.26
	pointsOfFunctionPlot[1_323].Y = 523.218

	pointsOfFunctionPlot[1_324].X = 6.27
	pointsOfFunctionPlot[1_324].Y = 528.477

	pointsOfFunctionPlot[1_325].X = 6.28
	pointsOfFunctionPlot[1_325].Y = 533.788

	pointsOfFunctionPlot[1_326].X = 6.29
	pointsOfFunctionPlot[1_326].Y = 539.153

	pointsOfFunctionPlot[1_327].X = 6.3
	pointsOfFunctionPlot[1_327].Y = 544.571

	pointsOfFunctionPlot[1_328].X = 6.31
	pointsOfFunctionPlot[1_328].Y = 550.044

	pointsOfFunctionPlot[1_329].X = 6.32
	pointsOfFunctionPlot[1_329].Y = 555.573

	pointsOfFunctionPlot[1_330].X = 6.33
	pointsOfFunctionPlot[1_330].Y = 561.156

	pointsOfFunctionPlot[1_331].X = 6.34
	pointsOfFunctionPlot[1_331].Y = 566.796

	pointsOfFunctionPlot[1_332].X = 6.35
	pointsOfFunctionPlot[1_332].Y = 572.492

	pointsOfFunctionPlot[1_333].X = 6.36
	pointsOfFunctionPlot[1_333].Y = 578.246

	pointsOfFunctionPlot[1_334].X = 6.37
	pointsOfFunctionPlot[1_334].Y = 584.057

	pointsOfFunctionPlot[1_335].X = 6.38
	pointsOfFunctionPlot[1_335].Y = 589.927

	pointsOfFunctionPlot[1_336].X = 6.39
	pointsOfFunctionPlot[1_336].Y = 595.856

	pointsOfFunctionPlot[1_337].X = 6.4
	pointsOfFunctionPlot[1_337].Y = 601.845

	pointsOfFunctionPlot[1_338].X = 6.41
	pointsOfFunctionPlot[1_338].Y = 607.893

	pointsOfFunctionPlot[1_339].X = 6.42
	pointsOfFunctionPlot[1_339].Y = 614.003

	pointsOfFunctionPlot[1_340].X = 6.43
	pointsOfFunctionPlot[1_340].Y = 620.173

	pointsOfFunctionPlot[1_341].X = 6.44
	pointsOfFunctionPlot[1_341].Y = 626.406

	pointsOfFunctionPlot[1_342].X = 6.45
	pointsOfFunctionPlot[1_342].Y = 632.702

	pointsOfFunctionPlot[1_343].X = 6.46
	pointsOfFunctionPlot[1_343].Y = 639.061

	pointsOfFunctionPlot[1_344].X = 6.47
	pointsOfFunctionPlot[1_344].Y = 645.483

	pointsOfFunctionPlot[1_345].X = 6.48
	pointsOfFunctionPlot[1_345].Y = 651.97

	pointsOfFunctionPlot[1_346].X = 6.49
	pointsOfFunctionPlot[1_346].Y = 658.523

	pointsOfFunctionPlot[1_347].X = 6.5
	pointsOfFunctionPlot[1_347].Y = 665.141

	pointsOfFunctionPlot[1_348].X = 6.51
	pointsOfFunctionPlot[1_348].Y = 671.826

	pointsOfFunctionPlot[1_349].X = 6.52
	pointsOfFunctionPlot[1_349].Y = 678.578

	pointsOfFunctionPlot[1_350].X = 6.53
	pointsOfFunctionPlot[1_350].Y = 685.398

	pointsOfFunctionPlot[1_351].X = 6.54
	pointsOfFunctionPlot[1_351].Y = 692.286

	pointsOfFunctionPlot[1_352].X = 6.55
	pointsOfFunctionPlot[1_352].Y = 699.244

	pointsOfFunctionPlot[1_353].X = 6.56
	pointsOfFunctionPlot[1_353].Y = 706.271

	pointsOfFunctionPlot[1_354].X = 6.57
	pointsOfFunctionPlot[1_354].Y = 713.369

	pointsOfFunctionPlot[1_355].X = 6.58
	pointsOfFunctionPlot[1_355].Y = 720.539

	pointsOfFunctionPlot[1_356].X = 6.59
	pointsOfFunctionPlot[1_356].Y = 727.78

	pointsOfFunctionPlot[1_357].X = 6.6
	pointsOfFunctionPlot[1_357].Y = 735.095

	pointsOfFunctionPlot[1_358].X = 6.61
	pointsOfFunctionPlot[1_358].Y = 742.483

	pointsOfFunctionPlot[1_359].X = 6.62
	pointsOfFunctionPlot[1_359].Y = 749.945

	pointsOfFunctionPlot[1_360].X = 6.63
	pointsOfFunctionPlot[1_360].Y = 757.482

	pointsOfFunctionPlot[1_361].X = 6.64
	pointsOfFunctionPlot[1_361].Y = 765.095

	pointsOfFunctionPlot[1_362].X = 6.65
	pointsOfFunctionPlot[1_362].Y = 772.784

	pointsOfFunctionPlot[1_363].X = 6.66
	pointsOfFunctionPlot[1_363].Y = 780.55

	pointsOfFunctionPlot[1_364].X = 6.67
	pointsOfFunctionPlot[1_364].Y = 788.395

	pointsOfFunctionPlot[1_365].X = 6.68
	pointsOfFunctionPlot[1_365].Y = 796.319

	pointsOfFunctionPlot[1_366].X = 6.69
	pointsOfFunctionPlot[1_366].Y = 804.322

	pointsOfFunctionPlot[1_367].X = 6.7
	pointsOfFunctionPlot[1_367].Y = 812.405

	pointsOfFunctionPlot[1_368].X = 6.71
	pointsOfFunctionPlot[1_368].Y = 820.57

	pointsOfFunctionPlot[1_369].X = 6.72
	pointsOfFunctionPlot[1_369].Y = 828.817

	pointsOfFunctionPlot[1_370].X = 6.73
	pointsOfFunctionPlot[1_370].Y = 837.147

	pointsOfFunctionPlot[1_371].X = 6.74
	pointsOfFunctionPlot[1_371].Y = 845.56

	pointsOfFunctionPlot[1_372].X = 6.75
	pointsOfFunctionPlot[1_372].Y = 854.058

	pointsOfFunctionPlot[1_373].X = 6.76
	pointsOfFunctionPlot[1_373].Y = 862.642

	pointsOfFunctionPlot[1_374].X = 6.77
	pointsOfFunctionPlot[1_374].Y = 871.311

	pointsOfFunctionPlot[1_375].X = 6.78
	pointsOfFunctionPlot[1_375].Y = 880.068

	pointsOfFunctionPlot[1_376].X = 6.79
	pointsOfFunctionPlot[1_376].Y = 888.913

	pointsOfFunctionPlot[1_377].X = 6.8
	pointsOfFunctionPlot[1_377].Y = 897.847

	pointsOfFunctionPlot[1_378].X = 6.81
	pointsOfFunctionPlot[1_378].Y = 906.87

	pointsOfFunctionPlot[1_379].X = 6.82
	pointsOfFunctionPlot[1_379].Y = 915.985

	pointsOfFunctionPlot[1_380].X = 6.83
	pointsOfFunctionPlot[1_380].Y = 925.19

	pointsOfFunctionPlot[1_381].X = 6.84
	pointsOfFunctionPlot[1_381].Y = 934.489

	pointsOfFunctionPlot[1_382].X = 6.85
	pointsOfFunctionPlot[1_382].Y = 943.88

	pointsOfFunctionPlot[1_383].X = 6.86
	pointsOfFunctionPlot[1_383].Y = 953.367

	pointsOfFunctionPlot[1_384].X = 6.87
	pointsOfFunctionPlot[1_384].Y = 962.948

	pointsOfFunctionPlot[1_385].X = 6.88
	pointsOfFunctionPlot[1_385].Y = 972.626

	pointsOfFunctionPlot[1_386].X = 6.89
	pointsOfFunctionPlot[1_386].Y = 982.401

	pointsOfFunctionPlot[1_387].X = 6.9
	pointsOfFunctionPlot[1_387].Y = 992.274

	pointsOfFunctionPlot[1_388].X = 6.91
	pointsOfFunctionPlot[1_388].Y = 1_002.247

	pointsOfFunctionPlot[1_389].X = 6.92
	pointsOfFunctionPlot[1_389].Y = 1_012.32

	pointsOfFunctionPlot[1_390].X = 6.93
	pointsOfFunctionPlot[1_390].Y = 1_022.494

	pointsOfFunctionPlot[1_391].X = 6.94
	pointsOfFunctionPlot[1_391].Y = 1_032.77

	pointsOfFunctionPlot[1_392].X = 6.95
	pointsOfFunctionPlot[1_392].Y = 1_043.149

	pointsOfFunctionPlot[1_393].X = 6.96
	pointsOfFunctionPlot[1_393].Y = 1_053.633

	pointsOfFunctionPlot[1_394].X = 6.97
	pointsOfFunctionPlot[1_394].Y = 1_064.222

	pointsOfFunctionPlot[1_395].X = 6.98
	pointsOfFunctionPlot[1_395].Y = 1_074.918

	pointsOfFunctionPlot[1_396].X = 6.99
	pointsOfFunctionPlot[1_396].Y = 1_085.721

	pointsOfFunctionPlot[1_397].X = 7.0
	pointsOfFunctionPlot[1_397].Y = 1_096.633

	pointsOfFunctionPlot[1_398].X = 7.01
	pointsOfFunctionPlot[1_398].Y = 1_107.654

	pointsOfFunctionPlot[1_399].X = 7.02
	pointsOfFunctionPlot[1_399].Y = 1_118.786

	pointsOfFunctionPlot[1_400].X = 7.03
	pointsOfFunctionPlot[1_400].Y = 1_130.03

	pointsOfFunctionPlot[1_401].X = 7.04
	pointsOfFunctionPlot[1_401].Y = 1_141.387

	pointsOfFunctionPlot[1_402].X = 7.05
	pointsOfFunctionPlot[1_402].Y = 1_152.858

	pointsOfFunctionPlot[1_403].X = 7.06
	pointsOfFunctionPlot[1_403].Y = 1_164.445

	pointsOfFunctionPlot[1_404].X = 7.07
	pointsOfFunctionPlot[1_404].Y = 1_176.148

	pointsOfFunctionPlot[1_405].X = 7.08
	pointsOfFunctionPlot[1_405].Y = 1_187.968

	pointsOfFunctionPlot[1_406].X = 7.09
	pointsOfFunctionPlot[1_406].Y = 1_199.907

	pointsOfFunctionPlot[1_407].X = 7.1
	pointsOfFunctionPlot[1_407].Y = 1_211.967

	pointsOfFunctionPlot[1_408].X = 7.11
	pointsOfFunctionPlot[1_408].Y = 1_224.147

	pointsOfFunctionPlot[1_409].X = 7.12
	pointsOfFunctionPlot[1_409].Y = 1_236.45

	pointsOfFunctionPlot[1_410].X = 7.13
	pointsOfFunctionPlot[1_410].Y = 1_248.877

	pointsOfFunctionPlot[1_411].X = 7.14
	pointsOfFunctionPlot[1_411].Y = 1_261.428

	pointsOfFunctionPlot[1_412].X = 7.15
	pointsOfFunctionPlot[1_412].Y = 1_274.106

	pointsOfFunctionPlot[1_413].X = 7.16
	pointsOfFunctionPlot[1_413].Y = 1_286.91

	pointsOfFunctionPlot[1_414].X = 7.17
	pointsOfFunctionPlot[1_414].Y = 1_299.844

	pointsOfFunctionPlot[1_415].X = 7.18
	pointsOfFunctionPlot[1_415].Y = 1_312.908

	pointsOfFunctionPlot[1_416].X = 7.19
	pointsOfFunctionPlot[1_416].Y = 1_326.103

	pointsOfFunctionPlot[1_417].X = 7.2
	pointsOfFunctionPlot[1_417].Y = 1_339.43

	pointsOfFunctionPlot[1_418].X = 7.21
	pointsOfFunctionPlot[1_418].Y = 1_352.892

	pointsOfFunctionPlot[1_419].X = 7.22
	pointsOfFunctionPlot[1_419].Y = 1_366.489

	pointsOfFunctionPlot[1_420].X = 7.23
	pointsOfFunctionPlot[1_420].Y = 1_380.222

	pointsOfFunctionPlot[1_421].X = 7.24
	pointsOfFunctionPlot[1_421].Y = 1_394.094

	pointsOfFunctionPlot[1_422].X = 7.25
	pointsOfFunctionPlot[1_422].Y = 1_408.104

	pointsOfFunctionPlot[1_423].X = 7.26
	pointsOfFunctionPlot[1_423].Y = 1_422.256

	pointsOfFunctionPlot[1_424].X = 7.27
	pointsOfFunctionPlot[1_424].Y = 1_436.55

	pointsOfFunctionPlot[1_425].X = 7.28
	pointsOfFunctionPlot[1_425].Y = 1_450.988

	pointsOfFunctionPlot[1_426].X = 7.29
	pointsOfFunctionPlot[1_426].Y = 1_465.57

	pointsOfFunctionPlot[1_427].X = 7.3
	pointsOfFunctionPlot[1_427].Y = 1_480.299

	pointsOfFunctionPlot[1_428].X = 7.31
	pointsOfFunctionPlot[1_428].Y = 1_495.177

	pointsOfFunctionPlot[1_429].X = 7.32
	pointsOfFunctionPlot[1_429].Y = 1_510.204

	pointsOfFunctionPlot[1_430].X = 7.33
	pointsOfFunctionPlot[1_430].Y = 1_525.381

	pointsOfFunctionPlot[1_431].X = 7.34
	pointsOfFunctionPlot[1_431].Y = 1_540.712

	pointsOfFunctionPlot[1_432].X = 7.35
	pointsOfFunctionPlot[1_432].Y = 1_556.196

	pointsOfFunctionPlot[1_433].X = 7.36
	pointsOfFunctionPlot[1_433].Y = 1_571.836

	pointsOfFunctionPlot[1_434].X = 7.37
	pointsOfFunctionPlot[1_434].Y = 1_587.633

	pointsOfFunctionPlot[1_435].X = 7.38
	pointsOfFunctionPlot[1_435].Y = 1_603.589

	pointsOfFunctionPlot[1_436].X = 7.39
	pointsOfFunctionPlot[1_436].Y = 1_619.706

	pointsOfFunctionPlot[1_437].X = 7.4
	pointsOfFunctionPlot[1_437].Y = 1_635.984

	pointsOfFunctionPlot[1_438].X = 7.41
	pointsOfFunctionPlot[1_438].Y = 1_652.426

	pointsOfFunctionPlot[1_439].X = 7.42
	pointsOfFunctionPlot[1_439].Y = 1_669.033

	pointsOfFunctionPlot[1_440].X = 7.43
	pointsOfFunctionPlot[1_440].Y = 1_685.807

	pointsOfFunctionPlot[1_441].X = 7.44
	pointsOfFunctionPlot[1_441].Y = 1_702.75

	pointsOfFunctionPlot[1_442].X = 7.45
	pointsOfFunctionPlot[1_442].Y = 1_719.863

	pointsOfFunctionPlot[1_443].X = 7.46
	pointsOfFunctionPlot[1_443].Y = 1_737.148

	pointsOfFunctionPlot[1_444].X = 7.47
	pointsOfFunctionPlot[1_444].Y = 1_754.606

	pointsOfFunctionPlot[1_445].X = 7.48
	pointsOfFunctionPlot[1_445].Y = 1_772.24

	pointsOfFunctionPlot[1_446].X = 7.49
	pointsOfFunctionPlot[1_446].Y = 1_790.052

	pointsOfFunctionPlot[1_447].X = 7.5
	pointsOfFunctionPlot[1_447].Y = 1_808.042

	pointsOfFunctionPlot[1_448].X = 7.51
	pointsOfFunctionPlot[1_448].Y = 1_826.213

	pointsOfFunctionPlot[1_449].X = 7.52
	pointsOfFunctionPlot[1_449].Y = 1_844.567

	pointsOfFunctionPlot[1_450].X = 7.53
	pointsOfFunctionPlot[1_450].Y = 1_863.105

	pointsOfFunctionPlot[1_451].X = 7.54
	pointsOfFunctionPlot[1_451].Y = 1_881.83

	pointsOfFunctionPlot[1_452].X = 7.55
	pointsOfFunctionPlot[1_452].Y = 1_900.742

	pointsOfFunctionPlot[1_453].X = 7.56
	pointsOfFunctionPlot[1_453].Y = 1_919.845

	pointsOfFunctionPlot[1_454].X = 7.57
	pointsOfFunctionPlot[1_454].Y = 1_939.14

	pointsOfFunctionPlot[1_455].X = 7.58
	pointsOfFunctionPlot[1_455].Y = 1_958.629

	pointsOfFunctionPlot[1_456].X = 7.59
	pointsOfFunctionPlot[1_456].Y = 1_978.313

	pointsOfFunctionPlot[1_457].X = 7.6
	pointsOfFunctionPlot[1_457].Y = 1_998.195

	pointsOfFunctionPlot[1_458].X = 7.61
	pointsOfFunctionPlot[1_458].Y = 2_018.278

	pointsOfFunctionPlot[1_459].X = 7.62
	pointsOfFunctionPlot[1_459].Y = 2_038.562

	pointsOfFunctionPlot[1_460].X = 7.63
	pointsOfFunctionPlot[1_460].Y = 2_059.05

	pointsOfFunctionPlot[1_461].X = 7.64
	pointsOfFunctionPlot[1_461].Y = 2_079.743

	pointsOfFunctionPlot[1_462].X = 7.65
	pointsOfFunctionPlot[1_462].Y = 2_100.645

	pointsOfFunctionPlot[1_463].X = 7.66
	pointsOfFunctionPlot[1_463].Y = 2_121.757

	pointsOfFunctionPlot[1_464].X = 7.67
	pointsOfFunctionPlot[1_464].Y = 2_143.081

	pointsOfFunctionPlot[1_465].X = 7.68
	pointsOfFunctionPlot[1_465].Y = 2_164.619

	pointsOfFunctionPlot[1_466].X = 7.69
	pointsOfFunctionPlot[1_466].Y = 2_186.374

	pointsOfFunctionPlot[1_467].X = 7.7
	pointsOfFunctionPlot[1_467].Y = 2_208.348

	pointsOfFunctionPlot[1_468].X = 7.71
	pointsOfFunctionPlot[1_468].Y = 2_230.542

	pointsOfFunctionPlot[1_469].X = 7.72
	pointsOfFunctionPlot[1_469].Y = 2_252.959

	pointsOfFunctionPlot[1_470].X = 7.73
	pointsOfFunctionPlot[1_470].Y = 2_275.602

	pointsOfFunctionPlot[1_471].X = 7.74
	pointsOfFunctionPlot[1_471].Y = 2_298.472

	pointsOfFunctionPlot[1_472].X = 7.75
	pointsOfFunctionPlot[1_472].Y = 2_321.572

	pointsOfFunctionPlot[1_473].X = 7.76
	pointsOfFunctionPlot[1_473].Y = 2_344.904

	pointsOfFunctionPlot[1_474].X = 7.77
	pointsOfFunctionPlot[1_474].Y = 2_368.471

	pointsOfFunctionPlot[1_475].X = 7.78
	pointsOfFunctionPlot[1_475].Y = 2_392.274

	pointsOfFunctionPlot[1_476].X = 7.79
	pointsOfFunctionPlot[1_476].Y = 2_416.317

	pointsOfFunctionPlot[1_477].X = 7.8
	pointsOfFunctionPlot[1_477].Y = 2_440.602

	pointsOfFunctionPlot[1_478].X = 7.81
	pointsOfFunctionPlot[1_478].Y = 2_465.13

	pointsOfFunctionPlot[1_479].X = 7.82
	pointsOfFunctionPlot[1_479].Y = 2_489.905

	pointsOfFunctionPlot[1_480].X = 7.83
	pointsOfFunctionPlot[1_480].Y = 2_514.929

	pointsOfFunctionPlot[1_481].X = 7.84
	pointsOfFunctionPlot[1_481].Y = 2_540.204

	pointsOfFunctionPlot[1_482].X = 7.85
	pointsOfFunctionPlot[1_482].Y = 2_565.734

	pointsOfFunctionPlot[1_483].X = 7.86
	pointsOfFunctionPlot[1_483].Y = 2_591.52

	pointsOfFunctionPlot[1_484].X = 7.87
	pointsOfFunctionPlot[1_484].Y = 2_617.565

	pointsOfFunctionPlot[1_485].X = 7.88
	pointsOfFunctionPlot[1_485].Y = 2_643.872

	pointsOfFunctionPlot[1_486].X = 7.89
	pointsOfFunctionPlot[1_486].Y = 2_670.443

	pointsOfFunctionPlot[1_487].X = 7.9
	pointsOfFunctionPlot[1_487].Y = 2_697.282

	pointsOfFunctionPlot[1_488].X = 7.91
	pointsOfFunctionPlot[1_488].Y = 2_724.39

	pointsOfFunctionPlot[1_489].X = 7.92
	pointsOfFunctionPlot[1_489].Y = 2_751.771

	pointsOfFunctionPlot[1_490].X = 7.93
	pointsOfFunctionPlot[1_490].Y = 2_779.426

	pointsOfFunctionPlot[1_491].X = 7.94
	pointsOfFunctionPlot[1_491].Y = 2_807.36

	pointsOfFunctionPlot[1_492].X = 7.95
	pointsOfFunctionPlot[1_492].Y = 2_835.575

	pointsOfFunctionPlot[1_493].X = 7.96
	pointsOfFunctionPlot[1_493].Y = 2_864.073

	pointsOfFunctionPlot[1_494].X = 7.97
	pointsOfFunctionPlot[1_494].Y = 2_892.857

	pointsOfFunctionPlot[1_495].X = 7.98
	pointsOfFunctionPlot[1_495].Y = 2_921.931

	pointsOfFunctionPlot[1_496].X = 7.99
	pointsOfFunctionPlot[1_496].Y = 2_951.297

	pointsOfFunctionPlot[1_497].X = 8.0
	pointsOfFunctionPlot[1_497].Y = 2_980.958

	pointsOfFunctionPlot[1_498].X = 8.01
	pointsOfFunctionPlot[1_498].Y = 3_010.917

	pointsOfFunctionPlot[1_499].X = 8.02
	pointsOfFunctionPlot[1_499].Y = 3_041.177

	pointsOfFunctionPlot[1_500].X = 8.03
	pointsOfFunctionPlot[1_500].Y = 3_071.741

	pointsOfFunctionPlot[1_501].X = 8.04
	pointsOfFunctionPlot[1_501].Y = 3_102.613

	pointsOfFunctionPlot[1_502].X = 8.05
	pointsOfFunctionPlot[1_502].Y = 3_133.795

	pointsOfFunctionPlot[1_503].X = 8.06
	pointsOfFunctionPlot[1_503].Y = 3_165.29

	pointsOfFunctionPlot[1_504].X = 8.07
	pointsOfFunctionPlot[1_504].Y = 3_197.101

	pointsOfFunctionPlot[1_505].X = 8.08
	pointsOfFunctionPlot[1_505].Y = 3_229.233

	pointsOfFunctionPlot[1_506].X = 8.09
	pointsOfFunctionPlot[1_506].Y = 3_261.687

	pointsOfFunctionPlot[1_507].X = 8.1
	pointsOfFunctionPlot[1_507].Y = 3_294.468

	pointsOfFunctionPlot[1_508].X = 8.11
	pointsOfFunctionPlot[1_508].Y = 3_327.578

	pointsOfFunctionPlot[1_509].X = 8.12
	pointsOfFunctionPlot[1_509].Y = 3_361.02

	pointsOfFunctionPlot[1_510].X = 8.13
	pointsOfFunctionPlot[1_510].Y = 3_394.799

	pointsOfFunctionPlot[1_511].X = 8.14
	pointsOfFunctionPlot[1_511].Y = 3_428.917

	pointsOfFunctionPlot[1_512].X = 8.15
	pointsOfFunctionPlot[1_512].Y = 3_463.379

	pointsOfFunctionPlot[1_513].X = 8.16
	pointsOfFunctionPlot[1_513].Y = 3_498.186

	pointsOfFunctionPlot[1_514].X = 8.17
	pointsOfFunctionPlot[1_514].Y = 3_533.344

	pointsOfFunctionPlot[1_515].X = 8.18
	pointsOfFunctionPlot[1_515].Y = 3_568.854

	pointsOfFunctionPlot[1_516].X = 8.19
	pointsOfFunctionPlot[1_516].Y = 3_604.722

	pointsOfFunctionPlot[1_517].X = 8.2
	pointsOfFunctionPlot[1_517].Y = 3_640.95

	pointsOfFunctionPlot[1_518].X = 8.21
	pointsOfFunctionPlot[1_518].Y = 3_677.542

	pointsOfFunctionPlot[1_519].X = 8.22
	pointsOfFunctionPlot[1_519].Y = 3_714.502

	pointsOfFunctionPlot[1_520].X = 8.23
	pointsOfFunctionPlot[1_520].Y = 3_751.833

	pointsOfFunctionPlot[1_521].X = 8.24
	pointsOfFunctionPlot[1_521].Y = 3_789.54

	pointsOfFunctionPlot[1_522].X = 8.25
	pointsOfFunctionPlot[1_522].Y = 3_827.625

	pointsOfFunctionPlot[1_523].X = 8.26
	pointsOfFunctionPlot[1_523].Y = 3_866.094

	pointsOfFunctionPlot[1_524].X = 8.27
	pointsOfFunctionPlot[1_524].Y = 3_904.949

	pointsOfFunctionPlot[1_525].X = 8.28
	pointsOfFunctionPlot[1_525].Y = 3_944.194

	pointsOfFunctionPlot[1_526].X = 8.29
	pointsOfFunctionPlot[1_526].Y = 3_983.834

	pointsOfFunctionPlot[1_527].X = 8.3
	pointsOfFunctionPlot[1_527].Y = 4_023.872

	pointsOfFunctionPlot[1_528].X = 8.31
	pointsOfFunctionPlot[1_528].Y = 4_064.313

	pointsOfFunctionPlot[1_529].X = 8.32
	pointsOfFunctionPlot[1_529].Y = 4_105.16

	pointsOfFunctionPlot[1_530].X = 8.33
	pointsOfFunctionPlot[1_530].Y = 4_164.417

	pointsOfFunctionPlot[1_531].X = 8.34
	pointsOfFunctionPlot[1_531].Y = 4_188.089

	pointsOfFunctionPlot[1_532].X = 8.35
	pointsOfFunctionPlot[1_532].Y = 4_230.18

	pointsOfFunctionPlot[1_533].X = 8.36
	pointsOfFunctionPlot[1_533].Y = 4_272.694

	pointsOfFunctionPlot[1_534].X = 8.37
	pointsOfFunctionPlot[1_534].Y = 4_315.636

	pointsOfFunctionPlot[1_535].X = 8.38
	pointsOfFunctionPlot[1_535].Y = 4_359.008

	pointsOfFunctionPlot[1_536].X = 8.39
	pointsOfFunctionPlot[1_536].Y = 4_402.817

	pointsOfFunctionPlot[1_537].X = 8.4
	pointsOfFunctionPlot[1_537].Y = 4_447.066

	pointsOfFunctionPlot[1_538].X = 8.41
	pointsOfFunctionPlot[1_538].Y = 4_491.76

	pointsOfFunctionPlot[1_539].X = 8.42
	pointsOfFunctionPlot[1_539].Y = 4_536.903

	pointsOfFunctionPlot[1_540].X = 8.43
	pointsOfFunctionPlot[1_540].Y = 4_582.5

	pointsOfFunctionPlot[1_541].X = 8.44
	pointsOfFunctionPlot[1_541].Y = 4_628.555

	pointsOfFunctionPlot[1_542].X = 8.45
	pointsOfFunctionPlot[1_542].Y = 4_675.072

	pointsOfFunctionPlot[1_543].X = 8.46
	pointsOfFunctionPlot[1_543].Y = 4_722.058

	pointsOfFunctionPlot[1_544].X = 8.47
	pointsOfFunctionPlot[1_544].Y = 4_769.515

	pointsOfFunctionPlot[1_545].X = 8.48
	pointsOfFunctionPlot[1_545].Y = 4_817.449

	pointsOfFunctionPlot[1_546].X = 8.49
	pointsOfFunctionPlot[1_546].Y = 4_865.866

	pointsOfFunctionPlot[1_547].X = 8.5
	pointsOfFunctionPlot[1_547].Y = 4_914.768

	pointsOfFunctionPlot[1_548].X = 8.51
	pointsOfFunctionPlot[1_548].Y = 4_914.768

	pointsOfFunctionPlot[1_549].X = 8.52
	pointsOfFunctionPlot[1_549].Y = 4_964.163

	pointsOfFunctionPlot[1_550].X = 8.53
	pointsOfFunctionPlot[1_550].Y = 5_014.053

	pointsOfFunctionPlot[1_551].X = 8.54
	pointsOfFunctionPlot[1_551].Y = 5_115.344

	pointsOfFunctionPlot[1_552].X = 8.55
	pointsOfFunctionPlot[1_552].Y = 5_166.754

	pointsOfFunctionPlot[1_553].X = 8.56
	pointsOfFunctionPlot[1_553].Y = 5_218.681

	pointsOfFunctionPlot[1_554].X = 8.57
	pointsOfFunctionPlot[1_554].Y = 5_271.129

	pointsOfFunctionPlot[1_555].X = 8.58
	pointsOfFunctionPlot[1_555].Y = 5_324.105

	pointsOfFunctionPlot[1_556].X = 8.59
	pointsOfFunctionPlot[1_556].Y = 5_377.613

	pointsOfFunctionPlot[1_557].X = 8.6
	pointsOfFunctionPlot[1_557].Y = 5_431.386

	pointsOfFunctionPlot[1_558].X = 8.61
	pointsOfFunctionPlot[1_558].Y = 5_486.248

	pointsOfFunctionPlot[1_559].X = 8.62
	pointsOfFunctionPlot[1_559].Y = 5_541.386

	pointsOfFunctionPlot[1_560].X = 8.63
	pointsOfFunctionPlot[1_560].Y = 5_597.078

	pointsOfFunctionPlot[1_561].X = 8.64
	pointsOfFunctionPlot[1_561].Y = 5_653.329

	pointsOfFunctionPlot[1_562].X = 8.65
	pointsOfFunctionPlot[1_562].Y = 5_710.146

	pointsOfFunctionPlot[1_563].X = 8.66
	pointsOfFunctionPlot[1_563].Y = 5_767.534

	pointsOfFunctionPlot[1_564].X = 8.67
	pointsOfFunctionPlot[1_564].Y = 5_825.499

	pointsOfFunctionPlot[1_565].X = 8.68
	pointsOfFunctionPlot[1_565].Y = 5_884.046

	pointsOfFunctionPlot[1_566].X = 8.69
	pointsOfFunctionPlot[1_566].Y = 5_943.182

	pointsOfFunctionPlot[1_567].X = 8.7
	pointsOfFunctionPlot[1_567].Y = 6_002.912

	pointsOfFunctionPlot[1_568].X = 8.71
	pointsOfFunctionPlot[1_568].Y = 6_063.242

	pointsOfFunctionPlot[1_569].X = 8.72
	pointsOfFunctionPlot[1_569].Y = 6_124.179

	pointsOfFunctionPlot[1_570].X = 8.73
	pointsOfFunctionPlot[1_570].Y = 6_185.728

	pointsOfFunctionPlot[1_571].X = 8.74
	pointsOfFunctionPlot[1_571].Y = 6_247.895

	pointsOfFunctionPlot[1_572].X = 8.75
	pointsOfFunctionPlot[1_572].Y = 6_310.688

	pointsOfFunctionPlot[1_573].X = 8.76
	pointsOfFunctionPlot[1_573].Y = 6_374.111

	pointsOfFunctionPlot[1_574].X = 8.77
	pointsOfFunctionPlot[1_574].Y = 6_438.172

	pointsOfFunctionPlot[1_575].X = 8.78
	pointsOfFunctionPlot[1_575].Y = 6_502.877

	pointsOfFunctionPlot[1_576].X = 8.79
	pointsOfFunctionPlot[1_576].Y = 6_568.232

	pointsOfFunctionPlot[1_577].X = 8.8
	pointsOfFunctionPlot[1_577].Y = 6_634.244

	pointsOfFunctionPlot[1_578].X = 8.81
	pointsOfFunctionPlot[1_578].Y = 6_700.919

	pointsOfFunctionPlot[1_579].X = 8.82
	pointsOfFunctionPlot[1_579].Y = 6_768.264

	pointsOfFunctionPlot[1_580].X = 8.83
	pointsOfFunctionPlot[1_580].Y = 6_836.286

	pointsOfFunctionPlot[1_581].X = 8.84
	pointsOfFunctionPlot[1_581].Y = 6_904.992

	pointsOfFunctionPlot[1_582].X = 8.85
	pointsOfFunctionPlot[1_582].Y = 6_974.389

	pointsOfFunctionPlot[1_583].X = 8.86
	pointsOfFunctionPlot[1_583].Y = 7_044.482

	pointsOfFunctionPlot[1_584].X = 8.87
	pointsOfFunctionPlot[1_584].Y = 7_115.281

	pointsOfFunctionPlot[1_585].X = 8.88
	pointsOfFunctionPlot[1_585].Y = 7_186.79

	pointsOfFunctionPlot[1_586].X = 8.89
	pointsOfFunctionPlot[1_586].Y = 7_259.019

	pointsOfFunctionPlot[1_587].X = 8.9
	pointsOfFunctionPlot[1_587].Y = 7_331.973

	pointsOfFunctionPlot[1_588].X = 8.91
	pointsOfFunctionPlot[1_588].Y = 7_405.661

	pointsOfFunctionPlot[1_589].X = 8.92
	pointsOfFunctionPlot[1_589].Y = 7_480.089

	pointsOfFunctionPlot[1_590].X = 8.93
	pointsOfFunctionPlot[1_590].Y = 7_555.265

	pointsOfFunctionPlot[1_591].X = 8.94
	pointsOfFunctionPlot[1_591].Y = 7_631.197

	pointsOfFunctionPlot[1_592].X = 8.95
	pointsOfFunctionPlot[1_592].Y = 7_707.891

	pointsOfFunctionPlot[1_593].X = 8.96
	pointsOfFunctionPlot[1_593].Y = 7_785.357

	pointsOfFunctionPlot[1_594].X = 8.97
	pointsOfFunctionPlot[1_594].Y = 7_863.601

	pointsOfFunctionPlot[1_595].X = 8.98
	pointsOfFunctionPlot[1_595].Y = 7_942.632

	pointsOfFunctionPlot[1_596].X = 8.99
	pointsOfFunctionPlot[1_596].Y = 8_022.456

	pointsOfFunctionPlot[1_597].X = 9.0
	pointsOfFunctionPlot[1_597].Y = 8_103.083

	pointsOfFunctionPlot[1_598].X = 9.01
	pointsOfFunctionPlot[1_598].Y = 8_184.521

	pointsOfFunctionPlot[1_599].X = 9.02
	pointsOfFunctionPlot[1_599].Y = 8_266.777

	pointsOfFunctionPlot[1_600].X = 9.03
	pointsOfFunctionPlot[1_600].Y = 8_349.859

	pointsOfFunctionPlot[1_601].X = 9.04
	pointsOfFunctionPlot[1_601].Y = 8_433.777

	pointsOfFunctionPlot[1_602].X = 9.05
	pointsOfFunctionPlot[1_602].Y = 8_518.537

	pointsOfFunctionPlot[1_603].X = 9.06
	pointsOfFunctionPlot[1_603].Y = 8_604.15

	pointsOfFunctionPlot[1_604].X = 9.07
	pointsOfFunctionPlot[1_604].Y = 8_690.623

	pointsOfFunctionPlot[1_605].X = 9.08
	pointsOfFunctionPlot[1_605].Y = 8_777.966

	pointsOfFunctionPlot[1_606].X = 9.09
	pointsOfFunctionPlot[1_606].Y = 8_866.186

	pointsOfFunctionPlot[1_607].X = 9.1
	pointsOfFunctionPlot[1_607].Y = 8_955.292

	pointsOfFunctionPlot[1_608].X = 9.11
	pointsOfFunctionPlot[1_608].Y = 9_045.294

	pointsOfFunctionPlot[1_609].X = 9.12
	pointsOfFunctionPlot[1_609].Y = 9_136.201

	pointsOfFunctionPlot[1_610].X = 9.13
	pointsOfFunctionPlot[1_610].Y = 9_228.022

	pointsOfFunctionPlot[1_611].X = 9.14
	pointsOfFunctionPlot[1_611].Y = 9_320.765

	pointsOfFunctionPlot[1_612].X = 9.15
	pointsOfFunctionPlot[1_612].Y = 9_414.44

	pointsOfFunctionPlot[1_613].X = 9.16
	pointsOfFunctionPlot[1_613].Y = 9_509.057

	pointsOfFunctionPlot[1_614].X = 9.17
	pointsOfFunctionPlot[1_614].Y = 9_604.624

	pointsOfFunctionPlot[1_615].X = 9.18
	pointsOfFunctionPlot[1_615].Y = 9_701.152

	pointsOfFunctionPlot[1_616].X = 9.19
	pointsOfFunctionPlot[1_616].Y = 9_798.651

	pointsOfFunctionPlot[1_617].X = 9.2
	pointsOfFunctionPlot[1_617].Y = 9_897.129

	pointsOfFunctionPlot[1_618].X = 9.21
	pointsOfFunctionPlot[1_618].Y = 9_996.596

	pointsOfFunctionPlot[1_619].X = 9.22
	pointsOfFunctionPlot[1_619].Y = 10_097.064

	pointsOfFunctionPlot[1_620].X = 9.23
	pointsOfFunctionPlot[1_620].Y = 10_198.541

	pointsOfFunctionPlot[1_621].X = 9.24
	pointsOfFunctionPlot[1_621].Y = 10_301.038

	pointsOfFunctionPlot[1_622].X = 9.25
	pointsOfFunctionPlot[1_622].Y = 10_404.565

	pointsOfFunctionPlot[1_623].X = 9.26
	pointsOfFunctionPlot[1_623].Y = 10_509.133

	pointsOfFunctionPlot[1_624].X = 9.27
	pointsOfFunctionPlot[1_624].Y = 10_614.751

	pointsOfFunctionPlot[1_625].X = 9.28
	pointsOfFunctionPlot[1_625].Y = 10_721.431

	pointsOfFunctionPlot[1_626].X = 9.29
	pointsOfFunctionPlot[1_626].Y = 10_829.184

	pointsOfFunctionPlot[1_627].X = 9.3
	pointsOfFunctionPlot[1_627].Y = 10_938.019

	pointsOfFunctionPlot[1_628].X = 9.31
	pointsOfFunctionPlot[1_628].Y = 11_047.948

	pointsOfFunctionPlot[1_629].X = 9.32
	pointsOfFunctionPlot[1_629].Y = 11_158.981

	pointsOfFunctionPlot[1_630].X = 9.33
	pointsOfFunctionPlot[1_630].Y = 11_271.131

	pointsOfFunctionPlot[1_631].X = 9.34
	pointsOfFunctionPlot[1_631].Y = 11_384.408

	pointsOfFunctionPlot[1_632].X = 9.35
	pointsOfFunctionPlot[1_632].Y = 11_498.823

	pointsOfFunctionPlot[1_633].X = 9.36
	pointsOfFunctionPlot[1_633].Y = 11_614.388

	pointsOfFunctionPlot[1_634].X = 9.37
	pointsOfFunctionPlot[1_634].Y = 11_731.115

	pointsOfFunctionPlot[1_635].X = 9.38
	pointsOfFunctionPlot[1_635].Y = 11_849.014

	pointsOfFunctionPlot[1_636].X = 9.39
	pointsOfFunctionPlot[1_636].Y = 11_968.099

	pointsOfFunctionPlot[1_637].X = 9.4
	pointsOfFunctionPlot[1_637].Y = 12_088.38

	pointsOfFunctionPlot[1_638].X = 9.41
	pointsOfFunctionPlot[1_638].Y = 12_209.871

	pointsOfFunctionPlot[1_639].X = 9.42
	pointsOfFunctionPlot[1_639].Y = 12_332.582

	pointsOfFunctionPlot[1_640].X = 9.43
	pointsOfFunctionPlot[1_640].Y = 12_456.526

	pointsOfFunctionPlot[1_641].X = 9.44
	pointsOfFunctionPlot[1_641].Y = 12_581.716

	pointsOfFunctionPlot[1_642].X = 9.45
	pointsOfFunctionPlot[1_642].Y = 12_708.165

	pointsOfFunctionPlot[1_643].X = 9.46
	pointsOfFunctionPlot[1_643].Y = 12_835.884

	pointsOfFunctionPlot[1_644].X = 9.47
	pointsOfFunctionPlot[1_644].Y = 12_964.887

	pointsOfFunctionPlot[1_645].X = 9.48
	pointsOfFunctionPlot[1_645].Y = 13_095.186

	pointsOfFunctionPlot[1_646].X = 9.49
	pointsOfFunctionPlot[1_646].Y = 13_226.795

	pointsOfFunctionPlot[1_647].X = 9.5
	pointsOfFunctionPlot[1_647].Y = 13_359.726

	pointsOfFunctionPlot[1_648].X = 9.51
	pointsOfFunctionPlot[1_648].Y = 13_493.994

	pointsOfFunctionPlot[1_649].X = 9.52
	pointsOfFunctionPlot[1_649].Y = 13_629.611

	pointsOfFunctionPlot[1_650].X = 9.53
	pointsOfFunctionPlot[1_650].Y = 13_766.591

	pointsOfFunctionPlot[1_651].X = 9.54
	pointsOfFunctionPlot[1_651].Y = 13_904.947

	pointsOfFunctionPlot[1_652].X = 9.55
	pointsOfFunctionPlot[1_652].Y = 14_044.694

	pointsOfFunctionPlot[1_653].X = 9.56
	pointsOfFunctionPlot[1_653].Y = 14_185.846

	pointsOfFunctionPlot[1_654].X = 9.57
	pointsOfFunctionPlot[1_654].Y = 14_328.416

	pointsOfFunctionPlot[1_655].X = 9.58
	pointsOfFunctionPlot[1_655].Y = 14_472.419

	pointsOfFunctionPlot[1_656].X = 9.59
	pointsOfFunctionPlot[1_656].Y = 14_617.869

	pointsOfFunctionPlot[1_657].X = 9.6
	pointsOfFunctionPlot[1_657].Y = 14_764.781

	pointsOfFunctionPlot[1_658].X = 9.61
	pointsOfFunctionPlot[1_658].Y = 14_913.17

	pointsOfFunctionPlot[1_659].X = 9.62
	pointsOfFunctionPlot[1_659].Y = 15_063.049

	pointsOfFunctionPlot[1_660].X = 9.63
	pointsOfFunctionPlot[1_660].Y = 15_214.436

	pointsOfFunctionPlot[1_661].X = 9.64
	pointsOfFunctionPlot[1_661].Y = 15_367.343

	pointsOfFunctionPlot[1_662].X = 9.65
	pointsOfFunctionPlot[1_662].Y = 15_521.788

	pointsOfFunctionPlot[1_663].X = 9.66
	pointsOfFunctionPlot[1_663].Y = 15_677.784

	pointsOfFunctionPlot[1_664].X = 9.67
	pointsOfFunctionPlot[1_664].Y = 15_835.349

	pointsOfFunctionPlot[1_665].X = 9.68
	pointsOfFunctionPlot[1_665].Y = 15_994.496

	pointsOfFunctionPlot[1_666].X = 9.69
	pointsOfFunctionPlot[1_666].Y = 16_155.244

	pointsOfFunctionPlot[1_667].X = 9.7
	pointsOfFunctionPlot[1_667].Y = 16_317.607

	pointsOfFunctionPlot[1_668].X = 9.71
	pointsOfFunctionPlot[1_668].Y = 16_481.601

	pointsOfFunctionPlot[1_669].X = 9.72
	pointsOfFunctionPlot[1_669].Y = 16_647.244

	pointsOfFunctionPlot[1_670].X = 9.73
	pointsOfFunctionPlot[1_670].Y = 16_814.552

	pointsOfFunctionPlot[1_671].X = 9.74
	pointsOfFunctionPlot[1_671].Y = 16_983.541

	pointsOfFunctionPlot[1_672].X = 9.75
	pointsOfFunctionPlot[1_672].Y = 17_154.228

	pointsOfFunctionPlot[1_673].X = 9.76
	pointsOfFunctionPlot[1_673].Y = 17_326.631

	pointsOfFunctionPlot[1_674].X = 9.77
	pointsOfFunctionPlot[1_674].Y = 17_500.767

	pointsOfFunctionPlot[1_675].X = 9.78
	pointsOfFunctionPlot[1_675].Y = 17_676.652

	pointsOfFunctionPlot[1_676].X = 9.79
	pointsOfFunctionPlot[1_676].Y = 17_854.306

	pointsOfFunctionPlot[1_677].X = 9.8
	pointsOfFunctionPlot[1_677].Y = 18_033.744

	pointsOfFunctionPlot[1_678].X = 9.81
	pointsOfFunctionPlot[1_678].Y = 18_214.987

	pointsOfFunctionPlot[1_679].X = 9.82
	pointsOfFunctionPlot[1_679].Y = 18_398.05

	pointsOfFunctionPlot[1_680].X = 9.83
	pointsOfFunctionPlot[1_680].Y = 18_582.954

	pointsOfFunctionPlot[1_681].X = 9.84
	pointsOfFunctionPlot[1_681].Y = 18_769.716

	pointsOfFunctionPlot[1_682].X = 9.85
	pointsOfFunctionPlot[1_682].Y = 18_958.354

	pointsOfFunctionPlot[1_683].X = 9.86
	pointsOfFunctionPlot[1_683].Y = 19_148.889

	pointsOfFunctionPlot[1_684].X = 9.87
	pointsOfFunctionPlot[1_684].Y = 19_341.339

	pointsOfFunctionPlot[1_685].X = 9.88
	pointsOfFunctionPlot[1_685].Y = 19_535.722

	pointsOfFunctionPlot[1_686].X = 9.89
	pointsOfFunctionPlot[1_686].Y = 19_732.059

	pointsOfFunctionPlot[1_687].X = 9.9
	pointsOfFunctionPlot[1_687].Y = 19_930.37

	pointsOfFunctionPlot[1_688].X = 9.91
	pointsOfFunctionPlot[1_688].Y = 20_130.674

	pointsOfFunctionPlot[1_689].X = 9.92
	pointsOfFunctionPlot[1_689].Y = 20_332.99

	pointsOfFunctionPlot[1_690].X = 9.93
	pointsOfFunctionPlot[1_690].Y = 20_537.34

	pointsOfFunctionPlot[1_691].X = 9.94
	pointsOfFunctionPlot[1_691].Y = 20_743.744

	pointsOfFunctionPlot[1_692].X = 9.95
	pointsOfFunctionPlot[1_692].Y = 20_952.222

	pointsOfFunctionPlot[1_693].X = 9.96
	pointsOfFunctionPlot[1_693].Y = 21_162.795

	pointsOfFunctionPlot[1_694].X = 9.97
	pointsOfFunctionPlot[1_694].Y = 21_375.485

	pointsOfFunctionPlot[1_695].X = 9.98
	pointsOfFunctionPlot[1_695].Y = 21_590.312

	pointsOfFunctionPlot[1_696].X = 9.99
	pointsOfFunctionPlot[1_696].Y = 21_807.298

	pointsOfFunctionPlot[1_697].X = 10.0
	pointsOfFunctionPlot[1_697].Y = 22_247.835










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function exp(x)"

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
		"exp-plot-01.png"); err != nil {

		panic(err)
	}
}
