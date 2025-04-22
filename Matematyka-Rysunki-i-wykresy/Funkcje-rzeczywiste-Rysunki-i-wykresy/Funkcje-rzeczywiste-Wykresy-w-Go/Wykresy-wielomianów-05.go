package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres wielomianu f(x) = 0.001 x^4 - 0.01 x^3 - 0.1 x^2 - x

	pointsOfPolynomialPlot := make(plotter.XYs, 551)

	pointsOfPolynomialPlot[0].X = 0.0
	pointsOfPolynomialPlot[0].Y = 0.0

	pointsOfPolynomialPlot[1].X = 0.01
	pointsOfPolynomialPlot[1].Y = -0.01

	pointsOfPolynomialPlot[2].X = 0.02
	pointsOfPolynomialPlot[2].Y = -0.02

	pointsOfPolynomialPlot[3].X = 0.03
	pointsOfPolynomialPlot[3].Y = -0.03

	pointsOfPolynomialPlot[4].X = 0.04
	pointsOfPolynomialPlot[4].Y = -0.04

	pointsOfPolynomialPlot[5].X = 0.05
	pointsOfPolynomialPlot[5].Y = -0.05

	pointsOfPolynomialPlot[6].X = 0.06
	pointsOfPolynomialPlot[6].Y = -0.06

	pointsOfPolynomialPlot[7].X = 0.07
	pointsOfPolynomialPlot[7].Y = -0.07

	pointsOfPolynomialPlot[8].X = 0.08
	pointsOfPolynomialPlot[8].Y = -0.08

	pointsOfPolynomialPlot[9].X = 0.09
	pointsOfPolynomialPlot[9].Y = -0.09

	pointsOfPolynomialPlot[10].X = 0.1
	pointsOfPolynomialPlot[10].Y = -0.101

	pointsOfPolynomialPlot[11].X = 0.11
	pointsOfPolynomialPlot[11].Y = -0.111

	pointsOfPolynomialPlot[12].X = 0.12
	pointsOfPolynomialPlot[12].Y = -0.121

	pointsOfPolynomialPlot[13].X = 0.13
	pointsOfPolynomialPlot[13].Y = -0.131

	pointsOfPolynomialPlot[14].X = 0.14
	pointsOfPolynomialPlot[14].Y = -0.142

	pointsOfPolynomialPlot[15].X = 0.15
	pointsOfPolynomialPlot[15].Y = -0.152

	pointsOfPolynomialPlot[16].X = 0.16
	pointsOfPolynomialPlot[16].Y = -0.162

	pointsOfPolynomialPlot[17].X = 0.17
	pointsOfPolynomialPlot[17].Y = -0.172

	pointsOfPolynomialPlot[18].X = 0.18
	pointsOfPolynomialPlot[18].Y = -0.183

	pointsOfPolynomialPlot[19].X = 0.19
	pointsOfPolynomialPlot[19].Y = -0.193

	pointsOfPolynomialPlot[20].X = 0.2
	pointsOfPolynomialPlot[20].Y = -0.204

	pointsOfPolynomialPlot[21].X = 0.21
	pointsOfPolynomialPlot[21].Y = -0.214

	pointsOfPolynomialPlot[22].X = 0.22
	pointsOfPolynomialPlot[22].Y = -0.224

	pointsOfPolynomialPlot[23].X = 0.23
	pointsOfPolynomialPlot[23].Y = -0.235

	pointsOfPolynomialPlot[24].X = 0.24
	pointsOfPolynomialPlot[24].Y = -0.245

	pointsOfPolynomialPlot[25].X = 0.25
	pointsOfPolynomialPlot[25].Y = -0.256

	pointsOfPolynomialPlot[26].X = 0.26
	pointsOfPolynomialPlot[26].Y = -0.266

	pointsOfPolynomialPlot[27].X = 0.27
	pointsOfPolynomialPlot[27].Y = -0.277

	pointsOfPolynomialPlot[28].X = 0.28
	pointsOfPolynomialPlot[28].Y = -0.288

	pointsOfPolynomialPlot[29].X = 0.29
	pointsOfPolynomialPlot[29].Y = -0.298

	pointsOfPolynomialPlot[30].X = 0.3
	pointsOfPolynomialPlot[30].Y = -0.309

	pointsOfPolynomialPlot[31].X = 0.31
	pointsOfPolynomialPlot[31].Y = -0.319

	pointsOfPolynomialPlot[32].X = 0.32
	pointsOfPolynomialPlot[32].Y = -0.33

	pointsOfPolynomialPlot[33].X = 0.33
	pointsOfPolynomialPlot[33].Y = -0.341

	pointsOfPolynomialPlot[34].X = 0.34
	pointsOfPolynomialPlot[34].Y = -0.351

	pointsOfPolynomialPlot[35].X = 0.35
	pointsOfPolynomialPlot[35].Y = -0.362

	pointsOfPolynomialPlot[36].X = 0.36
	pointsOfPolynomialPlot[36].Y = -0.373

	pointsOfPolynomialPlot[37].X = 0.37
	pointsOfPolynomialPlot[37].Y = -0.384

	pointsOfPolynomialPlot[38].X = 0.38
	pointsOfPolynomialPlot[38].Y = -0.395

	pointsOfPolynomialPlot[39].X = 0.39
	pointsOfPolynomialPlot[39].Y = -0.405

	pointsOfPolynomialPlot[40].X = 0.4
	pointsOfPolynomialPlot[40].Y = -0.416

	pointsOfPolynomialPlot[41].X = 0.41
	pointsOfPolynomialPlot[41].Y = -0.427

	pointsOfPolynomialPlot[42].X = 0.42
	pointsOfPolynomialPlot[42].Y = -0.438

	pointsOfPolynomialPlot[43].X = 0.43
	pointsOfPolynomialPlot[43].Y = -0.449

	pointsOfPolynomialPlot[44].X = 0.44
	pointsOfPolynomialPlot[44].Y = -0.46

	pointsOfPolynomialPlot[45].X = 0.45
	pointsOfPolynomialPlot[45].Y = -0.471

	pointsOfPolynomialPlot[46].X = 0.46
	pointsOfPolynomialPlot[46].Y = -0.482

	pointsOfPolynomialPlot[47].X = 0.47
	pointsOfPolynomialPlot[47].Y = -0.493

	pointsOfPolynomialPlot[48].X = 0.48
	pointsOfPolynomialPlot[48].Y = -0.504

	pointsOfPolynomialPlot[49].X = 0.49
	pointsOfPolynomialPlot[49].Y = -0.515

	pointsOfPolynomialPlot[50].X = 0.5
	pointsOfPolynomialPlot[50].Y = -0.526

	pointsOfPolynomialPlot[51].X = 0.51
	pointsOfPolynomialPlot[51].Y = -0.537

	pointsOfPolynomialPlot[52].X = 0.52
	pointsOfPolynomialPlot[52].Y = -0.549

	pointsOfPolynomialPlot[53].X = 0.53
	pointsOfPolynomialPlot[53].Y = -0.559

	pointsOfPolynomialPlot[54].X = 0.54
	pointsOfPolynomialPlot[54].Y = -0.57

	pointsOfPolynomialPlot[55].X = 0.55
	pointsOfPolynomialPlot[55].Y = -0.581

	pointsOfPolynomialPlot[56].X = 0.56
	pointsOfPolynomialPlot[56].Y = -0.593

	pointsOfPolynomialPlot[57].X = 0.57
	pointsOfPolynomialPlot[57].Y = -0.604

	pointsOfPolynomialPlot[58].X = 0.58
	pointsOfPolynomialPlot[58].Y = -0.615

	pointsOfPolynomialPlot[59].X = 0.59
	pointsOfPolynomialPlot[59].Y = -0.626

	pointsOfPolynomialPlot[60].X = 0.6
	pointsOfPolynomialPlot[60].Y = -0.638

	pointsOfPolynomialPlot[61].X = 0.61
	pointsOfPolynomialPlot[61].Y = -0.649

	pointsOfPolynomialPlot[62].X = 0.62
	pointsOfPolynomialPlot[62].Y = -0.66

	pointsOfPolynomialPlot[63].X = 0.63
	pointsOfPolynomialPlot[63].Y = -0.672

	pointsOfPolynomialPlot[64].X = 0.64
	pointsOfPolynomialPlot[64].Y = -0.683

	pointsOfPolynomialPlot[65].X = 0.65
	pointsOfPolynomialPlot[65].Y = -0.694

	pointsOfPolynomialPlot[66].X = 0.66
	pointsOfPolynomialPlot[66].Y = -0.706

	pointsOfPolynomialPlot[67].X = 0.67
	pointsOfPolynomialPlot[67].Y = -0.717

	pointsOfPolynomialPlot[68].X = 0.68
	pointsOfPolynomialPlot[68].Y = -0.729

	pointsOfPolynomialPlot[69].X = 0.69
	pointsOfPolynomialPlot[69].Y = -0.74

	pointsOfPolynomialPlot[70].X = 0.7
	pointsOfPolynomialPlot[70].Y = -0.752

	pointsOfPolynomialPlot[71].X = 0.71
	pointsOfPolynomialPlot[71].Y = -0.763

	pointsOfPolynomialPlot[72].X = 0.72
	pointsOfPolynomialPlot[72].Y = -0.775

	pointsOfPolynomialPlot[73].X = 0.73
	pointsOfPolynomialPlot[73].Y = -0.786

	pointsOfPolynomialPlot[74].X = 0.74
	pointsOfPolynomialPlot[74].Y = -0.798

	pointsOfPolynomialPlot[75].X = 0.75
	pointsOfPolynomialPlot[75].Y = -0.81

	pointsOfPolynomialPlot[76].X = 0.76
	pointsOfPolynomialPlot[76].Y = -0.821

	pointsOfPolynomialPlot[77].X = 0.77
	pointsOfPolynomialPlot[77].Y = -0.833

	pointsOfPolynomialPlot[78].X = 0.78
	pointsOfPolynomialPlot[78].Y = -0.845

	pointsOfPolynomialPlot[79].X = 0.79
	pointsOfPolynomialPlot[79].Y = -0.857

	pointsOfPolynomialPlot[80].X = 0.8
	pointsOfPolynomialPlot[80].Y = -0.868

	pointsOfPolynomialPlot[81].X = 0.81
	pointsOfPolynomialPlot[81].Y = -0.88

	pointsOfPolynomialPlot[82].X = 0.82
	pointsOfPolynomialPlot[82].Y = -0.892

	pointsOfPolynomialPlot[83].X = 0.83
	pointsOfPolynomialPlot[83].Y = -0.904

	pointsOfPolynomialPlot[84].X = 0.84
	pointsOfPolynomialPlot[84].Y = -0.916

	pointsOfPolynomialPlot[85].X = 0.85
	pointsOfPolynomialPlot[85].Y = -0.927

	pointsOfPolynomialPlot[86].X = 0.86
	pointsOfPolynomialPlot[86].Y = -0.939

	pointsOfPolynomialPlot[87].X = 0.87
	pointsOfPolynomialPlot[87].Y = -0.951

	pointsOfPolynomialPlot[88].X = 0.88
	pointsOfPolynomialPlot[88].Y = -0.963

	pointsOfPolynomialPlot[89].X = 0.89
	pointsOfPolynomialPlot[89].Y = -0.975

	pointsOfPolynomialPlot[90].X = 0.9
	pointsOfPolynomialPlot[90].Y = -0.987

	pointsOfPolynomialPlot[91].X = 0.91
	pointsOfPolynomialPlot[91].Y = -0.999

	pointsOfPolynomialPlot[92].X = 0.92
	pointsOfPolynomialPlot[92].Y = -1.011

	pointsOfPolynomialPlot[93].X = 0.93
	pointsOfPolynomialPlot[93].Y = -1.023

	pointsOfPolynomialPlot[94].X = 0.94
	pointsOfPolynomialPlot[94].Y = -1.035

	pointsOfPolynomialPlot[95].X = 0.95
	pointsOfPolynomialPlot[95].Y = -1.048

	pointsOfPolynomialPlot[96].X = 0.96
	pointsOfPolynomialPlot[96].Y = -1.06

	pointsOfPolynomialPlot[97].X = 0.97
	pointsOfPolynomialPlot[97].Y = -1.072

	pointsOfPolynomialPlot[98].X = 0.98
	pointsOfPolynomialPlot[98].Y = -1.084

	pointsOfPolynomialPlot[99].X = 0.99
	pointsOfPolynomialPlot[99].Y = -1.096

	pointsOfPolynomialPlot[100].X = 1.0
	pointsOfPolynomialPlot[100].Y = -1.109

	pointsOfPolynomialPlot[101].X = 1.01
	pointsOfPolynomialPlot[101].Y = -1.121

	pointsOfPolynomialPlot[102].X = 1.02
	pointsOfPolynomialPlot[102].Y = -1.133

	pointsOfPolynomialPlot[103].X = 1.03
	pointsOfPolynomialPlot[103].Y = -1.145

	pointsOfPolynomialPlot[104].X = 1.04
	pointsOfPolynomialPlot[104].Y = -1.158

	pointsOfPolynomialPlot[105].X = 1.05
	pointsOfPolynomialPlot[105].Y = -1.17

	pointsOfPolynomialPlot[106].X = 1.06
	pointsOfPolynomialPlot[106].Y = -1.183

	pointsOfPolynomialPlot[107].X = 1.07
	pointsOfPolynomialPlot[107].Y = -1.195

	pointsOfPolynomialPlot[108].X = 1.08
	pointsOfPolynomialPlot[108].Y = -1.207

	pointsOfPolynomialPlot[109].X = 1.09
	pointsOfPolynomialPlot[109].Y = -1.22

	pointsOfPolynomialPlot[110].X = 1.1
	pointsOfPolynomialPlot[110].Y = -1.232

	pointsOfPolynomialPlot[111].X = 1.11
	pointsOfPolynomialPlot[111].Y = -1.245

	pointsOfPolynomialPlot[112].X = 1.12
	pointsOfPolynomialPlot[112].Y = -1.257

	pointsOfPolynomialPlot[113].X = 1.13
	pointsOfPolynomialPlot[113].Y = -1.27

	pointsOfPolynomialPlot[114].X = 1.14
	pointsOfPolynomialPlot[114].Y = -1.283

	pointsOfPolynomialPlot[115].X = 1.15
	pointsOfPolynomialPlot[115].Y = -1.295

	pointsOfPolynomialPlot[116].X = 1.16
	pointsOfPolynomialPlot[116].Y = -1.308

	pointsOfPolynomialPlot[117].X = 1.17
	pointsOfPolynomialPlot[117].Y = -1.321

	pointsOfPolynomialPlot[118].X = 1.18
	pointsOfPolynomialPlot[118].Y = -1.333

	pointsOfPolynomialPlot[119].X = 1.19
	pointsOfPolynomialPlot[119].Y = -1.346

	pointsOfPolynomialPlot[120].X = 1.2
	pointsOfPolynomialPlot[120].Y = -1.359

	pointsOfPolynomialPlot[121].X = 1.21
	pointsOfPolynomialPlot[121].Y = -1.372

	pointsOfPolynomialPlot[122].X = 1.22
	pointsOfPolynomialPlot[122].Y = -1.384

	pointsOfPolynomialPlot[123].X = 1.23
	pointsOfPolynomialPlot[123].Y = -1.397

	pointsOfPolynomialPlot[124].X = 1.24
	pointsOfPolynomialPlot[124].Y = -1.41

	pointsOfPolynomialPlot[125].X = 1.25
	pointsOfPolynomialPlot[125].Y = -1.423

	pointsOfPolynomialPlot[126].X = 1.26
	pointsOfPolynomialPlot[126].Y = -1.436

	pointsOfPolynomialPlot[127].X = 1.27
	pointsOfPolynomialPlot[127].Y = -1.449

	pointsOfPolynomialPlot[128].X = 1.28
	pointsOfPolynomialPlot[128].Y = -1.462

	pointsOfPolynomialPlot[129].X = 1.29
	pointsOfPolynomialPlot[129].Y = -1.475

	pointsOfPolynomialPlot[130].X = 1.3
	pointsOfPolynomialPlot[130].Y = -1.488

	pointsOfPolynomialPlot[131].X = 1.31
	pointsOfPolynomialPlot[131].Y = -1.501

	pointsOfPolynomialPlot[132].X = 1.32
	pointsOfPolynomialPlot[132].Y = -1.514

	pointsOfPolynomialPlot[133].X = 1.33
	pointsOfPolynomialPlot[133].Y = -1.527

	pointsOfPolynomialPlot[134].X = 1.34
	pointsOfPolynomialPlot[134].Y = -1.54

	pointsOfPolynomialPlot[135].X = 1.35
	pointsOfPolynomialPlot[135].Y = -1.553

	pointsOfPolynomialPlot[136].X = 1.36
	pointsOfPolynomialPlot[136].Y = -1.566

	pointsOfPolynomialPlot[137].X = 1.37
	pointsOfPolynomialPlot[137].Y = -1.579

	pointsOfPolynomialPlot[138].X = 1.38
	pointsOfPolynomialPlot[138].Y = -1.593

	pointsOfPolynomialPlot[139].X = 1.39
	pointsOfPolynomialPlot[139].Y = -1.606

	pointsOfPolynomialPlot[140].X = 1.4
	pointsOfPolynomialPlot[140].Y = -1.619

	pointsOfPolynomialPlot[141].X = 1.41
	pointsOfPolynomialPlot[141].Y = -1.632

	pointsOfPolynomialPlot[142].X = 1.42
	pointsOfPolynomialPlot[142].Y = -1.646

	pointsOfPolynomialPlot[143].X = 1.43
	pointsOfPolynomialPlot[143].Y = -1.659

	pointsOfPolynomialPlot[144].X = 1.44
	pointsOfPolynomialPlot[144].Y = -1.672

	pointsOfPolynomialPlot[145].X = 1.45
	pointsOfPolynomialPlot[145].Y = -1.686

	pointsOfPolynomialPlot[146].X = 1.46
	pointsOfPolynomialPlot[146].Y = -1.699

	pointsOfPolynomialPlot[147].X = 1.47
	pointsOfPolynomialPlot[147].Y = -1.713

	pointsOfPolynomialPlot[148].X = 1.48
	pointsOfPolynomialPlot[148].Y = -1.726

	pointsOfPolynomialPlot[149].X = 1.49
	pointsOfPolynomialPlot[149].Y = -1.74

	pointsOfPolynomialPlot[150].X = 1.5
	pointsOfPolynomialPlot[150].Y = -1.753

	pointsOfPolynomialPlot[151].X = 1.51
	pointsOfPolynomialPlot[151].Y = -1.767

	pointsOfPolynomialPlot[152].X = 1.52
	pointsOfPolynomialPlot[152].Y = -1.78

	pointsOfPolynomialPlot[153].X = 1.53
	pointsOfPolynomialPlot[153].Y = -1.794

	pointsOfPolynomialPlot[154].X = 1.54
	pointsOfPolynomialPlot[154].Y = -1.808

	pointsOfPolynomialPlot[155].X = 1.55
	pointsOfPolynomialPlot[155].Y = -1.821

	pointsOfPolynomialPlot[156].X = 1.56
	pointsOfPolynomialPlot[156].Y = -1.835

	pointsOfPolynomialPlot[157].X = 1.57
	pointsOfPolynomialPlot[157].Y = -1.849

	pointsOfPolynomialPlot[158].X = 1.58
	pointsOfPolynomialPlot[158].Y = -1.862

	pointsOfPolynomialPlot[159].X = 1.59
	pointsOfPolynomialPlot[159].Y = -1.876

	pointsOfPolynomialPlot[160].X = 1.6
	pointsOfPolynomialPlot[160].Y = -1.89

	pointsOfPolynomialPlot[161].X = 1.61
	pointsOfPolynomialPlot[161].Y = -1.904

	pointsOfPolynomialPlot[162].X = 1.62
	pointsOfPolynomialPlot[162].Y = -1.918

	pointsOfPolynomialPlot[163].X = 1.63
	pointsOfPolynomialPlot[163].Y = -1.931

	pointsOfPolynomialPlot[164].X = 1.64
	pointsOfPolynomialPlot[164].Y = -1.945

	pointsOfPolynomialPlot[165].X = 1.65
	pointsOfPolynomialPlot[165].Y = -1.959

	pointsOfPolynomialPlot[166].X = 1.66
	pointsOfPolynomialPlot[166].Y = -1.973

	pointsOfPolynomialPlot[167].X = 1.67
	pointsOfPolynomialPlot[167].Y = -1.987

	pointsOfPolynomialPlot[168].X = 1.68
	pointsOfPolynomialPlot[168].Y = -2.001

	pointsOfPolynomialPlot[169].X = 1.69
	pointsOfPolynomialPlot[169].Y = -2.015

	pointsOfPolynomialPlot[170].X = 1.7
	pointsOfPolynomialPlot[170].Y = -2.029

	pointsOfPolynomialPlot[171].X = 1.71
	pointsOfPolynomialPlot[171].Y = -2.043

	pointsOfPolynomialPlot[172].X = 1.72
	pointsOfPolynomialPlot[172].Y = -2.058

	pointsOfPolynomialPlot[173].X = 1.73
	pointsOfPolynomialPlot[173].Y = -2.072

	pointsOfPolynomialPlot[174].X = 1.74
	pointsOfPolynomialPlot[174].Y = -2.086

	pointsOfPolynomialPlot[175].X = 1.75
	pointsOfPolynomialPlot[175].Y = -2.1

	pointsOfPolynomialPlot[176].X = 1.76
	pointsOfPolynomialPlot[176].Y = -2.114

	pointsOfPolynomialPlot[177].X = 1.77
	pointsOfPolynomialPlot[177].Y = -2.128

	pointsOfPolynomialPlot[178].X = 1.78
	pointsOfPolynomialPlot[178].Y = -2.143

	pointsOfPolynomialPlot[179].X = 1.79
	pointsOfPolynomialPlot[179].Y = -2.157

	pointsOfPolynomialPlot[180].X = 1.8
	pointsOfPolynomialPlot[180].Y = -2.171

	pointsOfPolynomialPlot[181].X = 1.81
	pointsOfPolynomialPlot[181].Y = -2.186

	pointsOfPolynomialPlot[182].X = 1.82
	pointsOfPolynomialPlot[182].Y = -2.2

	pointsOfPolynomialPlot[183].X = 1.83
	pointsOfPolynomialPlot[183].Y = -2.215

	pointsOfPolynomialPlot[184].X = 1.84
	pointsOfPolynomialPlot[184].Y = -2.229

	pointsOfPolynomialPlot[185].X = 1.85
	pointsOfPolynomialPlot[185].Y = -2.243

	pointsOfPolynomialPlot[186].X = 1.86
	pointsOfPolynomialPlot[186].Y = -2.258

	pointsOfPolynomialPlot[187].X = 1.87
	pointsOfPolynomialPlot[187].Y = -2.272

	pointsOfPolynomialPlot[188].X = 1.88
	pointsOfPolynomialPlot[188].Y = -2.287

	pointsOfPolynomialPlot[189].X = 1.89
	pointsOfPolynomialPlot[189].Y = -2.302

	pointsOfPolynomialPlot[190].X = 1.9
	pointsOfPolynomialPlot[190].Y = -2.316

	pointsOfPolynomialPlot[191].X = 1.91
	pointsOfPolynomialPlot[191].Y = -2.331

	pointsOfPolynomialPlot[192].X = 1.92
	pointsOfPolynomialPlot[192].Y = -2.345

	pointsOfPolynomialPlot[193].X = 1.93
	pointsOfPolynomialPlot[193].Y = -2.36

	pointsOfPolynomialPlot[194].X = 1.94
	pointsOfPolynomialPlot[194].Y = -2.375

	pointsOfPolynomialPlot[195].X = 1.95
	pointsOfPolynomialPlot[195].Y = -2.389

	pointsOfPolynomialPlot[196].X = 1.96
	pointsOfPolynomialPlot[196].Y = -2.404

	pointsOfPolynomialPlot[197].X = 1.97
	pointsOfPolynomialPlot[197].Y = -2.419

	pointsOfPolynomialPlot[198].X = 1.98
	pointsOfPolynomialPlot[198].Y = -2.434

	pointsOfPolynomialPlot[199].X = 1.99
	pointsOfPolynomialPlot[199].Y = -2.449

	pointsOfPolynomialPlot[200].X = 2.0
	pointsOfPolynomialPlot[200].Y = -2.464

	pointsOfPolynomialPlot[201].X = 2.01
	pointsOfPolynomialPlot[201].Y = -2.478

	pointsOfPolynomialPlot[202].X = 2.02
	pointsOfPolynomialPlot[202].Y = -2.493

	pointsOfPolynomialPlot[203].X = 2.03
	pointsOfPolynomialPlot[203].Y = -2.508

	pointsOfPolynomialPlot[204].X = 2.04
	pointsOfPolynomialPlot[204].Y = -2.523

	pointsOfPolynomialPlot[205].X = 2.05
	pointsOfPolynomialPlot[205].Y = -2.538

	pointsOfPolynomialPlot[206].X = 2.06
	pointsOfPolynomialPlot[206].Y = -2.553

	pointsOfPolynomialPlot[207].X = 2.07
	pointsOfPolynomialPlot[207].Y = -2.568

	pointsOfPolynomialPlot[208].X = 2.08
	pointsOfPolynomialPlot[208].Y = -2.583

	pointsOfPolynomialPlot[209].X = 2.09
	pointsOfPolynomialPlot[209].Y = -2.599

	pointsOfPolynomialPlot[210].X = 2.1
	pointsOfPolynomialPlot[210].Y = -2.614

	pointsOfPolynomialPlot[211].X = 2.11
	pointsOfPolynomialPlot[211].Y = -2.629

	pointsOfPolynomialPlot[212].X = 2.12
	pointsOfPolynomialPlot[212].Y = -2.644

	pointsOfPolynomialPlot[213].X = 2.13
	pointsOfPolynomialPlot[213].Y = -2.659

	pointsOfPolynomialPlot[214].X = 2.14
	pointsOfPolynomialPlot[214].Y = -2.675

	pointsOfPolynomialPlot[215].X = 2.15
	pointsOfPolynomialPlot[215].Y = -2.69

	pointsOfPolynomialPlot[216].X = 2.16
	pointsOfPolynomialPlot[216].Y = -2.705

	pointsOfPolynomialPlot[217].X = 2.17
	pointsOfPolynomialPlot[217].Y = -2.72

	pointsOfPolynomialPlot[218].X = 2.18
	pointsOfPolynomialPlot[218].Y = -2.736

	pointsOfPolynomialPlot[219].X = 2.19
	pointsOfPolynomialPlot[219].Y = -2.751

	pointsOfPolynomialPlot[220].X = 2.2
	pointsOfPolynomialPlot[220].Y = -2.767

	pointsOfPolynomialPlot[221].X = 2.21
	pointsOfPolynomialPlot[221].Y = -2.782

	pointsOfPolynomialPlot[222].X = 2.22
	pointsOfPolynomialPlot[222].Y = -2.798

	pointsOfPolynomialPlot[223].X = 2.23
	pointsOfPolynomialPlot[223].Y = -2.813

	pointsOfPolynomialPlot[224].X = 2.24
	pointsOfPolynomialPlot[224].Y = -2.829

	pointsOfPolynomialPlot[225].X = 2.25
	pointsOfPolynomialPlot[225].Y = -2.844

	pointsOfPolynomialPlot[226].X = 2.26
	pointsOfPolynomialPlot[226].Y = -2.86

	pointsOfPolynomialPlot[227].X = 2.27
	pointsOfPolynomialPlot[227].Y = -2.875

	pointsOfPolynomialPlot[228].X = 2.28
	pointsOfPolynomialPlot[228].Y = -2.891

	pointsOfPolynomialPlot[229].X = 2.29
	pointsOfPolynomialPlot[229].Y = -2.907

	pointsOfPolynomialPlot[230].X = 2.3
	pointsOfPolynomialPlot[230].Y = -2.922

	pointsOfPolynomialPlot[231].X = 2.31
	pointsOfPolynomialPlot[231].Y = -2.938

	pointsOfPolynomialPlot[232].X = 2.32
	pointsOfPolynomialPlot[232].Y = -2.954

	pointsOfPolynomialPlot[233].X = 2.33
	pointsOfPolynomialPlot[233].Y = -2.969

	pointsOfPolynomialPlot[234].X = 2.34
	pointsOfPolynomialPlot[234].Y = -2.985

	pointsOfPolynomialPlot[235].X = 2.35
	pointsOfPolynomialPlot[235].Y = -3.001

	pointsOfPolynomialPlot[236].X = 2.36
	pointsOfPolynomialPlot[236].Y = -3.017

	pointsOfPolynomialPlot[237].X = 2.37
	pointsOfPolynomialPlot[237].Y = -3.033

	pointsOfPolynomialPlot[238].X = 2.38
	pointsOfPolynomialPlot[238].Y = -3.049

	pointsOfPolynomialPlot[239].X = 2.39
	pointsOfPolynomialPlot[239].Y = -3.065

	pointsOfPolynomialPlot[240].X = 2.4
	pointsOfPolynomialPlot[240].Y = -3.081

	pointsOfPolynomialPlot[241].X = 2.41
	pointsOfPolynomialPlot[241].Y = -3.097

	pointsOfPolynomialPlot[242].X = 2.42
	pointsOfPolynomialPlot[242].Y = -3.113

	pointsOfPolynomialPlot[243].X = 2.43
	pointsOfPolynomialPlot[243].Y = -3.129

	pointsOfPolynomialPlot[244].X = 2.44
	pointsOfPolynomialPlot[244].Y = -3.145

	pointsOfPolynomialPlot[245].X = 2.45
	pointsOfPolynomialPlot[245].Y = -3.161

	pointsOfPolynomialPlot[246].X = 2.46
	pointsOfPolynomialPlot[246].Y = -3.177

	pointsOfPolynomialPlot[247].X = 2.47
	pointsOfPolynomialPlot[247].Y = -3.193

	pointsOfPolynomialPlot[248].X = 2.48
	pointsOfPolynomialPlot[248].Y = -3.209

	pointsOfPolynomialPlot[249].X = 2.49
	pointsOfPolynomialPlot[249].Y = -3.226

	pointsOfPolynomialPlot[250].X = 2.5
	pointsOfPolynomialPlot[250].Y = -3.242

	pointsOfPolynomialPlot[251].X = 2.51
	pointsOfPolynomialPlot[251].Y = -3.258

	pointsOfPolynomialPlot[252].X = 2.52
	pointsOfPolynomialPlot[252].Y = -3.274

	pointsOfPolynomialPlot[253].X = 2.53
	pointsOfPolynomialPlot[253].Y = -3.291

	pointsOfPolynomialPlot[254].X = 2.54
	pointsOfPolynomialPlot[254].Y = -3.307

	pointsOfPolynomialPlot[255].X = 2.55
	pointsOfPolynomialPlot[255].Y = -3.323

	pointsOfPolynomialPlot[256].X = 2.56
	pointsOfPolynomialPlot[256].Y = -3.34

	pointsOfPolynomialPlot[257].X = 2.57
	pointsOfPolynomialPlot[257].Y = -3.356

	pointsOfPolynomialPlot[258].X = 2.58
	pointsOfPolynomialPlot[258].Y = -3.373

	pointsOfPolynomialPlot[259].X = 2.59
	pointsOfPolynomialPlot[259].Y = -3.389

	pointsOfPolynomialPlot[260].X = 2.6
	pointsOfPolynomialPlot[260].Y = -3.406

	pointsOfPolynomialPlot[261].X = 2.61
	pointsOfPolynomialPlot[261].Y = -3.422

	pointsOfPolynomialPlot[262].X = 2.62
	pointsOfPolynomialPlot[262].Y = -3.439

	pointsOfPolynomialPlot[263].X = 2.63
	pointsOfPolynomialPlot[263].Y = -3.455

	pointsOfPolynomialPlot[264].X = 2.64
	pointsOfPolynomialPlot[264].Y = -3.472

	pointsOfPolynomialPlot[265].X = 2.65
	pointsOfPolynomialPlot[265].Y = -3.489

	pointsOfPolynomialPlot[266].X = 2.66
	pointsOfPolynomialPlot[266].Y = -3.505

	pointsOfPolynomialPlot[267].X = 2.67
	pointsOfPolynomialPlot[267].Y = -3.522

	pointsOfPolynomialPlot[268].X = 2.68
	pointsOfPolynomialPlot[268].Y = -3.539

	pointsOfPolynomialPlot[269].X = 2.69
	pointsOfPolynomialPlot[269].Y = -3.555

	pointsOfPolynomialPlot[270].X = 2.7
	pointsOfPolynomialPlot[270].Y = -3.572

	pointsOfPolynomialPlot[271].X = 2.71
	pointsOfPolynomialPlot[271].Y = -3.589

	pointsOfPolynomialPlot[272].X = 2.72
	pointsOfPolynomialPlot[272].Y = -3.606

	pointsOfPolynomialPlot[273].X = 2.73
	pointsOfPolynomialPlot[273].Y = -3.623

	pointsOfPolynomialPlot[274].X = 2.74
	pointsOfPolynomialPlot[274].Y = -3.64

	pointsOfPolynomialPlot[275].X = 2.75
	pointsOfPolynomialPlot[275].Y = -3.657

	pointsOfPolynomialPlot[276].X = 2.76
	pointsOfPolynomialPlot[276].Y = -3.674

	pointsOfPolynomialPlot[277].X = 2.77
	pointsOfPolynomialPlot[277].Y = -3.691

	pointsOfPolynomialPlot[278].X = 2.78
	pointsOfPolynomialPlot[278].Y = -3.708

	pointsOfPolynomialPlot[279].X = 2.79
	pointsOfPolynomialPlot[279].Y = -3.725

	pointsOfPolynomialPlot[280].X = 2.8
	pointsOfPolynomialPlot[280].Y = -3.742

	pointsOfPolynomialPlot[281].X = 2.81
	pointsOfPolynomialPlot[281].Y = -3.759

	pointsOfPolynomialPlot[282].X = 2.82
	pointsOfPolynomialPlot[282].Y = -3.776

	pointsOfPolynomialPlot[283].X = 2.83
	pointsOfPolynomialPlot[283].Y = -3.793

	pointsOfPolynomialPlot[284].X = 2.84
	pointsOfPolynomialPlot[284].Y = -3.81

	pointsOfPolynomialPlot[285].X = 2.85
	pointsOfPolynomialPlot[285].Y = -3.827

	pointsOfPolynomialPlot[286].X = 2.86
	pointsOfPolynomialPlot[286].Y = -3.845

	pointsOfPolynomialPlot[287].X = 2.87
	pointsOfPolynomialPlot[287].Y = -3.862

	pointsOfPolynomialPlot[288].X = 2.88
	pointsOfPolynomialPlot[288].Y = -3.879

	pointsOfPolynomialPlot[289].X = 2.89
	pointsOfPolynomialPlot[289].Y = -3.896

	pointsOfPolynomialPlot[290].X = 2.9
	pointsOfPolynomialPlot[290].Y = -3.914

	pointsOfPolynomialPlot[291].X = 2.91
	pointsOfPolynomialPlot[291].Y = -3.931

	pointsOfPolynomialPlot[292].X = 2.92
	pointsOfPolynomialPlot[292].Y = -3.948

	pointsOfPolynomialPlot[293].X = 2.93
	pointsOfPolynomialPlot[293].Y = -3.966

	pointsOfPolynomialPlot[294].X = 2.94
	pointsOfPolynomialPlot[294].Y = -3.983

	pointsOfPolynomialPlot[295].X = 2.95
	pointsOfPolynomialPlot[295].Y = -4.001

	pointsOfPolynomialPlot[296].X = 2.96
	pointsOfPolynomialPlot[296].Y = -4.018

	pointsOfPolynomialPlot[297].X = 2.97
	pointsOfPolynomialPlot[297].Y = -4.036

	pointsOfPolynomialPlot[298].X = 2.98
	pointsOfPolynomialPlot[298].Y = -4.053

	pointsOfPolynomialPlot[299].X = 2.99
	pointsOfPolynomialPlot[299].Y = -4.071

	pointsOfPolynomialPlot[300].X = 3.0
	pointsOfPolynomialPlot[300].Y = -4.089

	pointsOfPolynomialPlot[301].X = 3.01
	pointsOfPolynomialPlot[301].Y = -4.106

	pointsOfPolynomialPlot[302].X = 3.02
	pointsOfPolynomialPlot[302].Y = -4.123

	pointsOfPolynomialPlot[303].X = 3.03
	pointsOfPolynomialPlot[303].Y = -4.141

	pointsOfPolynomialPlot[304].X = 3.04
	pointsOfPolynomialPlot[304].Y = -4.159

	pointsOfPolynomialPlot[305].X = 3.05
	pointsOfPolynomialPlot[305].Y = -4.177

	pointsOfPolynomialPlot[306].X = 3.06
	pointsOfPolynomialPlot[306].Y = -4.195

	pointsOfPolynomialPlot[307].X = 3.07
	pointsOfPolynomialPlot[307].Y = -4.213

	pointsOfPolynomialPlot[308].X = 3.08
	pointsOfPolynomialPlot[308].Y = -4.23

	pointsOfPolynomialPlot[309].X = 3.09
	pointsOfPolynomialPlot[309].Y = -4.248

	pointsOfPolynomialPlot[310].X = 3.1
	pointsOfPolynomialPlot[310].Y = -4.266

	pointsOfPolynomialPlot[311].X = 3.11
	pointsOfPolynomialPlot[311].Y = -4.284

	pointsOfPolynomialPlot[312].X = 3.12
	pointsOfPolynomialPlot[312].Y = -4.302

	pointsOfPolynomialPlot[313].X = 3.13
	pointsOfPolynomialPlot[313].Y = -4.32

	pointsOfPolynomialPlot[314].X = 3.14
	pointsOfPolynomialPlot[314].Y = -4.338

	pointsOfPolynomialPlot[315].X = 3.15
	pointsOfPolynomialPlot[315].Y = -4.356

	pointsOfPolynomialPlot[316].X = 3.16
	pointsOfPolynomialPlot[316].Y = -4.374

	pointsOfPolynomialPlot[317].X = 3.17
	pointsOfPolynomialPlot[317].Y = -4.392

	pointsOfPolynomialPlot[318].X = 3.18
	pointsOfPolynomialPlot[318].Y = -4.41

	pointsOfPolynomialPlot[319].X = 3.19
	pointsOfPolynomialPlot[319].Y = -4.428

	pointsOfPolynomialPlot[320].X = 3.2
	pointsOfPolynomialPlot[320].Y = -4.446

	pointsOfPolynomialPlot[321].X = 3.21
	pointsOfPolynomialPlot[321].Y = -4.465

	pointsOfPolynomialPlot[322].X = 3.22
	pointsOfPolynomialPlot[322].Y = -4.483

	pointsOfPolynomialPlot[323].X = 3.23
	pointsOfPolynomialPlot[323].Y = -4.501

	pointsOfPolynomialPlot[324].X = 3.24
	pointsOfPolynomialPlot[324].Y = -4.519

	pointsOfPolynomialPlot[325].X = 3.25
	pointsOfPolynomialPlot[325].Y = -4.538

	pointsOfPolynomialPlot[326].X = 3.26
	pointsOfPolynomialPlot[326].Y = -4.556

	pointsOfPolynomialPlot[327].X = 3.27
	pointsOfPolynomialPlot[327].Y = -4.574

	pointsOfPolynomialPlot[328].X = 3.28
	pointsOfPolynomialPlot[328].Y = -4.593

	pointsOfPolynomialPlot[329].X = 3.29
	pointsOfPolynomialPlot[329].Y = -4.611

	pointsOfPolynomialPlot[330].X = 3.3
	pointsOfPolynomialPlot[330].Y = -4.629

	pointsOfPolynomialPlot[331].X = 3.31
	pointsOfPolynomialPlot[331].Y = -4.648

	pointsOfPolynomialPlot[332].X = 3.32
	pointsOfPolynomialPlot[332].Y = -4.666

	pointsOfPolynomialPlot[333].X = 3.33
	pointsOfPolynomialPlot[333].Y = -4.685

	pointsOfPolynomialPlot[334].X = 3.34
	pointsOfPolynomialPlot[334].Y = -4.703

	pointsOfPolynomialPlot[335].X = 3.35
	pointsOfPolynomialPlot[335].Y = -4.722

	pointsOfPolynomialPlot[336].X = 3.36
	pointsOfPolynomialPlot[336].Y = -4.74

	pointsOfPolynomialPlot[337].X = 3.37
	pointsOfPolynomialPlot[337].Y = -4.759

	pointsOfPolynomialPlot[338].X = 3.38
	pointsOfPolynomialPlot[338].Y = -4.778

	pointsOfPolynomialPlot[339].X = 3.39
	pointsOfPolynomialPlot[339].Y = -4.796

	pointsOfPolynomialPlot[340].X = 3.4
	pointsOfPolynomialPlot[340].Y = -4.815

	pointsOfPolynomialPlot[341].X = 3.41
	pointsOfPolynomialPlot[341].Y = -4.834

	pointsOfPolynomialPlot[342].X = 3.42
	pointsOfPolynomialPlot[342].Y = -4.852

	pointsOfPolynomialPlot[343].X = 3.43
	pointsOfPolynomialPlot[343].Y = -4.871

	pointsOfPolynomialPlot[344].X = 3.44
	pointsOfPolynomialPlot[344].Y = -4.89

	pointsOfPolynomialPlot[345].X = 3.45
	pointsOfPolynomialPlot[345].Y = -4.909

	pointsOfPolynomialPlot[346].X = 3.46
	pointsOfPolynomialPlot[346].Y = -4.928

	pointsOfPolynomialPlot[347].X = 3.47
	pointsOfPolynomialPlot[347].Y = -4.946

	pointsOfPolynomialPlot[348].X = 3.48
	pointsOfPolynomialPlot[348].Y = -4.965

	pointsOfPolynomialPlot[349].X = 3.49
	pointsOfPolynomialPlot[349].Y = -4.984

	pointsOfPolynomialPlot[350].X = 3.5
	pointsOfPolynomialPlot[350].Y = -5.003

	pointsOfPolynomialPlot[351].X = 3.51
	pointsOfPolynomialPlot[351].Y = -5.022

	pointsOfPolynomialPlot[352].X = 3.52
	pointsOfPolynomialPlot[352].Y = -5.041

	pointsOfPolynomialPlot[353].X = 3.53
	pointsOfPolynomialPlot[353].Y = -5.06

	pointsOfPolynomialPlot[354].X = 3.54
	pointsOfPolynomialPlot[354].Y = -5.079

	pointsOfPolynomialPlot[355].X = 3.55
	pointsOfPolynomialPlot[355].Y = -5.098

	pointsOfPolynomialPlot[356].X = 3.56
	pointsOfPolynomialPlot[356].Y = -5.117

	pointsOfPolynomialPlot[357].X = 3.57
	pointsOfPolynomialPlot[357].Y = -5.137

	pointsOfPolynomialPlot[358].X = 3.58
	pointsOfPolynomialPlot[358].Y = -5.156

	pointsOfPolynomialPlot[359].X = 3.59
	pointsOfPolynomialPlot[359].Y = -5.175

	pointsOfPolynomialPlot[360].X = 3.6
	pointsOfPolynomialPlot[360].Y = -5.194

	pointsOfPolynomialPlot[361].X = 3.61
	pointsOfPolynomialPlot[361].Y = -5.213

	pointsOfPolynomialPlot[362].X = 3.62
	pointsOfPolynomialPlot[362].Y = -5.233

	pointsOfPolynomialPlot[363].X = 3.63
	pointsOfPolynomialPlot[363].Y = -5.252

	pointsOfPolynomialPlot[364].X = 3.64
	pointsOfPolynomialPlot[364].Y = -5.271

	pointsOfPolynomialPlot[365].X = 3.65
	pointsOfPolynomialPlot[365].Y = -5.291

	pointsOfPolynomialPlot[366].X = 3.66
	pointsOfPolynomialPlot[366].Y = -5.31

	pointsOfPolynomialPlot[367].X = 3.67
	pointsOfPolynomialPlot[367].Y = -5.329

	pointsOfPolynomialPlot[368].X = 3.68
	pointsOfPolynomialPlot[368].Y = -5.349

	pointsOfPolynomialPlot[369].X = 3.69
	pointsOfPolynomialPlot[369].Y = -5.368

	pointsOfPolynomialPlot[370].X = 3.7
	pointsOfPolynomialPlot[370].Y = -5.388

	pointsOfPolynomialPlot[371].X = 3.71
	pointsOfPolynomialPlot[371].Y = -5.407

	pointsOfPolynomialPlot[372].X = 3.72
	pointsOfPolynomialPlot[372].Y = -5.427

	pointsOfPolynomialPlot[373].X = 3.73
	pointsOfPolynomialPlot[373].Y = -5.446

	pointsOfPolynomialPlot[374].X = 3.74
	pointsOfPolynomialPlot[374].Y = -5.466

	pointsOfPolynomialPlot[375].X = 3.75
	pointsOfPolynomialPlot[375].Y = -5.485

	pointsOfPolynomialPlot[376].X = 3.76
	pointsOfPolynomialPlot[376].Y = -5.505

	pointsOfPolynomialPlot[377].X = 3.77
	pointsOfPolynomialPlot[377].Y = -5.525

	pointsOfPolynomialPlot[378].X = 3.78
	pointsOfPolynomialPlot[378].Y = -5.544

	pointsOfPolynomialPlot[379].X = 3.79
	pointsOfPolynomialPlot[379].Y = -5.564

	pointsOfPolynomialPlot[380].X = 3.8
	pointsOfPolynomialPlot[380].Y = -5.584

	pointsOfPolynomialPlot[381].X = 3.81
	pointsOfPolynomialPlot[381].Y = -5.604

	pointsOfPolynomialPlot[382].X = 3.82
	pointsOfPolynomialPlot[382].Y = -5.623

	pointsOfPolynomialPlot[383].X = 3.83
	pointsOfPolynomialPlot[383].Y = -5.643

	pointsOfPolynomialPlot[384].X = 3.84
	pointsOfPolynomialPlot[384].Y = -5.663

	pointsOfPolynomialPlot[385].X = 3.85
	pointsOfPolynomialPlot[385].Y = -5.683

	pointsOfPolynomialPlot[386].X = 3.86
	pointsOfPolynomialPlot[386].Y = -5.703

	pointsOfPolynomialPlot[387].X = 3.87
	pointsOfPolynomialPlot[387].Y = -5.723

	pointsOfPolynomialPlot[388].X = 3.88
	pointsOfPolynomialPlot[388].Y = -5.742

	pointsOfPolynomialPlot[389].X = 3.89
	pointsOfPolynomialPlot[389].Y = -5.762

	pointsOfPolynomialPlot[390].X = 3.9
	pointsOfPolynomialPlot[390].Y = -5.782

	pointsOfPolynomialPlot[391].X = 3.91
	pointsOfPolynomialPlot[391].Y = -5.802

	pointsOfPolynomialPlot[392].X = 3.92
	pointsOfPolynomialPlot[392].Y = -5.822

	pointsOfPolynomialPlot[393].X = 3.93
	pointsOfPolynomialPlot[393].Y = -5.842

	pointsOfPolynomialPlot[394].X = 3.94
	pointsOfPolynomialPlot[394].Y = -5.863

	pointsOfPolynomialPlot[395].X = 3.95
	pointsOfPolynomialPlot[395].Y = -5.883

	pointsOfPolynomialPlot[396].X = 3.96
	pointsOfPolynomialPlot[396].Y = -5.903

	pointsOfPolynomialPlot[397].X = 3.97
	pointsOfPolynomialPlot[397].Y = -5.923

	pointsOfPolynomialPlot[398].X = 3.98
	pointsOfPolynomialPlot[398].Y = -5.943

	pointsOfPolynomialPlot[399].X = 3.99
	pointsOfPolynomialPlot[399].Y = -5.963

	pointsOfPolynomialPlot[400].X = 4.0
	pointsOfPolynomialPlot[400].Y = -5.984

	pointsOfPolynomialPlot[401].X = 4.01
	pointsOfPolynomialPlot[401].Y = -6.004

	pointsOfPolynomialPlot[402].X = 4.02
	pointsOfPolynomialPlot[402].Y = -6.024

	pointsOfPolynomialPlot[403].X = 4.03
	pointsOfPolynomialPlot[403].Y = -6.044

	pointsOfPolynomialPlot[404].X = 4.04
	pointsOfPolynomialPlot[404].Y = -6.065

	pointsOfPolynomialPlot[405].X = 4.05
	pointsOfPolynomialPlot[405].Y = -6.085

	pointsOfPolynomialPlot[406].X = 4.06
	pointsOfPolynomialPlot[406].Y = -6.105

	pointsOfPolynomialPlot[407].X = 4.07
	pointsOfPolynomialPlot[407].Y = -6.126

	pointsOfPolynomialPlot[408].X = 4.08
	pointsOfPolynomialPlot[408].Y = -6.146

	pointsOfPolynomialPlot[409].X = 4.09
	pointsOfPolynomialPlot[409].Y = -6.167

	pointsOfPolynomialPlot[410].X = 4.1
	pointsOfPolynomialPlot[410].Y = -6.187

	pointsOfPolynomialPlot[411].X = 4.11
	pointsOfPolynomialPlot[411].Y = -6.208

	pointsOfPolynomialPlot[412].X = 4.12
	pointsOfPolynomialPlot[412].Y = -6.228

	pointsOfPolynomialPlot[413].X = 4.13
	pointsOfPolynomialPlot[413].Y = -6.249

	pointsOfPolynomialPlot[414].X = 4.14
	pointsOfPolynomialPlot[414].Y = -6.269

	pointsOfPolynomialPlot[415].X = 4.15
	pointsOfPolynomialPlot[415].Y = -6.29

	pointsOfPolynomialPlot[416].X = 4.16
	pointsOfPolynomialPlot[416].Y = -6.311

	pointsOfPolynomialPlot[417].X = 4.17
	pointsOfPolynomialPlot[417].Y = -6.331

	pointsOfPolynomialPlot[418].X = 4.18
	pointsOfPolynomialPlot[418].Y = -6.352

	pointsOfPolynomialPlot[419].X = 4.19
	pointsOfPolynomialPlot[419].Y = -6.373

	pointsOfPolynomialPlot[420].X = 4.2
	pointsOfPolynomialPlot[420].Y = -6.393

	pointsOfPolynomialPlot[421].X = 4.21
	pointsOfPolynomialPlot[421].Y = -6.414

	pointsOfPolynomialPlot[422].X = 4.22
	pointsOfPolynomialPlot[422].Y = -6.435

	pointsOfPolynomialPlot[423].X = 4.23
	pointsOfPolynomialPlot[423].Y = -6.456

	pointsOfPolynomialPlot[424].X = 4.24
	pointsOfPolynomialPlot[424].Y = -6.476

	pointsOfPolynomialPlot[425].X = 4.25
	pointsOfPolynomialPlot[425].Y = -6.497

	pointsOfPolynomialPlot[426].X = 4.26
	pointsOfPolynomialPlot[426].Y = -6.518

	pointsOfPolynomialPlot[427].X = 4.27
	pointsOfPolynomialPlot[427].Y = -6.539

	pointsOfPolynomialPlot[428].X = 4.28
	pointsOfPolynomialPlot[428].Y = -6.56

	pointsOfPolynomialPlot[429].X = 4.29
	pointsOfPolynomialPlot[429].Y = -6.581

	pointsOfPolynomialPlot[430].X = 4.30
	pointsOfPolynomialPlot[430].Y = -6.602

	pointsOfPolynomialPlot[431].X = 4.31
	pointsOfPolynomialPlot[431].Y = -6.623

	pointsOfPolynomialPlot[432].X = 4.32
	pointsOfPolynomialPlot[432].Y = -6.644

	pointsOfPolynomialPlot[433].X = 4.33
	pointsOfPolynomialPlot[433].Y = -6.665

	pointsOfPolynomialPlot[434].X = 4.34
	pointsOfPolynomialPlot[434].Y = -6.686

	pointsOfPolynomialPlot[435].X = 4.35
	pointsOfPolynomialPlot[435].Y = -6.707

	pointsOfPolynomialPlot[436].X = 4.36
	pointsOfPolynomialPlot[436].Y = -6.728

	pointsOfPolynomialPlot[437].X = 4.37
	pointsOfPolynomialPlot[437].Y = -6.749

	pointsOfPolynomialPlot[438].X = 4.38
	pointsOfPolynomialPlot[438].Y = -6.77

	pointsOfPolynomialPlot[439].X = 4.39
	pointsOfPolynomialPlot[439].Y = -6.791

	pointsOfPolynomialPlot[440].X = 4.4
	pointsOfPolynomialPlot[440].Y = -6.813

	pointsOfPolynomialPlot[441].X = 4.41
	pointsOfPolynomialPlot[441].Y = -6.834

	pointsOfPolynomialPlot[442].X = 4.42
	pointsOfPolynomialPlot[442].Y = -6.855

	pointsOfPolynomialPlot[443].X = 4.43
	pointsOfPolynomialPlot[443].Y = -6.876

	pointsOfPolynomialPlot[444].X = 4.44
	pointsOfPolynomialPlot[444].Y = -6.898

	pointsOfPolynomialPlot[445].X = 4.45
	pointsOfPolynomialPlot[445].Y = -6.919

	pointsOfPolynomialPlot[446].X = 4.46
	pointsOfPolynomialPlot[446].Y = -6.94

	pointsOfPolynomialPlot[447].X = 4.47
	pointsOfPolynomialPlot[447].Y = -6.962

	pointsOfPolynomialPlot[448].X = 4.48
	pointsOfPolynomialPlot[448].Y = -6.983

	pointsOfPolynomialPlot[449].X = 4.49
	pointsOfPolynomialPlot[449].Y = -7.004

	pointsOfPolynomialPlot[450].X = 4.5
	pointsOfPolynomialPlot[450].Y = -7.026

	pointsOfPolynomialPlot[451].X = 4.51
	pointsOfPolynomialPlot[451].Y = -7.047

	pointsOfPolynomialPlot[452].X = 4.52
	pointsOfPolynomialPlot[452].Y = -7.069

	pointsOfPolynomialPlot[453].X = 4.53
	pointsOfPolynomialPlot[453].Y = -7.09

	pointsOfPolynomialPlot[454].X = 4.54
	pointsOfPolynomialPlot[454].Y = -7.112

	pointsOfPolynomialPlot[455].X = 4.55
	pointsOfPolynomialPlot[455].Y = -7.133

	pointsOfPolynomialPlot[456].X = 4.56
	pointsOfPolynomialPlot[456].Y = -7.155

	pointsOfPolynomialPlot[457].X = 4.57
	pointsOfPolynomialPlot[457].Y = -7.176

	pointsOfPolynomialPlot[458].X = 4.58
	pointsOfPolynomialPlot[458].Y = -7.198

	pointsOfPolynomialPlot[459].X = 4.59
	pointsOfPolynomialPlot[459].Y = -7.22

	pointsOfPolynomialPlot[460].X = 4.6
	pointsOfPolynomialPlot[460].Y = -7.241

	pointsOfPolynomialPlot[461].X = 4.61
	pointsOfPolynomialPlot[461].Y = -7.263

	pointsOfPolynomialPlot[462].X = 4.62
	pointsOfPolynomialPlot[462].Y = -7.285

	pointsOfPolynomialPlot[463].X = 4.63
	pointsOfPolynomialPlot[463].Y = -7.306

	pointsOfPolynomialPlot[464].X = 4.64
	pointsOfPolynomialPlot[464].Y = -7.328

	pointsOfPolynomialPlot[465].X = 4.65
	pointsOfPolynomialPlot[465].Y = -7.35

	pointsOfPolynomialPlot[466].X = 4.66
	pointsOfPolynomialPlot[466].Y = -7.371

	pointsOfPolynomialPlot[467].X = 4.67
	pointsOfPolynomialPlot[467].Y = -7.393

	pointsOfPolynomialPlot[468].X = 4.68
	pointsOfPolynomialPlot[468].Y = -7.415

	pointsOfPolynomialPlot[469].X = 4.69
	pointsOfPolynomialPlot[469].Y = -7.437

	pointsOfPolynomialPlot[470].X = 4.7
	pointsOfPolynomialPlot[470].Y = -7.459

	pointsOfPolynomialPlot[471].X = 4.71
	pointsOfPolynomialPlot[471].Y = -7.481

	pointsOfPolynomialPlot[472].X = 4.72
	pointsOfPolynomialPlot[472].Y = -7.503

	pointsOfPolynomialPlot[473].X = 4.73
	pointsOfPolynomialPlot[473].Y = -7.525

	pointsOfPolynomialPlot[474].X = 4.74
	pointsOfPolynomialPlot[474].Y = -7.546

	pointsOfPolynomialPlot[475].X = 4.75
	pointsOfPolynomialPlot[475].Y = -7.568

	pointsOfPolynomialPlot[476].X = 4.76
	pointsOfPolynomialPlot[476].Y = -7.59

	pointsOfPolynomialPlot[477].X = 4.77
	pointsOfPolynomialPlot[477].Y = -7.612

	pointsOfPolynomialPlot[478].X = 4.78
	pointsOfPolynomialPlot[478].Y = -7.634

	pointsOfPolynomialPlot[479].X = 4.79
	pointsOfPolynomialPlot[479].Y = -7.657

	pointsOfPolynomialPlot[480].X = 4.8
	pointsOfPolynomialPlot[480].Y = -7.679

	pointsOfPolynomialPlot[481].X = 4.81
	pointsOfPolynomialPlot[481].Y = -7.701

	pointsOfPolynomialPlot[482].X = 4.82
	pointsOfPolynomialPlot[482].Y = -7.723

	pointsOfPolynomialPlot[483].X = 4.83
	pointsOfPolynomialPlot[483].Y = -7.745

	pointsOfPolynomialPlot[484].X = 4.84
	pointsOfPolynomialPlot[484].Y = -7.767

	pointsOfPolynomialPlot[485].X = 4.85
	pointsOfPolynomialPlot[485].Y = -7.789

	pointsOfPolynomialPlot[486].X = 4.86
	pointsOfPolynomialPlot[486].Y = -7.812

	pointsOfPolynomialPlot[487].X = 4.87
	pointsOfPolynomialPlot[487].Y = -7.834

	pointsOfPolynomialPlot[488].X = 4.88
	pointsOfPolynomialPlot[488].Y = -7.856

	pointsOfPolynomialPlot[489].X = 4.89
	pointsOfPolynomialPlot[489].Y = -7.878

	pointsOfPolynomialPlot[490].X = 4.9
	pointsOfPolynomialPlot[490].Y = -7.901

	pointsOfPolynomialPlot[491].X = 4.91
	pointsOfPolynomialPlot[491].Y = -7.923

	pointsOfPolynomialPlot[492].X = 4.92
	pointsOfPolynomialPlot[492].Y = -7.945

	pointsOfPolynomialPlot[493].X = 4.93
	pointsOfPolynomialPlot[493].Y = -7.968

	pointsOfPolynomialPlot[494].X = 4.94
	pointsOfPolynomialPlot[494].Y = -7.99

	pointsOfPolynomialPlot[495].X = 4.95
	pointsOfPolynomialPlot[495].Y = -8.012

	pointsOfPolynomialPlot[496].X = 4.96
	pointsOfPolynomialPlot[496].Y = -8.035

	pointsOfPolynomialPlot[497].X = 4.97
	pointsOfPolynomialPlot[497].Y = -8.057

	pointsOfPolynomialPlot[498].X = 4.98
	pointsOfPolynomialPlot[498].Y = -8.08

	pointsOfPolynomialPlot[499].X = 4.99
	pointsOfPolynomialPlot[499].Y = -8.102

	pointsOfPolynomialPlot[500].X = 5.0
	pointsOfPolynomialPlot[500].Y = -8.125

	pointsOfPolynomialPlot[501].X = 5.01
	pointsOfPolynomialPlot[501].Y = -8.147

	pointsOfPolynomialPlot[502].X = 5.02
	pointsOfPolynomialPlot[502].Y = -8.17

	pointsOfPolynomialPlot[503].X = 5.03
	pointsOfPolynomialPlot[503].Y = -8.192

	pointsOfPolynomialPlot[504].X = 5.04
	pointsOfPolynomialPlot[504].Y = -8.215

	pointsOfPolynomialPlot[505].X = 5.05
	pointsOfPolynomialPlot[505].Y = -8.237

	pointsOfPolynomialPlot[506].X = 5.06
	pointsOfPolynomialPlot[506].Y = -8.26

	pointsOfPolynomialPlot[507].X = 5.07
	pointsOfPolynomialPlot[507].Y = -8.283

	pointsOfPolynomialPlot[508].X = 5.08
	pointsOfPolynomialPlot[508].Y = -8.305

	pointsOfPolynomialPlot[509].X = 5.09
	pointsOfPolynomialPlot[509].Y = -8.328

	pointsOfPolynomialPlot[510].X = 5.10
	pointsOfPolynomialPlot[510].Y = -8.351

	pointsOfPolynomialPlot[511].X = 5.11
	pointsOfPolynomialPlot[511].Y = -8.373

	pointsOfPolynomialPlot[512].X = 5.12
	pointsOfPolynomialPlot[512].Y = -8.396

	pointsOfPolynomialPlot[513].X = 5.13
	pointsOfPolynomialPlot[513].Y = -8.419

	pointsOfPolynomialPlot[514].X = 5.14
	pointsOfPolynomialPlot[514].Y = -8.441

	pointsOfPolynomialPlot[515].X = 5.15
	pointsOfPolynomialPlot[515].Y = -8.464

	pointsOfPolynomialPlot[516].X = 5.16
	pointsOfPolynomialPlot[516].Y = -8.487

	pointsOfPolynomialPlot[517].X = 5.17
	pointsOfPolynomialPlot[517].Y = -8.51

	pointsOfPolynomialPlot[518].X = 5.18
	pointsOfPolynomialPlot[518].Y = -8.533

	pointsOfPolynomialPlot[519].X = 5.19
	pointsOfPolynomialPlot[519].Y = -8.556

	pointsOfPolynomialPlot[520].X = 5.2
	pointsOfPolynomialPlot[520].Y = -8.578

	pointsOfPolynomialPlot[521].X = 5.21
	pointsOfPolynomialPlot[521].Y = -8.601

	pointsOfPolynomialPlot[522].X = 5.22
	pointsOfPolynomialPlot[522].Y = -8.624

	pointsOfPolynomialPlot[523].X = 5.23
	pointsOfPolynomialPlot[523].Y = -8.647

	pointsOfPolynomialPlot[524].X = 5.24
	pointsOfPolynomialPlot[524].Y = -8.67

	pointsOfPolynomialPlot[525].X = 5.25
	pointsOfPolynomialPlot[525].Y = -8.693

	pointsOfPolynomialPlot[526].X = 5.26
	pointsOfPolynomialPlot[526].Y = -8.716

	pointsOfPolynomialPlot[527].X = 5.27
	pointsOfPolynomialPlot[527].Y = -8.739

	pointsOfPolynomialPlot[528].X = 5.28
	pointsOfPolynomialPlot[528].Y = -8.762

	pointsOfPolynomialPlot[529].X = 5.29
	pointsOfPolynomialPlot[529].Y = -8.785

	pointsOfPolynomialPlot[530].X = 5.3
	pointsOfPolynomialPlot[530].Y = -8.808

	pointsOfPolynomialPlot[531].X = 5.31
	pointsOfPolynomialPlot[531].Y = -8.831

	pointsOfPolynomialPlot[532].X = 5.32
	pointsOfPolynomialPlot[532].Y = -8.854

	pointsOfPolynomialPlot[533].X = 5.33
	pointsOfPolynomialPlot[533].Y = -8.878

	pointsOfPolynomialPlot[534].X = 5.34
	pointsOfPolynomialPlot[534].Y = -8.901

	pointsOfPolynomialPlot[535].X = 5.35
	pointsOfPolynomialPlot[535].Y = -8.924

	pointsOfPolynomialPlot[536].X = 5.36
	pointsOfPolynomialPlot[536].Y = -8.947

	pointsOfPolynomialPlot[537].X = 5.37
	pointsOfPolynomialPlot[537].Y = -8.97

	pointsOfPolynomialPlot[538].X = 5.38
	pointsOfPolynomialPlot[538].Y = -8.993

	pointsOfPolynomialPlot[539].X = 5.39
	pointsOfPolynomialPlot[539].Y = -9.017

	pointsOfPolynomialPlot[540].X = 5.4
	pointsOfPolynomialPlot[540].Y = -9.04

	pointsOfPolynomialPlot[541].X = 5.41
	pointsOfPolynomialPlot[541].Y = -9.063

	pointsOfPolynomialPlot[542].X = 5.42
	pointsOfPolynomialPlot[542].Y = -9.086

	pointsOfPolynomialPlot[543].X = 5.43
	pointsOfPolynomialPlot[543].Y = -9.11

	pointsOfPolynomialPlot[544].X = 5.44
	pointsOfPolynomialPlot[544].Y = -9.133

	pointsOfPolynomialPlot[545].X = 5.45
	pointsOfPolynomialPlot[545].Y = -9.156

	pointsOfPolynomialPlot[546].X = 5.46
	pointsOfPolynomialPlot[546].Y = -9.18

	pointsOfPolynomialPlot[547].X = 5.47
	pointsOfPolynomialPlot[547].Y = -9.203

	pointsOfPolynomialPlot[548].X = 5.48
	pointsOfPolynomialPlot[548].Y = -9.226

	pointsOfPolynomialPlot[549].X = 5.49
	pointsOfPolynomialPlot[549].Y = -9.25

	pointsOfPolynomialPlot[550].X = 5.5
	pointsOfPolynomialPlot[550].Y = -9.273







































































































































































































































































































	wykresWielomianu := plot.New()

	wykresWielomianu.Title.Text = "Wykres funkcji f(x) = 0.001x^4 - 0.01x^3 - 0.1x^2 + x"

	wykresWielomianu.X.Label.Text = "x"
	wykresWielomianu.Y.Label.Text = "y"

	liniaWykresu, err := plotter.NewLine(pointsOfPolynomialPlot)

	if err != nil {
		panic(err)
	}

	liniaWykresu.LineStyle.Width = vg.Points(0.1)
	liniaWykresu.Color = color.RGBA{R: 200, G: 100, B: 100}

	wykresWielomianu.Add(liniaWykresu)
	wykresWielomianu.Legend.Add("f(x)", liniaWykresu)

	if err := wykresWielomianu.Save(10*vg.Inch, 10*vg.Inch,
		"Wykres-wielomianu-05.png"); err != nil {

		panic(err)
	}
}
