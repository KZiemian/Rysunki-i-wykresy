package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres funkcji f(x) = -(2/3)x^3 + x - 5cos(x) + 5
	// Jest to funkcja pierwotna funkcji g(x) = -2x^2 + 1 + 5sin(x),
	// spełniająca warunek f(0) = 0.

	punktyWykresuFunkcji := make(plotter.XYs, 1001)

	punktyWykresuFunkcji[0].X = 0.0
	punktyWykresuFunkcji[0].Y = 0.0

	punktyWykresuFunkcji[1].X = 0.01
	punktyWykresuFunkcji[1].Y = 0.01

	punktyWykresuFunkcji[2].X = 0.02
	punktyWykresuFunkcji[2].Y = 0.021

	punktyWykresuFunkcji[3].X = 0.03
	punktyWykresuFunkcji[3].Y = 0.032

	punktyWykresuFunkcji[4].X = 0.04
	punktyWykresuFunkcji[4].Y = 0.044

	punktyWykresuFunkcji[5].X = 0.05
	punktyWykresuFunkcji[5].Y = 0.056

	punktyWykresuFunkcji[6].X = 0.06
	punktyWykresuFunkcji[6].Y = 0.068

	punktyWykresuFunkcji[7].X = 0.07
	punktyWykresuFunkcji[7].Y = 0.082

	punktyWykresuFunkcji[8].X = 0.08
	punktyWykresuFunkcji[8].Y = 0.095

	punktyWykresuFunkcji[9].X = 0.09
	punktyWykresuFunkcji[9].Y = 0.109

	punktyWykresuFunkcji[10].X = 0.1
	punktyWykresuFunkcji[10].Y = 0.124

	punktyWykresuFunkcji[11].X = 0.11
	punktyWykresuFunkcji[11].Y = 0.139

	punktyWykresuFunkcji[12].X = 0.12
	punktyWykresuFunkcji[12].Y = 0.154

	punktyWykresuFunkcji[13].X = 0.13
	punktyWykresuFunkcji[13].Y = 0.17

	punktyWykresuFunkcji[14].X = 0.14
	punktyWykresuFunkcji[14].Y = 0.187

	punktyWykresuFunkcji[15].X = 0.15
	punktyWykresuFunkcji[15].Y = 0.203

	punktyWykresuFunkcji[16].X = 0.16
	punktyWykresuFunkcji[16].Y = 0.221

	punktyWykresuFunkcji[17].X = 0.17
	punktyWykresuFunkcji[17].Y = 0.238

	punktyWykresuFunkcji[18].X = 0.18
	punktyWykresuFunkcji[18].Y = 0.256

	punktyWykresuFunkcji[19].X = 0.19
	punktyWykresuFunkcji[19].Y = 0.275

	punktyWykresuFunkcji[20].X = 0.2
	punktyWykresuFunkcji[20].Y = 0.294

	punktyWykresuFunkcji[21].X = 0.21
	punktyWykresuFunkcji[21].Y = 0.313

	punktyWykresuFunkcji[22].X = 0.22
	punktyWykresuFunkcji[22].Y = 0.333

	punktyWykresuFunkcji[23].X = 0.23
	punktyWykresuFunkcji[23].Y = 0.353

	punktyWykresuFunkcji[24].X = 0.24
	punktyWykresuFunkcji[24].Y = 0.374

	punktyWykresuFunkcji[25].X = 0.25
	punktyWykresuFunkcji[25].Y = 0.395

	punktyWykresuFunkcji[26].X = 0.26
	punktyWykresuFunkcji[26].Y = 0.416

	punktyWykresuFunkcji[27].X = 0.27
	punktyWykresuFunkcji[27].Y = 0.438

	punktyWykresuFunkcji[28].X = 0.28
	punktyWykresuFunkcji[28].Y = 0.46

	punktyWykresuFunkcji[29].X = 0.29
	punktyWykresuFunkcji[29].Y = 0.482

	punktyWykresuFunkcji[30].X = 0.3
	punktyWykresuFunkcji[30].Y = 0.505

	punktyWykresuFunkcji[31].X = 0.31
	punktyWykresuFunkcji[31].Y = 0.528

	punktyWykresuFunkcji[32].X = 0.32
	punktyWykresuFunkcji[32].Y = 0.552

	punktyWykresuFunkcji[33].X = 0.33
	punktyWykresuFunkcji[33].Y = 0.575

	punktyWykresuFunkcji[34].X = 0.34
	punktyWykresuFunkcji[34].Y = 0.6

	punktyWykresuFunkcji[35].X = 0.35
	punktyWykresuFunkcji[35].Y = 0.624

	punktyWykresuFunkcji[36].X = 0.36
	punktyWykresuFunkcji[36].Y = 0.649

	punktyWykresuFunkcji[37].X = 0.37
	punktyWykresuFunkcji[37].Y = 0.674

	punktyWykresuFunkcji[38].X = 0.38
	punktyWykresuFunkcji[38].Y = 0.7

	punktyWykresuFunkcji[39].X = 0.39
	punktyWykresuFunkcji[39].Y = 0.725

	punktyWykresuFunkcji[40].X = 0.40
	punktyWykresuFunkcji[40].Y = 0.752

	punktyWykresuFunkcji[41].X = 0.41
	punktyWykresuFunkcji[41].Y = 0.778

	punktyWykresuFunkcji[42].X = 0.42
	punktyWykresuFunkcji[42].Y = 0.805

	punktyWykresuFunkcji[43].X = 0.43
	punktyWykresuFunkcji[43].Y = 0.832

	punktyWykresuFunkcji[44].X = 0.44
	punktyWykresuFunkcji[44].Y = 0.859

	punktyWykresuFunkcji[45].X = 0.45
	punktyWykresuFunkcji[45].Y = 0.887

	punktyWykresuFunkcji[46].X = 0.46
	punktyWykresuFunkcji[46].Y = 0.914

	punktyWykresuFunkcji[47].X = 0.47
	punktyWykresuFunkcji[47].Y = 0.942

	punktyWykresuFunkcji[48].X = 0.48
	punktyWykresuFunkcji[48].Y = 0.971

	punktyWykresuFunkcji[49].X = 0.49
	punktyWykresuFunkcji[49].Y = 0.999

	punktyWykresuFunkcji[50].X = 0.5
	punktyWykresuFunkcji[50].Y = 1.028

	punktyWykresuFunkcji[51].X = 0.51
	punktyWykresuFunkcji[51].Y = 1.057

	punktyWykresuFunkcji[52].X = 0.52
	punktyWykresuFunkcji[52].Y = 1.087

	punktyWykresuFunkcji[53].X = 0.53
	punktyWykresuFunkcji[53].Y = 1.116

	punktyWykresuFunkcji[54].X = 0.54
	punktyWykresuFunkcji[54].Y = 1.146

	punktyWykresuFunkcji[55].X = 0.55
	punktyWykresuFunkcji[55].Y = 1.176

	punktyWykresuFunkcji[56].X = 0.56
	punktyWykresuFunkcji[56].Y = 1.206

	punktyWykresuFunkcji[57].X = 0.57
	punktyWykresuFunkcji[57].Y = 1.237

	punktyWykresuFunkcji[58].X = 0.58
	punktyWykresuFunkcji[58].Y = 1.267

	punktyWykresuFunkcji[59].X = 0.59
	punktyWykresuFunkcji[59].Y = 1.298

	punktyWykresuFunkcji[60].X = 0.6
	punktyWykresuFunkcji[60].Y = 1.329

	punktyWykresuFunkcji[61].X = 0.61
	punktyWykresuFunkcji[61].Y = 1.36

	punktyWykresuFunkcji[62].X = 0.62
	punktyWykresuFunkcji[62].Y = 1.391

	punktyWykresuFunkcji[63].X = 0.63
	punktyWykresuFunkcji[63].Y = 1.423

	punktyWykresuFunkcji[64].X = 0.64
	punktyWykresuFunkcji[64].Y = 1.454

	punktyWykresuFunkcji[65].X = 0.65
	punktyWykresuFunkcji[65].Y = 1.486

	punktyWykresuFunkcji[66].X = 0.66
	punktyWykresuFunkcji[66].Y = 1.518

	punktyWykresuFunkcji[67].X = 0.67
	punktyWykresuFunkcji[67].Y = 1.55

	punktyWykresuFunkcji[68].X = 0.68
	punktyWykresuFunkcji[68].Y = 1.582

	punktyWykresuFunkcji[69].X = 0.69
	punktyWykresuFunkcji[69].Y = 1.614

	punktyWykresuFunkcji[70].X = 0.7
	punktyWykresuFunkcji[70].Y = 1.647

	punktyWykresuFunkcji[71].X = 0.71
	punktyWykresuFunkcji[71].Y = 1.679

	punktyWykresuFunkcji[72].X = 0.72
	punktyWykresuFunkcji[72].Y = 1.712

	punktyWykresuFunkcji[73].X = 0.73
	punktyWykresuFunkcji[73].Y = 1.744

	punktyWykresuFunkcji[74].X = 0.74
	punktyWykresuFunkcji[74].Y = 1.777

	punktyWykresuFunkcji[75].X = 0.75
	punktyWykresuFunkcji[75].Y = 1.81

	punktyWykresuFunkcji[76].X = 0.76
	punktyWykresuFunkcji[76].Y = 1.843

	punktyWykresuFunkcji[77].X = 0.77
	punktyWykresuFunkcji[77].Y = 1.876

	punktyWykresuFunkcji[78].X = 0.78
	punktyWykresuFunkcji[78].Y = 1.909

	punktyWykresuFunkcji[79].X = 0.79
	punktyWykresuFunkcji[79].Y = 1.942

	punktyWykresuFunkcji[80].X = 0.80
	punktyWykresuFunkcji[80].Y = 1.975

	punktyWykresuFunkcji[81].X = 0.81
	punktyWykresuFunkcji[81].Y = 2.008

	punktyWykresuFunkcji[82].X = 0.82
	punktyWykresuFunkcji[82].Y = 2.041

	punktyWykresuFunkcji[83].X = 0.83
	punktyWykresuFunkcji[83].Y = 2.074

	punktyWykresuFunkcji[84].X = 0.84
	punktyWykresuFunkcji[84].Y = 2.107

	punktyWykresuFunkcji[85].X = 0.85
	punktyWykresuFunkcji[85].Y = 2.14

	punktyWykresuFunkcji[86].X = 0.86
	punktyWykresuFunkcji[86].Y = 2.173

	punktyWykresuFunkcji[87].X = 0.87
	punktyWykresuFunkcji[87].Y = 2.206

	punktyWykresuFunkcji[88].X = 0.88
	punktyWykresuFunkcji[88].Y = 2.239

	punktyWykresuFunkcji[89].X = 0.89
	punktyWykresuFunkcji[89].Y = 2.273

	punktyWykresuFunkcji[90].X = 0.9
	punktyWykresuFunkcji[90].Y = 2.306

	punktyWykresuFunkcji[91].X = 0.91
	punktyWykresuFunkcji[91].Y = 2.338

	punktyWykresuFunkcji[92].X = 0.92
	punktyWykresuFunkcji[92].Y = 2.371

	punktyWykresuFunkcji[93].X = 0.93
	punktyWykresuFunkcji[93].Y = 2.404

	punktyWykresuFunkcji[94].X = 0.94
	punktyWykresuFunkcji[94].Y = 2.437

	punktyWykresuFunkcji[95].X = 0.95
	punktyWykresuFunkcji[95].Y = 2.47

	punktyWykresuFunkcji[96].X = 0.96
	punktyWykresuFunkcji[96].Y = 2.502

	punktyWykresuFunkcji[97].X = 0.97
	punktyWykresuFunkcji[97].Y = 2.535

	punktyWykresuFunkcji[98].X = 0.98
	punktyWykresuFunkcji[98].Y = 2.567

	punktyWykresuFunkcji[99].X = 0.99
	punktyWykresuFunkcji[99].Y = 2.599

	punktyWykresuFunkcji[100].X = 1.0
	punktyWykresuFunkcji[100].Y = 2.631

	punktyWykresuFunkcji[101].X = 1.01
	punktyWykresuFunkcji[101].Y = 2.663

	punktyWykresuFunkcji[102].X = 1.02
	punktyWykresuFunkcji[102].Y = 2.695

	punktyWykresuFunkcji[103].X = 1.03
	punktyWykresuFunkcji[103].Y = 2.727

	punktyWykresuFunkcji[104].X = 1.04
	punktyWykresuFunkcji[104].Y = 2.759

	punktyWykresuFunkcji[105].X = 1.05
	punktyWykresuFunkcji[105].Y = 2.79

	punktyWykresuFunkcji[106].X = 1.06
	punktyWykresuFunkcji[106].Y = 2.821

	punktyWykresuFunkcji[107].X = 1.07
	punktyWykresuFunkcji[107].Y = 2.852

	punktyWykresuFunkcji[108].X = 1.08
	punktyWykresuFunkcji[108].Y = 2.883

	punktyWykresuFunkcji[109].X = 1.09
	punktyWykresuFunkcji[109].Y = 2.914

	punktyWykresuFunkcji[110].X = 1.1
	punktyWykresuFunkcji[110].Y = 2.944

	punktyWykresuFunkcji[111].X = 1.11
	punktyWykresuFunkcji[111].Y = 2.974

	punktyWykresuFunkcji[112].X = 1.12
	punktyWykresuFunkcji[112].Y = 3.005

	punktyWykresuFunkcji[113].X = 1.13
	punktyWykresuFunkcji[113].Y = 3.034

	punktyWykresuFunkcji[114].X = 1.14
	punktyWykresuFunkcji[114].Y = 3.064

	punktyWykresuFunkcji[115].X = 1.15
	punktyWykresuFunkcji[115].Y = 3.093

	punktyWykresuFunkcji[116].X = 1.16
	punktyWykresuFunkcji[116].Y = 3.122

	punktyWykresuFunkcji[117].X = 1.17
	punktyWykresuFunkcji[117].Y = 3.151

	punktyWykresuFunkcji[118].X = 1.18
	punktyWykresuFunkcji[118].Y = 3.18

	punktyWykresuFunkcji[119].X = 1.19
	punktyWykresuFunkcji[119].Y = 3.208

	punktyWykresuFunkcji[120].X = 1.2
	punktyWykresuFunkcji[120].Y = 3.236

	punktyWykresuFunkcji[121].X = 1.21
	punktyWykresuFunkcji[121].Y = 3.263

	punktyWykresuFunkcji[122].X = 1.22
	punktyWykresuFunkcji[122].Y = 3.291

	punktyWykresuFunkcji[123].X = 1.23
	punktyWykresuFunkcji[123].Y = 3.318

	punktyWykresuFunkcji[124].X = 1.24
	punktyWykresuFunkcji[124].Y = 3.344

	punktyWykresuFunkcji[125].X = 1.25
	punktyWykresuFunkcji[125].Y = 3.371

	punktyWykresuFunkcji[126].X = 1.26
	punktyWykresuFunkcji[126].Y = 3.397

	punktyWykresuFunkcji[127].X = 1.27
	punktyWykresuFunkcji[127].Y = 3.423

	punktyWykresuFunkcji[128].X = 1.28
	punktyWykresuFunkcji[128].Y = 3.448

	punktyWykresuFunkcji[129].X = 1.29
	punktyWykresuFunkcji[129].Y = 3.473

	punktyWykresuFunkcji[130].X = 1.3
	punktyWykresuFunkcji[130].Y = 3.497

	punktyWykresuFunkcji[131].X = 1.31
	punktyWykresuFunkcji[131].Y = 3.522

	punktyWykresuFunkcji[132].X = 1.32
	punktyWykresuFunkcji[132].Y = 3.545

	punktyWykresuFunkcji[133].X = 1.33
	punktyWykresuFunkcji[133].Y = 3.569

	punktyWykresuFunkcji[134].X = 1.34
	punktyWykresuFunkcji[134].Y = 3.592

	punktyWykresuFunkcji[135].X = 1.35
	punktyWykresuFunkcji[135].Y = 3.614

	punktyWykresuFunkcji[136].X = 1.36
	punktyWykresuFunkcji[136].Y = 3.636

	punktyWykresuFunkcji[137].X = 1.37
	punktyWykresuFunkcji[137].Y = 3.658

	punktyWykresuFunkcji[138].X = 1.38
	punktyWykresuFunkcji[138].Y = 3.679

	punktyWykresuFunkcji[139].X = 1.39
	punktyWykresuFunkcji[139].Y = 3.7

	punktyWykresuFunkcji[140].X = 1.4
	punktyWykresuFunkcji[140].Y = 3.72

	punktyWykresuFunkcji[141].X = 1.41
	punktyWykresuFunkcji[141].Y = 3.74

	punktyWykresuFunkcji[142].X = 1.42
	punktyWykresuFunkcji[142].Y = 3.76

	punktyWykresuFunkcji[143].X = 1.43
	punktyWykresuFunkcji[143].Y = 3.778

	punktyWykresuFunkcji[144].X = 1.44
	punktyWykresuFunkcji[144].Y = 3.797

	punktyWykresuFunkcji[145].X = 1.45
	punktyWykresuFunkcji[145].Y = 3.815

	punktyWykresuFunkcji[146].X = 1.46
	punktyWykresuFunkcji[146].Y = 3.832

	punktyWykresuFunkcji[147].X = 1.47
	punktyWykresuFunkcji[147].Y = 3.849

	punktyWykresuFunkcji[148].X = 1.48
	punktyWykresuFunkcji[148].Y = 3.865

	punktyWykresuFunkcji[149].X = 1.49
	punktyWykresuFunkcji[149].Y = 3.881

	punktyWykresuFunkcji[150].X = 1.5
	punktyWykresuFunkcji[150].Y = 3.896

	punktyWykresuFunkcji[151].X = 1.51
	punktyWykresuFunkcji[151].Y = 3.91

	punktyWykresuFunkcji[152].X = 1.52
	punktyWykresuFunkcji[152].Y = 3.924

	punktyWykresuFunkcji[153].X = 1.53
	punktyWykresuFunkcji[153].Y = 3.938

	punktyWykresuFunkcji[154].X = 1.54
	punktyWykresuFunkcji[154].Y = 3.951

	punktyWykresuFunkcji[155].X = 1.55
	punktyWykresuFunkcji[155].Y = 3.963

	punktyWykresuFunkcji[156].X = 1.56
	punktyWykresuFunkcji[156].Y = 3.975

	punktyWykresuFunkcji[157].X = 1.57
	punktyWykresuFunkcji[157].Y = 3.986

	punktyWykresuFunkcji[158].X = 1.58
	punktyWykresuFunkcji[158].Y = 3.996

	punktyWykresuFunkcji[159].X = 1.59
	punktyWykresuFunkcji[159].Y = 4.006

	punktyWykresuFunkcji[160].X = 1.6
	punktyWykresuFunkcji[160].Y = 4.015

	punktyWykresuFunkcji[161].X = 1.61
	punktyWykresuFunkcji[161].Y = 4.023

	punktyWykresuFunkcji[162].X = 1.62
	punktyWykresuFunkcji[162].Y = 4.031

	punktyWykresuFunkcji[163].X = 1.63
	punktyWykresuFunkcji[163].Y = 4.038

	punktyWykresuFunkcji[164].X = 1.64
	punktyWykresuFunkcji[164].Y = 4.045

	punktyWykresuFunkcji[165].X = 1.65
	punktyWykresuFunkcji[165].Y = 4.05

	punktyWykresuFunkcji[166].X = 1.66
	punktyWykresuFunkcji[166].Y = 4.055

	punktyWykresuFunkcji[167].X = 1.67
	punktyWykresuFunkcji[167].Y = 4.06

	punktyWykresuFunkcji[168].X = 1.68
	punktyWykresuFunkcji[168].Y = 4.063

	punktyWykresuFunkcji[169].X = 1.69
	punktyWykresuFunkcji[169].Y = 4.066

	punktyWykresuFunkcji[170].X = 1.7
	punktyWykresuFunkcji[170].Y = 4.068

	punktyWykresuFunkcji[171].X = 1.71
	punktyWykresuFunkcji[171].Y = 4.07

	punktyWykresuFunkcji[172].X = 1.72
	punktyWykresuFunkcji[172].Y = 4.071

	punktyWykresuFunkcji[173].X = 1.73
	punktyWykresuFunkcji[173].Y = 4.07

	punktyWykresuFunkcji[174].X = 1.74
	punktyWykresuFunkcji[174].Y = 4.07

	punktyWykresuFunkcji[175].X = 1.75
	punktyWykresuFunkcji[175].Y = 4.068

	punktyWykresuFunkcji[176].X = 1.76
	punktyWykresuFunkcji[176].Y = 4.065

	punktyWykresuFunkcji[177].X = 1.77
	punktyWykresuFunkcji[177].Y = 4.062

	punktyWykresuFunkcji[178].X = 1.78
	punktyWykresuFunkcji[178].Y = 4.058

	punktyWykresuFunkcji[179].X = 1.79
	punktyWykresuFunkcji[179].Y = 4.053

	punktyWykresuFunkcji[180].X = 1.8
	punktyWykresuFunkcji[180].Y = 4.048

	punktyWykresuFunkcji[181].X = 1.81
	punktyWykresuFunkcji[181].Y = 4.041

	punktyWykresuFunkcji[182].X = 1.82
	punktyWykresuFunkcji[182].Y = 4.034

	punktyWykresuFunkcji[183].X = 1.83
	punktyWykresuFunkcji[183].Y = 4.025

	punktyWykresuFunkcji[184].X = 1.84
	punktyWykresuFunkcji[184].Y = 4.016

	punktyWykresuFunkcji[185].X = 1.85
	punktyWykresuFunkcji[185].Y = 4.006

	punktyWykresuFunkcji[186].X = 1.86
	punktyWykresuFunkcji[186].Y = 3.996

	punktyWykresuFunkcji[187].X = 1.87
	punktyWykresuFunkcji[187].Y = 3.984

	punktyWykresuFunkcji[188].X = 1.88
	punktyWykresuFunkcji[188].Y = 3.971

	punktyWykresuFunkcji[189].X = 1.89
	punktyWykresuFunkcji[189].Y = 3.958

	punktyWykresuFunkcji[190].X = 1.9
	punktyWykresuFunkcji[190].Y = 3.943

	punktyWykresuFunkcji[191].X = 1.91
	punktyWykresuFunkcji[191].Y = 3.928

	punktyWykresuFunkcji[192].X = 1.92
	punktyWykresuFunkcji[192].Y = 3.912

	punktyWykresuFunkcji[193].X = 1.93
	punktyWykresuFunkcji[193].Y = 3.894

	punktyWykresuFunkcji[194].X = 1.94
	punktyWykresuFunkcji[194].Y = 3.876

	punktyWykresuFunkcji[195].X = 1.95
	punktyWykresuFunkcji[195].Y = 3.857

	punktyWykresuFunkcji[196].X = 1.96
	punktyWykresuFunkcji[196].Y = 3.837

	punktyWykresuFunkcji[197].X = 1.97
	punktyWykresuFunkcji[197].Y = 3.816

	punktyWykresuFunkcji[198].X = 1.98
	punktyWykresuFunkcji[198].Y = 3.794

	punktyWykresuFunkcji[199].X = 1.99
	punktyWykresuFunkcji[199].Y = 3.771

	punktyWykresuFunkcji[200].X = 2.0
	punktyWykresuFunkcji[200].Y = 3.747

	punktyWykresuFunkcji[201].X = 2.01
	punktyWykresuFunkcji[201].Y = 3.722

	punktyWykresuFunkcji[202].X = 2.02
	punktyWykresuFunkcji[202].Y = 3.696

	punktyWykresuFunkcji[203].X = 2.03
	punktyWykresuFunkcji[203].Y = 3.669

	punktyWykresuFunkcji[204].X = 2.04
	punktyWykresuFunkcji[204].Y = 3.641

	punktyWykresuFunkcji[205].X = 2.05
	punktyWykresuFunkcji[205].Y = 3.612

	punktyWykresuFunkcji[206].X = 2.06
	punktyWykresuFunkcji[206].Y = 3.581

	punktyWykresuFunkcji[207].X = 2.07
	punktyWykresuFunkcji[207].Y = 3.55

	punktyWykresuFunkcji[208].X = 2.08
	punktyWykresuFunkcji[208].Y = 3.518

	punktyWykresuFunkcji[209].X = 2.09
	punktyWykresuFunkcji[209].Y = 3.484

	punktyWykresuFunkcji[210].X = 2.1
	punktyWykresuFunkcji[210].Y = 3.45

	punktyWykresuFunkcji[211].X = 2.11
	punktyWykresuFunkcji[211].Y = 3.414

	punktyWykresuFunkcji[212].X = 2.12
	punktyWykresuFunkcji[212].Y = 3.378

	punktyWykresuFunkcji[213].X = 2.13
	punktyWykresuFunkcji[213].Y = 3.34

	punktyWykresuFunkcji[214].X = 2.14
	punktyWykresuFunkcji[214].Y = 3.301

	punktyWykresuFunkcji[215].X = 2.15
	punktyWykresuFunkcji[215].Y = 3.261

	punktyWykresuFunkcji[216].X = 2.16
	punktyWykresuFunkcji[216].Y = 3.22

	punktyWykresuFunkcji[217].X = 2.17
	punktyWykresuFunkcji[217].Y = 3.177

	punktyWykresuFunkcji[218].X = 2.18
	punktyWykresuFunkcji[218].Y = 3.134

	punktyWykresuFunkcji[219].X = 2.19
	punktyWykresuFunkcji[219].Y = 3.089

	punktyWykresuFunkcji[220].X = 2.2
	punktyWykresuFunkcji[220].Y = 3.043

	punktyWykresuFunkcji[221].X = 2.21
	punktyWykresuFunkcji[221].Y = 2.996

	punktyWykresuFunkcji[222].X = 2.22
	punktyWykresuFunkcji[222].Y = 2.948

	punktyWykresuFunkcji[223].X = 2.23
	punktyWykresuFunkcji[223].Y = 2.899

	punktyWykresuFunkcji[224].X = 2.24
	punktyWykresuFunkcji[224].Y = 2.848

	punktyWykresuFunkcji[225].X = 2.25
	punktyWykresuFunkcji[225].Y = 2.797

	punktyWykresuFunkcji[226].X = 2.26
	punktyWykresuFunkcji[226].Y = 2.744

	punktyWykresuFunkcji[227].X = 2.27
	punktyWykresuFunkcji[227].Y = 2.69

	punktyWykresuFunkcji[228].X = 2.28
	punktyWykresuFunkcji[228].Y = 2.634

	punktyWykresuFunkcji[229].X = 2.29
	punktyWykresuFunkcji[229].Y = 2.577

	punktyWykresuFunkcji[230].X = 2.3
	punktyWykresuFunkcji[230].Y = 2.52

	punktyWykresuFunkcji[231].X = 2.31
	punktyWykresuFunkcji[231].Y = 2.46

	punktyWykresuFunkcji[232].X = 2.32
	punktyWykresuFunkcji[232].Y = 2.4

	punktyWykresuFunkcji[233].X = 2.33
	punktyWykresuFunkcji[233].Y = 2.338

	punktyWykresuFunkcji[234].X = 2.34
	punktyWykresuFunkcji[234].Y = 2.275

	punktyWykresuFunkcji[235].X = 2.35
	punktyWykresuFunkcji[235].Y = 2.211

	punktyWykresuFunkcji[236].X = 2.36
	punktyWykresuFunkcji[236].Y = 2.146

	punktyWykresuFunkcji[237].X = 2.37
	punktyWykresuFunkcji[237].Y = 2.079

	punktyWykresuFunkcji[238].X = 2.38
	punktyWykresuFunkcji[238].Y = 2.011

	punktyWykresuFunkcji[239].X = 2.39
	punktyWykresuFunkcji[239].Y = 1.941

	punktyWykresuFunkcji[240].X = 2.4
	punktyWykresuFunkcji[240].Y = 1.871

	punktyWykresuFunkcji[241].X = 2.41
	punktyWykresuFunkcji[241].Y = 1.798

	punktyWykresuFunkcji[242].X = 2.42
	punktyWykresuFunkcji[242].Y = 1.725

	punktyWykresuFunkcji[243].X = 2.43
	punktyWykresuFunkcji[243].Y = 1.65

	punktyWykresuFunkcji[244].X = 2.44
	punktyWykresuFunkcji[244].Y = 1.574

	punktyWykresuFunkcji[245].X = 2.45
	punktyWykresuFunkcji[245].Y = 1.497

	punktyWykresuFunkcji[246].X = 2.46
	punktyWykresuFunkcji[246].Y = 1.418

	punktyWykresuFunkcji[247].X = 2.47
	punktyWykresuFunkcji[247].Y = 1.338

	punktyWykresuFunkcji[248].X = 2.48
	punktyWykresuFunkcji[248].Y = 1.256

	punktyWykresuFunkcji[249].X = 2.49
	punktyWykresuFunkcji[249].Y = 1.173

	punktyWykresuFunkcji[250].X = 2.5
	punktyWykresuFunkcji[250].Y = 1.089

	punktyWykresuFunkcji[251].X = 2.51
	punktyWykresuFunkcji[251].Y = 1.003

	punktyWykresuFunkcji[252].X = 2.52
	punktyWykresuFunkcji[252].Y = 0.916

	punktyWykresuFunkcji[253].X = 2.53
	punktyWykresuFunkcji[253].Y = 0.827

	punktyWykresuFunkcji[254].X = 2.54
	punktyWykresuFunkcji[254].Y = 0.737

	punktyWykresuFunkcji[255].X = 2.55
	punktyWykresuFunkcji[255].Y = 0.646

	punktyWykresuFunkcji[256].X = 2.56
	punktyWykresuFunkcji[256].Y = 0.553

	punktyWykresuFunkcji[257].X = 2.57
	punktyWykresuFunkcji[257].Y = 0.458

	punktyWykresuFunkcji[258].X = 2.58
	punktyWykresuFunkcji[258].Y = 0.363

	punktyWykresuFunkcji[259].X = 2.59
	punktyWykresuFunkcji[259].Y = 0.265

	punktyWykresuFunkcji[260].X = 2.6
	punktyWykresuFunkcji[260].Y = 0.167

	punktyWykresuFunkcji[261].X = 2.61
	punktyWykresuFunkcji[261].Y = 0.067

	punktyWykresuFunkcji[262].X = 2.62
	punktyWykresuFunkcji[262].Y = -0.034

	punktyWykresuFunkcji[263].X = 2.63
	punktyWykresuFunkcji[263].Y = -0.137

	punktyWykresuFunkcji[264].X = 2.64
	punktyWykresuFunkcji[264].Y = -0.242

	punktyWykresuFunkcji[265].X = 2.65
	punktyWykresuFunkcji[265].Y = -0.348

	punktyWykresuFunkcji[266].X = 2.66
	punktyWykresuFunkcji[266].Y = -0.456

	punktyWykresuFunkcji[267].X = 2.67
	punktyWykresuFunkcji[267].Y = -0.565

	punktyWykresuFunkcji[268].X = 2.68
	punktyWykresuFunkcji[268].Y = -0.675

	punktyWykresuFunkcji[269].X = 2.69
	punktyWykresuFunkcji[269].Y = -0.788

	punktyWykresuFunkcji[270].X = 2.7
	punktyWykresuFunkcji[270].Y = -0.901

	punktyWykresuFunkcji[271].X = 2.71
	punktyWykresuFunkcji[271].Y = -1.016

	punktyWykresuFunkcji[272].X = 2.72
	punktyWykresuFunkcji[272].Y = -1.133

	punktyWykresuFunkcji[273].X = 2.73
	punktyWykresuFunkcji[273].Y = -1.251

	punktyWykresuFunkcji[274].X = 2.74
	punktyWykresuFunkcji[274].Y = -1.371

	punktyWykresuFunkcji[275].X = 2.75
	punktyWykresuFunkcji[275].Y = -1.493

	punktyWykresuFunkcji[276].X = 2.76
	punktyWykresuFunkcji[276].Y = -1.616

	punktyWykresuFunkcji[277].X = 2.77
	punktyWykresuFunkcji[277].Y = -1.74

	punktyWykresuFunkcji[278].X = 2.78
	punktyWykresuFunkcji[278].Y = -1.866

	punktyWykresuFunkcji[279].X = 2.79
	punktyWykresuFunkcji[279].Y = -1.994

	punktyWykresuFunkcji[280].X = 2.8
	punktyWykresuFunkcji[280].Y = -2.123

	punktyWykresuFunkcji[281].X = 2.81
	punktyWykresuFunkcji[281].Y = -2.254

	punktyWykresuFunkcji[282].X = 2.82
	punktyWykresuFunkcji[282].Y = -2.386

	punktyWykresuFunkcji[283].X = 2.83
	punktyWykresuFunkcji[283].Y = -2.52

	punktyWykresuFunkcji[284].X = 2.84
	punktyWykresuFunkcji[284].Y = -2.656

	punktyWykresuFunkcji[285].X = 2.85
	punktyWykresuFunkcji[285].Y = -2.793

	punktyWykresuFunkcji[286].X = 2.86
	punktyWykresuFunkcji[286].Y = -2.932

	punktyWykresuFunkcji[287].X = 2.87
	punktyWykresuFunkcji[287].Y = -3.073

	punktyWykresuFunkcji[288].X = 2.88
	punktyWykresuFunkcji[288].Y = -3.215

	punktyWykresuFunkcji[289].X = 2.89
	punktyWykresuFunkcji[289].Y = -3.359

	punktyWykresuFunkcji[290].X = 2.9
	punktyWykresuFunkcji[290].Y = -3.504

	punktyWykresuFunkcji[291].X = 2.91
	punktyWykresuFunkcji[291].Y = -3.651

	punktyWykresuFunkcji[292].X = 2.92
	punktyWykresuFunkcji[292].Y = -3.8

	punktyWykresuFunkcji[293].X = 2.93
	punktyWykresuFunkcji[293].Y = -3.95

	punktyWykresuFunkcji[294].X = 2.94
	punktyWykresuFunkcji[294].Y = -4.102

	punktyWykresuFunkcji[295].X = 2.95
	punktyWykresuFunkcji[295].Y = -4.256

	punktyWykresuFunkcji[296].X = 2.96
	punktyWykresuFunkcji[296].Y = -4.411

	punktyWykresuFunkcji[297].X = 2.97
	punktyWykresuFunkcji[297].Y = -4.568

	punktyWykresuFunkcji[298].X = 2.98
	punktyWykresuFunkcji[298].Y = -4.727

	punktyWykresuFunkcji[299].X = 2.99
	punktyWykresuFunkcji[299].Y = -4.887

	punktyWykresuFunkcji[300].X = 3.0
	punktyWykresuFunkcji[300].Y = -5.05

	punktyWykresuFunkcji[301].X = 3.01
	punktyWykresuFunkcji[301].Y = -5.213

	punktyWykresuFunkcji[302].X = 3.02
	punktyWykresuFunkcji[302].Y = -5.379

	punktyWykresuFunkcji[303].X = 3.03
	punktyWykresuFunkcji[303].Y = -5.546

	punktyWykresuFunkcji[304].X = 3.04
	punktyWykresuFunkcji[304].Y = -5.715

	punktyWykresuFunkcji[305].X = 3.05
	punktyWykresuFunkcji[305].Y = -5.886

	punktyWykresuFunkcji[306].X = 3.06
	punktyWykresuFunkcji[306].Y = -6.058

	punktyWykresuFunkcji[307].X = 3.07
	punktyWykresuFunkcji[307].Y = -6.232

	punktyWykresuFunkcji[308].X = 3.08
	punktyWykresuFunkcji[308].Y = -6.408

	punktyWykresuFunkcji[309].X = 3.09
	punktyWykresuFunkcji[309].Y = -6.585

	punktyWykresuFunkcji[310].X = 3.1
	punktyWykresuFunkcji[310].Y = -6.765

	punktyWykresuFunkcji[311].X = 3.11
	punktyWykresuFunkcji[311].Y = -6.946

	punktyWykresuFunkcji[312].X = 3.12
	punktyWykresuFunkcji[312].Y = -7.128

	punktyWykresuFunkcji[313].X = 3.13
	punktyWykresuFunkcji[313].Y = -7.313

	punktyWykresuFunkcji[314].X = 3.14
	punktyWykresuFunkcji[314].Y = -7.499

	punktyWykresuFunkcji[315].X = 3.15
	punktyWykresuFunkcji[315].Y = -7.687

	punktyWykresuFunkcji[316].X = 3.16
	punktyWykresuFunkcji[316].Y = -7.887

	punktyWykresuFunkcji[317].X = 3.17
	punktyWykresuFunkcji[317].Y = -8.068

	punktyWykresuFunkcji[318].X = 3.18
	punktyWykresuFunkcji[318].Y = -8.262

	punktyWykresuFunkcji[319].X = 3.19
	punktyWykresuFunkcji[319].Y = -8.457

	punktyWykresuFunkcji[320].X = 3.2
	punktyWykresuFunkcji[320].Y = -8.653

	punktyWykresuFunkcji[321].X = 3.21
	punktyWykresuFunkcji[321].Y = -8.852

	punktyWykresuFunkcji[322].X = 3.22
	punktyWykresuFunkcji[322].Y = -9.052

	punktyWykresuFunkcji[323].X = 3.23
	punktyWykresuFunkcji[323].Y = -9.255

	punktyWykresuFunkcji[324].X = 3.24
	punktyWykresuFunkcji[324].Y = -9.459

	punktyWykresuFunkcji[325].X = 3.25
	punktyWykresuFunkcji[325].Y = -9.664

	punktyWykresuFunkcji[326].X = 3.26
	punktyWykresuFunkcji[326].Y = -9.872

	punktyWykresuFunkcji[327].X = 3.27
	punktyWykresuFunkcji[327].Y = -10.081

	punktyWykresuFunkcji[328].X = 3.28
	punktyWykresuFunkcji[328].Y = -10.292

	punktyWykresuFunkcji[329].X = 3.29
	punktyWykresuFunkcji[329].Y = -10.505

	punktyWykresuFunkcji[330].X = 3.3
	punktyWykresuFunkcji[330].Y = -10.72

	punktyWykresuFunkcji[331].X = 3.31
	punktyWykresuFunkcji[331].Y = -10.937

	punktyWykresuFunkcji[332].X = 3.32
	punktyWykresuFunkcji[332].Y = -11.155

	punktyWykresuFunkcji[333].X = 3.33
	punktyWykresuFunkcji[333].Y = -11.375

	punktyWykresuFunkcji[334].X = 3.34
	punktyWykresuFunkcji[334].Y = -11.597

	punktyWykresuFunkcji[335].X = 3.35
	punktyWykresuFunkcji[335].Y = -11.821

	punktyWykresuFunkcji[336].X = 3.36
	punktyWykresuFunkcji[336].Y = -12.047

	punktyWykresuFunkcji[337].X = 3.37
	punktyWykresuFunkcji[337].Y = -12.275

	punktyWykresuFunkcji[338].X = 3.38
	punktyWykresuFunkcji[338].Y = -12.504

	punktyWykresuFunkcji[339].X = 3.39
	punktyWykresuFunkcji[339].Y = -12.735

	punktyWykresuFunkcji[340].X = 3.4
	punktyWykresuFunkcji[340].Y = -12.968

	punktyWykresuFunkcji[341].X = 3.41
	punktyWykresuFunkcji[341].Y = -13.203

	punktyWykresuFunkcji[342].X = 3.42
	punktyWykresuFunkcji[342].Y = -13.44

	punktyWykresuFunkcji[343].X = 3.43
	punktyWykresuFunkcji[343].Y = -13.678

	punktyWykresuFunkcji[344].X = 3.44
	punktyWykresuFunkcji[344].Y = -13.919

	punktyWykresuFunkcji[345].X = 3.45
	punktyWykresuFunkcji[345].Y = -14.161

	punktyWykresuFunkcji[346].X = 3.46
	punktyWykresuFunkcji[346].Y = -14.405

	punktyWykresuFunkcji[347].X = 3.47
	punktyWykresuFunkcji[347].Y = -14.651

	punktyWykresuFunkcji[348].X = 3.48
	punktyWykresuFunkcji[348].Y = -14.899

	punktyWykresuFunkcji[349].X = 3.49
	punktyWykresuFunkcji[349].Y = -15.149

	punktyWykresuFunkcji[350].X = 3.5
	punktyWykresuFunkcji[350].Y = -15.401

	punktyWykresuFunkcji[351].X = 3.51
	punktyWykresuFunkcji[351].Y = -15.654

	punktyWykresuFunkcji[352].X = 3.52
	punktyWykresuFunkcji[352].Y = -15.909

	punktyWykresuFunkcji[353].X = 3.53
	punktyWykresuFunkcji[353].Y = -16.167

	punktyWykresuFunkcji[354].X = 3.54
	punktyWykresuFunkcji[354].Y = -16.426

	punktyWykresuFunkcji[355].X = 3.55
	punktyWykresuFunkcji[355].Y = -16.687

	punktyWykresuFunkcji[356].X = 3.56
	punktyWykresuFunkcji[356].Y = -16.95

	punktyWykresuFunkcji[357].X = 3.57
	punktyWykresuFunkcji[357].Y = -17.214

	punktyWykresuFunkcji[358].X = 3.58
	punktyWykresuFunkcji[358].Y = -17.481

	punktyWykresuFunkcji[359].X = 3.59
	punktyWykresuFunkcji[359].Y = -17.749

	punktyWykresuFunkcji[360].X = 3.6
	punktyWykresuFunkcji[360].Y = -18.02

	punktyWykresuFunkcji[361].X = 3.61
	punktyWykresuFunkcji[361].Y = -18.292

	punktyWykresuFunkcji[362].X = 3.62
	punktyWykresuFunkcji[362].Y = -18.566

	punktyWykresuFunkcji[363].X = 3.63
	punktyWykresuFunkcji[363].Y = -18.842

	punktyWykresuFunkcji[364].X = 3.64
	punktyWykresuFunkcji[364].Y = -19.12

	punktyWykresuFunkcji[365].X = 3.65
	punktyWykresuFunkcji[365].Y = -19.4

	punktyWykresuFunkcji[366].X = 3.66
	punktyWykresuFunkcji[366].Y = -19.682

	punktyWykresuFunkcji[367].X = 3.67
	punktyWykresuFunkcji[367].Y = -19.965

	punktyWykresuFunkcji[368].X = 3.68
	punktyWykresuFunkcji[368].Y = -20.251

	punktyWykresuFunkcji[369].X = 3.69
	punktyWykresuFunkcji[369].Y = -20.538

	punktyWykresuFunkcji[370].X = 3.7
	punktyWykresuFunkcji[370].Y = -20.828

	punktyWykresuFunkcji[371].X = 3.71
	punktyWykresuFunkcji[371].Y = -21.119

	punktyWykresuFunkcji[372].X = 3.72
	punktyWykresuFunkcji[372].Y = -21.412

	punktyWykresuFunkcji[373].X = 3.73
	punktyWykresuFunkcji[373].Y = -21.707

	punktyWykresuFunkcji[374].X = 3.74
	punktyWykresuFunkcji[374].Y = -22.004

	punktyWykresuFunkcji[375].X = 3.75
	punktyWykresuFunkcji[375].Y = -22.303

	punktyWykresuFunkcji[376].X = 3.76
	punktyWykresuFunkcji[376].Y = -22.604

	punktyWykresuFunkcji[377].X = 3.77
	punktyWykresuFunkcji[377].Y = -22.906

	punktyWykresuFunkcji[378].X = 3.78
	punktyWykresuFunkcji[378].Y = -23.211

	punktyWykresuFunkcji[379].X = 3.79
	punktyWykresuFunkcji[379].Y = -23.518

	punktyWykresuFunkcji[380].X = 3.8
	punktyWykresuFunkcji[380].Y = -23.826

	punktyWykresuFunkcji[381].X = 3.81
	punktyWykresuFunkcji[381].Y = -24.136

	punktyWykresuFunkcji[382].X = 3.82
	punktyWykresuFunkcji[382].Y = -24.449

	punktyWykresuFunkcji[383].X = 3.83
	punktyWykresuFunkcji[383].Y = -24.763

	punktyWykresuFunkcji[384].X = 3.84
	punktyWykresuFunkcji[384].Y = -25.079

	punktyWykresuFunkcji[385].X = 3.85
	punktyWykresuFunkcji[385].Y = -25.397

	punktyWykresuFunkcji[386].X = 3.86
	punktyWykresuFunkcji[386].Y = -25.717

	punktyWykresuFunkcji[387].X = 3.87
	punktyWykresuFunkcji[387].Y = -26.039

	punktyWykresuFunkcji[388].X = 3.88
	punktyWykresuFunkcji[388].Y = -26.363

	punktyWykresuFunkcji[389].X = 3.89
	punktyWykresuFunkcji[389].Y = -26.688

	punktyWykresuFunkcji[390].X = 3.9
	punktyWykresuFunkcji[390].Y = -27.016

	punktyWykresuFunkcji[391].X = 3.91
	punktyWykresuFunkcji[391].Y = -27.345

	punktyWykresuFunkcji[392].X = 3.92
	punktyWykresuFunkcji[392].Y = -27.677

	punktyWykresuFunkcji[393].X = 3.93
	punktyWykresuFunkcji[393].Y = -28.01

	punktyWykresuFunkcji[394].X = 3.94
	punktyWykresuFunkcji[394].Y = -28.346

	punktyWykresuFunkcji[395].X = 3.95
	punktyWykresuFunkcji[395].Y = -28.683

	punktyWykresuFunkcji[396].X = 3.96
	punktyWykresuFunkcji[396].Y = -29.022

	punktyWykresuFunkcji[397].X = 3.97
	punktyWykresuFunkcji[397].Y = -29.363

	punktyWykresuFunkcji[398].X = 3.98
	punktyWykresuFunkcji[398].Y = -29.706

	punktyWykresuFunkcji[399].X = 3.99
	punktyWykresuFunkcji[399].Y = -30.051

	punktyWykresuFunkcji[400].X = 4.0
	punktyWykresuFunkcji[400].Y = -30.398

	punktyWykresuFunkcji[401].X = 4.01
	punktyWykresuFunkcji[401].Y = -30.747

	punktyWykresuFunkcji[402].X = 4.02
	punktyWykresuFunkcji[402].Y = -31.097

	punktyWykresuFunkcji[403].X = 4.03
	punktyWykresuFunkcji[403].Y = -31.45

	punktyWykresuFunkcji[404].X = 4.04
	punktyWykresuFunkcji[404].Y = -31.805

	punktyWykresuFunkcji[405].X = 4.05
	punktyWykresuFunkcji[405].Y = -32.161

	punktyWykresuFunkcji[406].X = 4.06
	punktyWykresuFunkcji[406].Y = -32.52

	punktyWykresuFunkcji[407].X = 4.07
	punktyWykresuFunkcji[407].Y = -32.88

	punktyWykresuFunkcji[408].X = 4.08
	punktyWykresuFunkcji[408].Y = -33.242

	punktyWykresuFunkcji[409].X = 4.09
	punktyWykresuFunkcji[409].Y = -33.607

	punktyWykresuFunkcji[410].X = 4.1
	punktyWykresuFunkcji[410].Y = -33.973

	punktyWykresuFunkcji[411].X = 4.11
	punktyWykresuFunkcji[411].Y = -34.341

	punktyWykresuFunkcji[412].X = 4.12
	punktyWykresuFunkcji[412].Y = -34.711

	punktyWykresuFunkcji[413].X = 4.13
	punktyWykresuFunkcji[413].Y = -35.083

	punktyWykresuFunkcji[414].X = 4.14
	punktyWykresuFunkcji[414].Y = -35.457

	punktyWykresuFunkcji[415].X = 4.15
	punktyWykresuFunkcji[415].Y = -35.832

	punktyWykresuFunkcji[416].X = 4.16
	punktyWykresuFunkcji[416].Y = -36.21

	punktyWykresuFunkcji[417].X = 4.17
	punktyWykresuFunkcji[417].Y = -36.59

	punktyWykresuFunkcji[418].X = 4.18
	punktyWykresuFunkcji[418].Y = -36.971

	punktyWykresuFunkcji[419].X = 4.19
	punktyWykresuFunkcji[419].Y = -37.355

	punktyWykresuFunkcji[420].X = 4.2
	punktyWykresuFunkcji[420].Y = -37.74

	punktyWykresuFunkcji[421].X = 4.21
	punktyWykresuFunkcji[421].Y = -38.128

	punktyWykresuFunkcji[422].X = 4.22
	punktyWykresuFunkcji[422].Y = -38.517

	punktyWykresuFunkcji[423].X = 4.23
	punktyWykresuFunkcji[423].Y = -38.908

	punktyWykresuFunkcji[424].X = 4.24
	punktyWykresuFunkcji[424].Y = -39.301

	punktyWykresuFunkcji[425].X = 4.25
	punktyWykresuFunkcji[425].Y = -39.696

	punktyWykresuFunkcji[426].X = 4.26
	punktyWykresuFunkcji[426].Y = -40.093

	punktyWykresuFunkcji[427].X = 4.27
	punktyWykresuFunkcji[427].Y = -40.492

	punktyWykresuFunkcji[428].X = 4.28
	punktyWykresuFunkcji[428].Y = -40.893

	punktyWykresuFunkcji[429].X = 4.29
	punktyWykresuFunkcji[429].Y = -41.296

	punktyWykresuFunkcji[430].X = 4.3
	punktyWykresuFunkcji[430].Y = -41.7

	punktyWykresuFunkcji[431].X = 4.31
	punktyWykresuFunkcji[431].Y = -42.107

	punktyWykresuFunkcji[432].X = 4.32
	punktyWykresuFunkcji[432].Y = -42.515

	punktyWykresuFunkcji[433].X = 4.33
	punktyWykresuFunkcji[433].Y = -42.926

	punktyWykresuFunkcji[434].X = 4.34
	punktyWykresuFunkcji[434].Y = -43.338

	punktyWykresuFunkcji[435].X = 4.35
	punktyWykresuFunkcji[435].Y = -43.752

	punktyWykresuFunkcji[436].X = 4.36
	punktyWykresuFunkcji[436].Y = -44.168

	punktyWykresuFunkcji[437].X = 4.37
	punktyWykresuFunkcji[437].Y = -44.586

	punktyWykresuFunkcji[438].X = 4.38
	punktyWykresuFunkcji[438].Y = -45.006

	punktyWykresuFunkcji[439].X = 4.39
	punktyWykresuFunkcji[439].Y = -45.428

	punktyWykresuFunkcji[440].X = 4.4
	punktyWykresuFunkcji[440].Y = -45.852

	punktyWykresuFunkcji[441].X = 4.41
	punktyWykresuFunkcji[441].Y = -46.278

	punktyWykresuFunkcji[442].X = 4.42
	punktyWykresuFunkcji[442].Y = -46.706

	punktyWykresuFunkcji[443].X = 4.43
	punktyWykresuFunkcji[443].Y = -47.135

	punktyWykresuFunkcji[444].X = 4.44
	punktyWykresuFunkcji[444].Y = -47.567

	punktyWykresuFunkcji[445].X = 4.45
	punktyWykresuFunkcji[445].Y = -48.0

	punktyWykresuFunkcji[446].X = 4.46
	punktyWykresuFunkcji[446].Y = -48.435

	punktyWykresuFunkcji[447].X = 4.47
	punktyWykresuFunkcji[447].Y = -48.872

	punktyWykresuFunkcji[448].X = 4.48
	punktyWykresuFunkcji[448].Y = -49.312

	punktyWykresuFunkcji[449].X = 4.49
	punktyWykresuFunkcji[449].Y = -49.753

	punktyWykresuFunkcji[450].X = 4.5
	punktyWykresuFunkcji[450].Y = -50.196

	punktyWykresuFunkcji[451].X = 4.51
	punktyWykresuFunkcji[451].Y = -50.64

	punktyWykresuFunkcji[452].X = 4.52
	punktyWykresuFunkcji[452].Y = -51.087

	punktyWykresuFunkcji[453].X = 4.53
	punktyWykresuFunkcji[453].Y = -51.536

	punktyWykresuFunkcji[454].X = 4.54
	punktyWykresuFunkcji[454].Y = -51.986

	punktyWykresuFunkcji[455].X = 4.55
	punktyWykresuFunkcji[455].Y = -52.439

	punktyWykresuFunkcji[456].X = 4.56
	punktyWykresuFunkcji[456].Y = -52.893

	punktyWykresuFunkcji[457].X = 4.57
	punktyWykresuFunkcji[457].Y = -53.349

	punktyWykresuFunkcji[458].X = 4.58
	punktyWykresuFunkcji[458].Y = -53.807

	punktyWykresuFunkcji[459].X = 4.59
	punktyWykresuFunkcji[459].Y = -54.267

	punktyWykresuFunkcji[460].X = 4.6
	punktyWykresuFunkcji[460].Y = -54.729

	punktyWykresuFunkcji[461].X = 4.61
	punktyWykresuFunkcji[461].Y = -55.193

	punktyWykresuFunkcji[462].X = 4.62
	punktyWykresuFunkcji[462].Y = -55.659

	punktyWykresuFunkcji[463].X = 4.63
	punktyWykresuFunkcji[463].Y = -56.127

	punktyWykresuFunkcji[464].X = 4.64
	punktyWykresuFunkcji[464].Y = -56.596

	punktyWykresuFunkcji[465].X = 4.65
	punktyWykresuFunkcji[465].Y = -57.067

	punktyWykresuFunkcji[466].X = 4.66
	punktyWykresuFunkcji[466].Y = -57.541

	punktyWykresuFunkcji[467].X = 4.67
	punktyWykresuFunkcji[467].Y = -58.016

	punktyWykresuFunkcji[468].X = 4.68
	punktyWykresuFunkcji[468].Y = -58.493

	punktyWykresuFunkcji[469].X = 4.69
	punktyWykresuFunkcji[469].Y = -58.972

	punktyWykresuFunkcji[470].X = 4.7
	punktyWykresuFunkcji[470].Y = -59.453

	punktyWykresuFunkcji[471].X = 4.71
	punktyWykresuFunkcji[471].Y = -59.936

	punktyWykresuFunkcji[472].X = 4.72
	punktyWykresuFunkcji[472].Y = -60.42

	punktyWykresuFunkcji[473].X = 4.73
	punktyWykresuFunkcji[473].Y = -60.907

	punktyWykresuFunkcji[474].X = 4.74
	punktyWykresuFunkcji[474].Y = -61.395

	punktyWykresuFunkcji[475].X = 4.75
	punktyWykresuFunkcji[475].Y = -61.885

	punktyWykresuFunkcji[476].X = 4.76
	punktyWykresuFunkcji[476].Y = -62.378

	punktyWykresuFunkcji[477].X = 4.77
	punktyWykresuFunkcji[477].Y = -62.872

	punktyWykresuFunkcji[478].X = 4.78
	punktyWykresuFunkcji[478].Y = -63.368

	punktyWykresuFunkcji[479].X = 4.79
	punktyWykresuFunkcji[479].Y = -63.865

	punktyWykresuFunkcji[480].X = 4.8
	punktyWykresuFunkcji[480].Y = -64.365

	punktyWykresuFunkcji[481].X = 4.81
	punktyWykresuFunkcji[481].Y = -64.867

	punktyWykresuFunkcji[482].X = 4.82
	punktyWykresuFunkcji[482].Y = -65.37

	punktyWykresuFunkcji[483].X = 4.83
	punktyWykresuFunkcji[483].Y = -65.875

	punktyWykresuFunkcji[484].X = 4.84
	punktyWykresuFunkcji[484].Y = -66.382

	punktyWykresuFunkcji[485].X = 4.85
	punktyWykresuFunkcji[485].Y = -66.891

	punktyWykresuFunkcji[486].X = 4.86
	punktyWykresuFunkcji[486].Y = -67.402

	punktyWykresuFunkcji[487].X = 4.87
	punktyWykresuFunkcji[487].Y = -67.915

	punktyWykresuFunkcji[488].X = 4.88
	punktyWykresuFunkcji[488].Y = -68.43

	punktyWykresuFunkcji[489].X = 4.89
	punktyWykresuFunkcji[489].Y = -68.946

	punktyWykresuFunkcji[490].X = 4.9
	punktyWykresuFunkcji[490].Y = -69.465

	punktyWykresuFunkcji[491].X = 4.91
	punktyWykresuFunkcji[491].Y = -69.985

	punktyWykresuFunkcji[492].X = 4.92
	punktyWykresuFunkcji[492].Y = -70.507

	punktyWykresuFunkcji[493].X = 4.93
	punktyWykresuFunkcji[493].Y = -71.031

	punktyWykresuFunkcji[494].X = 4.94
	punktyWykresuFunkcji[494].Y = -71.557

	punktyWykresuFunkcji[495].X = 4.95
	punktyWykresuFunkcji[495].Y = -72.085

	punktyWykresuFunkcji[496].X = 4.96
	punktyWykresuFunkcji[496].Y = -72.614

	punktyWykresuFunkcji[497].X = 4.97
	punktyWykresuFunkcji[497].Y = -73.146

	punktyWykresuFunkcji[498].X = 4.98
	punktyWykresuFunkcji[498].Y = -73.679

	punktyWykresuFunkcji[499].X = 4.99
	punktyWykresuFunkcji[499].Y = -74.214

	punktyWykresuFunkcji[500].X = 5.0
	punktyWykresuFunkcji[500].Y = -74.751

	punktyWykresuFunkcji[501].X = 5.01
	punktyWykresuFunkcji[501].Y = -75.29

	punktyWykresuFunkcji[502].X = 5.02
	punktyWykresuFunkcji[502].Y = -75.831

	punktyWykresuFunkcji[503].X = 5.03
	punktyWykresuFunkcji[503].Y = -76.373

	punktyWykresuFunkcji[504].X = 5.04
	punktyWykresuFunkcji[504].Y = -76.918

	punktyWykresuFunkcji[505].X = 5.05
	punktyWykresuFunkcji[505].Y = -77.464

	punktyWykresuFunkcji[506].X = 5.06
	punktyWykresuFunkcji[506].Y = -78.012

	punktyWykresuFunkcji[507].X = 5.07
	punktyWykresuFunkcji[507].Y = -78.562

	punktyWykresuFunkcji[508].X = 5.08
	punktyWykresuFunkcji[508].Y = -79.114

	punktyWykresuFunkcji[509].X = 5.09
	punktyWykresuFunkcji[509].Y = -79.668

	punktyWykresuFunkcji[510].X = 5.1
	punktyWykresuFunkcji[510].Y = -80.223

	punktyWykresuFunkcji[511].X = 5.11
	punktyWykresuFunkcji[511].Y = -80.781

	punktyWykresuFunkcji[512].X = 5.12
	punktyWykresuFunkcji[512].Y = -81.34

	punktyWykresuFunkcji[513].X = 5.13
	punktyWykresuFunkcji[513].Y = -81.901

	punktyWykresuFunkcji[514].X = 5.14
	punktyWykresuFunkcji[514].Y = -82.464

	punktyWykresuFunkcji[515].X = 5.15
	punktyWykresuFunkcji[515].Y = -83.029

	punktyWykresuFunkcji[516].X = 5.16
	punktyWykresuFunkcji[516].Y = -83.596

	punktyWykresuFunkcji[517].X = 5.17
	punktyWykresuFunkcji[517].Y = -84.164

	punktyWykresuFunkcji[518].X = 5.18
	punktyWykresuFunkcji[518].Y = -84.734

	punktyWykresuFunkcji[519].X = 5.19
	punktyWykresuFunkcji[519].Y = -85.307

	punktyWykresuFunkcji[520].X = 5.2
	punktyWykresuFunkcji[520].Y = -85.881

	punktyWykresuFunkcji[521].X = 5.21
	punktyWykresuFunkcji[521].Y = -86.457

	punktyWykresuFunkcji[522].X = 5.22
	punktyWykresuFunkcji[522].Y = -87.034

	punktyWykresuFunkcji[523].X = 5.23
	punktyWykresuFunkcji[523].Y = -87.614

	punktyWykresuFunkcji[524].X = 5.24
	punktyWykresuFunkcji[524].Y = -88.195

	punktyWykresuFunkcji[525].X = 5.25
	punktyWykresuFunkcji[525].Y = -88.779

	punktyWykresuFunkcji[526].X = 5.26
	punktyWykresuFunkcji[526].Y = -89.364

	punktyWykresuFunkcji[527].X = 5.27
	punktyWykresuFunkcji[527].Y = -89.951

	punktyWykresuFunkcji[528].X = 5.28
	punktyWykresuFunkcji[528].Y = -90.54

	punktyWykresuFunkcji[529].X = 5.29
	punktyWykresuFunkcji[529].Y = -91.13

	punktyWykresuFunkcji[530].X = 5.3
	punktyWykresuFunkcji[530].Y = -91.723

	punktyWykresuFunkcji[531].X = 5.31
	punktyWykresuFunkcji[531].Y = -92.317

	punktyWykresuFunkcji[532].X = 5.32
	punktyWykresuFunkcji[532].Y = -92.913

	punktyWykresuFunkcji[533].X = 5.33
	punktyWykresuFunkcji[533].Y = -93.511

	punktyWykresuFunkcji[534].X = 5.34
	punktyWykresuFunkcji[534].Y = -94.111

	punktyWykresuFunkcji[535].X = 5.35
	punktyWykresuFunkcji[535].Y = -94.713

	punktyWykresuFunkcji[536].X = 5.36
	punktyWykresuFunkcji[536].Y = -95.316

	punktyWykresuFunkcji[537].X = 5.37
	punktyWykresuFunkcji[537].Y = -95.922

	punktyWykresuFunkcji[538].X = 5.38
	punktyWykresuFunkcji[538].Y = -96.529

	punktyWykresuFunkcji[539].X = 5.39
	punktyWykresuFunkcji[539].Y = -97.138

	punktyWykresuFunkcji[540].X = 5.4
	punktyWykresuFunkcji[540].Y = -97.749

	punktyWykresuFunkcji[541].X = 5.41
	punktyWykresuFunkcji[541].Y = -98.362

	punktyWykresuFunkcji[542].X = 5.42
	punktyWykresuFunkcji[542].Y = -98.976

	punktyWykresuFunkcji[543].X = 5.43
	punktyWykresuFunkcji[543].Y = -99.593

	punktyWykresuFunkcji[544].X = 5.44
	punktyWykresuFunkcji[544].Y = -100.211

	punktyWykresuFunkcji[545].X = 5.45
	punktyWykresuFunkcji[545].Y = -100.831

	punktyWykresuFunkcji[546].X = 5.46
	punktyWykresuFunkcji[546].Y = -101.453

	punktyWykresuFunkcji[547].X = 5.47
	punktyWykresuFunkcji[547].Y = -102.077

	punktyWykresuFunkcji[548].X = 5.48
	punktyWykresuFunkcji[548].Y = -102.703

	punktyWykresuFunkcji[549].X = 5.49
	punktyWykresuFunkcji[549].Y = -103.33

	punktyWykresuFunkcji[550].X = 5.5
	punktyWykresuFunkcji[550].Y = -103.959

	punktyWykresuFunkcji[551].X = 5.51
	punktyWykresuFunkcji[551].Y = -104.591

	punktyWykresuFunkcji[552].X = 5.52
	punktyWykresuFunkcji[552].Y = -105.224

	punktyWykresuFunkcji[553].X = 5.53
	punktyWykresuFunkcji[553].Y = -105.859

	punktyWykresuFunkcji[554].X = 5.54
	punktyWykresuFunkcji[554].Y = -106.495

	punktyWykresuFunkcji[555].X = 5.55
	punktyWykresuFunkcji[555].Y = -107.134

	punktyWykresuFunkcji[556].X = 5.56
	punktyWykresuFunkcji[556].Y = -107.774

	punktyWykresuFunkcji[557].X = 5.57
	punktyWykresuFunkcji[557].Y = -108.417

	punktyWykresuFunkcji[558].X = 5.58
	punktyWykresuFunkcji[558].Y = -109.061

	punktyWykresuFunkcji[559].X = 5.59
	punktyWykresuFunkcji[559].Y = -109.707

	punktyWykresuFunkcji[560].X = 5.6
	punktyWykresuFunkcji[560].Y = -110.355

	punktyWykresuFunkcji[561].X = 5.61
	punktyWykresuFunkcji[561].Y = -111.004

	punktyWykresuFunkcji[562].X = 5.62
	punktyWykresuFunkcji[562].Y = -111.653

	punktyWykresuFunkcji[563].X = 5.63
	punktyWykresuFunkcji[563].Y = -112.309

	punktyWykresuFunkcji[564].X = 5.64
	punktyWykresuFunkcji[564].Y = -112.964

	punktyWykresuFunkcji[565].X = 5.65
	punktyWykresuFunkcji[565].Y = -113.622

	punktyWykresuFunkcji[566].X = 5.66
	punktyWykresuFunkcji[566].Y = -114.281

	punktyWykresuFunkcji[567].X = 5.67
	punktyWykresuFunkcji[567].Y = -114.941

	punktyWykresuFunkcji[568].X = 5.68
	punktyWykresuFunkcji[568].Y = -115.604

	punktyWykresuFunkcji[569].X = 5.69
	punktyWykresuFunkcji[569].Y = -116.269

	punktyWykresuFunkcji[570].X = 5.7
	punktyWykresuFunkcji[570].Y = -116.935

	punktyWykresuFunkcji[571].X = 5.71
	punktyWykresuFunkcji[571].Y = -117.603

	punktyWykresuFunkcji[572].X = 5.72
	punktyWykresuFunkcji[572].Y = -118.273

	punktyWykresuFunkcji[573].X = 5.73
	punktyWykresuFunkcji[573].Y = -118.945

	punktyWykresuFunkcji[574].X = 5.74
	punktyWykresuFunkcji[574].Y = -119.619

	punktyWykresuFunkcji[575].X = 5.75
	punktyWykresuFunkcji[575].Y = -120.295

	punktyWykresuFunkcji[576].X = 5.76
	punktyWykresuFunkcji[576].Y = -120.973

	punktyWykresuFunkcji[577].X = 5.77
	punktyWykresuFunkcji[577].Y = -121.652

	punktyWykresuFunkcji[578].X = 5.78
	punktyWykresuFunkcji[578].Y = -122.333

	punktyWykresuFunkcji[579].X = 5.79
	punktyWykresuFunkcji[579].Y = -123.017

	punktyWykresuFunkcji[580].X = 5.8
	punktyWykresuFunkcji[580].Y = -123.702

	punktyWykresuFunkcji[581].X = 5.81
	punktyWykresuFunkcji[581].Y = -124.389

	punktyWykresuFunkcji[582].X = 5.82
	punktyWykresuFunkcji[582].Y = -125.077

	punktyWykresuFunkcji[583].X = 5.83
	punktyWykresuFunkcji[583].Y = -125.768

	punktyWykresuFunkcji[584].X = 5.84
	punktyWykresuFunkcji[584].Y = -126.461

	punktyWykresuFunkcji[585].X = 5.85
	punktyWykresuFunkcji[585].Y = -127.155

	punktyWykresuFunkcji[586].X = 5.86
	punktyWykresuFunkcji[586].Y = -127.852

	punktyWykresuFunkcji[587].X = 5.87
	punktyWykresuFunkcji[587].Y = -128.55

	punktyWykresuFunkcji[588].X = 5.88
	punktyWykresuFunkcji[588].Y = -129.25

	punktyWykresuFunkcji[589].X = 5.89
	punktyWykresuFunkcji[589].Y = -129.952

	punktyWykresuFunkcji[590].X = 5.9
	punktyWykresuFunkcji[590].Y = -130.656

	punktyWykresuFunkcji[591].X = 5.91
	punktyWykresuFunkcji[591].Y = -131.362

	punktyWykresuFunkcji[592].X = 5.92
	punktyWykresuFunkcji[592].Y = -132.07

	punktyWykresuFunkcji[593].X = 5.93
	punktyWykresuFunkcji[593].Y = -132.779

	punktyWykresuFunkcji[594].X = 5.94
	punktyWykresuFunkcji[594].Y = -133.491

	punktyWykresuFunkcji[595].X = 5.95
	punktyWykresuFunkcji[595].Y = -134.204

	punktyWykresuFunkcji[596].X = 5.96
	punktyWykresuFunkcji[596].Y = -134.92

	punktyWykresuFunkcji[597].X = 5.97
	punktyWykresuFunkcji[597].Y = -135.637

	punktyWykresuFunkcji[598].X = 5.98
	punktyWykresuFunkcji[598].Y = -136.356

	punktyWykresuFunkcji[599].X = 5.99
	punktyWykresuFunkcji[599].Y = -137.077

	punktyWykresuFunkcji[600].X = 6.0
	punktyWykresuFunkcji[600].Y = -137.8

	punktyWykresuFunkcji[601].X = 6.01
	punktyWykresuFunkcji[601].Y = -138.525

	punktyWykresuFunkcji[602].X = 6.02
	punktyWykresuFunkcji[602].Y = -139.252

	punktyWykresuFunkcji[603].X = 6.03
	punktyWykresuFunkcji[603].Y = -139.981

	punktyWykresuFunkcji[604].X = 6.04
	punktyWykresuFunkcji[604].Y = -140.712

	punktyWykresuFunkcji[605].X = 6.05
	punktyWykresuFunkcji[605].Y = -141.444

	punktyWykresuFunkcji[606].X = 6.06
	punktyWykresuFunkcji[606].Y = -142.179

	punktyWykresuFunkcji[607].X = 6.07
	punktyWykresuFunkcji[607].Y = -142.915

	punktyWykresuFunkcji[608].X = 6.08
	punktyWykresuFunkcji[608].Y = -143.654

	punktyWykresuFunkcji[609].X = 6.09
	punktyWykresuFunkcji[609].Y = -144.394

	punktyWykresuFunkcji[610].X = 6.10
	punktyWykresuFunkcji[610].Y = -145.136

	punktyWykresuFunkcji[611].X = 6.11
	punktyWykresuFunkcji[611].Y = -145.881

	punktyWykresuFunkcji[612].X = 6.12
	punktyWykresuFunkcji[612].Y = -146.627

	punktyWykresuFunkcji[613].X = 6.13
	punktyWykresuFunkcji[613].Y = -147.375

	punktyWykresuFunkcji[614].X = 6.14
	punktyWykresuFunkcji[614].Y = -148.125

	punktyWykresuFunkcji[615].X = 6.15
	punktyWykresuFunkcji[615].Y = -148.877

	punktyWykresuFunkcji[616].X = 6.16
	punktyWykresuFunkcji[616].Y = -149.631

	punktyWykresuFunkcji[617].X = 6.17
	punktyWykresuFunkcji[617].Y = -150.387

	punktyWykresuFunkcji[618].X = 6.18
	punktyWykresuFunkcji[618].Y = -151.145

	punktyWykresuFunkcji[619].X = 6.19
	punktyWykresuFunkcji[619].Y = -151.905

	punktyWykresuFunkcji[620].X = 6.2
	punktyWykresuFunkcji[620].Y = -152.667

	punktyWykresuFunkcji[621].X = 6.21
	punktyWykresuFunkcji[621].Y = -153.431

	punktyWykresuFunkcji[622].X = 6.22
	punktyWykresuFunkcji[622].Y = -154.197

	punktyWykresuFunkcji[623].X = 6.23
	punktyWykresuFunkcji[623].Y = -154.965

	punktyWykresuFunkcji[624].X = 6.24
	punktyWykresuFunkcji[624].Y = -155.735

	punktyWykresuFunkcji[625].X = 6.25
	punktyWykresuFunkcji[625].Y = -156.507

	punktyWykresuFunkcji[626].X = 6.26
	punktyWykresuFunkcji[626].Y = -157.281

	punktyWykresuFunkcji[627].X = 6.27
	punktyWykresuFunkcji[627].Y = -158.057

	punktyWykresuFunkcji[628].X = 6.28
	punktyWykresuFunkcji[628].Y = -158.835

	punktyWykresuFunkcji[629].X = 6.29
	punktyWykresuFunkcji[629].Y = -159.615

	punktyWykresuFunkcji[630].X = 6.3
	punktyWykresuFunkcji[630].Y = -160.397

	punktyWykresuFunkcji[631].X = 6.31
	punktyWykresuFunkcji[631].Y = -161.181

	punktyWykresuFunkcji[632].X = 6.32
	punktyWykresuFunkcji[632].Y = -161.967

	punktyWykresuFunkcji[633].X = 6.33
	punktyWykresuFunkcji[633].Y = -162.755

	punktyWykresuFunkcji[634].X = 6.34
	punktyWykresuFunkcji[634].Y = -163.545

	punktyWykresuFunkcji[635].X = 6.35
	punktyWykresuFunkcji[635].Y = -164.337

	punktyWykresuFunkcji[636].X = 6.36
	punktyWykresuFunkcji[636].Y = -165.131

	punktyWykresuFunkcji[637].X = 6.37
	punktyWykresuFunkcji[637].Y = -165.927

	punktyWykresuFunkcji[638].X = 6.38
	punktyWykresuFunkcji[638].Y = -166.725

	punktyWykresuFunkcji[639].X = 6.39
	punktyWykresuFunkcji[639].Y = -167.526

	punktyWykresuFunkcji[640].X = 6.4
	punktyWykresuFunkcji[640].Y = -168.328

	punktyWykresuFunkcji[641].X = 6.41
	punktyWykresuFunkcji[641].Y = -169.132

	punktyWykresuFunkcji[642].X = 6.42
	punktyWykresuFunkcji[642].Y = -169.939

	punktyWykresuFunkcji[643].X = 6.43
	punktyWykresuFunkcji[643].Y = -170.747

	punktyWykresuFunkcji[644].X = 6.44
	punktyWykresuFunkcji[644].Y = -171.747

	punktyWykresuFunkcji[645].X = 6.45
	punktyWykresuFunkcji[645].Y = -172.371

	punktyWykresuFunkcji[646].X = 6.46
	punktyWykresuFunkcji[646].Y = -173.186

	punktyWykresuFunkcji[647].X = 6.47
	punktyWykresuFunkcji[647].Y = -174.002

	punktyWykresuFunkcji[648].X = 6.48
	punktyWykresuFunkcji[648].Y = -174.821

	punktyWykresuFunkcji[649].X = 6.49
	punktyWykresuFunkcji[649].Y = -175.642

	punktyWykresuFunkcji[650].X = 6.5
	punktyWykresuFunkcji[650].Y = -176.466

	punktyWykresuFunkcji[651].X = 6.51
	punktyWykresuFunkcji[651].Y = -177.291

	punktyWykresuFunkcji[652].X = 6.52
	punktyWykresuFunkcji[652].Y = -178.118

	punktyWykresuFunkcji[653].X = 6.53
	punktyWykresuFunkcji[653].Y = -178.948

	punktyWykresuFunkcji[654].X = 6.54
	punktyWykresuFunkcji[654].Y = -179.78

	punktyWykresuFunkcji[655].X = 6.55
	punktyWykresuFunkcji[655].Y = -180.613

	punktyWykresuFunkcji[656].X = 6.56
	punktyWykresuFunkcji[656].Y = -181.449

	punktyWykresuFunkcji[657].X = 6.57
	punktyWykresuFunkcji[657].Y = -182.287

	punktyWykresuFunkcji[658].X = 6.58
	punktyWykresuFunkcji[658].Y = -183.128

	punktyWykresuFunkcji[659].X = 6.59
	punktyWykresuFunkcji[659].Y = -183.97

	punktyWykresuFunkcji[660].X = 6.6
	punktyWykresuFunkcji[660].Y = -184.815

	punktyWykresuFunkcji[661].X = 6.61
	punktyWykresuFunkcji[661].Y = -185.661

	punktyWykresuFunkcji[662].X = 6.62
	punktyWykresuFunkcji[662].Y = -186.51

	punktyWykresuFunkcji[663].X = 6.63
	punktyWykresuFunkcji[663].Y = -187.361

	punktyWykresuFunkcji[664].X = 6.64
	punktyWykresuFunkcji[664].Y = -188.214

	punktyWykresuFunkcji[665].X = 6.65
	punktyWykresuFunkcji[665].Y = -189.07

	punktyWykresuFunkcji[666].X = 6.66
	punktyWykresuFunkcji[666].Y = -189.927

	punktyWykresuFunkcji[667].X = 6.67
	punktyWykresuFunkcji[667].Y = -190.787

	punktyWykresuFunkcji[668].X = 6.68
	punktyWykresuFunkcji[668].Y = -191.649

	punktyWykresuFunkcji[669].X = 6.69
	punktyWykresuFunkcji[669].Y = -192.513

	punktyWykresuFunkcji[670].X = 6.7
	punktyWykresuFunkcji[670].Y = -193.38

	punktyWykresuFunkcji[671].X = 6.71
	punktyWykresuFunkcji[671].Y = -194.249

	punktyWykresuFunkcji[672].X = 6.72
	punktyWykresuFunkcji[672].Y = -195.119

	punktyWykresuFunkcji[673].X = 6.73
	punktyWykresuFunkcji[673].Y = -195.993

	punktyWykresuFunkcji[674].X = 6.74
	punktyWykresuFunkcji[674].Y = -196.868

	punktyWykresuFunkcji[675].X = 6.75
	punktyWykresuFunkcji[675].Y = -197.746

	punktyWykresuFunkcji[676].X = 6.76
	punktyWykresuFunkcji[676].Y = -198.626

	punktyWykresuFunkcji[677].X = 6.77
	punktyWykresuFunkcji[677].Y = -199.508

	punktyWykresuFunkcji[678].X = 6.78
	punktyWykresuFunkcji[678].Y = -200.392

	punktyWykresuFunkcji[679].X = 6.79
	punktyWykresuFunkcji[679].Y = -201.279

	punktyWykresuFunkcji[680].X = 6.8
	punktyWykresuFunkcji[680].Y = -202.168

	punktyWykresuFunkcji[681].X = 6.81
	punktyWykresuFunkcji[681].Y = -203.059

	punktyWykresuFunkcji[682].X = 6.82
	punktyWykresuFunkcji[682].Y = -203.952

	punktyWykresuFunkcji[683].X = 6.83
	punktyWykresuFunkcji[683].Y = -204.848

	punktyWykresuFunkcji[684].X = 6.84
	punktyWykresuFunkcji[684].Y = -205.746

	punktyWykresuFunkcji[685].X = 6.85
	punktyWykresuFunkcji[685].Y = -206.647

	punktyWykresuFunkcji[686].X = 6.86
	punktyWykresuFunkcji[686].Y = -207.55

	punktyWykresuFunkcji[687].X = 6.87
	punktyWykresuFunkcji[687].Y = -208.455

	punktyWykresuFunkcji[688].X = 6.88
	punktyWykresuFunkcji[688].Y = -209.362

	punktyWykresuFunkcji[689].X = 6.89
	punktyWykresuFunkcji[689].Y = -210.272

	punktyWykresuFunkcji[690].X = 6.9
	punktyWykresuFunkcji[690].Y = -211.184

	punktyWykresuFunkcji[691].X = 6.91
	punktyWykresuFunkcji[691].Y = -212.098

	punktyWykresuFunkcji[692].X = 6.92
	punktyWykresuFunkcji[692].Y = -213.015

	punktyWykresuFunkcji[693].X = 6.93
	punktyWykresuFunkcji[693].Y = -213.934

	punktyWykresuFunkcji[694].X = 6.94
	punktyWykresuFunkcji[694].Y = -214.856

	punktyWykresuFunkcji[695].X = 6.95
	punktyWykresuFunkcji[695].Y = -215.78

	punktyWykresuFunkcji[696].X = 6.96
	punktyWykresuFunkcji[696].Y = -216.706

	punktyWykresuFunkcji[697].X = 6.97
	punktyWykresuFunkcji[697].Y = -217.635

	punktyWykresuFunkcji[698].X = 6.98
	punktyWykresuFunkcji[698].Y = -218.566

	punktyWykresuFunkcji[699].X = 6.99
	punktyWykresuFunkcji[699].Y = -219.5

	punktyWykresuFunkcji[700].X = 7.0
	punktyWykresuFunkcji[700].Y = -220.435

	punktyWykresuFunkcji[701].X = 7.01
	punktyWykresuFunkcji[701].Y = -221.374

	punktyWykresuFunkcji[702].X = 7.02
	punktyWykresuFunkcji[702].Y = -222.315

	punktyWykresuFunkcji[703].X = 7.03
	punktyWykresuFunkcji[703].Y = -223.258

	punktyWykresuFunkcji[704].X = 7.04
	punktyWykresuFunkcji[704].Y = -224.204

	punktyWykresuFunkcji[705].X = 7.05
	punktyWykresuFunkcji[705].Y = -225.152

	punktyWykresuFunkcji[706].X = 7.06
	punktyWykresuFunkcji[706].Y = -226.102

	punktyWykresuFunkcji[707].X = 7.07
	punktyWykresuFunkcji[707].Y = -227.055

	punktyWykresuFunkcji[708].X = 7.08
	punktyWykresuFunkcji[708].Y = -228.011

	punktyWykresuFunkcji[709].X = 7.09
	punktyWykresuFunkcji[709].Y = -228.969

	punktyWykresuFunkcji[710].X = 7.1
	punktyWykresuFunkcji[710].Y = -229.929

	punktyWykresuFunkcji[711].X = 7.11
	punktyWykresuFunkcji[711].Y = -230.892

	punktyWykresuFunkcji[712].X = 7.12
	punktyWykresuFunkcji[712].Y = -231.858

	punktyWykresuFunkcji[713].X = 7.13
	punktyWykresuFunkcji[713].Y = -232.826

	punktyWykresuFunkcji[714].X = 7.14
	punktyWykresuFunkcji[714].Y = -233.796

	punktyWykresuFunkcji[715].X = 7.15
	punktyWykresuFunkcji[715].Y = -234.77

	punktyWykresuFunkcji[716].X = 7.16
	punktyWykresuFunkcji[716].Y = -235.745

	punktyWykresuFunkcji[717].X = 7.17
	punktyWykresuFunkcji[717].Y = -236.723

	punktyWykresuFunkcji[718].X = 7.18
	punktyWykresuFunkcji[718].Y = -237.704

	punktyWykresuFunkcji[719].X = 7.19
	punktyWykresuFunkcji[719].Y = -238.687

	punktyWykresuFunkcji[720].X = 7.2
	punktyWykresuFunkcji[720].Y = -239.673

	punktyWykresuFunkcji[721].X = 7.21
	punktyWykresuFunkcji[721].Y = -240.661

	punktyWykresuFunkcji[722].X = 7.22
	punktyWykresuFunkcji[722].Y = -241.652

	punktyWykresuFunkcji[723].X = 7.23
	punktyWykresuFunkcji[723].Y = -242.646

	punktyWykresuFunkcji[724].X = 7.24
	punktyWykresuFunkcji[724].Y = -243.642

	punktyWykresuFunkcji[725].X = 7.25
	punktyWykresuFunkcji[725].Y = -244.641

	punktyWykresuFunkcji[726].X = 7.26
	punktyWykresuFunkcji[726].Y = -245.642

	punktyWykresuFunkcji[727].X = 7.27
	punktyWykresuFunkcji[727].Y = -246.646

	punktyWykresuFunkcji[728].X = 7.28
	punktyWykresuFunkcji[728].Y = -247.653

	punktyWykresuFunkcji[729].X = 7.29
	punktyWykresuFunkcji[729].Y = -248.662

	punktyWykresuFunkcji[730].X = 7.3
	punktyWykresuFunkcji[730].Y = -249.674

	punktyWykresuFunkcji[731].X = 7.31
	punktyWykresuFunkcji[731].Y = -250.689

	punktyWykresuFunkcji[732].X = 7.32
	punktyWykresuFunkcji[732].Y = -251.706

	punktyWykresuFunkcji[733].X = 7.33
	punktyWykresuFunkcji[733].Y = -252.726

	punktyWykresuFunkcji[734].X = 7.34
	punktyWykresuFunkcji[734].Y = -253.749

	punktyWykresuFunkcji[735].X = 7.35
	punktyWykresuFunkcji[735].Y = -254.774

	punktyWykresuFunkcji[736].X = 7.36
	punktyWykresuFunkcji[736].Y = -255.802

	punktyWykresuFunkcji[737].X = 7.37
	punktyWykresuFunkcji[737].Y = -256.833

	punktyWykresuFunkcji[738].X = 7.38
	punktyWykresuFunkcji[738].Y = -257.866

	punktyWykresuFunkcji[739].X = 7.39
	punktyWykresuFunkcji[739].Y = -258.902

	punktyWykresuFunkcji[740].X = 7.4
	punktyWykresuFunkcji[740].Y = -259.941

	punktyWykresuFunkcji[741].X = 7.41
	punktyWykresuFunkcji[741].Y = -260.983

	punktyWykresuFunkcji[742].X = 7.42
	punktyWykresuFunkcji[742].Y = -262.027

	punktyWykresuFunkcji[743].X = 7.43
	punktyWykresuFunkcji[743].Y = -263.075

	punktyWykresuFunkcji[744].X = 7.44
	punktyWykresuFunkcji[744].Y = -264.124

	punktyWykresuFunkcji[745].X = 7.45
	punktyWykresuFunkcji[745].Y = -265.177

	punktyWykresuFunkcji[746].X = 7.46
	punktyWykresuFunkcji[746].Y = -266.233

	punktyWykresuFunkcji[747].X = 7.47
	punktyWykresuFunkcji[747].Y = -267.291

	punktyWykresuFunkcji[748].X = 7.48
	punktyWykresuFunkcji[748].Y = -268.352

	punktyWykresuFunkcji[749].X = 7.49
	punktyWykresuFunkcji[749].Y = -269.416

	punktyWykresuFunkcji[750].X = 7.5
	punktyWykresuFunkcji[750].Y = -270.482

	punktyWykresuFunkcji[751].X = 7.51
	punktyWykresuFunkcji[751].Y = -271.552

	punktyWykresuFunkcji[752].X = 7.52
	punktyWykresuFunkcji[752].Y = -272.624

	punktyWykresuFunkcji[753].X = 7.53
	punktyWykresuFunkcji[753].Y = -273.7

	punktyWykresuFunkcji[754].X = 7.54
	punktyWykresuFunkcji[754].Y = -274.778

	punktyWykresuFunkcji[755].X = 7.55
	punktyWykresuFunkcji[755].Y = -275.858

	punktyWykresuFunkcji[756].X = 7.56
	punktyWykresuFunkcji[756].Y = -276.942

	punktyWykresuFunkcji[757].X = 7.57
	punktyWykresuFunkcji[757].Y = -278.029

	punktyWykresuFunkcji[758].X = 7.58
	punktyWykresuFunkcji[758].Y = -279.118

	punktyWykresuFunkcji[759].X = 7.59
	punktyWykresuFunkcji[759].Y = -280.211

	punktyWykresuFunkcji[760].X = 7.6
	punktyWykresuFunkcji[760].Y = -281.306

	punktyWykresuFunkcji[761].X = 7.61
	punktyWykresuFunkcji[761].Y = -282.404

	punktyWykresuFunkcji[762].X = 7.62
	punktyWykresuFunkcji[762].Y = -283.506

	punktyWykresuFunkcji[763].X = 7.63
	punktyWykresuFunkcji[763].Y = -284.61

	punktyWykresuFunkcji[764].X = 7.64
	punktyWykresuFunkcji[764].Y = -285.717

	punktyWykresuFunkcji[765].X = 7.65
	punktyWykresuFunkcji[765].Y = -286.827

	punktyWykresuFunkcji[766].X = 7.66
	punktyWykresuFunkcji[766].Y = -287.94

	punktyWykresuFunkcji[767].X = 7.67
	punktyWykresuFunkcji[767].Y = -289.056

	punktyWykresuFunkcji[768].X = 7.68
	punktyWykresuFunkcji[768].Y = -290.175

	punktyWykresuFunkcji[769].X = 7.69
	punktyWykresuFunkcji[769].Y = -291.297

	punktyWykresuFunkcji[770].X = 7.7
	punktyWykresuFunkcji[770].Y = -292.421

	punktyWykresuFunkcji[771].X = 7.71
	punktyWykresuFunkcji[771].Y = -293.549

	punktyWykresuFunkcji[772].X = 7.72
	punktyWykresuFunkcji[772].Y = -294.68

	punktyWykresuFunkcji[773].X = 7.73
	punktyWykresuFunkcji[773].Y = -295.814

	punktyWykresuFunkcji[774].X = 7.74
	punktyWykresuFunkcji[774].Y = -296.951

	punktyWykresuFunkcji[775].X = 7.75
	punktyWykresuFunkcji[775].Y = -298.091

	punktyWykresuFunkcji[776].X = 7.76
	punktyWykresuFunkcji[776].Y = -299.234

	punktyWykresuFunkcji[777].X = 7.77
	punktyWykresuFunkcji[777].Y = -300.38

	punktyWykresuFunkcji[778].X = 7.78
	punktyWykresuFunkcji[778].Y = -301.529

	punktyWykresuFunkcji[779].X = 7.79
	punktyWykresuFunkcji[779].Y = -302.682

	punktyWykresuFunkcji[780].X = 7.8
	punktyWykresuFunkcji[780].Y = -303.837

	punktyWykresuFunkcji[781].X = 7.81
	punktyWykresuFunkcji[781].Y = -304.995

	punktyWykresuFunkcji[782].X = 7.82
	punktyWykresuFunkcji[782].Y = -306.157

	punktyWykresuFunkcji[783].X = 7.83
	punktyWykresuFunkcji[783].Y = -307.322

	punktyWykresuFunkcji[784].X = 7.84
	punktyWykresuFunkcji[784].Y = -308.489

	punktyWykresuFunkcji[785].X = 7.85
	punktyWykresuFunkcji[785].Y = -309.66

	punktyWykresuFunkcji[786].X = 7.86
	punktyWykresuFunkcji[786].Y = -310.834

	punktyWykresuFunkcji[787].X = 7.87
	punktyWykresuFunkcji[787].Y = -312.011

	punktyWykresuFunkcji[788].X = 7.88
	punktyWykresuFunkcji[788].Y = -313.192

	punktyWykresuFunkcji[789].X = 7.89
	punktyWykresuFunkcji[789].Y = -314.375

	punktyWykresuFunkcji[790].X = 7.9
	punktyWykresuFunkcji[790].Y = -315.562

	punktyWykresuFunkcji[791].X = 7.91
	punktyWykresuFunkcji[791].Y = -316.752

	punktyWykresuFunkcji[792].X = 7.92
	punktyWykresuFunkcji[792].Y = -317.945

	punktyWykresuFunkcji[793].X = 7.93
	punktyWykresuFunkcji[793].Y = -319.141

	punktyWykresuFunkcji[794].X = 7.94
	punktyWykresuFunkcji[794].Y = -320.34

	punktyWykresuFunkcji[795].X = 7.95
	punktyWykresuFunkcji[795].Y = -321.543

	punktyWykresuFunkcji[796].X = 7.96
	punktyWykresuFunkcji[796].Y = -322.749

	punktyWykresuFunkcji[797].X = 7.97
	punktyWykresuFunkcji[797].Y = -323.958

	punktyWykresuFunkcji[798].X = 7.98
	punktyWykresuFunkcji[798].Y = -325.171

	punktyWykresuFunkcji[799].X = 7.99
	punktyWykresuFunkcji[799].Y = -326.386

	punktyWykresuFunkcji[800].X = 8.0
	punktyWykresuFunkcji[800].Y = -327.605

	punktyWykresuFunkcji[801].X = 8.01
	punktyWykresuFunkcji[801].Y = -328.827

	punktyWykresuFunkcji[802].X = 8.02
	punktyWykresuFunkcji[802].Y = -330.053

	punktyWykresuFunkcji[803].X = 8.03
	punktyWykresuFunkcji[803].Y = -331.281

	punktyWykresuFunkcji[804].X = 8.04
	punktyWykresuFunkcji[804].Y = -332.513

	punktyWykresuFunkcji[805].X = 8.05
	punktyWykresuFunkcji[805].Y = -333.749

	punktyWykresuFunkcji[806].X = 8.06
	punktyWykresuFunkcji[806].Y = -334.987

	punktyWykresuFunkcji[807].X = 8.07
	punktyWykresuFunkcji[807].Y = -336.229

	punktyWykresuFunkcji[808].X = 8.08
	punktyWykresuFunkcji[808].Y = -337.475

	punktyWykresuFunkcji[809].X = 8.09
	punktyWykresuFunkcji[809].Y = -338.723

	punktyWykresuFunkcji[810].X = 8.1
	punktyWykresuFunkcji[810].Y = -339.975

	punktyWykresuFunkcji[811].X = 8.11
	punktyWykresuFunkcji[811].Y = -341.231

	punktyWykresuFunkcji[812].X = 8.12
	punktyWykresuFunkcji[812].Y = -342.49

	punktyWykresuFunkcji[813].X = 8.13
	punktyWykresuFunkcji[813].Y = -343.752

	punktyWykresuFunkcji[814].X = 8.14
	punktyWykresuFunkcji[814].Y = -345.017

	punktyWykresuFunkcji[815].X = 8.15
	punktyWykresuFunkcji[815].Y = -346.286

	punktyWykresuFunkcji[816].X = 8.16
	punktyWykresuFunkcji[816].Y = -347.559

	punktyWykresuFunkcji[817].X = 8.17
	punktyWykresuFunkcji[817].Y = -348.834

	punktyWykresuFunkcji[818].X = 8.18
	punktyWykresuFunkcji[818].Y = -350.113

	punktyWykresuFunkcji[819].X = 8.19
	punktyWykresuFunkcji[819].Y = -351.396

	punktyWykresuFunkcji[820].X = 8.2
	punktyWykresuFunkcji[820].Y = -352.682

	punktyWykresuFunkcji[821].X = 8.21
	punktyWykresuFunkcji[821].Y = -353.972

	punktyWykresuFunkcji[822].X = 8.22
	punktyWykresuFunkcji[822].Y = -355.265

	punktyWykresuFunkcji[823].X = 8.23
	punktyWykresuFunkcji[823].Y = -356.561

	punktyWykresuFunkcji[824].X = 8.24
	punktyWykresuFunkcji[824].Y = -357.861

	punktyWykresuFunkcji[825].X = 8.25
	punktyWykresuFunkcji[825].Y = -359.164

	punktyWykresuFunkcji[826].X = 8.26
	punktyWykresuFunkcji[826].Y = -360.471

	punktyWykresuFunkcji[827].X = 8.27
	punktyWykresuFunkcji[827].Y = -361.781

	punktyWykresuFunkcji[828].X = 8.28
	punktyWykresuFunkcji[828].Y = -363.095

	punktyWykresuFunkcji[829].X = 8.29
	punktyWykresuFunkcji[829].Y = -364.413

	punktyWykresuFunkcji[830].X = 8.3
	punktyWykresuFunkcji[830].Y = -365.734

	punktyWykresuFunkcji[831].X = 8.31
	punktyWykresuFunkcji[831].Y = -367.058

	punktyWykresuFunkcji[832].X = 8.32
	punktyWykresuFunkcji[832].Y = -368.386

	punktyWykresuFunkcji[833].X = 8.33
	punktyWykresuFunkcji[833].Y = -369.718

	punktyWykresuFunkcji[834].X = 8.34
	punktyWykresuFunkcji[834].Y = -371.053

	punktyWykresuFunkcji[835].X = 8.35
	punktyWykresuFunkcji[835].Y = -372.391

	punktyWykresuFunkcji[836].X = 8.36
	punktyWykresuFunkcji[836].Y = -373.734

	punktyWykresuFunkcji[837].X = 8.37
	punktyWykresuFunkcji[837].Y = -375.08

	punktyWykresuFunkcji[838].X = 8.38
	punktyWykresuFunkcji[838].Y = -376.429

	punktyWykresuFunkcji[839].X = 8.39
	punktyWykresuFunkcji[839].Y = -377.782

	punktyWykresuFunkcji[840].X = 8.4
	punktyWykresuFunkcji[840].Y = -379.139

	punktyWykresuFunkcji[841].X = 8.41
	punktyWykresuFunkcji[841].Y = -380.499

	punktyWykresuFunkcji[842].X = 8.42
	punktyWykresuFunkcji[842].Y = -381.863

	punktyWykresuFunkcji[843].X = 8.43
	punktyWykresuFunkcji[843].Y = -383.23

	punktyWykresuFunkcji[844].X = 8.44
	punktyWykresuFunkcji[844].Y = -384.602

	punktyWykresuFunkcji[845].X = 8.45
	punktyWykresuFunkcji[845].Y = -385.976

	punktyWykresuFunkcji[846].X = 8.46
	punktyWykresuFunkcji[846].Y = -387.355

	punktyWykresuFunkcji[847].X = 8.47
	punktyWykresuFunkcji[847].Y = -388.737

	punktyWykresuFunkcji[848].X = 8.48
	punktyWykresuFunkcji[848].Y = -390.123

	punktyWykresuFunkcji[849].X = 8.49
	punktyWykresuFunkcji[849].Y = -391.513

	punktyWykresuFunkcji[850].X = 8.5
	punktyWykresuFunkcji[850].Y = -392.906

	punktyWykresuFunkcji[851].X = 8.51
	punktyWykresuFunkcji[851].Y = -394.303

	punktyWykresuFunkcji[852].X = 8.52
	punktyWykresuFunkcji[852].Y = -395.703

	punktyWykresuFunkcji[853].X = 8.53
	punktyWykresuFunkcji[853].Y = -397.108

	punktyWykresuFunkcji[854].X = 8.54
	punktyWykresuFunkcji[854].Y = -398.516

	punktyWykresuFunkcji[855].X = 8.55
	punktyWykresuFunkcji[855].Y = -399.928

	punktyWykresuFunkcji[856].X = 8.56
	punktyWykresuFunkcji[856].Y = -401.343

	punktyWykresuFunkcji[857].X = 8.57
	punktyWykresuFunkcji[857].Y = -402.762

	punktyWykresuFunkcji[858].X = 8.58
	punktyWykresuFunkcji[858].Y = -404.185

	punktyWykresuFunkcji[859].X = 8.59
	punktyWykresuFunkcji[859].Y = -405.612

	punktyWykresuFunkcji[860].X = 8.6
	punktyWykresuFunkcji[860].Y = -407.043

	punktyWykresuFunkcji[861].X = 8.61
	punktyWykresuFunkcji[861].Y = -408.477

	punktyWykresuFunkcji[862].X = 8.62
	punktyWykresuFunkcji[862].Y = -409.915

	punktyWykresuFunkcji[863].X = 8.63
	punktyWykresuFunkcji[863].Y = -411.357

	punktyWykresuFunkcji[864].X = 8.64
	punktyWykresuFunkcji[864].Y = -412.803

	punktyWykresuFunkcji[865].X = 8.65
	punktyWykresuFunkcji[865].Y = -414.253

	punktyWykresuFunkcji[866].X = 8.66
	punktyWykresuFunkcji[866].Y = -415.706

	punktyWykresuFunkcji[867].X = 8.67
	punktyWykresuFunkcji[867].Y = -417.163

	punktyWykresuFunkcji[868].X = 8.68
	punktyWykresuFunkcji[868].Y = -418.624

	punktyWykresuFunkcji[869].X = 8.69
	punktyWykresuFunkcji[869].Y = -420.089

	punktyWykresuFunkcji[870].X = 8.7
	punktyWykresuFunkcji[870].Y = -421.558

	punktyWykresuFunkcji[871].X = 8.71
	punktyWykresuFunkcji[871].Y = -423.03

	punktyWykresuFunkcji[872].X = 8.72
	punktyWykresuFunkcji[872].Y = -424.507

	punktyWykresuFunkcji[873].X = 8.73
	punktyWykresuFunkcji[873].Y = -425.987

	punktyWykresuFunkcji[874].X = 8.74
	punktyWykresuFunkcji[874].Y = -427.471

	punktyWykresuFunkcji[875].X = 8.75
	punktyWykresuFunkcji[875].Y = -428.959

	punktyWykresuFunkcji[876].X = 8.76
	punktyWykresuFunkcji[876].Y = -430.451

	punktyWykresuFunkcji[877].X = 8.77
	punktyWykresuFunkcji[877].Y = -431.947

	punktyWykresuFunkcji[878].X = 8.78
	punktyWykresuFunkcji[878].Y = -433.447

	punktyWykresuFunkcji[879].X = 8.79
	punktyWykresuFunkcji[879].Y = -434.951

	punktyWykresuFunkcji[880].X = 8.8
	punktyWykresuFunkcji[880].Y = -436.458

	punktyWykresuFunkcji[881].X = 8.81
	punktyWykresuFunkcji[881].Y = -437.97

	punktyWykresuFunkcji[882].X = 8.82
	punktyWykresuFunkcji[882].Y = -439.485

	punktyWykresuFunkcji[883].X = 8.83
	punktyWykresuFunkcji[883].Y = -441.005

	punktyWykresuFunkcji[884].X = 8.84
	punktyWykresuFunkcji[884].Y = -442.528

	punktyWykresuFunkcji[885].X = 8.85
	punktyWykresuFunkcji[885].Y = -444.055

	punktyWykresuFunkcji[886].X = 8.86
	punktyWykresuFunkcji[886].Y = -445.587

	punktyWykresuFunkcji[887].X = 8.87
	punktyWykresuFunkcji[887].Y = -447.122

	punktyWykresuFunkcji[888].X = 8.88
	punktyWykresuFunkcji[888].Y = -448.661

	punktyWykresuFunkcji[889].X = 8.89
	punktyWykresuFunkcji[889].Y = -450.204

	punktyWykresuFunkcji[890].X = 8.9
	punktyWykresuFunkcji[890].Y = -451.751

	punktyWykresuFunkcji[891].X = 8.91
	punktyWykresuFunkcji[891].Y = -453.302

	punktyWykresuFunkcji[892].X = 8.92
	punktyWykresuFunkcji[892].Y = -454.858

	punktyWykresuFunkcji[893].X = 8.93
	punktyWykresuFunkcji[893].Y = -456.417

	punktyWykresuFunkcji[894].X = 8.94
	punktyWykresuFunkcji[894].Y = -457.98

	punktyWykresuFunkcji[895].X = 8.95
	punktyWykresuFunkcji[895].Y = -459.547

	punktyWykresuFunkcji[896].X = 8.96
	punktyWykresuFunkcji[896].Y = -461.118

	punktyWykresuFunkcji[897].X = 8.97
	punktyWykresuFunkcji[897].Y = -462.693

	punktyWykresuFunkcji[898].X = 8.98
	punktyWykresuFunkcji[898].Y = -464.273

	punktyWykresuFunkcji[899].X = 8.99
	punktyWykresuFunkcji[899].Y = -465.856

	punktyWykresuFunkcji[900].X = 9.0
	punktyWykresuFunkcji[900].Y = -467.443

	punktyWykresuFunkcji[901].X = 9.01
	punktyWykresuFunkcji[901].Y = -469.035

	punktyWykresuFunkcji[902].X = 9.02
	punktyWykresuFunkcji[902].Y = -470.63

	punktyWykresuFunkcji[903].X = 9.03
	punktyWykresuFunkcji[903].Y = -472.23

	punktyWykresuFunkcji[904].X = 9.04
	punktyWykresuFunkcji[904].Y = -473.833

	punktyWykresuFunkcji[905].X = 9.05
	punktyWykresuFunkcji[905].Y = -475.441

	punktyWykresuFunkcji[906].X = 9.06
	punktyWykresuFunkcji[906].Y = -477.053

	punktyWykresuFunkcji[907].X = 9.07
	punktyWykresuFunkcji[907].Y = -478.669

	punktyWykresuFunkcji[908].X = 9.08
	punktyWykresuFunkcji[908].Y = -480.289

	punktyWykresuFunkcji[909].X = 9.09
	punktyWykresuFunkcji[909].Y = -481.913

	punktyWykresuFunkcji[910].X = 9.1
	punktyWykresuFunkcji[910].Y = -483.541

	punktyWykresuFunkcji[911].X = 9.11
	punktyWykresuFunkcji[911].Y = -485.173

	punktyWykresuFunkcji[912].X = 9.12
	punktyWykresuFunkcji[912].Y = -486.81

	punktyWykresuFunkcji[913].X = 9.13
	punktyWykresuFunkcji[913].Y = -488.45

	punktyWykresuFunkcji[914].X = 9.14
	punktyWykresuFunkcji[914].Y = -490.095

	punktyWykresuFunkcji[915].X = 9.15
	punktyWykresuFunkcji[915].Y = -491.744

	punktyWykresuFunkcji[916].X = 9.16
	punktyWykresuFunkcji[916].Y = -493.397

	punktyWykresuFunkcji[917].X = 9.17
	punktyWykresuFunkcji[917].Y = -495.054

	punktyWykresuFunkcji[918].X = 9.18
	punktyWykresuFunkcji[918].Y = -496.715

	punktyWykresuFunkcji[919].X = 9.19
	punktyWykresuFunkcji[919].Y = -498.381

	punktyWykresuFunkcji[920].X = 9.2
	punktyWykresuFunkcji[920].Y = -500.05

	punktyWykresuFunkcji[921].X = 9.21
	punktyWykresuFunkcji[921].Y = -501.724

	punktyWykresuFunkcji[922].X = 9.22
	punktyWykresuFunkcji[922].Y = -503.402

	punktyWykresuFunkcji[923].X = 9.23
	punktyWykresuFunkcji[923].Y = -505.084

	punktyWykresuFunkcji[924].X = 9.24
	punktyWykresuFunkcji[924].Y = -506.77

	punktyWykresuFunkcji[925].X = 9.25
	punktyWykresuFunkcji[925].Y = -508.461

	punktyWykresuFunkcji[926].X = 9.26
	punktyWykresuFunkcji[926].Y = -510.155

	punktyWykresuFunkcji[927].X = 9.27
	punktyWykresuFunkcji[927].Y = -511.854

	punktyWykresuFunkcji[928].X = 9.28
	punktyWykresuFunkcji[928].Y = -513.557

	punktyWykresuFunkcji[929].X = 9.29
	punktyWykresuFunkcji[929].Y = -515.264

	punktyWykresuFunkcji[930].X = 9.3
	punktyWykresuFunkcji[930].Y = -516.976

	punktyWykresuFunkcji[931].X = 9.31
	punktyWykresuFunkcji[931].Y = -518.692

	punktyWykresuFunkcji[932].X = 9.32
	punktyWykresuFunkcji[932].Y = -520.411

	punktyWykresuFunkcji[933].X = 9.33
	punktyWykresuFunkcji[933].Y = -522.136

	punktyWykresuFunkcji[934].X = 9.34
	punktyWykresuFunkcji[934].Y = -523.864

	punktyWykresuFunkcji[935].X = 9.35
	punktyWykresuFunkcji[935].Y = -525.597

	punktyWykresuFunkcji[936].X = 9.36
	punktyWykresuFunkcji[936].Y = -527.333

	punktyWykresuFunkcji[937].X = 9.37
	punktyWykresuFunkcji[937].Y = -529.074

	punktyWykresuFunkcji[938].X = 9.38
	punktyWykresuFunkcji[938].Y = -530.82

	punktyWykresuFunkcji[939].X = 9.39
	punktyWykresuFunkcji[939].Y = -532.569

	punktyWykresuFunkcji[940].X = 9.4
	punktyWykresuFunkcji[940].Y = -534.323

	punktyWykresuFunkcji[941].X = 9.41
	punktyWykresuFunkcji[941].Y = -536.081

	punktyWykresuFunkcji[942].X = 9.42
	punktyWykresuFunkcji[942].Y = -537.844

	punktyWykresuFunkcji[943].X = 9.43
	punktyWykresuFunkcji[943].Y = -539.61

	punktyWykresuFunkcji[944].X = 9.44
	punktyWykresuFunkcji[944].Y = -541.381

	punktyWykresuFunkcji[945].X = 9.45
	punktyWykresuFunkcji[945].Y = -543.156

	punktyWykresuFunkcji[946].X = 9.46
	punktyWykresuFunkcji[946].Y = -544.936

	punktyWykresuFunkcji[947].X = 9.47
	punktyWykresuFunkcji[947].Y = -546.72

	punktyWykresuFunkcji[948].X = 9.48
	punktyWykresuFunkcji[948].Y = -548.508

	punktyWykresuFunkcji[949].X = 9.49
	punktyWykresuFunkcji[949].Y = -550.3

	punktyWykresuFunkcji[950].X = 9.5
	punktyWykresuFunkcji[950].Y = -552.096

	punktyWykresuFunkcji[951].X = 9.51
	punktyWykresuFunkcji[951].Y = -553.897

	punktyWykresuFunkcji[952].X = 9.52
	punktyWykresuFunkcji[952].Y = -555.703

	punktyWykresuFunkcji[953].X = 9.53
	punktyWykresuFunkcji[953].Y = -557.512

	punktyWykresuFunkcji[954].X = 9.54
	punktyWykresuFunkcji[954].Y = -559.326

	punktyWykresuFunkcji[955].X = 9.55
	punktyWykresuFunkcji[955].Y = -561.144

	punktyWykresuFunkcji[956].X = 9.56
	punktyWykresuFunkcji[956].Y = -562.966

	punktyWykresuFunkcji[957].X = 9.57
	punktyWykresuFunkcji[957].Y = -564.793

	punktyWykresuFunkcji[958].X = 9.58
	punktyWykresuFunkcji[958].Y = -566.624

	punktyWykresuFunkcji[959].X = 9.59
	punktyWykresuFunkcji[959].Y = -568.46

	punktyWykresuFunkcji[960].X = 9.6
	punktyWykresuFunkcji[960].Y = -570.3

	punktyWykresuFunkcji[961].X = 9.61
	punktyWykresuFunkcji[961].Y = -572.144

	punktyWykresuFunkcji[962].X = 9.62
	punktyWykresuFunkcji[962].Y = -573.992

	punktyWykresuFunkcji[963].X = 9.63
	punktyWykresuFunkcji[963].Y = -575.845

	punktyWykresuFunkcji[964].X = 9.64
	punktyWykresuFunkcji[964].Y = -577.702

	punktyWykresuFunkcji[965].X = 9.65
	punktyWykresuFunkcji[965].Y = -579.563

	punktyWykresuFunkcji[966].X = 9.66
	punktyWykresuFunkcji[966].Y = -581.429

	punktyWykresuFunkcji[967].X = 9.67
	punktyWykresuFunkcji[967].Y = -583.299

	punktyWykresuFunkcji[968].X = 9.68
	punktyWykresuFunkcji[968].Y = -585.174

	punktyWykresuFunkcji[969].X = 9.69
	punktyWykresuFunkcji[969].Y = -587.053

	punktyWykresuFunkcji[970].X = 9.7
	punktyWykresuFunkcji[970].Y = -588.936

	punktyWykresuFunkcji[971].X = 9.71
	punktyWykresuFunkcji[971].Y = -590.823

	punktyWykresuFunkcji[972].X = 9.72
	punktyWykresuFunkcji[972].Y = -592.715

	punktyWykresuFunkcji[973].X = 9.73
	punktyWykresuFunkcji[973].Y = -594.612

	punktyWykresuFunkcji[974].X = 9.74
	punktyWykresuFunkcji[974].Y = -596.512

	punktyWykresuFunkcji[975].X = 9.75
	punktyWykresuFunkcji[975].Y = -598.417

	punktyWykresuFunkcji[976].X = 9.76
	punktyWykresuFunkcji[976].Y = -600.327

	punktyWykresuFunkcji[977].X = 9.77
	punktyWykresuFunkcji[977].Y = -602.24

	punktyWykresuFunkcji[978].X = 9.78
	punktyWykresuFunkcji[978].Y = -604.159

	punktyWykresuFunkcji[979].X = 9.79
	punktyWykresuFunkcji[979].Y = -606.081

	punktyWykresuFunkcji[980].X = 9.8
	punktyWykresuFunkcji[980].Y = -608.008

	punktyWykresuFunkcji[981].X = 9.81
	punktyWykresuFunkcji[981].Y = -609.939

	punktyWykresuFunkcji[982].X = 9.82
	punktyWykresuFunkcji[982].Y = -611.875

	punktyWykresuFunkcji[983].X = 9.83
	punktyWykresuFunkcji[983].Y = -613.815

	punktyWykresuFunkcji[984].X = 9.84
	punktyWykresuFunkcji[984].Y = -615.76

	punktyWykresuFunkcji[985].X = 9.85
	punktyWykresuFunkcji[985].Y = -617.709

	punktyWykresuFunkcji[986].X = 9.86
	punktyWykresuFunkcji[986].Y = -619.662

	punktyWykresuFunkcji[987].X = 9.87
	punktyWykresuFunkcji[987].Y = -621.62

	punktyWykresuFunkcji[988].X = 9.88
	punktyWykresuFunkcji[988].Y = -623.582

	punktyWykresuFunkcji[989].X = 9.89
	punktyWykresuFunkcji[989].Y = -625.548

	punktyWykresuFunkcji[990].X = 9.9
	punktyWykresuFunkcji[990].Y = -627.519

	punktyWykresuFunkcji[991].X = 9.91
	punktyWykresuFunkcji[991].Y = -629.494

	punktyWykresuFunkcji[992].X = 9.92
	punktyWykresuFunkcji[992].Y = -631.474

	punktyWykresuFunkcji[993].X = 9.93
	punktyWykresuFunkcji[993].Y = -633.458

	punktyWykresuFunkcji[994].X = 9.94
	punktyWykresuFunkcji[994].Y = -635.447

	punktyWykresuFunkcji[995].X = 9.95
	punktyWykresuFunkcji[995].Y = -637.439

	punktyWykresuFunkcji[996].X = 9.96
	punktyWykresuFunkcji[996].Y = -639.437

	punktyWykresuFunkcji[997].X = 9.97
	punktyWykresuFunkcji[997].Y = -641.438

	punktyWykresuFunkcji[998].X = 9.98
	punktyWykresuFunkcji[998].Y = -643.445

	punktyWykresuFunkcji[999].X = 9.99
	punktyWykresuFunkcji[999].Y = -645.455

	punktyWykresuFunkcji[1000].X = 10.0
	punktyWykresuFunkcji[1000].Y = -647.47




	wykresFunkcji := plot.New()

	wykresFunkcji.Title.Text = "Wykres funkcji f(x) = -(2/3)x^3 + x - 5cos(x) + 5"

	wykresFunkcji.X.Label.Text = "x"
	wykresFunkcji.Y.Label.Text = "y"

	liniaWykresu, err := plotter.NewLine(punktyWykresuFunkcji)

	if err != nil {
		panic(err)
	}

	liniaWykresu.LineStyle.Width = vg.Points(0.1)
	liniaWykresu.Color = color.RGBA{R: 200, G: 100, B: 100}

	wykresFunkcji.Add(liniaWykresu)
	wykresFunkcji.Legend.Add("f(x)", liniaWykresu)

	if err := wykresFunkcji.Save(10*vg.Inch, 10*vg.Inch,
		"Wykresy-funkcji-02.png"); err != nil {

		panic(err)
	}
}
