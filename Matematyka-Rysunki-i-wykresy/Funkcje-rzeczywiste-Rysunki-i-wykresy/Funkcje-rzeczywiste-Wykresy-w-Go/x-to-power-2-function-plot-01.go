package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function f(x) = x^2.

	pointsOfFunctionPlot := make(plotter.XYs, 2_001)

	pointsOfFunctionPlot[0].X = -10.0
	pointsOfFunctionPlot[0].Y = 100.0

	pointsOfFunctionPlot[1].X = -9.99
	pointsOfFunctionPlot[1].Y = 99.8

	pointsOfFunctionPlot[2].X = -9.98
	pointsOfFunctionPlot[2].Y = 99.6

	pointsOfFunctionPlot[3].X = -9.97
	pointsOfFunctionPlot[3].Y = 99.4

	pointsOfFunctionPlot[4].X = -9.96
	pointsOfFunctionPlot[4].Y = 99.201

	pointsOfFunctionPlot[5].X = -9.95
	pointsOfFunctionPlot[5].Y = 99.002

	pointsOfFunctionPlot[6].X = -9.94
	pointsOfFunctionPlot[6].Y = 98.803

	pointsOfFunctionPlot[7].X = -9.93
	pointsOfFunctionPlot[7].Y = 98.604

	pointsOfFunctionPlot[8].X = -9.92
	pointsOfFunctionPlot[8].Y = 98.406

	pointsOfFunctionPlot[9].X = -9.91
	pointsOfFunctionPlot[9].Y = 98.208

	pointsOfFunctionPlot[10].X = -9.90
	pointsOfFunctionPlot[10].Y = 98.01

	pointsOfFunctionPlot[11].X = -9.89
	pointsOfFunctionPlot[11].Y = 97.812

	pointsOfFunctionPlot[12].X = -9.88
	pointsOfFunctionPlot[12].Y = 97.614

	pointsOfFunctionPlot[13].X = -9.87
	pointsOfFunctionPlot[13].Y = 97.416

	pointsOfFunctionPlot[14].X = -9.86
	pointsOfFunctionPlot[14].Y = 97.219

	pointsOfFunctionPlot[15].X = -9.85
	pointsOfFunctionPlot[15].Y = 97.022

	pointsOfFunctionPlot[16].X = -9.84
	pointsOfFunctionPlot[16].Y = 96.825

	pointsOfFunctionPlot[17].X = -9.83
	pointsOfFunctionPlot[17].Y = 96.628

	pointsOfFunctionPlot[18].X = -9.82
	pointsOfFunctionPlot[18].Y = 96.432

	pointsOfFunctionPlot[19].X = -9.81
	pointsOfFunctionPlot[19].Y = 96.236

	pointsOfFunctionPlot[20].X = -9.80
	pointsOfFunctionPlot[20].Y = 96.04

	pointsOfFunctionPlot[21].X = -9.79
	pointsOfFunctionPlot[21].Y = 95.844

	pointsOfFunctionPlot[22].X = -9.78
	pointsOfFunctionPlot[22].Y = 95.648

	pointsOfFunctionPlot[23].X = -9.77
	pointsOfFunctionPlot[23].Y = 95.452

	pointsOfFunctionPlot[24].X = -9.76
	pointsOfFunctionPlot[24].Y = 95.257

	pointsOfFunctionPlot[25].X = -9.75
	pointsOfFunctionPlot[25].Y = 95.062

	pointsOfFunctionPlot[26].X = -9.74
	pointsOfFunctionPlot[26].Y = 94.867

	pointsOfFunctionPlot[27].X = -9.73
	pointsOfFunctionPlot[27].Y = 94.672

	pointsOfFunctionPlot[28].X = -9.72
	pointsOfFunctionPlot[28].Y = 94.478

	pointsOfFunctionPlot[29].X = -9.71
	pointsOfFunctionPlot[29].Y = 94.284

	pointsOfFunctionPlot[30].X = -9.70
	pointsOfFunctionPlot[30].Y = 94.089

	pointsOfFunctionPlot[31].X = -9.69
	pointsOfFunctionPlot[31].Y = 93.986

	pointsOfFunctionPlot[32].X = -9.68
	pointsOfFunctionPlot[32].Y = 93.702

	pointsOfFunctionPlot[33].X = -9.67
	pointsOfFunctionPlot[33].Y = 93.508

	pointsOfFunctionPlot[34].X = -9.66
	pointsOfFunctionPlot[34].Y = 93.315

	pointsOfFunctionPlot[35].X = -9.65
	pointsOfFunctionPlot[35].Y = 93.122

	pointsOfFunctionPlot[36].X = -9.64
	pointsOfFunctionPlot[36].Y = 92.929

	pointsOfFunctionPlot[37].X = -9.63
	pointsOfFunctionPlot[37].Y = 92.736

	pointsOfFunctionPlot[38].X = -9.62
	pointsOfFunctionPlot[38].Y = 92.544

	pointsOfFunctionPlot[39].X = -9.61
	pointsOfFunctionPlot[39].Y = 92.352

	pointsOfFunctionPlot[40].X = -9.60
	pointsOfFunctionPlot[40].Y = 92.16

	pointsOfFunctionPlot[41].X = -9.59
	pointsOfFunctionPlot[41].Y = 91.968

	pointsOfFunctionPlot[42].X = -9.58
	pointsOfFunctionPlot[42].Y = 91.776

	pointsOfFunctionPlot[43].X = -9.57
	pointsOfFunctionPlot[43].Y = 91.584

	pointsOfFunctionPlot[44].X = -9.56
	pointsOfFunctionPlot[44].Y = 91.393

	pointsOfFunctionPlot[45].X = -9.55
	pointsOfFunctionPlot[45].Y = 91.202

	pointsOfFunctionPlot[46].X = -9.54
	pointsOfFunctionPlot[46].Y = 91.011

	pointsOfFunctionPlot[47].X = -9.53
	pointsOfFunctionPlot[47].Y = 90.82

	pointsOfFunctionPlot[48].X = -9.52
	pointsOfFunctionPlot[48].Y = 90.63

	pointsOfFunctionPlot[49].X = -9.51
	pointsOfFunctionPlot[49].Y = 90.44

	pointsOfFunctionPlot[50].X = -9.50
	pointsOfFunctionPlot[50].Y = 90.25

	pointsOfFunctionPlot[51].X = -9.49
	pointsOfFunctionPlot[51].Y = 90.06

	pointsOfFunctionPlot[52].X = -9.48
	pointsOfFunctionPlot[52].Y = 89.87

	pointsOfFunctionPlot[53].X = -9.47
	pointsOfFunctionPlot[53].Y = 89.68

	pointsOfFunctionPlot[54].X = -9.46
	pointsOfFunctionPlot[54].Y = 89.491

	pointsOfFunctionPlot[55].X = -9.45
	pointsOfFunctionPlot[55].Y = 89.302

	pointsOfFunctionPlot[56].X = -9.44
	pointsOfFunctionPlot[56].Y = 89.113

	pointsOfFunctionPlot[57].X = -9.43
	pointsOfFunctionPlot[57].Y = 88.924

	pointsOfFunctionPlot[58].X = -9.42
	pointsOfFunctionPlot[58].Y = 88.736

	pointsOfFunctionPlot[59].X = -9.41
	pointsOfFunctionPlot[59].Y = 88.548

	pointsOfFunctionPlot[60].X = -9.40
	pointsOfFunctionPlot[60].Y = 88.36

	pointsOfFunctionPlot[61].X = -9.39
	pointsOfFunctionPlot[61].Y = 88.172

	pointsOfFunctionPlot[62].X = -9.38
	pointsOfFunctionPlot[62].Y = 87.984

	pointsOfFunctionPlot[63].X = -9.37
	pointsOfFunctionPlot[63].Y = 87.796

	pointsOfFunctionPlot[64].X = -9.36
	pointsOfFunctionPlot[64].Y = 87.609

	pointsOfFunctionPlot[65].X = -9.35
	pointsOfFunctionPlot[65].Y = 87.422

	pointsOfFunctionPlot[66].X = -9.34
	pointsOfFunctionPlot[66].Y = 87.235

	pointsOfFunctionPlot[67].X = -9.33
	pointsOfFunctionPlot[67].Y = 87.048

	pointsOfFunctionPlot[68].X = -9.32
	pointsOfFunctionPlot[68].Y = 86.862

	pointsOfFunctionPlot[69].X = -9.31
	pointsOfFunctionPlot[69].Y = 86.676

	pointsOfFunctionPlot[70].X = -9.30
	pointsOfFunctionPlot[70].Y = 86.49

	pointsOfFunctionPlot[71].X = -9.29
	pointsOfFunctionPlot[71].Y = 86.304

	pointsOfFunctionPlot[72].X = -9.28
	pointsOfFunctionPlot[72].Y = 86.118

	pointsOfFunctionPlot[73].X = -9.27
	pointsOfFunctionPlot[73].Y = 85.932

	pointsOfFunctionPlot[74].X = -9.26
	pointsOfFunctionPlot[74].Y = 85.747

	pointsOfFunctionPlot[75].X = -9.25
	pointsOfFunctionPlot[75].Y = 85.562

	pointsOfFunctionPlot[76].X = -9.24
	pointsOfFunctionPlot[76].Y = 85.377

	pointsOfFunctionPlot[77].X = -9.23
	pointsOfFunctionPlot[77].Y = 85.192

	pointsOfFunctionPlot[78].X = -9.22
	pointsOfFunctionPlot[78].Y = 85.008

	pointsOfFunctionPlot[79].X = -9.21
	pointsOfFunctionPlot[79].Y = 84.824

	pointsOfFunctionPlot[80].X = -9.20
	pointsOfFunctionPlot[80].Y = 84.64

	pointsOfFunctionPlot[81].X = -9.19
	pointsOfFunctionPlot[81].Y = 84.456

	pointsOfFunctionPlot[82].X = -9.18
	pointsOfFunctionPlot[82].Y = 84.272

	pointsOfFunctionPlot[83].X = -9.17
	pointsOfFunctionPlot[83].Y = 84.088

	pointsOfFunctionPlot[84].X = -9.16
	pointsOfFunctionPlot[84].Y = 83.905

	pointsOfFunctionPlot[85].X = -9.15
	pointsOfFunctionPlot[85].Y = 83.722

	pointsOfFunctionPlot[86].X = -9.14
	pointsOfFunctionPlot[86].Y = 83.539

	pointsOfFunctionPlot[87].X = -9.13
	pointsOfFunctionPlot[87].Y = 83.356

	pointsOfFunctionPlot[88].X = -9.12
	pointsOfFunctionPlot[88].Y = 83.174

	pointsOfFunctionPlot[89].X = -9.11
	pointsOfFunctionPlot[89].Y = 82.992

	pointsOfFunctionPlot[90].X = -9.10
	pointsOfFunctionPlot[90].Y = 82.81

	pointsOfFunctionPlot[91].X = -9.09
	pointsOfFunctionPlot[91].Y = 82.628

	pointsOfFunctionPlot[92].X = -9.08
	pointsOfFunctionPlot[92].Y = 82.446

	pointsOfFunctionPlot[93].X = -9.07
	pointsOfFunctionPlot[93].Y = 82.264

	pointsOfFunctionPlot[94].X = -9.06
	pointsOfFunctionPlot[94].Y = 82.083

	pointsOfFunctionPlot[95].X = -9.05
	pointsOfFunctionPlot[95].Y = 81.902

	pointsOfFunctionPlot[96].X = -9.04
	pointsOfFunctionPlot[96].Y = 81.721

	pointsOfFunctionPlot[97].X = -9.03
	pointsOfFunctionPlot[97].Y = 81.54

	pointsOfFunctionPlot[98].X = -9.02
	pointsOfFunctionPlot[98].Y = 81.36

	pointsOfFunctionPlot[99].X = -9.01
	pointsOfFunctionPlot[99].Y = 81.18

	pointsOfFunctionPlot[100].X = -9.0
	pointsOfFunctionPlot[100].Y = 81.0

	pointsOfFunctionPlot[101].X = -8.99
	pointsOfFunctionPlot[101].Y = 80.82

	pointsOfFunctionPlot[102].X = -8.98
	pointsOfFunctionPlot[102].Y = 80.64

	pointsOfFunctionPlot[103].X = -8.97
	pointsOfFunctionPlot[103].Y = 80.46

	pointsOfFunctionPlot[104].X = -8.96
	pointsOfFunctionPlot[104].Y = 80.281

	pointsOfFunctionPlot[105].X = -8.95
	pointsOfFunctionPlot[105].Y = 80.102

	pointsOfFunctionPlot[106].X = -8.94
	pointsOfFunctionPlot[106].Y = 79.923

	pointsOfFunctionPlot[107].X = -8.93
	pointsOfFunctionPlot[107].Y = 79.744

	pointsOfFunctionPlot[108].X = -8.92
	pointsOfFunctionPlot[108].Y = 79.566

	pointsOfFunctionPlot[109].X = -8.91
	pointsOfFunctionPlot[109].Y = 79.388

	pointsOfFunctionPlot[110].X = -8.90
	pointsOfFunctionPlot[110].Y = 79.21

	pointsOfFunctionPlot[111].X = -8.89
	pointsOfFunctionPlot[111].Y = 79.032

	pointsOfFunctionPlot[112].X = -8.88
	pointsOfFunctionPlot[112].Y = 78.854

	pointsOfFunctionPlot[113].X = -8.87
	pointsOfFunctionPlot[113].Y = 78.676

	pointsOfFunctionPlot[114].X = -8.86
	pointsOfFunctionPlot[114].Y = 78.499

	pointsOfFunctionPlot[115].X = -8.85
	pointsOfFunctionPlot[115].Y = 78.322

	pointsOfFunctionPlot[116].X = -8.84
	pointsOfFunctionPlot[116].Y = 78.145

	pointsOfFunctionPlot[117].X = -8.83
	pointsOfFunctionPlot[117].Y = 77.968

	pointsOfFunctionPlot[118].X = -8.82
	pointsOfFunctionPlot[118].Y = 77.792

	pointsOfFunctionPlot[119].X = -8.81
	pointsOfFunctionPlot[119].Y = 77.616

	pointsOfFunctionPlot[120].X = -8.80
	pointsOfFunctionPlot[120].Y = 77.44

	pointsOfFunctionPlot[121].X = -8.79
	pointsOfFunctionPlot[121].Y = 77.264

	pointsOfFunctionPlot[122].X = -8.78
	pointsOfFunctionPlot[122].Y = 77.088

	pointsOfFunctionPlot[123].X = -8.77
	pointsOfFunctionPlot[123].Y = 76.912

	pointsOfFunctionPlot[124].X = -8.76
	pointsOfFunctionPlot[124].Y = 76.737

	pointsOfFunctionPlot[125].X = -8.75
	pointsOfFunctionPlot[125].Y = 76.562

	pointsOfFunctionPlot[126].X = -8.74
	pointsOfFunctionPlot[126].Y = 76.387

	pointsOfFunctionPlot[127].X = -8.73
	pointsOfFunctionPlot[127].Y = 76.212

	pointsOfFunctionPlot[128].X = -8.72
	pointsOfFunctionPlot[128].Y = 76.038

	pointsOfFunctionPlot[129].X = -8.71
	pointsOfFunctionPlot[129].Y = 75.864

	pointsOfFunctionPlot[130].X = -8.70
	pointsOfFunctionPlot[130].Y = 75.69

	pointsOfFunctionPlot[131].X = -8.69
	pointsOfFunctionPlot[131].Y = 75.516

	pointsOfFunctionPlot[132].X = -8.68
	pointsOfFunctionPlot[132].Y = 75.342

	pointsOfFunctionPlot[133].X = -8.67
	pointsOfFunctionPlot[133].Y = 75.168

	pointsOfFunctionPlot[134].X = -8.66
	pointsOfFunctionPlot[134].Y = 74.995

	pointsOfFunctionPlot[135].X = -8.65
	pointsOfFunctionPlot[135].Y = 74.882

	pointsOfFunctionPlot[136].X = -8.64
	pointsOfFunctionPlot[136].Y = 74.649

	pointsOfFunctionPlot[137].X = -8.63
	pointsOfFunctionPlot[137].Y = 74.476

	pointsOfFunctionPlot[138].X = -8.62
	pointsOfFunctionPlot[138].Y = 74.304

	pointsOfFunctionPlot[139].X = -8.61
	pointsOfFunctionPlot[139].Y = 74.132

	pointsOfFunctionPlot[140].X = -8.60
	pointsOfFunctionPlot[140].Y = 73.96

	pointsOfFunctionPlot[141].X = -8.59
	pointsOfFunctionPlot[141].Y = 73.788

	pointsOfFunctionPlot[142].X = -8.58
	pointsOfFunctionPlot[142].Y = 73.616

	pointsOfFunctionPlot[143].X = -8.57
	pointsOfFunctionPlot[143].Y = 73.444

	pointsOfFunctionPlot[144].X = -8.56
	pointsOfFunctionPlot[144].Y = 73.273

	pointsOfFunctionPlot[145].X = -8.55
	pointsOfFunctionPlot[145].Y = 73.102

	pointsOfFunctionPlot[146].X = -8.54
	pointsOfFunctionPlot[146].Y = 72.931

	pointsOfFunctionPlot[147].X = -8.53
	pointsOfFunctionPlot[147].Y = 72.76

	pointsOfFunctionPlot[148].X = -8.52
	pointsOfFunctionPlot[148].Y = 72.59

	pointsOfFunctionPlot[149].X = -8.51
	pointsOfFunctionPlot[149].Y = 72.42

	pointsOfFunctionPlot[150].X = -8.50
	pointsOfFunctionPlot[150].Y = 72.25

	pointsOfFunctionPlot[151].X = -8.49
	pointsOfFunctionPlot[151].Y = 72.08

	pointsOfFunctionPlot[152].X = -8.48
	pointsOfFunctionPlot[152].Y = 71.91

	pointsOfFunctionPlot[153].X = -8.47
	pointsOfFunctionPlot[153].Y = 71.74

	pointsOfFunctionPlot[154].X = -8.46
	pointsOfFunctionPlot[154].Y = 71.57

	pointsOfFunctionPlot[155].X = -8.45
	pointsOfFunctionPlot[155].Y = 71.402

	pointsOfFunctionPlot[156].X = -8.44
	pointsOfFunctionPlot[156].Y = 71.233

	pointsOfFunctionPlot[157].X = -8.43
	pointsOfFunctionPlot[157].Y = 71.064

	pointsOfFunctionPlot[158].X = -8.42
	pointsOfFunctionPlot[158].Y = 70.896

	pointsOfFunctionPlot[159].X = -8.41
	pointsOfFunctionPlot[159].Y = 70.728

	pointsOfFunctionPlot[160].X = -8.40
	pointsOfFunctionPlot[160].Y = 70.56

	pointsOfFunctionPlot[161].X = -8.39
	pointsOfFunctionPlot[161].Y = 70.392

	pointsOfFunctionPlot[162].X = -8.38
	pointsOfFunctionPlot[162].Y = 70.224

	pointsOfFunctionPlot[163].X = -8.37
	pointsOfFunctionPlot[163].Y = 70.056

	pointsOfFunctionPlot[164].X = -8.36
	pointsOfFunctionPlot[164].Y = 69.889

	pointsOfFunctionPlot[165].X = -8.35
	pointsOfFunctionPlot[165].Y = 69.722

	pointsOfFunctionPlot[166].X = -8.34
	pointsOfFunctionPlot[166].Y = 69.555

	pointsOfFunctionPlot[167].X = -8.33
	pointsOfFunctionPlot[167].Y = 69.388

	pointsOfFunctionPlot[168].X = -8.32
	pointsOfFunctionPlot[168].Y = 69.222

	pointsOfFunctionPlot[169].X = -8.31
	pointsOfFunctionPlot[169].Y = 69.056

	pointsOfFunctionPlot[170].X = -8.30
	pointsOfFunctionPlot[170].Y = 68.89

	pointsOfFunctionPlot[171].X = -8.29
	pointsOfFunctionPlot[171].Y = 68.721

	pointsOfFunctionPlot[172].X = -8.28
	pointsOfFunctionPlot[172].Y = 68.558

	pointsOfFunctionPlot[173].X = -8.27
	pointsOfFunctionPlot[173].Y = 68.392

	pointsOfFunctionPlot[174].X = -8.26
	pointsOfFunctionPlot[174].Y = 68.227

	pointsOfFunctionPlot[175].X = -8.25
	pointsOfFunctionPlot[175].Y = 68.062

	pointsOfFunctionPlot[176].X = -8.24
	pointsOfFunctionPlot[176].Y = 67.897

	pointsOfFunctionPlot[177].X = -8.23
	pointsOfFunctionPlot[177].Y = 67.732

	pointsOfFunctionPlot[178].X = -8.22
	pointsOfFunctionPlot[178].Y = 67.568

	pointsOfFunctionPlot[179].X = -8.21
	pointsOfFunctionPlot[179].Y = 67.404

	pointsOfFunctionPlot[180].X = -8.20
	pointsOfFunctionPlot[180].Y = 67.24

	pointsOfFunctionPlot[181].X = -8.19
	pointsOfFunctionPlot[181].Y = 67.076

	pointsOfFunctionPlot[182].X = -8.18
	pointsOfFunctionPlot[182].Y = 66.912

	pointsOfFunctionPlot[183].X = -8.17
	pointsOfFunctionPlot[183].Y = 66.748

	pointsOfFunctionPlot[184].X = -8.16
	pointsOfFunctionPlot[184].Y = 66.585

	pointsOfFunctionPlot[185].X = -8.15
	pointsOfFunctionPlot[185].Y = 66.422

	pointsOfFunctionPlot[186].X = -8.14
	pointsOfFunctionPlot[186].Y = 66.259

	pointsOfFunctionPlot[187].X = -8.13
	pointsOfFunctionPlot[187].Y = 66.096

	pointsOfFunctionPlot[188].X = -8.12
	pointsOfFunctionPlot[188].Y = 65.934

	pointsOfFunctionPlot[189].X = -8.11
	pointsOfFunctionPlot[189].Y = 65.772

	pointsOfFunctionPlot[190].X = -8.10
	pointsOfFunctionPlot[190].Y = 65.61

	pointsOfFunctionPlot[191].X = -8.09
	pointsOfFunctionPlot[191].Y = 65.448

	pointsOfFunctionPlot[192].X = -8.08
	pointsOfFunctionPlot[192].Y = 65.286

	pointsOfFunctionPlot[193].X = -8.07
	pointsOfFunctionPlot[193].Y = 65.124

	pointsOfFunctionPlot[194].X = -8.06
	pointsOfFunctionPlot[194].Y = 64.963

	pointsOfFunctionPlot[195].X = -8.05
	pointsOfFunctionPlot[195].Y = 64.802

	pointsOfFunctionPlot[196].X = -8.04
	pointsOfFunctionPlot[196].Y = 64.641

	pointsOfFunctionPlot[197].X = -8.03
	pointsOfFunctionPlot[197].Y = 64.48

	pointsOfFunctionPlot[198].X = -8.02
	pointsOfFunctionPlot[198].Y = 64.32

	pointsOfFunctionPlot[199].X = -8.01
	pointsOfFunctionPlot[199].Y = 64.16

	pointsOfFunctionPlot[200].X = -8.0
	pointsOfFunctionPlot[200].Y = 64.0

	pointsOfFunctionPlot[201].X = -7.99
	pointsOfFunctionPlot[201].Y = 63.84

	pointsOfFunctionPlot[202].X = -7.98
	pointsOfFunctionPlot[202].Y = 63.68

	pointsOfFunctionPlot[203].X = -7.97
	pointsOfFunctionPlot[203].Y = 63.52

	pointsOfFunctionPlot[204].X = -7.96
	pointsOfFunctionPlot[204].Y = 63.361

	pointsOfFunctionPlot[205].X = -7.95
	pointsOfFunctionPlot[205].Y = 63.202

	pointsOfFunctionPlot[206].X = -7.94
	pointsOfFunctionPlot[206].Y = 63.043

	pointsOfFunctionPlot[207].X = -7.93
	pointsOfFunctionPlot[207].Y = 62.884

	pointsOfFunctionPlot[208].X = -7.92
	pointsOfFunctionPlot[208].Y = 62.726

	pointsOfFunctionPlot[209].X = -7.91
	pointsOfFunctionPlot[209].Y = 62.568

	pointsOfFunctionPlot[210].X = -7.90
	pointsOfFunctionPlot[210].Y = 62.41

	pointsOfFunctionPlot[211].X = -7.89
	pointsOfFunctionPlot[211].Y = 62.252

	pointsOfFunctionPlot[212].X = -7.88
	pointsOfFunctionPlot[212].Y = 62.094

	pointsOfFunctionPlot[213].X = -7.87
	pointsOfFunctionPlot[213].Y = 61.936

	pointsOfFunctionPlot[214].X = -7.86
	pointsOfFunctionPlot[214].Y = 61.779

	pointsOfFunctionPlot[215].X = -7.85
	pointsOfFunctionPlot[215].Y = 61.622

	pointsOfFunctionPlot[216].X = -7.84
	pointsOfFunctionPlot[216].Y = 61.465

	pointsOfFunctionPlot[217].X = -7.83
	pointsOfFunctionPlot[217].Y = 61.308

	pointsOfFunctionPlot[218].X = -7.82
	pointsOfFunctionPlot[218].Y = 61.152

	pointsOfFunctionPlot[219].X = -7.81
	pointsOfFunctionPlot[219].Y = 60.996

	pointsOfFunctionPlot[220].X = -7.80
	pointsOfFunctionPlot[220].Y = 60.84

	pointsOfFunctionPlot[221].X = -7.79
	pointsOfFunctionPlot[221].Y = 60.684

	pointsOfFunctionPlot[222].X = -7.78
	pointsOfFunctionPlot[222].Y = 60.528

	pointsOfFunctionPlot[223].X = -7.77
	pointsOfFunctionPlot[223].Y = 60.372

	pointsOfFunctionPlot[224].X = -7.76
	pointsOfFunctionPlot[224].Y = 60.217

	pointsOfFunctionPlot[225].X = -7.75
	pointsOfFunctionPlot[225].Y = 60.062

	pointsOfFunctionPlot[226].X = -7.74
	pointsOfFunctionPlot[226].Y = 59.907

	pointsOfFunctionPlot[227].X = -7.73
	pointsOfFunctionPlot[227].Y = 59.752

	pointsOfFunctionPlot[228].X = -7.72
	pointsOfFunctionPlot[228].Y = 59.598

	pointsOfFunctionPlot[229].X = -7.71
	pointsOfFunctionPlot[229].Y = 59.444

	pointsOfFunctionPlot[230].X = -7.70
	pointsOfFunctionPlot[230].Y = 59.29

	pointsOfFunctionPlot[231].X = -7.69
	pointsOfFunctionPlot[231].Y = 59.136

	pointsOfFunctionPlot[232].X = -7.68
	pointsOfFunctionPlot[232].Y = 58.982

	pointsOfFunctionPlot[233].X = -7.67
	pointsOfFunctionPlot[233].Y = 58.828

	pointsOfFunctionPlot[234].X = -7.66
	pointsOfFunctionPlot[234].Y = 58.675

	pointsOfFunctionPlot[235].X = -7.65
	pointsOfFunctionPlot[235].Y = 58.522

	pointsOfFunctionPlot[236].X = -7.64
	pointsOfFunctionPlot[236].Y = 58.369

	pointsOfFunctionPlot[237].X = -7.63
	pointsOfFunctionPlot[237].Y = 58.216

	pointsOfFunctionPlot[238].X = -7.62
	pointsOfFunctionPlot[238].Y = 58.064

	pointsOfFunctionPlot[239].X = -7.61
	pointsOfFunctionPlot[239].Y = 57.912

	pointsOfFunctionPlot[240].X = -7.60
	pointsOfFunctionPlot[240].Y = 57.76

	pointsOfFunctionPlot[241].X = -7.59
	pointsOfFunctionPlot[241].Y = 57.608

	pointsOfFunctionPlot[242].X = -7.58
	pointsOfFunctionPlot[242].Y = 57.456

	pointsOfFunctionPlot[243].X = -7.57
	pointsOfFunctionPlot[243].Y = 57.304

	pointsOfFunctionPlot[244].X = -7.56
	pointsOfFunctionPlot[244].Y = 57.153

	pointsOfFunctionPlot[245].X = -7.55
	pointsOfFunctionPlot[245].Y = 57.002

	pointsOfFunctionPlot[246].X = -7.54
	pointsOfFunctionPlot[246].Y = 56.851

	pointsOfFunctionPlot[247].X = -7.53
	pointsOfFunctionPlot[247].Y = 56.7

	pointsOfFunctionPlot[248].X = -7.52
	pointsOfFunctionPlot[248].Y = 56.55

	pointsOfFunctionPlot[249].X = -7.51
	pointsOfFunctionPlot[249].Y = 56.4

	pointsOfFunctionPlot[250].X = -7.50
	pointsOfFunctionPlot[250].Y = 56.25

	pointsOfFunctionPlot[251].X = -7.49
	pointsOfFunctionPlot[251].Y = 56.1

	pointsOfFunctionPlot[252].X = -7.48
	pointsOfFunctionPlot[252].Y = 55.95

	pointsOfFunctionPlot[253].X = -7.47
	pointsOfFunctionPlot[253].Y = 55.8

	pointsOfFunctionPlot[254].X = -7.46
	pointsOfFunctionPlot[254].Y = 55.651

	pointsOfFunctionPlot[255].X = -7.45
	pointsOfFunctionPlot[255].Y = 55.502

	pointsOfFunctionPlot[256].X = -7.44
	pointsOfFunctionPlot[256].Y = 55.353

	pointsOfFunctionPlot[257].X = -7.43
	pointsOfFunctionPlot[257].Y = 55.204

	pointsOfFunctionPlot[258].X = -7.42
	pointsOfFunctionPlot[258].Y = 55.056

	pointsOfFunctionPlot[259].X = -7.41
	pointsOfFunctionPlot[259].Y = 54.908

	pointsOfFunctionPlot[260].X = -7.40
	pointsOfFunctionPlot[260].Y = 54.76

	pointsOfFunctionPlot[261].X = -7.39
	pointsOfFunctionPlot[261].Y = 54.612

	pointsOfFunctionPlot[262].X = -7.38
	pointsOfFunctionPlot[262].Y = 54.464

	pointsOfFunctionPlot[263].X = -7.37
	pointsOfFunctionPlot[263].Y = 54.316

	pointsOfFunctionPlot[264].X = -7.36
	pointsOfFunctionPlot[264].Y = 54.169

	pointsOfFunctionPlot[265].X = -7.35
	pointsOfFunctionPlot[265].Y = 54.022

	pointsOfFunctionPlot[266].X = -7.34
	pointsOfFunctionPlot[266].Y = 53.875

	pointsOfFunctionPlot[267].X = -7.33
	pointsOfFunctionPlot[267].Y = 53.728

	pointsOfFunctionPlot[268].X = -7.32
	pointsOfFunctionPlot[268].Y = 53.582

	pointsOfFunctionPlot[269].X = -7.31
	pointsOfFunctionPlot[269].Y = 53.436

	pointsOfFunctionPlot[270].X = -7.30
	pointsOfFunctionPlot[270].Y = 53.29

	pointsOfFunctionPlot[271].X = -7.29
	pointsOfFunctionPlot[271].Y = 53.144

	pointsOfFunctionPlot[272].X = -7.28
	pointsOfFunctionPlot[272].Y = 52.998

	pointsOfFunctionPlot[273].X = -7.27
	pointsOfFunctionPlot[273].Y = 52.852

	pointsOfFunctionPlot[274].X = -7.26
	pointsOfFunctionPlot[274].Y = 52.707

	pointsOfFunctionPlot[275].X = -7.25
	pointsOfFunctionPlot[275].Y = 52.562

	pointsOfFunctionPlot[276].X = -7.24
	pointsOfFunctionPlot[276].Y = 52.417

	pointsOfFunctionPlot[277].X = -7.23
	pointsOfFunctionPlot[277].Y = 52.272

	pointsOfFunctionPlot[278].X = -7.22
	pointsOfFunctionPlot[278].Y = 52.128

	pointsOfFunctionPlot[279].X = -7.21
	pointsOfFunctionPlot[279].Y = 51.984

	pointsOfFunctionPlot[280].X = -7.20
	pointsOfFunctionPlot[280].Y = 51.84

	pointsOfFunctionPlot[281].X = -7.19
	pointsOfFunctionPlot[281].Y = 51.696

	pointsOfFunctionPlot[282].X = -7.18
	pointsOfFunctionPlot[282].Y = 51.552

	pointsOfFunctionPlot[283].X = -7.17
	pointsOfFunctionPlot[283].Y = 51.408

	pointsOfFunctionPlot[284].X = -7.16
	pointsOfFunctionPlot[284].Y = 51.265

	pointsOfFunctionPlot[285].X = -7.15
	pointsOfFunctionPlot[285].Y = 51.122

	pointsOfFunctionPlot[286].X = -7.14
	pointsOfFunctionPlot[286].Y = 50.979

	pointsOfFunctionPlot[287].X = -7.13
	pointsOfFunctionPlot[287].Y = 50.836

	pointsOfFunctionPlot[288].X = -7.12
	pointsOfFunctionPlot[288].Y = 50.694

	pointsOfFunctionPlot[289].X = -7.11
	pointsOfFunctionPlot[289].Y = 50.522

	pointsOfFunctionPlot[290].X = -7.10
	pointsOfFunctionPlot[290].Y = 50.41

	pointsOfFunctionPlot[291].X = -7.09
	pointsOfFunctionPlot[291].Y = 50.268

	pointsOfFunctionPlot[292].X = -7.08
	pointsOfFunctionPlot[292].Y = 50.126

	pointsOfFunctionPlot[293].X = -7.07
	pointsOfFunctionPlot[293].Y = 49.984

	pointsOfFunctionPlot[294].X = -7.06
	pointsOfFunctionPlot[294].Y = 49.843

	pointsOfFunctionPlot[295].X = -7.05
	pointsOfFunctionPlot[295].Y = 49.702

	pointsOfFunctionPlot[296].X = -7.04
	pointsOfFunctionPlot[296].Y = 49.561

	pointsOfFunctionPlot[297].X = -7.03
	pointsOfFunctionPlot[297].Y = 49.42

	pointsOfFunctionPlot[298].X = -7.02
	pointsOfFunctionPlot[298].Y = 49.28

	pointsOfFunctionPlot[299].X = -7.01
	pointsOfFunctionPlot[299].Y = 49.14

	pointsOfFunctionPlot[300].X = -7.0
	pointsOfFunctionPlot[300].Y = 49.0

	pointsOfFunctionPlot[301].X = -6.99
	pointsOfFunctionPlot[301].Y = 48.86

	pointsOfFunctionPlot[302].X = -6.98
	pointsOfFunctionPlot[302].Y = 48.72

	pointsOfFunctionPlot[303].X = -6.97
	pointsOfFunctionPlot[303].Y = 48.58

	pointsOfFunctionPlot[304].X = -6.96
	pointsOfFunctionPlot[304].Y = 48.441

	pointsOfFunctionPlot[305].X = -6.95
	pointsOfFunctionPlot[305].Y = 48.302

	pointsOfFunctionPlot[306].X = -6.94
	pointsOfFunctionPlot[306].Y = 48.163

	pointsOfFunctionPlot[307].X = -6.93
	pointsOfFunctionPlot[307].Y = 48.024

	pointsOfFunctionPlot[308].X = -6.92
	pointsOfFunctionPlot[308].Y = 47.886

	pointsOfFunctionPlot[309].X = -6.91
	pointsOfFunctionPlot[309].Y = 47.748

	pointsOfFunctionPlot[310].X = -6.90
	pointsOfFunctionPlot[310].Y = 47.61

	pointsOfFunctionPlot[311].X = -6.89
	pointsOfFunctionPlot[311].Y = 47.472

	pointsOfFunctionPlot[312].X = -6.88
	pointsOfFunctionPlot[312].Y = 47.334

	pointsOfFunctionPlot[313].X = -6.87
	pointsOfFunctionPlot[313].Y = 47.196

	pointsOfFunctionPlot[314].X = -6.86
	pointsOfFunctionPlot[314].Y = 47.059

	pointsOfFunctionPlot[315].X = -6.85
	pointsOfFunctionPlot[315].Y = 46.922

	pointsOfFunctionPlot[316].X = -6.84
	pointsOfFunctionPlot[316].Y = 46.785

	pointsOfFunctionPlot[317].X = -6.83
	pointsOfFunctionPlot[317].Y = 46.648

	pointsOfFunctionPlot[318].X = -6.82
	pointsOfFunctionPlot[318].Y = 46.512

	pointsOfFunctionPlot[319].X = -6.81
	pointsOfFunctionPlot[319].Y = 46.376

	pointsOfFunctionPlot[320].X = -6.80
	pointsOfFunctionPlot[320].Y = 46.24

	pointsOfFunctionPlot[321].X = -6.79
	pointsOfFunctionPlot[321].Y = 46.104

	pointsOfFunctionPlot[322].X = -6.78
	pointsOfFunctionPlot[322].Y = 45.968

	pointsOfFunctionPlot[323].X = -6.77
	pointsOfFunctionPlot[323].Y = 45.832

	pointsOfFunctionPlot[324].X = -6.76
	pointsOfFunctionPlot[324].Y = 45.697

	pointsOfFunctionPlot[325].X = -6.75
	pointsOfFunctionPlot[325].Y = 45.562

	pointsOfFunctionPlot[326].X = -6.74
	pointsOfFunctionPlot[326].Y = 45.427

	pointsOfFunctionPlot[327].X = -6.73
	pointsOfFunctionPlot[327].Y = 45.292

	pointsOfFunctionPlot[328].X = -6.72
	pointsOfFunctionPlot[328].Y = 45.158

	pointsOfFunctionPlot[329].X = -6.71
	pointsOfFunctionPlot[329].Y = 45.024

	pointsOfFunctionPlot[330].X = -6.70
	pointsOfFunctionPlot[330].Y = 44.89

	pointsOfFunctionPlot[331].X = -6.69
	pointsOfFunctionPlot[331].Y = 44.756

	pointsOfFunctionPlot[332].X = -6.68
	pointsOfFunctionPlot[332].Y = 44.622

	pointsOfFunctionPlot[333].X = -6.67
	pointsOfFunctionPlot[333].Y = 44.488

	pointsOfFunctionPlot[334].X = -6.66
	pointsOfFunctionPlot[334].Y = 44.355

	pointsOfFunctionPlot[335].X = -6.65
	pointsOfFunctionPlot[335].Y = 44.222

	pointsOfFunctionPlot[336].X = -6.64
	pointsOfFunctionPlot[336].Y = 44.089

	pointsOfFunctionPlot[337].X = -6.63
	pointsOfFunctionPlot[337].Y = 43.956

	pointsOfFunctionPlot[338].X = -6.62
	pointsOfFunctionPlot[338].Y = 43.824

	pointsOfFunctionPlot[339].X = -6.61
	pointsOfFunctionPlot[339].Y = 43.692

	pointsOfFunctionPlot[340].X = -6.60
	pointsOfFunctionPlot[340].Y = 43.56

	pointsOfFunctionPlot[341].X = -6.59
	pointsOfFunctionPlot[341].Y = 43.428

	pointsOfFunctionPlot[342].X = -6.58
	pointsOfFunctionPlot[342].Y = 43.296

	pointsOfFunctionPlot[343].X = -6.57
	pointsOfFunctionPlot[343].Y = 43.164

	pointsOfFunctionPlot[344].X = -6.56
	pointsOfFunctionPlot[344].Y = 43.033

	pointsOfFunctionPlot[345].X = -6.55
	pointsOfFunctionPlot[345].Y = 42.902

	pointsOfFunctionPlot[346].X = -6.54
	pointsOfFunctionPlot[346].Y = 42.771

	pointsOfFunctionPlot[347].X = -6.53
	pointsOfFunctionPlot[347].Y = 42.64

	pointsOfFunctionPlot[348].X = -6.52
	pointsOfFunctionPlot[348].Y = 42.51

	pointsOfFunctionPlot[349].X = -6.51
	pointsOfFunctionPlot[349].Y = 42.38

	pointsOfFunctionPlot[350].X = -6.50
	pointsOfFunctionPlot[350].Y = 42.25

	pointsOfFunctionPlot[351].X = -6.49
	pointsOfFunctionPlot[351].Y = 42.12

	pointsOfFunctionPlot[352].X = -6.48
	pointsOfFunctionPlot[352].Y = 41.99

	pointsOfFunctionPlot[353].X = -6.47
	pointsOfFunctionPlot[353].Y = 41.86

	pointsOfFunctionPlot[354].X = -6.46
	pointsOfFunctionPlot[354].Y = 41.731

	pointsOfFunctionPlot[355].X = -6.45
	pointsOfFunctionPlot[355].Y = 41.602

	pointsOfFunctionPlot[356].X = -6.44
	pointsOfFunctionPlot[356].Y = 41.473

	pointsOfFunctionPlot[357].X = -6.43
	pointsOfFunctionPlot[357].Y = 41.344

	pointsOfFunctionPlot[358].X = -6.42
	pointsOfFunctionPlot[358].Y = 41.216

	pointsOfFunctionPlot[359].X = -6.41
	pointsOfFunctionPlot[359].Y = 41.088

	pointsOfFunctionPlot[360].X = -6.40
	pointsOfFunctionPlot[360].Y = 40.96

	pointsOfFunctionPlot[361].X = -6.39
	pointsOfFunctionPlot[361].Y = 40.832

	pointsOfFunctionPlot[362].X = -6.38
	pointsOfFunctionPlot[362].Y = 40.704

	pointsOfFunctionPlot[363].X = -6.37
	pointsOfFunctionPlot[363].Y = 40.576

	pointsOfFunctionPlot[364].X = -6.36
	pointsOfFunctionPlot[364].Y = 40.449

	pointsOfFunctionPlot[365].X = -6.35
	pointsOfFunctionPlot[365].Y = 40.322

	pointsOfFunctionPlot[366].X = -6.34
	pointsOfFunctionPlot[366].Y = 40.195

	pointsOfFunctionPlot[367].X = -6.33
	pointsOfFunctionPlot[367].Y = 40.068

	pointsOfFunctionPlot[368].X = -6.32
	pointsOfFunctionPlot[368].Y = 39.942

	pointsOfFunctionPlot[369].X = -6.31
	pointsOfFunctionPlot[369].Y = 39.816

	pointsOfFunctionPlot[370].X = -6.30
	pointsOfFunctionPlot[370].Y = 39.69

	pointsOfFunctionPlot[371].X = -6.29
	pointsOfFunctionPlot[371].Y = 39.564

	pointsOfFunctionPlot[372].X = -6.28
	pointsOfFunctionPlot[372].Y = 39.438

	pointsOfFunctionPlot[373].X = -6.27
	pointsOfFunctionPlot[373].Y = 39.312

	pointsOfFunctionPlot[374].X = -6.26
	pointsOfFunctionPlot[374].Y = 39.187

	pointsOfFunctionPlot[375].X = -6.25
	pointsOfFunctionPlot[375].Y = 39.062

	pointsOfFunctionPlot[376].X = -6.24
	pointsOfFunctionPlot[376].Y = 38.937

	pointsOfFunctionPlot[377].X = -6.23
	pointsOfFunctionPlot[377].Y = 38.812

	pointsOfFunctionPlot[378].X = -6.22
	pointsOfFunctionPlot[378].Y = 38.688

	pointsOfFunctionPlot[379].X = -6.21
	pointsOfFunctionPlot[379].Y = 38.564

	pointsOfFunctionPlot[380].X = -6.20
	pointsOfFunctionPlot[380].Y = 38.44

	pointsOfFunctionPlot[381].X = -6.19
	pointsOfFunctionPlot[381].Y = 38.316

	pointsOfFunctionPlot[382].X = -6.18
	pointsOfFunctionPlot[382].Y = 38.192

	pointsOfFunctionPlot[383].X = -6.17
	pointsOfFunctionPlot[383].Y = 38.068

	pointsOfFunctionPlot[384].X = -6.16
	pointsOfFunctionPlot[384].Y = 37.945

	pointsOfFunctionPlot[385].X = -6.15
	pointsOfFunctionPlot[385].Y = 37.822

	pointsOfFunctionPlot[386].X = -6.14
	pointsOfFunctionPlot[386].Y = 37.699

	pointsOfFunctionPlot[387].X = -6.13
	pointsOfFunctionPlot[387].Y = 37.576

	pointsOfFunctionPlot[388].X = -6.12
	pointsOfFunctionPlot[388].Y = 37.454

	pointsOfFunctionPlot[389].X = -6.11
	pointsOfFunctionPlot[389].Y = 37.332

	pointsOfFunctionPlot[390].X = -6.10
	pointsOfFunctionPlot[390].Y = 37.21

	pointsOfFunctionPlot[391].X = -6.09
	pointsOfFunctionPlot[391].Y = 37.088

	pointsOfFunctionPlot[392].X = -6.08
	pointsOfFunctionPlot[392].Y = 36.966

	pointsOfFunctionPlot[393].X = -6.07
	pointsOfFunctionPlot[393].Y = 36.844

	pointsOfFunctionPlot[394].X = -6.06
	pointsOfFunctionPlot[394].Y = 36.723

	pointsOfFunctionPlot[395].X = -6.05
	pointsOfFunctionPlot[395].Y = 36.602

	pointsOfFunctionPlot[396].X = -6.04
	pointsOfFunctionPlot[396].Y = 36.481

	pointsOfFunctionPlot[397].X = -6.03
	pointsOfFunctionPlot[397].Y = 36.36

	pointsOfFunctionPlot[398].X = -6.02
	pointsOfFunctionPlot[398].Y = 36.24

	pointsOfFunctionPlot[399].X = -6.01
	pointsOfFunctionPlot[399].Y = 36.12

	pointsOfFunctionPlot[400].X = -6.0
	pointsOfFunctionPlot[400].Y = 36.0

	pointsOfFunctionPlot[401].X = -5.99
	pointsOfFunctionPlot[401].Y = 35.88

	pointsOfFunctionPlot[402].X = -5.98
	pointsOfFunctionPlot[402].Y = 35.76

	pointsOfFunctionPlot[403].X = -5.97
	pointsOfFunctionPlot[403].Y = 35.64

	pointsOfFunctionPlot[404].X = -5.96
	pointsOfFunctionPlot[404].Y = 35.521

	pointsOfFunctionPlot[405].X = -5.95
	pointsOfFunctionPlot[405].Y = 35.402

	pointsOfFunctionPlot[406].X = -5.94
	pointsOfFunctionPlot[406].Y = 35.283

	pointsOfFunctionPlot[407].X = -5.93
	pointsOfFunctionPlot[407].Y = 35.164

	pointsOfFunctionPlot[408].X = -5.92
	pointsOfFunctionPlot[408].Y = 35.046

	pointsOfFunctionPlot[409].X = -5.91
	pointsOfFunctionPlot[409].Y = 34.928

	pointsOfFunctionPlot[410].X = -5.90
	pointsOfFunctionPlot[410].Y = 34.81

	pointsOfFunctionPlot[411].X = -5.89
	pointsOfFunctionPlot[411].Y = 34.692

	pointsOfFunctionPlot[412].X = -5.88
	pointsOfFunctionPlot[412].Y = 34.574

	pointsOfFunctionPlot[413].X = -5.87
	pointsOfFunctionPlot[413].Y = 34.456

	pointsOfFunctionPlot[414].X = -5.86
	pointsOfFunctionPlot[414].Y = 34.339

	pointsOfFunctionPlot[415].X = -5.85
	pointsOfFunctionPlot[415].Y = 34.222

	pointsOfFunctionPlot[416].X = -5.84
	pointsOfFunctionPlot[416].Y = 34.105

	pointsOfFunctionPlot[417].X = -5.83
	pointsOfFunctionPlot[417].Y = 33.988

	pointsOfFunctionPlot[418].X = -5.82
	pointsOfFunctionPlot[418].Y = 33.872

	pointsOfFunctionPlot[419].X = -5.81
	pointsOfFunctionPlot[419].Y = 33.756

	pointsOfFunctionPlot[420].X = -5.80
	pointsOfFunctionPlot[420].Y = 33.64

	pointsOfFunctionPlot[421].X = -5.79
	pointsOfFunctionPlot[421].Y = 33.524

	pointsOfFunctionPlot[422].X = -5.78
	pointsOfFunctionPlot[422].Y = 33.408

	pointsOfFunctionPlot[423].X = -5.77
	pointsOfFunctionPlot[423].Y = 33.292

	pointsOfFunctionPlot[424].X = -5.76
	pointsOfFunctionPlot[424].Y = 33.177

	pointsOfFunctionPlot[425].X = -5.75
	pointsOfFunctionPlot[425].Y = 33.062

	pointsOfFunctionPlot[426].X = -5.74
	pointsOfFunctionPlot[426].Y = 32.947

	pointsOfFunctionPlot[427].X = -5.73
	pointsOfFunctionPlot[427].Y = 32.832

	pointsOfFunctionPlot[428].X = -5.72
	pointsOfFunctionPlot[428].Y = 32.718

	pointsOfFunctionPlot[429].X = -5.71
	pointsOfFunctionPlot[429].Y = 32.604

	pointsOfFunctionPlot[430].X = -5.70
	pointsOfFunctionPlot[430].Y = 32.49

	pointsOfFunctionPlot[431].X = -5.69
	pointsOfFunctionPlot[431].Y = 32.376

	pointsOfFunctionPlot[432].X = -5.68
	pointsOfFunctionPlot[432].Y = 32.262

	pointsOfFunctionPlot[433].X = -5.67
	pointsOfFunctionPlot[433].Y = 32.148

	pointsOfFunctionPlot[434].X = -5.66
	pointsOfFunctionPlot[434].Y = 32.035

	pointsOfFunctionPlot[435].X = -5.65
	pointsOfFunctionPlot[435].Y = 31.922

	pointsOfFunctionPlot[436].X = -5.64
	pointsOfFunctionPlot[436].Y = 31.809

	pointsOfFunctionPlot[437].X = -5.63
	pointsOfFunctionPlot[437].Y = 31.696

	pointsOfFunctionPlot[438].X = -5.62
	pointsOfFunctionPlot[438].Y = 31.584

	pointsOfFunctionPlot[439].X = -5.61
	pointsOfFunctionPlot[439].Y = 31.472

	pointsOfFunctionPlot[440].X = -5.60
	pointsOfFunctionPlot[440].Y = 31.36

	pointsOfFunctionPlot[441].X = -5.59
	pointsOfFunctionPlot[441].Y = 31.248

	pointsOfFunctionPlot[442].X = -5.58
	pointsOfFunctionPlot[442].Y = 31.136

	pointsOfFunctionPlot[443].X = -5.57
	pointsOfFunctionPlot[443].Y = 31.024

	pointsOfFunctionPlot[444].X = -5.56
	pointsOfFunctionPlot[444].Y = 30.913

	pointsOfFunctionPlot[445].X = -5.55
	pointsOfFunctionPlot[445].Y = 30.802

	pointsOfFunctionPlot[446].X = -5.54
	pointsOfFunctionPlot[446].Y = 30.691

	pointsOfFunctionPlot[447].X = -5.53
	pointsOfFunctionPlot[447].Y = 30.58

	pointsOfFunctionPlot[448].X = -5.52
	pointsOfFunctionPlot[448].Y = 30.47

	pointsOfFunctionPlot[449].X = -5.51
	pointsOfFunctionPlot[449].Y = 30.36

	pointsOfFunctionPlot[450].X = -5.50
	pointsOfFunctionPlot[450].Y = 30.25

	pointsOfFunctionPlot[451].X = -5.49
	pointsOfFunctionPlot[451].Y = 30.14

	pointsOfFunctionPlot[452].X = -5.48
	pointsOfFunctionPlot[452].Y = 30.03

	pointsOfFunctionPlot[453].X = -5.47
	pointsOfFunctionPlot[453].Y = 29.92

	pointsOfFunctionPlot[454].X = -5.46
	pointsOfFunctionPlot[454].Y = 29.811

	pointsOfFunctionPlot[455].X = -5.45
	pointsOfFunctionPlot[455].Y = 29.702

	pointsOfFunctionPlot[456].X = -5.44
	pointsOfFunctionPlot[456].Y = 29.593

	pointsOfFunctionPlot[457].X = -5.43
	pointsOfFunctionPlot[457].Y = 29.484

	pointsOfFunctionPlot[458].X = -5.42
	pointsOfFunctionPlot[458].Y = 29.376

	pointsOfFunctionPlot[459].X = -5.41
	pointsOfFunctionPlot[459].Y = 29.268

	pointsOfFunctionPlot[460].X = -5.40
	pointsOfFunctionPlot[460].Y = 29.16

	pointsOfFunctionPlot[461].X = -5.39
	pointsOfFunctionPlot[461].Y = 29.052

	pointsOfFunctionPlot[462].X = -5.38
	pointsOfFunctionPlot[462].Y = 28.944

	pointsOfFunctionPlot[463].X = -5.37
	pointsOfFunctionPlot[463].Y = 28.836

	pointsOfFunctionPlot[464].X = -5.36
	pointsOfFunctionPlot[464].Y = 28.729

	pointsOfFunctionPlot[465].X = -5.35
	pointsOfFunctionPlot[465].Y = 28.622

	pointsOfFunctionPlot[466].X = -5.34
	pointsOfFunctionPlot[466].Y = 28.515

	pointsOfFunctionPlot[467].X = -5.33
	pointsOfFunctionPlot[467].Y = 28.408

	pointsOfFunctionPlot[468].X = -5.32
	pointsOfFunctionPlot[468].Y = 28.302

	pointsOfFunctionPlot[469].X = -5.31
	pointsOfFunctionPlot[469].Y = 28.196

	pointsOfFunctionPlot[470].X = -5.30
	pointsOfFunctionPlot[470].Y = 28.09

	pointsOfFunctionPlot[471].X = -5.29
	pointsOfFunctionPlot[471].Y = 27.984

	pointsOfFunctionPlot[472].X = -5.28
	pointsOfFunctionPlot[472].Y = 27.878

	pointsOfFunctionPlot[473].X = -5.27
	pointsOfFunctionPlot[473].Y = 27.772

	pointsOfFunctionPlot[474].X = -5.26
	pointsOfFunctionPlot[474].Y = 27.667

	pointsOfFunctionPlot[475].X = -5.25
	pointsOfFunctionPlot[475].Y = 27.562

	pointsOfFunctionPlot[476].X = -5.24
	pointsOfFunctionPlot[476].Y = 27.457

	pointsOfFunctionPlot[477].X = -5.23
	pointsOfFunctionPlot[477].Y = 27.352

	pointsOfFunctionPlot[478].X = -5.22
	pointsOfFunctionPlot[478].Y = 27.248

	pointsOfFunctionPlot[479].X = -5.21
	pointsOfFunctionPlot[479].Y = 27.144

	pointsOfFunctionPlot[480].X = -5.20
	pointsOfFunctionPlot[480].Y = 27.04

	pointsOfFunctionPlot[481].X = -5.19
	pointsOfFunctionPlot[481].Y = 26.936

	pointsOfFunctionPlot[482].X = -5.18
	pointsOfFunctionPlot[482].Y = 26.832

	pointsOfFunctionPlot[483].X = -5.17
	pointsOfFunctionPlot[483].Y = 26.728

	pointsOfFunctionPlot[484].X = -5.16
	pointsOfFunctionPlot[484].Y = 26.625

	pointsOfFunctionPlot[485].X = -5.15
	pointsOfFunctionPlot[485].Y = 26.522

	pointsOfFunctionPlot[486].X = -5.14
	pointsOfFunctionPlot[486].Y = 26.419

	pointsOfFunctionPlot[487].X = -5.13
	pointsOfFunctionPlot[487].Y = 26.316

	pointsOfFunctionPlot[488].X = -5.12
	pointsOfFunctionPlot[488].Y = 26.214

	pointsOfFunctionPlot[489].X = -5.11
	pointsOfFunctionPlot[489].Y = 26.112

	pointsOfFunctionPlot[490].X = -5.10
	pointsOfFunctionPlot[490].Y = 26.01

	pointsOfFunctionPlot[491].X = -5.09
	pointsOfFunctionPlot[491].Y = 25.908

	pointsOfFunctionPlot[492].X = -5.08
	pointsOfFunctionPlot[492].Y = 25.806

	pointsOfFunctionPlot[493].X = -5.07
	pointsOfFunctionPlot[493].Y = 25.704

	pointsOfFunctionPlot[494].X = -5.06
	pointsOfFunctionPlot[494].Y = 25.603

	pointsOfFunctionPlot[495].X = -5.05
	pointsOfFunctionPlot[495].Y = 25.502

	pointsOfFunctionPlot[496].X = -5.04
	pointsOfFunctionPlot[496].Y = 25.401

	pointsOfFunctionPlot[497].X = -5.03
	pointsOfFunctionPlot[497].Y = 25.3

	pointsOfFunctionPlot[498].X = -5.02
	pointsOfFunctionPlot[498].Y = 25.2

	pointsOfFunctionPlot[499].X = -5.01
	pointsOfFunctionPlot[499].Y = 25.1

	pointsOfFunctionPlot[500].X = -5.0
	pointsOfFunctionPlot[500].Y = 25.0

	pointsOfFunctionPlot[501].X = -4.99
	pointsOfFunctionPlot[501].Y = 24.9

	pointsOfFunctionPlot[502].X = -4.98
	pointsOfFunctionPlot[502].Y = 24.8

	pointsOfFunctionPlot[503].X = -4.97
	pointsOfFunctionPlot[503].Y = 24.7

	pointsOfFunctionPlot[504].X = -4.96
	pointsOfFunctionPlot[504].Y = 24.601

	pointsOfFunctionPlot[505].X = -4.95
	pointsOfFunctionPlot[505].Y = 24.502

	pointsOfFunctionPlot[506].X = -4.94
	pointsOfFunctionPlot[506].Y = 24.403

	pointsOfFunctionPlot[507].X = -4.93
	pointsOfFunctionPlot[507].Y = 24.304

	pointsOfFunctionPlot[508].X = -4.92
	pointsOfFunctionPlot[508].Y = 24.206

	pointsOfFunctionPlot[509].X = -4.91
	pointsOfFunctionPlot[509].Y = 24.108

	pointsOfFunctionPlot[510].X = -4.90
	pointsOfFunctionPlot[510].Y = 24.01

	pointsOfFunctionPlot[511].X = -4.89
	pointsOfFunctionPlot[511].Y = 23.912

	pointsOfFunctionPlot[512].X = -4.88
	pointsOfFunctionPlot[512].Y = 23.814

	pointsOfFunctionPlot[513].X = -4.87
	pointsOfFunctionPlot[513].Y = 23.716

	pointsOfFunctionPlot[514].X = -4.86
	pointsOfFunctionPlot[514].Y = 23.619

	pointsOfFunctionPlot[515].X = -4.85
	pointsOfFunctionPlot[515].Y = 23.522

	pointsOfFunctionPlot[516].X = -4.84
	pointsOfFunctionPlot[516].Y = 23.425

	pointsOfFunctionPlot[517].X = -4.83
	pointsOfFunctionPlot[517].Y = 23.328

	pointsOfFunctionPlot[518].X = -4.82
	pointsOfFunctionPlot[518].Y = 23.232

	pointsOfFunctionPlot[519].X = -4.81
	pointsOfFunctionPlot[519].Y = 23.136

	pointsOfFunctionPlot[520].X = -4.80
	pointsOfFunctionPlot[520].Y = 23.04

	pointsOfFunctionPlot[521].X = -4.79
	pointsOfFunctionPlot[521].Y = 22.944

	pointsOfFunctionPlot[522].X = -4.78
	pointsOfFunctionPlot[522].Y = 22.848

	pointsOfFunctionPlot[523].X = -4.77
	pointsOfFunctionPlot[523].Y = 22.752

	pointsOfFunctionPlot[524].X = -4.76
	pointsOfFunctionPlot[524].Y = 22.657

	pointsOfFunctionPlot[525].X = -4.75
	pointsOfFunctionPlot[525].Y = 22.562

	pointsOfFunctionPlot[526].X = -4.74
	pointsOfFunctionPlot[526].Y = 22.467

	pointsOfFunctionPlot[527].X = -4.73
	pointsOfFunctionPlot[527].Y = 22.372

	pointsOfFunctionPlot[528].X = -4.72
	pointsOfFunctionPlot[528].Y = 22.278

	pointsOfFunctionPlot[529].X = -4.71
	pointsOfFunctionPlot[529].Y = 22.184

	pointsOfFunctionPlot[530].X = -4.70
	pointsOfFunctionPlot[530].Y = 22.09

	pointsOfFunctionPlot[531].X = -4.69
	pointsOfFunctionPlot[531].Y = 21.996

	pointsOfFunctionPlot[532].X = -4.68
	pointsOfFunctionPlot[532].Y = 21.902

	pointsOfFunctionPlot[533].X = -4.67
	pointsOfFunctionPlot[533].Y = 21.808

	pointsOfFunctionPlot[534].X = -4.66
	pointsOfFunctionPlot[534].Y = 21.715

	pointsOfFunctionPlot[535].X = -4.65
	pointsOfFunctionPlot[535].Y = 21.622

	pointsOfFunctionPlot[536].X = -4.64
	pointsOfFunctionPlot[536].Y = 21.529

	pointsOfFunctionPlot[537].X = -4.63
	pointsOfFunctionPlot[537].Y = 21.436

	pointsOfFunctionPlot[538].X = -4.62
	pointsOfFunctionPlot[538].Y = 21.344

	pointsOfFunctionPlot[539].X = -4.61
	pointsOfFunctionPlot[539].Y = 21.252

	pointsOfFunctionPlot[540].X = -4.60
	pointsOfFunctionPlot[540].Y = 21.16

	pointsOfFunctionPlot[541].X = -4.59
	pointsOfFunctionPlot[541].Y = 21.068

	pointsOfFunctionPlot[542].X = -4.58
	pointsOfFunctionPlot[542].Y = 20.976

	pointsOfFunctionPlot[543].X = -4.57
	pointsOfFunctionPlot[543].Y = 20.884

	pointsOfFunctionPlot[544].X = -4.56
	pointsOfFunctionPlot[544].Y = 20.793

	pointsOfFunctionPlot[545].X = -4.55
	pointsOfFunctionPlot[545].Y = 20.702

	pointsOfFunctionPlot[546].X = -4.54
	pointsOfFunctionPlot[546].Y = 20.611

	pointsOfFunctionPlot[547].X = -4.53
	pointsOfFunctionPlot[547].Y = 20.52

	pointsOfFunctionPlot[548].X = -4.52
	pointsOfFunctionPlot[548].Y = 20.43

	pointsOfFunctionPlot[549].X = -4.51
	pointsOfFunctionPlot[549].Y = 20.34

	pointsOfFunctionPlot[550].X = -4.50
	pointsOfFunctionPlot[550].Y = 20.25

	pointsOfFunctionPlot[551].X = -4.49
	pointsOfFunctionPlot[551].Y = 20.16

	pointsOfFunctionPlot[552].X = -4.48
	pointsOfFunctionPlot[552].Y = 20.07

	pointsOfFunctionPlot[553].X = -4.47
	pointsOfFunctionPlot[553].Y = 19.98

	pointsOfFunctionPlot[554].X = -4.46
	pointsOfFunctionPlot[554].Y = 19.891

	pointsOfFunctionPlot[555].X = -4.45
	pointsOfFunctionPlot[555].Y = 19.802

	pointsOfFunctionPlot[556].X = -4.44
	pointsOfFunctionPlot[556].Y = 19.713

	pointsOfFunctionPlot[557].X = -4.43
	pointsOfFunctionPlot[557].Y = 19.624

	pointsOfFunctionPlot[558].X = -4.42
	pointsOfFunctionPlot[558].Y = 19.536

	pointsOfFunctionPlot[559].X = -4.41
	pointsOfFunctionPlot[559].Y = 19.448

	pointsOfFunctionPlot[560].X = -4.40
	pointsOfFunctionPlot[560].Y = 19.36

	pointsOfFunctionPlot[561].X = -4.39
	pointsOfFunctionPlot[561].Y = 19.272

	pointsOfFunctionPlot[562].X = -4.38
	pointsOfFunctionPlot[562].Y = 19.184

	pointsOfFunctionPlot[563].X = -4.37
	pointsOfFunctionPlot[563].Y = 19.096

	pointsOfFunctionPlot[564].X = -4.36
	pointsOfFunctionPlot[564].Y = 19.009

	pointsOfFunctionPlot[565].X = -4.35
	pointsOfFunctionPlot[565].Y = 18.922

	pointsOfFunctionPlot[566].X = -4.34
	pointsOfFunctionPlot[566].Y = 18.835

	pointsOfFunctionPlot[567].X = -4.33
	pointsOfFunctionPlot[567].Y = 18.748

	pointsOfFunctionPlot[568].X = -4.32
	pointsOfFunctionPlot[568].Y = 18.662

	pointsOfFunctionPlot[569].X = -4.31
	pointsOfFunctionPlot[569].Y = 18.576

	pointsOfFunctionPlot[570].X = -4.30
	pointsOfFunctionPlot[570].Y = 18.49

	pointsOfFunctionPlot[571].X = -4.29
	pointsOfFunctionPlot[571].Y = 18.404

	pointsOfFunctionPlot[572].X = -4.28
	pointsOfFunctionPlot[572].Y = 18.318

	pointsOfFunctionPlot[573].X = -4.27
	pointsOfFunctionPlot[573].Y = 18.232

	pointsOfFunctionPlot[574].X = -4.26
	pointsOfFunctionPlot[574].Y = 18.147

	pointsOfFunctionPlot[575].X = -4.25
	pointsOfFunctionPlot[575].Y = 18.062

	pointsOfFunctionPlot[576].X = -4.24
	pointsOfFunctionPlot[576].Y = 17.977

	pointsOfFunctionPlot[577].X = -4.23
	pointsOfFunctionPlot[577].Y = 17.892

	pointsOfFunctionPlot[578].X = -4.22
	pointsOfFunctionPlot[578].Y = 17.808

	pointsOfFunctionPlot[579].X = -4.21
	pointsOfFunctionPlot[579].Y = 17.724

	pointsOfFunctionPlot[580].X = -4.20
	pointsOfFunctionPlot[580].Y = 17.64

	pointsOfFunctionPlot[581].X = -4.19
	pointsOfFunctionPlot[581].Y = 17.556

	pointsOfFunctionPlot[582].X = -4.18
	pointsOfFunctionPlot[582].Y = 17.472

	pointsOfFunctionPlot[583].X = -4.17
	pointsOfFunctionPlot[583].Y = 17.388

	pointsOfFunctionPlot[584].X = -4.16
	pointsOfFunctionPlot[584].Y = 17.305

	pointsOfFunctionPlot[585].X = -4.15
	pointsOfFunctionPlot[585].Y = 17.222

	pointsOfFunctionPlot[586].X = -4.14
	pointsOfFunctionPlot[586].Y = 17.139

	pointsOfFunctionPlot[587].X = -4.13
	pointsOfFunctionPlot[587].Y = 17.056

	pointsOfFunctionPlot[588].X = -4.12
	pointsOfFunctionPlot[588].Y = 16.974

	pointsOfFunctionPlot[589].X = -4.11
	pointsOfFunctionPlot[589].Y = 16.892

	pointsOfFunctionPlot[590].X = -4.10
	pointsOfFunctionPlot[590].Y = 16.81

	pointsOfFunctionPlot[591].X = -4.09
	pointsOfFunctionPlot[591].Y = 16.728

	pointsOfFunctionPlot[592].X = -4.08
	pointsOfFunctionPlot[592].Y = 16.646

	pointsOfFunctionPlot[593].X = -4.07
	pointsOfFunctionPlot[593].Y = 16.564

	pointsOfFunctionPlot[594].X = -4.06
	pointsOfFunctionPlot[594].Y = 16.483

	pointsOfFunctionPlot[595].X = -4.05
	pointsOfFunctionPlot[595].Y = 16.402

	pointsOfFunctionPlot[596].X = -4.04
	pointsOfFunctionPlot[596].Y = 16.321

	pointsOfFunctionPlot[597].X = -4.03
	pointsOfFunctionPlot[597].Y = 16.24

	pointsOfFunctionPlot[598].X = -4.02
	pointsOfFunctionPlot[598].Y = 16.16

	pointsOfFunctionPlot[599].X = -4.01
	pointsOfFunctionPlot[599].Y = 16.08

	pointsOfFunctionPlot[600].X = -4.0
	pointsOfFunctionPlot[600].Y = 16.0

	pointsOfFunctionPlot[601].X = -3.99
	pointsOfFunctionPlot[601].Y = 15.92

	pointsOfFunctionPlot[602].X = -3.98
	pointsOfFunctionPlot[602].Y = 15.84

	pointsOfFunctionPlot[603].X = -3.97
	pointsOfFunctionPlot[603].Y = 15.76

	pointsOfFunctionPlot[604].X = -3.96
	pointsOfFunctionPlot[604].Y = 15.681

	pointsOfFunctionPlot[605].X = -3.95
	pointsOfFunctionPlot[605].Y = 15.602

	pointsOfFunctionPlot[606].X = -3.94
	pointsOfFunctionPlot[606].Y = 15.523

	pointsOfFunctionPlot[607].X = -3.93
	pointsOfFunctionPlot[607].Y = 15.444

	pointsOfFunctionPlot[608].X = -3.92
	pointsOfFunctionPlot[608].Y = 15.366

	pointsOfFunctionPlot[609].X = -3.91
	pointsOfFunctionPlot[609].Y = 15.288

	pointsOfFunctionPlot[610].X = -3.90
	pointsOfFunctionPlot[610].Y = 15.21

	pointsOfFunctionPlot[611].X = -3.89
	pointsOfFunctionPlot[611].Y = 15.132

	pointsOfFunctionPlot[612].X = -3.88
	pointsOfFunctionPlot[612].Y = 15.054

	pointsOfFunctionPlot[613].X = -3.87
	pointsOfFunctionPlot[613].Y = 14.976

	pointsOfFunctionPlot[614].X = -3.86
	pointsOfFunctionPlot[614].Y = 14.899

	pointsOfFunctionPlot[615].X = -3.85
	pointsOfFunctionPlot[615].Y = 14.822

	pointsOfFunctionPlot[616].X = -3.84
	pointsOfFunctionPlot[616].Y = 14.745

	pointsOfFunctionPlot[617].X = -3.83
	pointsOfFunctionPlot[617].Y = 14.668

	pointsOfFunctionPlot[618].X = -3.82
	pointsOfFunctionPlot[618].Y = 14.592

	pointsOfFunctionPlot[619].X = -3.81
	pointsOfFunctionPlot[619].Y = 14.516

	pointsOfFunctionPlot[620].X = -3.80
	pointsOfFunctionPlot[620].Y = 14.44

	pointsOfFunctionPlot[621].X = -3.79
	pointsOfFunctionPlot[621].Y = 14.364

	pointsOfFunctionPlot[622].X = -3.78
	pointsOfFunctionPlot[622].Y = 14.288

	pointsOfFunctionPlot[623].X = -3.77
	pointsOfFunctionPlot[623].Y = 14.212

	pointsOfFunctionPlot[624].X = -3.76
	pointsOfFunctionPlot[624].Y = 14.137

	pointsOfFunctionPlot[625].X = -3.75
	pointsOfFunctionPlot[625].Y = 14.062

	pointsOfFunctionPlot[626].X = -3.74
	pointsOfFunctionPlot[626].Y = 13.987

	pointsOfFunctionPlot[627].X = -3.73
	pointsOfFunctionPlot[627].Y = 13.912

	pointsOfFunctionPlot[628].X = -3.72
	pointsOfFunctionPlot[628].Y = 13.838

	pointsOfFunctionPlot[629].X = -3.71
	pointsOfFunctionPlot[629].Y = 13.764

	pointsOfFunctionPlot[630].X = -3.70
	pointsOfFunctionPlot[630].Y = 13.69

	pointsOfFunctionPlot[631].X = -3.69
	pointsOfFunctionPlot[631].Y = 13.616

	pointsOfFunctionPlot[632].X = -3.68
	pointsOfFunctionPlot[632].Y = 13.542

	pointsOfFunctionPlot[633].X = -3.67
	pointsOfFunctionPlot[633].Y = 13.468

	pointsOfFunctionPlot[634].X = -3.66
	pointsOfFunctionPlot[634].Y = 13.395

	pointsOfFunctionPlot[635].X = -3.65
	pointsOfFunctionPlot[635].Y = 13.322

	pointsOfFunctionPlot[636].X = -3.64
	pointsOfFunctionPlot[636].Y = 13.249

	pointsOfFunctionPlot[637].X = -3.63
	pointsOfFunctionPlot[637].Y = 13.176

	pointsOfFunctionPlot[638].X = -3.62
	pointsOfFunctionPlot[638].Y = 13.104

	pointsOfFunctionPlot[639].X = -3.61
	pointsOfFunctionPlot[639].Y = 13.032

	pointsOfFunctionPlot[640].X = -3.60
	pointsOfFunctionPlot[640].Y = 12.96

	pointsOfFunctionPlot[641].X = -3.59
	pointsOfFunctionPlot[641].Y = 12.888

	pointsOfFunctionPlot[642].X = -3.58
	pointsOfFunctionPlot[642].Y = 12.816

	pointsOfFunctionPlot[643].X = -3.57
	pointsOfFunctionPlot[643].Y = 12.744

	pointsOfFunctionPlot[644].X = -3.56
	pointsOfFunctionPlot[644].Y = 12.673

	pointsOfFunctionPlot[645].X = -3.55
	pointsOfFunctionPlot[645].Y = 12.602

	pointsOfFunctionPlot[646].X = -3.54
	pointsOfFunctionPlot[646].Y = 12.531

	pointsOfFunctionPlot[647].X = -3.53
	pointsOfFunctionPlot[647].Y = 12.46

	pointsOfFunctionPlot[648].X = -3.52
	pointsOfFunctionPlot[648].Y = 12.39

	pointsOfFunctionPlot[649].X = -3.51
	pointsOfFunctionPlot[649].Y = 12.32

	pointsOfFunctionPlot[650].X = -3.50
	pointsOfFunctionPlot[650].Y = 12.25

	pointsOfFunctionPlot[651].X = -3.49
	pointsOfFunctionPlot[651].Y = 12.18

	pointsOfFunctionPlot[652].X = -3.48
	pointsOfFunctionPlot[652].Y = 12.11

	pointsOfFunctionPlot[653].X = -3.47
	pointsOfFunctionPlot[653].Y = 12.04

	pointsOfFunctionPlot[654].X = -3.46
	pointsOfFunctionPlot[654].Y = 11.971

	pointsOfFunctionPlot[655].X = -3.45
	pointsOfFunctionPlot[655].Y = 11.902

	pointsOfFunctionPlot[656].X = -3.44
	pointsOfFunctionPlot[656].Y = 11.833

	pointsOfFunctionPlot[657].X = -3.43
	pointsOfFunctionPlot[657].Y = 11.764

	pointsOfFunctionPlot[658].X = -3.42
	pointsOfFunctionPlot[658].Y = 11.696

	pointsOfFunctionPlot[659].X = -3.41
	pointsOfFunctionPlot[659].Y = 11.628

	pointsOfFunctionPlot[660].X = -3.40
	pointsOfFunctionPlot[660].Y = 11.56

	pointsOfFunctionPlot[661].X = -3.39
	pointsOfFunctionPlot[661].Y = 11.492

	pointsOfFunctionPlot[662].X = -3.38
	pointsOfFunctionPlot[662].Y = 11.424

	pointsOfFunctionPlot[663].X = -3.37
	pointsOfFunctionPlot[663].Y = 11.356

	pointsOfFunctionPlot[664].X = -3.36
	pointsOfFunctionPlot[664].Y = 11.289

	pointsOfFunctionPlot[665].X = -3.35
	pointsOfFunctionPlot[665].Y = 11.222

	pointsOfFunctionPlot[666].X = -3.34
	pointsOfFunctionPlot[666].Y = 11.155

	pointsOfFunctionPlot[667].X = -3.33
	pointsOfFunctionPlot[667].Y = 11.088

	pointsOfFunctionPlot[668].X = -3.32
	pointsOfFunctionPlot[668].Y = 11.022

	pointsOfFunctionPlot[669].X = -3.31
	pointsOfFunctionPlot[669].Y = 10.956

	pointsOfFunctionPlot[670].X = -3.30
	pointsOfFunctionPlot[670].Y = 10.89

	pointsOfFunctionPlot[671].X = -3.29
	pointsOfFunctionPlot[671].Y = 10.824

	pointsOfFunctionPlot[672].X = -3.28
	pointsOfFunctionPlot[672].Y = 10.758

	pointsOfFunctionPlot[673].X = -3.27
	pointsOfFunctionPlot[673].Y = 10.692

	pointsOfFunctionPlot[674].X = -3.26
	pointsOfFunctionPlot[674].Y = 10.627

	pointsOfFunctionPlot[675].X = -3.25
	pointsOfFunctionPlot[675].Y = 10.562

	pointsOfFunctionPlot[676].X = -3.24
	pointsOfFunctionPlot[676].Y = 10.497

	pointsOfFunctionPlot[677].X = -3.23
	pointsOfFunctionPlot[677].Y = 10.432

	pointsOfFunctionPlot[678].X = -3.22
	pointsOfFunctionPlot[678].Y = 10.368

	pointsOfFunctionPlot[679].X = -3.21
	pointsOfFunctionPlot[679].Y = 10.304

	pointsOfFunctionPlot[680].X = -3.20
	pointsOfFunctionPlot[680].Y = 10.24

	pointsOfFunctionPlot[681].X = -3.19
	pointsOfFunctionPlot[681].Y = 10.176

	pointsOfFunctionPlot[682].X = -3.18
	pointsOfFunctionPlot[682].Y = 10.112

	pointsOfFunctionPlot[683].X = -3.17
	pointsOfFunctionPlot[683].Y = 10.048

	pointsOfFunctionPlot[684].X = -3.16
	pointsOfFunctionPlot[684].Y = 9.985

	pointsOfFunctionPlot[685].X = -3.15
	pointsOfFunctionPlot[685].Y = 9.922

	pointsOfFunctionPlot[686].X = -3.14
	pointsOfFunctionPlot[686].Y = 9.859

	pointsOfFunctionPlot[687].X = -3.13
	pointsOfFunctionPlot[687].Y = 9.796

	pointsOfFunctionPlot[688].X = -3.12
	pointsOfFunctionPlot[688].Y = 9.734

	pointsOfFunctionPlot[689].X = -3.11
	pointsOfFunctionPlot[689].Y = 9.672

	pointsOfFunctionPlot[690].X = -3.10
	pointsOfFunctionPlot[690].Y = 9.61

	pointsOfFunctionPlot[691].X = -3.09
	pointsOfFunctionPlot[691].Y = 9.548

	pointsOfFunctionPlot[692].X = -3.08
	pointsOfFunctionPlot[692].Y = 9.486

	pointsOfFunctionPlot[693].X = -3.07
	pointsOfFunctionPlot[693].Y = 9.424

	pointsOfFunctionPlot[694].X = -3.06
	pointsOfFunctionPlot[694].Y = 9.363

	pointsOfFunctionPlot[695].X = -3.05
	pointsOfFunctionPlot[695].Y = 9.302

	pointsOfFunctionPlot[696].X = -3.04
	pointsOfFunctionPlot[696].Y = 9.241

	pointsOfFunctionPlot[697].X = -3.03
	pointsOfFunctionPlot[697].Y = 9.18

	pointsOfFunctionPlot[698].X = -3.02
	pointsOfFunctionPlot[698].Y = 9.12

	pointsOfFunctionPlot[699].X = -3.01
	pointsOfFunctionPlot[699].Y = 9.06

	pointsOfFunctionPlot[700].X = -3.0
	pointsOfFunctionPlot[700].Y = 9.0

	pointsOfFunctionPlot[701].X = -2.99
	pointsOfFunctionPlot[701].Y = 8.94

	pointsOfFunctionPlot[702].X = -2.98
	pointsOfFunctionPlot[702].Y = 8.88

	pointsOfFunctionPlot[703].X = -2.97
	pointsOfFunctionPlot[703].Y = 8.82

	pointsOfFunctionPlot[704].X = -2.96
	pointsOfFunctionPlot[704].Y = 8.761

	pointsOfFunctionPlot[705].X = -2.95
	pointsOfFunctionPlot[705].Y = 8.702

	pointsOfFunctionPlot[706].X = -2.94
	pointsOfFunctionPlot[706].Y = 8.643

	pointsOfFunctionPlot[707].X = -2.93
	pointsOfFunctionPlot[707].Y = 8.584

	pointsOfFunctionPlot[708].X = -2.92
	pointsOfFunctionPlot[708].Y = 8.526

	pointsOfFunctionPlot[709].X = -2.91
	pointsOfFunctionPlot[709].Y = 8.468

	pointsOfFunctionPlot[710].X = -2.90
	pointsOfFunctionPlot[710].Y = 8.41

	pointsOfFunctionPlot[711].X = -2.89
	pointsOfFunctionPlot[711].Y = 8.352

	pointsOfFunctionPlot[712].X = -2.88
	pointsOfFunctionPlot[712].Y = 8.294

	pointsOfFunctionPlot[713].X = -2.87
	pointsOfFunctionPlot[713].Y = 8.236

	pointsOfFunctionPlot[714].X = -2.86
	pointsOfFunctionPlot[714].Y = 8.179

	pointsOfFunctionPlot[715].X = -2.85
	pointsOfFunctionPlot[715].Y = 8.122

	pointsOfFunctionPlot[716].X = -2.84
	pointsOfFunctionPlot[716].Y = 8.065

	pointsOfFunctionPlot[717].X = -2.83
	pointsOfFunctionPlot[717].Y = 8.008

	pointsOfFunctionPlot[718].X = -2.82
	pointsOfFunctionPlot[718].Y = 7.952

	pointsOfFunctionPlot[719].X = -2.81
	pointsOfFunctionPlot[719].Y = 7.896

	pointsOfFunctionPlot[720].X = -2.80
	pointsOfFunctionPlot[720].Y = 7.84

	pointsOfFunctionPlot[721].X = -2.79
	pointsOfFunctionPlot[721].Y = 7.784

	pointsOfFunctionPlot[722].X = -2.78
	pointsOfFunctionPlot[722].Y = 7.728

	pointsOfFunctionPlot[723].X = -2.77
	pointsOfFunctionPlot[723].Y = 7.672

	pointsOfFunctionPlot[724].X = -2.76
	pointsOfFunctionPlot[724].Y = 7.617

	pointsOfFunctionPlot[725].X = -2.75
	pointsOfFunctionPlot[725].Y = 7.562

	pointsOfFunctionPlot[726].X = -2.74
	pointsOfFunctionPlot[726].Y = 7.507

	pointsOfFunctionPlot[727].X = -2.73
	pointsOfFunctionPlot[727].Y = 7.452

	pointsOfFunctionPlot[728].X = -2.72
	pointsOfFunctionPlot[728].Y = 7.398

	pointsOfFunctionPlot[729].X = -2.71
	pointsOfFunctionPlot[729].Y = 7.344

	pointsOfFunctionPlot[730].X = -2.70
	pointsOfFunctionPlot[730].Y = 7.29

	pointsOfFunctionPlot[731].X = -2.69
	pointsOfFunctionPlot[731].Y = 7.236

	pointsOfFunctionPlot[732].X = -2.68
	pointsOfFunctionPlot[732].Y = 7.182

	pointsOfFunctionPlot[733].X = -2.67
	pointsOfFunctionPlot[733].Y = 7.128

	pointsOfFunctionPlot[734].X = -2.66
	pointsOfFunctionPlot[734].Y = 7.075

	pointsOfFunctionPlot[735].X = -2.65
	pointsOfFunctionPlot[735].Y = 7.022

	pointsOfFunctionPlot[736].X = -2.64
	pointsOfFunctionPlot[736].Y = 6.969

	pointsOfFunctionPlot[737].X = -2.63
	pointsOfFunctionPlot[737].Y = 6.916

	pointsOfFunctionPlot[738].X = -2.62
	pointsOfFunctionPlot[738].Y = 6.864

	pointsOfFunctionPlot[739].X = -2.61
	pointsOfFunctionPlot[739].Y = 6.812

	pointsOfFunctionPlot[740].X = -2.60
	pointsOfFunctionPlot[740].Y = 6.76

	pointsOfFunctionPlot[741].X = -2.59
	pointsOfFunctionPlot[741].Y = 6.708

	pointsOfFunctionPlot[742].X = -2.58
	pointsOfFunctionPlot[742].Y = 6.656

	pointsOfFunctionPlot[743].X = -2.57
	pointsOfFunctionPlot[743].Y = 6.604

	pointsOfFunctionPlot[744].X = -2.56
	pointsOfFunctionPlot[744].Y = 6.553

	pointsOfFunctionPlot[745].X = -2.55
	pointsOfFunctionPlot[745].Y = 6.502

	pointsOfFunctionPlot[746].X = -2.54
	pointsOfFunctionPlot[746].Y = 6.451

	pointsOfFunctionPlot[747].X = -2.53
	pointsOfFunctionPlot[747].Y = 6.4

	pointsOfFunctionPlot[748].X = -2.52
	pointsOfFunctionPlot[748].Y = 6.35

	pointsOfFunctionPlot[749].X = -2.51
	pointsOfFunctionPlot[749].Y = 6.3

	pointsOfFunctionPlot[750].X = -2.50
	pointsOfFunctionPlot[750].Y = 6.25

	pointsOfFunctionPlot[751].X = -2.49
	pointsOfFunctionPlot[751].Y = 6.2

	pointsOfFunctionPlot[752].X = -2.48
	pointsOfFunctionPlot[752].Y = 6.15

	pointsOfFunctionPlot[753].X = -2.47
	pointsOfFunctionPlot[753].Y = 6.1

	pointsOfFunctionPlot[754].X = -2.46
	pointsOfFunctionPlot[754].Y = 6.051

	pointsOfFunctionPlot[755].X = -2.45
	pointsOfFunctionPlot[755].Y = 6.002

	pointsOfFunctionPlot[756].X = -2.44
	pointsOfFunctionPlot[756].Y = 5.953

	pointsOfFunctionPlot[757].X = -2.43
	pointsOfFunctionPlot[757].Y = 5.904

	pointsOfFunctionPlot[758].X = -2.42
	pointsOfFunctionPlot[758].Y = 5.856

	pointsOfFunctionPlot[759].X = -2.41
	pointsOfFunctionPlot[759].Y = 5.808

	pointsOfFunctionPlot[760].X = -2.40
	pointsOfFunctionPlot[760].Y = 5.76

	pointsOfFunctionPlot[761].X = -2.39
	pointsOfFunctionPlot[761].Y = 5.712

	pointsOfFunctionPlot[762].X = -2.38
	pointsOfFunctionPlot[762].Y = 5.664

	pointsOfFunctionPlot[763].X = -2.37
	pointsOfFunctionPlot[763].Y = 5.616

	pointsOfFunctionPlot[764].X = -2.36
	pointsOfFunctionPlot[764].Y = 5.569

	pointsOfFunctionPlot[765].X = -2.35
	pointsOfFunctionPlot[765].Y = 5.522

	pointsOfFunctionPlot[766].X = -2.34
	pointsOfFunctionPlot[766].Y = 5.475

	pointsOfFunctionPlot[767].X = -2.33
	pointsOfFunctionPlot[767].Y = 5.428

	pointsOfFunctionPlot[768].X = -2.32
	pointsOfFunctionPlot[768].Y = 5.382

	pointsOfFunctionPlot[769].X = -2.31
	pointsOfFunctionPlot[769].Y = 5.336

	pointsOfFunctionPlot[770].X = -2.30
	pointsOfFunctionPlot[770].Y = 5.29

	pointsOfFunctionPlot[771].X = -2.29
	pointsOfFunctionPlot[771].Y = 5.244

	pointsOfFunctionPlot[772].X = -2.28
	pointsOfFunctionPlot[772].Y = 5.198

	pointsOfFunctionPlot[773].X = -2.27
	pointsOfFunctionPlot[773].Y = 5.152

	pointsOfFunctionPlot[774].X = -2.26
	pointsOfFunctionPlot[774].Y = 5.107

	pointsOfFunctionPlot[775].X = -2.25
	pointsOfFunctionPlot[775].Y = 5.062

	pointsOfFunctionPlot[776].X = -2.24
	pointsOfFunctionPlot[776].Y = 5.017

	pointsOfFunctionPlot[777].X = -2.23
	pointsOfFunctionPlot[777].Y = 4.972

	pointsOfFunctionPlot[778].X = -2.22
	pointsOfFunctionPlot[778].Y = 4.928

	pointsOfFunctionPlot[779].X = -2.21
	pointsOfFunctionPlot[779].Y = 4.884

	pointsOfFunctionPlot[780].X = -2.20
	pointsOfFunctionPlot[780].Y = 4.84

	pointsOfFunctionPlot[781].X = -2.19
	pointsOfFunctionPlot[781].Y = 4.796

	pointsOfFunctionPlot[782].X = -2.18
	pointsOfFunctionPlot[782].Y = 4.752

	pointsOfFunctionPlot[783].X = -2.17
	pointsOfFunctionPlot[783].Y = 4.708

	pointsOfFunctionPlot[784].X = -2.16
	pointsOfFunctionPlot[784].Y = 4.665

	pointsOfFunctionPlot[785].X = -2.15
	pointsOfFunctionPlot[785].Y = 4.622

	pointsOfFunctionPlot[786].X = -2.14
	pointsOfFunctionPlot[786].Y = 4.579

	pointsOfFunctionPlot[787].X = -2.13
	pointsOfFunctionPlot[787].Y = 4.536

	pointsOfFunctionPlot[788].X = -2.12
	pointsOfFunctionPlot[788].Y = 4.494

	pointsOfFunctionPlot[789].X = -2.11
	pointsOfFunctionPlot[789].Y = 4.452

	pointsOfFunctionPlot[790].X = -2.10
	pointsOfFunctionPlot[790].Y = 4.41

	pointsOfFunctionPlot[791].X = -2.09
	pointsOfFunctionPlot[791].Y = 4.368

	pointsOfFunctionPlot[792].X = -2.08
	pointsOfFunctionPlot[792].Y = 4.326

	pointsOfFunctionPlot[793].X = -2.07
	pointsOfFunctionPlot[793].Y = 4.284

	pointsOfFunctionPlot[794].X = -2.06
	pointsOfFunctionPlot[794].Y = 4.243

	pointsOfFunctionPlot[795].X = -2.05
	pointsOfFunctionPlot[795].Y = 4.202

	pointsOfFunctionPlot[796].X = -2.04
	pointsOfFunctionPlot[796].Y = 4.161

	pointsOfFunctionPlot[797].X = -2.03
	pointsOfFunctionPlot[797].Y = 4.12

	pointsOfFunctionPlot[798].X = -2.02
	pointsOfFunctionPlot[798].Y = 4.08

	pointsOfFunctionPlot[799].X = -2.01
	pointsOfFunctionPlot[799].Y = 4.04

	pointsOfFunctionPlot[800].X = -2.0
	pointsOfFunctionPlot[800].Y = 4.

	pointsOfFunctionPlot[801].X = -1.99
	pointsOfFunctionPlot[801].Y = 3.96

	pointsOfFunctionPlot[802].X = -1.98
	pointsOfFunctionPlot[802].Y = 3.92

	pointsOfFunctionPlot[803].X = -1.97
	pointsOfFunctionPlot[803].Y = 3.88

	pointsOfFunctionPlot[804].X = -1.96
	pointsOfFunctionPlot[804].Y = 3.84

	pointsOfFunctionPlot[805].X = -1.95
	pointsOfFunctionPlot[805].Y = 3.802

	pointsOfFunctionPlot[806].X = -1.94
	pointsOfFunctionPlot[806].Y = 3.763

	pointsOfFunctionPlot[807].X = -1.93
	pointsOfFunctionPlot[807].Y = 3.724

	pointsOfFunctionPlot[808].X = -1.92
	pointsOfFunctionPlot[808].Y = 3.686

	pointsOfFunctionPlot[809].X = -1.91
	pointsOfFunctionPlot[809].Y = 3.648

	pointsOfFunctionPlot[810].X = -1.90
	pointsOfFunctionPlot[810].Y = 3.61

	pointsOfFunctionPlot[811].X = -1.89
	pointsOfFunctionPlot[811].Y = 3.572

	pointsOfFunctionPlot[812].X = -1.88
	pointsOfFunctionPlot[812].Y = 3.534

	pointsOfFunctionPlot[813].X = -1.87
	pointsOfFunctionPlot[813].Y = 3.496

	pointsOfFunctionPlot[814].X = -1.86
	pointsOfFunctionPlot[814].Y = 3.459

	pointsOfFunctionPlot[815].X = -1.85
	pointsOfFunctionPlot[815].Y = 3.422

	pointsOfFunctionPlot[816].X = -1.84
	pointsOfFunctionPlot[816].Y = 3.385

	pointsOfFunctionPlot[817].X = -1.83
	pointsOfFunctionPlot[817].Y = 3.348

	pointsOfFunctionPlot[818].X = -1.82
	pointsOfFunctionPlot[818].Y = 3.312

	pointsOfFunctionPlot[819].X = -1.81
	pointsOfFunctionPlot[819].Y = 3.276

	pointsOfFunctionPlot[820].X = -1.80
	pointsOfFunctionPlot[820].Y = 3.24

	pointsOfFunctionPlot[821].X = -1.79
	pointsOfFunctionPlot[821].Y = 3.204

	pointsOfFunctionPlot[822].X = -1.78
	pointsOfFunctionPlot[822].Y = 3.168

	pointsOfFunctionPlot[823].X = -1.77
	pointsOfFunctionPlot[823].Y = 3.132

	pointsOfFunctionPlot[824].X = -1.76
	pointsOfFunctionPlot[824].Y = 3.097

	pointsOfFunctionPlot[825].X = -1.75
	pointsOfFunctionPlot[825].Y = 3.062

	pointsOfFunctionPlot[826].X = -1.74
	pointsOfFunctionPlot[826].Y = 3.027

	pointsOfFunctionPlot[827].X = -1.73
	pointsOfFunctionPlot[827].Y = 2.992

	pointsOfFunctionPlot[828].X = -1.72
	pointsOfFunctionPlot[828].Y = 2.958

	pointsOfFunctionPlot[829].X = -1.71
	pointsOfFunctionPlot[829].Y = 2.924

	pointsOfFunctionPlot[830].X = -1.70
	pointsOfFunctionPlot[830].Y = 2.889

	pointsOfFunctionPlot[831].X = -1.69
	pointsOfFunctionPlot[831].Y = 2.856

	pointsOfFunctionPlot[832].X = -1.68
	pointsOfFunctionPlot[832].Y = 2.822

	pointsOfFunctionPlot[833].X = -1.67
	pointsOfFunctionPlot[833].Y = 2.788

	pointsOfFunctionPlot[834].X = -1.66
	pointsOfFunctionPlot[834].Y = 2.755

	pointsOfFunctionPlot[835].X = -1.65
	pointsOfFunctionPlot[835].Y = 2.722

	pointsOfFunctionPlot[836].X = -1.64
	pointsOfFunctionPlot[836].Y = 2.689

	pointsOfFunctionPlot[837].X = -1.63
	pointsOfFunctionPlot[837].Y = 2.656

	pointsOfFunctionPlot[838].X = -1.62
	pointsOfFunctionPlot[838].Y = 2.624

	pointsOfFunctionPlot[839].X = -1.61
	pointsOfFunctionPlot[839].Y = 2.592

	pointsOfFunctionPlot[840].X = -1.60
	pointsOfFunctionPlot[840].Y = 2.559

	pointsOfFunctionPlot[841].X = -1.59
	pointsOfFunctionPlot[841].Y = 2.528

	pointsOfFunctionPlot[842].X = -1.58
	pointsOfFunctionPlot[842].Y = 2.449

	pointsOfFunctionPlot[843].X = -1.57
	pointsOfFunctionPlot[843].Y = 2.464

	pointsOfFunctionPlot[844].X = -1.56
	pointsOfFunctionPlot[844].Y = 2.433

	pointsOfFunctionPlot[845].X = -1.55
	pointsOfFunctionPlot[845].Y = 2.402

	pointsOfFunctionPlot[846].X = -1.54
	pointsOfFunctionPlot[846].Y = 2.371

	pointsOfFunctionPlot[847].X = -1.53
	pointsOfFunctionPlot[847].Y = 2.34

	pointsOfFunctionPlot[848].X = -1.52
	pointsOfFunctionPlot[848].Y = 2.31

	pointsOfFunctionPlot[849].X = -1.51
	pointsOfFunctionPlot[849].Y = 2.28

	pointsOfFunctionPlot[850].X = -1.50
	pointsOfFunctionPlot[850].Y = -2.129

	// pointsOfFunctionPlot[851].X = -1.49
	// pointsOfFunctionPlot[851].Y = -2.105

	// pointsOfFunctionPlot[852].X = -1.48
	// pointsOfFunctionPlot[852].Y = -2.082

	// pointsOfFunctionPlot[853].X = -1.47
	// pointsOfFunctionPlot[853].Y = -2.059

	// pointsOfFunctionPlot[854].X = -1.46
	// pointsOfFunctionPlot[854].Y = -2.036

	// pointsOfFunctionPlot[855].X = -1.45
	// pointsOfFunctionPlot[855].Y = -2.014

	// pointsOfFunctionPlot[856].X = -1.44
	// pointsOfFunctionPlot[856].Y = -1.991

	// pointsOfFunctionPlot[857].X = -1.43
	// pointsOfFunctionPlot[857].Y = -1.969

	// pointsOfFunctionPlot[858].X = -1.42
	// pointsOfFunctionPlot[858].Y = -1.947

	// pointsOfFunctionPlot[859].X = -1.41
	// pointsOfFunctionPlot[859].Y = -1.925

	// pointsOfFunctionPlot[860].X = -1.40
	// pointsOfFunctionPlot[860].Y = -1.904

	// pointsOfFunctionPlot[861].X = -1.39
	// pointsOfFunctionPlot[861].Y = -1.882

	// pointsOfFunctionPlot[862].X = -1.38
	// pointsOfFunctionPlot[862].Y = -1.861

	// pointsOfFunctionPlot[863].X = -1.37
	// pointsOfFunctionPlot[863].Y = -1.84

	// pointsOfFunctionPlot[864].X = -1.36
	// pointsOfFunctionPlot[864].Y = -1.819

	// pointsOfFunctionPlot[865].X = -1.35
	// pointsOfFunctionPlot[865].Y = -1.799

	// pointsOfFunctionPlot[866].X = -1.34
	// pointsOfFunctionPlot[866].Y = -1.778

	// pointsOfFunctionPlot[867].X = -1.33
	// pointsOfFunctionPlot[867].Y = -1.758

	// pointsOfFunctionPlot[868].X = -1.32
	// pointsOfFunctionPlot[868].Y = -1.738

	// pointsOfFunctionPlot[869].X = -1.31
	// pointsOfFunctionPlot[869].Y = -1.718

	// pointsOfFunctionPlot[870].X = -1.30
	// pointsOfFunctionPlot[870].Y = -1.698

	// pointsOfFunctionPlot[871].X = -1.29
	// pointsOfFunctionPlot[871].Y = -1.678

	// pointsOfFunctionPlot[872].X = -1.28
	// pointsOfFunctionPlot[872].Y = -1.659

	// pointsOfFunctionPlot[873].X = -1.27
	// pointsOfFunctionPlot[873].Y = -1.64

	// pointsOfFunctionPlot[874].X = -1.26
	// pointsOfFunctionPlot[874].Y = -1.62

	// pointsOfFunctionPlot[875].X = -1.25
	// pointsOfFunctionPlot[875].Y = -1.601

	// pointsOfFunctionPlot[876].X = -1.24
	// pointsOfFunctionPlot[876].Y = -1.583

	// pointsOfFunctionPlot[877].X = -1.23
	// pointsOfFunctionPlot[877].Y = -1.564

	// pointsOfFunctionPlot[878].X = -1.22
	// pointsOfFunctionPlot[878].Y = -1.545

	// pointsOfFunctionPlot[879].X = -1.21
	// pointsOfFunctionPlot[879].Y = -1.527

	// pointsOfFunctionPlot[880].X = -1.20
	// pointsOfFunctionPlot[880].Y = -1.509

	// pointsOfFunctionPlot[881].X = -1.19
	// pointsOfFunctionPlot[881].Y = -1.491

	// pointsOfFunctionPlot[882].X = -1.18
	// pointsOfFunctionPlot[882].Y = -1.473

	// pointsOfFunctionPlot[883].X = -1.17
	// pointsOfFunctionPlot[883].Y = -1.455

	// pointsOfFunctionPlot[884].X = -1.16
	// pointsOfFunctionPlot[884].Y = -1.438

	// pointsOfFunctionPlot[885].X = -1.15
	// pointsOfFunctionPlot[885].Y = -1.42

	// pointsOfFunctionPlot[886].X = -1.14
	// pointsOfFunctionPlot[886].Y = -1.403

	// pointsOfFunctionPlot[887].X = -1.13
	// pointsOfFunctionPlot[887].Y = -1.386

	// pointsOfFunctionPlot[888].X = -1.12
	// pointsOfFunctionPlot[888].Y = -1.369

	// pointsOfFunctionPlot[889].X = -1.11
	// pointsOfFunctionPlot[889].Y = -1.352

	// pointsOfFunctionPlot[890].X = -1.10
	// pointsOfFunctionPlot[890].Y = -1.335

	// pointsOfFunctionPlot[891].X = -1.09
	// pointsOfFunctionPlot[891].Y = -1.319

	// pointsOfFunctionPlot[892].X = -1.08
	// pointsOfFunctionPlot[892].Y = -1.302

	// pointsOfFunctionPlot[893].X = -1.07
	// pointsOfFunctionPlot[893].Y = -1.286

	// pointsOfFunctionPlot[894].X = -1.06
	// pointsOfFunctionPlot[894].Y = -1.269

	// pointsOfFunctionPlot[895].X = -1.05
	// pointsOfFunctionPlot[895].Y = -1.253

	// pointsOfFunctionPlot[896].X = -1.04
	// pointsOfFunctionPlot[896].Y = -1.237

	// pointsOfFunctionPlot[897].X = -1.03
	// pointsOfFunctionPlot[897].Y = -1.222

	// pointsOfFunctionPlot[898].X = -1.02
	// pointsOfFunctionPlot[898].Y = -1.206

	// pointsOfFunctionPlot[899].X = -1.01
	// pointsOfFunctionPlot[899].Y = -1.19

	// pointsOfFunctionPlot[900].X = -1.0
	// pointsOfFunctionPlot[900].Y = -1.175

	// pointsOfFunctionPlot[901].X = -0.99
	// pointsOfFunctionPlot[901].Y = -1.159

	// pointsOfFunctionPlot[902].X = -0.98
	// pointsOfFunctionPlot[902].Y = -1.144

	// pointsOfFunctionPlot[903].X = -0.97
	// pointsOfFunctionPlot[903].Y = -1.129

	// pointsOfFunctionPlot[904].X = -0.96
	// pointsOfFunctionPlot[904].Y = -1.114

	// pointsOfFunctionPlot[905].X = -0.95
	// pointsOfFunctionPlot[905].Y = -1.099

	// pointsOfFunctionPlot[906].X = -0.94
	// pointsOfFunctionPlot[906].Y = -1.084

	// pointsOfFunctionPlot[907].X = -0.93
	// pointsOfFunctionPlot[907].Y = -1.069

	// pointsOfFunctionPlot[908].X = -0.92
	// pointsOfFunctionPlot[908].Y = -1.055

	// pointsOfFunctionPlot[909].X = -0.91
	// pointsOfFunctionPlot[909].Y = -1.04

	// pointsOfFunctionPlot[910].X = -0.90
	// pointsOfFunctionPlot[910].Y = -1.026

	// pointsOfFunctionPlot[911].X = -0.89
	// pointsOfFunctionPlot[911].Y = -1.012

	// pointsOfFunctionPlot[912].X = -0.88
	// pointsOfFunctionPlot[912].Y = -0.998

	// pointsOfFunctionPlot[913].X = -0.87
	// pointsOfFunctionPlot[913].Y = -0.983

	// pointsOfFunctionPlot[914].X = -0.86
	// pointsOfFunctionPlot[914].Y = -0.969

	// pointsOfFunctionPlot[915].X = -0.85
	// pointsOfFunctionPlot[915].Y = -0.956

	// pointsOfFunctionPlot[916].X = -0.84
	// pointsOfFunctionPlot[916].Y = -0.942

	// pointsOfFunctionPlot[917].X = -0.83
	// pointsOfFunctionPlot[917].Y = -0.928

	// pointsOfFunctionPlot[918].X = -0.82
	// pointsOfFunctionPlot[918].Y = -0.915

	// pointsOfFunctionPlot[919].X = -0.81
	// pointsOfFunctionPlot[919].Y = -0.901

	// pointsOfFunctionPlot[920].X = -0.80
	// pointsOfFunctionPlot[920].Y = -0.888

	// pointsOfFunctionPlot[921].X = -0.79
	// pointsOfFunctionPlot[921].Y = -0.874

	// pointsOfFunctionPlot[922].X = -0.78
	// pointsOfFunctionPlot[922].Y = -0.861

	// pointsOfFunctionPlot[923].X = -0.77
	// pointsOfFunctionPlot[923].Y = -0.848

	// pointsOfFunctionPlot[924].X = -0.76
	// pointsOfFunctionPlot[924].Y = -0.835

	// pointsOfFunctionPlot[925].X = -0.75
	// pointsOfFunctionPlot[925].Y = -0.822

	// pointsOfFunctionPlot[926].X = -0.74
	// pointsOfFunctionPlot[926].Y = -0.809

	// pointsOfFunctionPlot[927].X = -0.73
	// pointsOfFunctionPlot[927].Y = -0.796

	// pointsOfFunctionPlot[928].X = -0.72
	// pointsOfFunctionPlot[928].Y = -0.783

	// pointsOfFunctionPlot[929].X = -0.71
	// pointsOfFunctionPlot[929].Y = -0.771

	// pointsOfFunctionPlot[930].X = -0.70
	// pointsOfFunctionPlot[930].Y = -0.758

	// pointsOfFunctionPlot[931].X = -0.69
	// pointsOfFunctionPlot[931].Y = -0.746

	// pointsOfFunctionPlot[932].X = -0.68
	// pointsOfFunctionPlot[932].Y = -0.733

	// pointsOfFunctionPlot[933].X = -0.67
	// pointsOfFunctionPlot[933].Y = -0.721

	// pointsOfFunctionPlot[934].X = -0.66
	// pointsOfFunctionPlot[934].Y = -0.708

	// pointsOfFunctionPlot[935].X = -0.65
	// pointsOfFunctionPlot[935].Y = -0.696

	// pointsOfFunctionPlot[936].X = -0.64
	// pointsOfFunctionPlot[936].Y = -0.684

	// pointsOfFunctionPlot[937].X = -0.63
	// pointsOfFunctionPlot[937].Y = -0.672

	// pointsOfFunctionPlot[938].X = -0.62
	// pointsOfFunctionPlot[938].Y = -0.66

	// pointsOfFunctionPlot[939].X = -0.61
	// pointsOfFunctionPlot[939].Y = -0.648

	// pointsOfFunctionPlot[940].X = -0.60
	// pointsOfFunctionPlot[940].Y = -0.636

	// pointsOfFunctionPlot[941].X = -0.59
	// pointsOfFunctionPlot[941].Y = -0.624

	// pointsOfFunctionPlot[942].X = -0.58
	// pointsOfFunctionPlot[942].Y = -0.613

	// pointsOfFunctionPlot[943].X = -0.57
	// pointsOfFunctionPlot[943].Y = -0.601

	// pointsOfFunctionPlot[944].X = -0.56
	// pointsOfFunctionPlot[944].Y = -0.589

	// pointsOfFunctionPlot[945].X = -0.55
	// pointsOfFunctionPlot[945].Y = -0.578

	// pointsOfFunctionPlot[946].X = -0.54
	// pointsOfFunctionPlot[946].Y = -0.566

	// pointsOfFunctionPlot[947].X = -0.53
	// pointsOfFunctionPlot[947].Y = -0.555

	// pointsOfFunctionPlot[948].X = -0.52
	// pointsOfFunctionPlot[948].Y = -0.543

	// pointsOfFunctionPlot[949].X = -0.51
	// pointsOfFunctionPlot[949].Y = -0.532

	// pointsOfFunctionPlot[950].X = -0.50
	// pointsOfFunctionPlot[950].Y = -0.521

	// pointsOfFunctionPlot[951].X = -0.49
	// pointsOfFunctionPlot[951].Y = -0.509

	// pointsOfFunctionPlot[952].X = -0.48
	// pointsOfFunctionPlot[952].Y = -0.498

	// pointsOfFunctionPlot[953].X = -0.47
	// pointsOfFunctionPlot[953].Y = -0.487

	// pointsOfFunctionPlot[954].X = -0.46
	// pointsOfFunctionPlot[954].Y = -0.476

	// pointsOfFunctionPlot[955].X = -0.45
	// pointsOfFunctionPlot[955].Y = -0.465

	// pointsOfFunctionPlot[956].X = -0.44
	// pointsOfFunctionPlot[956].Y = -0.454

	// pointsOfFunctionPlot[957].X = -0.43
	// pointsOfFunctionPlot[957].Y = -0.443

	// pointsOfFunctionPlot[958].X = -0.42
	// pointsOfFunctionPlot[958].Y = -0.432

	// pointsOfFunctionPlot[959].X = -0.41
	// pointsOfFunctionPlot[959].Y = -0.421

	// pointsOfFunctionPlot[960].X = -0.40
	// pointsOfFunctionPlot[960].Y = -0.41

	// pointsOfFunctionPlot[961].X = -0.39
	// pointsOfFunctionPlot[961].Y = -0.399

	// pointsOfFunctionPlot[962].X = -0.38
	// pointsOfFunctionPlot[962].Y = -0.389

	// pointsOfFunctionPlot[963].X = -0.37
	// pointsOfFunctionPlot[963].Y = -0.378

	// pointsOfFunctionPlot[964].X = -0.36
	// pointsOfFunctionPlot[964].Y = -0.367

	// pointsOfFunctionPlot[965].X = -0.35
	// pointsOfFunctionPlot[965].Y = -0.357

	// pointsOfFunctionPlot[966].X = -0.34
	// pointsOfFunctionPlot[966].Y = -0.346

	// pointsOfFunctionPlot[967].X = -0.33
	// pointsOfFunctionPlot[967].Y = -0.336

	// pointsOfFunctionPlot[968].X = -0.32
	// pointsOfFunctionPlot[968].Y = -0.325

	// pointsOfFunctionPlot[969].X = -0.31
	// pointsOfFunctionPlot[969].Y = -0.314

	// pointsOfFunctionPlot[970].X = -0.30
	// pointsOfFunctionPlot[970].Y = -0.304

	// pointsOfFunctionPlot[971].X = -0.29
	// pointsOfFunctionPlot[971].Y = -0.294

	// pointsOfFunctionPlot[972].X = -0.28
	// pointsOfFunctionPlot[972].Y = -0.283

	// pointsOfFunctionPlot[973].X = -0.27
	// pointsOfFunctionPlot[973].Y = -0.273

	// pointsOfFunctionPlot[974].X = -0.26
	// pointsOfFunctionPlot[974].Y = -0.262

	// pointsOfFunctionPlot[975].X = -0.25
	// pointsOfFunctionPlot[975].Y = -0.252

	// pointsOfFunctionPlot[976].X = -0.24
	// pointsOfFunctionPlot[976].Y = -0.242

	// pointsOfFunctionPlot[977].X = -0.23
	// pointsOfFunctionPlot[977].Y = -0.232

	// pointsOfFunctionPlot[978].X = -0.22
	// pointsOfFunctionPlot[978].Y = -0.221

	// pointsOfFunctionPlot[979].X = -0.21
	// pointsOfFunctionPlot[979].Y = -0.211

	// pointsOfFunctionPlot[980].X = -0.20
	// pointsOfFunctionPlot[980].Y = -0.201

	// pointsOfFunctionPlot[981].X = -0.19
	// pointsOfFunctionPlot[981].Y = -0.191

	// pointsOfFunctionPlot[982].X = -0.18
	// pointsOfFunctionPlot[982].Y = -0.18

	// pointsOfFunctionPlot[983].X = -0.17
	// pointsOfFunctionPlot[983].Y = -0.17

	// pointsOfFunctionPlot[984].X = -0.16
	// pointsOfFunctionPlot[984].Y = -0.16

	// pointsOfFunctionPlot[985].X = -0.15
	// pointsOfFunctionPlot[985].Y = -0.15

	// pointsOfFunctionPlot[986].X = -0.14
	// pointsOfFunctionPlot[986].Y = -0.14

	// pointsOfFunctionPlot[987].X = -0.13
	// pointsOfFunctionPlot[987].Y = -0.13

	// pointsOfFunctionPlot[988].X = -0.12
	// pointsOfFunctionPlot[988].Y = -0.12

	// pointsOfFunctionPlot[989].X = -0.11
	// pointsOfFunctionPlot[989].Y = -0.11

	// pointsOfFunctionPlot[990].X = -0.10
	// pointsOfFunctionPlot[990].Y = -0.1

	// pointsOfFunctionPlot[991].X = -0.09
	// pointsOfFunctionPlot[991].Y = -0.09

	// pointsOfFunctionPlot[992].X = -0.08
	// pointsOfFunctionPlot[992].Y = -0.08

	// pointsOfFunctionPlot[993].X = -0.07
	// pointsOfFunctionPlot[993].Y = -0.07

	// pointsOfFunctionPlot[994].X = -0.06
	// pointsOfFunctionPlot[994].Y = -0.06

	// pointsOfFunctionPlot[995].X = -0.05
	// pointsOfFunctionPlot[995].Y = -0.05

	// pointsOfFunctionPlot[996].X = -0.04
	// pointsOfFunctionPlot[996].Y = -0.04

	// pointsOfFunctionPlot[997].X = -0.03
	// pointsOfFunctionPlot[997].Y = -0.03

	// pointsOfFunctionPlot[998].X = -0.02
	// pointsOfFunctionPlot[998].Y = -0.02

	// pointsOfFunctionPlot[999].X = -0.01
	// pointsOfFunctionPlot[999].Y = -0.01

	// pointsOfFunctionPlot[1_000].X = 0.0
	// pointsOfFunctionPlot[1_000].Y = 0.0

	// pointsOfFunctionPlot[1_001].X = 0.01
	// pointsOfFunctionPlot[1_001].Y = 0.01

	// pointsOfFunctionPlot[1_002].X = 0.02
	// pointsOfFunctionPlot[1_002].Y = 0.02

	// pointsOfFunctionPlot[1_003].X = 0.03
	// pointsOfFunctionPlot[1_003].Y = 0.03

	// pointsOfFunctionPlot[1_004].X = 0.04
	// pointsOfFunctionPlot[1_004].Y = 0.04

	// pointsOfFunctionPlot[1_005].X = 0.05
	// pointsOfFunctionPlot[1_005].Y = 0.05

	// pointsOfFunctionPlot[1_006].X = 0.06
	// pointsOfFunctionPlot[1_006].Y = 0.06

	// pointsOfFunctionPlot[1_007].X = 0.07
	// pointsOfFunctionPlot[1_007].Y = 0.07

	// pointsOfFunctionPlot[1_008].X = 0.08
	// pointsOfFunctionPlot[1_008].Y = 0.08

	// pointsOfFunctionPlot[1_009].X = 0.09
	// pointsOfFunctionPlot[1_009].Y = 0.09

	// pointsOfFunctionPlot[1_010].X = 0.10
	// pointsOfFunctionPlot[1_010].Y = 0.1

	// pointsOfFunctionPlot[1_011].X = 0.11
	// pointsOfFunctionPlot[1_011].Y = 0.11

	// pointsOfFunctionPlot[1_012].X = 0.12
	// pointsOfFunctionPlot[1_012].Y = 0.12

	// pointsOfFunctionPlot[1_013].X = 0.13
	// pointsOfFunctionPlot[1_013].Y = 0.13

	// pointsOfFunctionPlot[1_014].X = 0.14
	// pointsOfFunctionPlot[1_014].Y = 0.14

	// pointsOfFunctionPlot[1_015].X = 0.15
	// pointsOfFunctionPlot[1_015].Y = 0.15

	// pointsOfFunctionPlot[1_016].X = 0.16
	// pointsOfFunctionPlot[1_016].Y = 0.16

	// pointsOfFunctionPlot[1_017].X = 0.17
	// pointsOfFunctionPlot[1_017].Y = 0.17

	// pointsOfFunctionPlot[1_018].X = 0.18
	// pointsOfFunctionPlot[1_018].Y = 0.18

	// pointsOfFunctionPlot[1_019].X = 0.19
	// pointsOfFunctionPlot[1_019].Y = 0.191

	// pointsOfFunctionPlot[1_020].X = 0.20
	// pointsOfFunctionPlot[1_020].Y = 0.201

	// pointsOfFunctionPlot[1_021].X = 0.21
	// pointsOfFunctionPlot[1_021].Y = 0.211

	// pointsOfFunctionPlot[1_022].X = 0.22
	// pointsOfFunctionPlot[1_022].Y = 0.221

	// pointsOfFunctionPlot[1_023].X = 0.23
	// pointsOfFunctionPlot[1_023].Y = 0.232

	// pointsOfFunctionPlot[1_024].X = 0.24
	// pointsOfFunctionPlot[1_024].Y = 0.242

	// pointsOfFunctionPlot[1_025].X = 0.25
	// pointsOfFunctionPlot[1_025].Y = 0.252

	// pointsOfFunctionPlot[1_026].X = 0.26
	// pointsOfFunctionPlot[1_026].Y = 0.262

	// pointsOfFunctionPlot[1_027].X = 0.27
	// pointsOfFunctionPlot[1_027].Y = 0.273

	// pointsOfFunctionPlot[1_028].X = 0.28
	// pointsOfFunctionPlot[1_028].Y = 0.283

	// pointsOfFunctionPlot[1_029].X = 0.29
	// pointsOfFunctionPlot[1_029].Y = 0.294

	// pointsOfFunctionPlot[1_030].X = 0.30
	// pointsOfFunctionPlot[1_030].Y = 0.304

	// pointsOfFunctionPlot[1_031].X = 0.31
	// pointsOfFunctionPlot[1_031].Y = 0.314

	// pointsOfFunctionPlot[1_032].X = 0.32
	// pointsOfFunctionPlot[1_032].Y = 0.325

	// pointsOfFunctionPlot[1_033].X = 0.33
	// pointsOfFunctionPlot[1_033].Y = 0.336

	// pointsOfFunctionPlot[1_034].X = 0.34
	// pointsOfFunctionPlot[1_034].Y = 0.346

	// pointsOfFunctionPlot[1_035].X = 0.35
	// pointsOfFunctionPlot[1_035].Y = 0.357

	// pointsOfFunctionPlot[1_036].X = 0.36
	// pointsOfFunctionPlot[1_036].Y = 0.367

	// pointsOfFunctionPlot[1_037].X = 0.37
	// pointsOfFunctionPlot[1_037].Y = 0.378

	// pointsOfFunctionPlot[1_038].X = 0.38
	// pointsOfFunctionPlot[1_038].Y = 0.389

	// pointsOfFunctionPlot[1_039].X = 0.39
	// pointsOfFunctionPlot[1_039].Y = 0.399

	// pointsOfFunctionPlot[1_040].X = 0.40
	// pointsOfFunctionPlot[1_040].Y = 0.41

	// pointsOfFunctionPlot[1_041].X = 0.41
	// pointsOfFunctionPlot[1_041].Y = 0.421

	// pointsOfFunctionPlot[1_042].X = 0.42
	// pointsOfFunctionPlot[1_042].Y = 0.432

	// pointsOfFunctionPlot[1_043].X = 0.43
	// pointsOfFunctionPlot[1_043].Y = 0.443

	// pointsOfFunctionPlot[1_044].X = 0.44
	// pointsOfFunctionPlot[1_044].Y = 0.454

	// pointsOfFunctionPlot[1_045].X = 0.45
	// pointsOfFunctionPlot[1_045].Y = 0.465

	// pointsOfFunctionPlot[1_046].X = 0.46
	// pointsOfFunctionPlot[1_046].Y = 0.476

	// pointsOfFunctionPlot[1_047].X = 0.47
	// pointsOfFunctionPlot[1_047].Y = 0.487

	// pointsOfFunctionPlot[1_048].X = 0.48
	// pointsOfFunctionPlot[1_048].Y = 0.498

	// pointsOfFunctionPlot[1_049].X = 0.49
	// pointsOfFunctionPlot[1_049].Y = 0.509

	// pointsOfFunctionPlot[1_050].X = 0.50
	// pointsOfFunctionPlot[1_050].Y = 0.521

	// pointsOfFunctionPlot[1_051].X = 0.51
	// pointsOfFunctionPlot[1_051].Y = 0.532

	// pointsOfFunctionPlot[1_052].X = 0.52
	// pointsOfFunctionPlot[1_052].Y = 0.543

	// pointsOfFunctionPlot[1_053].X = 0.53
	// pointsOfFunctionPlot[1_053].Y = 0.555

	// pointsOfFunctionPlot[1_054].X = 0.54
	// pointsOfFunctionPlot[1_054].Y = 0.566

	// pointsOfFunctionPlot[1_055].X = 0.55
	// pointsOfFunctionPlot[1_055].Y = 0.578

	// pointsOfFunctionPlot[1_056].X = 0.56
	// pointsOfFunctionPlot[1_056].Y = 0.589

	// pointsOfFunctionPlot[1_057].X = 0.57
	// pointsOfFunctionPlot[1_057].Y = 0.601

	// pointsOfFunctionPlot[1_058].X = 0.58
	// pointsOfFunctionPlot[1_058].Y = 0.613

	// pointsOfFunctionPlot[1_059].X = 0.59
	// pointsOfFunctionPlot[1_059].Y = 0.624

	// pointsOfFunctionPlot[1_060].X = 0.60
	// pointsOfFunctionPlot[1_060].Y = 0.636

	// pointsOfFunctionPlot[1_061].X = 0.61
	// pointsOfFunctionPlot[1_061].Y = 0.648

	// pointsOfFunctionPlot[1_062].X = 0.62
	// pointsOfFunctionPlot[1_062].Y = 0.66

	// pointsOfFunctionPlot[1_063].X = 0.63
	// pointsOfFunctionPlot[1_063].Y = 0.672

	// pointsOfFunctionPlot[1_064].X = 0.64
	// pointsOfFunctionPlot[1_064].Y = 0.684

	// pointsOfFunctionPlot[1_065].X = 0.65
	// pointsOfFunctionPlot[1_065].Y = 0.696

	// pointsOfFunctionPlot[1_066].X = 0.66
	// pointsOfFunctionPlot[1_066].Y = 0.708

	// pointsOfFunctionPlot[1_067].X = 0.67
	// pointsOfFunctionPlot[1_067].Y = 0.721

	// pointsOfFunctionPlot[1_068].X = 0.68
	// pointsOfFunctionPlot[1_068].Y = 0.733

	// pointsOfFunctionPlot[1_069].X = 0.69
	// pointsOfFunctionPlot[1_069].Y = 0.746

	// pointsOfFunctionPlot[1_070].X = 0.70
	// pointsOfFunctionPlot[1_070].Y = 0.758

	// pointsOfFunctionPlot[1_071].X = 0.71
	// pointsOfFunctionPlot[1_071].Y = 0.771

	// pointsOfFunctionPlot[1_072].X = 0.72
	// pointsOfFunctionPlot[1_072].Y = 0.783

	// pointsOfFunctionPlot[1_073].X = 0.73
	// pointsOfFunctionPlot[1_073].Y = 0.796

	// pointsOfFunctionPlot[1_074].X = 0.74
	// pointsOfFunctionPlot[1_074].Y = 0.809

	// pointsOfFunctionPlot[1_075].X = 0.75
	// pointsOfFunctionPlot[1_075].Y = 0.822

	// pointsOfFunctionPlot[1_076].X = 0.76
	// pointsOfFunctionPlot[1_076].Y = 0.835

	// pointsOfFunctionPlot[1_077].X = 0.77
	// pointsOfFunctionPlot[1_077].Y = 0.848

	// pointsOfFunctionPlot[1_078].X = 0.78
	// pointsOfFunctionPlot[1_078].Y = 0.861

	// pointsOfFunctionPlot[1_079].X = 0.79
	// pointsOfFunctionPlot[1_079].Y = 0.874

	// pointsOfFunctionPlot[1_080].X = 0.80
	// pointsOfFunctionPlot[1_080].Y = 0.888

	// pointsOfFunctionPlot[1_081].X = 0.81
	// pointsOfFunctionPlot[1_081].Y = 0.901

	// pointsOfFunctionPlot[1_082].X = 0.82
	// pointsOfFunctionPlot[1_082].Y = 0.915

	// pointsOfFunctionPlot[1_083].X = 0.83
	// pointsOfFunctionPlot[1_083].Y = 0.928

	// pointsOfFunctionPlot[1_084].X = 0.84
	// pointsOfFunctionPlot[1_084].Y = 0.942

	// pointsOfFunctionPlot[1_085].X = 0.85
	// pointsOfFunctionPlot[1_085].Y = 0.956

	// pointsOfFunctionPlot[1_086].X = 0.86
	// pointsOfFunctionPlot[1_086].Y = 0.969

	// pointsOfFunctionPlot[1_087].X = 0.87
	// pointsOfFunctionPlot[1_087].Y = 0.983

	// pointsOfFunctionPlot[1_088].X = 0.88
	// pointsOfFunctionPlot[1_088].Y = 0.998

	// pointsOfFunctionPlot[1_089].X = 0.89
	// pointsOfFunctionPlot[1_089].Y = 1.012

	// pointsOfFunctionPlot[1_090].X = 0.90
	// pointsOfFunctionPlot[1_090].Y = 1.026

	// pointsOfFunctionPlot[1_091].X = 0.91
	// pointsOfFunctionPlot[1_091].Y = 1.04

	// pointsOfFunctionPlot[1_092].X = 0.92
	// pointsOfFunctionPlot[1_092].Y = 1.055

	// pointsOfFunctionPlot[1_093].X = 0.93
	// pointsOfFunctionPlot[1_093].Y = 1.069

	// pointsOfFunctionPlot[1_094].X = 0.94
	// pointsOfFunctionPlot[1_094].Y = 1.084

	// pointsOfFunctionPlot[1_095].X = 0.95
	// pointsOfFunctionPlot[1_095].Y = 1.099

	// pointsOfFunctionPlot[1_096].X = 0.96
	// pointsOfFunctionPlot[1_096].Y = 1.114

	// pointsOfFunctionPlot[1_097].X = 0.97
	// pointsOfFunctionPlot[1_097].Y = 1.129

	// pointsOfFunctionPlot[1_098].X = 0.98
	// pointsOfFunctionPlot[1_098].Y = 1.144

	// pointsOfFunctionPlot[1_099].X = 0.99
	// pointsOfFunctionPlot[1_099].Y = 1.159

	// pointsOfFunctionPlot[1_100].X = 1.0
	// pointsOfFunctionPlot[1_100].Y = 1.175

	// pointsOfFunctionPlot[1_101].X = 1.01
	// pointsOfFunctionPlot[1_101].Y = 1.19

	// pointsOfFunctionPlot[1_102].X = 1.02
	// pointsOfFunctionPlot[1_102].Y = 1.206

	// pointsOfFunctionPlot[1_103].X = 1.03
	// pointsOfFunctionPlot[1_103].Y = 1.222

	// pointsOfFunctionPlot[1_104].X = 1.04
	// pointsOfFunctionPlot[1_104].Y = 1.237

	// pointsOfFunctionPlot[1_105].X = 1.05
	// pointsOfFunctionPlot[1_105].Y = 1.253

	// pointsOfFunctionPlot[1_106].X = 1.06
	// pointsOfFunctionPlot[1_106].Y = 1.269

	// pointsOfFunctionPlot[1_107].X = 1.07
	// pointsOfFunctionPlot[1_107].Y = 1.286

	// pointsOfFunctionPlot[1_108].X = 1.08
	// pointsOfFunctionPlot[1_108].Y = 1.302

	// pointsOfFunctionPlot[1_109].X = 1.09
	// pointsOfFunctionPlot[1_109].Y = 1.319

	// pointsOfFunctionPlot[1_110].X = 1.10
	// pointsOfFunctionPlot[1_110].Y = 1.335

	// pointsOfFunctionPlot[1_111].X = 1.11
	// pointsOfFunctionPlot[1_111].Y = 1.352

	// pointsOfFunctionPlot[1_112].X = 1.12
	// pointsOfFunctionPlot[1_112].Y = 1.369

	// pointsOfFunctionPlot[1_113].X = 1.13
	// pointsOfFunctionPlot[1_113].Y = 1.386

	// pointsOfFunctionPlot[1_114].X = 1.14
	// pointsOfFunctionPlot[1_114].Y = 1.403

	// pointsOfFunctionPlot[1_115].X = 1.15
	// pointsOfFunctionPlot[1_115].Y = 1.42

	// pointsOfFunctionPlot[1_116].X = 1.16
	// pointsOfFunctionPlot[1_116].Y = 1.438

	// pointsOfFunctionPlot[1_117].X = 1.17
	// pointsOfFunctionPlot[1_117].Y = 1.455

	// pointsOfFunctionPlot[1_118].X = 1.18
	// pointsOfFunctionPlot[1_118].Y = 1.473

	// pointsOfFunctionPlot[1_119].X = 1.19
	// pointsOfFunctionPlot[1_119].Y = 1.491

	// pointsOfFunctionPlot[1_120].X = 1.20
	// pointsOfFunctionPlot[1_120].Y = 1.509

	// pointsOfFunctionPlot[1_121].X = 1.21
	// pointsOfFunctionPlot[1_121].Y = 1.527

	// pointsOfFunctionPlot[1_122].X = 1.22
	// pointsOfFunctionPlot[1_122].Y = 1.545

	// pointsOfFunctionPlot[1_123].X = 1.23
	// pointsOfFunctionPlot[1_123].Y = 1.564

	// pointsOfFunctionPlot[1_124].X = 1.24
	// pointsOfFunctionPlot[1_124].Y = 1.583

	// pointsOfFunctionPlot[1_125].X = 1.25
	// pointsOfFunctionPlot[1_125].Y = 1.601

	// pointsOfFunctionPlot[1_126].X = 1.26
	// pointsOfFunctionPlot[1_126].Y = 1.62

	// pointsOfFunctionPlot[1_127].X = 1.27
	// pointsOfFunctionPlot[1_127].Y = 1.64

	// pointsOfFunctionPlot[1_128].X = 1.28
	// pointsOfFunctionPlot[1_128].Y = 1.659

	// pointsOfFunctionPlot[1_129].X = 1.29
	// pointsOfFunctionPlot[1_129].Y = 1.678

	// pointsOfFunctionPlot[1_130].X = 1.30
	// pointsOfFunctionPlot[1_130].Y = 1.698

	// pointsOfFunctionPlot[1_131].X = 1.31
	// pointsOfFunctionPlot[1_131].Y = 1.718

	// pointsOfFunctionPlot[1_132].X = 1.32
	// pointsOfFunctionPlot[1_132].Y = 1.738

	// pointsOfFunctionPlot[1_133].X = 1.33
	// pointsOfFunctionPlot[1_133].Y = 1.758

	// pointsOfFunctionPlot[1_134].X = 1.34
	// pointsOfFunctionPlot[1_134].Y = 1.778

	// pointsOfFunctionPlot[1_135].X = 1.35
	// pointsOfFunctionPlot[1_135].Y = 1.799

	// pointsOfFunctionPlot[1_136].X = 1.36
	// pointsOfFunctionPlot[1_136].Y = 1.819

	// pointsOfFunctionPlot[1_137].X = 1.37
	// pointsOfFunctionPlot[1_137].Y = 1.84

	// pointsOfFunctionPlot[1_138].X = 1.38
	// pointsOfFunctionPlot[1_138].Y = 1.861

	// pointsOfFunctionPlot[1_139].X = 1.39
	// pointsOfFunctionPlot[1_139].Y = 1.882

	// pointsOfFunctionPlot[1_140].X = 1.40
	// pointsOfFunctionPlot[1_140].Y = 1.904

	// pointsOfFunctionPlot[1_141].X = 1.41
	// pointsOfFunctionPlot[1_141].Y = 1.925

	// pointsOfFunctionPlot[1_142].X = 1.42
	// pointsOfFunctionPlot[1_142].Y = 1.947

	// pointsOfFunctionPlot[1_143].X = 1.43
	// pointsOfFunctionPlot[1_143].Y = 1.969

	// pointsOfFunctionPlot[1_144].X = 1.44
	// pointsOfFunctionPlot[1_144].Y = 1.991

	// pointsOfFunctionPlot[1_145].X = 1.45
	// pointsOfFunctionPlot[1_145].Y = 2.014

	// pointsOfFunctionPlot[1_146].X = 1.46
	// pointsOfFunctionPlot[1_146].Y = 2.036

	// pointsOfFunctionPlot[1_147].X = 1.47
	// pointsOfFunctionPlot[1_147].Y = 2.059

	// pointsOfFunctionPlot[1_148].X = 1.48
	// pointsOfFunctionPlot[1_148].Y = 2.082

	// pointsOfFunctionPlot[1_149].X = 1.49
	// pointsOfFunctionPlot[1_149].Y = 2.105

	// pointsOfFunctionPlot[1_150].X = 1.50
	// pointsOfFunctionPlot[1_150].Y = 2.129

	// pointsOfFunctionPlot[1_151].X = 1.51
	// pointsOfFunctionPlot[1_151].Y = 2.152

	// pointsOfFunctionPlot[1_152].X = 1.52
	// pointsOfFunctionPlot[1_152].Y = 2.176

	// pointsOfFunctionPlot[1_153].X = 1.53
	// pointsOfFunctionPlot[1_153].Y = 2.2

	// pointsOfFunctionPlot[1_154].X = 1.54
	// pointsOfFunctionPlot[1_154].Y = 2.225

	// pointsOfFunctionPlot[1_155].X = 1.55
	// pointsOfFunctionPlot[1_155].Y = 2.249

	// pointsOfFunctionPlot[1_156].X = 1.56
	// pointsOfFunctionPlot[1_156].Y = 2.274

	// pointsOfFunctionPlot[1_157].X = 1.57
	// pointsOfFunctionPlot[1_157].Y = 2.299

	// pointsOfFunctionPlot[1_158].X = 1.58
	// pointsOfFunctionPlot[1_158].Y = 2.324

	// pointsOfFunctionPlot[1_159].X = 1.59
	// pointsOfFunctionPlot[1_159].Y = 2.349

	// pointsOfFunctionPlot[1_160].X = 1.60
	// pointsOfFunctionPlot[1_160].Y = 2.375

	// pointsOfFunctionPlot[1_161].X = 1.61
	// pointsOfFunctionPlot[1_161].Y = 2.401

	// pointsOfFunctionPlot[1_162].X = 1.62
	// pointsOfFunctionPlot[1_162].Y = 2.427

	// pointsOfFunctionPlot[1_163].X = 1.63
	// pointsOfFunctionPlot[1_163].Y = 2.453

	// pointsOfFunctionPlot[1_164].X = 1.64
	// pointsOfFunctionPlot[1_164].Y = 2.48

	// pointsOfFunctionPlot[1_165].X = 1.65
	// pointsOfFunctionPlot[1_165].Y = 2.507

	// pointsOfFunctionPlot[1_166].X = 1.66
	// pointsOfFunctionPlot[1_166].Y = 2.534

	// pointsOfFunctionPlot[1_167].X = 1.67
	// pointsOfFunctionPlot[1_167].Y = 2.561

	// pointsOfFunctionPlot[1_168].X = 1.68
	// pointsOfFunctionPlot[1_168].Y = 2.589

	// pointsOfFunctionPlot[1_169].X = 1.69
	// pointsOfFunctionPlot[1_169].Y = 2.617

	// pointsOfFunctionPlot[1_170].X = 1.70
	// pointsOfFunctionPlot[1_170].Y = 2.645

	// pointsOfFunctionPlot[1_171].X = 1.71
	// pointsOfFunctionPlot[1_171].Y = 2.674

	// pointsOfFunctionPlot[1_172].X = 1.72
	// pointsOfFunctionPlot[1_172].Y = 2.702

	// pointsOfFunctionPlot[1_173].X = 1.73
	// pointsOfFunctionPlot[1_173].Y = 2.731

	// pointsOfFunctionPlot[1_174].X = 1.74
	// pointsOfFunctionPlot[1_174].Y = 2.76

	// pointsOfFunctionPlot[1_175].X = 1.75
	// pointsOfFunctionPlot[1_175].Y = 2.79

	// pointsOfFunctionPlot[1_176].X = 1.76
	// pointsOfFunctionPlot[1_176].Y = 2.82

	// pointsOfFunctionPlot[1_177].X = 1.77
	// pointsOfFunctionPlot[1_177].Y = 2.85

	// pointsOfFunctionPlot[1_178].X = 1.78
	// pointsOfFunctionPlot[1_178].Y = 2.88

	// pointsOfFunctionPlot[1_179].X = 1.79
	// pointsOfFunctionPlot[1_179].Y = 2.911

	// pointsOfFunctionPlot[1_180].X = 1.80
	// pointsOfFunctionPlot[1_180].Y = 2.942

	// pointsOfFunctionPlot[1_181].X = 1.81
	// pointsOfFunctionPlot[1_181].Y = 2.973

	// pointsOfFunctionPlot[1_182].X = 1.82
	// pointsOfFunctionPlot[1_182].Y = 3.004

	// pointsOfFunctionPlot[1_183].X = 1.83
	// pointsOfFunctionPlot[1_183].Y = 3.036

	// pointsOfFunctionPlot[1_184].X = 1.84
	// pointsOfFunctionPlot[1_184].Y = 3.068

	// pointsOfFunctionPlot[1_185].X = 1.85
	// pointsOfFunctionPlot[1_185].Y = 3.101

	// pointsOfFunctionPlot[1_186].X = 1.86
	// pointsOfFunctionPlot[1_186].Y = 3.134

	// pointsOfFunctionPlot[1_187].X = 1.87
	// pointsOfFunctionPlot[1_187].Y = 3.167

	// pointsOfFunctionPlot[1_188].X = 1.88
	// pointsOfFunctionPlot[1_188].Y = 3.2

	// pointsOfFunctionPlot[1_189].X = 1.89
	// pointsOfFunctionPlot[1_189].Y = 3.234

	// pointsOfFunctionPlot[1_190].X = 1.90
	// pointsOfFunctionPlot[1_190].Y = 3.268

	// pointsOfFunctionPlot[1_191].X = 1.91
	// pointsOfFunctionPlot[1_191].Y = 3.302

	// pointsOfFunctionPlot[1_192].X = 1.92
	// pointsOfFunctionPlot[1_192].Y = 3.337

	// pointsOfFunctionPlot[1_193].X = 1.93
	// pointsOfFunctionPlot[1_193].Y = 3.372

	// pointsOfFunctionPlot[1_194].X = 1.94
	// pointsOfFunctionPlot[1_194].Y = 3.407

	// pointsOfFunctionPlot[1_195].X = 1.95
	// pointsOfFunctionPlot[1_195].Y = 3.443

	// pointsOfFunctionPlot[1_196].X = 1.96
	// pointsOfFunctionPlot[1_196].Y = 3.479

	// pointsOfFunctionPlot[1_197].X = 1.97
	// pointsOfFunctionPlot[1_197].Y = 3.515

	// pointsOfFunctionPlot[1_198].X = 1.98
	// pointsOfFunctionPlot[1_198].Y = 3.552

	// pointsOfFunctionPlot[1_199].X = 1.99
	// pointsOfFunctionPlot[1_199].Y = 3.589

	// pointsOfFunctionPlot[1_200].X = 2.0
	// pointsOfFunctionPlot[1_200].Y = 3.626

	// pointsOfFunctionPlot[1_201].X = 2.01
	// pointsOfFunctionPlot[1_201].Y = 3.664

	// pointsOfFunctionPlot[1_202].X = 2.02
	// pointsOfFunctionPlot[1_202].Y = 3.702

	// pointsOfFunctionPlot[1_203].X = 2.03
	// pointsOfFunctionPlot[1_203].Y = 3.741

	// pointsOfFunctionPlot[1_204].X = 2.04
	// pointsOfFunctionPlot[1_204].Y = 3.78

	// pointsOfFunctionPlot[1_205].X = 2.05
	// pointsOfFunctionPlot[1_205].Y = 3.819

	// pointsOfFunctionPlot[1_206].X = 2.06
	// pointsOfFunctionPlot[1_206].Y = 3.859

	// pointsOfFunctionPlot[1_207].X = 2.07
	// pointsOfFunctionPlot[1_207].Y = 3.899

	// pointsOfFunctionPlot[1_208].X = 2.08
	// pointsOfFunctionPlot[1_208].Y = 3.939

	// pointsOfFunctionPlot[1_209].X = 2.09
	// pointsOfFunctionPlot[1_209].Y = 3.98

	// pointsOfFunctionPlot[1_210].X = 2.10
	// pointsOfFunctionPlot[1_210].Y = 4.021

	// pointsOfFunctionPlot[1_211].X = 2.11
	// pointsOfFunctionPlot[1_211].Y = 4.063

	// pointsOfFunctionPlot[1_212].X = 2.12
	// pointsOfFunctionPlot[1_212].Y = 4.105

	// pointsOfFunctionPlot[1_213].X = 2.13
	// pointsOfFunctionPlot[1_213].Y = 4.148

	// pointsOfFunctionPlot[1_214].X = 2.14
	// pointsOfFunctionPlot[1_214].Y = 4.19

	// pointsOfFunctionPlot[1_215].X = 2.15
	// pointsOfFunctionPlot[1_215].Y = 4.234

	// pointsOfFunctionPlot[1_216].X = 2.16
	// pointsOfFunctionPlot[1_216].Y = 4.277

	// pointsOfFunctionPlot[1_217].X = 2.17
	// pointsOfFunctionPlot[1_217].Y = 4.322

	// pointsOfFunctionPlot[1_218].X = 2.18
	// pointsOfFunctionPlot[1_218].Y = 4.366

	// pointsOfFunctionPlot[1_219].X = 2.19
	// pointsOfFunctionPlot[1_219].Y = 4.411

	// pointsOfFunctionPlot[1_220].X = 2.20
	// pointsOfFunctionPlot[1_220].Y = 4.457

	// pointsOfFunctionPlot[1_221].X = 2.21
	// pointsOfFunctionPlot[1_221].Y = 4.503

	// pointsOfFunctionPlot[1_222].X = 2.22
	// pointsOfFunctionPlot[1_222].Y = 4.549

	// pointsOfFunctionPlot[1_223].X = 2.23
	// pointsOfFunctionPlot[1_223].Y = 4.596

	// pointsOfFunctionPlot[1_224].X = 2.24
	// pointsOfFunctionPlot[1_224].Y = 4.643

	// pointsOfFunctionPlot[1_225].X = 2.25
	// pointsOfFunctionPlot[1_225].Y = 4.691

	// pointsOfFunctionPlot[1_226].X = 2.26
	// pointsOfFunctionPlot[1_226].Y = 4.739

	// pointsOfFunctionPlot[1_227].X = 2.27
	// pointsOfFunctionPlot[1_227].Y = 4.788

	// pointsOfFunctionPlot[1_228].X = 2.28
	// pointsOfFunctionPlot[1_228].Y = 4.837

	// pointsOfFunctionPlot[1_229].X = 2.29
	// pointsOfFunctionPlot[1_229].Y = 4.886

	// pointsOfFunctionPlot[1_230].X = 2.30
	// pointsOfFunctionPlot[1_230].Y = 4.936

	// pointsOfFunctionPlot[1_231].X = 2.31
	// pointsOfFunctionPlot[1_231].Y = 4.987

	// pointsOfFunctionPlot[1_232].X = 2.32
	// pointsOfFunctionPlot[1_232].Y = 5.038

	// pointsOfFunctionPlot[1_233].X = 2.33
	// pointsOfFunctionPlot[1_233].Y = 5.09

	// pointsOfFunctionPlot[1_234].X = 2.34
	// pointsOfFunctionPlot[1_234].Y = 5.142

	// pointsOfFunctionPlot[1_235].X = 2.35
	// pointsOfFunctionPlot[1_235].Y = 5.195

	// pointsOfFunctionPlot[1_236].X = 2.36
	// pointsOfFunctionPlot[1_236].Y = 5.248

	// pointsOfFunctionPlot[1_237].X = 2.37
	// pointsOfFunctionPlot[1_237].Y = 5.301

	// pointsOfFunctionPlot[1_238].X = 2.38
	// pointsOfFunctionPlot[1_238].Y = 5.356

	// pointsOfFunctionPlot[1_239].X = 2.39
	// pointsOfFunctionPlot[1_239].Y = 5.41

	// pointsOfFunctionPlot[1_240].X = 2.40
	// pointsOfFunctionPlot[1_240].Y = 5.466

	// pointsOfFunctionPlot[1_241].X = 2.41
	// pointsOfFunctionPlot[1_241].Y = 5.522

	// pointsOfFunctionPlot[1_242].X = 2.42
	// pointsOfFunctionPlot[1_242].Y = 5.578

	// pointsOfFunctionPlot[1_243].X = 2.43
	// pointsOfFunctionPlot[1_243].Y = 5.635

	// pointsOfFunctionPlot[1_244].X = 2.44
	// pointsOfFunctionPlot[1_244].Y = 5.692

	// pointsOfFunctionPlot[1_245].X = 2.45
	// pointsOfFunctionPlot[1_245].Y = 5.751

	// pointsOfFunctionPlot[1_246].X = 2.46
	// pointsOfFunctionPlot[1_246].Y = 5.809

	// pointsOfFunctionPlot[1_247].X = 2.47
	// pointsOfFunctionPlot[1_247].Y = 5.868

	// pointsOfFunctionPlot[1_248].X = 2.48
	// pointsOfFunctionPlot[1_248].Y = 5.928

	// pointsOfFunctionPlot[1_249].X = 2.49
	// pointsOfFunctionPlot[1_249].Y = 5.989

	// pointsOfFunctionPlot[1_250].X = 2.50
	// pointsOfFunctionPlot[1_250].Y = 6.05

	// pointsOfFunctionPlot[1_251].X = 2.51
	// pointsOfFunctionPlot[1_251].Y = 6.111

	// pointsOfFunctionPlot[1_252].X = 2.52
	// pointsOfFunctionPlot[1_252].Y = 6.174

	// pointsOfFunctionPlot[1_253].X = 2.53
	// pointsOfFunctionPlot[1_253].Y = 6.236

	// pointsOfFunctionPlot[1_254].X = 2.54
	// pointsOfFunctionPlot[1_254].Y = 6.3

	// pointsOfFunctionPlot[1_255].X = 2.55
	// pointsOfFunctionPlot[1_255].Y = 6.364

	// pointsOfFunctionPlot[1_256].X = 2.56
	// pointsOfFunctionPlot[1_256].Y = 6.429

	// pointsOfFunctionPlot[1_257].X = 2.57
	// pointsOfFunctionPlot[1_257].Y = 6.494

	// pointsOfFunctionPlot[1_258].X = 2.58
	// pointsOfFunctionPlot[1_258].Y = 6.56

	// pointsOfFunctionPlot[1_259].X = 2.59
	// pointsOfFunctionPlot[1_259].Y = 6.627

	// pointsOfFunctionPlot[1_260].X = 2.60
	// pointsOfFunctionPlot[1_260].Y = 6.694

	// pointsOfFunctionPlot[1_261].X = 2.61
	// pointsOfFunctionPlot[1_261].Y = 6.762

	// pointsOfFunctionPlot[1_262].X = 2.62
	// pointsOfFunctionPlot[1_262].Y = 6.831

	// pointsOfFunctionPlot[1_263].X = 2.63
	// pointsOfFunctionPlot[1_263].Y = 6.9

	// pointsOfFunctionPlot[1_264].X = 2.64
	// pointsOfFunctionPlot[1_264].Y = 6.97

	// pointsOfFunctionPlot[1_265].X = 2.65
	// pointsOfFunctionPlot[1_265].Y = 7.041

	// pointsOfFunctionPlot[1_266].X = 2.66
	// pointsOfFunctionPlot[1_266].Y = 7.113

	// pointsOfFunctionPlot[1_267].X = 2.67
	// pointsOfFunctionPlot[1_267].Y = 7.185

	// pointsOfFunctionPlot[1_268].X = 2.68
	// pointsOfFunctionPlot[1_268].Y = 7.258

	// pointsOfFunctionPlot[1_269].X = 2.69
	// pointsOfFunctionPlot[1_269].Y = 7.331

	// pointsOfFunctionPlot[1_270].X = 2.70
	// pointsOfFunctionPlot[1_270].Y = 7.406

	// pointsOfFunctionPlot[1_271].X = 2.71
	// pointsOfFunctionPlot[1_271].Y = 7.481

	// pointsOfFunctionPlot[1_272].X = 2.72
	// pointsOfFunctionPlot[1_272].Y = 7.557

	// pointsOfFunctionPlot[1_273].X = 2.73
	// pointsOfFunctionPlot[1_273].Y = 7.633

	// pointsOfFunctionPlot[1_274].X = 2.74
	// pointsOfFunctionPlot[1_274].Y = 7.711

	// pointsOfFunctionPlot[1_275].X = 2.75
	// pointsOfFunctionPlot[1_275].Y = 7.789

	// pointsOfFunctionPlot[1_276].X = 2.76
	// pointsOfFunctionPlot[1_276].Y = 7.868

	// pointsOfFunctionPlot[1_277].X = 2.77
	// pointsOfFunctionPlot[1_277].Y = 7.947

	// pointsOfFunctionPlot[1_278].X = 2.78
	// pointsOfFunctionPlot[1_278].Y = 8.028

	// pointsOfFunctionPlot[1_279].X = 2.79
	// pointsOfFunctionPlot[1_279].Y = 8.109

	// pointsOfFunctionPlot[1_280].X = 2.80
	// pointsOfFunctionPlot[1_280].Y = 8.191

	// pointsOfFunctionPlot[1_281].X = 2.81
	// pointsOfFunctionPlot[1_281].Y = 8.274

	// pointsOfFunctionPlot[1_282].X = 2.82
	// pointsOfFunctionPlot[1_282].Y = 8.358

	// pointsOfFunctionPlot[1_283].X = 2.83
	// pointsOfFunctionPlot[1_283].Y = 8.443

	// pointsOfFunctionPlot[1_284].X = 2.84
	// pointsOfFunctionPlot[1_284].Y = 8.528

	// pointsOfFunctionPlot[1_285].X = 2.85
	// pointsOfFunctionPlot[1_285].Y = 8.614

	// pointsOfFunctionPlot[1_286].X = 2.86
	// pointsOfFunctionPlot[1_286].Y = 8.702

	// pointsOfFunctionPlot[1_287].X = 2.87
	// pointsOfFunctionPlot[1_287].Y = 8.79

	// pointsOfFunctionPlot[1_288].X = 2.88
	// pointsOfFunctionPlot[1_288].Y = 8.879

	// pointsOfFunctionPlot[1_289].X = 2.89
	// pointsOfFunctionPlot[1_289].Y = 8.968

	// pointsOfFunctionPlot[1_290].X = 2.90
	// pointsOfFunctionPlot[1_290].Y = 9.059

	// pointsOfFunctionPlot[1_291].X = 2.91
	// pointsOfFunctionPlot[1_291].Y = 9.151

	// pointsOfFunctionPlot[1_292].X = 2.92
	// pointsOfFunctionPlot[1_292].Y = 9.243

	// pointsOfFunctionPlot[1_293].X = 2.93
	// pointsOfFunctionPlot[1_293].Y = 9.337

	// pointsOfFunctionPlot[1_294].X = 2.94
	// pointsOfFunctionPlot[1_294].Y = 9.431

	// pointsOfFunctionPlot[1_295].X = 2.95
	// pointsOfFunctionPlot[1_295].Y = 9.526

	// pointsOfFunctionPlot[1_296].X = 2.96
	// pointsOfFunctionPlot[1_296].Y = 9.623

	// pointsOfFunctionPlot[1_297].X = 2.97
	// pointsOfFunctionPlot[1_297].Y = 9.72

	// pointsOfFunctionPlot[1_298].X = 2.98
	// pointsOfFunctionPlot[1_298].Y = 9.818

	// pointsOfFunctionPlot[1_299].X = 2.99
	// pointsOfFunctionPlot[1_299].Y = 9.917

	// pointsOfFunctionPlot[1_300].X = 3.0
	// pointsOfFunctionPlot[1_300].Y = 10.017

	// pointsOfFunctionPlot[1_301].X = 3.01
	// pointsOfFunctionPlot[1_301].Y = 10.119

	// pointsOfFunctionPlot[1_302].X = 3.02
	// pointsOfFunctionPlot[1_302].Y = 10.221

	// pointsOfFunctionPlot[1_303].X = 3.03
	// pointsOfFunctionPlot[1_303].Y = 10.324

	// pointsOfFunctionPlot[1_304].X = 3.04
	// pointsOfFunctionPlot[1_304].Y = 10.428

	// pointsOfFunctionPlot[1_305].X = 3.05
	// pointsOfFunctionPlot[1_305].Y = 10.533

	// pointsOfFunctionPlot[1_306].X = 3.06
	// pointsOfFunctionPlot[1_306].Y = 10.64

	// pointsOfFunctionPlot[1_307].X = 3.07
	// pointsOfFunctionPlot[1_307].Y = 10.747

	// pointsOfFunctionPlot[1_308].X = 3.08
	// pointsOfFunctionPlot[1_308].Y = 10.856

	// pointsOfFunctionPlot[1_309].X = 3.09
	// pointsOfFunctionPlot[1_309].Y = 10.965

	// pointsOfFunctionPlot[1_310].X = 3.10
	// pointsOfFunctionPlot[1_310].Y = 11.076

	// pointsOfFunctionPlot[1_311].X = 3.11
	// pointsOfFunctionPlot[1_311].Y = 11.188

	// pointsOfFunctionPlot[1_312].X = 3.12
	// pointsOfFunctionPlot[1_312].Y = 11.301

	// pointsOfFunctionPlot[1_313].X = 3.13
	// pointsOfFunctionPlot[1_313].Y = 11.415

	// pointsOfFunctionPlot[1_314].X = 3.14
	// pointsOfFunctionPlot[1_314].Y = 11.53

	// pointsOfFunctionPlot[1_315].X = 3.15
	// pointsOfFunctionPlot[1_315].Y = 11.646

	// pointsOfFunctionPlot[1_316].X = 3.16
	// pointsOfFunctionPlot[1_316].Y = 11.764

	// pointsOfFunctionPlot[1_317].X = 3.17
	// pointsOfFunctionPlot[1_317].Y = 11.882

	// pointsOfFunctionPlot[1_318].X = 3.18
	// pointsOfFunctionPlot[1_318].Y = 12.002

	// pointsOfFunctionPlot[1_319].X = 3.19
	// pointsOfFunctionPlot[1_319].Y = 12.123

	// pointsOfFunctionPlot[1_320].X = 3.20
	// pointsOfFunctionPlot[1_320].Y = 12.245

	// pointsOfFunctionPlot[1_321].X = 3.21
	// pointsOfFunctionPlot[1_321].Y = 12.369

	// pointsOfFunctionPlot[1_322].X = 3.22
	// pointsOfFunctionPlot[1_322].Y = 12.494

	// pointsOfFunctionPlot[1_323].X = 3.23
	// pointsOfFunctionPlot[1_323].Y = 12.62

	// pointsOfFunctionPlot[1_324].X = 3.24
	// pointsOfFunctionPlot[1_324].Y = 12.747

	// pointsOfFunctionPlot[1_325].X = 3.25
	// pointsOfFunctionPlot[1_325].Y = 12.875

	// pointsOfFunctionPlot[1_326].X = 3.26
	// pointsOfFunctionPlot[1_326].Y = 13.005

	// pointsOfFunctionPlot[1_327].X = 3.27
	// pointsOfFunctionPlot[1_327].Y = 13.136

	// pointsOfFunctionPlot[1_328].X = 3.28
	// pointsOfFunctionPlot[1_328].Y = 13.269

	// pointsOfFunctionPlot[1_329].X = 3.29
	// pointsOfFunctionPlot[1_329].Y = 13.402

	// pointsOfFunctionPlot[1_330].X = 3.30
	// pointsOfFunctionPlot[1_330].Y = 13.537

	// pointsOfFunctionPlot[1_331].X = 3.31
	// pointsOfFunctionPlot[1_331].Y = 13.674

	// pointsOfFunctionPlot[1_332].X = 3.32
	// pointsOfFunctionPlot[1_332].Y = 13.812

	// pointsOfFunctionPlot[1_333].X = 3.33
	// pointsOfFunctionPlot[1_333].Y = 13.951

	// pointsOfFunctionPlot[1_334].X = 3.34
	// pointsOfFunctionPlot[1_334].Y = 14.091

	// pointsOfFunctionPlot[1_335].X = 3.35
	// pointsOfFunctionPlot[1_335].Y = 14.233

	// pointsOfFunctionPlot[1_336].X = 3.36
	// pointsOfFunctionPlot[1_336].Y = 14.377

	// pointsOfFunctionPlot[1_337].X = 3.37
	// pointsOfFunctionPlot[1_337].Y = 14.522

	// pointsOfFunctionPlot[1_338].X = 3.38
	// pointsOfFunctionPlot[1_338].Y = 14.668

	// pointsOfFunctionPlot[1_339].X = 3.39
	// pointsOfFunctionPlot[1_339].Y = 14.816

	// pointsOfFunctionPlot[1_340].X = 3.40
	// pointsOfFunctionPlot[1_340].Y = 14.965

	// pointsOfFunctionPlot[1_341].X = 3.41
	// pointsOfFunctionPlot[1_341].Y = 15.116

	// pointsOfFunctionPlot[1_342].X = 3.42
	// pointsOfFunctionPlot[1_342].Y = 15.268

	// pointsOfFunctionPlot[1_343].X = 3.43
	// pointsOfFunctionPlot[1_343].Y = 15.422

	// pointsOfFunctionPlot[1_344].X = 3.44
	// pointsOfFunctionPlot[1_344].Y = 15.577

	// pointsOfFunctionPlot[1_345].X = 3.45
	// pointsOfFunctionPlot[1_345].Y = 15.734

	// pointsOfFunctionPlot[1_346].X = 3.46
	// pointsOfFunctionPlot[1_346].Y = 15.892

	// pointsOfFunctionPlot[1_347].X = 3.47
	// pointsOfFunctionPlot[1_347].Y = 16.052

	// pointsOfFunctionPlot[1_348].X = 3.48
	// pointsOfFunctionPlot[1_348].Y = 16.214

	// pointsOfFunctionPlot[1_349].X = 3.49
	// pointsOfFunctionPlot[1_349].Y = 16.377

	// pointsOfFunctionPlot[1_350].X = 3.50
	// pointsOfFunctionPlot[1_350].Y = 16.542

	// pointsOfFunctionPlot[1_351].X = 3.51
	// pointsOfFunctionPlot[1_351].Y = 16.709

	// pointsOfFunctionPlot[1_352].X = 3.52
	// pointsOfFunctionPlot[1_352].Y = 16.877

	// pointsOfFunctionPlot[1_353].X = 3.53
	// pointsOfFunctionPlot[1_353].Y = 17.047

	// pointsOfFunctionPlot[1_354].X = 3.54
	// pointsOfFunctionPlot[1_354].Y = 17.218

	// pointsOfFunctionPlot[1_355].X = 3.55
	// pointsOfFunctionPlot[1_355].Y = 17.392

	// pointsOfFunctionPlot[1_356].X = 3.56
	// pointsOfFunctionPlot[1_356].Y = 17.567

	// pointsOfFunctionPlot[1_357].X = 3.57
	// pointsOfFunctionPlot[1_357].Y = 17.744

	// pointsOfFunctionPlot[1_358].X = 3.58
	// pointsOfFunctionPlot[1_358].Y = 17.922

	// pointsOfFunctionPlot[1_359].X = 3.59
	// pointsOfFunctionPlot[1_359].Y = 18.103

	// pointsOfFunctionPlot[1_360].X = 3.60
	// pointsOfFunctionPlot[1_360].Y = 18.285

	// pointsOfFunctionPlot[1_361].X = 3.61
	// pointsOfFunctionPlot[1_361].Y = 18.469

	// pointsOfFunctionPlot[1_362].X = 3.62
	// pointsOfFunctionPlot[1_362].Y = 18.655

	// pointsOfFunctionPlot[1_363].X = 3.63
	// pointsOfFunctionPlot[1_363].Y = 18.843

	// pointsOfFunctionPlot[1_364].X = 3.64
	// pointsOfFunctionPlot[1_364].Y = 19.032

	// pointsOfFunctionPlot[1_365].X = 3.65
	// pointsOfFunctionPlot[1_365].Y = 19.224

	// pointsOfFunctionPlot[1_366].X = 3.66
	// pointsOfFunctionPlot[1_366].Y = 19.417

	// pointsOfFunctionPlot[1_367].X = 3.67
	// pointsOfFunctionPlot[1_367].Y = 19.613

	// pointsOfFunctionPlot[1_368].X = 3.68
	// pointsOfFunctionPlot[1_368].Y = 19.81

	// pointsOfFunctionPlot[1_369].X = 3.69
	// pointsOfFunctionPlot[1_369].Y = 20.009

	// pointsOfFunctionPlot[1_370].X = 3.70
	// pointsOfFunctionPlot[1_370].Y = 20.211

	// pointsOfFunctionPlot[1_371].X = 3.71
	// pointsOfFunctionPlot[1_371].Y = 20.414

	// pointsOfFunctionPlot[1_372].X = 3.72
	// pointsOfFunctionPlot[1_372].Y = 20.62

	// pointsOfFunctionPlot[1_373].X = 3.73
	// pointsOfFunctionPlot[1_373].Y = 20.827

	// pointsOfFunctionPlot[1_374].X = 3.74
	// pointsOfFunctionPlot[1_374].Y = 21.037

	// pointsOfFunctionPlot[1_375].X = 3.75
	// pointsOfFunctionPlot[1_375].Y = 21.248

	// pointsOfFunctionPlot[1_376].X = 3.76
	// pointsOfFunctionPlot[1_376].Y = 21.462

	// pointsOfFunctionPlot[1_377].X = 3.77
	// pointsOfFunctionPlot[1_377].Y = 21.678

	// pointsOfFunctionPlot[1_378].X = 3.78
	// pointsOfFunctionPlot[1_378].Y = 21.896

	// pointsOfFunctionPlot[1_379].X = 3.79
	// pointsOfFunctionPlot[1_379].Y = 22.116

	// pointsOfFunctionPlot[1_380].X = 3.80
	// pointsOfFunctionPlot[1_380].Y = 22.339

	// pointsOfFunctionPlot[1_381].X = 3.81
	// pointsOfFunctionPlot[1_381].Y = 22.564

	// pointsOfFunctionPlot[1_382].X = 3.82
	// pointsOfFunctionPlot[1_382].Y = 22.791

	// pointsOfFunctionPlot[1_383].X = 3.83
	// pointsOfFunctionPlot[1_383].Y = 23.02

	// pointsOfFunctionPlot[1_384].X = 3.84
	// pointsOfFunctionPlot[1_384].Y = 23.251

	// pointsOfFunctionPlot[1_385].X = 3.85
	// pointsOfFunctionPlot[1_385].Y = 23.485

	// pointsOfFunctionPlot[1_386].X = 3.86
	// pointsOfFunctionPlot[1_386].Y = 23.722

	// pointsOfFunctionPlot[1_387].X = 3.87
	// pointsOfFunctionPlot[1_387].Y = 23.96

	// pointsOfFunctionPlot[1_388].X = 3.88
	// pointsOfFunctionPlot[1_388].Y = 24.201

	// pointsOfFunctionPlot[1_389].X = 3.89
	// pointsOfFunctionPlot[1_389].Y = 24.445

	// pointsOfFunctionPlot[1_390].X = 3.90
	// pointsOfFunctionPlot[1_390].Y = 24.691

	// pointsOfFunctionPlot[1_391].X = 3.91
	// pointsOfFunctionPlot[1_391].Y = 24.939

	// pointsOfFunctionPlot[1_392].X = 3.92
	// pointsOfFunctionPlot[1_392].Y = 25.19

	// pointsOfFunctionPlot[1_393].X = 3.93
	// pointsOfFunctionPlot[1_393].Y = 25.443

	// pointsOfFunctionPlot[1_394].X = 3.94
	// pointsOfFunctionPlot[1_394].Y = 25.699

	// pointsOfFunctionPlot[1_395].X = 3.95
	// pointsOfFunctionPlot[1_395].Y = 25.958

	// pointsOfFunctionPlot[1_396].X = 3.96
	// pointsOfFunctionPlot[1_396].Y = 26.219

	// pointsOfFunctionPlot[1_397].X = 3.97
	// pointsOfFunctionPlot[1_397].Y = 26.482

	// pointsOfFunctionPlot[1_398].X = 3.98
	// pointsOfFunctionPlot[1_398].Y = 26.749

	// pointsOfFunctionPlot[1_399].X = 3.99
	// pointsOfFunctionPlot[1_399].Y = 27.018

	// pointsOfFunctionPlot[1_400].X = 4.0
	// pointsOfFunctionPlot[1_400].Y = 27.289

	// pointsOfFunctionPlot[1_401].X = 4.01
	// pointsOfFunctionPlot[1_401].Y = 27.564

	// pointsOfFunctionPlot[1_402].X = 4.02
	// pointsOfFunctionPlot[1_402].Y = 27.841

	// pointsOfFunctionPlot[1_403].X = 4.03
	// pointsOfFunctionPlot[1_403].Y = 28.121

	// pointsOfFunctionPlot[1_404].X = 4.04
	// pointsOfFunctionPlot[1_404].Y = 28.404

	// pointsOfFunctionPlot[1_405].X = 4.05
	// pointsOfFunctionPlot[1_405].Y = 28.69

	// pointsOfFunctionPlot[1_406].X = 4.06
	// pointsOfFunctionPlot[1_406].Y = 28.978

	// pointsOfFunctionPlot[1_407].X = 4.07
	// pointsOfFunctionPlot[1_407].Y = 29.269

	// pointsOfFunctionPlot[1_408].X = 4.08
	// pointsOfFunctionPlot[1_408].Y = 29.564

	// pointsOfFunctionPlot[1_409].X = 4.09
	// pointsOfFunctionPlot[1_409].Y = 29.861

	// pointsOfFunctionPlot[1_410].X = 4.10
	// pointsOfFunctionPlot[1_410].Y = 30.161

	// pointsOfFunctionPlot[1_411].X = 4.11
	// pointsOfFunctionPlot[1_411].Y = 30.465

	// pointsOfFunctionPlot[1_412].X = 4.12
	// pointsOfFunctionPlot[1_412].Y = 30.771

	// pointsOfFunctionPlot[1_413].X = 4.13
	// pointsOfFunctionPlot[1_413].Y = 31.08

	// pointsOfFunctionPlot[1_414].X = 4.14
	// pointsOfFunctionPlot[1_414].Y = 31.393

	// pointsOfFunctionPlot[1_415].X = 4.15
	// pointsOfFunctionPlot[1_415].Y = 31.709

	// pointsOfFunctionPlot[1_416].X = 4.16
	// pointsOfFunctionPlot[1_416].Y = 32.027

	// pointsOfFunctionPlot[1_417].X = 4.17
	// pointsOfFunctionPlot[1_417].Y = 32.349

	// pointsOfFunctionPlot[1_418].X = 4.18
	// pointsOfFunctionPlot[1_418].Y = 32.675

	// pointsOfFunctionPlot[1_419].X = 4.19
	// pointsOfFunctionPlot[1_419].Y = 33.003

	// pointsOfFunctionPlot[1_420].X = 4.20
	// pointsOfFunctionPlot[1_420].Y = 33.335

	// pointsOfFunctionPlot[1_421].X = 4.21
	// pointsOfFunctionPlot[1_421].Y = 33.67

	// pointsOfFunctionPlot[1_422].X = 4.22
	// pointsOfFunctionPlot[1_422].Y = 34.009

	// pointsOfFunctionPlot[1_423].X = 4.23
	// pointsOfFunctionPlot[1_423].Y = 34.351

	// pointsOfFunctionPlot[1_424].X = 4.24
	// pointsOfFunctionPlot[1_424].Y = 34.696

	// pointsOfFunctionPlot[1_425].X = 4.25
	// pointsOfFunctionPlot[1_425].Y = 35.045

	// pointsOfFunctionPlot[1_426].X = 4.26
	// pointsOfFunctionPlot[1_426].Y = 35.397

	// pointsOfFunctionPlot[1_427].X = 4.27
	// pointsOfFunctionPlot[1_427].Y = 35.753

	// pointsOfFunctionPlot[1_428].X = 4.28
	// pointsOfFunctionPlot[1_428].Y = 36.113

	// pointsOfFunctionPlot[1_429].X = 4.29
	// pointsOfFunctionPlot[1_429].Y = 36.476

	// pointsOfFunctionPlot[1_430].X = 4.30
	// pointsOfFunctionPlot[1_430].Y = 36.843

	// pointsOfFunctionPlot[1_431].X = 4.31
	// pointsOfFunctionPlot[1_431].Y = 37.213

	// pointsOfFunctionPlot[1_432].X = 4.32
	// pointsOfFunctionPlot[1_432].Y = 37.587

	// pointsOfFunctionPlot[1_433].X = 4.33
	// pointsOfFunctionPlot[1_433].Y = 37.965

	// pointsOfFunctionPlot[1_434].X = 4.34
	// pointsOfFunctionPlot[1_434].Y = 38.347

	// pointsOfFunctionPlot[1_435].X = 4.35
	// pointsOfFunctionPlot[1_435].Y = 38.732

	// pointsOfFunctionPlot[1_436].X = 4.36
	// pointsOfFunctionPlot[1_436].Y = 39.122

	// pointsOfFunctionPlot[1_437].X = 4.37
	// pointsOfFunctionPlot[1_437].Y = 39.515

	// pointsOfFunctionPlot[1_438].X = 4.38
	// pointsOfFunctionPlot[1_438].Y = 39.912

	// pointsOfFunctionPlot[1_439].X = 4.39
	// pointsOfFunctionPlot[1_439].Y = 40.314

	// pointsOfFunctionPlot[1_440].X = 4.40
	// pointsOfFunctionPlot[1_440].Y = 40.719

	// pointsOfFunctionPlot[1_441].X = 4.41
	// pointsOfFunctionPlot[1_441].Y = 41.128

	// pointsOfFunctionPlot[1_442].X = 4.42
	// pointsOfFunctionPlot[1_442].Y = 41.542

	// pointsOfFunctionPlot[1_443].X = 4.43
	// pointsOfFunctionPlot[1_443].Y = 41.959

	// pointsOfFunctionPlot[1_444].X = 4.44
	// pointsOfFunctionPlot[1_444].Y = 42.381

	// pointsOfFunctionPlot[1_445].X = 4.45
	// pointsOfFunctionPlot[1_445].Y = 42.807

	// pointsOfFunctionPlot[1_446].X = 4.46
	// pointsOfFunctionPlot[1_446].Y = 43.237

	// pointsOfFunctionPlot[1_447].X = 4.47
	// pointsOfFunctionPlot[1_447].Y = 43.672

	// pointsOfFunctionPlot[1_448].X = 4.48
	// pointsOfFunctionPlot[1_448].Y = 44.111

	// pointsOfFunctionPlot[1_449].X = 4.49
	// pointsOfFunctionPlot[1_449].Y = 44.555

	// pointsOfFunctionPlot[1_450].X = 4.50
	// pointsOfFunctionPlot[1_450].Y = 45.003

	// pointsOfFunctionPlot[1_451].X = 4.51
	// pointsOfFunctionPlot[1_451].Y = 45.455

	// pointsOfFunctionPlot[1_452].X = 4.52
	// pointsOfFunctionPlot[1_452].Y = 45.912

	// pointsOfFunctionPlot[1_453].X = 4.53
	// pointsOfFunctionPlot[1_453].Y = 46.373

	// pointsOfFunctionPlot[1_454].X = 4.54
	// pointsOfFunctionPlot[1_454].Y = 46.84

	// pointsOfFunctionPlot[1_455].X = 4.55
	// pointsOfFunctionPlot[1_455].Y = 47.31

	// pointsOfFunctionPlot[1_456].X = 4.56
	// pointsOfFunctionPlot[1_456].Y = 47.765

	// pointsOfFunctionPlot[1_457].X = 4.57
	// pointsOfFunctionPlot[1_457].Y = 48.266

	// pointsOfFunctionPlot[1_458].X = 4.58
	// pointsOfFunctionPlot[1_458].Y = 48.752

	// pointsOfFunctionPlot[1_459].X = 4.59
	// pointsOfFunctionPlot[1_459].Y = 49.242

	// pointsOfFunctionPlot[1_460].X = 4.60
	// pointsOfFunctionPlot[1_460].Y = 49.737

	// pointsOfFunctionPlot[1_461].X = 4.61
	// pointsOfFunctionPlot[1_461].Y = 50.237

	// pointsOfFunctionPlot[1_462].X = 4.62
	// pointsOfFunctionPlot[1_462].Y = 50.742

	// pointsOfFunctionPlot[1_463].X = 4.63
	// pointsOfFunctionPlot[1_463].Y = 51.252

	// pointsOfFunctionPlot[1_464].X = 4.64
	// pointsOfFunctionPlot[1_464].Y = 51.767

	// pointsOfFunctionPlot[1_465].X = 4.65
	// pointsOfFunctionPlot[1_465].Y = 52.287

	// pointsOfFunctionPlot[1_466].X = 4.66
	// pointsOfFunctionPlot[1_466].Y = 52.813

	// pointsOfFunctionPlot[1_467].X = 4.67
	// pointsOfFunctionPlot[1_467].Y = 53.344

	// pointsOfFunctionPlot[1_468].X = 4.68
	// pointsOfFunctionPlot[1_468].Y = 53.88

	// pointsOfFunctionPlot[1_469].X = 4.69
	// pointsOfFunctionPlot[1_469].Y = 54.421

	// pointsOfFunctionPlot[1_470].X = 4.70
	// pointsOfFunctionPlot[1_470].Y = 54.969

	// pointsOfFunctionPlot[1_471].X = 4.71
	// pointsOfFunctionPlot[1_471].Y = 55.521

	// pointsOfFunctionPlot[1_472].X = 4.72
	// pointsOfFunctionPlot[1_472].Y = 56.079

	// pointsOfFunctionPlot[1_473].X = 4.73
	// pointsOfFunctionPlot[1_473].Y = 56.643

	// pointsOfFunctionPlot[1_474].X = 4.74
	// pointsOfFunctionPlot[1_474].Y = 57.212

	// pointsOfFunctionPlot[1_475].X = 4.75
	// pointsOfFunctionPlot[1_475].Y = 57.787

	// pointsOfFunctionPlot[1_476].X = 4.76
	// pointsOfFunctionPlot[1_476].Y = 58.368

	// pointsOfFunctionPlot[1_477].X = 4.77
	// pointsOfFunctionPlot[1_477].Y = 58.955

	// pointsOfFunctionPlot[1_478].X = 4.78
	// pointsOfFunctionPlot[1_478].Y = 59.547

	// pointsOfFunctionPlot[1_479].X = 4.79
	// pointsOfFunctionPlot[1_479].Y = 60.146

	// pointsOfFunctionPlot[1_480].X = 4.80
	// pointsOfFunctionPlot[1_480].Y = 60.751

	// pointsOfFunctionPlot[1_481].X = 4.81
	// pointsOfFunctionPlot[1_481].Y = 61.361

	// pointsOfFunctionPlot[1_482].X = 4.82
	// pointsOfFunctionPlot[1_482].Y = 61.978

	// pointsOfFunctionPlot[1_483].X = 4.83
	// pointsOfFunctionPlot[1_483].Y = 62.601

	// pointsOfFunctionPlot[1_484].X = 4.84
	// pointsOfFunctionPlot[1_484].Y = 63.23

	// pointsOfFunctionPlot[1_485].X = 4.85
	// pointsOfFunctionPlot[1_485].Y = 63.866

	// pointsOfFunctionPlot[1_486].X = 4.86
	// pointsOfFunctionPlot[1_486].Y = 64.508

	// pointsOfFunctionPlot[1_487].X = 4.87
	// pointsOfFunctionPlot[1_487].Y = 65.156

	// pointsOfFunctionPlot[1_488].X = 4.88
	// pointsOfFunctionPlot[1_488].Y = 65.811

	// pointsOfFunctionPlot[1_489].X = 4.89
	// pointsOfFunctionPlot[1_489].Y = 66.473

	// pointsOfFunctionPlot[1_490].X = 4.90
	// pointsOfFunctionPlot[1_490].Y = 67.141

	// pointsOfFunctionPlot[1_491].X = 4.91
	// pointsOfFunctionPlot[1_491].Y = 67.816

	// pointsOfFunctionPlot[1_492].X = 4.92
	// pointsOfFunctionPlot[1_492].Y = 68.497

	// pointsOfFunctionPlot[1_493].X = 4.93
	// pointsOfFunctionPlot[1_493].Y = 69.186

	// pointsOfFunctionPlot[1_494].X = 4.94
	// pointsOfFunctionPlot[1_494].Y = 69.881

	// pointsOfFunctionPlot[1_495].X = 4.95
	// pointsOfFunctionPlot[1_495].Y = 70.583

	// pointsOfFunctionPlot[1_496].X = 4.96
	// pointsOfFunctionPlot[1_496].Y = 71.293

	// pointsOfFunctionPlot[1_497].X = 4.97
	// pointsOfFunctionPlot[1_497].Y = 72.009

	// pointsOfFunctionPlot[1_498].X = 4.98
	// pointsOfFunctionPlot[1_498].Y = 72.733

	// pointsOfFunctionPlot[1_499].X = 4.99
	// pointsOfFunctionPlot[1_499].Y = 73.464

	// pointsOfFunctionPlot[1_500].X = 5.0
	// pointsOfFunctionPlot[1_500].Y = 74.203

	// pointsOfFunctionPlot[1_501].X = 5.01
	// pointsOfFunctionPlot[1_501].Y = 74.949

	// pointsOfFunctionPlot[1_502].X = 5.02
	// pointsOfFunctionPlot[1_502].Y = 75.702

	// pointsOfFunctionPlot[1_503].X = 5.03
	// pointsOfFunctionPlot[1_503].Y = 76.463

	// pointsOfFunctionPlot[1_504].X = 5.04
	// pointsOfFunctionPlot[1_504].Y = 77.231

	// pointsOfFunctionPlot[1_505].X = 5.05
	// pointsOfFunctionPlot[1_505].Y = 78.008

	// pointsOfFunctionPlot[1_506].X = 5.06
	// pointsOfFunctionPlot[1_506].Y = 78.792

	// pointsOfFunctionPlot[1_507].X = 5.07
	// pointsOfFunctionPlot[1_507].Y = 79.584

	// pointsOfFunctionPlot[1_508].X = 5.08
	// pointsOfFunctionPlot[1_508].Y = 80.383

	// pointsOfFunctionPlot[1_509].X = 5.09
	// pointsOfFunctionPlot[1_509].Y = 81.191

	// pointsOfFunctionPlot[1_510].X = 5.10
	// pointsOfFunctionPlot[1_510].Y = 82.007

	// pointsOfFunctionPlot[1_511].X = 5.11
	// pointsOfFunctionPlot[1_511].Y = 82.832

	// pointsOfFunctionPlot[1_512].X = 5.12
	// pointsOfFunctionPlot[1_512].Y = 83.664

	// pointsOfFunctionPlot[1_513].X = 5.13
	// pointsOfFunctionPlot[1_513].Y = 84.505

	// pointsOfFunctionPlot[1_514].X = 5.14
	// pointsOfFunctionPlot[1_514].Y = 85.354

	// pointsOfFunctionPlot[1_515].X = 5.15
	// pointsOfFunctionPlot[1_515].Y = 86.212

	// pointsOfFunctionPlot[1_516].X = 5.16
	// pointsOfFunctionPlot[1_516].Y = 87.079

	// pointsOfFunctionPlot[1_517].X = 5.17
	// pointsOfFunctionPlot[1_517].Y = 87.954

	// pointsOfFunctionPlot[1_518].X = 5.18
	// pointsOfFunctionPlot[1_518].Y = 88.838

	// pointsOfFunctionPlot[1_519].X = 5.19
	// pointsOfFunctionPlot[1_519].Y = 89.731

	// pointsOfFunctionPlot[1_520].X = 5.20
	// pointsOfFunctionPlot[1_520].Y = 90.633

	// pointsOfFunctionPlot[1_521].X = 5.21
	// pointsOfFunctionPlot[1_521].Y = 91.544

	// pointsOfFunctionPlot[1_522].X = 5.22
	// pointsOfFunctionPlot[1_522].Y = 92.464

	// pointsOfFunctionPlot[1_523].X = 5.23
	// pointsOfFunctionPlot[1_523].Y = 93.393

	// pointsOfFunctionPlot[1_524].X = 5.24
	// pointsOfFunctionPlot[1_524].Y = 94.332

	// pointsOfFunctionPlot[1_525].X = 5.25
	// pointsOfFunctionPlot[1_525].Y = 95.28

	// pointsOfFunctionPlot[1_526].X = 5.26
	// pointsOfFunctionPlot[1_526].Y = 96.238

	// pointsOfFunctionPlot[1_527].X = 5.27
	// pointsOfFunctionPlot[1_527].Y = 97.205

	// pointsOfFunctionPlot[1_528].X = 5.28
	// pointsOfFunctionPlot[1_528].Y = 98.182

	// pointsOfFunctionPlot[1_529].X = 5.29
	// pointsOfFunctionPlot[1_529].Y = 99.169

	// pointsOfFunctionPlot[1_530].X = 5.30
	// pointsOfFunctionPlot[1_530].Y = 100.165

	// pointsOfFunctionPlot[1_531].X = 5.31
	// pointsOfFunctionPlot[1_531].Y = 101.172

	// pointsOfFunctionPlot[1_532].X = 5.32
	// pointsOfFunctionPlot[1_532].Y = 102.189

	// pointsOfFunctionPlot[1_533].X = 5.33
	// pointsOfFunctionPlot[1_533].Y = 103.216

	// pointsOfFunctionPlot[1_534].X = 5.34
	// pointsOfFunctionPlot[1_534].Y = 104.253

	// pointsOfFunctionPlot[1_535].X = 5.35
	// pointsOfFunctionPlot[1_535].Y = 105.301

	// pointsOfFunctionPlot[1_536].X = 5.36
	// pointsOfFunctionPlot[1_536].Y = 106.36

	// pointsOfFunctionPlot[1_537].X = 5.37
	// pointsOfFunctionPlot[1_537].Y = 107.429

	// pointsOfFunctionPlot[1_538].X = 5.38
	// pointsOfFunctionPlot[1_538].Y = 108.508

	// pointsOfFunctionPlot[1_539].X = 5.39
	// pointsOfFunctionPlot[1_539].Y = 109.599

	// pointsOfFunctionPlot[1_540].X = 5.40
	// pointsOfFunctionPlot[1_540].Y = 110.7

	// pointsOfFunctionPlot[1_541].X = 5.41
	// pointsOfFunctionPlot[1_541].Y = 111.813

	// pointsOfFunctionPlot[1_542].X = 5.42
	// pointsOfFunctionPlot[1_542].Y = 112.937

	// pointsOfFunctionPlot[1_543].X = 5.43
	// pointsOfFunctionPlot[1_543].Y = 114.072

	// pointsOfFunctionPlot[1_544].X = 5.44
	// pointsOfFunctionPlot[1_544].Y = 115.218

	// pointsOfFunctionPlot[1_545].X = 5.45
	// pointsOfFunctionPlot[1_545].Y = 116.376

	// pointsOfFunctionPlot[1_546].X = 5.46
	// pointsOfFunctionPlot[1_546].Y = 117.546

	// pointsOfFunctionPlot[1_547].X = 5.47
	// pointsOfFunctionPlot[1_547].Y = 118.727

	// pointsOfFunctionPlot[1_548].X = 5.48
	// pointsOfFunctionPlot[1_548].Y = 119.921

	// pointsOfFunctionPlot[1_549].X = 5.49
	// pointsOfFunctionPlot[1_549].Y = 121.126

	// pointsOfFunctionPlot[1_550].X = 5.50
	// pointsOfFunctionPlot[1_550].Y = 122.343

	// pointsOfFunctionPlot[1_551].X = 5.51
	// pointsOfFunctionPlot[1_551].Y = 123.573

	// pointsOfFunctionPlot[1_552].X = 5.52
	// pointsOfFunctionPlot[1_552].Y = 124.815

	// pointsOfFunctionPlot[1_553].X = 5.53
	// pointsOfFunctionPlot[1_553].Y = 126.069

	// pointsOfFunctionPlot[1_554].X = 5.54
	// pointsOfFunctionPlot[1_554].Y = 127.337

	// pointsOfFunctionPlot[1_555].X = 5.55
	// pointsOfFunctionPlot[1_555].Y = 128.616

	// pointsOfFunctionPlot[1_556].X = 5.56
	// pointsOfFunctionPlot[1_556].Y = 129.909

	// pointsOfFunctionPlot[1_557].X = 5.57
	// pointsOfFunctionPlot[1_557].Y = 131.215

	// pointsOfFunctionPlot[1_558].X = 5.58
	// pointsOfFunctionPlot[1_558].Y = 132.533

	// pointsOfFunctionPlot[1_559].X = 5.59
	// pointsOfFunctionPlot[1_559].Y = 133.865

	// pointsOfFunctionPlot[1_560].X = 5.60
	// pointsOfFunctionPlot[1_560].Y = 135.211

	// pointsOfFunctionPlot[1_561].X = 5.61
	// pointsOfFunctionPlot[1_561].Y = 136.57

	// pointsOfFunctionPlot[1_562].X = 5.62
	// pointsOfFunctionPlot[1_562].Y = 137.942

	// pointsOfFunctionPlot[1_563].X = 5.63
	// pointsOfFunctionPlot[1_563].Y = 139.329

	// pointsOfFunctionPlot[1_564].X = 5.64
	// pointsOfFunctionPlot[1_564].Y = 140.729

	// pointsOfFunctionPlot[1_565].X = 5.65
	// pointsOfFunctionPlot[1_565].Y = 142.143

	// pointsOfFunctionPlot[1_566].X = 5.66
	// pointsOfFunctionPlot[1_566].Y = 143.572

	// pointsOfFunctionPlot[1_567].X = 5.67
	// pointsOfFunctionPlot[1_567].Y = 145.015

	// pointsOfFunctionPlot[1_568].X = 5.68
	// pointsOfFunctionPlot[1_568].Y = 146.473

	// pointsOfFunctionPlot[1_569].X = 5.69
	// pointsOfFunctionPlot[1_569].Y = 147.945

	// pointsOfFunctionPlot[1_570].X = 5.70
	// pointsOfFunctionPlot[1_570].Y = 149.432

	// pointsOfFunctionPlot[1_571].X = 5.71
	// pointsOfFunctionPlot[1_571].Y = 150.933

	// pointsOfFunctionPlot[1_572].X = 5.72
	// pointsOfFunctionPlot[1_572].Y = 152.45

	// pointsOfFunctionPlot[1_573].X = 5.73
	// pointsOfFunctionPlot[1_573].Y = 153.983

	// pointsOfFunctionPlot[1_574].X = 5.74
	// pointsOfFunctionPlot[1_574].Y = 155.53

	// pointsOfFunctionPlot[1_575].X = 5.75
	// pointsOfFunctionPlot[1_575].Y = 157.093

	// pointsOfFunctionPlot[1_576].X = 5.76
	// pointsOfFunctionPlot[1_576].Y = 158.672

	// pointsOfFunctionPlot[1_577].X = 5.77
	// pointsOfFunctionPlot[1_577].Y = 160.267

	// pointsOfFunctionPlot[1_578].X = 5.78
	// pointsOfFunctionPlot[1_578].Y = 161.878

	// pointsOfFunctionPlot[1_579].X = 5.79
	// pointsOfFunctionPlot[1_579].Y = 163.504

	// pointsOfFunctionPlot[1_580].X = 5.80
	// pointsOfFunctionPlot[1_580].Y = 165.148

	// pointsOfFunctionPlot[1_581].X = 5.81
	// pointsOfFunctionPlot[1_581].Y = 166.484

	// pointsOfFunctionPlot[1_582].X = 5.82
	// pointsOfFunctionPlot[1_582].Y = 168.484

	// pointsOfFunctionPlot[1_583].X = 5.83
	// pointsOfFunctionPlot[1_583].Y = 170.177

	// pointsOfFunctionPlot[1_584].X = 5.84
	// pointsOfFunctionPlot[1_584].Y = 171.888

	// pointsOfFunctionPlot[1_585].X = 5.85
	// pointsOfFunctionPlot[1_585].Y = 173.615

	// pointsOfFunctionPlot[1_586].X = 5.86
	// pointsOfFunctionPlot[1_586].Y = 175.36

	// pointsOfFunctionPlot[1_587].X = 5.87
	// pointsOfFunctionPlot[1_587].Y = 177.123

	// pointsOfFunctionPlot[1_588].X = 5.88
	// pointsOfFunctionPlot[1_588].Y = 178.903

	// pointsOfFunctionPlot[1_589].X = 5.89
	// pointsOfFunctionPlot[1_589].Y = 180.701

	// pointsOfFunctionPlot[1_590].X = 5.90
	// pointsOfFunctionPlot[1_590].Y = 182.517

	// pointsOfFunctionPlot[1_591].X = 5.91
	// pointsOfFunctionPlot[1_591].Y = 184.351

	// pointsOfFunctionPlot[1_592].X = 5.92
	// pointsOfFunctionPlot[1_592].Y = 186.204

	// pointsOfFunctionPlot[1_593].X = 5.93
	// pointsOfFunctionPlot[1_593].Y = 188.075

	// pointsOfFunctionPlot[1_594].X = 5.94
	// pointsOfFunctionPlot[1_594].Y = 189.966

	// pointsOfFunctionPlot[1_595].X = 5.95
	// pointsOfFunctionPlot[1_595].Y = 191.875

	// pointsOfFunctionPlot[1_596].X = 5.96
	// pointsOfFunctionPlot[1_596].Y = 193.803

	// pointsOfFunctionPlot[1_597].X = 5.97
	// pointsOfFunctionPlot[1_597].Y = 195.751

	// pointsOfFunctionPlot[1_598].X = 5.98
	// pointsOfFunctionPlot[1_598].Y = 197.718

	// pointsOfFunctionPlot[1_599].X = 5.99
	// pointsOfFunctionPlot[1_599].Y = 199.706

	// pointsOfFunctionPlot[1_600].X = 6.0
	// pointsOfFunctionPlot[1_600].Y = 201.713

	// pointsOfFunctionPlot[1_601].X = 6.01
	// pointsOfFunctionPlot[1_601].Y = 203.74

	// pointsOfFunctionPlot[1_602].X = 6.02
	// pointsOfFunctionPlot[1_602].Y = 205.788

	// pointsOfFunctionPlot[1_603].X = 6.03
	// pointsOfFunctionPlot[1_603].Y = 207.856

	// pointsOfFunctionPlot[1_604].X = 6.04
	// pointsOfFunctionPlot[1_604].Y = 209.945

	// pointsOfFunctionPlot[1_605].X = 6.05
	// pointsOfFunctionPlot[1_605].Y = 213.055

	// pointsOfFunctionPlot[1_606].X = 6.06
	// pointsOfFunctionPlot[1_606].Y = 214.186

	// pointsOfFunctionPlot[1_607].X = 6.07
	// pointsOfFunctionPlot[1_607].Y = 216.339

	// pointsOfFunctionPlot[1_608].X = 6.08
	// pointsOfFunctionPlot[1_608].Y = 218.513

	// pointsOfFunctionPlot[1_609].X = 6.09
	// pointsOfFunctionPlot[1_609].Y = 220.709

	// pointsOfFunctionPlot[1_610].X = 6.10
	// pointsOfFunctionPlot[1_610].Y = 222.927

	// pointsOfFunctionPlot[1_611].X = 6.11
	// pointsOfFunctionPlot[1_611].Y = 225.168

	// pointsOfFunctionPlot[1_612].X = 6.12
	// pointsOfFunctionPlot[1_612].Y = 227.431

	// pointsOfFunctionPlot[1_613].X = 6.13
	// pointsOfFunctionPlot[1_613].Y = 229.716

	// pointsOfFunctionPlot[1_614].X = 6.14
	// pointsOfFunctionPlot[1_614].Y = 232.025

	// pointsOfFunctionPlot[1_615].X = 6.15
	// pointsOfFunctionPlot[1_615].Y = 234.357

	// pointsOfFunctionPlot[1_616].X = 6.16
	// pointsOfFunctionPlot[1_616].Y = 236.712

	// pointsOfFunctionPlot[1_617].X = 6.17
	// pointsOfFunctionPlot[1_617].Y = 239.092

	// pointsOfFunctionPlot[1_618].X = 6.18
	// pointsOfFunctionPlot[1_618].Y = 241.494

	// pointsOfFunctionPlot[1_619].X = 6.19
	// pointsOfFunctionPlot[1_619].Y = 243.992

	// pointsOfFunctionPlot[1_620].X = 6.20
	// pointsOfFunctionPlot[1_620].Y = 246.373

	// pointsOfFunctionPlot[1_621].X = 6.21
	// pointsOfFunctionPlot[1_621].Y = 248.849

	// pointsOfFunctionPlot[1_622].X = 6.22
	// pointsOfFunctionPlot[1_622].Y = 251.35

	// pointsOfFunctionPlot[1_623].X = 6.23
	// pointsOfFunctionPlot[1_623].Y = 253.876

	// pointsOfFunctionPlot[1_624].X = 6.24
	// pointsOfFunctionPlot[1_624].Y = 256.428

	// pointsOfFunctionPlot[1_625].X = 6.25
	// pointsOfFunctionPlot[1_625].Y = 259.005

	// pointsOfFunctionPlot[1_626].X = 6.26
	// pointsOfFunctionPlot[1_626].Y = 261.608

	// pointsOfFunctionPlot[1_627].X = 6.27
	// pointsOfFunctionPlot[1_627].Y = 264.237

	// pointsOfFunctionPlot[1_628].X = 6.28
	// pointsOfFunctionPlot[1_628].Y = 266.893

	// pointsOfFunctionPlot[1_629].X = 6.29
	// pointsOfFunctionPlot[1_629].Y = 269.575

	// pointsOfFunctionPlot[1_630].X = 6.30
	// pointsOfFunctionPlot[1_630].Y = 272.285

	// pointsOfFunctionPlot[1_631].X = 6.31
	// pointsOfFunctionPlot[1_631].Y = 275.021

	// pointsOfFunctionPlot[1_632].X = 6.32
	// pointsOfFunctionPlot[1_632].Y = 277.785

	// pointsOfFunctionPlot[1_633].X = 6.33
	// pointsOfFunctionPlot[1_633].Y = 280.577

	// pointsOfFunctionPlot[1_634].X = 6.34
	// pointsOfFunctionPlot[1_634].Y = 283.397

	// pointsOfFunctionPlot[1_635].X = 6.35
	// pointsOfFunctionPlot[1_635].Y = 286.245

	// pointsOfFunctionPlot[1_636].X = 6.36
	// pointsOfFunctionPlot[1_636].Y = 289.122

	// pointsOfFunctionPlot[1_637].X = 6.37
	// pointsOfFunctionPlot[1_637].Y = 292.028

	// pointsOfFunctionPlot[1_638].X = 6.38
	// pointsOfFunctionPlot[1_638].Y = 294.963

	// pointsOfFunctionPlot[1_639].X = 6.39
	// pointsOfFunctionPlot[1_639].Y = 297.927

	// pointsOfFunctionPlot[1_640].X = 6.40
	// pointsOfFunctionPlot[1_640].Y = 300.921

	// pointsOfFunctionPlot[1_641].X = 6.41
	// pointsOfFunctionPlot[1_641].Y = 303.946

	// pointsOfFunctionPlot[1_642].X = 6.42
	// pointsOfFunctionPlot[1_642].Y = 307.0

	// pointsOfFunctionPlot[1_643].X = 6.43
	// pointsOfFunctionPlot[1_643].Y = 310.086

	// pointsOfFunctionPlot[1_644].X = 6.44
	// pointsOfFunctionPlot[1_644].Y = 313.202

	// pointsOfFunctionPlot[1_645].X = 6.45
	// pointsOfFunctionPlot[1_645].Y = 316.35

	// pointsOfFunctionPlot[1_646].X = 6.46
	// pointsOfFunctionPlot[1_646].Y = 319.529

	// pointsOfFunctionPlot[1_647].X = 6.47
	// pointsOfFunctionPlot[1_647].Y = 322.741

	// pointsOfFunctionPlot[1_648].X = 6.48
	// pointsOfFunctionPlot[1_648].Y = 325.984

	// pointsOfFunctionPlot[1_649].X = 6.49
	// pointsOfFunctionPlot[1_649].Y = 329.26

	// pointsOfFunctionPlot[1_650].X = 6.50
	// pointsOfFunctionPlot[1_650].Y = 332.57

	// pointsOfFunctionPlot[1_651].X = 6.51
	// pointsOfFunctionPlot[1_651].Y = 335.912

	// pointsOfFunctionPlot[1_652].X = 6.52
	// pointsOfFunctionPlot[1_652].Y = 339.288

	// pointsOfFunctionPlot[1_653].X = 6.53
	// pointsOfFunctionPlot[1_653].Y = 342.698

	// pointsOfFunctionPlot[1_654].X = 6.54
	// pointsOfFunctionPlot[1_654].Y = 346.142

	// pointsOfFunctionPlot[1_655].X = 6.55
	// pointsOfFunctionPlot[1_655].Y = 349.621

	// pointsOfFunctionPlot[1_656].X = 6.56
	// pointsOfFunctionPlot[1_656].Y = 353.135

	// pointsOfFunctionPlot[1_657].X = 6.57
	// pointsOfFunctionPlot[1_657].Y = 356.684

	// pointsOfFunctionPlot[1_658].X = 6.58
	// pointsOfFunctionPlot[1_658].Y = 360.268

	// pointsOfFunctionPlot[1_659].X = 6.59
	// pointsOfFunctionPlot[1_659].Y = 363.889

	// pointsOfFunctionPlot[1_660].X = 6.60
	// pointsOfFunctionPlot[1_660].Y = 367.546

	// pointsOfFunctionPlot[1_661].X = 6.61
	// pointsOfFunctionPlot[1_661].Y = 371.24

	// pointsOfFunctionPlot[1_662].X = 6.62
	// pointsOfFunctionPlot[1_662].Y = 374.971

	// pointsOfFunctionPlot[1_663].X = 6.63
	// pointsOfFunctionPlot[1_663].Y = 378.74

	// pointsOfFunctionPlot[1_664].X = 6.64
	// pointsOfFunctionPlot[1_664].Y = 382.546

	// pointsOfFunctionPlot[1_665].X = 6.65
	// pointsOfFunctionPlot[1_665].Y = 386.391

	// pointsOfFunctionPlot[1_666].X = 6.66
	// pointsOfFunctionPlot[1_666].Y = 390.274

	// pointsOfFunctionPlot[1_667].X = 6.67
	// pointsOfFunctionPlot[1_667].Y = 394.197

	// pointsOfFunctionPlot[1_668].X = 6.68
	// pointsOfFunctionPlot[1_668].Y = 398.158

	// pointsOfFunctionPlot[1_669].X = 6.69
	// pointsOfFunctionPlot[1_669].Y = 402.16

	// pointsOfFunctionPlot[1_670].X = 6.70
	// pointsOfFunctionPlot[1_670].Y = 406.202

	// pointsOfFunctionPlot[1_671].X = 6.71
	// pointsOfFunctionPlot[1_671].Y = 410.284

	// pointsOfFunctionPlot[1_672].X = 6.72
	// pointsOfFunctionPlot[1_672].Y = 414.408

	// pointsOfFunctionPlot[1_673].X = 6.73
	// pointsOfFunctionPlot[1_673].Y = 418.573

	// pointsOfFunctionPlot[1_674].X = 6.74
	// pointsOfFunctionPlot[1_674].Y = 422.779

	// pointsOfFunctionPlot[1_675].X = 6.75
	// pointsOfFunctionPlot[1_675].Y = 427.028

	// pointsOfFunctionPlot[1_676].X = 6.76
	// pointsOfFunctionPlot[1_676].Y = 431.32

	// pointsOfFunctionPlot[1_677].X = 6.77
	// pointsOfFunctionPlot[1_677].Y = 435.655

	// pointsOfFunctionPlot[1_678].X = 6.78
	// pointsOfFunctionPlot[1_678].Y = 440.033

	// pointsOfFunctionPlot[1_679].X = 6.79
	// pointsOfFunctionPlot[1_679].Y = 444.456

	// pointsOfFunctionPlot[1_680].X = 6.80
	// pointsOfFunctionPlot[1_680].Y = 448.923

	// pointsOfFunctionPlot[1_681].X = 6.81
	// pointsOfFunctionPlot[1_681].Y = 453.434

	// pointsOfFunctionPlot[1_682].X = 6.82
	// pointsOfFunctionPlot[1_682].Y = 457.991

	// pointsOfFunctionPlot[1_683].X = 6.83
	// pointsOfFunctionPlot[1_683].Y = 462.594

	// pointsOfFunctionPlot[1_684].X = 6.84
	// pointsOfFunctionPlot[1_684].Y = 467.244

	// pointsOfFunctionPlot[1_685].X = 6.85
	// pointsOfFunctionPlot[1_685].Y = 471.939

	// pointsOfFunctionPlot[1_686].X = 6.86
	// pointsOfFunctionPlot[1_686].Y = 476.683

	// pointsOfFunctionPlot[1_687].X = 6.87
	// pointsOfFunctionPlot[1_687].Y = 481.312

	// pointsOfFunctionPlot[1_688].X = 6.88
	// pointsOfFunctionPlot[1_688].Y = 491.2

	// pointsOfFunctionPlot[1_689].X = 6.89
	// pointsOfFunctionPlot[1_689].Y = 496.136

	// pointsOfFunctionPlot[1_690].X = 6.90
	// pointsOfFunctionPlot[1_690].Y = 496.136

	// pointsOfFunctionPlot[1_691].X = 6.91
	// pointsOfFunctionPlot[1_691].Y = 501.123

	// pointsOfFunctionPlot[1_692].X = 6.92
	// pointsOfFunctionPlot[1_692].Y = 506.159

	// pointsOfFunctionPlot[1_693].X = 6.93
	// pointsOfFunctionPlot[1_693].Y = 511.246

	// pointsOfFunctionPlot[1_694].X = 6.94
	// pointsOfFunctionPlot[1_694].Y = 516.384

	// pointsOfFunctionPlot[1_695].X = 6.95
	// pointsOfFunctionPlot[1_695].Y = 521.574

	// pointsOfFunctionPlot[1_696].X = 6.96
	// pointsOfFunctionPlot[1_696].Y = 526.816

	// pointsOfFunctionPlot[1_697].X = 6.97
	// pointsOfFunctionPlot[1_697].Y = 532.11

	// pointsOfFunctionPlot[1_698].X = 6.98
	// pointsOfFunctionPlot[1_698].Y = 537.458

	// pointsOfFunctionPlot[1_699].X = 6.99
	// pointsOfFunctionPlot[1_699].Y = 542.86

	// pointsOfFunctionPlot[1_700].X = 7.0
	// pointsOfFunctionPlot[1_700].Y = 548.316

	// pointsOfFunctionPlot[1_701].X = 7.01
	// pointsOfFunctionPlot[1_701].Y = 553.826

	// pointsOfFunctionPlot[1_702].X = 7.02
	// pointsOfFunctionPlot[1_702].Y = 559.392

	// pointsOfFunctionPlot[1_703].X = 7.03
	// pointsOfFunctionPlot[1_703].Y = 565.014

	// pointsOfFunctionPlot[1_704].X = 7.04
	// pointsOfFunctionPlot[1_704].Y = 570.693

	// pointsOfFunctionPlot[1_705].X = 7.05
	// pointsOfFunctionPlot[1_705].Y = 576.428

	// pointsOfFunctionPlot[1_706].X = 7.06
	// pointsOfFunctionPlot[1_706].Y = 582.222

	// pointsOfFunctionPlot[1_707].X = 7.07
	// pointsOfFunctionPlot[1_707].Y = 588.073

	// pointsOfFunctionPlot[1_708].X = 7.08
	// pointsOfFunctionPlot[1_708].Y = 593.983

	// pointsOfFunctionPlot[1_709].X = 7.09
	// pointsOfFunctionPlot[1_709].Y = 599.953

	// pointsOfFunctionPlot[1_710].X = 7.10
	// pointsOfFunctionPlot[1_710].Y = 605.983

	// pointsOfFunctionPlot[1_711].X = 7.11
	// pointsOfFunctionPlot[1_711].Y = 612.073

	// pointsOfFunctionPlot[1_712].X = 7.12
	// pointsOfFunctionPlot[1_712].Y = 618.224

	// pointsOfFunctionPlot[1_713].X = 7.13
	// pointsOfFunctionPlot[1_713].Y = 624.438

	// pointsOfFunctionPlot[1_714].X = 7.14
	// pointsOfFunctionPlot[1_714].Y = 630.713

	// pointsOfFunctionPlot[1_715].X = 7.15
	// pointsOfFunctionPlot[1_715].Y = 637.052

	// pointsOfFunctionPlot[1_716].X = 7.16
	// pointsOfFunctionPlot[1_716].Y = 643.455

	// pointsOfFunctionPlot[1_717].X = 7.17
	// pointsOfFunctionPlot[1_717].Y = 649.921

	// pointsOfFunctionPlot[1_718].X = 7.18
	// pointsOfFunctionPlot[1_718].Y = 656.453

	// pointsOfFunctionPlot[1_719].X = 7.19
	// pointsOfFunctionPlot[1_719].Y = 663.051

	// pointsOfFunctionPlot[1_720].X = 7.20
	// pointsOfFunctionPlot[1_720].Y = 669.715

	// pointsOfFunctionPlot[1_721].X = 7.21
	// pointsOfFunctionPlot[1_721].Y = 676.445

	// pointsOfFunctionPlot[1_722].X = 7.22
	// pointsOfFunctionPlot[1_722].Y = 683.244

	// pointsOfFunctionPlot[1_723].X = 7.23
	// pointsOfFunctionPlot[1_723].Y = 690.11

	// pointsOfFunctionPlot[1_724].X = 7.24
	// pointsOfFunctionPlot[1_724].Y = 697.046

	// pointsOfFunctionPlot[1_725].X = 7.25
	// pointsOfFunctionPlot[1_725].Y = 704.052

	// pointsOfFunctionPlot[1_726].X = 7.26
	// pointsOfFunctionPlot[1_726].Y = 711.127

	// pointsOfFunctionPlot[1_727].X = 7.27
	// pointsOfFunctionPlot[1_727].Y = 718.274

	// pointsOfFunctionPlot[1_728].X = 7.28
	// pointsOfFunctionPlot[1_728].Y = 725.493

	// pointsOfFunctionPlot[1_729].X = 7.29
	// pointsOfFunctionPlot[1_729].Y = 732.785

	// pointsOfFunctionPlot[1_730].X = 7.30
	// pointsOfFunctionPlot[1_730].Y = 740.149

	// pointsOfFunctionPlot[1_731].X = 7.31
	// pointsOfFunctionPlot[1_731].Y = 747.588

	// pointsOfFunctionPlot[1_732].X = 7.32
	// pointsOfFunctionPlot[1_732].Y = 755.101

	// pointsOfFunctionPlot[1_733].X = 7.33
	// pointsOfFunctionPlot[1_733].Y = 762.69

	// pointsOfFunctionPlot[1_734].X = 7.34
	// pointsOfFunctionPlot[1_734].Y = 770.355

	// pointsOfFunctionPlot[1_735].X = 7.35
	// pointsOfFunctionPlot[1_735].Y = 778.097

	// pointsOfFunctionPlot[1_736].X = 7.36
	// pointsOfFunctionPlot[1_736].Y = 785.917

	// pointsOfFunctionPlot[1_737].X = 7.37
	// pointsOfFunctionPlot[1_737].Y = 793.816

	// pointsOfFunctionPlot[1_738].X = 7.38
	// pointsOfFunctionPlot[1_738].Y = 801.794

	// pointsOfFunctionPlot[1_739].X = 7.39
	// pointsOfFunctionPlot[1_739].Y = 809.852

	// pointsOfFunctionPlot[1_740].X = 7.40
	// pointsOfFunctionPlot[1_740].Y = 817.991

	// pointsOfFunctionPlot[1_741].X = 7.41
	// pointsOfFunctionPlot[1_741].Y = 826.212

	// pointsOfFunctionPlot[1_742].X = 7.42
	// pointsOfFunctionPlot[1_742].Y = 834.516

	// pointsOfFunctionPlot[1_743].X = 7.43
	// pointsOfFunctionPlot[1_743].Y = 842.903

	// pointsOfFunctionPlot[1_744].X = 7.44
	// pointsOfFunctionPlot[1_744].Y = 851.374

	// pointsOfFunctionPlot[1_745].X = 7.45
	// pointsOfFunctionPlot[1_745].Y = 859.931

	// pointsOfFunctionPlot[1_746].X = 7.46
	// pointsOfFunctionPlot[1_746].Y = 868.573

	// pointsOfFunctionPlot[1_747].X = 7.47
	// pointsOfFunctionPlot[1_747].Y = 877.303

	// pointsOfFunctionPlot[1_748].X = 7.48
	// pointsOfFunctionPlot[1_748].Y = 886.12

	// pointsOfFunctionPlot[1_749].X = 7.49
	// pointsOfFunctionPlot[1_749].Y = 895.025

	// pointsOfFunctionPlot[1_750].X = 7.50
	// pointsOfFunctionPlot[1_750].Y = 904.02

	// pointsOfFunctionPlot[1_751].X = 7.51
	// pointsOfFunctionPlot[1_751].Y = 913.106

	// pointsOfFunctionPlot[1_752].X = 7.52
	// pointsOfFunctionPlot[1_752].Y = 922.283

	// pointsOfFunctionPlot[1_753].X = 7.53
	// pointsOfFunctionPlot[1_753].Y = 931.552

	// pointsOfFunctionPlot[1_754].X = 7.54
	// pointsOfFunctionPlot[1_754].Y = 940.914

	// pointsOfFunctionPlot[1_755].X = 7.55
	// pointsOfFunctionPlot[1_755].Y = 950.371

	// pointsOfFunctionPlot[1_756].X = 7.56
	// pointsOfFunctionPlot[1_756].Y = 959.922

	// pointsOfFunctionPlot[1_757].X = 7.57
	// pointsOfFunctionPlot[1_757].Y = 969.569

	// pointsOfFunctionPlot[1_758].X = 7.58
	// pointsOfFunctionPlot[1_758].Y = 979.314

	// pointsOfFunctionPlot[1_759].X = 7.59
	// pointsOfFunctionPlot[1_759].Y = 989.156

	// pointsOfFunctionPlot[1_760].X = 7.60
	// pointsOfFunctionPlot[1_760].Y = 999.097

	// pointsOfFunctionPlot[1_761].X = 7.61
	// pointsOfFunctionPlot[1_761].Y = 1_009.138

	// pointsOfFunctionPlot[1_762].X = 7.62
	// pointsOfFunctionPlot[1_762].Y = 1_019.28

	// pointsOfFunctionPlot[1_763].X = 7.63
	// pointsOfFunctionPlot[1_763].Y = 1_029.524

	// pointsOfFunctionPlot[1_764].X = 7.64
	// pointsOfFunctionPlot[1_764].Y = 1_039.871

	// pointsOfFunctionPlot[1_765].X = 7.65
	// pointsOfFunctionPlot[1_765].Y = 1_050.322

	// pointsOfFunctionPlot[1_766].X = 7.66
	// pointsOfFunctionPlot[1_766].Y = 1_060.878

	// pointsOfFunctionPlot[1_767].X = 7.67
	// pointsOfFunctionPlot[1_767].Y = 1_071.54

	// pointsOfFunctionPlot[1_768].X = 7.68
	// pointsOfFunctionPlot[1_768].Y = 1_082.309

	// pointsOfFunctionPlot[1_769].X = 7.69
	// pointsOfFunctionPlot[1_769].Y = 1_093.187

	// pointsOfFunctionPlot[1_770].X = 7.70
	// pointsOfFunctionPlot[1_770].Y = 1_104.173

	// pointsOfFunctionPlot[1_771].X = 7.71
	// pointsOfFunctionPlot[1_771].Y = 1_115.27

	// pointsOfFunctionPlot[1_772].X = 7.72
	// pointsOfFunctionPlot[1_772].Y = 1_126.479

	// pointsOfFunctionPlot[1_773].X = 7.73
	// pointsOfFunctionPlot[1_773].Y = 1_137.8

	// pointsOfFunctionPlot[1_774].X = 7.74
	// pointsOfFunctionPlot[1_774].Y = 1_149.235

	// pointsOfFunctionPlot[1_775].X = 7.75
	// pointsOfFunctionPlot[1_775].Y = 1_160.785

	// pointsOfFunctionPlot[1_776].X = 7.76
	// pointsOfFunctionPlot[1_776].Y = 1_172.452

	// pointsOfFunctionPlot[1_777].X = 7.77
	// pointsOfFunctionPlot[1_777].Y = 1_184.235

	// pointsOfFunctionPlot[1_778].X = 7.78
	// pointsOfFunctionPlot[1_778].Y = 1_196.137

	// pointsOfFunctionPlot[1_779].X = 7.79
	// pointsOfFunctionPlot[1_779].Y = 1_208.158

	// pointsOfFunctionPlot[1_780].X = 7.80
	// pointsOfFunctionPlot[1_780].Y = 1_220.3

	// pointsOfFunctionPlot[1_781].X = 7.81
	// pointsOfFunctionPlot[1_781].Y = 1_232.565

	// pointsOfFunctionPlot[1_782].X = 7.82
	// pointsOfFunctionPlot[1_782].Y = 1_244.952

	// pointsOfFunctionPlot[1_783].X = 7.83
	// pointsOfFunctionPlot[1_783].Y = 1_257.464

	// pointsOfFunctionPlot[1_784].X = 7.84
	// pointsOfFunctionPlot[1_784].Y = 1_270.102

	// pointsOfFunctionPlot[1_785].X = 7.85
	// pointsOfFunctionPlot[1_785].Y = 1_282.866

	// pointsOfFunctionPlot[1_786].X = 7.86
	// pointsOfFunctionPlot[1_786].Y = 1_295.759

	// pointsOfFunctionPlot[1_787].X = 7.87
	// pointsOfFunctionPlot[1_787].Y = 1_308.782

	// pointsOfFunctionPlot[1_788].X = 7.88
	// pointsOfFunctionPlot[1_788].Y = 1_321.936

	// pointsOfFunctionPlot[1_789].X = 7.89
	// pointsOfFunctionPlot[1_789].Y = 1_335.221

	// pointsOfFunctionPlot[1_790].X = 7.90
	// pointsOfFunctionPlot[1_790].Y = 1_348.64

	// pointsOfFunctionPlot[1_791].X = 7.91
	// pointsOfFunctionPlot[1_791].Y = 1_362.195

	// pointsOfFunctionPlot[1_792].X = 7.92
	// pointsOfFunctionPlot[1_792].Y = 1_375.885

	// pointsOfFunctionPlot[1_793].X = 7.93
	// pointsOfFunctionPlot[1_793].Y = 1_389.713

	// pointsOfFunctionPlot[1_794].X = 7.94
	// pointsOfFunctionPlot[1_794].Y = 1_403.68

	// pointsOfFunctionPlot[1_795].X = 7.95
	// pointsOfFunctionPlot[1_795].Y = 1_417.787

	// pointsOfFunctionPlot[1_796].X = 7.96
	// pointsOfFunctionPlot[1_796].Y = 1_432.036

	// pointsOfFunctionPlot[1_797].X = 7.97
	// pointsOfFunctionPlot[1_797].Y = 1_446.428

	// pointsOfFunctionPlot[1_798].X = 7.98
	// pointsOfFunctionPlot[1_798].Y = 1_460.965

	// pointsOfFunctionPlot[1_799].X = 7.99
	// pointsOfFunctionPlot[1_799].Y = 1_475.648

	// pointsOfFunctionPlot[1_800].X = 8.0
	// pointsOfFunctionPlot[1_800].Y = 1_490.478

	// pointsOfFunctionPlot[1_801].X = 8.01
	// pointsOfFunctionPlot[1_801].Y = 1_505.458

	// pointsOfFunctionPlot[1_802].X = 8.02
	// pointsOfFunctionPlot[1_802].Y = 1_520.588

	// pointsOfFunctionPlot[1_803].X = 8.03
	// pointsOfFunctionPlot[1_803].Y = 1_535.87

	// pointsOfFunctionPlot[1_804].X = 8.04
	// pointsOfFunctionPlot[1_804].Y = 1_551.306

	// pointsOfFunctionPlot[1_805].X = 8.05
	// pointsOfFunctionPlot[1_805].Y = 1_566.897

	// pointsOfFunctionPlot[1_806].X = 8.06
	// pointsOfFunctionPlot[1_806].Y = 1_582.644

	// pointsOfFunctionPlot[1_807].X = 8.07
	// pointsOfFunctionPlot[1_807].Y = 1_598.55

	// pointsOfFunctionPlot[1_808].X = 8.08
	// pointsOfFunctionPlot[1_808].Y = 1_614.616

	// pointsOfFunctionPlot[1_809].X = 8.09
	// pointsOfFunctionPlot[1_809].Y = 1_630.843

	// pointsOfFunctionPlot[1_810].X = 8.10
	// pointsOfFunctionPlot[1_810].Y = 1_647.233

	// pointsOfFunctionPlot[1_811].X = 8.11
	// pointsOfFunctionPlot[1_811].Y = 1_663.788

	// pointsOfFunctionPlot[1_812].X = 8.12
	// pointsOfFunctionPlot[1_812].Y = 1_680.51

	// pointsOfFunctionPlot[1_813].X = 8.13
	// pointsOfFunctionPlot[1_813].Y = 1_697.399

	// pointsOfFunctionPlot[1_814].X = 8.14
	// pointsOfFunctionPlot[1_814].Y = 1_714.458

	// pointsOfFunctionPlot[1_815].X = 8.15
	// pointsOfFunctionPlot[1_815].Y = 1_731.689

	// pointsOfFunctionPlot[1_816].X = 8.16
	// pointsOfFunctionPlot[1_816].Y = 1_749.093

	// pointsOfFunctionPlot[1_817].X = 8.17
	// pointsOfFunctionPlot[1_817].Y = 1_766.671

	// pointsOfFunctionPlot[1_818].X = 8.18
	// pointsOfFunctionPlot[1_818].Y = 1_784.427

	// pointsOfFunctionPlot[1_819].X = 8.19
	// pointsOfFunctionPlot[1_819].Y = 1_802.36

	// pointsOfFunctionPlot[1_820].X = 8.20
	// pointsOfFunctionPlot[1_820].Y = 1_820.475

	// pointsOfFunctionPlot[1_821].X = 8.21
	// pointsOfFunctionPlot[1_821].Y = 1_838.771

	// pointsOfFunctionPlot[1_822].X = 8.22
	// pointsOfFunctionPlot[1_822].Y = 1_857.251

	// pointsOfFunctionPlot[1_823].X = 8.23
	// pointsOfFunctionPlot[1_823].Y = 1_875.916

	// pointsOfFunctionPlot[1_824].X = 8.24
	// pointsOfFunctionPlot[1_824].Y = 1_894.77

	// pointsOfFunctionPlot[1_825].X = 8.25
	// pointsOfFunctionPlot[1_825].Y = 1_913.812

	// pointsOfFunctionPlot[1_826].X = 8.26
	// pointsOfFunctionPlot[1_826].Y = 1_933.046

	// pointsOfFunctionPlot[1_827].X = 8.27
	// pointsOfFunctionPlot[1_827].Y = 1_952.474

	// pointsOfFunctionPlot[1_828].X = 8.28
	// pointsOfFunctionPlot[1_828].Y = 1_972.097

	// pointsOfFunctionPlot[1_829].X = 8.29
	// pointsOfFunctionPlot[1_829].Y = 1_991.916

	// pointsOfFunctionPlot[1_830].X = 8.30
	// pointsOfFunctionPlot[1_830].Y = 2_011.936

	// pointsOfFunctionPlot[1_831].X = 8.31
	// pointsOfFunctionPlot[1_831].Y = 2_032.156

	// pointsOfFunctionPlot[1_832].X = 8.32
	// pointsOfFunctionPlot[1_832].Y = 2_052.579

	// pointsOfFunctionPlot[1_833].X = 8.33
	// pointsOfFunctionPlot[1_833].Y = 2_073.208

	// pointsOfFunctionPlot[1_834].X = 8.34
	// pointsOfFunctionPlot[1_834].Y = 2_094.044

	// pointsOfFunctionPlot[1_835].X = 8.35
	// pointsOfFunctionPlot[1_835].Y = 2_115.09

	// pointsOfFunctionPlot[1_836].X = 8.36
	// pointsOfFunctionPlot[1_836].Y = 2_136.347

	// pointsOfFunctionPlot[1_837].X = 8.37
	// pointsOfFunctionPlot[1_837].Y = 2_157.817

	// pointsOfFunctionPlot[1_838].X = 8.38
	// pointsOfFunctionPlot[1_838].Y = 2_179.504

	// pointsOfFunctionPlot[1_839].X = 8.39
	// pointsOfFunctionPlot[1_839].Y = 2_201.408

	// pointsOfFunctionPlot[1_840].X = 8.40
	// pointsOfFunctionPlot[1_840].Y = 2_223.533

	// pointsOfFunctionPlot[1_841].X = 8.41
	// pointsOfFunctionPlot[1_841].Y = 2_245.88

	// pointsOfFunctionPlot[1_842].X = 8.42
	// pointsOfFunctionPlot[1_842].Y = 2_268.451

	// pointsOfFunctionPlot[1_843].X = 8.43
	// pointsOfFunctionPlot[1_843].Y = 2_291.249

	// pointsOfFunctionPlot[1_844].X = 8.44
	// pointsOfFunctionPlot[1_844].Y = 2_314.277

	// pointsOfFunctionPlot[1_845].X = 8.45
	// pointsOfFunctionPlot[1_845].Y = 2_337.536

	// pointsOfFunctionPlot[1_846].X = 8.46
	// pointsOfFunctionPlot[1_846].Y = 2_361.028

	// pointsOfFunctionPlot[1_847].X = 8.47
	// pointsOfFunctionPlot[1_847].Y = 2_384.757

	// pointsOfFunctionPlot[1_848].X = 8.48
	// pointsOfFunctionPlot[1_848].Y = 2_408.724

	// pointsOfFunctionPlot[1_849].X = 8.49
	// pointsOfFunctionPlot[1_849].Y = 2_432.932

	// pointsOfFunctionPlot[1_850].X = 8.50
	// pointsOfFunctionPlot[1_850].Y = 2_457.384

	// pointsOfFunctionPlot[1_851].X = 8.51
	// pointsOfFunctionPlot[1_851].Y = 2_482.081

	// pointsOfFunctionPlot[1_852].X = 8.52
	// pointsOfFunctionPlot[1_852].Y = 2_507.026

	// pointsOfFunctionPlot[1_853].X = 8.53
	// pointsOfFunctionPlot[1_853].Y = 2_532.222

	// pointsOfFunctionPlot[1_854].X = 8.54
	// pointsOfFunctionPlot[1_854].Y = 2_557.672

	// pointsOfFunctionPlot[1_855].X = 8.55
	// pointsOfFunctionPlot[1_855].Y = 2_583.377

	// pointsOfFunctionPlot[1_856].X = 8.56
	// pointsOfFunctionPlot[1_856].Y = 2_609.34

	// pointsOfFunctionPlot[1_857].X = 8.57
	// pointsOfFunctionPlot[1_857].Y = 2_635.564

	// pointsOfFunctionPlot[1_858].X = 8.58
	// pointsOfFunctionPlot[1_858].Y = 2_662.052

	// pointsOfFunctionPlot[1_859].X = 8.59
	// pointsOfFunctionPlot[1_859].Y = 2_688.806

	// pointsOfFunctionPlot[1_860].X = 8.60
	// pointsOfFunctionPlot[1_860].Y = 2_715.829

	// pointsOfFunctionPlot[1_861].X = 8.61
	// pointsOfFunctionPlot[1_861].Y = 2_743.124

	// pointsOfFunctionPlot[1_862].X = 8.62
	// pointsOfFunctionPlot[1_862].Y = 2_770.693

	// pointsOfFunctionPlot[1_863].X = 8.63
	// pointsOfFunctionPlot[1_863].Y = 2_798.539

	// pointsOfFunctionPlot[1_864].X = 8.64
	// pointsOfFunctionPlot[1_864].Y = 2_826.664

	// pointsOfFunctionPlot[1_865].X = 8.65
	// pointsOfFunctionPlot[1_865].Y = 2_855.073

	// pointsOfFunctionPlot[1_866].X = 8.66
	// pointsOfFunctionPlot[1_866].Y = 2_883.767

	// pointsOfFunctionPlot[1_867].X = 8.67
	// pointsOfFunctionPlot[1_867].Y = 2_912.749

	// pointsOfFunctionPlot[1_868].X = 8.68
	// pointsOfFunctionPlot[1_868].Y = 2_942.023

	// pointsOfFunctionPlot[1_869].X = 8.69
	// pointsOfFunctionPlot[1_869].Y = 2_971.591

	// pointsOfFunctionPlot[1_870].X = 8.70
	// pointsOfFunctionPlot[1_870].Y = 3_001.456

	// pointsOfFunctionPlot[1_871].X = 8.71
	// pointsOfFunctionPlot[1_871].Y = 3_031.621

	// pointsOfFunctionPlot[1_872].X = 8.72
	// pointsOfFunctionPlot[1_872].Y = 3_062.089

	// pointsOfFunctionPlot[1_873].X = 8.73
	// pointsOfFunctionPlot[1_873].Y = 3_092.863

	// pointsOfFunctionPlot[1_874].X = 8.74
	// pointsOfFunctionPlot[1_874].Y = 3_123.947

	// pointsOfFunctionPlot[1_875].X = 8.75
	// pointsOfFunctionPlot[1_875].Y = 3_155.343

	// pointsOfFunctionPlot[1_876].X = 8.76
	// pointsOfFunctionPlot[1_876].Y = 3_187.055

	// pointsOfFunctionPlot[1_877].X = 8.77
	// pointsOfFunctionPlot[1_877].Y = 3_219.086

	// pointsOfFunctionPlot[1_878].X = 8.78
	// pointsOfFunctionPlot[1_878].Y = 3_251.438

	// pointsOfFunctionPlot[1_879].X = 8.79
	// pointsOfFunctionPlot[1_879].Y = 3_284.116

	// pointsOfFunctionPlot[1_880].X = 8.80
	// pointsOfFunctionPlot[1_880].Y = 3_317.121

	// pointsOfFunctionPlot[1_881].X = 8.81
	// pointsOfFunctionPlot[1_881].Y = 3_350.459

	// pointsOfFunctionPlot[1_882].X = 8.82
	// pointsOfFunctionPlot[1_882].Y = 3_384.132

	// pointsOfFunctionPlot[1_883].X = 8.83
	// pointsOfFunctionPlot[1_883].Y = 3_418.143

	// pointsOfFunctionPlot[1_884].X = 8.84
	// pointsOfFunctionPlot[1_884].Y = 3_452.496

	// pointsOfFunctionPlot[1_885].X = 8.85
	// pointsOfFunctionPlot[1_885].Y = 3_487.194

	// pointsOfFunctionPlot[1_886].X = 8.86
	// pointsOfFunctionPlot[1_886].Y = 3_522.241

	// pointsOfFunctionPlot[1_887].X = 8.87
	// pointsOfFunctionPlot[1_887].Y = 3_557.64

	// pointsOfFunctionPlot[1_888].X = 8.88
	// pointsOfFunctionPlot[1_888].Y = 3_593.395

	// pointsOfFunctionPlot[1_889].X = 8.89
	// pointsOfFunctionPlot[1_889].Y = 3_629.509

	// pointsOfFunctionPlot[1_890].X = 8.90
	// pointsOfFunctionPlot[1_890].Y = 3_665.986

	// pointsOfFunctionPlot[1_891].X = 8.91
	// pointsOfFunctionPlot[1_891].Y = 3_702.83

	// pointsOfFunctionPlot[1_892].X = 8.92
	// pointsOfFunctionPlot[1_892].Y = 3_740.044

	// pointsOfFunctionPlot[1_893].X = 8.93
	// pointsOfFunctionPlot[1_893].Y = 3_777.632

	// pointsOfFunctionPlot[1_894].X = 8.94
	// pointsOfFunctionPlot[1_894].Y = 3_815.598

	// pointsOfFunctionPlot[1_895].X = 8.95
	// pointsOfFunctionPlot[1_895].Y = 3_853.945

	// pointsOfFunctionPlot[1_896].X = 8.96
	// pointsOfFunctionPlot[1_896].Y = 3_892.678

	// pointsOfFunctionPlot[1_897].X = 8.97
	// pointsOfFunctionPlot[1_897].Y = 3_931.8

	// pointsOfFunctionPlot[1_898].X = 8.98
	// pointsOfFunctionPlot[1_898].Y = 3_971.315

	// pointsOfFunctionPlot[1_899].X = 8.99
	// pointsOfFunctionPlot[1_899].Y = 4_011.228

	// pointsOfFunctionPlot[1_900].X = 9.0
	// pointsOfFunctionPlot[1_900].Y = 4_051.541

	// pointsOfFunctionPlot[1_901].X = 9.01
	// pointsOfFunctionPlot[1_901].Y = 4_092.26

	// pointsOfFunctionPlot[1_902].X = 9.02
	// pointsOfFunctionPlot[1_902].Y = 4_133.388

	// pointsOfFunctionPlot[1_903].X = 9.03
	// pointsOfFunctionPlot[1_903].Y = 4_174.929

	// pointsOfFunctionPlot[1_904].X = 9.04
	// pointsOfFunctionPlot[1_904].Y = 4_216.888

	// pointsOfFunctionPlot[1_905].X = 9.05
	// pointsOfFunctionPlot[1_905].Y = 4_259.268

	// pointsOfFunctionPlot[1_906].X = 9.06
	// pointsOfFunctionPlot[1_906].Y = 4_302.075

	// pointsOfFunctionPlot[1_907].X = 9.07
	// pointsOfFunctionPlot[1_907].Y = 4_345.311

	// pointsOfFunctionPlot[1_908].X = 9.08
	// pointsOfFunctionPlot[1_908].Y = 4_388.982

	// pointsOfFunctionPlot[1_909].X = 9.09
	// pointsOfFunctionPlot[1_909].Y = 4_433.092

	// pointsOfFunctionPlot[1_910].X = 9.10
	// pointsOfFunctionPlot[1_910].Y = 4_477.646

	// pointsOfFunctionPlot[1_911].X = 9.11
	// pointsOfFunctionPlot[1_911].Y = 4_522.647

	// pointsOfFunctionPlot[1_912].X = 9.12
	// pointsOfFunctionPlot[1_912].Y = 4_568.1

	// pointsOfFunctionPlot[1_913].X = 9.13
	// pointsOfFunctionPlot[1_913].Y = 4_614.01

	// pointsOfFunctionPlot[1_914].X = 9.14
	// pointsOfFunctionPlot[1_914].Y = 4_660.382

	// pointsOfFunctionPlot[1_915].X = 9.15
	// pointsOfFunctionPlot[1_915].Y = 4_707.22

	// pointsOfFunctionPlot[1_916].X = 9.16
	// pointsOfFunctionPlot[1_916].Y = 4_754.528

	// pointsOfFunctionPlot[1_917].X = 9.17
	// pointsOfFunctionPlot[1_917].Y = 4_802.312

	// pointsOfFunctionPlot[1_918].X = 9.18
	// pointsOfFunctionPlot[1_918].Y = 4_850.576

	// pointsOfFunctionPlot[1_919].X = 9.19
	// pointsOfFunctionPlot[1_919].Y = 4_899.325

	// pointsOfFunctionPlot[1_920].X = 9.20
	// pointsOfFunctionPlot[1_920].Y = 4_948.564

	// pointsOfFunctionPlot[1_921].X = 9.21
	// pointsOfFunctionPlot[1_921].Y = 4_998.298

	// pointsOfFunctionPlot[1_922].X = 9.22
	// pointsOfFunctionPlot[1_922].Y = 5_048.532

	// pointsOfFunctionPlot[1_923].X = 9.23
	// pointsOfFunctionPlot[1_923].Y = 5_099.27

	// pointsOfFunctionPlot[1_924].X = 9.24
	// pointsOfFunctionPlot[1_924].Y = 5_150.519

	// pointsOfFunctionPlot[1_925].X = 9.25
	// pointsOfFunctionPlot[1_925].Y = 5_202.282

	// pointsOfFunctionPlot[1_926].X = 9.26
	// pointsOfFunctionPlot[1_926].Y = 5_254.566

	// pointsOfFunctionPlot[1_927].X = 9.27
	// pointsOfFunctionPlot[1_927].Y = 5_307.375

	// pointsOfFunctionPlot[1_928].X = 9.28
	// pointsOfFunctionPlot[1_928].Y = 5_360.715

	// pointsOfFunctionPlot[1_929].X = 9.29
	// pointsOfFunctionPlot[1_929].Y = 5_414.592

	// pointsOfFunctionPlot[1_930].X = 9.30
	// pointsOfFunctionPlot[1_930].Y = 5_469.009

	// pointsOfFunctionPlot[1_931].X = 9.31
	// pointsOfFunctionPlot[1_931].Y = 5_523.974

	// pointsOfFunctionPlot[1_932].X = 9.32
	// pointsOfFunctionPlot[1_932].Y = 5_579.49

	// pointsOfFunctionPlot[1_933].X = 9.33
	// pointsOfFunctionPlot[1_933].Y = 5_635.565

	// pointsOfFunctionPlot[1_934].X = 9.34
	// pointsOfFunctionPlot[1_934].Y = 5_692

	// pointsOfFunctionPlot[1_935].X = 9.35
	// pointsOfFunctionPlot[1_935].Y = 5_749.411

	// pointsOfFunctionPlot[1_936].X = 9.36
	// pointsOfFunctionPlot[1_936].Y = 5_807.194

	// pointsOfFunctionPlot[1_937].X = 9.37
	// pointsOfFunctionPlot[1_937].Y = 5_865.557

	// pointsOfFunctionPlot[1_938].X = 9.38
	// pointsOfFunctionPlot[1_938].Y = 5_924.507

	// pointsOfFunctionPlot[1_939].X = 9.39
	// pointsOfFunctionPlot[1_939].Y = 5_984.049

	// pointsOfFunctionPlot[1_940].X = 9.40
	// pointsOfFunctionPlot[1_940].Y = 6_044.19

	// pointsOfFunctionPlot[1_941].X = 9.41
	// pointsOfFunctionPlot[1_941].Y = 6_104.935

	// pointsOfFunctionPlot[1_942].X = 9.42
	// pointsOfFunctionPlot[1_942].Y = 6_166.291

	// pointsOfFunctionPlot[1_943].X = 9.43
	// pointsOfFunctionPlot[1_943].Y = 6_228.263

	// pointsOfFunctionPlot[1_944].X = 9.44
	// pointsOfFunctionPlot[1_944].Y = 6_290.858

	// pointsOfFunctionPlot[1_945].X = 9.45
	// pointsOfFunctionPlot[1_945].Y = 6_354.082

	// pointsOfFunctionPlot[1_946].X = 9.46
	// pointsOfFunctionPlot[1_946].Y = 6_417.942

	// pointsOfFunctionPlot[1_947].X = 9.47
	// pointsOfFunctionPlot[1_947].Y = 6_482.443

	// pointsOfFunctionPlot[1_948].X = 9.48
	// pointsOfFunctionPlot[1_948].Y = 6_547.593

	// pointsOfFunctionPlot[1_949].X = 9.49
	// pointsOfFunctionPlot[1_949].Y = 6_613.397

	// pointsOfFunctionPlot[1_950].X = 9.50
	// pointsOfFunctionPlot[1_950].Y = 6_679.863

	// pointsOfFunctionPlot[1_951].X = 9.51
	// pointsOfFunctionPlot[1_951].Y = 6_746.997

	// pointsOfFunctionPlot[1_952].X = 9.52
	// pointsOfFunctionPlot[1_952].Y = 6_814.805

	// pointsOfFunctionPlot[1_953].X = 9.53
	// pointsOfFunctionPlot[1_953].Y = 6_883.295

	// pointsOfFunctionPlot[1_954].X = 9.54
	// pointsOfFunctionPlot[1_954].Y = 6_952.473

	// pointsOfFunctionPlot[1_955].X = 9.55
	// pointsOfFunctionPlot[1_955].Y = 7_022.347

	// pointsOfFunctionPlot[1_956].X = 9.56
	// pointsOfFunctionPlot[1_956].Y = 7_092.923

	// pointsOfFunctionPlot[1_957].X = 9.57
	// pointsOfFunctionPlot[1_957].Y = 7_164.208

	// pointsOfFunctionPlot[1_958].X = 9.58
	// pointsOfFunctionPlot[1_958].Y = 7_236.209

	// pointsOfFunctionPlot[1_959].X = 9.59
	// pointsOfFunctionPlot[1_959].Y = 7_308.934

	// pointsOfFunctionPlot[1_960].X = 9.60
	// pointsOfFunctionPlot[1_960].Y = 7_382.39

	// pointsOfFunctionPlot[1_961].X = 9.61
	// pointsOfFunctionPlot[1_961].Y = 7_456.585

	// pointsOfFunctionPlot[1_962].X = 9.62
	// pointsOfFunctionPlot[1_962].Y = 7_531.524

	// pointsOfFunctionPlot[1_963].X = 9.63
	// pointsOfFunctionPlot[1_963].Y = 7_607.218

	// pointsOfFunctionPlot[1_964].X = 9.64
	// pointsOfFunctionPlot[1_964].Y = 7_683.671

	// pointsOfFunctionPlot[1_965].X = 9.65
	// pointsOfFunctionPlot[1_965].Y = 7_760.894

	// pointsOfFunctionPlot[1_966].X = 9.66
	// pointsOfFunctionPlot[1_966].Y = 7_838.892

	// pointsOfFunctionPlot[1_967].X = 9.67
	// pointsOfFunctionPlot[1_967].Y = 7_917.674

	// pointsOfFunctionPlot[1_968].X = 9.68
	// pointsOfFunctionPlot[1_968].Y = 7_997.248

	// pointsOfFunctionPlot[1_969].X = 9.69
	// pointsOfFunctionPlot[1_969].Y = 8_077.622

	// pointsOfFunctionPlot[1_970].X = 9.70
	// pointsOfFunctionPlot[1_970].Y = 8_158.803

	// pointsOfFunctionPlot[1_971].X = 9.71
	// pointsOfFunctionPlot[1_971].Y = 8_240.8

	// pointsOfFunctionPlot[1_972].X = 9.72
	// pointsOfFunctionPlot[1_972].Y = 8_323.622

	// pointsOfFunctionPlot[1_973].X = 9.73
	// pointsOfFunctionPlot[1_973].Y = 8_407.276

	// pointsOfFunctionPlot[1_974].X = 9.74
	// pointsOfFunctionPlot[1_974].Y = 8_491.77

	// pointsOfFunctionPlot[1_975].X = 9.75
	// pointsOfFunctionPlot[1_975].Y = 8_577.114

	// pointsOfFunctionPlot[1_976].X = 9.76
	// pointsOfFunctionPlot[1_976].Y = 8_663.315

	// pointsOfFunctionPlot[1_977].X = 9.77
	// pointsOfFunctionPlot[1_977].Y = 8_750.383

	// pointsOfFunctionPlot[1_978].X = 9.78
	// pointsOfFunctionPlot[1_978].Y = 8_838.326

	// pointsOfFunctionPlot[1_979].X = 9.79
	// pointsOfFunctionPlot[1_979].Y = 8_927.153

	// pointsOfFunctionPlot[1_980].X = 9.80
	// pointsOfFunctionPlot[1_980].Y = 9_016.872

	// pointsOfFunctionPlot[1_981].X = 9.81
	// pointsOfFunctionPlot[1_981].Y = 9_107.493

	// pointsOfFunctionPlot[1_982].X = 9.82
	// pointsOfFunctionPlot[1_982].Y = 9_199.025

	// pointsOfFunctionPlot[1_983].X = 9.83
	// pointsOfFunctionPlot[1_983].Y = 9_291.477

	// pointsOfFunctionPlot[1_984].X = 9.84
	// pointsOfFunctionPlot[1_984].Y = 9_394.857

	// pointsOfFunctionPlot[1_985].X = 9.85
	// pointsOfFunctionPlot[1_985].Y = 9_479.177

	// pointsOfFunctionPlot[1_986].X = 9.86
	// pointsOfFunctionPlot[1_986].Y = 9_574.444

	// pointsOfFunctionPlot[1_987].X = 9.87
	// pointsOfFunctionPlot[1_987].Y = 9_670.669

	// pointsOfFunctionPlot[1_988].X = 9.88
	// pointsOfFunctionPlot[1_988].Y = 9_767.861

	// pointsOfFunctionPlot[1_989].X = 9.89
	// pointsOfFunctionPlot[1_989].Y = 9_866.029

	// pointsOfFunctionPlot[1_990].X = 9.90
	// pointsOfFunctionPlot[1_990].Y = 9_965.185

	// pointsOfFunctionPlot[1_991].X = 9.91
	// pointsOfFunctionPlot[1_991].Y = 10_065.336

	// pointsOfFunctionPlot[1_992].X = 9.92
	// pointsOfFunctionPlot[1_992].Y = 10_166.495

	// pointsOfFunctionPlot[1_993].X = 9.93
	// pointsOfFunctionPlot[1_993].Y = 10_268.67

	// pointsOfFunctionPlot[1_994].X = 9.94
	// pointsOfFunctionPlot[1_994].Y = 10_371.872

	// pointsOfFunctionPlot[1_995].X = 9.95
	// pointsOfFunctionPlot[1_995].Y = 10_476.111

	// pointsOfFunctionPlot[1_996].X = 9.96
	// pointsOfFunctionPlot[1_996].Y = 10_581.397

	// pointsOfFunctionPlot[1_997].X = 9.97
	// pointsOfFunctionPlot[1_997].Y = 10_687.742

	// pointsOfFunctionPlot[1_998].X = 9.98
	// pointsOfFunctionPlot[1_998].Y = 10_795.156

	// pointsOfFunctionPlot[1_999].X = 9.99
	// pointsOfFunctionPlot[1_999].Y = 10_903.649

	// pointsOfFunctionPlot[2_000].X = 10.0
	// pointsOfFunctionPlot[2_000].Y = 11_013.232










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function f(x) = x^2"

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
		"x-to-power-2-function-plot-01.png"); err != nil {

		panic(err)
	}
}
