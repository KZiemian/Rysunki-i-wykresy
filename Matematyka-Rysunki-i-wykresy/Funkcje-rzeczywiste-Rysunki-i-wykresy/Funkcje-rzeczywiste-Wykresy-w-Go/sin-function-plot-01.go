package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function sin(x).

	pointsOfFunctionPlot := make(plotter.XYs, 631)

	pointsOfFunctionPlot[0].X = 0.0
	pointsOfFunctionPlot[0].Y = 0.0

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = 0.01

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = 0.02

	pointsOfFunctionPlot[3].X = 0.03
	pointsOfFunctionPlot[3].Y = 0.03

	pointsOfFunctionPlot[4].X = 0.04
	pointsOfFunctionPlot[4].Y = 0.04

	pointsOfFunctionPlot[5].X = 0.05
	pointsOfFunctionPlot[5].Y = 0.05

	pointsOfFunctionPlot[6].X = 0.06
	pointsOfFunctionPlot[6].Y = 0.06

	pointsOfFunctionPlot[7].X = 0.07
	pointsOfFunctionPlot[7].Y = 0.069

	pointsOfFunctionPlot[8].X = 0.08
	pointsOfFunctionPlot[8].Y = 0.079

	pointsOfFunctionPlot[9].X = 0.09
	pointsOfFunctionPlot[9].Y = 0.089

	pointsOfFunctionPlot[10].X = 0.10
	pointsOfFunctionPlot[10].Y = 0.099

	pointsOfFunctionPlot[11].X = 0.11
	pointsOfFunctionPlot[11].Y = 0.109

	pointsOfFunctionPlot[12].X = 0.12
	pointsOfFunctionPlot[12].Y = 0.119

	pointsOfFunctionPlot[13].X = 0.13
	pointsOfFunctionPlot[13].Y = 0.129

	pointsOfFunctionPlot[14].X = 0.14
	pointsOfFunctionPlot[14].Y = 0.139

	pointsOfFunctionPlot[15].X = 0.15
	pointsOfFunctionPlot[15].Y = 0.149

	pointsOfFunctionPlot[16].X = 0.16
	pointsOfFunctionPlot[16].Y = 0.159

	pointsOfFunctionPlot[17].X = 0.17
	pointsOfFunctionPlot[17].Y = 0.169

	pointsOfFunctionPlot[18].X = 0.18
	pointsOfFunctionPlot[18].Y = 0.179

	pointsOfFunctionPlot[19].X = 0.19
	pointsOfFunctionPlot[19].Y = 0.188

	pointsOfFunctionPlot[20].X = 0.20
	pointsOfFunctionPlot[20].Y = 0.198

	pointsOfFunctionPlot[21].X = 0.21
	pointsOfFunctionPlot[21].Y = 0.208

	pointsOfFunctionPlot[22].X = 0.22
	pointsOfFunctionPlot[22].Y = 0.218

	pointsOfFunctionPlot[23].X = 0.23
	pointsOfFunctionPlot[23].Y = 0.228

	pointsOfFunctionPlot[24].X = 0.24
	pointsOfFunctionPlot[24].Y = 0.237

	pointsOfFunctionPlot[25].X = 0.25
	pointsOfFunctionPlot[25].Y = 0.247

	pointsOfFunctionPlot[26].X = 0.26
	pointsOfFunctionPlot[26].Y = 0.257

	pointsOfFunctionPlot[27].X = 0.27
	pointsOfFunctionPlot[27].Y = 0.266

	pointsOfFunctionPlot[28].X = 0.28
	pointsOfFunctionPlot[28].Y = 0.276

	pointsOfFunctionPlot[29].X = 0.29
	pointsOfFunctionPlot[29].Y = 0.286

	pointsOfFunctionPlot[30].X = 0.30
	pointsOfFunctionPlot[30].Y = 0.295

	pointsOfFunctionPlot[31].X = 0.31
	pointsOfFunctionPlot[31].Y = 0.305

	pointsOfFunctionPlot[32].X = 0.32
	pointsOfFunctionPlot[32].Y = 0.314

	pointsOfFunctionPlot[33].X = 0.33
	pointsOfFunctionPlot[33].Y = 0.324

	pointsOfFunctionPlot[34].X = 0.34
	pointsOfFunctionPlot[34].Y = 0.333

	pointsOfFunctionPlot[35].X = 0.35
	pointsOfFunctionPlot[35].Y = 0.342

	pointsOfFunctionPlot[36].X = 0.36
	pointsOfFunctionPlot[36].Y = 0.352

	pointsOfFunctionPlot[37].X = 0.37
	pointsOfFunctionPlot[37].Y = 0.361

	pointsOfFunctionPlot[38].X = 0.38
	pointsOfFunctionPlot[38].Y = 0.37

	pointsOfFunctionPlot[39].X = 0.39
	pointsOfFunctionPlot[39].Y = 0.38

	pointsOfFunctionPlot[40].X = 0.40
	pointsOfFunctionPlot[40].Y = 0.389

	pointsOfFunctionPlot[41].X = 0.41
	pointsOfFunctionPlot[41].Y = 0.398

	pointsOfFunctionPlot[42].X = 0.42
	pointsOfFunctionPlot[42].Y = 0.407

	pointsOfFunctionPlot[43].X = 0.43
	pointsOfFunctionPlot[43].Y = 0.416

	pointsOfFunctionPlot[44].X = 0.44
	pointsOfFunctionPlot[44].Y = 0.425

	pointsOfFunctionPlot[45].X = 0.45
	pointsOfFunctionPlot[45].Y = 0.435

	pointsOfFunctionPlot[46].X = 0.46
	pointsOfFunctionPlot[46].Y = 0.443

	pointsOfFunctionPlot[47].X = 0.47
	pointsOfFunctionPlot[47].Y = 0.452

	pointsOfFunctionPlot[48].X = 0.48
	pointsOfFunctionPlot[48].Y = 0.461

	pointsOfFunctionPlot[49].X = 0.49
	pointsOfFunctionPlot[49].Y = 0.47

	pointsOfFunctionPlot[50].X = 0.50
	pointsOfFunctionPlot[50].Y = 0.479

	pointsOfFunctionPlot[51].X = 0.51
	pointsOfFunctionPlot[51].Y = 0.488

	pointsOfFunctionPlot[52].X = 0.52
	pointsOfFunctionPlot[52].Y = 0.496

	pointsOfFunctionPlot[53].X = 0.53
	pointsOfFunctionPlot[53].Y = 0.505

	pointsOfFunctionPlot[54].X = 0.54
	pointsOfFunctionPlot[54].Y = 0.514

	pointsOfFunctionPlot[55].X = 0.55
	pointsOfFunctionPlot[55].Y = 0.522

	pointsOfFunctionPlot[56].X = 0.56
	pointsOfFunctionPlot[56].Y = 0.531

	pointsOfFunctionPlot[57].X = 0.57
	pointsOfFunctionPlot[57].Y = 0.539

	pointsOfFunctionPlot[58].X = 0.58
	pointsOfFunctionPlot[58].Y = 0.548

	pointsOfFunctionPlot[59].X = 0.59
	pointsOfFunctionPlot[59].Y = 0.556

	pointsOfFunctionPlot[60].X = 0.60
	pointsOfFunctionPlot[60].Y = 0.564

	pointsOfFunctionPlot[61].X = 0.61
	pointsOfFunctionPlot[61].Y = 0.572

	pointsOfFunctionPlot[62].X = 0.62
	pointsOfFunctionPlot[62].Y = 0.581

	pointsOfFunctionPlot[63].X = 0.63
	pointsOfFunctionPlot[63].Y = 0.589

	pointsOfFunctionPlot[64].X = 0.64
	pointsOfFunctionPlot[64].Y = 0.597

	pointsOfFunctionPlot[65].X = 0.65
	pointsOfFunctionPlot[65].Y = 0.605

	pointsOfFunctionPlot[66].X = 0.66
	pointsOfFunctionPlot[66].Y = 0.613

	pointsOfFunctionPlot[67].X = 0.67
	pointsOfFunctionPlot[67].Y = 0.621

	pointsOfFunctionPlot[68].X = 0.68
	pointsOfFunctionPlot[68].Y = 0.628

	pointsOfFunctionPlot[69].X = 0.69
	pointsOfFunctionPlot[69].Y = 0.636

	pointsOfFunctionPlot[70].X = 0.70
	pointsOfFunctionPlot[70].Y = 0.644

	pointsOfFunctionPlot[71].X = 0.71
	pointsOfFunctionPlot[71].Y = 0.651

	pointsOfFunctionPlot[72].X = 0.72
	pointsOfFunctionPlot[72].Y = 0.659

	pointsOfFunctionPlot[73].X = 0.73
	pointsOfFunctionPlot[73].Y = 0.666

	pointsOfFunctionPlot[74].X = 0.74
	pointsOfFunctionPlot[74].Y = 0.674

	pointsOfFunctionPlot[75].X = 0.75
	pointsOfFunctionPlot[75].Y = 0.681

	pointsOfFunctionPlot[76].X = 0.76
	pointsOfFunctionPlot[76].Y = 0.688

	pointsOfFunctionPlot[77].X = 0.77
	pointsOfFunctionPlot[77].Y = 0.696

	pointsOfFunctionPlot[78].X = 0.78
	pointsOfFunctionPlot[78].Y = 0.703

	pointsOfFunctionPlot[79].X = 0.79
	pointsOfFunctionPlot[79].Y = 0.71

	pointsOfFunctionPlot[80].X = 0.80
	pointsOfFunctionPlot[80].Y = 0.717

	pointsOfFunctionPlot[81].X = 0.81
	pointsOfFunctionPlot[81].Y = 0.724

	pointsOfFunctionPlot[82].X = 0.82
	pointsOfFunctionPlot[82].Y = 0.731

	pointsOfFunctionPlot[83].X = 0.83
	pointsOfFunctionPlot[83].Y = 0.737

	pointsOfFunctionPlot[84].X = 0.84
	pointsOfFunctionPlot[84].Y = 0.744

	pointsOfFunctionPlot[85].X = 0.85
	pointsOfFunctionPlot[85].Y = 0.751

	pointsOfFunctionPlot[86].X = 0.86
	pointsOfFunctionPlot[86].Y = 0.757

	pointsOfFunctionPlot[87].X = 0.87
	pointsOfFunctionPlot[87].Y = 0.764

	pointsOfFunctionPlot[88].X = 0.88
	pointsOfFunctionPlot[88].Y = 0.77

	pointsOfFunctionPlot[89].X = 0.89
	pointsOfFunctionPlot[89].Y = 0.777

	pointsOfFunctionPlot[90].X = 0.90
	pointsOfFunctionPlot[90].Y = 0.783

	pointsOfFunctionPlot[91].X = 0.91
	pointsOfFunctionPlot[91].Y = 0.789

	pointsOfFunctionPlot[92].X = 0.92
	pointsOfFunctionPlot[92].Y = 0.795

	pointsOfFunctionPlot[93].X = 0.93
	pointsOfFunctionPlot[93].Y = 0.801

	pointsOfFunctionPlot[94].X = 0.94
	pointsOfFunctionPlot[94].Y = 0.807

	pointsOfFunctionPlot[95].X = 0.95
	pointsOfFunctionPlot[95].Y = 0.813

	pointsOfFunctionPlot[96].X = 0.96
	pointsOfFunctionPlot[96].Y = 0.819

	pointsOfFunctionPlot[97].X = 0.97
	pointsOfFunctionPlot[97].Y = 0.824

	pointsOfFunctionPlot[98].X = 0.98
	pointsOfFunctionPlot[98].Y = 0.83

	pointsOfFunctionPlot[99].X = 0.99
	pointsOfFunctionPlot[99].Y = 0.836

	pointsOfFunctionPlot[100].X = 1.0
	pointsOfFunctionPlot[100].Y = 0.841

	pointsOfFunctionPlot[101].X = 1.01
	pointsOfFunctionPlot[101].Y = 0.846

	pointsOfFunctionPlot[102].X = 1.02
	pointsOfFunctionPlot[102].Y = 0.852

	pointsOfFunctionPlot[103].X = 1.03
	pointsOfFunctionPlot[103].Y = 0.857

	pointsOfFunctionPlot[104].X = 1.04
	pointsOfFunctionPlot[104].Y = 0.862

	pointsOfFunctionPlot[105].X = 1.05
	pointsOfFunctionPlot[105].Y = 0.867

	pointsOfFunctionPlot[106].X = 1.06
	pointsOfFunctionPlot[106].Y = 0.872

	pointsOfFunctionPlot[107].X = 1.07
	pointsOfFunctionPlot[107].Y = 0.877

	pointsOfFunctionPlot[108].X = 1.08
	pointsOfFunctionPlot[108].Y = 0.882

	pointsOfFunctionPlot[109].X = 1.09
	pointsOfFunctionPlot[109].Y = 0.886

	pointsOfFunctionPlot[110].X = 1.10
	pointsOfFunctionPlot[110].Y = 0.891

	pointsOfFunctionPlot[111].X = 1.11
	pointsOfFunctionPlot[111].Y = 0.895

	pointsOfFunctionPlot[112].X = 1.12
	pointsOfFunctionPlot[112].Y = 0.9

	pointsOfFunctionPlot[113].X = 1.13
	pointsOfFunctionPlot[113].Y = 0.904

	pointsOfFunctionPlot[114].X = 1.14
	pointsOfFunctionPlot[114].Y = 0.908

	pointsOfFunctionPlot[115].X = 1.15
	pointsOfFunctionPlot[115].Y = 0.912

	pointsOfFunctionPlot[116].X = 1.16
	pointsOfFunctionPlot[116].Y = 0.916

	pointsOfFunctionPlot[117].X = 1.17
	pointsOfFunctionPlot[117].Y = 0.92

	pointsOfFunctionPlot[118].X = 1.18
	pointsOfFunctionPlot[118].Y = 0.924

	pointsOfFunctionPlot[119].X = 1.19
	pointsOfFunctionPlot[119].Y = 0.928

	pointsOfFunctionPlot[120].X = 1.20
	pointsOfFunctionPlot[120].Y = 0.932

	pointsOfFunctionPlot[121].X = 1.21
	pointsOfFunctionPlot[121].Y = 0.933

	pointsOfFunctionPlot[122].X = 1.22
	pointsOfFunctionPlot[122].Y = 0.939

	pointsOfFunctionPlot[123].X = 1.23
	pointsOfFunctionPlot[123].Y = 0.942

	pointsOfFunctionPlot[124].X = 1.24
	pointsOfFunctionPlot[124].Y = 0.945

	pointsOfFunctionPlot[125].X = 1.25
	pointsOfFunctionPlot[125].Y = 0.949

	pointsOfFunctionPlot[126].X = 1.26
	pointsOfFunctionPlot[126].Y = 0.952

	pointsOfFunctionPlot[127].X = 1.27
	pointsOfFunctionPlot[127].Y = 0.955

	pointsOfFunctionPlot[128].X = 1.28
	pointsOfFunctionPlot[128].Y = 0.958

	pointsOfFunctionPlot[129].X = 1.29
	pointsOfFunctionPlot[129].Y = 0.96

	pointsOfFunctionPlot[130].X = 1.30
	pointsOfFunctionPlot[130].Y = 0.963

	pointsOfFunctionPlot[131].X = 1.31
	pointsOfFunctionPlot[131].Y = 0.966

	pointsOfFunctionPlot[132].X = 1.32
	pointsOfFunctionPlot[132].Y = 0.968

	pointsOfFunctionPlot[133].X = 1.33
	pointsOfFunctionPlot[133].Y = 0.971

	pointsOfFunctionPlot[134].X = 1.34
	pointsOfFunctionPlot[134].Y = 0.973

	pointsOfFunctionPlot[135].X = 1.35
	pointsOfFunctionPlot[135].Y = 0.975

	pointsOfFunctionPlot[136].X = 1.36
	pointsOfFunctionPlot[136].Y = 0.977

	pointsOfFunctionPlot[137].X = 1.37
	pointsOfFunctionPlot[137].Y = 0.979

	pointsOfFunctionPlot[138].X = 1.38
	pointsOfFunctionPlot[138].Y = 0.981

	pointsOfFunctionPlot[139].X = 1.39
	pointsOfFunctionPlot[139].Y = 0.983

	pointsOfFunctionPlot[140].X = 1.40
	pointsOfFunctionPlot[140].Y = 0.985

	pointsOfFunctionPlot[141].X = 1.41
	pointsOfFunctionPlot[141].Y = 0.987

	pointsOfFunctionPlot[142].X = 1.42
	pointsOfFunctionPlot[142].Y = 0.988

	pointsOfFunctionPlot[143].X = 1.43
	pointsOfFunctionPlot[143].Y = 0.99

	pointsOfFunctionPlot[144].X = 1.44
	pointsOfFunctionPlot[144].Y = 0.991

	pointsOfFunctionPlot[145].X = 1.45
	pointsOfFunctionPlot[145].Y = 0.992

	pointsOfFunctionPlot[146].X = 1.46
	pointsOfFunctionPlot[146].Y = 0.993

	pointsOfFunctionPlot[147].X = 1.47
	pointsOfFunctionPlot[147].Y = 0.994

	pointsOfFunctionPlot[148].X = 1.48
	pointsOfFunctionPlot[148].Y = 0.995

	pointsOfFunctionPlot[149].X = 1.49
	pointsOfFunctionPlot[149].Y = 0.996

	pointsOfFunctionPlot[150].X = 1.50
	pointsOfFunctionPlot[150].Y = 0.997

	pointsOfFunctionPlot[151].X = 1.51
	pointsOfFunctionPlot[151].Y = 0.998

	pointsOfFunctionPlot[152].X = 1.52
	pointsOfFunctionPlot[152].Y = 0.998

	pointsOfFunctionPlot[153].X = 1.53
	pointsOfFunctionPlot[153].Y = 0.999

	pointsOfFunctionPlot[154].X = 1.54
	pointsOfFunctionPlot[154].Y = 0.999

	pointsOfFunctionPlot[155].X = 1.55
	pointsOfFunctionPlot[155].Y = 0.999

	pointsOfFunctionPlot[156].X = 1.56
	pointsOfFunctionPlot[156].Y = 0.999

	pointsOfFunctionPlot[157].X = 1.57
	pointsOfFunctionPlot[157].Y = 1.0

	pointsOfFunctionPlot[158].X = 1.58
	pointsOfFunctionPlot[158].Y = 1.0

	pointsOfFunctionPlot[159].X = 1.59
	pointsOfFunctionPlot[159].Y = 0.999

	pointsOfFunctionPlot[160].X = 1.60
	pointsOfFunctionPlot[160].Y = 0.999

	pointsOfFunctionPlot[161].X = 1.61
	pointsOfFunctionPlot[161].Y = 0.999

	pointsOfFunctionPlot[162].X = 1.62
	pointsOfFunctionPlot[162].Y = 0.998

	pointsOfFunctionPlot[163].X = 1.63
	pointsOfFunctionPlot[163].Y = 0.998

	pointsOfFunctionPlot[164].X = 1.64
	pointsOfFunctionPlot[164].Y = 0.997

	pointsOfFunctionPlot[165].X = 1.65
	pointsOfFunctionPlot[165].Y = 0.996

	pointsOfFunctionPlot[166].X = 1.66
	pointsOfFunctionPlot[166].Y = 0.996

	pointsOfFunctionPlot[167].X = 1.67
	pointsOfFunctionPlot[167].Y = 0.995

	pointsOfFunctionPlot[168].X = 1.68
	pointsOfFunctionPlot[168].Y = 0.994

	pointsOfFunctionPlot[169].X = 1.69
	pointsOfFunctionPlot[169].Y = 0.992

	pointsOfFunctionPlot[170].X = 1.70
	pointsOfFunctionPlot[170].Y = 0.991

	pointsOfFunctionPlot[171].X = 1.71
	pointsOfFunctionPlot[171].Y = 0.99

	pointsOfFunctionPlot[172].X = 1.72
	pointsOfFunctionPlot[172].Y = 0.988

	pointsOfFunctionPlot[173].X = 1.73
	pointsOfFunctionPlot[173].Y = 0.987

	pointsOfFunctionPlot[174].X = 1.74
	pointsOfFunctionPlot[174].Y = 0.985

	pointsOfFunctionPlot[175].X = 1.75
	pointsOfFunctionPlot[175].Y = 0.984

	pointsOfFunctionPlot[176].X = 1.76
	pointsOfFunctionPlot[176].Y = 0.982

	pointsOfFunctionPlot[177].X = 1.77
	pointsOfFunctionPlot[177].Y = 0.98

	pointsOfFunctionPlot[178].X = 1.78
	pointsOfFunctionPlot[178].Y = 0.978

	pointsOfFunctionPlot[179].X = 1.79
	pointsOfFunctionPlot[179].Y = 0.976

	pointsOfFunctionPlot[180].X = 1.80
	pointsOfFunctionPlot[180].Y = 0.973

	pointsOfFunctionPlot[181].X = 1.81
	pointsOfFunctionPlot[181].Y = 0.971

	pointsOfFunctionPlot[182].X = 1.82
	pointsOfFunctionPlot[182].Y = 0.969

	pointsOfFunctionPlot[183].X = 1.83
	pointsOfFunctionPlot[183].Y = 0.966

	pointsOfFunctionPlot[184].X = 1.84
	pointsOfFunctionPlot[184].Y = 0.964

	pointsOfFunctionPlot[185].X = 1.85
	pointsOfFunctionPlot[185].Y = 0.961

	pointsOfFunctionPlot[186].X = 1.86
	pointsOfFunctionPlot[186].Y = 0.958

	pointsOfFunctionPlot[187].X = 1.87
	pointsOfFunctionPlot[187].Y = 0.955

	pointsOfFunctionPlot[188].X = 1.88
	pointsOfFunctionPlot[188].Y = 0.952

	pointsOfFunctionPlot[189].X = 1.89
	pointsOfFunctionPlot[189].Y = 0.949

	pointsOfFunctionPlot[190].X = 1.90
	pointsOfFunctionPlot[190].Y = 0.946

	pointsOfFunctionPlot[191].X = 1.91
	pointsOfFunctionPlot[191].Y = 0.943

	pointsOfFunctionPlot[192].X = 1.92
	pointsOfFunctionPlot[192].Y = 0.939

	pointsOfFunctionPlot[193].X = 1.93
	pointsOfFunctionPlot[193].Y = 0.936

	pointsOfFunctionPlot[194].X = 1.94
	pointsOfFunctionPlot[194].Y = 0.932

	pointsOfFunctionPlot[195].X = 1.95
	pointsOfFunctionPlot[195].Y = 0.929

	pointsOfFunctionPlot[196].X = 1.96
	pointsOfFunctionPlot[196].Y = 0.925

	pointsOfFunctionPlot[197].X = 1.97
	pointsOfFunctionPlot[197].Y = 0.921

	pointsOfFunctionPlot[198].X = 1.98
	pointsOfFunctionPlot[198].Y = 0.917

	pointsOfFunctionPlot[199].X = 1.99
	pointsOfFunctionPlot[199].Y = 0.913

	pointsOfFunctionPlot[200].X = 2.0
	pointsOfFunctionPlot[200].Y = 0.909

	pointsOfFunctionPlot[201].X = 2.01
	pointsOfFunctionPlot[201].Y = 0.905

	pointsOfFunctionPlot[202].X = 2.02
	pointsOfFunctionPlot[202].Y = 0.9

	pointsOfFunctionPlot[203].X = 2.03
	pointsOfFunctionPlot[203].Y = 0.896

	pointsOfFunctionPlot[204].X = 2.04
	pointsOfFunctionPlot[204].Y = 0.891

	pointsOfFunctionPlot[205].X = 2.05
	pointsOfFunctionPlot[205].Y = 0.887

	pointsOfFunctionPlot[206].X = 2.06
	pointsOfFunctionPlot[206].Y = 0.882

	pointsOfFunctionPlot[207].X = 2.07
	pointsOfFunctionPlot[207].Y = 0.878

	pointsOfFunctionPlot[208].X = 2.08
	pointsOfFunctionPlot[208].Y = 0.873

	pointsOfFunctionPlot[209].X = 2.09
	pointsOfFunctionPlot[209].Y = 0.868

	pointsOfFunctionPlot[210].X = 2.10
	pointsOfFunctionPlot[210].Y = 0.863

	pointsOfFunctionPlot[211].X = 2.11
	pointsOfFunctionPlot[211].Y = 0.858

	pointsOfFunctionPlot[212].X = 2.12
	pointsOfFunctionPlot[212].Y = 0.852

	pointsOfFunctionPlot[213].X = 2.13
	pointsOfFunctionPlot[213].Y = 0.847

	pointsOfFunctionPlot[214].X = 2.14
	pointsOfFunctionPlot[214].Y = 0.842

	pointsOfFunctionPlot[215].X = 2.15
	pointsOfFunctionPlot[215].Y = 0.836

	pointsOfFunctionPlot[216].X = 2.16
	pointsOfFunctionPlot[216].Y = 0.831

	pointsOfFunctionPlot[217].X = 2.17
	pointsOfFunctionPlot[217].Y = 0.825

	pointsOfFunctionPlot[218].X = 2.18
	pointsOfFunctionPlot[218].Y = 0.82

	pointsOfFunctionPlot[219].X = 2.19
	pointsOfFunctionPlot[219].Y = 0.814

	pointsOfFunctionPlot[220].X = 2.20
	pointsOfFunctionPlot[220].Y = 0.808

	pointsOfFunctionPlot[221].X = 2.21
	pointsOfFunctionPlot[221].Y = 0.802

	pointsOfFunctionPlot[222].X = 2.22
	pointsOfFunctionPlot[222].Y = 0.796

	pointsOfFunctionPlot[223].X = 2.23
	pointsOfFunctionPlot[223].Y = 0.79

	pointsOfFunctionPlot[224].X = 2.24
	pointsOfFunctionPlot[224].Y = 0.784

	pointsOfFunctionPlot[225].X = 2.25
	pointsOfFunctionPlot[225].Y = 0.778

	pointsOfFunctionPlot[226].X = 2.26
	pointsOfFunctionPlot[226].Y = 0.771

	pointsOfFunctionPlot[227].X = 2.27
	pointsOfFunctionPlot[227].Y = 0.765

	pointsOfFunctionPlot[228].X = 2.28
	pointsOfFunctionPlot[228].Y = 0.758

	pointsOfFunctionPlot[229].X = 2.29
	pointsOfFunctionPlot[229].Y = 0.752

	pointsOfFunctionPlot[230].X = 2.30
	pointsOfFunctionPlot[230].Y = 0.745

	pointsOfFunctionPlot[231].X = 2.31
	pointsOfFunctionPlot[231].Y = 0.739

	pointsOfFunctionPlot[232].X = 2.32
	pointsOfFunctionPlot[232].Y = 0.732

	pointsOfFunctionPlot[233].X = 2.33
	pointsOfFunctionPlot[233].Y = 0.725

	pointsOfFunctionPlot[234].X = 2.34
	pointsOfFunctionPlot[234].Y = 0.718

	pointsOfFunctionPlot[235].X = 2.35
	pointsOfFunctionPlot[235].Y = 0.711

	pointsOfFunctionPlot[236].X = 2.36
	pointsOfFunctionPlot[236].Y = 0.704

	pointsOfFunctionPlot[237].X = 2.37
	pointsOfFunctionPlot[237].Y = 0.697

	pointsOfFunctionPlot[238].X = 2.38
	pointsOfFunctionPlot[238].Y = 0.69

	pointsOfFunctionPlot[239].X = 2.39
	pointsOfFunctionPlot[239].Y = 0.682

	pointsOfFunctionPlot[240].X = 2.40
	pointsOfFunctionPlot[240].Y = 0.675

	pointsOfFunctionPlot[241].X = 2.41
	pointsOfFunctionPlot[241].Y = 0.668

	pointsOfFunctionPlot[242].X = 2.42
	pointsOfFunctionPlot[242].Y = 0.66

	pointsOfFunctionPlot[243].X = 2.43
	pointsOfFunctionPlot[243].Y = 0.653

	pointsOfFunctionPlot[244].X = 2.44
	pointsOfFunctionPlot[244].Y = 0.645

	pointsOfFunctionPlot[245].X = 2.45
	pointsOfFunctionPlot[245].Y = 0.637

	pointsOfFunctionPlot[246].X = 2.46
	pointsOfFunctionPlot[246].Y = 0.63

	pointsOfFunctionPlot[247].X = 2.47
	pointsOfFunctionPlot[247].Y = 0.622

	pointsOfFunctionPlot[248].X = 2.48
	pointsOfFunctionPlot[248].Y = 0.614

	pointsOfFunctionPlot[249].X = 2.49
	pointsOfFunctionPlot[249].Y = 0.606

	pointsOfFunctionPlot[250].X = 2.50
	pointsOfFunctionPlot[250].Y = 0.598

	pointsOfFunctionPlot[251].X = 2.51
	pointsOfFunctionPlot[251].Y = 0.59

	pointsOfFunctionPlot[252].X = 2.52
	pointsOfFunctionPlot[252].Y = 0.582

	pointsOfFunctionPlot[253].X = 2.53
	pointsOfFunctionPlot[253].Y = 0.574

	pointsOfFunctionPlot[254].X = 2.54
	pointsOfFunctionPlot[254].Y = 0.566

	pointsOfFunctionPlot[255].X = 2.55
	pointsOfFunctionPlot[255].Y = 0.557

	pointsOfFunctionPlot[256].X = 2.56
	pointsOfFunctionPlot[256].Y = 0.549

	pointsOfFunctionPlot[257].X = 2.57
	pointsOfFunctionPlot[257].Y = 0.541

	pointsOfFunctionPlot[258].X = 2.58
	pointsOfFunctionPlot[258].Y = 0.532

	pointsOfFunctionPlot[259].X = 2.59
	pointsOfFunctionPlot[259].Y = 0.524

	pointsOfFunctionPlot[260].X = 2.60
	pointsOfFunctionPlot[260].Y = 0.515

	pointsOfFunctionPlot[261].X = 2.61
	pointsOfFunctionPlot[261].Y = 0.506

	pointsOfFunctionPlot[262].X = 2.62
	pointsOfFunctionPlot[262].Y = 0.498

	pointsOfFunctionPlot[263].X = 2.63
	pointsOfFunctionPlot[263].Y = 0.489

	pointsOfFunctionPlot[264].X = 2.64
	pointsOfFunctionPlot[264].Y = 0.48

	pointsOfFunctionPlot[265].X = 2.65
	pointsOfFunctionPlot[265].Y = 0.472

	pointsOfFunctionPlot[266].X = 2.66
	pointsOfFunctionPlot[266].Y = 0.463

	pointsOfFunctionPlot[267].X = 2.67
	pointsOfFunctionPlot[267].Y = 0.454

	pointsOfFunctionPlot[268].X = 2.68
	pointsOfFunctionPlot[268].Y = 0.445

	pointsOfFunctionPlot[269].X = 2.69
	pointsOfFunctionPlot[269].Y = 0.436

	pointsOfFunctionPlot[270].X = 2.70
	pointsOfFunctionPlot[270].Y = 0.427

	pointsOfFunctionPlot[271].X = 2.71
	pointsOfFunctionPlot[271].Y = 0.418

	pointsOfFunctionPlot[272].X = 2.72
	pointsOfFunctionPlot[272].Y = 0.409

	pointsOfFunctionPlot[273].X = 2.73
	pointsOfFunctionPlot[273].Y = 0.4

	pointsOfFunctionPlot[274].X = 2.74
	pointsOfFunctionPlot[274].Y = 0.39

	pointsOfFunctionPlot[275].X = 2.75
	pointsOfFunctionPlot[275].Y = 0.381

	pointsOfFunctionPlot[276].X = 2.76
	pointsOfFunctionPlot[276].Y = 0.372

	pointsOfFunctionPlot[277].X = 2.77
	pointsOfFunctionPlot[277].Y = 0.363

	pointsOfFunctionPlot[278].X = 2.78
	pointsOfFunctionPlot[278].Y = 0.353

	pointsOfFunctionPlot[279].X = 2.79
	pointsOfFunctionPlot[279].Y = 0.344

	pointsOfFunctionPlot[280].X = 2.80
	pointsOfFunctionPlot[280].Y = 0.335

	pointsOfFunctionPlot[281].X = 2.81
	pointsOfFunctionPlot[281].Y = 0.325

	pointsOfFunctionPlot[282].X = 2.82
	pointsOfFunctionPlot[282].Y = 0.316

	pointsOfFunctionPlot[283].X = 2.83
	pointsOfFunctionPlot[283].Y = 0.306

	pointsOfFunctionPlot[284].X = 2.84
	pointsOfFunctionPlot[284].Y = 0.297

	pointsOfFunctionPlot[285].X = 2.85
	pointsOfFunctionPlot[285].Y = 0.287

	pointsOfFunctionPlot[286].X = 2.86
	pointsOfFunctionPlot[286].Y = 0.277

	pointsOfFunctionPlot[287].X = 2.87
	pointsOfFunctionPlot[287].Y = 0.268

	pointsOfFunctionPlot[288].X = 2.88
	pointsOfFunctionPlot[288].Y = 0.258

	pointsOfFunctionPlot[289].X = 2.89
	pointsOfFunctionPlot[289].Y = 0.248

	pointsOfFunctionPlot[290].X = 2.90
	pointsOfFunctionPlot[290].Y = 0.239

	pointsOfFunctionPlot[291].X = 2.91
	pointsOfFunctionPlot[291].Y = 0.229

	pointsOfFunctionPlot[292].X = 2.92
	pointsOfFunctionPlot[292].Y = 0.219

	pointsOfFunctionPlot[293].X = 2.93
	pointsOfFunctionPlot[293].Y = 0.21

	pointsOfFunctionPlot[294].X = 2.94
	pointsOfFunctionPlot[294].Y = 0.2

	pointsOfFunctionPlot[295].X = 2.95
	pointsOfFunctionPlot[295].Y = 0.19

	pointsOfFunctionPlot[296].X = 2.96
	pointsOfFunctionPlot[296].Y = 0.18

	pointsOfFunctionPlot[297].X = 2.97
	pointsOfFunctionPlot[297].Y = 0.17

	pointsOfFunctionPlot[298].X = 2.98
	pointsOfFunctionPlot[298].Y = 0.16

	pointsOfFunctionPlot[299].X = 2.99
	pointsOfFunctionPlot[299].Y = 0.151

	pointsOfFunctionPlot[300].X = 3.0
	pointsOfFunctionPlot[300].Y = 0.141

	pointsOfFunctionPlot[301].X = 3.01
	pointsOfFunctionPlot[301].Y = 0.131

	pointsOfFunctionPlot[302].X = 3.02
	pointsOfFunctionPlot[302].Y = 0.121

	pointsOfFunctionPlot[303].X = 3.03
	pointsOfFunctionPlot[303].Y = 0.111

	pointsOfFunctionPlot[304].X = 3.04
	pointsOfFunctionPlot[304].Y = 0.101

	pointsOfFunctionPlot[305].X = 3.05
	pointsOfFunctionPlot[305].Y = 0.091

	pointsOfFunctionPlot[306].X = 3.06
	pointsOfFunctionPlot[306].Y = 0.081

	pointsOfFunctionPlot[307].X = 3.07
	pointsOfFunctionPlot[307].Y = 0.071

	pointsOfFunctionPlot[308].X = 3.08
	pointsOfFunctionPlot[308].Y = 0.061

	pointsOfFunctionPlot[309].X = 3.09
	pointsOfFunctionPlot[309].Y = 0.051

	pointsOfFunctionPlot[310].X = 3.10
	pointsOfFunctionPlot[310].Y = 0.041

	pointsOfFunctionPlot[311].X = 3.11
	pointsOfFunctionPlot[311].Y = 0.031

	pointsOfFunctionPlot[312].X = 3.12
	pointsOfFunctionPlot[312].Y = 0.021

	pointsOfFunctionPlot[313].X = 3.13
	pointsOfFunctionPlot[313].Y = 0.011

	pointsOfFunctionPlot[314].X = 3.14
	pointsOfFunctionPlot[314].Y = 0.001

	pointsOfFunctionPlot[315].X = 3.15
	pointsOfFunctionPlot[315].Y = -0.008

	pointsOfFunctionPlot[316].X = 3.16
	pointsOfFunctionPlot[316].Y = -0.018

	pointsOfFunctionPlot[317].X = 3.17
	pointsOfFunctionPlot[317].Y = -0.028

	pointsOfFunctionPlot[318].X = 3.18
	pointsOfFunctionPlot[318].Y = -0.038

	pointsOfFunctionPlot[319].X = 3.19
	pointsOfFunctionPlot[319].Y = -0.048

	pointsOfFunctionPlot[320].X = 3.20
	pointsOfFunctionPlot[320].Y = -0.058

	pointsOfFunctionPlot[321].X = 3.21
	pointsOfFunctionPlot[321].Y = -0.068

	pointsOfFunctionPlot[322].X = 3.22
	pointsOfFunctionPlot[322].Y = -0.078

	pointsOfFunctionPlot[323].X = 3.23
	pointsOfFunctionPlot[323].Y = -0.088

	pointsOfFunctionPlot[324].X = 3.24
	pointsOfFunctionPlot[324].Y = -0.098

	pointsOfFunctionPlot[325].X = 3.25
	pointsOfFunctionPlot[325].Y = -0.108

	pointsOfFunctionPlot[326].X = 3.26
	pointsOfFunctionPlot[326].Y = -0.118

	pointsOfFunctionPlot[327].X = 3.27
	pointsOfFunctionPlot[327].Y = -0.128

	pointsOfFunctionPlot[328].X = 3.28
	pointsOfFunctionPlot[328].Y = -0.138

	pointsOfFunctionPlot[329].X = 3.29
	pointsOfFunctionPlot[329].Y = -0.147

	pointsOfFunctionPlot[330].X = 3.30
	pointsOfFunctionPlot[330].Y = -0.157

	pointsOfFunctionPlot[331].X = 3.31
	pointsOfFunctionPlot[331].Y = -0.167

	pointsOfFunctionPlot[332].X = 3.32
	pointsOfFunctionPlot[332].Y = -0.177

	pointsOfFunctionPlot[333].X = 3.33
	pointsOfFunctionPlot[333].Y = -0.187

	pointsOfFunctionPlot[334].X = 3.34
	pointsOfFunctionPlot[334].Y = -0.197

	pointsOfFunctionPlot[335].X = 3.35
	pointsOfFunctionPlot[335].Y = -0.206

	pointsOfFunctionPlot[336].X = 3.36
	pointsOfFunctionPlot[336].Y = -0.216

	pointsOfFunctionPlot[337].X = 3.37
	pointsOfFunctionPlot[337].Y = -0.226

	pointsOfFunctionPlot[338].X = 3.38
	pointsOfFunctionPlot[338].Y = -0.236

	pointsOfFunctionPlot[339].X = 3.39
	pointsOfFunctionPlot[339].Y = -0.245

	pointsOfFunctionPlot[340].X = 3.40
	pointsOfFunctionPlot[340].Y = -0.255

	pointsOfFunctionPlot[341].X = 3.41
	pointsOfFunctionPlot[341].Y = -0.265

	pointsOfFunctionPlot[342].X = 3.42
	pointsOfFunctionPlot[342].Y = -0.274

	pointsOfFunctionPlot[343].X = 3.43
	pointsOfFunctionPlot[343].Y = -0.284

	pointsOfFunctionPlot[344].X = 3.44
	pointsOfFunctionPlot[344].Y = -0.294

	pointsOfFunctionPlot[345].X = 3.45
	pointsOfFunctionPlot[345].Y = -0.303

	pointsOfFunctionPlot[346].X = 3.46
	pointsOfFunctionPlot[346].Y = -0.313

	pointsOfFunctionPlot[347].X = 3.47
	pointsOfFunctionPlot[347].Y = -0.322

	pointsOfFunctionPlot[348].X = 3.48
	pointsOfFunctionPlot[348].Y = -0.332

	pointsOfFunctionPlot[349].X = 3.49
	pointsOfFunctionPlot[349].Y = -0.341

	pointsOfFunctionPlot[350].X = 3.50
	pointsOfFunctionPlot[350].Y = -0.35

	pointsOfFunctionPlot[351].X = 3.51
	pointsOfFunctionPlot[351].Y = -0.36

	pointsOfFunctionPlot[352].X = 3.52
	pointsOfFunctionPlot[352].Y = -0.369

	pointsOfFunctionPlot[353].X = 3.53
	pointsOfFunctionPlot[353].Y = -0.378

	pointsOfFunctionPlot[354].X = 3.54
	pointsOfFunctionPlot[354].Y = -0.388

	pointsOfFunctionPlot[355].X = 3.55
	pointsOfFunctionPlot[355].Y = -0.397

	pointsOfFunctionPlot[356].X = 3.56
	pointsOfFunctionPlot[356].Y = -0.406

	pointsOfFunctionPlot[357].X = 3.57
	pointsOfFunctionPlot[357].Y = -0.415

	pointsOfFunctionPlot[358].X = 3.58
	pointsOfFunctionPlot[358].Y = -0.424

	pointsOfFunctionPlot[359].X = 3.59
	pointsOfFunctionPlot[359].Y = -0.433

	pointsOfFunctionPlot[360].X = 3.60
	pointsOfFunctionPlot[360].Y = -0.442

	pointsOfFunctionPlot[361].X = 3.61
	pointsOfFunctionPlot[361].Y = -0.451

	pointsOfFunctionPlot[362].X = 3.62
	pointsOfFunctionPlot[362].Y = -0.46

	pointsOfFunctionPlot[363].X = 3.63
	pointsOfFunctionPlot[363].Y = -0.469

	pointsOfFunctionPlot[364].X = 3.64
	pointsOfFunctionPlot[364].Y = -0.478

	pointsOfFunctionPlot[365].X = 3.65
	pointsOfFunctionPlot[365].Y = -0.486

	pointsOfFunctionPlot[366].X = 3.66
	pointsOfFunctionPlot[366].Y = -0.495

	pointsOfFunctionPlot[367].X = 3.67
	pointsOfFunctionPlot[367].Y = -0.504

	pointsOfFunctionPlot[368].X = 3.68
	pointsOfFunctionPlot[368].Y = -0.512

	pointsOfFunctionPlot[369].X = 3.69
	pointsOfFunctionPlot[369].Y = -0.521

	pointsOfFunctionPlot[370].X = 3.70
	pointsOfFunctionPlot[370].Y = -0.529

	pointsOfFunctionPlot[371].X = 3.71
	pointsOfFunctionPlot[371].Y = -0.538

	pointsOfFunctionPlot[372].X = 3.72
	pointsOfFunctionPlot[372].Y = -0.546

	pointsOfFunctionPlot[373].X = 3.73
	pointsOfFunctionPlot[373].Y = -0.555

	pointsOfFunctionPlot[374].X = 3.74
	pointsOfFunctionPlot[374].Y = -0.563

	pointsOfFunctionPlot[375].X = 3.75
	pointsOfFunctionPlot[375].Y = -0.571

	pointsOfFunctionPlot[376].X = 3.76
	pointsOfFunctionPlot[376].Y = -0.579

	pointsOfFunctionPlot[377].X = 3.77
	pointsOfFunctionPlot[377].Y = -0.587

	pointsOfFunctionPlot[378].X = 3.78
	pointsOfFunctionPlot[378].Y = -0.595

	pointsOfFunctionPlot[379].X = 3.79
	pointsOfFunctionPlot[379].Y = -0.603

	pointsOfFunctionPlot[380].X = 3.80
	pointsOfFunctionPlot[380].Y = -0.611

	pointsOfFunctionPlot[381].X = 3.81
	pointsOfFunctionPlot[381].Y = -0.619

	pointsOfFunctionPlot[382].X = 3.82
	pointsOfFunctionPlot[382].Y = -0.627

	pointsOfFunctionPlot[383].X = 3.83
	pointsOfFunctionPlot[383].Y = -0.635

	pointsOfFunctionPlot[384].X = 3.84
	pointsOfFunctionPlot[384].Y = -0.643

	pointsOfFunctionPlot[385].X = 3.85
	pointsOfFunctionPlot[385].Y = -0.65

	pointsOfFunctionPlot[386].X = 3.86
	pointsOfFunctionPlot[386].Y = -0.658

	pointsOfFunctionPlot[387].X = 3.87
	pointsOfFunctionPlot[387].Y = -0.665

	pointsOfFunctionPlot[388].X = 3.88
	pointsOfFunctionPlot[388].Y = -0.673

	pointsOfFunctionPlot[389].X = 3.89
	pointsOfFunctionPlot[389].Y = -0.68

	pointsOfFunctionPlot[390].X = 3.90
	pointsOfFunctionPlot[390].Y = -0.687

	pointsOfFunctionPlot[391].X = 3.91
	pointsOfFunctionPlot[391].Y = -0.694

	pointsOfFunctionPlot[392].X = 3.92
	pointsOfFunctionPlot[392].Y = -0.702

	pointsOfFunctionPlot[393].X = 3.93
	pointsOfFunctionPlot[393].Y = -0.709

	pointsOfFunctionPlot[394].X = 3.94
	pointsOfFunctionPlot[394].Y = -0.716

	pointsOfFunctionPlot[395].X = 3.95
	pointsOfFunctionPlot[395].Y = -0.723

	pointsOfFunctionPlot[396].X = 3.96
	pointsOfFunctionPlot[396].Y = -0.73

	pointsOfFunctionPlot[397].X = 3.97
	pointsOfFunctionPlot[397].Y = -0.736

	pointsOfFunctionPlot[398].X = 3.98
	pointsOfFunctionPlot[398].Y = -0.743

	pointsOfFunctionPlot[399].X = 3.99
	pointsOfFunctionPlot[399].Y = -0.75

	pointsOfFunctionPlot[400].X = 4.0
	pointsOfFunctionPlot[400].Y = -0.756

	pointsOfFunctionPlot[401].X = 4.01
	pointsOfFunctionPlot[401].Y = -0.763

	pointsOfFunctionPlot[402].X = 4.02
	pointsOfFunctionPlot[402].Y = -0.769

	pointsOfFunctionPlot[403].X = 4.03
	pointsOfFunctionPlot[403].Y = -0.776

	pointsOfFunctionPlot[404].X = 4.04
	pointsOfFunctionPlot[404].Y = -0.782

	pointsOfFunctionPlot[405].X = 4.05
	pointsOfFunctionPlot[405].Y = -0.788

	pointsOfFunctionPlot[406].X = 4.06
	pointsOfFunctionPlot[406].Y = -0.793

	pointsOfFunctionPlot[407].X = 4.07
	pointsOfFunctionPlot[407].Y = -0.8

	pointsOfFunctionPlot[408].X = 4.08
	pointsOfFunctionPlot[408].Y = -0.806

	pointsOfFunctionPlot[409].X = 4.09
	pointsOfFunctionPlot[409].Y = -0.812

	pointsOfFunctionPlot[410].X = 4.10
	pointsOfFunctionPlot[410].Y = -0.818

	pointsOfFunctionPlot[411].X = 4.11
	pointsOfFunctionPlot[411].Y = -0.824

	pointsOfFunctionPlot[412].X = 4.12
	pointsOfFunctionPlot[412].Y = -0.829

	pointsOfFunctionPlot[413].X = 4.13
	pointsOfFunctionPlot[413].Y = -0.835

	pointsOfFunctionPlot[414].X = 4.14
	pointsOfFunctionPlot[414].Y = -0.84

	pointsOfFunctionPlot[415].X = 4.15
	pointsOfFunctionPlot[415].Y = -0.846

	pointsOfFunctionPlot[416].X = 4.16
	pointsOfFunctionPlot[416].Y = -0.851

	pointsOfFunctionPlot[417].X = 4.17
	pointsOfFunctionPlot[417].Y = -0.856

	pointsOfFunctionPlot[418].X = 4.18
	pointsOfFunctionPlot[418].Y = -0.861

	pointsOfFunctionPlot[419].X = 4.19
	pointsOfFunctionPlot[419].Y = -0.866

	pointsOfFunctionPlot[420].X = 4.20
	pointsOfFunctionPlot[420].Y = -0.871

	pointsOfFunctionPlot[421].X = 4.21
	pointsOfFunctionPlot[421].Y = -0.874

	pointsOfFunctionPlot[422].X = 4.22
	pointsOfFunctionPlot[422].Y = -0.881

	pointsOfFunctionPlot[423].X = 4.23
	pointsOfFunctionPlot[423].Y = -0.885

	pointsOfFunctionPlot[424].X = 4.24
	pointsOfFunctionPlot[424].Y = -0.89

	pointsOfFunctionPlot[425].X = 4.25
	pointsOfFunctionPlot[425].Y = -0.895

	pointsOfFunctionPlot[426].X = 4.26
	pointsOfFunctionPlot[426].Y = -0.899

	pointsOfFunctionPlot[427].X = 4.27
	pointsOfFunctionPlot[427].Y = -0.903

	pointsOfFunctionPlot[428].X = 4.28
	pointsOfFunctionPlot[428].Y = -0.908

	pointsOfFunctionPlot[429].X = 4.29
	pointsOfFunctionPlot[429].Y = -0.912

	pointsOfFunctionPlot[430].X = 4.30
	pointsOfFunctionPlot[430].Y = -0.916

	pointsOfFunctionPlot[431].X = 4.31
	pointsOfFunctionPlot[431].Y = -0.92

	pointsOfFunctionPlot[432].X = 4.32
	pointsOfFunctionPlot[432].Y = -0.924

	pointsOfFunctionPlot[433].X = 4.33
	pointsOfFunctionPlot[433].Y = -0.927

	pointsOfFunctionPlot[434].X = 4.34
	pointsOfFunctionPlot[434].Y = -0.931

	pointsOfFunctionPlot[435].X = 4.35
	pointsOfFunctionPlot[435].Y = -0.935

	pointsOfFunctionPlot[436].X = 4.36
	pointsOfFunctionPlot[436].Y = -0.938

	pointsOfFunctionPlot[437].X = 4.37
	pointsOfFunctionPlot[437].Y = -0.942

	pointsOfFunctionPlot[438].X = 4.38
	pointsOfFunctionPlot[438].Y = -0.945

	pointsOfFunctionPlot[439].X = 4.39
	pointsOfFunctionPlot[439].Y = -0.948

	pointsOfFunctionPlot[440].X = 4.40
	pointsOfFunctionPlot[440].Y = -0.951

	pointsOfFunctionPlot[441].X = 4.41
	pointsOfFunctionPlot[441].Y = -0.954

	pointsOfFunctionPlot[442].X = 4.42
	pointsOfFunctionPlot[442].Y = -0.957

	pointsOfFunctionPlot[443].X = 4.43
	pointsOfFunctionPlot[443].Y = -0.96

	pointsOfFunctionPlot[444].X = 4.44
	pointsOfFunctionPlot[444].Y = -0.963

	pointsOfFunctionPlot[445].X = 4.45
	pointsOfFunctionPlot[445].Y = -0.965

	pointsOfFunctionPlot[446].X = 4.46
	pointsOfFunctionPlot[446].Y = -0.968

	pointsOfFunctionPlot[447].X = 4.47
	pointsOfFunctionPlot[447].Y = -0.97

	pointsOfFunctionPlot[448].X = 4.48
	pointsOfFunctionPlot[448].Y = -0.973

	pointsOfFunctionPlot[449].X = 4.49
	pointsOfFunctionPlot[449].Y = -0.975

	pointsOfFunctionPlot[450].X = 4.50
	pointsOfFunctionPlot[450].Y = -0.977

	pointsOfFunctionPlot[451].X = 4.51
	pointsOfFunctionPlot[451].Y = -0.979

	pointsOfFunctionPlot[452].X = 4.52
	pointsOfFunctionPlot[452].Y = -0.981

	pointsOfFunctionPlot[453].X = 4.53
	pointsOfFunctionPlot[453].Y = -0.983

	pointsOfFunctionPlot[454].X = 4.54
	pointsOfFunctionPlot[454].Y = -0.985

	pointsOfFunctionPlot[455].X = 4.55
	pointsOfFunctionPlot[455].Y = -0.986

	pointsOfFunctionPlot[456].X = 4.56
	pointsOfFunctionPlot[456].Y = -0.988

	pointsOfFunctionPlot[457].X = 4.57
	pointsOfFunctionPlot[457].Y = -0.989

	pointsOfFunctionPlot[458].X = 4.58
	pointsOfFunctionPlot[458].Y = -0.991

	pointsOfFunctionPlot[459].X = 4.59
	pointsOfFunctionPlot[459].Y = -0.992

	pointsOfFunctionPlot[460].X = 4.60
	pointsOfFunctionPlot[460].Y = -0.993

	pointsOfFunctionPlot[461].X = 4.61
	pointsOfFunctionPlot[461].Y = -0.994

	pointsOfFunctionPlot[462].X = 4.62
	pointsOfFunctionPlot[462].Y = -0.995

	pointsOfFunctionPlot[463].X = 4.63
	pointsOfFunctionPlot[463].Y = -0.996

	pointsOfFunctionPlot[464].X = 4.64
	pointsOfFunctionPlot[464].Y = -0.997

	pointsOfFunctionPlot[465].X = 4.65
	pointsOfFunctionPlot[465].Y = -0.998

	pointsOfFunctionPlot[466].X = 4.66
	pointsOfFunctionPlot[466].Y = -0.998

	pointsOfFunctionPlot[467].X = 4.67
	pointsOfFunctionPlot[467].Y = -0.999

	pointsOfFunctionPlot[468].X = 4.68
	pointsOfFunctionPlot[468].Y = -0.999

	pointsOfFunctionPlot[469].X = 4.69
	pointsOfFunctionPlot[469].Y = -0.999

	pointsOfFunctionPlot[470].X = 4.70
	pointsOfFunctionPlot[470].Y = -0.999

	pointsOfFunctionPlot[471].X = 4.71
	pointsOfFunctionPlot[471].Y = -1.0

	pointsOfFunctionPlot[472].X = 4.72
	pointsOfFunctionPlot[472].Y = -1.0

	pointsOfFunctionPlot[473].X = 4.73
	pointsOfFunctionPlot[473].Y = -0.999

	pointsOfFunctionPlot[474].X = 4.74
	pointsOfFunctionPlot[474].Y = -0.999

	pointsOfFunctionPlot[475].X = 4.75
	pointsOfFunctionPlot[475].Y = -0.999

	pointsOfFunctionPlot[476].X = 4.76
	pointsOfFunctionPlot[476].Y = -0.998

	pointsOfFunctionPlot[477].X = 4.77
	pointsOfFunctionPlot[477].Y = -0.998

	pointsOfFunctionPlot[478].X = 4.78
	pointsOfFunctionPlot[478].Y = -0.997

	pointsOfFunctionPlot[479].X = 4.79
	pointsOfFunctionPlot[479].Y = -0.997

	pointsOfFunctionPlot[480].X = 4.80
	pointsOfFunctionPlot[480].Y = -0.996

	pointsOfFunctionPlot[481].X = 4.81
	pointsOfFunctionPlot[481].Y = -0.995

	pointsOfFunctionPlot[482].X = 4.82
	pointsOfFunctionPlot[482].Y = -0.994

	pointsOfFunctionPlot[483].X = 4.83
	pointsOfFunctionPlot[483].Y = -0.993

	pointsOfFunctionPlot[484].X = 4.84
	pointsOfFunctionPlot[484].Y = -0.991

	pointsOfFunctionPlot[485].X = 4.85
	pointsOfFunctionPlot[485].Y = -0.99

	pointsOfFunctionPlot[486].X = 4.86
	pointsOfFunctionPlot[486].Y = -0.989

	pointsOfFunctionPlot[487].X = 4.87
	pointsOfFunctionPlot[487].Y = -0.987

	pointsOfFunctionPlot[488].X = 4.88
	pointsOfFunctionPlot[488].Y = -0.986

	pointsOfFunctionPlot[489].X = 4.89
	pointsOfFunctionPlot[489].Y = -0.984

	pointsOfFunctionPlot[490].X = 4.90
	pointsOfFunctionPlot[490].Y = -0.982

	pointsOfFunctionPlot[491].X = 4.91
	pointsOfFunctionPlot[491].Y = -0.98

	pointsOfFunctionPlot[492].X = 4.92
	pointsOfFunctionPlot[492].Y = -0.978

	pointsOfFunctionPlot[493].X = 4.93
	pointsOfFunctionPlot[493].Y = -0.976

	pointsOfFunctionPlot[494].X = 4.94
	pointsOfFunctionPlot[494].Y = -0.974

	pointsOfFunctionPlot[495].X = 4.95
	pointsOfFunctionPlot[495].Y = -0.971

	pointsOfFunctionPlot[496].X = 4.96
	pointsOfFunctionPlot[496].Y = -0.969

	pointsOfFunctionPlot[497].X = 4.97
	pointsOfFunctionPlot[497].Y = -0.967

	pointsOfFunctionPlot[498].X = 4.98
	pointsOfFunctionPlot[498].Y = -0.964

	pointsOfFunctionPlot[499].X = 4.99
	pointsOfFunctionPlot[499].Y = -0.961

	pointsOfFunctionPlot[500].X = 5.0
	pointsOfFunctionPlot[500].Y = -0.958

	pointsOfFunctionPlot[501].X = 5.01
	pointsOfFunctionPlot[501].Y = -0.956

	pointsOfFunctionPlot[502].X = 5.02
	pointsOfFunctionPlot[502].Y = -0.953

	pointsOfFunctionPlot[503].X = 5.03
	pointsOfFunctionPlot[503].Y = -0.95

	pointsOfFunctionPlot[504].X = 5.04
	pointsOfFunctionPlot[504].Y = -0.946

	pointsOfFunctionPlot[505].X = 5.05
	pointsOfFunctionPlot[505].Y = -0.943

	pointsOfFunctionPlot[506].X = 5.06
	pointsOfFunctionPlot[506].Y = -0.94

	pointsOfFunctionPlot[507].X = 5.07
	pointsOfFunctionPlot[507].Y = -0.936

	pointsOfFunctionPlot[508].X = 5.08
	pointsOfFunctionPlot[508].Y = -0.933

	pointsOfFunctionPlot[509].X = 5.09
	pointsOfFunctionPlot[509].Y = -0.929

	pointsOfFunctionPlot[510].X = 5.10
	pointsOfFunctionPlot[510].Y = -0.925

	pointsOfFunctionPlot[511].X = 5.11
	pointsOfFunctionPlot[511].Y = -0.922

	pointsOfFunctionPlot[512].X = 5.12
	pointsOfFunctionPlot[512].Y = -0.918

	pointsOfFunctionPlot[513].X = 5.13
	pointsOfFunctionPlot[513].Y = -0.914

	pointsOfFunctionPlot[514].X = 5.14
	pointsOfFunctionPlot[514].Y = -0.91

	pointsOfFunctionPlot[515].X = 5.15
	pointsOfFunctionPlot[515].Y = -0.905

	pointsOfFunctionPlot[516].X = 5.16
	pointsOfFunctionPlot[516].Y = -0.901

	pointsOfFunctionPlot[517].X = 5.17
	pointsOfFunctionPlot[517].Y = -0.897

	pointsOfFunctionPlot[518].X = 5.18
	pointsOfFunctionPlot[518].Y = -0.892

	pointsOfFunctionPlot[519].X = 5.19
	pointsOfFunctionPlot[519].Y = -0.888

	pointsOfFunctionPlot[520].X = 5.20
	pointsOfFunctionPlot[520].Y = -0.883

	pointsOfFunctionPlot[521].X = 5.21
	pointsOfFunctionPlot[521].Y = -0.878

	pointsOfFunctionPlot[522].X = 5.22
	pointsOfFunctionPlot[522].Y = -0.873

	pointsOfFunctionPlot[523].X = 5.23
	pointsOfFunctionPlot[523].Y = -0.869

	pointsOfFunctionPlot[524].X = 5.24
	pointsOfFunctionPlot[524].Y = -0.864

	pointsOfFunctionPlot[525].X = 5.25
	pointsOfFunctionPlot[525].Y = -0.858

	pointsOfFunctionPlot[526].X = 5.26
	pointsOfFunctionPlot[526].Y = -0.853

	pointsOfFunctionPlot[527].X = 5.27
	pointsOfFunctionPlot[527].Y = -0.848

	pointsOfFunctionPlot[528].X = 5.28
	pointsOfFunctionPlot[528].Y = -0.843

	pointsOfFunctionPlot[529].X = 5.29
	pointsOfFunctionPlot[529].Y = -0.837

	pointsOfFunctionPlot[530].X = 5.30
	pointsOfFunctionPlot[530].Y = -0.832

	pointsOfFunctionPlot[531].X = 5.31
	pointsOfFunctionPlot[531].Y = -0.826

	pointsOfFunctionPlot[532].X = 5.32
	pointsOfFunctionPlot[532].Y = -0.821

	pointsOfFunctionPlot[533].X = 5.33
	pointsOfFunctionPlot[533].Y = -0.815

	pointsOfFunctionPlot[534].X = 5.34
	pointsOfFunctionPlot[534].Y = -0.809

	pointsOfFunctionPlot[535].X = 5.35
	pointsOfFunctionPlot[535].Y = -0.803

	pointsOfFunctionPlot[536].X = 5.36
	pointsOfFunctionPlot[536].Y = -0.797

	pointsOfFunctionPlot[537].X = 5.37
	pointsOfFunctionPlot[537].Y = -0.791

	pointsOfFunctionPlot[538].X = 5.38
	pointsOfFunctionPlot[538].Y = -0.785

	pointsOfFunctionPlot[539].X = 5.39
	pointsOfFunctionPlot[539].Y = -0.779

	pointsOfFunctionPlot[540].X = 5.40
	pointsOfFunctionPlot[540].Y = -0.772

	pointsOfFunctionPlot[541].X = 5.41
	pointsOfFunctionPlot[541].Y = -0.766

	pointsOfFunctionPlot[542].X = 5.42
	pointsOfFunctionPlot[542].Y = -0.759

	pointsOfFunctionPlot[543].X = 5.43
	pointsOfFunctionPlot[543].Y = -0.753

	pointsOfFunctionPlot[544].X = 5.44
	pointsOfFunctionPlot[544].Y = -0.746

	pointsOfFunctionPlot[545].X = 5.45
	pointsOfFunctionPlot[545].Y = -0.74

	pointsOfFunctionPlot[546].X = 5.46
	pointsOfFunctionPlot[546].Y = -0.733

	pointsOfFunctionPlot[547].X = 5.47
	pointsOfFunctionPlot[547].Y = -0.726

	pointsOfFunctionPlot[548].X = 5.48
	pointsOfFunctionPlot[548].Y = -0.719

	pointsOfFunctionPlot[549].X = 5.49
	pointsOfFunctionPlot[549].Y = -0.712

	pointsOfFunctionPlot[550].X = 5.50
	pointsOfFunctionPlot[550].Y = -0.705

	pointsOfFunctionPlot[551].X = 5.51
	pointsOfFunctionPlot[551].Y = -0.698

	pointsOfFunctionPlot[552].X = 5.52
	pointsOfFunctionPlot[552].Y = -0.691

	pointsOfFunctionPlot[553].X = 5.53
	pointsOfFunctionPlot[553].Y = -0.684

	pointsOfFunctionPlot[554].X = 5.54
	pointsOfFunctionPlot[554].Y = -0.676

	pointsOfFunctionPlot[555].X = 5.55
	pointsOfFunctionPlot[555].Y = -0.669

	pointsOfFunctionPlot[556].X = 5.56
	pointsOfFunctionPlot[556].Y = -0.661

	pointsOfFunctionPlot[557].X = 5.57
	pointsOfFunctionPlot[557].Y = -0.654

	pointsOfFunctionPlot[558].X = 5.58
	pointsOfFunctionPlot[558].Y = -0.646

	pointsOfFunctionPlot[559].X = 5.59
	pointsOfFunctionPlot[559].Y = -0.639

	pointsOfFunctionPlot[560].X = 5.60
	pointsOfFunctionPlot[560].Y = -0.631

	pointsOfFunctionPlot[561].X = 5.61
	pointsOfFunctionPlot[561].Y = -0.623

	pointsOfFunctionPlot[562].X = 5.62
	pointsOfFunctionPlot[562].Y = -0.615

	pointsOfFunctionPlot[563].X = 5.63
	pointsOfFunctionPlot[563].Y = -0.607

	pointsOfFunctionPlot[564].X = 5.64
	pointsOfFunctionPlot[564].Y = -0.599

	pointsOfFunctionPlot[565].X = 5.65
	pointsOfFunctionPlot[565].Y = -0.591

	pointsOfFunctionPlot[566].X = 5.66
	pointsOfFunctionPlot[566].Y = -0.583

	pointsOfFunctionPlot[567].X = 5.67
	pointsOfFunctionPlot[567].Y = -0.575

	pointsOfFunctionPlot[568].X = 5.68
	pointsOfFunctionPlot[568].Y = -0.567

	pointsOfFunctionPlot[569].X = 5.69
	pointsOfFunctionPlot[569].Y = -0.559

	pointsOfFunctionPlot[570].X = 5.70
	pointsOfFunctionPlot[570].Y = -0.55

	pointsOfFunctionPlot[571].X = 5.71
	pointsOfFunctionPlot[571].Y = -0.542

	pointsOfFunctionPlot[572].X = 5.72
	pointsOfFunctionPlot[572].Y = -0.533

	pointsOfFunctionPlot[573].X = 5.73
	pointsOfFunctionPlot[573].Y = -0.525

	pointsOfFunctionPlot[574].X = 5.74
	pointsOfFunctionPlot[574].Y = -0.516

	pointsOfFunctionPlot[575].X = 5.75
	pointsOfFunctionPlot[575].Y = -0.508

	pointsOfFunctionPlot[576].X = 5.76
	pointsOfFunctionPlot[576].Y = -0.499

	pointsOfFunctionPlot[577].X = 5.77
	pointsOfFunctionPlot[577].Y = -0.491

	pointsOfFunctionPlot[578].X = 5.78
	pointsOfFunctionPlot[578].Y = -0.482

	pointsOfFunctionPlot[579].X = 5.79
	pointsOfFunctionPlot[579].Y = -0.473

	pointsOfFunctionPlot[580].X = 5.80
	pointsOfFunctionPlot[580].Y = -0.464

	pointsOfFunctionPlot[581].X = 5.81
	pointsOfFunctionPlot[581].Y = -0.455

	pointsOfFunctionPlot[582].X = 5.82
	pointsOfFunctionPlot[582].Y = -0.446

	pointsOfFunctionPlot[583].X = 5.83
	pointsOfFunctionPlot[583].Y = -0.437

	pointsOfFunctionPlot[584].X = 5.84
	pointsOfFunctionPlot[584].Y = -0.428

	pointsOfFunctionPlot[585].X = 5.85
	pointsOfFunctionPlot[585].Y = -0.419

	pointsOfFunctionPlot[586].X = 5.86
	pointsOfFunctionPlot[586].Y = -0.41

	pointsOfFunctionPlot[587].X = 5.87
	pointsOfFunctionPlot[587].Y = -0.401

	pointsOfFunctionPlot[588].X = 5.88
	pointsOfFunctionPlot[588].Y = -0.392

	pointsOfFunctionPlot[589].X = 5.89
	pointsOfFunctionPlot[589].Y = -0.383

	pointsOfFunctionPlot[590].X = 5.90
	pointsOfFunctionPlot[590].Y = -0.373

	pointsOfFunctionPlot[591].X = 5.91
	pointsOfFunctionPlot[591].Y = -0.364

	pointsOfFunctionPlot[592].X = 5.92
	pointsOfFunctionPlot[592].Y = -0.355

	pointsOfFunctionPlot[593].X = 5.93
	pointsOfFunctionPlot[593].Y = -0.345

	pointsOfFunctionPlot[594].X = 5.94
	pointsOfFunctionPlot[594].Y = -0.336

	pointsOfFunctionPlot[595].X = 5.95
	pointsOfFunctionPlot[595].Y = -0.327

	pointsOfFunctionPlot[596].X = 5.96
	pointsOfFunctionPlot[596].Y = -0.317

	pointsOfFunctionPlot[597].X = 5.97
	pointsOfFunctionPlot[597].Y = -0.308

	pointsOfFunctionPlot[598].X = 5.98
	pointsOfFunctionPlot[598].Y = -0.298

	pointsOfFunctionPlot[599].X = 5.99
	pointsOfFunctionPlot[599].Y = -0.289

	pointsOfFunctionPlot[600].X = 6.0
	pointsOfFunctionPlot[600].Y = -0.279

	pointsOfFunctionPlot[601].X = 6.01
	pointsOfFunctionPlot[601].Y = -0.269

	pointsOfFunctionPlot[602].X = 6.02
	pointsOfFunctionPlot[602].Y = -0.26

	pointsOfFunctionPlot[603].X = 6.03
	pointsOfFunctionPlot[603].Y = -0.25

	pointsOfFunctionPlot[604].X = 6.04
	pointsOfFunctionPlot[604].Y = -0.24

	pointsOfFunctionPlot[605].X = 6.05
	pointsOfFunctionPlot[605].Y = -0.231

	pointsOfFunctionPlot[606].X = 6.06
	pointsOfFunctionPlot[606].Y = -0.221

	pointsOfFunctionPlot[607].X = 6.07
	pointsOfFunctionPlot[607].Y = -0.211

	pointsOfFunctionPlot[608].X = 6.08
	pointsOfFunctionPlot[608].Y = -0.201

	pointsOfFunctionPlot[609].X = 6.09
	pointsOfFunctionPlot[609].Y = -0.192

	pointsOfFunctionPlot[610].X = 6.10
	pointsOfFunctionPlot[610].Y = -0.182

	pointsOfFunctionPlot[611].X = 6.11
	pointsOfFunctionPlot[611].Y = -0.172

	pointsOfFunctionPlot[612].X = 6.12
	pointsOfFunctionPlot[612].Y = -0.162

	pointsOfFunctionPlot[613].X = 6.13
	pointsOfFunctionPlot[613].Y = -0.152

	pointsOfFunctionPlot[614].X = 6.14
	pointsOfFunctionPlot[614].Y = -0.142

	pointsOfFunctionPlot[615].X = 6.15
	pointsOfFunctionPlot[615].Y = -0.132

	pointsOfFunctionPlot[616].X = 6.16
	pointsOfFunctionPlot[616].Y = -0.122

	pointsOfFunctionPlot[617].X = 6.17
	pointsOfFunctionPlot[617].Y = -0.112

	pointsOfFunctionPlot[618].X = 6.18
	pointsOfFunctionPlot[618].Y = -0.103

	pointsOfFunctionPlot[619].X = 6.19
	pointsOfFunctionPlot[619].Y = -0.093

	pointsOfFunctionPlot[620].X = 6.20
	pointsOfFunctionPlot[620].Y = -0.083

	pointsOfFunctionPlot[621].X = 6.21
	pointsOfFunctionPlot[621].Y = -0.073

	pointsOfFunctionPlot[622].X = 6.22
	pointsOfFunctionPlot[622].Y = -0.063

	pointsOfFunctionPlot[623].X = 6.23
	pointsOfFunctionPlot[623].Y = -0.053

	pointsOfFunctionPlot[624].X = 6.24
	pointsOfFunctionPlot[624].Y = -0.043

	pointsOfFunctionPlot[625].X = 6.25
	pointsOfFunctionPlot[625].Y = -0.033

	pointsOfFunctionPlot[626].X = 6.26
	pointsOfFunctionPlot[626].Y = -0.023

	pointsOfFunctionPlot[627].X = 6.27
	pointsOfFunctionPlot[627].Y = -0.013

	pointsOfFunctionPlot[628].X = 6.28
	pointsOfFunctionPlot[628].Y = -0.003

	pointsOfFunctionPlot[629].X = 6.29
	pointsOfFunctionPlot[629].Y = 0.006

	pointsOfFunctionPlot[630].X = 6.30
	pointsOfFunctionPlot[630].Y = 0.016





	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function sin(x)"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("sin(x)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"sin-function-plot-01.png"); err != nil {

		panic(err)
	}
}
