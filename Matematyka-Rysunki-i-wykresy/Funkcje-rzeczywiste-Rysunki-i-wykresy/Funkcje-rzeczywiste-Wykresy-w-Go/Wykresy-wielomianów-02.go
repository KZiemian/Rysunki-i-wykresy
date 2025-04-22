package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres wielomianu f(x) = x^4 - x^3 - x^2 + x

	punktyWykresuWielomianu := make(plotter.XYs, 901)

	punktyWykresuWielomianu[0].X = 0.0
	punktyWykresuWielomianu[0].Y = 0.0

	punktyWykresuWielomianu[1].X = 0.01
	punktyWykresuWielomianu[1].Y = 0.009

	punktyWykresuWielomianu[2].X = 0.02
	punktyWykresuWielomianu[2].Y = 0.019

	punktyWykresuWielomianu[3].X = 0.03
	punktyWykresuWielomianu[3].Y = 0.029

	punktyWykresuWielomianu[4].X = 0.04
	punktyWykresuWielomianu[4].Y = 0.038

	punktyWykresuWielomianu[5].X = 0.05
	punktyWykresuWielomianu[5].Y = 0.047

	punktyWykresuWielomianu[6].X = 0.06
	punktyWykresuWielomianu[6].Y = 0.056

	punktyWykresuWielomianu[7].X = 0.07
	punktyWykresuWielomianu[7].Y = 0.064

	punktyWykresuWielomianu[8].X = 0.08
	punktyWykresuWielomianu[8].Y = 0.073

	punktyWykresuWielomianu[9].X = 0.09
	punktyWykresuWielomianu[9].Y = 0.081

	punktyWykresuWielomianu[10].X = 0.1
	punktyWykresuWielomianu[10].Y = 0.089

	punktyWykresuWielomianu[11].X = 0.11
	punktyWykresuWielomianu[11].Y = 0.096

	punktyWykresuWielomianu[12].X = 0.12
	punktyWykresuWielomianu[12].Y = 0.104

	punktyWykresuWielomianu[13].X = 0.13
	punktyWykresuWielomianu[13].Y = 0.111

	punktyWykresuWielomianu[14].X = 0.14
	punktyWykresuWielomianu[14].Y = 0.118

	punktyWykresuWielomianu[15].X = 0.15
	punktyWykresuWielomianu[15].Y = 0.124

	punktyWykresuWielomianu[16].X = 0.16
	punktyWykresuWielomianu[16].Y = 0.131

	punktyWykresuWielomianu[17].X = 0.17
	punktyWykresuWielomianu[17].Y = 0.137

	punktyWykresuWielomianu[18].X = 0.18
	punktyWykresuWielomianu[18].Y = 0.142

	punktyWykresuWielomianu[19].X = 0.19
	punktyWykresuWielomianu[19].Y = 0.148

	punktyWykresuWielomianu[20].X = 0.2
	punktyWykresuWielomianu[20].Y = 0.153

	punktyWykresuWielomianu[21].X = 0.21
	punktyWykresuWielomianu[21].Y = 0.158

	punktyWykresuWielomianu[22].X = 0.22
	punktyWykresuWielomianu[22].Y = 0.163

	punktyWykresuWielomianu[23].X = 0.23
	punktyWykresuWielomianu[23].Y = 0.167

	punktyWykresuWielomianu[24].X = 0.24
	punktyWykresuWielomianu[24].Y = 0.171

	punktyWykresuWielomianu[25].X = 0.25
	punktyWykresuWielomianu[25].Y = 0.175

	punktyWykresuWielomianu[26].X = 0.26
	punktyWykresuWielomianu[26].Y = 0.179

	punktyWykresuWielomianu[27].X = 0.27
	punktyWykresuWielomianu[27].Y = 0.182

	punktyWykresuWielomianu[28].X = 0.28
	punktyWykresuWielomianu[28].Y = 0.185

	punktyWykresuWielomianu[29].X = 0.29
	punktyWykresuWielomianu[29].Y = 0.188

	punktyWykresuWielomianu[30].X = 0.3
	punktyWykresuWielomianu[30].Y = 0.191

	punktyWykresuWielomianu[31].X = 0.31
	punktyWykresuWielomianu[31].Y = 0.193

	punktyWykresuWielomianu[32].X = 0.32
	punktyWykresuWielomianu[32].Y = 0.195

	punktyWykresuWielomianu[33].X = 0.33
	punktyWykresuWielomianu[33].Y = 0.197

	punktyWykresuWielomianu[34].X = 0.34
	punktyWykresuWielomianu[34].Y = 0.198

	punktyWykresuWielomianu[35].X = 0.35
	punktyWykresuWielomianu[35].Y = 0.199

	punktyWykresuWielomianu[36].X = 0.36
	punktyWykresuWielomianu[36].Y = 0.2

	punktyWykresuWielomianu[37].X = 0.37
	punktyWykresuWielomianu[37].Y = 0.201

	punktyWykresuWielomianu[38].X = 0.38
	punktyWykresuWielomianu[38].Y = 0.201

	punktyWykresuWielomianu[39].X = 0.39
	punktyWykresuWielomianu[39].Y = 0.201

	punktyWykresuWielomianu[40].X = 0.4
	punktyWykresuWielomianu[40].Y = 0.201

	punktyWykresuWielomianu[41].X = 0.41
	punktyWykresuWielomianu[41].Y = 0.201

	punktyWykresuWielomianu[42].X = 0.42
	punktyWykresuWielomianu[42].Y = 0.2

	punktyWykresuWielomianu[43].X = 0.43
	punktyWykresuWielomianu[43].Y = 0.199

	punktyWykresuWielomianu[44].X = 0.44
	punktyWykresuWielomianu[44].Y = 0.198

	punktyWykresuWielomianu[45].X = 0.45
	punktyWykresuWielomianu[45].Y = 0.197

	punktyWykresuWielomianu[46].X = 0.46
	punktyWykresuWielomianu[46].Y = 0.195

	punktyWykresuWielomianu[47].X = 0.47
	punktyWykresuWielomianu[47].Y = 0.194

	punktyWykresuWielomianu[48].X = 0.48
	punktyWykresuWielomianu[48].Y = 0.192

	punktyWykresuWielomianu[49].X = 0.49
	punktyWykresuWielomianu[49].Y = 0.189

	punktyWykresuWielomianu[50].X = 0.5
	punktyWykresuWielomianu[50].Y = 0.187

	punktyWykresuWielomianu[51].X = 0.51
	punktyWykresuWielomianu[51].Y = 0.184

	punktyWykresuWielomianu[52].X = 0.52
	punktyWykresuWielomianu[52].Y = 0.182

	punktyWykresuWielomianu[53].X = 0.53
	punktyWykresuWielomianu[53].Y = 0.179

	punktyWykresuWielomianu[54].X = 0.54
	punktyWykresuWielomianu[54].Y = 0.176

	punktyWykresuWielomianu[55].X = 0.55
	punktyWykresuWielomianu[55].Y = 0.172

	punktyWykresuWielomianu[56].X = 0.56
	punktyWykresuWielomianu[56].Y = 0.169

	punktyWykresuWielomianu[57].X = 0.57
	punktyWykresuWielomianu[57].Y = 0.165

	punktyWykresuWielomianu[58].X = 0.58
	punktyWykresuWielomianu[58].Y = 0.161

	punktyWykresuWielomianu[59].X = 0.59
	punktyWykresuWielomianu[59].Y = 0.157

	punktyWykresuWielomianu[60].X = 0.6
	punktyWykresuWielomianu[60].Y = 0.153

	punktyWykresuWielomianu[61].X = 0.61
	punktyWykresuWielomianu[61].Y = 0.149

	punktyWykresuWielomianu[62].X = 0.62
	punktyWykresuWielomianu[62].Y = 0.145

	punktyWykresuWielomianu[63].X = 0.63
	punktyWykresuWielomianu[63].Y = 0.14

	punktyWykresuWielomianu[64].X = 0.64
	punktyWykresuWielomianu[64].Y = 0.136

	punktyWykresuWielomianu[65].X = 0.65
	punktyWykresuWielomianu[65].Y = 0.131

	punktyWykresuWielomianu[66].X = 0.66
	punktyWykresuWielomianu[66].Y = 0.126

	punktyWykresuWielomianu[67].X = 0.67
	punktyWykresuWielomianu[67].Y = 0.121

	punktyWykresuWielomianu[68].X = 0.68
	punktyWykresuWielomianu[68].Y = 0.117

	punktyWykresuWielomianu[69].X = 0.69
	punktyWykresuWielomianu[69].Y = 0.112

	punktyWykresuWielomianu[70].X = 0.7
	punktyWykresuWielomianu[70].Y = 0.107

	punktyWykresuWielomianu[71].X = 0.71
	punktyWykresuWielomianu[71].Y = 0.102

	punktyWykresuWielomianu[72].X = 0.72
	punktyWykresuWielomianu[72].Y = 0.097

	punktyWykresuWielomianu[73].X = 0.73
	punktyWykresuWielomianu[73].Y = 0.092

	punktyWykresuWielomianu[74].X = 0.74
	punktyWykresuWielomianu[74].Y = 0.087

	punktyWykresuWielomianu[75].X = 0.75
	punktyWykresuWielomianu[75].Y = 0.082

	punktyWykresuWielomianu[76].X = 0.76
	punktyWykresuWielomianu[76].Y = 0.077

	punktyWykresuWielomianu[77].X = 0.77
	punktyWykresuWielomianu[77].Y = 0.072

	punktyWykresuWielomianu[78].X = 0.78
	punktyWykresuWielomianu[78].Y = 0.067

	punktyWykresuWielomianu[79].X = 0.79
	punktyWykresuWielomianu[79].Y = 0.062

	punktyWykresuWielomianu[80].X = 0.8
	punktyWykresuWielomianu[80].Y = 0.057

	punktyWykresuWielomianu[81].X = 0.81
	punktyWykresuWielomianu[81].Y = 0.052

	punktyWykresuWielomianu[82].X = 0.82
	punktyWykresuWielomianu[82].Y = 0.048

	punktyWykresuWielomianu[83].X = 0.83
	punktyWykresuWielomianu[83].Y = 0.043

	punktyWykresuWielomianu[84].X = 0.84
	punktyWykresuWielomianu[84].Y = 0.039

	punktyWykresuWielomianu[85].X = 0.85
	punktyWykresuWielomianu[85].Y = 0.035

	punktyWykresuWielomianu[86].X = 0.86
	punktyWykresuWielomianu[86].Y = 0.031

	punktyWykresuWielomianu[87].X = 0.87
	punktyWykresuWielomianu[87].Y = 0.027

	punktyWykresuWielomianu[88].X = 0.88
	punktyWykresuWielomianu[88].Y = 0.023

	punktyWykresuWielomianu[89].X = 0.89
	punktyWykresuWielomianu[89].Y = 0.02

	punktyWykresuWielomianu[90].X = 0.9
	punktyWykresuWielomianu[90].Y = 0.017

	punktyWykresuWielomianu[91].X = 0.91
	punktyWykresuWielomianu[91].Y = 0.014

	punktyWykresuWielomianu[92].X = 0.92
	punktyWykresuWielomianu[92].Y = 0.011

	punktyWykresuWielomianu[93].X = 0.93
	punktyWykresuWielomianu[93].Y = 0.008

	punktyWykresuWielomianu[94].X = 0.94
	punktyWykresuWielomianu[94].Y = 0.006

	punktyWykresuWielomianu[95].X = 0.95
	punktyWykresuWielomianu[95].Y = 0.004

	punktyWykresuWielomianu[96].X = 0.96
	punktyWykresuWielomianu[96].Y = 0.003

	punktyWykresuWielomianu[97].X = 0.97
	punktyWykresuWielomianu[97].Y = 0.001

	punktyWykresuWielomianu[98].X = 0.98
	punktyWykresuWielomianu[98].Y = 0.0

	punktyWykresuWielomianu[99].X = 0.99
	punktyWykresuWielomianu[99].Y = 0.0

	punktyWykresuWielomianu[100].X = 1.0
	punktyWykresuWielomianu[100].Y = 0.0

	punktyWykresuWielomianu[101].X = 1.01
	punktyWykresuWielomianu[101].Y = 0.0

	punktyWykresuWielomianu[102].X = 1.02
	punktyWykresuWielomianu[102].Y = 0.0

	punktyWykresuWielomianu[103].X = 1.03
	punktyWykresuWielomianu[103].Y = 0.001

	punktyWykresuWielomianu[104].X = 1.04
	punktyWykresuWielomianu[104].Y = 0.003

	punktyWykresuWielomianu[105].X = 1.05
	punktyWykresuWielomianu[105].Y = 0.005

	punktyWykresuWielomianu[106].X = 1.06
	punktyWykresuWielomianu[106].Y = 0.007

	punktyWykresuWielomianu[107].X = 1.07
	punktyWykresuWielomianu[107].Y = 0.01

	punktyWykresuWielomianu[108].X = 1.08
	punktyWykresuWielomianu[108].Y = 0.014

	punktyWykresuWielomianu[109].X = 1.09
	punktyWykresuWielomianu[109].Y = 0.018

	punktyWykresuWielomianu[110].X = 1.1
	punktyWykresuWielomianu[110].Y = 0.023

	punktyWykresuWielomianu[111].X = 1.11
	punktyWykresuWielomianu[111].Y = 0.028

	punktyWykresuWielomianu[112].X = 1.12
	punktyWykresuWielomianu[112].Y = 0.034

	punktyWykresuWielomianu[113].X = 1.13
	punktyWykresuWielomianu[113].Y = 0.04

	punktyWykresuWielomianu[114].X = 1.14
	punktyWykresuWielomianu[114].Y = 0.047

	punktyWykresuWielomianu[115].X = 1.15
	punktyWykresuWielomianu[115].Y = 0.055

	punktyWykresuWielomianu[116].X = 1.16
	punktyWykresuWielomianu[116].Y = 0.064

	punktyWykresuWielomianu[117].X = 1.17
	punktyWykresuWielomianu[117].Y = 0.073

	punktyWykresuWielomianu[118].X = 1.18
	punktyWykresuWielomianu[118].Y = 0.083

	punktyWykresuWielomianu[119].X = 1.19
	punktyWykresuWielomianu[119].Y = 0.094

	punktyWykresuWielomianu[120].X = 1.2
	punktyWykresuWielomianu[120].Y = 0.105

	punktyWykresuWielomianu[121].X = 1.21
	punktyWykresuWielomianu[121].Y = 0.117

	punktyWykresuWielomianu[122].X = 1.22
	punktyWykresuWielomianu[122].Y = 0.131

	punktyWykresuWielomianu[123].X = 1.23
	punktyWykresuWielomianu[123].Y = 0.145

	punktyWykresuWielomianu[124].X = 1.24
	punktyWykresuWielomianu[124].Y = 0.16

	punktyWykresuWielomianu[125].X = 1.25
	punktyWykresuWielomianu[125].Y = 0.175

	punktyWykresuWielomianu[126].X = 1.26
	punktyWykresuWielomianu[126].Y = 0.192

	punktyWykresuWielomianu[127].X = 1.27
	punktyWykresuWielomianu[127].Y = 0.21

	punktyWykresuWielomianu[128].X = 1.28
	punktyWykresuWielomianu[128].Y = 0.228

	punktyWykresuWielomianu[129].X = 1.29
	punktyWykresuWielomianu[129].Y = 0.248

	punktyWykresuWielomianu[130].X = 1.3
	punktyWykresuWielomianu[130].Y = 0.269

	punktyWykresuWielomianu[131].X = 1.31
	punktyWykresuWielomianu[131].Y = 0.29

	punktyWykresuWielomianu[132].X = 1.32
	punktyWykresuWielomianu[132].Y = 0.313

	punktyWykresuWielomianu[133].X = 1.33
	punktyWykresuWielomianu[133].Y = 0.337

	punktyWykresuWielomianu[134].X = 1.34
	punktyWykresuWielomianu[134].Y = 0.362

	punktyWykresuWielomianu[135].X = 1.35
	punktyWykresuWielomianu[135].Y = 0.388

	punktyWykresuWielomianu[136].X = 1.36
	punktyWykresuWielomianu[136].Y = 0.416

	punktyWykresuWielomianu[137].X = 1.37
	punktyWykresuWielomianu[137].Y = 0.444

	punktyWykresuWielomianu[138].X = 1.38
	punktyWykresuWielomianu[138].Y = 0.474

	punktyWykresuWielomianu[139].X = 1.39
	punktyWykresuWielomianu[139].Y = 0.505

	punktyWykresuWielomianu[140].X = 1.4
	punktyWykresuWielomianu[140].Y = 0.537

	punktyWykresuWielomianu[141].X = 1.41
	punktyWykresuWielomianu[141].Y = 0.571

	punktyWykresuWielomianu[142].X = 1.42
	punktyWykresuWielomianu[142].Y = 0.606

	punktyWykresuWielomianu[143].X = 1.43
	punktyWykresuWielomianu[143].Y = 0.642

	punktyWykresuWielomianu[144].X = 1.44
	punktyWykresuWielomianu[144].Y = 0.68

	punktyWykresuWielomianu[145].X = 1.45
	punktyWykresuWielomianu[145].Y = 0.719

	punktyWykresuWielomianu[146].X = 1.46
	punktyWykresuWielomianu[146].Y = 0.76

	punktyWykresuWielomianu[147].X = 1.47
	punktyWykresuWielomianu[147].Y = 0.802

	punktyWykresuWielomianu[148].X = 1.48
	punktyWykresuWielomianu[148].Y = 0.845

	punktyWykresuWielomianu[149].X = 1.49
	punktyWykresuWielomianu[149].Y = 0.89

	punktyWykresuWielomianu[150].X = 1.5
	punktyWykresuWielomianu[150].Y = 0.937

	punktyWykresuWielomianu[151].X = 1.51
	punktyWykresuWielomianu[151].Y = 0.985

	punktyWykresuWielomianu[152].X = 1.52
	punktyWykresuWielomianu[152].Y = 1.035

	punktyWykresuWielomianu[153].X = 1.53
	punktyWykresuWielomianu[153].Y = 1.087

	punktyWykresuWielomianu[154].X = 1.54
	punktyWykresuWielomianu[154].Y = 1.14

	punktyWykresuWielomianu[155].X = 1.55
	punktyWykresuWielomianu[155].Y = 1.195

	punktyWykresuWielomianu[156].X = 1.56
	punktyWykresuWielomianu[156].Y = 1.252

	punktyWykresuWielomianu[157].X = 1.57
	punktyWykresuWielomianu[157].Y = 1.31

	punktyWykresuWielomianu[158].X = 1.58
	punktyWykresuWielomianu[158].Y = 1.371

	punktyWykresuWielomianu[159].X = 1.59
	punktyWykresuWielomianu[159].Y = 1.433

	punktyWykresuWielomianu[160].X = 1.6
	punktyWykresuWielomianu[160].Y = 1.497

	punktyWykresuWielomianu[161].X = 1.61
	punktyWykresuWielomianu[161].Y = 1.563

	punktyWykresuWielomianu[162].X = 1.62
	punktyWykresuWielomianu[162].Y = 1.631

	punktyWykresuWielomianu[163].X = 1.63
	punktyWykresuWielomianu[163].Y = 1.701

	punktyWykresuWielomianu[164].X = 1.64
	punktyWykresuWielomianu[164].Y = 1.773

	punktyWykresuWielomianu[165].X = 1.65
	punktyWykresuWielomianu[165].Y = 1.847

	punktyWykresuWielomianu[166].X = 1.66
	punktyWykresuWielomianu[166].Y = 1.923

	punktyWykresuWielomianu[167].X = 1.67
	punktyWykresuWielomianu[167].Y = 2.011

	punktyWykresuWielomianu[168].X = 1.68
	punktyWykresuWielomianu[168].Y = 2.081

	punktyWykresuWielomianu[169].X = 1.69
	punktyWykresuWielomianu[169].Y = 2.164

	punktyWykresuWielomianu[170].X = 1.7
	punktyWykresuWielomianu[170].Y = 2.249

	punktyWykresuWielomianu[171].X = 1.71
	punktyWykresuWielomianu[171].Y = 2.336

	punktyWykresuWielomianu[172].X = 1.72
	punktyWykresuWielomianu[172].Y = 2.425

	punktyWykresuWielomianu[173].X = 1.73
	punktyWykresuWielomianu[173].Y = 2.516

	punktyWykresuWielomianu[174].X = 1.74
	punktyWykresuWielomianu[174].Y = 2.61

	punktyWykresuWielomianu[175].X = 1.75
	punktyWykresuWielomianu[175].Y = 2.707

	punktyWykresuWielomianu[176].X = 1.76
	punktyWykresuWielomianu[176].Y = 2.805

	punktyWykresuWielomianu[177].X = 1.77
	punktyWykresuWielomianu[177].Y = 2.906

	punktyWykresuWielomianu[178].X = 1.78
	punktyWykresuWielomianu[178].Y = 3.01

	punktyWykresuWielomianu[179].X = 1.79
	punktyWykresuWielomianu[179].Y = 3.116

	punktyWykresuWielomianu[180].X = 1.8
	punktyWykresuWielomianu[180].Y = 3.225

	punktyWykresuWielomianu[181].X = 1.81
	punktyWykresuWielomianu[181].Y = 3.337

	punktyWykresuWielomianu[182].X = 1.82
	punktyWykresuWielomianu[182].Y = 3.451

	punktyWykresuWielomianu[183].X = 1.83
	punktyWykresuWielomianu[183].Y = 3.567

	punktyWykresuWielomianu[184].X = 1.84
	punktyWykresuWielomianu[184].Y = 3.687

	punktyWykresuWielomianu[185].X = 1.85
	punktyWykresuWielomianu[185].Y = 3.809

	punktyWykresuWielomianu[186].X = 1.86
	punktyWykresuWielomianu[186].Y = 3.934

	punktyWykresuWielomianu[187].X = 1.87
	punktyWykresuWielomianu[187].Y = 4.062

	punktyWykresuWielomianu[188].X = 1.88
	punktyWykresuWielomianu[188].Y = 4.192

	punktyWykresuWielomianu[189].X = 1.89
	punktyWykresuWielomianu[189].Y = 4.326

	punktyWykresuWielomianu[190].X = 1.9
	punktyWykresuWielomianu[190].Y = 4.463

	punktyWykresuWielomianu[191].X = 1.91
	punktyWykresuWielomianu[191].Y = 4.602

	punktyWykresuWielomianu[192].X = 1.92
	punktyWykresuWielomianu[192].Y = 4.745

	punktyWykresuWielomianu[193].X = 1.93
	punktyWykresuWielomianu[193].Y = 4.89

	punktyWykresuWielomianu[194].X = 1.94
	punktyWykresuWielomianu[194].Y = 5.039

	punktyWykresuWielomianu[195].X = 1.95
	punktyWykresuWielomianu[195].Y = 5.191

	punktyWykresuWielomianu[196].X = 1.96
	punktyWykresuWielomianu[196].Y = 5.346

	punktyWykresuWielomianu[197].X = 1.97
	punktyWykresuWielomianu[197].Y = 5.505

	punktyWykresuWielomianu[198].X = 1.98
	punktyWykresuWielomianu[198].Y = 5.666

	punktyWykresuWielomianu[199].X = 1.99
	punktyWykresuWielomianu[199].Y = 5.831

	punktyWykresuWielomianu[200].X = 2.0
	punktyWykresuWielomianu[200].Y = 6.0

	punktyWykresuWielomianu[201].X = 2.01
	punktyWykresuWielomianu[201].Y = 6.171

	punktyWykresuWielomianu[202].X = 2.02
	punktyWykresuWielomianu[202].Y = 6.346

	punktyWykresuWielomianu[203].X = 2.03
	punktyWykresuWielomianu[203].Y = 6.525

	punktyWykresuWielomianu[204].X = 2.04
	punktyWykresuWielomianu[204].Y = 6.707

	punktyWykresuWielomianu[205].X = 2.05
	punktyWykresuWielomianu[205].Y = 6.893

	punktyWykresuWielomianu[206].X = 2.06
	punktyWykresuWielomianu[206].Y = 7.082

	punktyWykresuWielomianu[207].X = 2.07
	punktyWykresuWielomianu[207].Y = 7.275

	punktyWykresuWielomianu[208].X = 2.08
	punktyWykresuWielomianu[208].Y = 7.472

	punktyWykresuWielomianu[209].X = 2.09
	punktyWykresuWielomianu[209].Y = 7.672

	punktyWykresuWielomianu[210].X = 2.1
	punktyWykresuWielomianu[210].Y = 7.877

	punktyWykresuWielomianu[211].X = 2.11
	punktyWykresuWielomianu[211].Y = 8.085

	punktyWykresuWielomianu[212].X = 2.12
	punktyWykresuWielomianu[212].Y = 8.297

	punktyWykresuWielomianu[213].X = 2.13
	punktyWykresuWielomianu[213].Y = 8.513

	punktyWykresuWielomianu[214].X = 2.14
	punktyWykresuWielomianu[214].Y = 8.732

	punktyWykresuWielomianu[215].X = 2.15
	punktyWykresuWielomianu[215].Y = 8.956

	punktyWykresuWielomianu[216].X = 2.16
	punktyWykresuWielomianu[216].Y = 9.184

	punktyWykresuWielomianu[217].X = 2.17
	punktyWykresuWielomianu[217].Y = 9.416

	punktyWykresuWielomianu[218].X = 2.18
	punktyWykresuWielomianu[218].Y = 9.652

	punktyWykresuWielomianu[219].X = 2.19
	punktyWykresuWielomianu[219].Y = 9.893

	punktyWykresuWielomianu[220].X = 2.2
	punktyWykresuWielomianu[220].Y = 10.137

	punktyWykresuWielomianu[221].X = 2.21
	punktyWykresuWielomianu[221].Y = 10.386

	punktyWykresuWielomianu[222].X = 2.22
	punktyWykresuWielomianu[222].Y = 10.639

	punktyWykresuWielomianu[223].X = 2.23
	punktyWykresuWielomianu[223].Y = 10.897

	punktyWykresuWielomianu[224].X = 2.24
	punktyWykresuWielomianu[224].Y = 11.159

	punktyWykresuWielomianu[225].X = 2.25
	punktyWykresuWielomianu[225].Y = 11.425

	punktyWykresuWielomianu[226].X = 2.26
	punktyWykresuWielomianu[226].Y = 11.696

	punktyWykresuWielomianu[227].X = 2.27
	punktyWykresuWielomianu[227].Y = 11.972

	punktyWykresuWielomianu[228].X = 2.28
	punktyWykresuWielomianu[228].Y = 12.252

	punktyWykresuWielomianu[229].X = 2.29
	punktyWykresuWielomianu[229].Y = 12.537

	punktyWykresuWielomianu[230].X = 2.3
	punktyWykresuWielomianu[230].Y = 12.827

	punktyWykresuWielomianu[231].X = 2.31
	punktyWykresuWielomianu[231].Y = 13.121

	punktyWykresuWielomianu[232].X = 2.32
	punktyWykresuWielomianu[232].Y = 13.42

	punktyWykresuWielomianu[233].X = 2.33
	punktyWykresuWielomianu[233].Y = 13.724

	punktyWykresuWielomianu[234].X = 2.34
	punktyWykresuWielomianu[234].Y = 14.033

	punktyWykresuWielomianu[235].X = 2.35
	punktyWykresuWielomianu[235].Y = 14.347

	punktyWykresuWielomianu[236].X = 2.36
	punktyWykresuWielomianu[236].Y = 14.666

	punktyWykresuWielomianu[237].X = 2.37
	punktyWykresuWielomianu[237].Y = 14.99

	punktyWykresuWielomianu[238].X = 2.38
	punktyWykresuWielomianu[238].Y = 15.319

	punktyWykresuWielomianu[239].X = 2.39
	punktyWykresuWielomianu[239].Y = 15.654

	punktyWykresuWielomianu[240].X = 2.4
	punktyWykresuWielomianu[240].Y = 15.993

	punktyWykresuWielomianu[241].X = 2.41
	punktyWykresuWielomianu[241].Y = 16.338

	punktyWykresuWielomianu[242].X = 2.42
	punktyWykresuWielomianu[242].Y = 16.688

	punktyWykresuWielomianu[243].X = 2.43
	punktyWykresuWielomianu[243].Y = 17.044

	punktyWykresuWielomianu[244].X = 2.44
	punktyWykresuWielomianu[244].Y = 17.405

	punktyWykresuWielomianu[245].X = 2.45
	punktyWykresuWielomianu[245].Y = 17.771

	punktyWykresuWielomianu[246].X = 2.46
	punktyWykresuWielomianu[246].Y = 18.143

	punktyWykresuWielomianu[247].X = 2.47
	punktyWykresuWielomianu[247].Y = 18.52

	punktyWykresuWielomianu[248].X = 2.48
	punktyWykresuWielomianu[248].Y = 18.904

	punktyWykresuWielomianu[249].X = 2.49
	punktyWykresuWielomianu[249].Y = 19.292

	punktyWykresuWielomianu[250].X = 2.5
	punktyWykresuWielomianu[250].Y = 19.687

	punktyWykresuWielomianu[251].X = 2.51
	punktyWykresuWielomianu[251].Y = 20.087

	punktyWykresuWielomianu[252].X = 2.52
	punktyWykresuWielomianu[252].Y = 20.494

	punktyWykresuWielomianu[253].X = 2.53
	punktyWykresuWielomianu[253].Y = 20.906

	punktyWykresuWielomianu[254].X = 2.54
	punktyWykresuWielomianu[254].Y = 21.324

	punktyWykresuWielomianu[255].X = 2.55
	punktyWykresuWielomianu[255].Y = 21.748

	punktyWykresuWielomianu[256].X = 2.56
	punktyWykresuWielomianu[256].Y = 22.178

	punktyWykresuWielomianu[257].X = 2.57
	punktyWykresuWielomianu[257].Y = 22.615

	punktyWykresuWielomianu[258].X = 2.58
	punktyWykresuWielomianu[258].Y = 23.057

	punktyWykresuWielomianu[259].X = 2.59
	punktyWykresuWielomianu[259].Y = 23.506

	punktyWykresuWielomianu[260].X = 2.6
	punktyWykresuWielomianu[260].Y = 23.961

	punktyWykresuWielomianu[261].X = 2.61
	punktyWykresuWielomianu[261].Y = 24.423

	punktyWykresuWielomianu[262].X = 2.62
	punktyWykresuWielomianu[262].Y = 24.89

	punktyWykresuWielomianu[263].X = 2.63
	punktyWykresuWielomianu[263].Y = 25.365

	punktyWykresuWielomianu[264].X = 2.64
	punktyWykresuWielomianu[264].Y = 25.846

	punktyWykresuWielomianu[265].X = 2.65
	punktyWykresuWielomianu[265].Y = 26.333

	punktyWykresuWielomianu[266].X = 2.66
	punktyWykresuWielomianu[266].Y = 26.827

	punktyWykresuWielomianu[267].X = 2.67
	punktyWykresuWielomianu[267].Y = 27.328

	punktyWykresuWielomianu[268].X = 2.68
	punktyWykresuWielomianu[268].Y = 27.835

	punktyWykresuWielomianu[269].X = 2.69
	punktyWykresuWielomianu[269].Y = 28.349

	punktyWykresuWielomianu[270].X = 2.7
	punktyWykresuWielomianu[270].Y = 28.871

	punktyWykresuWielomianu[271].X = 2.71
	punktyWykresuWielomianu[271].Y = 29.399

	punktyWykresuWielomianu[272].X = 2.72
	punktyWykresuWielomianu[272].Y = 29.934

	punktyWykresuWielomianu[273].X = 2.73
	punktyWykresuWielomianu[273].Y = 30.476

	punktyWykresuWielomianu[274].X = 2.74
	punktyWykresuWielomianu[274].Y = 31.025

	punktyWykresuWielomianu[275].X = 2.75
	punktyWykresuWielomianu[275].Y = 31.582

	punktyWykresuWielomianu[276].X = 2.76
	punktyWykresuWielomianu[276].Y = 32.145

	punktyWykresuWielomianu[277].X = 2.77
	punktyWykresuWielomianu[277].Y = 32.716

	punktyWykresuWielomianu[278].X = 2.78
	punktyWykresuWielomianu[278].Y = 33.294

	punktyWykresuWielomianu[279].X = 2.79
	punktyWykresuWielomianu[279].Y = 33.88

	punktyWykresuWielomianu[280].X = 2.8
	punktyWykresuWielomianu[280].Y = 34.473

	punktyWykresuWielomianu[281].X = 2.81
	punktyWykresuWielomianu[281].Y = 35.074

	punktyWykresuWielomianu[282].X = 2.82
	punktyWykresuWielomianu[282].Y = 35.682

	punktyWykresuWielomianu[283].X = 2.83
	punktyWykresuWielomianu[283].Y = 36.298

	punktyWykresuWielomianu[284].X = 2.84
	punktyWykresuWielomianu[284].Y = 36.922

	punktyWykresuWielomianu[285].X = 2.85
	punktyWykresuWielomianu[285].Y = 37.553

	punktyWykresuWielomianu[286].X = 2.86
	punktyWykresuWielomianu[286].Y = 38.192

	punktyWykresuWielomianu[287].X = 2.87
	punktyWykresuWielomianu[287].Y = 38.839

	punktyWykresuWielomianu[288].X = 2.88
	punktyWykresuWielomianu[288].Y = 39.494

	punktyWykresuWielomianu[289].X = 2.89
	punktyWykresuWielomianu[289].Y = 40.157

	punktyWykresuWielomianu[290].X = 2.9
	punktyWykresuWielomianu[290].Y = 40.829

	punktyWykresuWielomianu[291].X = 2.91
	punktyWykresuWielomianu[291].Y = 41.508

	punktyWykresuWielomianu[292].X = 2.92
	punktyWykresuWielomianu[292].Y = 42.196

	punktyWykresuWielomianu[293].X = 2.93
	punktyWykresuWielomianu[293].Y = 42.891

	punktyWykresuWielomianu[294].X = 2.94
	punktyWykresuWielomianu[294].Y = 43.596

	punktyWykresuWielomianu[295].X = 2.95
	punktyWykresuWielomianu[295].Y = 44.308

	punktyWykresuWielomianu[296].X = 2.96
	punktyWykresuWielomianu[296].Y = 45.029

	punktyWykresuWielomianu[297].X = 2.97
	punktyWykresuWielomianu[297].Y = 45.759

	punktyWykresuWielomianu[298].X = 2.98
	punktyWykresuWielomianu[298].Y = 46.497

	punktyWykresuWielomianu[299].X = 2.99
	punktyWykresuWielomianu[299].Y = 47.244

	punktyWykresuWielomianu[300].X = 3.0
	punktyWykresuWielomianu[300].Y = 48.0

	punktyWykresuWielomianu[301].X = 3.01
	punktyWykresuWielomianu[301].Y = 48.764

	punktyWykresuWielomianu[302].X = 3.02
	punktyWykresuWielomianu[302].Y = 49.537

	punktyWykresuWielomianu[303].X = 3.03
	punktyWykresuWielomianu[303].Y = 50.319

	punktyWykresuWielomianu[304].X = 3.04
	punktyWykresuWielomianu[304].Y = 51.111

	punktyWykresuWielomianu[305].X = 3.05
	punktyWykresuWielomianu[305].Y = 51.911

	punktyWykresuWielomianu[306].X = 3.06
	punktyWykresuWielomianu[306].Y = 52.72

	punktyWykresuWielomianu[307].X = 3.07
	punktyWykresuWielomianu[307].Y = 53.539

	punktyWykresuWielomianu[308].X = 3.08
	punktyWykresuWielomianu[308].Y = 54.367

	punktyWykresuWielomianu[309].X = 3.09
	punktyWykresuWielomianu[309].Y = 55.204

	punktyWykresuWielomianu[310].X = 3.1
	punktyWykresuWielomianu[310].Y = 56.051

	punktyWykresuWielomianu[311].X = 3.11
	punktyWykresuWielomianu[311].Y = 56.907

	punktyWykresuWielomianu[312].X = 3.12
	punktyWykresuWielomianu[312].Y = 57.772

	punktyWykresuWielomianu[313].X = 3.13
	punktyWykresuWielomianu[313].Y = 58.648

	punktyWykresuWielomianu[314].X = 3.14
	punktyWykresuWielomianu[314].Y = 59.533

	punktyWykresuWielomianu[315].X = 3.15
	punktyWykresuWielomianu[315].Y = 60.427

	punktyWykresuWielomianu[316].X = 3.16
	punktyWykresuWielomianu[316].Y = 61.332

	punktyWykresuWielomianu[317].X = 3.17
	punktyWykresuWielomianu[317].Y = 62.246

	punktyWykresuWielomianu[318].X = 3.18
	punktyWykresuWielomianu[318].Y = 63.17

	punktyWykresuWielomianu[319].X = 3.19
	punktyWykresuWielomianu[319].Y = 64.105

	punktyWykresuWielomianu[320].X = 3.2
	punktyWykresuWielomianu[320].Y = 65.049

	punktyWykresuWielomianu[321].X = 3.21
	punktyWykresuWielomianu[321].Y = 66.004

	punktyWykresuWielomianu[322].X = 3.22
	punktyWykresuWielomianu[322].Y = 66.969

	punktyWykresuWielomianu[323].X = 3.23
	punktyWykresuWielomianu[323].Y = 67.944

	punktyWykresuWielomianu[324].X = 3.24
	punktyWykresuWielomianu[324].Y = 68.929

	punktyWykresuWielomianu[325].X = 3.25
	punktyWykresuWielomianu[325].Y = 69.925

	punktyWykresuWielomianu[326].X = 3.26
	punktyWykresuWielomianu[326].Y = 70.932

	punktyWykresuWielomianu[327].X = 3.27
	punktyWykresuWielomianu[327].Y = 71.949

	punktyWykresuWielomianu[328].X = 3.28
	punktyWykresuWielomianu[328].Y = 72.977

	punktyWykresuWielomianu[329].X = 3.29
	punktyWykresuWielomianu[329].Y = 74.015

	punktyWykresuWielomianu[330].X = 3.3
	punktyWykresuWielomianu[330].Y = 75.065

	punktyWykresuWielomianu[331].X = 3.31
	punktyWykresuWielomianu[331].Y = 76.125

	punktyWykresuWielomianu[332].X = 3.32
	punktyWykresuWielomianu[332].Y = 77.196

	punktyWykresuWielomianu[333].X = 3.33
	punktyWykresuWielomianu[333].Y = 78.278

	punktyWykresuWielomianu[334].X = 3.34
	punktyWykresuWielomianu[334].Y = 79.372

	punktyWykresuWielomianu[335].X = 3.35
	punktyWykresuWielomianu[335].Y = 80.476

	punktyWykresuWielomianu[336].X = 3.36
	punktyWykresuWielomianu[336].Y = 81.592

	punktyWykresuWielomianu[337].X = 3.37
	punktyWykresuWielomianu[337].Y = 82.719

	punktyWykresuWielomianu[338].X = 3.38
	punktyWykresuWielomianu[338].Y = 83.858

	punktyWykresuWielomianu[339].X = 3.39
	punktyWykresuWielomianu[339].Y = 85.008

	punktyWykresuWielomianu[340].X = 3.4
	punktyWykresuWielomianu[340].Y = 86.169

	punktyWykresuWielomianu[341].X = 3.41
	punktyWykresuWielomianu[341].Y = 87.342

	punktyWykresuWielomianu[342].X = 3.42
	punktyWykresuWielomianu[342].Y = 88.527

	punktyWykresuWielomianu[343].X = 3.43
	punktyWykresuWielomianu[343].Y = 89.724

	punktyWykresuWielomianu[344].X = 3.44
	punktyWykresuWielomianu[344].Y = 90.932

	punktyWykresuWielomianu[345].X = 3.45
	punktyWykresuWielomianu[345].Y = 92.153

	punktyWykresuWielomianu[346].X = 3.46
	punktyWykresuWielomianu[346].Y = 93.385

	punktyWykresuWielomianu[347].X = 3.47
	punktyWykresuWielomianu[347].Y = 94.63

	punktyWykresuWielomianu[348].X = 3.48
	punktyWykresuWielomianu[348].Y = 95.887

	punktyWykresuWielomianu[349].X = 3.49
	punktyWykresuWielomianu[349].Y = 97.156

	punktyWykresuWielomianu[350].X = 3.5
	punktyWykresuWielomianu[350].Y = 98.437

	punktyWykresuWielomianu[351].X = 3.51
	punktyWykresuWielomianu[351].Y = 99.731

	punktyWykresuWielomianu[352].X = 3.52
	punktyWykresuWielomianu[352].Y = 101.037

	punktyWykresuWielomianu[353].X = 3.53
	punktyWykresuWielomianu[353].Y = 102.356

	punktyWykresuWielomianu[354].X = 3.54
	punktyWykresuWielomianu[354].Y = 103.687

	punktyWykresuWielomianu[355].X = 3.55
	punktyWykresuWielomianu[355].Y = 105.031

	punktyWykresuWielomianu[356].X = 3.56
	punktyWykresuWielomianu[356].Y = 106.388

	punktyWykresuWielomianu[357].X = 3.57
	punktyWykresuWielomianu[357].Y = 107.758

	punktyWykresuWielomianu[358].X = 3.58
	punktyWykresuWielomianu[358].Y = 109.141

	punktyWykresuWielomianu[359].X = 3.59
	punktyWykresuWielomianu[359].Y = 110.536

	punktyWykresuWielomianu[360].X = 3.6
	punktyWykresuWielomianu[360].Y = 111.945

	punktyWykresuWielomianu[361].X = 3.61
	punktyWykresuWielomianu[361].Y = 113.367

	punktyWykresuWielomianu[362].X = 3.62
	punktyWykresuWielomianu[362].Y = 114.803

	punktyWykresuWielomianu[363].X = 3.63
	punktyWykresuWielomianu[363].Y = 116.251

	punktyWykresuWielomianu[364].X = 3.64
	punktyWykresuWielomianu[364].Y = 117.713

	punktyWykresuWielomianu[365].X = 3.65
	punktyWykresuWielomianu[365].Y = 119.189

	punktyWykresuWielomianu[366].X = 3.66
	punktyWykresuWielomianu[366].Y = 120.678

	punktyWykresuWielomianu[367].X = 3.67
	punktyWykresuWielomianu[367].Y = 122.181

	punktyWykresuWielomianu[368].X = 3.68
	punktyWykresuWielomianu[368].Y = 123.698

	punktyWykresuWielomianu[369].X = 3.69
	punktyWykresuWielomianu[369].Y = 125.228

	punktyWykresuWielomianu[370].X = 3.7
	punktyWykresuWielomianu[370].Y = 126.773

	punktyWykresuWielomianu[371].X = 3.71
	punktyWykresuWielomianu[371].Y = 128.331

	punktyWykresuWielomianu[372].X = 3.72
	punktyWykresuWielomianu[372].Y = 129.904

	punktyWykresuWielomianu[373].X = 3.73
	punktyWykresuWielomianu[373].Y = 131.49

	punktyWykresuWielomianu[374].X = 3.74
	punktyWykresuWielomianu[374].Y = 133.091

	punktyWykresuWielomianu[375].X = 3.75
	punktyWykresuWielomianu[375].Y = 134.707

	punktyWykresuWielomianu[376].X = 3.76
	punktyWykresuWielomianu[376].Y = 136.336

	punktyWykresuWielomianu[377].X = 3.77
	punktyWykresuWielomianu[377].Y = 137.981

	punktyWykresuWielomianu[378].X = 3.78
	punktyWykresuWielomianu[378].Y = 139.639

	punktyWykresuWielomianu[379].X = 3.79
	punktyWykresuWielomianu[379].Y = 141.313

	punktyWykresuWielomianu[380].X = 3.8
	punktyWykresuWielomianu[380].Y = 143.001

	punktyWykresuWielomianu[381].X = 3.81
	punktyWykresuWielomianu[381].Y = 144.704

	punktyWykresuWielomianu[382].X = 3.82
	punktyWykresuWielomianu[382].Y = 146.422

	punktyWykresuWielomianu[383].X = 3.83
	punktyWykresuWielomianu[383].Y = 148.155

	punktyWykresuWielomianu[384].X = 3.84
	punktyWykresuWielomianu[384].Y = 149.904

	punktyWykresuWielomianu[385].X = 3.85
	punktyWykresuWielomianu[385].Y = 151.667

	punktyWykresuWielomianu[386].X = 3.86
	punktyWykresuWielomianu[386].Y = 153.446

	punktyWykresuWielomianu[387].X = 3.87
	punktyWykresuWielomianu[387].Y = 155.24

	punktyWykresuWielomianu[388].X = 3.88
	punktyWykresuWielomianu[388].Y = 157.049

	punktyWykresuWielomianu[389].X = 3.89
	punktyWykresuWielomianu[389].Y = 158.874

	punktyWykresuWielomianu[390].X = 3.9
	punktyWykresuWielomianu[390].Y = 160.715

	punktyWykresuWielomianu[391].X = 3.91
	punktyWykresuWielomianu[391].Y = 162.571

	punktyWykresuWielomianu[392].X = 3.92
	punktyWykresuWielomianu[392].Y = 164.443

	punktyWykresuWielomianu[393].X = 3.93
	punktyWykresuWielomianu[393].Y = 166.331

	punktyWykresuWielomianu[394].X = 3.94
	punktyWykresuWielomianu[394].Y = 168.235

	punktyWykresuWielomianu[395].X = 3.95
	punktyWykresuWielomianu[395].Y = 170.155

	punktyWykresuWielomianu[396].X = 3.96
	punktyWykresuWielomianu[396].Y = 172.091

	punktyWykresuWielomianu[397].X = 3.97
	punktyWykresuWielomianu[397].Y = 174.044

	punktyWykresuWielomianu[398].X = 3.98
	punktyWykresuWielomianu[398].Y = 176.013

	punktyWykresuWielomianu[399].X = 3.99
	punktyWykresuWielomianu[399].Y = 177.998

	punktyWykresuWielomianu[400].X = 4.0
	punktyWykresuWielomianu[400].Y = 180.0

	punktyWykresuWielomianu[401].X = 4.01
	punktyWykresuWielomianu[401].Y = 182.018

	punktyWykresuWielomianu[402].X = 4.02
	punktyWykresuWielomianu[402].Y = 184.053

	punktyWykresuWielomianu[403].X = 4.03
	punktyWykresuWielomianu[403].Y = 186.105

	punktyWykresuWielomianu[404].X = 4.04
	punktyWykresuWielomianu[404].Y = 188.173

	punktyWykresuWielomianu[405].X = 4.05
	punktyWykresuWielomianu[405].Y = 190.259

	punktyWykresuWielomianu[406].X = 4.06
	punktyWykresuWielomianu[406].Y = 192.362

	punktyWykresuWielomianu[407].X = 4.07
	punktyWykresuWielomianu[407].Y = 194.481

	punktyWykresuWielomianu[408].X = 4.08
	punktyWykresuWielomianu[408].Y = 196.618

	punktyWykresuWielomianu[409].X = 4.09
	punktyWykresuWielomianu[409].Y = 198.773

	punktyWykresuWielomianu[410].X = 4.1
	punktyWykresuWielomianu[410].Y = 200.945

	punktyWykresuWielomianu[411].X = 4.11
	punktyWykresuWielomianu[411].Y = 203.134

	punktyWykresuWielomianu[412].X = 4.12
	punktyWykresuWielomianu[412].Y = 205.341

	punktyWykresuWielomianu[413].X = 4.13
	punktyWykresuWielomianu[413].Y = 207.565

	punktyWykresuWielomianu[414].X = 4.14
	punktyWykresuWielomianu[414].Y = 209.808

	punktyWykresuWielomianu[415].X = 4.15
	punktyWykresuWielomianu[415].Y = 212.068

	punktyWykresuWielomianu[416].X = 4.16
	punktyWykresuWielomianu[416].Y = 214.346

	punktyWykresuWielomianu[417].X = 4.17
	punktyWykresuWielomianu[417].Y = 216.643

	punktyWykresuWielomianu[418].X = 4.18
	punktyWykresuWielomianu[418].Y = 218.957

	punktyWykresuWielomianu[419].X = 4.19
	punktyWykresuWielomianu[419].Y = 221.29

	punktyWykresuWielomianu[420].X = 4.2
	punktyWykresuWielomianu[420].Y = 223.641

	punktyWykresuWielomianu[421].X = 4.21
	punktyWykresuWielomianu[421].Y = 226.011

	punktyWykresuWielomianu[422].X = 4.22
	punktyWykresuWielomianu[422].Y = 228.399

	punktyWykresuWielomianu[423].X = 4.23
	punktyWykresuWielomianu[423].Y = 230.806

	punktyWykresuWielomianu[424].X = 4.24
	punktyWykresuWielomianu[424].Y = 233.231

	punktyWykresuWielomianu[425].X = 4.25
	punktyWykresuWielomianu[425].Y = 235.675

	punktyWykresuWielomianu[426].X = 4.26
	punktyWykresuWielomianu[426].Y = 238.675

	punktyWykresuWielomianu[427].X = 4.27
	punktyWykresuWielomianu[427].Y = 240.621

	punktyWykresuWielomianu[428].X = 4.28
	punktyWykresuWielomianu[428].Y = 243.122

	punktyWykresuWielomianu[429].X = 4.29
	punktyWykresuWielomianu[429].Y = 245.643

	punktyWykresuWielomianu[430].X = 4.3
	punktyWykresuWielomianu[430].Y = 248.183

	punktyWykresuWielomianu[431].X = 4.31
	punktyWykresuWielomianu[431].Y = 250.742

	punktyWykresuWielomianu[432].X = 4.32
	punktyWykresuWielomianu[432].Y = 253.321

	punktyWykresuWielomianu[433].X = 4.33
	punktyWykresuWielomianu[433].Y = 255.919

	punktyWykresuWielomianu[434].X = 4.34
	punktyWykresuWielomianu[434].Y = 258.537

	punktyWykresuWielomianu[435].X = 4.35
	punktyWykresuWielomianu[435].Y = 261.175

	punktyWykresuWielomianu[436].X = 4.36
	punktyWykresuWielomianu[436].Y = 263.833

	punktyWykresuWielomianu[437].X = 4.37
	punktyWykresuWielomianu[437].Y = 266.511

	punktyWykresuWielomianu[438].X = 4.38
	punktyWykresuWielomianu[438].Y = 269.209

	punktyWykresuWielomianu[439].X = 4.39
	punktyWykresuWielomianu[439].Y = 271.927

	punktyWykresuWielomianu[440].X = 4.4
	punktyWykresuWielomianu[440].Y = 274.665

	punktyWykresuWielomianu[441].X = 4.41
	punktyWykresuWielomianu[441].Y = 277.424

	punktyWykresuWielomianu[442].X = 4.42
	punktyWykresuWielomianu[442].Y = 280.203

	punktyWykresuWielomianu[443].X = 4.43
	punktyWykresuWielomianu[443].Y = 283.003

	punktyWykresuWielomianu[444].X = 4.44
	punktyWykresuWielomianu[444].Y = 285.824

	punktyWykresuWielomianu[445].X = 4.45
	punktyWykresuWielomianu[445].Y = 288.665

	punktyWykresuWielomianu[446].X = 4.46
	punktyWykresuWielomianu[446].Y = 291.527

	punktyWykresuWielomianu[447].X = 4.47
	punktyWykresuWielomianu[447].Y = 294.41

	punktyWykresuWielomianu[448].X = 4.48
	punktyWykresuWielomianu[448].Y = 297.315

	punktyWykresuWielomianu[449].X = 4.49
	punktyWykresuWielomianu[449].Y = 300.24

	punktyWykresuWielomianu[450].X = 4.5
	punktyWykresuWielomianu[450].Y = 303.187

	punktyWykresuWielomianu[451].X = 4.51
	punktyWykresuWielomianu[451].Y = 306.155

	punktyWykresuWielomianu[452].X = 4.52
	punktyWykresuWielomianu[452].Y = 309.145

	punktyWykresuWielomianu[453].X = 4.53
	punktyWykresuWielomianu[453].Y = 312.156

	punktyWykresuWielomianu[454].X = 4.54
	punktyWykresuWielomianu[454].Y = 315.189

	punktyWykresuWielomianu[455].X = 4.55
	punktyWykresuWielomianu[455].Y = 318.244

	punktyWykresuWielomianu[456].X = 4.56
	punktyWykresuWielomianu[456].Y = 321.321

	punktyWykresuWielomianu[457].X = 4.57
	punktyWykresuWielomianu[457].Y = 324.42

	punktyWykresuWielomianu[458].X = 4.58
	punktyWykresuWielomianu[458].Y = 327.541

	punktyWykresuWielomianu[459].X = 4.59
	punktyWykresuWielomianu[459].Y = 330.684

	punktyWykresuWielomianu[460].X = 4.6
	punktyWykresuWielomianu[460].Y = 333.849

	punktyWykresuWielomianu[461].X = 4.61
	punktyWykresuWielomianu[461].Y = 337.037

	punktyWykresuWielomianu[462].X = 4.62
	punktyWykresuWielomianu[462].Y = 340.247

	punktyWykresuWielomianu[463].X = 4.63
	punktyWykresuWielomianu[463].Y = 343.48

	punktyWykresuWielomianu[464].X = 4.64
	punktyWykresuWielomianu[464].Y = 346.736

	punktyWykresuWielomianu[465].X = 4.65
	punktyWykresuWielomianu[465].Y = 350.015

	punktyWykresuWielomianu[466].X = 4.66
	punktyWykresuWielomianu[466].Y = 353.317

	punktyWykresuWielomianu[467].X = 4.67
	punktyWykresuWielomianu[467].Y = 356.641

	punktyWykresuWielomianu[468].X = 4.68
	punktyWykresuWielomianu[468].Y = 359.989

	punktyWykresuWielomianu[469].X = 4.69
	punktyWykresuWielomianu[469].Y = 363.36

	punktyWykresuWielomianu[470].X = 4.7
	punktyWykresuWielomianu[470].Y = 366.755

	punktyWykresuWielomianu[471].X = 4.71
	punktyWykresuWielomianu[471].Y = 370.173

	punktyWykresuWielomianu[472].X = 4.72
	punktyWykresuWielomianu[472].Y = 373.614

	punktyWykresuWielomianu[473].X = 4.73
	punktyWykresuWielomianu[473].Y = 377.079

	punktyWykresuWielomianu[474].X = 4.74
	punktyWykresuWielomianu[474].Y = 380.569

	punktyWykresuWielomianu[475].X = 4.75
	punktyWykresuWielomianu[475].Y = 384.082

	punktyWykresuWielomianu[476].X = 4.76
	punktyWykresuWielomianu[476].Y = 387.619

	punktyWykresuWielomianu[477].X = 4.77
	punktyWykresuWielomianu[477].Y = 391.18

	punktyWykresuWielomianu[478].X = 4.78
	punktyWykresuWielomianu[478].Y = 394.765

	punktyWykresuWielomianu[479].X = 4.79
	punktyWykresuWielomianu[479].Y = 398.375

	punktyWykresuWielomianu[480].X = 4.8
	punktyWykresuWielomianu[480].Y = 402.009

	punktyWykresuWielomianu[481].X = 4.81
	punktyWykresuWielomianu[481].Y = 405.668

	punktyWykresuWielomianu[482].X = 4.82
	punktyWykresuWielomianu[482].Y = 409.351

	punktyWykresuWielomianu[483].X = 4.83
	punktyWykresuWielomianu[483].Y = 413.06

	punktyWykresuWielomianu[484].X = 4.84
	punktyWykresuWielomianu[484].Y = 416.793

	punktyWykresuWielomianu[485].X = 4.85
	punktyWykresuWielomianu[485].Y = 420.551

	punktyWykresuWielomianu[486].X = 4.86
	punktyWykresuWielomianu[486].Y = 424.334

	punktyWykresuWielomianu[487].X = 4.87
	punktyWykresuWielomianu[487].Y = 428.143

	punktyWykresuWielomianu[488].X = 4.88
	punktyWykresuWielomianu[488].Y = 431.977

	punktyWykresuWielomianu[489].X = 4.89
	punktyWykresuWielomianu[489].Y = 435.836

	punktyWykresuWielomianu[490].X = 4.9
	punktyWykresuWielomianu[490].Y = 439.721

	punktyWykresuWielomianu[491].X = 4.91
	punktyWykresuWielomianu[491].Y = 443.631

	punktyWykresuWielomianu[492].X = 4.92
	punktyWykresuWielomianu[492].Y = 447.567

	punktyWykresuWielomianu[493].X = 4.93
	punktyWykresuWielomianu[493].Y = 451.53

	punktyWykresuWielomianu[494].X = 4.94
	punktyWykresuWielomianu[494].Y = 455.518

	punktyWykresuWielomianu[495].X = 4.95
	punktyWykresuWielomianu[495].Y = 459.532

	punktyWykresuWielomianu[496].X = 4.96
	punktyWykresuWielomianu[496].Y = 463.573

	punktyWykresuWielomianu[497].X = 4.97
	punktyWykresuWielomianu[497].Y = 467.64

	punktyWykresuWielomianu[498].X = 4.98
	punktyWykresuWielomianu[498].Y = 471.733

	punktyWykresuWielomianu[499].X = 4.99
	punktyWykresuWielomianu[499].Y = 475.853

	punktyWykresuWielomianu[500].X = 5.0
	punktyWykresuWielomianu[500].Y = 480.0

	punktyWykresuWielomianu[501].X = 5.01
	punktyWykresuWielomianu[501].Y = 484.173

	punktyWykresuWielomianu[502].X = 5.02
	punktyWykresuWielomianu[502].Y = 488.373

	punktyWykresuWielomianu[503].X = 5.03
	punktyWykresuWielomianu[503].Y = 492.601

	punktyWykresuWielomianu[504].X = 5.04
	punktyWykresuWielomianu[504].Y = 496.855

	punktyWykresuWielomianu[505].X = 5.05
	punktyWykresuWielomianu[505].Y = 501.137

	punktyWykresuWielomianu[506].X = 5.06
	punktyWykresuWielomianu[506].Y = 505.446

	punktyWykresuWielomianu[507].X = 5.07
	punktyWykresuWielomianu[507].Y = 509.783

	punktyWykresuWielomianu[508].X = 5.08
	punktyWykresuWielomianu[508].Y = 514.147

	punktyWykresuWielomianu[509].X = 5.09
	punktyWykresuWielomianu[509].Y = 518.539

	punktyWykresuWielomianu[510].X = 5.1
	punktyWykresuWielomianu[510].Y = 522.959

	punktyWykresuWielomianu[511].X = 5.11
	punktyWykresuWielomianu[511].Y = 527.406

	punktyWykresuWielomianu[512].X = 5.12
	punktyWykresuWielomianu[512].Y = 531.882

	punktyWykresuWielomianu[513].X = 5.13
	punktyWykresuWielomianu[513].Y = 536.386

	punktyWykresuWielomianu[514].X = 5.14
	punktyWykresuWielomianu[514].Y = 540.918

	punktyWykresuWielomianu[515].X = 5.15
	punktyWykresuWielomianu[515].Y = 545.479

	punktyWykresuWielomianu[516].X = 5.16
	punktyWykresuWielomianu[516].Y = 550.068

	punktyWykresuWielomianu[517].X = 5.17
	punktyWykresuWielomianu[517].Y = 554.686

	punktyWykresuWielomianu[518].X = 5.18
	punktyWykresuWielomianu[518].Y = 559.333

	punktyWykresuWielomianu[519].X = 5.19
	punktyWykresuWielomianu[519].Y = 564.009

	punktyWykresuWielomianu[520].X = 5.2
	punktyWykresuWielomianu[520].Y = 568.713

	punktyWykresuWielomianu[521].X = 5.21
	punktyWykresuWielomianu[521].Y = 573.447

	punktyWykresuWielomianu[522].X = 5.22
	punktyWykresuWielomianu[522].Y = 578.21

	punktyWykresuWielomianu[523].X = 5.23
	punktyWykresuWielomianu[523].Y = 583.002

	punktyWykresuWielomianu[524].X = 5.24
	punktyWykresuWielomianu[524].Y = 587.824

	punktyWykresuWielomianu[525].X = 5.25
	punktyWykresuWielomianu[525].Y = 592.675

	punktyWykresuWielomianu[526].X = 5.26
	punktyWykresuWielomianu[526].Y = 597.556

	punktyWykresuWielomianu[527].X = 5.27
	punktyWykresuWielomianu[527].Y = 602.467

	punktyWykresuWielomianu[528].X = 5.28
	punktyWykresuWielomianu[528].Y = 607.408

	punktyWykresuWielomianu[529].X = 5.29
	punktyWykresuWielomianu[529].Y = 612.379

	punktyWykresuWielomianu[530].X = 5.3
	punktyWykresuWielomianu[530].Y = 617.381

	punktyWykresuWielomianu[531].X = 5.31
	punktyWykresuWielomianu[531].Y = 622.412

	punktyWykresuWielomianu[532].X = 5.32
	punktyWykresuWielomianu[532].Y = 627.474

	punktyWykresuWielomianu[533].X = 5.33
	punktyWykresuWielomianu[533].Y = 632.567

	punktyWykresuWielomianu[534].X = 5.34
	punktyWykresuWielomianu[534].Y = 637.69

	punktyWykresuWielomianu[535].X = 5.35
	punktyWykresuWielomianu[535].Y = 642.844

	punktyWykresuWielomianu[536].X = 5.36
	punktyWykresuWielomianu[536].Y = 648.029

	punktyWykresuWielomianu[537].X = 5.37
	punktyWykresuWielomianu[537].Y = 653.245

	punktyWykresuWielomianu[538].X = 5.38
	punktyWykresuWielomianu[538].Y = 658.493

	punktyWykresuWielomianu[539].X = 5.39
	punktyWykresuWielomianu[539].Y = 663.771

	punktyWykresuWielomianu[540].X = 5.4
	punktyWykresuWielomianu[540].Y = 669.081

	punktyWykresuWielomianu[541].X = 5.41
	punktyWykresuWielomianu[541].Y = 674.423

	punktyWykresuWielomianu[542].X = 5.42
	punktyWykresuWielomianu[542].Y = 679.796

	punktyWykresuWielomianu[543].X = 5.43
	punktyWykresuWielomianu[543].Y = 685.201

	punktyWykresuWielomianu[544].X = 5.44
	punktyWykresuWielomianu[544].Y = 690.638

	punktyWykresuWielomianu[545].X = 5.45
	punktyWykresuWielomianu[545].Y = 696.107

	punktyWykresuWielomianu[546].X = 5.46
	punktyWykresuWielomianu[546].Y = 701.608

	punktyWykresuWielomianu[547].X = 5.47
	punktyWykresuWielomianu[547].Y = 707.142

	punktyWykresuWielomianu[548].X = 5.48
	punktyWykresuWielomianu[548].Y = 712.707

	punktyWykresuWielomianu[549].X = 5.49
	punktyWykresuWielomianu[549].Y = 718.306

	punktyWykresuWielomianu[550].X = 5.5
	punktyWykresuWielomianu[550].Y = 723.937

	punktyWykresuWielomianu[551].X = 5.51
	punktyWykresuWielomianu[551].Y = 729.601

	punktyWykresuWielomianu[552].X = 5.52
	punktyWykresuWielomianu[552].Y = 735.298

	punktyWykresuWielomianu[553].X = 5.53
	punktyWykresuWielomianu[553].Y = 741.028

	punktyWykresuWielomianu[554].X = 5.54
	punktyWykresuWielomianu[554].Y = 746.791

	punktyWykresuWielomianu[555].X = 5.55
	punktyWykresuWielomianu[555].Y = 752.587

	punktyWykresuWielomianu[556].X = 5.56
	punktyWykresuWielomianu[556].Y = 758.417

	punktyWykresuWielomianu[557].X = 5.57
	punktyWykresuWielomianu[557].Y = 764.28

	punktyWykresuWielomianu[558].X = 5.58
	punktyWykresuWielomianu[558].Y = 770.177

	punktyWykresuWielomianu[559].X = 5.59
	punktyWykresuWielomianu[559].Y = 776.108

	punktyWykresuWielomianu[560].X = 5.6
	punktyWykresuWielomianu[560].Y = 782.073

	punktyWykresuWielomianu[561].X = 5.61
	punktyWykresuWielomianu[561].Y = 788.072

	punktyWykresuWielomianu[562].X = 5.62
	punktyWykresuWielomianu[562].Y = 794.105

	punktyWykresuWielomianu[563].X = 5.63
	punktyWykresuWielomianu[563].Y = 800.173

	punktyWykresuWielomianu[564].X = 5.64
	punktyWykresuWielomianu[564].Y = 806.274

	punktyWykresuWielomianu[565].X = 5.65
	punktyWykresuWielomianu[565].Y = 812.411

	punktyWykresuWielomianu[566].X = 5.66
	punktyWykresuWielomianu[566].Y = 818.582

	punktyWykresuWielomianu[567].X = 5.67
	punktyWykresuWielomianu[567].Y = 824.788

	punktyWykresuWielomianu[568].X = 5.68
	punktyWykresuWielomianu[568].Y = 831.029

	punktyWykresuWielomianu[569].X = 5.69
	punktyWykresuWielomianu[569].Y = 837.305

	punktyWykresuWielomianu[570].X = 5.7
	punktyWykresuWielomianu[570].Y = 843.617

	punktyWykresuWielomianu[571].X = 5.71
	punktyWykresuWielomianu[571].Y = 849.963

	punktyWykresuWielomianu[572].X = 5.72
	punktyWykresuWielomianu[572].Y = 856.346

	punktyWykresuWielomianu[573].X = 5.73
	punktyWykresuWielomianu[573].Y = 862.763

	punktyWykresuWielomianu[574].X = 5.74
	punktyWykresuWielomianu[574].Y = 869.217

	punktyWykresuWielomianu[575].X = 5.75
	punktyWykresuWielomianu[575].Y = 875.707

	punktyWykresuWielomianu[576].X = 5.76
	punktyWykresuWielomianu[576].Y = 882.232

	punktyWykresuWielomianu[577].X = 5.77
	punktyWykresuWielomianu[577].Y = 888.794

	punktyWykresuWielomianu[578].X = 5.78
	punktyWykresuWielomianu[578].Y = 895.392

	punktyWykresuWielomianu[579].X = 5.79
	punktyWykresuWielomianu[579].Y = 902.026

	punktyWykresuWielomianu[580].X = 5.8
	punktyWykresuWielomianu[580].Y = 908.697

	punktyWykresuWielomianu[581].X = 5.81
	punktyWykresuWielomianu[581].Y = 915.405

	punktyWykresuWielomianu[582].X = 5.82
	punktyWykresuWielomianu[582].Y = 922.149

	punktyWykresuWielomianu[583].X = 5.83
	punktyWykresuWielomianu[583].Y = 928.931

	punktyWykresuWielomianu[584].X = 5.84
	punktyWykresuWielomianu[584].Y = 935.749

	punktyWykresuWielomianu[585].X = 5.85
	punktyWykresuWielomianu[585].Y = 942.605

	punktyWykresuWielomianu[586].X = 5.86
	punktyWykresuWielomianu[586].Y = 949.498

	punktyWykresuWielomianu[587].X = 5.87
	punktyWykresuWielomianu[587].Y = 956.429

	punktyWykresuWielomianu[588].X = 5.88
	punktyWykresuWielomianu[588].Y = 963.397

	punktyWykresuWielomianu[589].X = 5.89
	punktyWykresuWielomianu[589].Y = 970.403

	punktyWykresuWielomianu[590].X = 5.9
	punktyWykresuWielomianu[590].Y = 977.447

	punktyWykresuWielomianu[591].X = 5.91
	punktyWykresuWielomianu[591].Y = 984.529

	punktyWykresuWielomianu[592].X = 5.92
	punktyWykresuWielomianu[592].Y = 991.649

	punktyWykresuWielomianu[593].X = 5.93
	punktyWykresuWielomianu[593].Y = 998.807

	punktyWykresuWielomianu[594].X = 5.94
	punktyWykresuWielomianu[594].Y = 1_006.004

	punktyWykresuWielomianu[595].X = 5.95
	punktyWykresuWielomianu[595].Y = 1_013.239

	punktyWykresuWielomianu[596].X = 5.96
	punktyWykresuWielomianu[596].Y = 1_020.513

	punktyWykresuWielomianu[597].X = 5.97
	punktyWykresuWielomianu[597].Y = 1_027.826

	punktyWykresuWielomianu[598].X = 5.98
	punktyWykresuWielomianu[598].Y = 1_035.178

	punktyWykresuWielomianu[599].X = 5.99
	punktyWykresuWielomianu[599].Y = 1_042.569

	punktyWykresuWielomianu[600].X = 6.0
	punktyWykresuWielomianu[600].Y = 1_050.0

	punktyWykresuWielomianu[601].X = 6.01
	punktyWykresuWielomianu[601].Y = 1_057.469

	punktyWykresuWielomianu[602].X = 6.02
	punktyWykresuWielomianu[602].Y = 1_064.979

	punktyWykresuWielomianu[603].X = 6.03
	punktyWykresuWielomianu[603].Y = 1_072.527

	punktyWykresuWielomianu[604].X = 6.04
	punktyWykresuWielomianu[604].Y = 1_080.116

	punktyWykresuWielomianu[605].X = 6.05
	punktyWykresuWielomianu[605].Y = 1_087.745

	punktyWykresuWielomianu[606].X = 6.06
	punktyWykresuWielomianu[606].Y = 1_095.414

	punktyWykresuWielomianu[607].X = 6.07
	punktyWykresuWielomianu[607].Y = 1_103.123

	punktyWykresuWielomianu[608].X = 6.08
	punktyWykresuWielomianu[608].Y = 1_110.872

	punktyWykresuWielomianu[609].X = 6.09
	punktyWykresuWielomianu[609].Y = 1_118.662

	punktyWykresuWielomianu[610].X = 6.1
	punktyWykresuWielomianu[610].Y = 1_126.493

	punktyWykresuWielomianu[611].X = 6.11
	punktyWykresuWielomianu[611].Y = 1_134.364

	punktyWykresuWielomianu[612].X = 6.12
	punktyWykresuWielomianu[612].Y = 1_142.276

	punktyWykresuWielomianu[613].X = 6.13
	punktyWykresuWielomianu[613].Y = 1_150.23

	punktyWykresuWielomianu[614].X = 6.14
	punktyWykresuWielomianu[614].Y = 1_158.224

	punktyWykresuWielomianu[615].X = 6.15
	punktyWykresuWielomianu[615].Y = 1_166.26

	punktyWykresuWielomianu[616].X = 6.16
	punktyWykresuWielomianu[616].Y = 1_174.338

	punktyWykresuWielomianu[617].X = 6.17
	punktyWykresuWielomianu[617].Y = 1_182.457

	punktyWykresuWielomianu[618].X = 6.18
	punktyWykresuWielomianu[618].Y = 1_190.618

	punktyWykresuWielomianu[619].X = 6.19
	punktyWykresuWielomianu[619].Y = 1_198.82

	punktyWykresuWielomianu[620].X = 6.2
	punktyWykresuWielomianu[620].Y = 1_207.065

	punktyWykresuWielomianu[621].X = 6.21
	punktyWykresuWielomianu[621].Y = 1_215.352

	punktyWykresuWielomianu[622].X = 6.22
	punktyWykresuWielomianu[622].Y = 1_223.682

	punktyWykresuWielomianu[623].X = 6.23
	punktyWykresuWielomianu[623].Y = 1_232.053

	punktyWykresuWielomianu[624].X = 6.24
	punktyWykresuWielomianu[624].Y = 1_240.468

	punktyWykresuWielomianu[625].X = 6.25
	punktyWykresuWielomianu[625].Y = 1_248.925

	punktyWykresuWielomianu[626].X = 6.26
	punktyWykresuWielomianu[626].Y = 1_257.426

	punktyWykresuWielomianu[627].X = 6.27
	punktyWykresuWielomianu[627].Y = 1_265.969

	punktyWykresuWielomianu[628].X = 6.28
	punktyWykresuWielomianu[628].Y = 1_274.555

	punktyWykresuWielomianu[629].X = 6.29
	punktyWykresuWielomianu[629].Y = 1_283.185

	punktyWykresuWielomianu[630].X = 6.3
	punktyWykresuWielomianu[630].Y = 1_291.859

	punktyWykresuWielomianu[631].X = 6.31
	punktyWykresuWielomianu[631].Y = 1_300.576

	punktyWykresuWielomianu[632].X = 6.32
	punktyWykresuWielomianu[632].Y = 1_309.336

	punktyWykresuWielomianu[633].X = 6.33
	punktyWykresuWielomianu[633].Y = 1_318.141

	punktyWykresuWielomianu[634].X = 6.34
	punktyWykresuWielomianu[634].Y = 1_326.99

	punktyWykresuWielomianu[635].X = 6.35
	punktyWykresuWielomianu[635].Y = 1_335.883

	punktyWykresuWielomianu[636].X = 6.36
	punktyWykresuWielomianu[636].Y = 1_344.821

	punktyWykresuWielomianu[637].X = 6.37
	punktyWykresuWielomianu[637].Y = 1_353.803

	punktyWykresuWielomianu[638].X = 6.38
	punktyWykresuWielomianu[638].Y = 1_362.829

	punktyWykresuWielomianu[639].X = 6.39
	punktyWykresuWielomianu[639].Y = 1_371.901

	punktyWykresuWielomianu[640].X = 6.4
	punktyWykresuWielomianu[640].Y = 1_381.017

	punktyWykresuWielomianu[641].X = 6.41
	punktyWykresuWielomianu[641].Y = 1_390.179

	punktyWykresuWielomianu[642].X = 6.42
	punktyWykresuWielomianu[642].Y = 1_399.385

	punktyWykresuWielomianu[643].X = 6.43
	punktyWykresuWielomianu[643].Y = 1_408.638

	punktyWykresuWielomianu[644].X = 6.44
	punktyWykresuWielomianu[644].Y = 1_417.935

	punktyWykresuWielomianu[645].X = 6.45
	punktyWykresuWielomianu[645].Y = 1_427.279

	punktyWykresuWielomianu[646].X = 6.46
	punktyWykresuWielomianu[646].Y = 1_436.668

	punktyWykresuWielomianu[647].X = 6.47
	punktyWykresuWielomianu[647].Y = 1_446.104

	punktyWykresuWielomianu[648].X = 6.48
	punktyWykresuWielomianu[648].Y = 1_455.585

	punktyWykresuWielomianu[649].X = 6.49
	punktyWykresuWielomianu[649].Y = 1_465.113

	punktyWykresuWielomianu[650].X = 6.5
	punktyWykresuWielomianu[650].Y = 1_474.687

	punktyWykresuWielomianu[651].X = 6.51
	punktyWykresuWielomianu[651].Y = 1_484.308

	punktyWykresuWielomianu[652].X = 6.52
	punktyWykresuWielomianu[652].Y = 1_493.975

	punktyWykresuWielomianu[653].X = 6.53
	punktyWykresuWielomianu[653].Y = 1_503.69

	punktyWykresuWielomianu[654].X = 6.54
	punktyWykresuWielomianu[654].Y = 1_513.451

	punktyWykresuWielomianu[655].X = 6.55
	punktyWykresuWielomianu[655].Y = 1_523.26

	punktyWykresuWielomianu[656].X = 6.56
	punktyWykresuWielomianu[656].Y = 1_533.116

	punktyWykresuWielomianu[657].X = 6.57
	punktyWykresuWielomianu[657].Y = 1_543.02

	punktyWykresuWielomianu[658].X = 6.58
	punktyWykresuWielomianu[658].Y = 1_552.971

	punktyWykresuWielomianu[659].X = 6.59
	punktyWykresuWielomianu[659].Y = 1_562.97

	punktyWykresuWielomianu[660].X = 6.6
	punktyWykresuWielomianu[660].Y = 1_573.017

	punktyWykresuWielomianu[661].X = 6.61
	punktyWykresuWielomianu[661].Y = 1_583.112

	punktyWykresuWielomianu[662].X = 6.62
	punktyWykresuWielomianu[662].Y = 1_593.256

	punktyWykresuWielomianu[663].X = 6.63
	punktyWykresuWielomianu[663].Y = 1_603.447

	punktyWykresuWielomianu[664].X = 6.64
	punktyWykresuWielomianu[664].Y = 1_613.688

	punktyWykresuWielomianu[665].X = 6.65
	punktyWykresuWielomianu[665].Y = 1_623.977

	punktyWykresuWielomianu[666].X = 6.66
	punktyWykresuWielomianu[666].Y = 1_634.315

	punktyWykresuWielomianu[667].X = 6.67
	punktyWykresuWielomianu[667].Y = 1_644.702

	punktyWykresuWielomianu[668].X = 6.68
	punktyWykresuWielomianu[668].Y = 1_655.138

	punktyWykresuWielomianu[669].X = 6.69
	punktyWykresuWielomianu[669].Y = 1_665.138

	punktyWykresuWielomianu[670].X = 6.7
	punktyWykresuWielomianu[670].Y = 1_676.159

	punktyWykresuWielomianu[671].X = 6.71
	punktyWykresuWielomianu[671].Y = 1_686.743

	punktyWykresuWielomianu[672].X = 6.72
	punktyWykresuWielomianu[672].Y = 1_697.378

	punktyWykresuWielomianu[673].X = 6.73
	punktyWykresuWielomianu[673].Y = 1_708.062

	punktyWykresuWielomianu[674].X = 6.74
	punktyWykresuWielomianu[674].Y = 1_718.797

	punktyWykresuWielomianu[675].X = 6.75
	punktyWykresuWielomianu[675].Y = 1_729.582

	punktyWykresuWielomianu[676].X = 6.76
	punktyWykresuWielomianu[676].Y = 1_740.417

	punktyWykresuWielomianu[677].X = 6.77
	punktyWykresuWielomianu[677].Y = 1_751.303

	punktyWykresuWielomianu[678].X = 6.78
	punktyWykresuWielomianu[678].Y = 1_762.239

	punktyWykresuWielomianu[679].X = 6.79
	punktyWykresuWielomianu[679].Y = 1_773.227

	punktyWykresuWielomianu[680].X = 6.8
	punktyWykresuWielomianu[680].Y = 1_784.265

	punktyWykresuWielomianu[681].X = 6.81
	punktyWykresuWielomianu[681].Y = 1_795.355

	punktyWykresuWielomianu[682].X = 6.82
	punktyWykresuWielomianu[682].Y = 1_806.496

	punktyWykresuWielomianu[683].X = 6.83
	punktyWykresuWielomianu[683].Y = 1_817.689

	punktyWykresuWielomianu[684].X = 6.84
	punktyWykresuWielomianu[684].Y = 1_828.933

	punktyWykresuWielomianu[685].X = 6.85
	punktyWykresuWielomianu[685].Y = 1_840.229

	punktyWykresuWielomianu[686].X = 6.86
	punktyWykresuWielomianu[686].Y = 1_851.577

	punktyWykresuWielomianu[687].X = 6.87
	punktyWykresuWielomianu[687].Y = 1_862.977

	punktyWykresuWielomianu[688].X = 6.88
	punktyWykresuWielomianu[688].Y = 1_874.43

	punktyWykresuWielomianu[689].X = 6.89
	punktyWykresuWielomianu[689].Y = 1_885.935

	punktyWykresuWielomianu[690].X = 6.9
	punktyWykresuWielomianu[690].Y = 1_897.493

	punktyWykresuWielomianu[691].X = 6.91
	punktyWykresuWielomianu[691].Y = 1_909.103

	punktyWykresuWielomianu[692].X = 6.92
	punktyWykresuWielomianu[692].Y = 1_920.767

	punktyWykresuWielomianu[693].X = 6.93
	punktyWykresuWielomianu[693].Y = 1_932.483

	punktyWykresuWielomianu[694].X = 6.94
	punktyWykresuWielomianu[694].Y = 1_944.253

	punktyWykresuWielomianu[695].X = 6.95
	punktyWykresuWielomianu[695].Y = 1_956.076

	punktyWykresuWielomianu[696].X = 6.96
	punktyWykresuWielomianu[696].Y = 1_967.953

	punktyWykresuWielomianu[697].X = 6.97
	punktyWykresuWielomianu[697].Y = 1_979.884

	punktyWykresuWielomianu[698].X = 6.98
	punktyWykresuWielomianu[698].Y = 1_991.868

	punktyWykresuWielomianu[699].X = 6.99
	punktyWykresuWielomianu[699].Y = 2_003.907

	punktyWykresuWielomianu[700].X = 7.0
	punktyWykresuWielomianu[700].Y = 2_016.0

	punktyWykresuWielomianu[701].X = 7.01
	punktyWykresuWielomianu[701].Y = 2_028.147

	punktyWykresuWielomianu[702].X = 7.02
	punktyWykresuWielomianu[702].Y = 2_040.349

	punktyWykresuWielomianu[703].X = 7.03
	punktyWykresuWielomianu[703].Y = 2_052.605

	punktyWykresuWielomianu[704].X = 7.04
	punktyWykresuWielomianu[704].Y = 2_064.916

	punktyWykresuWielomianu[705].X = 7.05
	punktyWykresuWielomianu[705].Y = 2_077.283

	punktyWykresuWielomianu[706].X = 7.06
	punktyWykresuWielomianu[706].Y = 2_089.705

	punktyWykresuWielomianu[707].X = 7.07
	punktyWykresuWielomianu[707].Y = 2_102.182

	punktyWykresuWielomianu[708].X = 7.08
	punktyWykresuWielomianu[708].Y = 2_114.714

	punktyWykresuWielomianu[709].X = 7.09
	punktyWykresuWielomianu[709].Y = 2_127.302

	punktyWykresuWielomianu[710].X = 7.1
	punktyWykresuWielomianu[710].Y = 2_139.947

	punktyWykresuWielomianu[711].X = 7.11
	punktyWykresuWielomianu[711].Y = 2_152.647

	punktyWykresuWielomianu[712].X = 7.12
	punktyWykresuWielomianu[712].Y = 2_165.403

	punktyWykresuWielomianu[713].X = 7.13
	punktyWykresuWielomianu[713].Y = 2_178.216

	punktyWykresuWielomianu[714].X = 7.14
	punktyWykresuWielomianu[714].Y = 2_191.085

	punktyWykresuWielomianu[715].X = 7.15
	punktyWykresuWielomianu[715].Y = 2_204.011

	punktyWykresuWielomianu[716].X = 7.16
	punktyWykresuWielomianu[716].Y = 2_216.994

	punktyWykresuWielomianu[717].X = 7.17
	punktyWykresuWielomianu[717].Y = 2_230.034

	punktyWykresuWielomianu[718].X = 7.18
	punktyWykresuWielomianu[718].Y = 2_243.131

	punktyWykresuWielomianu[719].X = 7.19
	punktyWykresuWielomianu[719].Y = 2_256.285

	punktyWykresuWielomianu[720].X = 7.2
	punktyWykresuWielomianu[720].Y = 2_269.497

	punktyWykresuWielomianu[721].X = 7.21
	punktyWykresuWielomianu[721].Y = 2_282.767

	punktyWykresuWielomianu[722].X = 7.22
	punktyWykresuWielomianu[722].Y = 2_296.094

	punktyWykresuWielomianu[723].X = 7.23
	punktyWykresuWielomianu[723].Y = 2_309.48

	punktyWykresuWielomianu[724].X = 7.24
	punktyWykresuWielomianu[724].Y = 2_322.923

	punktyWykresuWielomianu[725].X = 7.25
	punktyWykresuWielomianu[725].Y = 2_336.425

	punktyWykresuWielomianu[726].X = 7.26
	punktyWykresuWielomianu[726].Y = 2_349.986

	punktyWykresuWielomianu[727].X = 7.27
	punktyWykresuWielomianu[727].Y = 2_363.605

	punktyWykresuWielomianu[728].X = 7.28
	punktyWykresuWielomianu[728].Y = 2_377.283

	punktyWykresuWielomianu[729].X = 7.29
	punktyWykresuWielomianu[729].Y = 2_391.02

	punktyWykresuWielomianu[730].X = 7.3
	punktyWykresuWielomianu[730].Y = 2_404.817

	punktyWykresuWielomianu[731].X = 7.31
	punktyWykresuWielomianu[731].Y = 2_418.672

	punktyWykresuWielomianu[732].X = 7.32
	punktyWykresuWielomianu[732].Y = 2_432.588

	punktyWykresuWielomianu[733].X = 7.33
	punktyWykresuWielomianu[733].Y = 2_446.563

	punktyWykresuWielomianu[734].X = 7.34
	punktyWykresuWielomianu[734].Y = 2_460.597

	punktyWykresuWielomianu[735].X = 7.35
	punktyWykresuWielomianu[735].Y = 2_474.692

	punktyWykresuWielomianu[736].X = 7.36
	punktyWykresuWielomianu[736].Y = 2_488.847

	punktyWykresuWielomianu[737].X = 7.37
	punktyWykresuWielomianu[737].Y = 2_503.063

	punktyWykresuWielomianu[738].X = 7.38
	punktyWykresuWielomianu[738].Y = 2_517.339

	punktyWykresuWielomianu[739].X = 7.39
	punktyWykresuWielomianu[739].Y = 2_531.675

	punktyWykresuWielomianu[740].X = 7.4
	punktyWykresuWielomianu[740].Y = 2_546.073

	punktyWykresuWielomianu[741].X = 7.41
	punktyWykresuWielomianu[741].Y = 2_560.532

	punktyWykresuWielomianu[742].X = 7.42
	punktyWykresuWielomianu[742].Y = 2_575.052

	punktyWykresuWielomianu[743].X = 7.43
	punktyWykresuWielomianu[743].Y = 2_589.633

	punktyWykresuWielomianu[744].X = 7.44
	punktyWykresuWielomianu[744].Y = 2_604.276

	punktyWykresuWielomianu[745].X = 7.45
	punktyWykresuWielomianu[745].Y = 2_618.981

	punktyWykresuWielomianu[746].X = 7.46
	punktyWykresuWielomianu[746].Y = 2_633.748

	punktyWykresuWielomianu[747].X = 7.47
	punktyWykresuWielomianu[747].Y = 2_648.576

	punktyWykresuWielomianu[748].X = 7.48
	punktyWykresuWielomianu[748].Y = 2_663.467

	punktyWykresuWielomianu[749].X = 7.49
	punktyWykresuWielomianu[749].Y = 2_678.421

	punktyWykresuWielomianu[750].X = 7.5
	punktyWykresuWielomianu[750].Y = 2_693.437

	punktyWykresuWielomianu[751].X = 7.51
	punktyWykresuWielomianu[751].Y = 2_708.516

	punktyWykresuWielomianu[752].X = 7.52
	punktyWykresuWielomianu[752].Y = 2_723.658

	punktyWykresuWielomianu[753].X = 7.53
	punktyWykresuWielomianu[753].Y = 2_738.863

	punktyWykresuWielomianu[754].X = 7.54
	punktyWykresuWielomianu[754].Y = 2_754.131

	punktyWykresuWielomianu[755].X = 7.55
	punktyWykresuWielomianu[755].Y = 2_769.463

	punktyWykresuWielomianu[756].X = 7.56
	punktyWykresuWielomianu[756].Y = 2_784.859

	punktyWykresuWielomianu[757].X = 7.57
	punktyWykresuWielomianu[757].Y = 2_800.318

	punktyWykresuWielomianu[758].X = 7.58
	punktyWykresuWielomianu[758].Y = 2_815.842

	punktyWykresuWielomianu[759].X = 7.59
	punktyWykresuWielomianu[759].Y = 2_831.429

	punktyWykresuWielomianu[760].X = 7.6
	punktyWykresuWielomianu[760].Y = 2_847.081

	punktyWykresuWielomianu[761].X = 7.61
	punktyWykresuWielomianu[761].Y = 2_862.798

	punktyWykresuWielomianu[762].X = 7.62
	punktyWykresuWielomianu[762].Y = 2_878.579

	punktyWykresuWielomianu[763].X = 7.63
	punktyWykresuWielomianu[763].Y = 2_894.425

	punktyWykresuWielomianu[764].X = 7.64
	punktyWykresuWielomianu[764].Y = 2_910.336

	punktyWykresuWielomianu[765].X = 7.65
	punktyWykresuWielomianu[765].Y = 2_926.313

	punktyWykresuWielomianu[766].X = 7.66
	punktyWykresuWielomianu[766].Y = 2_942.355

	punktyWykresuWielomianu[767].X = 7.67
	punktyWykresuWielomianu[767].Y = 2_958.462

	punktyWykresuWielomianu[768].X = 7.68
	punktyWykresuWielomianu[768].Y = 2_974.636

	punktyWykresuWielomianu[769].X = 7.69
	punktyWykresuWielomianu[769].Y = 2_990.875

	punktyWykresuWielomianu[770].X = 7.7
	punktyWykresuWielomianu[770].Y = 3_007.181

	punktyWykresuWielomianu[771].X = 7.71
	punktyWykresuWielomianu[771].Y = 3_023.552

	punktyWykresuWielomianu[772].X = 7.72
	punktyWykresuWielomianu[772].Y = 3_039.991

	punktyWykresuWielomianu[773].X = 7.73
	punktyWykresuWielomianu[773].Y = 3_056.496

	punktyWykresuWielomianu[774].X = 7.74
	punktyWykresuWielomianu[774].Y = 3_073.068

	punktyWykresuWielomianu[775].X = 7.75
	punktyWykresuWielomianu[775].Y = 3_089.707

	punktyWykresuWielomianu[776].X = 7.76
	punktyWykresuWielomianu[776].Y = 3_106.413

	punktyWykresuWielomianu[777].X = 7.77
	punktyWykresuWielomianu[777].Y = 3_123.186

	punktyWykresuWielomianu[778].X = 7.78
	punktyWykresuWielomianu[778].Y = 3_140.027

	punktyWykresuWielomianu[779].X = 7.79
	punktyWykresuWielomianu[779].Y = 3_156.938

	punktyWykresuWielomianu[780].X = 7.8
	punktyWykresuWielomianu[780].Y = 3_173.913

	punktyWykresuWielomianu[781].X = 7.81
	punktyWykresuWielomianu[781].Y = 3_190.958

	punktyWykresuWielomianu[782].X = 7.82
	punktyWykresuWielomianu[782].Y = 3_208.071

	punktyWykresuWielomianu[783].X = 7.83
	punktyWykresuWielomianu[783].Y = 3_225.253

	punktyWykresuWielomianu[784].X = 7.84
	punktyWykresuWielomianu[784].Y = 3_242.504

	punktyWykresuWielomianu[785].X = 7.85
	punktyWykresuWielomianu[785].Y = 3_259.823

	punktyWykresuWielomianu[786].X = 7.86
	punktyWykresuWielomianu[786].Y = 3_277.211

	punktyWykresuWielomianu[787].X = 7.87
	punktyWykresuWielomianu[787].Y = 3_294.669

	punktyWykresuWielomianu[788].X = 7.88
	punktyWykresuWielomianu[788].Y = 3_312.196

	punktyWykresuWielomianu[789].X = 7.89
	punktyWykresuWielomianu[789].Y = 3_329.792

	punktyWykresuWielomianu[790].X = 7.9
	punktyWykresuWielomianu[790].Y = 3_347.459

	punktyWykresuWielomianu[791].X = 7.91
	punktyWykresuWielomianu[791].Y = 3_365.195

	punktyWykresuWielomianu[792].X = 7.92
	punktyWykresuWielomianu[792].Y = 3_383.001

	punktyWykresuWielomianu[793].X = 7.93
	punktyWykresuWielomianu[793].Y = 3_400.878

	punktyWykresuWielomianu[794].X = 7.94
	punktyWykresuWielomianu[794].Y = 3_418.825

	punktyWykresuWielomianu[795].X = 7.95
	punktyWykresuWielomianu[795].Y = 3_436.843

	punktyWykresuWielomianu[796].X = 7.96
	punktyWykresuWielomianu[796].Y = 3_454.932

	punktyWykresuWielomianu[797].X = 7.97
	punktyWykresuWielomianu[797].Y = 3_473.092

	punktyWykresuWielomianu[798].X = 7.98
	punktyWykresuWielomianu[798].Y = 3_491.323

	punktyWykresuWielomianu[799].X = 7.99
	punktyWykresuWielomianu[799].Y = 3_509.625

	punktyWykresuWielomianu[800].X = 8.0
	punktyWykresuWielomianu[800].Y = 3_528.0

	punktyWykresuWielomianu[801].X = 8.01
	punktyWykresuWielomianu[801].Y = 3_546.445

	punktyWykresuWielomianu[802].X = 8.02
	punktyWykresuWielomianu[802].Y = 3_564.963

	punktyWykresuWielomianu[803].X = 8.03
	punktyWykresuWielomianu[803].Y = 3_583.553

	punktyWykresuWielomianu[804].X = 8.04
	punktyWykresuWielomianu[804].Y = 3_602.216

	punktyWykresuWielomianu[805].X = 8.05
	punktyWykresuWielomianu[805].Y = 3_620.951

	punktyWykresuWielomianu[806].X = 8.06
	punktyWykresuWielomianu[806].Y = 3_639.759

	punktyWykresuWielomianu[807].X = 8.07
	punktyWykresuWielomianu[807].Y = 3_658.639

	punktyWykresuWielomianu[808].X = 8.08
	punktyWykresuWielomianu[808].Y = 3_677.593

	punktyWykresuWielomianu[809].X = 8.09
	punktyWykresuWielomianu[809].Y = 3_696.62

	punktyWykresuWielomianu[810].X = 8.1
	punktyWykresuWielomianu[810].Y = 3_715.721

	punktyWykresuWielomianu[811].X = 8.11
	punktyWykresuWielomianu[811].Y = 3_734.895

	punktyWykresuWielomianu[812].X = 8.12
	punktyWykresuWielomianu[812].Y = 3_754.143

	punktyWykresuWielomianu[813].X = 8.13
	punktyWykresuWielomianu[813].Y = 3_773.465

	punktyWykresuWielomianu[814].X = 8.14
	punktyWykresuWielomianu[814].Y = 3_792.861

	punktyWykresuWielomianu[815].X = 8.15
	punktyWykresuWielomianu[815].Y = 3_812.332

	punktyWykresuWielomianu[816].X = 8.16
	punktyWykresuWielomianu[816].Y = 3_831.878

	punktyWykresuWielomianu[817].X = 8.17
	punktyWykresuWielomianu[817].Y = 3_851.498

	punktyWykresuWielomianu[818].X = 8.18
	punktyWykresuWielomianu[818].Y = 3_871.193

	punktyWykresuWielomianu[819].X = 8.19
	punktyWykresuWielomianu[819].Y = 3_890.963

	punktyWykresuWielomianu[820].X = 8.2
	punktyWykresuWielomianu[820].Y = 3_910.809

	punktyWykresuWielomianu[821].X = 8.21
	punktyWykresuWielomianu[821].Y = 0.0

	punktyWykresuWielomianu[822].X = 8.22
	punktyWykresuWielomianu[822].Y = 0.0

	punktyWykresuWielomianu[823].X = 8.23
	punktyWykresuWielomianu[823].Y = 0.0

	punktyWykresuWielomianu[824].X = 8.24
	punktyWykresuWielomianu[824].Y = 0.0

	punktyWykresuWielomianu[825].X = 8.25
	punktyWykresuWielomianu[825].Y = 0.0

	punktyWykresuWielomianu[826].X = 8.26
	punktyWykresuWielomianu[826].Y = 0.0

	punktyWykresuWielomianu[827].X = 8.27
	punktyWykresuWielomianu[827].Y = 0.0

	punktyWykresuWielomianu[828].X = 8.28
	punktyWykresuWielomianu[828].Y = 0.0

	punktyWykresuWielomianu[829].X = 8.29
	punktyWykresuWielomianu[829].Y = 0.0

	punktyWykresuWielomianu[830].X = 8.3
	punktyWykresuWielomianu[830].Y = 0.0

	punktyWykresuWielomianu[831].X = 8.31
	punktyWykresuWielomianu[831].Y = 0.0

	punktyWykresuWielomianu[832].X = 8.32
	punktyWykresuWielomianu[832].Y = 0.0

	punktyWykresuWielomianu[833].X = 8.33
	punktyWykresuWielomianu[833].Y = 0.0

	punktyWykresuWielomianu[834].X = 8.34
	punktyWykresuWielomianu[834].Y = 0.0

	punktyWykresuWielomianu[835].X = 8.35
	punktyWykresuWielomianu[835].Y = 0.0

	punktyWykresuWielomianu[836].X = 8.36
	punktyWykresuWielomianu[836].Y = 0.0

	punktyWykresuWielomianu[837].X = 8.37
	punktyWykresuWielomianu[837].Y = 0.0

	punktyWykresuWielomianu[838].X = 8.38
	punktyWykresuWielomianu[838].Y = 0.0

	punktyWykresuWielomianu[839].X = 8.39
	punktyWykresuWielomianu[839].Y = 0.0

	punktyWykresuWielomianu[840].X = 8.4
	punktyWykresuWielomianu[840].Y = 0.0

	punktyWykresuWielomianu[841].X = 8.41
	punktyWykresuWielomianu[841].Y = 0.0

	punktyWykresuWielomianu[842].X = 8.42
	punktyWykresuWielomianu[842].Y = 0.0

	punktyWykresuWielomianu[843].X = 8.43
	punktyWykresuWielomianu[843].Y = 0.0

	punktyWykresuWielomianu[844].X = 8.44
	punktyWykresuWielomianu[844].Y = 0.0

	punktyWykresuWielomianu[845].X = 8.45
	punktyWykresuWielomianu[845].Y = 0.0

	punktyWykresuWielomianu[846].X = 8.46
	punktyWykresuWielomianu[846].Y = 0.0

	punktyWykresuWielomianu[847].X = 8.47
	punktyWykresuWielomianu[847].Y = 0.0

	punktyWykresuWielomianu[848].X = 8.48
	punktyWykresuWielomianu[848].Y = 0.0

	punktyWykresuWielomianu[849].X = 8.49
	punktyWykresuWielomianu[849].Y = 0.0

	punktyWykresuWielomianu[850].X = 8.5
	punktyWykresuWielomianu[850].Y = 0.0

	punktyWykresuWielomianu[851].X = 8.51
	punktyWykresuWielomianu[851].Y = 0.0

	punktyWykresuWielomianu[852].X = 8.52
	punktyWykresuWielomianu[852].Y = 0.0

	punktyWykresuWielomianu[853].X = 8.53
	punktyWykresuWielomianu[853].Y = 0.0

	punktyWykresuWielomianu[854].X = 8.54
	punktyWykresuWielomianu[854].Y = 0.0

	punktyWykresuWielomianu[855].X = 8.55
	punktyWykresuWielomianu[855].Y = 0.0

	punktyWykresuWielomianu[856].X = 8.56
	punktyWykresuWielomianu[856].Y = 0.0

	punktyWykresuWielomianu[857].X = 8.57
	punktyWykresuWielomianu[857].Y = 0.0

	punktyWykresuWielomianu[858].X = 8.58
	punktyWykresuWielomianu[858].Y = 0.0

	punktyWykresuWielomianu[859].X = 8.59
	punktyWykresuWielomianu[859].Y = 0.0

	punktyWykresuWielomianu[860].X = 8.6
	punktyWykresuWielomianu[860].Y = 0.0

	punktyWykresuWielomianu[861].X = 8.61
	punktyWykresuWielomianu[861].Y = 0.0

	punktyWykresuWielomianu[862].X = 8.62
	punktyWykresuWielomianu[862].Y = 0.0

	punktyWykresuWielomianu[863].X = 8.63
	punktyWykresuWielomianu[863].Y = 0.0

	punktyWykresuWielomianu[864].X = 8.64
	punktyWykresuWielomianu[864].Y = 0.0

	punktyWykresuWielomianu[865].X = 8.65
	punktyWykresuWielomianu[865].Y = 0.0

	punktyWykresuWielomianu[866].X = 8.66
	punktyWykresuWielomianu[866].Y = 0.0

	punktyWykresuWielomianu[867].X = 8.67
	punktyWykresuWielomianu[867].Y = 0.0

	punktyWykresuWielomianu[868].X = 8.68
	punktyWykresuWielomianu[868].Y = 0.0

	punktyWykresuWielomianu[869].X = 8.69
	punktyWykresuWielomianu[869].Y = 0.0

	punktyWykresuWielomianu[870].X = 8.7
	punktyWykresuWielomianu[870].Y = 5_003.483

	punktyWykresuWielomianu[871].X = 8.71
	punktyWykresuWielomianu[871].Y = 5_027.431

	punktyWykresuWielomianu[872].X = 8.72
	punktyWykresuWielomianu[872].Y = 5_051.465

	punktyWykresuWielomianu[873].X = 8.73
	punktyWykresuWielomianu[873].Y = 5_075.584

	punktyWykresuWielomianu[874].X = 8.74
	punktyWykresuWielomianu[874].Y = 5_099.79

	punktyWykresuWielomianu[875].X = 8.75
	punktyWykresuWielomianu[875].Y = 5_124.082

	punktyWykresuWielomianu[876].X = 8.76
	punktyWykresuWielomianu[876].Y = 5_148.46

	punktyWykresuWielomianu[877].X = 8.77
	punktyWykresuWielomianu[877].Y = 5_172.925

	punktyWykresuWielomianu[878].X = 8.78
	punktyWykresuWielomianu[878].Y = 5_197.476

	punktyWykresuWielomianu[879].X = 8.79
	punktyWykresuWielomianu[879].Y = 5_222.115

	punktyWykresuWielomianu[880].X = 8.8
	punktyWykresuWielomianu[880].Y = 5_246.841

	punktyWykresuWielomianu[881].X = 8.81
	punktyWykresuWielomianu[881].Y = 5_271.655

	punktyWykresuWielomianu[882].X = 8.82
	punktyWykresuWielomianu[882].Y = 5_296.556

	punktyWykresuWielomianu[883].X = 8.83
	punktyWykresuWielomianu[883].Y = 5_321.545

	punktyWykresuWielomianu[884].X = 8.84
	punktyWykresuWielomianu[884].Y = 5_346.622

	punktyWykresuWielomianu[885].X = 8.85
	punktyWykresuWielomianu[885].Y = 5_371.787

	punktyWykresuWielomianu[886].X = 8.86
	punktyWykresuWielomianu[886].Y = 5_397.041

	punktyWykresuWielomianu[887].X = 8.87
	punktyWykresuWielomianu[887].Y = 5_422.383

	punktyWykresuWielomianu[888].X = 8.88
	punktyWykresuWielomianu[888].Y = 5_447.814

	punktyWykresuWielomianu[889].X = 8.89
	punktyWykresuWielomianu[889].Y = 5_473.335

	punktyWykresuWielomianu[890].X = 8.9
	punktyWykresuWielomianu[890].Y = 5_498.945

	punktyWykresuWielomianu[891].X = 8.91
	punktyWykresuWielomianu[891].Y = 5_524.644

	punktyWykresuWielomianu[892].X = 8.92
	punktyWykresuWielomianu[892].Y = 5_550.433

	punktyWykresuWielomianu[893].X = 8.93
	punktyWykresuWielomianu[893].Y = 5_576.312

	punktyWykresuWielomianu[894].X = 8.94
	punktyWykresuWielomianu[894].Y = 5_602.281

	punktyWykresuWielomianu[895].X = 8.95
	punktyWykresuWielomianu[895].Y = 5_628.34

	punktyWykresuWielomianu[896].X = 8.96
	punktyWykresuWielomianu[896].Y = 5_654.49

	punktyWykresuWielomianu[897].X = 8.97
	punktyWykresuWielomianu[897].Y = 5_680.731

	punktyWykresuWielomianu[898].X = 8.98
	punktyWykresuWielomianu[898].Y = 5_707.062

	punktyWykresuWielomianu[899].X = 8.99
	punktyWykresuWielomianu[899].Y = 5_733.485

	punktyWykresuWielomianu[900].X = 9.0
	punktyWykresuWielomianu[900].Y = 5_760.0





































































	wykresWielomianu := plot.New()

	wykresWielomianu.Title.Text = "Wykres funkcji f(x) = x^4 - x^3 - x^2 + x"

	wykresWielomianu.X.Label.Text = "x"
	wykresWielomianu.Y.Label.Text = "y"

	liniaWykresu, err := plotter.NewLine(punktyWykresuWielomianu)

	if err != nil {
		panic(err)
	}

	liniaWykresu.LineStyle.Width = vg.Points(0.1)
	liniaWykresu.Color = color.RGBA{R: 200, G: 100, B: 100}

	wykresWielomianu.Add(liniaWykresu)
	wykresWielomianu.Legend.Add("f(x)", liniaWykresu)

	if err := wykresWielomianu.Save(10*vg.Inch, 10*vg.Inch,
		"Wykres-wielomianu-02.png"); err != nil {

		panic(err)
	}
}
