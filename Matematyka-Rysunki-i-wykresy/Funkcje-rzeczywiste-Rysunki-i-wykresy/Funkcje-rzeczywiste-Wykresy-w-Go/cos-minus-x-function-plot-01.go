package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function f(x) = cos(x) - x.

	pointsOfFunctionPlot := make(plotter.XYs, 158)

	pointsOfFunctionPlot[0].X = 0.0
	pointsOfFunctionPlot[0].Y = 1.0

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = 0.99

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = 0.978

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

	pointsOfFunctionPlot[10].X = 0.10
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

	pointsOfFunctionPlot[20].X = 0.20
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

	pointsOfFunctionPlot[30].X = 0.30
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

	pointsOfFunctionPlot[40].X = 0.40
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

	pointsOfFunctionPlot[50].X = 0.50
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

	pointsOfFunctionPlot[60].X = 0.60
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

	pointsOfFunctionPlot[70].X = 0.70
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

	pointsOfFunctionPlot[80].X = 0.80
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

	pointsOfFunctionPlot[90].X = 0.90
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

	pointsOfFunctionPlot[110].X = 1.10
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

	pointsOfFunctionPlot[120].X = 1.20
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

	pointsOfFunctionPlot[130].X = 1.30
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

	pointsOfFunctionPlot[140].X = 1.40
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

	pointsOfFunctionPlot[150].X = 1.50
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
		"cos-minus-x-function-plot-01.png"); err != nil {

		panic(err)
	}
}
