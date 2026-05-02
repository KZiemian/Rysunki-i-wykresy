package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function f(x) = x^(1/12).

	pointsOfFunctionPlot := make(plotter.XYs, 1_001)

	pointsOfFunctionPlot[0].X = 0.0
	pointsOfFunctionPlot[0].Y = 0.0

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = 0.681

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = 0.721

	pointsOfFunctionPlot[3].X = 0.03
	pointsOfFunctionPlot[3].Y = 0.746

	pointsOfFunctionPlot[4].X = 0.04
	pointsOfFunctionPlot[4].Y = 0.764

	pointsOfFunctionPlot[5].X = 0.05
	pointsOfFunctionPlot[5].Y = 0.779

	pointsOfFunctionPlot[6].X = 0.06
	pointsOfFunctionPlot[6].Y = 0.791

	pointsOfFunctionPlot[7].X = 0.07
	pointsOfFunctionPlot[7].Y = 0.801

	pointsOfFunctionPlot[8].X = 0.08
	pointsOfFunctionPlot[8].Y = 0.81

	pointsOfFunctionPlot[9].X = 0.09
	pointsOfFunctionPlot[9].Y = 0.818

	pointsOfFunctionPlot[10].X = 0.10
	pointsOfFunctionPlot[10].Y = 0.825

	pointsOfFunctionPlot[11].X = 0.11
	pointsOfFunctionPlot[11].Y = 0.831

	pointsOfFunctionPlot[12].X = 0.12
	pointsOfFunctionPlot[12].Y = 0.838

	pointsOfFunctionPlot[13].X = 0.13
	pointsOfFunctionPlot[13].Y = 0.843

	pointsOfFunctionPlot[14].X = 0.14
	pointsOfFunctionPlot[14].Y = 0.848

	pointsOfFunctionPlot[15].X = 0.15
	pointsOfFunctionPlot[15].Y = 0.853

	pointsOfFunctionPlot[16].X = 0.16
	pointsOfFunctionPlot[16].Y = 0.858

	pointsOfFunctionPlot[17].X = 0.17
	pointsOfFunctionPlot[17].Y = 0.862

	pointsOfFunctionPlot[18].X = 0.18
	pointsOfFunctionPlot[18].Y = 0.866

	pointsOfFunctionPlot[19].X = 0.19
	pointsOfFunctionPlot[19].Y = 0.87

	pointsOfFunctionPlot[20].X = 0.20
	pointsOfFunctionPlot[20].Y = 0.874

	pointsOfFunctionPlot[21].X = 0.21
	pointsOfFunctionPlot[21].Y = 0.878

	pointsOfFunctionPlot[22].X = 0.22
	pointsOfFunctionPlot[22].Y = 0.881

	pointsOfFunctionPlot[23].X = 0.23
	pointsOfFunctionPlot[23].Y = 0.884

	pointsOfFunctionPlot[24].X = 0.24
	pointsOfFunctionPlot[24].Y = 0.887

	pointsOfFunctionPlot[25].X = 0.25
	pointsOfFunctionPlot[25].Y = 0.89

	pointsOfFunctionPlot[26].X = 0.26
	pointsOfFunctionPlot[26].Y = 0.893

	pointsOfFunctionPlot[27].X = 0.27
	pointsOfFunctionPlot[27].Y = 0.896

	pointsOfFunctionPlot[28].X = 0.28
	pointsOfFunctionPlot[28].Y = 0.899

	pointsOfFunctionPlot[29].X = 0.29
	pointsOfFunctionPlot[29].Y = 0.901

	pointsOfFunctionPlot[30].X = 0.30
	pointsOfFunctionPlot[30].Y = 0.904

	pointsOfFunctionPlot[31].X = 0.31
	pointsOfFunctionPlot[31].Y = 0.907

	pointsOfFunctionPlot[32].X = 0.32
	pointsOfFunctionPlot[32].Y = 0.909

	pointsOfFunctionPlot[33].X = 0.33
	pointsOfFunctionPlot[33].Y = 0.911

	pointsOfFunctionPlot[34].X = 0.34
	pointsOfFunctionPlot[34].Y = 0.914

	pointsOfFunctionPlot[35].X = 0.35
	pointsOfFunctionPlot[35].Y = 0.916

	pointsOfFunctionPlot[36].X = 0.36
	pointsOfFunctionPlot[36].Y = 0.918

	pointsOfFunctionPlot[37].X = 0.37
	pointsOfFunctionPlot[37].Y = 0.92

	pointsOfFunctionPlot[38].X = 0.38
	pointsOfFunctionPlot[38].Y = 0.922

	pointsOfFunctionPlot[39].X = 0.39
	pointsOfFunctionPlot[39].Y = 0.924

	pointsOfFunctionPlot[40].X = 0.40
	pointsOfFunctionPlot[40].Y = 0.926

	pointsOfFunctionPlot[41].X = 0.41
	pointsOfFunctionPlot[41].Y = 0.928

	pointsOfFunctionPlot[42].X = 0.42
	pointsOfFunctionPlot[42].Y = 0.93

	pointsOfFunctionPlot[43].X = 0.43
	pointsOfFunctionPlot[43].Y = 0.932

	pointsOfFunctionPlot[44].X = 0.44
	pointsOfFunctionPlot[44].Y = 0.933

	pointsOfFunctionPlot[45].X = 0.45
	pointsOfFunctionPlot[45].Y = 0.935

	pointsOfFunctionPlot[46].X = 0.46
	pointsOfFunctionPlot[46].Y = 0.937

	pointsOfFunctionPlot[47].X = 0.47
	pointsOfFunctionPlot[47].Y = 0.939

	pointsOfFunctionPlot[48].X = 0.48
	pointsOfFunctionPlot[48].Y = 0.94

	pointsOfFunctionPlot[49].X = 0.49
	pointsOfFunctionPlot[49].Y = 0.942

	pointsOfFunctionPlot[50].X = 0.50
	pointsOfFunctionPlot[50].Y = 0.943

	pointsOfFunctionPlot[51].X = 0.51
	pointsOfFunctionPlot[51].Y = 0.945

	pointsOfFunctionPlot[52].X = 0.52
	pointsOfFunctionPlot[52].Y = 0.946

	pointsOfFunctionPlot[53].X = 0.53
	pointsOfFunctionPlot[53].Y = 0.948

	pointsOfFunctionPlot[54].X = 0.54
	pointsOfFunctionPlot[54].Y = 0.949

	pointsOfFunctionPlot[55].X = 0.55
	pointsOfFunctionPlot[55].Y = 0.951

	pointsOfFunctionPlot[56].X = 0.56
	pointsOfFunctionPlot[56].Y = 0.952

	pointsOfFunctionPlot[57].X = 0.57
	pointsOfFunctionPlot[57].Y = 0.954

	pointsOfFunctionPlot[58].X = 0.58
	pointsOfFunctionPlot[58].Y = 0.955

	pointsOfFunctionPlot[59].X = 0.59
	pointsOfFunctionPlot[59].Y = 0.956

	pointsOfFunctionPlot[60].X = 0.60
	pointsOfFunctionPlot[60].Y = 0.958

	pointsOfFunctionPlot[61].X = 0.61
	pointsOfFunctionPlot[61].Y = 0.959

	pointsOfFunctionPlot[62].X = 0.62
	pointsOfFunctionPlot[62].Y = 0.96

	pointsOfFunctionPlot[63].X = 0.63
	pointsOfFunctionPlot[63].Y = 0.962

	pointsOfFunctionPlot[64].X = 0.64
	pointsOfFunctionPlot[64].Y = 0.963

	pointsOfFunctionPlot[65].X = 0.65
	pointsOfFunctionPlot[65].Y = 0.964

	pointsOfFunctionPlot[66].X = 0.66
	pointsOfFunctionPlot[66].Y = 0.965

	pointsOfFunctionPlot[67].X = 0.67
	pointsOfFunctionPlot[67].Y = 0.967

	pointsOfFunctionPlot[68].X = 0.68
	pointsOfFunctionPlot[68].Y = 0.968

	pointsOfFunctionPlot[69].X = 0.69
	pointsOfFunctionPlot[69].Y = 0.969

	pointsOfFunctionPlot[70].X = 0.70
	pointsOfFunctionPlot[70].Y = 0.97

	pointsOfFunctionPlot[71].X = 0.71
	pointsOfFunctionPlot[71].Y = 0.971

	pointsOfFunctionPlot[72].X = 0.72
	pointsOfFunctionPlot[72].Y = 0.972

	pointsOfFunctionPlot[73].X = 0.73
	pointsOfFunctionPlot[73].Y = 0.974

	pointsOfFunctionPlot[74].X = 0.74
	pointsOfFunctionPlot[74].Y = 0.975

	pointsOfFunctionPlot[75].X = 0.75
	pointsOfFunctionPlot[75].Y = 0.976

	pointsOfFunctionPlot[76].X = 0.76
	pointsOfFunctionPlot[76].Y = 0.977

	pointsOfFunctionPlot[77].X = 0.77
	pointsOfFunctionPlot[77].Y = 0.978

	pointsOfFunctionPlot[78].X = 0.78
	pointsOfFunctionPlot[78].Y = 0.979

	pointsOfFunctionPlot[79].X = 0.79
	pointsOfFunctionPlot[79].Y = 0.98

	pointsOfFunctionPlot[80].X = 0.80
	pointsOfFunctionPlot[80].Y = 0.981

	pointsOfFunctionPlot[81].X = 0.81
	pointsOfFunctionPlot[81].Y = 0.982

	pointsOfFunctionPlot[82].X = 0.82
	pointsOfFunctionPlot[82].Y = 0.983

	pointsOfFunctionPlot[83].X = 0.83
	pointsOfFunctionPlot[83].Y = 0.984

	pointsOfFunctionPlot[84].X = 0.84
	pointsOfFunctionPlot[84].Y = 0.985

	pointsOfFunctionPlot[85].X = 0.85
	pointsOfFunctionPlot[85].Y = 0.986

	pointsOfFunctionPlot[86].X = 0.86
	pointsOfFunctionPlot[86].Y = 0.987

	pointsOfFunctionPlot[87].X = 0.87
	pointsOfFunctionPlot[87].Y = 0.988

	pointsOfFunctionPlot[88].X = 0.88
	pointsOfFunctionPlot[88].Y = 0.989

	pointsOfFunctionPlot[89].X = 0.89
	pointsOfFunctionPlot[89].Y = 0.99

	pointsOfFunctionPlot[90].X = 0.90
	pointsOfFunctionPlot[90].Y = 0.991

	pointsOfFunctionPlot[91].X = 0.91
	pointsOfFunctionPlot[91].Y = 0.992

	pointsOfFunctionPlot[92].X = 0.92
	pointsOfFunctionPlot[92].Y = 0.993

	pointsOfFunctionPlot[93].X = 0.93
	pointsOfFunctionPlot[93].Y = 0.993

	pointsOfFunctionPlot[94].X = 0.94
	pointsOfFunctionPlot[94].Y = 0.994

	pointsOfFunctionPlot[95].X = 0.95
	pointsOfFunctionPlot[95].Y = 0.995

	pointsOfFunctionPlot[96].X = 0.96
	pointsOfFunctionPlot[96].Y = 0.996

	pointsOfFunctionPlot[97].X = 0.97
	pointsOfFunctionPlot[97].Y = 0.997

	pointsOfFunctionPlot[98].X = 0.98
	pointsOfFunctionPlot[98].Y = 0.998

	pointsOfFunctionPlot[99].X = 0.99
	pointsOfFunctionPlot[99].Y = 0.999

	pointsOfFunctionPlot[100].X = 1.0
	pointsOfFunctionPlot[100].Y = 1.0

	pointsOfFunctionPlot[101].X = 1.01
	pointsOfFunctionPlot[101].Y = 1.0

	pointsOfFunctionPlot[102].X = 1.02
	pointsOfFunctionPlot[102].Y = 1.001

	pointsOfFunctionPlot[103].X = 1.03
	pointsOfFunctionPlot[103].Y = 1.002

	pointsOfFunctionPlot[104].X = 1.04
	pointsOfFunctionPlot[104].Y = 1.004

	pointsOfFunctionPlot[105].X = 1.05
	pointsOfFunctionPlot[105].Y = 1.004

	pointsOfFunctionPlot[106].X = 1.06
	pointsOfFunctionPlot[106].Y = 1.005

	pointsOfFunctionPlot[107].X = 1.07
	pointsOfFunctionPlot[107].Y = 1.006

	pointsOfFunctionPlot[108].X = 1.08
	pointsOfFunctionPlot[108].Y = 1.007

	pointsOfFunctionPlot[109].X = 1.09
	pointsOfFunctionPlot[109].Y = 1.007

	pointsOfFunctionPlot[110].X = 1.10
	pointsOfFunctionPlot[110].Y = 1.007

	pointsOfFunctionPlot[111].X = 1.11
	pointsOfFunctionPlot[111].Y = 1.008

	pointsOfFunctionPlot[112].X = 1.12
	pointsOfFunctionPlot[112].Y = 1.009

	pointsOfFunctionPlot[113].X = 1.13
	pointsOfFunctionPlot[113].Y = 1.01

	pointsOfFunctionPlot[114].X = 1.14
	pointsOfFunctionPlot[114].Y = 1.01

	pointsOfFunctionPlot[115].X = 1.15
	pointsOfFunctionPlot[115].Y = 1.011

	pointsOfFunctionPlot[116].X = 1.16
	pointsOfFunctionPlot[116].Y = 1.012

	pointsOfFunctionPlot[117].X = 1.17
	pointsOfFunctionPlot[117].Y = 1.013

	pointsOfFunctionPlot[118].X = 1.18
	pointsOfFunctionPlot[118].Y = 1.013

	pointsOfFunctionPlot[119].X = 1.19
	pointsOfFunctionPlot[119].Y = 1.014

	pointsOfFunctionPlot[120].X = 1.20
	pointsOfFunctionPlot[120].Y = 1.015

	pointsOfFunctionPlot[121].X = 1.21
	pointsOfFunctionPlot[121].Y = 1.016

	pointsOfFunctionPlot[122].X = 1.22
	pointsOfFunctionPlot[122].Y = 1.016

	pointsOfFunctionPlot[123].X = 1.23
	pointsOfFunctionPlot[123].Y = 1.017

	pointsOfFunctionPlot[124].X = 1.24
	pointsOfFunctionPlot[124].Y = 1.018

	pointsOfFunctionPlot[125].X = 1.25
	pointsOfFunctionPlot[125].Y = 1.018

	pointsOfFunctionPlot[126].X = 1.26
	pointsOfFunctionPlot[126].Y = 1.019

	pointsOfFunctionPlot[127].X = 1.27
	pointsOfFunctionPlot[127].Y = 1.02

	pointsOfFunctionPlot[128].X = 1.28
	pointsOfFunctionPlot[128].Y = 1.02

	pointsOfFunctionPlot[129].X = 1.29
	pointsOfFunctionPlot[129].Y = 1.021

	pointsOfFunctionPlot[130].X = 1.30
	pointsOfFunctionPlot[130].Y = 1.022

	pointsOfFunctionPlot[131].X = 1.31
	pointsOfFunctionPlot[131].Y = 1.022

	pointsOfFunctionPlot[132].X = 1.32
	pointsOfFunctionPlot[132].Y = 1.023

	pointsOfFunctionPlot[133].X = 1.33
	pointsOfFunctionPlot[133].Y = 1.024

	pointsOfFunctionPlot[134].X = 1.34
	pointsOfFunctionPlot[134].Y = 1.024

	pointsOfFunctionPlot[135].X = 1.35
	pointsOfFunctionPlot[135].Y = 1.025

	pointsOfFunctionPlot[136].X = 1.36
	pointsOfFunctionPlot[136].Y = 1.025

	pointsOfFunctionPlot[137].X = 1.37
	pointsOfFunctionPlot[137].Y = 1.026

	pointsOfFunctionPlot[138].X = 1.38
	pointsOfFunctionPlot[138].Y = 1.027

	pointsOfFunctionPlot[139].X = 1.39
	pointsOfFunctionPlot[139].Y = 1.027

	pointsOfFunctionPlot[140].X = 1.40
	pointsOfFunctionPlot[140].Y = 1.028

	pointsOfFunctionPlot[141].X = 1.41
	pointsOfFunctionPlot[141].Y = 1.029

	pointsOfFunctionPlot[142].X = 1.42
	pointsOfFunctionPlot[142].Y = 1.029

	pointsOfFunctionPlot[143].X = 1.43
	pointsOfFunctionPlot[143].Y = 1.03

	pointsOfFunctionPlot[144].X = 1.44
	pointsOfFunctionPlot[144].Y = 1.03

	pointsOfFunctionPlot[145].X = 1.45
	pointsOfFunctionPlot[145].Y = 1.031

	pointsOfFunctionPlot[146].X = 1.46
	pointsOfFunctionPlot[146].Y = 1.032

	pointsOfFunctionPlot[147].X = 1.47
	pointsOfFunctionPlot[147].Y = 1.032

	pointsOfFunctionPlot[148].X = 1.48
	pointsOfFunctionPlot[148].Y = 1.033

	pointsOfFunctionPlot[149].X = 1.49
	pointsOfFunctionPlot[149].Y = 1.033

	pointsOfFunctionPlot[150].X = 1.50
	pointsOfFunctionPlot[150].Y = 1.034

	pointsOfFunctionPlot[151].X = 1.51
	pointsOfFunctionPlot[151].Y = 1.034

	pointsOfFunctionPlot[152].X = 1.52
	pointsOfFunctionPlot[152].Y = 1.035

	pointsOfFunctionPlot[153].X = 1.53
	pointsOfFunctionPlot[153].Y = 1.036

	pointsOfFunctionPlot[154].X = 1.54
	pointsOfFunctionPlot[154].Y = 1.036

	pointsOfFunctionPlot[155].X = 1.55
	pointsOfFunctionPlot[155].Y = 1.037

	pointsOfFunctionPlot[156].X = 1.56
	pointsOfFunctionPlot[156].Y = 1.037

	pointsOfFunctionPlot[157].X = 1.57
	pointsOfFunctionPlot[157].Y = 1.038

	pointsOfFunctionPlot[158].X = 1.58
	pointsOfFunctionPlot[158].Y = 1.038

	pointsOfFunctionPlot[159].X = 1.59
	pointsOfFunctionPlot[159].Y = 1.039

	pointsOfFunctionPlot[160].X = 1.60
	pointsOfFunctionPlot[160].Y = 1.039

	pointsOfFunctionPlot[161].X = 1.61
	pointsOfFunctionPlot[161].Y = 1.04

	pointsOfFunctionPlot[162].X = 1.62
	pointsOfFunctionPlot[162].Y = 1.041

	pointsOfFunctionPlot[163].X = 1.63
	pointsOfFunctionPlot[163].Y = 1.041

	pointsOfFunctionPlot[164].X = 1.64
	pointsOfFunctionPlot[164].Y = 1.042

	pointsOfFunctionPlot[165].X = 1.65
	pointsOfFunctionPlot[165].Y = 1.042

	pointsOfFunctionPlot[166].X = 1.66
	pointsOfFunctionPlot[166].Y = 1.043

	pointsOfFunctionPlot[167].X = 1.67
	pointsOfFunctionPlot[167].Y = 1.043

	pointsOfFunctionPlot[168].X = 1.68
	pointsOfFunctionPlot[168].Y = 1.044

	pointsOfFunctionPlot[169].X = 1.69
	pointsOfFunctionPlot[169].Y = 1.044

	pointsOfFunctionPlot[170].X = 1.70
	pointsOfFunctionPlot[170].Y = 1.045

	pointsOfFunctionPlot[171].X = 1.71
	pointsOfFunctionPlot[171].Y = 1.045

	pointsOfFunctionPlot[172].X = 1.72
	pointsOfFunctionPlot[172].Y = 1.046

	pointsOfFunctionPlot[173].X = 1.73
	pointsOfFunctionPlot[173].Y = 1.046

	pointsOfFunctionPlot[174].X = 1.74
	pointsOfFunctionPlot[174].Y = 1.047

	pointsOfFunctionPlot[175].X = 1.75
	pointsOfFunctionPlot[175].Y = 1.047

	pointsOfFunctionPlot[176].X = 1.76
	pointsOfFunctionPlot[176].Y = 1.048

	pointsOfFunctionPlot[177].X = 1.77
	pointsOfFunctionPlot[177].Y = 1.048

	pointsOfFunctionPlot[178].X = 1.78
	pointsOfFunctionPlot[178].Y = 1.049

	pointsOfFunctionPlot[179].X = 1.79
	pointsOfFunctionPlot[179].Y = 1.049

	pointsOfFunctionPlot[180].X = 1.80
	pointsOfFunctionPlot[180].Y = 1.05

	pointsOfFunctionPlot[181].X = 1.81
	pointsOfFunctionPlot[181].Y = 1.05

	pointsOfFunctionPlot[182].X = 1.82
	pointsOfFunctionPlot[182].Y = 1.051

	pointsOfFunctionPlot[183].X = 1.83
	pointsOfFunctionPlot[183].Y = 1.051

	pointsOfFunctionPlot[184].X = 1.84
	pointsOfFunctionPlot[184].Y = 1.052

	pointsOfFunctionPlot[185].X = 1.85
	pointsOfFunctionPlot[185].Y = 1.052

	pointsOfFunctionPlot[186].X = 1.86
	pointsOfFunctionPlot[186].Y = 1.053

	pointsOfFunctionPlot[187].X = 1.87
	pointsOfFunctionPlot[187].Y = 1.053

	pointsOfFunctionPlot[188].X = 1.88
	pointsOfFunctionPlot[188].Y = 1.054

	pointsOfFunctionPlot[189].X = 1.89
	pointsOfFunctionPlot[189].Y = 1.054

	pointsOfFunctionPlot[190].X = 1.90
	pointsOfFunctionPlot[190].Y = 1.054

	pointsOfFunctionPlot[191].X = 1.91
	pointsOfFunctionPlot[191].Y = 1.055

	pointsOfFunctionPlot[192].X = 1.92
	pointsOfFunctionPlot[192].Y = 1.055

	pointsOfFunctionPlot[193].X = 1.93
	pointsOfFunctionPlot[193].Y = 1.056

	pointsOfFunctionPlot[194].X = 1.94
	pointsOfFunctionPlot[194].Y = 1.056

	pointsOfFunctionPlot[195].X = 1.95
	pointsOfFunctionPlot[195].Y = 1.057

	pointsOfFunctionPlot[196].X = 1.96
	pointsOfFunctionPlot[196].Y = 1.057

	pointsOfFunctionPlot[197].X = 1.97
	pointsOfFunctionPlot[197].Y = 1.058

	pointsOfFunctionPlot[198].X = 1.98
	pointsOfFunctionPlot[198].Y = 1.058

	pointsOfFunctionPlot[199].X = 1.99
	pointsOfFunctionPlot[199].Y = 1.059

	pointsOfFunctionPlot[200].X = 2.0
	pointsOfFunctionPlot[200].Y = 1.059

	pointsOfFunctionPlot[201].X = 2.01
	pointsOfFunctionPlot[201].Y = 1.059

	pointsOfFunctionPlot[202].X = 2.02
	pointsOfFunctionPlot[202].Y = 1.06

	pointsOfFunctionPlot[203].X = 2.03
	pointsOfFunctionPlot[203].Y = 1.06

	pointsOfFunctionPlot[204].X = 2.04
	pointsOfFunctionPlot[204].Y = 1.061

	pointsOfFunctionPlot[205].X = 2.05
	pointsOfFunctionPlot[205].Y = 1.061

	pointsOfFunctionPlot[206].X = 2.06
	pointsOfFunctionPlot[206].Y = 1.062

	pointsOfFunctionPlot[207].X = 2.07
	pointsOfFunctionPlot[207].Y = 1.062

	pointsOfFunctionPlot[208].X = 2.08
	pointsOfFunctionPlot[208].Y = 1.062

	pointsOfFunctionPlot[209].X = 2.09
	pointsOfFunctionPlot[209].Y = 1.062

	pointsOfFunctionPlot[210].X = 2.10
	pointsOfFunctionPlot[210].Y = 1.063

	pointsOfFunctionPlot[211].X = 2.11
	pointsOfFunctionPlot[211].Y = 1.063

	pointsOfFunctionPlot[212].X = 2.12
	pointsOfFunctionPlot[212].Y = 1.064

	pointsOfFunctionPlot[213].X = 2.13
	pointsOfFunctionPlot[213].Y = 1.064

	pointsOfFunctionPlot[214].X = 2.14
	pointsOfFunctionPlot[214].Y = 1.065

	pointsOfFunctionPlot[215].X = 2.15
	pointsOfFunctionPlot[215].Y = 1.065

	pointsOfFunctionPlot[216].X = 2.16
	pointsOfFunctionPlot[216].Y = 1.066

	pointsOfFunctionPlot[217].X = 2.17
	pointsOfFunctionPlot[217].Y = 1.066

	pointsOfFunctionPlot[218].X = 2.18
	pointsOfFunctionPlot[218].Y = 1.067

	pointsOfFunctionPlot[219].X = 2.19
	pointsOfFunctionPlot[219].Y = 1.067

	pointsOfFunctionPlot[220].X = 2.20
	pointsOfFunctionPlot[220].Y = 1.067

	pointsOfFunctionPlot[221].X = 2.21
	pointsOfFunctionPlot[221].Y = 1.068

	pointsOfFunctionPlot[222].X = 2.22
	pointsOfFunctionPlot[222].Y = 1.068

	pointsOfFunctionPlot[223].X = 2.23
	pointsOfFunctionPlot[223].Y = 1.069

	pointsOfFunctionPlot[224].X = 2.24
	pointsOfFunctionPlot[224].Y = 1.069

	pointsOfFunctionPlot[225].X = 2.25
	pointsOfFunctionPlot[225].Y = 1.069

	pointsOfFunctionPlot[226].X = 2.26
	pointsOfFunctionPlot[226].Y = 1.07

	pointsOfFunctionPlot[227].X = 2.27
	pointsOfFunctionPlot[227].Y = 1.07

	pointsOfFunctionPlot[228].X = 2.28
	pointsOfFunctionPlot[228].Y = 1.071

	pointsOfFunctionPlot[229].X = 2.29
	pointsOfFunctionPlot[229].Y = 1.071

	pointsOfFunctionPlot[230].X = 2.30
	pointsOfFunctionPlot[230].Y = 1.071

	pointsOfFunctionPlot[231].X = 2.31
	pointsOfFunctionPlot[231].Y = 1.072

	pointsOfFunctionPlot[232].X = 2.32
	pointsOfFunctionPlot[232].Y = 1.072

	pointsOfFunctionPlot[233].X = 2.33
	pointsOfFunctionPlot[233].Y = 1.073

	pointsOfFunctionPlot[234].X = 2.34
	pointsOfFunctionPlot[234].Y = 1.073

	pointsOfFunctionPlot[235].X = 2.35
	pointsOfFunctionPlot[235].Y = 1.073

	pointsOfFunctionPlot[236].X = 2.36
	pointsOfFunctionPlot[236].Y = 1.074

	pointsOfFunctionPlot[237].X = 2.37
	pointsOfFunctionPlot[237].Y = 1.074

	pointsOfFunctionPlot[238].X = 2.38
	pointsOfFunctionPlot[238].Y = 1.074

	pointsOfFunctionPlot[239].X = 2.39
	pointsOfFunctionPlot[239].Y = 1.075

	pointsOfFunctionPlot[240].X = 2.40
	pointsOfFunctionPlot[240].Y = 1.075

	pointsOfFunctionPlot[241].X = 2.41
	pointsOfFunctionPlot[241].Y = 1.076

	pointsOfFunctionPlot[242].X = 2.42
	pointsOfFunctionPlot[242].Y = 1.076

	pointsOfFunctionPlot[243].X = 2.43
	pointsOfFunctionPlot[243].Y = 1.076

	pointsOfFunctionPlot[244].X = 2.44
	pointsOfFunctionPlot[244].Y = 1.077

	pointsOfFunctionPlot[245].X = 2.45
	pointsOfFunctionPlot[245].Y = 1.077

	pointsOfFunctionPlot[246].X = 2.46
	pointsOfFunctionPlot[246].Y = 1.077

	pointsOfFunctionPlot[247].X = 2.47
	pointsOfFunctionPlot[247].Y = 1.078

	pointsOfFunctionPlot[248].X = 2.48
	pointsOfFunctionPlot[248].Y = 1.078

	pointsOfFunctionPlot[249].X = 2.49
	pointsOfFunctionPlot[249].Y = 1.078

	pointsOfFunctionPlot[250].X = 2.50
	pointsOfFunctionPlot[250].Y = 1.079

	pointsOfFunctionPlot[251].X = 2.51
	pointsOfFunctionPlot[251].Y = 1.079

	pointsOfFunctionPlot[252].X = 2.52
	pointsOfFunctionPlot[252].Y = 1.08

	pointsOfFunctionPlot[253].X = 2.53
	pointsOfFunctionPlot[253].Y = 1.08

	pointsOfFunctionPlot[254].X = 2.54
	pointsOfFunctionPlot[254].Y = 1.08

	pointsOfFunctionPlot[255].X = 2.55
	pointsOfFunctionPlot[255].Y = 1.081

	pointsOfFunctionPlot[256].X = 2.56
	pointsOfFunctionPlot[256].Y = 1.081

	pointsOfFunctionPlot[257].X = 2.57
	pointsOfFunctionPlot[257].Y = 1.081

	pointsOfFunctionPlot[258].X = 2.58
	pointsOfFunctionPlot[258].Y = 1.082

	pointsOfFunctionPlot[259].X = 2.59
	pointsOfFunctionPlot[259].Y = 1.082

	pointsOfFunctionPlot[260].X = 2.60
	pointsOfFunctionPlot[260].Y = 1.082

	pointsOfFunctionPlot[261].X = 2.61
	pointsOfFunctionPlot[261].Y = 1.083

	pointsOfFunctionPlot[262].X = 2.62
	pointsOfFunctionPlot[262].Y = 1.083

	pointsOfFunctionPlot[263].X = 2.63
	pointsOfFunctionPlot[263].Y = 1.083

	pointsOfFunctionPlot[264].X = 2.64
	pointsOfFunctionPlot[264].Y = 1.084

	pointsOfFunctionPlot[265].X = 2.65
	pointsOfFunctionPlot[265].Y = 1.084

	pointsOfFunctionPlot[266].X = 2.66
	pointsOfFunctionPlot[266].Y = 1.084

	pointsOfFunctionPlot[267].X = 2.67
	pointsOfFunctionPlot[267].Y = 1.085

	pointsOfFunctionPlot[268].X = 2.68
	pointsOfFunctionPlot[268].Y = 1.085

	pointsOfFunctionPlot[269].X = 2.69
	pointsOfFunctionPlot[269].Y = 1.085

	pointsOfFunctionPlot[270].X = 2.70
	pointsOfFunctionPlot[270].Y = 1.086

	pointsOfFunctionPlot[271].X = 2.71
	pointsOfFunctionPlot[271].Y = 1.086

	pointsOfFunctionPlot[272].X = 2.72
	pointsOfFunctionPlot[272].Y = 1.086

	pointsOfFunctionPlot[273].X = 2.73
	pointsOfFunctionPlot[273].Y = 1.087

	pointsOfFunctionPlot[274].X = 2.74
	pointsOfFunctionPlot[274].Y = 1.087

	pointsOfFunctionPlot[275].X = 2.75
	pointsOfFunctionPlot[275].Y = 1.087

	pointsOfFunctionPlot[276].X = 2.76
	pointsOfFunctionPlot[276].Y = 1.088

	pointsOfFunctionPlot[277].X = 2.77
	pointsOfFunctionPlot[277].Y = 1.088

	pointsOfFunctionPlot[278].X = 2.78
	pointsOfFunctionPlot[278].Y = 1.088

	pointsOfFunctionPlot[279].X = 2.79
	pointsOfFunctionPlot[279].Y = 1.089

	pointsOfFunctionPlot[280].X = 2.80
	pointsOfFunctionPlot[280].Y = 1.089

	pointsOfFunctionPlot[281].X = 2.81
	pointsOfFunctionPlot[281].Y = 1.089

	pointsOfFunctionPlot[282].X = 2.82
	pointsOfFunctionPlot[282].Y = 1.09

	pointsOfFunctionPlot[283].X = 2.83
	pointsOfFunctionPlot[283].Y = 1.09

	pointsOfFunctionPlot[284].X = 2.84
	pointsOfFunctionPlot[284].Y = 1.09

	pointsOfFunctionPlot[285].X = 2.85
	pointsOfFunctionPlot[285].Y = 1.091

	pointsOfFunctionPlot[286].X = 2.86
	pointsOfFunctionPlot[286].Y = 1.091

	pointsOfFunctionPlot[287].X = 2.87
	pointsOfFunctionPlot[287].Y = 1.091

	pointsOfFunctionPlot[288].X = 2.88
	pointsOfFunctionPlot[288].Y = 1.092

	pointsOfFunctionPlot[289].X = 2.89
	pointsOfFunctionPlot[289].Y = 1.092

	pointsOfFunctionPlot[290].X = 2.90
	pointsOfFunctionPlot[290].Y = 1.092

	pointsOfFunctionPlot[291].X = 2.91
	pointsOfFunctionPlot[291].Y = 1.093

	pointsOfFunctionPlot[292].X = 2.92
	pointsOfFunctionPlot[292].Y = 1.093

	pointsOfFunctionPlot[293].X = 2.93
	pointsOfFunctionPlot[293].Y = 1.093

	pointsOfFunctionPlot[294].X = 2.94
	pointsOfFunctionPlot[294].Y = 1.094

	pointsOfFunctionPlot[295].X = 2.95
	pointsOfFunctionPlot[295].Y = 1.094

	pointsOfFunctionPlot[296].X = 2.96
	pointsOfFunctionPlot[296].Y = 1.094

	pointsOfFunctionPlot[297].X = 2.97
	pointsOfFunctionPlot[297].Y = 1.094

	pointsOfFunctionPlot[298].X = 2.98
	pointsOfFunctionPlot[298].Y = 1.095

	pointsOfFunctionPlot[299].X = 2.99
	pointsOfFunctionPlot[299].Y = 1.095

	pointsOfFunctionPlot[300].X = 3.0
	pointsOfFunctionPlot[300].Y = 1.095

	pointsOfFunctionPlot[301].X = 3.01
	pointsOfFunctionPlot[301].Y = 1.096

	pointsOfFunctionPlot[302].X = 3.02
	pointsOfFunctionPlot[302].Y = 1.096

	pointsOfFunctionPlot[303].X = 3.03
	pointsOfFunctionPlot[303].Y = 1.096

	pointsOfFunctionPlot[304].X = 3.04
	pointsOfFunctionPlot[304].Y = 1.097

	pointsOfFunctionPlot[305].X = 3.05
	pointsOfFunctionPlot[305].Y = 1.097

	pointsOfFunctionPlot[306].X = 3.06
	pointsOfFunctionPlot[306].Y = 1.097

	pointsOfFunctionPlot[307].X = 3.07
	pointsOfFunctionPlot[307].Y = 1.097

	pointsOfFunctionPlot[308].X = 3.08
	pointsOfFunctionPlot[308].Y = 1.098

	pointsOfFunctionPlot[309].X = 3.09
	pointsOfFunctionPlot[309].Y = 1.098

	pointsOfFunctionPlot[310].X = 3.10
	pointsOfFunctionPlot[310].Y = 1.098

	pointsOfFunctionPlot[311].X = 3.11
	pointsOfFunctionPlot[311].Y = 1.099

	pointsOfFunctionPlot[312].X = 3.12
	pointsOfFunctionPlot[312].Y = 1.099

	pointsOfFunctionPlot[313].X = 3.13
	pointsOfFunctionPlot[313].Y = 1.099

	pointsOfFunctionPlot[314].X = 3.14
	pointsOfFunctionPlot[314].Y = 1.1

	pointsOfFunctionPlot[315].X = 3.15
	pointsOfFunctionPlot[315].Y = 1.1

	pointsOfFunctionPlot[316].X = 3.16
	pointsOfFunctionPlot[316].Y = 1.1

	pointsOfFunctionPlot[317].X = 3.17
	pointsOfFunctionPlot[317].Y = 1.1

	pointsOfFunctionPlot[318].X = 3.18
	pointsOfFunctionPlot[318].Y = 1.101

	pointsOfFunctionPlot[319].X = 3.19
	pointsOfFunctionPlot[319].Y = 1.101

	pointsOfFunctionPlot[320].X = 3.20
	pointsOfFunctionPlot[320].Y = 1.101

	pointsOfFunctionPlot[321].X = 3.21
	pointsOfFunctionPlot[321].Y = 1.102

	pointsOfFunctionPlot[322].X = 3.22
	pointsOfFunctionPlot[322].Y = 1.102

	pointsOfFunctionPlot[323].X = 3.23
	pointsOfFunctionPlot[323].Y = 1.102

	pointsOfFunctionPlot[324].X = 3.24
	pointsOfFunctionPlot[324].Y = 1.102

	pointsOfFunctionPlot[325].X = 3.25
	pointsOfFunctionPlot[325].Y = 1.103

	pointsOfFunctionPlot[326].X = 3.26
	pointsOfFunctionPlot[326].Y = 1.103

	pointsOfFunctionPlot[327].X = 3.27
	pointsOfFunctionPlot[327].Y = 1.103

	pointsOfFunctionPlot[328].X = 3.28
	pointsOfFunctionPlot[328].Y = 1.104

	pointsOfFunctionPlot[329].X = 3.29
	pointsOfFunctionPlot[329].Y = 1.104

	pointsOfFunctionPlot[330].X = 3.30
	pointsOfFunctionPlot[330].Y = 1.104

	pointsOfFunctionPlot[331].X = 3.31
	pointsOfFunctionPlot[331].Y = 1.104

	pointsOfFunctionPlot[332].X = 3.32
	pointsOfFunctionPlot[332].Y = 1.105

	pointsOfFunctionPlot[333].X = 3.33
	pointsOfFunctionPlot[333].Y = 1.105

	pointsOfFunctionPlot[334].X = 3.34
	pointsOfFunctionPlot[334].Y = 1.105

	pointsOfFunctionPlot[335].X = 3.35
	pointsOfFunctionPlot[335].Y = 1.105

	pointsOfFunctionPlot[336].X = 3.36
	pointsOfFunctionPlot[336].Y = 1.106

	pointsOfFunctionPlot[337].X = 3.37
	pointsOfFunctionPlot[337].Y = 1.106

	pointsOfFunctionPlot[338].X = 3.38
	pointsOfFunctionPlot[338].Y = 1.106

	pointsOfFunctionPlot[339].X = 3.39
	pointsOfFunctionPlot[339].Y = 1.107

	pointsOfFunctionPlot[340].X = 3.40
	pointsOfFunctionPlot[340].Y = 1.107

	pointsOfFunctionPlot[341].X = 3.41
	pointsOfFunctionPlot[341].Y = 1.107

	pointsOfFunctionPlot[342].X = 3.42
	pointsOfFunctionPlot[342].Y = 1.107

	pointsOfFunctionPlot[343].X = 3.43
	pointsOfFunctionPlot[343].Y = 1.108

	pointsOfFunctionPlot[344].X = 3.44
	pointsOfFunctionPlot[344].Y = 1.108

	pointsOfFunctionPlot[345].X = 3.45
	pointsOfFunctionPlot[345].Y = 1.108

	pointsOfFunctionPlot[346].X = 3.46
	pointsOfFunctionPlot[346].Y = 1.108

	pointsOfFunctionPlot[347].X = 3.47
	pointsOfFunctionPlot[347].Y = 1.109

	pointsOfFunctionPlot[348].X = 3.48
	pointsOfFunctionPlot[348].Y = 1.109

	pointsOfFunctionPlot[349].X = 3.49
	pointsOfFunctionPlot[349].Y = 1.109

	pointsOfFunctionPlot[350].X = 3.50
	pointsOfFunctionPlot[350].Y = 1.11

	pointsOfFunctionPlot[351].X = 3.51
	pointsOfFunctionPlot[351].Y = 1.11

	pointsOfFunctionPlot[352].X = 3.52
	pointsOfFunctionPlot[352].Y = 1.11

	pointsOfFunctionPlot[353].X = 3.53
	pointsOfFunctionPlot[353].Y = 1.11

	pointsOfFunctionPlot[354].X = 3.54
	pointsOfFunctionPlot[354].Y = 1.111

	pointsOfFunctionPlot[355].X = 3.55
	pointsOfFunctionPlot[355].Y = 1.111

	pointsOfFunctionPlot[356].X = 3.56
	pointsOfFunctionPlot[356].Y = 1.111

	pointsOfFunctionPlot[357].X = 3.57
	pointsOfFunctionPlot[357].Y = 1.111

	pointsOfFunctionPlot[358].X = 3.58
	pointsOfFunctionPlot[358].Y = 1.112

	pointsOfFunctionPlot[359].X = 3.59
	pointsOfFunctionPlot[359].Y = 1.112

	pointsOfFunctionPlot[360].X = 3.60
	pointsOfFunctionPlot[360].Y = 1.112

	pointsOfFunctionPlot[361].X = 3.61
	pointsOfFunctionPlot[361].Y = 1.112

	pointsOfFunctionPlot[362].X = 3.62
	pointsOfFunctionPlot[362].Y = 1.113

	pointsOfFunctionPlot[363].X = 3.63
	pointsOfFunctionPlot[363].Y = 1.113

	pointsOfFunctionPlot[364].X = 3.64
	pointsOfFunctionPlot[364].Y = 1.113

	pointsOfFunctionPlot[365].X = 3.65
	pointsOfFunctionPlot[365].Y = 1.113

	pointsOfFunctionPlot[366].X = 3.66
	pointsOfFunctionPlot[366].Y = 1.114

	pointsOfFunctionPlot[367].X = 3.67
	pointsOfFunctionPlot[367].Y = 1.114

	pointsOfFunctionPlot[368].X = 3.68
	pointsOfFunctionPlot[368].Y = 1.114

	pointsOfFunctionPlot[369].X = 3.69
	pointsOfFunctionPlot[369].Y = 1.114

	pointsOfFunctionPlot[370].X = 3.70
	pointsOfFunctionPlot[370].Y = 1.115

	pointsOfFunctionPlot[371].X = 3.71
	pointsOfFunctionPlot[371].Y = 1.115

	pointsOfFunctionPlot[372].X = 3.72
	pointsOfFunctionPlot[372].Y = 1.115

	pointsOfFunctionPlot[373].X = 3.73
	pointsOfFunctionPlot[373].Y = 1.115

	pointsOfFunctionPlot[374].X = 3.74
	pointsOfFunctionPlot[374].Y = 1.116

	pointsOfFunctionPlot[375].X = 3.75
	pointsOfFunctionPlot[375].Y = 1.116

	pointsOfFunctionPlot[376].X = 3.76
	pointsOfFunctionPlot[376].Y = 1.116

	pointsOfFunctionPlot[377].X = 3.77
	pointsOfFunctionPlot[377].Y = 1.116

	pointsOfFunctionPlot[378].X = 3.78
	pointsOfFunctionPlot[378].Y = 1.117

	pointsOfFunctionPlot[379].X = 3.79
	pointsOfFunctionPlot[379].Y = 1.117

	pointsOfFunctionPlot[380].X = 3.80
	pointsOfFunctionPlot[380].Y = 1.117

	pointsOfFunctionPlot[381].X = 3.81
	pointsOfFunctionPlot[381].Y = 1.117

	pointsOfFunctionPlot[382].X = 3.82
	pointsOfFunctionPlot[382].Y = 1.118

	pointsOfFunctionPlot[383].X = 3.83
	pointsOfFunctionPlot[383].Y = 1.118

	pointsOfFunctionPlot[384].X = 3.84
	pointsOfFunctionPlot[384].Y = 1.118

	pointsOfFunctionPlot[385].X = 3.85
	pointsOfFunctionPlot[385].Y = 1.118

	pointsOfFunctionPlot[386].X = 3.86
	pointsOfFunctionPlot[386].Y = 1.119

	pointsOfFunctionPlot[387].X = 3.87
	pointsOfFunctionPlot[387].Y = 1.119

	pointsOfFunctionPlot[388].X = 3.88
	pointsOfFunctionPlot[388].Y = 1.119

	pointsOfFunctionPlot[389].X = 3.89
	pointsOfFunctionPlot[389].Y = 1.119

	pointsOfFunctionPlot[390].X = 3.90
	pointsOfFunctionPlot[390].Y = 1.12

	pointsOfFunctionPlot[391].X = 3.91
	pointsOfFunctionPlot[391].Y = 1.12

	pointsOfFunctionPlot[392].X = 3.92
	pointsOfFunctionPlot[392].Y = 1.12

	pointsOfFunctionPlot[393].X = 3.93
	pointsOfFunctionPlot[393].Y = 1.12

	pointsOfFunctionPlot[394].X = 3.94
	pointsOfFunctionPlot[394].Y = 1.121

	pointsOfFunctionPlot[395].X = 3.95
	pointsOfFunctionPlot[395].Y = 1.121

	pointsOfFunctionPlot[396].X = 3.96
	pointsOfFunctionPlot[396].Y = 1.121

	pointsOfFunctionPlot[397].X = 3.97
	pointsOfFunctionPlot[397].Y = 1.121

	pointsOfFunctionPlot[398].X = 3.98
	pointsOfFunctionPlot[398].Y = 1.121

	pointsOfFunctionPlot[399].X = 3.99
	pointsOfFunctionPlot[399].Y = 1.122

	pointsOfFunctionPlot[400].X = 4.0
	pointsOfFunctionPlot[400].Y = 1.122

	pointsOfFunctionPlot[401].X = 4.01
	pointsOfFunctionPlot[401].Y = 1.122

	pointsOfFunctionPlot[402].X = 4.02
	pointsOfFunctionPlot[402].Y = 1.122

	pointsOfFunctionPlot[403].X = 4.03
	pointsOfFunctionPlot[403].Y = 1.123

	pointsOfFunctionPlot[404].X = 4.04
	pointsOfFunctionPlot[404].Y = 1.123

	pointsOfFunctionPlot[405].X = 4.05
	pointsOfFunctionPlot[405].Y = 1.123

	pointsOfFunctionPlot[406].X = 4.06
	pointsOfFunctionPlot[406].Y = 1.123

	pointsOfFunctionPlot[407].X = 4.07
	pointsOfFunctionPlot[407].Y = 1.124

	pointsOfFunctionPlot[408].X = 4.08
	pointsOfFunctionPlot[408].Y = 1.124

	pointsOfFunctionPlot[409].X = 4.09
	pointsOfFunctionPlot[409].Y = 1.124

	pointsOfFunctionPlot[410].X = 4.10
	pointsOfFunctionPlot[410].Y = 1.124

	pointsOfFunctionPlot[411].X = 4.11
	pointsOfFunctionPlot[411].Y = 1.125

	pointsOfFunctionPlot[412].X = 4.12
	pointsOfFunctionPlot[412].Y = 1.125

	pointsOfFunctionPlot[413].X = 4.13
	pointsOfFunctionPlot[413].Y = 1.125

	pointsOfFunctionPlot[414].X = 4.14
	pointsOfFunctionPlot[414].Y = 1.125

	pointsOfFunctionPlot[415].X = 4.15
	pointsOfFunctionPlot[415].Y = 1.125

	pointsOfFunctionPlot[416].X = 4.16
	pointsOfFunctionPlot[416].Y = 1.126

	pointsOfFunctionPlot[417].X = 4.17
	pointsOfFunctionPlot[417].Y = 1.126

	pointsOfFunctionPlot[418].X = 4.18
	pointsOfFunctionPlot[418].Y = 1.126

	pointsOfFunctionPlot[419].X = 4.19
	pointsOfFunctionPlot[419].Y = 1.126

	pointsOfFunctionPlot[420].X = 4.20
	pointsOfFunctionPlot[420].Y = 1.127

	pointsOfFunctionPlot[421].X = 4.21
	pointsOfFunctionPlot[421].Y = 1.127

	pointsOfFunctionPlot[422].X = 4.22
	pointsOfFunctionPlot[422].Y = 1.127

	pointsOfFunctionPlot[423].X = 4.23
	pointsOfFunctionPlot[423].Y = 1.127

	pointsOfFunctionPlot[424].X = 4.24
	pointsOfFunctionPlot[424].Y = 1.127

	pointsOfFunctionPlot[425].X = 4.25
	pointsOfFunctionPlot[425].Y = 1.128

	pointsOfFunctionPlot[426].X = 4.26
	pointsOfFunctionPlot[426].Y = 1.128

	pointsOfFunctionPlot[427].X = 4.27
	pointsOfFunctionPlot[427].Y = 1.128

	pointsOfFunctionPlot[428].X = 4.28
	pointsOfFunctionPlot[428].Y = 1.128

	pointsOfFunctionPlot[429].X = 4.29
	pointsOfFunctionPlot[429].Y = 1.129

	pointsOfFunctionPlot[430].X = 4.30
	pointsOfFunctionPlot[430].Y = 1.129

	pointsOfFunctionPlot[431].X = 4.31
	pointsOfFunctionPlot[431].Y = 1.129

	pointsOfFunctionPlot[432].X = 4.32
	pointsOfFunctionPlot[432].Y = 1.129

	pointsOfFunctionPlot[433].X = 4.33
	pointsOfFunctionPlot[433].Y = 1.129

	pointsOfFunctionPlot[434].X = 4.34
	pointsOfFunctionPlot[434].Y = 1.13

	pointsOfFunctionPlot[435].X = 4.35
	pointsOfFunctionPlot[435].Y = 1.13

	pointsOfFunctionPlot[436].X = 4.36
	pointsOfFunctionPlot[436].Y = 1.13

	pointsOfFunctionPlot[437].X = 4.37
	pointsOfFunctionPlot[437].Y = 1.13

	pointsOfFunctionPlot[438].X = 4.38
	pointsOfFunctionPlot[438].Y = 1.13

	pointsOfFunctionPlot[439].X = 4.39
	pointsOfFunctionPlot[439].Y = 1.131

	pointsOfFunctionPlot[440].X = 4.40
	pointsOfFunctionPlot[440].Y = 1.131

	pointsOfFunctionPlot[441].X = 4.41
	pointsOfFunctionPlot[441].Y = 1.131

	pointsOfFunctionPlot[442].X = 4.42
	pointsOfFunctionPlot[442].Y = 1.131

	pointsOfFunctionPlot[443].X = 4.43
	pointsOfFunctionPlot[443].Y = 1.13

	pointsOfFunctionPlot[444].X = 4.44
	pointsOfFunctionPlot[444].Y = 1.132

	pointsOfFunctionPlot[445].X = 4.45
	pointsOfFunctionPlot[445].Y = 1.132

	pointsOfFunctionPlot[446].X = 4.46
	pointsOfFunctionPlot[446].Y = 1.132

	pointsOfFunctionPlot[447].X = 4.47
	pointsOfFunctionPlot[447].Y = 1.132

	pointsOfFunctionPlot[448].X = 4.48
	pointsOfFunctionPlot[448].Y = 1.133

	pointsOfFunctionPlot[449].X = 4.49
	pointsOfFunctionPlot[449].Y = 1.133

	pointsOfFunctionPlot[450].X = 4.50
	pointsOfFunctionPlot[450].Y = 1.133

	pointsOfFunctionPlot[451].X = 4.51
	pointsOfFunctionPlot[451].Y = 1.133

	pointsOfFunctionPlot[452].X = 4.52
	pointsOfFunctionPlot[452].Y = 1.133

	pointsOfFunctionPlot[453].X = 4.53
	pointsOfFunctionPlot[453].Y = 1.134

	pointsOfFunctionPlot[454].X = 4.54
	pointsOfFunctionPlot[454].Y = 1.134

	pointsOfFunctionPlot[455].X = 4.55
	pointsOfFunctionPlot[455].Y = 1.134

	pointsOfFunctionPlot[456].X = 4.56
	pointsOfFunctionPlot[456].Y = 1.134

	pointsOfFunctionPlot[457].X = 4.57
	pointsOfFunctionPlot[457].Y = 1.134

	pointsOfFunctionPlot[458].X = 4.58
	pointsOfFunctionPlot[458].Y = 1.135

	pointsOfFunctionPlot[459].X = 4.59
	pointsOfFunctionPlot[459].Y = 1.135

	pointsOfFunctionPlot[460].X = 4.60
	pointsOfFunctionPlot[460].Y = 1.135

	pointsOfFunctionPlot[461].X = 4.61
	pointsOfFunctionPlot[461].Y = 1.135

	pointsOfFunctionPlot[462].X = 4.62
	pointsOfFunctionPlot[462].Y = 1.136

	pointsOfFunctionPlot[463].X = 4.63
	pointsOfFunctionPlot[463].Y = 1.136

	pointsOfFunctionPlot[464].X = 4.64
	pointsOfFunctionPlot[464].Y = 1.136

	pointsOfFunctionPlot[465].X = 4.65
	pointsOfFunctionPlot[465].Y = 1.136

	pointsOfFunctionPlot[466].X = 4.66
	pointsOfFunctionPlot[466].Y = 1.136

	pointsOfFunctionPlot[467].X = 4.67
	pointsOfFunctionPlot[467].Y = 1.137

	pointsOfFunctionPlot[468].X = 4.68
	pointsOfFunctionPlot[468].Y = 1.137

	pointsOfFunctionPlot[469].X = 4.69
	pointsOfFunctionPlot[469].Y = 1.137

	pointsOfFunctionPlot[470].X = 4.70
	pointsOfFunctionPlot[470].Y = 1.137

	pointsOfFunctionPlot[471].X = 4.71
	pointsOfFunctionPlot[471].Y = 1.137

	pointsOfFunctionPlot[472].X = 4.72
	pointsOfFunctionPlot[472].Y = 1.138

	pointsOfFunctionPlot[473].X = 4.73
	pointsOfFunctionPlot[473].Y = 1.138

	pointsOfFunctionPlot[474].X = 4.74
	pointsOfFunctionPlot[474].Y = 1.138

	pointsOfFunctionPlot[475].X = 4.75
	pointsOfFunctionPlot[475].Y = 1.138

	pointsOfFunctionPlot[476].X = 4.76
	pointsOfFunctionPlot[476].Y = 1.138

	pointsOfFunctionPlot[477].X = 4.77
	pointsOfFunctionPlot[477].Y = 1.139

	pointsOfFunctionPlot[478].X = 4.78
	pointsOfFunctionPlot[478].Y = 1.139

	pointsOfFunctionPlot[479].X = 4.79
	pointsOfFunctionPlot[479].Y = 1.139

	pointsOfFunctionPlot[480].X = 4.80
	pointsOfFunctionPlot[480].Y = 1.139

	pointsOfFunctionPlot[481].X = 4.81
	pointsOfFunctionPlot[481].Y = 1.139

	pointsOfFunctionPlot[482].X = 4.82
	pointsOfFunctionPlot[482].Y = 1.14

	pointsOfFunctionPlot[483].X = 4.83
	pointsOfFunctionPlot[483].Y = 1.14

	pointsOfFunctionPlot[484].X = 4.84
	pointsOfFunctionPlot[484].Y = 1.14

	pointsOfFunctionPlot[485].X = 4.85
	pointsOfFunctionPlot[485].Y = 1.14

	pointsOfFunctionPlot[486].X = 4.86
	pointsOfFunctionPlot[486].Y = 1.14

	pointsOfFunctionPlot[487].X = 4.87
	pointsOfFunctionPlot[487].Y = 1.141

	pointsOfFunctionPlot[488].X = 4.88
	pointsOfFunctionPlot[488].Y = 1.141

	pointsOfFunctionPlot[489].X = 4.89
	pointsOfFunctionPlot[489].Y = 1.141

	pointsOfFunctionPlot[490].X = 4.90
	pointsOfFunctionPlot[490].Y = 1.141

	pointsOfFunctionPlot[491].X = 4.91
	pointsOfFunctionPlot[491].Y = 1.141

	pointsOfFunctionPlot[492].X = 4.92
	pointsOfFunctionPlot[492].Y = 1.141

	pointsOfFunctionPlot[493].X = 4.93
	pointsOfFunctionPlot[493].Y = 1.142

	pointsOfFunctionPlot[494].X = 4.94
	pointsOfFunctionPlot[494].Y = 1.142

	pointsOfFunctionPlot[495].X = 4.95
	pointsOfFunctionPlot[495].Y = 1.142

	pointsOfFunctionPlot[496].X = 4.96
	pointsOfFunctionPlot[496].Y = 1.142

	pointsOfFunctionPlot[497].X = 4.97
	pointsOfFunctionPlot[497].Y = 1.142

	pointsOfFunctionPlot[498].X = 4.98
	pointsOfFunctionPlot[498].Y = 1.143

	pointsOfFunctionPlot[499].X = 4.99
	pointsOfFunctionPlot[499].Y = 1.143

	pointsOfFunctionPlot[500].X = 5.0
	pointsOfFunctionPlot[500].Y = 1.143

	pointsOfFunctionPlot[501].X = 5.01
	pointsOfFunctionPlot[501].Y = 1.143

	pointsOfFunctionPlot[502].X = 5.02
	pointsOfFunctionPlot[502].Y = 1.143

	pointsOfFunctionPlot[503].X = 5.03
	pointsOfFunctionPlot[503].Y = 1.144

	pointsOfFunctionPlot[504].X = 5.04
	pointsOfFunctionPlot[504].Y = 1.144

	pointsOfFunctionPlot[505].X = 5.05
	pointsOfFunctionPlot[505].Y = 1.144

	pointsOfFunctionPlot[506].X = 5.06
	pointsOfFunctionPlot[506].Y = 1.144

	pointsOfFunctionPlot[507].X = 5.07
	pointsOfFunctionPlot[507].Y = 1.144

	pointsOfFunctionPlot[508].X = 5.08
	pointsOfFunctionPlot[508].Y = 1.145

	pointsOfFunctionPlot[509].X = 5.09
	pointsOfFunctionPlot[509].Y = 1.145

	pointsOfFunctionPlot[510].X = 5.10
	pointsOfFunctionPlot[510].Y = 1.145

	pointsOfFunctionPlot[511].X = 5.11
	pointsOfFunctionPlot[511].Y = 1.145

	pointsOfFunctionPlot[512].X = 5.12
	pointsOfFunctionPlot[512].Y = 1.145

	pointsOfFunctionPlot[513].X = 5.13
	pointsOfFunctionPlot[513].Y = 1.145

	pointsOfFunctionPlot[514].X = 5.14
	pointsOfFunctionPlot[514].Y = 1.146

	pointsOfFunctionPlot[515].X = 5.15
	pointsOfFunctionPlot[515].Y = 1.146

	pointsOfFunctionPlot[516].X = 5.16
	pointsOfFunctionPlot[516].Y = 1.146

	pointsOfFunctionPlot[517].X = 5.17
	pointsOfFunctionPlot[517].Y = 1.146

	pointsOfFunctionPlot[518].X = 5.18
	pointsOfFunctionPlot[518].Y = 1.146

	pointsOfFunctionPlot[519].X = 5.19
	pointsOfFunctionPlot[519].Y = 1.147

	pointsOfFunctionPlot[520].X = 5.20
	pointsOfFunctionPlot[520].Y = 1.147

	pointsOfFunctionPlot[521].X = 5.21
	pointsOfFunctionPlot[521].Y = 1.147

	pointsOfFunctionPlot[522].X = 5.22
	pointsOfFunctionPlot[522].Y = 1.147

	pointsOfFunctionPlot[523].X = 5.23
	pointsOfFunctionPlot[523].Y = 1.147

	pointsOfFunctionPlot[524].X = 5.24
	pointsOfFunctionPlot[524].Y = 1.148

	pointsOfFunctionPlot[525].X = 5.25
	pointsOfFunctionPlot[525].Y = 1.148

	pointsOfFunctionPlot[526].X = 5.26
	pointsOfFunctionPlot[526].Y = 1.148

	pointsOfFunctionPlot[527].X = 5.27
	pointsOfFunctionPlot[527].Y = 1.148

	pointsOfFunctionPlot[528].X = 5.28
	pointsOfFunctionPlot[528].Y = 1.148

	pointsOfFunctionPlot[529].X = 5.29
	pointsOfFunctionPlot[529].Y = 1.148

	pointsOfFunctionPlot[530].X = 5.30
	pointsOfFunctionPlot[530].Y = 1.149

	pointsOfFunctionPlot[531].X = 5.31
	pointsOfFunctionPlot[531].Y = 1.149

	pointsOfFunctionPlot[532].X = 5.32
	pointsOfFunctionPlot[532].Y = 1.149

	pointsOfFunctionPlot[533].X = 5.33
	pointsOfFunctionPlot[533].Y = 1.149

	pointsOfFunctionPlot[534].X = 5.34
	pointsOfFunctionPlot[534].Y = 1.149

	pointsOfFunctionPlot[535].X = 5.35
	pointsOfFunctionPlot[535].Y = 1.149

	pointsOfFunctionPlot[536].X = 5.36
	pointsOfFunctionPlot[536].Y = 1.15

	pointsOfFunctionPlot[537].X = 5.37
	pointsOfFunctionPlot[537].Y = 1.15

	pointsOfFunctionPlot[538].X = 5.38
	pointsOfFunctionPlot[538].Y = 1.15

	pointsOfFunctionPlot[539].X = 5.39
	pointsOfFunctionPlot[539].Y = 1.15

	pointsOfFunctionPlot[540].X = 5.40
	pointsOfFunctionPlot[540].Y = 1.15

	pointsOfFunctionPlot[541].X = 5.41
	pointsOfFunctionPlot[541].Y = 1.151

	pointsOfFunctionPlot[542].X = 5.42
	pointsOfFunctionPlot[542].Y = 1.151

	pointsOfFunctionPlot[543].X = 5.43
	pointsOfFunctionPlot[543].Y = 1.151

	pointsOfFunctionPlot[544].X = 5.44
	pointsOfFunctionPlot[544].Y = 1.151

	pointsOfFunctionPlot[545].X = 5.45
	pointsOfFunctionPlot[545].Y = 1.151

	pointsOfFunctionPlot[546].X = 5.46
	pointsOfFunctionPlot[546].Y = 1.151

	pointsOfFunctionPlot[547].X = 5.47
	pointsOfFunctionPlot[547].Y = 1.152

	pointsOfFunctionPlot[548].X = 5.48
	pointsOfFunctionPlot[548].Y = 1.152

	pointsOfFunctionPlot[549].X = 5.49
	pointsOfFunctionPlot[549].Y = 1.152

	pointsOfFunctionPlot[550].X = 5.50
	pointsOfFunctionPlot[550].Y = 1.152

	pointsOfFunctionPlot[551].X = 5.51
	pointsOfFunctionPlot[551].Y = 1.152

	pointsOfFunctionPlot[552].X = 5.52
	pointsOfFunctionPlot[552].Y = 1.152

	pointsOfFunctionPlot[553].X = 5.53
	pointsOfFunctionPlot[553].Y = 1.153

	pointsOfFunctionPlot[554].X = 5.54
	pointsOfFunctionPlot[554].Y = 1.153

	pointsOfFunctionPlot[555].X = 5.55
	pointsOfFunctionPlot[555].Y = 1.153

	pointsOfFunctionPlot[556].X = 5.56
	pointsOfFunctionPlot[556].Y = 1.153

	pointsOfFunctionPlot[557].X = 5.57
	pointsOfFunctionPlot[557].Y = 1.153

	pointsOfFunctionPlot[558].X = 5.58
	pointsOfFunctionPlot[558].Y = 1.154

	pointsOfFunctionPlot[559].X = 5.59
	pointsOfFunctionPlot[559].Y = 1.154

	pointsOfFunctionPlot[560].X = 5.60
	pointsOfFunctionPlot[560].Y = 1.154

	pointsOfFunctionPlot[561].X = 5.61
	pointsOfFunctionPlot[561].Y = 1.154

	pointsOfFunctionPlot[562].X = 5.62
	pointsOfFunctionPlot[562].Y = 1.154

	pointsOfFunctionPlot[563].X = 5.63
	pointsOfFunctionPlot[563].Y = 1.154

	pointsOfFunctionPlot[564].X = 5.64
	pointsOfFunctionPlot[564].Y = 1.155

	pointsOfFunctionPlot[565].X = 5.65
	pointsOfFunctionPlot[565].Y = 1.155

	pointsOfFunctionPlot[566].X = 5.66
	pointsOfFunctionPlot[566].Y = 1.155

	pointsOfFunctionPlot[567].X = 5.67
	pointsOfFunctionPlot[567].Y = 1.155

	pointsOfFunctionPlot[568].X = 5.68
	pointsOfFunctionPlot[568].Y = 1.155

	pointsOfFunctionPlot[569].X = 5.69
	pointsOfFunctionPlot[569].Y = 1.155

	pointsOfFunctionPlot[570].X = 5.70
	pointsOfFunctionPlot[570].Y = 1.156

	pointsOfFunctionPlot[571].X = 5.71
	pointsOfFunctionPlot[571].Y = 1.156

	pointsOfFunctionPlot[572].X = 5.72
	pointsOfFunctionPlot[572].Y = 1.156

	pointsOfFunctionPlot[573].X = 5.73
	pointsOfFunctionPlot[573].Y = 1.156

	pointsOfFunctionPlot[574].X = 5.74
	pointsOfFunctionPlot[574].Y = 1.156

	pointsOfFunctionPlot[575].X = 5.75
	pointsOfFunctionPlot[575].Y = 1.156

	pointsOfFunctionPlot[576].X = 5.76
	pointsOfFunctionPlot[576].Y = 1.157

	pointsOfFunctionPlot[577].X = 5.77
	pointsOfFunctionPlot[577].Y = 1.157

	pointsOfFunctionPlot[578].X = 5.78
	pointsOfFunctionPlot[578].Y = 1.157

	pointsOfFunctionPlot[579].X = 5.79
	pointsOfFunctionPlot[579].Y = 1.157

	pointsOfFunctionPlot[580].X = 5.80
	pointsOfFunctionPlot[580].Y = 1.157

	pointsOfFunctionPlot[581].X = 5.81
	pointsOfFunctionPlot[581].Y = 1.157

	pointsOfFunctionPlot[582].X = 5.82
	pointsOfFunctionPlot[582].Y = 1.158

	pointsOfFunctionPlot[583].X = 5.83
	pointsOfFunctionPlot[583].Y = 1.158

	pointsOfFunctionPlot[584].X = 5.84
	pointsOfFunctionPlot[584].Y = 1.158

	pointsOfFunctionPlot[585].X = 5.85
	pointsOfFunctionPlot[585].Y = 1.158

	pointsOfFunctionPlot[586].X = 5.86
	pointsOfFunctionPlot[586].Y = 1.158

	pointsOfFunctionPlot[587].X = 5.87
	pointsOfFunctionPlot[587].Y = 1.158

	pointsOfFunctionPlot[588].X = 5.88
	pointsOfFunctionPlot[588].Y = 1.159

	pointsOfFunctionPlot[589].X = 5.89
	pointsOfFunctionPlot[589].Y = 1.159

	pointsOfFunctionPlot[590].X = 5.90
	pointsOfFunctionPlot[590].Y = 1.159

	pointsOfFunctionPlot[591].X = 5.91
	pointsOfFunctionPlot[591].Y = 1.159

	pointsOfFunctionPlot[592].X = 5.92
	pointsOfFunctionPlot[592].Y = 1.159

	pointsOfFunctionPlot[593].X = 5.93
	pointsOfFunctionPlot[593].Y = 1.159

	pointsOfFunctionPlot[594].X = 5.94
	pointsOfFunctionPlot[594].Y = 1.16

	pointsOfFunctionPlot[595].X = 5.95
	pointsOfFunctionPlot[595].Y = 1.16

	pointsOfFunctionPlot[596].X = 5.96
	pointsOfFunctionPlot[596].Y = 1.16

	pointsOfFunctionPlot[597].X = 5.97
	pointsOfFunctionPlot[597].Y = 1.16

	pointsOfFunctionPlot[598].X = 5.98
	pointsOfFunctionPlot[598].Y = 1.16

	pointsOfFunctionPlot[599].X = 5.99
	pointsOfFunctionPlot[599].Y = 1.16

	pointsOfFunctionPlot[600].X = 6.0
	pointsOfFunctionPlot[600].Y = 1.161

	pointsOfFunctionPlot[601].X = 6.01
	pointsOfFunctionPlot[601].Y = 1.161

	pointsOfFunctionPlot[602].X = 6.02
	pointsOfFunctionPlot[602].Y = 1.161

	pointsOfFunctionPlot[603].X = 6.03
	pointsOfFunctionPlot[603].Y = 1.161

	pointsOfFunctionPlot[604].X = 6.04
	pointsOfFunctionPlot[604].Y = 1.161

	pointsOfFunctionPlot[605].X = 6.05
	pointsOfFunctionPlot[605].Y = 1.161

	pointsOfFunctionPlot[606].X = 6.06
	pointsOfFunctionPlot[606].Y = 1.161

	pointsOfFunctionPlot[607].X = 6.07
	pointsOfFunctionPlot[607].Y = 1.162

	pointsOfFunctionPlot[608].X = 6.08
	pointsOfFunctionPlot[608].Y = 1.162

	pointsOfFunctionPlot[609].X = 6.09
	pointsOfFunctionPlot[609].Y = 1.162

	pointsOfFunctionPlot[610].X = 6.10
	pointsOfFunctionPlot[610].Y = 1.162

	pointsOfFunctionPlot[611].X = 6.11
	pointsOfFunctionPlot[611].Y = 1.162

	pointsOfFunctionPlot[612].X = 6.12
	pointsOfFunctionPlot[612].Y = 1.162

	pointsOfFunctionPlot[613].X = 6.13
	pointsOfFunctionPlot[613].Y = 1.163

	pointsOfFunctionPlot[614].X = 6.14
	pointsOfFunctionPlot[614].Y = 1.163

	pointsOfFunctionPlot[615].X = 6.15
	pointsOfFunctionPlot[615].Y = 1.163

	pointsOfFunctionPlot[616].X = 6.16
	pointsOfFunctionPlot[616].Y = 1.163

	pointsOfFunctionPlot[617].X = 6.17
	pointsOfFunctionPlot[617].Y = 1.163

	pointsOfFunctionPlot[618].X = 6.18
	pointsOfFunctionPlot[618].Y = 1.163

	pointsOfFunctionPlot[619].X = 6.19
	pointsOfFunctionPlot[619].Y = 1.164

	pointsOfFunctionPlot[620].X = 6.20
	pointsOfFunctionPlot[620].Y = 1.164

	pointsOfFunctionPlot[621].X = 6.21
	pointsOfFunctionPlot[621].Y = 1.164

	pointsOfFunctionPlot[622].X = 6.22
	pointsOfFunctionPlot[622].Y = 1.164

	pointsOfFunctionPlot[623].X = 6.23
	pointsOfFunctionPlot[623].Y = 1.164

	pointsOfFunctionPlot[624].X = 6.24
	pointsOfFunctionPlot[624].Y = 1.164

	pointsOfFunctionPlot[625].X = 6.25
	pointsOfFunctionPlot[625].Y = 1.164

	pointsOfFunctionPlot[626].X = 6.26
	pointsOfFunctionPlot[626].Y = 1.165

	pointsOfFunctionPlot[627].X = 6.27
	pointsOfFunctionPlot[627].Y = 1.165

	pointsOfFunctionPlot[628].X = 6.28
	pointsOfFunctionPlot[628].Y = 1.165

	pointsOfFunctionPlot[629].X = 6.29
	pointsOfFunctionPlot[629].Y = 1.165

	pointsOfFunctionPlot[630].X = 6.30
	pointsOfFunctionPlot[630].Y = 1.165

	pointsOfFunctionPlot[631].X = 6.31
	pointsOfFunctionPlot[631].Y = 1.165

	pointsOfFunctionPlot[632].X = 6.32
	pointsOfFunctionPlot[632].Y = 1.166

	pointsOfFunctionPlot[633].X = 6.33
	pointsOfFunctionPlot[633].Y = 1.166

	pointsOfFunctionPlot[634].X = 6.34
	pointsOfFunctionPlot[634].Y = 1.166

	pointsOfFunctionPlot[635].X = 6.35
	pointsOfFunctionPlot[635].Y = 1.166

	pointsOfFunctionPlot[636].X = 6.36
	pointsOfFunctionPlot[636].Y = 1.166

	pointsOfFunctionPlot[637].X = 6.37
	pointsOfFunctionPlot[637].Y = 1.166

	pointsOfFunctionPlot[638].X = 6.38
	pointsOfFunctionPlot[638].Y = 1.166

	pointsOfFunctionPlot[639].X = 6.39
	pointsOfFunctionPlot[639].Y = 1.167

	pointsOfFunctionPlot[640].X = 6.40
	pointsOfFunctionPlot[640].Y = 1.167

	pointsOfFunctionPlot[641].X = 6.41
	pointsOfFunctionPlot[641].Y = 1.167

	pointsOfFunctionPlot[642].X = 6.42
	pointsOfFunctionPlot[642].Y = 1.167

	pointsOfFunctionPlot[643].X = 6.43
	pointsOfFunctionPlot[643].Y = 1.167

	pointsOfFunctionPlot[644].X = 6.44
	pointsOfFunctionPlot[644].Y = 1.167

	pointsOfFunctionPlot[645].X = 6.45
	pointsOfFunctionPlot[645].Y = 1.168

	pointsOfFunctionPlot[646].X = 6.46
	pointsOfFunctionPlot[646].Y = 1.168

	pointsOfFunctionPlot[647].X = 6.47
	pointsOfFunctionPlot[647].Y = 1.168

	pointsOfFunctionPlot[648].X = 6.48
	pointsOfFunctionPlot[648].Y = 1.168

	pointsOfFunctionPlot[649].X = 6.49
	pointsOfFunctionPlot[649].Y = 1.168

	pointsOfFunctionPlot[650].X = 6.50
	pointsOfFunctionPlot[650].Y = 1.168

	pointsOfFunctionPlot[651].X = 6.51
	pointsOfFunctionPlot[651].Y = 1.168

	pointsOfFunctionPlot[652].X = 6.52
	pointsOfFunctionPlot[652].Y = 1.169

	pointsOfFunctionPlot[653].X = 6.53
	pointsOfFunctionPlot[653].Y = 1.169

	pointsOfFunctionPlot[654].X = 6.54
	pointsOfFunctionPlot[654].Y = 1.169

	pointsOfFunctionPlot[655].X = 6.55
	pointsOfFunctionPlot[655].Y = 1.169

	pointsOfFunctionPlot[656].X = 6.56
	pointsOfFunctionPlot[656].Y = 1.169

	pointsOfFunctionPlot[657].X = 6.57
	pointsOfFunctionPlot[657].Y = 1.169

	pointsOfFunctionPlot[658].X = 6.58
	pointsOfFunctionPlot[658].Y = 1.169

	pointsOfFunctionPlot[659].X = 6.59
	pointsOfFunctionPlot[659].Y = 1.17

	pointsOfFunctionPlot[660].X = 6.60
	pointsOfFunctionPlot[660].Y = 1.17

	pointsOfFunctionPlot[661].X = 6.61
	pointsOfFunctionPlot[661].Y = 1.17

	pointsOfFunctionPlot[662].X = 6.62
	pointsOfFunctionPlot[662].Y = 1.17

	pointsOfFunctionPlot[663].X = 6.63
	pointsOfFunctionPlot[663].Y = 1.17

	pointsOfFunctionPlot[664].X = 6.64
	pointsOfFunctionPlot[664].Y = 1.17

	pointsOfFunctionPlot[665].X = 6.65
	pointsOfFunctionPlot[665].Y = 1.171

	pointsOfFunctionPlot[666].X = 6.66
	pointsOfFunctionPlot[666].Y = 1.171

	pointsOfFunctionPlot[667].X = 6.67
	pointsOfFunctionPlot[667].Y = 1.171

	pointsOfFunctionPlot[668].X = 6.68
	pointsOfFunctionPlot[668].Y = 1.171

	pointsOfFunctionPlot[669].X = 6.69
	pointsOfFunctionPlot[669].Y = 1.171

	pointsOfFunctionPlot[670].X = 6.70
	pointsOfFunctionPlot[670].Y = 1.171

	pointsOfFunctionPlot[671].X = 6.71
	pointsOfFunctionPlot[671].Y = 1.171

	pointsOfFunctionPlot[672].X = 6.72
	pointsOfFunctionPlot[672].Y = 1.172

	pointsOfFunctionPlot[673].X = 6.73
	pointsOfFunctionPlot[673].Y = 1.172

	pointsOfFunctionPlot[674].X = 6.74
	pointsOfFunctionPlot[674].Y = 1.172

	pointsOfFunctionPlot[675].X = 6.75
	pointsOfFunctionPlot[675].Y = 1.172

	pointsOfFunctionPlot[676].X = 6.76
	pointsOfFunctionPlot[676].Y = 1.172

	pointsOfFunctionPlot[677].X = 6.77
	pointsOfFunctionPlot[677].Y = 1.172

	pointsOfFunctionPlot[678].X = 6.78
	pointsOfFunctionPlot[678].Y = 1.172

	pointsOfFunctionPlot[679].X = 6.79
	pointsOfFunctionPlot[679].Y = 1.173

	pointsOfFunctionPlot[680].X = 6.80
	pointsOfFunctionPlot[680].Y = 1.173

	pointsOfFunctionPlot[681].X = 6.81
	pointsOfFunctionPlot[681].Y = 1.173

	pointsOfFunctionPlot[682].X = 6.82
	pointsOfFunctionPlot[682].Y = 1.173

	pointsOfFunctionPlot[683].X = 6.83
	pointsOfFunctionPlot[683].Y = 1.173

	pointsOfFunctionPlot[684].X = 6.84
	pointsOfFunctionPlot[684].Y = 1.173

	pointsOfFunctionPlot[685].X = 6.85
	pointsOfFunctionPlot[685].Y = 1.173

	pointsOfFunctionPlot[686].X = 6.86
	pointsOfFunctionPlot[686].Y = 1.174

	pointsOfFunctionPlot[687].X = 6.87
	pointsOfFunctionPlot[687].Y = 1.174

	pointsOfFunctionPlot[688].X = 6.88
	pointsOfFunctionPlot[688].Y = 1.174

	pointsOfFunctionPlot[689].X = 6.89
	pointsOfFunctionPlot[689].Y = 1.174

	pointsOfFunctionPlot[690].X = 6.90
	pointsOfFunctionPlot[690].Y = 1.174

	pointsOfFunctionPlot[691].X = 6.91
	pointsOfFunctionPlot[691].Y = 1.174

	pointsOfFunctionPlot[692].X = 6.92
	pointsOfFunctionPlot[692].Y = 1.174

	pointsOfFunctionPlot[693].X = 6.93
	pointsOfFunctionPlot[693].Y = 1.175

	pointsOfFunctionPlot[694].X = 6.94
	pointsOfFunctionPlot[694].Y = 1.175

	pointsOfFunctionPlot[695].X = 6.95
	pointsOfFunctionPlot[695].Y = 1.175

	pointsOfFunctionPlot[696].X = 6.96
	pointsOfFunctionPlot[696].Y = 1.175

	pointsOfFunctionPlot[697].X = 6.97
	pointsOfFunctionPlot[697].Y = 1.175

	pointsOfFunctionPlot[698].X = 6.98
	pointsOfFunctionPlot[698].Y = 1.175

	pointsOfFunctionPlot[699].X = 6.99
	pointsOfFunctionPlot[699].Y = 1.175

	pointsOfFunctionPlot[700].X = 7.0
	pointsOfFunctionPlot[700].Y = 1.176

	pointsOfFunctionPlot[701].X = 7.01
	pointsOfFunctionPlot[701].Y = 1.176

	pointsOfFunctionPlot[702].X = 7.02
	pointsOfFunctionPlot[702].Y = 1.176

	pointsOfFunctionPlot[703].X = 7.03
	pointsOfFunctionPlot[703].Y = 1.176

	pointsOfFunctionPlot[704].X = 7.04
	pointsOfFunctionPlot[704].Y = 1.176

	pointsOfFunctionPlot[705].X = 7.05
	pointsOfFunctionPlot[705].Y = 1.176

	pointsOfFunctionPlot[706].X = 7.06
	pointsOfFunctionPlot[706].Y = 1.176

	pointsOfFunctionPlot[707].X = 7.07
	pointsOfFunctionPlot[707].Y = 1.177

	pointsOfFunctionPlot[708].X = 7.08
	pointsOfFunctionPlot[708].Y = 1.177

	pointsOfFunctionPlot[709].X = 7.09
	pointsOfFunctionPlot[709].Y = 1.177

	pointsOfFunctionPlot[710].X = 7.10
	pointsOfFunctionPlot[710].Y = 1.177

	pointsOfFunctionPlot[711].X = 7.11
	pointsOfFunctionPlot[711].Y = 1.177

	pointsOfFunctionPlot[712].X = 7.12
	pointsOfFunctionPlot[712].Y = 1.177

	pointsOfFunctionPlot[713].X = 7.13
	pointsOfFunctionPlot[713].Y = 1.177

	pointsOfFunctionPlot[714].X = 7.14
	pointsOfFunctionPlot[714].Y = 1.177

	pointsOfFunctionPlot[715].X = 7.15
	pointsOfFunctionPlot[715].Y = 1.178

	pointsOfFunctionPlot[716].X = 7.16
	pointsOfFunctionPlot[716].Y = 1.178

	pointsOfFunctionPlot[717].X = 7.17
	pointsOfFunctionPlot[717].Y = 1.178

	pointsOfFunctionPlot[718].X = 7.18
	pointsOfFunctionPlot[718].Y = 1.178

	pointsOfFunctionPlot[719].X = 7.19
	pointsOfFunctionPlot[719].Y = 1.178

	pointsOfFunctionPlot[720].X = 7.20
	pointsOfFunctionPlot[720].Y = 1.178

	pointsOfFunctionPlot[721].X = 7.21
	pointsOfFunctionPlot[721].Y = 1.178

	pointsOfFunctionPlot[722].X = 7.22
	pointsOfFunctionPlot[722].Y = 1.179

	pointsOfFunctionPlot[723].X = 7.23
	pointsOfFunctionPlot[723].Y = 1.179

	pointsOfFunctionPlot[724].X = 7.24
	pointsOfFunctionPlot[724].Y = 1.179

	pointsOfFunctionPlot[725].X = 7.25
	pointsOfFunctionPlot[725].Y = 1.179

	pointsOfFunctionPlot[726].X = 7.26
	pointsOfFunctionPlot[726].Y = 1.179

	pointsOfFunctionPlot[727].X = 7.27
	pointsOfFunctionPlot[727].Y = 1.179

	pointsOfFunctionPlot[728].X = 7.28
	pointsOfFunctionPlot[728].Y = 1.179

	pointsOfFunctionPlot[729].X = 7.29
	pointsOfFunctionPlot[729].Y = 1.18

	pointsOfFunctionPlot[730].X = 7.30
	pointsOfFunctionPlot[730].Y = 1.18

	pointsOfFunctionPlot[731].X = 7.31
	pointsOfFunctionPlot[731].Y = 1.18

	pointsOfFunctionPlot[732].X = 7.32
	pointsOfFunctionPlot[732].Y = 1.18

	pointsOfFunctionPlot[733].X = 7.33
	pointsOfFunctionPlot[733].Y = 1.18

	pointsOfFunctionPlot[734].X = 7.34
	pointsOfFunctionPlot[734].Y = 1.18

	pointsOfFunctionPlot[735].X = 7.35
	pointsOfFunctionPlot[735].Y = 1.18

	pointsOfFunctionPlot[736].X = 7.36
	pointsOfFunctionPlot[736].Y = 1.18

	pointsOfFunctionPlot[737].X = 7.37
	pointsOfFunctionPlot[737].Y = 1.181

	pointsOfFunctionPlot[738].X = 7.38
	pointsOfFunctionPlot[738].Y = 1.181

	pointsOfFunctionPlot[739].X = 7.39
	pointsOfFunctionPlot[739].Y = 1.181

	pointsOfFunctionPlot[740].X = 7.40
	pointsOfFunctionPlot[740].Y = 1.181

	pointsOfFunctionPlot[741].X = 7.41
	pointsOfFunctionPlot[741].Y = 1.181

	pointsOfFunctionPlot[742].X = 7.42
	pointsOfFunctionPlot[742].Y = 1.181

	pointsOfFunctionPlot[743].X = 7.43
	pointsOfFunctionPlot[743].Y = 1.181

	pointsOfFunctionPlot[744].X = 7.44
	pointsOfFunctionPlot[744].Y = 1.182

	pointsOfFunctionPlot[745].X = 7.45
	pointsOfFunctionPlot[745].Y = 1.182

	pointsOfFunctionPlot[746].X = 7.46
	pointsOfFunctionPlot[746].Y = 1.182

	pointsOfFunctionPlot[747].X = 7.47
	pointsOfFunctionPlot[747].Y = 1.182

	pointsOfFunctionPlot[748].X = 7.48
	pointsOfFunctionPlot[748].Y = 1.182

	pointsOfFunctionPlot[749].X = 7.49
	pointsOfFunctionPlot[749].Y = 1.182

	pointsOfFunctionPlot[750].X = 7.50
	pointsOfFunctionPlot[750].Y = 1.182

	pointsOfFunctionPlot[751].X = 7.51
	pointsOfFunctionPlot[751].Y = 1.182

	pointsOfFunctionPlot[752].X = 7.52
	pointsOfFunctionPlot[752].Y = 1.183

	pointsOfFunctionPlot[753].X = 7.53
	pointsOfFunctionPlot[753].Y = 1.183

	pointsOfFunctionPlot[754].X = 7.54
	pointsOfFunctionPlot[754].Y = 1.183

	pointsOfFunctionPlot[755].X = 7.55
	pointsOfFunctionPlot[755].Y = 1.183

	pointsOfFunctionPlot[756].X = 7.56
	pointsOfFunctionPlot[756].Y = 1.183

	pointsOfFunctionPlot[757].X = 7.57
	pointsOfFunctionPlot[757].Y = 1.183

	pointsOfFunctionPlot[758].X = 7.58
	pointsOfFunctionPlot[758].Y = 1.183

	pointsOfFunctionPlot[759].X = 7.59
	pointsOfFunctionPlot[759].Y = 1.184

	pointsOfFunctionPlot[760].X = 7.60
	pointsOfFunctionPlot[760].Y = 1.184

	pointsOfFunctionPlot[761].X = 7.61
	pointsOfFunctionPlot[761].Y = 1.184

	pointsOfFunctionPlot[762].X = 7.62
	pointsOfFunctionPlot[762].Y = 1.184

	pointsOfFunctionPlot[763].X = 7.63
	pointsOfFunctionPlot[763].Y = 1.184

	pointsOfFunctionPlot[764].X = 7.64
	pointsOfFunctionPlot[764].Y = 1.184

	pointsOfFunctionPlot[765].X = 7.65
	pointsOfFunctionPlot[765].Y = 1.184

	pointsOfFunctionPlot[766].X = 7.66
	pointsOfFunctionPlot[766].Y = 1.184

	pointsOfFunctionPlot[767].X = 7.67
	pointsOfFunctionPlot[767].Y = 1.185

	pointsOfFunctionPlot[768].X = 7.68
	pointsOfFunctionPlot[768].Y = 1.185

	pointsOfFunctionPlot[769].X = 7.69
	pointsOfFunctionPlot[769].Y = 1.185

	pointsOfFunctionPlot[770].X = 7.70
	pointsOfFunctionPlot[770].Y = 1.185

	pointsOfFunctionPlot[771].X = 7.71
	pointsOfFunctionPlot[771].Y = 1.185

	pointsOfFunctionPlot[772].X = 7.72
	pointsOfFunctionPlot[772].Y = 1.185

	pointsOfFunctionPlot[773].X = 7.73
	pointsOfFunctionPlot[773].Y = 1.185

	pointsOfFunctionPlot[774].X = 7.74
	pointsOfFunctionPlot[774].Y = 1.185

	pointsOfFunctionPlot[775].X = 7.75
	pointsOfFunctionPlot[775].Y = 1.186

	pointsOfFunctionPlot[776].X = 7.76
	pointsOfFunctionPlot[776].Y = 1.186

	pointsOfFunctionPlot[777].X = 7.77
	pointsOfFunctionPlot[777].Y = 1.186

	pointsOfFunctionPlot[778].X = 7.78
	pointsOfFunctionPlot[778].Y = 1.186

	pointsOfFunctionPlot[779].X = 7.79
	pointsOfFunctionPlot[779].Y = 1.186

	pointsOfFunctionPlot[780].X = 7.80
	pointsOfFunctionPlot[780].Y = 1.186

	pointsOfFunctionPlot[781].X = 7.81
	pointsOfFunctionPlot[781].Y = 1.186

	pointsOfFunctionPlot[782].X = 7.82
	pointsOfFunctionPlot[782].Y = 1.186

	pointsOfFunctionPlot[783].X = 7.83
	pointsOfFunctionPlot[783].Y = 1.187

	pointsOfFunctionPlot[784].X = 7.84
	pointsOfFunctionPlot[784].Y = 1.187

	pointsOfFunctionPlot[785].X = 7.85
	pointsOfFunctionPlot[785].Y = 1.187

	pointsOfFunctionPlot[786].X = 7.86
	pointsOfFunctionPlot[786].Y = 1.187

	pointsOfFunctionPlot[787].X = 7.87
	pointsOfFunctionPlot[787].Y = 1.187

	pointsOfFunctionPlot[788].X = 7.88
	pointsOfFunctionPlot[788].Y = 1.187

	pointsOfFunctionPlot[789].X = 7.89
	pointsOfFunctionPlot[789].Y = 1.187

	pointsOfFunctionPlot[790].X = 7.90
	pointsOfFunctionPlot[790].Y = 1.187

	pointsOfFunctionPlot[791].X = 7.91
	pointsOfFunctionPlot[791].Y = 1.188

	pointsOfFunctionPlot[792].X = 7.92
	pointsOfFunctionPlot[792].Y = 1.188

	pointsOfFunctionPlot[793].X = 7.93
	pointsOfFunctionPlot[793].Y = 1.188

	pointsOfFunctionPlot[794].X = 7.94
	pointsOfFunctionPlot[794].Y = 1.188

	pointsOfFunctionPlot[795].X = 7.95
	pointsOfFunctionPlot[795].Y = 1.188

	pointsOfFunctionPlot[796].X = 7.96
	pointsOfFunctionPlot[796].Y = 1.188

	pointsOfFunctionPlot[797].X = 7.97
	pointsOfFunctionPlot[797].Y = 1.188

	pointsOfFunctionPlot[798].X = 7.98
	pointsOfFunctionPlot[798].Y = 1.188

	pointsOfFunctionPlot[799].X = 7.99
	pointsOfFunctionPlot[799].Y = 1.189

	pointsOfFunctionPlot[800].X = 8.0
	pointsOfFunctionPlot[800].Y = 1.189

	pointsOfFunctionPlot[801].X = 8.01
	pointsOfFunctionPlot[801].Y = 1.189

	pointsOfFunctionPlot[802].X = 8.02
	pointsOfFunctionPlot[802].Y = 1.189

	pointsOfFunctionPlot[803].X = 8.03
	pointsOfFunctionPlot[803].Y = 1.189

	pointsOfFunctionPlot[804].X = 8.04
	pointsOfFunctionPlot[804].Y = 1.189

	pointsOfFunctionPlot[805].X = 8.05
	pointsOfFunctionPlot[805].Y = 1.189

	pointsOfFunctionPlot[806].X = 8.06
	pointsOfFunctionPlot[806].Y = 1.189

	pointsOfFunctionPlot[807].X = 8.07
	pointsOfFunctionPlot[807].Y = 1.19

	pointsOfFunctionPlot[808].X = 8.08
	pointsOfFunctionPlot[808].Y = 1.19

	pointsOfFunctionPlot[809].X = 8.09
	pointsOfFunctionPlot[809].Y = 1.19

	pointsOfFunctionPlot[810].X = 8.10
	pointsOfFunctionPlot[810].Y = 1.19

	pointsOfFunctionPlot[811].X = 8.11
	pointsOfFunctionPlot[811].Y = 1.19

	pointsOfFunctionPlot[812].X = 8.12
	pointsOfFunctionPlot[812].Y = 1.19

	pointsOfFunctionPlot[813].X = 8.13
	pointsOfFunctionPlot[813].Y = 1.19

	pointsOfFunctionPlot[814].X = 8.14
	pointsOfFunctionPlot[814].Y = 1.19

	pointsOfFunctionPlot[815].X = 8.15
	pointsOfFunctionPlot[815].Y = 1.191

	pointsOfFunctionPlot[816].X = 8.16
	pointsOfFunctionPlot[816].Y = 1.191

	pointsOfFunctionPlot[817].X = 8.17
	pointsOfFunctionPlot[817].Y = 1.191

	pointsOfFunctionPlot[818].X = 8.18
	pointsOfFunctionPlot[818].Y = 1.191

	pointsOfFunctionPlot[819].X = 8.19
	pointsOfFunctionPlot[819].Y = 1.191

	pointsOfFunctionPlot[820].X = 8.20
	pointsOfFunctionPlot[820].Y = 1.191

	pointsOfFunctionPlot[821].X = 8.21
	pointsOfFunctionPlot[821].Y = 1.191

	pointsOfFunctionPlot[822].X = 8.22
	pointsOfFunctionPlot[822].Y = 1.191

	pointsOfFunctionPlot[823].X = 8.23
	pointsOfFunctionPlot[823].Y = 1.192

	pointsOfFunctionPlot[824].X = 8.24
	pointsOfFunctionPlot[824].Y = 1.192

	pointsOfFunctionPlot[825].X = 8.25
	pointsOfFunctionPlot[825].Y = 1.192

	pointsOfFunctionPlot[826].X = 8.26
	pointsOfFunctionPlot[826].Y = 1.192

	pointsOfFunctionPlot[827].X = 8.27
	pointsOfFunctionPlot[827].Y = 1.192

	pointsOfFunctionPlot[828].X = 8.28
	pointsOfFunctionPlot[828].Y = 1.192

	pointsOfFunctionPlot[829].X = 8.29
	pointsOfFunctionPlot[829].Y = 1.192

	pointsOfFunctionPlot[830].X = 8.30
	pointsOfFunctionPlot[830].Y = 1.192

	pointsOfFunctionPlot[831].X = 8.31
	pointsOfFunctionPlot[831].Y = 1.192

	pointsOfFunctionPlot[832].X = 8.32
	pointsOfFunctionPlot[832].Y = 1.193

	pointsOfFunctionPlot[833].X = 8.33
	pointsOfFunctionPlot[833].Y = 1.193

	pointsOfFunctionPlot[834].X = 8.34
	pointsOfFunctionPlot[834].Y = 1.193

	pointsOfFunctionPlot[835].X = 8.35
	pointsOfFunctionPlot[835].Y = 1.193

	pointsOfFunctionPlot[836].X = 8.36
	pointsOfFunctionPlot[836].Y = 1.193

	pointsOfFunctionPlot[837].X = 8.37
	pointsOfFunctionPlot[837].Y = 1.193

	pointsOfFunctionPlot[838].X = 8.38
	pointsOfFunctionPlot[838].Y = 1.193

	pointsOfFunctionPlot[839].X = 8.39
	pointsOfFunctionPlot[839].Y = 1.193

	pointsOfFunctionPlot[840].X = 8.40
	pointsOfFunctionPlot[840].Y = 1.194

	pointsOfFunctionPlot[841].X = 8.41
	pointsOfFunctionPlot[841].Y = 1.194

	pointsOfFunctionPlot[842].X = 8.42
	pointsOfFunctionPlot[842].Y = 1.194

	pointsOfFunctionPlot[843].X = 8.43
	pointsOfFunctionPlot[843].Y = 1.194

	pointsOfFunctionPlot[844].X = 8.44
	pointsOfFunctionPlot[844].Y = 1.194

	pointsOfFunctionPlot[845].X = 8.45
	pointsOfFunctionPlot[845].Y = 1.194

	pointsOfFunctionPlot[846].X = 8.46
	pointsOfFunctionPlot[846].Y = 1.194

	pointsOfFunctionPlot[847].X = 8.47
	pointsOfFunctionPlot[847].Y = 1.194

	pointsOfFunctionPlot[848].X = 8.48
	pointsOfFunctionPlot[848].Y = 1.194

	pointsOfFunctionPlot[849].X = 8.49
	pointsOfFunctionPlot[849].Y = 1.195

	pointsOfFunctionPlot[850].X = 8.50
	pointsOfFunctionPlot[850].Y = 1.195

	pointsOfFunctionPlot[851].X = 8.51
	pointsOfFunctionPlot[851].Y = 1.195

	pointsOfFunctionPlot[852].X = 8.52
	pointsOfFunctionPlot[852].Y = 1.195

	pointsOfFunctionPlot[853].X = 8.53
	pointsOfFunctionPlot[853].Y = 1.195

	pointsOfFunctionPlot[854].X = 8.54
	pointsOfFunctionPlot[854].Y = 1.195

	pointsOfFunctionPlot[855].X = 8.55
	pointsOfFunctionPlot[855].Y = 1.195

	pointsOfFunctionPlot[856].X = 8.56
	pointsOfFunctionPlot[856].Y = 1.195

	pointsOfFunctionPlot[857].X = 8.57
	pointsOfFunctionPlot[857].Y = 1.196

	pointsOfFunctionPlot[858].X = 8.58
	pointsOfFunctionPlot[858].Y = 1.196

	pointsOfFunctionPlot[859].X = 8.59
	pointsOfFunctionPlot[859].Y = 1.196

	pointsOfFunctionPlot[860].X = 8.60
	pointsOfFunctionPlot[860].Y = 1.196

	pointsOfFunctionPlot[861].X = 8.61
	pointsOfFunctionPlot[861].Y = 1.196

	pointsOfFunctionPlot[862].X = 8.62
	pointsOfFunctionPlot[862].Y = 1.196

	pointsOfFunctionPlot[863].X = 8.63
	pointsOfFunctionPlot[863].Y = 1.196

	pointsOfFunctionPlot[864].X = 8.64
	pointsOfFunctionPlot[864].Y = 1.196

	pointsOfFunctionPlot[865].X = 8.65
	pointsOfFunctionPlot[865].Y = 1.196

	pointsOfFunctionPlot[866].X = 8.66
	pointsOfFunctionPlot[866].Y = 1.197

	pointsOfFunctionPlot[867].X = 8.67
	pointsOfFunctionPlot[867].Y = 1.197

	pointsOfFunctionPlot[868].X = 8.68
	pointsOfFunctionPlot[868].Y = 1.197

	pointsOfFunctionPlot[869].X = 8.69
	pointsOfFunctionPlot[869].Y = 1.197

	pointsOfFunctionPlot[870].X = 8.70
	pointsOfFunctionPlot[870].Y = 1.197

	pointsOfFunctionPlot[871].X = 8.71
	pointsOfFunctionPlot[871].Y = 1.197

	pointsOfFunctionPlot[872].X = 8.72
	pointsOfFunctionPlot[872].Y = 1.197

	pointsOfFunctionPlot[873].X = 8.73
	pointsOfFunctionPlot[873].Y = 1.197

	pointsOfFunctionPlot[874].X = 8.74
	pointsOfFunctionPlot[874].Y = 1.198

	pointsOfFunctionPlot[875].X = 8.75
	pointsOfFunctionPlot[875].Y = 1.198

	pointsOfFunctionPlot[876].X = 8.76
	pointsOfFunctionPlot[876].Y = 1.198

	pointsOfFunctionPlot[877].X = 8.77
	pointsOfFunctionPlot[877].Y = 1.198

	pointsOfFunctionPlot[878].X = 8.78
	pointsOfFunctionPlot[878].Y = 1.198

	pointsOfFunctionPlot[879].X = 8.79
	pointsOfFunctionPlot[879].Y = 1.198

	pointsOfFunctionPlot[880].X = 8.80
	pointsOfFunctionPlot[880].Y = 1.198

	pointsOfFunctionPlot[881].X = 8.81
	pointsOfFunctionPlot[881].Y = 1.198

	pointsOfFunctionPlot[882].X = 8.82
	pointsOfFunctionPlot[882].Y = 1.198

	pointsOfFunctionPlot[883].X = 8.83
	pointsOfFunctionPlot[883].Y = 1.199

	pointsOfFunctionPlot[884].X = 8.84
	pointsOfFunctionPlot[884].Y = 1.199

	pointsOfFunctionPlot[885].X = 8.85
	pointsOfFunctionPlot[885].Y = 1.199

	pointsOfFunctionPlot[886].X = 8.86
	pointsOfFunctionPlot[886].Y = 1.199

	pointsOfFunctionPlot[887].X = 8.87
	pointsOfFunctionPlot[887].Y = 1.199

	pointsOfFunctionPlot[888].X = 8.88
	pointsOfFunctionPlot[888].Y = 1.199

	pointsOfFunctionPlot[889].X = 8.89
	pointsOfFunctionPlot[889].Y = 1.199

	pointsOfFunctionPlot[890].X = 8.90
	pointsOfFunctionPlot[890].Y = 1.199

	pointsOfFunctionPlot[891].X = 8.91
	pointsOfFunctionPlot[891].Y = 1.199

	pointsOfFunctionPlot[892].X = 8.92
	pointsOfFunctionPlot[892].Y = 1.2

	pointsOfFunctionPlot[893].X = 8.93
	pointsOfFunctionPlot[893].Y = 1.2

	pointsOfFunctionPlot[894].X = 8.94
	pointsOfFunctionPlot[894].Y = 1.2

	pointsOfFunctionPlot[895].X = 8.95
	pointsOfFunctionPlot[895].Y = 1.2

	pointsOfFunctionPlot[896].X = 8.96
	pointsOfFunctionPlot[896].Y = 1.2

	pointsOfFunctionPlot[897].X = 8.97
	pointsOfFunctionPlot[897].Y = 1.2

	pointsOfFunctionPlot[898].X = 8.98
	pointsOfFunctionPlot[898].Y = 1.2

	pointsOfFunctionPlot[899].X = 8.99
	pointsOfFunctionPlot[899].Y = 1.2

	pointsOfFunctionPlot[900].X = 9.0
	pointsOfFunctionPlot[900].Y = 1.2

	pointsOfFunctionPlot[901].X = 9.01
	pointsOfFunctionPlot[901].Y = 1.201

	pointsOfFunctionPlot[902].X = 9.02
	pointsOfFunctionPlot[902].Y = 1.201

	pointsOfFunctionPlot[903].X = 9.03
	pointsOfFunctionPlot[903].Y = 1.201

	pointsOfFunctionPlot[904].X = 9.04
	pointsOfFunctionPlot[904].Y = 1.201

	pointsOfFunctionPlot[905].X = 9.05
	pointsOfFunctionPlot[905].Y = 1.201

	pointsOfFunctionPlot[906].X = 9.06
	pointsOfFunctionPlot[906].Y = 1.201

	pointsOfFunctionPlot[907].X = 9.07
	pointsOfFunctionPlot[907].Y = 1.201

	pointsOfFunctionPlot[908].X = 9.08
	pointsOfFunctionPlot[908].Y = 1.201

	pointsOfFunctionPlot[909].X = 9.09
	pointsOfFunctionPlot[909].Y = 1.201

	pointsOfFunctionPlot[910].X = 9.10
	pointsOfFunctionPlot[910].Y = 1.202

	pointsOfFunctionPlot[911].X = 9.11
	pointsOfFunctionPlot[911].Y = 1.202

	pointsOfFunctionPlot[912].X = 9.12
	pointsOfFunctionPlot[912].Y = 1.202

	pointsOfFunctionPlot[913].X = 9.13
	pointsOfFunctionPlot[913].Y = 1.202

	pointsOfFunctionPlot[914].X = 9.14
	pointsOfFunctionPlot[914].Y = 1.202

	pointsOfFunctionPlot[915].X = 9.15
	pointsOfFunctionPlot[915].Y = 1.202

	pointsOfFunctionPlot[916].X = 9.16
	pointsOfFunctionPlot[916].Y = 1.202

	pointsOfFunctionPlot[917].X = 9.17
	pointsOfFunctionPlot[917].Y = 1.202

	pointsOfFunctionPlot[918].X = 9.18
	pointsOfFunctionPlot[918].Y = 1.202

	pointsOfFunctionPlot[919].X = 9.19
	pointsOfFunctionPlot[919].Y = 1.203

	pointsOfFunctionPlot[920].X = 9.20
	pointsOfFunctionPlot[920].Y = 1.203

	pointsOfFunctionPlot[921].X = 9.21
	pointsOfFunctionPlot[921].Y = 1.203

	pointsOfFunctionPlot[922].X = 9.22
	pointsOfFunctionPlot[922].Y = 1.203

	pointsOfFunctionPlot[923].X = 9.23
	pointsOfFunctionPlot[923].Y = 1.203

	pointsOfFunctionPlot[924].X = 9.24
	pointsOfFunctionPlot[924].Y = 1.203

	pointsOfFunctionPlot[925].X = 9.25
	pointsOfFunctionPlot[925].Y = 1.203

	pointsOfFunctionPlot[926].X = 9.26
	pointsOfFunctionPlot[926].Y = 1.203

	pointsOfFunctionPlot[927].X = 9.27
	pointsOfFunctionPlot[927].Y = 1.203

	pointsOfFunctionPlot[928].X = 9.28
	pointsOfFunctionPlot[928].Y = 1.204

	pointsOfFunctionPlot[929].X = 9.29
	pointsOfFunctionPlot[929].Y = 1.204

	pointsOfFunctionPlot[930].X = 9.30
	pointsOfFunctionPlot[930].Y = 1.204

	pointsOfFunctionPlot[931].X = 9.31
	pointsOfFunctionPlot[931].Y = 1.204

	pointsOfFunctionPlot[932].X = 9.32
	pointsOfFunctionPlot[932].Y = 1.204

	pointsOfFunctionPlot[933].X = 9.33
	pointsOfFunctionPlot[933].Y = 1.204

	pointsOfFunctionPlot[934].X = 9.34
	pointsOfFunctionPlot[934].Y = 1.204

	pointsOfFunctionPlot[935].X = 9.35
	pointsOfFunctionPlot[935].Y = 1.204

	pointsOfFunctionPlot[936].X = 9.36
	pointsOfFunctionPlot[936].Y = 1.204

	pointsOfFunctionPlot[937].X = 9.37
	pointsOfFunctionPlot[937].Y = 1.204

	pointsOfFunctionPlot[938].X = 9.38
	pointsOfFunctionPlot[938].Y = 1.205

	pointsOfFunctionPlot[939].X = 9.39
	pointsOfFunctionPlot[939].Y = 1.205

	pointsOfFunctionPlot[940].X = 9.40
	pointsOfFunctionPlot[940].Y = 1.205

	pointsOfFunctionPlot[941].X = 9.41
	pointsOfFunctionPlot[941].Y = 1.205

	pointsOfFunctionPlot[942].X = 9.42
	pointsOfFunctionPlot[942].Y = 1.205

	pointsOfFunctionPlot[943].X = 9.43
	pointsOfFunctionPlot[943].Y = 1.205

	pointsOfFunctionPlot[944].X = 9.44
	pointsOfFunctionPlot[944].Y = 1.205

	pointsOfFunctionPlot[945].X = 9.45
	pointsOfFunctionPlot[945].Y = 1.205

	pointsOfFunctionPlot[946].X = 9.46
	pointsOfFunctionPlot[946].Y = 1.205

	pointsOfFunctionPlot[947].X = 9.47
	pointsOfFunctionPlot[947].Y = 1.206

	pointsOfFunctionPlot[948].X = 9.48
	pointsOfFunctionPlot[948].Y = 1.206

	pointsOfFunctionPlot[949].X = 9.49
	pointsOfFunctionPlot[949].Y = 1.206

	pointsOfFunctionPlot[950].X = 9.50
	pointsOfFunctionPlot[950].Y = 1.206

	pointsOfFunctionPlot[951].X = 9.51
	pointsOfFunctionPlot[951].Y = 1.206

	pointsOfFunctionPlot[952].X = 9.52
	pointsOfFunctionPlot[952].Y = 1.206

	pointsOfFunctionPlot[953].X = 9.53
	pointsOfFunctionPlot[953].Y = 1.206

	pointsOfFunctionPlot[954].X = 9.54
	pointsOfFunctionPlot[954].Y = 1.206

	pointsOfFunctionPlot[955].X = 9.55
	pointsOfFunctionPlot[955].Y = 1.206

	pointsOfFunctionPlot[956].X = 9.56
	pointsOfFunctionPlot[956].Y = 1.206

	pointsOfFunctionPlot[957].X = 9.57
	pointsOfFunctionPlot[957].Y = 1.207

	pointsOfFunctionPlot[958].X = 9.58
	pointsOfFunctionPlot[958].Y = 1.207

	pointsOfFunctionPlot[959].X = 9.59
	pointsOfFunctionPlot[959].Y = 1.207

	pointsOfFunctionPlot[960].X = 9.60
	pointsOfFunctionPlot[960].Y = 1.207

	pointsOfFunctionPlot[961].X = 9.61
	pointsOfFunctionPlot[961].Y = 1.207

	pointsOfFunctionPlot[962].X = 9.62
	pointsOfFunctionPlot[962].Y = 1.207

	pointsOfFunctionPlot[963].X = 9.63
	pointsOfFunctionPlot[963].Y = 1.207

	pointsOfFunctionPlot[964].X = 9.64
	pointsOfFunctionPlot[964].Y = 1.207

	pointsOfFunctionPlot[965].X = 9.65
	pointsOfFunctionPlot[965].Y = 1.207

	pointsOfFunctionPlot[966].X = 9.66
	pointsOfFunctionPlot[966].Y = 1.208

	pointsOfFunctionPlot[967].X = 9.67
	pointsOfFunctionPlot[967].Y = 1.208

	pointsOfFunctionPlot[968].X = 9.68
	pointsOfFunctionPlot[968].Y = 1.208

	pointsOfFunctionPlot[969].X = 9.69
	pointsOfFunctionPlot[969].Y = 1.208

	pointsOfFunctionPlot[970].X = 9.70
	pointsOfFunctionPlot[970].Y = 1.208

	pointsOfFunctionPlot[971].X = 9.71
	pointsOfFunctionPlot[971].Y = 1.208

	pointsOfFunctionPlot[972].X = 9.72
	pointsOfFunctionPlot[972].Y = 1.208

	pointsOfFunctionPlot[973].X = 9.73
	pointsOfFunctionPlot[973].Y = 1.208

	pointsOfFunctionPlot[974].X = 9.74
	pointsOfFunctionPlot[974].Y = 1.208

	pointsOfFunctionPlot[975].X = 9.75
	pointsOfFunctionPlot[975].Y = 1.208

	pointsOfFunctionPlot[976].X = 9.76
	pointsOfFunctionPlot[976].Y = 1.208

	pointsOfFunctionPlot[977].X = 9.77
	pointsOfFunctionPlot[977].Y = 1.209

	pointsOfFunctionPlot[978].X = 9.78
	pointsOfFunctionPlot[978].Y = 1.209

	pointsOfFunctionPlot[979].X = 9.79
	pointsOfFunctionPlot[979].Y = 1.209

	pointsOfFunctionPlot[980].X = 9.80
	pointsOfFunctionPlot[980].Y = 1.209

	pointsOfFunctionPlot[981].X = 9.81
	pointsOfFunctionPlot[981].Y = 1.209

	pointsOfFunctionPlot[982].X = 9.82
	pointsOfFunctionPlot[982].Y = 1.209

	pointsOfFunctionPlot[983].X = 9.83
	pointsOfFunctionPlot[983].Y = 1.209

	pointsOfFunctionPlot[984].X = 9.84
	pointsOfFunctionPlot[984].Y = 1.209

	pointsOfFunctionPlot[985].X = 9.85
	pointsOfFunctionPlot[985].Y = 1.21

	pointsOfFunctionPlot[986].X = 9.86
	pointsOfFunctionPlot[986].Y = 1.21

	pointsOfFunctionPlot[987].X = 9.87
	pointsOfFunctionPlot[987].Y = 1.21

	pointsOfFunctionPlot[988].X = 9.88
	pointsOfFunctionPlot[988].Y = 1.21

	pointsOfFunctionPlot[989].X = 9.89
	pointsOfFunctionPlot[989].Y = 1.21

	pointsOfFunctionPlot[990].X = 9.90
	pointsOfFunctionPlot[990].Y = 1.21

	pointsOfFunctionPlot[991].X = 9.91
	pointsOfFunctionPlot[991].Y = 1.21

	pointsOfFunctionPlot[992].X = 9.92
	pointsOfFunctionPlot[992].Y = 1.21

	pointsOfFunctionPlot[993].X = 9.93
	pointsOfFunctionPlot[993].Y = 1.21

	pointsOfFunctionPlot[994].X = 9.94
	pointsOfFunctionPlot[994].Y = 1.21

	pointsOfFunctionPlot[995].X = 9.95
	pointsOfFunctionPlot[995].Y = 1.211

	pointsOfFunctionPlot[996].X = 9.96
	pointsOfFunctionPlot[996].Y = 1.211

	pointsOfFunctionPlot[997].X = 9.97
	pointsOfFunctionPlot[997].Y = 1.211

	pointsOfFunctionPlot[998].X = 9.98
	pointsOfFunctionPlot[998].Y = 1.211

	pointsOfFunctionPlot[999].X = 9.99
	pointsOfFunctionPlot[999].Y = 1.211

	pointsOfFunctionPlot[1_000].X = 10.0
	pointsOfFunctionPlot[1_000].Y = 1.211










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function f(x) = x^(1/12)"

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
		"x-to-power-1-over-12-function-plot-01.png"); err != nil {

		panic(err)
	}
}
