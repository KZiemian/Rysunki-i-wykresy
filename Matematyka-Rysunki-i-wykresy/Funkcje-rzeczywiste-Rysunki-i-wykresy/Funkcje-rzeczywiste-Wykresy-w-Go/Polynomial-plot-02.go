package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of polynomial f(x) = x^4 - x^3 - x^2 + x.

	pointsOfPolynomialPlot := make(plotter.XYs, 1_001)

	pointsOfPolynomialPlot[0].X = 0.0
	pointsOfPolynomialPlot[0].Y = 0.0

	pointsOfPolynomialPlot[1].X = 0.01
	pointsOfPolynomialPlot[1].Y = 0.009

	pointsOfPolynomialPlot[2].X = 0.02
	pointsOfPolynomialPlot[2].Y = 0.019

	pointsOfPolynomialPlot[3].X = 0.03
	pointsOfPolynomialPlot[3].Y = 0.029

	pointsOfPolynomialPlot[4].X = 0.04
	pointsOfPolynomialPlot[4].Y = 0.038

	pointsOfPolynomialPlot[5].X = 0.05
	pointsOfPolynomialPlot[5].Y = 0.047

	pointsOfPolynomialPlot[6].X = 0.06
	pointsOfPolynomialPlot[6].Y = 0.056

	pointsOfPolynomialPlot[7].X = 0.07
	pointsOfPolynomialPlot[7].Y = 0.064

	pointsOfPolynomialPlot[8].X = 0.08
	pointsOfPolynomialPlot[8].Y = 0.073

	pointsOfPolynomialPlot[9].X = 0.09
	pointsOfPolynomialPlot[9].Y = 0.081

	pointsOfPolynomialPlot[10].X = 0.1
	pointsOfPolynomialPlot[10].Y = 0.089

	pointsOfPolynomialPlot[11].X = 0.11
	pointsOfPolynomialPlot[11].Y = 0.096

	pointsOfPolynomialPlot[12].X = 0.12
	pointsOfPolynomialPlot[12].Y = 0.104

	pointsOfPolynomialPlot[13].X = 0.13
	pointsOfPolynomialPlot[13].Y = 0.111

	pointsOfPolynomialPlot[14].X = 0.14
	pointsOfPolynomialPlot[14].Y = 0.118

	pointsOfPolynomialPlot[15].X = 0.15
	pointsOfPolynomialPlot[15].Y = 0.124

	pointsOfPolynomialPlot[16].X = 0.16
	pointsOfPolynomialPlot[16].Y = 0.131

	pointsOfPolynomialPlot[17].X = 0.17
	pointsOfPolynomialPlot[17].Y = 0.137

	pointsOfPolynomialPlot[18].X = 0.18
	pointsOfPolynomialPlot[18].Y = 0.142

	pointsOfPolynomialPlot[19].X = 0.19
	pointsOfPolynomialPlot[19].Y = 0.148

	pointsOfPolynomialPlot[20].X = 0.2
	pointsOfPolynomialPlot[20].Y = 0.153

	pointsOfPolynomialPlot[21].X = 0.21
	pointsOfPolynomialPlot[21].Y = 0.158

	pointsOfPolynomialPlot[22].X = 0.22
	pointsOfPolynomialPlot[22].Y = 0.163

	pointsOfPolynomialPlot[23].X = 0.23
	pointsOfPolynomialPlot[23].Y = 0.167

	pointsOfPolynomialPlot[24].X = 0.24
	pointsOfPolynomialPlot[24].Y = 0.171

	pointsOfPolynomialPlot[25].X = 0.25
	pointsOfPolynomialPlot[25].Y = 0.175

	pointsOfPolynomialPlot[26].X = 0.26
	pointsOfPolynomialPlot[26].Y = 0.179

	pointsOfPolynomialPlot[27].X = 0.27
	pointsOfPolynomialPlot[27].Y = 0.182

	pointsOfPolynomialPlot[28].X = 0.28
	pointsOfPolynomialPlot[28].Y = 0.185

	pointsOfPolynomialPlot[29].X = 0.29
	pointsOfPolynomialPlot[29].Y = 0.188

	pointsOfPolynomialPlot[30].X = 0.3
	pointsOfPolynomialPlot[30].Y = 0.191

	pointsOfPolynomialPlot[31].X = 0.31
	pointsOfPolynomialPlot[31].Y = 0.193

	pointsOfPolynomialPlot[32].X = 0.32
	pointsOfPolynomialPlot[32].Y = 0.195

	pointsOfPolynomialPlot[33].X = 0.33
	pointsOfPolynomialPlot[33].Y = 0.197

	pointsOfPolynomialPlot[34].X = 0.34
	pointsOfPolynomialPlot[34].Y = 0.198

	pointsOfPolynomialPlot[35].X = 0.35
	pointsOfPolynomialPlot[35].Y = 0.199

	pointsOfPolynomialPlot[36].X = 0.36
	pointsOfPolynomialPlot[36].Y = 0.2

	pointsOfPolynomialPlot[37].X = 0.37
	pointsOfPolynomialPlot[37].Y = 0.201

	pointsOfPolynomialPlot[38].X = 0.38
	pointsOfPolynomialPlot[38].Y = 0.201

	pointsOfPolynomialPlot[39].X = 0.39
	pointsOfPolynomialPlot[39].Y = 0.201

	pointsOfPolynomialPlot[40].X = 0.4
	pointsOfPolynomialPlot[40].Y = 0.201

	pointsOfPolynomialPlot[41].X = 0.41
	pointsOfPolynomialPlot[41].Y = 0.201

	pointsOfPolynomialPlot[42].X = 0.42
	pointsOfPolynomialPlot[42].Y = 0.2

	pointsOfPolynomialPlot[43].X = 0.43
	pointsOfPolynomialPlot[43].Y = 0.199

	pointsOfPolynomialPlot[44].X = 0.44
	pointsOfPolynomialPlot[44].Y = 0.198

	pointsOfPolynomialPlot[45].X = 0.45
	pointsOfPolynomialPlot[45].Y = 0.197

	pointsOfPolynomialPlot[46].X = 0.46
	pointsOfPolynomialPlot[46].Y = 0.195

	pointsOfPolynomialPlot[47].X = 0.47
	pointsOfPolynomialPlot[47].Y = 0.194

	pointsOfPolynomialPlot[48].X = 0.48
	pointsOfPolynomialPlot[48].Y = 0.192

	pointsOfPolynomialPlot[49].X = 0.49
	pointsOfPolynomialPlot[49].Y = 0.189

	pointsOfPolynomialPlot[50].X = 0.5
	pointsOfPolynomialPlot[50].Y = 0.187

	pointsOfPolynomialPlot[51].X = 0.51
	pointsOfPolynomialPlot[51].Y = 0.184

	pointsOfPolynomialPlot[52].X = 0.52
	pointsOfPolynomialPlot[52].Y = 0.182

	pointsOfPolynomialPlot[53].X = 0.53
	pointsOfPolynomialPlot[53].Y = 0.179

	pointsOfPolynomialPlot[54].X = 0.54
	pointsOfPolynomialPlot[54].Y = 0.176

	pointsOfPolynomialPlot[55].X = 0.55
	pointsOfPolynomialPlot[55].Y = 0.172

	pointsOfPolynomialPlot[56].X = 0.56
	pointsOfPolynomialPlot[56].Y = 0.169

	pointsOfPolynomialPlot[57].X = 0.57
	pointsOfPolynomialPlot[57].Y = 0.165

	pointsOfPolynomialPlot[58].X = 0.58
	pointsOfPolynomialPlot[58].Y = 0.161

	pointsOfPolynomialPlot[59].X = 0.59
	pointsOfPolynomialPlot[59].Y = 0.157

	pointsOfPolynomialPlot[60].X = 0.6
	pointsOfPolynomialPlot[60].Y = 0.153

	pointsOfPolynomialPlot[61].X = 0.61
	pointsOfPolynomialPlot[61].Y = 0.149

	pointsOfPolynomialPlot[62].X = 0.62
	pointsOfPolynomialPlot[62].Y = 0.145

	pointsOfPolynomialPlot[63].X = 0.63
	pointsOfPolynomialPlot[63].Y = 0.14

	pointsOfPolynomialPlot[64].X = 0.64
	pointsOfPolynomialPlot[64].Y = 0.136

	pointsOfPolynomialPlot[65].X = 0.65
	pointsOfPolynomialPlot[65].Y = 0.131

	pointsOfPolynomialPlot[66].X = 0.66
	pointsOfPolynomialPlot[66].Y = 0.126

	pointsOfPolynomialPlot[67].X = 0.67
	pointsOfPolynomialPlot[67].Y = 0.121

	pointsOfPolynomialPlot[68].X = 0.68
	pointsOfPolynomialPlot[68].Y = 0.117

	pointsOfPolynomialPlot[69].X = 0.69
	pointsOfPolynomialPlot[69].Y = 0.112

	pointsOfPolynomialPlot[70].X = 0.7
	pointsOfPolynomialPlot[70].Y = 0.107

	pointsOfPolynomialPlot[71].X = 0.71
	pointsOfPolynomialPlot[71].Y = 0.102

	pointsOfPolynomialPlot[72].X = 0.72
	pointsOfPolynomialPlot[72].Y = 0.097

	pointsOfPolynomialPlot[73].X = 0.73
	pointsOfPolynomialPlot[73].Y = 0.092

	pointsOfPolynomialPlot[74].X = 0.74
	pointsOfPolynomialPlot[74].Y = 0.087

	pointsOfPolynomialPlot[75].X = 0.75
	pointsOfPolynomialPlot[75].Y = 0.082

	pointsOfPolynomialPlot[76].X = 0.76
	pointsOfPolynomialPlot[76].Y = 0.077

	pointsOfPolynomialPlot[77].X = 0.77
	pointsOfPolynomialPlot[77].Y = 0.072

	pointsOfPolynomialPlot[78].X = 0.78
	pointsOfPolynomialPlot[78].Y = 0.067

	pointsOfPolynomialPlot[79].X = 0.79
	pointsOfPolynomialPlot[79].Y = 0.062

	pointsOfPolynomialPlot[80].X = 0.8
	pointsOfPolynomialPlot[80].Y = 0.057

	pointsOfPolynomialPlot[81].X = 0.81
	pointsOfPolynomialPlot[81].Y = 0.052

	pointsOfPolynomialPlot[82].X = 0.82
	pointsOfPolynomialPlot[82].Y = 0.048

	pointsOfPolynomialPlot[83].X = 0.83
	pointsOfPolynomialPlot[83].Y = 0.043

	pointsOfPolynomialPlot[84].X = 0.84
	pointsOfPolynomialPlot[84].Y = 0.039

	pointsOfPolynomialPlot[85].X = 0.85
	pointsOfPolynomialPlot[85].Y = 0.035

	pointsOfPolynomialPlot[86].X = 0.86
	pointsOfPolynomialPlot[86].Y = 0.031

	pointsOfPolynomialPlot[87].X = 0.87
	pointsOfPolynomialPlot[87].Y = 0.027

	pointsOfPolynomialPlot[88].X = 0.88
	pointsOfPolynomialPlot[88].Y = 0.023

	pointsOfPolynomialPlot[89].X = 0.89
	pointsOfPolynomialPlot[89].Y = 0.02

	pointsOfPolynomialPlot[90].X = 0.9
	pointsOfPolynomialPlot[90].Y = 0.017

	pointsOfPolynomialPlot[91].X = 0.91
	pointsOfPolynomialPlot[91].Y = 0.014

	pointsOfPolynomialPlot[92].X = 0.92
	pointsOfPolynomialPlot[92].Y = 0.011

	pointsOfPolynomialPlot[93].X = 0.93
	pointsOfPolynomialPlot[93].Y = 0.008

	pointsOfPolynomialPlot[94].X = 0.94
	pointsOfPolynomialPlot[94].Y = 0.006

	pointsOfPolynomialPlot[95].X = 0.95
	pointsOfPolynomialPlot[95].Y = 0.004

	pointsOfPolynomialPlot[96].X = 0.96
	pointsOfPolynomialPlot[96].Y = 0.003

	pointsOfPolynomialPlot[97].X = 0.97
	pointsOfPolynomialPlot[97].Y = 0.001

	pointsOfPolynomialPlot[98].X = 0.98
	pointsOfPolynomialPlot[98].Y = 0.0

	pointsOfPolynomialPlot[99].X = 0.99
	pointsOfPolynomialPlot[99].Y = 0.0

	pointsOfPolynomialPlot[100].X = 1.0
	pointsOfPolynomialPlot[100].Y = 0.0

	pointsOfPolynomialPlot[101].X = 1.01
	pointsOfPolynomialPlot[101].Y = 0.0

	pointsOfPolynomialPlot[102].X = 1.02
	pointsOfPolynomialPlot[102].Y = 0.0

	pointsOfPolynomialPlot[103].X = 1.03
	pointsOfPolynomialPlot[103].Y = 0.001

	pointsOfPolynomialPlot[104].X = 1.04
	pointsOfPolynomialPlot[104].Y = 0.003

	pointsOfPolynomialPlot[105].X = 1.05
	pointsOfPolynomialPlot[105].Y = 0.005

	pointsOfPolynomialPlot[106].X = 1.06
	pointsOfPolynomialPlot[106].Y = 0.007

	pointsOfPolynomialPlot[107].X = 1.07
	pointsOfPolynomialPlot[107].Y = 0.01

	pointsOfPolynomialPlot[108].X = 1.08
	pointsOfPolynomialPlot[108].Y = 0.014

	pointsOfPolynomialPlot[109].X = 1.09
	pointsOfPolynomialPlot[109].Y = 0.018

	pointsOfPolynomialPlot[110].X = 1.1
	pointsOfPolynomialPlot[110].Y = 0.023

	pointsOfPolynomialPlot[111].X = 1.11
	pointsOfPolynomialPlot[111].Y = 0.028

	pointsOfPolynomialPlot[112].X = 1.12
	pointsOfPolynomialPlot[112].Y = 0.034

	pointsOfPolynomialPlot[113].X = 1.13
	pointsOfPolynomialPlot[113].Y = 0.04

	pointsOfPolynomialPlot[114].X = 1.14
	pointsOfPolynomialPlot[114].Y = 0.047

	pointsOfPolynomialPlot[115].X = 1.15
	pointsOfPolynomialPlot[115].Y = 0.055

	pointsOfPolynomialPlot[116].X = 1.16
	pointsOfPolynomialPlot[116].Y = 0.064

	pointsOfPolynomialPlot[117].X = 1.17
	pointsOfPolynomialPlot[117].Y = 0.073

	pointsOfPolynomialPlot[118].X = 1.18
	pointsOfPolynomialPlot[118].Y = 0.083

	pointsOfPolynomialPlot[119].X = 1.19
	pointsOfPolynomialPlot[119].Y = 0.094

	pointsOfPolynomialPlot[120].X = 1.2
	pointsOfPolynomialPlot[120].Y = 0.105

	pointsOfPolynomialPlot[121].X = 1.21
	pointsOfPolynomialPlot[121].Y = 0.117

	pointsOfPolynomialPlot[122].X = 1.22
	pointsOfPolynomialPlot[122].Y = 0.131

	pointsOfPolynomialPlot[123].X = 1.23
	pointsOfPolynomialPlot[123].Y = 0.145

	pointsOfPolynomialPlot[124].X = 1.24
	pointsOfPolynomialPlot[124].Y = 0.16

	pointsOfPolynomialPlot[125].X = 1.25
	pointsOfPolynomialPlot[125].Y = 0.175

	pointsOfPolynomialPlot[126].X = 1.26
	pointsOfPolynomialPlot[126].Y = 0.192

	pointsOfPolynomialPlot[127].X = 1.27
	pointsOfPolynomialPlot[127].Y = 0.21

	pointsOfPolynomialPlot[128].X = 1.28
	pointsOfPolynomialPlot[128].Y = 0.228

	pointsOfPolynomialPlot[129].X = 1.29
	pointsOfPolynomialPlot[129].Y = 0.248

	pointsOfPolynomialPlot[130].X = 1.3
	pointsOfPolynomialPlot[130].Y = 0.269

	pointsOfPolynomialPlot[131].X = 1.31
	pointsOfPolynomialPlot[131].Y = 0.29

	pointsOfPolynomialPlot[132].X = 1.32
	pointsOfPolynomialPlot[132].Y = 0.313

	pointsOfPolynomialPlot[133].X = 1.33
	pointsOfPolynomialPlot[133].Y = 0.337

	pointsOfPolynomialPlot[134].X = 1.34
	pointsOfPolynomialPlot[134].Y = 0.362

	pointsOfPolynomialPlot[135].X = 1.35
	pointsOfPolynomialPlot[135].Y = 0.388

	pointsOfPolynomialPlot[136].X = 1.36
	pointsOfPolynomialPlot[136].Y = 0.416

	pointsOfPolynomialPlot[137].X = 1.37
	pointsOfPolynomialPlot[137].Y = 0.444

	pointsOfPolynomialPlot[138].X = 1.38
	pointsOfPolynomialPlot[138].Y = 0.474

	pointsOfPolynomialPlot[139].X = 1.39
	pointsOfPolynomialPlot[139].Y = 0.505

	pointsOfPolynomialPlot[140].X = 1.4
	pointsOfPolynomialPlot[140].Y = 0.537

	pointsOfPolynomialPlot[141].X = 1.41
	pointsOfPolynomialPlot[141].Y = 0.571

	pointsOfPolynomialPlot[142].X = 1.42
	pointsOfPolynomialPlot[142].Y = 0.606

	pointsOfPolynomialPlot[143].X = 1.43
	pointsOfPolynomialPlot[143].Y = 0.642

	pointsOfPolynomialPlot[144].X = 1.44
	pointsOfPolynomialPlot[144].Y = 0.68

	pointsOfPolynomialPlot[145].X = 1.45
	pointsOfPolynomialPlot[145].Y = 0.719

	pointsOfPolynomialPlot[146].X = 1.46
	pointsOfPolynomialPlot[146].Y = 0.76

	pointsOfPolynomialPlot[147].X = 1.47
	pointsOfPolynomialPlot[147].Y = 0.802

	pointsOfPolynomialPlot[148].X = 1.48
	pointsOfPolynomialPlot[148].Y = 0.845

	pointsOfPolynomialPlot[149].X = 1.49
	pointsOfPolynomialPlot[149].Y = 0.89

	pointsOfPolynomialPlot[150].X = 1.5
	pointsOfPolynomialPlot[150].Y = 0.937

	pointsOfPolynomialPlot[151].X = 1.51
	pointsOfPolynomialPlot[151].Y = 0.985

	pointsOfPolynomialPlot[152].X = 1.52
	pointsOfPolynomialPlot[152].Y = 1.035

	pointsOfPolynomialPlot[153].X = 1.53
	pointsOfPolynomialPlot[153].Y = 1.087

	pointsOfPolynomialPlot[154].X = 1.54
	pointsOfPolynomialPlot[154].Y = 1.14

	pointsOfPolynomialPlot[155].X = 1.55
	pointsOfPolynomialPlot[155].Y = 1.195

	pointsOfPolynomialPlot[156].X = 1.56
	pointsOfPolynomialPlot[156].Y = 1.252

	pointsOfPolynomialPlot[157].X = 1.57
	pointsOfPolynomialPlot[157].Y = 1.31

	pointsOfPolynomialPlot[158].X = 1.58
	pointsOfPolynomialPlot[158].Y = 1.371

	pointsOfPolynomialPlot[159].X = 1.59
	pointsOfPolynomialPlot[159].Y = 1.433

	pointsOfPolynomialPlot[160].X = 1.6
	pointsOfPolynomialPlot[160].Y = 1.497

	pointsOfPolynomialPlot[161].X = 1.61
	pointsOfPolynomialPlot[161].Y = 1.563

	pointsOfPolynomialPlot[162].X = 1.62
	pointsOfPolynomialPlot[162].Y = 1.631

	pointsOfPolynomialPlot[163].X = 1.63
	pointsOfPolynomialPlot[163].Y = 1.701

	pointsOfPolynomialPlot[164].X = 1.64
	pointsOfPolynomialPlot[164].Y = 1.773

	pointsOfPolynomialPlot[165].X = 1.65
	pointsOfPolynomialPlot[165].Y = 1.847

	pointsOfPolynomialPlot[166].X = 1.66
	pointsOfPolynomialPlot[166].Y = 1.923

	pointsOfPolynomialPlot[167].X = 1.67
	pointsOfPolynomialPlot[167].Y = 2.011

	pointsOfPolynomialPlot[168].X = 1.68
	pointsOfPolynomialPlot[168].Y = 2.081

	pointsOfPolynomialPlot[169].X = 1.69
	pointsOfPolynomialPlot[169].Y = 2.164

	pointsOfPolynomialPlot[170].X = 1.7
	pointsOfPolynomialPlot[170].Y = 2.249

	pointsOfPolynomialPlot[171].X = 1.71
	pointsOfPolynomialPlot[171].Y = 2.336

	pointsOfPolynomialPlot[172].X = 1.72
	pointsOfPolynomialPlot[172].Y = 2.425

	pointsOfPolynomialPlot[173].X = 1.73
	pointsOfPolynomialPlot[173].Y = 2.516

	pointsOfPolynomialPlot[174].X = 1.74
	pointsOfPolynomialPlot[174].Y = 2.61

	pointsOfPolynomialPlot[175].X = 1.75
	pointsOfPolynomialPlot[175].Y = 2.707

	pointsOfPolynomialPlot[176].X = 1.76
	pointsOfPolynomialPlot[176].Y = 2.805

	pointsOfPolynomialPlot[177].X = 1.77
	pointsOfPolynomialPlot[177].Y = 2.906

	pointsOfPolynomialPlot[178].X = 1.78
	pointsOfPolynomialPlot[178].Y = 3.01

	pointsOfPolynomialPlot[179].X = 1.79
	pointsOfPolynomialPlot[179].Y = 3.116

	pointsOfPolynomialPlot[180].X = 1.8
	pointsOfPolynomialPlot[180].Y = 3.225

	pointsOfPolynomialPlot[181].X = 1.81
	pointsOfPolynomialPlot[181].Y = 3.337

	pointsOfPolynomialPlot[182].X = 1.82
	pointsOfPolynomialPlot[182].Y = 3.451

	pointsOfPolynomialPlot[183].X = 1.83
	pointsOfPolynomialPlot[183].Y = 3.567

	pointsOfPolynomialPlot[184].X = 1.84
	pointsOfPolynomialPlot[184].Y = 3.687

	pointsOfPolynomialPlot[185].X = 1.85
	pointsOfPolynomialPlot[185].Y = 3.809

	pointsOfPolynomialPlot[186].X = 1.86
	pointsOfPolynomialPlot[186].Y = 3.934

	pointsOfPolynomialPlot[187].X = 1.87
	pointsOfPolynomialPlot[187].Y = 4.062

	pointsOfPolynomialPlot[188].X = 1.88
	pointsOfPolynomialPlot[188].Y = 4.192

	pointsOfPolynomialPlot[189].X = 1.89
	pointsOfPolynomialPlot[189].Y = 4.326

	pointsOfPolynomialPlot[190].X = 1.9
	pointsOfPolynomialPlot[190].Y = 4.463

	pointsOfPolynomialPlot[191].X = 1.91
	pointsOfPolynomialPlot[191].Y = 4.602

	pointsOfPolynomialPlot[192].X = 1.92
	pointsOfPolynomialPlot[192].Y = 4.745

	pointsOfPolynomialPlot[193].X = 1.93
	pointsOfPolynomialPlot[193].Y = 4.89

	pointsOfPolynomialPlot[194].X = 1.94
	pointsOfPolynomialPlot[194].Y = 5.039

	pointsOfPolynomialPlot[195].X = 1.95
	pointsOfPolynomialPlot[195].Y = 5.191

	pointsOfPolynomialPlot[196].X = 1.96
	pointsOfPolynomialPlot[196].Y = 5.346

	pointsOfPolynomialPlot[197].X = 1.97
	pointsOfPolynomialPlot[197].Y = 5.505

	pointsOfPolynomialPlot[198].X = 1.98
	pointsOfPolynomialPlot[198].Y = 5.666

	pointsOfPolynomialPlot[199].X = 1.99
	pointsOfPolynomialPlot[199].Y = 5.831

	pointsOfPolynomialPlot[200].X = 2.0
	pointsOfPolynomialPlot[200].Y = 6.0

	pointsOfPolynomialPlot[201].X = 2.01
	pointsOfPolynomialPlot[201].Y = 6.171

	pointsOfPolynomialPlot[202].X = 2.02
	pointsOfPolynomialPlot[202].Y = 6.346

	pointsOfPolynomialPlot[203].X = 2.03
	pointsOfPolynomialPlot[203].Y = 6.525

	pointsOfPolynomialPlot[204].X = 2.04
	pointsOfPolynomialPlot[204].Y = 6.707

	pointsOfPolynomialPlot[205].X = 2.05
	pointsOfPolynomialPlot[205].Y = 6.893

	pointsOfPolynomialPlot[206].X = 2.06
	pointsOfPolynomialPlot[206].Y = 7.082

	pointsOfPolynomialPlot[207].X = 2.07
	pointsOfPolynomialPlot[207].Y = 7.275

	pointsOfPolynomialPlot[208].X = 2.08
	pointsOfPolynomialPlot[208].Y = 7.472

	pointsOfPolynomialPlot[209].X = 2.09
	pointsOfPolynomialPlot[209].Y = 7.672

	pointsOfPolynomialPlot[210].X = 2.1
	pointsOfPolynomialPlot[210].Y = 7.877

	pointsOfPolynomialPlot[211].X = 2.11
	pointsOfPolynomialPlot[211].Y = 8.085

	pointsOfPolynomialPlot[212].X = 2.12
	pointsOfPolynomialPlot[212].Y = 8.297

	pointsOfPolynomialPlot[213].X = 2.13
	pointsOfPolynomialPlot[213].Y = 8.513

	pointsOfPolynomialPlot[214].X = 2.14
	pointsOfPolynomialPlot[214].Y = 8.732

	pointsOfPolynomialPlot[215].X = 2.15
	pointsOfPolynomialPlot[215].Y = 8.956

	pointsOfPolynomialPlot[216].X = 2.16
	pointsOfPolynomialPlot[216].Y = 9.184

	pointsOfPolynomialPlot[217].X = 2.17
	pointsOfPolynomialPlot[217].Y = 9.416

	pointsOfPolynomialPlot[218].X = 2.18
	pointsOfPolynomialPlot[218].Y = 9.652

	pointsOfPolynomialPlot[219].X = 2.19
	pointsOfPolynomialPlot[219].Y = 9.893

	pointsOfPolynomialPlot[220].X = 2.2
	pointsOfPolynomialPlot[220].Y = 10.137

	pointsOfPolynomialPlot[221].X = 2.21
	pointsOfPolynomialPlot[221].Y = 10.386

	pointsOfPolynomialPlot[222].X = 2.22
	pointsOfPolynomialPlot[222].Y = 10.639

	pointsOfPolynomialPlot[223].X = 2.23
	pointsOfPolynomialPlot[223].Y = 10.897

	pointsOfPolynomialPlot[224].X = 2.24
	pointsOfPolynomialPlot[224].Y = 11.159

	pointsOfPolynomialPlot[225].X = 2.25
	pointsOfPolynomialPlot[225].Y = 11.425

	pointsOfPolynomialPlot[226].X = 2.26
	pointsOfPolynomialPlot[226].Y = 11.696

	pointsOfPolynomialPlot[227].X = 2.27
	pointsOfPolynomialPlot[227].Y = 11.972

	pointsOfPolynomialPlot[228].X = 2.28
	pointsOfPolynomialPlot[228].Y = 12.252

	pointsOfPolynomialPlot[229].X = 2.29
	pointsOfPolynomialPlot[229].Y = 12.537

	pointsOfPolynomialPlot[230].X = 2.3
	pointsOfPolynomialPlot[230].Y = 12.827

	pointsOfPolynomialPlot[231].X = 2.31
	pointsOfPolynomialPlot[231].Y = 13.121

	pointsOfPolynomialPlot[232].X = 2.32
	pointsOfPolynomialPlot[232].Y = 13.42

	pointsOfPolynomialPlot[233].X = 2.33
	pointsOfPolynomialPlot[233].Y = 13.724

	pointsOfPolynomialPlot[234].X = 2.34
	pointsOfPolynomialPlot[234].Y = 14.033

	pointsOfPolynomialPlot[235].X = 2.35
	pointsOfPolynomialPlot[235].Y = 14.347

	pointsOfPolynomialPlot[236].X = 2.36
	pointsOfPolynomialPlot[236].Y = 14.666

	pointsOfPolynomialPlot[237].X = 2.37
	pointsOfPolynomialPlot[237].Y = 14.99

	pointsOfPolynomialPlot[238].X = 2.38
	pointsOfPolynomialPlot[238].Y = 15.319

	pointsOfPolynomialPlot[239].X = 2.39
	pointsOfPolynomialPlot[239].Y = 15.654

	pointsOfPolynomialPlot[240].X = 2.4
	pointsOfPolynomialPlot[240].Y = 15.993

	pointsOfPolynomialPlot[241].X = 2.41
	pointsOfPolynomialPlot[241].Y = 16.338

	pointsOfPolynomialPlot[242].X = 2.42
	pointsOfPolynomialPlot[242].Y = 16.688

	pointsOfPolynomialPlot[243].X = 2.43
	pointsOfPolynomialPlot[243].Y = 17.044

	pointsOfPolynomialPlot[244].X = 2.44
	pointsOfPolynomialPlot[244].Y = 17.405

	pointsOfPolynomialPlot[245].X = 2.45
	pointsOfPolynomialPlot[245].Y = 17.771

	pointsOfPolynomialPlot[246].X = 2.46
	pointsOfPolynomialPlot[246].Y = 18.143

	pointsOfPolynomialPlot[247].X = 2.47
	pointsOfPolynomialPlot[247].Y = 18.52

	pointsOfPolynomialPlot[248].X = 2.48
	pointsOfPolynomialPlot[248].Y = 18.904

	pointsOfPolynomialPlot[249].X = 2.49
	pointsOfPolynomialPlot[249].Y = 19.292

	pointsOfPolynomialPlot[250].X = 2.5
	pointsOfPolynomialPlot[250].Y = 19.687

	pointsOfPolynomialPlot[251].X = 2.51
	pointsOfPolynomialPlot[251].Y = 20.087

	pointsOfPolynomialPlot[252].X = 2.52
	pointsOfPolynomialPlot[252].Y = 20.494

	pointsOfPolynomialPlot[253].X = 2.53
	pointsOfPolynomialPlot[253].Y = 20.906

	pointsOfPolynomialPlot[254].X = 2.54
	pointsOfPolynomialPlot[254].Y = 21.324

	pointsOfPolynomialPlot[255].X = 2.55
	pointsOfPolynomialPlot[255].Y = 21.748

	pointsOfPolynomialPlot[256].X = 2.56
	pointsOfPolynomialPlot[256].Y = 22.178

	pointsOfPolynomialPlot[257].X = 2.57
	pointsOfPolynomialPlot[257].Y = 22.615

	pointsOfPolynomialPlot[258].X = 2.58
	pointsOfPolynomialPlot[258].Y = 23.057

	pointsOfPolynomialPlot[259].X = 2.59
	pointsOfPolynomialPlot[259].Y = 23.506

	pointsOfPolynomialPlot[260].X = 2.6
	pointsOfPolynomialPlot[260].Y = 23.961

	pointsOfPolynomialPlot[261].X = 2.61
	pointsOfPolynomialPlot[261].Y = 24.423

	pointsOfPolynomialPlot[262].X = 2.62
	pointsOfPolynomialPlot[262].Y = 24.89

	pointsOfPolynomialPlot[263].X = 2.63
	pointsOfPolynomialPlot[263].Y = 25.365

	pointsOfPolynomialPlot[264].X = 2.64
	pointsOfPolynomialPlot[264].Y = 25.846

	pointsOfPolynomialPlot[265].X = 2.65
	pointsOfPolynomialPlot[265].Y = 26.333

	pointsOfPolynomialPlot[266].X = 2.66
	pointsOfPolynomialPlot[266].Y = 26.827

	pointsOfPolynomialPlot[267].X = 2.67
	pointsOfPolynomialPlot[267].Y = 27.328

	pointsOfPolynomialPlot[268].X = 2.68
	pointsOfPolynomialPlot[268].Y = 27.835

	pointsOfPolynomialPlot[269].X = 2.69
	pointsOfPolynomialPlot[269].Y = 28.349

	pointsOfPolynomialPlot[270].X = 2.7
	pointsOfPolynomialPlot[270].Y = 28.871

	pointsOfPolynomialPlot[271].X = 2.71
	pointsOfPolynomialPlot[271].Y = 29.399

	pointsOfPolynomialPlot[272].X = 2.72
	pointsOfPolynomialPlot[272].Y = 29.934

	pointsOfPolynomialPlot[273].X = 2.73
	pointsOfPolynomialPlot[273].Y = 30.476

	pointsOfPolynomialPlot[274].X = 2.74
	pointsOfPolynomialPlot[274].Y = 31.025

	pointsOfPolynomialPlot[275].X = 2.75
	pointsOfPolynomialPlot[275].Y = 31.582

	pointsOfPolynomialPlot[276].X = 2.76
	pointsOfPolynomialPlot[276].Y = 32.145

	pointsOfPolynomialPlot[277].X = 2.77
	pointsOfPolynomialPlot[277].Y = 32.716

	pointsOfPolynomialPlot[278].X = 2.78
	pointsOfPolynomialPlot[278].Y = 33.294

	pointsOfPolynomialPlot[279].X = 2.79
	pointsOfPolynomialPlot[279].Y = 33.88

	pointsOfPolynomialPlot[280].X = 2.8
	pointsOfPolynomialPlot[280].Y = 34.473

	pointsOfPolynomialPlot[281].X = 2.81
	pointsOfPolynomialPlot[281].Y = 35.074

	pointsOfPolynomialPlot[282].X = 2.82
	pointsOfPolynomialPlot[282].Y = 35.682

	pointsOfPolynomialPlot[283].X = 2.83
	pointsOfPolynomialPlot[283].Y = 36.298

	pointsOfPolynomialPlot[284].X = 2.84
	pointsOfPolynomialPlot[284].Y = 36.922

	pointsOfPolynomialPlot[285].X = 2.85
	pointsOfPolynomialPlot[285].Y = 37.553

	pointsOfPolynomialPlot[286].X = 2.86
	pointsOfPolynomialPlot[286].Y = 38.192

	pointsOfPolynomialPlot[287].X = 2.87
	pointsOfPolynomialPlot[287].Y = 38.839

	pointsOfPolynomialPlot[288].X = 2.88
	pointsOfPolynomialPlot[288].Y = 39.494

	pointsOfPolynomialPlot[289].X = 2.89
	pointsOfPolynomialPlot[289].Y = 40.157

	pointsOfPolynomialPlot[290].X = 2.9
	pointsOfPolynomialPlot[290].Y = 40.829

	pointsOfPolynomialPlot[291].X = 2.91
	pointsOfPolynomialPlot[291].Y = 41.508

	pointsOfPolynomialPlot[292].X = 2.92
	pointsOfPolynomialPlot[292].Y = 42.196

	pointsOfPolynomialPlot[293].X = 2.93
	pointsOfPolynomialPlot[293].Y = 42.891

	pointsOfPolynomialPlot[294].X = 2.94
	pointsOfPolynomialPlot[294].Y = 43.596

	pointsOfPolynomialPlot[295].X = 2.95
	pointsOfPolynomialPlot[295].Y = 44.308

	pointsOfPolynomialPlot[296].X = 2.96
	pointsOfPolynomialPlot[296].Y = 45.029

	pointsOfPolynomialPlot[297].X = 2.97
	pointsOfPolynomialPlot[297].Y = 45.759

	pointsOfPolynomialPlot[298].X = 2.98
	pointsOfPolynomialPlot[298].Y = 46.497

	pointsOfPolynomialPlot[299].X = 2.99
	pointsOfPolynomialPlot[299].Y = 47.244

	pointsOfPolynomialPlot[300].X = 3.0
	pointsOfPolynomialPlot[300].Y = 48.0

	pointsOfPolynomialPlot[301].X = 3.01
	pointsOfPolynomialPlot[301].Y = 48.764

	pointsOfPolynomialPlot[302].X = 3.02
	pointsOfPolynomialPlot[302].Y = 49.537

	pointsOfPolynomialPlot[303].X = 3.03
	pointsOfPolynomialPlot[303].Y = 50.319

	pointsOfPolynomialPlot[304].X = 3.04
	pointsOfPolynomialPlot[304].Y = 51.111

	pointsOfPolynomialPlot[305].X = 3.05
	pointsOfPolynomialPlot[305].Y = 51.911

	pointsOfPolynomialPlot[306].X = 3.06
	pointsOfPolynomialPlot[306].Y = 52.72

	pointsOfPolynomialPlot[307].X = 3.07
	pointsOfPolynomialPlot[307].Y = 53.539

	pointsOfPolynomialPlot[308].X = 3.08
	pointsOfPolynomialPlot[308].Y = 54.367

	pointsOfPolynomialPlot[309].X = 3.09
	pointsOfPolynomialPlot[309].Y = 55.204

	pointsOfPolynomialPlot[310].X = 3.1
	pointsOfPolynomialPlot[310].Y = 56.051

	pointsOfPolynomialPlot[311].X = 3.11
	pointsOfPolynomialPlot[311].Y = 56.907

	pointsOfPolynomialPlot[312].X = 3.12
	pointsOfPolynomialPlot[312].Y = 57.772

	pointsOfPolynomialPlot[313].X = 3.13
	pointsOfPolynomialPlot[313].Y = 58.648

	pointsOfPolynomialPlot[314].X = 3.14
	pointsOfPolynomialPlot[314].Y = 59.533

	pointsOfPolynomialPlot[315].X = 3.15
	pointsOfPolynomialPlot[315].Y = 60.427

	pointsOfPolynomialPlot[316].X = 3.16
	pointsOfPolynomialPlot[316].Y = 61.332

	pointsOfPolynomialPlot[317].X = 3.17
	pointsOfPolynomialPlot[317].Y = 62.246

	pointsOfPolynomialPlot[318].X = 3.18
	pointsOfPolynomialPlot[318].Y = 63.17

	pointsOfPolynomialPlot[319].X = 3.19
	pointsOfPolynomialPlot[319].Y = 64.105

	pointsOfPolynomialPlot[320].X = 3.2
	pointsOfPolynomialPlot[320].Y = 65.049

	pointsOfPolynomialPlot[321].X = 3.21
	pointsOfPolynomialPlot[321].Y = 66.004

	pointsOfPolynomialPlot[322].X = 3.22
	pointsOfPolynomialPlot[322].Y = 66.969

	pointsOfPolynomialPlot[323].X = 3.23
	pointsOfPolynomialPlot[323].Y = 67.944

	pointsOfPolynomialPlot[324].X = 3.24
	pointsOfPolynomialPlot[324].Y = 68.929

	pointsOfPolynomialPlot[325].X = 3.25
	pointsOfPolynomialPlot[325].Y = 69.925

	pointsOfPolynomialPlot[326].X = 3.26
	pointsOfPolynomialPlot[326].Y = 70.932

	pointsOfPolynomialPlot[327].X = 3.27
	pointsOfPolynomialPlot[327].Y = 71.949

	pointsOfPolynomialPlot[328].X = 3.28
	pointsOfPolynomialPlot[328].Y = 72.977

	pointsOfPolynomialPlot[329].X = 3.29
	pointsOfPolynomialPlot[329].Y = 74.015

	pointsOfPolynomialPlot[330].X = 3.3
	pointsOfPolynomialPlot[330].Y = 75.065

	pointsOfPolynomialPlot[331].X = 3.31
	pointsOfPolynomialPlot[331].Y = 76.125

	pointsOfPolynomialPlot[332].X = 3.32
	pointsOfPolynomialPlot[332].Y = 77.196

	pointsOfPolynomialPlot[333].X = 3.33
	pointsOfPolynomialPlot[333].Y = 78.278

	pointsOfPolynomialPlot[334].X = 3.34
	pointsOfPolynomialPlot[334].Y = 79.372

	pointsOfPolynomialPlot[335].X = 3.35
	pointsOfPolynomialPlot[335].Y = 80.476

	pointsOfPolynomialPlot[336].X = 3.36
	pointsOfPolynomialPlot[336].Y = 81.592

	pointsOfPolynomialPlot[337].X = 3.37
	pointsOfPolynomialPlot[337].Y = 82.719

	pointsOfPolynomialPlot[338].X = 3.38
	pointsOfPolynomialPlot[338].Y = 83.858

	pointsOfPolynomialPlot[339].X = 3.39
	pointsOfPolynomialPlot[339].Y = 85.008

	pointsOfPolynomialPlot[340].X = 3.4
	pointsOfPolynomialPlot[340].Y = 86.169

	pointsOfPolynomialPlot[341].X = 3.41
	pointsOfPolynomialPlot[341].Y = 87.342

	pointsOfPolynomialPlot[342].X = 3.42
	pointsOfPolynomialPlot[342].Y = 88.527

	pointsOfPolynomialPlot[343].X = 3.43
	pointsOfPolynomialPlot[343].Y = 89.724

	pointsOfPolynomialPlot[344].X = 3.44
	pointsOfPolynomialPlot[344].Y = 90.932

	pointsOfPolynomialPlot[345].X = 3.45
	pointsOfPolynomialPlot[345].Y = 92.153

	pointsOfPolynomialPlot[346].X = 3.46
	pointsOfPolynomialPlot[346].Y = 93.385

	pointsOfPolynomialPlot[347].X = 3.47
	pointsOfPolynomialPlot[347].Y = 94.63

	pointsOfPolynomialPlot[348].X = 3.48
	pointsOfPolynomialPlot[348].Y = 95.887

	pointsOfPolynomialPlot[349].X = 3.49
	pointsOfPolynomialPlot[349].Y = 97.156

	pointsOfPolynomialPlot[350].X = 3.5
	pointsOfPolynomialPlot[350].Y = 98.437

	pointsOfPolynomialPlot[351].X = 3.51
	pointsOfPolynomialPlot[351].Y = 99.731

	pointsOfPolynomialPlot[352].X = 3.52
	pointsOfPolynomialPlot[352].Y = 101.037

	pointsOfPolynomialPlot[353].X = 3.53
	pointsOfPolynomialPlot[353].Y = 102.356

	pointsOfPolynomialPlot[354].X = 3.54
	pointsOfPolynomialPlot[354].Y = 103.687

	pointsOfPolynomialPlot[355].X = 3.55
	pointsOfPolynomialPlot[355].Y = 105.031

	pointsOfPolynomialPlot[356].X = 3.56
	pointsOfPolynomialPlot[356].Y = 106.388

	pointsOfPolynomialPlot[357].X = 3.57
	pointsOfPolynomialPlot[357].Y = 107.758

	pointsOfPolynomialPlot[358].X = 3.58
	pointsOfPolynomialPlot[358].Y = 109.141

	pointsOfPolynomialPlot[359].X = 3.59
	pointsOfPolynomialPlot[359].Y = 110.536

	pointsOfPolynomialPlot[360].X = 3.6
	pointsOfPolynomialPlot[360].Y = 111.945

	pointsOfPolynomialPlot[361].X = 3.61
	pointsOfPolynomialPlot[361].Y = 113.367

	pointsOfPolynomialPlot[362].X = 3.62
	pointsOfPolynomialPlot[362].Y = 114.803

	pointsOfPolynomialPlot[363].X = 3.63
	pointsOfPolynomialPlot[363].Y = 116.251

	pointsOfPolynomialPlot[364].X = 3.64
	pointsOfPolynomialPlot[364].Y = 117.713

	pointsOfPolynomialPlot[365].X = 3.65
	pointsOfPolynomialPlot[365].Y = 119.189

	pointsOfPolynomialPlot[366].X = 3.66
	pointsOfPolynomialPlot[366].Y = 120.678

	pointsOfPolynomialPlot[367].X = 3.67
	pointsOfPolynomialPlot[367].Y = 122.181

	pointsOfPolynomialPlot[368].X = 3.68
	pointsOfPolynomialPlot[368].Y = 123.698

	pointsOfPolynomialPlot[369].X = 3.69
	pointsOfPolynomialPlot[369].Y = 125.228

	pointsOfPolynomialPlot[370].X = 3.7
	pointsOfPolynomialPlot[370].Y = 126.773

	pointsOfPolynomialPlot[371].X = 3.71
	pointsOfPolynomialPlot[371].Y = 128.331

	pointsOfPolynomialPlot[372].X = 3.72
	pointsOfPolynomialPlot[372].Y = 129.904

	pointsOfPolynomialPlot[373].X = 3.73
	pointsOfPolynomialPlot[373].Y = 131.49

	pointsOfPolynomialPlot[374].X = 3.74
	pointsOfPolynomialPlot[374].Y = 133.091

	pointsOfPolynomialPlot[375].X = 3.75
	pointsOfPolynomialPlot[375].Y = 134.707

	pointsOfPolynomialPlot[376].X = 3.76
	pointsOfPolynomialPlot[376].Y = 136.336

	pointsOfPolynomialPlot[377].X = 3.77
	pointsOfPolynomialPlot[377].Y = 137.981

	pointsOfPolynomialPlot[378].X = 3.78
	pointsOfPolynomialPlot[378].Y = 139.639

	pointsOfPolynomialPlot[379].X = 3.79
	pointsOfPolynomialPlot[379].Y = 141.313

	pointsOfPolynomialPlot[380].X = 3.8
	pointsOfPolynomialPlot[380].Y = 143.001

	pointsOfPolynomialPlot[381].X = 3.81
	pointsOfPolynomialPlot[381].Y = 144.704

	pointsOfPolynomialPlot[382].X = 3.82
	pointsOfPolynomialPlot[382].Y = 146.422

	pointsOfPolynomialPlot[383].X = 3.83
	pointsOfPolynomialPlot[383].Y = 148.155

	pointsOfPolynomialPlot[384].X = 3.84
	pointsOfPolynomialPlot[384].Y = 149.904

	pointsOfPolynomialPlot[385].X = 3.85
	pointsOfPolynomialPlot[385].Y = 151.667

	pointsOfPolynomialPlot[386].X = 3.86
	pointsOfPolynomialPlot[386].Y = 153.446

	pointsOfPolynomialPlot[387].X = 3.87
	pointsOfPolynomialPlot[387].Y = 155.24

	pointsOfPolynomialPlot[388].X = 3.88
	pointsOfPolynomialPlot[388].Y = 157.049

	pointsOfPolynomialPlot[389].X = 3.89
	pointsOfPolynomialPlot[389].Y = 158.874

	pointsOfPolynomialPlot[390].X = 3.9
	pointsOfPolynomialPlot[390].Y = 160.715

	pointsOfPolynomialPlot[391].X = 3.91
	pointsOfPolynomialPlot[391].Y = 162.571

	pointsOfPolynomialPlot[392].X = 3.92
	pointsOfPolynomialPlot[392].Y = 164.443

	pointsOfPolynomialPlot[393].X = 3.93
	pointsOfPolynomialPlot[393].Y = 166.331

	pointsOfPolynomialPlot[394].X = 3.94
	pointsOfPolynomialPlot[394].Y = 168.235

	pointsOfPolynomialPlot[395].X = 3.95
	pointsOfPolynomialPlot[395].Y = 170.155

	pointsOfPolynomialPlot[396].X = 3.96
	pointsOfPolynomialPlot[396].Y = 172.091

	pointsOfPolynomialPlot[397].X = 3.97
	pointsOfPolynomialPlot[397].Y = 174.044

	pointsOfPolynomialPlot[398].X = 3.98
	pointsOfPolynomialPlot[398].Y = 176.013

	pointsOfPolynomialPlot[399].X = 3.99
	pointsOfPolynomialPlot[399].Y = 177.998

	pointsOfPolynomialPlot[400].X = 4.0
	pointsOfPolynomialPlot[400].Y = 180.0

	pointsOfPolynomialPlot[401].X = 4.01
	pointsOfPolynomialPlot[401].Y = 182.018

	pointsOfPolynomialPlot[402].X = 4.02
	pointsOfPolynomialPlot[402].Y = 184.053

	pointsOfPolynomialPlot[403].X = 4.03
	pointsOfPolynomialPlot[403].Y = 186.105

	pointsOfPolynomialPlot[404].X = 4.04
	pointsOfPolynomialPlot[404].Y = 188.173

	pointsOfPolynomialPlot[405].X = 4.05
	pointsOfPolynomialPlot[405].Y = 190.259

	pointsOfPolynomialPlot[406].X = 4.06
	pointsOfPolynomialPlot[406].Y = 192.362

	pointsOfPolynomialPlot[407].X = 4.07
	pointsOfPolynomialPlot[407].Y = 194.481

	pointsOfPolynomialPlot[408].X = 4.08
	pointsOfPolynomialPlot[408].Y = 196.618

	pointsOfPolynomialPlot[409].X = 4.09
	pointsOfPolynomialPlot[409].Y = 198.773

	pointsOfPolynomialPlot[410].X = 4.1
	pointsOfPolynomialPlot[410].Y = 200.945

	pointsOfPolynomialPlot[411].X = 4.11
	pointsOfPolynomialPlot[411].Y = 203.134

	pointsOfPolynomialPlot[412].X = 4.12
	pointsOfPolynomialPlot[412].Y = 205.341

	pointsOfPolynomialPlot[413].X = 4.13
	pointsOfPolynomialPlot[413].Y = 207.565

	pointsOfPolynomialPlot[414].X = 4.14
	pointsOfPolynomialPlot[414].Y = 209.808

	pointsOfPolynomialPlot[415].X = 4.15
	pointsOfPolynomialPlot[415].Y = 212.068

	pointsOfPolynomialPlot[416].X = 4.16
	pointsOfPolynomialPlot[416].Y = 214.346

	pointsOfPolynomialPlot[417].X = 4.17
	pointsOfPolynomialPlot[417].Y = 216.643

	pointsOfPolynomialPlot[418].X = 4.18
	pointsOfPolynomialPlot[418].Y = 218.957

	pointsOfPolynomialPlot[419].X = 4.19
	pointsOfPolynomialPlot[419].Y = 221.29

	pointsOfPolynomialPlot[420].X = 4.2
	pointsOfPolynomialPlot[420].Y = 223.641

	pointsOfPolynomialPlot[421].X = 4.21
	pointsOfPolynomialPlot[421].Y = 226.011

	pointsOfPolynomialPlot[422].X = 4.22
	pointsOfPolynomialPlot[422].Y = 228.399

	pointsOfPolynomialPlot[423].X = 4.23
	pointsOfPolynomialPlot[423].Y = 230.806

	pointsOfPolynomialPlot[424].X = 4.24
	pointsOfPolynomialPlot[424].Y = 233.231

	pointsOfPolynomialPlot[425].X = 4.25
	pointsOfPolynomialPlot[425].Y = 235.675

	pointsOfPolynomialPlot[426].X = 4.26
	pointsOfPolynomialPlot[426].Y = 238.675

	pointsOfPolynomialPlot[427].X = 4.27
	pointsOfPolynomialPlot[427].Y = 240.621

	pointsOfPolynomialPlot[428].X = 4.28
	pointsOfPolynomialPlot[428].Y = 243.122

	pointsOfPolynomialPlot[429].X = 4.29
	pointsOfPolynomialPlot[429].Y = 245.643

	pointsOfPolynomialPlot[430].X = 4.3
	pointsOfPolynomialPlot[430].Y = 248.183

	pointsOfPolynomialPlot[431].X = 4.31
	pointsOfPolynomialPlot[431].Y = 250.742

	pointsOfPolynomialPlot[432].X = 4.32
	pointsOfPolynomialPlot[432].Y = 253.321

	pointsOfPolynomialPlot[433].X = 4.33
	pointsOfPolynomialPlot[433].Y = 255.919

	pointsOfPolynomialPlot[434].X = 4.34
	pointsOfPolynomialPlot[434].Y = 258.537

	pointsOfPolynomialPlot[435].X = 4.35
	pointsOfPolynomialPlot[435].Y = 261.175

	pointsOfPolynomialPlot[436].X = 4.36
	pointsOfPolynomialPlot[436].Y = 263.833

	pointsOfPolynomialPlot[437].X = 4.37
	pointsOfPolynomialPlot[437].Y = 266.511

	pointsOfPolynomialPlot[438].X = 4.38
	pointsOfPolynomialPlot[438].Y = 269.209

	pointsOfPolynomialPlot[439].X = 4.39
	pointsOfPolynomialPlot[439].Y = 271.927

	pointsOfPolynomialPlot[440].X = 4.4
	pointsOfPolynomialPlot[440].Y = 274.665

	pointsOfPolynomialPlot[441].X = 4.41
	pointsOfPolynomialPlot[441].Y = 277.424

	pointsOfPolynomialPlot[442].X = 4.42
	pointsOfPolynomialPlot[442].Y = 280.203

	pointsOfPolynomialPlot[443].X = 4.43
	pointsOfPolynomialPlot[443].Y = 283.003

	pointsOfPolynomialPlot[444].X = 4.44
	pointsOfPolynomialPlot[444].Y = 285.824

	pointsOfPolynomialPlot[445].X = 4.45
	pointsOfPolynomialPlot[445].Y = 288.665

	pointsOfPolynomialPlot[446].X = 4.46
	pointsOfPolynomialPlot[446].Y = 291.527

	pointsOfPolynomialPlot[447].X = 4.47
	pointsOfPolynomialPlot[447].Y = 294.41

	pointsOfPolynomialPlot[448].X = 4.48
	pointsOfPolynomialPlot[448].Y = 297.315

	pointsOfPolynomialPlot[449].X = 4.49
	pointsOfPolynomialPlot[449].Y = 300.24

	pointsOfPolynomialPlot[450].X = 4.5
	pointsOfPolynomialPlot[450].Y = 303.187

	pointsOfPolynomialPlot[451].X = 4.51
	pointsOfPolynomialPlot[451].Y = 306.155

	pointsOfPolynomialPlot[452].X = 4.52
	pointsOfPolynomialPlot[452].Y = 309.145

	pointsOfPolynomialPlot[453].X = 4.53
	pointsOfPolynomialPlot[453].Y = 312.156

	pointsOfPolynomialPlot[454].X = 4.54
	pointsOfPolynomialPlot[454].Y = 315.189

	pointsOfPolynomialPlot[455].X = 4.55
	pointsOfPolynomialPlot[455].Y = 318.244

	pointsOfPolynomialPlot[456].X = 4.56
	pointsOfPolynomialPlot[456].Y = 321.321

	pointsOfPolynomialPlot[457].X = 4.57
	pointsOfPolynomialPlot[457].Y = 324.42

	pointsOfPolynomialPlot[458].X = 4.58
	pointsOfPolynomialPlot[458].Y = 327.541

	pointsOfPolynomialPlot[459].X = 4.59
	pointsOfPolynomialPlot[459].Y = 330.684

	pointsOfPolynomialPlot[460].X = 4.6
	pointsOfPolynomialPlot[460].Y = 333.849

	pointsOfPolynomialPlot[461].X = 4.61
	pointsOfPolynomialPlot[461].Y = 337.037

	pointsOfPolynomialPlot[462].X = 4.62
	pointsOfPolynomialPlot[462].Y = 340.247

	pointsOfPolynomialPlot[463].X = 4.63
	pointsOfPolynomialPlot[463].Y = 343.48

	pointsOfPolynomialPlot[464].X = 4.64
	pointsOfPolynomialPlot[464].Y = 346.736

	pointsOfPolynomialPlot[465].X = 4.65
	pointsOfPolynomialPlot[465].Y = 350.015

	pointsOfPolynomialPlot[466].X = 4.66
	pointsOfPolynomialPlot[466].Y = 353.317

	pointsOfPolynomialPlot[467].X = 4.67
	pointsOfPolynomialPlot[467].Y = 356.641

	pointsOfPolynomialPlot[468].X = 4.68
	pointsOfPolynomialPlot[468].Y = 359.989

	pointsOfPolynomialPlot[469].X = 4.69
	pointsOfPolynomialPlot[469].Y = 363.36

	pointsOfPolynomialPlot[470].X = 4.7
	pointsOfPolynomialPlot[470].Y = 366.755

	pointsOfPolynomialPlot[471].X = 4.71
	pointsOfPolynomialPlot[471].Y = 370.173

	pointsOfPolynomialPlot[472].X = 4.72
	pointsOfPolynomialPlot[472].Y = 373.614

	pointsOfPolynomialPlot[473].X = 4.73
	pointsOfPolynomialPlot[473].Y = 377.079

	pointsOfPolynomialPlot[474].X = 4.74
	pointsOfPolynomialPlot[474].Y = 380.569

	pointsOfPolynomialPlot[475].X = 4.75
	pointsOfPolynomialPlot[475].Y = 384.082

	pointsOfPolynomialPlot[476].X = 4.76
	pointsOfPolynomialPlot[476].Y = 387.619

	pointsOfPolynomialPlot[477].X = 4.77
	pointsOfPolynomialPlot[477].Y = 391.18

	pointsOfPolynomialPlot[478].X = 4.78
	pointsOfPolynomialPlot[478].Y = 394.765

	pointsOfPolynomialPlot[479].X = 4.79
	pointsOfPolynomialPlot[479].Y = 398.375

	pointsOfPolynomialPlot[480].X = 4.8
	pointsOfPolynomialPlot[480].Y = 402.009

	pointsOfPolynomialPlot[481].X = 4.81
	pointsOfPolynomialPlot[481].Y = 405.668

	pointsOfPolynomialPlot[482].X = 4.82
	pointsOfPolynomialPlot[482].Y = 409.351

	pointsOfPolynomialPlot[483].X = 4.83
	pointsOfPolynomialPlot[483].Y = 413.06

	pointsOfPolynomialPlot[484].X = 4.84
	pointsOfPolynomialPlot[484].Y = 416.793

	pointsOfPolynomialPlot[485].X = 4.85
	pointsOfPolynomialPlot[485].Y = 420.551

	pointsOfPolynomialPlot[486].X = 4.86
	pointsOfPolynomialPlot[486].Y = 424.334

	pointsOfPolynomialPlot[487].X = 4.87
	pointsOfPolynomialPlot[487].Y = 428.143

	pointsOfPolynomialPlot[488].X = 4.88
	pointsOfPolynomialPlot[488].Y = 431.977

	pointsOfPolynomialPlot[489].X = 4.89
	pointsOfPolynomialPlot[489].Y = 435.836

	pointsOfPolynomialPlot[490].X = 4.9
	pointsOfPolynomialPlot[490].Y = 439.721

	pointsOfPolynomialPlot[491].X = 4.91
	pointsOfPolynomialPlot[491].Y = 443.631

	pointsOfPolynomialPlot[492].X = 4.92
	pointsOfPolynomialPlot[492].Y = 447.567

	pointsOfPolynomialPlot[493].X = 4.93
	pointsOfPolynomialPlot[493].Y = 451.53

	pointsOfPolynomialPlot[494].X = 4.94
	pointsOfPolynomialPlot[494].Y = 455.518

	pointsOfPolynomialPlot[495].X = 4.95
	pointsOfPolynomialPlot[495].Y = 459.532

	pointsOfPolynomialPlot[496].X = 4.96
	pointsOfPolynomialPlot[496].Y = 463.573

	pointsOfPolynomialPlot[497].X = 4.97
	pointsOfPolynomialPlot[497].Y = 467.64

	pointsOfPolynomialPlot[498].X = 4.98
	pointsOfPolynomialPlot[498].Y = 471.733

	pointsOfPolynomialPlot[499].X = 4.99
	pointsOfPolynomialPlot[499].Y = 475.853

	pointsOfPolynomialPlot[500].X = 5.0
	pointsOfPolynomialPlot[500].Y = 480.0

	pointsOfPolynomialPlot[501].X = 5.01
	pointsOfPolynomialPlot[501].Y = 484.173

	pointsOfPolynomialPlot[502].X = 5.02
	pointsOfPolynomialPlot[502].Y = 488.373

	pointsOfPolynomialPlot[503].X = 5.03
	pointsOfPolynomialPlot[503].Y = 492.601

	pointsOfPolynomialPlot[504].X = 5.04
	pointsOfPolynomialPlot[504].Y = 496.855

	pointsOfPolynomialPlot[505].X = 5.05
	pointsOfPolynomialPlot[505].Y = 501.137

	pointsOfPolynomialPlot[506].X = 5.06
	pointsOfPolynomialPlot[506].Y = 505.446

	pointsOfPolynomialPlot[507].X = 5.07
	pointsOfPolynomialPlot[507].Y = 509.783

	pointsOfPolynomialPlot[508].X = 5.08
	pointsOfPolynomialPlot[508].Y = 514.147

	pointsOfPolynomialPlot[509].X = 5.09
	pointsOfPolynomialPlot[509].Y = 518.539

	pointsOfPolynomialPlot[510].X = 5.1
	pointsOfPolynomialPlot[510].Y = 522.959

	pointsOfPolynomialPlot[511].X = 5.11
	pointsOfPolynomialPlot[511].Y = 527.406

	pointsOfPolynomialPlot[512].X = 5.12
	pointsOfPolynomialPlot[512].Y = 531.882

	pointsOfPolynomialPlot[513].X = 5.13
	pointsOfPolynomialPlot[513].Y = 536.386

	pointsOfPolynomialPlot[514].X = 5.14
	pointsOfPolynomialPlot[514].Y = 540.918

	pointsOfPolynomialPlot[515].X = 5.15
	pointsOfPolynomialPlot[515].Y = 545.479

	pointsOfPolynomialPlot[516].X = 5.16
	pointsOfPolynomialPlot[516].Y = 550.068

	pointsOfPolynomialPlot[517].X = 5.17
	pointsOfPolynomialPlot[517].Y = 554.686

	pointsOfPolynomialPlot[518].X = 5.18
	pointsOfPolynomialPlot[518].Y = 559.333

	pointsOfPolynomialPlot[519].X = 5.19
	pointsOfPolynomialPlot[519].Y = 564.009

	pointsOfPolynomialPlot[520].X = 5.2
	pointsOfPolynomialPlot[520].Y = 568.713

	pointsOfPolynomialPlot[521].X = 5.21
	pointsOfPolynomialPlot[521].Y = 573.447

	pointsOfPolynomialPlot[522].X = 5.22
	pointsOfPolynomialPlot[522].Y = 578.21

	pointsOfPolynomialPlot[523].X = 5.23
	pointsOfPolynomialPlot[523].Y = 583.002

	pointsOfPolynomialPlot[524].X = 5.24
	pointsOfPolynomialPlot[524].Y = 587.824

	pointsOfPolynomialPlot[525].X = 5.25
	pointsOfPolynomialPlot[525].Y = 592.675

	pointsOfPolynomialPlot[526].X = 5.26
	pointsOfPolynomialPlot[526].Y = 597.556

	pointsOfPolynomialPlot[527].X = 5.27
	pointsOfPolynomialPlot[527].Y = 602.467

	pointsOfPolynomialPlot[528].X = 5.28
	pointsOfPolynomialPlot[528].Y = 607.408

	pointsOfPolynomialPlot[529].X = 5.29
	pointsOfPolynomialPlot[529].Y = 612.379

	pointsOfPolynomialPlot[530].X = 5.3
	pointsOfPolynomialPlot[530].Y = 617.381

	pointsOfPolynomialPlot[531].X = 5.31
	pointsOfPolynomialPlot[531].Y = 622.412

	pointsOfPolynomialPlot[532].X = 5.32
	pointsOfPolynomialPlot[532].Y = 627.474

	pointsOfPolynomialPlot[533].X = 5.33
	pointsOfPolynomialPlot[533].Y = 632.567

	pointsOfPolynomialPlot[534].X = 5.34
	pointsOfPolynomialPlot[534].Y = 637.69

	pointsOfPolynomialPlot[535].X = 5.35
	pointsOfPolynomialPlot[535].Y = 642.844

	pointsOfPolynomialPlot[536].X = 5.36
	pointsOfPolynomialPlot[536].Y = 648.029

	pointsOfPolynomialPlot[537].X = 5.37
	pointsOfPolynomialPlot[537].Y = 653.245

	pointsOfPolynomialPlot[538].X = 5.38
	pointsOfPolynomialPlot[538].Y = 658.493

	pointsOfPolynomialPlot[539].X = 5.39
	pointsOfPolynomialPlot[539].Y = 663.771

	pointsOfPolynomialPlot[540].X = 5.4
	pointsOfPolynomialPlot[540].Y = 669.081

	pointsOfPolynomialPlot[541].X = 5.41
	pointsOfPolynomialPlot[541].Y = 674.423

	pointsOfPolynomialPlot[542].X = 5.42
	pointsOfPolynomialPlot[542].Y = 679.796

	pointsOfPolynomialPlot[543].X = 5.43
	pointsOfPolynomialPlot[543].Y = 685.201

	pointsOfPolynomialPlot[544].X = 5.44
	pointsOfPolynomialPlot[544].Y = 690.638

	pointsOfPolynomialPlot[545].X = 5.45
	pointsOfPolynomialPlot[545].Y = 696.107

	pointsOfPolynomialPlot[546].X = 5.46
	pointsOfPolynomialPlot[546].Y = 701.608

	pointsOfPolynomialPlot[547].X = 5.47
	pointsOfPolynomialPlot[547].Y = 707.142

	pointsOfPolynomialPlot[548].X = 5.48
	pointsOfPolynomialPlot[548].Y = 712.707

	pointsOfPolynomialPlot[549].X = 5.49
	pointsOfPolynomialPlot[549].Y = 718.306

	pointsOfPolynomialPlot[550].X = 5.5
	pointsOfPolynomialPlot[550].Y = 723.937

	pointsOfPolynomialPlot[551].X = 5.51
	pointsOfPolynomialPlot[551].Y = 729.601

	pointsOfPolynomialPlot[552].X = 5.52
	pointsOfPolynomialPlot[552].Y = 735.298

	pointsOfPolynomialPlot[553].X = 5.53
	pointsOfPolynomialPlot[553].Y = 741.028

	pointsOfPolynomialPlot[554].X = 5.54
	pointsOfPolynomialPlot[554].Y = 746.791

	pointsOfPolynomialPlot[555].X = 5.55
	pointsOfPolynomialPlot[555].Y = 752.587

	pointsOfPolynomialPlot[556].X = 5.56
	pointsOfPolynomialPlot[556].Y = 758.417

	pointsOfPolynomialPlot[557].X = 5.57
	pointsOfPolynomialPlot[557].Y = 764.28

	pointsOfPolynomialPlot[558].X = 5.58
	pointsOfPolynomialPlot[558].Y = 770.177

	pointsOfPolynomialPlot[559].X = 5.59
	pointsOfPolynomialPlot[559].Y = 776.108

	pointsOfPolynomialPlot[560].X = 5.6
	pointsOfPolynomialPlot[560].Y = 782.073

	pointsOfPolynomialPlot[561].X = 5.61
	pointsOfPolynomialPlot[561].Y = 788.072

	pointsOfPolynomialPlot[562].X = 5.62
	pointsOfPolynomialPlot[562].Y = 794.105

	pointsOfPolynomialPlot[563].X = 5.63
	pointsOfPolynomialPlot[563].Y = 800.173

	pointsOfPolynomialPlot[564].X = 5.64
	pointsOfPolynomialPlot[564].Y = 806.274

	pointsOfPolynomialPlot[565].X = 5.65
	pointsOfPolynomialPlot[565].Y = 812.411

	pointsOfPolynomialPlot[566].X = 5.66
	pointsOfPolynomialPlot[566].Y = 818.582

	pointsOfPolynomialPlot[567].X = 5.67
	pointsOfPolynomialPlot[567].Y = 824.788

	pointsOfPolynomialPlot[568].X = 5.68
	pointsOfPolynomialPlot[568].Y = 831.029

	pointsOfPolynomialPlot[569].X = 5.69
	pointsOfPolynomialPlot[569].Y = 837.305

	pointsOfPolynomialPlot[570].X = 5.7
	pointsOfPolynomialPlot[570].Y = 843.617

	pointsOfPolynomialPlot[571].X = 5.71
	pointsOfPolynomialPlot[571].Y = 849.963

	pointsOfPolynomialPlot[572].X = 5.72
	pointsOfPolynomialPlot[572].Y = 856.346

	pointsOfPolynomialPlot[573].X = 5.73
	pointsOfPolynomialPlot[573].Y = 862.763

	pointsOfPolynomialPlot[574].X = 5.74
	pointsOfPolynomialPlot[574].Y = 869.217

	pointsOfPolynomialPlot[575].X = 5.75
	pointsOfPolynomialPlot[575].Y = 875.707

	pointsOfPolynomialPlot[576].X = 5.76
	pointsOfPolynomialPlot[576].Y = 882.232

	pointsOfPolynomialPlot[577].X = 5.77
	pointsOfPolynomialPlot[577].Y = 888.794

	pointsOfPolynomialPlot[578].X = 5.78
	pointsOfPolynomialPlot[578].Y = 895.392

	pointsOfPolynomialPlot[579].X = 5.79
	pointsOfPolynomialPlot[579].Y = 902.026

	pointsOfPolynomialPlot[580].X = 5.8
	pointsOfPolynomialPlot[580].Y = 908.697

	pointsOfPolynomialPlot[581].X = 5.81
	pointsOfPolynomialPlot[581].Y = 915.405

	pointsOfPolynomialPlot[582].X = 5.82
	pointsOfPolynomialPlot[582].Y = 922.149

	pointsOfPolynomialPlot[583].X = 5.83
	pointsOfPolynomialPlot[583].Y = 928.931

	pointsOfPolynomialPlot[584].X = 5.84
	pointsOfPolynomialPlot[584].Y = 935.749

	pointsOfPolynomialPlot[585].X = 5.85
	pointsOfPolynomialPlot[585].Y = 942.605

	pointsOfPolynomialPlot[586].X = 5.86
	pointsOfPolynomialPlot[586].Y = 949.498

	pointsOfPolynomialPlot[587].X = 5.87
	pointsOfPolynomialPlot[587].Y = 956.429

	pointsOfPolynomialPlot[588].X = 5.88
	pointsOfPolynomialPlot[588].Y = 963.397

	pointsOfPolynomialPlot[589].X = 5.89
	pointsOfPolynomialPlot[589].Y = 970.403

	pointsOfPolynomialPlot[590].X = 5.9
	pointsOfPolynomialPlot[590].Y = 977.447

	pointsOfPolynomialPlot[591].X = 5.91
	pointsOfPolynomialPlot[591].Y = 984.529

	pointsOfPolynomialPlot[592].X = 5.92
	pointsOfPolynomialPlot[592].Y = 991.649

	pointsOfPolynomialPlot[593].X = 5.93
	pointsOfPolynomialPlot[593].Y = 998.807

	pointsOfPolynomialPlot[594].X = 5.94
	pointsOfPolynomialPlot[594].Y = 1_006.004

	pointsOfPolynomialPlot[595].X = 5.95
	pointsOfPolynomialPlot[595].Y = 1_013.239

	pointsOfPolynomialPlot[596].X = 5.96
	pointsOfPolynomialPlot[596].Y = 1_020.513

	pointsOfPolynomialPlot[597].X = 5.97
	pointsOfPolynomialPlot[597].Y = 1_027.826

	pointsOfPolynomialPlot[598].X = 5.98
	pointsOfPolynomialPlot[598].Y = 1_035.178

	pointsOfPolynomialPlot[599].X = 5.99
	pointsOfPolynomialPlot[599].Y = 1_042.569

	pointsOfPolynomialPlot[600].X = 6.0
	pointsOfPolynomialPlot[600].Y = 1_050.0

	pointsOfPolynomialPlot[601].X = 6.01
	pointsOfPolynomialPlot[601].Y = 1_057.469

	pointsOfPolynomialPlot[602].X = 6.02
	pointsOfPolynomialPlot[602].Y = 1_064.979

	pointsOfPolynomialPlot[603].X = 6.03
	pointsOfPolynomialPlot[603].Y = 1_072.527

	pointsOfPolynomialPlot[604].X = 6.04
	pointsOfPolynomialPlot[604].Y = 1_080.116

	pointsOfPolynomialPlot[605].X = 6.05
	pointsOfPolynomialPlot[605].Y = 1_087.745

	pointsOfPolynomialPlot[606].X = 6.06
	pointsOfPolynomialPlot[606].Y = 1_095.414

	pointsOfPolynomialPlot[607].X = 6.07
	pointsOfPolynomialPlot[607].Y = 1_103.123

	pointsOfPolynomialPlot[608].X = 6.08
	pointsOfPolynomialPlot[608].Y = 1_110.872

	pointsOfPolynomialPlot[609].X = 6.09
	pointsOfPolynomialPlot[609].Y = 1_118.662

	pointsOfPolynomialPlot[610].X = 6.1
	pointsOfPolynomialPlot[610].Y = 1_126.493

	pointsOfPolynomialPlot[611].X = 6.11
	pointsOfPolynomialPlot[611].Y = 1_134.364

	pointsOfPolynomialPlot[612].X = 6.12
	pointsOfPolynomialPlot[612].Y = 1_142.276

	pointsOfPolynomialPlot[613].X = 6.13
	pointsOfPolynomialPlot[613].Y = 1_150.23

	pointsOfPolynomialPlot[614].X = 6.14
	pointsOfPolynomialPlot[614].Y = 1_158.224

	pointsOfPolynomialPlot[615].X = 6.15
	pointsOfPolynomialPlot[615].Y = 1_166.26

	pointsOfPolynomialPlot[616].X = 6.16
	pointsOfPolynomialPlot[616].Y = 1_174.338

	pointsOfPolynomialPlot[617].X = 6.17
	pointsOfPolynomialPlot[617].Y = 1_182.457

	pointsOfPolynomialPlot[618].X = 6.18
	pointsOfPolynomialPlot[618].Y = 1_190.618

	pointsOfPolynomialPlot[619].X = 6.19
	pointsOfPolynomialPlot[619].Y = 1_198.82

	pointsOfPolynomialPlot[620].X = 6.2
	pointsOfPolynomialPlot[620].Y = 1_207.065

	pointsOfPolynomialPlot[621].X = 6.21
	pointsOfPolynomialPlot[621].Y = 1_215.352

	pointsOfPolynomialPlot[622].X = 6.22
	pointsOfPolynomialPlot[622].Y = 1_223.682

	pointsOfPolynomialPlot[623].X = 6.23
	pointsOfPolynomialPlot[623].Y = 1_232.053

	pointsOfPolynomialPlot[624].X = 6.24
	pointsOfPolynomialPlot[624].Y = 1_240.468

	pointsOfPolynomialPlot[625].X = 6.25
	pointsOfPolynomialPlot[625].Y = 1_248.925

	pointsOfPolynomialPlot[626].X = 6.26
	pointsOfPolynomialPlot[626].Y = 1_257.426

	pointsOfPolynomialPlot[627].X = 6.27
	pointsOfPolynomialPlot[627].Y = 1_265.969

	pointsOfPolynomialPlot[628].X = 6.28
	pointsOfPolynomialPlot[628].Y = 1_274.555

	pointsOfPolynomialPlot[629].X = 6.29
	pointsOfPolynomialPlot[629].Y = 1_283.185

	pointsOfPolynomialPlot[630].X = 6.3
	pointsOfPolynomialPlot[630].Y = 1_291.859

	pointsOfPolynomialPlot[631].X = 6.31
	pointsOfPolynomialPlot[631].Y = 1_300.576

	pointsOfPolynomialPlot[632].X = 6.32
	pointsOfPolynomialPlot[632].Y = 1_309.336

	pointsOfPolynomialPlot[633].X = 6.33
	pointsOfPolynomialPlot[633].Y = 1_318.141

	pointsOfPolynomialPlot[634].X = 6.34
	pointsOfPolynomialPlot[634].Y = 1_326.99

	pointsOfPolynomialPlot[635].X = 6.35
	pointsOfPolynomialPlot[635].Y = 1_335.883

	pointsOfPolynomialPlot[636].X = 6.36
	pointsOfPolynomialPlot[636].Y = 1_344.821

	pointsOfPolynomialPlot[637].X = 6.37
	pointsOfPolynomialPlot[637].Y = 1_353.803

	pointsOfPolynomialPlot[638].X = 6.38
	pointsOfPolynomialPlot[638].Y = 1_362.829

	pointsOfPolynomialPlot[639].X = 6.39
	pointsOfPolynomialPlot[639].Y = 1_371.901

	pointsOfPolynomialPlot[640].X = 6.4
	pointsOfPolynomialPlot[640].Y = 1_381.017

	pointsOfPolynomialPlot[641].X = 6.41
	pointsOfPolynomialPlot[641].Y = 1_390.179

	pointsOfPolynomialPlot[642].X = 6.42
	pointsOfPolynomialPlot[642].Y = 1_399.385

	pointsOfPolynomialPlot[643].X = 6.43
	pointsOfPolynomialPlot[643].Y = 1_408.638

	pointsOfPolynomialPlot[644].X = 6.44
	pointsOfPolynomialPlot[644].Y = 1_417.935

	pointsOfPolynomialPlot[645].X = 6.45
	pointsOfPolynomialPlot[645].Y = 1_427.279

	pointsOfPolynomialPlot[646].X = 6.46
	pointsOfPolynomialPlot[646].Y = 1_436.668

	pointsOfPolynomialPlot[647].X = 6.47
	pointsOfPolynomialPlot[647].Y = 1_446.104

	pointsOfPolynomialPlot[648].X = 6.48
	pointsOfPolynomialPlot[648].Y = 1_455.585

	pointsOfPolynomialPlot[649].X = 6.49
	pointsOfPolynomialPlot[649].Y = 1_465.113

	pointsOfPolynomialPlot[650].X = 6.5
	pointsOfPolynomialPlot[650].Y = 1_474.687

	pointsOfPolynomialPlot[651].X = 6.51
	pointsOfPolynomialPlot[651].Y = 1_484.308

	pointsOfPolynomialPlot[652].X = 6.52
	pointsOfPolynomialPlot[652].Y = 1_493.975

	pointsOfPolynomialPlot[653].X = 6.53
	pointsOfPolynomialPlot[653].Y = 1_503.69

	pointsOfPolynomialPlot[654].X = 6.54
	pointsOfPolynomialPlot[654].Y = 1_513.451

	pointsOfPolynomialPlot[655].X = 6.55
	pointsOfPolynomialPlot[655].Y = 1_523.26

	pointsOfPolynomialPlot[656].X = 6.56
	pointsOfPolynomialPlot[656].Y = 1_533.116

	pointsOfPolynomialPlot[657].X = 6.57
	pointsOfPolynomialPlot[657].Y = 1_543.02

	pointsOfPolynomialPlot[658].X = 6.58
	pointsOfPolynomialPlot[658].Y = 1_552.971

	pointsOfPolynomialPlot[659].X = 6.59
	pointsOfPolynomialPlot[659].Y = 1_562.97

	pointsOfPolynomialPlot[660].X = 6.6
	pointsOfPolynomialPlot[660].Y = 1_573.017

	pointsOfPolynomialPlot[661].X = 6.61
	pointsOfPolynomialPlot[661].Y = 1_583.112

	pointsOfPolynomialPlot[662].X = 6.62
	pointsOfPolynomialPlot[662].Y = 1_593.256

	pointsOfPolynomialPlot[663].X = 6.63
	pointsOfPolynomialPlot[663].Y = 1_603.447

	pointsOfPolynomialPlot[664].X = 6.64
	pointsOfPolynomialPlot[664].Y = 1_613.688

	pointsOfPolynomialPlot[665].X = 6.65
	pointsOfPolynomialPlot[665].Y = 1_623.977

	pointsOfPolynomialPlot[666].X = 6.66
	pointsOfPolynomialPlot[666].Y = 1_634.315

	pointsOfPolynomialPlot[667].X = 6.67
	pointsOfPolynomialPlot[667].Y = 1_644.702

	pointsOfPolynomialPlot[668].X = 6.68
	pointsOfPolynomialPlot[668].Y = 1_655.138

	pointsOfPolynomialPlot[669].X = 6.69
	pointsOfPolynomialPlot[669].Y = 1_665.138

	pointsOfPolynomialPlot[670].X = 6.7
	pointsOfPolynomialPlot[670].Y = 1_676.159

	pointsOfPolynomialPlot[671].X = 6.71
	pointsOfPolynomialPlot[671].Y = 1_686.743

	pointsOfPolynomialPlot[672].X = 6.72
	pointsOfPolynomialPlot[672].Y = 1_697.378

	pointsOfPolynomialPlot[673].X = 6.73
	pointsOfPolynomialPlot[673].Y = 1_708.062

	pointsOfPolynomialPlot[674].X = 6.74
	pointsOfPolynomialPlot[674].Y = 1_718.797

	pointsOfPolynomialPlot[675].X = 6.75
	pointsOfPolynomialPlot[675].Y = 1_729.582

	pointsOfPolynomialPlot[676].X = 6.76
	pointsOfPolynomialPlot[676].Y = 1_740.417

	pointsOfPolynomialPlot[677].X = 6.77
	pointsOfPolynomialPlot[677].Y = 1_751.303

	pointsOfPolynomialPlot[678].X = 6.78
	pointsOfPolynomialPlot[678].Y = 1_762.239

	pointsOfPolynomialPlot[679].X = 6.79
	pointsOfPolynomialPlot[679].Y = 1_773.227

	pointsOfPolynomialPlot[680].X = 6.8
	pointsOfPolynomialPlot[680].Y = 1_784.265

	pointsOfPolynomialPlot[681].X = 6.81
	pointsOfPolynomialPlot[681].Y = 1_795.355

	pointsOfPolynomialPlot[682].X = 6.82
	pointsOfPolynomialPlot[682].Y = 1_806.496

	pointsOfPolynomialPlot[683].X = 6.83
	pointsOfPolynomialPlot[683].Y = 1_817.689

	pointsOfPolynomialPlot[684].X = 6.84
	pointsOfPolynomialPlot[684].Y = 1_828.933

	pointsOfPolynomialPlot[685].X = 6.85
	pointsOfPolynomialPlot[685].Y = 1_840.229

	pointsOfPolynomialPlot[686].X = 6.86
	pointsOfPolynomialPlot[686].Y = 1_851.577

	pointsOfPolynomialPlot[687].X = 6.87
	pointsOfPolynomialPlot[687].Y = 1_862.977

	pointsOfPolynomialPlot[688].X = 6.88
	pointsOfPolynomialPlot[688].Y = 1_874.43

	pointsOfPolynomialPlot[689].X = 6.89
	pointsOfPolynomialPlot[689].Y = 1_885.935

	pointsOfPolynomialPlot[690].X = 6.9
	pointsOfPolynomialPlot[690].Y = 1_897.493

	pointsOfPolynomialPlot[691].X = 6.91
	pointsOfPolynomialPlot[691].Y = 1_909.103

	pointsOfPolynomialPlot[692].X = 6.92
	pointsOfPolynomialPlot[692].Y = 1_920.767

	pointsOfPolynomialPlot[693].X = 6.93
	pointsOfPolynomialPlot[693].Y = 1_932.483

	pointsOfPolynomialPlot[694].X = 6.94
	pointsOfPolynomialPlot[694].Y = 1_944.253

	pointsOfPolynomialPlot[695].X = 6.95
	pointsOfPolynomialPlot[695].Y = 1_956.076

	pointsOfPolynomialPlot[696].X = 6.96
	pointsOfPolynomialPlot[696].Y = 1_967.953

	pointsOfPolynomialPlot[697].X = 6.97
	pointsOfPolynomialPlot[697].Y = 1_979.884

	pointsOfPolynomialPlot[698].X = 6.98
	pointsOfPolynomialPlot[698].Y = 1_991.868

	pointsOfPolynomialPlot[699].X = 6.99
	pointsOfPolynomialPlot[699].Y = 2_003.907

	pointsOfPolynomialPlot[700].X = 7.0
	pointsOfPolynomialPlot[700].Y = 2_016.0

	pointsOfPolynomialPlot[701].X = 7.01
	pointsOfPolynomialPlot[701].Y = 2_028.147

	pointsOfPolynomialPlot[702].X = 7.02
	pointsOfPolynomialPlot[702].Y = 2_040.349

	pointsOfPolynomialPlot[703].X = 7.03
	pointsOfPolynomialPlot[703].Y = 2_052.605

	pointsOfPolynomialPlot[704].X = 7.04
	pointsOfPolynomialPlot[704].Y = 2_064.916

	pointsOfPolynomialPlot[705].X = 7.05
	pointsOfPolynomialPlot[705].Y = 2_077.283

	pointsOfPolynomialPlot[706].X = 7.06
	pointsOfPolynomialPlot[706].Y = 2_089.705

	pointsOfPolynomialPlot[707].X = 7.07
	pointsOfPolynomialPlot[707].Y = 2_102.182

	pointsOfPolynomialPlot[708].X = 7.08
	pointsOfPolynomialPlot[708].Y = 2_114.714

	pointsOfPolynomialPlot[709].X = 7.09
	pointsOfPolynomialPlot[709].Y = 2_127.302

	pointsOfPolynomialPlot[710].X = 7.1
	pointsOfPolynomialPlot[710].Y = 2_139.947

	pointsOfPolynomialPlot[711].X = 7.11
	pointsOfPolynomialPlot[711].Y = 2_152.647

	pointsOfPolynomialPlot[712].X = 7.12
	pointsOfPolynomialPlot[712].Y = 2_165.403

	pointsOfPolynomialPlot[713].X = 7.13
	pointsOfPolynomialPlot[713].Y = 2_178.216

	pointsOfPolynomialPlot[714].X = 7.14
	pointsOfPolynomialPlot[714].Y = 2_191.085

	pointsOfPolynomialPlot[715].X = 7.15
	pointsOfPolynomialPlot[715].Y = 2_204.011

	pointsOfPolynomialPlot[716].X = 7.16
	pointsOfPolynomialPlot[716].Y = 2_216.994

	pointsOfPolynomialPlot[717].X = 7.17
	pointsOfPolynomialPlot[717].Y = 2_230.034

	pointsOfPolynomialPlot[718].X = 7.18
	pointsOfPolynomialPlot[718].Y = 2_243.131

	pointsOfPolynomialPlot[719].X = 7.19
	pointsOfPolynomialPlot[719].Y = 2_256.285

	pointsOfPolynomialPlot[720].X = 7.2
	pointsOfPolynomialPlot[720].Y = 2_269.497

	pointsOfPolynomialPlot[721].X = 7.21
	pointsOfPolynomialPlot[721].Y = 2_282.767

	pointsOfPolynomialPlot[722].X = 7.22
	pointsOfPolynomialPlot[722].Y = 2_296.094

	pointsOfPolynomialPlot[723].X = 7.23
	pointsOfPolynomialPlot[723].Y = 2_309.48

	pointsOfPolynomialPlot[724].X = 7.24
	pointsOfPolynomialPlot[724].Y = 2_322.923

	pointsOfPolynomialPlot[725].X = 7.25
	pointsOfPolynomialPlot[725].Y = 2_336.425

	pointsOfPolynomialPlot[726].X = 7.26
	pointsOfPolynomialPlot[726].Y = 2_349.986

	pointsOfPolynomialPlot[727].X = 7.27
	pointsOfPolynomialPlot[727].Y = 2_363.605

	pointsOfPolynomialPlot[728].X = 7.28
	pointsOfPolynomialPlot[728].Y = 2_377.283

	pointsOfPolynomialPlot[729].X = 7.29
	pointsOfPolynomialPlot[729].Y = 2_391.02

	pointsOfPolynomialPlot[730].X = 7.3
	pointsOfPolynomialPlot[730].Y = 2_404.817

	pointsOfPolynomialPlot[731].X = 7.31
	pointsOfPolynomialPlot[731].Y = 2_418.672

	pointsOfPolynomialPlot[732].X = 7.32
	pointsOfPolynomialPlot[732].Y = 2_432.588

	pointsOfPolynomialPlot[733].X = 7.33
	pointsOfPolynomialPlot[733].Y = 2_446.563

	pointsOfPolynomialPlot[734].X = 7.34
	pointsOfPolynomialPlot[734].Y = 2_460.597

	pointsOfPolynomialPlot[735].X = 7.35
	pointsOfPolynomialPlot[735].Y = 2_474.692

	pointsOfPolynomialPlot[736].X = 7.36
	pointsOfPolynomialPlot[736].Y = 2_488.847

	pointsOfPolynomialPlot[737].X = 7.37
	pointsOfPolynomialPlot[737].Y = 2_503.063

	pointsOfPolynomialPlot[738].X = 7.38
	pointsOfPolynomialPlot[738].Y = 2_517.339

	pointsOfPolynomialPlot[739].X = 7.39
	pointsOfPolynomialPlot[739].Y = 2_531.675

	pointsOfPolynomialPlot[740].X = 7.4
	pointsOfPolynomialPlot[740].Y = 2_546.073

	pointsOfPolynomialPlot[741].X = 7.41
	pointsOfPolynomialPlot[741].Y = 2_560.532

	pointsOfPolynomialPlot[742].X = 7.42
	pointsOfPolynomialPlot[742].Y = 2_575.052

	pointsOfPolynomialPlot[743].X = 7.43
	pointsOfPolynomialPlot[743].Y = 2_589.633

	pointsOfPolynomialPlot[744].X = 7.44
	pointsOfPolynomialPlot[744].Y = 2_604.276

	pointsOfPolynomialPlot[745].X = 7.45
	pointsOfPolynomialPlot[745].Y = 2_618.981

	pointsOfPolynomialPlot[746].X = 7.46
	pointsOfPolynomialPlot[746].Y = 2_633.748

	pointsOfPolynomialPlot[747].X = 7.47
	pointsOfPolynomialPlot[747].Y = 2_648.576

	pointsOfPolynomialPlot[748].X = 7.48
	pointsOfPolynomialPlot[748].Y = 2_663.467

	pointsOfPolynomialPlot[749].X = 7.49
	pointsOfPolynomialPlot[749].Y = 2_678.421

	pointsOfPolynomialPlot[750].X = 7.5
	pointsOfPolynomialPlot[750].Y = 2_693.437

	pointsOfPolynomialPlot[751].X = 7.51
	pointsOfPolynomialPlot[751].Y = 2_708.516

	pointsOfPolynomialPlot[752].X = 7.52
	pointsOfPolynomialPlot[752].Y = 2_723.658

	pointsOfPolynomialPlot[753].X = 7.53
	pointsOfPolynomialPlot[753].Y = 2_738.863

	pointsOfPolynomialPlot[754].X = 7.54
	pointsOfPolynomialPlot[754].Y = 2_754.131

	pointsOfPolynomialPlot[755].X = 7.55
	pointsOfPolynomialPlot[755].Y = 2_769.463

	pointsOfPolynomialPlot[756].X = 7.56
	pointsOfPolynomialPlot[756].Y = 2_784.859

	pointsOfPolynomialPlot[757].X = 7.57
	pointsOfPolynomialPlot[757].Y = 2_800.318

	pointsOfPolynomialPlot[758].X = 7.58
	pointsOfPolynomialPlot[758].Y = 2_815.842

	pointsOfPolynomialPlot[759].X = 7.59
	pointsOfPolynomialPlot[759].Y = 2_831.429

	pointsOfPolynomialPlot[760].X = 7.6
	pointsOfPolynomialPlot[760].Y = 2_847.081

	pointsOfPolynomialPlot[761].X = 7.61
	pointsOfPolynomialPlot[761].Y = 2_862.798

	pointsOfPolynomialPlot[762].X = 7.62
	pointsOfPolynomialPlot[762].Y = 2_878.579

	pointsOfPolynomialPlot[763].X = 7.63
	pointsOfPolynomialPlot[763].Y = 2_894.425

	pointsOfPolynomialPlot[764].X = 7.64
	pointsOfPolynomialPlot[764].Y = 2_910.336

	pointsOfPolynomialPlot[765].X = 7.65
	pointsOfPolynomialPlot[765].Y = 2_926.313

	pointsOfPolynomialPlot[766].X = 7.66
	pointsOfPolynomialPlot[766].Y = 2_942.355

	pointsOfPolynomialPlot[767].X = 7.67
	pointsOfPolynomialPlot[767].Y = 2_958.462

	pointsOfPolynomialPlot[768].X = 7.68
	pointsOfPolynomialPlot[768].Y = 2_974.636

	pointsOfPolynomialPlot[769].X = 7.69
	pointsOfPolynomialPlot[769].Y = 2_990.875

	pointsOfPolynomialPlot[770].X = 7.7
	pointsOfPolynomialPlot[770].Y = 3_007.181

	pointsOfPolynomialPlot[771].X = 7.71
	pointsOfPolynomialPlot[771].Y = 3_023.552

	pointsOfPolynomialPlot[772].X = 7.72
	pointsOfPolynomialPlot[772].Y = 3_039.991

	pointsOfPolynomialPlot[773].X = 7.73
	pointsOfPolynomialPlot[773].Y = 3_056.496

	pointsOfPolynomialPlot[774].X = 7.74
	pointsOfPolynomialPlot[774].Y = 3_073.068

	pointsOfPolynomialPlot[775].X = 7.75
	pointsOfPolynomialPlot[775].Y = 3_089.707

	pointsOfPolynomialPlot[776].X = 7.76
	pointsOfPolynomialPlot[776].Y = 3_106.413

	pointsOfPolynomialPlot[777].X = 7.77
	pointsOfPolynomialPlot[777].Y = 3_123.186

	pointsOfPolynomialPlot[778].X = 7.78
	pointsOfPolynomialPlot[778].Y = 3_140.027

	pointsOfPolynomialPlot[779].X = 7.79
	pointsOfPolynomialPlot[779].Y = 3_156.938

	pointsOfPolynomialPlot[780].X = 7.8
	pointsOfPolynomialPlot[780].Y = 3_173.913

	pointsOfPolynomialPlot[781].X = 7.81
	pointsOfPolynomialPlot[781].Y = 3_190.958

	pointsOfPolynomialPlot[782].X = 7.82
	pointsOfPolynomialPlot[782].Y = 3_208.071

	pointsOfPolynomialPlot[783].X = 7.83
	pointsOfPolynomialPlot[783].Y = 3_225.253

	pointsOfPolynomialPlot[784].X = 7.84
	pointsOfPolynomialPlot[784].Y = 3_242.504

	pointsOfPolynomialPlot[785].X = 7.85
	pointsOfPolynomialPlot[785].Y = 3_259.823

	pointsOfPolynomialPlot[786].X = 7.86
	pointsOfPolynomialPlot[786].Y = 3_277.211

	pointsOfPolynomialPlot[787].X = 7.87
	pointsOfPolynomialPlot[787].Y = 3_294.669

	pointsOfPolynomialPlot[788].X = 7.88
	pointsOfPolynomialPlot[788].Y = 3_312.196

	pointsOfPolynomialPlot[789].X = 7.89
	pointsOfPolynomialPlot[789].Y = 3_329.792

	pointsOfPolynomialPlot[790].X = 7.9
	pointsOfPolynomialPlot[790].Y = 3_347.459

	pointsOfPolynomialPlot[791].X = 7.91
	pointsOfPolynomialPlot[791].Y = 3_365.195

	pointsOfPolynomialPlot[792].X = 7.92
	pointsOfPolynomialPlot[792].Y = 3_383.001

	pointsOfPolynomialPlot[793].X = 7.93
	pointsOfPolynomialPlot[793].Y = 3_400.878

	pointsOfPolynomialPlot[794].X = 7.94
	pointsOfPolynomialPlot[794].Y = 3_418.825

	pointsOfPolynomialPlot[795].X = 7.95
	pointsOfPolynomialPlot[795].Y = 3_436.843

	pointsOfPolynomialPlot[796].X = 7.96
	pointsOfPolynomialPlot[796].Y = 3_454.932

	pointsOfPolynomialPlot[797].X = 7.97
	pointsOfPolynomialPlot[797].Y = 3_473.092

	pointsOfPolynomialPlot[798].X = 7.98
	pointsOfPolynomialPlot[798].Y = 3_491.323

	pointsOfPolynomialPlot[799].X = 7.99
	pointsOfPolynomialPlot[799].Y = 3_509.625

	pointsOfPolynomialPlot[800].X = 8.0
	pointsOfPolynomialPlot[800].Y = 3_528.0

	pointsOfPolynomialPlot[801].X = 8.01
	pointsOfPolynomialPlot[801].Y = 3_546.445

	pointsOfPolynomialPlot[802].X = 8.02
	pointsOfPolynomialPlot[802].Y = 3_564.963

	pointsOfPolynomialPlot[803].X = 8.03
	pointsOfPolynomialPlot[803].Y = 3_583.553

	pointsOfPolynomialPlot[804].X = 8.04
	pointsOfPolynomialPlot[804].Y = 3_602.216

	pointsOfPolynomialPlot[805].X = 8.05
	pointsOfPolynomialPlot[805].Y = 3_620.951

	pointsOfPolynomialPlot[806].X = 8.06
	pointsOfPolynomialPlot[806].Y = 3_639.759

	pointsOfPolynomialPlot[807].X = 8.07
	pointsOfPolynomialPlot[807].Y = 3_658.639

	pointsOfPolynomialPlot[808].X = 8.08
	pointsOfPolynomialPlot[808].Y = 3_677.593

	pointsOfPolynomialPlot[809].X = 8.09
	pointsOfPolynomialPlot[809].Y = 3_696.62

	pointsOfPolynomialPlot[810].X = 8.1
	pointsOfPolynomialPlot[810].Y = 3_715.721

	pointsOfPolynomialPlot[811].X = 8.11
	pointsOfPolynomialPlot[811].Y = 3_734.895

	pointsOfPolynomialPlot[812].X = 8.12
	pointsOfPolynomialPlot[812].Y = 3_754.143

	pointsOfPolynomialPlot[813].X = 8.13
	pointsOfPolynomialPlot[813].Y = 3_773.465

	pointsOfPolynomialPlot[814].X = 8.14
	pointsOfPolynomialPlot[814].Y = 3_792.861

	pointsOfPolynomialPlot[815].X = 8.15
	pointsOfPolynomialPlot[815].Y = 3_812.332

	pointsOfPolynomialPlot[816].X = 8.16
	pointsOfPolynomialPlot[816].Y = 3_831.878

	pointsOfPolynomialPlot[817].X = 8.17
	pointsOfPolynomialPlot[817].Y = 3_851.498

	pointsOfPolynomialPlot[818].X = 8.18
	pointsOfPolynomialPlot[818].Y = 3_871.193

	pointsOfPolynomialPlot[819].X = 8.19
	pointsOfPolynomialPlot[819].Y = 3_890.963

	pointsOfPolynomialPlot[820].X = 8.2
	pointsOfPolynomialPlot[820].Y = 3_910.809

	pointsOfPolynomialPlot[821].X = 8.21
	pointsOfPolynomialPlot[821].Y = 3_930.73

	pointsOfPolynomialPlot[822].X = 8.22
	pointsOfPolynomialPlot[822].Y = 3_950.728

	pointsOfPolynomialPlot[823].X = 8.23
	pointsOfPolynomialPlot[823].Y = 3_970.801

	pointsOfPolynomialPlot[824].X = 8.24
	pointsOfPolynomialPlot[824].Y = 3_990.95

	pointsOfPolynomialPlot[825].X = 8.25
	pointsOfPolynomialPlot[825].Y = 4_011.175

	pointsOfPolynomialPlot[826].X = 8.26
	pointsOfPolynomialPlot[826].Y = 4_031.477

	pointsOfPolynomialPlot[827].X = 8.27
	pointsOfPolynomialPlot[827].Y = 4_051.856

	pointsOfPolynomialPlot[828].X = 8.28
	pointsOfPolynomialPlot[828].Y = 4_072.312

	pointsOfPolynomialPlot[829].X = 8.29
	pointsOfPolynomialPlot[829].Y = 4_092.845

	pointsOfPolynomialPlot[830].X = 8.3
	pointsOfPolynomialPlot[830].Y = 4_114.455

	pointsOfPolynomialPlot[831].X = 8.31
	pointsOfPolynomialPlot[831].Y = 4_134.142

	pointsOfPolynomialPlot[832].X = 8.32
	pointsOfPolynomialPlot[832].Y = 4_154.907

	pointsOfPolynomialPlot[833].X = 8.33
	pointsOfPolynomialPlot[833].Y = 4_175.751

	pointsOfPolynomialPlot[834].X = 8.34
	pointsOfPolynomialPlot[834].Y = 4_196.672

	pointsOfPolynomialPlot[835].X = 8.35
	pointsOfPolynomialPlot[835].Y = 4_217.671

	pointsOfPolynomialPlot[836].X = 8.36
	pointsOfPolynomialPlot[836].Y = 4_238.749

	pointsOfPolynomialPlot[837].X = 8.37
	pointsOfPolynomialPlot[837].Y = 4_259.906

	pointsOfPolynomialPlot[838].X = 8.38
	pointsOfPolynomialPlot[838].Y = 4_281.141

	pointsOfPolynomialPlot[839].X = 8.39
	pointsOfPolynomialPlot[839].Y = 4_302.455

	pointsOfPolynomialPlot[840].X = 8.4
	pointsOfPolynomialPlot[840].Y = 4_323.849

	pointsOfPolynomialPlot[841].X = 8.41
	pointsOfPolynomialPlot[841].Y = 4_345.322

	pointsOfPolynomialPlot[842].X = 8.42
	pointsOfPolynomialPlot[842].Y = 4_366.875

	pointsOfPolynomialPlot[843].X = 8.43
	pointsOfPolynomialPlot[843].Y = 4_388.508

	pointsOfPolynomialPlot[844].X = 8.44
	pointsOfPolynomialPlot[844].Y = 4_410.22

	pointsOfPolynomialPlot[845].X = 8.45
	pointsOfPolynomialPlot[845].Y = 4_432.013

	pointsOfPolynomialPlot[846].X = 8.46
	pointsOfPolynomialPlot[846].Y = 4_453.886

	pointsOfPolynomialPlot[847].X = 8.47
	pointsOfPolynomialPlot[847].Y = 4_475.84

	pointsOfPolynomialPlot[848].X = 8.48
	pointsOfPolynomialPlot[848].Y = 4_497.875

	pointsOfPolynomialPlot[849].X = 8.49
	pointsOfPolynomialPlot[849].Y = 4_519.99

	pointsOfPolynomialPlot[850].X = 8.5
	pointsOfPolynomialPlot[850].Y = 4_542.187

	pointsOfPolynomialPlot[851].X = 8.51
	pointsOfPolynomialPlot[851].Y = 4_564.465

	pointsOfPolynomialPlot[852].X = 8.52
	pointsOfPolynomialPlot[852].Y = 4_586.825

	pointsOfPolynomialPlot[853].X = 8.53
	pointsOfPolynomialPlot[853].Y = 4_609.267

	pointsOfPolynomialPlot[854].X = 8.54
	pointsOfPolynomialPlot[854].Y = 4_631.79

	pointsOfPolynomialPlot[855].X = 8.55
	pointsOfPolynomialPlot[855].Y = 4_654.396

	pointsOfPolynomialPlot[856].X = 8.56
	pointsOfPolynomialPlot[856].Y = 4_677.084

	pointsOfPolynomialPlot[857].X = 8.57
	pointsOfPolynomialPlot[857].Y = 4_699.855

	pointsOfPolynomialPlot[858].X = 8.58
	pointsOfPolynomialPlot[858].Y = 4_722.709

	pointsOfPolynomialPlot[859].X = 8.59
	pointsOfPolynomialPlot[859].Y = 4_745.645

	pointsOfPolynomialPlot[860].X = 8.6
	pointsOfPolynomialPlot[860].Y = 4_768.665

	pointsOfPolynomialPlot[861].X = 8.61
	pointsOfPolynomialPlot[861].Y = 4_791.768

	pointsOfPolynomialPlot[862].X = 8.62
	pointsOfPolynomialPlot[862].Y = 4_814.955

	pointsOfPolynomialPlot[863].X = 8.63
	pointsOfPolynomialPlot[863].Y = 4_838.226

	pointsOfPolynomialPlot[864].X = 8.64
	pointsOfPolynomialPlot[864].Y = 4_861.58

	pointsOfPolynomialPlot[865].X = 8.65
	pointsOfPolynomialPlot[865].Y = 4_885.019

	pointsOfPolynomialPlot[866].X = 8.66
	pointsOfPolynomialPlot[866].Y = 4_908.542

	pointsOfPolynomialPlot[867].X = 8.67
	pointsOfPolynomialPlot[867].Y = 4_932.15

	pointsOfPolynomialPlot[868].X = 8.68
	pointsOfPolynomialPlot[868].Y = 4_955.842

	pointsOfPolynomialPlot[869].X = 8.69
	pointsOfPolynomialPlot[869].Y = 4_979.62

	pointsOfPolynomialPlot[870].X = 8.7
	pointsOfPolynomialPlot[870].Y = 5_003.483

	pointsOfPolynomialPlot[871].X = 8.71
	pointsOfPolynomialPlot[871].Y = 5_027.431

	pointsOfPolynomialPlot[872].X = 8.72
	pointsOfPolynomialPlot[872].Y = 5_051.465

	pointsOfPolynomialPlot[873].X = 8.73
	pointsOfPolynomialPlot[873].Y = 5_075.584

	pointsOfPolynomialPlot[874].X = 8.74
	pointsOfPolynomialPlot[874].Y = 5_099.79

	pointsOfPolynomialPlot[875].X = 8.75
	pointsOfPolynomialPlot[875].Y = 5_124.082

	pointsOfPolynomialPlot[876].X = 8.76
	pointsOfPolynomialPlot[876].Y = 5_148.46

	pointsOfPolynomialPlot[877].X = 8.77
	pointsOfPolynomialPlot[877].Y = 5_172.925

	pointsOfPolynomialPlot[878].X = 8.78
	pointsOfPolynomialPlot[878].Y = 5_197.476

	pointsOfPolynomialPlot[879].X = 8.79
	pointsOfPolynomialPlot[879].Y = 5_222.115

	pointsOfPolynomialPlot[880].X = 8.8
	pointsOfPolynomialPlot[880].Y = 5_246.841

	pointsOfPolynomialPlot[881].X = 8.81
	pointsOfPolynomialPlot[881].Y = 5_271.655

	pointsOfPolynomialPlot[882].X = 8.82
	pointsOfPolynomialPlot[882].Y = 5_296.556

	pointsOfPolynomialPlot[883].X = 8.83
	pointsOfPolynomialPlot[883].Y = 5_321.545

	pointsOfPolynomialPlot[884].X = 8.84
	pointsOfPolynomialPlot[884].Y = 5_346.622

	pointsOfPolynomialPlot[885].X = 8.85
	pointsOfPolynomialPlot[885].Y = 5_371.787

	pointsOfPolynomialPlot[886].X = 8.86
	pointsOfPolynomialPlot[886].Y = 5_397.041

	pointsOfPolynomialPlot[887].X = 8.87
	pointsOfPolynomialPlot[887].Y = 5_422.383

	pointsOfPolynomialPlot[888].X = 8.88
	pointsOfPolynomialPlot[888].Y = 5_447.814

	pointsOfPolynomialPlot[889].X = 8.89
	pointsOfPolynomialPlot[889].Y = 5_473.335

	pointsOfPolynomialPlot[890].X = 8.9
	pointsOfPolynomialPlot[890].Y = 5_498.945

	pointsOfPolynomialPlot[891].X = 8.91
	pointsOfPolynomialPlot[891].Y = 5_524.644

	pointsOfPolynomialPlot[892].X = 8.92
	pointsOfPolynomialPlot[892].Y = 5_550.433

	pointsOfPolynomialPlot[893].X = 8.93
	pointsOfPolynomialPlot[893].Y = 5_576.312

	pointsOfPolynomialPlot[894].X = 8.94
	pointsOfPolynomialPlot[894].Y = 5_602.281

	pointsOfPolynomialPlot[895].X = 8.95
	pointsOfPolynomialPlot[895].Y = 5_628.34

	pointsOfPolynomialPlot[896].X = 8.96
	pointsOfPolynomialPlot[896].Y = 5_654.49

	pointsOfPolynomialPlot[897].X = 8.97
	pointsOfPolynomialPlot[897].Y = 5_680.731

	pointsOfPolynomialPlot[898].X = 8.98
	pointsOfPolynomialPlot[898].Y = 5_707.062

	pointsOfPolynomialPlot[899].X = 8.99
	pointsOfPolynomialPlot[899].Y = 5_733.485

	pointsOfPolynomialPlot[900].X = 9.0
	pointsOfPolynomialPlot[900].Y = 5_760.0

	pointsOfPolynomialPlot[901].X = 9.01
	pointsOfPolynomialPlot[901].Y = 5_786.605

	pointsOfPolynomialPlot[902].X = 9.02
	pointsOfPolynomialPlot[902].Y = 5_813.303

	pointsOfPolynomialPlot[903].X = 9.03
	pointsOfPolynomialPlot[903].Y = 5_840.093

	pointsOfPolynomialPlot[904].X = 9.04
	pointsOfPolynomialPlot[904].Y = 5_866.975

	pointsOfPolynomialPlot[905].X = 9.05
	pointsOfPolynomialPlot[905].Y = 5_893.949

	pointsOfPolynomialPlot[906].X = 9.06
	pointsOfPolynomialPlot[906].Y = 5_921.016

	pointsOfPolynomialPlot[907].X = 9.07
	pointsOfPolynomialPlot[907].Y = 5_948.176

	pointsOfPolynomialPlot[908].X = 9.08
	pointsOfPolynomialPlot[908].Y = 5_975.429

	pointsOfPolynomialPlot[909].X = 9.09
	pointsOfPolynomialPlot[909].Y = 6_002.775

	pointsOfPolynomialPlot[910].X = 9.1
	pointsOfPolynomialPlot[910].Y = 6_030.215

	pointsOfPolynomialPlot[911].X = 9.11
	pointsOfPolynomialPlot[911].Y = 6_057.748

	pointsOfPolynomialPlot[912].X = 9.12
	pointsOfPolynomialPlot[912].Y = 6_085.375

	pointsOfPolynomialPlot[913].X = 9.13
	pointsOfPolynomialPlot[913].Y = 6_113.097

	pointsOfPolynomialPlot[914].X = 9.14
	pointsOfPolynomialPlot[914].Y = 6_140.913

	pointsOfPolynomialPlot[915].X = 9.15
	pointsOfPolynomialPlot[915].Y = 6_168.823

	pointsOfPolynomialPlot[916].X = 9.16
	pointsOfPolynomialPlot[916].Y = 6_196.828

	pointsOfPolynomialPlot[917].X = 9.17
	pointsOfPolynomialPlot[917].Y = 6_224.929

	pointsOfPolynomialPlot[918].X = 9.18
	pointsOfPolynomialPlot[918].Y = 6_253.415

	pointsOfPolynomialPlot[919].X = 9.19
	pointsOfPolynomialPlot[919].Y = 6_281.415

	pointsOfPolynomialPlot[920].X = 9.2
	pointsOfPolynomialPlot[920].Y = 6_309.801

	pointsOfPolynomialPlot[921].X = 9.21
	pointsOfPolynomialPlot[921].Y = 6_338.283

	pointsOfPolynomialPlot[922].X = 9.22
	pointsOfPolynomialPlot[922].Y = 6_366.862

	pointsOfPolynomialPlot[923].X = 9.23
	pointsOfPolynomialPlot[923].Y = 6_395.536

	pointsOfPolynomialPlot[924].X = 9.24
	pointsOfPolynomialPlot[924].Y = 6_424.308

	pointsOfPolynomialPlot[925].X = 9.25
	pointsOfPolynomialPlot[925].Y = 6_453.175

	pointsOfPolynomialPlot[926].X = 9.26
	pointsOfPolynomialPlot[926].Y = 6_482.14

	pointsOfPolynomialPlot[927].X = 9.27
	pointsOfPolynomialPlot[927].Y = 6_511.202

	pointsOfPolynomialPlot[928].X = 9.28
	pointsOfPolynomialPlot[928].Y = 6_540.361

	pointsOfPolynomialPlot[929].X = 9.29
	pointsOfPolynomialPlot[929].Y = 6_569.618

	pointsOfPolynomialPlot[930].X = 9.3
	pointsOfPolynomialPlot[930].Y = 6_598.973

	pointsOfPolynomialPlot[931].X = 9.31
	pointsOfPolynomialPlot[931].Y = 6_628.425

	pointsOfPolynomialPlot[932].X = 9.32
	pointsOfPolynomialPlot[932].Y = 6_657.976

	pointsOfPolynomialPlot[933].X = 9.33
	pointsOfPolynomialPlot[933].Y = 6_687.625

	pointsOfPolynomialPlot[934].X = 9.34
	pointsOfPolynomialPlot[934].Y = 6_717.22

	pointsOfPolynomialPlot[935].X = 9.35
	pointsOfPolynomialPlot[935].Y = 6_747.22

	pointsOfPolynomialPlot[936].X = 9.36
	pointsOfPolynomialPlot[936].Y = 6_777.166

	pointsOfPolynomialPlot[937].X = 9.37
	pointsOfPolynomialPlot[937].Y = 6_807.211

	pointsOfPolynomialPlot[938].X = 9.38
	pointsOfPolynomialPlot[938].Y = 6_837.356

	pointsOfPolynomialPlot[939].X = 9.39
	pointsOfPolynomialPlot[939].Y = 6_867.601

	pointsOfPolynomialPlot[940].X = 9.4
	pointsOfPolynomialPlot[940].Y = 6_897.945

	pointsOfPolynomialPlot[941].X = 9.41
	pointsOfPolynomialPlot[941].Y = 6_928.39

	pointsOfPolynomialPlot[942].X = 9.42
	pointsOfPolynomialPlot[942].Y = 6_958.935

	pointsOfPolynomialPlot[943].X = 9.43
	pointsOfPolynomialPlot[943].Y = 6_989.581

	pointsOfPolynomialPlot[944].X = 9.44
	pointsOfPolynomialPlot[944].Y = 7_020.327

	pointsOfPolynomialPlot[945].X = 9.45
	pointsOfPolynomialPlot[945].Y = 7_051.175

	pointsOfPolynomialPlot[946].X = 9.46
	pointsOfPolynomialPlot[946].Y = 7_082.124

	pointsOfPolynomialPlot[947].X = 9.47
	pointsOfPolynomialPlot[947].Y = 7_113.174

	pointsOfPolynomialPlot[948].X = 9.48
	pointsOfPolynomialPlot[948].Y = 7_144.327

	pointsOfPolynomialPlot[949].X = 9.49
	pointsOfPolynomialPlot[949].Y = 7_175.581

	pointsOfPolynomialPlot[950].X = 9.5
	pointsOfPolynomialPlot[950].Y = 7_206.937

	pointsOfPolynomialPlot[951].X = 9.51
	pointsOfPolynomialPlot[951].Y = 7_238.396

	pointsOfPolynomialPlot[952].X = 9.52
	pointsOfPolynomialPlot[952].Y = 7_269.957

	pointsOfPolynomialPlot[953].X = 9.53
	pointsOfPolynomialPlot[953].Y = 7_301.621

	pointsOfPolynomialPlot[954].X = 9.54
	pointsOfPolynomialPlot[954].Y = 7_333.389

	pointsOfPolynomialPlot[955].X = 9.55
	pointsOfPolynomialPlot[955].Y = 7_365.259

	pointsOfPolynomialPlot[956].X = 9.56
	pointsOfPolynomialPlot[956].Y = 7_397.233

	pointsOfPolynomialPlot[957].X = 9.57
	pointsOfPolynomialPlot[957].Y = 7_429.311

	pointsOfPolynomialPlot[958].X = 9.58
	pointsOfPolynomialPlot[958].Y = 7_461.493

	pointsOfPolynomialPlot[959].X = 9.59
	pointsOfPolynomialPlot[959].Y = 7_493.779

	pointsOfPolynomialPlot[960].X = 9.6
	pointsOfPolynomialPlot[960].Y = 7_526.169

	pointsOfPolynomialPlot[961].X = 9.61
	pointsOfPolynomialPlot[961].Y = 7_558.664

	pointsOfPolynomialPlot[962].X = 9.62
	pointsOfPolynomialPlot[962].Y = 7_591.264

	pointsOfPolynomialPlot[963].X = 9.63
	pointsOfPolynomialPlot[963].Y = 7_623.969

	pointsOfPolynomialPlot[964].X = 9.64
	pointsOfPolynomialPlot[964].Y = 7_656.779

	pointsOfPolynomialPlot[965].X = 9.65
	pointsOfPolynomialPlot[965].Y = 7_689.695

	pointsOfPolynomialPlot[966].X = 9.66
	pointsOfPolynomialPlot[966].Y = 7_722.716

	pointsOfPolynomialPlot[967].X = 9.67
	pointsOfPolynomialPlot[967].Y = 7_755.844

	pointsOfPolynomialPlot[968].X = 9.68
	pointsOfPolynomialPlot[968].Y = 7_789.078

	pointsOfPolynomialPlot[969].X = 9.69
	pointsOfPolynomialPlot[969].Y = 7_822.418

	pointsOfPolynomialPlot[970].X = 9.7
	pointsOfPolynomialPlot[970].Y = 7_855.865

	pointsOfPolynomialPlot[971].X = 9.71
	pointsOfPolynomialPlot[971].Y = 7_889.418

	pointsOfPolynomialPlot[972].X = 9.72
	pointsOfPolynomialPlot[972].Y = 7_923.079

	pointsOfPolynomialPlot[973].X = 9.73
	pointsOfPolynomialPlot[973].Y = 7_956.847

	pointsOfPolynomialPlot[974].X = 9.74
	pointsOfPolynomialPlot[974].Y = 7_990.723

	pointsOfPolynomialPlot[975].X = 9.75
	pointsOfPolynomialPlot[975].Y = 8_024.707

	pointsOfPolynomialPlot[976].X = 9.76
	pointsOfPolynomialPlot[976].Y = 8_058.798

	pointsOfPolynomialPlot[977].X = 9.77
	pointsOfPolynomialPlot[977].Y = 8_092.998

	pointsOfPolynomialPlot[978].X = 9.78
	pointsOfPolynomialPlot[978].Y = 8_127.306

	pointsOfPolynomialPlot[979].X = 9.79
	pointsOfPolynomialPlot[979].Y = 8_161.723

	pointsOfPolynomialPlot[980].X = 9.8
	pointsOfPolynomialPlot[980].Y = 8_196.249

	pointsOfPolynomialPlot[981].X = 9.81
	pointsOfPolynomialPlot[981].Y = 8_230.884

	pointsOfPolynomialPlot[982].X = 9.82
	pointsOfPolynomialPlot[982].Y = 8_265.629

	pointsOfPolynomialPlot[983].X = 9.83
	pointsOfPolynomialPlot[983].Y = 8_300.483

	pointsOfPolynomialPlot[984].X = 9.84
	pointsOfPolynomialPlot[984].Y = 8_335.447

	pointsOfPolynomialPlot[985].X = 9.85
	pointsOfPolynomialPlot[985].Y = 8_370.521

	pointsOfPolynomialPlot[986].X = 9.86
	pointsOfPolynomialPlot[986].Y = 8_405.705

	pointsOfPolynomialPlot[987].X = 9.87
	pointsOfPolynomialPlot[987].Y = 8_441.0

	pointsOfPolynomialPlot[988].X = 9.88
	pointsOfPolynomialPlot[988].Y = 8_476.406

	pointsOfPolynomialPlot[989].X = 9.89
	pointsOfPolynomialPlot[989].Y = 8_511.923

	pointsOfPolynomialPlot[990].X = 9.9
	pointsOfPolynomialPlot[990].Y = 8_547.551

	pointsOfPolynomialPlot[991].X = 9.91
	pointsOfPolynomialPlot[991].Y = 8_583.29

	pointsOfPolynomialPlot[992].X = 9.92
	pointsOfPolynomialPlot[992].Y = 8_619.141

	pointsOfPolynomialPlot[993].X = 9.93
	pointsOfPolynomialPlot[993].Y = 8_655.104

	pointsOfPolynomialPlot[994].X = 9.94
	pointsOfPolynomialPlot[994].Y = 8_691.18

	pointsOfPolynomialPlot[995].X = 9.95
	pointsOfPolynomialPlot[995].Y = 8_727.367

	pointsOfPolynomialPlot[996].X = 9.96
	pointsOfPolynomialPlot[996].Y = 8_763.667

	pointsOfPolynomialPlot[997].X = 9.97
	pointsOfPolynomialPlot[997].Y = 8_800.081

	pointsOfPolynomialPlot[998].X = 9.98
	pointsOfPolynomialPlot[998].Y = 8_836.607

	pointsOfPolynomialPlot[999].X = 9.99
	pointsOfPolynomialPlot[999].Y = 8_873.246

	pointsOfPolynomialPlot[1_000].X = 10.0
	pointsOfPolynomialPlot[1_000].Y = 8_910.0










	polynomialPlot := plot.New()

	polynomialPlot.Title.Text = "Plot of polynomial f(x) = x^4 - x^3 - x^2 + x"

	polynomialPlot.X.Label.Text = "x"
	polynomialPlot.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfPolynomialPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	polynomialPlot.Add(plotLine)
	polynomialPlot.Legend.Add("f(x)", plotLine)

	if err := polynomialPlot.Save(10*vg.Inch, 10*vg.Inch,
		"Polynomial-plot-02.png"); err != nil {

		panic(err)
	}
}
