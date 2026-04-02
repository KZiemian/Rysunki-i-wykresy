package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function arccos(x).

	pointsOfFunctionPlot := make(plotter.XYs, 201)

	pointsOfFunctionPlot[0].X = -1.0
	pointsOfFunctionPlot[0].Y = 3.141

	pointsOfFunctionPlot[1].X = -0.99
	pointsOfFunctionPlot[1].Y = 3.0

	pointsOfFunctionPlot[2].X = -0.98
	pointsOfFunctionPlot[2].Y = 2.941

	pointsOfFunctionPlot[3].X = -0.97
	pointsOfFunctionPlot[3].Y = 2.896

	pointsOfFunctionPlot[4].X = -0.96
	pointsOfFunctionPlot[4].Y = 2.857

	pointsOfFunctionPlot[5].X = -0.95
	pointsOfFunctionPlot[5].Y = 2.824

	pointsOfFunctionPlot[6].X = -0.94
	pointsOfFunctionPlot[6].Y = 2.793

	pointsOfFunctionPlot[7].X = -0.93
	pointsOfFunctionPlot[7].Y = 2.765

	pointsOfFunctionPlot[8].X = -0.92
	pointsOfFunctionPlot[8].Y = 2.738

	pointsOfFunctionPlot[9].X = -0.91
	pointsOfFunctionPlot[9].Y = 2.714

	pointsOfFunctionPlot[10].X = -0.90
	pointsOfFunctionPlot[10].Y = 2.69

	pointsOfFunctionPlot[11].X = -0.89
	pointsOfFunctionPlot[11].Y = 2.668

	pointsOfFunctionPlot[12].X = -0.88
	pointsOfFunctionPlot[12].Y = 2.646

	pointsOfFunctionPlot[13].X = -0.87
	pointsOfFunctionPlot[13].Y = 2.626

	pointsOfFunctionPlot[14].X = -0.86
	pointsOfFunctionPlot[14].Y = 2.606

	pointsOfFunctionPlot[15].X = -0.85
	pointsOfFunctionPlot[15].Y = 2.586

	pointsOfFunctionPlot[16].X = -0.84
	pointsOfFunctionPlot[16].Y = 2.568

	pointsOfFunctionPlot[17].X = -0.83
	pointsOfFunctionPlot[17].Y = 2.549

	pointsOfFunctionPlot[18].X = -0.82
	pointsOfFunctionPlot[18].Y = 2.532

	pointsOfFunctionPlot[19].X = -0.81
	pointsOfFunctionPlot[19].Y = 2.514

	pointsOfFunctionPlot[20].X = -0.80
	pointsOfFunctionPlot[20].Y = 2.498

	pointsOfFunctionPlot[21].X = -0.79
	pointsOfFunctionPlot[21].Y = 2.481

	pointsOfFunctionPlot[22].X = -0.78
	pointsOfFunctionPlot[22].Y = 2.465

	pointsOfFunctionPlot[23].X = -0.77
	pointsOfFunctionPlot[23].Y = 2.449

	pointsOfFunctionPlot[24].X = -0.76
	pointsOfFunctionPlot[24].Y = 2.434

	pointsOfFunctionPlot[25].X = -0.75
	pointsOfFunctionPlot[25].Y = 2.418

	pointsOfFunctionPlot[26].X = -0.74
	pointsOfFunctionPlot[26].Y = 2.403

	pointsOfFunctionPlot[27].X = -0.73
	pointsOfFunctionPlot[27].Y = 2.389

	pointsOfFunctionPlot[28].X = -0.72
	pointsOfFunctionPlot[28].Y = 2.374

	pointsOfFunctionPlot[29].X = -0.71
	pointsOfFunctionPlot[29].Y = 2.36

	pointsOfFunctionPlot[30].X = -0.70
	pointsOfFunctionPlot[30].Y = 2.346

	pointsOfFunctionPlot[31].X = -0.69
	pointsOfFunctionPlot[31].Y = 2.332

	pointsOfFunctionPlot[32].X = -0.68
	pointsOfFunctionPlot[32].Y = 2.318

	pointsOfFunctionPlot[33].X = -0.67
	pointsOfFunctionPlot[33].Y = 2.305

	pointsOfFunctionPlot[34].X = -0.66
	pointsOfFunctionPlot[34].Y = 2.291

	pointsOfFunctionPlot[35].X = -0.65
	pointsOfFunctionPlot[35].Y = 2.278

	pointsOfFunctionPlot[36].X = -0.64
	pointsOfFunctionPlot[36].Y = 2.265

	pointsOfFunctionPlot[37].X = -0.63
	pointsOfFunctionPlot[37].Y = 2.252

	pointsOfFunctionPlot[38].X = -0.62
	pointsOfFunctionPlot[38].Y = 2.239

	pointsOfFunctionPlot[39].X = -0.61
	pointsOfFunctionPlot[39].Y = 2.226

	pointsOfFunctionPlot[40].X = -0.60
	pointsOfFunctionPlot[40].Y = 2.214

	pointsOfFunctionPlot[41].X = -0.59
	pointsOfFunctionPlot[41].Y = 2.201

	pointsOfFunctionPlot[42].X = -0.58
	pointsOfFunctionPlot[42].Y = 2.189

	pointsOfFunctionPlot[43].X = -0.57
	pointsOfFunctionPlot[43].Y = 2.177

	pointsOfFunctionPlot[44].X = -0.56
	pointsOfFunctionPlot[44].Y = 2.165

	pointsOfFunctionPlot[45].X = -0.55
	pointsOfFunctionPlot[45].Y = 2.153

	pointsOfFunctionPlot[46].X = -0.54
	pointsOfFunctionPlot[46].Y = 2.141

	pointsOfFunctionPlot[47].X = -0.53
	pointsOfFunctionPlot[47].Y = 2.129

	pointsOfFunctionPlot[48].X = -0.52
	pointsOfFunctionPlot[48].Y = 2.117

	pointsOfFunctionPlot[49].X = -0.51
	pointsOfFunctionPlot[49].Y = 2.106

	pointsOfFunctionPlot[50].X = -0.50
	pointsOfFunctionPlot[50].Y = 2.094

	pointsOfFunctionPlot[51].X = -0.49
	pointsOfFunctionPlot[51].Y = 2.082

	pointsOfFunctionPlot[52].X = -0.48
	pointsOfFunctionPlot[52].Y = 2.071

	pointsOfFunctionPlot[53].X = -0.47
	pointsOfFunctionPlot[53].Y = 2.06

	pointsOfFunctionPlot[54].X = -0.46
	pointsOfFunctionPlot[54].Y = 2.048

	pointsOfFunctionPlot[55].X = -0.45
	pointsOfFunctionPlot[55].Y = 2.037

	pointsOfFunctionPlot[56].X = -0.44
	pointsOfFunctionPlot[56].Y = 2.026

	pointsOfFunctionPlot[57].X = -0.43
	pointsOfFunctionPlot[57].Y = 2.015

	pointsOfFunctionPlot[58].X = -0.42
	pointsOfFunctionPlot[58].Y = 2.004

	pointsOfFunctionPlot[59].X = -0.41
	pointsOfFunctionPlot[59].Y = 1.993

	pointsOfFunctionPlot[60].X = -0.40
	pointsOfFunctionPlot[60].Y = 1.982

	pointsOfFunctionPlot[61].X = -0.39
	pointsOfFunctionPlot[61].Y = 1.971

	pointsOfFunctionPlot[62].X = -0.38
	pointsOfFunctionPlot[62].Y = 1.96

	pointsOfFunctionPlot[63].X = -0.37
	pointsOfFunctionPlot[63].Y = 1.949

	pointsOfFunctionPlot[64].X = -0.36
	pointsOfFunctionPlot[64].Y = 1.939

	pointsOfFunctionPlot[65].X = -0.35
	pointsOfFunctionPlot[65].Y = 1.928

	pointsOfFunctionPlot[66].X = -0.34
	pointsOfFunctionPlot[66].Y = 1.917

	pointsOfFunctionPlot[67].X = -0.33
	pointsOfFunctionPlot[67].Y = 1.907

	pointsOfFunctionPlot[68].X = -0.32
	pointsOfFunctionPlot[68].Y = 1.896

	pointsOfFunctionPlot[69].X = -0.31
	pointsOfFunctionPlot[69].Y = 1.886

	pointsOfFunctionPlot[70].X = -0.30
	pointsOfFunctionPlot[70].Y = 1.875

	pointsOfFunctionPlot[71].X = -0.29
	pointsOfFunctionPlot[71].Y = 1.865

	pointsOfFunctionPlot[72].X = -0.28
	pointsOfFunctionPlot[72].Y = 1.854

	pointsOfFunctionPlot[73].X = -0.27
	pointsOfFunctionPlot[73].Y = 1.844

	pointsOfFunctionPlot[74].X = -0.26
	pointsOfFunctionPlot[74].Y = 1.833

	pointsOfFunctionPlot[75].X = -0.25
	pointsOfFunctionPlot[75].Y = 1.823

	pointsOfFunctionPlot[76].X = -0.24
	pointsOfFunctionPlot[76].Y = 1.813

	pointsOfFunctionPlot[77].X = -0.23
	pointsOfFunctionPlot[77].Y = 1.802

	pointsOfFunctionPlot[78].X = -0.22
	pointsOfFunctionPlot[78].Y = 1.792

	pointsOfFunctionPlot[79].X = -0.21
	pointsOfFunctionPlot[79].Y = 1.782

	pointsOfFunctionPlot[80].X = -0.20
	pointsOfFunctionPlot[80].Y = 1.772

	pointsOfFunctionPlot[81].X = -0.19
	pointsOfFunctionPlot[81].Y = 1.762

	pointsOfFunctionPlot[82].X = -0.18
	pointsOfFunctionPlot[82].Y = 1.751

	pointsOfFunctionPlot[83].X = -0.17
	pointsOfFunctionPlot[83].Y = 1.741

	pointsOfFunctionPlot[84].X = -0.16
	pointsOfFunctionPlot[84].Y = 1.731

	pointsOfFunctionPlot[85].X = -0.15
	pointsOfFunctionPlot[85].Y = 1.721

	pointsOfFunctionPlot[86].X = -0.14
	pointsOfFunctionPlot[86].Y = 1.711

	pointsOfFunctionPlot[87].X = -0.13
	pointsOfFunctionPlot[87].Y = 1.701

	pointsOfFunctionPlot[88].X = -0.12
	pointsOfFunctionPlot[88].Y = 1.691

	pointsOfFunctionPlot[89].X = -0.11
	pointsOfFunctionPlot[89].Y = 1.681

	pointsOfFunctionPlot[90].X = -0.10
	pointsOfFunctionPlot[90].Y = 1.671

	pointsOfFunctionPlot[91].X = -0.09
	pointsOfFunctionPlot[91].Y = 1.66

	pointsOfFunctionPlot[92].X = -0.08
	pointsOfFunctionPlot[92].Y = 1.65

	pointsOfFunctionPlot[93].X = -0.07
	pointsOfFunctionPlot[93].Y = 1.64

	pointsOfFunctionPlot[94].X = -0.06
	pointsOfFunctionPlot[94].Y = 1.63

	pointsOfFunctionPlot[95].X = -0.05
	pointsOfFunctionPlot[95].Y = 1.62

	pointsOfFunctionPlot[96].X = -0.04
	pointsOfFunctionPlot[96].Y = 1.61

	pointsOfFunctionPlot[97].X = -0.03
	pointsOfFunctionPlot[97].Y = 1.6

	pointsOfFunctionPlot[98].X = -0.02
	pointsOfFunctionPlot[98].Y = 1.59

	pointsOfFunctionPlot[99].X = -0.01
	pointsOfFunctionPlot[99].Y = 1.58

	pointsOfFunctionPlot[100].X = 0.0
	pointsOfFunctionPlot[100].Y = 1.57

	pointsOfFunctionPlot[101].X = 0.01
	pointsOfFunctionPlot[101].Y = 1.56

	pointsOfFunctionPlot[102].X = 0.02
	pointsOfFunctionPlot[102].Y = 1.55

	pointsOfFunctionPlot[103].X = 0.03
	pointsOfFunctionPlot[103].Y = 1.54

	pointsOfFunctionPlot[104].X = 0.04
	pointsOfFunctionPlot[104].Y = 1.53

	pointsOfFunctionPlot[105].X = 0.05
	pointsOfFunctionPlot[105].Y = 1.52

	pointsOfFunctionPlot[106].X = 0.06
	pointsOfFunctionPlot[106].Y = 1.51

	pointsOfFunctionPlot[107].X = 0.07
	pointsOfFunctionPlot[107].Y = 1.5

	pointsOfFunctionPlot[108].X = 0.08
	pointsOfFunctionPlot[108].Y = 1.49

	pointsOfFunctionPlot[109].X = 0.09
	pointsOfFunctionPlot[109].Y = 1.48

	pointsOfFunctionPlot[110].X = 0.10
	pointsOfFunctionPlot[110].Y = 1.47

	pointsOfFunctionPlot[111].X = 0.11
	pointsOfFunctionPlot[111].Y = 1.46

	pointsOfFunctionPlot[112].X = 0.12
	pointsOfFunctionPlot[112].Y = 1.45

	pointsOfFunctionPlot[113].X = 0.13
	pointsOfFunctionPlot[113].Y = 1.44

	pointsOfFunctionPlot[114].X = 0.14
	pointsOfFunctionPlot[114].Y = 1.43

	pointsOfFunctionPlot[115].X = 0.15
	pointsOfFunctionPlot[115].Y = 1.42

	pointsOfFunctionPlot[116].X = 0.16
	pointsOfFunctionPlot[116].Y = 1.41

	pointsOfFunctionPlot[117].X = 0.17
	pointsOfFunctionPlot[117].Y = 1.4

	pointsOfFunctionPlot[118].X = 0.18
	pointsOfFunctionPlot[118].Y = 1.389

	pointsOfFunctionPlot[119].X = 0.19
	pointsOfFunctionPlot[119].Y = 1.379

	pointsOfFunctionPlot[120].X = 0.20
	pointsOfFunctionPlot[120].Y = 1.369

	pointsOfFunctionPlot[121].X = 0.21
	pointsOfFunctionPlot[121].Y = 1.359

	pointsOfFunctionPlot[122].X = 0.22
	pointsOfFunctionPlot[122].Y = 1.349

	pointsOfFunctionPlot[123].X = 0.23
	pointsOfFunctionPlot[123].Y = 1.338

	pointsOfFunctionPlot[124].X = 0.24
	pointsOfFunctionPlot[124].Y = 1.328

	pointsOfFunctionPlot[125].X = 0.25
	pointsOfFunctionPlot[125].Y = 1.318

	pointsOfFunctionPlot[126].X = 0.26
	pointsOfFunctionPlot[126].Y = 1.307

	pointsOfFunctionPlot[127].X = 0.27
	pointsOfFunctionPlot[127].Y = 1.297

	pointsOfFunctionPlot[128].X = 0.28
	pointsOfFunctionPlot[128].Y = 1.287

	pointsOfFunctionPlot[129].X = 0.29
	pointsOfFunctionPlot[129].Y = 1.276

	pointsOfFunctionPlot[130].X = 0.30
	pointsOfFunctionPlot[130].Y = 1.266

	pointsOfFunctionPlot[131].X = 0.31
	pointsOfFunctionPlot[131].Y = 1.255

	pointsOfFunctionPlot[132].X = 0.32
	pointsOfFunctionPlot[132].Y = 1.245

	pointsOfFunctionPlot[133].X = 0.33
	pointsOfFunctionPlot[133].Y = 1.234

	pointsOfFunctionPlot[134].X = 0.34
	pointsOfFunctionPlot[134].Y = 1.223

	pointsOfFunctionPlot[135].X = 0.35
	pointsOfFunctionPlot[135].Y = 1.213

	pointsOfFunctionPlot[136].X = 0.36
	pointsOfFunctionPlot[136].Y = 1.202

	pointsOfFunctionPlot[137].X = 0.37
	pointsOfFunctionPlot[137].Y = 1.191

	pointsOfFunctionPlot[138].X = 0.38
	pointsOfFunctionPlot[138].Y = 1.181

	pointsOfFunctionPlot[139].X = 0.39
	pointsOfFunctionPlot[139].Y = 1.17

	pointsOfFunctionPlot[140].X = 0.40
	pointsOfFunctionPlot[140].Y = 1.159

	pointsOfFunctionPlot[141].X = 0.41
	pointsOfFunctionPlot[141].Y = 1.148

	pointsOfFunctionPlot[142].X = 0.42
	pointsOfFunctionPlot[142].Y = 1.137

	pointsOfFunctionPlot[143].X = 0.43
	pointsOfFunctionPlot[143].Y = 1.126

	pointsOfFunctionPlot[144].X = 0.44
	pointsOfFunctionPlot[144].Y = 1.115

	pointsOfFunctionPlot[145].X = 0.45
	pointsOfFunctionPlot[145].Y = 1.104

	pointsOfFunctionPlot[146].X = 0.46
	pointsOfFunctionPlot[146].Y = 1.092

	pointsOfFunctionPlot[147].X = 0.47
	pointsOfFunctionPlot[147].Y = 1.081

	pointsOfFunctionPlot[148].X = 0.48
	pointsOfFunctionPlot[148].Y = 1.07

	pointsOfFunctionPlot[149].X = 0.49
	pointsOfFunctionPlot[149].Y = 1.058

	pointsOfFunctionPlot[150].X = 0.50
	pointsOfFunctionPlot[150].Y = 1.047

	pointsOfFunctionPlot[151].X = 0.51
	pointsOfFunctionPlot[151].Y = 1.035

	pointsOfFunctionPlot[152].X = 0.52
	pointsOfFunctionPlot[152].Y = 1.023

	pointsOfFunctionPlot[153].X = 0.53
	pointsOfFunctionPlot[153].Y = 1.012

	pointsOfFunctionPlot[154].X = 0.54
	pointsOfFunctionPlot[154].Y = 1.0

	pointsOfFunctionPlot[155].X = 0.55
	pointsOfFunctionPlot[155].Y = 0.988

	pointsOfFunctionPlot[156].X = 0.56
	pointsOfFunctionPlot[156].Y = 0.976

	pointsOfFunctionPlot[157].X = 0.57
	pointsOfFunctionPlot[157].Y = 0.964

	pointsOfFunctionPlot[158].X = 0.58
	pointsOfFunctionPlot[158].Y = 0.952

	pointsOfFunctionPlot[159].X = 0.59
	pointsOfFunctionPlot[159].Y = 0.939

	pointsOfFunctionPlot[160].X = 0.60
	pointsOfFunctionPlot[160].Y = 0.927

	pointsOfFunctionPlot[161].X = 0.61
	pointsOfFunctionPlot[161].Y = 0.914

	pointsOfFunctionPlot[162].X = 0.62
	pointsOfFunctionPlot[162].Y = 0.902

	pointsOfFunctionPlot[163].X = 0.63
	pointsOfFunctionPlot[163].Y = 0.889

	pointsOfFunctionPlot[164].X = 0.64
	pointsOfFunctionPlot[164].Y = 0.876

	pointsOfFunctionPlot[165].X = 0.65
	pointsOfFunctionPlot[165].Y = 0.863

	pointsOfFunctionPlot[166].X = 0.66
	pointsOfFunctionPlot[166].Y = 0.85

	pointsOfFunctionPlot[167].X = 0.67
	pointsOfFunctionPlot[167].Y = 0.836

	pointsOfFunctionPlot[168].X = 0.68
	pointsOfFunctionPlot[168].Y = 0.823

	pointsOfFunctionPlot[169].X = 0.69
	pointsOfFunctionPlot[169].Y = 0.809

	pointsOfFunctionPlot[170].X = 0.70
	pointsOfFunctionPlot[170].Y = 0.795

	pointsOfFunctionPlot[171].X = 0.71
	pointsOfFunctionPlot[171].Y = 0.781

	pointsOfFunctionPlot[172].X = 0.72
	pointsOfFunctionPlot[172].Y = 0.767

	pointsOfFunctionPlot[173].X = 0.73
	pointsOfFunctionPlot[173].Y = 0.752

	pointsOfFunctionPlot[174].X = 0.74
	pointsOfFunctionPlot[174].Y = 0.737

	pointsOfFunctionPlot[175].X = 0.75
	pointsOfFunctionPlot[175].Y = 0.722

	pointsOfFunctionPlot[176].X = 0.76
	pointsOfFunctionPlot[176].Y = 0.707

	pointsOfFunctionPlot[177].X = 0.77
	pointsOfFunctionPlot[177].Y = 0.692

	pointsOfFunctionPlot[178].X = 0.78
	pointsOfFunctionPlot[178].Y = 0.676

	pointsOfFunctionPlot[179].X = 0.79
	pointsOfFunctionPlot[179].Y = 0.66

	pointsOfFunctionPlot[180].X = 0.80
	pointsOfFunctionPlot[180].Y = 0.643

	pointsOfFunctionPlot[181].X = 0.81
	pointsOfFunctionPlot[181].Y = 0.626

	pointsOfFunctionPlot[182].X = 0.82
	pointsOfFunctionPlot[182].Y = 0.609

	pointsOfFunctionPlot[183].X = 0.83
	pointsOfFunctionPlot[183].Y = 0.591

	pointsOfFunctionPlot[184].X = 0.84
	pointsOfFunctionPlot[184].Y = 0.573

	pointsOfFunctionPlot[185].X = 0.85
	pointsOfFunctionPlot[185].Y = 0.554

	pointsOfFunctionPlot[186].X = 0.86
	pointsOfFunctionPlot[186].Y = 0.535

	pointsOfFunctionPlot[187].X = 0.87
	pointsOfFunctionPlot[187].Y = 0.515

	pointsOfFunctionPlot[188].X = 0.88
	pointsOfFunctionPlot[188].Y = 0.494

	pointsOfFunctionPlot[189].X = 0.89
	pointsOfFunctionPlot[189].Y = 0.473

	pointsOfFunctionPlot[190].X = 0.90
	pointsOfFunctionPlot[190].Y = 0.451

	pointsOfFunctionPlot[191].X = 0.91
	pointsOfFunctionPlot[191].Y = 0.427

	pointsOfFunctionPlot[192].X = 0.92
	pointsOfFunctionPlot[192].Y = 0.402

	pointsOfFunctionPlot[193].X = 0.93
	pointsOfFunctionPlot[193].Y = 0.376

	pointsOfFunctionPlot[194].X = 0.94
	pointsOfFunctionPlot[194].Y = 0.348

	pointsOfFunctionPlot[195].X = 0.95
	pointsOfFunctionPlot[195].Y = 0.317

	pointsOfFunctionPlot[196].X = 0.96
	pointsOfFunctionPlot[196].Y = 0.283

	pointsOfFunctionPlot[197].X = 0.97
	pointsOfFunctionPlot[197].Y = 0.245

	pointsOfFunctionPlot[198].X = 0.98
	pointsOfFunctionPlot[198].Y = 0.2

	pointsOfFunctionPlot[199].X = 0.99
	pointsOfFunctionPlot[199].Y = 0.141

	pointsOfFunctionPlot[200].X = 1.0
	pointsOfFunctionPlot[200].Y = 0.0










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function arccos(x)"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("arccos(x)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"arccos-function-plot-01.png"); err != nil {

		panic(err)
	}
}
