package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function cos(x).

	pointsOfFunctionPlot := make(plotter.XYs, 321)

	pointsOfFunctionPlot[0].X = 0.0
	pointsOfFunctionPlot[0].Y = 1.0

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = 1.0

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = 0.999

	pointsOfFunctionPlot[3].X = 0.03
	pointsOfFunctionPlot[3].Y = 0.999

	pointsOfFunctionPlot[4].X = 0.04
	pointsOfFunctionPlot[4].Y = 0.999

	pointsOfFunctionPlot[5].X = 0.05
	pointsOfFunctionPlot[5].Y = 0.998

	pointsOfFunctionPlot[6].X = 0.06
	pointsOfFunctionPlot[6].Y = 0.998

	pointsOfFunctionPlot[7].X = 0.07
	pointsOfFunctionPlot[7].Y = 0.997

	pointsOfFunctionPlot[8].X = 0.08
	pointsOfFunctionPlot[8].Y = 0.996

	pointsOfFunctionPlot[9].X = 0.09
	pointsOfFunctionPlot[9].Y = 0.996

	pointsOfFunctionPlot[10].X = 0.1
	pointsOfFunctionPlot[10].Y = 0.995

	pointsOfFunctionPlot[11].X = 0.11
	pointsOfFunctionPlot[11].Y = 0.994

	pointsOfFunctionPlot[12].X = 0.12
	pointsOfFunctionPlot[12].Y = 0.992

	pointsOfFunctionPlot[13].X = 0.13
	pointsOfFunctionPlot[13].Y = 0.991

	pointsOfFunctionPlot[14].X = 0.14
	pointsOfFunctionPlot[14].Y = 0.99

	pointsOfFunctionPlot[15].X = 0.15
	pointsOfFunctionPlot[15].Y = 0.988

	pointsOfFunctionPlot[16].X = 0.16
	pointsOfFunctionPlot[16].Y = 0.987

	pointsOfFunctionPlot[17].X = 0.17
	pointsOfFunctionPlot[17].Y = 0.985

	pointsOfFunctionPlot[18].X = 0.18
	pointsOfFunctionPlot[18].Y = 0.983

	pointsOfFunctionPlot[19].X = 0.19
	pointsOfFunctionPlot[19].Y = 0.982

	pointsOfFunctionPlot[20].X = 0.2
	pointsOfFunctionPlot[20].Y = 0.98

	pointsOfFunctionPlot[21].X = 0.21
	pointsOfFunctionPlot[21].Y = 0.978

	pointsOfFunctionPlot[22].X = 0.22
	pointsOfFunctionPlot[22].Y = 0.975

	pointsOfFunctionPlot[23].X = 0.23
	pointsOfFunctionPlot[23].Y = 0.973

	pointsOfFunctionPlot[24].X = 0.24
	pointsOfFunctionPlot[24].Y = 0.971

	pointsOfFunctionPlot[25].X = 0.25
	pointsOfFunctionPlot[25].Y = 0.968

	pointsOfFunctionPlot[26].X = 0.26
	pointsOfFunctionPlot[26].Y = 0.966

	pointsOfFunctionPlot[27].X = 0.27
	pointsOfFunctionPlot[27].Y = 0.963

	pointsOfFunctionPlot[28].X = 0.28
	pointsOfFunctionPlot[28].Y = 0.961

	pointsOfFunctionPlot[29].X = 0.29
	pointsOfFunctionPlot[29].Y = 0.958

	pointsOfFunctionPlot[30].X = 0.3
	pointsOfFunctionPlot[30].Y = 0.955

	pointsOfFunctionPlot[31].X = 0.31
	pointsOfFunctionPlot[31].Y = 0.952

	pointsOfFunctionPlot[32].X = 0.32
	pointsOfFunctionPlot[32].Y = 0.949

	pointsOfFunctionPlot[33].X = 0.33
	pointsOfFunctionPlot[33].Y = 0.946

	pointsOfFunctionPlot[34].X = 0.34
	pointsOfFunctionPlot[34].Y = 0.942

	pointsOfFunctionPlot[35].X = 0.35
	pointsOfFunctionPlot[35].Y = 0.939

	pointsOfFunctionPlot[36].X = 0.36
	pointsOfFunctionPlot[36].Y = 0.935

	pointsOfFunctionPlot[37].X = 0.37
	pointsOfFunctionPlot[37].Y = 0.932

	pointsOfFunctionPlot[38].X = 0.38
	pointsOfFunctionPlot[38].Y = 0.928

	pointsOfFunctionPlot[39].X = 0.39
	pointsOfFunctionPlot[39].Y = 0.924

	pointsOfFunctionPlot[40].X = 0.40
	pointsOfFunctionPlot[40].Y = 0.921

	pointsOfFunctionPlot[41].X = 0.41
	pointsOfFunctionPlot[41].Y = 0.917

	pointsOfFunctionPlot[42].X = 0.42
	pointsOfFunctionPlot[42].Y = 0.913

	pointsOfFunctionPlot[43].X = 0.43
	pointsOfFunctionPlot[43].Y = 0.909

	pointsOfFunctionPlot[44].X = 0.44
	pointsOfFunctionPlot[44].Y = 0.904

	pointsOfFunctionPlot[45].X = 0.45
	pointsOfFunctionPlot[45].Y = 0.9

	pointsOfFunctionPlot[46].X = 0.46
	pointsOfFunctionPlot[46].Y = 0.896

	pointsOfFunctionPlot[47].X = 0.47
	pointsOfFunctionPlot[47].Y = 0.891

	pointsOfFunctionPlot[48].X = 0.48
	pointsOfFunctionPlot[48].Y = 0.887

	pointsOfFunctionPlot[49].X = 0.49
	pointsOfFunctionPlot[49].Y = 0.882

	pointsOfFunctionPlot[50].X = 0.5
	pointsOfFunctionPlot[50].Y = 0.877

	pointsOfFunctionPlot[51].X = 0.51
	pointsOfFunctionPlot[51].Y = 0.872

	pointsOfFunctionPlot[52].X = 0.52
	pointsOfFunctionPlot[52].Y = 0.867

	pointsOfFunctionPlot[53].X = 0.53
	pointsOfFunctionPlot[53].Y = 0.862

	pointsOfFunctionPlot[54].X = 0.54
	pointsOfFunctionPlot[54].Y = 0.857

	pointsOfFunctionPlot[55].X = 0.55
	pointsOfFunctionPlot[55].Y = 0.852

	pointsOfFunctionPlot[56].X = 0.56
	pointsOfFunctionPlot[56].Y = 0.847

	pointsOfFunctionPlot[57].X = 0.57
	pointsOfFunctionPlot[57].Y = 0.841

	pointsOfFunctionPlot[58].X = 0.58
	pointsOfFunctionPlot[58].Y = 0.836

	pointsOfFunctionPlot[59].X = 0.59
	pointsOfFunctionPlot[59].Y = 0.83

	pointsOfFunctionPlot[60].X = 0.6
	pointsOfFunctionPlot[60].Y = 0.825

	pointsOfFunctionPlot[61].X = 0.61
	pointsOfFunctionPlot[61].Y = 0.819

	pointsOfFunctionPlot[62].X = 0.62
	pointsOfFunctionPlot[62].Y = 0.813

	pointsOfFunctionPlot[63].X = 0.63
	pointsOfFunctionPlot[63].Y = 0.808

	pointsOfFunctionPlot[64].X = 0.64
	pointsOfFunctionPlot[64].Y = 0.802

	pointsOfFunctionPlot[65].X = 0.65
	pointsOfFunctionPlot[65].Y = 0.796

	pointsOfFunctionPlot[66].X = 0.66
	pointsOfFunctionPlot[66].Y = 0.79

	pointsOfFunctionPlot[67].X = 0.67
	pointsOfFunctionPlot[67].Y = 0.783

	pointsOfFunctionPlot[68].X = 0.68
	pointsOfFunctionPlot[68].Y = 0.777

	pointsOfFunctionPlot[69].X = 0.69
	pointsOfFunctionPlot[69].Y = 0.771

	pointsOfFunctionPlot[70].X = 0.7
	pointsOfFunctionPlot[70].Y = 0.764

	pointsOfFunctionPlot[71].X = 0.71
	pointsOfFunctionPlot[71].Y = 0.758

	pointsOfFunctionPlot[72].X = 0.72
	pointsOfFunctionPlot[72].Y = 0.751

	pointsOfFunctionPlot[73].X = 0.73
	pointsOfFunctionPlot[73].Y = 0.745

	pointsOfFunctionPlot[74].X = 0.74
	pointsOfFunctionPlot[74].Y = 0.738

	pointsOfFunctionPlot[75].X = 0.75
	pointsOfFunctionPlot[75].Y = 0.731

	pointsOfFunctionPlot[76].X = 0.76
	pointsOfFunctionPlot[76].Y = 0.724

	pointsOfFunctionPlot[77].X = 0.77
	pointsOfFunctionPlot[77].Y = 0.717

	pointsOfFunctionPlot[78].X = 0.78
	pointsOfFunctionPlot[78].Y = 0.71

	pointsOfFunctionPlot[79].X = 0.79
	pointsOfFunctionPlot[79].Y = 0.703

	pointsOfFunctionPlot[80].X = 0.8
	pointsOfFunctionPlot[80].Y = 0.696

	pointsOfFunctionPlot[81].X = 0.81
	pointsOfFunctionPlot[81].Y = 0.689

	pointsOfFunctionPlot[82].X = 0.82
	pointsOfFunctionPlot[82].Y = 0.682

	pointsOfFunctionPlot[83].X = 0.83
	pointsOfFunctionPlot[83].Y = 0.674

	pointsOfFunctionPlot[84].X = 0.84
	pointsOfFunctionPlot[84].Y = 0.667

	pointsOfFunctionPlot[85].X = 0.85
	pointsOfFunctionPlot[85].Y = 0.66

	pointsOfFunctionPlot[86].X = 0.86
	pointsOfFunctionPlot[86].Y = 0.652

	pointsOfFunctionPlot[87].X = 0.87
	pointsOfFunctionPlot[87].Y = 0.644

	pointsOfFunctionPlot[88].X = 0.88
	pointsOfFunctionPlot[88].Y = 0.637

	pointsOfFunctionPlot[89].X = 0.89
	pointsOfFunctionPlot[89].Y = 0.629

	pointsOfFunctionPlot[90].X = 0.9
	pointsOfFunctionPlot[90].Y = 0.621

	pointsOfFunctionPlot[91].X = 0.91
	pointsOfFunctionPlot[91].Y = 0.613

	pointsOfFunctionPlot[92].X = 0.92
	pointsOfFunctionPlot[92].Y = 0.605

	pointsOfFunctionPlot[93].X = 0.93
	pointsOfFunctionPlot[93].Y = 0.597

	pointsOfFunctionPlot[94].X = 0.94
	pointsOfFunctionPlot[94].Y = 0.589

	pointsOfFunctionPlot[95].X = 0.95
	pointsOfFunctionPlot[95].Y = 0.581

	pointsOfFunctionPlot[96].X = 0.96
	pointsOfFunctionPlot[96].Y = 0.573

	pointsOfFunctionPlot[97].X = 0.97
	pointsOfFunctionPlot[97].Y = 0.565

	pointsOfFunctionPlot[98].X = 0.98
	pointsOfFunctionPlot[98].Y = 0.557

	pointsOfFunctionPlot[99].X = 0.99
	pointsOfFunctionPlot[99].Y = 0.548

	pointsOfFunctionPlot[100].X = 1.0
	pointsOfFunctionPlot[100].Y = 0.54

	pointsOfFunctionPlot[101].X = 1.01
	pointsOfFunctionPlot[101].Y = 0.531

	pointsOfFunctionPlot[102].X = 1.02
	pointsOfFunctionPlot[102].Y = 0.523

	pointsOfFunctionPlot[103].X = 1.03
	pointsOfFunctionPlot[103].Y = 0.514

	pointsOfFunctionPlot[104].X = 1.04
	pointsOfFunctionPlot[104].Y = 0.506

	pointsOfFunctionPlot[105].X = 1.05
	pointsOfFunctionPlot[105].Y = 0.497

	pointsOfFunctionPlot[106].X = 1.06
	pointsOfFunctionPlot[106].Y = 0.488

	pointsOfFunctionPlot[107].X = 1.07
	pointsOfFunctionPlot[107].Y = 0.48

	pointsOfFunctionPlot[108].X = 1.08
	pointsOfFunctionPlot[108].Y = 0.471

	pointsOfFunctionPlot[109].X = 1.09
	pointsOfFunctionPlot[109].Y = 0.462

	pointsOfFunctionPlot[110].X = 1.1
	pointsOfFunctionPlot[110].Y = 0.453

	pointsOfFunctionPlot[111].X = 1.11
	pointsOfFunctionPlot[111].Y = 0.444

	pointsOfFunctionPlot[112].X = 1.12
	pointsOfFunctionPlot[112].Y = 0.435

	pointsOfFunctionPlot[113].X = 1.13
	pointsOfFunctionPlot[113].Y = 0.426

	pointsOfFunctionPlot[114].X = 1.14
	pointsOfFunctionPlot[114].Y = 0.417

	pointsOfFunctionPlot[115].X = 1.15
	pointsOfFunctionPlot[115].Y = 0.408

	pointsOfFunctionPlot[116].X = 1.16
	pointsOfFunctionPlot[116].Y = 0.399

	pointsOfFunctionPlot[117].X = 1.17
	pointsOfFunctionPlot[117].Y = 0.39

	pointsOfFunctionPlot[118].X = 1.18
	pointsOfFunctionPlot[118].Y = 0.38

	pointsOfFunctionPlot[119].X = 1.19
	pointsOfFunctionPlot[119].Y = 0.371

	pointsOfFunctionPlot[120].X = 1.2
	pointsOfFunctionPlot[120].Y = 0.362

	pointsOfFunctionPlot[121].X = 1.21
	pointsOfFunctionPlot[121].Y = 0.353

	pointsOfFunctionPlot[122].X = 1.22
	pointsOfFunctionPlot[122].Y = 0.343

	pointsOfFunctionPlot[123].X = 1.23
	pointsOfFunctionPlot[123].Y = 0.334

	pointsOfFunctionPlot[124].X = 1.24
	pointsOfFunctionPlot[124].Y = 0.324

	pointsOfFunctionPlot[125].X = 1.25
	pointsOfFunctionPlot[125].Y = 0.315

	pointsOfFunctionPlot[126].X = 1.26
	pointsOfFunctionPlot[126].Y = 0.305

	pointsOfFunctionPlot[127].X = 1.27
	pointsOfFunctionPlot[127].Y = 0.296

	pointsOfFunctionPlot[128].X = 1.28
	pointsOfFunctionPlot[128].Y = 0.286

	pointsOfFunctionPlot[129].X = 1.29
	pointsOfFunctionPlot[129].Y = 0.277

	pointsOfFunctionPlot[130].X = 1.3
	pointsOfFunctionPlot[130].Y = 0.267

	pointsOfFunctionPlot[131].X = 1.31
	pointsOfFunctionPlot[131].Y = 0.257

	pointsOfFunctionPlot[132].X = 1.32
	pointsOfFunctionPlot[132].Y = 0.248

	pointsOfFunctionPlot[133].X = 1.33
	pointsOfFunctionPlot[133].Y = 0.238

	pointsOfFunctionPlot[134].X = 1.34
	pointsOfFunctionPlot[134].Y = 0.228

	pointsOfFunctionPlot[135].X = 1.35
	pointsOfFunctionPlot[135].Y = 0.219

	pointsOfFunctionPlot[136].X = 1.36
	pointsOfFunctionPlot[136].Y = 0.209

	pointsOfFunctionPlot[137].X = 1.37
	pointsOfFunctionPlot[137].Y = 0.199

	pointsOfFunctionPlot[138].X = 1.38
	pointsOfFunctionPlot[138].Y = 0.189

	pointsOfFunctionPlot[139].X = 1.39
	pointsOfFunctionPlot[139].Y = 0.179

	pointsOfFunctionPlot[140].X = 1.4
	pointsOfFunctionPlot[140].Y = 0.17

	pointsOfFunctionPlot[141].X = 1.41
	pointsOfFunctionPlot[141].Y = 0.16

	pointsOfFunctionPlot[142].X = 1.42
	pointsOfFunctionPlot[142].Y = 0.15

	pointsOfFunctionPlot[143].X = 1.43
	pointsOfFunctionPlot[143].Y = 0.14

	pointsOfFunctionPlot[144].X = 1.44
	pointsOfFunctionPlot[144].Y = 0.13

	pointsOfFunctionPlot[145].X = 1.45
	pointsOfFunctionPlot[145].Y = 0.12

	pointsOfFunctionPlot[146].X = 1.46
	pointsOfFunctionPlot[146].Y = 0.11

	pointsOfFunctionPlot[147].X = 1.47
	pointsOfFunctionPlot[147].Y = 0.1

	pointsOfFunctionPlot[148].X = 1.48
	pointsOfFunctionPlot[148].Y = 0.09

	pointsOfFunctionPlot[149].X = 1.49
	pointsOfFunctionPlot[149].Y = 0.08

	pointsOfFunctionPlot[150].X = 1.5
	pointsOfFunctionPlot[150].Y = 0.07

	pointsOfFunctionPlot[151].X = 1.51
	pointsOfFunctionPlot[151].Y = 0.06

	pointsOfFunctionPlot[152].X = 1.52
	pointsOfFunctionPlot[152].Y = 0.05

	pointsOfFunctionPlot[153].X = 1.53
	pointsOfFunctionPlot[153].Y = 0.04

	pointsOfFunctionPlot[154].X = 1.54
	pointsOfFunctionPlot[154].Y = 0.03

	pointsOfFunctionPlot[155].X = 1.55
	pointsOfFunctionPlot[155].Y = 0.02

	pointsOfFunctionPlot[156].X = 1.56
	pointsOfFunctionPlot[156].Y = 0.01

	pointsOfFunctionPlot[157].X = 1.57
	pointsOfFunctionPlot[157].Y = 0.0

	pointsOfFunctionPlot[158].X = 1.58
	pointsOfFunctionPlot[158].Y = -0.009

	pointsOfFunctionPlot[159].X = 1.59
	pointsOfFunctionPlot[159].Y = -0.019

	pointsOfFunctionPlot[160].X = 1.6
	pointsOfFunctionPlot[160].Y = -0.029

	pointsOfFunctionPlot[161].X = 1.61
	pointsOfFunctionPlot[161].Y = -0.039

	pointsOfFunctionPlot[162].X = 1.62
	pointsOfFunctionPlot[162].Y = -0.049

	pointsOfFunctionPlot[163].X = 1.63
	pointsOfFunctionPlot[163].Y = -0.059

	pointsOfFunctionPlot[164].X = 1.64
	pointsOfFunctionPlot[164].Y = -0.069

	pointsOfFunctionPlot[165].X = 1.65
	pointsOfFunctionPlot[165].Y = -0.079

	pointsOfFunctionPlot[166].X = 1.66
	pointsOfFunctionPlot[166].Y = -0.089

	pointsOfFunctionPlot[167].X = 1.67
	pointsOfFunctionPlot[167].Y = -0.099

	pointsOfFunctionPlot[168].X = 1.68
	pointsOfFunctionPlot[168].Y = -0.109

	pointsOfFunctionPlot[169].X = 1.69
	pointsOfFunctionPlot[169].Y = -0.118

	pointsOfFunctionPlot[170].X = 1.7
	pointsOfFunctionPlot[170].Y = -0.128

	pointsOfFunctionPlot[171].X = 1.71
	pointsOfFunctionPlot[171].Y = -0.138

	pointsOfFunctionPlot[172].X = 1.72
	pointsOfFunctionPlot[172].Y = -0.148

	pointsOfFunctionPlot[173].X = 1.73
	pointsOfFunctionPlot[173].Y = -0.158

	pointsOfFunctionPlot[174].X = 1.74
	pointsOfFunctionPlot[174].Y = -0.168

	pointsOfFunctionPlot[175].X = 1.75
	pointsOfFunctionPlot[175].Y = -0.178

	pointsOfFunctionPlot[176].X = 1.76
	pointsOfFunctionPlot[176].Y = -0.188

	pointsOfFunctionPlot[177].X = 1.77
	pointsOfFunctionPlot[177].Y = -0.197

	pointsOfFunctionPlot[178].X = 1.78
	pointsOfFunctionPlot[178].Y = -0.207

	pointsOfFunctionPlot[179].X = 1.79
	pointsOfFunctionPlot[179].Y = -0.217

	pointsOfFunctionPlot[180].X = 1.8
	pointsOfFunctionPlot[180].Y = -0.227

	pointsOfFunctionPlot[181].X = 1.81
	pointsOfFunctionPlot[181].Y = -0.236

	pointsOfFunctionPlot[182].X = 1.82
	pointsOfFunctionPlot[182].Y = -0.246

	pointsOfFunctionPlot[183].X = 1.83
	pointsOfFunctionPlot[183].Y = -0.256

	pointsOfFunctionPlot[184].X = 1.84
	pointsOfFunctionPlot[184].Y = -0.266

	pointsOfFunctionPlot[185].X = 1.85
	pointsOfFunctionPlot[185].Y = -0.275

	pointsOfFunctionPlot[186].X = 1.86
	pointsOfFunctionPlot[186].Y = -0.285

	pointsOfFunctionPlot[187].X = 1.87
	pointsOfFunctionPlot[187].Y = -0.294

	pointsOfFunctionPlot[188].X = 1.88
	pointsOfFunctionPlot[188].Y = -0.304

	pointsOfFunctionPlot[189].X = 1.89
	pointsOfFunctionPlot[189].Y = -0.313

	pointsOfFunctionPlot[190].X = 1.9
	pointsOfFunctionPlot[190].Y = -0.323

	pointsOfFunctionPlot[191].X = 1.91
	pointsOfFunctionPlot[191].Y = -0.332

	pointsOfFunctionPlot[192].X = 1.92
	pointsOfFunctionPlot[192].Y = -0.342

	pointsOfFunctionPlot[193].X = 1.93
	pointsOfFunctionPlot[193].Y = -0.351

	pointsOfFunctionPlot[194].X = 1.94
	pointsOfFunctionPlot[194].Y = -0.36

	pointsOfFunctionPlot[195].X = 1.95
	pointsOfFunctionPlot[195].Y = -0.37

	pointsOfFunctionPlot[196].X = 1.96
	pointsOfFunctionPlot[196].Y = -0.379

	pointsOfFunctionPlot[197].X = 1.97
	pointsOfFunctionPlot[197].Y = -0.388

	pointsOfFunctionPlot[198].X = 1.98
	pointsOfFunctionPlot[198].Y = -0.397

	pointsOfFunctionPlot[199].X = 1.99
	pointsOfFunctionPlot[199].Y = -0.407

	pointsOfFunctionPlot[200].X = 2.0
	pointsOfFunctionPlot[200].Y = -0.416

	pointsOfFunctionPlot[201].X = 2.01
	pointsOfFunctionPlot[201].Y = -0.425

	pointsOfFunctionPlot[202].X = 2.02
	pointsOfFunctionPlot[202].Y = -0.434

	pointsOfFunctionPlot[203].X = 2.03
	pointsOfFunctionPlot[203].Y = -0.443

	pointsOfFunctionPlot[204].X = 2.04
	pointsOfFunctionPlot[204].Y = -0.452

	pointsOfFunctionPlot[205].X = 2.05
	pointsOfFunctionPlot[205].Y = -0.461

	pointsOfFunctionPlot[206].X = 2.06
	pointsOfFunctionPlot[206].Y = -0.469

	pointsOfFunctionPlot[207].X = 2.07
	pointsOfFunctionPlot[207].Y = -0.478

	pointsOfFunctionPlot[208].X = 2.08
	pointsOfFunctionPlot[208].Y = -0.487

	pointsOfFunctionPlot[209].X = 2.09
	pointsOfFunctionPlot[209].Y = -0.496

	pointsOfFunctionPlot[210].X = 2.1
	pointsOfFunctionPlot[210].Y = -0.504

	pointsOfFunctionPlot[211].X = 2.11
	pointsOfFunctionPlot[211].Y = -0.513

	pointsOfFunctionPlot[212].X = 2.12
	pointsOfFunctionPlot[212].Y = -0.522

	pointsOfFunctionPlot[213].X = 2.13
	pointsOfFunctionPlot[213].Y = -0.53

	pointsOfFunctionPlot[214].X = 2.14
	pointsOfFunctionPlot[214].Y = -0.539

	pointsOfFunctionPlot[215].X = 2.15
	pointsOfFunctionPlot[215].Y = -0.547

	pointsOfFunctionPlot[216].X = 2.16
	pointsOfFunctionPlot[216].Y = -0.555

	pointsOfFunctionPlot[217].X = 2.17
	pointsOfFunctionPlot[217].Y = -0.564

	pointsOfFunctionPlot[218].X = 2.18
	pointsOfFunctionPlot[218].Y = -0.572

	pointsOfFunctionPlot[219].X = 2.19
	pointsOfFunctionPlot[219].Y = -0.58

	pointsOfFunctionPlot[220].X = 2.2
	pointsOfFunctionPlot[220].Y = -0.588

	pointsOfFunctionPlot[221].X = 2.21
	pointsOfFunctionPlot[221].Y = -0.596

	pointsOfFunctionPlot[222].X = 2.22
	pointsOfFunctionPlot[222].Y = -0.604

	pointsOfFunctionPlot[223].X = 2.23
	pointsOfFunctionPlot[223].Y = -0.612

	pointsOfFunctionPlot[224].X = 2.24
	pointsOfFunctionPlot[224].Y = -0.62

	pointsOfFunctionPlot[225].X = 2.25
	pointsOfFunctionPlot[225].Y = -0.628

	pointsOfFunctionPlot[226].X = 2.26
	pointsOfFunctionPlot[226].Y = -0.635

	pointsOfFunctionPlot[227].X = 2.27
	pointsOfFunctionPlot[227].Y = -0.643

	pointsOfFunctionPlot[228].X = 2.28
	pointsOfFunctionPlot[228].Y = -0.651

	pointsOfFunctionPlot[229].X = 2.29
	pointsOfFunctionPlot[229].Y = -0.658

	pointsOfFunctionPlot[230].X = 2.3
	pointsOfFunctionPlot[230].Y = -0.666

	pointsOfFunctionPlot[231].X = 2.31
	pointsOfFunctionPlot[231].Y = -0.673

	pointsOfFunctionPlot[232].X = 2.32
	pointsOfFunctionPlot[232].Y = -0.681

	pointsOfFunctionPlot[233].X = 2.33
	pointsOfFunctionPlot[233].Y = -0.688

	pointsOfFunctionPlot[234].X = 2.34
	pointsOfFunctionPlot[234].Y = -0.695

	pointsOfFunctionPlot[235].X = 2.35
	pointsOfFunctionPlot[235].Y = -0.702

	pointsOfFunctionPlot[236].X = 2.36
	pointsOfFunctionPlot[236].Y = -0.709

	pointsOfFunctionPlot[237].X = 2.37
	pointsOfFunctionPlot[237].Y = -0.716

	pointsOfFunctionPlot[238].X = 2.38
	pointsOfFunctionPlot[238].Y = -0.723

	pointsOfFunctionPlot[239].X = 2.39
	pointsOfFunctionPlot[239].Y = -0.73

	pointsOfFunctionPlot[240].X = 2.4
	pointsOfFunctionPlot[240].Y = -0.737

	pointsOfFunctionPlot[241].X = 2.41
	pointsOfFunctionPlot[241].Y = -0.744

	pointsOfFunctionPlot[242].X = 2.42
	pointsOfFunctionPlot[242].Y = -0.75

	pointsOfFunctionPlot[243].X = 2.43
	pointsOfFunctionPlot[243].Y = -0.757

	pointsOfFunctionPlot[244].X = 2.44
	pointsOfFunctionPlot[244].Y = -0.763

	pointsOfFunctionPlot[245].X = 2.45
	pointsOfFunctionPlot[245].Y = -0.77

	pointsOfFunctionPlot[246].X = 2.46
	pointsOfFunctionPlot[246].Y = -0.776

	pointsOfFunctionPlot[247].X = 2.47
	pointsOfFunctionPlot[247].Y = -0.782

	pointsOfFunctionPlot[248].X = 2.48
	pointsOfFunctionPlot[248].Y = -0.789

	pointsOfFunctionPlot[249].X = 2.49
	pointsOfFunctionPlot[249].Y = -0.795

	pointsOfFunctionPlot[250].X = 2.5
	pointsOfFunctionPlot[250].Y = -0.801

	pointsOfFunctionPlot[251].X = 2.51
	pointsOfFunctionPlot[251].Y = -0.807

	pointsOfFunctionPlot[252].X = 2.52
	pointsOfFunctionPlot[252].Y = -0.813

	pointsOfFunctionPlot[253].X = 2.53
	pointsOfFunctionPlot[253].Y = -0.818

	pointsOfFunctionPlot[254].X = 2.54
	pointsOfFunctionPlot[254].Y = -0.824

	pointsOfFunctionPlot[255].X = 2.55
	pointsOfFunctionPlot[255].Y = -0.83

	pointsOfFunctionPlot[256].X = 2.56
	pointsOfFunctionPlot[256].Y = -0.835

	pointsOfFunctionPlot[257].X = 2.57
	pointsOfFunctionPlot[257].Y = -0.841

	pointsOfFunctionPlot[258].X = 2.58
	pointsOfFunctionPlot[258].Y = -0.846

	pointsOfFunctionPlot[259].X = 2.59
	pointsOfFunctionPlot[259].Y = -0.851

	pointsOfFunctionPlot[260].X = 2.6
	pointsOfFunctionPlot[260].Y = -0.856

	pointsOfFunctionPlot[261].X = 2.61
	pointsOfFunctionPlot[261].Y = -0.862

	pointsOfFunctionPlot[262].X = 2.62
	pointsOfFunctionPlot[262].Y = -0.867

	pointsOfFunctionPlot[263].X = 2.63
	pointsOfFunctionPlot[263].Y = -0.872

	pointsOfFunctionPlot[264].X = 2.64
	pointsOfFunctionPlot[264].Y = -0.876

	pointsOfFunctionPlot[265].X = 2.65
	pointsOfFunctionPlot[265].Y = -0.881

	pointsOfFunctionPlot[266].X = 2.66
	pointsOfFunctionPlot[266].Y = -0.886

	pointsOfFunctionPlot[267].X = 2.67
	pointsOfFunctionPlot[267].Y = -0.89

	pointsOfFunctionPlot[268].X = 2.68
	pointsOfFunctionPlot[268].Y = -0.895

	pointsOfFunctionPlot[269].X = 2.69
	pointsOfFunctionPlot[269].Y = -0.899

	pointsOfFunctionPlot[270].X = 2.7
	pointsOfFunctionPlot[270].Y = -0.904

	pointsOfFunctionPlot[271].X = 2.71
	pointsOfFunctionPlot[271].Y = -0.908

	pointsOfFunctionPlot[272].X = 2.72
	pointsOfFunctionPlot[272].Y = -0.912

	pointsOfFunctionPlot[273].X = 2.73
	pointsOfFunctionPlot[273].Y = -0.916

	pointsOfFunctionPlot[274].X = 2.74
	pointsOfFunctionPlot[274].Y = -0.92

	pointsOfFunctionPlot[275].X = 2.75
	pointsOfFunctionPlot[275].Y = -0.924

	pointsOfFunctionPlot[276].X = 2.76
	pointsOfFunctionPlot[276].Y = -0.928

	pointsOfFunctionPlot[277].X = 2.77
	pointsOfFunctionPlot[277].Y = -0.931

	pointsOfFunctionPlot[278].X = 2.78
	pointsOfFunctionPlot[278].Y = -0.935

	pointsOfFunctionPlot[279].X = 2.79
	pointsOfFunctionPlot[279].Y = -0.938

	pointsOfFunctionPlot[280].X = 2.8
	pointsOfFunctionPlot[280].Y = -0.942

	pointsOfFunctionPlot[281].X = 2.81
	pointsOfFunctionPlot[281].Y = -0.945

	pointsOfFunctionPlot[282].X = 2.82
	pointsOfFunctionPlot[282].Y = -0.948

	pointsOfFunctionPlot[283].X = 2.83
	pointsOfFunctionPlot[283].Y = -0.951

	pointsOfFunctionPlot[284].X = 2.84
	pointsOfFunctionPlot[284].Y = -0.954

	pointsOfFunctionPlot[285].X = 2.85
	pointsOfFunctionPlot[285].Y = -0.957

	pointsOfFunctionPlot[286].X = 2.86
	pointsOfFunctionPlot[286].Y = -0.96

	pointsOfFunctionPlot[287].X = 2.87
	pointsOfFunctionPlot[287].Y = -0.963

	pointsOfFunctionPlot[288].X = 2.88
	pointsOfFunctionPlot[288].Y = -0.966

	pointsOfFunctionPlot[289].X = 2.89
	pointsOfFunctionPlot[289].Y = -0.968

	pointsOfFunctionPlot[290].X = 2.9
	pointsOfFunctionPlot[290].Y = -0.971

	pointsOfFunctionPlot[291].X = 2.91
	pointsOfFunctionPlot[291].Y = -0.973

	pointsOfFunctionPlot[292].X = 2.92
	pointsOfFunctionPlot[292].Y = -0.975

	pointsOfFunctionPlot[293].X = 2.93
	pointsOfFunctionPlot[293].Y = -0.977

	pointsOfFunctionPlot[294].X = 2.94
	pointsOfFunctionPlot[294].Y = -0.979

	pointsOfFunctionPlot[295].X = 2.95
	pointsOfFunctionPlot[295].Y = -0.981

	pointsOfFunctionPlot[296].X = 2.96
	pointsOfFunctionPlot[296].Y = -0.983

	pointsOfFunctionPlot[297].X = 2.97
	pointsOfFunctionPlot[297].Y = -0.985

	pointsOfFunctionPlot[298].X = 2.98
	pointsOfFunctionPlot[298].Y = -0.987

	pointsOfFunctionPlot[299].X = 2.99
	pointsOfFunctionPlot[299].Y = -0.988

	pointsOfFunctionPlot[300].X = 3.0
	pointsOfFunctionPlot[300].Y = -0.99

	pointsOfFunctionPlot[301].X = 3.01
	pointsOfFunctionPlot[301].Y = -0.991

	pointsOfFunctionPlot[302].X = 3.02
	pointsOfFunctionPlot[302].Y = -0.992

	pointsOfFunctionPlot[303].X = 3.03
	pointsOfFunctionPlot[303].Y = -0.993

	pointsOfFunctionPlot[304].X = 3.04
	pointsOfFunctionPlot[304].Y = -0.994

	pointsOfFunctionPlot[305].X = 3.05
	pointsOfFunctionPlot[305].Y = -0.995

	pointsOfFunctionPlot[306].X = 3.06
	pointsOfFunctionPlot[306].Y = -0.996

	pointsOfFunctionPlot[307].X = 3.07
	pointsOfFunctionPlot[307].Y = -0.997

	pointsOfFunctionPlot[308].X = 3.08
	pointsOfFunctionPlot[308].Y = -0.998

	pointsOfFunctionPlot[309].X = 3.09
	pointsOfFunctionPlot[309].Y = -0.998

	pointsOfFunctionPlot[310].X = 3.1
	pointsOfFunctionPlot[310].Y = -0.999

	pointsOfFunctionPlot[311].X = 3.11
	pointsOfFunctionPlot[311].Y = -0.999

	pointsOfFunctionPlot[312].X = 3.12
	pointsOfFunctionPlot[312].Y = -0.999

	pointsOfFunctionPlot[313].X = 3.13
	pointsOfFunctionPlot[313].Y = -0.999

	pointsOfFunctionPlot[314].X = 3.14
	pointsOfFunctionPlot[314].Y = -1.0

	pointsOfFunctionPlot[315].X = 3.15
	pointsOfFunctionPlot[315].Y = -1.0

	pointsOfFunctionPlot[316].X = 3.16
	pointsOfFunctionPlot[316].Y = -0.999

	pointsOfFunctionPlot[317].X = 3.17
	pointsOfFunctionPlot[317].Y = -0.999

	pointsOfFunctionPlot[318].X = 3.18
	pointsOfFunctionPlot[318].Y = -0.999

	pointsOfFunctionPlot[319].X = 3.19
	pointsOfFunctionPlot[319].Y = -0.998

	pointsOfFunctionPlot[320].X = 3.2
	pointsOfFunctionPlot[320].Y = -0.997










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
		"cos-function-plot-01.png"); err != nil {

		panic(err)
	}
}
