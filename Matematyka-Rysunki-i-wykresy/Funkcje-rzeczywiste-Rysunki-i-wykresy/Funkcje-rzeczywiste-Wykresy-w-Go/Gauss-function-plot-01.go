package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function exp(-x^2).

	pointsOfFunctionPlot := make(plotter.XYs, 529)

	pointsOfFunctionPlot[0].X = -10.0
	pointsOfFunctionPlot[0].Y = 0.0

	pointsOfFunctionPlot[1].X = -2.64
	pointsOfFunctionPlot[1].Y = 0.0

	pointsOfFunctionPlot[2].X = -2.63
	pointsOfFunctionPlot[2].Y = 0.001

	pointsOfFunctionPlot[3].X = -2.62
	pointsOfFunctionPlot[3].Y = 0.001

	pointsOfFunctionPlot[4].X = -2.61
	pointsOfFunctionPlot[4].Y = 0.001

	pointsOfFunctionPlot[5].X = -2.60
	pointsOfFunctionPlot[5].Y = 0.001

	pointsOfFunctionPlot[6].X = -2.59
	pointsOfFunctionPlot[6].Y = 0.001

	pointsOfFunctionPlot[7].X = -2.58
	pointsOfFunctionPlot[7].Y = 0.001

	pointsOfFunctionPlot[8].X = -2.57
	pointsOfFunctionPlot[8].Y = 0.001

	pointsOfFunctionPlot[9].X = -2.56
	pointsOfFunctionPlot[9].Y = 0.001

	pointsOfFunctionPlot[10].X = -2.55
	pointsOfFunctionPlot[10].Y = 0.001

	pointsOfFunctionPlot[11].X = -2.54
	pointsOfFunctionPlot[11].Y = 0.001

	pointsOfFunctionPlot[12].X = -2.53
	pointsOfFunctionPlot[12].Y = 0.001

	pointsOfFunctionPlot[13].X = -2.52
	pointsOfFunctionPlot[13].Y = 0.001

	pointsOfFunctionPlot[14].X = -2.51
	pointsOfFunctionPlot[14].Y = 0.001

	pointsOfFunctionPlot[15].X = -2.50
	pointsOfFunctionPlot[15].Y = 0.001

	pointsOfFunctionPlot[16].X = -2.49
	pointsOfFunctionPlot[16].Y = 0.002

	pointsOfFunctionPlot[17].X = -2.48
	pointsOfFunctionPlot[17].Y = 0.002

	pointsOfFunctionPlot[18].X = -2.47
	pointsOfFunctionPlot[18].Y = 0.002

	pointsOfFunctionPlot[19].X = -2.46
	pointsOfFunctionPlot[19].Y = 0.002

	pointsOfFunctionPlot[20].X = -2.45
	pointsOfFunctionPlot[20].Y = 0.002

	pointsOfFunctionPlot[21].X = -2.44
	pointsOfFunctionPlot[21].Y = 0.002

	pointsOfFunctionPlot[22].X = -2.43
	pointsOfFunctionPlot[22].Y = 0.002

	pointsOfFunctionPlot[23].X = -2.42
	pointsOfFunctionPlot[23].Y = 0.002

	pointsOfFunctionPlot[24].X = -2.41
	pointsOfFunctionPlot[24].Y = 0.003

	pointsOfFunctionPlot[25].X = -2.40
	pointsOfFunctionPlot[25].Y = 0.003

	pointsOfFunctionPlot[26].X = -2.39
	pointsOfFunctionPlot[26].Y = 0.003

	pointsOfFunctionPlot[27].X = -2.38
	pointsOfFunctionPlot[27].Y = 0.003

	pointsOfFunctionPlot[28].X = -2.37
	pointsOfFunctionPlot[28].Y = 0.003

	pointsOfFunctionPlot[29].X = -2.36
	pointsOfFunctionPlot[29].Y = 0.003

	pointsOfFunctionPlot[30].X = -2.35
	pointsOfFunctionPlot[30].Y = 0.004

	pointsOfFunctionPlot[31].X = -2.34
	pointsOfFunctionPlot[31].Y = 0.004

	pointsOfFunctionPlot[32].X = -2.33
	pointsOfFunctionPlot[32].Y = 0.004

	pointsOfFunctionPlot[33].X = -2.32
	pointsOfFunctionPlot[33].Y = 0.004

	pointsOfFunctionPlot[34].X = -2.31
	pointsOfFunctionPlot[34].Y = 0.004

	pointsOfFunctionPlot[35].X = -2.30
	pointsOfFunctionPlot[35].Y = 0.005

	pointsOfFunctionPlot[36].X = -2.29
	pointsOfFunctionPlot[36].Y = 0.005

	pointsOfFunctionPlot[37].X = -2.28
	pointsOfFunctionPlot[37].Y = 0.005

	pointsOfFunctionPlot[38].X = -2.27
	pointsOfFunctionPlot[38].Y = 0.005

	pointsOfFunctionPlot[39].X = -2.26
	pointsOfFunctionPlot[39].Y = 0.006

	pointsOfFunctionPlot[40].X = -2.25
	pointsOfFunctionPlot[40].Y = 0.006

	pointsOfFunctionPlot[41].X = -2.24
	pointsOfFunctionPlot[41].Y = 0.006

	pointsOfFunctionPlot[42].X = -2.23
	pointsOfFunctionPlot[42].Y = 0.006

	pointsOfFunctionPlot[43].X = -2.22
	pointsOfFunctionPlot[43].Y = 0.007

	pointsOfFunctionPlot[44].X = -2.21
	pointsOfFunctionPlot[44].Y = 0.007

	pointsOfFunctionPlot[45].X = -2.20
	pointsOfFunctionPlot[45].Y = 0.007

	pointsOfFunctionPlot[46].X = -2.19
	pointsOfFunctionPlot[46].Y = 0.008

	pointsOfFunctionPlot[47].X = -2.18
	pointsOfFunctionPlot[47].Y = 0.008

	pointsOfFunctionPlot[48].X = -2.17
	pointsOfFunctionPlot[48].Y = 0.009

	pointsOfFunctionPlot[49].X = -2.16
	pointsOfFunctionPlot[49].Y = 0.009

	pointsOfFunctionPlot[50].X = -2.15
	pointsOfFunctionPlot[50].Y = 0.009

	pointsOfFunctionPlot[51].X = -2.14
	pointsOfFunctionPlot[51].Y = 0.01

	pointsOfFunctionPlot[52].X = -2.13
	pointsOfFunctionPlot[52].Y = 0.01

	pointsOfFunctionPlot[53].X = -2.12
	pointsOfFunctionPlot[53].Y = 0.011

	pointsOfFunctionPlot[54].X = -2.11
	pointsOfFunctionPlot[54].Y = 0.011

	pointsOfFunctionPlot[55].X = -2.10
	pointsOfFunctionPlot[55].Y = 0.012

	pointsOfFunctionPlot[56].X = -2.09
	pointsOfFunctionPlot[56].Y = 0.012

	pointsOfFunctionPlot[57].X = -2.08
	pointsOfFunctionPlot[57].Y = 0.013

	pointsOfFunctionPlot[58].X = -2.07
	pointsOfFunctionPlot[58].Y = 0.013

	pointsOfFunctionPlot[59].X = -2.06
	pointsOfFunctionPlot[59].Y = 0.014

	pointsOfFunctionPlot[60].X = -2.05
	pointsOfFunctionPlot[60].Y = 0.015

	pointsOfFunctionPlot[61].X = -2.04
	pointsOfFunctionPlot[61].Y = 0.015

	pointsOfFunctionPlot[62].X = -2.03
	pointsOfFunctionPlot[62].Y = 0.016

	pointsOfFunctionPlot[63].X = -2.02
	pointsOfFunctionPlot[63].Y = 0.016

	pointsOfFunctionPlot[64].X = -2.01
	pointsOfFunctionPlot[64].Y = 0.017

	pointsOfFunctionPlot[65].X = -2.0
	pointsOfFunctionPlot[65].Y = 0.018

	pointsOfFunctionPlot[66].X = -1.99
	pointsOfFunctionPlot[66].Y = 0.019

	pointsOfFunctionPlot[67].X = -1.98
	pointsOfFunctionPlot[67].Y = 0.019

	pointsOfFunctionPlot[68].X = -1.97
	pointsOfFunctionPlot[68].Y = 0.02

	pointsOfFunctionPlot[69].X = -1.96
	pointsOfFunctionPlot[69].Y = 0.021

	pointsOfFunctionPlot[70].X = -1.95
	pointsOfFunctionPlot[70].Y = 0.022

	pointsOfFunctionPlot[71].X = -1.94
	pointsOfFunctionPlot[71].Y = 0.023

	pointsOfFunctionPlot[72].X = -1.93
	pointsOfFunctionPlot[72].Y = 0.024

	pointsOfFunctionPlot[73].X = -1.92
	pointsOfFunctionPlot[73].Y = 0.025

	pointsOfFunctionPlot[74].X = -1.91
	pointsOfFunctionPlot[74].Y = 0.026

	pointsOfFunctionPlot[75].X = -1.90
	pointsOfFunctionPlot[75].Y = 0.027

	pointsOfFunctionPlot[76].X = -1.89
	pointsOfFunctionPlot[76].Y = 0.028

	pointsOfFunctionPlot[77].X = -1.88
	pointsOfFunctionPlot[77].Y = 0.029

	pointsOfFunctionPlot[78].X = -1.87
	pointsOfFunctionPlot[78].Y = 0.03

	pointsOfFunctionPlot[79].X = -1.86
	pointsOfFunctionPlot[79].Y = 0.031

	pointsOfFunctionPlot[80].X = -1.85
	pointsOfFunctionPlot[80].Y = 0.032

	pointsOfFunctionPlot[81].X = -1.84
	pointsOfFunctionPlot[81].Y = 0.033

	pointsOfFunctionPlot[82].X = -1.83
	pointsOfFunctionPlot[82].Y = 0.035

	pointsOfFunctionPlot[83].X = -1.82
	pointsOfFunctionPlot[83].Y = 0.036

	pointsOfFunctionPlot[84].X = -1.81
	pointsOfFunctionPlot[84].Y = 0.037

	pointsOfFunctionPlot[85].X = -1.80
	pointsOfFunctionPlot[85].Y = 0.039

	pointsOfFunctionPlot[86].X = -1.79
	pointsOfFunctionPlot[86].Y = 0.04

	pointsOfFunctionPlot[87].X = -1.78
	pointsOfFunctionPlot[87].Y = 0.042

	pointsOfFunctionPlot[88].X = -1.77
	pointsOfFunctionPlot[88].Y = 0.043

	pointsOfFunctionPlot[89].X = -1.76
	pointsOfFunctionPlot[89].Y = 0.045

	pointsOfFunctionPlot[90].X = -1.75
	pointsOfFunctionPlot[90].Y = 0.046

	pointsOfFunctionPlot[91].X = -1.74
	pointsOfFunctionPlot[91].Y = 0.048

	pointsOfFunctionPlot[92].X = -1.73
	pointsOfFunctionPlot[92].Y = 0.05

	pointsOfFunctionPlot[93].X = -1.72
	pointsOfFunctionPlot[93].Y = 0.051

	pointsOfFunctionPlot[94].X = -1.71
	pointsOfFunctionPlot[94].Y = 0.053

	pointsOfFunctionPlot[95].X = -1.70
	pointsOfFunctionPlot[95].Y = 0.055

	pointsOfFunctionPlot[96].X = -1.69
	pointsOfFunctionPlot[96].Y = 0.057

	pointsOfFunctionPlot[97].X = -1.68
	pointsOfFunctionPlot[97].Y = 0.059

	pointsOfFunctionPlot[98].X = -1.67
	pointsOfFunctionPlot[98].Y = 0.061

	pointsOfFunctionPlot[99].X = -1.66
	pointsOfFunctionPlot[99].Y = 0.063

	pointsOfFunctionPlot[100].X = -1.65
	pointsOfFunctionPlot[100].Y = 0.065

	pointsOfFunctionPlot[101].X = -1.64
	pointsOfFunctionPlot[101].Y = 0.067

	pointsOfFunctionPlot[102].X = -1.63
	pointsOfFunctionPlot[102].Y = 0.07

	pointsOfFunctionPlot[103].X = -1.62
	pointsOfFunctionPlot[103].Y = 0.072

	pointsOfFunctionPlot[104].X = -1.61
	pointsOfFunctionPlot[104].Y = 0.074

	pointsOfFunctionPlot[105].X = -1.60
	pointsOfFunctionPlot[105].Y = 0.077

	pointsOfFunctionPlot[106].X = -1.59
	pointsOfFunctionPlot[106].Y = 0.079

	pointsOfFunctionPlot[107].X = -1.58
	pointsOfFunctionPlot[107].Y = 0.082

	pointsOfFunctionPlot[108].X = -1.57
	pointsOfFunctionPlot[108].Y = 0.085

	pointsOfFunctionPlot[109].X = -1.56
	pointsOfFunctionPlot[109].Y = 0.087

	pointsOfFunctionPlot[110].X = -1.55
	pointsOfFunctionPlot[110].Y = 0.09

	pointsOfFunctionPlot[111].X = -1.54
	pointsOfFunctionPlot[111].Y = 0.093

	pointsOfFunctionPlot[112].X = -1.53
	pointsOfFunctionPlot[112].Y = 0.096

	pointsOfFunctionPlot[113].X = -1.52
	pointsOfFunctionPlot[113].Y = 0.099

	pointsOfFunctionPlot[114].X = -1.51
	pointsOfFunctionPlot[114].Y = 0.102

	pointsOfFunctionPlot[115].X = -1.50
	pointsOfFunctionPlot[115].Y = 0.105

	pointsOfFunctionPlot[116].X = -1.49
	pointsOfFunctionPlot[116].Y = 0.108

	pointsOfFunctionPlot[117].X = -1.48
	pointsOfFunctionPlot[117].Y = 0.111

	pointsOfFunctionPlot[118].X = -1.47
	pointsOfFunctionPlot[118].Y = 0.115

	pointsOfFunctionPlot[119].X = -1.46
	pointsOfFunctionPlot[119].Y = 0.118

	pointsOfFunctionPlot[120].X = -1.45
	pointsOfFunctionPlot[120].Y = 0.122

	pointsOfFunctionPlot[121].X = -1.44
	pointsOfFunctionPlot[121].Y = 0.125

	pointsOfFunctionPlot[122].X = -1.43
	pointsOfFunctionPlot[122].Y = 0.129

	pointsOfFunctionPlot[123].X = -1.42
	pointsOfFunctionPlot[123].Y = 0.133

	pointsOfFunctionPlot[124].X = -1.41
	pointsOfFunctionPlot[124].Y = 0.137

	pointsOfFunctionPlot[125].X = -1.40
	pointsOfFunctionPlot[125].Y = 0.14

	pointsOfFunctionPlot[126].X = -1.39
	pointsOfFunctionPlot[126].Y = 0.144

	pointsOfFunctionPlot[127].X = -1.38
	pointsOfFunctionPlot[127].Y = 0.148

	pointsOfFunctionPlot[128].X = -1.37
	pointsOfFunctionPlot[128].Y = 0.153

	pointsOfFunctionPlot[129].X = -1.36
	pointsOfFunctionPlot[129].Y = 0.157

	pointsOfFunctionPlot[130].X = -1.35
	pointsOfFunctionPlot[130].Y = 0.161

	pointsOfFunctionPlot[131].X = -1.34
	pointsOfFunctionPlot[131].Y = 0.166

	pointsOfFunctionPlot[132].X = -1.33
	pointsOfFunctionPlot[132].Y = 0.17

	pointsOfFunctionPlot[133].X = -1.32
	pointsOfFunctionPlot[133].Y = 0.175

	pointsOfFunctionPlot[134].X = -1.31
	pointsOfFunctionPlot[134].Y = 0.179

	pointsOfFunctionPlot[135].X = -1.30
	pointsOfFunctionPlot[135].Y = 0.184

	pointsOfFunctionPlot[136].X = -1.29
	pointsOfFunctionPlot[136].Y = 0.189

	pointsOfFunctionPlot[137].X = -1.28
	pointsOfFunctionPlot[137].Y = 0.194

	pointsOfFunctionPlot[138].X = -1.27
	pointsOfFunctionPlot[138].Y = 0.199

	pointsOfFunctionPlot[139].X = -1.26
	pointsOfFunctionPlot[139].Y = 0.204

	pointsOfFunctionPlot[140].X = -1.25
	pointsOfFunctionPlot[140].Y = 0.209

	pointsOfFunctionPlot[141].X = -1.24
	pointsOfFunctionPlot[141].Y = 0.214

	pointsOfFunctionPlot[142].X = -1.23
	pointsOfFunctionPlot[142].Y = 0.22

	pointsOfFunctionPlot[143].X = -1.22
	pointsOfFunctionPlot[143].Y = 0.225

	pointsOfFunctionPlot[144].X = -1.21
	pointsOfFunctionPlot[144].Y = 0.231

	pointsOfFunctionPlot[145].X = -1.20
	pointsOfFunctionPlot[145].Y = 0.236

	pointsOfFunctionPlot[146].X = -1.19
	pointsOfFunctionPlot[146].Y = 0.242

	pointsOfFunctionPlot[147].X = -1.18
	pointsOfFunctionPlot[147].Y = 0.248

	pointsOfFunctionPlot[148].X = -1.17
	pointsOfFunctionPlot[148].Y = 0.254

	pointsOfFunctionPlot[149].X = -1.16
	pointsOfFunctionPlot[149].Y = 0.26

	pointsOfFunctionPlot[150].X = -1.15
	pointsOfFunctionPlot[150].Y = 0.266

	pointsOfFunctionPlot[151].X = -1.14
	pointsOfFunctionPlot[151].Y = 0.272

	pointsOfFunctionPlot[152].X = -1.13
	pointsOfFunctionPlot[152].Y = 0.278

	pointsOfFunctionPlot[153].X = -1.12
	pointsOfFunctionPlot[153].Y = 0.285

	pointsOfFunctionPlot[154].X = -1.11
	pointsOfFunctionPlot[154].Y = 0.291

	pointsOfFunctionPlot[155].X = -1.10
	pointsOfFunctionPlot[155].Y = 0.298

	pointsOfFunctionPlot[156].X = -1.09
	pointsOfFunctionPlot[156].Y = 0.304

	pointsOfFunctionPlot[157].X = -1.08
	pointsOfFunctionPlot[157].Y = 0.311

	pointsOfFunctionPlot[158].X = -1.07
	pointsOfFunctionPlot[158].Y = 0.318

	pointsOfFunctionPlot[159].X = -1.06
	pointsOfFunctionPlot[159].Y = 0.325

	pointsOfFunctionPlot[160].X = -1.05
	pointsOfFunctionPlot[160].Y = 0.332

	pointsOfFunctionPlot[161].X = -1.04
	pointsOfFunctionPlot[161].Y = 0.339

	pointsOfFunctionPlot[162].X = -1.03
	pointsOfFunctionPlot[162].Y = 0.346

	pointsOfFunctionPlot[163].X = -1.02
	pointsOfFunctionPlot[163].Y = 0.353

	pointsOfFunctionPlot[164].X = -1.01
	pointsOfFunctionPlot[164].Y = 0.36

	pointsOfFunctionPlot[165].X = -1.0
	pointsOfFunctionPlot[165].Y = 0.367

	pointsOfFunctionPlot[166].X = -0.99
	pointsOfFunctionPlot[166].Y = 0.375

	pointsOfFunctionPlot[167].X = -0.98
	pointsOfFunctionPlot[167].Y = 0.382

	pointsOfFunctionPlot[168].X = -0.97
	pointsOfFunctionPlot[168].Y = 0.39

	pointsOfFunctionPlot[169].X = -0.96
	pointsOfFunctionPlot[169].Y = 0.397

	pointsOfFunctionPlot[170].X = -0.95
	pointsOfFunctionPlot[170].Y = 0.405

	pointsOfFunctionPlot[171].X = -0.94
	pointsOfFunctionPlot[171].Y = 0.413

	pointsOfFunctionPlot[172].X = -0.93
	pointsOfFunctionPlot[172].Y = 0.421

	pointsOfFunctionPlot[173].X = -0.92
	pointsOfFunctionPlot[173].Y = 0.429

	pointsOfFunctionPlot[174].X = -0.91
	pointsOfFunctionPlot[174].Y = 0.436

	pointsOfFunctionPlot[175].X = -0.90
	pointsOfFunctionPlot[175].Y = 0.444

	pointsOfFunctionPlot[176].X = -0.89
	pointsOfFunctionPlot[176].Y = 0.452

	pointsOfFunctionPlot[177].X = -0.88
	pointsOfFunctionPlot[177].Y = 0.461

	pointsOfFunctionPlot[178].X = -0.87
	pointsOfFunctionPlot[178].Y = 0.469

	pointsOfFunctionPlot[179].X = -0.86
	pointsOfFunctionPlot[179].Y = 0.477

	pointsOfFunctionPlot[180].X = -0.85
	pointsOfFunctionPlot[180].Y = 0.485

	pointsOfFunctionPlot[181].X = -0.84
	pointsOfFunctionPlot[181].Y = 0.493

	pointsOfFunctionPlot[182].X = -0.83
	pointsOfFunctionPlot[182].Y = 0.502

	pointsOfFunctionPlot[183].X = -0.82
	pointsOfFunctionPlot[183].Y = 0.51

	pointsOfFunctionPlot[184].X = -0.81
	pointsOfFunctionPlot[184].Y = 0.518

	pointsOfFunctionPlot[185].X = -0.80
	pointsOfFunctionPlot[185].Y = 0.527

	pointsOfFunctionPlot[186].X = -0.79
	pointsOfFunctionPlot[186].Y = 0.535

	pointsOfFunctionPlot[187].X = -0.78
	pointsOfFunctionPlot[187].Y = 0.544

	pointsOfFunctionPlot[188].X = -0.77
	pointsOfFunctionPlot[188].Y = 0.552

	pointsOfFunctionPlot[189].X = -0.76
	pointsOfFunctionPlot[189].Y = 0.561

	pointsOfFunctionPlot[190].X = -0.75
	pointsOfFunctionPlot[190].Y = 0.569

	pointsOfFunctionPlot[191].X = -0.74
	pointsOfFunctionPlot[191].Y = 0.578

	pointsOfFunctionPlot[192].X = -0.73
	pointsOfFunctionPlot[192].Y = 0.586

	pointsOfFunctionPlot[193].X = -0.72
	pointsOfFunctionPlot[193].Y = 0.595

	pointsOfFunctionPlot[194].X = -0.71
	pointsOfFunctionPlot[194].Y = 0.604

	pointsOfFunctionPlot[195].X = -0.70
	pointsOfFunctionPlot[195].Y = 0.612

	pointsOfFunctionPlot[196].X = -0.69
	pointsOfFunctionPlot[196].Y = 0.621

	pointsOfFunctionPlot[197].X = -0.68
	pointsOfFunctionPlot[197].Y = 0.629

	pointsOfFunctionPlot[198].X = -0.67
	pointsOfFunctionPlot[198].Y = 0.638

	pointsOfFunctionPlot[199].X = -0.66
	pointsOfFunctionPlot[199].Y = 0.646

	pointsOfFunctionPlot[200].X = -0.65
	pointsOfFunctionPlot[200].Y = 0.655

	pointsOfFunctionPlot[201].X = -0.64
	pointsOfFunctionPlot[201].Y = 0.663

	pointsOfFunctionPlot[202].X = -0.63
	pointsOfFunctionPlot[202].Y = 0.672

	pointsOfFunctionPlot[203].X = -0.62
	pointsOfFunctionPlot[203].Y = 0.68

	pointsOfFunctionPlot[204].X = -0.61
	pointsOfFunctionPlot[204].Y = 0.689

	pointsOfFunctionPlot[205].X = -0.60
	pointsOfFunctionPlot[205].Y = 0.697

	pointsOfFunctionPlot[206].X = -0.59
	pointsOfFunctionPlot[206].Y = 0.706

	pointsOfFunctionPlot[207].X = -0.58
	pointsOfFunctionPlot[207].Y = 0.714

	pointsOfFunctionPlot[208].X = -0.57
	pointsOfFunctionPlot[208].Y = 0.722

	pointsOfFunctionPlot[209].X = -0.56
	pointsOfFunctionPlot[209].Y = 0.73

	pointsOfFunctionPlot[210].X = -0.55
	pointsOfFunctionPlot[210].Y = 0.739

	pointsOfFunctionPlot[211].X = -0.54
	pointsOfFunctionPlot[211].Y = 0.747

	pointsOfFunctionPlot[212].X = -0.53
	pointsOfFunctionPlot[212].Y = 0.755

	pointsOfFunctionPlot[213].X = -0.52
	pointsOfFunctionPlot[213].Y = 0.763

	pointsOfFunctionPlot[214].X = -0.51
	pointsOfFunctionPlot[214].Y = 0.771

	pointsOfFunctionPlot[215].X = -0.50
	pointsOfFunctionPlot[215].Y = 0.778

	pointsOfFunctionPlot[216].X = -0.49
	pointsOfFunctionPlot[216].Y = 0.786

	pointsOfFunctionPlot[217].X = -0.48
	pointsOfFunctionPlot[217].Y = 0.794

	pointsOfFunctionPlot[218].X = -0.47
	pointsOfFunctionPlot[218].Y = 0.801

	pointsOfFunctionPlot[219].X = -0.46
	pointsOfFunctionPlot[219].Y = 0.809

	pointsOfFunctionPlot[220].X = -0.45
	pointsOfFunctionPlot[220].Y = 0.816

	pointsOfFunctionPlot[221].X = -0.44
	pointsOfFunctionPlot[221].Y = 0.824

	pointsOfFunctionPlot[222].X = -0.43
	pointsOfFunctionPlot[222].Y = 0.831

	pointsOfFunctionPlot[223].X = -0.42
	pointsOfFunctionPlot[223].Y = 0.838

	pointsOfFunctionPlot[224].X = -0.41
	pointsOfFunctionPlot[224].Y = 0.845

	pointsOfFunctionPlot[225].X = -0.40
	pointsOfFunctionPlot[225].Y = 0.852

	pointsOfFunctionPlot[226].X = -0.39
	pointsOfFunctionPlot[226].Y = 0.858

	pointsOfFunctionPlot[227].X = -0.38
	pointsOfFunctionPlot[227].Y = 0.865

	pointsOfFunctionPlot[228].X = -0.37
	pointsOfFunctionPlot[228].Y = 0.872

	pointsOfFunctionPlot[229].X = -0.36
	pointsOfFunctionPlot[229].Y = 0.878

	pointsOfFunctionPlot[230].X = -0.35
	pointsOfFunctionPlot[230].Y = 0.884

	pointsOfFunctionPlot[231].X = -0.34
	pointsOfFunctionPlot[231].Y = 0.89

	pointsOfFunctionPlot[232].X = -0.33
	pointsOfFunctionPlot[232].Y = 0.896

	pointsOfFunctionPlot[233].X = -0.32
	pointsOfFunctionPlot[233].Y = 0.902

	pointsOfFunctionPlot[234].X = -0.31
	pointsOfFunctionPlot[234].Y = 0.908

	pointsOfFunctionPlot[235].X = -0.30
	pointsOfFunctionPlot[235].Y = 0.913

	pointsOfFunctionPlot[236].X = -0.29
	pointsOfFunctionPlot[236].Y = 0.919

	pointsOfFunctionPlot[237].X = -0.28
	pointsOfFunctionPlot[237].Y = 0.924

	pointsOfFunctionPlot[238].X = -0.27
	pointsOfFunctionPlot[238].Y = 0.929

	pointsOfFunctionPlot[239].X = -0.26
	pointsOfFunctionPlot[239].Y = 0.934

	pointsOfFunctionPlot[240].X = -0.25
	pointsOfFunctionPlot[240].Y = 0.939

	pointsOfFunctionPlot[241].X = -0.24
	pointsOfFunctionPlot[241].Y = 0.944

	pointsOfFunctionPlot[242].X = -0.23
	pointsOfFunctionPlot[242].Y = 0.948

	pointsOfFunctionPlot[243].X = -0.22
	pointsOfFunctionPlot[243].Y = 0.952

	pointsOfFunctionPlot[244].X = -0.21
	pointsOfFunctionPlot[244].Y = 0.956

	pointsOfFunctionPlot[245].X = -0.20
	pointsOfFunctionPlot[245].Y = 0.96

	pointsOfFunctionPlot[246].X = -0.19
	pointsOfFunctionPlot[246].Y = 0.964

	pointsOfFunctionPlot[247].X = -0.18
	pointsOfFunctionPlot[247].Y = 0.968

	pointsOfFunctionPlot[248].X = -0.17
	pointsOfFunctionPlot[248].Y = 0.971

	pointsOfFunctionPlot[249].X = -0.16
	pointsOfFunctionPlot[249].Y = 0.974

	pointsOfFunctionPlot[250].X = -0.15
	pointsOfFunctionPlot[250].Y = 0.977

	pointsOfFunctionPlot[251].X = -0.14
	pointsOfFunctionPlot[251].Y = 0.98

	pointsOfFunctionPlot[252].X = -0.13
	pointsOfFunctionPlot[252].Y = 0.983

	pointsOfFunctionPlot[253].X = -0.12
	pointsOfFunctionPlot[253].Y = 0.985

	pointsOfFunctionPlot[254].X = -0.11
	pointsOfFunctionPlot[254].Y = 0.988

	pointsOfFunctionPlot[255].X = -0.10
	pointsOfFunctionPlot[255].Y = 0.99

	pointsOfFunctionPlot[256].X = -0.09
	pointsOfFunctionPlot[256].Y = 0.991

	pointsOfFunctionPlot[257].X = -0.08
	pointsOfFunctionPlot[257].Y = 0.993

	pointsOfFunctionPlot[258].X = -0.07
	pointsOfFunctionPlot[258].Y = 0.995

	pointsOfFunctionPlot[259].X = -0.06
	pointsOfFunctionPlot[259].Y = 0.996

	pointsOfFunctionPlot[260].X = -0.05
	pointsOfFunctionPlot[260].Y = 0.997

	pointsOfFunctionPlot[261].X = -0.04
	pointsOfFunctionPlot[261].Y = 0.998

	pointsOfFunctionPlot[262].X = -0.03
	pointsOfFunctionPlot[262].Y = 0.999

	pointsOfFunctionPlot[263].X = -0.02
	pointsOfFunctionPlot[263].Y = 0.999

	pointsOfFunctionPlot[264].X = -0.01
	pointsOfFunctionPlot[264].Y = 0.999

	pointsOfFunctionPlot[265].X = 0.0
	pointsOfFunctionPlot[265].Y = 1.0

	pointsOfFunctionPlot[266].X = 0.01
	pointsOfFunctionPlot[266].Y = 0.999

	pointsOfFunctionPlot[267].X = 0.02
	pointsOfFunctionPlot[267].Y = 0.999

	pointsOfFunctionPlot[268].X = 0.03
	pointsOfFunctionPlot[268].Y = 0.999

	pointsOfFunctionPlot[269].X = 0.04
	pointsOfFunctionPlot[269].Y = 0.998

	pointsOfFunctionPlot[270].X = 0.05
	pointsOfFunctionPlot[270].Y = 0.997

	pointsOfFunctionPlot[271].X = 0.06
	pointsOfFunctionPlot[271].Y = 0.996

	pointsOfFunctionPlot[272].X = 0.07
	pointsOfFunctionPlot[272].Y = 0.995

	pointsOfFunctionPlot[273].X = 0.08
	pointsOfFunctionPlot[273].Y = 0.993

	pointsOfFunctionPlot[274].X = 0.09
	pointsOfFunctionPlot[274].Y = 0.991

	pointsOfFunctionPlot[275].X = 0.10
	pointsOfFunctionPlot[275].Y = 0.99

	pointsOfFunctionPlot[276].X = 0.11
	pointsOfFunctionPlot[276].Y = 0.988

	pointsOfFunctionPlot[277].X = 0.12
	pointsOfFunctionPlot[277].Y = 0.985

	pointsOfFunctionPlot[278].X = 0.13
	pointsOfFunctionPlot[278].Y = 0.983

	pointsOfFunctionPlot[279].X = 0.14
	pointsOfFunctionPlot[279].Y = 0.98

	pointsOfFunctionPlot[280].X = 0.15
	pointsOfFunctionPlot[280].Y = 0.977

	pointsOfFunctionPlot[281].X = 0.16
	pointsOfFunctionPlot[281].Y = 0.974

	pointsOfFunctionPlot[282].X = 0.17
	pointsOfFunctionPlot[282].Y = 0.971

	pointsOfFunctionPlot[283].X = 0.18
	pointsOfFunctionPlot[283].Y = 0.968

	pointsOfFunctionPlot[284].X = 0.19
	pointsOfFunctionPlot[284].Y = 0.964

	pointsOfFunctionPlot[285].X = 0.20
	pointsOfFunctionPlot[285].Y = 0.96

	pointsOfFunctionPlot[286].X = 0.21
	pointsOfFunctionPlot[286].Y = 0.956

	pointsOfFunctionPlot[287].X = 0.22
	pointsOfFunctionPlot[287].Y = 0.952

	pointsOfFunctionPlot[288].X = 0.23
	pointsOfFunctionPlot[288].Y = 0.948

	pointsOfFunctionPlot[289].X = 0.24
	pointsOfFunctionPlot[289].Y = 0.944

	pointsOfFunctionPlot[290].X = 0.25
	pointsOfFunctionPlot[290].Y = 0.939

	pointsOfFunctionPlot[291].X = 0.26
	pointsOfFunctionPlot[291].Y = 0.934

	pointsOfFunctionPlot[292].X = 0.27
	pointsOfFunctionPlot[292].Y = 0.929

	pointsOfFunctionPlot[293].X = 0.28
	pointsOfFunctionPlot[293].Y = 0.924

	pointsOfFunctionPlot[294].X = 0.29
	pointsOfFunctionPlot[294].Y = 0.919

	pointsOfFunctionPlot[295].X = 0.30
	pointsOfFunctionPlot[295].Y = 0.913

	pointsOfFunctionPlot[296].X = 0.31
	pointsOfFunctionPlot[296].Y = 0.908

	pointsOfFunctionPlot[297].X = 0.32
	pointsOfFunctionPlot[297].Y = 0.902

	pointsOfFunctionPlot[298].X = 0.33
	pointsOfFunctionPlot[298].Y = 0.896

	pointsOfFunctionPlot[299].X = 0.34
	pointsOfFunctionPlot[299].Y = 0.89

	pointsOfFunctionPlot[300].X = 0.35
	pointsOfFunctionPlot[300].Y = 0.884

	pointsOfFunctionPlot[301].X = 0.36
	pointsOfFunctionPlot[301].Y = 0.878

	pointsOfFunctionPlot[302].X = 0.37
	pointsOfFunctionPlot[302].Y = 0.872

	pointsOfFunctionPlot[303].X = 0.38
	pointsOfFunctionPlot[303].Y = 0.865

	pointsOfFunctionPlot[304].X = 0.39
	pointsOfFunctionPlot[304].Y = 0.858

	pointsOfFunctionPlot[305].X = 0.40
	pointsOfFunctionPlot[305].Y = 0.852

	pointsOfFunctionPlot[306].X = 0.41
	pointsOfFunctionPlot[306].Y = 0.845

	pointsOfFunctionPlot[307].X = 0.42
	pointsOfFunctionPlot[307].Y = 0.838

	pointsOfFunctionPlot[308].X = 0.43
	pointsOfFunctionPlot[308].Y = 0.831

	pointsOfFunctionPlot[309].X = 0.44
	pointsOfFunctionPlot[309].Y = 0.824

	pointsOfFunctionPlot[310].X = 0.45
	pointsOfFunctionPlot[310].Y = 0.816

	pointsOfFunctionPlot[311].X = 0.46
	pointsOfFunctionPlot[311].Y = 0.809

	pointsOfFunctionPlot[312].X = 0.47
	pointsOfFunctionPlot[312].Y = 0.801

	pointsOfFunctionPlot[313].X = 0.48
	pointsOfFunctionPlot[313].Y = 0.794

	pointsOfFunctionPlot[314].X = 0.49
	pointsOfFunctionPlot[314].Y = 0.786

	pointsOfFunctionPlot[315].X = 0.50
	pointsOfFunctionPlot[315].Y = 0.778

	pointsOfFunctionPlot[316].X = 0.51
	pointsOfFunctionPlot[316].Y = 0.771

	pointsOfFunctionPlot[317].X = 0.52
	pointsOfFunctionPlot[317].Y = 0.763

	pointsOfFunctionPlot[318].X = 0.53
	pointsOfFunctionPlot[318].Y = 0.755

	pointsOfFunctionPlot[319].X = 0.54
	pointsOfFunctionPlot[319].Y = 0.747

	pointsOfFunctionPlot[320].X = 0.55
	pointsOfFunctionPlot[320].Y = 0.739

	pointsOfFunctionPlot[321].X = 0.56
	pointsOfFunctionPlot[321].Y = 0.73

	pointsOfFunctionPlot[322].X = 0.57
	pointsOfFunctionPlot[322].Y = 0.722

	pointsOfFunctionPlot[323].X = 0.58
	pointsOfFunctionPlot[323].Y = 0.714

	pointsOfFunctionPlot[324].X = 0.59
	pointsOfFunctionPlot[324].Y = 0.706

	pointsOfFunctionPlot[325].X = 0.60
	pointsOfFunctionPlot[325].Y = 0.697

	pointsOfFunctionPlot[326].X = 0.61
	pointsOfFunctionPlot[326].Y = 0.689

	pointsOfFunctionPlot[327].X = 0.62
	pointsOfFunctionPlot[327].Y = 0.68

	pointsOfFunctionPlot[328].X = 0.63
	pointsOfFunctionPlot[328].Y = 0.672

	pointsOfFunctionPlot[329].X = 0.64
	pointsOfFunctionPlot[329].Y = 0.663

	pointsOfFunctionPlot[330].X = 0.65
	pointsOfFunctionPlot[330].Y = 0.655

	pointsOfFunctionPlot[331].X = 0.66
	pointsOfFunctionPlot[331].Y = 0.646

	pointsOfFunctionPlot[332].X = 0.67
	pointsOfFunctionPlot[332].Y = 0.638

	pointsOfFunctionPlot[333].X = 0.68
	pointsOfFunctionPlot[333].Y = 0.629

	pointsOfFunctionPlot[334].X = 0.69
	pointsOfFunctionPlot[334].Y = 0.621

	pointsOfFunctionPlot[335].X = 0.70
	pointsOfFunctionPlot[335].Y = 0.612

	pointsOfFunctionPlot[336].X = 0.71
	pointsOfFunctionPlot[336].Y = 0.604

	pointsOfFunctionPlot[337].X = 0.72
	pointsOfFunctionPlot[337].Y = 0.595

	pointsOfFunctionPlot[338].X = 0.73
	pointsOfFunctionPlot[338].Y = 0.586

	pointsOfFunctionPlot[339].X = 0.74
	pointsOfFunctionPlot[339].Y = 0.578

	pointsOfFunctionPlot[340].X = 0.75
	pointsOfFunctionPlot[340].Y = 0.569

	pointsOfFunctionPlot[341].X = 0.76
	pointsOfFunctionPlot[341].Y = 0.561

	pointsOfFunctionPlot[342].X = 0.77
	pointsOfFunctionPlot[342].Y = 0.552

	pointsOfFunctionPlot[343].X = 0.78
	pointsOfFunctionPlot[343].Y = 0.554

	pointsOfFunctionPlot[344].X = 0.79
	pointsOfFunctionPlot[344].Y = 0.535

	pointsOfFunctionPlot[345].X = 0.80
	pointsOfFunctionPlot[345].Y = 0.527

	pointsOfFunctionPlot[346].X = 0.81
	pointsOfFunctionPlot[346].Y = 0.518

	pointsOfFunctionPlot[347].X = 0.82
	pointsOfFunctionPlot[347].Y = 0.51

	pointsOfFunctionPlot[348].X = 0.83
	pointsOfFunctionPlot[348].Y = 0.502

	pointsOfFunctionPlot[349].X = 0.84
	pointsOfFunctionPlot[349].Y = 0.493

	pointsOfFunctionPlot[350].X = 0.85
	pointsOfFunctionPlot[350].Y = 0.485

	pointsOfFunctionPlot[351].X = 0.86
	pointsOfFunctionPlot[351].Y = 0.477

	pointsOfFunctionPlot[352].X = 0.87
	pointsOfFunctionPlot[352].Y = 0.469

	pointsOfFunctionPlot[353].X = 0.88
	pointsOfFunctionPlot[353].Y = 0.461

	pointsOfFunctionPlot[354].X = 0.89
	pointsOfFunctionPlot[354].Y = 0.452

	pointsOfFunctionPlot[355].X = 0.90
	pointsOfFunctionPlot[355].Y = 0.444

	pointsOfFunctionPlot[356].X = 0.91
	pointsOfFunctionPlot[356].Y = 0.436

	pointsOfFunctionPlot[357].X = 0.92
	pointsOfFunctionPlot[357].Y = 0.429

	pointsOfFunctionPlot[358].X = 0.93
	pointsOfFunctionPlot[358].Y = 0.421

	pointsOfFunctionPlot[359].X = 0.94
	pointsOfFunctionPlot[359].Y = 0.413

	pointsOfFunctionPlot[360].X = 0.95
	pointsOfFunctionPlot[360].Y = 0.405

	pointsOfFunctionPlot[361].X = 0.96
	pointsOfFunctionPlot[361].Y = 0.397

	pointsOfFunctionPlot[362].X = 0.97
	pointsOfFunctionPlot[362].Y = 0.39

	pointsOfFunctionPlot[363].X = 0.98
	pointsOfFunctionPlot[363].Y = 0.382

	pointsOfFunctionPlot[364].X = 0.99
	pointsOfFunctionPlot[364].Y = 0.375

	pointsOfFunctionPlot[365].X = 1.0
	pointsOfFunctionPlot[365].Y = 0.367

	pointsOfFunctionPlot[366].X = 1.01
	pointsOfFunctionPlot[366].Y = 0.36

	pointsOfFunctionPlot[367].X = 1.02
	pointsOfFunctionPlot[367].Y = 0.353

	pointsOfFunctionPlot[368].X = 1.03
	pointsOfFunctionPlot[368].Y = 0.346

	pointsOfFunctionPlot[369].X = 1.04
	pointsOfFunctionPlot[369].Y = 0.339

	pointsOfFunctionPlot[370].X = 1.05
	pointsOfFunctionPlot[370].Y = 0.332

	pointsOfFunctionPlot[371].X = 1.06
	pointsOfFunctionPlot[371].Y = 0.325

	pointsOfFunctionPlot[372].X = 1.07
	pointsOfFunctionPlot[372].Y = 0.318

	pointsOfFunctionPlot[373].X = 1.08
	pointsOfFunctionPlot[373].Y = 0.311

	pointsOfFunctionPlot[374].X = 1.09
	pointsOfFunctionPlot[374].Y = 0.304

	pointsOfFunctionPlot[375].X = 1.10
	pointsOfFunctionPlot[375].Y = 0.298

	pointsOfFunctionPlot[376].X = 1.11
	pointsOfFunctionPlot[376].Y = 0.291

	pointsOfFunctionPlot[377].X = 1.12
	pointsOfFunctionPlot[377].Y = 0.285

	pointsOfFunctionPlot[378].X = 1.13
	pointsOfFunctionPlot[378].Y = 0.278

	pointsOfFunctionPlot[379].X = 1.14
	pointsOfFunctionPlot[379].Y = 0.272

	pointsOfFunctionPlot[380].X = 1.15
	pointsOfFunctionPlot[380].Y = 0.266

	pointsOfFunctionPlot[381].X = 1.16
	pointsOfFunctionPlot[381].Y = 0.26

	pointsOfFunctionPlot[382].X = 1.17
	pointsOfFunctionPlot[382].Y = 0.254

	pointsOfFunctionPlot[383].X = 1.18
	pointsOfFunctionPlot[383].Y = 0.248

	pointsOfFunctionPlot[384].X = 1.19
	pointsOfFunctionPlot[384].Y = 0.242

	pointsOfFunctionPlot[385].X = 1.20
	pointsOfFunctionPlot[385].Y = 0.236

	pointsOfFunctionPlot[386].X = 1.21
	pointsOfFunctionPlot[386].Y = 0.231

	pointsOfFunctionPlot[387].X = 1.22
	pointsOfFunctionPlot[387].Y = 0.225

	pointsOfFunctionPlot[388].X = 1.23
	pointsOfFunctionPlot[388].Y = 0.22

	pointsOfFunctionPlot[389].X = 1.24
	pointsOfFunctionPlot[389].Y = 0.214

	pointsOfFunctionPlot[390].X = 1.25
	pointsOfFunctionPlot[390].Y = 0.209

	pointsOfFunctionPlot[391].X = 1.26
	pointsOfFunctionPlot[391].Y = 0.204

	pointsOfFunctionPlot[392].X = 1.27
	pointsOfFunctionPlot[392].Y = 0.199

	pointsOfFunctionPlot[393].X = 1.28
	pointsOfFunctionPlot[393].Y = 0.194

	pointsOfFunctionPlot[394].X = 1.29
	pointsOfFunctionPlot[394].Y = 0.189

	pointsOfFunctionPlot[395].X = 1.30
	pointsOfFunctionPlot[395].Y = 0.184

	pointsOfFunctionPlot[396].X = 1.31
	pointsOfFunctionPlot[396].Y = 0.179

	pointsOfFunctionPlot[397].X = 1.32
	pointsOfFunctionPlot[397].Y = 0.175

	pointsOfFunctionPlot[398].X = 1.33
	pointsOfFunctionPlot[398].Y = 0.17

	pointsOfFunctionPlot[399].X = 1.34
	pointsOfFunctionPlot[399].Y = 0.166

	pointsOfFunctionPlot[400].X = 1.35
	pointsOfFunctionPlot[400].Y = 0.161

	pointsOfFunctionPlot[401].X = 1.36
	pointsOfFunctionPlot[401].Y = 0.157

	pointsOfFunctionPlot[402].X = 1.37
	pointsOfFunctionPlot[402].Y = 0.153

	pointsOfFunctionPlot[403].X = 1.38
	pointsOfFunctionPlot[403].Y = 0.148

	pointsOfFunctionPlot[404].X = 1.39
	pointsOfFunctionPlot[404].Y = 0.144

	pointsOfFunctionPlot[405].X = 1.40
	pointsOfFunctionPlot[405].Y = 0.14

	pointsOfFunctionPlot[406].X = 1.41
	pointsOfFunctionPlot[406].Y = 0.137

	pointsOfFunctionPlot[407].X = 1.42
	pointsOfFunctionPlot[407].Y = 0.133

	pointsOfFunctionPlot[408].X = 1.43
	pointsOfFunctionPlot[408].Y = 0.129

	pointsOfFunctionPlot[409].X = 1.44
	pointsOfFunctionPlot[409].Y = 0.125

	pointsOfFunctionPlot[410].X = 1.45
	pointsOfFunctionPlot[410].Y = 0.122

	pointsOfFunctionPlot[411].X = 1.46
	pointsOfFunctionPlot[411].Y = 0.118

	pointsOfFunctionPlot[412].X = 1.47
	pointsOfFunctionPlot[412].Y = 0.115

	pointsOfFunctionPlot[413].X = 1.48
	pointsOfFunctionPlot[413].Y = 0.111

	pointsOfFunctionPlot[414].X = 1.49
	pointsOfFunctionPlot[414].Y = 0.108

	pointsOfFunctionPlot[415].X = 1.50
	pointsOfFunctionPlot[415].Y = 0.105

	pointsOfFunctionPlot[416].X = 1.51
	pointsOfFunctionPlot[416].Y = 0.102

	pointsOfFunctionPlot[417].X = 1.52
	pointsOfFunctionPlot[417].Y = 0.099

	pointsOfFunctionPlot[418].X = 1.53
	pointsOfFunctionPlot[418].Y = 0.096

	pointsOfFunctionPlot[419].X = 1.54
	pointsOfFunctionPlot[419].Y = 0.093

	pointsOfFunctionPlot[420].X = 1.55
	pointsOfFunctionPlot[420].Y = 0.09

	pointsOfFunctionPlot[421].X = 1.56
	pointsOfFunctionPlot[421].Y = 0.087

	pointsOfFunctionPlot[422].X = 1.57
	pointsOfFunctionPlot[422].Y = 0.085

	pointsOfFunctionPlot[423].X = 1.58
	pointsOfFunctionPlot[423].Y = 0.082

	pointsOfFunctionPlot[424].X = 1.59
	pointsOfFunctionPlot[424].Y = 0.079

	pointsOfFunctionPlot[425].X = 1.60
	pointsOfFunctionPlot[425].Y = 0.077

	pointsOfFunctionPlot[426].X = 1.61
	pointsOfFunctionPlot[426].Y = 0.074

	pointsOfFunctionPlot[427].X = 1.62
	pointsOfFunctionPlot[427].Y = 0.072

	pointsOfFunctionPlot[428].X = 1.63
	pointsOfFunctionPlot[428].Y = 0.07

	pointsOfFunctionPlot[429].X = 1.64
	pointsOfFunctionPlot[429].Y = 0.067

	pointsOfFunctionPlot[430].X = 1.65
	pointsOfFunctionPlot[430].Y = 0.065

	pointsOfFunctionPlot[431].X = 1.66
	pointsOfFunctionPlot[431].Y = 0.063

	pointsOfFunctionPlot[432].X = 1.67
	pointsOfFunctionPlot[432].Y = 0.061

	pointsOfFunctionPlot[433].X = 1.68
	pointsOfFunctionPlot[433].Y = 0.059

	pointsOfFunctionPlot[434].X = 1.69
	pointsOfFunctionPlot[434].Y = 0.057

	pointsOfFunctionPlot[435].X = 1.70
	pointsOfFunctionPlot[435].Y = 0.055

	pointsOfFunctionPlot[436].X = 1.71
	pointsOfFunctionPlot[436].Y = 0.053

	pointsOfFunctionPlot[437].X = 1.72
	pointsOfFunctionPlot[437].Y = 0.051

	pointsOfFunctionPlot[438].X = 1.73
	pointsOfFunctionPlot[438].Y = 0.05

	pointsOfFunctionPlot[439].X = 1.74
	pointsOfFunctionPlot[439].Y = 0.048

	pointsOfFunctionPlot[440].X = 1.75
	pointsOfFunctionPlot[440].Y = 0.046

	pointsOfFunctionPlot[441].X = 1.76
	pointsOfFunctionPlot[441].Y = 0.045

	pointsOfFunctionPlot[442].X = 1.77
	pointsOfFunctionPlot[442].Y = 0.043

	pointsOfFunctionPlot[443].X = 1.78
	pointsOfFunctionPlot[443].Y = 0.042

	pointsOfFunctionPlot[444].X = 1.79
	pointsOfFunctionPlot[444].Y = 0.04

	pointsOfFunctionPlot[445].X = 1.80
	pointsOfFunctionPlot[445].Y = 0.039

	pointsOfFunctionPlot[446].X = 1.81
	pointsOfFunctionPlot[446].Y = 0.037

	pointsOfFunctionPlot[447].X = 1.82
	pointsOfFunctionPlot[447].Y = 0.036

	pointsOfFunctionPlot[448].X = 1.83
	pointsOfFunctionPlot[448].Y = 0.035

	pointsOfFunctionPlot[449].X = 1.84
	pointsOfFunctionPlot[449].Y = 0.033

	pointsOfFunctionPlot[450].X = 1.85
	pointsOfFunctionPlot[450].Y = 0.032

	pointsOfFunctionPlot[451].X = 1.86
	pointsOfFunctionPlot[451].Y = 0.031

	pointsOfFunctionPlot[452].X = 1.87
	pointsOfFunctionPlot[452].Y = 0.03

	pointsOfFunctionPlot[453].X = 1.88
	pointsOfFunctionPlot[453].Y = 0.029

	pointsOfFunctionPlot[454].X = 1.89
	pointsOfFunctionPlot[454].Y = 0.028

	pointsOfFunctionPlot[455].X = 1.90
	pointsOfFunctionPlot[455].Y = 0.027

	pointsOfFunctionPlot[456].X = 1.91
	pointsOfFunctionPlot[456].Y = 0.026

	pointsOfFunctionPlot[457].X = 1.92
	pointsOfFunctionPlot[457].Y = 0.025

	pointsOfFunctionPlot[458].X = 1.93
	pointsOfFunctionPlot[458].Y = 0.024

	pointsOfFunctionPlot[459].X = 1.94
	pointsOfFunctionPlot[459].Y = 0.023

	pointsOfFunctionPlot[460].X = 1.95
	pointsOfFunctionPlot[460].Y = 0.022

	pointsOfFunctionPlot[461].X = 1.96
	pointsOfFunctionPlot[461].Y = 0.021

	pointsOfFunctionPlot[462].X = 1.97
	pointsOfFunctionPlot[462].Y = 0.02

	pointsOfFunctionPlot[463].X = 1.98
	pointsOfFunctionPlot[463].Y = 0.019

	pointsOfFunctionPlot[463].X = 1.99
	pointsOfFunctionPlot[463].Y = 0.019

	pointsOfFunctionPlot[464].X = 2.0
	pointsOfFunctionPlot[464].Y = 0.018

	pointsOfFunctionPlot[465].X = 2.01
	pointsOfFunctionPlot[465].Y = 0.017

	pointsOfFunctionPlot[466].X = 2.02
	pointsOfFunctionPlot[466].Y = 0.016

	pointsOfFunctionPlot[467].X = 2.03
	pointsOfFunctionPlot[467].Y = 0.016

	pointsOfFunctionPlot[468].X = 2.04
	pointsOfFunctionPlot[468].Y = 0.015

	pointsOfFunctionPlot[469].X = 2.05
	pointsOfFunctionPlot[469].Y = 0.015

	pointsOfFunctionPlot[470].X = 2.06
	pointsOfFunctionPlot[470].Y = 0.014

	pointsOfFunctionPlot[471].X = 2.07
	pointsOfFunctionPlot[471].Y = 0.013

	pointsOfFunctionPlot[472].X = 2.08
	pointsOfFunctionPlot[472].Y = 0.013

	pointsOfFunctionPlot[473].X = 2.09
	pointsOfFunctionPlot[473].Y = 0.012

	pointsOfFunctionPlot[474].X = 2.10
	pointsOfFunctionPlot[474].Y = 0.012

	pointsOfFunctionPlot[475].X = 2.11
	pointsOfFunctionPlot[475].Y = 0.011

	pointsOfFunctionPlot[476].X = 2.12
	pointsOfFunctionPlot[476].Y = 0.011

	pointsOfFunctionPlot[477].X = 2.13
	pointsOfFunctionPlot[477].Y = 0.01

	pointsOfFunctionPlot[478].X = 2.14
	pointsOfFunctionPlot[478].Y = 0.01

	pointsOfFunctionPlot[479].X = 2.15
	pointsOfFunctionPlot[479].Y = 0.009

	pointsOfFunctionPlot[480].X = 2.16
	pointsOfFunctionPlot[480].Y = 0.009

	pointsOfFunctionPlot[481].X = 2.17
	pointsOfFunctionPlot[481].Y = 0.009

	pointsOfFunctionPlot[482].X = 2.18
	pointsOfFunctionPlot[482].Y = 0.008

	pointsOfFunctionPlot[483].X = 2.19
	pointsOfFunctionPlot[483].Y = 0.008

	pointsOfFunctionPlot[484].X = 2.20
	pointsOfFunctionPlot[484].Y = 0.007

	pointsOfFunctionPlot[485].X = 2.21
	pointsOfFunctionPlot[485].Y = 0.007

	pointsOfFunctionPlot[486].X = 2.22
	pointsOfFunctionPlot[486].Y = 0.007

	pointsOfFunctionPlot[487].X = 2.23
	pointsOfFunctionPlot[487].Y = 0.006

	pointsOfFunctionPlot[488].X = 2.24
	pointsOfFunctionPlot[488].Y = 0.006

	pointsOfFunctionPlot[489].X = 2.25
	pointsOfFunctionPlot[489].Y = 0.006

	pointsOfFunctionPlot[490].X = 2.26
	pointsOfFunctionPlot[490].Y = 0.006

	pointsOfFunctionPlot[491].X = 2.27
	pointsOfFunctionPlot[491].Y = 0.005

	pointsOfFunctionPlot[492].X = 2.28
	pointsOfFunctionPlot[492].Y = 0.005

	pointsOfFunctionPlot[492].X = 2.29
	pointsOfFunctionPlot[492].Y = 0.005

	pointsOfFunctionPlot[493].X = 2.30
	pointsOfFunctionPlot[493].Y = 0.005

	pointsOfFunctionPlot[494].X = 2.31
	pointsOfFunctionPlot[494].Y = 0.004

	pointsOfFunctionPlot[495].X = 2.32
	pointsOfFunctionPlot[495].Y = 0.004

	pointsOfFunctionPlot[496].X = 2.33
	pointsOfFunctionPlot[496].Y = 0.004

	pointsOfFunctionPlot[497].X = 2.34
	pointsOfFunctionPlot[497].Y = 0.004

	pointsOfFunctionPlot[498].X = 2.35
	pointsOfFunctionPlot[498].Y = 0.004

	pointsOfFunctionPlot[499].X = 2.36
	pointsOfFunctionPlot[499].Y = 0.003

	pointsOfFunctionPlot[500].X = 2.37
	pointsOfFunctionPlot[500].Y = 0.003

	pointsOfFunctionPlot[501].X = 2.38
	pointsOfFunctionPlot[501].Y = 0.003

	pointsOfFunctionPlot[502].X = 2.39
	pointsOfFunctionPlot[502].Y = 0.003

	pointsOfFunctionPlot[503].X = 2.40
	pointsOfFunctionPlot[503].Y = 0.003

	pointsOfFunctionPlot[504].X = 2.41
	pointsOfFunctionPlot[504].Y = 0.003

	pointsOfFunctionPlot[505].X = 2.42
	pointsOfFunctionPlot[505].Y = 0.002

	pointsOfFunctionPlot[506].X = 2.43
	pointsOfFunctionPlot[506].Y = 0.002

	pointsOfFunctionPlot[507].X = 2.44
	pointsOfFunctionPlot[507].Y = 0.002

	pointsOfFunctionPlot[508].X = 2.45
	pointsOfFunctionPlot[508].Y = 0.002

	pointsOfFunctionPlot[509].X = 2.46
	pointsOfFunctionPlot[509].Y = 0.002

	pointsOfFunctionPlot[510].X = 2.47
	pointsOfFunctionPlot[510].Y = 0.002

	pointsOfFunctionPlot[511].X = 2.48
	pointsOfFunctionPlot[511].Y = 0.002

	pointsOfFunctionPlot[512].X = 2.49
	pointsOfFunctionPlot[512].Y = 0.002

	pointsOfFunctionPlot[513].X = 2.50
	pointsOfFunctionPlot[513].Y = 0.001

	pointsOfFunctionPlot[514].X = 2.51
	pointsOfFunctionPlot[514].Y = 0.001

	pointsOfFunctionPlot[515].X = 2.52
	pointsOfFunctionPlot[515].Y = 0.001

	pointsOfFunctionPlot[516].X = 2.53
	pointsOfFunctionPlot[516].Y = 0.001

	pointsOfFunctionPlot[517].X = 2.54
	pointsOfFunctionPlot[517].Y = 0.001

	pointsOfFunctionPlot[518].X = 2.55
	pointsOfFunctionPlot[518].Y = 0.001

	pointsOfFunctionPlot[519].X = 2.56
	pointsOfFunctionPlot[519].Y = 0.001

	pointsOfFunctionPlot[520].X = 2.57
	pointsOfFunctionPlot[520].Y = 0.001

	pointsOfFunctionPlot[521].X = 2.58
	pointsOfFunctionPlot[521].Y = 0.001

	pointsOfFunctionPlot[522].X = 2.59
	pointsOfFunctionPlot[522].Y = 0.001

	pointsOfFunctionPlot[523].X = 2.60
	pointsOfFunctionPlot[523].Y = 0.001

	pointsOfFunctionPlot[524].X = 2.61
	pointsOfFunctionPlot[524].Y = 0.001

	pointsOfFunctionPlot[525].X = 2.62
	pointsOfFunctionPlot[525].Y = 0.001

	pointsOfFunctionPlot[526].X = 2.63
	pointsOfFunctionPlot[526].Y = 0.001

	pointsOfFunctionPlot[527].X = 2.64
	pointsOfFunctionPlot[527].Y = 0.0

	pointsOfFunctionPlot[528].X = 10.0
	pointsOfFunctionPlot[528].Y = 0.0








	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function exp(-x^2)"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("exp(-x^2)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"Gauss-function-plot-01.png"); err != nil {

		panic(err)
	}
}
