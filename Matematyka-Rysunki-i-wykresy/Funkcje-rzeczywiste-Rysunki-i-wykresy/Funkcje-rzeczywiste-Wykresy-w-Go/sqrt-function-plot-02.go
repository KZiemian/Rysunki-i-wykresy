package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function sqrt(x).

	pointsOfFunctionPlot := make(plotter.XYs, 2_001)

	pointsOfFunctionPlot[0].X = 0.0
	pointsOfFunctionPlot[0].Y = 0.0

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = 0.1

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = 0.141

	pointsOfFunctionPlot[3].X = 0.03
	pointsOfFunctionPlot[3].Y = 0.173

	pointsOfFunctionPlot[4].X = 0.04
	pointsOfFunctionPlot[4].Y = 0.2

	pointsOfFunctionPlot[5].X = 0.05
	pointsOfFunctionPlot[5].Y = 0.223

	pointsOfFunctionPlot[6].X = 0.06
	pointsOfFunctionPlot[6].Y = 0.244

	pointsOfFunctionPlot[7].X = 0.07
	pointsOfFunctionPlot[7].Y = 0.264

	pointsOfFunctionPlot[8].X = 0.08
	pointsOfFunctionPlot[8].Y = 0.283

	pointsOfFunctionPlot[9].X = 0.09
	pointsOfFunctionPlot[9].Y = 0.3

	pointsOfFunctionPlot[10].X = 0.10
	pointsOfFunctionPlot[10].Y = 0.316

	pointsOfFunctionPlot[11].X = 0.11
	pointsOfFunctionPlot[11].Y = 0.331

	pointsOfFunctionPlot[12].X = 0.12
	pointsOfFunctionPlot[12].Y = 0.346

	pointsOfFunctionPlot[13].X = 0.13
	pointsOfFunctionPlot[13].Y = 0.36

	pointsOfFunctionPlot[14].X = 0.14
	pointsOfFunctionPlot[14].Y = 0.374

	pointsOfFunctionPlot[15].X = 0.15
	pointsOfFunctionPlot[15].Y = 0.387

	pointsOfFunctionPlot[16].X = 0.16
	pointsOfFunctionPlot[16].Y = 0.4

	pointsOfFunctionPlot[17].X = 0.17
	pointsOfFunctionPlot[17].Y = 0.412

	pointsOfFunctionPlot[18].X = 0.18
	pointsOfFunctionPlot[18].Y = 0.424

	pointsOfFunctionPlot[19].X = 0.19
	pointsOfFunctionPlot[19].Y = 0.435

	pointsOfFunctionPlot[20].X = 0.20
	pointsOfFunctionPlot[20].Y = 0.447

	pointsOfFunctionPlot[21].X = 0.21
	pointsOfFunctionPlot[21].Y = 0.458

	pointsOfFunctionPlot[22].X = 0.22
	pointsOfFunctionPlot[22].Y = 0.469

	pointsOfFunctionPlot[23].X = 0.23
	pointsOfFunctionPlot[23].Y = 0.479

	pointsOfFunctionPlot[24].X = 0.24
	pointsOfFunctionPlot[24].Y = 0.489

	pointsOfFunctionPlot[25].X = 0.25
	pointsOfFunctionPlot[25].Y = 0.5

	pointsOfFunctionPlot[26].X = 0.26
	pointsOfFunctionPlot[26].Y = 0.509

	pointsOfFunctionPlot[27].X = 0.27
	pointsOfFunctionPlot[27].Y = 0.519

	pointsOfFunctionPlot[28].X = 0.28
	pointsOfFunctionPlot[28].Y = 0.529

	pointsOfFunctionPlot[29].X = 0.29
	pointsOfFunctionPlot[29].Y = 0.538

	pointsOfFunctionPlot[30].X = 0.30
	pointsOfFunctionPlot[30].Y = 0.547

	pointsOfFunctionPlot[31].X = 0.31
	pointsOfFunctionPlot[31].Y = 0.556

	pointsOfFunctionPlot[32].X = 0.32
	pointsOfFunctionPlot[32].Y = 0.565

	pointsOfFunctionPlot[33].X = 0.33
	pointsOfFunctionPlot[33].Y = 0.574

	pointsOfFunctionPlot[34].X = 0.34
	pointsOfFunctionPlot[34].Y = 0.583

	pointsOfFunctionPlot[35].X = 0.35
	pointsOfFunctionPlot[35].Y = 0.591

	pointsOfFunctionPlot[36].X = 0.36
	pointsOfFunctionPlot[36].Y = 0.6

	pointsOfFunctionPlot[37].X = 0.37
	pointsOfFunctionPlot[37].Y = 0.608

	pointsOfFunctionPlot[38].X = 0.38
	pointsOfFunctionPlot[38].Y = 0.616

	pointsOfFunctionPlot[39].X = 0.39
	pointsOfFunctionPlot[39].Y = 0.624

	pointsOfFunctionPlot[40].X = 0.40
	pointsOfFunctionPlot[40].Y = 0.632

	pointsOfFunctionPlot[41].X = 0.41
	pointsOfFunctionPlot[41].Y = 0.64

	pointsOfFunctionPlot[42].X = 0.42
	pointsOfFunctionPlot[42].Y = 0.648

	pointsOfFunctionPlot[43].X = 0.43
	pointsOfFunctionPlot[43].Y = 0.655

	pointsOfFunctionPlot[44].X = 0.44
	pointsOfFunctionPlot[44].Y = 0.663

	pointsOfFunctionPlot[45].X = 0.45
	pointsOfFunctionPlot[45].Y = 0.67

	pointsOfFunctionPlot[46].X = 0.46
	pointsOfFunctionPlot[46].Y = 0.678

	pointsOfFunctionPlot[47].X = 0.47
	pointsOfFunctionPlot[47].Y = 0.685

	pointsOfFunctionPlot[48].X = 0.48
	pointsOfFunctionPlot[48].Y = 0.692

	pointsOfFunctionPlot[49].X = 0.49
	pointsOfFunctionPlot[49].Y = 0.7

	pointsOfFunctionPlot[50].X = 0.50
	pointsOfFunctionPlot[50].Y = 0.707

	pointsOfFunctionPlot[51].X = 0.51
	pointsOfFunctionPlot[51].Y = 0.714

	pointsOfFunctionPlot[52].X = 0.52
	pointsOfFunctionPlot[52].Y = 0.721

	pointsOfFunctionPlot[53].X = 0.53
	pointsOfFunctionPlot[53].Y = 0.728

	pointsOfFunctionPlot[54].X = 0.54
	pointsOfFunctionPlot[54].Y = 0.734

	pointsOfFunctionPlot[55].X = 0.55
	pointsOfFunctionPlot[55].Y = 0.741

	pointsOfFunctionPlot[56].X = 0.56
	pointsOfFunctionPlot[56].Y = 0.748

	pointsOfFunctionPlot[57].X = 0.57
	pointsOfFunctionPlot[57].Y = 0.754

	pointsOfFunctionPlot[58].X = 0.58
	pointsOfFunctionPlot[58].Y = 0.761

	pointsOfFunctionPlot[59].X = 0.59
	pointsOfFunctionPlot[59].Y = 0.768

	pointsOfFunctionPlot[60].X = 0.60
	pointsOfFunctionPlot[60].Y = 0.774

	pointsOfFunctionPlot[61].X = 0.61
	pointsOfFunctionPlot[61].Y = 0.781

	pointsOfFunctionPlot[62].X = 0.62
	pointsOfFunctionPlot[62].Y = 0.787

	pointsOfFunctionPlot[63].X = 0.63
	pointsOfFunctionPlot[63].Y = 0.793

	pointsOfFunctionPlot[64].X = 0.64
	pointsOfFunctionPlot[64].Y = 0.8

	pointsOfFunctionPlot[65].X = 0.65
	pointsOfFunctionPlot[65].Y = 0.806

	pointsOfFunctionPlot[66].X = 0.66
	pointsOfFunctionPlot[66].Y = 0.812

	pointsOfFunctionPlot[67].X = 0.67
	pointsOfFunctionPlot[67].Y = 0.818

	pointsOfFunctionPlot[68].X = 0.68
	pointsOfFunctionPlot[68].Y = 0.824

	pointsOfFunctionPlot[69].X = 0.69
	pointsOfFunctionPlot[69].Y = 0.83

	pointsOfFunctionPlot[70].X = 0.70
	pointsOfFunctionPlot[70].Y = 0.836

	pointsOfFunctionPlot[71].X = 0.71
	pointsOfFunctionPlot[71].Y = 0.842

	pointsOfFunctionPlot[72].X = 0.72
	pointsOfFunctionPlot[72].Y = 0.848

	pointsOfFunctionPlot[73].X = 0.73
	pointsOfFunctionPlot[73].Y = 0.854

	pointsOfFunctionPlot[74].X = 0.74
	pointsOfFunctionPlot[74].Y = 0.86

	pointsOfFunctionPlot[75].X = 0.75
	pointsOfFunctionPlot[75].Y = 0.866

	pointsOfFunctionPlot[76].X = 0.76
	pointsOfFunctionPlot[76].Y = 0.871

	pointsOfFunctionPlot[77].X = 0.77
	pointsOfFunctionPlot[77].Y = 0.877

	pointsOfFunctionPlot[78].X = 0.78
	pointsOfFunctionPlot[78].Y = 0.883

	pointsOfFunctionPlot[79].X = 0.79
	pointsOfFunctionPlot[79].Y = 0.888

	pointsOfFunctionPlot[80].X = 0.80
	pointsOfFunctionPlot[80].Y = 0.894

	pointsOfFunctionPlot[81].X = 0.81
	pointsOfFunctionPlot[81].Y = 0.9

	pointsOfFunctionPlot[82].X = 0.82
	pointsOfFunctionPlot[82].Y = 0.905

	pointsOfFunctionPlot[83].X = 0.83
	pointsOfFunctionPlot[83].Y = 0.911

	pointsOfFunctionPlot[84].X = 0.84
	pointsOfFunctionPlot[84].Y = 0.916

	pointsOfFunctionPlot[85].X = 0.85
	pointsOfFunctionPlot[85].Y = 0.921

	pointsOfFunctionPlot[86].X = 0.86
	pointsOfFunctionPlot[86].Y = 0.927

	pointsOfFunctionPlot[87].X = 0.87
	pointsOfFunctionPlot[87].Y = 0.932

	pointsOfFunctionPlot[88].X = 0.88
	pointsOfFunctionPlot[88].Y = 0.938

	pointsOfFunctionPlot[89].X = 0.89
	pointsOfFunctionPlot[89].Y = 0.943

	pointsOfFunctionPlot[90].X = 0.90
	pointsOfFunctionPlot[90].Y = 0.948

	pointsOfFunctionPlot[91].X = 0.91
	pointsOfFunctionPlot[91].Y = 0.953

	pointsOfFunctionPlot[92].X = 0.92
	pointsOfFunctionPlot[92].Y = 0.959

	pointsOfFunctionPlot[93].X = 0.93
	pointsOfFunctionPlot[93].Y = 0.964

	pointsOfFunctionPlot[94].X = 0.94
	pointsOfFunctionPlot[94].Y = 0.969

	pointsOfFunctionPlot[95].X = 0.95
	pointsOfFunctionPlot[95].Y = 0.974

	pointsOfFunctionPlot[96].X = 0.96
	pointsOfFunctionPlot[96].Y = 0.979

	pointsOfFunctionPlot[97].X = 0.97
	pointsOfFunctionPlot[97].Y = 0.984

	pointsOfFunctionPlot[98].X = 0.98
	pointsOfFunctionPlot[98].Y = 0.989

	pointsOfFunctionPlot[99].X = 0.99
	pointsOfFunctionPlot[99].Y = 0.994

	pointsOfFunctionPlot[100].X = 1.0
	pointsOfFunctionPlot[100].Y = 1.0

	pointsOfFunctionPlot[101].X = 1.01
	pointsOfFunctionPlot[101].Y = 1.004

	pointsOfFunctionPlot[102].X = 1.02
	pointsOfFunctionPlot[102].Y = 1.009

	pointsOfFunctionPlot[103].X = 1.03
	pointsOfFunctionPlot[103].Y = 1.014

	pointsOfFunctionPlot[104].X = 1.04
	pointsOfFunctionPlot[104].Y = 1.019

	pointsOfFunctionPlot[105].X = 1.05
	pointsOfFunctionPlot[105].Y = 1.024

	pointsOfFunctionPlot[106].X = 1.06
	pointsOfFunctionPlot[106].Y = 1.029

	pointsOfFunctionPlot[107].X = 1.07
	pointsOfFunctionPlot[107].Y = 1.034

	pointsOfFunctionPlot[108].X = 1.08
	pointsOfFunctionPlot[108].Y = 1.039

	pointsOfFunctionPlot[109].X = 1.09
	pointsOfFunctionPlot[109].Y = 1.044

	pointsOfFunctionPlot[110].X = 1.10
	pointsOfFunctionPlot[110].Y = 1.048

	pointsOfFunctionPlot[111].X = 1.11
	pointsOfFunctionPlot[111].Y = 1.053

	pointsOfFunctionPlot[112].X = 1.12
	pointsOfFunctionPlot[112].Y = 1.058

	pointsOfFunctionPlot[113].X = 1.13
	pointsOfFunctionPlot[113].Y = 1.063

	pointsOfFunctionPlot[114].X = 1.14
	pointsOfFunctionPlot[114].Y = 1.067

	pointsOfFunctionPlot[115].X = 1.15
	pointsOfFunctionPlot[115].Y = 1.072

	pointsOfFunctionPlot[116].X = 1.16
	pointsOfFunctionPlot[116].Y = 1.077

	pointsOfFunctionPlot[117].X = 1.17
	pointsOfFunctionPlot[117].Y = 1.081

	pointsOfFunctionPlot[118].X = 1.18
	pointsOfFunctionPlot[118].Y = 1.086

	pointsOfFunctionPlot[119].X = 1.19
	pointsOfFunctionPlot[119].Y = 1.09

	pointsOfFunctionPlot[120].X = 1.20
	pointsOfFunctionPlot[120].Y = 1.095

	pointsOfFunctionPlot[121].X = 1.21
	pointsOfFunctionPlot[121].Y = 1.1

	pointsOfFunctionPlot[122].X = 1.22
	pointsOfFunctionPlot[122].Y = 1.104

	pointsOfFunctionPlot[123].X = 1.23
	pointsOfFunctionPlot[123].Y = 1.109

	pointsOfFunctionPlot[124].X = 1.24
	pointsOfFunctionPlot[124].Y = 1.113

	pointsOfFunctionPlot[125].X = 1.25
	pointsOfFunctionPlot[125].Y = 1.118

	pointsOfFunctionPlot[126].X = 1.26
	pointsOfFunctionPlot[126].Y = 1.122

	pointsOfFunctionPlot[127].X = 1.27
	pointsOfFunctionPlot[127].Y = 1.126

	pointsOfFunctionPlot[128].X = 1.28
	pointsOfFunctionPlot[128].Y = 1.131

	pointsOfFunctionPlot[129].X = 1.29
	pointsOfFunctionPlot[129].Y = 1.135

	pointsOfFunctionPlot[130].X = 1.30
	pointsOfFunctionPlot[130].Y = 1.14

	pointsOfFunctionPlot[131].X = 1.31
	pointsOfFunctionPlot[131].Y = 1.144

	pointsOfFunctionPlot[132].X = 1.32
	pointsOfFunctionPlot[132].Y = 1.148

	pointsOfFunctionPlot[133].X = 1.33
	pointsOfFunctionPlot[133].Y = 1.153

	pointsOfFunctionPlot[134].X = 1.34
	pointsOfFunctionPlot[134].Y = 1.157

	pointsOfFunctionPlot[135].X = 1.35
	pointsOfFunctionPlot[135].Y = 1.161

	pointsOfFunctionPlot[136].X = 1.36
	pointsOfFunctionPlot[136].Y = 1.166

	pointsOfFunctionPlot[137].X = 1.37
	pointsOfFunctionPlot[137].Y = 1.17

	pointsOfFunctionPlot[138].X = 1.38
	pointsOfFunctionPlot[138].Y = 1.174

	pointsOfFunctionPlot[139].X = 1.39
	pointsOfFunctionPlot[139].Y = 1.178

	pointsOfFunctionPlot[140].X = 1.40
	pointsOfFunctionPlot[140].Y = 1.183

	pointsOfFunctionPlot[141].X = 1.41
	pointsOfFunctionPlot[141].Y = 1.187

	pointsOfFunctionPlot[142].X = 1.42
	pointsOfFunctionPlot[142].Y = 1.191

	pointsOfFunctionPlot[143].X = 1.43
	pointsOfFunctionPlot[143].Y = 1.195

	pointsOfFunctionPlot[144].X = 1.44
	pointsOfFunctionPlot[144].Y = 1.2

	pointsOfFunctionPlot[145].X = 1.45
	pointsOfFunctionPlot[145].Y = 1.204

	pointsOfFunctionPlot[146].X = 1.46
	pointsOfFunctionPlot[146].Y = 1.208

	pointsOfFunctionPlot[147].X = 1.47
	pointsOfFunctionPlot[147].Y = 1.212

	pointsOfFunctionPlot[148].X = 1.48
	pointsOfFunctionPlot[148].Y = 1.216

	pointsOfFunctionPlot[149].X = 1.49
	pointsOfFunctionPlot[149].Y = 1.22

	pointsOfFunctionPlot[150].X = 1.50
	pointsOfFunctionPlot[150].Y = 1.224

	pointsOfFunctionPlot[151].X = 1.51
	pointsOfFunctionPlot[151].Y = 1.228

	pointsOfFunctionPlot[152].X = 1.52
	pointsOfFunctionPlot[152].Y = 1.232

	pointsOfFunctionPlot[153].X = 1.53
	pointsOfFunctionPlot[153].Y = 1.236

	pointsOfFunctionPlot[154].X = 1.54
	pointsOfFunctionPlot[154].Y = 1.24

	pointsOfFunctionPlot[155].X = 1.55
	pointsOfFunctionPlot[155].Y = 1.244

	pointsOfFunctionPlot[156].X = 1.56
	pointsOfFunctionPlot[156].Y = 1.248

	pointsOfFunctionPlot[157].X = 1.57
	pointsOfFunctionPlot[157].Y = 1.252

	pointsOfFunctionPlot[158].X = 1.58
	pointsOfFunctionPlot[158].Y = 1.256

	pointsOfFunctionPlot[159].X = 1.59
	pointsOfFunctionPlot[159].Y = 1.26

	pointsOfFunctionPlot[160].X = 1.60
	pointsOfFunctionPlot[160].Y = 1.264

	pointsOfFunctionPlot[161].X = 1.61
	pointsOfFunctionPlot[161].Y = 1.268

	pointsOfFunctionPlot[162].X = 1.62
	pointsOfFunctionPlot[162].Y = 1.272

	pointsOfFunctionPlot[163].X = 1.63
	pointsOfFunctionPlot[163].Y = 1.276

	pointsOfFunctionPlot[164].X = 1.64
	pointsOfFunctionPlot[164].Y = 1.28

	pointsOfFunctionPlot[165].X = 1.65
	pointsOfFunctionPlot[165].Y = 1.284

	pointsOfFunctionPlot[166].X = 1.66
	pointsOfFunctionPlot[166].Y = 1.288

	pointsOfFunctionPlot[167].X = 1.67
	pointsOfFunctionPlot[167].Y = 1.292

	pointsOfFunctionPlot[168].X = 1.68
	pointsOfFunctionPlot[168].Y = 1.296

	pointsOfFunctionPlot[169].X = 1.69
	pointsOfFunctionPlot[169].Y = 1.3

	pointsOfFunctionPlot[170].X = 1.70
	pointsOfFunctionPlot[170].Y = 1.303

	pointsOfFunctionPlot[171].X = 1.71
	pointsOfFunctionPlot[171].Y = 1.307

	pointsOfFunctionPlot[172].X = 1.72
	pointsOfFunctionPlot[172].Y = 1.311

	pointsOfFunctionPlot[173].X = 1.73
	pointsOfFunctionPlot[173].Y = 1.315

	pointsOfFunctionPlot[174].X = 1.74
	pointsOfFunctionPlot[174].Y = 1.319

	pointsOfFunctionPlot[175].X = 1.75
	pointsOfFunctionPlot[175].Y = 1.322

	pointsOfFunctionPlot[176].X = 1.76
	pointsOfFunctionPlot[176].Y = 1.326

	pointsOfFunctionPlot[177].X = 1.77
	pointsOfFunctionPlot[177].Y = 1.33

	pointsOfFunctionPlot[178].X = 1.78
	pointsOfFunctionPlot[178].Y = 1.334

	pointsOfFunctionPlot[179].X = 1.79
	pointsOfFunctionPlot[179].Y = 1.337

	pointsOfFunctionPlot[180].X = 1.80
	pointsOfFunctionPlot[180].Y = 1.341

	pointsOfFunctionPlot[181].X = 1.81
	pointsOfFunctionPlot[181].Y = 1.345

	pointsOfFunctionPlot[182].X = 1.82
	pointsOfFunctionPlot[182].Y = 1.349

	pointsOfFunctionPlot[183].X = 1.83
	pointsOfFunctionPlot[183].Y = 1.352

	pointsOfFunctionPlot[184].X = 1.84
	pointsOfFunctionPlot[184].Y = 1.356

	pointsOfFunctionPlot[185].X = 1.85
	pointsOfFunctionPlot[185].Y = 1.36

	pointsOfFunctionPlot[186].X = 1.86
	pointsOfFunctionPlot[186].Y = 1.363

	pointsOfFunctionPlot[187].X = 1.87
	pointsOfFunctionPlot[187].Y = 1.367

	pointsOfFunctionPlot[188].X = 1.88
	pointsOfFunctionPlot[188].Y = 1.371

	pointsOfFunctionPlot[189].X = 1.89
	pointsOfFunctionPlot[189].Y = 1.374

	pointsOfFunctionPlot[190].X = 1.90
	pointsOfFunctionPlot[190].Y = 1.378

	pointsOfFunctionPlot[191].X = 1.91
	pointsOfFunctionPlot[191].Y = 1.383

	pointsOfFunctionPlot[192].X = 1.92
	pointsOfFunctionPlot[192].Y = 1.385

	pointsOfFunctionPlot[193].X = 1.93
	pointsOfFunctionPlot[193].Y = 1.389

	pointsOfFunctionPlot[194].X = 1.94
	pointsOfFunctionPlot[194].Y = 1.392

	pointsOfFunctionPlot[195].X = 1.95
	pointsOfFunctionPlot[195].Y = 1.396

	pointsOfFunctionPlot[196].X = 1.96
	pointsOfFunctionPlot[196].Y = 1.4

	pointsOfFunctionPlot[197].X = 1.97
	pointsOfFunctionPlot[197].Y = 1.403

	pointsOfFunctionPlot[198].X = 1.98
	pointsOfFunctionPlot[198].Y = 1.407

	pointsOfFunctionPlot[199].X = 1.99
	pointsOfFunctionPlot[199].Y = 1.41

	pointsOfFunctionPlot[200].X = 2.0
	pointsOfFunctionPlot[200].Y = 1.414

	pointsOfFunctionPlot[201].X = 2.01
	pointsOfFunctionPlot[201].Y = 1.417

	pointsOfFunctionPlot[202].X = 2.02
	pointsOfFunctionPlot[202].Y = 1.421

	pointsOfFunctionPlot[203].X = 2.03
	pointsOfFunctionPlot[203].Y = 1.424

	pointsOfFunctionPlot[204].X = 2.04
	pointsOfFunctionPlot[204].Y = 1.428

	pointsOfFunctionPlot[205].X = 2.05
	pointsOfFunctionPlot[205].Y = 1.431

	pointsOfFunctionPlot[206].X = 2.06
	pointsOfFunctionPlot[206].Y = 1.435

	pointsOfFunctionPlot[207].X = 2.07
	pointsOfFunctionPlot[207].Y = 1.438

	pointsOfFunctionPlot[208].X = 2.08
	pointsOfFunctionPlot[208].Y = 1.442

	pointsOfFunctionPlot[209].X = 2.09
	pointsOfFunctionPlot[209].Y = 1.445

	pointsOfFunctionPlot[210].X = 2.10
	pointsOfFunctionPlot[210].Y = 1.449

	pointsOfFunctionPlot[211].X = 2.11
	pointsOfFunctionPlot[211].Y = 1.452

	pointsOfFunctionPlot[212].X = 2.12
	pointsOfFunctionPlot[212].Y = 1.456

	pointsOfFunctionPlot[213].X = 2.13
	pointsOfFunctionPlot[213].Y = 1.459

	pointsOfFunctionPlot[214].X = 2.14
	pointsOfFunctionPlot[214].Y = 1.462

	pointsOfFunctionPlot[215].X = 2.15
	pointsOfFunctionPlot[215].Y = 1.466

	pointsOfFunctionPlot[216].X = 2.16
	pointsOfFunctionPlot[216].Y = 1.469

	pointsOfFunctionPlot[217].X = 2.17
	pointsOfFunctionPlot[217].Y = 1.473

	pointsOfFunctionPlot[218].X = 2.18
	pointsOfFunctionPlot[218].Y = 1.476

	pointsOfFunctionPlot[219].X = 2.19
	pointsOfFunctionPlot[219].Y = 1.479

	pointsOfFunctionPlot[220].X = 2.20
	pointsOfFunctionPlot[220].Y = 1.483

	pointsOfFunctionPlot[221].X = 2.21
	pointsOfFunctionPlot[221].Y = 1.486

	pointsOfFunctionPlot[222].X = 2.22
	pointsOfFunctionPlot[222].Y = 1.489

	pointsOfFunctionPlot[223].X = 2.23
	pointsOfFunctionPlot[223].Y = 1.493

	pointsOfFunctionPlot[224].X = 2.24
	pointsOfFunctionPlot[224].Y = 1.496

	pointsOfFunctionPlot[225].X = 2.25
	pointsOfFunctionPlot[225].Y = 1.5

	pointsOfFunctionPlot[226].X = 2.26
	pointsOfFunctionPlot[226].Y = 1.503

	pointsOfFunctionPlot[227].X = 2.27
	pointsOfFunctionPlot[227].Y = 1.506

	pointsOfFunctionPlot[228].X = 2.28
	pointsOfFunctionPlot[228].Y = 1.509

	pointsOfFunctionPlot[229].X = 2.29
	pointsOfFunctionPlot[229].Y = 1.513

	pointsOfFunctionPlot[230].X = 2.30
	pointsOfFunctionPlot[230].Y = 1.516

	pointsOfFunctionPlot[231].X = 2.31
	pointsOfFunctionPlot[231].Y = 1.519

	pointsOfFunctionPlot[232].X = 2.32
	pointsOfFunctionPlot[232].Y = 1.523

	pointsOfFunctionPlot[233].X = 2.33
	pointsOfFunctionPlot[233].Y = 1.526

	pointsOfFunctionPlot[234].X = 2.34
	pointsOfFunctionPlot[234].Y = 1.529

	pointsOfFunctionPlot[235].X = 2.35
	pointsOfFunctionPlot[235].Y = 1.532

	pointsOfFunctionPlot[236].X = 2.36
	pointsOfFunctionPlot[236].Y = 1.536

	pointsOfFunctionPlot[237].X = 2.37
	pointsOfFunctionPlot[237].Y = 1.539

	pointsOfFunctionPlot[238].X = 2.38
	pointsOfFunctionPlot[238].Y = 1.542

	pointsOfFunctionPlot[239].X = 2.39
	pointsOfFunctionPlot[239].Y = 1.545

	pointsOfFunctionPlot[240].X = 2.40
	pointsOfFunctionPlot[240].Y = 1.549

	pointsOfFunctionPlot[241].X = 2.41
	pointsOfFunctionPlot[241].Y = 1.552

	pointsOfFunctionPlot[242].X = 2.42
	pointsOfFunctionPlot[242].Y = 1.555

	pointsOfFunctionPlot[243].X = 2.43
	pointsOfFunctionPlot[243].Y = 1.558

	pointsOfFunctionPlot[244].X = 2.44
	pointsOfFunctionPlot[244].Y = 1.562

	pointsOfFunctionPlot[245].X = 2.45
	pointsOfFunctionPlot[245].Y = 1.565

	pointsOfFunctionPlot[246].X = 2.46
	pointsOfFunctionPlot[246].Y = 1.568

	pointsOfFunctionPlot[247].X = 2.47
	pointsOfFunctionPlot[247].Y = 1.571

	pointsOfFunctionPlot[248].X = 2.48
	pointsOfFunctionPlot[248].Y = 1.574

	pointsOfFunctionPlot[249].X = 2.49
	pointsOfFunctionPlot[249].Y = 1.577

	pointsOfFunctionPlot[250].X = 2.50
	pointsOfFunctionPlot[250].Y = 1.581

	pointsOfFunctionPlot[251].X = 2.51
	pointsOfFunctionPlot[251].Y = 1.584

	pointsOfFunctionPlot[252].X = 2.52
	pointsOfFunctionPlot[252].Y = 1.587

	pointsOfFunctionPlot[253].X = 2.53
	pointsOfFunctionPlot[253].Y = 1.59

	pointsOfFunctionPlot[254].X = 2.54
	pointsOfFunctionPlot[254].Y = 1.593

	pointsOfFunctionPlot[255].X = 2.55
	pointsOfFunctionPlot[255].Y = 1.596

	pointsOfFunctionPlot[256].X = 2.56
	pointsOfFunctionPlot[256].Y = 1.6

	pointsOfFunctionPlot[257].X = 2.57
	pointsOfFunctionPlot[257].Y = 1.603

	pointsOfFunctionPlot[258].X = 2.58
	pointsOfFunctionPlot[258].Y = 1.606

	pointsOfFunctionPlot[259].X = 2.59
	pointsOfFunctionPlot[259].Y = 1.609

	pointsOfFunctionPlot[260].X = 2.60
	pointsOfFunctionPlot[260].Y = 1.612

	pointsOfFunctionPlot[261].X = 2.61
	pointsOfFunctionPlot[261].Y = 1.615

	pointsOfFunctionPlot[262].X = 2.62
	pointsOfFunctionPlot[262].Y = 1.618

	pointsOfFunctionPlot[263].X = 2.63
	pointsOfFunctionPlot[263].Y = 1.621

	pointsOfFunctionPlot[264].X = 2.64
	pointsOfFunctionPlot[264].Y = 1.624

	pointsOfFunctionPlot[265].X = 2.65
	pointsOfFunctionPlot[265].Y = 1.627

	pointsOfFunctionPlot[266].X = 2.66
	pointsOfFunctionPlot[266].Y = 1.63

	pointsOfFunctionPlot[267].X = 2.67
	pointsOfFunctionPlot[267].Y = 1.634

	pointsOfFunctionPlot[268].X = 2.68
	pointsOfFunctionPlot[268].Y = 1.637

	pointsOfFunctionPlot[269].X = 2.69
	pointsOfFunctionPlot[269].Y = 1.64

	pointsOfFunctionPlot[270].X = 2.70
	pointsOfFunctionPlot[270].Y = 1.643

	pointsOfFunctionPlot[271].X = 2.71
	pointsOfFunctionPlot[271].Y = 1.646

	pointsOfFunctionPlot[272].X = 2.72
	pointsOfFunctionPlot[272].Y = 1.649

	pointsOfFunctionPlot[273].X = 2.73
	pointsOfFunctionPlot[273].Y = 1.652

	pointsOfFunctionPlot[274].X = 2.74
	pointsOfFunctionPlot[274].Y = 1.655

	pointsOfFunctionPlot[275].X = 2.75
	pointsOfFunctionPlot[275].Y = 1.658

	pointsOfFunctionPlot[276].X = 2.76
	pointsOfFunctionPlot[276].Y = 1.661

	pointsOfFunctionPlot[277].X = 2.77
	pointsOfFunctionPlot[277].Y = 1.664

	pointsOfFunctionPlot[278].X = 2.78
	pointsOfFunctionPlot[278].Y = 1.667

	pointsOfFunctionPlot[279].X = 2.79
	pointsOfFunctionPlot[279].Y = 1.67

	pointsOfFunctionPlot[280].X = 2.80
	pointsOfFunctionPlot[280].Y = 1.673

	pointsOfFunctionPlot[281].X = 2.81
	pointsOfFunctionPlot[281].Y = 1.676

	pointsOfFunctionPlot[282].X = 2.82
	pointsOfFunctionPlot[282].Y = 1.679

	pointsOfFunctionPlot[283].X = 2.83
	pointsOfFunctionPlot[283].Y = 1.682

	pointsOfFunctionPlot[284].X = 2.84
	pointsOfFunctionPlot[284].Y = 1.685

	pointsOfFunctionPlot[285].X = 2.85
	pointsOfFunctionPlot[285].Y = 1.688

	pointsOfFunctionPlot[286].X = 2.86
	pointsOfFunctionPlot[286].Y = 1.691

	pointsOfFunctionPlot[287].X = 2.87
	pointsOfFunctionPlot[287].Y = 1.694

	pointsOfFunctionPlot[288].X = 2.88
	pointsOfFunctionPlot[288].Y = 1.697

	pointsOfFunctionPlot[289].X = 2.89
	pointsOfFunctionPlot[289].Y = 1.7

	pointsOfFunctionPlot[290].X = 2.90
	pointsOfFunctionPlot[290].Y = 1.702

	pointsOfFunctionPlot[291].X = 2.91
	pointsOfFunctionPlot[291].Y = 1.705

	pointsOfFunctionPlot[292].X = 2.92
	pointsOfFunctionPlot[292].Y = 1.708

	pointsOfFunctionPlot[293].X = 2.93
	pointsOfFunctionPlot[293].Y = 1.711

	pointsOfFunctionPlot[294].X = 2.94
	pointsOfFunctionPlot[294].Y = 1.714

	pointsOfFunctionPlot[295].X = 2.95
	pointsOfFunctionPlot[295].Y = 1.717

	pointsOfFunctionPlot[296].X = 2.96
	pointsOfFunctionPlot[296].Y = 1.72

	pointsOfFunctionPlot[297].X = 2.97
	pointsOfFunctionPlot[297].Y = 1.723

	pointsOfFunctionPlot[298].X = 2.98
	pointsOfFunctionPlot[298].Y = 1.726

	pointsOfFunctionPlot[299].X = 2.99
	pointsOfFunctionPlot[299].Y = 1.729

	pointsOfFunctionPlot[300].X = 3.0
	pointsOfFunctionPlot[300].Y = 1.732

	pointsOfFunctionPlot[301].X = 3.01
	pointsOfFunctionPlot[301].Y = 1.734

	pointsOfFunctionPlot[302].X = 3.02
	pointsOfFunctionPlot[302].Y = 1.737

	pointsOfFunctionPlot[303].X = 3.03
	pointsOfFunctionPlot[303].Y = 1.74

	pointsOfFunctionPlot[304].X = 3.04
	pointsOfFunctionPlot[304].Y = 1.743

	pointsOfFunctionPlot[305].X = 3.05
	pointsOfFunctionPlot[305].Y = 1.746

	pointsOfFunctionPlot[306].X = 3.06
	pointsOfFunctionPlot[306].Y = 1.749

	pointsOfFunctionPlot[307].X = 3.07
	pointsOfFunctionPlot[307].Y = 1.752

	pointsOfFunctionPlot[308].X = 3.08
	pointsOfFunctionPlot[308].Y = 1.754

	pointsOfFunctionPlot[309].X = 3.09
	pointsOfFunctionPlot[309].Y = 1.757

	pointsOfFunctionPlot[310].X = 3.10
	pointsOfFunctionPlot[310].Y = 1.76

	pointsOfFunctionPlot[311].X = 3.11
	pointsOfFunctionPlot[311].Y = 1.763

	pointsOfFunctionPlot[312].X = 3.12
	pointsOfFunctionPlot[312].Y = 1.766

	pointsOfFunctionPlot[313].X = 3.13
	pointsOfFunctionPlot[313].Y = 1.769

	pointsOfFunctionPlot[314].X = 3.14
	pointsOfFunctionPlot[314].Y = 1.772

	pointsOfFunctionPlot[315].X = 3.15
	pointsOfFunctionPlot[315].Y = 1.774

	pointsOfFunctionPlot[316].X = 3.16
	pointsOfFunctionPlot[316].Y = 1.777

	pointsOfFunctionPlot[317].X = 3.17
	pointsOfFunctionPlot[317].Y = 1.78

	pointsOfFunctionPlot[318].X = 3.18
	pointsOfFunctionPlot[318].Y = 1.783

	pointsOfFunctionPlot[319].X = 3.19
	pointsOfFunctionPlot[319].Y = 1.786

	pointsOfFunctionPlot[320].X = 3.20
	pointsOfFunctionPlot[320].Y = 1.788

	pointsOfFunctionPlot[321].X = 3.21
	pointsOfFunctionPlot[321].Y = 1.791

	pointsOfFunctionPlot[322].X = 3.22
	pointsOfFunctionPlot[322].Y = 1.794

	pointsOfFunctionPlot[323].X = 3.23
	pointsOfFunctionPlot[323].Y = 1.797

	pointsOfFunctionPlot[324].X = 3.24
	pointsOfFunctionPlot[324].Y = 1.8

	pointsOfFunctionPlot[325].X = 3.25
	pointsOfFunctionPlot[325].Y = 1.802

	pointsOfFunctionPlot[326].X = 3.26
	pointsOfFunctionPlot[326].Y = 1.805

	pointsOfFunctionPlot[327].X = 3.27
	pointsOfFunctionPlot[327].Y = 1.808

	pointsOfFunctionPlot[328].X = 3.28
	pointsOfFunctionPlot[328].Y = 1.811

	pointsOfFunctionPlot[329].X = 3.29
	pointsOfFunctionPlot[329].Y = 1.813

	pointsOfFunctionPlot[330].X = 3.30
	pointsOfFunctionPlot[330].Y = 1.816

	pointsOfFunctionPlot[331].X = 3.31
	pointsOfFunctionPlot[331].Y = 1.819

	pointsOfFunctionPlot[332].X = 3.32
	pointsOfFunctionPlot[332].Y = 1.822

	pointsOfFunctionPlot[333].X = 3.33
	pointsOfFunctionPlot[333].Y = 1.824

	pointsOfFunctionPlot[334].X = 3.34
	pointsOfFunctionPlot[334].Y = 1.827

	pointsOfFunctionPlot[335].X = 3.35
	pointsOfFunctionPlot[335].Y = 1.83

	pointsOfFunctionPlot[336].X = 3.36
	pointsOfFunctionPlot[336].Y = 1.833

	pointsOfFunctionPlot[337].X = 3.37
	pointsOfFunctionPlot[337].Y = 1.835

	pointsOfFunctionPlot[338].X = 3.38
	pointsOfFunctionPlot[338].Y = 1.838

	pointsOfFunctionPlot[339].X = 3.39
	pointsOfFunctionPlot[339].Y = 1.841

	pointsOfFunctionPlot[340].X = 3.40
	pointsOfFunctionPlot[340].Y = 1.843

	pointsOfFunctionPlot[341].X = 3.41
	pointsOfFunctionPlot[341].Y = 1.846

	pointsOfFunctionPlot[342].X = 3.42
	pointsOfFunctionPlot[342].Y = 1.849

	pointsOfFunctionPlot[343].X = 3.43
	pointsOfFunctionPlot[343].Y = 1.852

	pointsOfFunctionPlot[344].X = 3.44
	pointsOfFunctionPlot[344].Y = 1.854

	pointsOfFunctionPlot[345].X = 3.45
	pointsOfFunctionPlot[345].Y = 1.857

	pointsOfFunctionPlot[346].X = 3.46
	pointsOfFunctionPlot[346].Y = 1.86

	pointsOfFunctionPlot[347].X = 3.47
	pointsOfFunctionPlot[347].Y = 1.862

	pointsOfFunctionPlot[348].X = 3.48
	pointsOfFunctionPlot[348].Y = 1.865

	pointsOfFunctionPlot[349].X = 3.49
	pointsOfFunctionPlot[349].Y = 1.868

	pointsOfFunctionPlot[350].X = 3.50
	pointsOfFunctionPlot[350].Y = 1.87

	pointsOfFunctionPlot[351].X = 3.51
	pointsOfFunctionPlot[351].Y = 1.873

	pointsOfFunctionPlot[352].X = 3.52
	pointsOfFunctionPlot[352].Y = 1.876

	pointsOfFunctionPlot[353].X = 3.53
	pointsOfFunctionPlot[353].Y = 1.878

	pointsOfFunctionPlot[354].X = 3.54
	pointsOfFunctionPlot[354].Y = 1.881

	pointsOfFunctionPlot[355].X = 3.55
	pointsOfFunctionPlot[355].Y = 1.884

	pointsOfFunctionPlot[356].X = 3.56
	pointsOfFunctionPlot[356].Y = 1.886

	pointsOfFunctionPlot[357].X = 3.57
	pointsOfFunctionPlot[357].Y = 1.889

	pointsOfFunctionPlot[358].X = 3.58
	pointsOfFunctionPlot[358].Y = 1.892

	pointsOfFunctionPlot[359].X = 3.59
	pointsOfFunctionPlot[359].Y = 1.894

	pointsOfFunctionPlot[360].X = 3.60
	pointsOfFunctionPlot[360].Y = 1.897

	pointsOfFunctionPlot[361].X = 3.61
	pointsOfFunctionPlot[361].Y = 1.9

	pointsOfFunctionPlot[362].X = 3.62
	pointsOfFunctionPlot[362].Y = 1.902

	pointsOfFunctionPlot[363].X = 3.63
	pointsOfFunctionPlot[363].Y = 1.905

	pointsOfFunctionPlot[364].X = 3.64
	pointsOfFunctionPlot[364].Y = 1.907

	pointsOfFunctionPlot[365].X = 3.65
	pointsOfFunctionPlot[365].Y = 1.91

	pointsOfFunctionPlot[366].X = 3.66
	pointsOfFunctionPlot[366].Y = 1.913

	pointsOfFunctionPlot[367].X = 3.67
	pointsOfFunctionPlot[367].Y = 1.915

	pointsOfFunctionPlot[368].X = 3.68
	pointsOfFunctionPlot[368].Y = 1.918

	pointsOfFunctionPlot[369].X = 3.69
	pointsOfFunctionPlot[369].Y = 1.92

	pointsOfFunctionPlot[370].X = 3.70
	pointsOfFunctionPlot[370].Y = 1.923

	pointsOfFunctionPlot[371].X = 3.71
	pointsOfFunctionPlot[371].Y = 1.926

	pointsOfFunctionPlot[372].X = 3.72
	pointsOfFunctionPlot[372].Y = 1.928

	pointsOfFunctionPlot[373].X = 3.73
	pointsOfFunctionPlot[373].Y = 1.931

	pointsOfFunctionPlot[374].X = 3.74
	pointsOfFunctionPlot[374].Y = 1.933

	pointsOfFunctionPlot[375].X = 3.75
	pointsOfFunctionPlot[375].Y = 1.936

	pointsOfFunctionPlot[376].X = 3.76
	pointsOfFunctionPlot[376].Y = 1.939

	pointsOfFunctionPlot[377].X = 3.77
	pointsOfFunctionPlot[377].Y = 1.941

	pointsOfFunctionPlot[378].X = 3.78
	pointsOfFunctionPlot[378].Y = 1.944

	pointsOfFunctionPlot[379].X = 3.79
	pointsOfFunctionPlot[379].Y = 1.946

	pointsOfFunctionPlot[380].X = 3.80
	pointsOfFunctionPlot[380].Y = 1.949

	pointsOfFunctionPlot[381].X = 3.81
	pointsOfFunctionPlot[381].Y = 1.951

	pointsOfFunctionPlot[382].X = 3.82
	pointsOfFunctionPlot[382].Y = 1.954

	pointsOfFunctionPlot[383].X = 3.83
	pointsOfFunctionPlot[383].Y = 1.957

	pointsOfFunctionPlot[384].X = 3.84
	pointsOfFunctionPlot[384].Y = 1.959

	pointsOfFunctionPlot[385].X = 3.85
	pointsOfFunctionPlot[385].Y = 1.962

	pointsOfFunctionPlot[386].X = 3.86
	pointsOfFunctionPlot[386].Y = 1.964

	pointsOfFunctionPlot[387].X = 3.87
	pointsOfFunctionPlot[387].Y = 1.967

	pointsOfFunctionPlot[388].X = 3.88
	pointsOfFunctionPlot[388].Y = 1.969

	pointsOfFunctionPlot[389].X = 3.89
	pointsOfFunctionPlot[389].Y = 1.972

	pointsOfFunctionPlot[390].X = 3.90
	pointsOfFunctionPlot[390].Y = 1.974

	pointsOfFunctionPlot[391].X = 3.91
	pointsOfFunctionPlot[391].Y = 1.977

	pointsOfFunctionPlot[392].X = 3.92
	pointsOfFunctionPlot[392].Y = 1.979

	pointsOfFunctionPlot[393].X = 3.93
	pointsOfFunctionPlot[393].Y = 1.982

	pointsOfFunctionPlot[394].X = 3.94
	pointsOfFunctionPlot[394].Y = 1.984

	pointsOfFunctionPlot[395].X = 3.95
	pointsOfFunctionPlot[395].Y = 1.987

	pointsOfFunctionPlot[396].X = 3.96
	pointsOfFunctionPlot[396].Y = 1.989

	pointsOfFunctionPlot[397].X = 3.97
	pointsOfFunctionPlot[397].Y = 1.992

	pointsOfFunctionPlot[398].X = 3.98
	pointsOfFunctionPlot[398].Y = 1.994

	pointsOfFunctionPlot[399].X = 3.99
	pointsOfFunctionPlot[399].Y = 1.997

	pointsOfFunctionPlot[400].X = 4.0
	pointsOfFunctionPlot[400].Y = 2.0

	pointsOfFunctionPlot[401].X = 4.01
	pointsOfFunctionPlot[401].Y = 2.002

	pointsOfFunctionPlot[402].X = 4.02
	pointsOfFunctionPlot[402].Y = 2.004

	pointsOfFunctionPlot[403].X = 4.03
	pointsOfFunctionPlot[403].Y = 2.007

	pointsOfFunctionPlot[404].X = 4.04
	pointsOfFunctionPlot[404].Y = 2.009

	pointsOfFunctionPlot[405].X = 4.05
	pointsOfFunctionPlot[405].Y = 2.012

	pointsOfFunctionPlot[406].X = 4.06
	pointsOfFunctionPlot[406].Y = 2.014

	pointsOfFunctionPlot[407].X = 4.07
	pointsOfFunctionPlot[407].Y = 2.017

	pointsOfFunctionPlot[408].X = 4.08
	pointsOfFunctionPlot[408].Y = 2.019

	pointsOfFunctionPlot[409].X = 4.09
	pointsOfFunctionPlot[409].Y = 2.022

	pointsOfFunctionPlot[410].X = 4.10
	pointsOfFunctionPlot[410].Y = 2.024

	pointsOfFunctionPlot[411].X = 4.11
	pointsOfFunctionPlot[411].Y = 2.027

	pointsOfFunctionPlot[412].X = 4.12
	pointsOfFunctionPlot[412].Y = 2.029

	pointsOfFunctionPlot[413].X = 4.13
	pointsOfFunctionPlot[413].Y = 2.032

	pointsOfFunctionPlot[414].X = 4.14
	pointsOfFunctionPlot[414].Y = 2.034

	pointsOfFunctionPlot[415].X = 4.15
	pointsOfFunctionPlot[415].Y = 2.037

	pointsOfFunctionPlot[416].X = 4.16
	pointsOfFunctionPlot[416].Y = 2.039

	pointsOfFunctionPlot[417].X = 4.17
	pointsOfFunctionPlot[417].Y = 2.042

	pointsOfFunctionPlot[418].X = 4.18
	pointsOfFunctionPlot[418].Y = 2.044

	pointsOfFunctionPlot[419].X = 4.19
	pointsOfFunctionPlot[419].Y = 2.046

	pointsOfFunctionPlot[420].X = 4.20
	pointsOfFunctionPlot[420].Y = 2.049

	pointsOfFunctionPlot[421].X = 4.21
	pointsOfFunctionPlot[421].Y = 2.051

	pointsOfFunctionPlot[422].X = 4.22
	pointsOfFunctionPlot[422].Y = 2.054

	pointsOfFunctionPlot[423].X = 4.23
	pointsOfFunctionPlot[423].Y = 2.056

	pointsOfFunctionPlot[424].X = 4.24
	pointsOfFunctionPlot[424].Y = 2.059

	pointsOfFunctionPlot[425].X = 4.25
	pointsOfFunctionPlot[425].Y = 2.061

	pointsOfFunctionPlot[426].X = 4.26
	pointsOfFunctionPlot[426].Y = 2.063

	pointsOfFunctionPlot[427].X = 4.27
	pointsOfFunctionPlot[427].Y = 2.066

	pointsOfFunctionPlot[428].X = 4.28
	pointsOfFunctionPlot[428].Y = 2.068

	pointsOfFunctionPlot[429].X = 4.29
	pointsOfFunctionPlot[429].Y = 2.071

	pointsOfFunctionPlot[430].X = 4.30
	pointsOfFunctionPlot[430].Y = 2.073

	pointsOfFunctionPlot[431].X = 4.31
	pointsOfFunctionPlot[431].Y = 2.076

	pointsOfFunctionPlot[432].X = 4.32
	pointsOfFunctionPlot[432].Y = 2.078

	pointsOfFunctionPlot[433].X = 4.33
	pointsOfFunctionPlot[433].Y = 2.08

	pointsOfFunctionPlot[434].X = 4.34
	pointsOfFunctionPlot[434].Y = 2.083

	pointsOfFunctionPlot[435].X = 4.35
	pointsOfFunctionPlot[435].Y = 2.085

	pointsOfFunctionPlot[436].X = 4.36
	pointsOfFunctionPlot[436].Y = 2.088

	pointsOfFunctionPlot[437].X = 4.37
	pointsOfFunctionPlot[437].Y = 2.09

	pointsOfFunctionPlot[438].X = 4.38
	pointsOfFunctionPlot[438].Y = 2.092

	pointsOfFunctionPlot[439].X = 4.39
	pointsOfFunctionPlot[439].Y = 2.095

	pointsOfFunctionPlot[440].X = 4.40
	pointsOfFunctionPlot[440].Y = 2.097

	pointsOfFunctionPlot[441].X = 4.41
	pointsOfFunctionPlot[441].Y = 2.1

	pointsOfFunctionPlot[442].X = 4.42
	pointsOfFunctionPlot[442].Y = 2.102

	pointsOfFunctionPlot[443].X = 4.43
	pointsOfFunctionPlot[443].Y = 2.104

	pointsOfFunctionPlot[444].X = 4.44
	pointsOfFunctionPlot[444].Y = 2.107

	pointsOfFunctionPlot[445].X = 4.45
	pointsOfFunctionPlot[445].Y = 2.109

	pointsOfFunctionPlot[446].X = 4.46
	pointsOfFunctionPlot[446].Y = 2.111

	pointsOfFunctionPlot[447].X = 4.47
	pointsOfFunctionPlot[447].Y = 2.114

	pointsOfFunctionPlot[448].X = 4.48
	pointsOfFunctionPlot[448].Y = 2.116

	pointsOfFunctionPlot[449].X = 4.49
	pointsOfFunctionPlot[449].Y = 2.118

	pointsOfFunctionPlot[450].X = 4.50
	pointsOfFunctionPlot[450].Y = 2.121

	pointsOfFunctionPlot[451].X = 4.51
	pointsOfFunctionPlot[451].Y = 2.123

	pointsOfFunctionPlot[452].X = 4.52
	pointsOfFunctionPlot[452].Y = 2.126

	pointsOfFunctionPlot[453].X = 4.53
	pointsOfFunctionPlot[453].Y = 2.128

	pointsOfFunctionPlot[454].X = 4.54
	pointsOfFunctionPlot[454].Y = 2.13

	pointsOfFunctionPlot[455].X = 4.55
	pointsOfFunctionPlot[455].Y = 2.133

	pointsOfFunctionPlot[456].X = 4.56
	pointsOfFunctionPlot[456].Y = 2.135

	pointsOfFunctionPlot[457].X = 4.57
	pointsOfFunctionPlot[457].Y = 2.137

	pointsOfFunctionPlot[458].X = 4.58
	pointsOfFunctionPlot[458].Y = 2.14

	pointsOfFunctionPlot[459].X = 4.59
	pointsOfFunctionPlot[459].Y = 2.142

	pointsOfFunctionPlot[460].X = 4.60
	pointsOfFunctionPlot[460].Y = 2.144

	pointsOfFunctionPlot[461].X = 4.61
	pointsOfFunctionPlot[461].Y = 2.147

	pointsOfFunctionPlot[462].X = 4.62
	pointsOfFunctionPlot[462].Y = 2.149

	pointsOfFunctionPlot[463].X = 4.63
	pointsOfFunctionPlot[463].Y = 2.151

	pointsOfFunctionPlot[464].X = 4.64
	pointsOfFunctionPlot[464].Y = 2.154

	pointsOfFunctionPlot[465].X = 4.65
	pointsOfFunctionPlot[465].Y = 2.156

	pointsOfFunctionPlot[466].X = 4.66
	pointsOfFunctionPlot[466].Y = 2.158

	pointsOfFunctionPlot[467].X = 4.67
	pointsOfFunctionPlot[467].Y = 2.161

	pointsOfFunctionPlot[468].X = 4.68
	pointsOfFunctionPlot[468].Y = 2.163

	pointsOfFunctionPlot[469].X = 4.69
	pointsOfFunctionPlot[469].Y = 2.165

	pointsOfFunctionPlot[470].X = 4.70
	pointsOfFunctionPlot[470].Y = 2.167

	pointsOfFunctionPlot[471].X = 4.71
	pointsOfFunctionPlot[471].Y = 2.17

	pointsOfFunctionPlot[472].X = 4.72
	pointsOfFunctionPlot[472].Y = 2.172

	pointsOfFunctionPlot[473].X = 4.73
	pointsOfFunctionPlot[473].Y = 2.174

	pointsOfFunctionPlot[474].X = 4.74
	pointsOfFunctionPlot[474].Y = 2.177

	pointsOfFunctionPlot[475].X = 4.75
	pointsOfFunctionPlot[475].Y = 2.179

	pointsOfFunctionPlot[476].X = 4.76
	pointsOfFunctionPlot[476].Y = 2.181

	pointsOfFunctionPlot[477].X = 4.77
	pointsOfFunctionPlot[477].Y = 2.184

	pointsOfFunctionPlot[478].X = 4.78
	pointsOfFunctionPlot[478].Y = 2.186

	pointsOfFunctionPlot[479].X = 4.79
	pointsOfFunctionPlot[479].Y = 2.188

	pointsOfFunctionPlot[480].X = 4.80
	pointsOfFunctionPlot[480].Y = 2.19

	pointsOfFunctionPlot[481].X = 4.81
	pointsOfFunctionPlot[481].Y = 2.193

	pointsOfFunctionPlot[482].X = 4.82
	pointsOfFunctionPlot[482].Y = 2.195

	pointsOfFunctionPlot[483].X = 4.83
	pointsOfFunctionPlot[483].Y = 2.197

	pointsOfFunctionPlot[484].X = 4.84
	pointsOfFunctionPlot[484].Y = 2.2

	pointsOfFunctionPlot[485].X = 4.85
	pointsOfFunctionPlot[485].Y = 2.202

	pointsOfFunctionPlot[486].X = 4.86
	pointsOfFunctionPlot[486].Y = 2.204

	pointsOfFunctionPlot[487].X = 4.87
	pointsOfFunctionPlot[487].Y = 2.206

	pointsOfFunctionPlot[488].X = 4.88
	pointsOfFunctionPlot[488].Y = 2.209

	pointsOfFunctionPlot[489].X = 4.89
	pointsOfFunctionPlot[489].Y = 2.211

	pointsOfFunctionPlot[490].X = 4.90
	pointsOfFunctionPlot[490].Y = 2.213

	pointsOfFunctionPlot[491].X = 4.91
	pointsOfFunctionPlot[491].Y = 2.215

	pointsOfFunctionPlot[492].X = 4.92
	pointsOfFunctionPlot[492].Y = 2.218

	pointsOfFunctionPlot[493].X = 4.93
	pointsOfFunctionPlot[493].Y = 2.22

	pointsOfFunctionPlot[494].X = 4.94
	pointsOfFunctionPlot[494].Y = 2.222

	pointsOfFunctionPlot[495].X = 4.95
	pointsOfFunctionPlot[495].Y = 2.224

	pointsOfFunctionPlot[496].X = 4.96
	pointsOfFunctionPlot[496].Y = 2.227

	pointsOfFunctionPlot[497].X = 4.97
	pointsOfFunctionPlot[497].Y = 2.229

	pointsOfFunctionPlot[498].X = 4.98
	pointsOfFunctionPlot[498].Y = 2.231

	pointsOfFunctionPlot[499].X = 4.99
	pointsOfFunctionPlot[499].Y = 2.233

	pointsOfFunctionPlot[500].X = 5.0
	pointsOfFunctionPlot[500].Y = 2.236

	pointsOfFunctionPlot[501].X = 5.01
	pointsOfFunctionPlot[501].Y = 2.238

	pointsOfFunctionPlot[502].X = 5.02
	pointsOfFunctionPlot[502].Y = 2.24

	pointsOfFunctionPlot[503].X = 5.03
	pointsOfFunctionPlot[503].Y = 2.242

	pointsOfFunctionPlot[504].X = 5.04
	pointsOfFunctionPlot[504].Y = 2.244

	pointsOfFunctionPlot[505].X = 5.05
	pointsOfFunctionPlot[505].Y = 2.247

	pointsOfFunctionPlot[506].X = 5.06
	pointsOfFunctionPlot[506].Y = 2.249

	pointsOfFunctionPlot[507].X = 5.07
	pointsOfFunctionPlot[507].Y = 2.251

	pointsOfFunctionPlot[508].X = 5.08
	pointsOfFunctionPlot[508].Y = 2.253

	pointsOfFunctionPlot[509].X = 5.09
	pointsOfFunctionPlot[509].Y = 2.256

	pointsOfFunctionPlot[510].X = 5.10
	pointsOfFunctionPlot[510].Y = 2.258

	pointsOfFunctionPlot[511].X = 5.11
	pointsOfFunctionPlot[511].Y = 2.26

	pointsOfFunctionPlot[512].X = 5.12
	pointsOfFunctionPlot[512].Y = 2.262

	pointsOfFunctionPlot[513].X = 5.13
	pointsOfFunctionPlot[513].Y = 2.264

	pointsOfFunctionPlot[514].X = 5.14
	pointsOfFunctionPlot[514].Y = 2.267

	pointsOfFunctionPlot[515].X = 5.15
	pointsOfFunctionPlot[515].Y = 2.269

	pointsOfFunctionPlot[516].X = 5.16
	pointsOfFunctionPlot[516].Y = 2.271

	pointsOfFunctionPlot[517].X = 5.17
	pointsOfFunctionPlot[517].Y = 2.273

	pointsOfFunctionPlot[518].X = 5.18
	pointsOfFunctionPlot[518].Y = 2.275

	pointsOfFunctionPlot[519].X = 5.19
	pointsOfFunctionPlot[519].Y = 2.278

	pointsOfFunctionPlot[520].X = 5.20
	pointsOfFunctionPlot[520].Y = 2.28

	pointsOfFunctionPlot[521].X = 5.21
	pointsOfFunctionPlot[521].Y = 2.282

	pointsOfFunctionPlot[522].X = 5.22
	pointsOfFunctionPlot[522].Y = 2.284

	pointsOfFunctionPlot[523].X = 5.23
	pointsOfFunctionPlot[523].Y = 2.286

	pointsOfFunctionPlot[524].X = 5.24
	pointsOfFunctionPlot[524].Y = 2.289

	pointsOfFunctionPlot[525].X = 5.25
	pointsOfFunctionPlot[525].Y = 2.291

	pointsOfFunctionPlot[526].X = 5.26
	pointsOfFunctionPlot[526].Y = 2.293

	pointsOfFunctionPlot[527].X = 5.27
	pointsOfFunctionPlot[527].Y = 2.295

	pointsOfFunctionPlot[528].X = 5.28
	pointsOfFunctionPlot[528].Y = 2.297

	pointsOfFunctionPlot[529].X = 5.29
	pointsOfFunctionPlot[529].Y = 2.3

	pointsOfFunctionPlot[530].X = 5.30
	pointsOfFunctionPlot[530].Y = 2.302

	pointsOfFunctionPlot[531].X = 5.31
	pointsOfFunctionPlot[531].Y = 2.304

	pointsOfFunctionPlot[532].X = 5.32
	pointsOfFunctionPlot[532].Y = 2.306

	pointsOfFunctionPlot[533].X = 5.33
	pointsOfFunctionPlot[533].Y = 2.308

	pointsOfFunctionPlot[534].X = 5.34
	pointsOfFunctionPlot[534].Y = 2.31

	pointsOfFunctionPlot[535].X = 5.35
	pointsOfFunctionPlot[535].Y = 2.313

	pointsOfFunctionPlot[536].X = 5.36
	pointsOfFunctionPlot[536].Y = 2.315

	pointsOfFunctionPlot[537].X = 5.37
	pointsOfFunctionPlot[537].Y = 2.317

	pointsOfFunctionPlot[538].X = 5.38
	pointsOfFunctionPlot[538].Y = 2.319

	pointsOfFunctionPlot[539].X = 5.39
	pointsOfFunctionPlot[539].Y = 2.321

	pointsOfFunctionPlot[540].X = 5.40
	pointsOfFunctionPlot[540].Y = 2.323

	pointsOfFunctionPlot[541].X = 5.41
	pointsOfFunctionPlot[541].Y = 2.325

	pointsOfFunctionPlot[542].X = 5.42
	pointsOfFunctionPlot[542].Y = 2.328

	pointsOfFunctionPlot[543].X = 5.43
	pointsOfFunctionPlot[543].Y = 2.33

	pointsOfFunctionPlot[544].X = 5.44
	pointsOfFunctionPlot[544].Y = 2.332

	pointsOfFunctionPlot[545].X = 5.45
	pointsOfFunctionPlot[545].Y = 2.334

	pointsOfFunctionPlot[546].X = 5.46
	pointsOfFunctionPlot[546].Y = 2.336

	pointsOfFunctionPlot[547].X = 5.47
	pointsOfFunctionPlot[547].Y = 2.338

	pointsOfFunctionPlot[548].X = 5.48
	pointsOfFunctionPlot[548].Y = 2.34

	pointsOfFunctionPlot[549].X = 5.49
	pointsOfFunctionPlot[549].Y = 2.343

	pointsOfFunctionPlot[550].X = 5.50
	pointsOfFunctionPlot[550].Y = 2.345

	pointsOfFunctionPlot[551].X = 5.51
	pointsOfFunctionPlot[551].Y = 2.347

	pointsOfFunctionPlot[552].X = 5.52
	pointsOfFunctionPlot[552].Y = 2.349

	pointsOfFunctionPlot[553].X = 5.53
	pointsOfFunctionPlot[553].Y = 2.351

	pointsOfFunctionPlot[554].X = 5.54
	pointsOfFunctionPlot[554].Y = 2.353

	pointsOfFunctionPlot[555].X = 5.55
	pointsOfFunctionPlot[555].Y = 2.355

	pointsOfFunctionPlot[556].X = 5.56
	pointsOfFunctionPlot[556].Y = 2.357

	pointsOfFunctionPlot[557].X = 5.57
	pointsOfFunctionPlot[557].Y = 2.36

	pointsOfFunctionPlot[558].X = 5.58
	pointsOfFunctionPlot[558].Y = 2.362

	pointsOfFunctionPlot[559].X = 5.59
	pointsOfFunctionPlot[559].Y = 2.364

	pointsOfFunctionPlot[560].X = 5.60
	pointsOfFunctionPlot[560].Y = 2.366

	pointsOfFunctionPlot[561].X = 5.61
	pointsOfFunctionPlot[561].Y = 2.368

	pointsOfFunctionPlot[562].X = 5.62
	pointsOfFunctionPlot[562].Y = 2.37

	pointsOfFunctionPlot[563].X = 5.63
	pointsOfFunctionPlot[563].Y = 2.372

	pointsOfFunctionPlot[564].X = 5.64
	pointsOfFunctionPlot[564].Y = 2.374

	pointsOfFunctionPlot[565].X = 5.65
	pointsOfFunctionPlot[565].Y = 2.376

	pointsOfFunctionPlot[566].X = 5.66
	pointsOfFunctionPlot[566].Y = 2.379

	pointsOfFunctionPlot[567].X = 5.67
	pointsOfFunctionPlot[567].Y = 2.381

	pointsOfFunctionPlot[568].X = 5.68
	pointsOfFunctionPlot[568].Y = 2.383

	pointsOfFunctionPlot[569].X = 5.69
	pointsOfFunctionPlot[569].Y = 2.385

	pointsOfFunctionPlot[570].X = 5.70
	pointsOfFunctionPlot[570].Y = 2.387

	pointsOfFunctionPlot[571].X = 5.71
	pointsOfFunctionPlot[571].Y = 2.389

	pointsOfFunctionPlot[572].X = 5.72
	pointsOfFunctionPlot[572].Y = 2.391

	pointsOfFunctionPlot[573].X = 5.73
	pointsOfFunctionPlot[573].Y = 2.393

	pointsOfFunctionPlot[574].X = 5.74
	pointsOfFunctionPlot[574].Y = 2.395

	pointsOfFunctionPlot[575].X = 5.75
	pointsOfFunctionPlot[575].Y = 2.397

	pointsOfFunctionPlot[576].X = 5.76
	pointsOfFunctionPlot[576].Y = 2.4

	pointsOfFunctionPlot[577].X = 5.77
	pointsOfFunctionPlot[577].Y = 2.402

	pointsOfFunctionPlot[578].X = 5.78
	pointsOfFunctionPlot[578].Y = 2.404

	pointsOfFunctionPlot[579].X = 5.79
	pointsOfFunctionPlot[579].Y = 2.406

	pointsOfFunctionPlot[580].X = 5.80
	pointsOfFunctionPlot[580].Y = 2.408

	pointsOfFunctionPlot[581].X = 5.81
	pointsOfFunctionPlot[581].Y = 2.41

	pointsOfFunctionPlot[582].X = 5.82
	pointsOfFunctionPlot[582].Y = 2.412

	pointsOfFunctionPlot[583].X = 5.83
	pointsOfFunctionPlot[583].Y = 2.414

	pointsOfFunctionPlot[584].X = 5.84
	pointsOfFunctionPlot[584].Y = 2.416

	pointsOfFunctionPlot[585].X = 5.85
	pointsOfFunctionPlot[585].Y = 2.418

	pointsOfFunctionPlot[586].X = 5.86
	pointsOfFunctionPlot[586].Y = 2.42

	pointsOfFunctionPlot[587].X = 5.87
	pointsOfFunctionPlot[587].Y = 2.422

	pointsOfFunctionPlot[588].X = 5.88
	pointsOfFunctionPlot[588].Y = 2.424

	pointsOfFunctionPlot[589].X = 5.89
	pointsOfFunctionPlot[589].Y = 2.426

	pointsOfFunctionPlot[590].X = 5.90
	pointsOfFunctionPlot[590].Y = 2.428

	pointsOfFunctionPlot[591].X = 5.91
	pointsOfFunctionPlot[591].Y = 2.431

	pointsOfFunctionPlot[592].X = 5.92
	pointsOfFunctionPlot[592].Y = 2.433

	pointsOfFunctionPlot[593].X = 5.93
	pointsOfFunctionPlot[593].Y = 2.435

	pointsOfFunctionPlot[594].X = 5.94
	pointsOfFunctionPlot[594].Y = 2.437

	pointsOfFunctionPlot[595].X = 5.95
	pointsOfFunctionPlot[595].Y = 2.439

	pointsOfFunctionPlot[596].X = 5.96
	pointsOfFunctionPlot[596].Y = 2.441

	pointsOfFunctionPlot[597].X = 5.97
	pointsOfFunctionPlot[597].Y = 2.443

	pointsOfFunctionPlot[598].X = 5.98
	pointsOfFunctionPlot[598].Y = 2.445

	pointsOfFunctionPlot[599].X = 5.99
	pointsOfFunctionPlot[599].Y = 2.447

	pointsOfFunctionPlot[600].X = 6.0
	pointsOfFunctionPlot[600].Y = 2.449

	pointsOfFunctionPlot[601].X = 6.01
	pointsOfFunctionPlot[601].Y = 2.451

	pointsOfFunctionPlot[602].X = 6.02
	pointsOfFunctionPlot[602].Y = 2.453

	pointsOfFunctionPlot[603].X = 6.03
	pointsOfFunctionPlot[603].Y = 2.455

	pointsOfFunctionPlot[604].X = 6.04
	pointsOfFunctionPlot[604].Y = 2.457

	pointsOfFunctionPlot[605].X = 6.05
	pointsOfFunctionPlot[605].Y = 2.459

	pointsOfFunctionPlot[606].X = 6.06
	pointsOfFunctionPlot[606].Y = 2.461

	pointsOfFunctionPlot[607].X = 6.07
	pointsOfFunctionPlot[607].Y = 2.463

	pointsOfFunctionPlot[608].X = 6.08
	pointsOfFunctionPlot[608].Y = 2.465

	pointsOfFunctionPlot[609].X = 6.09
	pointsOfFunctionPlot[609].Y = 2.467

	pointsOfFunctionPlot[610].X = 6.10
	pointsOfFunctionPlot[610].Y = 2.469

	pointsOfFunctionPlot[611].X = 6.11
	pointsOfFunctionPlot[611].Y = 2.471

	pointsOfFunctionPlot[612].X = 6.12
	pointsOfFunctionPlot[612].Y = 2.473

	pointsOfFunctionPlot[613].X = 6.13
	pointsOfFunctionPlot[613].Y = 2.475

	pointsOfFunctionPlot[614].X = 6.14
	pointsOfFunctionPlot[614].Y = 2.477

	pointsOfFunctionPlot[615].X = 6.15
	pointsOfFunctionPlot[615].Y = 2.479

	pointsOfFunctionPlot[616].X = 6.16
	pointsOfFunctionPlot[616].Y = 2.481

	pointsOfFunctionPlot[617].X = 6.17
	pointsOfFunctionPlot[617].Y = 2.483

	pointsOfFunctionPlot[618].X = 6.18
	pointsOfFunctionPlot[618].Y = 2.485

	pointsOfFunctionPlot[619].X = 6.19
	pointsOfFunctionPlot[619].Y = 2.487

	pointsOfFunctionPlot[620].X = 6.20
	pointsOfFunctionPlot[620].Y = 2.489

	pointsOfFunctionPlot[621].X = 6.21
	pointsOfFunctionPlot[621].Y = 2.491

	pointsOfFunctionPlot[622].X = 6.22
	pointsOfFunctionPlot[622].Y = 2.493

	pointsOfFunctionPlot[623].X = 6.23
	pointsOfFunctionPlot[623].Y = 2.495

	pointsOfFunctionPlot[624].X = 6.24
	pointsOfFunctionPlot[624].Y = 2.497

	pointsOfFunctionPlot[625].X = 6.25
	pointsOfFunctionPlot[625].Y = 2.5

	pointsOfFunctionPlot[626].X = 6.26
	pointsOfFunctionPlot[626].Y = 2.501

	pointsOfFunctionPlot[627].X = 6.27
	pointsOfFunctionPlot[627].Y = 2.503

	pointsOfFunctionPlot[628].X = 6.28
	pointsOfFunctionPlot[628].Y = 2.505

	pointsOfFunctionPlot[629].X = 6.29
	pointsOfFunctionPlot[629].Y = 2.507

	pointsOfFunctionPlot[630].X = 6.30
	pointsOfFunctionPlot[630].Y = 2.509

	pointsOfFunctionPlot[631].X = 6.31
	pointsOfFunctionPlot[631].Y = 2.511

	pointsOfFunctionPlot[632].X = 6.32
	pointsOfFunctionPlot[632].Y = 2.513

	pointsOfFunctionPlot[633].X = 6.33
	pointsOfFunctionPlot[633].Y = 2.515

	pointsOfFunctionPlot[634].X = 6.34
	pointsOfFunctionPlot[634].Y = 2.517

	pointsOfFunctionPlot[635].X = 6.35
	pointsOfFunctionPlot[635].Y = 2.519

	pointsOfFunctionPlot[636].X = 6.36
	pointsOfFunctionPlot[636].Y = 2.521

	pointsOfFunctionPlot[637].X = 6.37
	pointsOfFunctionPlot[637].Y = 2.523

	pointsOfFunctionPlot[638].X = 6.38
	pointsOfFunctionPlot[638].Y = 2.525

	pointsOfFunctionPlot[639].X = 6.39
	pointsOfFunctionPlot[639].Y = 2.527

	pointsOfFunctionPlot[640].X = 6.40
	pointsOfFunctionPlot[640].Y = 2.529

	pointsOfFunctionPlot[641].X = 6.41
	pointsOfFunctionPlot[641].Y = 2.531

	pointsOfFunctionPlot[642].X = 6.42
	pointsOfFunctionPlot[642].Y = 2.533

	pointsOfFunctionPlot[643].X = 6.43
	pointsOfFunctionPlot[643].Y = 2.535

	pointsOfFunctionPlot[644].X = 6.44
	pointsOfFunctionPlot[644].Y = 2.537

	pointsOfFunctionPlot[645].X = 6.45
	pointsOfFunctionPlot[645].Y = 2.539

	pointsOfFunctionPlot[646].X = 6.46
	pointsOfFunctionPlot[646].Y = 2.541

	pointsOfFunctionPlot[647].X = 6.47
	pointsOfFunctionPlot[647].Y = 2.543

	pointsOfFunctionPlot[648].X = 6.48
	pointsOfFunctionPlot[648].Y = 2.545

	pointsOfFunctionPlot[649].X = 6.49
	pointsOfFunctionPlot[649].Y = 2.547

	pointsOfFunctionPlot[650].X = 6.50
	pointsOfFunctionPlot[650].Y = 2.549

	pointsOfFunctionPlot[651].X = 6.51
	pointsOfFunctionPlot[651].Y = 2.551

	pointsOfFunctionPlot[652].X = 6.52
	pointsOfFunctionPlot[652].Y = 2.553

	pointsOfFunctionPlot[653].X = 6.53
	pointsOfFunctionPlot[653].Y = 2.555

	pointsOfFunctionPlot[654].X = 6.54
	pointsOfFunctionPlot[654].Y = 2.557

	pointsOfFunctionPlot[655].X = 6.55
	pointsOfFunctionPlot[655].Y = 2.559

	pointsOfFunctionPlot[656].X = 6.56
	pointsOfFunctionPlot[656].Y = 2.561

	pointsOfFunctionPlot[657].X = 6.57
	pointsOfFunctionPlot[657].Y = 2.563

	pointsOfFunctionPlot[658].X = 6.58
	pointsOfFunctionPlot[658].Y = 2.565

	pointsOfFunctionPlot[659].X = 6.59
	pointsOfFunctionPlot[659].Y = 2.567

	pointsOfFunctionPlot[660].X = 6.60
	pointsOfFunctionPlot[660].Y = 2.569

	pointsOfFunctionPlot[661].X = 6.61
	pointsOfFunctionPlot[661].Y = 2.57

	pointsOfFunctionPlot[662].X = 6.62
	pointsOfFunctionPlot[662].Y = 2.572

	pointsOfFunctionPlot[663].X = 6.63
	pointsOfFunctionPlot[663].Y = 2.574

	pointsOfFunctionPlot[664].X = 6.64
	pointsOfFunctionPlot[664].Y = 2.576

	pointsOfFunctionPlot[665].X = 6.65
	pointsOfFunctionPlot[665].Y = 2.578

	pointsOfFunctionPlot[666].X = 6.66
	pointsOfFunctionPlot[666].Y = 2.58

	pointsOfFunctionPlot[667].X = 6.67
	pointsOfFunctionPlot[667].Y = 2.582

	pointsOfFunctionPlot[668].X = 6.68
	pointsOfFunctionPlot[668].Y = 2.584

	pointsOfFunctionPlot[669].X = 6.69
	pointsOfFunctionPlot[669].Y = 2.586

	pointsOfFunctionPlot[670].X = 6.70
	pointsOfFunctionPlot[670].Y = 2.588

	pointsOfFunctionPlot[671].X = 6.71
	pointsOfFunctionPlot[671].Y = 2.59

	pointsOfFunctionPlot[672].X = 6.72
	pointsOfFunctionPlot[672].Y = 2.592

	pointsOfFunctionPlot[673].X = 6.73
	pointsOfFunctionPlot[673].Y = 2.594

	pointsOfFunctionPlot[674].X = 6.74
	pointsOfFunctionPlot[674].Y = 2.596

	pointsOfFunctionPlot[675].X = 6.75
	pointsOfFunctionPlot[675].Y = 2.598

	pointsOfFunctionPlot[676].X = 6.76
	pointsOfFunctionPlot[676].Y = 2.6

	pointsOfFunctionPlot[677].X = 6.77
	pointsOfFunctionPlot[677].Y = 2.601

	pointsOfFunctionPlot[678].X = 6.78
	pointsOfFunctionPlot[678].Y = 2.603

	pointsOfFunctionPlot[679].X = 6.79
	pointsOfFunctionPlot[679].Y = 2.605

	pointsOfFunctionPlot[680].X = 6.80
	pointsOfFunctionPlot[680].Y = 2.607

	pointsOfFunctionPlot[681].X = 6.81
	pointsOfFunctionPlot[681].Y = 2.609

	pointsOfFunctionPlot[682].X = 6.82
	pointsOfFunctionPlot[682].Y = 2.611

	pointsOfFunctionPlot[683].X = 6.83
	pointsOfFunctionPlot[683].Y = 2.613

	pointsOfFunctionPlot[684].X = 6.84
	pointsOfFunctionPlot[684].Y = 2.615

	pointsOfFunctionPlot[685].X = 6.85
	pointsOfFunctionPlot[685].Y = 2.617

	pointsOfFunctionPlot[686].X = 6.86
	pointsOfFunctionPlot[686].Y = 2.619

	pointsOfFunctionPlot[687].X = 6.87
	pointsOfFunctionPlot[687].Y = 2.621

	pointsOfFunctionPlot[688].X = 6.88
	pointsOfFunctionPlot[688].Y = 2.622

	pointsOfFunctionPlot[689].X = 6.89
	pointsOfFunctionPlot[689].Y = 2.624

	pointsOfFunctionPlot[690].X = 6.90
	pointsOfFunctionPlot[690].Y = 2.626

	pointsOfFunctionPlot[691].X = 6.91
	pointsOfFunctionPlot[691].Y = 2.628

	pointsOfFunctionPlot[692].X = 6.92
	pointsOfFunctionPlot[692].Y = 2.63

	pointsOfFunctionPlot[693].X = 6.93
	pointsOfFunctionPlot[693].Y = 2.632

	pointsOfFunctionPlot[694].X = 6.94
	pointsOfFunctionPlot[694].Y = 2.634

	pointsOfFunctionPlot[695].X = 6.95
	pointsOfFunctionPlot[695].Y = 2.636

	pointsOfFunctionPlot[696].X = 6.96
	pointsOfFunctionPlot[696].Y = 2.638

	pointsOfFunctionPlot[697].X = 6.97
	pointsOfFunctionPlot[697].Y = 2.64

	pointsOfFunctionPlot[698].X = 6.98
	pointsOfFunctionPlot[698].Y = 2.641

	pointsOfFunctionPlot[699].X = 6.99
	pointsOfFunctionPlot[699].Y = 2.643

	pointsOfFunctionPlot[700].X = 7.0
	pointsOfFunctionPlot[700].Y = 2.645

	pointsOfFunctionPlot[701].X = 7.01
	pointsOfFunctionPlot[701].Y = 2.647

	pointsOfFunctionPlot[702].X = 7.02
	pointsOfFunctionPlot[702].Y = 2.649

	pointsOfFunctionPlot[703].X = 7.03
	pointsOfFunctionPlot[703].Y = 2.651

	pointsOfFunctionPlot[704].X = 7.04
	pointsOfFunctionPlot[704].Y = 2.653

	pointsOfFunctionPlot[705].X = 7.05
	pointsOfFunctionPlot[705].Y = 2.655

	pointsOfFunctionPlot[706].X = 7.06
	pointsOfFunctionPlot[706].Y = 2.657

	pointsOfFunctionPlot[707].X = 7.07
	pointsOfFunctionPlot[707].Y = 2.658

	pointsOfFunctionPlot[708].X = 7.08
	pointsOfFunctionPlot[708].Y = 2.66

	pointsOfFunctionPlot[709].X = 7.09
	pointsOfFunctionPlot[709].Y = 2.662

	pointsOfFunctionPlot[710].X = 7.10
	pointsOfFunctionPlot[710].Y = 2.664

	pointsOfFunctionPlot[711].X = 7.11
	pointsOfFunctionPlot[711].Y = 2.666

	pointsOfFunctionPlot[712].X = 7.12
	pointsOfFunctionPlot[712].Y = 2.668

	pointsOfFunctionPlot[713].X = 7.13
	pointsOfFunctionPlot[713].Y = 2.67

	pointsOfFunctionPlot[714].X = 7.14
	pointsOfFunctionPlot[714].Y = 2.672

	pointsOfFunctionPlot[715].X = 7.15
	pointsOfFunctionPlot[715].Y = 2.673

	pointsOfFunctionPlot[716].X = 7.16
	pointsOfFunctionPlot[716].Y = 2.675

	pointsOfFunctionPlot[717].X = 7.17
	pointsOfFunctionPlot[717].Y = 2.677

	pointsOfFunctionPlot[718].X = 7.18
	pointsOfFunctionPlot[718].Y = 2.679

	pointsOfFunctionPlot[719].X = 7.19
	pointsOfFunctionPlot[719].Y = 2.681

	pointsOfFunctionPlot[720].X = 7.20
	pointsOfFunctionPlot[720].Y = 2.683

	pointsOfFunctionPlot[721].X = 7.21
	pointsOfFunctionPlot[721].Y = 2.685

	pointsOfFunctionPlot[722].X = 7.22
	pointsOfFunctionPlot[722].Y = 2.687

	pointsOfFunctionPlot[723].X = 7.23
	pointsOfFunctionPlot[723].Y = 2.688

	pointsOfFunctionPlot[724].X = 7.24
	pointsOfFunctionPlot[724].Y = 2.69

	pointsOfFunctionPlot[725].X = 7.25
	pointsOfFunctionPlot[725].Y = 2.692

	pointsOfFunctionPlot[726].X = 7.26
	pointsOfFunctionPlot[726].Y = 2.694

	pointsOfFunctionPlot[727].X = 7.27
	pointsOfFunctionPlot[727].Y = 2.696

	pointsOfFunctionPlot[728].X = 7.28
	pointsOfFunctionPlot[728].Y = 2.698

	pointsOfFunctionPlot[729].X = 7.29
	pointsOfFunctionPlot[729].Y = 2.7

	pointsOfFunctionPlot[730].X = 7.30
	pointsOfFunctionPlot[730].Y = 2.701

	pointsOfFunctionPlot[731].X = 7.31
	pointsOfFunctionPlot[731].Y = 2.703

	pointsOfFunctionPlot[732].X = 7.32
	pointsOfFunctionPlot[732].Y = 2.705

	pointsOfFunctionPlot[733].X = 7.33
	pointsOfFunctionPlot[733].Y = 2.707

	pointsOfFunctionPlot[734].X = 7.34
	pointsOfFunctionPlot[734].Y = 2.709

	pointsOfFunctionPlot[735].X = 7.35
	pointsOfFunctionPlot[735].Y = 2.711

	pointsOfFunctionPlot[736].X = 7.36
	pointsOfFunctionPlot[736].Y = 2.712

	pointsOfFunctionPlot[737].X = 7.37
	pointsOfFunctionPlot[737].Y = 2.714

	pointsOfFunctionPlot[738].X = 7.38
	pointsOfFunctionPlot[738].Y = 2.716

	pointsOfFunctionPlot[739].X = 7.39
	pointsOfFunctionPlot[739].Y = 2.718

	pointsOfFunctionPlot[740].X = 7.40
	pointsOfFunctionPlot[740].Y = 2.72

	pointsOfFunctionPlot[741].X = 7.41
	pointsOfFunctionPlot[741].Y = 2.722

	pointsOfFunctionPlot[742].X = 7.42
	pointsOfFunctionPlot[742].Y = 2.723

	pointsOfFunctionPlot[743].X = 7.43
	pointsOfFunctionPlot[743].Y = 2.725

	pointsOfFunctionPlot[744].X = 7.44
	pointsOfFunctionPlot[744].Y = 2.727

	pointsOfFunctionPlot[745].X = 7.45
	pointsOfFunctionPlot[745].Y = 2.729

	pointsOfFunctionPlot[746].X = 7.46
	pointsOfFunctionPlot[746].Y = 2.731

	pointsOfFunctionPlot[747].X = 7.47
	pointsOfFunctionPlot[747].Y = 2.733

	pointsOfFunctionPlot[748].X = 7.48
	pointsOfFunctionPlot[748].Y = 2.734

	pointsOfFunctionPlot[749].X = 7.49
	pointsOfFunctionPlot[749].Y = 2.736

	pointsOfFunctionPlot[750].X = 7.50
	pointsOfFunctionPlot[750].Y = 2.738

	pointsOfFunctionPlot[751].X = 7.51
	pointsOfFunctionPlot[751].Y = 2.74

	pointsOfFunctionPlot[752].X = 7.52
	pointsOfFunctionPlot[752].Y = 2.742

	pointsOfFunctionPlot[753].X = 7.53
	pointsOfFunctionPlot[753].Y = 2.744

	pointsOfFunctionPlot[754].X = 7.54
	pointsOfFunctionPlot[754].Y = 2.745

	pointsOfFunctionPlot[755].X = 7.55
	pointsOfFunctionPlot[755].Y = 2.747

	pointsOfFunctionPlot[756].X = 7.56
	pointsOfFunctionPlot[756].Y = 2.749

	pointsOfFunctionPlot[757].X = 7.57
	pointsOfFunctionPlot[757].Y = 2.751

	pointsOfFunctionPlot[758].X = 7.58
	pointsOfFunctionPlot[758].Y = 2.753

	pointsOfFunctionPlot[759].X = 7.59
	pointsOfFunctionPlot[759].Y = 2.754

	pointsOfFunctionPlot[760].X = 7.60
	pointsOfFunctionPlot[760].Y = 2.756

	pointsOfFunctionPlot[761].X = 7.61
	pointsOfFunctionPlot[761].Y = 2.758

	pointsOfFunctionPlot[762].X = 7.62
	pointsOfFunctionPlot[762].Y = 2.76

	pointsOfFunctionPlot[763].X = 7.63
	pointsOfFunctionPlot[763].Y = 2.762

	pointsOfFunctionPlot[764].X = 7.64
	pointsOfFunctionPlot[764].Y = 2.764

	pointsOfFunctionPlot[765].X = 7.65
	pointsOfFunctionPlot[765].Y = 2.765

	pointsOfFunctionPlot[766].X = 7.66
	pointsOfFunctionPlot[766].Y = 2.767

	pointsOfFunctionPlot[767].X = 7.67
	pointsOfFunctionPlot[767].Y = 2.769

	pointsOfFunctionPlot[768].X = 7.68
	pointsOfFunctionPlot[768].Y = 2.771

	pointsOfFunctionPlot[769].X = 7.69
	pointsOfFunctionPlot[769].Y = 2.773

	pointsOfFunctionPlot[770].X = 7.70
	pointsOfFunctionPlot[770].Y = 2.774

	pointsOfFunctionPlot[771].X = 7.71
	pointsOfFunctionPlot[771].Y = 2.776

	pointsOfFunctionPlot[772].X = 7.72
	pointsOfFunctionPlot[772].Y = 2.778

	pointsOfFunctionPlot[773].X = 7.73
	pointsOfFunctionPlot[773].Y = 2.78

	pointsOfFunctionPlot[774].X = 7.74
	pointsOfFunctionPlot[774].Y = 2.782

	pointsOfFunctionPlot[775].X = 7.75
	pointsOfFunctionPlot[775].Y = 2.783

	pointsOfFunctionPlot[776].X = 7.76
	pointsOfFunctionPlot[776].Y = 2.785

	pointsOfFunctionPlot[777].X = 7.77
	pointsOfFunctionPlot[777].Y = 2.787

	pointsOfFunctionPlot[778].X = 7.78
	pointsOfFunctionPlot[778].Y = 2.789

	pointsOfFunctionPlot[779].X = 7.79
	pointsOfFunctionPlot[779].Y = 2.791

	pointsOfFunctionPlot[780].X = 7.80
	pointsOfFunctionPlot[780].Y = 2.792

	pointsOfFunctionPlot[781].X = 7.81
	pointsOfFunctionPlot[781].Y = 2.794

	pointsOfFunctionPlot[782].X = 7.82
	pointsOfFunctionPlot[782].Y = 2.796

	pointsOfFunctionPlot[783].X = 7.83
	pointsOfFunctionPlot[783].Y = 2.798

	pointsOfFunctionPlot[784].X = 7.84
	pointsOfFunctionPlot[784].Y = 2.8

	pointsOfFunctionPlot[785].X = 7.85
	pointsOfFunctionPlot[785].Y = 2.801

	pointsOfFunctionPlot[786].X = 7.86
	pointsOfFunctionPlot[786].Y = 2.803

	pointsOfFunctionPlot[787].X = 7.87
	pointsOfFunctionPlot[787].Y = 2.805

	pointsOfFunctionPlot[788].X = 7.88
	pointsOfFunctionPlot[788].Y = 2.807

	pointsOfFunctionPlot[789].X = 7.89
	pointsOfFunctionPlot[789].Y = 2.808

	pointsOfFunctionPlot[790].X = 7.90
	pointsOfFunctionPlot[790].Y = 2.81

	pointsOfFunctionPlot[791].X = 7.91
	pointsOfFunctionPlot[791].Y = 2.812

	pointsOfFunctionPlot[792].X = 7.92
	pointsOfFunctionPlot[792].Y = 2.814

	pointsOfFunctionPlot[793].X = 7.93
	pointsOfFunctionPlot[793].Y = 2.816

	pointsOfFunctionPlot[794].X = 7.94
	pointsOfFunctionPlot[794].Y = 2.817

	pointsOfFunctionPlot[795].X = 7.95
	pointsOfFunctionPlot[795].Y = 2.819

	pointsOfFunctionPlot[796].X = 7.96
	pointsOfFunctionPlot[796].Y = 2.821

	pointsOfFunctionPlot[797].X = 7.97
	pointsOfFunctionPlot[797].Y = 2.823

	pointsOfFunctionPlot[798].X = 7.98
	pointsOfFunctionPlot[798].Y = 2.824

	pointsOfFunctionPlot[799].X = 7.99
	pointsOfFunctionPlot[799].Y = 2.826

	pointsOfFunctionPlot[800].X = 8.0
	pointsOfFunctionPlot[800].Y = 2.828

	pointsOfFunctionPlot[801].X = 8.01
	pointsOfFunctionPlot[801].Y = 2.83

	pointsOfFunctionPlot[802].X = 8.02
	pointsOfFunctionPlot[802].Y = 2.831

	pointsOfFunctionPlot[803].X = 8.03
	pointsOfFunctionPlot[803].Y = 2.833

	pointsOfFunctionPlot[804].X = 8.04
	pointsOfFunctionPlot[804].Y = 2.835

	pointsOfFunctionPlot[805].X = 8.05
	pointsOfFunctionPlot[805].Y = 2.837

	pointsOfFunctionPlot[806].X = 8.06
	pointsOfFunctionPlot[806].Y = 2.839

	pointsOfFunctionPlot[807].X = 8.07
	pointsOfFunctionPlot[807].Y = 2.84

	pointsOfFunctionPlot[808].X = 8.08
	pointsOfFunctionPlot[808].Y = 2.842

	pointsOfFunctionPlot[809].X = 8.09
	pointsOfFunctionPlot[809].Y = 2.844

	pointsOfFunctionPlot[810].X = 8.10
	pointsOfFunctionPlot[810].Y = 2.846

	pointsOfFunctionPlot[811].X = 8.11
	pointsOfFunctionPlot[811].Y = 2.847

	pointsOfFunctionPlot[812].X = 8.12
	pointsOfFunctionPlot[812].Y = 2.849

	pointsOfFunctionPlot[813].X = 8.13
	pointsOfFunctionPlot[813].Y = 2.851

	pointsOfFunctionPlot[814].X = 8.14
	pointsOfFunctionPlot[814].Y = 2.853

	pointsOfFunctionPlot[815].X = 8.15
	pointsOfFunctionPlot[815].Y = 2.854

	pointsOfFunctionPlot[816].X = 8.16
	pointsOfFunctionPlot[816].Y = 2.856

	pointsOfFunctionPlot[817].X = 8.17
	pointsOfFunctionPlot[817].Y = 2.858

	pointsOfFunctionPlot[818].X = 8.18
	pointsOfFunctionPlot[818].Y = 2.86

	pointsOfFunctionPlot[819].X = 8.19
	pointsOfFunctionPlot[819].Y = 2.86

	pointsOfFunctionPlot[820].X = 8.20
	pointsOfFunctionPlot[820].Y = 2.863

	pointsOfFunctionPlot[821].X = 8.21
	pointsOfFunctionPlot[821].Y = 2.865

	pointsOfFunctionPlot[822].X = 8.22
	pointsOfFunctionPlot[822].Y = 2.867

	pointsOfFunctionPlot[823].X = 8.23
	pointsOfFunctionPlot[823].Y = 2.868

	pointsOfFunctionPlot[824].X = 8.24
	pointsOfFunctionPlot[824].Y = 2.87

	pointsOfFunctionPlot[825].X = 8.25
	pointsOfFunctionPlot[825].Y = 2.872

	pointsOfFunctionPlot[826].X = 8.26
	pointsOfFunctionPlot[826].Y = 2.874

	pointsOfFunctionPlot[827].X = 8.27
	pointsOfFunctionPlot[827].Y = 2.875

	pointsOfFunctionPlot[828].X = 8.28
	pointsOfFunctionPlot[828].Y = 2.877

	pointsOfFunctionPlot[829].X = 8.29
	pointsOfFunctionPlot[829].Y = 2.879

	pointsOfFunctionPlot[830].X = 8.30
	pointsOfFunctionPlot[830].Y = 2.88

	pointsOfFunctionPlot[831].X = 8.31
	pointsOfFunctionPlot[831].Y = 2.882

	pointsOfFunctionPlot[832].X = 8.32
	pointsOfFunctionPlot[832].Y = 2.884

	pointsOfFunctionPlot[833].X = 8.33
	pointsOfFunctionPlot[833].Y = 2.886

	pointsOfFunctionPlot[834].X = 8.34
	pointsOfFunctionPlot[834].Y = 2.887

	pointsOfFunctionPlot[835].X = 8.35
	pointsOfFunctionPlot[835].Y = 2.889

	pointsOfFunctionPlot[836].X = 8.36
	pointsOfFunctionPlot[836].Y = 2.891

	pointsOfFunctionPlot[837].X = 8.37
	pointsOfFunctionPlot[837].Y = 2.893

	pointsOfFunctionPlot[838].X = 8.38
	pointsOfFunctionPlot[838].Y = 2.894

	pointsOfFunctionPlot[839].X = 8.39
	pointsOfFunctionPlot[839].Y = 2.896

	pointsOfFunctionPlot[840].X = 8.40
	pointsOfFunctionPlot[840].Y = 2.898

	pointsOfFunctionPlot[841].X = 8.41
	pointsOfFunctionPlot[841].Y = 2.9

	pointsOfFunctionPlot[842].X = 8.42
	pointsOfFunctionPlot[842].Y = 2.901

	pointsOfFunctionPlot[843].X = 8.43
	pointsOfFunctionPlot[843].Y = 2.903

	pointsOfFunctionPlot[844].X = 8.44
	pointsOfFunctionPlot[844].Y = 2.905

	pointsOfFunctionPlot[845].X = 8.45
	pointsOfFunctionPlot[845].Y = 2.906

	pointsOfFunctionPlot[846].X = 8.46
	pointsOfFunctionPlot[846].Y = 2.908

	pointsOfFunctionPlot[847].X = 8.47
	pointsOfFunctionPlot[847].Y = 2.91

	pointsOfFunctionPlot[848].X = 8.48
	pointsOfFunctionPlot[848].Y = 2.912

	pointsOfFunctionPlot[849].X = 8.49
	pointsOfFunctionPlot[849].Y = 2.913

	pointsOfFunctionPlot[850].X = 8.50
	pointsOfFunctionPlot[850].Y = 2.915

	pointsOfFunctionPlot[851].X = 8.51
	pointsOfFunctionPlot[851].Y = 2.917

	pointsOfFunctionPlot[852].X = 8.52
	pointsOfFunctionPlot[852].Y = 2.918

	pointsOfFunctionPlot[853].X = 8.53
	pointsOfFunctionPlot[853].Y = 2.92

	pointsOfFunctionPlot[854].X = 8.54
	pointsOfFunctionPlot[854].Y = 2.922

	pointsOfFunctionPlot[855].X = 8.55
	pointsOfFunctionPlot[855].Y = 2.924

	pointsOfFunctionPlot[856].X = 8.56
	pointsOfFunctionPlot[856].Y = 2.925

	pointsOfFunctionPlot[857].X = 8.57
	pointsOfFunctionPlot[857].Y = 2.927

	pointsOfFunctionPlot[858].X = 8.58
	pointsOfFunctionPlot[858].Y = 2.929

	pointsOfFunctionPlot[859].X = 8.59
	pointsOfFunctionPlot[859].Y = 2.93

	pointsOfFunctionPlot[860].X = 8.60
	pointsOfFunctionPlot[860].Y = 2.932

	pointsOfFunctionPlot[861].X = 8.61
	pointsOfFunctionPlot[861].Y = 2.934

	pointsOfFunctionPlot[862].X = 8.62
	pointsOfFunctionPlot[862].Y = 2.935

	pointsOfFunctionPlot[863].X = 8.63
	pointsOfFunctionPlot[863].Y = 2.937

	pointsOfFunctionPlot[864].X = 8.64
	pointsOfFunctionPlot[864].Y = 2.939

	pointsOfFunctionPlot[865].X = 8.65
	pointsOfFunctionPlot[865].Y = 2.941

	pointsOfFunctionPlot[866].X = 8.66
	pointsOfFunctionPlot[866].Y = 2.942

	pointsOfFunctionPlot[867].X = 8.67
	pointsOfFunctionPlot[867].Y = 2.944

	pointsOfFunctionPlot[868].X = 8.68
	pointsOfFunctionPlot[868].Y = 2.946

	pointsOfFunctionPlot[869].X = 8.69
	pointsOfFunctionPlot[869].Y = 2.947

	pointsOfFunctionPlot[870].X = 8.70
	pointsOfFunctionPlot[870].Y = 2.949

	pointsOfFunctionPlot[871].X = 8.71
	pointsOfFunctionPlot[871].Y = 2.951

	pointsOfFunctionPlot[872].X = 8.72
	pointsOfFunctionPlot[872].Y = 2.952

	pointsOfFunctionPlot[873].X = 8.73
	pointsOfFunctionPlot[873].Y = 2.954

	pointsOfFunctionPlot[874].X = 8.74
	pointsOfFunctionPlot[874].Y = 2.956

	pointsOfFunctionPlot[875].X = 8.75
	pointsOfFunctionPlot[875].Y = 2.958

	pointsOfFunctionPlot[876].X = 8.76
	pointsOfFunctionPlot[876].Y = 2.959

	pointsOfFunctionPlot[877].X = 8.77
	pointsOfFunctionPlot[877].Y = 2.961

	pointsOfFunctionPlot[878].X = 8.78
	pointsOfFunctionPlot[878].Y = 2.963

	pointsOfFunctionPlot[879].X = 8.79
	pointsOfFunctionPlot[879].Y = 2.964

	pointsOfFunctionPlot[880].X = 8.80
	pointsOfFunctionPlot[880].Y = 2.966

	pointsOfFunctionPlot[881].X = 8.81
	pointsOfFunctionPlot[881].Y = 2.968

	pointsOfFunctionPlot[882].X = 8.82
	pointsOfFunctionPlot[882].Y = 2.969

	pointsOfFunctionPlot[883].X = 8.83
	pointsOfFunctionPlot[883].Y = 2.971

	pointsOfFunctionPlot[884].X = 8.84
	pointsOfFunctionPlot[884].Y = 2.973

	pointsOfFunctionPlot[885].X = 8.85
	pointsOfFunctionPlot[885].Y = 2.974

	pointsOfFunctionPlot[886].X = 8.86
	pointsOfFunctionPlot[886].Y = 2.976

	pointsOfFunctionPlot[887].X = 8.87
	pointsOfFunctionPlot[887].Y = 2.978

	pointsOfFunctionPlot[888].X = 8.88
	pointsOfFunctionPlot[888].Y = 2.979

	pointsOfFunctionPlot[889].X = 8.89
	pointsOfFunctionPlot[889].Y = 2.981

	pointsOfFunctionPlot[890].X = 8.90
	pointsOfFunctionPlot[890].Y = 2.983

	pointsOfFunctionPlot[891].X = 8.91
	pointsOfFunctionPlot[891].Y = 2.984

	pointsOfFunctionPlot[892].X = 8.92
	pointsOfFunctionPlot[892].Y = 2.986

	pointsOfFunctionPlot[893].X = 8.93
	pointsOfFunctionPlot[893].Y = 2.988

	pointsOfFunctionPlot[894].X = 8.94
	pointsOfFunctionPlot[894].Y = 2.989

	pointsOfFunctionPlot[895].X = 8.95
	pointsOfFunctionPlot[895].Y = 2.991

	pointsOfFunctionPlot[896].X = 8.96
	pointsOfFunctionPlot[896].Y = 2.993

	pointsOfFunctionPlot[897].X = 8.97
	pointsOfFunctionPlot[897].Y = 2.994

	pointsOfFunctionPlot[898].X = 8.98
	pointsOfFunctionPlot[898].Y = 2.996

	pointsOfFunctionPlot[899].X = 8.99
	pointsOfFunctionPlot[899].Y = 2.998

	pointsOfFunctionPlot[900].X = 9.0
	pointsOfFunctionPlot[900].Y = 3.0

	pointsOfFunctionPlot[901].X = 9.01
	pointsOfFunctionPlot[901].Y = 3.001

	pointsOfFunctionPlot[902].X = 9.02
	pointsOfFunctionPlot[902].Y = 3.003

	pointsOfFunctionPlot[903].X = 9.03
	pointsOfFunctionPlot[903].Y = 3.004

	pointsOfFunctionPlot[904].X = 9.04
	pointsOfFunctionPlot[904].Y = 3.006

	pointsOfFunctionPlot[905].X = 9.05
	pointsOfFunctionPlot[905].Y = 3.008

	pointsOfFunctionPlot[906].X = 9.06
	pointsOfFunctionPlot[906].Y = 3.009

	pointsOfFunctionPlot[907].X = 9.07
	pointsOfFunctionPlot[907].Y = 3.011

	pointsOfFunctionPlot[908].X = 9.08
	pointsOfFunctionPlot[908].Y = 3.013

	pointsOfFunctionPlot[909].X = 9.09
	pointsOfFunctionPlot[909].Y = 3.014

	pointsOfFunctionPlot[910].X = 9.10
	pointsOfFunctionPlot[910].Y = 3.016

	pointsOfFunctionPlot[911].X = 9.11
	pointsOfFunctionPlot[911].Y = 3.018

	pointsOfFunctionPlot[912].X = 9.12
	pointsOfFunctionPlot[912].Y = 3.019

	pointsOfFunctionPlot[913].X = 9.13
	pointsOfFunctionPlot[913].Y = 3.021

	pointsOfFunctionPlot[914].X = 9.14
	pointsOfFunctionPlot[914].Y = 3.023

	pointsOfFunctionPlot[915].X = 9.15
	pointsOfFunctionPlot[915].Y = 3.024

	pointsOfFunctionPlot[916].X = 9.16
	pointsOfFunctionPlot[916].Y = 3.026

	pointsOfFunctionPlot[917].X = 9.17
	pointsOfFunctionPlot[917].Y = 3.028

	pointsOfFunctionPlot[918].X = 9.18
	pointsOfFunctionPlot[918].Y = 3.029

	pointsOfFunctionPlot[919].X = 9.19
	pointsOfFunctionPlot[919].Y = 3.031

	pointsOfFunctionPlot[920].X = 9.20
	pointsOfFunctionPlot[920].Y = 3.033

	pointsOfFunctionPlot[921].X = 9.21
	pointsOfFunctionPlot[921].Y = 3.034

	pointsOfFunctionPlot[922].X = 9.22
	pointsOfFunctionPlot[922].Y = 3.036

	pointsOfFunctionPlot[923].X = 9.23
	pointsOfFunctionPlot[923].Y = 3.038

	pointsOfFunctionPlot[924].X = 9.24
	pointsOfFunctionPlot[924].Y = 3.039

	pointsOfFunctionPlot[925].X = 9.25
	pointsOfFunctionPlot[925].Y = 3.041

	pointsOfFunctionPlot[926].X = 9.26
	pointsOfFunctionPlot[926].Y = 3.043

	pointsOfFunctionPlot[927].X = 9.27
	pointsOfFunctionPlot[927].Y = 3.044

	pointsOfFunctionPlot[928].X = 9.28
	pointsOfFunctionPlot[928].Y = 3.046

	pointsOfFunctionPlot[929].X = 9.29
	pointsOfFunctionPlot[929].Y = 3.047

	pointsOfFunctionPlot[930].X = 9.30
	pointsOfFunctionPlot[930].Y = 3.049

	pointsOfFunctionPlot[931].X = 9.31
	pointsOfFunctionPlot[931].Y = 3.051

	pointsOfFunctionPlot[932].X = 9.32
	pointsOfFunctionPlot[932].Y = 3.052

	pointsOfFunctionPlot[933].X = 9.33
	pointsOfFunctionPlot[933].Y = 3.054

	pointsOfFunctionPlot[934].X = 9.34
	pointsOfFunctionPlot[934].Y = 3.056

	pointsOfFunctionPlot[935].X = 9.35
	pointsOfFunctionPlot[935].Y = 3.057

	pointsOfFunctionPlot[936].X = 9.36
	pointsOfFunctionPlot[936].Y = 3.059

	pointsOfFunctionPlot[937].X = 9.37
	pointsOfFunctionPlot[937].Y = 3.061

	pointsOfFunctionPlot[938].X = 9.38
	pointsOfFunctionPlot[938].Y = 3.062

	pointsOfFunctionPlot[939].X = 9.39
	pointsOfFunctionPlot[939].Y = 3.064

	pointsOfFunctionPlot[940].X = 9.40
	pointsOfFunctionPlot[940].Y = 3.065

	pointsOfFunctionPlot[941].X = 9.41
	pointsOfFunctionPlot[941].Y = 3.067

	pointsOfFunctionPlot[942].X = 9.42
	pointsOfFunctionPlot[942].Y = 3.069

	pointsOfFunctionPlot[943].X = 9.43
	pointsOfFunctionPlot[943].Y = 3.07

	pointsOfFunctionPlot[944].X = 9.44
	pointsOfFunctionPlot[944].Y = 3.072

	pointsOfFunctionPlot[945].X = 9.45
	pointsOfFunctionPlot[945].Y = 3.074

	pointsOfFunctionPlot[946].X = 9.46
	pointsOfFunctionPlot[946].Y = 3.075

	pointsOfFunctionPlot[947].X = 9.47
	pointsOfFunctionPlot[947].Y = 3.077

	pointsOfFunctionPlot[948].X = 9.48
	pointsOfFunctionPlot[948].Y = 3.078

	pointsOfFunctionPlot[949].X = 9.49
	pointsOfFunctionPlot[949].Y = 3.08

	pointsOfFunctionPlot[950].X = 9.50
	pointsOfFunctionPlot[950].Y = 3.082

	pointsOfFunctionPlot[951].X = 9.51
	pointsOfFunctionPlot[951].Y = 3.083

	pointsOfFunctionPlot[952].X = 9.52
	pointsOfFunctionPlot[952].Y = 3.085

	pointsOfFunctionPlot[953].X = 9.53
	pointsOfFunctionPlot[953].Y = 3.087

	pointsOfFunctionPlot[954].X = 9.54
	pointsOfFunctionPlot[954].Y = 3.088

	pointsOfFunctionPlot[955].X = 9.55
	pointsOfFunctionPlot[955].Y = 3.09

	pointsOfFunctionPlot[956].X = 9.56
	pointsOfFunctionPlot[956].Y = 3.091

	pointsOfFunctionPlot[957].X = 9.57
	pointsOfFunctionPlot[957].Y = 3.093

	pointsOfFunctionPlot[958].X = 9.58
	pointsOfFunctionPlot[958].Y = 3.095

	pointsOfFunctionPlot[959].X = 9.59
	pointsOfFunctionPlot[959].Y = 3.096

	pointsOfFunctionPlot[960].X = 9.60
	pointsOfFunctionPlot[960].Y = 3.098

	pointsOfFunctionPlot[961].X = 9.61
	pointsOfFunctionPlot[961].Y = 3.1

	pointsOfFunctionPlot[962].X = 9.62
	pointsOfFunctionPlot[962].Y = 3.101

	pointsOfFunctionPlot[963].X = 9.63
	pointsOfFunctionPlot[963].Y = 3.103

	pointsOfFunctionPlot[964].X = 9.64
	pointsOfFunctionPlot[964].Y = 3.104

	pointsOfFunctionPlot[965].X = 9.65
	pointsOfFunctionPlot[965].Y = 3.106

	pointsOfFunctionPlot[966].X = 9.66
	pointsOfFunctionPlot[966].Y = 3.108

	pointsOfFunctionPlot[967].X = 9.67
	pointsOfFunctionPlot[967].Y = 3.109

	pointsOfFunctionPlot[968].X = 9.68
	pointsOfFunctionPlot[968].Y = 3.111

	pointsOfFunctionPlot[969].X = 9.69
	pointsOfFunctionPlot[969].Y = 3.112

	pointsOfFunctionPlot[970].X = 9.70
	pointsOfFunctionPlot[970].Y = 3.114

	pointsOfFunctionPlot[971].X = 9.71
	pointsOfFunctionPlot[971].Y = 3.116

	pointsOfFunctionPlot[972].X = 9.72
	pointsOfFunctionPlot[972].Y = 3.117

	pointsOfFunctionPlot[973].X = 9.73
	pointsOfFunctionPlot[973].Y = 3.119

	pointsOfFunctionPlot[974].X = 9.74
	pointsOfFunctionPlot[974].Y = 3.12

	pointsOfFunctionPlot[975].X = 9.75
	pointsOfFunctionPlot[975].Y = 3.122

	pointsOfFunctionPlot[976].X = 9.76
	pointsOfFunctionPlot[976].Y = 3.124

	pointsOfFunctionPlot[977].X = 9.77
	pointsOfFunctionPlot[977].Y = 3.125

	pointsOfFunctionPlot[978].X = 9.78
	pointsOfFunctionPlot[978].Y = 3.127

	pointsOfFunctionPlot[979].X = 9.79
	pointsOfFunctionPlot[979].Y = 3.128

	pointsOfFunctionPlot[980].X = 9.80
	pointsOfFunctionPlot[980].Y = 3.13

	pointsOfFunctionPlot[981].X = 9.81
	pointsOfFunctionPlot[981].Y = 3.132

	pointsOfFunctionPlot[982].X = 9.82
	pointsOfFunctionPlot[982].Y = 3.133

	pointsOfFunctionPlot[983].X = 9.83
	pointsOfFunctionPlot[983].Y = 3.135

	pointsOfFunctionPlot[984].X = 9.84
	pointsOfFunctionPlot[984].Y = 3.136

	pointsOfFunctionPlot[985].X = 9.85
	pointsOfFunctionPlot[985].Y = 3.138

	pointsOfFunctionPlot[986].X = 9.86
	pointsOfFunctionPlot[986].Y = 3.14

	pointsOfFunctionPlot[987].X = 9.87
	pointsOfFunctionPlot[987].Y = 3.141

	pointsOfFunctionPlot[988].X = 9.88
	pointsOfFunctionPlot[988].Y = 3.143

	pointsOfFunctionPlot[989].X = 9.89
	pointsOfFunctionPlot[989].Y = 3.144

	pointsOfFunctionPlot[990].X = 9.90
	pointsOfFunctionPlot[990].Y = 3.146

	pointsOfFunctionPlot[991].X = 9.91
	pointsOfFunctionPlot[991].Y = 3.148

	pointsOfFunctionPlot[992].X = 9.92
	pointsOfFunctionPlot[992].Y = 3.149

	pointsOfFunctionPlot[993].X = 9.93
	pointsOfFunctionPlot[993].Y = 3.151

	pointsOfFunctionPlot[994].X = 9.94
	pointsOfFunctionPlot[994].Y = 3.152

	pointsOfFunctionPlot[995].X = 9.95
	pointsOfFunctionPlot[995].Y = 3.154

	pointsOfFunctionPlot[996].X = 9.96
	pointsOfFunctionPlot[996].Y = 3.155

	pointsOfFunctionPlot[997].X = 9.97
	pointsOfFunctionPlot[997].Y = 3.157

	pointsOfFunctionPlot[998].X = 9.98
	pointsOfFunctionPlot[998].Y = 3.159

	pointsOfFunctionPlot[999].X = 9.99
	pointsOfFunctionPlot[999].Y = 3.16

	pointsOfFunctionPlot[1_000].X = 10.0
	pointsOfFunctionPlot[1_000].Y = 3.162

	pointsOfFunctionPlot[1_001].X = 10.01
	pointsOfFunctionPlot[1_001].Y = 3.163

	pointsOfFunctionPlot[1_002].X = 10.02
	pointsOfFunctionPlot[1_002].Y = 3.165

	pointsOfFunctionPlot[1_003].X = 10.03
	pointsOfFunctionPlot[1_003].Y = 3.167

	pointsOfFunctionPlot[1_004].X = 10.04
	pointsOfFunctionPlot[1_004].Y = 3.168

	pointsOfFunctionPlot[1_005].X = 10.05
	pointsOfFunctionPlot[1_005].Y = 3.17

	pointsOfFunctionPlot[1_006].X = 10.06
	pointsOfFunctionPlot[1_006].Y = 3.171

	pointsOfFunctionPlot[1_007].X = 10.07
	pointsOfFunctionPlot[1_007].Y = 3.173

	pointsOfFunctionPlot[1_008].X = 10.08
	pointsOfFunctionPlot[1_008].Y = 3.174

	pointsOfFunctionPlot[1_009].X = 10.09
	pointsOfFunctionPlot[1_009].Y = 3.176

	pointsOfFunctionPlot[1_010].X = 10.10
	pointsOfFunctionPlot[1_010].Y = 3.178

	pointsOfFunctionPlot[1_011].X = 10.11
	pointsOfFunctionPlot[1_011].Y = 3.179

	pointsOfFunctionPlot[1_012].X = 10.12
	pointsOfFunctionPlot[1_012].Y = 3.181

	pointsOfFunctionPlot[1_013].X = 10.13
	pointsOfFunctionPlot[1_013].Y = 3.182

	pointsOfFunctionPlot[1_014].X = 10.14
	pointsOfFunctionPlot[1_014].Y = 3.184

	pointsOfFunctionPlot[1_015].X = 10.15
	pointsOfFunctionPlot[1_015].Y = 3.185

	pointsOfFunctionPlot[1_016].X = 10.16
	pointsOfFunctionPlot[1_016].Y = 3.187

	pointsOfFunctionPlot[1_017].X = 10.17
	pointsOfFunctionPlot[1_017].Y = 3.189

	pointsOfFunctionPlot[1_018].X = 10.18
	pointsOfFunctionPlot[1_018].Y = 3.19

	pointsOfFunctionPlot[1_019].X = 10.19
	pointsOfFunctionPlot[1_019].Y = 3.192

	pointsOfFunctionPlot[1_020].X = 10.20
	pointsOfFunctionPlot[1_020].Y = 3.193

	pointsOfFunctionPlot[1_021].X = 10.21
	pointsOfFunctionPlot[1_021].Y = 3.195

	pointsOfFunctionPlot[1_022].X = 10.22
	pointsOfFunctionPlot[1_022].Y = 3.196

	pointsOfFunctionPlot[1_023].X = 10.23
	pointsOfFunctionPlot[1_023].Y = 3.198

	pointsOfFunctionPlot[1_024].X = 10.24
	pointsOfFunctionPlot[1_024].Y = 3.2

	pointsOfFunctionPlot[1_025].X = 10.25
	pointsOfFunctionPlot[1_025].Y = 3.201

	pointsOfFunctionPlot[1_026].X = 10.26
	pointsOfFunctionPlot[1_026].Y = 3.203

	pointsOfFunctionPlot[1_027].X = 10.27
	pointsOfFunctionPlot[1_027].Y = 3.204

	pointsOfFunctionPlot[1_028].X = 10.28
	pointsOfFunctionPlot[1_028].Y = 3.206

	pointsOfFunctionPlot[1_029].X = 10.29
	pointsOfFunctionPlot[1_029].Y = 3.207

	pointsOfFunctionPlot[1_030].X = 10.30
	pointsOfFunctionPlot[1_030].Y = 3.209

	pointsOfFunctionPlot[1_031].X = 10.31
	pointsOfFunctionPlot[1_031].Y = 3.21

	pointsOfFunctionPlot[1_032].X = 10.32
	pointsOfFunctionPlot[1_032].Y = 3.212

	pointsOfFunctionPlot[1_033].X = 10.33
	pointsOfFunctionPlot[1_033].Y = 3.214

	pointsOfFunctionPlot[1_034].X = 10.34
	pointsOfFunctionPlot[1_034].Y = 3.215

	pointsOfFunctionPlot[1_035].X = 10.35
	pointsOfFunctionPlot[1_035].Y = 3.217

	pointsOfFunctionPlot[1_036].X = 10.36
	pointsOfFunctionPlot[1_036].Y = 3.218

	pointsOfFunctionPlot[1_037].X = 10.37
	pointsOfFunctionPlot[1_037].Y = 3.22

	pointsOfFunctionPlot[1_038].X = 10.38
	pointsOfFunctionPlot[1_038].Y = 3.221

	pointsOfFunctionPlot[1_039].X = 10.39
	pointsOfFunctionPlot[1_039].Y = 3.223

	pointsOfFunctionPlot[1_040].X = 10.40
	pointsOfFunctionPlot[1_040].Y = 3.224

	pointsOfFunctionPlot[1_041].X = 10.41
	pointsOfFunctionPlot[1_041].Y = 3.226

	pointsOfFunctionPlot[1_042].X = 10.42
	pointsOfFunctionPlot[1_042].Y = 3.228

	pointsOfFunctionPlot[1_043].X = 10.43
	pointsOfFunctionPlot[1_043].Y = 3.229

	pointsOfFunctionPlot[1_044].X = 10.44
	pointsOfFunctionPlot[1_044].Y = 3.231

	pointsOfFunctionPlot[1_045].X = 10.45
	pointsOfFunctionPlot[1_045].Y = 3.232

	pointsOfFunctionPlot[1_046].X = 10.46
	pointsOfFunctionPlot[1_046].Y = 3.234

	pointsOfFunctionPlot[1_047].X = 10.47
	pointsOfFunctionPlot[1_047].Y = 3.235

	pointsOfFunctionPlot[1_048].X = 10.48
	pointsOfFunctionPlot[1_048].Y = 3.237

	pointsOfFunctionPlot[1_049].X = 10.49
	pointsOfFunctionPlot[1_049].Y = 3.238

	pointsOfFunctionPlot[1_050].X = 10.50
	pointsOfFunctionPlot[1_050].Y = 3.24

	pointsOfFunctionPlot[1_051].X = 10.51
	pointsOfFunctionPlot[1_051].Y = 3.241

	pointsOfFunctionPlot[1_052].X = 10.52
	pointsOfFunctionPlot[1_052].Y = 3.243

	pointsOfFunctionPlot[1_053].X = 10.53
	pointsOfFunctionPlot[1_053].Y = 3.244

	pointsOfFunctionPlot[1_054].X = 10.54
	pointsOfFunctionPlot[1_054].Y = 3.246

	pointsOfFunctionPlot[1_055].X = 10.55
	pointsOfFunctionPlot[1_055].Y = 3.248

	pointsOfFunctionPlot[1_056].X = 10.56
	pointsOfFunctionPlot[1_056].Y = 3.249

	pointsOfFunctionPlot[1_057].X = 10.57
	pointsOfFunctionPlot[1_057].Y = 3.251

	pointsOfFunctionPlot[1_058].X = 10.58
	pointsOfFunctionPlot[1_058].Y = 3.252

	pointsOfFunctionPlot[1_059].X = 10.59
	pointsOfFunctionPlot[1_059].Y = 3.254

	pointsOfFunctionPlot[1_060].X = 10.60
	pointsOfFunctionPlot[1_060].Y = 3.255

	pointsOfFunctionPlot[1_061].X = 10.61
	pointsOfFunctionPlot[1_061].Y = 3.257

	pointsOfFunctionPlot[1_062].X = 10.62
	pointsOfFunctionPlot[1_062].Y = 3.258

	pointsOfFunctionPlot[1_063].X = 10.63
	pointsOfFunctionPlot[1_063].Y = 3.26

	pointsOfFunctionPlot[1_064].X = 10.64
	pointsOfFunctionPlot[1_064].Y = 3.261

	pointsOfFunctionPlot[1_065].X = 10.65
	pointsOfFunctionPlot[1_065].Y = 3.263

	pointsOfFunctionPlot[1_066].X = 10.66
	pointsOfFunctionPlot[1_066].Y = 3.264

	pointsOfFunctionPlot[1_067].X = 10.67
	pointsOfFunctionPlot[1_067].Y = 3.266

	pointsOfFunctionPlot[1_068].X = 10.68
	pointsOfFunctionPlot[1_068].Y = 3.268

	pointsOfFunctionPlot[1_069].X = 10.69
	pointsOfFunctionPlot[1_069].Y = 3.269

	pointsOfFunctionPlot[1_070].X = 10.70
	pointsOfFunctionPlot[1_070].Y = 3.271

	pointsOfFunctionPlot[1_071].X = 10.71
	pointsOfFunctionPlot[1_071].Y = 3.272

	pointsOfFunctionPlot[1_072].X = 10.72
	pointsOfFunctionPlot[1_072].Y = 3.274

	pointsOfFunctionPlot[1_073].X = 10.73
	pointsOfFunctionPlot[1_073].Y = 3.275

	pointsOfFunctionPlot[1_074].X = 10.74
	pointsOfFunctionPlot[1_074].Y = 3.277

	pointsOfFunctionPlot[1_075].X = 10.75
	pointsOfFunctionPlot[1_075].Y = 3.278

	pointsOfFunctionPlot[1_076].X = 10.76
	pointsOfFunctionPlot[1_076].Y = 3.28

	pointsOfFunctionPlot[1_077].X = 10.77
	pointsOfFunctionPlot[1_077].Y = 3.281

	pointsOfFunctionPlot[1_078].X = 10.78
	pointsOfFunctionPlot[1_078].Y = 3.283

	pointsOfFunctionPlot[1_079].X = 10.79
	pointsOfFunctionPlot[1_079].Y = 3.284

	pointsOfFunctionPlot[1_080].X = 10.80
	pointsOfFunctionPlot[1_080].Y = 3.286

	pointsOfFunctionPlot[1_081].X = 10.81
	pointsOfFunctionPlot[1_081].Y = 3.287

	pointsOfFunctionPlot[1_082].X = 10.82
	pointsOfFunctionPlot[1_082].Y = 3.289

	pointsOfFunctionPlot[1_083].X = 10.83
	pointsOfFunctionPlot[1_083].Y = 3.29

	pointsOfFunctionPlot[1_084].X = 10.84
	pointsOfFunctionPlot[1_084].Y = 3.292

	pointsOfFunctionPlot[1_085].X = 10.85
	pointsOfFunctionPlot[1_085].Y = 3.293

	pointsOfFunctionPlot[1_086].X = 10.86
	pointsOfFunctionPlot[1_086].Y = 3.295

	pointsOfFunctionPlot[1_087].X = 10.87
	pointsOfFunctionPlot[1_087].Y = 3.296

	pointsOfFunctionPlot[1_088].X = 10.88
	pointsOfFunctionPlot[1_088].Y = 3.298

	pointsOfFunctionPlot[1_089].X = 10.89
	pointsOfFunctionPlot[1_089].Y = 3.3

	pointsOfFunctionPlot[1_090].X = 10.90
	pointsOfFunctionPlot[1_090].Y = 3.301

	pointsOfFunctionPlot[1_091].X = 10.91
	pointsOfFunctionPlot[1_091].Y = 3.303

	pointsOfFunctionPlot[1_092].X = 10.92
	pointsOfFunctionPlot[1_092].Y = 3.304

	pointsOfFunctionPlot[1_093].X = 10.93
	pointsOfFunctionPlot[1_093].Y = 3.306

	pointsOfFunctionPlot[1_094].X = 10.94
	pointsOfFunctionPlot[1_094].Y = 3.307

	pointsOfFunctionPlot[1_095].X = 10.95
	pointsOfFunctionPlot[1_095].Y = 3.309

	pointsOfFunctionPlot[1_096].X = 10.96
	pointsOfFunctionPlot[1_096].Y = 3.31

	pointsOfFunctionPlot[1_097].X = 10.97
	pointsOfFunctionPlot[1_097].Y = 3.312

	pointsOfFunctionPlot[1_098].X = 10.98
	pointsOfFunctionPlot[1_098].Y = 3.313

	pointsOfFunctionPlot[1_099].X = 10.99
	pointsOfFunctionPlot[1_099].Y = 3.315

	pointsOfFunctionPlot[1_100].X = 11.0
	pointsOfFunctionPlot[1_100].Y = 3.316

	pointsOfFunctionPlot[1_101].X = 11.01
	pointsOfFunctionPlot[1_101].Y = 3.318

	pointsOfFunctionPlot[1_102].X = 11.02
	pointsOfFunctionPlot[1_102].Y = 3.319

	pointsOfFunctionPlot[1_103].X = 11.03
	pointsOfFunctionPlot[1_103].Y = 3.321

	pointsOfFunctionPlot[1_104].X = 11.04
	pointsOfFunctionPlot[1_104].Y = 3.322

	pointsOfFunctionPlot[1_105].X = 11.05
	pointsOfFunctionPlot[1_105].Y = 3.324

	pointsOfFunctionPlot[1_106].X = 11.06
	pointsOfFunctionPlot[1_106].Y = 3.325

	pointsOfFunctionPlot[1_107].X = 11.07
	pointsOfFunctionPlot[1_107].Y = 3.327

	pointsOfFunctionPlot[1_108].X = 11.08
	pointsOfFunctionPlot[1_108].Y = 3.328

	pointsOfFunctionPlot[1_109].X = 11.09
	pointsOfFunctionPlot[1_109].Y = 3.33

	pointsOfFunctionPlot[1_110].X = 11.10
	pointsOfFunctionPlot[1_110].Y = 3.331

	pointsOfFunctionPlot[1_111].X = 11.11
	pointsOfFunctionPlot[1_111].Y = 3.333

	pointsOfFunctionPlot[1_112].X = 11.12
	pointsOfFunctionPlot[1_112].Y = 3.334

	pointsOfFunctionPlot[1_113].X = 11.13
	pointsOfFunctionPlot[1_113].Y = 3.336

	pointsOfFunctionPlot[1_114].X = 11.14
	pointsOfFunctionPlot[1_114].Y = 3.337

	pointsOfFunctionPlot[1_115].X = 11.15
	pointsOfFunctionPlot[1_115].Y = 3.339

	pointsOfFunctionPlot[1_116].X = 11.16
	pointsOfFunctionPlot[1_116].Y = 3.34

	pointsOfFunctionPlot[1_117].X = 11.17
	pointsOfFunctionPlot[1_117].Y = 3.342

	pointsOfFunctionPlot[1_118].X = 11.18
	pointsOfFunctionPlot[1_118].Y = 3.343

	pointsOfFunctionPlot[1_119].X = 11.19
	pointsOfFunctionPlot[1_119].Y = 3.345

	pointsOfFunctionPlot[1_120].X = 11.20
	pointsOfFunctionPlot[1_120].Y = 3.346

	pointsOfFunctionPlot[1_121].X = 11.21
	pointsOfFunctionPlot[1_121].Y = 3.348

	pointsOfFunctionPlot[1_122].X = 11.22
	pointsOfFunctionPlot[1_122].Y = 3.349

	pointsOfFunctionPlot[1_123].X = 11.23
	pointsOfFunctionPlot[1_123].Y = 3.351

	pointsOfFunctionPlot[1_124].X = 11.24
	pointsOfFunctionPlot[1_124].Y = 3.352

	pointsOfFunctionPlot[1_125].X = 11.25
	pointsOfFunctionPlot[1_125].Y = 3.354

	pointsOfFunctionPlot[1_126].X = 11.26
	pointsOfFunctionPlot[1_126].Y = 3.355

	pointsOfFunctionPlot[1_127].X = 11.27
	pointsOfFunctionPlot[1_127].Y = 3.357

	pointsOfFunctionPlot[1_128].X = 11.28
	pointsOfFunctionPlot[1_128].Y = 3.358

	pointsOfFunctionPlot[1_129].X = 11.29
	pointsOfFunctionPlot[1_129].Y = 3.36

	pointsOfFunctionPlot[1_130].X = 11.30
	pointsOfFunctionPlot[1_130].Y = 3.361

	pointsOfFunctionPlot[1_131].X = 11.31
	pointsOfFunctionPlot[1_131].Y = 3.363

	pointsOfFunctionPlot[1_132].X = 11.32
	pointsOfFunctionPlot[1_132].Y = 3.364

	pointsOfFunctionPlot[1_133].X = 11.33
	pointsOfFunctionPlot[1_133].Y = 3.366

	pointsOfFunctionPlot[1_134].X = 11.34
	pointsOfFunctionPlot[1_134].Y = 3.367

	pointsOfFunctionPlot[1_135].X = 11.35
	pointsOfFunctionPlot[1_135].Y = 3.368

	pointsOfFunctionPlot[1_136].X = 11.36
	pointsOfFunctionPlot[1_136].Y = 3.37

	pointsOfFunctionPlot[1_137].X = 11.37
	pointsOfFunctionPlot[1_137].Y = 3.371

	pointsOfFunctionPlot[1_138].X = 11.38
	pointsOfFunctionPlot[1_138].Y = 3.373

	pointsOfFunctionPlot[1_139].X = 11.39
	pointsOfFunctionPlot[1_139].Y = 3.374

	pointsOfFunctionPlot[1_140].X = 11.40
	pointsOfFunctionPlot[1_140].Y = 3.376

	pointsOfFunctionPlot[1_141].X = 11.41
	pointsOfFunctionPlot[1_141].Y = 3.377

	pointsOfFunctionPlot[1_142].X = 11.42
	pointsOfFunctionPlot[1_142].Y = 3.379

	pointsOfFunctionPlot[1_143].X = 11.43
	pointsOfFunctionPlot[1_143].Y = 3.38

	pointsOfFunctionPlot[1_144].X = 11.44
	pointsOfFunctionPlot[1_144].Y = 3.382

	pointsOfFunctionPlot[1_145].X = 11.45
	pointsOfFunctionPlot[1_145].Y = 3.383

	pointsOfFunctionPlot[1_146].X = 11.46
	pointsOfFunctionPlot[1_146].Y = 3.385

	pointsOfFunctionPlot[1_147].X = 11.47
	pointsOfFunctionPlot[1_147].Y = 3.386

	pointsOfFunctionPlot[1_148].X = 11.48
	pointsOfFunctionPlot[1_148].Y = 3.388

	pointsOfFunctionPlot[1_149].X = 11.49
	pointsOfFunctionPlot[1_149].Y = 3.389

	pointsOfFunctionPlot[1_150].X = 11.50
	pointsOfFunctionPlot[1_150].Y = 3.391

	pointsOfFunctionPlot[1_151].X = 11.51
	pointsOfFunctionPlot[1_151].Y = 3.392

	pointsOfFunctionPlot[1_152].X = 11.52
	pointsOfFunctionPlot[1_152].Y = 3.394

	pointsOfFunctionPlot[1_153].X = 11.53
	pointsOfFunctionPlot[1_153].Y = 3.395

	pointsOfFunctionPlot[1_154].X = 11.54
	pointsOfFunctionPlot[1_154].Y = 3.397

	pointsOfFunctionPlot[1_155].X = 11.55
	pointsOfFunctionPlot[1_155].Y = 3.398

	pointsOfFunctionPlot[1_156].X = 11.56
	pointsOfFunctionPlot[1_156].Y = 3.4

	pointsOfFunctionPlot[1_157].X = 11.57
	pointsOfFunctionPlot[1_157].Y = 3.401

	pointsOfFunctionPlot[1_158].X = 11.58
	pointsOfFunctionPlot[1_158].Y = 3.402

	pointsOfFunctionPlot[1_159].X = 11.59
	pointsOfFunctionPlot[1_159].Y = 3.404

	pointsOfFunctionPlot[1_160].X = 11.60
	pointsOfFunctionPlot[1_160].Y = 3.405

	pointsOfFunctionPlot[1_161].X = 11.61
	pointsOfFunctionPlot[1_161].Y = 3.407

	pointsOfFunctionPlot[1_162].X = 11.62
	pointsOfFunctionPlot[1_162].Y = 3.408

	pointsOfFunctionPlot[1_163].X = 11.63
	pointsOfFunctionPlot[1_163].Y = 3.41

	pointsOfFunctionPlot[1_164].X = 11.64
	pointsOfFunctionPlot[1_164].Y = 3.411

	pointsOfFunctionPlot[1_165].X = 11.65
	pointsOfFunctionPlot[1_165].Y = 3.413

	pointsOfFunctionPlot[1_166].X = 11.66
	pointsOfFunctionPlot[1_166].Y = 3.414

	pointsOfFunctionPlot[1_167].X = 11.67
	pointsOfFunctionPlot[1_167].Y = 3.416

	pointsOfFunctionPlot[1_168].X = 11.68
	pointsOfFunctionPlot[1_168].Y = 3.417

	pointsOfFunctionPlot[1_169].X = 11.69
	pointsOfFunctionPlot[1_169].Y = 3.419

	pointsOfFunctionPlot[1_170].X = 11.70
	pointsOfFunctionPlot[1_170].Y = 3.42

	pointsOfFunctionPlot[1_171].X = 11.71
	pointsOfFunctionPlot[1_171].Y = 3.421

	pointsOfFunctionPlot[1_172].X = 11.72
	pointsOfFunctionPlot[1_172].Y = 3.423

	pointsOfFunctionPlot[1_173].X = 11.73
	pointsOfFunctionPlot[1_173].Y = 3.424

	pointsOfFunctionPlot[1_174].X = 11.74
	pointsOfFunctionPlot[1_174].Y = 3.426

	pointsOfFunctionPlot[1_175].X = 11.75
	pointsOfFunctionPlot[1_175].Y = 3.427

	pointsOfFunctionPlot[1_176].X = 11.76
	pointsOfFunctionPlot[1_176].Y = 3.429

	pointsOfFunctionPlot[1_177].X = 11.77
	pointsOfFunctionPlot[1_177].Y = 3.43

	pointsOfFunctionPlot[1_178].X = 11.78
	pointsOfFunctionPlot[1_178].Y = 3.432

	pointsOfFunctionPlot[1_179].X = 11.79
	pointsOfFunctionPlot[1_179].Y = 3.433

	pointsOfFunctionPlot[1_180].X = 11.80
	pointsOfFunctionPlot[1_180].Y = 3.435

	pointsOfFunctionPlot[1_181].X = 11.81
	pointsOfFunctionPlot[1_181].Y = 3.436

	pointsOfFunctionPlot[1_182].X = 11.82
	pointsOfFunctionPlot[1_182].Y = 3.438

	pointsOfFunctionPlot[1_183].X = 11.83
	pointsOfFunctionPlot[1_183].Y = 3.439

	pointsOfFunctionPlot[1_184].X = 11.84
	pointsOfFunctionPlot[1_184].Y = 3.44

	pointsOfFunctionPlot[1_185].X = 11.85
	pointsOfFunctionPlot[1_185].Y = 3.442

	pointsOfFunctionPlot[1_186].X = 11.86
	pointsOfFunctionPlot[1_186].Y = 3.443

	pointsOfFunctionPlot[1_187].X = 11.87
	pointsOfFunctionPlot[1_187].Y = 3.445

	pointsOfFunctionPlot[1_188].X = 11.88
	pointsOfFunctionPlot[1_188].Y = 3.446

	pointsOfFunctionPlot[1_189].X = 11.89
	pointsOfFunctionPlot[1_189].Y = 3.448

	pointsOfFunctionPlot[1_190].X = 11.90
	pointsOfFunctionPlot[1_190].Y = 3.449

	pointsOfFunctionPlot[1_191].X = 11.91
	pointsOfFunctionPlot[1_191].Y = 3.451

	pointsOfFunctionPlot[1_192].X = 11.92
	pointsOfFunctionPlot[1_192].Y = 3.452

	pointsOfFunctionPlot[1_193].X = 11.93
	pointsOfFunctionPlot[1_193].Y = 3.453

	pointsOfFunctionPlot[1_194].X = 11.94
	pointsOfFunctionPlot[1_194].Y = 3.455

	pointsOfFunctionPlot[1_195].X = 11.95
	pointsOfFunctionPlot[1_195].Y = 3.456

	pointsOfFunctionPlot[1_196].X = 11.96
	pointsOfFunctionPlot[1_196].Y = 3.458

	pointsOfFunctionPlot[1_197].X = 11.97
	pointsOfFunctionPlot[1_197].Y = 3.459

	pointsOfFunctionPlot[1_198].X = 11.98
	pointsOfFunctionPlot[1_198].Y = 3.461

	pointsOfFunctionPlot[1_199].X = 11.99
	pointsOfFunctionPlot[1_199].Y = 3.462

	pointsOfFunctionPlot[1_200].X = 12.0
	pointsOfFunctionPlot[1_200].Y = 3.464

	pointsOfFunctionPlot[1_201].X = 12.01
	pointsOfFunctionPlot[1_201].Y = 3.465

	pointsOfFunctionPlot[1_202].X = 12.02
	pointsOfFunctionPlot[1_202].Y = 3.466

	pointsOfFunctionPlot[1_203].X = 12.03
	pointsOfFunctionPlot[1_203].Y = 3.468

	pointsOfFunctionPlot[1_204].X = 12.04
	pointsOfFunctionPlot[1_204].Y = 3.469

	pointsOfFunctionPlot[1_205].X = 12.05
	pointsOfFunctionPlot[1_205].Y = 3.471

	pointsOfFunctionPlot[1_206].X = 12.06
	pointsOfFunctionPlot[1_206].Y = 3.472

	pointsOfFunctionPlot[1_207].X = 12.07
	pointsOfFunctionPlot[1_207].Y = 3.474

	pointsOfFunctionPlot[1_208].X = 12.08
	pointsOfFunctionPlot[1_208].Y = 3.475

	pointsOfFunctionPlot[1_209].X = 12.09
	pointsOfFunctionPlot[1_209].Y = 3.477

	pointsOfFunctionPlot[1_210].X = 12.10
	pointsOfFunctionPlot[1_210].Y = 3.478

	pointsOfFunctionPlot[1_211].X = 12.11
	pointsOfFunctionPlot[1_211].Y = 3.479

	pointsOfFunctionPlot[1_212].X = 12.12
	pointsOfFunctionPlot[1_212].Y = 3.481

	pointsOfFunctionPlot[1_213].X = 12.13
	pointsOfFunctionPlot[1_213].Y = 3.482

	pointsOfFunctionPlot[1_214].X = 12.14
	pointsOfFunctionPlot[1_214].Y = 3.484

	pointsOfFunctionPlot[1_215].X = 12.15
	pointsOfFunctionPlot[1_215].Y = 3.485

	pointsOfFunctionPlot[1_216].X = 12.16
	pointsOfFunctionPlot[1_216].Y = 3.487

	pointsOfFunctionPlot[1_217].X = 12.17
	pointsOfFunctionPlot[1_217].Y = 3.488

	pointsOfFunctionPlot[1_218].X = 12.18
	pointsOfFunctionPlot[1_218].Y = 3.489

	pointsOfFunctionPlot[1_219].X = 12.19
	pointsOfFunctionPlot[1_219].Y = 3.491

	pointsOfFunctionPlot[1_220].X = 12.20
	pointsOfFunctionPlot[1_220].Y = 3.492

	pointsOfFunctionPlot[1_221].X = 12.21
	pointsOfFunctionPlot[1_221].Y = 3.494

	pointsOfFunctionPlot[1_222].X = 12.22
	pointsOfFunctionPlot[1_222].Y = 3.495

	pointsOfFunctionPlot[1_223].X = 12.23
	pointsOfFunctionPlot[1_223].Y = 3.497

	pointsOfFunctionPlot[1_224].X = 12.24
	pointsOfFunctionPlot[1_224].Y = 3.498

	pointsOfFunctionPlot[1_225].X = 12.25
	pointsOfFunctionPlot[1_225].Y = 3.5

	pointsOfFunctionPlot[1_226].X = 12.26
	pointsOfFunctionPlot[1_226].Y = 3.501

	pointsOfFunctionPlot[1_227].X = 12.27
	pointsOfFunctionPlot[1_227].Y = 3.502

	pointsOfFunctionPlot[1_228].X = 12.28
	pointsOfFunctionPlot[1_228].Y = 3.504

	pointsOfFunctionPlot[1_229].X = 12.29
	pointsOfFunctionPlot[1_229].Y = 3.505

	pointsOfFunctionPlot[1_230].X = 12.30
	pointsOfFunctionPlot[1_230].Y = 3.507

	pointsOfFunctionPlot[1_231].X = 12.31
	pointsOfFunctionPlot[1_231].Y = 3.508

	pointsOfFunctionPlot[1_232].X = 12.32
	pointsOfFunctionPlot[1_232].Y = 3.509

	pointsOfFunctionPlot[1_233].X = 12.33
	pointsOfFunctionPlot[1_233].Y = 3.511

	pointsOfFunctionPlot[1_234].X = 12.34
	pointsOfFunctionPlot[1_234].Y = 3.512

	pointsOfFunctionPlot[1_235].X = 12.35
	pointsOfFunctionPlot[1_235].Y = 3.514

	pointsOfFunctionPlot[1_236].X = 12.36
	pointsOfFunctionPlot[1_236].Y = 3.515

	pointsOfFunctionPlot[1_237].X = 12.37
	pointsOfFunctionPlot[1_237].Y = 3.517

	pointsOfFunctionPlot[1_238].X = 12.38
	pointsOfFunctionPlot[1_238].Y = 3.518

	pointsOfFunctionPlot[1_239].X = 12.39
	pointsOfFunctionPlot[1_239].Y = 3.519

	pointsOfFunctionPlot[1_240].X = 12.40
	pointsOfFunctionPlot[1_240].Y = 3.521

	pointsOfFunctionPlot[1_241].X = 12.41
	pointsOfFunctionPlot[1_241].Y = 3.522

	pointsOfFunctionPlot[1_242].X = 12.42
	pointsOfFunctionPlot[1_242].Y = 3.524

	pointsOfFunctionPlot[1_243].X = 12.43
	pointsOfFunctionPlot[1_243].Y = 3.525

	pointsOfFunctionPlot[1_244].X = 12.44
	pointsOfFunctionPlot[1_244].Y = 3.527

	pointsOfFunctionPlot[1_245].X = 12.45
	pointsOfFunctionPlot[1_245].Y = 3.528

	pointsOfFunctionPlot[1_246].X = 12.46
	pointsOfFunctionPlot[1_246].Y = 3.529

	pointsOfFunctionPlot[1_247].X = 12.47
	pointsOfFunctionPlot[1_247].Y = 3.531

	pointsOfFunctionPlot[1_248].X = 12.48
	pointsOfFunctionPlot[1_248].Y = 3.532

	pointsOfFunctionPlot[1_249].X = 12.49
	pointsOfFunctionPlot[1_249].Y = 3.534

	pointsOfFunctionPlot[1_250].X = 12.50
	pointsOfFunctionPlot[1_250].Y = 3.535

	pointsOfFunctionPlot[1_251].X = 12.51
	pointsOfFunctionPlot[1_251].Y = 3.536

	pointsOfFunctionPlot[1_252].X = 12.52
	pointsOfFunctionPlot[1_252].Y = 3.538

	pointsOfFunctionPlot[1_253].X = 12.53
	pointsOfFunctionPlot[1_253].Y = 3.539

	pointsOfFunctionPlot[1_254].X = 12.54
	pointsOfFunctionPlot[1_254].Y = 3.541

	pointsOfFunctionPlot[1_255].X = 12.55
	pointsOfFunctionPlot[1_255].Y = 3.542

	pointsOfFunctionPlot[1_256].X = 12.56
	pointsOfFunctionPlot[1_256].Y = 3.544

	pointsOfFunctionPlot[1_257].X = 12.57
	pointsOfFunctionPlot[1_257].Y = 3.545

	pointsOfFunctionPlot[1_258].X = 12.58
	pointsOfFunctionPlot[1_258].Y = 3.546

	pointsOfFunctionPlot[1_259].X = 12.59
	pointsOfFunctionPlot[1_259].Y = 3.548

	pointsOfFunctionPlot[1_260].X = 12.60
	pointsOfFunctionPlot[1_260].Y = 3.549

	pointsOfFunctionPlot[1_261].X = 12.61
	pointsOfFunctionPlot[1_261].Y = 3.551

	pointsOfFunctionPlot[1_262].X = 12.62
	pointsOfFunctionPlot[1_262].Y = 3.552

	pointsOfFunctionPlot[1_263].X = 12.63
	pointsOfFunctionPlot[1_263].Y = 3.553

	pointsOfFunctionPlot[1_264].X = 12.64
	pointsOfFunctionPlot[1_264].Y = 3.555

	pointsOfFunctionPlot[1_265].X = 12.65
	pointsOfFunctionPlot[1_265].Y = 3.556

	pointsOfFunctionPlot[1_266].X = 12.66
	pointsOfFunctionPlot[1_266].Y = 3.558

	pointsOfFunctionPlot[1_267].X = 12.67
	pointsOfFunctionPlot[1_267].Y = 3.559

	pointsOfFunctionPlot[1_268].X = 12.68
	pointsOfFunctionPlot[1_268].Y = 3.56

	pointsOfFunctionPlot[1_269].X = 12.69
	pointsOfFunctionPlot[1_269].Y = 3.562

	pointsOfFunctionPlot[1_270].X = 12.70
	pointsOfFunctionPlot[1_270].Y = 3.563

	pointsOfFunctionPlot[1_271].X = 12.71
	pointsOfFunctionPlot[1_271].Y = 3.565

	pointsOfFunctionPlot[1_272].X = 12.72
	pointsOfFunctionPlot[1_272].Y = 3.566

	pointsOfFunctionPlot[1_273].X = 12.73
	pointsOfFunctionPlot[1_273].Y = 3.567

	pointsOfFunctionPlot[1_274].X = 12.74
	pointsOfFunctionPlot[1_274].Y = 3.569

	pointsOfFunctionPlot[1_275].X = 12.75
	pointsOfFunctionPlot[1_275].Y = 3.57

	pointsOfFunctionPlot[1_276].X = 12.76
	pointsOfFunctionPlot[1_276].Y = 3.572

	pointsOfFunctionPlot[1_277].X = 12.77
	pointsOfFunctionPlot[1_277].Y = 3.573

	pointsOfFunctionPlot[1_278].X = 12.78
	pointsOfFunctionPlot[1_278].Y = 3.574

	pointsOfFunctionPlot[1_279].X = 12.79
	pointsOfFunctionPlot[1_279].Y = 3.576

	pointsOfFunctionPlot[1_280].X = 12.80
	pointsOfFunctionPlot[1_280].Y = 3.577

	pointsOfFunctionPlot[1_281].X = 12.81
	pointsOfFunctionPlot[1_281].Y = 3.579

	pointsOfFunctionPlot[1_282].X = 12.82
	pointsOfFunctionPlot[1_282].Y = 3.58

	pointsOfFunctionPlot[1_283].X = 12.83
	pointsOfFunctionPlot[1_283].Y = 3.581

	pointsOfFunctionPlot[1_284].X = 12.84
	pointsOfFunctionPlot[1_284].Y = 3.583

	pointsOfFunctionPlot[1_285].X = 12.85
	pointsOfFunctionPlot[1_285].Y = 3.584

	pointsOfFunctionPlot[1_286].X = 12.86
	pointsOfFunctionPlot[1_286].Y = 3.586

	pointsOfFunctionPlot[1_287].X = 12.87
	pointsOfFunctionPlot[1_287].Y = 3.587

	pointsOfFunctionPlot[1_288].X = 12.88
	pointsOfFunctionPlot[1_288].Y = 3.588

	pointsOfFunctionPlot[1_289].X = 12.89
	pointsOfFunctionPlot[1_289].Y = 3.59

	pointsOfFunctionPlot[1_290].X = 12.90
	pointsOfFunctionPlot[1_290].Y = 3.591

	pointsOfFunctionPlot[1_291].X = 12.91
	pointsOfFunctionPlot[1_291].Y = 3.593

	pointsOfFunctionPlot[1_292].X = 12.92
	pointsOfFunctionPlot[1_292].Y = 3.594

	pointsOfFunctionPlot[1_293].X = 12.93
	pointsOfFunctionPlot[1_293].Y = 3.595

	pointsOfFunctionPlot[1_294].X = 12.94
	pointsOfFunctionPlot[1_294].Y = 3.597

	pointsOfFunctionPlot[1_295].X = 12.95
	pointsOfFunctionPlot[1_295].Y = 3.598

	pointsOfFunctionPlot[1_296].X = 12.96
	pointsOfFunctionPlot[1_296].Y = 3.6

	pointsOfFunctionPlot[1_297].X = 12.97
	pointsOfFunctionPlot[1_297].Y = 3.601

	pointsOfFunctionPlot[1_298].X = 12.98
	pointsOfFunctionPlot[1_298].Y = 3.602

	pointsOfFunctionPlot[1_299].X = 12.99
	pointsOfFunctionPlot[1_299].Y = 3.604

	pointsOfFunctionPlot[1_300].X = 13.0
	pointsOfFunctionPlot[1_300].Y = 3.605

	pointsOfFunctionPlot[1_301].X = 13.01
	pointsOfFunctionPlot[1_301].Y = 3.606

	pointsOfFunctionPlot[1_302].X = 13.02
	pointsOfFunctionPlot[1_302].Y = 3.608

	pointsOfFunctionPlot[1_303].X = 13.03
	pointsOfFunctionPlot[1_303].Y = 3.609

	pointsOfFunctionPlot[1_304].X = 13.04
	pointsOfFunctionPlot[1_304].Y = 3.611

	pointsOfFunctionPlot[1_305].X = 13.05
	pointsOfFunctionPlot[1_305].Y = 3.612

	pointsOfFunctionPlot[1_306].X = 13.06
	pointsOfFunctionPlot[1_306].Y = 3.613

	pointsOfFunctionPlot[1_307].X = 13.07
	pointsOfFunctionPlot[1_307].Y = 3.615

	pointsOfFunctionPlot[1_308].X = 13.08
	pointsOfFunctionPlot[1_308].Y = 3.616

	pointsOfFunctionPlot[1_309].X = 13.09
	pointsOfFunctionPlot[1_309].Y = 3.618

	pointsOfFunctionPlot[1_310].X = 13.10
	pointsOfFunctionPlot[1_310].Y = 3.619

	pointsOfFunctionPlot[1_311].X = 13.11
	pointsOfFunctionPlot[1_311].Y = 3.62

	pointsOfFunctionPlot[1_312].X = 13.12
	pointsOfFunctionPlot[1_312].Y = 3.622

	pointsOfFunctionPlot[1_313].X = 13.13
	pointsOfFunctionPlot[1_313].Y = 3.623

	pointsOfFunctionPlot[1_314].X = 13.14
	pointsOfFunctionPlot[1_314].Y = 3.624

	pointsOfFunctionPlot[1_315].X = 13.15
	pointsOfFunctionPlot[1_315].Y = 3.626

	pointsOfFunctionPlot[1_316].X = 13.16
	pointsOfFunctionPlot[1_316].Y = 3.627

	pointsOfFunctionPlot[1_317].X = 13.17
	pointsOfFunctionPlot[1_317].Y = 3.629

	pointsOfFunctionPlot[1_318].X = 13.18
	pointsOfFunctionPlot[1_318].Y = 3.63

	pointsOfFunctionPlot[1_319].X = 13.19
	pointsOfFunctionPlot[1_319].Y = 3.631

	pointsOfFunctionPlot[1_320].X = 13.20
	pointsOfFunctionPlot[1_320].Y = 3.633

	pointsOfFunctionPlot[1_321].X = 13.21
	pointsOfFunctionPlot[1_321].Y = 3.634

	pointsOfFunctionPlot[1_322].X = 13.22
	pointsOfFunctionPlot[1_322].Y = 3.635

	pointsOfFunctionPlot[1_323].X = 13.23
	pointsOfFunctionPlot[1_323].Y = 3.637

	pointsOfFunctionPlot[1_324].X = 13.24
	pointsOfFunctionPlot[1_324].Y = 3.638

	pointsOfFunctionPlot[1_325].X = 13.25
	pointsOfFunctionPlot[1_325].Y = 3.64

	pointsOfFunctionPlot[1_326].X = 13.26
	pointsOfFunctionPlot[1_326].Y = 3.641

	pointsOfFunctionPlot[1_327].X = 13.27
	pointsOfFunctionPlot[1_327].Y = 3.642

	pointsOfFunctionPlot[1_328].X = 13.28
	pointsOfFunctionPlot[1_328].Y = 3.644

	pointsOfFunctionPlot[1_329].X = 13.29
	pointsOfFunctionPlot[1_329].Y = 3.645

	pointsOfFunctionPlot[1_330].X = 13.30
	pointsOfFunctionPlot[1_330].Y = 3.646

	pointsOfFunctionPlot[1_331].X = 13.31
	pointsOfFunctionPlot[1_331].Y = 3.648

	pointsOfFunctionPlot[1_332].X = 13.32
	pointsOfFunctionPlot[1_332].Y = 3.649

	pointsOfFunctionPlot[1_333].X = 13.33
	pointsOfFunctionPlot[1_333].Y = 3.651

	pointsOfFunctionPlot[1_334].X = 13.34
	pointsOfFunctionPlot[1_334].Y = 3.652

	pointsOfFunctionPlot[1_335].X = 13.35
	pointsOfFunctionPlot[1_335].Y = 3.653

	pointsOfFunctionPlot[1_336].X = 13.36
	pointsOfFunctionPlot[1_336].Y = 3.655

	pointsOfFunctionPlot[1_337].X = 13.37
	pointsOfFunctionPlot[1_337].Y = 3.656

	pointsOfFunctionPlot[1_338].X = 13.38
	pointsOfFunctionPlot[1_338].Y = 3.657

	pointsOfFunctionPlot[1_339].X = 13.39
	pointsOfFunctionPlot[1_339].Y = 3.659

	pointsOfFunctionPlot[1_340].X = 13.40
	pointsOfFunctionPlot[1_340].Y = 3.66

	pointsOfFunctionPlot[1_341].X = 13.41
	pointsOfFunctionPlot[1_341].Y = 3.661

	pointsOfFunctionPlot[1_342].X = 13.42
	pointsOfFunctionPlot[1_342].Y = 3.663

	pointsOfFunctionPlot[1_343].X = 13.43
	pointsOfFunctionPlot[1_343].Y = 3.664

	pointsOfFunctionPlot[1_344].X = 13.44
	pointsOfFunctionPlot[1_344].Y = 3.666

	pointsOfFunctionPlot[1_345].X = 13.45
	pointsOfFunctionPlot[1_345].Y = 3.667

	pointsOfFunctionPlot[1_346].X = 13.46
	pointsOfFunctionPlot[1_346].Y = 3.668

	pointsOfFunctionPlot[1_347].X = 13.47
	pointsOfFunctionPlot[1_347].Y = 3.67

	pointsOfFunctionPlot[1_348].X = 13.48
	pointsOfFunctionPlot[1_348].Y = 3.671

	pointsOfFunctionPlot[1_349].X = 13.49
	pointsOfFunctionPlot[1_349].Y = 3.672

	pointsOfFunctionPlot[1_350].X = 13.50
	pointsOfFunctionPlot[1_350].Y = 3.674

	pointsOfFunctionPlot[1_351].X = 13.51
	pointsOfFunctionPlot[1_351].Y = 3.675

	pointsOfFunctionPlot[1_352].X = 13.52
	pointsOfFunctionPlot[1_352].Y = 3.676

	pointsOfFunctionPlot[1_353].X = 13.53
	pointsOfFunctionPlot[1_353].Y = 3.678

	pointsOfFunctionPlot[1_354].X = 13.54
	pointsOfFunctionPlot[1_354].Y = 3.679

	pointsOfFunctionPlot[1_355].X = 13.55
	pointsOfFunctionPlot[1_355].Y = 3.681

	pointsOfFunctionPlot[1_356].X = 13.56
	pointsOfFunctionPlot[1_356].Y = 3.682

	pointsOfFunctionPlot[1_357].X = 13.57
	pointsOfFunctionPlot[1_357].Y = 3.683

	pointsOfFunctionPlot[1_358].X = 13.58
	pointsOfFunctionPlot[1_358].Y = 3.685

	pointsOfFunctionPlot[1_359].X = 13.59
	pointsOfFunctionPlot[1_359].Y = 3.686

	pointsOfFunctionPlot[1_360].X = 13.60
	pointsOfFunctionPlot[1_360].Y = 3.687

	pointsOfFunctionPlot[1_361].X = 13.61
	pointsOfFunctionPlot[1_361].Y = 3.689

	pointsOfFunctionPlot[1_362].X = 13.62
	pointsOfFunctionPlot[1_362].Y = 3.69

	pointsOfFunctionPlot[1_363].X = 13.63
	pointsOfFunctionPlot[1_363].Y = 3.691

	pointsOfFunctionPlot[1_364].X = 13.64
	pointsOfFunctionPlot[1_364].Y = 3.693

	pointsOfFunctionPlot[1_365].X = 13.65
	pointsOfFunctionPlot[1_365].Y = 3.694

	pointsOfFunctionPlot[1_366].X = 13.66
	pointsOfFunctionPlot[1_366].Y = 3.695

	pointsOfFunctionPlot[1_367].X = 13.67
	pointsOfFunctionPlot[1_367].Y = 3.697

	pointsOfFunctionPlot[1_368].X = 13.68
	pointsOfFunctionPlot[1_368].Y = 3.698

	pointsOfFunctionPlot[1_369].X = 13.69
	pointsOfFunctionPlot[1_369].Y = 3.699

	pointsOfFunctionPlot[1_370].X = 13.70
	pointsOfFunctionPlot[1_370].Y = 3.701

	pointsOfFunctionPlot[1_371].X = 13.71
	pointsOfFunctionPlot[1_371].Y = 3.702

	pointsOfFunctionPlot[1_372].X = 13.72
	pointsOfFunctionPlot[1_372].Y = 3.704

	pointsOfFunctionPlot[1_373].X = 13.73
	pointsOfFunctionPlot[1_373].Y = 3.705

	pointsOfFunctionPlot[1_374].X = 13.74
	pointsOfFunctionPlot[1_374].Y = 3.706

	pointsOfFunctionPlot[1_375].X = 13.75
	pointsOfFunctionPlot[1_375].Y = 3.708

	pointsOfFunctionPlot[1_376].X = 13.76
	pointsOfFunctionPlot[1_376].Y = 3.709

	pointsOfFunctionPlot[1_377].X = 13.77
	pointsOfFunctionPlot[1_377].Y = 3.71

	pointsOfFunctionPlot[1_378].X = 13.78
	pointsOfFunctionPlot[1_378].Y = 3.712

	pointsOfFunctionPlot[1_379].X = 13.79
	pointsOfFunctionPlot[1_379].Y = 3.713

	pointsOfFunctionPlot[1_380].X = 13.80
	pointsOfFunctionPlot[1_380].Y = 3.714

	pointsOfFunctionPlot[1_381].X = 13.81
	pointsOfFunctionPlot[1_381].Y = 3.716

	pointsOfFunctionPlot[1_382].X = 13.82
	pointsOfFunctionPlot[1_382].Y = 3.717

	pointsOfFunctionPlot[1_383].X = 13.83
	pointsOfFunctionPlot[1_383].Y = 3.718

	pointsOfFunctionPlot[1_384].X = 13.84
	pointsOfFunctionPlot[1_384].Y = 3.72

	pointsOfFunctionPlot[1_385].X = 13.85
	pointsOfFunctionPlot[1_385].Y = 3.721

	pointsOfFunctionPlot[1_386].X = 13.86
	pointsOfFunctionPlot[1_386].Y = 3.722

	pointsOfFunctionPlot[1_387].X = 13.87
	pointsOfFunctionPlot[1_387].Y = 3.724

	pointsOfFunctionPlot[1_388].X = 13.88
	pointsOfFunctionPlot[1_388].Y = 3.725

	pointsOfFunctionPlot[1_389].X = 13.89
	pointsOfFunctionPlot[1_389].Y = 3.726

	pointsOfFunctionPlot[1_390].X = 13.90
	pointsOfFunctionPlot[1_390].Y = 3.728

	pointsOfFunctionPlot[1_391].X = 13.91
	pointsOfFunctionPlot[1_391].Y = 3.729

	pointsOfFunctionPlot[1_392].X = 13.92
	pointsOfFunctionPlot[1_392].Y = 3.73

	pointsOfFunctionPlot[1_393].X = 13.93
	pointsOfFunctionPlot[1_393].Y = 3.732

	pointsOfFunctionPlot[1_394].X = 13.94
	pointsOfFunctionPlot[1_394].Y = 3.733

	pointsOfFunctionPlot[1_395].X = 13.95
	pointsOfFunctionPlot[1_395].Y = 3.734

	pointsOfFunctionPlot[1_396].X = 13.96
	pointsOfFunctionPlot[1_396].Y = 3.736

	pointsOfFunctionPlot[1_397].X = 13.97
	pointsOfFunctionPlot[1_397].Y = 3.737

	pointsOfFunctionPlot[1_398].X = 13.98
	pointsOfFunctionPlot[1_398].Y = 3.738

	pointsOfFunctionPlot[1_399].X = 13.99
	pointsOfFunctionPlot[1_399].Y = 3.74

	pointsOfFunctionPlot[1_400].X = 14.0
	pointsOfFunctionPlot[1_400].Y = 3.741

	pointsOfFunctionPlot[1_401].X = 14.01
	pointsOfFunctionPlot[1_401].Y = 3.742

	pointsOfFunctionPlot[1_402].X = 14.02
	pointsOfFunctionPlot[1_402].Y = 3.744

	pointsOfFunctionPlot[1_403].X = 14.03
	pointsOfFunctionPlot[1_403].Y = 3.745

	pointsOfFunctionPlot[1_404].X = 14.04
	pointsOfFunctionPlot[1_404].Y = 3.746

	pointsOfFunctionPlot[1_405].X = 14.05
	pointsOfFunctionPlot[1_405].Y = 3.748

	pointsOfFunctionPlot[1_406].X = 14.06
	pointsOfFunctionPlot[1_406].Y = 3.749

	pointsOfFunctionPlot[1_407].X = 14.07
	pointsOfFunctionPlot[1_407].Y = 3.75

	pointsOfFunctionPlot[1_408].X = 14.08
	pointsOfFunctionPlot[1_408].Y = 3.752

	pointsOfFunctionPlot[1_409].X = 14.09
	pointsOfFunctionPlot[1_409].Y = 3.753

	pointsOfFunctionPlot[1_410].X = 14.10
	pointsOfFunctionPlot[1_410].Y = 3.754

	pointsOfFunctionPlot[1_411].X = 14.11
	pointsOfFunctionPlot[1_411].Y = 3.756

	pointsOfFunctionPlot[1_412].X = 14.12
	pointsOfFunctionPlot[1_412].Y = 3.757

	pointsOfFunctionPlot[1_413].X = 14.13
	pointsOfFunctionPlot[1_413].Y = 3.758

	pointsOfFunctionPlot[1_414].X = 14.14
	pointsOfFunctionPlot[1_414].Y = 3.76

	pointsOfFunctionPlot[1_415].X = 14.15
	pointsOfFunctionPlot[1_415].Y = 3.761

	pointsOfFunctionPlot[1_416].X = 14.16
	pointsOfFunctionPlot[1_416].Y = 3.762

	pointsOfFunctionPlot[1_417].X = 14.17
	pointsOfFunctionPlot[1_417].Y = 3.764

	pointsOfFunctionPlot[1_418].X = 14.18
	pointsOfFunctionPlot[1_418].Y = 3.765

	pointsOfFunctionPlot[1_419].X = 14.19
	pointsOfFunctionPlot[1_419].Y = 3.766

	pointsOfFunctionPlot[1_420].X = 14.20
	pointsOfFunctionPlot[1_420].Y = 3.768

	pointsOfFunctionPlot[1_421].X = 14.21
	pointsOfFunctionPlot[1_421].Y = 3.769

	pointsOfFunctionPlot[1_422].X = 14.22
	pointsOfFunctionPlot[1_422].Y = 3.77

	pointsOfFunctionPlot[1_423].X = 14.23
	pointsOfFunctionPlot[1_423].Y = 3.772

	pointsOfFunctionPlot[1_424].X = 14.24
	pointsOfFunctionPlot[1_424].Y = 3.773

	pointsOfFunctionPlot[1_425].X = 14.25
	pointsOfFunctionPlot[1_425].Y = 3.774

	pointsOfFunctionPlot[1_426].X = 14.26
	pointsOfFunctionPlot[1_426].Y = 3.776

	pointsOfFunctionPlot[1_427].X = 14.27
	pointsOfFunctionPlot[1_427].Y = 3.777

	pointsOfFunctionPlot[1_428].X = 14.28
	pointsOfFunctionPlot[1_428].Y = 3.778

	pointsOfFunctionPlot[1_429].X = 14.29
	pointsOfFunctionPlot[1_429].Y = 3.78

	pointsOfFunctionPlot[1_430].X = 14.30
	pointsOfFunctionPlot[1_430].Y = 3.781

	pointsOfFunctionPlot[1_431].X = 14.31
	pointsOfFunctionPlot[1_431].Y = 3.782

	pointsOfFunctionPlot[1_432].X = 14.32
	pointsOfFunctionPlot[1_432].Y = 3.784

	pointsOfFunctionPlot[1_433].X = 14.33
	pointsOfFunctionPlot[1_433].Y = 3.785

	pointsOfFunctionPlot[1_434].X = 14.34
	pointsOfFunctionPlot[1_434].Y = 3.786

	pointsOfFunctionPlot[1_435].X = 14.35
	pointsOfFunctionPlot[1_435].Y = 3.788

	pointsOfFunctionPlot[1_436].X = 14.36
	pointsOfFunctionPlot[1_436].Y = 3.789

	pointsOfFunctionPlot[1_437].X = 14.37
	pointsOfFunctionPlot[1_437].Y = 3.79

	pointsOfFunctionPlot[1_438].X = 14.38
	pointsOfFunctionPlot[1_438].Y = 3.792

	pointsOfFunctionPlot[1_439].X = 14.39
	pointsOfFunctionPlot[1_439].Y = 3.793

	pointsOfFunctionPlot[1_440].X = 14.40
	pointsOfFunctionPlot[1_440].Y = 3.794

	pointsOfFunctionPlot[1_441].X = 14.41
	pointsOfFunctionPlot[1_441].Y = 3.796

	pointsOfFunctionPlot[1_442].X = 14.42
	pointsOfFunctionPlot[1_442].Y = 3.797

	pointsOfFunctionPlot[1_443].X = 14.43
	pointsOfFunctionPlot[1_443].Y = 3.798

	pointsOfFunctionPlot[1_444].X = 14.44
	pointsOfFunctionPlot[1_444].Y = 3.8

	pointsOfFunctionPlot[1_445].X = 14.45
	pointsOfFunctionPlot[1_445].Y = 3.801

	pointsOfFunctionPlot[1_446].X = 14.46
	pointsOfFunctionPlot[1_446].Y = 3.802

	pointsOfFunctionPlot[1_447].X = 14.47
	pointsOfFunctionPlot[1_447].Y = 3.803

	pointsOfFunctionPlot[1_448].X = 14.48
	pointsOfFunctionPlot[1_448].Y = 3.805

	pointsOfFunctionPlot[1_449].X = 14.49
	pointsOfFunctionPlot[1_449].Y = 3.806

	pointsOfFunctionPlot[1_450].X = 14.50
	pointsOfFunctionPlot[1_450].Y = 3.807

	pointsOfFunctionPlot[1_451].X = 14.51
	pointsOfFunctionPlot[1_451].Y = 3.809

	pointsOfFunctionPlot[1_452].X = 14.52
	pointsOfFunctionPlot[1_452].Y = 3.81

	pointsOfFunctionPlot[1_453].X = 14.53
	pointsOfFunctionPlot[1_453].Y = 3.811

	pointsOfFunctionPlot[1_454].X = 14.54
	pointsOfFunctionPlot[1_454].Y = 3.813

	pointsOfFunctionPlot[1_455].X = 14.55
	pointsOfFunctionPlot[1_455].Y = 3.814

	pointsOfFunctionPlot[1_456].X = 14.56
	pointsOfFunctionPlot[1_456].Y = 3.815

	pointsOfFunctionPlot[1_457].X = 14.57
	pointsOfFunctionPlot[1_457].Y = 3.817

	pointsOfFunctionPlot[1_458].X = 14.58
	pointsOfFunctionPlot[1_458].Y = 3.818

	pointsOfFunctionPlot[1_459].X = 14.59
	pointsOfFunctionPlot[1_459].Y = 3.819

	pointsOfFunctionPlot[1_460].X = 14.60
	pointsOfFunctionPlot[1_460].Y = 3.82

	pointsOfFunctionPlot[1_461].X = 14.61
	pointsOfFunctionPlot[1_461].Y = 3.822

	pointsOfFunctionPlot[1_462].X = 14.62
	pointsOfFunctionPlot[1_462].Y = 3.823

	pointsOfFunctionPlot[1_463].X = 14.63
	pointsOfFunctionPlot[1_463].Y = 3.824

	pointsOfFunctionPlot[1_464].X = 14.64
	pointsOfFunctionPlot[1_464].Y = 3.826

	pointsOfFunctionPlot[1_465].X = 14.65
	pointsOfFunctionPlot[1_465].Y = 3.827

	pointsOfFunctionPlot[1_466].X = 14.66
	pointsOfFunctionPlot[1_466].Y = 3.828

	pointsOfFunctionPlot[1_467].X = 14.67
	pointsOfFunctionPlot[1_467].Y = 3.83

	pointsOfFunctionPlot[1_468].X = 14.68
	pointsOfFunctionPlot[1_468].Y = 3.831

	pointsOfFunctionPlot[1_469].X = 14.69
	pointsOfFunctionPlot[1_469].Y = 3.832

	pointsOfFunctionPlot[1_470].X = 14.70
	pointsOfFunctionPlot[1_470].Y = 3.834

	pointsOfFunctionPlot[1_471].X = 14.71
	pointsOfFunctionPlot[1_471].Y = 3.835

	pointsOfFunctionPlot[1_472].X = 14.72
	pointsOfFunctionPlot[1_472].Y = 3.836

	pointsOfFunctionPlot[1_473].X = 14.738
	pointsOfFunctionPlot[1_473].Y = 3.837

	pointsOfFunctionPlot[1_474].X = 14.74
	pointsOfFunctionPlot[1_474].Y = 3.839

	pointsOfFunctionPlot[1_475].X = 14.75
	pointsOfFunctionPlot[1_475].Y = 3.84

	pointsOfFunctionPlot[1_476].X = 14.76
	pointsOfFunctionPlot[1_476].Y = 3.841

	pointsOfFunctionPlot[1_477].X = 14.77
	pointsOfFunctionPlot[1_477].Y = 3.843

	pointsOfFunctionPlot[1_478].X = 14.78
	pointsOfFunctionPlot[1_478].Y = 3.844

	pointsOfFunctionPlot[1_479].X = 14.79
	pointsOfFunctionPlot[1_479].Y = 3.845

	pointsOfFunctionPlot[1_480].X = 14.80
	pointsOfFunctionPlot[1_480].Y = 3.847

	pointsOfFunctionPlot[1_481].X = 14.81
	pointsOfFunctionPlot[1_481].Y = 3.848

	pointsOfFunctionPlot[1_482].X = 14.82
	pointsOfFunctionPlot[1_482].Y = 3.849

	pointsOfFunctionPlot[1_483].X = 14.83
	pointsOfFunctionPlot[1_483].Y = 3.85

	pointsOfFunctionPlot[1_484].X = 14.84
	pointsOfFunctionPlot[1_484].Y = 3.852

	pointsOfFunctionPlot[1_485].X = 14.85
	pointsOfFunctionPlot[1_485].Y = 3.853

	pointsOfFunctionPlot[1_486].X = 14.86
	pointsOfFunctionPlot[1_486].Y = 3.854

	pointsOfFunctionPlot[1_487].X = 14.87
	pointsOfFunctionPlot[1_487].Y = 3.856

	pointsOfFunctionPlot[1_488].X = 14.88
	pointsOfFunctionPlot[1_488].Y = 3.857

	pointsOfFunctionPlot[1_489].X = 14.89
	pointsOfFunctionPlot[1_489].Y = 3.858

	pointsOfFunctionPlot[1_490].X = 14.90
	pointsOfFunctionPlot[1_490].Y = 3.86

	pointsOfFunctionPlot[1_491].X = 14.91
	pointsOfFunctionPlot[1_491].Y = 3.861

	pointsOfFunctionPlot[1_492].X = 14.92
	pointsOfFunctionPlot[1_492].Y = 3.862

	pointsOfFunctionPlot[1_493].X = 14.93
	pointsOfFunctionPlot[1_493].Y = 3.863

	pointsOfFunctionPlot[1_494].X = 14.94
	pointsOfFunctionPlot[1_494].Y = 3.865

	pointsOfFunctionPlot[1_495].X = 14.95
	pointsOfFunctionPlot[1_495].Y = 3.866

	pointsOfFunctionPlot[1_496].X = 14.96
	pointsOfFunctionPlot[1_496].Y = 3.867

	pointsOfFunctionPlot[1_497].X = 14.97
	pointsOfFunctionPlot[1_497].Y = 3.869

	pointsOfFunctionPlot[1_498].X = 14.98
	pointsOfFunctionPlot[1_498].Y = 3.87

	pointsOfFunctionPlot[1_499].X = 14.99
	pointsOfFunctionPlot[1_499].Y = 3.871

	pointsOfFunctionPlot[1_500].X = 15.0
	pointsOfFunctionPlot[1_500].Y = 3.872

	pointsOfFunctionPlot[1_501].X = 15.01
	pointsOfFunctionPlot[1_501].Y = 3.874

	pointsOfFunctionPlot[1_502].X = 15.02
	pointsOfFunctionPlot[1_502].Y = 3.875

	pointsOfFunctionPlot[1_503].X = 15.03
	pointsOfFunctionPlot[1_503].Y = 3.876

	pointsOfFunctionPlot[1_504].X = 15.04
	pointsOfFunctionPlot[1_504].Y = 3.878

	pointsOfFunctionPlot[1_505].X = 15.05
	pointsOfFunctionPlot[1_505].Y = 3.879

	pointsOfFunctionPlot[1_506].X = 15.06
	pointsOfFunctionPlot[1_506].Y = 3.88

	pointsOfFunctionPlot[1_507].X = 15.07
	pointsOfFunctionPlot[1_507].Y = 3.882

	pointsOfFunctionPlot[1_508].X = 15.08
	pointsOfFunctionPlot[1_508].Y = 3.883

	pointsOfFunctionPlot[1_509].X = 15.09
	pointsOfFunctionPlot[1_509].Y = 3.884

	pointsOfFunctionPlot[1_510].X = 15.10
	pointsOfFunctionPlot[1_510].Y = 3.885

	pointsOfFunctionPlot[1_511].X = 15.11
	pointsOfFunctionPlot[1_511].Y = 3.887

	pointsOfFunctionPlot[1_512].X = 15.12
	pointsOfFunctionPlot[1_512].Y = 3.888

	pointsOfFunctionPlot[1_513].X = 15.13
	pointsOfFunctionPlot[1_513].Y = 3.889

	pointsOfFunctionPlot[1_514].X = 15.14
	pointsOfFunctionPlot[1_514].Y = 3.891

	pointsOfFunctionPlot[1_515].X = 15.15
	pointsOfFunctionPlot[1_515].Y = 3.892

	pointsOfFunctionPlot[1_516].X = 15.16
	pointsOfFunctionPlot[1_516].Y = 3.893

	pointsOfFunctionPlot[1_517].X = 15.17
	pointsOfFunctionPlot[1_517].Y = 3.894

	pointsOfFunctionPlot[1_518].X = 15.18
	pointsOfFunctionPlot[1_518].Y = 3.896

	pointsOfFunctionPlot[1_519].X = 15.19
	pointsOfFunctionPlot[1_519].Y = 3.897

	pointsOfFunctionPlot[1_520].X = 15.20
	pointsOfFunctionPlot[1_520].Y = 3.898

	pointsOfFunctionPlot[1_521].X = 15.21
	pointsOfFunctionPlot[1_521].Y = 3.9

	pointsOfFunctionPlot[1_522].X = 15.22
	pointsOfFunctionPlot[1_522].Y = 3.901

	pointsOfFunctionPlot[1_523].X = 15.23
	pointsOfFunctionPlot[1_523].Y = 3.902

	pointsOfFunctionPlot[1_524].X = 15.24
	pointsOfFunctionPlot[1_524].Y = 3.903

	pointsOfFunctionPlot[1_525].X = 15.25
	pointsOfFunctionPlot[1_525].Y = 3.905

	pointsOfFunctionPlot[1_526].X = 15.26
	pointsOfFunctionPlot[1_526].Y = 3.906

	pointsOfFunctionPlot[1_527].X = 15.27
	pointsOfFunctionPlot[1_527].Y = 3.907

	pointsOfFunctionPlot[1_528].X = 15.28
	pointsOfFunctionPlot[1_528].Y = 3.908

	pointsOfFunctionPlot[1_529].X = 15.29
	pointsOfFunctionPlot[1_529].Y = 3.91

	pointsOfFunctionPlot[1_530].X = 15.30
	pointsOfFunctionPlot[1_530].Y = 3.911

	pointsOfFunctionPlot[1_531].X = 15.31
	pointsOfFunctionPlot[1_531].Y = 3.912

	pointsOfFunctionPlot[1_532].X = 15.32
	pointsOfFunctionPlot[1_532].Y = 3.914

	pointsOfFunctionPlot[1_533].X = 15.33
	pointsOfFunctionPlot[1_533].Y = 3.915

	pointsOfFunctionPlot[1_534].X = 15.34
	pointsOfFunctionPlot[1_534].Y = 3.916

	pointsOfFunctionPlot[1_535].X = 15.35
	pointsOfFunctionPlot[1_535].Y = 3.917

	pointsOfFunctionPlot[1_536].X = 15.36
	pointsOfFunctionPlot[1_536].Y = 3.919

	pointsOfFunctionPlot[1_537].X = 15.37
	pointsOfFunctionPlot[1_537].Y = 3.92

	pointsOfFunctionPlot[1_538].X = 15.38
	pointsOfFunctionPlot[1_538].Y = 3.921

	pointsOfFunctionPlot[1_539].X = 15.39
	pointsOfFunctionPlot[1_539].Y = 3.923

	pointsOfFunctionPlot[1_540].X = 15.40
	pointsOfFunctionPlot[1_540].Y = 3.924

	pointsOfFunctionPlot[1_541].X = 15.41
	pointsOfFunctionPlot[1_541].Y = 3.925

	pointsOfFunctionPlot[1_542].X = 15.42
	pointsOfFunctionPlot[1_542].Y = 3.926

	pointsOfFunctionPlot[1_543].X = 15.43
	pointsOfFunctionPlot[1_543].Y = 3.928

	pointsOfFunctionPlot[1_544].X = 15.44
	pointsOfFunctionPlot[1_544].Y = 3.929

	pointsOfFunctionPlot[1_545].X = 15.45
	pointsOfFunctionPlot[1_545].Y = 3.93

	pointsOfFunctionPlot[1_546].X = 15.46
	pointsOfFunctionPlot[1_546].Y = 3.931

	pointsOfFunctionPlot[1_547].X = 15.47
	pointsOfFunctionPlot[1_547].Y = 3.933

	pointsOfFunctionPlot[1_548].X = 15.48
	pointsOfFunctionPlot[1_548].Y = 3.934

	pointsOfFunctionPlot[1_549].X = 15.49
	pointsOfFunctionPlot[1_549].Y = 3.935

	pointsOfFunctionPlot[1_550].X = 15.50
	pointsOfFunctionPlot[1_550].Y = 3.937

	pointsOfFunctionPlot[1_551].X = 15.51
	pointsOfFunctionPlot[1_551].Y = 3.938

	pointsOfFunctionPlot[1_552].X = 15.52
	pointsOfFunctionPlot[1_552].Y = 3.939

	pointsOfFunctionPlot[1_553].X = 15.53
	pointsOfFunctionPlot[1_553].Y = 3.94

	pointsOfFunctionPlot[1_554].X = 15.54
	pointsOfFunctionPlot[1_554].Y = 3.942

	pointsOfFunctionPlot[1_555].X = 15.55
	pointsOfFunctionPlot[1_555].Y = 3.943

	pointsOfFunctionPlot[1_556].X = 15.56
	pointsOfFunctionPlot[1_556].Y = 3.944

	pointsOfFunctionPlot[1_557].X = 15.57
	pointsOfFunctionPlot[1_557].Y = 3.945

	pointsOfFunctionPlot[1_558].X = 15.58
	pointsOfFunctionPlot[1_558].Y = 3.947

	pointsOfFunctionPlot[1_559].X = 15.59
	pointsOfFunctionPlot[1_559].Y = 3.948

	pointsOfFunctionPlot[1_560].X = 15.60
	pointsOfFunctionPlot[1_560].Y = 3.949

	pointsOfFunctionPlot[1_561].X = 15.61
	pointsOfFunctionPlot[1_561].Y = 3.95

	pointsOfFunctionPlot[1_562].X = 15.62
	pointsOfFunctionPlot[1_562].Y = 3.952

	pointsOfFunctionPlot[1_563].X = 15.63
	pointsOfFunctionPlot[1_563].Y = 3.953

	pointsOfFunctionPlot[1_564].X = 15.64
	pointsOfFunctionPlot[1_564].Y = 3.954

	pointsOfFunctionPlot[1_565].X = 15.65
	pointsOfFunctionPlot[1_565].Y = 3.956

	pointsOfFunctionPlot[1_566].X = 15.66
	pointsOfFunctionPlot[1_566].Y = 3.957

	pointsOfFunctionPlot[1_567].X = 15.67
	pointsOfFunctionPlot[1_567].Y = 3.958

	pointsOfFunctionPlot[1_568].X = 15.68
	pointsOfFunctionPlot[1_568].Y = 3.959

	pointsOfFunctionPlot[1_569].X = 15.69
	pointsOfFunctionPlot[1_569].Y = 3.961

	pointsOfFunctionPlot[1_570].X = 15.70
	pointsOfFunctionPlot[1_570].Y = 3.962

	pointsOfFunctionPlot[1_571].X = 15.71
	pointsOfFunctionPlot[1_571].Y = 3.963

	pointsOfFunctionPlot[1_572].X = 15.72
	pointsOfFunctionPlot[1_572].Y = 3.964

	pointsOfFunctionPlot[1_573].X = 15.73
	pointsOfFunctionPlot[1_573].Y = 3.966

	pointsOfFunctionPlot[1_574].X = 15.74
	pointsOfFunctionPlot[1_574].Y = 3.967

	pointsOfFunctionPlot[1_575].X = 15.75
	pointsOfFunctionPlot[1_575].Y = 3.968

	pointsOfFunctionPlot[1_576].X = 15.76
	pointsOfFunctionPlot[1_576].Y = 3.969

	pointsOfFunctionPlot[1_577].X = 15.77
	pointsOfFunctionPlot[1_577].Y = 3.971

	pointsOfFunctionPlot[1_578].X = 15.78
	pointsOfFunctionPlot[1_578].Y = 3.972

	pointsOfFunctionPlot[1_579].X = 15.79
	pointsOfFunctionPlot[1_579].Y = 3.973

	pointsOfFunctionPlot[1_580].X = 15.80
	pointsOfFunctionPlot[1_580].Y = 3.974

	pointsOfFunctionPlot[1_581].X = 15.81
	pointsOfFunctionPlot[1_581].Y = 3.976

	pointsOfFunctionPlot[1_582].X = 15.82
	pointsOfFunctionPlot[1_582].Y = 3.977

	pointsOfFunctionPlot[1_583].X = 15.83
	pointsOfFunctionPlot[1_583].Y = 3.978

	pointsOfFunctionPlot[1_584].X = 15.84
	pointsOfFunctionPlot[1_584].Y = 3.979

	pointsOfFunctionPlot[1_585].X = 15.85
	pointsOfFunctionPlot[1_585].Y = 3.981

	pointsOfFunctionPlot[1_586].X = 15.86
	pointsOfFunctionPlot[1_586].Y = 3.982

	pointsOfFunctionPlot[1_587].X = 15.87
	pointsOfFunctionPlot[1_587].Y = 3.983

	pointsOfFunctionPlot[1_588].X = 15.88
	pointsOfFunctionPlot[1_588].Y = 3.984

	pointsOfFunctionPlot[1_589].X = 15.89
	pointsOfFunctionPlot[1_589].Y = 3.986

	pointsOfFunctionPlot[1_590].X = 15.90
	pointsOfFunctionPlot[1_590].Y = 3.987

	pointsOfFunctionPlot[1_591].X = 15.91
	pointsOfFunctionPlot[1_591].Y = 3.988

	pointsOfFunctionPlot[1_592].X = 15.92
	pointsOfFunctionPlot[1_592].Y = 3.989

	pointsOfFunctionPlot[1_593].X = 15.93
	pointsOfFunctionPlot[1_593].Y = 3.991

	pointsOfFunctionPlot[1_594].X = 15.94
	pointsOfFunctionPlot[1_594].Y = 3.992

	pointsOfFunctionPlot[1_595].X = 15.95
	pointsOfFunctionPlot[1_595].Y = 3.993

	pointsOfFunctionPlot[1_596].X = 15.96
	pointsOfFunctionPlot[1_596].Y = 3.994

	pointsOfFunctionPlot[1_597].X = 15.97
	pointsOfFunctionPlot[1_597].Y = 3.996

	pointsOfFunctionPlot[1_598].X = 15.98
	pointsOfFunctionPlot[1_598].Y = 3.997

	pointsOfFunctionPlot[1_599].X = 15.99
	pointsOfFunctionPlot[1_599].Y = 3.998

	pointsOfFunctionPlot[1_600].X = 16.0
	pointsOfFunctionPlot[1_600].Y = 4.0

	pointsOfFunctionPlot[1_601].X = 16.01
	pointsOfFunctionPlot[1_601].Y = 4.001

	pointsOfFunctionPlot[1_602].X = 16.02
	pointsOfFunctionPlot[1_602].Y = 4.002

	pointsOfFunctionPlot[1_603].X = 16.03
	pointsOfFunctionPlot[1_603].Y = 4.003

	pointsOfFunctionPlot[1_604].X = 16.04
	pointsOfFunctionPlot[1_604].Y = 4.004

	pointsOfFunctionPlot[1_605].X = 16.05
	pointsOfFunctionPlot[1_605].Y = 4.006

	pointsOfFunctionPlot[1_606].X = 16.06
	pointsOfFunctionPlot[1_606].Y = 4.007

	pointsOfFunctionPlot[1_607].X = 16.07
	pointsOfFunctionPlot[1_607].Y = 4.008

	pointsOfFunctionPlot[1_608].X = 16.08
	pointsOfFunctionPlot[1_608].Y = 4.009

	pointsOfFunctionPlot[1_609].X = 16.09
	pointsOfFunctionPlot[1_609].Y = 4.011

	pointsOfFunctionPlot[1_610].X = 16.10
	pointsOfFunctionPlot[1_610].Y = 4.012

	pointsOfFunctionPlot[1_611].X = 16.11
	pointsOfFunctionPlot[1_611].Y = 4.013

	pointsOfFunctionPlot[1_612].X = 16.12
	pointsOfFunctionPlot[1_612].Y = 4.014

	pointsOfFunctionPlot[1_613].X = 16.13
	pointsOfFunctionPlot[1_613].Y = 4.016

	pointsOfFunctionPlot[1_614].X = 16.14
	pointsOfFunctionPlot[1_614].Y = 4.017

	pointsOfFunctionPlot[1_615].X = 16.15
	pointsOfFunctionPlot[1_615].Y = 4.018

	pointsOfFunctionPlot[1_616].X = 16.16
	pointsOfFunctionPlot[1_616].Y = 4.019

	pointsOfFunctionPlot[1_617].X = 16.17
	pointsOfFunctionPlot[1_617].Y = 4.021

	pointsOfFunctionPlot[1_618].X = 16.18
	pointsOfFunctionPlot[1_618].Y = 4.022

	pointsOfFunctionPlot[1_619].X = 16.19
	pointsOfFunctionPlot[1_619].Y = 4.023

	pointsOfFunctionPlot[1_620].X = 16.20
	pointsOfFunctionPlot[1_620].Y = 4.024

	pointsOfFunctionPlot[1_621].X = 16.21
	pointsOfFunctionPlot[1_621].Y = 4.026

	pointsOfFunctionPlot[1_622].X = 16.22
	pointsOfFunctionPlot[1_622].Y = 4.027

	pointsOfFunctionPlot[1_623].X = 16.23
	pointsOfFunctionPlot[1_623].Y = 4.028

	pointsOfFunctionPlot[1_624].X = 16.24
	pointsOfFunctionPlot[1_624].Y = 4.029

	pointsOfFunctionPlot[1_625].X = 16.25
	pointsOfFunctionPlot[1_625].Y = 4.031

	pointsOfFunctionPlot[1_626].X = 16.26
	pointsOfFunctionPlot[1_626].Y = 4.032

	pointsOfFunctionPlot[1_627].X = 16.27
	pointsOfFunctionPlot[1_627].Y = 4.033

	pointsOfFunctionPlot[1_628].X = 16.28
	pointsOfFunctionPlot[1_628].Y = 4.034

	pointsOfFunctionPlot[1_629].X = 16.29
	pointsOfFunctionPlot[1_629].Y = 4.036

	pointsOfFunctionPlot[1_630].X = 16.30
	pointsOfFunctionPlot[1_630].Y = 4.037

	pointsOfFunctionPlot[1_631].X = 16.31
	pointsOfFunctionPlot[1_631].Y = 4.038

	pointsOfFunctionPlot[1_632].X = 16.32
	pointsOfFunctionPlot[1_632].Y = 4.039

	pointsOfFunctionPlot[1_633].X = 16.33
	pointsOfFunctionPlot[1_633].Y = 4.041

	pointsOfFunctionPlot[1_634].X = 16.34
	pointsOfFunctionPlot[1_634].Y = 4.042

	pointsOfFunctionPlot[1_635].X = 16.35
	pointsOfFunctionPlot[1_635].Y = 4.043

	pointsOfFunctionPlot[1_636].X = 16.36
	pointsOfFunctionPlot[1_636].Y = 4.044

	pointsOfFunctionPlot[1_637].X = 16.37
	pointsOfFunctionPlot[1_637].Y = 4.045

	pointsOfFunctionPlot[1_638].X = 16.38
	pointsOfFunctionPlot[1_638].Y = 4.047

	pointsOfFunctionPlot[1_639].X = 16.39
	pointsOfFunctionPlot[1_639].Y = 4.048

	pointsOfFunctionPlot[1_640].X = 16.40
	pointsOfFunctionPlot[1_640].Y = 4.049

	pointsOfFunctionPlot[1_641].X = 16.41
	pointsOfFunctionPlot[1_641].Y = 4.05

	pointsOfFunctionPlot[1_642].X = 16.42
	pointsOfFunctionPlot[1_642].Y = 4.052

	pointsOfFunctionPlot[1_643].X = 16.43
	pointsOfFunctionPlot[1_643].Y = 4.053

	pointsOfFunctionPlot[1_644].X = 16.44
	pointsOfFunctionPlot[1_644].Y = 4.054

	pointsOfFunctionPlot[1_645].X = 16.45
	pointsOfFunctionPlot[1_645].Y = 4.055

	pointsOfFunctionPlot[1_646].X = 16.46
	pointsOfFunctionPlot[1_646].Y = 4.057

	pointsOfFunctionPlot[1_647].X = 16.47
	pointsOfFunctionPlot[1_647].Y = 4.058

	pointsOfFunctionPlot[1_648].X = 16.48
	pointsOfFunctionPlot[1_648].Y = 4.059

	pointsOfFunctionPlot[1_649].X = 16.49
	pointsOfFunctionPlot[1_649].Y = 4.06

	pointsOfFunctionPlot[1_650].X = 16.50
	pointsOfFunctionPlot[1_650].Y = 4.062

	pointsOfFunctionPlot[1_651].X = 16.51
	pointsOfFunctionPlot[1_651].Y = 4.063

	pointsOfFunctionPlot[1_652].X = 16.52
	pointsOfFunctionPlot[1_652].Y = 4.064

	pointsOfFunctionPlot[1_653].X = 16.53
	pointsOfFunctionPlot[1_653].Y = 4.065

	pointsOfFunctionPlot[1_654].X = 16.54
	pointsOfFunctionPlot[1_654].Y = 4.066

	pointsOfFunctionPlot[1_655].X = 16.55
	pointsOfFunctionPlot[1_655].Y = 4.068

	pointsOfFunctionPlot[1_656].X = 16.56
	pointsOfFunctionPlot[1_656].Y = 4.069

	pointsOfFunctionPlot[1_657].X = 16.57
	pointsOfFunctionPlot[1_657].Y = 4.07

	pointsOfFunctionPlot[1_658].X = 16.58
	pointsOfFunctionPlot[1_658].Y = 4.071

	pointsOfFunctionPlot[1_659].X = 16.59
	pointsOfFunctionPlot[1_659].Y = 4.073

	pointsOfFunctionPlot[1_660].X = 16.60
	pointsOfFunctionPlot[1_660].Y = 4.074

	pointsOfFunctionPlot[1_661].X = 16.61
	pointsOfFunctionPlot[1_661].Y = 4.075

	pointsOfFunctionPlot[1_662].X = 16.62
	pointsOfFunctionPlot[1_662].Y = 4.076

	pointsOfFunctionPlot[1_663].X = 16.63
	pointsOfFunctionPlot[1_663].Y = 4.077

	pointsOfFunctionPlot[1_664].X = 16.64
	pointsOfFunctionPlot[1_664].Y = 4.079

	pointsOfFunctionPlot[1_665].X = 16.65
	pointsOfFunctionPlot[1_665].Y = 4.08

	pointsOfFunctionPlot[1_666].X = 16.66
	pointsOfFunctionPlot[1_666].Y = 4.081

	pointsOfFunctionPlot[1_667].X = 16.67
	pointsOfFunctionPlot[1_667].Y = 4.082

	pointsOfFunctionPlot[1_668].X = 16.68
	pointsOfFunctionPlot[1_668].Y = 4.084

	pointsOfFunctionPlot[1_669].X = 16.69
	pointsOfFunctionPlot[1_669].Y = 4.085

	pointsOfFunctionPlot[1_670].X = 16.70
	pointsOfFunctionPlot[1_670].Y = 4.086

	pointsOfFunctionPlot[1_671].X = 16.71
	pointsOfFunctionPlot[1_671].Y = 4.087

	pointsOfFunctionPlot[1_672].X = 16.72
	pointsOfFunctionPlot[1_672].Y = 4.089

	pointsOfFunctionPlot[1_673].X = 16.73
	pointsOfFunctionPlot[1_673].Y = 4.09

	pointsOfFunctionPlot[1_674].X = 16.74
	pointsOfFunctionPlot[1_674].Y = 4.091

	pointsOfFunctionPlot[1_675].X = 16.75
	pointsOfFunctionPlot[1_675].Y = 4.092

	pointsOfFunctionPlot[1_676].X = 16.76
	pointsOfFunctionPlot[1_676].Y = 4.093

	pointsOfFunctionPlot[1_677].X = 16.77
	pointsOfFunctionPlot[1_677].Y = 4.095

	pointsOfFunctionPlot[1_678].X = 16.78
	pointsOfFunctionPlot[1_678].Y = 4.096

	pointsOfFunctionPlot[1_679].X = 16.79
	pointsOfFunctionPlot[1_679].Y = 4.097

	pointsOfFunctionPlot[1_680].X = 16.80
	pointsOfFunctionPlot[1_680].Y = 4.098

	pointsOfFunctionPlot[1_681].X = 16.81
	pointsOfFunctionPlot[1_681].Y = 4.1

	pointsOfFunctionPlot[1_682].X = 16.82
	pointsOfFunctionPlot[1_682].Y = 4.101

	pointsOfFunctionPlot[1_683].X = 16.83
	pointsOfFunctionPlot[1_683].Y = 4.102

	pointsOfFunctionPlot[1_684].X = 16.84
	pointsOfFunctionPlot[1_684].Y = 4.103

	pointsOfFunctionPlot[1_685].X = 16.85
	pointsOfFunctionPlot[1_685].Y = 4.104

	pointsOfFunctionPlot[1_686].X = 16.86
	pointsOfFunctionPlot[1_686].Y = 4.106

	pointsOfFunctionPlot[1_687].X = 16.87
	pointsOfFunctionPlot[1_687].Y = 4.107

	pointsOfFunctionPlot[1_688].X = 16.88
	pointsOfFunctionPlot[1_688].Y = 4.108

	pointsOfFunctionPlot[1_689].X = 16.89
	pointsOfFunctionPlot[1_689].Y = 4.109

	pointsOfFunctionPlot[1_690].X = 16.90
	pointsOfFunctionPlot[1_690].Y = 4.11

	pointsOfFunctionPlot[1_691].X = 16.91
	pointsOfFunctionPlot[1_691].Y = 4.112

	pointsOfFunctionPlot[1_692].X = 16.92
	pointsOfFunctionPlot[1_692].Y = 4.113

	pointsOfFunctionPlot[1_693].X = 16.93
	pointsOfFunctionPlot[1_693].Y = 4.114

	pointsOfFunctionPlot[1_694].X = 16.94
	pointsOfFunctionPlot[1_694].Y = 4.115

	pointsOfFunctionPlot[1_695].X = 16.95
	pointsOfFunctionPlot[1_695].Y = 4.117

	pointsOfFunctionPlot[1_696].X = 16.96
	pointsOfFunctionPlot[1_696].Y = 4.118

	pointsOfFunctionPlot[1_697].X = 16.97
	pointsOfFunctionPlot[1_697].Y = 4.119

	pointsOfFunctionPlot[1_698].X = 16.98
	pointsOfFunctionPlot[1_698].Y = 4.12

	pointsOfFunctionPlot[1_699].X = 16.99
	pointsOfFunctionPlot[1_699].Y = 4.121

	pointsOfFunctionPlot[1_700].X = 17.0
	pointsOfFunctionPlot[1_700].Y = 4.123

	pointsOfFunctionPlot[1_701].X = 17.01
	pointsOfFunctionPlot[1_701].Y = 4.124

	pointsOfFunctionPlot[1_702].X = 17.02
	pointsOfFunctionPlot[1_702].Y = 4.125

	pointsOfFunctionPlot[1_703].X = 17.03
	pointsOfFunctionPlot[1_703].Y = 4.126

	pointsOfFunctionPlot[1_704].X = 17.04
	pointsOfFunctionPlot[1_704].Y = 4.127

	pointsOfFunctionPlot[1_705].X = 17.05
	pointsOfFunctionPlot[1_705].Y = 4.129

	pointsOfFunctionPlot[1_706].X = 17.06
	pointsOfFunctionPlot[1_706].Y = 4.13

	pointsOfFunctionPlot[1_707].X = 17.07
	pointsOfFunctionPlot[1_707].Y = 4.131

	pointsOfFunctionPlot[1_708].X = 17.08
	pointsOfFunctionPlot[1_708].Y = 4.132

	pointsOfFunctionPlot[1_709].X = 17.09
	pointsOfFunctionPlot[1_709].Y = 4.134

	pointsOfFunctionPlot[1_710].X = 17.10
	pointsOfFunctionPlot[1_710].Y = 4.135

	pointsOfFunctionPlot[1_711].X = 17.11
	pointsOfFunctionPlot[1_711].Y = 4.136

	pointsOfFunctionPlot[1_712].X = 17.12
	pointsOfFunctionPlot[1_712].Y = 4.137

	pointsOfFunctionPlot[1_713].X = 17.13
	pointsOfFunctionPlot[1_713].Y = 4.138

	pointsOfFunctionPlot[1_714].X = 17.14
	pointsOfFunctionPlot[1_714].Y = 4.14

	pointsOfFunctionPlot[1_715].X = 17.15
	pointsOfFunctionPlot[1_715].Y = 4.141

	pointsOfFunctionPlot[1_716].X = 17.16
	pointsOfFunctionPlot[1_716].Y = 4.142

	pointsOfFunctionPlot[1_717].X = 17.17
	pointsOfFunctionPlot[1_717].Y = 4.143

	pointsOfFunctionPlot[1_718].X = 17.18
	pointsOfFunctionPlot[1_718].Y = 4.144

	pointsOfFunctionPlot[1_719].X = 17.19
	pointsOfFunctionPlot[1_719].Y = 4.146

	pointsOfFunctionPlot[1_720].X = 17.20
	pointsOfFunctionPlot[1_720].Y = 4.147

	pointsOfFunctionPlot[1_721].X = 17.21
	pointsOfFunctionPlot[1_721].Y = 4.148

	pointsOfFunctionPlot[1_722].X = 17.22
	pointsOfFunctionPlot[1_722].Y = 4.149

	pointsOfFunctionPlot[1_723].X = 17.23
	pointsOfFunctionPlot[1_723].Y = 4.15

	pointsOfFunctionPlot[1_724].X = 17.24
	pointsOfFunctionPlot[1_724].Y = 4.152

	pointsOfFunctionPlot[1_725].X = 17.25
	pointsOfFunctionPlot[1_725].Y = 4.153

	pointsOfFunctionPlot[1_726].X = 17.26
	pointsOfFunctionPlot[1_726].Y = 4.154

	pointsOfFunctionPlot[1_727].X = 17.27
	pointsOfFunctionPlot[1_727].Y = 4.155

	pointsOfFunctionPlot[1_728].X = 17.28
	pointsOfFunctionPlot[1_728].Y = 4.156

	pointsOfFunctionPlot[1_729].X = 17.29
	pointsOfFunctionPlot[1_729].Y = 4.158

	pointsOfFunctionPlot[1_730].X = 17.30
	pointsOfFunctionPlot[1_730].Y = 4.159

	pointsOfFunctionPlot[1_731].X = 17.31
	pointsOfFunctionPlot[1_731].Y = 4.16

	pointsOfFunctionPlot[1_732].X = 17.32
	pointsOfFunctionPlot[1_732].Y = 4.161

	pointsOfFunctionPlot[1_733].X = 17.33
	pointsOfFunctionPlot[1_733].Y = 4.162

	pointsOfFunctionPlot[1_734].X = 17.34
	pointsOfFunctionPlot[1_734].Y = 4.164

	pointsOfFunctionPlot[1_735].X = 17.35
	pointsOfFunctionPlot[1_735].Y = 4.165

	pointsOfFunctionPlot[1_736].X = 17.36
	pointsOfFunctionPlot[1_736].Y = 4.166

	pointsOfFunctionPlot[1_737].X = 17.37
	pointsOfFunctionPlot[1_737].Y = 4.167

	pointsOfFunctionPlot[1_738].X = 17.38
	pointsOfFunctionPlot[1_738].Y = 4.168

	pointsOfFunctionPlot[1_739].X = 17.39
	pointsOfFunctionPlot[1_739].Y = 4.17

	pointsOfFunctionPlot[1_740].X = 17.40
	pointsOfFunctionPlot[1_740].Y = 4.171

	pointsOfFunctionPlot[1_741].X = 17.41
	pointsOfFunctionPlot[1_741].Y = 4.172

	pointsOfFunctionPlot[1_742].X = 17.42
	pointsOfFunctionPlot[1_742].Y = 4.173

	pointsOfFunctionPlot[1_743].X = 17.43
	pointsOfFunctionPlot[1_743].Y = 4.174

	pointsOfFunctionPlot[1_744].X = 17.44
	pointsOfFunctionPlot[1_744].Y = 4.176

	pointsOfFunctionPlot[1_745].X = 17.45
	pointsOfFunctionPlot[1_745].Y = 4.177

	pointsOfFunctionPlot[1_746].X = 17.46
	pointsOfFunctionPlot[1_746].Y = 4.178

	pointsOfFunctionPlot[1_747].X = 17.47
	pointsOfFunctionPlot[1_747].Y = 4.179

	pointsOfFunctionPlot[1_748].X = 17.48
	pointsOfFunctionPlot[1_748].Y = 4.18

	pointsOfFunctionPlot[1_749].X = 17.49
	pointsOfFunctionPlot[1_749].Y = 4.182

	pointsOfFunctionPlot[1_750].X = 17.50
	pointsOfFunctionPlot[1_750].Y = 4.183

	pointsOfFunctionPlot[1_751].X = 17.51
	pointsOfFunctionPlot[1_751].Y = 4.184

	pointsOfFunctionPlot[1_752].X = 17.52
	pointsOfFunctionPlot[1_752].Y = 4.185

	pointsOfFunctionPlot[1_753].X = 17.53
	pointsOfFunctionPlot[1_753].Y = 4.186

	pointsOfFunctionPlot[1_754].X = 17.54
	pointsOfFunctionPlot[1_754].Y = 4.188

	pointsOfFunctionPlot[1_755].X = 17.55
	pointsOfFunctionPlot[1_755].Y = 4.189

	pointsOfFunctionPlot[1_756].X = 17.56
	pointsOfFunctionPlot[1_756].Y = 4.19

	pointsOfFunctionPlot[1_757].X = 17.57
	pointsOfFunctionPlot[1_757].Y = 4.191

	pointsOfFunctionPlot[1_758].X = 17.58
	pointsOfFunctionPlot[1_758].Y = 4.192

	pointsOfFunctionPlot[1_759].X = 17.59
	pointsOfFunctionPlot[1_759].Y = 4.194

	pointsOfFunctionPlot[1_760].X = 17.60
	pointsOfFunctionPlot[1_760].Y = 4.195

	pointsOfFunctionPlot[1_761].X = 17.61
	pointsOfFunctionPlot[1_761].Y = 4.196

	pointsOfFunctionPlot[1_762].X = 17.62
	pointsOfFunctionPlot[1_762].Y = 4.197

	pointsOfFunctionPlot[1_763].X = 17.63
	pointsOfFunctionPlot[1_763].Y = 4.198

	pointsOfFunctionPlot[1_764].X = 17.64
	pointsOfFunctionPlot[1_764].Y = 4.2

	pointsOfFunctionPlot[1_765].X = 17.65
	pointsOfFunctionPlot[1_765].Y = 4.201

	pointsOfFunctionPlot[1_766].X = 17.66
	pointsOfFunctionPlot[1_766].Y = 4.202

	pointsOfFunctionPlot[1_767].X = 17.67
	pointsOfFunctionPlot[1_767].Y = 4.203

	pointsOfFunctionPlot[1_768].X = 17.68
	pointsOfFunctionPlot[1_768].Y = 4.204

	pointsOfFunctionPlot[1_769].X = 17.69
	pointsOfFunctionPlot[1_769].Y = 4.205

	pointsOfFunctionPlot[1_770].X = 17.70
	pointsOfFunctionPlot[1_770].Y = 4.207

	pointsOfFunctionPlot[1_771].X = 17.71
	pointsOfFunctionPlot[1_771].Y = 4.208

	pointsOfFunctionPlot[1_772].X = 17.72
	pointsOfFunctionPlot[1_772].Y = 4.209

	pointsOfFunctionPlot[1_773].X = 17.73
	pointsOfFunctionPlot[1_773].Y = 4.21

	pointsOfFunctionPlot[1_774].X = 17.74
	pointsOfFunctionPlot[1_774].Y = 4.211

	pointsOfFunctionPlot[1_775].X = 17.75
	pointsOfFunctionPlot[1_775].Y = 4.213

	pointsOfFunctionPlot[1_776].X = 17.76
	pointsOfFunctionPlot[1_776].Y = 4.214

	pointsOfFunctionPlot[1_777].X = 17.77
	pointsOfFunctionPlot[1_777].Y = 4.215

	pointsOfFunctionPlot[1_778].X = 17.78
	pointsOfFunctionPlot[1_778].Y = 4.216

	pointsOfFunctionPlot[1_779].X = 17.79
	pointsOfFunctionPlot[1_779].Y = 4.217

	pointsOfFunctionPlot[1_780].X = 17.80
	pointsOfFunctionPlot[1_780].Y = 4.219

	pointsOfFunctionPlot[1_781].X = 17.81
	pointsOfFunctionPlot[1_781].Y = 4.22

	pointsOfFunctionPlot[1_782].X = 17.82
	pointsOfFunctionPlot[1_782].Y = 4.221

	pointsOfFunctionPlot[1_783].X = 17.83
	pointsOfFunctionPlot[1_783].Y = 4.222

	pointsOfFunctionPlot[1_784].X = 17.84
	pointsOfFunctionPlot[1_784].Y = 4.223

	pointsOfFunctionPlot[1_785].X = 17.85
	pointsOfFunctionPlot[1_785].Y = 4.224

	pointsOfFunctionPlot[1_786].X = 17.86
	pointsOfFunctionPlot[1_786].Y = 4.226

	pointsOfFunctionPlot[1_787].X = 17.87
	pointsOfFunctionPlot[1_787].Y = 4.227

	pointsOfFunctionPlot[1_788].X = 17.88
	pointsOfFunctionPlot[1_788].Y = 4.228

	pointsOfFunctionPlot[1_789].X = 17.89
	pointsOfFunctionPlot[1_789].Y = 4.229

	pointsOfFunctionPlot[1_790].X = 17.90
	pointsOfFunctionPlot[1_790].Y = 4.23

	pointsOfFunctionPlot[1_791].X = 17.91
	pointsOfFunctionPlot[1_791].Y = 4.232

	pointsOfFunctionPlot[1_792].X = 17.92
	pointsOfFunctionPlot[1_792].Y = 4.233

	pointsOfFunctionPlot[1_793].X = 17.93
	pointsOfFunctionPlot[1_793].Y = 4.234

	pointsOfFunctionPlot[1_794].X = 17.94
	pointsOfFunctionPlot[1_794].Y = 4.235

	pointsOfFunctionPlot[1_795].X = 17.95
	pointsOfFunctionPlot[1_795].Y = 4.236

	pointsOfFunctionPlot[1_796].X = 17.96
	pointsOfFunctionPlot[1_796].Y = 4.237

	pointsOfFunctionPlot[1_797].X = 17.97
	pointsOfFunctionPlot[1_797].Y = 4.239

	pointsOfFunctionPlot[1_798].X = 17.98
	pointsOfFunctionPlot[1_798].Y = 4.24

	pointsOfFunctionPlot[1_799].X = 17.99
	pointsOfFunctionPlot[1_799].Y = 4.241

	pointsOfFunctionPlot[1_800].X = 18.0
	pointsOfFunctionPlot[1_800].Y = 4.242

	pointsOfFunctionPlot[1_801].X = 18.01
	pointsOfFunctionPlot[1_801].Y = 4.243

	pointsOfFunctionPlot[1_802].X = 18.02
	pointsOfFunctionPlot[1_802].Y = 4.244

	pointsOfFunctionPlot[1_803].X = 18.03
	pointsOfFunctionPlot[1_803].Y = 4.246

	pointsOfFunctionPlot[1_804].X = 18.04
	pointsOfFunctionPlot[1_804].Y = 4.247

	pointsOfFunctionPlot[1_805].X = 18.05
	pointsOfFunctionPlot[1_805].Y = 4.248

	pointsOfFunctionPlot[1_806].X = 18.06
	pointsOfFunctionPlot[1_806].Y = 4.249

	pointsOfFunctionPlot[1_807].X = 18.07
	pointsOfFunctionPlot[1_807].Y = 4.25

	pointsOfFunctionPlot[1_808].X = 18.08
	pointsOfFunctionPlot[1_808].Y = 4.252

	pointsOfFunctionPlot[1_809].X = 18.09
	pointsOfFunctionPlot[1_809].Y = 4.253

	pointsOfFunctionPlot[1_810].X = 18.10
	pointsOfFunctionPlot[1_810].Y = 4.254

	pointsOfFunctionPlot[1_811].X = 18.11
	pointsOfFunctionPlot[1_811].Y = 4.255

	pointsOfFunctionPlot[1_812].X = 18.12
	pointsOfFunctionPlot[1_812].Y = 4.256

	pointsOfFunctionPlot[1_813].X = 18.13
	pointsOfFunctionPlot[1_813].Y = 4.257

	pointsOfFunctionPlot[1_814].X = 18.14
	pointsOfFunctionPlot[1_814].Y = 4.259

	pointsOfFunctionPlot[1_815].X = 18.15
	pointsOfFunctionPlot[1_815].Y = 4.26

	pointsOfFunctionPlot[1_816].X = 18.16
	pointsOfFunctionPlot[1_816].Y = 4.261

	pointsOfFunctionPlot[1_817].X = 18.17
	pointsOfFunctionPlot[1_817].Y = 4.262

	pointsOfFunctionPlot[1_818].X = 18.18
	pointsOfFunctionPlot[1_818].Y = 4.263

	pointsOfFunctionPlot[1_819].X = 18.19
	pointsOfFunctionPlot[1_819].Y = 4.264

	pointsOfFunctionPlot[1_820].X = 18.20
	pointsOfFunctionPlot[1_820].Y = 4.266

	pointsOfFunctionPlot[1_821].X = 18.21
	pointsOfFunctionPlot[1_821].Y = 4.267

	pointsOfFunctionPlot[1_822].X = 18.22
	pointsOfFunctionPlot[1_822].Y = 4.268

	pointsOfFunctionPlot[1_823].X = 18.23
	pointsOfFunctionPlot[1_823].Y = 4.269

	pointsOfFunctionPlot[1_824].X = 18.24
	pointsOfFunctionPlot[1_824].Y = 4.27

	pointsOfFunctionPlot[1_825].X = 18.25
	pointsOfFunctionPlot[1_825].Y = 4.272

	pointsOfFunctionPlot[1_826].X = 18.26
	pointsOfFunctionPlot[1_826].Y = 4.273

	pointsOfFunctionPlot[1_827].X = 18.27
	pointsOfFunctionPlot[1_827].Y = 4.274

	pointsOfFunctionPlot[1_828].X = 18.28
	pointsOfFunctionPlot[1_828].Y = 4.275

	pointsOfFunctionPlot[1_829].X = 18.29
	pointsOfFunctionPlot[1_829].Y = 4.276

	pointsOfFunctionPlot[1_830].X = 18.30
	pointsOfFunctionPlot[1_830].Y = 4.277

	pointsOfFunctionPlot[1_831].X = 18.31
	pointsOfFunctionPlot[1_831].Y = 4.279

	pointsOfFunctionPlot[1_832].X = 18.32
	pointsOfFunctionPlot[1_832].Y = 4.28

	pointsOfFunctionPlot[1_833].X = 18.33
	pointsOfFunctionPlot[1_833].Y = 4.281

	pointsOfFunctionPlot[1_834].X = 18.34
	pointsOfFunctionPlot[1_834].Y = 4.282

	pointsOfFunctionPlot[1_835].X = 18.35
	pointsOfFunctionPlot[1_835].Y = 4.283

	pointsOfFunctionPlot[1_836].X = 18.36
	pointsOfFunctionPlot[1_836].Y = 4.284

	pointsOfFunctionPlot[1_837].X = 18.37
	pointsOfFunctionPlot[1_837].Y = 4.286

	pointsOfFunctionPlot[1_838].X = 18.38
	pointsOfFunctionPlot[1_838].Y = 4.287

	pointsOfFunctionPlot[1_839].X = 18.39
	pointsOfFunctionPlot[1_839].Y = 4.288

	pointsOfFunctionPlot[1_840].X = 18.40
	pointsOfFunctionPlot[1_840].Y = 4.289

	pointsOfFunctionPlot[1_841].X = 18.41
	pointsOfFunctionPlot[1_841].Y = 4.29

	pointsOfFunctionPlot[1_842].X = 18.42
	pointsOfFunctionPlot[1_842].Y = 4.291

	pointsOfFunctionPlot[1_843].X = 18.43
	pointsOfFunctionPlot[1_843].Y = 4.293

	pointsOfFunctionPlot[1_844].X = 18.44
	pointsOfFunctionPlot[1_844].Y = 4.294

	pointsOfFunctionPlot[1_845].X = 18.45
	pointsOfFunctionPlot[1_845].Y = 4.295

	pointsOfFunctionPlot[1_846].X = 18.46
	pointsOfFunctionPlot[1_846].Y = 4.296

	pointsOfFunctionPlot[1_847].X = 18.47
	pointsOfFunctionPlot[1_847].Y = 4.297

	pointsOfFunctionPlot[1_848].X = 18.48
	pointsOfFunctionPlot[1_848].Y = 4.298

	pointsOfFunctionPlot[1_849].X = 18.49
	pointsOfFunctionPlot[1_849].Y = 4.3

	pointsOfFunctionPlot[1_850].X = 18.50
	pointsOfFunctionPlot[1_850].Y = 4.301

	pointsOfFunctionPlot[1_851].X = 18.51
	pointsOfFunctionPlot[1_851].Y = 4.302

	pointsOfFunctionPlot[1_852].X = 18.52
	pointsOfFunctionPlot[1_852].Y = 4.303

	pointsOfFunctionPlot[1_853].X = 18.53
	pointsOfFunctionPlot[1_853].Y = 4.304

	pointsOfFunctionPlot[1_854].X = 18.54
	pointsOfFunctionPlot[1_854].Y = 4.305

	pointsOfFunctionPlot[1_855].X = 18.55
	pointsOfFunctionPlot[1_855].Y = 4.306

	pointsOfFunctionPlot[1_856].X = 18.56
	pointsOfFunctionPlot[1_856].Y = 4.308

	pointsOfFunctionPlot[1_857].X = 18.57
	pointsOfFunctionPlot[1_857].Y = 4.309

	pointsOfFunctionPlot[1_858].X = 18.58
	pointsOfFunctionPlot[1_858].Y = 4.31

	pointsOfFunctionPlot[1_859].X = 18.59
	pointsOfFunctionPlot[1_859].Y = 4.311

	pointsOfFunctionPlot[1_860].X = 18.60
	pointsOfFunctionPlot[1_860].Y = 4.312

	pointsOfFunctionPlot[1_861].X = 18.61
	pointsOfFunctionPlot[1_861].Y = 4.313

	pointsOfFunctionPlot[1_862].X = 18.62
	pointsOfFunctionPlot[1_862].Y = 4.315

	pointsOfFunctionPlot[1_863].X = 18.63
	pointsOfFunctionPlot[1_863].Y = 4.316

	pointsOfFunctionPlot[1_864].X = 18.64
	pointsOfFunctionPlot[1_864].Y = 4.317

	pointsOfFunctionPlot[1_865].X = 18.65
	pointsOfFunctionPlot[1_865].Y = 4.318

	pointsOfFunctionPlot[1_866].X = 18.66
	pointsOfFunctionPlot[1_866].Y = 4.319

	pointsOfFunctionPlot[1_867].X = 18.67
	pointsOfFunctionPlot[1_867].Y = 4.32

	pointsOfFunctionPlot[1_868].X = 18.68
	pointsOfFunctionPlot[1_868].Y = 4.322

	pointsOfFunctionPlot[1_869].X = 18.69
	pointsOfFunctionPlot[1_869].Y = 4.323

	pointsOfFunctionPlot[1_870].X = 18.70
	pointsOfFunctionPlot[1_870].Y = 4.324

	pointsOfFunctionPlot[1_871].X = 18.71
	pointsOfFunctionPlot[1_871].Y = 4.325

	pointsOfFunctionPlot[1_872].X = 18.72
	pointsOfFunctionPlot[1_872].Y = 4.326

	pointsOfFunctionPlot[1_873].X = 18.73
	pointsOfFunctionPlot[1_873].Y = 4.327

	pointsOfFunctionPlot[1_874].X = 18.74
	pointsOfFunctionPlot[1_874].Y = 4.328

	pointsOfFunctionPlot[1_875].X = 18.75
	pointsOfFunctionPlot[1_875].Y = 4.33

	pointsOfFunctionPlot[1_876].X = 18.76
	pointsOfFunctionPlot[1_876].Y = 4.331

	pointsOfFunctionPlot[1_877].X = 18.77
	pointsOfFunctionPlot[1_877].Y = 4.332

	pointsOfFunctionPlot[1_878].X = 18.78
	pointsOfFunctionPlot[1_878].Y = 4.333

	pointsOfFunctionPlot[1_879].X = 18.79
	pointsOfFunctionPlot[1_879].Y = 4.334

	pointsOfFunctionPlot[1_880].X = 18.80
	pointsOfFunctionPlot[1_880].Y = 4.335

	pointsOfFunctionPlot[1_881].X = 18.81
	pointsOfFunctionPlot[1_881].Y = 4.337

	pointsOfFunctionPlot[1_882].X = 18.82
	pointsOfFunctionPlot[1_882].Y = 4.338

	pointsOfFunctionPlot[1_883].X = 18.83
	pointsOfFunctionPlot[1_883].Y = 4.339

	pointsOfFunctionPlot[1_884].X = 18.84
	pointsOfFunctionPlot[1_884].Y = 4.34

	pointsOfFunctionPlot[1_885].X = 18.85
	pointsOfFunctionPlot[1_885].Y = 4.341

	pointsOfFunctionPlot[1_886].X = 18.86
	pointsOfFunctionPlot[1_886].Y = 4.342

	pointsOfFunctionPlot[1_887].X = 18.87
	pointsOfFunctionPlot[1_887].Y = 4.343

	pointsOfFunctionPlot[1_888].X = 18.88
	pointsOfFunctionPlot[1_888].Y = 4.345

	pointsOfFunctionPlot[1_889].X = 18.89
	pointsOfFunctionPlot[1_889].Y = 4.346

	pointsOfFunctionPlot[1_890].X = 18.90
	pointsOfFunctionPlot[1_890].Y = 4.347

	pointsOfFunctionPlot[1_891].X = 18.91
	pointsOfFunctionPlot[1_891].Y = 4.348

	pointsOfFunctionPlot[1_892].X = 18.92
	pointsOfFunctionPlot[1_892].Y = 4.349

	pointsOfFunctionPlot[1_893].X = 18.93
	pointsOfFunctionPlot[1_893].Y = 4.35

	pointsOfFunctionPlot[1_894].X = 18.94
	pointsOfFunctionPlot[1_894].Y = 4.352

	pointsOfFunctionPlot[1_895].X = 18.95
	pointsOfFunctionPlot[1_895].Y = 4.353

	pointsOfFunctionPlot[1_896].X = 18.96
	pointsOfFunctionPlot[1_896].Y = 4.354

	pointsOfFunctionPlot[1_897].X = 18.97
	pointsOfFunctionPlot[1_897].Y = 4.355

	pointsOfFunctionPlot[1_898].X = 18.98
	pointsOfFunctionPlot[1_898].Y = 4.356

	pointsOfFunctionPlot[1_899].X = 18.99
	pointsOfFunctionPlot[1_899].Y = 4.357

	pointsOfFunctionPlot[1_900].X = 19.0
	pointsOfFunctionPlot[1_900].Y = 4.358

	pointsOfFunctionPlot[1_901].X = 19.01
	pointsOfFunctionPlot[1_901].Y = 4.36

	pointsOfFunctionPlot[1_902].X = 19.02
	pointsOfFunctionPlot[1_902].Y = 4.361

	pointsOfFunctionPlot[1_903].X = 19.03
	pointsOfFunctionPlot[1_903].Y = 4.362

	pointsOfFunctionPlot[1_904].X = 19.04
	pointsOfFunctionPlot[1_904].Y = 4.363

	pointsOfFunctionPlot[1_905].X = 19.05
	pointsOfFunctionPlot[1_905].Y = 4.364

	pointsOfFunctionPlot[1_906].X = 19.06
	pointsOfFunctionPlot[1_906].Y = 4.365

	pointsOfFunctionPlot[1_907].X = 19.07
	pointsOfFunctionPlot[1_907].Y = 4.366

	pointsOfFunctionPlot[1_908].X = 19.08
	pointsOfFunctionPlot[1_908].Y = 4.368

	pointsOfFunctionPlot[1_909].X = 19.09
	pointsOfFunctionPlot[1_909].Y = 4.369

	pointsOfFunctionPlot[1_910].X = 19.10
	pointsOfFunctionPlot[1_910].Y = 4.37

	pointsOfFunctionPlot[1_911].X = 19.11
	pointsOfFunctionPlot[1_911].Y = 4.371

	pointsOfFunctionPlot[1_912].X = 19.12
	pointsOfFunctionPlot[1_912].Y = 4.372

	pointsOfFunctionPlot[1_913].X = 19.13
	pointsOfFunctionPlot[1_913].Y = 4.373

	pointsOfFunctionPlot[1_914].X = 19.14
	pointsOfFunctionPlot[1_914].Y = 4.374

	pointsOfFunctionPlot[1_915].X = 19.15
	pointsOfFunctionPlot[1_915].Y = 4.376

	pointsOfFunctionPlot[1_916].X = 19.16
	pointsOfFunctionPlot[1_916].Y = 4.377

	pointsOfFunctionPlot[1_917].X = 19.17
	pointsOfFunctionPlot[1_917].Y = 4.378

	pointsOfFunctionPlot[1_918].X = 19.18
	pointsOfFunctionPlot[1_918].Y = 4.379

	pointsOfFunctionPlot[1_919].X = 19.19
	pointsOfFunctionPlot[1_919].Y = 4.38

	pointsOfFunctionPlot[1_920].X = 19.20
	pointsOfFunctionPlot[1_920].Y = 4.381

	pointsOfFunctionPlot[1_921].X = 19.21
	pointsOfFunctionPlot[1_921].Y = 4.382

	pointsOfFunctionPlot[1_922].X = 19.22
	pointsOfFunctionPlot[1_922].Y = 4.384

	pointsOfFunctionPlot[1_923].X = 19.23
	pointsOfFunctionPlot[1_923].Y = 4.385

	pointsOfFunctionPlot[1_924].X = 19.24
	pointsOfFunctionPlot[1_924].Y = 4.386

	pointsOfFunctionPlot[1_925].X = 19.25
	pointsOfFunctionPlot[1_925].Y = 4.387

	pointsOfFunctionPlot[1_926].X = 19.26
	pointsOfFunctionPlot[1_926].Y = 4.388

	pointsOfFunctionPlot[1_927].X = 19.27
	pointsOfFunctionPlot[1_927].Y = 4.389

	pointsOfFunctionPlot[1_928].X = 19.28
	pointsOfFunctionPlot[1_928].Y = 4.39

	pointsOfFunctionPlot[1_929].X = 19.29
	pointsOfFunctionPlot[1_929].Y = 4.392

	pointsOfFunctionPlot[1_930].X = 19.30
	pointsOfFunctionPlot[1_930].Y = 4.393

	pointsOfFunctionPlot[1_931].X = 19.31
	pointsOfFunctionPlot[1_931].Y = 4.394

	pointsOfFunctionPlot[1_932].X = 19.32
	pointsOfFunctionPlot[1_932].Y = 4.395

	pointsOfFunctionPlot[1_933].X = 19.33
	pointsOfFunctionPlot[1_933].Y = 4.396

	pointsOfFunctionPlot[1_934].X = 19.34
	pointsOfFunctionPlot[1_934].Y = 4.397

	pointsOfFunctionPlot[1_935].X = 19.35
	pointsOfFunctionPlot[1_935].Y = 4.398

	pointsOfFunctionPlot[1_936].X = 19.36
	pointsOfFunctionPlot[1_936].Y = 4.4

	pointsOfFunctionPlot[1_937].X = 19.37
	pointsOfFunctionPlot[1_937].Y = 4.401

	pointsOfFunctionPlot[1_938].X = 19.38
	pointsOfFunctionPlot[1_938].Y = 4.402

	pointsOfFunctionPlot[1_939].X = 19.39
	pointsOfFunctionPlot[1_939].Y = 4.403

	pointsOfFunctionPlot[1_940].X = 19.40
	pointsOfFunctionPlot[1_940].Y = 4.404

	pointsOfFunctionPlot[1_941].X = 19.41
	pointsOfFunctionPlot[1_941].Y = 4.405

	pointsOfFunctionPlot[1_942].X = 19.42
	pointsOfFunctionPlot[1_942].Y = 4.406

	pointsOfFunctionPlot[1_943].X = 19.43
	pointsOfFunctionPlot[1_943].Y = 4.407

	pointsOfFunctionPlot[1_944].X = 19.44
	pointsOfFunctionPlot[1_944].Y = 4.409

	pointsOfFunctionPlot[1_945].X = 19.45
	pointsOfFunctionPlot[1_945].Y = 4.41

	pointsOfFunctionPlot[1_946].X = 19.46
	pointsOfFunctionPlot[1_946].Y = 4.411

	pointsOfFunctionPlot[1_947].X = 19.47
	pointsOfFunctionPlot[1_947].Y = 4.412

	pointsOfFunctionPlot[1_948].X = 19.48
	pointsOfFunctionPlot[1_948].Y = 4.413

	pointsOfFunctionPlot[1_949].X = 19.49
	pointsOfFunctionPlot[1_949].Y = 4.414

	pointsOfFunctionPlot[1_950].X = 19.50
	pointsOfFunctionPlot[1_950].Y = 4.415

	pointsOfFunctionPlot[1_951].X = 19.51
	pointsOfFunctionPlot[1_951].Y = 4.417

	pointsOfFunctionPlot[1_952].X = 19.52
	pointsOfFunctionPlot[1_952].Y = 4.418

	pointsOfFunctionPlot[1_953].X = 19.53
	pointsOfFunctionPlot[1_953].Y = 4.419

	pointsOfFunctionPlot[1_954].X = 19.54
	pointsOfFunctionPlot[1_954].Y = 4.42

	pointsOfFunctionPlot[1_955].X = 19.55
	pointsOfFunctionPlot[1_955].Y = 4.421

	pointsOfFunctionPlot[1_956].X = 19.56
	pointsOfFunctionPlot[1_956].Y = 4.422

	pointsOfFunctionPlot[1_957].X = 19.57
	pointsOfFunctionPlot[1_957].Y = 4.423

	pointsOfFunctionPlot[1_958].X = 19.58
	pointsOfFunctionPlot[1_958].Y = 4.424

	pointsOfFunctionPlot[1_959].X = 19.59
	pointsOfFunctionPlot[1_959].Y = 4.426

	pointsOfFunctionPlot[1_960].X = 19.60
	pointsOfFunctionPlot[1_960].Y = 4.427

	pointsOfFunctionPlot[1_961].X = 19.61
	pointsOfFunctionPlot[1_961].Y = 4.428

	pointsOfFunctionPlot[1_962].X = 19.62
	pointsOfFunctionPlot[1_962].Y = 4.429

	pointsOfFunctionPlot[1_963].X = 19.63
	pointsOfFunctionPlot[1_963].Y = 4.43

	pointsOfFunctionPlot[1_964].X = 19.64
	pointsOfFunctionPlot[1_964].Y = 4.431

	pointsOfFunctionPlot[1_965].X = 19.65
	pointsOfFunctionPlot[1_965].Y = 4.432

	pointsOfFunctionPlot[1_966].X = 19.66
	pointsOfFunctionPlot[1_966].Y = 4.433

	pointsOfFunctionPlot[1_967].X = 19.67
	pointsOfFunctionPlot[1_967].Y = 4.435

	pointsOfFunctionPlot[1_968].X = 19.68
	pointsOfFunctionPlot[1_968].Y = 4.436

	pointsOfFunctionPlot[1_969].X = 19.69
	pointsOfFunctionPlot[1_969].Y = 4.437

	pointsOfFunctionPlot[1_970].X = 19.70
	pointsOfFunctionPlot[1_970].Y = 4.438

	pointsOfFunctionPlot[1_971].X = 19.71
	pointsOfFunctionPlot[1_971].Y = 4.439

	pointsOfFunctionPlot[1_972].X = 19.72
	pointsOfFunctionPlot[1_972].Y = 4.44

	pointsOfFunctionPlot[1_973].X = 19.73
	pointsOfFunctionPlot[1_973].Y = 4.441

	pointsOfFunctionPlot[1_974].X = 19.74
	pointsOfFunctionPlot[1_974].Y = 4.442

	pointsOfFunctionPlot[1_975].X = 19.75
	pointsOfFunctionPlot[1_975].Y = 4.444

	pointsOfFunctionPlot[1_976].X = 19.76
	pointsOfFunctionPlot[1_976].Y = 4.445

	pointsOfFunctionPlot[1_977].X = 19.77
	pointsOfFunctionPlot[1_977].Y = 4.446

	pointsOfFunctionPlot[1_978].X = 19.78
	pointsOfFunctionPlot[1_978].Y = 4.447

	pointsOfFunctionPlot[1_979].X = 19.79
	pointsOfFunctionPlot[1_979].Y = 4.448

	pointsOfFunctionPlot[1_980].X = 19.80
	pointsOfFunctionPlot[1_980].Y = 4.449

	pointsOfFunctionPlot[1_981].X = 19.81
	pointsOfFunctionPlot[1_981].Y = 4.45

	pointsOfFunctionPlot[1_982].X = 19.82
	pointsOfFunctionPlot[1_982].Y = 4.451

	pointsOfFunctionPlot[1_983].X = 19.83
	pointsOfFunctionPlot[1_983].Y = 4.453

	pointsOfFunctionPlot[1_984].X = 19.84
	pointsOfFunctionPlot[1_984].Y = 4.454

	pointsOfFunctionPlot[1_985].X = 19.85
	pointsOfFunctionPlot[1_985].Y = 4.455

	pointsOfFunctionPlot[1_986].X = 19.86
	pointsOfFunctionPlot[1_986].Y = 4.456

	pointsOfFunctionPlot[1_987].X = 19.87
	pointsOfFunctionPlot[1_987].Y = 4.457

	pointsOfFunctionPlot[1_988].X = 19.88
	pointsOfFunctionPlot[1_988].Y = 4.458

	pointsOfFunctionPlot[1_989].X = 19.89
	pointsOfFunctionPlot[1_989].Y = 4.459

	pointsOfFunctionPlot[1_990].X = 19.90
	pointsOfFunctionPlot[1_990].Y = 4.46

	pointsOfFunctionPlot[1_991].X = 19.91
	pointsOfFunctionPlot[1_991].Y = 4.462

	pointsOfFunctionPlot[1_992].X = 19.92
	pointsOfFunctionPlot[1_992].Y = 4.463

	pointsOfFunctionPlot[1_993].X = 19.93
	pointsOfFunctionPlot[1_993].Y = 4.464

	pointsOfFunctionPlot[1_994].X = 19.94
	pointsOfFunctionPlot[1_994].Y = 4.465

	pointsOfFunctionPlot[1_995].X = 19.95
	pointsOfFunctionPlot[1_995].Y = 4.466

	pointsOfFunctionPlot[1_996].X = 19.96
	pointsOfFunctionPlot[1_996].Y = 4.467

	pointsOfFunctionPlot[1_997].X = 19.97
	pointsOfFunctionPlot[1_997].Y = 4.468

	pointsOfFunctionPlot[1_998].X = 19.98
	pointsOfFunctionPlot[1_998].Y = 4.469

	pointsOfFunctionPlot[1_999].X = 19.99
	pointsOfFunctionPlot[1_999].Y = 4.471

	pointsOfFunctionPlot[2_000].X = 20.0
	pointsOfFunctionPlot[2_000].Y = 4.472










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function sqrt(x)"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("sqrt(x)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"sqrt-function-plot-02.png"); err != nil {

		panic(err)
	}
}
