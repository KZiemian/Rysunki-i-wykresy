package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function f(x) = exp(x) - ln(x).

	pointsOfFunctionPlot := make(plotter.XYs, 501)

	pointsOfFunctionPlot[0].X = 0.001
	pointsOfFunctionPlot[0].Y = -5.906

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = -3.595

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = -2.891

	pointsOfFunctionPlot[3].X = 0.03
	pointsOfFunctionPlot[3].Y = -2.476

	pointsOfFunctionPlot[4].X = 0.04
	pointsOfFunctionPlot[4].Y = -2.178

	pointsOfFunctionPlot[5].X = 0.05
	pointsOfFunctionPlot[5].Y = -1.944

	pointsOfFunctionPlot[6].X = 0.06
	pointsOfFunctionPlot[6].Y = -1.751

	pointsOfFunctionPlot[7].X = 0.07
	pointsOfFunctionPlot[7].Y = -1.586

	pointsOfFunctionPlot[8].X = 0.08
	pointsOfFunctionPlot[8].Y = -1.442

	pointsOfFunctionPlot[9].X = 0.09
	pointsOfFunctionPlot[9].Y = -1.313

	pointsOfFunctionPlot[10].X = 0.10
	pointsOfFunctionPlot[10].Y = -1.197

	pointsOfFunctionPlot[11].X = 0.11
	pointsOfFunctionPlot[11].Y = -1.09

	pointsOfFunctionPlot[12].X = 0.12
	pointsOfFunctionPlot[12].Y = -0.992

	pointsOfFunctionPlot[13].X = 0.13
	pointsOfFunctionPlot[13].Y = -0.901

	pointsOfFunctionPlot[14].X = 0.14
	pointsOfFunctionPlot[14].Y = -0.815

	pointsOfFunctionPlot[15].X = 0.15
	pointsOfFunctionPlot[15].Y = -0.735

	pointsOfFunctionPlot[16].X = 0.16
	pointsOfFunctionPlot[16].Y = -0.659

	pointsOfFunctionPlot[17].X = 0.17
	pointsOfFunctionPlot[17].Y = -0.586

	pointsOfFunctionPlot[18].X = 0.18
	pointsOfFunctionPlot[18].Y = -0.517

	pointsOfFunctionPlot[19].X = 0.19
	pointsOfFunctionPlot[19].Y = -0.451

	pointsOfFunctionPlot[20].X = 0.20
	pointsOfFunctionPlot[20].Y = -0.388

	pointsOfFunctionPlot[21].X = 0.21
	pointsOfFunctionPlot[21].Y = -0.326

	pointsOfFunctionPlot[22].X = 0.22
	pointsOfFunctionPlot[22].Y = -0.268

	pointsOfFunctionPlot[23].X = 0.23
	pointsOfFunctionPlot[23].Y = -0.211

	pointsOfFunctionPlot[24].X = 0.24
	pointsOfFunctionPlot[24].Y = -0.155

	pointsOfFunctionPlot[25].X = 0.25
	pointsOfFunctionPlot[25].Y = -0.102

	pointsOfFunctionPlot[26].X = 0.26
	pointsOfFunctionPlot[26].Y = -0.05

	pointsOfFunctionPlot[27].X = 0.27
	pointsOfFunctionPlot[27].Y = 0.0

	pointsOfFunctionPlot[28].X = 0.28
	pointsOfFunctionPlot[28].Y = 0.05

	pointsOfFunctionPlot[29].X = 0.29
	pointsOfFunctionPlot[29].Y = 0.098

	pointsOfFunctionPlot[30].X = 0.30
	pointsOfFunctionPlot[30].Y = 0.145

	pointsOfFunctionPlot[31].X = 0.31
	pointsOfFunctionPlot[31].Y = 0.192

	pointsOfFunctionPlot[32].X = 0.32
	pointsOfFunctionPlot[32].Y = 0.237

	pointsOfFunctionPlot[33].X = 0.33
	pointsOfFunctionPlot[33].Y = 0.282

	pointsOfFunctionPlot[34].X = 0.34
	pointsOfFunctionPlot[34].Y = 0.326

	pointsOfFunctionPlot[35].X = 0.35
	pointsOfFunctionPlot[35].Y = 0.369

	pointsOfFunctionPlot[36].X = 0.36
	pointsOfFunctionPlot[36].Y = 0.411

	pointsOfFunctionPlot[37].X = 0.37
	pointsOfFunctionPlot[37].Y = 0.453

	pointsOfFunctionPlot[38].X = 0.38
	pointsOfFunctionPlot[38].Y = 0.494

	pointsOfFunctionPlot[39].X = 0.39
	pointsOfFunctionPlot[39].Y = 0.535

	pointsOfFunctionPlot[40].X = 0.40
	pointsOfFunctionPlot[40].Y = 0.575

	pointsOfFunctionPlot[41].X = 0.41
	pointsOfFunctionPlot[41].Y = 0.615

	pointsOfFunctionPlot[42].X = 0.42
	pointsOfFunctionPlot[42].Y = 0.654

	pointsOfFunctionPlot[43].X = 0.43
	pointsOfFunctionPlot[43].Y = 0.693

	pointsOfFunctionPlot[44].X = 0.44
	pointsOfFunctionPlot[44].Y = 0.731

	pointsOfFunctionPlot[45].X = 0.45
	pointsOfFunctionPlot[45].Y = 0.769

	pointsOfFunctionPlot[46].X = 0.46
	pointsOfFunctionPlot[46].Y = 0.807

	pointsOfFunctionPlot[47].X = 0.47
	pointsOfFunctionPlot[47].Y = 0.844

	pointsOfFunctionPlot[48].X = 0.48
	pointsOfFunctionPlot[48].Y = 0.882

	pointsOfFunctionPlot[49].X = 0.49
	pointsOfFunctionPlot[49].Y = 0.918

	pointsOfFunctionPlot[50].X = 0.50
	pointsOfFunctionPlot[50].Y = 0.955

	pointsOfFunctionPlot[51].X = 0.51
	pointsOfFunctionPlot[51].Y = 0.991

	pointsOfFunctionPlot[52].X = 0.52
	pointsOfFunctionPlot[52].Y = 1.028

	pointsOfFunctionPlot[53].X = 0.53
	pointsOfFunctionPlot[53].Y = 1.064

	pointsOfFunctionPlot[54].X = 0.54
	pointsOfFunctionPlot[54].Y = 1.099

	pointsOfFunctionPlot[55].X = 0.55
	pointsOfFunctionPlot[55].Y = 1.135

	pointsOfFunctionPlot[56].X = 0.56
	pointsOfFunctionPlot[56].Y = 1.17

	pointsOfFunctionPlot[57].X = 0.57
	pointsOfFunctionPlot[57].Y = 1.206

	pointsOfFunctionPlot[58].X = 0.58
	pointsOfFunctionPlot[58].Y = 1.241

	pointsOfFunctionPlot[59].X = 0.59
	pointsOfFunctionPlot[59].Y = 1.276

	pointsOfFunctionPlot[60].X = 0.60
	pointsOfFunctionPlot[60].Y = 1.311

	pointsOfFunctionPlot[61].X = 0.61
	pointsOfFunctionPlot[61].Y = 1.346

	pointsOfFunctionPlot[62].X = 0.62
	pointsOfFunctionPlot[62].Y = 1.38

	pointsOfFunctionPlot[63].X = 0.63
	pointsOfFunctionPlot[63].Y = 1.415

	pointsOfFunctionPlot[64].X = 0.64
	pointsOfFunctionPlot[64].Y = 1.45

	pointsOfFunctionPlot[65].X = 0.65
	pointsOfFunctionPlot[65].Y = 1.484

	pointsOfFunctionPlot[66].X = 0.66
	pointsOfFunctionPlot[66].Y = 1.519

	pointsOfFunctionPlot[67].X = 0.67
	pointsOfFunctionPlot[67].Y = 1.553

	pointsOfFunctionPlot[68].X = 0.68
	pointsOfFunctionPlot[68].Y = 1.588

	pointsOfFunctionPlot[69].X = 0.69
	pointsOfFunctionPlot[69].Y = 1.622

	pointsOfFunctionPlot[70].X = 0.70
	pointsOfFunctionPlot[70].Y = 1.657

	pointsOfFunctionPlot[71].X = 0.71
	pointsOfFunctionPlot[71].Y = 1.691

	pointsOfFunctionPlot[72].X = 0.72
	pointsOfFunctionPlot[72].Y = 1.725

	pointsOfFunctionPlot[73].X = 0.73
	pointsOfFunctionPlot[73].Y = 1.76

	pointsOfFunctionPlot[74].X = 0.74
	pointsOfFunctionPlot[74].Y = 1.794

	pointsOfFunctionPlot[75].X = 0.75
	pointsOfFunctionPlot[75].Y = 1.829

	pointsOfFunctionPlot[76].X = 0.76
	pointsOfFunctionPlot[76].Y = 1.863

	pointsOfFunctionPlot[77].X = 0.77
	pointsOfFunctionPlot[77].Y = 1.898

	pointsOfFunctionPlot[78].X = 0.78
	pointsOfFunctionPlot[78].Y = 1.933

	pointsOfFunctionPlot[79].X = 0.79
	pointsOfFunctionPlot[79].Y = 1.967

	pointsOfFunctionPlot[80].X = 0.80
	pointsOfFunctionPlot[80].Y = 2.002

	pointsOfFunctionPlot[81].X = 0.81
	pointsOfFunctionPlot[81].Y = 2.037

	pointsOfFunctionPlot[82].X = 0.82
	pointsOfFunctionPlot[82].Y = 2.072

	pointsOfFunctionPlot[83].X = 0.83
	pointsOfFunctionPlot[83].Y = 2.106

	pointsOfFunctionPlot[84].X = 0.84
	pointsOfFunctionPlot[84].Y = 2.142

	pointsOfFunctionPlot[85].X = 0.85
	pointsOfFunctionPlot[85].Y = 2.177

	pointsOfFunctionPlot[86].X = 0.86
	pointsOfFunctionPlot[86].Y = 2.212

	pointsOfFunctionPlot[87].X = 0.87
	pointsOfFunctionPlot[87].Y = 2.247

	pointsOfFunctionPlot[88].X = 0.88
	pointsOfFunctionPlot[88].Y = 2.283

	pointsOfFunctionPlot[89].X = 0.89
	pointsOfFunctionPlot[89].Y = 2.318

	pointsOfFunctionPlot[90].X = 0.90
	pointsOfFunctionPlot[90].Y = 2.354

	pointsOfFunctionPlot[91].X = 0.91
	pointsOfFunctionPlot[91].Y = 2.39

	pointsOfFunctionPlot[92].X = 0.92
	pointsOfFunctionPlot[92].Y = 2.425

	pointsOfFunctionPlot[93].X = 0.93
	pointsOfFunctionPlot[93].Y = 2.461

	pointsOfFunctionPlot[94].X = 0.94
	pointsOfFunctionPlot[94].Y = 2.498

	pointsOfFunctionPlot[95].X = 0.95
	pointsOfFunctionPlot[95].Y = 2.534

	pointsOfFunctionPlot[96].X = 0.96
	pointsOfFunctionPlot[96].Y = 2.57

	pointsOfFunctionPlot[97].X = 0.97
	pointsOfFunctionPlot[97].Y = 2.607

	pointsOfFunctionPlot[98].X = 0.98
	pointsOfFunctionPlot[98].Y = 2.644

	pointsOfFunctionPlot[99].X = 0.99
	pointsOfFunctionPlot[99].Y = 2.681

	pointsOfFunctionPlot[100].X = 1.0
	pointsOfFunctionPlot[100].Y = 2.718

	pointsOfFunctionPlot[101].X = 1.01
	pointsOfFunctionPlot[101].Y = 2.755

	pointsOfFunctionPlot[102].X = 1.02
	pointsOfFunctionPlot[102].Y = 2.792

	pointsOfFunctionPlot[103].X = 1.03
	pointsOfFunctionPlot[103].Y = 2.83

	pointsOfFunctionPlot[104].X = 1.04
	pointsOfFunctionPlot[104].Y = 2.868

	pointsOfFunctionPlot[105].X = 1.05
	pointsOfFunctionPlot[105].Y = 2.906

	pointsOfFunctionPlot[106].X = 1.06
	pointsOfFunctionPlot[106].Y = 2.944

	pointsOfFunctionPlot[107].X = 1.07
	pointsOfFunctionPlot[107].Y = 2.983

	pointsOfFunctionPlot[108].X = 1.08
	pointsOfFunctionPlot[108].Y = 3.021

	pointsOfFunctionPlot[109].X = 1.09
	pointsOfFunctionPlot[109].Y = 3.06

	pointsOfFunctionPlot[110].X = 1.10
	pointsOfFunctionPlot[110].Y = 3.099

	pointsOfFunctionPlot[111].X = 1.11
	pointsOfFunctionPlot[111].Y = 3.138

	pointsOfFunctionPlot[112].X = 1.12
	pointsOfFunctionPlot[112].Y = 3.178

	pointsOfFunctionPlot[113].X = 1.13
	pointsOfFunctionPlot[113].Y = 3.217

	pointsOfFunctionPlot[114].X = 1.14
	pointsOfFunctionPlot[114].Y = 3.257

	pointsOfFunctionPlot[115].X = 1.15
	pointsOfFunctionPlot[115].Y = 3.297

	pointsOfFunctionPlot[116].X = 1.16
	pointsOfFunctionPlot[116].Y = 3.338

	pointsOfFunctionPlot[117].X = 1.17
	pointsOfFunctionPlot[117].Y = 3.378

	pointsOfFunctionPlot[118].X = 1.18
	pointsOfFunctionPlot[118].Y = 3.419

	pointsOfFunctionPlot[119].X = 1.19
	pointsOfFunctionPlot[119].Y = 3.461

	pointsOfFunctionPlot[120].X = 1.20
	pointsOfFunctionPlot[120].Y = 3.502

	pointsOfFunctionPlot[121].X = 1.21
	pointsOfFunctionPlot[121].Y = 3.544

	pointsOfFunctionPlot[122].X = 1.22
	pointsOfFunctionPlot[122].Y = 3.586

	pointsOfFunctionPlot[123].X = 1.23
	pointsOfFunctionPlot[123].Y = 3.628

	pointsOfFunctionPlot[124].X = 1.24
	pointsOfFunctionPlot[124].Y = 3.67

	pointsOfFunctionPlot[125].X = 1.25
	pointsOfFunctionPlot[125].Y = 3.713

	pointsOfFunctionPlot[126].X = 1.26
	pointsOfFunctionPlot[126].Y = 3.756

	pointsOfFunctionPlot[127].X = 1.27
	pointsOfFunctionPlot[127].Y = 3.799

	pointsOfFunctionPlot[128].X = 1.28
	pointsOfFunctionPlot[128].Y = 3.843

	pointsOfFunctionPlot[129].X = 1.29
	pointsOfFunctionPlot[129].Y = 3.887

	pointsOfFunctionPlot[130].X = 1.30
	pointsOfFunctionPlot[130].Y = 3.931

	pointsOfFunctionPlot[131].X = 1.31
	pointsOfFunctionPlot[131].Y = 3.976

	pointsOfFunctionPlot[132].X = 1.32
	pointsOfFunctionPlot[132].Y = 4.021

	pointsOfFunctionPlot[133].X = 1.33
	pointsOfFunctionPlot[133].Y = 4.066

	pointsOfFunctionPlot[134].X = 1.34
	pointsOfFunctionPlot[134].Y = 4.111

	pointsOfFunctionPlot[135].X = 1.35
	pointsOfFunctionPlot[135].Y = 4.157

	pointsOfFunctionPlot[136].X = 1.36
	pointsOfFunctionPlot[136].Y = 4.203

	pointsOfFunctionPlot[137].X = 1.37
	pointsOfFunctionPlot[137].Y = 4.25

	pointsOfFunctionPlot[138].X = 1.38
	pointsOfFunctionPlot[138].Y = 4.296

	pointsOfFunctionPlot[139].X = 1.39
	pointsOfFunctionPlot[139].Y = 4.344

	pointsOfFunctionPlot[140].X = 1.40
	pointsOfFunctionPlot[140].Y = 4.391

	pointsOfFunctionPlot[141].X = 1.41
	pointsOfFunctionPlot[141].Y = 4.439

	pointsOfFunctionPlot[142].X = 1.42
	pointsOfFunctionPlot[142].Y = 4.487

	pointsOfFunctionPlot[143].X = 1.43
	pointsOfFunctionPlot[143].Y = 4.536

	pointsOfFunctionPlot[144].X = 1.44
	pointsOfFunctionPlot[144].Y = 4.585

	pointsOfFunctionPlot[145].X = 1.45
	pointsOfFunctionPlot[145].Y = 4.634

	pointsOfFunctionPlot[146].X = 1.46
	pointsOfFunctionPlot[146].Y = 4.684

	pointsOfFunctionPlot[147].X = 1.47
	pointsOfFunctionPlot[147].Y = 4.734

	pointsOfFunctionPlot[148].X = 1.48
	pointsOfFunctionPlot[148].Y = 4.784

	pointsOfFunctionPlot[149].X = 1.49
	pointsOfFunctionPlot[149].Y = 4.835

	pointsOfFunctionPlot[150].X = 1.50
	pointsOfFunctionPlot[150].Y = 4.887

	pointsOfFunctionPlot[151].X = 1.51
	pointsOfFunctionPlot[151].Y = 4.938

	pointsOfFunctionPlot[152].X = 1.52
	pointsOfFunctionPlot[152].Y = 4.99

	pointsOfFunctionPlot[153].X = 1.53
	pointsOfFunctionPlot[153].Y = 5.043

	pointsOfFunctionPlot[154].X = 1.54
	pointsOfFunctionPlot[154].Y = 5.096

	pointsOfFunctionPlot[155].X = 1.55
	pointsOfFunctionPlot[155].Y = 5.149

	pointsOfFunctionPlot[156].X = 1.56
	pointsOfFunctionPlot[156].Y = 5.203

	pointsOfFunctionPlot[157].X = 1.57
	pointsOfFunctionPlot[157].Y = 5.257

	pointsOfFunctionPlot[158].X = 1.58
	pointsOfFunctionPlot[158].Y = 5.312

	pointsOfFunctionPlot[159].X = 1.59
	pointsOfFunctionPlot[159].Y = 5.367

	pointsOfFunctionPlot[160].X = 1.60
	pointsOfFunctionPlot[160].Y = 5.423

	pointsOfFunctionPlot[161].X = 1.61
	pointsOfFunctionPlot[161].Y = 5.479

	pointsOfFunctionPlot[162].X = 1.62
	pointsOfFunctionPlot[162].Y = 5.535

	pointsOfFunctionPlot[163].X = 1.63
	pointsOfFunctionPlot[163].Y = 5.592

	pointsOfFunctionPlot[164].X = 1.64
	pointsOfFunctionPlot[164].Y = 5.649

	pointsOfFunctionPlot[165].X = 1.65
	pointsOfFunctionPlot[165].Y = 5.707

	pointsOfFunctionPlot[166].X = 1.66
	pointsOfFunctionPlot[166].Y = 5.766

	pointsOfFunctionPlot[167].X = 1.67
	pointsOfFunctionPlot[167].Y = 5.824

	pointsOfFunctionPlot[168].X = 1.68
	pointsOfFunctionPlot[168].Y = 5.884

	pointsOfFunctionPlot[169].X = 1.69
	pointsOfFunctionPlot[169].Y = 5.944

	pointsOfFunctionPlot[170].X = 1.70
	pointsOfFunctionPlot[170].Y = 6.004

	pointsOfFunctionPlot[171].X = 1.71
	pointsOfFunctionPlot[171].Y = 6.065

	pointsOfFunctionPlot[172].X = 1.72
	pointsOfFunctionPlot[172].Y = 6.126

	pointsOfFunctionPlot[173].X = 1.73
	pointsOfFunctionPlot[173].Y = 6.188

	pointsOfFunctionPlot[174].X = 1.74
	pointsOfFunctionPlot[174].Y = 6.251

	pointsOfFunctionPlot[175].X = 1.75
	pointsOfFunctionPlot[175].Y = 6.314

	pointsOfFunctionPlot[176].X = 1.76
	pointsOfFunctionPlot[176].Y = 6.377

	pointsOfFunctionPlot[177].X = 1.77
	pointsOfFunctionPlot[177].Y = 6.441

	pointsOfFunctionPlot[178].X = 1.78
	pointsOfFunctionPlot[178].Y = 6.506

	pointsOfFunctionPlot[179].X = 1.79
	pointsOfFunctionPlot[179].Y = 6.571

	pointsOfFunctionPlot[180].X = 1.80
	pointsOfFunctionPlot[180].Y = 6.637

	pointsOfFunctionPlot[181].X = 1.81
	pointsOfFunctionPlot[181].Y = 6.703

	pointsOfFunctionPlot[182].X = 1.82
	pointsOfFunctionPlot[182].Y = 6.77

	pointsOfFunctionPlot[183].X = 1.83
	pointsOfFunctionPlot[183].Y = 6.838

	pointsOfFunctionPlot[184].X = 1.84
	pointsOfFunctionPlot[184].Y = 6.906

	pointsOfFunctionPlot[185].X = 1.85
	pointsOfFunctionPlot[185].Y = 6.975

	pointsOfFunctionPlot[186].X = 1.86
	pointsOfFunctionPlot[186].Y = 7.044

	pointsOfFunctionPlot[187].X = 1.87
	pointsOfFunctionPlot[187].Y = 7.114

	pointsOfFunctionPlot[188].X = 1.88
	pointsOfFunctionPlot[188].Y = 7.184

	pointsOfFunctionPlot[189].X = 1.89
	pointsOfFunctionPlot[189].Y = 7.255

	pointsOfFunctionPlot[190].X = 1.90
	pointsOfFunctionPlot[190].Y = 7.327

	pointsOfFunctionPlot[191].X = 1.91
	pointsOfFunctionPlot[191].Y = 7.4

	pointsOfFunctionPlot[192].X = 1.92
	pointsOfFunctionPlot[192].Y = 7.473

	pointsOfFunctionPlot[193].X = 1.93
	pointsOfFunctionPlot[193].Y = 7.547

	pointsOfFunctionPlot[194].X = 1.94
	pointsOfFunctionPlot[194].Y = 7.621

	pointsOfFunctionPlot[195].X = 1.95
	pointsOfFunctionPlot[195].Y = 7.696

	pointsOfFunctionPlot[196].X = 1.96
	pointsOfFunctionPlot[196].Y = 7.772

	pointsOfFunctionPlot[197].X = 1.97
	pointsOfFunctionPlot[197].Y = 7.848

	pointsOfFunctionPlot[198].X = 1.98
	pointsOfFunctionPlot[198].Y = 7.925

	pointsOfFunctionPlot[199].X = 1.99
	pointsOfFunctionPlot[199].Y = 8.003

	pointsOfFunctionPlot[200].X = 2.0
	pointsOfFunctionPlot[200].Y = 8.082

	pointsOfFunctionPlot[201].X = 2.01
	pointsOfFunctionPlot[201].Y = 8.161

	pointsOfFunctionPlot[202].X = 2.02
	pointsOfFunctionPlot[202].Y = 8.241

	pointsOfFunctionPlot[203].X = 2.03
	pointsOfFunctionPlot[203].Y = 8.322

	pointsOfFunctionPlot[204].X = 2.04
	pointsOfFunctionPlot[204].Y = 8.403

	pointsOfFunctionPlot[205].X = 2.05
	pointsOfFunctionPlot[205].Y = 8.485

	pointsOfFunctionPlot[206].X = 2.06
	pointsOfFunctionPlot[206].Y = 8.568

	pointsOfFunctionPlot[207].X = 2.07
	pointsOfFunctionPlot[207].Y = 8.652

	pointsOfFunctionPlot[208].X = 2.08
	pointsOfFunctionPlot[208].Y = 8.736

	pointsOfFunctionPlot[209].X = 2.09
	pointsOfFunctionPlot[209].Y = 8.822

	pointsOfFunctionPlot[210].X = 2.10
	pointsOfFunctionPlot[210].Y = 8.908

	pointsOfFunctionPlot[211].X = 2.11
	pointsOfFunctionPlot[211].Y = 8.994

	pointsOfFunctionPlot[212].X = 2.12
	pointsOfFunctionPlot[212].Y = 9.082

	pointsOfFunctionPlot[213].X = 2.13
	pointsOfFunctionPlot[213].Y = 9.17

	pointsOfFunctionPlot[214].X = 2.14
	pointsOfFunctionPlot[214].Y = 9.26

	pointsOfFunctionPlot[215].X = 2.15
	pointsOfFunctionPlot[215].Y = 9.35

	pointsOfFunctionPlot[216].X = 2.16
	pointsOfFunctionPlot[216].Y = 9.441

	pointsOfFunctionPlot[217].X = 2.17
	pointsOfFunctionPlot[217].Y = 9.533

	pointsOfFunctionPlot[218].X = 2.18
	pointsOfFunctionPlot[218].Y = 9.625

	pointsOfFunctionPlot[219].X = 2.19
	pointsOfFunctionPlot[219].Y = 9.719

	pointsOfFunctionPlot[220].X = 2.20
	pointsOfFunctionPlot[220].Y = 9.813

	pointsOfFunctionPlot[221].X = 2.21
	pointsOfFunctionPlot[221].Y = 9.908

	pointsOfFunctionPlot[222].X = 2.22
	pointsOfFunctionPlot[222].Y = 10.004

	pointsOfFunctionPlot[223].X = 2.23
	pointsOfFunctionPlot[223].Y = 10.101

	pointsOfFunctionPlot[224].X = 2.24
	pointsOfFunctionPlot[224].Y = 10.199

	pointsOfFunctionPlot[225].X = 2.25
	pointsOfFunctionPlot[225].Y = 10.298

	pointsOfFunctionPlot[226].X = 2.26
	pointsOfFunctionPlot[226].Y = 10.398

	pointsOfFunctionPlot[227].X = 2.27
	pointsOfFunctionPlot[227].Y = 10.499

	pointsOfFunctionPlot[228].X = 2.28
	pointsOfFunctionPlot[228].Y = 10.6

	pointsOfFunctionPlot[229].X = 2.29
	pointsOfFunctionPlot[229].Y = 10.703

	pointsOfFunctionPlot[230].X = 2.30
	pointsOfFunctionPlot[230].Y = 10.807

	pointsOfFunctionPlot[231].X = 2.31
	pointsOfFunctionPlot[231].Y = 10.911

	pointsOfFunctionPlot[232].X = 2.32
	pointsOfFunctionPlot[232].Y = 11.017

	pointsOfFunctionPlot[233].X = 2.33
	pointsOfFunctionPlot[233].Y = 11.123

	pointsOfFunctionPlot[234].X = 2.34
	pointsOfFunctionPlot[234].Y = 11.231

	pointsOfFunctionPlot[235].X = 2.35
	pointsOfFunctionPlot[235].Y = 11.339

	pointsOfFunctionPlot[236].X = 2.36
	pointsOfFunctionPlot[236].Y = 11.449

	pointsOfFunctionPlot[237].X = 2.37
	pointsOfFunctionPlot[237].Y = 11.56

	pointsOfFunctionPlot[238].X = 2.38
	pointsOfFunctionPlot[238].Y = 11.672

	pointsOfFunctionPlot[239].X = 2.39
	pointsOfFunctionPlot[239].Y = 11.784

	pointsOfFunctionPlot[240].X = 2.40
	pointsOfFunctionPlot[240].Y = 11.898

	pointsOfFunctionPlot[241].X = 2.41
	pointsOfFunctionPlot[241].Y = 12.013

	pointsOfFunctionPlot[242].X = 2.42
	pointsOfFunctionPlot[242].Y = 12.129

	pointsOfFunctionPlot[243].X = 2.43
	pointsOfFunctionPlot[243].Y = 12.246

	pointsOfFunctionPlot[244].X = 2.44
	pointsOfFunctionPlot[244].Y = 12.365

	pointsOfFunctionPlot[245].X = 2.45
	pointsOfFunctionPlot[245].Y = 12.484

	pointsOfFunctionPlot[246].X = 2.46
	pointsOfFunctionPlot[246].Y = 12.604

	pointsOfFunctionPlot[247].X = 2.47
	pointsOfFunctionPlot[247].Y = 12.726

	pointsOfFunctionPlot[248].X = 2.48
	pointsOfFunctionPlot[248].Y = 12.849

	pointsOfFunctionPlot[249].X = 2.49
	pointsOfFunctionPlot[249].Y = 12.973

	pointsOfFunctionPlot[250].X = 2.50
	pointsOfFunctionPlot[250].Y = 13.098

	pointsOfFunctionPlot[251].X = 2.51
	pointsOfFunctionPlot[251].Y = 13.225

	pointsOfFunctionPlot[252].X = 2.52
	pointsOfFunctionPlot[252].Y = 13.352

	pointsOfFunctionPlot[253].X = 2.53
	pointsOfFunctionPlot[253].Y = 13.481

	pointsOfFunctionPlot[254].X = 2.54
	pointsOfFunctionPlot[254].Y = 13.611

	pointsOfFunctionPlot[255].X = 2.55
	pointsOfFunctionPlot[255].Y = 13.743

	pointsOfFunctionPlot[256].X = 2.56
	pointsOfFunctionPlot[256].Y = 13.875

	pointsOfFunctionPlot[257].X = 2.57
	pointsOfFunctionPlot[257].Y = 14.009

	pointsOfFunctionPlot[258].X = 2.58
	pointsOfFunctionPlot[258].Y = 14.144

	pointsOfFunctionPlot[259].X = 2.59
	pointsOfFunctionPlot[259].Y = 14.281

	pointsOfFunctionPlot[260].X = 2.60
	pointsOfFunctionPlot[260].Y = 14.419

	pointsOfFunctionPlot[261].X = 2.61
	pointsOfFunctionPlot[261].Y = 14.558

	pointsOfFunctionPlot[262].X = 2.62
	pointsOfFunctionPlot[262].Y = 14.698

	pointsOfFunctionPlot[263].X = 2.63
	pointsOfFunctionPlot[263].Y = 14.84

	pointsOfFunctionPlot[264].X = 2.64
	pointsOfFunctionPlot[264].Y = 14.983

	pointsOfFunctionPlot[265].X = 2.65
	pointsOfFunctionPlot[265].Y = 15.128

	pointsOfFunctionPlot[266].X = 2.66
	pointsOfFunctionPlot[266].Y = 15.274

	pointsOfFunctionPlot[267].X = 2.67
	pointsOfFunctionPlot[267].Y = 15.57

	pointsOfFunctionPlot[268].X = 2.68
	pointsOfFunctionPlot[268].Y = 15.721

	pointsOfFunctionPlot[269].X = 2.69
	pointsOfFunctionPlot[269].Y = 15.721

	pointsOfFunctionPlot[270].X = 2.70
	pointsOfFunctionPlot[270].Y = 15.872

	pointsOfFunctionPlot[271].X = 2.71
	pointsOfFunctionPlot[271].Y = 16.026

	pointsOfFunctionPlot[272].X = 2.72
	pointsOfFunctionPlot[272].Y = 16.18

	pointsOfFunctionPlot[273].X = 2.73
	pointsOfFunctionPlot[273].Y = 16.337

	pointsOfFunctionPlot[274].X = 2.74
	pointsOfFunctionPlot[274].Y = 16.494

	pointsOfFunctionPlot[275].X = 2.75
	pointsOfFunctionPlot[275].Y = 16.654

	pointsOfFunctionPlot[276].X = 2.76
	pointsOfFunctionPlot[276].Y = 16.815

	pointsOfFunctionPlot[277].X = 2.77
	pointsOfFunctionPlot[277].Y = 16.977

	pointsOfFunctionPlot[278].X = 2.78
	pointsOfFunctionPlot[278].Y = 17.141

	pointsOfFunctionPlot[279].X = 2.79
	pointsOfFunctionPlot[279].Y = 17.307

	pointsOfFunctionPlot[280].X = 2.80
	pointsOfFunctionPlot[280].Y = 17.474

	pointsOfFunctionPlot[281].X = 2.81
	pointsOfFunctionPlot[281].Y = 17.643

	pointsOfFunctionPlot[282].X = 2.82
	pointsOfFunctionPlot[282].Y = 17.813

	pointsOfFunctionPlot[283].X = 2.83
	pointsOfFunctionPlot[283].Y = 17.985

	pointsOfFunctionPlot[284].X = 2.84
	pointsOfFunctionPlot[284].Y = 18.159

	pointsOfFunctionPlot[285].X = 2.85
	pointsOfFunctionPlot[285].Y = 18.335

	pointsOfFunctionPlot[286].X = 2.86
	pointsOfFunctionPlot[286].Y = 18.512

	pointsOfFunctionPlot[287].X = 2.87
	pointsOfFunctionPlot[287].Y = 18.691

	pointsOfFunctionPlot[288].X = 2.88
	pointsOfFunctionPlot[288].Y = 18.872

	pointsOfFunctionPlot[289].X = 2.89
	pointsOfFunctionPlot[289].Y = 19.054

	pointsOfFunctionPlot[290].X = 2.90
	pointsOfFunctionPlot[290].Y = 19.238

	pointsOfFunctionPlot[291].X = 2.91
	pointsOfFunctionPlot[291].Y = 19.424

	pointsOfFunctionPlot[292].X = 2.92
	pointsOfFunctionPlot[292].Y = 19.612

	pointsOfFunctionPlot[293].X = 2.93
	pointsOfFunctionPlot[293].Y = 19.802

	pointsOfFunctionPlot[294].X = 2.94
	pointsOfFunctionPlot[294].Y = 19.994

	pointsOfFunctionPlot[295].X = 2.95
	pointsOfFunctionPlot[295].Y = 20.187

	pointsOfFunctionPlot[296].X = 2.96
	pointsOfFunctionPlot[296].Y = 20.383

	pointsOfFunctionPlot[297].X = 2.97
	pointsOfFunctionPlot[297].Y = 20.58

	pointsOfFunctionPlot[298].X = 2.98
	pointsOfFunctionPlot[298].Y = 20.779

	pointsOfFunctionPlot[299].X = 2.99
	pointsOfFunctionPlot[299].Y = 20.98

	pointsOfFunctionPlot[300].X = 3.0
	pointsOfFunctionPlot[300].Y = 21.184

	pointsOfFunctionPlot[301].X = 3.01
	pointsOfFunctionPlot[301].Y = 21.389

	pointsOfFunctionPlot[302].X = 3.02
	pointsOfFunctionPlot[302].Y = 21.596

	pointsOfFunctionPlot[303].X = 3.03
	pointsOfFunctionPlot[303].Y = 21.805

	pointsOfFunctionPlot[304].X = 3.04
	pointsOfFunctionPlot[304].Y = 22.017

	pointsOfFunctionPlot[305].X = 3.05
	pointsOfFunctionPlot[305].Y = 22.23

	pointsOfFunctionPlot[306].X = 3.06
	pointsOfFunctionPlot[306].Y = 22.445

	pointsOfFunctionPlot[307].X = 3.07
	pointsOfFunctionPlot[307].Y = 22.663

	pointsOfFunctionPlot[308].X = 3.08
	pointsOfFunctionPlot[308].Y = 22.883

	pointsOfFunctionPlot[309].X = 3.09
	pointsOfFunctionPlot[309].Y = 23.1105

	pointsOfFunctionPlot[310].X = 3.10
	pointsOfFunctionPlot[310].Y = 23.329

	pointsOfFunctionPlot[311].X = 3.11
	pointsOfFunctionPlot[311].Y = 23.555

	pointsOfFunctionPlot[312].X = 3.12
	pointsOfFunctionPlot[312].Y = 23.784

	pointsOfFunctionPlot[313].X = 3.13
	pointsOfFunctionPlot[313].Y = 24.015

	pointsOfFunctionPlot[314].X = 3.14
	pointsOfFunctionPlot[314].Y = 24.248

	pointsOfFunctionPlot[315].X = 3.15
	pointsOfFunctionPlot[315].Y = 24.483

	pointsOfFunctionPlot[316].X = 3.16
	pointsOfFunctionPlot[316].Y = 24.721

	pointsOfFunctionPlot[317].X = 3.17
	pointsOfFunctionPlot[317].Y = 24.861

	pointsOfFunctionPlot[318].X = 3.18
	pointsOfFunctionPlot[318].Y = 25.203

	pointsOfFunctionPlot[319].X = 3.19
	pointsOfFunctionPlot[319].Y = 25.448

	pointsOfFunctionPlot[320].X = 3.20
	pointsOfFunctionPlot[320].Y = 25.695

	pointsOfFunctionPlot[321].X = 3.21
	pointsOfFunctionPlot[321].Y = 25.945

	pointsOfFunctionPlot[322].X = 3.22
	pointsOfFunctionPlot[322].Y = 26.197

	pointsOfFunctionPlot[323].X = 3.23
	pointsOfFunctionPlot[323].Y = 26.452

	pointsOfFunctionPlot[324].X = 3.24
	pointsOfFunctionPlot[324].Y = 26.709

	pointsOfFunctionPlot[325].X = 3.25
	pointsOfFunctionPlot[325].Y = 26.968

	pointsOfFunctionPlot[326].X = 3.26
	pointsOfFunctionPlot[326].Y = 27.231

	pointsOfFunctionPlot[327].X = 3.27
	pointsOfFunctionPlot[327].Y = 27.496

	pointsOfFunctionPlot[328].X = 3.28
	pointsOfFunctionPlot[328].Y = 27.763

	pointsOfFunctionPlot[329].X = 3.29
	pointsOfFunctionPlot[329].Y = 28.033

	pointsOfFunctionPlot[330].X = 3.30
	pointsOfFunctionPlot[330].Y = 28.306

	pointsOfFunctionPlot[331].X = 3.31
	pointsOfFunctionPlot[331].Y = 28.582

	pointsOfFunctionPlot[332].X = 3.32
	pointsOfFunctionPlot[332].Y = 28.86

	pointsOfFunctionPlot[333].X = 3.33
	pointsOfFunctionPlot[333].Y = 29.141

	pointsOfFunctionPlot[334].X = 3.34
	pointsOfFunctionPlot[334].Y = 29.425

	pointsOfFunctionPlot[335].X = 3.35
	pointsOfFunctionPlot[335].Y = 29.711

	pointsOfFunctionPlot[336].X = 3.36
	pointsOfFunctionPlot[336].Y = 30.001

	pointsOfFunctionPlot[337].X = 3.37
	pointsOfFunctionPlot[337].Y = 30.293

	pointsOfFunctionPlot[338].X = 3.38
	pointsOfFunctionPlot[338].Y = 30.588

	pointsOfFunctionPlot[339].X = 3.39
	pointsOfFunctionPlot[339].Y = 30.886

	pointsOfFunctionPlot[340].X = 3.40
	pointsOfFunctionPlot[340].Y = 31.187

	pointsOfFunctionPlot[341].X = 3.41
	pointsOfFunctionPlot[341].Y = 31.491

	pointsOfFunctionPlot[342].X = 3.42
	pointsOfFunctionPlot[342].Y = 31.799

	pointsOfFunctionPlot[343].X = 3.43
	pointsOfFunctionPlot[343].Y = 32.109

	pointsOfFunctionPlot[344].X = 3.44
	pointsOfFunctionPlot[344].Y = 32.422

	pointsOfFunctionPlot[345].X = 3.45
	pointsOfFunctionPlot[345].Y = 32.738

	pointsOfFunctionPlot[346].X = 3.46
	pointsOfFunctionPlot[346].Y = 33.058

	pointsOfFunctionPlot[347].X = 3.47
	pointsOfFunctionPlot[347].Y = 33.38

	pointsOfFunctionPlot[348].X = 3.48
	pointsOfFunctionPlot[348].Y = 33.706

	pointsOfFunctionPlot[349].X = 3.49
	pointsOfFunctionPlot[349].Y = 34.035

	pointsOfFunctionPlot[350].X = 3.50
	pointsOfFunctionPlot[350].Y = 34.368

	pointsOfFunctionPlot[351].X = 3.51
	pointsOfFunctionPlot[351].Y = 34.703

	pointsOfFunctionPlot[352].X = 3.52
	pointsOfFunctionPlot[352].Y = 35.042

	pointsOfFunctionPlot[353].X = 3.53
	pointsOfFunctionPlot[353].Y = 35.385

	pointsOfFunctionPlot[354].X = 3.54
	pointsOfFunctionPlot[354].Y = 35.731

	pointsOfFunctionPlot[355].X = 3.55
	pointsOfFunctionPlot[355].Y = 36.08

	pointsOfFunctionPlot[356].X = 3.56
	pointsOfFunctionPlot[356].Y = 36.432

	pointsOfFunctionPlot[357].X = 3.57
	pointsOfFunctionPlot[357].Y = 36.789

	pointsOfFunctionPlot[358].X = 3.58
	pointsOfFunctionPlot[358].Y = 37.148

	pointsOfFunctionPlot[359].X = 3.59
	pointsOfFunctionPlot[359].Y = 37.512

	pointsOfFunctionPlot[360].X = 3.60
	pointsOfFunctionPlot[360].Y = 37.879

	pointsOfFunctionPlot[361].X = 3.61
	pointsOfFunctionPlot[361].Y = 38.249

	pointsOfFunctionPlot[362].X = 3.62
	pointsOfFunctionPlot[362].Y = 38.624

	pointsOfFunctionPlot[363].X = 3.63
	pointsOfFunctionPlot[363].Y = 39.002

	pointsOfFunctionPlot[364].X = 3.64
	pointsOfFunctionPlot[364].Y = 39.383

	pointsOfFunctionPlot[365].X = 3.65
	pointsOfFunctionPlot[365].Y = 39.769

	pointsOfFunctionPlot[366].X = 3.66
	pointsOfFunctionPlot[366].Y = 40.158

	pointsOfFunctionPlot[367].X = 3.67
	pointsOfFunctionPlot[367].Y = 40.552

	pointsOfFunctionPlot[368].X = 3.68
	pointsOfFunctionPlot[368].Y = 40.949

	pointsOfFunctionPlot[369].X = 3.69
	pointsOfFunctionPlot[369].Y = 41.35

	pointsOfFunctionPlot[370].X = 3.70
	pointsOfFunctionPlot[370].Y = 41.755

	pointsOfFunctionPlot[371].X = 3.71
	pointsOfFunctionPlot[371].Y = 42.164

	pointsOfFunctionPlot[372].X = 3.72
	pointsOfFunctionPlot[372].Y = 42.578

	pointsOfFunctionPlot[373].X = 3.73
	pointsOfFunctionPlot[373].Y = 42.995

	pointsOfFunctionPlot[374].X = 3.74
	pointsOfFunctionPlot[374].Y = 43.417

	pointsOfFunctionPlot[375].X = 3.75
	pointsOfFunctionPlot[375].Y = 43.842

	pointsOfFunctionPlot[376].X = 3.76
	pointsOfFunctionPlot[376].Y = 44.272

	pointsOfFunctionPlot[377].X = 3.77
	pointsOfFunctionPlot[377].Y = 44.707

	pointsOfFunctionPlot[378].X = 3.78
	pointsOfFunctionPlot[378].Y = 45.145

	pointsOfFunctionPlot[379].X = 3.79
	pointsOfFunctionPlot[379].Y = 45.588

	pointsOfFunctionPlot[380].X = 3.80
	pointsOfFunctionPlot[380].Y = 46.036

	pointsOfFunctionPlot[381].X = 3.81
	pointsOfFunctionPlot[381].Y = 46.488

	pointsOfFunctionPlot[382].X = 3.82
	pointsOfFunctionPlot[382].Y = 46.944

	pointsOfFunctionPlot[383].X = 3.83
	pointsOfFunctionPlot[383].Y = 47.405

	pointsOfFunctionPlot[384].X = 3.84
	pointsOfFunctionPlot[384].Y = 47.87

	pointsOfFunctionPlot[385].X = 3.85
	pointsOfFunctionPlot[385].Y = 48.341

	pointsOfFunctionPlot[386].X = 3.86
	pointsOfFunctionPlot[386].Y = 48.816

	pointsOfFunctionPlot[387].X = 3.87
	pointsOfFunctionPlot[387].Y = 49.295

	pointsOfFunctionPlot[388].X = 3.88
	pointsOfFunctionPlot[388].Y = 49.78

	pointsOfFunctionPlot[389].X = 3.89
	pointsOfFunctionPlot[389].Y = 50.269

	pointsOfFunctionPlot[390].X = 3.90
	pointsOfFunctionPlot[390].Y = 50.763

	pointsOfFunctionPlot[391].X = 3.91
	pointsOfFunctionPlot[391].Y = 51.262

	pointsOfFunctionPlot[392].X = 3.92
	pointsOfFunctionPlot[392].Y = 51.766

	pointsOfFunctionPlot[393].X = 3.93
	pointsOfFunctionPlot[393].Y = 52.275

	pointsOfFunctionPlot[394].X = 3.94
	pointsOfFunctionPlot[394].Y = 52.789

	pointsOfFunctionPlot[395].X = 3.95
	pointsOfFunctionPlot[395].Y = 53.309

	pointsOfFunctionPlot[396].X = 3.96
	pointsOfFunctionPlot[396].Y = 53.833

	pointsOfFunctionPlot[397].X = 3.97
	pointsOfFunctionPlot[397].Y = 54.363

	pointsOfFunctionPlot[398].X = 3.98
	pointsOfFunctionPlot[398].Y = 54.898

	pointsOfFunctionPlot[399].X = 3.99
	pointsOfFunctionPlot[399].Y = 55.438

	pointsOfFunctionPlot[400].X = 4.0
	pointsOfFunctionPlot[400].Y = 55.984

	pointsOfFunctionPlot[401].X = 4.01
	pointsOfFunctionPlot[401].Y = 56.535

	pointsOfFunctionPlot[402].X = 4.02
	pointsOfFunctionPlot[402].Y = 57.092

	pointsOfFunctionPlot[403].X = 4.03
	pointsOfFunctionPlot[403].Y = 57.654

	pointsOfFunctionPlot[404].X = 4.04
	pointsOfFunctionPlot[404].Y = 58.222

	pointsOfFunctionPlot[405].X = 4.05
	pointsOfFunctionPlot[405].Y = 58.796

	pointsOfFunctionPlot[406].X = 4.06
	pointsOfFunctionPlot[406].Y = 59.375

	pointsOfFunctionPlot[407].X = 4.07
	pointsOfFunctionPlot[407].Y = 59.96

	pointsOfFunctionPlot[408].X = 4.08
	pointsOfFunctionPlot[408].Y = 60.551

	pointsOfFunctionPlot[409].X = 4.09
	pointsOfFunctionPlot[409].Y = 61.148

	pointsOfFunctionPlot[410].X = 4.10
	pointsOfFunctionPlot[410].Y = 61.751

	pointsOfFunctionPlot[411].X = 4.11
	pointsOfFunctionPlot[411].Y = 62.36

	pointsOfFunctionPlot[412].X = 4.12
	pointsOfFunctionPlot[412].Y = 62.975

	pointsOfFunctionPlot[413].X = 4.13
	pointsOfFunctionPlot[413].Y = 63.596

	pointsOfFunctionPlot[414].X = 4.14
	pointsOfFunctionPlot[414].Y = 64.223

	pointsOfFunctionPlot[415].X = 4.15
	pointsOfFunctionPlot[415].Y = 64.857

	pointsOfFunctionPlot[416].X = 4.16
	pointsOfFunctionPlot[416].Y = 65.497

	pointsOfFunctionPlot[417].X = 4.17
	pointsOfFunctionPlot[417].Y = 66.143

	pointsOfFunctionPlot[418].X = 4.18
	pointsOfFunctionPlot[418].Y = 66.796

	pointsOfFunctionPlot[419].X = 4.19
	pointsOfFunctionPlot[419].Y = 67.455

	pointsOfFunctionPlot[420].X = 4.20
	pointsOfFunctionPlot[420].Y = 68.121

	pointsOfFunctionPlot[421].X = 4.21
	pointsOfFunctionPlot[421].Y = 68.794

	pointsOfFunctionPlot[422].X = 4.22
	pointsOfFunctionPlot[422].Y = 69.473

	pointsOfFunctionPlot[423].X = 4.23
	pointsOfFunctionPlot[423].Y = 70.159

	pointsOfFunctionPlot[424].X = 4.24
	pointsOfFunctionPlot[424].Y = 70.852

	pointsOfFunctionPlot[425].X = 4.25
	pointsOfFunctionPlot[425].Y = 71.552

	pointsOfFunctionPlot[426].X = 4.26
	pointsOfFunctionPlot[426].Y = 72.259

	pointsOfFunctionPlot[427].X = 4.27
	pointsOfFunctionPlot[427].Y = 72.973

	pointsOfFunctionPlot[428].X = 4.28
	pointsOfFunctionPlot[428].Y = 73.694

	pointsOfFunctionPlot[429].X = 4.29
	pointsOfFunctionPlot[429].Y = 74.422

	pointsOfFunctionPlot[430].X = 4.30
	pointsOfFunctionPlot[430].Y = 75.158

	pointsOfFunctionPlot[431].X = 4.31
	pointsOfFunctionPlot[431].Y = 75.901

	pointsOfFunctionPlot[432].X = 4.32
	pointsOfFunctionPlot[432].Y = 76.651

	pointsOfFunctionPlot[433].X = 4.33
	pointsOfFunctionPlot[433].Y = 77.409

	pointsOfFunctionPlot[434].X = 4.34
	pointsOfFunctionPlot[434].Y = 78.175

	pointsOfFunctionPlot[435].X = 4.35
	pointsOfFunctionPlot[435].Y = 78.948

	pointsOfFunctionPlot[436].X = 4.36
	pointsOfFunctionPlot[436].Y = 79.729

	pointsOfFunctionPlot[437].X = 4.37
	pointsOfFunctionPlot[437].Y = 80.518

	pointsOfFunctionPlot[438].X = 4.38
	pointsOfFunctionPlot[438].Y = 81.315

	pointsOfFunctionPlot[439].X = 4.39
	pointsOfFunctionPlot[439].Y = 82.119

	pointsOfFunctionPlot[440].X = 4.40
	pointsOfFunctionPlot[440].Y = 82.932

	pointsOfFunctionPlot[441].X = 4.41
	pointsOfFunctionPlot[441].Y = 83.753

	pointsOfFunctionPlot[442].X = 4.42
	pointsOfFunctionPlot[442].Y = 84.583

	pointsOfFunctionPlot[443].X = 4.43
	pointsOfFunctionPlot[443].Y = 85.419

	pointsOfFunctionPlot[444].X = 4.44
	pointsOfFunctionPlot[444].Y = 86.265

	pointsOfFunctionPlot[445].X = 4.45
	pointsOfFunctionPlot[445].Y = 87.119

	pointsOfFunctionPlot[446].X = 4.46
	pointsOfFunctionPlot[446].Y = 87.982

	pointsOfFunctionPlot[447].X = 4.47
	pointsOfFunctionPlot[447].Y = 88.854

	pointsOfFunctionPlot[448].X = 4.48
	pointsOfFunctionPlot[448].Y = 89.734

	pointsOfFunctionPlot[449].X = 4.49
	pointsOfFunctionPlot[449].Y = 90.623

	pointsOfFunctionPlot[450].X = 4.50
	pointsOfFunctionPlot[450].Y = 91.521

	pointsOfFunctionPlot[451].X = 4.51
	pointsOfFunctionPlot[451].Y = 92.428

	pointsOfFunctionPlot[452].X = 4.52
	pointsOfFunctionPlot[452].Y = 93.344

	pointsOfFunctionPlot[453].X = 4.53
	pointsOfFunctionPlot[453].Y = 94.269

	pointsOfFunctionPlot[454].X = 4.54
	pointsOfFunctionPlot[454].Y = 95.203

	pointsOfFunctionPlot[455].X = 4.55
	pointsOfFunctionPlot[455].Y = 96.147

	pointsOfFunctionPlot[456].X = 4.56
	pointsOfFunctionPlot[456].Y = 97.1

	pointsOfFunctionPlot[457].X = 4.57
	pointsOfFunctionPlot[457].Y = 98.063

	pointsOfFunctionPlot[458].X = 4.58
	pointsOfFunctionPlot[458].Y = 99.036

	pointsOfFunctionPlot[459].X = 4.59
	pointsOfFunctionPlot[459].Y = 100.018

	pointsOfFunctionPlot[460].X = 4.60
	pointsOfFunctionPlot[460].Y = 101.01

	pointsOfFunctionPlot[461].X = 4.61
	pointsOfFunctionPlot[461].Y = 102.012

	pointsOfFunctionPlot[462].X = 4.62
	pointsOfFunctionPlot[462].Y = 103.024

	pointsOfFunctionPlot[463].X = 4.63
	pointsOfFunctionPlot[463].Y = 104.046

	pointsOfFunctionPlot[464].X = 4.64
	pointsOfFunctionPlot[464].Y = 105.079

	pointsOfFunctionPlot[465].X = 4.65
	pointsOfFunctionPlot[465].Y = 106.121

	pointsOfFunctionPlot[466].X = 4.66
	pointsOfFunctionPlot[466].Y = 107.175

	pointsOfFunctionPlot[467].X = 4.67
	pointsOfFunctionPlot[467].Y = 108.238

	pointsOfFunctionPlot[468].X = 4.68
	pointsOfFunctionPlot[468].Y = 109.313

	pointsOfFunctionPlot[469].X = 4.69
	pointsOfFunctionPlot[469].Y = 110.398

	pointsOfFunctionPlot[470].X = 4.70
	pointsOfFunctionPlot[470].Y = 111.494

	pointsOfFunctionPlot[471].X = 4.71
	pointsOfFunctionPlot[471].Y = 112.601

	pointsOfFunctionPlot[472].X = 4.72
	pointsOfFunctionPlot[472].Y = 113.72

	pointsOfFunctionPlot[473].X = 4.73
	pointsOfFunctionPlot[473].Y = 114.849

	pointsOfFunctionPlot[474].X = 4.74
	pointsOfFunctionPlot[474].Y = 115.99

	pointsOfFunctionPlot[475].X = 4.75
	pointsOfFunctionPlot[475].Y = 117.142

	pointsOfFunctionPlot[476].X = 4.76
	pointsOfFunctionPlot[476].Y = 118.306

	pointsOfFunctionPlot[477].X = 4.77
	pointsOfFunctionPlot[477].Y = 119.481

	pointsOfFunctionPlot[478].X = 4.78
	pointsOfFunctionPlot[478].Y = 120.668

	pointsOfFunctionPlot[479].X = 4.79
	pointsOfFunctionPlot[479].Y = 121.867

	pointsOfFunctionPlot[480].X = 4.80
	pointsOfFunctionPlot[480].Y = 123.079

	pointsOfFunctionPlot[481].X = 4.81
	pointsOfFunctionPlot[481].Y = 124.302

	pointsOfFunctionPlot[482].X = 4.82
	pointsOfFunctionPlot[482].Y = 125.537

	pointsOfFunctionPlot[483].X = 4.83
	pointsOfFunctionPlot[483].Y = 126.785

	pointsOfFunctionPlot[484].X = 4.84
	pointsOfFunctionPlot[484].Y = 128.046

	pointsOfFunctionPlot[485].X = 4.85
	pointsOfFunctionPlot[485].Y = 129.319

	pointsOfFunctionPlot[486].X = 4.86
	pointsOfFunctionPlot[486].Y = 130.605

	pointsOfFunctionPlot[487].X = 4.87
	pointsOfFunctionPlot[487].Y = 131.904

	pointsOfFunctionPlot[488].X = 4.88
	pointsOfFunctionPlot[488].Y = 133.215

	pointsOfFunctionPlot[489].X = 4.89
	pointsOfFunctionPlot[489].Y = 134.54

	pointsOfFunctionPlot[490].X = 4.90
	pointsOfFunctionPlot[490].Y = 135.879

	pointsOfFunctionPlot[491].X = 4.91
	pointsOfFunctionPlot[491].Y = 137.23

	pointsOfFunctionPlot[492].X = 4.92
	pointsOfFunctionPlot[492].Y = 138.595

	pointsOfFunctionPlot[493].X = 4.93
	pointsOfFunctionPlot[493].Y = 139.974

	pointsOfFunctionPlot[494].X = 4.94
	pointsOfFunctionPlot[494].Y = 141.367

	pointsOfFunctionPlot[495].X = 4.95
	pointsOfFunctionPlot[495].Y = 142.774

	pointsOfFunctionPlot[496].X = 4.96
	pointsOfFunctionPlot[496].Y = 144.195

	pointsOfFunctionPlot[497].X = 4.97
	pointsOfFunctionPlot[497].Y = 145.63

	pointsOfFunctionPlot[498].X = 4.98
	pointsOfFunctionPlot[498].Y = 147.079

	pointsOfFunctionPlot[499].X = 4.99
	pointsOfFunctionPlot[499].Y = 148.543

	pointsOfFunctionPlot[500].X = 5.0
	pointsOfFunctionPlot[500].Y = 150.022










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function f(x) = exp(x) - ln(x)"

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
		"exp-minus-ln-function-plot-01.png"); err != nil {

		panic(err)
	}
}
