package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of polynomial f(x) = x^5 - 19.222 x^4 + 69.903 x^3 -
	// - 10.34 x^2 - 15.722 x + 9.177.

	pointsOfPolynomialPlotu := make(plotter.XYs, 1_001)

	pointsOfPolynomialPlotu[0].X = 0.0
	pointsOfPolynomialPlotu[0].Y = 9.177

	pointsOfPolynomialPlotu[1].X = 0.01
	pointsOfPolynomialPlotu[1].Y = 9.018

	pointsOfPolynomialPlotu[2].X = 0.02
	pointsOfPolynomialPlotu[2].Y = 8.859

	pointsOfPolynomialPlotu[3].X = 0.03
	pointsOfPolynomialPlotu[3].Y = 8.697

	pointsOfPolynomialPlotu[4].X = 0.04
	pointsOfPolynomialPlotu[4].Y = 8.536

	pointsOfPolynomialPlotu[5].X = 0.05
	pointsOfPolynomialPlotu[5].Y = 8.373

	pointsOfPolynomialPlotu[6].X = 0.06
	pointsOfPolynomialPlotu[6].Y = 8.211

	pointsOfPolynomialPlotu[7].X = 0.07
	pointsOfPolynomialPlotu[7].Y = 8.049

	pointsOfPolynomialPlotu[8].X = 0.08
	pointsOfPolynomialPlotu[8].Y = 7.888

	pointsOfPolynomialPlotu[9].X = 0.09
	pointsOfPolynomialPlotu[9].Y = 7.728

	pointsOfPolynomialPlotu[10].X = 0.10
	pointsOfPolynomialPlotu[10].Y = 7.569

	pointsOfPolynomialPlotu[11].X = 0.11
	pointsOfPolynomialPlotu[11].Y = 7.412

	pointsOfPolynomialPlotu[12].X = 0.12
	pointsOfPolynomialPlotu[12].Y = 7.258

	pointsOfPolynomialPlotu[13].X = 0.13
	pointsOfPolynomialPlotu[13].Y = 7.106

	pointsOfPolynomialPlotu[14].X = 0.14
	pointsOfPolynomialPlotu[14].Y = 6.957

	pointsOfPolynomialPlotu[15].X = 0.15
	pointsOfPolynomialPlotu[15].Y = 6.812

	pointsOfPolynomialPlotu[16].X = 0.16
	pointsOfPolynomialPlotu[16].Y = 6.67

	pointsOfPolynomialPlotu[17].X = 0.17
	pointsOfPolynomialPlotu[17].Y = 6.532

	pointsOfPolynomialPlotu[18].X = 0.18
	pointsOfPolynomialPlotu[18].Y = 6.399

	pointsOfPolynomialPlotu[19].X = 0.19
	pointsOfPolynomialPlotu[19].Y = 6.271

	pointsOfPolynomialPlotu[20].X = 0.20
	pointsOfPolynomialPlotu[20].Y = 6.147

	pointsOfPolynomialPlotu[21].X = 0.21
	pointsOfPolynomialPlotu[21].Y = 6.029

	pointsOfPolynomialPlotu[22].X = 0.22
	pointsOfPolynomialPlotu[22].Y = 5.917

	pointsOfPolynomialPlotu[23].X = 0.23
	pointsOfPolynomialPlotu[23].Y = 5.811

	pointsOfPolynomialPlotu[24].X = 0.24
	pointsOfPolynomialPlotu[24].Y = 5.711

	pointsOfPolynomialPlotu[25].X = 0.25
	pointsOfPolynomialPlotu[25].Y = 5.618

	pointsOfPolynomialPlotu[26].X = 0.26
	pointsOfPolynomialPlotu[26].Y = 5.532

	pointsOfPolynomialPlotu[27].X = 0.27
	pointsOfPolynomialPlotu[27].Y = 5.453

	pointsOfPolynomialPlotu[28].X = 0.28
	pointsOfPolynomialPlotu[28].Y = 5.382

	pointsOfPolynomialPlotu[29].X = 0.29
	pointsOfPolynomialPlotu[29].Y = 5.318

	pointsOfPolynomialPlotu[30].X = 0.30
	pointsOfPolynomialPlotu[30].Y = 5.263

	pointsOfPolynomialPlotu[31].X = 0.31
	pointsOfPolynomialPlotu[31].Y = 5.217

	pointsOfPolynomialPlotu[32].X = 0.32
	pointsOfPolynomialPlotu[32].Y = 5.179

	pointsOfPolynomialPlotu[33].X = 0.33
	pointsOfPolynomialPlotu[33].Y = 5.15

	pointsOfPolynomialPlotu[34].X = 0.34
	pointsOfPolynomialPlotu[34].Y = 5.131

	pointsOfPolynomialPlotu[35].X = 0.35
	pointsOfPolynomialPlotu[35].Y = 5.121

	pointsOfPolynomialPlotu[36].X = 0.36
	pointsOfPolynomialPlotu[36].Y = 5.121

	pointsOfPolynomialPlotu[37].X = 0.37
	pointsOfPolynomialPlotu[37].Y = 5.131

	pointsOfPolynomialPlotu[38].X = 0.38
	pointsOfPolynomialPlotu[38].Y = 5.152

	pointsOfPolynomialPlotu[39].X = 0.39
	pointsOfPolynomialPlotu[39].Y = 5.183

	pointsOfPolynomialPlotu[40].X = 0.40
	pointsOfPolynomialPlotu[40].Y = 5.225

	pointsOfPolynomialPlotu[41].X = 0.41
	pointsOfPolynomialPlotu[41].Y = 5.278

	pointsOfPolynomialPlotu[42].X = 0.42
	pointsOfPolynomialPlotu[42].Y = 5.343

	pointsOfPolynomialPlotu[43].X = 0.43
	pointsOfPolynomialPlotu[43].Y = 5.419

	pointsOfPolynomialPlotu[44].X = 0.44
	pointsOfPolynomialPlotu[44].Y = 5.507

	pointsOfPolynomialPlotu[45].X = 0.45
	pointsOfPolynomialPlotu[45].Y = 5.607

	pointsOfPolynomialPlotu[46].X = 0.46
	pointsOfPolynomialPlotu[46].Y = 5.720

	pointsOfPolynomialPlotu[47].X = 0.47
	pointsOfPolynomialPlotu[47].Y = 5.845

	pointsOfPolynomialPlotu[48].X = 0.48
	pointsOfPolynomialPlotu[48].Y = 5.983

	pointsOfPolynomialPlotu[49].X = 0.49
	pointsOfPolynomialPlotu[49].Y = 6.134

	pointsOfPolynomialPlotu[50].X = 0.50
	pointsOfPolynomialPlotu[50].Y = 6.298

	pointsOfPolynomialPlotu[51].X = 0.51
	pointsOfPolynomialPlotu[51].Y = 6.475

	pointsOfPolynomialPlotu[52].X = 0.52
	pointsOfPolynomialPlotu[52].Y = 6.666

	pointsOfPolynomialPlotu[53].X = 0.53
	pointsOfPolynomialPlotu[53].Y = 6.871

	pointsOfPolynomialPlotu[54].X = 0.54
	pointsOfPolynomialPlotu[54].Y = 7.089

	pointsOfPolynomialPlotu[55].X = 0.55
	pointsOfPolynomialPlotu[55].Y = 7.322

	pointsOfPolynomialPlotu[56].X = 0.56
	pointsOfPolynomialPlotu[56].Y = 7.569

	pointsOfPolynomialPlotu[57].X = 0.57
	pointsOfPolynomialPlotu[57].Y = 7.831

	pointsOfPolynomialPlotu[58].X = 0.58
	pointsOfPolynomialPlotu[58].Y = 8.107

	pointsOfPolynomialPlotu[59].X = 0.59
	pointsOfPolynomialPlotu[59].Y = 8.399

	pointsOfPolynomialPlotu[60].X = 0.60
	pointsOfPolynomialPlotu[60].Y = 8.705

	pointsOfPolynomialPlotu[61].X = 0.61
	pointsOfPolynomialPlotu[61].Y = 9.027

	pointsOfPolynomialPlotu[62].X = 0.62
	pointsOfPolynomialPlotu[62].Y = 9.364

	pointsOfPolynomialPlotu[63].X = 0.63
	pointsOfPolynomialPlotu[63].Y = 9.716

	pointsOfPolynomialPlotu[64].X = 0.64
	pointsOfPolynomialPlotu[64].Y = 10.084

	pointsOfPolynomialPlotu[65].X = 0.65
	pointsOfPolynomialPlotu[65].Y = 10.469

	pointsOfPolynomialPlotu[66].X = 0.66
	pointsOfPolynomialPlotu[66].Y = 10.869

	pointsOfPolynomialPlotu[67].X = 0.67
	pointsOfPolynomialPlotu[67].Y = 11.285

	pointsOfPolynomialPlotu[68].X = 0.68
	pointsOfPolynomialPlotu[68].Y = 11.717

	pointsOfPolynomialPlotu[69].X = 0.69
	pointsOfPolynomialPlotu[69].Y = 12.166

	pointsOfPolynomialPlotu[70].X = 0.70
	pointsOfPolynomialPlotu[70].Y = 12.632

	pointsOfPolynomialPlotu[71].X = 0.71
	pointsOfPolynomialPlotu[71].Y = 13.114

	pointsOfPolynomialPlotu[72].X = 0.72
	pointsOfPolynomialPlotu[72].Y = 13.612

	pointsOfPolynomialPlotu[73].X = 0.73
	pointsOfPolynomialPlotu[73].Y = 14.128

	pointsOfPolynomialPlotu[74].X = 0.74
	pointsOfPolynomialPlotu[74].Y = 14.661

	pointsOfPolynomialPlotu[75].X = 0.75
	pointsOfPolynomialPlotu[75].Y = 15.211

	pointsOfPolynomialPlotu[76].X = 0.76
	pointsOfPolynomialPlotu[76].Y = 15.778

	pointsOfPolynomialPlotu[77].X = 0.77
	pointsOfPolynomialPlotu[77].Y = 16.363

	pointsOfPolynomialPlotu[78].X = 0.78
	pointsOfPolynomialPlotu[78].Y = 16.965

	pointsOfPolynomialPlotu[79].X = 0.79
	pointsOfPolynomialPlotu[79].Y = 17.584

	pointsOfPolynomialPlotu[80].X = 0.80
	pointsOfPolynomialPlotu[80].Y = 18.222

	pointsOfPolynomialPlotu[81].X = 0.81
	pointsOfPolynomialPlotu[81].Y = 18.876

	pointsOfPolynomialPlotu[82].X = 0.82
	pointsOfPolynomialPlotu[82].Y = 19.549

	pointsOfPolynomialPlotu[83].X = 0.83
	pointsOfPolynomialPlotu[83].Y = 20.24

	pointsOfPolynomialPlotu[84].X = 0.84
	pointsOfPolynomialPlotu[84].Y = 20.949

	pointsOfPolynomialPlotu[85].X = 0.85
	pointsOfPolynomialPlotu[85].Y = 21.675

	pointsOfPolynomialPlotu[86].X = 0.86
	pointsOfPolynomialPlotu[86].Y = 22.42

	pointsOfPolynomialPlotu[87].X = 0.87
	pointsOfPolynomialPlotu[87].Y = 23.183

	pointsOfPolynomialPlotu[88].X = 0.88
	pointsOfPolynomialPlotu[88].Y = 23.965

	pointsOfPolynomialPlotu[89].X = 0.89
	pointsOfPolynomialPlotu[89].Y = 24.764

	pointsOfPolynomialPlotu[90].X = 0.90
	pointsOfPolynomialPlotu[90].Y = 25.582

	pointsOfPolynomialPlotu[91].X = 0.91
	pointsOfPolynomialPlotu[91].Y = 26.419

	pointsOfPolynomialPlotu[92].X = 0.92
	pointsOfPolynomialPlotu[92].Y = 27.274

	pointsOfPolynomialPlotu[93].X = 0.93
	pointsOfPolynomialPlotu[93].Y = 28.147

	pointsOfPolynomialPlotu[94].X = 0.94
	pointsOfPolynomialPlotu[94].Y = 29.04

	pointsOfPolynomialPlotu[95].X = 0.95
	pointsOfPolynomialPlotu[95].Y = 29.95

	pointsOfPolynomialPlotu[96].X = 0.96
	pointsOfPolynomialPlotu[96].Y = 30.88

	pointsOfPolynomialPlotu[97].X = 0.97
	pointsOfPolynomialPlotu[97].Y = 31.828

	pointsOfPolynomialPlotu[98].X = 0.98
	pointsOfPolynomialPlotu[98].Y = 32.795

	pointsOfPolynomialPlotu[99].X = 0.99
	pointsOfPolynomialPlotu[99].Y = 33.78

	pointsOfPolynomialPlotu[100].X = 1.0
	pointsOfPolynomialPlotu[100].Y = 34.785

	pointsOfPolynomialPlotu[101].X = 1.01
	pointsOfPolynomialPlotu[101].Y = 35.808

	pointsOfPolynomialPlotu[102].X = 1.02
	pointsOfPolynomialPlotu[102].Y = 36.85

	pointsOfPolynomialPlotu[103].X = 1.03
	pointsOfPolynomialPlotu[103].Y = 37.91

	pointsOfPolynomialPlotu[104].X = 1.04
	pointsOfPolynomialPlotu[104].Y = 38.99

	pointsOfPolynomialPlotu[105].X = 1.05
	pointsOfPolynomialPlotu[105].Y = 40.089

	pointsOfPolynomialPlotu[106].X = 1.06
	pointsOfPolynomialPlotu[106].Y = 41.206

	pointsOfPolynomialPlotu[107].X = 1.07
	pointsOfPolynomialPlotu[107].Y = 42.342

	pointsOfPolynomialPlotu[108].X = 1.08
	pointsOfPolynomialPlotu[108].Y = 43.497

	pointsOfPolynomialPlotu[109].X = 1.09
	pointsOfPolynomialPlotu[109].Y = 44.671

	pointsOfPolynomialPlotu[110].X = 1.10
	pointsOfPolynomialPlotu[110].Y = 45.863

	pointsOfPolynomialPlotu[111].X = 1.11
	pointsOfPolynomialPlotu[111].Y = 47.075

	pointsOfPolynomialPlotu[112].X = 1.12
	pointsOfPolynomialPlotu[112].Y = 48.305

	pointsOfPolynomialPlotu[113].X = 1.13
	pointsOfPolynomialPlotu[113].Y = 49.554

	pointsOfPolynomialPlotu[114].X = 1.14
	pointsOfPolynomialPlotu[114].Y = 50.822

	pointsOfPolynomialPlotu[115].X = 1.15
	pointsOfPolynomialPlotu[115].Y = 52.108

	pointsOfPolynomialPlotu[116].X = 1.16
	pointsOfPolynomialPlotu[116].Y = 53.413

	pointsOfPolynomialPlotu[117].X = 1.17
	pointsOfPolynomialPlotu[117].Y = 54.737

	pointsOfPolynomialPlotu[118].X = 1.18
	pointsOfPolynomialPlotu[118].Y = 56.079

	pointsOfPolynomialPlotu[119].X = 1.19
	pointsOfPolynomialPlotu[119].Y = 57.44

	pointsOfPolynomialPlotu[120].X = 1.20
	pointsOfPolynomialPlotu[120].Y = 58.843

	pointsOfPolynomialPlotu[121].X = 1.21
	pointsOfPolynomialPlotu[121].Y = 60.241

	pointsOfPolynomialPlotu[122].X = 1.22
	pointsOfPolynomialPlotu[122].Y = 61.658

	pointsOfPolynomialPlotu[123].X = 1.23
	pointsOfPolynomialPlotu[123].Y = 63.094

	pointsOfPolynomialPlotu[124].X = 1.24
	pointsOfPolynomialPlotu[124].Y = 64.548

	pointsOfPolynomialPlotu[125].X = 1.25
	pointsOfPolynomialPlotu[125].Y = 66.02

	pointsOfPolynomialPlotu[126].X = 1.26
	pointsOfPolynomialPlotu[126].Y = 67.511

	pointsOfPolynomialPlotu[127].X = 1.27
	pointsOfPolynomialPlotu[127].Y = 69.019

	pointsOfPolynomialPlotu[128].X = 1.28
	pointsOfPolynomialPlotu[128].Y = 70.546

	pointsOfPolynomialPlotu[129].X = 1.29
	pointsOfPolynomialPlotu[129].Y = 72.091

	pointsOfPolynomialPlotu[130].X = 1.30
	pointsOfPolynomialPlotu[130].Y = 73.653

	pointsOfPolynomialPlotu[131].X = 1.31
	pointsOfPolynomialPlotu[131].Y = 75.234

	pointsOfPolynomialPlotu[132].X = 1.32
	pointsOfPolynomialPlotu[132].Y = 76.832

	pointsOfPolynomialPlotu[133].X = 1.33
	pointsOfPolynomialPlotu[133].Y = 78.448

	pointsOfPolynomialPlotu[134].X = 1.34
	pointsOfPolynomialPlotu[134].Y = 80.082

	pointsOfPolynomialPlotu[135].X = 1.35
	pointsOfPolynomialPlotu[135].Y = 81.733

	pointsOfPolynomialPlotu[136].X = 1.36
	pointsOfPolynomialPlotu[136].Y = 83.401

	pointsOfPolynomialPlotu[137].X = 1.37
	pointsOfPolynomialPlotu[137].Y = 85.087

	pointsOfPolynomialPlotu[138].X = 1.38
	pointsOfPolynomialPlotu[138].Y = 86.791

	pointsOfPolynomialPlotu[139].X = 1.39
	pointsOfPolynomialPlotu[139].Y = 88.511

	pointsOfPolynomialPlotu[140].X = 1.40
	pointsOfPolynomialPlotu[140].Y = 90.248

	pointsOfPolynomialPlotu[141].X = 1.41
	pointsOfPolynomialPlotu[141].Y = 92.002

	pointsOfPolynomialPlotu[142].X = 1.42
	pointsOfPolynomialPlotu[142].Y = 93.774

	pointsOfPolynomialPlotu[143].X = 1.43
	pointsOfPolynomialPlotu[143].Y = 95.561

	pointsOfPolynomialPlotu[144].X = 1.44
	pointsOfPolynomialPlotu[144].Y = 97.366

	pointsOfPolynomialPlotu[145].X = 1.45
	pointsOfPolynomialPlotu[145].Y = 99.187

	pointsOfPolynomialPlotu[146].X = 1.46
	pointsOfPolynomialPlotu[146].Y = 101.024

	pointsOfPolynomialPlotu[147].X = 1.47
	pointsOfPolynomialPlotu[147].Y = 102.877

	pointsOfPolynomialPlotu[148].X = 1.48
	pointsOfPolynomialPlotu[148].Y = 104.747

	pointsOfPolynomialPlotu[149].X = 1.49
	pointsOfPolynomialPlotu[149].Y = 106.632

	pointsOfPolynomialPlotu[150].X = 1.50
	pointsOfPolynomialPlotu[150].Y = 108.534

	pointsOfPolynomialPlotu[151].X = 1.51
	pointsOfPolynomialPlotu[151].Y = 110.451

	pointsOfPolynomialPlotu[152].X = 1.52
	pointsOfPolynomialPlotu[152].Y = 112.383

	pointsOfPolynomialPlotu[153].X = 1.53
	pointsOfPolynomialPlotu[153].Y = 114.331

	pointsOfPolynomialPlotu[154].X = 1.54
	pointsOfPolynomialPlotu[154].Y = 116.294

	pointsOfPolynomialPlotu[155].X = 1.55
	pointsOfPolynomialPlotu[155].Y = 118.273

	pointsOfPolynomialPlotu[156].X = 1.56
	pointsOfPolynomialPlotu[156].Y = 120.266

	pointsOfPolynomialPlotu[157].X = 1.57
	pointsOfPolynomialPlotu[157].Y = 122.274

	pointsOfPolynomialPlotu[158].X = 1.58
	pointsOfPolynomialPlotu[158].Y = 124.297

	pointsOfPolynomialPlotu[159].X = 1.59
	pointsOfPolynomialPlotu[159].Y = 126.334

	pointsOfPolynomialPlotu[160].X = 1.60
	pointsOfPolynomialPlotu[160].Y = 128.386

	pointsOfPolynomialPlotu[161].X = 1.61
	pointsOfPolynomialPlotu[161].Y = 130.452

	pointsOfPolynomialPlotu[162].X = 1.62
	pointsOfPolynomialPlotu[162].Y = 132.532

	pointsOfPolynomialPlotu[163].X = 1.63
	pointsOfPolynomialPlotu[163].Y = 134.626

	pointsOfPolynomialPlotu[164].X = 1.64
	pointsOfPolynomialPlotu[164].Y = 136.733

	pointsOfPolynomialPlotu[165].X = 1.65
	pointsOfPolynomialPlotu[165].Y = 138.854

	pointsOfPolynomialPlotu[166].X = 1.66
	pointsOfPolynomialPlotu[166].Y = 140.988

	pointsOfPolynomialPlotu[167].X = 1.67
	pointsOfPolynomialPlotu[167].Y = 143.135

	pointsOfPolynomialPlotu[168].X = 1.68
	pointsOfPolynomialPlotu[168].Y = 145.296

	pointsOfPolynomialPlotu[169].X = 1.69
	pointsOfPolynomialPlotu[169].Y = 147.469

	pointsOfPolynomialPlotu[170].X = 1.70
	pointsOfPolynomialPlotu[170].Y = 149.654

	pointsOfPolynomialPlotu[171].X = 1.71
	pointsOfPolynomialPlotu[171].Y = 151.853

	pointsOfPolynomialPlotu[172].X = 1.72
	pointsOfPolynomialPlotu[172].Y = 154.063

	pointsOfPolynomialPlotu[173].X = 1.73
	pointsOfPolynomialPlotu[173].Y = 156.285

	pointsOfPolynomialPlotu[174].X = 1.74
	pointsOfPolynomialPlotu[174].Y = 158.519

	pointsOfPolynomialPlotu[175].X = 1.75
	pointsOfPolynomialPlotu[175].Y = 160.765

	pointsOfPolynomialPlotu[176].X = 1.76
	pointsOfPolynomialPlotu[176].Y = 163.022

	pointsOfPolynomialPlotu[177].X = 1.77
	pointsOfPolynomialPlotu[177].Y = 165.29

	pointsOfPolynomialPlotu[178].X = 1.78
	pointsOfPolynomialPlotu[178].Y = 167.57

	pointsOfPolynomialPlotu[179].X = 1.79
	pointsOfPolynomialPlotu[179].Y = 169.86

	pointsOfPolynomialPlotu[180].X = 1.80
	pointsOfPolynomialPlotu[180].Y = 172.16

	pointsOfPolynomialPlotu[181].X = 1.81
	pointsOfPolynomialPlotu[181].Y = 174.471

	pointsOfPolynomialPlotu[182].X = 1.82
	pointsOfPolynomialPlotu[182].Y = 176.793

	pointsOfPolynomialPlotu[183].X = 1.83
	pointsOfPolynomialPlotu[183].Y = 179.124

	pointsOfPolynomialPlotu[184].X = 1.84
	pointsOfPolynomialPlotu[184].Y = 181.465

	pointsOfPolynomialPlotu[185].X = 1.85
	pointsOfPolynomialPlotu[185].Y = 183.815

	pointsOfPolynomialPlotu[186].X = 1.86
	pointsOfPolynomialPlotu[186].Y = 186.174

	pointsOfPolynomialPlotu[187].X = 1.87
	pointsOfPolynomialPlotu[187].Y = 188.543

	pointsOfPolynomialPlotu[188].X = 1.88
	pointsOfPolynomialPlotu[188].Y = 190.92

	pointsOfPolynomialPlotu[189].X = 1.89
	pointsOfPolynomialPlotu[189].Y = 193.306

	pointsOfPolynomialPlotu[190].X = 1.90
	pointsOfPolynomialPlotu[190].Y = 195.7

	pointsOfPolynomialPlotu[191].X = 1.91
	pointsOfPolynomialPlotu[191].Y = 198.102

	pointsOfPolynomialPlotu[192].X = 1.92
	pointsOfPolynomialPlotu[192].Y = 200.512

	pointsOfPolynomialPlotu[193].X = 1.93
	pointsOfPolynomialPlotu[193].Y = 202.93

	pointsOfPolynomialPlotu[194].X = 1.94
	pointsOfPolynomialPlotu[194].Y = 205.355

	pointsOfPolynomialPlotu[195].X = 1.95
	pointsOfPolynomialPlotu[195].Y = 207.787

	pointsOfPolynomialPlotu[196].X = 1.96
	pointsOfPolynomialPlotu[196].Y = 210.226

	pointsOfPolynomialPlotu[197].X = 1.97
	pointsOfPolynomialPlotu[197].Y = 212.671

	pointsOfPolynomialPlotu[198].X = 1.98
	pointsOfPolynomialPlotu[198].Y = 215.123

	pointsOfPolynomialPlotu[199].X = 1.99
	pointsOfPolynomialPlotu[199].Y = 217.581

	pointsOfPolynomialPlotu[200].X = 2.0
	pointsOfPolynomialPlotu[200].Y = 220.045

	pointsOfPolynomialPlotu[201].X = 2.01
	pointsOfPolynomialPlotu[201].Y = 222.514

	pointsOfPolynomialPlotu[202].X = 2.02
	pointsOfPolynomialPlotu[202].Y = 224.988

	pointsOfPolynomialPlotu[203].X = 2.03
	pointsOfPolynomialPlotu[203].Y = 227.468

	pointsOfPolynomialPlotu[204].X = 2.04
	pointsOfPolynomialPlotu[204].Y = 229.952

	pointsOfPolynomialPlotu[205].X = 2.05
	pointsOfPolynomialPlotu[205].Y = 232.441

	pointsOfPolynomialPlotu[206].X = 2.06
	pointsOfPolynomialPlotu[206].Y = 234.934

	pointsOfPolynomialPlotu[207].X = 2.07
	pointsOfPolynomialPlotu[207].Y = 237.431

	pointsOfPolynomialPlotu[208].X = 2.08
	pointsOfPolynomialPlotu[208].Y = 239.931

	pointsOfPolynomialPlotu[209].X = 2.09
	pointsOfPolynomialPlotu[209].Y = 242.435

	pointsOfPolynomialPlotu[210].X = 2.10
	pointsOfPolynomialPlotu[210].Y = 244.942

	pointsOfPolynomialPlotu[211].X = 2.11
	pointsOfPolynomialPlotu[211].Y = 247.452

	pointsOfPolynomialPlotu[212].X = 2.12
	pointsOfPolynomialPlotu[212].Y = 249.964

	pointsOfPolynomialPlotu[213].X = 2.13
	pointsOfPolynomialPlotu[213].Y = 252.479

	pointsOfPolynomialPlotu[214].X = 2.14
	pointsOfPolynomialPlotu[214].Y = 254.996

	pointsOfPolynomialPlotu[215].X = 2.15
	pointsOfPolynomialPlotu[215].Y = 257.514

	pointsOfPolynomialPlotu[216].X = 2.16
	pointsOfPolynomialPlotu[216].Y = 260.033

	pointsOfPolynomialPlotu[217].X = 2.17
	pointsOfPolynomialPlotu[217].Y = 262.554

	pointsOfPolynomialPlotu[218].X = 2.18
	pointsOfPolynomialPlotu[218].Y = 265.075

	pointsOfPolynomialPlotu[219].X = 2.19
	pointsOfPolynomialPlotu[219].Y = 267.597

	pointsOfPolynomialPlotu[220].X = 2.20
	pointsOfPolynomialPlotu[220].Y = 270.119

	pointsOfPolynomialPlotu[221].X = 2.21
	pointsOfPolynomialPlotu[221].Y = 272.641

	pointsOfPolynomialPlotu[222].X = 2.22
	pointsOfPolynomialPlotu[222].Y = 275.162

	pointsOfPolynomialPlotu[223].X = 2.23
	pointsOfPolynomialPlotu[223].Y = 277.683

	pointsOfPolynomialPlotu[224].X = 2.24
	pointsOfPolynomialPlotu[224].Y = 280.203

	pointsOfPolynomialPlotu[225].X = 2.25
	pointsOfPolynomialPlotu[225].Y = 282.721

	pointsOfPolynomialPlotu[226].X = 2.26
	pointsOfPolynomialPlotu[226].Y = 285.237

	pointsOfPolynomialPlotu[227].X = 2.27
	pointsOfPolynomialPlotu[227].Y = 287.752

	pointsOfPolynomialPlotu[228].X = 2.28
	pointsOfPolynomialPlotu[228].Y = 290.264

	pointsOfPolynomialPlotu[229].X = 2.29
	pointsOfPolynomialPlotu[229].Y = 292.774

	pointsOfPolynomialPlotu[230].X = 2.30
	pointsOfPolynomialPlotu[230].Y = 295.28

	pointsOfPolynomialPlotu[231].X = 2.31
	pointsOfPolynomialPlotu[231].Y = 297.784

	pointsOfPolynomialPlotu[232].X = 2.32
	pointsOfPolynomialPlotu[232].Y = 300.283

	pointsOfPolynomialPlotu[233].X = 2.33
	pointsOfPolynomialPlotu[233].Y = 302.779

	pointsOfPolynomialPlotu[234].X = 2.34
	pointsOfPolynomialPlotu[234].Y = 305.27

	pointsOfPolynomialPlotu[235].X = 2.35
	pointsOfPolynomialPlotu[235].Y = 307.757

	pointsOfPolynomialPlotu[236].X = 2.36
	pointsOfPolynomialPlotu[236].Y = 310.239

	pointsOfPolynomialPlotu[237].X = 2.37
	pointsOfPolynomialPlotu[237].Y = 312.716

	pointsOfPolynomialPlotu[238].X = 2.38
	pointsOfPolynomialPlotu[238].Y = 315.187

	pointsOfPolynomialPlotu[239].X = 2.39
	pointsOfPolynomialPlotu[239].Y = 317.652

	pointsOfPolynomialPlotu[240].X = 2.40
	pointsOfPolynomialPlotu[240].Y = 320.111

	pointsOfPolynomialPlotu[241].X = 2.41
	pointsOfPolynomialPlotu[241].Y = 322.563

	pointsOfPolynomialPlotu[242].X = 2.42
	pointsOfPolynomialPlotu[242].Y = 325.008

	pointsOfPolynomialPlotu[243].X = 2.43
	pointsOfPolynomialPlotu[243].Y = 327.446

	pointsOfPolynomialPlotu[244].X = 2.44
	pointsOfPolynomialPlotu[244].Y = 329.877

	pointsOfPolynomialPlotu[245].X = 2.45
	pointsOfPolynomialPlotu[245].Y = 332.299

	pointsOfPolynomialPlotu[246].X = 2.46
	pointsOfPolynomialPlotu[246].Y = 334.713

	pointsOfPolynomialPlotu[247].X = 2.47
	pointsOfPolynomialPlotu[247].Y = 337.118

	pointsOfPolynomialPlotu[248].X = 2.48
	pointsOfPolynomialPlotu[248].Y = 339.514

	pointsOfPolynomialPlotu[249].X = 2.49
	pointsOfPolynomialPlotu[249].Y = 341.901

	pointsOfPolynomialPlotu[250].X = 2.50
	pointsOfPolynomialPlotu[250].Y = 344.278

	pointsOfPolynomialPlotu[251].X = 2.51
	pointsOfPolynomialPlotu[251].Y = 346.645

	pointsOfPolynomialPlotu[252].X = 2.52
	pointsOfPolynomialPlotu[252].Y = 349.001

	pointsOfPolynomialPlotu[253].X = 2.53
	pointsOfPolynomialPlotu[253].Y = 351.347

	pointsOfPolynomialPlotu[254].X = 2.54
	pointsOfPolynomialPlotu[254].Y = 353.681

	pointsOfPolynomialPlotu[255].X = 2.55
	pointsOfPolynomialPlotu[255].Y = 356.004

	pointsOfPolynomialPlotu[256].X = 2.56
	pointsOfPolynomialPlotu[256].Y = 358.314

	pointsOfPolynomialPlotu[257].X = 2.57
	pointsOfPolynomialPlotu[257].Y = 360.613

	pointsOfPolynomialPlotu[258].X = 2.58
	pointsOfPolynomialPlotu[258].Y = 362.899

	pointsOfPolynomialPlotu[259].X = 2.59
	pointsOfPolynomialPlotu[259].Y = 365.171

	pointsOfPolynomialPlotu[260].X = 2.60
	pointsOfPolynomialPlotu[260].Y = 367.431

	pointsOfPolynomialPlotu[261].X = 2.61
	pointsOfPolynomialPlotu[261].Y = 369.676

	pointsOfPolynomialPlotu[262].X = 2.62
	pointsOfPolynomialPlotu[262].Y = 371.907

	pointsOfPolynomialPlotu[263].X = 2.63
	pointsOfPolynomialPlotu[263].Y = 374.124

	pointsOfPolynomialPlotu[264].X = 2.64
	pointsOfPolynomialPlotu[264].Y = 376.326

	pointsOfPolynomialPlotu[265].X = 2.65
	pointsOfPolynomialPlotu[265].Y = 378.513

	pointsOfPolynomialPlotu[266].X = 2.66
	pointsOfPolynomialPlotu[266].Y = 380.684

	pointsOfPolynomialPlotu[267].X = 2.67
	pointsOfPolynomialPlotu[267].Y = 382.838

	pointsOfPolynomialPlotu[268].X = 2.68
	pointsOfPolynomialPlotu[268].Y = 384.977

	pointsOfPolynomialPlotu[269].X = 2.69
	pointsOfPolynomialPlotu[269].Y = 387.098

	pointsOfPolynomialPlotu[270].X = 2.70
	pointsOfPolynomialPlotu[270].Y = 389.202

	pointsOfPolynomialPlotu[271].X = 2.71
	pointsOfPolynomialPlotu[271].Y = 391.289

	pointsOfPolynomialPlotu[272].X = 2.72
	pointsOfPolynomialPlotu[272].Y = 393.358

	pointsOfPolynomialPlotu[273].X = 2.73
	pointsOfPolynomialPlotu[273].Y = 395.408

	pointsOfPolynomialPlotu[274].X = 2.74
	pointsOfPolynomialPlotu[274].Y = 397.44

	pointsOfPolynomialPlotu[275].X = 2.75
	pointsOfPolynomialPlotu[275].Y = 399.452

	pointsOfPolynomialPlotu[276].X = 2.76
	pointsOfPolynomialPlotu[276].Y = 401.445

	pointsOfPolynomialPlotu[277].X = 2.77
	pointsOfPolynomialPlotu[277].Y = 403.417

	pointsOfPolynomialPlotu[278].X = 2.78
	pointsOfPolynomialPlotu[278].Y = 405.37

	pointsOfPolynomialPlotu[279].X = 2.79
	pointsOfPolynomialPlotu[279].Y = 407.301

	pointsOfPolynomialPlotu[280].X = 2.80
	pointsOfPolynomialPlotu[280].Y = 409.212

	pointsOfPolynomialPlotu[281].X = 2.81
	pointsOfPolynomialPlotu[281].Y = 411.101

	pointsOfPolynomialPlotu[282].X = 2.82
	pointsOfPolynomialPlotu[282].Y = 412.968

	pointsOfPolynomialPlotu[283].X = 2.83
	pointsOfPolynomialPlotu[283].Y = 414.812

	pointsOfPolynomialPlotu[284].X = 2.84
	pointsOfPolynomialPlotu[284].Y = 416.634

	pointsOfPolynomialPlotu[285].X = 2.85
	pointsOfPolynomialPlotu[285].Y = 418.433

	pointsOfPolynomialPlotu[286].X = 2.86
	pointsOfPolynomialPlotu[286].Y = 420.208

	pointsOfPolynomialPlotu[287].X = 2.87
	pointsOfPolynomialPlotu[287].Y = 421.959

	pointsOfPolynomialPlotu[288].X = 2.88
	pointsOfPolynomialPlotu[288].Y = 423.685

	pointsOfPolynomialPlotu[289].X = 2.89
	pointsOfPolynomialPlotu[289].Y = 425.387

	pointsOfPolynomialPlotu[290].X = 2.90
	pointsOfPolynomialPlotu[290].Y = 427.064

	pointsOfPolynomialPlotu[291].X = 2.91
	pointsOfPolynomialPlotu[291].Y = 428.714

	pointsOfPolynomialPlotu[292].X = 2.92
	pointsOfPolynomialPlotu[292].Y = 430.339

	pointsOfPolynomialPlotu[293].X = 2.93
	pointsOfPolynomialPlotu[293].Y = 431.938

	pointsOfPolynomialPlotu[294].X = 2.94
	pointsOfPolynomialPlotu[294].Y = 433.509

	pointsOfPolynomialPlotu[295].X = 2.95
	pointsOfPolynomialPlotu[295].Y = 435.053

	pointsOfPolynomialPlotu[296].X = 2.96
	pointsOfPolynomialPlotu[296].Y = 436.57

	pointsOfPolynomialPlotu[297].X = 2.97
	pointsOfPolynomialPlotu[297].Y = 438.058

	pointsOfPolynomialPlotu[298].X = 2.98
	pointsOfPolynomialPlotu[298].Y = 439.518

	pointsOfPolynomialPlotu[299].X = 2.99
	pointsOfPolynomialPlotu[299].Y = 440.948

	pointsOfPolynomialPlotu[300].X = 3.0
	pointsOfPolynomialPlotu[300].Y = 442.35

	pointsOfPolynomialPlotu[301].X = 3.01
	pointsOfPolynomialPlotu[301].Y = 443.721

	pointsOfPolynomialPlotu[302].X = 3.02
	pointsOfPolynomialPlotu[302].Y = 445.062

	pointsOfPolynomialPlotu[303].X = 3.03
	pointsOfPolynomialPlotu[303].Y = 446.373

	pointsOfPolynomialPlotu[304].X = 3.04
	pointsOfPolynomialPlotu[304].Y = 447.652

	pointsOfPolynomialPlotu[305].X = 3.05
	pointsOfPolynomialPlotu[305].Y = 448.9

	pointsOfPolynomialPlotu[306].X = 3.06
	pointsOfPolynomialPlotu[306].Y = 450.116

	pointsOfPolynomialPlotu[307].X = 3.07
	pointsOfPolynomialPlotu[307].Y = 451.299

	pointsOfPolynomialPlotu[308].X = 3.08
	pointsOfPolynomialPlotu[308].Y = 452.45

	pointsOfPolynomialPlotu[309].X = 3.09
	pointsOfPolynomialPlotu[309].Y = 453.567

	pointsOfPolynomialPlotu[310].X = 3.10
	pointsOfPolynomialPlotu[310].Y = 454.651

	pointsOfPolynomialPlotu[311].X = 3.11
	pointsOfPolynomialPlotu[311].Y = 455.7

	pointsOfPolynomialPlotu[312].X = 3.12
	pointsOfPolynomialPlotu[312].Y = 456.715

	pointsOfPolynomialPlotu[313].X = 3.13
	pointsOfPolynomialPlotu[313].Y = 457.695

	pointsOfPolynomialPlotu[314].X = 3.14
	pointsOfPolynomialPlotu[314].Y = 458.639

	pointsOfPolynomialPlotu[315].X = 3.15
	pointsOfPolynomialPlotu[315].Y = 459.548

	pointsOfPolynomialPlotu[316].X = 3.16
	pointsOfPolynomialPlotu[316].Y = 460.42

	pointsOfPolynomialPlotu[317].X = 3.17
	pointsOfPolynomialPlotu[317].Y = 461.256

	pointsOfPolynomialPlotu[318].X = 3.18
	pointsOfPolynomialPlotu[318].Y = 462.054

	pointsOfPolynomialPlotu[319].X = 3.19
	pointsOfPolynomialPlotu[319].Y = 462.815

	pointsOfPolynomialPlotu[320].X = 3.20
	pointsOfPolynomialPlotu[320].Y = 463.538

	pointsOfPolynomialPlotu[321].X = 3.21
	pointsOfPolynomialPlotu[321].Y = 464.222

	pointsOfPolynomialPlotu[322].X = 3.22
	pointsOfPolynomialPlotu[322].Y = 464.867

	pointsOfPolynomialPlotu[323].X = 3.23
	pointsOfPolynomialPlotu[323].Y = 465.473

	pointsOfPolynomialPlotu[324].X = 3.24
	pointsOfPolynomialPlotu[324].Y = 466.038

	pointsOfPolynomialPlotu[325].X = 3.25
	pointsOfPolynomialPlotu[325].Y = 466.564

	pointsOfPolynomialPlotu[326].X = 3.26
	pointsOfPolynomialPlotu[326].Y = 467.049

	pointsOfPolynomialPlotu[327].X = 3.27
	pointsOfPolynomialPlotu[327].Y = 467.493

	pointsOfPolynomialPlotu[328].X = 3.28
	pointsOfPolynomialPlotu[328].Y = 467.895

	pointsOfPolynomialPlotu[329].X = 3.29
	pointsOfPolynomialPlotu[329].Y = 468.255

	pointsOfPolynomialPlotu[330].X = 3.30
	pointsOfPolynomialPlotu[330].Y = 468.572

	pointsOfPolynomialPlotu[331].X = 3.31
	pointsOfPolynomialPlotu[331].Y = 468.846

	pointsOfPolynomialPlotu[332].X = 3.32
	pointsOfPolynomialPlotu[332].Y = 469.078

	pointsOfPolynomialPlotu[333].X = 3.33
	pointsOfPolynomialPlotu[333].Y = 469.265

	pointsOfPolynomialPlotu[334].X = 3.34
	pointsOfPolynomialPlotu[334].Y = 469.407

	pointsOfPolynomialPlotu[335].X = 3.35
	pointsOfPolynomialPlotu[335].Y = 469.505

	pointsOfPolynomialPlotu[336].X = 3.36
	pointsOfPolynomialPlotu[336].Y = 469.558

	pointsOfPolynomialPlotu[337].X = 3.37
	pointsOfPolynomialPlotu[337].Y = 469.565

	pointsOfPolynomialPlotu[338].X = 3.38
	pointsOfPolynomialPlotu[338].Y = 469.526

	pointsOfPolynomialPlotu[339].X = 3.39
	pointsOfPolynomialPlotu[339].Y = 469.441

	pointsOfPolynomialPlotu[340].X = 3.40
	pointsOfPolynomialPlotu[340].Y = 469.308

	pointsOfPolynomialPlotu[341].X = 3.41
	pointsOfPolynomialPlotu[341].Y = 469.128

	pointsOfPolynomialPlotu[342].X = 3.42
	pointsOfPolynomialPlotu[342].Y = 468.9

	pointsOfPolynomialPlotu[343].X = 3.43
	pointsOfPolynomialPlotu[343].Y = 468.623

	pointsOfPolynomialPlotu[344].X = 3.44
	pointsOfPolynomialPlotu[344].Y = 468.298

	pointsOfPolynomialPlotu[345].X = 3.45
	pointsOfPolynomialPlotu[345].Y = 467.923

	pointsOfPolynomialPlotu[346].X = 3.46
	pointsOfPolynomialPlotu[346].Y = 467.498

	pointsOfPolynomialPlotu[347].X = 3.47
	pointsOfPolynomialPlotu[347].Y = 467.024

	pointsOfPolynomialPlotu[348].X = 3.48
	pointsOfPolynomialPlotu[348].Y = 466.498

	pointsOfPolynomialPlotu[349].X = 3.49
	pointsOfPolynomialPlotu[349].Y = 465.921

	pointsOfPolynomialPlotu[350].X = 3.50
	pointsOfPolynomialPlotu[350].Y = 465.293

	pointsOfPolynomialPlotu[351].X = 3.51
	pointsOfPolynomialPlotu[351].Y = 464.613

	pointsOfPolynomialPlotu[352].X = 3.52
	pointsOfPolynomialPlotu[352].Y = 463.88

	pointsOfPolynomialPlotu[353].X = 3.53
	pointsOfPolynomialPlotu[353].Y = 463.094

	pointsOfPolynomialPlotu[354].X = 3.54
	pointsOfPolynomialPlotu[354].Y = 462.254

	pointsOfPolynomialPlotu[355].X = 3.55
	pointsOfPolynomialPlotu[355].Y = 461.361

	pointsOfPolynomialPlotu[356].X = 3.56
	pointsOfPolynomialPlotu[356].Y = 460.413

	pointsOfPolynomialPlotu[357].X = 3.57
	pointsOfPolynomialPlotu[357].Y = 459.411

	pointsOfPolynomialPlotu[358].X = 3.58
	pointsOfPolynomialPlotu[358].Y = 458.353

	pointsOfPolynomialPlotu[359].X = 3.59
	pointsOfPolynomialPlotu[359].Y = 457.239

	pointsOfPolynomialPlotu[360].X = 3.60
	pointsOfPolynomialPlotu[360].Y = 456.069

	pointsOfPolynomialPlotu[361].X = 3.61
	pointsOfPolynomialPlotu[361].Y = 454.843

	pointsOfPolynomialPlotu[362].X = 3.62
	pointsOfPolynomialPlotu[362].Y = 453.559

	pointsOfPolynomialPlotu[363].X = 3.63
	pointsOfPolynomialPlotu[363].Y = 452.217

	pointsOfPolynomialPlotu[364].X = 3.64
	pointsOfPolynomialPlotu[364].Y = 450.818

	pointsOfPolynomialPlotu[365].X = 3.65
	pointsOfPolynomialPlotu[365].Y = 449.36

	pointsOfPolynomialPlotu[366].X = 3.66
	pointsOfPolynomialPlotu[366].Y = 447.843

	pointsOfPolynomialPlotu[367].X = 3.67
	pointsOfPolynomialPlotu[367].Y = 446.266

	pointsOfPolynomialPlotu[368].X = 3.68
	pointsOfPolynomialPlotu[368].Y = 444.629

	pointsOfPolynomialPlotu[369].X = 3.69
	pointsOfPolynomialPlotu[369].Y = 442.932

	pointsOfPolynomialPlotu[370].X = 3.70
	pointsOfPolynomialPlotu[370].Y = 441.175

	pointsOfPolynomialPlotu[371].X = 3.71
	pointsOfPolynomialPlotu[371].Y = 439.355

	pointsOfPolynomialPlotu[372].X = 3.72
	pointsOfPolynomialPlotu[372].Y = 437.474

	pointsOfPolynomialPlotu[373].X = 3.73
	pointsOfPolynomialPlotu[373].Y = 435.531

	pointsOfPolynomialPlotu[374].X = 3.74
	pointsOfPolynomialPlotu[374].Y = 433.525

	pointsOfPolynomialPlotu[375].X = 3.75
	pointsOfPolynomialPlotu[375].Y = 431.455

	pointsOfPolynomialPlotu[376].X = 3.76
	pointsOfPolynomialPlotu[376].Y = 429.322

	pointsOfPolynomialPlotu[377].X = 3.77
	pointsOfPolynomialPlotu[377].Y = 427.125

	pointsOfPolynomialPlotu[378].X = 3.78
	pointsOfPolynomialPlotu[378].Y = 424.863

	pointsOfPolynomialPlotu[379].X = 3.79
	pointsOfPolynomialPlotu[379].Y = 422.536

	pointsOfPolynomialPlotu[380].X = 3.80
	pointsOfPolynomialPlotu[380].Y = 420.144

	pointsOfPolynomialPlotu[381].X = 3.81
	pointsOfPolynomialPlotu[381].Y = 417.686

	pointsOfPolynomialPlotu[382].X = 3.82
	pointsOfPolynomialPlotu[382].Y = 415.161

	pointsOfPolynomialPlotu[383].X = 3.83
	pointsOfPolynomialPlotu[383].Y = 412.569

	pointsOfPolynomialPlotu[384].X = 3.84
	pointsOfPolynomialPlotu[384].Y = 409.909

	pointsOfPolynomialPlotu[385].X = 3.85
	pointsOfPolynomialPlotu[385].Y = 407.182

	pointsOfPolynomialPlotu[386].X = 3.86
	pointsOfPolynomialPlotu[386].Y = 404.386

	pointsOfPolynomialPlotu[387].X = 3.87
	pointsOfPolynomialPlotu[387].Y = 401.522

	pointsOfPolynomialPlotu[388].X = 3.88
	pointsOfPolynomialPlotu[388].Y = 398.588

	pointsOfPolynomialPlotu[389].X = 3.89
	pointsOfPolynomialPlotu[389].Y = 395.585

	pointsOfPolynomialPlotu[390].X = 3.90
	pointsOfPolynomialPlotu[390].Y = 392.511

	pointsOfPolynomialPlotu[391].X = 3.91
	pointsOfPolynomialPlotu[391].Y = 389.367

	pointsOfPolynomialPlotu[392].X = 3.92
	pointsOfPolynomialPlotu[392].Y = 386.151

	pointsOfPolynomialPlotu[393].X = 3.93
	pointsOfPolynomialPlotu[393].Y = 382.864

	pointsOfPolynomialPlotu[394].X = 3.94
	pointsOfPolynomialPlotu[394].Y = 379.505

	pointsOfPolynomialPlotu[395].X = 3.95
	pointsOfPolynomialPlotu[395].Y = 376.073

	pointsOfPolynomialPlotu[396].X = 3.96
	pointsOfPolynomialPlotu[396].Y = 372.568

	pointsOfPolynomialPlotu[397].X = 3.97
	pointsOfPolynomialPlotu[397].Y = 368.989

	pointsOfPolynomialPlotu[398].X = 3.98
	pointsOfPolynomialPlotu[398].Y = 365.337

	pointsOfPolynomialPlotu[399].X = 3.99
	pointsOfPolynomialPlotu[399].Y = 361.61

	pointsOfPolynomialPlotu[400].X = 4.0
	pointsOfPolynomialPlotu[400].Y = 357.809

	pointsOfPolynomialPlotu[401].X = 4.01
	pointsOfPolynomialPlotu[401].Y = 353.931

	pointsOfPolynomialPlotu[402].X = 4.02
	pointsOfPolynomialPlotu[402].Y = 349.979

	pointsOfPolynomialPlotu[403].X = 4.03
	pointsOfPolynomialPlotu[403].Y = 345.949

	pointsOfPolynomialPlotu[404].X = 4.04
	pointsOfPolynomialPlotu[404].Y = 341.843

	pointsOfPolynomialPlotu[405].X = 4.05
	pointsOfPolynomialPlotu[405].Y = 337.66

	pointsOfPolynomialPlotu[406].X = 4.06
	pointsOfPolynomialPlotu[406].Y = 333.399

	pointsOfPolynomialPlotu[407].X = 4.07
	pointsOfPolynomialPlotu[407].Y = 329.06

	pointsOfPolynomialPlotu[408].X = 4.08
	pointsOfPolynomialPlotu[408].Y = 324.643

	pointsOfPolynomialPlotu[409].X = 4.09
	pointsOfPolynomialPlotu[409].Y = 320.146

	pointsOfPolynomialPlotu[410].X = 4.10
	pointsOfPolynomialPlotu[410].Y = 315.57

	pointsOfPolynomialPlotu[411].X = 4.11
	pointsOfPolynomialPlotu[411].Y = 310.914

	pointsOfPolynomialPlotu[412].X = 4.12
	pointsOfPolynomialPlotu[412].Y = 306.177

	pointsOfPolynomialPlotu[413].X = 4.13
	pointsOfPolynomialPlotu[413].Y = 301.359

	pointsOfPolynomialPlotu[414].X = 4.14
	pointsOfPolynomialPlotu[414].Y = 296.46

	pointsOfPolynomialPlotu[415].X = 4.15
	pointsOfPolynomialPlotu[415].Y = 291.479

	pointsOfPolynomialPlotu[416].X = 4.16
	pointsOfPolynomialPlotu[416].Y = 286.416

	pointsOfPolynomialPlotu[417].X = 4.17
	pointsOfPolynomialPlotu[417].Y = 281.27

	pointsOfPolynomialPlotu[418].X = 4.18
	pointsOfPolynomialPlotu[418].Y = 276.04

	pointsOfPolynomialPlotu[419].X = 4.19
	pointsOfPolynomialPlotu[419].Y = 270.727

	pointsOfPolynomialPlotu[420].X = 4.20
	pointsOfPolynomialPlotu[420].Y = 265.33

	pointsOfPolynomialPlotu[421].X = 4.21
	pointsOfPolynomialPlotu[421].Y = 259.848

	pointsOfPolynomialPlotu[422].X = 4.22
	pointsOfPolynomialPlotu[422].Y = 254.282

	pointsOfPolynomialPlotu[423].X = 4.23
	pointsOfPolynomialPlotu[423].Y = 248.629

	pointsOfPolynomialPlotu[424].X = 4.24
	pointsOfPolynomialPlotu[424].Y = 242.891

	pointsOfPolynomialPlotu[425].X = 4.25
	pointsOfPolynomialPlotu[425].Y = 237.066

	pointsOfPolynomialPlotu[426].X = 4.26
	pointsOfPolynomialPlotu[426].Y = 231.154

	pointsOfPolynomialPlotu[427].X = 4.27
	pointsOfPolynomialPlotu[427].Y = 225.155

	pointsOfPolynomialPlotu[428].X = 4.28
	pointsOfPolynomialPlotu[428].Y = 219.068

	pointsOfPolynomialPlotu[429].X = 4.29
	pointsOfPolynomialPlotu[429].Y = 212.892

	pointsOfPolynomialPlotu[430].X = 4.30
	pointsOfPolynomialPlotu[430].Y = 206.628

	pointsOfPolynomialPlotu[431].X = 4.31
	pointsOfPolynomialPlotu[431].Y = 200.275

	pointsOfPolynomialPlotu[432].X = 4.32
	pointsOfPolynomialPlotu[432].Y = 193.832

	pointsOfPolynomialPlotu[433].X = 4.33
	pointsOfPolynomialPlotu[433].Y = 187.299

	pointsOfPolynomialPlotu[434].X = 4.34
	pointsOfPolynomialPlotu[434].Y = 180.675

	pointsOfPolynomialPlotu[435].X = 4.35
	pointsOfPolynomialPlotu[435].Y = 173.961

	pointsOfPolynomialPlotu[436].X = 4.36
	pointsOfPolynomialPlotu[436].Y = 167.155

	pointsOfPolynomialPlotu[437].X = 4.37
	pointsOfPolynomialPlotu[437].Y = 160.257

	pointsOfPolynomialPlotu[438].X = 4.38
	pointsOfPolynomialPlotu[438].Y = 153.266

	pointsOfPolynomialPlotu[439].X = 4.39
	pointsOfPolynomialPlotu[439].Y = 146.183

	pointsOfPolynomialPlotu[440].X = 4.40
	pointsOfPolynomialPlotu[440].Y = 139.007

	pointsOfPolynomialPlotu[441].X = 4.41
	pointsOfPolynomialPlotu[441].Y = 131.736

	pointsOfPolynomialPlotu[442].X = 4.42
	pointsOfPolynomialPlotu[442].Y = 124.372

	pointsOfPolynomialPlotu[443].X = 4.43
	pointsOfPolynomialPlotu[443].Y = 116.913

	pointsOfPolynomialPlotu[444].X = 4.44
	pointsOfPolynomialPlotu[444].Y = 109.359

	pointsOfPolynomialPlotu[445].X = 4.45
	pointsOfPolynomialPlotu[445].Y = 101.709

	pointsOfPolynomialPlotu[446].X = 4.46
	pointsOfPolynomialPlotu[446].Y = 93.964

	pointsOfPolynomialPlotu[447].X = 4.47
	pointsOfPolynomialPlotu[447].Y = 86.122

	pointsOfPolynomialPlotu[448].X = 4.48
	pointsOfPolynomialPlotu[448].Y = 78.183

	pointsOfPolynomialPlotu[449].X = 4.49
	pointsOfPolynomialPlotu[449].Y = 70.147

	pointsOfPolynomialPlotu[450].X = 4.50
	pointsOfPolynomialPlotu[450].Y = 62.013

	pointsOfPolynomialPlotu[451].X = 4.51
	pointsOfPolynomialPlotu[451].Y = 53.781

	pointsOfPolynomialPlotu[452].X = 4.52
	pointsOfPolynomialPlotu[452].Y = 45.451

	pointsOfPolynomialPlotu[453].X = 4.53
	pointsOfPolynomialPlotu[453].Y = 37.021

	pointsOfPolynomialPlotu[454].X = 4.54
	pointsOfPolynomialPlotu[454].Y = 28.492

	pointsOfPolynomialPlotu[455].X = 4.55
	pointsOfPolynomialPlotu[455].Y = 19.863

	pointsOfPolynomialPlotu[456].X = 4.56
	pointsOfPolynomialPlotu[456].Y = 11.133

	pointsOfPolynomialPlotu[457].X = 4.57
	pointsOfPolynomialPlotu[457].Y = 2.303

	pointsOfPolynomialPlotu[458].X = 4.58
	pointsOfPolynomialPlotu[458].Y = -6.627

	pointsOfPolynomialPlotu[459].X = 4.59
	pointsOfPolynomialPlotu[459].Y = -15.661

	pointsOfPolynomialPlotu[460].X = 4.60
	pointsOfPolynomialPlotu[460].Y = -24.796

	pointsOfPolynomialPlotu[461].X = 4.61
	pointsOfPolynomialPlotu[461].Y = -34.034

	pointsOfPolynomialPlotu[462].X = 4.62
	pointsOfPolynomialPlotu[462].Y = -43.375

	pointsOfPolynomialPlotu[463].X = 4.63
	pointsOfPolynomialPlotu[463].Y = -52.819

	pointsOfPolynomialPlotu[464].X = 4.64
	pointsOfPolynomialPlotu[464].Y = -62.367

	pointsOfPolynomialPlotu[465].X = 4.65
	pointsOfPolynomialPlotu[465].Y = -72.019

	pointsOfPolynomialPlotu[466].X = 4.66
	pointsOfPolynomialPlotu[466].Y = -81.776

	pointsOfPolynomialPlotu[467].X = 4.67
	pointsOfPolynomialPlotu[467].Y = -91.639

	pointsOfPolynomialPlotu[468].X = 4.68
	pointsOfPolynomialPlotu[468].Y = -101.606

	pointsOfPolynomialPlotu[469].X = 4.69
	pointsOfPolynomialPlotu[469].Y = -111.68

	pointsOfPolynomialPlotu[470].X = 4.70
	pointsOfPolynomialPlotu[470].Y = -121.86

	pointsOfPolynomialPlotu[471].X = 4.71
	pointsOfPolynomialPlotu[471].Y = -132.147

	pointsOfPolynomialPlotu[472].X = 4.72
	pointsOfPolynomialPlotu[472].Y = -142.541

	pointsOfPolynomialPlotu[473].X = 4.73
	pointsOfPolynomialPlotu[473].Y = -153.043

	pointsOfPolynomialPlotu[474].X = 4.74
	pointsOfPolynomialPlotu[474].Y = -163.653

	pointsOfPolynomialPlotu[475].X = 4.75
	pointsOfPolynomialPlotu[475].Y = -174.372

	pointsOfPolynomialPlotu[476].X = 4.76
	pointsOfPolynomialPlotu[476].Y = -185.199

	pointsOfPolynomialPlotu[477].X = 4.77
	pointsOfPolynomialPlotu[477].Y = -196.136

	pointsOfPolynomialPlotu[478].X = 4.78
	pointsOfPolynomialPlotu[478].Y = -207.183

	pointsOfPolynomialPlotu[479].X = 4.79
	pointsOfPolynomialPlotu[479].Y = -218.339

	pointsOfPolynomialPlotu[480].X = 4.80
	pointsOfPolynomialPlotu[480].Y = -229.607

	pointsOfPolynomialPlotu[481].X = 4.81
	pointsOfPolynomialPlotu[481].Y = -240.985

	pointsOfPolynomialPlotu[482].X = 4.82
	pointsOfPolynomialPlotu[482].Y = -252.475

	pointsOfPolynomialPlotu[483].X = 4.83
	pointsOfPolynomialPlotu[483].Y = -264.077

	pointsOfPolynomialPlotu[484].X = 4.84
	pointsOfPolynomialPlotu[484].Y = -275.79

	pointsOfPolynomialPlotu[485].X = 4.85
	pointsOfPolynomialPlotu[485].Y = -287.617

	pointsOfPolynomialPlotu[486].X = 4.86
	pointsOfPolynomialPlotu[486].Y = -299.557

	pointsOfPolynomialPlotu[487].X = 4.87
	pointsOfPolynomialPlotu[487].Y = -311.61

	pointsOfPolynomialPlotu[488].X = 4.88
	pointsOfPolynomialPlotu[488].Y = -323.777

	pointsOfPolynomialPlotu[489].X = 4.89
	pointsOfPolynomialPlotu[489].Y = -336.058

	pointsOfPolynomialPlotu[490].X = 4.90
	pointsOfPolynomialPlotu[490].Y = -348.454

	pointsOfPolynomialPlotu[491].X = 4.91
	pointsOfPolynomialPlotu[491].Y = -360.965

	pointsOfPolynomialPlotu[492].X = 4.92
	pointsOfPolynomialPlotu[492].Y = -373.591

	pointsOfPolynomialPlotu[493].X = 4.93
	pointsOfPolynomialPlotu[493].Y = -386.333

	pointsOfPolynomialPlotu[494].X = 4.94
	pointsOfPolynomialPlotu[494].Y = -399.192

	pointsOfPolynomialPlotu[495].X = 4.95
	pointsOfPolynomialPlotu[495].Y = -412.167

	pointsOfPolynomialPlotu[496].X = 4.96
	pointsOfPolynomialPlotu[496].Y = -425.26

	pointsOfPolynomialPlotu[497].X = 4.97
	pointsOfPolynomialPlotu[497].Y = -438.469

	pointsOfPolynomialPlotu[498].X = 4.98
	pointsOfPolynomialPlotu[498].Y = -451.797

	pointsOfPolynomialPlotu[499].X = 4.99
	pointsOfPolynomialPlotu[499].Y = -465.243

	pointsOfPolynomialPlotu[500].X = 5.0
	pointsOfPolynomialPlotu[500].Y = -478.808

	pointsOfPolynomialPlotu[501].X = 5.01
	pointsOfPolynomialPlotu[501].Y = -492.491

	pointsOfPolynomialPlotu[502].X = 5.02
	pointsOfPolynomialPlotu[502].Y = -506.294

	pointsOfPolynomialPlotu[503].X = 5.03
	pointsOfPolynomialPlotu[503].Y = -520.217

	pointsOfPolynomialPlotu[504].X = 5.04
	pointsOfPolynomialPlotu[504].Y = -534.26

	pointsOfPolynomialPlotu[505].X = 5.05
	pointsOfPolynomialPlotu[505].Y = -548.423

	pointsOfPolynomialPlotu[506].X = 5.06
	pointsOfPolynomialPlotu[506].Y = -562.708

	pointsOfPolynomialPlotu[507].X = 5.07
	pointsOfPolynomialPlotu[507].Y = -577.113

	pointsOfPolynomialPlotu[508].X = 5.08
	pointsOfPolynomialPlotu[508].Y = -591.641

	pointsOfPolynomialPlotu[509].X = 5.09
	pointsOfPolynomialPlotu[509].Y = -606.29

	pointsOfPolynomialPlotu[510].X = 5.10
	pointsOfPolynomialPlotu[510].Y = -621.062

	pointsOfPolynomialPlotu[511].X = 5.11
	pointsOfPolynomialPlotu[511].Y = -635.957

	pointsOfPolynomialPlotu[512].X = 5.12
	pointsOfPolynomialPlotu[512].Y = -650.975

	pointsOfPolynomialPlotu[513].X = 5.13
	pointsOfPolynomialPlotu[513].Y = -666.116

	pointsOfPolynomialPlotu[514].X = 5.14
	pointsOfPolynomialPlotu[514].Y = -681.382

	pointsOfPolynomialPlotu[515].X = 5.15
	pointsOfPolynomialPlotu[515].Y = -696.772

	pointsOfPolynomialPlotu[516].X = 5.16
	pointsOfPolynomialPlotu[516].Y = -712.286

	pointsOfPolynomialPlotu[517].X = 5.17
	pointsOfPolynomialPlotu[517].Y = -727.925

	pointsOfPolynomialPlotu[518].X = 5.18
	pointsOfPolynomialPlotu[518].Y = -743.69

	pointsOfPolynomialPlotu[519].X = 5.19
	pointsOfPolynomialPlotu[519].Y = -759.581

	pointsOfPolynomialPlotu[520].X = 5.20
	pointsOfPolynomialPlotu[520].Y = -775.597

	pointsOfPolynomialPlotu[521].X = 5.21
	pointsOfPolynomialPlotu[521].Y = -791.741

	pointsOfPolynomialPlotu[522].X = 5.22
	pointsOfPolynomialPlotu[522].Y = -808.011

	pointsOfPolynomialPlotu[523].X = 5.23
	pointsOfPolynomialPlotu[523].Y = -824.408

	pointsOfPolynomialPlotu[524].X = 5.24
	pointsOfPolynomialPlotu[524].Y = -840.932

	pointsOfPolynomialPlotu[525].X = 5.25
	pointsOfPolynomialPlotu[525].Y = -857.585

	pointsOfPolynomialPlotu[526].X = 5.26
	pointsOfPolynomialPlotu[526].Y = -874.366

	pointsOfPolynomialPlotu[527].X = 5.27
	pointsOfPolynomialPlotu[527].Y = -891.275

	pointsOfPolynomialPlotu[528].X = 5.28
	pointsOfPolynomialPlotu[528].Y = -908.314

	pointsOfPolynomialPlotu[529].X = 5.29
	pointsOfPolynomialPlotu[529].Y = -925.481

	pointsOfPolynomialPlotu[530].X = 5.30
	pointsOfPolynomialPlotu[530].Y = -942.778

	pointsOfPolynomialPlotu[531].X = 5.31
	pointsOfPolynomialPlotu[531].Y = -960.206

	pointsOfPolynomialPlotu[532].X = 5.32
	pointsOfPolynomialPlotu[532].Y = -977.763

	pointsOfPolynomialPlotu[533].X = 5.33
	pointsOfPolynomialPlotu[533].Y = -995.451

	pointsOfPolynomialPlotu[534].X = 5.34
	pointsOfPolynomialPlotu[534].Y = -1_013.27

	pointsOfPolynomialPlotu[535].X = 5.35
	pointsOfPolynomialPlotu[535].Y = -1_031.221

	pointsOfPolynomialPlotu[536].X = 5.36
	pointsOfPolynomialPlotu[536].Y = -1_049.303

	pointsOfPolynomialPlotu[537].X = 5.37
	pointsOfPolynomialPlotu[537].Y = -1_067.517

	pointsOfPolynomialPlotu[538].X = 5.38
	pointsOfPolynomialPlotu[538].Y = -1_085.863

	pointsOfPolynomialPlotu[539].X = 5.39
	pointsOfPolynomialPlotu[539].Y = -1_104.342

	pointsOfPolynomialPlotu[540].X = 5.40
	pointsOfPolynomialPlotu[540].Y = -1_122.954

	pointsOfPolynomialPlotu[541].X = 5.41
	pointsOfPolynomialPlotu[541].Y = -1_141.699

	pointsOfPolynomialPlotu[542].X = 5.42
	pointsOfPolynomialPlotu[542].Y = -1_160.578

	pointsOfPolynomialPlotu[543].X = 5.43
	pointsOfPolynomialPlotu[543].Y = -1_179.59

	pointsOfPolynomialPlotu[544].X = 5.44
	pointsOfPolynomialPlotu[544].Y = -1_198.737

	pointsOfPolynomialPlotu[545].X = 5.45
	pointsOfPolynomialPlotu[545].Y = -1_218.018

	pointsOfPolynomialPlotu[546].X = 5.46
	pointsOfPolynomialPlotu[546].Y = -1_237.435

	pointsOfPolynomialPlotu[547].X = 5.47
	pointsOfPolynomialPlotu[547].Y = -1_256.986

	pointsOfPolynomialPlotu[548].X = 5.48
	pointsOfPolynomialPlotu[548].Y = -1_276.673

	pointsOfPolynomialPlotu[549].X = 5.49
	pointsOfPolynomialPlotu[549].Y = -1_296.496

	pointsOfPolynomialPlotu[550].X = 5.50
	pointsOfPolynomialPlotu[550].Y = -1_316.455

	pointsOfPolynomialPlotu[551].X = 5.51
	pointsOfPolynomialPlotu[551].Y = -1_336.55

	pointsOfPolynomialPlotu[552].X = 5.52
	pointsOfPolynomialPlotu[552].Y = -1_356.782

	pointsOfPolynomialPlotu[553].X = 5.53
	pointsOfPolynomialPlotu[553].Y = -1_377.15

	pointsOfPolynomialPlotu[554].X = 5.54
	pointsOfPolynomialPlotu[554].Y = -1_397.657

	pointsOfPolynomialPlotu[555].X = 5.55
	pointsOfPolynomialPlotu[555].Y = -1_418.3

	pointsOfPolynomialPlotu[556].X = 5.56
	pointsOfPolynomialPlotu[556].Y = -1_439.082

	pointsOfPolynomialPlotu[557].X = 5.57
	pointsOfPolynomialPlotu[557].Y = -1_460.002

	pointsOfPolynomialPlotu[558].X = 5.58
	pointsOfPolynomialPlotu[558].Y = -1_481.06

	pointsOfPolynomialPlotu[559].X = 5.59
	pointsOfPolynomialPlotu[559].Y = -1_502.257

	pointsOfPolynomialPlotu[560].X = 5.60
	pointsOfPolynomialPlotu[560].Y = -1_523.593

	pointsOfPolynomialPlotu[561].X = 5.61
	pointsOfPolynomialPlotu[561].Y = -1_545.069

	pointsOfPolynomialPlotu[562].X = 5.62
	pointsOfPolynomialPlotu[562].Y = -1_566.684

	pointsOfPolynomialPlotu[563].X = 5.63
	pointsOfPolynomialPlotu[563].Y = -1_588.439

	pointsOfPolynomialPlotu[564].X = 5.64
	pointsOfPolynomialPlotu[564].Y = -1_610.334

	pointsOfPolynomialPlotu[565].X = 5.65
	pointsOfPolynomialPlotu[565].Y = -1_632.369

	pointsOfPolynomialPlotu[566].X = 5.66
	pointsOfPolynomialPlotu[566].Y = -1_654.545

	pointsOfPolynomialPlotu[567].X = 5.67
	pointsOfPolynomialPlotu[567].Y = -1_676.863

	pointsOfPolynomialPlotu[568].X = 5.68
	pointsOfPolynomialPlotu[568].Y = -1_699.321

	pointsOfPolynomialPlotu[569].X = 5.69
	pointsOfPolynomialPlotu[569].Y = -1_721.921

	pointsOfPolynomialPlotu[570].X = 5.70
	pointsOfPolynomialPlotu[570].Y = -1_744.663

	pointsOfPolynomialPlotu[571].X = 5.71
	pointsOfPolynomialPlotu[571].Y = -1_767.547

	pointsOfPolynomialPlotu[572].X = 5.72
	pointsOfPolynomialPlotu[572].Y = -1_790.573

	pointsOfPolynomialPlotu[573].X = 5.73
	pointsOfPolynomialPlotu[573].Y = -1_813.741

	pointsOfPolynomialPlotu[574].X = 5.74
	pointsOfPolynomialPlotu[574].Y = -1_837.053

	pointsOfPolynomialPlotu[575].X = 5.75
	pointsOfPolynomialPlotu[575].Y = -1_860.507

	pointsOfPolynomialPlotu[576].X = 5.76
	pointsOfPolynomialPlotu[576].Y = -1_884.105

	pointsOfPolynomialPlotu[577].X = 5.77
	pointsOfPolynomialPlotu[577].Y = -1_907.847

	pointsOfPolynomialPlotu[578].X = 5.78
	pointsOfPolynomialPlotu[578].Y = -1_931.732

	pointsOfPolynomialPlotu[579].X = 5.79
	pointsOfPolynomialPlotu[579].Y = -1_955.761

	pointsOfPolynomialPlotu[580].X = 5.80
	pointsOfPolynomialPlotu[580].Y = -1_979.935

	pointsOfPolynomialPlotu[581].X = 5.81
	pointsOfPolynomialPlotu[581].Y = -2_004.253

	pointsOfPolynomialPlotu[582].X = 5.82
	pointsOfPolynomialPlotu[582].Y = -2_028.716

	pointsOfPolynomialPlotu[583].X = 5.83
	pointsOfPolynomialPlotu[583].Y = -2_053.323

	pointsOfPolynomialPlotu[584].X = 5.84
	pointsOfPolynomialPlotu[584].Y = -2_078.076

	pointsOfPolynomialPlotu[585].X = 5.85
	pointsOfPolynomialPlotu[585].Y = -2_102.975

	pointsOfPolynomialPlotu[586].X = 5.86
	pointsOfPolynomialPlotu[586].Y = -2_128.019

	pointsOfPolynomialPlotu[587].X = 5.87
	pointsOfPolynomialPlotu[587].Y = -2_153.21

	pointsOfPolynomialPlotu[588].X = 5.88
	pointsOfPolynomialPlotu[588].Y = -2_178.546

	pointsOfPolynomialPlotu[589].X = 5.89
	pointsOfPolynomialPlotu[589].Y = -2_204.029

	pointsOfPolynomialPlotu[590].X = 5.90
	pointsOfPolynomialPlotu[590].Y = -2_229.658

	pointsOfPolynomialPlotu[591].X = 5.91
	pointsOfPolynomialPlotu[591].Y = -2_255.434

	pointsOfPolynomialPlotu[592].X = 5.92
	pointsOfPolynomialPlotu[592].Y = -2_281.357

	pointsOfPolynomialPlotu[593].X = 5.93
	pointsOfPolynomialPlotu[593].Y = -2_307.427

	pointsOfPolynomialPlotu[594].X = 5.94
	pointsOfPolynomialPlotu[594].Y = -2_333.645

	pointsOfPolynomialPlotu[595].X = 5.95
	pointsOfPolynomialPlotu[595].Y = -2_360.01

	pointsOfPolynomialPlotu[596].X = 5.96
	pointsOfPolynomialPlotu[596].Y = -2_386.524

	pointsOfPolynomialPlotu[597].X = 5.97
	pointsOfPolynomialPlotu[597].Y = -2_413.185

	pointsOfPolynomialPlotu[598].X = 5.98
	pointsOfPolynomialPlotu[598].Y = -2_439.994

	pointsOfPolynomialPlotu[599].X = 5.99
	pointsOfPolynomialPlotu[599].Y = -2_466.952

	pointsOfPolynomialPlotu[600].X = 6.0
	pointsOfPolynomialPlotu[600].Y = -2_494.059

	pointsOfPolynomialPlotu[601].X = 6.01
	pointsOfPolynomialPlotu[601].Y = -2_521.314

	pointsOfPolynomialPlotu[602].X = 6.02
	pointsOfPolynomialPlotu[602].Y = -2_548.718

	pointsOfPolynomialPlotu[603].X = 6.03
	pointsOfPolynomialPlotu[603].Y = -2_576.272

	pointsOfPolynomialPlotu[604].X = 6.04
	pointsOfPolynomialPlotu[604].Y = -2_603.974

	pointsOfPolynomialPlotu[605].X = 6.05
	pointsOfPolynomialPlotu[605].Y = -2_631.827

	pointsOfPolynomialPlotu[606].X = 6.06
	pointsOfPolynomialPlotu[606].Y = -2_659.829

	pointsOfPolynomialPlotu[607].X = 6.07
	pointsOfPolynomialPlotu[607].Y = -2_687.981

	pointsOfPolynomialPlotu[608].X = 6.08
	pointsOfPolynomialPlotu[608].Y = -2_716.283

	pointsOfPolynomialPlotu[609].X = 6.09
	pointsOfPolynomialPlotu[609].Y = -2_744.735

	pointsOfPolynomialPlotu[610].X = 6.10
	pointsOfPolynomialPlotu[610].Y = -2_773.338

	pointsOfPolynomialPlotu[611].X = 6.11
	pointsOfPolynomialPlotu[611].Y = -2_802.091

	pointsOfPolynomialPlotu[612].X = 6.12
	pointsOfPolynomialPlotu[612].Y = -2_830.995

	pointsOfPolynomialPlotu[613].X = 6.13
	pointsOfPolynomialPlotu[613].Y = -2_860.05

	pointsOfPolynomialPlotu[614].X = 6.14
	pointsOfPolynomialPlotu[614].Y = -2_889.256

	pointsOfPolynomialPlotu[615].X = 6.15
	pointsOfPolynomialPlotu[615].Y = -2_918.613

	pointsOfPolynomialPlotu[616].X = 6.16
	pointsOfPolynomialPlotu[616].Y = -2_948.121

	pointsOfPolynomialPlotu[617].X = 6.17
	pointsOfPolynomialPlotu[617].Y = -2_977.781

	pointsOfPolynomialPlotu[618].X = 6.18
	pointsOfPolynomialPlotu[618].Y = -3_007.593

	pointsOfPolynomialPlotu[619].X = 6.19
	pointsOfPolynomialPlotu[619].Y = -3_037.556

	pointsOfPolynomialPlotu[620].X = 6.20
	pointsOfPolynomialPlotu[620].Y = -3_067.671

	pointsOfPolynomialPlotu[621].X = 6.21
	pointsOfPolynomialPlotu[621].Y = -3_097.938

	pointsOfPolynomialPlotu[622].X = 6.22
	pointsOfPolynomialPlotu[622].Y = -3_128.358

	pointsOfPolynomialPlotu[623].X = 6.23
	pointsOfPolynomialPlotu[623].Y = -3_158.929

	pointsOfPolynomialPlotu[624].X = 6.24
	pointsOfPolynomialPlotu[624].Y = -3_189.654

	pointsOfPolynomialPlotu[625].X = 6.25
	pointsOfPolynomialPlotu[625].Y = -3_220.53

	pointsOfPolynomialPlotu[626].X = 6.26
	pointsOfPolynomialPlotu[626].Y = -3_251.56

	pointsOfPolynomialPlotu[627].X = 6.27
	pointsOfPolynomialPlotu[627].Y = -3_282.742

	pointsOfPolynomialPlotu[628].X = 6.28
	pointsOfPolynomialPlotu[628].Y = -3_314.077

	pointsOfPolynomialPlotu[629].X = 6.29
	pointsOfPolynomialPlotu[629].Y = -3_345.565

	pointsOfPolynomialPlotu[630].X = 6.30
	pointsOfPolynomialPlotu[630].Y = -3_377.207

	pointsOfPolynomialPlotu[631].X = 6.31
	pointsOfPolynomialPlotu[631].Y = -3_409.001

	pointsOfPolynomialPlotu[632].X = 6.32
	pointsOfPolynomialPlotu[632].Y = -3_440.949

	pointsOfPolynomialPlotu[633].X = 6.33
	pointsOfPolynomialPlotu[633].Y = -3_473.714

	pointsOfPolynomialPlotu[634].X = 6.34
	pointsOfPolynomialPlotu[634].Y = -3_505.276

	pointsOfPolynomialPlotu[635].X = 6.35
	pointsOfPolynomialPlotu[635].Y = -3_537.714

	pointsOfPolynomialPlotu[636].X = 6.36
	pointsOfPolynomialPlotu[636].Y = -3_570.276

	pointsOfPolynomialPlotu[637].X = 6.37
	pointsOfPolynomialPlotu[637].Y = -3_602.992

	pointsOfPolynomialPlotu[638].X = 6.38
	pointsOfPolynomialPlotu[638].Y = -3_635.862

	pointsOfPolynomialPlotu[639].X = 6.39
	pointsOfPolynomialPlotu[639].Y = -3_668.886

	pointsOfPolynomialPlotu[640].X = 6.40
	pointsOfPolynomialPlotu[640].Y = -3_702.064

	pointsOfPolynomialPlotu[641].X = 6.41
	pointsOfPolynomialPlotu[641].Y = -3_735.396

	pointsOfPolynomialPlotu[642].X = 6.42
	pointsOfPolynomialPlotu[642].Y = -3_768.883

	pointsOfPolynomialPlotu[643].X = 6.43
	pointsOfPolynomialPlotu[643].Y = -3_802.523

	pointsOfPolynomialPlotu[644].X = 6.44
	pointsOfPolynomialPlotu[644].Y = -3_836.319

	pointsOfPolynomialPlotu[645].X = 6.45
	pointsOfPolynomialPlotu[645].Y = -3_870.268

	pointsOfPolynomialPlotu[646].X = 6.46
	pointsOfPolynomialPlotu[646].Y = -3_904.372

	pointsOfPolynomialPlotu[647].X = 6.47
	pointsOfPolynomialPlotu[647].Y = -3_938.631

	pointsOfPolynomialPlotu[648].X = 6.48
	pointsOfPolynomialPlotu[648].Y = -3_973.044

	pointsOfPolynomialPlotu[649].X = 6.49
	pointsOfPolynomialPlotu[649].Y = -4_007.612

	pointsOfPolynomialPlotu[650].X = 6.50
	pointsOfPolynomialPlotu[650].Y = -4_042.334

	pointsOfPolynomialPlotu[651].X = 6.51
	pointsOfPolynomialPlotu[651].Y = -4_077.212

	pointsOfPolynomialPlotu[652].X = 6.52
	pointsOfPolynomialPlotu[652].Y = -4_112.244

	pointsOfPolynomialPlotu[653].X = 6.53
	pointsOfPolynomialPlotu[653].Y = -4_147.431

	pointsOfPolynomialPlotu[654].X = 6.54
	pointsOfPolynomialPlotu[654].Y = -4_182.772

	pointsOfPolynomialPlotu[655].X = 6.55
	pointsOfPolynomialPlotu[655].Y = -4_218.269

	pointsOfPolynomialPlotu[656].X = 6.56
	pointsOfPolynomialPlotu[656].Y = -4_253.921

	pointsOfPolynomialPlotu[657].X = 6.57
	pointsOfPolynomialPlotu[657].Y = -4_289.727

	pointsOfPolynomialPlotu[658].X = 6.58
	pointsOfPolynomialPlotu[658].Y = -4_325.689

	pointsOfPolynomialPlotu[659].X = 6.59
	pointsOfPolynomialPlotu[659].Y = -4_361.805

	pointsOfPolynomialPlotu[660].X = 6.60
	pointsOfPolynomialPlotu[660].Y = -4_398.077

	pointsOfPolynomialPlotu[661].X = 6.61
	pointsOfPolynomialPlotu[661].Y = -4_434.504

	pointsOfPolynomialPlotu[662].X = 6.62
	pointsOfPolynomialPlotu[662].Y = -4_471.085

	pointsOfPolynomialPlotu[663].X = 6.63
	pointsOfPolynomialPlotu[663].Y = -4_507.822

	pointsOfPolynomialPlotu[664].X = 6.64
	pointsOfPolynomialPlotu[664].Y = -4_544.714

	pointsOfPolynomialPlotu[665].X = 6.65
	pointsOfPolynomialPlotu[665].Y = -4_581.761

	pointsOfPolynomialPlotu[666].X = 6.66
	pointsOfPolynomialPlotu[666].Y = -4_618.962

	pointsOfPolynomialPlotu[667].X = 6.67
	pointsOfPolynomialPlotu[667].Y = -4_656.319

	pointsOfPolynomialPlotu[668].X = 6.68
	pointsOfPolynomialPlotu[668].Y = -4_693.831

	pointsOfPolynomialPlotu[669].X = 6.69
	pointsOfPolynomialPlotu[669].Y = -4_731.498

	pointsOfPolynomialPlotu[670].X = 6.70
	pointsOfPolynomialPlotu[670].Y = -4_769.32

	pointsOfPolynomialPlotu[671].X = 6.71
	pointsOfPolynomialPlotu[671].Y = -4_807.297

	pointsOfPolynomialPlotu[672].X = 6.72
	pointsOfPolynomialPlotu[672].Y = -4_845.429

	pointsOfPolynomialPlotu[673].X = 6.73
	pointsOfPolynomialPlotu[673].Y = -4_883.716

	pointsOfPolynomialPlotu[674].X = 6.74
	pointsOfPolynomialPlotu[674].Y = -4_922.158

	pointsOfPolynomialPlotu[675].X = 6.75
	pointsOfPolynomialPlotu[675].Y = -4_960.754

	pointsOfPolynomialPlotu[676].X = 6.76
	pointsOfPolynomialPlotu[676].Y = -4_999.506

	pointsOfPolynomialPlotu[677].X = 6.77
	pointsOfPolynomialPlotu[677].Y = -5_038.412

	pointsOfPolynomialPlotu[678].X = 6.78
	pointsOfPolynomialPlotu[678].Y = -5_077.473

	pointsOfPolynomialPlotu[679].X = 6.79
	pointsOfPolynomialPlotu[679].Y = -5_116.689

	pointsOfPolynomialPlotu[680].X = 6.80
	pointsOfPolynomialPlotu[680].Y = -5_156.059

	pointsOfPolynomialPlotu[681].X = 6.81
	pointsOfPolynomialPlotu[681].Y = -5_195.584

	pointsOfPolynomialPlotu[682].X = 6.82
	pointsOfPolynomialPlotu[682].Y = -5_235.263

	pointsOfPolynomialPlotu[683].X = 6.83
	pointsOfPolynomialPlotu[683].Y = -5_275.097

	pointsOfPolynomialPlotu[684].X = 6.84
	pointsOfPolynomialPlotu[684].Y = -5_315.085

	pointsOfPolynomialPlotu[685].X = 6.85
	pointsOfPolynomialPlotu[685].Y = -5_355.228

	pointsOfPolynomialPlotu[686].X = 6.86
	pointsOfPolynomialPlotu[686].Y = -5_395.525

	pointsOfPolynomialPlotu[687].X = 6.87
	pointsOfPolynomialPlotu[687].Y = -5_435.976

	pointsOfPolynomialPlotu[688].X = 6.88
	pointsOfPolynomialPlotu[688].Y = -5_476.581

	pointsOfPolynomialPlotu[689].X = 6.89
	pointsOfPolynomialPlotu[689].Y = -5_517.34

	pointsOfPolynomialPlotu[690].X = 6.90
	pointsOfPolynomialPlotu[690].Y = -5_558.254

	pointsOfPolynomialPlotu[691].X = 6.91
	pointsOfPolynomialPlotu[691].Y = -5_599.321

	pointsOfPolynomialPlotu[692].X = 6.92
	pointsOfPolynomialPlotu[692].Y = -5_640.541

	pointsOfPolynomialPlotu[693].X = 6.93
	pointsOfPolynomialPlotu[693].Y = -5_681.916

	pointsOfPolynomialPlotu[694].X = 6.94
	pointsOfPolynomialPlotu[694].Y = -5_723.444

	pointsOfPolynomialPlotu[695].X = 6.95
	pointsOfPolynomialPlotu[695].Y = -5_765.125

	pointsOfPolynomialPlotu[696].X = 6.96
	pointsOfPolynomialPlotu[696].Y = -5_806.96

	pointsOfPolynomialPlotu[697].X = 6.97
	pointsOfPolynomialPlotu[697].Y = -5_848.948

	pointsOfPolynomialPlotu[698].X = 6.98
	pointsOfPolynomialPlotu[698].Y = -5_891.089

	pointsOfPolynomialPlotu[699].X = 6.99
	pointsOfPolynomialPlotu[699].Y = -5_933.383

	pointsOfPolynomialPlotu[700].X = 7.0
	pointsOfPolynomialPlotu[700].Y = -5_975.83

	pointsOfPolynomialPlotu[701].X = 7.01
	pointsOfPolynomialPlotu[701].Y = -6_018.429

	pointsOfPolynomialPlotu[702].X = 7.02
	pointsOfPolynomialPlotu[702].Y = -6_061.181

	pointsOfPolynomialPlotu[703].X = 7.03
	pointsOfPolynomialPlotu[703].Y = -6_104.086

	pointsOfPolynomialPlotu[704].X = 7.04
	pointsOfPolynomialPlotu[704].Y = -6_147.143

	pointsOfPolynomialPlotu[705].X = 7.05
	pointsOfPolynomialPlotu[705].Y = -6_190.352

	pointsOfPolynomialPlotu[706].X = 7.06
	pointsOfPolynomialPlotu[706].Y = -6_233.713

	pointsOfPolynomialPlotu[707].X = 7.07
	pointsOfPolynomialPlotu[707].Y = -6_277.226

	pointsOfPolynomialPlotu[708].X = 7.08
	pointsOfPolynomialPlotu[708].Y = -6_320.891

	pointsOfPolynomialPlotu[709].X = 7.09
	pointsOfPolynomialPlotu[709].Y = -6_364.707

	pointsOfPolynomialPlotu[710].X = 7.10
	pointsOfPolynomialPlotu[710].Y = -6_408.675

	pointsOfPolynomialPlotu[711].X = 7.11
	pointsOfPolynomialPlotu[711].Y = -6_452.794

	pointsOfPolynomialPlotu[712].X = 7.12
	pointsOfPolynomialPlotu[712].Y = -6_497.064

	pointsOfPolynomialPlotu[713].X = 7.13
	pointsOfPolynomialPlotu[713].Y = -6_541.485

	pointsOfPolynomialPlotu[714].X = 7.14
	pointsOfPolynomialPlotu[714].Y = -6_586.057

	pointsOfPolynomialPlotu[715].X = 7.15
	pointsOfPolynomialPlotu[715].Y = -6_630.779

	pointsOfPolynomialPlotu[716].X = 7.16
	pointsOfPolynomialPlotu[716].Y = -6_675.652

	pointsOfPolynomialPlotu[717].X = 7.17
	pointsOfPolynomialPlotu[717].Y = -6_720.674

	pointsOfPolynomialPlotu[718].X = 7.18
	pointsOfPolynomialPlotu[718].Y = -6_765.847

	pointsOfPolynomialPlotu[719].X = 7.19
	pointsOfPolynomialPlotu[719].Y = -6_811.169

	pointsOfPolynomialPlotu[720].X = 7.20
	pointsOfPolynomialPlotu[720].Y = -6_856.641

	pointsOfPolynomialPlotu[721].X = 7.21
	pointsOfPolynomialPlotu[721].Y = -6_902.263

	pointsOfPolynomialPlotu[722].X = 7.22
	pointsOfPolynomialPlotu[722].Y = -6_948.033

	pointsOfPolynomialPlotu[723].X = 7.23
	pointsOfPolynomialPlotu[723].Y = -6_993.952

	pointsOfPolynomialPlotu[724].X = 7.24
	pointsOfPolynomialPlotu[724].Y = -7_040.021

	pointsOfPolynomialPlotu[725].X = 7.25
	pointsOfPolynomialPlotu[725].Y = -7_086.237

	pointsOfPolynomialPlotu[726].X = 7.26
	pointsOfPolynomialPlotu[726].Y = -7_132.602

	pointsOfPolynomialPlotu[727].X = 7.27
	pointsOfPolynomialPlotu[727].Y = -7_179.115

	pointsOfPolynomialPlotu[728].X = 7.28
	pointsOfPolynomialPlotu[728].Y = -7_225.776

	pointsOfPolynomialPlotu[729].X = 7.29
	pointsOfPolynomialPlotu[729].Y = -7_272.584

	pointsOfPolynomialPlotu[730].X = 7.30
	pointsOfPolynomialPlotu[730].Y = -7_319.539

	pointsOfPolynomialPlotu[731].X = 7.31
	pointsOfPolynomialPlotu[731].Y = -7_366.642

	pointsOfPolynomialPlotu[732].X = 7.32
	pointsOfPolynomialPlotu[732].Y = -7_413.891

	pointsOfPolynomialPlotu[733].X = 7.33
	pointsOfPolynomialPlotu[733].Y = -7_461.287

	pointsOfPolynomialPlotu[734].X = 7.34
	pointsOfPolynomialPlotu[734].Y = -7_508.93

	pointsOfPolynomialPlotu[735].X = 7.35
	pointsOfPolynomialPlotu[735].Y = -7_556.518

	pointsOfPolynomialPlotu[736].X = 7.36
	pointsOfPolynomialPlotu[736].Y = -7_604.352

	pointsOfPolynomialPlotu[737].X = 7.37
	pointsOfPolynomialPlotu[737].Y = -7_652.332

	pointsOfPolynomialPlotu[738].X = 7.38
	pointsOfPolynomialPlotu[738].Y = -7_700.456

	pointsOfPolynomialPlotu[739].X = 7.39
	pointsOfPolynomialPlotu[739].Y = -7_748.726

	pointsOfPolynomialPlotu[740].X = 7.40
	pointsOfPolynomialPlotu[740].Y = -7_797.141

	pointsOfPolynomialPlotu[741].X = 7.41
	pointsOfPolynomialPlotu[741].Y = -7_845.699

	pointsOfPolynomialPlotu[742].X = 7.42
	pointsOfPolynomialPlotu[742].Y = -7_894.402

	pointsOfPolynomialPlotu[743].X = 7.43
	pointsOfPolynomialPlotu[743].Y = -7_943.249

	pointsOfPolynomialPlotu[744].X = 7.44
	pointsOfPolynomialPlotu[744].Y = -7_992.239

	pointsOfPolynomialPlotu[745].X = 7.45
	pointsOfPolynomialPlotu[745].Y = -8_041.372

	pointsOfPolynomialPlotu[746].X = 7.46
	pointsOfPolynomialPlotu[746].Y = -8_090.648

	pointsOfPolynomialPlotu[747].X = 7.47
	pointsOfPolynomialPlotu[747].Y = -8_140.067

	pointsOfPolynomialPlotu[748].X = 7.48
	pointsOfPolynomialPlotu[748].Y = -8_189.628

	pointsOfPolynomialPlotu[749].X = 7.49
	pointsOfPolynomialPlotu[749].Y = -8_239.331

	pointsOfPolynomialPlotu[750].X = 7.50
	pointsOfPolynomialPlotu[750].Y = -8_289.175

	pointsOfPolynomialPlotu[751].X = 7.51
	pointsOfPolynomialPlotu[751].Y = -8_339.161

	pointsOfPolynomialPlotu[752].X = 7.52
	pointsOfPolynomialPlotu[752].Y = -8_389.287

	pointsOfPolynomialPlotu[753].X = 7.53
	pointsOfPolynomialPlotu[753].Y = -8_439.554

	pointsOfPolynomialPlotu[754].X = 7.54
	pointsOfPolynomialPlotu[754].Y = -8_489.961

	pointsOfPolynomialPlotu[755].X = 7.55
	pointsOfPolynomialPlotu[755].Y = -8_540.509

	pointsOfPolynomialPlotu[756].X = 7.56
	pointsOfPolynomialPlotu[756].Y = -8_591.195

	pointsOfPolynomialPlotu[757].X = 7.57
	pointsOfPolynomialPlotu[757].Y = -8_642.021

	pointsOfPolynomialPlotu[758].X = 7.58
	pointsOfPolynomialPlotu[758].Y = -8_692.986

	pointsOfPolynomialPlotu[759].X = 7.59
	pointsOfPolynomialPlotu[759].Y = -8_744.089

	pointsOfPolynomialPlotu[760].X = 7.60
	pointsOfPolynomialPlotu[760].Y = -8_795.33

	pointsOfPolynomialPlotu[761].X = 7.61
	pointsOfPolynomialPlotu[761].Y = -8_846.709

	pointsOfPolynomialPlotu[762].X = 7.62
	pointsOfPolynomialPlotu[762].Y = -8_898.225

	pointsOfPolynomialPlotu[763].X = 7.63
	pointsOfPolynomialPlotu[763].Y = -8_949.877

	pointsOfPolynomialPlotu[764].X = 7.64
	pointsOfPolynomialPlotu[764].Y = -9_001.667

	pointsOfPolynomialPlotu[765].X = 7.65
	pointsOfPolynomialPlotu[765].Y = -9_053.593

	pointsOfPolynomialPlotu[766].X = 7.66
	pointsOfPolynomialPlotu[766].Y = -9_105.654

	pointsOfPolynomialPlotu[767].X = 7.67
	pointsOfPolynomialPlotu[767].Y = -9_157.85

	pointsOfPolynomialPlotu[768].X = 7.68
	pointsOfPolynomialPlotu[768].Y = -9_210.182

	pointsOfPolynomialPlotu[769].X = 7.69
	pointsOfPolynomialPlotu[769].Y = -9_262.648

	pointsOfPolynomialPlotu[770].X = 7.70
	pointsOfPolynomialPlotu[770].Y = -9_315.248

	pointsOfPolynomialPlotu[771].X = 7.71
	pointsOfPolynomialPlotu[771].Y = -9_367.982

	pointsOfPolynomialPlotu[772].X = 7.72
	pointsOfPolynomialPlotu[772].Y = -9_420.849

	pointsOfPolynomialPlotu[773].X = 7.73
	pointsOfPolynomialPlotu[773].Y = -9_473.849

	pointsOfPolynomialPlotu[774].X = 7.74
	pointsOfPolynomialPlotu[774].Y = -9_526.981

	pointsOfPolynomialPlotu[775].X = 7.75
	pointsOfPolynomialPlotu[775].Y = -9_580.245

	pointsOfPolynomialPlotu[776].X = 7.76
	pointsOfPolynomialPlotu[776].Y = -9_633.64

	pointsOfPolynomialPlotu[777].X = 7.77
	pointsOfPolynomialPlotu[777].Y = -9_687.167

	pointsOfPolynomialPlotu[778].X = 7.78
	pointsOfPolynomialPlotu[778].Y = -9_740.824

	pointsOfPolynomialPlotu[779].X = 7.79
	pointsOfPolynomialPlotu[779].Y = -9_794.611

	pointsOfPolynomialPlotu[780].X = 7.80
	pointsOfPolynomialPlotu[780].Y = -9_848.528

	pointsOfPolynomialPlotu[781].X = 7.81
	pointsOfPolynomialPlotu[781].Y = -9_902.574

	pointsOfPolynomialPlotu[782].X = 7.82
	pointsOfPolynomialPlotu[782].Y = -9_956.749

	pointsOfPolynomialPlotu[783].X = 7.83
	pointsOfPolynomialPlotu[783].Y = -10_011.052

	pointsOfPolynomialPlotu[784].X = 7.84
	pointsOfPolynomialPlotu[784].Y = -10_065.483

	pointsOfPolynomialPlotu[785].X = 7.85
	pointsOfPolynomialPlotu[785].Y = -10_120.041

	pointsOfPolynomialPlotu[786].X = 7.86
	pointsOfPolynomialPlotu[786].Y = -10_174.726

	pointsOfPolynomialPlotu[787].X = 7.87
	pointsOfPolynomialPlotu[787].Y = -10_229.537

	pointsOfPolynomialPlotu[788].X = 7.88
	pointsOfPolynomialPlotu[788].Y = -10_284.473

	pointsOfPolynomialPlotu[789].X = 7.89
	pointsOfPolynomialPlotu[789].Y = -10_339.535

	pointsOfPolynomialPlotu[790].X = 7.90
	pointsOfPolynomialPlotu[790].Y = -10_394.722

	pointsOfPolynomialPlotu[791].X = 7.91
	pointsOfPolynomialPlotu[791].Y = -10_450.033

	pointsOfPolynomialPlotu[792].X = 7.92
	pointsOfPolynomialPlotu[792].Y = -10_505.468

	pointsOfPolynomialPlotu[793].X = 7.93
	pointsOfPolynomialPlotu[793].Y = -10_561.026

	pointsOfPolynomialPlotu[794].X = 7.94
	pointsOfPolynomialPlotu[794].Y = -10_616.706

	pointsOfPolynomialPlotu[795].X = 7.95
	pointsOfPolynomialPlotu[795].Y = -10_672.509

	pointsOfPolynomialPlotu[796].X = 7.96
	pointsOfPolynomialPlotu[796].Y = -10_728.433

	pointsOfPolynomialPlotu[797].X = 7.97
	pointsOfPolynomialPlotu[797].Y = -10_784.478

	pointsOfPolynomialPlotu[798].X = 7.98
	pointsOfPolynomialPlotu[798].Y = -10_840.644

	pointsOfPolynomialPlotu[799].X = 7.99
	pointsOfPolynomialPlotu[799].Y = -10_896.93

	pointsOfPolynomialPlotu[800].X = 8.0
	pointsOfPolynomialPlotu[800].Y = -10_953.335

	pointsOfPolynomialPlotu[801].X = 8.01
	pointsOfPolynomialPlotu[801].Y = -11_009.858

	pointsOfPolynomialPlotu[802].X = 8.02
	pointsOfPolynomialPlotu[802].Y = -11_066.5

	pointsOfPolynomialPlotu[803].X = 8.03
	pointsOfPolynomialPlotu[803].Y = -11_123.26

	pointsOfPolynomialPlotu[804].X = 8.04
	pointsOfPolynomialPlotu[804].Y = -11_180.136

	pointsOfPolynomialPlotu[805].X = 8.05
	pointsOfPolynomialPlotu[805].Y = -11_237.129

	pointsOfPolynomialPlotu[806].X = 8.06
	pointsOfPolynomialPlotu[806].Y = -11_294.238

	pointsOfPolynomialPlotu[807].X = 8.07
	pointsOfPolynomialPlotu[807].Y = -11_351.463

	pointsOfPolynomialPlotu[808].X = 8.08
	pointsOfPolynomialPlotu[808].Y = -11_408.802

	pointsOfPolynomialPlotu[809].X = 8.09
	pointsOfPolynomialPlotu[809].Y = -11_466.255

	pointsOfPolynomialPlotu[810].X = 8.10
	pointsOfPolynomialPlotu[810].Y = -11_523.821

	pointsOfPolynomialPlotu[811].X = 8.11
	pointsOfPolynomialPlotu[811].Y = -11_581.5

	pointsOfPolynomialPlotu[812].X = 8.12
	pointsOfPolynomialPlotu[812].Y = -11_639.292

	pointsOfPolynomialPlotu[813].X = 8.13
	pointsOfPolynomialPlotu[813].Y = -11_697.195

	pointsOfPolynomialPlotu[814].X = 8.14
	pointsOfPolynomialPlotu[814].Y = -11_755.209

	pointsOfPolynomialPlotu[815].X = 8.15
	pointsOfPolynomialPlotu[815].Y = -11_813.333

	pointsOfPolynomialPlotu[816].X = 8.16
	pointsOfPolynomialPlotu[816].Y = -11_871.568

	pointsOfPolynomialPlotu[817].X = 8.17
	pointsOfPolynomialPlotu[817].Y = -11_929.911

	pointsOfPolynomialPlotu[818].X = 8.18
	pointsOfPolynomialPlotu[818].Y = -11_988.362

	pointsOfPolynomialPlotu[819].X = 8.19
	pointsOfPolynomialPlotu[819].Y = -12_046.921

	pointsOfPolynomialPlotu[820].X = 8.20
	pointsOfPolynomialPlotu[820].Y = -12_105.588

	pointsOfPolynomialPlotu[821].X = 8.21
	pointsOfPolynomialPlotu[821].Y = -12_164.36

	pointsOfPolynomialPlotu[822].X = 8.22
	pointsOfPolynomialPlotu[822].Y = -12_223.239

	pointsOfPolynomialPlotu[823].X = 8.23
	pointsOfPolynomialPlotu[823].Y = -12_282.222

	pointsOfPolynomialPlotu[824].X = 8.24
	pointsOfPolynomialPlotu[824].Y = -12_341.31

	pointsOfPolynomialPlotu[825].X = 8.25
	pointsOfPolynomialPlotu[825].Y = -12_400.501

	pointsOfPolynomialPlotu[826].X = 8.26
	pointsOfPolynomialPlotu[826].Y = -12_459.796

	pointsOfPolynomialPlotu[827].X = 8.27
	pointsOfPolynomialPlotu[827].Y = -12_519.193

	pointsOfPolynomialPlotu[828].X = 8.28
	pointsOfPolynomialPlotu[828].Y = -12_578.691

	pointsOfPolynomialPlotu[829].X = 8.29
	pointsOfPolynomialPlotu[829].Y = -12_638.29

	pointsOfPolynomialPlotu[830].X = 8.30
	pointsOfPolynomialPlotu[830].Y = -12_697.989

	pointsOfPolynomialPlotu[831].X = 8.31
	pointsOfPolynomialPlotu[831].Y = -12_757.788

	pointsOfPolynomialPlotu[832].X = 8.32
	pointsOfPolynomialPlotu[832].Y = -12_817.685

	pointsOfPolynomialPlotu[833].X = 8.33
	pointsOfPolynomialPlotu[833].Y = -12_877.681

	pointsOfPolynomialPlotu[834].X = 8.34
	pointsOfPolynomialPlotu[834].Y = -12_937.773

	pointsOfPolynomialPlotu[835].X = 8.35
	pointsOfPolynomialPlotu[835].Y = -12_997.962

	pointsOfPolynomialPlotu[836].X = 8.36
	pointsOfPolynomialPlotu[836].Y = -13_058.247

	pointsOfPolynomialPlotu[837].X = 8.37
	pointsOfPolynomialPlotu[837].Y = -13_118.627

	pointsOfPolynomialPlotu[838].X = 8.38
	pointsOfPolynomialPlotu[838].Y = -13_179.101

	pointsOfPolynomialPlotu[839].X = 8.39
	pointsOfPolynomialPlotu[839].Y = -13_239.668

	pointsOfPolynomialPlotu[840].X = 8.40
	pointsOfPolynomialPlotu[840].Y = -13_300.329

	pointsOfPolynomialPlotu[841].X = 8.41
	pointsOfPolynomialPlotu[841].Y = -13_361.081

	pointsOfPolynomialPlotu[842].X = 8.42
	pointsOfPolynomialPlotu[842].Y = -13_421.924

	pointsOfPolynomialPlotu[843].X = 8.43
	pointsOfPolynomialPlotu[843].Y = -13_482.857

	pointsOfPolynomialPlotu[844].X = 8.44
	pointsOfPolynomialPlotu[844].Y = -13_543.881

	pointsOfPolynomialPlotu[845].X = 8.45
	pointsOfPolynomialPlotu[845].Y = -13_604.992

	pointsOfPolynomialPlotu[846].X = 8.46
	pointsOfPolynomialPlotu[846].Y = -13_666.192

	pointsOfPolynomialPlotu[847].X = 8.47
	pointsOfPolynomialPlotu[847].Y = -13_727.479

	pointsOfPolynomialPlotu[848].X = 8.48
	pointsOfPolynomialPlotu[848].Y = -13_788.852

	pointsOfPolynomialPlotu[849].X = 8.49
	pointsOfPolynomialPlotu[849].Y = -13_850.311

	pointsOfPolynomialPlotu[850].X = 8.50
	pointsOfPolynomialPlotu[850].Y = -13_911.855

	pointsOfPolynomialPlotu[851].X = 8.51
	pointsOfPolynomialPlotu[851].Y = -13_973.482

	pointsOfPolynomialPlotu[852].X = 8.52
	pointsOfPolynomialPlotu[852].Y = -14_035.193

	pointsOfPolynomialPlotu[853].X = 8.53
	pointsOfPolynomialPlotu[853].Y = -14_096.985

	pointsOfPolynomialPlotu[854].X = 8.54
	pointsOfPolynomialPlotu[854].Y = -14_158.859

	pointsOfPolynomialPlotu[855].X = 8.55
	pointsOfPolynomialPlotu[855].Y = -14_220.813

	pointsOfPolynomialPlotu[856].X = 8.56
	pointsOfPolynomialPlotu[856].Y = -14_282.847

	pointsOfPolynomialPlotu[857].X = 8.57
	pointsOfPolynomialPlotu[857].Y = -14_344.96

	pointsOfPolynomialPlotu[858].X = 8.58
	pointsOfPolynomialPlotu[858].Y = -14_407.151

	pointsOfPolynomialPlotu[859].X = 8.59
	pointsOfPolynomialPlotu[859].Y = -14_469.419

	pointsOfPolynomialPlotu[860].X = 8.60
	pointsOfPolynomialPlotu[860].Y = -14_531.762

	pointsOfPolynomialPlotu[861].X = 8.61
	pointsOfPolynomialPlotu[861].Y = -14_594.181

	pointsOfPolynomialPlotu[862].X = 8.62
	pointsOfPolynomialPlotu[862].Y = -14_656.675

	pointsOfPolynomialPlotu[863].X = 8.63
	pointsOfPolynomialPlotu[863].Y = -14_719.242

	pointsOfPolynomialPlotu[864].X = 8.64
	pointsOfPolynomialPlotu[864].Y = -14_781.881

	pointsOfPolynomialPlotu[865].X = 8.65
	pointsOfPolynomialPlotu[865].Y = -14_844.592

	pointsOfPolynomialPlotu[866].X = 8.66
	pointsOfPolynomialPlotu[866].Y = -14_907.374

	pointsOfPolynomialPlotu[867].X = 8.67
	pointsOfPolynomialPlotu[867].Y = -14_970.226

	pointsOfPolynomialPlotu[868].X = 8.68
	pointsOfPolynomialPlotu[868].Y = -15_033.146

	pointsOfPolynomialPlotu[869].X = 8.69
	pointsOfPolynomialPlotu[869].Y = -15_096.134

	pointsOfPolynomialPlotu[870].X = 8.70
	pointsOfPolynomialPlotu[870].Y = -15_159.14

	pointsOfPolynomialPlotu[871].X = 8.71
	pointsOfPolynomialPlotu[871].Y = -15_222.311

	pointsOfPolynomialPlotu[872].X = 8.72
	pointsOfPolynomialPlotu[872].Y = -15_285.498

	pointsOfPolynomialPlotu[873].X = 8.73
	pointsOfPolynomialPlotu[873].Y = -15_348.749

	pointsOfPolynomialPlotu[874].X = 8.74
	pointsOfPolynomialPlotu[874].Y = -15_412.063

	pointsOfPolynomialPlotu[875].X = 8.75
	pointsOfPolynomialPlotu[875].Y = -15_475.439

	pointsOfPolynomialPlotu[876].X = 8.76
	pointsOfPolynomialPlotu[876].Y = -15_538.876

	pointsOfPolynomialPlotu[877].X = 8.77
	pointsOfPolynomialPlotu[877].Y = -15_602.374

	pointsOfPolynomialPlotu[878].X = 8.78
	pointsOfPolynomialPlotu[878].Y = -15_665.931

	pointsOfPolynomialPlotu[879].X = 8.79
	pointsOfPolynomialPlotu[879].Y = -15_729.546

	pointsOfPolynomialPlotu[880].X = 8.80
	pointsOfPolynomialPlotu[880].Y = -15_793.219

	pointsOfPolynomialPlotu[881].X = 8.81
	pointsOfPolynomialPlotu[881].Y = -15_856.948

	pointsOfPolynomialPlotu[882].X = 8.82
	pointsOfPolynomialPlotu[882].Y = -15_920.732

	pointsOfPolynomialPlotu[883].X = 8.83
	pointsOfPolynomialPlotu[883].Y = -15_984.571

	pointsOfPolynomialPlotu[884].X = 8.84
	pointsOfPolynomialPlotu[884].Y = -16_048.462

	pointsOfPolynomialPlotu[885].X = 8.85
	pointsOfPolynomialPlotu[885].Y = -16_112.406

	pointsOfPolynomialPlotu[886].X = 8.86
	pointsOfPolynomialPlotu[886].Y = -16_176.401

	pointsOfPolynomialPlotu[887].X = 8.87
	pointsOfPolynomialPlotu[887].Y = -16_240.447

	pointsOfPolynomialPlotu[888].X = 8.88
	pointsOfPolynomialPlotu[888].Y = -16_304.541

	pointsOfPolynomialPlotu[889].X = 8.89
	pointsOfPolynomialPlotu[889].Y = -16_368.683

	pointsOfPolynomialPlotu[890].X = 8.90
	pointsOfPolynomialPlotu[890].Y = -16_432.873

	pointsOfPolynomialPlotu[891].X = 8.91
	pointsOfPolynomialPlotu[891].Y = -16_497.108

	pointsOfPolynomialPlotu[892].X = 8.92
	pointsOfPolynomialPlotu[892].Y = -16_561.389

	pointsOfPolynomialPlotu[893].X = 8.93
	pointsOfPolynomialPlotu[893].Y = -16_625.713

	pointsOfPolynomialPlotu[894].X = 8.94
	pointsOfPolynomialPlotu[894].Y = -16_690.079

	pointsOfPolynomialPlotu[895].X = 8.95
	pointsOfPolynomialPlotu[895].Y = -16_754.488

	pointsOfPolynomialPlotu[896].X = 8.96
	pointsOfPolynomialPlotu[896].Y = -16_818.937

	pointsOfPolynomialPlotu[897].X = 8.97
	pointsOfPolynomialPlotu[897].Y = -16_883.425

	pointsOfPolynomialPlotu[898].X = 8.98
	pointsOfPolynomialPlotu[898].Y = -16_947.952

	pointsOfPolynomialPlotu[899].X = 8.99
	pointsOfPolynomialPlotu[899].Y = -17_012.516

	pointsOfPolynomialPlotu[900].X = 9.0
	pointsOfPolynomialPlotu[900].Y = -17_077.116

	pointsOfPolynomialPlotu[901].X = 9.01
	pointsOfPolynomialPlotu[901].Y = -17_141.75

	pointsOfPolynomialPlotu[902].X = 9.02
	pointsOfPolynomialPlotu[902].Y = -17_206.419

	pointsOfPolynomialPlotu[903].X = 9.03
	pointsOfPolynomialPlotu[903].Y = -17_271.121

	pointsOfPolynomialPlotu[904].X = 9.04
	pointsOfPolynomialPlotu[904].Y = -17_335.854

	pointsOfPolynomialPlotu[905].X = 9.05
	pointsOfPolynomialPlotu[905].Y = -17_400.617

	pointsOfPolynomialPlotu[906].X = 9.06
	pointsOfPolynomialPlotu[906].Y = -17_465.41

	pointsOfPolynomialPlotu[907].X = 9.07
	pointsOfPolynomialPlotu[907].Y = -17_530.231

	pointsOfPolynomialPlotu[908].X = 9.08
	pointsOfPolynomialPlotu[908].Y = -17_595.079

	pointsOfPolynomialPlotu[909].X = 9.09
	pointsOfPolynomialPlotu[909].Y = -17_659.952

	pointsOfPolynomialPlotu[910].X = 9.10
	pointsOfPolynomialPlotu[910].Y = -17_724.85

	pointsOfPolynomialPlotu[911].X = 9.11
	pointsOfPolynomialPlotu[911].Y = -17_789.771

	pointsOfPolynomialPlotu[912].X = 9.12
	pointsOfPolynomialPlotu[912].Y = -17_854.715

	pointsOfPolynomialPlotu[913].X = 9.13
	pointsOfPolynomialPlotu[913].Y = -17_919.68

	pointsOfPolynomialPlotu[914].X = 9.14
	pointsOfPolynomialPlotu[914].Y = -17_984.664

	pointsOfPolynomialPlotu[915].X = 9.15
	pointsOfPolynomialPlotu[915].Y = -18_049.667

	pointsOfPolynomialPlotu[916].X = 9.16
	pointsOfPolynomialPlotu[916].Y = -18_114.687

	pointsOfPolynomialPlotu[917].X = 9.17
	pointsOfPolynomialPlotu[917].Y = -18_179.724

	pointsOfPolynomialPlotu[918].X = 9.18
	pointsOfPolynomialPlotu[918].Y = -18_244.775

	pointsOfPolynomialPlotu[919].X = 9.19
	pointsOfPolynomialPlotu[919].Y = -18_309.84

	pointsOfPolynomialPlotu[920].X = 9.20
	pointsOfPolynomialPlotu[920].Y = -18_374.918

	pointsOfPolynomialPlotu[921].X = 9.21
	pointsOfPolynomialPlotu[921].Y = -18_440.006

	pointsOfPolynomialPlotu[922].X = 9.22
	pointsOfPolynomialPlotu[922].Y = -18_505.105

	pointsOfPolynomialPlotu[923].X = 9.23
	pointsOfPolynomialPlotu[923].Y = -18_570.212

	pointsOfPolynomialPlotu[924].X = 9.24
	pointsOfPolynomialPlotu[924].Y = -18_635.327

	pointsOfPolynomialPlotu[925].X = 9.25
	pointsOfPolynomialPlotu[925].Y = -18_700.447

	pointsOfPolynomialPlotu[926].X = 9.26
	pointsOfPolynomialPlotu[926].Y = -18_765.573

	pointsOfPolynomialPlotu[927].X = 9.27
	pointsOfPolynomialPlotu[927].Y = -18_830.702

	pointsOfPolynomialPlotu[928].X = 9.28
	pointsOfPolynomialPlotu[928].Y = -18_895.833

	pointsOfPolynomialPlotu[929].X = 9.29
	pointsOfPolynomialPlotu[929].Y = -18_960.965

	pointsOfPolynomialPlotu[930].X = 9.30
	pointsOfPolynomialPlotu[930].Y = -19_026.097

	pointsOfPolynomialPlotu[931].X = 9.31
	pointsOfPolynomialPlotu[931].Y = -19_091.227

	pointsOfPolynomialPlotu[932].X = 9.32
	pointsOfPolynomialPlotu[932].Y = -19_156.354

	pointsOfPolynomialPlotu[933].X = 9.33
	pointsOfPolynomialPlotu[933].Y = -19_221.4777

	pointsOfPolynomialPlotu[934].X = 9.34
	pointsOfPolynomialPlotu[934].Y = -19_286.594

	pointsOfPolynomialPlotu[935].X = 9.35
	pointsOfPolynomialPlotu[935].Y = -19_351.704

	pointsOfPolynomialPlotu[936].X = 9.36
	pointsOfPolynomialPlotu[936].Y = -19_416.805

	pointsOfPolynomialPlotu[937].X = 9.37
	pointsOfPolynomialPlotu[937].Y = -19_481.897

	pointsOfPolynomialPlotu[938].X = 9.38
	pointsOfPolynomialPlotu[938].Y = -19_546.978

	pointsOfPolynomialPlotu[939].X = 9.39
	pointsOfPolynomialPlotu[939].Y = -19_612.047

	pointsOfPolynomialPlotu[940].X = 9.40
	pointsOfPolynomialPlotu[940].Y = -19_677.101

	pointsOfPolynomialPlotu[941].X = 9.41
	pointsOfPolynomialPlotu[941].Y = -19_742.141

	pointsOfPolynomialPlotu[942].X = 9.42
	pointsOfPolynomialPlotu[942].Y = -19_807.163

	pointsOfPolynomialPlotu[943].X = 9.43
	pointsOfPolynomialPlotu[943].Y = -19_872.168

	pointsOfPolynomialPlotu[944].X = 9.44
	pointsOfPolynomialPlotu[944].Y = -19_937.154

	pointsOfPolynomialPlotu[945].X = 9.45
	pointsOfPolynomialPlotu[945].Y = -20_002.118

	pointsOfPolynomialPlotu[946].X = 9.46
	pointsOfPolynomialPlotu[946].Y = -20_067.061

	pointsOfPolynomialPlotu[947].X = 9.47
	pointsOfPolynomialPlotu[947].Y = -20_131.979

	pointsOfPolynomialPlotu[948].X = 9.48
	pointsOfPolynomialPlotu[948].Y = -20_196.873

	pointsOfPolynomialPlotu[949].X = 9.49
	pointsOfPolynomialPlotu[949].Y = -20_261.74

	pointsOfPolynomialPlotu[950].X = 9.50
	pointsOfPolynomialPlotu[950].Y = -20_326.58

	pointsOfPolynomialPlotu[951].X = 9.51
	pointsOfPolynomialPlotu[951].Y = -20_391.389

	pointsOfPolynomialPlotu[952].X = 9.52
	pointsOfPolynomialPlotu[952].Y = -20_456.168

	pointsOfPolynomialPlotu[953].X = 9.53
	pointsOfPolynomialPlotu[953].Y = -20_520.915

	pointsOfPolynomialPlotu[954].X = 9.54
	pointsOfPolynomialPlotu[954].Y = -20_585.628

	pointsOfPolynomialPlotu[955].X = 9.55
	pointsOfPolynomialPlotu[955].Y = -20_650.306

	pointsOfPolynomialPlotu[956].X = 9.56
	pointsOfPolynomialPlotu[956].Y = -20_714.947

	pointsOfPolynomialPlotu[957].X = 9.57
	pointsOfPolynomialPlotu[957].Y = -20_779.55

	pointsOfPolynomialPlotu[958].X = 9.58
	pointsOfPolynomialPlotu[958].Y = -20_844.113

	pointsOfPolynomialPlotu[959].X = 9.59
	pointsOfPolynomialPlotu[959].Y = -20_908.634

	pointsOfPolynomialPlotu[960].X = 9.60
	pointsOfPolynomialPlotu[960].Y = -20_973.114

	pointsOfPolynomialPlotu[961].X = 9.61
	pointsOfPolynomialPlotu[961].Y = -21_037.548

	pointsOfPolynomialPlotu[962].X = 9.62
	pointsOfPolynomialPlotu[962].Y = -21_101.937

	pointsOfPolynomialPlotu[963].X = 9.63
	pointsOfPolynomialPlotu[963].Y = -21_166.279

	pointsOfPolynomialPlotu[964].X = 9.64
	pointsOfPolynomialPlotu[964].Y = -21_230.572

	pointsOfPolynomialPlotu[965].X = 9.65
	pointsOfPolynomialPlotu[965].Y = -21_294.815

	pointsOfPolynomialPlotu[966].X = 9.66
	pointsOfPolynomialPlotu[966].Y = -21_359.005

	pointsOfPolynomialPlotu[967].X = 9.67
	pointsOfPolynomialPlotu[967].Y = -21_423.142

	pointsOfPolynomialPlotu[968].X = 9.68
	pointsOfPolynomialPlotu[968].Y = -21_487.225

	pointsOfPolynomialPlotu[969].X = 9.69
	pointsOfPolynomialPlotu[969].Y = -21_551.25

	pointsOfPolynomialPlotu[970].X = 9.70
	pointsOfPolynomialPlotu[970].Y = -21_615.217

	pointsOfPolynomialPlotu[971].X = 9.71
	pointsOfPolynomialPlotu[971].Y = -21_679.125

	pointsOfPolynomialPlotu[972].X = 9.72
	pointsOfPolynomialPlotu[972].Y = -21_742.971

	pointsOfPolynomialPlotu[973].X = 9.73
	pointsOfPolynomialPlotu[973].Y = -21_806.754

	pointsOfPolynomialPlotu[974].X = 9.74
	pointsOfPolynomialPlotu[974].Y = -21_870.472

	pointsOfPolynomialPlotu[975].X = 9.75
	pointsOfPolynomialPlotu[975].Y = -21_934.124

	pointsOfPolynomialPlotu[976].X = 9.76
	pointsOfPolynomialPlotu[976].Y = -21_997.709

	pointsOfPolynomialPlotu[977].X = 9.77
	pointsOfPolynomialPlotu[977].Y = -22_061.224

	pointsOfPolynomialPlotu[978].X = 9.78
	pointsOfPolynomialPlotu[978].Y = -22_124.668

	pointsOfPolynomialPlotu[979].X = 9.79
	pointsOfPolynomialPlotu[979].Y = -22_188.039

	pointsOfPolynomialPlotu[980].X = 9.80
	pointsOfPolynomialPlotu[980].Y = -22_251.335

	pointsOfPolynomialPlotu[981].X = 9.81
	pointsOfPolynomialPlotu[981].Y = -22_314.556

	pointsOfPolynomialPlotu[982].X = 9.82
	pointsOfPolynomialPlotu[982].Y = -22_377.699

	pointsOfPolynomialPlotu[983].X = 9.83
	pointsOfPolynomialPlotu[983].Y = -22_440.763

	pointsOfPolynomialPlotu[984].X = 9.84
	pointsOfPolynomialPlotu[984].Y = -22_503.745

	pointsOfPolynomialPlotu[985].X = 9.85
	pointsOfPolynomialPlotu[985].Y = -22_566.645

	pointsOfPolynomialPlotu[986].X = 9.86
	pointsOfPolynomialPlotu[986].Y = -22_629.46

	pointsOfPolynomialPlotu[987].X = 9.87
	pointsOfPolynomialPlotu[987].Y = -22_692.189

	pointsOfPolynomialPlotu[988].X = 9.88
	pointsOfPolynomialPlotu[988].Y = -22_754.831

	pointsOfPolynomialPlotu[989].X = 9.89
	pointsOfPolynomialPlotu[989].Y = -22_817.382

	pointsOfPolynomialPlotu[990].X = 9.90
	pointsOfPolynomialPlotu[990].Y = -22_879.843

	pointsOfPolynomialPlotu[991].X = 9.91
	pointsOfPolynomialPlotu[991].Y = -22_942.21

	pointsOfPolynomialPlotu[992].X = 9.92
	pointsOfPolynomialPlotu[992].Y = -23_004.483

	pointsOfPolynomialPlotu[993].X = 9.93
	pointsOfPolynomialPlotu[993].Y = -23_066.659

	pointsOfPolynomialPlotu[994].X = 9.94
	pointsOfPolynomialPlotu[994].Y = -23_128.737

	pointsOfPolynomialPlotu[995].X = 9.95
	pointsOfPolynomialPlotu[995].Y = -23_190.715

	pointsOfPolynomialPlotu[996].X = 9.96
	pointsOfPolynomialPlotu[996].Y = -23_252.591

	pointsOfPolynomialPlotu[997].X = 9.97
	pointsOfPolynomialPlotu[997].Y = -23_314.364

	pointsOfPolynomialPlotu[998].X = 9.98
	pointsOfPolynomialPlotu[998].Y = -23_376.031

	pointsOfPolynomialPlotu[999].X = 9.99
	pointsOfPolynomialPlotu[999].Y = -23_437.591

	pointsOfPolynomialPlotu[1_000].X = 10.0
	pointsOfPolynomialPlotu[1_000].Y = -23_499.043










	polynomialPlot := plot.New()

	polynomialPlot.Title.Text = "Plot of polynomial f(x) = x^5 - 19.222x^4 + 69.903x^3 - 10.34x^2 - 15.722x + 9.177"

	polynomialPlot.X.Label.Text = "x"
	polynomialPlot.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfPolynomialPlotu)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	polynomialPlot.Add(plotLine)
	polynomialPlot.Legend.Add("f(x)", plotLine)

	if err := polynomialPlot.Save(10*vg.Inch, 10*vg.Inch,
		"Polynomial-function-plot-04.png"); err != nil {

		panic(err)
	}
}
