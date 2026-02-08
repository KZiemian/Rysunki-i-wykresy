package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function 2^x.

	pointsOfFunctionPlot := make(plotter.XYs, 2_001)

	pointsOfFunctionPlot[0].X = -10.0
	pointsOfFunctionPlot[0].Y = 0.001

	pointsOfFunctionPlot[1].X = -9.99
	pointsOfFunctionPlot[1].Y = 0.001

	pointsOfFunctionPlot[2].X = -9.98
	pointsOfFunctionPlot[2].Y = 0.001

	pointsOfFunctionPlot[3].X = -9.97
	pointsOfFunctionPlot[3].Y = 0.001

	pointsOfFunctionPlot[4].X = -9.96
	pointsOfFunctionPlot[4].Y = 0.001

	pointsOfFunctionPlot[5].X = -9.95
	pointsOfFunctionPlot[5].Y = 0.001

	pointsOfFunctionPlot[6].X = -9.94
	pointsOfFunctionPlot[6].Y = 0.001

	pointsOfFunctionPlot[7].X = -9.93
	pointsOfFunctionPlot[7].Y = 0.001

	pointsOfFunctionPlot[8].X = -9.92
	pointsOfFunctionPlot[8].Y = 0.001

	pointsOfFunctionPlot[9].X = -9.91
	pointsOfFunctionPlot[9].Y = 0.001

	pointsOfFunctionPlot[10].X = -9.90
	pointsOfFunctionPlot[10].Y = 0.001

	pointsOfFunctionPlot[11].X = -9.89
	pointsOfFunctionPlot[11].Y = 0.001

	pointsOfFunctionPlot[12].X = -9.88
	pointsOfFunctionPlot[12].Y = 0.001

	pointsOfFunctionPlot[13].X = -9.87
	pointsOfFunctionPlot[13].Y = 0.001

	pointsOfFunctionPlot[14].X = -9.86
	pointsOfFunctionPlot[14].Y = 0.001

	pointsOfFunctionPlot[15].X = -9.85
	pointsOfFunctionPlot[15].Y = 0.001

	pointsOfFunctionPlot[16].X = -9.84
	pointsOfFunctionPlot[16].Y = 0.001

	pointsOfFunctionPlot[17].X = -9.83
	pointsOfFunctionPlot[17].Y = 0.001

	pointsOfFunctionPlot[18].X = -9.82
	pointsOfFunctionPlot[18].Y = 0.001

	pointsOfFunctionPlot[19].X = -9.81
	pointsOfFunctionPlot[19].Y = 0.001

	pointsOfFunctionPlot[20].X = -9.80
	pointsOfFunctionPlot[20].Y = 0.001

	pointsOfFunctionPlot[21].X = -9.79
	pointsOfFunctionPlot[21].Y = 0.001

	pointsOfFunctionPlot[22].X = -9.78
	pointsOfFunctionPlot[22].Y = 0.001

	pointsOfFunctionPlot[23].X = -9.77
	pointsOfFunctionPlot[23].Y = 0.001

	pointsOfFunctionPlot[24].X = -9.76
	pointsOfFunctionPlot[24].Y = 0.001

	pointsOfFunctionPlot[25].X = -9.75
	pointsOfFunctionPlot[25].Y = 0.001

	pointsOfFunctionPlot[26].X = -9.74
	pointsOfFunctionPlot[26].Y = 0.001

	pointsOfFunctionPlot[27].X = -9.73
	pointsOfFunctionPlot[27].Y = 0.001

	pointsOfFunctionPlot[28].X = -9.72
	pointsOfFunctionPlot[28].Y = 0.001

	pointsOfFunctionPlot[29].X = -9.71
	pointsOfFunctionPlot[29].Y = 0.001

	pointsOfFunctionPlot[30].X = -9.70
	pointsOfFunctionPlot[30].Y = 0.001

	pointsOfFunctionPlot[31].X = -9.69
	pointsOfFunctionPlot[31].Y = 0.001

	pointsOfFunctionPlot[32].X = -9.68
	pointsOfFunctionPlot[32].Y = 0.001

	pointsOfFunctionPlot[33].X = -9.67
	pointsOfFunctionPlot[33].Y = 0.001

	pointsOfFunctionPlot[34].X = -9.66
	pointsOfFunctionPlot[34].Y = 0.001

	pointsOfFunctionPlot[35].X = -9.65
	pointsOfFunctionPlot[35].Y = 0.001

	pointsOfFunctionPlot[36].X = -9.64
	pointsOfFunctionPlot[36].Y = 0.001

	pointsOfFunctionPlot[37].X = -9.63
	pointsOfFunctionPlot[37].Y = 0.001

	pointsOfFunctionPlot[38].X = -9.62
	pointsOfFunctionPlot[38].Y = 0.001

	pointsOfFunctionPlot[39].X = -9.61
	pointsOfFunctionPlot[39].Y = 0.001

	pointsOfFunctionPlot[40].X = -9.60
	pointsOfFunctionPlot[40].Y = 0.001

	pointsOfFunctionPlot[41].X = -9.59
	pointsOfFunctionPlot[41].Y = 0.001

	pointsOfFunctionPlot[42].X = -9.58
	pointsOfFunctionPlot[42].Y = 0.001

	pointsOfFunctionPlot[43].X = -9.57
	pointsOfFunctionPlot[43].Y = 0.001

	pointsOfFunctionPlot[44].X = -9.56
	pointsOfFunctionPlot[44].Y = 0.001

	pointsOfFunctionPlot[45].X = -9.55
	pointsOfFunctionPlot[45].Y = 0.001

	pointsOfFunctionPlot[46].X = -9.54
	pointsOfFunctionPlot[46].Y = 0.001

	pointsOfFunctionPlot[47].X = -9.53
	pointsOfFunctionPlot[47].Y = 0.001

	pointsOfFunctionPlot[48].X = -9.52
	pointsOfFunctionPlot[48].Y = 0.001

	pointsOfFunctionPlot[49].X = -9.51
	pointsOfFunctionPlot[49].Y = 0.001

	pointsOfFunctionPlot[50].X = -9.50
	pointsOfFunctionPlot[50].Y = 0.001

	pointsOfFunctionPlot[51].X = -9.49
	pointsOfFunctionPlot[51].Y = 0.001

	pointsOfFunctionPlot[52].X = -9.48
	pointsOfFunctionPlot[52].Y = 0.001

	pointsOfFunctionPlot[53].X = -9.47
	pointsOfFunctionPlot[53].Y = 0.001

	pointsOfFunctionPlot[54].X = -9.46
	pointsOfFunctionPlot[54].Y = 0.001

	pointsOfFunctionPlot[55].X = -9.45
	pointsOfFunctionPlot[55].Y = 0.001

	pointsOfFunctionPlot[56].X = -9.44
	pointsOfFunctionPlot[56].Y = 0.001

	pointsOfFunctionPlot[57].X = -9.43
	pointsOfFunctionPlot[57].Y = 0.001

	pointsOfFunctionPlot[58].X = -9.42
	pointsOfFunctionPlot[58].Y = 0.001

	pointsOfFunctionPlot[59].X = -9.41
	pointsOfFunctionPlot[59].Y = 0.001

	pointsOfFunctionPlot[60].X = -9.40
	pointsOfFunctionPlot[60].Y = 0.001

	pointsOfFunctionPlot[61].X = -9.39
	pointsOfFunctionPlot[61].Y = 0.001

	pointsOfFunctionPlot[62].X = -9.38
	pointsOfFunctionPlot[62].Y = 0.001

	pointsOfFunctionPlot[63].X = -9.37
	pointsOfFunctionPlot[63].Y = 0.001

	pointsOfFunctionPlot[64].X = -9.36
	pointsOfFunctionPlot[64].Y = 0.001

	pointsOfFunctionPlot[65].X = -9.35
	pointsOfFunctionPlot[65].Y = 0.001

	pointsOfFunctionPlot[66].X = -9.34
	pointsOfFunctionPlot[66].Y = 0.001

	pointsOfFunctionPlot[67].X = -9.33
	pointsOfFunctionPlot[67].Y = 0.001

	pointsOfFunctionPlot[68].X = -9.32
	pointsOfFunctionPlot[68].Y = 0.001

	pointsOfFunctionPlot[69].X = -9.31
	pointsOfFunctionPlot[69].Y = 0.001

	pointsOfFunctionPlot[70].X = -9.30
	pointsOfFunctionPlot[70].Y = 0.001

	pointsOfFunctionPlot[71].X = -9.29
	pointsOfFunctionPlot[71].Y = 0.001

	pointsOfFunctionPlot[72].X = -9.28
	pointsOfFunctionPlot[72].Y = 0.001

	pointsOfFunctionPlot[73].X = -9.27
	pointsOfFunctionPlot[73].Y = 0.001

	pointsOfFunctionPlot[74].X = -9.26
	pointsOfFunctionPlot[74].Y = 0.001

	pointsOfFunctionPlot[75].X = -9.25
	pointsOfFunctionPlot[75].Y = 0.001

	pointsOfFunctionPlot[76].X = -9.24
	pointsOfFunctionPlot[76].Y = 0.001

	pointsOfFunctionPlot[77].X = -9.23
	pointsOfFunctionPlot[77].Y = 0.001

	pointsOfFunctionPlot[78].X = -9.22
	pointsOfFunctionPlot[78].Y = 0.001

	pointsOfFunctionPlot[79].X = -9.21
	pointsOfFunctionPlot[79].Y = 0.001

	pointsOfFunctionPlot[80].X = -9.20
	pointsOfFunctionPlot[80].Y = 0.001

	pointsOfFunctionPlot[81].X = -9.19
	pointsOfFunctionPlot[81].Y = 0.001

	pointsOfFunctionPlot[82].X = -9.18
	pointsOfFunctionPlot[82].Y = 0.001

	pointsOfFunctionPlot[83].X = -9.17
	pointsOfFunctionPlot[83].Y = 0.001

	pointsOfFunctionPlot[84].X = -9.16
	pointsOfFunctionPlot[84].Y = 0.001

	pointsOfFunctionPlot[85].X = -9.15
	pointsOfFunctionPlot[85].Y = 0.001

	pointsOfFunctionPlot[86].X = -9.14
	pointsOfFunctionPlot[86].Y = 0.001

	pointsOfFunctionPlot[87].X = -9.13
	pointsOfFunctionPlot[87].Y = 0.001

	pointsOfFunctionPlot[88].X = -9.12
	pointsOfFunctionPlot[88].Y = 0.001

	pointsOfFunctionPlot[89].X = -9.11
	pointsOfFunctionPlot[89].Y = 0.001

	pointsOfFunctionPlot[90].X = -9.10
	pointsOfFunctionPlot[90].Y = 0.001

	pointsOfFunctionPlot[91].X = -9.09
	pointsOfFunctionPlot[91].Y = 0.001

	pointsOfFunctionPlot[92].X = -9.08
	pointsOfFunctionPlot[92].Y = 0.001

	pointsOfFunctionPlot[93].X = -9.07
	pointsOfFunctionPlot[93].Y = 0.001

	pointsOfFunctionPlot[94].X = -9.06
	pointsOfFunctionPlot[94].Y = 0.001

	pointsOfFunctionPlot[95].X = -9.05
	pointsOfFunctionPlot[95].Y = 0.001

	pointsOfFunctionPlot[96].X = -9.04
	pointsOfFunctionPlot[96].Y = 0.001

	pointsOfFunctionPlot[97].X = -9.03
	pointsOfFunctionPlot[97].Y = 0.001

	pointsOfFunctionPlot[98].X = -9.02
	pointsOfFunctionPlot[98].Y = 0.001

	pointsOfFunctionPlot[99].X = -9.01
	pointsOfFunctionPlot[99].Y = 0.001

	pointsOfFunctionPlot[100].X = -9.0
	pointsOfFunctionPlot[100].Y = 0.002

	pointsOfFunctionPlot[101].X = -8.99
	pointsOfFunctionPlot[101].Y = 0.002

	pointsOfFunctionPlot[102].X = -8.98
	pointsOfFunctionPlot[102].Y = 0.002

	pointsOfFunctionPlot[103].X = -8.97
	pointsOfFunctionPlot[103].Y = 0.002

	pointsOfFunctionPlot[104].X = -8.96
	pointsOfFunctionPlot[104].Y = 0.002

	pointsOfFunctionPlot[105].X = -8.95
	pointsOfFunctionPlot[105].Y = 0.002

	pointsOfFunctionPlot[106].X = -8.94
	pointsOfFunctionPlot[106].Y = 0.002

	pointsOfFunctionPlot[107].X = -8.93
	pointsOfFunctionPlot[107].Y = 0.002

	pointsOfFunctionPlot[108].X = -8.92
	pointsOfFunctionPlot[108].Y = 0.002

	pointsOfFunctionPlot[109].X = -8.91
	pointsOfFunctionPlot[109].Y = 0.002

	pointsOfFunctionPlot[110].X = -8.90
	pointsOfFunctionPlot[110].Y = 0.002

	pointsOfFunctionPlot[111].X = -8.89
	pointsOfFunctionPlot[111].Y = 0.002

	pointsOfFunctionPlot[112].X = -8.88
	pointsOfFunctionPlot[112].Y = 0.002

	pointsOfFunctionPlot[113].X = -8.87
	pointsOfFunctionPlot[113].Y = 0.002

	pointsOfFunctionPlot[114].X = -8.86
	pointsOfFunctionPlot[114].Y = 0.002

	pointsOfFunctionPlot[115].X = -8.85
	pointsOfFunctionPlot[115].Y = 0.002

	pointsOfFunctionPlot[116].X = -8.84
	pointsOfFunctionPlot[116].Y = 0.002

	pointsOfFunctionPlot[117].X = -8.83
	pointsOfFunctionPlot[117].Y = 0.002

	pointsOfFunctionPlot[118].X = -8.82
	pointsOfFunctionPlot[118].Y = 0.002

	pointsOfFunctionPlot[119].X = -8.81
	pointsOfFunctionPlot[119].Y = 0.002

	pointsOfFunctionPlot[120].X = -8.80
	pointsOfFunctionPlot[120].Y = 0.002

	pointsOfFunctionPlot[121].X = -8.79
	pointsOfFunctionPlot[121].Y = 0.002

	pointsOfFunctionPlot[122].X = -8.78
	pointsOfFunctionPlot[122].Y = 0.002

	pointsOfFunctionPlot[123].X = -8.77
	pointsOfFunctionPlot[123].Y = 0.002

	pointsOfFunctionPlot[124].X = -8.76
	pointsOfFunctionPlot[124].Y = 0.002

	pointsOfFunctionPlot[125].X = -8.75
	pointsOfFunctionPlot[125].Y = 0.002

	pointsOfFunctionPlot[126].X = -8.74
	pointsOfFunctionPlot[126].Y = 0.002

	pointsOfFunctionPlot[127].X = -8.73
	pointsOfFunctionPlot[127].Y = 0.002

	pointsOfFunctionPlot[128].X = -8.72
	pointsOfFunctionPlot[128].Y = 0.002

	pointsOfFunctionPlot[129].X = -8.71
	pointsOfFunctionPlot[129].Y = 0.002

	pointsOfFunctionPlot[130].X = -8.70
	pointsOfFunctionPlot[130].Y = 0.002

	pointsOfFunctionPlot[131].X = -8.69
	pointsOfFunctionPlot[131].Y = 0.002

	pointsOfFunctionPlot[132].X = -8.68
	pointsOfFunctionPlot[132].Y = 0.002

	pointsOfFunctionPlot[133].X = -8.67
	pointsOfFunctionPlot[133].Y = 0.002

	pointsOfFunctionPlot[134].X = -8.66
	pointsOfFunctionPlot[134].Y = 0.002

	pointsOfFunctionPlot[135].X = -8.65
	pointsOfFunctionPlot[135].Y = 0.002

	pointsOfFunctionPlot[136].X = -8.64
	pointsOfFunctionPlot[136].Y = 0.002

	pointsOfFunctionPlot[137].X = -8.63
	pointsOfFunctionPlot[137].Y = 0.002

	pointsOfFunctionPlot[138].X = -8.62
	pointsOfFunctionPlot[138].Y = 0.002

	pointsOfFunctionPlot[139].X = -8.61
	pointsOfFunctionPlot[139].Y = 0.002

	pointsOfFunctionPlot[140].X = -8.60
	pointsOfFunctionPlot[140].Y = 0.002

	pointsOfFunctionPlot[141].X = -8.59
	pointsOfFunctionPlot[141].Y = 0.002

	pointsOfFunctionPlot[142].X = -8.58
	pointsOfFunctionPlot[142].Y = 0.002

	pointsOfFunctionPlot[143].X = -8.57
	pointsOfFunctionPlot[143].Y = 0.002

	pointsOfFunctionPlot[144].X = -8.56
	pointsOfFunctionPlot[144].Y = 0.002

	pointsOfFunctionPlot[145].X = -8.55
	pointsOfFunctionPlot[145].Y = 0.002

	pointsOfFunctionPlot[146].X = -8.54
	pointsOfFunctionPlot[146].Y = 0.002

	pointsOfFunctionPlot[147].X = -8.53
	pointsOfFunctionPlot[147].Y = 0.002

	pointsOfFunctionPlot[148].X = -8.52
	pointsOfFunctionPlot[148].Y = 0.002

	pointsOfFunctionPlot[149].X = -8.51
	pointsOfFunctionPlot[149].Y = 0.002

	pointsOfFunctionPlot[150].X = -8.50
	pointsOfFunctionPlot[150].Y = 0.002

	pointsOfFunctionPlot[151].X = -8.49
	pointsOfFunctionPlot[151].Y = 0.002

	pointsOfFunctionPlot[152].X = -8.48
	pointsOfFunctionPlot[152].Y = 0.002

	pointsOfFunctionPlot[153].X = -8.47
	pointsOfFunctionPlot[153].Y = 0.002

	pointsOfFunctionPlot[154].X = -8.46
	pointsOfFunctionPlot[154].Y = 0.002

	pointsOfFunctionPlot[155].X = -8.45
	pointsOfFunctionPlot[155].Y = 0.002

	pointsOfFunctionPlot[156].X = -8.44
	pointsOfFunctionPlot[156].Y = 0.002

	pointsOfFunctionPlot[157].X = -8.43
	pointsOfFunctionPlot[157].Y = 0.002

	pointsOfFunctionPlot[158].X = -8.42
	pointsOfFunctionPlot[158].Y = 0.002

	pointsOfFunctionPlot[159].X = -8.41
	pointsOfFunctionPlot[159].Y = 0.002

	pointsOfFunctionPlot[160].X = -8.40
	pointsOfFunctionPlot[160].Y = 0.003

	pointsOfFunctionPlot[161].X = -8.39
	pointsOfFunctionPlot[161].Y = 0.003

	pointsOfFunctionPlot[162].X = -8.38
	pointsOfFunctionPlot[162].Y = 0.003

	pointsOfFunctionPlot[163].X = -8.37
	pointsOfFunctionPlot[163].Y = 0.003

	pointsOfFunctionPlot[164].X = -8.36
	pointsOfFunctionPlot[164].Y = 0.003

	pointsOfFunctionPlot[165].X = -8.35
	pointsOfFunctionPlot[165].Y = 0.003

	pointsOfFunctionPlot[166].X = -8.34
	pointsOfFunctionPlot[166].Y = 0.003

	pointsOfFunctionPlot[167].X = -8.33
	pointsOfFunctionPlot[167].Y = 0.003

	pointsOfFunctionPlot[168].X = -8.32
	pointsOfFunctionPlot[168].Y = 0.003

	pointsOfFunctionPlot[169].X = -8.31
	pointsOfFunctionPlot[169].Y = 0.003

	pointsOfFunctionPlot[170].X = -8.30
	pointsOfFunctionPlot[170].Y = 0.003

	pointsOfFunctionPlot[171].X = -8.29
	pointsOfFunctionPlot[171].Y = 0.003

	pointsOfFunctionPlot[172].X = -8.28
	pointsOfFunctionPlot[172].Y = 0.003

	pointsOfFunctionPlot[173].X = -8.27
	pointsOfFunctionPlot[173].Y = 0.003

	pointsOfFunctionPlot[174].X = -8.26
	pointsOfFunctionPlot[174].Y = 0.003

	pointsOfFunctionPlot[175].X = -8.25
	pointsOfFunctionPlot[175].Y = 0.003

	pointsOfFunctionPlot[176].X = -8.24
	pointsOfFunctionPlot[176].Y = 0.003

	pointsOfFunctionPlot[177].X = -8.23
	pointsOfFunctionPlot[177].Y = 0.003

	pointsOfFunctionPlot[178].X = -8.22
	pointsOfFunctionPlot[178].Y = 0.003

	pointsOfFunctionPlot[179].X = -8.21
	pointsOfFunctionPlot[179].Y = 0.003

	pointsOfFunctionPlot[180].X = -8.20
	pointsOfFunctionPlot[180].Y = 0.003

	pointsOfFunctionPlot[181].X = -8.19
	pointsOfFunctionPlot[181].Y = 0.003

	pointsOfFunctionPlot[182].X = -8.18
	pointsOfFunctionPlot[182].Y = 0.003

	pointsOfFunctionPlot[183].X = -8.17
	pointsOfFunctionPlot[183].Y = 0.003

	pointsOfFunctionPlot[184].X = -8.16
	pointsOfFunctionPlot[184].Y = 0.003

	pointsOfFunctionPlot[185].X = -8.15
	pointsOfFunctionPlot[185].Y = 0.003

	pointsOfFunctionPlot[186].X = -8.14
	pointsOfFunctionPlot[186].Y = 0.003

	pointsOfFunctionPlot[187].X = -8.13
	pointsOfFunctionPlot[187].Y = 0.003

	pointsOfFunctionPlot[188].X = -8.12
	pointsOfFunctionPlot[188].Y = 0.003

	pointsOfFunctionPlot[189].X = -8.11
	pointsOfFunctionPlot[189].Y = 0.003

	pointsOfFunctionPlot[190].X = -8.10
	pointsOfFunctionPlot[190].Y = 0.003

	pointsOfFunctionPlot[191].X = -8.09
	pointsOfFunctionPlot[191].Y = 0.003

	pointsOfFunctionPlot[192].X = -8.08
	pointsOfFunctionPlot[192].Y = 0.003

	pointsOfFunctionPlot[193].X = -8.07
	pointsOfFunctionPlot[193].Y = 0.003

	pointsOfFunctionPlot[194].X = -8.06
	pointsOfFunctionPlot[194].Y = 0.003

	pointsOfFunctionPlot[195].X = -8.05
	pointsOfFunctionPlot[195].Y = 0.003

	pointsOfFunctionPlot[196].X = -8.04
	pointsOfFunctionPlot[196].Y = 0.003

	pointsOfFunctionPlot[197].X = -8.03
	pointsOfFunctionPlot[197].Y = 0.003

	pointsOfFunctionPlot[198].X = -8.02
	pointsOfFunctionPlot[198].Y = 0.003

	pointsOfFunctionPlot[199].X = -8.01
	pointsOfFunctionPlot[199].Y = 0.003

	pointsOfFunctionPlot[200].X = -8.0
	pointsOfFunctionPlot[200].Y = 0.003

	pointsOfFunctionPlot[201].X = -7.99
	pointsOfFunctionPlot[201].Y = 0.003

	pointsOfFunctionPlot[202].X = -7.98
	pointsOfFunctionPlot[202].Y = 0.004

	pointsOfFunctionPlot[203].X = -7.97
	pointsOfFunctionPlot[203].Y = 0.004

	pointsOfFunctionPlot[204].X = -7.96
	pointsOfFunctionPlot[204].Y = 0.004

	pointsOfFunctionPlot[205].X = -7.95
	pointsOfFunctionPlot[205].Y = 0.004

	pointsOfFunctionPlot[206].X = -7.94
	pointsOfFunctionPlot[206].Y = 0.004

	pointsOfFunctionPlot[207].X = -7.93
	pointsOfFunctionPlot[207].Y = 0.004

	pointsOfFunctionPlot[208].X = -7.92
	pointsOfFunctionPlot[208].Y = 0.004

	pointsOfFunctionPlot[209].X = -7.91
	pointsOfFunctionPlot[209].Y = 0.004

	pointsOfFunctionPlot[210].X = -7.90
	pointsOfFunctionPlot[210].Y = 0.004

	pointsOfFunctionPlot[211].X = -7.89
	pointsOfFunctionPlot[211].Y = 0.004

	pointsOfFunctionPlot[212].X = -7.88
	pointsOfFunctionPlot[212].Y = 0.004

	pointsOfFunctionPlot[213].X = -7.87
	pointsOfFunctionPlot[213].Y = 0.004

	pointsOfFunctionPlot[214].X = -7.86
	pointsOfFunctionPlot[214].Y = 0.004

	pointsOfFunctionPlot[215].X = -7.85
	pointsOfFunctionPlot[215].Y = 0.004

	pointsOfFunctionPlot[216].X = -7.84
	pointsOfFunctionPlot[216].Y = 0.004

	pointsOfFunctionPlot[217].X = -7.83
	pointsOfFunctionPlot[217].Y = 0.004

	pointsOfFunctionPlot[218].X = -7.82
	pointsOfFunctionPlot[218].Y = 0.004

	pointsOfFunctionPlot[219].X = -7.81
	pointsOfFunctionPlot[219].Y = 0.004

	pointsOfFunctionPlot[220].X = -7.80
	pointsOfFunctionPlot[220].Y = 0.004

	pointsOfFunctionPlot[221].X = -7.79
	pointsOfFunctionPlot[221].Y = 0.004

	pointsOfFunctionPlot[222].X = -7.78
	pointsOfFunctionPlot[222].Y = 0.004

	pointsOfFunctionPlot[223].X = -7.77
	pointsOfFunctionPlot[223].Y = 0.004

	pointsOfFunctionPlot[224].X = -7.76
	pointsOfFunctionPlot[224].Y = 0.004

	pointsOfFunctionPlot[225].X = -7.75
	pointsOfFunctionPlot[225].Y = 0.004

	pointsOfFunctionPlot[226].X = -7.74
	pointsOfFunctionPlot[226].Y = 0.004

	pointsOfFunctionPlot[227].X = -7.73
	pointsOfFunctionPlot[227].Y = 0.004

	pointsOfFunctionPlot[228].X = -7.72
	pointsOfFunctionPlot[228].Y = 0.004

	pointsOfFunctionPlot[229].X = -7.71
	pointsOfFunctionPlot[229].Y = 0.004

	pointsOfFunctionPlot[230].X = -7.70
	pointsOfFunctionPlot[230].Y = 0.004

	pointsOfFunctionPlot[231].X = -7.69
	pointsOfFunctionPlot[231].Y = 0.004

	pointsOfFunctionPlot[232].X = -7.68
	pointsOfFunctionPlot[232].Y = 0.004

	pointsOfFunctionPlot[233].X = -7.67
	pointsOfFunctionPlot[233].Y = 0.004

	pointsOfFunctionPlot[234].X = -7.66
	pointsOfFunctionPlot[234].Y = 0.004

	pointsOfFunctionPlot[235].X = -7.65
	pointsOfFunctionPlot[235].Y = 0.005

	pointsOfFunctionPlot[236].X = -7.64
	pointsOfFunctionPlot[236].Y = 0.005

	pointsOfFunctionPlot[237].X = -7.63
	pointsOfFunctionPlot[237].Y = 0.005

	pointsOfFunctionPlot[238].X = -7.62
	pointsOfFunctionPlot[238].Y = 0.005

	pointsOfFunctionPlot[239].X = -7.61
	pointsOfFunctionPlot[239].Y = 0.005

	pointsOfFunctionPlot[240].X = -7.60
	pointsOfFunctionPlot[240].Y = 0.005

	pointsOfFunctionPlot[241].X = -7.59
	pointsOfFunctionPlot[241].Y = 0.005

	pointsOfFunctionPlot[242].X = -7.58
	pointsOfFunctionPlot[242].Y = 0.005

	pointsOfFunctionPlot[243].X = -7.57
	pointsOfFunctionPlot[243].Y = 0.005

	pointsOfFunctionPlot[244].X = -7.56
	pointsOfFunctionPlot[244].Y = 0.005

	pointsOfFunctionPlot[245].X = -7.55
	pointsOfFunctionPlot[245].Y = 0.005

	pointsOfFunctionPlot[246].X = -7.54
	pointsOfFunctionPlot[246].Y = 0.005

	pointsOfFunctionPlot[247].X = -7.53
	pointsOfFunctionPlot[247].Y = 0.005

	pointsOfFunctionPlot[248].X = -7.52
	pointsOfFunctionPlot[248].Y = 0.005

	pointsOfFunctionPlot[249].X = -7.51
	pointsOfFunctionPlot[249].Y = 0.005

	pointsOfFunctionPlot[250].X = -7.50
	pointsOfFunctionPlot[250].Y = 0.005

	pointsOfFunctionPlot[251].X = -7.49
	pointsOfFunctionPlot[251].Y = 0.005

	pointsOfFunctionPlot[252].X = -7.48
	pointsOfFunctionPlot[252].Y = 0.005

	pointsOfFunctionPlot[253].X = -7.47
	pointsOfFunctionPlot[253].Y = 0.005

	pointsOfFunctionPlot[254].X = -7.46
	pointsOfFunctionPlot[254].Y = 0.005

	pointsOfFunctionPlot[255].X = -7.45
	pointsOfFunctionPlot[255].Y = 0.005

	pointsOfFunctionPlot[256].X = -7.44
	pointsOfFunctionPlot[256].Y = 0.005

	pointsOfFunctionPlot[257].X = -7.43
	pointsOfFunctionPlot[257].Y = 0.005

	pointsOfFunctionPlot[258].X = -7.42
	pointsOfFunctionPlot[258].Y = 0.005

	pointsOfFunctionPlot[259].X = -7.41
	pointsOfFunctionPlot[259].Y = 0.005

	pointsOfFunctionPlot[260].X = -7.40
	pointsOfFunctionPlot[260].Y = 0.005

	pointsOfFunctionPlot[261].X = -7.39
	pointsOfFunctionPlot[261].Y = 0.006

	pointsOfFunctionPlot[262].X = -7.38
	pointsOfFunctionPlot[262].Y = 0.006

	pointsOfFunctionPlot[263].X = -7.37
	pointsOfFunctionPlot[263].Y = 0.006

	pointsOfFunctionPlot[264].X = -7.36
	pointsOfFunctionPlot[264].Y = 0.006

	pointsOfFunctionPlot[265].X = -7.35
	pointsOfFunctionPlot[265].Y = 0.006

	pointsOfFunctionPlot[266].X = -7.34
	pointsOfFunctionPlot[266].Y = 0.006

	pointsOfFunctionPlot[267].X = -7.33
	pointsOfFunctionPlot[267].Y = 0.006

	pointsOfFunctionPlot[268].X = -7.32
	pointsOfFunctionPlot[268].Y = 0.006

	pointsOfFunctionPlot[269].X = -7.31
	pointsOfFunctionPlot[269].Y = 0.006

	pointsOfFunctionPlot[270].X = -7.30
	pointsOfFunctionPlot[270].Y = 0.006

	pointsOfFunctionPlot[271].X = -7.29
	pointsOfFunctionPlot[271].Y = 0.006

	pointsOfFunctionPlot[272].X = -7.28
	pointsOfFunctionPlot[272].Y = 0.006

	pointsOfFunctionPlot[273].X = -7.27
	pointsOfFunctionPlot[273].Y = 0.006

	pointsOfFunctionPlot[274].X = -7.26
	pointsOfFunctionPlot[274].Y = 0.006

	pointsOfFunctionPlot[275].X = -7.25
	pointsOfFunctionPlot[275].Y = 0.006

	pointsOfFunctionPlot[276].X = -7.24
	pointsOfFunctionPlot[276].Y = 0.006

	pointsOfFunctionPlot[277].X = -7.23
	pointsOfFunctionPlot[277].Y = 0.006

	pointsOfFunctionPlot[278].X = -7.22
	pointsOfFunctionPlot[278].Y = 0.006

	pointsOfFunctionPlot[279].X = -7.21
	pointsOfFunctionPlot[279].Y = 0.006

	pointsOfFunctionPlot[280].X = -7.20
	pointsOfFunctionPlot[280].Y = 0.006

	pointsOfFunctionPlot[281].X = -7.19
	pointsOfFunctionPlot[281].Y = 0.006

	pointsOfFunctionPlot[282].X = -7.18
	pointsOfFunctionPlot[282].Y = 0.006

	pointsOfFunctionPlot[283].X = -7.17
	pointsOfFunctionPlot[283].Y = 0.006

	pointsOfFunctionPlot[284].X = -7.16
	pointsOfFunctionPlot[284].Y = 0.007

	pointsOfFunctionPlot[285].X = -7.15
	pointsOfFunctionPlot[285].Y = 0.007

	pointsOfFunctionPlot[286].X = -7.14
	pointsOfFunctionPlot[286].Y = 0.007

	pointsOfFunctionPlot[287].X = -7.13
	pointsOfFunctionPlot[287].Y = 0.007

	pointsOfFunctionPlot[288].X = -7.12
	pointsOfFunctionPlot[288].Y = 0.007

	pointsOfFunctionPlot[289].X = -7.11
	pointsOfFunctionPlot[289].Y = 0.007

	pointsOfFunctionPlot[290].X = -7.10
	pointsOfFunctionPlot[290].Y = 0.007

	pointsOfFunctionPlot[291].X = -7.09
	pointsOfFunctionPlot[291].Y = 0.007

	pointsOfFunctionPlot[292].X = -7.08
	pointsOfFunctionPlot[292].Y = 0.007

	pointsOfFunctionPlot[293].X = -7.07
	pointsOfFunctionPlot[293].Y = 0.007

	pointsOfFunctionPlot[294].X = -7.06
	pointsOfFunctionPlot[294].Y = 0.007

	pointsOfFunctionPlot[295].X = -7.05
	pointsOfFunctionPlot[295].Y = 0.007

	pointsOfFunctionPlot[296].X = -7.04
	pointsOfFunctionPlot[296].Y = 0.007

	pointsOfFunctionPlot[297].X = -7.03
	pointsOfFunctionPlot[297].Y = 0.007

	pointsOfFunctionPlot[298].X = -7.02
	pointsOfFunctionPlot[298].Y = 0.007

	pointsOfFunctionPlot[299].X = -7.01
	pointsOfFunctionPlot[299].Y = 0.007

	pointsOfFunctionPlot[300].X = -7.0
	pointsOfFunctionPlot[300].Y = 0.007

	pointsOfFunctionPlot[301].X = -6.99
	pointsOfFunctionPlot[301].Y = 0.007

	pointsOfFunctionPlot[302].X = -6.98
	pointsOfFunctionPlot[302].Y = 0.007

	pointsOfFunctionPlot[303].X = -6.97
	pointsOfFunctionPlot[303].Y = 0.008

	pointsOfFunctionPlot[304].X = -6.96
	pointsOfFunctionPlot[304].Y = 0.008

	pointsOfFunctionPlot[305].X = -6.95
	pointsOfFunctionPlot[305].Y = 0.008

	pointsOfFunctionPlot[306].X = -6.94
	pointsOfFunctionPlot[306].Y = 0.008

	pointsOfFunctionPlot[307].X = -6.93
	pointsOfFunctionPlot[307].Y = 0.008

	pointsOfFunctionPlot[308].X = -6.92
	pointsOfFunctionPlot[308].Y = 0.008

	pointsOfFunctionPlot[309].X = -6.91
	pointsOfFunctionPlot[309].Y = 0.008

	pointsOfFunctionPlot[310].X = -6.90
	pointsOfFunctionPlot[310].Y = 0.008

	pointsOfFunctionPlot[311].X = -6.89
	pointsOfFunctionPlot[311].Y = 0.008

	pointsOfFunctionPlot[312].X = -6.88
	pointsOfFunctionPlot[312].Y = 0.008

	pointsOfFunctionPlot[313].X = -6.87
	pointsOfFunctionPlot[313].Y = 0.008

	pointsOfFunctionPlot[314].X = -6.86
	pointsOfFunctionPlot[314].Y = 0.008

	pointsOfFunctionPlot[315].X = -6.85
	pointsOfFunctionPlot[315].Y = 0.008

	pointsOfFunctionPlot[316].X = -6.84
	pointsOfFunctionPlot[316].Y = 0.008

	pointsOfFunctionPlot[317].X = -6.83
	pointsOfFunctionPlot[317].Y = 0.008

	pointsOfFunctionPlot[318].X = -6.82
	pointsOfFunctionPlot[318].Y = 0.008

	pointsOfFunctionPlot[319].X = -6.81
	pointsOfFunctionPlot[319].Y = 0.008

	pointsOfFunctionPlot[320].X = -6.80
	pointsOfFunctionPlot[320].Y = 0.009

	pointsOfFunctionPlot[321].X = -6.79
	pointsOfFunctionPlot[321].Y = 0.009

	pointsOfFunctionPlot[322].X = -6.78
	pointsOfFunctionPlot[322].Y = 0.009

	pointsOfFunctionPlot[323].X = -6.77
	pointsOfFunctionPlot[323].Y = 0.009

	pointsOfFunctionPlot[324].X = -6.76
	pointsOfFunctionPlot[324].Y = 0.009

	pointsOfFunctionPlot[325].X = -6.75
	pointsOfFunctionPlot[325].Y = 0.009

	pointsOfFunctionPlot[326].X = -6.74
	pointsOfFunctionPlot[326].Y = 0.009

	pointsOfFunctionPlot[327].X = -6.73
	pointsOfFunctionPlot[327].Y = 0.009

	pointsOfFunctionPlot[328].X = -6.72
	pointsOfFunctionPlot[328].Y = 0.009

	pointsOfFunctionPlot[329].X = -6.71
	pointsOfFunctionPlot[329].Y = 0.009

	pointsOfFunctionPlot[330].X = -6.70
	pointsOfFunctionPlot[330].Y = 0.009

	pointsOfFunctionPlot[331].X = -6.69
	pointsOfFunctionPlot[331].Y = 0.009

	pointsOfFunctionPlot[332].X = -6.68
	pointsOfFunctionPlot[332].Y = 0.009

	pointsOfFunctionPlot[333].X = -6.67
	pointsOfFunctionPlot[333].Y = 0.009

	pointsOfFunctionPlot[334].X = -6.66
	pointsOfFunctionPlot[334].Y = 0.009

	pointsOfFunctionPlot[335].X = -6.65
	pointsOfFunctionPlot[335].Y = 0.01

	pointsOfFunctionPlot[336].X = -6.64
	pointsOfFunctionPlot[336].Y = 0.01

	pointsOfFunctionPlot[337].X = -6.63
	pointsOfFunctionPlot[337].Y = 0.01

	pointsOfFunctionPlot[338].X = -6.62
	pointsOfFunctionPlot[338].Y = 0.01

	pointsOfFunctionPlot[339].X = -6.61
	pointsOfFunctionPlot[339].Y = 0.01

	pointsOfFunctionPlot[340].X = -6.60
	pointsOfFunctionPlot[340].Y = 0.01

	pointsOfFunctionPlot[341].X = -6.59
	pointsOfFunctionPlot[341].Y = 0.01

	pointsOfFunctionPlot[342].X = -6.58
	pointsOfFunctionPlot[342].Y = 0.01

	pointsOfFunctionPlot[343].X = -6.57
	pointsOfFunctionPlot[343].Y = 0.01

	pointsOfFunctionPlot[344].X = -6.56
	pointsOfFunctionPlot[344].Y = 0.01

	pointsOfFunctionPlot[345].X = -6.55
	pointsOfFunctionPlot[345].Y = 0.01

	pointsOfFunctionPlot[346].X = -6.54
	pointsOfFunctionPlot[346].Y = 0.01

	pointsOfFunctionPlot[347].X = -6.53
	pointsOfFunctionPlot[347].Y = 0.01

	pointsOfFunctionPlot[348].X = -6.52
	pointsOfFunctionPlot[348].Y = 0.01

	pointsOfFunctionPlot[349].X = -6.51
	pointsOfFunctionPlot[349].Y = 0.011

	pointsOfFunctionPlot[350].X = -6.50
	pointsOfFunctionPlot[350].Y = 0.011

	pointsOfFunctionPlot[351].X = -6.49
	pointsOfFunctionPlot[351].Y = 0.011

	pointsOfFunctionPlot[352].X = -6.48
	pointsOfFunctionPlot[352].Y = 0.011

	pointsOfFunctionPlot[353].X = -6.47
	pointsOfFunctionPlot[353].Y = 0.011

	pointsOfFunctionPlot[354].X = -6.46
	pointsOfFunctionPlot[354].Y = 0.011

	pointsOfFunctionPlot[355].X = -6.45
	pointsOfFunctionPlot[355].Y = 0.011

	pointsOfFunctionPlot[356].X = -6.44
	pointsOfFunctionPlot[356].Y = 0.011

	pointsOfFunctionPlot[357].X = -6.43
	pointsOfFunctionPlot[357].Y = 0.011

	pointsOfFunctionPlot[358].X = -6.42
	pointsOfFunctionPlot[358].Y = 0.011

	pointsOfFunctionPlot[359].X = -6.41
	pointsOfFunctionPlot[359].Y = 0.011

	pointsOfFunctionPlot[360].X = -6.40
	pointsOfFunctionPlot[360].Y = 0.011

	pointsOfFunctionPlot[361].X = -6.39
	pointsOfFunctionPlot[361].Y = 0.011

	pointsOfFunctionPlot[362].X = -6.38
	pointsOfFunctionPlot[362].Y = 0.012

	pointsOfFunctionPlot[363].X = -6.37
	pointsOfFunctionPlot[363].Y = 0.012

	pointsOfFunctionPlot[364].X = -6.36
	pointsOfFunctionPlot[364].Y = 0.012

	pointsOfFunctionPlot[365].X = -6.35
	pointsOfFunctionPlot[365].Y = 0.012

	pointsOfFunctionPlot[366].X = -6.34
	pointsOfFunctionPlot[366].Y = 0.012

	pointsOfFunctionPlot[367].X = -6.33
	pointsOfFunctionPlot[367].Y = 0.012

	pointsOfFunctionPlot[368].X = -6.32
	pointsOfFunctionPlot[368].Y = 0.012

	pointsOfFunctionPlot[369].X = -6.31
	pointsOfFunctionPlot[369].Y = 0.012

	pointsOfFunctionPlot[370].X = -6.30
	pointsOfFunctionPlot[370].Y = 0.012

	pointsOfFunctionPlot[371].X = -6.29
	pointsOfFunctionPlot[371].Y = 0.012

	pointsOfFunctionPlot[372].X = -6.28
	pointsOfFunctionPlot[372].Y = 0.012

	pointsOfFunctionPlot[373].X = -6.27
	pointsOfFunctionPlot[373].Y = 0.013

	pointsOfFunctionPlot[374].X = -6.26
	pointsOfFunctionPlot[374].Y = 0.013

	pointsOfFunctionPlot[375].X = -6.25
	pointsOfFunctionPlot[375].Y = 0.013

	pointsOfFunctionPlot[376].X = -6.24
	pointsOfFunctionPlot[376].Y = 0.013

	pointsOfFunctionPlot[377].X = -6.23
	pointsOfFunctionPlot[377].Y = 0.013

	pointsOfFunctionPlot[378].X = -6.22
	pointsOfFunctionPlot[378].Y = 0.013

	pointsOfFunctionPlot[379].X = -6.21
	pointsOfFunctionPlot[379].Y = 0.013

	pointsOfFunctionPlot[380].X = -6.20
	pointsOfFunctionPlot[380].Y = 0.013

	pointsOfFunctionPlot[381].X = -6.19
	pointsOfFunctionPlot[381].Y = 0.013

	pointsOfFunctionPlot[382].X = -6.18
	pointsOfFunctionPlot[382].Y = 0.013

	pointsOfFunctionPlot[383].X = -6.17
	pointsOfFunctionPlot[383].Y = 0.013

	pointsOfFunctionPlot[384].X = -6.16
	pointsOfFunctionPlot[384].Y = 0.014

	pointsOfFunctionPlot[385].X = -6.15
	pointsOfFunctionPlot[385].Y = 0.014

	pointsOfFunctionPlot[386].X = -6.14
	pointsOfFunctionPlot[386].Y = 0.014

	pointsOfFunctionPlot[387].X = -6.13
	pointsOfFunctionPlot[387].Y = 0.014

	pointsOfFunctionPlot[388].X = -6.12
	pointsOfFunctionPlot[388].Y = 0.014

	pointsOfFunctionPlot[389].X = -6.11
	pointsOfFunctionPlot[389].Y = 0.014

	pointsOfFunctionPlot[390].X = -6.10
	pointsOfFunctionPlot[390].Y = 0.014

	pointsOfFunctionPlot[391].X = -6.09
	pointsOfFunctionPlot[391].Y = 0.014

	pointsOfFunctionPlot[392].X = -6.08
	pointsOfFunctionPlot[392].Y = 0.014

	pointsOfFunctionPlot[393].X = -6.07
	pointsOfFunctionPlot[393].Y = 0.014

	pointsOfFunctionPlot[394].X = -6.06
	pointsOfFunctionPlot[394].Y = 0.015

	pointsOfFunctionPlot[395].X = -6.05
	pointsOfFunctionPlot[395].Y = 0.015

	pointsOfFunctionPlot[396].X = -6.04
	pointsOfFunctionPlot[396].Y = 0.015

	pointsOfFunctionPlot[397].X = -6.03
	pointsOfFunctionPlot[397].Y = 0.015

	pointsOfFunctionPlot[398].X = -6.02
	pointsOfFunctionPlot[398].Y = 0.015

	pointsOfFunctionPlot[399].X = -6.01
	pointsOfFunctionPlot[399].Y = 0.015

	pointsOfFunctionPlot[400].X = -6.0
	pointsOfFunctionPlot[400].Y = 0.015

	pointsOfFunctionPlot[401].X = -5.99
	pointsOfFunctionPlot[401].Y = 0.015

	pointsOfFunctionPlot[402].X = -5.98
	pointsOfFunctionPlot[402].Y = 0.015

	pointsOfFunctionPlot[403].X = -5.97
	pointsOfFunctionPlot[403].Y = 0.016

	pointsOfFunctionPlot[404].X = -5.96
	pointsOfFunctionPlot[404].Y = 0.016

	pointsOfFunctionPlot[405].X = -5.95
	pointsOfFunctionPlot[405].Y = 0.016

	pointsOfFunctionPlot[406].X = -5.94
	pointsOfFunctionPlot[406].Y = 0.016

	pointsOfFunctionPlot[407].X = -5.93
	pointsOfFunctionPlot[407].Y = 0.016

	pointsOfFunctionPlot[408].X = -5.92
	pointsOfFunctionPlot[408].Y = 0.016

	pointsOfFunctionPlot[409].X = -5.91
	pointsOfFunctionPlot[409].Y = 0.016

	pointsOfFunctionPlot[410].X = -5.90
	pointsOfFunctionPlot[410].Y = 0.016

	pointsOfFunctionPlot[411].X = -5.89
	pointsOfFunctionPlot[411].Y = 0.016

	pointsOfFunctionPlot[412].X = -5.88
	pointsOfFunctionPlot[412].Y = 0.017

	pointsOfFunctionPlot[413].X = -5.87
	pointsOfFunctionPlot[413].Y = 0.017

	pointsOfFunctionPlot[414].X = -5.86
	pointsOfFunctionPlot[414].Y = 0.017

	pointsOfFunctionPlot[415].X = -5.85
	pointsOfFunctionPlot[415].Y = 0.017

	pointsOfFunctionPlot[416].X = -5.84
	pointsOfFunctionPlot[416].Y = 0.017

	pointsOfFunctionPlot[417].X = -5.82
	pointsOfFunctionPlot[417].Y = 0.017

	pointsOfFunctionPlot[418].X = -5.82
	pointsOfFunctionPlot[418].Y = 0.017

	pointsOfFunctionPlot[419].X = -5.81
	pointsOfFunctionPlot[419].Y = 0.017

	pointsOfFunctionPlot[420].X = -5.80
	pointsOfFunctionPlot[420].Y = 0.017

	pointsOfFunctionPlot[421].X = -5.79
	pointsOfFunctionPlot[421].Y = 0.018

	pointsOfFunctionPlot[422].X = -5.78
	pointsOfFunctionPlot[422].Y = 0.018

	pointsOfFunctionPlot[423].X = -5.77
	pointsOfFunctionPlot[423].Y = 0.018

	pointsOfFunctionPlot[424].X = -5.76
	pointsOfFunctionPlot[424].Y = 0.018

	pointsOfFunctionPlot[425].X = -5.75
	pointsOfFunctionPlot[425].Y = 0.018

	pointsOfFunctionPlot[426].X = -5.74
	pointsOfFunctionPlot[426].Y = 0.018

	pointsOfFunctionPlot[427].X = -5.73
	pointsOfFunctionPlot[427].Y = 0.018

	pointsOfFunctionPlot[428].X = -5.72
	pointsOfFunctionPlot[428].Y = 0.019

	pointsOfFunctionPlot[429].X = -5.71
	pointsOfFunctionPlot[429].Y = 0.019

	pointsOfFunctionPlot[430].X = -5.70
	pointsOfFunctionPlot[430].Y = 0.019

	pointsOfFunctionPlot[431].X = -5.69
	pointsOfFunctionPlot[431].Y = 0.019

	pointsOfFunctionPlot[432].X = -5.68
	pointsOfFunctionPlot[432].Y = 0.019

	pointsOfFunctionPlot[433].X = -5.67
	pointsOfFunctionPlot[433].Y = 0.019

	pointsOfFunctionPlot[434].X = -5.66
	pointsOfFunctionPlot[434].Y = 0.019

	pointsOfFunctionPlot[435].X = -5.65
	pointsOfFunctionPlot[435].Y = 0.019

	pointsOfFunctionPlot[436].X = -5.64
	pointsOfFunctionPlot[436].Y = 0.02

	pointsOfFunctionPlot[437].X = -5.63
	pointsOfFunctionPlot[437].Y = 0.02

	pointsOfFunctionPlot[438].X = -5.62
	pointsOfFunctionPlot[438].Y = 0.02

	pointsOfFunctionPlot[439].X = -5.61
	pointsOfFunctionPlot[439].Y = 0.02

	pointsOfFunctionPlot[440].X = -5.60
	pointsOfFunctionPlot[440].Y = 0.02

	pointsOfFunctionPlot[441].X = -5.59
	pointsOfFunctionPlot[441].Y = 0.02

	pointsOfFunctionPlot[442].X = -5.58
	pointsOfFunctionPlot[442].Y = 0.02

	pointsOfFunctionPlot[443].X = -5.57
	pointsOfFunctionPlot[443].Y = 0.021

	pointsOfFunctionPlot[444].X = -5.56
	pointsOfFunctionPlot[444].Y = 0.021

	pointsOfFunctionPlot[445].X = -5.55
	pointsOfFunctionPlot[445].Y = 0.021

	pointsOfFunctionPlot[446].X = -5.54
	pointsOfFunctionPlot[446].Y = 0.021

	pointsOfFunctionPlot[447].X = -5.53
	pointsOfFunctionPlot[447].Y = 0.021

	pointsOfFunctionPlot[448].X = -5.52
	pointsOfFunctionPlot[448].Y = 0.021

	pointsOfFunctionPlot[449].X = -5.51
	pointsOfFunctionPlot[449].Y = 0.021

	pointsOfFunctionPlot[450].X = -5.50
	pointsOfFunctionPlot[450].Y = 0.022

	pointsOfFunctionPlot[451].X = -5.49
	pointsOfFunctionPlot[451].Y = 0.022

	pointsOfFunctionPlot[452].X = -5.48
	pointsOfFunctionPlot[452].Y = 0.022

	pointsOfFunctionPlot[453].X = -5.47
	pointsOfFunctionPlot[453].Y = 0.022

	pointsOfFunctionPlot[454].X = -5.46
	pointsOfFunctionPlot[454].Y = 0.022

	pointsOfFunctionPlot[455].X = -5.45
	pointsOfFunctionPlot[455].Y = 0.022

	pointsOfFunctionPlot[456].X = -5.44
	pointsOfFunctionPlot[456].Y = 0.023

	pointsOfFunctionPlot[457].X = -5.43
	pointsOfFunctionPlot[457].Y = 0.023

	pointsOfFunctionPlot[458].X = -5.42
	pointsOfFunctionPlot[458].Y = 0.023

	pointsOfFunctionPlot[459].X = -5.41
	pointsOfFunctionPlot[459].Y = 0.023

	pointsOfFunctionPlot[460].X = -5.40
	pointsOfFunctionPlot[460].Y = 0.023

	pointsOfFunctionPlot[461].X = -5.39
	pointsOfFunctionPlot[461].Y = 0.023

	pointsOfFunctionPlot[462].X = -5.38
	pointsOfFunctionPlot[462].Y = 0.024

	pointsOfFunctionPlot[463].X = -5.37
	pointsOfFunctionPlot[463].Y = 0.024

	pointsOfFunctionPlot[464].X = -5.36
	pointsOfFunctionPlot[464].Y = 0.024

	pointsOfFunctionPlot[465].X = -5.35
	pointsOfFunctionPlot[465].Y = 0.024

	pointsOfFunctionPlot[466].X = -5.34
	pointsOfFunctionPlot[466].Y = 0.024

	pointsOfFunctionPlot[467].X = -5.33
	pointsOfFunctionPlot[467].Y = 0.024

	pointsOfFunctionPlot[468].X = -5.32
	pointsOfFunctionPlot[468].Y = 0.025

	pointsOfFunctionPlot[469].X = -5.31
	pointsOfFunctionPlot[469].Y = 0.025

	pointsOfFunctionPlot[470].X = -5.30
	pointsOfFunctionPlot[470].Y = 0.025

	pointsOfFunctionPlot[471].X = -5.29
	pointsOfFunctionPlot[471].Y = 0.025

	pointsOfFunctionPlot[472].X = -5.28
	pointsOfFunctionPlot[472].Y = 0.025

	pointsOfFunctionPlot[473].X = -5.27
	pointsOfFunctionPlot[473].Y = 0.025

	pointsOfFunctionPlot[474].X = -5.26
	pointsOfFunctionPlot[474].Y = 0.026

	pointsOfFunctionPlot[475].X = -5.25
	pointsOfFunctionPlot[475].Y = 0.026

	pointsOfFunctionPlot[476].X = -5.24
	pointsOfFunctionPlot[476].Y = 0.026

	pointsOfFunctionPlot[477].X = -5.23
	pointsOfFunctionPlot[477].Y = 0.026

	pointsOfFunctionPlot[478].X = -5.22
	pointsOfFunctionPlot[478].Y = 0.026

	pointsOfFunctionPlot[479].X = -5.21
	pointsOfFunctionPlot[479].Y = 0.027

	pointsOfFunctionPlot[480].X = -5.20
	pointsOfFunctionPlot[480].Y = 0.027

	pointsOfFunctionPlot[481].X = -5.19
	pointsOfFunctionPlot[481].Y = 0.027

	pointsOfFunctionPlot[482].X = -5.18
	pointsOfFunctionPlot[482].Y = 0.027

	pointsOfFunctionPlot[483].X = -5.17
	pointsOfFunctionPlot[483].Y = 0.027

	pointsOfFunctionPlot[484].X = -5.16
	pointsOfFunctionPlot[484].Y = 0.028

	pointsOfFunctionPlot[485].X = -5.15
	pointsOfFunctionPlot[485].Y = 0.028

	pointsOfFunctionPlot[486].X = -5.14
	pointsOfFunctionPlot[486].Y = 0.028

	pointsOfFunctionPlot[487].X = -5.13
	pointsOfFunctionPlot[487].Y = 0.028

	pointsOfFunctionPlot[488].X = -5.12
	pointsOfFunctionPlot[488].Y = 0.028

	pointsOfFunctionPlot[489].X = -5.11
	pointsOfFunctionPlot[489].Y = 0.029

	pointsOfFunctionPlot[490].X = -5.10
	pointsOfFunctionPlot[490].Y = 0.029

	pointsOfFunctionPlot[491].X = -5.09
	pointsOfFunctionPlot[491].Y = 0.029

	pointsOfFunctionPlot[492].X = -5.08
	pointsOfFunctionPlot[492].Y = 0.029

	pointsOfFunctionPlot[493].X = -5.07
	pointsOfFunctionPlot[493].Y = 0.029

	pointsOfFunctionPlot[494].X = -5.06
	pointsOfFunctionPlot[494].Y = 0.03

	pointsOfFunctionPlot[495].X = -5.05
	pointsOfFunctionPlot[495].Y = 0.03

	pointsOfFunctionPlot[496].X = -5.04
	pointsOfFunctionPlot[496].Y = 0.03

	pointsOfFunctionPlot[497].X = -5.03
	pointsOfFunctionPlot[497].Y = 0.03

	pointsOfFunctionPlot[498].X = -5.02
	pointsOfFunctionPlot[498].Y = 0.031

	pointsOfFunctionPlot[499].X = -5.01
	pointsOfFunctionPlot[499].Y = 0.031

	pointsOfFunctionPlot[500].X = -5.0
	pointsOfFunctionPlot[500].Y = 0.031

	pointsOfFunctionPlot[501].X = -4.99
	pointsOfFunctionPlot[501].Y = 0.031

	pointsOfFunctionPlot[502].X = -4.98
	pointsOfFunctionPlot[502].Y = 0.031

	pointsOfFunctionPlot[503].X = -4.97
	pointsOfFunctionPlot[503].Y = 0.031

	pointsOfFunctionPlot[504].X = -4.96
	pointsOfFunctionPlot[504].Y = 0.032

	pointsOfFunctionPlot[505].X = -4.95
	pointsOfFunctionPlot[505].Y = 0.032

	pointsOfFunctionPlot[506].X = -4.94
	pointsOfFunctionPlot[506].Y = 0.032

	pointsOfFunctionPlot[507].X = -4.93
	pointsOfFunctionPlot[507].Y = 0.032

	pointsOfFunctionPlot[508].X = -4.92
	pointsOfFunctionPlot[508].Y = 0.033

	pointsOfFunctionPlot[509].X = -4.91
	pointsOfFunctionPlot[509].Y = 0.033

	pointsOfFunctionPlot[510].X = -4.90
	pointsOfFunctionPlot[510].Y = 0.033

	pointsOfFunctionPlot[511].X = -4.89
	pointsOfFunctionPlot[511].Y = 0.033

	pointsOfFunctionPlot[512].X = -4.88
	pointsOfFunctionPlot[512].Y = 0.034

	pointsOfFunctionPlot[513].X = -4.87
	pointsOfFunctionPlot[513].Y = 0.034

	pointsOfFunctionPlot[514].X = -4.86
	pointsOfFunctionPlot[514].Y = 0.034

	pointsOfFunctionPlot[515].X = -4.85
	pointsOfFunctionPlot[515].Y = 0.034

	pointsOfFunctionPlot[516].X = -4.84
	pointsOfFunctionPlot[516].Y = 0.034

	pointsOfFunctionPlot[517].X = -4.83
	pointsOfFunctionPlot[517].Y = 0.035

	pointsOfFunctionPlot[518].X = -4.82
	pointsOfFunctionPlot[518].Y = 0.035

	pointsOfFunctionPlot[519].X = -4.81
	pointsOfFunctionPlot[519].Y = 0.035

	pointsOfFunctionPlot[520].X = -4.80
	pointsOfFunctionPlot[520].Y = 0.035

	pointsOfFunctionPlot[521].X = -4.79
	pointsOfFunctionPlot[521].Y = 0.036

	pointsOfFunctionPlot[522].X = -4.78
	pointsOfFunctionPlot[522].Y = 0.036

	pointsOfFunctionPlot[523].X = -4.77
	pointsOfFunctionPlot[523].Y = 0.036

	pointsOfFunctionPlot[524].X = -4.76
	pointsOfFunctionPlot[524].Y = 0.036

	pointsOfFunctionPlot[525].X = -4.75
	pointsOfFunctionPlot[525].Y = 0.037

	pointsOfFunctionPlot[526].X = -4.74
	pointsOfFunctionPlot[526].Y = 0.037

	pointsOfFunctionPlot[527].X = -4.73
	pointsOfFunctionPlot[527].Y = 0.037

	pointsOfFunctionPlot[528].X = -4.72
	pointsOfFunctionPlot[528].Y = 0.037

	pointsOfFunctionPlot[529].X = -4.71
	pointsOfFunctionPlot[529].Y = 0.038

	pointsOfFunctionPlot[530].X = -4.70
	pointsOfFunctionPlot[530].Y = 0.038

	pointsOfFunctionPlot[531].X = -4.69
	pointsOfFunctionPlot[531].Y = 0.038

	pointsOfFunctionPlot[532].X = -4.68
	pointsOfFunctionPlot[532].Y = 0.039

	pointsOfFunctionPlot[533].X = -4.67
	pointsOfFunctionPlot[533].Y = 0.039

	pointsOfFunctionPlot[534].X = -4.66
	pointsOfFunctionPlot[534].Y = 0.039

	pointsOfFunctionPlot[535].X = -4.65
	pointsOfFunctionPlot[535].Y = 0.039

	pointsOfFunctionPlot[536].X = -4.64
	pointsOfFunctionPlot[536].Y = 0.04

	pointsOfFunctionPlot[537].X = -4.63
	pointsOfFunctionPlot[537].Y = 0.04

	pointsOfFunctionPlot[538].X = -4.62
	pointsOfFunctionPlot[538].Y = 0.04

	pointsOfFunctionPlot[539].X = -4.61
	pointsOfFunctionPlot[539].Y = 0.04

	pointsOfFunctionPlot[540].X = -4.60
	pointsOfFunctionPlot[540].Y = 0.041

	pointsOfFunctionPlot[541].X = -4.59
	pointsOfFunctionPlot[541].Y = 0.041

	pointsOfFunctionPlot[542].X = -4.58
	pointsOfFunctionPlot[542].Y = 0.041

	pointsOfFunctionPlot[543].X = -4.57
	pointsOfFunctionPlot[543].Y = 0.041

	pointsOfFunctionPlot[544].X = -4.56
	pointsOfFunctionPlot[544].Y = 0.042

	pointsOfFunctionPlot[545].X = -4.55
	pointsOfFunctionPlot[545].Y = 0.042

	pointsOfFunctionPlot[546].X = -4.54
	pointsOfFunctionPlot[546].Y = 0.043

	pointsOfFunctionPlot[547].X = -4.53
	pointsOfFunctionPlot[547].Y = 0.043

	pointsOfFunctionPlot[548].X = -4.52
	pointsOfFunctionPlot[548].Y = 0.043

	pointsOfFunctionPlot[549].X = -4.51
	pointsOfFunctionPlot[549].Y = 0.043

	pointsOfFunctionPlot[550].X = -4.50
	pointsOfFunctionPlot[550].Y = 0.044

	pointsOfFunctionPlot[551].X = -4.49
	pointsOfFunctionPlot[551].Y = 0.044

	pointsOfFunctionPlot[552].X = -4.48
	pointsOfFunctionPlot[552].Y = 0.044

	pointsOfFunctionPlot[553].X = -4.47
	pointsOfFunctionPlot[553].Y = 0.045

	pointsOfFunctionPlot[554].X = -4.46
	pointsOfFunctionPlot[554].Y = 0.045

	pointsOfFunctionPlot[555].X = -4.45
	pointsOfFunctionPlot[555].Y = 0.045

	pointsOfFunctionPlot[556].X = -4.44
	pointsOfFunctionPlot[556].Y = 0.046

	pointsOfFunctionPlot[557].X = -4.43
	pointsOfFunctionPlot[557].Y = 0.046

	pointsOfFunctionPlot[558].X = -4.42
	pointsOfFunctionPlot[558].Y = 0.046

	pointsOfFunctionPlot[559].X = -4.41
	pointsOfFunctionPlot[559].Y = 0.047

	pointsOfFunctionPlot[560].X = -4.40
	pointsOfFunctionPlot[560].Y = 0.047

	pointsOfFunctionPlot[561].X = -4.39
	pointsOfFunctionPlot[561].Y = 0.047

	pointsOfFunctionPlot[562].X = -4.38
	pointsOfFunctionPlot[562].Y = 0.048

	pointsOfFunctionPlot[563].X = -4.37
	pointsOfFunctionPlot[563].Y = 0.048

	pointsOfFunctionPlot[564].X = -4.36
	pointsOfFunctionPlot[564].Y = 0.048

	pointsOfFunctionPlot[565].X = -4.35
	pointsOfFunctionPlot[565].Y = 0.049

	pointsOfFunctionPlot[566].X = -4.34
	pointsOfFunctionPlot[566].Y = 0.049

	pointsOfFunctionPlot[567].X = -4.33
	pointsOfFunctionPlot[567].Y = 0.049

	pointsOfFunctionPlot[568].X = -4.32
	pointsOfFunctionPlot[568].Y = 0.05

	pointsOfFunctionPlot[569].X = -4.31
	pointsOfFunctionPlot[569].Y = 0.05

	pointsOfFunctionPlot[570].X = -4.30
	pointsOfFunctionPlot[570].Y = 0.05

	pointsOfFunctionPlot[571].X = -4.29
	pointsOfFunctionPlot[571].Y = 0.051

	pointsOfFunctionPlot[572].X = -4.28
	pointsOfFunctionPlot[572].Y = 0.051

	pointsOfFunctionPlot[573].X = -4.27
	pointsOfFunctionPlot[573].Y = 0.051

	pointsOfFunctionPlot[574].X = -4.26
	pointsOfFunctionPlot[574].Y = 0.052

	pointsOfFunctionPlot[575].X = -4.25
	pointsOfFunctionPlot[575].Y = 0.052

	pointsOfFunctionPlot[576].X = -4.24
	pointsOfFunctionPlot[576].Y = 0.052

	pointsOfFunctionPlot[577].X = -4.23
	pointsOfFunctionPlot[577].Y = 0.053

	pointsOfFunctionPlot[578].X = -4.22
	pointsOfFunctionPlot[578].Y = 0.053

	pointsOfFunctionPlot[579].X = -4.21
	pointsOfFunctionPlot[579].Y = 0.054

	pointsOfFunctionPlot[580].X = -4.20
	pointsOfFunctionPlot[580].Y = 0.045

	pointsOfFunctionPlot[581].X = -4.19
	pointsOfFunctionPlot[581].Y = 0.054

	pointsOfFunctionPlot[582].X = -4.18
	pointsOfFunctionPlot[582].Y = 0.055

	pointsOfFunctionPlot[583].X = -4.17
	pointsOfFunctionPlot[583].Y = 0.055

	pointsOfFunctionPlot[584].X = -4.16
	pointsOfFunctionPlot[584].Y = 0.055

	pointsOfFunctionPlot[585].X = -4.15
	pointsOfFunctionPlot[585].Y = 0.056

	pointsOfFunctionPlot[586].X = -4.14
	pointsOfFunctionPlot[586].Y = 0.056

	pointsOfFunctionPlot[587].X = -4.13
	pointsOfFunctionPlot[587].Y = 0.057

	pointsOfFunctionPlot[588].X = -4.12
	pointsOfFunctionPlot[588].Y = 0.057

	pointsOfFunctionPlot[589].X = -4.11
	pointsOfFunctionPlot[589].Y = 0.057

	pointsOfFunctionPlot[590].X = -4.10
	pointsOfFunctionPlot[590].Y = 0.058

	pointsOfFunctionPlot[591].X = -4.09
	pointsOfFunctionPlot[591].Y = 0.058

	pointsOfFunctionPlot[592].X = -4.08
	pointsOfFunctionPlot[592].Y = 0.059

	pointsOfFunctionPlot[593].X = -4.07
	pointsOfFunctionPlot[593].Y = 0.059

	pointsOfFunctionPlot[594].X = -4.06
	pointsOfFunctionPlot[594].Y = 0.06

	pointsOfFunctionPlot[595].X = -4.05
	pointsOfFunctionPlot[595].Y = 0.06

	pointsOfFunctionPlot[596].X = -4.04
	pointsOfFunctionPlot[596].Y = 0.06

	pointsOfFunctionPlot[597].X = -4.03
	pointsOfFunctionPlot[597].Y = 0.061

	pointsOfFunctionPlot[598].X = -4.02
	pointsOfFunctionPlot[598].Y = 0.061

	pointsOfFunctionPlot[599].X = -4.01
	pointsOfFunctionPlot[599].Y = 0.062

	pointsOfFunctionPlot[600].X = -4.0
	pointsOfFunctionPlot[600].Y = 0.062

	pointsOfFunctionPlot[601].X = -3.99
	pointsOfFunctionPlot[601].Y = 0.062

	pointsOfFunctionPlot[602].X = -3.98
	pointsOfFunctionPlot[602].Y = 0.063

	pointsOfFunctionPlot[603].X = -3.97
	pointsOfFunctionPlot[603].Y = 0.063

	pointsOfFunctionPlot[604].X = -3.96
	pointsOfFunctionPlot[604].Y = 0.064

	pointsOfFunctionPlot[605].X = -3.95
	pointsOfFunctionPlot[605].Y = 0.064

	pointsOfFunctionPlot[606].X = -3.94
	pointsOfFunctionPlot[606].Y = 0.065

	pointsOfFunctionPlot[607].X = -3.93
	pointsOfFunctionPlot[607].Y = 0.065

	pointsOfFunctionPlot[608].X = -3.92
	pointsOfFunctionPlot[608].Y = 0.066

	pointsOfFunctionPlot[609].X = -3.91
	pointsOfFunctionPlot[609].Y = 0.066

	pointsOfFunctionPlot[610].X = -3.90
	pointsOfFunctionPlot[610].Y = 0.067

	pointsOfFunctionPlot[611].X = -3.89
	pointsOfFunctionPlot[611].Y = 0.067

	pointsOfFunctionPlot[612].X = -3.88
	pointsOfFunctionPlot[612].Y = 0.067

	pointsOfFunctionPlot[613].X = -3.87
	pointsOfFunctionPlot[613].Y = 0.068

	pointsOfFunctionPlot[614].X = -3.86
	pointsOfFunctionPlot[614].Y = 0.068

	pointsOfFunctionPlot[615].X = -3.85
	pointsOfFunctionPlot[615].Y = 0.069

	pointsOfFunctionPlot[616].X = -3.84
	pointsOfFunctionPlot[616].Y = 0.069

	pointsOfFunctionPlot[617].X = -3.83
	pointsOfFunctionPlot[617].Y = 0.07

	pointsOfFunctionPlot[618].X = -3.82
	pointsOfFunctionPlot[618].Y = 0.07

	pointsOfFunctionPlot[619].X = -3.81
	pointsOfFunctionPlot[619].Y = 0.071

	pointsOfFunctionPlot[620].X = -3.80
	pointsOfFunctionPlot[620].Y = 0.071

	pointsOfFunctionPlot[621].X = -3.79
	pointsOfFunctionPlot[621].Y = 0.072

	pointsOfFunctionPlot[622].X = -3.78
	pointsOfFunctionPlot[622].Y = 0.073

	pointsOfFunctionPlot[623].X = -3.77
	pointsOfFunctionPlot[623].Y = 0.073

	pointsOfFunctionPlot[624].X = -3.76
	pointsOfFunctionPlot[624].Y = 0.073

	pointsOfFunctionPlot[625].X = -3.75
	pointsOfFunctionPlot[625].Y = 0.074

	pointsOfFunctionPlot[626].X = -3.74
	pointsOfFunctionPlot[626].Y = 0.074

	pointsOfFunctionPlot[627].X = -3.73
	pointsOfFunctionPlot[627].Y = 0.075

	pointsOfFunctionPlot[628].X = -3.72
	pointsOfFunctionPlot[628].Y = 0.075

	pointsOfFunctionPlot[629].X = -3.71
	pointsOfFunctionPlot[629].Y = 0.076

	pointsOfFunctionPlot[630].X = -3.70
	pointsOfFunctionPlot[630].Y = 0.076

	pointsOfFunctionPlot[631].X = -3.69
	pointsOfFunctionPlot[631].Y = 0.077

	pointsOfFunctionPlot[632].X = -3.68
	pointsOfFunctionPlot[632].Y = 0.078

	pointsOfFunctionPlot[633].X = -3.67
	pointsOfFunctionPlot[633].Y = 0.078

	pointsOfFunctionPlot[634].X = -3.66
	pointsOfFunctionPlot[634].Y = 0.079

	pointsOfFunctionPlot[635].X = -3.65
	pointsOfFunctionPlot[635].Y = 0.079

	pointsOfFunctionPlot[636].X = -3.64
	pointsOfFunctionPlot[636].Y = 0.08

	pointsOfFunctionPlot[637].X = -3.63
	pointsOfFunctionPlot[637].Y = 0.08

	pointsOfFunctionPlot[638].X = -3.62
	pointsOfFunctionPlot[638].Y = 0.081

	pointsOfFunctionPlot[639].X = -3.61
	pointsOfFunctionPlot[639].Y = 0.081

	pointsOfFunctionPlot[640].X = -3.60
	pointsOfFunctionPlot[640].Y = 0.082

	pointsOfFunctionPlot[641].X = -3.59
	pointsOfFunctionPlot[641].Y = 0.083

	pointsOfFunctionPlot[642].X = -3.58
	pointsOfFunctionPlot[642].Y = 0.083

	pointsOfFunctionPlot[643].X = -3.57
	pointsOfFunctionPlot[643].Y = 0.084

	pointsOfFunctionPlot[644].X = -3.56
	pointsOfFunctionPlot[644].Y = 0.084

	pointsOfFunctionPlot[645].X = -3.55
	pointsOfFunctionPlot[645].Y = 0.085

	pointsOfFunctionPlot[646].X = -3.54
	pointsOfFunctionPlot[646].Y = 0.086

	pointsOfFunctionPlot[647].X = -3.53
	pointsOfFunctionPlot[647].Y = 0.086

	pointsOfFunctionPlot[648].X = -3.52
	pointsOfFunctionPlot[648].Y = 0.087

	pointsOfFunctionPlot[649].X = -3.51
	pointsOfFunctionPlot[649].Y = 0.087

	pointsOfFunctionPlot[650].X = -3.50
	pointsOfFunctionPlot[650].Y = 0.088

	pointsOfFunctionPlot[651].X = -3.49
	pointsOfFunctionPlot[651].Y = 0.089

	pointsOfFunctionPlot[652].X = -3.48
	pointsOfFunctionPlot[652].Y = 0.089

	pointsOfFunctionPlot[653].X = -3.47
	pointsOfFunctionPlot[653].Y = 0.09

	pointsOfFunctionPlot[654].X = -3.46
	pointsOfFunctionPlot[654].Y = 0.09

	pointsOfFunctionPlot[655].X = -3.45
	pointsOfFunctionPlot[655].Y = 0.091

	pointsOfFunctionPlot[656].X = -3.44
	pointsOfFunctionPlot[656].Y = 0.092

	pointsOfFunctionPlot[657].X = -3.43
	pointsOfFunctionPlot[657].Y = 0.092

	pointsOfFunctionPlot[658].X = -3.42
	pointsOfFunctionPlot[658].Y = 0.093

	pointsOfFunctionPlot[659].X = -3.41
	pointsOfFunctionPlot[659].Y = 0.094

	pointsOfFunctionPlot[660].X = -3.40
	pointsOfFunctionPlot[660].Y = 0.094

	pointsOfFunctionPlot[661].X = -3.39
	pointsOfFunctionPlot[661].Y = 0.095

	pointsOfFunctionPlot[662].X = -3.38
	pointsOfFunctionPlot[662].Y = 0.096

	pointsOfFunctionPlot[663].X = -3.37
	pointsOfFunctionPlot[663].Y = 0.096

	pointsOfFunctionPlot[664].X = -3.36
	pointsOfFunctionPlot[664].Y = 0.097

	pointsOfFunctionPlot[665].X = -3.35
	pointsOfFunctionPlot[665].Y = 0.098

	pointsOfFunctionPlot[666].X = -3.34
	pointsOfFunctionPlot[666].Y = 0.098

	pointsOfFunctionPlot[667].X = -3.33
	pointsOfFunctionPlot[667].Y = 0.099

	pointsOfFunctionPlot[668].X = -3.32
	pointsOfFunctionPlot[668].Y = 0.1

	pointsOfFunctionPlot[669].X = -3.31
	pointsOfFunctionPlot[669].Y = 0.1

	pointsOfFunctionPlot[670].X = -3.30
	pointsOfFunctionPlot[670].Y = 0.101

	pointsOfFunctionPlot[671].X = -3.29
	pointsOfFunctionPlot[671].Y = 0.102

	pointsOfFunctionPlot[672].X = -3.28
	pointsOfFunctionPlot[672].Y = 0.102

	pointsOfFunctionPlot[673].X = -3.27
	pointsOfFunctionPlot[673].Y = 0.103

	pointsOfFunctionPlot[674].X = -3.26
	pointsOfFunctionPlot[674].Y = 0.104

	pointsOfFunctionPlot[675].X = -3.25
	pointsOfFunctionPlot[675].Y = 0.105

	pointsOfFunctionPlot[676].X = -3.24
	pointsOfFunctionPlot[676].Y = 0.105

	pointsOfFunctionPlot[677].X = -3.23
	pointsOfFunctionPlot[677].Y = 0.106

	pointsOfFunctionPlot[678].X = -3.22
	pointsOfFunctionPlot[678].Y = 0.107

	pointsOfFunctionPlot[679].X = -3.21
	pointsOfFunctionPlot[679].Y = 0.108

	pointsOfFunctionPlot[680].X = -3.20
	pointsOfFunctionPlot[680].Y = 0.108

	pointsOfFunctionPlot[681].X = -3.19
	pointsOfFunctionPlot[681].Y = 0.109

	pointsOfFunctionPlot[682].X = -3.18
	pointsOfFunctionPlot[682].Y = 0.11

	pointsOfFunctionPlot[683].X = -3.17
	pointsOfFunctionPlot[683].Y = 0.111

	pointsOfFunctionPlot[684].X = -3.16
	pointsOfFunctionPlot[684].Y = 0.111

	pointsOfFunctionPlot[685].X = -3.15
	pointsOfFunctionPlot[685].Y = 0.112

	pointsOfFunctionPlot[686].X = -3.14
	pointsOfFunctionPlot[686].Y = 0.113

	pointsOfFunctionPlot[687].X = -3.13
	pointsOfFunctionPlot[687].Y = 0.114

	pointsOfFunctionPlot[688].X = -3.12
	pointsOfFunctionPlot[688].Y = 0.115

	pointsOfFunctionPlot[689].X = -3.11
	pointsOfFunctionPlot[689].Y = 0.115

	pointsOfFunctionPlot[690].X = -3.10
	pointsOfFunctionPlot[690].Y = 0.116

	pointsOfFunctionPlot[691].X = -3.09
	pointsOfFunctionPlot[691].Y = 0.117

	pointsOfFunctionPlot[692].X = -3.08
	pointsOfFunctionPlot[692].Y = 0.118

	pointsOfFunctionPlot[693].X = -3.07
	pointsOfFunctionPlot[693].Y = 0.119

	pointsOfFunctionPlot[694].X = -3.06
	pointsOfFunctionPlot[694].Y = 0.119

	pointsOfFunctionPlot[695].X = -3.05
	pointsOfFunctionPlot[695].Y = 0.12

	pointsOfFunctionPlot[696].X = -3.04
	pointsOfFunctionPlot[696].Y = 0.121

	pointsOfFunctionPlot[697].X = -3.03
	pointsOfFunctionPlot[697].Y = 0.122

	pointsOfFunctionPlot[698].X = -3.02
	pointsOfFunctionPlot[698].Y = 0.123

	pointsOfFunctionPlot[699].X = -3.01
	pointsOfFunctionPlot[699].Y = 0.124

	pointsOfFunctionPlot[700].X = -3.0
	pointsOfFunctionPlot[700].Y = 0.125

	pointsOfFunctionPlot[701].X = -2.99
	pointsOfFunctionPlot[701].Y = 0.125

	pointsOfFunctionPlot[702].X = -2.98
	pointsOfFunctionPlot[702].Y = 0.126

	pointsOfFunctionPlot[703].X = -2.97
	pointsOfFunctionPlot[703].Y = 0.127

	pointsOfFunctionPlot[704].X = -2.96
	pointsOfFunctionPlot[704].Y = 0.128

	pointsOfFunctionPlot[705].X = -2.95
	pointsOfFunctionPlot[705].Y = 0.129

	pointsOfFunctionPlot[706].X = -2.94
	pointsOfFunctionPlot[706].Y = 0.13

	pointsOfFunctionPlot[707].X = -2.93
	pointsOfFunctionPlot[707].Y = 0.131

	pointsOfFunctionPlot[708].X = -2.92
	pointsOfFunctionPlot[708].Y = 0.132

	pointsOfFunctionPlot[709].X = -2.91
	pointsOfFunctionPlot[709].Y = 0.133

	pointsOfFunctionPlot[710].X = -2.90
	pointsOfFunctionPlot[710].Y = 0.134

	pointsOfFunctionPlot[711].X = -2.89
	pointsOfFunctionPlot[711].Y = 0.134

	pointsOfFunctionPlot[712].X = -2.88
	pointsOfFunctionPlot[712].Y = 0.135

	pointsOfFunctionPlot[713].X = -2.87
	pointsOfFunctionPlot[713].Y = 0.136

	pointsOfFunctionPlot[714].X = -2.86
	pointsOfFunctionPlot[714].Y = 0.137

	pointsOfFunctionPlot[715].X = -2.85
	pointsOfFunctionPlot[715].Y = 0.138

	pointsOfFunctionPlot[716].X = -2.84
	pointsOfFunctionPlot[716].Y = 0.139

	pointsOfFunctionPlot[717].X = -2.83
	pointsOfFunctionPlot[717].Y = 0.14

	pointsOfFunctionPlot[718].X = -2.82
	pointsOfFunctionPlot[718].Y = 0.141

	pointsOfFunctionPlot[719].X = -2.81
	pointsOfFunctionPlot[719].Y = 0.142

	pointsOfFunctionPlot[720].X = -2.80
	pointsOfFunctionPlot[720].Y = 0.143

	pointsOfFunctionPlot[721].X = -2.79
	pointsOfFunctionPlot[721].Y = 0.144

	pointsOfFunctionPlot[722].X = -2.78
	pointsOfFunctionPlot[722].Y = 0.145

	pointsOfFunctionPlot[723].X = -2.77
	pointsOfFunctionPlot[723].Y = 0.146

	pointsOfFunctionPlot[724].X = -2.76
	pointsOfFunctionPlot[724].Y = 0.147

	pointsOfFunctionPlot[725].X = -2.75
	pointsOfFunctionPlot[725].Y = 0.148

	pointsOfFunctionPlot[726].X = -2.74
	pointsOfFunctionPlot[726].Y = 0.149

	pointsOfFunctionPlot[727].X = -2.73
	pointsOfFunctionPlot[727].Y = 0.15

	pointsOfFunctionPlot[728].X = -2.72
	pointsOfFunctionPlot[728].Y = 0.151

	pointsOfFunctionPlot[729].X = -2.71
	pointsOfFunctionPlot[729].Y = 0.152

	pointsOfFunctionPlot[730].X = -2.70
	pointsOfFunctionPlot[730].Y = 0.153

	pointsOfFunctionPlot[731].X = -2.69
	pointsOfFunctionPlot[731].Y = 0.155

	pointsOfFunctionPlot[732].X = -2.68
	pointsOfFunctionPlot[732].Y = 0.156

	pointsOfFunctionPlot[733].X = -2.67
	pointsOfFunctionPlot[733].Y = 0.157

	pointsOfFunctionPlot[734].X = -2.66
	pointsOfFunctionPlot[734].Y = 0.158

	pointsOfFunctionPlot[735].X = -2.65
	pointsOfFunctionPlot[735].Y = 0.159

	pointsOfFunctionPlot[736].X = -2.64
	pointsOfFunctionPlot[736].Y = 0.16

	pointsOfFunctionPlot[737].X = -2.63
	pointsOfFunctionPlot[737].Y = 0.161

	pointsOfFunctionPlot[738].X = -2.62
	pointsOfFunctionPlot[738].Y = 0.162

	pointsOfFunctionPlot[739].X = -2.61
	pointsOfFunctionPlot[739].Y = 0.163

	pointsOfFunctionPlot[740].X = -2.60
	pointsOfFunctionPlot[740].Y = 0.164

	pointsOfFunctionPlot[741].X = -2.59
	pointsOfFunctionPlot[741].Y = 0.166

	pointsOfFunctionPlot[742].X = -2.58
	pointsOfFunctionPlot[742].Y = 0.167

	pointsOfFunctionPlot[743].X = -2.57
	pointsOfFunctionPlot[743].Y = 0.168

	pointsOfFunctionPlot[744].X = -2.56
	pointsOfFunctionPlot[744].Y = 0.169

	pointsOfFunctionPlot[745].X = -2.55
	pointsOfFunctionPlot[745].Y = 0.17

	pointsOfFunctionPlot[746].X = -2.54
	pointsOfFunctionPlot[746].Y = 0.171

	pointsOfFunctionPlot[747].X = -2.53
	pointsOfFunctionPlot[747].Y = 0.173

	pointsOfFunctionPlot[748].X = -2.52
	pointsOfFunctionPlot[748].Y = 0.174

	pointsOfFunctionPlot[749].X = -2.51
	pointsOfFunctionPlot[749].Y = 0.175

	pointsOfFunctionPlot[750].X = -2.50
	pointsOfFunctionPlot[750].Y = 0.176

	pointsOfFunctionPlot[751].X = -2.49
	pointsOfFunctionPlot[751].Y = 0.178

	pointsOfFunctionPlot[752].X = -2.48
	pointsOfFunctionPlot[752].Y = 0.179

	pointsOfFunctionPlot[753].X = -2.47
	pointsOfFunctionPlot[753].Y = 0.18

	pointsOfFunctionPlot[754].X = -2.46
	pointsOfFunctionPlot[754].Y = 0.181

	pointsOfFunctionPlot[755].X = -2.45
	pointsOfFunctionPlot[755].Y = 0.183

	pointsOfFunctionPlot[756].X = -2.44
	pointsOfFunctionPlot[756].Y = 0.184

	pointsOfFunctionPlot[757].X = -2.43
	pointsOfFunctionPlot[757].Y = 0.185

	pointsOfFunctionPlot[758].X = -2.42
	pointsOfFunctionPlot[758].Y = 0.186

	pointsOfFunctionPlot[759].X = -2.41
	pointsOfFunctionPlot[759].Y = 0.188

	pointsOfFunctionPlot[760].X = -2.40
	pointsOfFunctionPlot[760].Y = 0.189

	pointsOfFunctionPlot[761].X = -2.39
	pointsOfFunctionPlot[761].Y = 0.19

	pointsOfFunctionPlot[762].X = -2.38
	pointsOfFunctionPlot[762].Y = 0.192

	pointsOfFunctionPlot[763].X = -2.37
	pointsOfFunctionPlot[763].Y = 0.193

	pointsOfFunctionPlot[764].X = -2.36
	pointsOfFunctionPlot[764].Y = 0.194

	pointsOfFunctionPlot[765].X = -2.35
	pointsOfFunctionPlot[765].Y = 0.196

	pointsOfFunctionPlot[766].X = -2.34
	pointsOfFunctionPlot[766].Y = 0.197

	pointsOfFunctionPlot[767].X = -2.33
	pointsOfFunctionPlot[767].Y = 0.198

	pointsOfFunctionPlot[768].X = -2.32
	pointsOfFunctionPlot[768].Y = 0.2

	pointsOfFunctionPlot[769].X = -2.31
	pointsOfFunctionPlot[769].Y = 0.201

	pointsOfFunctionPlot[770].X = -2.30
	pointsOfFunctionPlot[770].Y = 0.203

	pointsOfFunctionPlot[771].X = -2.29
	pointsOfFunctionPlot[771].Y = 0.204

	pointsOfFunctionPlot[772].X = -2.28
	pointsOfFunctionPlot[772].Y = 0.205

	pointsOfFunctionPlot[773].X = -2.27
	pointsOfFunctionPlot[773].Y = 0.207

	pointsOfFunctionPlot[774].X = -2.26
	pointsOfFunctionPlot[774].Y = 0.208

	pointsOfFunctionPlot[775].X = -2.25
	pointsOfFunctionPlot[775].Y = 0.21

	pointsOfFunctionPlot[776].X = -2.24
	pointsOfFunctionPlot[776].Y = 0.211

	pointsOfFunctionPlot[777].X = -2.23
	pointsOfFunctionPlot[777].Y = 0.213

	pointsOfFunctionPlot[778].X = -2.22
	pointsOfFunctionPlot[778].Y = 0.214

	pointsOfFunctionPlot[779].X = -2.21
	pointsOfFunctionPlot[779].Y = 0.216

	pointsOfFunctionPlot[780].X = -2.20
	pointsOfFunctionPlot[780].Y = 0.217

	pointsOfFunctionPlot[781].X = -2.19
	pointsOfFunctionPlot[781].Y = 0.219

	pointsOfFunctionPlot[782].X = -2.18
	pointsOfFunctionPlot[782].Y = 0.22

	pointsOfFunctionPlot[783].X = -2.17
	pointsOfFunctionPlot[783].Y = 0.222

	pointsOfFunctionPlot[784].X = -2.16
	pointsOfFunctionPlot[784].Y = 0.223

	pointsOfFunctionPlot[785].X = -2.15
	pointsOfFunctionPlot[785].Y = 0.225

	pointsOfFunctionPlot[786].X = -2.14
	pointsOfFunctionPlot[786].Y = 0.226

	pointsOfFunctionPlot[787].X = -2.13
	pointsOfFunctionPlot[787].Y = 0.228

	pointsOfFunctionPlot[788].X = -2.12
	pointsOfFunctionPlot[788].Y = 0.23

	pointsOfFunctionPlot[789].X = -2.11
	pointsOfFunctionPlot[789].Y = 0.231

	pointsOfFunctionPlot[790].X = -2.10
	pointsOfFunctionPlot[790].Y = 0.233

	pointsOfFunctionPlot[791].X = -2.09
	pointsOfFunctionPlot[791].Y = 0.234

	pointsOfFunctionPlot[792].X = -2.08
	pointsOfFunctionPlot[792].Y = 0.236

	pointsOfFunctionPlot[793].X = -2.07
	pointsOfFunctionPlot[793].Y = 0.238

	pointsOfFunctionPlot[794].X = -2.06
	pointsOfFunctionPlot[794].Y = 0.239

	pointsOfFunctionPlot[795].X = -2.05
	pointsOfFunctionPlot[795].Y = 0.241

	pointsOfFunctionPlot[796].X = -2.04
	pointsOfFunctionPlot[796].Y = 0.243

	pointsOfFunctionPlot[797].X = -2.03
	pointsOfFunctionPlot[797].Y = 0.244

	pointsOfFunctionPlot[798].X = -2.02
	pointsOfFunctionPlot[798].Y = 0.246

	pointsOfFunctionPlot[799].X = -2.01
	pointsOfFunctionPlot[799].Y = 0.248

	pointsOfFunctionPlot[800].X = -2.0
	pointsOfFunctionPlot[800].Y = 0.25

	pointsOfFunctionPlot[801].X = -1.99
	pointsOfFunctionPlot[801].Y = 0.251

	pointsOfFunctionPlot[802].X = -1.98
	pointsOfFunctionPlot[802].Y = 0.253

	pointsOfFunctionPlot[803].X = -1.97
	pointsOfFunctionPlot[803].Y = 0.255

	pointsOfFunctionPlot[804].X = -1.96
	pointsOfFunctionPlot[804].Y = 0.257

	pointsOfFunctionPlot[805].X = -1.95
	pointsOfFunctionPlot[805].Y = 0.258

	pointsOfFunctionPlot[806].X = -1.94
	pointsOfFunctionPlot[806].Y = 0.26

	pointsOfFunctionPlot[807].X = -1.93
	pointsOfFunctionPlot[807].Y = 0.262

	pointsOfFunctionPlot[808].X = -1.92
	pointsOfFunctionPlot[808].Y = 0.264

	pointsOfFunctionPlot[809].X = -1.91
	pointsOfFunctionPlot[809].Y = 0.266

	pointsOfFunctionPlot[810].X = -1.90
	pointsOfFunctionPlot[810].Y = 0.267

	pointsOfFunctionPlot[811].X = -1.89
	pointsOfFunctionPlot[811].Y = 0.269

	pointsOfFunctionPlot[812].X = -1.88
	pointsOfFunctionPlot[812].Y = 0.271

	pointsOfFunctionPlot[813].X = -1.87
	pointsOfFunctionPlot[813].Y = 0.273

	pointsOfFunctionPlot[814].X = -1.86
	pointsOfFunctionPlot[814].Y = 0.275

	pointsOfFunctionPlot[815].X = -1.85
	pointsOfFunctionPlot[815].Y = 0.277

	pointsOfFunctionPlot[816].X = -1.84
	pointsOfFunctionPlot[816].Y = 0.279

	pointsOfFunctionPlot[817].X = -1.83
	pointsOfFunctionPlot[817].Y = 0.281

	pointsOfFunctionPlot[818].X = -1.82
	pointsOfFunctionPlot[818].Y = 0.283

	pointsOfFunctionPlot[819].X = -1.81
	pointsOfFunctionPlot[819].Y = 0.285

	pointsOfFunctionPlot[820].X = -1.80
	pointsOfFunctionPlot[820].Y = 0.287

	pointsOfFunctionPlot[821].X = -1.79
	pointsOfFunctionPlot[821].Y = 0.289

	pointsOfFunctionPlot[822].X = -1.78
	pointsOfFunctionPlot[822].Y = 0.291

	pointsOfFunctionPlot[823].X = -1.77
	pointsOfFunctionPlot[823].Y = 0.293

	pointsOfFunctionPlot[824].X = -1.76
	pointsOfFunctionPlot[824].Y = 0.295

	pointsOfFunctionPlot[825].X = -1.75
	pointsOfFunctionPlot[825].Y = 0.297

	pointsOfFunctionPlot[826].X = -1.74
	pointsOfFunctionPlot[826].Y = 0.299

	pointsOfFunctionPlot[827].X = -1.73
	pointsOfFunctionPlot[827].Y = 0.301

	pointsOfFunctionPlot[828].X = -1.72
	pointsOfFunctionPlot[828].Y = 0.303

	pointsOfFunctionPlot[829].X = -1.71
	pointsOfFunctionPlot[829].Y = 0.305

	pointsOfFunctionPlot[830].X = -1.70
	pointsOfFunctionPlot[830].Y = 0.307

	pointsOfFunctionPlot[831].X = -1.69
	pointsOfFunctionPlot[831].Y = 0.309

	pointsOfFunctionPlot[832].X = -1.68
	pointsOfFunctionPlot[832].Y = 0.312

	pointsOfFunctionPlot[833].X = -1.67
	pointsOfFunctionPlot[833].Y = 0.314

	pointsOfFunctionPlot[834].X = -1.66
	pointsOfFunctionPlot[834].Y = 0.316

	pointsOfFunctionPlot[835].X = -1.65
	pointsOfFunctionPlot[835].Y = 0.318

	pointsOfFunctionPlot[836].X = -1.64
	pointsOfFunctionPlot[836].Y = 0.32

	pointsOfFunctionPlot[837].X = -1.63
	pointsOfFunctionPlot[837].Y = 0.323

	pointsOfFunctionPlot[838].X = -1.62
	pointsOfFunctionPlot[838].Y = 0.325

	pointsOfFunctionPlot[839].X = -1.61
	pointsOfFunctionPlot[839].Y = 0.327

	pointsOfFunctionPlot[840].X = -1.60
	pointsOfFunctionPlot[840].Y = 0.329

	pointsOfFunctionPlot[841].X = -1.59
	pointsOfFunctionPlot[841].Y = 0.332

	pointsOfFunctionPlot[842].X = -1.58
	pointsOfFunctionPlot[842].Y = 0.334

	pointsOfFunctionPlot[843].X = -1.57
	pointsOfFunctionPlot[843].Y = 0.336

	pointsOfFunctionPlot[844].X = -1.56
	pointsOfFunctionPlot[844].Y = 0.339

	pointsOfFunctionPlot[845].X = -1.55
	pointsOfFunctionPlot[845].Y = 0.341

	pointsOfFunctionPlot[846].X = -1.54
	pointsOfFunctionPlot[846].Y = 0.343

	pointsOfFunctionPlot[847].X = -1.53
	pointsOfFunctionPlot[847].Y = 0.346

	pointsOfFunctionPlot[848].X = -1.52
	pointsOfFunctionPlot[848].Y = 0.348

	pointsOfFunctionPlot[849].X = -1.51
	pointsOfFunctionPlot[849].Y = 0.351

	pointsOfFunctionPlot[850].X = -1.50
	pointsOfFunctionPlot[850].Y = 0.353

	pointsOfFunctionPlot[851].X = -1.49
	pointsOfFunctionPlot[851].Y = 0.356

	pointsOfFunctionPlot[852].X = -1.48
	pointsOfFunctionPlot[852].Y = 0.358

	pointsOfFunctionPlot[853].X = -1.47
	pointsOfFunctionPlot[853].Y = 0.361

	pointsOfFunctionPlot[854].X = -1.46
	pointsOfFunctionPlot[854].Y = 0.363

	pointsOfFunctionPlot[855].X = -1.45
	pointsOfFunctionPlot[855].Y = 0.366

	pointsOfFunctionPlot[856].X = -1.44
	pointsOfFunctionPlot[856].Y = 0.368

	pointsOfFunctionPlot[857].X = -1.43
	pointsOfFunctionPlot[857].Y = 0.371

	pointsOfFunctionPlot[858].X = -1.42
	pointsOfFunctionPlot[858].Y = 0.373

	pointsOfFunctionPlot[859].X = -1.41
	pointsOfFunctionPlot[859].Y = 0.376

	pointsOfFunctionPlot[860].X = -1.40
	pointsOfFunctionPlot[860].Y = 0.378

	pointsOfFunctionPlot[861].X = -1.39
	pointsOfFunctionPlot[861].Y = 0.381

	pointsOfFunctionPlot[862].X = -1.38
	pointsOfFunctionPlot[862].Y = 0.384

	pointsOfFunctionPlot[863].X = -1.37
	pointsOfFunctionPlot[863].Y = 0.386

	pointsOfFunctionPlot[864].X = -1.36
	pointsOfFunctionPlot[864].Y = 0.389

	pointsOfFunctionPlot[865].X = -1.35
	pointsOfFunctionPlot[865].Y = 0.392

	pointsOfFunctionPlot[866].X = -1.34
	pointsOfFunctionPlot[866].Y = 0.395

	pointsOfFunctionPlot[867].X = -1.33
	pointsOfFunctionPlot[867].Y = 0.397

	pointsOfFunctionPlot[868].X = -1.32
	pointsOfFunctionPlot[868].Y = 0.4

	pointsOfFunctionPlot[869].X = -1.31
	pointsOfFunctionPlot[869].Y = 0.403

	pointsOfFunctionPlot[870].X = -1.30
	pointsOfFunctionPlot[870].Y = 0.406

	pointsOfFunctionPlot[871].X = -1.29
	pointsOfFunctionPlot[871].Y = 0.409

	pointsOfFunctionPlot[872].X = -1.28
	pointsOfFunctionPlot[872].Y = 0.411

	pointsOfFunctionPlot[873].X = -1.27
	pointsOfFunctionPlot[873].Y = 0.414

	pointsOfFunctionPlot[874].X = -1.26
	pointsOfFunctionPlot[874].Y = 0.417

	pointsOfFunctionPlot[875].X = -1.25
	pointsOfFunctionPlot[875].Y = 0.42

	pointsOfFunctionPlot[876].X = -1.24
	pointsOfFunctionPlot[876].Y = 0.423

	pointsOfFunctionPlot[877].X = -1.23
	pointsOfFunctionPlot[877].Y = 0.426

	pointsOfFunctionPlot[878].X = -1.22
	pointsOfFunctionPlot[878].Y = 0.429

	pointsOfFunctionPlot[879].X = -1.21
	pointsOfFunctionPlot[879].Y = 0.432

	pointsOfFunctionPlot[880].X = -1.20
	pointsOfFunctionPlot[880].Y = 0.435

	pointsOfFunctionPlot[881].X = -1.19
	pointsOfFunctionPlot[881].Y = 0.438

	pointsOfFunctionPlot[882].X = -1.18
	pointsOfFunctionPlot[882].Y = 0.441

	pointsOfFunctionPlot[883].X = -1.17
	pointsOfFunctionPlot[883].Y = 0.444

	pointsOfFunctionPlot[884].X = -1.16
	pointsOfFunctionPlot[884].Y = 0.447

	pointsOfFunctionPlot[885].X = -1.15
	pointsOfFunctionPlot[885].Y = 0.45

	pointsOfFunctionPlot[886].X = -1.14
	pointsOfFunctionPlot[886].Y = 0.453

	pointsOfFunctionPlot[887].X = -1.13
	pointsOfFunctionPlot[887].Y = 0.456

	pointsOfFunctionPlot[888].X = -1.12
	pointsOfFunctionPlot[888].Y = 0.46

	pointsOfFunctionPlot[889].X = -1.11
	pointsOfFunctionPlot[889].Y = 0.463

	pointsOfFunctionPlot[890].X = -1.10
	pointsOfFunctionPlot[890].Y = 0.466

	pointsOfFunctionPlot[891].X = -1.09
	pointsOfFunctionPlot[891].Y = 0.469

	pointsOfFunctionPlot[892].X = -1.08
	pointsOfFunctionPlot[892].Y = 0.473

	pointsOfFunctionPlot[893].X = -1.07
	pointsOfFunctionPlot[893].Y = 0.476

	pointsOfFunctionPlot[894].X = -1.06
	pointsOfFunctionPlot[894].Y = 0.479

	pointsOfFunctionPlot[895].X = -1.05
	pointsOfFunctionPlot[895].Y = 0.483

	pointsOfFunctionPlot[896].X = -1.04
	pointsOfFunctionPlot[896].Y = 0.486

	pointsOfFunctionPlot[897].X = -1.03
	pointsOfFunctionPlot[897].Y = 0.489

	pointsOfFunctionPlot[898].X = -1.02
	pointsOfFunctionPlot[898].Y = 0.493

	pointsOfFunctionPlot[899].X = -1.01
	pointsOfFunctionPlot[899].Y = 0.496

	pointsOfFunctionPlot[900].X = -1.0
	pointsOfFunctionPlot[900].Y = 0.5

	pointsOfFunctionPlot[901].X = -0.99
	pointsOfFunctionPlot[901].Y = 0.503

	pointsOfFunctionPlot[902].X = -0.98
	pointsOfFunctionPlot[902].Y = 0.507

	pointsOfFunctionPlot[903].X = -0.97
	pointsOfFunctionPlot[903].Y = 0.51

	pointsOfFunctionPlot[904].X = -0.96
	pointsOfFunctionPlot[904].Y = 0.514

	pointsOfFunctionPlot[905].X = -0.95
	pointsOfFunctionPlot[905].Y = 0.517

	pointsOfFunctionPlot[906].X = -0.94
	pointsOfFunctionPlot[906].Y = 0.521

	pointsOfFunctionPlot[907].X = -0.93
	pointsOfFunctionPlot[907].Y = 0.524

	pointsOfFunctionPlot[908].X = -0.92
	pointsOfFunctionPlot[908].Y = 0.528

	pointsOfFunctionPlot[909].X = -0.91
	pointsOfFunctionPlot[909].Y = 0.532

	pointsOfFunctionPlot[910].X = -0.90
	pointsOfFunctionPlot[910].Y = 0.535

	pointsOfFunctionPlot[911].X = -0.89
	pointsOfFunctionPlot[911].Y = 0.539

	pointsOfFunctionPlot[912].X = -0.88
	pointsOfFunctionPlot[912].Y = 0.543

	pointsOfFunctionPlot[913].X = -0.87
	pointsOfFunctionPlot[913].Y = 0.547

	pointsOfFunctionPlot[914].X = -0.86
	pointsOfFunctionPlot[914].Y = 0.551

	pointsOfFunctionPlot[915].X = -0.85
	pointsOfFunctionPlot[915].Y = 0.554

	pointsOfFunctionPlot[916].X = -0.84
	pointsOfFunctionPlot[916].Y = 0.558

	pointsOfFunctionPlot[917].X = -0.83
	pointsOfFunctionPlot[917].Y = 0.562

	pointsOfFunctionPlot[918].X = -0.82
	pointsOfFunctionPlot[918].Y = 0.566

	pointsOfFunctionPlot[919].X = -0.81
	pointsOfFunctionPlot[919].Y = 0.57

	pointsOfFunctionPlot[920].X = -0.80
	pointsOfFunctionPlot[920].Y = 0.574

	pointsOfFunctionPlot[921].X = -0.79
	pointsOfFunctionPlot[921].Y = 0.578

	pointsOfFunctionPlot[922].X = -0.78
	pointsOfFunctionPlot[922].Y = 0.582

	pointsOfFunctionPlot[923].X = -0.77
	pointsOfFunctionPlot[923].Y = 0.586

	pointsOfFunctionPlot[924].X = -0.76
	pointsOfFunctionPlot[924].Y = 0.59

	pointsOfFunctionPlot[925].X = -0.75
	pointsOfFunctionPlot[925].Y = 0.594

	pointsOfFunctionPlot[926].X = -0.74
	pointsOfFunctionPlot[926].Y = 0.598

	pointsOfFunctionPlot[927].X = -0.73
	pointsOfFunctionPlot[927].Y = 0.602

	pointsOfFunctionPlot[928].X = -0.72
	pointsOfFunctionPlot[928].Y = 0.607

	pointsOfFunctionPlot[929].X = -0.71
	pointsOfFunctionPlot[929].Y = 0.611

	pointsOfFunctionPlot[930].X = -0.70
	pointsOfFunctionPlot[930].Y = 0.615

	pointsOfFunctionPlot[931].X = -0.69
	pointsOfFunctionPlot[931].Y = 0.619

	pointsOfFunctionPlot[932].X = -0.68
	pointsOfFunctionPlot[932].Y = 0.624

	pointsOfFunctionPlot[933].X = -0.67
	pointsOfFunctionPlot[933].Y = 0.628

	pointsOfFunctionPlot[934].X = -0.66
	pointsOfFunctionPlot[934].Y = 0.632

	pointsOfFunctionPlot[935].X = -0.65
	pointsOfFunctionPlot[935].Y = 0.637

	pointsOfFunctionPlot[936].X = -0.64
	pointsOfFunctionPlot[936].Y = 0.641

	pointsOfFunctionPlot[937].X = -0.63
	pointsOfFunctionPlot[937].Y = 0.646

	pointsOfFunctionPlot[938].X = -0.62
	pointsOfFunctionPlot[938].Y = 0.65

	pointsOfFunctionPlot[939].X = -0.61
	pointsOfFunctionPlot[939].Y = 0.655

	pointsOfFunctionPlot[940].X = -0.60
	pointsOfFunctionPlot[940].Y = 0.659

	pointsOfFunctionPlot[941].X = -0.59
	pointsOfFunctionPlot[941].Y = 0.664

	pointsOfFunctionPlot[942].X = -0.58
	pointsOfFunctionPlot[942].Y = 0.669

	pointsOfFunctionPlot[943].X = -0.57
	pointsOfFunctionPlot[943].Y = 0.673

	pointsOfFunctionPlot[944].X = -0.56
	pointsOfFunctionPlot[944].Y = 0.678

	pointsOfFunctionPlot[945].X = -0.55
	pointsOfFunctionPlot[945].Y = 0.683

	pointsOfFunctionPlot[946].X = -0.54
	pointsOfFunctionPlot[946].Y = 0.687

	pointsOfFunctionPlot[947].X = -0.53
	pointsOfFunctionPlot[947].Y = 0.692

	pointsOfFunctionPlot[948].X = -0.52
	pointsOfFunctionPlot[948].Y = 0.697

	pointsOfFunctionPlot[949].X = -0.51
	pointsOfFunctionPlot[949].Y = 0.702

	pointsOfFunctionPlot[950].X = -0.50
	pointsOfFunctionPlot[950].Y = 0.707

	pointsOfFunctionPlot[951].X = -0.49
	pointsOfFunctionPlot[951].Y = 0.712

	pointsOfFunctionPlot[952].X = -0.48
	pointsOfFunctionPlot[952].Y = 0.717

	pointsOfFunctionPlot[953].X = -0.47
	pointsOfFunctionPlot[953].Y = 0.722

	pointsOfFunctionPlot[954].X = -0.46
	pointsOfFunctionPlot[954].Y = 0.727

	pointsOfFunctionPlot[955].X = -0.45
	pointsOfFunctionPlot[955].Y = 0.732

	pointsOfFunctionPlot[956].X = -0.44
	pointsOfFunctionPlot[956].Y = 0.737

	pointsOfFunctionPlot[957].X = -0.43
	pointsOfFunctionPlot[957].Y = 0.742

	pointsOfFunctionPlot[958].X = -0.42
	pointsOfFunctionPlot[958].Y = 0.747

	pointsOfFunctionPlot[959].X = -0.41
	pointsOfFunctionPlot[959].Y = 0.752

	pointsOfFunctionPlot[960].X = -0.40
	pointsOfFunctionPlot[960].Y = 0.757

	pointsOfFunctionPlot[961].X = -0.39
	pointsOfFunctionPlot[961].Y = 0.763

	pointsOfFunctionPlot[962].X = -0.38
	pointsOfFunctionPlot[962].Y = 0.768

	pointsOfFunctionPlot[963].X = -0.37
	pointsOfFunctionPlot[963].Y = 0.773

	pointsOfFunctionPlot[964].X = -0.36
	pointsOfFunctionPlot[964].Y = 0.779

	pointsOfFunctionPlot[965].X = -0.35
	pointsOfFunctionPlot[965].Y = 0.784

	pointsOfFunctionPlot[966].X = -0.34
	pointsOfFunctionPlot[966].Y = 0.79

	pointsOfFunctionPlot[967].X = -0.33
	pointsOfFunctionPlot[967].Y = 0.795

	pointsOfFunctionPlot[968].X = -0.32
	pointsOfFunctionPlot[968].Y = 0.801

	pointsOfFunctionPlot[969].X = -0.31
	pointsOfFunctionPlot[969].Y = 0.806

	pointsOfFunctionPlot[970].X = -0.30
	pointsOfFunctionPlot[970].Y = 0.812

	pointsOfFunctionPlot[971].X = -0.29
	pointsOfFunctionPlot[971].Y = 0.817

	pointsOfFunctionPlot[972].X = -0.28
	pointsOfFunctionPlot[972].Y = 0.823

	pointsOfFunctionPlot[973].X = -0.27
	pointsOfFunctionPlot[973].Y = 0.829

	pointsOfFunctionPlot[974].X = -0.26
	pointsOfFunctionPlot[974].Y = 0.835

	pointsOfFunctionPlot[975].X = -0.25
	pointsOfFunctionPlot[975].Y = 0.84

	pointsOfFunctionPlot[976].X = -0.24
	pointsOfFunctionPlot[976].Y = 0.846

	pointsOfFunctionPlot[977].X = -0.23
	pointsOfFunctionPlot[977].Y = 0.852

	pointsOfFunctionPlot[978].X = -0.22
	pointsOfFunctionPlot[978].Y = 0.858

	pointsOfFunctionPlot[979].X = -0.21
	pointsOfFunctionPlot[979].Y = 0.864

	pointsOfFunctionPlot[980].X = -0.20
	pointsOfFunctionPlot[980].Y = 0.87

	pointsOfFunctionPlot[981].X = -0.19
	pointsOfFunctionPlot[981].Y = 0.876

	pointsOfFunctionPlot[982].X = -0.18
	pointsOfFunctionPlot[982].Y = 0.882

	pointsOfFunctionPlot[983].X = -0.17
	pointsOfFunctionPlot[983].Y = 0.888

	pointsOfFunctionPlot[984].X = -0.16
	pointsOfFunctionPlot[984].Y = 0.895

	pointsOfFunctionPlot[985].X = -0.15
	pointsOfFunctionPlot[985].Y = 0.901

	pointsOfFunctionPlot[986].X = -0.14
	pointsOfFunctionPlot[986].Y = 0.907

	pointsOfFunctionPlot[987].X = -0.13
	pointsOfFunctionPlot[987].Y = 0.913

	pointsOfFunctionPlot[988].X = -0.12
	pointsOfFunctionPlot[988].Y = 0.92

	pointsOfFunctionPlot[989].X = -0.11
	pointsOfFunctionPlot[989].Y = 0.926

	pointsOfFunctionPlot[990].X = -0.10
	pointsOfFunctionPlot[990].Y = 0.933

	pointsOfFunctionPlot[991].X = -0.09
	pointsOfFunctionPlot[991].Y = 0.946

	pointsOfFunctionPlot[992].X = -0.08
	pointsOfFunctionPlot[992].Y = 0.946

	pointsOfFunctionPlot[993].X = -0.07
	pointsOfFunctionPlot[993].Y = 0.952

	pointsOfFunctionPlot[994].X = -0.06
	pointsOfFunctionPlot[994].Y = 0.959

	pointsOfFunctionPlot[995].X = -0.05
	pointsOfFunctionPlot[995].Y = 0.965

	pointsOfFunctionPlot[996].X = -0.04
	pointsOfFunctionPlot[996].Y = 0.972

	pointsOfFunctionPlot[997].X = -0.03
	pointsOfFunctionPlot[997].Y = 0.979

	pointsOfFunctionPlot[998].X = -0.02
	pointsOfFunctionPlot[998].Y = 0.986

	pointsOfFunctionPlot[999].X = -0.01
	pointsOfFunctionPlot[999].Y = 0.993

	pointsOfFunctionPlot[1_000].X = 0.0
	pointsOfFunctionPlot[1_000].Y = 1.0

	pointsOfFunctionPlot[1_001].X = 0.01
	pointsOfFunctionPlot[1_001].Y = 1.007

	pointsOfFunctionPlot[1_002].X = 0.02
	pointsOfFunctionPlot[1_002].Y = 1.014

	pointsOfFunctionPlot[1_003].X = 0.03
	pointsOfFunctionPlot[1_003].Y = 1.021

	pointsOfFunctionPlot[1_004].X = 0.04
	pointsOfFunctionPlot[1_004].Y = 1.028

	pointsOfFunctionPlot[1_005].X = 0.05
	pointsOfFunctionPlot[1_005].Y = 1.035

	pointsOfFunctionPlot[1_006].X = 0.06
	pointsOfFunctionPlot[1_006].Y = 1.042

	pointsOfFunctionPlot[1_007].X = 0.07
	pointsOfFunctionPlot[1_007].Y = 1.049

	pointsOfFunctionPlot[1_008].X = 0.08
	pointsOfFunctionPlot[1_008].Y = 1.057

	pointsOfFunctionPlot[1_009].X = 0.09
	pointsOfFunctionPlot[1_009].Y = 1.064

	pointsOfFunctionPlot[1_010].X = 0.10
	pointsOfFunctionPlot[1_010].Y = 1.071

	pointsOfFunctionPlot[1_011].X = 0.11
	pointsOfFunctionPlot[1_011].Y = 1.079

	pointsOfFunctionPlot[1_012].X = 0.12
	pointsOfFunctionPlot[1_012].Y = 1.086

	pointsOfFunctionPlot[1_013].X = 0.13
	pointsOfFunctionPlot[1_013].Y = 1.094

	pointsOfFunctionPlot[1_014].X = 0.14
	pointsOfFunctionPlot[1_014].Y = 1.101

	pointsOfFunctionPlot[1_015].X = 0.15
	pointsOfFunctionPlot[1_015].Y = 1.109

	pointsOfFunctionPlot[1_016].X = 0.16
	pointsOfFunctionPlot[1_016].Y = 1.117

	pointsOfFunctionPlot[1_017].X = 0.17
	pointsOfFunctionPlot[1_017].Y = 1.125

	pointsOfFunctionPlot[1_018].X = 0.18
	pointsOfFunctionPlot[1_018].Y = 1.132

	pointsOfFunctionPlot[1_019].X = 0.19
	pointsOfFunctionPlot[1_019].Y = 1.14

	pointsOfFunctionPlot[1_020].X = 0.20
	pointsOfFunctionPlot[1_020].Y = 1.148

	pointsOfFunctionPlot[1_021].X = 0.21
	pointsOfFunctionPlot[1_021].Y = 1.156

	pointsOfFunctionPlot[1_022].X = 0.22
	pointsOfFunctionPlot[1_022].Y = 1.164

	pointsOfFunctionPlot[1_023].X = 0.23
	pointsOfFunctionPlot[1_023].Y = 1.172

	pointsOfFunctionPlot[1_024].X = 0.24
	pointsOfFunctionPlot[1_024].Y = 1.181

	pointsOfFunctionPlot[1_025].X = 0.25
	pointsOfFunctionPlot[1_025].Y = 1.189

	pointsOfFunctionPlot[1_026].X = 0.26
	pointsOfFunctionPlot[1_026].Y = 1.197

	pointsOfFunctionPlot[1_027].X = 0.27
	pointsOfFunctionPlot[1_027].Y = 1.205

	pointsOfFunctionPlot[1_028].X = 0.28
	pointsOfFunctionPlot[1_028].Y = 1.214

	pointsOfFunctionPlot[1_029].X = 0.29
	pointsOfFunctionPlot[1_029].Y = 1.222

	pointsOfFunctionPlot[1_030].X = 0.30
	pointsOfFunctionPlot[1_030].Y = 1.231

	pointsOfFunctionPlot[1_031].X = 0.31
	pointsOfFunctionPlot[1_031].Y = 1.239

	pointsOfFunctionPlot[1_032].X = 0.32
	pointsOfFunctionPlot[1_032].Y = 1.248

	pointsOfFunctionPlot[1_033].X = 0.33
	pointsOfFunctionPlot[1_033].Y = 1.257

	pointsOfFunctionPlot[1_034].X = 0.34
	pointsOfFunctionPlot[1_034].Y = 1.265

	pointsOfFunctionPlot[1_035].X = 0.35
	pointsOfFunctionPlot[1_035].Y = 1.274

	pointsOfFunctionPlot[1_036].X = 0.36
	pointsOfFunctionPlot[1_036].Y = 1.283

	pointsOfFunctionPlot[1_037].X = 0.37
	pointsOfFunctionPlot[1_037].Y = 1.292

	pointsOfFunctionPlot[1_038].X = 0.38
	pointsOfFunctionPlot[1_038].Y = 1.301

	pointsOfFunctionPlot[1_039].X = 0.39
	pointsOfFunctionPlot[1_039].Y = 1.31

	pointsOfFunctionPlot[1_040].X = 0.40
	pointsOfFunctionPlot[1_040].Y = 1.319

	pointsOfFunctionPlot[1_041].X = 0.41
	pointsOfFunctionPlot[1_041].Y = 1.328

	pointsOfFunctionPlot[1_042].X = 0.42
	pointsOfFunctionPlot[1_042].Y = 1.337

	pointsOfFunctionPlot[1_043].X = 0.43
	pointsOfFunctionPlot[1_043].Y = 1.347

	pointsOfFunctionPlot[1_044].X = 0.44
	pointsOfFunctionPlot[1_044].Y = 1.356

	pointsOfFunctionPlot[1_045].X = 0.45
	pointsOfFunctionPlot[1_045].Y = 1.366

	pointsOfFunctionPlot[1_046].X = 0.46
	pointsOfFunctionPlot[1_046].Y = 1.375

	pointsOfFunctionPlot[1_047].X = 0.47
	pointsOfFunctionPlot[1_047].Y = 1.385

	pointsOfFunctionPlot[1_048].X = 0.48
	pointsOfFunctionPlot[1_048].Y = 1.394

	pointsOfFunctionPlot[1_049].X = 0.49
	pointsOfFunctionPlot[1_049].Y = 1.404

	pointsOfFunctionPlot[1_050].X = 0.50
	pointsOfFunctionPlot[1_050].Y = 1.414

	pointsOfFunctionPlot[1_051].X = 0.51
	pointsOfFunctionPlot[1_051].Y = 1.424

	pointsOfFunctionPlot[1_052].X = 0.52
	pointsOfFunctionPlot[1_052].Y = 1.434

	pointsOfFunctionPlot[1_053].X = 0.53
	pointsOfFunctionPlot[1_053].Y = 1.443

	pointsOfFunctionPlot[1_054].X = 0.54
	pointsOfFunctionPlot[1_054].Y = 1.454

	pointsOfFunctionPlot[1_055].X = 0.55
	pointsOfFunctionPlot[1_055].Y = 1.464

	pointsOfFunctionPlot[1_056].X = 0.56
	pointsOfFunctionPlot[1_056].Y = 1.474

	pointsOfFunctionPlot[1_057].X = 0.57
	pointsOfFunctionPlot[1_057].Y = 1.484

	pointsOfFunctionPlot[1_058].X = 0.58
	pointsOfFunctionPlot[1_058].Y = 1.494

	pointsOfFunctionPlot[1_059].X = 0.59
	pointsOfFunctionPlot[1_059].Y = 1.505

	pointsOfFunctionPlot[1_060].X = 0.60
	pointsOfFunctionPlot[1_060].Y = 1.515

	pointsOfFunctionPlot[1_061].X = 0.61
	pointsOfFunctionPlot[1_061].Y = 1.526

	pointsOfFunctionPlot[1_062].X = 0.62
	pointsOfFunctionPlot[1_062].Y = 1.536

	pointsOfFunctionPlot[1_063].X = 0.63
	pointsOfFunctionPlot[1_063].Y = 1.547

	pointsOfFunctionPlot[1_064].X = 0.64
	pointsOfFunctionPlot[1_064].Y = 1.558

	pointsOfFunctionPlot[1_065].X = 0.65
	pointsOfFunctionPlot[1_065].Y = 1.569

	pointsOfFunctionPlot[1_066].X = 0.66
	pointsOfFunctionPlot[1_066].Y = 1.58

	pointsOfFunctionPlot[1_067].X = 0.67
	pointsOfFunctionPlot[1_067].Y = 1.591

	pointsOfFunctionPlot[1_068].X = 0.68
	pointsOfFunctionPlot[1_068].Y = 1.602

	pointsOfFunctionPlot[1_069].X = 0.69
	pointsOfFunctionPlot[1_069].Y = 1.613

	pointsOfFunctionPlot[1_070].X = 0.70
	pointsOfFunctionPlot[1_070].Y = 1.624

	pointsOfFunctionPlot[1_071].X = 0.71
	pointsOfFunctionPlot[1_071].Y = 1.635

	pointsOfFunctionPlot[1_072].X = 0.72
	pointsOfFunctionPlot[1_072].Y = 1.647

	pointsOfFunctionPlot[1_073].X = 0.73
	pointsOfFunctionPlot[1_073].Y = 1.658

	pointsOfFunctionPlot[1_074].X = 0.74
	pointsOfFunctionPlot[1_074].Y = 1.67

	pointsOfFunctionPlot[1_075].X = 0.75
	pointsOfFunctionPlot[1_075].Y = 1.681

	pointsOfFunctionPlot[1_076].X = 0.76
	pointsOfFunctionPlot[1_076].Y = 1.693

	pointsOfFunctionPlot[1_077].X = 0.77
	pointsOfFunctionPlot[1_077].Y = 1.705

	pointsOfFunctionPlot[1_078].X = 0.78
	pointsOfFunctionPlot[1_078].Y = 1.717

	pointsOfFunctionPlot[1_079].X = 0.79
	pointsOfFunctionPlot[1_079].Y = 1.729

	pointsOfFunctionPlot[1_080].X = 0.80
	pointsOfFunctionPlot[1_080].Y = 1.741

	pointsOfFunctionPlot[1_081].X = 0.81
	pointsOfFunctionPlot[1_081].Y = 1.753

	pointsOfFunctionPlot[1_082].X = 0.82
	pointsOfFunctionPlot[1_082].Y = 1.765

	pointsOfFunctionPlot[1_083].X = 0.83
	pointsOfFunctionPlot[1_083].Y = 1.777

	pointsOfFunctionPlot[1_084].X = 0.84
	pointsOfFunctionPlot[1_084].Y = 1.79

	pointsOfFunctionPlot[1_085].X = 0.85
	pointsOfFunctionPlot[1_085].Y = 1.802

	pointsOfFunctionPlot[1_086].X = 0.86
	pointsOfFunctionPlot[1_086].Y = 1.815

	pointsOfFunctionPlot[1_087].X = 0.87
	pointsOfFunctionPlot[1_087].Y = 1.827

	pointsOfFunctionPlot[1_088].X = 0.88
	pointsOfFunctionPlot[1_088].Y = 1.84

	pointsOfFunctionPlot[1_089].X = 0.89
	pointsOfFunctionPlot[1_089].Y = 1.853

	pointsOfFunctionPlot[1_090].X = 0.90
	pointsOfFunctionPlot[1_090].Y = 1.866

	pointsOfFunctionPlot[1_091].X = 0.91
	pointsOfFunctionPlot[1_091].Y = 1.879

	pointsOfFunctionPlot[1_092].X = 0.92
	pointsOfFunctionPlot[1_092].Y = 1.892

	pointsOfFunctionPlot[1_093].X = 0.93
	pointsOfFunctionPlot[1_093].Y = 1.905

	pointsOfFunctionPlot[1_094].X = 0.94
	pointsOfFunctionPlot[1_094].Y = 1.918

	pointsOfFunctionPlot[1_095].X = 0.95
	pointsOfFunctionPlot[1_095].Y = 1.931

	pointsOfFunctionPlot[1_096].X = 0.96
	pointsOfFunctionPlot[1_096].Y = 1.945

	pointsOfFunctionPlot[1_097].X = 0.97
	pointsOfFunctionPlot[1_097].Y = 1.958

	pointsOfFunctionPlot[1_098].X = 0.98
	pointsOfFunctionPlot[1_098].Y = 1.972

	pointsOfFunctionPlot[1_099].X = 0.99
	pointsOfFunctionPlot[1_099].Y = 1.986

	pointsOfFunctionPlot[1_100].X = 1.0
	pointsOfFunctionPlot[1_100].Y = 2.0

	pointsOfFunctionPlot[1_101].X = 1.01
	pointsOfFunctionPlot[1_101].Y = 2.013

	pointsOfFunctionPlot[1_102].X = 1.02
	pointsOfFunctionPlot[1_102].Y = 2.027

	pointsOfFunctionPlot[1_103].X = 1.03
	pointsOfFunctionPlot[1_103].Y = 2.042

	pointsOfFunctionPlot[1_104].X = 1.04
	pointsOfFunctionPlot[1_104].Y = 2.056

	pointsOfFunctionPlot[1_105].X = 1.05
	pointsOfFunctionPlot[1_105].Y = 2.07

	pointsOfFunctionPlot[1_106].X = 1.06
	pointsOfFunctionPlot[1_106].Y = 2.084

	pointsOfFunctionPlot[1_107].X = 1.07
	pointsOfFunctionPlot[1_107].Y = 2.099

	pointsOfFunctionPlot[1_108].X = 1.08
	pointsOfFunctionPlot[1_108].Y = 2.114

	pointsOfFunctionPlot[1_109].X = 1.09
	pointsOfFunctionPlot[1_109].Y = 2.128

	pointsOfFunctionPlot[1_110].X = 1.10
	pointsOfFunctionPlot[1_110].Y = 2.143

	pointsOfFunctionPlot[1_111].X = 1.11
	pointsOfFunctionPlot[1_111].Y = 2.158

	pointsOfFunctionPlot[1_112].X = 1.12
	pointsOfFunctionPlot[1_112].Y = 2.173

	pointsOfFunctionPlot[1_113].X = 1.13
	pointsOfFunctionPlot[1_113].Y = 2.188

	pointsOfFunctionPlot[1_114].X = 1.14
	pointsOfFunctionPlot[1_114].Y = 2.203

	pointsOfFunctionPlot[1_115].X = 1.15
	pointsOfFunctionPlot[1_115].Y = 2.219

	pointsOfFunctionPlot[1_116].X = 1.16
	pointsOfFunctionPlot[1_116].Y = 2.234

	pointsOfFunctionPlot[1_117].X = 1.17
	pointsOfFunctionPlot[1_117].Y = 2.25

	pointsOfFunctionPlot[1_118].X = 1.18
	pointsOfFunctionPlot[1_118].Y = 2.265

	pointsOfFunctionPlot[1_119].X = 1.19
	pointsOfFunctionPlot[1_119].Y = 2.281

	pointsOfFunctionPlot[1_120].X = 1.20
	pointsOfFunctionPlot[1_120].Y = 2.297

	pointsOfFunctionPlot[1_121].X = 1.21
	pointsOfFunctionPlot[1_121].Y = 2.313

	pointsOfFunctionPlot[1_122].X = 1.22
	pointsOfFunctionPlot[1_122].Y = 2.329

	pointsOfFunctionPlot[1_123].X = 1.23
	pointsOfFunctionPlot[1_123].Y = 2.345

	pointsOfFunctionPlot[1_124].X = 1.24
	pointsOfFunctionPlot[1_124].Y = 2.362

	pointsOfFunctionPlot[1_125].X = 1.25
	pointsOfFunctionPlot[1_125].Y = 2.378

	pointsOfFunctionPlot[1_126].X = 1.26
	pointsOfFunctionPlot[1_126].Y = 2.395

	pointsOfFunctionPlot[1_127].X = 1.27
	pointsOfFunctionPlot[1_127].Y = 2.411

	pointsOfFunctionPlot[1_128].X = 1.28
	pointsOfFunctionPlot[1_128].Y = 2.428

	pointsOfFunctionPlot[1_129].X = 1.29
	pointsOfFunctionPlot[1_129].Y = 2.445

	pointsOfFunctionPlot[1_130].X = 1.30
	pointsOfFunctionPlot[1_130].Y = 2.462

	pointsOfFunctionPlot[1_131].X = 1.31
	pointsOfFunctionPlot[1_131].Y = 2.479

	pointsOfFunctionPlot[1_132].X = 1.32
	pointsOfFunctionPlot[1_132].Y = 2.496

	pointsOfFunctionPlot[1_133].X = 1.33
	pointsOfFunctionPlot[1_133].Y = 2.514

	pointsOfFunctionPlot[1_134].X = 1.34
	pointsOfFunctionPlot[1_134].Y = 2.531

	pointsOfFunctionPlot[1_135].X = 1.35
	pointsOfFunctionPlot[1_135].Y = 2.549

	pointsOfFunctionPlot[1_136].X = 1.36
	pointsOfFunctionPlot[1_136].Y = 2.566

	pointsOfFunctionPlot[1_137].X = 1.37
	pointsOfFunctionPlot[1_137].Y = 2.584

	pointsOfFunctionPlot[1_138].X = 1.38
	pointsOfFunctionPlot[1_138].Y = 2.602

	pointsOfFunctionPlot[1_139].X = 1.39
	pointsOfFunctionPlot[1_139].Y = 2.62

	pointsOfFunctionPlot[1_140].X = 1.40
	pointsOfFunctionPlot[1_140].Y = 2.639

	pointsOfFunctionPlot[1_141].X = 1.41
	pointsOfFunctionPlot[1_141].Y = 2.657

	pointsOfFunctionPlot[1_142].X = 1.42
	pointsOfFunctionPlot[1_142].Y = 2.675

	pointsOfFunctionPlot[1_143].X = 1.43
	pointsOfFunctionPlot[1_143].Y = 2.694

	pointsOfFunctionPlot[1_144].X = 1.44
	pointsOfFunctionPlot[1_144].Y = 2.713

	pointsOfFunctionPlot[1_145].X = 1.45
	pointsOfFunctionPlot[1_145].Y = 2.732

	pointsOfFunctionPlot[1_146].X = 1.46
	pointsOfFunctionPlot[1_146].Y = 2.751

	pointsOfFunctionPlot[1_147].X = 1.47
	pointsOfFunctionPlot[1_147].Y = 2.77

	pointsOfFunctionPlot[1_148].X = 1.48
	pointsOfFunctionPlot[1_148].Y = 2.789

	pointsOfFunctionPlot[1_149].X = 1.49
	pointsOfFunctionPlot[1_149].Y = 2.808

	pointsOfFunctionPlot[1_150].X = 1.50
	pointsOfFunctionPlot[1_150].Y = 2.828

	pointsOfFunctionPlot[1_151].X = 1.51
	pointsOfFunctionPlot[1_151].Y = 2.848

	pointsOfFunctionPlot[1_152].X = 1.52
	pointsOfFunctionPlot[1_152].Y = 2.867

	pointsOfFunctionPlot[1_153].X = 1.53
	pointsOfFunctionPlot[1_153].Y = 2.887

	pointsOfFunctionPlot[1_154].X = 1.54
	pointsOfFunctionPlot[1_154].Y = 2.907

	pointsOfFunctionPlot[1_155].X = 1.55
	pointsOfFunctionPlot[1_155].Y = 2.928

	pointsOfFunctionPlot[1_156].X = 1.56
	pointsOfFunctionPlot[1_156].Y = 2.948

	pointsOfFunctionPlot[1_157].X = 1.57
	pointsOfFunctionPlot[1_157].Y = 2.969

	pointsOfFunctionPlot[1_158].X = 1.58
	pointsOfFunctionPlot[1_158].Y = 2.989

	pointsOfFunctionPlot[1_159].X = 1.59
	pointsOfFunctionPlot[1_159].Y = 3.01

	pointsOfFunctionPlot[1_160].X = 1.60
	pointsOfFunctionPlot[1_160].Y = 3.031

	pointsOfFunctionPlot[1_161].X = 1.61
	pointsOfFunctionPlot[1_161].Y = 3.052

	pointsOfFunctionPlot[1_162].X = 1.62
	pointsOfFunctionPlot[1_162].Y = 3.073

	pointsOfFunctionPlot[1_163].X = 1.63
	pointsOfFunctionPlot[1_163].Y = 3.095

	pointsOfFunctionPlot[1_164].X = 1.64
	pointsOfFunctionPlot[1_164].Y = 3.116

	pointsOfFunctionPlot[1_165].X = 1.65
	pointsOfFunctionPlot[1_165].Y = 3.138

	pointsOfFunctionPlot[1_166].X = 1.66
	pointsOfFunctionPlot[1_166].Y = 3.16

	pointsOfFunctionPlot[1_167].X = 1.67
	pointsOfFunctionPlot[1_167].Y = 3.182

	pointsOfFunctionPlot[1_168].X = 1.68
	pointsOfFunctionPlot[1_168].Y = 3.204

	pointsOfFunctionPlot[1_169].X = 1.69
	pointsOfFunctionPlot[1_169].Y = 3.226

	pointsOfFunctionPlot[1_170].X = 1.70
	pointsOfFunctionPlot[1_170].Y = 3.249

	pointsOfFunctionPlot[1_171].X = 1.71
	pointsOfFunctionPlot[1_171].Y = 3.271

	pointsOfFunctionPlot[1_172].X = 1.72
	pointsOfFunctionPlot[1_172].Y = 3.294

	pointsOfFunctionPlot[1_173].X = 1.73
	pointsOfFunctionPlot[1_173].Y = 3.317

	pointsOfFunctionPlot[1_174].X = 1.74
	pointsOfFunctionPlot[1_174].Y = 3.34

	pointsOfFunctionPlot[1_175].X = 1.75
	pointsOfFunctionPlot[1_175].Y = 3.363

	pointsOfFunctionPlot[1_176].X = 1.76
	pointsOfFunctionPlot[1_176].Y = 3.387

	pointsOfFunctionPlot[1_177].X = 1.77
	pointsOfFunctionPlot[1_177].Y = 3.41

	pointsOfFunctionPlot[1_178].X = 1.78
	pointsOfFunctionPlot[1_178].Y = 3.434

	pointsOfFunctionPlot[1_179].X = 1.79
	pointsOfFunctionPlot[1_179].Y = 3.458

	pointsOfFunctionPlot[1_180].X = 1.80
	pointsOfFunctionPlot[1_180].Y = 3.482

	pointsOfFunctionPlot[1_181].X = 1.81
	pointsOfFunctionPlot[1_181].Y = 3.506

	pointsOfFunctionPlot[1_182].X = 1.82
	pointsOfFunctionPlot[1_182].Y = 3.53

	pointsOfFunctionPlot[1_183].X = 1.83
	pointsOfFunctionPlot[1_183].Y = 3.555

	pointsOfFunctionPlot[1_184].X = 1.84
	pointsOfFunctionPlot[1_184].Y = 3.58

	pointsOfFunctionPlot[1_185].X = 1.85
	pointsOfFunctionPlot[1_185].Y = 3.605

	pointsOfFunctionPlot[1_186].X = 1.86
	pointsOfFunctionPlot[1_186].Y = 3.63

	pointsOfFunctionPlot[1_187].X = 1.87
	pointsOfFunctionPlot[1_187].Y = 3.655

	pointsOfFunctionPlot[1_188].X = 1.88
	pointsOfFunctionPlot[1_188].Y = 3.68

	pointsOfFunctionPlot[1_189].X = 1.89
	pointsOfFunctionPlot[1_189].Y = 3.706

	pointsOfFunctionPlot[1_190].X = 1.90
	pointsOfFunctionPlot[1_190].Y = 3.732

	pointsOfFunctionPlot[1_191].X = 1.91
	pointsOfFunctionPlot[1_191].Y = 3.758

	pointsOfFunctionPlot[1_192].X = 1.92
	pointsOfFunctionPlot[1_192].Y = 3.784

	pointsOfFunctionPlot[1_193].X = 1.93
	pointsOfFunctionPlot[1_193].Y = 3.81

	pointsOfFunctionPlot[1_194].X = 1.94
	pointsOfFunctionPlot[1_194].Y = 3.837

	pointsOfFunctionPlot[1_195].X = 1.95
	pointsOfFunctionPlot[1_195].Y = 3.863

	pointsOfFunctionPlot[1_196].X = 1.96
	pointsOfFunctionPlot[1_196].Y = 3.89

	pointsOfFunctionPlot[1_197].X = 1.97
	pointsOfFunctionPlot[1_197].Y = 3.917

	pointsOfFunctionPlot[1_198].X = 1.98
	pointsOfFunctionPlot[1_198].Y = 3.944

	pointsOfFunctionPlot[1_199].X = 1.99
	pointsOfFunctionPlot[1_199].Y = 3.972

	pointsOfFunctionPlot[1_200].X = 2.0
	pointsOfFunctionPlot[1_200].Y = 4.0

	pointsOfFunctionPlot[1_201].X = 2.01
	pointsOfFunctionPlot[1_201].Y = 4.027

	pointsOfFunctionPlot[1_202].X = 2.02
	pointsOfFunctionPlot[1_202].Y = 4.055

	pointsOfFunctionPlot[1_203].X = 2.03
	pointsOfFunctionPlot[1_203].Y = 4.084

	pointsOfFunctionPlot[1_204].X = 2.04
	pointsOfFunctionPlot[1_204].Y = 4.112

	pointsOfFunctionPlot[1_205].X = 2.05
	pointsOfFunctionPlot[1_205].Y = 4.141

	pointsOfFunctionPlot[1_206].X = 2.06
	pointsOfFunctionPlot[1_206].Y = 4.169

	pointsOfFunctionPlot[1_207].X = 2.07
	pointsOfFunctionPlot[1_207].Y = 4.198

	pointsOfFunctionPlot[1_208].X = 2.08
	pointsOfFunctionPlot[1_208].Y = 4.228

	pointsOfFunctionPlot[1_209].X = 2.09
	pointsOfFunctionPlot[1_209].Y = 4.257

	pointsOfFunctionPlot[1_210].X = 2.10
	pointsOfFunctionPlot[1_210].Y = 4.287

	pointsOfFunctionPlot[1_211].X = 2.11
	pointsOfFunctionPlot[1_211].Y = 4.316

	pointsOfFunctionPlot[1_212].X = 2.12
	pointsOfFunctionPlot[1_212].Y = 4.346

	pointsOfFunctionPlot[1_213].X = 2.13
	pointsOfFunctionPlot[1_213].Y = 4.377

	pointsOfFunctionPlot[1_214].X = 2.14
	pointsOfFunctionPlot[1_214].Y = 4.407

	pointsOfFunctionPlot[1_215].X = 2.15
	pointsOfFunctionPlot[1_215].Y = 4.438

	pointsOfFunctionPlot[1_216].X = 2.16
	pointsOfFunctionPlot[1_216].Y = 4.469

	pointsOfFunctionPlot[1_217].X = 2.17
	pointsOfFunctionPlot[1_217].Y = 4.5

	pointsOfFunctionPlot[1_218].X = 2.18
	pointsOfFunctionPlot[1_218].Y = 4.531

	pointsOfFunctionPlot[1_219].X = 2.19
	pointsOfFunctionPlot[1_219].Y = 4.563

	pointsOfFunctionPlot[1_220].X = 2.20
	pointsOfFunctionPlot[1_220].Y = 4.594

	pointsOfFunctionPlot[1_221].X = 2.21
	pointsOfFunctionPlot[1_221].Y = 4.626

	pointsOfFunctionPlot[1_222].X = 2.22
	pointsOfFunctionPlot[1_222].Y = 4.658

	pointsOfFunctionPlot[1_223].X = 2.23
	pointsOfFunctionPlot[1_223].Y = 4.691

	pointsOfFunctionPlot[1_224].X = 2.24
	pointsOfFunctionPlot[1_224].Y = 4.724

	pointsOfFunctionPlot[1_225].X = 2.25
	pointsOfFunctionPlot[1_225].Y = 4.756

	pointsOfFunctionPlot[1_226].X = 2.26
	pointsOfFunctionPlot[1_226].Y = 4.789

	pointsOfFunctionPlot[1_227].X = 2.27
	pointsOfFunctionPlot[1_227].Y = 4.823

	pointsOfFunctionPlot[1_228].X = 2.28
	pointsOfFunctionPlot[1_228].Y = 4.856

	pointsOfFunctionPlot[1_229].X = 2.29
	pointsOfFunctionPlot[1_229].Y = 4.89

	pointsOfFunctionPlot[1_230].X = 2.30
	pointsOfFunctionPlot[1_230].Y = 4.924

	pointsOfFunctionPlot[1_231].X = 2.31
	pointsOfFunctionPlot[1_231].Y = 4.958

	pointsOfFunctionPlot[1_232].X = 2.32
	pointsOfFunctionPlot[1_232].Y = 4.993

	pointsOfFunctionPlot[1_233].X = 2.33
	pointsOfFunctionPlot[1_233].Y = 5.028

	pointsOfFunctionPlot[1_234].X = 2.34
	pointsOfFunctionPlot[1_234].Y = 5.063

	pointsOfFunctionPlot[1_235].X = 2.35
	pointsOfFunctionPlot[1_235].Y = 5.098

	pointsOfFunctionPlot[1_236].X = 2.36
	pointsOfFunctionPlot[1_236].Y = 5.133

	pointsOfFunctionPlot[1_237].X = 2.37
	pointsOfFunctionPlot[1_237].Y = 5.169

	pointsOfFunctionPlot[1_238].X = 2.38
	pointsOfFunctionPlot[1_238].Y = 5.205

	pointsOfFunctionPlot[1_239].X = 2.39
	pointsOfFunctionPlot[1_239].Y = 5.241

	pointsOfFunctionPlot[1_240].X = 2.40
	pointsOfFunctionPlot[1_240].Y = 5.278

	pointsOfFunctionPlot[1_241].X = 2.41
	pointsOfFunctionPlot[1_241].Y = 5.314

	pointsOfFunctionPlot[1_242].X = 2.42
	pointsOfFunctionPlot[1_242].Y = 5.351

	pointsOfFunctionPlot[1_243].X = 2.43
	pointsOfFunctionPlot[1_243].Y = 5.388

	pointsOfFunctionPlot[1_244].X = 2.44
	pointsOfFunctionPlot[1_244].Y = 5.426

	pointsOfFunctionPlot[1_245].X = 2.45
	pointsOfFunctionPlot[1_245].Y = 5.464

	pointsOfFunctionPlot[1_246].X = 2.46
	pointsOfFunctionPlot[1_246].Y = 5.502

	pointsOfFunctionPlot[1_247].X = 2.47
	pointsOfFunctionPlot[1_247].Y = 5.54

	pointsOfFunctionPlot[1_248].X = 2.48
	pointsOfFunctionPlot[1_248].Y = 5.579

	pointsOfFunctionPlot[1_249].X = 2.49
	pointsOfFunctionPlot[1_249].Y = 5.617

	pointsOfFunctionPlot[1_250].X = 2.50
	pointsOfFunctionPlot[1_250].Y = 5.656

	pointsOfFunctionPlot[1_251].X = 2.51
	pointsOfFunctionPlot[1_251].Y = 5.696

	pointsOfFunctionPlot[1_252].X = 2.52
	pointsOfFunctionPlot[1_252].Y = 5.735

	pointsOfFunctionPlot[1_253].X = 2.53
	pointsOfFunctionPlot[1_253].Y = 5.775

	pointsOfFunctionPlot[1_254].X = 2.54
	pointsOfFunctionPlot[1_254].Y = 5.815

	pointsOfFunctionPlot[1_255].X = 2.55
	pointsOfFunctionPlot[1_255].Y = 5.856

	pointsOfFunctionPlot[1_256].X = 2.56
	pointsOfFunctionPlot[1_256].Y = 5.897

	pointsOfFunctionPlot[1_257].X = 2.57
	pointsOfFunctionPlot[1_257].Y = 5.938

	pointsOfFunctionPlot[1_258].X = 2.58
	pointsOfFunctionPlot[1_258].Y = 5.979

	pointsOfFunctionPlot[1_259].X = 2.59
	pointsOfFunctionPlot[1_259].Y = 6.021

	pointsOfFunctionPlot[1_260].X = 2.60
	pointsOfFunctionPlot[1_260].Y = 6.062

	pointsOfFunctionPlot[1_261].X = 2.61
	pointsOfFunctionPlot[1_261].Y = 6.105

	pointsOfFunctionPlot[1_262].X = 2.62
	pointsOfFunctionPlot[1_262].Y = 6.147

	pointsOfFunctionPlot[1_263].X = 2.63
	pointsOfFunctionPlot[1_263].Y = 6.19

	pointsOfFunctionPlot[1_264].X = 2.64
	pointsOfFunctionPlot[1_264].Y = 6.233

	pointsOfFunctionPlot[1_265].X = 2.65
	pointsOfFunctionPlot[1_265].Y = 6.276

	pointsOfFunctionPlot[1_266].X = 2.66
	pointsOfFunctionPlot[1_266].Y = 6.32

	pointsOfFunctionPlot[1_267].X = 2.67
	pointsOfFunctionPlot[1_267].Y = 6.364

	pointsOfFunctionPlot[1_268].X = 2.68
	pointsOfFunctionPlot[1_268].Y = 6.408

	pointsOfFunctionPlot[1_269].X = 2.69
	pointsOfFunctionPlot[1_269].Y = 6.453

	pointsOfFunctionPlot[1_270].X = 2.70
	pointsOfFunctionPlot[1_270].Y = 6.498

	pointsOfFunctionPlot[1_271].X = 2.71
	pointsOfFunctionPlot[1_271].Y = 6.543

	pointsOfFunctionPlot[1_272].X = 2.72
	pointsOfFunctionPlot[1_272].Y = 6.588

	pointsOfFunctionPlot[1_273].X = 2.73
	pointsOfFunctionPlot[1_273].Y = 6.634

	pointsOfFunctionPlot[1_274].X = 2.74
	pointsOfFunctionPlot[1_274].Y = 6.68

	pointsOfFunctionPlot[1_275].X = 2.75
	pointsOfFunctionPlot[1_275].Y = 6.727

	pointsOfFunctionPlot[1_276].X = 2.76
	pointsOfFunctionPlot[1_276].Y = 6.774

	pointsOfFunctionPlot[1_277].X = 2.77
	pointsOfFunctionPlot[1_277].Y = 6.821

	pointsOfFunctionPlot[1_278].X = 2.78
	pointsOfFunctionPlot[1_278].Y = 6.868

	pointsOfFunctionPlot[1_279].X = 2.79
	pointsOfFunctionPlot[1_279].Y = 6.916

	pointsOfFunctionPlot[1_280].X = 2.80
	pointsOfFunctionPlot[1_280].Y = 6.964

	pointsOfFunctionPlot[1_281].X = 2.81
	pointsOfFunctionPlot[1_281].Y = 7.012

	pointsOfFunctionPlot[1_282].X = 2.82
	pointsOfFunctionPlot[1_282].Y = 7.061

	pointsOfFunctionPlot[1_283].X = 2.83
	pointsOfFunctionPlot[1_283].Y = 7.11

	pointsOfFunctionPlot[1_284].X = 2.84
	pointsOfFunctionPlot[1_284].Y = 7.16

	pointsOfFunctionPlot[1_285].X = 2.85
	pointsOfFunctionPlot[1_285].Y = 7.21

	pointsOfFunctionPlot[1_286].X = 2.86
	pointsOfFunctionPlot[1_286].Y = 7.26

	pointsOfFunctionPlot[1_287].X = 2.87
	pointsOfFunctionPlot[1_287].Y = 7.31

	pointsOfFunctionPlot[1_288].X = 2.88
	pointsOfFunctionPlot[1_288].Y = 7.361

	pointsOfFunctionPlot[1_289].X = 2.89
	pointsOfFunctionPlot[1_289].Y = 7.412

	pointsOfFunctionPlot[1_290].X = 2.90
	pointsOfFunctionPlot[1_290].Y = 7.464

	pointsOfFunctionPlot[1_291].X = 2.91
	pointsOfFunctionPlot[1_291].Y = 7.516

	pointsOfFunctionPlot[1_292].X = 2.92
	pointsOfFunctionPlot[1_292].Y = 7.568

	pointsOfFunctionPlot[1_293].X = 2.93
	pointsOfFunctionPlot[1_293].Y = 7.621

	pointsOfFunctionPlot[1_294].X = 2.94
	pointsOfFunctionPlot[1_294].Y = 7.674

	pointsOfFunctionPlot[1_295].X = 2.95
	pointsOfFunctionPlot[1_295].Y = 7.727

	pointsOfFunctionPlot[1_296].X = 2.96
	pointsOfFunctionPlot[1_296].Y = 7.781

	pointsOfFunctionPlot[1_297].X = 2.97
	pointsOfFunctionPlot[1_297].Y = 7.835

	pointsOfFunctionPlot[1_298].X = 2.98
	pointsOfFunctionPlot[1_298].Y = 7.889

	pointsOfFunctionPlot[1_299].X = 2.99
	pointsOfFunctionPlot[1_299].Y = 7.944

	pointsOfFunctionPlot[1_300].X = 3.0
	pointsOfFunctionPlot[1_300].Y = 8.0

	pointsOfFunctionPlot[1_301].X = 3.01
	pointsOfFunctionPlot[1_301].Y = 8.055

	pointsOfFunctionPlot[1_302].X = 3.02
	pointsOfFunctionPlot[1_302].Y = 8.111

	pointsOfFunctionPlot[1_303].X = 3.03
	pointsOfFunctionPlot[1_303].Y = 8.168

	pointsOfFunctionPlot[1_304].X = 3.04
	pointsOfFunctionPlot[1_304].Y = 8.224

	pointsOfFunctionPlot[1_305].X = 3.05
	pointsOfFunctionPlot[1_305].Y = 8.282

	pointsOfFunctionPlot[1_306].X = 3.06
	pointsOfFunctionPlot[1_306].Y = 8.339

	pointsOfFunctionPlot[1_307].X = 3.07
	pointsOfFunctionPlot[1_307].Y = 8.397

	pointsOfFunctionPlot[1_308].X = 3.08
	pointsOfFunctionPlot[1_308].Y = 8.456

	pointsOfFunctionPlot[1_309].X = 3.09
	pointsOfFunctionPlot[1_309].Y = 8.515

	pointsOfFunctionPlot[1_310].X = 3.10
	pointsOfFunctionPlot[1_310].Y = 8.574

	pointsOfFunctionPlot[1_311].X = 3.11
	pointsOfFunctionPlot[1_311].Y = 8.633

	pointsOfFunctionPlot[1_312].X = 3.12
	pointsOfFunctionPlot[1_312].Y = 8.693

	pointsOfFunctionPlot[1_313].X = 3.13
	pointsOfFunctionPlot[1_313].Y = 8.754

	pointsOfFunctionPlot[1_314].X = 3.14
	pointsOfFunctionPlot[1_314].Y = 8.815

	pointsOfFunctionPlot[1_315].X = 3.15
	pointsOfFunctionPlot[1_315].Y = 8.876

	pointsOfFunctionPlot[1_316].X = 3.16
	pointsOfFunctionPlot[1_316].Y = 8.938

	pointsOfFunctionPlot[1_317].X = 3.17
	pointsOfFunctionPlot[1_317].Y = 9.005

	pointsOfFunctionPlot[1_318].X = 3.18
	pointsOfFunctionPlot[1_318].Y = 9.063

	pointsOfFunctionPlot[1_319].X = 3.19
	pointsOfFunctionPlot[1_319].Y = 9.126

	pointsOfFunctionPlot[1_320].X = 3.20
	pointsOfFunctionPlot[1_320].Y = 9.189

	pointsOfFunctionPlot[1_321].X = 3.21
	pointsOfFunctionPlot[1_321].Y = 9.253

	pointsOfFunctionPlot[1_322].X = 3.22
	pointsOfFunctionPlot[1_322].Y = 9.317

	pointsOfFunctionPlot[1_323].X = 3.23
	pointsOfFunctionPlot[1_323].Y = 9.382

	pointsOfFunctionPlot[1_324].X = 3.24
	pointsOfFunctionPlot[1_324].Y = 9.447

	pointsOfFunctionPlot[1_325].X = 3.25
	pointsOfFunctionPlot[1_325].Y = 9.513

	pointsOfFunctionPlot[1_326].X = 3.26
	pointsOfFunctionPlot[1_326].Y = 9.579

	pointsOfFunctionPlot[1_327].X = 3.27
	pointsOfFunctionPlot[1_327].Y = 9.646

	pointsOfFunctionPlot[1_328].X = 3.28
	pointsOfFunctionPlot[1_328].Y = 9.713

	pointsOfFunctionPlot[1_329].X = 3.29
	pointsOfFunctionPlot[1_329].Y = 9.781

	pointsOfFunctionPlot[1_330].X = 3.30
	pointsOfFunctionPlot[1_330].Y = 9.849

	pointsOfFunctionPlot[1_331].X = 3.31
	pointsOfFunctionPlot[1_331].Y = 9.917

	pointsOfFunctionPlot[1_332].X = 3.32
	pointsOfFunctionPlot[1_332].Y = 9.986

	pointsOfFunctionPlot[1_333].X = 3.33
	pointsOfFunctionPlot[1_333].Y = 10.056

	pointsOfFunctionPlot[1_334].X = 3.34
	pointsOfFunctionPlot[1_334].Y = 10.126

	pointsOfFunctionPlot[1_335].X = 3.35
	pointsOfFunctionPlot[1_335].Y = 10.196

	pointsOfFunctionPlot[1_336].X = 3.36
	pointsOfFunctionPlot[1_336].Y = 10.267

	pointsOfFunctionPlot[1_337].X = 3.37
	pointsOfFunctionPlot[1_337].Y = 10.338

	pointsOfFunctionPlot[1_338].X = 3.38
	pointsOfFunctionPlot[1_338].Y = 10.41

	pointsOfFunctionPlot[1_339].X = 3.39
	pointsOfFunctionPlot[1_339].Y = 10.483

	pointsOfFunctionPlot[1_340].X = 3.40
	pointsOfFunctionPlot[1_340].Y = 10.556

	pointsOfFunctionPlot[1_341].X = 3.41
	pointsOfFunctionPlot[1_341].Y = 10.629

	pointsOfFunctionPlot[1_342].X = 3.42
	pointsOfFunctionPlot[1_342].Y = 10.703

	pointsOfFunctionPlot[1_343].X = 3.43
	pointsOfFunctionPlot[1_343].Y = 10.777

	pointsOfFunctionPlot[1_344].X = 3.44
	pointsOfFunctionPlot[1_344].Y = 10.852

	pointsOfFunctionPlot[1_345].X = 3.45
	pointsOfFunctionPlot[1_345].Y = 10.928

	pointsOfFunctionPlot[1_346].X = 3.46
	pointsOfFunctionPlot[1_346].Y = 11.004

	pointsOfFunctionPlot[1_347].X = 3.47
	pointsOfFunctionPlot[1_347].Y = 11.08

	pointsOfFunctionPlot[1_348].X = 3.48
	pointsOfFunctionPlot[1_348].Y = 11.157

	pointsOfFunctionPlot[1_349].X = 3.49
	pointsOfFunctionPlot[1_349].Y = 11.235

	pointsOfFunctionPlot[1_350].X = 3.50
	pointsOfFunctionPlot[1_350].Y = 11.313

	pointsOfFunctionPlot[1_351].X = 3.51
	pointsOfFunctionPlot[1_351].Y = 11.393

	pointsOfFunctionPlot[1_352].X = 3.52
	pointsOfFunctionPlot[1_352].Y = 11.471

	pointsOfFunctionPlot[1_353].X = 3.53
	pointsOfFunctionPlot[1_353].Y = 11.551

	pointsOfFunctionPlot[1_354].X = 3.54
	pointsOfFunctionPlot[1_354].Y = 11.631

	pointsOfFunctionPlot[1_355].X = 3.55
	pointsOfFunctionPlot[1_355].Y = 11.712

	pointsOfFunctionPlot[1_356].X = 3.56
	pointsOfFunctionPlot[1_356].Y = 11.794

	pointsOfFunctionPlot[1_357].X = 3.57
	pointsOfFunctionPlot[1_357].Y = 11.876

	pointsOfFunctionPlot[1_358].X = 3.58
	pointsOfFunctionPlot[1_358].Y = 11.958

	pointsOfFunctionPlot[1_359].X = 3.59
	pointsOfFunctionPlot[1_359].Y = 12.042

	pointsOfFunctionPlot[1_360].X = 3.60
	pointsOfFunctionPlot[1_360].Y = 12.125

	pointsOfFunctionPlot[1_361].X = 3.61
	pointsOfFunctionPlot[1_361].Y = 12.21

	pointsOfFunctionPlot[1_362].X = 3.62
	pointsOfFunctionPlot[1_362].Y = 12.295

	pointsOfFunctionPlot[1_363].X = 3.63
	pointsOfFunctionPlot[1_363].Y = 12.38

	pointsOfFunctionPlot[1_364].X = 3.64
	pointsOfFunctionPlot[1_364].Y = 12.466

	pointsOfFunctionPlot[1_365].X = 3.65
	pointsOfFunctionPlot[1_365].Y = 12.553

	pointsOfFunctionPlot[1_366].X = 3.66
	pointsOfFunctionPlot[1_366].Y = 12.64

	pointsOfFunctionPlot[1_367].X = 3.67
	pointsOfFunctionPlot[1_367].Y = 12.728

	pointsOfFunctionPlot[1_368].X = 3.68
	pointsOfFunctionPlot[1_368].Y = 12.817

	pointsOfFunctionPlot[1_369].X = 3.69
	pointsOfFunctionPlot[1_369].Y = 12.906

	pointsOfFunctionPlot[1_370].X = 3.70
	pointsOfFunctionPlot[1_370].Y = 12.996

	pointsOfFunctionPlot[1_371].X = 3.71
	pointsOfFunctionPlot[1_371].Y = 13.086

	pointsOfFunctionPlot[1_372].X = 3.72
	pointsOfFunctionPlot[1_372].Y = 13.177

	pointsOfFunctionPlot[1_373].X = 3.73
	pointsOfFunctionPlot[1_373].Y = 13.269

	pointsOfFunctionPlot[1_374].X = 3.74
	pointsOfFunctionPlot[1_374].Y = 13.361

	pointsOfFunctionPlot[1_375].X = 3.75
	pointsOfFunctionPlot[1_375].Y = 13.454

	pointsOfFunctionPlot[1_376].X = 3.76
	pointsOfFunctionPlot[1_376].Y = 13.547

	pointsOfFunctionPlot[1_377].X = 3.77
	pointsOfFunctionPlot[1_377].Y = 13.642

	pointsOfFunctionPlot[1_378].X = 3.78
	pointsOfFunctionPlot[1_378].Y = 13.737

	pointsOfFunctionPlot[1_379].X = 3.79
	pointsOfFunctionPlot[1_379].Y = 13.832

	pointsOfFunctionPlot[1_380].X = 3.80
	pointsOfFunctionPlot[1_380].Y = 13.928

	pointsOfFunctionPlot[1_381].X = 3.81
	pointsOfFunctionPlot[1_381].Y = 14.025

	pointsOfFunctionPlot[1_382].X = 3.82
	pointsOfFunctionPlot[1_382].Y = 14.123

	pointsOfFunctionPlot[1_383].X = 3.83
	pointsOfFunctionPlot[1_383].Y = 14.221

	pointsOfFunctionPlot[1_384].X = 3.84
	pointsOfFunctionPlot[1_384].Y = 14.32

	pointsOfFunctionPlot[1_385].X = 3.85
	pointsOfFunctionPlot[1_385].Y = 14.42

	pointsOfFunctionPlot[1_386].X = 3.86
	pointsOfFunctionPlot[1_386].Y = 14.52

	pointsOfFunctionPlot[1_387].X = 3.87
	pointsOfFunctionPlot[1_387].Y = 14.621

	pointsOfFunctionPlot[1_388].X = 3.88
	pointsOfFunctionPlot[1_388].Y = 14.723

	pointsOfFunctionPlot[1_389].X = 3.89
	pointsOfFunctionPlot[1_389].Y = 14.825

	pointsOfFunctionPlot[1_390].X = 3.90
	pointsOfFunctionPlot[1_390].Y = 14.928

	pointsOfFunctionPlot[1_391].X = 3.91
	pointsOfFunctionPlot[1_391].Y = 15.032

	pointsOfFunctionPlot[1_392].X = 3.92
	pointsOfFunctionPlot[1_392].Y = 15.136

	pointsOfFunctionPlot[1_393].X = 3.93
	pointsOfFunctionPlot[1_393].Y = 15.242

	pointsOfFunctionPlot[1_394].X = 3.94
	pointsOfFunctionPlot[1_394].Y = 15.348

	pointsOfFunctionPlot[1_395].X = 3.95
	pointsOfFunctionPlot[1_395].Y = 15.455

	pointsOfFunctionPlot[1_396].X = 3.96
	pointsOfFunctionPlot[1_396].Y = 15.562

	pointsOfFunctionPlot[1_397].X = 3.97
	pointsOfFunctionPlot[1_397].Y = 15.67

	pointsOfFunctionPlot[1_398].X = 3.98
	pointsOfFunctionPlot[1_398].Y = 15.779

	pointsOfFunctionPlot[1_399].X = 3.99
	pointsOfFunctionPlot[1_399].Y = 15.889

	pointsOfFunctionPlot[1_400].X = 4.0
	pointsOfFunctionPlot[1_400].Y = 16.0

	pointsOfFunctionPlot[1_401].X = 4.01
	pointsOfFunctionPlot[1_401].Y = 16.111

	pointsOfFunctionPlot[1_402].X = 4.02
	pointsOfFunctionPlot[1_402].Y = 16.223

	pointsOfFunctionPlot[1_403].X = 4.03
	pointsOfFunctionPlot[1_403].Y = 16.336

	pointsOfFunctionPlot[1_404].X = 4.04
	pointsOfFunctionPlot[1_404].Y = 16.449

	pointsOfFunctionPlot[1_405].X = 4.05
	pointsOfFunctionPlot[1_405].Y = 16.564

	pointsOfFunctionPlot[1_406].X = 4.06
	pointsOfFunctionPlot[1_406].Y = 16.679

	pointsOfFunctionPlot[1_407].X = 4.07
	pointsOfFunctionPlot[1_407].Y = 16.795

	pointsOfFunctionPlot[1_408].X = 4.08
	pointsOfFunctionPlot[1_408].Y = 16.912

	pointsOfFunctionPlot[1_409].X = 4.09
	pointsOfFunctionPlot[1_409].Y = 17.029

	pointsOfFunctionPlot[1_410].X = 4.10
	pointsOfFunctionPlot[1_410].Y = 17.148

	pointsOfFunctionPlot[1_411].X = 4.11
	pointsOfFunctionPlot[1_411].Y = 17.267

	pointsOfFunctionPlot[1_412].X = 4.12
	pointsOfFunctionPlot[1_412].Y = 17.387

	pointsOfFunctionPlot[1_413].X = 4.13
	pointsOfFunctionPlot[1_413].Y = 17.508

	pointsOfFunctionPlot[1_414].X = 4.14
	pointsOfFunctionPlot[1_414].Y = 17.63

	pointsOfFunctionPlot[1_415].X = 4.15
	pointsOfFunctionPlot[1_415].Y = 17.753

	pointsOfFunctionPlot[1_416].X = 4.16
	pointsOfFunctionPlot[1_416].Y = 17.876

	pointsOfFunctionPlot[1_417].X = 4.17
	pointsOfFunctionPlot[1_417].Y = 18.0

	pointsOfFunctionPlot[1_418].X = 4.18
	pointsOfFunctionPlot[1_418].Y = 18.126

	pointsOfFunctionPlot[1_419].X = 4.19
	pointsOfFunctionPlot[1_419].Y = 18.252

	pointsOfFunctionPlot[1_420].X = 4.20
	pointsOfFunctionPlot[1_420].Y = 18.379

	pointsOfFunctionPlot[1_421].X = 4.21
	pointsOfFunctionPlot[1_421].Y = 18.507

	pointsOfFunctionPlot[1_422].X = 4.22
	pointsOfFunctionPlot[1_422].Y = 18.635

	pointsOfFunctionPlot[1_423].X = 4.23
	pointsOfFunctionPlot[1_423].Y = 18.765

	pointsOfFunctionPlot[1_424].X = 4.24
	pointsOfFunctionPlot[1_424].Y = 18.895

	pointsOfFunctionPlot[1_425].X = 4.25
	pointsOfFunctionPlot[1_425].Y = 19.027

	pointsOfFunctionPlot[1_426].X = 4.26
	pointsOfFunctionPlot[1_426].Y = 19.159

	pointsOfFunctionPlot[1_427].X = 4.27
	pointsOfFunctionPlot[1_427].Y = 19.292

	pointsOfFunctionPlot[1_428].X = 4.28
	pointsOfFunctionPlot[1_428].Y = 19.427

	pointsOfFunctionPlot[1_429].X = 4.29
	pointsOfFunctionPlot[1_429].Y = 19.562

	pointsOfFunctionPlot[1_430].X = 4.30
	pointsOfFunctionPlot[1_430].Y = 19.698

	pointsOfFunctionPlot[1_431].X = 4.31
	pointsOfFunctionPlot[1_431].Y = 19.835

	pointsOfFunctionPlot[1_432].X = 4.32
	pointsOfFunctionPlot[1_432].Y = 19.973

	pointsOfFunctionPlot[1_433].X = 4.33
	pointsOfFunctionPlot[1_433].Y = 20.112

	pointsOfFunctionPlot[1_434].X = 4.34
	pointsOfFunctionPlot[1_434].Y = 20.252

	pointsOfFunctionPlot[1_435].X = 4.35
	pointsOfFunctionPlot[1_435].Y = 20.393

	pointsOfFunctionPlot[1_436].X = 4.36
	pointsOfFunctionPlot[1_436].Y = 20.534

	pointsOfFunctionPlot[1_437].X = 4.37
	pointsOfFunctionPlot[1_437].Y = 20.677

	pointsOfFunctionPlot[1_438].X = 4.38
	pointsOfFunctionPlot[1_438].Y = 20.821

	pointsOfFunctionPlot[1_439].X = 4.39
	pointsOfFunctionPlot[1_439].Y = 20.966

	pointsOfFunctionPlot[1_440].X = 4.40
	pointsOfFunctionPlot[1_440].Y = 21.112

	pointsOfFunctionPlot[1_441].X = 4.41
	pointsOfFunctionPlot[1_441].Y = 21.259

	pointsOfFunctionPlot[1_442].X = 4.42
	pointsOfFunctionPlot[1_442].Y = 21.406

	pointsOfFunctionPlot[1_443].X = 4.43
	pointsOfFunctionPlot[1_443].Y = 21.555

	pointsOfFunctionPlot[1_444].X = 4.44
	pointsOfFunctionPlot[1_444].Y = 21.705

	pointsOfFunctionPlot[1_445].X = 4.45
	pointsOfFunctionPlot[1_445].Y = 21.856

	pointsOfFunctionPlot[1_446].X = 4.46
	pointsOfFunctionPlot[1_446].Y = 22.008

	pointsOfFunctionPlot[1_447].X = 4.47
	pointsOfFunctionPlot[1_447].Y = 22.161

	pointsOfFunctionPlot[1_448].X = 4.48
	pointsOfFunctionPlot[1_448].Y = 22.315

	pointsOfFunctionPlot[1_449].X = 4.49
	pointsOfFunctionPlot[1_449].Y = 22.471

	pointsOfFunctionPlot[1_450].X = 4.50
	pointsOfFunctionPlot[1_450].Y = 22.627

	pointsOfFunctionPlot[1_451].X = 4.51
	pointsOfFunctionPlot[1_451].Y = 22.784

	pointsOfFunctionPlot[1_452].X = 4.52
	pointsOfFunctionPlot[1_452].Y = 22.943

	pointsOfFunctionPlot[1_453].X = 4.53
	pointsOfFunctionPlot[1_453].Y = 23.102

	pointsOfFunctionPlot[1_454].X = 4.54
	pointsOfFunctionPlot[1_454].Y = 23.263

	pointsOfFunctionPlot[1_455].X = 4.55
	pointsOfFunctionPlot[1_455].Y = 23.425

	pointsOfFunctionPlot[1_456].X = 4.56
	pointsOfFunctionPlot[1_456].Y = 23.588

	pointsOfFunctionPlot[1_457].X = 4.57
	pointsOfFunctionPlot[1_457].Y = 23.752

	pointsOfFunctionPlot[1_458].X = 4.58
	pointsOfFunctionPlot[1_458].Y = 23.917

	pointsOfFunctionPlot[1_459].X = 4.59
	pointsOfFunctionPlot[1_459].Y = 24.083

	pointsOfFunctionPlot[1_460].X = 4.60
	pointsOfFunctionPlot[1_460].Y = 24.251

	pointsOfFunctionPlot[1_461].X = 4.61
	pointsOfFunctionPlot[1_461].Y = 24.42

	pointsOfFunctionPlot[1_462].X = 4.62
	pointsOfFunctionPlot[1_462].Y = 24.59

	pointsOfFunctionPlot[1_463].X = 4.63
	pointsOfFunctionPlot[1_463].Y = 24.761

	pointsOfFunctionPlot[1_464].X = 4.64
	pointsOfFunctionPlot[1_464].Y = 24.933

	pointsOfFunctionPlot[1_465].X = 4.65
	pointsOfFunctionPlot[1_465].Y = 25.106

	pointsOfFunctionPlot[1_466].X = 4.66
	pointsOfFunctionPlot[1_466].Y = 25.281

	pointsOfFunctionPlot[1_467].X = 4.67
	pointsOfFunctionPlot[1_467].Y = 25.457

	pointsOfFunctionPlot[1_468].X = 4.68
	pointsOfFunctionPlot[1_468].Y = 25.634

	pointsOfFunctionPlot[1_469].X = 4.69
	pointsOfFunctionPlot[1_469].Y = 25.812

	pointsOfFunctionPlot[1_470].X = 4.70
	pointsOfFunctionPlot[1_470].Y = 25.992

	pointsOfFunctionPlot[1_471].X = 4.71
	pointsOfFunctionPlot[1_471].Y = 26.172

	pointsOfFunctionPlot[1_472].X = 4.72
	pointsOfFunctionPlot[1_472].Y = 26.354

	pointsOfFunctionPlot[1_473].X = 4.73
	pointsOfFunctionPlot[1_473].Y = 26.538

	pointsOfFunctionPlot[1_474].X = 4.74
	pointsOfFunctionPlot[1_474].Y = 26.722

	pointsOfFunctionPlot[1_475].X = 4.75
	pointsOfFunctionPlot[1_475].Y = 26.908

	pointsOfFunctionPlot[1_476].X = 4.76
	pointsOfFunctionPlot[1_476].Y = 27.095

	pointsOfFunctionPlot[1_477].X = 4.77
	pointsOfFunctionPlot[1_477].Y = 27.284

	pointsOfFunctionPlot[1_478].X = 4.78
	pointsOfFunctionPlot[1_478].Y = 27.474

	pointsOfFunctionPlot[1_479].X = 4.79
	pointsOfFunctionPlot[1_479].Y = 27.665

	pointsOfFunctionPlot[1_480].X = 4.80
	pointsOfFunctionPlot[1_480].Y = 27.857

	pointsOfFunctionPlot[1_481].X = 4.81
	pointsOfFunctionPlot[1_481].Y = 28.051

	pointsOfFunctionPlot[1_482].X = 4.82
	pointsOfFunctionPlot[1_482].Y = 28.246

	pointsOfFunctionPlot[1_483].X = 4.83
	pointsOfFunctionPlot[1_483].Y = 28.443

	pointsOfFunctionPlot[1_484].X = 4.84
	pointsOfFunctionPlot[1_484].Y = 28.64

	pointsOfFunctionPlot[1_485].X = 4.85
	pointsOfFunctionPlot[1_485].Y = 28.84

	pointsOfFunctionPlot[1_486].X = 4.86
	pointsOfFunctionPlot[1_486].Y = 29.04

	pointsOfFunctionPlot[1_487].X = 4.87
	pointsOfFunctionPlot[1_487].Y = 29.242

	pointsOfFunctionPlot[1_488].X = 4.88
	pointsOfFunctionPlot[1_488].Y = 29.446

	pointsOfFunctionPlot[1_489].X = 4.89
	pointsOfFunctionPlot[1_489].Y = 29.65

	pointsOfFunctionPlot[1_490].X = 4.90
	pointsOfFunctionPlot[1_490].Y = 29.857

	pointsOfFunctionPlot[1_491].X = 4.91
	pointsOfFunctionPlot[1_491].Y = 30.064

	pointsOfFunctionPlot[1_492].X = 4.92
	pointsOfFunctionPlot[1_492].Y = 30.273

	pointsOfFunctionPlot[1_493].X = 4.93
	pointsOfFunctionPlot[1_493].Y = 30.484

	pointsOfFunctionPlot[1_494].X = 4.94
	pointsOfFunctionPlot[1_494].Y = 30.696

	pointsOfFunctionPlot[1_495].X = 4.95
	pointsOfFunctionPlot[1_495].Y = 30.91

	pointsOfFunctionPlot[1_496].X = 4.96
	pointsOfFunctionPlot[1_496].Y = 31.125

	pointsOfFunctionPlot[1_497].X = 4.97
	pointsOfFunctionPlot[1_497].Y = 31.341

	pointsOfFunctionPlot[1_498].X = 4.98
	pointsOfFunctionPlot[1_498].Y = 31.559

	pointsOfFunctionPlot[1_499].X = 4.99
	pointsOfFunctionPlot[1_499].Y = 31.779

	pointsOfFunctionPlot[1_500].X = 5.0
	pointsOfFunctionPlot[1_500].Y = 32.0

	pointsOfFunctionPlot[1_501].X = 5.01
	pointsOfFunctionPlot[1_501].Y = 32.222

	pointsOfFunctionPlot[1_502].X = 5.02
	pointsOfFunctionPlot[1_502].Y = 32.446

	pointsOfFunctionPlot[1_503].X = 5.03
	pointsOfFunctionPlot[1_503].Y = 32.672

	pointsOfFunctionPlot[1_504].X = 5.04
	pointsOfFunctionPlot[1_504].Y = 32.899

	pointsOfFunctionPlot[1_505].X = 5.05
	pointsOfFunctionPlot[1_505].Y = 33.128

	pointsOfFunctionPlot[1_506].X = 5.06
	pointsOfFunctionPlot[1_506].Y = 33.358

	pointsOfFunctionPlot[1_507].X = 5.07
	pointsOfFunctionPlot[1_507].Y = 33.59

	pointsOfFunctionPlot[1_508].X = 5.08
	pointsOfFunctionPlot[1_508].Y = 33.824

	pointsOfFunctionPlot[1_509].X = 5.09
	pointsOfFunctionPlot[1_509].Y = 34.059

	pointsOfFunctionPlot[1_510].X = 5.10
	pointsOfFunctionPlot[1_510].Y = 34.296

	pointsOfFunctionPlot[1_511].X = 5.11
	pointsOfFunctionPlot[1_511].Y = 34.535

	pointsOfFunctionPlot[1_512].X = 5.12
	pointsOfFunctionPlot[1_512].Y = 34.775

	pointsOfFunctionPlot[1_513].X = 5.13
	pointsOfFunctionPlot[1_513].Y = 35.017

	pointsOfFunctionPlot[1_514].X = 5.14
	pointsOfFunctionPlot[1_514].Y = 35.261

	pointsOfFunctionPlot[1_515].X = 5.15
	pointsOfFunctionPlot[1_515].Y = 35.506

	pointsOfFunctionPlot[1_516].X = 5.16
	pointsOfFunctionPlot[1_516].Y = 35.753

	pointsOfFunctionPlot[1_517].X = 5.17
	pointsOfFunctionPlot[1_517].Y = 36.001

	pointsOfFunctionPlot[1_518].X = 5.18
	pointsOfFunctionPlot[1_518].Y = 36.252

	pointsOfFunctionPlot[1_519].X = 5.19
	pointsOfFunctionPlot[1_519].Y = 36.504

	pointsOfFunctionPlot[1_520].X = 5.20
	pointsOfFunctionPlot[1_520].Y = 36.758

	pointsOfFunctionPlot[1_521].X = 5.21
	pointsOfFunctionPlot[1_521].Y = 37.014

	pointsOfFunctionPlot[1_522].X = 5.22
	pointsOfFunctionPlot[1_522].Y = 37.271

	pointsOfFunctionPlot[1_523].X = 5.23
	pointsOfFunctionPlot[1_523].Y = 37.53

	pointsOfFunctionPlot[1_524].X = 5.24
	pointsOfFunctionPlot[1_524].Y = 37.791

	pointsOfFunctionPlot[1_525].X = 5.25
	pointsOfFunctionPlot[1_525].Y = 38.054

	pointsOfFunctionPlot[1_526].X = 5.26
	pointsOfFunctionPlot[1_526].Y = 38.319

	pointsOfFunctionPlot[1_527].X = 5.27
	pointsOfFunctionPlot[1_527].Y = 38.585

	pointsOfFunctionPlot[1_528].X = 5.28
	pointsOfFunctionPlot[1_528].Y = 38.854

	pointsOfFunctionPlot[1_529].X = 5.29
	pointsOfFunctionPlot[1_529].Y = 39.124

	pointsOfFunctionPlot[1_530].X = 5.30
	pointsOfFunctionPlot[1_530].Y = 39.396

	pointsOfFunctionPlot[1_531].X = 5.31
	pointsOfFunctionPlot[1_531].Y = 39.67

	pointsOfFunctionPlot[1_532].X = 5.32
	pointsOfFunctionPlot[1_532].Y = 39.946

	pointsOfFunctionPlot[1_533].X = 5.33
	pointsOfFunctionPlot[1_533].Y = 40.224

	pointsOfFunctionPlot[1_534].X = 5.34
	pointsOfFunctionPlot[1_534].Y = 40.504

	pointsOfFunctionPlot[1_535].X = 5.35
	pointsOfFunctionPlot[1_535].Y = 40.785

	pointsOfFunctionPlot[1_536].X = 5.36
	pointsOfFunctionPlot[1_536].Y = 41.069

	pointsOfFunctionPlot[1_537].X = 5.37
	pointsOfFunctionPlot[1_537].Y = 41.355

	pointsOfFunctionPlot[1_538].X = 5.38
	pointsOfFunctionPlot[1_538].Y = 41.642

	pointsOfFunctionPlot[1_539].X = 5.39
	pointsOfFunctionPlot[1_539].Y = 41.932

	pointsOfFunctionPlot[1_540].X = 5.40
	pointsOfFunctionPlot[1_540].Y = 42.224

	pointsOfFunctionPlot[1_541].X = 5.41
	pointsOfFunctionPlot[1_541].Y = 42.517

	pointsOfFunctionPlot[1_542].X = 5.42
	pointsOfFunctionPlot[1_542].Y = 42.813

	pointsOfFunctionPlot[1_543].X = 5.43
	pointsOfFunctionPlot[1_543].Y = 43.111

	pointsOfFunctionPlot[1_544].X = 5.44
	pointsOfFunctionPlot[1_544].Y = 43.411

	pointsOfFunctionPlot[1_545].X = 5.45
	pointsOfFunctionPlot[1_545].Y = 43.713

	pointsOfFunctionPlot[1_546].X = 5.46
	pointsOfFunctionPlot[1_546].Y = 44.017

	pointsOfFunctionPlot[1_547].X = 5.47
	pointsOfFunctionPlot[1_547].Y = 44.323

	pointsOfFunctionPlot[1_548].X = 5.48
	pointsOfFunctionPlot[1_548].Y = 44.631

	pointsOfFunctionPlot[1_549].X = 5.49
	pointsOfFunctionPlot[1_549].Y = 44.942

	pointsOfFunctionPlot[1_550].X = 5.50
	pointsOfFunctionPlot[1_550].Y = 45.254

	pointsOfFunctionPlot[1_551].X = 5.51
	pointsOfFunctionPlot[1_551].Y = 45.569

	pointsOfFunctionPlot[1_552].X = 5.52
	pointsOfFunctionPlot[1_552].Y = 45.886

	pointsOfFunctionPlot[1_553].X = 5.53
	pointsOfFunctionPlot[1_553].Y = 46.205

	pointsOfFunctionPlot[1_554].X = 5.54
	pointsOfFunctionPlot[1_554].Y = 46.527

	pointsOfFunctionPlot[1_555].X = 5.55
	pointsOfFunctionPlot[1_555].Y = 46.85

	pointsOfFunctionPlot[1_556].X = 5.56
	pointsOfFunctionPlot[1_556].Y = 47.176

	pointsOfFunctionPlot[1_557].X = 5.57
	pointsOfFunctionPlot[1_557].Y = 47.504

	pointsOfFunctionPlot[1_558].X = 5.58
	pointsOfFunctionPlot[1_558].Y = 47.835

	pointsOfFunctionPlot[1_559].X = 5.59
	pointsOfFunctionPlot[1_559].Y = 48.167

	pointsOfFunctionPlot[1_560].X = 5.60
	pointsOfFunctionPlot[1_560].Y = 48.502

	pointsOfFunctionPlot[1_561].X = 5.61
	pointsOfFunctionPlot[1_561].Y = 48.84

	pointsOfFunctionPlot[1_562].X = 5.62
	pointsOfFunctionPlot[1_562].Y = 49.18

	pointsOfFunctionPlot[1_563].X = 5.63
	pointsOfFunctionPlot[1_563].Y = 49.522

	pointsOfFunctionPlot[1_564].X = 5.64
	pointsOfFunctionPlot[1_564].Y = 49.866

	pointsOfFunctionPlot[1_565].X = 5.65
	pointsOfFunctionPlot[1_565].Y = 50.213

	pointsOfFunctionPlot[1_566].X = 5.66
	pointsOfFunctionPlot[1_566].Y = 50.562

	pointsOfFunctionPlot[1_567].X = 5.67
	pointsOfFunctionPlot[1_567].Y = 50.914

	pointsOfFunctionPlot[1_568].X = 5.68
	pointsOfFunctionPlot[1_568].Y = 51.268

	pointsOfFunctionPlot[1_569].X = 5.69
	pointsOfFunctionPlot[1_569].Y = 51.625

	pointsOfFunctionPlot[1_570].X = 5.70
	pointsOfFunctionPlot[1_570].Y = 51.984

	pointsOfFunctionPlot[1_571].X = 5.71
	pointsOfFunctionPlot[1_571].Y = 52.345

	pointsOfFunctionPlot[1_572].X = 5.72
	pointsOfFunctionPlot[1_572].Y = 52.709

	pointsOfFunctionPlot[1_573].X = 5.73
	pointsOfFunctionPlot[1_573].Y = 53.076

	pointsOfFunctionPlot[1_574].X = 5.74
	pointsOfFunctionPlot[1_574].Y = 53.445

	pointsOfFunctionPlot[1_575].X = 5.75
	pointsOfFunctionPlot[1_575].Y = 53.817

	pointsOfFunctionPlot[1_576].X = 5.76
	pointsOfFunctionPlot[1_576].Y = 54.191

	pointsOfFunctionPlot[1_577].X = 5.77
	pointsOfFunctionPlot[1_577].Y = 54.568

	pointsOfFunctionPlot[1_578].X = 5.78
	pointsOfFunctionPlot[1_578].Y = 54.948

	pointsOfFunctionPlot[1_579].X = 5.79
	pointsOfFunctionPlot[1_579].Y = 55.33

	pointsOfFunctionPlot[1_580].X = 5.80
	pointsOfFunctionPlot[1_580].Y = 55.715

	pointsOfFunctionPlot[1_581].X = 5.81
	pointsOfFunctionPlot[1_581].Y = 56.102

	pointsOfFunctionPlot[1_582].X = 5.82
	pointsOfFunctionPlot[1_582].Y = 56.493

	pointsOfFunctionPlot[1_583].X = 5.83
	pointsOfFunctionPlot[1_583].Y = 56.885

	pointsOfFunctionPlot[1_584].X = 5.84
	pointsOfFunctionPlot[1_584].Y = 57.281

	pointsOfFunctionPlot[1_585].X = 5.85
	pointsOfFunctionPlot[1_585].Y = 57.68

	pointsOfFunctionPlot[1_586].X = 5.86
	pointsOfFunctionPlot[1_586].Y = 58.081

	pointsOfFunctionPlot[1_587].X = 5.87
	pointsOfFunctionPlot[1_587].Y = 58.485

	pointsOfFunctionPlot[1_588].X = 5.88
	pointsOfFunctionPlot[1_588].Y = 58.892

	pointsOfFunctionPlot[1_589].X = 5.89
	pointsOfFunctionPlot[1_589].Y = 59.301

	pointsOfFunctionPlot[1_590].X = 5.90
	pointsOfFunctionPlot[1_590].Y = 59.714

	pointsOfFunctionPlot[1_591].X = 5.91
	pointsOfFunctionPlot[1_591].Y = 60.129

	pointsOfFunctionPlot[1_592].X = 5.92
	pointsOfFunctionPlot[1_592].Y = 60.547

	pointsOfFunctionPlot[1_593].X = 5.93
	pointsOfFunctionPlot[1_593].Y = 60.968

	pointsOfFunctionPlot[1_594].X = 5.94
	pointsOfFunctionPlot[1_594].Y = 61.392

	pointsOfFunctionPlot[1_595].X = 5.95
	pointsOfFunctionPlot[1_595].Y = 61.819

	pointsOfFunctionPlot[1_596].X = 5.96
	pointsOfFunctionPlot[1_596].Y = 62.249

	pointsOfFunctionPlot[1_597].X = 5.97
	pointsOfFunctionPlot[1_597].Y = 62.682

	pointsOfFunctionPlot[1_598].X = 5.98
	pointsOfFunctionPlot[1_598].Y = 63.118

	pointsOfFunctionPlot[1_599].X = 5.99
	pointsOfFunctionPlot[1_599].Y = 63.557

	pointsOfFunctionPlot[1_600].X = 6.0
	pointsOfFunctionPlot[1_600].Y = 64.0

	pointsOfFunctionPlot[1_601].X = 6.01
	pointsOfFunctionPlot[1_601].Y = 64.445

	pointsOfFunctionPlot[1_602].X = 6.02
	pointsOfFunctionPlot[1_602].Y = 64.893

	pointsOfFunctionPlot[1_603].X = 6.03
	pointsOfFunctionPlot[1_603].Y = 65.344

	pointsOfFunctionPlot[1_604].X = 6.04
	pointsOfFunctionPlot[1_604].Y = 65.799

	pointsOfFunctionPlot[1_605].X = 6.05
	pointsOfFunctionPlot[1_605].Y = 66.257

	pointsOfFunctionPlot[1_606].X = 6.06
	pointsOfFunctionPlot[1_606].Y = 66.717

	pointsOfFunctionPlot[1_607].X = 6.07
	pointsOfFunctionPlot[1_607].Y = 67.181

	pointsOfFunctionPlot[1_608].X = 6.08
	pointsOfFunctionPlot[1_608].Y = 67.649

	pointsOfFunctionPlot[1_609].X = 6.09
	pointsOfFunctionPlot[1_609].Y = 68.119

	pointsOfFunctionPlot[1_610].X = 6.10
	pointsOfFunctionPlot[1_610].Y = 68.593

	pointsOfFunctionPlot[1_611].X = 6.11
	pointsOfFunctionPlot[1_611].Y = 69.07

	pointsOfFunctionPlot[1_612].X = 6.12
	pointsOfFunctionPlot[1_612].Y = 69.551

	pointsOfFunctionPlot[1_613].X = 6.13
	pointsOfFunctionPlot[1_613].Y = 70.034

	pointsOfFunctionPlot[1_614].X = 6.14
	pointsOfFunctionPlot[1_614].Y = 70.521

	pointsOfFunctionPlot[1_615].X = 6.15
	pointsOfFunctionPlot[1_615].Y = 71.012

	pointsOfFunctionPlot[1_616].X = 6.16
	pointsOfFunctionPlot[1_616].Y = 71.506

	pointsOfFunctionPlot[1_617].X = 6.17
	pointsOfFunctionPlot[1_617].Y = 72.003

	pointsOfFunctionPlot[1_618].X = 6.18
	pointsOfFunctionPlot[1_618].Y = 72.504

	pointsOfFunctionPlot[1_619].X = 6.19
	pointsOfFunctionPlot[1_619].Y = 73.008

	pointsOfFunctionPlot[1_620].X = 6.20
	pointsOfFunctionPlot[1_620].Y = 73.516

	pointsOfFunctionPlot[1_621].X = 6.21
	pointsOfFunctionPlot[1_621].Y = 74.028

	pointsOfFunctionPlot[1_622].X = 6.22
	pointsOfFunctionPlot[1_622].Y = 74.542

	pointsOfFunctionPlot[1_623].X = 6.23
	pointsOfFunctionPlot[1_623].Y = 75.061

	pointsOfFunctionPlot[1_624].X = 6.24
	pointsOfFunctionPlot[1_624].Y = 75.583

	pointsOfFunctionPlot[1_625].X = 6.25
	pointsOfFunctionPlot[1_625].Y = 76.109

	pointsOfFunctionPlot[1_626].X = 6.26
	pointsOfFunctionPlot[1_626].Y = 76.638

	pointsOfFunctionPlot[1_627].X = 6.27
	pointsOfFunctionPlot[1_627].Y = 77.171

	pointsOfFunctionPlot[1_628].X = 6.28
	pointsOfFunctionPlot[1_628].Y = 77.708

	pointsOfFunctionPlot[1_629].X = 6.29
	pointsOfFunctionPlot[1_629].Y = 78.249

	pointsOfFunctionPlot[1_630].X = 6.30
	pointsOfFunctionPlot[1_630].Y = 78.793

	pointsOfFunctionPlot[1_631].X = 6.31
	pointsOfFunctionPlot[1_631].Y = 79.341

	pointsOfFunctionPlot[1_632].X = 6.32
	pointsOfFunctionPlot[1_632].Y = 79.893

	pointsOfFunctionPlot[1_633].X = 6.33
	pointsOfFunctionPlot[1_633].Y = 80.448

	pointsOfFunctionPlot[1_634].X = 6.34
	pointsOfFunctionPlot[1_634].Y = 81.008

	pointsOfFunctionPlot[1_635].X = 6.35
	pointsOfFunctionPlot[1_635].Y = 81.571

	pointsOfFunctionPlot[1_636].X = 6.36
	pointsOfFunctionPlot[1_636].Y = 82.139

	pointsOfFunctionPlot[1_637].X = 6.37
	pointsOfFunctionPlot[1_637].Y = 82.71

	pointsOfFunctionPlot[1_638].X = 6.38
	pointsOfFunctionPlot[1_638].Y = 83.285

	pointsOfFunctionPlot[1_639].X = 6.39
	pointsOfFunctionPlot[1_639].Y = 83.865

	pointsOfFunctionPlot[1_640].X = 6.40
	pointsOfFunctionPlot[1_640].Y = 84.448

	pointsOfFunctionPlot[1_641].X = 6.41
	pointsOfFunctionPlot[1_641].Y = 85.035

	pointsOfFunctionPlot[1_642].X = 6.42
	pointsOfFunctionPlot[1_642].Y = 85.627

	pointsOfFunctionPlot[1_643].X = 6.43
	pointsOfFunctionPlot[1_643].Y = 86.222

	pointsOfFunctionPlot[1_644].X = 6.44
	pointsOfFunctionPlot[1_644].Y = 86.822

	pointsOfFunctionPlot[1_645].X = 6.45
	pointsOfFunctionPlot[1_645].Y = 87.426

	pointsOfFunctionPlot[1_646].X = 6.46
	pointsOfFunctionPlot[1_646].Y = 88.034

	pointsOfFunctionPlot[1_647].X = 6.47
	pointsOfFunctionPlot[1_647].Y = 88.647

	pointsOfFunctionPlot[1_648].X = 6.48
	pointsOfFunctionPlot[1_648].Y = 89.263

	pointsOfFunctionPlot[1_649].X = 6.49
	pointsOfFunctionPlot[1_649].Y = 89.884

	pointsOfFunctionPlot[1_650].X = 6.50
	pointsOfFunctionPlot[1_650].Y = 90.509

	pointsOfFunctionPlot[1_651].X = 6.51
	pointsOfFunctionPlot[1_651].Y = 91.139

	pointsOfFunctionPlot[1_652].X = 6.52
	pointsOfFunctionPlot[1_652].Y = 91.773

	pointsOfFunctionPlot[1_653].X = 6.53
	pointsOfFunctionPlot[1_653].Y = 92.411

	pointsOfFunctionPlot[1_654].X = 6.54
	pointsOfFunctionPlot[1_654].Y = 93.054

	pointsOfFunctionPlot[1_655].X = 6.55
	pointsOfFunctionPlot[1_655].Y = 93.701

	pointsOfFunctionPlot[1_656].X = 6.56
	pointsOfFunctionPlot[1_656].Y = 94.353

	pointsOfFunctionPlot[1_657].X = 6.57
	pointsOfFunctionPlot[1_657].Y = 95.009

	pointsOfFunctionPlot[1_658].X = 6.58
	pointsOfFunctionPlot[1_658].Y = 95.67

	pointsOfFunctionPlot[1_659].X = 6.59
	pointsOfFunctionPlot[1_659].Y = 96.335

	pointsOfFunctionPlot[1_660].X = 6.60
	pointsOfFunctionPlot[1_660].Y = 97.005

	pointsOfFunctionPlot[1_661].X = 6.61
	pointsOfFunctionPlot[1_661].Y = 97.68

	pointsOfFunctionPlot[1_662].X = 6.62
	pointsOfFunctionPlot[1_662].Y = 98.36

	pointsOfFunctionPlot[1_663].X = 6.63
	pointsOfFunctionPlot[1_663].Y = 99.044

	pointsOfFunctionPlot[1_664].X = 6.64
	pointsOfFunctionPlot[1_664].Y = 99.733

	pointsOfFunctionPlot[1_665].X = 6.65
	pointsOfFunctionPlot[1_665].Y = 100.426

	pointsOfFunctionPlot[1_666].X = 6.66
	pointsOfFunctionPlot[1_666].Y = 101.125

	pointsOfFunctionPlot[1_667].X = 6.67
	pointsOfFunctionPlot[1_667].Y = 101.828

	pointsOfFunctionPlot[1_668].X = 6.68
	pointsOfFunctionPlot[1_668].Y = 102.536

	pointsOfFunctionPlot[1_669].X = 6.69
	pointsOfFunctionPlot[1_669].Y = 103.25

	pointsOfFunctionPlot[1_670].X = 6.70
	pointsOfFunctionPlot[1_670].Y = 103.968

	pointsOfFunctionPlot[1_671].X = 6.71
	pointsOfFunctionPlot[1_671].Y = 104.691

	pointsOfFunctionPlot[1_672].X = 6.72
	pointsOfFunctionPlot[1_672].Y = 105.419

	pointsOfFunctionPlot[1_673].X = 6.73
	pointsOfFunctionPlot[1_673].Y = 106.152

	pointsOfFunctionPlot[1_674].X = 6.74
	pointsOfFunctionPlot[1_674].Y = 106.891

	pointsOfFunctionPlot[1_675].X = 6.75
	pointsOfFunctionPlot[1_675].Y = 107.634

	pointsOfFunctionPlot[1_676].X = 6.76
	pointsOfFunctionPlot[1_676].Y = 108.383

	pointsOfFunctionPlot[1_677].X = 6.77
	pointsOfFunctionPlot[1_677].Y = 109.137

	pointsOfFunctionPlot[1_678].X = 6.78
	pointsOfFunctionPlot[1_678].Y = 109.896

	pointsOfFunctionPlot[1_679].X = 6.79
	pointsOfFunctionPlot[1_679].Y = 110.66

	pointsOfFunctionPlot[1_680].X = 6.80
	pointsOfFunctionPlot[1_680].Y = 111.43

	pointsOfFunctionPlot[1_681].X = 6.81
	pointsOfFunctionPlot[1_681].Y = 112.205

	pointsOfFunctionPlot[1_682].X = 6.82
	pointsOfFunctionPlot[1_682].Y = 112.986

	pointsOfFunctionPlot[1_683].X = 6.83
	pointsOfFunctionPlot[1_683].Y = 113.771

	pointsOfFunctionPlot[1_684].X = 6.84
	pointsOfFunctionPlot[1_684].Y = 114.563

	pointsOfFunctionPlot[1_685].X = 6.85
	pointsOfFunctionPlot[1_685].Y = 115.36

	pointsOfFunctionPlot[1_686].X = 6.86
	pointsOfFunctionPlot[1_686].Y = 116.162

	pointsOfFunctionPlot[1_687].X = 6.87
	pointsOfFunctionPlot[1_687].Y = 116.162

	pointsOfFunctionPlot[1_688].X = 6.88
	pointsOfFunctionPlot[1_688].Y = 117.784

	pointsOfFunctionPlot[1_689].X = 6.89
	pointsOfFunctionPlot[1_689].Y = 118.603

	pointsOfFunctionPlot[1_690].X = 6.90
	pointsOfFunctionPlot[1_690].Y = 119.428

	pointsOfFunctionPlot[1_691].X = 6.91
	pointsOfFunctionPlot[1_691].Y = 120.258

	pointsOfFunctionPlot[1_692].X = 6.92
	pointsOfFunctionPlot[1_692].Y = 121.095

	pointsOfFunctionPlot[1_693].X = 6.93
	pointsOfFunctionPlot[1_693].Y = 121.937

	pointsOfFunctionPlot[1_694].X = 6.94
	pointsOfFunctionPlot[1_694].Y = 122.785

	pointsOfFunctionPlot[1_695].X = 6.95
	pointsOfFunctionPlot[1_695].Y = 123.639

	pointsOfFunctionPlot[1_696].X = 6.96
	pointsOfFunctionPlot[1_696].Y = 124.499

	pointsOfFunctionPlot[1_697].X = 6.97
	pointsOfFunctionPlot[1_697].Y = 125.365

	pointsOfFunctionPlot[1_698].X = 6.98
	pointsOfFunctionPlot[1_698].Y = 126.237

	pointsOfFunctionPlot[1_699].X = 6.99
	pointsOfFunctionPlot[1_699].Y = 127.115

	pointsOfFunctionPlot[1_700].X = 7.0
	pointsOfFunctionPlot[1_700].Y = 128.0

	pointsOfFunctionPlot[1_701].X = 7.01
	pointsOfFunctionPlot[1_701].Y = 128.89

	pointsOfFunctionPlot[1_702].X = 7.02
	pointsOfFunctionPlot[1_702].Y = 129.786

	pointsOfFunctionPlot[1_703].X = 7.03
	pointsOfFunctionPlot[1_703].Y = 130.689

	pointsOfFunctionPlot[1_704].X = 7.04
	pointsOfFunctionPlot[1_704].Y = 131.598

	pointsOfFunctionPlot[1_705].X = 7.05
	pointsOfFunctionPlot[1_705].Y = 132.513

	pointsOfFunctionPlot[1_706].X = 7.06
	pointsOfFunctionPlot[1_706].Y = 133.435

	pointsOfFunctionPlot[1_707].X = 7.07
	pointsOfFunctionPlot[1_707].Y = 134.363

	pointsOfFunctionPlot[1_708].X = 7.08
	pointsOfFunctionPlot[1_708].Y = 135.298

	pointsOfFunctionPlot[1_709].X = 7.09
	pointsOfFunctionPlot[1_709].Y = 136.239

	pointsOfFunctionPlot[1_710].X = 7.10
	pointsOfFunctionPlot[1_710].Y = 137.187

	pointsOfFunctionPlot[1_711].X = 7.11
	pointsOfFunctionPlot[1_711].Y = 138.141

	pointsOfFunctionPlot[1_712].X = 7.12
	pointsOfFunctionPlot[1_712].Y = 139.102

	pointsOfFunctionPlot[1_713].X = 7.13
	pointsOfFunctionPlot[1_713].Y = 140.069

	pointsOfFunctionPlot[1_714].X = 7.14
	pointsOfFunctionPlot[1_714].Y = 141.043

	pointsOfFunctionPlot[1_715].X = 7.15
	pointsOfFunctionPlot[1_715].Y = 142.024

	pointsOfFunctionPlot[1_716].X = 7.16
	pointsOfFunctionPlot[1_716].Y = 143.012

	pointsOfFunctionPlot[1_717].X = 7.17
	pointsOfFunctionPlot[1_717].Y = 144.007

	pointsOfFunctionPlot[1_718].X = 7.18
	pointsOfFunctionPlot[1_718].Y = 145.009

	pointsOfFunctionPlot[1_719].X = 7.19
	pointsOfFunctionPlot[1_719].Y = 146.017

	pointsOfFunctionPlot[1_720].X = 7.20
	pointsOfFunctionPlot[1_720].Y = 147.033

	pointsOfFunctionPlot[1_721].X = 7.21
	pointsOfFunctionPlot[1_721].Y = 148.056

	pointsOfFunctionPlot[1_722].X = 7.22
	pointsOfFunctionPlot[1_722].Y = 149.085

	pointsOfFunctionPlot[1_723].X = 7.23
	pointsOfFunctionPlot[1_723].Y = 150.122

	pointsOfFunctionPlot[1_724].X = 7.24
	pointsOfFunctionPlot[1_724].Y = 151.167

	pointsOfFunctionPlot[1_725].X = 7.25
	pointsOfFunctionPlot[1_725].Y = 152.218

	pointsOfFunctionPlot[1_726].X = 7.26
	pointsOfFunctionPlot[1_726].Y = 153.277

	pointsOfFunctionPlot[1_727].X = 7.27
	pointsOfFunctionPlot[1_727].Y = 154.343

	pointsOfFunctionPlot[1_728].X = 7.28
	pointsOfFunctionPlot[1_728].Y = 155.416

	pointsOfFunctionPlot[1_729].X = 7.29
	pointsOfFunctionPlot[1_729].Y = 156.498

	pointsOfFunctionPlot[1_730].X = 7.30
	pointsOfFunctionPlot[1_730].Y = 157.586

	pointsOfFunctionPlot[1_731].X = 7.31
	pointsOfFunctionPlot[1_731].Y = 158.682

	pointsOfFunctionPlot[1_732].X = 7.32
	pointsOfFunctionPlot[1_732].Y = 159.786

	pointsOfFunctionPlot[1_733].X = 7.33
	pointsOfFunctionPlot[1_733].Y = 160.897

	pointsOfFunctionPlot[1_734].X = 7.34
	pointsOfFunctionPlot[1_734].Y = 162.016

	pointsOfFunctionPlot[1_735].X = 7.35
	pointsOfFunctionPlot[1_735].Y = 163.143

	pointsOfFunctionPlot[1_736].X = 7.36
	pointsOfFunctionPlot[1_736].Y = 164.278

	pointsOfFunctionPlot[1_737].X = 7.37
	pointsOfFunctionPlot[1_737].Y = 165.421

	pointsOfFunctionPlot[1_738].X = 7.38
	pointsOfFunctionPlot[1_738].Y = 166.571

	pointsOfFunctionPlot[1_739].X = 7.39
	pointsOfFunctionPlot[1_739].Y = 167.73

	pointsOfFunctionPlot[1_740].X = 7.40
	pointsOfFunctionPlot[1_740].Y = 168.897

	pointsOfFunctionPlot[1_741].X = 7.41
	pointsOfFunctionPlot[1_741].Y = 170.071

	pointsOfFunctionPlot[1_742].X = 7.42
	pointsOfFunctionPlot[1_742].Y = 171.254

	pointsOfFunctionPlot[1_743].X = 7.43
	pointsOfFunctionPlot[1_743].Y = 172.445

	pointsOfFunctionPlot[1_744].X = 7.44
	pointsOfFunctionPlot[1_744].Y = 173.645

	pointsOfFunctionPlot[1_745].X = 7.45
	pointsOfFunctionPlot[1_745].Y = 174.853

	pointsOfFunctionPlot[1_746].X = 7.46
	pointsOfFunctionPlot[1_746].Y = 176.069

	pointsOfFunctionPlot[1_747].X = 7.47
	pointsOfFunctionPlot[1_747].Y = 177.294

	pointsOfFunctionPlot[1_748].X = 7.48
	pointsOfFunctionPlot[1_748].Y = 178.527

	pointsOfFunctionPlot[1_749].X = 7.49
	pointsOfFunctionPlot[1_749].Y = 179.768

	pointsOfFunctionPlot[1_750].X = 7.50
	pointsOfFunctionPlot[1_750].Y = 181.019

	pointsOfFunctionPlot[1_751].X = 7.51
	pointsOfFunctionPlot[1_751].Y = 182.278

	pointsOfFunctionPlot[1_752].X = 7.52
	pointsOfFunctionPlot[1_752].Y = 183.546

	pointsOfFunctionPlot[1_753].X = 7.53
	pointsOfFunctionPlot[1_753].Y = 184.822

	pointsOfFunctionPlot[1_754].X = 7.54
	pointsOfFunctionPlot[1_754].Y = 186.108

	pointsOfFunctionPlot[1_755].X = 7.55
	pointsOfFunctionPlot[1_755].Y = 187.403

	pointsOfFunctionPlot[1_756].X = 7.56
	pointsOfFunctionPlot[1_756].Y = 188.706

	pointsOfFunctionPlot[1_757].X = 7.57
	pointsOfFunctionPlot[1_757].Y = 190.019

	pointsOfFunctionPlot[1_758].X = 7.58
	pointsOfFunctionPlot[1_758].Y = 191.34

	pointsOfFunctionPlot[1_759].X = 7.59
	pointsOfFunctionPlot[1_759].Y = 192.671

	pointsOfFunctionPlot[1_760].X = 7.60
	pointsOfFunctionPlot[1_760].Y = 194.011

	pointsOfFunctionPlot[1_761].X = 7.61
	pointsOfFunctionPlot[1_761].Y = 195.361

	pointsOfFunctionPlot[1_762].X = 7.62
	pointsOfFunctionPlot[1_762].Y = 196.72

	pointsOfFunctionPlot[1_763].X = 7.63
	pointsOfFunctionPlot[1_763].Y = 198.088

	pointsOfFunctionPlot[1_764].X = 7.64
	pointsOfFunctionPlot[1_764].Y = 199.466

	pointsOfFunctionPlot[1_765].X = 7.65
	pointsOfFunctionPlot[1_765].Y = 200.853

	pointsOfFunctionPlot[1_766].X = 7.66
	pointsOfFunctionPlot[1_766].Y = 202.25

	pointsOfFunctionPlot[1_767].X = 7.67
	pointsOfFunctionPlot[1_767].Y = 203.657

	pointsOfFunctionPlot[1_768].X = 7.68
	pointsOfFunctionPlot[1_768].Y = 205.073

	pointsOfFunctionPlot[1_769].X = 7.69
	pointsOfFunctionPlot[1_769].Y = 206.5

	pointsOfFunctionPlot[1_770].X = 7.70
	pointsOfFunctionPlot[1_770].Y = 207.936

	pointsOfFunctionPlot[1_771].X = 7.71
	pointsOfFunctionPlot[1_771].Y = 209.382

	pointsOfFunctionPlot[1_772].X = 7.72
	pointsOfFunctionPlot[1_772].Y = 210.839

	pointsOfFunctionPlot[1_773].X = 7.73
	pointsOfFunctionPlot[1_773].Y = 212.305

	pointsOfFunctionPlot[1_774].X = 7.74
	pointsOfFunctionPlot[1_774].Y = 213.782

	pointsOfFunctionPlot[1_775].X = 7.75
	pointsOfFunctionPlot[1_775].Y = 215.269

	pointsOfFunctionPlot[1_776].X = 7.76
	pointsOfFunctionPlot[1_776].Y = 216.766

	pointsOfFunctionPlot[1_777].X = 7.77
	pointsOfFunctionPlot[1_777].Y = 218.274

	pointsOfFunctionPlot[1_778].X = 7.78
	pointsOfFunctionPlot[1_778].Y = 219.792

	pointsOfFunctionPlot[1_779].X = 7.79
	pointsOfFunctionPlot[1_779].Y = 221.321

	pointsOfFunctionPlot[1_780].X = 7.80
	pointsOfFunctionPlot[1_780].Y = 222.86

	pointsOfFunctionPlot[1_781].X = 7.81
	pointsOfFunctionPlot[1_781].Y = 224.411

	pointsOfFunctionPlot[1_782].X = 7.82
	pointsOfFunctionPlot[1_782].Y = 225.972

	pointsOfFunctionPlot[1_783].X = 7.83
	pointsOfFunctionPlot[1_783].Y = 227.543

	pointsOfFunctionPlot[1_784].X = 7.84
	pointsOfFunctionPlot[1_784].Y = 229.126

	pointsOfFunctionPlot[1_785].X = 7.85
	pointsOfFunctionPlot[1_785].Y = 230.72

	pointsOfFunctionPlot[1_786].X = 7.86
	pointsOfFunctionPlot[1_786].Y = 232.324

	pointsOfFunctionPlot[1_787].X = 7.87
	pointsOfFunctionPlot[1_787].Y = 233.94

	pointsOfFunctionPlot[1_788].X = 7.88
	pointsOfFunctionPlot[1_788].Y = 235.568

	pointsOfFunctionPlot[1_789].X = 7.89
	pointsOfFunctionPlot[1_789].Y = 237.206

	pointsOfFunctionPlot[1_790].X = 7.90
	pointsOfFunctionPlot[1_790].Y = 238.856

	pointsOfFunctionPlot[1_791].X = 7.91
	pointsOfFunctionPlot[1_791].Y = 240.517

	pointsOfFunctionPlot[1_792].X = 7.92
	pointsOfFunctionPlot[1_792].Y = 242.19

	pointsOfFunctionPlot[1_793].X = 7.93
	pointsOfFunctionPlot[1_793].Y = 243.875

	pointsOfFunctionPlot[1_794].X = 7.94
	pointsOfFunctionPlot[1_794].Y = 245.571

	pointsOfFunctionPlot[1_795].X = 7.95
	pointsOfFunctionPlot[1_795].Y = 247.279

	pointsOfFunctionPlot[1_796].X = 7.96
	pointsOfFunctionPlot[1_796].Y = 248.999

	pointsOfFunctionPlot[1_797].X = 7.97
	pointsOfFunctionPlot[1_797].Y = 250.731

	pointsOfFunctionPlot[1_798].X = 7.98
	pointsOfFunctionPlot[1_798].Y = 252.475

	pointsOfFunctionPlot[1_799].X = 7.99
	pointsOfFunctionPlot[1_799].Y = 254.231

	pointsOfFunctionPlot[1_800].X = 8.0
	pointsOfFunctionPlot[1_800].Y = 256.0

	pointsOfFunctionPlot[1_801].X = 8.01
	pointsOfFunctionPlot[1_801].Y = 257.78

	pointsOfFunctionPlot[1_802].X = 8.02
	pointsOfFunctionPlot[1_802].Y = 259.78

	pointsOfFunctionPlot[1_803].X = 8.03
	pointsOfFunctionPlot[1_803].Y = 261.379

	pointsOfFunctionPlot[1_804].X = 8.04
	pointsOfFunctionPlot[1_804].Y = 263.197

	pointsOfFunctionPlot[1_805].X = 8.05
	pointsOfFunctionPlot[1_805].Y = 265.027

	pointsOfFunctionPlot[1_806].X = 8.06
	pointsOfFunctionPlot[1_806].Y = 266.871

	pointsOfFunctionPlot[1_807].X = 8.07
	pointsOfFunctionPlot[1_807].Y = 268.727

	pointsOfFunctionPlot[1_808].X = 8.08
	pointsOfFunctionPlot[1_808].Y = 270.596

	pointsOfFunctionPlot[1_809].X = 8.09
	pointsOfFunctionPlot[1_809].Y = 272.478

	pointsOfFunctionPlot[1_810].X = 8.10
	pointsOfFunctionPlot[1_810].Y = 274.374

	pointsOfFunctionPlot[1_811].X = 8.11
	pointsOfFunctionPlot[1_811].Y = 276.282

	pointsOfFunctionPlot[1_812].X = 8.12
	pointsOfFunctionPlot[1_812].Y = 278.204

	pointsOfFunctionPlot[1_813].X = 8.13
	pointsOfFunctionPlot[1_813].Y = 280.139

	pointsOfFunctionPlot[1_814].X = 8.14
	pointsOfFunctionPlot[1_814].Y = 282.087

	pointsOfFunctionPlot[1_815].X = 8.15
	pointsOfFunctionPlot[1_815].Y = 284.049

	pointsOfFunctionPlot[1_816].X = 8.16
	pointsOfFunctionPlot[1_816].Y = 286.025

	pointsOfFunctionPlot[1_817].X = 8.17
	pointsOfFunctionPlot[1_817].Y = 288.015

	pointsOfFunctionPlot[1_818].X = 8.18
	pointsOfFunctionPlot[1_818].Y = 290.018

	pointsOfFunctionPlot[1_819].X = 8.19
	pointsOfFunctionPlot[1_819].Y = 292.035

	pointsOfFunctionPlot[1_820].X = 8.20
	pointsOfFunctionPlot[1_820].Y = 294.066

	pointsOfFunctionPlot[1_821].X = 8.21
	pointsOfFunctionPlot[1_821].Y = 296.112

	pointsOfFunctionPlot[1_822].X = 8.22
	pointsOfFunctionPlot[1_822].Y = 298.171

	pointsOfFunctionPlot[1_823].X = 8.23
	pointsOfFunctionPlot[1_823].Y = 300.245

	pointsOfFunctionPlot[1_824].X = 8.24
	pointsOfFunctionPlot[1_824].Y = 302.334

	pointsOfFunctionPlot[1_825].X = 8.25
	pointsOfFunctionPlot[1_825].Y = 304.437

	pointsOfFunctionPlot[1_826].X = 8.26
	pointsOfFunctionPlot[1_826].Y = 306.554

	pointsOfFunctionPlot[1_827].X = 8.27
	pointsOfFunctionPlot[1_827].Y = 308.686

	pointsOfFunctionPlot[1_828].X = 8.28
	pointsOfFunctionPlot[1_828].Y = 310.833

	pointsOfFunctionPlot[1_829].X = 8.29
	pointsOfFunctionPlot[1_829].Y = 312.995

	pointsOfFunctionPlot[1_830].X = 8.30
	pointsOfFunctionPlot[1_830].Y = 315.173

	pointsOfFunctionPlot[1_831].X = 8.31
	pointsOfFunctionPlot[1_831].Y = 317.365

	pointsOfFunctionPlot[1_832].X = 8.32
	pointsOfFunctionPlot[1_832].Y = 319.572

	pointsOfFunctionPlot[1_833].X = 8.33
	pointsOfFunctionPlot[1_833].Y = 321.795

	pointsOfFunctionPlot[1_834].X = 8.34
	pointsOfFunctionPlot[1_834].Y = 324.033

	pointsOfFunctionPlot[1_835].X = 8.35
	pointsOfFunctionPlot[1_835].Y = 326.287

	pointsOfFunctionPlot[1_836].X = 8.36
	pointsOfFunctionPlot[1_836].Y = 328.557

	pointsOfFunctionPlot[1_837].X = 8.37
	pointsOfFunctionPlot[1_837].Y = 330.842

	pointsOfFunctionPlot[1_838].X = 8.38
	pointsOfFunctionPlot[1_838].Y = 333.143

	pointsOfFunctionPlot[1_839].X = 8.39
	pointsOfFunctionPlot[1_839].Y = 335.46

	pointsOfFunctionPlot[1_840].X = 8.40
	pointsOfFunctionPlot[1_840].Y = 337.794

	pointsOfFunctionPlot[1_841].X = 8.41
	pointsOfFunctionPlot[1_841].Y = 340.143

	pointsOfFunctionPlot[1_842].X = 8.42
	pointsOfFunctionPlot[1_842].Y = 342.509

	pointsOfFunctionPlot[1_843].X = 8.43
	pointsOfFunctionPlot[1_843].Y = 344.891

	pointsOfFunctionPlot[1_844].X = 8.44
	pointsOfFunctionPlot[1_844].Y = 347.29

	pointsOfFunctionPlot[1_845].X = 8.45
	pointsOfFunctionPlot[1_845].Y = 349.706

	pointsOfFunctionPlot[1_846].X = 8.46
	pointsOfFunctionPlot[1_846].Y = 352.138

	pointsOfFunctionPlot[1_847].X = 8.47
	pointsOfFunctionPlot[1_847].Y = 354.588

	pointsOfFunctionPlot[1_848].X = 8.48
	pointsOfFunctionPlot[1_848].Y = 357.054

	pointsOfFunctionPlot[1_849].X = 8.49
	pointsOfFunctionPlot[1_849].Y = 359.537

	pointsOfFunctionPlot[1_850].X = 8.50
	pointsOfFunctionPlot[1_850].Y = 362.038

	pointsOfFunctionPlot[1_851].X = 8.51
	pointsOfFunctionPlot[1_851].Y = 364.556

	pointsOfFunctionPlot[1_852].X = 8.52
	pointsOfFunctionPlot[1_852].Y = 367.092

	pointsOfFunctionPlot[1_853].X = 8.53
	pointsOfFunctionPlot[1_853].Y = 369.645

	pointsOfFunctionPlot[1_854].X = 8.54
	pointsOfFunctionPlot[1_854].Y = 372.217

	pointsOfFunctionPlot[1_855].X = 8.55
	pointsOfFunctionPlot[1_855].Y = 374.805

	pointsOfFunctionPlot[1_856].X = 8.56
	pointsOfFunctionPlot[1_856].Y = 377.412

	pointsOfFunctionPlot[1_857].X = 8.57
	pointsOfFunctionPlot[1_857].Y = 380.038

	pointsOfFunctionPlot[1_858].X = 8.58
	pointsOfFunctionPlot[1_858].Y = 382.681

	pointsOfFunctionPlot[1_859].X = 8.59
	pointsOfFunctionPlot[1_859].Y = 385.343

	pointsOfFunctionPlot[1_860].X = 8.60
	pointsOfFunctionPlot[1_860].Y = 388.023

	pointsOfFunctionPlot[1_861].X = 8.61
	pointsOfFunctionPlot[1_861].Y = 390.722

	pointsOfFunctionPlot[1_862].X = 8.62
	pointsOfFunctionPlot[1_862].Y = 393.44

	pointsOfFunctionPlot[1_863].X = 8.63
	pointsOfFunctionPlot[1_863].Y = 396.176

	pointsOfFunctionPlot[1_864].X = 8.64
	pointsOfFunctionPlot[1_864].Y = 398.932

	pointsOfFunctionPlot[1_865].X = 8.65
	pointsOfFunctionPlot[1_865].Y = 401.707

	pointsOfFunctionPlot[1_866].X = 8.66
	pointsOfFunctionPlot[1_866].Y = 404.501

	pointsOfFunctionPlot[1_867].X = 8.67
	pointsOfFunctionPlot[1_867].Y = 407.314

	pointsOfFunctionPlot[1_868].X = 8.68
	pointsOfFunctionPlot[1_868].Y = 410.147

	pointsOfFunctionPlot[1_869].X = 8.69
	pointsOfFunctionPlot[1_869].Y = 413.0

	pointsOfFunctionPlot[1_870].X = 8.70
	pointsOfFunctionPlot[1_870].Y = 415.873

	pointsOfFunctionPlot[1_871].X = 8.71
	pointsOfFunctionPlot[1_871].Y = 418.765

	pointsOfFunctionPlot[1_872].X = 8.72
	pointsOfFunctionPlot[1_872].Y = 421.678

	pointsOfFunctionPlot[1_873].X = 8.73
	pointsOfFunctionPlot[1_873].Y = 424.611

	pointsOfFunctionPlot[1_874].X = 8.74
	pointsOfFunctionPlot[1_874].Y = 427.565

	pointsOfFunctionPlot[1_875].X = 8.75
	pointsOfFunctionPlot[1_875].Y = 430.539

	pointsOfFunctionPlot[1_876].X = 8.76
	pointsOfFunctionPlot[1_876].Y = 433.533

	pointsOfFunctionPlot[1_877].X = 8.77
	pointsOfFunctionPlot[1_877].Y = 436.549

	pointsOfFunctionPlot[1_878].X = 8.78
	pointsOfFunctionPlot[1_878].Y = 439.585

	pointsOfFunctionPlot[1_879].X = 8.79
	pointsOfFunctionPlot[1_879].Y = 442.643

	pointsOfFunctionPlot[1_880].X = 8.80
	pointsOfFunctionPlot[1_880].Y = 445.721

	pointsOfFunctionPlot[1_881].X = 8.81
	pointsOfFunctionPlot[1_881].Y = 448.822

	pointsOfFunctionPlot[1_882].X = 8.82
	pointsOfFunctionPlot[1_882].Y = 451.943

	pointsOfFunctionPlot[1_883].X = 8.83
	pointsOfFunctionPlot[1_883].Y = 455.087

	pointsOfFunctionPlot[1_884].X = 8.84
	pointsOfFunctionPlot[1_884].Y = 458.252

	pointsOfFunctionPlot[1_885].X = 8.85
	pointsOfFunctionPlot[1_885].Y = 461.44

	pointsOfFunctionPlot[1_886].X = 8.86
	pointsOfFunctionPlot[1_886].Y = 464.649

	pointsOfFunctionPlot[1_887].X = 8.87
	pointsOfFunctionPlot[1_887].Y = 467.881

	pointsOfFunctionPlot[1_888].X = 8.88
	pointsOfFunctionPlot[1_888].Y = 471.136

	pointsOfFunctionPlot[1_889].X = 8.89
	pointsOfFunctionPlot[1_889].Y = 474.413

	pointsOfFunctionPlot[1_890].X = 8.90
	pointsOfFunctionPlot[1_890].Y = 477.712

	pointsOfFunctionPlot[1_891].X = 8.91
	pointsOfFunctionPlot[1_891].Y = 481.035

	pointsOfFunctionPlot[1_892].X = 8.92
	pointsOfFunctionPlot[1_892].Y = 484.381

	pointsOfFunctionPlot[1_893].X = 8.93
	pointsOfFunctionPlot[1_893].Y = 487.75

	pointsOfFunctionPlot[1_894].X = 8.94
	pointsOfFunctionPlot[1_894].Y = 491.143

	pointsOfFunctionPlot[1_895].X = 8.95
	pointsOfFunctionPlot[1_895].Y = 494.559

	pointsOfFunctionPlot[1_896].X = 8.96
	pointsOfFunctionPlot[1_896].Y = 497.999

	pointsOfFunctionPlot[1_897].X = 8.97
	pointsOfFunctionPlot[1_897].Y = 501.463

	pointsOfFunctionPlot[1_898].X = 8.98
	pointsOfFunctionPlot[1_898].Y = 504.951

	pointsOfFunctionPlot[1_899].X = 8.99
	pointsOfFunctionPlot[1_899].Y = 508.463

	pointsOfFunctionPlot[1_900].X = 9.0
	pointsOfFunctionPlot[1_900].Y = 512.0

	pointsOfFunctionPlot[1_901].X = 9.01
	pointsOfFunctionPlot[1_901].Y = 515.561

	pointsOfFunctionPlot[1_902].X = 9.02
	pointsOfFunctionPlot[1_902].Y = 519.147

	pointsOfFunctionPlot[1_903].X = 9.03
	pointsOfFunctionPlot[1_903].Y = 522.758

	pointsOfFunctionPlot[1_904].X = 9.04
	pointsOfFunctionPlot[1_904].Y = 526.394

	pointsOfFunctionPlot[1_905].X = 9.05
	pointsOfFunctionPlot[1_905].Y = 530.055

	pointsOfFunctionPlot[1_906].X = 9.06
	pointsOfFunctionPlot[1_906].Y = 533.742

	pointsOfFunctionPlot[1_907].X = 9.07
	pointsOfFunctionPlot[1_907].Y = 537.454

	pointsOfFunctionPlot[1_908].X = 9.08
	pointsOfFunctionPlot[1_908].Y = 541.193

	pointsOfFunctionPlot[1_909].X = 9.09
	pointsOfFunctionPlot[1_909].Y = 544.957

	pointsOfFunctionPlot[1_910].X = 9.10
	pointsOfFunctionPlot[1_910].Y = 548.748

	pointsOfFunctionPlot[1_911].X = 9.11
	pointsOfFunctionPlot[1_911].Y = 552.564

	pointsOfFunctionPlot[1_912].X = 9.12
	pointsOfFunctionPlot[1_912].Y = 556.408

	pointsOfFunctionPlot[1_913].X = 9.13
	pointsOfFunctionPlot[1_913].Y = 560.278

	pointsOfFunctionPlot[1_914].X = 9.14
	pointsOfFunctionPlot[1_914].Y = 564.175

	pointsOfFunctionPlot[1_915].X = 9.15
	pointsOfFunctionPlot[1_915].Y = 568.099

	pointsOfFunctionPlot[1_916].X = 9.16
	pointsOfFunctionPlot[1_916].Y = 572.051

	pointsOfFunctionPlot[1_917].X = 9.17
	pointsOfFunctionPlot[1_917].Y = 576.029

	pointsOfFunctionPlot[1_918].X = 9.18
	pointsOfFunctionPlot[1_918].Y = 580.036

	pointsOfFunctionPlot[1_919].X = 9.19
	pointsOfFunctionPlot[1_919].Y = 584.071

	pointsOfFunctionPlot[1_920].X = 9.20
	pointsOfFunctionPlot[1_920].Y = 588.133

	pointsOfFunctionPlot[1_921].X = 9.21
	pointsOfFunctionPlot[1_921].Y = 592.224

	pointsOfFunctionPlot[1_922].X = 9.22
	pointsOfFunctionPlot[1_922].Y = 596.343

	pointsOfFunctionPlot[1_923].X = 9.23
	pointsOfFunctionPlot[1_923].Y = 600.491

	pointsOfFunctionPlot[1_924].X = 9.24
	pointsOfFunctionPlot[1_924].Y = 604.668

	pointsOfFunctionPlot[1_925].X = 9.25
	pointsOfFunctionPlot[1_925].Y = 608.874

	pointsOfFunctionPlot[1_926].X = 9.26
	pointsOfFunctionPlot[1_926].Y = 613.109

	pointsOfFunctionPlot[1_927].X = 9.27
	pointsOfFunctionPlot[1_927].Y = 617.373

	pointsOfFunctionPlot[1_928].X = 9.28
	pointsOfFunctionPlot[1_928].Y = 621.667

	pointsOfFunctionPlot[1_929].X = 9.29
	pointsOfFunctionPlot[1_929].Y = 625.991

	pointsOfFunctionPlot[1_930].X = 9.30
	pointsOfFunctionPlot[1_930].Y = 630.345

	pointsOfFunctionPlot[1_931].X = 9.31
	pointsOfFunctionPlot[1_931].Y = 634.73

	pointsOfFunctionPlot[1_932].X = 9.32
	pointsOfFunctionPlot[1_932].Y = 639.145

	pointsOfFunctionPlot[1_933].X = 9.33
	pointsOfFunctionPlot[1_933].Y = 643.59

	pointsOfFunctionPlot[1_934].X = 9.34
	pointsOfFunctionPlot[1_934].Y = 648.067

	pointsOfFunctionPlot[1_935].X = 9.35
	pointsOfFunctionPlot[1_935].Y = 652.575

	pointsOfFunctionPlot[1_936].X = 9.36
	pointsOfFunctionPlot[1_936].Y = 657.114

	pointsOfFunctionPlot[1_937].X = 9.37
	pointsOfFunctionPlot[1_937].Y = 661.684

	pointsOfFunctionPlot[1_938].X = 9.38
	pointsOfFunctionPlot[1_938].Y = 666.287

	pointsOfFunctionPlot[1_939].X = 9.39
	pointsOfFunctionPlot[1_939].Y = 670.921

	pointsOfFunctionPlot[1_940].X = 9.40
	pointsOfFunctionPlot[1_940].Y = 675.588

	pointsOfFunctionPlot[1_941].X = 9.41
	pointsOfFunctionPlot[1_941].Y = 680.287

	pointsOfFunctionPlot[1_942].X = 9.42
	pointsOfFunctionPlot[1_942].Y = 685.018

	pointsOfFunctionPlot[1_943].X = 9.43
	pointsOfFunctionPlot[1_943].Y = 689.783

	pointsOfFunctionPlot[1_944].X = 9.44
	pointsOfFunctionPlot[1_944].Y = 694.581

	pointsOfFunctionPlot[1_945].X = 9.45
	pointsOfFunctionPlot[1_945].Y = 699.412

	pointsOfFunctionPlot[1_946].X = 9.46
	pointsOfFunctionPlot[1_946].Y = 704.277

	pointsOfFunctionPlot[1_947].X = 9.47
	pointsOfFunctionPlot[1_947].Y = 709.176

	pointsOfFunctionPlot[1_948].X = 9.48
	pointsOfFunctionPlot[1_948].Y = 714.108

	pointsOfFunctionPlot[1_949].X = 9.49
	pointsOfFunctionPlot[1_949].Y = 719.075

	pointsOfFunctionPlot[1_950].X = 9.50
	pointsOfFunctionPlot[1_950].Y = 724.077

	pointsOfFunctionPlot[1_951].X = 9.51
	pointsOfFunctionPlot[1_951].Y = 729.113

	pointsOfFunctionPlot[1_952].X = 9.52
	pointsOfFunctionPlot[1_952].Y = 734.185

	pointsOfFunctionPlot[1_953].X = 9.53
	pointsOfFunctionPlot[1_953].Y = 739.291

	pointsOfFunctionPlot[1_954].X = 9.54
	pointsOfFunctionPlot[1_954].Y = 744.433

	pointsOfFunctionPlot[1_955].X = 9.55
	pointsOfFunctionPlot[1_955].Y = 749.611

	pointsOfFunctionPlot[1_956].X = 9.56
	pointsOfFunctionPlot[1_956].Y = 754.825

	pointsOfFunctionPlot[1_957].X = 9.57
	pointsOfFunctionPlot[1_957].Y = 760.076

	pointsOfFunctionPlot[1_958].X = 9.58
	pointsOfFunctionPlot[1_958].Y = 765.362

	pointsOfFunctionPlot[1_959].X = 9.59
	pointsOfFunctionPlot[1_959].Y = 770.686

	pointsOfFunctionPlot[1_960].X = 9.60
	pointsOfFunctionPlot[1_960].Y = 776.046

	pointsOfFunctionPlot[1_961].X = 9.61
	pointsOfFunctionPlot[1_961].Y = 781.444

	pointsOfFunctionPlot[1_962].X = 9.62
	pointsOfFunctionPlot[1_962].Y = 786.88

	pointsOfFunctionPlot[1_963].X = 9.63
	pointsOfFunctionPlot[1_963].Y = 792.353

	pointsOfFunctionPlot[1_964].X = 9.64
	pointsOfFunctionPlot[1_964].Y = 797.864

	pointsOfFunctionPlot[1_965].X = 9.65
	pointsOfFunctionPlot[1_965].Y = 803.414

	pointsOfFunctionPlot[1_966].X = 9.66
	pointsOfFunctionPlot[1_966].Y = 809.002

	pointsOfFunctionPlot[1_967].X = 9.67
	pointsOfFunctionPlot[1_967].Y = 814.629

	pointsOfFunctionPlot[1_968].X = 9.68
	pointsOfFunctionPlot[1_968].Y = 820.295

	pointsOfFunctionPlot[1_969].X = 9.69
	pointsOfFunctionPlot[1_969].Y = 826.001

	pointsOfFunctionPlot[1_970].X = 9.70
	pointsOfFunctionPlot[1_970].Y = 831.746

	pointsOfFunctionPlot[1_971].X = 9.71
	pointsOfFunctionPlot[1_971].Y = 837.531

	pointsOfFunctionPlot[1_972].X = 9.72
	pointsOfFunctionPlot[1_972].Y = 843.357

	pointsOfFunctionPlot[1_973].X = 9.73
	pointsOfFunctionPlot[1_973].Y = 849.223

	pointsOfFunctionPlot[1_974].X = 9.74
	pointsOfFunctionPlot[1_974].Y = 855.13

	pointsOfFunctionPlot[1_975].X = 9.75
	pointsOfFunctionPlot[1_975].Y = 861.077

	pointsOfFunctionPlot[1_976].X = 9.76
	pointsOfFunctionPlot[1_976].Y = 867.067

	pointsOfFunctionPlot[1_977].X = 9.77
	pointsOfFunctionPlot[1_977].Y = 873.098

	pointsOfFunctionPlot[1_978].X = 9.78
	pointsOfFunctionPlot[1_978].Y = 879.171

	pointsOfFunctionPlot[1_979].X = 9.79
	pointsOfFunctionPlot[1_979].Y = 885.286

	pointsOfFunctionPlot[1_980].X = 9.80
	pointsOfFunctionPlot[1_980].Y = 891.443

	pointsOfFunctionPlot[1_981].X = 9.81
	pointsOfFunctionPlot[1_981].Y = 897.644

	pointsOfFunctionPlot[1_982].X = 9.82
	pointsOfFunctionPlot[1_982].Y = 903.887

	pointsOfFunctionPlot[1_983].X = 9.83
	pointsOfFunctionPlot[1_983].Y = 910.174

	pointsOfFunctionPlot[1_984].X = 9.84
	pointsOfFunctionPlot[1_984].Y = 916.505

	pointsOfFunctionPlot[1_985].X = 9.85
	pointsOfFunctionPlot[1_985].Y = 922.88

	pointsOfFunctionPlot[1_986].X = 9.86
	pointsOfFunctionPlot[1_986].Y = 929.299

	pointsOfFunctionPlot[1_987].X = 9.87
	pointsOfFunctionPlot[1_987].Y = 935.763

	pointsOfFunctionPlot[1_988].X = 9.88
	pointsOfFunctionPlot[1_988].Y = 942.272

	pointsOfFunctionPlot[1_989].X = 9.89
	pointsOfFunctionPlot[1_989].Y = 948.826

	pointsOfFunctionPlot[1_990].X = 9.90
	pointsOfFunctionPlot[1_990].Y = 955.425

	pointsOfFunctionPlot[1_991].X = 9.91
	pointsOfFunctionPlot[1_991].Y = 962.071

	pointsOfFunctionPlot[1_992].X = 9.92
	pointsOfFunctionPlot[1_992].Y = 968.763

	pointsOfFunctionPlot[1_993].X = 9.93
	pointsOfFunctionPlot[1_993].Y = 975.501

	pointsOfFunctionPlot[1_994].X = 9.94
	pointsOfFunctionPlot[1_994].Y = 982.286

	pointsOfFunctionPlot[1_995].X = 9.95
	pointsOfFunctionPlot[1_995].Y = 989.118

	pointsOfFunctionPlot[1_996].X = 9.96
	pointsOfFunctionPlot[1_996].Y = 995.998

	pointsOfFunctionPlot[1_997].X = 9.97
	pointsOfFunctionPlot[1_997].Y = 1_002.926

	pointsOfFunctionPlot[1_998].X = 9.98
	pointsOfFunctionPlot[1_998].Y = 1_009.902

	pointsOfFunctionPlot[1_999].X = 9.99
	pointsOfFunctionPlot[1_999].Y = 1_016.926

	pointsOfFunctionPlot[2_000].X = 10.0
	pointsOfFunctionPlot[2_000].Y = 1_024.0










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function 2^x"

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
		"Power-of-2-function-plot-01.png"); err != nil {

		panic(err)
	}
}
