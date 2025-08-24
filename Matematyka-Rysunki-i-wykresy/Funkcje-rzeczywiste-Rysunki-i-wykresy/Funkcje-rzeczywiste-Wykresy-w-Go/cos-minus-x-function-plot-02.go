package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function f(x) = cos(x) - x.

	pointsOfFunctionPlot := make(plotter.XYs, 1_261)

	pointsOfFunctionPlot[0].X = 0.0
	pointsOfFunctionPlot[0].Y = 1.0

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = 0.99

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = 0.979

	pointsOfFunctionPlot[3].X = 0.03
	pointsOfFunctionPlot[3].Y = 0.969

	pointsOfFunctionPlot[4].X = 0.04
	pointsOfFunctionPlot[4].Y = 0.959

	pointsOfFunctionPlot[5].X = 0.05
	pointsOfFunctionPlot[5].Y = 0.948

	pointsOfFunctionPlot[6].X = 0.06
	pointsOfFunctionPlot[6].Y = 0.938

	pointsOfFunctionPlot[7].X = 0.07
	pointsOfFunctionPlot[7].Y = 0.927

	pointsOfFunctionPlot[8].X = 0.08
	pointsOfFunctionPlot[8].Y = 0.916

	pointsOfFunctionPlot[9].X = 0.09
	pointsOfFunctionPlot[9].Y = 0.906

	pointsOfFunctionPlot[10].X = 0.1
	pointsOfFunctionPlot[10].Y = 0.895

	pointsOfFunctionPlot[11].X = 0.11
	pointsOfFunctionPlot[11].Y = 0.884

	pointsOfFunctionPlot[12].X = 0.12
	pointsOfFunctionPlot[12].Y = 0.872

	pointsOfFunctionPlot[13].X = 0.13
	pointsOfFunctionPlot[13].Y = 0.861

	pointsOfFunctionPlot[14].X = 0.14
	pointsOfFunctionPlot[14].Y = 0.85

	pointsOfFunctionPlot[15].X = 0.15
	pointsOfFunctionPlot[15].Y = 0.838

	pointsOfFunctionPlot[16].X = 0.16
	pointsOfFunctionPlot[16].Y = 0.827

	pointsOfFunctionPlot[17].X = 0.17
	pointsOfFunctionPlot[17].Y = 0.815

	pointsOfFunctionPlot[18].X = 0.18
	pointsOfFunctionPlot[18].Y = 0.803

	pointsOfFunctionPlot[19].X = 0.19
	pointsOfFunctionPlot[19].Y = 0.792

	pointsOfFunctionPlot[20].X = 0.2
	pointsOfFunctionPlot[20].Y = 0.78

	pointsOfFunctionPlot[21].X = 0.21
	pointsOfFunctionPlot[21].Y = 0.768

	pointsOfFunctionPlot[22].X = 0.22
	pointsOfFunctionPlot[22].Y = 0.755

	pointsOfFunctionPlot[23].X = 0.23
	pointsOfFunctionPlot[23].Y = 0.743

	pointsOfFunctionPlot[24].X = 0.24
	pointsOfFunctionPlot[24].Y = 0.731

	pointsOfFunctionPlot[25].X = 0.25
	pointsOfFunctionPlot[25].Y = 0.718

	pointsOfFunctionPlot[26].X = 0.26
	pointsOfFunctionPlot[26].Y = 0.706

	pointsOfFunctionPlot[27].X = 0.27
	pointsOfFunctionPlot[27].Y = 0.693

	pointsOfFunctionPlot[28].X = 0.28
	pointsOfFunctionPlot[28].Y = 0.681

	pointsOfFunctionPlot[29].X = 0.29
	pointsOfFunctionPlot[29].Y = 0.668

	pointsOfFunctionPlot[30].X = 0.3
	pointsOfFunctionPlot[30].Y = 0.655

	pointsOfFunctionPlot[31].X = 0.31
	pointsOfFunctionPlot[31].Y = 0.642

	pointsOfFunctionPlot[32].X = 0.32
	pointsOfFunctionPlot[32].Y = 0.629

	pointsOfFunctionPlot[33].X = 0.33
	pointsOfFunctionPlot[33].Y = 0.616

	pointsOfFunctionPlot[34].X = 0.34
	pointsOfFunctionPlot[34].Y = 0.602

	pointsOfFunctionPlot[35].X = 0.35
	pointsOfFunctionPlot[35].Y = 0.589

	pointsOfFunctionPlot[36].X = 0.36
	pointsOfFunctionPlot[36].Y = 0.575

	pointsOfFunctionPlot[37].X = 0.37
	pointsOfFunctionPlot[37].Y = 0.562

	pointsOfFunctionPlot[38].X = 0.38
	pointsOfFunctionPlot[38].Y = 0.548

	pointsOfFunctionPlot[39].X = 0.39
	pointsOfFunctionPlot[39].Y = 0.534

	pointsOfFunctionPlot[40].X = 0.4
	pointsOfFunctionPlot[40].Y = 0.521

	pointsOfFunctionPlot[41].X = 0.41
	pointsOfFunctionPlot[41].Y = 0.507

	pointsOfFunctionPlot[42].X = 0.42
	pointsOfFunctionPlot[42].Y = 0.493

	pointsOfFunctionPlot[43].X = 0.43
	pointsOfFunctionPlot[43].Y = 0.479

	pointsOfFunctionPlot[44].X = 0.44
	pointsOfFunctionPlot[44].Y = 0.464

	pointsOfFunctionPlot[45].X = 0.45
	pointsOfFunctionPlot[45].Y = 0.45

	pointsOfFunctionPlot[46].X = 0.46
	pointsOfFunctionPlot[46].Y = 0.436

	pointsOfFunctionPlot[47].X = 0.47
	pointsOfFunctionPlot[47].Y = 0.421

	pointsOfFunctionPlot[48].X = 0.48
	pointsOfFunctionPlot[48].Y = 0.407

	pointsOfFunctionPlot[49].X = 0.49
	pointsOfFunctionPlot[49].Y = 0.392

	pointsOfFunctionPlot[50].X = 0.5
	pointsOfFunctionPlot[50].Y = 0.377

	pointsOfFunctionPlot[51].X = 0.51
	pointsOfFunctionPlot[51].Y = 0.362

	pointsOfFunctionPlot[52].X = 0.52
	pointsOfFunctionPlot[52].Y = 0.347

	pointsOfFunctionPlot[53].X = 0.53
	pointsOfFunctionPlot[53].Y = 0.332

	pointsOfFunctionPlot[54].X = 0.54
	pointsOfFunctionPlot[54].Y = 0.317

	pointsOfFunctionPlot[55].X = 0.55
	pointsOfFunctionPlot[55].Y = 0.302

	pointsOfFunctionPlot[56].X = 0.56
	pointsOfFunctionPlot[56].Y = 0.287

	pointsOfFunctionPlot[57].X = 0.57
	pointsOfFunctionPlot[57].Y = 0.271

	pointsOfFunctionPlot[58].X = 0.58
	pointsOfFunctionPlot[58].Y = 0.256

	pointsOfFunctionPlot[59].X = 0.59
	pointsOfFunctionPlot[59].Y = 0.24

	pointsOfFunctionPlot[60].X = 0.6
	pointsOfFunctionPlot[60].Y = 0.225

	pointsOfFunctionPlot[61].X = 0.61
	pointsOfFunctionPlot[61].Y = 0.209

	pointsOfFunctionPlot[62].X = 0.62
	pointsOfFunctionPlot[62].Y = 0.193

	pointsOfFunctionPlot[63].X = 0.63
	pointsOfFunctionPlot[63].Y = 0.178

	pointsOfFunctionPlot[64].X = 0.64
	pointsOfFunctionPlot[64].Y = 0.162

	pointsOfFunctionPlot[65].X = 0.65
	pointsOfFunctionPlot[65].Y = 0.146

	pointsOfFunctionPlot[66].X = 0.66
	pointsOfFunctionPlot[66].Y = 0.13

	pointsOfFunctionPlot[67].X = 0.67
	pointsOfFunctionPlot[67].Y = 0.113

	pointsOfFunctionPlot[68].X = 0.68
	pointsOfFunctionPlot[68].Y = 0.097

	pointsOfFunctionPlot[69].X = 0.69
	pointsOfFunctionPlot[69].Y = 0.081

	pointsOfFunctionPlot[70].X = 0.7
	pointsOfFunctionPlot[70].Y = 0.064

	pointsOfFunctionPlot[71].X = 0.71
	pointsOfFunctionPlot[71].Y = 0.048

	pointsOfFunctionPlot[72].X = 0.72
	pointsOfFunctionPlot[72].Y = 0.031

	pointsOfFunctionPlot[73].X = 0.73
	pointsOfFunctionPlot[73].Y = 0.015

	pointsOfFunctionPlot[74].X = 0.74
	pointsOfFunctionPlot[74].Y = -0.001

	pointsOfFunctionPlot[75].X = 0.75
	pointsOfFunctionPlot[75].Y = -0.018

	pointsOfFunctionPlot[76].X = 0.76
	pointsOfFunctionPlot[76].Y = -0.035

	pointsOfFunctionPlot[77].X = 0.77
	pointsOfFunctionPlot[77].Y = -0.052

	pointsOfFunctionPlot[78].X = 0.78
	pointsOfFunctionPlot[78].Y = -0.069

	pointsOfFunctionPlot[79].X = 0.79
	pointsOfFunctionPlot[79].Y = -0.086

	pointsOfFunctionPlot[80].X = 0.8
	pointsOfFunctionPlot[80].Y = -0.103

	pointsOfFunctionPlot[81].X = 0.81
	pointsOfFunctionPlot[81].Y = -0.12

	pointsOfFunctionPlot[82].X = 0.82
	pointsOfFunctionPlot[82].Y = -0.137

	pointsOfFunctionPlot[83].X = 0.83
	pointsOfFunctionPlot[83].Y = -0.155

	pointsOfFunctionPlot[84].X = 0.84
	pointsOfFunctionPlot[84].Y = -0.172

	pointsOfFunctionPlot[85].X = 0.85
	pointsOfFunctionPlot[85].Y = -0.19

	pointsOfFunctionPlot[86].X = 0.86
	pointsOfFunctionPlot[86].Y = -0.207

	pointsOfFunctionPlot[87].X = 0.87
	pointsOfFunctionPlot[87].Y = -0.225

	pointsOfFunctionPlot[88].X = 0.88
	pointsOfFunctionPlot[88].Y = -0.242

	pointsOfFunctionPlot[89].X = 0.89
	pointsOfFunctionPlot[89].Y = -0.26

	pointsOfFunctionPlot[90].X = 0.9
	pointsOfFunctionPlot[90].Y = -0.278

	pointsOfFunctionPlot[91].X = 0.91
	pointsOfFunctionPlot[91].Y = -0.296

	pointsOfFunctionPlot[92].X = 0.92
	pointsOfFunctionPlot[92].Y = -0.314

	pointsOfFunctionPlot[93].X = 0.93
	pointsOfFunctionPlot[93].Y = -0.332

	pointsOfFunctionPlot[94].X = 0.94
	pointsOfFunctionPlot[94].Y = -0.35

	pointsOfFunctionPlot[95].X = 0.95
	pointsOfFunctionPlot[95].Y = -0.368

	pointsOfFunctionPlot[96].X = 0.96
	pointsOfFunctionPlot[96].Y = -0.386

	pointsOfFunctionPlot[97].X = 0.97
	pointsOfFunctionPlot[97].Y = -0.404

	pointsOfFunctionPlot[98].X = 0.98
	pointsOfFunctionPlot[98].Y = -0.423

	pointsOfFunctionPlot[99].X = 0.99
	pointsOfFunctionPlot[99].Y = -0.441

	pointsOfFunctionPlot[100].X = 1.0
	pointsOfFunctionPlot[100].Y = -0.459

	pointsOfFunctionPlot[101].X = 1.01
	pointsOfFunctionPlot[101].Y = -0.478

	pointsOfFunctionPlot[102].X = 1.02
	pointsOfFunctionPlot[102].Y = -0.496

	pointsOfFunctionPlot[103].X = 1.03
	pointsOfFunctionPlot[103].Y = -0.515

	pointsOfFunctionPlot[104].X = 1.04
	pointsOfFunctionPlot[104].Y = -0.533

	pointsOfFunctionPlot[105].X = 1.05
	pointsOfFunctionPlot[105].Y = -0.552

	pointsOfFunctionPlot[106].X = 1.06
	pointsOfFunctionPlot[106].Y = -0.571

	pointsOfFunctionPlot[107].X = 1.07
	pointsOfFunctionPlot[107].Y = -0.589

	pointsOfFunctionPlot[108].X = 1.08
	pointsOfFunctionPlot[108].Y = -0.608

	pointsOfFunctionPlot[109].X = 1.09
	pointsOfFunctionPlot[109].Y = -0.627

	pointsOfFunctionPlot[110].X = 1.1
	pointsOfFunctionPlot[110].Y = -0.646

	pointsOfFunctionPlot[111].X = 1.11
	pointsOfFunctionPlot[111].Y = -0.665

	pointsOfFunctionPlot[112].X = 1.12
	pointsOfFunctionPlot[112].Y = -0.684

	pointsOfFunctionPlot[113].X = 1.13
	pointsOfFunctionPlot[113].Y = -0.703

	pointsOfFunctionPlot[114].X = 1.14
	pointsOfFunctionPlot[114].Y = -0.722

	pointsOfFunctionPlot[115].X = 1.15
	pointsOfFunctionPlot[115].Y = -0.741

	pointsOfFunctionPlot[116].X = 1.16
	pointsOfFunctionPlot[116].Y = -0.76

	pointsOfFunctionPlot[117].X = 1.17
	pointsOfFunctionPlot[117].Y = -0.779

	pointsOfFunctionPlot[118].X = 1.18
	pointsOfFunctionPlot[118].Y = -0.799

	pointsOfFunctionPlot[119].X = 1.19
	pointsOfFunctionPlot[119].Y = -0.818

	pointsOfFunctionPlot[120].X = 1.2
	pointsOfFunctionPlot[120].Y = -0.837

	pointsOfFunctionPlot[121].X = 1.21
	pointsOfFunctionPlot[121].Y = -0.857

	pointsOfFunctionPlot[122].X = 1.22
	pointsOfFunctionPlot[122].Y = -0.876

	pointsOfFunctionPlot[123].X = 1.23
	pointsOfFunctionPlot[123].Y = -0.895

	pointsOfFunctionPlot[124].X = 1.24
	pointsOfFunctionPlot[124].Y = -0.915

	pointsOfFunctionPlot[125].X = 1.25
	pointsOfFunctionPlot[125].Y = -0.934

	pointsOfFunctionPlot[126].X = 1.26
	pointsOfFunctionPlot[126].Y = -0.954

	pointsOfFunctionPlot[127].X = 1.27
	pointsOfFunctionPlot[127].Y = -0.973

	pointsOfFunctionPlot[128].X = 1.28
	pointsOfFunctionPlot[128].Y = -0.993

	pointsOfFunctionPlot[129].X = 1.29
	pointsOfFunctionPlot[129].Y = -1.012

	pointsOfFunctionPlot[130].X = 1.3
	pointsOfFunctionPlot[130].Y = -1.032

	pointsOfFunctionPlot[131].X = 1.31
	pointsOfFunctionPlot[131].Y = -1.052

	pointsOfFunctionPlot[132].X = 1.32
	pointsOfFunctionPlot[132].Y = -1.071

	pointsOfFunctionPlot[133].X = 1.33
	pointsOfFunctionPlot[133].Y = -1.091

	pointsOfFunctionPlot[134].X = 1.34
	pointsOfFunctionPlot[134].Y = -1.111

	pointsOfFunctionPlot[135].X = 1.35
	pointsOfFunctionPlot[135].Y = -1.131

	pointsOfFunctionPlot[136].X = 1.36
	pointsOfFunctionPlot[136].Y = -1.15

	pointsOfFunctionPlot[137].X = 1.37
	pointsOfFunctionPlot[137].Y = -1.17

	pointsOfFunctionPlot[138].X = 1.38
	pointsOfFunctionPlot[138].Y = -1.19

	pointsOfFunctionPlot[139].X = 1.39
	pointsOfFunctionPlot[139].Y = -1.21

	pointsOfFunctionPlot[140].X = 1.4
	pointsOfFunctionPlot[140].Y = -1.23

	pointsOfFunctionPlot[141].X = 1.41
	pointsOfFunctionPlot[141].Y = -1.249

	pointsOfFunctionPlot[142].X = 1.42
	pointsOfFunctionPlot[142].Y = -1.269

	pointsOfFunctionPlot[143].X = 1.43
	pointsOfFunctionPlot[143].Y = -1.289

	pointsOfFunctionPlot[144].X = 1.44
	pointsOfFunctionPlot[144].Y = -1.309

	pointsOfFunctionPlot[145].X = 1.45
	pointsOfFunctionPlot[145].Y = -1.329

	pointsOfFunctionPlot[146].X = 1.46
	pointsOfFunctionPlot[146].Y = -1.349

	pointsOfFunctionPlot[147].X = 1.47
	pointsOfFunctionPlot[147].Y = -1.369

	pointsOfFunctionPlot[148].X = 1.48
	pointsOfFunctionPlot[148].Y = -1.389

	pointsOfFunctionPlot[149].X = 1.49
	pointsOfFunctionPlot[149].Y = -1.409

	pointsOfFunctionPlot[150].X = 1.5
	pointsOfFunctionPlot[150].Y = -1.429

	pointsOfFunctionPlot[151].X = 1.51
	pointsOfFunctionPlot[151].Y = -1.449

	pointsOfFunctionPlot[152].X = 1.52
	pointsOfFunctionPlot[152].Y = -1.469

	pointsOfFunctionPlot[153].X = 1.53
	pointsOfFunctionPlot[153].Y = -1.489

	pointsOfFunctionPlot[154].X = 1.54
	pointsOfFunctionPlot[154].Y = -1.509

	pointsOfFunctionPlot[155].X = 1.55
	pointsOfFunctionPlot[155].Y = -1.529

	pointsOfFunctionPlot[156].X = 1.56
	pointsOfFunctionPlot[156].Y = -1.549

	pointsOfFunctionPlot[157].X = 1.57
	pointsOfFunctionPlot[157].Y = -1.569

	pointsOfFunctionPlot[158].X = 1.58
	pointsOfFunctionPlot[158].Y = -1.589

	pointsOfFunctionPlot[159].X = 1.59
	pointsOfFunctionPlot[159].Y = -1.609

	pointsOfFunctionPlot[160].X = 1.6
	pointsOfFunctionPlot[160].Y = -1.629

	pointsOfFunctionPlot[161].X = 1.61
	pointsOfFunctionPlot[161].Y = -1.649

	pointsOfFunctionPlot[162].X = 1.62
	pointsOfFunctionPlot[162].Y = -1.669

	pointsOfFunctionPlot[163].X = 1.63
	pointsOfFunctionPlot[163].Y = -1.689

	pointsOfFunctionPlot[164].X = 1.64
	pointsOfFunctionPlot[164].Y = -1.709

	pointsOfFunctionPlot[165].X = 1.65
	pointsOfFunctionPlot[165].Y = -1.729

	pointsOfFunctionPlot[166].X = 1.66
	pointsOfFunctionPlot[166].Y = -1.749

	pointsOfFunctionPlot[167].X = 1.67
	pointsOfFunctionPlot[167].Y = -1.769

	pointsOfFunctionPlot[168].X = 1.68
	pointsOfFunctionPlot[168].Y = -1.789

	pointsOfFunctionPlot[169].X = 1.69
	pointsOfFunctionPlot[169].Y = -1.808

	pointsOfFunctionPlot[170].X = 1.7
	pointsOfFunctionPlot[170].Y = -1.828

	pointsOfFunctionPlot[171].X = 1.71
	pointsOfFunctionPlot[171].Y = -1.848

	pointsOfFunctionPlot[172].X = 1.72
	pointsOfFunctionPlot[172].Y = -1.868

	pointsOfFunctionPlot[173].X = 1.73
	pointsOfFunctionPlot[173].Y = -1.888

	pointsOfFunctionPlot[174].X = 1.74
	pointsOfFunctionPlot[174].Y = -1.908

	pointsOfFunctionPlot[175].X = 1.75
	pointsOfFunctionPlot[175].Y = -1.928

	pointsOfFunctionPlot[176].X = 1.76
	pointsOfFunctionPlot[176].Y = -1.948

	pointsOfFunctionPlot[177].X = 1.77
	pointsOfFunctionPlot[177].Y = -1.967

	pointsOfFunctionPlot[178].X = 1.78
	pointsOfFunctionPlot[178].Y = -1.987

	pointsOfFunctionPlot[179].X = 1.79
	pointsOfFunctionPlot[179].Y = -2.007

	pointsOfFunctionPlot[180].X = 1.8
	pointsOfFunctionPlot[180].Y = -2.027

	pointsOfFunctionPlot[181].X = 1.81
	pointsOfFunctionPlot[181].Y = -2.046

	pointsOfFunctionPlot[182].X = 1.82
	pointsOfFunctionPlot[182].Y = -2.066

	pointsOfFunctionPlot[183].X = 1.83
	pointsOfFunctionPlot[183].Y = -2.086

	pointsOfFunctionPlot[184].X = 1.84
	pointsOfFunctionPlot[184].Y = -2.106

	pointsOfFunctionPlot[185].X = 1.85
	pointsOfFunctionPlot[185].Y = -2.125

	pointsOfFunctionPlot[186].X = 1.86
	pointsOfFunctionPlot[186].Y = -2.145

	pointsOfFunctionPlot[187].X = 1.87
	pointsOfFunctionPlot[187].Y = -2.164

	pointsOfFunctionPlot[188].X = 1.88
	pointsOfFunctionPlot[188].Y = -2.184

	pointsOfFunctionPlot[189].X = 1.89
	pointsOfFunctionPlot[189].Y = -2.203

	pointsOfFunctionPlot[190].X = 1.9
	pointsOfFunctionPlot[190].Y = -2.223

	pointsOfFunctionPlot[191].X = 1.91
	pointsOfFunctionPlot[191].Y = -2.242

	pointsOfFunctionPlot[192].X = 1.92
	pointsOfFunctionPlot[192].Y = -2.262

	pointsOfFunctionPlot[193].X = 1.93
	pointsOfFunctionPlot[193].Y = -2.281

	pointsOfFunctionPlot[194].X = 1.94
	pointsOfFunctionPlot[194].Y = -2.3

	pointsOfFunctionPlot[195].X = 1.95
	pointsOfFunctionPlot[195].Y = -2.32

	pointsOfFunctionPlot[196].X = 1.96
	pointsOfFunctionPlot[196].Y = -2.339

	pointsOfFunctionPlot[197].X = 1.97
	pointsOfFunctionPlot[197].Y = -2.358

	pointsOfFunctionPlot[198].X = 1.98
	pointsOfFunctionPlot[198].Y = -2.377

	pointsOfFunctionPlot[199].X = 1.99
	pointsOfFunctionPlot[199].Y = -2.397

	pointsOfFunctionPlot[200].X = 2.0
	pointsOfFunctionPlot[200].Y = -2.416

	pointsOfFunctionPlot[201].X = 2.01
	pointsOfFunctionPlot[201].Y = -2.435

	pointsOfFunctionPlot[202].X = 2.02
	pointsOfFunctionPlot[202].Y = -2.454

	pointsOfFunctionPlot[203].X = 2.03
	pointsOfFunctionPlot[203].Y = -2.473

	pointsOfFunctionPlot[204].X = 2.04
	pointsOfFunctionPlot[204].Y = -2.492

	pointsOfFunctionPlot[205].X = 2.05
	pointsOfFunctionPlot[205].Y = -2.511

	pointsOfFunctionPlot[206].X = 2.06
	pointsOfFunctionPlot[206].Y = -2.529

	pointsOfFunctionPlot[207].X = 2.07
	pointsOfFunctionPlot[207].Y = -2.548

	pointsOfFunctionPlot[208].X = 2.08
	pointsOfFunctionPlot[208].Y = -2.567

	pointsOfFunctionPlot[209].X = 2.09
	pointsOfFunctionPlot[209].Y = -2.586

	pointsOfFunctionPlot[210].X = 2.1
	pointsOfFunctionPlot[210].Y = -2.604

	pointsOfFunctionPlot[211].X = 2.11
	pointsOfFunctionPlot[211].Y = -2.623

	pointsOfFunctionPlot[212].X = 2.12
	pointsOfFunctionPlot[212].Y = -2.642

	pointsOfFunctionPlot[213].X = 2.13
	pointsOfFunctionPlot[213].Y = -2.66

	pointsOfFunctionPlot[214].X = 2.14
	pointsOfFunctionPlot[214].Y = -2.679

	pointsOfFunctionPlot[215].X = 2.15
	pointsOfFunctionPlot[215].Y = -2.697

	pointsOfFunctionPlot[216].X = 2.16
	pointsOfFunctionPlot[216].Y = -2.715

	pointsOfFunctionPlot[217].X = 2.17
	pointsOfFunctionPlot[217].Y = -2.734

	pointsOfFunctionPlot[218].X = 2.18
	pointsOfFunctionPlot[218].Y = -2.752

	pointsOfFunctionPlot[219].X = 2.19
	pointsOfFunctionPlot[219].Y = -2.77

	pointsOfFunctionPlot[220].X = 2.2
	pointsOfFunctionPlot[220].Y = -2.788

	pointsOfFunctionPlot[221].X = 2.21
	pointsOfFunctionPlot[221].Y = -2.806

	pointsOfFunctionPlot[222].X = 2.22
	pointsOfFunctionPlot[222].Y = -2.824

	pointsOfFunctionPlot[223].X = 2.23
	pointsOfFunctionPlot[223].Y = -2.842

	pointsOfFunctionPlot[224].X = 2.24
	pointsOfFunctionPlot[224].Y = -2.86

	pointsOfFunctionPlot[225].X = 2.25
	pointsOfFunctionPlot[225].Y = -2.878

	pointsOfFunctionPlot[226].X = 2.26
	pointsOfFunctionPlot[226].Y = -2.895

	pointsOfFunctionPlot[227].X = 2.27
	pointsOfFunctionPlot[227].Y = -2.913

	pointsOfFunctionPlot[228].X = 2.28
	pointsOfFunctionPlot[228].Y = -2.931

	pointsOfFunctionPlot[229].X = 2.29
	pointsOfFunctionPlot[229].Y = -2.948

	pointsOfFunctionPlot[230].X = 2.3
	pointsOfFunctionPlot[230].Y = -2.966

	pointsOfFunctionPlot[231].X = 2.31
	pointsOfFunctionPlot[231].Y = -2.983

	pointsOfFunctionPlot[232].X = 2.32
	pointsOfFunctionPlot[232].Y = -3.001

	pointsOfFunctionPlot[233].X = 2.33
	pointsOfFunctionPlot[233].Y = -3.018

	pointsOfFunctionPlot[234].X = 2.34
	pointsOfFunctionPlot[234].Y = -3.035

	pointsOfFunctionPlot[235].X = 2.35
	pointsOfFunctionPlot[235].Y = -3.052

	pointsOfFunctionPlot[236].X = 2.36
	pointsOfFunctionPlot[236].Y = -3.069

	pointsOfFunctionPlot[237].X = 2.37
	pointsOfFunctionPlot[237].Y = -3.086

	pointsOfFunctionPlot[238].X = 2.38
	pointsOfFunctionPlot[238].Y = -3.103

	pointsOfFunctionPlot[239].X = 2.39
	pointsOfFunctionPlot[239].Y = -3.12

	pointsOfFunctionPlot[240].X = 2.4
	pointsOfFunctionPlot[240].Y = -3.137

	pointsOfFunctionPlot[241].X = 2.41
	pointsOfFunctionPlot[241].Y = -3.154

	pointsOfFunctionPlot[242].X = 2.42
	pointsOfFunctionPlot[242].Y = -3.17

	pointsOfFunctionPlot[243].X = 2.43
	pointsOfFunctionPlot[243].Y = -3.187

	pointsOfFunctionPlot[244].X = 2.44
	pointsOfFunctionPlot[244].Y = -3.203

	pointsOfFunctionPlot[245].X = 2.45
	pointsOfFunctionPlot[245].Y = -3.22

	pointsOfFunctionPlot[246].X = 2.46
	pointsOfFunctionPlot[246].Y = -3.236

	pointsOfFunctionPlot[247].X = 2.47
	pointsOfFunctionPlot[247].Y = -3.252

	pointsOfFunctionPlot[248].X = 2.48
	pointsOfFunctionPlot[248].Y = -3.269

	pointsOfFunctionPlot[249].X = 2.49
	pointsOfFunctionPlot[249].Y = -3.285

	pointsOfFunctionPlot[250].X = 2.5
	pointsOfFunctionPlot[250].Y = -3.301

	pointsOfFunctionPlot[251].X = 2.51
	pointsOfFunctionPlot[251].Y = -3.317

	pointsOfFunctionPlot[252].X = 2.52
	pointsOfFunctionPlot[252].Y = -3.333

	pointsOfFunctionPlot[253].X = 2.53
	pointsOfFunctionPlot[253].Y = -3.348

	pointsOfFunctionPlot[254].X = 2.54
	pointsOfFunctionPlot[254].Y = -3.364

	pointsOfFunctionPlot[255].X = 2.55
	pointsOfFunctionPlot[255].Y = -3.38

	pointsOfFunctionPlot[256].X = 2.56
	pointsOfFunctionPlot[256].Y = -3.395

	pointsOfFunctionPlot[257].X = 2.57
	pointsOfFunctionPlot[257].Y = -3.411

	pointsOfFunctionPlot[258].X = 2.58
	pointsOfFunctionPlot[258].Y = -3.426

	pointsOfFunctionPlot[259].X = 2.59
	pointsOfFunctionPlot[259].Y = -3.441

	pointsOfFunctionPlot[260].X = 2.6
	pointsOfFunctionPlot[260].Y = -3.456

	pointsOfFunctionPlot[261].X = 2.61
	pointsOfFunctionPlot[261].Y = -3.472

	pointsOfFunctionPlot[262].X = 2.62
	pointsOfFunctionPlot[262].Y = -3.487

	pointsOfFunctionPlot[263].X = 2.63
	pointsOfFunctionPlot[263].Y = -3.502

	pointsOfFunctionPlot[264].X = 2.64
	pointsOfFunctionPlot[264].Y = -3.516

	pointsOfFunctionPlot[265].X = 2.65
	pointsOfFunctionPlot[265].Y = -3.531

	pointsOfFunctionPlot[266].X = 2.66
	pointsOfFunctionPlot[266].Y = -3.546

	pointsOfFunctionPlot[267].X = 2.67
	pointsOfFunctionPlot[267].Y = -3.56

	pointsOfFunctionPlot[268].X = 2.68
	pointsOfFunctionPlot[268].Y = -3.575

	pointsOfFunctionPlot[269].X = 2.69
	pointsOfFunctionPlot[269].Y = -3.589

	pointsOfFunctionPlot[270].X = 2.7
	pointsOfFunctionPlot[270].Y = -3.604

	pointsOfFunctionPlot[271].X = 2.71
	pointsOfFunctionPlot[271].Y = -3.618

	pointsOfFunctionPlot[272].X = 2.72
	pointsOfFunctionPlot[272].Y = -3.632

	pointsOfFunctionPlot[273].X = 2.73
	pointsOfFunctionPlot[273].Y = -3.646

	pointsOfFunctionPlot[274].X = 2.74
	pointsOfFunctionPlot[274].Y = -3.66

	pointsOfFunctionPlot[275].X = 2.75
	pointsOfFunctionPlot[275].Y = -3.674

	pointsOfFunctionPlot[276].X = 2.76
	pointsOfFunctionPlot[276].Y = -3.688

	pointsOfFunctionPlot[277].X = 2.77
	pointsOfFunctionPlot[277].Y = -3.701

	pointsOfFunctionPlot[278].X = 2.78
	pointsOfFunctionPlot[278].Y = -3.715

	pointsOfFunctionPlot[279].X = 2.79
	pointsOfFunctionPlot[279].Y = -3.728

	pointsOfFunctionPlot[280].X = 2.8
	pointsOfFunctionPlot[280].Y = -3.742

	pointsOfFunctionPlot[281].X = 2.81
	pointsOfFunctionPlot[281].Y = -3.755

	pointsOfFunctionPlot[282].X = 2.82
	pointsOfFunctionPlot[282].Y = -3.768

	pointsOfFunctionPlot[283].X = 2.83
	pointsOfFunctionPlot[283].Y = -3.781

	pointsOfFunctionPlot[284].X = 2.84
	pointsOfFunctionPlot[284].Y = -3.794

	pointsOfFunctionPlot[285].X = 2.85
	pointsOfFunctionPlot[285].Y = -3.807

	pointsOfFunctionPlot[286].X = 2.86
	pointsOfFunctionPlot[286].Y = -3.82

	pointsOfFunctionPlot[287].X = 2.87
	pointsOfFunctionPlot[287].Y = -3.833

	pointsOfFunctionPlot[288].X = 2.88
	pointsOfFunctionPlot[288].Y = -3.846

	pointsOfFunctionPlot[289].X = 2.89
	pointsOfFunctionPlot[289].Y = -3.858

	pointsOfFunctionPlot[290].X = 2.9
	pointsOfFunctionPlot[290].Y = -3.871

	pointsOfFunctionPlot[291].X = 2.91
	pointsOfFunctionPlot[291].Y = -3.883

	pointsOfFunctionPlot[292].X = 2.92
	pointsOfFunctionPlot[292].Y = -3.895

	pointsOfFunctionPlot[293].X = 2.93
	pointsOfFunctionPlot[293].Y = -3.907

	pointsOfFunctionPlot[294].X = 2.94
	pointsOfFunctionPlot[294].Y = -3.919

	pointsOfFunctionPlot[295].X = 2.95
	pointsOfFunctionPlot[295].Y = -3.931

	pointsOfFunctionPlot[296].X = 2.96
	pointsOfFunctionPlot[296].Y = -3.943

	pointsOfFunctionPlot[297].X = 2.97
	pointsOfFunctionPlot[297].Y = -3.955

	pointsOfFunctionPlot[298].X = 2.98
	pointsOfFunctionPlot[298].Y = -3.967

	pointsOfFunctionPlot[299].X = 2.99
	pointsOfFunctionPlot[299].Y = -3.978

	pointsOfFunctionPlot[300].X = 3.0
	pointsOfFunctionPlot[300].Y = -3.99

	pointsOfFunctionPlot[301].X = 3.01
	pointsOfFunctionPlot[301].Y = -4.001

	pointsOfFunctionPlot[302].X = 3.02
	pointsOfFunctionPlot[302].Y = -4.012

	pointsOfFunctionPlot[303].X = 3.03
	pointsOfFunctionPlot[303].Y = -4.023

	pointsOfFunctionPlot[304].X = 3.04
	pointsOfFunctionPlot[304].Y = -4.034

	pointsOfFunctionPlot[305].X = 3.05
	pointsOfFunctionPlot[305].Y = -4.045

	pointsOfFunctionPlot[306].X = 3.06
	pointsOfFunctionPlot[306].Y = -4.056

	pointsOfFunctionPlot[307].X = 3.07
	pointsOfFunctionPlot[307].Y = -4.067

	pointsOfFunctionPlot[308].X = 3.08
	pointsOfFunctionPlot[308].Y = -4.078

	pointsOfFunctionPlot[309].X = 3.09
	pointsOfFunctionPlot[309].Y = -4.088

	pointsOfFunctionPlot[310].X = 3.1
	pointsOfFunctionPlot[310].Y = -4.099

	pointsOfFunctionPlot[311].X = 3.11
	pointsOfFunctionPlot[311].Y = -4.109

	pointsOfFunctionPlot[312].X = 3.12
	pointsOfFunctionPlot[312].Y = -4.119

	pointsOfFunctionPlot[313].X = 3.13
	pointsOfFunctionPlot[313].Y = -4.129

	pointsOfFunctionPlot[314].X = 3.14
	pointsOfFunctionPlot[314].Y = -4.14

	pointsOfFunctionPlot[315].X = 3.15
	pointsOfFunctionPlot[315].Y = -4.15

	pointsOfFunctionPlot[316].X = 3.16
	pointsOfFunctionPlot[316].Y = -4.159

	pointsOfFunctionPlot[317].X = 3.17
	pointsOfFunctionPlot[317].Y = -4.169

	pointsOfFunctionPlot[318].X = 3.18
	pointsOfFunctionPlot[318].Y = -4.179

	pointsOfFunctionPlot[319].X = 3.19
	pointsOfFunctionPlot[319].Y = -4.188

	pointsOfFunctionPlot[320].X = 3.2
	pointsOfFunctionPlot[320].Y = -4.198

	pointsOfFunctionPlot[321].X = 3.21
	pointsOfFunctionPlot[321].Y = -4.207

	pointsOfFunctionPlot[322].X = 3.22
	pointsOfFunctionPlot[322].Y = -4.216

	pointsOfFunctionPlot[323].X = 3.23
	pointsOfFunctionPlot[323].Y = -4.226

	pointsOfFunctionPlot[324].X = 3.24
	pointsOfFunctionPlot[324].Y = -4.235

	pointsOfFunctionPlot[325].X = 3.25
	pointsOfFunctionPlot[325].Y = -4.244

	pointsOfFunctionPlot[326].X = 3.26
	pointsOfFunctionPlot[326].Y = -4.253

	pointsOfFunctionPlot[327].X = 3.27
	pointsOfFunctionPlot[327].Y = -4.261

	pointsOfFunctionPlot[328].X = 3.28
	pointsOfFunctionPlot[328].Y = -4.27

	pointsOfFunctionPlot[329].X = 3.29
	pointsOfFunctionPlot[329].Y = -4.279

	pointsOfFunctionPlot[330].X = 3.3
	pointsOfFunctionPlot[330].Y = -4.287

	pointsOfFunctionPlot[331].X = 3.31
	pointsOfFunctionPlot[331].Y = -4.295

	pointsOfFunctionPlot[332].X = 3.32
	pointsOfFunctionPlot[332].Y = -4.304

	pointsOfFunctionPlot[333].X = 3.33
	pointsOfFunctionPlot[333].Y = -4.312

	pointsOfFunctionPlot[334].X = 3.34
	pointsOfFunctionPlot[334].Y = -4.32

	pointsOfFunctionPlot[335].X = 3.35
	pointsOfFunctionPlot[335].Y = -4.328

	pointsOfFunctionPlot[336].X = 3.36
	pointsOfFunctionPlot[336].Y = -4.336

	pointsOfFunctionPlot[337].X = 3.37
	pointsOfFunctionPlot[337].Y = -4.334

	pointsOfFunctionPlot[338].X = 3.38
	pointsOfFunctionPlot[338].Y = -4.351

	pointsOfFunctionPlot[339].X = 3.39
	pointsOfFunctionPlot[339].Y = -4.359

	pointsOfFunctionPlot[340].X = 3.4
	pointsOfFunctionPlot[340].Y = -4.366

	pointsOfFunctionPlot[341].X = 3.41
	pointsOfFunctionPlot[341].Y = -4.374

	pointsOfFunctionPlot[342].X = 3.42
	pointsOfFunctionPlot[342].Y = -4.381

	pointsOfFunctionPlot[343].X = 3.43
	pointsOfFunctionPlot[343].Y = -4.388

	pointsOfFunctionPlot[344].X = 3.44
	pointsOfFunctionPlot[344].Y = -4.395

	pointsOfFunctionPlot[345].X = 3.45
	pointsOfFunctionPlot[345].Y = -4.402

	pointsOfFunctionPlot[346].X = 3.46
	pointsOfFunctionPlot[346].Y = -4.409

	pointsOfFunctionPlot[347].X = 3.47
	pointsOfFunctionPlot[347].Y = -4.416

	pointsOfFunctionPlot[348].X = 3.48
	pointsOfFunctionPlot[348].Y = -4.423

	pointsOfFunctionPlot[349].X = 3.49
	pointsOfFunctionPlot[349].Y = -4.429

	pointsOfFunctionPlot[350].X = 3.5
	pointsOfFunctionPlot[350].Y = -4.436

	pointsOfFunctionPlot[351].X = 3.51
	pointsOfFunctionPlot[351].Y = -4.442

	pointsOfFunctionPlot[352].X = 3.52
	pointsOfFunctionPlot[352].Y = -4.449

	pointsOfFunctionPlot[353].X = 3.53
	pointsOfFunctionPlot[353].Y = -4.455

	pointsOfFunctionPlot[354].X = 3.54
	pointsOfFunctionPlot[354].Y = -4.461

	pointsOfFunctionPlot[355].X = 3.55
	pointsOfFunctionPlot[355].Y = -4.467

	pointsOfFunctionPlot[356].X = 3.56
	pointsOfFunctionPlot[356].Y = -4.473

	pointsOfFunctionPlot[357].X = 3.57
	pointsOfFunctionPlot[357].Y = -4.479

	pointsOfFunctionPlot[358].X = 3.58
	pointsOfFunctionPlot[358].Y = -4.485

	pointsOfFunctionPlot[359].X = 3.59
	pointsOfFunctionPlot[359].Y = -4.491

	pointsOfFunctionPlot[360].X = 3.6
	pointsOfFunctionPlot[360].Y = -4.496

	pointsOfFunctionPlot[361].X = 3.61
	pointsOfFunctionPlot[361].Y = -4.502

	pointsOfFunctionPlot[362].X = 3.62
	pointsOfFunctionPlot[362].Y = -4.507

	pointsOfFunctionPlot[363].X = 3.63
	pointsOfFunctionPlot[363].Y = -4.513

	pointsOfFunctionPlot[364].X = 3.64
	pointsOfFunctionPlot[364].Y = -4.518

	pointsOfFunctionPlot[365].X = 3.65
	pointsOfFunctionPlot[365].Y = -4.523

	pointsOfFunctionPlot[366].X = 3.66
	pointsOfFunctionPlot[366].Y = -4.528

	pointsOfFunctionPlot[367].X = 3.67
	pointsOfFunctionPlot[367].Y = -4.533

	pointsOfFunctionPlot[368].X = 3.68
	pointsOfFunctionPlot[368].Y = -4.538

	pointsOfFunctionPlot[369].X = 3.69
	pointsOfFunctionPlot[369].Y = -4.543

	pointsOfFunctionPlot[370].X = 3.7
	pointsOfFunctionPlot[370].Y = -4.548

	pointsOfFunctionPlot[371].X = 3.71
	pointsOfFunctionPlot[371].Y = -4.552

	pointsOfFunctionPlot[372].X = 3.72
	pointsOfFunctionPlot[372].Y = -4.557

	pointsOfFunctionPlot[373].X = 3.73
	pointsOfFunctionPlot[373].Y = -4.561

	pointsOfFunctionPlot[374].X = 3.74
	pointsOfFunctionPlot[374].Y = -4.566

	pointsOfFunctionPlot[375].X = 3.75
	pointsOfFunctionPlot[375].Y = -4.57

	pointsOfFunctionPlot[376].X = 3.76
	pointsOfFunctionPlot[376].Y = -4.574

	pointsOfFunctionPlot[377].X = 3.77
	pointsOfFunctionPlot[377].Y = -4.579

	pointsOfFunctionPlot[378].X = 3.78
	pointsOfFunctionPlot[378].Y = -4.583

	pointsOfFunctionPlot[379].X = 3.79
	pointsOfFunctionPlot[379].Y = -4.587

	pointsOfFunctionPlot[380].X = 3.8
	pointsOfFunctionPlot[380].Y = -4.591

	pointsOfFunctionPlot[381].X = 3.81
	pointsOfFunctionPlot[381].Y = -4.594

	pointsOfFunctionPlot[382].X = 3.82
	pointsOfFunctionPlot[382].Y = -4.598

	pointsOfFunctionPlot[383].X = 3.83
	pointsOfFunctionPlot[383].Y = -4.602

	pointsOfFunctionPlot[384].X = 3.84
	pointsOfFunctionPlot[384].Y = -4.605

	pointsOfFunctionPlot[385].X = 3.85
	pointsOfFunctionPlot[385].Y = -4.609

	pointsOfFunctionPlot[386].X = 3.86
	pointsOfFunctionPlot[386].Y = -4.612

	pointsOfFunctionPlot[387].X = 3.87
	pointsOfFunctionPlot[387].Y = -4.616

	pointsOfFunctionPlot[388].X = 3.88
	pointsOfFunctionPlot[388].Y = -4.619

	pointsOfFunctionPlot[389].X = 3.89
	pointsOfFunctionPlot[389].Y = -4.622

	pointsOfFunctionPlot[390].X = 3.9
	pointsOfFunctionPlot[390].Y = -4.625

	pointsOfFunctionPlot[391].X = 3.91
	pointsOfFunctionPlot[391].Y = -4.629

	pointsOfFunctionPlot[392].X = 3.92
	pointsOfFunctionPlot[392].Y = -4.632

	pointsOfFunctionPlot[393].X = 3.93
	pointsOfFunctionPlot[393].Y = -4.635

	pointsOfFunctionPlot[394].X = 3.94
	pointsOfFunctionPlot[394].Y = -4.637

	pointsOfFunctionPlot[395].X = 3.95
	pointsOfFunctionPlot[395].Y = -4.64

	pointsOfFunctionPlot[396].X = 3.96
	pointsOfFunctionPlot[396].Y = -4.643

	pointsOfFunctionPlot[397].X = 3.97
	pointsOfFunctionPlot[397].Y = -4.646

	pointsOfFunctionPlot[398].X = 3.98
	pointsOfFunctionPlot[398].Y = -4.648

	pointsOfFunctionPlot[399].X = 3.99
	pointsOfFunctionPlot[399].Y = -4.651

	pointsOfFunctionPlot[400].X = 4.0
	pointsOfFunctionPlot[400].Y = -4.653

	pointsOfFunctionPlot[401].X = 4.01
	pointsOfFunctionPlot[401].Y = -4.656

	pointsOfFunctionPlot[402].X = 4.02
	pointsOfFunctionPlot[402].Y = -4.658

	pointsOfFunctionPlot[403].X = 4.03
	pointsOfFunctionPlot[403].Y = -4.66

	pointsOfFunctionPlot[404].X = 4.04
	pointsOfFunctionPlot[404].Y = -4.662

	pointsOfFunctionPlot[405].X = 4.05
	pointsOfFunctionPlot[405].Y = -4.665

	pointsOfFunctionPlot[406].X = 4.06
	pointsOfFunctionPlot[406].Y = -4.667

	pointsOfFunctionPlot[407].X = 4.07
	pointsOfFunctionPlot[407].Y = -4.669

	pointsOfFunctionPlot[408].X = 4.08
	pointsOfFunctionPlot[408].Y = -4.671

	pointsOfFunctionPlot[409].X = 4.09
	pointsOfFunctionPlot[409].Y = -4.673

	pointsOfFunctionPlot[410].X = 4.1
	pointsOfFunctionPlot[410].Y = -4.674

	pointsOfFunctionPlot[411].X = 4.11
	pointsOfFunctionPlot[411].Y = -4.676

	pointsOfFunctionPlot[412].X = 4.12
	pointsOfFunctionPlot[412].Y = -4.678

	pointsOfFunctionPlot[413].X = 4.13
	pointsOfFunctionPlot[413].Y = -4.68

	pointsOfFunctionPlot[414].X = 4.14
	pointsOfFunctionPlot[414].Y = -4.681

	pointsOfFunctionPlot[415].X = 4.15
	pointsOfFunctionPlot[415].Y = -4.683

	pointsOfFunctionPlot[416].X = 4.16
	pointsOfFunctionPlot[416].Y = -4.684

	pointsOfFunctionPlot[417].X = 4.17
	pointsOfFunctionPlot[417].Y = -4.686

	pointsOfFunctionPlot[418].X = 4.18
	pointsOfFunctionPlot[418].Y = -4.687

	pointsOfFunctionPlot[419].X = 4.19
	pointsOfFunctionPlot[419].Y = -4.689

	pointsOfFunctionPlot[420].X = 4.2
	pointsOfFunctionPlot[420].Y = -4.69

	pointsOfFunctionPlot[421].X = 4.21
	pointsOfFunctionPlot[421].Y = -4.69

	pointsOfFunctionPlot[422].X = 4.22
	pointsOfFunctionPlot[422].Y = -4.692

	pointsOfFunctionPlot[423].X = 4.23
	pointsOfFunctionPlot[423].Y = -4.693

	pointsOfFunctionPlot[424].X = 4.24
	pointsOfFunctionPlot[424].Y = -4.695

	pointsOfFunctionPlot[425].X = 4.25
	pointsOfFunctionPlot[425].Y = -4.696

	pointsOfFunctionPlot[426].X = 4.26
	pointsOfFunctionPlot[426].Y = -4.697

	pointsOfFunctionPlot[427].X = 4.27
	pointsOfFunctionPlot[427].Y = -4.698

	pointsOfFunctionPlot[428].X = 4.28
	pointsOfFunctionPlot[428].Y = -4.699

	pointsOfFunctionPlot[429].X = 4.29
	pointsOfFunctionPlot[429].Y = -4.699

	pointsOfFunctionPlot[430].X = 4.3
	pointsOfFunctionPlot[430].Y = -4.7

	pointsOfFunctionPlot[431].X = 4.31
	pointsOfFunctionPlot[431].Y = -4.701

	pointsOfFunctionPlot[432].X = 4.32
	pointsOfFunctionPlot[432].Y = -4.702

	pointsOfFunctionPlot[433].X = 4.33
	pointsOfFunctionPlot[433].Y = -4.703

	pointsOfFunctionPlot[434].X = 4.34
	pointsOfFunctionPlot[434].Y = -4.703

	pointsOfFunctionPlot[435].X = 4.35
	pointsOfFunctionPlot[435].Y = -4.704

	pointsOfFunctionPlot[436].X = 4.36
	pointsOfFunctionPlot[436].Y = -4.705

	pointsOfFunctionPlot[437].X = 4.37
	pointsOfFunctionPlot[437].Y = -4.705

	pointsOfFunctionPlot[438].X = 4.38
	pointsOfFunctionPlot[438].Y = -4.706

	pointsOfFunctionPlot[439].X = 4.39
	pointsOfFunctionPlot[439].Y = -4.706

	pointsOfFunctionPlot[440].X = 4.4
	pointsOfFunctionPlot[440].Y = -4.707

	pointsOfFunctionPlot[441].X = 4.41
	pointsOfFunctionPlot[441].Y = -4.707

	pointsOfFunctionPlot[442].X = 4.42
	pointsOfFunctionPlot[442].Y = -4.708

	pointsOfFunctionPlot[443].X = 4.43
	pointsOfFunctionPlot[443].Y = -4.708

	pointsOfFunctionPlot[444].X = 4.44
	pointsOfFunctionPlot[444].Y = -4.709

	pointsOfFunctionPlot[445].X = 4.45
	pointsOfFunctionPlot[445].Y = -4.709

	pointsOfFunctionPlot[446].X = 4.46
	pointsOfFunctionPlot[446].Y = -4.709

	pointsOfFunctionPlot[447].X = 4.47
	pointsOfFunctionPlot[447].Y = -4.71

	pointsOfFunctionPlot[448].X = 4.48
	pointsOfFunctionPlot[448].Y = -4.71

	pointsOfFunctionPlot[449].X = 4.49
	pointsOfFunctionPlot[449].Y = -4.71

	pointsOfFunctionPlot[450].X = 4.5
	pointsOfFunctionPlot[450].Y = -4.71

	pointsOfFunctionPlot[451].X = 4.51
	pointsOfFunctionPlot[451].Y = -4.711

	pointsOfFunctionPlot[452].X = 4.52
	pointsOfFunctionPlot[452].Y = -4.711

	pointsOfFunctionPlot[453].X = 4.53
	pointsOfFunctionPlot[453].Y = -4.711

	pointsOfFunctionPlot[454].X = 4.54
	pointsOfFunctionPlot[454].Y = -4.711

	pointsOfFunctionPlot[455].X = 4.55
	pointsOfFunctionPlot[455].Y = -4.711

	pointsOfFunctionPlot[456].X = 4.56
	pointsOfFunctionPlot[456].Y = -4.711

	pointsOfFunctionPlot[457].X = 4.57
	pointsOfFunctionPlot[457].Y = -4.711

	pointsOfFunctionPlot[458].X = 4.58
	pointsOfFunctionPlot[458].Y = -4.712

	pointsOfFunctionPlot[459].X = 4.59
	pointsOfFunctionPlot[459].Y = -4.712

	pointsOfFunctionPlot[460].X = 4.6
	pointsOfFunctionPlot[460].Y = -4.712

	pointsOfFunctionPlot[461].X = 4.61
	pointsOfFunctionPlot[461].Y = -4.712

	pointsOfFunctionPlot[462].X = 4.62
	pointsOfFunctionPlot[462].Y = -4.712

	pointsOfFunctionPlot[463].X = 4.63
	pointsOfFunctionPlot[463].Y = -4.712

	pointsOfFunctionPlot[464].X = 4.64
	pointsOfFunctionPlot[464].Y = -4.712

	pointsOfFunctionPlot[465].X = 4.65
	pointsOfFunctionPlot[465].Y = -4.712

	pointsOfFunctionPlot[466].X = 4.66
	pointsOfFunctionPlot[466].Y = -4.712

	pointsOfFunctionPlot[467].X = 4.67
	pointsOfFunctionPlot[467].Y = -4.712

	pointsOfFunctionPlot[468].X = 4.68
	pointsOfFunctionPlot[468].Y = -4.712

	pointsOfFunctionPlot[469].X = 4.69
	pointsOfFunctionPlot[469].Y = -4.712

	pointsOfFunctionPlot[470].X = 4.7
	pointsOfFunctionPlot[470].Y = -4.712

	pointsOfFunctionPlot[471].X = 4.71
	pointsOfFunctionPlot[471].Y = -4.712

	pointsOfFunctionPlot[472].X = 4.72
	pointsOfFunctionPlot[472].Y = -4.712

	pointsOfFunctionPlot[473].X = 4.73
	pointsOfFunctionPlot[473].Y = -4.712

	pointsOfFunctionPlot[474].X = 4.74
	pointsOfFunctionPlot[474].Y = -4.712

	pointsOfFunctionPlot[475].X = 4.75
	pointsOfFunctionPlot[475].Y = -4.712

	pointsOfFunctionPlot[476].X = 4.76
	pointsOfFunctionPlot[476].Y = -4.712

	pointsOfFunctionPlot[477].X = 4.77
	pointsOfFunctionPlot[477].Y = -4.712

	pointsOfFunctionPlot[478].X = 4.78
	pointsOfFunctionPlot[478].Y = -4.712

	pointsOfFunctionPlot[479].X = 4.79
	pointsOfFunctionPlot[479].Y = -4.712

	pointsOfFunctionPlot[480].X = 4.8
	pointsOfFunctionPlot[480].Y = -4.712

	pointsOfFunctionPlot[481].X = 4.81
	pointsOfFunctionPlot[481].Y = -4.712

	pointsOfFunctionPlot[482].X = 4.82
	pointsOfFunctionPlot[482].Y = -4.712

	pointsOfFunctionPlot[483].X = 4.83
	pointsOfFunctionPlot[483].Y = -4.712

	pointsOfFunctionPlot[484].X = 4.84
	pointsOfFunctionPlot[484].Y = -4.712

	pointsOfFunctionPlot[485].X = 4.85
	pointsOfFunctionPlot[485].Y = -4.712

	pointsOfFunctionPlot[486].X = 4.86
	pointsOfFunctionPlot[486].Y = -4.712

	pointsOfFunctionPlot[487].X = 4.87
	pointsOfFunctionPlot[487].Y = -4.713

	pointsOfFunctionPlot[488].X = 4.88
	pointsOfFunctionPlot[488].Y = -4.713

	pointsOfFunctionPlot[489].X = 4.89
	pointsOfFunctionPlot[489].Y = -4.713

	pointsOfFunctionPlot[490].X = 4.9
	pointsOfFunctionPlot[490].Y = -4.713

	pointsOfFunctionPlot[491].X = 4.91
	pointsOfFunctionPlot[491].Y = -4.713

	pointsOfFunctionPlot[492].X = 4.92
	pointsOfFunctionPlot[492].Y = -4.713

	pointsOfFunctionPlot[493].X = 4.93
	pointsOfFunctionPlot[493].Y = -4.714

	pointsOfFunctionPlot[494].X = 4.94
	pointsOfFunctionPlot[494].Y = -4.714

	pointsOfFunctionPlot[495].X = 4.95
	pointsOfFunctionPlot[495].Y = -4.714

	pointsOfFunctionPlot[496].X = 4.96
	pointsOfFunctionPlot[496].Y = -4.714

	pointsOfFunctionPlot[497].X = 4.97
	pointsOfFunctionPlot[497].Y = -4.715

	pointsOfFunctionPlot[498].X = 4.98
	pointsOfFunctionPlot[498].Y = -4.715

	pointsOfFunctionPlot[499].X = 4.99
	pointsOfFunctionPlot[499].Y = -4.715

	pointsOfFunctionPlot[500].X = 5.0
	pointsOfFunctionPlot[500].Y = -4.716

	pointsOfFunctionPlot[501].X = 5.01
	pointsOfFunctionPlot[501].Y = -4.716

	pointsOfFunctionPlot[502].X = 5.02
	pointsOfFunctionPlot[502].Y = -4.717

	pointsOfFunctionPlot[503].X = 5.03
	pointsOfFunctionPlot[503].Y = -4.717

	pointsOfFunctionPlot[504].X = 5.04
	pointsOfFunctionPlot[504].Y = -4.718

	pointsOfFunctionPlot[505].X = 5.05
	pointsOfFunctionPlot[505].Y = -4.718

	pointsOfFunctionPlot[506].X = 5.06
	pointsOfFunctionPlot[506].Y = -4.719

	pointsOfFunctionPlot[507].X = 5.07
	pointsOfFunctionPlot[507].Y = -4.72

	pointsOfFunctionPlot[508].X = 5.08
	pointsOfFunctionPlot[508].Y = -4.72

	pointsOfFunctionPlot[509].X = 5.09
	pointsOfFunctionPlot[509].Y = -4.721

	pointsOfFunctionPlot[510].X = 5.1
	pointsOfFunctionPlot[510].Y = -4.722

	pointsOfFunctionPlot[511].X = 5.11
	pointsOfFunctionPlot[511].Y = -4.722

	pointsOfFunctionPlot[512].X = 5.12
	pointsOfFunctionPlot[512].Y = -4.723

	pointsOfFunctionPlot[513].X = 5.13
	pointsOfFunctionPlot[513].Y = -4.724

	pointsOfFunctionPlot[514].X = 5.14
	pointsOfFunctionPlot[514].Y = -4.725

	pointsOfFunctionPlot[515].X = 5.15
	pointsOfFunctionPlot[515].Y = -4.726

	pointsOfFunctionPlot[516].X = 5.16
	pointsOfFunctionPlot[516].Y = -4.727

	pointsOfFunctionPlot[517].X = 5.17
	pointsOfFunctionPlot[517].Y = -4.728

	pointsOfFunctionPlot[518].X = 5.18
	pointsOfFunctionPlot[518].Y = -4.729

	pointsOfFunctionPlot[519].X = 5.19
	pointsOfFunctionPlot[519].Y = -4.73

	pointsOfFunctionPlot[520].X = 5.2
	pointsOfFunctionPlot[520].Y = -4.731

	pointsOfFunctionPlot[521].X = 5.21
	pointsOfFunctionPlot[521].Y = -4.732

	pointsOfFunctionPlot[522].X = 5.22
	pointsOfFunctionPlot[522].Y = -4.733

	pointsOfFunctionPlot[523].X = 5.23
	pointsOfFunctionPlot[523].Y = -4.735

	pointsOfFunctionPlot[524].X = 5.24
	pointsOfFunctionPlot[524].Y = -4.736

	pointsOfFunctionPlot[525].X = 5.25
	pointsOfFunctionPlot[525].Y = -4.737

	pointsOfFunctionPlot[526].X = 5.26
	pointsOfFunctionPlot[526].Y = -4.739

	pointsOfFunctionPlot[527].X = 5.27
	pointsOfFunctionPlot[527].Y = -4.74

	pointsOfFunctionPlot[528].X = 5.28
	pointsOfFunctionPlot[528].Y = -4.742

	pointsOfFunctionPlot[529].X = 5.29
	pointsOfFunctionPlot[529].Y = -4.744

	pointsOfFunctionPlot[530].X = 5.3
	pointsOfFunctionPlot[530].Y = -4.745

	pointsOfFunctionPlot[531].X = 5.31
	pointsOfFunctionPlot[531].Y = -4.747

	pointsOfFunctionPlot[532].X = 5.32
	pointsOfFunctionPlot[532].Y = -4.749

	pointsOfFunctionPlot[533].X = 5.33
	pointsOfFunctionPlot[533].Y = -4.75

	pointsOfFunctionPlot[534].X = 5.34
	pointsOfFunctionPlot[534].Y = -4.752

	pointsOfFunctionPlot[535].X = 5.35
	pointsOfFunctionPlot[535].Y = -4.754

	pointsOfFunctionPlot[536].X = 5.36
	pointsOfFunctionPlot[536].Y = -4.756

	pointsOfFunctionPlot[537].X = 5.37
	pointsOfFunctionPlot[537].Y = -4.758

	pointsOfFunctionPlot[538].X = 5.38
	pointsOfFunctionPlot[538].Y = -4.76

	pointsOfFunctionPlot[539].X = 5.39
	pointsOfFunctionPlot[539].Y = -4.763

	pointsOfFunctionPlot[540].X = 5.4
	pointsOfFunctionPlot[540].Y = -4.765

	pointsOfFunctionPlot[541].X = 5.41
	pointsOfFunctionPlot[541].Y = -4.767

	pointsOfFunctionPlot[542].X = 5.42
	pointsOfFunctionPlot[542].Y = -4.77

	pointsOfFunctionPlot[543].X = 5.43
	pointsOfFunctionPlot[543].Y = -4.772

	pointsOfFunctionPlot[544].X = 5.44
	pointsOfFunctionPlot[544].Y = -4.774

	pointsOfFunctionPlot[545].X = 5.45
	pointsOfFunctionPlot[545].Y = -4.777

	pointsOfFunctionPlot[546].X = 5.46
	pointsOfFunctionPlot[546].Y = -4.78

	pointsOfFunctionPlot[547].X = 5.47
	pointsOfFunctionPlot[547].Y = -4.782

	pointsOfFunctionPlot[548].X = 5.48
	pointsOfFunctionPlot[548].Y = -4.785

	pointsOfFunctionPlot[549].X = 5.49
	pointsOfFunctionPlot[549].Y = -4.788

	pointsOfFunctionPlot[550].X = 5.5
	pointsOfFunctionPlot[550].Y = -4.791

	pointsOfFunctionPlot[551].X = 5.51
	pointsOfFunctionPlot[551].Y = -4.794

	pointsOfFunctionPlot[552].X = 5.52
	pointsOfFunctionPlot[552].Y = -4.797

	pointsOfFunctionPlot[553].X = 5.53
	pointsOfFunctionPlot[553].Y = -4.8

	pointsOfFunctionPlot[554].X = 5.54
	pointsOfFunctionPlot[554].Y = -4.803

	pointsOfFunctionPlot[555].X = 5.55
	pointsOfFunctionPlot[555].Y = -4.807

	pointsOfFunctionPlot[556].X = 5.56
	pointsOfFunctionPlot[556].Y = -4.81

	pointsOfFunctionPlot[557].X = 5.57
	pointsOfFunctionPlot[557].Y = -4.813

	pointsOfFunctionPlot[558].X = 5.58
	pointsOfFunctionPlot[558].Y = -4.817

	pointsOfFunctionPlot[559].X = 5.59
	pointsOfFunctionPlot[559].Y = -4.82

	pointsOfFunctionPlot[560].X = 5.6
	pointsOfFunctionPlot[560].Y = -4.824

	pointsOfFunctionPlot[561].X = 5.61
	pointsOfFunctionPlot[561].Y = -4.828

	pointsOfFunctionPlot[562].X = 5.62
	pointsOfFunctionPlot[562].Y = -4.832

	pointsOfFunctionPlot[563].X = 5.63
	pointsOfFunctionPlot[563].Y = -4.835

	pointsOfFunctionPlot[564].X = 5.64
	pointsOfFunctionPlot[564].Y = -4.839

	pointsOfFunctionPlot[565].X = 5.65
	pointsOfFunctionPlot[565].Y = -4.843

	pointsOfFunctionPlot[566].X = 5.66
	pointsOfFunctionPlot[566].Y = -4.848

	pointsOfFunctionPlot[567].X = 5.67
	pointsOfFunctionPlot[567].Y = -4.852

	pointsOfFunctionPlot[568].X = 5.68
	pointsOfFunctionPlot[568].Y = -4.856

	pointsOfFunctionPlot[569].X = 5.69
	pointsOfFunctionPlot[569].Y = -4.86

	pointsOfFunctionPlot[570].X = 5.7
	pointsOfFunctionPlot[570].Y = -4.865

	pointsOfFunctionPlot[571].X = 5.71
	pointsOfFunctionPlot[571].Y = -4.869

	pointsOfFunctionPlot[572].X = 5.72
	pointsOfFunctionPlot[572].Y = -4.874

	pointsOfFunctionPlot[573].X = 5.73
	pointsOfFunctionPlot[573].Y = -4.879

	pointsOfFunctionPlot[574].X = 5.74
	pointsOfFunctionPlot[574].Y = -4.883

	pointsOfFunctionPlot[575].X = 5.75
	pointsOfFunctionPlot[575].Y = -4.888

	pointsOfFunctionPlot[576].X = 5.76
	pointsOfFunctionPlot[576].Y = -4.893

	pointsOfFunctionPlot[577].X = 5.77
	pointsOfFunctionPlot[577].Y = -4.898

	pointsOfFunctionPlot[578].X = 5.78
	pointsOfFunctionPlot[578].Y = -4.903

	pointsOfFunctionPlot[579].X = 5.79
	pointsOfFunctionPlot[579].Y = -4.909

	pointsOfFunctionPlot[580].X = 5.8
	pointsOfFunctionPlot[580].Y = -4.914

	pointsOfFunctionPlot[581].X = 5.81
	pointsOfFunctionPlot[581].Y = -4.919

	pointsOfFunctionPlot[582].X = 5.82
	pointsOfFunctionPlot[582].Y = -4.925

	pointsOfFunctionPlot[583].X = 5.83
	pointsOfFunctionPlot[583].Y = -4.93

	pointsOfFunctionPlot[584].X = 5.84
	pointsOfFunctionPlot[584].Y = -4.936

	pointsOfFunctionPlot[585].X = 5.85
	pointsOfFunctionPlot[585].Y = -4.942

	pointsOfFunctionPlot[586].X = 5.86
	pointsOfFunctionPlot[586].Y = -4.948

	pointsOfFunctionPlot[587].X = 5.87
	pointsOfFunctionPlot[587].Y = -4.954

	pointsOfFunctionPlot[588].X = 5.88
	pointsOfFunctionPlot[588].Y = -4.96

	pointsOfFunctionPlot[589].X = 5.89
	pointsOfFunctionPlot[589].Y = -4.966

	pointsOfFunctionPlot[590].X = 5.9
	pointsOfFunctionPlot[590].Y = -4.972

	pointsOfFunctionPlot[591].X = 5.91
	pointsOfFunctionPlot[591].Y = -4.978

	pointsOfFunctionPlot[592].X = 5.92
	pointsOfFunctionPlot[592].Y = -4.985

	pointsOfFunctionPlot[593].X = 5.93
	pointsOfFunctionPlot[593].Y = -4.991

	pointsOfFunctionPlot[594].X = 5.94
	pointsOfFunctionPlot[594].Y = -4.998

	pointsOfFunctionPlot[595].X = 5.95
	pointsOfFunctionPlot[595].Y = -5.005

	pointsOfFunctionPlot[596].X = 5.96
	pointsOfFunctionPlot[596].Y = -5.011

	pointsOfFunctionPlot[597].X = 5.97
	pointsOfFunctionPlot[597].Y = -5.018

	pointsOfFunctionPlot[598].X = 5.98
	pointsOfFunctionPlot[598].Y = -5.025

	pointsOfFunctionPlot[599].X = 5.99
	pointsOfFunctionPlot[599].Y = -5.032

	pointsOfFunctionPlot[600].X = 6.0
	pointsOfFunctionPlot[600].Y = -5.039

	pointsOfFunctionPlot[601].X = 6.01
	pointsOfFunctionPlot[601].Y = -5.047

	pointsOfFunctionPlot[602].X = 6.02
	pointsOfFunctionPlot[602].Y = -5.054

	pointsOfFunctionPlot[603].X = 6.03
	pointsOfFunctionPlot[603].Y = -5.061

	pointsOfFunctionPlot[604].X = 6.04
	pointsOfFunctionPlot[604].Y = -5.069

	pointsOfFunctionPlot[605].X = 6.05
	pointsOfFunctionPlot[605].Y = -5.077

	pointsOfFunctionPlot[606].X = 6.06
	pointsOfFunctionPlot[606].Y = -5.084

	pointsOfFunctionPlot[607].X = 6.07
	pointsOfFunctionPlot[607].Y = -5.092

	pointsOfFunctionPlot[608].X = 6.08
	pointsOfFunctionPlot[608].Y = -5.1

	pointsOfFunctionPlot[609].X = 6.09
	pointsOfFunctionPlot[609].Y = -5.108

	pointsOfFunctionPlot[610].X = 6.1
	pointsOfFunctionPlot[610].Y = -5.116

	pointsOfFunctionPlot[611].X = 6.11
	pointsOfFunctionPlot[611].Y = -5.125

	pointsOfFunctionPlot[612].X = 6.12
	pointsOfFunctionPlot[612].Y = -5.133

	pointsOfFunctionPlot[613].X = 6.13
	pointsOfFunctionPlot[613].Y = -5.141

	pointsOfFunctionPlot[614].X = 6.14
	pointsOfFunctionPlot[614].Y = -5.15

	pointsOfFunctionPlot[615].X = 6.15
	pointsOfFunctionPlot[615].Y = -5.158

	pointsOfFunctionPlot[616].X = 6.16
	pointsOfFunctionPlot[616].Y = -5.167

	pointsOfFunctionPlot[617].X = 6.17
	pointsOfFunctionPlot[617].Y = -5.176

	pointsOfFunctionPlot[618].X = 6.18
	pointsOfFunctionPlot[618].Y = -5.185

	pointsOfFunctionPlot[619].X = 6.19
	pointsOfFunctionPlot[619].Y = -5.194

	pointsOfFunctionPlot[620].X = 6.2
	pointsOfFunctionPlot[620].Y = -5.203

	pointsOfFunctionPlot[621].X = 6.21
	pointsOfFunctionPlot[621].Y = -5.212

	pointsOfFunctionPlot[622].X = 6.22
	pointsOfFunctionPlot[622].Y = -5.222

	pointsOfFunctionPlot[623].X = 6.23
	pointsOfFunctionPlot[623].Y = -5.231

	pointsOfFunctionPlot[624].X = 6.24
	pointsOfFunctionPlot[624].Y = -5.24

	pointsOfFunctionPlot[625].X = 6.25
	pointsOfFunctionPlot[625].Y = -5.25

	pointsOfFunctionPlot[626].X = 6.26
	pointsOfFunctionPlot[626].Y = -5.26

	pointsOfFunctionPlot[627].X = 6.27
	pointsOfFunctionPlot[627].Y = -5.27

	pointsOfFunctionPlot[628].X = 6.28
	pointsOfFunctionPlot[628].Y = -5.28

	pointsOfFunctionPlot[629].X = 6.29
	pointsOfFunctionPlot[629].Y = -5.29

	pointsOfFunctionPlot[630].X = 6.3
	pointsOfFunctionPlot[630].Y = -5.3

	pointsOfFunctionPlot[631].X = 6.31
	pointsOfFunctionPlot[631].Y = -5.31

	pointsOfFunctionPlot[632].X = 6.32
	pointsOfFunctionPlot[632].Y = -5.32

	pointsOfFunctionPlot[633].X = 6.33
	pointsOfFunctionPlot[633].Y = -5.331

	pointsOfFunctionPlot[634].X = 6.34
	pointsOfFunctionPlot[634].Y = -5.341

	pointsOfFunctionPlot[635].X = 6.35
	pointsOfFunctionPlot[635].Y = -5.352

	pointsOfFunctionPlot[636].X = 6.36
	pointsOfFunctionPlot[636].Y = -5.362

	pointsOfFunctionPlot[637].X = 6.37
	pointsOfFunctionPlot[637].Y = -5.373

	pointsOfFunctionPlot[638].X = 6.38
	pointsOfFunctionPlot[638].Y = -5.384

	pointsOfFunctionPlot[639].X = 6.39
	pointsOfFunctionPlot[639].Y = -5.395

	pointsOfFunctionPlot[640].X = 6.4
	pointsOfFunctionPlot[640].Y = -5.406

	pointsOfFunctionPlot[641].X = 6.41
	pointsOfFunctionPlot[641].Y = -5.418

	pointsOfFunctionPlot[642].X = 6.42
	pointsOfFunctionPlot[642].Y = -5.429

	pointsOfFunctionPlot[643].X = 6.43
	pointsOfFunctionPlot[643].Y = -5.44

	pointsOfFunctionPlot[644].X = 6.44
	pointsOfFunctionPlot[644].Y = -5.452

	pointsOfFunctionPlot[645].X = 6.45
	pointsOfFunctionPlot[645].Y = -5.463

	pointsOfFunctionPlot[646].X = 6.46
	pointsOfFunctionPlot[646].Y = -5.475

	pointsOfFunctionPlot[647].X = 6.47
	pointsOfFunctionPlot[647].Y = -5.487

	pointsOfFunctionPlot[648].X = 6.48
	pointsOfFunctionPlot[648].Y = -5.499

	pointsOfFunctionPlot[649].X = 6.49
	pointsOfFunctionPlot[649].Y = -5.511

	pointsOfFunctionPlot[650].X = 6.5
	pointsOfFunctionPlot[650].Y = -5.523

	pointsOfFunctionPlot[651].X = 6.51
	pointsOfFunctionPlot[651].Y = -5.535

	pointsOfFunctionPlot[652].X = 6.52
	pointsOfFunctionPlot[652].Y = -5.547

	pointsOfFunctionPlot[653].X = 6.53
	pointsOfFunctionPlot[653].Y = -5.56

	pointsOfFunctionPlot[654].X = 6.54
	pointsOfFunctionPlot[654].Y = -5.572

	pointsOfFunctionPlot[655].X = 6.55
	pointsOfFunctionPlot[655].Y = -5.585

	pointsOfFunctionPlot[656].X = 6.56
	pointsOfFunctionPlot[656].Y = -5.598

	pointsOfFunctionPlot[657].X = 6.57
	pointsOfFunctionPlot[657].Y = -5.61

	pointsOfFunctionPlot[658].X = 6.58
	pointsOfFunctionPlot[658].Y = -5.623

	pointsOfFunctionPlot[659].X = 6.59
	pointsOfFunctionPlot[659].Y = -5.636

	pointsOfFunctionPlot[660].X = 6.6
	pointsOfFunctionPlot[660].Y = -5.649

	pointsOfFunctionPlot[661].X = 6.61
	pointsOfFunctionPlot[661].Y = -5.662

	pointsOfFunctionPlot[662].X = 6.62
	pointsOfFunctionPlot[662].Y = -5.676

	pointsOfFunctionPlot[663].X = 6.63
	pointsOfFunctionPlot[663].Y = -5.689

	pointsOfFunctionPlot[664].X = 6.64
	pointsOfFunctionPlot[664].Y = -5.703

	pointsOfFunctionPlot[665].X = 6.65
	pointsOfFunctionPlot[665].Y = -5.716

	pointsOfFunctionPlot[666].X = 6.66
	pointsOfFunctionPlot[666].Y = -5.73

	pointsOfFunctionPlot[667].X = 6.67
	pointsOfFunctionPlot[667].Y = -5.743

	pointsOfFunctionPlot[668].X = 6.68
	pointsOfFunctionPlot[668].Y = -5.757

	pointsOfFunctionPlot[669].X = 6.69
	pointsOfFunctionPlot[669].Y = -5.771

	pointsOfFunctionPlot[670].X = 6.7
	pointsOfFunctionPlot[670].Y = -5.785

	pointsOfFunctionPlot[671].X = 6.71
	pointsOfFunctionPlot[671].Y = -5.799

	pointsOfFunctionPlot[672].X = 6.72
	pointsOfFunctionPlot[672].Y = -5.813

	pointsOfFunctionPlot[673].X = 6.73
	pointsOfFunctionPlot[673].Y = -5.828

	pointsOfFunctionPlot[674].X = 6.74
	pointsOfFunctionPlot[674].Y = -5.842

	pointsOfFunctionPlot[675].X = 6.75
	pointsOfFunctionPlot[675].Y = -5.857

	pointsOfFunctionPlot[676].X = 6.76
	pointsOfFunctionPlot[676].Y = -5.871

	pointsOfFunctionPlot[677].X = 6.77
	pointsOfFunctionPlot[677].Y = -5.886

	pointsOfFunctionPlot[678].X = 6.78
	pointsOfFunctionPlot[678].Y = -5.9

	pointsOfFunctionPlot[679].X = 6.79
	pointsOfFunctionPlot[679].Y = -5.915

	pointsOfFunctionPlot[680].X = 6.8
	pointsOfFunctionPlot[680].Y = -5.93

	pointsOfFunctionPlot[681].X = 6.81
	pointsOfFunctionPlot[681].Y = -5.945

	pointsOfFunctionPlot[682].X = 6.82
	pointsOfFunctionPlot[682].Y = -5.96

	pointsOfFunctionPlot[683].X = 6.83
	pointsOfFunctionPlot[683].Y = -5.975

	pointsOfFunctionPlot[684].X = 6.84
	pointsOfFunctionPlot[684].Y = -5.991

	pointsOfFunctionPlot[685].X = 6.85
	pointsOfFunctionPlot[685].Y = -6.006

	pointsOfFunctionPlot[686].X = 6.86
	pointsOfFunctionPlot[686].Y = -6.021

	pointsOfFunctionPlot[687].X = 6.87
	pointsOfFunctionPlot[687].Y = -6.037

	pointsOfFunctionPlot[688].X = 6.88
	pointsOfFunctionPlot[688].Y = -6.052

	pointsOfFunctionPlot[689].X = 6.89
	pointsOfFunctionPlot[689].Y = -6.068

	pointsOfFunctionPlot[690].X = 6.9
	pointsOfFunctionPlot[690].Y = -6.084

	pointsOfFunctionPlot[691].X = 6.91
	pointsOfFunctionPlot[691].Y = -6.1

	pointsOfFunctionPlot[692].X = 6.92
	pointsOfFunctionPlot[692].Y = -6.116

	pointsOfFunctionPlot[693].X = 6.93
	pointsOfFunctionPlot[693].Y = -6.132

	pointsOfFunctionPlot[694].X = 6.94
	pointsOfFunctionPlot[694].Y = -6.148

	pointsOfFunctionPlot[695].X = 6.95
	pointsOfFunctionPlot[695].Y = -6.164

	pointsOfFunctionPlot[696].X = 6.96
	pointsOfFunctionPlot[696].Y = -6.18

	pointsOfFunctionPlot[697].X = 6.97
	pointsOfFunctionPlot[697].Y = -6.196

	pointsOfFunctionPlot[698].X = 6.98
	pointsOfFunctionPlot[698].Y = -6.213

	pointsOfFunctionPlot[699].X = 6.99
	pointsOfFunctionPlot[699].Y = -6.229

	pointsOfFunctionPlot[700].X = 7.0
	pointsOfFunctionPlot[700].Y = -6.246

	pointsOfFunctionPlot[701].X = 7.01
	pointsOfFunctionPlot[701].Y = -6.262

	pointsOfFunctionPlot[702].X = 7.02
	pointsOfFunctionPlot[702].Y = -6.279

	pointsOfFunctionPlot[703].X = 7.03
	pointsOfFunctionPlot[703].Y = -6.296

	pointsOfFunctionPlot[704].X = 7.04
	pointsOfFunctionPlot[704].Y = -6.313

	pointsOfFunctionPlot[705].X = 7.05
	pointsOfFunctionPlot[705].Y = -6.329

	pointsOfFunctionPlot[706].X = 7.06
	pointsOfFunctionPlot[706].Y = -6.346

	pointsOfFunctionPlot[707].X = 7.07
	pointsOfFunctionPlot[707].Y = -6.363

	pointsOfFunctionPlot[708].X = 7.08
	pointsOfFunctionPlot[708].Y = -6.381

	pointsOfFunctionPlot[709].X = 7.09
	pointsOfFunctionPlot[709].Y = -6.398

	pointsOfFunctionPlot[710].X = 7.1
	pointsOfFunctionPlot[710].Y = -6.415

	pointsOfFunctionPlot[711].X = 7.11
	pointsOfFunctionPlot[711].Y = -6.432

	pointsOfFunctionPlot[712].X = 7.12
	pointsOfFunctionPlot[712].Y = -6.45

	pointsOfFunctionPlot[713].X = 7.13
	pointsOfFunctionPlot[713].Y = -6.467

	pointsOfFunctionPlot[714].X = 7.14
	pointsOfFunctionPlot[714].Y = -6.485

	pointsOfFunctionPlot[715].X = 7.15
	pointsOfFunctionPlot[715].Y = -6.502

	pointsOfFunctionPlot[716].X = 7.16
	pointsOfFunctionPlot[716].Y = -6.52

	pointsOfFunctionPlot[717].X = 7.17
	pointsOfFunctionPlot[717].Y = -6.538

	pointsOfFunctionPlot[718].X = 7.18
	pointsOfFunctionPlot[718].Y = -6.555

	pointsOfFunctionPlot[719].X = 7.19
	pointsOfFunctionPlot[719].Y = -6.573

	pointsOfFunctionPlot[720].X = 7.2
	pointsOfFunctionPlot[720].Y = -6.591

	pointsOfFunctionPlot[721].X = 7.21
	pointsOfFunctionPlot[721].Y = -6.609

	pointsOfFunctionPlot[722].X = 7.22
	pointsOfFunctionPlot[722].Y = -6.627

	pointsOfFunctionPlot[723].X = 7.23
	pointsOfFunctionPlot[723].Y = -6.645

	pointsOfFunctionPlot[724].X = 7.24
	pointsOfFunctionPlot[724].Y = -6.663

	pointsOfFunctionPlot[725].X = 7.25
	pointsOfFunctionPlot[725].Y = -6.682

	pointsOfFunctionPlot[726].X = 7.26
	pointsOfFunctionPlot[726].Y = -6.7

	pointsOfFunctionPlot[727].X = 7.27
	pointsOfFunctionPlot[727].Y = -6.718

	pointsOfFunctionPlot[728].X = 7.28
	pointsOfFunctionPlot[728].Y = -6.737

	pointsOfFunctionPlot[729].X = 7.29
	pointsOfFunctionPlot[729].Y = -6.755

	pointsOfFunctionPlot[730].X = 7.3
	pointsOfFunctionPlot[730].Y = -6.773

	pointsOfFunctionPlot[731].X = 7.31
	pointsOfFunctionPlot[731].Y = -6.792

	pointsOfFunctionPlot[732].X = 7.32
	pointsOfFunctionPlot[732].Y = -6.811

	pointsOfFunctionPlot[733].X = 7.33
	pointsOfFunctionPlot[733].Y = -6.829

	pointsOfFunctionPlot[734].X = 7.34
	pointsOfFunctionPlot[734].Y = -6.848

	pointsOfFunctionPlot[735].X = 7.35
	pointsOfFunctionPlot[735].Y = -6.867

	pointsOfFunctionPlot[736].X = 7.36
	pointsOfFunctionPlot[736].Y = -6.885

	pointsOfFunctionPlot[737].X = 7.37
	pointsOfFunctionPlot[737].Y = -6.904

	pointsOfFunctionPlot[738].X = 7.38
	pointsOfFunctionPlot[738].Y = -6.923

	pointsOfFunctionPlot[739].X = 7.39
	pointsOfFunctionPlot[739].Y = -6.942

	pointsOfFunctionPlot[740].X = 7.4
	pointsOfFunctionPlot[740].Y = -6.961

	pointsOfFunctionPlot[741].X = 7.41
	pointsOfFunctionPlot[741].Y = -6.98

	pointsOfFunctionPlot[742].X = 7.42
	pointsOfFunctionPlot[742].Y = -6.999

	pointsOfFunctionPlot[743].X = 7.43
	pointsOfFunctionPlot[743].Y = -7.018

	pointsOfFunctionPlot[744].X = 7.44
	pointsOfFunctionPlot[744].Y = -7.037

	pointsOfFunctionPlot[745].X = 7.45
	pointsOfFunctionPlot[745].Y = -7.056

	pointsOfFunctionPlot[746].X = 7.46
	pointsOfFunctionPlot[746].Y = -7.076

	pointsOfFunctionPlot[747].X = 7.47
	pointsOfFunctionPlot[747].Y = -7.095

	pointsOfFunctionPlot[748].X = 7.48
	pointsOfFunctionPlot[748].Y = -7.114

	pointsOfFunctionPlot[749].X = 7.49
	pointsOfFunctionPlot[749].Y = -7.134

	pointsOfFunctionPlot[750].X = 7.5
	pointsOfFunctionPlot[750].Y = -7.153

	pointsOfFunctionPlot[751].X = 7.51
	pointsOfFunctionPlot[751].Y = -7.172

	pointsOfFunctionPlot[752].X = 7.52
	pointsOfFunctionPlot[752].Y = -7.192

	pointsOfFunctionPlot[753].X = 7.53
	pointsOfFunctionPlot[753].Y = -7.211

	pointsOfFunctionPlot[754].X = 7.54
	pointsOfFunctionPlot[754].Y = -7.231

	pointsOfFunctionPlot[755].X = 7.55
	pointsOfFunctionPlot[755].Y = -7.25

	pointsOfFunctionPlot[756].X = 7.56
	pointsOfFunctionPlot[756].Y = -7.27

	pointsOfFunctionPlot[757].X = 7.57
	pointsOfFunctionPlot[757].Y = -7.289

	pointsOfFunctionPlot[758].X = 7.58
	pointsOfFunctionPlot[758].Y = -7.309

	pointsOfFunctionPlot[759].X = 7.59
	pointsOfFunctionPlot[759].Y = -7.329

	pointsOfFunctionPlot[760].X = 7.6
	pointsOfFunctionPlot[760].Y = -7.348

	pointsOfFunctionPlot[761].X = 7.61
	pointsOfFunctionPlot[761].Y = -7.368

	pointsOfFunctionPlot[762].X = 7.62
	pointsOfFunctionPlot[762].Y = -7.388

	pointsOfFunctionPlot[763].X = 7.63
	pointsOfFunctionPlot[763].Y = -7.407

	pointsOfFunctionPlot[764].X = 7.64
	pointsOfFunctionPlot[764].Y = -7.427

	pointsOfFunctionPlot[765].X = 7.65
	pointsOfFunctionPlot[765].Y = -7.447

	pointsOfFunctionPlot[766].X = 7.66
	pointsOfFunctionPlot[766].Y = -7.467

	pointsOfFunctionPlot[767].X = 7.67
	pointsOfFunctionPlot[767].Y = -7.487

	pointsOfFunctionPlot[768].X = 7.68
	pointsOfFunctionPlot[768].Y = -7.506

	pointsOfFunctionPlot[769].X = 7.69
	pointsOfFunctionPlot[769].Y = -7.526

	pointsOfFunctionPlot[770].X = 7.7
	pointsOfFunctionPlot[770].Y = -7.546

	pointsOfFunctionPlot[771].X = 7.71
	pointsOfFunctionPlot[771].Y = -7.566

	pointsOfFunctionPlot[772].X = 7.72
	pointsOfFunctionPlot[772].Y = -7.586

	pointsOfFunctionPlot[773].X = 7.73
	pointsOfFunctionPlot[773].Y = -7.606

	pointsOfFunctionPlot[774].X = 7.74
	pointsOfFunctionPlot[774].Y = -7.626

	pointsOfFunctionPlot[775].X = 7.75
	pointsOfFunctionPlot[775].Y = -7.646

	pointsOfFunctionPlot[776].X = 7.76
	pointsOfFunctionPlot[776].Y = -7.666

	pointsOfFunctionPlot[777].X = 7.77
	pointsOfFunctionPlot[777].Y = -7.686

	pointsOfFunctionPlot[778].X = 7.78
	pointsOfFunctionPlot[778].Y = -7.706

	pointsOfFunctionPlot[779].X = 7.79
	pointsOfFunctionPlot[779].Y = -7.726

	pointsOfFunctionPlot[780].X = 7.8
	pointsOfFunctionPlot[780].Y = -7.746

	pointsOfFunctionPlot[781].X = 7.81
	pointsOfFunctionPlot[781].Y = -7.766

	pointsOfFunctionPlot[782].X = 7.82
	pointsOfFunctionPlot[782].Y = -7.786

	pointsOfFunctionPlot[783].X = 7.83
	pointsOfFunctionPlot[783].Y = -7.806

	pointsOfFunctionPlot[784].X = 7.84
	pointsOfFunctionPlot[784].Y = -7.826

	pointsOfFunctionPlot[785].X = 7.85
	pointsOfFunctionPlot[785].Y = -7.846

	pointsOfFunctionPlot[786].X = 7.86
	pointsOfFunctionPlot[786].Y = -7.866

	pointsOfFunctionPlot[787].X = 7.87
	pointsOfFunctionPlot[787].Y = -7.886

	pointsOfFunctionPlot[788].X = 7.88
	pointsOfFunctionPlot[788].Y = -7.906

	pointsOfFunctionPlot[789].X = 7.89
	pointsOfFunctionPlot[789].Y = -7.926

	pointsOfFunctionPlot[790].X = 7.9
	pointsOfFunctionPlot[790].Y = -7.946

	pointsOfFunctionPlot[791].X = 7.91
	pointsOfFunctionPlot[791].Y = -7.966

	pointsOfFunctionPlot[792].X = 7.92
	pointsOfFunctionPlot[792].Y = -7.986

	pointsOfFunctionPlot[793].X = 7.93
	pointsOfFunctionPlot[793].Y = -8.005

	pointsOfFunctionPlot[794].X = 7.94
	pointsOfFunctionPlot[794].Y = -8.025

	pointsOfFunctionPlot[795].X = 7.95
	pointsOfFunctionPlot[795].Y = -8.045

	pointsOfFunctionPlot[796].X = 7.96
	pointsOfFunctionPlot[796].Y = -8.065

	pointsOfFunctionPlot[797].X = 7.97
	pointsOfFunctionPlot[797].Y = -8.085

	pointsOfFunctionPlot[798].X = 7.98
	pointsOfFunctionPlot[798].Y = -8.105

	pointsOfFunctionPlot[799].X = 7.99
	pointsOfFunctionPlot[799].Y = -8.125

	pointsOfFunctionPlot[800].X = 8.0
	pointsOfFunctionPlot[800].Y = -8.145

	pointsOfFunctionPlot[801].X = 8.01
	pointsOfFunctionPlot[801].Y = -8.165

	pointsOfFunctionPlot[802].X = 8.02
	pointsOfFunctionPlot[802].Y = -8.185

	pointsOfFunctionPlot[803].X = 8.03
	pointsOfFunctionPlot[803].Y = -8.205

	pointsOfFunctionPlot[804].X = 8.04
	pointsOfFunctionPlot[804].Y = -8.224

	pointsOfFunctionPlot[805].X = 8.05
	pointsOfFunctionPlot[805].Y = -8.244

	pointsOfFunctionPlot[806].X = 8.06
	pointsOfFunctionPlot[806].Y = -8.264

	pointsOfFunctionPlot[807].X = 8.07
	pointsOfFunctionPlot[807].Y = -8.284

	pointsOfFunctionPlot[808].X = 8.08
	pointsOfFunctionPlot[808].Y = -8.304

	pointsOfFunctionPlot[809].X = 8.09
	pointsOfFunctionPlot[809].Y = -8.323

	pointsOfFunctionPlot[810].X = 8.1
	pointsOfFunctionPlot[810].Y = -8.343

	pointsOfFunctionPlot[811].X = 8.11
	pointsOfFunctionPlot[811].Y = -8.363

	pointsOfFunctionPlot[812].X = 8.12
	pointsOfFunctionPlot[812].Y = -8.382

	pointsOfFunctionPlot[813].X = 8.13
	pointsOfFunctionPlot[813].Y = -8.402

	pointsOfFunctionPlot[814].X = 8.14
	pointsOfFunctionPlot[814].Y = -8.422

	pointsOfFunctionPlot[815].X = 8.15
	pointsOfFunctionPlot[815].Y = -8.441

	pointsOfFunctionPlot[816].X = 8.16
	pointsOfFunctionPlot[816].Y = -8.461

	pointsOfFunctionPlot[817].X = 8.17
	pointsOfFunctionPlot[817].Y = -8.48

	pointsOfFunctionPlot[818].X = 8.18
	pointsOfFunctionPlot[818].Y = -8.5

	pointsOfFunctionPlot[819].X = 8.19
	pointsOfFunctionPlot[819].Y = -8.519

	pointsOfFunctionPlot[820].X = 8.2
	pointsOfFunctionPlot[820].Y = -8.539

	pointsOfFunctionPlot[821].X = 8.21
	pointsOfFunctionPlot[821].Y = -8.558

	pointsOfFunctionPlot[822].X = 8.22
	pointsOfFunctionPlot[822].Y = -8.577

	pointsOfFunctionPlot[823].X = 8.23
	pointsOfFunctionPlot[823].Y = -8.597

	pointsOfFunctionPlot[824].X = 8.24
	pointsOfFunctionPlot[824].Y = -8.616

	pointsOfFunctionPlot[825].X = 8.25
	pointsOfFunctionPlot[825].Y = -8.635

	pointsOfFunctionPlot[826].X = 8.26
	pointsOfFunctionPlot[826].Y = -8.655

	pointsOfFunctionPlot[827].X = 8.27
	pointsOfFunctionPlot[827].Y = -8.674

	pointsOfFunctionPlot[828].X = 8.28
	pointsOfFunctionPlot[828].Y = -8.693

	pointsOfFunctionPlot[829].X = 8.29
	pointsOfFunctionPlot[829].Y = -8.712

	pointsOfFunctionPlot[830].X = 8.3
	pointsOfFunctionPlot[830].Y = -8.731

	pointsOfFunctionPlot[831].X = 8.31
	pointsOfFunctionPlot[831].Y = -8.75

	pointsOfFunctionPlot[832].X = 8.32
	pointsOfFunctionPlot[832].Y = -8.769

	pointsOfFunctionPlot[833].X = 8.33
	pointsOfFunctionPlot[833].Y = -8.788

	pointsOfFunctionPlot[834].X = 8.34
	pointsOfFunctionPlot[834].Y = -8.807

	pointsOfFunctionPlot[835].X = 8.35
	pointsOfFunctionPlot[835].Y = -8.825

	pointsOfFunctionPlot[836].X = 8.36
	pointsOfFunctionPlot[836].Y = -8.844

	pointsOfFunctionPlot[837].X = 8.37
	pointsOfFunctionPlot[837].Y = -8.863

	pointsOfFunctionPlot[838].X = 8.38
	pointsOfFunctionPlot[838].Y = -8.882

	pointsOfFunctionPlot[839].X = 8.39
	pointsOfFunctionPlot[839].Y = -8.9

	pointsOfFunctionPlot[840].X = 8.4
	pointsOfFunctionPlot[840].Y = -8.919

	pointsOfFunctionPlot[841].X = 8.41
	pointsOfFunctionPlot[841].Y = -8.937

	pointsOfFunctionPlot[842].X = 8.42
	pointsOfFunctionPlot[842].Y = -8.956

	pointsOfFunctionPlot[843].X = 8.43
	pointsOfFunctionPlot[843].Y = -8.974

	pointsOfFunctionPlot[844].X = 8.44
	pointsOfFunctionPlot[844].Y = -8.993

	pointsOfFunctionPlot[845].X = 8.45
	pointsOfFunctionPlot[845].Y = -9.011

	pointsOfFunctionPlot[846].X = 8.46
	pointsOfFunctionPlot[846].Y = -9.029

	pointsOfFunctionPlot[847].X = 8.47
	pointsOfFunctionPlot[847].Y = -9.047

	pointsOfFunctionPlot[848].X = 8.48
	pointsOfFunctionPlot[848].Y = -9.065

	pointsOfFunctionPlot[849].X = 8.49
	pointsOfFunctionPlot[849].Y = -9.084

	pointsOfFunctionPlot[850].X = 8.5
	pointsOfFunctionPlot[850].Y = -9.102

	pointsOfFunctionPlot[851].X = 8.51
	pointsOfFunctionPlot[851].Y = -9.12

	pointsOfFunctionPlot[852].X = 8.52
	pointsOfFunctionPlot[852].Y = -9.137

	pointsOfFunctionPlot[853].X = 8.53
	pointsOfFunctionPlot[853].Y = -9.155

	pointsOfFunctionPlot[854].X = 8.54
	pointsOfFunctionPlot[854].Y = -9.173

	pointsOfFunctionPlot[855].X = 8.55
	pointsOfFunctionPlot[855].Y = -9.191

	pointsOfFunctionPlot[856].X = 8.56
	pointsOfFunctionPlot[856].Y = -9.208

	pointsOfFunctionPlot[857].X = 8.57
	pointsOfFunctionPlot[857].Y = -9.226

	pointsOfFunctionPlot[858].X = 8.58
	pointsOfFunctionPlot[858].Y = -9.243

	pointsOfFunctionPlot[859].X = 8.59
	pointsOfFunctionPlot[859].Y = -9.261

	pointsOfFunctionPlot[860].X = 8.6
	pointsOfFunctionPlot[860].Y = -9.278

	pointsOfFunctionPlot[861].X = 8.61
	pointsOfFunctionPlot[861].Y = -9.296

	pointsOfFunctionPlot[862].X = 8.62
	pointsOfFunctionPlot[862].Y = -9.313

	pointsOfFunctionPlot[863].X = 8.63
	pointsOfFunctionPlot[863].Y = -9.33

	pointsOfFunctionPlot[864].X = 8.64
	pointsOfFunctionPlot[864].Y = -9.347

	pointsOfFunctionPlot[865].X = 8.65
	pointsOfFunctionPlot[865].Y = -9.364

	pointsOfFunctionPlot[866].X = 8.66
	pointsOfFunctionPlot[866].Y = -9.381

	pointsOfFunctionPlot[867].X = 8.67
	pointsOfFunctionPlot[867].Y = -9.398

	pointsOfFunctionPlot[868].X = 8.68
	pointsOfFunctionPlot[868].Y = -9.415

	pointsOfFunctionPlot[869].X = 8.69
	pointsOfFunctionPlot[869].Y = -9.432

	pointsOfFunctionPlot[870].X = 8.7
	pointsOfFunctionPlot[870].Y = -9.448

	pointsOfFunctionPlot[871].X = 8.71
	pointsOfFunctionPlot[871].Y = -9.465

	pointsOfFunctionPlot[872].X = 8.72
	pointsOfFunctionPlot[872].Y = -9.481

	pointsOfFunctionPlot[873].X = 8.73
	pointsOfFunctionPlot[873].Y = -9.498

	pointsOfFunctionPlot[874].X = 8.74
	pointsOfFunctionPlot[874].Y = -9.514

	pointsOfFunctionPlot[875].X = 8.75
	pointsOfFunctionPlot[875].Y = -9.53

	pointsOfFunctionPlot[876].X = 8.76
	pointsOfFunctionPlot[876].Y = -9.547

	pointsOfFunctionPlot[877].X = 8.77
	pointsOfFunctionPlot[877].Y = -9.563

	pointsOfFunctionPlot[878].X = 8.78
	pointsOfFunctionPlot[878].Y = -9.579

	pointsOfFunctionPlot[879].X = 8.79
	pointsOfFunctionPlot[879].Y = -9.595

	pointsOfFunctionPlot[880].X = 8.8
	pointsOfFunctionPlot[880].Y = -9.611

	pointsOfFunctionPlot[881].X = 8.81
	pointsOfFunctionPlot[881].Y = -9.626

	pointsOfFunctionPlot[882].X = 8.82
	pointsOfFunctionPlot[882].Y = -9.642

	pointsOfFunctionPlot[883].X = 8.83
	pointsOfFunctionPlot[883].Y = -9.658

	pointsOfFunctionPlot[884].X = 8.84
	pointsOfFunctionPlot[884].Y = -9.673

	pointsOfFunctionPlot[885].X = 8.85
	pointsOfFunctionPlot[885].Y = -9.689

	pointsOfFunctionPlot[886].X = 8.86
	pointsOfFunctionPlot[886].Y = -9.704

	pointsOfFunctionPlot[887].X = 8.87
	pointsOfFunctionPlot[887].Y = -9.72

	pointsOfFunctionPlot[888].X = 8.88
	pointsOfFunctionPlot[888].Y = -9.735

	pointsOfFunctionPlot[889].X = 8.89
	pointsOfFunctionPlot[889].Y = -9.75

	pointsOfFunctionPlot[890].X = 8.9
	pointsOfFunctionPlot[890].Y = -9.765

	pointsOfFunctionPlot[891].X = 8.91
	pointsOfFunctionPlot[891].Y = -9.78

	pointsOfFunctionPlot[892].X = 8.92
	pointsOfFunctionPlot[892].Y = -9.795

	pointsOfFunctionPlot[893].X = 8.93
	pointsOfFunctionPlot[893].Y = -9.81

	pointsOfFunctionPlot[894].X = 8.94
	pointsOfFunctionPlot[894].Y = -9.824

	pointsOfFunctionPlot[895].X = 8.95
	pointsOfFunctionPlot[895].Y = -9.839

	pointsOfFunctionPlot[896].X = 8.96
	pointsOfFunctionPlot[896].Y = -9.853

	pointsOfFunctionPlot[897].X = 8.97
	pointsOfFunctionPlot[897].Y = -9.868

	pointsOfFunctionPlot[898].X = 8.98
	pointsOfFunctionPlot[898].Y = -9.882

	pointsOfFunctionPlot[899].X = 8.99
	pointsOfFunctionPlot[899].Y = -9.897

	pointsOfFunctionPlot[900].X = 9.0
	pointsOfFunctionPlot[900].Y = -9.911

	pointsOfFunctionPlot[901].X = 9.01
	pointsOfFunctionPlot[901].Y = -9.925

	pointsOfFunctionPlot[902].X = 9.02
	pointsOfFunctionPlot[902].Y = -9.939

	pointsOfFunctionPlot[903].X = 9.03
	pointsOfFunctionPlot[903].Y = -9.953

	pointsOfFunctionPlot[904].X = 9.04
	pointsOfFunctionPlot[904].Y = -9.966

	pointsOfFunctionPlot[905].X = 9.05
	pointsOfFunctionPlot[905].Y = -9.98

	pointsOfFunctionPlot[906].X = 9.06
	pointsOfFunctionPlot[906].Y = -9.994

	pointsOfFunctionPlot[907].X = 9.07
	pointsOfFunctionPlot[907].Y = -10.007

	pointsOfFunctionPlot[908].X = 9.08
	pointsOfFunctionPlot[908].Y = -10.021

	pointsOfFunctionPlot[909].X = 9.09
	pointsOfFunctionPlot[909].Y = -10.034

	pointsOfFunctionPlot[910].X = 9.1
	pointsOfFunctionPlot[910].Y = -10.047

	pointsOfFunctionPlot[911].X = 9.11
	pointsOfFunctionPlot[911].Y = -10.06

	pointsOfFunctionPlot[912].X = 9.12
	pointsOfFunctionPlot[912].Y = -10.073

	pointsOfFunctionPlot[913].X = 9.13
	pointsOfFunctionPlot[913].Y = -10.086

	pointsOfFunctionPlot[914].X = 9.14
	pointsOfFunctionPlot[914].Y = -10.099

	pointsOfFunctionPlot[915].X = 9.15
	pointsOfFunctionPlot[915].Y = -10.112

	pointsOfFunctionPlot[916].X = 9.16
	pointsOfFunctionPlot[916].Y = -10.125

	pointsOfFunctionPlot[917].X = 9.17
	pointsOfFunctionPlot[917].Y = -10.137

	pointsOfFunctionPlot[918].X = 9.18
	pointsOfFunctionPlot[918].Y = -10.15

	pointsOfFunctionPlot[919].X = 9.19
	pointsOfFunctionPlot[919].Y = -10.162

	pointsOfFunctionPlot[920].X = 9.2
	pointsOfFunctionPlot[920].Y = -10.174

	pointsOfFunctionPlot[921].X = 9.21
	pointsOfFunctionPlot[921].Y = -10.187

	pointsOfFunctionPlot[922].X = 9.22
	pointsOfFunctionPlot[922].Y = -10.199

	pointsOfFunctionPlot[923].X = 9.23
	pointsOfFunctionPlot[923].Y = -10.211

	pointsOfFunctionPlot[924].X = 9.24
	pointsOfFunctionPlot[924].Y = -10.223

	pointsOfFunctionPlot[925].X = 9.25
	pointsOfFunctionPlot[925].Y = -10.234

	pointsOfFunctionPlot[926].X = 9.26
	pointsOfFunctionPlot[926].Y = -10.246

	pointsOfFunctionPlot[927].X = 9.27
	pointsOfFunctionPlot[927].Y = -10.258

	pointsOfFunctionPlot[928].X = 9.28
	pointsOfFunctionPlot[928].Y = -10.269

	pointsOfFunctionPlot[929].X = 9.29
	pointsOfFunctionPlot[929].Y = -10.28

	pointsOfFunctionPlot[930].X = 9.3
	pointsOfFunctionPlot[930].Y = -10.292

	pointsOfFunctionPlot[931].X = 9.31
	pointsOfFunctionPlot[931].Y = -10.303

	pointsOfFunctionPlot[932].X = 9.32
	pointsOfFunctionPlot[932].Y = -10.314

	pointsOfFunctionPlot[933].X = 9.33
	pointsOfFunctionPlot[933].Y = -10.325

	pointsOfFunctionPlot[934].X = 9.34
	pointsOfFunctionPlot[934].Y = -10.336

	pointsOfFunctionPlot[935].X = 9.35
	pointsOfFunctionPlot[935].Y = -10.347

	pointsOfFunctionPlot[936].X = 9.36
	pointsOfFunctionPlot[936].Y = -10.357

	pointsOfFunctionPlot[937].X = 9.37
	pointsOfFunctionPlot[937].Y = -10.368

	pointsOfFunctionPlot[938].X = 9.38
	pointsOfFunctionPlot[938].Y = -10.379

	pointsOfFunctionPlot[939].X = 9.39
	pointsOfFunctionPlot[939].Y = -10.389

	pointsOfFunctionPlot[940].X = 9.4
	pointsOfFunctionPlot[940].Y = -10.399

	pointsOfFunctionPlot[941].X = 9.41
	pointsOfFunctionPlot[941].Y = -10.409

	pointsOfFunctionPlot[942].X = 9.42
	pointsOfFunctionPlot[942].Y = -10.42

	pointsOfFunctionPlot[943].X = 9.43
	pointsOfFunctionPlot[943].Y = -10.43

	pointsOfFunctionPlot[944].X = 9.44
	pointsOfFunctionPlot[944].Y = -10.439

	pointsOfFunctionPlot[945].X = 9.45
	pointsOfFunctionPlot[945].Y = -10.449

	pointsOfFunctionPlot[946].X = 9.46
	pointsOfFunctionPlot[946].Y = -10.459

	pointsOfFunctionPlot[947].X = 9.47
	pointsOfFunctionPlot[947].Y = -10.469

	pointsOfFunctionPlot[948].X = 9.48
	pointsOfFunctionPlot[948].Y = -10.478

	pointsOfFunctionPlot[949].X = 9.49
	pointsOfFunctionPlot[949].Y = -10.487

	pointsOfFunctionPlot[950].X = 9.5
	pointsOfFunctionPlot[950].Y = -10.497

	pointsOfFunctionPlot[951].X = 9.51
	pointsOfFunctionPlot[951].Y = -10.506

	pointsOfFunctionPlot[952].X = 9.52
	pointsOfFunctionPlot[952].Y = -10.515

	pointsOfFunctionPlot[953].X = 9.53
	pointsOfFunctionPlot[953].Y = -10.524

	pointsOfFunctionPlot[954].X = 9.54
	pointsOfFunctionPlot[954].Y = -10.533

	pointsOfFunctionPlot[955].X = 9.55
	pointsOfFunctionPlot[955].Y = -10.542

	pointsOfFunctionPlot[956].X = 9.56
	pointsOfFunctionPlot[956].Y = -10.55

	pointsOfFunctionPlot[957].X = 9.57
	pointsOfFunctionPlot[957].Y = -10.559

	pointsOfFunctionPlot[958].X = 9.58
	pointsOfFunctionPlot[958].Y = -10.568

	pointsOfFunctionPlot[959].X = 9.59
	pointsOfFunctionPlot[959].Y = -10.576

	pointsOfFunctionPlot[960].X = 9.6
	pointsOfFunctionPlot[960].Y = -10.584

	pointsOfFunctionPlot[961].X = 9.61
	pointsOfFunctionPlot[961].Y = -10.592

	pointsOfFunctionPlot[962].X = 9.62
	pointsOfFunctionPlot[962].Y = -10.601

	pointsOfFunctionPlot[963].X = 9.63
	pointsOfFunctionPlot[963].Y = -10.609

	pointsOfFunctionPlot[964].X = 9.64
	pointsOfFunctionPlot[964].Y = -10.616

	pointsOfFunctionPlot[965].X = 9.65
	pointsOfFunctionPlot[965].Y = -10.624

	pointsOfFunctionPlot[966].X = 9.66
	pointsOfFunctionPlot[966].Y = -10.632

	pointsOfFunctionPlot[967].X = 9.67
	pointsOfFunctionPlot[967].Y = -10.64

	pointsOfFunctionPlot[968].X = 9.68
	pointsOfFunctionPlot[968].Y = -10.647

	pointsOfFunctionPlot[969].X = 9.69
	pointsOfFunctionPlot[969].Y = -10.655

	pointsOfFunctionPlot[970].X = 9.7
	pointsOfFunctionPlot[970].Y = -10.662

	pointsOfFunctionPlot[971].X = 9.71
	pointsOfFunctionPlot[971].Y = -10.669

	pointsOfFunctionPlot[972].X = 9.72
	pointsOfFunctionPlot[972].Y = -10.676

	pointsOfFunctionPlot[973].X = 9.73
	pointsOfFunctionPlot[973].Y = -10.683

	pointsOfFunctionPlot[974].X = 9.74
	pointsOfFunctionPlot[974].Y = -10.69

	pointsOfFunctionPlot[975].X = 9.75
	pointsOfFunctionPlot[975].Y = -10.697

	pointsOfFunctionPlot[976].X = 9.76
	pointsOfFunctionPlot[976].Y = -10.704

	pointsOfFunctionPlot[977].X = 9.77
	pointsOfFunctionPlot[977].Y = -10.711

	pointsOfFunctionPlot[978].X = 9.78
	pointsOfFunctionPlot[978].Y = -10.717

	pointsOfFunctionPlot[979].X = 9.79
	pointsOfFunctionPlot[979].Y = -10.724

	pointsOfFunctionPlot[980].X = 9.8
	pointsOfFunctionPlot[980].Y = -10.73

	pointsOfFunctionPlot[981].X = 9.81
	pointsOfFunctionPlot[981].Y = -10.73

	pointsOfFunctionPlot[982].X = 9.82
	pointsOfFunctionPlot[982].Y = -10.736

	pointsOfFunctionPlot[983].X = 9.83
	pointsOfFunctionPlot[983].Y = -10.742

	pointsOfFunctionPlot[984].X = 9.84
	pointsOfFunctionPlot[984].Y = -10.755

	pointsOfFunctionPlot[985].X = 9.85
	pointsOfFunctionPlot[985].Y = -10.76

	pointsOfFunctionPlot[986].X = 9.86
	pointsOfFunctionPlot[986].Y = -10.766

	pointsOfFunctionPlot[987].X = 9.87
	pointsOfFunctionPlot[987].Y = -10.772

	pointsOfFunctionPlot[988].X = 9.88
	pointsOfFunctionPlot[988].Y = -10.778

	pointsOfFunctionPlot[989].X = 9.89
	pointsOfFunctionPlot[989].Y = -10.783

	pointsOfFunctionPlot[990].X = 9.9
	pointsOfFunctionPlot[990].Y = -10.789

	pointsOfFunctionPlot[991].X = 9.91
	pointsOfFunctionPlot[991].Y = -10.794

	pointsOfFunctionPlot[992].X = 9.92
	pointsOfFunctionPlot[992].Y = -10.799

	pointsOfFunctionPlot[993].X = 9.93
	pointsOfFunctionPlot[993].Y = -10.805

	pointsOfFunctionPlot[994].X = 9.94
	pointsOfFunctionPlot[994].Y = -10.81

	pointsOfFunctionPlot[995].X = 9.95
	pointsOfFunctionPlot[995].Y = -10.815

	pointsOfFunctionPlot[996].X = 9.96
	pointsOfFunctionPlot[996].Y = -10.82

	pointsOfFunctionPlot[997].X = 9.97
	pointsOfFunctionPlot[997].Y = -10.825

	pointsOfFunctionPlot[998].X = 9.98
	pointsOfFunctionPlot[998].Y = -10.829

	pointsOfFunctionPlot[999].X = 9.99
	pointsOfFunctionPlot[999].Y = -10.834

	pointsOfFunctionPlot[1_000].X = 10.0
	pointsOfFunctionPlot[1_000].Y = -10.839

	pointsOfFunctionPlot[1_001].X = 10.01
	pointsOfFunctionPlot[1_001].Y = -10.843

	pointsOfFunctionPlot[1_002].X = 10.02
	pointsOfFunctionPlot[1_002].Y = -10.848

	pointsOfFunctionPlot[1_003].X = 10.03
	pointsOfFunctionPlot[1_003].Y = -10.852

	pointsOfFunctionPlot[1_004].X = 10.04
	pointsOfFunctionPlot[1_004].Y = -10.856

	pointsOfFunctionPlot[1_005].X = 10.05
	pointsOfFunctionPlot[1_005].Y = -10.86

	pointsOfFunctionPlot[1_006].X = 10.06
	pointsOfFunctionPlot[1_006].Y = -10.864

	pointsOfFunctionPlot[1_007].X = 10.07
	pointsOfFunctionPlot[1_007].Y = -10.869

	pointsOfFunctionPlot[1_008].X = 10.08
	pointsOfFunctionPlot[1_008].Y = -10.872

	pointsOfFunctionPlot[1_009].X = 10.09
	pointsOfFunctionPlot[1_009].Y = -10.876

	pointsOfFunctionPlot[1_010].X = 10.1
	pointsOfFunctionPlot[1_010].Y = -10.88

	pointsOfFunctionPlot[1_011].X = 10.11
	pointsOfFunctionPlot[1_011].Y = -10.884

	pointsOfFunctionPlot[1_012].X = 10.12
	pointsOfFunctionPlot[1_012].Y = -10.887

	pointsOfFunctionPlot[1_013].X = 10.13
	pointsOfFunctionPlot[1_013].Y = -10.891

	pointsOfFunctionPlot[1_014].X = 10.14
	pointsOfFunctionPlot[1_014].Y = -10.894

	pointsOfFunctionPlot[1_015].X = 10.15
	pointsOfFunctionPlot[1_015].Y = -10.898

	pointsOfFunctionPlot[1_016].X = 10.16
	pointsOfFunctionPlot[1_016].Y = -10.901

	pointsOfFunctionPlot[1_017].X = 10.17
	pointsOfFunctionPlot[1_017].Y = -10.904

	pointsOfFunctionPlot[1_018].X = 10.18
	pointsOfFunctionPlot[1_018].Y = -10.908

	pointsOfFunctionPlot[1_019].X = 10.19
	pointsOfFunctionPlot[1_019].Y = -10.911

	pointsOfFunctionPlot[1_020].X = 10.2
	pointsOfFunctionPlot[1_020].Y = -10.914

	pointsOfFunctionPlot[1_021].X = 10.21
	pointsOfFunctionPlot[1_021].Y = -10.917

	pointsOfFunctionPlot[1_022].X = 10.22
	pointsOfFunctionPlot[1_022].Y = -10.92

	pointsOfFunctionPlot[1_023].X = 10.23
	pointsOfFunctionPlot[1_023].Y = -10.923

	pointsOfFunctionPlot[1_024].X = 10.24
	pointsOfFunctionPlot[1_024].Y = -10.925

	pointsOfFunctionPlot[1_025].X = 10.25
	pointsOfFunctionPlot[1_025].Y = -10.928

	pointsOfFunctionPlot[1_026].X = 10.26
	pointsOfFunctionPlot[1_026].Y = -10.931

	pointsOfFunctionPlot[1_027].X = 10.27
	pointsOfFunctionPlot[1_027].Y = -10.933

	pointsOfFunctionPlot[1_028].X = 10.28
	pointsOfFunctionPlot[1_028].Y = -10.936

	pointsOfFunctionPlot[1_029].X = 10.29
	pointsOfFunctionPlot[1_029].Y = -10.938

	pointsOfFunctionPlot[1_030].X = 10.3
	pointsOfFunctionPlot[1_030].Y = -10.94

	pointsOfFunctionPlot[1_031].X = 10.31
	pointsOfFunctionPlot[1_031].Y = -10.943

	pointsOfFunctionPlot[1_032].X = 10.32
	pointsOfFunctionPlot[1_032].Y = -10.945

	pointsOfFunctionPlot[1_033].X = 10.33
	pointsOfFunctionPlot[1_033].Y = -10.947

	pointsOfFunctionPlot[1_034].X = 10.34
	pointsOfFunctionPlot[1_034].Y = -10.949

	pointsOfFunctionPlot[1_035].X = 10.35
	pointsOfFunctionPlot[1_035].Y = -10.951

	pointsOfFunctionPlot[1_036].X = 10.36
	pointsOfFunctionPlot[1_036].Y = -10.953

	pointsOfFunctionPlot[1_037].X = 10.37
	pointsOfFunctionPlot[1_037].Y = -10.955

	pointsOfFunctionPlot[1_038].X = 10.38
	pointsOfFunctionPlot[1_038].Y = -10.957

	pointsOfFunctionPlot[1_039].X = 10.39
	pointsOfFunctionPlot[1_039].Y = -10.959

	pointsOfFunctionPlot[1_040].X = 10.40
	pointsOfFunctionPlot[1_040].Y = -10.961

	pointsOfFunctionPlot[1_041].X = 10.41
	pointsOfFunctionPlot[1_041].Y = -10.962

	pointsOfFunctionPlot[1_042].X = 10.42
	pointsOfFunctionPlot[1_042].Y = -10.964

	pointsOfFunctionPlot[1_043].X = 10.43
	pointsOfFunctionPlot[1_043].Y = -10.965

	pointsOfFunctionPlot[1_044].X = 10.44
	pointsOfFunctionPlot[1_044].Y = -10.967

	pointsOfFunctionPlot[1_045].X = 10.45
	pointsOfFunctionPlot[1_045].Y = -10.968

	pointsOfFunctionPlot[1_046].X = 10.46
	pointsOfFunctionPlot[1_046].Y = -10.97

	pointsOfFunctionPlot[1_047].X = 10.47
	pointsOfFunctionPlot[1_047].Y = -10.971

	pointsOfFunctionPlot[1_048].X = 10.48
	pointsOfFunctionPlot[1_048].Y = -10.973

	pointsOfFunctionPlot[1_049].X = 10.49
	pointsOfFunctionPlot[1_049].Y = -10.974

	pointsOfFunctionPlot[1_050].X = 10.5
	pointsOfFunctionPlot[1_050].Y = -10.975

	pointsOfFunctionPlot[1_051].X = 10.51
	pointsOfFunctionPlot[1_051].Y = -10.976

	pointsOfFunctionPlot[1_052].X = 10.52
	pointsOfFunctionPlot[1_052].Y = -10.977

	pointsOfFunctionPlot[1_053].X = 10.53
	pointsOfFunctionPlot[1_053].Y = -10.978

	pointsOfFunctionPlot[1_054].X = 10.54
	pointsOfFunctionPlot[1_054].Y = -10.98

	pointsOfFunctionPlot[1_055].X = 10.55
	pointsOfFunctionPlot[1_055].Y = -10.981

	pointsOfFunctionPlot[1_056].X = 10.56
	pointsOfFunctionPlot[1_056].Y = -10.981

	pointsOfFunctionPlot[1_057].X = 10.57
	pointsOfFunctionPlot[1_057].Y = -10.982

	pointsOfFunctionPlot[1_058].X = 10.58
	pointsOfFunctionPlot[1_058].Y = -10.983

	pointsOfFunctionPlot[1_059].X = 10.59
	pointsOfFunctionPlot[1_059].Y = -10.984

	pointsOfFunctionPlot[1_060].X = 10.6
	pointsOfFunctionPlot[1_060].Y = -10.985

	pointsOfFunctionPlot[1_061].X = 10.61
	pointsOfFunctionPlot[1_061].Y = -10.986

	pointsOfFunctionPlot[1_062].X = 10.62
	pointsOfFunctionPlot[1_062].Y = -10.986

	pointsOfFunctionPlot[1_063].X = 10.63
	pointsOfFunctionPlot[1_063].Y = -10.987

	pointsOfFunctionPlot[1_064].X = 10.64
	pointsOfFunctionPlot[1_064].Y = -10.988

	pointsOfFunctionPlot[1_065].X = 10.65
	pointsOfFunctionPlot[1_065].Y = -10.988

	pointsOfFunctionPlot[1_066].X = 10.66
	pointsOfFunctionPlot[1_066].Y = -10.989

	pointsOfFunctionPlot[1_067].X = 10.67
	pointsOfFunctionPlot[1_067].Y = -10.989

	pointsOfFunctionPlot[1_068].X = 10.68
	pointsOfFunctionPlot[1_068].Y = -10.99

	pointsOfFunctionPlot[1_069].X = 10.69
	pointsOfFunctionPlot[1_069].Y = -10.99

	pointsOfFunctionPlot[1_070].X = 10.7
	pointsOfFunctionPlot[1_070].Y = -10.991

	pointsOfFunctionPlot[1_071].X = 10.71
	pointsOfFunctionPlot[1_071].Y = -10.991

	pointsOfFunctionPlot[1_072].X = 10.72
	pointsOfFunctionPlot[1_072].Y = -10.992

	pointsOfFunctionPlot[1_073].X = 10.73
	pointsOfFunctionPlot[1_073].Y = -10.992

	pointsOfFunctionPlot[1_074].X = 10.74
	pointsOfFunctionPlot[1_074].Y = -10.992

	pointsOfFunctionPlot[1_075].X = 10.75
	pointsOfFunctionPlot[1_075].Y = -10.993

	pointsOfFunctionPlot[1_076].X = 10.76
	pointsOfFunctionPlot[1_076].Y = -10.993

	pointsOfFunctionPlot[1_077].X = 10.77
	pointsOfFunctionPlot[1_077].Y = -10.993

	pointsOfFunctionPlot[1_078].X = 10.78
	pointsOfFunctionPlot[1_078].Y = -10.993

	pointsOfFunctionPlot[1_079].X = 10.79
	pointsOfFunctionPlot[1_079].Y = -10.994

	pointsOfFunctionPlot[1_080].X = 10.8
	pointsOfFunctionPlot[1_080].Y = -10.994

	pointsOfFunctionPlot[1_081].X = 10.81
	pointsOfFunctionPlot[1_081].Y = -10.994

	pointsOfFunctionPlot[1_082].X = 10.82
	pointsOfFunctionPlot[1_082].Y = -10.994

	pointsOfFunctionPlot[1_083].X = 10.83
	pointsOfFunctionPlot[1_083].Y = -10.994

	pointsOfFunctionPlot[1_084].X = 10.84
	pointsOfFunctionPlot[1_084].Y = -10.994

	pointsOfFunctionPlot[1_085].X = 10.85
	pointsOfFunctionPlot[1_085].Y = -10.995

	pointsOfFunctionPlot[1_086].X = 10.86
	pointsOfFunctionPlot[1_086].Y = -10.995

	pointsOfFunctionPlot[1_087].X = 10.87
	pointsOfFunctionPlot[1_087].Y = -10.995

	pointsOfFunctionPlot[1_088].X = 10.88
	pointsOfFunctionPlot[1_088].Y = -10.995

	pointsOfFunctionPlot[1_089].X = 10.89
	pointsOfFunctionPlot[1_089].Y = -10.995

	pointsOfFunctionPlot[1_090].X = 10.9
	pointsOfFunctionPlot[1_090].Y = -10.995

	pointsOfFunctionPlot[1_091].X = 10.91
	pointsOfFunctionPlot[1_091].Y = -10.995

	pointsOfFunctionPlot[1_092].X = 10.92
	pointsOfFunctionPlot[1_092].Y = -10.995

	pointsOfFunctionPlot[1_093].X = 10.93
	pointsOfFunctionPlot[1_093].Y = -10.995

	pointsOfFunctionPlot[1_094].X = 10.94
	pointsOfFunctionPlot[1_094].Y = -10.995

	pointsOfFunctionPlot[1_095].X = 10.95
	pointsOfFunctionPlot[1_095].Y = -10.995

	pointsOfFunctionPlot[1_096].X = 10.96
	pointsOfFunctionPlot[1_096].Y = -10.995

	pointsOfFunctionPlot[1_097].X = 10.97
	pointsOfFunctionPlot[1_097].Y = -10.995

	pointsOfFunctionPlot[1_098].X = 10.98
	pointsOfFunctionPlot[1_098].Y = -10.995

	pointsOfFunctionPlot[1_099].X = 10.99
	pointsOfFunctionPlot[1_099].Y = -10.995

	pointsOfFunctionPlot[1_100].X = 11.0
	pointsOfFunctionPlot[1_100].Y = -10.995

	pointsOfFunctionPlot[1_101].X = 11.01
	pointsOfFunctionPlot[1_101].Y = -10.995

	pointsOfFunctionPlot[1_102].X = 11.02
	pointsOfFunctionPlot[1_102].Y = -10.995

	pointsOfFunctionPlot[1_103].X = 11.03
	pointsOfFunctionPlot[1_103].Y = -10.995

	pointsOfFunctionPlot[1_104].X = 11.04
	pointsOfFunctionPlot[1_104].Y = -10.995

	pointsOfFunctionPlot[1_105].X = 11.05
	pointsOfFunctionPlot[1_105].Y = -10.995

	pointsOfFunctionPlot[1_106].X = 11.06
	pointsOfFunctionPlot[1_106].Y = -10.995

	pointsOfFunctionPlot[1_107].X = 11.07
	pointsOfFunctionPlot[1_107].Y = -10.995

	pointsOfFunctionPlot[1_108].X = 11.08
	pointsOfFunctionPlot[1_108].Y = -10.995

	pointsOfFunctionPlot[1_109].X = 11.09
	pointsOfFunctionPlot[1_109].Y = -10.995

	pointsOfFunctionPlot[1_110].X = 11.1
	pointsOfFunctionPlot[1_110].Y = -10.995

	pointsOfFunctionPlot[1_111].X = 11.11
	pointsOfFunctionPlot[1_111].Y = -10.995

	pointsOfFunctionPlot[1_112].X = 11.12
	pointsOfFunctionPlot[1_112].Y = -10.995

	pointsOfFunctionPlot[1_113].X = 11.13
	pointsOfFunctionPlot[1_113].Y = -10.996

	pointsOfFunctionPlot[1_114].X = 11.14
	pointsOfFunctionPlot[1_114].Y = -10.996

	pointsOfFunctionPlot[1_115].X = 11.15
	pointsOfFunctionPlot[1_115].Y = -10.996

	pointsOfFunctionPlot[1_116].X = 11.16
	pointsOfFunctionPlot[1_116].Y = -10.996

	pointsOfFunctionPlot[1_117].X = 11.17
	pointsOfFunctionPlot[1_117].Y = -10.996

	pointsOfFunctionPlot[1_118].X = 11.18
	pointsOfFunctionPlot[1_118].Y = -10.996

	pointsOfFunctionPlot[1_119].X = 11.19
	pointsOfFunctionPlot[1_119].Y = -10.996

	pointsOfFunctionPlot[1_120].X = 11.2
	pointsOfFunctionPlot[1_120].Y = -10.997

	pointsOfFunctionPlot[1_121].X = 11.21
	pointsOfFunctionPlot[1_121].Y = -10.997

	pointsOfFunctionPlot[1_122].X = 11.22
	pointsOfFunctionPlot[1_122].Y = -10.997

	pointsOfFunctionPlot[1_123].X = 11.23
	pointsOfFunctionPlot[1_123].Y = -10.997

	pointsOfFunctionPlot[1_124].X = 11.24
	pointsOfFunctionPlot[1_124].Y = -10.998

	pointsOfFunctionPlot[1_125].X = 11.25
	pointsOfFunctionPlot[1_125].Y = -10.998

	pointsOfFunctionPlot[1_126].X = 11.26
	pointsOfFunctionPlot[1_126].Y = -10.998

	pointsOfFunctionPlot[1_127].X = 11.27
	pointsOfFunctionPlot[1_127].Y = -10.999

	pointsOfFunctionPlot[1_128].X = 11.28
	pointsOfFunctionPlot[1_128].Y = -10.999

	pointsOfFunctionPlot[1_129].X = 11.29
	pointsOfFunctionPlot[1_129].Y = -10.999

	pointsOfFunctionPlot[1_130].X = 11.3
	pointsOfFunctionPlot[1_130].Y = -11.0

	pointsOfFunctionPlot[1_131].X = 11.31
	pointsOfFunctionPlot[1_131].Y = -11.0

	pointsOfFunctionPlot[1_132].X = 11.32
	pointsOfFunctionPlot[1_132].Y = -11.001

	pointsOfFunctionPlot[1_133].X = 11.33
	pointsOfFunctionPlot[1_133].Y = -11.001

	pointsOfFunctionPlot[1_134].X = 11.34
	pointsOfFunctionPlot[1_134].Y = -11.002

	pointsOfFunctionPlot[1_135].X = 11.35
	pointsOfFunctionPlot[1_135].Y = -11.002

	pointsOfFunctionPlot[1_136].X = 11.36
	pointsOfFunctionPlot[1_136].Y = -11.003

	pointsOfFunctionPlot[1_137].X = 11.37
	pointsOfFunctionPlot[1_137].Y = -11.004

	pointsOfFunctionPlot[1_138].X = 11.38
	pointsOfFunctionPlot[1_138].Y = -11.005

	pointsOfFunctionPlot[1_139].X = 11.39
	pointsOfFunctionPlot[1_139].Y = -11.005

	pointsOfFunctionPlot[1_140].X = 11.40
	pointsOfFunctionPlot[1_140].Y = -11.006

	pointsOfFunctionPlot[1_141].X = 11.41
	pointsOfFunctionPlot[1_141].Y = -11.007

	pointsOfFunctionPlot[1_142].X = 11.42
	pointsOfFunctionPlot[1_142].Y = -11.008

	pointsOfFunctionPlot[1_143].X = 11.43
	pointsOfFunctionPlot[1_143].Y = -11.009

	pointsOfFunctionPlot[1_144].X = 11.44
	pointsOfFunctionPlot[1_144].Y = -11.01

	pointsOfFunctionPlot[1_145].X = 11.45
	pointsOfFunctionPlot[1_145].Y = -11.011

	pointsOfFunctionPlot[1_146].X = 11.46
	pointsOfFunctionPlot[1_146].Y = -11.012

	pointsOfFunctionPlot[1_147].X = 11.47
	pointsOfFunctionPlot[1_147].Y = -11.013

	pointsOfFunctionPlot[1_148].X = 11.48
	pointsOfFunctionPlot[1_148].Y = -11.014

	pointsOfFunctionPlot[1_149].X = 11.49
	pointsOfFunctionPlot[1_149].Y = -11.015

	pointsOfFunctionPlot[1_150].X = 11.5
	pointsOfFunctionPlot[1_150].Y = -11.016

	pointsOfFunctionPlot[1_151].X = 11.51
	pointsOfFunctionPlot[1_151].Y = -11.018

	pointsOfFunctionPlot[1_152].X = 11.52
	pointsOfFunctionPlot[1_152].Y = -11.019

	pointsOfFunctionPlot[1_153].X = 11.53
	pointsOfFunctionPlot[1_153].Y = -11.02

	pointsOfFunctionPlot[1_154].X = 11.54
	pointsOfFunctionPlot[1_154].Y = -11.022

	pointsOfFunctionPlot[1_155].X = 11.55
	pointsOfFunctionPlot[1_155].Y = -11.023

	pointsOfFunctionPlot[1_156].X = 11.56
	pointsOfFunctionPlot[1_156].Y = -11.025

	pointsOfFunctionPlot[1_157].X = 11.57
	pointsOfFunctionPlot[1_157].Y = -11.026

	pointsOfFunctionPlot[1_158].X = 11.58
	pointsOfFunctionPlot[1_158].Y = -11.028

	pointsOfFunctionPlot[1_159].X = 11.59
	pointsOfFunctionPlot[1_159].Y = -11.03

	pointsOfFunctionPlot[1_160].X = 11.6
	pointsOfFunctionPlot[1_160].Y = -11.031

	pointsOfFunctionPlot[1_161].X = 11.61
	pointsOfFunctionPlot[1_161].Y = -11.033

	pointsOfFunctionPlot[1_162].X = 11.62
	pointsOfFunctionPlot[1_162].Y = -11.035

	pointsOfFunctionPlot[1_163].X = 11.63
	pointsOfFunctionPlot[1_163].Y = -11.037

	pointsOfFunctionPlot[1_164].X = 11.64
	pointsOfFunctionPlot[1_164].Y = -11.039

	pointsOfFunctionPlot[1_165].X = 11.65
	pointsOfFunctionPlot[1_165].Y = -11.041

	pointsOfFunctionPlot[1_166].X = 11.66
	pointsOfFunctionPlot[1_166].Y = -11.043

	pointsOfFunctionPlot[1_167].X = 11.67
	pointsOfFunctionPlot[1_167].Y = -11.045

	pointsOfFunctionPlot[1_168].X = 11.68
	pointsOfFunctionPlot[1_168].Y = -11.047

	pointsOfFunctionPlot[1_169].X = 11.69
	pointsOfFunctionPlot[1_169].Y = -11.05

	pointsOfFunctionPlot[1_170].X = 11.7
	pointsOfFunctionPlot[1_170].Y = -11.052

	pointsOfFunctionPlot[1_171].X = 11.71
	pointsOfFunctionPlot[1_171].Y = -11.054

	pointsOfFunctionPlot[1_172].X = 11.72
	pointsOfFunctionPlot[1_172].Y = -11.057

	pointsOfFunctionPlot[1_173].X = 11.73
	pointsOfFunctionPlot[1_173].Y = -11.059

	pointsOfFunctionPlot[1_174].X = 11.74
	pointsOfFunctionPlot[1_174].Y = -11.062

	pointsOfFunctionPlot[1_175].X = 11.75
	pointsOfFunctionPlot[1_175].Y = -11.065

	pointsOfFunctionPlot[1_176].X = 11.76
	pointsOfFunctionPlot[1_176].Y = -11.067

	pointsOfFunctionPlot[1_177].X = 11.77
	pointsOfFunctionPlot[1_177].Y = -11.07

	pointsOfFunctionPlot[1_178].X = 11.78
	pointsOfFunctionPlot[1_178].Y = -11.073

	pointsOfFunctionPlot[1_179].X = 11.79
	pointsOfFunctionPlot[1_179].Y = -11.076

	pointsOfFunctionPlot[1_180].X = 11.8
	pointsOfFunctionPlot[1_180].Y = -11.079

	pointsOfFunctionPlot[1_181].X = 11.81
	pointsOfFunctionPlot[1_181].Y = -11.082

	pointsOfFunctionPlot[1_182].X = 11.82
	pointsOfFunctionPlot[1_182].Y = -11.085

	pointsOfFunctionPlot[1_183].X = 11.83
	pointsOfFunctionPlot[1_183].Y = -11.089

	pointsOfFunctionPlot[1_184].X = 11.84
	pointsOfFunctionPlot[1_184].Y = -11.092

	pointsOfFunctionPlot[1_185].X = 11.85
	pointsOfFunctionPlot[1_185].Y = -11.095

	pointsOfFunctionPlot[1_186].X = 11.86
	pointsOfFunctionPlot[1_186].Y = -11.099

	pointsOfFunctionPlot[1_187].X = 11.87
	pointsOfFunctionPlot[1_187].Y = -11.102

	pointsOfFunctionPlot[1_188].X = 11.88
	pointsOfFunctionPlot[1_188].Y = -11.106

	pointsOfFunctionPlot[1_189].X = 11.89
	pointsOfFunctionPlot[1_189].Y = -11.11

	pointsOfFunctionPlot[1_190].X = 11.9
	pointsOfFunctionPlot[1_190].Y = -11.113

	pointsOfFunctionPlot[1_191].X = 11.91
	pointsOfFunctionPlot[1_191].Y = -11.117

	pointsOfFunctionPlot[1_192].X = 11.92
	pointsOfFunctionPlot[1_192].Y = -11.121

	pointsOfFunctionPlot[1_193].X = 11.93
	pointsOfFunctionPlot[1_193].Y = -11.125

	pointsOfFunctionPlot[1_194].X = 11.94
	pointsOfFunctionPlot[1_194].Y = -11.129

	pointsOfFunctionPlot[1_195].X = 11.95
	pointsOfFunctionPlot[1_195].Y = -11.134

	pointsOfFunctionPlot[1_196].X = 11.96
	pointsOfFunctionPlot[1_196].Y = -11.138

	pointsOfFunctionPlot[1_197].X = 11.97
	pointsOfFunctionPlot[1_197].Y = -11.142

	pointsOfFunctionPlot[1_198].X = 11.98
	pointsOfFunctionPlot[1_198].Y = -11.147

	pointsOfFunctionPlot[1_199].X = 11.99
	pointsOfFunctionPlot[1_199].Y = -11.151

	pointsOfFunctionPlot[1_200].X = 12.0
	pointsOfFunctionPlot[1_200].Y = -11.156

	pointsOfFunctionPlot[1_201].X = 12.01
	pointsOfFunctionPlot[1_201].Y = -11.16

	pointsOfFunctionPlot[1_202].X = 12.02
	pointsOfFunctionPlot[1_202].Y = -11.165

	pointsOfFunctionPlot[1_203].X = 12.03
	pointsOfFunctionPlot[1_203].Y = -11.17

	pointsOfFunctionPlot[1_204].X = 12.04
	pointsOfFunctionPlot[1_204].Y = -11.175

	pointsOfFunctionPlot[1_205].X = 12.05
	pointsOfFunctionPlot[1_205].Y = -11.18

	pointsOfFunctionPlot[1_206].X = 12.06
	pointsOfFunctionPlot[1_206].Y = -11.185

	pointsOfFunctionPlot[1_207].X = 12.07
	pointsOfFunctionPlot[1_207].Y = -11.19

	pointsOfFunctionPlot[1_208].X = 12.08
	pointsOfFunctionPlot[1_208].Y = -11.196

	pointsOfFunctionPlot[1_209].X = 12.09
	pointsOfFunctionPlot[1_209].Y = -11.201

	pointsOfFunctionPlot[1_210].X = 12.1
	pointsOfFunctionPlot[1_210].Y = -11.206

	pointsOfFunctionPlot[1_211].X = 12.11
	pointsOfFunctionPlot[1_211].Y = -11.212

	pointsOfFunctionPlot[1_212].X = 12.12
	pointsOfFunctionPlot[1_212].Y = -11.218

	pointsOfFunctionPlot[1_213].X = 12.13
	pointsOfFunctionPlot[1_213].Y = -11.223

	pointsOfFunctionPlot[1_214].X = 12.14
	pointsOfFunctionPlot[1_214].Y = -11.229

	pointsOfFunctionPlot[1_215].X = 12.15
	pointsOfFunctionPlot[1_215].Y = -11.235

	pointsOfFunctionPlot[1_216].X = 12.16
	pointsOfFunctionPlot[1_216].Y = -11.241

	pointsOfFunctionPlot[1_217].X = 12.17
	pointsOfFunctionPlot[1_217].Y = -11.247

	pointsOfFunctionPlot[1_218].X = 12.18
	pointsOfFunctionPlot[1_218].Y = -11.253

	pointsOfFunctionPlot[1_219].X = 12.19
	pointsOfFunctionPlot[1_219].Y = -11.26

	pointsOfFunctionPlot[1_220].X = 12.2
	pointsOfFunctionPlot[1_220].Y = -11.266

	pointsOfFunctionPlot[1_221].X = 12.21
	pointsOfFunctionPlot[1_221].Y = -11.272

	pointsOfFunctionPlot[1_222].X = 12.22
	pointsOfFunctionPlot[1_222].Y = -11.279

	pointsOfFunctionPlot[1_223].X = 12.23
	pointsOfFunctionPlot[1_223].Y = -11.286

	pointsOfFunctionPlot[1_224].X = 12.24
	pointsOfFunctionPlot[1_224].Y = -11.292

	pointsOfFunctionPlot[1_225].X = 12.25
	pointsOfFunctionPlot[1_225].Y = -11.299

	pointsOfFunctionPlot[1_226].X = 12.26
	pointsOfFunctionPlot[1_226].Y = -11.306

	pointsOfFunctionPlot[1_227].X = 12.27
	pointsOfFunctionPlot[1_227].Y = -11.313

	pointsOfFunctionPlot[1_228].X = 12.28
	pointsOfFunctionPlot[1_228].Y = -11.32

	pointsOfFunctionPlot[1_229].X = 12.29
	pointsOfFunctionPlot[1_229].Y = -11.327

	pointsOfFunctionPlot[1_230].X = 12.3
	pointsOfFunctionPlot[1_230].Y = -11.335

	pointsOfFunctionPlot[1_231].X = 12.31
	pointsOfFunctionPlot[1_231].Y = -11.342

	pointsOfFunctionPlot[1_232].X = 12.32
	pointsOfFunctionPlot[1_232].Y = -11.35

	pointsOfFunctionPlot[1_233].X = 12.33
	pointsOfFunctionPlot[1_233].Y = -11.357

	pointsOfFunctionPlot[1_234].X = 12.34
	pointsOfFunctionPlot[1_234].Y = -11.365

	pointsOfFunctionPlot[1_235].X = 12.35
	pointsOfFunctionPlot[1_235].Y = -11.373

	pointsOfFunctionPlot[1_236].X = 12.36
	pointsOfFunctionPlot[1_236].Y = -11.381

	pointsOfFunctionPlot[1_237].X = 12.37
	pointsOfFunctionPlot[1_237].Y = -11.389

	pointsOfFunctionPlot[1_238].X = 12.38
	pointsOfFunctionPlot[1_238].Y = -11.397

	pointsOfFunctionPlot[1_239].X = 12.39
	pointsOfFunctionPlot[1_239].Y = -11.405

	pointsOfFunctionPlot[1_240].X = 12.4
	pointsOfFunctionPlot[1_240].Y = -11.413

	pointsOfFunctionPlot[1_241].X = 12.41
	pointsOfFunctionPlot[1_241].Y = -11.422

	pointsOfFunctionPlot[1_242].X = 12.42
	pointsOfFunctionPlot[1_242].Y = -11.43

	pointsOfFunctionPlot[1_243].X = 12.43
	pointsOfFunctionPlot[1_243].Y = -11.439

	pointsOfFunctionPlot[1_244].X = 12.44
	pointsOfFunctionPlot[1_244].Y = -11.448

	pointsOfFunctionPlot[1_245].X = 12.45
	pointsOfFunctionPlot[1_245].Y = -11.456

	pointsOfFunctionPlot[1_246].X = 12.46
	pointsOfFunctionPlot[1_246].Y = -11.465

	pointsOfFunctionPlot[1_247].X = 12.47
	pointsOfFunctionPlot[1_247].Y = -11.474

	pointsOfFunctionPlot[1_248].X = 12.48
	pointsOfFunctionPlot[1_248].Y = -11.483

	pointsOfFunctionPlot[1_249].X = 12.49
	pointsOfFunctionPlot[1_249].Y = -11.492

	pointsOfFunctionPlot[1_250].X = 12.5
	pointsOfFunctionPlot[1_250].Y = -11.502

	pointsOfFunctionPlot[1_251].X = 12.51
	pointsOfFunctionPlot[1_251].Y = -11.511

	pointsOfFunctionPlot[1_252].X = 12.52
	pointsOfFunctionPlot[1_252].Y = -11.521

	pointsOfFunctionPlot[1_253].X = 12.53
	pointsOfFunctionPlot[1_253].Y = -11.53

	pointsOfFunctionPlot[1_254].X = 12.54
	pointsOfFunctionPlot[1_254].Y = -11.54

	pointsOfFunctionPlot[1_255].X = 12.55
	pointsOfFunctionPlot[1_255].Y = -11.55

	pointsOfFunctionPlot[1_256].X = 12.56
	pointsOfFunctionPlot[1_256].Y = -11.56

	pointsOfFunctionPlot[1_257].X = 12.57
	pointsOfFunctionPlot[1_257].Y = -11.57

	pointsOfFunctionPlot[1_258].X = 12.58
	pointsOfFunctionPlot[1_258].Y = -11.58

	pointsOfFunctionPlot[1_259].X = 12.59
	pointsOfFunctionPlot[1_259].Y = -11.59

	pointsOfFunctionPlot[1_260].X = 12.6
	pointsOfFunctionPlot[1_260].Y = -11.6










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function f(x) = cos(x) - x"

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
		"cos-minus-x-function-plot-02.png"); err != nil {

		panic(err)
	}
}
