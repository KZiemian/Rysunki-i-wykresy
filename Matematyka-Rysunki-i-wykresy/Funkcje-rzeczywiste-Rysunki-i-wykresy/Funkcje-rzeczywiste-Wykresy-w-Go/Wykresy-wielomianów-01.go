package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres wielomianu f(x) = x^3 - x^2 + 2x + 1

	punktyWykresuWielomianu := make(plotter.XYs, 1_001)

	punktyWykresuWielomianu[0].X = 0.0
	punktyWykresuWielomianu[0].Y = 1.0

	punktyWykresuWielomianu[1].X = 0.01
	punktyWykresuWielomianu[1].Y = 1.019

	punktyWykresuWielomianu[2].X = 0.02
	punktyWykresuWielomianu[2].Y = 1.039

	punktyWykresuWielomianu[3].X = 0.03
	punktyWykresuWielomianu[3].Y = 1.059

	punktyWykresuWielomianu[4].X = 0.04
	punktyWykresuWielomianu[4].Y = 1.078

	punktyWykresuWielomianu[5].X = 0.05
	punktyWykresuWielomianu[5].Y = 1.097

	punktyWykresuWielomianu[6].X = 0.06
	punktyWykresuWielomianu[6].Y = 1.116

	punktyWykresuWielomianu[7].X = 0.07
	punktyWykresuWielomianu[7].Y = 1.135

	punktyWykresuWielomianu[8].X = 0.08
	punktyWykresuWielomianu[8].Y = 1.154

	punktyWykresuWielomianu[9].X = 0.09
	punktyWykresuWielomianu[9].Y = 1.172

	punktyWykresuWielomianu[10].X = 0.1
	punktyWykresuWielomianu[10].Y = 1.191

	punktyWykresuWielomianu[11].X = 0.11
	punktyWykresuWielomianu[11].Y = 1.209

	punktyWykresuWielomianu[12].X = 0.12
	punktyWykresuWielomianu[12].Y = 1.227

	punktyWykresuWielomianu[13].X = 0.13
	punktyWykresuWielomianu[13].Y = 1.245

	punktyWykresuWielomianu[14].X = 0.14
	punktyWykresuWielomianu[14].Y = 1.263

	punktyWykresuWielomianu[15].X = 0.15
	punktyWykresuWielomianu[15].Y = 1.28

	punktyWykresuWielomianu[16].X = 0.16
	punktyWykresuWielomianu[16].Y = 1.298

	punktyWykresuWielomianu[17].X = 0.17
	punktyWykresuWielomianu[17].Y = 1.316

	punktyWykresuWielomianu[18].X = 0.18
	punktyWykresuWielomianu[18].Y = 1.333

	punktyWykresuWielomianu[19].X = 0.19
	punktyWykresuWielomianu[19].Y = 1.351

	punktyWykresuWielomianu[20].X = 0.2
	punktyWykresuWielomianu[20].Y = 1.368

	punktyWykresuWielomianu[21].X = 0.21
	punktyWykresuWielomianu[21].Y = 1.385

	punktyWykresuWielomianu[22].X = 0.22
	punktyWykresuWielomianu[22].Y = 1.402

	punktyWykresuWielomianu[23].X = 0.23
	punktyWykresuWielomianu[23].Y = 1.419

	punktyWykresuWielomianu[24].X = 0.24
	punktyWykresuWielomianu[24].Y = 1.436

	punktyWykresuWielomianu[25].X = 0.25
	punktyWykresuWielomianu[25].Y = 1.453

	punktyWykresuWielomianu[26].X = 0.26
	punktyWykresuWielomianu[26].Y = 1.47

	punktyWykresuWielomianu[27].X = 0.27
	punktyWykresuWielomianu[27].Y = 1.487

	punktyWykresuWielomianu[28].X = 0.28
	punktyWykresuWielomianu[28].Y = 1.504

	punktyWykresuWielomianu[29].X = 0.29
	punktyWykresuWielomianu[29].Y = 1.520

	punktyWykresuWielomianu[30].X = 0.3
	punktyWykresuWielomianu[30].Y = 1.537

	punktyWykresuWielomianu[31].X = 0.31
	punktyWykresuWielomianu[31].Y = 1.554

	punktyWykresuWielomianu[32].X = 0.32
	punktyWykresuWielomianu[32].Y = 1.570

	punktyWykresuWielomianu[33].X = 0.33
	punktyWykresuWielomianu[33].Y = 1.587

	punktyWykresuWielomianu[34].X = 0.34
	punktyWykresuWielomianu[34].Y = 1.604

	punktyWykresuWielomianu[35].X = 0.35
	punktyWykresuWielomianu[35].Y = 1.62

	punktyWykresuWielomianu[36].X = 0.36
	punktyWykresuWielomianu[36].Y = 1.636

	punktyWykresuWielomianu[37].X = 0.37
	punktyWykresuWielomianu[37].Y = 1.654

	punktyWykresuWielomianu[38].X = 0.38
	punktyWykresuWielomianu[38].Y = 1.67

	punktyWykresuWielomianu[39].X = 0.39
	punktyWykresuWielomianu[39].Y = 1.687

	punktyWykresuWielomianu[40].X = 0.4
	punktyWykresuWielomianu[40].Y = 1.704

	punktyWykresuWielomianu[41].X = 0.41
	punktyWykresuWielomianu[41].Y = 1.721

	punktyWykresuWielomianu[42].X = 0.42
	punktyWykresuWielomianu[42].Y = 1.738

	punktyWykresuWielomianu[43].X = 0.43
	punktyWykresuWielomianu[43].Y = 1.755

	punktyWykresuWielomianu[44].X = 0.44
	punktyWykresuWielomianu[44].Y = 1.772

	punktyWykresuWielomianu[45].X = 0.45
	punktyWykresuWielomianu[45].Y = 1.788

	punktyWykresuWielomianu[46].X = 0.46
	punktyWykresuWielomianu[46].Y = 1.806

	punktyWykresuWielomianu[47].X = 0.47
	punktyWykresuWielomianu[47].Y = 1.823

	punktyWykresuWielomianu[48].X = 0.48
	punktyWykresuWielomianu[48].Y = 1.840

	punktyWykresuWielomianu[49].X = 0.49
	punktyWykresuWielomianu[49].Y = 1.858

	punktyWykresuWielomianu[50].X = 0.5
	punktyWykresuWielomianu[50].Y = 1.875

	punktyWykresuWielomianu[51].X = 0.51
	punktyWykresuWielomianu[51].Y = 1.893

	punktyWykresuWielomianu[52].X = 0.52
	punktyWykresuWielomianu[52].Y = 1.91

	punktyWykresuWielomianu[53].X = 0.53
	punktyWykresuWielomianu[53].Y = 1.928

	punktyWykresuWielomianu[54].X = 0.54
	punktyWykresuWielomianu[54].Y = 1.946

	punktyWykresuWielomianu[55].X = 0.55
	punktyWykresuWielomianu[55].Y = 1.963

	punktyWykresuWielomianu[56].X = 0.56
	punktyWykresuWielomianu[56].Y = 1.982

	punktyWykresuWielomianu[57].X = 0.57
	punktyWykresuWielomianu[57].Y = 2.0

	punktyWykresuWielomianu[58].X = 0.58
	punktyWykresuWielomianu[58].Y = 2.019

	punktyWykresuWielomianu[59].X = 0.59
	punktyWykresuWielomianu[59].Y = 2.037

	punktyWykresuWielomianu[60].X = 0.6
	punktyWykresuWielomianu[60].Y = 2.056

	punktyWykresuWielomianu[61].X = 0.61
	punktyWykresuWielomianu[61].Y = 2.075

	punktyWykresuWielomianu[62].X = 0.62
	punktyWykresuWielomianu[62].Y = 2.094

	punktyWykresuWielomianu[63].X = 0.63
	punktyWykresuWielomianu[63].Y = 2.113

	punktyWykresuWielomianu[64].X = 0.64
	punktyWykresuWielomianu[64].Y = 2.133

	punktyWykresuWielomianu[65].X = 0.65
	punktyWykresuWielomianu[65].Y = 2.152

	punktyWykresuWielomianu[66].X = 0.66
	punktyWykresuWielomianu[66].Y = 2.171

	punktyWykresuWielomianu[67].X = 0.67
	punktyWykresuWielomianu[67].Y = 2.191

	punktyWykresuWielomianu[68].X = 0.68
	punktyWykresuWielomianu[68].Y = 2.212

	punktyWykresuWielomianu[69].X = 0.69
	punktyWykresuWielomianu[69].Y = 2.232

	punktyWykresuWielomianu[70].X = 0.7
	punktyWykresuWielomianu[70].Y = 2.253

	punktyWykresuWielomianu[71].X = 0.71
	punktyWykresuWielomianu[71].Y = 2.273

	punktyWykresuWielomianu[72].X = 0.72
	punktyWykresuWielomianu[72].Y = 2.294

	punktyWykresuWielomianu[73].X = 0.73
	punktyWykresuWielomianu[73].Y = 2.316

	punktyWykresuWielomianu[74].X = 0.74
	punktyWykresuWielomianu[74].Y = 2.337

	punktyWykresuWielomianu[75].X = 0.75
	punktyWykresuWielomianu[75].Y = 2.359

	punktyWykresuWielomianu[76].X = 0.76
	punktyWykresuWielomianu[76].Y = 2.381

	punktyWykresuWielomianu[77].X = 0.77
	punktyWykresuWielomianu[77].Y = 2.403

	punktyWykresuWielomianu[78].X = 0.78
	punktyWykresuWielomianu[78].Y = 2.426

	punktyWykresuWielomianu[79].X = 0.79
	punktyWykresuWielomianu[79].Y = 2.448

	punktyWykresuWielomianu[80].X = 0.8
	punktyWykresuWielomianu[80].Y = 2.472

	punktyWykresuWielomianu[81].X = 0.81
	punktyWykresuWielomianu[81].Y = 2.495

	punktyWykresuWielomianu[82].X = 0.82
	punktyWykresuWielomianu[82].Y = 2.519

	punktyWykresuWielomianu[83].X = 0.83
	punktyWykresuWielomianu[83].Y = 2.542

	punktyWykresuWielomianu[84].X = 0.84
	punktyWykresuWielomianu[84].Y = 2.567

	punktyWykresuWielomianu[85].X = 0.85
	punktyWykresuWielomianu[85].Y = 2.591

	punktyWykresuWielomianu[86].X = 0.86
	punktyWykresuWielomianu[86].Y = 2.616

	punktyWykresuWielomianu[87].X = 0.87
	punktyWykresuWielomianu[87].Y = 2.641

	punktyWykresuWielomianu[88].X = 0.88
	punktyWykresuWielomianu[88].Y = 2.667

	punktyWykresuWielomianu[89].X = 0.89
	punktyWykresuWielomianu[89].Y = 2.692

	punktyWykresuWielomianu[90].X = 0.9
	punktyWykresuWielomianu[90].Y = 2.719

	punktyWykresuWielomianu[91].X = 0.91
	punktyWykresuWielomianu[91].Y = 2.745

	punktyWykresuWielomianu[92].X = 0.92
	punktyWykresuWielomianu[92].Y = 2.772

	punktyWykresuWielomianu[93].X = 0.93
	punktyWykresuWielomianu[93].Y = 2.799

	punktyWykresuWielomianu[94].X = 0.94
	punktyWykresuWielomianu[94].Y = 2.827

	punktyWykresuWielomianu[95].X = 0.95
	punktyWykresuWielomianu[95].Y = 2.854

	punktyWykresuWielomianu[96].X = 0.96
	punktyWykresuWielomianu[96].Y = 2.883

	punktyWykresuWielomianu[97].X = 0.97
	punktyWykresuWielomianu[97].Y = 2.911

	punktyWykresuWielomianu[98].X = 0.98
	punktyWykresuWielomianu[98].Y = 2.94

	punktyWykresuWielomianu[99].X = 0.99
	punktyWykresuWielomianu[99].Y = 2.97

	punktyWykresuWielomianu[100].X = 1.0
	punktyWykresuWielomianu[100].Y = 3.0

	punktyWykresuWielomianu[101].X = 1.01
	punktyWykresuWielomianu[101].Y = 3.03

	punktyWykresuWielomianu[102].X = 1.02
	punktyWykresuWielomianu[102].Y = 3.06

	punktyWykresuWielomianu[103].X = 1.03
	punktyWykresuWielomianu[103].Y = 3.091

	punktyWykresuWielomianu[104].X = 1.04
	punktyWykresuWielomianu[104].Y = 3.123

	punktyWykresuWielomianu[105].X = 1.05
	punktyWykresuWielomianu[105].Y = 3.155

	punktyWykresuWielomianu[106].X = 1.06
	punktyWykresuWielomianu[106].Y = 3.187

	punktyWykresuWielomianu[107].X = 1.07
	punktyWykresuWielomianu[107].Y = 3.22

	punktyWykresuWielomianu[108].X = 1.08
	punktyWykresuWielomianu[108].Y = 3.253

	punktyWykresuWielomianu[109].X = 1.09
	punktyWykresuWielomianu[109].Y = 3.286

	punktyWykresuWielomianu[110].X = 1.1
	punktyWykresuWielomianu[110].Y = 3.321

	punktyWykresuWielomianu[111].X = 1.11
	punktyWykresuWielomianu[111].Y = 3.355

	punktyWykresuWielomianu[112].X = 1.12
	punktyWykresuWielomianu[112].Y = 3.39

	punktyWykresuWielomianu[113].X = 1.13
	punktyWykresuWielomianu[113].Y = 3.426

	punktyWykresuWielomianu[114].X = 1.14
	punktyWykresuWielomianu[114].Y = 3.461

	punktyWykresuWielomianu[115].X = 1.15
	punktyWykresuWielomianu[115].Y = 3.498

	punktyWykresuWielomianu[116].X = 1.16
	punktyWykresuWielomianu[116].Y = 3.535

	punktyWykresuWielomianu[117].X = 1.17
	punktyWykresuWielomianu[117].Y = 3.572

	punktyWykresuWielomianu[118].X = 1.18
	punktyWykresuWielomianu[118].Y = 3.61

	punktyWykresuWielomianu[119].X = 1.19
	punktyWykresuWielomianu[119].Y = 3.649

	punktyWykresuWielomianu[120].X = 1.2
	punktyWykresuWielomianu[120].Y = 3.688

	punktyWykresuWielomianu[121].X = 1.21
	punktyWykresuWielomianu[121].Y = 3.727

	punktyWykresuWielomianu[122].X = 1.22
	punktyWykresuWielomianu[122].Y = 3.767

	punktyWykresuWielomianu[123].X = 1.23
	punktyWykresuWielomianu[123].Y = 3.808

	punktyWykresuWielomianu[124].X = 1.24
	punktyWykresuWielomianu[124].Y = 3.849

	punktyWykresuWielomianu[125].X = 1.25
	punktyWykresuWielomianu[125].Y = 3.89

	punktyWykresuWielomianu[126].X = 1.26
	punktyWykresuWielomianu[126].Y = 3.932

	punktyWykresuWielomianu[127].X = 1.27
	punktyWykresuWielomianu[127].Y = 3.975

	punktyWykresuWielomianu[128].X = 1.28
	punktyWykresuWielomianu[128].Y = 4.018

	punktyWykresuWielomianu[129].X = 1.29
	punktyWykresuWielomianu[129].Y = 4.062

	punktyWykresuWielomianu[130].X = 1.3
	punktyWykresuWielomianu[130].Y = 4.107

	punktyWykresuWielomianu[131].X = 1.31
	punktyWykresuWielomianu[131].Y = 4.107

	punktyWykresuWielomianu[132].X = 1.32
	punktyWykresuWielomianu[132].Y = 4.197

	punktyWykresuWielomianu[133].X = 1.33
	punktyWykresuWielomianu[133].Y = 4.243

	punktyWykresuWielomianu[134].X = 1.34
	punktyWykresuWielomianu[134].Y = 4.29

	punktyWykresuWielomianu[135].X = 1.35
	punktyWykresuWielomianu[135].Y = 4.337

	punktyWykresuWielomianu[136].X = 1.36
	punktyWykresuWielomianu[136].Y = 4.385

	punktyWykresuWielomianu[137].X = 1.37
	punktyWykresuWielomianu[137].Y = 4.434

	punktyWykresuWielomianu[138].X = 1.38
	punktyWykresuWielomianu[138].Y = 4.483

	punktyWykresuWielomianu[139].X = 1.39
	punktyWykresuWielomianu[139].Y = 4.533

	punktyWykresuWielomianu[140].X = 1.4
	punktyWykresuWielomianu[140].Y = 4.584

	punktyWykresuWielomianu[141].X = 1.41
	punktyWykresuWielomianu[141].Y = 4.635

	punktyWykresuWielomianu[142].X = 1.42
	punktyWykresuWielomianu[142].Y = 4.686

	punktyWykresuWielomianu[143].X = 1.43
	punktyWykresuWielomianu[143].Y = 4.739

	punktyWykresuWielomianu[144].X = 1.44
	punktyWykresuWielomianu[144].Y = 4.792

	punktyWykresuWielomianu[145].X = 1.45
	punktyWykresuWielomianu[145].Y = 4.846

	punktyWykresuWielomianu[146].X = 1.46
	punktyWykresuWielomianu[146].Y = 4.9

	punktyWykresuWielomianu[147].X = 1.47
	punktyWykresuWielomianu[147].Y = 4.955

	punktyWykresuWielomianu[146].X = 1.48
	punktyWykresuWielomianu[146].Y = 5.011

	punktyWykresuWielomianu[146].X = 1.49
	punktyWykresuWielomianu[146].Y = 5.067

	punktyWykresuWielomianu[150].X = 1.5
	punktyWykresuWielomianu[150].Y = 5.125

	punktyWykresuWielomianu[151].X = 1.51
	punktyWykresuWielomianu[151].Y = 5.182

	punktyWykresuWielomianu[152].X = 1.52
	punktyWykresuWielomianu[152].Y = 5.241

	punktyWykresuWielomianu[153].X = 1.53
	punktyWykresuWielomianu[153].Y = 5.3

	punktyWykresuWielomianu[154].X = 1.5
	punktyWykresuWielomianu[154].Y = 5.36

	punktyWykresuWielomianu[155].X = 1.55
	punktyWykresuWielomianu[155].Y = 5.421

	punktyWykresuWielomianu[156].X = 1.56
	punktyWykresuWielomianu[156].Y = 5.482

	punktyWykresuWielomianu[157].X = 1.57
	punktyWykresuWielomianu[157].Y = 5.545

	punktyWykresuWielomianu[158].X = 1.58
	punktyWykresuWielomianu[158].Y = 5.607

	punktyWykresuWielomianu[159].X = 1.59
	punktyWykresuWielomianu[159].Y = 5.671

	punktyWykresuWielomianu[160].X = 1.6
	punktyWykresuWielomianu[160].Y = 5.736

	punktyWykresuWielomianu[161].X = 1.61
	punktyWykresuWielomianu[161].Y = 5.801

	punktyWykresuWielomianu[162].X = 1.62
	punktyWykresuWielomianu[162].Y = 5.867

	punktyWykresuWielomianu[163].X = 1.63
	punktyWykresuWielomianu[163].Y = 5.933

	punktyWykresuWielomianu[164].X = 1.64
	punktyWykresuWielomianu[164].Y = 6.001

	punktyWykresuWielomianu[165].X = 1.65
	punktyWykresuWielomianu[165].Y = 6.069

	punktyWykresuWielomianu[166].X = 1.66
	punktyWykresuWielomianu[166].Y = 6.138

	punktyWykresuWielomianu[167].X = 1.67
	punktyWykresuWielomianu[167].Y = 6.208

	punktyWykresuWielomianu[168].X = 1.68
	punktyWykresuWielomianu[168].Y = 6.279

	punktyWykresuWielomianu[169].X = 1.69
	punktyWykresuWielomianu[169].Y = 6.35

	punktyWykresuWielomianu[170].X = 1.7
	punktyWykresuWielomianu[170].Y = 6.423

	punktyWykresuWielomianu[171].X = 1.71
	punktyWykresuWielomianu[171].Y = 6.496

	punktyWykresuWielomianu[172].X = 1.72
	punktyWykresuWielomianu[172].Y = 6.57

	punktyWykresuWielomianu[173].X = 1.73
	punktyWykresuWielomianu[173].Y = 6.644

	punktyWykresuWielomianu[174].X = 1.74
	punktyWykresuWielomianu[174].Y = 6.72

	punktyWykresuWielomianu[175].X = 1.75
	punktyWykresuWielomianu[175].Y = 6.796

	punktyWykresuWielomianu[176].X = 1.76
	punktyWykresuWielomianu[176].Y = 6.874

	punktyWykresuWielomianu[177].X = 1.77
	punktyWykresuWielomianu[177].Y = 6.952

	punktyWykresuWielomianu[178].X = 1.78
	punktyWykresuWielomianu[178].Y = 7.031

	punktyWykresuWielomianu[179].X = 1.79
	punktyWykresuWielomianu[179].Y = 7.111

	punktyWykresuWielomianu[180].X = 1.8
	punktyWykresuWielomianu[180].Y = 7.192

	punktyWykresuWielomianu[181].X = 1.81
	punktyWykresuWielomianu[181].Y = 7.273

	punktyWykresuWielomianu[182].X = 1.82
	punktyWykresuWielomianu[182].Y = 7.356

	punktyWykresuWielomianu[183].X = 1.83
	punktyWykresuWielomianu[183].Y = 7.439

	punktyWykresuWielomianu[184].X = 1.84
	punktyWykresuWielomianu[184].Y = 7.523

	punktyWykresuWielomianu[185].X = 1.85
	punktyWykresuWielomianu[185].Y = 7.609

	punktyWykresuWielomianu[186].X = 1.86
	punktyWykresuWielomianu[186].Y = 7.695

	punktyWykresuWielomianu[187].X = 1.87
	punktyWykresuWielomianu[187].Y = 7.782

	punktyWykresuWielomianu[188].X = 1.88
	punktyWykresuWielomianu[188].Y = 7.87

	punktyWykresuWielomianu[189].X = 1.89
	punktyWykresuWielomianu[189].Y = 7.959

	punktyWykresuWielomianu[190].X = 1.9
	punktyWykresuWielomianu[190].Y = 8.049

	punktyWykresuWielomianu[191].X = 1.91
	punktyWykresuWielomianu[191].Y = 8.139

	punktyWykresuWielomianu[192].X = 1.92
	punktyWykresuWielomianu[192].Y = 8.231

	punktyWykresuWielomianu[193].X = 1.93
	punktyWykresuWielomianu[193].Y = 8.324

	punktyWykresuWielomianu[194].X = 1.94
	punktyWykresuWielomianu[194].Y = 8.417

	punktyWykresuWielomianu[195].X = 1.95
	punktyWykresuWielomianu[195].Y = 8.512

	punktyWykresuWielomianu[196].X = 1.96
	punktyWykresuWielomianu[196].Y = 8.607

	punktyWykresuWielomianu[197].X = 1.97
	punktyWykresuWielomianu[197].Y = 8.704

	punktyWykresuWielomianu[198].X = 1.98
	punktyWykresuWielomianu[198].Y = 8.802

	punktyWykresuWielomianu[199].X = 1.99
	punktyWykresuWielomianu[199].Y = 8.9

	punktyWykresuWielomianu[200].X = 2.0
	punktyWykresuWielomianu[200].Y = 9.0

	punktyWykresuWielomianu[201].X = 2.01
	punktyWykresuWielomianu[201].Y = 9.1

	punktyWykresuWielomianu[202].X = 2.02
	punktyWykresuWielomianu[202].Y = 9.202

	punktyWykresuWielomianu[203].X = 2.03
	punktyWykresuWielomianu[203].Y = 9.304

	punktyWykresuWielomianu[204].X = 2.04
	punktyWykresuWielomianu[204].Y = 9.408

	punktyWykresuWielomianu[205].X = 2.05
	punktyWykresuWielomianu[205].Y = 9.512

	punktyWykresuWielomianu[206].X = 2.06
	punktyWykresuWielomianu[206].Y = 9.618

	punktyWykresuWielomianu[207].X = 2.07
	punktyWykresuWielomianu[207].Y = 9.724

	punktyWykresuWielomianu[208].X = 2.08
	punktyWykresuWielomianu[208].Y = 9.832

	punktyWykresuWielomianu[209].X = 2.09
	punktyWykresuWielomianu[209].Y = 9.941

	punktyWykresuWielomianu[210].X = 2.1
	punktyWykresuWielomianu[210].Y = 10.051

	punktyWykresuWielomianu[211].X = 2.11
	punktyWykresuWielomianu[211].Y = 10.161

	punktyWykresuWielomianu[212].X = 2.12
	punktyWykresuWielomianu[212].Y = 10.273

	punktyWykresuWielomianu[213].X = 2.13
	punktyWykresuWielomianu[213].Y = 10.386

	punktyWykresuWielomianu[214].X = 2.14
	punktyWykresuWielomianu[214].Y = 10.5

	punktyWykresuWielomianu[215].X = 2.15
	punktyWykresuWielomianu[215].Y = 10.615

	punktyWykresuWielomianu[216].X = 2.16
	punktyWykresuWielomianu[216].Y = 10.732

	punktyWykresuWielomianu[217].X = 2.17
	punktyWykresuWielomianu[217].Y = 10.849

	punktyWykresuWielomianu[218].X = 2.18
	punktyWykresuWielomianu[218].Y = 10.967

	punktyWykresuWielomianu[219].X = 2.19
	punktyWykresuWielomianu[219].Y = 11.087

	punktyWykresuWielomianu[220].X = 2.2
	punktyWykresuWielomianu[220].Y = 11.208

	punktyWykresuWielomianu[221].X = 2.21
	punktyWykresuWielomianu[221].Y = 11.329

	punktyWykresuWielomianu[222].X = 2.22
	punktyWykresuWielomianu[222].Y = 11.452

	punktyWykresuWielomianu[223].X = 2.23
	punktyWykresuWielomianu[223].Y = 11.576

	punktyWykresuWielomianu[224].X = 2.24
	punktyWykresuWielomianu[224].Y = 11.701

	punktyWykresuWielomianu[225].X = 2.25
	punktyWykresuWielomianu[225].Y = 11.828

	punktyWykresuWielomianu[226].X = 2.26
	punktyWykresuWielomianu[226].Y = 11.955

	punktyWykresuWielomianu[227].X = 2.27
	punktyWykresuWielomianu[227].Y = 12.084

	punktyWykresuWielomianu[228].X = 2.28
	punktyWykresuWielomianu[228].Y = 12.214

	punktyWykresuWielomianu[229].X = 2.29
	punktyWykresuWielomianu[229].Y = 12.344

	punktyWykresuWielomianu[230].X = 2.3
	punktyWykresuWielomianu[230].Y = 12.477

	punktyWykresuWielomianu[231].X = 2.31
	punktyWykresuWielomianu[231].Y = 12.61

	punktyWykresuWielomianu[232].X = 2.32
	punktyWykresuWielomianu[232].Y = 12.744

	punktyWykresuWielomianu[233].X = 2.33
	punktyWykresuWielomianu[233].Y = 12.88

	punktyWykresuWielomianu[234].X = 2.34
	punktyWykresuWielomianu[234].Y = 12.017

	punktyWykresuWielomianu[235].X = 2.35
	punktyWykresuWielomianu[235].Y = 13.155

	punktyWykresuWielomianu[236].X = 2.36
	punktyWykresuWielomianu[236].Y = 13.294

	punktyWykresuWielomianu[237].X = 2.37
	punktyWykresuWielomianu[237].Y = 13.435

	punktyWykresuWielomianu[238].X = 2.38
	punktyWykresuWielomianu[238].Y = 13.576

	punktyWykresuWielomianu[239].X = 2.39
	punktyWykresuWielomianu[239].Y = 13.719

	punktyWykresuWielomianu[240].X = 2.4
	punktyWykresuWielomianu[240].Y = 13.864

	punktyWykresuWielomianu[241].X = 2.41
	punktyWykresuWielomianu[241].Y = 14.009

	punktyWykresuWielomianu[242].X = 2.42
	punktyWykresuWielomianu[242].Y = 14.156

	punktyWykresuWielomianu[243].X = 2.43
	punktyWykresuWielomianu[243].Y = 14.304

	punktyWykresuWielomianu[244].X = 2.44
	punktyWykresuWielomianu[244].Y = 14.453

	punktyWykresuWielomianu[245].X = 2.45
	punktyWykresuWielomianu[245].Y = 14.603

	punktyWykresuWielomianu[246].X = 2.46
	punktyWykresuWielomianu[246].Y = 14.755

	punktyWykresuWielomianu[247].X = 2.47
	punktyWykresuWielomianu[247].Y = 14.908

	punktyWykresuWielomianu[248].X = 2.48
	punktyWykresuWielomianu[248].Y = 15.062

	punktyWykresuWielomianu[249].X = 2.49
	punktyWykresuWielomianu[249].Y = 15.218

	punktyWykresuWielomianu[250].X = 2.5
	punktyWykresuWielomianu[250].Y = 15.375

	punktyWykresuWielomianu[251].X = 2.51
	punktyWykresuWielomianu[251].Y = 15.533

	punktyWykresuWielomianu[252].X = 2.52
	punktyWykresuWielomianu[252].Y = 15.692

	punktyWykresuWielomianu[253].X = 2.53
	punktyWykresuWielomianu[253].Y = 15.853

	punktyWykresuWielomianu[254].X = 2.54
	punktyWykresuWielomianu[254].Y = 16.015

	punktyWykresuWielomianu[255].X = 2.55
	punktyWykresuWielomianu[255].Y = 16.178

	punktyWykresuWielomianu[256].X = 2.56
	punktyWykresuWielomianu[256].Y = 16.343

	punktyWykresuWielomianu[257].X = 2.57
	punktyWykresuWielomianu[257].Y = 16.509

	punktyWykresuWielomianu[258].X = 2.58
	punktyWykresuWielomianu[258].Y = 16.677

	punktyWykresuWielomianu[259].X = 2.59
	punktyWykresuWielomianu[259].Y = 16.845

	punktyWykresuWielomianu[260].X = 2.6
	punktyWykresuWielomianu[260].Y = 17.016

	punktyWykresuWielomianu[261].X = 2.61
	punktyWykresuWielomianu[261].Y = 17.186

	punktyWykresuWielomianu[262].X = 2.62
	punktyWykresuWielomianu[262].Y = 17.36

	punktyWykresuWielomianu[263].X = 2.63
	punktyWykresuWielomianu[263].Y = 17.534

	punktyWykresuWielomianu[264].X = 2.64
	punktyWykresuWielomianu[264].Y = 17.71

	punktyWykresuWielomianu[265].X = 2.65
	punktyWykresuWielomianu[265].Y = 17.887

	punktyWykresuWielomianu[266].X = 2.66
	punktyWykresuWielomianu[266].Y = 18.065

	punktyWykresuWielomianu[267].X = 2.67
	punktyWykresuWielomianu[267].Y = 18.245

	punktyWykresuWielomianu[268].X = 2.68
	punktyWykresuWielomianu[268].Y = 18.426

	punktyWykresuWielomianu[269].X = 2.69
	punktyWykresuWielomianu[269].Y = 18.609

	punktyWykresuWielomianu[270].X = 2.7
	punktyWykresuWielomianu[270].Y = 18.793

	punktyWykresuWielomianu[271].X = 2.71
	punktyWykresuWielomianu[271].Y = 18.978

	punktyWykresuWielomianu[272].X = 2.72
	punktyWykresuWielomianu[272].Y = 19.165

	punktyWykresuWielomianu[273].X = 2.73
	punktyWykresuWielomianu[273].Y = 19.353

	punktyWykresuWielomianu[274].X = 2.74
	punktyWykresuWielomianu[274].Y = 19.543

	punktyWykresuWielomianu[275].X = 2.75
	punktyWykresuWielomianu[275].Y = 19.734

	punktyWykresuWielomianu[276].X = 2.76
	punktyWykresuWielomianu[276].Y = 19.927

	punktyWykresuWielomianu[277].X = 2.77
	punktyWykresuWielomianu[277].Y = 20.121

	punktyWykresuWielomianu[278].X = 2.78
	punktyWykresuWielomianu[278].Y = 20.316

	punktyWykresuWielomianu[279].X = 2.79
	punktyWykresuWielomianu[279].Y = 20.513

	punktyWykresuWielomianu[280].X = 2.8
	punktyWykresuWielomianu[280].Y = 20.712

	punktyWykresuWielomianu[281].X = 2.81
	punktyWykresuWielomianu[281].Y = 20.911

	punktyWykresuWielomianu[282].X = 2.82
	punktyWykresuWielomianu[282].Y = 21.113

	punktyWykresuWielomianu[283].X = 2.83
	punktyWykresuWielomianu[283].Y = 21.316

	punktyWykresuWielomianu[284].X = 2.84
	punktyWykresuWielomianu[284].Y = 21.52

	punktyWykresuWielomianu[285].X = 2.85
	punktyWykresuWielomianu[285].Y = 21.726

	punktyWykresuWielomianu[286].X = 2.86
	punktyWykresuWielomianu[286].Y = 21.934

	punktyWykresuWielomianu[287].X = 2.87
	punktyWykresuWielomianu[287].Y = 22.143

	punktyWykresuWielomianu[288].X = 2.88
	punktyWykresuWielomianu[288].Y = 22.353

	punktyWykresuWielomianu[289].X = 2.89
	punktyWykresuWielomianu[289].Y = 22.565

	punktyWykresuWielomianu[290].X = 2.9
	punktyWykresuWielomianu[290].Y = 22.779

	punktyWykresuWielomianu[291].X = 2.91
	punktyWykresuWielomianu[291].Y = 22.994

	punktyWykresuWielomianu[292].X = 2.92
	punktyWykresuWielomianu[292].Y = 23.21

	punktyWykresuWielomianu[293].X = 2.93
	punktyWykresuWielomianu[293].Y = 23.428

	punktyWykresuWielomianu[294].X = 2.94
	punktyWykresuWielomianu[294].Y = 23.648

	punktyWykresuWielomianu[295].X = 2.95
	punktyWykresuWielomianu[295].Y = 23.869

	punktyWykresuWielomianu[296].X = 2.96
	punktyWykresuWielomianu[296].Y = 24.092

	punktyWykresuWielomianu[297].X = 2.97
	punktyWykresuWielomianu[297].Y = 24.317

	punktyWykresuWielomianu[298].X = 2.98
	punktyWykresuWielomianu[298].Y = 24.543

	punktyWykresuWielomianu[299].X = 2.99
	punktyWykresuWielomianu[299].Y = 24.77

	punktyWykresuWielomianu[300].X = 3.0
	punktyWykresuWielomianu[300].Y = 25.0

	punktyWykresuWielomianu[301].X = 3.01
	punktyWykresuWielomianu[301].Y = 25.23

	punktyWykresuWielomianu[302].X = 3.02
	punktyWykresuWielomianu[302].Y = 25.463

	punktyWykresuWielomianu[303].X = 3.03
	punktyWykresuWielomianu[303].Y = 25.697

	punktyWykresuWielomianu[304].X = 3.04
	punktyWykresuWielomianu[304].Y = 25.932

	punktyWykresuWielomianu[305].X = 3.05
	punktyWykresuWielomianu[305].Y = 26.17

	punktyWykresuWielomianu[306].X = 3.06
	punktyWykresuWielomianu[306].Y = 26.409

	punktyWykresuWielomianu[307].X = 3.07
	punktyWykresuWielomianu[307].Y = 26.649

	punktyWykresuWielomianu[308].X = 3.08
	punktyWykresuWielomianu[308].Y = 26.891

	punktyWykresuWielomianu[309].X = 3.09
	punktyWykresuWielomianu[309].Y = 27.135

	punktyWykresuWielomianu[310].X = 3.1
	punktyWykresuWielomianu[310].Y = 27.381

	punktyWykresuWielomianu[311].X = 3.11
	punktyWykresuWielomianu[311].Y = 27.628

	punktyWykresuWielomianu[312].X = 3.12
	punktyWykresuWielomianu[312].Y = 27.876

	punktyWykresuWielomianu[313].X = 3.13
	punktyWykresuWielomianu[313].Y = 28.127

	punktyWykresuWielomianu[314].X = 3.14
	punktyWykresuWielomianu[314].Y = 28.379

	punktyWykresuWielomianu[315].X = 3.15
	punktyWykresuWielomianu[315].Y = 28.633

	punktyWykresuWielomianu[316].X = 3.16
	punktyWykresuWielomianu[316].Y = 28.888

	punktyWykresuWielomianu[317].X = 3.17
	punktyWykresuWielomianu[317].Y = 29.146

	punktyWykresuWielomianu[318].X = 3.18
	punktyWykresuWielomianu[318].Y = 29.405

	punktyWykresuWielomianu[319].X = 3.19
	punktyWykresuWielomianu[319].Y = 29.665

	punktyWykresuWielomianu[320].X = 3.2
	punktyWykresuWielomianu[320].Y = 29.928

	punktyWykresuWielomianu[321].X = 3.21
	punktyWykresuWielomianu[321].Y = 30.192

	punktyWykresuWielomianu[322].X = 3.22
	punktyWykresuWielomianu[322].Y = 30.457

	punktyWykresuWielomianu[323].X = 3.23
	punktyWykresuWielomianu[323].Y = 30.725

	punktyWykresuWielomianu[324].X = 3.24
	punktyWykresuWielomianu[324].Y = 30.994

	punktyWykresuWielomianu[325].X = 3.25
	punktyWykresuWielomianu[325].Y = 31.265

	punktyWykresuWielomianu[326].X = 3.26
	punktyWykresuWielomianu[326].Y = 31.538

	punktyWykresuWielomianu[327].X = 3.27
	punktyWykresuWielomianu[327].Y = 31.812

	punktyWykresuWielomianu[328].X = 3.28
	punktyWykresuWielomianu[328].Y = 32.089

	punktyWykresuWielomianu[329].X = 3.29
	punktyWykresuWielomianu[329].Y = 32.367

	punktyWykresuWielomianu[330].X = 3.3
	punktyWykresuWielomianu[330].Y = 32.647

	punktyWykresuWielomianu[331].X = 3.31
	punktyWykresuWielomianu[331].Y = 32.928

	punktyWykresuWielomianu[332].X = 3.32
	punktyWykresuWielomianu[332].Y = 33.212

	punktyWykresuWielomianu[333].X = 3.33
	punktyWykresuWielomianu[333].Y = 33.497

	punktyWykresuWielomianu[334].X = 3.34
	punktyWykresuWielomianu[334].Y = 33.784

	punktyWykresuWielomianu[335].X = 3.35
	punktyWykresuWielomianu[335].Y = 34.072

	punktyWykresuWielomianu[336].X = 3.36
	punktyWykresuWielomianu[336].Y = 34.363

	punktyWykresuWielomianu[337].X = 3.37
	punktyWykresuWielomianu[337].Y = 34.655

	punktyWykresuWielomianu[338].X = 3.38
	punktyWykresuWielomianu[338].Y = 34.95

	punktyWykresuWielomianu[339].X = 3.39
	punktyWykresuWielomianu[339].Y = 35.246

	punktyWykresuWielomianu[340].X = 3.4
	punktyWykresuWielomianu[340].Y = 35.544

	punktyWykresuWielomianu[341].X = 3.41
	punktyWykresuWielomianu[341].Y = 35.843

	punktyWykresuWielomianu[342].X = 3.42
	punktyWykresuWielomianu[342].Y = 36.145

	punktyWykresuWielomianu[343].X = 3.43
	punktyWykresuWielomianu[343].Y = 36.448

	punktyWykresuWielomianu[344].X = 3.44
	punktyWykresuWielomianu[344].Y = 37.754

	punktyWykresuWielomianu[345].X = 3.45
	punktyWykresuWielomianu[345].Y = 37.061

	punktyWykresuWielomianu[346].X = 3.46
	punktyWykresuWielomianu[346].Y = 37.37

	punktyWykresuWielomianu[347].X = 3.47
	punktyWykresuWielomianu[347].Y = 37.681

	punktyWykresuWielomianu[348].X = 3.48
	punktyWykresuWielomianu[348].Y = 37.993

	punktyWykresuWielomianu[349].X = 3.49
	punktyWykresuWielomianu[349].Y = 38.308

	punktyWykresuWielomianu[350].X = 3.5
	punktyWykresuWielomianu[350].Y = 38.625

	punktyWykresuWielomianu[351].X = 3.51
	punktyWykresuWielomianu[351].Y = 38.943

	punktyWykresuWielomianu[352].X = 3.52
	punktyWykresuWielomianu[352].Y = 39.263

	punktyWykresuWielomianu[353].X = 3.53
	punktyWykresuWielomianu[353].Y = 39.586

	punktyWykresuWielomianu[354].X = 3.54
	punktyWykresuWielomianu[354].Y = 39.91

	punktyWykresuWielomianu[355].X = 3.55
	punktyWykresuWielomianu[355].Y = 40.236

	punktyWykresuWielomianu[356].X = 3.56
	punktyWykresuWielomianu[356].Y = 40.564

	punktyWykresuWielomianu[357].X = 3.57
	punktyWykresuWielomianu[357].Y = 40.894

	punktyWykresuWielomianu[358].X = 3.58
	punktyWykresuWielomianu[358].Y = 41.226

	punktyWykresuWielomianu[359].X = 3.59
	punktyWykresuWielomianu[359].Y = 41.56

	punktyWykresuWielomianu[360].X = 3.6
	punktyWykresuWielomianu[360].Y = 41.896

	punktyWykresuWielomianu[361].X = 3.61
	punktyWykresuWielomianu[361].Y = 42.233

	punktyWykresuWielomianu[362].X = 3.62
	punktyWykresuWielomianu[362].Y = 42.573

	punktyWykresuWielomianu[363].X = 3.65
	punktyWykresuWielomianu[363].Y = 42.915

	punktyWykresuWielomianu[364].X = 3.64
	punktyWykresuWielomianu[364].Y = 43.258

	punktyWykresuWielomianu[365].X = 3.65
	punktyWykresuWielomianu[365].Y = 43.604

	punktyWykresuWielomianu[366].X = 3.66
	punktyWykresuWielomianu[366].Y = 43.952

	punktyWykresuWielomianu[367].X = 3.67
	punktyWykresuWielomianu[367].Y = 44.302

	punktyWykresuWielomianu[368].X = 3.68
	punktyWykresuWielomianu[368].Y = 44.653

	punktyWykresuWielomianu[369].X = 3.69
	punktyWykresuWielomianu[369].Y = 45.007

	punktyWykresuWielomianu[370].X = 3.7
	punktyWykresuWielomianu[370].Y = 45.363

	punktyWykresuWielomianu[371].X = 3.71
	punktyWykresuWielomianu[371].Y = 45.72

	punktyWykresuWielomianu[372].X = 3.72
	punktyWykresuWielomianu[372].Y = 46.08

	punktyWykresuWielomianu[373].X = 3.73
	punktyWykresuWielomianu[373].Y = 46.442

	punktyWykresuWielomianu[374].X = 3.74
	punktyWykresuWielomianu[374].Y = 46.806

	punktyWykresuWielomianu[375].X = 3.75
	punktyWykresuWielomianu[375].Y = 47.171

	punktyWykresuWielomianu[376].X = 3.76
	punktyWykresuWielomianu[376].Y = 47.539

	punktyWykresuWielomianu[377].X = 3.77
	punktyWykresuWielomianu[377].Y = 47.909

	punktyWykresuWielomianu[378].X = 3.78
	punktyWykresuWielomianu[378].Y = 48.281

	punktyWykresuWielomianu[379].X = 3.79
	punktyWykresuWielomianu[379].Y = 48.655

	punktyWykresuWielomianu[380].X = 3.8
	punktyWykresuWielomianu[380].Y = 49.032

	punktyWykresuWielomianu[381].X = 3.81
	punktyWykresuWielomianu[381].Y = 49.41

	punktyWykresuWielomianu[382].X = 3.82
	punktyWykresuWielomianu[382].Y = 49.79

	punktyWykresuWielomianu[383].X = 3.83
	punktyWykresuWielomianu[383].Y = 50.173

	punktyWykresuWielomianu[384].X = 3.84
	punktyWykresuWielomianu[384].Y = 50.557

	punktyWykresuWielomianu[385].X = 3.85
	punktyWykresuWielomianu[385].Y = 50.944

	punktyWykresuWielomianu[386].X = 3.86
	punktyWykresuWielomianu[386].Y = 51.332

	punktyWykresuWielomianu[387].X = 3.87
	punktyWykresuWielomianu[387].Y = 51.723

	punktyWykresuWielomianu[388].X = 3.88
	punktyWykresuWielomianu[388].Y = 52.116

	punktyWykresuWielomianu[389].X = 3.89
	punktyWykresuWielomianu[389].Y = 52.511

	punktyWykresuWielomianu[390].X = 3.9
	punktyWykresuWielomianu[390].Y = 52.909

	punktyWykresuWielomianu[391].X = 3.91
	punktyWykresuWielomianu[391].Y = 53.308

	punktyWykresuWielomianu[392].X = 3.92
	punktyWykresuWielomianu[392].Y = 53.709

	punktyWykresuWielomianu[393].X = 3.93
	punktyWykresuWielomianu[393].Y = 54.113

	punktyWykresuWielomianu[394].X = 3.94
	punktyWykresuWielomianu[394].Y = 54.519

	punktyWykresuWielomianu[395].X = 3.95
	punktyWykresuWielomianu[395].Y = 54.927

	punktyWykresuWielomianu[396].X = 3.96
	punktyWykresuWielomianu[396].Y = 55.337

	punktyWykresuWielomianu[397].X = 3.97
	punktyWykresuWielomianu[397].Y = 55.749

	punktyWykresuWielomianu[398].X = 3.98
	punktyWykresuWielomianu[398].Y = 56.164

	punktyWykresuWielomianu[399].X = 3.99
	punktyWykresuWielomianu[399].Y = 56.581

	punktyWykresuWielomianu[400].X = 4.0
	punktyWykresuWielomianu[400].Y = 57.0

	punktyWykresuWielomianu[401].X = 4.01
	punktyWykresuWielomianu[401].Y = 57.421

	punktyWykresuWielomianu[402].X = 4.02
	punktyWykresuWielomianu[402].Y = 57.844

	punktyWykresuWielomianu[403].X = 4.03
	punktyWykresuWielomianu[403].Y = 58.269

	punktyWykresuWielomianu[404].X = 4.04
	punktyWykresuWielomianu[404].Y = 58.697

	punktyWykresuWielomianu[405].X = 4.05
	punktyWykresuWielomianu[405].Y = 59.127

	punktyWykresuWielomianu[406].X = 4.06
	punktyWykresuWielomianu[406].Y = 59.559

	punktyWykresuWielomianu[407].X = 4.07
	punktyWykresuWielomianu[407].Y = 59.994

	punktyWykresuWielomianu[408].X = 4.08
	punktyWykresuWielomianu[408].Y = 60.43

	punktyWykresuWielomianu[409].X = 4.09
	punktyWykresuWielomianu[409].Y = 60.869

	punktyWykresuWielomianu[410].X = 4.1
	punktyWykresuWielomianu[410].Y = 61.311

	punktyWykresuWielomianu[411].X = 4.11
	punktyWykresuWielomianu[411].Y = 61.754

	punktyWykresuWielomianu[412].X = 4.12
	punktyWykresuWielomianu[412].Y = 62.2

	punktyWykresuWielomianu[413].X = 4.13
	punktyWykresuWielomianu[413].Y = 62.648

	punktyWykresuWielomianu[414].X = 4.14
	punktyWykresuWielomianu[414].Y = 63.098

	punktyWykresuWielomianu[415].X = 4.15
	punktyWykresuWielomianu[415].Y = 63.55

	punktyWykresuWielomianu[416].X = 4.16
	punktyWykresuWielomianu[416].Y = 64.005

	punktyWykresuWielomianu[417].X = 4.17
	punktyWykresuWielomianu[417].Y = 64.462

	punktyWykresuWielomianu[418].X = 4.18
	punktyWykresuWielomianu[418].Y = 64.922

	punktyWykresuWielomianu[419].X = 4.19
	punktyWykresuWielomianu[419].Y = 64.384

	punktyWykresuWielomianu[420].X = 4.2
	punktyWykresuWielomianu[420].Y = 65.848

	punktyWykresuWielomianu[421].X = 4.21
	punktyWykresuWielomianu[421].Y = 66.314

	punktyWykresuWielomianu[422].X = 4.22
	punktyWykresuWielomianu[422].Y = 66.783

	punktyWykresuWielomianu[423].X = 4.23
	punktyWykresuWielomianu[423].Y = 67.254

	punktyWykresuWielomianu[424].X = 4.24
	punktyWykresuWielomianu[424].Y = 67.727

	punktyWykresuWielomianu[425].X = 4.25
	punktyWykresuWielomianu[425].Y = 68.203

	punktyWykresuWielomianu[426].X = 4.26
	punktyWykresuWielomianu[426].Y = 68.681

	punktyWykresuWielomianu[427].X = 4.27
	punktyWykresuWielomianu[427].Y = 69.161

	punktyWykresuWielomianu[428].X = 4.28
	punktyWykresuWielomianu[428].Y = 69.644

	punktyWykresuWielomianu[429].X = 4.29
	punktyWykresuWielomianu[429].Y = 70.129

	punktyWykresuWielomianu[430].X = 4.3
	punktyWykresuWielomianu[430].Y = 70.617

	punktyWykresuWielomianu[431].X = 4.31
	punktyWykresuWielomianu[431].Y = 71.106

	punktyWykresuWielomianu[432].X = 4.32
	punktyWykresuWielomianu[432].Y = 71.599

	punktyWykresuWielomianu[433].X = 4.33
	punktyWykresuWielomianu[433].Y = 72.093

	punktyWykresuWielomianu[434].X = 4.34
	punktyWykresuWielomianu[434].Y = 72.59

	punktyWykresuWielomianu[435].X = 4.35
	punktyWykresuWielomianu[435].Y = 73.09

	punktyWykresuWielomianu[436].X = 4.36
	punktyWykresuWielomianu[436].Y = 72.592

	punktyWykresuWielomianu[437].X = 4.37
	punktyWykresuWielomianu[437].Y = 74.096

	punktyWykresuWielomianu[438].X = 4.38
	punktyWykresuWielomianu[438].Y = 74.603

	punktyWykresuWielomianu[439].X = 4.39
	punktyWykresuWielomianu[439].Y = 75.112

	punktyWykresuWielomianu[440].X = 4.4
	punktyWykresuWielomianu[440].Y = 75.624

	punktyWykresuWielomianu[441].X = 4.41
	punktyWykresuWielomianu[441].Y = 76.138

	punktyWykresuWielomianu[442].X = 4.42
	punktyWykresuWielomianu[442].Y = 76.654

	punktyWykresuWielomianu[443].X = 4.43
	punktyWykresuWielomianu[443].Y = 77.173

	punktyWykresuWielomianu[444].X = 4.44
	punktyWykresuWielomianu[444].Y = 77.694

	punktyWykresuWielomianu[445].X = 4.45
	punktyWykresuWielomianu[445].Y = 78.218

	punktyWykresuWielomianu[446].X = 4.46
	punktyWykresuWielomianu[446].Y = 78.744

	punktyWykresuWielomianu[447].X = 4.47
	punktyWykresuWielomianu[447].Y = 79.273

	punktyWykresuWielomianu[448].X = 4.48
	punktyWykresuWielomianu[448].Y = 79.805

	punktyWykresuWielomianu[449].X = 4.49
	punktyWykresuWielomianu[449].Y = 80.338

	punktyWykresuWielomianu[450].X = 4.5
	punktyWykresuWielomianu[450].Y = 80.875

	punktyWykresuWielomianu[451].X = 4.51
	punktyWykresuWielomianu[451].Y = 81.413

	punktyWykresuWielomianu[452].X = 4.52
	punktyWykresuWielomianu[452].Y = 81.955

	punktyWykresuWielomianu[453].X = 4.53
	punktyWykresuWielomianu[453].Y = 82.498

	punktyWykresuWielomianu[454].X = 4.54
	punktyWykresuWielomianu[454].Y = 83.045

	punktyWykresuWielomianu[455].X = 4.55
	punktyWykresuWielomianu[455].Y = 83.593

	punktyWykresuWielomianu[456].X = 4.56
	punktyWykresuWielomianu[456].Y = 84.145

	punktyWykresuWielomianu[457].X = 4.57
	punktyWykresuWielomianu[457].Y = 84.699

	punktyWykresuWielomianu[458].X = 4.58
	punktyWykresuWielomianu[458].Y = 85.255

	punktyWykresuWielomianu[459].X = 4.59
	punktyWykresuWielomianu[459].Y = 85.814

	punktyWykresuWielomianu[460].X = 4.6
	punktyWykresuWielomianu[460].Y = 86.376

	punktyWykresuWielomianu[461].X = 4.61
	punktyWykresuWielomianu[461].Y = 86.94

	punktyWykresuWielomianu[462].X = 4.62
	punktyWykresuWielomianu[462].Y = 87.506

	punktyWykresuWielomianu[463].X = 4.63
	punktyWykresuWielomianu[463].Y = 88.075

	punktyWykresuWielomianu[464].X = 4.64
	punktyWykresuWielomianu[464].Y = 88.647

	punktyWykresuWielomianu[465].X = 4.65
	punktyWykresuWielomianu[465].Y = 89.222

	punktyWykresuWielomianu[466].X = 4.66
	punktyWykresuWielomianu[466].Y = 89.799

	punktyWykresuWielomianu[467].X = 4.67
	punktyWykresuWielomianu[467].Y = 90.378

	punktyWykresuWielomianu[468].X = 4.68
	punktyWykresuWielomianu[468].Y = 90.96

	punktyWykresuWielomianu[469].X = 4.69
	punktyWykresuWielomianu[469].Y = 91.545

	punktyWykresuWielomianu[470].X = 4.7
	punktyWykresuWielomianu[470].Y = 92.133

	punktyWykresuWielomianu[471].X = 4.71
	punktyWykresuWielomianu[471].Y = 92.723

	punktyWykresuWielomianu[472].X = 4.72
	punktyWykresuWielomianu[472].Y = 93.315

	punktyWykresuWielomianu[473].X = 4.73
	punktyWykresuWielomianu[473].Y = 93.91

	punktyWykresuWielomianu[474].X = 4.74
	punktyWykresuWielomianu[474].Y = 94.508

	punktyWykresuWielomianu[475].X = 4.75
	punktyWykresuWielomianu[475].Y = 95.109

	punktyWykresuWielomianu[476].X = 4.76
	punktyWykresuWielomianu[476].Y = 95.712

	punktyWykresuWielomianu[477].X = 4.77
	punktyWykresuWielomianu[477].Y = 96.318

	punktyWykresuWielomianu[478].X = 4.78
	punktyWykresuWielomianu[478].Y = 96.927

	punktyWykresuWielomianu[479].X = 4.79
	punktyWykresuWielomianu[479].Y = 97.538

	punktyWykresuWielomianu[480].X = 4.8
	punktyWykresuWielomianu[480].Y = 98.152

	punktyWykresuWielomianu[481].X = 4.81
	punktyWykresuWielomianu[481].Y = 98.769

	punktyWykresuWielomianu[482].X = 4.82
	punktyWykresuWielomianu[482].Y = 99.387

	punktyWykresuWielomianu[483].X = 4.83
	punktyWykresuWielomianu[483].Y = 100.009

	punktyWykresuWielomianu[484].X = 4.84
	punktyWykresuWielomianu[484].Y = 100.634

	punktyWykresuWielomianu[485].X = 4.85
	punktyWykresuWielomianu[485].Y = 101.261

	punktyWykresuWielomianu[486].X = 4.86
	punktyWykresuWielomianu[486].Y = 101.891

	punktyWykresuWielomianu[487].X = 4.87
	punktyWykresuWielomianu[487].Y = 102.524

	punktyWykresuWielomianu[488].X = 4.88
	punktyWykresuWielomianu[488].Y = 103.159

	punktyWykresuWielomianu[489].X = 4.89
	punktyWykresuWielomianu[489].Y = 103.798

	punktyWykresuWielomianu[490].X = 4.9
	punktyWykresuWielomianu[490].Y = 104.439

	punktyWykresuWielomianu[491].X = 4.91
	punktyWykresuWielomianu[491].Y = 105.082

	punktyWykresuWielomianu[492].X = 4.92
	punktyWykresuWielomianu[492].Y = 105.729

	punktyWykresuWielomianu[493].X = 4.93
	punktyWykresuWielomianu[493].Y = 106.378

	punktyWykresuWielomianu[494].X = 4.94
	punktyWykresuWielomianu[494].Y = 107.03

	punktyWykresuWielomianu[495].X = 4.95
	punktyWykresuWielomianu[495].Y = 107.684

	punktyWykresuWielomianu[496].X = 4.96
	punktyWykresuWielomianu[496].Y = 108.342

	punktyWykresuWielomianu[497].X = 4.97
	punktyWykresuWielomianu[497].Y = 109.002

	punktyWykresuWielomianu[498].X = 4.98
	punktyWykresuWielomianu[498].Y = 109.665

	punktyWykresuWielomianu[499].X = 4.99
	punktyWykresuWielomianu[499].Y = 110.331

	punktyWykresuWielomianu[500].X = 5.0
	punktyWykresuWielomianu[500].Y = 111.0

	punktyWykresuWielomianu[501].X = 5.01
	punktyWykresuWielomianu[501].Y = 111.671

	punktyWykresuWielomianu[502].X = 5.02
	punktyWykresuWielomianu[502].Y = 112.345

	punktyWykresuWielomianu[503].X = 5.03
	punktyWykresuWielomianu[503].Y = 113.022

	punktyWykresuWielomianu[504].X = 5.04
	punktyWykresuWielomianu[504].Y = 113.702

	punktyWykresuWielomianu[505].X = 5.05
	punktyWykresuWielomianu[505].Y = 114.385

	punktyWykresuWielomianu[506].X = 5.06
	punktyWykresuWielomianu[506].Y = 115.07

	punktyWykresuWielomianu[507].X = 5.07
	punktyWykresuWielomianu[507].Y = 115.758

	punktyWykresuWielomianu[508].X = 5.08
	punktyWykresuWielomianu[508].Y = 116.45

	punktyWykresuWielomianu[509].X = 5.09
	punktyWykresuWielomianu[509].Y = 117.144

	punktyWykresuWielomianu[510].X = 5.1
	punktyWykresuWielomianu[510].Y = 117.841

	punktyWykresuWielomianu[511].X = 5.11
	punktyWykresuWielomianu[511].Y = 118.54

	punktyWykresuWielomianu[512].X = 5.12
	punktyWykresuWielomianu[512].Y = 119.243

	punktyWykresuWielomianu[513].X = 5.13
	punktyWykresuWielomianu[513].Y = 120.657

	punktyWykresuWielomianu[514].X = 5.14
	punktyWykresuWielomianu[514].Y = 120.657

	punktyWykresuWielomianu[515].X = 5.15
	punktyWykresuWielomianu[515].Y = 121.368

	punktyWykresuWielomianu[516].X = 5.16
	punktyWykresuWielomianu[516].Y = 122.082

	punktyWykresuWielomianu[517].X = 5.17
	punktyWykresuWielomianu[517].Y = 122.799

	punktyWykresuWielomianu[518].X = 5.18
	punktyWykresuWielomianu[518].Y = 123.519

	punktyWykresuWielomianu[519].X = 5.19
	punktyWykresuWielomianu[519].Y = 124.242

	punktyWykresuWielomianu[520].X = 5.2
	punktyWykresuWielomianu[520].Y = 124.968

	punktyWykresuWielomianu[521].X = 5.21
	punktyWykresuWielomianu[521].Y = 125.696

	punktyWykresuWielomianu[522].X = 5.22
	punktyWykresuWielomianu[522].Y = 126.64

	punktyWykresuWielomianu[523].X = 5.23
	punktyWykresuWielomianu[523].Y = 127.162

	punktyWykresuWielomianu[524].X = 5.24
	punktyWykresuWielomianu[524].Y = 127.9

	punktyWykresuWielomianu[525].X = 5.25
	punktyWykresuWielomianu[525].Y = 128.64

	punktyWykresuWielomianu[526].X = 5.26
	punktyWykresuWielomianu[526].Y = 129.631

	punktyWykresuWielomianu[527].X = 5.27
	punktyWykresuWielomianu[527].Y = 130.13

	punktyWykresuWielomianu[528].X = 5.28
	punktyWykresuWielomianu[528].Y = 130.879

	punktyWykresuWielomianu[529].X = 5.29
	punktyWykresuWielomianu[529].Y = 131.631

	punktyWykresuWielomianu[530].X = 5.3
	punktyWykresuWielomianu[530].Y = 132.387

	punktyWykresuWielomianu[531].X = 5.31
	punktyWykresuWielomianu[531].Y = 133.145

	punktyWykresuWielomianu[532].X = 5.32
	punktyWykresuWielomianu[532].Y = 133.67

	punktyWykresuWielomianu[533].X = 5.33
	punktyWykresuWielomianu[533].Y = 134.67

	punktyWykresuWielomianu[534].X = 5.34
	punktyWykresuWielomianu[534].Y = 135.437

	punktyWykresuWielomianu[535].X = 5.35
	punktyWykresuWielomianu[535].Y = 136.207

	punktyWykresuWielomianu[536].X = 5.36
	punktyWykresuWielomianu[536].Y = 136.981

	punktyWykresuWielomianu[537].X = 5.37
	punktyWykresuWielomianu[537].Y = 137.757

	punktyWykresuWielomianu[538].X = 5.38
	punktyWykresuWielomianu[538].Y = 138.536

	punktyWykresuWielomianu[539].X = 5.39
	punktyWykresuWielomianu[539].Y = 139.318

	punktyWykresuWielomianu[540].X = 5.4
	punktyWykresuWielomianu[540].Y = 140.104

	punktyWykresuWielomianu[541].X = 5.41
	punktyWykresuWielomianu[541].Y = 140.892

	punktyWykresuWielomianu[542].X = 5.42
	punktyWykresuWielomianu[542].Y = 141.683

	punktyWykresuWielomianu[543].X = 5.43
	punktyWykresuWielomianu[543].Y = 142.478

	punktyWykresuWielomianu[544].X = 5.44
	punktyWykresuWielomianu[544].Y = 143.275

	punktyWykresuWielomianu[545].X = 5.45
	punktyWykresuWielomianu[545].Y = 144.076

	punktyWykresuWielomianu[546].X = 5.46
	punktyWykresuWielomianu[546].Y = 144.879

	punktyWykresuWielomianu[547].X = 5.47
	punktyWykresuWielomianu[547].Y = 145.686

	punktyWykresuWielomianu[548].X = 5.48
	punktyWykresuWielomianu[548].Y = 146.496

	punktyWykresuWielomianu[549].X = 5.49
	punktyWykresuWielomianu[549].Y = 147.309

	punktyWykresuWielomianu[550].X = 5.5
	punktyWykresuWielomianu[550].Y = 148.125

	punktyWykresuWielomianu[551].X = 5.51
	punktyWykresuWielomianu[551].Y = 148.944

	punktyWykresuWielomianu[552].X = 5.52
	punktyWykresuWielomianu[552].Y = 149.766

	punktyWykresuWielomianu[553].X = 5.53
	punktyWykresuWielomianu[553].Y = 150.951

	punktyWykresuWielomianu[554].X = 5.54
	punktyWykresuWielomianu[554].Y = 151.419

	punktyWykresuWielomianu[555].X = 5.55
	punktyWykresuWielomianu[555].Y = 152.251

	punktyWykresuWielomianu[556].X = 5.56
	punktyWykresuWielomianu[556].Y = 153.086

	punktyWykresuWielomianu[557].X = 5.57
	punktyWykresuWielomianu[557].Y = 153.923

	punktyWykresuWielomianu[558].X = 5.58
	punktyWykresuWielomianu[558].Y = 154.764

	punktyWykresuWielomianu[559].X = 5.59
	punktyWykresuWielomianu[559].Y = 155.608

	punktyWykresuWielomianu[560].X = 5.6
	punktyWykresuWielomianu[560].Y = 156.456

	punktyWykresuWielomianu[561].X = 5.61
	punktyWykresuWielomianu[561].Y = 157.306

	punktyWykresuWielomianu[562].X = 5.62
	punktyWykresuWielomianu[562].Y = 158.159

	punktyWykresuWielomianu[563].X = 5.63
	punktyWykresuWielomianu[563].Y = 159.016

	punktyWykresuWielomianu[564].X = 5.64
	punktyWykresuWielomianu[564].Y = 159.876

	punktyWykresuWielomianu[565].X = 5.65
	punktyWykresuWielomianu[565].Y = 160.739

	punktyWykresuWielomianu[566].X = 5.66
	punktyWykresuWielomianu[566].Y = 161.605

	punktyWykresuWielomianu[567].X = 5.67
	punktyWykresuWielomianu[567].Y = 162.475

	punktyWykresuWielomianu[568].X = 5.68
	punktyWykresuWielomianu[568].Y = 163.348

	punktyWykresuWielomianu[569].X = 5.69
	punktyWykresuWielomianu[569].Y = 164.223

	punktyWykresuWielomianu[570].X = 5.7
	punktyWykresuWielomianu[570].Y = 165.103

	punktyWykresuWielomianu[571].X = 5.71
	punktyWykresuWielomianu[571].Y = 165.985

	punktyWykresuWielomianu[572].X = 5.72
	punktyWykresuWielomianu[572].Y = 166.87

	punktyWykresuWielomianu[573].X = 5.73
	punktyWykresuWielomianu[573].Y = 167.759

	punktyWykresuWielomianu[574].X = 5.74
	punktyWykresuWielomianu[574].Y = 168.651

	punktyWykresuWielomianu[575].X = 5.75
	punktyWykresuWielomianu[575].Y = 169.546

	punktyWykresuWielomianu[576].X = 5.76
	punktyWykresuWielomianu[576].Y = 170.445

	punktyWykresuWielomianu[577].X = 5.77
	punktyWykresuWielomianu[577].Y = 171.347

	punktyWykresuWielomianu[578].X = 5.78
	punktyWykresuWielomianu[578].Y = 172.252

	punktyWykresuWielomianu[579].X = 5.79
	punktyWykresuWielomianu[579].Y = 173.16

	punktyWykresuWielomianu[580].X = 5.8
	punktyWykresuWielomianu[580].Y = 174.072

	punktyWykresuWielomianu[581].X = 5.81
	punktyWykresuWielomianu[581].Y = 174.986

	punktyWykresuWielomianu[582].X = 5.82
	punktyWykresuWielomianu[582].Y = 175.905

	punktyWykresuWielomianu[583].X = 5.83
	punktyWykresuWielomianu[583].Y = 176.826

	punktyWykresuWielomianu[584].X = 5.84
	punktyWykresuWielomianu[584].Y = 177.751

	punktyWykresuWielomianu[585].X = 5.85
	punktyWykresuWielomianu[585].Y = 178.679

	punktyWykresuWielomianu[586].X = 5.86
	punktyWykresuWielomianu[586].Y = 179.61

	punktyWykresuWielomianu[587].X = 5.87
	punktyWykresuWielomianu[587].Y = 180.545

	punktyWykresuWielomianu[588].X = 5.88
	punktyWykresuWielomianu[588].Y = 181.483

	punktyWykresuWielomianu[589].X = 5.89
	punktyWykresuWielomianu[589].Y = 182.424

	punktyWykresuWielomianu[590].X = 5.9
	punktyWykresuWielomianu[590].Y = 183.369

	punktyWykresuWielomianu[591].X = 5.91
	punktyWykresuWielomianu[591].Y = 184.317

	punktyWykresuWielomianu[592].X = 5.92
	punktyWykresuWielomianu[592].Y = 185.268

	punktyWykresuWielomianu[593].X = 5.93
	punktyWykresuWielomianu[593].Y = 186.223

	punktyWykresuWielomianu[594].X = 5.94
	punktyWykresuWielomianu[594].Y = 187.181

	punktyWykresuWielomianu[595].X = 5.95
	punktyWykresuWielomianu[595].Y = 188.142

	punktyWykresuWielomianu[596].X = 5.96
	punktyWykresuWielomianu[596].Y = 189.107

	punktyWykresuWielomianu[597].X = 5.97
	punktyWykresuWielomianu[597].Y = 190.075

	punktyWykresuWielomianu[598].X = 5.98
	punktyWykresuWielomianu[598].Y = 191.046

	punktyWykresuWielomianu[599].X = 5.99
	punktyWykresuWielomianu[599].Y = 192.021

	punktyWykresuWielomianu[600].X = 6.0
	punktyWykresuWielomianu[600].Y = 193.0

	punktyWykresuWielomianu[601].X = 6.01
	punktyWykresuWielomianu[601].Y = 193.981

	punktyWykresuWielomianu[602].X = 6.02
	punktyWykresuWielomianu[602].Y = 194.966

	punktyWykresuWielomianu[603].X = 6.03
	punktyWykresuWielomianu[603].Y = 195.955

	punktyWykresuWielomianu[604].X = 6.04
	punktyWykresuWielomianu[604].Y = 196.947

	punktyWykresuWielomianu[605].X = 6.05
	punktyWykresuWielomianu[605].Y = 197.942

	punktyWykresuWielomianu[606].X = 6.06
	punktyWykresuWielomianu[606].Y = 198.941

	punktyWykresuWielomianu[607].X = 6.07
	punktyWykresuWielomianu[607].Y = 199.943

	punktyWykresuWielomianu[608].X = 6.08
	punktyWykresuWielomianu[608].Y = 200.949

	punktyWykresuWielomianu[609].X = 6.09
	punktyWykresuWielomianu[609].Y = 201.958

	punktyWykresuWielomianu[610].X = 6.1
	punktyWykresuWielomianu[610].Y = 202.971

	punktyWykresuWielomianu[611].X = 6.11
	punktyWykresuWielomianu[611].Y = 203.987

	punktyWykresuWielomianu[612].X = 6.12
	punktyWykresuWielomianu[612].Y = 205.006

	punktyWykresuWielomianu[613].X = 6.13
	punktyWykresuWielomianu[613].Y = 206.029

	punktyWykresuWielomianu[614].X = 6.14
	punktyWykresuWielomianu[614].Y = 207.055

	punktyWykresuWielomianu[615].X = 6.15
	punktyWykresuWielomianu[615].Y = 208.085

	punktyWykresuWielomianu[616].X = 6.16
	punktyWykresuWielomianu[616].Y = 209.119

	punktyWykresuWielomianu[617].X = 6.17
	punktyWykresuWielomianu[617].Y = 210.156

	punktyWykresuWielomianu[618].X = 6.18
	punktyWykresuWielomianu[618].Y = 211.196

	punktyWykresuWielomianu[619].X = 6.19
	punktyWykresuWielomianu[619].Y = 212.24

	punktyWykresuWielomianu[620].X = 6.2
	punktyWykresuWielomianu[620].Y = 213.288

	punktyWykresuWielomianu[621].X = 6.21
	punktyWykresuWielomianu[621].Y = 214.339

	punktyWykresuWielomianu[622].X = 6.22
	punktyWykresuWielomianu[622].Y = 215.393

	punktyWykresuWielomianu[623].X = 6.23
	punktyWykresuWielomianu[623].Y = 216.451

	punktyWykresuWielomianu[624].X = 6.24
	punktyWykresuWielomianu[624].Y = 217.513

	punktyWykresuWielomianu[625].X = 6.25
	punktyWykresuWielomianu[625].Y = 218.578

	punktyWykresuWielomianu[626].X = 6.26
	punktyWykresuWielomianu[626].Y = 219.646

	punktyWykresuWielomianu[627].X = 6.27
	punktyWykresuWielomianu[627].Y = 220.719

	punktyWykresuWielomianu[628].X = 6.28
	punktyWykresuWielomianu[628].Y = 221.794

	punktyWykresuWielomianu[629].X = 6.29
	punktyWykresuWielomianu[629].Y = 222.874

	punktyWykresuWielomianu[630].X = 6.3
	punktyWykresuWielomianu[630].Y = 223.957

	punktyWykresuWielomianu[631].X = 6.31
	punktyWykresuWielomianu[631].Y = 225.043

	punktyWykresuWielomianu[632].X = 6.32
	punktyWykresuWielomianu[632].Y = 226.133

	punktyWykresuWielomianu[633].X = 6.33
	punktyWykresuWielomianu[633].Y = 227.227

	punktyWykresuWielomianu[634].X = 6.33
	punktyWykresuWielomianu[634].Y = 228.324

	punktyWykresuWielomianu[635].X = 6.35
	punktyWykresuWielomianu[635].Y = 229.425

	punktyWykresuWielomianu[636].X = 6.36
	punktyWykresuWielomianu[636].Y = 230.529

	punktyWykresuWielomianu[637].X = 6.37
	punktyWykresuWielomianu[637].Y = 231.638

	punktyWykresuWielomianu[638].X = 6.38
	punktyWykresuWielomianu[638].Y = 232.749

	punktyWykresuWielomianu[639].X = 6.39
	punktyWykresuWielomianu[639].Y = 233.865

	punktyWykresuWielomianu[640].X = 6.4
	punktyWykresuWielomianu[640].Y = 234.984

	punktyWykresuWielomianu[641].X = 6.41
	punktyWykresuWielomianu[641].Y = 236.106

	punktyWykresuWielomianu[642].X = 6.42
	punktyWykresuWielomianu[642].Y = 237.232

	punktyWykresuWielomianu[643].X = 6.43
	punktyWykresuWielomianu[643].Y = 238.362

	punktyWykresuWielomianu[644].X = 6.44
	punktyWykresuWielomianu[644].Y = 239.496

	punktyWykresuWielomianu[645].X = 6.45
	punktyWykresuWielomianu[645].Y = 240.633

	punktyWykresuWielomianu[646].X = 6.46
	punktyWykresuWielomianu[646].Y = 241.774

	punktyWykresuWielomianu[647].X = 6.47
	punktyWykresuWielomianu[647].Y = 242.919

	punktyWykresuWielomianu[648].X = 6.48
	punktyWykresuWielomianu[648].Y = 244.067

	punktyWykresuWielomianu[649].X = 6.49
	punktyWykresuWielomianu[649].Y = 245.219

	punktyWykresuWielomianu[650].X = 6.5
	punktyWykresuWielomianu[650].Y = 246.375

	punktyWykresuWielomianu[651].X = 6.51
	punktyWykresuWielomianu[651].Y = 247.534

	punktyWykresuWielomianu[652].X = 6.52
	punktyWykresuWielomianu[652].Y = 248.697

	punktyWykresuWielomianu[653].X = 6.53
	punktyWykresuWielomianu[653].Y = 249.864

	punktyWykresuWielomianu[654].X = 6.54
	punktyWykresuWielomianu[654].Y = 251.034

	punktyWykresuWielomianu[655].X = 6.55
	punktyWykresuWielomianu[655].Y = 252.208

	punktyWykresuWielomianu[656].X = 6.56
	punktyWykresuWielomianu[656].Y = 253.386

	punktyWykresuWielomianu[657].X = 6.57
	punktyWykresuWielomianu[657].Y = 254.568

	punktyWykresuWielomianu[658].X = 6.58
	punktyWykresuWielomianu[658].Y = 255.753

	punktyWykresuWielomianu[659].X = 6.59
	punktyWykresuWielomianu[659].Y = 256.943

	punktyWykresuWielomianu[660].X = 6.6
	punktyWykresuWielomianu[660].Y = 258.136

	punktyWykresuWielomianu[661].X = 6.61
	punktyWykresuWielomianu[661].Y = 259.332

	punktyWykresuWielomianu[662].X = 6.62
	punktyWykresuWielomianu[662].Y = 260.533

	punktyWykresuWielomianu[663].X = 6.63
	punktyWykresuWielomianu[663].Y = 261.737

	punktyWykresuWielomianu[664].X = 6.64
	punktyWykresuWielomianu[664].Y = 262.945

	punktyWykresuWielomianu[665].X = 6.65
	punktyWykresuWielomianu[665].Y = 264.157

	punktyWykresuWielomianu[666].X = 6.66
	punktyWykresuWielomianu[666].Y = 265.372

	punktyWykresuWielomianu[667].X = 6.67
	punktyWykresuWielomianu[667].Y = 266.592

	punktyWykresuWielomianu[668].X = 6.68
	punktyWykresuWielomianu[668].Y = 267.815

	punktyWykresuWielomianu[669].X = 6.69
	punktyWykresuWielomianu[669].Y = 269.042

	punktyWykresuWielomianu[670].X = 6.7
	punktyWykresuWielomianu[670].Y = 270.273

	punktyWykresuWielomianu[671].X = 6.71
	punktyWykresuWielomianu[671].Y = 271.507

	punktyWykresuWielomianu[672].X = 6.72
	punktyWykresuWielomianu[672].Y = 272.746

	punktyWykresuWielomianu[673].X = 6.73
	punktyWykresuWielomianu[673].Y = 273.988

	punktyWykresuWielomianu[674].X = 6.74
	punktyWykresuWielomianu[674].Y = 275.234

	punktyWykresuWielomianu[675].X = 6.75
	punktyWykresuWielomianu[675].Y = 276.484

	punktyWykresuWielomianu[676].X = 6.76
	punktyWykresuWielomianu[676].Y = 277.738

	punktyWykresuWielomianu[677].X = 6.77
	punktyWykresuWielomianu[677].Y = 278.995

	punktyWykresuWielomianu[678].X = 6.78
	punktyWykresuWielomianu[678].Y = 280.257

	punktyWykresuWielomianu[679].X = 6.79
	punktyWykresuWielomianu[679].Y = 281.522

	punktyWykresuWielomianu[680].X = 6.8
	punktyWykresuWielomianu[680].Y = 282.792

	punktyWykresuWielomianu[681].X = 6.81
	punktyWykresuWielomianu[681].Y = 284.065

	punktyWykresuWielomianu[682].X = 6.82
	punktyWykresuWielomianu[682].Y = 285.342

	punktyWykresuWielomianu[683].X = 6.83
	punktyWykresuWielomianu[683].Y = 286.623

	punktyWykresuWielomianu[684].X = 6.84
	punktyWykresuWielomianu[684].Y = 287.907

	punktyWykresuWielomianu[685].X = 6.85
	punktyWykresuWielomianu[685].Y = 289.196

	punktyWykresuWielomianu[686].X = 6.86
	punktyWykresuWielomianu[686].Y = 290.489

	punktyWykresuWielomianu[687].X = 6.87
	punktyWykresuWielomianu[687].Y = 291.785

	punktyWykresuWielomianu[688].X = 6.88
	punktyWykresuWielomianu[688].Y = 293.086

	punktyWykresuWielomianu[689].X = 6.89
	punktyWykresuWielomianu[689].Y = 294.39

	punktyWykresuWielomianu[690].X = 6.9
	punktyWykresuWielomianu[690].Y = 295.699

	punktyWykresuWielomianu[691].X = 6.91
	punktyWykresuWielomianu[691].Y = 297.011

	punktyWykresuWielomianu[692].X = 6.92
	punktyWykresuWielomianu[692].Y = 298.327

	punktyWykresuWielomianu[693].X = 6.93
	punktyWykresuWielomianu[693].Y = 299.647

	punktyWykresuWielomianu[694].X = 6.94
	punktyWykresuWielomianu[694].Y = 300.971

	punktyWykresuWielomianu[695].X = 6.95
	punktyWykresuWielomianu[695].Y = 302.299

	punktyWykresuWielomianu[696].X = 6.96
	punktyWykresuWielomianu[696].Y = 303.631

	punktyWykresuWielomianu[697].X = 6.97
	punktyWykresuWielomianu[697].Y = 304.968

	punktyWykresuWielomianu[698].X = 6.98
	punktyWykresuWielomianu[698].Y = 306.308

	punktyWykresuWielomianu[699].X = 6.99
	punktyWykresuWielomianu[699].Y = 307.652

	punktyWykresuWielomianu[700].X = 7.0
	punktyWykresuWielomianu[700].Y = 309.0

	punktyWykresuWielomianu[701].X = 7.01
	punktyWykresuWielomianu[701].Y = 310.352

	punktyWykresuWielomianu[702].X = 7.02
	punktyWykresuWielomianu[702].Y = 311.708

	punktyWykresuWielomianu[703].X = 7.03
	punktyWykresuWielomianu[703].Y = 313.068

	punktyWykresuWielomianu[704].X = 7.04
	punktyWykresuWielomianu[704].Y = 314.432

	punktyWykresuWielomianu[705].X = 7.05
	punktyWykresuWielomianu[705].Y = 315.8

	punktyWykresuWielomianu[706].X = 7.06
	punktyWykresuWielomianu[706].Y = 317.172

	punktyWykresuWielomianu[707].X = 7.07
	punktyWykresuWielomianu[707].Y = 318.548

	punktyWykresuWielomianu[708].X = 7.08
	punktyWykresuWielomianu[708].Y = 319.928

	punktyWykresuWielomianu[709].X = 7.09
	punktyWykresuWielomianu[709].Y = 321.312

	punktyWykresuWielomianu[710].X = 7.1
	punktyWykresuWielomianu[710].Y = 322.701

	punktyWykresuWielomianu[711].X = 7.11
	punktyWykresuWielomianu[711].Y = 324.093

	punktyWykresuWielomianu[712].X = 7.12
	punktyWykresuWielomianu[712].Y = 325.489

	punktyWykresuWielomianu[713].X = 7.13
	punktyWykresuWielomianu[713].Y = 326.89

	punktyWykresuWielomianu[714].X = 7.14
	punktyWykresuWielomianu[714].Y = 328.294

	punktyWykresuWielomianu[715].X = 7.15
	punktyWykresuWielomianu[715].Y = 329.703

	punktyWykresuWielomianu[716].X = 7.16
	punktyWykresuWielomianu[716].Y = 331.116

	punktyWykresuWielomianu[717].X = 7.17
	punktyWykresuWielomianu[717].Y = 332.532

	punktyWykresuWielomianu[718].X = 7.18
	punktyWykresuWielomianu[718].Y = 333.953

	punktyWykresuWielomianu[719].X = 7.19
	punktyWykresuWielomianu[719].Y = 335.378

	punktyWykresuWielomianu[720].X = 7.2
	punktyWykresuWielomianu[720].Y = 336.808

	punktyWykresuWielomianu[721].X = 7.21
	punktyWykresuWielomianu[721].Y = 338.241

	punktyWykresuWielomianu[722].X = 7.22
	punktyWykresuWielomianu[722].Y = 339.678

	punktyWykresuWielomianu[723].X = 7.23
	punktyWykresuWielomianu[723].Y = 341.12

	punktyWykresuWielomianu[724].X = 7.24
	punktyWykresuWielomianu[724].Y = 342.565

	punktyWykresuWielomianu[725].X = 7.25
	punktyWykresuWielomianu[725].Y = 344.015

	punktyWykresuWielomianu[726].X = 7.26
	punktyWykresuWielomianu[726].Y = 345.469

	punktyWykresuWielomianu[727].X = 7.27
	punktyWykresuWielomianu[727].Y = 346.927

	punktyWykresuWielomianu[728].X = 7.28
	punktyWykresuWielomianu[728].Y = 348.39

	punktyWykresuWielomianu[729].X = 7.29
	punktyWykresuWielomianu[729].Y = 349.856

	punktyWykresuWielomianu[730].X = 7.3
	punktyWykresuWielomianu[730].Y = 351.327

	punktyWykresuWielomianu[731].X = 7.31
	punktyWykresuWielomianu[731].Y = 352.801

	punktyWykresuWielomianu[732].X = 7.32
	punktyWykresuWielomianu[732].Y = 354.28

	punktyWykresuWielomianu[733].X = 7.33
	punktyWykresuWielomianu[733].Y = 355.763

	punktyWykresuWielomianu[734].X = 7.34
	punktyWykresuWielomianu[734].Y = 357.251

	punktyWykresuWielomianu[735].X = 7.35
	punktyWykresuWielomianu[735].Y = 358.742

	punktyWykresuWielomianu[736].X = 7.36
	punktyWykresuWielomianu[736].Y = 360.238

	punktyWykresuWielomianu[737].X = 7.37
	punktyWykresuWielomianu[737].Y = 361.738

	punktyWykresuWielomianu[738].X = 7.38
	punktyWykresuWielomianu[738].Y = 363.242

	punktyWykresuWielomianu[739].X = 7.39
	punktyWykresuWielomianu[739].Y = 364.751

	punktyWykresuWielomianu[740].X = 7.4
	punktyWykresuWielomianu[740].Y = 366.264

	punktyWykresuWielomianu[741].X = 7.41
	punktyWykresuWielomianu[741].Y = 367.78

	punktyWykresuWielomianu[742].X = 7.42
	punktyWykresuWielomianu[742].Y = 369.302

	punktyWykresuWielomianu[743].X = 7.43
	punktyWykresuWielomianu[743].Y = 370.827

	punktyWykresuWielomianu[744].X = 7.44
	punktyWykresuWielomianu[744].Y = 372.357

	punktyWykresuWielomianu[745].X = 7.45
	punktyWykresuWielomianu[745].Y = 373.891

	punktyWykresuWielomianu[746].X = 7.46
	punktyWykresuWielomianu[746].Y = 375.429

	punktyWykresuWielomianu[747].X = 7.47
	punktyWykresuWielomianu[747].Y = 376.971

	punktyWykresuWielomianu[748].X = 7.48
	punktyWykresuWielomianu[748].Y = 378.518

	punktyWykresuWielomianu[749].X = 7.49
	punktyWykresuWielomianu[749].Y = 380.069

	punktyWykresuWielomianu[750].X = 7.5
	punktyWykresuWielomianu[750].Y = 381.625

	punktyWykresuWielomianu[751].X = 7.51
	punktyWykresuWielomianu[751].Y = 383.184

	punktyWykresuWielomianu[752].X = 7.52
	punktyWykresuWielomianu[752].Y = 384.748

	punktyWykresuWielomianu[753].X = 7.53
	punktyWykresuWielomianu[753].Y = 386.316

	punktyWykresuWielomianu[754].X = 7.54
	punktyWykresuWielomianu[754].Y = 387.889

	punktyWykresuWielomianu[755].X = 7.55
	punktyWykresuWielomianu[755].Y = 389.466

	punktyWykresuWielomianu[756].X = 7.56
	punktyWykresuWielomianu[756].Y = 391.047

	punktyWykresuWielomianu[757].X = 7.57
	punktyWykresuWielomianu[757].Y = 392.633

	punktyWykresuWielomianu[758].X = 7.58
	punktyWykresuWielomianu[758].Y = 394.223

	punktyWykresuWielomianu[759].X = 7.59
	punktyWykresuWielomianu[759].Y = 395.817

	punktyWykresuWielomianu[760].X = 7.6
	punktyWykresuWielomianu[760].Y = 397.416

	punktyWykresuWielomianu[761].X = 7.61
	punktyWykresuWielomianu[761].Y = 399.019

	punktyWykresuWielomianu[762].X = 7.62
	punktyWykresuWielomianu[762].Y = 400.626

	punktyWykresuWielomianu[763].X = 7.63
	punktyWykresuWielomianu[763].Y = 402.238

	punktyWykresuWielomianu[764].X = 7.64
	punktyWykresuWielomianu[764].Y = 403.854

	punktyWykresuWielomianu[765].X = 7.65
	punktyWykresuWielomianu[765].Y = 405.474

	punktyWykresuWielomianu[766].X = 7.66
	punktyWykresuWielomianu[766].Y = 407.099

	punktyWykresuWielomianu[767].X = 7.67
	punktyWykresuWielomianu[767].Y = 408.728

	punktyWykresuWielomianu[768].X = 7.68
	punktyWykresuWielomianu[768].Y = 410.362

	punktyWykresuWielomianu[769].X = 7.69
	punktyWykresuWielomianu[769].Y = 412.0

	punktyWykresuWielomianu[770].X = 7.7
	punktyWykresuWielomianu[770].Y = 413.643

	punktyWykresuWielomianu[771].X = 7.71
	punktyWykresuWielomianu[771].Y = 415.289

	punktyWykresuWielomianu[772].X = 7.72
	punktyWykresuWielomianu[772].Y = 416.941

	punktyWykresuWielomianu[773].X = 7.73
	punktyWykresuWielomianu[773].Y = 418.597

	punktyWykresuWielomianu[774].X = 7.74
	punktyWykresuWielomianu[774].Y = 420.257

	punktyWykresuWielomianu[775].X = 7.75
	punktyWykresuWielomianu[775].Y = 421.921

	punktyWykresuWielomianu[776].X = 7.76
	punktyWykresuWielomianu[776].Y = 423.591

	punktyWykresuWielomianu[777].X = 7.77
	punktyWykresuWielomianu[777].Y = 425.264

	punktyWykresuWielomianu[778].X = 7.78
	punktyWykresuWielomianu[778].Y = 426.942

	punktyWykresuWielomianu[779].X = 7.79
	punktyWykresuWielomianu[779].Y = 428.625

	punktyWykresuWielomianu[780].X = 7.8
	punktyWykresuWielomianu[780].Y = 430.312

	punktyWykresuWielomianu[781].X = 7.81
	punktyWykresuWielomianu[781].Y = 432.003

	punktyWykresuWielomianu[782].X = 7.82
	punktyWykresuWielomianu[782].Y = 433.699

	punktyWykresuWielomianu[783].X = 7.83
	punktyWykresuWielomianu[783].Y = 435.399

	punktyWykresuWielomianu[784].X = 7.84
	punktyWykresuWielomianu[784].Y = 437.104

	punktyWykresuWielomianu[785].X = 7.85
	punktyWykresuWielomianu[785].Y = 438.814

	punktyWykresuWielomianu[786].X = 7.86
	punktyWykresuWielomianu[786].Y = 440.528

	punktyWykresuWielomianu[787].X = 7.87
	punktyWykresuWielomianu[787].Y = 442.246

	punktyWykresuWielomianu[788].X = 7.88
	punktyWykresuWielomianu[788].Y = 443.969

	punktyWykresuWielomianu[789].X = 7.89
	punktyWykresuWielomianu[789].Y = 445.697

	punktyWykresuWielomianu[790].X = 7.9
	punktyWykresuWielomianu[790].Y = 447.429

	punktyWykresuWielomianu[791].X = 7.91
	punktyWykresuWielomianu[791].Y = 449.165

	punktyWykresuWielomianu[792].X = 7.92
	punktyWykresuWielomianu[792].Y = 450.906

	punktyWykresuWielomianu[793].X = 7.93
	punktyWykresuWielomianu[793].Y = 452.652

	punktyWykresuWielomianu[794].X = 7.94
	punktyWykresuWielomianu[794].Y = 454.402

	punktyWykresuWielomianu[795].X = 7.95
	punktyWykresuWielomianu[795].Y = 456.157

	punktyWykresuWielomianu[796].X = 7.96
	punktyWykresuWielomianu[796].Y = 457.916

	punktyWykresuWielomianu[797].X = 7.97
	punktyWykresuWielomianu[797].Y = 459.68

	punktyWykresuWielomianu[798].X = 7.98
	punktyWykresuWielomianu[798].Y = 461.449

	punktyWykresuWielomianu[799].X = 7.99
	punktyWykresuWielomianu[799].Y = 463.222

	punktyWykresuWielomianu[800].X = 8.0
	punktyWykresuWielomianu[800].Y = 465.0

	punktyWykresuWielomianu[801].X = 8.01
	punktyWykresuWielomianu[801].Y = 466.782

	punktyWykresuWielomianu[802].X = 8.02
	punktyWykresuWielomianu[802].Y = 468.569

	punktyWykresuWielomianu[803].X = 8.03
	punktyWykresuWielomianu[803].Y = 470.36

	punktyWykresuWielomianu[804].X = 8.04
	punktyWykresuWielomianu[804].Y = 472.156

	punktyWykresuWielomianu[805].X = 8.05
	punktyWykresuWielomianu[805].Y = 473.957

	punktyWykresuWielomianu[806].X = 8.06
	punktyWykresuWielomianu[806].Y = 475.763

	punktyWykresuWielomianu[807].X = 8.07
	punktyWykresuWielomianu[807].Y = 477.573

	punktyWykresuWielomianu[808].X = 8.08
	punktyWykresuWielomianu[808].Y = 479.387

	punktyWykresuWielomianu[809].X = 8.09
	punktyWykresuWielomianu[809].Y = 481.207

	punktyWykresuWielomianu[810].X = 8.1
	punktyWykresuWielomianu[810].Y = 483.031

	punktyWykresuWielomianu[811].X = 8.11
	punktyWykresuWielomianu[811].Y = 484.859

	punktyWykresuWielomianu[812].X = 8.12
	punktyWykresuWielomianu[812].Y = 486.859

	punktyWykresuWielomianu[813].X = 8.13
	punktyWykresuWielomianu[813].Y = 488.53

	punktyWykresuWielomianu[814].X = 8.14
	punktyWykresuWielomianu[814].Y = 490.373

	punktyWykresuWielomianu[815].X = 8.15
	punktyWykresuWielomianu[815].Y = 492.22

	punktyWykresuWielomianu[816].X = 8.16
	punktyWykresuWielomianu[816].Y = 494.072

	punktyWykresuWielomianu[817].X = 8.17
	punktyWykresuWielomianu[817].Y = 495.929

	punktyWykresuWielomianu[818].X = 8.18
	punktyWykresuWielomianu[818].Y = 497.791

	punktyWykresuWielomianu[819].X = 8.19
	punktyWykresuWielomianu[819].Y = 499.657

	punktyWykresuWielomianu[820].X = 8.2
	punktyWykresuWielomianu[820].Y = 501.528

	punktyWykresuWielomianu[821].X = 8.21
	punktyWykresuWielomianu[821].Y = 503.403

	punktyWykresuWielomianu[822].X = 8.22
	punktyWykresuWielomianu[822].Y = 505.282

	punktyWykresuWielomianu[823].X = 8.23
	punktyWykresuWielomianu[823].Y = 507.168

	punktyWykresuWielomianu[824].X = 8.24
	punktyWykresuWielomianu[824].Y = 509.058

	punktyWykresuWielomianu[825].X = 8.25
	punktyWykresuWielomianu[825].Y = 510.953

	punktyWykresuWielomianu[826].X = 8.26
	punktyWykresuWielomianu[826].Y = 512.852

	punktyWykresuWielomianu[827].X = 8.27
	punktyWykresuWielomianu[827].Y = 514.756

	punktyWykresuWielomianu[828].X = 8.28
	punktyWykresuWielomianu[828].Y = 516.665

	punktyWykresuWielomianu[829].X = 8.29
	punktyWykresuWielomianu[829].Y = 518.578

	punktyWykresuWielomianu[830].X = 8.3
	punktyWykresuWielomianu[830].Y = 520.497

	punktyWykresuWielomianu[831].X = 8.31
	punktyWykresuWielomianu[831].Y = 522.42

	punktyWykresuWielomianu[832].X = 8.32
	punktyWykresuWielomianu[832].Y = 524.349

	punktyWykresuWielomianu[833].X = 8.33
	punktyWykresuWielomianu[833].Y = 526.28

	punktyWykresuWielomianu[834].X = 8.34
	punktyWykresuWielomianu[834].Y = 528.218

	punktyWykresuWielomianu[835].X = 8.35
	punktyWykresuWielomianu[835].Y = 530.16

	punktyWykresuWielomianu[836].X = 8.36
	punktyWykresuWielomianu[836].Y = 532.107

	punktyWykresuWielomianu[837].X = 8.37
	punktyWykresuWielomianu[837].Y = 534.218

	punktyWykresuWielomianu[838].X = 8.38
	punktyWykresuWielomianu[838].Y = 536.016

	punktyWykresuWielomianu[839].X = 8.39
	punktyWykresuWielomianu[839].Y = 537.977

	punktyWykresuWielomianu[840].X = 8.4
	punktyWykresuWielomianu[840].Y = 539.944

	punktyWykresuWielomianu[841].X = 8.41
	punktyWykresuWielomianu[841].Y = 541.915

	punktyWykresuWielomianu[842].X = 8.42
	punktyWykresuWielomianu[842].Y = 543.891

	punktyWykresuWielomianu[843].X = 8.43
	punktyWykresuWielomianu[843].Y = 545.872

	punktyWykresuWielomianu[844].X = 8.44
	punktyWykresuWielomianu[844].Y = 547.858

	punktyWykresuWielomianu[845].X = 8.45
	punktyWykresuWielomianu[845].Y = 549.848

	punktyWykresuWielomianu[846].X = 8.46
	punktyWykresuWielomianu[846].Y = 551.844

	punktyWykresuWielomianu[847].X = 8.47
	punktyWykresuWielomianu[847].Y = 553.844

	punktyWykresuWielomianu[848].X = 8.48
	punktyWykresuWielomianu[848].Y = 555.849

	punktyWykresuWielomianu[849].X = 8.49
	punktyWykresuWielomianu[849].Y = 557.859

	punktyWykresuWielomianu[850].X = 8.5
	punktyWykresuWielomianu[850].Y = 559.875

	punktyWykresuWielomianu[851].X = 8.51
	punktyWykresuWielomianu[851].Y = 561.895

	punktyWykresuWielomianu[852].X = 8.52
	punktyWykresuWielomianu[852].Y = 563.919

	punktyWykresuWielomianu[853].X = 8.53
	punktyWykresuWielomianu[853].Y = 565.949

	punktyWykresuWielomianu[854].X = 8.54
	punktyWykresuWielomianu[854].Y = 567.984

	punktyWykresuWielomianu[855].X = 8.55
	punktyWykresuWielomianu[855].Y = 570.023

	punktyWykresuWielomianu[856].X = 8.56
	punktyWykresuWielomianu[856].Y = 572.068

	punktyWykresuWielomianu[857].X = 8.57
	punktyWykresuWielomianu[857].Y = 574.117

	punktyWykresuWielomianu[858].X = 8.58
	punktyWykresuWielomianu[858].Y = 576.172

	punktyWykresuWielomianu[859].X = 8.59
	punktyWykresuWielomianu[859].Y = 578.231

	punktyWykresuWielomianu[860].X = 8.6
	punktyWykresuWielomianu[860].Y = 580.296

	punktyWykresuWielomianu[861].X = 8.61
	punktyWykresuWielomianu[861].Y = 582.365

	punktyWykresuWielomianu[862].X = 8.62
	punktyWykresuWielomianu[862].Y = 584.439

	punktyWykresuWielomianu[863].X = 8.63
	punktyWykresuWielomianu[863].Y = 586.518

	punktyWykresuWielomianu[864].X = 8.64
	punktyWykresuWielomianu[864].Y = 588.602

	punktyWykresuWielomianu[865].X = 8.65
	punktyWykresuWielomianu[865].Y = 590.692

	punktyWykresuWielomianu[866].X = 8.66
	punktyWykresuWielomianu[866].Y = 592.786

	punktyWykresuWielomianu[867].X = 8.67
	punktyWykresuWielomianu[867].Y = 594.885

	punktyWykresuWielomianu[868].X = 8.68
	punktyWykresuWielomianu[868].Y = 596.989

	punktyWykresuWielomianu[869].X = 8.69
	punktyWykresuWielomianu[869].Y = 599.098

	punktyWykresuWielomianu[870].X = 8.7
	punktyWykresuWielomianu[870].Y = 601.213

	punktyWykresuWielomianu[871].X = 8.71
	punktyWykresuWielomianu[871].Y = 603.332

	punktyWykresuWielomianu[872].X = 8.72
	punktyWykresuWielomianu[872].Y = 605.456

	punktyWykresuWielomianu[873].X = 8.73
	punktyWykresuWielomianu[873].Y = 607.585

	punktyWykresuWielomianu[874].X = 8.74
	punktyWykresuWielomianu[874].Y = 609.72

	punktyWykresuWielomianu[875].X = 8.75
	punktyWykresuWielomianu[875].Y = 611.859

	punktyWykresuWielomianu[876].X = 8.76
	punktyWykresuWielomianu[876].Y = 614.003

	punktyWykresuWielomianu[877].X = 8.77
	punktyWykresuWielomianu[877].Y = 616.153

	punktyWykresuWielomianu[878].X = 8.78
	punktyWykresuWielomianu[878].Y = 618.307

	punktyWykresuWielomianu[879].X = 8.79
	punktyWykresuWielomianu[879].Y = 620.467

	punktyWykresuWielomianu[880].X = 8.8
	punktyWykresuWielomianu[880].Y = 622.632

	punktyWykresuWielomianu[881].X = 8.81
	punktyWykresuWielomianu[881].Y = 624.801

	punktyWykresuWielomianu[882].X = 8.82
	punktyWykresuWielomianu[882].Y = 626.976

	punktyWykresuWielomianu[883].X = 8.83
	punktyWykresuWielomianu[883].Y = 629.156

	punktyWykresuWielomianu[884].X = 8.84
	punktyWykresuWielomianu[884].Y = 631.341

	punktyWykresuWielomianu[885].X = 8.85
	punktyWykresuWielomianu[885].Y = 633.531

	punktyWykresuWielomianu[886].X = 8.86
	punktyWykresuWielomianu[886].Y = 635.726

	punktyWykresuWielomianu[887].X = 8.87
	punktyWykresuWielomianu[887].Y = 637.927

	punktyWykresuWielomianu[888].X = 8.88
	punktyWykresuWielomianu[888].Y = 640.132

	punktyWykresuWielomianu[889].X = 8.89
	punktyWykresuWielomianu[889].Y = 642.343

	punktyWykresuWielomianu[890].X = 8.9
	punktyWykresuWielomianu[890].Y = 644.559

	punktyWykresuWielomianu[891].X = 8.91
	punktyWykresuWielomianu[891].Y = 646.779

	punktyWykresuWielomianu[892].X = 8.92
	punktyWykresuWielomianu[892].Y = 649.005

	punktyWykresuWielomianu[893].X = 8.93
	punktyWykresuWielomianu[893].Y = 651.237

	punktyWykresuWielomianu[894].X = 8.94
	punktyWykresuWielomianu[894].Y = 653.473

	punktyWykresuWielomianu[895].X = 8.95
	punktyWykresuWielomianu[895].Y = 655.714

	punktyWykresuWielomianu[896].X = 8.96
	punktyWykresuWielomianu[896].Y = 657.961

	punktyWykresuWielomianu[897].X = 8.97
	punktyWykresuWielomianu[897].Y = 660.213

	punktyWykresuWielomianu[898].X = 8.98
	punktyWykresuWielomianu[898].Y = 662.47

	punktyWykresuWielomianu[899].X = 8.99
	punktyWykresuWielomianu[899].Y = 664.732

	punktyWykresuWielomianu[900].X = 9.0
	punktyWykresuWielomianu[900].Y = 667.0

	punktyWykresuWielomianu[901].X = 9.01
	punktyWykresuWielomianu[901].Y = 669.272

	punktyWykresuWielomianu[902].X = 9.02
	punktyWykresuWielomianu[902].Y = 671.55

	punktyWykresuWielomianu[903].X = 9.03
	punktyWykresuWielomianu[903].Y = 673.833

	punktyWykresuWielomianu[904].X = 9.04
	punktyWykresuWielomianu[904].Y = 676.121

	punktyWykresuWielomianu[905].X = 9.05
	punktyWykresuWielomianu[905].Y = 678.415

	punktyWykresuWielomianu[906].X = 9.06
	punktyWykresuWielomianu[906].Y = 680.713

	punktyWykresuWielomianu[907].X = 9.07
	punktyWykresuWielomianu[907].Y = 683.017

	punktyWykresuWielomianu[908].X = 9.08
	punktyWykresuWielomianu[908].Y = 685.326

	punktyWykresuWielomianu[909].X = 9.09
	punktyWykresuWielomianu[909].Y = 687.641

	punktyWykresuWielomianu[910].X = 9.1
	punktyWykresuWielomianu[910].Y = 689.961

	punktyWykresuWielomianu[911].X = 9.11
	punktyWykresuWielomianu[911].Y = 692.285

	punktyWykresuWielomianu[912].X = 9.12
	punktyWykresuWielomianu[912].Y = 694.616

	punktyWykresuWielomianu[913].X = 9.13
	punktyWykresuWielomianu[913].Y = 696.951

	punktyWykresuWielomianu[914].X = 9.14
	punktyWykresuWielomianu[914].Y = 699.292

	punktyWykresuWielomianu[915].X = 9.15
	punktyWykresuWielomianu[915].Y = 701.638

	punktyWykresuWielomianu[916].X = 9.16
	punktyWykresuWielomianu[916].Y = 703.989

	punktyWykresuWielomianu[917].X = 9.17
	punktyWykresuWielomianu[917].Y = 706.346

	punktyWykresuWielomianu[918].X = 9.18
	punktyWykresuWielomianu[918].Y = 708.708

	punktyWykresuWielomianu[919].X = 9.19
	punktyWykresuWielomianu[919].Y = 711.075

	punktyWykresuWielomianu[920].X = 9.2
	punktyWykresuWielomianu[920].Y = 713.448

	punktyWykresuWielomianu[921].X = 9.21
	punktyWykresuWielomianu[921].Y = 715.825

	punktyWykresuWielomianu[922].X = 9.22
	punktyWykresuWielomianu[922].Y = 718.209

	punktyWykresuWielomianu[923].X = 9.23
	punktyWykresuWielomianu[923].Y = 720.597

	punktyWykresuWielomianu[924].X = 9.24
	punktyWykresuWielomianu[924].Y = 722.991

	punktyWykresuWielomianu[925].X = 9.25
	punktyWykresuWielomianu[925].Y = 725.39

	punktyWykresuWielomianu[926].X = 9.26
	punktyWykresuWielomianu[926].Y = 727.795

	punktyWykresuWielomianu[927].X = 9.27
	punktyWykresuWielomianu[927].Y = 730.205

	punktyWykresuWielomianu[928].X = 9.28
	punktyWykresuWielomianu[928].Y = 732.62

	punktyWykresuWielomianu[929].X = 9.29
	punktyWykresuWielomianu[929].Y = 735.041

	punktyWykresuWielomianu[930].X = 9.3
	punktyWykresuWielomianu[930].Y = 737.467

	punktyWykresuWielomianu[931].X = 9.31
	punktyWykresuWielomianu[931].Y = 739.898

	punktyWykresuWielomianu[932].X = 9.32
	punktyWykresuWielomianu[932].Y = 742.335

	punktyWykresuWielomianu[933].X = 9.33
	punktyWykresuWielomianu[933].Y = 744.777

	punktyWykresuWielomianu[934].X = 9.34
	punktyWykresuWielomianu[934].Y = 747.224

	punktyWykresuWielomianu[935].X = 9.35
	punktyWykresuWielomianu[935].Y = 749.677

	punktyWykresuWielomianu[936].X = 9.36
	punktyWykresuWielomianu[936].Y = 752.136

	punktyWykresuWielomianu[937].X = 9.37
	punktyWykresuWielomianu[937].Y = 754.6

	punktyWykresuWielomianu[938].X = 9.38
	punktyWykresuWielomianu[938].Y = 757.069

	punktyWykresuWielomianu[939].X = 9.39
	punktyWykresuWielomianu[939].Y = 759.543

	punktyWykresuWielomianu[940].X = 9.4
	punktyWykresuWielomianu[940].Y = 762.024

	punktyWykresuWielomianu[941].X = 9.41
	punktyWykresuWielomianu[941].Y = 764.509

	punktyWykresuWielomianu[942].X = 9.42
	punktyWykresuWielomianu[942].Y = 767.0

	punktyWykresuWielomianu[943].X = 9.43
	punktyWykresuWielomianu[943].Y = 769.496

	punktyWykresuWielomianu[944].X = 9.44
	punktyWykresuWielomianu[944].Y = 771.998

	punktyWykresuWielomianu[945].X = 9.45
	punktyWykresuWielomianu[945].Y = 774.506

	punktyWykresuWielomianu[946].X = 9.46
	punktyWykresuWielomianu[946].Y = 777.018

	punktyWykresuWielomianu[947].X = 9.47
	punktyWykresuWielomianu[947].Y = 779.537

	punktyWykresuWielomianu[948].X = 9.48
	punktyWykresuWielomianu[948].Y = 782.061

	punktyWykresuWielomianu[949].X = 9.49
	punktyWykresuWielomianu[949].Y = 784.59

	punktyWykresuWielomianu[950].X = 9.5
	punktyWykresuWielomianu[950].Y = 787.125

	punktyWykresuWielomianu[951].X = 9.51
	punktyWykresuWielomianu[951].Y = 789.665

	punktyWykresuWielomianu[952].X = 9.52
	punktyWykresuWielomianu[952].Y = 792.211

	punktyWykresuWielomianu[953].X = 9.53
	punktyWykresuWielomianu[953].Y = 794.762

	punktyWykresuWielomianu[954].X = 9.54
	punktyWykresuWielomianu[954].Y = 797.319

	punktyWykresuWielomianu[955].X = 9.55
	punktyWykresuWielomianu[955].Y = 799.881

	punktyWykresuWielomianu[956].X = 9.56
	punktyWykresuWielomianu[956].Y = 802.449

	punktyWykresuWielomianu[957].X = 9.57
	punktyWykresuWielomianu[957].Y = 805.022

	punktyWykresuWielomianu[958].X = 9.58
	punktyWykresuWielomianu[958].Y = 807.601

	punktyWykresuWielomianu[959].X = 9.59
	punktyWykresuWielomianu[959].Y = 810.186

	punktyWykresuWielomianu[960].X = 9.6
	punktyWykresuWielomianu[960].Y = 812.776

	punktyWykresuWielomianu[961].X = 9.61
	punktyWykresuWielomianu[961].Y = 815.371

	punktyWykresuWielomianu[962].X = 9.62
	punktyWykresuWielomianu[962].Y = 817.972

	punktyWykresuWielomianu[963].X = 9.63
	punktyWykresuWielomianu[963].Y = 820.579

	punktyWykresuWielomianu[964].X = 9.64
	punktyWykresuWielomianu[964].Y = 823.191

	punktyWykresuWielomianu[965].X = 9.65
	punktyWykresuWielomianu[965].Y = 825.809

	punktyWykresuWielomianu[966].X = 9.66
	punktyWykresuWielomianu[966].Y = 828.433

	punktyWykresuWielomianu[967].X = 9.67
	punktyWykresuWielomianu[967].Y = 831.062

	punktyWykresuWielomianu[968].X = 9.68
	punktyWykresuWielomianu[968].Y = 833.696

	punktyWykresuWielomianu[969].X = 9.69
	punktyWykresuWielomianu[969].Y = 836.337

	punktyWykresuWielomianu[970].X = 9.7
	punktyWykresuWielomianu[970].Y = 838.983

	punktyWykresuWielomianu[971].X = 9.71
	punktyWykresuWielomianu[971].Y = 841.634

	punktyWykresuWielomianu[972].X = 9.72
	punktyWykresuWielomianu[972].Y = 844.291

	punktyWykresuWielomianu[973].X = 9.73
	punktyWykresuWielomianu[973].Y = 846.954

	punktyWykresuWielomianu[974].X = 9.74
	punktyWykresuWielomianu[974].Y = 849.622

	punktyWykresuWielomianu[975].X = 9.75
	punktyWykresuWielomianu[975].Y = 852.296

	punktyWykresuWielomianu[976].X = 9.76
	punktyWykresuWielomianu[976].Y = 854.976

	punktyWykresuWielomianu[977].X = 9.77
	punktyWykresuWielomianu[977].Y = 857.661

	punktyWykresuWielomianu[978].X = 9.78
	punktyWykresuWielomianu[978].Y = 860.353

	punktyWykresuWielomianu[979].X = 9.79
	punktyWykresuWielomianu[979].Y = 863.049

	punktyWykresuWielomianu[980].X = 9.8
	punktyWykresuWielomianu[980].Y = 865.752

	punktyWykresuWielomianu[981].X = 9.81
	punktyWykresuWielomianu[981].Y = 868.46

	punktyWykresuWielomianu[982].X = 9.82
	punktyWykresuWielomianu[982].Y = 871.173

	punktyWykresuWielomianu[983].X = 9.83
	punktyWykresuWielomianu[983].Y = 873.893

	punktyWykresuWielomianu[984].X = 9.84
	punktyWykresuWielomianu[984].Y = 876.618

	punktyWykresuWielomianu[985].X = 9.85
	punktyWykresuWielomianu[985].Y = 879.349

	punktyWykresuWielomianu[986].X = 9.86
	punktyWykresuWielomianu[986].Y = 882.085

	punktyWykresuWielomianu[987].X = 9.87
	punktyWykresuWielomianu[987].Y = 884.827

	punktyWykresuWielomianu[988].X = 9.88
	punktyWykresuWielomianu[988].Y = 887.575

	punktyWykresuWielomianu[989].X = 9.89
	punktyWykresuWielomianu[989].Y = 890.329

	punktyWykresuWielomianu[990].X = 9.9
	punktyWykresuWielomianu[990].Y = 893.089

	punktyWykresuWielomianu[991].X = 9.91
	punktyWykresuWielomianu[991].Y = 895.854

	punktyWykresuWielomianu[992].X = 9.92
	punktyWykresuWielomianu[992].Y = 898.625

	punktyWykresuWielomianu[993].X = 9.93
	punktyWykresuWielomianu[993].Y = 901.401

	punktyWykresuWielomianu[994].X = 9.94
	punktyWykresuWielomianu[994].Y = 904.183

	punktyWykresuWielomianu[995].X = 9.95
	punktyWykresuWielomianu[995].Y = 906.972

	punktyWykresuWielomianu[996].X = 9.96
	punktyWykresuWielomianu[996].Y = 909.766

	punktyWykresuWielomianu[997].X = 9.97
	punktyWykresuWielomianu[997].Y = 912.566

	punktyWykresuWielomianu[998].X = 9.98
	punktyWykresuWielomianu[998].Y = 915.371

	punktyWykresuWielomianu[999].X = 9.99
	punktyWykresuWielomianu[999].Y = 918.182

	punktyWykresuWielomianu[1000].X = 10.0
	punktyWykresuWielomianu[1000].Y = 921.0



	wykresWielomianu := plot.New()

	wykresWielomianu.Title.Text = "Wykres funkcji f(x) = x^3 - x^2 + 2x + 1"

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
		"Wykres-wielomianu-01.png"); err != nil {

		panic(err)
	}
}
