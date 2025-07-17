package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres wielomianu f(x) = x^4 - 20x^3 + 33.75x^2 - 235x + 89.0625

	pointsOfPolynomialPlot := make(plotter.XYs, 1_001)

	pointsOfPolynomialPlot[0].X = 0.0
	pointsOfPolynomialPlot[0].Y = 89.0625

	pointsOfPolynomialPlot[1].X = 0.01
	pointsOfPolynomialPlot[1].Y = 86.715

	pointsOfPolynomialPlot[2].X = 0.02
	pointsOfPolynomialPlot[2].Y = 84.375

	pointsOfPolynomialPlot[3].X = 0.03
	pointsOfPolynomialPlot[3].Y = 82.0423

	pointsOfPolynomialPlot[4].X = 0.04
	pointsOfPolynomialPlot[4].Y = 79.715

	pointsOfPolynomialPlot[5].X = 0.05
	pointsOfPolynomialPlot[5].Y = 77.394

	pointsOfPolynomialPlot[6].X = 0.06
	pointsOfPolynomialPlot[6].Y = 75.079

	pointsOfPolynomialPlot[7].X = 0.07
	pointsOfPolynomialPlot[7].Y = 72.771

	pointsOfPolynomialPlot[8].X = 0.08
	pointsOfPolynomialPlot[8].Y = 70.468

	pointsOfPolynomialPlot[9].X = 0.09
	pointsOfPolynomialPlot[9].Y = 68.171

	pointsOfPolynomialPlot[10].X = 0.1
	pointsOfPolynomialPlot[10].Y = 65.88

	pointsOfPolynomialPlot[11].X = 0.11
	pointsOfPolynomialPlot[11].Y = 63.594

	pointsOfPolynomialPlot[12].X = 0.12
	pointsOfPolynomialPlot[12].Y = 61.314

	pointsOfPolynomialPlot[13].X = 0.13
	pointsOfPolynomialPlot[13].Y = 59.039

	pointsOfPolynomialPlot[14].X = 0.14
	pointsOfPolynomialPlot[14].Y = 56.769

	pointsOfPolynomialPlot[15].X = 0.15
	pointsOfPolynomialPlot[15].Y = 54.504

	pointsOfPolynomialPlot[16].X = 0.16
	pointsOfPolynomialPlot[16].Y = 52.245

	pointsOfPolynomialPlot[17].X = 0.17
	pointsOfPolynomialPlot[17].Y = 49.99

	pointsOfPolynomialPlot[18].X = 0.18
	pointsOfPolynomialPlot[18].Y = 47.74

	pointsOfPolynomialPlot[19].X = 0.19
	pointsOfPolynomialPlot[19].Y = 45.495

	pointsOfPolynomialPlot[20].X = 0.2
	pointsOfPolynomialPlot[20].Y = 43.254

	pointsOfPolynomialPlot[21].X = 0.21
	pointsOfPolynomialPlot[21].Y = 41.017

	pointsOfPolynomialPlot[22].X = 0.22
	pointsOfPolynomialPlot[22].Y = 38.785

	pointsOfPolynomialPlot[23].X = 0.23
	pointsOfPolynomialPlot[23].Y = 36.557

	pointsOfPolynomialPlot[24].X = 0.24
	pointsOfPolynomialPlot[24].Y = 34.333

	pointsOfPolynomialPlot[25].X = 0.25
	pointsOfPolynomialPlot[25].Y = 32.113

	pointsOfPolynomialPlot[26].X = 0.26
	pointsOfPolynomialPlot[26].Y = 29.897

	pointsOfPolynomialPlot[27].X = 0.27
	pointsOfPolynomialPlot[27].Y = 27.684

	pointsOfPolynomialPlot[28].X = 0.28
	pointsOfPolynomialPlot[28].Y = 25.475

	pointsOfPolynomialPlot[29].X = 0.29
	pointsOfPolynomialPlot[29].Y = 23.27

	pointsOfPolynomialPlot[30].X = 0.3
	pointsOfPolynomialPlot[30].Y = 21.068

	pointsOfPolynomialPlot[31].X = 0.31
	pointsOfPolynomialPlot[31].Y = 18.869

	pointsOfPolynomialPlot[32].X = 0.32
	pointsOfPolynomialPlot[32].Y = 16.673

	pointsOfPolynomialPlot[33].X = 0.33
	pointsOfPolynomialPlot[33].Y = 14.481

	pointsOfPolynomialPlot[34].X = 0.34
	pointsOfPolynomialPlot[34].Y = 12.291

	pointsOfPolynomialPlot[35].X = 0.35
	pointsOfPolynomialPlot[35].Y = 10.104

	pointsOfPolynomialPlot[36].X = 0.36
	pointsOfPolynomialPlot[36].Y = 7.92

	pointsOfPolynomialPlot[37].X = 0.37
	pointsOfPolynomialPlot[37].Y = 5.738

	pointsOfPolynomialPlot[38].X = 0.38
	pointsOfPolynomialPlot[38].Y = 3.559

	pointsOfPolynomialPlot[39].X = 0.39
	pointsOfPolynomialPlot[39].Y = 1.382

	pointsOfPolynomialPlot[40].X = 0.4
	pointsOfPolynomialPlot[40].Y = -0.791

	pointsOfPolynomialPlot[41].X = 0.41
	pointsOfPolynomialPlot[41].Y = -2.964

	pointsOfPolynomialPlot[42].X = 0.42
	pointsOfPolynomialPlot[42].Y = -5.134

	pointsOfPolynomialPlot[43].X = 0.43
	pointsOfPolynomialPlot[43].Y = -7.303

	pointsOfPolynomialPlot[44].X = 0.44
	pointsOfPolynomialPlot[44].Y = -9.469

	pointsOfPolynomialPlot[45].X = 0.45
	pointsOfPolynomialPlot[45].Y = -11.634

	pointsOfPolynomialPlot[46].X = 0.46
	pointsOfPolynomialPlot[46].Y = -13.797

	pointsOfPolynomialPlot[47].X = 0.47
	pointsOfPolynomialPlot[47].Y = -15.959

	pointsOfPolynomialPlot[48].X = 0.48
	pointsOfPolynomialPlot[48].Y = -18.12

	pointsOfPolynomialPlot[49].X = 0.49
	pointsOfPolynomialPlot[49].Y = -20.279

	pointsOfPolynomialPlot[50].X = 0.5
	pointsOfPolynomialPlot[50].Y = -22.437

	pointsOfPolynomialPlot[51].X = 0.51
	pointsOfPolynomialPlot[51].Y = -24.594

	pointsOfPolynomialPlot[52].X = 0.52
	pointsOfPolynomialPlot[52].Y = -26.75

	pointsOfPolynomialPlot[53].X = 0.53
	pointsOfPolynomialPlot[53].Y = -28.905

	pointsOfPolynomialPlot[54].X = 0.54
	pointsOfPolynomialPlot[54].Y = -31.06

	pointsOfPolynomialPlot[55].X = 0.55
	pointsOfPolynomialPlot[55].Y = -33.214

	pointsOfPolynomialPlot[56].X = 0.56
	pointsOfPolynomialPlot[56].Y = -35.367

	pointsOfPolynomialPlot[57].X = 0.57
	pointsOfPolynomialPlot[57].Y = -37.52

	pointsOfPolynomialPlot[58].X = 0.58
	pointsOfPolynomialPlot[58].Y = -39.673

	pointsOfPolynomialPlot[59].X = 0.59
	pointsOfPolynomialPlot[59].Y = -41.825

	pointsOfPolynomialPlot[60].X = 0.6
	pointsOfPolynomialPlot[60].Y = -43.977

	pointsOfPolynomialPlot[61].X = 0.61
	pointsOfPolynomialPlot[61].Y = -46.13

	pointsOfPolynomialPlot[62].X = 0.62
	pointsOfPolynomialPlot[62].Y = -48.282

	pointsOfPolynomialPlot[63].X = 0.63
	pointsOfPolynomialPlot[63].Y = -50.435

	pointsOfPolynomialPlot[64].X = 0.64
	pointsOfPolynomialPlot[64].Y = -52.588

	pointsOfPolynomialPlot[65].X = 0.65
	pointsOfPolynomialPlot[65].Y = -54.742

	pointsOfPolynomialPlot[66].X = 0.66
	pointsOfPolynomialPlot[66].Y = -56.896

	pointsOfPolynomialPlot[67].X = 0.67
	pointsOfPolynomialPlot[67].Y = -59.05

	pointsOfPolynomialPlot[68].X = 0.68
	pointsOfPolynomialPlot[68].Y = -61.206

	pointsOfPolynomialPlot[69].X = 0.69
	pointsOfPolynomialPlot[69].Y = -63.362

	pointsOfPolynomialPlot[70].X = 0.7
	pointsOfPolynomialPlot[70].Y = -65.519

	pointsOfPolynomialPlot[71].X = 0.71
	pointsOfPolynomialPlot[71].Y = -67.678

	pointsOfPolynomialPlot[72].X = 0.72
	pointsOfPolynomialPlot[72].Y = -69.837

	pointsOfPolynomialPlot[73].X = 0.73
	pointsOfPolynomialPlot[73].Y = -71.998

	pointsOfPolynomialPlot[74].X = 0.74
	pointsOfPolynomialPlot[74].Y = -74.16

	pointsOfPolynomialPlot[75].X = 0.75
	pointsOfPolynomialPlot[75].Y = -76.324

	pointsOfPolynomialPlot[76].X = 0.76
	pointsOfPolynomialPlot[76].Y = -78.489

	pointsOfPolynomialPlot[77].X = 0.77
	pointsOfPolynomialPlot[77].Y = -80.656

	pointsOfPolynomialPlot[78].X = 0.78
	pointsOfPolynomialPlot[78].Y = -82.824

	pointsOfPolynomialPlot[79].X = 0.79
	pointsOfPolynomialPlot[79].Y = -84.995

	pointsOfPolynomialPlot[80].X = 0.8
	pointsOfPolynomialPlot[80].Y = -87.167

	pointsOfPolynomialPlot[81].X = 0.81
	pointsOfPolynomialPlot[81].Y = -89.342

	pointsOfPolynomialPlot[82].X = 0.82
	pointsOfPolynomialPlot[82].Y = -91.519

	pointsOfPolynomialPlot[83].X = 0.83
	pointsOfPolynomialPlot[83].Y = -93.698

	pointsOfPolynomialPlot[84].X = 0.84
	pointsOfPolynomialPlot[84].Y = -95.879

	pointsOfPolynomialPlot[85].X = 0.85
	pointsOfPolynomialPlot[85].Y = -98.063

	pointsOfPolynomialPlot[86].X = 0.86
	pointsOfPolynomialPlot[86].Y = -100.25

	pointsOfPolynomialPlot[87].X = 0.87
	pointsOfPolynomialPlot[87].Y = -102.439

	pointsOfPolynomialPlot[88].X = 0.88
	pointsOfPolynomialPlot[88].Y = -104.631

	pointsOfPolynomialPlot[89].X = 0.89
	pointsOfPolynomialPlot[89].Y = -106.826

	pointsOfPolynomialPlot[90].X = 0.9
	pointsOfPolynomialPlot[90].Y = -109.023

	pointsOfPolynomialPlot[91].X = 0.91
	pointsOfPolynomialPlot[91].Y = -111.224

	pointsOfPolynomialPlot[92].X = 0.92
	pointsOfPolynomialPlot[92].Y = -113.428

	pointsOfPolynomialPlot[93].X = 0.93
	pointsOfPolynomialPlot[93].Y = -115.636

	pointsOfPolynomialPlot[94].X = 0.94
	pointsOfPolynomialPlot[94].Y = -117.846

	pointsOfPolynomialPlot[95].X = 0.95
	pointsOfPolynomialPlot[95].Y = -120.061

	pointsOfPolynomialPlot[96].X = 0.96
	pointsOfPolynomialPlot[96].Y = -122.278

	pointsOfPolynomialPlot[97].X = 0.97
	pointsOfPolynomialPlot[97].Y = -124.5

	pointsOfPolynomialPlot[98].X = 0.98
	pointsOfPolynomialPlot[98].Y = -126.725

	pointsOfPolynomialPlot[99].X = 0.99
	pointsOfPolynomialPlot[99].Y = -128.954

	pointsOfPolynomialPlot[100].X = 1.0
	pointsOfPolynomialPlot[100].Y = -131.187

	pointsOfPolynomialPlot[101].X = 1.01
	pointsOfPolynomialPlot[101].Y = -133.424

	pointsOfPolynomialPlot[102].X = 1.02
	pointsOfPolynomialPlot[102].Y = -135.665

	pointsOfPolynomialPlot[103].X = 1.03
	pointsOfPolynomialPlot[103].Y = -137.911

	pointsOfPolynomialPlot[104].X = 1.04
	pointsOfPolynomialPlot[104].Y = -140.16

	pointsOfPolynomialPlot[105].X = 1.05
	pointsOfPolynomialPlot[105].Y = -142.415

	pointsOfPolynomialPlot[106].X = 1.06
	pointsOfPolynomialPlot[106].Y = -144.673

	pointsOfPolynomialPlot[107].X = 1.07
	pointsOfPolynomialPlot[107].Y = -146.937

	pointsOfPolynomialPlot[108].X = 1.08
	pointsOfPolynomialPlot[108].Y = -149.205

	pointsOfPolynomialPlot[109].X = 1.09
	pointsOfPolynomialPlot[109].Y = -151.478

	pointsOfPolynomialPlot[110].X = 1.1
	pointsOfPolynomialPlot[110].Y = -153.755

	pointsOfPolynomialPlot[111].X = 1.11
	pointsOfPolynomialPlot[111].Y = -156.038

	pointsOfPolynomialPlot[112].X = 1.12
	pointsOfPolynomialPlot[112].Y = -158.326

	pointsOfPolynomialPlot[113].X = 1.13
	pointsOfPolynomialPlot[113].Y = -160.619

	pointsOfPolynomialPlot[114].X = 1.14
	pointsOfPolynomialPlot[114].Y = -162.917

	pointsOfPolynomialPlot[115].X = 1.15
	pointsOfPolynomialPlot[115].Y = -165.221

	pointsOfPolynomialPlot[116].X = 1.16
	pointsOfPolynomialPlot[116].Y = -167.53

	pointsOfPolynomialPlot[117].X = 1.17
	pointsOfPolynomialPlot[117].Y = -169.845

	pointsOfPolynomialPlot[118].X = 1.18
	pointsOfPolynomialPlot[118].Y = -172.165

	pointsOfPolynomialPlot[119].X = 1.19
	pointsOfPolynomialPlot[119].Y = -174.492

	pointsOfPolynomialPlot[120].X = 1.2
	pointsOfPolynomialPlot[120].Y = -176.823

	pointsOfPolynomialPlot[121].X = 1.21
	pointsOfPolynomialPlot[121].Y = -179.161

	pointsOfPolynomialPlot[122].X = 1.22
	pointsOfPolynomialPlot[122].Y = -181.505

	pointsOfPolynomialPlot[123].X = 1.23
	pointsOfPolynomialPlot[123].Y = -183.855

	pointsOfPolynomialPlot[124].X = 1.24
	pointsOfPolynomialPlot[124].Y = -186.211

	pointsOfPolynomialPlot[125].X = 1.25
	pointsOfPolynomialPlot[125].Y = -188.574

	pointsOfPolynomialPlot[126].X = 1.26
	pointsOfPolynomialPlot[126].Y = -190.943

	pointsOfPolynomialPlot[127].X = 1.27
	pointsOfPolynomialPlot[127].Y = -193.318

	pointsOfPolynomialPlot[128].X = 1.28
	pointsOfPolynomialPlot[128].Y = -195.7

	pointsOfPolynomialPlot[129].X = 1.29
	pointsOfPolynomialPlot[129].Y = -198.088

	pointsOfPolynomialPlot[130].X = 1.3
	pointsOfPolynomialPlot[130].Y = -200.483

	pointsOfPolynomialPlot[131].X = 1.31
	pointsOfPolynomialPlot[131].Y = -202.885

	pointsOfPolynomialPlot[132].X = 1.32
	pointsOfPolynomialPlot[132].Y = -205.294

	pointsOfPolynomialPlot[133].X = 1.33
	pointsOfPolynomialPlot[133].Y = -207.71

	pointsOfPolynomialPlot[134].X = 1.34
	pointsOfPolynomialPlot[134].Y = -210.133

	pointsOfPolynomialPlot[135].X = 1.35
	pointsOfPolynomialPlot[135].Y = -212.564

	pointsOfPolynomialPlot[136].X = 1.36
	pointsOfPolynomialPlot[136].Y = -215.001

	pointsOfPolynomialPlot[137].X = 1.37
	pointsOfPolynomialPlot[137].Y = -217.446

	pointsOfPolynomialPlot[138].X = 1.38
	pointsOfPolynomialPlot[138].Y = -219.898

	pointsOfPolynomialPlot[139].X = 1.39
	pointsOfPolynomialPlot[139].Y = -222.358

	pointsOfPolynomialPlot[140].X = 1.4
	pointsOfPolynomialPlot[140].Y = -224.825

	pointsOfPolynomialPlot[141].X = 1.41
	pointsOfPolynomialPlot[141].Y = -227.301

	pointsOfPolynomialPlot[142].X = 1.42
	pointsOfPolynomialPlot[142].Y = -229.783

	pointsOfPolynomialPlot[143].X = 1.43
	pointsOfPolynomialPlot[143].Y = -232.274

	pointsOfPolynomialPlot[144].X = 1.44
	pointsOfPolynomialPlot[144].Y = -234.773

	pointsOfPolynomialPlot[145].X = 1.45
	pointsOfPolynomialPlot[145].Y = -237.28

	pointsOfPolynomialPlot[146].X = 1.46
	pointsOfPolynomialPlot[146].Y = -239.795

	pointsOfPolynomialPlot[147].X = 1.47
	pointsOfPolynomialPlot[147].Y = -242.318

	pointsOfPolynomialPlot[148].X = 1.48
	pointsOfPolynomialPlot[148].Y = -244.849

	pointsOfPolynomialPlot[149].X = 1.49
	pointsOfPolynomialPlot[149].Y = -247.389

	pointsOfPolynomialPlot[150].X = 1.5
	pointsOfPolynomialPlot[150].Y = -249.937

	pointsOfPolynomialPlot[151].X = 1.51
	pointsOfPolynomialPlot[151].Y = -252.494

	pointsOfPolynomialPlot[152].X = 1.52
	pointsOfPolynomialPlot[152].Y = -255.059

	pointsOfPolynomialPlot[153].X = 1.53
	pointsOfPolynomialPlot[153].Y = -257.633

	pointsOfPolynomialPlot[154].X = 1.54
	pointsOfPolynomialPlot[154].Y = -260.216

	pointsOfPolynomialPlot[155].X = 1.55
	pointsOfPolynomialPlot[155].Y = -262.808

	pointsOfPolynomialPlot[156].X = 1.56
	pointsOfPolynomialPlot[156].Y = -265.409

	pointsOfPolynomialPlot[157].X = 1.57
	pointsOfPolynomialPlot[157].Y = -268.019

	pointsOfPolynomialPlot[158].X = 1.58
	pointsOfPolynomialPlot[158].Y = -270.638

	pointsOfPolynomialPlot[159].X = 1.59
	pointsOfPolynomialPlot[159].Y = -273.903

	pointsOfPolynomialPlot[160].X = 1.6
	pointsOfPolynomialPlot[160].Y = -275.903

	pointsOfPolynomialPlot[161].X = 1.61
	pointsOfPolynomialPlot[161].Y = -278.55

	pointsOfPolynomialPlot[162].X = 1.62
	pointsOfPolynomialPlot[162].Y = -281.207

	pointsOfPolynomialPlot[163].X = 1.63
	pointsOfPolynomialPlot[163].Y = -283.872

	pointsOfPolynomialPlot[164].X = 1.64
	pointsOfPolynomialPlot[164].Y = -286.548

	pointsOfPolynomialPlot[165].X = 1.65
	pointsOfPolynomialPlot[165].Y = -289.233

	pointsOfPolynomialPlot[166].X = 1.66
	pointsOfPolynomialPlot[166].Y = -291.928

	pointsOfPolynomialPlot[167].X = 1.67
	pointsOfPolynomialPlot[167].Y = -294.633

	pointsOfPolynomialPlot[168].X = 1.68
	pointsOfPolynomialPlot[168].Y = -297.348

	pointsOfPolynomialPlot[169].X = 1.69
	pointsOfPolynomialPlot[169].Y = -300.073

	pointsOfPolynomialPlot[170].X = 1.7
	pointsOfPolynomialPlot[170].Y = -302.807

	pointsOfPolynomialPlot[171].X = 1.71
	pointsOfPolynomialPlot[171].Y = -305.553

	pointsOfPolynomialPlot[172].X = 1.72
	pointsOfPolynomialPlot[172].Y = -308.308

	pointsOfPolynomialPlot[173].X = 1.73
	pointsOfPolynomialPlot[173].Y = -311.074

	pointsOfPolynomialPlot[174].X = 1.74
	pointsOfPolynomialPlot[174].Y = -313.85

	pointsOfPolynomialPlot[175].X = 1.75
	pointsOfPolynomialPlot[175].Y = -316.636

	pointsOfPolynomialPlot[176].X = 1.76
	pointsOfPolynomialPlot[176].Y = -319.433

	pointsOfPolynomialPlot[177].X = 1.77
	pointsOfPolynomialPlot[177].Y = -322.241

	pointsOfPolynomialPlot[178].X = 1.78
	pointsOfPolynomialPlot[178].Y = -325.06

	pointsOfPolynomialPlot[179].X = 1.79
	pointsOfPolynomialPlot[179].Y = -327.889

	pointsOfPolynomialPlot[180].X = 1.8
	pointsOfPolynomialPlot[180].Y = -330.729

	pointsOfPolynomialPlot[181].X = 1.81
	pointsOfPolynomialPlot[181].Y = -333.581

	pointsOfPolynomialPlot[182].X = 1.82
	pointsOfPolynomialPlot[182].Y = -336.443

	pointsOfPolynomialPlot[183].X = 1.83
	pointsOfPolynomialPlot[183].Y = -339.316

	pointsOfPolynomialPlot[184].X = 1.84
	pointsOfPolynomialPlot[184].Y = -342.201

	pointsOfPolynomialPlot[185].X = 1.85
	pointsOfPolynomialPlot[185].Y = -345.097

	pointsOfPolynomialPlot[186].X = 1.86
	pointsOfPolynomialPlot[186].Y = -348.004

	pointsOfPolynomialPlot[187].X = 1.87
	pointsOfPolynomialPlot[187].Y = -350.922

	pointsOfPolynomialPlot[188].X = 1.88
	pointsOfPolynomialPlot[188].Y = -353.853

	pointsOfPolynomialPlot[189].X = 1.89
	pointsOfPolynomialPlot[189].Y = -356.794

	pointsOfPolynomialPlot[190].X = 1.9
	pointsOfPolynomialPlot[190].Y = -359.747

	pointsOfPolynomialPlot[191].X = 1.91
	pointsOfPolynomialPlot[191].Y = -362.712

	pointsOfPolynomialPlot[192].X = 1.92
	pointsOfPolynomialPlot[192].Y = -365.689

	pointsOfPolynomialPlot[193].X = 1.93
	pointsOfPolynomialPlot[193].Y = -368.678

	pointsOfPolynomialPlot[194].X = 1.94
	pointsOfPolynomialPlot[194].Y = -371.679

	pointsOfPolynomialPlot[195].X = 1.95
	pointsOfPolynomialPlot[195].Y = -374.691

	pointsOfPolynomialPlot[196].X = 1.96
	pointsOfPolynomialPlot[196].Y = -377.716

	pointsOfPolynomialPlot[197].X = 1.97
	pointsOfPolynomialPlot[197].Y = -380.753

	pointsOfPolynomialPlot[198].X = 1.98
	pointsOfPolynomialPlot[198].Y = -383.802

	pointsOfPolynomialPlot[199].X = 1.99
	pointsOfPolynomialPlot[199].Y = -386.863

	pointsOfPolynomialPlot[200].X = 2.0
	pointsOfPolynomialPlot[200].Y = -389.937

	pointsOfPolynomialPlot[201].X = 2.01
	pointsOfPolynomialPlot[201].Y = -393.023

	pointsOfPolynomialPlot[202].X = 2.02
	pointsOfPolynomialPlot[202].Y = -396.122

	pointsOfPolynomialPlot[203].X = 2.03
	pointsOfPolynomialPlot[203].Y = -399.233

	pointsOfPolynomialPlot[204].X = 2.04
	pointsOfPolynomialPlot[204].Y = -402.357

	pointsOfPolynomialPlot[205].X = 2.05
	pointsOfPolynomialPlot[205].Y = -405.494

	pointsOfPolynomialPlot[206].X = 2.06
	pointsOfPolynomialPlot[206].Y = -408.644

	pointsOfPolynomialPlot[207].X = 2.07
	pointsOfPolynomialPlot[207].Y = -411.806

	pointsOfPolynomialPlot[208].X = 2.08
	pointsOfPolynomialPlot[208].Y = -414.982

	pointsOfPolynomialPlot[209].X = 2.09
	pointsOfPolynomialPlot[209].Y = -418.17

	pointsOfPolynomialPlot[210].X = 2.10
	pointsOfPolynomialPlot[210].Y = -421.371

	pointsOfPolynomialPlot[211].X = 2.11
	pointsOfPolynomialPlot[211].Y = -424.586

	pointsOfPolynomialPlot[212].X = 2.12
	pointsOfPolynomialPlot[212].Y = -427.814

	pointsOfPolynomialPlot[213].X = 2.13
	pointsOfPolynomialPlot[213].Y = -431.055

	pointsOfPolynomialPlot[214].X = 2.14
	pointsOfPolynomialPlot[214].Y = -434.31

	pointsOfPolynomialPlot[215].X = 2.15
	pointsOfPolynomialPlot[215].Y = -437.578

	pointsOfPolynomialPlot[216].X = 2.16
	pointsOfPolynomialPlot[216].Y = -440.859

	pointsOfPolynomialPlot[217].X = 2.17
	pointsOfPolynomialPlot[217].Y = -444.154

	pointsOfPolynomialPlot[218].X = 2.18
	pointsOfPolynomialPlot[218].Y = -447.463

	pointsOfPolynomialPlot[219].X = 2.19
	pointsOfPolynomialPlot[219].Y = -450.785

	pointsOfPolynomialPlot[220].X = 2.2
	pointsOfPolynomialPlot[220].Y = -454.121

	pointsOfPolynomialPlot[221].X = 2.21
	pointsOfPolynomialPlot[221].Y = -457.471

	pointsOfPolynomialPlot[222].X = 2.22
	pointsOfPolynomialPlot[222].Y = -460.835

	pointsOfPolynomialPlot[223].X = 2.23
	pointsOfPolynomialPlot[223].Y = -464.213

	pointsOfPolynomialPlot[224].X = 2.24
	pointsOfPolynomialPlot[224].Y = -467.605

	pointsOfPolynomialPlot[225].X = 2.25
	pointsOfPolynomialPlot[225].Y = -471.011

	pointsOfPolynomialPlot[226].X = 2.26
	pointsOfPolynomialPlot[226].Y = -474.431

	pointsOfPolynomialPlot[227].X = 2.27
	pointsOfPolynomialPlot[227].Y = -477.866

	pointsOfPolynomialPlot[228].X = 2.28
	pointsOfPolynomialPlot[228].Y = -481.315

	pointsOfPolynomialPlot[229].X = 2.29
	pointsOfPolynomialPlot[229].Y = -484.778

	pointsOfPolynomialPlot[230].X = 2.3
	pointsOfPolynomialPlot[230].Y = -488.255

	pointsOfPolynomialPlot[231].X = 2.31
	pointsOfPolynomialPlot[231].Y = -491.748

	pointsOfPolynomialPlot[232].X = 2.32
	pointsOfPolynomialPlot[232].Y = -495.254

	pointsOfPolynomialPlot[233].X = 2.33
	pointsOfPolynomialPlot[233].Y = -498.775

	pointsOfPolynomialPlot[234].X = 2.34
	pointsOfPolynomialPlot[234].Y = -502.311

	pointsOfPolynomialPlot[235].X = 2.35
	pointsOfPolynomialPlot[235].Y = -505.862

	pointsOfPolynomialPlot[236].X = 2.36
	pointsOfPolynomialPlot[236].Y = -509.428

	pointsOfPolynomialPlot[237].X = 2.37
	pointsOfPolynomialPlot[237].Y = -513.008

	pointsOfPolynomialPlot[238].X = 2.38
	pointsOfPolynomialPlot[238].Y = -516.604

	pointsOfPolynomialPlot[239].X = 2.39
	pointsOfPolynomialPlot[239].Y = -520.214

	pointsOfPolynomialPlot[240].X = 2.4
	pointsOfPolynomialPlot[240].Y = -523.839

	pointsOfPolynomialPlot[241].X = 2.41
	pointsOfPolynomialPlot[241].Y = -527.48

	pointsOfPolynomialPlot[242].X = 2.42
	pointsOfPolynomialPlot[242].Y = -531.136

	pointsOfPolynomialPlot[243].X = 2.43
	pointsOfPolynomialPlot[243].Y = -534.807

	pointsOfPolynomialPlot[244].X = 2.44
	pointsOfPolynomialPlot[244].Y = -538.493

	pointsOfPolynomialPlot[245].X = 2.45
	pointsOfPolynomialPlot[245].Y = -542.195

	pointsOfPolynomialPlot[246].X = 2.46
	pointsOfPolynomialPlot[246].Y = -545.912

	pointsOfPolynomialPlot[247].X = 2.47
	pointsOfPolynomialPlot[247].Y = -549.645

	pointsOfPolynomialPlot[248].X = 2.48
	pointsOfPolynomialPlot[248].Y = -553.393

	pointsOfPolynomialPlot[249].X = 2.49
	pointsOfPolynomialPlot[249].Y = -557.157

	pointsOfPolynomialPlot[250].X = 2.5
	pointsOfPolynomialPlot[250].Y = -560.937

	pointsOfPolynomialPlot[251].X = 2.51
	pointsOfPolynomialPlot[251].Y = -564.732

	pointsOfPolynomialPlot[252].X = 2.52
	pointsOfPolynomialPlot[252].Y = -568.544

	pointsOfPolynomialPlot[253].X = 2.53
	pointsOfPolynomialPlot[253].Y = -572.371

	pointsOfPolynomialPlot[254].X = 2.54
	pointsOfPolynomialPlot[254].Y = -576.214

	pointsOfPolynomialPlot[255].X = 2.55
	pointsOfPolynomialPlot[255].Y = -580.073

	pointsOfPolynomialPlot[256].X = 2.56
	pointsOfPolynomialPlot[256].Y = -583.948

	pointsOfPolynomialPlot[257].X = 2.57
	pointsOfPolynomialPlot[257].Y = -587.839

	pointsOfPolynomialPlot[258].X = 2.58
	pointsOfPolynomialPlot[258].Y = -591.746

	pointsOfPolynomialPlot[259].X = 2.59
	pointsOfPolynomialPlot[259].Y = -595.67

	pointsOfPolynomialPlot[260].X = 2.6
	pointsOfPolynomialPlot[260].Y = -599.609

	pointsOfPolynomialPlot[261].X = 2.61
	pointsOfPolynomialPlot[261].Y = -603.566

	pointsOfPolynomialPlot[262].X = 2.62
	pointsOfPolynomialPlot[262].Y = -607.538

	pointsOfPolynomialPlot[263].X = 2.63
	pointsOfPolynomialPlot[263].Y = -611.527

	pointsOfPolynomialPlot[264].X = 2.64
	pointsOfPolynomialPlot[264].Y = -615.533

	pointsOfPolynomialPlot[265].X = 2.65
	pointsOfPolynomialPlot[265].Y = -619.555

	pointsOfPolynomialPlot[266].X = 2.66
	pointsOfPolynomialPlot[266].Y = -623.593

	pointsOfPolynomialPlot[267].X = 2.67
	pointsOfPolynomialPlot[267].Y = -627.649

	pointsOfPolynomialPlot[268].X = 2.68
	pointsOfPolynomialPlot[268].Y = -631.721

	pointsOfPolynomialPlot[269].X = 2.69
	pointsOfPolynomialPlot[269].Y = -635.81

	pointsOfPolynomialPlot[270].X = 2.7
	pointsOfPolynomialPlot[270].Y = -639.915

	pointsOfPolynomialPlot[271].X = 2.71
	pointsOfPolynomialPlot[271].Y = -644.038

	pointsOfPolynomialPlot[272].X = 2.72
	pointsOfPolynomialPlot[272].Y = -648.178

	pointsOfPolynomialPlot[273].X = 2.73
	pointsOfPolynomialPlot[273].Y = -652.334

	pointsOfPolynomialPlot[274].X = 2.74
	pointsOfPolynomialPlot[274].Y = -656.508

	pointsOfPolynomialPlot[275].X = 2.75
	pointsOfPolynomialPlot[275].Y = -660.699

	pointsOfPolynomialPlot[276].X = 2.76
	pointsOfPolynomialPlot[276].Y = -664.907

	pointsOfPolynomialPlot[277].X = 2.77
	pointsOfPolynomialPlot[277].Y = -669.132

	pointsOfPolynomialPlot[278].X = 2.78
	pointsOfPolynomialPlot[278].Y = -673.374

	pointsOfPolynomialPlot[279].X = 2.79
	pointsOfPolynomialPlot[279].Y = -677.634

	pointsOfPolynomialPlot[280].X = 2.8
	pointsOfPolynomialPlot[280].Y = -681.911

	pointsOfPolynomialPlot[281].X = 2.81
	pointsOfPolynomialPlot[281].Y = -686.206

	pointsOfPolynomialPlot[282].X = 2.82
	pointsOfPolynomialPlot[282].Y = -690.518

	pointsOfPolynomialPlot[283].X = 2.83
	pointsOfPolynomialPlot[283].Y = -694.848

	pointsOfPolynomialPlot[284].X = 2.84
	pointsOfPolynomialPlot[284].Y = -699.195

	pointsOfPolynomialPlot[285].X = 2.85
	pointsOfPolynomialPlot[285].Y = -703.56

	pointsOfPolynomialPlot[286].X = 2.86
	pointsOfPolynomialPlot[286].Y = -707.943

	pointsOfPolynomialPlot[287].X = 2.87
	pointsOfPolynomialPlot[287].Y = -712.343

	pointsOfPolynomialPlot[288].X = 2.88
	pointsOfPolynomialPlot[288].Y = -716.761

	pointsOfPolynomialPlot[289].X = 2.89
	pointsOfPolynomialPlot[289].Y = -721.197

	pointsOfPolynomialPlot[290].X = 2.9
	pointsOfPolynomialPlot[290].Y = -725.651

	pointsOfPolynomialPlot[291].X = 2.91
	pointsOfPolynomialPlot[291].Y = -730.123

	pointsOfPolynomialPlot[292].X = 2.92
	pointsOfPolynomialPlot[292].Y = -734.613

	pointsOfPolynomialPlot[293].X = 2.93
	pointsOfPolynomialPlot[293].Y = -739.121

	pointsOfPolynomialPlot[294].X = 2.94
	pointsOfPolynomialPlot[294].Y = -743.647

	pointsOfPolynomialPlot[295].X = 2.95
	pointsOfPolynomialPlot[295].Y = -748.192

	pointsOfPolynomialPlot[296].X = 2.96
	pointsOfPolynomialPlot[296].Y = -752.754

	pointsOfPolynomialPlot[297].X = 2.97
	pointsOfPolynomialPlot[297].Y = -757.335

	pointsOfPolynomialPlot[298].X = 2.98
	pointsOfPolynomialPlot[298].Y = -761.934

	pointsOfPolynomialPlot[299].X = 2.99
	pointsOfPolynomialPlot[299].Y = -766.551

	pointsOfPolynomialPlot[300].X = 3.0
	pointsOfPolynomialPlot[300].Y = -771.187

	pointsOfPolynomialPlot[301].X = 3.01
	pointsOfPolynomialPlot[301].Y = -775.841

	pointsOfPolynomialPlot[302].X = 3.02
	pointsOfPolynomialPlot[302].Y = -780.514

	pointsOfPolynomialPlot[303].X = 3.03
	pointsOfPolynomialPlot[303].Y = -785.205

	pointsOfPolynomialPlot[304].X = 3.04
	pointsOfPolynomialPlot[304].Y = -789.915

	pointsOfPolynomialPlot[305].X = 3.05
	pointsOfPolynomialPlot[305].Y = -794.644

	pointsOfPolynomialPlot[306].X = 3.06
	pointsOfPolynomialPlot[306].Y = -799.391

	pointsOfPolynomialPlot[307].X = 3.07
	pointsOfPolynomialPlot[307].Y = -804.157

	pointsOfPolynomialPlot[308].X = 3.08
	pointsOfPolynomialPlot[308].Y = -808.942

	pointsOfPolynomialPlot[309].X = 3.09
	pointsOfPolynomialPlot[309].Y = -813.745

	pointsOfPolynomialPlot[310].X = 3.1
	pointsOfPolynomialPlot[310].Y = -818.745

	pointsOfPolynomialPlot[311].X = 3.11
	pointsOfPolynomialPlot[311].Y = -823.409

	pointsOfPolynomialPlot[312].X = 3.12
	pointsOfPolynomialPlot[312].Y = -828.269

	pointsOfPolynomialPlot[313].X = 3.13
	pointsOfPolynomialPlot[313].Y = -833.148

	pointsOfPolynomialPlot[314].X = 3.14
	pointsOfPolynomialPlot[314].Y = -838.047

	pointsOfPolynomialPlot[315].X = 3.15
	pointsOfPolynomialPlot[315].Y = -842.964

	pointsOfPolynomialPlot[316].X = 3.16
	pointsOfPolynomialPlot[316].Y = -847.901

	pointsOfPolynomialPlot[317].X = 3.17
	pointsOfPolynomialPlot[317].Y = -852.857

	pointsOfPolynomialPlot[318].X = 3.18
	pointsOfPolynomialPlot[318].Y = -857.832

	pointsOfPolynomialPlot[319].X = 3.19
	pointsOfPolynomialPlot[319].Y = -862.826

	pointsOfPolynomialPlot[320].X = 3.2
	pointsOfPolynomialPlot[320].Y = -867.839

	pointsOfPolynomialPlot[321].X = 3.21
	pointsOfPolynomialPlot[321].Y = -872.872

	pointsOfPolynomialPlot[322].X = 3.22
	pointsOfPolynomialPlot[322].Y = -877.925

	pointsOfPolynomialPlot[323].X = 3.23
	pointsOfPolynomialPlot[323].Y = -882.997

	pointsOfPolynomialPlot[324].X = 3.24
	pointsOfPolynomialPlot[324].Y = -888.088

	pointsOfPolynomialPlot[325].X = 3.25
	pointsOfPolynomialPlot[325].Y = -893.329

	pointsOfPolynomialPlot[326].X = 3.26
	pointsOfPolynomialPlot[326].Y = -898.329

	pointsOfPolynomialPlot[327].X = 3.27
	pointsOfPolynomialPlot[327].Y = -903.479

	pointsOfPolynomialPlot[328].X = 3.28
	pointsOfPolynomialPlot[328].Y = -908.649

	pointsOfPolynomialPlot[329].X = 3.29
	pointsOfPolynomialPlot[329].Y = -913.838

	pointsOfPolynomialPlot[330].X = 3.3
	pointsOfPolynomialPlot[330].Y = -919.047

	pointsOfPolynomialPlot[331].X = 3.31
	pointsOfPolynomialPlot[331].Y = -924.276

	pointsOfPolynomialPlot[332].X = 3.32
	pointsOfPolynomialPlot[332].Y = -929.525

	pointsOfPolynomialPlot[333].X = 3.33
	pointsOfPolynomialPlot[333].Y = -934.794

	pointsOfPolynomialPlot[334].X = 3.34
	pointsOfPolynomialPlot[334].Y = -940.082

	pointsOfPolynomialPlot[335].X = 3.35
	pointsOfPolynomialPlot[335].Y = -945.391

	pointsOfPolynomialPlot[336].X = 3.36
	pointsOfPolynomialPlot[336].Y = -950.719

	pointsOfPolynomialPlot[337].X = 3.37
	pointsOfPolynomialPlot[337].Y = -956.068

	pointsOfPolynomialPlot[338].X = 3.38
	pointsOfPolynomialPlot[338].Y = -961.436

	pointsOfPolynomialPlot[339].X = 3.39
	pointsOfPolynomialPlot[339].Y = -966.825

	pointsOfPolynomialPlot[340].X = 3.4
	pointsOfPolynomialPlot[340].Y = -972.233

	pointsOfPolynomialPlot[341].X = 3.41
	pointsOfPolynomialPlot[341].Y = -977.662

	pointsOfPolynomialPlot[342].X = 3.42
	pointsOfPolynomialPlot[342].Y = -983.112

	pointsOfPolynomialPlot[343].X = 3.43
	pointsOfPolynomialPlot[343].Y = -988.581

	pointsOfPolynomialPlot[344].X = 3.44
	pointsOfPolynomialPlot[344].Y = -994.071

	pointsOfPolynomialPlot[345].X = 3.45
	pointsOfPolynomialPlot[345].Y = -999.581

	pointsOfPolynomialPlot[346].X = 3.46
	pointsOfPolynomialPlot[346].Y = -1_005.111

	pointsOfPolynomialPlot[347].X = 3.47
	pointsOfPolynomialPlot[347].Y = -1_010.662

	pointsOfPolynomialPlot[348].X = 3.48
	pointsOfPolynomialPlot[348].Y = -1_016.233

	pointsOfPolynomialPlot[349].X = 3.49
	pointsOfPolynomialPlot[349].Y = -1_021.825

	pointsOfPolynomialPlot[350].X = 3.5
	pointsOfPolynomialPlot[350].Y = -1_027.437

	pointsOfPolynomialPlot[351].X = 3.51
	pointsOfPolynomialPlot[351].Y = -1_033.07

	pointsOfPolynomialPlot[352].X = 3.52
	pointsOfPolynomialPlot[352].Y = -1_038.723

	pointsOfPolynomialPlot[353].X = 3.53
	pointsOfPolynomialPlot[353].Y = -1_044.397

	pointsOfPolynomialPlot[354].X = 3.54
	pointsOfPolynomialPlot[354].Y = -1_050.092

	pointsOfPolynomialPlot[355].X = 3.55
	pointsOfPolynomialPlot[355].Y = -1_055.807

	pointsOfPolynomialPlot[356].X = 3.56
	pointsOfPolynomialPlot[356].Y = -1_061.543

	pointsOfPolynomialPlot[357].X = 3.57
	pointsOfPolynomialPlot[357].Y = -1_067.3

	pointsOfPolynomialPlot[358].X = 3.58
	pointsOfPolynomialPlot[358].Y = -1_073.078

	pointsOfPolynomialPlot[359].X = 3.59
	pointsOfPolynomialPlot[359].Y = -1_078.876

	pointsOfPolynomialPlot[360].X = 3.6
	pointsOfPolynomialPlot[360].Y = -1_084.695

	pointsOfPolynomialPlot[361].X = 3.61
	pointsOfPolynomialPlot[361].Y = -1_090.536

	pointsOfPolynomialPlot[362].X = 3.62
	pointsOfPolynomialPlot[362].Y = -1_096.397

	pointsOfPolynomialPlot[363].X = 3.63
	pointsOfPolynomialPlot[363].Y = -1_102.279

	pointsOfPolynomialPlot[364].X = 3.64
	pointsOfPolynomialPlot[364].Y = -1_108.182

	pointsOfPolynomialPlot[365].X = 3.65
	pointsOfPolynomialPlot[365].Y = -1_114.106

	pointsOfPolynomialPlot[366].X = 3.66
	pointsOfPolynomialPlot[366].Y = -1_120.051

	pointsOfPolynomialPlot[367].X = 3.67
	pointsOfPolynomialPlot[367].Y = -1_126.018

	pointsOfPolynomialPlot[368].X = 3.68
	pointsOfPolynomialPlot[368].Y = -1_132.005

	pointsOfPolynomialPlot[369].X = 3.69
	pointsOfPolynomialPlot[369].Y = -1_138.014

	pointsOfPolynomialPlot[370].X = 3.7
	pointsOfPolynomialPlot[370].Y = -1_144.043

	pointsOfPolynomialPlot[371].X = 3.71
	pointsOfPolynomialPlot[371].Y = -1_150.094

	pointsOfPolynomialPlot[372].X = 3.72
	pointsOfPolynomialPlot[372].Y = -1_156.167

	pointsOfPolynomialPlot[373].X = 3.73
	pointsOfPolynomialPlot[373].Y = -1_162.26

	pointsOfPolynomialPlot[374].X = 3.74
	pointsOfPolynomialPlot[374].Y = -1_168.375

	pointsOfPolynomialPlot[375].X = 3.75
	pointsOfPolynomialPlot[375].Y = -1_174.511

	pointsOfPolynomialPlot[376].X = 3.76
	pointsOfPolynomialPlot[376].Y = -1_180.669

	pointsOfPolynomialPlot[377].X = 3.77
	pointsOfPolynomialPlot[377].Y = -1_186.848

	pointsOfPolynomialPlot[378].X = 3.78
	pointsOfPolynomialPlot[378].Y = -1_193.048

	pointsOfPolynomialPlot[379].X = 3.79
	pointsOfPolynomialPlot[379].Y = -1_199.27

	pointsOfPolynomialPlot[380].X = 3.8
	pointsOfPolynomialPlot[380].Y = -1_205.513

	pointsOfPolynomialPlot[381].X = 3.81
	pointsOfPolynomialPlot[381].Y = -1_211.778

	pointsOfPolynomialPlot[382].X = 3.82
	pointsOfPolynomialPlot[382].Y = -1_218.065

	pointsOfPolynomialPlot[383].X = 3.83
	pointsOfPolynomialPlot[383].Y = -1_224.373

	pointsOfPolynomialPlot[384].X = 3.84
	pointsOfPolynomialPlot[384].Y = -1_230.702

	pointsOfPolynomialPlot[385].X = 3.85
	pointsOfPolynomialPlot[385].Y = -1_237.054

	pointsOfPolynomialPlot[386].X = 3.86
	pointsOfPolynomialPlot[386].Y = -1_243.427

	pointsOfPolynomialPlot[387].X = 3.87
	pointsOfPolynomialPlot[387].Y = -1_249.821

	pointsOfPolynomialPlot[388].X = 3.88
	pointsOfPolynomialPlot[388].Y = -1_256.238

	pointsOfPolynomialPlot[389].X = 3.89
	pointsOfPolynomialPlot[389].Y = -1_262.676

	pointsOfPolynomialPlot[390].X = 3.9
	pointsOfPolynomialPlot[390].Y = -1_269.135

	pointsOfPolynomialPlot[391].X = 3.91
	pointsOfPolynomialPlot[391].Y = -1_275.617

	pointsOfPolynomialPlot[392].X = 3.92
	pointsOfPolynomialPlot[392].Y = -1_282.121

	pointsOfPolynomialPlot[393].X = 3.93
	pointsOfPolynomialPlot[393].Y = -1_288.646

	pointsOfPolynomialPlot[394].X = 3.94
	pointsOfPolynomialPlot[394].Y = -1_295.193

	pointsOfPolynomialPlot[395].X = 3.95
	pointsOfPolynomialPlot[395].Y = -1_301.762

	pointsOfPolynomialPlot[396].X = 3.96
	pointsOfPolynomialPlot[396].Y = -1_308.353

	pointsOfPolynomialPlot[397].X = 3.97
	pointsOfPolynomialPlot[397].Y = -1_314.966

	pointsOfPolynomialPlot[398].X = 3.98
	pointsOfPolynomialPlot[398].Y = -1_321.601

	pointsOfPolynomialPlot[399].X = 3.99
	pointsOfPolynomialPlot[399].Y = -1_328.258

	pointsOfPolynomialPlot[400].X = 4.0
	pointsOfPolynomialPlot[400].Y = -1_334.937

	pointsOfPolynomialPlot[401].X = 4.01
	pointsOfPolynomialPlot[401].Y = -1_341.638

	pointsOfPolynomialPlot[402].X = 4.02
	pointsOfPolynomialPlot[402].Y = -1_348.361

	pointsOfPolynomialPlot[403].X = 4.03
	pointsOfPolynomialPlot[403].Y = -1_355.106

	pointsOfPolynomialPlot[404].X = 4.04
	pointsOfPolynomialPlot[404].Y = -1_361.874

	pointsOfPolynomialPlot[405].X = 4.05
	pointsOfPolynomialPlot[405].Y = -1_368.663

	pointsOfPolynomialPlot[406].X = 4.06
	pointsOfPolynomialPlot[406].Y = -1_375.475

	pointsOfPolynomialPlot[407].X = 4.07
	pointsOfPolynomialPlot[407].Y = -1_382.309

	pointsOfPolynomialPlot[408].X = 4.08
	pointsOfPolynomialPlot[408].Y = -1_389.165

	pointsOfPolynomialPlot[409].X = 4.09
	pointsOfPolynomialPlot[409].Y = -1_396.043

	pointsOfPolynomialPlot[410].X = 4.1
	pointsOfPolynomialPlot[410].Y = -1_402.943

	pointsOfPolynomialPlot[411].X = 4.11
	pointsOfPolynomialPlot[411].Y = -1_409.866

	pointsOfPolynomialPlot[412].X = 4.12
	pointsOfPolynomialPlot[412].Y = -1_416.811

	pointsOfPolynomialPlot[413].X = 4.13
	pointsOfPolynomialPlot[413].Y = -1_423.779

	pointsOfPolynomialPlot[414].X = 4.14
	pointsOfPolynomialPlot[414].Y = -1_430.769

	pointsOfPolynomialPlot[415].X = 4.15
	pointsOfPolynomialPlot[415].Y = -1_437.781

	pointsOfPolynomialPlot[416].X = 4.16
	pointsOfPolynomialPlot[416].Y = -1_444.815

	pointsOfPolynomialPlot[417].X = 4.17
	pointsOfPolynomialPlot[417].Y = -1_451.872

	pointsOfPolynomialPlot[418].X = 4.18
	pointsOfPolynomialPlot[418].Y = -1_458.951

	pointsOfPolynomialPlot[419].X = 4.19
	pointsOfPolynomialPlot[419].Y = -1_466.053

	pointsOfPolynomialPlot[420].X = 4.20
	pointsOfPolynomialPlot[420].Y = -1_473.177

	pointsOfPolynomialPlot[421].X = 4.21
	pointsOfPolynomialPlot[421].Y = -1_480.324

	pointsOfPolynomialPlot[422].X = 4.22
	pointsOfPolynomialPlot[422].Y = -1_487.493

	pointsOfPolynomialPlot[423].X = 4.23
	pointsOfPolynomialPlot[423].Y = -1_494.685

	pointsOfPolynomialPlot[424].X = 4.24
	pointsOfPolynomialPlot[424].Y = -1_501.899

	pointsOfPolynomialPlot[425].X = 4.25
	pointsOfPolynomialPlot[425].Y = -1_509.136

	pointsOfPolynomialPlot[426].X = 4.26
	pointsOfPolynomialPlot[426].Y = -1_516.396

	pointsOfPolynomialPlot[427].X = 4.27
	pointsOfPolynomialPlot[427].Y = -1_523.678

	pointsOfPolynomialPlot[428].X = 4.28
	pointsOfPolynomialPlot[428].Y = -1_530.982

	pointsOfPolynomialPlot[429].X = 4.29
	pointsOfPolynomialPlot[429].Y = -1_538.31

	pointsOfPolynomialPlot[430].X = 4.30
	pointsOfPolynomialPlot[430].Y = -1_545.659

	pointsOfPolynomialPlot[431].X = 4.31
	pointsOfPolynomialPlot[431].Y = -1_553.032

	pointsOfPolynomialPlot[432].X = 4.32
	pointsOfPolynomialPlot[432].Y = -1_560.427

	pointsOfPolynomialPlot[433].X = 4.33
	pointsOfPolynomialPlot[433].Y = -1_567.845

	pointsOfPolynomialPlot[434].X = 4.34
	pointsOfPolynomialPlot[434].Y = -1_575.286

	pointsOfPolynomialPlot[435].X = 4.35
	pointsOfPolynomialPlot[435].Y = -1_582.749

	pointsOfPolynomialPlot[436].X = 4.36
	pointsOfPolynomialPlot[436].Y = -1_590.235

	pointsOfPolynomialPlot[437].X = 4.37
	pointsOfPolynomialPlot[437].Y = -1_597.744

	pointsOfPolynomialPlot[438].X = 4.38
	pointsOfPolynomialPlot[438].Y = -1_605.276

	pointsOfPolynomialPlot[439].X = 4.39
	pointsOfPolynomialPlot[439].Y = -1_612.83

	pointsOfPolynomialPlot[440].X = 4.40
	pointsOfPolynomialPlot[440].Y = -1_620.407

	pointsOfPolynomialPlot[441].X = 4.41
	pointsOfPolynomialPlot[441].Y = -1_628.008

	pointsOfPolynomialPlot[442].X = 4.42
	pointsOfPolynomialPlot[442].Y = -1_635.63

	pointsOfPolynomialPlot[443].X = 4.43
	pointsOfPolynomialPlot[443].Y = -1_643.276

	pointsOfPolynomialPlot[444].X = 4.44
	pointsOfPolynomialPlot[444].Y = -1_650.945

	pointsOfPolynomialPlot[445].X = 4.45
	pointsOfPolynomialPlot[445].Y = -1_658.636

	pointsOfPolynomialPlot[446].X = 4.46
	pointsOfPolynomialPlot[446].Y = -1_666.351

	pointsOfPolynomialPlot[447].X = 4.47
	pointsOfPolynomialPlot[447].Y = -1_674.088

	pointsOfPolynomialPlot[448].X = 4.48
	pointsOfPolynomialPlot[448].Y = -1_681.848

	pointsOfPolynomialPlot[449].X = 4.49
	pointsOfPolynomialPlot[449].Y = -1_689.631

	pointsOfPolynomialPlot[450].X = 4.5
	pointsOfPolynomialPlot[450].Y = -1_697.437

	pointsOfPolynomialPlot[451].X = 4.51
	pointsOfPolynomialPlot[451].Y = -1_705.266

	pointsOfPolynomialPlot[452].X = 4.52
	pointsOfPolynomialPlot[452].Y = -1_713.118

	pointsOfPolynomialPlot[453].X = 4.53
	pointsOfPolynomialPlot[453].Y = -1_720.993

	pointsOfPolynomialPlot[454].X = 4.54
	pointsOfPolynomialPlot[454].Y = -1_728.891

	pointsOfPolynomialPlot[455].X = 4.55
	pointsOfPolynomialPlot[455].Y = -1_736.812

	pointsOfPolynomialPlot[456].X = 4.56
	pointsOfPolynomialPlot[456].Y = -1_744.756

	pointsOfPolynomialPlot[457].X = 4.57
	pointsOfPolynomialPlot[457].Y = -1_752.722

	pointsOfPolynomialPlot[458].X = 4.58
	pointsOfPolynomialPlot[458].Y = -1_760.712

	pointsOfPolynomialPlot[459].X = 4.59
	pointsOfPolynomialPlot[459].Y = -1_768.725

	pointsOfPolynomialPlot[460].X = 4.6
	pointsOfPolynomialPlot[460].Y = -1_776.761

	pointsOfPolynomialPlot[461].X = 4.61
	pointsOfPolynomialPlot[461].Y = -1_784.821

	pointsOfPolynomialPlot[462].X = 4.62
	pointsOfPolynomialPlot[462].Y = -1_792.309

	pointsOfPolynomialPlot[463].X = 4.63
	pointsOfPolynomialPlot[463].Y = -1_801.008

	pointsOfPolynomialPlot[464].X = 4.64
	pointsOfPolynomialPlot[464].Y = -1_809.136

	pointsOfPolynomialPlot[465].X = 4.65
	pointsOfPolynomialPlot[465].Y = -1_817.288

	pointsOfPolynomialPlot[466].X = 4.66
	pointsOfPolynomialPlot[466].Y = -1_825.462

	pointsOfPolynomialPlot[467].X = 4.67
	pointsOfPolynomialPlot[467].Y = -1_833.66

	pointsOfPolynomialPlot[468].X = 4.68
	pointsOfPolynomialPlot[468].Y = -1_841.881

	pointsOfPolynomialPlot[469].X = 4.69
	pointsOfPolynomialPlot[469].Y = -1_850.124

	pointsOfPolynomialPlot[470].X = 4.7
	pointsOfPolynomialPlot[470].Y = -1_858.391

	pointsOfPolynomialPlot[471].X = 4.71
	pointsOfPolynomialPlot[471].Y = -1_866.682

	pointsOfPolynomialPlot[472].X = 4.72
	pointsOfPolynomialPlot[472].Y = -1_874.995

	pointsOfPolynomialPlot[473].X = 4.73
	pointsOfPolynomialPlot[473].Y = -1_883.331

	pointsOfPolynomialPlot[474].X = 4.74
	pointsOfPolynomialPlot[474].Y = -1_891.691

	pointsOfPolynomialPlot[475].X = 4.75
	pointsOfPolynomialPlot[475].Y = -1_900.074

	pointsOfPolynomialPlot[476].X = 4.76
	pointsOfPolynomialPlot[476].Y = -1_908.48

	pointsOfPolynomialPlot[477].X = 4.77
	pointsOfPolynomialPlot[477].Y = -1_916.909

	pointsOfPolynomialPlot[478].X = 4.78
	pointsOfPolynomialPlot[478].Y = -1_925.361

	pointsOfPolynomialPlot[479].X = 4.79
	pointsOfPolynomialPlot[479].Y = -1_933.837

	pointsOfPolynomialPlot[480].X = 4.8
	pointsOfPolynomialPlot[480].Y = -1_942.335

	pointsOfPolynomialPlot[481].X = 4.81
	pointsOfPolynomialPlot[481].Y = -1_950.857

	pointsOfPolynomialPlot[482].X = 4.82
	pointsOfPolynomialPlot[482].Y = -1_959.403

	pointsOfPolynomialPlot[483].X = 4.83
	pointsOfPolynomialPlot[483].Y = -1_967.971

	pointsOfPolynomialPlot[484].X = 4.84
	pointsOfPolynomialPlot[484].Y = -1_976.562

	pointsOfPolynomialPlot[485].X = 4.85
	pointsOfPolynomialPlot[485].Y = -1_985.177

	pointsOfPolynomialPlot[486].X = 4.86
	pointsOfPolynomialPlot[486].Y = -1_993.815

	pointsOfPolynomialPlot[487].X = 4.87
	pointsOfPolynomialPlot[487].Y = -2_002.476

	pointsOfPolynomialPlot[488].X = 4.88
	pointsOfPolynomialPlot[488].Y = -2_011.161

	pointsOfPolynomialPlot[489].X = 4.89
	pointsOfPolynomialPlot[489].Y = -2_019.869

	pointsOfPolynomialPlot[490].X = 4.9
	pointsOfPolynomialPlot[490].Y = -2_028.599

	pointsOfPolynomialPlot[491].X = 4.91
	pointsOfPolynomialPlot[491].Y = -2_037.354

	pointsOfPolynomialPlot[492].X = 4.92
	pointsOfPolynomialPlot[492].Y = -2_046.131

	pointsOfPolynomialPlot[493].X = 4.93
	pointsOfPolynomialPlot[493].Y = -2_054.932

	pointsOfPolynomialPlot[494].X = 4.94
	pointsOfPolynomialPlot[494].Y = -2_063.756

	pointsOfPolynomialPlot[495].X = 4.95
	pointsOfPolynomialPlot[495].Y = -2_072.603

	pointsOfPolynomialPlot[496].X = 4.96
	pointsOfPolynomialPlot[496].Y = -2_081.473

	pointsOfPolynomialPlot[497].X = 4.97
	pointsOfPolynomialPlot[497].Y = -2_090.367

	pointsOfPolynomialPlot[498].X = 4.98
	pointsOfPolynomialPlot[498].Y = -2_099.284

	pointsOfPolynomialPlot[499].X = 4.99
	pointsOfPolynomialPlot[499].Y = -2_108.224

	pointsOfPolynomialPlot[500].X = 5.0
	pointsOfPolynomialPlot[500].Y = -2_117.187

	pointsOfPolynomialPlot[501].X = 5.01
	pointsOfPolynomialPlot[501].Y = -2_122.174

	pointsOfPolynomialPlot[502].X = 5.02
	pointsOfPolynomialPlot[502].Y = -2_135.184

	pointsOfPolynomialPlot[503].X = 5.03
	pointsOfPolynomialPlot[503].Y = -2_144.217

	pointsOfPolynomialPlot[504].X = 5.04
	pointsOfPolynomialPlot[504].Y = -2_153.273

	pointsOfPolynomialPlot[505].X = 5.05
	pointsOfPolynomialPlot[505].Y = -2_162.353

	pointsOfPolynomialPlot[506].X = 5.06
	pointsOfPolynomialPlot[506].Y = -2_171.456

	pointsOfPolynomialPlot[507].X = 5.07
	pointsOfPolynomialPlot[507].Y = -2_180.582

	pointsOfPolynomialPlot[508].X = 5.08
	pointsOfPolynomialPlot[508].Y = -2_189.731

	pointsOfPolynomialPlot[509].X = 5.09
	pointsOfPolynomialPlot[509].Y = -2_198.904

	pointsOfPolynomialPlot[510].X = 5.1
	pointsOfPolynomialPlot[510].Y = -2_208.099

	pointsOfPolynomialPlot[511].X = 5.11
	pointsOfPolynomialPlot[511].Y = -2_217.319

	pointsOfPolynomialPlot[512].X = 5.12
	pointsOfPolynomialPlot[512].Y = -2_226.561

	pointsOfPolynomialPlot[513].X = 5.13
	pointsOfPolynomialPlot[513].Y = -2_235.826

	pointsOfPolynomialPlot[514].X = 5.14
	pointsOfPolynomialPlot[514].Y = -2_245.115

	pointsOfPolynomialPlot[515].X = 5.15
	pointsOfPolynomialPlot[515].Y = -2_254.427

	pointsOfPolynomialPlot[516].X = 5.16
	pointsOfPolynomialPlot[516].Y = -2_263.762

	pointsOfPolynomialPlot[517].X = 5.17
	pointsOfPolynomialPlot[517].Y = -2_273.121

	pointsOfPolynomialPlot[518].X = 5.18
	pointsOfPolynomialPlot[518].Y = -2_282.503

	pointsOfPolynomialPlot[519].X = 5.19
	pointsOfPolynomialPlot[519].Y = -2_291.907

	pointsOfPolynomialPlot[520].X = 5.2
	pointsOfPolynomialPlot[520].Y = -2_301.335

	pointsOfPolynomialPlot[521].X = 5.21
	pointsOfPolynomialPlot[521].Y = -2_310.787

	pointsOfPolynomialPlot[522].X = 5.22
	pointsOfPolynomialPlot[522].Y = -2_320.261

	pointsOfPolynomialPlot[523].X = 5.23
	pointsOfPolynomialPlot[523].Y = -2_329.759

	pointsOfPolynomialPlot[524].X = 5.24
	pointsOfPolynomialPlot[524].Y = -2_339.28

	pointsOfPolynomialPlot[525].X = 5.25
	pointsOfPolynomialPlot[525].Y = -2_348.824

	pointsOfPolynomialPlot[526].X = 5.26
	pointsOfPolynomialPlot[526].Y = -2_358.391

	pointsOfPolynomialPlot[527].X = 5.27
	pointsOfPolynomialPlot[527].Y = -2_367.981

	pointsOfPolynomialPlot[528].X = 5.28
	pointsOfPolynomialPlot[528].Y = -2_377.595

	pointsOfPolynomialPlot[529].X = 5.29
	pointsOfPolynomialPlot[529].Y = -2_387.232

	pointsOfPolynomialPlot[530].X = 5.3
	pointsOfPolynomialPlot[530].Y = -2_396.891

	pointsOfPolynomialPlot[531].X = 5.31
	pointsOfPolynomialPlot[531].Y = -2_406.574

	pointsOfPolynomialPlot[532].X = 5.32
	pointsOfPolynomialPlot[532].Y = -2_416.281

	pointsOfPolynomialPlot[533].X = 5.33
	pointsOfPolynomialPlot[533].Y = -2_426.01

	pointsOfPolynomialPlot[534].X = 5.34
	pointsOfPolynomialPlot[534].Y = -2_435.762

	pointsOfPolynomialPlot[535].X = 5.35
	pointsOfPolynomialPlot[535].Y = -2_445.538

	pointsOfPolynomialPlot[536].X = 5.36
	pointsOfPolynomialPlot[536].Y = -2_455.336

	pointsOfPolynomialPlot[537].X = 5.37
	pointsOfPolynomialPlot[537].Y = -2_465.158

	pointsOfPolynomialPlot[538].X = 5.38
	pointsOfPolynomialPlot[538].Y = -2_475.003

	pointsOfPolynomialPlot[539].X = 5.39
	pointsOfPolynomialPlot[539].Y = -2_484.871

	pointsOfPolynomialPlot[540].X = 5.4
	pointsOfPolynomialPlot[540].Y = -2_494.761

	pointsOfPolynomialPlot[541].X = 5.41
	pointsOfPolynomialPlot[541].Y = -2_504.675

	pointsOfPolynomialPlot[542].X = 5.42
	pointsOfPolynomialPlot[542].Y = -2_514.612

	pointsOfPolynomialPlot[543].X = 5.43
	pointsOfPolynomialPlot[543].Y = -2_524.572

	pointsOfPolynomialPlot[544].X = 5.44
	pointsOfPolynomialPlot[544].Y = -2_534.556

	pointsOfPolynomialPlot[545].X = 5.45
	pointsOfPolynomialPlot[545].Y = -2_544.562

	pointsOfPolynomialPlot[546].X = 5.46
	pointsOfPolynomialPlot[546].Y = -2_554.591

	pointsOfPolynomialPlot[547].X = 5.47
	pointsOfPolynomialPlot[547].Y = -2_564.643

	pointsOfPolynomialPlot[548].X = 5.48
	pointsOfPolynomialPlot[548].Y = -2_574.718

	pointsOfPolynomialPlot[549].X = 5.49
	pointsOfPolynomialPlot[549].Y = -2_584.816

	pointsOfPolynomialPlot[550].X = 5.5
	pointsOfPolynomialPlot[550].Y = -2_594.937

	pointsOfPolynomialPlot[551].X = 5.51
	pointsOfPolynomialPlot[551].Y = -2_605.081

	pointsOfPolynomialPlot[552].X = 5.52
	pointsOfPolynomialPlot[552].Y = -2_615.248

	pointsOfPolynomialPlot[553].X = 5.53
	pointsOfPolynomialPlot[553].Y = -2_625.438

	pointsOfPolynomialPlot[554].X = 5.54
	pointsOfPolynomialPlot[554].Y = -2_635.651

	pointsOfPolynomialPlot[555].X = 5.55
	pointsOfPolynomialPlot[555].Y = -2_645.886

	pointsOfPolynomialPlot[556].X = 5.56
	pointsOfPolynomialPlot[556].Y = -2_656.145

	pointsOfPolynomialPlot[557].X = 5.57
	pointsOfPolynomialPlot[557].Y = -2_666.426

	pointsOfPolynomialPlot[558].X = 5.58
	pointsOfPolynomialPlot[558].Y = -2_676.73

	pointsOfPolynomialPlot[559].X = 5.59
	pointsOfPolynomialPlot[559].Y = -2_686.058

	pointsOfPolynomialPlot[560].X = 5.6
	pointsOfPolynomialPlot[560].Y = -2_697.407

	pointsOfPolynomialPlot[561].X = 5.61
	pointsOfPolynomialPlot[561].Y = -2_707.78

	pointsOfPolynomialPlot[562].X = 5.62
	pointsOfPolynomialPlot[562].Y = -2_718.176

	pointsOfPolynomialPlot[563].X = 5.63
	pointsOfPolynomialPlot[563].Y = -2_728.594

	pointsOfPolynomialPlot[564].X = 5.64
	pointsOfPolynomialPlot[564].Y = -2_739.035

	pointsOfPolynomialPlot[565].X = 5.65
	pointsOfPolynomialPlot[565].Y = -2_749.499

	pointsOfPolynomialPlot[566].X = 5.66
	pointsOfPolynomialPlot[566].Y = -2_759.986

	pointsOfPolynomialPlot[567].X = 5.67
	pointsOfPolynomialPlot[567].Y = -2_770.495

	pointsOfPolynomialPlot[568].X = 5.68
	pointsOfPolynomialPlot[568].Y = -2_781.027

	pointsOfPolynomialPlot[569].X = 5.69
	pointsOfPolynomialPlot[569].Y = -2_791.582

	pointsOfPolynomialPlot[570].X = 5.7
	pointsOfPolynomialPlot[570].Y = -2_802.159

	pointsOfPolynomialPlot[571].X = 5.71
	pointsOfPolynomialPlot[571].Y = -2_812.76

	pointsOfPolynomialPlot[572].X = 5.72
	pointsOfPolynomialPlot[572].Y = -2_823.382

	pointsOfPolynomialPlot[573].X = 5.73
	pointsOfPolynomialPlot[573].Y = -2_834.028

	pointsOfPolynomialPlot[574].X = 5.74
	pointsOfPolynomialPlot[574].Y = -2_844.696

	pointsOfPolynomialPlot[575].X = 5.75
	pointsOfPolynomialPlot[575].Y = -2_855.386

	pointsOfPolynomialPlot[576].X = 5.76
	pointsOfPolynomialPlot[576].Y = -2_866.099

	pointsOfPolynomialPlot[577].X = 5.77
	pointsOfPolynomialPlot[577].Y = -2_876.835

	pointsOfPolynomialPlot[578].X = 5.78
	pointsOfPolynomialPlot[578].Y = -2_887.593

	pointsOfPolynomialPlot[579].X = 5.79
	pointsOfPolynomialPlot[579].Y = -2_898.374

	pointsOfPolynomialPlot[580].X = 5.8
	pointsOfPolynomialPlot[580].Y = -2_909.177

	pointsOfPolynomialPlot[581].X = 5.81
	pointsOfPolynomialPlot[581].Y = -2_920.003

	pointsOfPolynomialPlot[582].X = 5.82
	pointsOfPolynomialPlot[582].Y = -2_930.851

	pointsOfPolynomialPlot[583].X = 5.83
	pointsOfPolynomialPlot[583].Y = -2_941.722

	pointsOfPolynomialPlot[584].X = 5.84
	pointsOfPolynomialPlot[584].Y = -2_952.615

	pointsOfPolynomialPlot[585].X = 5.85
	pointsOfPolynomialPlot[585].Y = -2_963.531

	pointsOfPolynomialPlot[586].X = 5.86
	pointsOfPolynomialPlot[586].Y = -2_974.469

	pointsOfPolynomialPlot[587].X = 5.87
	pointsOfPolynomialPlot[587].Y = -2_985.429

	pointsOfPolynomialPlot[588].X = 5.88
	pointsOfPolynomialPlot[588].Y = -2_996.411

	pointsOfPolynomialPlot[589].X = 5.89
	pointsOfPolynomialPlot[589].Y = -3_007.416

	pointsOfPolynomialPlot[590].X = 5.9
	pointsOfPolynomialPlot[590].Y = -3_018.443

	pointsOfPolynomialPlot[591].X = 5.91
	pointsOfPolynomialPlot[591].Y = -3_029.493

	pointsOfPolynomialPlot[592].X = 5.92
	pointsOfPolynomialPlot[592].Y = -3_040.565

	pointsOfPolynomialPlot[593].X = 5.93
	pointsOfPolynomialPlot[593].Y = -3_051.659

	pointsOfPolynomialPlot[594].X = 5.94
	pointsOfPolynomialPlot[594].Y = -3_062.775

	pointsOfPolynomialPlot[595].X = 5.95
	pointsOfPolynomialPlot[595].Y = -3_073.913

	pointsOfPolynomialPlot[596].X = 5.96
	pointsOfPolynomialPlot[596].Y = -3_085.074

	pointsOfPolynomialPlot[597].X = 5.97
	pointsOfPolynomialPlot[597].Y = -3_096.256

	pointsOfPolynomialPlot[598].X = 5.98
	pointsOfPolynomialPlot[598].Y = -3_107.461

	pointsOfPolynomialPlot[599].X = 5.99
	pointsOfPolynomialPlot[599].Y = -3_118.688

	pointsOfPolynomialPlot[600].X = 6.0
	pointsOfPolynomialPlot[600].Y = -3_129.937

	pointsOfPolynomialPlot[601].X = 6.01
	pointsOfPolynomialPlot[601].Y = -3_141.208

	pointsOfPolynomialPlot[602].X = 6.02
	pointsOfPolynomialPlot[602].Y = -3_152.501

	pointsOfPolynomialPlot[603].X = 6.03
	pointsOfPolynomialPlot[603].Y = -3_163.816

	pointsOfPolynomialPlot[604].X = 6.04
	pointsOfPolynomialPlot[604].Y = -3_175.153

	pointsOfPolynomialPlot[605].X = 6.05
	pointsOfPolynomialPlot[605].Y = -3_186.512

	pointsOfPolynomialPlot[606].X = 6.06
	pointsOfPolynomialPlot[606].Y = -3_197.893

	pointsOfPolynomialPlot[607].X = 6.07
	pointsOfPolynomialPlot[607].Y = -3_209.296

	pointsOfPolynomialPlot[608].X = 6.08
	pointsOfPolynomialPlot[608].Y = -3_220.721

	pointsOfPolynomialPlot[609].X = 6.09
	pointsOfPolynomialPlot[609].Y = -3_232.167

	pointsOfPolynomialPlot[610].X = 6.1
	pointsOfPolynomialPlot[610].Y = -3_243.635

	pointsOfPolynomialPlot[611].X = 6.11
	pointsOfPolynomialPlot[611].Y = -3_255.126

	pointsOfPolynomialPlot[612].X = 6.12
	pointsOfPolynomialPlot[612].Y = -3_266.638

	pointsOfPolynomialPlot[613].X = 6.13
	pointsOfPolynomialPlot[613].Y = -3_278.171

	pointsOfPolynomialPlot[614].X = 6.14
	pointsOfPolynomialPlot[614].Y = -3_289.727

	pointsOfPolynomialPlot[615].X = 6.15
	pointsOfPolynomialPlot[615].Y = -3_301.304

	pointsOfPolynomialPlot[616].X = 6.16
	pointsOfPolynomialPlot[616].Y = -3_312.902

	pointsOfPolynomialPlot[617].X = 6.17
	pointsOfPolynomialPlot[617].Y = -3_324.523

	pointsOfPolynomialPlot[618].X = 6.18
	pointsOfPolynomialPlot[618].Y = -3_336.165

	pointsOfPolynomialPlot[619].X = 6.19
	pointsOfPolynomialPlot[619].Y = -3_347.828

	pointsOfPolynomialPlot[620].X = 6.2
	pointsOfPolynomialPlot[620].Y = -3_359.513

	pointsOfPolynomialPlot[621].X = 6.21
	pointsOfPolynomialPlot[621].Y = -3_371.22

	pointsOfPolynomialPlot[622].X = 6.22
	pointsOfPolynomialPlot[622].Y = -3_382.948

	pointsOfPolynomialPlot[623].X = 6.23
	pointsOfPolynomialPlot[623].Y = -3_394.698

	pointsOfPolynomialPlot[624].X = 6.24
	pointsOfPolynomialPlot[624].Y = -3_406.469

	pointsOfPolynomialPlot[625].X = 6.25
	pointsOfPolynomialPlot[625].Y = -3_418.261

	pointsOfPolynomialPlot[626].X = 6.26
	pointsOfPolynomialPlot[626].Y = -3_430.075

	pointsOfPolynomialPlot[627].X = 6.27
	pointsOfPolynomialPlot[627].Y = -3_441.91

	pointsOfPolynomialPlot[628].X = 6.28
	pointsOfPolynomialPlot[628].Y = -3_453.767

	pointsOfPolynomialPlot[629].X = 6.29
	pointsOfPolynomialPlot[629].Y = -3_465.644

	pointsOfPolynomialPlot[630].X = 6.3
	pointsOfPolynomialPlot[630].Y = -3_477.543

	pointsOfPolynomialPlot[631].X = 6.31
	pointsOfPolynomialPlot[631].Y = -3_489.464

	pointsOfPolynomialPlot[632].X = 6.32
	pointsOfPolynomialPlot[632].Y = -3_501.405

	pointsOfPolynomialPlot[633].X = 6.33
	pointsOfPolynomialPlot[633].Y = -3_513.368

	pointsOfPolynomialPlot[634].X = 6.34
	pointsOfPolynomialPlot[634].Y = -3_525.351

	pointsOfPolynomialPlot[635].X = 6.35
	pointsOfPolynomialPlot[635].Y = -3_537.356

	pointsOfPolynomialPlot[636].X = 6.36
	pointsOfPolynomialPlot[636].Y = -3_549.382

	pointsOfPolynomialPlot[637].X = 6.37
	pointsOfPolynomialPlot[637].Y = -3_561.429

	pointsOfPolynomialPlot[638].X = 6.38
	pointsOfPolynomialPlot[638].Y = -3_573.497

	pointsOfPolynomialPlot[639].X = 6.39
	pointsOfPolynomialPlot[639].Y = -3_585.586

	pointsOfPolynomialPlot[640].X = 6.4
	pointsOfPolynomialPlot[640].Y = -3_597.695

	pointsOfPolynomialPlot[641].X = 6.41
	pointsOfPolynomialPlot[641].Y = -3_609.826

	pointsOfPolynomialPlot[642].X = 6.42
	pointsOfPolynomialPlot[642].Y = -3_621.978

	pointsOfPolynomialPlot[643].X = 6.43
	pointsOfPolynomialPlot[643].Y = -3_634.15

	pointsOfPolynomialPlot[644].X = 6.44
	pointsOfPolynomialPlot[644].Y = -3_646.343

	pointsOfPolynomialPlot[645].X = 6.45
	pointsOfPolynomialPlot[645].Y = -3_658.557

	pointsOfPolynomialPlot[646].X = 6.46
	pointsOfPolynomialPlot[646].Y = -3_670.792

	pointsOfPolynomialPlot[647].X = 6.47
	pointsOfPolynomialPlot[647].Y = -3_683.047

	pointsOfPolynomialPlot[648].X = 6.48
	pointsOfPolynomialPlot[648].Y = -3_695.323

	pointsOfPolynomialPlot[649].X = 6.49
	pointsOfPolynomialPlot[649].Y = -3_707.62

	pointsOfPolynomialPlot[650].X = 6.5
	pointsOfPolynomialPlot[650].Y = -3_719.937

	pointsOfPolynomialPlot[651].X = 6.51
	pointsOfPolynomialPlot[651].Y = -3_732.275

	pointsOfPolynomialPlot[652].X = 6.52
	pointsOfPolynomialPlot[652].Y = -3_744.633

	pointsOfPolynomialPlot[653].X = 6.53
	pointsOfPolynomialPlot[653].Y = -3_757.012

	pointsOfPolynomialPlot[654].X = 6.54
	pointsOfPolynomialPlot[654].Y = -3_769.411

	pointsOfPolynomialPlot[655].X = 6.55
	pointsOfPolynomialPlot[655].Y = -3_781.831

	pointsOfPolynomialPlot[656].X = 6.56
	pointsOfPolynomialPlot[656].Y = -3_794.271

	pointsOfPolynomialPlot[657].X = 6.57
	pointsOfPolynomialPlot[657].Y = -3_806.731

	pointsOfPolynomialPlot[658].X = 6.58
	pointsOfPolynomialPlot[658].Y = -3_819.212

	pointsOfPolynomialPlot[659].X = 6.59
	pointsOfPolynomialPlot[659].Y = -3_831.712

	pointsOfPolynomialPlot[660].X = 6.6
	pointsOfPolynomialPlot[660].Y = -3_844.233

	pointsOfPolynomialPlot[661].X = 6.61
	pointsOfPolynomialPlot[661].Y = -3_856.775

	pointsOfPolynomialPlot[662].X = 6.62
	pointsOfPolynomialPlot[662].Y = -3_869.336

	pointsOfPolynomialPlot[663].X = 6.63
	pointsOfPolynomialPlot[663].Y = -3_881.918

	pointsOfPolynomialPlot[664].X = 6.64
	pointsOfPolynomialPlot[664].Y = -3_894.519

	pointsOfPolynomialPlot[665].X = 6.65
	pointsOfPolynomialPlot[665].Y = -3_907.141

	pointsOfPolynomialPlot[666].X = 6.66
	pointsOfPolynomialPlot[666].Y = -3_919.782

	pointsOfPolynomialPlot[667].X = 6.67
	pointsOfPolynomialPlot[667].Y = -3_932.444

	pointsOfPolynomialPlot[668].X = 6.68
	pointsOfPolynomialPlot[668].Y = -3_945.125

	pointsOfPolynomialPlot[669].X = 6.69
	pointsOfPolynomialPlot[669].Y = -3_957.826

	pointsOfPolynomialPlot[670].X = 6.7
	pointsOfPolynomialPlot[670].Y = -3_970.547

	pointsOfPolynomialPlot[671].X = 6.71
	pointsOfPolynomialPlot[671].Y = -3_983.288

	pointsOfPolynomialPlot[672].X = 6.72
	pointsOfPolynomialPlot[672].Y = -3_996.049

	pointsOfPolynomialPlot[673].X = 6.73
	pointsOfPolynomialPlot[673].Y = -4_008.829

	pointsOfPolynomialPlot[674].X = 6.74
	pointsOfPolynomialPlot[674].Y = -4_021.629

	pointsOfPolynomialPlot[675].X = 6.75
	pointsOfPolynomialPlot[675].Y = -4_034.449

	pointsOfPolynomialPlot[676].X = 6.76
	pointsOfPolynomialPlot[676].Y = -4_047.288

	pointsOfPolynomialPlot[677].X = 6.77
	pointsOfPolynomialPlot[677].Y = -4_060.147

	pointsOfPolynomialPlot[678].X = 6.78
	pointsOfPolynomialPlot[678].Y = -4_073.025

	pointsOfPolynomialPlot[679].X = 6.79
	pointsOfPolynomialPlot[679].Y = -4_085.922

	pointsOfPolynomialPlot[680].X = 6.8
	pointsOfPolynomialPlot[680].Y = -4_098.839

	pointsOfPolynomialPlot[681].X = 6.81
	pointsOfPolynomialPlot[681].Y = -4_111.776

	pointsOfPolynomialPlot[682].X = 6.82
	pointsOfPolynomialPlot[682].Y = -4_124.732

	pointsOfPolynomialPlot[683].X = 6.83
	pointsOfPolynomialPlot[683].Y = -4_137.707

	pointsOfPolynomialPlot[684].X = 6.84
	pointsOfPolynomialPlot[684].Y = -4_150.701

	pointsOfPolynomialPlot[685].X = 6.85
	pointsOfPolynomialPlot[685].Y = -4_163.714

	pointsOfPolynomialPlot[686].X = 6.86
	pointsOfPolynomialPlot[686].Y = -4_176.747

	pointsOfPolynomialPlot[687].X = 6.87
	pointsOfPolynomialPlot[687].Y = -4_189.798

	pointsOfPolynomialPlot[688].X = 6.88
	pointsOfPolynomialPlot[688].Y = -4_202.869

	pointsOfPolynomialPlot[689].X = 6.89
	pointsOfPolynomialPlot[689].Y = -4_215.959

	pointsOfPolynomialPlot[690].X = 6.9
	pointsOfPolynomialPlot[690].Y = -4_229.067

	pointsOfPolynomialPlot[691].X = 6.91
	pointsOfPolynomialPlot[691].Y = -4_242.195

	pointsOfPolynomialPlot[692].X = 6.92
	pointsOfPolynomialPlot[692].Y = -4_255.342

	pointsOfPolynomialPlot[693].X = 6.93
	pointsOfPolynomialPlot[693].Y = -4_268.507

	pointsOfPolynomialPlot[694].X = 6.94
	pointsOfPolynomialPlot[694].Y = -4_281.691

	pointsOfPolynomialPlot[695].X = 6.95
	pointsOfPolynomialPlot[695].Y = -4_294.894

	pointsOfPolynomialPlot[696].X = 6.96
	pointsOfPolynomialPlot[696].Y = -4_308.115

	pointsOfPolynomialPlot[697].X = 6.97
	pointsOfPolynomialPlot[697].Y = -4_321.355

	pointsOfPolynomialPlot[698].X = 6.98
	pointsOfPolynomialPlot[698].Y = -4_334.614

	pointsOfPolynomialPlot[699].X = 6.99
	pointsOfPolynomialPlot[699].Y = -4_347.891

	pointsOfPolynomialPlot[700].X = 7.0
	pointsOfPolynomialPlot[700].Y = -4_361.187

	pointsOfPolynomialPlot[701].X = 7.01
	pointsOfPolynomialPlot[701].Y = -4_374.501

	pointsOfPolynomialPlot[702].X = 7.02
	pointsOfPolynomialPlot[702].Y = -4_387.834

	pointsOfPolynomialPlot[703].X = 7.03
	pointsOfPolynomialPlot[703].Y = -4_401.185

	pointsOfPolynomialPlot[704].X = 7.04
	pointsOfPolynomialPlot[704].Y = -4_414.554

	pointsOfPolynomialPlot[705].X = 7.05
	pointsOfPolynomialPlot[705].Y = -4_427.942

	pointsOfPolynomialPlot[706].X = 7.06
	pointsOfPolynomialPlot[706].Y = -4_441.347

	pointsOfPolynomialPlot[707].X = 7.07
	pointsOfPolynomialPlot[707].Y = -4_454.771

	pointsOfPolynomialPlot[708].X = 7.08
	pointsOfPolynomialPlot[708].Y = -4_468.213

	pointsOfPolynomialPlot[709].X = 7.09
	pointsOfPolynomialPlot[709].Y = -4_481.673

	pointsOfPolynomialPlot[710].X = 7.1
	pointsOfPolynomialPlot[710].Y = -4_495.151

	pointsOfPolynomialPlot[711].X = 7.11
	pointsOfPolynomialPlot[711].Y = -4_508.647

	pointsOfPolynomialPlot[712].X = 7.12
	pointsOfPolynomialPlot[712].Y = -4_522.161

	pointsOfPolynomialPlot[713].X = 7.13
	pointsOfPolynomialPlot[713].Y = -4_535.693

	pointsOfPolynomialPlot[714].X = 7.14
	pointsOfPolynomialPlot[714].Y = -4_549.243

	pointsOfPolynomialPlot[715].X = 7.15
	pointsOfPolynomialPlot[715].Y = -4_562.81

	pointsOfPolynomialPlot[716].X = 7.16
	pointsOfPolynomialPlot[716].Y = -4_576.395

	pointsOfPolynomialPlot[717].X = 7.17
	pointsOfPolynomialPlot[717].Y = -4_589.998

	pointsOfPolynomialPlot[718].X = 7.18
	pointsOfPolynomialPlot[718].Y = -4_603.618

	pointsOfPolynomialPlot[719].X = 7.19
	pointsOfPolynomialPlot[719].Y = -4_617.256

	pointsOfPolynomialPlot[720].X = 7.2
	pointsOfPolynomialPlot[720].Y = -4_630.911

	pointsOfPolynomialPlot[721].X = 7.21
	pointsOfPolynomialPlot[721].Y = -4_644.584

	pointsOfPolynomialPlot[722].X = 7.22
	pointsOfPolynomialPlot[722].Y = -4_658.274

	pointsOfPolynomialPlot[723].X = 7.23
	pointsOfPolynomialPlot[723].Y = -4_671.982

	pointsOfPolynomialPlot[724].X = 7.24
	pointsOfPolynomialPlot[724].Y = -4_685.707

	pointsOfPolynomialPlot[725].X = 7.25
	pointsOfPolynomialPlot[725].Y = -4_699.449

	pointsOfPolynomialPlot[726].X = 7.26
	pointsOfPolynomialPlot[726].Y = -4_713.208

	pointsOfPolynomialPlot[727].X = 7.27
	pointsOfPolynomialPlot[727].Y = -4_726.984

	pointsOfPolynomialPlot[728].X = 7.28
	pointsOfPolynomialPlot[728].Y = -4_740.778

	pointsOfPolynomialPlot[729].X = 7.29
	pointsOfPolynomialPlot[729].Y = -4_754.588

	pointsOfPolynomialPlot[730].X = 7.3
	pointsOfPolynomialPlot[730].Y = -4_768.415

	pointsOfPolynomialPlot[731].X = 7.31
	pointsOfPolynomialPlot[731].Y = -4_782.26

	pointsOfPolynomialPlot[732].X = 7.32
	pointsOfPolynomialPlot[732].Y = -4_796.121

	pointsOfPolynomialPlot[733].X = 7.33
	pointsOfPolynomialPlot[733].Y = -4_809.999

	pointsOfPolynomialPlot[734].X = 7.34
	pointsOfPolynomialPlot[734].Y = -4_823.893

	pointsOfPolynomialPlot[735].X = 7.35
	pointsOfPolynomialPlot[735].Y = -4_837.805

	pointsOfPolynomialPlot[736].X = 7.36
	pointsOfPolynomialPlot[736].Y = -4_851.733

	pointsOfPolynomialPlot[737].X = 7.37
	pointsOfPolynomialPlot[737].Y = -4_865.677

	pointsOfPolynomialPlot[738].X = 7.38
	pointsOfPolynomialPlot[738].Y = -4_879.638

	pointsOfPolynomialPlot[739].X = 7.39
	pointsOfPolynomialPlot[739].Y = -4_893.616

	pointsOfPolynomialPlot[740].X = 7.4
	pointsOfPolynomialPlot[740].Y = -4_907.609

	pointsOfPolynomialPlot[741].X = 7.41
	pointsOfPolynomialPlot[741].Y = -4_921.62

	pointsOfPolynomialPlot[742].X = 7.42
	pointsOfPolynomialPlot[742].Y = -4_935.646

	pointsOfPolynomialPlot[743].X = 7.43
	pointsOfPolynomialPlot[743].Y = -4_949.689

	pointsOfPolynomialPlot[744].X = 7.44
	pointsOfPolynomialPlot[744].Y = -4_963.748

	pointsOfPolynomialPlot[745].X = 7.45
	pointsOfPolynomialPlot[745].Y = -4_977.823

	pointsOfPolynomialPlot[746].X = 7.46
	pointsOfPolynomialPlot[746].Y = -4_991.914

	pointsOfPolynomialPlot[747].X = 7.47
	pointsOfPolynomialPlot[747].Y = -5_006.021

	pointsOfPolynomialPlot[748].X = 7.48
	pointsOfPolynomialPlot[748].Y = -5_020.144

	pointsOfPolynomialPlot[749].X = 7.49
	pointsOfPolynomialPlot[749].Y = -5_034.282

	pointsOfPolynomialPlot[750].X = 7.5
	pointsOfPolynomialPlot[750].Y = -5_048.437

	pointsOfPolynomialPlot[751].X = 7.51
	pointsOfPolynomialPlot[751].Y = -5_062.607

	pointsOfPolynomialPlot[752].X = 7.52
	pointsOfPolynomialPlot[752].Y = -5_076.793

	pointsOfPolynomialPlot[753].X = 7.53
	pointsOfPolynomialPlot[753].Y = -5_090.995

	pointsOfPolynomialPlot[754].X = 7.54
	pointsOfPolynomialPlot[754].Y = -5_105.212

	pointsOfPolynomialPlot[755].X = 7.55
	pointsOfPolynomialPlot[755].Y = -5_119.445

	pointsOfPolynomialPlot[756].X = 7.56
	pointsOfPolynomialPlot[756].Y = -5_133.693

	pointsOfPolynomialPlot[757].X = 7.57
	pointsOfPolynomialPlot[757].Y = -5_147.957

	pointsOfPolynomialPlot[758].X = 7.58
	pointsOfPolynomialPlot[758].Y = -5_162.236

	pointsOfPolynomialPlot[759].X = 7.59
	pointsOfPolynomialPlot[759].Y = -5_176.53

	pointsOfPolynomialPlot[760].X = 7.6
	pointsOfPolynomialPlot[760].Y = -5_190.839

	pointsOfPolynomialPlot[761].X = 7.61
	pointsOfPolynomialPlot[761].Y = -5_205.164

	pointsOfPolynomialPlot[762].X = 7.62
	pointsOfPolynomialPlot[762].Y = -5_219.504

	pointsOfPolynomialPlot[763].X = 7.63
	pointsOfPolynomialPlot[763].Y = -5_233.858

	pointsOfPolynomialPlot[764].X = 7.64
	pointsOfPolynomialPlot[764].Y = -5_248.228

	pointsOfPolynomialPlot[765].X = 7.65
	pointsOfPolynomialPlot[765].Y = -5_262.612

	pointsOfPolynomialPlot[766].X = 7.66
	pointsOfPolynomialPlot[766].Y = -5_277.011

	pointsOfPolynomialPlot[767].X = 7.67
	pointsOfPolynomialPlot[767].Y = -5_291.425

	pointsOfPolynomialPlot[768].X = 7.68
	pointsOfPolynomialPlot[768].Y = -5_305.854

	pointsOfPolynomialPlot[769].X = 7.69
	pointsOfPolynomialPlot[769].Y = -5_320.298

	pointsOfPolynomialPlot[770].X = 7.7
	pointsOfPolynomialPlot[770].Y = -5_334.755

	pointsOfPolynomialPlot[771].X = 7.71
	pointsOfPolynomialPlot[771].Y = -5_349.228

	pointsOfPolynomialPlot[772].X = 7.72
	pointsOfPolynomialPlot[772].Y = -5_363.715

	pointsOfPolynomialPlot[773].X = 7.73
	pointsOfPolynomialPlot[773].Y = -5_378.216

	pointsOfPolynomialPlot[774].X = 7.74
	pointsOfPolynomialPlot[774].Y = -5_392.731

	pointsOfPolynomialPlot[775].X = 7.75
	pointsOfPolynomialPlot[775].Y = -5_407.261

	pointsOfPolynomialPlot[776].X = 7.76
	pointsOfPolynomialPlot[776].Y = -5_421.805

	pointsOfPolynomialPlot[777].X = 7.77
	pointsOfPolynomialPlot[777].Y = -5_436.363

	pointsOfPolynomialPlot[778].X = 7.78
	pointsOfPolynomialPlot[778].Y = -5_450.935

	pointsOfPolynomialPlot[779].X = 7.79
	pointsOfPolynomialPlot[779].Y = -5_465.521

	pointsOfPolynomialPlot[780].X = 7.8
	pointsOfPolynomialPlot[780].Y = -5_480.121

	pointsOfPolynomialPlot[781].X = 7.81
	pointsOfPolynomialPlot[781].Y = -5_494.735

	pointsOfPolynomialPlot[782].X = 7.82
	pointsOfPolynomialPlot[782].Y = -5_509.363

	pointsOfPolynomialPlot[783].X = 7.83
	pointsOfPolynomialPlot[783].Y = -5_524.004

	pointsOfPolynomialPlot[784].X = 7.84
	pointsOfPolynomialPlot[784].Y = -5_538.659

	pointsOfPolynomialPlot[785].X = 7.85
	pointsOfPolynomialPlot[785].Y = -5_553.328

	pointsOfPolynomialPlot[786].X = 7.86
	pointsOfPolynomialPlot[786].Y = -5_568.01

	pointsOfPolynomialPlot[787].X = 7.87
	pointsOfPolynomialPlot[787].Y = -5_582.705

	pointsOfPolynomialPlot[788].X = 7.88
	pointsOfPolynomialPlot[788].Y = -5_597.414

	pointsOfPolynomialPlot[789].X = 7.89
	pointsOfPolynomialPlot[789].Y = -5_612.136

	pointsOfPolynomialPlot[790].X = 7.9
	pointsOfPolynomialPlot[790].Y = -5_626.871

	pointsOfPolynomialPlot[791].X = 7.91
	pointsOfPolynomialPlot[791].Y = -5_641.62

	pointsOfPolynomialPlot[792].X = 7.92
	pointsOfPolynomialPlot[792].Y = -5_656.382

	pointsOfPolynomialPlot[793].X = 7.93
	pointsOfPolynomialPlot[793].Y = -5_671.156

	pointsOfPolynomialPlot[794].X = 7.94
	pointsOfPolynomialPlot[794].Y = -5_685.944

	pointsOfPolynomialPlot[795].X = 7.95
	pointsOfPolynomialPlot[795].Y = -5_700.744

	pointsOfPolynomialPlot[796].X = 7.96
	pointsOfPolynomialPlot[796].Y = -5_715.557

	pointsOfPolynomialPlot[797].X = 7.97
	pointsOfPolynomialPlot[797].Y = -5_730.383

	pointsOfPolynomialPlot[798].X = 7.98
	pointsOfPolynomialPlot[798].Y = -5_745.222

	pointsOfPolynomialPlot[799].X = 7.99
	pointsOfPolynomialPlot[799].Y = -5_760.073

	pointsOfPolynomialPlot[800].X = 8.0
	pointsOfPolynomialPlot[800].Y = -5_774.937

	pointsOfPolynomialPlot[801].X = 8.01
	pointsOfPolynomialPlot[801].Y = -5_789.813

	pointsOfPolynomialPlot[802].X = 8.02
	pointsOfPolynomialPlot[802].Y = -5_804.702

	pointsOfPolynomialPlot[803].X = 8.03
	pointsOfPolynomialPlot[803].Y = -5_819.603

	pointsOfPolynomialPlot[804].X = 8.04
	pointsOfPolynomialPlot[804].Y = -5_834.516

	pointsOfPolynomialPlot[805].X = 8.05
	pointsOfPolynomialPlot[805].Y = -5_849.441

	pointsOfPolynomialPlot[806].X = 8.06
	pointsOfPolynomialPlot[806].Y = -5_864.379

	pointsOfPolynomialPlot[807].X = 8.07
	pointsOfPolynomialPlot[807].Y = -5_879.328

	pointsOfPolynomialPlot[808].X = 8.08
	pointsOfPolynomialPlot[808].Y = -5_894.289

	pointsOfPolynomialPlot[809].X = 8.09
	pointsOfPolynomialPlot[809].Y = -5_909.262

	pointsOfPolynomialPlot[810].X = 8.1
	pointsOfPolynomialPlot[810].Y = -5_925.247

	pointsOfPolynomialPlot[811].X = 8.11
	pointsOfPolynomialPlot[811].Y = -5_939.244

	pointsOfPolynomialPlot[812].X = 8.12
	pointsOfPolynomialPlot[812].Y = -5_954.253

	pointsOfPolynomialPlot[813].X = 8.13
	pointsOfPolynomialPlot[813].Y = -5_969.272

	pointsOfPolynomialPlot[814].X = 8.14
	pointsOfPolynomialPlot[814].Y = -5_984.304

	pointsOfPolynomialPlot[815].X = 8.15
	pointsOfPolynomialPlot[815].Y = -5_999.347

	pointsOfPolynomialPlot[816].X = 8.16
	pointsOfPolynomialPlot[816].Y = -6_014.401

	pointsOfPolynomialPlot[817].X = 8.17
	pointsOfPolynomialPlot[817].Y = -6_029.466

	pointsOfPolynomialPlot[818].X = 8.18
	pointsOfPolynomialPlot[818].Y = -6_044.543

	pointsOfPolynomialPlot[819].X = 8.19
	pointsOfPolynomialPlot[819].Y = -6_059.631

	pointsOfPolynomialPlot[820].X = 8.2
	pointsOfPolynomialPlot[820].Y = -6_074.729

	pointsOfPolynomialPlot[821].X = 8.21
	pointsOfPolynomialPlot[821].Y = -6_089.839

	pointsOfPolynomialPlot[822].X = 8.22
	pointsOfPolynomialPlot[822].Y = -6_104.96

	pointsOfPolynomialPlot[823].X = 8.23
	pointsOfPolynomialPlot[823].Y = -6_120.091

	pointsOfPolynomialPlot[824].X = 8.24
	pointsOfPolynomialPlot[824].Y = -6_135.233

	pointsOfPolynomialPlot[825].X = 8.25
	pointsOfPolynomialPlot[825].Y = -6_150.386

	pointsOfPolynomialPlot[826].X = 8.26
	pointsOfPolynomialPlot[826].Y = -6_165.55

	pointsOfPolynomialPlot[827].X = 8.27
	pointsOfPolynomialPlot[827].Y = -6_180.724

	pointsOfPolynomialPlot[828].X = 8.28
	pointsOfPolynomialPlot[828].Y = -6_195.908

	pointsOfPolynomialPlot[829].X = 8.29
	pointsOfPolynomialPlot[829].Y = -6_211.103

	pointsOfPolynomialPlot[830].X = 8.3
	pointsOfPolynomialPlot[830].Y = -6_226.307

	pointsOfPolynomialPlot[831].X = 8.31
	pointsOfPolynomialPlot[831].Y = -6_241.523

	pointsOfPolynomialPlot[832].X = 8.32
	pointsOfPolynomialPlot[832].Y = -6_256.748

	pointsOfPolynomialPlot[833].X = 8.33
	pointsOfPolynomialPlot[833].Y = -6_271.983

	pointsOfPolynomialPlot[834].X = 8.34
	pointsOfPolynomialPlot[834].Y = -6_287.228

	pointsOfPolynomialPlot[835].X = 8.35
	pointsOfPolynomialPlot[835].Y = -6_302.483

	pointsOfPolynomialPlot[836].X = 8.36
	pointsOfPolynomialPlot[836].Y = -6_317.748

	pointsOfPolynomialPlot[837].X = 8.37
	pointsOfPolynomialPlot[837].Y = -6_333.022

	pointsOfPolynomialPlot[838].X = 8.38
	pointsOfPolynomialPlot[838].Y = -6_348.307

	pointsOfPolynomialPlot[839].X = 8.39
	pointsOfPolynomialPlot[839].Y = -6_363.6

	pointsOfPolynomialPlot[840].X = 8.4
	pointsOfPolynomialPlot[840].Y = -6_378.903

	pointsOfPolynomialPlot[841].X = 8.41
	pointsOfPolynomialPlot[841].Y = -6_394.216

	pointsOfPolynomialPlot[842].X = 8.42
	pointsOfPolynomialPlot[842].Y = -6_409.538

	pointsOfPolynomialPlot[843].X = 8.43
	pointsOfPolynomialPlot[843].Y = -6_424.869

	pointsOfPolynomialPlot[844].X = 8.44
	pointsOfPolynomialPlot[844].Y = -6_440.209

	pointsOfPolynomialPlot[845].X = 8.45
	pointsOfPolynomialPlot[845].Y = -6_455.558

	pointsOfPolynomialPlot[846].X = 8.46
	pointsOfPolynomialPlot[846].Y = -6_470.916

	pointsOfPolynomialPlot[847].X = 8.47
	pointsOfPolynomialPlot[847].Y = -6_486.283

	pointsOfPolynomialPlot[848].X = 8.48
	pointsOfPolynomialPlot[848].Y = -6_501.659

	pointsOfPolynomialPlot[849].X = 8.49
	pointsOfPolynomialPlot[849].Y = -6_517.044

	pointsOfPolynomialPlot[850].X = 8.5
	pointsOfPolynomialPlot[850].Y = -6_532.437

	pointsOfPolynomialPlot[851].X = 8.51
	pointsOfPolynomialPlot[851].Y = -6_547.839

	pointsOfPolynomialPlot[852].X = 8.52
	pointsOfPolynomialPlot[852].Y = -6_563.249

	pointsOfPolynomialPlot[853].X = 8.53
	pointsOfPolynomialPlot[853].Y = -6_578.668

	pointsOfPolynomialPlot[854].X = 8.54
	pointsOfPolynomialPlot[854].Y = -6_594.095

	pointsOfPolynomialPlot[855].X = 8.55
	pointsOfPolynomialPlot[855].Y = -6_609.53

	pointsOfPolynomialPlot[856].X = 8.56
	pointsOfPolynomialPlot[856].Y = -6_624.973

	pointsOfPolynomialPlot[857].X = 8.57
	pointsOfPolynomialPlot[857].Y = -6_640.424

	pointsOfPolynomialPlot[858].X = 8.58
	pointsOfPolynomialPlot[858].Y = -6_655.883

	pointsOfPolynomialPlot[859].X = 8.59
	pointsOfPolynomialPlot[859].Y = -6_671.351

	pointsOfPolynomialPlot[860].X = 8.6
	pointsOfPolynomialPlot[860].Y = -6_686.825

	pointsOfPolynomialPlot[861].X = 8.61
	pointsOfPolynomialPlot[861].Y = -6_702.308

	pointsOfPolynomialPlot[862].X = 8.62
	pointsOfPolynomialPlot[862].Y = -6_717.798

	pointsOfPolynomialPlot[863].X = 8.63
	pointsOfPolynomialPlot[863].Y = -6_733.296

	pointsOfPolynomialPlot[864].X = 8.64
	pointsOfPolynomialPlot[864].Y = -6_748.801

	pointsOfPolynomialPlot[865].X = 8.65
	pointsOfPolynomialPlot[865].Y = -6_764.314

	pointsOfPolynomialPlot[866].X = 8.66
	pointsOfPolynomialPlot[866].Y = -6_779.833

	pointsOfPolynomialPlot[867].X = 8.67
	pointsOfPolynomialPlot[867].Y = -6_795.36

	pointsOfPolynomialPlot[868].X = 8.68
	pointsOfPolynomialPlot[868].Y = -6_810.894

	pointsOfPolynomialPlot[869].X = 8.69
	pointsOfPolynomialPlot[869].Y = -6_826.435

	pointsOfPolynomialPlot[870].X = 8.7
	pointsOfPolynomialPlot[870].Y = -6_841.983

	pointsOfPolynomialPlot[871].X = 8.71
	pointsOfPolynomialPlot[871].Y = -6_857.538

	pointsOfPolynomialPlot[872].X = 8.72
	pointsOfPolynomialPlot[872].Y = -6_873.1

	pointsOfPolynomialPlot[873].X = 8.73
	pointsOfPolynomialPlot[873].Y = -6_888.668

	pointsOfPolynomialPlot[874].X = 8.74
	pointsOfPolynomialPlot[874].Y = -6_904.243

	pointsOfPolynomialPlot[875].X = 8.75
	pointsOfPolynomialPlot[875].Y = -6_919.824

	pointsOfPolynomialPlot[876].X = 8.76
	pointsOfPolynomialPlot[876].Y = -6_935.411

	pointsOfPolynomialPlot[877].X = 8.77
	pointsOfPolynomialPlot[877].Y = -6_951.005

	pointsOfPolynomialPlot[878].X = 8.78
	pointsOfPolynomialPlot[878].Y = -6_966.605

	pointsOfPolynomialPlot[879].X = 8.79
	pointsOfPolynomialPlot[879].Y = -6_982.211

	pointsOfPolynomialPlot[880].X = 8.8
	pointsOfPolynomialPlot[880].Y = -6_997.823

	pointsOfPolynomialPlot[881].X = 8.81
	pointsOfPolynomialPlot[881].Y = -7_013.442

	pointsOfPolynomialPlot[882].X = 8.82
	pointsOfPolynomialPlot[882].Y = -7_029.065

	pointsOfPolynomialPlot[883].X = 8.83
	pointsOfPolynomialPlot[883].Y = -7_044.695

	pointsOfPolynomialPlot[884].X = 8.84
	pointsOfPolynomialPlot[884].Y = -7_060.33

	pointsOfPolynomialPlot[885].X = 8.85
	pointsOfPolynomialPlot[885].Y = -7_075.971

	pointsOfPolynomialPlot[886].X = 8.86
	pointsOfPolynomialPlot[886].Y = -7_091.617

	pointsOfPolynomialPlot[887].X = 8.87
	pointsOfPolynomialPlot[887].Y = -7_107.269

	pointsOfPolynomialPlot[888].X = 8.88
	pointsOfPolynomialPlot[888].Y = -7_122.926

	pointsOfPolynomialPlot[889].X = 8.89
	pointsOfPolynomialPlot[889].Y = -7_138.588

	pointsOfPolynomialPlot[890].X = 8.9
	pointsOfPolynomialPlot[890].Y = -7_154.255

	pointsOfPolynomialPlot[891].X = 8.91
	pointsOfPolynomialPlot[891].Y = -7_169.928

	pointsOfPolynomialPlot[892].X = 8.92
	pointsOfPolynomialPlot[892].Y = -7_185.605

	pointsOfPolynomialPlot[893].X = 8.93
	pointsOfPolynomialPlot[893].Y = -7_201.287

	pointsOfPolynomialPlot[894].X = 8.94
	pointsOfPolynomialPlot[894].Y = -7_216.973

	pointsOfPolynomialPlot[895].X = 8.95
	pointsOfPolynomialPlot[895].Y = -7_232.665

	pointsOfPolynomialPlot[896].X = 8.96
	pointsOfPolynomialPlot[896].Y = -7_248.36

	pointsOfPolynomialPlot[897].X = 8.97
	pointsOfPolynomialPlot[897].Y = -7_264.061

	pointsOfPolynomialPlot[898].X = 8.98
	pointsOfPolynomialPlot[898].Y = -7_279.765

	pointsOfPolynomialPlot[899].X = 8.99
	pointsOfPolynomialPlot[899].Y = -7_295.474

	pointsOfPolynomialPlot[900].X = 9.0
	pointsOfPolynomialPlot[900].Y = -7_311.187

	pointsOfPolynomialPlot[901].X = 9.01
	pointsOfPolynomialPlot[901].Y = -7_326.904

	pointsOfPolynomialPlot[902].X = 9.02
	pointsOfPolynomialPlot[902].Y = -7_342.625

	pointsOfPolynomialPlot[903].X = 9.03
	pointsOfPolynomialPlot[903].Y = -7_358.35

	pointsOfPolynomialPlot[904].X = 9.04
	pointsOfPolynomialPlot[904].Y = -7_374.078

	pointsOfPolynomialPlot[905].X = 9.05
	pointsOfPolynomialPlot[905].Y = -7_389.811

	pointsOfPolynomialPlot[906].X = 9.06
	pointsOfPolynomialPlot[906].Y = -7_405.546

	pointsOfPolynomialPlot[907].X = 9.07
	pointsOfPolynomialPlot[907].Y = -7_421.286

	pointsOfPolynomialPlot[908].X = 9.08
	pointsOfPolynomialPlot[908].Y = -7_437.028

	pointsOfPolynomialPlot[909].X = 9.09
	pointsOfPolynomialPlot[909].Y = -7_452.774

	pointsOfPolynomialPlot[910].X = 9.1
	pointsOfPolynomialPlot[910].Y = -7_468.523

	pointsOfPolynomialPlot[911].X = 9.11
	pointsOfPolynomialPlot[911].Y = -7_484.276

	pointsOfPolynomialPlot[912].X = 9.12
	pointsOfPolynomialPlot[912].Y = -7_500.031

	pointsOfPolynomialPlot[913].X = 9.13
	pointsOfPolynomialPlot[913].Y = -7_515.789

	pointsOfPolynomialPlot[914].X = 9.14
	pointsOfPolynomialPlot[914].Y = -7_531.55

	pointsOfPolynomialPlot[915].X = 9.15
	pointsOfPolynomialPlot[915].Y = -7_547.313

	pointsOfPolynomialPlot[916].X = 9.16
	pointsOfPolynomialPlot[916].Y = -7_563.079

	pointsOfPolynomialPlot[917].X = 9.17
	pointsOfPolynomialPlot[917].Y = -7_578.848

	pointsOfPolynomialPlot[918].X = 9.18
	pointsOfPolynomialPlot[918].Y = -7_594.619

	pointsOfPolynomialPlot[919].X = 9.19
	pointsOfPolynomialPlot[919].Y = -7_610.392

	pointsOfPolynomialPlot[920].X = 9.2
	pointsOfPolynomialPlot[920].Y = -7_626.167

	pointsOfPolynomialPlot[921].X = 9.21
	pointsOfPolynomialPlot[921].Y = -7_641.945

	pointsOfPolynomialPlot[922].X = 9.22
	pointsOfPolynomialPlot[922].Y = -7_657.724

	pointsOfPolynomialPlot[923].X = 9.23
	pointsOfPolynomialPlot[923].Y = -7_673.506

	pointsOfPolynomialPlot[924].X = 9.24
	pointsOfPolynomialPlot[924].Y = -7_689.289

	pointsOfPolynomialPlot[925].X = 9.25
	pointsOfPolynomialPlot[925].Y = -7_705.074

	pointsOfPolynomialPlot[926].X = 9.26
	pointsOfPolynomialPlot[926].Y = -7_720.86

	pointsOfPolynomialPlot[927].X = 9.27
	pointsOfPolynomialPlot[927].Y = -7_736.648

	pointsOfPolynomialPlot[928].X = 9.28
	pointsOfPolynomialPlot[928].Y = -7_752.437

	pointsOfPolynomialPlot[929].X = 9.29
	pointsOfPolynomialPlot[929].Y = -7_768.228

	pointsOfPolynomialPlot[930].X = 9.3
	pointsOfPolynomialPlot[930].Y = -7_784.019

	pointsOfPolynomialPlot[931].X = 9.31
	pointsOfPolynomialPlot[931].Y = -7_799.812

	pointsOfPolynomialPlot[932].X = 9.32
	pointsOfPolynomialPlot[932].Y = -7_815.606

	pointsOfPolynomialPlot[933].X = 9.33
	pointsOfPolynomialPlot[933].Y = -7_831.4

	pointsOfPolynomialPlot[934].X = 9.34
	pointsOfPolynomialPlot[934].Y = -7_847.196

	pointsOfPolynomialPlot[935].X = 9.35
	pointsOfPolynomialPlot[935].Y = -7_862.992

	pointsOfPolynomialPlot[936].X = 9.36
	pointsOfPolynomialPlot[936].Y = -7_878.788

	pointsOfPolynomialPlot[937].X = 9.37
	pointsOfPolynomialPlot[937].Y = -7_894.585

	pointsOfPolynomialPlot[938].X = 9.38
	pointsOfPolynomialPlot[938].Y = -7_910.382

	pointsOfPolynomialPlot[939].X = 9.39
	pointsOfPolynomialPlot[939].Y = -7_926.18

	pointsOfPolynomialPlot[940].X = 9.4
	pointsOfPolynomialPlot[940].Y = -7_941.977

	pointsOfPolynomialPlot[941].X = 9.41
	pointsOfPolynomialPlot[941].Y = -7_957.775

	pointsOfPolynomialPlot[942].X = 9.42
	pointsOfPolynomialPlot[942].Y = -7_973.573

	pointsOfPolynomialPlot[943].X = 9.43
	pointsOfPolynomialPlot[943].Y = -7_989.37

	pointsOfPolynomialPlot[944].X = 9.44
	pointsOfPolynomialPlot[944].Y = -8_005.167

	pointsOfPolynomialPlot[945].X = 9.45
	pointsOfPolynomialPlot[945].Y = -8_020.964

	pointsOfPolynomialPlot[946].X = 9.46
	pointsOfPolynomialPlot[946].Y = -8_036.76

	pointsOfPolynomialPlot[947].X = 9.47
	pointsOfPolynomialPlot[947].Y = -8_052.555

	pointsOfPolynomialPlot[948].X = 9.48
	pointsOfPolynomialPlot[948].Y = -8_068.35

	pointsOfPolynomialPlot[949].X = 9.49
	pointsOfPolynomialPlot[949].Y = -8_084.144

	pointsOfPolynomialPlot[950].X = 9.5
	pointsOfPolynomialPlot[950].Y = -8_099.937

	pointsOfPolynomialPlot[951].X = 9.51
	pointsOfPolynomialPlot[951].Y = -8_115.729

	pointsOfPolynomialPlot[952].X = 9.52
	pointsOfPolynomialPlot[952].Y = -8_131.52

	pointsOfPolynomialPlot[953].X = 9.53
	pointsOfPolynomialPlot[953].Y = -8_147.309

	pointsOfPolynomialPlot[954].X = 9.54
	pointsOfPolynomialPlot[954].Y = -8_163.097

	pointsOfPolynomialPlot[955].X = 9.55
	pointsOfPolynomialPlot[955].Y = -8_178.884

	pointsOfPolynomialPlot[956].X = 9.56
	pointsOfPolynomialPlot[956].Y = -8_194.669

	pointsOfPolynomialPlot[957].X = 9.57
	pointsOfPolynomialPlot[957].Y = -8_210.453

	pointsOfPolynomialPlot[958].X = 9.58
	pointsOfPolynomialPlot[958].Y = -8_226.234

	pointsOfPolynomialPlot[959].X = 9.59
	pointsOfPolynomialPlot[959].Y = -8_242.014

	pointsOfPolynomialPlot[960].X = 9.6
	pointsOfPolynomialPlot[960].Y = -8_257.791

	pointsOfPolynomialPlot[961].X = 9.61
	pointsOfPolynomialPlot[961].Y = -8_273.34

	pointsOfPolynomialPlot[962].X = 9.62
	pointsOfPolynomialPlot[962].Y = -8_289.34

	pointsOfPolynomialPlot[963].X = 9.63
	pointsOfPolynomialPlot[963].Y = -8_305.111

	pointsOfPolynomialPlot[964].X = 9.64
	pointsOfPolynomialPlot[964].Y = -8_320.879

	pointsOfPolynomialPlot[965].X = 9.65
	pointsOfPolynomialPlot[965].Y = -8_336.645

	pointsOfPolynomialPlot[966].X = 9.66
	pointsOfPolynomialPlot[966].Y = -8_352.408

	pointsOfPolynomialPlot[967].X = 9.67
	pointsOfPolynomialPlot[967].Y = -8_368.169

	pointsOfPolynomialPlot[968].X = 9.68
	pointsOfPolynomialPlot[968].Y = -8_383.926

	pointsOfPolynomialPlot[969].X = 9.69
	pointsOfPolynomialPlot[969].Y = -8_399.68

	pointsOfPolynomialPlot[970].X = 9.7
	pointsOfPolynomialPlot[970].Y = -8_415.431

	pointsOfPolynomialPlot[971].X = 9.71
	pointsOfPolynomialPlot[971].Y = -8_431.179

	pointsOfPolynomialPlot[972].X = 9.72
	pointsOfPolynomialPlot[972].Y = -8_446.924

	pointsOfPolynomialPlot[973].X = 9.73
	pointsOfPolynomialPlot[973].Y = -8_462.665

	pointsOfPolynomialPlot[974].X = 9.74
	pointsOfPolynomialPlot[974].Y = -8_478.403

	pointsOfPolynomialPlot[975].X = 9.75
	pointsOfPolynomialPlot[975].Y = -8_494.136

	pointsOfPolynomialPlot[976].X = 9.76
	pointsOfPolynomialPlot[976].Y = -8_509.866

	pointsOfPolynomialPlot[977].X = 9.77
	pointsOfPolynomialPlot[977].Y = -8_525.592

	pointsOfPolynomialPlot[978].X = 9.78
	pointsOfPolynomialPlot[978].Y = -8_541.314

	pointsOfPolynomialPlot[979].X = 9.79
	pointsOfPolynomialPlot[979].Y = -8_557.032

	pointsOfPolynomialPlot[980].X = 9.8
	pointsOfPolynomialPlot[980].Y = -8_572.745

	pointsOfPolynomialPlot[981].X = 9.81
	pointsOfPolynomialPlot[981].Y = -8_588.455

	pointsOfPolynomialPlot[982].X = 9.82
	pointsOfPolynomialPlot[982].Y = -8_604.159

	pointsOfPolynomialPlot[983].X = 9.83
	pointsOfPolynomialPlot[983].Y = -8_619.859

	pointsOfPolynomialPlot[984].X = 9.84
	pointsOfPolynomialPlot[984].Y = -8_635.554

	pointsOfPolynomialPlot[985].X = 9.85
	pointsOfPolynomialPlot[985].Y = -8_651.245

	pointsOfPolynomialPlot[986].X = 9.86
	pointsOfPolynomialPlot[986].Y = -8_666.93

	pointsOfPolynomialPlot[987].X = 9.87
	pointsOfPolynomialPlot[987].Y = -8_682.61

	pointsOfPolynomialPlot[988].X = 9.88
	pointsOfPolynomialPlot[988].Y = -8_698.285

	pointsOfPolynomialPlot[989].X = 9.89
	pointsOfPolynomialPlot[989].Y = -8_713.955

	pointsOfPolynomialPlot[990].X = 9.9
	pointsOfPolynomialPlot[990].Y = -8_729.619

	pointsOfPolynomialPlot[991].X = 9.91
	pointsOfPolynomialPlot[991].Y = -8_745.278

	pointsOfPolynomialPlot[992].X = 9.92
	pointsOfPolynomialPlot[992].Y = -8_760.931

	pointsOfPolynomialPlot[993].X = 9.93
	pointsOfPolynomialPlot[993].Y = -8_776.579

	pointsOfPolynomialPlot[994].X = 9.94
	pointsOfPolynomialPlot[994].Y = -8_792.22

	pointsOfPolynomialPlot[995].X = 9.95
	pointsOfPolynomialPlot[995].Y = -8_807.855

	pointsOfPolynomialPlot[996].X = 9.96
	pointsOfPolynomialPlot[996].Y = -8_823.484

	pointsOfPolynomialPlot[997].X = 9.97
	pointsOfPolynomialPlot[997].Y = -8_839.107

	pointsOfPolynomialPlot[998].X = 9.98
	pointsOfPolynomialPlot[998].Y = -8_854.724

	pointsOfPolynomialPlot[999].X = 9.99
	pointsOfPolynomialPlot[999].Y = -8_870.334

	pointsOfPolynomialPlot[1_000].X = 10.0
	pointsOfPolynomialPlot[1_000].Y = -8_885.937










	polynomialPlot := plot.New()

	polynomialPlot.Title.Text = "Wykres funkcji f(x) = x^4 - 20x^3 + 33.75x^2 - 235x + 89.0625"

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
		"Polynomial-plot-03.png"); err != nil {

		panic(err)
	}
}
