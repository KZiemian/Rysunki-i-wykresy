package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function tan(x).

	pointsOfFunctionPlot := make(plotter.XYs, 313)

	pointsOfFunctionPlot[0].X = -1.56
	pointsOfFunctionPlot[0].Y = -92.62

	pointsOfFunctionPlot[1].X = -1.55
	pointsOfFunctionPlot[1].Y = -48.078

	pointsOfFunctionPlot[2].X = -1.54
	pointsOfFunctionPlot[2].Y = -32.461

	pointsOfFunctionPlot[3].X = -1.53
	pointsOfFunctionPlot[3].Y = -24.498

	pointsOfFunctionPlot[4].X = -1.52
	pointsOfFunctionPlot[4].Y = -19.669

	pointsOfFunctionPlot[5].X = -1.51
	pointsOfFunctionPlot[5].Y = -16.428

	pointsOfFunctionPlot[6].X = -1.50
	pointsOfFunctionPlot[6].Y = -14.101

	pointsOfFunctionPlot[7].X = -1.49
	pointsOfFunctionPlot[7].Y = -12.349

	pointsOfFunctionPlot[8].X = -1.48
	pointsOfFunctionPlot[8].Y = -10.983

	pointsOfFunctionPlot[9].X = -1.47
	pointsOfFunctionPlot[9].Y = -9.887

	pointsOfFunctionPlot[10].X = -1.46
	pointsOfFunctionPlot[10].Y = -8.988

	pointsOfFunctionPlot[11].X = -1.45
	pointsOfFunctionPlot[11].Y = -8.238

	pointsOfFunctionPlot[12].X = -1.44
	pointsOfFunctionPlot[12].Y = -7.601

	pointsOfFunctionPlot[13].X = -1.43
	pointsOfFunctionPlot[13].Y = -7.055

	pointsOfFunctionPlot[14].X = -1.42
	pointsOfFunctionPlot[14].Y = -6.581

	pointsOfFunctionPlot[15].X = -1.41
	pointsOfFunctionPlot[15].Y = -6.165

	pointsOfFunctionPlot[16].X = -1.40
	pointsOfFunctionPlot[16].Y = -5.797

	pointsOfFunctionPlot[17].X = -1.39
	pointsOfFunctionPlot[17].Y = -5.47

	pointsOfFunctionPlot[18].X = -1.38
	pointsOfFunctionPlot[18].Y = -5.177

	pointsOfFunctionPlot[19].X = -1.37
	pointsOfFunctionPlot[19].Y = -4.913

	pointsOfFunctionPlot[20].X = -1.36
	pointsOfFunctionPlot[20].Y = -4.673

	pointsOfFunctionPlot[21].X = -1.35
	pointsOfFunctionPlot[21].Y = -4.455

	pointsOfFunctionPlot[22].X = -1.34
	pointsOfFunctionPlot[22].Y = -4.255

	pointsOfFunctionPlot[23].X = -1.33
	pointsOfFunctionPlot[23].Y = -4.072

	pointsOfFunctionPlot[24].X = -1.32
	pointsOfFunctionPlot[24].Y = -3.903

	pointsOfFunctionPlot[25].X = -1.31
	pointsOfFunctionPlot[25].Y = -3.747

	pointsOfFunctionPlot[26].X = -1.30
	pointsOfFunctionPlot[26].Y = -3.602

	pointsOfFunctionPlot[27].X = -1.29
	pointsOfFunctionPlot[27].Y = -3.467

	pointsOfFunctionPlot[28].X = -1.28
	pointsOfFunctionPlot[28].Y = -3.341

	pointsOfFunctionPlot[29].X = -1.27
	pointsOfFunctionPlot[29].Y = -3.223

	pointsOfFunctionPlot[30].X = -1.26
	pointsOfFunctionPlot[30].Y = -3.113

	pointsOfFunctionPlot[31].X = -1.25
	pointsOfFunctionPlot[31].Y = -3.009

	pointsOfFunctionPlot[32].X = -1.24
	pointsOfFunctionPlot[32].Y = -2.911

	pointsOfFunctionPlot[33].X = -1.23
	pointsOfFunctionPlot[33].Y = -2.819

	pointsOfFunctionPlot[34].X = -1.22
	pointsOfFunctionPlot[34].Y = -2.732

	pointsOfFunctionPlot[35].X = -1.21
	pointsOfFunctionPlot[35].Y = -2.65

	pointsOfFunctionPlot[36].X = -1.20
	pointsOfFunctionPlot[36].Y = -2.572

	pointsOfFunctionPlot[37].X = -1.19
	pointsOfFunctionPlot[37].Y = -2.497

	pointsOfFunctionPlot[38].X = -1.18
	pointsOfFunctionPlot[38].Y = -2.427

	pointsOfFunctionPlot[39].X = -1.17
	pointsOfFunctionPlot[39].Y = -2.36

	pointsOfFunctionPlot[40].X = -1.16
	pointsOfFunctionPlot[40].Y = -2.295

	pointsOfFunctionPlot[41].X = -1.15
	pointsOfFunctionPlot[41].Y = -2.234

	pointsOfFunctionPlot[42].X = -1.14
	pointsOfFunctionPlot[42].Y = -2.175

	pointsOfFunctionPlot[43].X = -1.13
	pointsOfFunctionPlot[43].Y = -2.119

	pointsOfFunctionPlot[44].X = -1.12
	pointsOfFunctionPlot[44].Y = -2.066

	pointsOfFunctionPlot[45].X = -1.11
	pointsOfFunctionPlot[45].Y = -2.014

	pointsOfFunctionPlot[46].X = -1.10
	pointsOfFunctionPlot[46].Y = -1.964

	pointsOfFunctionPlot[47].X = -1.09
	pointsOfFunctionPlot[47].Y = -1.917

	pointsOfFunctionPlot[48].X = -1.08
	pointsOfFunctionPlot[48].Y = -1.871

	pointsOfFunctionPlot[49].X = -1.07
	pointsOfFunctionPlot[49].Y = -1.827

	pointsOfFunctionPlot[50].X = -1.06
	pointsOfFunctionPlot[50].Y = -1.784

	pointsOfFunctionPlot[51].X = -1.05
	pointsOfFunctionPlot[51].Y = -1.743

	pointsOfFunctionPlot[52].X = -1.04
	pointsOfFunctionPlot[52].Y = -1.703

	pointsOfFunctionPlot[53].X = -1.03
	pointsOfFunctionPlot[53].Y = -1.665

	pointsOfFunctionPlot[54].X = -1.02
	pointsOfFunctionPlot[54].Y = -1.628

	pointsOfFunctionPlot[55].X = -1.01
	pointsOfFunctionPlot[55].Y = -1.592

	pointsOfFunctionPlot[56].X = -1.0
	pointsOfFunctionPlot[56].Y = -1.557

	pointsOfFunctionPlot[57].X = -0.99
	pointsOfFunctionPlot[57].Y = -1.523

	pointsOfFunctionPlot[58].X = -0.98
	pointsOfFunctionPlot[58].Y = -1.491

	pointsOfFunctionPlot[59].X = -0.97
	pointsOfFunctionPlot[59].Y = -1.459

	pointsOfFunctionPlot[60].X = -0.96
	pointsOfFunctionPlot[60].Y = -1.428

	pointsOfFunctionPlot[61].X = -0.95
	pointsOfFunctionPlot[61].Y = -1.398

	pointsOfFunctionPlot[62].X = -0.94
	pointsOfFunctionPlot[62].Y = -1.369

	pointsOfFunctionPlot[63].X = -0.93
	pointsOfFunctionPlot[63].Y = -1.34

	pointsOfFunctionPlot[64].X = -0.92
	pointsOfFunctionPlot[64].Y = -1.313

	pointsOfFunctionPlot[65].X = -0.91
	pointsOfFunctionPlot[65].Y = -1.286

	pointsOfFunctionPlot[66].X = -0.90
	pointsOfFunctionPlot[66].Y = -1.26

	pointsOfFunctionPlot[67].X = -0.89
	pointsOfFunctionPlot[67].Y = -1.234

	pointsOfFunctionPlot[68].X = -0.88
	pointsOfFunctionPlot[68].Y = -1.209

	pointsOfFunctionPlot[69].X = -0.87
	pointsOfFunctionPlot[69].Y = -1.185

	pointsOfFunctionPlot[70].X = -0.86
	pointsOfFunctionPlot[70].Y = -1.161

	pointsOfFunctionPlot[71].X = -0.85
	pointsOfFunctionPlot[71].Y = -1.138

	pointsOfFunctionPlot[72].X = -0.84
	pointsOfFunctionPlot[72].Y = -1.115

	pointsOfFunctionPlot[73].X = -0.83
	pointsOfFunctionPlot[73].Y = -1.093

	pointsOfFunctionPlot[74].X = -0.82
	pointsOfFunctionPlot[74].Y = -1.071

	pointsOfFunctionPlot[75].X = -0.81
	pointsOfFunctionPlot[75].Y = -1.05

	pointsOfFunctionPlot[76].X = -0.80
	pointsOfFunctionPlot[76].Y = -1.029

	pointsOfFunctionPlot[77].X = -0.79
	pointsOfFunctionPlot[77].Y = -1.009

	pointsOfFunctionPlot[78].X = -0.78
	pointsOfFunctionPlot[78].Y = -0.989

	pointsOfFunctionPlot[79].X = -0.77
	pointsOfFunctionPlot[79].Y = -0.969

	pointsOfFunctionPlot[80].X = -0.76
	pointsOfFunctionPlot[80].Y = -0.95

	pointsOfFunctionPlot[81].X = -0.75
	pointsOfFunctionPlot[81].Y = -0.931

	pointsOfFunctionPlot[82].X = -0.74
	pointsOfFunctionPlot[82].Y = -0.913

	pointsOfFunctionPlot[83].X = -0.73
	pointsOfFunctionPlot[83].Y = -0.894

	pointsOfFunctionPlot[84].X = -0.72
	pointsOfFunctionPlot[84].Y = -0.877

	pointsOfFunctionPlot[85].X = -0.71
	pointsOfFunctionPlot[85].Y = -0.859

	pointsOfFunctionPlot[86].X = -0.70
	pointsOfFunctionPlot[86].Y = -0.842

	pointsOfFunctionPlot[87].X = -0.69
	pointsOfFunctionPlot[87].Y = -0.825

	pointsOfFunctionPlot[88].X = -0.68
	pointsOfFunctionPlot[88].Y = -0.808

	pointsOfFunctionPlot[89].X = -0.67
	pointsOfFunctionPlot[89].Y = -0.792

	pointsOfFunctionPlot[90].X = -0.66
	pointsOfFunctionPlot[90].Y = -0.776

	pointsOfFunctionPlot[91].X = -0.65
	pointsOfFunctionPlot[91].Y = -0.76

	pointsOfFunctionPlot[92].X = -0.64
	pointsOfFunctionPlot[92].Y = -0.744

	pointsOfFunctionPlot[93].X = -0.63
	pointsOfFunctionPlot[93].Y = -0.729

	pointsOfFunctionPlot[94].X = -0.62
	pointsOfFunctionPlot[94].Y = -0.713

	pointsOfFunctionPlot[95].X = -0.61
	pointsOfFunctionPlot[95].Y = -0.698

	pointsOfFunctionPlot[96].X = -0.60
	pointsOfFunctionPlot[96].Y = -0.684

	pointsOfFunctionPlot[97].X = -0.59
	pointsOfFunctionPlot[97].Y = -0.669

	pointsOfFunctionPlot[98].X = -0.58
	pointsOfFunctionPlot[98].Y = -0.655

	pointsOfFunctionPlot[99].X = -0.57
	pointsOfFunctionPlot[99].Y = -0.641

	pointsOfFunctionPlot[100].X = -0.56
	pointsOfFunctionPlot[100].Y = -0.626

	pointsOfFunctionPlot[101].X = -0.55
	pointsOfFunctionPlot[101].Y = -0.613

	pointsOfFunctionPlot[102].X = -0.54
	pointsOfFunctionPlot[102].Y = -0.599

	pointsOfFunctionPlot[103].X = -0.53
	pointsOfFunctionPlot[103].Y = -0.585

	pointsOfFunctionPlot[104].X = -0.52
	pointsOfFunctionPlot[104].Y = -0.572

	pointsOfFunctionPlot[105].X = -0.51
	pointsOfFunctionPlot[105].Y = -0.559

	pointsOfFunctionPlot[106].X = -0.50
	pointsOfFunctionPlot[106].Y = -0.546

	pointsOfFunctionPlot[107].X = -0.49
	pointsOfFunctionPlot[107].Y = -0.533

	pointsOfFunctionPlot[108].X = -0.48
	pointsOfFunctionPlot[108].Y = -0.52

	pointsOfFunctionPlot[109].X = -0.47
	pointsOfFunctionPlot[109].Y = -0.508

	pointsOfFunctionPlot[110].X = -0.46
	pointsOfFunctionPlot[110].Y = -0.495

	pointsOfFunctionPlot[111].X = -0.45
	pointsOfFunctionPlot[111].Y = -0.483

	pointsOfFunctionPlot[112].X = -0.44
	pointsOfFunctionPlot[112].Y = -0.47

	pointsOfFunctionPlot[113].X = -0.43
	pointsOfFunctionPlot[113].Y = -0.458

	pointsOfFunctionPlot[114].X = -0.42
	pointsOfFunctionPlot[114].Y = -0.446

	pointsOfFunctionPlot[115].X = -0.41
	pointsOfFunctionPlot[115].Y = -0.434

	pointsOfFunctionPlot[116].X = -0.40
	pointsOfFunctionPlot[116].Y = -0.422

	pointsOfFunctionPlot[117].X = -0.39
	pointsOfFunctionPlot[117].Y = -0.411

	pointsOfFunctionPlot[118].X = -0.38
	pointsOfFunctionPlot[118].Y = -0.399

	pointsOfFunctionPlot[119].X = -0.37
	pointsOfFunctionPlot[119].Y = -0.387

	pointsOfFunctionPlot[120].X = -0.36
	pointsOfFunctionPlot[120].Y = -0.376

	pointsOfFunctionPlot[121].X = -0.35
	pointsOfFunctionPlot[121].Y = -0.365

	pointsOfFunctionPlot[122].X = -0.34
	pointsOfFunctionPlot[122].Y = -0.353

	pointsOfFunctionPlot[123].X = -0.33
	pointsOfFunctionPlot[123].Y = -0.342

	pointsOfFunctionPlot[124].X = -0.32
	pointsOfFunctionPlot[124].Y = -0.331

	pointsOfFunctionPlot[125].X = -0.31
	pointsOfFunctionPlot[125].Y = -0.32

	pointsOfFunctionPlot[126].X = -0.30
	pointsOfFunctionPlot[126].Y = -0.309

	pointsOfFunctionPlot[127].X = -0.29
	pointsOfFunctionPlot[127].Y = -0.298

	pointsOfFunctionPlot[128].X = -0.28
	pointsOfFunctionPlot[128].Y = -0.287

	pointsOfFunctionPlot[129].X = -0.27
	pointsOfFunctionPlot[129].Y = -0.276

	pointsOfFunctionPlot[130].X = -0.26
	pointsOfFunctionPlot[130].Y = -0.266

	pointsOfFunctionPlot[131].X = -0.25
	pointsOfFunctionPlot[131].Y = -0.255

	pointsOfFunctionPlot[132].X = -0.24
	pointsOfFunctionPlot[132].Y = -0.244

	pointsOfFunctionPlot[133].X = -0.23
	pointsOfFunctionPlot[133].Y = -0.234

	pointsOfFunctionPlot[134].X = -0.22
	pointsOfFunctionPlot[134].Y = -0.223

	pointsOfFunctionPlot[135].X = -0.21
	pointsOfFunctionPlot[135].Y = -0.213

	pointsOfFunctionPlot[136].X = -0.20
	pointsOfFunctionPlot[136].Y = -0.202

	pointsOfFunctionPlot[137].X = -0.19
	pointsOfFunctionPlot[137].Y = -0.192

	pointsOfFunctionPlot[138].X = -0.18
	pointsOfFunctionPlot[138].Y = -0.182

	pointsOfFunctionPlot[139].X = -0.17
	pointsOfFunctionPlot[139].Y = -0.171

	pointsOfFunctionPlot[140].X = -0.16
	pointsOfFunctionPlot[140].Y = -0.161

	pointsOfFunctionPlot[141].X = -0.15
	pointsOfFunctionPlot[141].Y = -0.151

	pointsOfFunctionPlot[142].X = -0.14
	pointsOfFunctionPlot[142].Y = -0.14

	pointsOfFunctionPlot[143].X = -0.13
	pointsOfFunctionPlot[143].Y = -0.13

	pointsOfFunctionPlot[144].X = -0.12
	pointsOfFunctionPlot[144].Y = -0.12

	pointsOfFunctionPlot[145].X = -0.11
	pointsOfFunctionPlot[145].Y = -0.11

	pointsOfFunctionPlot[146].X = -0.10
	pointsOfFunctionPlot[146].Y = -0.1

	pointsOfFunctionPlot[147].X = -0.09
	pointsOfFunctionPlot[147].Y = -0.09

	pointsOfFunctionPlot[148].X = -0.08
	pointsOfFunctionPlot[148].Y = -0.08

	pointsOfFunctionPlot[149].X = -0.07
	pointsOfFunctionPlot[149].Y = -0.07

	pointsOfFunctionPlot[150].X = -0.06
	pointsOfFunctionPlot[150].Y = -0.06

	pointsOfFunctionPlot[151].X = -0.05
	pointsOfFunctionPlot[151].Y = -0.05

	pointsOfFunctionPlot[152].X = -0.04
	pointsOfFunctionPlot[152].Y = -0.04

	pointsOfFunctionPlot[153].X = -0.03
	pointsOfFunctionPlot[153].Y = -0.03

	pointsOfFunctionPlot[154].X = -0.02
	pointsOfFunctionPlot[154].Y = -0.02

	pointsOfFunctionPlot[155].X = -0.01
	pointsOfFunctionPlot[155].Y = -0.01

	pointsOfFunctionPlot[156].X = 0.0
	pointsOfFunctionPlot[156].Y = 0.0

	pointsOfFunctionPlot[157].X = 0.01
	pointsOfFunctionPlot[157].Y = 0.01

	pointsOfFunctionPlot[158].X = 0.02
	pointsOfFunctionPlot[158].Y = 0.02

	pointsOfFunctionPlot[159].X = 0.03
	pointsOfFunctionPlot[159].Y = 0.03

	pointsOfFunctionPlot[160].X = 0.04
	pointsOfFunctionPlot[160].Y = 0.04

	pointsOfFunctionPlot[161].X = 0.05
	pointsOfFunctionPlot[161].Y = 0.05

	pointsOfFunctionPlot[162].X = 0.06
	pointsOfFunctionPlot[162].Y = 0.06

	pointsOfFunctionPlot[163].X = 0.07
	pointsOfFunctionPlot[163].Y = 0.07

	pointsOfFunctionPlot[164].X = 0.08
	pointsOfFunctionPlot[164].Y = 0.08

	pointsOfFunctionPlot[165].X = 0.09
	pointsOfFunctionPlot[165].Y = 0.09

	pointsOfFunctionPlot[166].X = 0.10
	pointsOfFunctionPlot[166].Y = 0.10

	pointsOfFunctionPlot[167].X = 0.11
	pointsOfFunctionPlot[167].Y = 0.11

	pointsOfFunctionPlot[168].X = 0.12
	pointsOfFunctionPlot[168].Y = 0.12

	pointsOfFunctionPlot[169].X = 0.13
	pointsOfFunctionPlot[169].Y = 0.13

	pointsOfFunctionPlot[170].X = 0.14
	pointsOfFunctionPlot[170].Y = 0.14

	pointsOfFunctionPlot[171].X = 0.15
	pointsOfFunctionPlot[171].Y = 0.151

	pointsOfFunctionPlot[172].X = 0.16
	pointsOfFunctionPlot[172].Y = 0.161

	pointsOfFunctionPlot[173].X = 0.17
	pointsOfFunctionPlot[173].Y = 0.171

	pointsOfFunctionPlot[174].X = 0.18
	pointsOfFunctionPlot[174].Y = 0.182

	pointsOfFunctionPlot[175].X = 0.19
	pointsOfFunctionPlot[175].Y = 0.192

	pointsOfFunctionPlot[176].X = 0.20
	pointsOfFunctionPlot[176].Y = 0.202

	pointsOfFunctionPlot[177].X = 0.21
	pointsOfFunctionPlot[177].Y = 0.213

	pointsOfFunctionPlot[178].X = 0.22
	pointsOfFunctionPlot[178].Y = 0.223

	pointsOfFunctionPlot[179].X = 0.23
	pointsOfFunctionPlot[179].Y = 0.234

	pointsOfFunctionPlot[180].X = 0.24
	pointsOfFunctionPlot[180].Y = 0.244

	pointsOfFunctionPlot[181].X = 0.25
	pointsOfFunctionPlot[181].Y = 0.255

	pointsOfFunctionPlot[182].X = 0.26
	pointsOfFunctionPlot[182].Y = 0.266

	pointsOfFunctionPlot[183].X = 0.27
	pointsOfFunctionPlot[183].Y = 0.276

	pointsOfFunctionPlot[184].X = 0.28
	pointsOfFunctionPlot[184].Y = 0.287

	pointsOfFunctionPlot[185].X = 0.29
	pointsOfFunctionPlot[185].Y = 0.298

	pointsOfFunctionPlot[186].X = 0.30
	pointsOfFunctionPlot[186].Y = 0.309

	pointsOfFunctionPlot[187].X = 0.31
	pointsOfFunctionPlot[187].Y = 0.32

	pointsOfFunctionPlot[188].X = 0.32
	pointsOfFunctionPlot[188].Y = 0.331

	pointsOfFunctionPlot[189].X = 0.33
	pointsOfFunctionPlot[189].Y = 0.342

	pointsOfFunctionPlot[190].X = 0.34
	pointsOfFunctionPlot[190].Y = 0.353

	pointsOfFunctionPlot[191].X = 0.35
	pointsOfFunctionPlot[191].Y = 0.365

	pointsOfFunctionPlot[192].X = 0.36
	pointsOfFunctionPlot[192].Y = 0.376

	pointsOfFunctionPlot[193].X = 0.37
	pointsOfFunctionPlot[193].Y = 0.387

	pointsOfFunctionPlot[194].X = 0.38
	pointsOfFunctionPlot[194].Y = 0.399

	pointsOfFunctionPlot[195].X = 0.39
	pointsOfFunctionPlot[195].Y = 0.411

	pointsOfFunctionPlot[196].X = 0.40
	pointsOfFunctionPlot[196].Y = 0.422

	pointsOfFunctionPlot[197].X = 0.41
	pointsOfFunctionPlot[197].Y = 0.434

	pointsOfFunctionPlot[198].X = 0.42
	pointsOfFunctionPlot[198].Y = 0.446

	pointsOfFunctionPlot[199].X = 0.43
	pointsOfFunctionPlot[199].Y = 0.458

	pointsOfFunctionPlot[200].X = 0.44
	pointsOfFunctionPlot[200].Y = 0.47

	pointsOfFunctionPlot[201].X = 0.45
	pointsOfFunctionPlot[201].Y = 0.483

	pointsOfFunctionPlot[202].X = 0.46
	pointsOfFunctionPlot[202].Y = 0.495

	pointsOfFunctionPlot[203].X = 0.47
	pointsOfFunctionPlot[203].Y = 0.508

	pointsOfFunctionPlot[204].X = 0.48
	pointsOfFunctionPlot[204].Y = 0.52

	pointsOfFunctionPlot[205].X = 0.49
	pointsOfFunctionPlot[205].Y = 0.533

	pointsOfFunctionPlot[206].X = 0.50
	pointsOfFunctionPlot[206].Y = 0.546

	pointsOfFunctionPlot[207].X = 0.51
	pointsOfFunctionPlot[207].Y = 0.559

	pointsOfFunctionPlot[208].X = 0.52
	pointsOfFunctionPlot[208].Y = 0.572

	pointsOfFunctionPlot[209].X = 0.53
	pointsOfFunctionPlot[209].Y = 0.585

	pointsOfFunctionPlot[210].X = 0.54
	pointsOfFunctionPlot[210].Y = 0.599

	pointsOfFunctionPlot[211].X = 0.55
	pointsOfFunctionPlot[211].Y = 0.613

	pointsOfFunctionPlot[212].X = 0.56
	pointsOfFunctionPlot[212].Y = 0.626

	pointsOfFunctionPlot[213].X = 0.57
	pointsOfFunctionPlot[213].Y = 0.641

	pointsOfFunctionPlot[214].X = 0.58
	pointsOfFunctionPlot[214].Y = 0.655

	pointsOfFunctionPlot[215].X = 0.59
	pointsOfFunctionPlot[215].Y = 0.669

	pointsOfFunctionPlot[216].X = 0.60
	pointsOfFunctionPlot[216].Y = 0.684

	pointsOfFunctionPlot[217].X = 0.61
	pointsOfFunctionPlot[217].Y = 0.698

	pointsOfFunctionPlot[218].X = 0.62
	pointsOfFunctionPlot[218].Y = 0.713

	pointsOfFunctionPlot[219].X = 0.63
	pointsOfFunctionPlot[219].Y = 0.729

	pointsOfFunctionPlot[220].X = 0.64
	pointsOfFunctionPlot[220].Y = 0.744

	pointsOfFunctionPlot[221].X = 0.65
	pointsOfFunctionPlot[221].Y = 0.76

	pointsOfFunctionPlot[222].X = 0.66
	pointsOfFunctionPlot[222].Y = 0.776

	pointsOfFunctionPlot[223].X = 0.67
	pointsOfFunctionPlot[223].Y = 0.792

	pointsOfFunctionPlot[224].X = 0.68
	pointsOfFunctionPlot[224].Y = 0.808

	pointsOfFunctionPlot[225].X = 0.69
	pointsOfFunctionPlot[225].Y = 0.825

	pointsOfFunctionPlot[226].X = 0.70
	pointsOfFunctionPlot[226].Y = 0.842

	pointsOfFunctionPlot[227].X = 0.71
	pointsOfFunctionPlot[227].Y = 0.859

	pointsOfFunctionPlot[228].X = 0.72
	pointsOfFunctionPlot[228].Y = 0.877

	pointsOfFunctionPlot[229].X = 0.73
	pointsOfFunctionPlot[229].Y = 0.894

	pointsOfFunctionPlot[230].X = 0.74
	pointsOfFunctionPlot[230].Y = 0.913

	pointsOfFunctionPlot[231].X = 0.75
	pointsOfFunctionPlot[231].Y = 0.931

	pointsOfFunctionPlot[232].X = 0.76
	pointsOfFunctionPlot[232].Y = 0.95

	pointsOfFunctionPlot[233].X = 0.77
	pointsOfFunctionPlot[233].Y = 0.969

	pointsOfFunctionPlot[234].X = 0.78
	pointsOfFunctionPlot[234].Y = 0.989

	pointsOfFunctionPlot[235].X = 0.79
	pointsOfFunctionPlot[235].Y = 1.009

	pointsOfFunctionPlot[236].X = 0.80
	pointsOfFunctionPlot[236].Y = 1.029

	pointsOfFunctionPlot[237].X = 0.81
	pointsOfFunctionPlot[237].Y = 1.05

	pointsOfFunctionPlot[238].X = 0.82
	pointsOfFunctionPlot[238].Y = 1.071

	pointsOfFunctionPlot[239].X = 0.83
	pointsOfFunctionPlot[239].Y = 1.093

	pointsOfFunctionPlot[240].X = 0.84
	pointsOfFunctionPlot[240].Y = 1.115

	pointsOfFunctionPlot[241].X = 0.85
	pointsOfFunctionPlot[241].Y = 1.138

	pointsOfFunctionPlot[242].X = 0.86
	pointsOfFunctionPlot[242].Y = 1.161

	pointsOfFunctionPlot[243].X = 0.87
	pointsOfFunctionPlot[243].Y = 1.185

	pointsOfFunctionPlot[244].X = 0.88
	pointsOfFunctionPlot[244].Y = 1.209

	pointsOfFunctionPlot[245].X = 0.89
	pointsOfFunctionPlot[245].Y = 1.234

	pointsOfFunctionPlot[246].X = 0.90
	pointsOfFunctionPlot[246].Y = 1.26

	pointsOfFunctionPlot[247].X = 0.91
	pointsOfFunctionPlot[247].Y = 1.286

	pointsOfFunctionPlot[248].X = 0.92
	pointsOfFunctionPlot[248].Y = 1.313

	pointsOfFunctionPlot[249].X = 0.93
	pointsOfFunctionPlot[249].Y = 1.34

	pointsOfFunctionPlot[250].X = 0.94
	pointsOfFunctionPlot[250].Y = 1.369

	pointsOfFunctionPlot[251].X = 0.95
	pointsOfFunctionPlot[251].Y = 1.398

	pointsOfFunctionPlot[252].X = 0.96
	pointsOfFunctionPlot[252].Y = 1.428

	pointsOfFunctionPlot[253].X = 0.97
	pointsOfFunctionPlot[253].Y = 1.459

	pointsOfFunctionPlot[254].X = 0.98
	pointsOfFunctionPlot[254].Y = 1.491

	pointsOfFunctionPlot[255].X = 0.99
	pointsOfFunctionPlot[255].Y = 1.523

	pointsOfFunctionPlot[256].X = 1.0
	pointsOfFunctionPlot[256].Y = 1.557

	pointsOfFunctionPlot[257].X = 1.01
	pointsOfFunctionPlot[257].Y = 1.592

	pointsOfFunctionPlot[258].X = 1.02
	pointsOfFunctionPlot[258].Y = 1.628

	pointsOfFunctionPlot[259].X = 1.03
	pointsOfFunctionPlot[259].Y = 1.665

	pointsOfFunctionPlot[260].X = 1.04
	pointsOfFunctionPlot[260].Y = 1.703

	pointsOfFunctionPlot[261].X = 1.05
	pointsOfFunctionPlot[261].Y = 1.743

	pointsOfFunctionPlot[262].X = 1.06
	pointsOfFunctionPlot[262].Y = 1.784

	pointsOfFunctionPlot[263].X = 1.07
	pointsOfFunctionPlot[263].Y = 1.827

	pointsOfFunctionPlot[264].X = 1.08
	pointsOfFunctionPlot[264].Y = 1.871

	pointsOfFunctionPlot[265].X = 1.09
	pointsOfFunctionPlot[265].Y = 1.917

	pointsOfFunctionPlot[266].X = 1.10
	pointsOfFunctionPlot[266].Y = 1.964

	pointsOfFunctionPlot[267].X = 1.11
	pointsOfFunctionPlot[267].Y = 2.014

	pointsOfFunctionPlot[268].X = 1.12
	pointsOfFunctionPlot[268].Y = 2.066

	pointsOfFunctionPlot[269].X = 1.13
	pointsOfFunctionPlot[269].Y = 2.119

	pointsOfFunctionPlot[270].X = 1.14
	pointsOfFunctionPlot[270].Y = 2.175

	pointsOfFunctionPlot[271].X = 1.15
	pointsOfFunctionPlot[271].Y = 2.234

	pointsOfFunctionPlot[272].X = 1.16
	pointsOfFunctionPlot[272].Y = 2.295

	pointsOfFunctionPlot[273].X = 1.17
	pointsOfFunctionPlot[273].Y = 2.36

	pointsOfFunctionPlot[274].X = 1.18
	pointsOfFunctionPlot[274].Y = 2.427

	pointsOfFunctionPlot[275].X = 1.19
	pointsOfFunctionPlot[275].Y = 2.497

	pointsOfFunctionPlot[276].X = 1.20
	pointsOfFunctionPlot[276].Y = 2.572

	pointsOfFunctionPlot[277].X = 1.21
	pointsOfFunctionPlot[277].Y = 2.65

	pointsOfFunctionPlot[278].X = 1.22
	pointsOfFunctionPlot[278].Y = 2.732

	pointsOfFunctionPlot[279].X = 1.23
	pointsOfFunctionPlot[279].Y = 2.819

	pointsOfFunctionPlot[280].X = 1.24
	pointsOfFunctionPlot[280].Y = 2.911

	pointsOfFunctionPlot[281].X = 1.25
	pointsOfFunctionPlot[281].Y = 3.009

	pointsOfFunctionPlot[282].X = 1.26
	pointsOfFunctionPlot[282].Y = 3.113

	pointsOfFunctionPlot[283].X = 1.27
	pointsOfFunctionPlot[283].Y = 3.223

	pointsOfFunctionPlot[284].X = 1.28
	pointsOfFunctionPlot[284].Y = 3.341

	pointsOfFunctionPlot[285].X = 1.29
	pointsOfFunctionPlot[285].Y = 3.467

	pointsOfFunctionPlot[286].X = 1.30
	pointsOfFunctionPlot[286].Y = 3.602

	pointsOfFunctionPlot[287].X = 1.31
	pointsOfFunctionPlot[287].Y = 3.747

	pointsOfFunctionPlot[288].X = 1.32
	pointsOfFunctionPlot[288].Y = 3.903

	pointsOfFunctionPlot[289].X = 1.33
	pointsOfFunctionPlot[289].Y = 4.072

	pointsOfFunctionPlot[290].X = 1.34
	pointsOfFunctionPlot[290].Y = 4.255

	pointsOfFunctionPlot[291].X = 1.35
	pointsOfFunctionPlot[291].Y = 4.455

	pointsOfFunctionPlot[292].X = 1.36
	pointsOfFunctionPlot[292].Y = 4.673

	pointsOfFunctionPlot[293].X = 1.37
	pointsOfFunctionPlot[293].Y = 4.913

	pointsOfFunctionPlot[294].X = 1.38
	pointsOfFunctionPlot[294].Y = 5.177

	pointsOfFunctionPlot[295].X = 1.39
	pointsOfFunctionPlot[295].Y = 5.47

	pointsOfFunctionPlot[296].X = 1.40
	pointsOfFunctionPlot[296].Y = 5.797

	pointsOfFunctionPlot[297].X = 1.41
	pointsOfFunctionPlot[297].Y = 6.165

	pointsOfFunctionPlot[298].X = 1.42
	pointsOfFunctionPlot[298].Y = 6.581

	pointsOfFunctionPlot[299].X = 1.43
	pointsOfFunctionPlot[299].Y = 7.055

	pointsOfFunctionPlot[300].X = 1.44
	pointsOfFunctionPlot[300].Y = 7.601

	pointsOfFunctionPlot[301].X = 1.45
	pointsOfFunctionPlot[301].Y = 8.238

	pointsOfFunctionPlot[302].X = 1.46
	pointsOfFunctionPlot[302].Y = 8.988

	pointsOfFunctionPlot[303].X = 1.47
	pointsOfFunctionPlot[303].Y = 9.887

	pointsOfFunctionPlot[304].X = 1.48
	pointsOfFunctionPlot[304].Y = 10.983

	pointsOfFunctionPlot[305].X = 1.49
	pointsOfFunctionPlot[305].Y = 12.349

	pointsOfFunctionPlot[306].X = 1.50
	pointsOfFunctionPlot[306].Y = 14.101

	pointsOfFunctionPlot[307].X = 1.51
	pointsOfFunctionPlot[307].Y = 16.428

	pointsOfFunctionPlot[308].X = 1.52
	pointsOfFunctionPlot[308].Y = 19.669

	pointsOfFunctionPlot[309].X = 1.53
	pointsOfFunctionPlot[309].Y = 24.498

	pointsOfFunctionPlot[310].X = 1.54
	pointsOfFunctionPlot[310].Y = 32.461

	pointsOfFunctionPlot[311].X = 1.55
	pointsOfFunctionPlot[311].Y = 48.078

	pointsOfFunctionPlot[312].X = 1.56
	pointsOfFunctionPlot[312].Y = 92.62

	// pointsOfFunctionPlot[313].X = -6.87
	// pointsOfFunctionPlot[313].Y = 0.008

	// pointsOfFunctionPlot[314].X = -6.86
	// pointsOfFunctionPlot[314].Y = 0.008

	// pointsOfFunctionPlot[315].X = -6.85
	// pointsOfFunctionPlot[315].Y = 0.008

	// pointsOfFunctionPlot[316].X = -6.84
	// pointsOfFunctionPlot[316].Y = 0.008

	// pointsOfFunctionPlot[317].X = -6.83
	// pointsOfFunctionPlot[317].Y = 0.008

	// pointsOfFunctionPlot[318].X = -6.82
	// pointsOfFunctionPlot[318].Y = 0.008

	// pointsOfFunctionPlot[319].X = -6.81
	// pointsOfFunctionPlot[319].Y = 0.008

	// pointsOfFunctionPlot[320].X = -6.8
	// pointsOfFunctionPlot[320].Y = 0.009

	// pointsOfFunctionPlot[321].X = -6.79
	// pointsOfFunctionPlot[321].Y = 0.009

	// pointsOfFunctionPlot[322].X = -6.78
	// pointsOfFunctionPlot[322].Y = 0.009

	// pointsOfFunctionPlot[323].X = -6.77
	// pointsOfFunctionPlot[323].Y = 0.009

	// pointsOfFunctionPlot[324].X = -6.76
	// pointsOfFunctionPlot[324].Y = 0.009

	// pointsOfFunctionPlot[325].X = -6.75
	// pointsOfFunctionPlot[325].Y = 0.009

	// pointsOfFunctionPlot[326].X = -6.74
	// pointsOfFunctionPlot[326].Y = 0.009

	// pointsOfFunctionPlot[327].X = -6.73
	// pointsOfFunctionPlot[327].Y = 0.009

	// pointsOfFunctionPlot[328].X = -6.72
	// pointsOfFunctionPlot[328].Y = 0.009

	// pointsOfFunctionPlot[329].X = -6.71
	// pointsOfFunctionPlot[329].Y = 0.009

	// pointsOfFunctionPlot[330].X = -6.7
	// pointsOfFunctionPlot[330].Y = 0.009

	// pointsOfFunctionPlot[331].X = -6.69
	// pointsOfFunctionPlot[331].Y = 0.009

	// pointsOfFunctionPlot[332].X = -6.68
	// pointsOfFunctionPlot[332].Y = 0.009

	// pointsOfFunctionPlot[333].X = -6.67
	// pointsOfFunctionPlot[333].Y = 0.009

	// pointsOfFunctionPlot[334].X = -6.66
	// pointsOfFunctionPlot[334].Y = 0.009

	// pointsOfFunctionPlot[335].X = -6.65
	// pointsOfFunctionPlot[335].Y = 0.01

	// pointsOfFunctionPlot[336].X = -6.64
	// pointsOfFunctionPlot[336].Y = 0.01

	// pointsOfFunctionPlot[337].X = -6.63
	// pointsOfFunctionPlot[337].Y = 0.01

	// pointsOfFunctionPlot[338].X = -6.62
	// pointsOfFunctionPlot[338].Y = 0.01

	// pointsOfFunctionPlot[339].X = -6.61
	// pointsOfFunctionPlot[339].Y = 0.01

	// pointsOfFunctionPlot[340].X = -6.6
	// pointsOfFunctionPlot[340].Y = 0.01

	// pointsOfFunctionPlot[341].X = -6.59
	// pointsOfFunctionPlot[341].Y = 0.01

	// pointsOfFunctionPlot[342].X = -6.58
	// pointsOfFunctionPlot[342].Y = 0.01

	// pointsOfFunctionPlot[343].X = -6.57
	// pointsOfFunctionPlot[343].Y = 0.01

	// pointsOfFunctionPlot[344].X = -6.56
	// pointsOfFunctionPlot[344].Y = 0.01

	// pointsOfFunctionPlot[345].X = -6.55
	// pointsOfFunctionPlot[345].Y = 0.01

	// pointsOfFunctionPlot[346].X = -6.54
	// pointsOfFunctionPlot[346].Y = 0.01

	// pointsOfFunctionPlot[347].X = -6.53
	// pointsOfFunctionPlot[347].Y = 0.01

	// pointsOfFunctionPlot[348].X = -6.52
	// pointsOfFunctionPlot[348].Y = 0.01

	// pointsOfFunctionPlot[349].X = -6.51
	// pointsOfFunctionPlot[349].Y = 0.011

	// pointsOfFunctionPlot[350].X = -6.5
	// pointsOfFunctionPlot[350].Y = 0.011

	// pointsOfFunctionPlot[351].X = -6.49
	// pointsOfFunctionPlot[351].Y = 0.011

	// pointsOfFunctionPlot[352].X = -6.48
	// pointsOfFunctionPlot[352].Y = 0.011

	// pointsOfFunctionPlot[353].X = -6.47
	// pointsOfFunctionPlot[353].Y = 0.011

	// pointsOfFunctionPlot[354].X = -6.46
	// pointsOfFunctionPlot[354].Y = 0.011

	// pointsOfFunctionPlot[355].X = -6.45
	// pointsOfFunctionPlot[355].Y = 0.011

	// pointsOfFunctionPlot[356].X = -6.44
	// pointsOfFunctionPlot[356].Y = 0.011

	// pointsOfFunctionPlot[357].X = -6.43
	// pointsOfFunctionPlot[357].Y = 0.011

	// pointsOfFunctionPlot[358].X = -6.42
	// pointsOfFunctionPlot[358].Y = 0.011

	// pointsOfFunctionPlot[359].X = -6.41
	// pointsOfFunctionPlot[359].Y = 0.011

	// pointsOfFunctionPlot[360].X = -6.4
	// pointsOfFunctionPlot[360].Y = 0.011

	// pointsOfFunctionPlot[361].X = -6.39
	// pointsOfFunctionPlot[361].Y = 0.011

	// pointsOfFunctionPlot[362].X = -6.38
	// pointsOfFunctionPlot[362].Y = 0.012

	// pointsOfFunctionPlot[363].X = -6.37
	// pointsOfFunctionPlot[363].Y = 0.012

	// pointsOfFunctionPlot[364].X = -6.36
	// pointsOfFunctionPlot[364].Y = 0.012

	// pointsOfFunctionPlot[365].X = -6.35
	// pointsOfFunctionPlot[365].Y = 0.012

	// pointsOfFunctionPlot[366].X = -6.34
	// pointsOfFunctionPlot[366].Y = 0.012

	// pointsOfFunctionPlot[367].X = -6.33
	// pointsOfFunctionPlot[367].Y = 0.012

	// pointsOfFunctionPlot[368].X = -6.32
	// pointsOfFunctionPlot[368].Y = 0.012

	// pointsOfFunctionPlot[369].X = -6.31
	// pointsOfFunctionPlot[369].Y = 0.012

	// pointsOfFunctionPlot[370].X = -6.3
	// pointsOfFunctionPlot[370].Y = 0.012

	// pointsOfFunctionPlot[371].X = -6.29
	// pointsOfFunctionPlot[371].Y = 0.012

	// pointsOfFunctionPlot[372].X = -6.28
	// pointsOfFunctionPlot[372].Y = 0.012

	// pointsOfFunctionPlot[373].X = -6.27
	// pointsOfFunctionPlot[373].Y = 0.013

	// pointsOfFunctionPlot[374].X = -6.26
	// pointsOfFunctionPlot[374].Y = 0.013

	// pointsOfFunctionPlot[375].X = -6.25
	// pointsOfFunctionPlot[375].Y = 0.013

	// pointsOfFunctionPlot[376].X = -6.24
	// pointsOfFunctionPlot[376].Y = 0.013

	// pointsOfFunctionPlot[377].X = -6.23
	// pointsOfFunctionPlot[377].Y = 0.013

	// pointsOfFunctionPlot[378].X = -6.22
	// pointsOfFunctionPlot[378].Y = 0.013

	// pointsOfFunctionPlot[379].X = -6.21
	// pointsOfFunctionPlot[379].Y = 0.013

	// pointsOfFunctionPlot[380].X = -6.2
	// pointsOfFunctionPlot[380].Y = 0.013

	// pointsOfFunctionPlot[381].X = -6.19
	// pointsOfFunctionPlot[381].Y = 0.013

	// pointsOfFunctionPlot[382].X = -6.18
	// pointsOfFunctionPlot[382].Y = 0.013

	// pointsOfFunctionPlot[383].X = -6.17
	// pointsOfFunctionPlot[383].Y = 0.013

	// pointsOfFunctionPlot[384].X = -6.16
	// pointsOfFunctionPlot[384].Y = 0.014

	// pointsOfFunctionPlot[385].X = -6.15
	// pointsOfFunctionPlot[385].Y = 0.014

	// pointsOfFunctionPlot[386].X = -6.14
	// pointsOfFunctionPlot[386].Y = 0.014

	// pointsOfFunctionPlot[387].X = -6.13
	// pointsOfFunctionPlot[387].Y = 0.014

	// pointsOfFunctionPlot[388].X = -6.12
	// pointsOfFunctionPlot[388].Y = 0.014

	// pointsOfFunctionPlot[389].X = -6.11
	// pointsOfFunctionPlot[389].Y = 0.014

	// pointsOfFunctionPlot[390].X = -6.1
	// pointsOfFunctionPlot[390].Y = 0.014

	// pointsOfFunctionPlot[391].X = -6.09
	// pointsOfFunctionPlot[391].Y = 0.014

	// pointsOfFunctionPlot[392].X = -6.08
	// pointsOfFunctionPlot[392].Y = 0.014

	// pointsOfFunctionPlot[393].X = -6.07
	// pointsOfFunctionPlot[393].Y = 0.014

	// pointsOfFunctionPlot[394].X = -6.06
	// pointsOfFunctionPlot[394].Y = 0.015

	// pointsOfFunctionPlot[395].X = -6.05
	// pointsOfFunctionPlot[395].Y = 0.015

	// pointsOfFunctionPlot[396].X = -6.04
	// pointsOfFunctionPlot[396].Y = 0.015

	// pointsOfFunctionPlot[397].X = -6.03
	// pointsOfFunctionPlot[397].Y = 0.015

	// pointsOfFunctionPlot[398].X = -6.02
	// pointsOfFunctionPlot[398].Y = 0.015

	// pointsOfFunctionPlot[399].X = -6.01
	// pointsOfFunctionPlot[399].Y = 0.015

	// pointsOfFunctionPlot[400].X = -6.0
	// pointsOfFunctionPlot[400].Y = 0.015

	// pointsOfFunctionPlot[401].X = -5.99
	// pointsOfFunctionPlot[401].Y = 0.015

	// pointsOfFunctionPlot[402].X = -5.98
	// pointsOfFunctionPlot[402].Y = 0.015

	// pointsOfFunctionPlot[403].X = -5.97
	// pointsOfFunctionPlot[403].Y = 0.016

	// pointsOfFunctionPlot[404].X = -5.96
	// pointsOfFunctionPlot[404].Y = 0.016

	// pointsOfFunctionPlot[405].X = -5.95
	// pointsOfFunctionPlot[405].Y = 0.016

	// pointsOfFunctionPlot[406].X = -5.94
	// pointsOfFunctionPlot[406].Y = 0.016

	// pointsOfFunctionPlot[407].X = -5.93
	// pointsOfFunctionPlot[407].Y = 0.016

	// pointsOfFunctionPlot[408].X = -5.92
	// pointsOfFunctionPlot[408].Y = 0.016

	// pointsOfFunctionPlot[409].X = -5.91
	// pointsOfFunctionPlot[409].Y = 0.016

	// pointsOfFunctionPlot[410].X = -5.9
	// pointsOfFunctionPlot[410].Y = 0.016

	// pointsOfFunctionPlot[411].X = -5.89
	// pointsOfFunctionPlot[411].Y = 0.016

	// pointsOfFunctionPlot[412].X = -5.88
	// pointsOfFunctionPlot[412].Y = 0.017

	// pointsOfFunctionPlot[413].X = -5.87
	// pointsOfFunctionPlot[413].Y = 0.017

	// pointsOfFunctionPlot[414].X = -5.86
	// pointsOfFunctionPlot[414].Y = 0.017

	// pointsOfFunctionPlot[415].X = -5.85
	// pointsOfFunctionPlot[415].Y = 0.017

	// pointsOfFunctionPlot[416].X = -5.84
	// pointsOfFunctionPlot[416].Y = 0.017

	// pointsOfFunctionPlot[417].X = -5.82
	// pointsOfFunctionPlot[417].Y = 0.017

	// pointsOfFunctionPlot[418].X = -5.82
	// pointsOfFunctionPlot[418].Y = 0.017

	// pointsOfFunctionPlot[419].X = -5.81
	// pointsOfFunctionPlot[419].Y = 0.017

	// pointsOfFunctionPlot[420].X = -5.8
	// pointsOfFunctionPlot[420].Y = 0.017

	// pointsOfFunctionPlot[421].X = -5.79
	// pointsOfFunctionPlot[421].Y = 0.018

	// pointsOfFunctionPlot[422].X = -5.78
	// pointsOfFunctionPlot[422].Y = 0.018

	// pointsOfFunctionPlot[423].X = -5.77
	// pointsOfFunctionPlot[423].Y = 0.018

	// pointsOfFunctionPlot[424].X = -5.76
	// pointsOfFunctionPlot[424].Y = 0.018

	// pointsOfFunctionPlot[425].X = -5.75
	// pointsOfFunctionPlot[425].Y = 0.018

	// pointsOfFunctionPlot[426].X = -5.74
	// pointsOfFunctionPlot[426].Y = 0.018

	// pointsOfFunctionPlot[427].X = -5.73
	// pointsOfFunctionPlot[427].Y = 0.018

	// pointsOfFunctionPlot[428].X = -5.72
	// pointsOfFunctionPlot[428].Y = 0.019

	// pointsOfFunctionPlot[429].X = -5.71
	// pointsOfFunctionPlot[429].Y = 0.019

	// pointsOfFunctionPlot[430].X = -5.7
	// pointsOfFunctionPlot[430].Y = 0.019

	// pointsOfFunctionPlot[431].X = -5.69
	// pointsOfFunctionPlot[431].Y = 0.019

	// pointsOfFunctionPlot[432].X = -5.68
	// pointsOfFunctionPlot[432].Y = 0.019

	// pointsOfFunctionPlot[433].X = -5.67
	// pointsOfFunctionPlot[433].Y = 0.019

	// pointsOfFunctionPlot[434].X = -5.66
	// pointsOfFunctionPlot[434].Y = 0.019

	// pointsOfFunctionPlot[435].X = -5.65
	// pointsOfFunctionPlot[435].Y = 0.019

	// pointsOfFunctionPlot[436].X = -5.64
	// pointsOfFunctionPlot[436].Y = 0.02

	// pointsOfFunctionPlot[437].X = -5.63
	// pointsOfFunctionPlot[437].Y = 0.02

	// pointsOfFunctionPlot[438].X = -5.62
	// pointsOfFunctionPlot[438].Y = 0.02

	// pointsOfFunctionPlot[439].X = -5.61
	// pointsOfFunctionPlot[439].Y = 0.02

	// pointsOfFunctionPlot[440].X = -5.6
	// pointsOfFunctionPlot[440].Y = 0.02

	// pointsOfFunctionPlot[441].X = -5.59
	// pointsOfFunctionPlot[441].Y = 0.02

	// pointsOfFunctionPlot[442].X = -5.58
	// pointsOfFunctionPlot[442].Y = 0.02

	// pointsOfFunctionPlot[443].X = -5.57
	// pointsOfFunctionPlot[443].Y = 0.021

	// pointsOfFunctionPlot[444].X = -5.56
	// pointsOfFunctionPlot[444].Y = 0.021

	// pointsOfFunctionPlot[445].X = -5.55
	// pointsOfFunctionPlot[445].Y = 0.021

	// pointsOfFunctionPlot[446].X = -5.54
	// pointsOfFunctionPlot[446].Y = 0.021

	// pointsOfFunctionPlot[447].X = -5.53
	// pointsOfFunctionPlot[447].Y = 0.021

	// pointsOfFunctionPlot[448].X = -5.52
	// pointsOfFunctionPlot[448].Y = 0.021

	// pointsOfFunctionPlot[449].X = -5.51
	// pointsOfFunctionPlot[449].Y = 0.021

	// pointsOfFunctionPlot[450].X = -5.5
	// pointsOfFunctionPlot[450].Y = 0.022

	// pointsOfFunctionPlot[451].X = -5.49
	// pointsOfFunctionPlot[451].Y = 0.022

	// pointsOfFunctionPlot[452].X = -5.48
	// pointsOfFunctionPlot[452].Y = 0.022

	// pointsOfFunctionPlot[453].X = -5.47
	// pointsOfFunctionPlot[453].Y = 0.022

	// pointsOfFunctionPlot[454].X = -5.46
	// pointsOfFunctionPlot[454].Y = 0.022

	// pointsOfFunctionPlot[455].X = -5.45
	// pointsOfFunctionPlot[455].Y = 0.022

	// pointsOfFunctionPlot[456].X = -5.44
	// pointsOfFunctionPlot[456].Y = 0.023

	// pointsOfFunctionPlot[457].X = -5.43
	// pointsOfFunctionPlot[457].Y = 0.023

	// pointsOfFunctionPlot[458].X = -5.42
	// pointsOfFunctionPlot[458].Y = 0.023

	// pointsOfFunctionPlot[459].X = -5.41
	// pointsOfFunctionPlot[459].Y = 0.023

	// pointsOfFunctionPlot[460].X = -5.4
	// pointsOfFunctionPlot[460].Y = 0.023

	// pointsOfFunctionPlot[461].X = -5.39
	// pointsOfFunctionPlot[461].Y = 0.023

	// pointsOfFunctionPlot[462].X = -5.38
	// pointsOfFunctionPlot[462].Y = 0.024

	// pointsOfFunctionPlot[463].X = -5.37
	// pointsOfFunctionPlot[463].Y = 0.024

	// pointsOfFunctionPlot[464].X = -5.36
	// pointsOfFunctionPlot[464].Y = 0.024

	// pointsOfFunctionPlot[465].X = -5.35
	// pointsOfFunctionPlot[465].Y = 0.024

	// pointsOfFunctionPlot[466].X = -5.34
	// pointsOfFunctionPlot[466].Y = 0.024

	// pointsOfFunctionPlot[467].X = -5.33
	// pointsOfFunctionPlot[467].Y = 0.024

	// pointsOfFunctionPlot[468].X = -5.32
	// pointsOfFunctionPlot[468].Y = 0.025

	// pointsOfFunctionPlot[469].X = -5.31
	// pointsOfFunctionPlot[469].Y = 0.025

	// pointsOfFunctionPlot[470].X = -5.3
	// pointsOfFunctionPlot[470].Y = 0.025

	// pointsOfFunctionPlot[471].X = -5.29
	// pointsOfFunctionPlot[471].Y = 0.025

	// pointsOfFunctionPlot[472].X = -5.28
	// pointsOfFunctionPlot[472].Y = 0.025

	// pointsOfFunctionPlot[473].X = -5.27
	// pointsOfFunctionPlot[473].Y = 0.025

	// pointsOfFunctionPlot[474].X = -5.26
	// pointsOfFunctionPlot[474].Y = 0.026

	// pointsOfFunctionPlot[475].X = -5.25
	// pointsOfFunctionPlot[475].Y = 0.026

	// pointsOfFunctionPlot[476].X = -5.24
	// pointsOfFunctionPlot[476].Y = 0.026

	// pointsOfFunctionPlot[477].X = -5.23
	// pointsOfFunctionPlot[477].Y = 0.026

	// pointsOfFunctionPlot[478].X = -5.22
	// pointsOfFunctionPlot[478].Y = 0.026

	// pointsOfFunctionPlot[479].X = -5.21
	// pointsOfFunctionPlot[479].Y = 0.027

	// pointsOfFunctionPlot[480].X = -5.2
	// pointsOfFunctionPlot[480].Y = 0.027

	// pointsOfFunctionPlot[481].X = -5.19
	// pointsOfFunctionPlot[481].Y = 0.027

	// pointsOfFunctionPlot[482].X = -5.18
	// pointsOfFunctionPlot[482].Y = 0.027

	// pointsOfFunctionPlot[483].X = -5.17
	// pointsOfFunctionPlot[483].Y = 0.027

	// pointsOfFunctionPlot[484].X = -5.16
	// pointsOfFunctionPlot[484].Y = 0.028

	// pointsOfFunctionPlot[485].X = -5.15
	// pointsOfFunctionPlot[485].Y = 0.028

	// pointsOfFunctionPlot[486].X = -5.14
	// pointsOfFunctionPlot[486].Y = 0.028

	// pointsOfFunctionPlot[487].X = -5.13
	// pointsOfFunctionPlot[487].Y = 0.028

	// pointsOfFunctionPlot[488].X = -5.12
	// pointsOfFunctionPlot[488].Y = 0.028

	// pointsOfFunctionPlot[489].X = -5.11
	// pointsOfFunctionPlot[489].Y = 0.029

	// pointsOfFunctionPlot[490].X = -5.1
	// pointsOfFunctionPlot[490].Y = 0.029

	// pointsOfFunctionPlot[491].X = -5.09
	// pointsOfFunctionPlot[491].Y = 0.029

	// pointsOfFunctionPlot[492].X = -5.08
	// pointsOfFunctionPlot[492].Y = 0.029

	// pointsOfFunctionPlot[493].X = -5.07
	// pointsOfFunctionPlot[493].Y = 0.029

	// pointsOfFunctionPlot[494].X = -5.06
	// pointsOfFunctionPlot[494].Y = 0.03

	// pointsOfFunctionPlot[495].X = -5.05
	// pointsOfFunctionPlot[495].Y = 0.03

	// pointsOfFunctionPlot[496].X = -5.04
	// pointsOfFunctionPlot[496].Y = 0.03

	// pointsOfFunctionPlot[497].X = -5.03
	// pointsOfFunctionPlot[497].Y = 0.03

	// pointsOfFunctionPlot[498].X = -5.02
	// pointsOfFunctionPlot[498].Y = 0.031

	// pointsOfFunctionPlot[499].X = -5.01

	// pointsOfFunctionPlot[1_300].X = 3.0
	// pointsOfFunctionPlot[1_300].Y = 8.0

	// pointsOfFunctionPlot[1_301].X = 3.01
	// pointsOfFunctionPlot[1_301].Y = 8.055

	// pointsOfFunctionPlot[1_302].X = 3.02
	// pointsOfFunctionPlot[1_302].Y = 8.111

	// pointsOfFunctionPlot[1_303].X = 3.03
	// pointsOfFunctionPlot[1_303].Y = 8.168

	// pointsOfFunctionPlot[1_304].X = 3.04
	// pointsOfFunctionPlot[1_304].Y = 8.224

	// pointsOfFunctionPlot[1_305].X = 3.05
	// pointsOfFunctionPlot[1_305].Y = 8.282

	// pointsOfFunctionPlot[1_306].X = 3.06
	// pointsOfFunctionPlot[1_306].Y = 8.339

	// pointsOfFunctionPlot[1_307].X = 3.07
	// pointsOfFunctionPlot[1_307].Y = 8.397

	// pointsOfFunctionPlot[1_308].X = 3.08
	// pointsOfFunctionPlot[1_308].Y = 8.456

	// pointsOfFunctionPlot[1_309].X = 3.09
	// pointsOfFunctionPlot[1_309].Y = 8.515

	// pointsOfFunctionPlot[1_310].X = 3.1
	// pointsOfFunctionPlot[1_310].Y = 8.574

	// pointsOfFunctionPlot[1_311].X = 3.11
	// pointsOfFunctionPlot[1_311].Y = 8.633

	// pointsOfFunctionPlot[1_312].X = 3.12
	// pointsOfFunctionPlot[1_312].Y = 8.693

	// pointsOfFunctionPlot[1_313].X = 3.13
	// pointsOfFunctionPlot[1_313].Y = 8.754

	// pointsOfFunctionPlot[1_314].X = 3.14
	// pointsOfFunctionPlot[1_314].Y = 8.815

	// pointsOfFunctionPlot[1_315].X = 3.15
	// pointsOfFunctionPlot[1_315].Y = 8.876

	// pointsOfFunctionPlot[1_316].X = 3.16
	// pointsOfFunctionPlot[1_316].Y = 8.938

	// pointsOfFunctionPlot[1_317].X = 3.17
	// pointsOfFunctionPlot[1_317].Y = 9.005

	// pointsOfFunctionPlot[1_318].X = 3.18
	// pointsOfFunctionPlot[1_318].Y = 9.063

	// pointsOfFunctionPlot[1_319].X = 3.19
	// pointsOfFunctionPlot[1_319].Y = 9.126

	// pointsOfFunctionPlot[1_320].X = 3.2
	// pointsOfFunctionPlot[1_320].Y = 9.189

	// pointsOfFunctionPlot[1_321].X = 3.21
	// pointsOfFunctionPlot[1_321].Y = 9.253

	// pointsOfFunctionPlot[1_322].X = 3.22
	// pointsOfFunctionPlot[1_322].Y = 9.317

	// pointsOfFunctionPlot[1_323].X = 3.23
	// pointsOfFunctionPlot[1_323].Y = 9.382

	// pointsOfFunctionPlot[1_324].X = 3.24
	// pointsOfFunctionPlot[1_324].Y = 9.447

	// pointsOfFunctionPlot[1_325].X = 3.25
	// pointsOfFunctionPlot[1_325].Y = 9.513

	// pointsOfFunctionPlot[1_326].X = 3.26
	// pointsOfFunctionPlot[1_326].Y = 9.579

	// pointsOfFunctionPlot[1_327].X = 3.27
	// pointsOfFunctionPlot[1_327].Y = 9.646

	// pointsOfFunctionPlot[1_328].X = 3.28
	// pointsOfFunctionPlot[1_328].Y = 9.713

	// pointsOfFunctionPlot[1_329].X = 3.29
	// pointsOfFunctionPlot[1_329].Y = 9.781

	// pointsOfFunctionPlot[1_330].X = 3.3
	// pointsOfFunctionPlot[1_330].Y = 9.849

	// pointsOfFunctionPlot[1_331].X = 3.31
	// pointsOfFunctionPlot[1_331].Y = 9.917

	// pointsOfFunctionPlot[1_332].X = 3.32
	// pointsOfFunctionPlot[1_332].Y = 9.986

	// pointsOfFunctionPlot[1_333].X = 3.33
	// pointsOfFunctionPlot[1_333].Y = 10.056

	// pointsOfFunctionPlot[1_334].X = 3.34
	// pointsOfFunctionPlot[1_334].Y = 10.126

	// pointsOfFunctionPlot[1_335].X = 3.35
	// pointsOfFunctionPlot[1_335].Y = 10.196

	// pointsOfFunctionPlot[1_336].X = 3.36
	// pointsOfFunctionPlot[1_336].Y = 10.267

	// pointsOfFunctionPlot[1_337].X = 3.37
	// pointsOfFunctionPlot[1_337].Y = 10.338

	// pointsOfFunctionPlot[1_338].X = 3.38
	// pointsOfFunctionPlot[1_338].Y = 10.41

	// pointsOfFunctionPlot[1_339].X = 3.39
	// pointsOfFunctionPlot[1_339].Y = 10.483

	// pointsOfFunctionPlot[1_340].X = 3.4
	// pointsOfFunctionPlot[1_340].Y = 10.556

	// pointsOfFunctionPlot[1_341].X = 3.41
	// pointsOfFunctionPlot[1_341].Y = 10.629

	// pointsOfFunctionPlot[1_342].X = 3.42
	// pointsOfFunctionPlot[1_342].Y = 10.703

	// pointsOfFunctionPlot[1_343].X = 3.43
	// pointsOfFunctionPlot[1_343].Y = 10.777

	// pointsOfFunctionPlot[1_344].X = 3.44
	// pointsOfFunctionPlot[1_344].Y = 10.852

	// pointsOfFunctionPlot[1_345].X = 3.45
	// pointsOfFunctionPlot[1_345].Y = 10.928

	// pointsOfFunctionPlot[1_346].X = 3.46
	// pointsOfFunctionPlot[1_346].Y = 11.004

	// pointsOfFunctionPlot[1_347].X = 3.47
	// pointsOfFunctionPlot[1_347].Y = 11.08

	// pointsOfFunctionPlot[1_348].X = 3.48
	// pointsOfFunctionPlot[1_348].Y = 11.157

	// pointsOfFunctionPlot[1_349].X = 3.49
	// pointsOfFunctionPlot[1_349].Y = 11.235

	// pointsOfFunctionPlot[1_350].X = 3.5
	// pointsOfFunctionPlot[1_350].Y = 11.313

	// pointsOfFunctionPlot[1_351].X = 3.51
	// pointsOfFunctionPlot[1_351].Y = 11.393

	// pointsOfFunctionPlot[1_352].X = 3.52
	// pointsOfFunctionPlot[1_352].Y = 11.471

	// pointsOfFunctionPlot[1_353].X = 3.53
	// pointsOfFunctionPlot[1_353].Y = 11.551

	// pointsOfFunctionPlot[1_354].X = 3.54
	// pointsOfFunctionPlot[1_354].Y = 11.631

	// pointsOfFunctionPlot[1_355].X = 3.55
	// pointsOfFunctionPlot[1_355].Y = 11.712

	// pointsOfFunctionPlot[1_356].X = 3.56
	// pointsOfFunctionPlot[1_356].Y = 11.794

	// pointsOfFunctionPlot[1_357].X = 3.57
	// pointsOfFunctionPlot[1_357].Y = 11.876

	// pointsOfFunctionPlot[1_358].X = 3.58
	// pointsOfFunctionPlot[1_358].Y = 11.958

	// pointsOfFunctionPlot[1_359].X = 3.59
	// pointsOfFunctionPlot[1_359].Y = 12.042

	// pointsOfFunctionPlot[1_360].X = 3.6
	// pointsOfFunctionPlot[1_360].Y = 12.125

	// pointsOfFunctionPlot[1_361].X = 3.61
	// pointsOfFunctionPlot[1_361].Y = 12.21

	// pointsOfFunctionPlot[1_362].X = 3.62
	// pointsOfFunctionPlot[1_362].Y = 12.295

	// pointsOfFunctionPlot[1_363].X = 3.63
	// pointsOfFunctionPlot[1_363].Y = 12.38

	// pointsOfFunctionPlot[1_364].X = 3.64
	// pointsOfFunctionPlot[1_364].Y = 12.466

	// pointsOfFunctionPlot[1_365].X = 3.65
	// pointsOfFunctionPlot[1_365].Y = 12.553

	// pointsOfFunctionPlot[1_366].X = 3.66
	// pointsOfFunctionPlot[1_366].Y = 12.64

	// pointsOfFunctionPlot[1_367].X = 3.67
	// pointsOfFunctionPlot[1_367].Y = 12.728

	// pointsOfFunctionPlot[1_368].X = 3.68
	// pointsOfFunctionPlot[1_368].Y = 12.817

	// pointsOfFunctionPlot[1_369].X = 3.69
	// pointsOfFunctionPlot[1_369].Y = 12.906

	// pointsOfFunctionPlot[1_370].X = 3.7
	// pointsOfFunctionPlot[1_370].Y = 12.996

	// pointsOfFunctionPlot[1_371].X = 3.71
	// pointsOfFunctionPlot[1_371].Y = 13.086

	// pointsOfFunctionPlot[1_372].X = 3.72
	// pointsOfFunctionPlot[1_372].Y = 13.177

	// pointsOfFunctionPlot[1_373].X = 3.73
	// pointsOfFunctionPlot[1_373].Y = 13.269

	// pointsOfFunctionPlot[1_374].X = 3.74
	// pointsOfFunctionPlot[1_374].Y = 13.361

	// pointsOfFunctionPlot[1_375].X = 3.75
	// pointsOfFunctionPlot[1_375].Y = 13.454

	// pointsOfFunctionPlot[1_376].X = 3.76
	// pointsOfFunctionPlot[1_376].Y = 13.547

	// pointsOfFunctionPlot[1_377].X = 3.77
	// pointsOfFunctionPlot[1_377].Y = 13.642

	// pointsOfFunctionPlot[1_378].X = 3.78
	// pointsOfFunctionPlot[1_378].Y = 13.737

	// pointsOfFunctionPlot[1_379].X = 3.79
	// pointsOfFunctionPlot[1_379].Y = 13.832

	// pointsOfFunctionPlot[1_380].X = 3.8
	// pointsOfFunctionPlot[1_380].Y = 13.928

	// pointsOfFunctionPlot[1_381].X = 3.81
	// pointsOfFunctionPlot[1_381].Y = 14.025

	// pointsOfFunctionPlot[1_382].X = 3.82
	// pointsOfFunctionPlot[1_382].Y = 14.123

	// pointsOfFunctionPlot[1_383].X = 3.83
	// pointsOfFunctionPlot[1_383].Y = 14.221

	// pointsOfFunctionPlot[1_384].X = 3.84
	// pointsOfFunctionPlot[1_384].Y = 14.32

	// pointsOfFunctionPlot[1_385].X = 3.85
	// pointsOfFunctionPlot[1_385].Y = 14.42

	// pointsOfFunctionPlot[1_386].X = 3.86
	// pointsOfFunctionPlot[1_386].Y = 14.52

	// pointsOfFunctionPlot[1_387].X = 3.87
	// pointsOfFunctionPlot[1_387].Y = 14.621

	// pointsOfFunctionPlot[1_388].X = 3.88
	// pointsOfFunctionPlot[1_388].Y = 14.723

	// pointsOfFunctionPlot[1_389].X = 3.89
	// pointsOfFunctionPlot[1_389].Y = 14.825

	// pointsOfFunctionPlot[1_390].X = 3.9
	// pointsOfFunctionPlot[1_390].Y = 14.928

	// pointsOfFunctionPlot[1_391].X = 3.91
	// pointsOfFunctionPlot[1_391].Y = 15.032

	// pointsOfFunctionPlot[1_392].X = 3.92
	// pointsOfFunctionPlot[1_392].Y = 15.136

	// pointsOfFunctionPlot[1_393].X = 3.93
	// pointsOfFunctionPlot[1_393].Y = 15.242

	// pointsOfFunctionPlot[1_394].X = 3.94
	// pointsOfFunctionPlot[1_394].Y = 15.348

	// pointsOfFunctionPlot[1_395].X = 3.95
	// pointsOfFunctionPlot[1_395].Y = 15.455

	// pointsOfFunctionPlot[1_396].X = 3.96
	// pointsOfFunctionPlot[1_396].Y = 15.562

	// pointsOfFunctionPlot[1_397].X = 3.97
	// pointsOfFunctionPlot[1_397].Y = 15.67

	// pointsOfFunctionPlot[1_398].X = 3.98
	// pointsOfFunctionPlot[1_398].Y = 15.779

	// pointsOfFunctionPlot[1_399].X = 3.99
	// pointsOfFunctionPlot[1_399].Y = 15.889

	// pointsOfFunctionPlot[1_400].X = 4.0
	// pointsOfFunctionPlot[1_400].Y = 16.0

	// pointsOfFunctionPlot[1_401].X = 4.01
	// pointsOfFunctionPlot[1_401].Y = 16.111

	// pointsOfFunctionPlot[1_402].X = 4.02
	// pointsOfFunctionPlot[1_402].Y = 16.223

	// pointsOfFunctionPlot[1_403].X = 4.03
	// pointsOfFunctionPlot[1_403].Y = 16.336

	// pointsOfFunctionPlot[1_404].X = 4.04
	// pointsOfFunctionPlot[1_404].Y = 16.449

	// pointsOfFunctionPlot[1_405].X = 4.05
	// pointsOfFunctionPlot[1_405].Y = 16.564

	// pointsOfFunctionPlot[1_406].X = 4.06
	// pointsOfFunctionPlot[1_406].Y = 16.679

	// pointsOfFunctionPlot[1_407].X = 4.07
	// pointsOfFunctionPlot[1_407].Y = 16.795

	// pointsOfFunctionPlot[1_408].X = 4.08
	// pointsOfFunctionPlot[1_408].Y = 16.912

	// pointsOfFunctionPlot[1_409].X = 4.09
	// pointsOfFunctionPlot[1_409].Y = 17.029

	// pointsOfFunctionPlot[1_410].X = 4.1
	// pointsOfFunctionPlot[1_410].Y = 17.148

	// pointsOfFunctionPlot[1_411].X = 4.11
	// pointsOfFunctionPlot[1_411].Y = 17.267

	// pointsOfFunctionPlot[1_412].X = 4.12
	// pointsOfFunctionPlot[1_412].Y = 17.387

	// pointsOfFunctionPlot[1_413].X = 4.13
	// pointsOfFunctionPlot[1_413].Y = 17.508

	// pointsOfFunctionPlot[1_414].X = 4.14
	// pointsOfFunctionPlot[1_414].Y = 17.63

	// pointsOfFunctionPlot[1_415].X = 4.15
	// pointsOfFunctionPlot[1_415].Y = 17.753

	// pointsOfFunctionPlot[1_416].X = 4.16
	// pointsOfFunctionPlot[1_416].Y = 17.876

	// pointsOfFunctionPlot[1_417].X = 4.17
	// pointsOfFunctionPlot[1_417].Y = 18.0

	// pointsOfFunctionPlot[1_418].X = 4.18
	// pointsOfFunctionPlot[1_418].Y = 18.126

	// pointsOfFunctionPlot[1_419].X = 4.19
	// pointsOfFunctionPlot[1_419].Y = 18.252

	// pointsOfFunctionPlot[1_420].X = 4.2
	// pointsOfFunctionPlot[1_420].Y = 18.379

	// pointsOfFunctionPlot[1_421].X = 4.21
	// pointsOfFunctionPlot[1_421].Y = 18.507

	// pointsOfFunctionPlot[1_422].X = 4.22
	// pointsOfFunctionPlot[1_422].Y = 18.635

	// pointsOfFunctionPlot[1_423].X = 4.23
	// pointsOfFunctionPlot[1_423].Y = 18.765

	// pointsOfFunctionPlot[1_424].X = 4.24
	// pointsOfFunctionPlot[1_424].Y = 18.895

	// pointsOfFunctionPlot[1_425].X = 4.25
	// pointsOfFunctionPlot[1_425].Y = 19.027

	// pointsOfFunctionPlot[1_426].X = 4.26
	// pointsOfFunctionPlot[1_426].Y = 19.159

	// pointsOfFunctionPlot[1_427].X = 4.27
	// pointsOfFunctionPlot[1_427].Y = 19.292

	// pointsOfFunctionPlot[1_428].X = 4.28
	// pointsOfFunctionPlot[1_428].Y = 19.427

	// pointsOfFunctionPlot[1_429].X = 4.29
	// pointsOfFunctionPlot[1_429].Y = 19.562

	// pointsOfFunctionPlot[1_430].X = 4.3
	// pointsOfFunctionPlot[1_430].Y = 19.698

	// pointsOfFunctionPlot[1_431].X = 4.31
	// pointsOfFunctionPlot[1_431].Y = 19.835

	// pointsOfFunctionPlot[1_432].X = 4.32
	// pointsOfFunctionPlot[1_432].Y = 19.973

	// pointsOfFunctionPlot[1_433].X = 4.33
	// pointsOfFunctionPlot[1_433].Y = 20.112

	// pointsOfFunctionPlot[1_434].X = 4.34
	// pointsOfFunctionPlot[1_434].Y = 20.252

	// pointsOfFunctionPlot[1_435].X = 4.35
	// pointsOfFunctionPlot[1_435].Y = 20.393

	// pointsOfFunctionPlot[1_436].X = 4.36
	// pointsOfFunctionPlot[1_436].Y = 20.534

	// pointsOfFunctionPlot[1_437].X = 4.37
	// pointsOfFunctionPlot[1_437].Y = 20.677

	// pointsOfFunctionPlot[1_438].X = 4.38
	// pointsOfFunctionPlot[1_438].Y = 20.821

	// pointsOfFunctionPlot[1_439].X = 4.39
	// pointsOfFunctionPlot[1_439].Y = 20.966

	// pointsOfFunctionPlot[1_440].X = 4.4
	// pointsOfFunctionPlot[1_440].Y = 21.112

	// pointsOfFunctionPlot[1_441].X = 4.41
	// pointsOfFunctionPlot[1_441].Y = 21.259

	// pointsOfFunctionPlot[1_442].X = 4.42
	// pointsOfFunctionPlot[1_442].Y = 21.406

	// pointsOfFunctionPlot[1_443].X = 4.43
	// pointsOfFunctionPlot[1_443].Y = 21.555

	// pointsOfFunctionPlot[1_444].X = 4.44
	// pointsOfFunctionPlot[1_444].Y = 21.705

	// pointsOfFunctionPlot[1_445].X = 4.45
	// pointsOfFunctionPlot[1_445].Y = 21.856

	// pointsOfFunctionPlot[1_446].X = 4.46
	// pointsOfFunctionPlot[1_446].Y = 22.008

	// pointsOfFunctionPlot[1_447].X = 4.47
	// pointsOfFunctionPlot[1_447].Y = 22.161

	// pointsOfFunctionPlot[1_448].X = 4.48
	// pointsOfFunctionPlot[1_448].Y = 22.315

	// pointsOfFunctionPlot[1_449].X = 4.49
	// pointsOfFunctionPlot[1_449].Y = 22.471

	// pointsOfFunctionPlot[1_450].X = 4.5
	// pointsOfFunctionPlot[1_450].Y = 22.627

	// pointsOfFunctionPlot[1_451].X = 4.51
	// pointsOfFunctionPlot[1_451].Y = 22.784

	// pointsOfFunctionPlot[1_452].X = 4.52
	// pointsOfFunctionPlot[1_452].Y = 22.943

	// pointsOfFunctionPlot[1_453].X = 4.53
	// pointsOfFunctionPlot[1_453].Y = 23.102

	// pointsOfFunctionPlot[1_454].X = 4.54
	// pointsOfFunctionPlot[1_454].Y = 23.263

	// pointsOfFunctionPlot[1_455].X = 4.55
	// pointsOfFunctionPlot[1_455].Y = 23.425

	// pointsOfFunctionPlot[1_456].X = 4.56
	// pointsOfFunctionPlot[1_456].Y = 23.588

	// pointsOfFunctionPlot[1_457].X = 4.57
	// pointsOfFunctionPlot[1_457].Y = 23.752

	// pointsOfFunctionPlot[1_458].X = 4.58
	// pointsOfFunctionPlot[1_458].Y = 23.917

	// pointsOfFunctionPlot[1_459].X = 4.59
	// pointsOfFunctionPlot[1_459].Y = 24.083

	// pointsOfFunctionPlot[1_460].X = 4.6
	// pointsOfFunctionPlot[1_460].Y = 24.251

	// pointsOfFunctionPlot[1_461].X = 4.61
	// pointsOfFunctionPlot[1_461].Y = 24.42

	// pointsOfFunctionPlot[1_462].X = 4.62
	// pointsOfFunctionPlot[1_462].Y = 24.59

	// pointsOfFunctionPlot[1_463].X = 4.63
	// pointsOfFunctionPlot[1_463].Y = 24.761

	// pointsOfFunctionPlot[1_464].X = 4.64
	// pointsOfFunctionPlot[1_464].Y = 24.933

	// pointsOfFunctionPlot[1_465].X = 4.65
	// pointsOfFunctionPlot[1_465].Y = 25.106

	// pointsOfFunctionPlot[1_466].X = 4.66
	// pointsOfFunctionPlot[1_466].Y = 25.281

	// pointsOfFunctionPlot[1_467].X = 4.67
	// pointsOfFunctionPlot[1_467].Y = 25.457

	// pointsOfFunctionPlot[1_468].X = 4.68
	// pointsOfFunctionPlot[1_468].Y = 25.634

	// pointsOfFunctionPlot[1_469].X = 4.69
	// pointsOfFunctionPlot[1_469].Y = 25.812

	// pointsOfFunctionPlot[1_470].X = 4.7
	// pointsOfFunctionPlot[1_470].Y = 25.992

	// pointsOfFunctionPlot[1_471].X = 4.71
	// pointsOfFunctionPlot[1_471].Y = 26.172

	// pointsOfFunctionPlot[1_472].X = 4.72
	// pointsOfFunctionPlot[1_472].Y = 26.354

	// pointsOfFunctionPlot[1_473].X = 4.73
	// pointsOfFunctionPlot[1_473].Y = 26.538

	// pointsOfFunctionPlot[1_474].X = 4.74
	// pointsOfFunctionPlot[1_474].Y = 26.722

	// pointsOfFunctionPlot[1_475].X = 4.75
	// pointsOfFunctionPlot[1_475].Y = 26.908

	// pointsOfFunctionPlot[1_476].X = 4.76
	// pointsOfFunctionPlot[1_476].Y = 27.095

	// pointsOfFunctionPlot[1_477].X = 4.77
	// pointsOfFunctionPlot[1_477].Y = 27.284

	// pointsOfFunctionPlot[1_478].X = 4.78
	// pointsOfFunctionPlot[1_478].Y = 27.474

	// pointsOfFunctionPlot[1_479].X = 4.79
	// pointsOfFunctionPlot[1_479].Y = 27.665

	// pointsOfFunctionPlot[1_480].X = 4.8
	// pointsOfFunctionPlot[1_480].Y = 27.857

	// pointsOfFunctionPlot[1_481].X = 4.81
	// pointsOfFunctionPlot[1_481].Y = 28.051

	// pointsOfFunctionPlot[1_482].X = 4.82
	// pointsOfFunctionPlot[1_482].Y = 28.246

	// pointsOfFunctionPlot[1_483].X = 4.83
	// pointsOfFunctionPlot[1_483].Y = 28.443

	// pointsOfFunctionPlot[1_484].X = 4.84
	// pointsOfFunctionPlot[1_484].Y = 28.64

	// pointsOfFunctionPlot[1_485].X = 4.85
	// pointsOfFunctionPlot[1_485].Y = 28.84

	// pointsOfFunctionPlot[1_486].X = 4.86
	// pointsOfFunctionPlot[1_486].Y = 29.04

	// pointsOfFunctionPlot[1_487].X = 4.87
	// pointsOfFunctionPlot[1_487].Y = 29.242

	// pointsOfFunctionPlot[1_488].X = 4.88
	// pointsOfFunctionPlot[1_488].Y = 29.446

	// pointsOfFunctionPlot[1_489].X = 4.89
	// pointsOfFunctionPlot[1_489].Y = 29.65

	// pointsOfFunctionPlot[1_490].X = 4.9
	// pointsOfFunctionPlot[1_490].Y = 29.857

	// pointsOfFunctionPlot[1_491].X = 4.91
	// pointsOfFunctionPlot[1_491].Y = 30.064

	// pointsOfFunctionPlot[1_492].X = 4.92
	// pointsOfFunctionPlot[1_492].Y = 30.273

	// pointsOfFunctionPlot[1_493].X = 4.93
	// pointsOfFunctionPlot[1_493].Y = 30.484

	// pointsOfFunctionPlot[1_494].X = 4.94
	// pointsOfFunctionPlot[1_494].Y = 30.696

	// pointsOfFunctionPlot[1_495].X = 4.95
	// pointsOfFunctionPlot[1_495].Y = 30.91

	// pointsOfFunctionPlot[1_496].X = 4.96
	// pointsOfFunctionPlot[1_496].Y = 31.125

	// pointsOfFunctionPlot[1_497].X = 4.97
	// pointsOfFunctionPlot[1_497].Y = 31.341

	// pointsOfFunctionPlot[1_498].X = 4.98
	// pointsOfFunctionPlot[1_498].Y = 31.559

	// pointsOfFunctionPlot[1_499].X = 4.99
	// pointsOfFunctionPlot[1_499].Y = 31.779

	// pointsOfFunctionPlot[1_500].X = 5.0
	// pointsOfFunctionPlot[1_500].Y = 32.0

	// pointsOfFunctionPlot[1_501].X = 5.01
	// pointsOfFunctionPlot[1_501].Y = 32.222

	// pointsOfFunctionPlot[1_502].X = 5.02
	// pointsOfFunctionPlot[1_502].Y = 32.446

	// pointsOfFunctionPlot[1_503].X = 5.03
	// pointsOfFunctionPlot[1_503].Y = 32.672

	// pointsOfFunctionPlot[1_504].X = 5.04
	// pointsOfFunctionPlot[1_504].Y = 32.899

	// pointsOfFunctionPlot[1_505].X = 5.05
	// pointsOfFunctionPlot[1_505].Y = 33.128

	// pointsOfFunctionPlot[1_506].X = 5.06
	// pointsOfFunctionPlot[1_506].Y = 33.358

	// pointsOfFunctionPlot[1_507].X = 5.07
	// pointsOfFunctionPlot[1_507].Y = 33.59

	// pointsOfFunctionPlot[1_508].X = 5.08
	// pointsOfFunctionPlot[1_508].Y = 33.824

	// pointsOfFunctionPlot[1_509].X = 5.09
	// pointsOfFunctionPlot[1_509].Y = 34.059

	// pointsOfFunctionPlot[1_510].X = 5.1
	// pointsOfFunctionPlot[1_510].Y = 34.296

	// pointsOfFunctionPlot[1_511].X = 5.11
	// pointsOfFunctionPlot[1_511].Y = 34.535

	// pointsOfFunctionPlot[1_512].X = 5.12
	// pointsOfFunctionPlot[1_512].Y = 34.775

	// pointsOfFunctionPlot[1_513].X = 5.13
	// pointsOfFunctionPlot[1_513].Y = 35.017

	// pointsOfFunctionPlot[1_514].X = 5.14
	// pointsOfFunctionPlot[1_514].Y = 35.261

	// pointsOfFunctionPlot[1_515].X = 5.15
	// pointsOfFunctionPlot[1_515].Y = 35.506

	// pointsOfFunctionPlot[1_516].X = 5.16
	// pointsOfFunctionPlot[1_516].Y = 35.753

	// pointsOfFunctionPlot[1_517].X = 5.17
	// pointsOfFunctionPlot[1_517].Y = 36.001

	// pointsOfFunctionPlot[1_518].X = 5.18
	// pointsOfFunctionPlot[1_518].Y = 36.252

	// pointsOfFunctionPlot[1_519].X = 5.19
	// pointsOfFunctionPlot[1_519].Y = 36.504

	// pointsOfFunctionPlot[1_520].X = 5.2
	// pointsOfFunctionPlot[1_520].Y = 36.758

	// pointsOfFunctionPlot[1_521].X = 5.21
	// pointsOfFunctionPlot[1_521].Y = 37.014

	// pointsOfFunctionPlot[1_522].X = 5.22
	// pointsOfFunctionPlot[1_522].Y = 37.271

	// pointsOfFunctionPlot[1_523].X = 5.23
	// pointsOfFunctionPlot[1_523].Y = 37.53

	// pointsOfFunctionPlot[1_524].X = 5.24
	// pointsOfFunctionPlot[1_524].Y = 37.791

	// pointsOfFunctionPlot[1_525].X = 5.25
	// pointsOfFunctionPlot[1_525].Y = 38.054

	// pointsOfFunctionPlot[1_526].X = 5.26
	// pointsOfFunctionPlot[1_526].Y = 38.319

	// pointsOfFunctionPlot[1_527].X = 5.27
	// pointsOfFunctionPlot[1_527].Y = 38.585

	// pointsOfFunctionPlot[1_528].X = 5.28
	// pointsOfFunctionPlot[1_528].Y = 38.854

	// pointsOfFunctionPlot[1_529].X = 5.29
	// pointsOfFunctionPlot[1_529].Y = 39.124

	// pointsOfFunctionPlot[1_530].X = 5.3
	// pointsOfFunctionPlot[1_530].Y = 39.396

	// pointsOfFunctionPlot[1_531].X = 5.31
	// pointsOfFunctionPlot[1_531].Y = 39.67

	// pointsOfFunctionPlot[1_532].X = 5.32
	// pointsOfFunctionPlot[1_532].Y = 39.946

	// pointsOfFunctionPlot[1_533].X = 5.33
	// pointsOfFunctionPlot[1_533].Y = 40.224

	// pointsOfFunctionPlot[1_534].X = 5.34
	// pointsOfFunctionPlot[1_534].Y = 40.504

	// pointsOfFunctionPlot[1_535].X = 5.35
	// pointsOfFunctionPlot[1_535].Y = 40.785

	// pointsOfFunctionPlot[1_536].X = 5.36
	// pointsOfFunctionPlot[1_536].Y = 41.069

	// pointsOfFunctionPlot[1_537].X = 5.37
	// pointsOfFunctionPlot[1_537].Y = 41.355

	// pointsOfFunctionPlot[1_538].X = 5.38
	// pointsOfFunctionPlot[1_538].Y = 41.642

	// pointsOfFunctionPlot[1_539].X = 5.39
	// pointsOfFunctionPlot[1_539].Y = 41.932

	// pointsOfFunctionPlot[1_540].X = 5.4
	// pointsOfFunctionPlot[1_540].Y = 42.224

	// pointsOfFunctionPlot[1_541].X = 5.41
	// pointsOfFunctionPlot[1_541].Y = 42.517

	// pointsOfFunctionPlot[1_542].X = 5.42
	// pointsOfFunctionPlot[1_542].Y = 42.813

	// pointsOfFunctionPlot[1_543].X = 5.43
	// pointsOfFunctionPlot[1_543].Y = 43.111

	// pointsOfFunctionPlot[1_544].X = 5.44
	// pointsOfFunctionPlot[1_544].Y = 43.411

	// pointsOfFunctionPlot[1_545].X = 5.45
	// pointsOfFunctionPlot[1_545].Y = 43.713

	// pointsOfFunctionPlot[1_546].X = 5.46
	// pointsOfFunctionPlot[1_546].Y = 44.017

	// pointsOfFunctionPlot[1_547].X = 5.47
	// pointsOfFunctionPlot[1_547].Y = 44.323

	// pointsOfFunctionPlot[1_548].X = 5.48
	// pointsOfFunctionPlot[1_548].Y = 44.631

	// pointsOfFunctionPlot[1_549].X = 5.49
	// pointsOfFunctionPlot[1_549].Y = 44.942

	// pointsOfFunctionPlot[1_550].X = 5.5
	// pointsOfFunctionPlot[1_550].Y = 45.254

	// pointsOfFunctionPlot[1_551].X = 5.51
	// pointsOfFunctionPlot[1_551].Y = 45.569

	// pointsOfFunctionPlot[1_552].X = 5.52
	// pointsOfFunctionPlot[1_552].Y = 45.886

	// pointsOfFunctionPlot[1_553].X = 5.53
	// pointsOfFunctionPlot[1_553].Y = 46.205

	// pointsOfFunctionPlot[1_554].X = 5.54
	// pointsOfFunctionPlot[1_554].Y = 46.527

	// pointsOfFunctionPlot[1_555].X = 5.55
	// pointsOfFunctionPlot[1_555].Y = 46.85

	// pointsOfFunctionPlot[1_556].X = 5.56
	// pointsOfFunctionPlot[1_556].Y = 47.176

	// pointsOfFunctionPlot[1_557].X = 5.57
	// pointsOfFunctionPlot[1_557].Y = 47.504

	// pointsOfFunctionPlot[1_558].X = 5.58
	// pointsOfFunctionPlot[1_558].Y = 47.835

	// pointsOfFunctionPlot[1_559].X = 5.59
	// pointsOfFunctionPlot[1_559].Y = 48.167

	// pointsOfFunctionPlot[1_560].X = 5.6
	// pointsOfFunctionPlot[1_560].Y = 48.502

	// pointsOfFunctionPlot[1_561].X = 5.61
	// pointsOfFunctionPlot[1_561].Y = 48.84

	// pointsOfFunctionPlot[1_562].X = 5.62
	// pointsOfFunctionPlot[1_562].Y = 49.18

	// pointsOfFunctionPlot[1_563].X = 5.63
	// pointsOfFunctionPlot[1_563].Y = 49.522

	// pointsOfFunctionPlot[1_564].X = 5.64
	// pointsOfFunctionPlot[1_564].Y = 49.866

	// pointsOfFunctionPlot[1_565].X = 5.65
	// pointsOfFunctionPlot[1_565].Y = 50.213

	// pointsOfFunctionPlot[1_566].X = 5.66
	// pointsOfFunctionPlot[1_566].Y = 50.562

	// pointsOfFunctionPlot[1_567].X = 5.67
	// pointsOfFunctionPlot[1_567].Y = 50.914

	// pointsOfFunctionPlot[1_568].X = 5.68
	// pointsOfFunctionPlot[1_568].Y = 51.268

	// pointsOfFunctionPlot[1_569].X = 5.69
	// pointsOfFunctionPlot[1_569].Y = 51.625

	// pointsOfFunctionPlot[1_570].X = 5.7
	// pointsOfFunctionPlot[1_570].Y = 51.984

	// pointsOfFunctionPlot[1_571].X = 5.71
	// pointsOfFunctionPlot[1_571].Y = 52.345

	// pointsOfFunctionPlot[1_572].X = 5.72
	// pointsOfFunctionPlot[1_572].Y = 52.709

	// pointsOfFunctionPlot[1_573].X = 5.73
	// pointsOfFunctionPlot[1_573].Y = 53.076

	// pointsOfFunctionPlot[1_574].X = 5.74
	// pointsOfFunctionPlot[1_574].Y = 53.445

	// pointsOfFunctionPlot[1_575].X = 5.75
	// pointsOfFunctionPlot[1_575].Y = 53.817

	// pointsOfFunctionPlot[1_576].X = 5.76
	// pointsOfFunctionPlot[1_576].Y = 54.191

	// pointsOfFunctionPlot[1_577].X = 5.77
	// pointsOfFunctionPlot[1_577].Y = 54.568

	// pointsOfFunctionPlot[1_578].X = 5.78
	// pointsOfFunctionPlot[1_578].Y = 54.948

	// pointsOfFunctionPlot[1_579].X = 5.79
	// pointsOfFunctionPlot[1_579].Y = 55.33

	// pointsOfFunctionPlot[1_580].X = 5.8
	// pointsOfFunctionPlot[1_580].Y = 55.715

	// pointsOfFunctionPlot[1_581].X = 5.81
	// pointsOfFunctionPlot[1_581].Y = 56.102

	// pointsOfFunctionPlot[1_582].X = 5.82
	// pointsOfFunctionPlot[1_582].Y = 56.493

	// pointsOfFunctionPlot[1_583].X = 5.83
	// pointsOfFunctionPlot[1_583].Y = 56.885

	// pointsOfFunctionPlot[1_584].X = 5.84
	// pointsOfFunctionPlot[1_584].Y = 57.281

	// pointsOfFunctionPlot[1_585].X = 5.85
	// pointsOfFunctionPlot[1_585].Y = 57.68

	// pointsOfFunctionPlot[1_586].X = 5.86
	// pointsOfFunctionPlot[1_586].Y = 58.081

	// pointsOfFunctionPlot[1_587].X = 5.87
	// pointsOfFunctionPlot[1_587].Y = 58.485

	// pointsOfFunctionPlot[1_588].X = 5.88
	// pointsOfFunctionPlot[1_588].Y = 58.892

	// pointsOfFunctionPlot[1_589].X = 5.89
	// pointsOfFunctionPlot[1_589].Y = 59.301

	// pointsOfFunctionPlot[1_590].X = 5.9
	// pointsOfFunctionPlot[1_590].Y = 59.714

	// pointsOfFunctionPlot[1_591].X = 5.91
	// pointsOfFunctionPlot[1_591].Y = 60.129

	// pointsOfFunctionPlot[1_592].X = 5.92
	// pointsOfFunctionPlot[1_592].Y = 60.547

	// pointsOfFunctionPlot[1_593].X = 5.93
	// pointsOfFunctionPlot[1_593].Y = 60.968

	// pointsOfFunctionPlot[1_594].X = 5.94
	// pointsOfFunctionPlot[1_594].Y = 61.392

	// pointsOfFunctionPlot[1_595].X = 5.95
	// pointsOfFunctionPlot[1_595].Y = 61.819

	// pointsOfFunctionPlot[1_596].X = 5.96
	// pointsOfFunctionPlot[1_596].Y = 62.249

	// pointsOfFunctionPlot[1_597].X = 5.97
	// pointsOfFunctionPlot[1_597].Y = 62.682

	// pointsOfFunctionPlot[1_598].X = 5.98
	// pointsOfFunctionPlot[1_598].Y = 63.118












	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function tan(x)"

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
		"tan-function-plot-01.png"); err != nil {

		panic(err)
	}
}
