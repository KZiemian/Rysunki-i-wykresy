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

	pointsOfFunctionPlot := make(plotter.XYs, 501)

	pointsOfFunctionPlot[0].X = -2.50
	pointsOfFunctionPlot[0].Y = -0.999

	pointsOfFunctionPlot[1].X = -2.49
	pointsOfFunctionPlot[1].Y = -0.999

	pointsOfFunctionPlot[2].X = -2.48
	pointsOfFunctionPlot[2].Y = -0.999

	pointsOfFunctionPlot[3].X = -2.47
	pointsOfFunctionPlot[3].Y = -0.999

	pointsOfFunctionPlot[4].X = -2.46
	pointsOfFunctionPlot[4].Y = -0.999

	pointsOfFunctionPlot[5].X = -2.45
	pointsOfFunctionPlot[5].Y = -0.999

	pointsOfFunctionPlot[6].X = -2.44
	pointsOfFunctionPlot[6].Y = -0.999

	pointsOfFunctionPlot[7].X = -2.43
	pointsOfFunctionPlot[7].Y = -0.999

	pointsOfFunctionPlot[8].X = -2.42
	pointsOfFunctionPlot[8].Y = -0.999

	pointsOfFunctionPlot[9].X = -2.41
	pointsOfFunctionPlot[9].Y = -0.999

	pointsOfFunctionPlot[10].X = -2.40
	pointsOfFunctionPlot[10].Y = -0.999

	pointsOfFunctionPlot[11].X = -2.39
	pointsOfFunctionPlot[11].Y = -0.999

	pointsOfFunctionPlot[12].X = -2.38
	pointsOfFunctionPlot[12].Y = -0.999

	pointsOfFunctionPlot[13].X = -2.37
	pointsOfFunctionPlot[13].Y = -0.999

	pointsOfFunctionPlot[14].X = -2.36
	pointsOfFunctionPlot[14].Y = -0.999

	pointsOfFunctionPlot[15].X = -2.35
	pointsOfFunctionPlot[15].Y = -0.999

	pointsOfFunctionPlot[16].X = -2.34
	pointsOfFunctionPlot[16].Y = -0.999

	pointsOfFunctionPlot[17].X = -2.33
	pointsOfFunctionPlot[17].Y = -0.999

	pointsOfFunctionPlot[18].X = -2.32
	pointsOfFunctionPlot[18].Y = -0.998

	pointsOfFunctionPlot[19].X = -2.31
	pointsOfFunctionPlot[19].Y = -0.998

	pointsOfFunctionPlot[20].X = -2.30
	pointsOfFunctionPlot[20].Y = -0.998

	pointsOfFunctionPlot[21].X = -2.29
	pointsOfFunctionPlot[21].Y = -0.998

	pointsOfFunctionPlot[22].X = -2.28
	pointsOfFunctionPlot[22].Y = -0.998

	pointsOfFunctionPlot[23].X = -2.27
	pointsOfFunctionPlot[23].Y = -0.998

	pointsOfFunctionPlot[24].X = -2.26
	pointsOfFunctionPlot[24].Y = -0.998

	pointsOfFunctionPlot[25].X = -2.25
	pointsOfFunctionPlot[25].Y = -0.998

	pointsOfFunctionPlot[26].X = -2.24
	pointsOfFunctionPlot[26].Y = -0.998

	pointsOfFunctionPlot[27].X = -2.23
	pointsOfFunctionPlot[27].Y = -0.998

	pointsOfFunctionPlot[28].X = -2.22
	pointsOfFunctionPlot[28].Y = -0.998

	pointsOfFunctionPlot[29].X = -2.21
	pointsOfFunctionPlot[29].Y = -0.998

	pointsOfFunctionPlot[30].X = -2.20
	pointsOfFunctionPlot[30].Y = -0.998

	pointsOfFunctionPlot[31].X = -2.19
	pointsOfFunctionPlot[31].Y = -0.998

	pointsOfFunctionPlot[32].X = -2.18
	pointsOfFunctionPlot[32].Y = -0.997

	pointsOfFunctionPlot[33].X = -2.17
	pointsOfFunctionPlot[33].Y = -0.997

	pointsOfFunctionPlot[34].X = -2.16
	pointsOfFunctionPlot[34].Y = -0.997

	pointsOfFunctionPlot[35].X = -2.15
	pointsOfFunctionPlot[35].Y = -0.997

	pointsOfFunctionPlot[36].X = -2.14
	pointsOfFunctionPlot[36].Y = -0.997

	pointsOfFunctionPlot[37].X = -2.13
	pointsOfFunctionPlot[37].Y = -0.997

	pointsOfFunctionPlot[38].X = -2.12
	pointsOfFunctionPlot[38].Y = -0.997

	pointsOfFunctionPlot[39].X = -2.11
	pointsOfFunctionPlot[39].Y = -0.997

	pointsOfFunctionPlot[40].X = -2.10
	pointsOfFunctionPlot[40].Y = -0.997

	pointsOfFunctionPlot[41].X = -2.09
	pointsOfFunctionPlot[41].Y = -0.996

	pointsOfFunctionPlot[42].X = -2.08
	pointsOfFunctionPlot[42].Y = -0.996

	pointsOfFunctionPlot[43].X = -2.07
	pointsOfFunctionPlot[43].Y = -0.996

	pointsOfFunctionPlot[44].X = -2.06
	pointsOfFunctionPlot[44].Y = -0.996

	pointsOfFunctionPlot[45].X = -2.05
	pointsOfFunctionPlot[45].Y = -0.996

	pointsOfFunctionPlot[46].X = -2.04
	pointsOfFunctionPlot[46].Y = -0.996

	pointsOfFunctionPlot[47].X = -2.03
	pointsOfFunctionPlot[47].Y = -0.995

	pointsOfFunctionPlot[48].X = -2.02
	pointsOfFunctionPlot[48].Y = -0.995

	pointsOfFunctionPlot[49].X = -2.01
	pointsOfFunctionPlot[49].Y = -0.995

	pointsOfFunctionPlot[50].X = -2.0
	pointsOfFunctionPlot[50].Y = -0.995

	pointsOfFunctionPlot[51].X = -1.99
	pointsOfFunctionPlot[51].Y = -0.995

	pointsOfFunctionPlot[52].X = -1.98
	pointsOfFunctionPlot[52].Y = -0.994

	pointsOfFunctionPlot[53].X = -1.97
	pointsOfFunctionPlot[53].Y = -0.994

	pointsOfFunctionPlot[54].X = -1.96
	pointsOfFunctionPlot[54].Y = -0.994

	pointsOfFunctionPlot[55].X = -1.95
	pointsOfFunctionPlot[55].Y = -0.994

	pointsOfFunctionPlot[56].X = -1.94
	pointsOfFunctionPlot[56].Y = -0.993

	pointsOfFunctionPlot[57].X = -1.93
	pointsOfFunctionPlot[57].Y = -0.993

	pointsOfFunctionPlot[58].X = -1.92
	pointsOfFunctionPlot[58].Y = -0.993

	pointsOfFunctionPlot[59].X = -1.91
	pointsOfFunctionPlot[59].Y = -0.993

	pointsOfFunctionPlot[60].X = -1.90
	pointsOfFunctionPlot[60].Y = -0.992

	pointsOfFunctionPlot[61].X = -1.89
	pointsOfFunctionPlot[61].Y = -0.992

	pointsOfFunctionPlot[62].X = -1.88
	pointsOfFunctionPlot[62].Y = -0.992

	pointsOfFunctionPlot[63].X = -1.87
	pointsOfFunctionPlot[63].Y = -0.991

	pointsOfFunctionPlot[64].X = -1.86
	pointsOfFunctionPlot[64].Y = -0.991

	pointsOfFunctionPlot[65].X = -1.85
	pointsOfFunctionPlot[65].Y = -0.991

	pointsOfFunctionPlot[66].X = -1.84
	pointsOfFunctionPlot[66].Y = -0.99

	pointsOfFunctionPlot[67].X = -1.83
	pointsOfFunctionPlot[67].Y = -0.99

	pointsOfFunctionPlot[68].X = -1.82
	pointsOfFunctionPlot[68].Y = -0.989

	pointsOfFunctionPlot[69].X = -1.81
	pointsOfFunctionPlot[69].Y = -0.989

	pointsOfFunctionPlot[70].X = -1.80
	pointsOfFunctionPlot[70].Y = -0.989

	pointsOfFunctionPlot[71].X = -1.79
	pointsOfFunctionPlot[71].Y = -0.988

	pointsOfFunctionPlot[72].X = -1.78
	pointsOfFunctionPlot[72].Y = -0.988

	pointsOfFunctionPlot[73].X = -1.77
	pointsOfFunctionPlot[73].Y = -0.987

	pointsOfFunctionPlot[74].X = -1.76
	pointsOfFunctionPlot[74].Y = -0.987

	pointsOfFunctionPlot[75].X = -1.75
	pointsOfFunctionPlot[75].Y = -0.986

	pointsOfFunctionPlot[76].X = -1.74
	pointsOfFunctionPlot[76].Y = -0.986

	pointsOfFunctionPlot[77].X = -1.73
	pointsOfFunctionPlot[77].Y = -0.985

	pointsOfFunctionPlot[78].X = -1.72
	pointsOfFunctionPlot[78].Y = -0.985

	pointsOfFunctionPlot[79].X = -1.71
	pointsOfFunctionPlot[79].Y = -0.984

	pointsOfFunctionPlot[80].X = -1.7
	pointsOfFunctionPlot[80].Y = -0.983

	pointsOfFunctionPlot[81].X = -1.69
	pointsOfFunctionPlot[81].Y = -0.983

	pointsOfFunctionPlot[82].X = -1.68
	pointsOfFunctionPlot[82].Y = -0.982

	pointsOfFunctionPlot[83].X = -1.67
	pointsOfFunctionPlot[83].Y = -0.981

	pointsOfFunctionPlot[84].X = -1.66
	pointsOfFunctionPlot[84].Y = -0.981

	pointsOfFunctionPlot[85].X = -1.65
	pointsOfFunctionPlot[85].Y = -0.98

	pointsOfFunctionPlot[86].X = -1.64
	pointsOfFunctionPlot[86].Y = -0.979

	pointsOfFunctionPlot[87].X = -1.63
	pointsOfFunctionPlot[87].Y = -0.978

	pointsOfFunctionPlot[88].X = -1.62
	pointsOfFunctionPlot[88].Y = -0.978

	pointsOfFunctionPlot[89].X = -1.61
	pointsOfFunctionPlot[89].Y = -0.977

	pointsOfFunctionPlot[90].X = -1.60
	pointsOfFunctionPlot[90].Y = -0.976

	pointsOfFunctionPlot[91].X = -1.59
	pointsOfFunctionPlot[91].Y = -0.975

	pointsOfFunctionPlot[92].X = -1.58
	pointsOfFunctionPlot[92].Y = -0.974

	pointsOfFunctionPlot[93].X = -1.57
	pointsOfFunctionPlot[93].Y = -0.973

	pointsOfFunctionPlot[94].X = -1.56
	pointsOfFunctionPlot[94].Y = -0.972

	pointsOfFunctionPlot[95].X = -1.55
	pointsOfFunctionPlot[95].Y = -0.971

	pointsOfFunctionPlot[96].X = -1.54
	pointsOfFunctionPlot[96].Y = -0.97

	pointsOfFunctionPlot[97].X = -1.53
	pointsOfFunctionPlot[97].Y = -0.969

	pointsOfFunctionPlot[98].X = -1.52
	pointsOfFunctionPlot[98].Y = -0.968

	pointsOfFunctionPlot[99].X = -1.51
	pointsOfFunctionPlot[99].Y = -0.967

	pointsOfFunctionPlot[100].X = -1.50
	pointsOfFunctionPlot[100].Y = -0.966

	pointsOfFunctionPlot[101].X = -1.49
	pointsOfFunctionPlot[101].Y = -0.964

	pointsOfFunctionPlot[102].X = -1.48
	pointsOfFunctionPlot[102].Y = -0.963

	pointsOfFunctionPlot[103].X = -1.47
	pointsOfFunctionPlot[103].Y = -0.962

	pointsOfFunctionPlot[104].X = -1.46
	pointsOfFunctionPlot[104].Y = -0.961

	pointsOfFunctionPlot[105].X = -1.45
	pointsOfFunctionPlot[105].Y = -0.959

	pointsOfFunctionPlot[106].X = -1.44
	pointsOfFunctionPlot[106].Y = -0.958

	pointsOfFunctionPlot[107].X = -1.43
	pointsOfFunctionPlot[107].Y = -0.956

	pointsOfFunctionPlot[108].X = -1.42
	pointsOfFunctionPlot[108].Y = -0.955

	pointsOfFunctionPlot[109].X = -1.41
	pointsOfFunctionPlot[109].Y = -0.953

	pointsOfFunctionPlot[110].X = -1.40
	pointsOfFunctionPlot[110].Y = -0.952

	pointsOfFunctionPlot[111].X = -1.39
	pointsOfFunctionPlot[111].Y = -0.95

	pointsOfFunctionPlot[112].X = -1.38
	pointsOfFunctionPlot[112].Y = -0.949

	pointsOfFunctionPlot[113].X = -1.37
	pointsOfFunctionPlot[113].Y = -0.947

	pointsOfFunctionPlot[114].X = -1.36
	pointsOfFunctionPlot[114].Y = -0.945

	pointsOfFunctionPlot[115].X = -1.35
	pointsOfFunctionPlot[115].Y = -0.943

	pointsOfFunctionPlot[116].X = -1.34
	pointsOfFunctionPlot[116].Y = -0.941

	pointsOfFunctionPlot[117].X = -1.33
	pointsOfFunctionPlot[117].Y = -0.94

	pointsOfFunctionPlot[118].X = -1.32
	pointsOfFunctionPlot[118].Y = -0.938

	pointsOfFunctionPlot[119].X = -1.31
	pointsOfFunctionPlot[119].Y = -0.936

	pointsOfFunctionPlot[120].X = -1.30
	pointsOfFunctionPlot[120].Y = -0.934

	pointsOfFunctionPlot[121].X = -1.29
	pointsOfFunctionPlot[121].Y = -0.931

	pointsOfFunctionPlot[122].X = -1.28
	pointsOfFunctionPlot[122].Y = -0.929

	pointsOfFunctionPlot[123].X = -1.27
	pointsOfFunctionPlot[123].Y = -0.927

	pointsOfFunctionPlot[124].X = -1.26
	pointsOfFunctionPlot[124].Y = -0.925

	pointsOfFunctionPlot[125].X = -1.25
	pointsOfFunctionPlot[125].Y = -0.922

	pointsOfFunctionPlot[126].X = -1.24
	pointsOfFunctionPlot[126].Y = -0.92

	pointsOfFunctionPlot[127].X = -1.23
	pointsOfFunctionPlot[127].Y = -0.918

	pointsOfFunctionPlot[128].X = -1.22
	pointsOfFunctionPlot[128].Y = -0.915

	pointsOfFunctionPlot[129].X = -1.21
	pointsOfFunctionPlot[129].Y = -0.912

	pointsOfFunctionPlot[130].X = -1.20
	pointsOfFunctionPlot[130].Y = -0.91

	pointsOfFunctionPlot[131].X = -1.19
	pointsOfFunctionPlot[131].Y = -0.907

	pointsOfFunctionPlot[132].X = -1.18
	pointsOfFunctionPlot[132].Y = -0.904

	pointsOfFunctionPlot[133].X = -1.17
	pointsOfFunctionPlot[133].Y = -0.902

	pointsOfFunctionPlot[134].X = -1.16
	pointsOfFunctionPlot[134].Y = -0.899

	pointsOfFunctionPlot[135].X = -1.15
	pointsOfFunctionPlot[135].Y = -0.896

	pointsOfFunctionPlot[136].X = -1.14
	pointsOfFunctionPlot[136].Y = -0.893

	pointsOfFunctionPlot[137].X = -1.13
	pointsOfFunctionPlot[137].Y = -0.889

	pointsOfFunctionPlot[138].X = -1.12
	pointsOfFunctionPlot[138].Y = -0.886

	pointsOfFunctionPlot[139].X = -1.11
	pointsOfFunctionPlot[139].Y = -0.883

	pointsOfFunctionPlot[140].X = -1.10
	pointsOfFunctionPlot[140].Y = -0.88

	pointsOfFunctionPlot[141].X = -1.09
	pointsOfFunctionPlot[141].Y = -0.876

	pointsOfFunctionPlot[142].X = -1.08
	pointsOfFunctionPlot[142].Y = -0.873

	pointsOfFunctionPlot[143].X = -1.07
	pointsOfFunctionPlot[143].Y = -0.869

	pointsOfFunctionPlot[144].X = -1.06
	pointsOfFunctionPlot[144].Y = -0.866

	pointsOfFunctionPlot[145].X = -1.05
	pointsOfFunctionPlot[145].Y = -0.862

	pointsOfFunctionPlot[146].X = -1.04
	pointsOfFunctionPlot[146].Y = -0.858

	pointsOfFunctionPlot[147].X = -1.03
	pointsOfFunctionPlot[147].Y = -0.854

	pointsOfFunctionPlot[148].X = -1.02
	pointsOfFunctionPlot[148].Y = -0.85

	pointsOfFunctionPlot[149].X = -1.01
	pointsOfFunctionPlot[149].Y = -0.846

	pointsOfFunctionPlot[150].X = -1.0
	pointsOfFunctionPlot[150].Y = -0.842

	pointsOfFunctionPlot[151].X = -0.99
	pointsOfFunctionPlot[151].Y = -0.838

	pointsOfFunctionPlot[152].X = -0.98
	pointsOfFunctionPlot[152].Y = -0.834

	pointsOfFunctionPlot[153].X = -0.97
	pointsOfFunctionPlot[153].Y = -0.829

	pointsOfFunctionPlot[154].X = -0.96
	pointsOfFunctionPlot[154].Y = -0.825

	pointsOfFunctionPlot[155].X = -0.95
	pointsOfFunctionPlot[155].Y = -0.82

	pointsOfFunctionPlot[156].X = -0.94
	pointsOfFunctionPlot[156].Y = -0.816

	pointsOfFunctionPlot[157].X = -0.93
	pointsOfFunctionPlot[157].Y = -0.811

	pointsOfFunctionPlot[158].X = -0.92
	pointsOfFunctionPlot[158].Y = -0.806

	pointsOfFunctionPlot[159].X = -0.91
	pointsOfFunctionPlot[159].Y = -0.801

	pointsOfFunctionPlot[160].X = -0.90
	pointsOfFunctionPlot[160].Y = -0.796

	pointsOfFunctionPlot[161].X = -0.89
	pointsOfFunctionPlot[161].Y = -0.791

	pointsOfFunctionPlot[162].X = -0.88
	pointsOfFunctionPlot[162].Y = -0.786

	pointsOfFunctionPlot[163].X = -0.87
	pointsOfFunctionPlot[163].Y = -0.781

	pointsOfFunctionPlot[164].X = -0.86
	pointsOfFunctionPlot[164].Y = -0.776

	pointsOfFunctionPlot[165].X = -0.85
	pointsOfFunctionPlot[165].Y = -0.77

	pointsOfFunctionPlot[166].X = -0.84
	pointsOfFunctionPlot[166].Y = -0.765

	pointsOfFunctionPlot[167].X = -0.83
	pointsOfFunctionPlot[167].Y = -0.759

	pointsOfFunctionPlot[168].X = -0.82
	pointsOfFunctionPlot[168].Y = -0.753

	pointsOfFunctionPlot[169].X = -0.81
	pointsOfFunctionPlot[169].Y = -0.748

	pointsOfFunctionPlot[170].X = -0.80
	pointsOfFunctionPlot[170].Y = -0.742

	pointsOfFunctionPlot[171].X = -0.79
	pointsOfFunctionPlot[171].Y = -0.736

	pointsOfFunctionPlot[172].X = -0.78
	pointsOfFunctionPlot[172].Y = -0.73

	pointsOfFunctionPlot[173].X = -0.77
	pointsOfFunctionPlot[173].Y = -0.723

	pointsOfFunctionPlot[174].X = -0.76
	pointsOfFunctionPlot[174].Y = -0.717

	pointsOfFunctionPlot[175].X = -0.75
	pointsOfFunctionPlot[175].Y = -0.711

	pointsOfFunctionPlot[176].X = -0.74
	pointsOfFunctionPlot[176].Y = -0.704

	pointsOfFunctionPlot[177].X = -0.73
	pointsOfFunctionPlot[177].Y = -0.698

	pointsOfFunctionPlot[178].X = -0.72
	pointsOfFunctionPlot[178].Y = -0.691

	pointsOfFunctionPlot[179].X = -0.71
	pointsOfFunctionPlot[179].Y = -0.684

	pointsOfFunctionPlot[180].X = -0.7
	pointsOfFunctionPlot[180].Y = -0.677

	pointsOfFunctionPlot[181].X = -0.69
	pointsOfFunctionPlot[181].Y = -0.67

	pointsOfFunctionPlot[182].X = -0.68
	pointsOfFunctionPlot[182].Y = -0.663

	pointsOfFunctionPlot[183].X = -0.67
	pointsOfFunctionPlot[183].Y = -0.656

	pointsOfFunctionPlot[184].X = -0.66
	pointsOfFunctionPlot[184].Y = -0.649

	pointsOfFunctionPlot[185].X = -0.65
	pointsOfFunctionPlot[185].Y = -0.642

	pointsOfFunctionPlot[186].X = -0.64
	pointsOfFunctionPlot[186].Y = -0.634

	pointsOfFunctionPlot[187].X = -0.63
	pointsOfFunctionPlot[187].Y = -0.627

	pointsOfFunctionPlot[188].X = -0.62
	pointsOfFunctionPlot[188].Y = -0.619

	pointsOfFunctionPlot[189].X = -0.61
	pointsOfFunctionPlot[189].Y = -0.611

	pointsOfFunctionPlot[190].X = -0.60
	pointsOfFunctionPlot[190].Y = -0.603

	pointsOfFunctionPlot[191].X = -0.59
	pointsOfFunctionPlot[191].Y = -0.595

	pointsOfFunctionPlot[192].X = -0.58
	pointsOfFunctionPlot[192].Y = -0.587

	pointsOfFunctionPlot[193].X = -0.57
	pointsOfFunctionPlot[193].Y = -0.579

	pointsOfFunctionPlot[194].X = -0.56
	pointsOfFunctionPlot[194].Y = -0.571

	pointsOfFunctionPlot[195].X = -0.55
	pointsOfFunctionPlot[195].Y = -0.563

	pointsOfFunctionPlot[196].X = -0.54
	pointsOfFunctionPlot[196].Y = -0.554

	pointsOfFunctionPlot[197].X = -0.53
	pointsOfFunctionPlot[197].Y = -0.546

	pointsOfFunctionPlot[198].X = -0.52
	pointsOfFunctionPlot[198].Y = -0.537

	pointsOfFunctionPlot[199].X = -0.51
	pointsOfFunctionPlot[199].Y = -0.529

	pointsOfFunctionPlot[200].X = -0.50
	pointsOfFunctionPlot[200].Y = -0.52

	pointsOfFunctionPlot[201].X = -0.49
	pointsOfFunctionPlot[201].Y = -0.511

	pointsOfFunctionPlot[202].X = -0.48
	pointsOfFunctionPlot[202].Y = -0.502

	pointsOfFunctionPlot[203].X = -0.47
	pointsOfFunctionPlot[203].Y = -0.493

	pointsOfFunctionPlot[204].X = -0.46
	pointsOfFunctionPlot[204].Y = -0.484

	pointsOfFunctionPlot[205].X = -0.45
	pointsOfFunctionPlot[205].Y = -0.475

	pointsOfFunctionPlot[206].X = -0.44
	pointsOfFunctionPlot[206].Y = -0.466

	pointsOfFunctionPlot[207].X = -0.43
	pointsOfFunctionPlot[207].Y = -0.456

	pointsOfFunctionPlot[208].X = -0.42
	pointsOfFunctionPlot[208].Y = -0.447

	pointsOfFunctionPlot[209].X = -0.41
	pointsOfFunctionPlot[209].Y = -0.437

	pointsOfFunctionPlot[210].X = -0.40
	pointsOfFunctionPlot[210].Y = -0.428

	pointsOfFunctionPlot[211].X = -0.39
	pointsOfFunctionPlot[211].Y = -0.418

	pointsOfFunctionPlot[212].X = -0.38
	pointsOfFunctionPlot[212].Y = -0.409

	pointsOfFunctionPlot[213].X = -0.37
	pointsOfFunctionPlot[213].Y = -0.399

	pointsOfFunctionPlot[214].X = -0.36
	pointsOfFunctionPlot[214].Y = -0.389

	pointsOfFunctionPlot[215].X = -0.35
	pointsOfFunctionPlot[215].Y = -0.379

	pointsOfFunctionPlot[216].X = -0.34
	pointsOfFunctionPlot[216].Y = -0.369

	pointsOfFunctionPlot[217].X = -0.33
	pointsOfFunctionPlot[217].Y = -0.359

	pointsOfFunctionPlot[218].X = -0.32
	pointsOfFunctionPlot[218].Y = -0.349

	pointsOfFunctionPlot[219].X = -0.31
	pointsOfFunctionPlot[219].Y = -0.338

	pointsOfFunctionPlot[220].X = -0.30
	pointsOfFunctionPlot[220].Y = -0.328

	pointsOfFunctionPlot[221].X = -0.29
	pointsOfFunctionPlot[221].Y = -0.318

	pointsOfFunctionPlot[222].X = -0.28
	pointsOfFunctionPlot[222].Y = -0.307

	pointsOfFunctionPlot[223].X = -0.27
	pointsOfFunctionPlot[223].Y = -0.297

	pointsOfFunctionPlot[224].X = -0.26
	pointsOfFunctionPlot[224].Y = -0.286

	pointsOfFunctionPlot[225].X = -0.25
	pointsOfFunctionPlot[225].Y = -0.276

	pointsOfFunctionPlot[226].X = -0.24
	pointsOfFunctionPlot[226].Y = -0.265

	pointsOfFunctionPlot[227].X = -0.23
	pointsOfFunctionPlot[227].Y = -0.255

	pointsOfFunctionPlot[228].X = -0.22
	pointsOfFunctionPlot[228].Y = -0.244

	pointsOfFunctionPlot[229].X = -0.21
	pointsOfFunctionPlot[229].Y = -0.233

	pointsOfFunctionPlot[230].X = -0.20
	pointsOfFunctionPlot[230].Y = -0.222

	pointsOfFunctionPlot[231].X = -0.19
	pointsOfFunctionPlot[231].Y = -0.211

	pointsOfFunctionPlot[232].X = -0.18
	pointsOfFunctionPlot[232].Y = -0.2

	pointsOfFunctionPlot[233].X = -0.17
	pointsOfFunctionPlot[233].Y = -0.189

	pointsOfFunctionPlot[234].X = -0.16
	pointsOfFunctionPlot[234].Y = -0.179

	pointsOfFunctionPlot[235].X = -0.15
	pointsOfFunctionPlot[235].Y = -0.167

	pointsOfFunctionPlot[236].X = -0.14
	pointsOfFunctionPlot[236].Y = -0.156

	pointsOfFunctionPlot[237].X = -0.13
	pointsOfFunctionPlot[237].Y = -0.145

	pointsOfFunctionPlot[238].X = -0.12
	pointsOfFunctionPlot[238].Y = -0.134

	pointsOfFunctionPlot[239].X = -0.11
	pointsOfFunctionPlot[239].Y = -0.123

	pointsOfFunctionPlot[240].X = -0.10
	pointsOfFunctionPlot[240].Y = -0.112

	pointsOfFunctionPlot[241].X = -0.09
	pointsOfFunctionPlot[241].Y = -0.101

	pointsOfFunctionPlot[242].X = -0.08
	pointsOfFunctionPlot[242].Y = -0.09

	pointsOfFunctionPlot[243].X = -0.07
	pointsOfFunctionPlot[243].Y = -0.078

	pointsOfFunctionPlot[244].X = -0.06
	pointsOfFunctionPlot[244].Y = -0.067

	pointsOfFunctionPlot[245].X = -0.05
	pointsOfFunctionPlot[245].Y = -0.056

	pointsOfFunctionPlot[246].X = -0.04
	pointsOfFunctionPlot[246].Y = -0.045

	pointsOfFunctionPlot[247].X = -0.03
	pointsOfFunctionPlot[247].Y = -0.033

	pointsOfFunctionPlot[248].X = -0.02
	pointsOfFunctionPlot[248].Y = -0.022

	pointsOfFunctionPlot[249].X = -0.01
	pointsOfFunctionPlot[249].Y = -0.011

	pointsOfFunctionPlot[250].X = 0.0
	pointsOfFunctionPlot[250].Y = 0.0

	pointsOfFunctionPlot[251].X = 0.01
	pointsOfFunctionPlot[251].Y = 0.011

	pointsOfFunctionPlot[252].X = 0.02
	pointsOfFunctionPlot[252].Y = 0.022

	pointsOfFunctionPlot[253].X = 0.03
	pointsOfFunctionPlot[253].Y = 0.033

	pointsOfFunctionPlot[254].X = 0.04
	pointsOfFunctionPlot[254].Y = 0.045

	pointsOfFunctionPlot[255].X = 0.05
	pointsOfFunctionPlot[255].Y = 0.056

	pointsOfFunctionPlot[256].X = 0.06
	pointsOfFunctionPlot[256].Y = 0.067

	pointsOfFunctionPlot[257].X = 0.07
	pointsOfFunctionPlot[257].Y = 0.078

	pointsOfFunctionPlot[258].X = 0.08
	pointsOfFunctionPlot[258].Y = 0.09

	pointsOfFunctionPlot[259].X = 0.09
	pointsOfFunctionPlot[259].Y = 0.101

	pointsOfFunctionPlot[260].X = 0.10
	pointsOfFunctionPlot[260].Y = 0.112

	pointsOfFunctionPlot[261].X = 0.11
	pointsOfFunctionPlot[261].Y = 0.123

	pointsOfFunctionPlot[262].X = 0.12
	pointsOfFunctionPlot[262].Y = 0.134

	pointsOfFunctionPlot[263].X = 0.13
	pointsOfFunctionPlot[263].Y = 0.145

	pointsOfFunctionPlot[264].X = 0.14
	pointsOfFunctionPlot[264].Y = 0.156

	pointsOfFunctionPlot[265].X = 0.15
	pointsOfFunctionPlot[265].Y = 0.167

	pointsOfFunctionPlot[266].X = 0.16
	pointsOfFunctionPlot[266].Y = 0.179

	pointsOfFunctionPlot[267].X = 0.17
	pointsOfFunctionPlot[267].Y = 0.189

	pointsOfFunctionPlot[268].X = 0.18
	pointsOfFunctionPlot[268].Y = 0.2

	pointsOfFunctionPlot[269].X = 0.19
	pointsOfFunctionPlot[269].Y = 0.211

	pointsOfFunctionPlot[270].X = 0.20
	pointsOfFunctionPlot[270].Y = 0.222

	pointsOfFunctionPlot[271].X = 0.21
	pointsOfFunctionPlot[271].Y = 0.233

	pointsOfFunctionPlot[272].X = 0.22
	pointsOfFunctionPlot[272].Y = 0.244

	pointsOfFunctionPlot[273].X = 0.23
	pointsOfFunctionPlot[273].Y = 0.255

	pointsOfFunctionPlot[274].X = 0.24
	pointsOfFunctionPlot[274].Y = 0.265

	pointsOfFunctionPlot[275].X = 0.25
	pointsOfFunctionPlot[275].Y = 0.276

	pointsOfFunctionPlot[276].X = 0.26
	pointsOfFunctionPlot[276].Y = 0.286

	pointsOfFunctionPlot[277].X = 0.27
	pointsOfFunctionPlot[277].Y = 0.297

	pointsOfFunctionPlot[278].X = 0.28
	pointsOfFunctionPlot[278].Y = 0.307

	pointsOfFunctionPlot[279].X = 0.29
	pointsOfFunctionPlot[279].Y = 0.318

	pointsOfFunctionPlot[280].X = 0.30
	pointsOfFunctionPlot[280].Y = 0.328

	pointsOfFunctionPlot[281].X = 0.31
	pointsOfFunctionPlot[281].Y = 0.338

	pointsOfFunctionPlot[282].X = 0.32
	pointsOfFunctionPlot[282].Y = 0.349

	pointsOfFunctionPlot[283].X = 0.33
	pointsOfFunctionPlot[283].Y = 0.359

	pointsOfFunctionPlot[284].X = 0.34
	pointsOfFunctionPlot[284].Y = 0.369

	pointsOfFunctionPlot[285].X = 0.35
	pointsOfFunctionPlot[285].Y = 0.379

	pointsOfFunctionPlot[286].X = 0.36
	pointsOfFunctionPlot[286].Y = 0.389

	pointsOfFunctionPlot[287].X = 0.37
	pointsOfFunctionPlot[287].Y = 0.399

	pointsOfFunctionPlot[288].X = 0.38
	pointsOfFunctionPlot[288].Y = 0.409

	pointsOfFunctionPlot[289].X = 0.39
	pointsOfFunctionPlot[289].Y = 0.418

	pointsOfFunctionPlot[290].X = 0.40
	pointsOfFunctionPlot[290].Y = 0.428

	pointsOfFunctionPlot[291].X = 0.41
	pointsOfFunctionPlot[291].Y = 0.437

	pointsOfFunctionPlot[292].X = 0.42
	pointsOfFunctionPlot[292].Y = 0.447

	pointsOfFunctionPlot[293].X = 0.43
	pointsOfFunctionPlot[293].Y = 0.456

	pointsOfFunctionPlot[294].X = 0.44
	pointsOfFunctionPlot[294].Y = 0.466

	pointsOfFunctionPlot[295].X = 0.45
	pointsOfFunctionPlot[295].Y = 0.475

	pointsOfFunctionPlot[296].X = 0.46
	pointsOfFunctionPlot[296].Y = 0.484

	pointsOfFunctionPlot[297].X = 0.47
	pointsOfFunctionPlot[297].Y = 0.493

	pointsOfFunctionPlot[298].X = 0.48
	pointsOfFunctionPlot[298].Y = 0.502

	pointsOfFunctionPlot[299].X = 0.49
	pointsOfFunctionPlot[299].Y = 0.511

	pointsOfFunctionPlot[300].X = 0.50
	pointsOfFunctionPlot[300].Y = 0.52

	pointsOfFunctionPlot[301].X = 0.51
	pointsOfFunctionPlot[301].Y = 0.529

	pointsOfFunctionPlot[302].X = 0.52
	pointsOfFunctionPlot[302].Y = 0.537

	pointsOfFunctionPlot[303].X = 0.53
	pointsOfFunctionPlot[303].Y = 0.546

	pointsOfFunctionPlot[304].X = 0.54
	pointsOfFunctionPlot[304].Y = 0.554

	pointsOfFunctionPlot[305].X = 0.55
	pointsOfFunctionPlot[305].Y = 0.563

	pointsOfFunctionPlot[306].X = 0.56
	pointsOfFunctionPlot[306].Y = 0.571

	pointsOfFunctionPlot[307].X = 0.57
	pointsOfFunctionPlot[307].Y = 0.579

	pointsOfFunctionPlot[308].X = 0.58
	pointsOfFunctionPlot[308].Y = 0.587

	pointsOfFunctionPlot[309].X = 0.59
	pointsOfFunctionPlot[309].Y = 0.595

	pointsOfFunctionPlot[310].X = 0.60
	pointsOfFunctionPlot[310].Y = 0.603

	pointsOfFunctionPlot[311].X = 0.61
	pointsOfFunctionPlot[311].Y = 0.611

	pointsOfFunctionPlot[312].X = 0.62
	pointsOfFunctionPlot[312].Y = 0.619

	pointsOfFunctionPlot[313].X = 0.63
	pointsOfFunctionPlot[313].Y = 0.627

	pointsOfFunctionPlot[314].X = 0.64
	pointsOfFunctionPlot[314].Y = 0.634

	pointsOfFunctionPlot[315].X = 0.65
	pointsOfFunctionPlot[315].Y = 0.642

	pointsOfFunctionPlot[316].X = 0.66
	pointsOfFunctionPlot[316].Y = 0.649

	pointsOfFunctionPlot[317].X = 0.67
	pointsOfFunctionPlot[317].Y = 0.656

	pointsOfFunctionPlot[318].X = 0.68
	pointsOfFunctionPlot[318].Y = 0.663

	pointsOfFunctionPlot[319].X = 0.69
	pointsOfFunctionPlot[319].Y = 0.67

	pointsOfFunctionPlot[320].X = 0.70
	pointsOfFunctionPlot[320].Y = 0.677

	pointsOfFunctionPlot[321].X = 0.71
	pointsOfFunctionPlot[321].Y = 0.684

	pointsOfFunctionPlot[322].X = 0.72
	pointsOfFunctionPlot[322].Y = 0.691

	pointsOfFunctionPlot[323].X = 0.73
	pointsOfFunctionPlot[323].Y = 0.698

	pointsOfFunctionPlot[324].X = 0.74
	pointsOfFunctionPlot[324].Y = 0.704

	pointsOfFunctionPlot[325].X = 0.75
	pointsOfFunctionPlot[325].Y = 0.711

	pointsOfFunctionPlot[326].X = 0.76
	pointsOfFunctionPlot[326].Y = 0.717

	pointsOfFunctionPlot[327].X = 0.77
	pointsOfFunctionPlot[327].Y = 0.723

	pointsOfFunctionPlot[328].X = 0.78
	pointsOfFunctionPlot[328].Y = 0.73

	pointsOfFunctionPlot[329].X = 0.79
	pointsOfFunctionPlot[329].Y = 0.736

	pointsOfFunctionPlot[330].X = 0.80
	pointsOfFunctionPlot[330].Y = 0.742

	pointsOfFunctionPlot[331].X = 0.81
	pointsOfFunctionPlot[331].Y = 0.748

	pointsOfFunctionPlot[332].X = 0.82
	pointsOfFunctionPlot[332].Y = 0.753

	pointsOfFunctionPlot[333].X = 0.83
	pointsOfFunctionPlot[333].Y = 0.759

	pointsOfFunctionPlot[334].X = 0.84
	pointsOfFunctionPlot[334].Y = 0.765

	pointsOfFunctionPlot[335].X = 0.85
	pointsOfFunctionPlot[335].Y = 0.77

	pointsOfFunctionPlot[336].X = 0.86
	pointsOfFunctionPlot[336].Y = 0.776

	pointsOfFunctionPlot[337].X = 0.87
	pointsOfFunctionPlot[337].Y = 0.781

	pointsOfFunctionPlot[338].X = 0.88
	pointsOfFunctionPlot[338].Y = 0.786

	pointsOfFunctionPlot[339].X = 0.89
	pointsOfFunctionPlot[339].Y = 0.791

	pointsOfFunctionPlot[340].X = 0.90
	pointsOfFunctionPlot[340].Y = 0.796

	pointsOfFunctionPlot[341].X = 0.91
	pointsOfFunctionPlot[341].Y = 0.801

	pointsOfFunctionPlot[342].X = 0.92
	pointsOfFunctionPlot[342].Y = 0.806

	pointsOfFunctionPlot[343].X = 0.93
	pointsOfFunctionPlot[343].Y = 0.811

	pointsOfFunctionPlot[344].X = 0.94
	pointsOfFunctionPlot[344].Y = 0.816

	pointsOfFunctionPlot[345].X = 0.95
	pointsOfFunctionPlot[345].Y = 0.82

	pointsOfFunctionPlot[346].X = 0.96
	pointsOfFunctionPlot[346].Y = 0.825

	pointsOfFunctionPlot[347].X = 0.97
	pointsOfFunctionPlot[347].Y = 0.829

	pointsOfFunctionPlot[348].X = 0.98
	pointsOfFunctionPlot[348].Y = 0.834

	pointsOfFunctionPlot[349].X = 0.99
	pointsOfFunctionPlot[349].Y = 0.838

	pointsOfFunctionPlot[350].X = 1.0
	pointsOfFunctionPlot[350].Y = 0.842

	pointsOfFunctionPlot[351].X = 1.01
	pointsOfFunctionPlot[351].Y = 0.846

	pointsOfFunctionPlot[352].X = 1.02
	pointsOfFunctionPlot[352].Y = 0.85

	pointsOfFunctionPlot[353].X = 1.03
	pointsOfFunctionPlot[353].Y = 0.854

	pointsOfFunctionPlot[354].X = 1.04
	pointsOfFunctionPlot[354].Y = 0.858

	pointsOfFunctionPlot[355].X = 1.05
	pointsOfFunctionPlot[355].Y = 0.862

	pointsOfFunctionPlot[356].X = 1.06
	pointsOfFunctionPlot[356].Y = 0.866

	pointsOfFunctionPlot[357].X = 1.07
	pointsOfFunctionPlot[357].Y = 0.869

	pointsOfFunctionPlot[358].X = 1.08
	pointsOfFunctionPlot[358].Y = 0.873

	pointsOfFunctionPlot[359].X = 1.09
	pointsOfFunctionPlot[359].Y = 0.876

	pointsOfFunctionPlot[360].X = 1.10
	pointsOfFunctionPlot[360].Y = 0.88

	pointsOfFunctionPlot[361].X = 1.11
	pointsOfFunctionPlot[361].Y = 0.883

	pointsOfFunctionPlot[362].X = 1.12
	pointsOfFunctionPlot[362].Y = 0.886

	pointsOfFunctionPlot[363].X = 1.13
	pointsOfFunctionPlot[363].Y = 0.889

	pointsOfFunctionPlot[364].X = 1.14
	pointsOfFunctionPlot[364].Y = 0.893

	pointsOfFunctionPlot[365].X = 1.15
	pointsOfFunctionPlot[365].Y = 0.896

	pointsOfFunctionPlot[366].X = 1.16
	pointsOfFunctionPlot[366].Y = 0.899

	pointsOfFunctionPlot[367].X = 1.17
	pointsOfFunctionPlot[367].Y = 0.902

	pointsOfFunctionPlot[368].X = 1.18
	pointsOfFunctionPlot[368].Y = 0.904

	pointsOfFunctionPlot[369].X = 1.19
	pointsOfFunctionPlot[369].Y = 0.907

	pointsOfFunctionPlot[370].X = 1.20
	pointsOfFunctionPlot[370].Y = 0.91

	pointsOfFunctionPlot[371].X = 1.21
	pointsOfFunctionPlot[371].Y = 0.912

	pointsOfFunctionPlot[372].X = 1.22
	pointsOfFunctionPlot[372].Y = 0.915

	pointsOfFunctionPlot[373].X = 1.23
	pointsOfFunctionPlot[373].Y = 0.918

	pointsOfFunctionPlot[374].X = 1.24
	pointsOfFunctionPlot[374].Y = 0.92

	pointsOfFunctionPlot[375].X = 1.25
	pointsOfFunctionPlot[375].Y = 0.922

	pointsOfFunctionPlot[376].X = 1.26
	pointsOfFunctionPlot[376].Y = 0.925

	pointsOfFunctionPlot[377].X = 1.27
	pointsOfFunctionPlot[377].Y = 0.927

	pointsOfFunctionPlot[378].X = 1.28
	pointsOfFunctionPlot[378].Y = 0.929

	pointsOfFunctionPlot[379].X = 1.29
	pointsOfFunctionPlot[379].Y = 0.931

	pointsOfFunctionPlot[380].X = 1.30
	pointsOfFunctionPlot[380].Y = 0.934

	pointsOfFunctionPlot[381].X = 1.31
	pointsOfFunctionPlot[381].Y = 0.936

	pointsOfFunctionPlot[382].X = 1.32
	pointsOfFunctionPlot[382].Y = 0.938

	pointsOfFunctionPlot[383].X = 1.33
	pointsOfFunctionPlot[383].Y = 0.94

	pointsOfFunctionPlot[384].X = 1.34
	pointsOfFunctionPlot[384].Y = 0.941

	pointsOfFunctionPlot[385].X = 1.35
	pointsOfFunctionPlot[385].Y = 0.943

	pointsOfFunctionPlot[386].X = 1.36
	pointsOfFunctionPlot[386].Y = 0.945

	pointsOfFunctionPlot[387].X = 1.37
	pointsOfFunctionPlot[387].Y = 0.947

	pointsOfFunctionPlot[388].X = 1.38
	pointsOfFunctionPlot[388].Y = 0.949

	pointsOfFunctionPlot[389].X = 1.39
	pointsOfFunctionPlot[389].Y = 0.95

	pointsOfFunctionPlot[390].X = 1.40
	pointsOfFunctionPlot[390].Y = 0.952

	pointsOfFunctionPlot[391].X = 1.41
	pointsOfFunctionPlot[391].Y = 0.953

	pointsOfFunctionPlot[392].X = 1.42
	pointsOfFunctionPlot[392].Y = 0.955

	pointsOfFunctionPlot[393].X = 1.43
	pointsOfFunctionPlot[393].Y = 0.956

	pointsOfFunctionPlot[394].X = 1.44
	pointsOfFunctionPlot[394].Y = 0.958

	pointsOfFunctionPlot[395].X = 1.45
	pointsOfFunctionPlot[395].Y = 0.959

	pointsOfFunctionPlot[396].X = 1.46
	pointsOfFunctionPlot[396].Y = 0.961

	pointsOfFunctionPlot[397].X = 1.47
	pointsOfFunctionPlot[397].Y = 0.962

	pointsOfFunctionPlot[398].X = 1.48
	pointsOfFunctionPlot[398].Y = 0.963

	pointsOfFunctionPlot[399].X = 1.49
	pointsOfFunctionPlot[399].Y = 0.964

	pointsOfFunctionPlot[400].X = 1.50
	pointsOfFunctionPlot[400].Y = 0.966

	pointsOfFunctionPlot[401].X = 1.51
	pointsOfFunctionPlot[401].Y = 0.967

	pointsOfFunctionPlot[402].X = 1.52
	pointsOfFunctionPlot[402].Y = 0.968

	pointsOfFunctionPlot[403].X = 1.53
	pointsOfFunctionPlot[403].Y = 0.969

	pointsOfFunctionPlot[404].X = 1.54
	pointsOfFunctionPlot[404].Y = 0.97

	pointsOfFunctionPlot[405].X = 1.55
	pointsOfFunctionPlot[405].Y = 0.971

	pointsOfFunctionPlot[406].X = 1.56
	pointsOfFunctionPlot[406].Y = 0.972

	pointsOfFunctionPlot[407].X = 1.57
	pointsOfFunctionPlot[407].Y = 0.973

	pointsOfFunctionPlot[408].X = 1.58
	pointsOfFunctionPlot[408].Y = 0.974

	pointsOfFunctionPlot[409].X = 1.59
	pointsOfFunctionPlot[409].Y = 0.975

	pointsOfFunctionPlot[410].X = 1.60
	pointsOfFunctionPlot[410].Y = 0.976

	pointsOfFunctionPlot[411].X = 1.61
	pointsOfFunctionPlot[411].Y = 0.977

	pointsOfFunctionPlot[412].X = 1.62
	pointsOfFunctionPlot[412].Y = 0.978

	pointsOfFunctionPlot[413].X = 1.63
	pointsOfFunctionPlot[413].Y = 0.978

	pointsOfFunctionPlot[414].X = 1.64
	pointsOfFunctionPlot[414].Y = 0.979

	pointsOfFunctionPlot[415].X = 1.65
	pointsOfFunctionPlot[415].Y = 0.98

	pointsOfFunctionPlot[416].X = 1.66
	pointsOfFunctionPlot[416].Y = 0.981

	pointsOfFunctionPlot[417].X = 1.67
	pointsOfFunctionPlot[417].Y = 0.981

	pointsOfFunctionPlot[418].X = 1.68
	pointsOfFunctionPlot[418].Y = 0.982

	pointsOfFunctionPlot[419].X = 1.69
	pointsOfFunctionPlot[419].Y = 0.983

	pointsOfFunctionPlot[420].X = 1.70
	pointsOfFunctionPlot[420].Y = 0.983

	pointsOfFunctionPlot[421].X = 1.71
	pointsOfFunctionPlot[421].Y = 0.984

	pointsOfFunctionPlot[422].X = 1.72
	pointsOfFunctionPlot[422].Y = 0.985

	pointsOfFunctionPlot[423].X = 1.73
	pointsOfFunctionPlot[423].Y = 0.985

	pointsOfFunctionPlot[424].X = 1.74
	pointsOfFunctionPlot[424].Y = 0.986

	pointsOfFunctionPlot[425].X = 1.75
	pointsOfFunctionPlot[425].Y = 0.986

	pointsOfFunctionPlot[426].X = 1.76
	pointsOfFunctionPlot[426].Y = 0.987

	pointsOfFunctionPlot[427].X = 1.77
	pointsOfFunctionPlot[427].Y = 0.987

	pointsOfFunctionPlot[428].X = 1.78
	pointsOfFunctionPlot[428].Y = 0.988

	pointsOfFunctionPlot[429].X = 1.79
	pointsOfFunctionPlot[429].Y = 0.988

	pointsOfFunctionPlot[430].X = 1.80
	pointsOfFunctionPlot[430].Y = 0.989

	pointsOfFunctionPlot[431].X = 1.81
	pointsOfFunctionPlot[431].Y = 0.989

	pointsOfFunctionPlot[432].X = 1.82
	pointsOfFunctionPlot[432].Y = 0.989

	pointsOfFunctionPlot[433].X = 1.83
	pointsOfFunctionPlot[433].Y = 0.99

	pointsOfFunctionPlot[434].X = 1.84
	pointsOfFunctionPlot[434].Y = 0.99

	pointsOfFunctionPlot[435].X = 1.85
	pointsOfFunctionPlot[435].Y = 0.991

	pointsOfFunctionPlot[436].X = 1.86
	pointsOfFunctionPlot[436].Y = 0.991

	pointsOfFunctionPlot[437].X = 1.87
	pointsOfFunctionPlot[437].Y = 0.991

	pointsOfFunctionPlot[438].X = 1.88
	pointsOfFunctionPlot[438].Y = 0.992

	pointsOfFunctionPlot[439].X = 1.89
	pointsOfFunctionPlot[439].Y = 0.992

	pointsOfFunctionPlot[440].X = 1.90
	pointsOfFunctionPlot[440].Y = 0.992

	pointsOfFunctionPlot[441].X = 1.91
	pointsOfFunctionPlot[441].Y = 0.993

	pointsOfFunctionPlot[442].X = 1.92
	pointsOfFunctionPlot[442].Y = 0.993

	pointsOfFunctionPlot[443].X = 1.93
	pointsOfFunctionPlot[443].Y = 0.993

	pointsOfFunctionPlot[444].X = 1.94
	pointsOfFunctionPlot[444].Y = 0.993

	pointsOfFunctionPlot[445].X = 1.95
	pointsOfFunctionPlot[445].Y = 0.994

	pointsOfFunctionPlot[446].X = 1.96
	pointsOfFunctionPlot[446].Y = 0.994

	pointsOfFunctionPlot[447].X = 1.97
	pointsOfFunctionPlot[447].Y = 0.994

	pointsOfFunctionPlot[448].X = 1.98
	pointsOfFunctionPlot[448].Y = 0.994

	pointsOfFunctionPlot[449].X = 1.99
	pointsOfFunctionPlot[449].Y = 0.995

	pointsOfFunctionPlot[450].X = 2.0
	pointsOfFunctionPlot[450].Y = 0.995

	pointsOfFunctionPlot[451].X = 2.01
	pointsOfFunctionPlot[451].Y = 0.995

	pointsOfFunctionPlot[452].X = 2.02
	pointsOfFunctionPlot[452].Y = 0.995

	pointsOfFunctionPlot[453].X = 2.03
	pointsOfFunctionPlot[453].Y = 0.995

	pointsOfFunctionPlot[454].X = 2.04
	pointsOfFunctionPlot[454].Y = 0.996

	pointsOfFunctionPlot[455].X = 2.05
	pointsOfFunctionPlot[455].Y = 0.996

	pointsOfFunctionPlot[456].X = 2.06
	pointsOfFunctionPlot[456].Y = 0.996

	pointsOfFunctionPlot[457].X = 2.07
	pointsOfFunctionPlot[457].Y = 0.996

	pointsOfFunctionPlot[458].X = 2.08
	pointsOfFunctionPlot[458].Y = 0.996

	pointsOfFunctionPlot[459].X = 2.09
	pointsOfFunctionPlot[459].Y = 0.996

	pointsOfFunctionPlot[460].X = 2.10
	pointsOfFunctionPlot[460].Y = 0.997

	pointsOfFunctionPlot[461].X = 2.11
	pointsOfFunctionPlot[461].Y = 0.997

	pointsOfFunctionPlot[462].X = 2.12
	pointsOfFunctionPlot[462].Y = 0.997

	pointsOfFunctionPlot[463].X = 2.13
	pointsOfFunctionPlot[463].Y = 0.997

	pointsOfFunctionPlot[464].X = 2.14
	pointsOfFunctionPlot[464].Y = 0.997

	pointsOfFunctionPlot[465].X = 2.15
	pointsOfFunctionPlot[465].Y = 0.997

	pointsOfFunctionPlot[466].X = 2.16
	pointsOfFunctionPlot[466].Y = 0.997

	pointsOfFunctionPlot[467].X = 2.17
	pointsOfFunctionPlot[467].Y = 0.997

	pointsOfFunctionPlot[468].X = 2.18
	pointsOfFunctionPlot[468].Y = 0.997

	pointsOfFunctionPlot[469].X = 2.19
	pointsOfFunctionPlot[469].Y = 0.998

	pointsOfFunctionPlot[470].X = 2.20
	pointsOfFunctionPlot[470].Y = 0.998

	pointsOfFunctionPlot[471].X = 2.21
	pointsOfFunctionPlot[471].Y = 0.998

	pointsOfFunctionPlot[472].X = 2.22
	pointsOfFunctionPlot[472].Y = 0.998

	pointsOfFunctionPlot[473].X = 2.23
	pointsOfFunctionPlot[473].Y = 0.998

	pointsOfFunctionPlot[474].X = 2.24
	pointsOfFunctionPlot[474].Y = 0.998

	pointsOfFunctionPlot[475].X = 2.25
	pointsOfFunctionPlot[475].Y = 0.998

	pointsOfFunctionPlot[476].X = 2.26
	pointsOfFunctionPlot[476].Y = 0.998

	pointsOfFunctionPlot[477].X = 2.27
	pointsOfFunctionPlot[477].Y = 0.998

	pointsOfFunctionPlot[478].X = 2.28
	pointsOfFunctionPlot[478].Y = 0.998

	pointsOfFunctionPlot[479].X = 2.29
	pointsOfFunctionPlot[479].Y = 0.998

	pointsOfFunctionPlot[480].X = 2.30
	pointsOfFunctionPlot[480].Y = 0.998

	pointsOfFunctionPlot[481].X = 2.31
	pointsOfFunctionPlot[481].Y = 0.998

	pointsOfFunctionPlot[482].X = 2.32
	pointsOfFunctionPlot[482].Y = 0.998

	pointsOfFunctionPlot[483].X = 2.33
	pointsOfFunctionPlot[483].Y = 0.999

	pointsOfFunctionPlot[484].X = 2.34
	pointsOfFunctionPlot[484].Y = 0.999

	pointsOfFunctionPlot[485].X = 2.35
	pointsOfFunctionPlot[485].Y = 0.999

	pointsOfFunctionPlot[486].X = 2.36
	pointsOfFunctionPlot[486].Y = 0.999

	pointsOfFunctionPlot[487].X = 2.37
	pointsOfFunctionPlot[487].Y = 0.999

	pointsOfFunctionPlot[488].X = 2.38
	pointsOfFunctionPlot[488].Y = 0.999

	pointsOfFunctionPlot[489].X = 2.39
	pointsOfFunctionPlot[489].Y = 0.999

	pointsOfFunctionPlot[490].X = 2.40
	pointsOfFunctionPlot[490].Y = 0.999

	pointsOfFunctionPlot[491].X = 2.41
	pointsOfFunctionPlot[491].Y = 0.999

	pointsOfFunctionPlot[492].X = 2.42
	pointsOfFunctionPlot[492].Y = 0.999

	pointsOfFunctionPlot[493].X = 2.43
	pointsOfFunctionPlot[493].Y = 0.999

	pointsOfFunctionPlot[494].X = 2.44
	pointsOfFunctionPlot[494].Y = 0.999

	pointsOfFunctionPlot[495].X = 2.45
	pointsOfFunctionPlot[495].Y = 0.999

	pointsOfFunctionPlot[496].X = 2.46
	pointsOfFunctionPlot[496].Y = 0.999

	pointsOfFunctionPlot[497].X = 2.47
	pointsOfFunctionPlot[497].Y = 0.999

	pointsOfFunctionPlot[498].X = 2.48
	pointsOfFunctionPlot[498].Y = 0.999

	pointsOfFunctionPlot[499].X = 2.49
	pointsOfFunctionPlot[499].Y = 0.999

	pointsOfFunctionPlot[500].X = 2.50
	pointsOfFunctionPlot[500].Y = 0.999










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
		"erf-function-plot-02.png"); err != nil {

		panic(err)
	}
}
