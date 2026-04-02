package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of Gamma function.

	pointsOfFunctionPlot := make(plotter.XYs, 701)

	pointsOfFunctionPlot[0].X = 0.001
	pointsOfFunctionPlot[0].Y = 999.423

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = 99.432

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = 49.442

	pointsOfFunctionPlot[3].X = 0.03
	pointsOfFunctionPlot[3].Y = 32.784

	pointsOfFunctionPlot[4].X = 0.04
	pointsOfFunctionPlot[4].Y = 24.46

	pointsOfFunctionPlot[5].X = 0.05
	pointsOfFunctionPlot[5].Y = 19.47

	pointsOfFunctionPlot[6].X = 0.06
	pointsOfFunctionPlot[6].Y = 16.145

	pointsOfFunctionPlot[7].X = 0.07
	pointsOfFunctionPlot[7].Y = 13.773

	pointsOfFunctionPlot[8].X = 0.08
	pointsOfFunctionPlot[8].Y = 11.996

	pointsOfFunctionPlot[9].X = 0.09
	pointsOfFunctionPlot[9].Y = 10.616

	pointsOfFunctionPlot[10].X = 0.10
	pointsOfFunctionPlot[10].Y = 9.513

	pointsOfFunctionPlot[11].X = 0.11
	pointsOfFunctionPlot[11].Y = 8.612

	pointsOfFunctionPlot[12].X = 0.12
	pointsOfFunctionPlot[12].Y = 7.863

	pointsOfFunctionPlot[13].X = 0.13
	pointsOfFunctionPlot[13].Y = 7.23

	pointsOfFunctionPlot[14].X = 0.14
	pointsOfFunctionPlot[14].Y = 6.688

	pointsOfFunctionPlot[15].X = 0.15
	pointsOfFunctionPlot[15].Y = 6.22

	pointsOfFunctionPlot[16].X = 0.16
	pointsOfFunctionPlot[16].Y = 5.811

	pointsOfFunctionPlot[17].X = 0.17
	pointsOfFunctionPlot[17].Y = 5.451

	pointsOfFunctionPlot[18].X = 0.18
	pointsOfFunctionPlot[18].Y = 5.131

	pointsOfFunctionPlot[19].X = 0.19
	pointsOfFunctionPlot[19].Y = 4.846

	pointsOfFunctionPlot[20].X = 0.20
	pointsOfFunctionPlot[20].Y = 4.59

	pointsOfFunctionPlot[21].X = 0.21
	pointsOfFunctionPlot[21].Y = 4.359

	pointsOfFunctionPlot[22].X = 0.22
	pointsOfFunctionPlot[22].Y = 4.15

	pointsOfFunctionPlot[23].X = 0.23
	pointsOfFunctionPlot[23].Y = 3.959

	pointsOfFunctionPlot[24].X = 0.24
	pointsOfFunctionPlot[24].Y = 3.785

	pointsOfFunctionPlot[25].X = 0.25
	pointsOfFunctionPlot[25].Y = 3.625

	pointsOfFunctionPlot[26].X = 0.26
	pointsOfFunctionPlot[26].Y = 3.478

	pointsOfFunctionPlot[27].X = 0.27
	pointsOfFunctionPlot[27].Y = 3.342

	pointsOfFunctionPlot[28].X = 0.28
	pointsOfFunctionPlot[28].Y = 3.216

	pointsOfFunctionPlot[29].X = 0.29
	pointsOfFunctionPlot[29].Y = 3.1

	pointsOfFunctionPlot[30].X = 0.30
	pointsOfFunctionPlot[30].Y = 2.991

	pointsOfFunctionPlot[31].X = 0.31
	pointsOfFunctionPlot[31].Y = 2.89

	pointsOfFunctionPlot[32].X = 0.32
	pointsOfFunctionPlot[32].Y = 2.795

	pointsOfFunctionPlot[33].X = 0.33
	pointsOfFunctionPlot[33].Y = 2.707

	pointsOfFunctionPlot[34].X = 0.34
	pointsOfFunctionPlot[34].Y = 2.624

	pointsOfFunctionPlot[35].X = 0.35
	pointsOfFunctionPlot[35].Y = 2.546

	pointsOfFunctionPlot[36].X = 0.36
	pointsOfFunctionPlot[36].Y = 2.472

	pointsOfFunctionPlot[37].X = 0.37
	pointsOfFunctionPlot[37].Y = 2.403

	pointsOfFunctionPlot[38].X = 0.38
	pointsOfFunctionPlot[38].Y = 2.338

	pointsOfFunctionPlot[39].X = 0.39
	pointsOfFunctionPlot[39].Y = 2.276

	pointsOfFunctionPlot[40].X = 0.40
	pointsOfFunctionPlot[40].Y = 2.218

	pointsOfFunctionPlot[41].X = 0.41
	pointsOfFunctionPlot[41].Y = 2.162

	pointsOfFunctionPlot[42].X = 0.42
	pointsOfFunctionPlot[42].Y = 2.11

	pointsOfFunctionPlot[43].X = 0.43
	pointsOfFunctionPlot[43].Y = 2.06

	pointsOfFunctionPlot[44].X = 0.44
	pointsOfFunctionPlot[44].Y = 2.013

	pointsOfFunctionPlot[45].X = 0.45
	pointsOfFunctionPlot[45].Y = 1.968

	pointsOfFunctionPlot[46].X = 0.46
	pointsOfFunctionPlot[46].Y = 1.925

	pointsOfFunctionPlot[47].X = 0.47
	pointsOfFunctionPlot[47].Y = 1.884

	pointsOfFunctionPlot[48].X = 0.48
	pointsOfFunctionPlot[48].Y = 1.845

	pointsOfFunctionPlot[49].X = 0.49
	pointsOfFunctionPlot[49].Y = 1.808

	pointsOfFunctionPlot[50].X = 0.50
	pointsOfFunctionPlot[50].Y = 1.772

	pointsOfFunctionPlot[51].X = 0.51
	pointsOfFunctionPlot[51].Y = 1.738

	pointsOfFunctionPlot[52].X = 0.52
	pointsOfFunctionPlot[52].Y = 1.705

	pointsOfFunctionPlot[53].X = 0.53
	pointsOfFunctionPlot[53].Y = 1.674

	pointsOfFunctionPlot[54].X = 0.54
	pointsOfFunctionPlot[54].Y = 1.644

	pointsOfFunctionPlot[55].X = 0.55
	pointsOfFunctionPlot[55].Y = 1.616

	pointsOfFunctionPlot[56].X = 0.56
	pointsOfFunctionPlot[56].Y = 1.588

	pointsOfFunctionPlot[57].X = 0.57
	pointsOfFunctionPlot[57].Y = 1.562

	pointsOfFunctionPlot[58].X = 0.58
	pointsOfFunctionPlot[58].Y = 1.536

	pointsOfFunctionPlot[59].X = 0.59
	pointsOfFunctionPlot[59].Y = 1.512

	pointsOfFunctionPlot[60].X = 0.60
	pointsOfFunctionPlot[60].Y = 1.489

	pointsOfFunctionPlot[61].X = 0.61
	pointsOfFunctionPlot[61].Y = 1.466

	pointsOfFunctionPlot[62].X = 0.62
	pointsOfFunctionPlot[62].Y = 1.445

	pointsOfFunctionPlot[63].X = 0.63
	pointsOfFunctionPlot[63].Y = 1.424

	pointsOfFunctionPlot[64].X = 0.64
	pointsOfFunctionPlot[64].Y = 1.404

	pointsOfFunctionPlot[65].X = 0.65
	pointsOfFunctionPlot[65].Y = 1.384

	pointsOfFunctionPlot[66].X = 0.66
	pointsOfFunctionPlot[66].Y = 1.366

	pointsOfFunctionPlot[67].X = 0.67
	pointsOfFunctionPlot[67].Y = 1.348

	pointsOfFunctionPlot[68].X = 0.68
	pointsOfFunctionPlot[68].Y = 1.33

	pointsOfFunctionPlot[69].X = 0.69
	pointsOfFunctionPlot[69].Y = 1.314

	pointsOfFunctionPlot[70].X = 0.70
	pointsOfFunctionPlot[70].Y = 1.298

	pointsOfFunctionPlot[71].X = 0.71
	pointsOfFunctionPlot[71].Y = 1.282

	pointsOfFunctionPlot[72].X = 0.72
	pointsOfFunctionPlot[72].Y = 1.267

	pointsOfFunctionPlot[73].X = 0.73
	pointsOfFunctionPlot[73].Y = 1.252

	pointsOfFunctionPlot[74].X = 0.74
	pointsOfFunctionPlot[74].Y = 1.238

	pointsOfFunctionPlot[75].X = 0.75
	pointsOfFunctionPlot[75].Y = 1.225

	pointsOfFunctionPlot[76].X = 0.76
	pointsOfFunctionPlot[76].Y = 1.212

	pointsOfFunctionPlot[77].X = 0.77
	pointsOfFunctionPlot[77].Y = 1.199

	pointsOfFunctionPlot[78].X = 0.78
	pointsOfFunctionPlot[78].Y = 1.187

	pointsOfFunctionPlot[79].X = 0.79
	pointsOfFunctionPlot[79].Y = 1.175

	pointsOfFunctionPlot[80].X = 0.80
	pointsOfFunctionPlot[80].Y = 1.164

	pointsOfFunctionPlot[81].X = 0.81
	pointsOfFunctionPlot[81].Y = 1.153

	pointsOfFunctionPlot[82].X = 0.82
	pointsOfFunctionPlot[82].Y = 1.142

	pointsOfFunctionPlot[83].X = 0.83
	pointsOfFunctionPlot[83].Y = 1.132

	pointsOfFunctionPlot[84].X = 0.84
	pointsOfFunctionPlot[84].Y = 1.122

	pointsOfFunctionPlot[85].X = 0.85
	pointsOfFunctionPlot[85].Y = 1.112

	pointsOfFunctionPlot[86].X = 0.86
	pointsOfFunctionPlot[86].Y = 1.103

	pointsOfFunctionPlot[87].X = 0.87
	pointsOfFunctionPlot[87].Y = 1.094

	pointsOfFunctionPlot[88].X = 0.88
	pointsOfFunctionPlot[88].Y = 1.085

	pointsOfFunctionPlot[89].X = 0.89
	pointsOfFunctionPlot[89].Y = 1.076

	pointsOfFunctionPlot[90].X = 0.90
	pointsOfFunctionPlot[90].Y = 1.068

	pointsOfFunctionPlot[91].X = 0.91
	pointsOfFunctionPlot[91].Y = 1.06

	pointsOfFunctionPlot[92].X = 0.92
	pointsOfFunctionPlot[92].Y = 1.053

	pointsOfFunctionPlot[93].X = 0.93
	pointsOfFunctionPlot[93].Y = 1.045

	pointsOfFunctionPlot[94].X = 0.94
	pointsOfFunctionPlot[94].Y = 1.038

	pointsOfFunctionPlot[95].X = 0.95
	pointsOfFunctionPlot[95].Y = 1.031

	pointsOfFunctionPlot[96].X = 0.96
	pointsOfFunctionPlot[96].Y = 1.024

	pointsOfFunctionPlot[97].X = 0.97
	pointsOfFunctionPlot[97].Y = 1.018

	pointsOfFunctionPlot[98].X = 0.98
	pointsOfFunctionPlot[98].Y = 1.011

	pointsOfFunctionPlot[99].X = 0.99
	pointsOfFunctionPlot[99].Y = 1.005

	pointsOfFunctionPlot[100].X = 1.0
	pointsOfFunctionPlot[100].Y = 1.0

	pointsOfFunctionPlot[101].X = 1.01
	pointsOfFunctionPlot[101].Y = 0.994

	pointsOfFunctionPlot[102].X = 1.02
	pointsOfFunctionPlot[102].Y = 0.988

	pointsOfFunctionPlot[103].X = 1.03
	pointsOfFunctionPlot[103].Y = 0.983

	pointsOfFunctionPlot[104].X = 1.04
	pointsOfFunctionPlot[104].Y = 0.978

	pointsOfFunctionPlot[105].X = 1.05
	pointsOfFunctionPlot[105].Y = 0.973

	pointsOfFunctionPlot[106].X = 1.06
	pointsOfFunctionPlot[106].Y = 0.968

	pointsOfFunctionPlot[107].X = 1.07
	pointsOfFunctionPlot[107].Y = 0.964

	pointsOfFunctionPlot[108].X = 1.08
	pointsOfFunctionPlot[108].Y = 0.959

	pointsOfFunctionPlot[109].X = 1.09
	pointsOfFunctionPlot[109].Y = 0.955

	pointsOfFunctionPlot[110].X = 1.10
	pointsOfFunctionPlot[110].Y = 0.951

	pointsOfFunctionPlot[111].X = 1.11
	pointsOfFunctionPlot[111].Y = 0.947

	pointsOfFunctionPlot[112].X = 1.12
	pointsOfFunctionPlot[112].Y = 0.943

	pointsOfFunctionPlot[113].X = 1.13
	pointsOfFunctionPlot[113].Y = 0.939

	pointsOfFunctionPlot[114].X = 1.14
	pointsOfFunctionPlot[114].Y = 0.936

	pointsOfFunctionPlot[115].X = 1.15
	pointsOfFunctionPlot[115].Y = 0.933

	pointsOfFunctionPlot[116].X = 1.16
	pointsOfFunctionPlot[116].Y = 0.929

	pointsOfFunctionPlot[117].X = 1.17
	pointsOfFunctionPlot[117].Y = 0.926

	pointsOfFunctionPlot[118].X = 1.18
	pointsOfFunctionPlot[118].Y = 0.923

	pointsOfFunctionPlot[119].X = 1.19
	pointsOfFunctionPlot[119].Y = 0.92

	pointsOfFunctionPlot[120].X = 1.20
	pointsOfFunctionPlot[120].Y = 0.918

	pointsOfFunctionPlot[121].X = 1.21
	pointsOfFunctionPlot[121].Y = 0.915

	pointsOfFunctionPlot[122].X = 1.22
	pointsOfFunctionPlot[122].Y = 0.913

	pointsOfFunctionPlot[123].X = 1.23
	pointsOfFunctionPlot[123].Y = 0.91

	pointsOfFunctionPlot[124].X = 1.24
	pointsOfFunctionPlot[124].Y = 0.908

	pointsOfFunctionPlot[125].X = 1.25
	pointsOfFunctionPlot[125].Y = 0.906

	pointsOfFunctionPlot[126].X = 1.26
	pointsOfFunctionPlot[126].Y = 0.904

	pointsOfFunctionPlot[127].X = 1.27
	pointsOfFunctionPlot[127].Y = 0.902

	pointsOfFunctionPlot[128].X = 1.28
	pointsOfFunctionPlot[128].Y = 0.9

	pointsOfFunctionPlot[129].X = 1.29
	pointsOfFunctionPlot[129].Y = 0.899

	pointsOfFunctionPlot[130].X = 1.30
	pointsOfFunctionPlot[130].Y = 0.897

	pointsOfFunctionPlot[131].X = 1.31
	pointsOfFunctionPlot[131].Y = 0.896

	pointsOfFunctionPlot[132].X = 1.32
	pointsOfFunctionPlot[132].Y = 0.894

	pointsOfFunctionPlot[133].X = 1.33
	pointsOfFunctionPlot[133].Y = 0.893

	pointsOfFunctionPlot[134].X = 1.34
	pointsOfFunctionPlot[134].Y = 0.892

	pointsOfFunctionPlot[135].X = 1.35
	pointsOfFunctionPlot[135].Y = 0.891

	pointsOfFunctionPlot[136].X = 1.36
	pointsOfFunctionPlot[136].Y = 0.89

	pointsOfFunctionPlot[137].X = 1.37
	pointsOfFunctionPlot[137].Y = 0.889

	pointsOfFunctionPlot[138].X = 1.38
	pointsOfFunctionPlot[138].Y = 0.888

	pointsOfFunctionPlot[139].X = 1.39
	pointsOfFunctionPlot[139].Y = 0.887

	pointsOfFunctionPlot[140].X = 1.40
	pointsOfFunctionPlot[140].Y = 0.887

	pointsOfFunctionPlot[141].X = 1.41
	pointsOfFunctionPlot[141].Y = 0.886

	pointsOfFunctionPlot[142].X = 1.42
	pointsOfFunctionPlot[142].Y = 0.886

	pointsOfFunctionPlot[143].X = 1.43
	pointsOfFunctionPlot[143].Y = 0.886

	pointsOfFunctionPlot[144].X = 1.44
	pointsOfFunctionPlot[144].Y = 0.885

	pointsOfFunctionPlot[145].X = 1.45
	pointsOfFunctionPlot[145].Y = 0.885

	pointsOfFunctionPlot[146].X = 1.46
	pointsOfFunctionPlot[146].Y = 0.885

	pointsOfFunctionPlot[147].X = 1.47
	pointsOfFunctionPlot[147].Y = 0.885

	pointsOfFunctionPlot[148].X = 1.48
	pointsOfFunctionPlot[148].Y = 0.885

	pointsOfFunctionPlot[149].X = 1.49
	pointsOfFunctionPlot[149].Y = 0.885

	pointsOfFunctionPlot[150].X = 1.50
	pointsOfFunctionPlot[150].Y = 0.886

	pointsOfFunctionPlot[151].X = 1.51
	pointsOfFunctionPlot[151].Y = 0.886

	pointsOfFunctionPlot[152].X = 1.52
	pointsOfFunctionPlot[152].Y = 0.887

	pointsOfFunctionPlot[153].X = 1.53
	pointsOfFunctionPlot[153].Y = 0.887

	pointsOfFunctionPlot[154].X = 1.54
	pointsOfFunctionPlot[154].Y = 0.888

	pointsOfFunctionPlot[155].X = 1.55
	pointsOfFunctionPlot[155].Y = 0.888

	pointsOfFunctionPlot[156].X = 1.56
	pointsOfFunctionPlot[156].Y = 0.889

	pointsOfFunctionPlot[157].X = 1.57
	pointsOfFunctionPlot[157].Y = 0.89

	pointsOfFunctionPlot[158].X = 1.58
	pointsOfFunctionPlot[158].Y = 0.891

	pointsOfFunctionPlot[159].X = 1.59
	pointsOfFunctionPlot[159].Y = 0.892

	pointsOfFunctionPlot[160].X = 1.60
	pointsOfFunctionPlot[160].Y = 0.893

	pointsOfFunctionPlot[161].X = 1.61
	pointsOfFunctionPlot[161].Y = 0.894

	pointsOfFunctionPlot[162].X = 1.62
	pointsOfFunctionPlot[162].Y = 0.895

	pointsOfFunctionPlot[163].X = 1.63
	pointsOfFunctionPlot[163].Y = 0.897

	pointsOfFunctionPlot[164].X = 1.64
	pointsOfFunctionPlot[164].Y = 0.898

	pointsOfFunctionPlot[165].X = 1.65
	pointsOfFunctionPlot[165].Y = 0.9

	pointsOfFunctionPlot[166].X = 1.66
	pointsOfFunctionPlot[166].Y = 0.901

	pointsOfFunctionPlot[167].X = 1.67
	pointsOfFunctionPlot[167].Y = 0.903

	pointsOfFunctionPlot[168].X = 1.68
	pointsOfFunctionPlot[168].Y = 0.905

	pointsOfFunctionPlot[169].X = 1.69
	pointsOfFunctionPlot[169].Y = 0.906

	pointsOfFunctionPlot[170].X = 1.70
	pointsOfFunctionPlot[170].Y = 0.908

	pointsOfFunctionPlot[171].X = 1.71
	pointsOfFunctionPlot[171].Y = 0.91

	pointsOfFunctionPlot[172].X = 1.72
	pointsOfFunctionPlot[172].Y = 0.912

	pointsOfFunctionPlot[173].X = 1.73
	pointsOfFunctionPlot[173].Y = 0.914

	pointsOfFunctionPlot[174].X = 1.74
	pointsOfFunctionPlot[174].Y = 0.916

	pointsOfFunctionPlot[175].X = 1.75
	pointsOfFunctionPlot[175].Y = 0.919

	pointsOfFunctionPlot[176].X = 1.76
	pointsOfFunctionPlot[176].Y = 0.921

	pointsOfFunctionPlot[177].X = 1.77
	pointsOfFunctionPlot[177].Y = 0.923

	pointsOfFunctionPlot[178].X = 1.78
	pointsOfFunctionPlot[178].Y = 0.926

	pointsOfFunctionPlot[179].X = 1.79
	pointsOfFunctionPlot[179].Y = 0.928

	pointsOfFunctionPlot[180].X = 1.80
	pointsOfFunctionPlot[180].Y = 0.931

	pointsOfFunctionPlot[181].X = 1.81
	pointsOfFunctionPlot[181].Y = 0.934

	pointsOfFunctionPlot[182].X = 1.82
	pointsOfFunctionPlot[182].Y = 0.936

	pointsOfFunctionPlot[183].X = 1.83
	pointsOfFunctionPlot[183].Y = 0.939

	pointsOfFunctionPlot[184].X = 1.84
	pointsOfFunctionPlot[184].Y = 0.942

	pointsOfFunctionPlot[185].X = 1.85
	pointsOfFunctionPlot[185].Y = 0.945

	pointsOfFunctionPlot[186].X = 1.86
	pointsOfFunctionPlot[186].Y = 0.948

	pointsOfFunctionPlot[187].X = 1.87
	pointsOfFunctionPlot[187].Y = 0.951

	pointsOfFunctionPlot[188].X = 1.88
	pointsOfFunctionPlot[188].Y = 0.955

	pointsOfFunctionPlot[189].X = 1.89
	pointsOfFunctionPlot[189].Y = 0.958

	pointsOfFunctionPlot[190].X = 1.90
	pointsOfFunctionPlot[190].Y = 0.961

	pointsOfFunctionPlot[191].X = 1.91
	pointsOfFunctionPlot[191].Y = 0.965

	pointsOfFunctionPlot[192].X = 1.92
	pointsOfFunctionPlot[192].Y = 0.968

	pointsOfFunctionPlot[193].X = 1.93
	pointsOfFunctionPlot[193].Y = 0.972

	pointsOfFunctionPlot[194].X = 1.94
	pointsOfFunctionPlot[194].Y = 0.976

	pointsOfFunctionPlot[195].X = 1.95
	pointsOfFunctionPlot[195].Y = 0.979

	pointsOfFunctionPlot[196].X = 1.96
	pointsOfFunctionPlot[196].Y = 0.983

	pointsOfFunctionPlot[197].X = 1.97
	pointsOfFunctionPlot[197].Y = 0.987

	pointsOfFunctionPlot[198].X = 1.98
	pointsOfFunctionPlot[198].Y = 0.991

	pointsOfFunctionPlot[199].X = 1.99
	pointsOfFunctionPlot[199].Y = 0.995

	pointsOfFunctionPlot[200].X = 2.0
	pointsOfFunctionPlot[200].Y = 1.0

	pointsOfFunctionPlot[201].X = 2.01
	pointsOfFunctionPlot[201].Y = 1.004

	pointsOfFunctionPlot[202].X = 2.02
	pointsOfFunctionPlot[202].Y = 1.008

	pointsOfFunctionPlot[203].X = 2.03
	pointsOfFunctionPlot[203].Y = 1.013

	pointsOfFunctionPlot[204].X = 2.04
	pointsOfFunctionPlot[204].Y = 1.017

	pointsOfFunctionPlot[205].X = 2.05
	pointsOfFunctionPlot[205].Y = 1.022

	pointsOfFunctionPlot[206].X = 2.06
	pointsOfFunctionPlot[206].Y = 1.026

	pointsOfFunctionPlot[207].X = 2.07
	pointsOfFunctionPlot[207].Y = 1.031

	pointsOfFunctionPlot[208].X = 2.08
	pointsOfFunctionPlot[208].Y = 1.036

	pointsOfFunctionPlot[209].X = 2.09
	pointsOfFunctionPlot[209].Y = 1.041

	pointsOfFunctionPlot[210].X = 2.10
	pointsOfFunctionPlot[210].Y = 1.046

	pointsOfFunctionPlot[211].X = 2.11
	pointsOfFunctionPlot[211].Y = 1.051

	pointsOfFunctionPlot[212].X = 2.12
	pointsOfFunctionPlot[212].Y = 1.056

	pointsOfFunctionPlot[213].X = 2.13
	pointsOfFunctionPlot[213].Y = 1.062

	pointsOfFunctionPlot[214].X = 2.14
	pointsOfFunctionPlot[214].Y = 1.067

	pointsOfFunctionPlot[215].X = 2.15
	pointsOfFunctionPlot[215].Y = 1.072

	pointsOfFunctionPlot[216].X = 2.16
	pointsOfFunctionPlot[216].Y = 1.078

	pointsOfFunctionPlot[217].X = 2.17
	pointsOfFunctionPlot[217].Y = 1.084

	pointsOfFunctionPlot[218].X = 2.18
	pointsOfFunctionPlot[218].Y = 1.089

	pointsOfFunctionPlot[219].X = 2.19
	pointsOfFunctionPlot[219].Y = 1.095

	pointsOfFunctionPlot[220].X = 2.20
	pointsOfFunctionPlot[220].Y = 1.101

	pointsOfFunctionPlot[221].X = 2.21
	pointsOfFunctionPlot[221].Y = 1.107

	pointsOfFunctionPlot[222].X = 2.22
	pointsOfFunctionPlot[222].Y = 1.113

	pointsOfFunctionPlot[223].X = 2.23
	pointsOfFunctionPlot[223].Y = 1.12

	pointsOfFunctionPlot[224].X = 2.24
	pointsOfFunctionPlot[224].Y = 1.126

	pointsOfFunctionPlot[225].X = 2.25
	pointsOfFunctionPlot[225].Y = 1.133

	pointsOfFunctionPlot[226].X = 2.26
	pointsOfFunctionPlot[226].Y = 1.139

	pointsOfFunctionPlot[227].X = 2.27
	pointsOfFunctionPlot[227].Y = 1.146

	pointsOfFunctionPlot[228].X = 2.28
	pointsOfFunctionPlot[228].Y = 1.152

	pointsOfFunctionPlot[229].X = 2.29
	pointsOfFunctionPlot[229].Y = 1.159

	pointsOfFunctionPlot[230].X = 2.30
	pointsOfFunctionPlot[230].Y = 1.166

	pointsOfFunctionPlot[231].X = 2.31
	pointsOfFunctionPlot[231].Y = 1.173

	pointsOfFunctionPlot[232].X = 2.32
	pointsOfFunctionPlot[232].Y = 1.18

	pointsOfFunctionPlot[233].X = 2.33
	pointsOfFunctionPlot[233].Y = 1.188

	pointsOfFunctionPlot[234].X = 2.34
	pointsOfFunctionPlot[234].Y = 1.195

	pointsOfFunctionPlot[235].X = 2.35
	pointsOfFunctionPlot[235].Y = 1.203

	pointsOfFunctionPlot[236].X = 2.36
	pointsOfFunctionPlot[236].Y = 1.21

	pointsOfFunctionPlot[237].X = 2.37
	pointsOfFunctionPlot[237].Y = 1.218

	pointsOfFunctionPlot[238].X = 2.38
	pointsOfFunctionPlot[238].Y = 1.226

	pointsOfFunctionPlot[239].X = 2.39
	pointsOfFunctionPlot[239].Y = 1.234

	pointsOfFunctionPlot[240].X = 2.40
	pointsOfFunctionPlot[240].Y = 1.242

	pointsOfFunctionPlot[241].X = 2.41
	pointsOfFunctionPlot[241].Y = 1.25

	pointsOfFunctionPlot[242].X = 2.42
	pointsOfFunctionPlot[242].Y = 1.258

	pointsOfFunctionPlot[243].X = 2.43
	pointsOfFunctionPlot[243].Y = 1.267

	pointsOfFunctionPlot[244].X = 2.44
	pointsOfFunctionPlot[244].Y = 1.275

	pointsOfFunctionPlot[245].X = 2.45
	pointsOfFunctionPlot[245].Y = 1.284

	pointsOfFunctionPlot[246].X = 2.46
	pointsOfFunctionPlot[246].Y = 1.292

	pointsOfFunctionPlot[247].X = 2.47
	pointsOfFunctionPlot[247].Y = 1.301

	pointsOfFunctionPlot[248].X = 2.48
	pointsOfFunctionPlot[248].Y = 1.31

	pointsOfFunctionPlot[249].X = 2.49
	pointsOfFunctionPlot[249].Y = 1.32

	pointsOfFunctionPlot[250].X = 2.50
	pointsOfFunctionPlot[250].Y = 1.329

	pointsOfFunctionPlot[251].X = 2.51
	pointsOfFunctionPlot[251].Y = 1.338

	pointsOfFunctionPlot[252].X = 2.52
	pointsOfFunctionPlot[252].Y = 1.348

	pointsOfFunctionPlot[253].X = 2.53
	pointsOfFunctionPlot[253].Y = 1.357

	pointsOfFunctionPlot[254].X = 2.54
	pointsOfFunctionPlot[254].Y = 1.367

	pointsOfFunctionPlot[255].X = 2.55
	pointsOfFunctionPlot[255].Y = 1.377

	pointsOfFunctionPlot[256].X = 2.56
	pointsOfFunctionPlot[256].Y = 1.387

	pointsOfFunctionPlot[257].X = 2.57
	pointsOfFunctionPlot[257].Y = 1.398

	pointsOfFunctionPlot[258].X = 2.58
	pointsOfFunctionPlot[258].Y = 1.408

	pointsOfFunctionPlot[259].X = 2.59
	pointsOfFunctionPlot[259].Y = 1.418

	pointsOfFunctionPlot[260].X = 2.60
	pointsOfFunctionPlot[260].Y = 1.429

	pointsOfFunctionPlot[261].X = 2.61
	pointsOfFunctionPlot[261].Y = 1.44

	pointsOfFunctionPlot[262].X = 2.62
	pointsOfFunctionPlot[262].Y = 1.451

	pointsOfFunctionPlot[263].X = 2.63
	pointsOfFunctionPlot[263].Y = 1.462

	pointsOfFunctionPlot[264].X = 2.64
	pointsOfFunctionPlot[264].Y = 1.473

	pointsOfFunctionPlot[265].X = 2.65
	pointsOfFunctionPlot[265].Y = 1.485

	pointsOfFunctionPlot[266].X = 2.66
	pointsOfFunctionPlot[266].Y = 1.496

	pointsOfFunctionPlot[267].X = 2.67
	pointsOfFunctionPlot[267].Y = 1.508

	pointsOfFunctionPlot[268].X = 2.68
	pointsOfFunctionPlot[268].Y = 1.52

	pointsOfFunctionPlot[269].X = 2.69
	pointsOfFunctionPlot[269].Y = 1.532

	pointsOfFunctionPlot[270].X = 2.70
	pointsOfFunctionPlot[270].Y = 1.544

	pointsOfFunctionPlot[271].X = 2.71
	pointsOfFunctionPlot[271].Y = 1.557

	pointsOfFunctionPlot[272].X = 2.72
	pointsOfFunctionPlot[272].Y = 1.569

	pointsOfFunctionPlot[273].X = 2.73
	pointsOfFunctionPlot[273].Y = 1.582

	pointsOfFunctionPlot[274].X = 2.74
	pointsOfFunctionPlot[274].Y = 1.595

	pointsOfFunctionPlot[275].X = 2.75
	pointsOfFunctionPlot[275].Y = 1.608

	pointsOfFunctionPlot[276].X = 2.76
	pointsOfFunctionPlot[276].Y = 1.621

	pointsOfFunctionPlot[277].X = 2.77
	pointsOfFunctionPlot[277].Y = 1.635

	pointsOfFunctionPlot[278].X = 2.78
	pointsOfFunctionPlot[278].Y = 1.648

	pointsOfFunctionPlot[279].X = 2.79
	pointsOfFunctionPlot[279].Y = 1.662

	pointsOfFunctionPlot[280].X = 2.80
	pointsOfFunctionPlot[280].Y = 1.676

	pointsOfFunctionPlot[281].X = 2.81
	pointsOfFunctionPlot[281].Y = 1.69

	pointsOfFunctionPlot[282].X = 2.82
	pointsOfFunctionPlot[282].Y = 1.705

	pointsOfFunctionPlot[283].X = 2.83
	pointsOfFunctionPlot[283].Y = 1.719

	pointsOfFunctionPlot[284].X = 2.84
	pointsOfFunctionPlot[284].Y = 1.734

	pointsOfFunctionPlot[285].X = 2.85
	pointsOfFunctionPlot[285].Y = 1.749

	pointsOfFunctionPlot[286].X = 2.86
	pointsOfFunctionPlot[286].Y = 1.764

	pointsOfFunctionPlot[287].X = 2.87
	pointsOfFunctionPlot[287].Y = 1.779

	pointsOfFunctionPlot[288].X = 2.88
	pointsOfFunctionPlot[288].Y = 1.795

	pointsOfFunctionPlot[289].X = 2.89
	pointsOfFunctionPlot[289].Y = 1.811

	pointsOfFunctionPlot[290].X = 2.90
	pointsOfFunctionPlot[290].Y = 1.827

	pointsOfFunctionPlot[291].X = 2.91
	pointsOfFunctionPlot[291].Y = 1.843

	pointsOfFunctionPlot[292].X = 2.92
	pointsOfFunctionPlot[292].Y = 1.86

	pointsOfFunctionPlot[293].X = 2.93
	pointsOfFunctionPlot[293].Y = 1.876

	pointsOfFunctionPlot[294].X = 2.94
	pointsOfFunctionPlot[294].Y = 1.893

	pointsOfFunctionPlot[295].X = 2.95
	pointsOfFunctionPlot[295].Y = 1.91

	pointsOfFunctionPlot[296].X = 2.96
	pointsOfFunctionPlot[296].Y = 1.928

	pointsOfFunctionPlot[297].X = 2.97
	pointsOfFunctionPlot[297].Y = 1.945

	pointsOfFunctionPlot[298].X = 2.98
	pointsOfFunctionPlot[298].Y = 1.963

	pointsOfFunctionPlot[299].X = 2.99
	pointsOfFunctionPlot[299].Y = 1.981

	pointsOfFunctionPlot[300].X = 3.0
	pointsOfFunctionPlot[300].Y = 2.0

	pointsOfFunctionPlot[301].X = 3.01
	pointsOfFunctionPlot[301].Y = 2.018

	pointsOfFunctionPlot[302].X = 3.02
	pointsOfFunctionPlot[302].Y = 2.037

	pointsOfFunctionPlot[303].X = 3.03
	pointsOfFunctionPlot[303].Y = 2.056

	pointsOfFunctionPlot[304].X = 3.04
	pointsOfFunctionPlot[304].Y = 2.075

	pointsOfFunctionPlot[305].X = 3.05
	pointsOfFunctionPlot[305].Y = 2.095

	pointsOfFunctionPlot[306].X = 3.06
	pointsOfFunctionPlot[306].Y = 2.115

	pointsOfFunctionPlot[307].X = 3.07
	pointsOfFunctionPlot[307].Y = 2.135

	pointsOfFunctionPlot[308].X = 3.08
	pointsOfFunctionPlot[308].Y = 2.155

	pointsOfFunctionPlot[309].X = 3.09
	pointsOfFunctionPlot[309].Y = 2.176

	pointsOfFunctionPlot[310].X = 3.10
	pointsOfFunctionPlot[310].Y = 2.197

	pointsOfFunctionPlot[311].X = 3.11
	pointsOfFunctionPlot[311].Y = 2.218

	pointsOfFunctionPlot[312].X = 3.12
	pointsOfFunctionPlot[312].Y = 2.24

	pointsOfFunctionPlot[313].X = 3.13
	pointsOfFunctionPlot[313].Y = 2.262

	pointsOfFunctionPlot[314].X = 3.14
	pointsOfFunctionPlot[314].Y = 2.284

	pointsOfFunctionPlot[315].X = 3.15
	pointsOfFunctionPlot[315].Y = 2.306

	pointsOfFunctionPlot[316].X = 3.16
	pointsOfFunctionPlot[316].Y = 2.329

	pointsOfFunctionPlot[317].X = 3.17
	pointsOfFunctionPlot[317].Y = 2.352

	pointsOfFunctionPlot[318].X = 3.18
	pointsOfFunctionPlot[318].Y = 2.376

	pointsOfFunctionPlot[319].X = 3.19
	pointsOfFunctionPlot[319].Y = 2.399

	pointsOfFunctionPlot[320].X = 3.20
	pointsOfFunctionPlot[320].Y = 2.423

	pointsOfFunctionPlot[321].X = 3.21
	pointsOfFunctionPlot[321].Y = 2.448

	pointsOfFunctionPlot[322].X = 3.22
	pointsOfFunctionPlot[322].Y = 2.473

	pointsOfFunctionPlot[323].X = 3.23
	pointsOfFunctionPlot[323].Y = 2.498

	pointsOfFunctionPlot[324].X = 3.24
	pointsOfFunctionPlot[324].Y = 2.523

	pointsOfFunctionPlot[325].X = 3.25
	pointsOfFunctionPlot[325].Y = 2.549

	pointsOfFunctionPlot[326].X = 3.26
	pointsOfFunctionPlot[326].Y = 2.575

	pointsOfFunctionPlot[327].X = 3.27
	pointsOfFunctionPlot[327].Y = 2.601

	pointsOfFunctionPlot[328].X = 3.28
	pointsOfFunctionPlot[328].Y = 2.628

	pointsOfFunctionPlot[329].X = 3.29
	pointsOfFunctionPlot[329].Y = 2.655

	pointsOfFunctionPlot[330].X = 3.30
	pointsOfFunctionPlot[330].Y = 2.683

	pointsOfFunctionPlot[331].X = 3.31
	pointsOfFunctionPlot[331].Y = 2.711

	pointsOfFunctionPlot[332].X = 3.32
	pointsOfFunctionPlot[332].Y = 2.739

	pointsOfFunctionPlot[333].X = 3.33
	pointsOfFunctionPlot[333].Y = 2.768

	pointsOfFunctionPlot[334].X = 3.34
	pointsOfFunctionPlot[334].Y = 2.797

	pointsOfFunctionPlot[335].X = 3.35
	pointsOfFunctionPlot[335].Y = 2.827

	pointsOfFunctionPlot[336].X = 3.36
	pointsOfFunctionPlot[336].Y = 2.857

	pointsOfFunctionPlot[337].X = 3.37
	pointsOfFunctionPlot[337].Y = 2.887

	pointsOfFunctionPlot[338].X = 3.38
	pointsOfFunctionPlot[338].Y = 2.918

	pointsOfFunctionPlot[339].X = 3.39
	pointsOfFunctionPlot[339].Y = 2.949

	pointsOfFunctionPlot[340].X = 3.40
	pointsOfFunctionPlot[340].Y = 2.981

	pointsOfFunctionPlot[341].X = 3.41
	pointsOfFunctionPlot[341].Y = 3.013

	pointsOfFunctionPlot[342].X = 3.42
	pointsOfFunctionPlot[342].Y = 3.045

	pointsOfFunctionPlot[343].X = 3.43
	pointsOfFunctionPlot[343].Y = 3.078

	pointsOfFunctionPlot[344].X = 3.44
	pointsOfFunctionPlot[344].Y = 3.112

	pointsOfFunctionPlot[345].X = 3.45
	pointsOfFunctionPlot[345].Y = 3.146

	pointsOfFunctionPlot[346].X = 3.46
	pointsOfFunctionPlot[346].Y = 3.18

	pointsOfFunctionPlot[347].X = 3.47
	pointsOfFunctionPlot[347].Y = 3.215

	pointsOfFunctionPlot[348].X = 3.48
	pointsOfFunctionPlot[348].Y = 3.251

	pointsOfFunctionPlot[349].X = 3.49
	pointsOfFunctionPlot[349].Y = 3.286

	pointsOfFunctionPlot[350].X = 3.50
	pointsOfFunctionPlot[350].Y = 3.323

	pointsOfFunctionPlot[351].X = 3.51
	pointsOfFunctionPlot[351].Y = 3.36

	pointsOfFunctionPlot[352].X = 3.52
	pointsOfFunctionPlot[352].Y = 3.397

	pointsOfFunctionPlot[353].X = 3.53
	pointsOfFunctionPlot[353].Y = 3.435

	pointsOfFunctionPlot[354].X = 3.54
	pointsOfFunctionPlot[354].Y = 3.474

	pointsOfFunctionPlot[355].X = 3.55
	pointsOfFunctionPlot[355].Y = 3.513

	pointsOfFunctionPlot[356].X = 3.56
	pointsOfFunctionPlot[356].Y = 3.552

	pointsOfFunctionPlot[357].X = 3.57
	pointsOfFunctionPlot[357].Y = 3.593

	pointsOfFunctionPlot[358].X = 3.58
	pointsOfFunctionPlot[358].Y = 3.633

	pointsOfFunctionPlot[359].X = 3.59
	pointsOfFunctionPlot[359].Y = 3.675

	pointsOfFunctionPlot[360].X = 3.60
	pointsOfFunctionPlot[360].Y = 3.717

	pointsOfFunctionPlot[361].X = 3.61
	pointsOfFunctionPlot[361].Y = 3.759

	pointsOfFunctionPlot[362].X = 3.62
	pointsOfFunctionPlot[362].Y = 3.802

	pointsOfFunctionPlot[363].X = 3.63
	pointsOfFunctionPlot[363].Y = 3.846

	pointsOfFunctionPlot[364].X = 3.64
	pointsOfFunctionPlot[364].Y = 3.89

	pointsOfFunctionPlot[365].X = 3.65
	pointsOfFunctionPlot[365].Y = 3.935

	pointsOfFunctionPlot[366].X = 3.66
	pointsOfFunctionPlot[366].Y = 3.981

	pointsOfFunctionPlot[367].X = 3.67
	pointsOfFunctionPlot[367].Y = 4.027

	pointsOfFunctionPlot[368].X = 3.68
	pointsOfFunctionPlot[368].Y = 4.074

	pointsOfFunctionPlot[369].X = 3.69
	pointsOfFunctionPlot[369].Y = 4.122

	pointsOfFunctionPlot[370].X = 3.70
	pointsOfFunctionPlot[370].Y = 4.17

	pointsOfFunctionPlot[371].X = 3.71
	pointsOfFunctionPlot[371].Y = 4.219

	pointsOfFunctionPlot[372].X = 3.72
	pointsOfFunctionPlot[372].Y = 4.269

	pointsOfFunctionPlot[373].X = 3.73
	pointsOfFunctionPlot[373].Y = 4.319

	pointsOfFunctionPlot[374].X = 3.74
	pointsOfFunctionPlot[374].Y = 4.371

	pointsOfFunctionPlot[375].X = 3.75
	pointsOfFunctionPlot[375].Y = 4.422

	pointsOfFunctionPlot[376].X = 3.76
	pointsOfFunctionPlot[376].Y = 4.475

	pointsOfFunctionPlot[377].X = 3.77
	pointsOfFunctionPlot[377].Y = 4.529

	pointsOfFunctionPlot[378].X = 3.78
	pointsOfFunctionPlot[378].Y = 4.583

	pointsOfFunctionPlot[379].X = 3.79
	pointsOfFunctionPlot[379].Y = 4.638

	pointsOfFunctionPlot[380].X = 3.80
	pointsOfFunctionPlot[380].Y = 4.694

	pointsOfFunctionPlot[381].X = 3.81
	pointsOfFunctionPlot[381].Y = 4.75

	pointsOfFunctionPlot[382].X = 3.82
	pointsOfFunctionPlot[382].Y = 4.808

	pointsOfFunctionPlot[383].X = 3.83
	pointsOfFunctionPlot[383].Y = 4.866

	pointsOfFunctionPlot[384].X = 3.84
	pointsOfFunctionPlot[384].Y = 4.925

	pointsOfFunctionPlot[385].X = 3.85
	pointsOfFunctionPlot[385].Y = 4.985

	pointsOfFunctionPlot[386].X = 3.86
	pointsOfFunctionPlot[386].Y = 5.046

	pointsOfFunctionPlot[387].X = 3.87
	pointsOfFunctionPlot[387].Y = 5.108

	pointsOfFunctionPlot[388].X = 3.88
	pointsOfFunctionPlot[388].Y = 5.171

	pointsOfFunctionPlot[389].X = 3.89
	pointsOfFunctionPlot[389].Y = 5.234

	pointsOfFunctionPlot[390].X = 3.90
	pointsOfFunctionPlot[390].Y = 5.299

	pointsOfFunctionPlot[391].X = 3.91
	pointsOfFunctionPlot[391].Y = 5.364

	pointsOfFunctionPlot[392].X = 3.92
	pointsOfFunctionPlot[392].Y = 5.431

	pointsOfFunctionPlot[393].X = 3.93
	pointsOfFunctionPlot[393].Y = 5.498

	pointsOfFunctionPlot[394].X = 3.94
	pointsOfFunctionPlot[394].Y = 5.567

	pointsOfFunctionPlot[395].X = 3.95
	pointsOfFunctionPlot[395].Y = 5.636

	pointsOfFunctionPlot[396].X = 3.96
	pointsOfFunctionPlot[396].Y = 5.707

	pointsOfFunctionPlot[397].X = 3.97
	pointsOfFunctionPlot[397].Y = 5.778

	pointsOfFunctionPlot[398].X = 3.98
	pointsOfFunctionPlot[398].Y = 5.851

	pointsOfFunctionPlot[399].X = 3.99
	pointsOfFunctionPlot[399].Y = 5.925

	pointsOfFunctionPlot[400].X = 4.0
	pointsOfFunctionPlot[400].Y = 6.0

	pointsOfFunctionPlot[401].X = 4.01
	pointsOfFunctionPlot[401].Y = 6.075

	pointsOfFunctionPlot[402].X = 4.02
	pointsOfFunctionPlot[402].Y = 6.152

	pointsOfFunctionPlot[403].X = 4.03
	pointsOfFunctionPlot[403].Y = 6.231

	pointsOfFunctionPlot[404].X = 4.04
	pointsOfFunctionPlot[404].Y = 6.31

	pointsOfFunctionPlot[405].X = 4.05
	pointsOfFunctionPlot[405].Y = 6.391

	pointsOfFunctionPlot[406].X = 4.06
	pointsOfFunctionPlot[406].Y = 6.472

	pointsOfFunctionPlot[407].X = 4.07
	pointsOfFunctionPlot[407].Y = 6.555

	pointsOfFunctionPlot[408].X = 4.08
	pointsOfFunctionPlot[408].Y = 6.64

	pointsOfFunctionPlot[409].X = 4.09
	pointsOfFunctionPlot[409].Y = 6.725

	pointsOfFunctionPlot[410].X = 4.10
	pointsOfFunctionPlot[410].Y = 6.812

	pointsOfFunctionPlot[411].X = 4.11
	pointsOfFunctionPlot[411].Y = 6.9

	pointsOfFunctionPlot[412].X = 4.12
	pointsOfFunctionPlot[412].Y = 6.99

	pointsOfFunctionPlot[413].X = 4.13
	pointsOfFunctionPlot[413].Y = 7.081

	pointsOfFunctionPlot[414].X = 4.14
	pointsOfFunctionPlot[414].Y = 7.173

	pointsOfFunctionPlot[415].X = 4.15
	pointsOfFunctionPlot[415].Y = 7.266

	pointsOfFunctionPlot[416].X = 4.16
	pointsOfFunctionPlot[416].Y = 7.361

	pointsOfFunctionPlot[417].X = 4.17
	pointsOfFunctionPlot[417].Y = 7.458

	pointsOfFunctionPlot[418].X = 4.18
	pointsOfFunctionPlot[418].Y = 7.556

	pointsOfFunctionPlot[419].X = 4.19
	pointsOfFunctionPlot[419].Y = 7.655

	pointsOfFunctionPlot[420].X = 4.20
	pointsOfFunctionPlot[420].Y = 7.756

	pointsOfFunctionPlot[421].X = 4.21
	pointsOfFunctionPlot[421].Y = 7.859

	pointsOfFunctionPlot[422].X = 4.22
	pointsOfFunctionPlot[422].Y = 7.963

	pointsOfFunctionPlot[423].X = 4.23
	pointsOfFunctionPlot[423].Y = 8.068

	pointsOfFunctionPlot[424].X = 4.24
	pointsOfFunctionPlot[424].Y = 8.176

	pointsOfFunctionPlot[425].X = 4.25
	pointsOfFunctionPlot[425].Y = 8.285

	pointsOfFunctionPlot[426].X = 4.26
	pointsOfFunctionPlot[426].Y = 8.395

	pointsOfFunctionPlot[427].X = 4.27
	pointsOfFunctionPlot[427].Y = 8.507

	pointsOfFunctionPlot[428].X = 4.28
	pointsOfFunctionPlot[428].Y = 8.621

	pointsOfFunctionPlot[429].X = 4.29
	pointsOfFunctionPlot[429].Y = 8.737

	pointsOfFunctionPlot[430].X = 4.30
	pointsOfFunctionPlot[430].Y = 8.855

	pointsOfFunctionPlot[431].X = 4.31
	pointsOfFunctionPlot[431].Y = 8.974

	pointsOfFunctionPlot[432].X = 4.32
	pointsOfFunctionPlot[432].Y = 9.095

	pointsOfFunctionPlot[433].X = 4.33
	pointsOfFunctionPlot[433].Y = 9.219

	pointsOfFunctionPlot[434].X = 4.34
	pointsOfFunctionPlot[434].Y = 9.344

	pointsOfFunctionPlot[435].X = 4.35
	pointsOfFunctionPlot[435].Y = 9.471

	pointsOfFunctionPlot[436].X = 4.36
	pointsOfFunctionPlot[436].Y = 9.599

	pointsOfFunctionPlot[437].X = 4.37
	pointsOfFunctionPlot[437].Y = 9.73

	pointsOfFunctionPlot[438].X = 4.38
	pointsOfFunctionPlot[438].Y = 9.863

	pointsOfFunctionPlot[439].X = 4.39
	pointsOfFunctionPlot[439].Y = 9.998

	pointsOfFunctionPlot[440].X = 4.40
	pointsOfFunctionPlot[440].Y = 10.136

	pointsOfFunctionPlot[441].X = 4.41
	pointsOfFunctionPlot[441].Y = 10.275

	pointsOfFunctionPlot[442].X = 4.42
	pointsOfFunctionPlot[442].Y = 10.416

	pointsOfFunctionPlot[443].X = 4.43
	pointsOfFunctionPlot[443].Y = 10.56

	pointsOfFunctionPlot[444].X = 4.44
	pointsOfFunctionPlot[444].Y = 10.706

	pointsOfFunctionPlot[445].X = 4.45
	pointsOfFunctionPlot[445].Y = 10.854

	pointsOfFunctionPlot[446].X = 4.46
	pointsOfFunctionPlot[446].Y = 11.005

	pointsOfFunctionPlot[447].X = 4.47
	pointsOfFunctionPlot[447].Y = 11.158

	pointsOfFunctionPlot[448].X = 4.48
	pointsOfFunctionPlot[448].Y = 11.313

	pointsOfFunctionPlot[449].X = 4.49
	pointsOfFunctionPlot[449].Y = 11.471

	pointsOfFunctionPlot[450].X = 4.50
	pointsOfFunctionPlot[450].Y = 11.631

	pointsOfFunctionPlot[451].X = 4.51
	pointsOfFunctionPlot[451].Y = 11.794

	pointsOfFunctionPlot[452].X = 4.52
	pointsOfFunctionPlot[452].Y = 11.959

	pointsOfFunctionPlot[453].X = 4.53
	pointsOfFunctionPlot[453].Y = 12.127

	pointsOfFunctionPlot[454].X = 4.54
	pointsOfFunctionPlot[454].Y = 12.298

	pointsOfFunctionPlot[455].X = 4.55
	pointsOfFunctionPlot[455].Y = 12.472

	pointsOfFunctionPlot[456].X = 4.56
	pointsOfFunctionPlot[456].Y = 12.648

	pointsOfFunctionPlot[457].X = 4.57
	pointsOfFunctionPlot[457].Y = 12.827

	pointsOfFunctionPlot[458].X = 4.58
	pointsOfFunctionPlot[458].Y = 13.008

	pointsOfFunctionPlot[459].X = 4.59
	pointsOfFunctionPlot[459].Y = 13.193

	pointsOfFunctionPlot[460].X = 4.60
	pointsOfFunctionPlot[460].Y = 13.381

	pointsOfFunctionPlot[461].X = 4.61
	pointsOfFunctionPlot[461].Y = 13.571

	pointsOfFunctionPlot[462].X = 4.62
	pointsOfFunctionPlot[462].Y = 13.765

	pointsOfFunctionPlot[463].X = 4.63
	pointsOfFunctionPlot[463].Y = 13.962

	pointsOfFunctionPlot[464].X = 4.64
	pointsOfFunctionPlot[464].Y = 14.162

	pointsOfFunctionPlot[465].X = 4.65
	pointsOfFunctionPlot[465].Y = 14.365

	pointsOfFunctionPlot[466].X = 4.66
	pointsOfFunctionPlot[466].Y = 14.571

	pointsOfFunctionPlot[467].X = 4.67
	pointsOfFunctionPlot[467].Y = 14.781

	pointsOfFunctionPlot[468].X = 4.68
	pointsOfFunctionPlot[468].Y = 14.994

	pointsOfFunctionPlot[469].X = 4.69
	pointsOfFunctionPlot[469].Y = 15.211

	pointsOfFunctionPlot[470].X = 4.70
	pointsOfFunctionPlot[470].Y = 15.431

	pointsOfFunctionPlot[471].X = 4.71
	pointsOfFunctionPlot[471].Y = 15.655

	pointsOfFunctionPlot[472].X = 4.72
	pointsOfFunctionPlot[472].Y = 15.882

	pointsOfFunctionPlot[473].X = 4.73
	pointsOfFunctionPlot[473].Y = 16.113

	pointsOfFunctionPlot[474].X = 4.74
	pointsOfFunctionPlot[474].Y = 16.347

	pointsOfFunctionPlot[475].X = 4.75
	pointsOfFunctionPlot[475].Y = 16.586

	pointsOfFunctionPlot[476].X = 4.76
	pointsOfFunctionPlot[476].Y = 16.828

	pointsOfFunctionPlot[477].X = 4.77
	pointsOfFunctionPlot[477].Y = 17.074

	pointsOfFunctionPlot[478].X = 4.78
	pointsOfFunctionPlot[478].Y = 17.325

	pointsOfFunctionPlot[479].X = 4.79
	pointsOfFunctionPlot[479].Y = 17.579

	pointsOfFunctionPlot[480].X = 4.80
	pointsOfFunctionPlot[480].Y = 17.837

	pointsOfFunctionPlot[481].X = 4.81
	pointsOfFunctionPlot[481].Y = 18.1

	pointsOfFunctionPlot[482].X = 4.82
	pointsOfFunctionPlot[482].Y = 18.367

	pointsOfFunctionPlot[483].X = 4.83
	pointsOfFunctionPlot[483].Y = 18.638

	pointsOfFunctionPlot[484].X = 4.84
	pointsOfFunctionPlot[484].Y = 18.914

	pointsOfFunctionPlot[485].X = 4.85
	pointsOfFunctionPlot[485].Y = 19.195

	pointsOfFunctionPlot[486].X = 4.86
	pointsOfFunctionPlot[486].Y = 19.48

	pointsOfFunctionPlot[487].X = 4.87
	pointsOfFunctionPlot[487].Y = 19.769

	pointsOfFunctionPlot[488].X = 4.88
	pointsOfFunctionPlot[488].Y = 20.064

	pointsOfFunctionPlot[489].X = 4.89
	pointsOfFunctionPlot[489].Y = 20.363

	pointsOfFunctionPlot[490].X = 4.90
	pointsOfFunctionPlot[490].Y = 20.667

	pointsOfFunctionPlot[491].X = 4.91
	pointsOfFunctionPlot[491].Y = 20.976

	pointsOfFunctionPlot[492].X = 4.92
	pointsOfFunctionPlot[492].Y = 21.29

	pointsOfFunctionPlot[493].X = 4.93
	pointsOfFunctionPlot[493].Y = 21.61

	pointsOfFunctionPlot[494].X = 4.94
	pointsOfFunctionPlot[494].Y = 21.935

	pointsOfFunctionPlot[495].X = 4.95
	pointsOfFunctionPlot[495].Y = 22.265

	pointsOfFunctionPlot[496].X = 4.96
	pointsOfFunctionPlot[496].Y = 22.6

	pointsOfFunctionPlot[497].X = 4.97
	pointsOfFunctionPlot[497].Y = 22.942

	pointsOfFunctionPlot[498].X = 4.98
	pointsOfFunctionPlot[498].Y = 23.288

	pointsOfFunctionPlot[499].X = 4.99
	pointsOfFunctionPlot[499].Y = 23.641

	pointsOfFunctionPlot[500].X = 5.0
	pointsOfFunctionPlot[500].Y = 24.0

	pointsOfFunctionPlot[501].X = 5.01
	pointsOfFunctionPlot[501].Y = 24.364

	pointsOfFunctionPlot[502].X = 5.02
	pointsOfFunctionPlot[502].Y = 24.735

	pointsOfFunctionPlot[503].X = 5.03
	pointsOfFunctionPlot[503].Y = 25.111

	pointsOfFunctionPlot[504].X = 5.04
	pointsOfFunctionPlot[504].Y = 25.494

	pointsOfFunctionPlot[505].X = 5.05
	pointsOfFunctionPlot[505].Y = 25.884

	pointsOfFunctionPlot[506].X = 5.06
	pointsOfFunctionPlot[506].Y = 26.28

	pointsOfFunctionPlot[507].X = 5.07
	pointsOfFunctionPlot[507].Y = 26.682

	pointsOfFunctionPlot[508].X = 5.08
	pointsOfFunctionPlot[508].Y = 27.092

	pointsOfFunctionPlot[509].X = 5.09
	pointsOfFunctionPlot[509].Y = 27.508

	pointsOfFunctionPlot[510].X = 5.10
	pointsOfFunctionPlot[510].Y = 27.931

	pointsOfFunctionPlot[511].X = 5.11
	pointsOfFunctionPlot[511].Y = 28.362

	pointsOfFunctionPlot[512].X = 5.12
	pointsOfFunctionPlot[512].Y = 28.799

	pointsOfFunctionPlot[513].X = 5.13
	pointsOfFunctionPlot[513].Y = 29.244

	pointsOfFunctionPlot[514].X = 5.14
	pointsOfFunctionPlot[514].Y = 29.697

	pointsOfFunctionPlot[515].X = 5.15
	pointsOfFunctionPlot[515].Y = 30.157

	pointsOfFunctionPlot[516].X = 5.16
	pointsOfFunctionPlot[516].Y = 30.625

	pointsOfFunctionPlot[517].X = 5.17
	pointsOfFunctionPlot[517].Y = 31.101

	pointsOfFunctionPlot[518].X = 5.18
	pointsOfFunctionPlot[518].Y = 31.585

	pointsOfFunctionPlot[519].X = 5.19
	pointsOfFunctionPlot[519].Y = 32.077

	pointsOfFunctionPlot[520].X = 5.20
	pointsOfFunctionPlot[520].Y = 32.578

	pointsOfFunctionPlot[521].X = 5.21
	pointsOfFunctionPlot[521].Y = 33.087

	pointsOfFunctionPlot[522].X = 5.22
	pointsOfFunctionPlot[522].Y = 33.604

	pointsOfFunctionPlot[523].X = 5.23
	pointsOfFunctionPlot[523].Y = 34.131

	pointsOfFunctionPlot[524].X = 5.24
	pointsOfFunctionPlot[524].Y = 34.666

	pointsOfFunctionPlot[525].X = 5.25
	pointsOfFunctionPlot[525].Y = 35.211

	pointsOfFunctionPlot[526].X = 5.26
	pointsOfFunctionPlot[526].Y = 35.765

	pointsOfFunctionPlot[527].X = 5.27
	pointsOfFunctionPlot[527].Y = 36.329

	pointsOfFunctionPlot[528].X = 5.28
	pointsOfFunctionPlot[528].Y = 36.902

	pointsOfFunctionPlot[529].X = 5.29
	pointsOfFunctionPlot[529].Y = 37.485

	pointsOfFunctionPlot[530].X = 5.30
	pointsOfFunctionPlot[530].Y = 38.077

	pointsOfFunctionPlot[531].X = 5.31
	pointsOfFunctionPlot[531].Y = 38.681

	pointsOfFunctionPlot[532].X = 5.32
	pointsOfFunctionPlot[532].Y = 39.294

	pointsOfFunctionPlot[533].X = 5.33
	pointsOfFunctionPlot[533].Y = 39.918

	pointsOfFunctionPlot[534].X = 5.34
	pointsOfFunctionPlot[534].Y = 40.553

	pointsOfFunctionPlot[535].X = 5.35
	pointsOfFunctionPlot[535].Y = 41.199

	pointsOfFunctionPlot[536].X = 5.36
	pointsOfFunctionPlot[536].Y = 41.855

	pointsOfFunctionPlot[537].X = 5.37
	pointsOfFunctionPlot[537].Y = 42.524

	pointsOfFunctionPlot[538].X = 5.38
	pointsOfFunctionPlot[538].Y = 43.203

	pointsOfFunctionPlot[539].X = 5.39
	pointsOfFunctionPlot[539].Y = 43.895

	pointsOfFunctionPlot[540].X = 5.40
	pointsOfFunctionPlot[540].Y = 44.598

	pointsOfFunctionPlot[541].X = 5.41
	pointsOfFunctionPlot[541].Y = 45.314

	pointsOfFunctionPlot[542].X = 5.42
	pointsOfFunctionPlot[542].Y = 46.042

	pointsOfFunctionPlot[543].X = 5.43
	pointsOfFunctionPlot[543].Y = 46.783

	pointsOfFunctionPlot[544].X = 5.44
	pointsOfFunctionPlot[544].Y = 47.537

	pointsOfFunctionPlot[545].X = 5.45
	pointsOfFunctionPlot[545].Y = 48.303

	pointsOfFunctionPlot[546].X = 5.46
	pointsOfFunctionPlot[546].Y = 49.083

	pointsOfFunctionPlot[547].X = 5.47
	pointsOfFunctionPlot[547].Y = 49.877

	pointsOfFunctionPlot[548].X = 5.48
	pointsOfFunctionPlot[548].Y = 50.685

	pointsOfFunctionPlot[549].X = 5.49
	pointsOfFunctionPlot[549].Y = 51.506

	pointsOfFunctionPlot[550].X = 5.50
	pointsOfFunctionPlot[550].Y = 52.342

	pointsOfFunctionPlot[551].X = 5.51
	pointsOfFunctionPlot[551].Y = 53.193

	pointsOfFunctionPlot[552].X = 5.52
	pointsOfFunctionPlot[552].Y = 54.058

	pointsOfFunctionPlot[553].X = 5.53
	pointsOfFunctionPlot[553].Y = 54.939

	pointsOfFunctionPlot[554].X = 5.54
	pointsOfFunctionPlot[554].Y = 55.835

	pointsOfFunctionPlot[555].X = 5.55
	pointsOfFunctionPlot[555].Y = 56.747

	pointsOfFunctionPlot[556].X = 5.56
	pointsOfFunctionPlot[556].Y = 57.675

	pointsOfFunctionPlot[557].X = 5.57
	pointsOfFunctionPlot[557].Y = 58.62

	pointsOfFunctionPlot[558].X = 5.58
	pointsOfFunctionPlot[558].Y = 59.58

	pointsOfFunctionPlot[559].X = 5.59
	pointsOfFunctionPlot[559].Y = 60.558

	pointsOfFunctionPlot[560].X = 5.60
	pointsOfFunctionPlot[560].Y = 61.553

	pointsOfFunctionPlot[561].X = 5.61
	pointsOfFunctionPlot[561].Y = 62.566

	pointsOfFunctionPlot[562].X = 5.62
	pointsOfFunctionPlot[562].Y = 63.597

	pointsOfFunctionPlot[563].X = 5.63
	pointsOfFunctionPlot[563].Y = 64.645

	pointsOfFunctionPlot[564].X = 5.64
	pointsOfFunctionPlot[564].Y = 65.713

	pointsOfFunctionPlot[565].X = 5.65
	pointsOfFunctionPlot[565].Y = 66.799

	pointsOfFunctionPlot[566].X = 5.66
	pointsOfFunctionPlot[566].Y = 67.905

	pointsOfFunctionPlot[567].X = 5.67
	pointsOfFunctionPlot[567].Y = 69.03

	pointsOfFunctionPlot[568].X = 5.68
	pointsOfFunctionPlot[568].Y = 70.175

	pointsOfFunctionPlot[569].X = 5.69
	pointsOfFunctionPlot[569].Y = 71.341

	pointsOfFunctionPlot[570].X = 5.70
	pointsOfFunctionPlot[570].Y = 72.527

	pointsOfFunctionPlot[571].X = 5.71
	pointsOfFunctionPlot[571].Y = 73.735

	pointsOfFunctionPlot[572].X = 5.72
	pointsOfFunctionPlot[572].Y = 74.964

	pointsOfFunctionPlot[573].X = 5.73
	pointsOfFunctionPlot[573].Y = 76.215

	pointsOfFunctionPlot[574].X = 5.74
	pointsOfFunctionPlot[574].Y = 77.488

	pointsOfFunctionPlot[575].X = 5.75
	pointsOfFunctionPlot[575].Y = 78.784

	pointsOfFunctionPlot[576].X = 5.76
	pointsOfFunctionPlot[576].Y = 80.103

	pointsOfFunctionPlot[577].X = 5.77
	pointsOfFunctionPlot[577].Y = 81.446

	pointsOfFunctionPlot[578].X = 5.78
	pointsOfFunctionPlot[578].Y = 82.813

	pointsOfFunctionPlot[579].X = 5.79
	pointsOfFunctionPlot[579].Y = 84.205

	pointsOfFunctionPlot[580].X = 5.80
	pointsOfFunctionPlot[580].Y = 85.621

	pointsOfFunctionPlot[581].X = 5.81
	pointsOfFunctionPlot[581].Y = 87.063

	pointsOfFunctionPlot[582].X = 5.82
	pointsOfFunctionPlot[582].Y = 88.531

	pointsOfFunctionPlot[583].X = 5.83
	pointsOfFunctionPlot[583].Y = 90.026

	pointsOfFunctionPlot[584].X = 5.84
	pointsOfFunctionPlot[584].Y = 91.547

	pointsOfFunctionPlot[585].X = 5.85
	pointsOfFunctionPlot[585].Y = 93.096

	pointsOfFunctionPlot[586].X = 5.86
	pointsOfFunctionPlot[586].Y = 94.672

	pointsOfFunctionPlot[587].X = 5.87
	pointsOfFunctionPlot[587].Y = 96.278

	pointsOfFunctionPlot[588].X = 5.88
	pointsOfFunctionPlot[588].Y = 97.912

	pointsOfFunctionPlot[589].X = 5.89
	pointsOfFunctionPlot[589].Y = 99.576

	pointsOfFunctionPlot[590].X = 5.90
	pointsOfFunctionPlot[590].Y = 101.27

	pointsOfFunctionPlot[591].X = 5.91
	pointsOfFunctionPlot[591].Y = 102.994

	pointsOfFunctionPlot[592].X = 5.92
	pointsOfFunctionPlot[592].Y = 104.75

	pointsOfFunctionPlot[593].X = 5.93
	pointsOfFunctionPlot[593].Y = 106.538

	pointsOfFunctionPlot[594].X = 5.94
	pointsOfFunctionPlot[594].Y = 108.359

	pointsOfFunctionPlot[595].X = 5.95
	pointsOfFunctionPlot[595].Y = 110.212

	pointsOfFunctionPlot[596].X = 5.96
	pointsOfFunctionPlot[596].Y = 112.1

	pointsOfFunctionPlot[597].X = 5.97
	pointsOfFunctionPlot[597].Y = 114.021

	pointsOfFunctionPlot[598].X = 5.98
	pointsOfFunctionPlot[598].Y = 115.978

	pointsOfFunctionPlot[599].X = 5.99
	pointsOfFunctionPlot[599].Y = 117.971

	pointsOfFunctionPlot[600].X = 6.0
	pointsOfFunctionPlot[600].Y = 120.0

	pointsOfFunctionPlot[601].X = 6.01
	pointsOfFunctionPlot[601].Y = 122.066

	pointsOfFunctionPlot[602].X = 6.02
	pointsOfFunctionPlot[602].Y = 124.169

	pointsOfFunctionPlot[603].X = 6.03
	pointsOfFunctionPlot[603].Y = 126.312

	pointsOfFunctionPlot[604].X = 6.04
	pointsOfFunctionPlot[604].Y = 128.493

	pointsOfFunctionPlot[605].X = 6.05
	pointsOfFunctionPlot[605].Y = 130.715

	pointsOfFunctionPlot[606].X = 6.06
	pointsOfFunctionPlot[606].Y = 132.978

	pointsOfFunctionPlot[607].X = 6.07
	pointsOfFunctionPlot[607].Y = 135.282

	pointsOfFunctionPlot[608].X = 6.08
	pointsOfFunctionPlot[608].Y = 137.628

	pointsOfFunctionPlot[609].X = 6.09
	pointsOfFunctionPlot[609].Y = 140.018

	pointsOfFunctionPlot[610].X = 6.10
	pointsOfFunctionPlot[610].Y = 142.451

	pointsOfFunctionPlot[611].X = 6.11
	pointsOfFunctionPlot[611].Y = 144.93

	pointsOfFunctionPlot[612].X = 6.12
	pointsOfFunctionPlot[612].Y = 147.454

	pointsOfFunctionPlot[613].X = 6.13
	pointsOfFunctionPlot[613].Y = 150.025

	pointsOfFunctionPlot[614].X = 6.14
	pointsOfFunctionPlot[614].Y = 152.644

	pointsOfFunctionPlot[615].X = 6.15
	pointsOfFunctionPlot[615].Y = 155.311

	pointsOfFunctionPlot[616].X = 6.16
	pointsOfFunctionPlot[616].Y = 158.027

	pointsOfFunctionPlot[617].X = 6.17
	pointsOfFunctionPlot[617].Y = 160.794

	pointsOfFunctionPlot[618].X = 6.18
	pointsOfFunctionPlot[618].Y = 163.612

	pointsOfFunctionPlot[619].X = 6.19
	pointsOfFunctionPlot[619].Y = 166.482

	pointsOfFunctionPlot[620].X = 6.20
	pointsOfFunctionPlot[620].Y = 169.406

	pointsOfFunctionPlot[621].X = 6.21
	pointsOfFunctionPlot[621].Y = 172.384

	pointsOfFunctionPlot[622].X = 6.22
	pointsOfFunctionPlot[622].Y = 175.417

	pointsOfFunctionPlot[623].X = 6.23
	pointsOfFunctionPlot[623].Y = 178.507

	pointsOfFunctionPlot[624].X = 6.24
	pointsOfFunctionPlot[624].Y = 181.654

	pointsOfFunctionPlot[625].X = 6.25
	pointsOfFunctionPlot[625].Y = 184.86

	pointsOfFunctionPlot[626].X = 6.26
	pointsOfFunctionPlot[626].Y = 188.126

	pointsOfFunctionPlot[627].X = 6.27
	pointsOfFunctionPlot[627].Y = 191.454

	pointsOfFunctionPlot[628].X = 6.28
	pointsOfFunctionPlot[628].Y = 194.843

	pointsOfFunctionPlot[629].X = 6.29
	pointsOfFunctionPlot[629].Y = 198.295

	pointsOfFunctionPlot[630].X = 6.30
	pointsOfFunctionPlot[630].Y = 201.813

	pointsOfFunctionPlot[631].X = 6.31
	pointsOfFunctionPlot[631].Y = 205.396

	pointsOfFunctionPlot[632].X = 6.32
	pointsOfFunctionPlot[632].Y = 209.046

	pointsOfFunctionPlot[633].X = 6.33
	pointsOfFunctionPlot[633].Y = 212.765

	pointsOfFunctionPlot[634].X = 6.34
	pointsOfFunctionPlot[634].Y = 216.554

	pointsOfFunctionPlot[635].X = 6.35
	pointsOfFunctionPlot[635].Y = 220.414

	pointsOfFunctionPlot[636].X = 6.36
	pointsOfFunctionPlot[636].Y = 224.347

	pointsOfFunctionPlot[637].X = 6.37
	pointsOfFunctionPlot[637].Y = 228.354

	pointsOfFunctionPlot[638].X = 6.38
	pointsOfFunctionPlot[638].Y = 232.436

	pointsOfFunctionPlot[639].X = 6.39
	pointsOfFunctionPlot[639].Y = 236.595

	pointsOfFunctionPlot[640].X = 6.40
	pointsOfFunctionPlot[640].Y = 240.833

	pointsOfFunctionPlot[641].X = 6.41
	pointsOfFunctionPlot[641].Y = 245.5151

	pointsOfFunctionPlot[642].X = 6.42
	pointsOfFunctionPlot[642].Y = 249.551

	pointsOfFunctionPlot[643].X = 6.43
	pointsOfFunctionPlot[643].Y = 254.033

	pointsOfFunctionPlot[644].X = 6.44
	pointsOfFunctionPlot[644].Y = 258.601

	pointsOfFunctionPlot[645].X = 6.45
	pointsOfFunctionPlot[645].Y = 263.255

	pointsOfFunctionPlot[646].X = 6.46
	pointsOfFunctionPlot[646].Y = 267.997

	pointsOfFunctionPlot[647].X = 6.47
	pointsOfFunctionPlot[647].Y = 272.83

	pointsOfFunctionPlot[648].X = 6.48
	pointsOfFunctionPlot[648].Y = 277.754

	pointsOfFunctionPlot[649].X = 6.49
	pointsOfFunctionPlot[649].Y = 282.772

	pointsOfFunctionPlot[650].X = 6.50
	pointsOfFunctionPlot[650].Y = 287.885

	pointsOfFunctionPlot[651].X = 6.51
	pointsOfFunctionPlot[651].Y = 293.095

	pointsOfFunctionPlot[652].X = 6.52
	pointsOfFunctionPlot[652].Y = 298.405

	pointsOfFunctionPlot[653].X = 6.53
	pointsOfFunctionPlot[653].Y = 303.816

	pointsOfFunctionPlot[654].X = 6.54
	pointsOfFunctionPlot[654].Y = 309.33

	pointsOfFunctionPlot[655].X = 6.55
	pointsOfFunctionPlot[655].Y = 314.95

	pointsOfFunctionPlot[656].X = 6.56
	pointsOfFunctionPlot[656].Y = 320.677

	pointsOfFunctionPlot[657].X = 6.57
	pointsOfFunctionPlot[657].Y = 326.513

	pointsOfFunctionPlot[658].X = 6.58
	pointsOfFunctionPlot[658].Y = 332.461

	pointsOfFunctionPlot[659].X = 6.59
	pointsOfFunctionPlot[659].Y = 338.523

	pointsOfFunctionPlot[660].X = 6.60
	pointsOfFunctionPlot[660].Y = 344.701

	pointsOfFunctionPlot[661].X = 6.61
	pointsOfFunctionPlot[661].Y = 350.998

	pointsOfFunctionPlot[662].X = 6.62
	pointsOfFunctionPlot[662].Y = 357.416

	pointsOfFunctionPlot[663].X = 6.63
	pointsOfFunctionPlot[663].Y = 363.956

	pointsOfFunctionPlot[664].X = 6.64
	pointsOfFunctionPlot[664].Y = 370.623

	pointsOfFunctionPlot[665].X = 6.65
	pointsOfFunctionPlot[665].Y = 377.418

	pointsOfFunctionPlot[666].X = 6.66
	pointsOfFunctionPlot[666].Y = 384.343

	pointsOfFunctionPlot[667].X = 6.67
	pointsOfFunctionPlot[667].Y = 391.402

	pointsOfFunctionPlot[668].X = 6.68
	pointsOfFunctionPlot[668].Y = 398.598

	pointsOfFunctionPlot[669].X = 6.69
	pointsOfFunctionPlot[669].Y = 405.931

	pointsOfFunctionPlot[670].X = 6.70
	pointsOfFunctionPlot[670].Y = 413.407

	pointsOfFunctionPlot[671].X = 6.71
	pointsOfFunctionPlot[671].Y = 421.027

	pointsOfFunctionPlot[672].X = 6.72
	pointsOfFunctionPlot[672].Y = 428.794

	pointsOfFunctionPlot[673].X = 6.73
	pointsOfFunctionPlot[673].Y = 436.712

	pointsOfFunctionPlot[674].X = 6.74
	pointsOfFunctionPlot[674].Y = 444.783

	pointsOfFunctionPlot[675].X = 6.75
	pointsOfFunctionPlot[675].Y = 453.01

	pointsOfFunctionPlot[676].X = 6.76
	pointsOfFunctionPlot[676].Y = 461.397

	pointsOfFunctionPlot[677].X = 6.77
	pointsOfFunctionPlot[677].Y = 469.947

	pointsOfFunctionPlot[678].X = 6.78
	pointsOfFunctionPlot[678].Y = 478.663

	pointsOfFunctionPlot[679].X = 6.79
	pointsOfFunctionPlot[679].Y = 487.548

	pointsOfFunctionPlot[680].X = 6.80
	pointsOfFunctionPlot[680].Y = 496.606

	pointsOfFunctionPlot[681].X = 6.81
	pointsOfFunctionPlot[681].Y = 505.84

	pointsOfFunctionPlot[682].X = 6.82
	pointsOfFunctionPlot[682].Y = 515.254

	pointsOfFunctionPlot[683].X = 6.83
	pointsOfFunctionPlot[683].Y = 524.851

	pointsOfFunctionPlot[684].X = 6.84
	pointsOfFunctionPlot[684].Y = 534.636

	pointsOfFunctionPlot[685].X = 6.85
	pointsOfFunctionPlot[685].Y = 544.612

	pointsOfFunctionPlot[686].X = 6.86
	pointsOfFunctionPlot[686].Y = 554.782

	pointsOfFunctionPlot[687].X = 6.87
	pointsOfFunctionPlot[687].Y = 565.152

	pointsOfFunctionPlot[688].X = 6.88
	pointsOfFunctionPlot[688].Y = 575.724

	pointsOfFunctionPlot[689].X = 6.89
	pointsOfFunctionPlot[689].Y = 586.503

	pointsOfFunctionPlot[690].X = 6.90
	pointsOfFunctionPlot[690].Y = 597.494

	pointsOfFunctionPlot[691].X = 6.91
	pointsOfFunctionPlot[691].Y = 608.699

	pointsOfFunctionPlot[692].X = 6.92
	pointsOfFunctionPlot[692].Y = 620.125

	pointsOfFunctionPlot[693].X = 6.93
	pointsOfFunctionPlot[693].Y = 631.775

	pointsOfFunctionPlot[694].X = 6.94
	pointsOfFunctionPlot[694].Y = 643.654

	pointsOfFunctionPlot[695].X = 6.95
	pointsOfFunctionPlot[695].Y = 655.766

	pointsOfFunctionPlot[696].X = 6.96
	pointsOfFunctionPlot[696].Y = 668.116

	pointsOfFunctionPlot[697].X = 6.97
	pointsOfFunctionPlot[697].Y = 680.71

	pointsOfFunctionPlot[698].X = 6.98
	pointsOfFunctionPlot[698].Y = 693.552

	pointsOfFunctionPlot[699].X = 6.99
	pointsOfFunctionPlot[699].Y = 706.646

	pointsOfFunctionPlot[700].X = 7.0
	pointsOfFunctionPlot[700].Y = 720.0










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of Gamma function"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("Gamma(x)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"Gamma-function-plot-01.png"); err != nil {

		panic(err)
	}
}
