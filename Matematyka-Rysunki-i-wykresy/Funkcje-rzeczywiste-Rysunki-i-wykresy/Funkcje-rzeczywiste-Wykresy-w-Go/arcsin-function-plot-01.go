package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function arcsin(x).

	pointsOfFunctionPlot := make(plotter.XYs, 201)

	pointsOfFunctionPlot[0].X = -1.0
	pointsOfFunctionPlot[0].Y = -1.57

	pointsOfFunctionPlot[1].X = -0.99
	pointsOfFunctionPlot[1].Y = -1.429

	pointsOfFunctionPlot[2].X = -0.98
	pointsOfFunctionPlot[2].Y = -1.37

	pointsOfFunctionPlot[3].X = -0.97
	pointsOfFunctionPlot[3].Y = -1.325

	pointsOfFunctionPlot[4].X = -0.96
	pointsOfFunctionPlot[4].Y = -1.287

	pointsOfFunctionPlot[5].X = -0.95
	pointsOfFunctionPlot[5].Y = -1.253

	pointsOfFunctionPlot[6].X = -0.94
	pointsOfFunctionPlot[6].Y = -1.222

	pointsOfFunctionPlot[7].X = -0.93
	pointsOfFunctionPlot[7].Y = -1.194

	pointsOfFunctionPlot[8].X = -0.92
	pointsOfFunctionPlot[8].Y = -1.168

	pointsOfFunctionPlot[9].X = -0.91
	pointsOfFunctionPlot[9].Y = -1.143

	pointsOfFunctionPlot[10].X = -0.90
	pointsOfFunctionPlot[10].Y = -1.119

	pointsOfFunctionPlot[11].X = -0.89
	pointsOfFunctionPlot[11].Y = -1.097

	pointsOfFunctionPlot[12].X = -0.88
	pointsOfFunctionPlot[12].Y = -1.075

	pointsOfFunctionPlot[13].X = -0.87
	pointsOfFunctionPlot[13].Y = -1.055

	pointsOfFunctionPlot[14].X = -0.86
	pointsOfFunctionPlot[14].Y = -1.035

	pointsOfFunctionPlot[15].X = -0.85
	pointsOfFunctionPlot[15].Y = -1.016

	pointsOfFunctionPlot[16].X = -0.84
	pointsOfFunctionPlot[16].Y = -0.997

	pointsOfFunctionPlot[17].X = -0.83
	pointsOfFunctionPlot[17].Y = -0.979

	pointsOfFunctionPlot[18].X = -0.82
	pointsOfFunctionPlot[18].Y = -0.961

	pointsOfFunctionPlot[19].X = -0.81
	pointsOfFunctionPlot[19].Y = -0.944

	pointsOfFunctionPlot[20].X = -0.80
	pointsOfFunctionPlot[20].Y = -0.927

	pointsOfFunctionPlot[21].X = -0.79
	pointsOfFunctionPlot[21].Y = -0.91

	pointsOfFunctionPlot[22].X = -0.78
	pointsOfFunctionPlot[22].Y = -0.894

	pointsOfFunctionPlot[23].X = -0.77
	pointsOfFunctionPlot[23].Y = -0.878

	pointsOfFunctionPlot[24].X = -0.76
	pointsOfFunctionPlot[24].Y = -0.863

	pointsOfFunctionPlot[25].X = -0.75
	pointsOfFunctionPlot[25].Y = -0.848

	pointsOfFunctionPlot[26].X = -0.74
	pointsOfFunctionPlot[26].Y = -0.833

	pointsOfFunctionPlot[27].X = -0.73
	pointsOfFunctionPlot[27].Y = -0.818

	pointsOfFunctionPlot[28].X = -0.72
	pointsOfFunctionPlot[28].Y = -0.803

	pointsOfFunctionPlot[29].X = -0.71
	pointsOfFunctionPlot[29].Y = -0.789

	pointsOfFunctionPlot[30].X = -0.70
	pointsOfFunctionPlot[30].Y = -0.775

	pointsOfFunctionPlot[31].X = -0.69
	pointsOfFunctionPlot[31].Y = -0.761

	pointsOfFunctionPlot[32].X = -0.68
	pointsOfFunctionPlot[32].Y = -0.747

	pointsOfFunctionPlot[33].X = -0.67
	pointsOfFunctionPlot[33].Y = -0.734

	pointsOfFunctionPlot[34].X = -0.66
	pointsOfFunctionPlot[34].Y = -0.72

	pointsOfFunctionPlot[35].X = -0.65
	pointsOfFunctionPlot[35].Y = -0.707

	pointsOfFunctionPlot[36].X = -0.64
	pointsOfFunctionPlot[36].Y = -0.694

	pointsOfFunctionPlot[37].X = -0.63
	pointsOfFunctionPlot[37].Y = -0.681

	pointsOfFunctionPlot[38].X = -0.62
	pointsOfFunctionPlot[38].Y = -0.668

	pointsOfFunctionPlot[39].X = -0.61
	pointsOfFunctionPlot[39].Y = -0.656

	pointsOfFunctionPlot[40].X = -0.60
	pointsOfFunctionPlot[40].Y = -0.643

	pointsOfFunctionPlot[41].X = -0.59
	pointsOfFunctionPlot[41].Y = -0.631

	pointsOfFunctionPlot[42].X = -0.58
	pointsOfFunctionPlot[42].Y = -0.618

	pointsOfFunctionPlot[43].X = -0.57
	pointsOfFunctionPlot[43].Y = -0.606

	pointsOfFunctionPlot[44].X = -0.56
	pointsOfFunctionPlot[44].Y = -0.594

	pointsOfFunctionPlot[45].X = -0.55
	pointsOfFunctionPlot[45].Y = -0.582

	pointsOfFunctionPlot[46].X = -0.54
	pointsOfFunctionPlot[46].Y = -0.57

	pointsOfFunctionPlot[47].X = -0.53
	pointsOfFunctionPlot[47].Y = -0.558

	pointsOfFunctionPlot[48].X = -0.52
	pointsOfFunctionPlot[48].Y = -0.546

	pointsOfFunctionPlot[49].X = -0.51
	pointsOfFunctionPlot[49].Y = -0.535

	pointsOfFunctionPlot[50].X = -0.50
	pointsOfFunctionPlot[50].Y = -0.523

	pointsOfFunctionPlot[51].X = -0.49
	pointsOfFunctionPlot[51].Y = -0.512

	pointsOfFunctionPlot[52].X = -0.48
	pointsOfFunctionPlot[52].Y = -0.5

	pointsOfFunctionPlot[53].X = -0.47
	pointsOfFunctionPlot[53].Y = -0.489

	pointsOfFunctionPlot[54].X = -0.46
	pointsOfFunctionPlot[54].Y = -0.478

	pointsOfFunctionPlot[55].X = -0.45
	pointsOfFunctionPlot[55].Y = -0.466

	pointsOfFunctionPlot[56].X = -0.44
	pointsOfFunctionPlot[56].Y = -0.455

	pointsOfFunctionPlot[57].X = -0.43
	pointsOfFunctionPlot[57].Y = -0.444

	pointsOfFunctionPlot[58].X = -0.42
	pointsOfFunctionPlot[58].Y = -0.433

	pointsOfFunctionPlot[59].X = -0.41
	pointsOfFunctionPlot[59].Y = -0.422

	pointsOfFunctionPlot[60].X = -0.40
	pointsOfFunctionPlot[60].Y = -0.4

	pointsOfFunctionPlot[61].X = -0.39
	pointsOfFunctionPlot[61].Y = -0.4

	pointsOfFunctionPlot[62].X = -0.38
	pointsOfFunctionPlot[62].Y = -0.389

	pointsOfFunctionPlot[63].X = -0.37
	pointsOfFunctionPlot[63].Y = -0.379

	pointsOfFunctionPlot[64].X = -0.36
	pointsOfFunctionPlot[64].Y = -0.368

	pointsOfFunctionPlot[65].X = -0.35
	pointsOfFunctionPlot[65].Y = -0.357

	pointsOfFunctionPlot[66].X = -0.34
	pointsOfFunctionPlot[66].Y = -0.346

	pointsOfFunctionPlot[67].X = -0.33
	pointsOfFunctionPlot[67].Y = -0.336

	pointsOfFunctionPlot[68].X = -0.32
	pointsOfFunctionPlot[68].Y = -0.325

	pointsOfFunctionPlot[69].X = -0.31
	pointsOfFunctionPlot[69].Y = -0.315

	pointsOfFunctionPlot[70].X = -0.30
	pointsOfFunctionPlot[70].Y = -0.304

	pointsOfFunctionPlot[71].X = -0.29
	pointsOfFunctionPlot[71].Y = -0.294

	pointsOfFunctionPlot[72].X = -0.28
	pointsOfFunctionPlot[72].Y = -0.283

	pointsOfFunctionPlot[73].X = -0.27
	pointsOfFunctionPlot[73].Y = -0.273

	pointsOfFunctionPlot[74].X = -0.26
	pointsOfFunctionPlot[74].Y = -0.263

	pointsOfFunctionPlot[75].X = -0.25
	pointsOfFunctionPlot[75].Y = -0.252

	pointsOfFunctionPlot[76].X = -0.24
	pointsOfFunctionPlot[76].Y = -0.242

	pointsOfFunctionPlot[77].X = -0.23
	pointsOfFunctionPlot[77].Y = -0.232

	pointsOfFunctionPlot[78].X = -0.22
	pointsOfFunctionPlot[78].Y = -0.221

	pointsOfFunctionPlot[79].X = -0.21
	pointsOfFunctionPlot[79].Y = -0.211

	pointsOfFunctionPlot[80].X = -0.20
	pointsOfFunctionPlot[80].Y = -0.201

	pointsOfFunctionPlot[81].X = -0.19
	pointsOfFunctionPlot[81].Y = -0.191

	pointsOfFunctionPlot[82].X = -0.18
	pointsOfFunctionPlot[82].Y = -0.181

	pointsOfFunctionPlot[83].X = -0.17
	pointsOfFunctionPlot[83].Y = -0.17

	pointsOfFunctionPlot[84].X = -0.16
	pointsOfFunctionPlot[84].Y = -0.16

	pointsOfFunctionPlot[85].X = -0.15
	pointsOfFunctionPlot[85].Y = -0.15

	pointsOfFunctionPlot[86].X = -0.14
	pointsOfFunctionPlot[86].Y = -0.14

	pointsOfFunctionPlot[87].X = -0.13
	pointsOfFunctionPlot[87].Y = -0.13

	pointsOfFunctionPlot[88].X = -0.12
	pointsOfFunctionPlot[88].Y = -0.12

	pointsOfFunctionPlot[89].X = -0.11
	pointsOfFunctionPlot[89].Y = -0.11

	pointsOfFunctionPlot[90].X = -0.10
	pointsOfFunctionPlot[90].Y = -0.1

	pointsOfFunctionPlot[91].X = -0.09
	pointsOfFunctionPlot[91].Y = -0.09

	pointsOfFunctionPlot[92].X = -0.08
	pointsOfFunctionPlot[92].Y = -0.08

	pointsOfFunctionPlot[93].X = -0.07
	pointsOfFunctionPlot[93].Y = -0.07

	pointsOfFunctionPlot[94].X = -0.06
	pointsOfFunctionPlot[94].Y = -0.06

	pointsOfFunctionPlot[95].X = -0.05
	pointsOfFunctionPlot[95].Y = -0.05

	pointsOfFunctionPlot[96].X = -0.04
	pointsOfFunctionPlot[96].Y = -0.04

	pointsOfFunctionPlot[97].X = -0.03
	pointsOfFunctionPlot[97].Y = -0.03

	pointsOfFunctionPlot[98].X = -0.02
	pointsOfFunctionPlot[98].Y = -0.02

	pointsOfFunctionPlot[99].X = -0.01
	pointsOfFunctionPlot[99].Y = -0.01

	pointsOfFunctionPlot[100].X = 0.0
	pointsOfFunctionPlot[100].Y = 0.0

	pointsOfFunctionPlot[101].X = 0.01
	pointsOfFunctionPlot[101].Y = 0.01

	pointsOfFunctionPlot[102].X = 0.02
	pointsOfFunctionPlot[102].Y = 0.02

	pointsOfFunctionPlot[103].X = 0.03
	pointsOfFunctionPlot[103].Y = 0.03

	pointsOfFunctionPlot[104].X = 0.04
	pointsOfFunctionPlot[104].Y = 0.04

	pointsOfFunctionPlot[105].X = 0.05
	pointsOfFunctionPlot[105].Y = 0.05

	pointsOfFunctionPlot[106].X = 0.06
	pointsOfFunctionPlot[106].Y = 0.06

	pointsOfFunctionPlot[107].X = 0.07
	pointsOfFunctionPlot[107].Y = 0.07

	pointsOfFunctionPlot[108].X = 0.08
	pointsOfFunctionPlot[108].Y = 0.08

	pointsOfFunctionPlot[109].X = 0.09
	pointsOfFunctionPlot[109].Y = 0.09

	pointsOfFunctionPlot[110].X = 0.10
	pointsOfFunctionPlot[110].Y = 0.1

	pointsOfFunctionPlot[111].X = 0.11
	pointsOfFunctionPlot[111].Y = 0.11

	pointsOfFunctionPlot[112].X = 0.12
	pointsOfFunctionPlot[112].Y = 0.12

	pointsOfFunctionPlot[113].X = 0.13
	pointsOfFunctionPlot[113].Y = 0.13

	pointsOfFunctionPlot[114].X = 0.14
	pointsOfFunctionPlot[114].Y = 0.14

	pointsOfFunctionPlot[115].X = 0.15
	pointsOfFunctionPlot[115].Y = 0.15

	pointsOfFunctionPlot[116].X = 0.16
	pointsOfFunctionPlot[116].Y = 0.16

	pointsOfFunctionPlot[117].X = 0.17
	pointsOfFunctionPlot[117].Y = 0.17

	pointsOfFunctionPlot[118].X = 0.18
	pointsOfFunctionPlot[118].Y = 0.18

	pointsOfFunctionPlot[119].X = 0.19
	pointsOfFunctionPlot[119].Y = 0.19

	pointsOfFunctionPlot[120].X = 0.20
	pointsOfFunctionPlot[120].Y = 0.201

	pointsOfFunctionPlot[121].X = 0.21
	pointsOfFunctionPlot[121].Y = 0.211

	pointsOfFunctionPlot[122].X = 0.22
	pointsOfFunctionPlot[122].Y = 0.221

	pointsOfFunctionPlot[123].X = 0.23
	pointsOfFunctionPlot[123].Y = 0.232

	pointsOfFunctionPlot[124].X = 0.24
	pointsOfFunctionPlot[124].Y = 0.242

	pointsOfFunctionPlot[125].X = 0.25
	pointsOfFunctionPlot[125].Y = 0.252

	pointsOfFunctionPlot[126].X = 0.26
	pointsOfFunctionPlot[126].Y = 0.263

	pointsOfFunctionPlot[127].X = 0.27
	pointsOfFunctionPlot[127].Y = 0.273

	pointsOfFunctionPlot[128].X = 0.28
	pointsOfFunctionPlot[128].Y = 0.283

	pointsOfFunctionPlot[129].X = 0.29
	pointsOfFunctionPlot[129].Y = 0.294

	pointsOfFunctionPlot[130].X = 0.30
	pointsOfFunctionPlot[130].Y = 0.304

	pointsOfFunctionPlot[131].X = 0.31
	pointsOfFunctionPlot[131].Y = 0.315

	pointsOfFunctionPlot[132].X = 0.32
	pointsOfFunctionPlot[132].Y = 0.325

	pointsOfFunctionPlot[133].X = 0.33
	pointsOfFunctionPlot[133].Y = 0.336

	pointsOfFunctionPlot[134].X = 0.34
	pointsOfFunctionPlot[134].Y = 0.346

	pointsOfFunctionPlot[135].X = 0.35
	pointsOfFunctionPlot[135].Y = 0.357

	pointsOfFunctionPlot[136].X = 0.36
	pointsOfFunctionPlot[136].Y = 0.368

	pointsOfFunctionPlot[137].X = 0.37
	pointsOfFunctionPlot[137].Y = 0.379

	pointsOfFunctionPlot[138].X = 0.38
	pointsOfFunctionPlot[138].Y = 0.389

	pointsOfFunctionPlot[139].X = 0.39
	pointsOfFunctionPlot[139].Y = 0.4

	pointsOfFunctionPlot[140].X = 0.40
	pointsOfFunctionPlot[140].Y = 0.411

	pointsOfFunctionPlot[141].X = 0.41
	pointsOfFunctionPlot[141].Y = 0.422

	pointsOfFunctionPlot[142].X = 0.42
	pointsOfFunctionPlot[142].Y = 0.433

	pointsOfFunctionPlot[143].X = 0.43
	pointsOfFunctionPlot[143].Y = 0.444

	pointsOfFunctionPlot[144].X = 0.44
	pointsOfFunctionPlot[144].Y = 0.455

	pointsOfFunctionPlot[145].X = 0.45
	pointsOfFunctionPlot[145].Y = 0.466

	pointsOfFunctionPlot[146].X = 0.46
	pointsOfFunctionPlot[146].Y = 0.478

	pointsOfFunctionPlot[147].X = 0.47
	pointsOfFunctionPlot[147].Y = 0.489

	pointsOfFunctionPlot[148].X = 0.48
	pointsOfFunctionPlot[148].Y = 0.5

	pointsOfFunctionPlot[149].X = 0.49
	pointsOfFunctionPlot[149].Y = 0.512

	pointsOfFunctionPlot[150].X = 0.50
	pointsOfFunctionPlot[150].Y = 0.523

	pointsOfFunctionPlot[151].X = 0.51
	pointsOfFunctionPlot[151].Y = 0.535

	pointsOfFunctionPlot[152].X = 0.52
	pointsOfFunctionPlot[152].Y = 0.546

	pointsOfFunctionPlot[153].X = 0.53
	pointsOfFunctionPlot[153].Y = 0.558

	pointsOfFunctionPlot[154].X = 0.54
	pointsOfFunctionPlot[154].Y = 0.57

	pointsOfFunctionPlot[155].X = 0.55
	pointsOfFunctionPlot[155].Y = 0.582

	pointsOfFunctionPlot[156].X = 0.56
	pointsOfFunctionPlot[156].Y = 0.594

	pointsOfFunctionPlot[157].X = 0.57
	pointsOfFunctionPlot[157].Y = 0.606

	pointsOfFunctionPlot[158].X = 0.58
	pointsOfFunctionPlot[158].Y = 0.618

	pointsOfFunctionPlot[159].X = 0.59
	pointsOfFunctionPlot[159].Y = 0.631

	pointsOfFunctionPlot[160].X = 0.60
	pointsOfFunctionPlot[160].Y = 0.643

	pointsOfFunctionPlot[161].X = 0.61
	pointsOfFunctionPlot[161].Y = 0.656

	pointsOfFunctionPlot[162].X = 0.62
	pointsOfFunctionPlot[162].Y = 0.668

	pointsOfFunctionPlot[163].X = 0.63
	pointsOfFunctionPlot[163].Y = 0.681

	pointsOfFunctionPlot[164].X = 0.64
	pointsOfFunctionPlot[164].Y = 0.694

	pointsOfFunctionPlot[165].X = 0.65
	pointsOfFunctionPlot[165].Y = 0.707

	pointsOfFunctionPlot[166].X = 0.66
	pointsOfFunctionPlot[166].Y = 0.72

	pointsOfFunctionPlot[167].X = 0.67
	pointsOfFunctionPlot[167].Y = 0.734

	pointsOfFunctionPlot[168].X = 0.68
	pointsOfFunctionPlot[168].Y = 0.747

	pointsOfFunctionPlot[169].X = 0.69
	pointsOfFunctionPlot[169].Y = 0.761

	pointsOfFunctionPlot[170].X = 0.70
	pointsOfFunctionPlot[170].Y = 0.775

	pointsOfFunctionPlot[171].X = 0.71
	pointsOfFunctionPlot[171].Y = 0.789

	pointsOfFunctionPlot[172].X = 0.72
	pointsOfFunctionPlot[172].Y = 0.803

	pointsOfFunctionPlot[173].X = 0.73
	pointsOfFunctionPlot[173].Y = 0.818

	pointsOfFunctionPlot[174].X = 0.74
	pointsOfFunctionPlot[174].Y = 0.833

	pointsOfFunctionPlot[175].X = 0.75
	pointsOfFunctionPlot[175].Y = 0.848

	pointsOfFunctionPlot[176].X = 0.76
	pointsOfFunctionPlot[176].Y = 0.863

	pointsOfFunctionPlot[177].X = 0.77
	pointsOfFunctionPlot[177].Y = 0.878

	pointsOfFunctionPlot[178].X = 0.78
	pointsOfFunctionPlot[178].Y = 0.894

	pointsOfFunctionPlot[179].X = 0.79
	pointsOfFunctionPlot[179].Y = 0.91

	pointsOfFunctionPlot[180].X = 0.80
	pointsOfFunctionPlot[180].Y = 0.927

	pointsOfFunctionPlot[181].X = 0.81
	pointsOfFunctionPlot[181].Y = 0.944

	pointsOfFunctionPlot[182].X = 0.82
	pointsOfFunctionPlot[182].Y = 0.961

	pointsOfFunctionPlot[183].X = 0.83
	pointsOfFunctionPlot[183].Y = 0.979

	pointsOfFunctionPlot[184].X = 0.84
	pointsOfFunctionPlot[184].Y = 0.997

	pointsOfFunctionPlot[185].X = 0.85
	pointsOfFunctionPlot[185].Y = 1.016

	pointsOfFunctionPlot[186].X = 0.86
	pointsOfFunctionPlot[186].Y = 1.035

	pointsOfFunctionPlot[187].X = 0.87
	pointsOfFunctionPlot[187].Y = 1.055

	pointsOfFunctionPlot[188].X = 0.88
	pointsOfFunctionPlot[188].Y = 1.075

	pointsOfFunctionPlot[189].X = 0.89
	pointsOfFunctionPlot[189].Y = 1.097

	pointsOfFunctionPlot[190].X = 0.90
	pointsOfFunctionPlot[190].Y = 1.119

	pointsOfFunctionPlot[191].X = 0.91
	pointsOfFunctionPlot[191].Y = 1.143

	pointsOfFunctionPlot[192].X = 0.92
	pointsOfFunctionPlot[192].Y = 1.168

	pointsOfFunctionPlot[193].X = 0.93
	pointsOfFunctionPlot[193].Y = 1.194

	pointsOfFunctionPlot[194].X = 0.94
	pointsOfFunctionPlot[194].Y = 1.222

	pointsOfFunctionPlot[195].X = 0.95
	pointsOfFunctionPlot[195].Y = 1.253

	pointsOfFunctionPlot[196].X = 0.96
	pointsOfFunctionPlot[196].Y = 1.287

	pointsOfFunctionPlot[197].X = 0.97
	pointsOfFunctionPlot[197].Y = 1.325

	pointsOfFunctionPlot[198].X = 0.98
	pointsOfFunctionPlot[198].Y = 1.37

	pointsOfFunctionPlot[199].X = 0.99
	pointsOfFunctionPlot[199].Y = 1.429

	pointsOfFunctionPlot[200].X = 1.0
	pointsOfFunctionPlot[200].Y = 1.57










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function arcsin(x)"

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
		"arcsin-function-plot-01.png"); err != nil {

		panic(err)
	}
}
