package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres wielomianu f(x) = x^5 - 19.222 x^4 + 69.903 x^3
	// - 10.34 x^2 - 15.722 x + 9.177

	punktyWykresuWielomianu := make(plotter.XYs, 801)

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
	punktyWykresuWielomianu[120].Y = 58.82

	// punktyWykresuWielomianu[121].X = 1.21
	// punktyWykresuWielomianu[121].Y = -179.161

	// punktyWykresuWielomianu[122].X = 1.22
	// punktyWykresuWielomianu[122].Y = -181.505

	// punktyWykresuWielomianu[123].X = 1.23
	// punktyWykresuWielomianu[123].Y = -183.855

	// punktyWykresuWielomianu[124].X = 1.24
	// punktyWykresuWielomianu[124].Y = -186.211

	// punktyWykresuWielomianu[125].X = 1.25
	// punktyWykresuWielomianu[125].Y = -188.574

	// punktyWykresuWielomianu[126].X = 1.26
	// punktyWykresuWielomianu[126].Y = -190.943

	// punktyWykresuWielomianu[127].X = 1.27
	// punktyWykresuWielomianu[127].Y = -193.318

	// punktyWykresuWielomianu[128].X = 1.28
	// punktyWykresuWielomianu[128].Y = -195.7

	// punktyWykresuWielomianu[129].X = 1.29
	// punktyWykresuWielomianu[129].Y = -198.088

	// punktyWykresuWielomianu[130].X = 1.3
	// punktyWykresuWielomianu[130].Y = -200.483

	// punktyWykresuWielomianu[131].X = 1.31
	// punktyWykresuWielomianu[131].Y = -202.885

	// punktyWykresuWielomianu[132].X = 1.32
	// punktyWykresuWielomianu[132].Y = -205.294

	// punktyWykresuWielomianu[133].X = 1.33
	// punktyWykresuWielomianu[133].Y = -207.71

	// punktyWykresuWielomianu[134].X = 1.34
	// punktyWykresuWielomianu[134].Y = -210.133

	// punktyWykresuWielomianu[135].X = 1.35
	// punktyWykresuWielomianu[135].Y = -212.564

	// punktyWykresuWielomianu[136].X = 1.36
	// punktyWykresuWielomianu[136].Y = -215.001

	// punktyWykresuWielomianu[137].X = 1.37
	// punktyWykresuWielomianu[137].Y = -217.446

	// punktyWykresuWielomianu[138].X = 1.38
	// punktyWykresuWielomianu[138].Y = -219.898

	// punktyWykresuWielomianu[139].X = 1.39
	// punktyWykresuWielomianu[139].Y = -222.358

	// punktyWykresuWielomianu[140].X = 1.4
	// punktyWykresuWielomianu[140].Y = -224.825

	// punktyWykresuWielomianu[141].X = 1.41
	// punktyWykresuWielomianu[141].Y = -227.301

	// punktyWykresuWielomianu[142].X = 1.42
	// punktyWykresuWielomianu[142].Y = -229.783

	// punktyWykresuWielomianu[143].X = 1.43
	// punktyWykresuWielomianu[143].Y = -232.274

	// punktyWykresuWielomianu[144].X = 1.44
	// punktyWykresuWielomianu[144].Y = -234.773

	// punktyWykresuWielomianu[145].X = 1.45
	// punktyWykresuWielomianu[145].Y = -237.28

	// punktyWykresuWielomianu[146].X = 1.46
	// punktyWykresuWielomianu[146].Y = -239.795

	// punktyWykresuWielomianu[147].X = 1.47
	// punktyWykresuWielomianu[147].Y = -242.318

	// punktyWykresuWielomianu[148].X = 1.48
	// punktyWykresuWielomianu[148].Y = -244.849

	// punktyWykresuWielomianu[149].X = 1.49
	// punktyWykresuWielomianu[149].Y = -247.389

	// punktyWykresuWielomianu[150].X = 1.5
	// punktyWykresuWielomianu[150].Y = -249.937

	// punktyWykresuWielomianu[151].X = 1.51
	// punktyWykresuWielomianu[151].Y = -252.494

	// punktyWykresuWielomianu[152].X = 1.52
	// punktyWykresuWielomianu[152].Y = -255.059

	// punktyWykresuWielomianu[153].X = 1.53
	// punktyWykresuWielomianu[153].Y = -257.633

	// punktyWykresuWielomianu[154].X = 1.54
	// punktyWykresuWielomianu[154].Y = -260.216

	// punktyWykresuWielomianu[155].X = 1.55
	// punktyWykresuWielomianu[155].Y = -262.808

	// punktyWykresuWielomianu[156].X = 1.56
	// punktyWykresuWielomianu[156].Y = -265.409

	// punktyWykresuWielomianu[157].X = 1.57
	// punktyWykresuWielomianu[157].Y = -268.019

	// punktyWykresuWielomianu[158].X = 1.58
	// punktyWykresuWielomianu[158].Y = -270.638

	// punktyWykresuWielomianu[159].X = 1.59
	// punktyWykresuWielomianu[159].Y = -273.903

	// punktyWykresuWielomianu[160].X = 1.6
	// punktyWykresuWielomianu[160].Y = -275.903

	// punktyWykresuWielomianu[161].X = 1.61
	// punktyWykresuWielomianu[161].Y = -278.55

	// punktyWykresuWielomianu[162].X = 1.62
	// punktyWykresuWielomianu[162].Y = -281.207

	// punktyWykresuWielomianu[163].X = 1.63
	// punktyWykresuWielomianu[163].Y = -283.872

	// punktyWykresuWielomianu[164].X = 1.64
	// punktyWykresuWielomianu[164].Y = -286.548

	// punktyWykresuWielomianu[165].X = 1.65
	// punktyWykresuWielomianu[165].Y = -289.233

	// punktyWykresuWielomianu[166].X = 1.66
	// punktyWykresuWielomianu[166].Y = -291.928

	// punktyWykresuWielomianu[167].X = 1.67
	// punktyWykresuWielomianu[167].Y = -294.633

	// punktyWykresuWielomianu[168].X = 1.68
	// punktyWykresuWielomianu[168].Y = -297.348

	// punktyWykresuWielomianu[169].X = 1.69
	// punktyWykresuWielomianu[169].Y = -300.073

	// punktyWykresuWielomianu[170].X = 1.7
	// punktyWykresuWielomianu[170].Y = -302.807

	// punktyWykresuWielomianu[171].X = 1.71
	// punktyWykresuWielomianu[171].Y = -305.553

	// punktyWykresuWielomianu[172].X = 1.72
	// punktyWykresuWielomianu[172].Y = -308.308

	// punktyWykresuWielomianu[173].X = 1.73
	// punktyWykresuWielomianu[173].Y = -311.074

	// punktyWykresuWielomianu[174].X = 1.74
	// punktyWykresuWielomianu[174].Y = -313.85

	// punktyWykresuWielomianu[175].X = 1.75
	// punktyWykresuWielomianu[175].Y = -316.636

	// punktyWykresuWielomianu[176].X = 1.76
	// punktyWykresuWielomianu[176].Y = -319.433

	// punktyWykresuWielomianu[177].X = 1.77
	// punktyWykresuWielomianu[177].Y = -322.241

	// punktyWykresuWielomianu[178].X = 1.78
	// punktyWykresuWielomianu[178].Y = -325.06

	// punktyWykresuWielomianu[179].X = 1.79
	// punktyWykresuWielomianu[179].Y = -327.889

	// punktyWykresuWielomianu[180].X = 1.8
	// punktyWykresuWielomianu[180].Y = -330.729

	// punktyWykresuWielomianu[181].X = 1.81
	// punktyWykresuWielomianu[181].Y = -333.581

	// punktyWykresuWielomianu[182].X = 1.82
	// punktyWykresuWielomianu[182].Y = -336.443

	// punktyWykresuWielomianu[183].X = 1.83
	// punktyWykresuWielomianu[183].Y = -339.316

	// punktyWykresuWielomianu[184].X = 1.84
	// punktyWykresuWielomianu[184].Y = -342.201

	// punktyWykresuWielomianu[185].X = 1.85
	// punktyWykresuWielomianu[185].Y = -345.097

	// punktyWykresuWielomianu[186].X = 1.86
	// punktyWykresuWielomianu[186].Y = -348.004

	// punktyWykresuWielomianu[187].X = 1.87
	// punktyWykresuWielomianu[187].Y = -350.922

	// punktyWykresuWielomianu[188].X = 1.88
	// punktyWykresuWielomianu[188].Y = -353.853

	// punktyWykresuWielomianu[189].X = 1.89
	// punktyWykresuWielomianu[189].Y = -356.794

	// punktyWykresuWielomianu[190].X = 1.9
	// punktyWykresuWielomianu[190].Y = -359.747

	// punktyWykresuWielomianu[191].X = 1.91
	// punktyWykresuWielomianu[191].Y = -362.712

	// punktyWykresuWielomianu[192].X = 1.92
	// punktyWykresuWielomianu[192].Y = -365.689

	// punktyWykresuWielomianu[193].X = 1.93
	// punktyWykresuWielomianu[193].Y = -368.678

	// punktyWykresuWielomianu[194].X = 1.94
	// punktyWykresuWielomianu[194].Y = -371.679

	// punktyWykresuWielomianu[195].X = 1.95
	// punktyWykresuWielomianu[195].Y = -374.691

	// punktyWykresuWielomianu[196].X = 1.96
	// punktyWykresuWielomianu[196].Y = -377.716

	// punktyWykresuWielomianu[197].X = 1.97
	// punktyWykresuWielomianu[197].Y = -380.753

	// punktyWykresuWielomianu[198].X = 1.98
	// punktyWykresuWielomianu[198].Y = -383.802

	// punktyWykresuWielomianu[199].X = 1.99
	// punktyWykresuWielomianu[199].Y = -386.863

	// punktyWykresuWielomianu[200].X = 2.0
	// punktyWykresuWielomianu[200].Y = -389.937

	// punktyWykresuWielomianu[201].X = 2.01
	// punktyWykresuWielomianu[201].Y = -393.023

	// punktyWykresuWielomianu[202].X = 2.02
	// punktyWykresuWielomianu[202].Y = -396.122

	// punktyWykresuWielomianu[203].X = 2.03
	// punktyWykresuWielomianu[203].Y = -399.233

	// punktyWykresuWielomianu[204].X = 2.04
	// punktyWykresuWielomianu[204].Y = -402.357

	// punktyWykresuWielomianu[205].X = 2.05
	// punktyWykresuWielomianu[205].Y = -405.494

	// punktyWykresuWielomianu[206].X = 2.06
	// punktyWykresuWielomianu[206].Y = -408.644

	// punktyWykresuWielomianu[207].X = 2.07
	// punktyWykresuWielomianu[207].Y = -411.806

	// punktyWykresuWielomianu[208].X = 2.08
	// punktyWykresuWielomianu[208].Y = -414.982

	// punktyWykresuWielomianu[209].X = 2.09
	// punktyWykresuWielomianu[209].Y = -418.17

	// punktyWykresuWielomianu[210].X = 2.10
	// punktyWykresuWielomianu[210].Y = -421.371

	// punktyWykresuWielomianu[211].X = 2.11
	// punktyWykresuWielomianu[211].Y = -424.586

	// punktyWykresuWielomianu[212].X = 2.12
	// punktyWykresuWielomianu[212].Y = -427.814

	// punktyWykresuWielomianu[213].X = 2.13
	// punktyWykresuWielomianu[213].Y = -431.055

	// punktyWykresuWielomianu[214].X = 2.14
	// punktyWykresuWielomianu[214].Y = -434.31

	// punktyWykresuWielomianu[215].X = 2.15
	// punktyWykresuWielomianu[215].Y = -437.578

	// punktyWykresuWielomianu[216].X = 2.16
	// punktyWykresuWielomianu[216].Y = -440.859

	// punktyWykresuWielomianu[217].X = 2.17
	// punktyWykresuWielomianu[217].Y = -444.154

	// punktyWykresuWielomianu[218].X = 2.18
	// punktyWykresuWielomianu[218].Y = -447.463

	// punktyWykresuWielomianu[219].X = 2.19
	// punktyWykresuWielomianu[219].Y = -450.785

	// punktyWykresuWielomianu[220].X = 2.2
	// punktyWykresuWielomianu[220].Y = -454.121

	// punktyWykresuWielomianu[221].X = 2.21
	// punktyWykresuWielomianu[221].Y = -457.471

	// punktyWykresuWielomianu[222].X = 2.22
	// punktyWykresuWielomianu[222].Y = -460.835

	// punktyWykresuWielomianu[223].X = 2.23
	// punktyWykresuWielomianu[223].Y = -464.213

	// punktyWykresuWielomianu[224].X = 2.24
	// punktyWykresuWielomianu[224].Y = -467.605

	// punktyWykresuWielomianu[225].X = 2.25
	// punktyWykresuWielomianu[225].Y = -471.011

	// punktyWykresuWielomianu[226].X = 2.26
	// punktyWykresuWielomianu[226].Y = -474.431

	// punktyWykresuWielomianu[227].X = 2.27
	// punktyWykresuWielomianu[227].Y = -477.866

	// punktyWykresuWielomianu[228].X = 2.28
	// punktyWykresuWielomianu[228].Y = -481.315

	// punktyWykresuWielomianu[229].X = 2.29
	// punktyWykresuWielomianu[229].Y = -484.778

	// punktyWykresuWielomianu[230].X = 2.3
	// punktyWykresuWielomianu[230].Y = -488.255

	// punktyWykresuWielomianu[231].X = 2.31
	// punktyWykresuWielomianu[231].Y = -491.748

	// punktyWykresuWielomianu[232].X = 2.32
	// punktyWykresuWielomianu[232].Y = -495.254

	// punktyWykresuWielomianu[233].X = 2.33
	// punktyWykresuWielomianu[233].Y = -498.775

	// punktyWykresuWielomianu[234].X = 2.34
	// punktyWykresuWielomianu[234].Y = -502.311

	// punktyWykresuWielomianu[235].X = 2.35
	// punktyWykresuWielomianu[235].Y = -505.862

	// punktyWykresuWielomianu[236].X = 2.36
	// punktyWykresuWielomianu[236].Y = -509.428

	// punktyWykresuWielomianu[237].X = 2.37
	// punktyWykresuWielomianu[237].Y = -513.008

	// punktyWykresuWielomianu[238].X = 2.38
	// punktyWykresuWielomianu[238].Y = -516.604

	// punktyWykresuWielomianu[239].X = 2.39
	// punktyWykresuWielomianu[239].Y = -520.214

	// punktyWykresuWielomianu[240].X = 2.4
	// punktyWykresuWielomianu[240].Y = -523.839

	// punktyWykresuWielomianu[241].X = 2.41
	// punktyWykresuWielomianu[241].Y = -527.48

	// punktyWykresuWielomianu[242].X = 2.42
	// punktyWykresuWielomianu[242].Y = -531.136

	// punktyWykresuWielomianu[243].X = 2.43
	// punktyWykresuWielomianu[243].Y = -534.807

	// punktyWykresuWielomianu[244].X = 2.44
	// punktyWykresuWielomianu[244].Y = -538.493

	// punktyWykresuWielomianu[245].X = 2.45
	// punktyWykresuWielomianu[245].Y = -542.195

	// punktyWykresuWielomianu[246].X = 2.46
	// punktyWykresuWielomianu[246].Y = -545.912

	// punktyWykresuWielomianu[247].X = 2.47
	// punktyWykresuWielomianu[247].Y = -549.645

	// punktyWykresuWielomianu[248].X = 2.48
	// punktyWykresuWielomianu[248].Y = -553.393

	// punktyWykresuWielomianu[249].X = 2.49
	// punktyWykresuWielomianu[249].Y = -557.157

	// punktyWykresuWielomianu[250].X = 2.5
	// punktyWykresuWielomianu[250].Y = -560.937

	// punktyWykresuWielomianu[251].X = 2.51
	// punktyWykresuWielomianu[251].Y = -564.732

	// punktyWykresuWielomianu[252].X = 2.52
	// punktyWykresuWielomianu[252].Y = -568.544

	// punktyWykresuWielomianu[253].X = 2.53
	// punktyWykresuWielomianu[253].Y = -572.371

	// punktyWykresuWielomianu[254].X = 2.54
	// punktyWykresuWielomianu[254].Y = -576.214

	// punktyWykresuWielomianu[255].X = 2.55
	// punktyWykresuWielomianu[255].Y = -580.073

	// punktyWykresuWielomianu[256].X = 2.56
	// punktyWykresuWielomianu[256].Y = -583.948

	// punktyWykresuWielomianu[257].X = 2.57
	// punktyWykresuWielomianu[257].Y = -587.839

	// punktyWykresuWielomianu[258].X = 2.58
	// punktyWykresuWielomianu[258].Y = -591.746

	// punktyWykresuWielomianu[259].X = 2.59
	// punktyWykresuWielomianu[259].Y = -595.67

	// punktyWykresuWielomianu[260].X = 2.6
	// punktyWykresuWielomianu[260].Y = -599.609

	// punktyWykresuWielomianu[261].X = 2.61
	// punktyWykresuWielomianu[261].Y = -603.566

	// punktyWykresuWielomianu[262].X = 2.62
	// punktyWykresuWielomianu[262].Y = -607.538

	// punktyWykresuWielomianu[263].X = 2.63
	// punktyWykresuWielomianu[263].Y = -611.527

	// punktyWykresuWielomianu[264].X = 2.64
	// punktyWykresuWielomianu[264].Y = -615.533

	// punktyWykresuWielomianu[265].X = 2.65
	// punktyWykresuWielomianu[265].Y = -619.555

	// punktyWykresuWielomianu[266].X = 2.66
	// punktyWykresuWielomianu[266].Y = -623.593

	// punktyWykresuWielomianu[267].X = 2.67
	// punktyWykresuWielomianu[267].Y = -627.649

	// punktyWykresuWielomianu[268].X = 2.68
	// punktyWykresuWielomianu[268].Y = -631.721

	// punktyWykresuWielomianu[269].X = 2.69
	// punktyWykresuWielomianu[269].Y = -635.81

	// punktyWykresuWielomianu[270].X = 2.7
	// punktyWykresuWielomianu[270].Y = -639.915

	// punktyWykresuWielomianu[271].X = 2.71
	// punktyWykresuWielomianu[271].Y = -644.038

	// punktyWykresuWielomianu[272].X = 2.72
	// punktyWykresuWielomianu[272].Y = -648.178

	// punktyWykresuWielomianu[273].X = 2.73
	// punktyWykresuWielomianu[273].Y = -652.334

	// punktyWykresuWielomianu[274].X = 2.74
	// punktyWykresuWielomianu[274].Y = -656.508

	// punktyWykresuWielomianu[275].X = 2.75
	// punktyWykresuWielomianu[275].Y = -660.699

	// punktyWykresuWielomianu[276].X = 2.76
	// punktyWykresuWielomianu[276].Y = -664.907

	// punktyWykresuWielomianu[277].X = 2.77
	// punktyWykresuWielomianu[277].Y = -669.132

	// punktyWykresuWielomianu[278].X = 2.78
	// punktyWykresuWielomianu[278].Y = -673.374

	// punktyWykresuWielomianu[279].X = 2.79
	// punktyWykresuWielomianu[279].Y = -677.634

	// punktyWykresuWielomianu[280].X = 2.8
	// punktyWykresuWielomianu[280].Y = -681.911

	// punktyWykresuWielomianu[281].X = 2.81
	// punktyWykresuWielomianu[281].Y = -686.206

	// punktyWykresuWielomianu[282].X = 2.82
	// punktyWykresuWielomianu[282].Y = -690.518

	// punktyWykresuWielomianu[283].X = 2.83
	// punktyWykresuWielomianu[283].Y = -694.848

	// punktyWykresuWielomianu[284].X = 2.84
	// punktyWykresuWielomianu[284].Y = -699.195

	// punktyWykresuWielomianu[285].X = 2.85
	// punktyWykresuWielomianu[285].Y = -703.56

	// punktyWykresuWielomianu[286].X = 2.86
	// punktyWykresuWielomianu[286].Y = -707.943

	// punktyWykresuWielomianu[287].X = 2.87
	// punktyWykresuWielomianu[287].Y = -712.343

	// punktyWykresuWielomianu[288].X = 2.88
	// punktyWykresuWielomianu[288].Y = -716.761

	// punktyWykresuWielomianu[289].X = 2.89
	// punktyWykresuWielomianu[289].Y = -721.197

	// punktyWykresuWielomianu[290].X = 2.9
	// punktyWykresuWielomianu[290].Y = -725.651

	// punktyWykresuWielomianu[291].X = 2.91
	// punktyWykresuWielomianu[291].Y = -730.123

	// punktyWykresuWielomianu[292].X = 2.92
	// punktyWykresuWielomianu[292].Y = -734.613

	// punktyWykresuWielomianu[293].X = 2.93
	// punktyWykresuWielomianu[293].Y = -739.121

	// punktyWykresuWielomianu[294].X = 2.94
	// punktyWykresuWielomianu[294].Y = -743.647

	// punktyWykresuWielomianu[295].X = 2.95
	// punktyWykresuWielomianu[295].Y = -748.192

	// punktyWykresuWielomianu[296].X = 2.96
	// punktyWykresuWielomianu[296].Y = -752.754

	// punktyWykresuWielomianu[297].X = 2.97
	// punktyWykresuWielomianu[297].Y = -757.335

	// punktyWykresuWielomianu[298].X = 2.98
	// punktyWykresuWielomianu[298].Y = -761.934

	// punktyWykresuWielomianu[299].X = 2.99
	// punktyWykresuWielomianu[299].Y = -766.551

	// punktyWykresuWielomianu[300].X = 3.0
	// punktyWykresuWielomianu[300].Y = -771.187

	// punktyWykresuWielomianu[301].X = 3.01
	// punktyWykresuWielomianu[301].Y = -775.841

	// punktyWykresuWielomianu[302].X = 3.02
	// punktyWykresuWielomianu[302].Y = -780.514

	// punktyWykresuWielomianu[303].X = 3.03
	// punktyWykresuWielomianu[303].Y = -785.205

	// punktyWykresuWielomianu[304].X = 3.04
	// punktyWykresuWielomianu[304].Y = -789.915

	// punktyWykresuWielomianu[305].X = 3.05
	// punktyWykresuWielomianu[305].Y = -794.644

	// punktyWykresuWielomianu[306].X = 3.06
	// punktyWykresuWielomianu[306].Y = -799.391

	// punktyWykresuWielomianu[307].X = 3.07
	// punktyWykresuWielomianu[307].Y = -804.157

	// punktyWykresuWielomianu[308].X = 3.08
	// punktyWykresuWielomianu[308].Y = -808.942

	// punktyWykresuWielomianu[309].X = 3.09
	// punktyWykresuWielomianu[309].Y = -813.745

	// punktyWykresuWielomianu[310].X = 3.1
	// punktyWykresuWielomianu[310].Y = -818.745

	// punktyWykresuWielomianu[311].X = 3.11
	// punktyWykresuWielomianu[311].Y = -823.409

	// punktyWykresuWielomianu[312].X = 3.12
	// punktyWykresuWielomianu[312].Y = -828.269

	// punktyWykresuWielomianu[313].X = 3.13
	// punktyWykresuWielomianu[313].Y = -833.148

	// punktyWykresuWielomianu[314].X = 3.14
	// punktyWykresuWielomianu[314].Y = -838.047

	// punktyWykresuWielomianu[315].X = 3.15
	// punktyWykresuWielomianu[315].Y = -842.964

	// punktyWykresuWielomianu[316].X = 3.16
	// punktyWykresuWielomianu[316].Y = -847.901

	// punktyWykresuWielomianu[317].X = 3.17
	// punktyWykresuWielomianu[317].Y = -852.857

	// punktyWykresuWielomianu[318].X = 3.18
	// punktyWykresuWielomianu[318].Y = -857.832

	// punktyWykresuWielomianu[319].X = 3.19
	// punktyWykresuWielomianu[319].Y = -862.826

	// punktyWykresuWielomianu[320].X = 3.2
	// punktyWykresuWielomianu[320].Y = -867.839

	// punktyWykresuWielomianu[321].X = 3.21
	// punktyWykresuWielomianu[321].Y = -872.872

	// punktyWykresuWielomianu[322].X = 3.22
	// punktyWykresuWielomianu[322].Y = -877.925

	// punktyWykresuWielomianu[323].X = 3.23
	// punktyWykresuWielomianu[323].Y = -882.997

	// punktyWykresuWielomianu[324].X = 3.24
	// punktyWykresuWielomianu[324].Y = -888.088

	// punktyWykresuWielomianu[325].X = 3.25
	// punktyWykresuWielomianu[325].Y = -893.329

	// punktyWykresuWielomianu[326].X = 3.26
	// punktyWykresuWielomianu[326].Y = -898.329

	// punktyWykresuWielomianu[327].X = 3.27
	// punktyWykresuWielomianu[327].Y = -903.479

	// punktyWykresuWielomianu[328].X = 3.28
	// punktyWykresuWielomianu[328].Y = -908.649

	// punktyWykresuWielomianu[329].X = 3.29
	// punktyWykresuWielomianu[329].Y = -913.838

	// punktyWykresuWielomianu[330].X = 3.3
	// punktyWykresuWielomianu[330].Y = -919.047

	// punktyWykresuWielomianu[331].X = 3.31
	// punktyWykresuWielomianu[331].Y = -924.276

	// punktyWykresuWielomianu[332].X = 3.32
	// punktyWykresuWielomianu[332].Y = -929.525

	// punktyWykresuWielomianu[333].X = 3.33
	// punktyWykresuWielomianu[333].Y = -934.794

	// punktyWykresuWielomianu[334].X = 3.34
	// punktyWykresuWielomianu[334].Y = -940.082

	// punktyWykresuWielomianu[335].X = 3.35
	// punktyWykresuWielomianu[335].Y = -945.391

	// punktyWykresuWielomianu[336].X = 3.36
	// punktyWykresuWielomianu[336].Y = -950.719

	// punktyWykresuWielomianu[337].X = 3.37
	// punktyWykresuWielomianu[337].Y = -956.068

	// punktyWykresuWielomianu[338].X = 3.38
	// punktyWykresuWielomianu[338].Y = -961.436

	// punktyWykresuWielomianu[339].X = 3.39
	// punktyWykresuWielomianu[339].Y = -966.825

	// punktyWykresuWielomianu[340].X = 3.4
	// punktyWykresuWielomianu[340].Y = -972.233

	// punktyWykresuWielomianu[341].X = 3.41
	// punktyWykresuWielomianu[341].Y = -977.662

	// punktyWykresuWielomianu[342].X = 3.42
	// punktyWykresuWielomianu[342].Y = -983.112

	// punktyWykresuWielomianu[343].X = 3.43
	// punktyWykresuWielomianu[343].Y = -988.581

	// punktyWykresuWielomianu[344].X = 3.44
	// punktyWykresuWielomianu[344].Y = -994.071

	// punktyWykresuWielomianu[345].X = 3.45
	// punktyWykresuWielomianu[345].Y = -999.581

	// punktyWykresuWielomianu[346].X = 3.46
	// punktyWykresuWielomianu[346].Y = -1005.111

	// punktyWykresuWielomianu[347].X = 3.47
	// punktyWykresuWielomianu[347].Y = -1010.662

	// punktyWykresuWielomianu[348].X = 3.48
	// punktyWykresuWielomianu[348].Y = -1016.233

	// punktyWykresuWielomianu[349].X = 3.49
	// punktyWykresuWielomianu[349].Y = -1021.825

	// punktyWykresuWielomianu[350].X = 3.5
	// punktyWykresuWielomianu[350].Y = -1027.437

	// punktyWykresuWielomianu[351].X = 3.51
	// punktyWykresuWielomianu[351].Y = -1033.07

	// punktyWykresuWielomianu[352].X = 3.52
	// punktyWykresuWielomianu[352].Y = -1038.723

	// punktyWykresuWielomianu[353].X = 3.53
	// punktyWykresuWielomianu[353].Y = -1044.397

	// punktyWykresuWielomianu[354].X = 3.54
	// punktyWykresuWielomianu[354].Y = -1050.092

	// punktyWykresuWielomianu[355].X = 3.55
	// punktyWykresuWielomianu[355].Y = -1055.807

	// punktyWykresuWielomianu[356].X = 3.56
	// punktyWykresuWielomianu[356].Y = -1061.543

	// punktyWykresuWielomianu[357].X = 3.57
	// punktyWykresuWielomianu[357].Y = -1067.3

	// punktyWykresuWielomianu[358].X = 3.58
	// punktyWykresuWielomianu[358].Y = -1073.078

	// punktyWykresuWielomianu[359].X = 3.59
	// punktyWykresuWielomianu[359].Y = -1078.876

	// punktyWykresuWielomianu[360].X = 3.6
	// punktyWykresuWielomianu[360].Y = -1084.695

	// punktyWykresuWielomianu[361].X = 3.61
	// punktyWykresuWielomianu[361].Y = -1090.536

	// punktyWykresuWielomianu[362].X = 3.62
	// punktyWykresuWielomianu[362].Y = -1096.397

	// punktyWykresuWielomianu[363].X = 3.63
	// punktyWykresuWielomianu[363].Y = -1102.279

	// punktyWykresuWielomianu[364].X = 3.64
	// punktyWykresuWielomianu[364].Y = -1108.182

	// punktyWykresuWielomianu[365].X = 3.65
	// punktyWykresuWielomianu[365].Y = -1114.106

	// punktyWykresuWielomianu[366].X = 3.66
	// punktyWykresuWielomianu[366].Y = -1120.051

	// punktyWykresuWielomianu[367].X = 3.67
	// punktyWykresuWielomianu[367].Y = -1126.018

	// punktyWykresuWielomianu[368].X = 3.68
	// punktyWykresuWielomianu[368].Y = -1132.005

	// punktyWykresuWielomianu[369].X = 3.69
	// punktyWykresuWielomianu[369].Y = -1138.014

	// punktyWykresuWielomianu[370].X = 3.7
	// punktyWykresuWielomianu[370].Y = -1144.043

	// punktyWykresuWielomianu[371].X = 3.71
	// punktyWykresuWielomianu[371].Y = -1150.094

	// punktyWykresuWielomianu[372].X = 3.72
	// punktyWykresuWielomianu[372].Y = -1156.167

	// punktyWykresuWielomianu[373].X = 3.73
	// punktyWykresuWielomianu[373].Y = -1162.26

	// punktyWykresuWielomianu[374].X = 3.74
	// punktyWykresuWielomianu[374].Y = -1168.375

	// punktyWykresuWielomianu[375].X = 3.75
	// punktyWykresuWielomianu[375].Y = -1174.511

	// punktyWykresuWielomianu[376].X = 3.76
	// punktyWykresuWielomianu[376].Y = -1180.669

	// punktyWykresuWielomianu[377].X = 3.77
	// punktyWykresuWielomianu[377].Y = -1186.848

	// punktyWykresuWielomianu[378].X = 3.78
	// punktyWykresuWielomianu[378].Y = -1193.048

	// punktyWykresuWielomianu[379].X = 3.79
	// punktyWykresuWielomianu[379].Y = -1199.27

	// punktyWykresuWielomianu[380].X = 3.8
	// punktyWykresuWielomianu[380].Y = -1205.513

	// punktyWykresuWielomianu[381].X = 3.81
	// punktyWykresuWielomianu[381].Y = -1211.778

	// punktyWykresuWielomianu[382].X = 3.82
	// punktyWykresuWielomianu[382].Y = -1218.065

	// punktyWykresuWielomianu[383].X = 3.83
	// punktyWykresuWielomianu[383].Y = -1224.373

	// punktyWykresuWielomianu[384].X = 3.84
	// punktyWykresuWielomianu[384].Y = -1230.702

	// punktyWykresuWielomianu[385].X = 3.85
	// punktyWykresuWielomianu[385].Y = -1237.054

	// punktyWykresuWielomianu[386].X = 3.86
	// punktyWykresuWielomianu[386].Y = -1243.427

	// punktyWykresuWielomianu[387].X = 3.87
	// punktyWykresuWielomianu[387].Y = -1249.821

	// punktyWykresuWielomianu[388].X = 3.88
	// punktyWykresuWielomianu[388].Y = -1256.238

	// punktyWykresuWielomianu[389].X = 3.89
	// punktyWykresuWielomianu[389].Y = -1262.676

	// punktyWykresuWielomianu[390].X = 3.9
	// punktyWykresuWielomianu[390].Y = -1269.135

	// punktyWykresuWielomianu[391].X = 3.91
	// punktyWykresuWielomianu[391].Y = -1275.617

	// punktyWykresuWielomianu[392].X = 3.92
	// punktyWykresuWielomianu[392].Y = -1282.121

	// punktyWykresuWielomianu[393].X = 3.93
	// punktyWykresuWielomianu[393].Y = -1288.646

	// punktyWykresuWielomianu[394].X = 3.94
	// punktyWykresuWielomianu[394].Y = -1295.193

	// punktyWykresuWielomianu[395].X = 3.95
	// punktyWykresuWielomianu[395].Y = -1301.762

	// punktyWykresuWielomianu[396].X = 3.96
	// punktyWykresuWielomianu[396].Y = -1308.353

	// punktyWykresuWielomianu[397].X = 3.97
	// punktyWykresuWielomianu[397].Y = -1314.966

	// punktyWykresuWielomianu[398].X = 3.98
	// punktyWykresuWielomianu[398].Y = -1321.601

	// punktyWykresuWielomianu[399].X = 3.99
	// punktyWykresuWielomianu[399].Y = -1328.258

	// punktyWykresuWielomianu[400].X = 4.0
	// punktyWykresuWielomianu[400].Y = -1334.937

	// punktyWykresuWielomianu[401].X = 4.01
	// punktyWykresuWielomianu[401].Y = -1341.638

	// punktyWykresuWielomianu[402].X = 4.02
	// punktyWykresuWielomianu[402].Y = -1348.361

	// punktyWykresuWielomianu[403].X = 4.03
	// punktyWykresuWielomianu[403].Y = -1355.106

	// punktyWykresuWielomianu[404].X = 4.04
	// punktyWykresuWielomianu[404].Y = -1361.874

	// punktyWykresuWielomianu[405].X = 4.05
	// punktyWykresuWielomianu[405].Y = -1368.663

	// punktyWykresuWielomianu[406].X = 4.06
	// punktyWykresuWielomianu[406].Y = -1375.475

	// punktyWykresuWielomianu[407].X = 4.07
	// punktyWykresuWielomianu[407].Y = -1382.309

	// punktyWykresuWielomianu[408].X = 4.08
	// punktyWykresuWielomianu[408].Y = -1389.165

	// punktyWykresuWielomianu[409].X = 4.09
	// punktyWykresuWielomianu[409].Y = -1396.043

	// punktyWykresuWielomianu[410].X = 4.1
	// punktyWykresuWielomianu[410].Y = -1402.943

	// punktyWykresuWielomianu[411].X = 4.11
	// punktyWykresuWielomianu[411].Y = -1409.866

	// punktyWykresuWielomianu[412].X = 4.12
	// punktyWykresuWielomianu[412].Y = -1416.811

	// punktyWykresuWielomianu[413].X = 4.13
	// punktyWykresuWielomianu[413].Y = -1423.779

	// punktyWykresuWielomianu[414].X = 4.14
	// punktyWykresuWielomianu[414].Y = -1430.769

	// punktyWykresuWielomianu[415].X = 4.15
	// punktyWykresuWielomianu[415].Y = -1437.781

	// punktyWykresuWielomianu[416].X = 4.16
	// punktyWykresuWielomianu[416].Y = -1444.815

	// punktyWykresuWielomianu[417].X = 4.17
	// punktyWykresuWielomianu[417].Y = -1451.872

	// punktyWykresuWielomianu[418].X = 4.18
	// punktyWykresuWielomianu[418].Y = -1458.951

	// punktyWykresuWielomianu[419].X = 4.19
	// punktyWykresuWielomianu[419].Y = -1466.053

	// punktyWykresuWielomianu[420].X = 4.20
	// punktyWykresuWielomianu[420].Y = -1473.177

	// punktyWykresuWielomianu[421].X = 4.21
	// punktyWykresuWielomianu[421].Y = -1480.324

	// punktyWykresuWielomianu[422].X = 4.22
	// punktyWykresuWielomianu[422].Y = -1487.493

	// punktyWykresuWielomianu[423].X = 4.23
	// punktyWykresuWielomianu[423].Y = -1494.685

	// punktyWykresuWielomianu[424].X = 4.24
	// punktyWykresuWielomianu[424].Y = -1501.899

	// punktyWykresuWielomianu[425].X = 4.25
	// punktyWykresuWielomianu[425].Y = -1509.136

	// punktyWykresuWielomianu[426].X = 4.26
	// punktyWykresuWielomianu[426].Y = -1516.396

	// punktyWykresuWielomianu[427].X = 4.27
	// punktyWykresuWielomianu[427].Y = -1523.678

	// punktyWykresuWielomianu[428].X = 4.28
	// punktyWykresuWielomianu[428].Y = -1530.982

	// punktyWykresuWielomianu[429].X = 4.29
	// punktyWykresuWielomianu[429].Y = -1538.31

	// punktyWykresuWielomianu[430].X = 4.30
	// punktyWykresuWielomianu[430].Y = -1545.659

	// punktyWykresuWielomianu[431].X = 4.31
	// punktyWykresuWielomianu[431].Y = -1553.032

	// punktyWykresuWielomianu[432].X = 4.32
	// punktyWykresuWielomianu[432].Y = -1560.427

	// punktyWykresuWielomianu[433].X = 4.33
	// punktyWykresuWielomianu[433].Y = -1567.845

	// punktyWykresuWielomianu[434].X = 4.34
	// punktyWykresuWielomianu[434].Y = -1575.286

	// punktyWykresuWielomianu[435].X = 4.35
	// punktyWykresuWielomianu[435].Y = -1582.749

	// punktyWykresuWielomianu[436].X = 4.36
	// punktyWykresuWielomianu[436].Y = -1590.235

	// punktyWykresuWielomianu[437].X = 4.37
	// punktyWykresuWielomianu[437].Y = -1597.744

	// punktyWykresuWielomianu[438].X = 4.38
	// punktyWykresuWielomianu[438].Y = -1605.276

	// punktyWykresuWielomianu[439].X = 4.39
	// punktyWykresuWielomianu[439].Y = -1612.83

	// punktyWykresuWielomianu[440].X = 4.40
	// punktyWykresuWielomianu[440].Y = -1620.407

	// punktyWykresuWielomianu[441].X = 4.41
	// punktyWykresuWielomianu[441].Y = -1628.008

	// punktyWykresuWielomianu[442].X = 4.42
	// punktyWykresuWielomianu[442].Y = -1635.63

	// punktyWykresuWielomianu[443].X = 4.43
	// punktyWykresuWielomianu[443].Y = -1643.276

	// punktyWykresuWielomianu[444].X = 4.44
	// punktyWykresuWielomianu[444].Y = -1650.945

	// punktyWykresuWielomianu[445].X = 4.45
	// punktyWykresuWielomianu[445].Y = -1658.636

	// punktyWykresuWielomianu[446].X = 4.46
	// punktyWykresuWielomianu[446].Y = -1666.351

	// punktyWykresuWielomianu[447].X = 4.47
	// punktyWykresuWielomianu[447].Y = -1674.088

	// punktyWykresuWielomianu[448].X = 4.48
	// punktyWykresuWielomianu[448].Y = -1681.848

	// punktyWykresuWielomianu[449].X = 4.49
	// punktyWykresuWielomianu[449].Y = -1689.631

	// punktyWykresuWielomianu[450].X = 4.5
	// punktyWykresuWielomianu[450].Y = -1697.437

	// punktyWykresuWielomianu[451].X = 4.51
	// punktyWykresuWielomianu[451].Y = -1705.266

	// punktyWykresuWielomianu[452].X = 4.52
	// punktyWykresuWielomianu[452].Y = -1713.118

	// punktyWykresuWielomianu[453].X = 4.53
	// punktyWykresuWielomianu[453].Y = -1720.993

	// punktyWykresuWielomianu[454].X = 4.54
	// punktyWykresuWielomianu[454].Y = -1728.891

	// punktyWykresuWielomianu[455].X = 4.55
	// punktyWykresuWielomianu[455].Y = -1736.812

	// punktyWykresuWielomianu[456].X = 4.56
	// punktyWykresuWielomianu[456].Y = -1744.756

	// punktyWykresuWielomianu[457].X = 4.57
	// punktyWykresuWielomianu[457].Y = -1752.722

	// punktyWykresuWielomianu[458].X = 4.58
	// punktyWykresuWielomianu[458].Y = -1760.712

	// punktyWykresuWielomianu[459].X = 4.59
	// punktyWykresuWielomianu[459].Y = -1768.725

	// punktyWykresuWielomianu[460].X = 4.6
	// punktyWykresuWielomianu[460].Y = -1776.761





	// punktyWykresuWielomianu[491].X = 4.91
	// punktyWykresuWielomianu[491].Y = -2037.354

	// punktyWykresuWielomianu[492].X = 4.92
	// punktyWykresuWielomianu[492].Y = -2046.131

	// punktyWykresuWielomianu[493].X = 4.93
	// punktyWykresuWielomianu[493].Y = -2054.932

	// punktyWykresuWielomianu[494].X = 4.94
	// punktyWykresuWielomianu[494].Y = -2063.756

	// punktyWykresuWielomianu[495].X = 4.95
	// punktyWykresuWielomianu[495].Y = -2072.603

	// punktyWykresuWielomianu[496].X = 4.96
	// punktyWykresuWielomianu[496].Y = -2081.473

	// punktyWykresuWielomianu[497].X = 4.97
	// punktyWykresuWielomianu[497].Y = -2090.367

	// punktyWykresuWielomianu[498].X = 4.98
	// punktyWykresuWielomianu[498].Y = -2099.284

	// punktyWykresuWielomianu[499].X = 4.99
	// punktyWykresuWielomianu[499].Y = -2108.224

	// punktyWykresuWielomianu[500].X = 5.0
	// punktyWykresuWielomianu[500].Y = -2117.187




	// punktyWykresuWielomianu[791].X = 7.91
	// punktyWykresuWielomianu[791].Y = -5641.62

	// punktyWykresuWielomianu[792].X = 7.92
	// punktyWykresuWielomianu[792].Y = -5656

	// punktyWykresuWielomianu[793].X = 7.93
	// punktyWykresuWielomianu[793].Y = -5671.156

	// punktyWykresuWielomianu[794].X = 7.94
	// punktyWykresuWielomianu[794].Y = -5685.944

	// punktyWykresuWielomianu[795].X = 7.95
	// punktyWykresuWielomianu[795].Y = -5700.744

	// punktyWykresuWielomianu[796].X = 7.96
	// punktyWykresuWielomianu[796].Y = -5715.557

	// punktyWykresuWielomianu[797].X = 7.97
	// punktyWykresuWielomianu[797].Y = -5730.383

	// punktyWykresuWielomianu[798].X = 7.98
	// punktyWykresuWielomianu[798].Y = -5745.222

	// punktyWykresuWielomianu[799].X = 7.99
	// punktyWykresuWielomianu[799].Y = -5760.073

	// punktyWykresuWielomianu[800].X = 8.0
	// punktyWykresuWielomianu[800].Y = -5774.937















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
