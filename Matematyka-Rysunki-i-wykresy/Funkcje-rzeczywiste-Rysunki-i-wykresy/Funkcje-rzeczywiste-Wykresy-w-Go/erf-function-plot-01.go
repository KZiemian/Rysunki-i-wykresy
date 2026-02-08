package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function erf(x) (error function, integral of Gauss
	// function).

	pointsOfFunctionPlot := make(plotter.XYs, 1_189)

	pointsOfFunctionPlot[0].X = -10.0
	pointsOfFunctionPlot[0].Y = -1.0

	pointsOfFunctionPlot[1].X = -5.93
	pointsOfFunctionPlot[1].Y = -1.0

	pointsOfFunctionPlot[2].X = -5.92
	pointsOfFunctionPlot[2].Y = -0.999

	pointsOfFunctionPlot[3].X = -5.91
	pointsOfFunctionPlot[3].Y = -0.999

	pointsOfFunctionPlot[4].X = -5.90
	pointsOfFunctionPlot[4].Y = -0.999

	pointsOfFunctionPlot[5].X = -5.89
	pointsOfFunctionPlot[5].Y = -0.999

	pointsOfFunctionPlot[6].X = -5.88
	pointsOfFunctionPlot[6].Y = -0.999

	pointsOfFunctionPlot[7].X = -5.87
	pointsOfFunctionPlot[7].Y = -0.999

	pointsOfFunctionPlot[8].X = -5.86
	pointsOfFunctionPlot[8].Y = -0.999

	pointsOfFunctionPlot[9].X = -5.85
	pointsOfFunctionPlot[9].Y = -0.999

	pointsOfFunctionPlot[10].X = -5.84
	pointsOfFunctionPlot[10].Y = -0.999

	pointsOfFunctionPlot[11].X = -5.83
	pointsOfFunctionPlot[11].Y = -0.999

	pointsOfFunctionPlot[12].X = -5.82
	pointsOfFunctionPlot[12].Y = -0.999

	pointsOfFunctionPlot[13].X = -5.81
	pointsOfFunctionPlot[13].Y = -0.999

	pointsOfFunctionPlot[14].X = -5.80
	pointsOfFunctionPlot[14].Y = -0.999

	pointsOfFunctionPlot[15].X = -5.79
	pointsOfFunctionPlot[15].Y = -0.999

	pointsOfFunctionPlot[16].X = -5.78
	pointsOfFunctionPlot[16].Y = -0.999

	pointsOfFunctionPlot[17].X = -5.77
	pointsOfFunctionPlot[17].Y = -0.999

	pointsOfFunctionPlot[18].X = -5.76
	pointsOfFunctionPlot[18].Y = -0.999

	pointsOfFunctionPlot[19].X = -5.75
	pointsOfFunctionPlot[19].Y = -0.999

	pointsOfFunctionPlot[20].X = -5.74
	pointsOfFunctionPlot[20].Y = -0.999

	pointsOfFunctionPlot[21].X = -5.73
	pointsOfFunctionPlot[21].Y = -0.999

	pointsOfFunctionPlot[22].X = -5.72
	pointsOfFunctionPlot[22].Y = -0.999

	pointsOfFunctionPlot[23].X = -5.71
	pointsOfFunctionPlot[23].Y = -0.999

	pointsOfFunctionPlot[24].X = -5.70
	pointsOfFunctionPlot[24].Y = -0.999

	pointsOfFunctionPlot[25].X = -5.69
	pointsOfFunctionPlot[25].Y = -0.999

	pointsOfFunctionPlot[26].X = -5.68
	pointsOfFunctionPlot[26].Y = -0.999

	pointsOfFunctionPlot[27].X = -5.67
	pointsOfFunctionPlot[27].Y = -0.999

	pointsOfFunctionPlot[28].X = -5.66
	pointsOfFunctionPlot[28].Y = -0.999

	pointsOfFunctionPlot[29].X = -5.65
	pointsOfFunctionPlot[29].Y = -0.999

	pointsOfFunctionPlot[30].X = -5.64
	pointsOfFunctionPlot[30].Y = -0.999

	pointsOfFunctionPlot[31].X = -5.63
	pointsOfFunctionPlot[31].Y = -0.999

	pointsOfFunctionPlot[32].X = -5.62
	pointsOfFunctionPlot[32].Y = -0.999

	pointsOfFunctionPlot[33].X = -5.61
	pointsOfFunctionPlot[33].Y = -0.999

	pointsOfFunctionPlot[34].X = -5.60
	pointsOfFunctionPlot[34].Y = -0.999

	pointsOfFunctionPlot[35].X = -5.59
	pointsOfFunctionPlot[35].Y = -0.999

	pointsOfFunctionPlot[36].X = -5.58
	pointsOfFunctionPlot[36].Y = -0.999

	pointsOfFunctionPlot[37].X = -5.57
	pointsOfFunctionPlot[37].Y = -0.999

	pointsOfFunctionPlot[38].X = -5.56
	pointsOfFunctionPlot[38].Y = -0.999

	pointsOfFunctionPlot[39].X = -5.55
	pointsOfFunctionPlot[39].Y = -0.999

	pointsOfFunctionPlot[40].X = -5.54
	pointsOfFunctionPlot[40].Y = -0.999

	pointsOfFunctionPlot[41].X = -5.53
	pointsOfFunctionPlot[41].Y = -0.999

	pointsOfFunctionPlot[42].X = -5.52
	pointsOfFunctionPlot[42].Y = -0.999

	pointsOfFunctionPlot[43].X = -5.51
	pointsOfFunctionPlot[43].Y = -0.999

	pointsOfFunctionPlot[44].X = -5.50
	pointsOfFunctionPlot[44].Y = -0.999

	pointsOfFunctionPlot[45].X = -5.49
	pointsOfFunctionPlot[45].Y = -0.999

	pointsOfFunctionPlot[46].X = -5.48
	pointsOfFunctionPlot[46].Y = -0.999

	pointsOfFunctionPlot[47].X = -5.47
	pointsOfFunctionPlot[47].Y = -0.999

	pointsOfFunctionPlot[48].X = -5.46
	pointsOfFunctionPlot[48].Y = -0.999

	pointsOfFunctionPlot[49].X = -5.45
	pointsOfFunctionPlot[49].Y = -0.999

	pointsOfFunctionPlot[50].X = -5.44
	pointsOfFunctionPlot[50].Y = -0.999

	pointsOfFunctionPlot[51].X = -5.43
	pointsOfFunctionPlot[51].Y = -0.999

	pointsOfFunctionPlot[52].X = -5.42
	pointsOfFunctionPlot[52].Y = -0.999

	pointsOfFunctionPlot[53].X = -5.41
	pointsOfFunctionPlot[53].Y = -0.999

	pointsOfFunctionPlot[54].X = -5.40
	pointsOfFunctionPlot[54].Y = -0.999

	pointsOfFunctionPlot[55].X = -5.39
	pointsOfFunctionPlot[55].Y = -0.999

	pointsOfFunctionPlot[56].X = -5.38
	pointsOfFunctionPlot[56].Y = -0.999

	pointsOfFunctionPlot[57].X = -5.37
	pointsOfFunctionPlot[57].Y = -0.999

	pointsOfFunctionPlot[58].X = -5.36
	pointsOfFunctionPlot[58].Y = -0.999

	pointsOfFunctionPlot[59].X = -5.35
	pointsOfFunctionPlot[59].Y = -0.999

	pointsOfFunctionPlot[60].X = -5.34
	pointsOfFunctionPlot[60].Y = -0.999

	pointsOfFunctionPlot[61].X = -5.33
	pointsOfFunctionPlot[61].Y = -0.999

	pointsOfFunctionPlot[62].X = -5.32
	pointsOfFunctionPlot[62].Y = -0.999

	pointsOfFunctionPlot[63].X = -5.31
	pointsOfFunctionPlot[63].Y = -0.999

	pointsOfFunctionPlot[64].X = -5.30
	pointsOfFunctionPlot[64].Y = -0.999

	pointsOfFunctionPlot[65].X = -5.29
	pointsOfFunctionPlot[65].Y = -0.999

	pointsOfFunctionPlot[66].X = -5.28
	pointsOfFunctionPlot[66].Y = -0.999

	pointsOfFunctionPlot[67].X = -5.27
	pointsOfFunctionPlot[67].Y = -0.999

	pointsOfFunctionPlot[68].X = -5.26
	pointsOfFunctionPlot[68].Y = -0.999

	pointsOfFunctionPlot[69].X = -5.25
	pointsOfFunctionPlot[69].Y = -0.999

	pointsOfFunctionPlot[70].X = -5.24
	pointsOfFunctionPlot[70].Y = -0.999

	pointsOfFunctionPlot[71].X = -5.23
	pointsOfFunctionPlot[71].Y = -0.999

	pointsOfFunctionPlot[72].X = -5.22
	pointsOfFunctionPlot[72].Y = -0.999

	pointsOfFunctionPlot[73].X = -5.21
	pointsOfFunctionPlot[73].Y = -0.999

	pointsOfFunctionPlot[74].X = -5.20
	pointsOfFunctionPlot[74].Y = -0.999

	pointsOfFunctionPlot[75].X = -5.19
	pointsOfFunctionPlot[75].Y = -0.999

	pointsOfFunctionPlot[76].X = -5.18
	pointsOfFunctionPlot[76].Y = -0.999

	pointsOfFunctionPlot[77].X = -5.17
	pointsOfFunctionPlot[77].Y = -0.999

	pointsOfFunctionPlot[78].X = -5.16
	pointsOfFunctionPlot[78].Y = -0.999

	pointsOfFunctionPlot[79].X = -5.15
	pointsOfFunctionPlot[79].Y = -0.999

	pointsOfFunctionPlot[80].X = -5.14
	pointsOfFunctionPlot[80].Y = -0.999

	pointsOfFunctionPlot[81].X = -5.13
	pointsOfFunctionPlot[81].Y = -0.999

	pointsOfFunctionPlot[82].X = -5.12
	pointsOfFunctionPlot[82].Y = -0.999

	pointsOfFunctionPlot[83].X = -5.11
	pointsOfFunctionPlot[83].Y = -0.999

	pointsOfFunctionPlot[84].X = -5.10
	pointsOfFunctionPlot[84].Y = -0.999

	pointsOfFunctionPlot[85].X = -5.09
	pointsOfFunctionPlot[85].Y = -0.999

	pointsOfFunctionPlot[86].X = -5.08
	pointsOfFunctionPlot[86].Y = -0.999

	pointsOfFunctionPlot[87].X = -5.07
	pointsOfFunctionPlot[87].Y = -0.999

	pointsOfFunctionPlot[88].X = -5.06
	pointsOfFunctionPlot[88].Y = -0.999

	pointsOfFunctionPlot[89].X = -5.05
	pointsOfFunctionPlot[89].Y = -0.999

	pointsOfFunctionPlot[90].X = -5.04
	pointsOfFunctionPlot[90].Y = -0.999

	pointsOfFunctionPlot[91].X = -5.03
	pointsOfFunctionPlot[91].Y = -0.999

	pointsOfFunctionPlot[92].X = -5.02
	pointsOfFunctionPlot[92].Y = -0.999

	pointsOfFunctionPlot[93].X = -5.01
	pointsOfFunctionPlot[93].Y = -0.999

	pointsOfFunctionPlot[94].X = -5.0
	pointsOfFunctionPlot[94].Y = -0.999

	pointsOfFunctionPlot[95].X = -4.99
	pointsOfFunctionPlot[95].Y = -0.999

	pointsOfFunctionPlot[96].X = -4.98
	pointsOfFunctionPlot[96].Y = -0.999

	pointsOfFunctionPlot[97].X = -4.97
	pointsOfFunctionPlot[97].Y = -0.999

	pointsOfFunctionPlot[98].X = -4.96
	pointsOfFunctionPlot[98].Y = -0.999

	pointsOfFunctionPlot[99].X = -4.95
	pointsOfFunctionPlot[99].Y = -0.999

	pointsOfFunctionPlot[100].X = -4.94
	pointsOfFunctionPlot[100].Y = -0.999

	pointsOfFunctionPlot[101].X = -4.93
	pointsOfFunctionPlot[101].Y = -0.999

	pointsOfFunctionPlot[102].X = -4.92
	pointsOfFunctionPlot[102].Y = -0.999

	pointsOfFunctionPlot[103].X = -4.91
	pointsOfFunctionPlot[103].Y = -0.999

	pointsOfFunctionPlot[104].X = -4.90
	pointsOfFunctionPlot[104].Y = -0.999

	pointsOfFunctionPlot[105].X = -4.89
	pointsOfFunctionPlot[105].Y = -0.999

	pointsOfFunctionPlot[106].X = -4.88
	pointsOfFunctionPlot[106].Y = -0.999

	pointsOfFunctionPlot[107].X = -4.87
	pointsOfFunctionPlot[107].Y = -0.999

	pointsOfFunctionPlot[108].X = -4.86
	pointsOfFunctionPlot[108].Y = -0.999

	pointsOfFunctionPlot[109].X = -4.85
	pointsOfFunctionPlot[109].Y = -0.999

	pointsOfFunctionPlot[110].X = -4.84
	pointsOfFunctionPlot[110].Y = -0.999

	pointsOfFunctionPlot[111].X = -4.83
	pointsOfFunctionPlot[111].Y = -0.999

	pointsOfFunctionPlot[112].X = -4.82
	pointsOfFunctionPlot[112].Y = -0.999

	pointsOfFunctionPlot[113].X = -4.81
	pointsOfFunctionPlot[113].Y = -0.999

	pointsOfFunctionPlot[114].X = -4.80
	pointsOfFunctionPlot[114].Y = -0.999

	pointsOfFunctionPlot[115].X = -4.79
	pointsOfFunctionPlot[115].Y = -0.999

	pointsOfFunctionPlot[116].X = -4.78
	pointsOfFunctionPlot[116].Y = -0.999

	pointsOfFunctionPlot[117].X = -4.77
	pointsOfFunctionPlot[117].Y = -0.999

	pointsOfFunctionPlot[118].X = -4.76
	pointsOfFunctionPlot[118].Y = -0.999

	pointsOfFunctionPlot[119].X = -4.75
	pointsOfFunctionPlot[119].Y = -0.999

	pointsOfFunctionPlot[120].X = -4.74
	pointsOfFunctionPlot[120].Y = -0.999

	pointsOfFunctionPlot[121].X = -4.73
	pointsOfFunctionPlot[121].Y = -0.999

	pointsOfFunctionPlot[122].X = -4.72
	pointsOfFunctionPlot[122].Y = -0.999

	pointsOfFunctionPlot[123].X = -4.71
	pointsOfFunctionPlot[123].Y = -0.999

	pointsOfFunctionPlot[124].X = -4.70
	pointsOfFunctionPlot[124].Y = -0.999

	pointsOfFunctionPlot[125].X = -4.69
	pointsOfFunctionPlot[125].Y = -0.999

	pointsOfFunctionPlot[126].X = -4.68
	pointsOfFunctionPlot[126].Y = -0.999

	pointsOfFunctionPlot[127].X = -4.67
	pointsOfFunctionPlot[127].Y = -0.999

	pointsOfFunctionPlot[128].X = -4.66
	pointsOfFunctionPlot[128].Y = -0.999

	pointsOfFunctionPlot[129].X = -4.65
	pointsOfFunctionPlot[129].Y = -0.999

	pointsOfFunctionPlot[130].X = -4.64
	pointsOfFunctionPlot[130].Y = -0.999

	pointsOfFunctionPlot[131].X = -4.63
	pointsOfFunctionPlot[131].Y = -0.999

	pointsOfFunctionPlot[132].X = -4.62
	pointsOfFunctionPlot[132].Y = -0.999

	pointsOfFunctionPlot[133].X = -4.61
	pointsOfFunctionPlot[133].Y = -0.999

	pointsOfFunctionPlot[134].X = -4.60
	pointsOfFunctionPlot[134].Y = -0.999

	pointsOfFunctionPlot[135].X = -4.59
	pointsOfFunctionPlot[135].Y = -0.999

	pointsOfFunctionPlot[136].X = -4.58
	pointsOfFunctionPlot[136].Y = -0.999

	pointsOfFunctionPlot[137].X = -4.57
	pointsOfFunctionPlot[137].Y = -0.999

	pointsOfFunctionPlot[138].X = -4.56
	pointsOfFunctionPlot[138].Y = -0.999

	pointsOfFunctionPlot[139].X = -4.55
	pointsOfFunctionPlot[139].Y = -0.999

	pointsOfFunctionPlot[140].X = -4.54
	pointsOfFunctionPlot[140].Y = -0.999

	pointsOfFunctionPlot[141].X = -4.53
	pointsOfFunctionPlot[141].Y = -0.999

	pointsOfFunctionPlot[142].X = -4.52
	pointsOfFunctionPlot[142].Y = -0.999

	pointsOfFunctionPlot[143].X = -4.51
	pointsOfFunctionPlot[143].Y = -0.999

	pointsOfFunctionPlot[144].X = -4.50
	pointsOfFunctionPlot[144].Y = -0.999

	pointsOfFunctionPlot[145].X = -4.49
	pointsOfFunctionPlot[145].Y = -0.999

	pointsOfFunctionPlot[146].X = -4.48
	pointsOfFunctionPlot[146].Y = -0.999

	pointsOfFunctionPlot[147].X = -4.47
	pointsOfFunctionPlot[147].Y = -0.999

	pointsOfFunctionPlot[148].X = -4.46
	pointsOfFunctionPlot[148].Y = -0.999

	pointsOfFunctionPlot[149].X = -4.45
	pointsOfFunctionPlot[149].Y = -0.999

	pointsOfFunctionPlot[150].X = -4.44
	pointsOfFunctionPlot[150].Y = -0.999

	pointsOfFunctionPlot[151].X = -4.43
	pointsOfFunctionPlot[151].Y = -0.999

	pointsOfFunctionPlot[152].X = -4.42
	pointsOfFunctionPlot[152].Y = -0.999

	pointsOfFunctionPlot[153].X = -4.41
	pointsOfFunctionPlot[153].Y = -0.999

	pointsOfFunctionPlot[154].X = -4.40
	pointsOfFunctionPlot[154].Y = -0.999

	pointsOfFunctionPlot[155].X = -4.39
	pointsOfFunctionPlot[155].Y = -0.999

	pointsOfFunctionPlot[156].X = -4.38
	pointsOfFunctionPlot[156].Y = -0.999

	pointsOfFunctionPlot[157].X = -4.37
	pointsOfFunctionPlot[157].Y = -0.999

	pointsOfFunctionPlot[158].X = -4.36
	pointsOfFunctionPlot[158].Y = -0.999

	pointsOfFunctionPlot[159].X = -4.35
	pointsOfFunctionPlot[159].Y = -0.999

	pointsOfFunctionPlot[160].X = -4.34
	pointsOfFunctionPlot[160].Y = -0.999

	pointsOfFunctionPlot[161].X = -4.33
	pointsOfFunctionPlot[161].Y = -0.999

	pointsOfFunctionPlot[162].X = -4.32
	pointsOfFunctionPlot[162].Y = -0.999

	pointsOfFunctionPlot[163].X = -4.31
	pointsOfFunctionPlot[163].Y = -0.999

	pointsOfFunctionPlot[164].X = -4.30
	pointsOfFunctionPlot[164].Y = -0.999

	pointsOfFunctionPlot[165].X = -4.29
	pointsOfFunctionPlot[165].Y = -0.999

	pointsOfFunctionPlot[166].X = -4.28
	pointsOfFunctionPlot[166].Y = -0.999

	pointsOfFunctionPlot[167].X = -4.27
	pointsOfFunctionPlot[167].Y = -0.999

	pointsOfFunctionPlot[168].X = -4.26
	pointsOfFunctionPlot[168].Y = -0.999

	pointsOfFunctionPlot[169].X = -4.25
	pointsOfFunctionPlot[169].Y = -0.999

	pointsOfFunctionPlot[170].X = -4.24
	pointsOfFunctionPlot[170].Y = -0.999

	pointsOfFunctionPlot[171].X = -4.23
	pointsOfFunctionPlot[171].Y = -0.999

	pointsOfFunctionPlot[172].X = -4.22
	pointsOfFunctionPlot[172].Y = -0.999

	pointsOfFunctionPlot[173].X = -4.21
	pointsOfFunctionPlot[173].Y = -0.999

	pointsOfFunctionPlot[174].X = -4.20
	pointsOfFunctionPlot[174].Y = -0.999

	pointsOfFunctionPlot[175].X = -4.19
	pointsOfFunctionPlot[175].Y = -0.999

	pointsOfFunctionPlot[176].X = -4.18
	pointsOfFunctionPlot[176].Y = -0.999

	pointsOfFunctionPlot[177].X = -4.17
	pointsOfFunctionPlot[177].Y = -0.999

	pointsOfFunctionPlot[178].X = -4.16
	pointsOfFunctionPlot[178].Y = -0.999

	pointsOfFunctionPlot[179].X = -4.15
	pointsOfFunctionPlot[179].Y = -0.999

	pointsOfFunctionPlot[180].X = -4.14
	pointsOfFunctionPlot[180].Y = -0.999

	pointsOfFunctionPlot[181].X = -4.13
	pointsOfFunctionPlot[181].Y = -0.999

	pointsOfFunctionPlot[182].X = -4.12
	pointsOfFunctionPlot[182].Y = -0.999

	pointsOfFunctionPlot[183].X = -4.11
	pointsOfFunctionPlot[183].Y = -0.999

	pointsOfFunctionPlot[184].X = -4.10
	pointsOfFunctionPlot[184].Y = -0.999

	pointsOfFunctionPlot[185].X = -4.09
	pointsOfFunctionPlot[185].Y = -0.999

	pointsOfFunctionPlot[186].X = -4.08
	pointsOfFunctionPlot[186].Y = -0.999

	pointsOfFunctionPlot[187].X = -4.07
	pointsOfFunctionPlot[187].Y = -0.999

	pointsOfFunctionPlot[188].X = -4.06
	pointsOfFunctionPlot[188].Y = -0.999

	pointsOfFunctionPlot[189].X = -4.05
	pointsOfFunctionPlot[189].Y = -0.999

	pointsOfFunctionPlot[190].X = -4.04
	pointsOfFunctionPlot[190].Y = -0.999

	pointsOfFunctionPlot[191].X = -4.03
	pointsOfFunctionPlot[191].Y = -0.999

	pointsOfFunctionPlot[192].X = -4.02
	pointsOfFunctionPlot[192].Y = -0.999

	pointsOfFunctionPlot[193].X = -4.01
	pointsOfFunctionPlot[193].Y = -0.999

	pointsOfFunctionPlot[194].X = -4.0
	pointsOfFunctionPlot[194].Y = -0.999

	pointsOfFunctionPlot[195].X = -3.99
	pointsOfFunctionPlot[195].Y = -0.999

	pointsOfFunctionPlot[196].X = -3.98
	pointsOfFunctionPlot[196].Y = -0.999

	pointsOfFunctionPlot[197].X = -3.97
	pointsOfFunctionPlot[197].Y = -0.999

	pointsOfFunctionPlot[198].X = -3.96
	pointsOfFunctionPlot[198].Y = -0.999

	pointsOfFunctionPlot[199].X = -3.95
	pointsOfFunctionPlot[199].Y = -0.999

	pointsOfFunctionPlot[200].X = -3.94
	pointsOfFunctionPlot[200].Y = -0.999

	pointsOfFunctionPlot[201].X = -3.93
	pointsOfFunctionPlot[201].Y = -0.999

	pointsOfFunctionPlot[202].X = -3.92
	pointsOfFunctionPlot[202].Y = -0.999

	pointsOfFunctionPlot[203].X = -3.91
	pointsOfFunctionPlot[203].Y = -0.999

	pointsOfFunctionPlot[204].X = -3.90
	pointsOfFunctionPlot[204].Y = -0.999

	pointsOfFunctionPlot[205].X = -3.89
	pointsOfFunctionPlot[205].Y = -0.999

	pointsOfFunctionPlot[206].X = -3.88
	pointsOfFunctionPlot[206].Y = -0.999

	pointsOfFunctionPlot[207].X = -3.87
	pointsOfFunctionPlot[207].Y = -0.999

	pointsOfFunctionPlot[208].X = -3.86
	pointsOfFunctionPlot[208].Y = -0.999

	pointsOfFunctionPlot[209].X = -3.85
	pointsOfFunctionPlot[209].Y = -0.999

	pointsOfFunctionPlot[210].X = -3.84
	pointsOfFunctionPlot[210].Y = -0.999

	pointsOfFunctionPlot[211].X = -3.83
	pointsOfFunctionPlot[211].Y = -0.999

	pointsOfFunctionPlot[212].X = -3.82
	pointsOfFunctionPlot[212].Y = -0.999

	pointsOfFunctionPlot[213].X = -3.81
	pointsOfFunctionPlot[213].Y = -0.999

	pointsOfFunctionPlot[214].X = -3.80
	pointsOfFunctionPlot[214].Y = -0.999

	pointsOfFunctionPlot[215].X = -3.79
	pointsOfFunctionPlot[215].Y = -0.999

	pointsOfFunctionPlot[216].X = -3.78
	pointsOfFunctionPlot[216].Y = -0.999

	pointsOfFunctionPlot[217].X = -3.77
	pointsOfFunctionPlot[217].Y = -0.999

	pointsOfFunctionPlot[218].X = -3.76
	pointsOfFunctionPlot[218].Y = -0.999

	pointsOfFunctionPlot[219].X = -3.75
	pointsOfFunctionPlot[219].Y = -0.999

	pointsOfFunctionPlot[220].X = -3.74
	pointsOfFunctionPlot[220].Y = -0.999

	pointsOfFunctionPlot[221].X = -3.73
	pointsOfFunctionPlot[221].Y = -0.999

	pointsOfFunctionPlot[222].X = -3.72
	pointsOfFunctionPlot[222].Y = -0.999

	pointsOfFunctionPlot[223].X = -3.71
	pointsOfFunctionPlot[223].Y = -0.999

	pointsOfFunctionPlot[224].X = -3.70
	pointsOfFunctionPlot[224].Y = -0.999

	pointsOfFunctionPlot[225].X = -3.69
	pointsOfFunctionPlot[225].Y = -0.999

	pointsOfFunctionPlot[226].X = -3.68
	pointsOfFunctionPlot[226].Y = -0.999

	pointsOfFunctionPlot[227].X = -3.67
	pointsOfFunctionPlot[227].Y = -0.999

	pointsOfFunctionPlot[228].X = -3.66
	pointsOfFunctionPlot[228].Y = -0.999

	pointsOfFunctionPlot[229].X = -3.65
	pointsOfFunctionPlot[229].Y = -0.999

	pointsOfFunctionPlot[230].X = -3.64
	pointsOfFunctionPlot[230].Y = -0.999

	pointsOfFunctionPlot[231].X = -3.63
	pointsOfFunctionPlot[231].Y = -0.999

	pointsOfFunctionPlot[232].X = -3.62
	pointsOfFunctionPlot[232].Y = -0.999

	pointsOfFunctionPlot[233].X = -3.61
	pointsOfFunctionPlot[233].Y = -0.999

	pointsOfFunctionPlot[234].X = -3.60
	pointsOfFunctionPlot[234].Y = -0.999

	pointsOfFunctionPlot[235].X = -3.59
	pointsOfFunctionPlot[235].Y = -0.999

	pointsOfFunctionPlot[236].X = -3.58
	pointsOfFunctionPlot[236].Y = -0.999

	pointsOfFunctionPlot[237].X = -3.57
	pointsOfFunctionPlot[237].Y = -0.999

	pointsOfFunctionPlot[238].X = -3.56
	pointsOfFunctionPlot[238].Y = -0.999

	pointsOfFunctionPlot[239].X = -3.55
	pointsOfFunctionPlot[239].Y = -0.999

	pointsOfFunctionPlot[240].X = -3.54
	pointsOfFunctionPlot[240].Y = -0.999

	pointsOfFunctionPlot[241].X = -3.53
	pointsOfFunctionPlot[241].Y = -0.999

	pointsOfFunctionPlot[242].X = -3.52
	pointsOfFunctionPlot[242].Y = -0.999

	pointsOfFunctionPlot[243].X = -3.51
	pointsOfFunctionPlot[243].Y = -0.999

	pointsOfFunctionPlot[244].X = -3.50
	pointsOfFunctionPlot[244].Y = -0.999

	pointsOfFunctionPlot[245].X = -3.49
	pointsOfFunctionPlot[245].Y = -0.999

	pointsOfFunctionPlot[246].X = -3.48
	pointsOfFunctionPlot[246].Y = -0.999

	pointsOfFunctionPlot[247].X = -3.47
	pointsOfFunctionPlot[247].Y = -0.999

	pointsOfFunctionPlot[248].X = -3.46
	pointsOfFunctionPlot[248].Y = -0.999

	pointsOfFunctionPlot[249].X = -3.45
	pointsOfFunctionPlot[249].Y = -0.999

	pointsOfFunctionPlot[250].X = -3.44
	pointsOfFunctionPlot[250].Y = -0.999

	pointsOfFunctionPlot[251].X = -3.43
	pointsOfFunctionPlot[251].Y = -0.999

	pointsOfFunctionPlot[252].X = -3.42
	pointsOfFunctionPlot[252].Y = -0.999

	pointsOfFunctionPlot[253].X = -3.41
	pointsOfFunctionPlot[253].Y = -0.999

	pointsOfFunctionPlot[254].X = -3.40
	pointsOfFunctionPlot[254].Y = -0.999

	pointsOfFunctionPlot[255].X = -3.39
	pointsOfFunctionPlot[255].Y = -0.999

	pointsOfFunctionPlot[256].X = -3.38
	pointsOfFunctionPlot[256].Y = -0.999

	pointsOfFunctionPlot[257].X = -3.37
	pointsOfFunctionPlot[257].Y = -0.999

	pointsOfFunctionPlot[258].X = -3.36
	pointsOfFunctionPlot[258].Y = -0.999

	pointsOfFunctionPlot[259].X = -3.35
	pointsOfFunctionPlot[259].Y = -0.999

	pointsOfFunctionPlot[260].X = -3.34
	pointsOfFunctionPlot[260].Y = -0.999

	pointsOfFunctionPlot[261].X = -3.33
	pointsOfFunctionPlot[261].Y = -0.999

	pointsOfFunctionPlot[262].X = -3.32
	pointsOfFunctionPlot[262].Y = -0.999

	pointsOfFunctionPlot[263].X = -3.31
	pointsOfFunctionPlot[263].Y = -0.999

	pointsOfFunctionPlot[264].X = -3.30
	pointsOfFunctionPlot[264].Y = -0.999

	pointsOfFunctionPlot[265].X = -3.29
	pointsOfFunctionPlot[265].Y = -0.999

	pointsOfFunctionPlot[266].X = -3.28
	pointsOfFunctionPlot[266].Y = -0.999

	pointsOfFunctionPlot[267].X = -3.27
	pointsOfFunctionPlot[267].Y = -0.999

	pointsOfFunctionPlot[268].X = -3.26
	pointsOfFunctionPlot[268].Y = -0.999

	pointsOfFunctionPlot[269].X = -3.25
	pointsOfFunctionPlot[269].Y = -0.999

	pointsOfFunctionPlot[270].X = -3.24
	pointsOfFunctionPlot[270].Y = -0.999

	pointsOfFunctionPlot[271].X = -3.23
	pointsOfFunctionPlot[271].Y = -0.999

	pointsOfFunctionPlot[272].X = -3.22
	pointsOfFunctionPlot[272].Y = -0.999

	pointsOfFunctionPlot[273].X = -3.21
	pointsOfFunctionPlot[273].Y = -0.999

	pointsOfFunctionPlot[274].X = -3.20
	pointsOfFunctionPlot[274].Y = -0.999

	pointsOfFunctionPlot[275].X = -3.19
	pointsOfFunctionPlot[275].Y = -0.999

	pointsOfFunctionPlot[276].X = -3.18
	pointsOfFunctionPlot[276].Y = -0.999

	pointsOfFunctionPlot[277].X = -3.17
	pointsOfFunctionPlot[277].Y = -0.999

	pointsOfFunctionPlot[278].X = -3.16
	pointsOfFunctionPlot[278].Y = -0.999

	pointsOfFunctionPlot[279].X = -3.15
	pointsOfFunctionPlot[279].Y = -0.999

	pointsOfFunctionPlot[280].X = -3.14
	pointsOfFunctionPlot[280].Y = -0.999

	pointsOfFunctionPlot[281].X = -3.13
	pointsOfFunctionPlot[281].Y = -0.999

	pointsOfFunctionPlot[282].X = -3.12
	pointsOfFunctionPlot[282].Y = -0.999

	pointsOfFunctionPlot[283].X = -3.11
	pointsOfFunctionPlot[283].Y = -0.999

	pointsOfFunctionPlot[284].X = -3.10
	pointsOfFunctionPlot[284].Y = -0.999

	pointsOfFunctionPlot[285].X = -3.09
	pointsOfFunctionPlot[285].Y = -0.999

	pointsOfFunctionPlot[286].X = -3.08
	pointsOfFunctionPlot[286].Y = -0.999

	pointsOfFunctionPlot[287].X = -3.07
	pointsOfFunctionPlot[287].Y = -0.999

	pointsOfFunctionPlot[288].X = -3.06
	pointsOfFunctionPlot[288].Y = -0.999

	pointsOfFunctionPlot[289].X = -3.05
	pointsOfFunctionPlot[289].Y = -0.999

	pointsOfFunctionPlot[290].X = -3.04
	pointsOfFunctionPlot[290].Y = -0.999

	pointsOfFunctionPlot[291].X = -3.03
	pointsOfFunctionPlot[291].Y = -0.999

	pointsOfFunctionPlot[292].X = -3.02
	pointsOfFunctionPlot[292].Y = -0.999

	pointsOfFunctionPlot[293].X = -3.01
	pointsOfFunctionPlot[293].Y = -0.999

	pointsOfFunctionPlot[294].X = -3.0
	pointsOfFunctionPlot[294].Y = -0.999

	pointsOfFunctionPlot[295].X = -2.99
	pointsOfFunctionPlot[295].Y = -0.999

	pointsOfFunctionPlot[296].X = -2.98
	pointsOfFunctionPlot[296].Y = -0.999

	pointsOfFunctionPlot[297].X = -2.97
	pointsOfFunctionPlot[297].Y = -0.999

	pointsOfFunctionPlot[298].X = -2.96
	pointsOfFunctionPlot[298].Y = -0.999

	pointsOfFunctionPlot[299].X = -2.95
	pointsOfFunctionPlot[299].Y = -0.999

	pointsOfFunctionPlot[300].X = -2.94
	pointsOfFunctionPlot[300].Y = -0.999

	pointsOfFunctionPlot[301].X = -2.93
	pointsOfFunctionPlot[301].Y = -0.999

	pointsOfFunctionPlot[302].X = -2.92
	pointsOfFunctionPlot[302].Y = -0.999

	pointsOfFunctionPlot[303].X = -2.91
	pointsOfFunctionPlot[303].Y = -0.999

	pointsOfFunctionPlot[304].X = -2.90
	pointsOfFunctionPlot[304].Y = -0.999

	pointsOfFunctionPlot[305].X = -2.89
	pointsOfFunctionPlot[305].Y = -0.999

	pointsOfFunctionPlot[306].X = -2.88
	pointsOfFunctionPlot[306].Y = -0.999

	pointsOfFunctionPlot[307].X = -2.87
	pointsOfFunctionPlot[307].Y = -0.999

	pointsOfFunctionPlot[308].X = -2.86
	pointsOfFunctionPlot[308].Y = -0.999

	pointsOfFunctionPlot[309].X = -2.85
	pointsOfFunctionPlot[309].Y = -0.999

	pointsOfFunctionPlot[310].X = -2.84
	pointsOfFunctionPlot[310].Y = -0.999

	pointsOfFunctionPlot[311].X = -2.83
	pointsOfFunctionPlot[311].Y = -0.999

	pointsOfFunctionPlot[312].X = -2.82
	pointsOfFunctionPlot[312].Y = -0.999

	pointsOfFunctionPlot[313].X = -2.81
	pointsOfFunctionPlot[313].Y = -0.999

	pointsOfFunctionPlot[314].X = -2.80
	pointsOfFunctionPlot[314].Y = -0.999

	pointsOfFunctionPlot[315].X = -2.79
	pointsOfFunctionPlot[315].Y = -0.999

	pointsOfFunctionPlot[316].X = -2.78
	pointsOfFunctionPlot[316].Y = -0.999

	pointsOfFunctionPlot[317].X = -2.77
	pointsOfFunctionPlot[317].Y = -0.999

	pointsOfFunctionPlot[318].X = -2.76
	pointsOfFunctionPlot[318].Y = -0.999

	pointsOfFunctionPlot[319].X = -2.75
	pointsOfFunctionPlot[319].Y = -0.999

	pointsOfFunctionPlot[320].X = -2.74
	pointsOfFunctionPlot[320].Y = -0.999

	pointsOfFunctionPlot[321].X = -2.73
	pointsOfFunctionPlot[321].Y = -0.999

	pointsOfFunctionPlot[322].X = -2.72
	pointsOfFunctionPlot[322].Y = -0.999

	pointsOfFunctionPlot[323].X = -2.71
	pointsOfFunctionPlot[323].Y = -0.999

	pointsOfFunctionPlot[324].X = -2.70
	pointsOfFunctionPlot[324].Y = -0.999

	pointsOfFunctionPlot[325].X = -2.69
	pointsOfFunctionPlot[325].Y = -0.999

	pointsOfFunctionPlot[326].X = -2.68
	pointsOfFunctionPlot[326].Y = -0.999

	pointsOfFunctionPlot[327].X = -2.67
	pointsOfFunctionPlot[327].Y = -0.999

	pointsOfFunctionPlot[328].X = -2.66
	pointsOfFunctionPlot[328].Y = -0.999

	pointsOfFunctionPlot[329].X = -2.65
	pointsOfFunctionPlot[329].Y = -0.999

	pointsOfFunctionPlot[330].X = -2.64
	pointsOfFunctionPlot[330].Y = -0.999

	pointsOfFunctionPlot[331].X = -2.63
	pointsOfFunctionPlot[331].Y = -0.999

	pointsOfFunctionPlot[332].X = -2.62
	pointsOfFunctionPlot[332].Y = -0.999

	pointsOfFunctionPlot[333].X = -2.61
	pointsOfFunctionPlot[333].Y = -0.999

	pointsOfFunctionPlot[334].X = -2.60
	pointsOfFunctionPlot[334].Y = -0.999

	pointsOfFunctionPlot[335].X = -2.59
	pointsOfFunctionPlot[335].Y = -0.999

	pointsOfFunctionPlot[336].X = -2.58
	pointsOfFunctionPlot[336].Y = -0.999

	pointsOfFunctionPlot[337].X = -2.57
	pointsOfFunctionPlot[337].Y = -0.999

	pointsOfFunctionPlot[338].X = -2.56
	pointsOfFunctionPlot[338].Y = -0.999

	pointsOfFunctionPlot[339].X = -2.55
	pointsOfFunctionPlot[339].Y = -0.999

	pointsOfFunctionPlot[340].X = -2.54
	pointsOfFunctionPlot[340].Y = -0.999

	pointsOfFunctionPlot[341].X = -2.53
	pointsOfFunctionPlot[341].Y = -0.999

	pointsOfFunctionPlot[342].X = -2.52
	pointsOfFunctionPlot[342].Y = -0.999

	pointsOfFunctionPlot[343].X = -2.51
	pointsOfFunctionPlot[343].Y = -0.999

	pointsOfFunctionPlot[344].X = -2.50
	pointsOfFunctionPlot[344].Y = -0.999

	pointsOfFunctionPlot[345].X = -2.49
	pointsOfFunctionPlot[345].Y = -0.999

	pointsOfFunctionPlot[346].X = -2.48
	pointsOfFunctionPlot[346].Y = -0.999

	pointsOfFunctionPlot[347].X = -2.47
	pointsOfFunctionPlot[347].Y = -0.999

	pointsOfFunctionPlot[348].X = -2.46
	pointsOfFunctionPlot[348].Y = -0.999

	pointsOfFunctionPlot[349].X = -2.45
	pointsOfFunctionPlot[349].Y = -0.999

	pointsOfFunctionPlot[350].X = -2.44
	pointsOfFunctionPlot[350].Y = -0.999

	pointsOfFunctionPlot[351].X = -2.43
	pointsOfFunctionPlot[351].Y = -0.999

	pointsOfFunctionPlot[352].X = -2.42
	pointsOfFunctionPlot[352].Y = -0.999

	pointsOfFunctionPlot[353].X = -2.41
	pointsOfFunctionPlot[353].Y = -0.999

	pointsOfFunctionPlot[354].X = -2.40
	pointsOfFunctionPlot[354].Y = -0.999

	pointsOfFunctionPlot[355].X = -2.39
	pointsOfFunctionPlot[355].Y = -0.999

	pointsOfFunctionPlot[356].X = -2.38
	pointsOfFunctionPlot[356].Y = -0.999

	pointsOfFunctionPlot[357].X = -2.37
	pointsOfFunctionPlot[357].Y = -0.999

	pointsOfFunctionPlot[358].X = -2.36
	pointsOfFunctionPlot[358].Y = -0.999

	pointsOfFunctionPlot[359].X = -2.35
	pointsOfFunctionPlot[359].Y = -0.999

	pointsOfFunctionPlot[360].X = -2.34
	pointsOfFunctionPlot[360].Y = -0.999

	pointsOfFunctionPlot[361].X = -2.33
	pointsOfFunctionPlot[361].Y = -0.999

	pointsOfFunctionPlot[362].X = -2.32
	pointsOfFunctionPlot[362].Y = -0.998

	pointsOfFunctionPlot[363].X = -2.31
	pointsOfFunctionPlot[363].Y = -0.998

	pointsOfFunctionPlot[364].X = -2.30
	pointsOfFunctionPlot[364].Y = -0.998

	pointsOfFunctionPlot[365].X = -2.29
	pointsOfFunctionPlot[365].Y = -0.998

	pointsOfFunctionPlot[366].X = -2.28
	pointsOfFunctionPlot[366].Y = -0.998

	pointsOfFunctionPlot[367].X = -2.27
	pointsOfFunctionPlot[367].Y = -0.998

	pointsOfFunctionPlot[368].X = -2.26
	pointsOfFunctionPlot[368].Y = -0.998

	pointsOfFunctionPlot[369].X = -2.25
	pointsOfFunctionPlot[369].Y = -0.998

	pointsOfFunctionPlot[370].X = -2.24
	pointsOfFunctionPlot[370].Y = -0.998

	pointsOfFunctionPlot[371].X = -2.23
	pointsOfFunctionPlot[371].Y = -0.998

	pointsOfFunctionPlot[372].X = -2.22
	pointsOfFunctionPlot[372].Y = -0.998

	pointsOfFunctionPlot[373].X = -2.21
	pointsOfFunctionPlot[373].Y = -0.998

	pointsOfFunctionPlot[374].X = -2.20
	pointsOfFunctionPlot[374].Y = -0.998

	pointsOfFunctionPlot[375].X = -2.19
	pointsOfFunctionPlot[375].Y = -0.998

	pointsOfFunctionPlot[376].X = -2.18
	pointsOfFunctionPlot[376].Y = -0.997

	pointsOfFunctionPlot[377].X = -2.17
	pointsOfFunctionPlot[377].Y = -0.997

	pointsOfFunctionPlot[378].X = -2.16
	pointsOfFunctionPlot[378].Y = -0.997

	pointsOfFunctionPlot[379].X = -2.15
	pointsOfFunctionPlot[379].Y = -0.997

	pointsOfFunctionPlot[380].X = -2.14
	pointsOfFunctionPlot[380].Y = -0.997

	pointsOfFunctionPlot[381].X = -2.13
	pointsOfFunctionPlot[381].Y = -0.997

	pointsOfFunctionPlot[382].X = -2.12
	pointsOfFunctionPlot[382].Y = -0.997

	pointsOfFunctionPlot[383].X = -2.11
	pointsOfFunctionPlot[383].Y = -0.997

	pointsOfFunctionPlot[384].X = -2.10
	pointsOfFunctionPlot[384].Y = -0.997

	pointsOfFunctionPlot[385].X = -2.09
	pointsOfFunctionPlot[385].Y = -0.996

	pointsOfFunctionPlot[386].X = -2.08
	pointsOfFunctionPlot[386].Y = -0.996

	pointsOfFunctionPlot[387].X = -2.07
	pointsOfFunctionPlot[387].Y = -0.996

	pointsOfFunctionPlot[388].X = -2.06
	pointsOfFunctionPlot[388].Y = -0.996

	pointsOfFunctionPlot[389].X = -2.05
	pointsOfFunctionPlot[389].Y = -0.996

	pointsOfFunctionPlot[390].X = -2.04
	pointsOfFunctionPlot[390].Y = -0.996

	pointsOfFunctionPlot[391].X = -2.03
	pointsOfFunctionPlot[391].Y = -0.995

	pointsOfFunctionPlot[392].X = -2.02
	pointsOfFunctionPlot[392].Y = -0.995

	pointsOfFunctionPlot[393].X = -2.01
	pointsOfFunctionPlot[393].Y = -0.995

	pointsOfFunctionPlot[394].X = -2.0
	pointsOfFunctionPlot[394].Y = -0.995

	pointsOfFunctionPlot[395].X = -1.99
	pointsOfFunctionPlot[395].Y = -0.995

	pointsOfFunctionPlot[396].X = -1.98
	pointsOfFunctionPlot[396].Y = -0.994

	pointsOfFunctionPlot[397].X = -1.97
	pointsOfFunctionPlot[397].Y = -0.994

	pointsOfFunctionPlot[398].X = -1.96
	pointsOfFunctionPlot[398].Y = -0.994

	pointsOfFunctionPlot[399].X = -1.95
	pointsOfFunctionPlot[399].Y = -0.994

	pointsOfFunctionPlot[400].X = -1.94
	pointsOfFunctionPlot[400].Y = -0.993

	pointsOfFunctionPlot[401].X = -1.93
	pointsOfFunctionPlot[401].Y = -0.993

	pointsOfFunctionPlot[402].X = -1.92
	pointsOfFunctionPlot[402].Y = -0.993

	pointsOfFunctionPlot[403].X = -1.91
	pointsOfFunctionPlot[403].Y = -0.993

	pointsOfFunctionPlot[404].X = -1.90
	pointsOfFunctionPlot[404].Y = -0.992

	pointsOfFunctionPlot[405].X = -1.89
	pointsOfFunctionPlot[405].Y = -0.992

	pointsOfFunctionPlot[406].X = -1.88
	pointsOfFunctionPlot[406].Y = -0.992

	pointsOfFunctionPlot[407].X = -1.87
	pointsOfFunctionPlot[407].Y = -0.991

	pointsOfFunctionPlot[408].X = -1.86
	pointsOfFunctionPlot[408].Y = -0.991

	pointsOfFunctionPlot[409].X = -1.85
	pointsOfFunctionPlot[409].Y = -0.991

	pointsOfFunctionPlot[410].X = -1.84
	pointsOfFunctionPlot[410].Y = -0.99

	pointsOfFunctionPlot[411].X = -1.83
	pointsOfFunctionPlot[411].Y = -0.99

	pointsOfFunctionPlot[412].X = -1.82
	pointsOfFunctionPlot[412].Y = -0.989

	pointsOfFunctionPlot[413].X = -1.81
	pointsOfFunctionPlot[413].Y = -0.989

	pointsOfFunctionPlot[414].X = -1.80
	pointsOfFunctionPlot[414].Y = -0.989

	pointsOfFunctionPlot[415].X = -1.79
	pointsOfFunctionPlot[415].Y = -0.988

	pointsOfFunctionPlot[416].X = -1.78
	pointsOfFunctionPlot[416].Y = -0.988

	pointsOfFunctionPlot[417].X = -1.77
	pointsOfFunctionPlot[417].Y = -0.987

	pointsOfFunctionPlot[418].X = -1.76
	pointsOfFunctionPlot[418].Y = -0.987

	pointsOfFunctionPlot[419].X = -1.75
	pointsOfFunctionPlot[419].Y = -0.986

	pointsOfFunctionPlot[420].X = -1.74
	pointsOfFunctionPlot[420].Y = -0.986

	pointsOfFunctionPlot[421].X = -1.73
	pointsOfFunctionPlot[421].Y = -0.985

	pointsOfFunctionPlot[422].X = -1.72
	pointsOfFunctionPlot[422].Y = -0.985

	pointsOfFunctionPlot[423].X = -1.71
	pointsOfFunctionPlot[423].Y = -0.984

	pointsOfFunctionPlot[424].X = -1.70
	pointsOfFunctionPlot[424].Y = -0.983

	pointsOfFunctionPlot[425].X = -1.69
	pointsOfFunctionPlot[425].Y = -0.983

	pointsOfFunctionPlot[426].X = -1.68
	pointsOfFunctionPlot[426].Y = -0.982

	pointsOfFunctionPlot[427].X = -1.67
	pointsOfFunctionPlot[427].Y = -0.981

	pointsOfFunctionPlot[428].X = -1.66
	pointsOfFunctionPlot[428].Y = -0.981

	pointsOfFunctionPlot[429].X = -1.65
	pointsOfFunctionPlot[429].Y = -0.98

	pointsOfFunctionPlot[430].X = -1.64
	pointsOfFunctionPlot[430].Y = -0.979

	pointsOfFunctionPlot[431].X = -1.63
	pointsOfFunctionPlot[431].Y = -0.978

	pointsOfFunctionPlot[432].X = -1.62
	pointsOfFunctionPlot[432].Y = -0.978

	pointsOfFunctionPlot[433].X = -1.61
	pointsOfFunctionPlot[433].Y = -0.977

	pointsOfFunctionPlot[434].X = -1.60
	pointsOfFunctionPlot[434].Y = -0.976

	pointsOfFunctionPlot[435].X = -1.59
	pointsOfFunctionPlot[435].Y = -0.975

	pointsOfFunctionPlot[436].X = -1.58
	pointsOfFunctionPlot[436].Y = -0.974

	pointsOfFunctionPlot[437].X = -1.57
	pointsOfFunctionPlot[437].Y = -0.973

	pointsOfFunctionPlot[438].X = -1.56
	pointsOfFunctionPlot[438].Y = -0.972

	pointsOfFunctionPlot[439].X = -1.55
	pointsOfFunctionPlot[439].Y = -0.971

	pointsOfFunctionPlot[440].X = -1.54
	pointsOfFunctionPlot[440].Y = -0.97

	pointsOfFunctionPlot[441].X = -1.53
	pointsOfFunctionPlot[441].Y = -0.969

	pointsOfFunctionPlot[442].X = -1.52
	pointsOfFunctionPlot[442].Y = -0.968

	pointsOfFunctionPlot[443].X = -1.51
	pointsOfFunctionPlot[443].Y = -0.967

	pointsOfFunctionPlot[444].X = -1.50
	pointsOfFunctionPlot[444].Y = -0.966

	pointsOfFunctionPlot[445].X = -1.49
	pointsOfFunctionPlot[445].Y = -0.964

	pointsOfFunctionPlot[446].X = -1.48
	pointsOfFunctionPlot[446].Y = -0.963

	pointsOfFunctionPlot[447].X = -1.47
	pointsOfFunctionPlot[447].Y = -0.962

	pointsOfFunctionPlot[448].X = -1.46
	pointsOfFunctionPlot[448].Y = -0.961

	pointsOfFunctionPlot[449].X = -1.45
	pointsOfFunctionPlot[449].Y = -0.959

	pointsOfFunctionPlot[450].X = -1.44
	pointsOfFunctionPlot[450].Y = -0.958

	pointsOfFunctionPlot[451].X = -1.43
	pointsOfFunctionPlot[451].Y = -0.956

	pointsOfFunctionPlot[452].X = -1.42
	pointsOfFunctionPlot[452].Y = -0.955

	pointsOfFunctionPlot[453].X = -1.41
	pointsOfFunctionPlot[453].Y = -0.953

	pointsOfFunctionPlot[454].X = -1.40
	pointsOfFunctionPlot[454].Y = -0.952

	pointsOfFunctionPlot[455].X = -1.39
	pointsOfFunctionPlot[455].Y = -0.95

	pointsOfFunctionPlot[456].X = -1.38
	pointsOfFunctionPlot[456].Y = -0.949

	pointsOfFunctionPlot[457].X = -1.37
	pointsOfFunctionPlot[457].Y = -0.947

	pointsOfFunctionPlot[458].X = -1.36
	pointsOfFunctionPlot[458].Y = -0.945

	pointsOfFunctionPlot[459].X = -1.35
	pointsOfFunctionPlot[459].Y = -0.943

	pointsOfFunctionPlot[460].X = -1.34
	pointsOfFunctionPlot[460].Y = -0.941

	pointsOfFunctionPlot[461].X = -1.33
	pointsOfFunctionPlot[461].Y = -0.94

	pointsOfFunctionPlot[462].X = -1.32
	pointsOfFunctionPlot[462].Y = -0.938

	pointsOfFunctionPlot[463].X = -1.31
	pointsOfFunctionPlot[463].Y = -0.936

	pointsOfFunctionPlot[464].X = -1.30
	pointsOfFunctionPlot[464].Y = -0.934

	pointsOfFunctionPlot[465].X = -1.29
	pointsOfFunctionPlot[465].Y = -0.931

	pointsOfFunctionPlot[466].X = -1.28
	pointsOfFunctionPlot[466].Y = -0.929

	pointsOfFunctionPlot[467].X = -1.27
	pointsOfFunctionPlot[467].Y = -0.927

	pointsOfFunctionPlot[468].X = -1.26
	pointsOfFunctionPlot[468].Y = -0.925

	pointsOfFunctionPlot[469].X = -1.25
	pointsOfFunctionPlot[469].Y = -0.922

	pointsOfFunctionPlot[470].X = -1.24
	pointsOfFunctionPlot[470].Y = -0.92

	pointsOfFunctionPlot[471].X = -1.23
	pointsOfFunctionPlot[471].Y = -0.918

	pointsOfFunctionPlot[472].X = -1.22
	pointsOfFunctionPlot[472].Y = -0.915

	pointsOfFunctionPlot[473].X = -1.21
	pointsOfFunctionPlot[473].Y = -0.912

	pointsOfFunctionPlot[474].X = -1.20
	pointsOfFunctionPlot[474].Y = -0.91

	pointsOfFunctionPlot[475].X = -1.19
	pointsOfFunctionPlot[475].Y = -0.907

	pointsOfFunctionPlot[476].X = -1.18
	pointsOfFunctionPlot[476].Y = -0.904

	pointsOfFunctionPlot[477].X = -1.17
	pointsOfFunctionPlot[477].Y = -0.902

	pointsOfFunctionPlot[478].X = -1.16
	pointsOfFunctionPlot[478].Y = -0.899

	pointsOfFunctionPlot[479].X = -1.15
	pointsOfFunctionPlot[479].Y = -0.896

	pointsOfFunctionPlot[480].X = -1.14
	pointsOfFunctionPlot[480].Y = -0.893

	pointsOfFunctionPlot[481].X = -1.13
	pointsOfFunctionPlot[481].Y = -0.889

	pointsOfFunctionPlot[482].X = -1.12
	pointsOfFunctionPlot[482].Y = -0.886

	pointsOfFunctionPlot[483].X = -1.11
	pointsOfFunctionPlot[483].Y = -0.883

	pointsOfFunctionPlot[484].X = -1.10
	pointsOfFunctionPlot[484].Y = -0.88

	pointsOfFunctionPlot[485].X = -1.09
	pointsOfFunctionPlot[485].Y = -0.876

	pointsOfFunctionPlot[486].X = -1.08
	pointsOfFunctionPlot[486].Y = -0.873

	pointsOfFunctionPlot[487].X = -1.07
	pointsOfFunctionPlot[487].Y = -0.869

	pointsOfFunctionPlot[488].X = -1.06
	pointsOfFunctionPlot[488].Y = -0.866

	pointsOfFunctionPlot[489].X = -1.05
	pointsOfFunctionPlot[489].Y = -0.862

	pointsOfFunctionPlot[490].X = -1.04
	pointsOfFunctionPlot[490].Y = -0.858

	pointsOfFunctionPlot[491].X = -1.03
	pointsOfFunctionPlot[491].Y = -0.854

	pointsOfFunctionPlot[492].X = -1.02
	pointsOfFunctionPlot[492].Y = -0.85

	pointsOfFunctionPlot[493].X = -1.01
	pointsOfFunctionPlot[493].Y = -0.846

	pointsOfFunctionPlot[494].X = -1.0
	pointsOfFunctionPlot[494].Y = -0.842

	pointsOfFunctionPlot[495].X = -0.99
	pointsOfFunctionPlot[495].Y = -0.838

	pointsOfFunctionPlot[496].X = -0.98
	pointsOfFunctionPlot[496].Y = -0.834

	pointsOfFunctionPlot[497].X = -0.97
	pointsOfFunctionPlot[497].Y = -0.829

	pointsOfFunctionPlot[498].X = -0.96
	pointsOfFunctionPlot[498].Y = -0.825

	pointsOfFunctionPlot[499].X = -0.95
	pointsOfFunctionPlot[499].Y = -0.82

	pointsOfFunctionPlot[500].X = -0.94
	pointsOfFunctionPlot[500].Y = -0.816

	pointsOfFunctionPlot[501].X = -0.93
	pointsOfFunctionPlot[501].Y = -0.811

	pointsOfFunctionPlot[502].X = -0.92
	pointsOfFunctionPlot[502].Y = -0.806

	pointsOfFunctionPlot[503].X = -0.91
	pointsOfFunctionPlot[503].Y = -0.801

	pointsOfFunctionPlot[504].X = -0.90
	pointsOfFunctionPlot[504].Y = -0.796

	pointsOfFunctionPlot[505].X = -0.89
	pointsOfFunctionPlot[505].Y = -0.791

	pointsOfFunctionPlot[506].X = -0.88
	pointsOfFunctionPlot[506].Y = -0.786

	pointsOfFunctionPlot[507].X = -0.87
	pointsOfFunctionPlot[507].Y = -0.781

	pointsOfFunctionPlot[508].X = -0.86
	pointsOfFunctionPlot[508].Y = -0.776

	pointsOfFunctionPlot[509].X = -0.85
	pointsOfFunctionPlot[509].Y = -0.77

	pointsOfFunctionPlot[510].X = -0.84
	pointsOfFunctionPlot[510].Y = -0.765

	pointsOfFunctionPlot[511].X = -0.83
	pointsOfFunctionPlot[511].Y = -0.759

	pointsOfFunctionPlot[512].X = -0.82
	pointsOfFunctionPlot[512].Y = -0.753

	pointsOfFunctionPlot[513].X = -0.81
	pointsOfFunctionPlot[513].Y = -0.748

	pointsOfFunctionPlot[514].X = -0.80
	pointsOfFunctionPlot[514].Y = -0.742

	pointsOfFunctionPlot[515].X = -0.79
	pointsOfFunctionPlot[515].Y = -0.736

	pointsOfFunctionPlot[516].X = -0.78
	pointsOfFunctionPlot[516].Y = -0.73

	pointsOfFunctionPlot[517].X = -0.77
	pointsOfFunctionPlot[517].Y = -0.723

	pointsOfFunctionPlot[518].X = -0.76
	pointsOfFunctionPlot[518].Y = -0.717

	pointsOfFunctionPlot[519].X = -0.75
	pointsOfFunctionPlot[519].Y = -0.711

	pointsOfFunctionPlot[520].X = -0.74
	pointsOfFunctionPlot[520].Y = -0.704

	pointsOfFunctionPlot[521].X = -0.73
	pointsOfFunctionPlot[521].Y = -0.698

	pointsOfFunctionPlot[522].X = -0.72
	pointsOfFunctionPlot[522].Y = -0.691

	pointsOfFunctionPlot[523].X = -0.71
	pointsOfFunctionPlot[523].Y = -0.684

	pointsOfFunctionPlot[524].X = -0.70
	pointsOfFunctionPlot[524].Y = -0.677

	pointsOfFunctionPlot[525].X = -0.69
	pointsOfFunctionPlot[525].Y = -0.67

	pointsOfFunctionPlot[526].X = -0.68
	pointsOfFunctionPlot[526].Y = -0.663

	pointsOfFunctionPlot[527].X = -0.67
	pointsOfFunctionPlot[527].Y = -0.656

	pointsOfFunctionPlot[528].X = -0.66
	pointsOfFunctionPlot[528].Y = -0.649

	pointsOfFunctionPlot[529].X = -0.65
	pointsOfFunctionPlot[529].Y = -0.642

	pointsOfFunctionPlot[530].X = -0.64
	pointsOfFunctionPlot[530].Y = -0.634

	pointsOfFunctionPlot[531].X = -0.63
	pointsOfFunctionPlot[531].Y = -0.627

	pointsOfFunctionPlot[532].X = -0.62
	pointsOfFunctionPlot[532].Y = -0.619

	pointsOfFunctionPlot[533].X = -0.61
	pointsOfFunctionPlot[533].Y = -0.611

	pointsOfFunctionPlot[534].X = -0.60
	pointsOfFunctionPlot[534].Y = -0.603

	pointsOfFunctionPlot[535].X = -0.59
	pointsOfFunctionPlot[535].Y = -0.595

	pointsOfFunctionPlot[536].X = -0.58
	pointsOfFunctionPlot[536].Y = -0.587

	pointsOfFunctionPlot[537].X = -0.57
	pointsOfFunctionPlot[537].Y = -0.579

	pointsOfFunctionPlot[538].X = -0.56
	pointsOfFunctionPlot[538].Y = -0.571

	pointsOfFunctionPlot[539].X = -0.55
	pointsOfFunctionPlot[539].Y = -0.563

	pointsOfFunctionPlot[540].X = -0.54
	pointsOfFunctionPlot[540].Y = -0.554

	pointsOfFunctionPlot[541].X = -0.53
	pointsOfFunctionPlot[541].Y = -0.546

	pointsOfFunctionPlot[542].X = -0.52
	pointsOfFunctionPlot[542].Y = -0.537

	pointsOfFunctionPlot[543].X = -0.51
	pointsOfFunctionPlot[543].Y = -0.529

	pointsOfFunctionPlot[544].X = -0.50
	pointsOfFunctionPlot[544].Y = -0.52

	pointsOfFunctionPlot[545].X = -0.49
	pointsOfFunctionPlot[545].Y = -0.511

	pointsOfFunctionPlot[546].X = -0.48
	pointsOfFunctionPlot[546].Y = -0.502

	pointsOfFunctionPlot[547].X = -0.47
	pointsOfFunctionPlot[547].Y = -0.493

	pointsOfFunctionPlot[548].X = -0.46
	pointsOfFunctionPlot[548].Y = -0.484

	pointsOfFunctionPlot[549].X = -0.45
	pointsOfFunctionPlot[549].Y = -0.475

	pointsOfFunctionPlot[550].X = -0.44
	pointsOfFunctionPlot[550].Y = -0.466

	pointsOfFunctionPlot[551].X = -0.43
	pointsOfFunctionPlot[551].Y = -0.456

	pointsOfFunctionPlot[552].X = -0.42
	pointsOfFunctionPlot[552].Y = -0.447

	pointsOfFunctionPlot[553].X = -0.41
	pointsOfFunctionPlot[553].Y = -0.437

	pointsOfFunctionPlot[554].X = -0.40
	pointsOfFunctionPlot[554].Y = -0.428

	pointsOfFunctionPlot[555].X = -0.39
	pointsOfFunctionPlot[555].Y = -0.418

	pointsOfFunctionPlot[556].X = -0.38
	pointsOfFunctionPlot[556].Y = -0.409

	pointsOfFunctionPlot[557].X = -0.37
	pointsOfFunctionPlot[557].Y = -0.399

	pointsOfFunctionPlot[558].X = -0.36
	pointsOfFunctionPlot[558].Y = -0.389

	pointsOfFunctionPlot[559].X = -0.35
	pointsOfFunctionPlot[559].Y = -0.379

	pointsOfFunctionPlot[560].X = -0.34
	pointsOfFunctionPlot[560].Y = -0.369

	pointsOfFunctionPlot[561].X = -0.33
	pointsOfFunctionPlot[561].Y = -0.359

	pointsOfFunctionPlot[562].X = -0.32
	pointsOfFunctionPlot[562].Y = -0.349

	pointsOfFunctionPlot[563].X = -0.31
	pointsOfFunctionPlot[563].Y = -0.338

	pointsOfFunctionPlot[564].X = -0.30
	pointsOfFunctionPlot[564].Y = -0.328

	pointsOfFunctionPlot[565].X = -0.29
	pointsOfFunctionPlot[565].Y = -0.318

	pointsOfFunctionPlot[566].X = -0.28
	pointsOfFunctionPlot[566].Y = -0.307

	pointsOfFunctionPlot[567].X = -0.27
	pointsOfFunctionPlot[567].Y = -0.297

	pointsOfFunctionPlot[568].X = -0.26
	pointsOfFunctionPlot[568].Y = -0.286

	pointsOfFunctionPlot[569].X = -0.25
	pointsOfFunctionPlot[569].Y = -0.276

	pointsOfFunctionPlot[570].X = -0.24
	pointsOfFunctionPlot[570].Y = -0.265

	pointsOfFunctionPlot[571].X = -0.23
	pointsOfFunctionPlot[571].Y = -0.255

	pointsOfFunctionPlot[572].X = -0.22
	pointsOfFunctionPlot[572].Y = -0.244

	pointsOfFunctionPlot[573].X = -0.21
	pointsOfFunctionPlot[573].Y = -0.233

	pointsOfFunctionPlot[574].X = -0.20
	pointsOfFunctionPlot[574].Y = -0.222

	pointsOfFunctionPlot[575].X = -0.19
	pointsOfFunctionPlot[575].Y = -0.211

	pointsOfFunctionPlot[576].X = -0.18
	pointsOfFunctionPlot[576].Y = -0.2

	pointsOfFunctionPlot[577].X = -0.17
	pointsOfFunctionPlot[577].Y = -0.189

	pointsOfFunctionPlot[578].X = -0.16
	pointsOfFunctionPlot[578].Y = -0.179

	pointsOfFunctionPlot[579].X = -0.15
	pointsOfFunctionPlot[579].Y = -0.167

	pointsOfFunctionPlot[580].X = -0.14
	pointsOfFunctionPlot[580].Y = -0.156

	pointsOfFunctionPlot[581].X = -0.13
	pointsOfFunctionPlot[581].Y = -0.145

	pointsOfFunctionPlot[582].X = -0.12
	pointsOfFunctionPlot[582].Y = -0.134

	pointsOfFunctionPlot[583].X = -0.11
	pointsOfFunctionPlot[583].Y = -0.123

	pointsOfFunctionPlot[584].X = -0.10
	pointsOfFunctionPlot[584].Y = -0.112

	pointsOfFunctionPlot[585].X = -0.09
	pointsOfFunctionPlot[585].Y = -0.101

	pointsOfFunctionPlot[586].X = -0.08
	pointsOfFunctionPlot[586].Y = -0.09

	pointsOfFunctionPlot[587].X = -0.07
	pointsOfFunctionPlot[587].Y = -0.078

	pointsOfFunctionPlot[588].X = -0.06
	pointsOfFunctionPlot[588].Y = -0.067

	pointsOfFunctionPlot[589].X = -0.05
	pointsOfFunctionPlot[589].Y = -0.056

	pointsOfFunctionPlot[590].X = -0.04
	pointsOfFunctionPlot[590].Y = -0.045

	pointsOfFunctionPlot[591].X = -0.03
	pointsOfFunctionPlot[591].Y = -0.033

	pointsOfFunctionPlot[592].X = -0.02
	pointsOfFunctionPlot[592].Y = -0.022

	pointsOfFunctionPlot[593].X = -0.01
	pointsOfFunctionPlot[593].Y = -0.011

	pointsOfFunctionPlot[594].X = 0.0
	pointsOfFunctionPlot[594].Y = 0.0

	pointsOfFunctionPlot[595].X = 0.01
	pointsOfFunctionPlot[595].Y = 0.011

	pointsOfFunctionPlot[596].X = 0.02
	pointsOfFunctionPlot[596].Y = 0.022

	pointsOfFunctionPlot[597].X = 0.03
	pointsOfFunctionPlot[597].Y = 0.033

	pointsOfFunctionPlot[598].X = 0.04
	pointsOfFunctionPlot[598].Y = 0.045

	pointsOfFunctionPlot[599].X = 0.05
	pointsOfFunctionPlot[599].Y = 0.056

	pointsOfFunctionPlot[600].X = 0.06
	pointsOfFunctionPlot[600].Y = 0.067

	pointsOfFunctionPlot[601].X = 0.07
	pointsOfFunctionPlot[601].Y = 0.078

	pointsOfFunctionPlot[602].X = 0.08
	pointsOfFunctionPlot[602].Y = 0.09

	pointsOfFunctionPlot[603].X = 0.09
	pointsOfFunctionPlot[603].Y = 0.101

	pointsOfFunctionPlot[604].X = 0.10
	pointsOfFunctionPlot[604].Y = 0.112

	pointsOfFunctionPlot[605].X = 0.11
	pointsOfFunctionPlot[605].Y = 0.123

	pointsOfFunctionPlot[606].X = 0.12
	pointsOfFunctionPlot[606].Y = 0.134

	pointsOfFunctionPlot[607].X = 0.13
	pointsOfFunctionPlot[607].Y = 0.145

	pointsOfFunctionPlot[608].X = 0.14
	pointsOfFunctionPlot[608].Y = 0.156

	pointsOfFunctionPlot[609].X = 0.15
	pointsOfFunctionPlot[609].Y = 0.167

	pointsOfFunctionPlot[610].X = 0.16
	pointsOfFunctionPlot[610].Y = 0.179

	pointsOfFunctionPlot[611].X = 0.17
	pointsOfFunctionPlot[611].Y = 0.189

	pointsOfFunctionPlot[612].X = 0.18
	pointsOfFunctionPlot[612].Y = 0.2

	pointsOfFunctionPlot[613].X = 0.19
	pointsOfFunctionPlot[613].Y = 0.211

	pointsOfFunctionPlot[614].X = 0.20
	pointsOfFunctionPlot[614].Y = 0.222

	pointsOfFunctionPlot[615].X = 0.21
	pointsOfFunctionPlot[615].Y = 0.233

	pointsOfFunctionPlot[616].X = 0.22
	pointsOfFunctionPlot[616].Y = 0.244

	pointsOfFunctionPlot[617].X = 0.23
	pointsOfFunctionPlot[617].Y = 0.255

	pointsOfFunctionPlot[618].X = 0.24
	pointsOfFunctionPlot[618].Y = 0.265

	pointsOfFunctionPlot[619].X = 0.25
	pointsOfFunctionPlot[619].Y = 0.276

	pointsOfFunctionPlot[620].X = 0.26
	pointsOfFunctionPlot[620].Y = 0.286

	pointsOfFunctionPlot[621].X = 0.27
	pointsOfFunctionPlot[621].Y = 0.297

	pointsOfFunctionPlot[622].X = 0.28
	pointsOfFunctionPlot[622].Y = 0.307

	pointsOfFunctionPlot[623].X = 0.29
	pointsOfFunctionPlot[623].Y = 0.318

	pointsOfFunctionPlot[624].X = 0.30
	pointsOfFunctionPlot[624].Y = 0.328

	pointsOfFunctionPlot[625].X = 0.31
	pointsOfFunctionPlot[625].Y = 0.338

	pointsOfFunctionPlot[626].X = 0.32
	pointsOfFunctionPlot[626].Y = 0.349

	pointsOfFunctionPlot[627].X = 0.33
	pointsOfFunctionPlot[627].Y = 0.359

	pointsOfFunctionPlot[628].X = 0.34
	pointsOfFunctionPlot[628].Y = 0.369

	pointsOfFunctionPlot[629].X = 0.35
	pointsOfFunctionPlot[629].Y = 0.379

	pointsOfFunctionPlot[630].X = 0.36
	pointsOfFunctionPlot[630].Y = 0.389

	pointsOfFunctionPlot[631].X = 0.37
	pointsOfFunctionPlot[631].Y = 0.399

	pointsOfFunctionPlot[632].X = 0.38
	pointsOfFunctionPlot[632].Y = 0.409

	pointsOfFunctionPlot[633].X = 0.39
	pointsOfFunctionPlot[633].Y = 0.418

	pointsOfFunctionPlot[634].X = 0.40
	pointsOfFunctionPlot[634].Y = 0.428

	pointsOfFunctionPlot[635].X = 0.41
	pointsOfFunctionPlot[635].Y = 0.437

	pointsOfFunctionPlot[636].X = 0.42
	pointsOfFunctionPlot[636].Y = 0.447

	pointsOfFunctionPlot[637].X = 0.43
	pointsOfFunctionPlot[637].Y = 0.456

	pointsOfFunctionPlot[638].X = 0.44
	pointsOfFunctionPlot[638].Y = 0.466

	pointsOfFunctionPlot[639].X = 0.45
	pointsOfFunctionPlot[639].Y = 0.475

	pointsOfFunctionPlot[640].X = 0.46
	pointsOfFunctionPlot[640].Y = 0.484

	pointsOfFunctionPlot[641].X = 0.47
	pointsOfFunctionPlot[641].Y = 0.493

	pointsOfFunctionPlot[642].X = 0.48
	pointsOfFunctionPlot[642].Y = 0.502

	pointsOfFunctionPlot[643].X = 0.49
	pointsOfFunctionPlot[643].Y = 0.511

	pointsOfFunctionPlot[644].X = 0.50
	pointsOfFunctionPlot[644].Y = 0.520

	pointsOfFunctionPlot[645].X = 0.51
	pointsOfFunctionPlot[645].Y = 0.529

	pointsOfFunctionPlot[646].X = 0.52
	pointsOfFunctionPlot[646].Y = 0.537

	pointsOfFunctionPlot[647].X = 0.53
	pointsOfFunctionPlot[647].Y = 0.546

	pointsOfFunctionPlot[648].X = 0.54
	pointsOfFunctionPlot[648].Y = 0.554

	pointsOfFunctionPlot[649].X = 0.55
	pointsOfFunctionPlot[649].Y = 0.563

	pointsOfFunctionPlot[650].X = 0.56
	pointsOfFunctionPlot[650].Y = 0.571

	pointsOfFunctionPlot[651].X = 0.57
	pointsOfFunctionPlot[651].Y = 0.579

	pointsOfFunctionPlot[652].X = 0.58
	pointsOfFunctionPlot[652].Y = 0.587

	pointsOfFunctionPlot[653].X = 0.59
	pointsOfFunctionPlot[653].Y = 0.595

	pointsOfFunctionPlot[654].X = 0.60
	pointsOfFunctionPlot[654].Y = 0.603

	pointsOfFunctionPlot[655].X = 0.61
	pointsOfFunctionPlot[655].Y = 0.611

	pointsOfFunctionPlot[656].X = 0.62
	pointsOfFunctionPlot[656].Y = 0.619

	pointsOfFunctionPlot[657].X = 0.63
	pointsOfFunctionPlot[657].Y = 0.627

	pointsOfFunctionPlot[658].X = 0.64
	pointsOfFunctionPlot[658].Y = 0.634

	pointsOfFunctionPlot[659].X = 0.65
	pointsOfFunctionPlot[659].Y = 0.642

	pointsOfFunctionPlot[660].X = 0.66
	pointsOfFunctionPlot[660].Y = 0.649

	pointsOfFunctionPlot[661].X = 0.67
	pointsOfFunctionPlot[661].Y = 0.656

	pointsOfFunctionPlot[662].X = 0.68
	pointsOfFunctionPlot[662].Y = 0.663

	pointsOfFunctionPlot[663].X = 0.69
	pointsOfFunctionPlot[663].Y = 0.67

	pointsOfFunctionPlot[664].X = 0.70
	pointsOfFunctionPlot[664].Y = 0.677

	pointsOfFunctionPlot[665].X = 0.71
	pointsOfFunctionPlot[665].Y = 0.684

	pointsOfFunctionPlot[666].X = 0.72
	pointsOfFunctionPlot[666].Y = 0.691

	pointsOfFunctionPlot[667].X = 0.73
	pointsOfFunctionPlot[667].Y = 0.698

	pointsOfFunctionPlot[668].X = 0.74
	pointsOfFunctionPlot[668].Y = 0.704

	pointsOfFunctionPlot[669].X = 0.75
	pointsOfFunctionPlot[669].Y = 0.711

	pointsOfFunctionPlot[670].X = 0.76
	pointsOfFunctionPlot[670].Y = 0.717

	pointsOfFunctionPlot[671].X = 0.77
	pointsOfFunctionPlot[671].Y = 0.723

	pointsOfFunctionPlot[672].X = 0.78
	pointsOfFunctionPlot[672].Y = 0.73

	pointsOfFunctionPlot[673].X = 0.79
	pointsOfFunctionPlot[673].Y = 0.736

	pointsOfFunctionPlot[674].X = 0.80
	pointsOfFunctionPlot[674].Y = 0.742

	pointsOfFunctionPlot[675].X = 0.81
	pointsOfFunctionPlot[675].Y = 0.748

	pointsOfFunctionPlot[676].X = 0.82
	pointsOfFunctionPlot[676].Y = 0.753

	pointsOfFunctionPlot[677].X = 0.83
	pointsOfFunctionPlot[677].Y = 0.759

	pointsOfFunctionPlot[678].X = 0.84
	pointsOfFunctionPlot[678].Y = 0.765

	pointsOfFunctionPlot[679].X = 0.85
	pointsOfFunctionPlot[679].Y = 0.77

	pointsOfFunctionPlot[680].X = 0.86
	pointsOfFunctionPlot[680].Y = 0.776

	pointsOfFunctionPlot[681].X = 0.87
	pointsOfFunctionPlot[681].Y = 0.781

	pointsOfFunctionPlot[682].X = 0.88
	pointsOfFunctionPlot[682].Y = 0.786

	pointsOfFunctionPlot[683].X = 0.89
	pointsOfFunctionPlot[683].Y = 0.791

	pointsOfFunctionPlot[684].X = 0.90
	pointsOfFunctionPlot[684].Y = 0.796

	pointsOfFunctionPlot[685].X = 0.91
	pointsOfFunctionPlot[685].Y = 0.801

	pointsOfFunctionPlot[686].X = 0.92
	pointsOfFunctionPlot[686].Y = 0.806

	pointsOfFunctionPlot[687].X = 0.93
	pointsOfFunctionPlot[687].Y = 0.811

	pointsOfFunctionPlot[688].X = 0.94
	pointsOfFunctionPlot[688].Y = 0.816

	pointsOfFunctionPlot[689].X = 0.95
	pointsOfFunctionPlot[689].Y = 0.82

	pointsOfFunctionPlot[690].X = 0.96
	pointsOfFunctionPlot[690].Y = 0.825

	pointsOfFunctionPlot[691].X = 0.97
	pointsOfFunctionPlot[691].Y = 0.829

	pointsOfFunctionPlot[692].X = 0.98
	pointsOfFunctionPlot[692].Y = 0.834

	pointsOfFunctionPlot[693].X = 0.99
	pointsOfFunctionPlot[693].Y = 0.838

	pointsOfFunctionPlot[694].X = 1.0
	pointsOfFunctionPlot[694].Y = 0.842

	pointsOfFunctionPlot[695].X = 1.01
	pointsOfFunctionPlot[695].Y = 0.846

	pointsOfFunctionPlot[696].X = 1.02
	pointsOfFunctionPlot[696].Y = 0.85

	pointsOfFunctionPlot[697].X = 1.03
	pointsOfFunctionPlot[697].Y = 0.854

	pointsOfFunctionPlot[698].X = 1.04
	pointsOfFunctionPlot[698].Y = 0.858

	pointsOfFunctionPlot[699].X = 1.05
	pointsOfFunctionPlot[699].Y = 0.862

	pointsOfFunctionPlot[700].X = 1.06
	pointsOfFunctionPlot[700].Y = 0.866

	pointsOfFunctionPlot[701].X = 1.07
	pointsOfFunctionPlot[701].Y = 0.869

	pointsOfFunctionPlot[702].X = 1.08
	pointsOfFunctionPlot[702].Y = 0.873

	pointsOfFunctionPlot[703].X = 1.09
	pointsOfFunctionPlot[703].Y = 0.876

	pointsOfFunctionPlot[704].X = 1.10
	pointsOfFunctionPlot[704].Y = 0.88

	pointsOfFunctionPlot[705].X = 1.11
	pointsOfFunctionPlot[705].Y = 0.883

	pointsOfFunctionPlot[706].X = 1.12
	pointsOfFunctionPlot[706].Y = 0.886

	pointsOfFunctionPlot[707].X = 1.13
	pointsOfFunctionPlot[707].Y = 0.889

	pointsOfFunctionPlot[708].X = 1.14
	pointsOfFunctionPlot[708].Y = 0.893

	pointsOfFunctionPlot[709].X = 1.15
	pointsOfFunctionPlot[709].Y = 0.896

	pointsOfFunctionPlot[710].X = 1.16
	pointsOfFunctionPlot[710].Y = 0.899

	pointsOfFunctionPlot[711].X = 1.17
	pointsOfFunctionPlot[711].Y = 0.902

	pointsOfFunctionPlot[712].X = 1.18
	pointsOfFunctionPlot[712].Y = 0.904

	pointsOfFunctionPlot[713].X = 1.19
	pointsOfFunctionPlot[713].Y = 0.907

	pointsOfFunctionPlot[714].X = 1.20
	pointsOfFunctionPlot[714].Y = 0.91

	pointsOfFunctionPlot[715].X = 1.21
	pointsOfFunctionPlot[715].Y = 0.912

	pointsOfFunctionPlot[716].X = 1.22
	pointsOfFunctionPlot[716].Y = 0.915

	pointsOfFunctionPlot[717].X = 1.23
	pointsOfFunctionPlot[717].Y = 0.918

	pointsOfFunctionPlot[718].X = 1.24
	pointsOfFunctionPlot[718].Y = 0.92

	pointsOfFunctionPlot[719].X = 1.25
	pointsOfFunctionPlot[719].Y = 0.922

	pointsOfFunctionPlot[720].X = 1.26
	pointsOfFunctionPlot[720].Y = 0.925

	pointsOfFunctionPlot[721].X = 1.27
	pointsOfFunctionPlot[721].Y = 0.927

	pointsOfFunctionPlot[722].X = 1.28
	pointsOfFunctionPlot[722].Y = 0.929

	pointsOfFunctionPlot[723].X = 1.29
	pointsOfFunctionPlot[723].Y = 0.931

	pointsOfFunctionPlot[724].X = 1.30
	pointsOfFunctionPlot[724].Y = 0.934

	pointsOfFunctionPlot[725].X = 1.31
	pointsOfFunctionPlot[725].Y = 0.936

	pointsOfFunctionPlot[726].X = 1.32
	pointsOfFunctionPlot[726].Y = 0.938

	pointsOfFunctionPlot[727].X = 1.33
	pointsOfFunctionPlot[727].Y = 0.94

	pointsOfFunctionPlot[728].X = 1.34
	pointsOfFunctionPlot[728].Y = 0.941

	pointsOfFunctionPlot[729].X = 1.35
	pointsOfFunctionPlot[729].Y = 0.943

	pointsOfFunctionPlot[730].X = 1.36
	pointsOfFunctionPlot[730].Y = 0.945

	pointsOfFunctionPlot[731].X = 1.37
	pointsOfFunctionPlot[731].Y = 0.947

	pointsOfFunctionPlot[732].X = 1.38
	pointsOfFunctionPlot[732].Y = 0.949

	pointsOfFunctionPlot[733].X = 1.39
	pointsOfFunctionPlot[733].Y = 0.95

	pointsOfFunctionPlot[734].X = 1.40
	pointsOfFunctionPlot[734].Y = 0.952

	pointsOfFunctionPlot[735].X = 1.41
	pointsOfFunctionPlot[735].Y = 0.953

	pointsOfFunctionPlot[736].X = 1.42
	pointsOfFunctionPlot[736].Y = 0.955

	pointsOfFunctionPlot[737].X = 1.43
	pointsOfFunctionPlot[737].Y = 0.956

	pointsOfFunctionPlot[738].X = 1.44
	pointsOfFunctionPlot[738].Y = 0.958

	pointsOfFunctionPlot[739].X = 1.45
	pointsOfFunctionPlot[739].Y = 0.959

	pointsOfFunctionPlot[740].X = 1.46
	pointsOfFunctionPlot[740].Y = 0.961

	pointsOfFunctionPlot[741].X = 1.47
	pointsOfFunctionPlot[741].Y = 0.962

	pointsOfFunctionPlot[742].X = 1.48
	pointsOfFunctionPlot[742].Y = 0.963

	pointsOfFunctionPlot[743].X = 1.49
	pointsOfFunctionPlot[743].Y = 0.964

	pointsOfFunctionPlot[744].X = 1.50
	pointsOfFunctionPlot[744].Y = 0.966

	pointsOfFunctionPlot[745].X = 1.51
	pointsOfFunctionPlot[745].Y = 0.967

	pointsOfFunctionPlot[746].X = 1.52
	pointsOfFunctionPlot[746].Y = 0.968

	pointsOfFunctionPlot[747].X = 1.53
	pointsOfFunctionPlot[747].Y = 0.969

	pointsOfFunctionPlot[748].X = 1.54
	pointsOfFunctionPlot[748].Y = 0.97

	pointsOfFunctionPlot[749].X = 1.55
	pointsOfFunctionPlot[749].Y = 0.971

	pointsOfFunctionPlot[750].X = 1.56
	pointsOfFunctionPlot[750].Y = 0.972

	pointsOfFunctionPlot[751].X = 1.57
	pointsOfFunctionPlot[751].Y = 0.973

	pointsOfFunctionPlot[752].X = 1.58
	pointsOfFunctionPlot[752].Y = 0.974

	pointsOfFunctionPlot[753].X = 1.59
	pointsOfFunctionPlot[753].Y = 0.975

	pointsOfFunctionPlot[754].X = 1.60
	pointsOfFunctionPlot[754].Y = 0.976

	pointsOfFunctionPlot[755].X = 1.61
	pointsOfFunctionPlot[755].Y = 0.977

	pointsOfFunctionPlot[756].X = 1.62
	pointsOfFunctionPlot[756].Y = 0.978

	pointsOfFunctionPlot[757].X = 1.63
	pointsOfFunctionPlot[757].Y = 0.978

	pointsOfFunctionPlot[758].X = 1.64
	pointsOfFunctionPlot[758].Y = 0.979

	pointsOfFunctionPlot[759].X = 1.65
	pointsOfFunctionPlot[759].Y = 0.98

	pointsOfFunctionPlot[760].X = 1.66
	pointsOfFunctionPlot[760].Y = 0.981

	pointsOfFunctionPlot[761].X = 1.67
	pointsOfFunctionPlot[761].Y = 0.981

	pointsOfFunctionPlot[762].X = 1.68
	pointsOfFunctionPlot[762].Y = 0.982

	pointsOfFunctionPlot[763].X = 1.69
	pointsOfFunctionPlot[763].Y = 0.983

	pointsOfFunctionPlot[764].X = 1.70
	pointsOfFunctionPlot[764].Y = 0.983

	pointsOfFunctionPlot[765].X = 1.71
	pointsOfFunctionPlot[765].Y = 0.984

	pointsOfFunctionPlot[766].X = 1.72
	pointsOfFunctionPlot[766].Y = 0.985

	pointsOfFunctionPlot[767].X = 1.73
	pointsOfFunctionPlot[767].Y = 0.985

	pointsOfFunctionPlot[768].X = 1.74
	pointsOfFunctionPlot[768].Y = 0.986

	pointsOfFunctionPlot[769].X = 1.75
	pointsOfFunctionPlot[769].Y = 0.986

	pointsOfFunctionPlot[770].X = 1.76
	pointsOfFunctionPlot[770].Y = 0.987

	pointsOfFunctionPlot[771].X = 1.77
	pointsOfFunctionPlot[771].Y = 0.987

	pointsOfFunctionPlot[772].X = 1.78
	pointsOfFunctionPlot[772].Y = 0.988

	pointsOfFunctionPlot[773].X = 1.79
	pointsOfFunctionPlot[773].Y = 0.988

	pointsOfFunctionPlot[774].X = 1.80
	pointsOfFunctionPlot[774].Y = 0.989

	pointsOfFunctionPlot[775].X = 1.81
	pointsOfFunctionPlot[775].Y = 0.989

	pointsOfFunctionPlot[776].X = 1.82
	pointsOfFunctionPlot[776].Y = 0.989

	pointsOfFunctionPlot[777].X = 1.83
	pointsOfFunctionPlot[777].Y = 0.99

	pointsOfFunctionPlot[778].X = 1.84
	pointsOfFunctionPlot[778].Y = 0.99

	pointsOfFunctionPlot[779].X = 1.85
	pointsOfFunctionPlot[779].Y = 0.991

	pointsOfFunctionPlot[780].X = 1.86
	pointsOfFunctionPlot[780].Y = 0.991

	pointsOfFunctionPlot[781].X = 1.87
	pointsOfFunctionPlot[781].Y = 0.991

	pointsOfFunctionPlot[782].X = 1.88
	pointsOfFunctionPlot[782].Y = 0.992

	pointsOfFunctionPlot[783].X = 1.89
	pointsOfFunctionPlot[783].Y = 0.992

	pointsOfFunctionPlot[784].X = 1.90
	pointsOfFunctionPlot[784].Y = 0.992

	pointsOfFunctionPlot[785].X = 1.91
	pointsOfFunctionPlot[785].Y = 0.993

	pointsOfFunctionPlot[786].X = 1.92
	pointsOfFunctionPlot[786].Y = 0.993

	pointsOfFunctionPlot[787].X = 1.93
	pointsOfFunctionPlot[787].Y = 0.993

	pointsOfFunctionPlot[788].X = 1.94
	pointsOfFunctionPlot[788].Y = 0.993

	pointsOfFunctionPlot[789].X = 1.95
	pointsOfFunctionPlot[789].Y = 0.994

	pointsOfFunctionPlot[790].X = 1.96
	pointsOfFunctionPlot[790].Y = 0.994

	pointsOfFunctionPlot[791].X = 1.97
	pointsOfFunctionPlot[791].Y = 0.994

	pointsOfFunctionPlot[792].X = 1.98
	pointsOfFunctionPlot[792].Y = 0.994

	pointsOfFunctionPlot[793].X = 1.99
	pointsOfFunctionPlot[793].Y = 0.995

	pointsOfFunctionPlot[794].X = 2.0
	pointsOfFunctionPlot[794].Y = 0.995

	pointsOfFunctionPlot[795].X = 2.01
	pointsOfFunctionPlot[795].Y = 0.995

	pointsOfFunctionPlot[796].X = 2.02
	pointsOfFunctionPlot[796].Y = 0.995

	pointsOfFunctionPlot[797].X = 2.03
	pointsOfFunctionPlot[797].Y = 0.995

	pointsOfFunctionPlot[798].X = 2.04
	pointsOfFunctionPlot[798].Y = 0.996

	pointsOfFunctionPlot[799].X = 2.05
	pointsOfFunctionPlot[799].Y = 0.996

	pointsOfFunctionPlot[800].X = 2.06
	pointsOfFunctionPlot[800].Y = 0.996

	pointsOfFunctionPlot[801].X = 2.07
	pointsOfFunctionPlot[801].Y = 0.996

	pointsOfFunctionPlot[802].X = 2.08
	pointsOfFunctionPlot[802].Y = 0.996

	pointsOfFunctionPlot[803].X = 2.09
	pointsOfFunctionPlot[803].Y = 0.996

	pointsOfFunctionPlot[804].X = 2.10
	pointsOfFunctionPlot[804].Y = 0.997

	pointsOfFunctionPlot[805].X = 2.11
	pointsOfFunctionPlot[805].Y = 0.997

	pointsOfFunctionPlot[806].X = 2.12
	pointsOfFunctionPlot[806].Y = 0.997

	pointsOfFunctionPlot[807].X = 2.13
	pointsOfFunctionPlot[807].Y = 0.997

	pointsOfFunctionPlot[808].X = 2.14
	pointsOfFunctionPlot[808].Y = 0.997

	pointsOfFunctionPlot[809].X = 2.15
	pointsOfFunctionPlot[809].Y = 0.997

	pointsOfFunctionPlot[810].X = 2.16
	pointsOfFunctionPlot[810].Y = 0.997

	pointsOfFunctionPlot[811].X = 2.17
	pointsOfFunctionPlot[811].Y = 0.997

	pointsOfFunctionPlot[812].X = 2.18
	pointsOfFunctionPlot[812].Y = 0.997

	pointsOfFunctionPlot[813].X = 2.19
	pointsOfFunctionPlot[813].Y = 0.998

	pointsOfFunctionPlot[814].X = 2.20
	pointsOfFunctionPlot[814].Y = 0.998

	pointsOfFunctionPlot[815].X = 2.21
	pointsOfFunctionPlot[815].Y = 0.998

	pointsOfFunctionPlot[816].X = 2.22
	pointsOfFunctionPlot[816].Y = 0.998

	pointsOfFunctionPlot[817].X = 2.23
	pointsOfFunctionPlot[817].Y = 0.998

	pointsOfFunctionPlot[818].X = 2.24
	pointsOfFunctionPlot[818].Y = 0.998

	pointsOfFunctionPlot[819].X = 2.25
	pointsOfFunctionPlot[819].Y = 0.998

	pointsOfFunctionPlot[820].X = 2.26
	pointsOfFunctionPlot[820].Y = 0.998

	pointsOfFunctionPlot[821].X = 2.27
	pointsOfFunctionPlot[821].Y = 0.998

	pointsOfFunctionPlot[822].X = 2.28
	pointsOfFunctionPlot[822].Y = 0.998

	pointsOfFunctionPlot[823].X = 2.29
	pointsOfFunctionPlot[823].Y = 0.998

	pointsOfFunctionPlot[824].X = 2.30
	pointsOfFunctionPlot[824].Y = 0.998

	pointsOfFunctionPlot[825].X = 2.31
	pointsOfFunctionPlot[825].Y = 0.998

	pointsOfFunctionPlot[826].X = 2.32
	pointsOfFunctionPlot[826].Y = 0.998

	pointsOfFunctionPlot[827].X = 2.33
	pointsOfFunctionPlot[827].Y = 0.999

	pointsOfFunctionPlot[828].X = 2.34
	pointsOfFunctionPlot[828].Y = 0.999

	pointsOfFunctionPlot[829].X = 2.35
	pointsOfFunctionPlot[829].Y = 0.999

	pointsOfFunctionPlot[830].X = 2.36
	pointsOfFunctionPlot[830].Y = 0.999

	pointsOfFunctionPlot[831].X = 2.37
	pointsOfFunctionPlot[831].Y = 0.999

	pointsOfFunctionPlot[832].X = 2.38
	pointsOfFunctionPlot[832].Y = 0.999

	pointsOfFunctionPlot[833].X = 2.39
	pointsOfFunctionPlot[833].Y = 0.999

	pointsOfFunctionPlot[834].X = 2.40
	pointsOfFunctionPlot[834].Y = 0.999

	pointsOfFunctionPlot[835].X = 2.41
	pointsOfFunctionPlot[835].Y = 0.999

	pointsOfFunctionPlot[836].X = 2.42
	pointsOfFunctionPlot[836].Y = 0.999

	pointsOfFunctionPlot[837].X = 2.43
	pointsOfFunctionPlot[837].Y = 0.999

	pointsOfFunctionPlot[838].X = 2.44
	pointsOfFunctionPlot[838].Y = 0.999

	pointsOfFunctionPlot[839].X = 2.45
	pointsOfFunctionPlot[839].Y = 0.999

	pointsOfFunctionPlot[840].X = 2.46
	pointsOfFunctionPlot[840].Y = 0.999

	pointsOfFunctionPlot[841].X = 2.47
	pointsOfFunctionPlot[841].Y = 0.999

	pointsOfFunctionPlot[842].X = 2.48
	pointsOfFunctionPlot[842].Y = 0.999

	pointsOfFunctionPlot[843].X = 2.49
	pointsOfFunctionPlot[843].Y = 0.999

	pointsOfFunctionPlot[844].X = 2.50
	pointsOfFunctionPlot[844].Y = 0.999

	pointsOfFunctionPlot[845].X = 2.51
	pointsOfFunctionPlot[845].Y = 0.999

	pointsOfFunctionPlot[846].X = 2.52
	pointsOfFunctionPlot[846].Y = 0.999

	pointsOfFunctionPlot[847].X = 2.53
	pointsOfFunctionPlot[847].Y = 0.999

	pointsOfFunctionPlot[848].X = 2.54
	pointsOfFunctionPlot[848].Y = 0.999

	pointsOfFunctionPlot[849].X = 2.55
	pointsOfFunctionPlot[849].Y = 0.999

	pointsOfFunctionPlot[850].X = 2.56
	pointsOfFunctionPlot[850].Y = 0.999

	pointsOfFunctionPlot[851].X = 2.57
	pointsOfFunctionPlot[851].Y = 0.999

	pointsOfFunctionPlot[852].X = 2.58
	pointsOfFunctionPlot[852].Y = 0.999

	pointsOfFunctionPlot[853].X = 2.59
	pointsOfFunctionPlot[853].Y = 0.999

	pointsOfFunctionPlot[854].X = 2.60
	pointsOfFunctionPlot[854].Y = 0.999

	pointsOfFunctionPlot[855].X = 2.61
	pointsOfFunctionPlot[855].Y = 0.999

	pointsOfFunctionPlot[856].X = 2.62
	pointsOfFunctionPlot[856].Y = 0.999

	pointsOfFunctionPlot[857].X = 2.63
	pointsOfFunctionPlot[857].Y = 0.999

	pointsOfFunctionPlot[858].X = 2.64
	pointsOfFunctionPlot[858].Y = 0.999

	pointsOfFunctionPlot[859].X = 2.65
	pointsOfFunctionPlot[859].Y = 0.999

	pointsOfFunctionPlot[860].X = 2.66
	pointsOfFunctionPlot[860].Y = 0.999

	pointsOfFunctionPlot[861].X = 2.67
	pointsOfFunctionPlot[861].Y = 0.999

	pointsOfFunctionPlot[862].X = 2.68
	pointsOfFunctionPlot[862].Y = 0.999

	pointsOfFunctionPlot[863].X = 2.69
	pointsOfFunctionPlot[863].Y = 0.999

	pointsOfFunctionPlot[864].X = 2.70
	pointsOfFunctionPlot[864].Y = 0.999

	pointsOfFunctionPlot[865].X = 2.71
	pointsOfFunctionPlot[865].Y = 0.999

	pointsOfFunctionPlot[866].X = 2.72
	pointsOfFunctionPlot[866].Y = 0.999

	pointsOfFunctionPlot[867].X = 2.73
	pointsOfFunctionPlot[867].Y = 0.999

	pointsOfFunctionPlot[868].X = 2.74
	pointsOfFunctionPlot[868].Y = 0.999

	pointsOfFunctionPlot[869].X = 2.75
	pointsOfFunctionPlot[869].Y = 0.999

	pointsOfFunctionPlot[870].X = 2.76
	pointsOfFunctionPlot[870].Y = 0.999

	pointsOfFunctionPlot[871].X = 2.77
	pointsOfFunctionPlot[871].Y = 0.999

	pointsOfFunctionPlot[872].X = 2.78
	pointsOfFunctionPlot[872].Y = 0.999

	pointsOfFunctionPlot[873].X = 2.79
	pointsOfFunctionPlot[873].Y = 0.999

	pointsOfFunctionPlot[874].X = 2.80
	pointsOfFunctionPlot[874].Y = 0.999

	pointsOfFunctionPlot[875].X = 2.81
	pointsOfFunctionPlot[875].Y = 0.999

	pointsOfFunctionPlot[876].X = 2.82
	pointsOfFunctionPlot[876].Y = 0.999

	pointsOfFunctionPlot[877].X = 2.83
	pointsOfFunctionPlot[877].Y = 0.999

	pointsOfFunctionPlot[878].X = 2.84
	pointsOfFunctionPlot[878].Y = 0.999

	pointsOfFunctionPlot[879].X = 2.85
	pointsOfFunctionPlot[879].Y = 0.999

	pointsOfFunctionPlot[880].X = 2.86
	pointsOfFunctionPlot[880].Y = 0.999

	pointsOfFunctionPlot[881].X = 2.87
	pointsOfFunctionPlot[881].Y = 0.999

	pointsOfFunctionPlot[882].X = 2.88
	pointsOfFunctionPlot[882].Y = 0.999

	pointsOfFunctionPlot[883].X = 2.89
	pointsOfFunctionPlot[883].Y = 0.999

	pointsOfFunctionPlot[884].X = 2.90
	pointsOfFunctionPlot[884].Y = 0.999

	pointsOfFunctionPlot[885].X = 2.91
	pointsOfFunctionPlot[885].Y = 0.999

	pointsOfFunctionPlot[886].X = 2.92
	pointsOfFunctionPlot[886].Y = 0.999

	pointsOfFunctionPlot[887].X = 2.93
	pointsOfFunctionPlot[887].Y = 0.999

	pointsOfFunctionPlot[888].X = 2.94
	pointsOfFunctionPlot[888].Y = 0.999

	pointsOfFunctionPlot[889].X = 2.95
	pointsOfFunctionPlot[889].Y = 0.999

	pointsOfFunctionPlot[890].X = 2.96
	pointsOfFunctionPlot[890].Y = 0.999

	pointsOfFunctionPlot[891].X = 2.97
	pointsOfFunctionPlot[891].Y = 0.999

	pointsOfFunctionPlot[892].X = 2.98
	pointsOfFunctionPlot[892].Y = 0.999

	pointsOfFunctionPlot[893].X = 2.99
	pointsOfFunctionPlot[893].Y = 0.999

	pointsOfFunctionPlot[894].X = 3.0
	pointsOfFunctionPlot[894].Y = 0.999

	pointsOfFunctionPlot[895].X = 3.01
	pointsOfFunctionPlot[895].Y = 0.999

	pointsOfFunctionPlot[896].X = 3.02
	pointsOfFunctionPlot[896].Y = 0.999

	pointsOfFunctionPlot[897].X = 3.03
	pointsOfFunctionPlot[897].Y = 0.999

	pointsOfFunctionPlot[898].X = 3.04
	pointsOfFunctionPlot[898].Y = 0.999

	pointsOfFunctionPlot[899].X = 3.05
	pointsOfFunctionPlot[899].Y = 0.999

	pointsOfFunctionPlot[900].X = 3.06
	pointsOfFunctionPlot[900].Y = 0.999

	pointsOfFunctionPlot[901].X = 3.07
	pointsOfFunctionPlot[901].Y = 0.999

	pointsOfFunctionPlot[902].X = 3.08
	pointsOfFunctionPlot[902].Y = 0.999

	pointsOfFunctionPlot[903].X = 3.09
	pointsOfFunctionPlot[903].Y = 0.999

	pointsOfFunctionPlot[904].X = 3.10
	pointsOfFunctionPlot[904].Y = 0.999

	pointsOfFunctionPlot[905].X = 3.11
	pointsOfFunctionPlot[905].Y = 0.999

	pointsOfFunctionPlot[906].X = 3.12
	pointsOfFunctionPlot[906].Y = 0.999

	pointsOfFunctionPlot[907].X = 3.13
	pointsOfFunctionPlot[907].Y = 0.999

	pointsOfFunctionPlot[908].X = 3.14
	pointsOfFunctionPlot[908].Y = 0.999

	pointsOfFunctionPlot[909].X = 3.15
	pointsOfFunctionPlot[909].Y = 0.999

	pointsOfFunctionPlot[910].X = 3.16
	pointsOfFunctionPlot[910].Y = 0.999

	pointsOfFunctionPlot[911].X = 3.17
	pointsOfFunctionPlot[911].Y = 0.999

	pointsOfFunctionPlot[912].X = 3.18
	pointsOfFunctionPlot[912].Y = 0.999

	pointsOfFunctionPlot[913].X = 3.19
	pointsOfFunctionPlot[913].Y = 0.999

	pointsOfFunctionPlot[914].X = 3.20
	pointsOfFunctionPlot[914].Y = 0.999

	pointsOfFunctionPlot[915].X = 3.21
	pointsOfFunctionPlot[915].Y = 0.999

	pointsOfFunctionPlot[916].X = 3.22
	pointsOfFunctionPlot[916].Y = 0.999

	pointsOfFunctionPlot[917].X = 3.23
	pointsOfFunctionPlot[917].Y = 0.999

	pointsOfFunctionPlot[918].X = 3.24
	pointsOfFunctionPlot[918].Y = 0.999

	pointsOfFunctionPlot[919].X = 3.25
	pointsOfFunctionPlot[919].Y = 0.999

	pointsOfFunctionPlot[920].X = 3.26
	pointsOfFunctionPlot[920].Y = 0.999

	pointsOfFunctionPlot[921].X = 3.27
	pointsOfFunctionPlot[921].Y = 0.999

	pointsOfFunctionPlot[922].X = 3.28
	pointsOfFunctionPlot[922].Y = 0.999

	pointsOfFunctionPlot[923].X = 3.29
	pointsOfFunctionPlot[923].Y = 0.999

	pointsOfFunctionPlot[924].X = 3.30
	pointsOfFunctionPlot[924].Y = 0.999

	pointsOfFunctionPlot[925].X = 3.31
	pointsOfFunctionPlot[925].Y = 0.999

	pointsOfFunctionPlot[926].X = 3.32
	pointsOfFunctionPlot[926].Y = 0.999

	pointsOfFunctionPlot[927].X = 3.33
	pointsOfFunctionPlot[927].Y = 0.999

	pointsOfFunctionPlot[928].X = 3.34
	pointsOfFunctionPlot[928].Y = 0.999

	pointsOfFunctionPlot[929].X = 3.35
	pointsOfFunctionPlot[929].Y = 0.999

	pointsOfFunctionPlot[930].X = 3.36
	pointsOfFunctionPlot[930].Y = 0.999

	pointsOfFunctionPlot[931].X = 3.37
	pointsOfFunctionPlot[931].Y = 0.999

	pointsOfFunctionPlot[932].X = 3.38
	pointsOfFunctionPlot[932].Y = 0.999

	pointsOfFunctionPlot[933].X = 3.39
	pointsOfFunctionPlot[933].Y = 0.999

	pointsOfFunctionPlot[934].X = 3.40
	pointsOfFunctionPlot[934].Y = 0.999

	pointsOfFunctionPlot[935].X = 3.41
	pointsOfFunctionPlot[935].Y = 0.999

	pointsOfFunctionPlot[936].X = 3.42
	pointsOfFunctionPlot[936].Y = 0.999

	pointsOfFunctionPlot[937].X = 3.43
	pointsOfFunctionPlot[937].Y = 0.999

	pointsOfFunctionPlot[938].X = 3.44
	pointsOfFunctionPlot[938].Y = 0.999

	pointsOfFunctionPlot[939].X = 3.45
	pointsOfFunctionPlot[939].Y = 0.999

	pointsOfFunctionPlot[940].X = 3.46
	pointsOfFunctionPlot[940].Y = 0.999

	pointsOfFunctionPlot[941].X = 3.47
	pointsOfFunctionPlot[941].Y = 0.999

	pointsOfFunctionPlot[942].X = 3.48
	pointsOfFunctionPlot[942].Y = 0.999

	pointsOfFunctionPlot[943].X = 3.49
	pointsOfFunctionPlot[943].Y = 0.999

	pointsOfFunctionPlot[944].X = 3.50
	pointsOfFunctionPlot[944].Y = 0.999

	pointsOfFunctionPlot[945].X = 3.51
	pointsOfFunctionPlot[945].Y = 0.999

	pointsOfFunctionPlot[946].X = 3.52
	pointsOfFunctionPlot[946].Y = 0.999

	pointsOfFunctionPlot[947].X = 3.53
	pointsOfFunctionPlot[947].Y = 0.999

	pointsOfFunctionPlot[948].X = 3.54
	pointsOfFunctionPlot[948].Y = 0.999

	pointsOfFunctionPlot[949].X = 3.55
	pointsOfFunctionPlot[949].Y = 0.999

	pointsOfFunctionPlot[950].X = 3.56
	pointsOfFunctionPlot[950].Y = 0.999

	pointsOfFunctionPlot[951].X = 3.57
	pointsOfFunctionPlot[951].Y = 0.999

	pointsOfFunctionPlot[952].X = 3.58
	pointsOfFunctionPlot[952].Y = 0.999

	pointsOfFunctionPlot[953].X = 3.59
	pointsOfFunctionPlot[953].Y = 0.999

	pointsOfFunctionPlot[954].X = 3.60
	pointsOfFunctionPlot[954].Y = 0.999

	pointsOfFunctionPlot[955].X = 3.61
	pointsOfFunctionPlot[955].Y = 0.999

	pointsOfFunctionPlot[956].X = 3.62
	pointsOfFunctionPlot[956].Y = 0.999

	pointsOfFunctionPlot[957].X = 3.63
	pointsOfFunctionPlot[957].Y = 0.999

	pointsOfFunctionPlot[958].X = 3.64
	pointsOfFunctionPlot[958].Y = 0.999

	pointsOfFunctionPlot[959].X = 3.65
	pointsOfFunctionPlot[959].Y = 0.999

	pointsOfFunctionPlot[960].X = 3.66
	pointsOfFunctionPlot[960].Y = 0.999

	pointsOfFunctionPlot[961].X = 3.67
	pointsOfFunctionPlot[961].Y = 0.999

	pointsOfFunctionPlot[962].X = 3.68
	pointsOfFunctionPlot[962].Y = 0.999

	pointsOfFunctionPlot[963].X = 3.69
	pointsOfFunctionPlot[963].Y = 0.999

	pointsOfFunctionPlot[964].X = 3.70
	pointsOfFunctionPlot[964].Y = 0.999

	pointsOfFunctionPlot[965].X = 3.71
	pointsOfFunctionPlot[965].Y = 0.999

	pointsOfFunctionPlot[966].X = 3.72
	pointsOfFunctionPlot[966].Y = 0.999

	pointsOfFunctionPlot[967].X = 3.73
	pointsOfFunctionPlot[967].Y = 0.999

	pointsOfFunctionPlot[968].X = 3.74
	pointsOfFunctionPlot[968].Y = 0.999

	pointsOfFunctionPlot[969].X = 3.75
	pointsOfFunctionPlot[969].Y = 0.999

	pointsOfFunctionPlot[970].X = 3.76
	pointsOfFunctionPlot[970].Y = 0.999

	pointsOfFunctionPlot[971].X = 3.77
	pointsOfFunctionPlot[971].Y = 0.999

	pointsOfFunctionPlot[972].X = 3.78
	pointsOfFunctionPlot[972].Y = 0.999

	pointsOfFunctionPlot[973].X = 3.79
	pointsOfFunctionPlot[973].Y = 0.999

	pointsOfFunctionPlot[974].X = 3.80
	pointsOfFunctionPlot[974].Y = 0.999

	pointsOfFunctionPlot[975].X = 3.81
	pointsOfFunctionPlot[975].Y = 0.999

	pointsOfFunctionPlot[976].X = 3.82
	pointsOfFunctionPlot[976].Y = 0.999

	pointsOfFunctionPlot[977].X = 3.83
	pointsOfFunctionPlot[977].Y = 0.999

	pointsOfFunctionPlot[978].X = 3.84
	pointsOfFunctionPlot[978].Y = 0.999

	pointsOfFunctionPlot[979].X = 3.85
	pointsOfFunctionPlot[979].Y = 0.999

	pointsOfFunctionPlot[980].X = 3.86
	pointsOfFunctionPlot[980].Y = 0.999

	pointsOfFunctionPlot[981].X = 3.87
	pointsOfFunctionPlot[981].Y = 0.999

	pointsOfFunctionPlot[982].X = 3.88
	pointsOfFunctionPlot[982].Y = 0.999

	pointsOfFunctionPlot[983].X = 3.89
	pointsOfFunctionPlot[983].Y = 0.999

	pointsOfFunctionPlot[984].X = 3.90
	pointsOfFunctionPlot[984].Y = 0.999

	pointsOfFunctionPlot[985].X = 3.91
	pointsOfFunctionPlot[985].Y = 0.999

	pointsOfFunctionPlot[986].X = 3.92
	pointsOfFunctionPlot[986].Y = 0.999

	pointsOfFunctionPlot[987].X = 3.93
	pointsOfFunctionPlot[987].Y = 0.999

	pointsOfFunctionPlot[988].X = 3.94
	pointsOfFunctionPlot[988].Y = 0.999

	pointsOfFunctionPlot[989].X = 3.95
	pointsOfFunctionPlot[989].Y = 0.999

	pointsOfFunctionPlot[990].X = 3.96
	pointsOfFunctionPlot[990].Y = 0.999

	pointsOfFunctionPlot[991].X = 3.97
	pointsOfFunctionPlot[991].Y = 0.999

	pointsOfFunctionPlot[992].X = 3.98
	pointsOfFunctionPlot[992].Y = 0.999

	pointsOfFunctionPlot[993].X = 3.99
	pointsOfFunctionPlot[993].Y = 0.999

	pointsOfFunctionPlot[994].X = 4.0
	pointsOfFunctionPlot[994].Y = 0.999

	pointsOfFunctionPlot[995].X = 4.01
	pointsOfFunctionPlot[995].Y = 0.999

	pointsOfFunctionPlot[996].X = 4.02
	pointsOfFunctionPlot[996].Y = 0.999

	pointsOfFunctionPlot[997].X = 4.03
	pointsOfFunctionPlot[997].Y = 0.999

	pointsOfFunctionPlot[998].X = 4.04
	pointsOfFunctionPlot[998].Y = 0.999

	pointsOfFunctionPlot[999].X = 4.05
	pointsOfFunctionPlot[999].Y = 0.999

	pointsOfFunctionPlot[1_000].X = 4.06
	pointsOfFunctionPlot[1_000].Y = 0.999

	pointsOfFunctionPlot[1_001].X = 4.07
	pointsOfFunctionPlot[1_001].Y = 0.999

	pointsOfFunctionPlot[1_002].X = 4.08
	pointsOfFunctionPlot[1_002].Y = 0.999

	pointsOfFunctionPlot[1_003].X = 4.09
	pointsOfFunctionPlot[1_003].Y = 0.999

	pointsOfFunctionPlot[1_004].X = 4.10
	pointsOfFunctionPlot[1_004].Y = 0.999

	pointsOfFunctionPlot[1_005].X = 4.11
	pointsOfFunctionPlot[1_005].Y = 0.999

	pointsOfFunctionPlot[1_006].X = 4.12
	pointsOfFunctionPlot[1_006].Y = 0.999

	pointsOfFunctionPlot[1_007].X = 4.13
	pointsOfFunctionPlot[1_007].Y = 0.999

	pointsOfFunctionPlot[1_008].X = 4.14
	pointsOfFunctionPlot[1_008].Y = 0.999

	pointsOfFunctionPlot[1_009].X = 4.15
	pointsOfFunctionPlot[1_009].Y = 0.999

	pointsOfFunctionPlot[1_010].X = 4.16
	pointsOfFunctionPlot[1_010].Y = 0.999

	pointsOfFunctionPlot[1_011].X = 4.17
	pointsOfFunctionPlot[1_011].Y = 0.999

	pointsOfFunctionPlot[1_012].X = 4.18
	pointsOfFunctionPlot[1_012].Y = 0.999

	pointsOfFunctionPlot[1_013].X = 4.19
	pointsOfFunctionPlot[1_013].Y = 0.999

	pointsOfFunctionPlot[1_014].X = 4.20
	pointsOfFunctionPlot[1_014].Y = 0.999

	pointsOfFunctionPlot[1_015].X = 4.21
	pointsOfFunctionPlot[1_015].Y = 0.999

	pointsOfFunctionPlot[1_016].X = 4.22
	pointsOfFunctionPlot[1_016].Y = 0.999

	pointsOfFunctionPlot[1_017].X = 4.23
	pointsOfFunctionPlot[1_017].Y = 0.999

	pointsOfFunctionPlot[1_018].X = 4.24
	pointsOfFunctionPlot[1_018].Y = 0.999

	pointsOfFunctionPlot[1_019].X = 4.25
	pointsOfFunctionPlot[1_019].Y = 0.999

	pointsOfFunctionPlot[1_020].X = 4.26
	pointsOfFunctionPlot[1_020].Y = 0.999

	pointsOfFunctionPlot[1_021].X = 4.27
	pointsOfFunctionPlot[1_021].Y = 0.999

	pointsOfFunctionPlot[1_022].X = 4.28
	pointsOfFunctionPlot[1_022].Y = 0.999

	pointsOfFunctionPlot[1_023].X = 4.29
	pointsOfFunctionPlot[1_023].Y = 0.999

	pointsOfFunctionPlot[1_024].X = 4.30
	pointsOfFunctionPlot[1_024].Y = 0.999

	pointsOfFunctionPlot[1_025].X = 4.31
	pointsOfFunctionPlot[1_025].Y = 0.999

	pointsOfFunctionPlot[1_026].X = 4.32
	pointsOfFunctionPlot[1_026].Y = 0.999

	pointsOfFunctionPlot[1_027].X = 4.33
	pointsOfFunctionPlot[1_027].Y = 0.999

	pointsOfFunctionPlot[1_028].X = 4.34
	pointsOfFunctionPlot[1_028].Y = 0.999

	pointsOfFunctionPlot[1_029].X = 4.35
	pointsOfFunctionPlot[1_029].Y = 0.999

	pointsOfFunctionPlot[1_030].X = 4.36
	pointsOfFunctionPlot[1_030].Y = 0.999

	pointsOfFunctionPlot[1_031].X = 4.37
	pointsOfFunctionPlot[1_031].Y = 0.999

	pointsOfFunctionPlot[1_032].X = 4.38
	pointsOfFunctionPlot[1_032].Y = 0.999

	pointsOfFunctionPlot[1_033].X = 4.39
	pointsOfFunctionPlot[1_033].Y = 0.999

	pointsOfFunctionPlot[1_034].X = 4.40
	pointsOfFunctionPlot[1_034].Y = 0.999

	pointsOfFunctionPlot[1_035].X = 4.41
	pointsOfFunctionPlot[1_035].Y = 0.999

	pointsOfFunctionPlot[1_036].X = 4.42
	pointsOfFunctionPlot[1_036].Y = 0.999

	pointsOfFunctionPlot[1_037].X = 4.43
	pointsOfFunctionPlot[1_037].Y = 0.999

	pointsOfFunctionPlot[1_038].X = 4.44
	pointsOfFunctionPlot[1_038].Y = 0.999

	pointsOfFunctionPlot[1_039].X = 4.45
	pointsOfFunctionPlot[1_039].Y = 0.999

	pointsOfFunctionPlot[1_040].X = 4.46
	pointsOfFunctionPlot[1_040].Y = 0.999

	pointsOfFunctionPlot[1_041].X = 4.47
	pointsOfFunctionPlot[1_041].Y = 0.999

	pointsOfFunctionPlot[1_042].X = 4.48
	pointsOfFunctionPlot[1_042].Y = 0.999

	pointsOfFunctionPlot[1_043].X = 4.49
	pointsOfFunctionPlot[1_043].Y = 0.999

	pointsOfFunctionPlot[1_044].X = 4.50
	pointsOfFunctionPlot[1_044].Y = 0.999

	pointsOfFunctionPlot[1_045].X = 4.51
	pointsOfFunctionPlot[1_045].Y = 0.999

	pointsOfFunctionPlot[1_046].X = 4.52
	pointsOfFunctionPlot[1_046].Y = 0.999

	pointsOfFunctionPlot[1_047].X = 4.53
	pointsOfFunctionPlot[1_047].Y = 0.999

	pointsOfFunctionPlot[1_048].X = 4.54
	pointsOfFunctionPlot[1_048].Y = 0.999

	pointsOfFunctionPlot[1_049].X = 4.55
	pointsOfFunctionPlot[1_049].Y = 0.999

	pointsOfFunctionPlot[1_050].X = 4.56
	pointsOfFunctionPlot[1_050].Y = 0.999

	pointsOfFunctionPlot[1_051].X = 4.57
	pointsOfFunctionPlot[1_051].Y = 0.999

	pointsOfFunctionPlot[1_052].X = 4.58
	pointsOfFunctionPlot[1_052].Y = 0.999

	pointsOfFunctionPlot[1_053].X = 4.59
	pointsOfFunctionPlot[1_053].Y = 0.999

	pointsOfFunctionPlot[1_054].X = 4.60
	pointsOfFunctionPlot[1_054].Y = 0.999

	pointsOfFunctionPlot[1_055].X = 4.61
	pointsOfFunctionPlot[1_055].Y = 0.999

	pointsOfFunctionPlot[1_056].X = 4.62
	pointsOfFunctionPlot[1_056].Y = 0.999

	pointsOfFunctionPlot[1_057].X = 4.63
	pointsOfFunctionPlot[1_057].Y = 0.999

	pointsOfFunctionPlot[1_058].X = 4.64
	pointsOfFunctionPlot[1_058].Y = 0.999

	pointsOfFunctionPlot[1_059].X = 4.65
	pointsOfFunctionPlot[1_059].Y = 0.999

	pointsOfFunctionPlot[1_060].X = 4.66
	pointsOfFunctionPlot[1_060].Y = 0.999

	pointsOfFunctionPlot[1_061].X = 4.67
	pointsOfFunctionPlot[1_061].Y = 0.999

	pointsOfFunctionPlot[1_062].X = 4.68
	pointsOfFunctionPlot[1_062].Y = 0.999

	pointsOfFunctionPlot[1_063].X = 4.69
	pointsOfFunctionPlot[1_063].Y = 0.999

	pointsOfFunctionPlot[1_064].X = 4.70
	pointsOfFunctionPlot[1_064].Y = 0.999

	pointsOfFunctionPlot[1_065].X = 4.71
	pointsOfFunctionPlot[1_065].Y = 0.999

	pointsOfFunctionPlot[1_066].X = 4.72
	pointsOfFunctionPlot[1_066].Y = 0.999

	pointsOfFunctionPlot[1_067].X = 4.73
	pointsOfFunctionPlot[1_067].Y = 0.999

	pointsOfFunctionPlot[1_068].X = 4.74
	pointsOfFunctionPlot[1_068].Y = 0.999

	pointsOfFunctionPlot[1_069].X = 4.75
	pointsOfFunctionPlot[1_069].Y = 0.999

	pointsOfFunctionPlot[1_070].X = 4.76
	pointsOfFunctionPlot[1_070].Y = 0.999

	pointsOfFunctionPlot[1_071].X = 4.77
	pointsOfFunctionPlot[1_071].Y = 0.999

	pointsOfFunctionPlot[1_072].X = 4.78
	pointsOfFunctionPlot[1_072].Y = 0.999

	pointsOfFunctionPlot[1_073].X = 4.79
	pointsOfFunctionPlot[1_073].Y = 0.999

	pointsOfFunctionPlot[1_074].X = 4.80
	pointsOfFunctionPlot[1_074].Y = 0.999

	pointsOfFunctionPlot[1_075].X = 4.81
	pointsOfFunctionPlot[1_075].Y = 0.999

	pointsOfFunctionPlot[1_076].X = 4.82
	pointsOfFunctionPlot[1_076].Y = 0.999

	pointsOfFunctionPlot[1_077].X = 4.83
	pointsOfFunctionPlot[1_077].Y = 0.999

	pointsOfFunctionPlot[1_078].X = 4.84
	pointsOfFunctionPlot[1_078].Y = 0.999

	pointsOfFunctionPlot[1_079].X = 4.85
	pointsOfFunctionPlot[1_079].Y = 0.999

	pointsOfFunctionPlot[1_080].X = 4.86
	pointsOfFunctionPlot[1_080].Y = 0.999

	pointsOfFunctionPlot[1_081].X = 4.87
	pointsOfFunctionPlot[1_081].Y = 0.999

	pointsOfFunctionPlot[1_082].X = 4.88
	pointsOfFunctionPlot[1_082].Y = 0.999

	pointsOfFunctionPlot[1_083].X = 4.89
	pointsOfFunctionPlot[1_083].Y = 0.999

	pointsOfFunctionPlot[1_084].X = 4.90
	pointsOfFunctionPlot[1_084].Y = 0.999

	pointsOfFunctionPlot[1_085].X = 4.91
	pointsOfFunctionPlot[1_085].Y = 0.999

	pointsOfFunctionPlot[1_086].X = 4.92
	pointsOfFunctionPlot[1_086].Y = 0.999

	pointsOfFunctionPlot[1_087].X = 4.93
	pointsOfFunctionPlot[1_087].Y = 0.999

	pointsOfFunctionPlot[1_088].X = 4.94
	pointsOfFunctionPlot[1_088].Y = 0.999

	pointsOfFunctionPlot[1_089].X = 4.95
	pointsOfFunctionPlot[1_089].Y = 0.999

	pointsOfFunctionPlot[1_090].X = 4.96
	pointsOfFunctionPlot[1_090].Y = 0.999

	pointsOfFunctionPlot[1_091].X = 4.97
	pointsOfFunctionPlot[1_091].Y = 0.999

	pointsOfFunctionPlot[1_092].X = 4.98
	pointsOfFunctionPlot[1_092].Y = 0.999

	pointsOfFunctionPlot[1_093].X = 4.99
	pointsOfFunctionPlot[1_093].Y = 0.999

	pointsOfFunctionPlot[1_094].X = 5.0
	pointsOfFunctionPlot[1_094].Y = 0.999

	pointsOfFunctionPlot[1_095].X = 5.01
	pointsOfFunctionPlot[1_095].Y = 0.999

	pointsOfFunctionPlot[1_096].X = 5.02
	pointsOfFunctionPlot[1_096].Y = 0.999

	pointsOfFunctionPlot[1_097].X = 5.03
	pointsOfFunctionPlot[1_097].Y = 0.999

	pointsOfFunctionPlot[1_098].X = 5.04
	pointsOfFunctionPlot[1_098].Y = 0.999

	pointsOfFunctionPlot[1_099].X = 5.05
	pointsOfFunctionPlot[1_099].Y = 0.999

	pointsOfFunctionPlot[1_100].X = 5.06
	pointsOfFunctionPlot[1_100].Y = 0.999

	pointsOfFunctionPlot[1_101].X = 5.07
	pointsOfFunctionPlot[1_101].Y = 0.999

	pointsOfFunctionPlot[1_102].X = 5.08
	pointsOfFunctionPlot[1_102].Y = 0.999

	pointsOfFunctionPlot[1_103].X = 5.09
	pointsOfFunctionPlot[1_103].Y = 0.999

	pointsOfFunctionPlot[1_104].X = 5.10
	pointsOfFunctionPlot[1_104].Y = 0.999

	pointsOfFunctionPlot[1_105].X = 5.11
	pointsOfFunctionPlot[1_105].Y = 0.999

	pointsOfFunctionPlot[1_106].X = 5.12
	pointsOfFunctionPlot[1_106].Y = 0.999

	pointsOfFunctionPlot[1_107].X = 5.13
	pointsOfFunctionPlot[1_107].Y = 0.999

	pointsOfFunctionPlot[1_108].X = 5.14
	pointsOfFunctionPlot[1_108].Y = 0.999

	pointsOfFunctionPlot[1_109].X = 5.15
	pointsOfFunctionPlot[1_109].Y = 0.999

	pointsOfFunctionPlot[1_110].X = 5.16
	pointsOfFunctionPlot[1_110].Y = 0.999

	pointsOfFunctionPlot[1_111].X = 5.17
	pointsOfFunctionPlot[1_111].Y = 0.999

	pointsOfFunctionPlot[1_112].X = 5.18
	pointsOfFunctionPlot[1_112].Y = 0.999

	pointsOfFunctionPlot[1_113].X = 5.19
	pointsOfFunctionPlot[1_113].Y = 0.999

	pointsOfFunctionPlot[1_114].X = 5.20
	pointsOfFunctionPlot[1_114].Y = 0.999

	pointsOfFunctionPlot[1_115].X = 5.21
	pointsOfFunctionPlot[1_115].Y = 0.999

	pointsOfFunctionPlot[1_116].X = 5.22
	pointsOfFunctionPlot[1_116].Y = 0.999

	pointsOfFunctionPlot[1_117].X = 5.23
	pointsOfFunctionPlot[1_117].Y = 0.999

	pointsOfFunctionPlot[1_118].X = 5.24
	pointsOfFunctionPlot[1_118].Y = 0.999

	pointsOfFunctionPlot[1_119].X = 5.25
	pointsOfFunctionPlot[1_119].Y = 0.999

	pointsOfFunctionPlot[1_120].X = 5.26
	pointsOfFunctionPlot[1_120].Y = 0.999

	pointsOfFunctionPlot[1_121].X = 5.27
	pointsOfFunctionPlot[1_121].Y = 0.999

	pointsOfFunctionPlot[1_122].X = 5.28
	pointsOfFunctionPlot[1_122].Y = 0.999

	pointsOfFunctionPlot[1_123].X = 5.29
	pointsOfFunctionPlot[1_123].Y = 0.999

	pointsOfFunctionPlot[1_124].X = 5.30
	pointsOfFunctionPlot[1_124].Y = 0.999

	pointsOfFunctionPlot[1_125].X = 5.31
	pointsOfFunctionPlot[1_125].Y = 0.999

	pointsOfFunctionPlot[1_126].X = 5.32
	pointsOfFunctionPlot[1_126].Y = 0.999

	pointsOfFunctionPlot[1_127].X = 5.33
	pointsOfFunctionPlot[1_127].Y = 0.999

	pointsOfFunctionPlot[1_128].X = 5.34
	pointsOfFunctionPlot[1_128].Y = 0.999

	pointsOfFunctionPlot[1_129].X = 5.35
	pointsOfFunctionPlot[1_129].Y = 0.999

	pointsOfFunctionPlot[1_130].X = 5.36
	pointsOfFunctionPlot[1_130].Y = 0.999

	pointsOfFunctionPlot[1_131].X = 5.37
	pointsOfFunctionPlot[1_131].Y = 0.999

	pointsOfFunctionPlot[1_132].X = 5.38
	pointsOfFunctionPlot[1_132].Y = 0.999

	pointsOfFunctionPlot[1_133].X = 5.39
	pointsOfFunctionPlot[1_133].Y = 0.999

	pointsOfFunctionPlot[1_134].X = 5.40
	pointsOfFunctionPlot[1_134].Y = 0.999

	pointsOfFunctionPlot[1_135].X = 5.41
	pointsOfFunctionPlot[1_135].Y = 0.999

	pointsOfFunctionPlot[1_136].X = 5.42
	pointsOfFunctionPlot[1_136].Y = 0.999

	pointsOfFunctionPlot[1_137].X = 5.43
	pointsOfFunctionPlot[1_137].Y = 0.999

	pointsOfFunctionPlot[1_138].X = 5.44
	pointsOfFunctionPlot[1_138].Y = 0.999

	pointsOfFunctionPlot[1_139].X = 5.45
	pointsOfFunctionPlot[1_139].Y = 0.999

	pointsOfFunctionPlot[1_140].X = 5.46
	pointsOfFunctionPlot[1_140].Y = 0.999

	pointsOfFunctionPlot[1_141].X = 5.47
	pointsOfFunctionPlot[1_141].Y = 0.999

	pointsOfFunctionPlot[1_142].X = 5.48
	pointsOfFunctionPlot[1_142].Y = 0.999

	pointsOfFunctionPlot[1_143].X = 5.49
	pointsOfFunctionPlot[1_143].Y = 0.999

	pointsOfFunctionPlot[1_144].X = 5.50
	pointsOfFunctionPlot[1_144].Y = 0.999

	pointsOfFunctionPlot[1_145].X = 5.51
	pointsOfFunctionPlot[1_145].Y = 0.999

	pointsOfFunctionPlot[1_146].X = 5.52
	pointsOfFunctionPlot[1_146].Y = 0.999

	pointsOfFunctionPlot[1_147].X = 5.53
	pointsOfFunctionPlot[1_147].Y = 0.999

	pointsOfFunctionPlot[1_148].X = 5.54
	pointsOfFunctionPlot[1_148].Y = 0.999

	pointsOfFunctionPlot[1_149].X = 5.55
	pointsOfFunctionPlot[1_149].Y = 0.999

	pointsOfFunctionPlot[1_150].X = 5.56
	pointsOfFunctionPlot[1_150].Y = 0.999

	pointsOfFunctionPlot[1_151].X = 5.57
	pointsOfFunctionPlot[1_151].Y = 0.999

	pointsOfFunctionPlot[1_152].X = 5.58
	pointsOfFunctionPlot[1_152].Y = 0.999

	pointsOfFunctionPlot[1_153].X = 5.59
	pointsOfFunctionPlot[1_153].Y = 0.999

	pointsOfFunctionPlot[1_154].X = 5.60
	pointsOfFunctionPlot[1_154].Y = 0.999

	pointsOfFunctionPlot[1_155].X = 5.61
	pointsOfFunctionPlot[1_155].Y = 0.999

	pointsOfFunctionPlot[1_156].X = 5.62
	pointsOfFunctionPlot[1_156].Y = 0.999

	pointsOfFunctionPlot[1_157].X = 5.63
	pointsOfFunctionPlot[1_157].Y = 0.999

	pointsOfFunctionPlot[1_158].X = 5.64
	pointsOfFunctionPlot[1_158].Y = 0.999

	pointsOfFunctionPlot[1_159].X = 5.65
	pointsOfFunctionPlot[1_159].Y = 0.999

	pointsOfFunctionPlot[1_160].X = 5.66
	pointsOfFunctionPlot[1_160].Y = 0.999

	pointsOfFunctionPlot[1_161].X = 5.67
	pointsOfFunctionPlot[1_161].Y = 0.999

	pointsOfFunctionPlot[1_162].X = 5.68
	pointsOfFunctionPlot[1_162].Y = 0.999

	pointsOfFunctionPlot[1_163].X = 5.69
	pointsOfFunctionPlot[1_163].Y = 0.999

	pointsOfFunctionPlot[1_164].X = 5.70
	pointsOfFunctionPlot[1_164].Y = 0.999

	pointsOfFunctionPlot[1_165].X = 5.71
	pointsOfFunctionPlot[1_165].Y = 0.999

	pointsOfFunctionPlot[1_166].X = 5.72
	pointsOfFunctionPlot[1_166].Y = 0.999

	pointsOfFunctionPlot[1_167].X = 5.73
	pointsOfFunctionPlot[1_167].Y = 0.999

	pointsOfFunctionPlot[1_168].X = 5.74
	pointsOfFunctionPlot[1_168].Y = 0.999

	pointsOfFunctionPlot[1_169].X = 5.75
	pointsOfFunctionPlot[1_169].Y = 0.999

	pointsOfFunctionPlot[1_170].X = 5.76
	pointsOfFunctionPlot[1_170].Y = 0.999

	pointsOfFunctionPlot[1_171].X = 5.77
	pointsOfFunctionPlot[1_171].Y = 0.999

	pointsOfFunctionPlot[1_172].X = 5.78
	pointsOfFunctionPlot[1_172].Y = 0.999

	pointsOfFunctionPlot[1_173].X = 5.79
	pointsOfFunctionPlot[1_173].Y = 0.999

	pointsOfFunctionPlot[1_174].X = 5.80
	pointsOfFunctionPlot[1_174].Y = 0.999

	pointsOfFunctionPlot[1_175].X = 5.81
	pointsOfFunctionPlot[1_175].Y = 0.999

	pointsOfFunctionPlot[1_176].X = 5.82
	pointsOfFunctionPlot[1_176].Y = 0.999

	pointsOfFunctionPlot[1_177].X = 5.83
	pointsOfFunctionPlot[1_177].Y = 0.999

	pointsOfFunctionPlot[1_178].X = 5.84
	pointsOfFunctionPlot[1_178].Y = 0.999

	pointsOfFunctionPlot[1_179].X = 5.85
	pointsOfFunctionPlot[1_179].Y = 0.999

	pointsOfFunctionPlot[1_180].X = 5.86
	pointsOfFunctionPlot[1_180].Y = 0.999

	pointsOfFunctionPlot[1_181].X = 5.87
	pointsOfFunctionPlot[1_181].Y = 0.999

	pointsOfFunctionPlot[1_182].X = 5.88
	pointsOfFunctionPlot[1_182].Y = 0.999

	pointsOfFunctionPlot[1_183].X = 5.89
	pointsOfFunctionPlot[1_183].Y = 0.999

	pointsOfFunctionPlot[1_184].X = 5.90
	pointsOfFunctionPlot[1_184].Y = 0.999

	pointsOfFunctionPlot[1_185].X = 5.91
	pointsOfFunctionPlot[1_185].Y = 0.999

	pointsOfFunctionPlot[1_186].X = 5.92
	pointsOfFunctionPlot[1_186].Y = 0.999

	pointsOfFunctionPlot[1_187].X = 5.93
	pointsOfFunctionPlot[1_187].Y = 1.0

	pointsOfFunctionPlot[1_188].X = 10.0
	pointsOfFunctionPlot[1_188].Y = 1.0








	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function erf(x)"

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
		"erf-function-plot-01.png"); err != nil {

		panic(err)
	}
}
