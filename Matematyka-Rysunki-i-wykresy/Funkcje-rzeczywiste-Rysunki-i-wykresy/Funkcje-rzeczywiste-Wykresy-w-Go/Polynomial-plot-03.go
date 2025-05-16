package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres wielomianu f(x) = x^4 - 20x^3 + 33.75x^2 - 235x + 89.0625

	punktyWykresuWielomianu := make(plotter.XYs, 1_001)

	punktyWykresuWielomianu[0].X = 0.0
	punktyWykresuWielomianu[0].Y = 89.0625

	punktyWykresuWielomianu[1].X = 0.01
	punktyWykresuWielomianu[1].Y = 86.715

	punktyWykresuWielomianu[2].X = 0.02
	punktyWykresuWielomianu[2].Y = 84.375

	punktyWykresuWielomianu[3].X = 0.03
	punktyWykresuWielomianu[3].Y = 82.0423

	punktyWykresuWielomianu[4].X = 0.04
	punktyWykresuWielomianu[4].Y = 79.715

	punktyWykresuWielomianu[5].X = 0.05
	punktyWykresuWielomianu[5].Y = 77.394

	punktyWykresuWielomianu[6].X = 0.06
	punktyWykresuWielomianu[6].Y = 75.079

	punktyWykresuWielomianu[7].X = 0.07
	punktyWykresuWielomianu[7].Y = 72.771

	punktyWykresuWielomianu[8].X = 0.08
	punktyWykresuWielomianu[8].Y = 70.468

	punktyWykresuWielomianu[9].X = 0.09
	punktyWykresuWielomianu[9].Y = 68.171

	punktyWykresuWielomianu[10].X = 0.1
	punktyWykresuWielomianu[10].Y = 65.88

	punktyWykresuWielomianu[11].X = 0.11
	punktyWykresuWielomianu[11].Y = 63.594

	punktyWykresuWielomianu[12].X = 0.12
	punktyWykresuWielomianu[12].Y = 61.314

	punktyWykresuWielomianu[13].X = 0.13
	punktyWykresuWielomianu[13].Y = 59.039

	punktyWykresuWielomianu[14].X = 0.14
	punktyWykresuWielomianu[14].Y = 56.769

	punktyWykresuWielomianu[15].X = 0.15
	punktyWykresuWielomianu[15].Y = 54.504

	punktyWykresuWielomianu[16].X = 0.16
	punktyWykresuWielomianu[16].Y = 52.245

	punktyWykresuWielomianu[17].X = 0.17
	punktyWykresuWielomianu[17].Y = 49.99

	punktyWykresuWielomianu[18].X = 0.18
	punktyWykresuWielomianu[18].Y = 47.74

	punktyWykresuWielomianu[19].X = 0.19
	punktyWykresuWielomianu[19].Y = 45.495

	punktyWykresuWielomianu[20].X = 0.2
	punktyWykresuWielomianu[20].Y = 43.254

	punktyWykresuWielomianu[21].X = 0.21
	punktyWykresuWielomianu[21].Y = 41.017

	punktyWykresuWielomianu[22].X = 0.22
	punktyWykresuWielomianu[22].Y = 38.785

	punktyWykresuWielomianu[23].X = 0.23
	punktyWykresuWielomianu[23].Y = 36.557

	punktyWykresuWielomianu[24].X = 0.24
	punktyWykresuWielomianu[24].Y = 34.333

	punktyWykresuWielomianu[25].X = 0.25
	punktyWykresuWielomianu[25].Y = 32.113

	punktyWykresuWielomianu[26].X = 0.26
	punktyWykresuWielomianu[26].Y = 29.897

	punktyWykresuWielomianu[27].X = 0.27
	punktyWykresuWielomianu[27].Y = 27.684

	punktyWykresuWielomianu[28].X = 0.28
	punktyWykresuWielomianu[28].Y = 25.475

	punktyWykresuWielomianu[29].X = 0.29
	punktyWykresuWielomianu[29].Y = 23.27

	punktyWykresuWielomianu[30].X = 0.3
	punktyWykresuWielomianu[30].Y = 21.068

	punktyWykresuWielomianu[31].X = 0.31
	punktyWykresuWielomianu[31].Y = 18.869

	punktyWykresuWielomianu[32].X = 0.32
	punktyWykresuWielomianu[32].Y = 16.673

	punktyWykresuWielomianu[33].X = 0.33
	punktyWykresuWielomianu[33].Y = 14.481

	punktyWykresuWielomianu[34].X = 0.34
	punktyWykresuWielomianu[34].Y = 12.291

	punktyWykresuWielomianu[35].X = 0.35
	punktyWykresuWielomianu[35].Y = 10.104

	punktyWykresuWielomianu[36].X = 0.36
	punktyWykresuWielomianu[36].Y = 7.92

	punktyWykresuWielomianu[37].X = 0.37
	punktyWykresuWielomianu[37].Y = 5.738

	punktyWykresuWielomianu[38].X = 0.38
	punktyWykresuWielomianu[38].Y = 3.559

	punktyWykresuWielomianu[39].X = 0.39
	punktyWykresuWielomianu[39].Y = 1.382

	punktyWykresuWielomianu[40].X = 0.4
	punktyWykresuWielomianu[40].Y = -0.791

	punktyWykresuWielomianu[41].X = 0.41
	punktyWykresuWielomianu[41].Y = -2.964

	punktyWykresuWielomianu[42].X = 0.42
	punktyWykresuWielomianu[42].Y = -5.134

	punktyWykresuWielomianu[43].X = 0.43
	punktyWykresuWielomianu[43].Y = -7.303

	punktyWykresuWielomianu[44].X = 0.44
	punktyWykresuWielomianu[44].Y = -9.469

	punktyWykresuWielomianu[45].X = 0.45
	punktyWykresuWielomianu[45].Y = -11.634

	punktyWykresuWielomianu[46].X = 0.46
	punktyWykresuWielomianu[46].Y = -13.797

	punktyWykresuWielomianu[47].X = 0.47
	punktyWykresuWielomianu[47].Y = -15.959

	punktyWykresuWielomianu[48].X = 0.48
	punktyWykresuWielomianu[48].Y = -18.12

	punktyWykresuWielomianu[49].X = 0.49
	punktyWykresuWielomianu[49].Y = -20.279

	punktyWykresuWielomianu[50].X = 0.5
	punktyWykresuWielomianu[50].Y = -22.437

	punktyWykresuWielomianu[51].X = 0.51
	punktyWykresuWielomianu[51].Y = -24.594

	punktyWykresuWielomianu[52].X = 0.52
	punktyWykresuWielomianu[52].Y = -26.75

	punktyWykresuWielomianu[53].X = 0.53
	punktyWykresuWielomianu[53].Y = -28.905

	punktyWykresuWielomianu[54].X = 0.54
	punktyWykresuWielomianu[54].Y = -31.06

	punktyWykresuWielomianu[55].X = 0.55
	punktyWykresuWielomianu[55].Y = -33.214

	punktyWykresuWielomianu[56].X = 0.56
	punktyWykresuWielomianu[56].Y = -35.367

	punktyWykresuWielomianu[57].X = 0.57
	punktyWykresuWielomianu[57].Y = -37.52

	punktyWykresuWielomianu[58].X = 0.58
	punktyWykresuWielomianu[58].Y = -39.673

	punktyWykresuWielomianu[59].X = 0.59
	punktyWykresuWielomianu[59].Y = -41.825

	punktyWykresuWielomianu[60].X = 0.6
	punktyWykresuWielomianu[60].Y = -43.977

	punktyWykresuWielomianu[61].X = 0.61
	punktyWykresuWielomianu[61].Y = -46.13

	punktyWykresuWielomianu[62].X = 0.62
	punktyWykresuWielomianu[62].Y = -48.282

	punktyWykresuWielomianu[63].X = 0.63
	punktyWykresuWielomianu[63].Y = -50.435

	punktyWykresuWielomianu[64].X = 0.64
	punktyWykresuWielomianu[64].Y = -52.588

	punktyWykresuWielomianu[65].X = 0.65
	punktyWykresuWielomianu[65].Y = -54.742

	punktyWykresuWielomianu[66].X = 0.66
	punktyWykresuWielomianu[66].Y = -56.896

	punktyWykresuWielomianu[67].X = 0.67
	punktyWykresuWielomianu[67].Y = -59.05

	punktyWykresuWielomianu[68].X = 0.68
	punktyWykresuWielomianu[68].Y = -61.206

	punktyWykresuWielomianu[69].X = 0.69
	punktyWykresuWielomianu[69].Y = -63.362

	punktyWykresuWielomianu[70].X = 0.7
	punktyWykresuWielomianu[70].Y = -65.519

	punktyWykresuWielomianu[71].X = 0.71
	punktyWykresuWielomianu[71].Y = -67.678

	punktyWykresuWielomianu[72].X = 0.72
	punktyWykresuWielomianu[72].Y = -69.837

	punktyWykresuWielomianu[73].X = 0.73
	punktyWykresuWielomianu[73].Y = -71.998

	punktyWykresuWielomianu[74].X = 0.74
	punktyWykresuWielomianu[74].Y = -74.16

	punktyWykresuWielomianu[75].X = 0.75
	punktyWykresuWielomianu[75].Y = -76.324

	punktyWykresuWielomianu[76].X = 0.76
	punktyWykresuWielomianu[76].Y = -78.489

	punktyWykresuWielomianu[77].X = 0.77
	punktyWykresuWielomianu[77].Y = -80.656

	punktyWykresuWielomianu[78].X = 0.78
	punktyWykresuWielomianu[78].Y = -82.824

	punktyWykresuWielomianu[79].X = 0.79
	punktyWykresuWielomianu[79].Y = -84.995

	punktyWykresuWielomianu[80].X = 0.8
	punktyWykresuWielomianu[80].Y = -87.167

	punktyWykresuWielomianu[81].X = 0.81
	punktyWykresuWielomianu[81].Y = -89.342

	punktyWykresuWielomianu[82].X = 0.82
	punktyWykresuWielomianu[82].Y = -91.519

	punktyWykresuWielomianu[83].X = 0.83
	punktyWykresuWielomianu[83].Y = -93.698

	punktyWykresuWielomianu[84].X = 0.84
	punktyWykresuWielomianu[84].Y = -95.879

	punktyWykresuWielomianu[85].X = 0.85
	punktyWykresuWielomianu[85].Y = -98.063

	punktyWykresuWielomianu[86].X = 0.86
	punktyWykresuWielomianu[86].Y = -100.25

	punktyWykresuWielomianu[87].X = 0.87
	punktyWykresuWielomianu[87].Y = -102.439

	punktyWykresuWielomianu[88].X = 0.88
	punktyWykresuWielomianu[88].Y = -104.631

	punktyWykresuWielomianu[89].X = 0.89
	punktyWykresuWielomianu[89].Y = -106.826

	punktyWykresuWielomianu[90].X = 0.9
	punktyWykresuWielomianu[90].Y = -109.023

	punktyWykresuWielomianu[91].X = 0.91
	punktyWykresuWielomianu[91].Y = -111.224

	punktyWykresuWielomianu[92].X = 0.92
	punktyWykresuWielomianu[92].Y = -113.428

	punktyWykresuWielomianu[93].X = 0.93
	punktyWykresuWielomianu[93].Y = -115.636

	punktyWykresuWielomianu[94].X = 0.94
	punktyWykresuWielomianu[94].Y = -117.846

	punktyWykresuWielomianu[95].X = 0.95
	punktyWykresuWielomianu[95].Y = -120.061

	punktyWykresuWielomianu[96].X = 0.96
	punktyWykresuWielomianu[96].Y = -122.278

	punktyWykresuWielomianu[97].X = 0.97
	punktyWykresuWielomianu[97].Y = -124.5

	punktyWykresuWielomianu[98].X = 0.98
	punktyWykresuWielomianu[98].Y = -126.725

	punktyWykresuWielomianu[99].X = 0.99
	punktyWykresuWielomianu[99].Y = -128.954

	punktyWykresuWielomianu[100].X = 1.0
	punktyWykresuWielomianu[100].Y = -131.187

	punktyWykresuWielomianu[101].X = 1.01
	punktyWykresuWielomianu[101].Y = -133.424

	punktyWykresuWielomianu[102].X = 1.02
	punktyWykresuWielomianu[102].Y = -135.665

	punktyWykresuWielomianu[103].X = 1.03
	punktyWykresuWielomianu[103].Y = -137.911

	punktyWykresuWielomianu[104].X = 1.04
	punktyWykresuWielomianu[104].Y = -140.16

	punktyWykresuWielomianu[105].X = 1.05
	punktyWykresuWielomianu[105].Y = -142.415

	punktyWykresuWielomianu[106].X = 1.06
	punktyWykresuWielomianu[106].Y = -144.673

	punktyWykresuWielomianu[107].X = 1.07
	punktyWykresuWielomianu[107].Y = -146.937

	punktyWykresuWielomianu[108].X = 1.08
	punktyWykresuWielomianu[108].Y = -149.205

	punktyWykresuWielomianu[109].X = 1.09
	punktyWykresuWielomianu[109].Y = -151.478

	punktyWykresuWielomianu[110].X = 1.1
	punktyWykresuWielomianu[110].Y = -153.755

	punktyWykresuWielomianu[111].X = 1.11
	punktyWykresuWielomianu[111].Y = -156.038

	punktyWykresuWielomianu[112].X = 1.12
	punktyWykresuWielomianu[112].Y = -158.326

	punktyWykresuWielomianu[113].X = 1.13
	punktyWykresuWielomianu[113].Y = -160.619

	punktyWykresuWielomianu[114].X = 1.14
	punktyWykresuWielomianu[114].Y = -162.917

	punktyWykresuWielomianu[115].X = 1.15
	punktyWykresuWielomianu[115].Y = -165.221

	punktyWykresuWielomianu[116].X = 1.16
	punktyWykresuWielomianu[116].Y = -167.53

	punktyWykresuWielomianu[117].X = 1.17
	punktyWykresuWielomianu[117].Y = -169.845

	punktyWykresuWielomianu[118].X = 1.18
	punktyWykresuWielomianu[118].Y = -172.165

	punktyWykresuWielomianu[119].X = 1.19
	punktyWykresuWielomianu[119].Y = -174.492

	punktyWykresuWielomianu[120].X = 1.2
	punktyWykresuWielomianu[120].Y = -176.823

	punktyWykresuWielomianu[121].X = 1.21
	punktyWykresuWielomianu[121].Y = -179.161

	punktyWykresuWielomianu[122].X = 1.22
	punktyWykresuWielomianu[122].Y = -181.505

	punktyWykresuWielomianu[123].X = 1.23
	punktyWykresuWielomianu[123].Y = -183.855

	punktyWykresuWielomianu[124].X = 1.24
	punktyWykresuWielomianu[124].Y = -186.211

	punktyWykresuWielomianu[125].X = 1.25
	punktyWykresuWielomianu[125].Y = -188.574

	punktyWykresuWielomianu[126].X = 1.26
	punktyWykresuWielomianu[126].Y = -190.943

	punktyWykresuWielomianu[127].X = 1.27
	punktyWykresuWielomianu[127].Y = -193.318

	punktyWykresuWielomianu[128].X = 1.28
	punktyWykresuWielomianu[128].Y = -195.7

	punktyWykresuWielomianu[129].X = 1.29
	punktyWykresuWielomianu[129].Y = -198.088

	punktyWykresuWielomianu[130].X = 1.3
	punktyWykresuWielomianu[130].Y = -200.483

	punktyWykresuWielomianu[131].X = 1.31
	punktyWykresuWielomianu[131].Y = -202.885

	punktyWykresuWielomianu[132].X = 1.32
	punktyWykresuWielomianu[132].Y = -205.294

	punktyWykresuWielomianu[133].X = 1.33
	punktyWykresuWielomianu[133].Y = -207.71

	punktyWykresuWielomianu[134].X = 1.34
	punktyWykresuWielomianu[134].Y = -210.133

	punktyWykresuWielomianu[135].X = 1.35
	punktyWykresuWielomianu[135].Y = -212.564

	punktyWykresuWielomianu[136].X = 1.36
	punktyWykresuWielomianu[136].Y = -215.001

	punktyWykresuWielomianu[137].X = 1.37
	punktyWykresuWielomianu[137].Y = -217.446

	punktyWykresuWielomianu[138].X = 1.38
	punktyWykresuWielomianu[138].Y = -219.898

	punktyWykresuWielomianu[139].X = 1.39
	punktyWykresuWielomianu[139].Y = -222.358

	punktyWykresuWielomianu[140].X = 1.4
	punktyWykresuWielomianu[140].Y = -224.825

	punktyWykresuWielomianu[141].X = 1.41
	punktyWykresuWielomianu[141].Y = -227.301

	punktyWykresuWielomianu[142].X = 1.42
	punktyWykresuWielomianu[142].Y = -229.783

	punktyWykresuWielomianu[143].X = 1.43
	punktyWykresuWielomianu[143].Y = -232.274

	punktyWykresuWielomianu[144].X = 1.44
	punktyWykresuWielomianu[144].Y = -234.773

	punktyWykresuWielomianu[145].X = 1.45
	punktyWykresuWielomianu[145].Y = -237.28

	punktyWykresuWielomianu[146].X = 1.46
	punktyWykresuWielomianu[146].Y = -239.795

	punktyWykresuWielomianu[147].X = 1.47
	punktyWykresuWielomianu[147].Y = -242.318

	punktyWykresuWielomianu[148].X = 1.48
	punktyWykresuWielomianu[148].Y = -244.849

	punktyWykresuWielomianu[149].X = 1.49
	punktyWykresuWielomianu[149].Y = -247.389

	punktyWykresuWielomianu[150].X = 1.5
	punktyWykresuWielomianu[150].Y = -249.937

	punktyWykresuWielomianu[151].X = 1.51
	punktyWykresuWielomianu[151].Y = -252.494

	punktyWykresuWielomianu[152].X = 1.52
	punktyWykresuWielomianu[152].Y = -255.059

	punktyWykresuWielomianu[153].X = 1.53
	punktyWykresuWielomianu[153].Y = -257.633

	punktyWykresuWielomianu[154].X = 1.54
	punktyWykresuWielomianu[154].Y = -260.216

	punktyWykresuWielomianu[155].X = 1.55
	punktyWykresuWielomianu[155].Y = -262.808

	punktyWykresuWielomianu[156].X = 1.56
	punktyWykresuWielomianu[156].Y = -265.409

	punktyWykresuWielomianu[157].X = 1.57
	punktyWykresuWielomianu[157].Y = -268.019

	punktyWykresuWielomianu[158].X = 1.58
	punktyWykresuWielomianu[158].Y = -270.638

	punktyWykresuWielomianu[159].X = 1.59
	punktyWykresuWielomianu[159].Y = -273.903

	punktyWykresuWielomianu[160].X = 1.6
	punktyWykresuWielomianu[160].Y = -275.903

	punktyWykresuWielomianu[161].X = 1.61
	punktyWykresuWielomianu[161].Y = -278.55

	punktyWykresuWielomianu[162].X = 1.62
	punktyWykresuWielomianu[162].Y = -281.207

	punktyWykresuWielomianu[163].X = 1.63
	punktyWykresuWielomianu[163].Y = -283.872

	punktyWykresuWielomianu[164].X = 1.64
	punktyWykresuWielomianu[164].Y = -286.548

	punktyWykresuWielomianu[165].X = 1.65
	punktyWykresuWielomianu[165].Y = -289.233

	punktyWykresuWielomianu[166].X = 1.66
	punktyWykresuWielomianu[166].Y = -291.928

	punktyWykresuWielomianu[167].X = 1.67
	punktyWykresuWielomianu[167].Y = -294.633

	punktyWykresuWielomianu[168].X = 1.68
	punktyWykresuWielomianu[168].Y = -297.348

	punktyWykresuWielomianu[169].X = 1.69
	punktyWykresuWielomianu[169].Y = -300.073

	punktyWykresuWielomianu[170].X = 1.7
	punktyWykresuWielomianu[170].Y = -302.807

	punktyWykresuWielomianu[171].X = 1.71
	punktyWykresuWielomianu[171].Y = -305.553

	punktyWykresuWielomianu[172].X = 1.72
	punktyWykresuWielomianu[172].Y = -308.308

	punktyWykresuWielomianu[173].X = 1.73
	punktyWykresuWielomianu[173].Y = -311.074

	punktyWykresuWielomianu[174].X = 1.74
	punktyWykresuWielomianu[174].Y = -313.85

	punktyWykresuWielomianu[175].X = 1.75
	punktyWykresuWielomianu[175].Y = -316.636

	punktyWykresuWielomianu[176].X = 1.76
	punktyWykresuWielomianu[176].Y = -319.433

	punktyWykresuWielomianu[177].X = 1.77
	punktyWykresuWielomianu[177].Y = -322.241

	punktyWykresuWielomianu[178].X = 1.78
	punktyWykresuWielomianu[178].Y = -325.06

	punktyWykresuWielomianu[179].X = 1.79
	punktyWykresuWielomianu[179].Y = -327.889

	punktyWykresuWielomianu[180].X = 1.8
	punktyWykresuWielomianu[180].Y = -330.729

	punktyWykresuWielomianu[181].X = 1.81
	punktyWykresuWielomianu[181].Y = -333.581

	punktyWykresuWielomianu[182].X = 1.82
	punktyWykresuWielomianu[182].Y = -336.443

	punktyWykresuWielomianu[183].X = 1.83
	punktyWykresuWielomianu[183].Y = -339.316

	punktyWykresuWielomianu[184].X = 1.84
	punktyWykresuWielomianu[184].Y = -342.201

	punktyWykresuWielomianu[185].X = 1.85
	punktyWykresuWielomianu[185].Y = -345.097

	punktyWykresuWielomianu[186].X = 1.86
	punktyWykresuWielomianu[186].Y = -348.004

	punktyWykresuWielomianu[187].X = 1.87
	punktyWykresuWielomianu[187].Y = -350.922

	punktyWykresuWielomianu[188].X = 1.88
	punktyWykresuWielomianu[188].Y = -353.853

	punktyWykresuWielomianu[189].X = 1.89
	punktyWykresuWielomianu[189].Y = -356.794

	punktyWykresuWielomianu[190].X = 1.9
	punktyWykresuWielomianu[190].Y = -359.747

	punktyWykresuWielomianu[191].X = 1.91
	punktyWykresuWielomianu[191].Y = -362.712

	punktyWykresuWielomianu[192].X = 1.92
	punktyWykresuWielomianu[192].Y = -365.689

	punktyWykresuWielomianu[193].X = 1.93
	punktyWykresuWielomianu[193].Y = -368.678

	punktyWykresuWielomianu[194].X = 1.94
	punktyWykresuWielomianu[194].Y = -371.679

	punktyWykresuWielomianu[195].X = 1.95
	punktyWykresuWielomianu[195].Y = -374.691

	punktyWykresuWielomianu[196].X = 1.96
	punktyWykresuWielomianu[196].Y = -377.716

	punktyWykresuWielomianu[197].X = 1.97
	punktyWykresuWielomianu[197].Y = -380.753

	punktyWykresuWielomianu[198].X = 1.98
	punktyWykresuWielomianu[198].Y = -383.802

	punktyWykresuWielomianu[199].X = 1.99
	punktyWykresuWielomianu[199].Y = -386.863

	punktyWykresuWielomianu[200].X = 2.0
	punktyWykresuWielomianu[200].Y = -389.937

	punktyWykresuWielomianu[201].X = 2.01
	punktyWykresuWielomianu[201].Y = -393.023

	punktyWykresuWielomianu[202].X = 2.02
	punktyWykresuWielomianu[202].Y = -396.122

	punktyWykresuWielomianu[203].X = 2.03
	punktyWykresuWielomianu[203].Y = -399.233

	punktyWykresuWielomianu[204].X = 2.04
	punktyWykresuWielomianu[204].Y = -402.357

	punktyWykresuWielomianu[205].X = 2.05
	punktyWykresuWielomianu[205].Y = -405.494

	punktyWykresuWielomianu[206].X = 2.06
	punktyWykresuWielomianu[206].Y = -408.644

	punktyWykresuWielomianu[207].X = 2.07
	punktyWykresuWielomianu[207].Y = -411.806

	punktyWykresuWielomianu[208].X = 2.08
	punktyWykresuWielomianu[208].Y = -414.982

	punktyWykresuWielomianu[209].X = 2.09
	punktyWykresuWielomianu[209].Y = -418.17

	punktyWykresuWielomianu[210].X = 2.10
	punktyWykresuWielomianu[210].Y = -421.371

	punktyWykresuWielomianu[211].X = 2.11
	punktyWykresuWielomianu[211].Y = -424.586

	punktyWykresuWielomianu[212].X = 2.12
	punktyWykresuWielomianu[212].Y = -427.814

	punktyWykresuWielomianu[213].X = 2.13
	punktyWykresuWielomianu[213].Y = -431.055

	punktyWykresuWielomianu[214].X = 2.14
	punktyWykresuWielomianu[214].Y = -434.31

	punktyWykresuWielomianu[215].X = 2.15
	punktyWykresuWielomianu[215].Y = -437.578

	punktyWykresuWielomianu[216].X = 2.16
	punktyWykresuWielomianu[216].Y = -440.859

	punktyWykresuWielomianu[217].X = 2.17
	punktyWykresuWielomianu[217].Y = -444.154

	punktyWykresuWielomianu[218].X = 2.18
	punktyWykresuWielomianu[218].Y = -447.463

	punktyWykresuWielomianu[219].X = 2.19
	punktyWykresuWielomianu[219].Y = -450.785

	punktyWykresuWielomianu[220].X = 2.2
	punktyWykresuWielomianu[220].Y = -454.121

	punktyWykresuWielomianu[221].X = 2.21
	punktyWykresuWielomianu[221].Y = -457.471

	punktyWykresuWielomianu[222].X = 2.22
	punktyWykresuWielomianu[222].Y = -460.835

	punktyWykresuWielomianu[223].X = 2.23
	punktyWykresuWielomianu[223].Y = -464.213

	punktyWykresuWielomianu[224].X = 2.24
	punktyWykresuWielomianu[224].Y = -467.605

	punktyWykresuWielomianu[225].X = 2.25
	punktyWykresuWielomianu[225].Y = -471.011

	punktyWykresuWielomianu[226].X = 2.26
	punktyWykresuWielomianu[226].Y = -474.431

	punktyWykresuWielomianu[227].X = 2.27
	punktyWykresuWielomianu[227].Y = -477.866

	punktyWykresuWielomianu[228].X = 2.28
	punktyWykresuWielomianu[228].Y = -481.315

	punktyWykresuWielomianu[229].X = 2.29
	punktyWykresuWielomianu[229].Y = -484.778

	punktyWykresuWielomianu[230].X = 2.3
	punktyWykresuWielomianu[230].Y = -488.255

	punktyWykresuWielomianu[231].X = 2.31
	punktyWykresuWielomianu[231].Y = -491.748

	punktyWykresuWielomianu[232].X = 2.32
	punktyWykresuWielomianu[232].Y = -495.254

	punktyWykresuWielomianu[233].X = 2.33
	punktyWykresuWielomianu[233].Y = -498.775

	punktyWykresuWielomianu[234].X = 2.34
	punktyWykresuWielomianu[234].Y = -502.311

	punktyWykresuWielomianu[235].X = 2.35
	punktyWykresuWielomianu[235].Y = -505.862

	punktyWykresuWielomianu[236].X = 2.36
	punktyWykresuWielomianu[236].Y = -509.428

	punktyWykresuWielomianu[237].X = 2.37
	punktyWykresuWielomianu[237].Y = -513.008

	punktyWykresuWielomianu[238].X = 2.38
	punktyWykresuWielomianu[238].Y = -516.604

	punktyWykresuWielomianu[239].X = 2.39
	punktyWykresuWielomianu[239].Y = -520.214

	punktyWykresuWielomianu[240].X = 2.4
	punktyWykresuWielomianu[240].Y = -523.839

	punktyWykresuWielomianu[241].X = 2.41
	punktyWykresuWielomianu[241].Y = -527.48

	punktyWykresuWielomianu[242].X = 2.42
	punktyWykresuWielomianu[242].Y = -531.136

	punktyWykresuWielomianu[243].X = 2.43
	punktyWykresuWielomianu[243].Y = -534.807

	punktyWykresuWielomianu[244].X = 2.44
	punktyWykresuWielomianu[244].Y = -538.493

	punktyWykresuWielomianu[245].X = 2.45
	punktyWykresuWielomianu[245].Y = -542.195

	punktyWykresuWielomianu[246].X = 2.46
	punktyWykresuWielomianu[246].Y = -545.912

	punktyWykresuWielomianu[247].X = 2.47
	punktyWykresuWielomianu[247].Y = -549.645

	punktyWykresuWielomianu[248].X = 2.48
	punktyWykresuWielomianu[248].Y = -553.393

	punktyWykresuWielomianu[249].X = 2.49
	punktyWykresuWielomianu[249].Y = -557.157

	punktyWykresuWielomianu[250].X = 2.5
	punktyWykresuWielomianu[250].Y = -560.937

	punktyWykresuWielomianu[251].X = 2.51
	punktyWykresuWielomianu[251].Y = -564.732

	punktyWykresuWielomianu[252].X = 2.52
	punktyWykresuWielomianu[252].Y = -568.544

	punktyWykresuWielomianu[253].X = 2.53
	punktyWykresuWielomianu[253].Y = -572.371

	punktyWykresuWielomianu[254].X = 2.54
	punktyWykresuWielomianu[254].Y = -576.214

	punktyWykresuWielomianu[255].X = 2.55
	punktyWykresuWielomianu[255].Y = -580.073

	punktyWykresuWielomianu[256].X = 2.56
	punktyWykresuWielomianu[256].Y = -583.948

	punktyWykresuWielomianu[257].X = 2.57
	punktyWykresuWielomianu[257].Y = -587.839

	punktyWykresuWielomianu[258].X = 2.58
	punktyWykresuWielomianu[258].Y = -591.746

	punktyWykresuWielomianu[259].X = 2.59
	punktyWykresuWielomianu[259].Y = -595.67

	punktyWykresuWielomianu[260].X = 2.6
	punktyWykresuWielomianu[260].Y = -599.609

	punktyWykresuWielomianu[261].X = 2.61
	punktyWykresuWielomianu[261].Y = -603.566

	punktyWykresuWielomianu[262].X = 2.62
	punktyWykresuWielomianu[262].Y = -607.538

	punktyWykresuWielomianu[263].X = 2.63
	punktyWykresuWielomianu[263].Y = -611.527

	punktyWykresuWielomianu[264].X = 2.64
	punktyWykresuWielomianu[264].Y = -615.533

	punktyWykresuWielomianu[265].X = 2.65
	punktyWykresuWielomianu[265].Y = -619.555

	punktyWykresuWielomianu[266].X = 2.66
	punktyWykresuWielomianu[266].Y = -623.593

	punktyWykresuWielomianu[267].X = 2.67
	punktyWykresuWielomianu[267].Y = -627.649

	punktyWykresuWielomianu[268].X = 2.68
	punktyWykresuWielomianu[268].Y = -631.721

	punktyWykresuWielomianu[269].X = 2.69
	punktyWykresuWielomianu[269].Y = -635.81

	punktyWykresuWielomianu[270].X = 2.7
	punktyWykresuWielomianu[270].Y = -639.915

	punktyWykresuWielomianu[271].X = 2.71
	punktyWykresuWielomianu[271].Y = -644.038

	punktyWykresuWielomianu[272].X = 2.72
	punktyWykresuWielomianu[272].Y = -648.178

	punktyWykresuWielomianu[273].X = 2.73
	punktyWykresuWielomianu[273].Y = -652.334

	punktyWykresuWielomianu[274].X = 2.74
	punktyWykresuWielomianu[274].Y = -656.508

	punktyWykresuWielomianu[275].X = 2.75
	punktyWykresuWielomianu[275].Y = -660.699

	punktyWykresuWielomianu[276].X = 2.76
	punktyWykresuWielomianu[276].Y = -664.907

	punktyWykresuWielomianu[277].X = 2.77
	punktyWykresuWielomianu[277].Y = -669.132

	punktyWykresuWielomianu[278].X = 2.78
	punktyWykresuWielomianu[278].Y = -673.374

	punktyWykresuWielomianu[279].X = 2.79
	punktyWykresuWielomianu[279].Y = -677.634

	punktyWykresuWielomianu[280].X = 2.8
	punktyWykresuWielomianu[280].Y = -681.911

	punktyWykresuWielomianu[281].X = 2.81
	punktyWykresuWielomianu[281].Y = -686.206

	punktyWykresuWielomianu[282].X = 2.82
	punktyWykresuWielomianu[282].Y = -690.518

	punktyWykresuWielomianu[283].X = 2.83
	punktyWykresuWielomianu[283].Y = -694.848

	punktyWykresuWielomianu[284].X = 2.84
	punktyWykresuWielomianu[284].Y = -699.195

	punktyWykresuWielomianu[285].X = 2.85
	punktyWykresuWielomianu[285].Y = -703.56

	punktyWykresuWielomianu[286].X = 2.86
	punktyWykresuWielomianu[286].Y = -707.943

	punktyWykresuWielomianu[287].X = 2.87
	punktyWykresuWielomianu[287].Y = -712.343

	punktyWykresuWielomianu[288].X = 2.88
	punktyWykresuWielomianu[288].Y = -716.761

	punktyWykresuWielomianu[289].X = 2.89
	punktyWykresuWielomianu[289].Y = -721.197

	punktyWykresuWielomianu[290].X = 2.9
	punktyWykresuWielomianu[290].Y = -725.651

	punktyWykresuWielomianu[291].X = 2.91
	punktyWykresuWielomianu[291].Y = -730.123

	punktyWykresuWielomianu[292].X = 2.92
	punktyWykresuWielomianu[292].Y = -734.613

	punktyWykresuWielomianu[293].X = 2.93
	punktyWykresuWielomianu[293].Y = -739.121

	punktyWykresuWielomianu[294].X = 2.94
	punktyWykresuWielomianu[294].Y = -743.647

	punktyWykresuWielomianu[295].X = 2.95
	punktyWykresuWielomianu[295].Y = -748.192

	punktyWykresuWielomianu[296].X = 2.96
	punktyWykresuWielomianu[296].Y = -752.754

	punktyWykresuWielomianu[297].X = 2.97
	punktyWykresuWielomianu[297].Y = -757.335

	punktyWykresuWielomianu[298].X = 2.98
	punktyWykresuWielomianu[298].Y = -761.934

	punktyWykresuWielomianu[299].X = 2.99
	punktyWykresuWielomianu[299].Y = -766.551

	punktyWykresuWielomianu[300].X = 3.0
	punktyWykresuWielomianu[300].Y = -771.187

	punktyWykresuWielomianu[301].X = 3.01
	punktyWykresuWielomianu[301].Y = -775.841

	punktyWykresuWielomianu[302].X = 3.02
	punktyWykresuWielomianu[302].Y = -780.514

	punktyWykresuWielomianu[303].X = 3.03
	punktyWykresuWielomianu[303].Y = -785.205

	punktyWykresuWielomianu[304].X = 3.04
	punktyWykresuWielomianu[304].Y = -789.915

	punktyWykresuWielomianu[305].X = 3.05
	punktyWykresuWielomianu[305].Y = -794.644

	punktyWykresuWielomianu[306].X = 3.06
	punktyWykresuWielomianu[306].Y = -799.391

	punktyWykresuWielomianu[307].X = 3.07
	punktyWykresuWielomianu[307].Y = -804.157

	punktyWykresuWielomianu[308].X = 3.08
	punktyWykresuWielomianu[308].Y = -808.942

	punktyWykresuWielomianu[309].X = 3.09
	punktyWykresuWielomianu[309].Y = -813.745

	punktyWykresuWielomianu[310].X = 3.1
	punktyWykresuWielomianu[310].Y = -818.745

	punktyWykresuWielomianu[311].X = 3.11
	punktyWykresuWielomianu[311].Y = -823.409

	punktyWykresuWielomianu[312].X = 3.12
	punktyWykresuWielomianu[312].Y = -828.269

	punktyWykresuWielomianu[313].X = 3.13
	punktyWykresuWielomianu[313].Y = -833.148

	punktyWykresuWielomianu[314].X = 3.14
	punktyWykresuWielomianu[314].Y = -838.047

	punktyWykresuWielomianu[315].X = 3.15
	punktyWykresuWielomianu[315].Y = -842.964

	punktyWykresuWielomianu[316].X = 3.16
	punktyWykresuWielomianu[316].Y = -847.901

	punktyWykresuWielomianu[317].X = 3.17
	punktyWykresuWielomianu[317].Y = -852.857

	punktyWykresuWielomianu[318].X = 3.18
	punktyWykresuWielomianu[318].Y = -857.832

	punktyWykresuWielomianu[319].X = 3.19
	punktyWykresuWielomianu[319].Y = -862.826

	punktyWykresuWielomianu[320].X = 3.2
	punktyWykresuWielomianu[320].Y = -867.839

	punktyWykresuWielomianu[321].X = 3.21
	punktyWykresuWielomianu[321].Y = -872.872

	punktyWykresuWielomianu[322].X = 3.22
	punktyWykresuWielomianu[322].Y = -877.925

	punktyWykresuWielomianu[323].X = 3.23
	punktyWykresuWielomianu[323].Y = -882.997

	punktyWykresuWielomianu[324].X = 3.24
	punktyWykresuWielomianu[324].Y = -888.088

	punktyWykresuWielomianu[325].X = 3.25
	punktyWykresuWielomianu[325].Y = -893.329

	punktyWykresuWielomianu[326].X = 3.26
	punktyWykresuWielomianu[326].Y = -898.329

	punktyWykresuWielomianu[327].X = 3.27
	punktyWykresuWielomianu[327].Y = -903.479

	punktyWykresuWielomianu[328].X = 3.28
	punktyWykresuWielomianu[328].Y = -908.649

	punktyWykresuWielomianu[329].X = 3.29
	punktyWykresuWielomianu[329].Y = -913.838

	punktyWykresuWielomianu[330].X = 3.3
	punktyWykresuWielomianu[330].Y = -919.047

	punktyWykresuWielomianu[331].X = 3.31
	punktyWykresuWielomianu[331].Y = -924.276

	punktyWykresuWielomianu[332].X = 3.32
	punktyWykresuWielomianu[332].Y = -929.525

	punktyWykresuWielomianu[333].X = 3.33
	punktyWykresuWielomianu[333].Y = -934.794

	punktyWykresuWielomianu[334].X = 3.34
	punktyWykresuWielomianu[334].Y = -940.082

	punktyWykresuWielomianu[335].X = 3.35
	punktyWykresuWielomianu[335].Y = -945.391

	punktyWykresuWielomianu[336].X = 3.36
	punktyWykresuWielomianu[336].Y = -950.719

	punktyWykresuWielomianu[337].X = 3.37
	punktyWykresuWielomianu[337].Y = -956.068

	punktyWykresuWielomianu[338].X = 3.38
	punktyWykresuWielomianu[338].Y = -961.436

	punktyWykresuWielomianu[339].X = 3.39
	punktyWykresuWielomianu[339].Y = -966.825

	punktyWykresuWielomianu[340].X = 3.4
	punktyWykresuWielomianu[340].Y = -972.233

	punktyWykresuWielomianu[341].X = 3.41
	punktyWykresuWielomianu[341].Y = -977.662

	punktyWykresuWielomianu[342].X = 3.42
	punktyWykresuWielomianu[342].Y = -983.112

	punktyWykresuWielomianu[343].X = 3.43
	punktyWykresuWielomianu[343].Y = -988.581

	punktyWykresuWielomianu[344].X = 3.44
	punktyWykresuWielomianu[344].Y = -994.071

	punktyWykresuWielomianu[345].X = 3.45
	punktyWykresuWielomianu[345].Y = -999.581

	punktyWykresuWielomianu[346].X = 3.46
	punktyWykresuWielomianu[346].Y = -1_005.111

	punktyWykresuWielomianu[347].X = 3.47
	punktyWykresuWielomianu[347].Y = -1_010.662

	punktyWykresuWielomianu[348].X = 3.48
	punktyWykresuWielomianu[348].Y = -1_016.233

	punktyWykresuWielomianu[349].X = 3.49
	punktyWykresuWielomianu[349].Y = -1_021.825

	punktyWykresuWielomianu[350].X = 3.5
	punktyWykresuWielomianu[350].Y = -1_027.437

	punktyWykresuWielomianu[351].X = 3.51
	punktyWykresuWielomianu[351].Y = -1_033.07

	punktyWykresuWielomianu[352].X = 3.52
	punktyWykresuWielomianu[352].Y = -1_038.723

	punktyWykresuWielomianu[353].X = 3.53
	punktyWykresuWielomianu[353].Y = -1_044.397

	punktyWykresuWielomianu[354].X = 3.54
	punktyWykresuWielomianu[354].Y = -1_050.092

	punktyWykresuWielomianu[355].X = 3.55
	punktyWykresuWielomianu[355].Y = -1_055.807

	punktyWykresuWielomianu[356].X = 3.56
	punktyWykresuWielomianu[356].Y = -1_061.543

	punktyWykresuWielomianu[357].X = 3.57
	punktyWykresuWielomianu[357].Y = -1_067.3

	punktyWykresuWielomianu[358].X = 3.58
	punktyWykresuWielomianu[358].Y = -1_073.078

	punktyWykresuWielomianu[359].X = 3.59
	punktyWykresuWielomianu[359].Y = -1_078.876

	punktyWykresuWielomianu[360].X = 3.6
	punktyWykresuWielomianu[360].Y = -1_084.695

	punktyWykresuWielomianu[361].X = 3.61
	punktyWykresuWielomianu[361].Y = -1_090.536

	punktyWykresuWielomianu[362].X = 3.62
	punktyWykresuWielomianu[362].Y = -1_096.397

	punktyWykresuWielomianu[363].X = 3.63
	punktyWykresuWielomianu[363].Y = -1_102.279

	punktyWykresuWielomianu[364].X = 3.64
	punktyWykresuWielomianu[364].Y = -1_108.182

	punktyWykresuWielomianu[365].X = 3.65
	punktyWykresuWielomianu[365].Y = -1_114.106

	punktyWykresuWielomianu[366].X = 3.66
	punktyWykresuWielomianu[366].Y = -1_120.051

	punktyWykresuWielomianu[367].X = 3.67
	punktyWykresuWielomianu[367].Y = -1_126.018

	punktyWykresuWielomianu[368].X = 3.68
	punktyWykresuWielomianu[368].Y = -1_132.005

	punktyWykresuWielomianu[369].X = 3.69
	punktyWykresuWielomianu[369].Y = -1_138.014

	punktyWykresuWielomianu[370].X = 3.7
	punktyWykresuWielomianu[370].Y = -1_144.043

	punktyWykresuWielomianu[371].X = 3.71
	punktyWykresuWielomianu[371].Y = -1_150.094

	punktyWykresuWielomianu[372].X = 3.72
	punktyWykresuWielomianu[372].Y = -1_156.167

	punktyWykresuWielomianu[373].X = 3.73
	punktyWykresuWielomianu[373].Y = -1_162.26

	punktyWykresuWielomianu[374].X = 3.74
	punktyWykresuWielomianu[374].Y = -1_168.375

	punktyWykresuWielomianu[375].X = 3.75
	punktyWykresuWielomianu[375].Y = -1_174.511

	punktyWykresuWielomianu[376].X = 3.76
	punktyWykresuWielomianu[376].Y = -1_180.669

	punktyWykresuWielomianu[377].X = 3.77
	punktyWykresuWielomianu[377].Y = -1_186.848

	punktyWykresuWielomianu[378].X = 3.78
	punktyWykresuWielomianu[378].Y = -1_193.048

	punktyWykresuWielomianu[379].X = 3.79
	punktyWykresuWielomianu[379].Y = -1_199.27

	punktyWykresuWielomianu[380].X = 3.8
	punktyWykresuWielomianu[380].Y = -1_205.513

	punktyWykresuWielomianu[381].X = 3.81
	punktyWykresuWielomianu[381].Y = -1_211.778

	punktyWykresuWielomianu[382].X = 3.82
	punktyWykresuWielomianu[382].Y = -1_218.065

	punktyWykresuWielomianu[383].X = 3.83
	punktyWykresuWielomianu[383].Y = -1_224.373

	punktyWykresuWielomianu[384].X = 3.84
	punktyWykresuWielomianu[384].Y = -1_230.702

	punktyWykresuWielomianu[385].X = 3.85
	punktyWykresuWielomianu[385].Y = -1_237.054

	punktyWykresuWielomianu[386].X = 3.86
	punktyWykresuWielomianu[386].Y = -1_243.427

	punktyWykresuWielomianu[387].X = 3.87
	punktyWykresuWielomianu[387].Y = -1_249.821

	punktyWykresuWielomianu[388].X = 3.88
	punktyWykresuWielomianu[388].Y = -1_256.238

	punktyWykresuWielomianu[389].X = 3.89
	punktyWykresuWielomianu[389].Y = -1_262.676

	punktyWykresuWielomianu[390].X = 3.9
	punktyWykresuWielomianu[390].Y = -1_269.135

	punktyWykresuWielomianu[391].X = 3.91
	punktyWykresuWielomianu[391].Y = -1_275.617

	punktyWykresuWielomianu[392].X = 3.92
	punktyWykresuWielomianu[392].Y = -1_282.121

	punktyWykresuWielomianu[393].X = 3.93
	punktyWykresuWielomianu[393].Y = -1_288.646

	punktyWykresuWielomianu[394].X = 3.94
	punktyWykresuWielomianu[394].Y = -1_295.193

	punktyWykresuWielomianu[395].X = 3.95
	punktyWykresuWielomianu[395].Y = -1_301.762

	punktyWykresuWielomianu[396].X = 3.96
	punktyWykresuWielomianu[396].Y = -1_308.353

	punktyWykresuWielomianu[397].X = 3.97
	punktyWykresuWielomianu[397].Y = -1_314.966

	punktyWykresuWielomianu[398].X = 3.98
	punktyWykresuWielomianu[398].Y = -1_321.601

	punktyWykresuWielomianu[399].X = 3.99
	punktyWykresuWielomianu[399].Y = -1_328.258

	punktyWykresuWielomianu[400].X = 4.0
	punktyWykresuWielomianu[400].Y = -1_334.937

	punktyWykresuWielomianu[401].X = 4.01
	punktyWykresuWielomianu[401].Y = -1_341.638

	punktyWykresuWielomianu[402].X = 4.02
	punktyWykresuWielomianu[402].Y = -1_348.361

	punktyWykresuWielomianu[403].X = 4.03
	punktyWykresuWielomianu[403].Y = -1_355.106

	punktyWykresuWielomianu[404].X = 4.04
	punktyWykresuWielomianu[404].Y = -1_361.874

	punktyWykresuWielomianu[405].X = 4.05
	punktyWykresuWielomianu[405].Y = -1_368.663

	punktyWykresuWielomianu[406].X = 4.06
	punktyWykresuWielomianu[406].Y = -1_375.475

	punktyWykresuWielomianu[407].X = 4.07
	punktyWykresuWielomianu[407].Y = -1_382.309

	punktyWykresuWielomianu[408].X = 4.08
	punktyWykresuWielomianu[408].Y = -1_389.165

	punktyWykresuWielomianu[409].X = 4.09
	punktyWykresuWielomianu[409].Y = -1_396.043

	punktyWykresuWielomianu[410].X = 4.1
	punktyWykresuWielomianu[410].Y = -1_402.943

	punktyWykresuWielomianu[411].X = 4.11
	punktyWykresuWielomianu[411].Y = -1_409.866

	punktyWykresuWielomianu[412].X = 4.12
	punktyWykresuWielomianu[412].Y = -1_416.811

	punktyWykresuWielomianu[413].X = 4.13
	punktyWykresuWielomianu[413].Y = -1_423.779

	punktyWykresuWielomianu[414].X = 4.14
	punktyWykresuWielomianu[414].Y = -1_430.769

	punktyWykresuWielomianu[415].X = 4.15
	punktyWykresuWielomianu[415].Y = -1_437.781

	punktyWykresuWielomianu[416].X = 4.16
	punktyWykresuWielomianu[416].Y = -1_444.815

	punktyWykresuWielomianu[417].X = 4.17
	punktyWykresuWielomianu[417].Y = -1_451.872

	punktyWykresuWielomianu[418].X = 4.18
	punktyWykresuWielomianu[418].Y = -1_458.951

	punktyWykresuWielomianu[419].X = 4.19
	punktyWykresuWielomianu[419].Y = -1_466.053

	punktyWykresuWielomianu[420].X = 4.20
	punktyWykresuWielomianu[420].Y = -1_473.177

	punktyWykresuWielomianu[421].X = 4.21
	punktyWykresuWielomianu[421].Y = -1_480.324

	punktyWykresuWielomianu[422].X = 4.22
	punktyWykresuWielomianu[422].Y = -1_487.493

	punktyWykresuWielomianu[423].X = 4.23
	punktyWykresuWielomianu[423].Y = -1_494.685

	punktyWykresuWielomianu[424].X = 4.24
	punktyWykresuWielomianu[424].Y = -1_501.899

	punktyWykresuWielomianu[425].X = 4.25
	punktyWykresuWielomianu[425].Y = -1_509.136

	punktyWykresuWielomianu[426].X = 4.26
	punktyWykresuWielomianu[426].Y = -1_516.396

	punktyWykresuWielomianu[427].X = 4.27
	punktyWykresuWielomianu[427].Y = -1_523.678

	punktyWykresuWielomianu[428].X = 4.28
	punktyWykresuWielomianu[428].Y = -1_530.982

	punktyWykresuWielomianu[429].X = 4.29
	punktyWykresuWielomianu[429].Y = -1_538.31

	punktyWykresuWielomianu[430].X = 4.30
	punktyWykresuWielomianu[430].Y = -1_545.659

	punktyWykresuWielomianu[431].X = 4.31
	punktyWykresuWielomianu[431].Y = -1_553.032

	punktyWykresuWielomianu[432].X = 4.32
	punktyWykresuWielomianu[432].Y = -1_560.427

	punktyWykresuWielomianu[433].X = 4.33
	punktyWykresuWielomianu[433].Y = -1_567.845

	punktyWykresuWielomianu[434].X = 4.34
	punktyWykresuWielomianu[434].Y = -1_575.286

	punktyWykresuWielomianu[435].X = 4.35
	punktyWykresuWielomianu[435].Y = -1_582.749

	punktyWykresuWielomianu[436].X = 4.36
	punktyWykresuWielomianu[436].Y = -1_590.235

	punktyWykresuWielomianu[437].X = 4.37
	punktyWykresuWielomianu[437].Y = -1_597.744

	punktyWykresuWielomianu[438].X = 4.38
	punktyWykresuWielomianu[438].Y = -1_605.276

	punktyWykresuWielomianu[439].X = 4.39
	punktyWykresuWielomianu[439].Y = -1_612.83

	punktyWykresuWielomianu[440].X = 4.40
	punktyWykresuWielomianu[440].Y = -1_620.407

	punktyWykresuWielomianu[441].X = 4.41
	punktyWykresuWielomianu[441].Y = -1_628.008

	punktyWykresuWielomianu[442].X = 4.42
	punktyWykresuWielomianu[442].Y = -1_635.63

	punktyWykresuWielomianu[443].X = 4.43
	punktyWykresuWielomianu[443].Y = -1_643.276

	punktyWykresuWielomianu[444].X = 4.44
	punktyWykresuWielomianu[444].Y = -1_650.945

	punktyWykresuWielomianu[445].X = 4.45
	punktyWykresuWielomianu[445].Y = -1_658.636

	punktyWykresuWielomianu[446].X = 4.46
	punktyWykresuWielomianu[446].Y = -1_666.351

	punktyWykresuWielomianu[447].X = 4.47
	punktyWykresuWielomianu[447].Y = -1_674.088

	punktyWykresuWielomianu[448].X = 4.48
	punktyWykresuWielomianu[448].Y = -1_681.848

	punktyWykresuWielomianu[449].X = 4.49
	punktyWykresuWielomianu[449].Y = -1_689.631

	punktyWykresuWielomianu[450].X = 4.5
	punktyWykresuWielomianu[450].Y = -1_697.437

	punktyWykresuWielomianu[451].X = 4.51
	punktyWykresuWielomianu[451].Y = -1_705.266

	punktyWykresuWielomianu[452].X = 4.52
	punktyWykresuWielomianu[452].Y = -1_713.118

	punktyWykresuWielomianu[453].X = 4.53
	punktyWykresuWielomianu[453].Y = -1_720.993

	punktyWykresuWielomianu[454].X = 4.54
	punktyWykresuWielomianu[454].Y = -1_728.891

	punktyWykresuWielomianu[455].X = 4.55
	punktyWykresuWielomianu[455].Y = -1_736.812

	punktyWykresuWielomianu[456].X = 4.56
	punktyWykresuWielomianu[456].Y = -1_744.756

	punktyWykresuWielomianu[457].X = 4.57
	punktyWykresuWielomianu[457].Y = -1_752.722

	punktyWykresuWielomianu[458].X = 4.58
	punktyWykresuWielomianu[458].Y = -1_760.712

	punktyWykresuWielomianu[459].X = 4.59
	punktyWykresuWielomianu[459].Y = -1_768.725

	punktyWykresuWielomianu[460].X = 4.6
	punktyWykresuWielomianu[460].Y = -1_776.761

	punktyWykresuWielomianu[461].X = 4.61
	punktyWykresuWielomianu[461].Y = -1_784.821

	punktyWykresuWielomianu[462].X = 4.62
	punktyWykresuWielomianu[462].Y = -1_792.309

	punktyWykresuWielomianu[463].X = 4.63
	punktyWykresuWielomianu[463].Y = -1_801.008

	punktyWykresuWielomianu[464].X = 4.64
	punktyWykresuWielomianu[464].Y = -1_809.136

	punktyWykresuWielomianu[465].X = 4.65
	punktyWykresuWielomianu[465].Y = -1_817.288

	punktyWykresuWielomianu[466].X = 4.66
	punktyWykresuWielomianu[466].Y = -1_825.462

	punktyWykresuWielomianu[467].X = 4.67
	punktyWykresuWielomianu[467].Y = -1_833.66

	punktyWykresuWielomianu[468].X = 4.68
	punktyWykresuWielomianu[468].Y = -1_841.881

	punktyWykresuWielomianu[469].X = 4.69
	punktyWykresuWielomianu[469].Y = -1_850.124

	punktyWykresuWielomianu[470].X = 4.7
	punktyWykresuWielomianu[470].Y = -1_858.391

	punktyWykresuWielomianu[471].X = 4.71
	punktyWykresuWielomianu[471].Y = -1_866.682

	punktyWykresuWielomianu[472].X = 4.72
	punktyWykresuWielomianu[472].Y = -1_874.995

	punktyWykresuWielomianu[473].X = 4.73
	punktyWykresuWielomianu[473].Y = -1_883.331

	punktyWykresuWielomianu[474].X = 4.74
	punktyWykresuWielomianu[474].Y = -1_891.691

	punktyWykresuWielomianu[475].X = 4.75
	punktyWykresuWielomianu[475].Y = -1_900.074

	punktyWykresuWielomianu[476].X = 4.76
	punktyWykresuWielomianu[476].Y = -1_908.48

	punktyWykresuWielomianu[477].X = 4.77
	punktyWykresuWielomianu[477].Y = -1_916.909

	punktyWykresuWielomianu[478].X = 4.78
	punktyWykresuWielomianu[478].Y = -1_925.361

	punktyWykresuWielomianu[479].X = 4.79
	punktyWykresuWielomianu[479].Y = -1_933.837

	punktyWykresuWielomianu[480].X = 4.8
	punktyWykresuWielomianu[480].Y = -1_942.335

	punktyWykresuWielomianu[481].X = 4.81
	punktyWykresuWielomianu[481].Y = -1_950.857

	punktyWykresuWielomianu[482].X = 4.82
	punktyWykresuWielomianu[482].Y = -1_959.403

	punktyWykresuWielomianu[483].X = 4.83
	punktyWykresuWielomianu[483].Y = -1_967.971

	punktyWykresuWielomianu[484].X = 4.84
	punktyWykresuWielomianu[484].Y = -1_976.562

	punktyWykresuWielomianu[485].X = 4.85
	punktyWykresuWielomianu[485].Y = -1_985.177

	punktyWykresuWielomianu[486].X = 4.86
	punktyWykresuWielomianu[486].Y = -1_993.815

	punktyWykresuWielomianu[487].X = 4.87
	punktyWykresuWielomianu[487].Y = -2_002.476

	punktyWykresuWielomianu[488].X = 4.88
	punktyWykresuWielomianu[488].Y = -2_011.161

	punktyWykresuWielomianu[489].X = 4.89
	punktyWykresuWielomianu[489].Y = -2_019.869

	punktyWykresuWielomianu[490].X = 4.9
	punktyWykresuWielomianu[490].Y = -2_028.599

	punktyWykresuWielomianu[491].X = 4.91
	punktyWykresuWielomianu[491].Y = -2_037.354

	punktyWykresuWielomianu[492].X = 4.92
	punktyWykresuWielomianu[492].Y = -2_046.131

	punktyWykresuWielomianu[493].X = 4.93
	punktyWykresuWielomianu[493].Y = -2_054.932

	punktyWykresuWielomianu[494].X = 4.94
	punktyWykresuWielomianu[494].Y = -2_063.756

	punktyWykresuWielomianu[495].X = 4.95
	punktyWykresuWielomianu[495].Y = -2_072.603

	punktyWykresuWielomianu[496].X = 4.96
	punktyWykresuWielomianu[496].Y = -2_081.473

	punktyWykresuWielomianu[497].X = 4.97
	punktyWykresuWielomianu[497].Y = -2_090.367

	punktyWykresuWielomianu[498].X = 4.98
	punktyWykresuWielomianu[498].Y = -2_099.284

	punktyWykresuWielomianu[499].X = 4.99
	punktyWykresuWielomianu[499].Y = -2_108.224

	punktyWykresuWielomianu[500].X = 5.0
	punktyWykresuWielomianu[500].Y = -2_117.187

	punktyWykresuWielomianu[501].X = 5.01
	punktyWykresuWielomianu[501].Y = -2_122.174

	punktyWykresuWielomianu[502].X = 5.02
	punktyWykresuWielomianu[502].Y = -2_135.184

	punktyWykresuWielomianu[503].X = 5.03
	punktyWykresuWielomianu[503].Y = -2_144.217

	punktyWykresuWielomianu[504].X = 5.04
	punktyWykresuWielomianu[504].Y = -2_153.273

	punktyWykresuWielomianu[505].X = 5.05
	punktyWykresuWielomianu[505].Y = -2_162.353

	punktyWykresuWielomianu[506].X = 5.06
	punktyWykresuWielomianu[506].Y = -2_171.456

	punktyWykresuWielomianu[507].X = 5.07
	punktyWykresuWielomianu[507].Y = -2_180.582

	punktyWykresuWielomianu[508].X = 5.08
	punktyWykresuWielomianu[508].Y = -2_189.731

	punktyWykresuWielomianu[509].X = 5.09
	punktyWykresuWielomianu[509].Y = -2_198.904

	punktyWykresuWielomianu[510].X = 5.1
	punktyWykresuWielomianu[510].Y = -2_208.099

	punktyWykresuWielomianu[511].X = 5.11
	punktyWykresuWielomianu[511].Y = -2_217.319

	punktyWykresuWielomianu[512].X = 5.12
	punktyWykresuWielomianu[512].Y = -2_226.561

	punktyWykresuWielomianu[513].X = 5.13
	punktyWykresuWielomianu[513].Y = -2_235.826

	punktyWykresuWielomianu[514].X = 5.14
	punktyWykresuWielomianu[514].Y = -2_245.115

	punktyWykresuWielomianu[515].X = 5.15
	punktyWykresuWielomianu[515].Y = -2_254.427

	punktyWykresuWielomianu[516].X = 5.16
	punktyWykresuWielomianu[516].Y = -2_263.762

	punktyWykresuWielomianu[517].X = 5.17
	punktyWykresuWielomianu[517].Y = -2_273.121

	punktyWykresuWielomianu[518].X = 5.18
	punktyWykresuWielomianu[518].Y = -2_282.503

	punktyWykresuWielomianu[519].X = 5.19
	punktyWykresuWielomianu[519].Y = -2_291.907

	punktyWykresuWielomianu[520].X = 5.2
	punktyWykresuWielomianu[520].Y = -2_301.335

	punktyWykresuWielomianu[521].X = 5.21
	punktyWykresuWielomianu[521].Y = -2_310.787

	punktyWykresuWielomianu[522].X = 5.22
	punktyWykresuWielomianu[522].Y = -2_320.261

	punktyWykresuWielomianu[523].X = 5.23
	punktyWykresuWielomianu[523].Y = -2_329.759

	punktyWykresuWielomianu[524].X = 5.24
	punktyWykresuWielomianu[524].Y = -2_339.28

	punktyWykresuWielomianu[525].X = 5.25
	punktyWykresuWielomianu[525].Y = -2_348.824

	punktyWykresuWielomianu[526].X = 5.26
	punktyWykresuWielomianu[526].Y = -2_358.391

	punktyWykresuWielomianu[527].X = 5.27
	punktyWykresuWielomianu[527].Y = -2_367.981

	punktyWykresuWielomianu[528].X = 5.28
	punktyWykresuWielomianu[528].Y = -2_377.595

	punktyWykresuWielomianu[529].X = 5.29
	punktyWykresuWielomianu[529].Y = -2_387.232

	punktyWykresuWielomianu[530].X = 5.3
	punktyWykresuWielomianu[530].Y = -2_396.891

	punktyWykresuWielomianu[531].X = 5.31
	punktyWykresuWielomianu[531].Y = -2_406.574

	punktyWykresuWielomianu[532].X = 5.32
	punktyWykresuWielomianu[532].Y = -2_416.281

	punktyWykresuWielomianu[533].X = 5.33
	punktyWykresuWielomianu[533].Y = -2_426.01

	punktyWykresuWielomianu[534].X = 5.34
	punktyWykresuWielomianu[534].Y = -2_435.762

	punktyWykresuWielomianu[535].X = 5.35
	punktyWykresuWielomianu[535].Y = -2_445.538

	punktyWykresuWielomianu[536].X = 5.36
	punktyWykresuWielomianu[536].Y = -2_455.336

	punktyWykresuWielomianu[537].X = 5.37
	punktyWykresuWielomianu[537].Y = -2_465.158

	punktyWykresuWielomianu[538].X = 5.38
	punktyWykresuWielomianu[538].Y = -2_475.003

	punktyWykresuWielomianu[539].X = 5.39
	punktyWykresuWielomianu[539].Y = -2_484.871

	punktyWykresuWielomianu[540].X = 5.4
	punktyWykresuWielomianu[540].Y = -2_494.761

	punktyWykresuWielomianu[541].X = 5.41
	punktyWykresuWielomianu[541].Y = -2_504.675

	punktyWykresuWielomianu[542].X = 5.42
	punktyWykresuWielomianu[542].Y = -2_514.612

	punktyWykresuWielomianu[543].X = 5.43
	punktyWykresuWielomianu[543].Y = -2_524.572

	punktyWykresuWielomianu[544].X = 5.44
	punktyWykresuWielomianu[544].Y = -2_534.556

	punktyWykresuWielomianu[545].X = 5.45
	punktyWykresuWielomianu[545].Y = -2_544.562

	punktyWykresuWielomianu[546].X = 5.46
	punktyWykresuWielomianu[546].Y = -2_554.591

	punktyWykresuWielomianu[547].X = 5.47
	punktyWykresuWielomianu[547].Y = -2_564.643

	punktyWykresuWielomianu[548].X = 5.48
	punktyWykresuWielomianu[548].Y = -2_574.718

	punktyWykresuWielomianu[549].X = 5.49
	punktyWykresuWielomianu[549].Y = -2_584.816

	punktyWykresuWielomianu[550].X = 5.5
	punktyWykresuWielomianu[550].Y = -2_594.937

	punktyWykresuWielomianu[551].X = 5.51
	punktyWykresuWielomianu[551].Y = -2_605.081

	punktyWykresuWielomianu[552].X = 5.52
	punktyWykresuWielomianu[552].Y = -2_615.248

	punktyWykresuWielomianu[553].X = 5.53
	punktyWykresuWielomianu[553].Y = -2_625.438

	punktyWykresuWielomianu[554].X = 5.54
	punktyWykresuWielomianu[554].Y = -2_635.651

	punktyWykresuWielomianu[555].X = 5.55
	punktyWykresuWielomianu[555].Y = -2_645.886

	punktyWykresuWielomianu[556].X = 5.56
	punktyWykresuWielomianu[556].Y = -2_656.145

	punktyWykresuWielomianu[557].X = 5.57
	punktyWykresuWielomianu[557].Y = -2_666.426

	punktyWykresuWielomianu[558].X = 5.58
	punktyWykresuWielomianu[558].Y = -2_676.73

	punktyWykresuWielomianu[559].X = 5.59
	punktyWykresuWielomianu[559].Y = -2_686.058

	punktyWykresuWielomianu[560].X = 5.6
	punktyWykresuWielomianu[560].Y = -2_697.407

	punktyWykresuWielomianu[561].X = 5.61
	punktyWykresuWielomianu[561].Y = -2_707.78

	punktyWykresuWielomianu[562].X = 5.62
	punktyWykresuWielomianu[562].Y = -2_718.176

	punktyWykresuWielomianu[563].X = 5.63
	punktyWykresuWielomianu[563].Y = -2_728.594

	punktyWykresuWielomianu[564].X = 5.64
	punktyWykresuWielomianu[564].Y = -2_739.035

	punktyWykresuWielomianu[565].X = 5.65
	punktyWykresuWielomianu[565].Y = -2_749.499

	punktyWykresuWielomianu[566].X = 5.66
	punktyWykresuWielomianu[566].Y = -2_759.986

	punktyWykresuWielomianu[567].X = 5.67
	punktyWykresuWielomianu[567].Y = -2_770.495

	punktyWykresuWielomianu[568].X = 5.68
	punktyWykresuWielomianu[568].Y = -2_781.027

	punktyWykresuWielomianu[569].X = 5.69
	punktyWykresuWielomianu[569].Y = -2_791.582

	punktyWykresuWielomianu[570].X = 5.7
	punktyWykresuWielomianu[570].Y = -2_802.159

	punktyWykresuWielomianu[571].X = 5.71
	punktyWykresuWielomianu[571].Y = -2_812.76

	punktyWykresuWielomianu[572].X = 5.72
	punktyWykresuWielomianu[572].Y = -2_823.382

	punktyWykresuWielomianu[573].X = 5.73
	punktyWykresuWielomianu[573].Y = -2_834.028

	punktyWykresuWielomianu[574].X = 5.74
	punktyWykresuWielomianu[574].Y = -2_844.696

	punktyWykresuWielomianu[575].X = 5.75
	punktyWykresuWielomianu[575].Y = -2_855.386

	punktyWykresuWielomianu[576].X = 5.76
	punktyWykresuWielomianu[576].Y = -2_866.099

	punktyWykresuWielomianu[577].X = 5.77
	punktyWykresuWielomianu[577].Y = -2_876.835

	punktyWykresuWielomianu[578].X = 5.78
	punktyWykresuWielomianu[578].Y = -2_887.593

	punktyWykresuWielomianu[579].X = 5.79
	punktyWykresuWielomianu[579].Y = -2_898.374

	punktyWykresuWielomianu[580].X = 5.8
	punktyWykresuWielomianu[580].Y = -2_909.177

	punktyWykresuWielomianu[581].X = 5.81
	punktyWykresuWielomianu[581].Y = -2_920.003

	punktyWykresuWielomianu[582].X = 5.82
	punktyWykresuWielomianu[582].Y = -2_930.851

	punktyWykresuWielomianu[583].X = 5.83
	punktyWykresuWielomianu[583].Y = -2_941.722

	punktyWykresuWielomianu[584].X = 5.84
	punktyWykresuWielomianu[584].Y = -2_952.615

	punktyWykresuWielomianu[585].X = 5.85
	punktyWykresuWielomianu[585].Y = -2_963.531

	punktyWykresuWielomianu[586].X = 5.86
	punktyWykresuWielomianu[586].Y = -2_974.469

	punktyWykresuWielomianu[587].X = 5.87
	punktyWykresuWielomianu[587].Y = -2_985.429

	punktyWykresuWielomianu[588].X = 5.88
	punktyWykresuWielomianu[588].Y = -2_996.411

	punktyWykresuWielomianu[589].X = 5.89
	punktyWykresuWielomianu[589].Y = -3_007.416

	punktyWykresuWielomianu[590].X = 5.9
	punktyWykresuWielomianu[590].Y = -3_018.443

	punktyWykresuWielomianu[591].X = 5.91
	punktyWykresuWielomianu[591].Y = -3_029.493

	punktyWykresuWielomianu[592].X = 5.92
	punktyWykresuWielomianu[592].Y = -3_040.565

	punktyWykresuWielomianu[593].X = 5.93
	punktyWykresuWielomianu[593].Y = -3_051.659

	punktyWykresuWielomianu[594].X = 5.94
	punktyWykresuWielomianu[594].Y = -3_062.775

	punktyWykresuWielomianu[595].X = 5.95
	punktyWykresuWielomianu[595].Y = -3_073.913

	punktyWykresuWielomianu[596].X = 5.96
	punktyWykresuWielomianu[596].Y = -3_085.074

	punktyWykresuWielomianu[597].X = 5.97
	punktyWykresuWielomianu[597].Y = -3_096.256

	punktyWykresuWielomianu[598].X = 5.98
	punktyWykresuWielomianu[598].Y = -3_107.461

	punktyWykresuWielomianu[599].X = 5.99
	punktyWykresuWielomianu[599].Y = -3_118.688

	punktyWykresuWielomianu[600].X = 6.0
	punktyWykresuWielomianu[600].Y = -3_129.937

	punktyWykresuWielomianu[601].X = 6.01
	punktyWykresuWielomianu[601].Y = -3_141.208

	punktyWykresuWielomianu[602].X = 6.02
	punktyWykresuWielomianu[602].Y = -3_152.501

	punktyWykresuWielomianu[603].X = 6.03
	punktyWykresuWielomianu[603].Y = -3_163.816

	punktyWykresuWielomianu[604].X = 6.04
	punktyWykresuWielomianu[604].Y = -3_175.153

	punktyWykresuWielomianu[605].X = 6.05
	punktyWykresuWielomianu[605].Y = -3_186.512

	punktyWykresuWielomianu[606].X = 6.06
	punktyWykresuWielomianu[606].Y = -3_197.893

	punktyWykresuWielomianu[607].X = 6.07
	punktyWykresuWielomianu[607].Y = -3_209.296

	punktyWykresuWielomianu[608].X = 6.08
	punktyWykresuWielomianu[608].Y = -3_220.721

	punktyWykresuWielomianu[609].X = 6.09
	punktyWykresuWielomianu[609].Y = -3_232.167

	punktyWykresuWielomianu[610].X = 6.1
	punktyWykresuWielomianu[610].Y = -3_243.635

	punktyWykresuWielomianu[611].X = 6.11
	punktyWykresuWielomianu[611].Y = -3_255.126

	punktyWykresuWielomianu[612].X = 6.12
	punktyWykresuWielomianu[612].Y = -3_266.638

	punktyWykresuWielomianu[613].X = 6.13
	punktyWykresuWielomianu[613].Y = -3_278.171

	punktyWykresuWielomianu[614].X = 6.14
	punktyWykresuWielomianu[614].Y = -3_289.727

	punktyWykresuWielomianu[615].X = 6.15
	punktyWykresuWielomianu[615].Y = -3_301.304

	punktyWykresuWielomianu[616].X = 6.16
	punktyWykresuWielomianu[616].Y = -3_312.902

	punktyWykresuWielomianu[617].X = 6.17
	punktyWykresuWielomianu[617].Y = -3_324.523

	punktyWykresuWielomianu[618].X = 6.18
	punktyWykresuWielomianu[618].Y = -3_336.165

	punktyWykresuWielomianu[619].X = 6.19
	punktyWykresuWielomianu[619].Y = -3_347.828

	punktyWykresuWielomianu[620].X = 6.2
	punktyWykresuWielomianu[620].Y = -3_359.513

	punktyWykresuWielomianu[621].X = 6.21
	punktyWykresuWielomianu[621].Y = -3_371.22

	punktyWykresuWielomianu[622].X = 6.22
	punktyWykresuWielomianu[622].Y = -3_382.948

	punktyWykresuWielomianu[623].X = 6.23
	punktyWykresuWielomianu[623].Y = -3_394.698

	punktyWykresuWielomianu[624].X = 6.24
	punktyWykresuWielomianu[624].Y = -3_406.469

	punktyWykresuWielomianu[625].X = 6.25
	punktyWykresuWielomianu[625].Y = -3_418.261

	punktyWykresuWielomianu[626].X = 6.26
	punktyWykresuWielomianu[626].Y = -3_430.075

	punktyWykresuWielomianu[627].X = 6.27
	punktyWykresuWielomianu[627].Y = -3_441.91

	punktyWykresuWielomianu[628].X = 6.28
	punktyWykresuWielomianu[628].Y = -3_453.767

	punktyWykresuWielomianu[629].X = 6.29
	punktyWykresuWielomianu[629].Y = -3_465.644

	punktyWykresuWielomianu[630].X = 6.3
	punktyWykresuWielomianu[630].Y = -3_477.543

	punktyWykresuWielomianu[631].X = 6.31
	punktyWykresuWielomianu[631].Y = -3_489.464

	punktyWykresuWielomianu[632].X = 6.32
	punktyWykresuWielomianu[632].Y = -3_501.405

	punktyWykresuWielomianu[633].X = 6.33
	punktyWykresuWielomianu[633].Y = -3_513.368

	punktyWykresuWielomianu[634].X = 6.34
	punktyWykresuWielomianu[634].Y = -3_525.351

	punktyWykresuWielomianu[635].X = 6.35
	punktyWykresuWielomianu[635].Y = -3_537.356

	punktyWykresuWielomianu[636].X = 6.36
	punktyWykresuWielomianu[636].Y = -3_549.382

	punktyWykresuWielomianu[637].X = 6.37
	punktyWykresuWielomianu[637].Y = -3_561.429

	punktyWykresuWielomianu[638].X = 6.38
	punktyWykresuWielomianu[638].Y = -3_573.497

	punktyWykresuWielomianu[639].X = 6.39
	punktyWykresuWielomianu[639].Y = -3_585.586

	punktyWykresuWielomianu[640].X = 6.4
	punktyWykresuWielomianu[640].Y = -3_597.695

	punktyWykresuWielomianu[641].X = 6.41
	punktyWykresuWielomianu[641].Y = -3_609.826

	punktyWykresuWielomianu[642].X = 6.42
	punktyWykresuWielomianu[642].Y = -3_621.978

	punktyWykresuWielomianu[643].X = 6.43
	punktyWykresuWielomianu[643].Y = -3_634.15

	punktyWykresuWielomianu[644].X = 6.44
	punktyWykresuWielomianu[644].Y = -3_646.343

	punktyWykresuWielomianu[645].X = 6.45
	punktyWykresuWielomianu[645].Y = -3_658.557

	punktyWykresuWielomianu[646].X = 6.46
	punktyWykresuWielomianu[646].Y = -3_670.792

	punktyWykresuWielomianu[647].X = 6.47
	punktyWykresuWielomianu[647].Y = -3_683.047

	punktyWykresuWielomianu[648].X = 6.48
	punktyWykresuWielomianu[648].Y = -3_695.323

	punktyWykresuWielomianu[649].X = 6.49
	punktyWykresuWielomianu[649].Y = -3_707.62

	punktyWykresuWielomianu[650].X = 6.5
	punktyWykresuWielomianu[650].Y = -3_719.937

	punktyWykresuWielomianu[651].X = 6.51
	punktyWykresuWielomianu[651].Y = -3_732.275

	punktyWykresuWielomianu[652].X = 6.52
	punktyWykresuWielomianu[652].Y = -3_744.633

	punktyWykresuWielomianu[653].X = 6.53
	punktyWykresuWielomianu[653].Y = -3_757.012

	punktyWykresuWielomianu[654].X = 6.54
	punktyWykresuWielomianu[654].Y = -3_769.411

	punktyWykresuWielomianu[655].X = 6.55
	punktyWykresuWielomianu[655].Y = -3_781.831

	punktyWykresuWielomianu[656].X = 6.56
	punktyWykresuWielomianu[656].Y = -3_794.271

	punktyWykresuWielomianu[657].X = 6.57
	punktyWykresuWielomianu[657].Y = -3_806.731

	punktyWykresuWielomianu[658].X = 6.58
	punktyWykresuWielomianu[658].Y = -3_819.212

	punktyWykresuWielomianu[659].X = 6.59
	punktyWykresuWielomianu[659].Y = -3_831.712

	punktyWykresuWielomianu[660].X = 6.6
	punktyWykresuWielomianu[660].Y = -3_844.233

	punktyWykresuWielomianu[661].X = 6.61
	punktyWykresuWielomianu[661].Y = -3_856.775

	punktyWykresuWielomianu[662].X = 6.62
	punktyWykresuWielomianu[662].Y = -3_869.336

	punktyWykresuWielomianu[663].X = 6.63
	punktyWykresuWielomianu[663].Y = -3_881.918

	punktyWykresuWielomianu[664].X = 6.64
	punktyWykresuWielomianu[664].Y = -3_894.519

	punktyWykresuWielomianu[665].X = 6.65
	punktyWykresuWielomianu[665].Y = -3_907.141

	punktyWykresuWielomianu[666].X = 6.66
	punktyWykresuWielomianu[666].Y = -3_919.782

	punktyWykresuWielomianu[667].X = 6.67
	punktyWykresuWielomianu[667].Y = -3_932.444

	punktyWykresuWielomianu[668].X = 6.68
	punktyWykresuWielomianu[668].Y = -3_945.125

	punktyWykresuWielomianu[669].X = 6.69
	punktyWykresuWielomianu[669].Y = -3_957.826

	punktyWykresuWielomianu[670].X = 6.7
	punktyWykresuWielomianu[670].Y = -3_970.547

	punktyWykresuWielomianu[671].X = 6.71
	punktyWykresuWielomianu[671].Y = -3_983.288

	punktyWykresuWielomianu[672].X = 6.72
	punktyWykresuWielomianu[672].Y = -3_996.049

	punktyWykresuWielomianu[673].X = 6.73
	punktyWykresuWielomianu[673].Y = -4_008.829

	punktyWykresuWielomianu[674].X = 6.74
	punktyWykresuWielomianu[674].Y = -4_021.629

	punktyWykresuWielomianu[675].X = 6.75
	punktyWykresuWielomianu[675].Y = -4_034.449

	punktyWykresuWielomianu[676].X = 6.76
	punktyWykresuWielomianu[676].Y = -4_047.288

	punktyWykresuWielomianu[677].X = 6.77
	punktyWykresuWielomianu[677].Y = -4_060.147

	punktyWykresuWielomianu[678].X = 6.78
	punktyWykresuWielomianu[678].Y = -4_073.025

	punktyWykresuWielomianu[679].X = 6.79
	punktyWykresuWielomianu[679].Y = -4_085.922

	punktyWykresuWielomianu[680].X = 6.8
	punktyWykresuWielomianu[680].Y = -4_098.839

	punktyWykresuWielomianu[681].X = 6.81
	punktyWykresuWielomianu[681].Y = -4_111.776

	punktyWykresuWielomianu[682].X = 6.82
	punktyWykresuWielomianu[682].Y = -4_124.732

	punktyWykresuWielomianu[683].X = 6.83
	punktyWykresuWielomianu[683].Y = -4_137.707

	punktyWykresuWielomianu[684].X = 6.84
	punktyWykresuWielomianu[684].Y = -4_150.701

	punktyWykresuWielomianu[685].X = 6.85
	punktyWykresuWielomianu[685].Y = -4_163.714

	punktyWykresuWielomianu[686].X = 6.86
	punktyWykresuWielomianu[686].Y = -4_176.747

	punktyWykresuWielomianu[687].X = 6.87
	punktyWykresuWielomianu[687].Y = -4_189.798

	punktyWykresuWielomianu[688].X = 6.88
	punktyWykresuWielomianu[688].Y = -4_202.869

	punktyWykresuWielomianu[689].X = 6.89
	punktyWykresuWielomianu[689].Y = -4_215.959

	punktyWykresuWielomianu[690].X = 6.9
	punktyWykresuWielomianu[690].Y = -4_229.067

	punktyWykresuWielomianu[691].X = 6.91
	punktyWykresuWielomianu[691].Y = -4_242.195

	punktyWykresuWielomianu[692].X = 6.92
	punktyWykresuWielomianu[692].Y = -4_255.342

	punktyWykresuWielomianu[693].X = 6.93
	punktyWykresuWielomianu[693].Y = -4_268.507

	punktyWykresuWielomianu[694].X = 6.94
	punktyWykresuWielomianu[694].Y = -4_281.691

	punktyWykresuWielomianu[695].X = 6.95
	punktyWykresuWielomianu[695].Y = -4_294.894

	punktyWykresuWielomianu[696].X = 6.96
	punktyWykresuWielomianu[696].Y = -4_308.115

	punktyWykresuWielomianu[697].X = 6.97
	punktyWykresuWielomianu[697].Y = -4_321.355

	punktyWykresuWielomianu[698].X = 6.98
	punktyWykresuWielomianu[698].Y = -4_334.614

	punktyWykresuWielomianu[699].X = 6.99
	punktyWykresuWielomianu[699].Y = -4_347.891

	punktyWykresuWielomianu[700].X = 7.0
	punktyWykresuWielomianu[700].Y = -4_361.187

	punktyWykresuWielomianu[701].X = 7.01
	punktyWykresuWielomianu[701].Y = -4_374.501

	punktyWykresuWielomianu[702].X = 7.02
	punktyWykresuWielomianu[702].Y = -4_387.834

	punktyWykresuWielomianu[703].X = 7.03
	punktyWykresuWielomianu[703].Y = -4_401.185

	punktyWykresuWielomianu[704].X = 7.04
	punktyWykresuWielomianu[704].Y = -4_414.554

	punktyWykresuWielomianu[705].X = 7.05
	punktyWykresuWielomianu[705].Y = -4_427.942

	punktyWykresuWielomianu[706].X = 7.06
	punktyWykresuWielomianu[706].Y = -4_441.347

	punktyWykresuWielomianu[707].X = 7.07
	punktyWykresuWielomianu[707].Y = -4_454.771

	punktyWykresuWielomianu[708].X = 7.08
	punktyWykresuWielomianu[708].Y = -4_468.213

	punktyWykresuWielomianu[709].X = 7.09
	punktyWykresuWielomianu[709].Y = -4_481.673

	punktyWykresuWielomianu[710].X = 7.1
	punktyWykresuWielomianu[710].Y = -4_495.151

	punktyWykresuWielomianu[711].X = 7.11
	punktyWykresuWielomianu[711].Y = -4_508.647

	punktyWykresuWielomianu[712].X = 7.12
	punktyWykresuWielomianu[712].Y = -4_522.161

	punktyWykresuWielomianu[713].X = 7.13
	punktyWykresuWielomianu[713].Y = -4_535.693

	punktyWykresuWielomianu[714].X = 7.14
	punktyWykresuWielomianu[714].Y = -4_549.243

	punktyWykresuWielomianu[715].X = 7.15
	punktyWykresuWielomianu[715].Y = -4_562.81

	punktyWykresuWielomianu[716].X = 7.16
	punktyWykresuWielomianu[716].Y = -4_576.395

	punktyWykresuWielomianu[717].X = 7.17
	punktyWykresuWielomianu[717].Y = -4_589.998

	punktyWykresuWielomianu[718].X = 7.18
	punktyWykresuWielomianu[718].Y = -4_603.618

	punktyWykresuWielomianu[719].X = 7.19
	punktyWykresuWielomianu[719].Y = -4_617.256

	punktyWykresuWielomianu[720].X = 7.2
	punktyWykresuWielomianu[720].Y = -4_630.911

	punktyWykresuWielomianu[721].X = 7.21
	punktyWykresuWielomianu[721].Y = -4_644.584

	punktyWykresuWielomianu[722].X = 7.22
	punktyWykresuWielomianu[722].Y = -4_658.274

	punktyWykresuWielomianu[723].X = 7.23
	punktyWykresuWielomianu[723].Y = -4_671.982

	punktyWykresuWielomianu[724].X = 7.24
	punktyWykresuWielomianu[724].Y = -4_685.707

	punktyWykresuWielomianu[725].X = 7.25
	punktyWykresuWielomianu[725].Y = -4_699.449

	punktyWykresuWielomianu[726].X = 7.26
	punktyWykresuWielomianu[726].Y = -4_713.208

	punktyWykresuWielomianu[727].X = 7.27
	punktyWykresuWielomianu[727].Y = -4_726.984

	punktyWykresuWielomianu[728].X = 7.28
	punktyWykresuWielomianu[728].Y = -4_740.778

	punktyWykresuWielomianu[729].X = 7.29
	punktyWykresuWielomianu[729].Y = -4_754.588

	punktyWykresuWielomianu[730].X = 7.3
	punktyWykresuWielomianu[730].Y = -4_768.415

	punktyWykresuWielomianu[731].X = 7.31
	punktyWykresuWielomianu[731].Y = -4_782.26

	punktyWykresuWielomianu[732].X = 7.32
	punktyWykresuWielomianu[732].Y = -4_796.121

	punktyWykresuWielomianu[733].X = 7.33
	punktyWykresuWielomianu[733].Y = -4_809.999

	punktyWykresuWielomianu[734].X = 7.34
	punktyWykresuWielomianu[734].Y = -4_823.893

	punktyWykresuWielomianu[735].X = 7.35
	punktyWykresuWielomianu[735].Y = -4_837.805

	punktyWykresuWielomianu[736].X = 7.36
	punktyWykresuWielomianu[736].Y = -4_851.733

	punktyWykresuWielomianu[737].X = 7.37
	punktyWykresuWielomianu[737].Y = -4_865.677

	punktyWykresuWielomianu[738].X = 7.38
	punktyWykresuWielomianu[738].Y = -4_879.638

	punktyWykresuWielomianu[739].X = 7.39
	punktyWykresuWielomianu[739].Y = -4_893.616

	punktyWykresuWielomianu[740].X = 7.4
	punktyWykresuWielomianu[740].Y = -4_907.609

	punktyWykresuWielomianu[741].X = 7.41
	punktyWykresuWielomianu[741].Y = -4_921.62

	punktyWykresuWielomianu[742].X = 7.42
	punktyWykresuWielomianu[742].Y = -4_935.646

	punktyWykresuWielomianu[743].X = 7.43
	punktyWykresuWielomianu[743].Y = -4_949.689

	punktyWykresuWielomianu[744].X = 7.44
	punktyWykresuWielomianu[744].Y = -4_963.748

	punktyWykresuWielomianu[745].X = 7.45
	punktyWykresuWielomianu[745].Y = -4_977.823

	punktyWykresuWielomianu[746].X = 7.46
	punktyWykresuWielomianu[746].Y = -4_991.914

	punktyWykresuWielomianu[747].X = 7.47
	punktyWykresuWielomianu[747].Y = -5_006.021

	punktyWykresuWielomianu[748].X = 7.48
	punktyWykresuWielomianu[748].Y = -5_020.144

	punktyWykresuWielomianu[749].X = 7.49
	punktyWykresuWielomianu[749].Y = -5_034.282

	punktyWykresuWielomianu[750].X = 7.5
	punktyWykresuWielomianu[750].Y = -5_048.437

	punktyWykresuWielomianu[751].X = 7.51
	punktyWykresuWielomianu[751].Y = -5_062.607

	punktyWykresuWielomianu[752].X = 7.52
	punktyWykresuWielomianu[752].Y = -5_076.793

	punktyWykresuWielomianu[753].X = 7.53
	punktyWykresuWielomianu[753].Y = -5_090.995

	punktyWykresuWielomianu[754].X = 7.54
	punktyWykresuWielomianu[754].Y = -5_105.212

	punktyWykresuWielomianu[755].X = 7.55
	punktyWykresuWielomianu[755].Y = -5_119.445

	punktyWykresuWielomianu[756].X = 7.56
	punktyWykresuWielomianu[756].Y = -5_133.693

	punktyWykresuWielomianu[757].X = 7.57
	punktyWykresuWielomianu[757].Y = -5_147.957

	punktyWykresuWielomianu[758].X = 7.58
	punktyWykresuWielomianu[758].Y = -5_162.236

	punktyWykresuWielomianu[759].X = 7.59
	punktyWykresuWielomianu[759].Y = -5_176.53

	punktyWykresuWielomianu[760].X = 7.6
	punktyWykresuWielomianu[760].Y = -5_190.839

	punktyWykresuWielomianu[761].X = 7.61
	punktyWykresuWielomianu[761].Y = -5_205.164

	punktyWykresuWielomianu[762].X = 7.62
	punktyWykresuWielomianu[762].Y = -5_219.504

	punktyWykresuWielomianu[763].X = 7.63
	punktyWykresuWielomianu[763].Y = -5_233.858

	punktyWykresuWielomianu[764].X = 7.64
	punktyWykresuWielomianu[764].Y = -5_248.228

	punktyWykresuWielomianu[765].X = 7.65
	punktyWykresuWielomianu[765].Y = -5_262.612

	punktyWykresuWielomianu[766].X = 7.66
	punktyWykresuWielomianu[766].Y = -5_277.011

	punktyWykresuWielomianu[767].X = 7.67
	punktyWykresuWielomianu[767].Y = -5_291.425

	punktyWykresuWielomianu[768].X = 7.68
	punktyWykresuWielomianu[768].Y = -5_305.854

	punktyWykresuWielomianu[769].X = 7.69
	punktyWykresuWielomianu[769].Y = -5_320.298

	punktyWykresuWielomianu[770].X = 7.7
	punktyWykresuWielomianu[770].Y = -5_334.755

	punktyWykresuWielomianu[771].X = 7.71
	punktyWykresuWielomianu[771].Y = -5_349.228

	punktyWykresuWielomianu[772].X = 7.72
	punktyWykresuWielomianu[772].Y = -5_363.715

	punktyWykresuWielomianu[773].X = 7.73
	punktyWykresuWielomianu[773].Y = -5_378.216

	punktyWykresuWielomianu[774].X = 7.74
	punktyWykresuWielomianu[774].Y = -5_392.731

	punktyWykresuWielomianu[775].X = 7.75
	punktyWykresuWielomianu[775].Y = -5_407.261

	punktyWykresuWielomianu[776].X = 7.76
	punktyWykresuWielomianu[776].Y = -5_421.805

	punktyWykresuWielomianu[777].X = 7.77
	punktyWykresuWielomianu[777].Y = -5_436.363

	punktyWykresuWielomianu[778].X = 7.78
	punktyWykresuWielomianu[778].Y = -5_450.935

	punktyWykresuWielomianu[779].X = 7.79
	punktyWykresuWielomianu[779].Y = -5_465.521

	punktyWykresuWielomianu[780].X = 7.8
	punktyWykresuWielomianu[780].Y = -5_480.121

	punktyWykresuWielomianu[781].X = 7.81
	punktyWykresuWielomianu[781].Y = -5_494.735

	punktyWykresuWielomianu[782].X = 7.82
	punktyWykresuWielomianu[782].Y = -5_509.363

	punktyWykresuWielomianu[783].X = 7.83
	punktyWykresuWielomianu[783].Y = -5_524.004

	punktyWykresuWielomianu[784].X = 7.84
	punktyWykresuWielomianu[784].Y = -5_538.659

	punktyWykresuWielomianu[785].X = 7.85
	punktyWykresuWielomianu[785].Y = -5_553.328

	punktyWykresuWielomianu[786].X = 7.86
	punktyWykresuWielomianu[786].Y = -5_568.01

	punktyWykresuWielomianu[787].X = 7.87
	punktyWykresuWielomianu[787].Y = -5_582.705

	punktyWykresuWielomianu[788].X = 7.88
	punktyWykresuWielomianu[788].Y = -5_597.414

	punktyWykresuWielomianu[789].X = 7.89
	punktyWykresuWielomianu[789].Y = -5_612.136

	punktyWykresuWielomianu[790].X = 7.9
	punktyWykresuWielomianu[790].Y = -5_626.871

	punktyWykresuWielomianu[791].X = 7.91
	punktyWykresuWielomianu[791].Y = -5_641.62

	punktyWykresuWielomianu[792].X = 7.92
	punktyWykresuWielomianu[792].Y = -5_656.382

	punktyWykresuWielomianu[793].X = 7.93
	punktyWykresuWielomianu[793].Y = -5_671.156

	punktyWykresuWielomianu[794].X = 7.94
	punktyWykresuWielomianu[794].Y = -5_685.944

	punktyWykresuWielomianu[795].X = 7.95
	punktyWykresuWielomianu[795].Y = -5_700.744

	punktyWykresuWielomianu[796].X = 7.96
	punktyWykresuWielomianu[796].Y = -5_715.557

	punktyWykresuWielomianu[797].X = 7.97
	punktyWykresuWielomianu[797].Y = -5_730.383

	punktyWykresuWielomianu[798].X = 7.98
	punktyWykresuWielomianu[798].Y = -5_745.222

	punktyWykresuWielomianu[799].X = 7.99
	punktyWykresuWielomianu[799].Y = -5_760.073

	punktyWykresuWielomianu[800].X = 8.0
	punktyWykresuWielomianu[800].Y = -5_774.937

	punktyWykresuWielomianu[801].X = 8.01
	punktyWykresuWielomianu[801].Y = -5_789.813

	punktyWykresuWielomianu[802].X = 8.02
	punktyWykresuWielomianu[802].Y = -5_804.702

	punktyWykresuWielomianu[803].X = 8.03
	punktyWykresuWielomianu[803].Y = -5_819.603

	punktyWykresuWielomianu[804].X = 8.04
	punktyWykresuWielomianu[804].Y = -5_834.516

	punktyWykresuWielomianu[805].X = 8.05
	punktyWykresuWielomianu[805].Y = -5_849.441

	punktyWykresuWielomianu[806].X = 8.06
	punktyWykresuWielomianu[806].Y = -5_864.379

	punktyWykresuWielomianu[807].X = 8.07
	punktyWykresuWielomianu[807].Y = -5_879.328

	punktyWykresuWielomianu[808].X = 8.08
	punktyWykresuWielomianu[808].Y = -5_894.289

	punktyWykresuWielomianu[809].X = 8.09
	punktyWykresuWielomianu[809].Y = -5_909.262

	punktyWykresuWielomianu[810].X = 8.1
	punktyWykresuWielomianu[810].Y = -5_925.247

	punktyWykresuWielomianu[811].X = 8.11
	punktyWykresuWielomianu[811].Y = -5_939.244

	punktyWykresuWielomianu[812].X = 8.12
	punktyWykresuWielomianu[812].Y = -5_954.253

	punktyWykresuWielomianu[813].X = 8.13
	punktyWykresuWielomianu[813].Y = -5_969.272

	punktyWykresuWielomianu[814].X = 8.14
	punktyWykresuWielomianu[814].Y = -5_984.304

	punktyWykresuWielomianu[815].X = 8.15
	punktyWykresuWielomianu[815].Y = -5_999.347

	punktyWykresuWielomianu[816].X = 8.16
	punktyWykresuWielomianu[816].Y = -6_014.401

	punktyWykresuWielomianu[817].X = 8.17
	punktyWykresuWielomianu[817].Y = -6_029.466

	punktyWykresuWielomianu[818].X = 8.18
	punktyWykresuWielomianu[818].Y = -6_044.543

	punktyWykresuWielomianu[819].X = 8.19
	punktyWykresuWielomianu[819].Y = -6_059.631

	punktyWykresuWielomianu[820].X = 8.2
	punktyWykresuWielomianu[820].Y = -6_074.729

	punktyWykresuWielomianu[821].X = 8.21
	punktyWykresuWielomianu[821].Y = -6_089.839

	punktyWykresuWielomianu[822].X = 8.22
	punktyWykresuWielomianu[822].Y = -6_104.96

	punktyWykresuWielomianu[823].X = 8.23
	punktyWykresuWielomianu[823].Y = -6_120.091

	punktyWykresuWielomianu[824].X = 8.24
	punktyWykresuWielomianu[824].Y = -6_135.233

	punktyWykresuWielomianu[825].X = 8.25
	punktyWykresuWielomianu[825].Y = -6_150.386

	punktyWykresuWielomianu[826].X = 8.26
	punktyWykresuWielomianu[826].Y = -6_165.55

	punktyWykresuWielomianu[827].X = 8.27
	punktyWykresuWielomianu[827].Y = -6_180.724

	punktyWykresuWielomianu[828].X = 8.28
	punktyWykresuWielomianu[828].Y = -6_195.908

	punktyWykresuWielomianu[829].X = 8.29
	punktyWykresuWielomianu[829].Y = -6_211.103

	punktyWykresuWielomianu[830].X = 8.3
	punktyWykresuWielomianu[830].Y = -6_226.307

	punktyWykresuWielomianu[831].X = 8.31
	punktyWykresuWielomianu[831].Y = -6_241.523

	punktyWykresuWielomianu[832].X = 8.32
	punktyWykresuWielomianu[832].Y = -6_256.748

	punktyWykresuWielomianu[833].X = 8.33
	punktyWykresuWielomianu[833].Y = -6_271.983

	punktyWykresuWielomianu[834].X = 8.34
	punktyWykresuWielomianu[834].Y = -6_287.228

	punktyWykresuWielomianu[835].X = 8.35
	punktyWykresuWielomianu[835].Y = -6_302.483

	punktyWykresuWielomianu[836].X = 8.36
	punktyWykresuWielomianu[836].Y = -6_317.748

	punktyWykresuWielomianu[837].X = 8.37
	punktyWykresuWielomianu[837].Y = -6_333.022

	punktyWykresuWielomianu[838].X = 8.38
	punktyWykresuWielomianu[838].Y = -6_348.307

	punktyWykresuWielomianu[839].X = 8.39
	punktyWykresuWielomianu[839].Y = -6_363.6

	punktyWykresuWielomianu[840].X = 8.4
	punktyWykresuWielomianu[840].Y = -6_378.903

	punktyWykresuWielomianu[841].X = 8.41
	punktyWykresuWielomianu[841].Y = -6_394.216

	punktyWykresuWielomianu[842].X = 8.42
	punktyWykresuWielomianu[842].Y = -6_409.538

	punktyWykresuWielomianu[843].X = 8.43
	punktyWykresuWielomianu[843].Y = -6_424.869

	punktyWykresuWielomianu[844].X = 8.44
	punktyWykresuWielomianu[844].Y = -6_440.209

	punktyWykresuWielomianu[845].X = 8.45
	punktyWykresuWielomianu[845].Y = -6_455.558

	punktyWykresuWielomianu[846].X = 8.46
	punktyWykresuWielomianu[846].Y = -6_470.916

	punktyWykresuWielomianu[847].X = 8.47
	punktyWykresuWielomianu[847].Y = -6_486.283

	punktyWykresuWielomianu[848].X = 8.48
	punktyWykresuWielomianu[848].Y = -6_501.659

	punktyWykresuWielomianu[849].X = 8.49
	punktyWykresuWielomianu[849].Y = -6_517.044

	punktyWykresuWielomianu[850].X = 8.5
	punktyWykresuWielomianu[850].Y = -6_532.437

	punktyWykresuWielomianu[851].X = 8.51
	punktyWykresuWielomianu[851].Y = -6_547.839

	punktyWykresuWielomianu[852].X = 8.52
	punktyWykresuWielomianu[852].Y = -6_563.249

	punktyWykresuWielomianu[853].X = 8.53
	punktyWykresuWielomianu[853].Y = -6_578.668

	punktyWykresuWielomianu[854].X = 8.54
	punktyWykresuWielomianu[854].Y = -6_594.095

	punktyWykresuWielomianu[855].X = 8.55
	punktyWykresuWielomianu[855].Y = -6_609.53

	punktyWykresuWielomianu[856].X = 8.56
	punktyWykresuWielomianu[856].Y = -6_624.973

	punktyWykresuWielomianu[857].X = 8.57
	punktyWykresuWielomianu[857].Y = -6_640.424

	punktyWykresuWielomianu[858].X = 8.58
	punktyWykresuWielomianu[858].Y = -6_655.883

	punktyWykresuWielomianu[859].X = 8.59
	punktyWykresuWielomianu[859].Y = -6_671.351

	punktyWykresuWielomianu[860].X = 8.6
	punktyWykresuWielomianu[860].Y = -6_686.825

	punktyWykresuWielomianu[861].X = 8.61
	punktyWykresuWielomianu[861].Y = -6_702.308

	punktyWykresuWielomianu[862].X = 8.62
	punktyWykresuWielomianu[862].Y = -6_717.798

	punktyWykresuWielomianu[863].X = 8.63
	punktyWykresuWielomianu[863].Y = -6_733.296

	punktyWykresuWielomianu[864].X = 8.64
	punktyWykresuWielomianu[864].Y = -6_748.801

	punktyWykresuWielomianu[865].X = 8.65
	punktyWykresuWielomianu[865].Y = -6_764.314

	punktyWykresuWielomianu[866].X = 8.66
	punktyWykresuWielomianu[866].Y = -6_779.833

	punktyWykresuWielomianu[867].X = 8.67
	punktyWykresuWielomianu[867].Y = -6_795.36

	punktyWykresuWielomianu[868].X = 8.68
	punktyWykresuWielomianu[868].Y = -6_810.894

	punktyWykresuWielomianu[869].X = 8.69
	punktyWykresuWielomianu[869].Y = -6_826.435

	punktyWykresuWielomianu[870].X = 8.7
	punktyWykresuWielomianu[870].Y = -6_841.983

	punktyWykresuWielomianu[871].X = 8.71
	punktyWykresuWielomianu[871].Y = -6_857.538

	punktyWykresuWielomianu[872].X = 8.72
	punktyWykresuWielomianu[872].Y = -6_873.1

	punktyWykresuWielomianu[873].X = 8.73
	punktyWykresuWielomianu[873].Y = -6_888.668

	punktyWykresuWielomianu[874].X = 8.74
	punktyWykresuWielomianu[874].Y = -6_904.243

	punktyWykresuWielomianu[875].X = 8.75
	punktyWykresuWielomianu[875].Y = -6_919.824

	punktyWykresuWielomianu[876].X = 8.76
	punktyWykresuWielomianu[876].Y = -6_935.411

	punktyWykresuWielomianu[877].X = 8.77
	punktyWykresuWielomianu[877].Y = -6_951.005

	punktyWykresuWielomianu[878].X = 8.78
	punktyWykresuWielomianu[878].Y = -6_966.605

	punktyWykresuWielomianu[879].X = 8.79
	punktyWykresuWielomianu[879].Y = -6_982.211

	punktyWykresuWielomianu[880].X = 8.8
	punktyWykresuWielomianu[880].Y = -6_997.823

	punktyWykresuWielomianu[881].X = 8.81
	punktyWykresuWielomianu[881].Y = -7_013.442

	punktyWykresuWielomianu[882].X = 8.82
	punktyWykresuWielomianu[882].Y = -7_029.065

	punktyWykresuWielomianu[883].X = 8.83
	punktyWykresuWielomianu[883].Y = -7_044.695

	punktyWykresuWielomianu[884].X = 8.84
	punktyWykresuWielomianu[884].Y = -7_060.33

	punktyWykresuWielomianu[885].X = 8.85
	punktyWykresuWielomianu[885].Y = -7_075.971

	punktyWykresuWielomianu[886].X = 8.86
	punktyWykresuWielomianu[886].Y = -7_091.617

	punktyWykresuWielomianu[887].X = 8.87
	punktyWykresuWielomianu[887].Y = -7_107.269

	punktyWykresuWielomianu[888].X = 8.88
	punktyWykresuWielomianu[888].Y = -7_122.926

	punktyWykresuWielomianu[889].X = 8.89
	punktyWykresuWielomianu[889].Y = -7_138.588

	punktyWykresuWielomianu[890].X = 8.9
	punktyWykresuWielomianu[890].Y = -7_154.255

	punktyWykresuWielomianu[891].X = 8.91
	punktyWykresuWielomianu[891].Y = -7_169.928

	punktyWykresuWielomianu[892].X = 8.92
	punktyWykresuWielomianu[892].Y = -7_185.605

	punktyWykresuWielomianu[893].X = 8.93
	punktyWykresuWielomianu[893].Y = -7_201.287

	punktyWykresuWielomianu[894].X = 8.94
	punktyWykresuWielomianu[894].Y = -7_216.973

	punktyWykresuWielomianu[895].X = 8.95
	punktyWykresuWielomianu[895].Y = -7_232.665

	punktyWykresuWielomianu[896].X = 8.96
	punktyWykresuWielomianu[896].Y = -7_248.36

	punktyWykresuWielomianu[897].X = 8.97
	punktyWykresuWielomianu[897].Y = -7_264.061

	punktyWykresuWielomianu[898].X = 8.98
	punktyWykresuWielomianu[898].Y = -7_279.765

	punktyWykresuWielomianu[899].X = 8.99
	punktyWykresuWielomianu[899].Y = -7_295.474

	punktyWykresuWielomianu[900].X = 9.0
	punktyWykresuWielomianu[900].Y = -7_311.187

	punktyWykresuWielomianu[901].X = 9.01
	punktyWykresuWielomianu[901].Y = -7_326.904

	punktyWykresuWielomianu[902].X = 9.02
	punktyWykresuWielomianu[902].Y = -7_342.625

	punktyWykresuWielomianu[903].X = 9.03
	punktyWykresuWielomianu[903].Y = -7_358.35

	punktyWykresuWielomianu[904].X = 9.04
	punktyWykresuWielomianu[904].Y = -7_374.078

	punktyWykresuWielomianu[905].X = 9.05
	punktyWykresuWielomianu[905].Y = -7_389.811

	punktyWykresuWielomianu[906].X = 9.06
	punktyWykresuWielomianu[906].Y = -7_405.546

	punktyWykresuWielomianu[907].X = 9.07
	punktyWykresuWielomianu[907].Y = -7_421.286

	punktyWykresuWielomianu[908].X = 9.08
	punktyWykresuWielomianu[908].Y = -7_437.028

	punktyWykresuWielomianu[909].X = 9.09
	punktyWykresuWielomianu[909].Y = -7_452.774

	punktyWykresuWielomianu[910].X = 9.1
	punktyWykresuWielomianu[910].Y = -7_468.523

	punktyWykresuWielomianu[911].X = 9.11
	punktyWykresuWielomianu[911].Y = -7_484.276

	punktyWykresuWielomianu[912].X = 9.12
	punktyWykresuWielomianu[912].Y = -7_500.031

	punktyWykresuWielomianu[913].X = 9.13
	punktyWykresuWielomianu[913].Y = -7_515.789

	punktyWykresuWielomianu[914].X = 9.14
	punktyWykresuWielomianu[914].Y = -7_531.55

	punktyWykresuWielomianu[915].X = 9.15
	punktyWykresuWielomianu[915].Y = -7_547.313

	punktyWykresuWielomianu[916].X = 9.16
	punktyWykresuWielomianu[916].Y = -7_563.079

	punktyWykresuWielomianu[917].X = 9.17
	punktyWykresuWielomianu[917].Y = -7_578.848

	punktyWykresuWielomianu[918].X = 9.18
	punktyWykresuWielomianu[918].Y = -7_594.619

	punktyWykresuWielomianu[919].X = 9.19
	punktyWykresuWielomianu[919].Y = -7_610.392

	punktyWykresuWielomianu[920].X = 9.2
	punktyWykresuWielomianu[920].Y = -7_626.167

	punktyWykresuWielomianu[921].X = 9.21
	punktyWykresuWielomianu[921].Y = -7_641.945

	punktyWykresuWielomianu[922].X = 9.22
	punktyWykresuWielomianu[922].Y = -7_657.724

	punktyWykresuWielomianu[923].X = 9.23
	punktyWykresuWielomianu[923].Y = -7_673.506

	punktyWykresuWielomianu[924].X = 9.24
	punktyWykresuWielomianu[924].Y = -7_689.289

	punktyWykresuWielomianu[925].X = 9.25
	punktyWykresuWielomianu[925].Y = -7_705.074

	punktyWykresuWielomianu[926].X = 9.26
	punktyWykresuWielomianu[926].Y = -7_720.86

	punktyWykresuWielomianu[927].X = 9.27
	punktyWykresuWielomianu[927].Y = -7_736.648

	punktyWykresuWielomianu[928].X = 9.28
	punktyWykresuWielomianu[928].Y = -7_752.437

	punktyWykresuWielomianu[929].X = 9.29
	punktyWykresuWielomianu[929].Y = -7_768.228

	punktyWykresuWielomianu[930].X = 9.3
	punktyWykresuWielomianu[930].Y = -7_784.019

	punktyWykresuWielomianu[931].X = 9.31
	punktyWykresuWielomianu[931].Y = -7_799.812

	punktyWykresuWielomianu[932].X = 9.32
	punktyWykresuWielomianu[932].Y = -7_815.606

	punktyWykresuWielomianu[933].X = 9.33
	punktyWykresuWielomianu[933].Y = -7_831.4

	punktyWykresuWielomianu[934].X = 9.34
	punktyWykresuWielomianu[934].Y = -7_847.196

	punktyWykresuWielomianu[935].X = 9.35
	punktyWykresuWielomianu[935].Y = -7_862.992

	punktyWykresuWielomianu[936].X = 9.36
	punktyWykresuWielomianu[936].Y = -7_878.788

	punktyWykresuWielomianu[937].X = 9.37
	punktyWykresuWielomianu[937].Y = -7_894.585

	punktyWykresuWielomianu[938].X = 9.38
	punktyWykresuWielomianu[938].Y = -7_910.382

	punktyWykresuWielomianu[939].X = 9.39
	punktyWykresuWielomianu[939].Y = -7_926.18

	punktyWykresuWielomianu[940].X = 9.4
	punktyWykresuWielomianu[940].Y = -7_941.977

	punktyWykresuWielomianu[941].X = 9.41
	punktyWykresuWielomianu[941].Y = -7_957.775

	punktyWykresuWielomianu[942].X = 9.42
	punktyWykresuWielomianu[942].Y = -7_973.573

	punktyWykresuWielomianu[943].X = 9.43
	punktyWykresuWielomianu[943].Y = -7_989.37

	punktyWykresuWielomianu[944].X = 9.44
	punktyWykresuWielomianu[944].Y = -8_005.167

	punktyWykresuWielomianu[945].X = 9.45
	punktyWykresuWielomianu[945].Y = -8_020.964

	punktyWykresuWielomianu[946].X = 9.46
	punktyWykresuWielomianu[946].Y = -8_036.76

	punktyWykresuWielomianu[947].X = 9.47
	punktyWykresuWielomianu[947].Y = -8_052.555

	punktyWykresuWielomianu[948].X = 9.48
	punktyWykresuWielomianu[948].Y = -8_068.35

	punktyWykresuWielomianu[949].X = 9.49
	punktyWykresuWielomianu[949].Y = -8_084.144

	punktyWykresuWielomianu[950].X = 9.5
	punktyWykresuWielomianu[950].Y = -8_099.937

	punktyWykresuWielomianu[951].X = 9.51
	punktyWykresuWielomianu[951].Y = -8_115.729

	punktyWykresuWielomianu[952].X = 9.52
	punktyWykresuWielomianu[952].Y = -8_131.52

	punktyWykresuWielomianu[953].X = 9.53
	punktyWykresuWielomianu[953].Y = -8_147.309

	punktyWykresuWielomianu[954].X = 9.54
	punktyWykresuWielomianu[954].Y = -8_163.097

	punktyWykresuWielomianu[955].X = 9.55
	punktyWykresuWielomianu[955].Y = -8_178.884

	punktyWykresuWielomianu[956].X = 9.56
	punktyWykresuWielomianu[956].Y = -8_194.669

	punktyWykresuWielomianu[957].X = 9.57
	punktyWykresuWielomianu[957].Y = -8_210.453

	punktyWykresuWielomianu[958].X = 9.58
	punktyWykresuWielomianu[958].Y = -8_226.234

	punktyWykresuWielomianu[959].X = 9.59
	punktyWykresuWielomianu[959].Y = -8_242.014

	punktyWykresuWielomianu[960].X = 9.6
	punktyWykresuWielomianu[960].Y = -8_257.791

	punktyWykresuWielomianu[961].X = 9.61
	punktyWykresuWielomianu[961].Y = -8_273.34

	punktyWykresuWielomianu[962].X = 9.62
	punktyWykresuWielomianu[962].Y = -8_289.34

	punktyWykresuWielomianu[963].X = 9.63
	punktyWykresuWielomianu[963].Y = -8_305.111

	punktyWykresuWielomianu[964].X = 9.64
	punktyWykresuWielomianu[964].Y = -8_320.879

	punktyWykresuWielomianu[965].X = 9.65
	punktyWykresuWielomianu[965].Y = -8_336.645

	punktyWykresuWielomianu[966].X = 9.66
	punktyWykresuWielomianu[966].Y = -8_352.408

	punktyWykresuWielomianu[967].X = 9.67
	punktyWykresuWielomianu[967].Y = -8_368.169

	punktyWykresuWielomianu[968].X = 9.68
	punktyWykresuWielomianu[968].Y = -8_383.926

	punktyWykresuWielomianu[969].X = 9.69
	punktyWykresuWielomianu[969].Y = -8_399.68

	punktyWykresuWielomianu[970].X = 9.7
	punktyWykresuWielomianu[970].Y = -8_415.431

	punktyWykresuWielomianu[971].X = 9.71
	punktyWykresuWielomianu[971].Y = -8_431.179

	punktyWykresuWielomianu[972].X = 9.72
	punktyWykresuWielomianu[972].Y = -8_446.924

	punktyWykresuWielomianu[973].X = 9.73
	punktyWykresuWielomianu[973].Y = -8_462.665

	punktyWykresuWielomianu[974].X = 9.74
	punktyWykresuWielomianu[974].Y = -8_478.403

	punktyWykresuWielomianu[975].X = 9.75
	punktyWykresuWielomianu[975].Y = -8_494.136

	punktyWykresuWielomianu[976].X = 9.76
	punktyWykresuWielomianu[976].Y = -8_509.866

	punktyWykresuWielomianu[977].X = 9.77
	punktyWykresuWielomianu[977].Y = -8_525.592

	punktyWykresuWielomianu[978].X = 9.78
	punktyWykresuWielomianu[978].Y = -8_541.314

	punktyWykresuWielomianu[979].X = 9.79
	punktyWykresuWielomianu[979].Y = -8_557.032

	punktyWykresuWielomianu[980].X = 9.8
	punktyWykresuWielomianu[980].Y = -8_572.745

	punktyWykresuWielomianu[981].X = 9.81
	punktyWykresuWielomianu[981].Y = -8_588.455

	punktyWykresuWielomianu[982].X = 9.82
	punktyWykresuWielomianu[982].Y = -8_604.159

	punktyWykresuWielomianu[983].X = 9.83
	punktyWykresuWielomianu[983].Y = -8_619.859

	punktyWykresuWielomianu[984].X = 9.84
	punktyWykresuWielomianu[984].Y = -8_635.554

	punktyWykresuWielomianu[985].X = 9.85
	punktyWykresuWielomianu[985].Y = -8_651.245

	punktyWykresuWielomianu[986].X = 9.86
	punktyWykresuWielomianu[986].Y = -8_666.93

	punktyWykresuWielomianu[987].X = 9.87
	punktyWykresuWielomianu[987].Y = -8_682.61

	punktyWykresuWielomianu[988].X = 9.88
	punktyWykresuWielomianu[988].Y = -8_698.285

	punktyWykresuWielomianu[989].X = 9.89
	punktyWykresuWielomianu[989].Y = -8_713.955

	punktyWykresuWielomianu[990].X = 9.9
	punktyWykresuWielomianu[990].Y = -8_729.619

	punktyWykresuWielomianu[991].X = 9.91
	punktyWykresuWielomianu[991].Y = -8_745.278

	punktyWykresuWielomianu[992].X = 9.92
	punktyWykresuWielomianu[992].Y = -8_760.931

	punktyWykresuWielomianu[993].X = 9.93
	punktyWykresuWielomianu[993].Y = -8_776.579

	punktyWykresuWielomianu[994].X = 9.94
	punktyWykresuWielomianu[994].Y = -8_792.22

	punktyWykresuWielomianu[995].X = 9.95
	punktyWykresuWielomianu[995].Y = -8_807.855

	punktyWykresuWielomianu[996].X = 9.96
	punktyWykresuWielomianu[996].Y = -8_823.484

	punktyWykresuWielomianu[997].X = 9.97
	punktyWykresuWielomianu[997].Y = -8_839.107

	punktyWykresuWielomianu[998].X = 9.98
	punktyWykresuWielomianu[998].Y = -8_854.724

	punktyWykresuWielomianu[999].X = 9.99
	punktyWykresuWielomianu[999].Y = -8_870.334

	punktyWykresuWielomianu[1_000].X = 10.0
	punktyWykresuWielomianu[1_000].Y = -8_885.937










	wykresWielomianu := plot.New()

	wykresWielomianu.Title.Text = "Wykres funkcji f(x) = x^4 - 20x^3 + 33.75x^2 - 235x + 89.0625"

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
		"Wykres-wielomianu-03.png"); err != nil {

		panic(err)
	}
}
