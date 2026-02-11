package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of Gamma function.

	pointsOfFunctionPlot1 := make(plotter.XYs, 101)

	pointsOfFunctionPlot1[0].X = -1.999
	pointsOfFunctionPlot1[0].Y = 500.462

	pointsOfFunctionPlot1[1].X = -1.99
	pointsOfFunctionPlot1[1].Y = 50.47

	pointsOfFunctionPlot1[2].X = -1.98
	pointsOfFunctionPlot1[2].Y = 25.48

	pointsOfFunctionPlot1[3].X = -1.97
	pointsOfFunctionPlot1[3].Y = 17.156

	pointsOfFunctionPlot1[4].X = -1.96
	pointsOfFunctionPlot1[4].Y = 13.0

	pointsOfFunctionPlot1[5].X = -1.95
	pointsOfFunctionPlot1[5].Y = 10.51

	pointsOfFunctionPlot1[6].X = -1.94
	pointsOfFunctionPlot1[6].Y = 8.853

	pointsOfFunctionPlot1[7].X = -1.93
	pointsOfFunctionPlot1[7].Y = 7.673

	pointsOfFunctionPlot1[8].X = -1.92
	pointsOfFunctionPlot1[8].Y = 6.791

	pointsOfFunctionPlot1[9].X = -1.91
	pointsOfFunctionPlot1[9].Y = 6.107

	pointsOfFunctionPlot1[10].X = -1.90
	pointsOfFunctionPlot1[10].Y = 5.563

	pointsOfFunctionPlot1[11].X = -1.89
	pointsOfFunctionPlot1[11].Y = 5.12

	pointsOfFunctionPlot1[12].X = -1.88
	pointsOfFunctionPlot1[12].Y = 4.752

	pointsOfFunctionPlot1[13].X = -1.87
	pointsOfFunctionPlot1[13].Y = 4.444

	pointsOfFunctionPlot1[14].X = -1.86
	pointsOfFunctionPlot1[14].Y = 4.181

	pointsOfFunctionPlot1[15].X = -1.85
	pointsOfFunctionPlot1[15].Y = 3.955

	pointsOfFunctionPlot1[16].X = -1.84
	pointsOfFunctionPlot1[16].Y = 3.759

	pointsOfFunctionPlot1[17].X = -1.83
	pointsOfFunctionPlot1[17].Y = 3.588

	pointsOfFunctionPlot1[18].X = -1.82
	pointsOfFunctionPlot1[18].Y = 3.438

	pointsOfFunctionPlot1[19].X = -1.81
	pointsOfFunctionPlot1[19].Y = 3.305

	pointsOfFunctionPlot1[20].X = -1.80
	pointsOfFunctionPlot1[20].Y = 3.188

	pointsOfFunctionPlot1[21].X = -1.79
	pointsOfFunctionPlot1[21].Y = 3.083

	pointsOfFunctionPlot1[22].X = -1.78
	pointsOfFunctionPlot1[22].Y = 2.989

	pointsOfFunctionPlot1[23].X = -1.77
	pointsOfFunctionPlot1[23].Y = 2.905

	pointsOfFunctionPlot1[24].X = -1.76
	pointsOfFunctionPlot1[24].Y = 2.83

	pointsOfFunctionPlot1[25].X = -1.75
	pointsOfFunctionPlot1[25].Y = 2.762

	pointsOfFunctionPlot1[26].X = -1.74
	pointsOfFunctionPlot1[26].Y = 2.701

	pointsOfFunctionPlot1[27].X = -1.73
	pointsOfFunctionPlot1[27].Y = 2.646

	pointsOfFunctionPlot1[28].X = -1.72
	pointsOfFunctionPlot1[28].Y = 2.597

	pointsOfFunctionPlot1[29].X = -1.71
	pointsOfFunctionPlot1[29].Y = 2.553

	pointsOfFunctionPlot1[30].X = -1.70
	pointsOfFunctionPlot1[30].Y = 2.513

	pointsOfFunctionPlot1[31].X = -1.69
	pointsOfFunctionPlot1[31].Y = 2.478

	pointsOfFunctionPlot1[32].X = -1.68
	pointsOfFunctionPlot1[32].Y = 2.447

	pointsOfFunctionPlot1[33].X = -1.67
	pointsOfFunctionPlot1[33].Y = 2.419

	pointsOfFunctionPlot1[34].X = -1.66
	pointsOfFunctionPlot1[34].Y = 2.395

	pointsOfFunctionPlot1[35].X = -1.65
	pointsOfFunctionPlot1[35].Y = 2.374

	pointsOfFunctionPlot1[36].X = -1.64
	pointsOfFunctionPlot1[36].Y = 2.355

	pointsOfFunctionPlot1[37].X = -1.63
	pointsOfFunctionPlot1[37].Y = 2.34

	pointsOfFunctionPlot1[38].X = -1.62
	pointsOfFunctionPlot1[38].Y = 2.328

	pointsOfFunctionPlot1[39].X = -1.61
	pointsOfFunctionPlot1[39].Y = 2.318

	pointsOfFunctionPlot1[40].X = -1.60
	pointsOfFunctionPlot1[40].Y = 2.31

	pointsOfFunctionPlot1[41].X = -1.59
	pointsOfFunctionPlot1[41].Y = 2.305

	pointsOfFunctionPlot1[42].X = -1.58
	pointsOfFunctionPlot1[42].Y = 2.302

	pointsOfFunctionPlot1[43].X = -1.57
	pointsOfFunctionPlot1[43].Y = 2.302

	pointsOfFunctionPlot1[44].X = -1.56
	pointsOfFunctionPlot1[44].Y = 2.304

	pointsOfFunctionPlot1[45].X = -1.55
	pointsOfFunctionPlot1[45].Y = 2.308

	pointsOfFunctionPlot1[46].X = -1.54
	pointsOfFunctionPlot1[46].Y = 2.315

	pointsOfFunctionPlot1[47].X = -1.53
	pointsOfFunctionPlot1[47].Y = 2.323

	pointsOfFunctionPlot1[48].X = -1.52
	pointsOfFunctionPlot1[48].Y = 2.334

	pointsOfFunctionPlot1[49].X = -1.51
	pointsOfFunctionPlot1[49].Y = 2.347

	pointsOfFunctionPlot1[50].X = -1.50
	pointsOfFunctionPlot1[50].Y = 2.363

	pointsOfFunctionPlot1[51].X = -1.49
	pointsOfFunctionPlot1[51].Y = 2.381

	pointsOfFunctionPlot1[52].X = -1.48
	pointsOfFunctionPlot1[52].Y = 2.401

	pointsOfFunctionPlot1[53].X = -1.47
	pointsOfFunctionPlot1[53].Y = 2.423

	pointsOfFunctionPlot1[54].X = -1.46
	pointsOfFunctionPlot1[54].Y = 2.449

	pointsOfFunctionPlot1[55].X = -1.45
	pointsOfFunctionPlot1[55].Y = 2.476

	pointsOfFunctionPlot1[56].X = -1.44
	pointsOfFunctionPlot1[56].Y = 2.507

	pointsOfFunctionPlot1[57].X = -1.43
	pointsOfFunctionPlot1[57].Y = 2.54

	pointsOfFunctionPlot1[58].X = -1.42
	pointsOfFunctionPlot1[58].Y = 2.577

	pointsOfFunctionPlot1[59].X = -1.41
	pointsOfFunctionPlot1[59].Y = 2.616

	pointsOfFunctionPlot1[60].X = -1.40
	pointsOfFunctionPlot1[60].Y = 2.659

	pointsOfFunctionPlot1[61].X = -1.39
	pointsOfFunctionPlot1[61].Y = 2.705

	pointsOfFunctionPlot1[62].X = -1.38
	pointsOfFunctionPlot1[62].Y = 2.755

	pointsOfFunctionPlot1[63].X = -1.37
	pointsOfFunctionPlot1[63].Y = 2.809

	pointsOfFunctionPlot1[64].X = -1.36
	pointsOfFunctionPlot1[64].Y = 2.867

	pointsOfFunctionPlot1[65].X = -1.35
	pointsOfFunctionPlot1[65].Y = 2.93

	pointsOfFunctionPlot1[66].X = -1.34
	pointsOfFunctionPlot1[66].Y = 2.998

	pointsOfFunctionPlot1[67].X = -1.33
	pointsOfFunctionPlot1[67].Y = 3.071

	pointsOfFunctionPlot1[68].X = -1.32
	pointsOfFunctionPlot1[68].Y = 3.15

	pointsOfFunctionPlot1[69].X = -1.31
	pointsOfFunctionPlot1[69].Y = 3.236

	pointsOfFunctionPlot1[70].X = -1.30
	pointsOfFunctionPlot1[70].Y = 3.328

	pointsOfFunctionPlot1[71].X = -1.29
	pointsOfFunctionPlot1[71].Y = 3.428

	pointsOfFunctionPlot1[72].X = -1.28
	pointsOfFunctionPlot1[72].Y = 3.536

	pointsOfFunctionPlot1[73].X = -1.27
	pointsOfFunctionPlot1[73].Y = 3.654

	pointsOfFunctionPlot1[74].X = -1.26
	pointsOfFunctionPlot1[74].Y = 3.781

	pointsOfFunctionPlot1[75].X = -1.25
	pointsOfFunctionPlot1[75].Y = 3.921

	pointsOfFunctionPlot1[76].X = -1.24
	pointsOfFunctionPlot1[76].Y = 4.073

	pointsOfFunctionPlot1[77].X = -1.23
	pointsOfFunctionPlot1[77].Y = 4.24

	pointsOfFunctionPlot1[78].X = -1.22
	pointsOfFunctionPlot1[78].Y = 4.424

	pointsOfFunctionPlot1[79].X = -1.21
	pointsOfFunctionPlot1[79].Y = 4.626

	pointsOfFunctionPlot1[80].X = -1.20
	pointsOfFunctionPlot1[80].Y = 4.85

	pointsOfFunctionPlot1[81].X = -1.19
	pointsOfFunctionPlot1[81].Y = 5.1

	pointsOfFunctionPlot1[82].X = -1.18
	pointsOfFunctionPlot1[82].Y = 5.378

	pointsOfFunctionPlot1[83].X = -1.17
	pointsOfFunctionPlot1[83].Y = 5.692

	pointsOfFunctionPlot1[84].X = -1.16
	pointsOfFunctionPlot1[84].Y = 6.046

	pointsOfFunctionPlot1[85].X = -1.15
	pointsOfFunctionPlot1[85].Y = 6.449

	pointsOfFunctionPlot1[86].X = -1.14
	pointsOfFunctionPlot1[86].Y = 6.911

	pointsOfFunctionPlot1[87].X = -1.13
	pointsOfFunctionPlot1[87].Y = 7.447

	pointsOfFunctionPlot1[88].X = -1.12
	pointsOfFunctionPlot1[88].Y = 8.075

	pointsOfFunctionPlot1[89].X = -1.11
	pointsOfFunctionPlot1[89].Y = 8.819

	pointsOfFunctionPlot1[90].X = -1.10
	pointsOfFunctionPlot1[90].Y = 9.714

	pointsOfFunctionPlot1[91].X = -1.09
	pointsOfFunctionPlot1[91].Y = 10.812

	pointsOfFunctionPlot1[92].X = -1.08
	pointsOfFunctionPlot1[92].Y = 12.187

	pointsOfFunctionPlot1[93].X = -1.07
	pointsOfFunctionPlot1[93].Y = 13.959

	pointsOfFunctionPlot1[94].X = -1.06
	pointsOfFunctionPlot1[94].Y = 16.327

	pointsOfFunctionPlot1[95].X = -1.05
	pointsOfFunctionPlot1[95].Y = 19.646

	pointsOfFunctionPlot1[96].X = -1.04
	pointsOfFunctionPlot1[96].Y = 24.632

	pointsOfFunctionPlot1[97].X = -1.03
	pointsOfFunctionPlot1[97].Y = 32.952

	pointsOfFunctionPlot1[98].X = -1.02
	pointsOfFunctionPlot1[98].Y = 49.605

	pointsOfFunctionPlot1[99].X = -1.01
	pointsOfFunctionPlot1[99].Y = 99.591

	pointsOfFunctionPlot1[100].X = -1.001
	pointsOfFunctionPlot1[100].Y = 999.578





	pointsOfFunctionPlot2 := make(plotter.XYs, 101)

	pointsOfFunctionPlot2[0].X = -0.999
	pointsOfFunctionPlot2[0].Y = -1_000.424

	pointsOfFunctionPlot2[1].X = -0.99
	pointsOfFunctionPlot2[1].Y = -100.436

	pointsOfFunctionPlot2[2].X = -0.98
	pointsOfFunctionPlot2[2].Y = -50.451

	pointsOfFunctionPlot2[3].X = -0.97
	pointsOfFunctionPlot2[3].Y = -33.798

	pointsOfFunctionPlot2[4].X = -0.96
	pointsOfFunctionPlot2[4].Y = -25.48

	pointsOfFunctionPlot2[5].X = -0.95
	pointsOfFunctionPlot2[5].Y = -20.494

	pointsOfFunctionPlot2[6].X = -0.94
	pointsOfFunctionPlot2[6].Y = -17.176

	pointsOfFunctionPlot2[7].X = -0.93
	pointsOfFunctionPlot2[7].Y = -14.81

	pointsOfFunctionPlot2[8].X = -0.92
	pointsOfFunctionPlot2[8].Y = -13.039

	pointsOfFunctionPlot2[9].X = -0.91
	pointsOfFunctionPlot2[9].Y = -11.666

	pointsOfFunctionPlot2[10].X = -0.90
	pointsOfFunctionPlot2[10].Y = -10.57

	pointsOfFunctionPlot2[11].X = -0.89
	pointsOfFunctionPlot2[11].Y = -9.677

	pointsOfFunctionPlot2[12].X = -0.88
	pointsOfFunctionPlot2[12].Y = -8.935

	pointsOfFunctionPlot2[13].X = -0.87
	pointsOfFunctionPlot2[13].Y = -8.31

	pointsOfFunctionPlot2[14].X = -0.86
	pointsOfFunctionPlot2[14].Y = -7.777

	pointsOfFunctionPlot2[15].X = -0.85
	pointsOfFunctionPlot2[15].Y = -7.317

	pointsOfFunctionPlot2[16].X = -0.84
	pointsOfFunctionPlot2[16].Y = -6.918

	pointsOfFunctionPlot2[17].X = -0.83
	pointsOfFunctionPlot2[17].Y = -6.567

	pointsOfFunctionPlot2[18].X = -0.82
	pointsOfFunctionPlot2[18].Y = -6.258

	pointsOfFunctionPlot2[19].X = -0.81
	pointsOfFunctionPlot2[19].Y = -5.983

	pointsOfFunctionPlot2[20].X = -0.80
	pointsOfFunctionPlot2[20].Y = -5.738

	pointsOfFunctionPlot2[21].X = -0.79
	pointsOfFunctionPlot2[21].Y = -5.518

	pointsOfFunctionPlot2[22].X = -0.78
	pointsOfFunctionPlot2[22].Y = -5.321

	pointsOfFunctionPlot2[23].X = -0.77
	pointsOfFunctionPlot2[23].Y = -5.142

	pointsOfFunctionPlot2[24].X = -0.76
	pointsOfFunctionPlot2[24].Y = -4.98

	pointsOfFunctionPlot2[25].X = -0.75
	pointsOfFunctionPlot2[25].Y = -4.834

	pointsOfFunctionPlot2[26].X = -0.74
	pointsOfFunctionPlot2[26].Y = -4.7

	pointsOfFunctionPlot2[27].X = -0.73
	pointsOfFunctionPlot2[27].Y = -4.578

	pointsOfFunctionPlot2[28].X = -0.72
	pointsOfFunctionPlot2[28].Y = -4.467

	pointsOfFunctionPlot2[29].X = -0.71
	pointsOfFunctionPlot2[29].Y = -4.366

	pointsOfFunctionPlot2[30].X = -0.70
	pointsOfFunctionPlot2[30].Y = -4.273

	pointsOfFunctionPlot2[31].X = -0.69
	pointsOfFunctionPlot2[31].Y = -4.188

	pointsOfFunctionPlot2[32].X = -0.68
	pointsOfFunctionPlot2[32].Y = -4.111

	pointsOfFunctionPlot2[33].X = -0.67
	pointsOfFunctionPlot2[33].Y = -4.04

	pointsOfFunctionPlot2[34].X = -0.66
	pointsOfFunctionPlot2[34].Y = -3.976

	pointsOfFunctionPlot2[35].X = -0.65
	pointsOfFunctionPlot2[35].Y = -3.917

	pointsOfFunctionPlot2[36].X = -0.64
	pointsOfFunctionPlot2[36].Y = -3.863

	pointsOfFunctionPlot2[37].X = -0.63
	pointsOfFunctionPlot2[37].Y = -3.815

	pointsOfFunctionPlot2[38].X = -0.62
	pointsOfFunctionPlot2[38].Y = -3.771

	pointsOfFunctionPlot2[39].X = -0.61
	pointsOfFunctionPlot2[39].Y = -3.732

	pointsOfFunctionPlot2[40].X = -0.60
	pointsOfFunctionPlot2[40].Y = -3.696

	pointsOfFunctionPlot2[41].X = -0.59
	pointsOfFunctionPlot2[41].Y = -3.665

	pointsOfFunctionPlot2[42].X = -0.58
	pointsOfFunctionPlot2[42].Y = -3.638

	pointsOfFunctionPlot2[43].X = -0.57
	pointsOfFunctionPlot2[43].Y = -3.614

	pointsOfFunctionPlot2[44].X = -0.56
	pointsOfFunctionPlot2[44].Y = -3.594

	pointsOfFunctionPlot2[45].X = -0.55
	pointsOfFunctionPlot2[45].Y = -3.578

	pointsOfFunctionPlot2[46].X = -0.54
	pointsOfFunctionPlot2[46].Y = -3.565

	pointsOfFunctionPlot2[47].X = -0.53
	pointsOfFunctionPlot2[47].Y = -3.555

	pointsOfFunctionPlot2[48].X = -0.52
	pointsOfFunctionPlot2[48].Y = -3.548

	pointsOfFunctionPlot2[49].X = -0.51
	pointsOfFunctionPlot2[49].Y = -3.545

	pointsOfFunctionPlot2[50].X = -0.50
	pointsOfFunctionPlot2[50].Y = -3.544

	pointsOfFunctionPlot2[51].X = -0.49
	pointsOfFunctionPlot2[51].Y = -3.547

	pointsOfFunctionPlot2[52].X = -0.48
	pointsOfFunctionPlot2[52].Y = -3.553

	pointsOfFunctionPlot2[53].X = -0.47
	pointsOfFunctionPlot2[53].Y = -3.563

	pointsOfFunctionPlot2[54].X = -0.46
	pointsOfFunctionPlot2[54].Y = -3.575

	pointsOfFunctionPlot2[55].X = -0.45
	pointsOfFunctionPlot2[55].Y = -3.591

	pointsOfFunctionPlot2[56].X = -0.44
	pointsOfFunctionPlot2[56].Y = -3.61

	pointsOfFunctionPlot2[57].X = -0.43
	pointsOfFunctionPlot2[57].Y = -3.633

	pointsOfFunctionPlot2[58].X = -0.42
	pointsOfFunctionPlot2[58].Y = -3.659

	pointsOfFunctionPlot2[59].X = -0.41
	pointsOfFunctionPlot2[59].Y = -3.689

	pointsOfFunctionPlot2[60].X = -0.40
	pointsOfFunctionPlot2[60].Y = -3.722

	pointsOfFunctionPlot2[61].X = -0.39
	pointsOfFunctionPlot2[61].Y = -3.76

	pointsOfFunctionPlot2[62].X = -0.38
	pointsOfFunctionPlot2[62].Y = -3.802

	pointsOfFunctionPlot2[63].X = -0.37
	pointsOfFunctionPlot2[63].Y = -3.849

	pointsOfFunctionPlot2[64].X = -0.36
	pointsOfFunctionPlot2[64].Y = -3.9

	pointsOfFunctionPlot2[65].X = -0.35
	pointsOfFunctionPlot2[65].Y = -3.956

	pointsOfFunctionPlot2[66].X = -0.34
	pointsOfFunctionPlot2[66].Y = -4.018

	pointsOfFunctionPlot2[67].X = -0.33
	pointsOfFunctionPlot2[67].Y = -4.085

	pointsOfFunctionPlot2[68].X = -0.32
	pointsOfFunctionPlot2[68].Y = -4.159

	pointsOfFunctionPlot2[69].X = -0.31
	pointsOfFunctionPlot2[69].Y = -4.239

	pointsOfFunctionPlot2[70].X = -0.30
	pointsOfFunctionPlot2[70].Y = -4.326

	pointsOfFunctionPlot2[71].X = -0.29
	pointsOfFunctionPlot2[71].Y = -4.422

	pointsOfFunctionPlot2[72].X = -0.28
	pointsOfFunctionPlot2[72].Y = -4.526

	pointsOfFunctionPlot2[73].X = -0.27
	pointsOfFunctionPlot2[73].Y = -4.64

	pointsOfFunctionPlot2[74].X = -0.26
	pointsOfFunctionPlot2[74].Y = -4.765

	pointsOfFunctionPlot2[75].X = -0.25
	pointsOfFunctionPlot2[75].Y = -4.901

	pointsOfFunctionPlot2[76].X = -0.24
	pointsOfFunctionPlot2[76].Y = -5.051

	pointsOfFunctionPlot2[77].X = -0.23
	pointsOfFunctionPlot2[77].Y = -5.216

	pointsOfFunctionPlot2[78].X = -0.22
	pointsOfFunctionPlot2[78].Y = -5.397

	pointsOfFunctionPlot2[79].X = -0.21
	pointsOfFunctionPlot2[79].Y = -5.598

	pointsOfFunctionPlot2[80].X = -0.20
	pointsOfFunctionPlot2[80].Y = -5.821

	pointsOfFunctionPlot2[81].X = -0.19
	pointsOfFunctionPlot2[81].Y = -6.069

	pointsOfFunctionPlot2[82].X = -0.18
	pointsOfFunctionPlot2[82].Y = -6.347

	pointsOfFunctionPlot2[83].X = -0.17
	pointsOfFunctionPlot2[83].Y = -6.659

	pointsOfFunctionPlot2[84].X = -0.16
	pointsOfFunctionPlot2[84].Y = -7.013

	pointsOfFunctionPlot2[85].X = -0.15
	pointsOfFunctionPlot2[85].Y = -7.416

	pointsOfFunctionPlot2[86].X = -0.14
	pointsOfFunctionPlot2[86].Y = -7.879

	pointsOfFunctionPlot2[87].X = -0.13
	pointsOfFunctionPlot2[87].Y = -8.415

	pointsOfFunctionPlot2[88].X = -0.12
	pointsOfFunctionPlot2[88].Y = -9.044

	pointsOfFunctionPlot2[89].X = -0.11
	pointsOfFunctionPlot2[89].Y = -9.789

	pointsOfFunctionPlot2[90].X = -0.10
	pointsOfFunctionPlot2[90].Y = -10.686

	pointsOfFunctionPlot2[91].X = -0.09
	pointsOfFunctionPlot2[91].Y = -11.785

	pointsOfFunctionPlot2[92].X = -0.08
	pointsOfFunctionPlot2[92].Y = -13.162

	pointsOfFunctionPlot2[93].X = -0.07
	pointsOfFunctionPlot2[93].Y = -14.936

	pointsOfFunctionPlot2[94].X = -0.06
	pointsOfFunctionPlot2[94].Y = -17.306

	pointsOfFunctionPlot2[95].X = -0.05
	pointsOfFunctionPlot2[95].Y = -20.629

	pointsOfFunctionPlot2[96].X = -0.04
	pointsOfFunctionPlot2[96].Y = -25.618

	pointsOfFunctionPlot2[97].X = -0.03
	pointsOfFunctionPlot2[97].Y = -33.941

	pointsOfFunctionPlot2[98].X = -0.02
	pointsOfFunctionPlot2[98].Y = -50.597

	pointsOfFunctionPlot2[99].X = -0.01
	pointsOfFunctionPlot2[99].Y = -100.587

	pointsOfFunctionPlot2[100].X = -0.001
	pointsOfFunctionPlot2[100].Y = -1_000.578





	pointsOfFunctionPlot3 := make(plotter.XYs, 101)

	pointsOfFunctionPlot3[0].X = 0.001
	pointsOfFunctionPlot3[0].Y = 999.423

	pointsOfFunctionPlot3[1].X = 0.01
	pointsOfFunctionPlot3[1].Y = 99.432

	pointsOfFunctionPlot3[2].X = 0.02
	pointsOfFunctionPlot3[2].Y = 49.442

	pointsOfFunctionPlot3[3].X = 0.03
	pointsOfFunctionPlot3[3].Y = 32.784

	pointsOfFunctionPlot3[4].X = 0.04
	pointsOfFunctionPlot3[4].Y = 24.46

	pointsOfFunctionPlot3[5].X = 0.05
	pointsOfFunctionPlot3[5].Y = 19.47

	pointsOfFunctionPlot3[6].X = 0.06
	pointsOfFunctionPlot3[6].Y = 16.145

	pointsOfFunctionPlot3[7].X = 0.07
	pointsOfFunctionPlot3[7].Y = 13.773

	pointsOfFunctionPlot3[8].X = 0.08
	pointsOfFunctionPlot3[8].Y = 11.996

	pointsOfFunctionPlot3[9].X = 0.09
	pointsOfFunctionPlot3[9].Y = 10.616

	pointsOfFunctionPlot3[10].X = 0.10
	pointsOfFunctionPlot3[10].Y = 9.513

	pointsOfFunctionPlot3[11].X = 0.11
	pointsOfFunctionPlot3[11].Y = 8.612

	pointsOfFunctionPlot3[12].X = 0.12
	pointsOfFunctionPlot3[12].Y = 7.863

	pointsOfFunctionPlot3[13].X = 0.13
	pointsOfFunctionPlot3[13].Y = 7.23

	pointsOfFunctionPlot3[14].X = 0.14
	pointsOfFunctionPlot3[14].Y = 6.688

	pointsOfFunctionPlot3[15].X = 0.15
	pointsOfFunctionPlot3[15].Y = 6.22

	pointsOfFunctionPlot3[16].X = 0.16
	pointsOfFunctionPlot3[16].Y = 5.811

	pointsOfFunctionPlot3[17].X = 0.17
	pointsOfFunctionPlot3[17].Y = 5.451

	pointsOfFunctionPlot3[18].X = 0.18
	pointsOfFunctionPlot3[18].Y = 5.131

	pointsOfFunctionPlot3[19].X = 0.19
	pointsOfFunctionPlot3[19].Y = 4.846

	pointsOfFunctionPlot3[20].X = 0.20
	pointsOfFunctionPlot3[20].Y = 4.59

	pointsOfFunctionPlot3[21].X = 0.21
	pointsOfFunctionPlot3[21].Y = 4.359

	pointsOfFunctionPlot3[22].X = 0.22
	pointsOfFunctionPlot3[22].Y = 4.15

	pointsOfFunctionPlot3[23].X = 0.23
	pointsOfFunctionPlot3[23].Y = 3.959

	pointsOfFunctionPlot3[24].X = 0.24
	pointsOfFunctionPlot3[24].Y = 3.785

	pointsOfFunctionPlot3[25].X = 0.25
	pointsOfFunctionPlot3[25].Y = 3.625

	pointsOfFunctionPlot3[26].X = 0.26
	pointsOfFunctionPlot3[26].Y = 3.478

	pointsOfFunctionPlot3[27].X = 0.27
	pointsOfFunctionPlot3[27].Y = 3.342

	pointsOfFunctionPlot3[28].X = 0.28
	pointsOfFunctionPlot3[28].Y = 3.216

	pointsOfFunctionPlot3[29].X = 0.29
	pointsOfFunctionPlot3[29].Y = 3.1

	pointsOfFunctionPlot3[30].X = 0.30
	pointsOfFunctionPlot3[30].Y = 2.991

	pointsOfFunctionPlot3[31].X = 0.31
	pointsOfFunctionPlot3[31].Y = 2.89

	pointsOfFunctionPlot3[32].X = 0.32
	pointsOfFunctionPlot3[32].Y = 2.795

	pointsOfFunctionPlot3[33].X = 0.33
	pointsOfFunctionPlot3[33].Y = 2.707

	pointsOfFunctionPlot3[34].X = 0.34
	pointsOfFunctionPlot3[34].Y = 2.624

	pointsOfFunctionPlot3[35].X = 0.35
	pointsOfFunctionPlot3[35].Y = 2.546

	pointsOfFunctionPlot3[36].X = 0.36
	pointsOfFunctionPlot3[36].Y = 2.472

	pointsOfFunctionPlot3[37].X = 0.37
	pointsOfFunctionPlot3[37].Y = 2.403

	pointsOfFunctionPlot3[38].X = 0.38
	pointsOfFunctionPlot3[38].Y = 2.338

	pointsOfFunctionPlot3[39].X = 0.39
	pointsOfFunctionPlot3[39].Y = 2.276

	pointsOfFunctionPlot3[40].X = 0.40
	pointsOfFunctionPlot3[40].Y = 2.218

	pointsOfFunctionPlot3[41].X = 0.41
	pointsOfFunctionPlot3[41].Y = 2.162

	pointsOfFunctionPlot3[42].X = 0.42
	pointsOfFunctionPlot3[42].Y = 2.11

	pointsOfFunctionPlot3[43].X = 0.43
	pointsOfFunctionPlot3[43].Y = 2.06

	pointsOfFunctionPlot3[44].X = 0.44
	pointsOfFunctionPlot3[44].Y = 2.013

	pointsOfFunctionPlot3[45].X = 0.45
	pointsOfFunctionPlot3[45].Y = 1.968

	pointsOfFunctionPlot3[46].X = 0.46
	pointsOfFunctionPlot3[46].Y = 1.925

	pointsOfFunctionPlot3[47].X = 0.47
	pointsOfFunctionPlot3[47].Y = 1.884

	pointsOfFunctionPlot3[48].X = 0.48
	pointsOfFunctionPlot3[48].Y = 1.845

	pointsOfFunctionPlot3[49].X = 0.49
	pointsOfFunctionPlot3[49].Y = 1.808

	pointsOfFunctionPlot3[50].X = 0.50
	pointsOfFunctionPlot3[50].Y = 1.772

	pointsOfFunctionPlot3[51].X = 0.51
	pointsOfFunctionPlot3[51].Y = 1.738

	pointsOfFunctionPlot3[52].X = 0.52
	pointsOfFunctionPlot3[52].Y = 1.705

	pointsOfFunctionPlot3[53].X = 0.53
	pointsOfFunctionPlot3[53].Y = 1.674

	pointsOfFunctionPlot3[54].X = 0.54
	pointsOfFunctionPlot3[54].Y = 1.644

	pointsOfFunctionPlot3[55].X = 0.55
	pointsOfFunctionPlot3[55].Y = 1.616

	pointsOfFunctionPlot3[56].X = 0.56
	pointsOfFunctionPlot3[56].Y = 1.588

	pointsOfFunctionPlot3[57].X = 0.57
	pointsOfFunctionPlot3[57].Y = 1.562

	pointsOfFunctionPlot3[58].X = 0.58
	pointsOfFunctionPlot3[58].Y = 1.536

	pointsOfFunctionPlot3[59].X = 0.59
	pointsOfFunctionPlot3[59].Y = 1.512

	pointsOfFunctionPlot3[60].X = 0.60
	pointsOfFunctionPlot3[60].Y = 1.489

	pointsOfFunctionPlot3[61].X = 0.61
	pointsOfFunctionPlot3[61].Y = 1.466

	pointsOfFunctionPlot3[62].X = 0.62
	pointsOfFunctionPlot3[62].Y = 1.445

	pointsOfFunctionPlot3[63].X = 0.63
	pointsOfFunctionPlot3[63].Y = 1.424

	pointsOfFunctionPlot3[64].X = 0.64
	pointsOfFunctionPlot3[64].Y = 1.404

	pointsOfFunctionPlot3[65].X = 0.65
	pointsOfFunctionPlot3[65].Y = 1.384

	pointsOfFunctionPlot3[66].X = 0.66
	pointsOfFunctionPlot3[66].Y = 1.366

	pointsOfFunctionPlot3[67].X = 0.67
	pointsOfFunctionPlot3[67].Y = 1.348

	pointsOfFunctionPlot3[68].X = 0.68
	pointsOfFunctionPlot3[68].Y = 1.33

	pointsOfFunctionPlot3[69].X = 0.69
	pointsOfFunctionPlot3[69].Y = 1.314

	pointsOfFunctionPlot3[70].X = 0.70
	pointsOfFunctionPlot3[70].Y = 1.298

	pointsOfFunctionPlot3[71].X = 0.71
	pointsOfFunctionPlot3[71].Y = 1.282

	pointsOfFunctionPlot3[72].X = 0.72
	pointsOfFunctionPlot3[72].Y = 1.267

	pointsOfFunctionPlot3[73].X = 0.73
	pointsOfFunctionPlot3[73].Y = 1.252

	pointsOfFunctionPlot3[74].X = 0.74
	pointsOfFunctionPlot3[74].Y = 1.238

	pointsOfFunctionPlot3[75].X = 0.75
	pointsOfFunctionPlot3[75].Y = 1.225

	pointsOfFunctionPlot3[76].X = 0.76
	pointsOfFunctionPlot3[76].Y = 1.212

	pointsOfFunctionPlot3[77].X = 0.77
	pointsOfFunctionPlot3[77].Y = 1.199

	pointsOfFunctionPlot3[78].X = 0.78
	pointsOfFunctionPlot3[78].Y = 1.187

	pointsOfFunctionPlot3[79].X = 0.79
	pointsOfFunctionPlot3[79].Y = 1.175

	pointsOfFunctionPlot3[80].X = 0.80
	pointsOfFunctionPlot3[80].Y = 1.164

	pointsOfFunctionPlot3[81].X = 0.81
	pointsOfFunctionPlot3[81].Y = 1.153

	pointsOfFunctionPlot3[82].X = 0.82
	pointsOfFunctionPlot3[82].Y = 1.142

	pointsOfFunctionPlot3[83].X = 0.83
	pointsOfFunctionPlot3[83].Y = 1.132

	pointsOfFunctionPlot3[84].X = 0.84
	pointsOfFunctionPlot3[84].Y = 1.122

	pointsOfFunctionPlot3[85].X = 0.85
	pointsOfFunctionPlot3[85].Y = 1.112

	pointsOfFunctionPlot3[86].X = 0.86
	pointsOfFunctionPlot3[86].Y = 1.103

	pointsOfFunctionPlot3[87].X = 0.87
	pointsOfFunctionPlot3[87].Y = 1.094

	pointsOfFunctionPlot3[88].X = 0.88
	pointsOfFunctionPlot3[88].Y = 1.085

	pointsOfFunctionPlot3[89].X = 0.89
	pointsOfFunctionPlot3[89].Y = 1.076

	pointsOfFunctionPlot3[90].X = 0.90
	pointsOfFunctionPlot3[90].Y = 1.068

	pointsOfFunctionPlot3[91].X = 0.91
	pointsOfFunctionPlot3[91].Y = 1.06

	pointsOfFunctionPlot3[92].X = 0.92
	pointsOfFunctionPlot3[92].Y = 1.053

	pointsOfFunctionPlot3[93].X = 0.93
	pointsOfFunctionPlot3[93].Y = 1.045

	pointsOfFunctionPlot3[94].X = 0.94
	pointsOfFunctionPlot3[94].Y = 1.038

	pointsOfFunctionPlot3[95].X = 0.95
	pointsOfFunctionPlot3[95].Y = 1.031

	pointsOfFunctionPlot3[96].X = 0.96
	pointsOfFunctionPlot3[96].Y = 1.024

	pointsOfFunctionPlot3[97].X = 0.97
	pointsOfFunctionPlot3[97].Y = 1.018

	pointsOfFunctionPlot3[98].X = 0.98
	pointsOfFunctionPlot3[98].Y = 1.011

	pointsOfFunctionPlot3[99].X = 0.99
	pointsOfFunctionPlot3[99].Y = 1.005

	pointsOfFunctionPlot3[100].X = 1.0
	pointsOfFunctionPlot3[100].Y = 1.0







	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of Gamma function"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot1)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("f(x)", plotLine)



	plotLine, err = plotter.NewLine(pointsOfFunctionPlot2)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)



	plotLine, err = plotter.NewLine(pointsOfFunctionPlot3)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)



	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"Gamma-function-plot-02.png"); err != nil {

		panic(err)
	}
}
