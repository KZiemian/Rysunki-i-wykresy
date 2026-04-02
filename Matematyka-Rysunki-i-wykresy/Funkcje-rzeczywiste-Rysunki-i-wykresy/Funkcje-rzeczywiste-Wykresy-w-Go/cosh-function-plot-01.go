package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function cosh(x).

	pointsOfFunctionPlot := make(plotter.XYs, 1_001)

	pointsOfFunctionPlot[0].X = -5.0
	pointsOfFunctionPlot[0].Y = 74.209

	pointsOfFunctionPlot[1].X = -4.99
	pointsOfFunctionPlot[1].Y = 73.471

	pointsOfFunctionPlot[2].X = -4.98
	pointsOfFunctionPlot[2].Y = 72.74

	pointsOfFunctionPlot[3].X = -4.97
	pointsOfFunctionPlot[3].Y = 72.016

	pointsOfFunctionPlot[4].X = -4.96
	pointsOfFunctionPlot[4].Y = 71.3

	pointsOfFunctionPlot[5].X = -4.95
	pointsOfFunctionPlot[5].Y = 70.591

	pointsOfFunctionPlot[6].X = -4.94
	pointsOfFunctionPlot[6].Y = 69.888

	pointsOfFunctionPlot[7].X = -4.93
	pointsOfFunctionPlot[7].Y = 69.193

	pointsOfFunctionPlot[8].X = -4.92
	pointsOfFunctionPlot[8].Y = 68.504

	pointsOfFunctionPlot[9].X = -4.91
	pointsOfFunctionPlot[9].Y = 67.823

	pointsOfFunctionPlot[10].X = -4.90
	pointsOfFunctionPlot[10].Y = 67.148

	pointsOfFunctionPlot[11].X = -4.89
	pointsOfFunctionPlot[11].Y = 66.48

	pointsOfFunctionPlot[12].X = -4.88
	pointsOfFunctionPlot[12].Y = 65.819

	pointsOfFunctionPlot[13].X = -4.87
	pointsOfFunctionPlot[13].Y = 65.164

	pointsOfFunctionPlot[14].X = -4.86
	pointsOfFunctionPlot[14].Y = 64.515

	pointsOfFunctionPlot[15].X = -4.85
	pointsOfFunctionPlot[15].Y = 63.874

	pointsOfFunctionPlot[16].X = -4.84
	pointsOfFunctionPlot[16].Y = 63.238

	pointsOfFunctionPlot[17].X = -4.83
	pointsOfFunctionPlot[17].Y = 62.609

	pointsOfFunctionPlot[18].X = -4.82
	pointsOfFunctionPlot[18].Y = 61.986

	pointsOfFunctionPlot[19].X = -4.81
	pointsOfFunctionPlot[19].Y = 61.369

	pointsOfFunctionPlot[20].X = -4.80
	pointsOfFunctionPlot[20].Y = 60.759

	pointsOfFunctionPlot[21].X = -4.79
	pointsOfFunctionPlot[21].Y = 60.154

	pointsOfFunctionPlot[22].X = -4.78
	pointsOfFunctionPlot[22].Y = 59.556

	pointsOfFunctionPlot[23].X = -4.77
	pointsOfFunctionPlot[23].Y = 58.963

	pointsOfFunctionPlot[24].X = -4.76
	pointsOfFunctionPlot[24].Y = 58.377

	pointsOfFunctionPlot[25].X = -4.75
	pointsOfFunctionPlot[25].Y = 57.796

	pointsOfFunctionPlot[26].X = -4.74
	pointsOfFunctionPlot[26].Y = 57.221

	pointsOfFunctionPlot[27].X = -4.73
	pointsOfFunctionPlot[27].Y = 56.652

	pointsOfFunctionPlot[28].X = -4.72
	pointsOfFunctionPlot[28].Y = 56.088

	pointsOfFunctionPlot[29].X = -4.71
	pointsOfFunctionPlot[29].Y = 55.53

	pointsOfFunctionPlot[30].X = -4.70
	pointsOfFunctionPlot[30].Y = 54.978

	pointsOfFunctionPlot[31].X = -4.69
	pointsOfFunctionPlot[31].Y = 54.431

	pointsOfFunctionPlot[32].X = -4.68
	pointsOfFunctionPlot[32].Y = 53.889

	pointsOfFunctionPlot[33].X = -4.67
	pointsOfFunctionPlot[33].Y = 53.353

	pointsOfFunctionPlot[34].X = -4.66
	pointsOfFunctionPlot[34].Y = 52.822

	pointsOfFunctionPlot[35].X = -4.65
	pointsOfFunctionPlot[35].Y = 52.297

	pointsOfFunctionPlot[36].X = -4.64
	pointsOfFunctionPlot[36].Y = 51.777

	pointsOfFunctionPlot[37].X = -4.63
	pointsOfFunctionPlot[37].Y = 51.261

	pointsOfFunctionPlot[38].X = -4.62
	pointsOfFunctionPlot[38].Y = 50.751

	pointsOfFunctionPlot[39].X = -4.61
	pointsOfFunctionPlot[39].Y = 50.247

	pointsOfFunctionPlot[40].X = -4.60
	pointsOfFunctionPlot[40].Y = 49.747

	pointsOfFunctionPlot[41].X = -4.59
	pointsOfFunctionPlot[41].Y = 49.252

	pointsOfFunctionPlot[42].X = -4.58
	pointsOfFunctionPlot[42].Y = 48.762

	pointsOfFunctionPlot[43].X = -4.57
	pointsOfFunctionPlot[43].Y = 48.277

	pointsOfFunctionPlot[44].X = -4.56
	pointsOfFunctionPlot[44].Y = 47.796

	pointsOfFunctionPlot[45].X = -4.55
	pointsOfFunctionPlot[45].Y = 47.321

	pointsOfFunctionPlot[46].X = -4.54
	pointsOfFunctionPlot[46].Y = 46.85

	pointsOfFunctionPlot[47].X = -4.53
	pointsOfFunctionPlot[47].Y = 46.384

	pointsOfFunctionPlot[48].X = -4.52
	pointsOfFunctionPlot[48].Y = 45.923

	pointsOfFunctionPlot[49].X = -4.51
	pointsOfFunctionPlot[49].Y = 45.466

	pointsOfFunctionPlot[50].X = -4.50
	pointsOfFunctionPlot[50].Y = 45.014

	pointsOfFunctionPlot[51].X = -4.49
	pointsOfFunctionPlot[51].Y = 44.566

	pointsOfFunctionPlot[52].X = -4.48
	pointsOfFunctionPlot[52].Y = 44.123

	pointsOfFunctionPlot[53].X = -4.47
	pointsOfFunctionPlot[53].Y = 43.684

	pointsOfFunctionPlot[54].X = -4.46
	pointsOfFunctionPlot[54].Y = 43.249

	pointsOfFunctionPlot[55].X = -4.45
	pointsOfFunctionPlot[55].Y = 42.819

	pointsOfFunctionPlot[56].X = -4.44
	pointsOfFunctionPlot[56].Y = 42.393

	pointsOfFunctionPlot[57].X = -4.43
	pointsOfFunctionPlot[57].Y = 41.971

	pointsOfFunctionPlot[58].X = -4.42
	pointsOfFunctionPlot[58].Y = 41.554

	pointsOfFunctionPlot[59].X = -4.41
	pointsOfFunctionPlot[59].Y = 41.14

	pointsOfFunctionPlot[60].X = -4.40
	pointsOfFunctionPlot[60].Y = 40.731

	pointsOfFunctionPlot[61].X = -4.39
	pointsOfFunctionPlot[61].Y = 40.326

	pointsOfFunctionPlot[62].X = -4.38
	pointsOfFunctionPlot[62].Y = 39.925

	pointsOfFunctionPlot[63].X = -4.37
	pointsOfFunctionPlot[63].Y = 39.528

	pointsOfFunctionPlot[64].X = -4.36
	pointsOfFunctionPlot[64].Y = 39.134

	pointsOfFunctionPlot[65].X = -4.35
	pointsOfFunctionPlot[65].Y = 38.745

	pointsOfFunctionPlot[66].X = -4.34
	pointsOfFunctionPlot[66].Y = 38.36

	pointsOfFunctionPlot[67].X = -4.33
	pointsOfFunctionPlot[67].Y = 37.978

	pointsOfFunctionPlot[68].X = -4.32
	pointsOfFunctionPlot[68].Y = 37.6

	pointsOfFunctionPlot[69].X = -4.31
	pointsOfFunctionPlot[69].Y = 37.226

	pointsOfFunctionPlot[70].X = -4.30
	pointsOfFunctionPlot[70].Y = 36.856

	pointsOfFunctionPlot[71].X = -4.29
	pointsOfFunctionPlot[71].Y = 36.49

	pointsOfFunctionPlot[72].X = -4.28
	pointsOfFunctionPlot[72].Y = 36.127

	pointsOfFunctionPlot[73].X = -4.27
	pointsOfFunctionPlot[73].Y = 35.767

	pointsOfFunctionPlot[74].X = -4.26
	pointsOfFunctionPlot[74].Y = 35.412

	pointsOfFunctionPlot[75].X = -4.25
	pointsOfFunctionPlot[75].Y = 35.059

	pointsOfFunctionPlot[76].X = -4.24
	pointsOfFunctionPlot[76].Y = 34.711

	pointsOfFunctionPlot[77].X = -4.23
	pointsOfFunctionPlot[77].Y = 34.365

	pointsOfFunctionPlot[78].X = -4.22
	pointsOfFunctionPlot[78].Y = 34.024

	pointsOfFunctionPlot[79].X = -4.21
	pointsOfFunctionPlot[79].Y = 33.685

	pointsOfFunctionPlot[80].X = -4.20
	pointsOfFunctionPlot[80].Y = 33.35

	pointsOfFunctionPlot[81].X = -4.19
	pointsOfFunctionPlot[81].Y = 33.018

	pointsOfFunctionPlot[82].X = -4.18
	pointsOfFunctionPlot[82].Y = 32.69

	pointsOfFunctionPlot[83].X = -4.17
	pointsOfFunctionPlot[83].Y = 32.365

	pointsOfFunctionPlot[84].X = -4.16
	pointsOfFunctionPlot[84].Y = 32.043

	pointsOfFunctionPlot[85].X = -4.15
	pointsOfFunctionPlot[85].Y = 31.724

	pointsOfFunctionPlot[86].X = -4.14
	pointsOfFunctionPlot[86].Y = 31.409

	pointsOfFunctionPlot[87].X = -4.13
	pointsOfFunctionPlot[87].Y = 31.097

	pointsOfFunctionPlot[88].X = -4.12
	pointsOfFunctionPlot[88].Y = 30.787

	pointsOfFunctionPlot[89].X = -4.11
	pointsOfFunctionPlot[89].Y = 30.481

	pointsOfFunctionPlot[90].X = -4.10
	pointsOfFunctionPlot[90].Y = 30.178

	pointsOfFunctionPlot[91].X = -4.09
	pointsOfFunctionPlot[91].Y = 29.878

	pointsOfFunctionPlot[92].X = -4.08
	pointsOfFunctionPlot[92].Y = 29.581

	pointsOfFunctionPlot[93].X = -4.07
	pointsOfFunctionPlot[93].Y = 29.287

	pointsOfFunctionPlot[94].X = -4.06
	pointsOfFunctionPlot[94].Y = 28.995

	pointsOfFunctionPlot[95].X = -4.05
	pointsOfFunctionPlot[95].Y = 28.707

	pointsOfFunctionPlot[96].X = -4.04
	pointsOfFunctionPlot[96].Y = 28.421

	pointsOfFunctionPlot[97].X = -4.03
	pointsOfFunctionPlot[97].Y = 28.139

	pointsOfFunctionPlot[98].X = -4.02
	pointsOfFunctionPlot[98].Y = 27.859

	pointsOfFunctionPlot[99].X = -4.01
	pointsOfFunctionPlot[99].Y = 27.582

	pointsOfFunctionPlot[100].X = -4.0
	pointsOfFunctionPlot[100].Y = 27.308

	pointsOfFunctionPlot[101].X = -3.99
	pointsOfFunctionPlot[101].Y = 27.036

	pointsOfFunctionPlot[102].X = -3.98
	pointsOfFunctionPlot[102].Y = 26.767

	pointsOfFunctionPlot[103].X = -3.97
	pointsOfFunctionPlot[103].Y = 26.501

	pointsOfFunctionPlot[104].X = -3.96
	pointsOfFunctionPlot[104].Y = 26.238

	pointsOfFunctionPlot[105].X = -3.95
	pointsOfFunctionPlot[105].Y = 25.977

	pointsOfFunctionPlot[106].X = -3.94
	pointsOfFunctionPlot[106].Y = 25.719

	pointsOfFunctionPlot[107].X = -3.93
	pointsOfFunctionPlot[107].Y = 25.463

	pointsOfFunctionPlot[108].X = -3.92
	pointsOfFunctionPlot[108].Y = 25.21

	pointsOfFunctionPlot[109].X = -3.91
	pointsOfFunctionPlot[109].Y = 24.959

	pointsOfFunctionPlot[110].X = -3.90
	pointsOfFunctionPlot[110].Y = 24.711

	pointsOfFunctionPlot[111].X = -3.89
	pointsOfFunctionPlot[111].Y = 24.465

	pointsOfFunctionPlot[112].X = -3.88
	pointsOfFunctionPlot[112].Y = 24.222

	pointsOfFunctionPlot[113].X = -3.87
	pointsOfFunctionPlot[113].Y = 23.981

	pointsOfFunctionPlot[114].X = -3.86
	pointsOfFunctionPlot[114].Y = 23.722

	pointsOfFunctionPlot[115].X = -3.85
	pointsOfFunctionPlot[115].Y = 23.507

	pointsOfFunctionPlot[116].X = -3.84
	pointsOfFunctionPlot[116].Y = 23.273

	pointsOfFunctionPlot[117].X = -3.83
	pointsOfFunctionPlot[117].Y = 23.042

	pointsOfFunctionPlot[118].X = -3.82
	pointsOfFunctionPlot[118].Y = 22.813

	pointsOfFunctionPlot[119].X = -3.81
	pointsOfFunctionPlot[119].Y = 22.586

	pointsOfFunctionPlot[120].X = -3.80
	pointsOfFunctionPlot[120].Y = 22.361

	pointsOfFunctionPlot[121].X = -3.79
	pointsOfFunctionPlot[121].Y = 22.139

	pointsOfFunctionPlot[122].X = -3.78
	pointsOfFunctionPlot[122].Y = 21.919

	pointsOfFunctionPlot[123].X = -3.77
	pointsOfFunctionPlot[123].Y = 21.701

	pointsOfFunctionPlot[124].X = -3.76
	pointsOfFunctionPlot[124].Y = 21.485

	pointsOfFunctionPlot[125].X = -3.75
	pointsOfFunctionPlot[125].Y = 21.272

	pointsOfFunctionPlot[126].X = -3.74
	pointsOfFunctionPlot[126].Y = 21.06

	pointsOfFunctionPlot[127].X = -3.73
	pointsOfFunctionPlot[127].Y = 20.851

	pointsOfFunctionPlot[128].X = -3.72
	pointsOfFunctionPlot[128].Y = 20.644

	pointsOfFunctionPlot[129].X = -3.71
	pointsOfFunctionPlot[129].Y = 20.439

	pointsOfFunctionPlot[130].X = -3.70
	pointsOfFunctionPlot[130].Y = 20.236

	pointsOfFunctionPlot[131].X = -3.69
	pointsOfFunctionPlot[131].Y = 20.034

	pointsOfFunctionPlot[132].X = -3.68
	pointsOfFunctionPlot[132].Y = 19.835

	pointsOfFunctionPlot[133].X = -3.67
	pointsOfFunctionPlot[133].Y = 19.638

	pointsOfFunctionPlot[134].X = -3.66
	pointsOfFunctionPlot[134].Y = 19.443

	pointsOfFunctionPlot[135].X = -3.65
	pointsOfFunctionPlot[135].Y = 19.25

	pointsOfFunctionPlot[136].X = -3.64
	pointsOfFunctionPlot[136].Y = 19.059

	pointsOfFunctionPlot[137].X = -3.63
	pointsOfFunctionPlot[137].Y = 18.869

	pointsOfFunctionPlot[138].X = -3.62
	pointsOfFunctionPlot[138].Y = 18.682

	pointsOfFunctionPlot[139].X = -3.61
	pointsOfFunctionPlot[139].Y = 18.496

	pointsOfFunctionPlot[140].X = -3.60
	pointsOfFunctionPlot[140].Y = 18.312

	pointsOfFunctionPlot[141].X = -3.59
	pointsOfFunctionPlot[141].Y = 18.13

	pointsOfFunctionPlot[142].X = -3.58
	pointsOfFunctionPlot[142].Y = 17.95

	pointsOfFunctionPlot[143].X = -3.57
	pointsOfFunctionPlot[143].Y = 17.772

	pointsOfFunctionPlot[144].X = -3.56
	pointsOfFunctionPlot[144].Y = 17.595

	pointsOfFunctionPlot[145].X = -3.55
	pointsOfFunctionPlot[145].Y = 17.421

	pointsOfFunctionPlot[146].X = -3.54
	pointsOfFunctionPlot[146].Y = 17.247

	pointsOfFunctionPlot[147].X = -3.53
	pointsOfFunctionPlot[147].Y = 17.076

	pointsOfFunctionPlot[148].X = -3.52
	pointsOfFunctionPlot[148].Y = 16.907

	pointsOfFunctionPlot[149].X = -3.51
	pointsOfFunctionPlot[149].Y = 16.739

	pointsOfFunctionPlot[150].X = -3.50
	pointsOfFunctionPlot[150].Y = 16.572

	pointsOfFunctionPlot[151].X = -3.49
	pointsOfFunctionPlot[151].Y = 16.408

	pointsOfFunctionPlot[152].X = -3.48
	pointsOfFunctionPlot[152].Y = 16.245

	pointsOfFunctionPlot[153].X = -3.47
	pointsOfFunctionPlot[153].Y = 16.083

	pointsOfFunctionPlot[154].X = -3.46
	pointsOfFunctionPlot[154].Y = 15.924

	pointsOfFunctionPlot[155].X = -3.45
	pointsOfFunctionPlot[155].Y = 15.766

	pointsOfFunctionPlot[156].X = -3.44
	pointsOfFunctionPlot[156].Y = 15.609

	pointsOfFunctionPlot[157].X = -3.43
	pointsOfFunctionPlot[157].Y = 15.454

	pointsOfFunctionPlot[158].X = -3.42
	pointsOfFunctionPlot[158].Y = 15.301

	pointsOfFunctionPlot[159].X = -3.41
	pointsOfFunctionPlot[159].Y = 15.149

	pointsOfFunctionPlot[160].X = -3.40
	pointsOfFunctionPlot[160].Y = 14.998

	pointsOfFunctionPlot[161].X = -3.39
	pointsOfFunctionPlot[161].Y = 14.849

	pointsOfFunctionPlot[162].X = -3.38
	pointsOfFunctionPlot[162].Y = 14.702

	pointsOfFunctionPlot[163].X = -3.37
	pointsOfFunctionPlot[163].Y = 14.556

	pointsOfFunctionPlot[164].X = -3.36
	pointsOfFunctionPlot[164].Y = 14.411

	pointsOfFunctionPlot[165].X = -3.35
	pointsOfFunctionPlot[165].Y = 14.268

	pointsOfFunctionPlot[166].X = -3.34
	pointsOfFunctionPlot[166].Y = 14.127

	pointsOfFunctionPlot[167].X = -3.33
	pointsOfFunctionPlot[167].Y = 13.987

	pointsOfFunctionPlot[168].X = -3.32
	pointsOfFunctionPlot[168].Y = 13.848

	pointsOfFunctionPlot[169].X = -3.31
	pointsOfFunctionPlot[169].Y = 13.71

	pointsOfFunctionPlot[170].X = -3.30
	pointsOfFunctionPlot[170].Y = 13.574

	pointsOfFunctionPlot[171].X = -3.29
	pointsOfFunctionPlot[171].Y = 13.44

	pointsOfFunctionPlot[172].X = -3.28
	pointsOfFunctionPlot[172].Y = 13.306

	pointsOfFunctionPlot[173].X = -3.27
	pointsOfFunctionPlot[173].Y = 13.174

	pointsOfFunctionPlot[174].X = -3.26
	pointsOfFunctionPlot[174].Y = 13.043

	pointsOfFunctionPlot[175].X = -3.25
	pointsOfFunctionPlot[175].Y = 12.914

	pointsOfFunctionPlot[176].X = -3.24
	pointsOfFunctionPlot[176].Y = 12.786

	pointsOfFunctionPlot[177].X = -3.23
	pointsOfFunctionPlot[177].Y = 12.659

	pointsOfFunctionPlot[178].X = -3.22
	pointsOfFunctionPlot[178].Y = 12.534

	pointsOfFunctionPlot[179].X = -3.21
	pointsOfFunctionPlot[179].Y = 12.409

	pointsOfFunctionPlot[180].X = -3.20
	pointsOfFunctionPlot[180].Y = 12.286

	pointsOfFunctionPlot[181].X = -3.19
	pointsOfFunctionPlot[181].Y = 12.164

	pointsOfFunctionPlot[182].X = -3.18
	pointsOfFunctionPlot[182].Y = 12.044

	pointsOfFunctionPlot[183].X = -3.17
	pointsOfFunctionPlot[183].Y = 11.924

	pointsOfFunctionPlot[184].X = -3.16
	pointsOfFunctionPlot[184].Y = 11.806

	pointsOfFunctionPlot[185].X = -3.15
	pointsOfFunctionPlot[185].Y = 11.689

	pointsOfFunctionPlot[186].X = -3.14
	pointsOfFunctionPlot[186].Y = 11.573

	pointsOfFunctionPlot[187].X = -3.13
	pointsOfFunctionPlot[187].Y = 11.458

	pointsOfFunctionPlot[188].X = -3.12
	pointsOfFunctionPlot[188].Y = 11.345

	pointsOfFunctionPlot[189].X = -3.11
	pointsOfFunctionPlot[189].Y = 11.232

	pointsOfFunctionPlot[190].X = -3.10
	pointsOfFunctionPlot[190].Y = 11.121

	pointsOfFunctionPlot[191].X = -3.09
	pointsOfFunctionPlot[191].Y = 11.011

	pointsOfFunctionPlot[192].X = -3.08
	pointsOfFunctionPlot[192].Y = 10.902

	pointsOfFunctionPlot[193].X = -3.07
	pointsOfFunctionPlot[193].Y = 10.794

	pointsOfFunctionPlot[194].X = -3.06
	pointsOfFunctionPlot[194].Y = 10.687

	pointsOfFunctionPlot[195].X = -3.05
	pointsOfFunctionPlot[195].Y = 10.581

	pointsOfFunctionPlot[196].X = -3.04
	pointsOfFunctionPlot[196].Y = 10.476

	pointsOfFunctionPlot[197].X = -3.03
	pointsOfFunctionPlot[197].Y = 10.372

	pointsOfFunctionPlot[198].X = -3.02
	pointsOfFunctionPlot[198].Y = 10.27

	pointsOfFunctionPlot[199].X = -3.01
	pointsOfFunctionPlot[199].Y = 10.168

	pointsOfFunctionPlot[200].X = -3.0
	pointsOfFunctionPlot[200].Y = 10.067

	pointsOfFunctionPlot[201].X = -2.99
	pointsOfFunctionPlot[201].Y = 9.967

	pointsOfFunctionPlot[202].X = -2.98
	pointsOfFunctionPlot[202].Y = 9.869

	pointsOfFunctionPlot[203].X = -2.97
	pointsOfFunctionPlot[203].Y = 9.771

	pointsOfFunctionPlot[204].X = -2.96
	pointsOfFunctionPlot[204].Y = 9.674

	pointsOfFunctionPlot[205].X = -2.95
	pointsOfFunctionPlot[205].Y = 9.579

	pointsOfFunctionPlot[206].X = -2.94
	pointsOfFunctionPlot[206].Y = 9.484

	pointsOfFunctionPlot[207].X = -2.93
	pointsOfFunctionPlot[207].Y = 9.39

	pointsOfFunctionPlot[208].X = -2.92
	pointsOfFunctionPlot[208].Y = 9.297

	pointsOfFunctionPlot[209].X = -2.91
	pointsOfFunctionPlot[209].Y = 9.205

	pointsOfFunctionPlot[210].X = -2.90
	pointsOfFunctionPlot[210].Y = 9.144

	pointsOfFunctionPlot[211].X = -2.89
	pointsOfFunctionPlot[211].Y = 9.024

	pointsOfFunctionPlot[212].X = -2.88
	pointsOfFunctionPlot[212].Y = 8.935

	pointsOfFunctionPlot[213].X = -2.87
	pointsOfFunctionPlot[213].Y = 8.846

	pointsOfFunctionPlot[214].X = -2.86
	pointsOfFunctionPlot[214].Y = 8.759

	pointsOfFunctionPlot[215].X = -2.85
	pointsOfFunctionPlot[215].Y = 8.672

	pointsOfFunctionPlot[216].X = -2.84
	pointsOfFunctionPlot[216].Y = 8.587

	pointsOfFunctionPlot[217].X = -2.83
	pointsOfFunctionPlot[217].Y = 8.502

	pointsOfFunctionPlot[218].X = -2.82
	pointsOfFunctionPlot[218].Y = 8.418

	pointsOfFunctionPlot[219].X = -2.81
	pointsOfFunctionPlot[219].Y = 8.335

	pointsOfFunctionPlot[220].X = -2.80
	pointsOfFunctionPlot[220].Y = 8.252

	pointsOfFunctionPlot[221].X = -2.79
	pointsOfFunctionPlot[221].Y = 8.171

	pointsOfFunctionPlot[222].X = -2.78
	pointsOfFunctionPlot[222].Y = 8.09

	pointsOfFunctionPlot[223].X = -2.77
	pointsOfFunctionPlot[223].Y = 8.01

	pointsOfFunctionPlot[224].X = -2.76
	pointsOfFunctionPlot[224].Y = 7.931

	pointsOfFunctionPlot[225].X = -2.75
	pointsOfFunctionPlot[225].Y = 7.853

	pointsOfFunctionPlot[226].X = -2.74
	pointsOfFunctionPlot[226].Y = 7.775

	pointsOfFunctionPlot[227].X = -2.73
	pointsOfFunctionPlot[227].Y = 7.699

	pointsOfFunctionPlot[228].X = -2.72
	pointsOfFunctionPlot[228].Y = 7.623

	pointsOfFunctionPlot[229].X = -2.71
	pointsOfFunctionPlot[229].Y = 7.547

	pointsOfFunctionPlot[230].X = -2.70
	pointsOfFunctionPlot[230].Y = 7.473

	pointsOfFunctionPlot[231].X = -2.69
	pointsOfFunctionPlot[231].Y = 7.399

	pointsOfFunctionPlot[232].X = -2.68
	pointsOfFunctionPlot[232].Y = 7.326

	pointsOfFunctionPlot[233].X = -2.67
	pointsOfFunctionPlot[233].Y = 7.254

	pointsOfFunctionPlot[234].X = -2.66
	pointsOfFunctionPlot[234].Y = 7.183

	pointsOfFunctionPlot[235].X = -2.65
	pointsOfFunctionPlot[235].Y = 7.112

	pointsOfFunctionPlot[236].X = -2.64
	pointsOfFunctionPlot[236].Y = 7.042

	pointsOfFunctionPlot[237].X = -2.63
	pointsOfFunctionPlot[237].Y = 6.972

	pointsOfFunctionPlot[238].X = -2.62
	pointsOfFunctionPlot[238].Y = 6.904

	pointsOfFunctionPlot[239].X = -2.61
	pointsOfFunctionPlot[239].Y = 6.836

	pointsOfFunctionPlot[240].X = -2.60
	pointsOfFunctionPlot[240].Y = 6.769

	pointsOfFunctionPlot[241].X = -2.59
	pointsOfFunctionPlot[241].Y = 6.702

	pointsOfFunctionPlot[242].X = -2.58
	pointsOfFunctionPlot[242].Y = 6.636

	pointsOfFunctionPlot[243].X = -2.57
	pointsOfFunctionPlot[243].Y = 6.571

	pointsOfFunctionPlot[244].X = -2.56
	pointsOfFunctionPlot[244].Y = 6.506

	pointsOfFunctionPlot[245].X = -2.55
	pointsOfFunctionPlot[245].Y = 6.442

	pointsOfFunctionPlot[246].X = -2.54
	pointsOfFunctionPlot[246].Y = 6.379

	pointsOfFunctionPlot[247].X = -2.53
	pointsOfFunctionPlot[247].Y = 6.316

	pointsOfFunctionPlot[248].X = -2.52
	pointsOfFunctionPlot[248].Y = 6.252

	pointsOfFunctionPlot[249].X = -2.51
	pointsOfFunctionPlot[249].Y = 6.193

	pointsOfFunctionPlot[250].X = -2.50
	pointsOfFunctionPlot[250].Y = 6.132

	pointsOfFunctionPlot[251].X = -2.49
	pointsOfFunctionPlot[251].Y = 6.072

	pointsOfFunctionPlot[252].X = -2.48
	pointsOfFunctionPlot[252].Y = 6.012

	pointsOfFunctionPlot[253].X = -2.47
	pointsOfFunctionPlot[253].Y = 5.953

	pointsOfFunctionPlot[254].X = -2.46
	pointsOfFunctionPlot[254].Y = 5.895

	pointsOfFunctionPlot[255].X = -2.45
	pointsOfFunctionPlot[255].Y = 5.837

	pointsOfFunctionPlot[256].X = -2.44
	pointsOfFunctionPlot[256].Y = 5.78

	pointsOfFunctionPlot[257].X = -2.43
	pointsOfFunctionPlot[257].Y = 5.723

	pointsOfFunctionPlot[258].X = -2.42
	pointsOfFunctionPlot[258].Y = 5.667

	pointsOfFunctionPlot[259].X = -2.41
	pointsOfFunctionPlot[259].Y = 5.611

	pointsOfFunctionPlot[260].X = -2.40
	pointsOfFunctionPlot[260].Y = 5.556

	pointsOfFunctionPlot[261].X = -2.39
	pointsOfFunctionPlot[261].Y = 5.502

	pointsOfFunctionPlot[262].X = -2.38
	pointsOfFunctionPlot[262].Y = 5.448

	pointsOfFunctionPlot[263].X = -2.37
	pointsOfFunctionPlot[263].Y = 5.395

	pointsOfFunctionPlot[264].X = -2.36
	pointsOfFunctionPlot[264].Y = 5.342

	pointsOfFunctionPlot[265].X = -2.35
	pointsOfFunctionPlot[265].Y = 5.29

	pointsOfFunctionPlot[266].X = -2.34
	pointsOfFunctionPlot[266].Y = 5.238

	pointsOfFunctionPlot[267].X = -2.33
	pointsOfFunctionPlot[267].Y = 5.187

	pointsOfFunctionPlot[268].X = -2.32
	pointsOfFunctionPlot[268].Y = 5.136

	pointsOfFunctionPlot[269].X = -2.31
	pointsOfFunctionPlot[269].Y = 5.086

	pointsOfFunctionPlot[270].X = -2.30
	pointsOfFunctionPlot[270].Y = 5.037

	pointsOfFunctionPlot[271].X = -2.29
	pointsOfFunctionPlot[271].Y = 4.988

	pointsOfFunctionPlot[272].X = -2.28
	pointsOfFunctionPlot[272].Y = 4.939

	pointsOfFunctionPlot[273].X = -2.27
	pointsOfFunctionPlot[273].Y = 4.891

	pointsOfFunctionPlot[274].X = -2.26
	pointsOfFunctionPlot[274].Y = 4.843

	pointsOfFunctionPlot[275].X = -2.25
	pointsOfFunctionPlot[275].Y = 4.796

	pointsOfFunctionPlot[276].X = -2.24
	pointsOfFunctionPlot[276].Y = 4.749

	pointsOfFunctionPlot[277].X = -2.23
	pointsOfFunctionPlot[277].Y = 4.703

	pointsOfFunctionPlot[278].X = -2.22
	pointsOfFunctionPlot[278].Y = 4.657

	pointsOfFunctionPlot[279].X = -2.21
	pointsOfFunctionPlot[279].Y = 4.612

	pointsOfFunctionPlot[280].X = -2.20
	pointsOfFunctionPlot[280].Y = 4.567

	pointsOfFunctionPlot[281].X = -2.19
	pointsOfFunctionPlot[281].Y = 4.523

	pointsOfFunctionPlot[282].X = -2.18
	pointsOfFunctionPlot[282].Y = 4.479

	pointsOfFunctionPlot[283].X = -2.17
	pointsOfFunctionPlot[283].Y = 4.436

	pointsOfFunctionPlot[284].X = -2.16
	pointsOfFunctionPlot[284].Y = 4.393

	pointsOfFunctionPlot[285].X = -2.15
	pointsOfFunctionPlot[285].Y = 4.35

	pointsOfFunctionPlot[286].X = -2.14
	pointsOfFunctionPlot[286].Y = 4.308

	pointsOfFunctionPlot[287].X = -2.13
	pointsOfFunctionPlot[287].Y = 4.266

	pointsOfFunctionPlot[288].X = -2.12
	pointsOfFunctionPlot[288].Y = 4.224

	pointsOfFunctionPlot[289].X = -2.11
	pointsOfFunctionPlot[289].Y = 4.184

	pointsOfFunctionPlot[290].X = -2.10
	pointsOfFunctionPlot[290].Y = 4.144

	pointsOfFunctionPlot[291].X = -2.09
	pointsOfFunctionPlot[291].Y = 4.104

	pointsOfFunctionPlot[292].X = -2.08
	pointsOfFunctionPlot[292].Y = 4.064

	pointsOfFunctionPlot[293].X = -2.07
	pointsOfFunctionPlot[293].Y = 4.025

	pointsOfFunctionPlot[294].X = -2.06
	pointsOfFunctionPlot[294].Y = 3.986

	pointsOfFunctionPlot[295].X = -2.05
	pointsOfFunctionPlot[295].Y = 3.948

	pointsOfFunctionPlot[296].X = -2.04
	pointsOfFunctionPlot[296].Y = 3.91

	pointsOfFunctionPlot[297].X = -2.03
	pointsOfFunctionPlot[297].Y = 3.872

	pointsOfFunctionPlot[298].X = -2.02
	pointsOfFunctionPlot[298].Y = 3.835

	pointsOfFunctionPlot[299].X = -2.01
	pointsOfFunctionPlot[299].Y = 3.798

	pointsOfFunctionPlot[300].X = -2.0
	pointsOfFunctionPlot[300].Y = 3.762

	pointsOfFunctionPlot[301].X = -1.99
	pointsOfFunctionPlot[301].Y = 3.726

	pointsOfFunctionPlot[302].X = -1.98
	pointsOfFunctionPlot[302].Y = 3.69

	pointsOfFunctionPlot[303].X = -1.97
	pointsOfFunctionPlot[303].Y = 3.655

	pointsOfFunctionPlot[304].X = -1.96
	pointsOfFunctionPlot[304].Y = 3.62

	pointsOfFunctionPlot[305].X = -1.95
	pointsOfFunctionPlot[305].Y = 3.585

	pointsOfFunctionPlot[306].X = -1.94
	pointsOfFunctionPlot[306].Y = 3.551

	pointsOfFunctionPlot[307].X = -1.93
	pointsOfFunctionPlot[307].Y = 3.517

	pointsOfFunctionPlot[308].X = -1.92
	pointsOfFunctionPlot[308].Y = 3.483

	pointsOfFunctionPlot[309].X = -1.91
	pointsOfFunctionPlot[309].Y = 3.45

	pointsOfFunctionPlot[310].X = -1.90
	pointsOfFunctionPlot[310].Y = 3.417

	pointsOfFunctionPlot[311].X = -1.89
	pointsOfFunctionPlot[311].Y = 3.385

	pointsOfFunctionPlot[312].X = -1.88
	pointsOfFunctionPlot[312].Y = 3.353

	pointsOfFunctionPlot[313].X = -1.87
	pointsOfFunctionPlot[313].Y = 3.321

	pointsOfFunctionPlot[314].X = -1.86
	pointsOfFunctionPlot[314].Y = 3.289

	pointsOfFunctionPlot[315].X = -1.85
	pointsOfFunctionPlot[315].Y = 3.258

	pointsOfFunctionPlot[316].X = -1.84
	pointsOfFunctionPlot[316].Y = 3.227

	pointsOfFunctionPlot[317].X = -1.83
	pointsOfFunctionPlot[317].Y = 3.197

	pointsOfFunctionPlot[318].X = -1.82
	pointsOfFunctionPlot[318].Y = 3.166

	pointsOfFunctionPlot[319].X = -1.81
	pointsOfFunctionPlot[319].Y = 3.137

	pointsOfFunctionPlot[320].X = -1.80
	pointsOfFunctionPlot[320].Y = 3.107

	pointsOfFunctionPlot[321].X = -1.79
	pointsOfFunctionPlot[321].Y = 3.078

	pointsOfFunctionPlot[322].X = -1.78
	pointsOfFunctionPlot[322].Y = 3.049

	pointsOfFunctionPlot[323].X = -1.77
	pointsOfFunctionPlot[323].Y = 3.02

	pointsOfFunctionPlot[324].X = -1.76
	pointsOfFunctionPlot[324].Y = 2.992

	pointsOfFunctionPlot[325].X = -1.75
	pointsOfFunctionPlot[325].Y = 2.964

	pointsOfFunctionPlot[326].X = -1.74
	pointsOfFunctionPlot[326].Y = 2.936

	pointsOfFunctionPlot[327].X = -1.73
	pointsOfFunctionPlot[327].Y = 2.908

	pointsOfFunctionPlot[328].X = -1.72
	pointsOfFunctionPlot[328].Y = 2.881

	pointsOfFunctionPlot[329].X = -1.71
	pointsOfFunctionPlot[329].Y = 2.854

	pointsOfFunctionPlot[330].X = -1.70
	pointsOfFunctionPlot[330].Y = 2.828

	pointsOfFunctionPlot[331].X = -1.69
	pointsOfFunctionPlot[331].Y = 2.802

	pointsOfFunctionPlot[332].X = -1.68
	pointsOfFunctionPlot[332].Y = 2.775

	pointsOfFunctionPlot[333].X = -1.67
	pointsOfFunctionPlot[333].Y = 2.75

	pointsOfFunctionPlot[334].X = -1.66
	pointsOfFunctionPlot[334].Y = 2.724

	pointsOfFunctionPlot[335].X = -1.65
	pointsOfFunctionPlot[335].Y = 2.699

	pointsOfFunctionPlot[336].X = -1.64
	pointsOfFunctionPlot[336].Y = 2.674

	pointsOfFunctionPlot[337].X = -1.63
	pointsOfFunctionPlot[337].Y = 2.649

	pointsOfFunctionPlot[338].X = -1.62
	pointsOfFunctionPlot[338].Y = 2.625

	pointsOfFunctionPlot[339].X = -1.61
	pointsOfFunctionPlot[339].Y = 2.601

	pointsOfFunctionPlot[340].X = -1.60
	pointsOfFunctionPlot[340].Y = 2.577

	pointsOfFunctionPlot[341].X = -1.59
	pointsOfFunctionPlot[341].Y = 2.553

	pointsOfFunctionPlot[342].X = -1.58
	pointsOfFunctionPlot[342].Y = 2.53

	pointsOfFunctionPlot[343].X = -1.57
	pointsOfFunctionPlot[343].Y = 2.507

	pointsOfFunctionPlot[344].X = -1.56
	pointsOfFunctionPlot[344].Y = 2.484

	pointsOfFunctionPlot[345].X = -1.55
	pointsOfFunctionPlot[345].Y = 2.461

	pointsOfFunctionPlot[346].X = -1.54
	pointsOfFunctionPlot[346].Y = 2.439

	pointsOfFunctionPlot[347].X = -1.53
	pointsOfFunctionPlot[347].Y = 2.417

	pointsOfFunctionPlot[348].X = -1.52
	pointsOfFunctionPlot[348].Y = 2.395

	pointsOfFunctionPlot[349].X = -1.51
	pointsOfFunctionPlot[349].Y = 2.373

	pointsOfFunctionPlot[350].X = -1.50
	pointsOfFunctionPlot[350].Y = 2.352

	pointsOfFunctionPlot[351].X = -1.49
	pointsOfFunctionPlot[351].Y = 2.331

	pointsOfFunctionPlot[352].X = -1.48
	pointsOfFunctionPlot[352].Y = 2.31

	pointsOfFunctionPlot[353].X = -1.47
	pointsOfFunctionPlot[353].Y = 2.289

	pointsOfFunctionPlot[354].X = -1.46
	pointsOfFunctionPlot[354].Y = 2.269

	pointsOfFunctionPlot[355].X = -1.45
	pointsOfFunctionPlot[355].Y = 2.248

	pointsOfFunctionPlot[356].X = -1.44
	pointsOfFunctionPlot[356].Y = 2.228

	pointsOfFunctionPlot[357].X = -1.43
	pointsOfFunctionPlot[357].Y = 2.209

	pointsOfFunctionPlot[358].X = -1.42
	pointsOfFunctionPlot[358].Y = 2.189

	pointsOfFunctionPlot[359].X = -1.41
	pointsOfFunctionPlot[359].Y = 2.17

	pointsOfFunctionPlot[360].X = -1.40
	pointsOfFunctionPlot[360].Y = 2.15

	pointsOfFunctionPlot[361].X = -1.39
	pointsOfFunctionPlot[361].Y = 2.131

	pointsOfFunctionPlot[362].X = -1.38
	pointsOfFunctionPlot[362].Y = 2.113

	pointsOfFunctionPlot[363].X = -1.37
	pointsOfFunctionPlot[363].Y = 2.094

	pointsOfFunctionPlot[364].X = -1.36
	pointsOfFunctionPlot[364].Y = 2.076

	pointsOfFunctionPlot[365].X = -1.35
	pointsOfFunctionPlot[365].Y = 2.058

	pointsOfFunctionPlot[366].X = -1.34
	pointsOfFunctionPlot[366].Y = 2.04

	pointsOfFunctionPlot[367].X = -1.33
	pointsOfFunctionPlot[367].Y = 2.022

	pointsOfFunctionPlot[368].X = -1.32
	pointsOfFunctionPlot[368].Y = 2.005

	pointsOfFunctionPlot[369].X = -1.31
	pointsOfFunctionPlot[369].Y = 1.987

	pointsOfFunctionPlot[370].X = -1.30
	pointsOfFunctionPlot[370].Y = 1.97

	pointsOfFunctionPlot[371].X = -1.29
	pointsOfFunctionPlot[371].Y = 1.954

	pointsOfFunctionPlot[372].X = -1.28
	pointsOfFunctionPlot[372].Y = 1.937

	pointsOfFunctionPlot[373].X = -1.27
	pointsOfFunctionPlot[373].Y = 1.92

	pointsOfFunctionPlot[374].X = -1.26
	pointsOfFunctionPlot[374].Y = 1.904

	pointsOfFunctionPlot[375].X = -1.25
	pointsOfFunctionPlot[375].Y = 1.888

	pointsOfFunctionPlot[376].X = -1.24
	pointsOfFunctionPlot[376].Y = 1.872

	pointsOfFunctionPlot[377].X = -1.23
	pointsOfFunctionPlot[377].Y = 1.856

	pointsOfFunctionPlot[378].X = -1.22
	pointsOfFunctionPlot[378].Y = 1.841

	pointsOfFunctionPlot[379].X = -1.21
	pointsOfFunctionPlot[379].Y = 1.825

	pointsOfFunctionPlot[380].X = -1.20
	pointsOfFunctionPlot[380].Y = 1.81

	pointsOfFunctionPlot[381].X = -1.19
	pointsOfFunctionPlot[381].Y = 1.795

	pointsOfFunctionPlot[382].X = -1.18
	pointsOfFunctionPlot[382].Y = 1.78

	pointsOfFunctionPlot[383].X = -1.17
	pointsOfFunctionPlot[383].Y = 1.766

	pointsOfFunctionPlot[384].X = -1.16
	pointsOfFunctionPlot[384].Y = 1.751

	pointsOfFunctionPlot[385].X = -1.15
	pointsOfFunctionPlot[385].Y = 1.737

	pointsOfFunctionPlot[386].X = -1.14
	pointsOfFunctionPlot[386].Y = 1.723

	pointsOfFunctionPlot[387].X = -1.13
	pointsOfFunctionPlot[387].Y = 1.709

	pointsOfFunctionPlot[388].X = -1.12
	pointsOfFunctionPlot[388].Y = 1.695

	pointsOfFunctionPlot[389].X = -1.11
	pointsOfFunctionPlot[389].Y = 1.681

	pointsOfFunctionPlot[390].X = -1.10
	pointsOfFunctionPlot[390].Y = 1.668

	pointsOfFunctionPlot[391].X = -1.09
	pointsOfFunctionPlot[391].Y = 1.655

	pointsOfFunctionPlot[392].X = -1.08
	pointsOfFunctionPlot[392].Y = 1.642

	pointsOfFunctionPlot[393].X = -1.07
	pointsOfFunctionPlot[393].Y = 1.629

	pointsOfFunctionPlot[394].X = -1.06
	pointsOfFunctionPlot[394].Y = 1.616

	pointsOfFunctionPlot[395].X = -1.05
	pointsOfFunctionPlot[395].Y = 1.603

	pointsOfFunctionPlot[396].X = -1.04
	pointsOfFunctionPlot[396].Y = 1.591

	pointsOfFunctionPlot[397].X = -1.03
	pointsOfFunctionPlot[397].Y = 1.579

	pointsOfFunctionPlot[398].X = -1.02
	pointsOfFunctionPlot[398].Y = 1.566

	pointsOfFunctionPlot[399].X = -1.01
	pointsOfFunctionPlot[399].Y = 1.554

	pointsOfFunctionPlot[400].X = -1.0
	pointsOfFunctionPlot[400].Y = 1.543

	pointsOfFunctionPlot[401].X = -0.99
	pointsOfFunctionPlot[401].Y = 1.531

	pointsOfFunctionPlot[402].X = -0.98
	pointsOfFunctionPlot[402].Y = 1.519

	pointsOfFunctionPlot[403].X = -0.97
	pointsOfFunctionPlot[403].Y = 1.508

	pointsOfFunctionPlot[404].X = -0.96
	pointsOfFunctionPlot[404].Y = 1.497

	pointsOfFunctionPlot[405].X = -0.95
	pointsOfFunctionPlot[405].Y = 1.486

	pointsOfFunctionPlot[406].X = -0.94
	pointsOfFunctionPlot[406].Y = 1.475

	pointsOfFunctionPlot[407].X = -0.93
	pointsOfFunctionPlot[407].Y = 1.464

	pointsOfFunctionPlot[408].X = -0.92
	pointsOfFunctionPlot[408].Y = 1.453

	pointsOfFunctionPlot[409].X = -0.91
	pointsOfFunctionPlot[409].Y = 1.443

	pointsOfFunctionPlot[410].X = -0.90
	pointsOfFunctionPlot[410].Y = 1.433

	pointsOfFunctionPlot[411].X = -0.89
	pointsOfFunctionPlot[411].Y = 1.422

	pointsOfFunctionPlot[412].X = -0.88
	pointsOfFunctionPlot[412].Y = 1.412

	pointsOfFunctionPlot[413].X = -0.87
	pointsOfFunctionPlot[413].Y = 1.402

	pointsOfFunctionPlot[414].X = -0.86
	pointsOfFunctionPlot[414].Y = 1.393

	pointsOfFunctionPlot[415].X = -0.85
	pointsOfFunctionPlot[415].Y = 1.383

	pointsOfFunctionPlot[416].X = -0.84
	pointsOfFunctionPlot[416].Y = 1.374

	pointsOfFunctionPlot[417].X = -0.83
	pointsOfFunctionPlot[417].Y = 1.364

	pointsOfFunctionPlot[418].X = -0.82
	pointsOfFunctionPlot[418].Y = 1.355

	pointsOfFunctionPlot[419].X = -0.81
	pointsOfFunctionPlot[419].Y = 1.346

	pointsOfFunctionPlot[420].X = -0.80
	pointsOfFunctionPlot[420].Y = 1.337

	pointsOfFunctionPlot[421].X = -0.79
	pointsOfFunctionPlot[421].Y = 1.328

	pointsOfFunctionPlot[422].X = -0.78
	pointsOfFunctionPlot[422].Y = 1.319

	pointsOfFunctionPlot[423].X = -0.77
	pointsOfFunctionPlot[423].Y = 1.311

	pointsOfFunctionPlot[424].X = -0.76
	pointsOfFunctionPlot[424].Y = 1.302

	pointsOfFunctionPlot[425].X = -0.75
	pointsOfFunctionPlot[425].Y = 1.294

	pointsOfFunctionPlot[426].X = -0.74
	pointsOfFunctionPlot[426].Y = 1.286

	pointsOfFunctionPlot[427].X = -0.73
	pointsOfFunctionPlot[427].Y = 1.278

	pointsOfFunctionPlot[428].X = -0.72
	pointsOfFunctionPlot[428].Y = 1.27

	pointsOfFunctionPlot[429].X = -0.71
	pointsOfFunctionPlot[429].Y = 1.262

	pointsOfFunctionPlot[430].X = -0.70
	pointsOfFunctionPlot[430].Y = 1.255

	pointsOfFunctionPlot[431].X = -0.69
	pointsOfFunctionPlot[431].Y = 1.247

	pointsOfFunctionPlot[432].X = -0.68
	pointsOfFunctionPlot[432].Y = 1.24

	pointsOfFunctionPlot[433].X = -0.67
	pointsOfFunctionPlot[433].Y = 1.232

	pointsOfFunctionPlot[434].X = -0.66
	pointsOfFunctionPlot[434].Y = 1.225

	pointsOfFunctionPlot[435].X = -0.65
	pointsOfFunctionPlot[435].Y = 1.218

	pointsOfFunctionPlot[436].X = -0.64
	pointsOfFunctionPlot[436].Y = 1.211

	pointsOfFunctionPlot[437].X = -0.63
	pointsOfFunctionPlot[437].Y = 1.205

	pointsOfFunctionPlot[438].X = -0.62
	pointsOfFunctionPlot[438].Y = 1.198

	pointsOfFunctionPlot[439].X = -0.61
	pointsOfFunctionPlot[439].Y = 1.191

	pointsOfFunctionPlot[440].X = -0.60
	pointsOfFunctionPlot[440].Y = 1.185

	pointsOfFunctionPlot[441].X = -0.59
	pointsOfFunctionPlot[441].Y = 1.179

	pointsOfFunctionPlot[442].X = -0.58
	pointsOfFunctionPlot[442].Y = 1.172

	pointsOfFunctionPlot[443].X = -0.57
	pointsOfFunctionPlot[443].Y = 1.166

	pointsOfFunctionPlot[444].X = -0.56
	pointsOfFunctionPlot[444].Y = 1.16

	pointsOfFunctionPlot[445].X = -0.55
	pointsOfFunctionPlot[445].Y = 1.155

	pointsOfFunctionPlot[446].X = -0.54
	pointsOfFunctionPlot[446].Y = 1.149

	pointsOfFunctionPlot[447].X = -0.53
	pointsOfFunctionPlot[447].Y = 1.143

	pointsOfFunctionPlot[448].X = -0.52
	pointsOfFunctionPlot[448].Y = 1.138

	pointsOfFunctionPlot[449].X = -0.51
	pointsOfFunctionPlot[449].Y = 1.132

	pointsOfFunctionPlot[450].X = -0.50
	pointsOfFunctionPlot[450].Y = 1.127

	pointsOfFunctionPlot[451].X = -0.49
	pointsOfFunctionPlot[451].Y = 1.122

	pointsOfFunctionPlot[452].X = -0.48
	pointsOfFunctionPlot[452].Y = 1.117

	pointsOfFunctionPlot[453].X = -0.47
	pointsOfFunctionPlot[453].Y = 1.112

	pointsOfFunctionPlot[454].X = -0.46
	pointsOfFunctionPlot[454].Y = 1.107

	pointsOfFunctionPlot[455].X = -0.45
	pointsOfFunctionPlot[455].Y = 1.102

	pointsOfFunctionPlot[456].X = -0.44
	pointsOfFunctionPlot[456].Y = 1.098

	pointsOfFunctionPlot[457].X = -0.43
	pointsOfFunctionPlot[457].Y = 1.093

	pointsOfFunctionPlot[458].X = -0.42
	pointsOfFunctionPlot[458].Y = 1.089

	pointsOfFunctionPlot[459].X = -0.41
	pointsOfFunctionPlot[459].Y = 1.085

	pointsOfFunctionPlot[460].X = -0.40
	pointsOfFunctionPlot[460].Y = 1.081

	pointsOfFunctionPlot[461].X = -0.39
	pointsOfFunctionPlot[461].Y = 1.077

	pointsOfFunctionPlot[462].X = -0.38
	pointsOfFunctionPlot[462].Y = 1.073

	pointsOfFunctionPlot[463].X = -0.37
	pointsOfFunctionPlot[463].Y = 1.069

	pointsOfFunctionPlot[464].X = -0.36
	pointsOfFunctionPlot[464].Y = 1.065

	pointsOfFunctionPlot[465].X = -0.35
	pointsOfFunctionPlot[465].Y = 1.061

	pointsOfFunctionPlot[466].X = -0.34
	pointsOfFunctionPlot[466].Y = 1.058

	pointsOfFunctionPlot[467].X = -0.33
	pointsOfFunctionPlot[467].Y = 1.054

	pointsOfFunctionPlot[468].X = -0.32
	pointsOfFunctionPlot[468].Y = 1.051

	pointsOfFunctionPlot[469].X = -0.31
	pointsOfFunctionPlot[469].Y = 1.048

	pointsOfFunctionPlot[470].X = -0.30
	pointsOfFunctionPlot[470].Y = 1.045

	pointsOfFunctionPlot[471].X = -0.29
	pointsOfFunctionPlot[471].Y = 1.042

	pointsOfFunctionPlot[472].X = -0.28
	pointsOfFunctionPlot[472].Y = 1.039

	pointsOfFunctionPlot[473].X = -0.27
	pointsOfFunctionPlot[473].Y = 1.036

	pointsOfFunctionPlot[474].X = -0.26
	pointsOfFunctionPlot[474].Y = 1.033

	pointsOfFunctionPlot[475].X = -0.25
	pointsOfFunctionPlot[475].Y = 1.031

	pointsOfFunctionPlot[476].X = -0.24
	pointsOfFunctionPlot[476].Y = 1.028

	pointsOfFunctionPlot[477].X = -0.23
	pointsOfFunctionPlot[477].Y = 1.026

	pointsOfFunctionPlot[478].X = -0.22
	pointsOfFunctionPlot[478].Y = 1.024

	pointsOfFunctionPlot[479].X = -0.21
	pointsOfFunctionPlot[479].Y = 1.022

	pointsOfFunctionPlot[480].X = -0.20
	pointsOfFunctionPlot[480].Y = 1.02

	pointsOfFunctionPlot[481].X = -0.19
	pointsOfFunctionPlot[481].Y = 1.018

	pointsOfFunctionPlot[482].X = -0.18
	pointsOfFunctionPlot[482].Y = 1.016

	pointsOfFunctionPlot[483].X = -0.17
	pointsOfFunctionPlot[483].Y = 1.014

	pointsOfFunctionPlot[484].X = -0.16
	pointsOfFunctionPlot[484].Y = 1.012

	pointsOfFunctionPlot[485].X = -0.15
	pointsOfFunctionPlot[485].Y = 1.011

	pointsOfFunctionPlot[486].X = -0.14
	pointsOfFunctionPlot[486].Y = 1.009

	pointsOfFunctionPlot[487].X = -0.13
	pointsOfFunctionPlot[487].Y = 1.008

	pointsOfFunctionPlot[488].X = -0.12
	pointsOfFunctionPlot[488].Y = 1.007

	pointsOfFunctionPlot[489].X = -0.11
	pointsOfFunctionPlot[489].Y = 1.006

	pointsOfFunctionPlot[490].X = -0.10
	pointsOfFunctionPlot[490].Y = 1.005

	pointsOfFunctionPlot[491].X = -0.09
	pointsOfFunctionPlot[491].Y = 1.004

	pointsOfFunctionPlot[492].X = -0.08
	pointsOfFunctionPlot[492].Y = 1.003

	pointsOfFunctionPlot[493].X = -0.07
	pointsOfFunctionPlot[493].Y = 1.002

	pointsOfFunctionPlot[494].X = -0.06
	pointsOfFunctionPlot[494].Y = 1.001

	pointsOfFunctionPlot[495].X = -0.05
	pointsOfFunctionPlot[495].Y = 1.001

	pointsOfFunctionPlot[496].X = -0.04
	pointsOfFunctionPlot[496].Y = 1.0

	pointsOfFunctionPlot[497].X = -0.03
	pointsOfFunctionPlot[497].Y = 1.0

	pointsOfFunctionPlot[498].X = -0.02
	pointsOfFunctionPlot[498].Y = 1.0

	pointsOfFunctionPlot[499].X = -0.01
	pointsOfFunctionPlot[499].Y = 1.0

	pointsOfFunctionPlot[500].X = 0.0
	pointsOfFunctionPlot[500].Y = 1.0

	pointsOfFunctionPlot[501].X = 0.01
	pointsOfFunctionPlot[501].Y = 1.0

	pointsOfFunctionPlot[502].X = 0.02
	pointsOfFunctionPlot[502].Y = 1.0

	pointsOfFunctionPlot[503].X = 0.03
	pointsOfFunctionPlot[503].Y = 1.0

	pointsOfFunctionPlot[504].X = 0.04
	pointsOfFunctionPlot[504].Y = 1.0

	pointsOfFunctionPlot[505].X = 0.05
	pointsOfFunctionPlot[505].Y = 1.001

	pointsOfFunctionPlot[506].X = 0.06
	pointsOfFunctionPlot[506].Y = 1.001

	pointsOfFunctionPlot[507].X = 0.07
	pointsOfFunctionPlot[507].Y = 1.002

	pointsOfFunctionPlot[508].X = 0.08
	pointsOfFunctionPlot[508].Y = 1.003

	pointsOfFunctionPlot[509].X = 0.09
	pointsOfFunctionPlot[509].Y = 1.004

	pointsOfFunctionPlot[510].X = 0.10
	pointsOfFunctionPlot[510].Y = 1.005

	pointsOfFunctionPlot[511].X = 0.11
	pointsOfFunctionPlot[511].Y = 1.006

	pointsOfFunctionPlot[512].X = 0.12
	pointsOfFunctionPlot[512].Y = 1.007

	pointsOfFunctionPlot[513].X = 0.13
	pointsOfFunctionPlot[513].Y = 1.008

	pointsOfFunctionPlot[514].X = 0.14
	pointsOfFunctionPlot[514].Y = 1.009

	pointsOfFunctionPlot[515].X = 0.15
	pointsOfFunctionPlot[515].Y = 1.011

	pointsOfFunctionPlot[516].X = 0.16
	pointsOfFunctionPlot[516].Y = 1.012

	pointsOfFunctionPlot[517].X = 0.17
	pointsOfFunctionPlot[517].Y = 1.014

	pointsOfFunctionPlot[518].X = 0.18
	pointsOfFunctionPlot[518].Y = 1.016

	pointsOfFunctionPlot[519].X = 0.19
	pointsOfFunctionPlot[519].Y = 1.018

	pointsOfFunctionPlot[520].X = 0.20
	pointsOfFunctionPlot[520].Y = 1.02

	pointsOfFunctionPlot[521].X = 0.21
	pointsOfFunctionPlot[521].Y = 1.022

	pointsOfFunctionPlot[522].X = 0.22
	pointsOfFunctionPlot[522].Y = 1.024

	pointsOfFunctionPlot[523].X = 0.23
	pointsOfFunctionPlot[523].Y = 1.026

	pointsOfFunctionPlot[524].X = 0.24
	pointsOfFunctionPlot[524].Y = 1.028

	pointsOfFunctionPlot[525].X = 0.25
	pointsOfFunctionPlot[525].Y = 1.031

	pointsOfFunctionPlot[526].X = 0.26
	pointsOfFunctionPlot[526].Y = 1.033

	pointsOfFunctionPlot[527].X = 0.27
	pointsOfFunctionPlot[527].Y = 1.036

	pointsOfFunctionPlot[528].X = 0.28
	pointsOfFunctionPlot[528].Y = 1.039

	pointsOfFunctionPlot[529].X = 0.29
	pointsOfFunctionPlot[529].Y = 1.042

	pointsOfFunctionPlot[530].X = 0.30
	pointsOfFunctionPlot[530].Y = 1.045

	pointsOfFunctionPlot[531].X = 0.31
	pointsOfFunctionPlot[531].Y = 1.048

	pointsOfFunctionPlot[532].X = 0.32
	pointsOfFunctionPlot[532].Y = 1.051

	pointsOfFunctionPlot[533].X = 0.33
	pointsOfFunctionPlot[533].Y = 1.054

	pointsOfFunctionPlot[534].X = 0.34
	pointsOfFunctionPlot[534].Y = 1.058

	pointsOfFunctionPlot[535].X = 0.35
	pointsOfFunctionPlot[535].Y = 1.061

	pointsOfFunctionPlot[536].X = 0.36
	pointsOfFunctionPlot[536].Y = 1.065

	pointsOfFunctionPlot[537].X = 0.37
	pointsOfFunctionPlot[537].Y = 1.069

	pointsOfFunctionPlot[538].X = 0.38
	pointsOfFunctionPlot[538].Y = 1.073

	pointsOfFunctionPlot[539].X = 0.39
	pointsOfFunctionPlot[539].Y = 1.077

	pointsOfFunctionPlot[540].X = 0.40
	pointsOfFunctionPlot[540].Y = 1.081

	pointsOfFunctionPlot[541].X = 0.41
	pointsOfFunctionPlot[541].Y = 1.085

	pointsOfFunctionPlot[542].X = 0.42
	pointsOfFunctionPlot[542].Y = 1.089

	pointsOfFunctionPlot[543].X = 0.43
	pointsOfFunctionPlot[543].Y = 1.093

	pointsOfFunctionPlot[544].X = 0.44
	pointsOfFunctionPlot[544].Y = 1.098

	pointsOfFunctionPlot[545].X = 0.45
	pointsOfFunctionPlot[545].Y = 1.102

	pointsOfFunctionPlot[546].X = 0.46
	pointsOfFunctionPlot[546].Y = 1.107

	pointsOfFunctionPlot[547].X = 0.47
	pointsOfFunctionPlot[547].Y = 1.112

	pointsOfFunctionPlot[548].X = 0.48
	pointsOfFunctionPlot[548].Y = 1.117

	pointsOfFunctionPlot[549].X = 0.49
	pointsOfFunctionPlot[549].Y = 1.122

	pointsOfFunctionPlot[550].X = 0.50
	pointsOfFunctionPlot[550].Y = 1.127

	pointsOfFunctionPlot[551].X = 0.51
	pointsOfFunctionPlot[551].Y = 1.132

	pointsOfFunctionPlot[552].X = 0.52
	pointsOfFunctionPlot[552].Y = 1.138

	pointsOfFunctionPlot[553].X = 0.53
	pointsOfFunctionPlot[553].Y = 1.143

	pointsOfFunctionPlot[554].X = 0.54
	pointsOfFunctionPlot[554].Y = 1.149

	pointsOfFunctionPlot[555].X = 0.55
	pointsOfFunctionPlot[555].Y = 1.155

	pointsOfFunctionPlot[556].X = 0.56
	pointsOfFunctionPlot[556].Y = 1.16

	pointsOfFunctionPlot[557].X = 0.57
	pointsOfFunctionPlot[557].Y = 1.166

	pointsOfFunctionPlot[558].X = 0.58
	pointsOfFunctionPlot[558].Y = 1.172

	pointsOfFunctionPlot[559].X = 0.59
	pointsOfFunctionPlot[559].Y = 1.179

	pointsOfFunctionPlot[560].X = 0.60
	pointsOfFunctionPlot[560].Y = 1.185

	pointsOfFunctionPlot[561].X = 0.61
	pointsOfFunctionPlot[561].Y = 1.191

	pointsOfFunctionPlot[562].X = 0.62
	pointsOfFunctionPlot[562].Y = 1.198

	pointsOfFunctionPlot[563].X = 0.63
	pointsOfFunctionPlot[563].Y = 1.205

	pointsOfFunctionPlot[564].X = 0.64
	pointsOfFunctionPlot[564].Y = 1.211

	pointsOfFunctionPlot[565].X = 0.65
	pointsOfFunctionPlot[565].Y = 1.218

	pointsOfFunctionPlot[566].X = 0.66
	pointsOfFunctionPlot[566].Y = 1.225

	pointsOfFunctionPlot[567].X = 0.67
	pointsOfFunctionPlot[567].Y = 1.232

	pointsOfFunctionPlot[568].X = 0.68
	pointsOfFunctionPlot[568].Y = 1.24

	pointsOfFunctionPlot[569].X = 0.69
	pointsOfFunctionPlot[569].Y = 1.247

	pointsOfFunctionPlot[570].X = 0.70
	pointsOfFunctionPlot[570].Y = 1.255

	pointsOfFunctionPlot[571].X = 0.71
	pointsOfFunctionPlot[571].Y = 1.262

	pointsOfFunctionPlot[572].X = 0.72
	pointsOfFunctionPlot[572].Y = 1.27

	pointsOfFunctionPlot[573].X = 0.73
	pointsOfFunctionPlot[573].Y = 1.278

	pointsOfFunctionPlot[574].X = 0.74
	pointsOfFunctionPlot[574].Y = 1.286

	pointsOfFunctionPlot[575].X = 0.75
	pointsOfFunctionPlot[575].Y = 1.294

	pointsOfFunctionPlot[576].X = 0.76
	pointsOfFunctionPlot[576].Y = 1.302

	pointsOfFunctionPlot[577].X = 0.77
	pointsOfFunctionPlot[577].Y = 1.311

	pointsOfFunctionPlot[578].X = 0.78
	pointsOfFunctionPlot[578].Y = 1.319

	pointsOfFunctionPlot[579].X = 0.79
	pointsOfFunctionPlot[579].Y = 1.328

	pointsOfFunctionPlot[580].X = 0.80
	pointsOfFunctionPlot[580].Y = 1.337

	pointsOfFunctionPlot[581].X = 0.81
	pointsOfFunctionPlot[581].Y = 1.346

	pointsOfFunctionPlot[582].X = 0.82
	pointsOfFunctionPlot[582].Y = 1.355

	pointsOfFunctionPlot[583].X = 0.83
	pointsOfFunctionPlot[583].Y = 1.364

	pointsOfFunctionPlot[584].X = 0.84
	pointsOfFunctionPlot[584].Y = 1.374

	pointsOfFunctionPlot[585].X = 0.85
	pointsOfFunctionPlot[585].Y = 1.383

	pointsOfFunctionPlot[586].X = 0.86
	pointsOfFunctionPlot[586].Y = 1.393

	pointsOfFunctionPlot[587].X = 0.87
	pointsOfFunctionPlot[587].Y = 1.402

	pointsOfFunctionPlot[588].X = 0.88
	pointsOfFunctionPlot[588].Y = 1.412

	pointsOfFunctionPlot[589].X = 0.89
	pointsOfFunctionPlot[589].Y = 1.422

	pointsOfFunctionPlot[590].X = 0.90
	pointsOfFunctionPlot[590].Y = 1.433

	pointsOfFunctionPlot[591].X = 0.91
	pointsOfFunctionPlot[591].Y = 1.443

	pointsOfFunctionPlot[592].X = 0.92
	pointsOfFunctionPlot[592].Y = 1.453

	pointsOfFunctionPlot[593].X = 0.93
	pointsOfFunctionPlot[593].Y = 1.464

	pointsOfFunctionPlot[594].X = 0.94
	pointsOfFunctionPlot[594].Y = 1.475

	pointsOfFunctionPlot[595].X = 0.95
	pointsOfFunctionPlot[595].Y = 1.486

	pointsOfFunctionPlot[596].X = 0.96
	pointsOfFunctionPlot[596].Y = 1.497

	pointsOfFunctionPlot[597].X = 0.97
	pointsOfFunctionPlot[597].Y = 1.508

	pointsOfFunctionPlot[598].X = 0.98
	pointsOfFunctionPlot[598].Y = 1.519

	pointsOfFunctionPlot[599].X = 0.99
	pointsOfFunctionPlot[599].Y = 1.531

	pointsOfFunctionPlot[600].X = 1.0
	pointsOfFunctionPlot[600].Y = 1.543

	pointsOfFunctionPlot[601].X = 1.01
	pointsOfFunctionPlot[601].Y = 1.554

	pointsOfFunctionPlot[602].X = 1.02
	pointsOfFunctionPlot[602].Y = 1.566

	pointsOfFunctionPlot[603].X = 1.03
	pointsOfFunctionPlot[603].Y = 1.579

	pointsOfFunctionPlot[604].X = 1.04
	pointsOfFunctionPlot[604].Y = 1.591

	pointsOfFunctionPlot[605].X = 1.05
	pointsOfFunctionPlot[605].Y = 1.603

	pointsOfFunctionPlot[606].X = 1.06
	pointsOfFunctionPlot[606].Y = 1.616

	pointsOfFunctionPlot[607].X = 1.07
	pointsOfFunctionPlot[607].Y = 1.629

	pointsOfFunctionPlot[608].X = 1.08
	pointsOfFunctionPlot[608].Y = 1.642

	pointsOfFunctionPlot[609].X = 1.09
	pointsOfFunctionPlot[609].Y = 1.655

	pointsOfFunctionPlot[610].X = 1.10
	pointsOfFunctionPlot[610].Y = 1.668

	pointsOfFunctionPlot[611].X = 1.11
	pointsOfFunctionPlot[611].Y = 1.681

	pointsOfFunctionPlot[612].X = 1.12
	pointsOfFunctionPlot[612].Y = 1.695

	pointsOfFunctionPlot[613].X = 1.13
	pointsOfFunctionPlot[613].Y = 1.709

	pointsOfFunctionPlot[614].X = 1.14
	pointsOfFunctionPlot[614].Y = 1.723

	pointsOfFunctionPlot[615].X = 1.15
	pointsOfFunctionPlot[615].Y = 1.737

	pointsOfFunctionPlot[616].X = 1.16
	pointsOfFunctionPlot[616].Y = 1.751

	pointsOfFunctionPlot[617].X = 1.17
	pointsOfFunctionPlot[617].Y = 1.766

	pointsOfFunctionPlot[618].X = 1.18
	pointsOfFunctionPlot[618].Y = 1.78

	pointsOfFunctionPlot[619].X = 1.19
	pointsOfFunctionPlot[619].Y = 1.795

	pointsOfFunctionPlot[620].X = 1.20
	pointsOfFunctionPlot[620].Y = 1.81

	pointsOfFunctionPlot[621].X = 1.21
	pointsOfFunctionPlot[621].Y = 1.825

	pointsOfFunctionPlot[622].X = 1.22
	pointsOfFunctionPlot[622].Y = 1.841

	pointsOfFunctionPlot[623].X = 1.23
	pointsOfFunctionPlot[623].Y = 1.856

	pointsOfFunctionPlot[624].X = 1.24
	pointsOfFunctionPlot[624].Y = 1.872

	pointsOfFunctionPlot[625].X = 1.25
	pointsOfFunctionPlot[625].Y = 1.888

	pointsOfFunctionPlot[626].X = 1.26
	pointsOfFunctionPlot[626].Y = 1.904

	pointsOfFunctionPlot[627].X = 1.27
	pointsOfFunctionPlot[627].Y = 1.92

	pointsOfFunctionPlot[628].X = 1.28
	pointsOfFunctionPlot[628].Y = 1.937

	pointsOfFunctionPlot[629].X = 1.29
	pointsOfFunctionPlot[629].Y = 1.954

	pointsOfFunctionPlot[630].X = 1.30
	pointsOfFunctionPlot[630].Y = 1.97

	pointsOfFunctionPlot[631].X = 1.31
	pointsOfFunctionPlot[631].Y = 1.987

	pointsOfFunctionPlot[632].X = 1.32
	pointsOfFunctionPlot[632].Y = 2.005

	pointsOfFunctionPlot[633].X = 1.33
	pointsOfFunctionPlot[633].Y = 2.022

	pointsOfFunctionPlot[634].X = 1.34
	pointsOfFunctionPlot[634].Y = 2.04

	pointsOfFunctionPlot[635].X = 1.35
	pointsOfFunctionPlot[635].Y = 2.058

	pointsOfFunctionPlot[636].X = 1.36
	pointsOfFunctionPlot[636].Y = 2.076

	pointsOfFunctionPlot[637].X = 1.37
	pointsOfFunctionPlot[637].Y = 2.094

	pointsOfFunctionPlot[638].X = 1.38
	pointsOfFunctionPlot[638].Y = 2.113

	pointsOfFunctionPlot[639].X = 1.39
	pointsOfFunctionPlot[639].Y = 2.131

	pointsOfFunctionPlot[640].X = 1.40
	pointsOfFunctionPlot[640].Y = 2.15

	pointsOfFunctionPlot[641].X = 1.41
	pointsOfFunctionPlot[641].Y = 2.17

	pointsOfFunctionPlot[642].X = 1.42
	pointsOfFunctionPlot[642].Y = 2.189

	pointsOfFunctionPlot[643].X = 1.43
	pointsOfFunctionPlot[643].Y = 2.209

	pointsOfFunctionPlot[644].X = 1.44
	pointsOfFunctionPlot[644].Y = 2.228

	pointsOfFunctionPlot[645].X = 1.45
	pointsOfFunctionPlot[645].Y = 2.248

	pointsOfFunctionPlot[646].X = 1.46
	pointsOfFunctionPlot[646].Y = 2.269

	pointsOfFunctionPlot[647].X = 1.47
	pointsOfFunctionPlot[647].Y = 2.289

	pointsOfFunctionPlot[648].X = 1.48
	pointsOfFunctionPlot[648].Y = 2.31

	pointsOfFunctionPlot[649].X = 1.49
	pointsOfFunctionPlot[649].Y = 2.331

	pointsOfFunctionPlot[650].X = 1.50
	pointsOfFunctionPlot[650].Y = 2.352

	pointsOfFunctionPlot[651].X = 1.51
	pointsOfFunctionPlot[651].Y = 2.373

	pointsOfFunctionPlot[652].X = 1.52
	pointsOfFunctionPlot[652].Y = 2.395

	pointsOfFunctionPlot[653].X = 1.53
	pointsOfFunctionPlot[653].Y = 2.417

	pointsOfFunctionPlot[654].X = 1.54
	pointsOfFunctionPlot[654].Y = 2.439

	pointsOfFunctionPlot[655].X = 1.55
	pointsOfFunctionPlot[655].Y = 2.461

	pointsOfFunctionPlot[656].X = 1.56
	pointsOfFunctionPlot[656].Y = 2.484

	pointsOfFunctionPlot[657].X = 1.57
	pointsOfFunctionPlot[657].Y = 2.507

	pointsOfFunctionPlot[658].X = 1.58
	pointsOfFunctionPlot[658].Y = 2.53

	pointsOfFunctionPlot[659].X = 1.59
	pointsOfFunctionPlot[659].Y = 2.553

	pointsOfFunctionPlot[660].X = 1.60
	pointsOfFunctionPlot[660].Y = 2.577

	pointsOfFunctionPlot[661].X = 1.61
	pointsOfFunctionPlot[661].Y = 2.601

	pointsOfFunctionPlot[662].X = 1.62
	pointsOfFunctionPlot[662].Y = 2.625

	pointsOfFunctionPlot[663].X = 1.63
	pointsOfFunctionPlot[663].Y = 2.649

	pointsOfFunctionPlot[664].X = 1.64
	pointsOfFunctionPlot[664].Y = 2.674

	pointsOfFunctionPlot[665].X = 1.65
	pointsOfFunctionPlot[665].Y = 2.699

	pointsOfFunctionPlot[666].X = 1.66
	pointsOfFunctionPlot[666].Y = 2.724

	pointsOfFunctionPlot[667].X = 1.67
	pointsOfFunctionPlot[667].Y = 2.75

	pointsOfFunctionPlot[668].X = 1.68
	pointsOfFunctionPlot[668].Y = 2.775

	pointsOfFunctionPlot[669].X = 1.69
	pointsOfFunctionPlot[669].Y = 2.802

	pointsOfFunctionPlot[670].X = 1.70
	pointsOfFunctionPlot[670].Y = 2.828

	pointsOfFunctionPlot[671].X = 1.71
	pointsOfFunctionPlot[671].Y = 2.854

	pointsOfFunctionPlot[672].X = 1.72
	pointsOfFunctionPlot[672].Y = 2.881

	pointsOfFunctionPlot[673].X = 1.73
	pointsOfFunctionPlot[673].Y = 2.908

	pointsOfFunctionPlot[674].X = 1.74
	pointsOfFunctionPlot[674].Y = 2.936

	pointsOfFunctionPlot[675].X = 1.75
	pointsOfFunctionPlot[675].Y = 2.964

	pointsOfFunctionPlot[676].X = 1.76
	pointsOfFunctionPlot[676].Y = 2.992

	pointsOfFunctionPlot[677].X = 1.77
	pointsOfFunctionPlot[677].Y = 3.02

	pointsOfFunctionPlot[678].X = 1.78
	pointsOfFunctionPlot[678].Y = 3.049

	pointsOfFunctionPlot[679].X = 1.79
	pointsOfFunctionPlot[679].Y = 3.078

	pointsOfFunctionPlot[680].X = 1.80
	pointsOfFunctionPlot[680].Y = 3.107

	pointsOfFunctionPlot[681].X = 1.81
	pointsOfFunctionPlot[681].Y = 3.137

	pointsOfFunctionPlot[682].X = 1.82
	pointsOfFunctionPlot[682].Y = 3.166

	pointsOfFunctionPlot[683].X = 1.83
	pointsOfFunctionPlot[683].Y = 3.197

	pointsOfFunctionPlot[684].X = 1.84
	pointsOfFunctionPlot[684].Y = 3.227

	pointsOfFunctionPlot[685].X = 1.85
	pointsOfFunctionPlot[685].Y = 3.258

	pointsOfFunctionPlot[686].X = 1.86
	pointsOfFunctionPlot[686].Y = 3.289

	pointsOfFunctionPlot[687].X = 1.87
	pointsOfFunctionPlot[687].Y = 3.321

	pointsOfFunctionPlot[688].X = 1.88
	pointsOfFunctionPlot[688].Y = 3.353

	pointsOfFunctionPlot[689].X = 1.89
	pointsOfFunctionPlot[689].Y = 3.385

	pointsOfFunctionPlot[690].X = 1.90
	pointsOfFunctionPlot[690].Y = 3.417

	pointsOfFunctionPlot[691].X = 1.91
	pointsOfFunctionPlot[691].Y = 3.45

	pointsOfFunctionPlot[692].X = 1.92
	pointsOfFunctionPlot[692].Y = 3.483

	pointsOfFunctionPlot[693].X = 1.93
	pointsOfFunctionPlot[693].Y = 3.517

	pointsOfFunctionPlot[694].X = 1.94
	pointsOfFunctionPlot[694].Y = 3.551

	pointsOfFunctionPlot[695].X = 1.95
	pointsOfFunctionPlot[695].Y = 3.585

	pointsOfFunctionPlot[696].X = 1.96
	pointsOfFunctionPlot[696].Y = 3.62

	pointsOfFunctionPlot[697].X = 1.97
	pointsOfFunctionPlot[697].Y = 3.655

	pointsOfFunctionPlot[698].X = 1.98
	pointsOfFunctionPlot[698].Y = 3.69

	pointsOfFunctionPlot[699].X = 1.99
	pointsOfFunctionPlot[699].Y = 3.726

	pointsOfFunctionPlot[700].X = 2.0
	pointsOfFunctionPlot[700].Y = 3.762

	pointsOfFunctionPlot[701].X = 2.01
	pointsOfFunctionPlot[701].Y = 3.798

	pointsOfFunctionPlot[702].X = 2.02
	pointsOfFunctionPlot[702].Y = 3.835

	pointsOfFunctionPlot[703].X = 2.03
	pointsOfFunctionPlot[703].Y = 3.872

	pointsOfFunctionPlot[704].X = 2.04
	pointsOfFunctionPlot[704].Y = 3.91

	pointsOfFunctionPlot[705].X = 2.05
	pointsOfFunctionPlot[705].Y = 3.948

	pointsOfFunctionPlot[706].X = 2.06
	pointsOfFunctionPlot[706].Y = 3.986

	pointsOfFunctionPlot[707].X = 2.07
	pointsOfFunctionPlot[707].Y = 4.025

	pointsOfFunctionPlot[708].X = 2.08
	pointsOfFunctionPlot[708].Y = 4.064

	pointsOfFunctionPlot[709].X = 2.09
	pointsOfFunctionPlot[709].Y = 4.104

	pointsOfFunctionPlot[710].X = 2.10
	pointsOfFunctionPlot[710].Y = 4.144

	pointsOfFunctionPlot[711].X = 2.11
	pointsOfFunctionPlot[711].Y = 4.184

	pointsOfFunctionPlot[712].X = 2.12
	pointsOfFunctionPlot[712].Y = 4.225

	pointsOfFunctionPlot[713].X = 2.13
	pointsOfFunctionPlot[713].Y = 4.266

	pointsOfFunctionPlot[714].X = 2.14
	pointsOfFunctionPlot[714].Y = 4.308

	pointsOfFunctionPlot[715].X = 2.15
	pointsOfFunctionPlot[715].Y = 4.35

	pointsOfFunctionPlot[716].X = 2.16
	pointsOfFunctionPlot[716].Y = 4.393

	pointsOfFunctionPlot[717].X = 2.17
	pointsOfFunctionPlot[717].Y = 4.436

	pointsOfFunctionPlot[718].X = 2.18
	pointsOfFunctionPlot[718].Y = 4.479

	pointsOfFunctionPlot[719].X = 2.19
	pointsOfFunctionPlot[719].Y = 4.523

	pointsOfFunctionPlot[720].X = 2.20
	pointsOfFunctionPlot[720].Y = 4.567

	pointsOfFunctionPlot[721].X = 2.21
	pointsOfFunctionPlot[721].Y = 4.612

	pointsOfFunctionPlot[722].X = 2.22
	pointsOfFunctionPlot[722].Y = 4.657

	pointsOfFunctionPlot[723].X = 2.23
	pointsOfFunctionPlot[723].Y = 4.703

	pointsOfFunctionPlot[724].X = 2.24
	pointsOfFunctionPlot[724].Y = 4.749

	pointsOfFunctionPlot[725].X = 2.25
	pointsOfFunctionPlot[725].Y = 4.796

	pointsOfFunctionPlot[726].X = 2.26
	pointsOfFunctionPlot[726].Y = 4.843

	pointsOfFunctionPlot[727].X = 2.27
	pointsOfFunctionPlot[727].Y = 4.891

	pointsOfFunctionPlot[728].X = 2.28
	pointsOfFunctionPlot[728].Y = 4.939

	pointsOfFunctionPlot[729].X = 2.29
	pointsOfFunctionPlot[729].Y = 4.988

	pointsOfFunctionPlot[730].X = 2.30
	pointsOfFunctionPlot[730].Y = 5.037

	pointsOfFunctionPlot[731].X = 2.31
	pointsOfFunctionPlot[731].Y = 5.086

	pointsOfFunctionPlot[732].X = 2.32
	pointsOfFunctionPlot[732].Y = 5.136

	pointsOfFunctionPlot[733].X = 2.33
	pointsOfFunctionPlot[733].Y = 5.187

	pointsOfFunctionPlot[734].X = 2.34
	pointsOfFunctionPlot[734].Y = 5.238

	pointsOfFunctionPlot[735].X = 2.35
	pointsOfFunctionPlot[735].Y = 5.29

	pointsOfFunctionPlot[736].X = 2.36
	pointsOfFunctionPlot[736].Y = 5.342

	pointsOfFunctionPlot[737].X = 2.37
	pointsOfFunctionPlot[737].Y = 5.395

	pointsOfFunctionPlot[738].X = 2.38
	pointsOfFunctionPlot[738].Y = 5.448

	pointsOfFunctionPlot[739].X = 2.39
	pointsOfFunctionPlot[739].Y = 5.502

	pointsOfFunctionPlot[740].X = 2.40
	pointsOfFunctionPlot[740].Y = 5.556

	pointsOfFunctionPlot[741].X = 2.41
	pointsOfFunctionPlot[741].Y = 5.611

	pointsOfFunctionPlot[742].X = 2.42
	pointsOfFunctionPlot[742].Y = 5.667

	pointsOfFunctionPlot[743].X = 2.43
	pointsOfFunctionPlot[743].Y = 5.723

	pointsOfFunctionPlot[744].X = 2.44
	pointsOfFunctionPlot[744].Y = 5.78

	pointsOfFunctionPlot[745].X = 2.45
	pointsOfFunctionPlot[745].Y = 5.837

	pointsOfFunctionPlot[746].X = 2.46
	pointsOfFunctionPlot[746].Y = 5.895

	pointsOfFunctionPlot[747].X = 2.47
	pointsOfFunctionPlot[747].Y = 5.953

	pointsOfFunctionPlot[748].X = 2.48
	pointsOfFunctionPlot[748].Y = 6.012

	pointsOfFunctionPlot[749].X = 2.49
	pointsOfFunctionPlot[749].Y = 6.072

	pointsOfFunctionPlot[750].X = 2.50
	pointsOfFunctionPlot[750].Y = 6.132

	pointsOfFunctionPlot[751].X = 2.51
	pointsOfFunctionPlot[751].Y = 6.193

	pointsOfFunctionPlot[752].X = 2.52
	pointsOfFunctionPlot[752].Y = 6.254

	pointsOfFunctionPlot[753].X = 2.53
	pointsOfFunctionPlot[753].Y = 6.316

	pointsOfFunctionPlot[754].X = 2.54
	pointsOfFunctionPlot[754].Y = 6.379

	pointsOfFunctionPlot[755].X = 2.55
	pointsOfFunctionPlot[755].Y = 6.442

	pointsOfFunctionPlot[756].X = 2.56
	pointsOfFunctionPlot[756].Y = 6.506

	pointsOfFunctionPlot[757].X = 2.57
	pointsOfFunctionPlot[757].Y = 6.571

	pointsOfFunctionPlot[758].X = 2.58
	pointsOfFunctionPlot[758].Y = 6.636

	pointsOfFunctionPlot[759].X = 2.59
	pointsOfFunctionPlot[759].Y = 6.702

	pointsOfFunctionPlot[760].X = 2.60
	pointsOfFunctionPlot[760].Y = 6.769

	pointsOfFunctionPlot[761].X = 2.61
	pointsOfFunctionPlot[761].Y = 6.836

	pointsOfFunctionPlot[762].X = 2.62
	pointsOfFunctionPlot[762].Y = 6.904

	pointsOfFunctionPlot[763].X = 2.63
	pointsOfFunctionPlot[763].Y = 6.972

	pointsOfFunctionPlot[764].X = 2.64
	pointsOfFunctionPlot[764].Y = 7.042

	pointsOfFunctionPlot[765].X = 2.65
	pointsOfFunctionPlot[765].Y = 7.112

	pointsOfFunctionPlot[766].X = 2.66
	pointsOfFunctionPlot[766].Y = 7.183

	pointsOfFunctionPlot[767].X = 2.67
	pointsOfFunctionPlot[767].Y = 7.254

	pointsOfFunctionPlot[768].X = 2.68
	pointsOfFunctionPlot[768].Y = 7.326

	pointsOfFunctionPlot[769].X = 2.69
	pointsOfFunctionPlot[769].Y = 7.399

	pointsOfFunctionPlot[770].X = 2.70
	pointsOfFunctionPlot[770].Y = 7.473

	pointsOfFunctionPlot[771].X = 2.71
	pointsOfFunctionPlot[771].Y = 7.547

	pointsOfFunctionPlot[772].X = 2.72
	pointsOfFunctionPlot[772].Y = 7.623

	pointsOfFunctionPlot[773].X = 2.73
	pointsOfFunctionPlot[773].Y = 7.699

	pointsOfFunctionPlot[774].X = 2.74
	pointsOfFunctionPlot[774].Y = 7.775

	pointsOfFunctionPlot[775].X = 2.75
	pointsOfFunctionPlot[775].Y = 7.853

	pointsOfFunctionPlot[776].X = 2.76
	pointsOfFunctionPlot[776].Y = 7.931

	pointsOfFunctionPlot[777].X = 2.77
	pointsOfFunctionPlot[777].Y = 8.01

	pointsOfFunctionPlot[778].X = 2.78
	pointsOfFunctionPlot[778].Y = 8.09

	pointsOfFunctionPlot[779].X = 2.79
	pointsOfFunctionPlot[779].Y = 8.171

	pointsOfFunctionPlot[780].X = 2.80
	pointsOfFunctionPlot[780].Y = 8.252

	pointsOfFunctionPlot[781].X = 2.81
	pointsOfFunctionPlot[781].Y = 8.335

	pointsOfFunctionPlot[782].X = 2.82
	pointsOfFunctionPlot[782].Y = 8.418

	pointsOfFunctionPlot[783].X = 2.83
	pointsOfFunctionPlot[783].Y = 8.502

	pointsOfFunctionPlot[784].X = 2.84
	pointsOfFunctionPlot[784].Y = 8.587

	pointsOfFunctionPlot[785].X = 2.85
	pointsOfFunctionPlot[785].Y = 8.672

	pointsOfFunctionPlot[786].X = 2.86
	pointsOfFunctionPlot[786].Y = 8.759

	pointsOfFunctionPlot[787].X = 2.87
	pointsOfFunctionPlot[787].Y = 8.846

	pointsOfFunctionPlot[788].X = 2.88
	pointsOfFunctionPlot[788].Y = 8.935

	pointsOfFunctionPlot[789].X = 2.89
	pointsOfFunctionPlot[789].Y = 9.024

	pointsOfFunctionPlot[790].X = 2.90
	pointsOfFunctionPlot[790].Y = 9.114

	pointsOfFunctionPlot[791].X = 2.91
	pointsOfFunctionPlot[791].Y = 9.025

	pointsOfFunctionPlot[792].X = 2.92
	pointsOfFunctionPlot[792].Y = 9.297

	pointsOfFunctionPlot[793].X = 2.93
	pointsOfFunctionPlot[793].Y = 9.39

	pointsOfFunctionPlot[794].X = 2.94
	pointsOfFunctionPlot[794].Y = 9.484

	pointsOfFunctionPlot[795].X = 2.95
	pointsOfFunctionPlot[795].Y = 9.579

	pointsOfFunctionPlot[796].X = 2.96
	pointsOfFunctionPlot[796].Y = 9.674

	pointsOfFunctionPlot[797].X = 2.97
	pointsOfFunctionPlot[797].Y = 9.771

	pointsOfFunctionPlot[798].X = 2.98
	pointsOfFunctionPlot[798].Y = 9.869

	pointsOfFunctionPlot[799].X = 2.99
	pointsOfFunctionPlot[799].Y = 9.967

	pointsOfFunctionPlot[800].X = 3.0
	pointsOfFunctionPlot[800].Y = 10.067

	pointsOfFunctionPlot[801].X = 3.01
	pointsOfFunctionPlot[801].Y = 10.168

	pointsOfFunctionPlot[802].X = 3.02
	pointsOfFunctionPlot[802].Y = 10.27

	pointsOfFunctionPlot[803].X = 3.03
	pointsOfFunctionPlot[803].Y = 10.372

	pointsOfFunctionPlot[804].X = 3.04
	pointsOfFunctionPlot[804].Y = 10.476

	pointsOfFunctionPlot[805].X = 3.05
	pointsOfFunctionPlot[805].Y = 10.581

	pointsOfFunctionPlot[806].X = 3.06
	pointsOfFunctionPlot[806].Y = 10.687

	pointsOfFunctionPlot[807].X = 3.07
	pointsOfFunctionPlot[807].Y = 10.794

	pointsOfFunctionPlot[808].X = 3.08
	pointsOfFunctionPlot[808].Y = 10.902

	pointsOfFunctionPlot[809].X = 3.09
	pointsOfFunctionPlot[809].Y = 11.011

	pointsOfFunctionPlot[810].X = 3.10
	pointsOfFunctionPlot[810].Y = 11.121

	pointsOfFunctionPlot[811].X = 3.11
	pointsOfFunctionPlot[811].Y = 11.232

	pointsOfFunctionPlot[812].X = 3.12
	pointsOfFunctionPlot[812].Y = 11.345

	pointsOfFunctionPlot[813].X = 3.13
	pointsOfFunctionPlot[813].Y = 11.458

	pointsOfFunctionPlot[814].X = 3.14
	pointsOfFunctionPlot[814].Y = 11.573

	pointsOfFunctionPlot[815].X = 3.15
	pointsOfFunctionPlot[815].Y = 11.689

	pointsOfFunctionPlot[816].X = 3.16
	pointsOfFunctionPlot[816].Y = 11.806

	pointsOfFunctionPlot[817].X = 3.17
	pointsOfFunctionPlot[817].Y = 11.924

	pointsOfFunctionPlot[818].X = 3.18
	pointsOfFunctionPlot[818].Y = 12.044

	pointsOfFunctionPlot[819].X = 3.19
	pointsOfFunctionPlot[819].Y = 12.164

	pointsOfFunctionPlot[820].X = 3.20
	pointsOfFunctionPlot[820].Y = 12.286

	pointsOfFunctionPlot[821].X = 3.21
	pointsOfFunctionPlot[821].Y = 12.409

	pointsOfFunctionPlot[822].X = 3.22
	pointsOfFunctionPlot[822].Y = 12.53

	pointsOfFunctionPlot[823].X = 3.23
	pointsOfFunctionPlot[823].Y = 12.659

	pointsOfFunctionPlot[824].X = 3.24
	pointsOfFunctionPlot[824].Y = 12.786

	pointsOfFunctionPlot[825].X = 3.25
	pointsOfFunctionPlot[825].Y = 12.914

	pointsOfFunctionPlot[826].X = 3.26
	pointsOfFunctionPlot[826].Y = 13.043

	pointsOfFunctionPlot[827].X = 3.27
	pointsOfFunctionPlot[827].Y = 13.174

	pointsOfFunctionPlot[828].X = 3.28
	pointsOfFunctionPlot[828].Y = 13.306

	pointsOfFunctionPlot[829].X = 3.29
	pointsOfFunctionPlot[829].Y = 13.44

	pointsOfFunctionPlot[830].X = 3.30
	pointsOfFunctionPlot[830].Y = 13.574

	pointsOfFunctionPlot[831].X = 3.31
	pointsOfFunctionPlot[831].Y = 13.71

	pointsOfFunctionPlot[832].X = 3.32
	pointsOfFunctionPlot[832].Y = 13.848

	pointsOfFunctionPlot[833].X = 3.33
	pointsOfFunctionPlot[833].Y = 13.987

	pointsOfFunctionPlot[834].X = 3.34
	pointsOfFunctionPlot[834].Y = 14.127

	pointsOfFunctionPlot[835].X = 3.35
	pointsOfFunctionPlot[835].Y = 14.268

	pointsOfFunctionPlot[836].X = 3.36
	pointsOfFunctionPlot[836].Y = 14.411

	pointsOfFunctionPlot[837].X = 3.37
	pointsOfFunctionPlot[837].Y = 14.556

	pointsOfFunctionPlot[838].X = 3.38
	pointsOfFunctionPlot[838].Y = 14.702

	pointsOfFunctionPlot[839].X = 3.39
	pointsOfFunctionPlot[839].Y = 14.849

	pointsOfFunctionPlot[840].X = 3.40
	pointsOfFunctionPlot[840].Y = 14.998

	pointsOfFunctionPlot[841].X = 3.41
	pointsOfFunctionPlot[841].Y = 15.149

	pointsOfFunctionPlot[842].X = 3.42
	pointsOfFunctionPlot[842].Y = 15.301

	pointsOfFunctionPlot[843].X = 3.43
	pointsOfFunctionPlot[843].Y = 15.454

	pointsOfFunctionPlot[844].X = 3.44
	pointsOfFunctionPlot[844].Y = 15.609

	pointsOfFunctionPlot[845].X = 3.45
	pointsOfFunctionPlot[845].Y = 15.766

	pointsOfFunctionPlot[846].X = 3.46
	pointsOfFunctionPlot[846].Y = 15.924

	pointsOfFunctionPlot[847].X = 3.47
	pointsOfFunctionPlot[847].Y = 16.083

	pointsOfFunctionPlot[848].X = 3.48
	pointsOfFunctionPlot[848].Y = 16.245

	pointsOfFunctionPlot[849].X = 3.49
	pointsOfFunctionPlot[849].Y = 16.408

	pointsOfFunctionPlot[850].X = 3.50
	pointsOfFunctionPlot[850].Y = 16.572

	pointsOfFunctionPlot[851].X = 3.51
	pointsOfFunctionPlot[851].Y = 16.739

	pointsOfFunctionPlot[852].X = 3.52
	pointsOfFunctionPlot[852].Y = 16.907

	pointsOfFunctionPlot[853].X = 3.53
	pointsOfFunctionPlot[853].Y = 17.076

	pointsOfFunctionPlot[854].X = 3.54
	pointsOfFunctionPlot[854].Y = 17.247

	pointsOfFunctionPlot[855].X = 3.55
	pointsOfFunctionPlot[855].Y = 17.421

	pointsOfFunctionPlot[856].X = 3.56
	pointsOfFunctionPlot[856].Y = 17.595

	pointsOfFunctionPlot[857].X = 3.57
	pointsOfFunctionPlot[857].Y = 17.772

	pointsOfFunctionPlot[858].X = 3.58
	pointsOfFunctionPlot[858].Y = 17.95

	pointsOfFunctionPlot[859].X = 3.59
	pointsOfFunctionPlot[859].Y = 18.13

	pointsOfFunctionPlot[860].X = 3.60
	pointsOfFunctionPlot[860].Y = 18.312

	pointsOfFunctionPlot[861].X = 3.61
	pointsOfFunctionPlot[861].Y = 18.496

	pointsOfFunctionPlot[862].X = 3.62
	pointsOfFunctionPlot[862].Y = 18.682

	pointsOfFunctionPlot[863].X = 3.63
	pointsOfFunctionPlot[863].Y = 18.869

	pointsOfFunctionPlot[864].X = 3.64
	pointsOfFunctionPlot[864].Y = 19.059

	pointsOfFunctionPlot[865].X = 3.65
	pointsOfFunctionPlot[865].Y = 19.25

	pointsOfFunctionPlot[866].X = 3.66
	pointsOfFunctionPlot[866].Y = 19.443

	pointsOfFunctionPlot[867].X = 3.67
	pointsOfFunctionPlot[867].Y = 19.638

	pointsOfFunctionPlot[868].X = 3.68
	pointsOfFunctionPlot[868].Y = 19.835

	pointsOfFunctionPlot[869].X = 3.69
	pointsOfFunctionPlot[869].Y = 20.034

	pointsOfFunctionPlot[870].X = 3.70
	pointsOfFunctionPlot[870].Y = 20.236

	pointsOfFunctionPlot[871].X = 3.71
	pointsOfFunctionPlot[871].Y = 20.439

	pointsOfFunctionPlot[872].X = 3.72
	pointsOfFunctionPlot[872].Y = 20.644

	pointsOfFunctionPlot[873].X = 3.73
	pointsOfFunctionPlot[873].Y = 20.851

	pointsOfFunctionPlot[874].X = 3.74
	pointsOfFunctionPlot[874].Y = 21.06

	pointsOfFunctionPlot[875].X = 3.75
	pointsOfFunctionPlot[875].Y = 21.272

	pointsOfFunctionPlot[876].X = 3.76
	pointsOfFunctionPlot[876].Y = 21.485

	pointsOfFunctionPlot[877].X = 3.77
	pointsOfFunctionPlot[877].Y = 21.701

	pointsOfFunctionPlot[878].X = 3.78
	pointsOfFunctionPlot[878].Y = 21.919

	pointsOfFunctionPlot[879].X = 3.79
	pointsOfFunctionPlot[879].Y = 22.139

	pointsOfFunctionPlot[880].X = 3.80
	pointsOfFunctionPlot[880].Y = 22.361

	pointsOfFunctionPlot[881].X = 3.81
	pointsOfFunctionPlot[881].Y = 22.586

	pointsOfFunctionPlot[882].X = 3.82
	pointsOfFunctionPlot[882].Y = 22.813

	pointsOfFunctionPlot[883].X = 3.83
	pointsOfFunctionPlot[883].Y = 23.042

	pointsOfFunctionPlot[884].X = 3.84
	pointsOfFunctionPlot[884].Y = 23.273

	pointsOfFunctionPlot[885].X = 3.85
	pointsOfFunctionPlot[885].Y = 23.507

	pointsOfFunctionPlot[886].X = 3.86
	pointsOfFunctionPlot[886].Y = 23.743

	pointsOfFunctionPlot[887].X = 3.87
	pointsOfFunctionPlot[887].Y = 23.981

	pointsOfFunctionPlot[888].X = 3.88
	pointsOfFunctionPlot[888].Y = 24.222

	pointsOfFunctionPlot[889].X = 3.89
	pointsOfFunctionPlot[889].Y = 24.465

	pointsOfFunctionPlot[890].X = 3.90
	pointsOfFunctionPlot[890].Y = 24.711

	pointsOfFunctionPlot[891].X = 3.91
	pointsOfFunctionPlot[891].Y = 24.959

	pointsOfFunctionPlot[892].X = 3.92
	pointsOfFunctionPlot[892].Y = 25.21

	pointsOfFunctionPlot[893].X = 3.93
	pointsOfFunctionPlot[893].Y = 25.463

	pointsOfFunctionPlot[894].X = 3.94
	pointsOfFunctionPlot[894].Y = 25.719

	pointsOfFunctionPlot[895].X = 3.95
	pointsOfFunctionPlot[895].Y = 25.977

	pointsOfFunctionPlot[896].X = 3.96
	pointsOfFunctionPlot[896].Y = 26.238

	pointsOfFunctionPlot[897].X = 3.97
	pointsOfFunctionPlot[897].Y = 26.501

	pointsOfFunctionPlot[898].X = 3.98
	pointsOfFunctionPlot[898].Y = 26.767

	pointsOfFunctionPlot[899].X = 3.99
	pointsOfFunctionPlot[899].Y = 27.036

	pointsOfFunctionPlot[900].X = 4.0
	pointsOfFunctionPlot[900].Y = 27.308

	pointsOfFunctionPlot[901].X = 4.01
	pointsOfFunctionPlot[901].Y = 27.582

	pointsOfFunctionPlot[902].X = 4.02
	pointsOfFunctionPlot[902].Y = 27.859

	pointsOfFunctionPlot[903].X = 4.03
	pointsOfFunctionPlot[903].Y = 28.139

	pointsOfFunctionPlot[904].X = 4.04
	pointsOfFunctionPlot[904].Y = 28.421

	pointsOfFunctionPlot[905].X = 4.05
	pointsOfFunctionPlot[905].Y = 28.707

	pointsOfFunctionPlot[906].X = 4.06
	pointsOfFunctionPlot[906].Y = 28.995

	pointsOfFunctionPlot[907].X = 4.07
	pointsOfFunctionPlot[907].Y = 29.287

	pointsOfFunctionPlot[908].X = 4.08
	pointsOfFunctionPlot[908].Y = 29.581

	pointsOfFunctionPlot[909].X = 4.09
	pointsOfFunctionPlot[909].Y = 29.878

	pointsOfFunctionPlot[910].X = 4.10
	pointsOfFunctionPlot[910].Y = 30.178

	pointsOfFunctionPlot[911].X = 4.11
	pointsOfFunctionPlot[911].Y = 30.481

	pointsOfFunctionPlot[912].X = 4.12
	pointsOfFunctionPlot[912].Y = 30.787

	pointsOfFunctionPlot[913].X = 4.13
	pointsOfFunctionPlot[913].Y = 31.097

	pointsOfFunctionPlot[914].X = 4.14
	pointsOfFunctionPlot[914].Y = 31.409

	pointsOfFunctionPlot[915].X = 4.15
	pointsOfFunctionPlot[915].Y = 31.724

	pointsOfFunctionPlot[916].X = 4.16
	pointsOfFunctionPlot[916].Y = 32.043

	pointsOfFunctionPlot[917].X = 4.17
	pointsOfFunctionPlot[917].Y = 32.365

	pointsOfFunctionPlot[918].X = 4.18
	pointsOfFunctionPlot[918].Y = 32.69

	pointsOfFunctionPlot[919].X = 4.19
	pointsOfFunctionPlot[919].Y = 33.018

	pointsOfFunctionPlot[920].X = 4.20
	pointsOfFunctionPlot[920].Y = 33.35

	pointsOfFunctionPlot[921].X = 4.21
	pointsOfFunctionPlot[921].Y = 33.685

	pointsOfFunctionPlot[922].X = 4.22
	pointsOfFunctionPlot[922].Y = 34.024

	pointsOfFunctionPlot[923].X = 4.23
	pointsOfFunctionPlot[923].Y = 34.365

	pointsOfFunctionPlot[924].X = 4.24
	pointsOfFunctionPlot[924].Y = 34.711

	pointsOfFunctionPlot[925].X = 4.25
	pointsOfFunctionPlot[925].Y = 35.059

	pointsOfFunctionPlot[926].X = 4.26
	pointsOfFunctionPlot[926].Y = 35.412

	pointsOfFunctionPlot[927].X = 4.27
	pointsOfFunctionPlot[927].Y = 35.767

	pointsOfFunctionPlot[928].X = 4.28
	pointsOfFunctionPlot[928].Y = 36.127

	pointsOfFunctionPlot[929].X = 4.29
	pointsOfFunctionPlot[929].Y = 36.49

	pointsOfFunctionPlot[930].X = 4.30
	pointsOfFunctionPlot[930].Y = 36.856

	pointsOfFunctionPlot[931].X = 4.31
	pointsOfFunctionPlot[931].Y = 37.226

	pointsOfFunctionPlot[932].X = 4.32
	pointsOfFunctionPlot[932].Y = 37.6

	pointsOfFunctionPlot[933].X = 4.33
	pointsOfFunctionPlot[933].Y = 37.978

	pointsOfFunctionPlot[934].X = 4.34
	pointsOfFunctionPlot[934].Y = 38.36

	pointsOfFunctionPlot[935].X = 4.35
	pointsOfFunctionPlot[935].Y = 38.745

	pointsOfFunctionPlot[936].X = 4.36
	pointsOfFunctionPlot[936].Y = 39.134

	pointsOfFunctionPlot[937].X = 4.37
	pointsOfFunctionPlot[937].Y = 39.528

	pointsOfFunctionPlot[938].X = 4.38
	pointsOfFunctionPlot[938].Y = 39.925

	pointsOfFunctionPlot[939].X = 4.39
	pointsOfFunctionPlot[939].Y = 40.326

	pointsOfFunctionPlot[940].X = 4.40
	pointsOfFunctionPlot[940].Y = 40.731

	pointsOfFunctionPlot[941].X = 4.41
	pointsOfFunctionPlot[941].Y = 41.14

	pointsOfFunctionPlot[942].X = 4.42
	pointsOfFunctionPlot[942].Y = 41.554

	pointsOfFunctionPlot[943].X = 4.43
	pointsOfFunctionPlot[943].Y = 41.971

	pointsOfFunctionPlot[944].X = 4.44
	pointsOfFunctionPlot[944].Y = 42.393

	pointsOfFunctionPlot[945].X = 4.45
	pointsOfFunctionPlot[945].Y = 42.819

	pointsOfFunctionPlot[946].X = 4.46
	pointsOfFunctionPlot[946].Y = 43.249

	pointsOfFunctionPlot[947].X = 4.47
	pointsOfFunctionPlot[947].Y = 43.684

	pointsOfFunctionPlot[948].X = 4.48
	pointsOfFunctionPlot[948].Y = 44.123

	pointsOfFunctionPlot[949].X = 4.49
	pointsOfFunctionPlot[949].Y = 44.566

	pointsOfFunctionPlot[950].X = 4.50
	pointsOfFunctionPlot[950].Y = 45.014

	pointsOfFunctionPlot[951].X = 4.51
	pointsOfFunctionPlot[951].Y = 45.466

	pointsOfFunctionPlot[952].X = 4.52
	pointsOfFunctionPlot[952].Y = 45.923

	pointsOfFunctionPlot[953].X = 4.53
	pointsOfFunctionPlot[953].Y = 46.384

	pointsOfFunctionPlot[954].X = 4.54
	pointsOfFunctionPlot[954].Y = 46.85

	pointsOfFunctionPlot[955].X = 4.55
	pointsOfFunctionPlot[955].Y = 47.321

	pointsOfFunctionPlot[956].X = 4.56
	pointsOfFunctionPlot[956].Y = 47.796

	pointsOfFunctionPlot[957].X = 4.57
	pointsOfFunctionPlot[957].Y = 48.277

	pointsOfFunctionPlot[958].X = 4.58
	pointsOfFunctionPlot[958].Y = 48.762

	pointsOfFunctionPlot[959].X = 4.59
	pointsOfFunctionPlot[959].Y = 49.252

	pointsOfFunctionPlot[960].X = 4.60
	pointsOfFunctionPlot[960].Y = 49.747

	pointsOfFunctionPlot[961].X = 4.61
	pointsOfFunctionPlot[961].Y = 50.247

	pointsOfFunctionPlot[962].X = 4.62
	pointsOfFunctionPlot[962].Y = 50.751

	pointsOfFunctionPlot[963].X = 4.63
	pointsOfFunctionPlot[963].Y = 51.261

	pointsOfFunctionPlot[964].X = 4.64
	pointsOfFunctionPlot[964].Y = 51.77

	pointsOfFunctionPlot[965].X = 4.65
	pointsOfFunctionPlot[965].Y = 52.297

	pointsOfFunctionPlot[966].X = 4.66
	pointsOfFunctionPlot[966].Y = 52.822

	pointsOfFunctionPlot[967].X = 4.67
	pointsOfFunctionPlot[967].Y = 53.353

	pointsOfFunctionPlot[968].X = 4.68
	pointsOfFunctionPlot[968].Y = 53.889

	pointsOfFunctionPlot[969].X = 4.69
	pointsOfFunctionPlot[969].Y = 54.431

	pointsOfFunctionPlot[970].X = 4.70
	pointsOfFunctionPlot[970].Y = 54.978

	pointsOfFunctionPlot[971].X = 4.71
	pointsOfFunctionPlot[971].Y = 55.53

	pointsOfFunctionPlot[972].X = 4.72
	pointsOfFunctionPlot[972].Y = 56.088

	pointsOfFunctionPlot[973].X = 4.73
	pointsOfFunctionPlot[973].Y = 56.652

	pointsOfFunctionPlot[974].X = 4.74
	pointsOfFunctionPlot[974].Y = 57.221

	pointsOfFunctionPlot[975].X = 4.75
	pointsOfFunctionPlot[975].Y = 57.796

	pointsOfFunctionPlot[976].X = 4.76
	pointsOfFunctionPlot[976].Y = 58.377

	pointsOfFunctionPlot[977].X = 4.77
	pointsOfFunctionPlot[977].Y = 58.963

	pointsOfFunctionPlot[978].X = 4.78
	pointsOfFunctionPlot[978].Y = 59.556

	pointsOfFunctionPlot[979].X = 4.79
	pointsOfFunctionPlot[979].Y = 60.154

	pointsOfFunctionPlot[980].X = 4.80
	pointsOfFunctionPlot[980].Y = 60.759

	pointsOfFunctionPlot[981].X = 4.81
	pointsOfFunctionPlot[981].Y = 61.369

	pointsOfFunctionPlot[982].X = 4.82
	pointsOfFunctionPlot[982].Y = 61.986

	pointsOfFunctionPlot[983].X = 4.83
	pointsOfFunctionPlot[983].Y = 62.609

	pointsOfFunctionPlot[984].X = 4.84
	pointsOfFunctionPlot[984].Y = 63.238

	pointsOfFunctionPlot[985].X = 4.85
	pointsOfFunctionPlot[985].Y = 63.874

	pointsOfFunctionPlot[986].X = 4.86
	pointsOfFunctionPlot[986].Y = 64.515

	pointsOfFunctionPlot[987].X = 4.87
	pointsOfFunctionPlot[987].Y = 65.164

	pointsOfFunctionPlot[988].X = 4.88
	pointsOfFunctionPlot[988].Y = 65.819

	pointsOfFunctionPlot[989].X = 4.89
	pointsOfFunctionPlot[989].Y = 66.48

	pointsOfFunctionPlot[990].X = 4.90
	pointsOfFunctionPlot[990].Y = 67.148

	pointsOfFunctionPlot[991].X = 4.91
	pointsOfFunctionPlot[991].Y = 67.823

	pointsOfFunctionPlot[992].X = 4.92
	pointsOfFunctionPlot[992].Y = 68.504

	pointsOfFunctionPlot[993].X = 4.93
	pointsOfFunctionPlot[993].Y = 69.193

	pointsOfFunctionPlot[994].X = 4.94
	pointsOfFunctionPlot[994].Y = 69.888

	pointsOfFunctionPlot[995].X = 4.95
	pointsOfFunctionPlot[995].Y = 70.591

	pointsOfFunctionPlot[996].X = 4.96
	pointsOfFunctionPlot[996].Y = 71.3

	pointsOfFunctionPlot[997].X = 4.97
	pointsOfFunctionPlot[997].Y = 72.016

	pointsOfFunctionPlot[998].X = 4.98
	pointsOfFunctionPlot[998].Y = 72.74

	pointsOfFunctionPlot[999].X = 4.99
	pointsOfFunctionPlot[999].Y = 73.471

	pointsOfFunctionPlot[1_000].X = 5.0
	pointsOfFunctionPlot[1_000].Y = 74.209










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function cosh(x)"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("cosh(x)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"cosh-function-plot-01.png"); err != nil {

		panic(err)
	}
}
