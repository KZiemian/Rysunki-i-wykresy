package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres funkcji f(x) = cos(x)

	pointsOfFunctionPlot := make(plotter.XYs, 321)

	pointsOfFunctionPlot[0].X = 0.0
	pointsOfFunctionPlot[0].Y = 0.0

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = 0.01

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = 0.021

	pointsOfFunctionPlot[3].X = 0.03
	pointsOfFunctionPlot[3].Y = 0.032

	pointsOfFunctionPlot[4].X = 0.04
	pointsOfFunctionPlot[4].Y = 0.043

	pointsOfFunctionPlot[5].X = 0.05
	pointsOfFunctionPlot[5].Y = 0.056

	pointsOfFunctionPlot[6].X = 0.06
	pointsOfFunctionPlot[6].Y = 0.068

	pointsOfFunctionPlot[7].X = 0.07
	pointsOfFunctionPlot[7].Y = 0.081

	pointsOfFunctionPlot[8].X = 0.08
	pointsOfFunctionPlot[8].Y = 0.095

	pointsOfFunctionPlot[9].X = 0.09
	pointsOfFunctionPlot[9].Y = 0.109

	pointsOfFunctionPlot[10].X = 0.1
	pointsOfFunctionPlot[10].Y = 0.124

	pointsOfFunctionPlot[11].X = 0.11
	pointsOfFunctionPlot[11].Y = 0.138

	pointsOfFunctionPlot[12].X = 0.12
	pointsOfFunctionPlot[12].Y = 0.154

	pointsOfFunctionPlot[13].X = 0.13
	pointsOfFunctionPlot[13].Y = 0.17

	pointsOfFunctionPlot[14].X = 0.14
	pointsOfFunctionPlot[14].Y = 0.186

	pointsOfFunctionPlot[15].X = 0.15
	pointsOfFunctionPlot[15].Y = 0.202

	pointsOfFunctionPlot[16].X = 0.16
	pointsOfFunctionPlot[16].Y = 0.219

	pointsOfFunctionPlot[17].X = 0.17
	pointsOfFunctionPlot[17].Y = 0.237

	pointsOfFunctionPlot[18].X = 0.18
	pointsOfFunctionPlot[18].Y = 0.255

	pointsOfFunctionPlot[19].X = 0.19
	pointsOfFunctionPlot[19].Y = 0.273

	pointsOfFunctionPlot[20].X = 0.2
	pointsOfFunctionPlot[20].Y = 0.292

	pointsOfFunctionPlot[21].X = 0.21
	pointsOfFunctionPlot[21].Y = 0.311

	pointsOfFunctionPlot[22].X = 0.22
	pointsOfFunctionPlot[22].Y = 0.33

	pointsOfFunctionPlot[23].X = 0.23
	pointsOfFunctionPlot[23].Y = 0.35

	pointsOfFunctionPlot[24].X = 0.24
	pointsOfFunctionPlot[24].Y = 0.37

	pointsOfFunctionPlot[25].X = 0.25
	pointsOfFunctionPlot[25].Y = 0.39

	pointsOfFunctionPlot[26].X = 0.26
	pointsOfFunctionPlot[26].Y = 0.411

	pointsOfFunctionPlot[27].X = 0.27
	pointsOfFunctionPlot[27].Y = 0.432

	pointsOfFunctionPlot[28].X = 0.28
	pointsOfFunctionPlot[28].Y = 0.454

	pointsOfFunctionPlot[29].X = 0.29
	pointsOfFunctionPlot[29].Y = 0.476

	pointsOfFunctionPlot[30].X = 0.3
	pointsOfFunctionPlot[30].Y = 0.498

	pointsOfFunctionPlot[31].X = 0.31
	pointsOfFunctionPlot[31].Y = 0.52

	pointsOfFunctionPlot[32].X = 0.32
	pointsOfFunctionPlot[32].Y = 0.543

	pointsOfFunctionPlot[33].X = 0.33
	pointsOfFunctionPlot[33].Y = 0.566

	pointsOfFunctionPlot[34].X = 0.34
	pointsOfFunctionPlot[34].Y = 0.59

	pointsOfFunctionPlot[35].X = 0.35
	pointsOfFunctionPlot[35].Y = 0.614

	pointsOfFunctionPlot[36].X = 0.36
	pointsOfFunctionPlot[36].Y = 0.638

	pointsOfFunctionPlot[37].X = 0.37
	pointsOfFunctionPlot[37].Y = 0.662

	pointsOfFunctionPlot[38].X = 0.38
	pointsOfFunctionPlot[38].Y = 0.687

	pointsOfFunctionPlot[39].X = 0.39
	pointsOfFunctionPlot[39].Y = 0.711

	pointsOfFunctionPlot[40].X = 0.40
	pointsOfFunctionPlot[40].Y = 0.737

	pointsOfFunctionPlot[41].X = 0.41
	pointsOfFunctionPlot[41].Y = 0.762

	pointsOfFunctionPlot[42].X = 0.42
	pointsOfFunctionPlot[42].Y = 0.788

	pointsOfFunctionPlot[43].X = 0.43
	pointsOfFunctionPlot[43].Y = 0.814

	pointsOfFunctionPlot[44].X = 0.44
	pointsOfFunctionPlot[44].Y = 0.84

	pointsOfFunctionPlot[45].X = 0.45
	pointsOfFunctionPlot[45].Y = 0.866

	pointsOfFunctionPlot[46].X = 0.46
	pointsOfFunctionPlot[46].Y = 0.893

	pointsOfFunctionPlot[47].X = 0.47
	pointsOfFunctionPlot[47].Y = 0.92

	pointsOfFunctionPlot[48].X = 0.48
	pointsOfFunctionPlot[48].Y = 0.947

	pointsOfFunctionPlot[49].X = 0.49
	pointsOfFunctionPlot[49].Y = 0.975

	pointsOfFunctionPlot[50].X = 0.5
	pointsOfFunctionPlot[50].Y = 1.002

	pointsOfFunctionPlot[51].X = 0.51
	pointsOfFunctionPlot[51].Y = 1.03

	pointsOfFunctionPlot[52].X = 0.52
	pointsOfFunctionPlot[52].Y = 1.058

	pointsOfFunctionPlot[53].X = 0.53
	pointsOfFunctionPlot[53].Y = 1.086

	pointsOfFunctionPlot[54].X = 0.54
	pointsOfFunctionPlot[54].Y = 1.115

	pointsOfFunctionPlot[55].X = 0.55
	pointsOfFunctionPlot[55].Y = 1.143

	pointsOfFunctionPlot[56].X = 0.56
	pointsOfFunctionPlot[56].Y = 1.172

	pointsOfFunctionPlot[57].X = 0.57
	pointsOfFunctionPlot[57].Y = 1.201

	pointsOfFunctionPlot[58].X = 0.58
	pointsOfFunctionPlot[58].Y = 1.23

	pointsOfFunctionPlot[59].X = 0.59
	pointsOfFunctionPlot[59].Y = 1.26

	pointsOfFunctionPlot[60].X = 0.6
	pointsOfFunctionPlot[60].Y = 1.289

	pointsOfFunctionPlot[61].X = 0.61
	pointsOfFunctionPlot[61].Y = 1.319

	pointsOfFunctionPlot[62].X = 0.62
	pointsOfFunctionPlot[62].Y = 1.349

	pointsOfFunctionPlot[63].X = 0.63
	pointsOfFunctionPlot[63].Y = 1.379

	pointsOfFunctionPlot[64].X = 0.64
	pointsOfFunctionPlot[64].Y = 1.409

	pointsOfFunctionPlot[65].X = 0.65
	pointsOfFunctionPlot[65].Y = 1.439

	pointsOfFunctionPlot[66].X = 0.66
	pointsOfFunctionPlot[66].Y = 1.47

	pointsOfFunctionPlot[67].X = 0.67
	pointsOfFunctionPlot[67].Y = 1.5

	pointsOfFunctionPlot[68].X = 0.68
	pointsOfFunctionPlot[68].Y = 1.531

	pointsOfFunctionPlot[69].X = 0.69
	pointsOfFunctionPlot[69].Y = 1.561

	pointsOfFunctionPlot[70].X = 0.7
	pointsOfFunctionPlot[70].Y = 1.592

	pointsOfFunctionPlot[71].X = 0.71
	pointsOfFunctionPlot[71].Y = 1.623

	pointsOfFunctionPlot[72].X = 0.72
	pointsOfFunctionPlot[72].Y = 1.654

	pointsOfFunctionPlot[73].X = 0.73
	pointsOfFunctionPlot[73].Y = 1.686

	pointsOfFunctionPlot[74].X = 0.74
	pointsOfFunctionPlot[74].Y = 1.717

	pointsOfFunctionPlot[75].X = 0.75
	pointsOfFunctionPlot[75].Y = 1.748

	pointsOfFunctionPlot[76].X = 0.76
	pointsOfFunctionPlot[76].Y = 1.78

	pointsOfFunctionPlot[77].X = 0.77
	pointsOfFunctionPlot[77].Y = 1.811

	pointsOfFunctionPlot[78].X = 0.78
	pointsOfFunctionPlot[78].Y = 1.843

	pointsOfFunctionPlot[79].X = 0.79
	pointsOfFunctionPlot[79].Y = 1.875

	pointsOfFunctionPlot[80].X = 0.8
	pointsOfFunctionPlot[80].Y = 1.906

	pointsOfFunctionPlot[81].X = 0.81
	pointsOfFunctionPlot[81].Y = 1.938

	pointsOfFunctionPlot[82].X = 0.82
	pointsOfFunctionPlot[82].Y = 1.97

	pointsOfFunctionPlot[83].X = 0.83
	pointsOfFunctionPlot[83].Y = 2.002

	pointsOfFunctionPlot[84].X = 0.84
	pointsOfFunctionPlot[84].Y = 2.034

	pointsOfFunctionPlot[85].X = 0.85
	pointsOfFunctionPlot[85].Y = 2.066

	pointsOfFunctionPlot[86].X = 0.86
	pointsOfFunctionPlot[86].Y = 2.098

	pointsOfFunctionPlot[87].X = 0.87
	pointsOfFunctionPlot[87].Y = 2.13

	pointsOfFunctionPlot[88].X = 0.88
	pointsOfFunctionPlot[88].Y = 2.162

	pointsOfFunctionPlot[89].X = 0.89
	pointsOfFunctionPlot[89].Y = 2.194

	pointsOfFunctionPlot[90].X = 0.9
	pointsOfFunctionPlot[90].Y = 2.227

	pointsOfFunctionPlot[91].X = 0.91
	pointsOfFunctionPlot[91].Y = 2.259

	pointsOfFunctionPlot[92].X = 0.92
	pointsOfFunctionPlot[92].Y = 2.291

	pointsOfFunctionPlot[93].X = 0.93
	pointsOfFunctionPlot[93].Y = 2.323

	pointsOfFunctionPlot[94].X = 0.94
	pointsOfFunctionPlot[94].Y = 2.355

	pointsOfFunctionPlot[95].X = 0.95
	pointsOfFunctionPlot[95].Y = 2.387

	pointsOfFunctionPlot[96].X = 0.96
	pointsOfFunctionPlot[96].Y = 2.42

	pointsOfFunctionPlot[97].X = 0.97
	pointsOfFunctionPlot[97].Y = 2.452

	pointsOfFunctionPlot[98].X = 0.98
	pointsOfFunctionPlot[98].Y = 2.484

	pointsOfFunctionPlot[99].X = 0.99
	pointsOfFunctionPlot[99].Y = 2.516

	pointsOfFunctionPlot[100].X = 1.0
	pointsOfFunctionPlot[100].Y = 2.548

	pointsOfFunctionPlot[101].X = 1.01
	pointsOfFunctionPlot[101].Y = 2.58

	pointsOfFunctionPlot[102].X = 1.02
	pointsOfFunctionPlot[102].Y = 2.612

	pointsOfFunctionPlot[103].X = 1.03
	pointsOfFunctionPlot[103].Y = 2.644

	pointsOfFunctionPlot[104].X = 1.04
	pointsOfFunctionPlot[104].Y = 2.676

	pointsOfFunctionPlot[105].X = 1.05
	pointsOfFunctionPlot[105].Y = 2.708

	pointsOfFunctionPlot[106].X = 1.06
	pointsOfFunctionPlot[106].Y = 2.74

	pointsOfFunctionPlot[107].X = 1.07
	pointsOfFunctionPlot[107].Y = 2.772

	pointsOfFunctionPlot[108].X = 1.08
	pointsOfFunctionPlot[108].Y = 2.803

	pointsOfFunctionPlot[109].X = 1.09
	pointsOfFunctionPlot[109].Y = 2.835

	pointsOfFunctionPlot[110].X = 1.1
	pointsOfFunctionPlot[110].Y = 2.867

	pointsOfFunctionPlot[111].X = 1.11
	pointsOfFunctionPlot[111].Y = 2.898

	pointsOfFunctionPlot[112].X = 1.12
	pointsOfFunctionPlot[112].Y = 2.93

	pointsOfFunctionPlot[113].X = 1.13
	pointsOfFunctionPlot[113].Y = 2.961

	pointsOfFunctionPlot[114].X = 1.14
	pointsOfFunctionPlot[114].Y = 2.992

	pointsOfFunctionPlot[115].X = 1.15
	pointsOfFunctionPlot[115].Y = 3.023

	pointsOfFunctionPlot[116].X = 1.16
	pointsOfFunctionPlot[116].Y = 3.055

	pointsOfFunctionPlot[117].X = 1.17
	pointsOfFunctionPlot[117].Y = 3.086

	pointsOfFunctionPlot[118].X = 1.18
	pointsOfFunctionPlot[118].Y = 3.117

	pointsOfFunctionPlot[119].X = 1.19
	pointsOfFunctionPlot[119].Y = 3.147

	pointsOfFunctionPlot[120].X = 1.2
	pointsOfFunctionPlot[120].Y = 3.178

	pointsOfFunctionPlot[121].X = 1.21
	pointsOfFunctionPlot[121].Y = 3.209

	pointsOfFunctionPlot[122].X = 1.22
	pointsOfFunctionPlot[122].Y = 3.239

	pointsOfFunctionPlot[123].X = 1.23
	pointsOfFunctionPlot[123].Y = 3.27

	pointsOfFunctionPlot[124].X = 1.24
	pointsOfFunctionPlot[124].Y = 3.3

	pointsOfFunctionPlot[125].X = 1.25
	pointsOfFunctionPlot[125].Y = 3.33

	pointsOfFunctionPlot[126].X = 1.26
	pointsOfFunctionPlot[126].Y = 3.36

	pointsOfFunctionPlot[127].X = 1.27
	pointsOfFunctionPlot[127].Y = 3.39

	pointsOfFunctionPlot[128].X = 1.28
	pointsOfFunctionPlot[128].Y = 3.42

	pointsOfFunctionPlot[129].X = 1.29
	pointsOfFunctionPlot[129].Y = 3.45

	pointsOfFunctionPlot[130].X = 1.3
	pointsOfFunctionPlot[130].Y = 3.479

	pointsOfFunctionPlot[131].X = 1.31
	pointsOfFunctionPlot[131].Y = 3.508

	pointsOfFunctionPlot[132].X = 1.32
	pointsOfFunctionPlot[132].Y = 3.538

	pointsOfFunctionPlot[133].X = 1.33
	pointsOfFunctionPlot[133].Y = 3.567

	pointsOfFunctionPlot[134].X = 1.34
	pointsOfFunctionPlot[134].Y = 3.596

	pointsOfFunctionPlot[135].X = 1.35
	pointsOfFunctionPlot[135].Y = 3.625

	pointsOfFunctionPlot[136].X = 1.36
	pointsOfFunctionPlot[136].Y = 3.653

	pointsOfFunctionPlot[137].X = 1.37
	pointsOfFunctionPlot[137].Y = 3.682

	pointsOfFunctionPlot[138].X = 1.38
	pointsOfFunctionPlot[138].Y = 3.71

	pointsOfFunctionPlot[139].X = 1.39
	pointsOfFunctionPlot[139].Y = 3.738

	pointsOfFunctionPlot[140].X = 1.4
	pointsOfFunctionPlot[140].Y = 3.766

	pointsOfFunctionPlot[141].X = 1.41
	pointsOfFunctionPlot[141].Y = 3.794

	pointsOfFunctionPlot[142].X = 1.42
	pointsOfFunctionPlot[142].Y = 3.822

	pointsOfFunctionPlot[143].X = 1.43
	pointsOfFunctionPlot[143].Y = 3.849

	pointsOfFunctionPlot[144].X = 1.44
	pointsOfFunctionPlot[144].Y = 3.876

	pointsOfFunctionPlot[145].X = 1.45
	pointsOfFunctionPlot[145].Y = 3.904

	pointsOfFunctionPlot[146].X = 1.46
	pointsOfFunctionPlot[146].Y = 3.93

	pointsOfFunctionPlot[147].X = 1.47
	pointsOfFunctionPlot[147].Y = 3.957

	pointsOfFunctionPlot[148].X = 1.48
	pointsOfFunctionPlot[148].Y = 3.984

	pointsOfFunctionPlot[149].X = 1.49
	pointsOfFunctionPlot[149].Y = 4.01

	pointsOfFunctionPlot[150].X = 1.5
	pointsOfFunctionPlot[150].Y = 0.0

	pointsOfFunctionPlot[151].X = 1.51
	pointsOfFunctionPlot[151].Y = 3.91

	pointsOfFunctionPlot[152].X = 1.52
	pointsOfFunctionPlot[152].Y = 3.924

	pointsOfFunctionPlot[153].X = 1.53
	pointsOfFunctionPlot[153].Y = 3.938

	pointsOfFunctionPlot[154].X = 1.54
	pointsOfFunctionPlot[154].Y = 3.951

	pointsOfFunctionPlot[155].X = 1.55
	pointsOfFunctionPlot[155].Y = 3.963

	pointsOfFunctionPlot[156].X = 1.56
	pointsOfFunctionPlot[156].Y = 3.975

	pointsOfFunctionPlot[157].X = 1.57
	pointsOfFunctionPlot[157].Y = 3.986

	pointsOfFunctionPlot[158].X = 1.58
	pointsOfFunctionPlot[158].Y = 3.996

	pointsOfFunctionPlot[159].X = 1.59
	pointsOfFunctionPlot[159].Y = 4.006

	pointsOfFunctionPlot[160].X = 1.6
	pointsOfFunctionPlot[160].Y = 4.015

	pointsOfFunctionPlot[161].X = 1.61
	pointsOfFunctionPlot[161].Y = 4.023

	pointsOfFunctionPlot[162].X = 1.62
	pointsOfFunctionPlot[162].Y = 4.031

	pointsOfFunctionPlot[163].X = 1.63
	pointsOfFunctionPlot[163].Y = 4.038

	pointsOfFunctionPlot[164].X = 1.64
	pointsOfFunctionPlot[164].Y = 4.045

	pointsOfFunctionPlot[165].X = 1.65
	pointsOfFunctionPlot[165].Y = 4.05

	pointsOfFunctionPlot[166].X = 1.66
	pointsOfFunctionPlot[166].Y = 4.055

	pointsOfFunctionPlot[167].X = 1.67
	pointsOfFunctionPlot[167].Y = 4.06

	pointsOfFunctionPlot[168].X = 1.68
	pointsOfFunctionPlot[168].Y = 4.063

	pointsOfFunctionPlot[169].X = 1.69
	pointsOfFunctionPlot[169].Y = 4.066

	pointsOfFunctionPlot[170].X = 1.7
	pointsOfFunctionPlot[170].Y = 4.068

	pointsOfFunctionPlot[171].X = 1.71
	pointsOfFunctionPlot[171].Y = 4.07

	pointsOfFunctionPlot[172].X = 1.72
	pointsOfFunctionPlot[172].Y = 4.071

	pointsOfFunctionPlot[173].X = 1.73
	pointsOfFunctionPlot[173].Y = 4.07

	pointsOfFunctionPlot[174].X = 1.74
	pointsOfFunctionPlot[174].Y = 4.07

	pointsOfFunctionPlot[175].X = 1.75
	pointsOfFunctionPlot[175].Y = 4.068

	pointsOfFunctionPlot[176].X = 1.76
	pointsOfFunctionPlot[176].Y = 4.065

	pointsOfFunctionPlot[177].X = 1.77
	pointsOfFunctionPlot[177].Y = 4.062

	pointsOfFunctionPlot[178].X = 1.78
	pointsOfFunctionPlot[178].Y = 4.058

	pointsOfFunctionPlot[179].X = 1.79
	pointsOfFunctionPlot[179].Y = 4.053

	pointsOfFunctionPlot[180].X = 1.8
	pointsOfFunctionPlot[180].Y = 4.048

	pointsOfFunctionPlot[181].X = 1.81
	pointsOfFunctionPlot[181].Y = 4.041

	pointsOfFunctionPlot[182].X = 1.82
	pointsOfFunctionPlot[182].Y = 4.034

	pointsOfFunctionPlot[183].X = 1.83
	pointsOfFunctionPlot[183].Y = 4.025

	pointsOfFunctionPlot[184].X = 1.84
	pointsOfFunctionPlot[184].Y = 4.016

	pointsOfFunctionPlot[185].X = 1.85
	pointsOfFunctionPlot[185].Y = 4.006

	pointsOfFunctionPlot[186].X = 1.86
	pointsOfFunctionPlot[186].Y = 3.996

	pointsOfFunctionPlot[187].X = 1.87
	pointsOfFunctionPlot[187].Y = 3.984

	pointsOfFunctionPlot[188].X = 1.88
	pointsOfFunctionPlot[188].Y = 3.971

	pointsOfFunctionPlot[189].X = 1.89
	pointsOfFunctionPlot[189].Y = 3.958

	pointsOfFunctionPlot[190].X = 1.9
	pointsOfFunctionPlot[190].Y = 3.943

	pointsOfFunctionPlot[191].X = 1.91
	pointsOfFunctionPlot[191].Y = 3.928

	pointsOfFunctionPlot[192].X = 1.92
	pointsOfFunctionPlot[192].Y = 3.912

	pointsOfFunctionPlot[193].X = 1.93
	pointsOfFunctionPlot[193].Y = 3.894

	pointsOfFunctionPlot[194].X = 1.94
	pointsOfFunctionPlot[194].Y = 3.876

	pointsOfFunctionPlot[195].X = 1.95
	pointsOfFunctionPlot[195].Y = 3.857

	pointsOfFunctionPlot[196].X = 1.96
	pointsOfFunctionPlot[196].Y = 3.837

	pointsOfFunctionPlot[197].X = 1.97
	pointsOfFunctionPlot[197].Y = 3.816

	pointsOfFunctionPlot[198].X = 1.98
	pointsOfFunctionPlot[198].Y = 3.794

	pointsOfFunctionPlot[199].X = 1.99
	pointsOfFunctionPlot[199].Y = 3.771

	pointsOfFunctionPlot[200].X = 2.0
	pointsOfFunctionPlot[200].Y = 3.747

	pointsOfFunctionPlot[201].X = 2.01
	pointsOfFunctionPlot[201].Y = 3.722

	pointsOfFunctionPlot[202].X = 2.02
	pointsOfFunctionPlot[202].Y = 3.696

	pointsOfFunctionPlot[203].X = 2.03
	pointsOfFunctionPlot[203].Y = 3.669

	pointsOfFunctionPlot[204].X = 2.04
	pointsOfFunctionPlot[204].Y = 3.641

	pointsOfFunctionPlot[205].X = 2.05
	pointsOfFunctionPlot[205].Y = 3.612

	pointsOfFunctionPlot[206].X = 2.06
	pointsOfFunctionPlot[206].Y = 3.581

	pointsOfFunctionPlot[207].X = 2.07
	pointsOfFunctionPlot[207].Y = 3.55

	pointsOfFunctionPlot[208].X = 2.08
	pointsOfFunctionPlot[208].Y = 3.518

	pointsOfFunctionPlot[209].X = 2.09
	pointsOfFunctionPlot[209].Y = 3.484

	pointsOfFunctionPlot[210].X = 2.1
	pointsOfFunctionPlot[210].Y = 3.45

	pointsOfFunctionPlot[211].X = 2.11
	pointsOfFunctionPlot[211].Y = 3.414

	pointsOfFunctionPlot[212].X = 2.12
	pointsOfFunctionPlot[212].Y = 3.378

	pointsOfFunctionPlot[213].X = 2.13
	pointsOfFunctionPlot[213].Y = 3.34

	pointsOfFunctionPlot[214].X = 2.14
	pointsOfFunctionPlot[214].Y = 3.301

	pointsOfFunctionPlot[215].X = 2.15
	pointsOfFunctionPlot[215].Y = 3.261

	pointsOfFunctionPlot[216].X = 2.16
	pointsOfFunctionPlot[216].Y = 3.22

	pointsOfFunctionPlot[217].X = 2.17
	pointsOfFunctionPlot[217].Y = 3.177

	pointsOfFunctionPlot[218].X = 2.18
	pointsOfFunctionPlot[218].Y = 3.134

	pointsOfFunctionPlot[219].X = 2.19
	pointsOfFunctionPlot[219].Y = 3.089

	pointsOfFunctionPlot[220].X = 2.2
	pointsOfFunctionPlot[220].Y = 3.043

	pointsOfFunctionPlot[221].X = 2.21
	pointsOfFunctionPlot[221].Y = 2.996

	pointsOfFunctionPlot[222].X = 2.22
	pointsOfFunctionPlot[222].Y = 2.948

	pointsOfFunctionPlot[223].X = 2.23
	pointsOfFunctionPlot[223].Y = 2.899

	pointsOfFunctionPlot[224].X = 2.24
	pointsOfFunctionPlot[224].Y = 2.848

	pointsOfFunctionPlot[225].X = 2.25
	pointsOfFunctionPlot[225].Y = 2.797

	pointsOfFunctionPlot[226].X = 2.26
	pointsOfFunctionPlot[226].Y = 2.744

	pointsOfFunctionPlot[227].X = 2.27
	pointsOfFunctionPlot[227].Y = 2.69

	pointsOfFunctionPlot[228].X = 2.28
	pointsOfFunctionPlot[228].Y = 2.634

	pointsOfFunctionPlot[229].X = 2.29
	pointsOfFunctionPlot[229].Y = 2.577

	pointsOfFunctionPlot[230].X = 2.3
	pointsOfFunctionPlot[230].Y = 2.52

	pointsOfFunctionPlot[231].X = 2.31
	pointsOfFunctionPlot[231].Y = 2.46

	pointsOfFunctionPlot[232].X = 2.32
	pointsOfFunctionPlot[232].Y = 2.4

	pointsOfFunctionPlot[233].X = 2.33
	pointsOfFunctionPlot[233].Y = 2.338

	pointsOfFunctionPlot[234].X = 2.34
	pointsOfFunctionPlot[234].Y = 2.275

	pointsOfFunctionPlot[235].X = 2.35
	pointsOfFunctionPlot[235].Y = 2.211

	pointsOfFunctionPlot[236].X = 2.36
	pointsOfFunctionPlot[236].Y = 2.146

	pointsOfFunctionPlot[237].X = 2.37
	pointsOfFunctionPlot[237].Y = 2.079

	pointsOfFunctionPlot[238].X = 2.38
	pointsOfFunctionPlot[238].Y = 2.011

	pointsOfFunctionPlot[239].X = 2.39
	pointsOfFunctionPlot[239].Y = 1.941

	pointsOfFunctionPlot[240].X = 2.4
	pointsOfFunctionPlot[240].Y = 1.871

	pointsOfFunctionPlot[241].X = 2.41
	pointsOfFunctionPlot[241].Y = 1.798

	pointsOfFunctionPlot[242].X = 2.42
	pointsOfFunctionPlot[242].Y = 1.725

	pointsOfFunctionPlot[243].X = 2.43
	pointsOfFunctionPlot[243].Y = 1.65

	pointsOfFunctionPlot[244].X = 2.44
	pointsOfFunctionPlot[244].Y = 1.574

	pointsOfFunctionPlot[245].X = 2.45
	pointsOfFunctionPlot[245].Y = 1.497

	pointsOfFunctionPlot[246].X = 2.46
	pointsOfFunctionPlot[246].Y = 1.418

	pointsOfFunctionPlot[247].X = 2.47
	pointsOfFunctionPlot[247].Y = 1.338

	pointsOfFunctionPlot[248].X = 2.48
	pointsOfFunctionPlot[248].Y = 1.256

	pointsOfFunctionPlot[249].X = 2.49
	pointsOfFunctionPlot[249].Y = 1.173

	pointsOfFunctionPlot[250].X = 2.5
	pointsOfFunctionPlot[250].Y = 1.089

	pointsOfFunctionPlot[251].X = 2.51
	pointsOfFunctionPlot[251].Y = 1.003

	pointsOfFunctionPlot[252].X = 2.52
	pointsOfFunctionPlot[252].Y = 0.916

	pointsOfFunctionPlot[253].X = 2.53
	pointsOfFunctionPlot[253].Y = 0.827

	pointsOfFunctionPlot[254].X = 2.54
	pointsOfFunctionPlot[254].Y = 0.737

	pointsOfFunctionPlot[255].X = 2.55
	pointsOfFunctionPlot[255].Y = 0.646

	pointsOfFunctionPlot[256].X = 2.56
	pointsOfFunctionPlot[256].Y = 0.553

	pointsOfFunctionPlot[257].X = 2.57
	pointsOfFunctionPlot[257].Y = 0.458

	pointsOfFunctionPlot[258].X = 2.58
	pointsOfFunctionPlot[258].Y = 0.363

	pointsOfFunctionPlot[259].X = 2.59
	pointsOfFunctionPlot[259].Y = 0.265

	pointsOfFunctionPlot[260].X = 2.6
	pointsOfFunctionPlot[260].Y = 0.167

	pointsOfFunctionPlot[261].X = 2.61
	pointsOfFunctionPlot[261].Y = 0.067

	pointsOfFunctionPlot[262].X = 2.62
	pointsOfFunctionPlot[262].Y = -0.034

	pointsOfFunctionPlot[263].X = 2.63
	pointsOfFunctionPlot[263].Y = -0.137

	pointsOfFunctionPlot[264].X = 2.64
	pointsOfFunctionPlot[264].Y = -0.242

	pointsOfFunctionPlot[265].X = 2.65
	pointsOfFunctionPlot[265].Y = -0.348

	pointsOfFunctionPlot[266].X = 2.66
	pointsOfFunctionPlot[266].Y = -0.456

	pointsOfFunctionPlot[267].X = 2.67
	pointsOfFunctionPlot[267].Y = -0.565

	pointsOfFunctionPlot[268].X = 2.68
	pointsOfFunctionPlot[268].Y = -0.675

	pointsOfFunctionPlot[269].X = 2.69
	pointsOfFunctionPlot[269].Y = -0.788

	pointsOfFunctionPlot[270].X = 2.7
	pointsOfFunctionPlot[270].Y = -0.901

	pointsOfFunctionPlot[271].X = 2.71
	pointsOfFunctionPlot[271].Y = -1.016

	pointsOfFunctionPlot[272].X = 2.72
	pointsOfFunctionPlot[272].Y = -1.133

	pointsOfFunctionPlot[273].X = 2.73
	pointsOfFunctionPlot[273].Y = -1.251

	pointsOfFunctionPlot[274].X = 2.74
	pointsOfFunctionPlot[274].Y = -1.371

	pointsOfFunctionPlot[275].X = 2.75
	pointsOfFunctionPlot[275].Y = -1.493

	pointsOfFunctionPlot[276].X = 2.76
	pointsOfFunctionPlot[276].Y = -1.616

	pointsOfFunctionPlot[277].X = 2.77
	pointsOfFunctionPlot[277].Y = -1.74

	pointsOfFunctionPlot[278].X = 2.78
	pointsOfFunctionPlot[278].Y = -1.866

	pointsOfFunctionPlot[279].X = 2.79
	pointsOfFunctionPlot[279].Y = -1.994

	pointsOfFunctionPlot[280].X = 2.8
	pointsOfFunctionPlot[280].Y = -2.123

	pointsOfFunctionPlot[281].X = 2.81
	pointsOfFunctionPlot[281].Y = -2.254

	pointsOfFunctionPlot[282].X = 2.82
	pointsOfFunctionPlot[282].Y = -2.386

	pointsOfFunctionPlot[283].X = 2.83
	pointsOfFunctionPlot[283].Y = -2.52

	pointsOfFunctionPlot[284].X = 2.84
	pointsOfFunctionPlot[284].Y = -2.656

	pointsOfFunctionPlot[285].X = 2.85
	pointsOfFunctionPlot[285].Y = -2.793

	pointsOfFunctionPlot[286].X = 2.86
	pointsOfFunctionPlot[286].Y = -2.932

	pointsOfFunctionPlot[287].X = 2.87
	pointsOfFunctionPlot[287].Y = -3.073

	pointsOfFunctionPlot[288].X = 2.88
	pointsOfFunctionPlot[288].Y = -3.215

	pointsOfFunctionPlot[289].X = 2.89
	pointsOfFunctionPlot[289].Y = -3.359

	pointsOfFunctionPlot[290].X = 2.9
	pointsOfFunctionPlot[290].Y = -3.504

	pointsOfFunctionPlot[291].X = 2.91
	pointsOfFunctionPlot[291].Y = -3.651

	pointsOfFunctionPlot[292].X = 2.92
	pointsOfFunctionPlot[292].Y = -3.8

	pointsOfFunctionPlot[293].X = 2.93
	pointsOfFunctionPlot[293].Y = -3.95

	pointsOfFunctionPlot[294].X = 2.94
	pointsOfFunctionPlot[294].Y = -4.102

	pointsOfFunctionPlot[295].X = 2.95
	pointsOfFunctionPlot[295].Y = -4.256

	pointsOfFunctionPlot[296].X = 2.96
	pointsOfFunctionPlot[296].Y = -4.411

	pointsOfFunctionPlot[297].X = 2.97
	pointsOfFunctionPlot[297].Y = -4.568

	pointsOfFunctionPlot[298].X = 2.98
	pointsOfFunctionPlot[298].Y = -4.727

	pointsOfFunctionPlot[299].X = 2.99
	pointsOfFunctionPlot[299].Y = -4.887

	pointsOfFunctionPlot[300].X = 3.0
	pointsOfFunctionPlot[300].Y = -5.05

	pointsOfFunctionPlot[301].X = 3.01
	pointsOfFunctionPlot[301].Y = -5.213

	pointsOfFunctionPlot[302].X = 3.02
	pointsOfFunctionPlot[302].Y = -5.379

	pointsOfFunctionPlot[303].X = 3.03
	pointsOfFunctionPlot[303].Y = -5.546

	pointsOfFunctionPlot[304].X = 3.04
	pointsOfFunctionPlot[304].Y = -5.715

	pointsOfFunctionPlot[305].X = 3.05
	pointsOfFunctionPlot[305].Y = -5.886

	pointsOfFunctionPlot[306].X = 3.06
	pointsOfFunctionPlot[306].Y = -6.058

	pointsOfFunctionPlot[307].X = 3.07
	pointsOfFunctionPlot[307].Y = -6.232

	pointsOfFunctionPlot[308].X = 3.08
	pointsOfFunctionPlot[308].Y = -6.408

	pointsOfFunctionPlot[309].X = 3.09
	pointsOfFunctionPlot[309].Y = -6.585

	pointsOfFunctionPlot[310].X = 3.1
	pointsOfFunctionPlot[310].Y = -6.765

	pointsOfFunctionPlot[311].X = 3.11
	pointsOfFunctionPlot[311].Y = -6.946

	pointsOfFunctionPlot[312].X = 3.12
	pointsOfFunctionPlot[312].Y = -7.128

	pointsOfFunctionPlot[313].X = 3.13
	pointsOfFunctionPlot[313].Y = -7.313

	pointsOfFunctionPlot[314].X = 3.14
	pointsOfFunctionPlot[314].Y = -7.499

	pointsOfFunctionPlot[315].X = 3.15
	pointsOfFunctionPlot[315].Y = -7.687

	pointsOfFunctionPlot[316].X = 3.16
	pointsOfFunctionPlot[316].Y = -7.887

	pointsOfFunctionPlot[317].X = 3.17
	pointsOfFunctionPlot[317].Y = -8.068

	pointsOfFunctionPlot[318].X = 3.18
	pointsOfFunctionPlot[318].Y = -8.262

	pointsOfFunctionPlot[319].X = 3.19
	pointsOfFunctionPlot[319].Y = -8.457

	pointsOfFunctionPlot[320].X = 3.2
	pointsOfFunctionPlot[320].Y = -8.653










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function cos(x)"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("cos(x)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"cos-plot-01.png"); err != nil {

		panic(err)
	}
}
