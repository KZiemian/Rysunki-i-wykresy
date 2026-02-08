package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of polynomial f(x) = x^3 - x^2 + 2x + 1.

	pointsOfPolynomialPlot := make(plotter.XYs, 1_001)

	pointsOfPolynomialPlot[0].X = 0.0
	pointsOfPolynomialPlot[0].Y = 1.0

	pointsOfPolynomialPlot[1].X = 0.01
	pointsOfPolynomialPlot[1].Y = 1.019

	pointsOfPolynomialPlot[2].X = 0.02
	pointsOfPolynomialPlot[2].Y = 1.039

	pointsOfPolynomialPlot[3].X = 0.03
	pointsOfPolynomialPlot[3].Y = 1.059

	pointsOfPolynomialPlot[4].X = 0.04
	pointsOfPolynomialPlot[4].Y = 1.078

	pointsOfPolynomialPlot[5].X = 0.05
	pointsOfPolynomialPlot[5].Y = 1.097

	pointsOfPolynomialPlot[6].X = 0.06
	pointsOfPolynomialPlot[6].Y = 1.116

	pointsOfPolynomialPlot[7].X = 0.07
	pointsOfPolynomialPlot[7].Y = 1.135

	pointsOfPolynomialPlot[8].X = 0.08
	pointsOfPolynomialPlot[8].Y = 1.154

	pointsOfPolynomialPlot[9].X = 0.09
	pointsOfPolynomialPlot[9].Y = 1.172

	pointsOfPolynomialPlot[10].X = 0.10
	pointsOfPolynomialPlot[10].Y = 1.191

	pointsOfPolynomialPlot[11].X = 0.11
	pointsOfPolynomialPlot[11].Y = 1.209

	pointsOfPolynomialPlot[12].X = 0.12
	pointsOfPolynomialPlot[12].Y = 1.227

	pointsOfPolynomialPlot[13].X = 0.13
	pointsOfPolynomialPlot[13].Y = 1.245

	pointsOfPolynomialPlot[14].X = 0.14
	pointsOfPolynomialPlot[14].Y = 1.263

	pointsOfPolynomialPlot[15].X = 0.15
	pointsOfPolynomialPlot[15].Y = 1.28

	pointsOfPolynomialPlot[16].X = 0.16
	pointsOfPolynomialPlot[16].Y = 1.298

	pointsOfPolynomialPlot[17].X = 0.17
	pointsOfPolynomialPlot[17].Y = 1.316

	pointsOfPolynomialPlot[18].X = 0.18
	pointsOfPolynomialPlot[18].Y = 1.333

	pointsOfPolynomialPlot[19].X = 0.19
	pointsOfPolynomialPlot[19].Y = 1.351

	pointsOfPolynomialPlot[20].X = 0.20
	pointsOfPolynomialPlot[20].Y = 1.368

	pointsOfPolynomialPlot[21].X = 0.21
	pointsOfPolynomialPlot[21].Y = 1.385

	pointsOfPolynomialPlot[22].X = 0.22
	pointsOfPolynomialPlot[22].Y = 1.402

	pointsOfPolynomialPlot[23].X = 0.23
	pointsOfPolynomialPlot[23].Y = 1.419

	pointsOfPolynomialPlot[24].X = 0.24
	pointsOfPolynomialPlot[24].Y = 1.436

	pointsOfPolynomialPlot[25].X = 0.25
	pointsOfPolynomialPlot[25].Y = 1.453

	pointsOfPolynomialPlot[26].X = 0.26
	pointsOfPolynomialPlot[26].Y = 1.47

	pointsOfPolynomialPlot[27].X = 0.27
	pointsOfPolynomialPlot[27].Y = 1.487

	pointsOfPolynomialPlot[28].X = 0.28
	pointsOfPolynomialPlot[28].Y = 1.504

	pointsOfPolynomialPlot[29].X = 0.29
	pointsOfPolynomialPlot[29].Y = 1.520

	pointsOfPolynomialPlot[30].X = 0.30
	pointsOfPolynomialPlot[30].Y = 1.537

	pointsOfPolynomialPlot[31].X = 0.31
	pointsOfPolynomialPlot[31].Y = 1.554

	pointsOfPolynomialPlot[32].X = 0.32
	pointsOfPolynomialPlot[32].Y = 1.570

	pointsOfPolynomialPlot[33].X = 0.33
	pointsOfPolynomialPlot[33].Y = 1.587

	pointsOfPolynomialPlot[34].X = 0.34
	pointsOfPolynomialPlot[34].Y = 1.604

	pointsOfPolynomialPlot[35].X = 0.35
	pointsOfPolynomialPlot[35].Y = 1.62

	pointsOfPolynomialPlot[36].X = 0.36
	pointsOfPolynomialPlot[36].Y = 1.636

	pointsOfPolynomialPlot[37].X = 0.37
	pointsOfPolynomialPlot[37].Y = 1.654

	pointsOfPolynomialPlot[38].X = 0.38
	pointsOfPolynomialPlot[38].Y = 1.67

	pointsOfPolynomialPlot[39].X = 0.39
	pointsOfPolynomialPlot[39].Y = 1.687

	pointsOfPolynomialPlot[40].X = 0.40
	pointsOfPolynomialPlot[40].Y = 1.704

	pointsOfPolynomialPlot[41].X = 0.41
	pointsOfPolynomialPlot[41].Y = 1.721

	pointsOfPolynomialPlot[42].X = 0.42
	pointsOfPolynomialPlot[42].Y = 1.738

	pointsOfPolynomialPlot[43].X = 0.43
	pointsOfPolynomialPlot[43].Y = 1.755

	pointsOfPolynomialPlot[44].X = 0.44
	pointsOfPolynomialPlot[44].Y = 1.772

	pointsOfPolynomialPlot[45].X = 0.45
	pointsOfPolynomialPlot[45].Y = 1.788

	pointsOfPolynomialPlot[46].X = 0.46
	pointsOfPolynomialPlot[46].Y = 1.806

	pointsOfPolynomialPlot[47].X = 0.47
	pointsOfPolynomialPlot[47].Y = 1.823

	pointsOfPolynomialPlot[48].X = 0.48
	pointsOfPolynomialPlot[48].Y = 1.840

	pointsOfPolynomialPlot[49].X = 0.49
	pointsOfPolynomialPlot[49].Y = 1.858

	pointsOfPolynomialPlot[50].X = 0.50
	pointsOfPolynomialPlot[50].Y = 1.875

	pointsOfPolynomialPlot[51].X = 0.51
	pointsOfPolynomialPlot[51].Y = 1.893

	pointsOfPolynomialPlot[52].X = 0.52
	pointsOfPolynomialPlot[52].Y = 1.91

	pointsOfPolynomialPlot[53].X = 0.53
	pointsOfPolynomialPlot[53].Y = 1.928

	pointsOfPolynomialPlot[54].X = 0.54
	pointsOfPolynomialPlot[54].Y = 1.946

	pointsOfPolynomialPlot[55].X = 0.55
	pointsOfPolynomialPlot[55].Y = 1.963

	pointsOfPolynomialPlot[56].X = 0.56
	pointsOfPolynomialPlot[56].Y = 1.982

	pointsOfPolynomialPlot[57].X = 0.57
	pointsOfPolynomialPlot[57].Y = 2.0

	pointsOfPolynomialPlot[58].X = 0.58
	pointsOfPolynomialPlot[58].Y = 2.019

	pointsOfPolynomialPlot[59].X = 0.59
	pointsOfPolynomialPlot[59].Y = 2.037

	pointsOfPolynomialPlot[60].X = 0.60
	pointsOfPolynomialPlot[60].Y = 2.056

	pointsOfPolynomialPlot[61].X = 0.61
	pointsOfPolynomialPlot[61].Y = 2.075

	pointsOfPolynomialPlot[62].X = 0.62
	pointsOfPolynomialPlot[62].Y = 2.094

	pointsOfPolynomialPlot[63].X = 0.63
	pointsOfPolynomialPlot[63].Y = 2.113

	pointsOfPolynomialPlot[64].X = 0.64
	pointsOfPolynomialPlot[64].Y = 2.133

	pointsOfPolynomialPlot[65].X = 0.65
	pointsOfPolynomialPlot[65].Y = 2.152

	pointsOfPolynomialPlot[66].X = 0.66
	pointsOfPolynomialPlot[66].Y = 2.171

	pointsOfPolynomialPlot[67].X = 0.67
	pointsOfPolynomialPlot[67].Y = 2.191

	pointsOfPolynomialPlot[68].X = 0.68
	pointsOfPolynomialPlot[68].Y = 2.212

	pointsOfPolynomialPlot[69].X = 0.69
	pointsOfPolynomialPlot[69].Y = 2.232

	pointsOfPolynomialPlot[70].X = 0.70
	pointsOfPolynomialPlot[70].Y = 2.253

	pointsOfPolynomialPlot[71].X = 0.71
	pointsOfPolynomialPlot[71].Y = 2.273

	pointsOfPolynomialPlot[72].X = 0.72
	pointsOfPolynomialPlot[72].Y = 2.294

	pointsOfPolynomialPlot[73].X = 0.73
	pointsOfPolynomialPlot[73].Y = 2.316

	pointsOfPolynomialPlot[74].X = 0.74
	pointsOfPolynomialPlot[74].Y = 2.337

	pointsOfPolynomialPlot[75].X = 0.75
	pointsOfPolynomialPlot[75].Y = 2.359

	pointsOfPolynomialPlot[76].X = 0.76
	pointsOfPolynomialPlot[76].Y = 2.381

	pointsOfPolynomialPlot[77].X = 0.77
	pointsOfPolynomialPlot[77].Y = 2.403

	pointsOfPolynomialPlot[78].X = 0.78
	pointsOfPolynomialPlot[78].Y = 2.426

	pointsOfPolynomialPlot[79].X = 0.79
	pointsOfPolynomialPlot[79].Y = 2.448

	pointsOfPolynomialPlot[80].X = 0.80
	pointsOfPolynomialPlot[80].Y = 2.472

	pointsOfPolynomialPlot[81].X = 0.81
	pointsOfPolynomialPlot[81].Y = 2.495

	pointsOfPolynomialPlot[82].X = 0.82
	pointsOfPolynomialPlot[82].Y = 2.519

	pointsOfPolynomialPlot[83].X = 0.83
	pointsOfPolynomialPlot[83].Y = 2.542

	pointsOfPolynomialPlot[84].X = 0.84
	pointsOfPolynomialPlot[84].Y = 2.567

	pointsOfPolynomialPlot[85].X = 0.85
	pointsOfPolynomialPlot[85].Y = 2.591

	pointsOfPolynomialPlot[86].X = 0.86
	pointsOfPolynomialPlot[86].Y = 2.616

	pointsOfPolynomialPlot[87].X = 0.87
	pointsOfPolynomialPlot[87].Y = 2.641

	pointsOfPolynomialPlot[88].X = 0.88
	pointsOfPolynomialPlot[88].Y = 2.667

	pointsOfPolynomialPlot[89].X = 0.89
	pointsOfPolynomialPlot[89].Y = 2.692

	pointsOfPolynomialPlot[90].X = 0.90
	pointsOfPolynomialPlot[90].Y = 2.719

	pointsOfPolynomialPlot[91].X = 0.91
	pointsOfPolynomialPlot[91].Y = 2.745

	pointsOfPolynomialPlot[92].X = 0.92
	pointsOfPolynomialPlot[92].Y = 2.772

	pointsOfPolynomialPlot[93].X = 0.93
	pointsOfPolynomialPlot[93].Y = 2.799

	pointsOfPolynomialPlot[94].X = 0.94
	pointsOfPolynomialPlot[94].Y = 2.827

	pointsOfPolynomialPlot[95].X = 0.95
	pointsOfPolynomialPlot[95].Y = 2.854

	pointsOfPolynomialPlot[96].X = 0.96
	pointsOfPolynomialPlot[96].Y = 2.883

	pointsOfPolynomialPlot[97].X = 0.97
	pointsOfPolynomialPlot[97].Y = 2.911

	pointsOfPolynomialPlot[98].X = 0.98
	pointsOfPolynomialPlot[98].Y = 2.94

	pointsOfPolynomialPlot[99].X = 0.99
	pointsOfPolynomialPlot[99].Y = 2.97

	pointsOfPolynomialPlot[100].X = 1.0
	pointsOfPolynomialPlot[100].Y = 3.0

	pointsOfPolynomialPlot[101].X = 1.01
	pointsOfPolynomialPlot[101].Y = 3.03

	pointsOfPolynomialPlot[102].X = 1.02
	pointsOfPolynomialPlot[102].Y = 3.06

	pointsOfPolynomialPlot[103].X = 1.03
	pointsOfPolynomialPlot[103].Y = 3.091

	pointsOfPolynomialPlot[104].X = 1.04
	pointsOfPolynomialPlot[104].Y = 3.123

	pointsOfPolynomialPlot[105].X = 1.05
	pointsOfPolynomialPlot[105].Y = 3.155

	pointsOfPolynomialPlot[106].X = 1.06
	pointsOfPolynomialPlot[106].Y = 3.187

	pointsOfPolynomialPlot[107].X = 1.07
	pointsOfPolynomialPlot[107].Y = 3.22

	pointsOfPolynomialPlot[108].X = 1.08
	pointsOfPolynomialPlot[108].Y = 3.253

	pointsOfPolynomialPlot[109].X = 1.09
	pointsOfPolynomialPlot[109].Y = 3.286

	pointsOfPolynomialPlot[110].X = 1.10
	pointsOfPolynomialPlot[110].Y = 3.321

	pointsOfPolynomialPlot[111].X = 1.11
	pointsOfPolynomialPlot[111].Y = 3.355

	pointsOfPolynomialPlot[112].X = 1.12
	pointsOfPolynomialPlot[112].Y = 3.39

	pointsOfPolynomialPlot[113].X = 1.13
	pointsOfPolynomialPlot[113].Y = 3.426

	pointsOfPolynomialPlot[114].X = 1.14
	pointsOfPolynomialPlot[114].Y = 3.461

	pointsOfPolynomialPlot[115].X = 1.15
	pointsOfPolynomialPlot[115].Y = 3.498

	pointsOfPolynomialPlot[116].X = 1.16
	pointsOfPolynomialPlot[116].Y = 3.535

	pointsOfPolynomialPlot[117].X = 1.17
	pointsOfPolynomialPlot[117].Y = 3.572

	pointsOfPolynomialPlot[118].X = 1.18
	pointsOfPolynomialPlot[118].Y = 3.61

	pointsOfPolynomialPlot[119].X = 1.19
	pointsOfPolynomialPlot[119].Y = 3.649

	pointsOfPolynomialPlot[120].X = 1.20
	pointsOfPolynomialPlot[120].Y = 3.688

	pointsOfPolynomialPlot[121].X = 1.21
	pointsOfPolynomialPlot[121].Y = 3.727

	pointsOfPolynomialPlot[122].X = 1.22
	pointsOfPolynomialPlot[122].Y = 3.767

	pointsOfPolynomialPlot[123].X = 1.23
	pointsOfPolynomialPlot[123].Y = 3.808

	pointsOfPolynomialPlot[124].X = 1.24
	pointsOfPolynomialPlot[124].Y = 3.849

	pointsOfPolynomialPlot[125].X = 1.25
	pointsOfPolynomialPlot[125].Y = 3.89

	pointsOfPolynomialPlot[126].X = 1.26
	pointsOfPolynomialPlot[126].Y = 3.932

	pointsOfPolynomialPlot[127].X = 1.27
	pointsOfPolynomialPlot[127].Y = 3.975

	pointsOfPolynomialPlot[128].X = 1.28
	pointsOfPolynomialPlot[128].Y = 4.018

	pointsOfPolynomialPlot[129].X = 1.29
	pointsOfPolynomialPlot[129].Y = 4.062

	pointsOfPolynomialPlot[130].X = 1.30
	pointsOfPolynomialPlot[130].Y = 4.107

	pointsOfPolynomialPlot[131].X = 1.31
	pointsOfPolynomialPlot[131].Y = 4.107

	pointsOfPolynomialPlot[132].X = 1.32
	pointsOfPolynomialPlot[132].Y = 4.197

	pointsOfPolynomialPlot[133].X = 1.33
	pointsOfPolynomialPlot[133].Y = 4.243

	pointsOfPolynomialPlot[134].X = 1.34
	pointsOfPolynomialPlot[134].Y = 4.29

	pointsOfPolynomialPlot[135].X = 1.35
	pointsOfPolynomialPlot[135].Y = 4.337

	pointsOfPolynomialPlot[136].X = 1.36
	pointsOfPolynomialPlot[136].Y = 4.385

	pointsOfPolynomialPlot[137].X = 1.37
	pointsOfPolynomialPlot[137].Y = 4.434

	pointsOfPolynomialPlot[138].X = 1.38
	pointsOfPolynomialPlot[138].Y = 4.483

	pointsOfPolynomialPlot[139].X = 1.39
	pointsOfPolynomialPlot[139].Y = 4.533

	pointsOfPolynomialPlot[140].X = 1.40
	pointsOfPolynomialPlot[140].Y = 4.584

	pointsOfPolynomialPlot[141].X = 1.41
	pointsOfPolynomialPlot[141].Y = 4.635

	pointsOfPolynomialPlot[142].X = 1.42
	pointsOfPolynomialPlot[142].Y = 4.686

	pointsOfPolynomialPlot[143].X = 1.43
	pointsOfPolynomialPlot[143].Y = 4.739

	pointsOfPolynomialPlot[144].X = 1.44
	pointsOfPolynomialPlot[144].Y = 4.792

	pointsOfPolynomialPlot[145].X = 1.45
	pointsOfPolynomialPlot[145].Y = 4.846

	pointsOfPolynomialPlot[146].X = 1.46
	pointsOfPolynomialPlot[146].Y = 4.9

	pointsOfPolynomialPlot[147].X = 1.47
	pointsOfPolynomialPlot[147].Y = 4.955

	pointsOfPolynomialPlot[146].X = 1.48
	pointsOfPolynomialPlot[146].Y = 5.011

	pointsOfPolynomialPlot[146].X = 1.49
	pointsOfPolynomialPlot[146].Y = 5.067

	pointsOfPolynomialPlot[150].X = 1.50
	pointsOfPolynomialPlot[150].Y = 5.125

	pointsOfPolynomialPlot[151].X = 1.51
	pointsOfPolynomialPlot[151].Y = 5.182

	pointsOfPolynomialPlot[152].X = 1.52
	pointsOfPolynomialPlot[152].Y = 5.241

	pointsOfPolynomialPlot[153].X = 1.53
	pointsOfPolynomialPlot[153].Y = 5.3

	pointsOfPolynomialPlot[154].X = 1.5
	pointsOfPolynomialPlot[154].Y = 5.36

	pointsOfPolynomialPlot[155].X = 1.55
	pointsOfPolynomialPlot[155].Y = 5.421

	pointsOfPolynomialPlot[156].X = 1.56
	pointsOfPolynomialPlot[156].Y = 5.482

	pointsOfPolynomialPlot[157].X = 1.57
	pointsOfPolynomialPlot[157].Y = 5.545

	pointsOfPolynomialPlot[158].X = 1.58
	pointsOfPolynomialPlot[158].Y = 5.607

	pointsOfPolynomialPlot[159].X = 1.59
	pointsOfPolynomialPlot[159].Y = 5.671

	pointsOfPolynomialPlot[160].X = 1.60
	pointsOfPolynomialPlot[160].Y = 5.736

	pointsOfPolynomialPlot[161].X = 1.61
	pointsOfPolynomialPlot[161].Y = 5.801

	pointsOfPolynomialPlot[162].X = 1.62
	pointsOfPolynomialPlot[162].Y = 5.867

	pointsOfPolynomialPlot[163].X = 1.63
	pointsOfPolynomialPlot[163].Y = 5.933

	pointsOfPolynomialPlot[164].X = 1.64
	pointsOfPolynomialPlot[164].Y = 6.001

	pointsOfPolynomialPlot[165].X = 1.65
	pointsOfPolynomialPlot[165].Y = 6.069

	pointsOfPolynomialPlot[166].X = 1.66
	pointsOfPolynomialPlot[166].Y = 6.138

	pointsOfPolynomialPlot[167].X = 1.67
	pointsOfPolynomialPlot[167].Y = 6.208

	pointsOfPolynomialPlot[168].X = 1.68
	pointsOfPolynomialPlot[168].Y = 6.279

	pointsOfPolynomialPlot[169].X = 1.69
	pointsOfPolynomialPlot[169].Y = 6.35

	pointsOfPolynomialPlot[170].X = 1.70
	pointsOfPolynomialPlot[170].Y = 6.423

	pointsOfPolynomialPlot[171].X = 1.71
	pointsOfPolynomialPlot[171].Y = 6.496

	pointsOfPolynomialPlot[172].X = 1.72
	pointsOfPolynomialPlot[172].Y = 6.57

	pointsOfPolynomialPlot[173].X = 1.73
	pointsOfPolynomialPlot[173].Y = 6.644

	pointsOfPolynomialPlot[174].X = 1.74
	pointsOfPolynomialPlot[174].Y = 6.72

	pointsOfPolynomialPlot[175].X = 1.75
	pointsOfPolynomialPlot[175].Y = 6.796

	pointsOfPolynomialPlot[176].X = 1.76
	pointsOfPolynomialPlot[176].Y = 6.874

	pointsOfPolynomialPlot[177].X = 1.77
	pointsOfPolynomialPlot[177].Y = 6.952

	pointsOfPolynomialPlot[178].X = 1.78
	pointsOfPolynomialPlot[178].Y = 7.031

	pointsOfPolynomialPlot[179].X = 1.79
	pointsOfPolynomialPlot[179].Y = 7.111

	pointsOfPolynomialPlot[180].X = 1.80
	pointsOfPolynomialPlot[180].Y = 7.192

	pointsOfPolynomialPlot[181].X = 1.81
	pointsOfPolynomialPlot[181].Y = 7.273

	pointsOfPolynomialPlot[182].X = 1.82
	pointsOfPolynomialPlot[182].Y = 7.356

	pointsOfPolynomialPlot[183].X = 1.83
	pointsOfPolynomialPlot[183].Y = 7.439

	pointsOfPolynomialPlot[184].X = 1.84
	pointsOfPolynomialPlot[184].Y = 7.523

	pointsOfPolynomialPlot[185].X = 1.85
	pointsOfPolynomialPlot[185].Y = 7.609

	pointsOfPolynomialPlot[186].X = 1.86
	pointsOfPolynomialPlot[186].Y = 7.695

	pointsOfPolynomialPlot[187].X = 1.87
	pointsOfPolynomialPlot[187].Y = 7.782

	pointsOfPolynomialPlot[188].X = 1.88
	pointsOfPolynomialPlot[188].Y = 7.87

	pointsOfPolynomialPlot[189].X = 1.89
	pointsOfPolynomialPlot[189].Y = 7.959

	pointsOfPolynomialPlot[190].X = 1.90
	pointsOfPolynomialPlot[190].Y = 8.049

	pointsOfPolynomialPlot[191].X = 1.91
	pointsOfPolynomialPlot[191].Y = 8.139

	pointsOfPolynomialPlot[192].X = 1.92
	pointsOfPolynomialPlot[192].Y = 8.231

	pointsOfPolynomialPlot[193].X = 1.93
	pointsOfPolynomialPlot[193].Y = 8.324

	pointsOfPolynomialPlot[194].X = 1.94
	pointsOfPolynomialPlot[194].Y = 8.417

	pointsOfPolynomialPlot[195].X = 1.95
	pointsOfPolynomialPlot[195].Y = 8.512

	pointsOfPolynomialPlot[196].X = 1.96
	pointsOfPolynomialPlot[196].Y = 8.607

	pointsOfPolynomialPlot[197].X = 1.97
	pointsOfPolynomialPlot[197].Y = 8.704

	pointsOfPolynomialPlot[198].X = 1.98
	pointsOfPolynomialPlot[198].Y = 8.802

	pointsOfPolynomialPlot[199].X = 1.99
	pointsOfPolynomialPlot[199].Y = 8.9

	pointsOfPolynomialPlot[200].X = 2.0
	pointsOfPolynomialPlot[200].Y = 9.0

	pointsOfPolynomialPlot[201].X = 2.01
	pointsOfPolynomialPlot[201].Y = 9.1

	pointsOfPolynomialPlot[202].X = 2.02
	pointsOfPolynomialPlot[202].Y = 9.202

	pointsOfPolynomialPlot[203].X = 2.03
	pointsOfPolynomialPlot[203].Y = 9.304

	pointsOfPolynomialPlot[204].X = 2.04
	pointsOfPolynomialPlot[204].Y = 9.408

	pointsOfPolynomialPlot[205].X = 2.05
	pointsOfPolynomialPlot[205].Y = 9.512

	pointsOfPolynomialPlot[206].X = 2.06
	pointsOfPolynomialPlot[206].Y = 9.618

	pointsOfPolynomialPlot[207].X = 2.07
	pointsOfPolynomialPlot[207].Y = 9.724

	pointsOfPolynomialPlot[208].X = 2.08
	pointsOfPolynomialPlot[208].Y = 9.832

	pointsOfPolynomialPlot[209].X = 2.09
	pointsOfPolynomialPlot[209].Y = 9.941

	pointsOfPolynomialPlot[210].X = 2.10
	pointsOfPolynomialPlot[210].Y = 10.051

	pointsOfPolynomialPlot[211].X = 2.11
	pointsOfPolynomialPlot[211].Y = 10.161

	pointsOfPolynomialPlot[212].X = 2.12
	pointsOfPolynomialPlot[212].Y = 10.273

	pointsOfPolynomialPlot[213].X = 2.13
	pointsOfPolynomialPlot[213].Y = 10.386

	pointsOfPolynomialPlot[214].X = 2.14
	pointsOfPolynomialPlot[214].Y = 10.5

	pointsOfPolynomialPlot[215].X = 2.15
	pointsOfPolynomialPlot[215].Y = 10.615

	pointsOfPolynomialPlot[216].X = 2.16
	pointsOfPolynomialPlot[216].Y = 10.732

	pointsOfPolynomialPlot[217].X = 2.17
	pointsOfPolynomialPlot[217].Y = 10.849

	pointsOfPolynomialPlot[218].X = 2.18
	pointsOfPolynomialPlot[218].Y = 10.967

	pointsOfPolynomialPlot[219].X = 2.19
	pointsOfPolynomialPlot[219].Y = 11.087

	pointsOfPolynomialPlot[220].X = 2.20
	pointsOfPolynomialPlot[220].Y = 11.208

	pointsOfPolynomialPlot[221].X = 2.21
	pointsOfPolynomialPlot[221].Y = 11.329

	pointsOfPolynomialPlot[222].X = 2.22
	pointsOfPolynomialPlot[222].Y = 11.452

	pointsOfPolynomialPlot[223].X = 2.23
	pointsOfPolynomialPlot[223].Y = 11.576

	pointsOfPolynomialPlot[224].X = 2.24
	pointsOfPolynomialPlot[224].Y = 11.701

	pointsOfPolynomialPlot[225].X = 2.25
	pointsOfPolynomialPlot[225].Y = 11.828

	pointsOfPolynomialPlot[226].X = 2.26
	pointsOfPolynomialPlot[226].Y = 11.955

	pointsOfPolynomialPlot[227].X = 2.27
	pointsOfPolynomialPlot[227].Y = 12.084

	pointsOfPolynomialPlot[228].X = 2.28
	pointsOfPolynomialPlot[228].Y = 12.214

	pointsOfPolynomialPlot[229].X = 2.29
	pointsOfPolynomialPlot[229].Y = 12.344

	pointsOfPolynomialPlot[230].X = 2.30
	pointsOfPolynomialPlot[230].Y = 12.477

	pointsOfPolynomialPlot[231].X = 2.31
	pointsOfPolynomialPlot[231].Y = 12.61

	pointsOfPolynomialPlot[232].X = 2.32
	pointsOfPolynomialPlot[232].Y = 12.744

	pointsOfPolynomialPlot[233].X = 2.33
	pointsOfPolynomialPlot[233].Y = 12.88

	pointsOfPolynomialPlot[234].X = 2.34
	pointsOfPolynomialPlot[234].Y = 12.017

	pointsOfPolynomialPlot[235].X = 2.35
	pointsOfPolynomialPlot[235].Y = 13.155

	pointsOfPolynomialPlot[236].X = 2.36
	pointsOfPolynomialPlot[236].Y = 13.294

	pointsOfPolynomialPlot[237].X = 2.37
	pointsOfPolynomialPlot[237].Y = 13.435

	pointsOfPolynomialPlot[238].X = 2.38
	pointsOfPolynomialPlot[238].Y = 13.576

	pointsOfPolynomialPlot[239].X = 2.39
	pointsOfPolynomialPlot[239].Y = 13.719

	pointsOfPolynomialPlot[240].X = 2.40
	pointsOfPolynomialPlot[240].Y = 13.864

	pointsOfPolynomialPlot[241].X = 2.41
	pointsOfPolynomialPlot[241].Y = 14.009

	pointsOfPolynomialPlot[242].X = 2.42
	pointsOfPolynomialPlot[242].Y = 14.156

	pointsOfPolynomialPlot[243].X = 2.43
	pointsOfPolynomialPlot[243].Y = 14.304

	pointsOfPolynomialPlot[244].X = 2.44
	pointsOfPolynomialPlot[244].Y = 14.453

	pointsOfPolynomialPlot[245].X = 2.45
	pointsOfPolynomialPlot[245].Y = 14.603

	pointsOfPolynomialPlot[246].X = 2.46
	pointsOfPolynomialPlot[246].Y = 14.755

	pointsOfPolynomialPlot[247].X = 2.47
	pointsOfPolynomialPlot[247].Y = 14.908

	pointsOfPolynomialPlot[248].X = 2.48
	pointsOfPolynomialPlot[248].Y = 15.062

	pointsOfPolynomialPlot[249].X = 2.49
	pointsOfPolynomialPlot[249].Y = 15.218

	pointsOfPolynomialPlot[250].X = 2.50
	pointsOfPolynomialPlot[250].Y = 15.375

	pointsOfPolynomialPlot[251].X = 2.51
	pointsOfPolynomialPlot[251].Y = 15.533

	pointsOfPolynomialPlot[252].X = 2.52
	pointsOfPolynomialPlot[252].Y = 15.692

	pointsOfPolynomialPlot[253].X = 2.53
	pointsOfPolynomialPlot[253].Y = 15.853

	pointsOfPolynomialPlot[254].X = 2.54
	pointsOfPolynomialPlot[254].Y = 16.015

	pointsOfPolynomialPlot[255].X = 2.55
	pointsOfPolynomialPlot[255].Y = 16.178

	pointsOfPolynomialPlot[256].X = 2.56
	pointsOfPolynomialPlot[256].Y = 16.343

	pointsOfPolynomialPlot[257].X = 2.57
	pointsOfPolynomialPlot[257].Y = 16.509

	pointsOfPolynomialPlot[258].X = 2.58
	pointsOfPolynomialPlot[258].Y = 16.677

	pointsOfPolynomialPlot[259].X = 2.59
	pointsOfPolynomialPlot[259].Y = 16.845

	pointsOfPolynomialPlot[260].X = 2.60
	pointsOfPolynomialPlot[260].Y = 17.016

	pointsOfPolynomialPlot[261].X = 2.61
	pointsOfPolynomialPlot[261].Y = 17.186

	pointsOfPolynomialPlot[262].X = 2.62
	pointsOfPolynomialPlot[262].Y = 17.36

	pointsOfPolynomialPlot[263].X = 2.63
	pointsOfPolynomialPlot[263].Y = 17.534

	pointsOfPolynomialPlot[264].X = 2.64
	pointsOfPolynomialPlot[264].Y = 17.71

	pointsOfPolynomialPlot[265].X = 2.65
	pointsOfPolynomialPlot[265].Y = 17.887

	pointsOfPolynomialPlot[266].X = 2.66
	pointsOfPolynomialPlot[266].Y = 18.065

	pointsOfPolynomialPlot[267].X = 2.67
	pointsOfPolynomialPlot[267].Y = 18.245

	pointsOfPolynomialPlot[268].X = 2.68
	pointsOfPolynomialPlot[268].Y = 18.426

	pointsOfPolynomialPlot[269].X = 2.69
	pointsOfPolynomialPlot[269].Y = 18.609

	pointsOfPolynomialPlot[270].X = 2.70
	pointsOfPolynomialPlot[270].Y = 18.793

	pointsOfPolynomialPlot[271].X = 2.71
	pointsOfPolynomialPlot[271].Y = 18.978

	pointsOfPolynomialPlot[272].X = 2.72
	pointsOfPolynomialPlot[272].Y = 19.165

	pointsOfPolynomialPlot[273].X = 2.73
	pointsOfPolynomialPlot[273].Y = 19.353

	pointsOfPolynomialPlot[274].X = 2.74
	pointsOfPolynomialPlot[274].Y = 19.543

	pointsOfPolynomialPlot[275].X = 2.75
	pointsOfPolynomialPlot[275].Y = 19.734

	pointsOfPolynomialPlot[276].X = 2.76
	pointsOfPolynomialPlot[276].Y = 19.927

	pointsOfPolynomialPlot[277].X = 2.77
	pointsOfPolynomialPlot[277].Y = 20.121

	pointsOfPolynomialPlot[278].X = 2.78
	pointsOfPolynomialPlot[278].Y = 20.316

	pointsOfPolynomialPlot[279].X = 2.79
	pointsOfPolynomialPlot[279].Y = 20.513

	pointsOfPolynomialPlot[280].X = 2.80
	pointsOfPolynomialPlot[280].Y = 20.712

	pointsOfPolynomialPlot[281].X = 2.81
	pointsOfPolynomialPlot[281].Y = 20.911

	pointsOfPolynomialPlot[282].X = 2.82
	pointsOfPolynomialPlot[282].Y = 21.113

	pointsOfPolynomialPlot[283].X = 2.83
	pointsOfPolynomialPlot[283].Y = 21.316

	pointsOfPolynomialPlot[284].X = 2.84
	pointsOfPolynomialPlot[284].Y = 21.52

	pointsOfPolynomialPlot[285].X = 2.85
	pointsOfPolynomialPlot[285].Y = 21.726

	pointsOfPolynomialPlot[286].X = 2.86
	pointsOfPolynomialPlot[286].Y = 21.934

	pointsOfPolynomialPlot[287].X = 2.87
	pointsOfPolynomialPlot[287].Y = 22.143

	pointsOfPolynomialPlot[288].X = 2.88
	pointsOfPolynomialPlot[288].Y = 22.353

	pointsOfPolynomialPlot[289].X = 2.89
	pointsOfPolynomialPlot[289].Y = 22.565

	pointsOfPolynomialPlot[290].X = 2.90
	pointsOfPolynomialPlot[290].Y = 22.779

	pointsOfPolynomialPlot[291].X = 2.91
	pointsOfPolynomialPlot[291].Y = 22.994

	pointsOfPolynomialPlot[292].X = 2.92
	pointsOfPolynomialPlot[292].Y = 23.21

	pointsOfPolynomialPlot[293].X = 2.93
	pointsOfPolynomialPlot[293].Y = 23.428

	pointsOfPolynomialPlot[294].X = 2.94
	pointsOfPolynomialPlot[294].Y = 23.648

	pointsOfPolynomialPlot[295].X = 2.95
	pointsOfPolynomialPlot[295].Y = 23.869

	pointsOfPolynomialPlot[296].X = 2.96
	pointsOfPolynomialPlot[296].Y = 24.092

	pointsOfPolynomialPlot[297].X = 2.97
	pointsOfPolynomialPlot[297].Y = 24.317

	pointsOfPolynomialPlot[298].X = 2.98
	pointsOfPolynomialPlot[298].Y = 24.543

	pointsOfPolynomialPlot[299].X = 2.99
	pointsOfPolynomialPlot[299].Y = 24.77

	pointsOfPolynomialPlot[300].X = 3.0
	pointsOfPolynomialPlot[300].Y = 25.0

	pointsOfPolynomialPlot[301].X = 3.01
	pointsOfPolynomialPlot[301].Y = 25.23

	pointsOfPolynomialPlot[302].X = 3.02
	pointsOfPolynomialPlot[302].Y = 25.463

	pointsOfPolynomialPlot[303].X = 3.03
	pointsOfPolynomialPlot[303].Y = 25.697

	pointsOfPolynomialPlot[304].X = 3.04
	pointsOfPolynomialPlot[304].Y = 25.932

	pointsOfPolynomialPlot[305].X = 3.05
	pointsOfPolynomialPlot[305].Y = 26.17

	pointsOfPolynomialPlot[306].X = 3.06
	pointsOfPolynomialPlot[306].Y = 26.409

	pointsOfPolynomialPlot[307].X = 3.07
	pointsOfPolynomialPlot[307].Y = 26.649

	pointsOfPolynomialPlot[308].X = 3.08
	pointsOfPolynomialPlot[308].Y = 26.891

	pointsOfPolynomialPlot[309].X = 3.09
	pointsOfPolynomialPlot[309].Y = 27.135

	pointsOfPolynomialPlot[310].X = 3.10
	pointsOfPolynomialPlot[310].Y = 27.381

	pointsOfPolynomialPlot[311].X = 3.11
	pointsOfPolynomialPlot[311].Y = 27.628

	pointsOfPolynomialPlot[312].X = 3.12
	pointsOfPolynomialPlot[312].Y = 27.876

	pointsOfPolynomialPlot[313].X = 3.13
	pointsOfPolynomialPlot[313].Y = 28.127

	pointsOfPolynomialPlot[314].X = 3.14
	pointsOfPolynomialPlot[314].Y = 28.379

	pointsOfPolynomialPlot[315].X = 3.15
	pointsOfPolynomialPlot[315].Y = 28.633

	pointsOfPolynomialPlot[316].X = 3.16
	pointsOfPolynomialPlot[316].Y = 28.888

	pointsOfPolynomialPlot[317].X = 3.17
	pointsOfPolynomialPlot[317].Y = 29.146

	pointsOfPolynomialPlot[318].X = 3.18
	pointsOfPolynomialPlot[318].Y = 29.405

	pointsOfPolynomialPlot[319].X = 3.19
	pointsOfPolynomialPlot[319].Y = 29.665

	pointsOfPolynomialPlot[320].X = 3.20
	pointsOfPolynomialPlot[320].Y = 29.928

	pointsOfPolynomialPlot[321].X = 3.21
	pointsOfPolynomialPlot[321].Y = 30.192

	pointsOfPolynomialPlot[322].X = 3.22
	pointsOfPolynomialPlot[322].Y = 30.457

	pointsOfPolynomialPlot[323].X = 3.23
	pointsOfPolynomialPlot[323].Y = 30.725

	pointsOfPolynomialPlot[324].X = 3.24
	pointsOfPolynomialPlot[324].Y = 30.994

	pointsOfPolynomialPlot[325].X = 3.25
	pointsOfPolynomialPlot[325].Y = 31.265

	pointsOfPolynomialPlot[326].X = 3.26
	pointsOfPolynomialPlot[326].Y = 31.538

	pointsOfPolynomialPlot[327].X = 3.27
	pointsOfPolynomialPlot[327].Y = 31.812

	pointsOfPolynomialPlot[328].X = 3.28
	pointsOfPolynomialPlot[328].Y = 32.089

	pointsOfPolynomialPlot[329].X = 3.29
	pointsOfPolynomialPlot[329].Y = 32.367

	pointsOfPolynomialPlot[330].X = 3.30
	pointsOfPolynomialPlot[330].Y = 32.647

	pointsOfPolynomialPlot[331].X = 3.31
	pointsOfPolynomialPlot[331].Y = 32.928

	pointsOfPolynomialPlot[332].X = 3.32
	pointsOfPolynomialPlot[332].Y = 33.212

	pointsOfPolynomialPlot[333].X = 3.33
	pointsOfPolynomialPlot[333].Y = 33.497

	pointsOfPolynomialPlot[334].X = 3.34
	pointsOfPolynomialPlot[334].Y = 33.784

	pointsOfPolynomialPlot[335].X = 3.35
	pointsOfPolynomialPlot[335].Y = 34.072

	pointsOfPolynomialPlot[336].X = 3.36
	pointsOfPolynomialPlot[336].Y = 34.363

	pointsOfPolynomialPlot[337].X = 3.37
	pointsOfPolynomialPlot[337].Y = 34.655

	pointsOfPolynomialPlot[338].X = 3.38
	pointsOfPolynomialPlot[338].Y = 34.95

	pointsOfPolynomialPlot[339].X = 3.39
	pointsOfPolynomialPlot[339].Y = 35.246

	pointsOfPolynomialPlot[340].X = 3.40
	pointsOfPolynomialPlot[340].Y = 35.544

	pointsOfPolynomialPlot[341].X = 3.41
	pointsOfPolynomialPlot[341].Y = 35.843

	pointsOfPolynomialPlot[342].X = 3.42
	pointsOfPolynomialPlot[342].Y = 36.145

	pointsOfPolynomialPlot[343].X = 3.43
	pointsOfPolynomialPlot[343].Y = 36.448

	pointsOfPolynomialPlot[344].X = 3.44
	pointsOfPolynomialPlot[344].Y = 37.754

	pointsOfPolynomialPlot[345].X = 3.45
	pointsOfPolynomialPlot[345].Y = 37.061

	pointsOfPolynomialPlot[346].X = 3.46
	pointsOfPolynomialPlot[346].Y = 37.37

	pointsOfPolynomialPlot[347].X = 3.47
	pointsOfPolynomialPlot[347].Y = 37.681

	pointsOfPolynomialPlot[348].X = 3.48
	pointsOfPolynomialPlot[348].Y = 37.993

	pointsOfPolynomialPlot[349].X = 3.49
	pointsOfPolynomialPlot[349].Y = 38.308

	pointsOfPolynomialPlot[350].X = 3.50
	pointsOfPolynomialPlot[350].Y = 38.625

	pointsOfPolynomialPlot[351].X = 3.51
	pointsOfPolynomialPlot[351].Y = 38.943

	pointsOfPolynomialPlot[352].X = 3.52
	pointsOfPolynomialPlot[352].Y = 39.263

	pointsOfPolynomialPlot[353].X = 3.53
	pointsOfPolynomialPlot[353].Y = 39.586

	pointsOfPolynomialPlot[354].X = 3.54
	pointsOfPolynomialPlot[354].Y = 39.91

	pointsOfPolynomialPlot[355].X = 3.55
	pointsOfPolynomialPlot[355].Y = 40.236

	pointsOfPolynomialPlot[356].X = 3.56
	pointsOfPolynomialPlot[356].Y = 40.564

	pointsOfPolynomialPlot[357].X = 3.57
	pointsOfPolynomialPlot[357].Y = 40.894

	pointsOfPolynomialPlot[358].X = 3.58
	pointsOfPolynomialPlot[358].Y = 41.226

	pointsOfPolynomialPlot[359].X = 3.59
	pointsOfPolynomialPlot[359].Y = 41.56

	pointsOfPolynomialPlot[360].X = 3.60
	pointsOfPolynomialPlot[360].Y = 41.896

	pointsOfPolynomialPlot[361].X = 3.61
	pointsOfPolynomialPlot[361].Y = 42.233

	pointsOfPolynomialPlot[362].X = 3.62
	pointsOfPolynomialPlot[362].Y = 42.573

	pointsOfPolynomialPlot[363].X = 3.65
	pointsOfPolynomialPlot[363].Y = 42.915

	pointsOfPolynomialPlot[364].X = 3.64
	pointsOfPolynomialPlot[364].Y = 43.258

	pointsOfPolynomialPlot[365].X = 3.65
	pointsOfPolynomialPlot[365].Y = 43.604

	pointsOfPolynomialPlot[366].X = 3.66
	pointsOfPolynomialPlot[366].Y = 43.952

	pointsOfPolynomialPlot[367].X = 3.67
	pointsOfPolynomialPlot[367].Y = 44.302

	pointsOfPolynomialPlot[368].X = 3.68
	pointsOfPolynomialPlot[368].Y = 44.653

	pointsOfPolynomialPlot[369].X = 3.69
	pointsOfPolynomialPlot[369].Y = 45.007

	pointsOfPolynomialPlot[370].X = 3.70
	pointsOfPolynomialPlot[370].Y = 45.363

	pointsOfPolynomialPlot[371].X = 3.71
	pointsOfPolynomialPlot[371].Y = 45.72

	pointsOfPolynomialPlot[372].X = 3.72
	pointsOfPolynomialPlot[372].Y = 46.08

	pointsOfPolynomialPlot[373].X = 3.73
	pointsOfPolynomialPlot[373].Y = 46.442

	pointsOfPolynomialPlot[374].X = 3.74
	pointsOfPolynomialPlot[374].Y = 46.806

	pointsOfPolynomialPlot[375].X = 3.75
	pointsOfPolynomialPlot[375].Y = 47.171

	pointsOfPolynomialPlot[376].X = 3.76
	pointsOfPolynomialPlot[376].Y = 47.539

	pointsOfPolynomialPlot[377].X = 3.77
	pointsOfPolynomialPlot[377].Y = 47.909

	pointsOfPolynomialPlot[378].X = 3.78
	pointsOfPolynomialPlot[378].Y = 48.281

	pointsOfPolynomialPlot[379].X = 3.79
	pointsOfPolynomialPlot[379].Y = 48.655

	pointsOfPolynomialPlot[380].X = 3.80
	pointsOfPolynomialPlot[380].Y = 49.032

	pointsOfPolynomialPlot[381].X = 3.81
	pointsOfPolynomialPlot[381].Y = 49.41

	pointsOfPolynomialPlot[382].X = 3.82
	pointsOfPolynomialPlot[382].Y = 49.79

	pointsOfPolynomialPlot[383].X = 3.83
	pointsOfPolynomialPlot[383].Y = 50.173

	pointsOfPolynomialPlot[384].X = 3.84
	pointsOfPolynomialPlot[384].Y = 50.557

	pointsOfPolynomialPlot[385].X = 3.85
	pointsOfPolynomialPlot[385].Y = 50.944

	pointsOfPolynomialPlot[386].X = 3.86
	pointsOfPolynomialPlot[386].Y = 51.332

	pointsOfPolynomialPlot[387].X = 3.87
	pointsOfPolynomialPlot[387].Y = 51.723

	pointsOfPolynomialPlot[388].X = 3.88
	pointsOfPolynomialPlot[388].Y = 52.116

	pointsOfPolynomialPlot[389].X = 3.89
	pointsOfPolynomialPlot[389].Y = 52.511

	pointsOfPolynomialPlot[390].X = 3.90
	pointsOfPolynomialPlot[390].Y = 52.909

	pointsOfPolynomialPlot[391].X = 3.91
	pointsOfPolynomialPlot[391].Y = 53.308

	pointsOfPolynomialPlot[392].X = 3.92
	pointsOfPolynomialPlot[392].Y = 53.709

	pointsOfPolynomialPlot[393].X = 3.93
	pointsOfPolynomialPlot[393].Y = 54.113

	pointsOfPolynomialPlot[394].X = 3.94
	pointsOfPolynomialPlot[394].Y = 54.519

	pointsOfPolynomialPlot[395].X = 3.95
	pointsOfPolynomialPlot[395].Y = 54.927

	pointsOfPolynomialPlot[396].X = 3.96
	pointsOfPolynomialPlot[396].Y = 55.337

	pointsOfPolynomialPlot[397].X = 3.97
	pointsOfPolynomialPlot[397].Y = 55.749

	pointsOfPolynomialPlot[398].X = 3.98
	pointsOfPolynomialPlot[398].Y = 56.164

	pointsOfPolynomialPlot[399].X = 3.99
	pointsOfPolynomialPlot[399].Y = 56.581

	pointsOfPolynomialPlot[400].X = 4.0
	pointsOfPolynomialPlot[400].Y = 57.0

	pointsOfPolynomialPlot[401].X = 4.01
	pointsOfPolynomialPlot[401].Y = 57.421

	pointsOfPolynomialPlot[402].X = 4.02
	pointsOfPolynomialPlot[402].Y = 57.844

	pointsOfPolynomialPlot[403].X = 4.03
	pointsOfPolynomialPlot[403].Y = 58.269

	pointsOfPolynomialPlot[404].X = 4.04
	pointsOfPolynomialPlot[404].Y = 58.697

	pointsOfPolynomialPlot[405].X = 4.05
	pointsOfPolynomialPlot[405].Y = 59.127

	pointsOfPolynomialPlot[406].X = 4.06
	pointsOfPolynomialPlot[406].Y = 59.559

	pointsOfPolynomialPlot[407].X = 4.07
	pointsOfPolynomialPlot[407].Y = 59.994

	pointsOfPolynomialPlot[408].X = 4.08
	pointsOfPolynomialPlot[408].Y = 60.43

	pointsOfPolynomialPlot[409].X = 4.09
	pointsOfPolynomialPlot[409].Y = 60.869

	pointsOfPolynomialPlot[410].X = 4.10
	pointsOfPolynomialPlot[410].Y = 61.311

	pointsOfPolynomialPlot[411].X = 4.11
	pointsOfPolynomialPlot[411].Y = 61.754

	pointsOfPolynomialPlot[412].X = 4.12
	pointsOfPolynomialPlot[412].Y = 62.2

	pointsOfPolynomialPlot[413].X = 4.13
	pointsOfPolynomialPlot[413].Y = 62.648

	pointsOfPolynomialPlot[414].X = 4.14
	pointsOfPolynomialPlot[414].Y = 63.098

	pointsOfPolynomialPlot[415].X = 4.15
	pointsOfPolynomialPlot[415].Y = 63.55

	pointsOfPolynomialPlot[416].X = 4.16
	pointsOfPolynomialPlot[416].Y = 64.005

	pointsOfPolynomialPlot[417].X = 4.17
	pointsOfPolynomialPlot[417].Y = 64.462

	pointsOfPolynomialPlot[418].X = 4.18
	pointsOfPolynomialPlot[418].Y = 64.922

	pointsOfPolynomialPlot[419].X = 4.19
	pointsOfPolynomialPlot[419].Y = 64.384

	pointsOfPolynomialPlot[420].X = 4.20
	pointsOfPolynomialPlot[420].Y = 65.848

	pointsOfPolynomialPlot[421].X = 4.21
	pointsOfPolynomialPlot[421].Y = 66.314

	pointsOfPolynomialPlot[422].X = 4.22
	pointsOfPolynomialPlot[422].Y = 66.783

	pointsOfPolynomialPlot[423].X = 4.23
	pointsOfPolynomialPlot[423].Y = 67.254

	pointsOfPolynomialPlot[424].X = 4.24
	pointsOfPolynomialPlot[424].Y = 67.727

	pointsOfPolynomialPlot[425].X = 4.25
	pointsOfPolynomialPlot[425].Y = 68.203

	pointsOfPolynomialPlot[426].X = 4.26
	pointsOfPolynomialPlot[426].Y = 68.681

	pointsOfPolynomialPlot[427].X = 4.27
	pointsOfPolynomialPlot[427].Y = 69.161

	pointsOfPolynomialPlot[428].X = 4.28
	pointsOfPolynomialPlot[428].Y = 69.644

	pointsOfPolynomialPlot[429].X = 4.29
	pointsOfPolynomialPlot[429].Y = 70.129

	pointsOfPolynomialPlot[430].X = 4.30
	pointsOfPolynomialPlot[430].Y = 70.617

	pointsOfPolynomialPlot[431].X = 4.31
	pointsOfPolynomialPlot[431].Y = 71.106

	pointsOfPolynomialPlot[432].X = 4.32
	pointsOfPolynomialPlot[432].Y = 71.599

	pointsOfPolynomialPlot[433].X = 4.33
	pointsOfPolynomialPlot[433].Y = 72.093

	pointsOfPolynomialPlot[434].X = 4.34
	pointsOfPolynomialPlot[434].Y = 72.59

	pointsOfPolynomialPlot[435].X = 4.35
	pointsOfPolynomialPlot[435].Y = 73.09

	pointsOfPolynomialPlot[436].X = 4.36
	pointsOfPolynomialPlot[436].Y = 72.592

	pointsOfPolynomialPlot[437].X = 4.37
	pointsOfPolynomialPlot[437].Y = 74.096

	pointsOfPolynomialPlot[438].X = 4.38
	pointsOfPolynomialPlot[438].Y = 74.603

	pointsOfPolynomialPlot[439].X = 4.39
	pointsOfPolynomialPlot[439].Y = 75.112

	pointsOfPolynomialPlot[440].X = 4.40
	pointsOfPolynomialPlot[440].Y = 75.624

	pointsOfPolynomialPlot[441].X = 4.41
	pointsOfPolynomialPlot[441].Y = 76.138

	pointsOfPolynomialPlot[442].X = 4.42
	pointsOfPolynomialPlot[442].Y = 76.654

	pointsOfPolynomialPlot[443].X = 4.43
	pointsOfPolynomialPlot[443].Y = 77.173

	pointsOfPolynomialPlot[444].X = 4.44
	pointsOfPolynomialPlot[444].Y = 77.694

	pointsOfPolynomialPlot[445].X = 4.45
	pointsOfPolynomialPlot[445].Y = 78.218

	pointsOfPolynomialPlot[446].X = 4.46
	pointsOfPolynomialPlot[446].Y = 78.744

	pointsOfPolynomialPlot[447].X = 4.47
	pointsOfPolynomialPlot[447].Y = 79.273

	pointsOfPolynomialPlot[448].X = 4.48
	pointsOfPolynomialPlot[448].Y = 79.805

	pointsOfPolynomialPlot[449].X = 4.49
	pointsOfPolynomialPlot[449].Y = 80.338

	pointsOfPolynomialPlot[450].X = 4.50
	pointsOfPolynomialPlot[450].Y = 80.875

	pointsOfPolynomialPlot[451].X = 4.51
	pointsOfPolynomialPlot[451].Y = 81.413

	pointsOfPolynomialPlot[452].X = 4.52
	pointsOfPolynomialPlot[452].Y = 81.955

	pointsOfPolynomialPlot[453].X = 4.53
	pointsOfPolynomialPlot[453].Y = 82.498

	pointsOfPolynomialPlot[454].X = 4.54
	pointsOfPolynomialPlot[454].Y = 83.045

	pointsOfPolynomialPlot[455].X = 4.55
	pointsOfPolynomialPlot[455].Y = 83.593

	pointsOfPolynomialPlot[456].X = 4.56
	pointsOfPolynomialPlot[456].Y = 84.145

	pointsOfPolynomialPlot[457].X = 4.57
	pointsOfPolynomialPlot[457].Y = 84.699

	pointsOfPolynomialPlot[458].X = 4.58
	pointsOfPolynomialPlot[458].Y = 85.255

	pointsOfPolynomialPlot[459].X = 4.59
	pointsOfPolynomialPlot[459].Y = 85.814

	pointsOfPolynomialPlot[460].X = 4.60
	pointsOfPolynomialPlot[460].Y = 86.376

	pointsOfPolynomialPlot[461].X = 4.61
	pointsOfPolynomialPlot[461].Y = 86.94

	pointsOfPolynomialPlot[462].X = 4.62
	pointsOfPolynomialPlot[462].Y = 87.506

	pointsOfPolynomialPlot[463].X = 4.63
	pointsOfPolynomialPlot[463].Y = 88.075

	pointsOfPolynomialPlot[464].X = 4.64
	pointsOfPolynomialPlot[464].Y = 88.647

	pointsOfPolynomialPlot[465].X = 4.65
	pointsOfPolynomialPlot[465].Y = 89.222

	pointsOfPolynomialPlot[466].X = 4.66
	pointsOfPolynomialPlot[466].Y = 89.799

	pointsOfPolynomialPlot[467].X = 4.67
	pointsOfPolynomialPlot[467].Y = 90.378

	pointsOfPolynomialPlot[468].X = 4.68
	pointsOfPolynomialPlot[468].Y = 90.96

	pointsOfPolynomialPlot[469].X = 4.69
	pointsOfPolynomialPlot[469].Y = 91.545

	pointsOfPolynomialPlot[470].X = 4.70
	pointsOfPolynomialPlot[470].Y = 92.133

	pointsOfPolynomialPlot[471].X = 4.71
	pointsOfPolynomialPlot[471].Y = 92.723

	pointsOfPolynomialPlot[472].X = 4.72
	pointsOfPolynomialPlot[472].Y = 93.315

	pointsOfPolynomialPlot[473].X = 4.73
	pointsOfPolynomialPlot[473].Y = 93.91

	pointsOfPolynomialPlot[474].X = 4.74
	pointsOfPolynomialPlot[474].Y = 94.508

	pointsOfPolynomialPlot[475].X = 4.75
	pointsOfPolynomialPlot[475].Y = 95.109

	pointsOfPolynomialPlot[476].X = 4.76
	pointsOfPolynomialPlot[476].Y = 95.712

	pointsOfPolynomialPlot[477].X = 4.77
	pointsOfPolynomialPlot[477].Y = 96.318

	pointsOfPolynomialPlot[478].X = 4.78
	pointsOfPolynomialPlot[478].Y = 96.927

	pointsOfPolynomialPlot[479].X = 4.79
	pointsOfPolynomialPlot[479].Y = 97.538

	pointsOfPolynomialPlot[480].X = 4.80
	pointsOfPolynomialPlot[480].Y = 98.152

	pointsOfPolynomialPlot[481].X = 4.81
	pointsOfPolynomialPlot[481].Y = 98.769

	pointsOfPolynomialPlot[482].X = 4.82
	pointsOfPolynomialPlot[482].Y = 99.387

	pointsOfPolynomialPlot[483].X = 4.83
	pointsOfPolynomialPlot[483].Y = 100.009

	pointsOfPolynomialPlot[484].X = 4.84
	pointsOfPolynomialPlot[484].Y = 100.634

	pointsOfPolynomialPlot[485].X = 4.85
	pointsOfPolynomialPlot[485].Y = 101.261

	pointsOfPolynomialPlot[486].X = 4.86
	pointsOfPolynomialPlot[486].Y = 101.891

	pointsOfPolynomialPlot[487].X = 4.87
	pointsOfPolynomialPlot[487].Y = 102.524

	pointsOfPolynomialPlot[488].X = 4.88
	pointsOfPolynomialPlot[488].Y = 103.159

	pointsOfPolynomialPlot[489].X = 4.89
	pointsOfPolynomialPlot[489].Y = 103.798

	pointsOfPolynomialPlot[490].X = 4.90
	pointsOfPolynomialPlot[490].Y = 104.439

	pointsOfPolynomialPlot[491].X = 4.91
	pointsOfPolynomialPlot[491].Y = 105.082

	pointsOfPolynomialPlot[492].X = 4.92
	pointsOfPolynomialPlot[492].Y = 105.729

	pointsOfPolynomialPlot[493].X = 4.93
	pointsOfPolynomialPlot[493].Y = 106.378

	pointsOfPolynomialPlot[494].X = 4.94
	pointsOfPolynomialPlot[494].Y = 107.03

	pointsOfPolynomialPlot[495].X = 4.95
	pointsOfPolynomialPlot[495].Y = 107.684

	pointsOfPolynomialPlot[496].X = 4.96
	pointsOfPolynomialPlot[496].Y = 108.342

	pointsOfPolynomialPlot[497].X = 4.97
	pointsOfPolynomialPlot[497].Y = 109.002

	pointsOfPolynomialPlot[498].X = 4.98
	pointsOfPolynomialPlot[498].Y = 109.665

	pointsOfPolynomialPlot[499].X = 4.99
	pointsOfPolynomialPlot[499].Y = 110.331

	pointsOfPolynomialPlot[500].X = 5.0
	pointsOfPolynomialPlot[500].Y = 111.0

	pointsOfPolynomialPlot[501].X = 5.01
	pointsOfPolynomialPlot[501].Y = 111.671

	pointsOfPolynomialPlot[502].X = 5.02
	pointsOfPolynomialPlot[502].Y = 112.345

	pointsOfPolynomialPlot[503].X = 5.03
	pointsOfPolynomialPlot[503].Y = 113.022

	pointsOfPolynomialPlot[504].X = 5.04
	pointsOfPolynomialPlot[504].Y = 113.702

	pointsOfPolynomialPlot[505].X = 5.05
	pointsOfPolynomialPlot[505].Y = 114.385

	pointsOfPolynomialPlot[506].X = 5.06
	pointsOfPolynomialPlot[506].Y = 115.07

	pointsOfPolynomialPlot[507].X = 5.07
	pointsOfPolynomialPlot[507].Y = 115.758

	pointsOfPolynomialPlot[508].X = 5.08
	pointsOfPolynomialPlot[508].Y = 116.45

	pointsOfPolynomialPlot[509].X = 5.09
	pointsOfPolynomialPlot[509].Y = 117.144

	pointsOfPolynomialPlot[510].X = 5.10
	pointsOfPolynomialPlot[510].Y = 117.841

	pointsOfPolynomialPlot[511].X = 5.11
	pointsOfPolynomialPlot[511].Y = 118.54

	pointsOfPolynomialPlot[512].X = 5.12
	pointsOfPolynomialPlot[512].Y = 119.243

	pointsOfPolynomialPlot[513].X = 5.13
	pointsOfPolynomialPlot[513].Y = 120.657

	pointsOfPolynomialPlot[514].X = 5.14
	pointsOfPolynomialPlot[514].Y = 120.657

	pointsOfPolynomialPlot[515].X = 5.15
	pointsOfPolynomialPlot[515].Y = 121.368

	pointsOfPolynomialPlot[516].X = 5.16
	pointsOfPolynomialPlot[516].Y = 122.082

	pointsOfPolynomialPlot[517].X = 5.17
	pointsOfPolynomialPlot[517].Y = 122.799

	pointsOfPolynomialPlot[518].X = 5.18
	pointsOfPolynomialPlot[518].Y = 123.519

	pointsOfPolynomialPlot[519].X = 5.19
	pointsOfPolynomialPlot[519].Y = 124.242

	pointsOfPolynomialPlot[520].X = 5.20
	pointsOfPolynomialPlot[520].Y = 124.968

	pointsOfPolynomialPlot[521].X = 5.21
	pointsOfPolynomialPlot[521].Y = 125.696

	pointsOfPolynomialPlot[522].X = 5.22
	pointsOfPolynomialPlot[522].Y = 126.64

	pointsOfPolynomialPlot[523].X = 5.23
	pointsOfPolynomialPlot[523].Y = 127.162

	pointsOfPolynomialPlot[524].X = 5.24
	pointsOfPolynomialPlot[524].Y = 127.9

	pointsOfPolynomialPlot[525].X = 5.25
	pointsOfPolynomialPlot[525].Y = 128.64

	pointsOfPolynomialPlot[526].X = 5.26
	pointsOfPolynomialPlot[526].Y = 129.631

	pointsOfPolynomialPlot[527].X = 5.27
	pointsOfPolynomialPlot[527].Y = 130.13

	pointsOfPolynomialPlot[528].X = 5.28
	pointsOfPolynomialPlot[528].Y = 130.879

	pointsOfPolynomialPlot[529].X = 5.29
	pointsOfPolynomialPlot[529].Y = 131.631

	pointsOfPolynomialPlot[530].X = 5.30
	pointsOfPolynomialPlot[530].Y = 132.387

	pointsOfPolynomialPlot[531].X = 5.31
	pointsOfPolynomialPlot[531].Y = 133.145

	pointsOfPolynomialPlot[532].X = 5.32
	pointsOfPolynomialPlot[532].Y = 133.67

	pointsOfPolynomialPlot[533].X = 5.33
	pointsOfPolynomialPlot[533].Y = 134.67

	pointsOfPolynomialPlot[534].X = 5.34
	pointsOfPolynomialPlot[534].Y = 135.437

	pointsOfPolynomialPlot[535].X = 5.35
	pointsOfPolynomialPlot[535].Y = 136.207

	pointsOfPolynomialPlot[536].X = 5.36
	pointsOfPolynomialPlot[536].Y = 136.981

	pointsOfPolynomialPlot[537].X = 5.37
	pointsOfPolynomialPlot[537].Y = 137.757

	pointsOfPolynomialPlot[538].X = 5.38
	pointsOfPolynomialPlot[538].Y = 138.536

	pointsOfPolynomialPlot[539].X = 5.39
	pointsOfPolynomialPlot[539].Y = 139.318

	pointsOfPolynomialPlot[540].X = 5.40
	pointsOfPolynomialPlot[540].Y = 140.104

	pointsOfPolynomialPlot[541].X = 5.41
	pointsOfPolynomialPlot[541].Y = 140.892

	pointsOfPolynomialPlot[542].X = 5.42
	pointsOfPolynomialPlot[542].Y = 141.683

	pointsOfPolynomialPlot[543].X = 5.43
	pointsOfPolynomialPlot[543].Y = 142.478

	pointsOfPolynomialPlot[544].X = 5.44
	pointsOfPolynomialPlot[544].Y = 143.275

	pointsOfPolynomialPlot[545].X = 5.45
	pointsOfPolynomialPlot[545].Y = 144.076

	pointsOfPolynomialPlot[546].X = 5.46
	pointsOfPolynomialPlot[546].Y = 144.879

	pointsOfPolynomialPlot[547].X = 5.47
	pointsOfPolynomialPlot[547].Y = 145.686

	pointsOfPolynomialPlot[548].X = 5.48
	pointsOfPolynomialPlot[548].Y = 146.496

	pointsOfPolynomialPlot[549].X = 5.49
	pointsOfPolynomialPlot[549].Y = 147.309

	pointsOfPolynomialPlot[550].X = 5.50
	pointsOfPolynomialPlot[550].Y = 148.125

	pointsOfPolynomialPlot[551].X = 5.51
	pointsOfPolynomialPlot[551].Y = 148.944

	pointsOfPolynomialPlot[552].X = 5.52
	pointsOfPolynomialPlot[552].Y = 149.766

	pointsOfPolynomialPlot[553].X = 5.53
	pointsOfPolynomialPlot[553].Y = 150.951

	pointsOfPolynomialPlot[554].X = 5.54
	pointsOfPolynomialPlot[554].Y = 151.419

	pointsOfPolynomialPlot[555].X = 5.55
	pointsOfPolynomialPlot[555].Y = 152.251

	pointsOfPolynomialPlot[556].X = 5.56
	pointsOfPolynomialPlot[556].Y = 153.086

	pointsOfPolynomialPlot[557].X = 5.57
	pointsOfPolynomialPlot[557].Y = 153.923

	pointsOfPolynomialPlot[558].X = 5.58
	pointsOfPolynomialPlot[558].Y = 154.764

	pointsOfPolynomialPlot[559].X = 5.59
	pointsOfPolynomialPlot[559].Y = 155.608

	pointsOfPolynomialPlot[560].X = 5.60
	pointsOfPolynomialPlot[560].Y = 156.456

	pointsOfPolynomialPlot[561].X = 5.61
	pointsOfPolynomialPlot[561].Y = 157.306

	pointsOfPolynomialPlot[562].X = 5.62
	pointsOfPolynomialPlot[562].Y = 158.159

	pointsOfPolynomialPlot[563].X = 5.63
	pointsOfPolynomialPlot[563].Y = 159.016

	pointsOfPolynomialPlot[564].X = 5.64
	pointsOfPolynomialPlot[564].Y = 159.876

	pointsOfPolynomialPlot[565].X = 5.65
	pointsOfPolynomialPlot[565].Y = 160.739

	pointsOfPolynomialPlot[566].X = 5.66
	pointsOfPolynomialPlot[566].Y = 161.605

	pointsOfPolynomialPlot[567].X = 5.67
	pointsOfPolynomialPlot[567].Y = 162.475

	pointsOfPolynomialPlot[568].X = 5.68
	pointsOfPolynomialPlot[568].Y = 163.348

	pointsOfPolynomialPlot[569].X = 5.69
	pointsOfPolynomialPlot[569].Y = 164.223

	pointsOfPolynomialPlot[570].X = 5.70
	pointsOfPolynomialPlot[570].Y = 165.103

	pointsOfPolynomialPlot[571].X = 5.71
	pointsOfPolynomialPlot[571].Y = 165.985

	pointsOfPolynomialPlot[572].X = 5.72
	pointsOfPolynomialPlot[572].Y = 166.87

	pointsOfPolynomialPlot[573].X = 5.73
	pointsOfPolynomialPlot[573].Y = 167.759

	pointsOfPolynomialPlot[574].X = 5.74
	pointsOfPolynomialPlot[574].Y = 168.651

	pointsOfPolynomialPlot[575].X = 5.75
	pointsOfPolynomialPlot[575].Y = 169.546

	pointsOfPolynomialPlot[576].X = 5.76
	pointsOfPolynomialPlot[576].Y = 170.445

	pointsOfPolynomialPlot[577].X = 5.77
	pointsOfPolynomialPlot[577].Y = 171.347

	pointsOfPolynomialPlot[578].X = 5.78
	pointsOfPolynomialPlot[578].Y = 172.252

	pointsOfPolynomialPlot[579].X = 5.79
	pointsOfPolynomialPlot[579].Y = 173.16

	pointsOfPolynomialPlot[580].X = 5.80
	pointsOfPolynomialPlot[580].Y = 174.072

	pointsOfPolynomialPlot[581].X = 5.81
	pointsOfPolynomialPlot[581].Y = 174.986

	pointsOfPolynomialPlot[582].X = 5.82
	pointsOfPolynomialPlot[582].Y = 175.905

	pointsOfPolynomialPlot[583].X = 5.83
	pointsOfPolynomialPlot[583].Y = 176.826

	pointsOfPolynomialPlot[584].X = 5.84
	pointsOfPolynomialPlot[584].Y = 177.751

	pointsOfPolynomialPlot[585].X = 5.85
	pointsOfPolynomialPlot[585].Y = 178.679

	pointsOfPolynomialPlot[586].X = 5.86
	pointsOfPolynomialPlot[586].Y = 179.61

	pointsOfPolynomialPlot[587].X = 5.87
	pointsOfPolynomialPlot[587].Y = 180.545

	pointsOfPolynomialPlot[588].X = 5.88
	pointsOfPolynomialPlot[588].Y = 181.483

	pointsOfPolynomialPlot[589].X = 5.89
	pointsOfPolynomialPlot[589].Y = 182.424

	pointsOfPolynomialPlot[590].X = 5.90
	pointsOfPolynomialPlot[590].Y = 183.369

	pointsOfPolynomialPlot[591].X = 5.91
	pointsOfPolynomialPlot[591].Y = 184.317

	pointsOfPolynomialPlot[592].X = 5.92
	pointsOfPolynomialPlot[592].Y = 185.268

	pointsOfPolynomialPlot[593].X = 5.93
	pointsOfPolynomialPlot[593].Y = 186.223

	pointsOfPolynomialPlot[594].X = 5.94
	pointsOfPolynomialPlot[594].Y = 187.181

	pointsOfPolynomialPlot[595].X = 5.95
	pointsOfPolynomialPlot[595].Y = 188.142

	pointsOfPolynomialPlot[596].X = 5.96
	pointsOfPolynomialPlot[596].Y = 189.107

	pointsOfPolynomialPlot[597].X = 5.97
	pointsOfPolynomialPlot[597].Y = 190.075

	pointsOfPolynomialPlot[598].X = 5.98
	pointsOfPolynomialPlot[598].Y = 191.046

	pointsOfPolynomialPlot[599].X = 5.99
	pointsOfPolynomialPlot[599].Y = 192.021

	pointsOfPolynomialPlot[600].X = 6.0
	pointsOfPolynomialPlot[600].Y = 193.0

	pointsOfPolynomialPlot[601].X = 6.01
	pointsOfPolynomialPlot[601].Y = 193.981

	pointsOfPolynomialPlot[602].X = 6.02
	pointsOfPolynomialPlot[602].Y = 194.966

	pointsOfPolynomialPlot[603].X = 6.03
	pointsOfPolynomialPlot[603].Y = 195.955

	pointsOfPolynomialPlot[604].X = 6.04
	pointsOfPolynomialPlot[604].Y = 196.947

	pointsOfPolynomialPlot[605].X = 6.05
	pointsOfPolynomialPlot[605].Y = 197.942

	pointsOfPolynomialPlot[606].X = 6.06
	pointsOfPolynomialPlot[606].Y = 198.941

	pointsOfPolynomialPlot[607].X = 6.07
	pointsOfPolynomialPlot[607].Y = 199.943

	pointsOfPolynomialPlot[608].X = 6.08
	pointsOfPolynomialPlot[608].Y = 200.949

	pointsOfPolynomialPlot[609].X = 6.09
	pointsOfPolynomialPlot[609].Y = 201.958

	pointsOfPolynomialPlot[610].X = 6.10
	pointsOfPolynomialPlot[610].Y = 202.971

	pointsOfPolynomialPlot[611].X = 6.11
	pointsOfPolynomialPlot[611].Y = 203.987

	pointsOfPolynomialPlot[612].X = 6.12
	pointsOfPolynomialPlot[612].Y = 205.006

	pointsOfPolynomialPlot[613].X = 6.13
	pointsOfPolynomialPlot[613].Y = 206.029

	pointsOfPolynomialPlot[614].X = 6.14
	pointsOfPolynomialPlot[614].Y = 207.055

	pointsOfPolynomialPlot[615].X = 6.15
	pointsOfPolynomialPlot[615].Y = 208.085

	pointsOfPolynomialPlot[616].X = 6.16
	pointsOfPolynomialPlot[616].Y = 209.119

	pointsOfPolynomialPlot[617].X = 6.17
	pointsOfPolynomialPlot[617].Y = 210.156

	pointsOfPolynomialPlot[618].X = 6.18
	pointsOfPolynomialPlot[618].Y = 211.196

	pointsOfPolynomialPlot[619].X = 6.19
	pointsOfPolynomialPlot[619].Y = 212.24

	pointsOfPolynomialPlot[620].X = 6.20
	pointsOfPolynomialPlot[620].Y = 213.288

	pointsOfPolynomialPlot[621].X = 6.21
	pointsOfPolynomialPlot[621].Y = 214.339

	pointsOfPolynomialPlot[622].X = 6.22
	pointsOfPolynomialPlot[622].Y = 215.393

	pointsOfPolynomialPlot[623].X = 6.23
	pointsOfPolynomialPlot[623].Y = 216.451

	pointsOfPolynomialPlot[624].X = 6.24
	pointsOfPolynomialPlot[624].Y = 217.513

	pointsOfPolynomialPlot[625].X = 6.25
	pointsOfPolynomialPlot[625].Y = 218.578

	pointsOfPolynomialPlot[626].X = 6.26
	pointsOfPolynomialPlot[626].Y = 219.646

	pointsOfPolynomialPlot[627].X = 6.27
	pointsOfPolynomialPlot[627].Y = 220.719

	pointsOfPolynomialPlot[628].X = 6.28
	pointsOfPolynomialPlot[628].Y = 221.794

	pointsOfPolynomialPlot[629].X = 6.29
	pointsOfPolynomialPlot[629].Y = 222.874

	pointsOfPolynomialPlot[630].X = 6.30
	pointsOfPolynomialPlot[630].Y = 223.957

	pointsOfPolynomialPlot[631].X = 6.31
	pointsOfPolynomialPlot[631].Y = 225.043

	pointsOfPolynomialPlot[632].X = 6.32
	pointsOfPolynomialPlot[632].Y = 226.133

	pointsOfPolynomialPlot[633].X = 6.33
	pointsOfPolynomialPlot[633].Y = 227.227

	pointsOfPolynomialPlot[634].X = 6.33
	pointsOfPolynomialPlot[634].Y = 228.324

	pointsOfPolynomialPlot[635].X = 6.35
	pointsOfPolynomialPlot[635].Y = 229.425

	pointsOfPolynomialPlot[636].X = 6.36
	pointsOfPolynomialPlot[636].Y = 230.529

	pointsOfPolynomialPlot[637].X = 6.37
	pointsOfPolynomialPlot[637].Y = 231.638

	pointsOfPolynomialPlot[638].X = 6.38
	pointsOfPolynomialPlot[638].Y = 232.749

	pointsOfPolynomialPlot[639].X = 6.39
	pointsOfPolynomialPlot[639].Y = 233.865

	pointsOfPolynomialPlot[640].X = 6.40
	pointsOfPolynomialPlot[640].Y = 234.984

	pointsOfPolynomialPlot[641].X = 6.41
	pointsOfPolynomialPlot[641].Y = 236.106

	pointsOfPolynomialPlot[642].X = 6.42
	pointsOfPolynomialPlot[642].Y = 237.232

	pointsOfPolynomialPlot[643].X = 6.43
	pointsOfPolynomialPlot[643].Y = 238.362

	pointsOfPolynomialPlot[644].X = 6.44
	pointsOfPolynomialPlot[644].Y = 239.496

	pointsOfPolynomialPlot[645].X = 6.45
	pointsOfPolynomialPlot[645].Y = 240.633

	pointsOfPolynomialPlot[646].X = 6.46
	pointsOfPolynomialPlot[646].Y = 241.774

	pointsOfPolynomialPlot[647].X = 6.47
	pointsOfPolynomialPlot[647].Y = 242.919

	pointsOfPolynomialPlot[648].X = 6.48
	pointsOfPolynomialPlot[648].Y = 244.067

	pointsOfPolynomialPlot[649].X = 6.49
	pointsOfPolynomialPlot[649].Y = 245.219

	pointsOfPolynomialPlot[650].X = 6.50
	pointsOfPolynomialPlot[650].Y = 246.375

	pointsOfPolynomialPlot[651].X = 6.51
	pointsOfPolynomialPlot[651].Y = 247.534

	pointsOfPolynomialPlot[652].X = 6.52
	pointsOfPolynomialPlot[652].Y = 248.697

	pointsOfPolynomialPlot[653].X = 6.53
	pointsOfPolynomialPlot[653].Y = 249.864

	pointsOfPolynomialPlot[654].X = 6.54
	pointsOfPolynomialPlot[654].Y = 251.034

	pointsOfPolynomialPlot[655].X = 6.55
	pointsOfPolynomialPlot[655].Y = 252.208

	pointsOfPolynomialPlot[656].X = 6.56
	pointsOfPolynomialPlot[656].Y = 253.386

	pointsOfPolynomialPlot[657].X = 6.57
	pointsOfPolynomialPlot[657].Y = 254.568

	pointsOfPolynomialPlot[658].X = 6.58
	pointsOfPolynomialPlot[658].Y = 255.753

	pointsOfPolynomialPlot[659].X = 6.59
	pointsOfPolynomialPlot[659].Y = 256.943

	pointsOfPolynomialPlot[660].X = 6.60
	pointsOfPolynomialPlot[660].Y = 258.136

	pointsOfPolynomialPlot[661].X = 6.61
	pointsOfPolynomialPlot[661].Y = 259.332

	pointsOfPolynomialPlot[662].X = 6.62
	pointsOfPolynomialPlot[662].Y = 260.533

	pointsOfPolynomialPlot[663].X = 6.63
	pointsOfPolynomialPlot[663].Y = 261.737

	pointsOfPolynomialPlot[664].X = 6.64
	pointsOfPolynomialPlot[664].Y = 262.945

	pointsOfPolynomialPlot[665].X = 6.65
	pointsOfPolynomialPlot[665].Y = 264.157

	pointsOfPolynomialPlot[666].X = 6.66
	pointsOfPolynomialPlot[666].Y = 265.372

	pointsOfPolynomialPlot[667].X = 6.67
	pointsOfPolynomialPlot[667].Y = 266.592

	pointsOfPolynomialPlot[668].X = 6.68
	pointsOfPolynomialPlot[668].Y = 267.815

	pointsOfPolynomialPlot[669].X = 6.69
	pointsOfPolynomialPlot[669].Y = 269.042

	pointsOfPolynomialPlot[670].X = 6.70
	pointsOfPolynomialPlot[670].Y = 270.273

	pointsOfPolynomialPlot[671].X = 6.71
	pointsOfPolynomialPlot[671].Y = 271.507

	pointsOfPolynomialPlot[672].X = 6.72
	pointsOfPolynomialPlot[672].Y = 272.746

	pointsOfPolynomialPlot[673].X = 6.73
	pointsOfPolynomialPlot[673].Y = 273.988

	pointsOfPolynomialPlot[674].X = 6.74
	pointsOfPolynomialPlot[674].Y = 275.234

	pointsOfPolynomialPlot[675].X = 6.75
	pointsOfPolynomialPlot[675].Y = 276.484

	pointsOfPolynomialPlot[676].X = 6.76
	pointsOfPolynomialPlot[676].Y = 277.738

	pointsOfPolynomialPlot[677].X = 6.77
	pointsOfPolynomialPlot[677].Y = 278.995

	pointsOfPolynomialPlot[678].X = 6.78
	pointsOfPolynomialPlot[678].Y = 280.257

	pointsOfPolynomialPlot[679].X = 6.79
	pointsOfPolynomialPlot[679].Y = 281.522

	pointsOfPolynomialPlot[680].X = 6.80
	pointsOfPolynomialPlot[680].Y = 282.792

	pointsOfPolynomialPlot[681].X = 6.81
	pointsOfPolynomialPlot[681].Y = 284.065

	pointsOfPolynomialPlot[682].X = 6.82
	pointsOfPolynomialPlot[682].Y = 285.342

	pointsOfPolynomialPlot[683].X = 6.83
	pointsOfPolynomialPlot[683].Y = 286.623

	pointsOfPolynomialPlot[684].X = 6.84
	pointsOfPolynomialPlot[684].Y = 287.907

	pointsOfPolynomialPlot[685].X = 6.85
	pointsOfPolynomialPlot[685].Y = 289.196

	pointsOfPolynomialPlot[686].X = 6.86
	pointsOfPolynomialPlot[686].Y = 290.489

	pointsOfPolynomialPlot[687].X = 6.87
	pointsOfPolynomialPlot[687].Y = 291.785

	pointsOfPolynomialPlot[688].X = 6.88
	pointsOfPolynomialPlot[688].Y = 293.086

	pointsOfPolynomialPlot[689].X = 6.89
	pointsOfPolynomialPlot[689].Y = 294.39

	pointsOfPolynomialPlot[690].X = 6.90
	pointsOfPolynomialPlot[690].Y = 295.699

	pointsOfPolynomialPlot[691].X = 6.91
	pointsOfPolynomialPlot[691].Y = 297.011

	pointsOfPolynomialPlot[692].X = 6.92
	pointsOfPolynomialPlot[692].Y = 298.327

	pointsOfPolynomialPlot[693].X = 6.93
	pointsOfPolynomialPlot[693].Y = 299.647

	pointsOfPolynomialPlot[694].X = 6.94
	pointsOfPolynomialPlot[694].Y = 300.971

	pointsOfPolynomialPlot[695].X = 6.95
	pointsOfPolynomialPlot[695].Y = 302.299

	pointsOfPolynomialPlot[696].X = 6.96
	pointsOfPolynomialPlot[696].Y = 303.631

	pointsOfPolynomialPlot[697].X = 6.97
	pointsOfPolynomialPlot[697].Y = 304.968

	pointsOfPolynomialPlot[698].X = 6.98
	pointsOfPolynomialPlot[698].Y = 306.308

	pointsOfPolynomialPlot[699].X = 6.99
	pointsOfPolynomialPlot[699].Y = 307.652

	pointsOfPolynomialPlot[700].X = 7.0
	pointsOfPolynomialPlot[700].Y = 309.0

	pointsOfPolynomialPlot[701].X = 7.01
	pointsOfPolynomialPlot[701].Y = 310.352

	pointsOfPolynomialPlot[702].X = 7.02
	pointsOfPolynomialPlot[702].Y = 311.708

	pointsOfPolynomialPlot[703].X = 7.03
	pointsOfPolynomialPlot[703].Y = 313.068

	pointsOfPolynomialPlot[704].X = 7.04
	pointsOfPolynomialPlot[704].Y = 314.432

	pointsOfPolynomialPlot[705].X = 7.05
	pointsOfPolynomialPlot[705].Y = 315.8

	pointsOfPolynomialPlot[706].X = 7.06
	pointsOfPolynomialPlot[706].Y = 317.172

	pointsOfPolynomialPlot[707].X = 7.07
	pointsOfPolynomialPlot[707].Y = 318.548

	pointsOfPolynomialPlot[708].X = 7.08
	pointsOfPolynomialPlot[708].Y = 319.928

	pointsOfPolynomialPlot[709].X = 7.09
	pointsOfPolynomialPlot[709].Y = 321.312

	pointsOfPolynomialPlot[710].X = 7.10
	pointsOfPolynomialPlot[710].Y = 322.701

	pointsOfPolynomialPlot[711].X = 7.11
	pointsOfPolynomialPlot[711].Y = 324.093

	pointsOfPolynomialPlot[712].X = 7.12
	pointsOfPolynomialPlot[712].Y = 325.489

	pointsOfPolynomialPlot[713].X = 7.13
	pointsOfPolynomialPlot[713].Y = 326.89

	pointsOfPolynomialPlot[714].X = 7.14
	pointsOfPolynomialPlot[714].Y = 328.294

	pointsOfPolynomialPlot[715].X = 7.15
	pointsOfPolynomialPlot[715].Y = 329.703

	pointsOfPolynomialPlot[716].X = 7.16
	pointsOfPolynomialPlot[716].Y = 331.116

	pointsOfPolynomialPlot[717].X = 7.17
	pointsOfPolynomialPlot[717].Y = 332.532

	pointsOfPolynomialPlot[718].X = 7.18
	pointsOfPolynomialPlot[718].Y = 333.953

	pointsOfPolynomialPlot[719].X = 7.19
	pointsOfPolynomialPlot[719].Y = 335.378

	pointsOfPolynomialPlot[720].X = 7.20
	pointsOfPolynomialPlot[720].Y = 336.808

	pointsOfPolynomialPlot[721].X = 7.21
	pointsOfPolynomialPlot[721].Y = 338.241

	pointsOfPolynomialPlot[722].X = 7.22
	pointsOfPolynomialPlot[722].Y = 339.678

	pointsOfPolynomialPlot[723].X = 7.23
	pointsOfPolynomialPlot[723].Y = 341.12

	pointsOfPolynomialPlot[724].X = 7.24
	pointsOfPolynomialPlot[724].Y = 342.565

	pointsOfPolynomialPlot[725].X = 7.25
	pointsOfPolynomialPlot[725].Y = 344.015

	pointsOfPolynomialPlot[726].X = 7.26
	pointsOfPolynomialPlot[726].Y = 345.469

	pointsOfPolynomialPlot[727].X = 7.27
	pointsOfPolynomialPlot[727].Y = 346.927

	pointsOfPolynomialPlot[728].X = 7.28
	pointsOfPolynomialPlot[728].Y = 348.39

	pointsOfPolynomialPlot[729].X = 7.29
	pointsOfPolynomialPlot[729].Y = 349.856

	pointsOfPolynomialPlot[730].X = 7.30
	pointsOfPolynomialPlot[730].Y = 351.327

	pointsOfPolynomialPlot[731].X = 7.31
	pointsOfPolynomialPlot[731].Y = 352.801

	pointsOfPolynomialPlot[732].X = 7.32
	pointsOfPolynomialPlot[732].Y = 354.28

	pointsOfPolynomialPlot[733].X = 7.33
	pointsOfPolynomialPlot[733].Y = 355.763

	pointsOfPolynomialPlot[734].X = 7.34
	pointsOfPolynomialPlot[734].Y = 357.251

	pointsOfPolynomialPlot[735].X = 7.35
	pointsOfPolynomialPlot[735].Y = 358.742

	pointsOfPolynomialPlot[736].X = 7.36
	pointsOfPolynomialPlot[736].Y = 360.238

	pointsOfPolynomialPlot[737].X = 7.37
	pointsOfPolynomialPlot[737].Y = 361.738

	pointsOfPolynomialPlot[738].X = 7.38
	pointsOfPolynomialPlot[738].Y = 363.242

	pointsOfPolynomialPlot[739].X = 7.39
	pointsOfPolynomialPlot[739].Y = 364.751

	pointsOfPolynomialPlot[740].X = 7.40
	pointsOfPolynomialPlot[740].Y = 366.264

	pointsOfPolynomialPlot[741].X = 7.41
	pointsOfPolynomialPlot[741].Y = 367.78

	pointsOfPolynomialPlot[742].X = 7.42
	pointsOfPolynomialPlot[742].Y = 369.302

	pointsOfPolynomialPlot[743].X = 7.43
	pointsOfPolynomialPlot[743].Y = 370.827

	pointsOfPolynomialPlot[744].X = 7.44
	pointsOfPolynomialPlot[744].Y = 372.357

	pointsOfPolynomialPlot[745].X = 7.45
	pointsOfPolynomialPlot[745].Y = 373.891

	pointsOfPolynomialPlot[746].X = 7.46
	pointsOfPolynomialPlot[746].Y = 375.429

	pointsOfPolynomialPlot[747].X = 7.47
	pointsOfPolynomialPlot[747].Y = 376.971

	pointsOfPolynomialPlot[748].X = 7.48
	pointsOfPolynomialPlot[748].Y = 378.518

	pointsOfPolynomialPlot[749].X = 7.49
	pointsOfPolynomialPlot[749].Y = 380.069

	pointsOfPolynomialPlot[750].X = 7.50
	pointsOfPolynomialPlot[750].Y = 381.625

	pointsOfPolynomialPlot[751].X = 7.51
	pointsOfPolynomialPlot[751].Y = 383.184

	pointsOfPolynomialPlot[752].X = 7.52
	pointsOfPolynomialPlot[752].Y = 384.748

	pointsOfPolynomialPlot[753].X = 7.53
	pointsOfPolynomialPlot[753].Y = 386.316

	pointsOfPolynomialPlot[754].X = 7.54
	pointsOfPolynomialPlot[754].Y = 387.889

	pointsOfPolynomialPlot[755].X = 7.55
	pointsOfPolynomialPlot[755].Y = 389.466

	pointsOfPolynomialPlot[756].X = 7.56
	pointsOfPolynomialPlot[756].Y = 391.047

	pointsOfPolynomialPlot[757].X = 7.57
	pointsOfPolynomialPlot[757].Y = 392.633

	pointsOfPolynomialPlot[758].X = 7.58
	pointsOfPolynomialPlot[758].Y = 394.223

	pointsOfPolynomialPlot[759].X = 7.59
	pointsOfPolynomialPlot[759].Y = 395.817

	pointsOfPolynomialPlot[760].X = 7.60
	pointsOfPolynomialPlot[760].Y = 397.416

	pointsOfPolynomialPlot[761].X = 7.61
	pointsOfPolynomialPlot[761].Y = 399.019

	pointsOfPolynomialPlot[762].X = 7.62
	pointsOfPolynomialPlot[762].Y = 400.626

	pointsOfPolynomialPlot[763].X = 7.63
	pointsOfPolynomialPlot[763].Y = 402.238

	pointsOfPolynomialPlot[764].X = 7.64
	pointsOfPolynomialPlot[764].Y = 403.854

	pointsOfPolynomialPlot[765].X = 7.65
	pointsOfPolynomialPlot[765].Y = 405.474

	pointsOfPolynomialPlot[766].X = 7.66
	pointsOfPolynomialPlot[766].Y = 407.099

	pointsOfPolynomialPlot[767].X = 7.67
	pointsOfPolynomialPlot[767].Y = 408.728

	pointsOfPolynomialPlot[768].X = 7.68
	pointsOfPolynomialPlot[768].Y = 410.362

	pointsOfPolynomialPlot[769].X = 7.69
	pointsOfPolynomialPlot[769].Y = 412.0

	pointsOfPolynomialPlot[770].X = 7.70
	pointsOfPolynomialPlot[770].Y = 413.643

	pointsOfPolynomialPlot[771].X = 7.71
	pointsOfPolynomialPlot[771].Y = 415.289

	pointsOfPolynomialPlot[772].X = 7.72
	pointsOfPolynomialPlot[772].Y = 416.941

	pointsOfPolynomialPlot[773].X = 7.73
	pointsOfPolynomialPlot[773].Y = 418.597

	pointsOfPolynomialPlot[774].X = 7.74
	pointsOfPolynomialPlot[774].Y = 420.257

	pointsOfPolynomialPlot[775].X = 7.75
	pointsOfPolynomialPlot[775].Y = 421.921

	pointsOfPolynomialPlot[776].X = 7.76
	pointsOfPolynomialPlot[776].Y = 423.591

	pointsOfPolynomialPlot[777].X = 7.77
	pointsOfPolynomialPlot[777].Y = 425.264

	pointsOfPolynomialPlot[778].X = 7.78
	pointsOfPolynomialPlot[778].Y = 426.942

	pointsOfPolynomialPlot[779].X = 7.79
	pointsOfPolynomialPlot[779].Y = 428.625

	pointsOfPolynomialPlot[780].X = 7.80
	pointsOfPolynomialPlot[780].Y = 430.312

	pointsOfPolynomialPlot[781].X = 7.81
	pointsOfPolynomialPlot[781].Y = 432.003

	pointsOfPolynomialPlot[782].X = 7.82
	pointsOfPolynomialPlot[782].Y = 433.699

	pointsOfPolynomialPlot[783].X = 7.83
	pointsOfPolynomialPlot[783].Y = 435.399

	pointsOfPolynomialPlot[784].X = 7.84
	pointsOfPolynomialPlot[784].Y = 437.104

	pointsOfPolynomialPlot[785].X = 7.85
	pointsOfPolynomialPlot[785].Y = 438.814

	pointsOfPolynomialPlot[786].X = 7.86
	pointsOfPolynomialPlot[786].Y = 440.528

	pointsOfPolynomialPlot[787].X = 7.87
	pointsOfPolynomialPlot[787].Y = 442.246

	pointsOfPolynomialPlot[788].X = 7.88
	pointsOfPolynomialPlot[788].Y = 443.969

	pointsOfPolynomialPlot[789].X = 7.89
	pointsOfPolynomialPlot[789].Y = 445.697

	pointsOfPolynomialPlot[790].X = 7.90
	pointsOfPolynomialPlot[790].Y = 447.429

	pointsOfPolynomialPlot[791].X = 7.91
	pointsOfPolynomialPlot[791].Y = 449.165

	pointsOfPolynomialPlot[792].X = 7.92
	pointsOfPolynomialPlot[792].Y = 450.906

	pointsOfPolynomialPlot[793].X = 7.93
	pointsOfPolynomialPlot[793].Y = 452.652

	pointsOfPolynomialPlot[794].X = 7.94
	pointsOfPolynomialPlot[794].Y = 454.402

	pointsOfPolynomialPlot[795].X = 7.95
	pointsOfPolynomialPlot[795].Y = 456.157

	pointsOfPolynomialPlot[796].X = 7.96
	pointsOfPolynomialPlot[796].Y = 457.916

	pointsOfPolynomialPlot[797].X = 7.97
	pointsOfPolynomialPlot[797].Y = 459.68

	pointsOfPolynomialPlot[798].X = 7.98
	pointsOfPolynomialPlot[798].Y = 461.449

	pointsOfPolynomialPlot[799].X = 7.99
	pointsOfPolynomialPlot[799].Y = 463.222

	pointsOfPolynomialPlot[800].X = 8.0
	pointsOfPolynomialPlot[800].Y = 465.0

	pointsOfPolynomialPlot[801].X = 8.01
	pointsOfPolynomialPlot[801].Y = 466.782

	pointsOfPolynomialPlot[802].X = 8.02
	pointsOfPolynomialPlot[802].Y = 468.569

	pointsOfPolynomialPlot[803].X = 8.03
	pointsOfPolynomialPlot[803].Y = 470.36

	pointsOfPolynomialPlot[804].X = 8.04
	pointsOfPolynomialPlot[804].Y = 472.156

	pointsOfPolynomialPlot[805].X = 8.05
	pointsOfPolynomialPlot[805].Y = 473.957

	pointsOfPolynomialPlot[806].X = 8.06
	pointsOfPolynomialPlot[806].Y = 475.763

	pointsOfPolynomialPlot[807].X = 8.07
	pointsOfPolynomialPlot[807].Y = 477.573

	pointsOfPolynomialPlot[808].X = 8.08
	pointsOfPolynomialPlot[808].Y = 479.387

	pointsOfPolynomialPlot[809].X = 8.09
	pointsOfPolynomialPlot[809].Y = 481.207

	pointsOfPolynomialPlot[810].X = 8.10
	pointsOfPolynomialPlot[810].Y = 483.031

	pointsOfPolynomialPlot[811].X = 8.11
	pointsOfPolynomialPlot[811].Y = 484.859

	pointsOfPolynomialPlot[812].X = 8.12
	pointsOfPolynomialPlot[812].Y = 486.859

	pointsOfPolynomialPlot[813].X = 8.13
	pointsOfPolynomialPlot[813].Y = 488.53

	pointsOfPolynomialPlot[814].X = 8.14
	pointsOfPolynomialPlot[814].Y = 490.373

	pointsOfPolynomialPlot[815].X = 8.15
	pointsOfPolynomialPlot[815].Y = 492.22

	pointsOfPolynomialPlot[816].X = 8.16
	pointsOfPolynomialPlot[816].Y = 494.072

	pointsOfPolynomialPlot[817].X = 8.17
	pointsOfPolynomialPlot[817].Y = 495.929

	pointsOfPolynomialPlot[818].X = 8.18
	pointsOfPolynomialPlot[818].Y = 497.791

	pointsOfPolynomialPlot[819].X = 8.19
	pointsOfPolynomialPlot[819].Y = 499.657

	pointsOfPolynomialPlot[820].X = 8.20
	pointsOfPolynomialPlot[820].Y = 501.528

	pointsOfPolynomialPlot[821].X = 8.21
	pointsOfPolynomialPlot[821].Y = 503.403

	pointsOfPolynomialPlot[822].X = 8.22
	pointsOfPolynomialPlot[822].Y = 505.282

	pointsOfPolynomialPlot[823].X = 8.23
	pointsOfPolynomialPlot[823].Y = 507.168

	pointsOfPolynomialPlot[824].X = 8.24
	pointsOfPolynomialPlot[824].Y = 509.058

	pointsOfPolynomialPlot[825].X = 8.25
	pointsOfPolynomialPlot[825].Y = 510.953

	pointsOfPolynomialPlot[826].X = 8.26
	pointsOfPolynomialPlot[826].Y = 512.852

	pointsOfPolynomialPlot[827].X = 8.27
	pointsOfPolynomialPlot[827].Y = 514.756

	pointsOfPolynomialPlot[828].X = 8.28
	pointsOfPolynomialPlot[828].Y = 516.665

	pointsOfPolynomialPlot[829].X = 8.29
	pointsOfPolynomialPlot[829].Y = 518.578

	pointsOfPolynomialPlot[830].X = 8.30
	pointsOfPolynomialPlot[830].Y = 520.497

	pointsOfPolynomialPlot[831].X = 8.31
	pointsOfPolynomialPlot[831].Y = 522.42

	pointsOfPolynomialPlot[832].X = 8.32
	pointsOfPolynomialPlot[832].Y = 524.349

	pointsOfPolynomialPlot[833].X = 8.33
	pointsOfPolynomialPlot[833].Y = 526.28

	pointsOfPolynomialPlot[834].X = 8.34
	pointsOfPolynomialPlot[834].Y = 528.218

	pointsOfPolynomialPlot[835].X = 8.35
	pointsOfPolynomialPlot[835].Y = 530.16

	pointsOfPolynomialPlot[836].X = 8.36
	pointsOfPolynomialPlot[836].Y = 532.107

	pointsOfPolynomialPlot[837].X = 8.37
	pointsOfPolynomialPlot[837].Y = 534.218

	pointsOfPolynomialPlot[838].X = 8.38
	pointsOfPolynomialPlot[838].Y = 536.016

	pointsOfPolynomialPlot[839].X = 8.39
	pointsOfPolynomialPlot[839].Y = 537.977

	pointsOfPolynomialPlot[840].X = 8.40
	pointsOfPolynomialPlot[840].Y = 539.944

	pointsOfPolynomialPlot[841].X = 8.41
	pointsOfPolynomialPlot[841].Y = 541.915

	pointsOfPolynomialPlot[842].X = 8.42
	pointsOfPolynomialPlot[842].Y = 543.891

	pointsOfPolynomialPlot[843].X = 8.43
	pointsOfPolynomialPlot[843].Y = 545.872

	pointsOfPolynomialPlot[844].X = 8.44
	pointsOfPolynomialPlot[844].Y = 547.858

	pointsOfPolynomialPlot[845].X = 8.45
	pointsOfPolynomialPlot[845].Y = 549.848

	pointsOfPolynomialPlot[846].X = 8.46
	pointsOfPolynomialPlot[846].Y = 551.844

	pointsOfPolynomialPlot[847].X = 8.47
	pointsOfPolynomialPlot[847].Y = 553.844

	pointsOfPolynomialPlot[848].X = 8.48
	pointsOfPolynomialPlot[848].Y = 555.849

	pointsOfPolynomialPlot[849].X = 8.49
	pointsOfPolynomialPlot[849].Y = 557.859

	pointsOfPolynomialPlot[850].X = 8.50
	pointsOfPolynomialPlot[850].Y = 559.875

	pointsOfPolynomialPlot[851].X = 8.51
	pointsOfPolynomialPlot[851].Y = 561.895

	pointsOfPolynomialPlot[852].X = 8.52
	pointsOfPolynomialPlot[852].Y = 563.919

	pointsOfPolynomialPlot[853].X = 8.53
	pointsOfPolynomialPlot[853].Y = 565.949

	pointsOfPolynomialPlot[854].X = 8.54
	pointsOfPolynomialPlot[854].Y = 567.984

	pointsOfPolynomialPlot[855].X = 8.55
	pointsOfPolynomialPlot[855].Y = 570.023

	pointsOfPolynomialPlot[856].X = 8.56
	pointsOfPolynomialPlot[856].Y = 572.068

	pointsOfPolynomialPlot[857].X = 8.57
	pointsOfPolynomialPlot[857].Y = 574.117

	pointsOfPolynomialPlot[858].X = 8.58
	pointsOfPolynomialPlot[858].Y = 576.172

	pointsOfPolynomialPlot[859].X = 8.59
	pointsOfPolynomialPlot[859].Y = 578.231

	pointsOfPolynomialPlot[860].X = 8.60
	pointsOfPolynomialPlot[860].Y = 580.296

	pointsOfPolynomialPlot[861].X = 8.61
	pointsOfPolynomialPlot[861].Y = 582.365

	pointsOfPolynomialPlot[862].X = 8.62
	pointsOfPolynomialPlot[862].Y = 584.439

	pointsOfPolynomialPlot[863].X = 8.63
	pointsOfPolynomialPlot[863].Y = 586.518

	pointsOfPolynomialPlot[864].X = 8.64
	pointsOfPolynomialPlot[864].Y = 588.602

	pointsOfPolynomialPlot[865].X = 8.65
	pointsOfPolynomialPlot[865].Y = 590.692

	pointsOfPolynomialPlot[866].X = 8.66
	pointsOfPolynomialPlot[866].Y = 592.786

	pointsOfPolynomialPlot[867].X = 8.67
	pointsOfPolynomialPlot[867].Y = 594.885

	pointsOfPolynomialPlot[868].X = 8.68
	pointsOfPolynomialPlot[868].Y = 596.989

	pointsOfPolynomialPlot[869].X = 8.69
	pointsOfPolynomialPlot[869].Y = 599.098

	pointsOfPolynomialPlot[870].X = 8.70
	pointsOfPolynomialPlot[870].Y = 601.213

	pointsOfPolynomialPlot[871].X = 8.71
	pointsOfPolynomialPlot[871].Y = 603.332

	pointsOfPolynomialPlot[872].X = 8.72
	pointsOfPolynomialPlot[872].Y = 605.456

	pointsOfPolynomialPlot[873].X = 8.73
	pointsOfPolynomialPlot[873].Y = 607.585

	pointsOfPolynomialPlot[874].X = 8.74
	pointsOfPolynomialPlot[874].Y = 609.72

	pointsOfPolynomialPlot[875].X = 8.75
	pointsOfPolynomialPlot[875].Y = 611.859

	pointsOfPolynomialPlot[876].X = 8.76
	pointsOfPolynomialPlot[876].Y = 614.003

	pointsOfPolynomialPlot[877].X = 8.77
	pointsOfPolynomialPlot[877].Y = 616.153

	pointsOfPolynomialPlot[878].X = 8.78
	pointsOfPolynomialPlot[878].Y = 618.307

	pointsOfPolynomialPlot[879].X = 8.79
	pointsOfPolynomialPlot[879].Y = 620.467

	pointsOfPolynomialPlot[880].X = 8.80
	pointsOfPolynomialPlot[880].Y = 622.632

	pointsOfPolynomialPlot[881].X = 8.81
	pointsOfPolynomialPlot[881].Y = 624.801

	pointsOfPolynomialPlot[882].X = 8.82
	pointsOfPolynomialPlot[882].Y = 626.976

	pointsOfPolynomialPlot[883].X = 8.83
	pointsOfPolynomialPlot[883].Y = 629.156

	pointsOfPolynomialPlot[884].X = 8.84
	pointsOfPolynomialPlot[884].Y = 631.341

	pointsOfPolynomialPlot[885].X = 8.85
	pointsOfPolynomialPlot[885].Y = 633.531

	pointsOfPolynomialPlot[886].X = 8.86
	pointsOfPolynomialPlot[886].Y = 635.726

	pointsOfPolynomialPlot[887].X = 8.87
	pointsOfPolynomialPlot[887].Y = 637.927

	pointsOfPolynomialPlot[888].X = 8.88
	pointsOfPolynomialPlot[888].Y = 640.132

	pointsOfPolynomialPlot[889].X = 8.89
	pointsOfPolynomialPlot[889].Y = 642.343

	pointsOfPolynomialPlot[890].X = 8.90
	pointsOfPolynomialPlot[890].Y = 644.559

	pointsOfPolynomialPlot[891].X = 8.91
	pointsOfPolynomialPlot[891].Y = 646.779

	pointsOfPolynomialPlot[892].X = 8.92
	pointsOfPolynomialPlot[892].Y = 649.005

	pointsOfPolynomialPlot[893].X = 8.93
	pointsOfPolynomialPlot[893].Y = 651.237

	pointsOfPolynomialPlot[894].X = 8.94
	pointsOfPolynomialPlot[894].Y = 653.473

	pointsOfPolynomialPlot[895].X = 8.95
	pointsOfPolynomialPlot[895].Y = 655.714

	pointsOfPolynomialPlot[896].X = 8.96
	pointsOfPolynomialPlot[896].Y = 657.961

	pointsOfPolynomialPlot[897].X = 8.97
	pointsOfPolynomialPlot[897].Y = 660.213

	pointsOfPolynomialPlot[898].X = 8.98
	pointsOfPolynomialPlot[898].Y = 662.47

	pointsOfPolynomialPlot[899].X = 8.99
	pointsOfPolynomialPlot[899].Y = 664.732

	pointsOfPolynomialPlot[900].X = 9.0
	pointsOfPolynomialPlot[900].Y = 667.0

	pointsOfPolynomialPlot[901].X = 9.01
	pointsOfPolynomialPlot[901].Y = 669.272

	pointsOfPolynomialPlot[902].X = 9.02
	pointsOfPolynomialPlot[902].Y = 671.55

	pointsOfPolynomialPlot[903].X = 9.03
	pointsOfPolynomialPlot[903].Y = 673.833

	pointsOfPolynomialPlot[904].X = 9.04
	pointsOfPolynomialPlot[904].Y = 676.121

	pointsOfPolynomialPlot[905].X = 9.05
	pointsOfPolynomialPlot[905].Y = 678.415

	pointsOfPolynomialPlot[906].X = 9.06
	pointsOfPolynomialPlot[906].Y = 680.713

	pointsOfPolynomialPlot[907].X = 9.07
	pointsOfPolynomialPlot[907].Y = 683.017

	pointsOfPolynomialPlot[908].X = 9.08
	pointsOfPolynomialPlot[908].Y = 685.326

	pointsOfPolynomialPlot[909].X = 9.09
	pointsOfPolynomialPlot[909].Y = 687.641

	pointsOfPolynomialPlot[910].X = 9.10
	pointsOfPolynomialPlot[910].Y = 689.961

	pointsOfPolynomialPlot[911].X = 9.11
	pointsOfPolynomialPlot[911].Y = 692.285

	pointsOfPolynomialPlot[912].X = 9.12
	pointsOfPolynomialPlot[912].Y = 694.616

	pointsOfPolynomialPlot[913].X = 9.13
	pointsOfPolynomialPlot[913].Y = 696.951

	pointsOfPolynomialPlot[914].X = 9.14
	pointsOfPolynomialPlot[914].Y = 699.292

	pointsOfPolynomialPlot[915].X = 9.15
	pointsOfPolynomialPlot[915].Y = 701.638

	pointsOfPolynomialPlot[916].X = 9.16
	pointsOfPolynomialPlot[916].Y = 703.989

	pointsOfPolynomialPlot[917].X = 9.17
	pointsOfPolynomialPlot[917].Y = 706.346

	pointsOfPolynomialPlot[918].X = 9.18
	pointsOfPolynomialPlot[918].Y = 708.708

	pointsOfPolynomialPlot[919].X = 9.19
	pointsOfPolynomialPlot[919].Y = 711.075

	pointsOfPolynomialPlot[920].X = 9.20
	pointsOfPolynomialPlot[920].Y = 713.448

	pointsOfPolynomialPlot[921].X = 9.21
	pointsOfPolynomialPlot[921].Y = 715.825

	pointsOfPolynomialPlot[922].X = 9.22
	pointsOfPolynomialPlot[922].Y = 718.209

	pointsOfPolynomialPlot[923].X = 9.23
	pointsOfPolynomialPlot[923].Y = 720.597

	pointsOfPolynomialPlot[924].X = 9.24
	pointsOfPolynomialPlot[924].Y = 722.991

	pointsOfPolynomialPlot[925].X = 9.25
	pointsOfPolynomialPlot[925].Y = 725.39

	pointsOfPolynomialPlot[926].X = 9.26
	pointsOfPolynomialPlot[926].Y = 727.795

	pointsOfPolynomialPlot[927].X = 9.27
	pointsOfPolynomialPlot[927].Y = 730.205

	pointsOfPolynomialPlot[928].X = 9.28
	pointsOfPolynomialPlot[928].Y = 732.62

	pointsOfPolynomialPlot[929].X = 9.29
	pointsOfPolynomialPlot[929].Y = 735.041

	pointsOfPolynomialPlot[930].X = 9.30
	pointsOfPolynomialPlot[930].Y = 737.467

	pointsOfPolynomialPlot[931].X = 9.31
	pointsOfPolynomialPlot[931].Y = 739.898

	pointsOfPolynomialPlot[932].X = 9.32
	pointsOfPolynomialPlot[932].Y = 742.335

	pointsOfPolynomialPlot[933].X = 9.33
	pointsOfPolynomialPlot[933].Y = 744.777

	pointsOfPolynomialPlot[934].X = 9.34
	pointsOfPolynomialPlot[934].Y = 747.224

	pointsOfPolynomialPlot[935].X = 9.35
	pointsOfPolynomialPlot[935].Y = 749.677

	pointsOfPolynomialPlot[936].X = 9.36
	pointsOfPolynomialPlot[936].Y = 752.136

	pointsOfPolynomialPlot[937].X = 9.37
	pointsOfPolynomialPlot[937].Y = 754.6

	pointsOfPolynomialPlot[938].X = 9.38
	pointsOfPolynomialPlot[938].Y = 757.069

	pointsOfPolynomialPlot[939].X = 9.39
	pointsOfPolynomialPlot[939].Y = 759.543

	pointsOfPolynomialPlot[940].X = 9.40
	pointsOfPolynomialPlot[940].Y = 762.024

	pointsOfPolynomialPlot[941].X = 9.41
	pointsOfPolynomialPlot[941].Y = 764.509

	pointsOfPolynomialPlot[942].X = 9.42
	pointsOfPolynomialPlot[942].Y = 767.0

	pointsOfPolynomialPlot[943].X = 9.43
	pointsOfPolynomialPlot[943].Y = 769.496

	pointsOfPolynomialPlot[944].X = 9.44
	pointsOfPolynomialPlot[944].Y = 771.998

	pointsOfPolynomialPlot[945].X = 9.45
	pointsOfPolynomialPlot[945].Y = 774.506

	pointsOfPolynomialPlot[946].X = 9.46
	pointsOfPolynomialPlot[946].Y = 777.018

	pointsOfPolynomialPlot[947].X = 9.47
	pointsOfPolynomialPlot[947].Y = 779.537

	pointsOfPolynomialPlot[948].X = 9.48
	pointsOfPolynomialPlot[948].Y = 782.061

	pointsOfPolynomialPlot[949].X = 9.49
	pointsOfPolynomialPlot[949].Y = 784.59

	pointsOfPolynomialPlot[950].X = 9.50
	pointsOfPolynomialPlot[950].Y = 787.125

	pointsOfPolynomialPlot[951].X = 9.51
	pointsOfPolynomialPlot[951].Y = 789.665

	pointsOfPolynomialPlot[952].X = 9.52
	pointsOfPolynomialPlot[952].Y = 792.211

	pointsOfPolynomialPlot[953].X = 9.53
	pointsOfPolynomialPlot[953].Y = 794.762

	pointsOfPolynomialPlot[954].X = 9.54
	pointsOfPolynomialPlot[954].Y = 797.319

	pointsOfPolynomialPlot[955].X = 9.55
	pointsOfPolynomialPlot[955].Y = 799.881

	pointsOfPolynomialPlot[956].X = 9.56
	pointsOfPolynomialPlot[956].Y = 802.449

	pointsOfPolynomialPlot[957].X = 9.57
	pointsOfPolynomialPlot[957].Y = 805.022

	pointsOfPolynomialPlot[958].X = 9.58
	pointsOfPolynomialPlot[958].Y = 807.601

	pointsOfPolynomialPlot[959].X = 9.59
	pointsOfPolynomialPlot[959].Y = 810.186

	pointsOfPolynomialPlot[960].X = 9.60
	pointsOfPolynomialPlot[960].Y = 812.776

	pointsOfPolynomialPlot[961].X = 9.61
	pointsOfPolynomialPlot[961].Y = 815.371

	pointsOfPolynomialPlot[962].X = 9.62
	pointsOfPolynomialPlot[962].Y = 817.972

	pointsOfPolynomialPlot[963].X = 9.63
	pointsOfPolynomialPlot[963].Y = 820.579

	pointsOfPolynomialPlot[964].X = 9.64
	pointsOfPolynomialPlot[964].Y = 823.191

	pointsOfPolynomialPlot[965].X = 9.65
	pointsOfPolynomialPlot[965].Y = 825.809

	pointsOfPolynomialPlot[966].X = 9.66
	pointsOfPolynomialPlot[966].Y = 828.433

	pointsOfPolynomialPlot[967].X = 9.67
	pointsOfPolynomialPlot[967].Y = 831.062

	pointsOfPolynomialPlot[968].X = 9.68
	pointsOfPolynomialPlot[968].Y = 833.696

	pointsOfPolynomialPlot[969].X = 9.69
	pointsOfPolynomialPlot[969].Y = 836.337

	pointsOfPolynomialPlot[970].X = 9.70
	pointsOfPolynomialPlot[970].Y = 838.983

	pointsOfPolynomialPlot[971].X = 9.71
	pointsOfPolynomialPlot[971].Y = 841.634

	pointsOfPolynomialPlot[972].X = 9.72
	pointsOfPolynomialPlot[972].Y = 844.291

	pointsOfPolynomialPlot[973].X = 9.73
	pointsOfPolynomialPlot[973].Y = 846.954

	pointsOfPolynomialPlot[974].X = 9.74
	pointsOfPolynomialPlot[974].Y = 849.622

	pointsOfPolynomialPlot[975].X = 9.75
	pointsOfPolynomialPlot[975].Y = 852.296

	pointsOfPolynomialPlot[976].X = 9.76
	pointsOfPolynomialPlot[976].Y = 854.976

	pointsOfPolynomialPlot[977].X = 9.77
	pointsOfPolynomialPlot[977].Y = 857.661

	pointsOfPolynomialPlot[978].X = 9.78
	pointsOfPolynomialPlot[978].Y = 860.353

	pointsOfPolynomialPlot[979].X = 9.79
	pointsOfPolynomialPlot[979].Y = 863.049

	pointsOfPolynomialPlot[980].X = 9.80
	pointsOfPolynomialPlot[980].Y = 865.752

	pointsOfPolynomialPlot[981].X = 9.81
	pointsOfPolynomialPlot[981].Y = 868.46

	pointsOfPolynomialPlot[982].X = 9.82
	pointsOfPolynomialPlot[982].Y = 871.173

	pointsOfPolynomialPlot[983].X = 9.83
	pointsOfPolynomialPlot[983].Y = 873.893

	pointsOfPolynomialPlot[984].X = 9.84
	pointsOfPolynomialPlot[984].Y = 876.618

	pointsOfPolynomialPlot[985].X = 9.85
	pointsOfPolynomialPlot[985].Y = 879.349

	pointsOfPolynomialPlot[986].X = 9.86
	pointsOfPolynomialPlot[986].Y = 882.085

	pointsOfPolynomialPlot[987].X = 9.87
	pointsOfPolynomialPlot[987].Y = 884.827

	pointsOfPolynomialPlot[988].X = 9.88
	pointsOfPolynomialPlot[988].Y = 887.575

	pointsOfPolynomialPlot[989].X = 9.89
	pointsOfPolynomialPlot[989].Y = 890.329

	pointsOfPolynomialPlot[990].X = 9.90
	pointsOfPolynomialPlot[990].Y = 893.089

	pointsOfPolynomialPlot[991].X = 9.91
	pointsOfPolynomialPlot[991].Y = 895.854

	pointsOfPolynomialPlot[992].X = 9.92
	pointsOfPolynomialPlot[992].Y = 898.625

	pointsOfPolynomialPlot[993].X = 9.93
	pointsOfPolynomialPlot[993].Y = 901.401

	pointsOfPolynomialPlot[994].X = 9.94
	pointsOfPolynomialPlot[994].Y = 904.183

	pointsOfPolynomialPlot[995].X = 9.95
	pointsOfPolynomialPlot[995].Y = 906.972

	pointsOfPolynomialPlot[996].X = 9.96
	pointsOfPolynomialPlot[996].Y = 909.766

	pointsOfPolynomialPlot[997].X = 9.97
	pointsOfPolynomialPlot[997].Y = 912.566

	pointsOfPolynomialPlot[998].X = 9.98
	pointsOfPolynomialPlot[998].Y = 915.371

	pointsOfPolynomialPlot[999].X = 9.99
	pointsOfPolynomialPlot[999].Y = 918.182

	pointsOfPolynomialPlot[1_000].X = 10.0
	pointsOfPolynomialPlot[1_000].Y = 921.0










	polynomialPlot := plot.New()

	polynomialPlot.Title.Text = "Plot of polynomial f(x) = x^3 - x^2 + 2x + 1"

	polynomialPlot.X.Label.Text = "x"
	polynomialPlot.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfPolynomialPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	polynomialPlot.Add(plotLine)
	polynomialPlot.Legend.Add("f(x)", plotLine)

	if err := polynomialPlot.Save(10*vg.Inch, 10*vg.Inch,
		"Polynomial-function-plot-06.png"); err != nil {

		panic(err)
	}
}
