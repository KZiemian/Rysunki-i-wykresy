package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres wielomianu f(x) = x^5 - 19.222 x^4 + 69.903 x^3 -
	// 10.34 x^2 - 15.722 x + 9.177

	punktyWykresuWielomianu := make(plotter.XYs, 1_001)

	punktyWykresuWielomianu[0].X = 0.0
	punktyWykresuWielomianu[0].Y = 9.177

	punktyWykresuWielomianu[1].X = 0.01
	punktyWykresuWielomianu[1].Y = 9.018

	punktyWykresuWielomianu[2].X = 0.02
	punktyWykresuWielomianu[2].Y = 8.859

	punktyWykresuWielomianu[3].X = 0.03
	punktyWykresuWielomianu[3].Y = 8.697

	punktyWykresuWielomianu[4].X = 0.04
	punktyWykresuWielomianu[4].Y = 8.536

	punktyWykresuWielomianu[5].X = 0.05
	punktyWykresuWielomianu[5].Y = 8.373

	punktyWykresuWielomianu[6].X = 0.06
	punktyWykresuWielomianu[6].Y = 8.211

	punktyWykresuWielomianu[7].X = 0.07
	punktyWykresuWielomianu[7].Y = 8.049

	punktyWykresuWielomianu[8].X = 0.08
	punktyWykresuWielomianu[8].Y = 7.888

	punktyWykresuWielomianu[9].X = 0.09
	punktyWykresuWielomianu[9].Y = 7.728

	punktyWykresuWielomianu[10].X = 0.1
	punktyWykresuWielomianu[10].Y = 7.569

	punktyWykresuWielomianu[11].X = 0.11
	punktyWykresuWielomianu[11].Y = 7.412

	punktyWykresuWielomianu[12].X = 0.12
	punktyWykresuWielomianu[12].Y = 7.258

	punktyWykresuWielomianu[13].X = 0.13
	punktyWykresuWielomianu[13].Y = 7.106

	punktyWykresuWielomianu[14].X = 0.14
	punktyWykresuWielomianu[14].Y = 6.957

	punktyWykresuWielomianu[15].X = 0.15
	punktyWykresuWielomianu[15].Y = 6.812

	punktyWykresuWielomianu[16].X = 0.16
	punktyWykresuWielomianu[16].Y = 6.67

	punktyWykresuWielomianu[17].X = 0.17
	punktyWykresuWielomianu[17].Y = 6.532

	punktyWykresuWielomianu[18].X = 0.18
	punktyWykresuWielomianu[18].Y = 6.399

	punktyWykresuWielomianu[19].X = 0.19
	punktyWykresuWielomianu[19].Y = 6.271

	punktyWykresuWielomianu[20].X = 0.2
	punktyWykresuWielomianu[20].Y = 6.147

	punktyWykresuWielomianu[21].X = 0.21
	punktyWykresuWielomianu[21].Y = 6.029

	punktyWykresuWielomianu[22].X = 0.22
	punktyWykresuWielomianu[22].Y = 5.917

	punktyWykresuWielomianu[23].X = 0.23
	punktyWykresuWielomianu[23].Y = 5.811

	punktyWykresuWielomianu[24].X = 0.24
	punktyWykresuWielomianu[24].Y = 5.711

	punktyWykresuWielomianu[25].X = 0.25
	punktyWykresuWielomianu[25].Y = 5.618

	punktyWykresuWielomianu[26].X = 0.26
	punktyWykresuWielomianu[26].Y = 5.532

	punktyWykresuWielomianu[27].X = 0.27
	punktyWykresuWielomianu[27].Y = 5.453

	punktyWykresuWielomianu[28].X = 0.28
	punktyWykresuWielomianu[28].Y = 5.382

	punktyWykresuWielomianu[29].X = 0.29
	punktyWykresuWielomianu[29].Y = 5.318

	punktyWykresuWielomianu[30].X = 0.3
	punktyWykresuWielomianu[30].Y = 5.263

	punktyWykresuWielomianu[31].X = 0.31
	punktyWykresuWielomianu[31].Y = 5.217

	punktyWykresuWielomianu[32].X = 0.32
	punktyWykresuWielomianu[32].Y = 5.179

	punktyWykresuWielomianu[33].X = 0.33
	punktyWykresuWielomianu[33].Y = 5.15

	punktyWykresuWielomianu[34].X = 0.34
	punktyWykresuWielomianu[34].Y = 5.131

	punktyWykresuWielomianu[35].X = 0.35
	punktyWykresuWielomianu[35].Y = 5.121

	punktyWykresuWielomianu[36].X = 0.36
	punktyWykresuWielomianu[36].Y = 5.121

	punktyWykresuWielomianu[37].X = 0.37
	punktyWykresuWielomianu[37].Y = 5.131

	punktyWykresuWielomianu[38].X = 0.38
	punktyWykresuWielomianu[38].Y = 5.152

	punktyWykresuWielomianu[39].X = 0.39
	punktyWykresuWielomianu[39].Y = 5.183

	punktyWykresuWielomianu[40].X = 0.4
	punktyWykresuWielomianu[40].Y = 5.225

	punktyWykresuWielomianu[41].X = 0.41
	punktyWykresuWielomianu[41].Y = 5.278

	punktyWykresuWielomianu[42].X = 0.42
	punktyWykresuWielomianu[42].Y = 5.343

	punktyWykresuWielomianu[43].X = 0.43
	punktyWykresuWielomianu[43].Y = 5.419

	punktyWykresuWielomianu[44].X = 0.44
	punktyWykresuWielomianu[44].Y = 5.507

	punktyWykresuWielomianu[45].X = 0.45
	punktyWykresuWielomianu[45].Y = 5.607

	punktyWykresuWielomianu[46].X = 0.46
	punktyWykresuWielomianu[46].Y = 5.720

	punktyWykresuWielomianu[47].X = 0.47
	punktyWykresuWielomianu[47].Y = 5.845

	punktyWykresuWielomianu[48].X = 0.48
	punktyWykresuWielomianu[48].Y = 5.983

	punktyWykresuWielomianu[49].X = 0.49
	punktyWykresuWielomianu[49].Y = 6.134

	punktyWykresuWielomianu[50].X = 0.5
	punktyWykresuWielomianu[50].Y = 6.298

	punktyWykresuWielomianu[51].X = 0.51
	punktyWykresuWielomianu[51].Y = 6.475

	punktyWykresuWielomianu[52].X = 0.52
	punktyWykresuWielomianu[52].Y = 6.666

	punktyWykresuWielomianu[53].X = 0.53
	punktyWykresuWielomianu[53].Y = 6.871

	punktyWykresuWielomianu[54].X = 0.54
	punktyWykresuWielomianu[54].Y = 7.089

	punktyWykresuWielomianu[55].X = 0.55
	punktyWykresuWielomianu[55].Y = 7.322

	punktyWykresuWielomianu[56].X = 0.56
	punktyWykresuWielomianu[56].Y = 7.569

	punktyWykresuWielomianu[57].X = 0.57
	punktyWykresuWielomianu[57].Y = 7.831

	punktyWykresuWielomianu[58].X = 0.58
	punktyWykresuWielomianu[58].Y = 8.107

	punktyWykresuWielomianu[59].X = 0.59
	punktyWykresuWielomianu[59].Y = 8.399

	punktyWykresuWielomianu[60].X = 0.6
	punktyWykresuWielomianu[60].Y = 8.705

	punktyWykresuWielomianu[61].X = 0.61
	punktyWykresuWielomianu[61].Y = 9.027

	punktyWykresuWielomianu[62].X = 0.62
	punktyWykresuWielomianu[62].Y = 9.364

	punktyWykresuWielomianu[63].X = 0.63
	punktyWykresuWielomianu[63].Y = 9.716

	punktyWykresuWielomianu[64].X = 0.64
	punktyWykresuWielomianu[64].Y = 10.084

	punktyWykresuWielomianu[65].X = 0.65
	punktyWykresuWielomianu[65].Y = 10.469

	punktyWykresuWielomianu[66].X = 0.66
	punktyWykresuWielomianu[66].Y = 10.869

	punktyWykresuWielomianu[67].X = 0.67
	punktyWykresuWielomianu[67].Y = 11.285

	punktyWykresuWielomianu[68].X = 0.68
	punktyWykresuWielomianu[68].Y = 11.717

	punktyWykresuWielomianu[69].X = 0.69
	punktyWykresuWielomianu[69].Y = 12.166

	punktyWykresuWielomianu[70].X = 0.7
	punktyWykresuWielomianu[70].Y = 12.632

	punktyWykresuWielomianu[71].X = 0.71
	punktyWykresuWielomianu[71].Y = 13.114

	punktyWykresuWielomianu[72].X = 0.72
	punktyWykresuWielomianu[72].Y = 13.612

	punktyWykresuWielomianu[73].X = 0.73
	punktyWykresuWielomianu[73].Y = 14.128

	punktyWykresuWielomianu[74].X = 0.74
	punktyWykresuWielomianu[74].Y = 14.661

	punktyWykresuWielomianu[75].X = 0.75
	punktyWykresuWielomianu[75].Y = 15.211

	punktyWykresuWielomianu[76].X = 0.76
	punktyWykresuWielomianu[76].Y = 15.778

	punktyWykresuWielomianu[77].X = 0.77
	punktyWykresuWielomianu[77].Y = 16.363

	punktyWykresuWielomianu[78].X = 0.78
	punktyWykresuWielomianu[78].Y = 16.965

	punktyWykresuWielomianu[79].X = 0.79
	punktyWykresuWielomianu[79].Y = 17.584

	punktyWykresuWielomianu[80].X = 0.8
	punktyWykresuWielomianu[80].Y = 18.222

	punktyWykresuWielomianu[81].X = 0.81
	punktyWykresuWielomianu[81].Y = 18.876

	punktyWykresuWielomianu[82].X = 0.82
	punktyWykresuWielomianu[82].Y = 19.549

	punktyWykresuWielomianu[83].X = 0.83
	punktyWykresuWielomianu[83].Y = 20.24

	punktyWykresuWielomianu[84].X = 0.84
	punktyWykresuWielomianu[84].Y = 20.949

	punktyWykresuWielomianu[85].X = 0.85
	punktyWykresuWielomianu[85].Y = 21.675

	punktyWykresuWielomianu[86].X = 0.86
	punktyWykresuWielomianu[86].Y = 22.42

	punktyWykresuWielomianu[87].X = 0.87
	punktyWykresuWielomianu[87].Y = 23.183

	punktyWykresuWielomianu[88].X = 0.88
	punktyWykresuWielomianu[88].Y = 23.965

	punktyWykresuWielomianu[89].X = 0.89
	punktyWykresuWielomianu[89].Y = 24.764

	punktyWykresuWielomianu[90].X = 0.9
	punktyWykresuWielomianu[90].Y = 25.582

	punktyWykresuWielomianu[91].X = 0.91
	punktyWykresuWielomianu[91].Y = 26.419

	punktyWykresuWielomianu[92].X = 0.92
	punktyWykresuWielomianu[92].Y = 27.274

	punktyWykresuWielomianu[93].X = 0.93
	punktyWykresuWielomianu[93].Y = 28.147

	punktyWykresuWielomianu[94].X = 0.94
	punktyWykresuWielomianu[94].Y = 29.04

	punktyWykresuWielomianu[95].X = 0.95
	punktyWykresuWielomianu[95].Y = 29.95

	punktyWykresuWielomianu[96].X = 0.96
	punktyWykresuWielomianu[96].Y = 30.88

	punktyWykresuWielomianu[97].X = 0.97
	punktyWykresuWielomianu[97].Y = 31.828

	punktyWykresuWielomianu[98].X = 0.98
	punktyWykresuWielomianu[98].Y = 32.795

	punktyWykresuWielomianu[99].X = 0.99
	punktyWykresuWielomianu[99].Y = 33.78

	punktyWykresuWielomianu[100].X = 1.0
	punktyWykresuWielomianu[100].Y = 34.785

	punktyWykresuWielomianu[101].X = 1.01
	punktyWykresuWielomianu[101].Y = 35.808

	punktyWykresuWielomianu[102].X = 1.02
	punktyWykresuWielomianu[102].Y = 36.85

	punktyWykresuWielomianu[103].X = 1.03
	punktyWykresuWielomianu[103].Y = 37.91

	punktyWykresuWielomianu[104].X = 1.04
	punktyWykresuWielomianu[104].Y = 38.99

	punktyWykresuWielomianu[105].X = 1.05
	punktyWykresuWielomianu[105].Y = 40.089

	punktyWykresuWielomianu[106].X = 1.06
	punktyWykresuWielomianu[106].Y = 41.206

	punktyWykresuWielomianu[107].X = 1.07
	punktyWykresuWielomianu[107].Y = 42.342

	punktyWykresuWielomianu[108].X = 1.08
	punktyWykresuWielomianu[108].Y = 43.497

	punktyWykresuWielomianu[109].X = 1.09
	punktyWykresuWielomianu[109].Y = 44.671

	punktyWykresuWielomianu[110].X = 1.1
	punktyWykresuWielomianu[110].Y = 45.863

	punktyWykresuWielomianu[111].X = 1.11
	punktyWykresuWielomianu[111].Y = 47.075

	punktyWykresuWielomianu[112].X = 1.12
	punktyWykresuWielomianu[112].Y = 48.305

	punktyWykresuWielomianu[113].X = 1.13
	punktyWykresuWielomianu[113].Y = 49.554

	punktyWykresuWielomianu[114].X = 1.14
	punktyWykresuWielomianu[114].Y = 50.822

	punktyWykresuWielomianu[115].X = 1.15
	punktyWykresuWielomianu[115].Y = 52.108

	punktyWykresuWielomianu[116].X = 1.16
	punktyWykresuWielomianu[116].Y = 53.413

	punktyWykresuWielomianu[117].X = 1.17
	punktyWykresuWielomianu[117].Y = 54.737

	punktyWykresuWielomianu[118].X = 1.18
	punktyWykresuWielomianu[118].Y = 56.079

	punktyWykresuWielomianu[119].X = 1.19
	punktyWykresuWielomianu[119].Y = 57.44

	punktyWykresuWielomianu[120].X = 1.2
	punktyWykresuWielomianu[120].Y = 58.843

	punktyWykresuWielomianu[121].X = 1.21
	punktyWykresuWielomianu[121].Y = 60.241

	punktyWykresuWielomianu[122].X = 1.22
	punktyWykresuWielomianu[122].Y = 61.658

	punktyWykresuWielomianu[123].X = 1.23
	punktyWykresuWielomianu[123].Y = 63.094

	punktyWykresuWielomianu[124].X = 1.24
	punktyWykresuWielomianu[124].Y = 64.548

	punktyWykresuWielomianu[125].X = 1.25
	punktyWykresuWielomianu[125].Y = 66.02

	punktyWykresuWielomianu[126].X = 1.26
	punktyWykresuWielomianu[126].Y = 67.511

	punktyWykresuWielomianu[127].X = 1.27
	punktyWykresuWielomianu[127].Y = 69.019

	punktyWykresuWielomianu[128].X = 1.28
	punktyWykresuWielomianu[128].Y = 70.546

	punktyWykresuWielomianu[129].X = 1.29
	punktyWykresuWielomianu[129].Y = 72.091

	punktyWykresuWielomianu[130].X = 1.3
	punktyWykresuWielomianu[130].Y = 73.653

	punktyWykresuWielomianu[131].X = 1.31
	punktyWykresuWielomianu[131].Y = 75.234

	punktyWykresuWielomianu[132].X = 1.32
	punktyWykresuWielomianu[132].Y = 76.832

	punktyWykresuWielomianu[133].X = 1.33
	punktyWykresuWielomianu[133].Y = 78.448

	punktyWykresuWielomianu[134].X = 1.34
	punktyWykresuWielomianu[134].Y = 80.082

	punktyWykresuWielomianu[135].X = 1.35
	punktyWykresuWielomianu[135].Y = 81.733

	punktyWykresuWielomianu[136].X = 1.36
	punktyWykresuWielomianu[136].Y = 83.401

	punktyWykresuWielomianu[137].X = 1.37
	punktyWykresuWielomianu[137].Y = 85.087

	punktyWykresuWielomianu[138].X = 1.38
	punktyWykresuWielomianu[138].Y = 86.791

	punktyWykresuWielomianu[139].X = 1.39
	punktyWykresuWielomianu[139].Y = 88.511

	punktyWykresuWielomianu[140].X = 1.4
	punktyWykresuWielomianu[140].Y = 90.248

	punktyWykresuWielomianu[141].X = 1.41
	punktyWykresuWielomianu[141].Y = 92.002

	punktyWykresuWielomianu[142].X = 1.42
	punktyWykresuWielomianu[142].Y = 93.774

	punktyWykresuWielomianu[143].X = 1.43
	punktyWykresuWielomianu[143].Y = 95.561

	punktyWykresuWielomianu[144].X = 1.44
	punktyWykresuWielomianu[144].Y = 97.366

	punktyWykresuWielomianu[145].X = 1.45
	punktyWykresuWielomianu[145].Y = 99.187

	punktyWykresuWielomianu[146].X = 1.46
	punktyWykresuWielomianu[146].Y = 101.024

	punktyWykresuWielomianu[147].X = 1.47
	punktyWykresuWielomianu[147].Y = 102.877

	punktyWykresuWielomianu[148].X = 1.48
	punktyWykresuWielomianu[148].Y = 104.747

	punktyWykresuWielomianu[149].X = 1.49
	punktyWykresuWielomianu[149].Y = 106.632

	punktyWykresuWielomianu[150].X = 1.5
	punktyWykresuWielomianu[150].Y = 108.534

	punktyWykresuWielomianu[151].X = 1.51
	punktyWykresuWielomianu[151].Y = 110.451

	punktyWykresuWielomianu[152].X = 1.52
	punktyWykresuWielomianu[152].Y = 112.383

	punktyWykresuWielomianu[153].X = 1.53
	punktyWykresuWielomianu[153].Y = 114.331

	punktyWykresuWielomianu[154].X = 1.54
	punktyWykresuWielomianu[154].Y = 116.294

	punktyWykresuWielomianu[155].X = 1.55
	punktyWykresuWielomianu[155].Y = 118.273

	punktyWykresuWielomianu[156].X = 1.56
	punktyWykresuWielomianu[156].Y = 120.266

	punktyWykresuWielomianu[157].X = 1.57
	punktyWykresuWielomianu[157].Y = 122.274

	punktyWykresuWielomianu[158].X = 1.58
	punktyWykresuWielomianu[158].Y = 124.297

	punktyWykresuWielomianu[159].X = 1.59
	punktyWykresuWielomianu[159].Y = 126.334

	punktyWykresuWielomianu[160].X = 1.6
	punktyWykresuWielomianu[160].Y = 128.386

	punktyWykresuWielomianu[161].X = 1.61
	punktyWykresuWielomianu[161].Y = 130.452

	punktyWykresuWielomianu[162].X = 1.62
	punktyWykresuWielomianu[162].Y = 132.532

	punktyWykresuWielomianu[163].X = 1.63
	punktyWykresuWielomianu[163].Y = 134.626

	punktyWykresuWielomianu[164].X = 1.64
	punktyWykresuWielomianu[164].Y = 136.733

	punktyWykresuWielomianu[165].X = 1.65
	punktyWykresuWielomianu[165].Y = 138.854

	punktyWykresuWielomianu[166].X = 1.66
	punktyWykresuWielomianu[166].Y = 140.988

	punktyWykresuWielomianu[167].X = 1.67
	punktyWykresuWielomianu[167].Y = 143.135

	punktyWykresuWielomianu[168].X = 1.68
	punktyWykresuWielomianu[168].Y = 145.296

	punktyWykresuWielomianu[169].X = 1.69
	punktyWykresuWielomianu[169].Y = 147.469

	punktyWykresuWielomianu[170].X = 1.7
	punktyWykresuWielomianu[170].Y = 149.654

	punktyWykresuWielomianu[171].X = 1.71
	punktyWykresuWielomianu[171].Y = 151.853

	punktyWykresuWielomianu[172].X = 1.72
	punktyWykresuWielomianu[172].Y = 154.063

	punktyWykresuWielomianu[173].X = 1.73
	punktyWykresuWielomianu[173].Y = 156.285

	punktyWykresuWielomianu[174].X = 1.74
	punktyWykresuWielomianu[174].Y = 158.519

	punktyWykresuWielomianu[175].X = 1.75
	punktyWykresuWielomianu[175].Y = 160.765

	punktyWykresuWielomianu[176].X = 1.76
	punktyWykresuWielomianu[176].Y = 163.022

	punktyWykresuWielomianu[177].X = 1.77
	punktyWykresuWielomianu[177].Y = 165.29

	punktyWykresuWielomianu[178].X = 1.78
	punktyWykresuWielomianu[178].Y = 167.57

	punktyWykresuWielomianu[179].X = 1.79
	punktyWykresuWielomianu[179].Y = 169.86

	punktyWykresuWielomianu[180].X = 1.8
	punktyWykresuWielomianu[180].Y = 172.16

	punktyWykresuWielomianu[181].X = 1.81
	punktyWykresuWielomianu[181].Y = 174.471

	punktyWykresuWielomianu[182].X = 1.82
	punktyWykresuWielomianu[182].Y = 176.793

	punktyWykresuWielomianu[183].X = 1.83
	punktyWykresuWielomianu[183].Y = 179.124

	punktyWykresuWielomianu[184].X = 1.84
	punktyWykresuWielomianu[184].Y = 181.465

	punktyWykresuWielomianu[185].X = 1.85
	punktyWykresuWielomianu[185].Y = 183.815

	punktyWykresuWielomianu[186].X = 1.86
	punktyWykresuWielomianu[186].Y = 186.174

	punktyWykresuWielomianu[187].X = 1.87
	punktyWykresuWielomianu[187].Y = 188.543

	punktyWykresuWielomianu[188].X = 1.88
	punktyWykresuWielomianu[188].Y = 190.92

	punktyWykresuWielomianu[189].X = 1.89
	punktyWykresuWielomianu[189].Y = 193.306

	punktyWykresuWielomianu[190].X = 1.9
	punktyWykresuWielomianu[190].Y = 195.7

	punktyWykresuWielomianu[191].X = 1.91
	punktyWykresuWielomianu[191].Y = 198.102

	punktyWykresuWielomianu[192].X = 1.92
	punktyWykresuWielomianu[192].Y = 200.512

	punktyWykresuWielomianu[193].X = 1.93
	punktyWykresuWielomianu[193].Y = 202.93

	punktyWykresuWielomianu[194].X = 1.94
	punktyWykresuWielomianu[194].Y = 205.355

	punktyWykresuWielomianu[195].X = 1.95
	punktyWykresuWielomianu[195].Y = 207.787

	punktyWykresuWielomianu[196].X = 1.96
	punktyWykresuWielomianu[196].Y = 210.226

	punktyWykresuWielomianu[197].X = 1.97
	punktyWykresuWielomianu[197].Y = 212.671

	punktyWykresuWielomianu[198].X = 1.98
	punktyWykresuWielomianu[198].Y = 215.123

	punktyWykresuWielomianu[199].X = 1.99
	punktyWykresuWielomianu[199].Y = 217.581

	punktyWykresuWielomianu[200].X = 2.0
	punktyWykresuWielomianu[200].Y = 220.045

	punktyWykresuWielomianu[201].X = 2.01
	punktyWykresuWielomianu[201].Y = 222.514

	punktyWykresuWielomianu[202].X = 2.02
	punktyWykresuWielomianu[202].Y = 224.988

	punktyWykresuWielomianu[203].X = 2.03
	punktyWykresuWielomianu[203].Y = 227.468

	punktyWykresuWielomianu[204].X = 2.04
	punktyWykresuWielomianu[204].Y = 229.952

	punktyWykresuWielomianu[205].X = 2.05
	punktyWykresuWielomianu[205].Y = 232.441

	punktyWykresuWielomianu[206].X = 2.06
	punktyWykresuWielomianu[206].Y = 234.934

	punktyWykresuWielomianu[207].X = 2.07
	punktyWykresuWielomianu[207].Y = 237.431

	punktyWykresuWielomianu[208].X = 2.08
	punktyWykresuWielomianu[208].Y = 239.931

	punktyWykresuWielomianu[209].X = 2.09
	punktyWykresuWielomianu[209].Y = 242.435

	punktyWykresuWielomianu[210].X = 2.10
	punktyWykresuWielomianu[210].Y = 244.942

	punktyWykresuWielomianu[211].X = 2.11
	punktyWykresuWielomianu[211].Y = 247.452

	punktyWykresuWielomianu[212].X = 2.12
	punktyWykresuWielomianu[212].Y = 249.964

	punktyWykresuWielomianu[213].X = 2.13
	punktyWykresuWielomianu[213].Y = 252.479

	punktyWykresuWielomianu[214].X = 2.14
	punktyWykresuWielomianu[214].Y = 254.996

	punktyWykresuWielomianu[215].X = 2.15
	punktyWykresuWielomianu[215].Y = 257.514

	punktyWykresuWielomianu[216].X = 2.16
	punktyWykresuWielomianu[216].Y = 260.033

	punktyWykresuWielomianu[217].X = 2.17
	punktyWykresuWielomianu[217].Y = 262.554

	punktyWykresuWielomianu[218].X = 2.18
	punktyWykresuWielomianu[218].Y = 265.075

	punktyWykresuWielomianu[219].X = 2.19
	punktyWykresuWielomianu[219].Y = 267.597

	punktyWykresuWielomianu[220].X = 2.2
	punktyWykresuWielomianu[220].Y = 270.119

	punktyWykresuWielomianu[221].X = 2.21
	punktyWykresuWielomianu[221].Y = 272.641

	punktyWykresuWielomianu[222].X = 2.22
	punktyWykresuWielomianu[222].Y = 275.162

	punktyWykresuWielomianu[223].X = 2.23
	punktyWykresuWielomianu[223].Y = 277.683

	punktyWykresuWielomianu[224].X = 2.24
	punktyWykresuWielomianu[224].Y = 280.203

	punktyWykresuWielomianu[225].X = 2.25
	punktyWykresuWielomianu[225].Y = 282.721

	punktyWykresuWielomianu[226].X = 2.26
	punktyWykresuWielomianu[226].Y = 285.237

	punktyWykresuWielomianu[227].X = 2.27
	punktyWykresuWielomianu[227].Y = 287.752

	punktyWykresuWielomianu[228].X = 2.28
	punktyWykresuWielomianu[228].Y = 290.264

	punktyWykresuWielomianu[229].X = 2.29
	punktyWykresuWielomianu[229].Y = 292.774

	punktyWykresuWielomianu[230].X = 2.3
	punktyWykresuWielomianu[230].Y = 295.28

	punktyWykresuWielomianu[231].X = 2.31
	punktyWykresuWielomianu[231].Y = 297.784

	punktyWykresuWielomianu[232].X = 2.32
	punktyWykresuWielomianu[232].Y = 300.283

	punktyWykresuWielomianu[233].X = 2.33
	punktyWykresuWielomianu[233].Y = 302.779

	punktyWykresuWielomianu[234].X = 2.34
	punktyWykresuWielomianu[234].Y = 305.27

	punktyWykresuWielomianu[235].X = 2.35
	punktyWykresuWielomianu[235].Y = 307.757

	punktyWykresuWielomianu[236].X = 2.36
	punktyWykresuWielomianu[236].Y = 310.239

	punktyWykresuWielomianu[237].X = 2.37
	punktyWykresuWielomianu[237].Y = 312.716

	punktyWykresuWielomianu[238].X = 2.38
	punktyWykresuWielomianu[238].Y = 315.187

	punktyWykresuWielomianu[239].X = 2.39
	punktyWykresuWielomianu[239].Y = 317.652

	punktyWykresuWielomianu[240].X = 2.4
	punktyWykresuWielomianu[240].Y = 320.111

	punktyWykresuWielomianu[241].X = 2.41
	punktyWykresuWielomianu[241].Y = 322.563

	punktyWykresuWielomianu[242].X = 2.42
	punktyWykresuWielomianu[242].Y = 325.008

	punktyWykresuWielomianu[243].X = 2.43
	punktyWykresuWielomianu[243].Y = 327.446

	punktyWykresuWielomianu[244].X = 2.44
	punktyWykresuWielomianu[244].Y = 329.877

	punktyWykresuWielomianu[245].X = 2.45
	punktyWykresuWielomianu[245].Y = 332.299

	punktyWykresuWielomianu[246].X = 2.46
	punktyWykresuWielomianu[246].Y = 334.713

	punktyWykresuWielomianu[247].X = 2.47
	punktyWykresuWielomianu[247].Y = 337.118

	punktyWykresuWielomianu[248].X = 2.48
	punktyWykresuWielomianu[248].Y = 339.514

	punktyWykresuWielomianu[249].X = 2.49
	punktyWykresuWielomianu[249].Y = 341.901

	punktyWykresuWielomianu[250].X = 2.5
	punktyWykresuWielomianu[250].Y = 344.278

	punktyWykresuWielomianu[251].X = 2.51
	punktyWykresuWielomianu[251].Y = 346.645

	punktyWykresuWielomianu[252].X = 2.52
	punktyWykresuWielomianu[252].Y = 349.001

	punktyWykresuWielomianu[253].X = 2.53
	punktyWykresuWielomianu[253].Y = 351.347

	punktyWykresuWielomianu[254].X = 2.54
	punktyWykresuWielomianu[254].Y = 353.681

	punktyWykresuWielomianu[255].X = 2.55
	punktyWykresuWielomianu[255].Y = 356.004

	punktyWykresuWielomianu[256].X = 2.56
	punktyWykresuWielomianu[256].Y = 358.314

	punktyWykresuWielomianu[257].X = 2.57
	punktyWykresuWielomianu[257].Y = 360.613

	punktyWykresuWielomianu[258].X = 2.58
	punktyWykresuWielomianu[258].Y = 362.899

	punktyWykresuWielomianu[259].X = 2.59
	punktyWykresuWielomianu[259].Y = 365.171

	punktyWykresuWielomianu[260].X = 2.6
	punktyWykresuWielomianu[260].Y = 367.431

	punktyWykresuWielomianu[261].X = 2.61
	punktyWykresuWielomianu[261].Y = 369.676

	punktyWykresuWielomianu[262].X = 2.62
	punktyWykresuWielomianu[262].Y = 371.907

	punktyWykresuWielomianu[263].X = 2.63
	punktyWykresuWielomianu[263].Y = 374.124

	punktyWykresuWielomianu[264].X = 2.64
	punktyWykresuWielomianu[264].Y = 376.326

	punktyWykresuWielomianu[265].X = 2.65
	punktyWykresuWielomianu[265].Y = 378.513

	punktyWykresuWielomianu[266].X = 2.66
	punktyWykresuWielomianu[266].Y = 380.684

	punktyWykresuWielomianu[267].X = 2.67
	punktyWykresuWielomianu[267].Y = 382.838

	punktyWykresuWielomianu[268].X = 2.68
	punktyWykresuWielomianu[268].Y = 384.977

	punktyWykresuWielomianu[269].X = 2.69
	punktyWykresuWielomianu[269].Y = 387.098

	punktyWykresuWielomianu[270].X = 2.7
	punktyWykresuWielomianu[270].Y = 389.202

	punktyWykresuWielomianu[271].X = 2.71
	punktyWykresuWielomianu[271].Y = 391.289

	punktyWykresuWielomianu[272].X = 2.72
	punktyWykresuWielomianu[272].Y = 393.358

	punktyWykresuWielomianu[273].X = 2.73
	punktyWykresuWielomianu[273].Y = 395.408

	punktyWykresuWielomianu[274].X = 2.74
	punktyWykresuWielomianu[274].Y = 397.44

	punktyWykresuWielomianu[275].X = 2.75
	punktyWykresuWielomianu[275].Y = 399.452

	punktyWykresuWielomianu[276].X = 2.76
	punktyWykresuWielomianu[276].Y = 401.445

	punktyWykresuWielomianu[277].X = 2.77
	punktyWykresuWielomianu[277].Y = 403.417

	punktyWykresuWielomianu[278].X = 2.78
	punktyWykresuWielomianu[278].Y = 405.37

	punktyWykresuWielomianu[279].X = 2.79
	punktyWykresuWielomianu[279].Y = 407.301

	punktyWykresuWielomianu[280].X = 2.8
	punktyWykresuWielomianu[280].Y = 409.212

	punktyWykresuWielomianu[281].X = 2.81
	punktyWykresuWielomianu[281].Y = 411.101

	punktyWykresuWielomianu[282].X = 2.82
	punktyWykresuWielomianu[282].Y = 412.968

	punktyWykresuWielomianu[283].X = 2.83
	punktyWykresuWielomianu[283].Y = 414.812

	punktyWykresuWielomianu[284].X = 2.84
	punktyWykresuWielomianu[284].Y = 416.634

	punktyWykresuWielomianu[285].X = 2.85
	punktyWykresuWielomianu[285].Y = 418.433

	punktyWykresuWielomianu[286].X = 2.86
	punktyWykresuWielomianu[286].Y = 420.208

	punktyWykresuWielomianu[287].X = 2.87
	punktyWykresuWielomianu[287].Y = 421.959

	punktyWykresuWielomianu[288].X = 2.88
	punktyWykresuWielomianu[288].Y = 423.685

	punktyWykresuWielomianu[289].X = 2.89
	punktyWykresuWielomianu[289].Y = 425.387

	punktyWykresuWielomianu[290].X = 2.9
	punktyWykresuWielomianu[290].Y = 427.064

	punktyWykresuWielomianu[291].X = 2.91
	punktyWykresuWielomianu[291].Y = 428.714

	punktyWykresuWielomianu[292].X = 2.92
	punktyWykresuWielomianu[292].Y = 430.339

	punktyWykresuWielomianu[293].X = 2.93
	punktyWykresuWielomianu[293].Y = 431.938

	punktyWykresuWielomianu[294].X = 2.94
	punktyWykresuWielomianu[294].Y = 433.509

	punktyWykresuWielomianu[295].X = 2.95
	punktyWykresuWielomianu[295].Y = 435.053

	punktyWykresuWielomianu[296].X = 2.96
	punktyWykresuWielomianu[296].Y = 436.57

	punktyWykresuWielomianu[297].X = 2.97
	punktyWykresuWielomianu[297].Y = 438.058

	punktyWykresuWielomianu[298].X = 2.98
	punktyWykresuWielomianu[298].Y = 439.518

	punktyWykresuWielomianu[299].X = 2.99
	punktyWykresuWielomianu[299].Y = 440.948

	punktyWykresuWielomianu[300].X = 3.0
	punktyWykresuWielomianu[300].Y = 442.35

	punktyWykresuWielomianu[301].X = 3.01
	punktyWykresuWielomianu[301].Y = 443.721

	punktyWykresuWielomianu[302].X = 3.02
	punktyWykresuWielomianu[302].Y = 445.062

	punktyWykresuWielomianu[303].X = 3.03
	punktyWykresuWielomianu[303].Y = 446.373

	punktyWykresuWielomianu[304].X = 3.04
	punktyWykresuWielomianu[304].Y = 447.652

	punktyWykresuWielomianu[305].X = 3.05
	punktyWykresuWielomianu[305].Y = 448.9

	punktyWykresuWielomianu[306].X = 3.06
	punktyWykresuWielomianu[306].Y = 450.116

	punktyWykresuWielomianu[307].X = 3.07
	punktyWykresuWielomianu[307].Y = 451.299

	punktyWykresuWielomianu[308].X = 3.08
	punktyWykresuWielomianu[308].Y = 452.45

	punktyWykresuWielomianu[309].X = 3.09
	punktyWykresuWielomianu[309].Y = 453.567

	punktyWykresuWielomianu[310].X = 3.1
	punktyWykresuWielomianu[310].Y = 454.651

	punktyWykresuWielomianu[311].X = 3.11
	punktyWykresuWielomianu[311].Y = 455.7

	punktyWykresuWielomianu[312].X = 3.12
	punktyWykresuWielomianu[312].Y = 456.715

	punktyWykresuWielomianu[313].X = 3.13
	punktyWykresuWielomianu[313].Y = 457.695

	punktyWykresuWielomianu[314].X = 3.14
	punktyWykresuWielomianu[314].Y = 458.639

	punktyWykresuWielomianu[315].X = 3.15
	punktyWykresuWielomianu[315].Y = 459.548

	punktyWykresuWielomianu[316].X = 3.16
	punktyWykresuWielomianu[316].Y = 460.42

	punktyWykresuWielomianu[317].X = 3.17
	punktyWykresuWielomianu[317].Y = 461.256

	punktyWykresuWielomianu[318].X = 3.18
	punktyWykresuWielomianu[318].Y = 462.054

	punktyWykresuWielomianu[319].X = 3.19
	punktyWykresuWielomianu[319].Y = 462.815

	punktyWykresuWielomianu[320].X = 3.2
	punktyWykresuWielomianu[320].Y = 463.538

	punktyWykresuWielomianu[321].X = 3.21
	punktyWykresuWielomianu[321].Y = 464.222

	punktyWykresuWielomianu[322].X = 3.22
	punktyWykresuWielomianu[322].Y = 464.867

	punktyWykresuWielomianu[323].X = 3.23
	punktyWykresuWielomianu[323].Y = 465.473

	punktyWykresuWielomianu[324].X = 3.24
	punktyWykresuWielomianu[324].Y = 466.038

	punktyWykresuWielomianu[325].X = 3.25
	punktyWykresuWielomianu[325].Y = 466.564

	punktyWykresuWielomianu[326].X = 3.26
	punktyWykresuWielomianu[326].Y = 467.049

	punktyWykresuWielomianu[327].X = 3.27
	punktyWykresuWielomianu[327].Y = 467.493

	punktyWykresuWielomianu[328].X = 3.28
	punktyWykresuWielomianu[328].Y = 467.895

	punktyWykresuWielomianu[329].X = 3.29
	punktyWykresuWielomianu[329].Y = 468.255

	punktyWykresuWielomianu[330].X = 3.3
	punktyWykresuWielomianu[330].Y = 468.572

	punktyWykresuWielomianu[331].X = 3.31
	punktyWykresuWielomianu[331].Y = 468.846

	punktyWykresuWielomianu[332].X = 3.32
	punktyWykresuWielomianu[332].Y = 469.078

	punktyWykresuWielomianu[333].X = 3.33
	punktyWykresuWielomianu[333].Y = 469.265

	punktyWykresuWielomianu[334].X = 3.34
	punktyWykresuWielomianu[334].Y = 469.407

	punktyWykresuWielomianu[335].X = 3.35
	punktyWykresuWielomianu[335].Y = 469.505

	punktyWykresuWielomianu[336].X = 3.36
	punktyWykresuWielomianu[336].Y = 469.558

	punktyWykresuWielomianu[337].X = 3.37
	punktyWykresuWielomianu[337].Y = 469.565

	punktyWykresuWielomianu[338].X = 3.38
	punktyWykresuWielomianu[338].Y = 469.526

	punktyWykresuWielomianu[339].X = 3.39
	punktyWykresuWielomianu[339].Y = 469.441

	punktyWykresuWielomianu[340].X = 3.4
	punktyWykresuWielomianu[340].Y = 469.308

	punktyWykresuWielomianu[341].X = 3.41
	punktyWykresuWielomianu[341].Y = 469.128

	punktyWykresuWielomianu[342].X = 3.42
	punktyWykresuWielomianu[342].Y = 468.9

	punktyWykresuWielomianu[343].X = 3.43
	punktyWykresuWielomianu[343].Y = 468.623

	punktyWykresuWielomianu[344].X = 3.44
	punktyWykresuWielomianu[344].Y = 468.298

	punktyWykresuWielomianu[345].X = 3.45
	punktyWykresuWielomianu[345].Y = 467.923

	punktyWykresuWielomianu[346].X = 3.46
	punktyWykresuWielomianu[346].Y = 467.498

	punktyWykresuWielomianu[347].X = 3.47
	punktyWykresuWielomianu[347].Y = 467.024

	punktyWykresuWielomianu[348].X = 3.48
	punktyWykresuWielomianu[348].Y = 466.498

	punktyWykresuWielomianu[349].X = 3.49
	punktyWykresuWielomianu[349].Y = 465.921

	punktyWykresuWielomianu[350].X = 3.5
	punktyWykresuWielomianu[350].Y = 465.293

	punktyWykresuWielomianu[351].X = 3.51
	punktyWykresuWielomianu[351].Y = 464.613

	punktyWykresuWielomianu[352].X = 3.52
	punktyWykresuWielomianu[352].Y = 463.88

	punktyWykresuWielomianu[353].X = 3.53
	punktyWykresuWielomianu[353].Y = 463.094

	punktyWykresuWielomianu[354].X = 3.54
	punktyWykresuWielomianu[354].Y = 462.254

	punktyWykresuWielomianu[355].X = 3.55
	punktyWykresuWielomianu[355].Y = 461.361

	punktyWykresuWielomianu[356].X = 3.56
	punktyWykresuWielomianu[356].Y = 460.413

	punktyWykresuWielomianu[357].X = 3.57
	punktyWykresuWielomianu[357].Y = 459.411

	punktyWykresuWielomianu[358].X = 3.58
	punktyWykresuWielomianu[358].Y = 458.353

	punktyWykresuWielomianu[359].X = 3.59
	punktyWykresuWielomianu[359].Y = 457.239

	punktyWykresuWielomianu[360].X = 3.6
	punktyWykresuWielomianu[360].Y = 456.069

	punktyWykresuWielomianu[361].X = 3.61
	punktyWykresuWielomianu[361].Y = 454.843

	punktyWykresuWielomianu[362].X = 3.62
	punktyWykresuWielomianu[362].Y = 453.559

	punktyWykresuWielomianu[363].X = 3.63
	punktyWykresuWielomianu[363].Y = 452.217

	punktyWykresuWielomianu[364].X = 3.64
	punktyWykresuWielomianu[364].Y = 450.818

	punktyWykresuWielomianu[365].X = 3.65
	punktyWykresuWielomianu[365].Y = 449.36

	punktyWykresuWielomianu[366].X = 3.66
	punktyWykresuWielomianu[366].Y = 447.843

	punktyWykresuWielomianu[367].X = 3.67
	punktyWykresuWielomianu[367].Y = 446.266

	punktyWykresuWielomianu[368].X = 3.68
	punktyWykresuWielomianu[368].Y = 444.629

	punktyWykresuWielomianu[369].X = 3.69
	punktyWykresuWielomianu[369].Y = 442.932

	punktyWykresuWielomianu[370].X = 3.7
	punktyWykresuWielomianu[370].Y = 441.175

	punktyWykresuWielomianu[371].X = 3.71
	punktyWykresuWielomianu[371].Y = 439.355

	punktyWykresuWielomianu[372].X = 3.72
	punktyWykresuWielomianu[372].Y = 437.474

	punktyWykresuWielomianu[373].X = 3.73
	punktyWykresuWielomianu[373].Y = 435.531

	punktyWykresuWielomianu[374].X = 3.74
	punktyWykresuWielomianu[374].Y = 433.525

	punktyWykresuWielomianu[375].X = 3.75
	punktyWykresuWielomianu[375].Y = 431.455

	punktyWykresuWielomianu[376].X = 3.76
	punktyWykresuWielomianu[376].Y = 429.322

	punktyWykresuWielomianu[377].X = 3.77
	punktyWykresuWielomianu[377].Y = 427.125

	punktyWykresuWielomianu[378].X = 3.78
	punktyWykresuWielomianu[378].Y = 424.863

	punktyWykresuWielomianu[379].X = 3.79
	punktyWykresuWielomianu[379].Y = 422.536

	punktyWykresuWielomianu[380].X = 3.8
	punktyWykresuWielomianu[380].Y = 420.144

	punktyWykresuWielomianu[381].X = 3.81
	punktyWykresuWielomianu[381].Y = 417.686

	punktyWykresuWielomianu[382].X = 3.82
	punktyWykresuWielomianu[382].Y = 415.161

	punktyWykresuWielomianu[383].X = 3.83
	punktyWykresuWielomianu[383].Y = 412.569

	punktyWykresuWielomianu[384].X = 3.84
	punktyWykresuWielomianu[384].Y = 409.909

	punktyWykresuWielomianu[385].X = 3.85
	punktyWykresuWielomianu[385].Y = 407.182

	punktyWykresuWielomianu[386].X = 3.86
	punktyWykresuWielomianu[386].Y = 404.386

	punktyWykresuWielomianu[387].X = 3.87
	punktyWykresuWielomianu[387].Y = 401.522

	punktyWykresuWielomianu[388].X = 3.88
	punktyWykresuWielomianu[388].Y = 398.588

	punktyWykresuWielomianu[389].X = 3.89
	punktyWykresuWielomianu[389].Y = 395.585

	punktyWykresuWielomianu[390].X = 3.9
	punktyWykresuWielomianu[390].Y = 392.511

	punktyWykresuWielomianu[391].X = 3.91
	punktyWykresuWielomianu[391].Y = 389.367

	punktyWykresuWielomianu[392].X = 3.92
	punktyWykresuWielomianu[392].Y = 386.151

	punktyWykresuWielomianu[393].X = 3.93
	punktyWykresuWielomianu[393].Y = 382.864

	punktyWykresuWielomianu[394].X = 3.94
	punktyWykresuWielomianu[394].Y = 379.505

	punktyWykresuWielomianu[395].X = 3.95
	punktyWykresuWielomianu[395].Y = 376.073

	punktyWykresuWielomianu[396].X = 3.96
	punktyWykresuWielomianu[396].Y = 372.568

	punktyWykresuWielomianu[397].X = 3.97
	punktyWykresuWielomianu[397].Y = 368.989

	punktyWykresuWielomianu[398].X = 3.98
	punktyWykresuWielomianu[398].Y = 365.337

	punktyWykresuWielomianu[399].X = 3.99
	punktyWykresuWielomianu[399].Y = 361.61

	punktyWykresuWielomianu[400].X = 4.0
	punktyWykresuWielomianu[400].Y = 357.809

	punktyWykresuWielomianu[401].X = 4.01
	punktyWykresuWielomianu[401].Y = 353.931

	punktyWykresuWielomianu[402].X = 4.02
	punktyWykresuWielomianu[402].Y = 349.979

	punktyWykresuWielomianu[403].X = 4.03
	punktyWykresuWielomianu[403].Y = 345.949

	punktyWykresuWielomianu[404].X = 4.04
	punktyWykresuWielomianu[404].Y = 341.843

	punktyWykresuWielomianu[405].X = 4.05
	punktyWykresuWielomianu[405].Y = 337.66

	punktyWykresuWielomianu[406].X = 4.06
	punktyWykresuWielomianu[406].Y = 333.399

	punktyWykresuWielomianu[407].X = 4.07
	punktyWykresuWielomianu[407].Y = 329.06

	punktyWykresuWielomianu[408].X = 4.08
	punktyWykresuWielomianu[408].Y = 324.643

	punktyWykresuWielomianu[409].X = 4.09
	punktyWykresuWielomianu[409].Y = 320.146

	punktyWykresuWielomianu[410].X = 4.1
	punktyWykresuWielomianu[410].Y = 315.57

	punktyWykresuWielomianu[411].X = 4.11
	punktyWykresuWielomianu[411].Y = 310.914

	punktyWykresuWielomianu[412].X = 4.12
	punktyWykresuWielomianu[412].Y = 306.177

	punktyWykresuWielomianu[413].X = 4.13
	punktyWykresuWielomianu[413].Y = 301.359

	punktyWykresuWielomianu[414].X = 4.14
	punktyWykresuWielomianu[414].Y = 296.46

	punktyWykresuWielomianu[415].X = 4.15
	punktyWykresuWielomianu[415].Y = 291.479

	punktyWykresuWielomianu[416].X = 4.16
	punktyWykresuWielomianu[416].Y = 286.416

	punktyWykresuWielomianu[417].X = 4.17
	punktyWykresuWielomianu[417].Y = 281.27

	punktyWykresuWielomianu[418].X = 4.18
	punktyWykresuWielomianu[418].Y = 276.04

	punktyWykresuWielomianu[419].X = 4.19
	punktyWykresuWielomianu[419].Y = 270.727

	punktyWykresuWielomianu[420].X = 4.20
	punktyWykresuWielomianu[420].Y = 265.33

	punktyWykresuWielomianu[421].X = 4.21
	punktyWykresuWielomianu[421].Y = 259.848

	punktyWykresuWielomianu[422].X = 4.22
	punktyWykresuWielomianu[422].Y = 254.282

	punktyWykresuWielomianu[423].X = 4.23
	punktyWykresuWielomianu[423].Y = 248.629

	punktyWykresuWielomianu[424].X = 4.24
	punktyWykresuWielomianu[424].Y = 242.891

	punktyWykresuWielomianu[425].X = 4.25
	punktyWykresuWielomianu[425].Y = 237.066

	punktyWykresuWielomianu[426].X = 4.26
	punktyWykresuWielomianu[426].Y = 231.154

	punktyWykresuWielomianu[427].X = 4.27
	punktyWykresuWielomianu[427].Y = 225.155

	punktyWykresuWielomianu[428].X = 4.28
	punktyWykresuWielomianu[428].Y = 219.068

	punktyWykresuWielomianu[429].X = 4.29
	punktyWykresuWielomianu[429].Y = 212.892

	punktyWykresuWielomianu[430].X = 4.30
	punktyWykresuWielomianu[430].Y = 206.628

	punktyWykresuWielomianu[431].X = 4.31
	punktyWykresuWielomianu[431].Y = 200.275

	punktyWykresuWielomianu[432].X = 4.32
	punktyWykresuWielomianu[432].Y = 193.832

	punktyWykresuWielomianu[433].X = 4.33
	punktyWykresuWielomianu[433].Y = 187.299

	punktyWykresuWielomianu[434].X = 4.34
	punktyWykresuWielomianu[434].Y = 180.675

	punktyWykresuWielomianu[435].X = 4.35
	punktyWykresuWielomianu[435].Y = 173.961

	punktyWykresuWielomianu[436].X = 4.36
	punktyWykresuWielomianu[436].Y = 167.155

	punktyWykresuWielomianu[437].X = 4.37
	punktyWykresuWielomianu[437].Y = 160.257

	punktyWykresuWielomianu[438].X = 4.38
	punktyWykresuWielomianu[438].Y = 153.266

	punktyWykresuWielomianu[439].X = 4.39
	punktyWykresuWielomianu[439].Y = 146.183

	punktyWykresuWielomianu[440].X = 4.40
	punktyWykresuWielomianu[440].Y = 139.007

	punktyWykresuWielomianu[441].X = 4.41
	punktyWykresuWielomianu[441].Y = 131.736

	punktyWykresuWielomianu[442].X = 4.42
	punktyWykresuWielomianu[442].Y = 124.372

	punktyWykresuWielomianu[443].X = 4.43
	punktyWykresuWielomianu[443].Y = 116.913

	punktyWykresuWielomianu[444].X = 4.44
	punktyWykresuWielomianu[444].Y = 109.359

	punktyWykresuWielomianu[445].X = 4.45
	punktyWykresuWielomianu[445].Y = 101.709

	punktyWykresuWielomianu[446].X = 4.46
	punktyWykresuWielomianu[446].Y = 93.964

	punktyWykresuWielomianu[447].X = 4.47
	punktyWykresuWielomianu[447].Y = 86.122

	punktyWykresuWielomianu[448].X = 4.48
	punktyWykresuWielomianu[448].Y = 78.183

	punktyWykresuWielomianu[449].X = 4.49
	punktyWykresuWielomianu[449].Y = 70.147

	punktyWykresuWielomianu[450].X = 4.5
	punktyWykresuWielomianu[450].Y = 62.013

	punktyWykresuWielomianu[451].X = 4.51
	punktyWykresuWielomianu[451].Y = 53.781

	punktyWykresuWielomianu[452].X = 4.52
	punktyWykresuWielomianu[452].Y = 45.451

	punktyWykresuWielomianu[453].X = 4.53
	punktyWykresuWielomianu[453].Y = 37.021

	punktyWykresuWielomianu[454].X = 4.54
	punktyWykresuWielomianu[454].Y = 28.492

	punktyWykresuWielomianu[455].X = 4.55
	punktyWykresuWielomianu[455].Y = 19.863

	punktyWykresuWielomianu[456].X = 4.56
	punktyWykresuWielomianu[456].Y = 11.133

	punktyWykresuWielomianu[457].X = 4.57
	punktyWykresuWielomianu[457].Y = 2.303

	punktyWykresuWielomianu[458].X = 4.58
	punktyWykresuWielomianu[458].Y = -6.627

	punktyWykresuWielomianu[459].X = 4.59
	punktyWykresuWielomianu[459].Y = -15.661

	punktyWykresuWielomianu[460].X = 4.6
	punktyWykresuWielomianu[460].Y = -24.796

	punktyWykresuWielomianu[461].X = 4.61
	punktyWykresuWielomianu[461].Y = -34.034

	punktyWykresuWielomianu[462].X = 4.62
	punktyWykresuWielomianu[462].Y = -43.375

	punktyWykresuWielomianu[463].X = 4.63
	punktyWykresuWielomianu[463].Y = -52.819

	punktyWykresuWielomianu[464].X = 4.64
	punktyWykresuWielomianu[464].Y = -62.367

	punktyWykresuWielomianu[465].X = 4.65
	punktyWykresuWielomianu[465].Y = -72.019

	punktyWykresuWielomianu[466].X = 4.66
	punktyWykresuWielomianu[466].Y = -81.776

	punktyWykresuWielomianu[467].X = 4.67
	punktyWykresuWielomianu[467].Y = -91.639

	punktyWykresuWielomianu[468].X = 4.68
	punktyWykresuWielomianu[468].Y = -101.606

	punktyWykresuWielomianu[469].X = 4.69
	punktyWykresuWielomianu[469].Y = -111.68

	punktyWykresuWielomianu[470].X = 4.7
	punktyWykresuWielomianu[470].Y = -121.86

	punktyWykresuWielomianu[471].X = 4.71
	punktyWykresuWielomianu[471].Y = -132.147

	punktyWykresuWielomianu[472].X = 4.72
	punktyWykresuWielomianu[472].Y = -142.541

	punktyWykresuWielomianu[473].X = 4.73
	punktyWykresuWielomianu[473].Y = -153.043

	punktyWykresuWielomianu[474].X = 4.74
	punktyWykresuWielomianu[474].Y = -163.653

	punktyWykresuWielomianu[475].X = 4.75
	punktyWykresuWielomianu[475].Y = -174.372

	punktyWykresuWielomianu[476].X = 4.76
	punktyWykresuWielomianu[476].Y = -185.199

	punktyWykresuWielomianu[477].X = 4.77
	punktyWykresuWielomianu[477].Y = -196.136

	punktyWykresuWielomianu[478].X = 4.78
	punktyWykresuWielomianu[478].Y = -207.183

	punktyWykresuWielomianu[479].X = 4.79
	punktyWykresuWielomianu[479].Y = -218.339

	punktyWykresuWielomianu[480].X = 4.8
	punktyWykresuWielomianu[480].Y = -229.607

	punktyWykresuWielomianu[481].X = 4.81
	punktyWykresuWielomianu[481].Y = -240.985

	punktyWykresuWielomianu[482].X = 4.82
	punktyWykresuWielomianu[482].Y = -252.475

	punktyWykresuWielomianu[483].X = 4.83
	punktyWykresuWielomianu[483].Y = -264.077

	punktyWykresuWielomianu[484].X = 4.84
	punktyWykresuWielomianu[484].Y = -275.79

	punktyWykresuWielomianu[485].X = 4.85
	punktyWykresuWielomianu[485].Y = -287.617

	punktyWykresuWielomianu[486].X = 4.86
	punktyWykresuWielomianu[486].Y = -299.557

	punktyWykresuWielomianu[487].X = 4.87
	punktyWykresuWielomianu[487].Y = -311.61

	punktyWykresuWielomianu[488].X = 4.88
	punktyWykresuWielomianu[488].Y = -323.777

	punktyWykresuWielomianu[489].X = 4.89
	punktyWykresuWielomianu[489].Y = -336.058

	punktyWykresuWielomianu[490].X = 4.9
	punktyWykresuWielomianu[490].Y = -348.454

	punktyWykresuWielomianu[491].X = 4.91
	punktyWykresuWielomianu[491].Y = -360.965

	punktyWykresuWielomianu[492].X = 4.92
	punktyWykresuWielomianu[492].Y = -373.591

	punktyWykresuWielomianu[493].X = 4.93
	punktyWykresuWielomianu[493].Y = -386.333

	punktyWykresuWielomianu[494].X = 4.94
	punktyWykresuWielomianu[494].Y = -399.192

	punktyWykresuWielomianu[495].X = 4.95
	punktyWykresuWielomianu[495].Y = -412.167

	punktyWykresuWielomianu[496].X = 4.96
	punktyWykresuWielomianu[496].Y = -425.26

	punktyWykresuWielomianu[497].X = 4.97
	punktyWykresuWielomianu[497].Y = -438.469

	punktyWykresuWielomianu[498].X = 4.98
	punktyWykresuWielomianu[498].Y = -451.797

	punktyWykresuWielomianu[499].X = 4.99
	punktyWykresuWielomianu[499].Y = -465.243

	punktyWykresuWielomianu[500].X = 5.0
	punktyWykresuWielomianu[500].Y = -478.808

	punktyWykresuWielomianu[501].X = 5.01
	punktyWykresuWielomianu[501].Y = -492.491

	punktyWykresuWielomianu[502].X = 5.02
	punktyWykresuWielomianu[502].Y = -506.294

	punktyWykresuWielomianu[503].X = 5.03
	punktyWykresuWielomianu[503].Y = -520.217

	punktyWykresuWielomianu[504].X = 5.04
	punktyWykresuWielomianu[504].Y = -534.26

	punktyWykresuWielomianu[505].X = 5.05
	punktyWykresuWielomianu[505].Y = -548.423

	punktyWykresuWielomianu[506].X = 5.06
	punktyWykresuWielomianu[506].Y = -562.708

	punktyWykresuWielomianu[507].X = 5.07
	punktyWykresuWielomianu[507].Y = -577.113

	punktyWykresuWielomianu[508].X = 5.08
	punktyWykresuWielomianu[508].Y = -591.641

	punktyWykresuWielomianu[509].X = 5.09
	punktyWykresuWielomianu[509].Y = -606.29

	punktyWykresuWielomianu[510].X = 5.1
	punktyWykresuWielomianu[510].Y = -621.062

	punktyWykresuWielomianu[511].X = 5.11
	punktyWykresuWielomianu[511].Y = -635.957

	punktyWykresuWielomianu[512].X = 5.12
	punktyWykresuWielomianu[512].Y = -650.975

	punktyWykresuWielomianu[513].X = 5.13
	punktyWykresuWielomianu[513].Y = -666.116

	punktyWykresuWielomianu[514].X = 5.14
	punktyWykresuWielomianu[514].Y = -681.382

	punktyWykresuWielomianu[515].X = 5.15
	punktyWykresuWielomianu[515].Y = -696.772

	punktyWykresuWielomianu[516].X = 5.16
	punktyWykresuWielomianu[516].Y = -712.286

	punktyWykresuWielomianu[517].X = 5.17
	punktyWykresuWielomianu[517].Y = -727.925

	punktyWykresuWielomianu[518].X = 5.18
	punktyWykresuWielomianu[518].Y = -743.69

	punktyWykresuWielomianu[519].X = 5.19
	punktyWykresuWielomianu[519].Y = -759.581

	punktyWykresuWielomianu[520].X = 5.2
	punktyWykresuWielomianu[520].Y = -775.597

	punktyWykresuWielomianu[521].X = 5.21
	punktyWykresuWielomianu[521].Y = -791.741

	punktyWykresuWielomianu[522].X = 5.22
	punktyWykresuWielomianu[522].Y = -808.011

	punktyWykresuWielomianu[523].X = 5.23
	punktyWykresuWielomianu[523].Y = -824.408

	punktyWykresuWielomianu[524].X = 5.24
	punktyWykresuWielomianu[524].Y = -840.932

	punktyWykresuWielomianu[525].X = 5.25
	punktyWykresuWielomianu[525].Y = -857.585

	punktyWykresuWielomianu[526].X = 5.26
	punktyWykresuWielomianu[526].Y = -874.366

	punktyWykresuWielomianu[527].X = 5.27
	punktyWykresuWielomianu[527].Y = -891.275

	punktyWykresuWielomianu[528].X = 5.28
	punktyWykresuWielomianu[528].Y = -908.314

	punktyWykresuWielomianu[529].X = 5.29
	punktyWykresuWielomianu[529].Y = -925.481

	punktyWykresuWielomianu[530].X = 5.3
	punktyWykresuWielomianu[530].Y = -942.778

	punktyWykresuWielomianu[531].X = 5.31
	punktyWykresuWielomianu[531].Y = -960.206

	punktyWykresuWielomianu[532].X = 5.32
	punktyWykresuWielomianu[532].Y = -977.763

	punktyWykresuWielomianu[533].X = 5.33
	punktyWykresuWielomianu[533].Y = -995.451

	punktyWykresuWielomianu[534].X = 5.34
	punktyWykresuWielomianu[534].Y = -1_013.27

	punktyWykresuWielomianu[535].X = 5.35
	punktyWykresuWielomianu[535].Y = -1_031.221

	punktyWykresuWielomianu[536].X = 5.36
	punktyWykresuWielomianu[536].Y = -1_049.303

	punktyWykresuWielomianu[537].X = 5.37
	punktyWykresuWielomianu[537].Y = -1_067.517

	punktyWykresuWielomianu[538].X = 5.38
	punktyWykresuWielomianu[538].Y = -1_085.863

	punktyWykresuWielomianu[539].X = 5.39
	punktyWykresuWielomianu[539].Y = -1_104.342

	punktyWykresuWielomianu[540].X = 5.4
	punktyWykresuWielomianu[540].Y = -1_122.954

	punktyWykresuWielomianu[541].X = 5.41
	punktyWykresuWielomianu[541].Y = -1_141.699

	punktyWykresuWielomianu[542].X = 5.42
	punktyWykresuWielomianu[542].Y = -1_160.578

	punktyWykresuWielomianu[543].X = 5.43
	punktyWykresuWielomianu[543].Y = -1_179.59

	punktyWykresuWielomianu[544].X = 5.44
	punktyWykresuWielomianu[544].Y = -1_198.737

	punktyWykresuWielomianu[545].X = 5.45
	punktyWykresuWielomianu[545].Y = -1_218.018

	punktyWykresuWielomianu[546].X = 5.46
	punktyWykresuWielomianu[546].Y = -1_237.435

	punktyWykresuWielomianu[547].X = 5.47
	punktyWykresuWielomianu[547].Y = -1_256.986

	punktyWykresuWielomianu[548].X = 5.48
	punktyWykresuWielomianu[548].Y = -1_276.673

	punktyWykresuWielomianu[549].X = 5.49
	punktyWykresuWielomianu[549].Y = -1_296.496

	punktyWykresuWielomianu[550].X = 5.5
	punktyWykresuWielomianu[550].Y = -1_316.455

	punktyWykresuWielomianu[551].X = 5.51
	punktyWykresuWielomianu[551].Y = -1_336.55

	punktyWykresuWielomianu[552].X = 5.52
	punktyWykresuWielomianu[552].Y = -1_356.782

	punktyWykresuWielomianu[553].X = 5.53
	punktyWykresuWielomianu[553].Y = -1_377.15

	punktyWykresuWielomianu[554].X = 5.54
	punktyWykresuWielomianu[554].Y = -1_397.657

	punktyWykresuWielomianu[555].X = 5.55
	punktyWykresuWielomianu[555].Y = -1_418.3

	punktyWykresuWielomianu[556].X = 5.56
	punktyWykresuWielomianu[556].Y = -1_439.082

	punktyWykresuWielomianu[557].X = 5.57
	punktyWykresuWielomianu[557].Y = -1_460.002

	punktyWykresuWielomianu[558].X = 5.58
	punktyWykresuWielomianu[558].Y = -1_481.06

	punktyWykresuWielomianu[559].X = 5.59
	punktyWykresuWielomianu[559].Y = -1_502.257

	punktyWykresuWielomianu[560].X = 5.6
	punktyWykresuWielomianu[560].Y = -1_523.593

	punktyWykresuWielomianu[561].X = 5.61
	punktyWykresuWielomianu[561].Y = -1_545.069

	punktyWykresuWielomianu[562].X = 5.62
	punktyWykresuWielomianu[562].Y = -1_566.684

	punktyWykresuWielomianu[563].X = 5.63
	punktyWykresuWielomianu[563].Y = -1_588.439

	punktyWykresuWielomianu[564].X = 5.64
	punktyWykresuWielomianu[564].Y = -1_610.334

	punktyWykresuWielomianu[565].X = 5.65
	punktyWykresuWielomianu[565].Y = -1_632.369

	punktyWykresuWielomianu[566].X = 5.66
	punktyWykresuWielomianu[566].Y = -1_654.545

	punktyWykresuWielomianu[567].X = 5.67
	punktyWykresuWielomianu[567].Y = -1_676.863

	punktyWykresuWielomianu[568].X = 5.68
	punktyWykresuWielomianu[568].Y = -1_699.321

	punktyWykresuWielomianu[569].X = 5.69
	punktyWykresuWielomianu[569].Y = -1_721.921

	punktyWykresuWielomianu[570].X = 5.7
	punktyWykresuWielomianu[570].Y = -1_744.663

	punktyWykresuWielomianu[571].X = 5.71
	punktyWykresuWielomianu[571].Y = -1_767.547

	punktyWykresuWielomianu[572].X = 5.72
	punktyWykresuWielomianu[572].Y = -1_790.573

	punktyWykresuWielomianu[573].X = 5.73
	punktyWykresuWielomianu[573].Y = -1_813.741

	punktyWykresuWielomianu[574].X = 5.74
	punktyWykresuWielomianu[574].Y = -1_837.053

	punktyWykresuWielomianu[575].X = 5.75
	punktyWykresuWielomianu[575].Y = -1_860.507

	punktyWykresuWielomianu[576].X = 5.76
	punktyWykresuWielomianu[576].Y = -1_884.105

	punktyWykresuWielomianu[577].X = 5.77
	punktyWykresuWielomianu[577].Y = -1_907.847

	punktyWykresuWielomianu[578].X = 5.78
	punktyWykresuWielomianu[578].Y = -1_931.732

	punktyWykresuWielomianu[579].X = 5.79
	punktyWykresuWielomianu[579].Y = -1_955.761

	punktyWykresuWielomianu[580].X = 5.8
	punktyWykresuWielomianu[580].Y = -1_979.935

	punktyWykresuWielomianu[581].X = 5.81
	punktyWykresuWielomianu[581].Y = -2_004.253

	punktyWykresuWielomianu[582].X = 5.82
	punktyWykresuWielomianu[582].Y = -2_028.716

	punktyWykresuWielomianu[583].X = 5.83
	punktyWykresuWielomianu[583].Y = -2_053.323

	punktyWykresuWielomianu[584].X = 5.84
	punktyWykresuWielomianu[584].Y = -2_078.076

	punktyWykresuWielomianu[585].X = 5.85
	punktyWykresuWielomianu[585].Y = -2_102.975

	punktyWykresuWielomianu[586].X = 5.86
	punktyWykresuWielomianu[586].Y = -2_128.019

	punktyWykresuWielomianu[587].X = 5.87
	punktyWykresuWielomianu[587].Y = -2_153.21

	punktyWykresuWielomianu[588].X = 5.88
	punktyWykresuWielomianu[588].Y = -2_178.546

	punktyWykresuWielomianu[589].X = 5.89
	punktyWykresuWielomianu[589].Y = -2_204.029

	punktyWykresuWielomianu[590].X = 5.9
	punktyWykresuWielomianu[590].Y = -2_229.658

	punktyWykresuWielomianu[591].X = 5.91
	punktyWykresuWielomianu[591].Y = -2_255.434

	punktyWykresuWielomianu[592].X = 5.92
	punktyWykresuWielomianu[592].Y = -2_281.357

	punktyWykresuWielomianu[593].X = 5.93
	punktyWykresuWielomianu[593].Y = -2_307.427

	punktyWykresuWielomianu[594].X = 5.94
	punktyWykresuWielomianu[594].Y = -2_333.645

	punktyWykresuWielomianu[595].X = 5.95
	punktyWykresuWielomianu[595].Y = -2_360.01

	punktyWykresuWielomianu[596].X = 5.96
	punktyWykresuWielomianu[596].Y = -2_386.524

	punktyWykresuWielomianu[597].X = 5.97
	punktyWykresuWielomianu[597].Y = -2_413.185

	punktyWykresuWielomianu[598].X = 5.98
	punktyWykresuWielomianu[598].Y = -2_439.994

	punktyWykresuWielomianu[599].X = 5.99
	punktyWykresuWielomianu[599].Y = -2_466.952

	punktyWykresuWielomianu[600].X = 6.0
	punktyWykresuWielomianu[600].Y = -2_494.059

	punktyWykresuWielomianu[601].X = 6.01
	punktyWykresuWielomianu[601].Y = -2_521.314

	punktyWykresuWielomianu[602].X = 6.02
	punktyWykresuWielomianu[602].Y = -2_548.718

	punktyWykresuWielomianu[603].X = 6.03
	punktyWykresuWielomianu[603].Y = -2_576.272

	punktyWykresuWielomianu[604].X = 6.04
	punktyWykresuWielomianu[604].Y = -2_603.974

	punktyWykresuWielomianu[605].X = 6.05
	punktyWykresuWielomianu[605].Y = -2_631.827

	punktyWykresuWielomianu[606].X = 6.06
	punktyWykresuWielomianu[606].Y = -2_659.829

	punktyWykresuWielomianu[607].X = 6.07
	punktyWykresuWielomianu[607].Y = -2_687.981

	punktyWykresuWielomianu[608].X = 6.08
	punktyWykresuWielomianu[608].Y = -2_716.283

	punktyWykresuWielomianu[609].X = 6.09
	punktyWykresuWielomianu[609].Y = -2_744.735

	punktyWykresuWielomianu[610].X = 6.1
	punktyWykresuWielomianu[610].Y = -2_773.338

	punktyWykresuWielomianu[611].X = 6.11
	punktyWykresuWielomianu[611].Y = -2_802.091

	punktyWykresuWielomianu[612].X = 6.12
	punktyWykresuWielomianu[612].Y = -2_830.995

	punktyWykresuWielomianu[613].X = 6.13
	punktyWykresuWielomianu[613].Y = -2_860.05

	punktyWykresuWielomianu[614].X = 6.14
	punktyWykresuWielomianu[614].Y = -2_889.256

	punktyWykresuWielomianu[615].X = 6.15
	punktyWykresuWielomianu[615].Y = -2_918.613

	punktyWykresuWielomianu[616].X = 6.16
	punktyWykresuWielomianu[616].Y = -2_948.121

	punktyWykresuWielomianu[617].X = 6.17
	punktyWykresuWielomianu[617].Y = -2_977.781

	punktyWykresuWielomianu[618].X = 6.18
	punktyWykresuWielomianu[618].Y = -3_007.593

	punktyWykresuWielomianu[619].X = 6.19
	punktyWykresuWielomianu[619].Y = -3_037.556

	punktyWykresuWielomianu[620].X = 6.2
	punktyWykresuWielomianu[620].Y = -3_067.671

	punktyWykresuWielomianu[621].X = 6.21
	punktyWykresuWielomianu[621].Y = -3_097.938

	punktyWykresuWielomianu[622].X = 6.22
	punktyWykresuWielomianu[622].Y = -3_128.358

	punktyWykresuWielomianu[623].X = 6.23
	punktyWykresuWielomianu[623].Y = -3_158.929

	punktyWykresuWielomianu[624].X = 6.24
	punktyWykresuWielomianu[624].Y = -3_189.654

	punktyWykresuWielomianu[625].X = 6.25
	punktyWykresuWielomianu[625].Y = -3_220.53

	punktyWykresuWielomianu[626].X = 6.26
	punktyWykresuWielomianu[626].Y = -3_251.56

	punktyWykresuWielomianu[627].X = 6.27
	punktyWykresuWielomianu[627].Y = -3_282.742

	punktyWykresuWielomianu[628].X = 6.28
	punktyWykresuWielomianu[628].Y = -3_314.077

	punktyWykresuWielomianu[629].X = 6.29
	punktyWykresuWielomianu[629].Y = -3_345.565

	punktyWykresuWielomianu[630].X = 6.3
	punktyWykresuWielomianu[630].Y = -3_377.207

	punktyWykresuWielomianu[631].X = 6.31
	punktyWykresuWielomianu[631].Y = -3_409.001

	punktyWykresuWielomianu[632].X = 6.32
	punktyWykresuWielomianu[632].Y = -3_440.949

	punktyWykresuWielomianu[633].X = 6.33
	punktyWykresuWielomianu[633].Y = -3_473.714

	punktyWykresuWielomianu[634].X = 6.34
	punktyWykresuWielomianu[634].Y = -3_505.276

	punktyWykresuWielomianu[635].X = 6.35
	punktyWykresuWielomianu[635].Y = -3_537.714

	punktyWykresuWielomianu[636].X = 6.36
	punktyWykresuWielomianu[636].Y = -3_570.276

	punktyWykresuWielomianu[637].X = 6.37
	punktyWykresuWielomianu[637].Y = -3_602.992

	punktyWykresuWielomianu[638].X = 6.38
	punktyWykresuWielomianu[638].Y = -3_635.862

	punktyWykresuWielomianu[639].X = 6.39
	punktyWykresuWielomianu[639].Y = -3_668.886

	punktyWykresuWielomianu[640].X = 6.4
	punktyWykresuWielomianu[640].Y = -3_702.064

	punktyWykresuWielomianu[641].X = 6.41
	punktyWykresuWielomianu[641].Y = -3_735.396

	punktyWykresuWielomianu[642].X = 6.42
	punktyWykresuWielomianu[642].Y = -3_768.883

	punktyWykresuWielomianu[643].X = 6.43
	punktyWykresuWielomianu[643].Y = -3_802.523

	punktyWykresuWielomianu[644].X = 6.44
	punktyWykresuWielomianu[644].Y = -3_836.319

	punktyWykresuWielomianu[645].X = 6.45
	punktyWykresuWielomianu[645].Y = -3_870.268

	punktyWykresuWielomianu[646].X = 6.46
	punktyWykresuWielomianu[646].Y = -3_904.372

	punktyWykresuWielomianu[647].X = 6.47
	punktyWykresuWielomianu[647].Y = -3_938.631

	punktyWykresuWielomianu[648].X = 6.48
	punktyWykresuWielomianu[648].Y = -3_973.044

	punktyWykresuWielomianu[649].X = 6.49
	punktyWykresuWielomianu[649].Y = -4_007.612

	punktyWykresuWielomianu[650].X = 6.5
	punktyWykresuWielomianu[650].Y = -4_042.334

	punktyWykresuWielomianu[651].X = 6.51
	punktyWykresuWielomianu[651].Y = -4_077.212

	punktyWykresuWielomianu[652].X = 6.52
	punktyWykresuWielomianu[652].Y = -4_112.244

	punktyWykresuWielomianu[653].X = 6.53
	punktyWykresuWielomianu[653].Y = -4_147.431

	punktyWykresuWielomianu[654].X = 6.54
	punktyWykresuWielomianu[654].Y = -4_182.772

	punktyWykresuWielomianu[655].X = 6.55
	punktyWykresuWielomianu[655].Y = -4_218.269

	punktyWykresuWielomianu[656].X = 6.56
	punktyWykresuWielomianu[656].Y = -4_253.921

	punktyWykresuWielomianu[657].X = 6.57
	punktyWykresuWielomianu[657].Y = -4_289.727

	punktyWykresuWielomianu[658].X = 6.58
	punktyWykresuWielomianu[658].Y = -4_325.689

	punktyWykresuWielomianu[659].X = 6.59
	punktyWykresuWielomianu[659].Y = -4_361.805

	punktyWykresuWielomianu[660].X = 6.6
	punktyWykresuWielomianu[660].Y = -4_398.077

	punktyWykresuWielomianu[661].X = 6.61
	punktyWykresuWielomianu[661].Y = -4_434.504

	punktyWykresuWielomianu[662].X = 6.62
	punktyWykresuWielomianu[662].Y = -4_471.085

	punktyWykresuWielomianu[663].X = 6.63
	punktyWykresuWielomianu[663].Y = -4_507.822

	punktyWykresuWielomianu[664].X = 6.64
	punktyWykresuWielomianu[664].Y = -4_544.714

	punktyWykresuWielomianu[665].X = 6.65
	punktyWykresuWielomianu[665].Y = -4_581.761

	punktyWykresuWielomianu[666].X = 6.66
	punktyWykresuWielomianu[666].Y = -4_618.962

	punktyWykresuWielomianu[667].X = 6.67
	punktyWykresuWielomianu[667].Y = -4_656.319

	punktyWykresuWielomianu[668].X = 6.68
	punktyWykresuWielomianu[668].Y = -4_693.831

	punktyWykresuWielomianu[669].X = 6.69
	punktyWykresuWielomianu[669].Y = -4_731.498

	punktyWykresuWielomianu[670].X = 6.7
	punktyWykresuWielomianu[670].Y = -4_769.32

	punktyWykresuWielomianu[671].X = 6.71
	punktyWykresuWielomianu[671].Y = -4_807.297

	punktyWykresuWielomianu[672].X = 6.72
	punktyWykresuWielomianu[672].Y = -4_845.429

	punktyWykresuWielomianu[673].X = 6.73
	punktyWykresuWielomianu[673].Y = -4_883.716

	punktyWykresuWielomianu[674].X = 6.74
	punktyWykresuWielomianu[674].Y = -4_922.158

	punktyWykresuWielomianu[675].X = 6.75
	punktyWykresuWielomianu[675].Y = -4_960.754

	punktyWykresuWielomianu[676].X = 6.76
	punktyWykresuWielomianu[676].Y = -4_999.506

	punktyWykresuWielomianu[677].X = 6.77
	punktyWykresuWielomianu[677].Y = -5_038.412

	punktyWykresuWielomianu[678].X = 6.78
	punktyWykresuWielomianu[678].Y = -5_077.473

	punktyWykresuWielomianu[679].X = 6.79
	punktyWykresuWielomianu[679].Y = -5_116.689

	punktyWykresuWielomianu[680].X = 6.8
	punktyWykresuWielomianu[680].Y = -5_156.059

	punktyWykresuWielomianu[681].X = 6.81
	punktyWykresuWielomianu[681].Y = -5_195.584

	punktyWykresuWielomianu[682].X = 6.82
	punktyWykresuWielomianu[682].Y = -5_235.263

	punktyWykresuWielomianu[683].X = 6.83
	punktyWykresuWielomianu[683].Y = -5_275.097

	punktyWykresuWielomianu[684].X = 6.84
	punktyWykresuWielomianu[684].Y = -5_315.085

	punktyWykresuWielomianu[685].X = 6.85
	punktyWykresuWielomianu[685].Y = -5_355.228

	punktyWykresuWielomianu[686].X = 6.86
	punktyWykresuWielomianu[686].Y = -5_395.525

	punktyWykresuWielomianu[687].X = 6.87
	punktyWykresuWielomianu[687].Y = -5_435.976

	punktyWykresuWielomianu[688].X = 6.88
	punktyWykresuWielomianu[688].Y = -5_476.581

	punktyWykresuWielomianu[689].X = 6.89
	punktyWykresuWielomianu[689].Y = -5_517.34

	punktyWykresuWielomianu[690].X = 6.9
	punktyWykresuWielomianu[690].Y = -5_558.254

	punktyWykresuWielomianu[691].X = 6.91
	punktyWykresuWielomianu[691].Y = -5_599.321

	punktyWykresuWielomianu[692].X = 6.92
	punktyWykresuWielomianu[692].Y = -5_640.541

	punktyWykresuWielomianu[693].X = 6.93
	punktyWykresuWielomianu[693].Y = -5_681.916

	punktyWykresuWielomianu[694].X = 6.94
	punktyWykresuWielomianu[694].Y = -5_723.444

	punktyWykresuWielomianu[695].X = 6.95
	punktyWykresuWielomianu[695].Y = -5_765.125

	punktyWykresuWielomianu[696].X = 6.96
	punktyWykresuWielomianu[696].Y = -5_806.96

	punktyWykresuWielomianu[697].X = 6.97
	punktyWykresuWielomianu[697].Y = -5_848.948

	punktyWykresuWielomianu[698].X = 6.98
	punktyWykresuWielomianu[698].Y = -5_891.089

	punktyWykresuWielomianu[699].X = 6.99
	punktyWykresuWielomianu[699].Y = -5_933.383

	punktyWykresuWielomianu[700].X = 7.0
	punktyWykresuWielomianu[700].Y = -5_975.83

	punktyWykresuWielomianu[701].X = 7.01
	punktyWykresuWielomianu[701].Y = -6_018.429

	punktyWykresuWielomianu[702].X = 7.02
	punktyWykresuWielomianu[702].Y = -6_061.181

	punktyWykresuWielomianu[703].X = 7.03
	punktyWykresuWielomianu[703].Y = -6_104.086

	punktyWykresuWielomianu[704].X = 7.04
	punktyWykresuWielomianu[704].Y = -6_147.143

	punktyWykresuWielomianu[705].X = 7.05
	punktyWykresuWielomianu[705].Y = -6_190.352

	punktyWykresuWielomianu[706].X = 7.06
	punktyWykresuWielomianu[706].Y = -6_233.713

	punktyWykresuWielomianu[707].X = 7.07
	punktyWykresuWielomianu[707].Y = -6_277.226

	punktyWykresuWielomianu[708].X = 7.08
	punktyWykresuWielomianu[708].Y = -6_320.891

	punktyWykresuWielomianu[709].X = 7.09
	punktyWykresuWielomianu[709].Y = -6_364.707

	punktyWykresuWielomianu[710].X = 7.1
	punktyWykresuWielomianu[710].Y = -6_408.675

	punktyWykresuWielomianu[711].X = 7.11
	punktyWykresuWielomianu[711].Y = -6_452.794

	punktyWykresuWielomianu[712].X = 7.12
	punktyWykresuWielomianu[712].Y = -6_497.064

	punktyWykresuWielomianu[713].X = 7.13
	punktyWykresuWielomianu[713].Y = -6_541.485

	punktyWykresuWielomianu[714].X = 7.14
	punktyWykresuWielomianu[714].Y = -6_586.057

	punktyWykresuWielomianu[715].X = 7.15
	punktyWykresuWielomianu[715].Y = -6_630.779

	punktyWykresuWielomianu[716].X = 7.16
	punktyWykresuWielomianu[716].Y = -6_675.652

	punktyWykresuWielomianu[717].X = 7.17
	punktyWykresuWielomianu[717].Y = -6_720.674

	punktyWykresuWielomianu[718].X = 7.18
	punktyWykresuWielomianu[718].Y = -6_765.847

	punktyWykresuWielomianu[719].X = 7.19
	punktyWykresuWielomianu[719].Y = -6_811.169

	punktyWykresuWielomianu[720].X = 7.2
	punktyWykresuWielomianu[720].Y = -6_856.641

	punktyWykresuWielomianu[721].X = 7.21
	punktyWykresuWielomianu[721].Y = -6_902.263

	punktyWykresuWielomianu[722].X = 7.22
	punktyWykresuWielomianu[722].Y = -6_948.033

	punktyWykresuWielomianu[723].X = 7.23
	punktyWykresuWielomianu[723].Y = -6_993.952

	punktyWykresuWielomianu[724].X = 7.24
	punktyWykresuWielomianu[724].Y = -7_040.021

	punktyWykresuWielomianu[725].X = 7.25
	punktyWykresuWielomianu[725].Y = -7_086.237

	punktyWykresuWielomianu[726].X = 7.26
	punktyWykresuWielomianu[726].Y = -7_132.602

	punktyWykresuWielomianu[727].X = 7.27
	punktyWykresuWielomianu[727].Y = -7_179.115

	punktyWykresuWielomianu[728].X = 7.28
	punktyWykresuWielomianu[728].Y = -7_225.776

	punktyWykresuWielomianu[729].X = 7.29
	punktyWykresuWielomianu[729].Y = -7_272.584

	punktyWykresuWielomianu[730].X = 7.3
	punktyWykresuWielomianu[730].Y = -7_319.539

	punktyWykresuWielomianu[731].X = 7.31
	punktyWykresuWielomianu[731].Y = -7_366.642

	punktyWykresuWielomianu[732].X = 7.32
	punktyWykresuWielomianu[732].Y = -7_413.891

	punktyWykresuWielomianu[733].X = 7.33
	punktyWykresuWielomianu[733].Y = -7_461.287

	punktyWykresuWielomianu[734].X = 7.34
	punktyWykresuWielomianu[734].Y = -7_508.93

	punktyWykresuWielomianu[735].X = 7.35
	punktyWykresuWielomianu[735].Y = -7_556.518

	punktyWykresuWielomianu[736].X = 7.36
	punktyWykresuWielomianu[736].Y = -7_604.352

	punktyWykresuWielomianu[737].X = 7.37
	punktyWykresuWielomianu[737].Y = -7_652.332

	punktyWykresuWielomianu[738].X = 7.38
	punktyWykresuWielomianu[738].Y = -7_700.456

	punktyWykresuWielomianu[739].X = 7.39
	punktyWykresuWielomianu[739].Y = -7_748.726

	punktyWykresuWielomianu[740].X = 7.4
	punktyWykresuWielomianu[740].Y = -7_797.141

	punktyWykresuWielomianu[741].X = 7.41
	punktyWykresuWielomianu[741].Y = -7_845.699

	punktyWykresuWielomianu[742].X = 7.42
	punktyWykresuWielomianu[742].Y = -7_894.402

	punktyWykresuWielomianu[743].X = 7.43
	punktyWykresuWielomianu[743].Y = -7_943.249

	punktyWykresuWielomianu[744].X = 7.44
	punktyWykresuWielomianu[744].Y = -7_992.239

	punktyWykresuWielomianu[745].X = 7.45
	punktyWykresuWielomianu[745].Y = -8_041.372

	punktyWykresuWielomianu[746].X = 7.46
	punktyWykresuWielomianu[746].Y = -8_090.648

	punktyWykresuWielomianu[747].X = 7.47
	punktyWykresuWielomianu[747].Y = -8_140.067

	punktyWykresuWielomianu[748].X = 7.48
	punktyWykresuWielomianu[748].Y = -8_189.628

	punktyWykresuWielomianu[749].X = 7.49
	punktyWykresuWielomianu[749].Y = -8_239.331

	punktyWykresuWielomianu[750].X = 7.5
	punktyWykresuWielomianu[750].Y = -8_289.175

	punktyWykresuWielomianu[751].X = 7.51
	punktyWykresuWielomianu[751].Y = -8_339.161

	punktyWykresuWielomianu[752].X = 7.52
	punktyWykresuWielomianu[752].Y = -8_389.287

	punktyWykresuWielomianu[753].X = 7.53
	punktyWykresuWielomianu[753].Y = -8_439.554

	punktyWykresuWielomianu[754].X = 7.54
	punktyWykresuWielomianu[754].Y = -8_489.961

	punktyWykresuWielomianu[755].X = 7.55
	punktyWykresuWielomianu[755].Y = -8_540.509

	punktyWykresuWielomianu[756].X = 7.56
	punktyWykresuWielomianu[756].Y = -8_591.195

	punktyWykresuWielomianu[757].X = 7.57
	punktyWykresuWielomianu[757].Y = -8_642.021

	punktyWykresuWielomianu[758].X = 7.58
	punktyWykresuWielomianu[758].Y = -8_692.986

	punktyWykresuWielomianu[759].X = 7.59
	punktyWykresuWielomianu[759].Y = -8_744.089

	punktyWykresuWielomianu[760].X = 7.6
	punktyWykresuWielomianu[760].Y = -8_795.33

	punktyWykresuWielomianu[761].X = 7.61
	punktyWykresuWielomianu[761].Y = -8_846.709

	punktyWykresuWielomianu[762].X = 7.62
	punktyWykresuWielomianu[762].Y = -8_898.225

	punktyWykresuWielomianu[763].X = 7.63
	punktyWykresuWielomianu[763].Y = -8_949.877

	punktyWykresuWielomianu[764].X = 7.64
	punktyWykresuWielomianu[764].Y = -9_001.667

	punktyWykresuWielomianu[765].X = 7.65
	punktyWykresuWielomianu[765].Y = -9_053.593

	punktyWykresuWielomianu[766].X = 7.66
	punktyWykresuWielomianu[766].Y = -9_105.654

	punktyWykresuWielomianu[767].X = 7.67
	punktyWykresuWielomianu[767].Y = -9_157.85

	punktyWykresuWielomianu[768].X = 7.68
	punktyWykresuWielomianu[768].Y = -9_210.182

	punktyWykresuWielomianu[769].X = 7.69
	punktyWykresuWielomianu[769].Y = -9_262.648

	punktyWykresuWielomianu[770].X = 7.7
	punktyWykresuWielomianu[770].Y = -9_315.248

	punktyWykresuWielomianu[771].X = 7.71
	punktyWykresuWielomianu[771].Y = -9_367.982

	punktyWykresuWielomianu[772].X = 7.72
	punktyWykresuWielomianu[772].Y = -9_420.849

	punktyWykresuWielomianu[773].X = 7.73
	punktyWykresuWielomianu[773].Y = -9_473.849

	punktyWykresuWielomianu[774].X = 7.74
	punktyWykresuWielomianu[774].Y = -9_526.981

	punktyWykresuWielomianu[775].X = 7.75
	punktyWykresuWielomianu[775].Y = -9_580.245

	punktyWykresuWielomianu[776].X = 7.76
	punktyWykresuWielomianu[776].Y = -9_633.64

	punktyWykresuWielomianu[777].X = 7.77
	punktyWykresuWielomianu[777].Y = -9_687.167

	punktyWykresuWielomianu[778].X = 7.78
	punktyWykresuWielomianu[778].Y = -9_740.824

	punktyWykresuWielomianu[779].X = 7.79
	punktyWykresuWielomianu[779].Y = -9_794.611

	punktyWykresuWielomianu[780].X = 7.8
	punktyWykresuWielomianu[780].Y = -9_848.528

	punktyWykresuWielomianu[781].X = 7.81
	punktyWykresuWielomianu[781].Y = -9_902.574

	punktyWykresuWielomianu[782].X = 7.82
	punktyWykresuWielomianu[782].Y = -9_956.749

	punktyWykresuWielomianu[783].X = 7.83
	punktyWykresuWielomianu[783].Y = -10_011.052

	punktyWykresuWielomianu[784].X = 7.84
	punktyWykresuWielomianu[784].Y = -10_065.483

	punktyWykresuWielomianu[785].X = 7.85
	punktyWykresuWielomianu[785].Y = -10_120.041

	punktyWykresuWielomianu[786].X = 7.86
	punktyWykresuWielomianu[786].Y = -10_174.726

	punktyWykresuWielomianu[787].X = 7.87
	punktyWykresuWielomianu[787].Y = -10_229.537

	punktyWykresuWielomianu[788].X = 7.88
	punktyWykresuWielomianu[788].Y = -10_284.473

	punktyWykresuWielomianu[789].X = 7.89
	punktyWykresuWielomianu[789].Y = -10_339.535

	punktyWykresuWielomianu[790].X = 7.9
	punktyWykresuWielomianu[790].Y = -10_394.722

	punktyWykresuWielomianu[791].X = 7.91
	punktyWykresuWielomianu[791].Y = -10_450.033

	punktyWykresuWielomianu[792].X = 7.92
	punktyWykresuWielomianu[792].Y = -10_505.468

	punktyWykresuWielomianu[793].X = 7.93
	punktyWykresuWielomianu[793].Y = -10_561.026

	punktyWykresuWielomianu[794].X = 7.94
	punktyWykresuWielomianu[794].Y = -10_616.706

	punktyWykresuWielomianu[795].X = 7.95
	punktyWykresuWielomianu[795].Y = -10_672.509

	punktyWykresuWielomianu[796].X = 7.96
	punktyWykresuWielomianu[796].Y = -10_728.433

	punktyWykresuWielomianu[797].X = 7.97
	punktyWykresuWielomianu[797].Y = -10_784.478

	punktyWykresuWielomianu[798].X = 7.98
	punktyWykresuWielomianu[798].Y = -10_840.644

	punktyWykresuWielomianu[799].X = 7.99
	punktyWykresuWielomianu[799].Y = -10_896.93

	punktyWykresuWielomianu[800].X = 8.0
	punktyWykresuWielomianu[800].Y = -10_953.335

	punktyWykresuWielomianu[801].X = 8.01
	punktyWykresuWielomianu[801].Y = -11_009.858

	punktyWykresuWielomianu[802].X = 8.02
	punktyWykresuWielomianu[802].Y = -11_066.5

	punktyWykresuWielomianu[803].X = 8.03
	punktyWykresuWielomianu[803].Y = -11_123.26

	punktyWykresuWielomianu[804].X = 8.04
	punktyWykresuWielomianu[804].Y = -11_180.136

	punktyWykresuWielomianu[805].X = 8.05
	punktyWykresuWielomianu[805].Y = -11_237.129

	punktyWykresuWielomianu[806].X = 8.06
	punktyWykresuWielomianu[806].Y = -11_294.238

	punktyWykresuWielomianu[807].X = 8.07
	punktyWykresuWielomianu[807].Y = -11_351.463

	punktyWykresuWielomianu[808].X = 8.08
	punktyWykresuWielomianu[808].Y = -11_408.802

	punktyWykresuWielomianu[809].X = 8.09
	punktyWykresuWielomianu[809].Y = -11_466.255

	punktyWykresuWielomianu[810].X = 8.1
	punktyWykresuWielomianu[810].Y = -11_523.821

	punktyWykresuWielomianu[811].X = 8.11
	punktyWykresuWielomianu[811].Y = -11_581.5

	punktyWykresuWielomianu[812].X = 8.12
	punktyWykresuWielomianu[812].Y = -11_639.292

	punktyWykresuWielomianu[813].X = 8.13
	punktyWykresuWielomianu[813].Y = -11_697.195

	punktyWykresuWielomianu[814].X = 8.14
	punktyWykresuWielomianu[814].Y = -11_755.209

	punktyWykresuWielomianu[815].X = 8.15
	punktyWykresuWielomianu[815].Y = -11_813.333

	punktyWykresuWielomianu[816].X = 8.16
	punktyWykresuWielomianu[816].Y = -11_871.568

	punktyWykresuWielomianu[817].X = 8.17
	punktyWykresuWielomianu[817].Y = -11_929.911

	punktyWykresuWielomianu[818].X = 8.18
	punktyWykresuWielomianu[818].Y = -11_988.362

	punktyWykresuWielomianu[819].X = 8.19
	punktyWykresuWielomianu[819].Y = -12_046.921

	punktyWykresuWielomianu[820].X = 8.2
	punktyWykresuWielomianu[820].Y = -12_105.588

	punktyWykresuWielomianu[821].X = 8.21
	punktyWykresuWielomianu[821].Y = -12_164.36

	punktyWykresuWielomianu[822].X = 8.22
	punktyWykresuWielomianu[822].Y = -12_223.239

	punktyWykresuWielomianu[823].X = 8.23
	punktyWykresuWielomianu[823].Y = -12_282.222

	punktyWykresuWielomianu[824].X = 8.24
	punktyWykresuWielomianu[824].Y = -12_341.31

	punktyWykresuWielomianu[825].X = 8.25
	punktyWykresuWielomianu[825].Y = -12_400.501

	punktyWykresuWielomianu[826].X = 8.26
	punktyWykresuWielomianu[826].Y = -12_459.796

	punktyWykresuWielomianu[827].X = 8.27
	punktyWykresuWielomianu[827].Y = -12_519.193

	punktyWykresuWielomianu[828].X = 8.28
	punktyWykresuWielomianu[828].Y = -12_578.691

	punktyWykresuWielomianu[829].X = 8.29
	punktyWykresuWielomianu[829].Y = -12_638.29

	punktyWykresuWielomianu[830].X = 8.3
	punktyWykresuWielomianu[830].Y = -12_697.989

	punktyWykresuWielomianu[831].X = 8.31
	punktyWykresuWielomianu[831].Y = -12_757.788

	punktyWykresuWielomianu[832].X = 8.32
	punktyWykresuWielomianu[832].Y = -12_817.685

	punktyWykresuWielomianu[833].X = 8.33
	punktyWykresuWielomianu[833].Y = -12_877.681

	punktyWykresuWielomianu[834].X = 8.34
	punktyWykresuWielomianu[834].Y = -12_937.773

	punktyWykresuWielomianu[835].X = 8.35
	punktyWykresuWielomianu[835].Y = -12_997.962

	punktyWykresuWielomianu[836].X = 8.36
	punktyWykresuWielomianu[836].Y = -13_058.247

	punktyWykresuWielomianu[837].X = 8.37
	punktyWykresuWielomianu[837].Y = -13_118.627

	punktyWykresuWielomianu[838].X = 8.38
	punktyWykresuWielomianu[838].Y = -13_179.101

	punktyWykresuWielomianu[839].X = 8.39
	punktyWykresuWielomianu[839].Y = -13_239.668

	punktyWykresuWielomianu[840].X = 8.4
	punktyWykresuWielomianu[840].Y = -13_300.329

	punktyWykresuWielomianu[841].X = 8.41
	punktyWykresuWielomianu[841].Y = -13_361.081

	punktyWykresuWielomianu[842].X = 8.42
	punktyWykresuWielomianu[842].Y = -13_421.924

	punktyWykresuWielomianu[843].X = 8.43
	punktyWykresuWielomianu[843].Y = -13_482.857

	punktyWykresuWielomianu[844].X = 8.44
	punktyWykresuWielomianu[844].Y = -13_543.881

	punktyWykresuWielomianu[845].X = 8.45
	punktyWykresuWielomianu[845].Y = -13_604.992

	punktyWykresuWielomianu[846].X = 8.46
	punktyWykresuWielomianu[846].Y = -13_666.192

	punktyWykresuWielomianu[847].X = 8.47
	punktyWykresuWielomianu[847].Y = -13_727.479

	punktyWykresuWielomianu[848].X = 8.48
	punktyWykresuWielomianu[848].Y = -13_788.852

	punktyWykresuWielomianu[849].X = 8.49
	punktyWykresuWielomianu[849].Y = -13_850.311

	punktyWykresuWielomianu[850].X = 8.5
	punktyWykresuWielomianu[850].Y = -13_911.855

	punktyWykresuWielomianu[851].X = 8.51
	punktyWykresuWielomianu[851].Y = -13_973.482

	punktyWykresuWielomianu[852].X = 8.52
	punktyWykresuWielomianu[852].Y = -14_035.193

	punktyWykresuWielomianu[853].X = 8.53
	punktyWykresuWielomianu[853].Y = -14_096.985

	punktyWykresuWielomianu[854].X = 8.54
	punktyWykresuWielomianu[854].Y = -14_158.859

	punktyWykresuWielomianu[855].X = 8.55
	punktyWykresuWielomianu[855].Y = -14_220.813

	punktyWykresuWielomianu[856].X = 8.56
	punktyWykresuWielomianu[856].Y = -14_282.847

	punktyWykresuWielomianu[857].X = 8.57
	punktyWykresuWielomianu[857].Y = -14_344.96

	punktyWykresuWielomianu[858].X = 8.58
	punktyWykresuWielomianu[858].Y = -14_407.151

	punktyWykresuWielomianu[859].X = 8.59
	punktyWykresuWielomianu[859].Y = -14_469.419

	punktyWykresuWielomianu[860].X = 8.6
	punktyWykresuWielomianu[860].Y = -14_531.762

	punktyWykresuWielomianu[861].X = 8.61
	punktyWykresuWielomianu[861].Y = -14_594.181

	punktyWykresuWielomianu[862].X = 8.62
	punktyWykresuWielomianu[862].Y = -14_656.675

	punktyWykresuWielomianu[863].X = 8.63
	punktyWykresuWielomianu[863].Y = -14_719.242

	punktyWykresuWielomianu[864].X = 8.64
	punktyWykresuWielomianu[864].Y = -14_781.881

	punktyWykresuWielomianu[865].X = 8.65
	punktyWykresuWielomianu[865].Y = -14_844.592

	punktyWykresuWielomianu[866].X = 8.66
	punktyWykresuWielomianu[866].Y = -14_907.374

	punktyWykresuWielomianu[867].X = 8.67
	punktyWykresuWielomianu[867].Y = -14_970.226

	punktyWykresuWielomianu[868].X = 8.68
	punktyWykresuWielomianu[868].Y = -15_033.146

	punktyWykresuWielomianu[869].X = 8.69
	punktyWykresuWielomianu[869].Y = -15_096.134

	punktyWykresuWielomianu[870].X = 8.7
	punktyWykresuWielomianu[870].Y = -15_159.14

	punktyWykresuWielomianu[871].X = 8.71
	punktyWykresuWielomianu[871].Y = -15_222.311

	punktyWykresuWielomianu[872].X = 8.72
	punktyWykresuWielomianu[872].Y = -15_285.498

	punktyWykresuWielomianu[873].X = 8.73
	punktyWykresuWielomianu[873].Y = -15_348.749

	punktyWykresuWielomianu[874].X = 8.74
	punktyWykresuWielomianu[874].Y = -15_412.063

	punktyWykresuWielomianu[875].X = 8.75
	punktyWykresuWielomianu[875].Y = -15_475.439

	punktyWykresuWielomianu[876].X = 8.76
	punktyWykresuWielomianu[876].Y = -15_538.876

	punktyWykresuWielomianu[877].X = 8.77
	punktyWykresuWielomianu[877].Y = -15_602.374

	punktyWykresuWielomianu[878].X = 8.78
	punktyWykresuWielomianu[878].Y = -15_665.931

	punktyWykresuWielomianu[879].X = 8.79
	punktyWykresuWielomianu[879].Y = -15_729.546

	punktyWykresuWielomianu[880].X = 8.8
	punktyWykresuWielomianu[880].Y = -15_793.219

	punktyWykresuWielomianu[881].X = 8.81
	punktyWykresuWielomianu[881].Y = -15_856.948

	punktyWykresuWielomianu[882].X = 8.82
	punktyWykresuWielomianu[882].Y = -15_920.732

	punktyWykresuWielomianu[883].X = 8.83
	punktyWykresuWielomianu[883].Y = -15_984.571

	punktyWykresuWielomianu[884].X = 8.84
	punktyWykresuWielomianu[884].Y = -16_048.462

	punktyWykresuWielomianu[885].X = 8.85
	punktyWykresuWielomianu[885].Y = -16_112.406

	punktyWykresuWielomianu[886].X = 8.86
	punktyWykresuWielomianu[886].Y = -16_176.401

	punktyWykresuWielomianu[887].X = 8.87
	punktyWykresuWielomianu[887].Y = -16_240.447

	punktyWykresuWielomianu[888].X = 8.88
	punktyWykresuWielomianu[888].Y = -16_304.541

	punktyWykresuWielomianu[889].X = 8.89
	punktyWykresuWielomianu[889].Y = -16_368.683

	punktyWykresuWielomianu[890].X = 8.9
	punktyWykresuWielomianu[890].Y = -16_432.873

	punktyWykresuWielomianu[891].X = 8.91
	punktyWykresuWielomianu[891].Y = -16_497.108

	punktyWykresuWielomianu[892].X = 8.92
	punktyWykresuWielomianu[892].Y = -16_561.389

	punktyWykresuWielomianu[893].X = 8.93
	punktyWykresuWielomianu[893].Y = -16_625.713

	punktyWykresuWielomianu[894].X = 8.94
	punktyWykresuWielomianu[894].Y = -16_690.079

	punktyWykresuWielomianu[895].X = 8.95
	punktyWykresuWielomianu[895].Y = -16_754.488

	punktyWykresuWielomianu[896].X = 8.96
	punktyWykresuWielomianu[896].Y = -16_818.937

	punktyWykresuWielomianu[897].X = 8.97
	punktyWykresuWielomianu[897].Y = -16_883.425

	punktyWykresuWielomianu[898].X = 8.98
	punktyWykresuWielomianu[898].Y = -16_947.952

	punktyWykresuWielomianu[899].X = 8.99
	punktyWykresuWielomianu[899].Y = -17_012.516

	punktyWykresuWielomianu[900].X = 9.0
	punktyWykresuWielomianu[900].Y = -17_077.116

	punktyWykresuWielomianu[901].X = 9.01
	punktyWykresuWielomianu[901].Y = -17_141.75

	punktyWykresuWielomianu[902].X = 9.02
	punktyWykresuWielomianu[902].Y = -17_206.419

	punktyWykresuWielomianu[903].X = 9.03
	punktyWykresuWielomianu[903].Y = -17_271.121

	punktyWykresuWielomianu[904].X = 9.04
	punktyWykresuWielomianu[904].Y = -17_335.854

	punktyWykresuWielomianu[905].X = 9.05
	punktyWykresuWielomianu[905].Y = -17_400.617

	punktyWykresuWielomianu[906].X = 9.06
	punktyWykresuWielomianu[906].Y = -17_465.41

	punktyWykresuWielomianu[907].X = 9.07
	punktyWykresuWielomianu[907].Y = -17_530.231

	punktyWykresuWielomianu[908].X = 9.08
	punktyWykresuWielomianu[908].Y = -17_595.079

	punktyWykresuWielomianu[909].X = 9.09
	punktyWykresuWielomianu[909].Y = -17_659.952

	punktyWykresuWielomianu[910].X = 9.1
	punktyWykresuWielomianu[910].Y = -17_724.85

	punktyWykresuWielomianu[911].X = 9.11
	punktyWykresuWielomianu[911].Y = -17_789.771

	punktyWykresuWielomianu[912].X = 9.12
	punktyWykresuWielomianu[912].Y = -17_854.715

	punktyWykresuWielomianu[913].X = 9.13
	punktyWykresuWielomianu[913].Y = -17_919.68

	punktyWykresuWielomianu[914].X = 9.14
	punktyWykresuWielomianu[914].Y = -17_984.664

	punktyWykresuWielomianu[915].X = 9.15
	punktyWykresuWielomianu[915].Y = -18_049.667

	punktyWykresuWielomianu[916].X = 9.16
	punktyWykresuWielomianu[916].Y = -18_114.687

	punktyWykresuWielomianu[917].X = 9.17
	punktyWykresuWielomianu[917].Y = -18_179.724

	punktyWykresuWielomianu[918].X = 9.18
	punktyWykresuWielomianu[918].Y = -18_244.775

	punktyWykresuWielomianu[919].X = 9.19
	punktyWykresuWielomianu[919].Y = -18_309.84

	punktyWykresuWielomianu[920].X = 9.2
	punktyWykresuWielomianu[920].Y = -18_374.918

	punktyWykresuWielomianu[921].X = 9.21
	punktyWykresuWielomianu[921].Y = -18_440.006

	punktyWykresuWielomianu[922].X = 9.22
	punktyWykresuWielomianu[922].Y = -18_505.105

	punktyWykresuWielomianu[923].X = 9.23
	punktyWykresuWielomianu[923].Y = -18_570.212

	punktyWykresuWielomianu[924].X = 9.24
	punktyWykresuWielomianu[924].Y = -18_635.327

	punktyWykresuWielomianu[925].X = 9.25
	punktyWykresuWielomianu[925].Y = -18_700.447

	punktyWykresuWielomianu[926].X = 9.26
	punktyWykresuWielomianu[926].Y = -18_765.573

	punktyWykresuWielomianu[927].X = 9.27
	punktyWykresuWielomianu[927].Y = -18_830.702

	punktyWykresuWielomianu[928].X = 9.28
	punktyWykresuWielomianu[928].Y = -18_895.833

	punktyWykresuWielomianu[929].X = 9.29
	punktyWykresuWielomianu[929].Y = -18_960.965

	punktyWykresuWielomianu[930].X = 9.3
	punktyWykresuWielomianu[930].Y = -19_026.097

	punktyWykresuWielomianu[931].X = 9.31
	punktyWykresuWielomianu[931].Y = -19_091.227

	punktyWykresuWielomianu[932].X = 9.32
	punktyWykresuWielomianu[932].Y = -19_156.354

	punktyWykresuWielomianu[933].X = 9.33
	punktyWykresuWielomianu[933].Y = -19_221.4777

	punktyWykresuWielomianu[934].X = 9.34
	punktyWykresuWielomianu[934].Y = -19_286.594

	punktyWykresuWielomianu[935].X = 9.35
	punktyWykresuWielomianu[935].Y = -19_351.704

	punktyWykresuWielomianu[936].X = 9.36
	punktyWykresuWielomianu[936].Y = -19_416.805

	punktyWykresuWielomianu[937].X = 9.37
	punktyWykresuWielomianu[937].Y = -19_481.897

	punktyWykresuWielomianu[938].X = 9.38
	punktyWykresuWielomianu[938].Y = -19_546.978

	punktyWykresuWielomianu[939].X = 9.39
	punktyWykresuWielomianu[939].Y = -19_612.047

	punktyWykresuWielomianu[940].X = 9.4
	punktyWykresuWielomianu[940].Y = -19_677.101

	punktyWykresuWielomianu[941].X = 9.41
	punktyWykresuWielomianu[941].Y = -19_742.141

	punktyWykresuWielomianu[942].X = 9.42
	punktyWykresuWielomianu[942].Y = -19_807.163

	punktyWykresuWielomianu[943].X = 9.43
	punktyWykresuWielomianu[943].Y = -19_872.168

	punktyWykresuWielomianu[944].X = 9.44
	punktyWykresuWielomianu[944].Y = -19_937.154

	punktyWykresuWielomianu[945].X = 9.45
	punktyWykresuWielomianu[945].Y = -20_002.118

	punktyWykresuWielomianu[946].X = 9.46
	punktyWykresuWielomianu[946].Y = -20_067.061

	punktyWykresuWielomianu[947].X = 9.47
	punktyWykresuWielomianu[947].Y = -20_131.979

	punktyWykresuWielomianu[948].X = 9.48
	punktyWykresuWielomianu[948].Y = -20_196.873

	punktyWykresuWielomianu[949].X = 9.49
	punktyWykresuWielomianu[949].Y = -20_261.74

	punktyWykresuWielomianu[950].X = 9.5
	punktyWykresuWielomianu[950].Y = -20_326.58

	punktyWykresuWielomianu[951].X = 9.51
	punktyWykresuWielomianu[951].Y = -20_391.389

	punktyWykresuWielomianu[952].X = 9.52
	punktyWykresuWielomianu[952].Y = -20_456.168

	punktyWykresuWielomianu[953].X = 9.53
	punktyWykresuWielomianu[953].Y = -20_520.915

	punktyWykresuWielomianu[954].X = 9.54
	punktyWykresuWielomianu[954].Y = -20_585.628

	punktyWykresuWielomianu[955].X = 9.55
	punktyWykresuWielomianu[955].Y = -20_650.306

	punktyWykresuWielomianu[956].X = 9.56
	punktyWykresuWielomianu[956].Y = -20_714.947

	punktyWykresuWielomianu[957].X = 9.57
	punktyWykresuWielomianu[957].Y = -20_779.55

	punktyWykresuWielomianu[958].X = 9.58
	punktyWykresuWielomianu[958].Y = -20_844.113

	punktyWykresuWielomianu[959].X = 9.59
	punktyWykresuWielomianu[959].Y = -20_908.634

	punktyWykresuWielomianu[960].X = 9.6
	punktyWykresuWielomianu[960].Y = -20_973.114

	punktyWykresuWielomianu[961].X = 9.61
	punktyWykresuWielomianu[961].Y = -21_037.548

	punktyWykresuWielomianu[962].X = 9.62
	punktyWykresuWielomianu[962].Y = -21_101.937

	punktyWykresuWielomianu[963].X = 9.63
	punktyWykresuWielomianu[963].Y = -21_166.279

	punktyWykresuWielomianu[964].X = 9.64
	punktyWykresuWielomianu[964].Y = -21_230.572

	punktyWykresuWielomianu[965].X = 9.65
	punktyWykresuWielomianu[965].Y = -21_294.815

	punktyWykresuWielomianu[966].X = 9.66
	punktyWykresuWielomianu[966].Y = -21_359.005

	punktyWykresuWielomianu[967].X = 9.67
	punktyWykresuWielomianu[967].Y = -21_423.142

	punktyWykresuWielomianu[968].X = 9.68
	punktyWykresuWielomianu[968].Y = -21_487.225

	punktyWykresuWielomianu[969].X = 9.69
	punktyWykresuWielomianu[969].Y = -21_551.25

	punktyWykresuWielomianu[970].X = 9.7
	punktyWykresuWielomianu[970].Y = -21_615.217

	punktyWykresuWielomianu[971].X = 9.71
	punktyWykresuWielomianu[971].Y = -21_679.125

	punktyWykresuWielomianu[972].X = 9.72
	punktyWykresuWielomianu[972].Y = -21_742.971

	punktyWykresuWielomianu[973].X = 9.73
	punktyWykresuWielomianu[973].Y = -21_806.754

	punktyWykresuWielomianu[974].X = 9.74
	punktyWykresuWielomianu[974].Y = -21_870.472

	punktyWykresuWielomianu[975].X = 9.75
	punktyWykresuWielomianu[975].Y = -21_934.124

	punktyWykresuWielomianu[976].X = 9.76
	punktyWykresuWielomianu[976].Y = -21_997.709

	punktyWykresuWielomianu[977].X = 9.77
	punktyWykresuWielomianu[977].Y = -22_061.224

	punktyWykresuWielomianu[978].X = 9.78
	punktyWykresuWielomianu[978].Y = -22_124.668

	punktyWykresuWielomianu[979].X = 9.79
	punktyWykresuWielomianu[979].Y = -22_188.039

	punktyWykresuWielomianu[980].X = 9.8
	punktyWykresuWielomianu[980].Y = -22_251.335

	punktyWykresuWielomianu[981].X = 9.81
	punktyWykresuWielomianu[981].Y = -22_314.556

	punktyWykresuWielomianu[982].X = 9.82
	punktyWykresuWielomianu[982].Y = -22_377.699

	punktyWykresuWielomianu[983].X = 9.83
	punktyWykresuWielomianu[983].Y = -22_440.763

	punktyWykresuWielomianu[984].X = 9.84
	punktyWykresuWielomianu[984].Y = -22_503.745

	punktyWykresuWielomianu[985].X = 9.85
	punktyWykresuWielomianu[985].Y = -22_566.645

	punktyWykresuWielomianu[986].X = 9.86
	punktyWykresuWielomianu[986].Y = -22_629.46

	punktyWykresuWielomianu[987].X = 9.87
	punktyWykresuWielomianu[987].Y = -22_692.189

	punktyWykresuWielomianu[988].X = 9.88
	punktyWykresuWielomianu[988].Y = -22_754.831

	punktyWykresuWielomianu[989].X = 9.89
	punktyWykresuWielomianu[989].Y = -22_817.382

	punktyWykresuWielomianu[990].X = 9.9
	punktyWykresuWielomianu[990].Y = -22_879.843

	punktyWykresuWielomianu[991].X = 9.91
	punktyWykresuWielomianu[991].Y = -22_942.21

	punktyWykresuWielomianu[992].X = 9.92
	punktyWykresuWielomianu[992].Y = -23_004.483

	punktyWykresuWielomianu[993].X = 9.93
	punktyWykresuWielomianu[993].Y = -23_066.659

	punktyWykresuWielomianu[994].X = 9.94
	punktyWykresuWielomianu[994].Y = -23_128.737

	punktyWykresuWielomianu[995].X = 9.95
	punktyWykresuWielomianu[995].Y = -23_190.715

	punktyWykresuWielomianu[996].X = 9.96
	punktyWykresuWielomianu[996].Y = -23_252.591

	punktyWykresuWielomianu[997].X = 9.97
	punktyWykresuWielomianu[997].Y = -23_314.364

	punktyWykresuWielomianu[998].X = 9.98
	punktyWykresuWielomianu[998].Y = -23_376.031

	punktyWykresuWielomianu[999].X = 9.99
	punktyWykresuWielomianu[999].Y = -23_437.591

	punktyWykresuWielomianu[1_000].X = 10.0
	punktyWykresuWielomianu[1_000].Y = -23_499.043










	wykresWielomianu := plot.New()

	wykresWielomianu.Title.Text = "Wykres funkcji f(x) = x^5 - 19.222x^4 + 69.903x^3 - 10.34x^2 - 15.722x + 9.177"

	wykresWielomianu.X.Label.Text = "x"
	wykresWielomianu.Y.Label.Text = "y"

	liniaWykresu, err := plotter.NewLine(punktyWykresuWielomianu)

	if err != nil {
		panic(err)
	}

	liniaWykresu.LineStyle.Width = vg.Points(0.1)
	liniaWykresu.Color = color.RGBA{R: 200, G: 100, B: 100}

	wykresWielomianu.Add(liniaWykresu)
	wykresWielomianu.Legend.Add("f(x)", liniaWykresu)

	if err := wykresWielomianu.Save(10*vg.Inch, 10*vg.Inch,
		"Wykres-wielomianu-04.png"); err != nil {

		panic(err)
	}
}
