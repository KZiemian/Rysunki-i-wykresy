package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of polynomial f(x) = 0.000001 x^6 - + 0.001 x^5 -
	// - 0.01 x^4 - 0.1 x^3 + x^2 + x - 10.0.

	pointsOfPolynomialPlot := make(plotter.XYs, 2_001)

	pointsOfPolynomialPlot[0].X = 0.0
	pointsOfPolynomialPlot[0].Y = -10.0

	pointsOfPolynomialPlot[1].X = 0.01
	pointsOfPolynomialPlot[1].Y = -9.989

	pointsOfPolynomialPlot[2].X = 0.02
	pointsOfPolynomialPlot[2].Y = -9.979

	pointsOfPolynomialPlot[3].X = 0.03
	pointsOfPolynomialPlot[3].Y = -9.969

	pointsOfPolynomialPlot[4].X = 0.04
	pointsOfPolynomialPlot[4].Y = -9.958

	pointsOfPolynomialPlot[5].X = 0.05
	pointsOfPolynomialPlot[5].Y = -9.947

	pointsOfPolynomialPlot[6].X = 0.06
	pointsOfPolynomialPlot[6].Y = -9.936

	pointsOfPolynomialPlot[7].X = 0.07
	pointsOfPolynomialPlot[7].Y = -9.925

	pointsOfPolynomialPlot[8].X = 0.08
	pointsOfPolynomialPlot[8].Y = -9.913

	pointsOfPolynomialPlot[9].X = 0.09
	pointsOfPolynomialPlot[9].Y = -9.902

	pointsOfPolynomialPlot[10].X = 0.1
	pointsOfPolynomialPlot[10].Y = -9.89

	pointsOfPolynomialPlot[11].X = 0.11
	pointsOfPolynomialPlot[11].Y = -9.878

	pointsOfPolynomialPlot[12].X = 0.12
	pointsOfPolynomialPlot[12].Y = -9.865

	pointsOfPolynomialPlot[13].X = 0.13
	pointsOfPolynomialPlot[13].Y = -9.853

	pointsOfPolynomialPlot[14].X = 0.14
	pointsOfPolynomialPlot[14].Y = -9.84

	pointsOfPolynomialPlot[15].X = 0.15
	pointsOfPolynomialPlot[15].Y = -9.827

	pointsOfPolynomialPlot[16].X = 0.16
	pointsOfPolynomialPlot[16].Y = -9.814

	pointsOfPolynomialPlot[17].X = 0.17
	pointsOfPolynomialPlot[17].Y = -9.801

	pointsOfPolynomialPlot[18].X = 0.18
	pointsOfPolynomialPlot[18].Y = -9.788

	pointsOfPolynomialPlot[19].X = 0.19
	pointsOfPolynomialPlot[19].Y = -9.774

	pointsOfPolynomialPlot[20].X = 0.2
	pointsOfPolynomialPlot[20].Y = -9.76

	pointsOfPolynomialPlot[21].X = 0.21
	pointsOfPolynomialPlot[21].Y = -9.746

	pointsOfPolynomialPlot[22].X = 0.22
	pointsOfPolynomialPlot[22].Y = -9.732

	pointsOfPolynomialPlot[23].X = 0.23
	pointsOfPolynomialPlot[23].Y = -9.718

	pointsOfPolynomialPlot[24].X = 0.24
	pointsOfPolynomialPlot[24].Y = -9.703

	pointsOfPolynomialPlot[25].X = 0.25
	pointsOfPolynomialPlot[25].Y = -9.689

	pointsOfPolynomialPlot[26].X = 0.26
	pointsOfPolynomialPlot[26].Y = -9.674

	pointsOfPolynomialPlot[27].X = 0.27
	pointsOfPolynomialPlot[27].Y = -9.659

	pointsOfPolynomialPlot[28].X = 0.28
	pointsOfPolynomialPlot[28].Y = -9.643

	pointsOfPolynomialPlot[29].X = 0.29
	pointsOfPolynomialPlot[29].Y = -9.628

	pointsOfPolynomialPlot[30].X = 0.3
	pointsOfPolynomialPlot[30].Y = -9.612

	pointsOfPolynomialPlot[31].X = 0.31
	pointsOfPolynomialPlot[31].Y = -9.597

	pointsOfPolynomialPlot[32].X = 0.32
	pointsOfPolynomialPlot[32].Y = -9.581

	pointsOfPolynomialPlot[33].X = 0.33
	pointsOfPolynomialPlot[33].Y = -9.564

	pointsOfPolynomialPlot[34].X = 0.34
	pointsOfPolynomialPlot[34].Y = -9.548

	pointsOfPolynomialPlot[35].X = 0.35
	pointsOfPolynomialPlot[35].Y = -9.531

	pointsOfPolynomialPlot[36].X = 0.36
	pointsOfPolynomialPlot[36].Y = -9.515

	pointsOfPolynomialPlot[37].X = 0.37
	pointsOfPolynomialPlot[37].Y = -9.498

	pointsOfPolynomialPlot[38].X = 0.38
	pointsOfPolynomialPlot[38].Y = -9.481

	pointsOfPolynomialPlot[39].X = 0.39
	pointsOfPolynomialPlot[39].Y = -9.464

	pointsOfPolynomialPlot[40].X = 0.4
	pointsOfPolynomialPlot[40].Y = -9.446

	pointsOfPolynomialPlot[41].X = 0.41
	pointsOfPolynomialPlot[41].Y = -9.429

	pointsOfPolynomialPlot[42].X = 0.42
	pointsOfPolynomialPlot[42].Y = -9.411

	pointsOfPolynomialPlot[43].X = 0.43
	pointsOfPolynomialPlot[43].Y = -9.393

	pointsOfPolynomialPlot[44].X = 0.44
	pointsOfPolynomialPlot[44].Y = -9.375

	pointsOfPolynomialPlot[45].X = 0.45
	pointsOfPolynomialPlot[45].Y = -9.357

	pointsOfPolynomialPlot[46].X = 0.46
	pointsOfPolynomialPlot[46].Y = -9.338

	pointsOfPolynomialPlot[47].X = 0.47
	pointsOfPolynomialPlot[47].Y = -9.319

	pointsOfPolynomialPlot[48].X = 0.48
	pointsOfPolynomialPlot[48].Y = -9.301

	pointsOfPolynomialPlot[49].X = 0.49
	pointsOfPolynomialPlot[49].Y = -9.282

	pointsOfPolynomialPlot[50].X = 0.5
	pointsOfPolynomialPlot[50].Y = -9.263

	pointsOfPolynomialPlot[51].X = 0.51
	pointsOfPolynomialPlot[51].Y = -9.243

	pointsOfPolynomialPlot[52].X = 0.52
	pointsOfPolynomialPlot[52].Y = -9.224

	pointsOfPolynomialPlot[53].X = 0.53
	pointsOfPolynomialPlot[53].Y = -9.204

	pointsOfPolynomialPlot[54].X = 0.54
	pointsOfPolynomialPlot[54].Y = -9.185

	pointsOfPolynomialPlot[55].X = 0.55
	pointsOfPolynomialPlot[55].Y = -9.165

	pointsOfPolynomialPlot[56].X = 0.56
	pointsOfPolynomialPlot[56].Y = -9.144

	pointsOfPolynomialPlot[57].X = 0.57
	pointsOfPolynomialPlot[57].Y = -9.124

	pointsOfPolynomialPlot[58].X = 0.58
	pointsOfPolynomialPlot[58].Y = -9.104

	pointsOfPolynomialPlot[59].X = 0.59
	pointsOfPolynomialPlot[59].Y = -9.083

	pointsOfPolynomialPlot[60].X = 0.6
	pointsOfPolynomialPlot[60].Y = -9.062

	pointsOfPolynomialPlot[61].X = 0.61
	pointsOfPolynomialPlot[61].Y = -9.041

	pointsOfPolynomialPlot[62].X = 0.62
	pointsOfPolynomialPlot[62].Y = -9.02

	pointsOfPolynomialPlot[63].X = 0.63
	pointsOfPolynomialPlot[63].Y = -8.999

	pointsOfPolynomialPlot[64].X = 0.64
	pointsOfPolynomialPlot[64].Y = -8.978

	pointsOfPolynomialPlot[65].X = 0.65
	pointsOfPolynomialPlot[65].Y = -8.956

	pointsOfPolynomialPlot[66].X = 0.66
	pointsOfPolynomialPlot[66].Y = -8.934

	pointsOfPolynomialPlot[67].X = 0.67
	pointsOfPolynomialPlot[67].Y = -8.913

	pointsOfPolynomialPlot[68].X = 0.68
	pointsOfPolynomialPlot[68].Y = -8.891

	pointsOfPolynomialPlot[69].X = 0.69
	pointsOfPolynomialPlot[69].Y = -8.868

	pointsOfPolynomialPlot[70].X = 0.7
	pointsOfPolynomialPlot[70].Y = -8.846

	pointsOfPolynomialPlot[71].X = 0.71
	pointsOfPolynomialPlot[71].Y = -8.824

	pointsOfPolynomialPlot[72].X = 0.72
	pointsOfPolynomialPlot[72].Y = -8.801

	pointsOfPolynomialPlot[73].X = 0.73
	pointsOfPolynomialPlot[73].Y = -8.778

	pointsOfPolynomialPlot[74].X = 0.74
	pointsOfPolynomialPlot[74].Y = -8.755

	pointsOfPolynomialPlot[75].X = 0.75
	pointsOfPolynomialPlot[75].Y = -8.732

	pointsOfPolynomialPlot[76].X = 0.76
	pointsOfPolynomialPlot[76].Y = -8.709

	pointsOfPolynomialPlot[77].X = 0.77
	pointsOfPolynomialPlot[77].Y = -8.686

	pointsOfPolynomialPlot[78].X = 0.78
	pointsOfPolynomialPlot[78].Y = -8.662

	pointsOfPolynomialPlot[79].X = 0.79
	pointsOfPolynomialPlot[79].Y = -8.638

	pointsOfPolynomialPlot[80].X = 0.8
	pointsOfPolynomialPlot[80].Y = -8.615

	pointsOfPolynomialPlot[81].X = 0.81
	pointsOfPolynomialPlot[81].Y = -8.591

	pointsOfPolynomialPlot[82].X = 0.82
	pointsOfPolynomialPlot[82].Y = -8.566

	pointsOfPolynomialPlot[83].X = 0.83
	pointsOfPolynomialPlot[83].Y = -8.542

	pointsOfPolynomialPlot[84].X = 0.84
	pointsOfPolynomialPlot[84].Y = -8.518

	pointsOfPolynomialPlot[85].X = 0.85
	pointsOfPolynomialPlot[85].Y = -8.493

	pointsOfPolynomialPlot[86].X = 0.86
	pointsOfPolynomialPlot[86].Y = -8.469

	pointsOfPolynomialPlot[87].X = 0.87
	pointsOfPolynomialPlot[87].Y = -8.444

	pointsOfPolynomialPlot[88].X = 0.88
	pointsOfPolynomialPlot[88].Y = -8.419

	pointsOfPolynomialPlot[89].X = 0.89
	pointsOfPolynomialPlot[89].Y = -8.394

	pointsOfPolynomialPlot[90].X = 0.9
	pointsOfPolynomialPlot[90].Y = -8.368

	pointsOfPolynomialPlot[91].X = 0.91
	pointsOfPolynomialPlot[91].Y = -8.343

	pointsOfPolynomialPlot[92].X = 0.92
	pointsOfPolynomialPlot[92].Y = -8.318

	pointsOfPolynomialPlot[93].X = 0.93
	pointsOfPolynomialPlot[93].Y = -8.292

	pointsOfPolynomialPlot[94].X = 0.94
	pointsOfPolynomialPlot[94].Y = -8.266

	pointsOfPolynomialPlot[95].X = 0.95
	pointsOfPolynomialPlot[95].Y = -8.24

	pointsOfPolynomialPlot[96].X = 0.96
	pointsOfPolynomialPlot[96].Y = -8.214

	pointsOfPolynomialPlot[97].X = 0.97
	pointsOfPolynomialPlot[97].Y = -8.188

	pointsOfPolynomialPlot[98].X = 0.98
	pointsOfPolynomialPlot[98].Y = -8.162

	pointsOfPolynomialPlot[99].X = 0.99
	pointsOfPolynomialPlot[99].Y = -8.135

	pointsOfPolynomialPlot[100].X = 1.0
	pointsOfPolynomialPlot[100].Y = -8.109

	pointsOfPolynomialPlot[101].X = 1.01
	pointsOfPolynomialPlot[101].Y = -8.082

	pointsOfPolynomialPlot[102].X = 1.02
	pointsOfPolynomialPlot[102].Y = -8.055

	pointsOfPolynomialPlot[103].X = 1.03
	pointsOfPolynomialPlot[103].Y = -8.028

	pointsOfPolynomialPlot[104].X = 1.04
	pointsOfPolynomialPlot[104].Y = -8.001

	pointsOfPolynomialPlot[105].X = 1.05
	pointsOfPolynomialPlot[105].Y = -7.974

	pointsOfPolynomialPlot[106].X = 1.06
	pointsOfPolynomialPlot[106].Y = -7.946

	pointsOfPolynomialPlot[107].X = 1.07
	pointsOfPolynomialPlot[107].Y = -7.919

	pointsOfPolynomialPlot[108].X = 1.08
	pointsOfPolynomialPlot[108].Y = -7.891

	pointsOfPolynomialPlot[109].X = 1.09
	pointsOfPolynomialPlot[109].Y = -7.864

	pointsOfPolynomialPlot[110].X = 1.1
	pointsOfPolynomialPlot[110].Y = -7.836

	pointsOfPolynomialPlot[111].X = 1.11
	pointsOfPolynomialPlot[111].Y = -7.808

	pointsOfPolynomialPlot[112].X = 1.12
	pointsOfPolynomialPlot[112].Y = -7.78

	pointsOfPolynomialPlot[113].X = 1.13
	pointsOfPolynomialPlot[113].Y = -7.751

	pointsOfPolynomialPlot[114].X = 1.14
	pointsOfPolynomialPlot[114].Y = -7.723

	pointsOfPolynomialPlot[115].X = 1.15
	pointsOfPolynomialPlot[115].Y = -7.695

	pointsOfPolynomialPlot[116].X = 1.16
	pointsOfPolynomialPlot[116].Y = -7.666

	pointsOfPolynomialPlot[117].X = 1.17
	pointsOfPolynomialPlot[117].Y = -7.637

	pointsOfPolynomialPlot[118].X = 1.18
	pointsOfPolynomialPlot[118].Y = -7.609

	pointsOfPolynomialPlot[119].X = 1.19
	pointsOfPolynomialPlot[119].Y = -7.58

	pointsOfPolynomialPlot[120].X = 1.2
	pointsOfPolynomialPlot[120].Y = -7.551

	pointsOfPolynomialPlot[121].X = 1.21
	pointsOfPolynomialPlot[121].Y = -7.521

	pointsOfPolynomialPlot[122].X = 1.22
	pointsOfPolynomialPlot[122].Y = -7.492

	pointsOfPolynomialPlot[123].X = 1.23
	pointsOfPolynomialPlot[123].Y = -7.463

	pointsOfPolynomialPlot[124].X = 1.24
	pointsOfPolynomialPlot[124].Y = -7.433

	pointsOfPolynomialPlot[125].X = 1.25
	pointsOfPolynomialPlot[125].Y = -7.404

	pointsOfPolynomialPlot[126].X = 1.26
	pointsOfPolynomialPlot[126].Y = -7.374

	pointsOfPolynomialPlot[127].X = 1.27
	pointsOfPolynomialPlot[127].Y = -7.344

	pointsOfPolynomialPlot[128].X = 1.28
	pointsOfPolynomialPlot[128].Y = -7.314

	pointsOfPolynomialPlot[129].X = 1.29
	pointsOfPolynomialPlot[129].Y = -7.284

	pointsOfPolynomialPlot[130].X = 1.3
	pointsOfPolynomialPlot[130].Y = -7.254

	pointsOfPolynomialPlot[131].X = 1.31
	pointsOfPolynomialPlot[131].Y = -7.224

	pointsOfPolynomialPlot[132].X = 1.32
	pointsOfPolynomialPlot[132].Y = -7.193

	pointsOfPolynomialPlot[133].X = 1.33
	pointsOfPolynomialPlot[133].Y = -7.163

	pointsOfPolynomialPlot[134].X = 1.34
	pointsOfPolynomialPlot[134].Y = -7.132

	pointsOfPolynomialPlot[135].X = 1.35
	pointsOfPolynomialPlot[135].Y = -7.102

	pointsOfPolynomialPlot[136].X = 1.36
	pointsOfPolynomialPlot[136].Y = -7.071

	pointsOfPolynomialPlot[137].X = 1.37
	pointsOfPolynomialPlot[137].Y = -7.04

	pointsOfPolynomialPlot[138].X = 1.38
	pointsOfPolynomialPlot[138].Y = -7.009

	pointsOfPolynomialPlot[139].X = 1.39
	pointsOfPolynomialPlot[139].Y = -6.978

	pointsOfPolynomialPlot[140].X = 1.4
	pointsOfPolynomialPlot[140].Y = -6.947

	pointsOfPolynomialPlot[141].X = 1.41
	pointsOfPolynomialPlot[141].Y = -6.916

	pointsOfPolynomialPlot[142].X = 1.42
	pointsOfPolynomialPlot[142].Y = -6.884

	pointsOfPolynomialPlot[143].X = 1.43
	pointsOfPolynomialPlot[143].Y = -6.853

	pointsOfPolynomialPlot[144].X = 1.44
	pointsOfPolynomialPlot[144].Y = -6.821

	pointsOfPolynomialPlot[145].X = 1.45
	pointsOfPolynomialPlot[145].Y = -6.79

	pointsOfPolynomialPlot[146].X = 1.46
	pointsOfPolynomialPlot[146].Y = -6.758

	pointsOfPolynomialPlot[147].X = 1.47
	pointsOfPolynomialPlot[147].Y = -6.726

	pointsOfPolynomialPlot[148].X = 1.48
	pointsOfPolynomialPlot[148].Y = -6.694

	pointsOfPolynomialPlot[149].X = 1.49
	pointsOfPolynomialPlot[149].Y = -6.662

	pointsOfPolynomialPlot[150].X = 1.5
	pointsOfPolynomialPlot[150].Y = -6.63

	pointsOfPolynomialPlot[151].X = 1.51
	pointsOfPolynomialPlot[151].Y = -6.598

	pointsOfPolynomialPlot[152].X = 1.52
	pointsOfPolynomialPlot[152].Y = -6.566

	pointsOfPolynomialPlot[153].X = 1.53
	pointsOfPolynomialPlot[153].Y = -6.533

	pointsOfPolynomialPlot[154].X = 1.54
	pointsOfPolynomialPlot[154].Y = -6.501

	pointsOfPolynomialPlot[155].X = 1.55
	pointsOfPolynomialPlot[155].Y = -6.468

	pointsOfPolynomialPlot[156].X = 1.56
	pointsOfPolynomialPlot[156].Y = -6.436

	pointsOfPolynomialPlot[157].X = 1.57
	pointsOfPolynomialPlot[157].Y = -6.403

	pointsOfPolynomialPlot[158].X = 1.58
	pointsOfPolynomialPlot[158].Y = -6.37

	pointsOfPolynomialPlot[159].X = 1.59
	pointsOfPolynomialPlot[159].Y = -6.337

	pointsOfPolynomialPlot[160].X = 1.6
	pointsOfPolynomialPlot[160].Y = -6.304

	pointsOfPolynomialPlot[161].X = 1.61
	pointsOfPolynomialPlot[161].Y = -6.271

	pointsOfPolynomialPlot[162].X = 1.62
	pointsOfPolynomialPlot[162].Y = -6.238

	pointsOfPolynomialPlot[163].X = 1.63
	pointsOfPolynomialPlot[163].Y = -6.205

	pointsOfPolynomialPlot[164].X = 1.64
	pointsOfPolynomialPlot[164].Y = -6.172

	pointsOfPolynomialPlot[165].X = 1.65
	pointsOfPolynomialPlot[165].Y = -6.138

	pointsOfPolynomialPlot[166].X = 1.66
	pointsOfPolynomialPlot[166].Y = -6.105

	pointsOfPolynomialPlot[167].X = 1.67
	pointsOfPolynomialPlot[167].Y = -6.071

	pointsOfPolynomialPlot[168].X = 1.68
	pointsOfPolynomialPlot[168].Y = -6.038

	pointsOfPolynomialPlot[169].X = 1.69
	pointsOfPolynomialPlot[169].Y = -6.004

	pointsOfPolynomialPlot[170].X = 1.7
	pointsOfPolynomialPlot[170].Y = -5.97

	pointsOfPolynomialPlot[171].X = 1.71
	pointsOfPolynomialPlot[171].Y = -5.936

	pointsOfPolynomialPlot[172].X = 1.72
	pointsOfPolynomialPlot[172].Y = -5.909

	pointsOfPolynomialPlot[173].X = 1.73
	pointsOfPolynomialPlot[173].Y = -5.868

	pointsOfPolynomialPlot[174].X = 1.74
	pointsOfPolynomialPlot[174].Y = -5.834

	pointsOfPolynomialPlot[175].X = 1.75
	pointsOfPolynomialPlot[175].Y = -5.8

	pointsOfPolynomialPlot[176].X = 1.76
	pointsOfPolynomialPlot[176].Y = -5.766

	pointsOfPolynomialPlot[177].X = 1.77
	pointsOfPolynomialPlot[177].Y = -5.732

	pointsOfPolynomialPlot[178].X = 1.78
	pointsOfPolynomialPlot[178].Y = -5.698

	pointsOfPolynomialPlot[179].X = 1.79
	pointsOfPolynomialPlot[179].Y = -5.663

	pointsOfPolynomialPlot[180].X = 1.8
	pointsOfPolynomialPlot[180].Y = -5.629

	pointsOfPolynomialPlot[181].X = 1.81
	pointsOfPolynomialPlot[181].Y = -5.594

	pointsOfPolynomialPlot[182].X = 1.82
	pointsOfPolynomialPlot[182].Y = -5.56

	pointsOfPolynomialPlot[183].X = 1.83
	pointsOfPolynomialPlot[183].Y = -5.525

	pointsOfPolynomialPlot[184].X = 1.84
	pointsOfPolynomialPlot[184].Y = -5.49

	pointsOfPolynomialPlot[185].X = 1.85
	pointsOfPolynomialPlot[185].Y = -5.456

	pointsOfPolynomialPlot[186].X = 1.86
	pointsOfPolynomialPlot[186].Y = -5.421

	pointsOfPolynomialPlot[187].X = 1.87
	pointsOfPolynomialPlot[187].Y = -5.386

	pointsOfPolynomialPlot[188].X = 1.88
	pointsOfPolynomialPlot[188].Y = -5.351

	pointsOfPolynomialPlot[189].X = 1.89
	pointsOfPolynomialPlot[189].Y = -5.316

	pointsOfPolynomialPlot[190].X = 1.9
	pointsOfPolynomialPlot[190].Y = -5.281

	pointsOfPolynomialPlot[191].X = 1.91
	pointsOfPolynomialPlot[191].Y = -5.246

	pointsOfPolynomialPlot[192].X = 1.92
	pointsOfPolynomialPlot[192].Y = -5.211

	pointsOfPolynomialPlot[193].X = 1.93
	pointsOfPolynomialPlot[193].Y = -5.175

	pointsOfPolynomialPlot[194].X = 1.94
	pointsOfPolynomialPlot[194].Y = -5.14

	pointsOfPolynomialPlot[195].X = 1.95
	pointsOfPolynomialPlot[195].Y = -5.105

	pointsOfPolynomialPlot[196].X = 1.96
	pointsOfPolynomialPlot[196].Y = -5.07

	pointsOfPolynomialPlot[197].X = 1.97
	pointsOfPolynomialPlot[197].Y = -5.034

	pointsOfPolynomialPlot[198].X = 1.98
	pointsOfPolynomialPlot[198].Y = -4.999

	pointsOfPolynomialPlot[199].X = 1.99
	pointsOfPolynomialPlot[199].Y = -4.963

	pointsOfPolynomialPlot[200].X = 2.0
	pointsOfPolynomialPlot[200].Y = -4.927

	pointsOfPolynomialPlot[201].X = 2.01
	pointsOfPolynomialPlot[201].Y = -4.892

	pointsOfPolynomialPlot[202].X = 2.02
	pointsOfPolynomialPlot[202].Y = -4.856

	pointsOfPolynomialPlot[203].X = 2.03
	pointsOfPolynomialPlot[203].Y = -4.82

	pointsOfPolynomialPlot[204].X = 2.04
	pointsOfPolynomialPlot[204].Y = -4.785

	pointsOfPolynomialPlot[205].X = 2.05
	pointsOfPolynomialPlot[205].Y = -4.749

	pointsOfPolynomialPlot[206].X = 2.06
	pointsOfPolynomialPlot[206].Y = -4.713

	pointsOfPolynomialPlot[207].X = 2.07
	pointsOfPolynomialPlot[207].Y = -4.677

	pointsOfPolynomialPlot[208].X = 2.08
	pointsOfPolynomialPlot[208].Y = -4.641

	pointsOfPolynomialPlot[209].X = 2.09
	pointsOfPolynomialPlot[209].Y = -4.605

	pointsOfPolynomialPlot[210].X = 2.1
	pointsOfPolynomialPlot[210].Y = -4.569

	pointsOfPolynomialPlot[211].X = 2.11
	pointsOfPolynomialPlot[211].Y = -4.533

	pointsOfPolynomialPlot[212].X = 2.12
	pointsOfPolynomialPlot[212].Y = -4.497

	pointsOfPolynomialPlot[213].X = 2.13
	pointsOfPolynomialPlot[213].Y = -4.461

	pointsOfPolynomialPlot[214].X = 2.14
	pointsOfPolynomialPlot[214].Y = -4.425

	pointsOfPolynomialPlot[215].X = 2.15
	pointsOfPolynomialPlot[215].Y = -4.389

	pointsOfPolynomialPlot[216].X = 2.16
	pointsOfPolynomialPlot[216].Y = -4.352

	pointsOfPolynomialPlot[217].X = 2.17
	pointsOfPolynomialPlot[217].Y = -4.316

	pointsOfPolynomialPlot[218].X = 2.18
	pointsOfPolynomialPlot[218].Y = -4.28

	pointsOfPolynomialPlot[219].X = 2.19
	pointsOfPolynomialPlot[219].Y = -4.243

	pointsOfPolynomialPlot[220].X = 2.2
	pointsOfPolynomialPlot[220].Y = -4.207

	pointsOfPolynomialPlot[221].X = 2.21
	pointsOfPolynomialPlot[221].Y = -4.171

	pointsOfPolynomialPlot[222].X = 2.22
	pointsOfPolynomialPlot[222].Y = -4.134

	pointsOfPolynomialPlot[223].X = 2.23
	pointsOfPolynomialPlot[223].Y = -4.098

	pointsOfPolynomialPlot[224].X = 2.24
	pointsOfPolynomialPlot[224].Y = -4.061

	pointsOfPolynomialPlot[225].X = 2.25
	pointsOfPolynomialPlot[225].Y = -4.025

	pointsOfPolynomialPlot[226].X = 2.26
	pointsOfPolynomialPlot[226].Y = -3.988

	pointsOfPolynomialPlot[227].X = 2.27
	pointsOfPolynomialPlot[227].Y = -3.951

	pointsOfPolynomialPlot[228].X = 2.28
	pointsOfPolynomialPlot[228].Y = -3.915

	pointsOfPolynomialPlot[229].X = 2.29
	pointsOfPolynomialPlot[229].Y = -3.878

	pointsOfPolynomialPlot[230].X = 2.3
	pointsOfPolynomialPlot[230].Y = -3.842

	pointsOfPolynomialPlot[231].X = 2.31
	pointsOfPolynomialPlot[231].Y = -3.805

	pointsOfPolynomialPlot[232].X = 2.32
	pointsOfPolynomialPlot[232].Y = -3.768

	pointsOfPolynomialPlot[233].X = 2.33
	pointsOfPolynomialPlot[233].Y = -3.731

	pointsOfPolynomialPlot[234].X = 2.34
	pointsOfPolynomialPlot[234].Y = -3.695

	pointsOfPolynomialPlot[235].X = 2.35
	pointsOfPolynomialPlot[235].Y = -3.658

	pointsOfPolynomialPlot[236].X = 2.36
	pointsOfPolynomialPlot[236].Y = -3.621

	pointsOfPolynomialPlot[237].X = 2.37
	pointsOfPolynomialPlot[237].Y = -3.584

	pointsOfPolynomialPlot[238].X = 2.38
	pointsOfPolynomialPlot[238].Y = -3.548

	pointsOfPolynomialPlot[239].X = 2.39
	pointsOfPolynomialPlot[239].Y = -3.511

	pointsOfPolynomialPlot[240].X = 2.4
	pointsOfPolynomialPlot[240].Y = -3.474

	pointsOfPolynomialPlot[241].X = 2.41
	pointsOfPolynomialPlot[241].Y = -3.437

	pointsOfPolynomialPlot[242].X = 2.42
	pointsOfPolynomialPlot[242].Y = -3.4

	pointsOfPolynomialPlot[243].X = 2.43
	pointsOfPolynomialPlot[243].Y = -3.363

	pointsOfPolynomialPlot[244].X = 2.44
	pointsOfPolynomialPlot[244].Y = -3.326

	pointsOfPolynomialPlot[245].X = 2.45
	pointsOfPolynomialPlot[245].Y = -3.289

	pointsOfPolynomialPlot[246].X = 2.46
	pointsOfPolynomialPlot[246].Y = -3.253

	pointsOfPolynomialPlot[247].X = 2.47
	pointsOfPolynomialPlot[247].Y = -3.216

	pointsOfPolynomialPlot[248].X = 2.48
	pointsOfPolynomialPlot[248].Y = -3.179

	pointsOfPolynomialPlot[249].X = 2.49
	pointsOfPolynomialPlot[249].Y = -3.142

	pointsOfPolynomialPlot[250].X = 2.5
	pointsOfPolynomialPlot[250].Y = -3.105

	pointsOfPolynomialPlot[251].X = 2.51
	pointsOfPolynomialPlot[251].Y = -3.068

	pointsOfPolynomialPlot[252].X = 2.52
	pointsOfPolynomialPlot[252].Y = -3.031

	pointsOfPolynomialPlot[253].X = 2.53
	pointsOfPolynomialPlot[253].Y = -2.994

	pointsOfPolynomialPlot[254].X = 2.54
	pointsOfPolynomialPlot[254].Y = -2.957

	pointsOfPolynomialPlot[255].X = 2.55
	pointsOfPolynomialPlot[255].Y = -2.92

	pointsOfPolynomialPlot[256].X = 2.56
	pointsOfPolynomialPlot[256].Y = -2.883

	pointsOfPolynomialPlot[257].X = 2.57
	pointsOfPolynomialPlot[257].Y = -2.846

	pointsOfPolynomialPlot[258].X = 2.58
	pointsOfPolynomialPlot[258].Y = -2.809

	pointsOfPolynomialPlot[259].X = 2.59
	pointsOfPolynomialPlot[259].Y = -2.772

	pointsOfPolynomialPlot[260].X = 2.6
	pointsOfPolynomialPlot[260].Y = -2.735

	pointsOfPolynomialPlot[261].X = 2.61
	pointsOfPolynomialPlot[261].Y = -2.698

	pointsOfPolynomialPlot[262].X = 2.62
	pointsOfPolynomialPlot[262].Y = -2.661

	pointsOfPolynomialPlot[263].X = 2.63
	pointsOfPolynomialPlot[263].Y = -2.624

	pointsOfPolynomialPlot[264].X = 2.64
	pointsOfPolynomialPlot[264].Y = -2.587

	pointsOfPolynomialPlot[265].X = 2.65
	pointsOfPolynomialPlot[265].Y = -2.55

	pointsOfPolynomialPlot[266].X = 2.66
	pointsOfPolynomialPlot[266].Y = -2.513

	pointsOfPolynomialPlot[267].X = 2.67
	pointsOfPolynomialPlot[267].Y = -2.476

	pointsOfPolynomialPlot[268].X = 2.68
	pointsOfPolynomialPlot[268].Y = -2.439

	pointsOfPolynomialPlot[269].X = 2.69
	pointsOfPolynomialPlot[269].Y = -2.402

	pointsOfPolynomialPlot[270].X = 2.7
	pointsOfPolynomialPlot[270].Y = -2.365

	pointsOfPolynomialPlot[271].X = 2.71
	pointsOfPolynomialPlot[271].Y = -2.328

	pointsOfPolynomialPlot[272].X = 2.72
	pointsOfPolynomialPlot[272].Y = -2.292

	pointsOfPolynomialPlot[273].X = 2.73
	pointsOfPolynomialPlot[273].Y = -2.255

	pointsOfPolynomialPlot[274].X = 2.74
	pointsOfPolynomialPlot[274].Y = -2.218

	pointsOfPolynomialPlot[275].X = 2.75
	pointsOfPolynomialPlot[275].Y = -2.181

	pointsOfPolynomialPlot[276].X = 2.76
	pointsOfPolynomialPlot[276].Y = -2.144

	pointsOfPolynomialPlot[277].X = 2.77
	pointsOfPolynomialPlot[277].Y = -2.107

	pointsOfPolynomialPlot[278].X = 2.78
	pointsOfPolynomialPlot[278].Y = -2.07

	pointsOfPolynomialPlot[279].X = 2.79
	pointsOfPolynomialPlot[279].Y = -2.034

	pointsOfPolynomialPlot[280].X = 2.8
	pointsOfPolynomialPlot[280].Y = -1.997

	pointsOfPolynomialPlot[281].X = 2.81
	pointsOfPolynomialPlot[281].Y = -1.96

	pointsOfPolynomialPlot[282].X = 2.82
	pointsOfPolynomialPlot[282].Y = -1.923

	pointsOfPolynomialPlot[283].X = 2.83
	pointsOfPolynomialPlot[283].Y = -1.887

	pointsOfPolynomialPlot[284].X = 2.84
	pointsOfPolynomialPlot[284].Y = -1.85

	pointsOfPolynomialPlot[285].X = 2.85
	pointsOfPolynomialPlot[285].Y = -1.813

	pointsOfPolynomialPlot[286].X = 2.86
	pointsOfPolynomialPlot[286].Y = -1.776

	pointsOfPolynomialPlot[287].X = 2.87
	pointsOfPolynomialPlot[287].Y = -1.74

	pointsOfPolynomialPlot[288].X = 2.88
	pointsOfPolynomialPlot[288].Y = -1.703

	pointsOfPolynomialPlot[289].X = 2.89
	pointsOfPolynomialPlot[289].Y = -1.667

	pointsOfPolynomialPlot[290].X = 2.9
	pointsOfPolynomialPlot[290].Y = -1.63

	pointsOfPolynomialPlot[291].X = 2.91
	pointsOfPolynomialPlot[291].Y = -1.593

	pointsOfPolynomialPlot[292].X = 2.92
	pointsOfPolynomialPlot[292].Y = -1.557

	pointsOfPolynomialPlot[293].X = 2.93
	pointsOfPolynomialPlot[293].Y = -1.52

	pointsOfPolynomialPlot[294].X = 2.94
	pointsOfPolynomialPlot[294].Y = -1.484

	pointsOfPolynomialPlot[295].X = 2.95
	pointsOfPolynomialPlot[295].Y = -1.448

	pointsOfPolynomialPlot[296].X = 2.96
	pointsOfPolynomialPlot[296].Y = -1.411

	pointsOfPolynomialPlot[297].X = 2.97
	pointsOfPolynomialPlot[297].Y = -1.375

	pointsOfPolynomialPlot[298].X = 2.98
	pointsOfPolynomialPlot[298].Y = -1.338

	pointsOfPolynomialPlot[299].X = 2.99
	pointsOfPolynomialPlot[299].Y = -1.302

	pointsOfPolynomialPlot[300].X = 3.0
	pointsOfPolynomialPlot[300].Y = -1.266

	pointsOfPolynomialPlot[301].X = 3.01
	pointsOfPolynomialPlot[301].Y = -1.23

	pointsOfPolynomialPlot[302].X = 3.02
	pointsOfPolynomialPlot[302].Y = -1.193

	pointsOfPolynomialPlot[303].X = 3.03
	pointsOfPolynomialPlot[303].Y = -1.157

	pointsOfPolynomialPlot[304].X = 3.04
	pointsOfPolynomialPlot[304].Y = -1.121

	pointsOfPolynomialPlot[305].X = 3.05
	pointsOfPolynomialPlot[305].Y = -1.085

	pointsOfPolynomialPlot[306].X = 3.06
	pointsOfPolynomialPlot[306].Y = -1.049

	pointsOfPolynomialPlot[307].X = 3.07
	pointsOfPolynomialPlot[307].Y = -1.013

	pointsOfPolynomialPlot[308].X = 3.08
	pointsOfPolynomialPlot[308].Y = -0.977

	pointsOfPolynomialPlot[309].X = 3.09
	pointsOfPolynomialPlot[309].Y = -0.941

	pointsOfPolynomialPlot[310].X = 3.1
	pointsOfPolynomialPlot[310].Y = -0.905

	pointsOfPolynomialPlot[311].X = 3.11
	pointsOfPolynomialPlot[311].Y = -0.869

	pointsOfPolynomialPlot[312].X = 3.12
	pointsOfPolynomialPlot[312].Y = -0.833

	pointsOfPolynomialPlot[313].X = 3.13
	pointsOfPolynomialPlot[313].Y = -0.798

	pointsOfPolynomialPlot[314].X = 3.14
	pointsOfPolynomialPlot[314].Y = -0.762

	pointsOfPolynomialPlot[315].X = 3.15
	pointsOfPolynomialPlot[315].Y = -0.726

	pointsOfPolynomialPlot[316].X = 3.16
	pointsOfPolynomialPlot[316].Y = -0.69

	pointsOfPolynomialPlot[317].X = 3.17
	pointsOfPolynomialPlot[317].Y = -0.655

	pointsOfPolynomialPlot[318].X = 3.18
	pointsOfPolynomialPlot[318].Y = -0.619

	pointsOfPolynomialPlot[319].X = 3.19
	pointsOfPolynomialPlot[319].Y = -0.584

	pointsOfPolynomialPlot[320].X = 3.2
	pointsOfPolynomialPlot[320].Y = -0.548

	pointsOfPolynomialPlot[321].X = 3.21
	pointsOfPolynomialPlot[321].Y = -0.513

	pointsOfPolynomialPlot[322].X = 3.22
	pointsOfPolynomialPlot[322].Y = -0.478

	pointsOfPolynomialPlot[323].X = 3.23
	pointsOfPolynomialPlot[323].Y = -0.442

	pointsOfPolynomialPlot[324].X = 3.24
	pointsOfPolynomialPlot[324].Y = -0.407

	pointsOfPolynomialPlot[325].X = 3.25
	pointsOfPolynomialPlot[325].Y = -0.372

	pointsOfPolynomialPlot[326].X = 3.26
	pointsOfPolynomialPlot[326].Y = -0.337

	pointsOfPolynomialPlot[327].X = 3.27
	pointsOfPolynomialPlot[327].Y = -0.302

	pointsOfPolynomialPlot[328].X = 3.28
	pointsOfPolynomialPlot[328].Y = -0.266

	pointsOfPolynomialPlot[329].X = 3.29
	pointsOfPolynomialPlot[329].Y = -0.231

	pointsOfPolynomialPlot[330].X = 3.3
	pointsOfPolynomialPlot[330].Y = -0.197

	pointsOfPolynomialPlot[331].X = 3.31
	pointsOfPolynomialPlot[331].Y = -0.162

	pointsOfPolynomialPlot[332].X = 3.32
	pointsOfPolynomialPlot[332].Y = -0.127

	pointsOfPolynomialPlot[333].X = 3.33
	pointsOfPolynomialPlot[333].Y = -0.092

	pointsOfPolynomialPlot[334].X = 3.34
	pointsOfPolynomialPlot[334].Y = -0.057

	pointsOfPolynomialPlot[335].X = 3.35
	pointsOfPolynomialPlot[335].Y = -0.023

	pointsOfPolynomialPlot[336].X = 3.36
	pointsOfPolynomialPlot[336].Y = 0.011

	pointsOfPolynomialPlot[337].X = 3.37
	pointsOfPolynomialPlot[337].Y = 0.046

	pointsOfPolynomialPlot[338].X = 3.38
	pointsOfPolynomialPlot[338].Y = 0.08

	pointsOfPolynomialPlot[339].X = 3.39
	pointsOfPolynomialPlot[339].Y = 0.114

	pointsOfPolynomialPlot[340].X = 3.4
	pointsOfPolynomialPlot[340].Y = 0.149

	pointsOfPolynomialPlot[341].X = 3.41
	pointsOfPolynomialPlot[341].Y = 0.183

	pointsOfPolynomialPlot[342].X = 3.42
	pointsOfPolynomialPlot[342].Y = 0.217

	pointsOfPolynomialPlot[343].X = 3.43
	pointsOfPolynomialPlot[343].Y = 0.251

	pointsOfPolynomialPlot[344].X = 3.44
	pointsOfPolynomialPlot[344].Y = 0.285

	pointsOfPolynomialPlot[345].X = 3.45
	pointsOfPolynomialPlot[345].Y = 0.319

	pointsOfPolynomialPlot[346].X = 3.46
	pointsOfPolynomialPlot[346].Y = 0.353

	pointsOfPolynomialPlot[347].X = 3.47
	pointsOfPolynomialPlot[347].Y = 0.387

	pointsOfPolynomialPlot[348].X = 3.48
	pointsOfPolynomialPlot[348].Y = 0.421

	pointsOfPolynomialPlot[349].X = 3.49
	pointsOfPolynomialPlot[349].Y = 0.455

	pointsOfPolynomialPlot[350].X = 3.5
	pointsOfPolynomialPlot[350].Y = 0.488

	pointsOfPolynomialPlot[351].X = 3.51
	pointsOfPolynomialPlot[351].Y = 0.522

	pointsOfPolynomialPlot[352].X = 3.52
	pointsOfPolynomialPlot[352].Y = 0.556

	pointsOfPolynomialPlot[353].X = 3.53
	pointsOfPolynomialPlot[353].Y = 0.589

	pointsOfPolynomialPlot[354].X = 3.54
	pointsOfPolynomialPlot[354].Y = 0.622

	pointsOfPolynomialPlot[355].X = 3.55
	pointsOfPolynomialPlot[355].Y = 0.656

	pointsOfPolynomialPlot[356].X = 3.56
	pointsOfPolynomialPlot[356].Y = 0.689

	pointsOfPolynomialPlot[357].X = 3.57
	pointsOfPolynomialPlot[357].Y = 0.722

	pointsOfPolynomialPlot[358].X = 3.58
	pointsOfPolynomialPlot[358].Y = 0.755

	pointsOfPolynomialPlot[359].X = 3.59
	pointsOfPolynomialPlot[359].Y = 0.788

	pointsOfPolynomialPlot[360].X = 3.6
	pointsOfPolynomialPlot[360].Y = 0.821

	pointsOfPolynomialPlot[361].X = 3.61
	pointsOfPolynomialPlot[361].Y = 0.854

	pointsOfPolynomialPlot[362].X = 3.62
	pointsOfPolynomialPlot[362].Y = 0.887

	pointsOfPolynomialPlot[363].X = 3.63
	pointsOfPolynomialPlot[363].Y = 0.919

	pointsOfPolynomialPlot[364].X = 3.64
	pointsOfPolynomialPlot[364].Y = 0.952

	pointsOfPolynomialPlot[365].X = 3.65
	pointsOfPolynomialPlot[365].Y = 0.985

	pointsOfPolynomialPlot[366].X = 3.66
	pointsOfPolynomialPlot[366].Y = 1.017

	pointsOfPolynomialPlot[367].X = 3.67
	pointsOfPolynomialPlot[367].Y = 1.049

	pointsOfPolynomialPlot[368].X = 3.68
	pointsOfPolynomialPlot[368].Y = 1.082

	pointsOfPolynomialPlot[369].X = 3.69
	pointsOfPolynomialPlot[369].Y = 1.114

	pointsOfPolynomialPlot[370].X = 3.7
	pointsOfPolynomialPlot[370].Y = 1.146

	pointsOfPolynomialPlot[371].X = 3.71
	pointsOfPolynomialPlot[371].Y = 1.178

	pointsOfPolynomialPlot[372].X = 3.72
	pointsOfPolynomialPlot[372].Y = 1.21

	pointsOfPolynomialPlot[373].X = 3.73
	pointsOfPolynomialPlot[373].Y = 1.242

	pointsOfPolynomialPlot[374].X = 3.74
	pointsOfPolynomialPlot[374].Y = 1.274

	pointsOfPolynomialPlot[375].X = 3.75
	pointsOfPolynomialPlot[375].Y = 1.305

	pointsOfPolynomialPlot[376].X = 3.76
	pointsOfPolynomialPlot[376].Y = 1.337

	pointsOfPolynomialPlot[377].X = 3.77
	pointsOfPolynomialPlot[377].Y = 1.369

	pointsOfPolynomialPlot[378].X = 3.78
	pointsOfPolynomialPlot[378].Y = 1.4

	pointsOfPolynomialPlot[379].X = 3.79
	pointsOfPolynomialPlot[379].Y = 1.431

	pointsOfPolynomialPlot[380].X = 3.8
	pointsOfPolynomialPlot[380].Y = 1.463

	pointsOfPolynomialPlot[381].X = 3.81
	pointsOfPolynomialPlot[381].Y = 1.494

	pointsOfPolynomialPlot[382].X = 3.82
	pointsOfPolynomialPlot[382].Y = 1.525

	pointsOfPolynomialPlot[383].X = 3.83
	pointsOfPolynomialPlot[383].Y = 1.556

	pointsOfPolynomialPlot[384].X = 3.84
	pointsOfPolynomialPlot[384].Y = 1.587

	pointsOfPolynomialPlot[385].X = 3.85
	pointsOfPolynomialPlot[385].Y = 1.617

	pointsOfPolynomialPlot[386].X = 3.86
	pointsOfPolynomialPlot[386].Y = 1.648

	pointsOfPolynomialPlot[387].X = 3.87
	pointsOfPolynomialPlot[387].Y = 1.679

	pointsOfPolynomialPlot[388].X = 3.88
	pointsOfPolynomialPlot[388].Y = 1.709

	pointsOfPolynomialPlot[389].X = 3.89
	pointsOfPolynomialPlot[389].Y = 1.74

	pointsOfPolynomialPlot[390].X = 3.9
	pointsOfPolynomialPlot[390].Y = 1.77

	pointsOfPolynomialPlot[391].X = 3.91
	pointsOfPolynomialPlot[391].Y = 1.8

	pointsOfPolynomialPlot[392].X = 3.92
	pointsOfPolynomialPlot[392].Y = 1.83

	pointsOfPolynomialPlot[393].X = 3.93
	pointsOfPolynomialPlot[393].Y = 1.86

	pointsOfPolynomialPlot[394].X = 3.94
	pointsOfPolynomialPlot[394].Y = 1.89

	pointsOfPolynomialPlot[395].X = 3.95
	pointsOfPolynomialPlot[395].Y = 1.92

	pointsOfPolynomialPlot[396].X = 3.96
	pointsOfPolynomialPlot[396].Y = 1.95

	pointsOfPolynomialPlot[397].X = 3.97
	pointsOfPolynomialPlot[397].Y = 1.979

	pointsOfPolynomialPlot[398].X = 3.98
	pointsOfPolynomialPlot[398].Y = 2.009

	pointsOfPolynomialPlot[399].X = 3.99
	pointsOfPolynomialPlot[399].Y = 2.038

	pointsOfPolynomialPlot[400].X = 4.0
	pointsOfPolynomialPlot[400].Y = 2.068

	pointsOfPolynomialPlot[401].X = 4.01
	pointsOfPolynomialPlot[401].Y = 2.097

	pointsOfPolynomialPlot[402].X = 4.02
	pointsOfPolynomialPlot[402].Y = 2.126

	pointsOfPolynomialPlot[403].X = 4.03
	pointsOfPolynomialPlot[403].Y = 2.155

	pointsOfPolynomialPlot[404].X = 4.04
	pointsOfPolynomialPlot[404].Y = 2.184

	pointsOfPolynomialPlot[405].X = 4.05
	pointsOfPolynomialPlot[405].Y = 2.213

	pointsOfPolynomialPlot[406].X = 4.06
	pointsOfPolynomialPlot[406].Y = 2.241

	pointsOfPolynomialPlot[407].X = 4.07
	pointsOfPolynomialPlot[407].Y = 2.27

	pointsOfPolynomialPlot[408].X = 4.08
	pointsOfPolynomialPlot[408].Y = 2.298

	pointsOfPolynomialPlot[409].X = 4.09
	pointsOfPolynomialPlot[409].Y = 2.327

	pointsOfPolynomialPlot[410].X = 4.1
	pointsOfPolynomialPlot[410].Y = 2.355

	pointsOfPolynomialPlot[411].X = 4.11
	pointsOfPolynomialPlot[411].Y = 2.383

	pointsOfPolynomialPlot[412].X = 4.12
	pointsOfPolynomialPlot[412].Y = 2.411

	pointsOfPolynomialPlot[413].X = 4.13
	pointsOfPolynomialPlot[413].Y = 2.439

	pointsOfPolynomialPlot[414].X = 4.14
	pointsOfPolynomialPlot[414].Y = 2.467

	pointsOfPolynomialPlot[415].X = 4.15
	pointsOfPolynomialPlot[415].Y = 2.495

	pointsOfPolynomialPlot[416].X = 4.16
	pointsOfPolynomialPlot[416].Y = 2.522

	pointsOfPolynomialPlot[417].X = 4.17
	pointsOfPolynomialPlot[417].Y = 2.55

	pointsOfPolynomialPlot[418].X = 4.18
	pointsOfPolynomialPlot[418].Y = 2.577

	pointsOfPolynomialPlot[419].X = 4.19
	pointsOfPolynomialPlot[419].Y = 2.604

	pointsOfPolynomialPlot[420].X = 4.2
	pointsOfPolynomialPlot[420].Y = 2.631

	pointsOfPolynomialPlot[421].X = 4.21
	pointsOfPolynomialPlot[421].Y = 2.658

	pointsOfPolynomialPlot[422].X = 4.22
	pointsOfPolynomialPlot[422].Y = 2.685

	pointsOfPolynomialPlot[423].X = 4.23
	pointsOfPolynomialPlot[423].Y = 2.712

	pointsOfPolynomialPlot[424].X = 4.24
	pointsOfPolynomialPlot[424].Y = 2.739

	pointsOfPolynomialPlot[425].X = 4.25
	pointsOfPolynomialPlot[425].Y = 2.765

	pointsOfPolynomialPlot[426].X = 4.26
	pointsOfPolynomialPlot[426].Y = 2.792

	pointsOfPolynomialPlot[427].X = 4.27
	pointsOfPolynomialPlot[427].Y = 2.818

	pointsOfPolynomialPlot[428].X = 4.28
	pointsOfPolynomialPlot[428].Y = 2.844

	pointsOfPolynomialPlot[429].X = 4.29
	pointsOfPolynomialPlot[429].Y = 2.87

	pointsOfPolynomialPlot[430].X = 4.30
	pointsOfPolynomialPlot[430].Y = 2.896

	pointsOfPolynomialPlot[431].X = 4.31
	pointsOfPolynomialPlot[431].Y = 2.922

	pointsOfPolynomialPlot[432].X = 4.32
	pointsOfPolynomialPlot[432].Y = 2.948

	pointsOfPolynomialPlot[433].X = 4.33
	pointsOfPolynomialPlot[433].Y = 2.974

	pointsOfPolynomialPlot[434].X = 4.34
	pointsOfPolynomialPlot[434].Y = 2.999

	pointsOfPolynomialPlot[435].X = 4.35
	pointsOfPolynomialPlot[435].Y = 3.024

	pointsOfPolynomialPlot[436].X = 4.36
	pointsOfPolynomialPlot[436].Y = 3.05

	pointsOfPolynomialPlot[437].X = 4.37
	pointsOfPolynomialPlot[437].Y = 3.075

	pointsOfPolynomialPlot[438].X = 4.38
	pointsOfPolynomialPlot[438].Y = 3.1

	pointsOfPolynomialPlot[439].X = 4.39
	pointsOfPolynomialPlot[439].Y = 3.125

	pointsOfPolynomialPlot[440].X = 4.4
	pointsOfPolynomialPlot[440].Y = 3.149

	pointsOfPolynomialPlot[441].X = 4.41
	pointsOfPolynomialPlot[441].Y = 3.174

	pointsOfPolynomialPlot[442].X = 4.42
	pointsOfPolynomialPlot[442].Y = 3.199

	pointsOfPolynomialPlot[443].X = 4.43
	pointsOfPolynomialPlot[443].Y = 3.223

	pointsOfPolynomialPlot[444].X = 4.44
	pointsOfPolynomialPlot[444].Y = 3.247

	pointsOfPolynomialPlot[445].X = 4.45
	pointsOfPolynomialPlot[445].Y = 3.271

	pointsOfPolynomialPlot[446].X = 4.46
	pointsOfPolynomialPlot[446].Y = 3.295

	pointsOfPolynomialPlot[447].X = 4.47
	pointsOfPolynomialPlot[447].Y = 3.319

	pointsOfPolynomialPlot[448].X = 4.48
	pointsOfPolynomialPlot[448].Y = 3.343

	pointsOfPolynomialPlot[449].X = 4.49
	pointsOfPolynomialPlot[449].Y = 3.367

	pointsOfPolynomialPlot[450].X = 4.5
	pointsOfPolynomialPlot[450].Y = 3.39

	pointsOfPolynomialPlot[451].X = 4.51
	pointsOfPolynomialPlot[451].Y = 3.413

	pointsOfPolynomialPlot[452].X = 4.52
	pointsOfPolynomialPlot[452].Y = 3.437

	pointsOfPolynomialPlot[453].X = 4.53
	pointsOfPolynomialPlot[453].Y = 3.46

	pointsOfPolynomialPlot[454].X = 4.54
	pointsOfPolynomialPlot[454].Y = 3.483

	pointsOfPolynomialPlot[455].X = 4.55
	pointsOfPolynomialPlot[455].Y = 3.505

	pointsOfPolynomialPlot[456].X = 4.56
	pointsOfPolynomialPlot[456].Y = 3.528

	pointsOfPolynomialPlot[457].X = 4.57
	pointsOfPolynomialPlot[457].Y = 3.551

	pointsOfPolynomialPlot[458].X = 4.58
	pointsOfPolynomialPlot[458].Y = 3.573

	pointsOfPolynomialPlot[459].X = 4.59
	pointsOfPolynomialPlot[459].Y = 3.595

	pointsOfPolynomialPlot[460].X = 4.6
	pointsOfPolynomialPlot[460].Y = 3.618

	pointsOfPolynomialPlot[461].X = 4.61
	pointsOfPolynomialPlot[461].Y = 3.64

	pointsOfPolynomialPlot[462].X = 4.62
	pointsOfPolynomialPlot[462].Y = 3.662

	pointsOfPolynomialPlot[463].X = 4.63
	pointsOfPolynomialPlot[463].Y = 3.683

	pointsOfPolynomialPlot[464].X = 4.64
	pointsOfPolynomialPlot[464].Y = 3.705

	pointsOfPolynomialPlot[465].X = 4.65
	pointsOfPolynomialPlot[465].Y = 3.726

	pointsOfPolynomialPlot[466].X = 4.66
	pointsOfPolynomialPlot[466].Y = 3.748

	pointsOfPolynomialPlot[467].X = 4.67
	pointsOfPolynomialPlot[467].Y = 3.769

	pointsOfPolynomialPlot[468].X = 4.68
	pointsOfPolynomialPlot[468].Y = 3.79

	pointsOfPolynomialPlot[469].X = 4.69
	pointsOfPolynomialPlot[469].Y = 3.811

	pointsOfPolynomialPlot[470].X = 4.7
	pointsOfPolynomialPlot[470].Y = 3.832

	pointsOfPolynomialPlot[471].X = 4.71
	pointsOfPolynomialPlot[471].Y = 3.852

	pointsOfPolynomialPlot[472].X = 4.72
	pointsOfPolynomialPlot[472].Y = 3.873

	pointsOfPolynomialPlot[473].X = 4.73
	pointsOfPolynomialPlot[473].Y = 3.893

	pointsOfPolynomialPlot[474].X = 4.74
	pointsOfPolynomialPlot[474].Y = 3.914

	pointsOfPolynomialPlot[475].X = 4.75
	pointsOfPolynomialPlot[475].Y = 3.934

	pointsOfPolynomialPlot[476].X = 4.76
	pointsOfPolynomialPlot[476].Y = 3.954

	pointsOfPolynomialPlot[477].X = 4.77
	pointsOfPolynomialPlot[477].Y = 3.974

	pointsOfPolynomialPlot[478].X = 4.78
	pointsOfPolynomialPlot[478].Y = 3.993

	pointsOfPolynomialPlot[479].X = 4.79
	pointsOfPolynomialPlot[479].Y = 4.013

	pointsOfPolynomialPlot[480].X = 4.8
	pointsOfPolynomialPlot[480].Y = 4.032

	pointsOfPolynomialPlot[481].X = 4.81
	pointsOfPolynomialPlot[481].Y = 4.051

	pointsOfPolynomialPlot[482].X = 4.82
	pointsOfPolynomialPlot[482].Y = 4.071

	pointsOfPolynomialPlot[483].X = 4.83
	pointsOfPolynomialPlot[483].Y = 4.09

	pointsOfPolynomialPlot[484].X = 4.84
	pointsOfPolynomialPlot[484].Y = 4.108

	pointsOfPolynomialPlot[485].X = 4.85
	pointsOfPolynomialPlot[485].Y = 4.127

	pointsOfPolynomialPlot[486].X = 4.86
	pointsOfPolynomialPlot[486].Y = 4.146

	pointsOfPolynomialPlot[487].X = 4.87
	pointsOfPolynomialPlot[487].Y = 4.164

	pointsOfPolynomialPlot[488].X = 4.88
	pointsOfPolynomialPlot[488].Y = 4.182

	pointsOfPolynomialPlot[489].X = 4.89
	pointsOfPolynomialPlot[489].Y = 4.2

	pointsOfPolynomialPlot[490].X = 4.9
	pointsOfPolynomialPlot[490].Y = 4.218

	pointsOfPolynomialPlot[491].X = 4.91
	pointsOfPolynomialPlot[491].Y = 4.236

	pointsOfPolynomialPlot[492].X = 4.92
	pointsOfPolynomialPlot[492].Y = 4.254

	pointsOfPolynomialPlot[493].X = 4.93
	pointsOfPolynomialPlot[493].Y = 4.272

	pointsOfPolynomialPlot[494].X = 4.94
	pointsOfPolynomialPlot[494].Y = 4.289

	pointsOfPolynomialPlot[495].X = 4.95
	pointsOfPolynomialPlot[495].Y = 4.306

	pointsOfPolynomialPlot[496].X = 4.96
	pointsOfPolynomialPlot[496].Y = 4.323

	pointsOfPolynomialPlot[497].X = 4.97
	pointsOfPolynomialPlot[497].Y = 4.34

	pointsOfPolynomialPlot[498].X = 4.98
	pointsOfPolynomialPlot[498].Y = 4.357

	pointsOfPolynomialPlot[499].X = 4.99
	pointsOfPolynomialPlot[499].Y = 4.374

	pointsOfPolynomialPlot[500].X = 5.0
	pointsOfPolynomialPlot[500].Y = 4.39

	pointsOfPolynomialPlot[501].X = 5.01
	pointsOfPolynomialPlot[501].Y = 4.407

	pointsOfPolynomialPlot[502].X = 5.02
	pointsOfPolynomialPlot[502].Y = 4.423

	pointsOfPolynomialPlot[503].X = 5.03
	pointsOfPolynomialPlot[503].Y = 4.439

	pointsOfPolynomialPlot[504].X = 5.04
	pointsOfPolynomialPlot[504].Y = 4.455

	pointsOfPolynomialPlot[505].X = 5.05
	pointsOfPolynomialPlot[505].Y = 4.471

	pointsOfPolynomialPlot[506].X = 5.06
	pointsOfPolynomialPlot[506].Y = 4.486

	pointsOfPolynomialPlot[507].X = 5.07
	pointsOfPolynomialPlot[507].Y = 4.502

	pointsOfPolynomialPlot[508].X = 5.08
	pointsOfPolynomialPlot[508].Y = 4.517

	pointsOfPolynomialPlot[509].X = 5.09
	pointsOfPolynomialPlot[509].Y = 4.532

	pointsOfPolynomialPlot[510].X = 5.10
	pointsOfPolynomialPlot[510].Y = 4.547

	pointsOfPolynomialPlot[511].X = 5.11
	pointsOfPolynomialPlot[511].Y = 4.562

	pointsOfPolynomialPlot[512].X = 5.12
	pointsOfPolynomialPlot[512].Y = 4.577

	pointsOfPolynomialPlot[513].X = 5.13
	pointsOfPolynomialPlot[513].Y = 4.591

	pointsOfPolynomialPlot[514].X = 5.14
	pointsOfPolynomialPlot[514].Y = 4.606

	pointsOfPolynomialPlot[515].X = 5.15
	pointsOfPolynomialPlot[515].Y = 4.62

	pointsOfPolynomialPlot[516].X = 5.16
	pointsOfPolynomialPlot[516].Y = 4.634

	pointsOfPolynomialPlot[517].X = 5.17
	pointsOfPolynomialPlot[517].Y = 4.648

	pointsOfPolynomialPlot[518].X = 5.18
	pointsOfPolynomialPlot[518].Y = 4.662

	pointsOfPolynomialPlot[519].X = 5.19
	pointsOfPolynomialPlot[519].Y = 4.675

	pointsOfPolynomialPlot[520].X = 5.2
	pointsOfPolynomialPlot[520].Y = 4.689

	pointsOfPolynomialPlot[521].X = 5.21
	pointsOfPolynomialPlot[521].Y = 4.702

	pointsOfPolynomialPlot[522].X = 5.22
	pointsOfPolynomialPlot[522].Y = 4.715

	pointsOfPolynomialPlot[523].X = 5.23
	pointsOfPolynomialPlot[523].Y = 4.729

	pointsOfPolynomialPlot[524].X = 5.24
	pointsOfPolynomialPlot[524].Y = 4.741

	pointsOfPolynomialPlot[525].X = 5.25
	pointsOfPolynomialPlot[525].Y = 4.754

	pointsOfPolynomialPlot[526].X = 5.26
	pointsOfPolynomialPlot[526].Y = 4.767

	pointsOfPolynomialPlot[527].X = 5.27
	pointsOfPolynomialPlot[527].Y = 4.779

	pointsOfPolynomialPlot[528].X = 5.28
	pointsOfPolynomialPlot[528].Y = 4.791

	pointsOfPolynomialPlot[529].X = 5.29
	pointsOfPolynomialPlot[529].Y = 4.804

	pointsOfPolynomialPlot[530].X = 5.3
	pointsOfPolynomialPlot[530].Y = 4.815

	pointsOfPolynomialPlot[531].X = 5.31
	pointsOfPolynomialPlot[531].Y = 4.827

	pointsOfPolynomialPlot[532].X = 5.32
	pointsOfPolynomialPlot[532].Y = 4.839

	pointsOfPolynomialPlot[533].X = 5.33
	pointsOfPolynomialPlot[533].Y = 4.85

	pointsOfPolynomialPlot[534].X = 5.34
	pointsOfPolynomialPlot[534].Y = 4.862

	pointsOfPolynomialPlot[535].X = 5.35
	pointsOfPolynomialPlot[535].Y = 4.873

	pointsOfPolynomialPlot[536].X = 5.36
	pointsOfPolynomialPlot[536].Y = 4.884

	pointsOfPolynomialPlot[537].X = 5.37
	pointsOfPolynomialPlot[537].Y = 4.895

	pointsOfPolynomialPlot[538].X = 5.38
	pointsOfPolynomialPlot[538].Y = 4.906

	pointsOfPolynomialPlot[539].X = 5.39
	pointsOfPolynomialPlot[539].Y = 4.916

	pointsOfPolynomialPlot[540].X = 5.4
	pointsOfPolynomialPlot[540].Y = 4.927

	pointsOfPolynomialPlot[541].X = 5.41
	pointsOfPolynomialPlot[541].Y = 4.937

	pointsOfPolynomialPlot[542].X = 5.42
	pointsOfPolynomialPlot[542].Y = 4.947

	pointsOfPolynomialPlot[543].X = 5.43
	pointsOfPolynomialPlot[543].Y = 4.957

	pointsOfPolynomialPlot[544].X = 5.44
	pointsOfPolynomialPlot[544].Y = 4.967

	pointsOfPolynomialPlot[545].X = 5.45
	pointsOfPolynomialPlot[545].Y = 4.976

	pointsOfPolynomialPlot[546].X = 5.46
	pointsOfPolynomialPlot[546].Y = 4.986

	pointsOfPolynomialPlot[547].X = 5.47
	pointsOfPolynomialPlot[547].Y = 4.995

	pointsOfPolynomialPlot[548].X = 5.48
	pointsOfPolynomialPlot[548].Y = 5.004

	pointsOfPolynomialPlot[549].X = 5.49
	pointsOfPolynomialPlot[549].Y = 5.013

	pointsOfPolynomialPlot[550].X = 5.5
	pointsOfPolynomialPlot[550].Y = 5.022

	pointsOfPolynomialPlot[551].X = 5.51
	pointsOfPolynomialPlot[551].Y = 5.031

	pointsOfPolynomialPlot[552].X = 5.52
	pointsOfPolynomialPlot[552].Y = 5.039

	pointsOfPolynomialPlot[553].X = 5.53
	pointsOfPolynomialPlot[553].Y = 5.048

	pointsOfPolynomialPlot[554].X = 5.54
	pointsOfPolynomialPlot[554].Y = 5.056

	pointsOfPolynomialPlot[555].X = 5.55
	pointsOfPolynomialPlot[555].Y = 5.064

	pointsOfPolynomialPlot[556].X = 5.56
	pointsOfPolynomialPlot[556].Y = 5.072

	pointsOfPolynomialPlot[557].X = 5.57
	pointsOfPolynomialPlot[557].Y = 5.079

	pointsOfPolynomialPlot[558].X = 5.58
	pointsOfPolynomialPlot[558].Y = 5.087

	pointsOfPolynomialPlot[559].X = 5.59
	pointsOfPolynomialPlot[559].Y = 5.094

	pointsOfPolynomialPlot[560].X = 5.6
	pointsOfPolynomialPlot[560].Y = 5.102

	pointsOfPolynomialPlot[561].X = 5.61
	pointsOfPolynomialPlot[561].Y = 5.109

	pointsOfPolynomialPlot[562].X = 5.62
	pointsOfPolynomialPlot[562].Y = 5.116

	pointsOfPolynomialPlot[563].X = 5.63
	pointsOfPolynomialPlot[563].Y = 5.122

	pointsOfPolynomialPlot[564].X = 5.64
	pointsOfPolynomialPlot[564].Y = 5.129

	pointsOfPolynomialPlot[565].X = 5.65
	pointsOfPolynomialPlot[565].Y = 5.136

	pointsOfPolynomialPlot[566].X = 5.66
	pointsOfPolynomialPlot[566].Y = 5.142

	pointsOfPolynomialPlot[567].X = 5.67
	pointsOfPolynomialPlot[567].Y = 5.148

	pointsOfPolynomialPlot[568].X = 5.68
	pointsOfPolynomialPlot[568].Y = 5.154

	pointsOfPolynomialPlot[569].X = 5.69
	pointsOfPolynomialPlot[569].Y = 5.16

	pointsOfPolynomialPlot[570].X = 5.7
	pointsOfPolynomialPlot[570].Y = 5.165

	pointsOfPolynomialPlot[571].X = 5.71
	pointsOfPolynomialPlot[571].Y = 5.171

	pointsOfPolynomialPlot[572].X = 5.72
	pointsOfPolynomialPlot[572].Y = 5.176

	pointsOfPolynomialPlot[573].X = 5.73
	pointsOfPolynomialPlot[573].Y = 5.182

	pointsOfPolynomialPlot[574].X = 5.74
	pointsOfPolynomialPlot[574].Y = 5.187

	pointsOfPolynomialPlot[575].X = 5.75
	pointsOfPolynomialPlot[575].Y = 5.191

	pointsOfPolynomialPlot[576].X = 5.76
	pointsOfPolynomialPlot[576].Y = 5.196

	pointsOfPolynomialPlot[577].X = 5.77
	pointsOfPolynomialPlot[577].Y = 5.201

	pointsOfPolynomialPlot[578].X = 5.78
	pointsOfPolynomialPlot[578].Y = 5.205

	pointsOfPolynomialPlot[579].X = 5.79
	pointsOfPolynomialPlot[579].Y = 5.209

	pointsOfPolynomialPlot[580].X = 5.8
	pointsOfPolynomialPlot[580].Y = 5.213

	pointsOfPolynomialPlot[581].X = 5.81
	pointsOfPolynomialPlot[581].Y = 5.217

	pointsOfPolynomialPlot[582].X = 5.82
	pointsOfPolynomialPlot[582].Y = 5.221

	pointsOfPolynomialPlot[583].X = 5.83
	pointsOfPolynomialPlot[583].Y = 5.225

	pointsOfPolynomialPlot[584].X = 5.84
	pointsOfPolynomialPlot[584].Y = 5.228

	pointsOfPolynomialPlot[585].X = 5.85
	pointsOfPolynomialPlot[585].Y = 5.232

	pointsOfPolynomialPlot[586].X = 5.86
	pointsOfPolynomialPlot[586].Y = 5.235

	pointsOfPolynomialPlot[587].X = 5.87
	pointsOfPolynomialPlot[587].Y = 5.238

	pointsOfPolynomialPlot[588].X = 5.88
	pointsOfPolynomialPlot[588].Y = 5.241

	pointsOfPolynomialPlot[589].X = 5.89
	pointsOfPolynomialPlot[589].Y = 5.243

	pointsOfPolynomialPlot[590].X = 5.9
	pointsOfPolynomialPlot[590].Y = 5.246

	pointsOfPolynomialPlot[591].X = 5.91
	pointsOfPolynomialPlot[591].Y = 5.248

	pointsOfPolynomialPlot[592].X = 5.92
	pointsOfPolynomialPlot[592].Y = 5.25

	pointsOfPolynomialPlot[593].X = 5.93
	pointsOfPolynomialPlot[593].Y = 5.252

	pointsOfPolynomialPlot[594].X = 5.94
	pointsOfPolynomialPlot[594].Y = 5.254

	pointsOfPolynomialPlot[595].X = 5.95
	pointsOfPolynomialPlot[595].Y = 5.256

	pointsOfPolynomialPlot[596].X = 5.96
	pointsOfPolynomialPlot[596].Y = 5.257

	pointsOfPolynomialPlot[597].X = 5.97
	pointsOfPolynomialPlot[597].Y = 5.259

	pointsOfPolynomialPlot[598].X = 5.98
	pointsOfPolynomialPlot[598].Y = 5.26

	pointsOfPolynomialPlot[599].X = 5.99
	pointsOfPolynomialPlot[599].Y = 5.261

	pointsOfPolynomialPlot[600].X = 6.0
	pointsOfPolynomialPlot[600].Y = 5.262

	pointsOfPolynomialPlot[601].X = 6.01
	pointsOfPolynomialPlot[601].Y = 5.263

	pointsOfPolynomialPlot[602].X = 6.02
	pointsOfPolynomialPlot[602].Y = 5.264

	pointsOfPolynomialPlot[603].X = 6.03
	pointsOfPolynomialPlot[603].Y = 5.264

	pointsOfPolynomialPlot[604].X = 6.04
	pointsOfPolynomialPlot[604].Y = 5.264

	pointsOfPolynomialPlot[605].X = 6.05
	pointsOfPolynomialPlot[605].Y = 5.265

	pointsOfPolynomialPlot[606].X = 6.06
	pointsOfPolynomialPlot[606].Y = 5.265

	pointsOfPolynomialPlot[607].X = 6.07
	pointsOfPolynomialPlot[607].Y = 5.264

	pointsOfPolynomialPlot[608].X = 6.08
	pointsOfPolynomialPlot[608].Y = 5.264

	pointsOfPolynomialPlot[609].X = 6.09
	pointsOfPolynomialPlot[609].Y = 5.264

	pointsOfPolynomialPlot[610].X = 6.1
	pointsOfPolynomialPlot[610].Y = 5.263

	pointsOfPolynomialPlot[611].X = 6.11
	pointsOfPolynomialPlot[611].Y = 5.262

	pointsOfPolynomialPlot[612].X = 6.12
	pointsOfPolynomialPlot[612].Y = 5.261

	pointsOfPolynomialPlot[613].X = 6.13
	pointsOfPolynomialPlot[613].Y = 5.26

	pointsOfPolynomialPlot[614].X = 6.14
	pointsOfPolynomialPlot[614].Y = 5.259

	pointsOfPolynomialPlot[615].X = 6.15
	pointsOfPolynomialPlot[615].Y = 5.258

	pointsOfPolynomialPlot[616].X = 6.16
	pointsOfPolynomialPlot[616].Y = 5.256

	pointsOfPolynomialPlot[617].X = 6.17
	pointsOfPolynomialPlot[617].Y = 5.255

	pointsOfPolynomialPlot[618].X = 6.18
	pointsOfPolynomialPlot[618].Y = 5.253

	pointsOfPolynomialPlot[619].X = 6.19
	pointsOfPolynomialPlot[619].Y = 5.251

	pointsOfPolynomialPlot[620].X = 6.2
	pointsOfPolynomialPlot[620].Y = 5.249

	pointsOfPolynomialPlot[621].X = 6.21
	pointsOfPolynomialPlot[621].Y = 5.246

	pointsOfPolynomialPlot[622].X = 6.22
	pointsOfPolynomialPlot[622].Y = 5.244

	pointsOfPolynomialPlot[623].X = 6.23
	pointsOfPolynomialPlot[623].Y = 5.241

	pointsOfPolynomialPlot[624].X = 6.24
	pointsOfPolynomialPlot[624].Y = 5.238

	pointsOfPolynomialPlot[625].X = 6.25
	pointsOfPolynomialPlot[625].Y = 5.236

	pointsOfPolynomialPlot[626].X = 6.26
	pointsOfPolynomialPlot[626].Y = 5.232

	pointsOfPolynomialPlot[627].X = 6.27
	pointsOfPolynomialPlot[627].Y = 5.229

	pointsOfPolynomialPlot[628].X = 6.28
	pointsOfPolynomialPlot[628].Y = 5.226

	pointsOfPolynomialPlot[629].X = 6.29
	pointsOfPolynomialPlot[629].Y = 5.222

	pointsOfPolynomialPlot[630].X = 6.3
	pointsOfPolynomialPlot[630].Y = 5.219

	pointsOfPolynomialPlot[631].X = 6.31
	pointsOfPolynomialPlot[631].Y = 5.215

	pointsOfPolynomialPlot[632].X = 6.32
	pointsOfPolynomialPlot[632].Y = 5.211

	pointsOfPolynomialPlot[633].X = 6.33
	pointsOfPolynomialPlot[633].Y = 5.207

	pointsOfPolynomialPlot[634].X = 6.34
	pointsOfPolynomialPlot[634].Y = 5.203

	pointsOfPolynomialPlot[635].X = 6.35
	pointsOfPolynomialPlot[635].Y = 5.198

	pointsOfPolynomialPlot[636].X = 6.36
	pointsOfPolynomialPlot[636].Y = 5.194

	pointsOfPolynomialPlot[637].X = 6.37
	pointsOfPolynomialPlot[637].Y = 5.189

	pointsOfPolynomialPlot[638].X = 6.38
	pointsOfPolynomialPlot[638].Y = 5.184

	pointsOfPolynomialPlot[639].X = 6.39
	pointsOfPolynomialPlot[639].Y = 5.179

	pointsOfPolynomialPlot[640].X = 6.4
	pointsOfPolynomialPlot[640].Y = 5.174

	pointsOfPolynomialPlot[641].X = 6.41
	pointsOfPolynomialPlot[641].Y = 5.169

	pointsOfPolynomialPlot[642].X = 6.42
	pointsOfPolynomialPlot[642].Y = 5.163

	pointsOfPolynomialPlot[643].X = 6.43
	pointsOfPolynomialPlot[643].Y = 5.158

	pointsOfPolynomialPlot[644].X = 6.44
	pointsOfPolynomialPlot[644].Y = 5.152

	pointsOfPolynomialPlot[645].X = 6.45
	pointsOfPolynomialPlot[645].Y = 5.146

	pointsOfPolynomialPlot[646].X = 6.46
	pointsOfPolynomialPlot[646].Y = 5.14

	pointsOfPolynomialPlot[647].X = 6.47
	pointsOfPolynomialPlot[647].Y = 5.134

	pointsOfPolynomialPlot[648].X = 6.48
	pointsOfPolynomialPlot[648].Y = 5.128

	pointsOfPolynomialPlot[649].X = 6.49
	pointsOfPolynomialPlot[649].Y = 5.121

	pointsOfPolynomialPlot[650].X = 6.5
	pointsOfPolynomialPlot[650].Y = 5.115

	pointsOfPolynomialPlot[651].X = 6.51
	pointsOfPolynomialPlot[651].Y = 5.108

	pointsOfPolynomialPlot[652].X = 6.52
	pointsOfPolynomialPlot[652].Y = 5.101

	pointsOfPolynomialPlot[653].X = 6.53
	pointsOfPolynomialPlot[653].Y = 5.094

	pointsOfPolynomialPlot[654].X = 6.54
	pointsOfPolynomialPlot[654].Y = 5.087

	pointsOfPolynomialPlot[655].X = 6.55
	pointsOfPolynomialPlot[655].Y = 5.08

	pointsOfPolynomialPlot[656].X = 6.56
	pointsOfPolynomialPlot[656].Y = 5.072

	pointsOfPolynomialPlot[657].X = 6.57
	pointsOfPolynomialPlot[657].Y = 5.065

	pointsOfPolynomialPlot[658].X = 6.58
	pointsOfPolynomialPlot[658].Y = 5.057

	pointsOfPolynomialPlot[659].X = 6.59
	pointsOfPolynomialPlot[659].Y = 5.049

	pointsOfPolynomialPlot[660].X = 6.6
	pointsOfPolynomialPlot[660].Y = 5.041

	pointsOfPolynomialPlot[661].X = 6.61
	pointsOfPolynomialPlot[661].Y = 5.033

	pointsOfPolynomialPlot[662].X = 6.62
	pointsOfPolynomialPlot[662].Y = 5.025

	pointsOfPolynomialPlot[663].X = 6.63
	pointsOfPolynomialPlot[663].Y = 5.016

	pointsOfPolynomialPlot[664].X = 6.64
	pointsOfPolynomialPlot[664].Y = 5.008

	pointsOfPolynomialPlot[665].X = 6.65
	pointsOfPolynomialPlot[665].Y = 4.999

	pointsOfPolynomialPlot[666].X = 6.66
	pointsOfPolynomialPlot[666].Y = 4.99

	pointsOfPolynomialPlot[667].X = 6.67
	pointsOfPolynomialPlot[667].Y = 4.981

	pointsOfPolynomialPlot[668].X = 6.68
	pointsOfPolynomialPlot[668].Y = 4.972

	pointsOfPolynomialPlot[669].X = 6.69
	pointsOfPolynomialPlot[669].Y = 4.963

	pointsOfPolynomialPlot[670].X = 6.7
	pointsOfPolynomialPlot[670].Y = 4.954

	pointsOfPolynomialPlot[671].X = 6.71
	pointsOfPolynomialPlot[671].Y = 4.944

	pointsOfPolynomialPlot[672].X = 6.72
	pointsOfPolynomialPlot[672].Y = 4.935

	pointsOfPolynomialPlot[673].X = 6.73
	pointsOfPolynomialPlot[673].Y = 4.925

	pointsOfPolynomialPlot[674].X = 6.74
	pointsOfPolynomialPlot[674].Y = 4.915

	pointsOfPolynomialPlot[675].X = 6.75
	pointsOfPolynomialPlot[675].Y = 4.905

	pointsOfPolynomialPlot[676].X = 6.76
	pointsOfPolynomialPlot[676].Y = 4.895

	pointsOfPolynomialPlot[677].X = 6.77
	pointsOfPolynomialPlot[677].Y = 4.885

	pointsOfPolynomialPlot[678].X = 6.78
	pointsOfPolynomialPlot[678].Y = 4.874

	pointsOfPolynomialPlot[679].X = 6.79
	pointsOfPolynomialPlot[679].Y = 4.864

	pointsOfPolynomialPlot[680].X = 6.8
	pointsOfPolynomialPlot[680].Y = 4.853

	pointsOfPolynomialPlot[681].X = 6.81
	pointsOfPolynomialPlot[681].Y = 4.842

	pointsOfPolynomialPlot[682].X = 6.82
	pointsOfPolynomialPlot[682].Y = 4.831

	pointsOfPolynomialPlot[683].X = 6.83
	pointsOfPolynomialPlot[683].Y = 4.82

	pointsOfPolynomialPlot[684].X = 6.84
	pointsOfPolynomialPlot[684].Y = 4.809

	pointsOfPolynomialPlot[685].X = 6.85
	pointsOfPolynomialPlot[685].Y = 4.798

	pointsOfPolynomialPlot[686].X = 6.86
	pointsOfPolynomialPlot[686].Y = 4.787

	pointsOfPolynomialPlot[687].X = 6.87
	pointsOfPolynomialPlot[687].Y = 4.775

	pointsOfPolynomialPlot[688].X = 6.88
	pointsOfPolynomialPlot[688].Y = 4.763

	pointsOfPolynomialPlot[689].X = 6.89
	pointsOfPolynomialPlot[689].Y = 4.752

	pointsOfPolynomialPlot[690].X = 6.9
	pointsOfPolynomialPlot[690].Y = 4.74

	pointsOfPolynomialPlot[691].X = 6.91
	pointsOfPolynomialPlot[691].Y = 4.728

	pointsOfPolynomialPlot[692].X = 6.92
	pointsOfPolynomialPlot[692].Y = 4.716

	pointsOfPolynomialPlot[693].X = 6.93
	pointsOfPolynomialPlot[693].Y = 4.703

	pointsOfPolynomialPlot[694].X = 6.94
	pointsOfPolynomialPlot[694].Y = 4.691

	pointsOfPolynomialPlot[695].X = 6.95
	pointsOfPolynomialPlot[695].Y = 4.678

	pointsOfPolynomialPlot[696].X = 6.96
	pointsOfPolynomialPlot[696].Y = 4.666

	pointsOfPolynomialPlot[697].X = 6.97
	pointsOfPolynomialPlot[697].Y = 4.653

	pointsOfPolynomialPlot[698].X = 6.98
	pointsOfPolynomialPlot[698].Y = 4.64

	pointsOfPolynomialPlot[699].X = 6.99
	pointsOfPolynomialPlot[699].Y = 4.627

	pointsOfPolynomialPlot[700].X = 7.0
	pointsOfPolynomialPlot[700].Y = 4.614

	pointsOfPolynomialPlot[701].X = 7.01
	pointsOfPolynomialPlot[701].Y = 4.601

	pointsOfPolynomialPlot[702].X = 7.02
	pointsOfPolynomialPlot[702].Y = 4.588

	pointsOfPolynomialPlot[703].X = 7.03
	pointsOfPolynomialPlot[703].Y = 4.574

	pointsOfPolynomialPlot[704].X = 7.04
	pointsOfPolynomialPlot[704].Y = 4.561

	pointsOfPolynomialPlot[705].X = 7.05
	pointsOfPolynomialPlot[705].Y = 4.547

	pointsOfPolynomialPlot[706].X = 7.06
	pointsOfPolynomialPlot[706].Y = 4.533

	pointsOfPolynomialPlot[707].X = 7.07
	pointsOfPolynomialPlot[707].Y = 4.519

	pointsOfPolynomialPlot[708].X = 7.08
	pointsOfPolynomialPlot[708].Y = 4.505

	pointsOfPolynomialPlot[709].X = 7.09
	pointsOfPolynomialPlot[709].Y = 4.491

	pointsOfPolynomialPlot[710].X = 7.1
	pointsOfPolynomialPlot[710].Y = 4.477

	pointsOfPolynomialPlot[711].X = 7.11
	pointsOfPolynomialPlot[711].Y = 4.463

	pointsOfPolynomialPlot[712].X = 7.12
	pointsOfPolynomialPlot[712].Y = 4.448

	pointsOfPolynomialPlot[713].X = 7.13
	pointsOfPolynomialPlot[713].Y = 4.434

	pointsOfPolynomialPlot[714].X = 7.14
	pointsOfPolynomialPlot[714].Y = 4.419

	pointsOfPolynomialPlot[715].X = 7.15
	pointsOfPolynomialPlot[715].Y = 4.405

	pointsOfPolynomialPlot[716].X = 7.16
	pointsOfPolynomialPlot[716].Y = 4.39

	pointsOfPolynomialPlot[717].X = 7.17
	pointsOfPolynomialPlot[717].Y = 4.375

	pointsOfPolynomialPlot[718].X = 7.18
	pointsOfPolynomialPlot[718].Y = 4.36

	pointsOfPolynomialPlot[719].X = 7.19
	pointsOfPolynomialPlot[719].Y = 4.345

	pointsOfPolynomialPlot[720].X = 7.2
	pointsOfPolynomialPlot[720].Y = 4.329

	pointsOfPolynomialPlot[721].X = 7.21
	pointsOfPolynomialPlot[721].Y = 4.314

	pointsOfPolynomialPlot[722].X = 7.22
	pointsOfPolynomialPlot[722].Y = 4.299

	pointsOfPolynomialPlot[723].X = 7.23
	pointsOfPolynomialPlot[723].Y = 4.283

	pointsOfPolynomialPlot[724].X = 7.24
	pointsOfPolynomialPlot[724].Y = 4.267

	pointsOfPolynomialPlot[725].X = 7.25
	pointsOfPolynomialPlot[725].Y = 4.252

	pointsOfPolynomialPlot[726].X = 7.26
	pointsOfPolynomialPlot[726].Y = 4.236

	pointsOfPolynomialPlot[727].X = 7.27
	pointsOfPolynomialPlot[727].Y = 4.22

	pointsOfPolynomialPlot[728].X = 7.28
	pointsOfPolynomialPlot[728].Y = 4.204

	pointsOfPolynomialPlot[729].X = 7.29
	pointsOfPolynomialPlot[729].Y = 4.188

	pointsOfPolynomialPlot[730].X = 7.3
	pointsOfPolynomialPlot[730].Y = 4.172

	pointsOfPolynomialPlot[731].X = 7.31
	pointsOfPolynomialPlot[731].Y = 4.155

	pointsOfPolynomialPlot[732].X = 7.32
	pointsOfPolynomialPlot[732].Y = 4.139

	pointsOfPolynomialPlot[733].X = 7.33
	pointsOfPolynomialPlot[733].Y = 4.123

	pointsOfPolynomialPlot[734].X = 7.34
	pointsOfPolynomialPlot[734].Y = 4.106

	pointsOfPolynomialPlot[735].X = 7.35
	pointsOfPolynomialPlot[735].Y = 4.089

	pointsOfPolynomialPlot[736].X = 7.36
	pointsOfPolynomialPlot[736].Y = 4.073

	pointsOfPolynomialPlot[737].X = 7.37
	pointsOfPolynomialPlot[737].Y = 4.056

	pointsOfPolynomialPlot[738].X = 7.38
	pointsOfPolynomialPlot[738].Y = 4.039

	pointsOfPolynomialPlot[739].X = 7.39
	pointsOfPolynomialPlot[739].Y = 4.022

	pointsOfPolynomialPlot[740].X = 7.4
	pointsOfPolynomialPlot[740].Y = 4.005

	pointsOfPolynomialPlot[741].X = 7.41
	pointsOfPolynomialPlot[741].Y = 3.988

	pointsOfPolynomialPlot[742].X = 7.42
	pointsOfPolynomialPlot[742].Y = 3.97

	pointsOfPolynomialPlot[743].X = 7.43
	pointsOfPolynomialPlot[743].Y = 3.953

	pointsOfPolynomialPlot[744].X = 7.44
	pointsOfPolynomialPlot[744].Y = 3.936

	pointsOfPolynomialPlot[745].X = 7.45
	pointsOfPolynomialPlot[745].Y = 3.918

	pointsOfPolynomialPlot[746].X = 7.46
	pointsOfPolynomialPlot[746].Y = 3.901

	pointsOfPolynomialPlot[747].X = 7.47
	pointsOfPolynomialPlot[747].Y = 3.883

	pointsOfPolynomialPlot[748].X = 7.48
	pointsOfPolynomialPlot[748].Y = 3.865

	pointsOfPolynomialPlot[749].X = 7.49
	pointsOfPolynomialPlot[749].Y = 3.848

	pointsOfPolynomialPlot[750].X = 7.5
	pointsOfPolynomialPlot[750].Y = 3.83

	pointsOfPolynomialPlot[751].X = 7.51
	pointsOfPolynomialPlot[751].Y = 3.812

	pointsOfPolynomialPlot[752].X = 7.52
	pointsOfPolynomialPlot[752].Y = 3.794

	pointsOfPolynomialPlot[753].X = 7.53
	pointsOfPolynomialPlot[753].Y = 3.776

	pointsOfPolynomialPlot[754].X = 7.54
	pointsOfPolynomialPlot[754].Y = 3.758

	pointsOfPolynomialPlot[755].X = 7.55
	pointsOfPolynomialPlot[755].Y = 3.74

	pointsOfPolynomialPlot[756].X = 7.56
	pointsOfPolynomialPlot[756].Y = 3.721

	pointsOfPolynomialPlot[757].X = 7.57
	pointsOfPolynomialPlot[757].Y = 3.703

	pointsOfPolynomialPlot[758].X = 7.58
	pointsOfPolynomialPlot[758].Y = 3.685

	pointsOfPolynomialPlot[759].X = 7.59
	pointsOfPolynomialPlot[759].Y = 3.666

	pointsOfPolynomialPlot[760].X = 7.6
	pointsOfPolynomialPlot[760].Y = 3.648

	pointsOfPolynomialPlot[761].X = 7.61
	pointsOfPolynomialPlot[761].Y = 3.629

	pointsOfPolynomialPlot[762].X = 7.62
	pointsOfPolynomialPlot[762].Y = 3.611

	pointsOfPolynomialPlot[763].X = 7.63
	pointsOfPolynomialPlot[763].Y = 3.592

	pointsOfPolynomialPlot[764].X = 7.64
	pointsOfPolynomialPlot[764].Y = 3.573

	pointsOfPolynomialPlot[765].X = 7.65
	pointsOfPolynomialPlot[765].Y = 3.554

	pointsOfPolynomialPlot[766].X = 7.66
	pointsOfPolynomialPlot[766].Y = 3.535

	pointsOfPolynomialPlot[767].X = 7.67
	pointsOfPolynomialPlot[767].Y = 3.517

	pointsOfPolynomialPlot[768].X = 7.68
	pointsOfPolynomialPlot[768].Y = 3.498

	pointsOfPolynomialPlot[769].X = 7.69
	pointsOfPolynomialPlot[769].Y = 3.479

	pointsOfPolynomialPlot[770].X = 7.7
	pointsOfPolynomialPlot[770].Y = 3.459

	pointsOfPolynomialPlot[771].X = 7.71
	pointsOfPolynomialPlot[771].Y = 3.44

	pointsOfPolynomialPlot[772].X = 7.72
	pointsOfPolynomialPlot[772].Y = 3.421

	pointsOfPolynomialPlot[773].X = 7.73
	pointsOfPolynomialPlot[773].Y = 3.402

	pointsOfPolynomialPlot[774].X = 7.74
	pointsOfPolynomialPlot[774].Y = 3.382

	pointsOfPolynomialPlot[775].X = 7.75
	pointsOfPolynomialPlot[775].Y = 3.363

	pointsOfPolynomialPlot[776].X = 7.76
	pointsOfPolynomialPlot[776].Y = 3.344

	pointsOfPolynomialPlot[777].X = 7.77
	pointsOfPolynomialPlot[777].Y = 3.325

	pointsOfPolynomialPlot[778].X = 7.78
	pointsOfPolynomialPlot[778].Y = 3.305

	pointsOfPolynomialPlot[779].X = 7.79
	pointsOfPolynomialPlot[779].Y = 3.286

	pointsOfPolynomialPlot[780].X = 7.8
	pointsOfPolynomialPlot[780].Y = 3.266

	pointsOfPolynomialPlot[781].X = 7.81
	pointsOfPolynomialPlot[781].Y = 3.247

	pointsOfPolynomialPlot[782].X = 7.82
	pointsOfPolynomialPlot[782].Y = 3.227

	pointsOfPolynomialPlot[783].X = 7.83
	pointsOfPolynomialPlot[783].Y = 3.207

	pointsOfPolynomialPlot[784].X = 7.84
	pointsOfPolynomialPlot[784].Y = 3.188

	pointsOfPolynomialPlot[785].X = 7.85
	pointsOfPolynomialPlot[785].Y = 3.168

	pointsOfPolynomialPlot[786].X = 7.86
	pointsOfPolynomialPlot[786].Y = 3.148

	pointsOfPolynomialPlot[787].X = 7.87
	pointsOfPolynomialPlot[787].Y = 3.129

	pointsOfPolynomialPlot[788].X = 7.88
	pointsOfPolynomialPlot[788].Y = 3.109

	pointsOfPolynomialPlot[789].X = 7.89
	pointsOfPolynomialPlot[789].Y = 3.089

	pointsOfPolynomialPlot[790].X = 7.9
	pointsOfPolynomialPlot[790].Y = 3.069

	pointsOfPolynomialPlot[791].X = 7.91
	pointsOfPolynomialPlot[791].Y = 3.049

	pointsOfPolynomialPlot[792].X = 7.92
	pointsOfPolynomialPlot[792].Y = 3.029

	pointsOfPolynomialPlot[793].X = 7.93
	pointsOfPolynomialPlot[793].Y = 3.01

	pointsOfPolynomialPlot[794].X = 7.94
	pointsOfPolynomialPlot[794].Y = 2.99

	pointsOfPolynomialPlot[795].X = 7.95
	pointsOfPolynomialPlot[795].Y = 2.97

	pointsOfPolynomialPlot[796].X = 7.96
	pointsOfPolynomialPlot[796].Y = 2.95

	pointsOfPolynomialPlot[797].X = 7.97
	pointsOfPolynomialPlot[797].Y = 2.93

	pointsOfPolynomialPlot[798].X = 7.98
	pointsOfPolynomialPlot[798].Y = 2.91

	pointsOfPolynomialPlot[799].X = 7.99
	pointsOfPolynomialPlot[799].Y = 2.89

	pointsOfPolynomialPlot[800].X = 8.0
	pointsOfPolynomialPlot[800].Y = 2.87

	pointsOfPolynomialPlot[801].X = 8.01
	pointsOfPolynomialPlot[801].Y = 2.85

	pointsOfPolynomialPlot[802].X = 8.02
	pointsOfPolynomialPlot[802].Y = 2.83

	pointsOfPolynomialPlot[803].X = 8.03
	pointsOfPolynomialPlot[803].Y = 2.81

	pointsOfPolynomialPlot[804].X = 8.04
	pointsOfPolynomialPlot[804].Y = 2.789

	pointsOfPolynomialPlot[805].X = 8.05
	pointsOfPolynomialPlot[805].Y = 2.769

	pointsOfPolynomialPlot[806].X = 8.06
	pointsOfPolynomialPlot[806].Y = 2.749

	pointsOfPolynomialPlot[807].X = 8.07
	pointsOfPolynomialPlot[807].Y = 2.729

	pointsOfPolynomialPlot[808].X = 8.08
	pointsOfPolynomialPlot[808].Y = 2.709

	pointsOfPolynomialPlot[809].X = 8.09
	pointsOfPolynomialPlot[809].Y = 2.689

	pointsOfPolynomialPlot[810].X = 8.1
	pointsOfPolynomialPlot[810].Y = 2.669

	pointsOfPolynomialPlot[811].X = 8.11
	pointsOfPolynomialPlot[811].Y = 2.649

	pointsOfPolynomialPlot[812].X = 8.12
	pointsOfPolynomialPlot[812].Y = 2.629

	pointsOfPolynomialPlot[813].X = 8.13
	pointsOfPolynomialPlot[813].Y = 2.609

	pointsOfPolynomialPlot[814].X = 8.14
	pointsOfPolynomialPlot[814].Y = 2.589

	pointsOfPolynomialPlot[815].X = 8.15
	pointsOfPolynomialPlot[815].Y = 2.569

	pointsOfPolynomialPlot[816].X = 8.16
	pointsOfPolynomialPlot[816].Y = 2.549

	pointsOfPolynomialPlot[817].X = 8.17
	pointsOfPolynomialPlot[817].Y = 2.529

	pointsOfPolynomialPlot[818].X = 8.18
	pointsOfPolynomialPlot[818].Y = 2.509

	pointsOfPolynomialPlot[819].X = 8.19
	pointsOfPolynomialPlot[819].Y = 2.489

	pointsOfPolynomialPlot[820].X = 8.2
	pointsOfPolynomialPlot[820].Y = 2.469

	pointsOfPolynomialPlot[821].X = 8.21
	pointsOfPolynomialPlot[821].Y = 2.449

	pointsOfPolynomialPlot[822].X = 8.22
	pointsOfPolynomialPlot[822].Y = 2.429

	pointsOfPolynomialPlot[823].X = 8.23
	pointsOfPolynomialPlot[823].Y = 2.409

	pointsOfPolynomialPlot[824].X = 8.24
	pointsOfPolynomialPlot[824].Y = 2.389

	pointsOfPolynomialPlot[825].X = 8.25
	pointsOfPolynomialPlot[825].Y = 2.369

	pointsOfPolynomialPlot[826].X = 8.26
	pointsOfPolynomialPlot[826].Y = 2.349

	pointsOfPolynomialPlot[827].X = 8.27
	pointsOfPolynomialPlot[827].Y = 2.329

	pointsOfPolynomialPlot[828].X = 8.28
	pointsOfPolynomialPlot[828].Y = 2.309

	pointsOfPolynomialPlot[829].X = 8.29
	pointsOfPolynomialPlot[829].Y = 2.29

	pointsOfPolynomialPlot[830].X = 8.3
	pointsOfPolynomialPlot[830].Y = 2.27

	pointsOfPolynomialPlot[831].X = 8.31
	pointsOfPolynomialPlot[831].Y = 2.25

	pointsOfPolynomialPlot[832].X = 8.32
	pointsOfPolynomialPlot[832].Y = 2.23

	pointsOfPolynomialPlot[833].X = 8.33
	pointsOfPolynomialPlot[833].Y = 2.211

	pointsOfPolynomialPlot[834].X = 8.34
	pointsOfPolynomialPlot[834].Y = 2.191

	pointsOfPolynomialPlot[835].X = 8.35
	pointsOfPolynomialPlot[835].Y = 2.172

	pointsOfPolynomialPlot[836].X = 8.36
	pointsOfPolynomialPlot[836].Y = 2.152

	pointsOfPolynomialPlot[837].X = 8.37
	pointsOfPolynomialPlot[837].Y = 2.133

	pointsOfPolynomialPlot[838].X = 8.38
	pointsOfPolynomialPlot[838].Y = 2.133

	pointsOfPolynomialPlot[839].X = 8.39
	pointsOfPolynomialPlot[839].Y = 2.113

	pointsOfPolynomialPlot[840].X = 8.4
	pointsOfPolynomialPlot[840].Y = 2.075

	pointsOfPolynomialPlot[841].X = 8.41
	pointsOfPolynomialPlot[841].Y = 2.055

	pointsOfPolynomialPlot[842].X = 8.42
	pointsOfPolynomialPlot[842].Y = 2.036

	pointsOfPolynomialPlot[843].X = 8.43
	pointsOfPolynomialPlot[843].Y = 2.017

	pointsOfPolynomialPlot[844].X = 8.44
	pointsOfPolynomialPlot[844].Y = 1.998

	pointsOfPolynomialPlot[845].X = 8.45
	pointsOfPolynomialPlot[845].Y = 1.979

	pointsOfPolynomialPlot[846].X = 8.46
	pointsOfPolynomialPlot[846].Y = 1.96

	pointsOfPolynomialPlot[847].X = 8.47
	pointsOfPolynomialPlot[847].Y = 1.941

	pointsOfPolynomialPlot[848].X = 8.48
	pointsOfPolynomialPlot[848].Y = 1.922

	pointsOfPolynomialPlot[849].X = 8.49
	pointsOfPolynomialPlot[849].Y = 1.903

	pointsOfPolynomialPlot[850].X = 8.5
	pointsOfPolynomialPlot[850].Y = 1.884

	pointsOfPolynomialPlot[851].X = 8.51
	pointsOfPolynomialPlot[851].Y = 1.865

	pointsOfPolynomialPlot[852].X = 8.52
	pointsOfPolynomialPlot[852].Y = 1.847

	pointsOfPolynomialPlot[853].X = 8.53
	pointsOfPolynomialPlot[853].Y = 1.828

	pointsOfPolynomialPlot[854].X = 8.54
	pointsOfPolynomialPlot[854].Y = 1.81

	pointsOfPolynomialPlot[855].X = 8.55
	pointsOfPolynomialPlot[855].Y = 1.791

	pointsOfPolynomialPlot[856].X = 8.56
	pointsOfPolynomialPlot[856].Y = 1.773

	pointsOfPolynomialPlot[857].X = 8.57
	pointsOfPolynomialPlot[857].Y = 1.755

	pointsOfPolynomialPlot[858].X = 8.58
	pointsOfPolynomialPlot[858].Y = 1.737

	pointsOfPolynomialPlot[859].X = 8.59
	pointsOfPolynomialPlot[859].Y = 1.718

	pointsOfPolynomialPlot[860].X = 8.6
	pointsOfPolynomialPlot[860].Y = 1.7

	pointsOfPolynomialPlot[861].X = 8.61
	pointsOfPolynomialPlot[861].Y = 1.682

	pointsOfPolynomialPlot[862].X = 8.62
	pointsOfPolynomialPlot[862].Y = 1.665

	pointsOfPolynomialPlot[863].X = 8.63
	pointsOfPolynomialPlot[863].Y = 1.647

	pointsOfPolynomialPlot[864].X = 8.64
	pointsOfPolynomialPlot[864].Y = 1.629

	pointsOfPolynomialPlot[865].X = 8.65
	pointsOfPolynomialPlot[865].Y = 1.612

	pointsOfPolynomialPlot[866].X = 8.66
	pointsOfPolynomialPlot[866].Y = 1.594

	pointsOfPolynomialPlot[867].X = 8.67
	pointsOfPolynomialPlot[867].Y = 1.577

	pointsOfPolynomialPlot[868].X = 8.68
	pointsOfPolynomialPlot[868].Y = 1.559

	pointsOfPolynomialPlot[869].X = 8.69
	pointsOfPolynomialPlot[869].Y = 1.542

	pointsOfPolynomialPlot[870].X = 8.7
	pointsOfPolynomialPlot[870].Y = 1.525

	pointsOfPolynomialPlot[871].X = 8.71
	pointsOfPolynomialPlot[871].Y = 1.508

	pointsOfPolynomialPlot[872].X = 8.72
	pointsOfPolynomialPlot[872].Y = 1.491

	pointsOfPolynomialPlot[873].X = 8.73
	pointsOfPolynomialPlot[873].Y = 1.475

	pointsOfPolynomialPlot[874].X = 8.74
	pointsOfPolynomialPlot[874].Y = 1.458

	pointsOfPolynomialPlot[875].X = 8.75
	pointsOfPolynomialPlot[875].Y = 1.441

	pointsOfPolynomialPlot[876].X = 8.76
	pointsOfPolynomialPlot[876].Y = 1.425

	pointsOfPolynomialPlot[877].X = 8.77
	pointsOfPolynomialPlot[877].Y = 1.409

	pointsOfPolynomialPlot[878].X = 8.78
	pointsOfPolynomialPlot[878].Y = 1.392

	pointsOfPolynomialPlot[879].X = 8.79
	pointsOfPolynomialPlot[879].Y = 1.376

	pointsOfPolynomialPlot[880].X = 8.8
	pointsOfPolynomialPlot[880].Y = 1.36

	pointsOfPolynomialPlot[881].X = 8.81
	pointsOfPolynomialPlot[881].Y = 1.345

	pointsOfPolynomialPlot[882].X = 8.82
	pointsOfPolynomialPlot[882].Y = 1.329

	pointsOfPolynomialPlot[883].X = 8.83
	pointsOfPolynomialPlot[883].Y = 1.313

	pointsOfPolynomialPlot[884].X = 8.84
	pointsOfPolynomialPlot[884].Y = 1.298

	pointsOfPolynomialPlot[885].X = 8.85
	pointsOfPolynomialPlot[885].Y = 1.283

	pointsOfPolynomialPlot[886].X = 8.86
	pointsOfPolynomialPlot[886].Y = 1.267

	pointsOfPolynomialPlot[887].X = 8.87
	pointsOfPolynomialPlot[887].Y = 1.252

	pointsOfPolynomialPlot[888].X = 8.88
	pointsOfPolynomialPlot[888].Y = 1.237

	pointsOfPolynomialPlot[889].X = 8.89
	pointsOfPolynomialPlot[889].Y = 1.223

	pointsOfPolynomialPlot[890].X = 8.9
	pointsOfPolynomialPlot[890].Y = 1.208

	pointsOfPolynomialPlot[891].X = 8.91
	pointsOfPolynomialPlot[891].Y = 1.194

	pointsOfPolynomialPlot[892].X = 8.92
	pointsOfPolynomialPlot[892].Y = 1.179

	pointsOfPolynomialPlot[893].X = 8.93
	pointsOfPolynomialPlot[893].Y = 1.165

	pointsOfPolynomialPlot[894].X = 8.94
	pointsOfPolynomialPlot[894].Y = 1.151

	pointsOfPolynomialPlot[895].X = 8.95
	pointsOfPolynomialPlot[895].Y = 1.137

	pointsOfPolynomialPlot[896].X = 8.96
	pointsOfPolynomialPlot[896].Y = 1.123

	pointsOfPolynomialPlot[897].X = 8.97
	pointsOfPolynomialPlot[897].Y = 1.11

	pointsOfPolynomialPlot[898].X = 8.98
	pointsOfPolynomialPlot[898].Y = 1.096

	pointsOfPolynomialPlot[899].X = 8.99
	pointsOfPolynomialPlot[899].Y = 1.083

	pointsOfPolynomialPlot[900].X = 9.0
	pointsOfPolynomialPlot[900].Y = 1.07

	pointsOfPolynomialPlot[901].X = 9.01
	pointsOfPolynomialPlot[901].Y = 1.057

	pointsOfPolynomialPlot[902].X = 9.02
	pointsOfPolynomialPlot[902].Y = 1.044

	pointsOfPolynomialPlot[903].X = 9.03
	pointsOfPolynomialPlot[903].Y = 1.032

	pointsOfPolynomialPlot[904].X = 9.04
	pointsOfPolynomialPlot[904].Y = 1.019

	pointsOfPolynomialPlot[905].X = 9.05
	pointsOfPolynomialPlot[905].Y = 1.007

	pointsOfPolynomialPlot[906].X = 9.06
	pointsOfPolynomialPlot[906].Y = 0.995

	pointsOfPolynomialPlot[907].X = 9.07
	pointsOfPolynomialPlot[907].Y = 0.983

	pointsOfPolynomialPlot[908].X = 9.08
	pointsOfPolynomialPlot[908].Y = 0.971

	pointsOfPolynomialPlot[909].X = 9.09
	pointsOfPolynomialPlot[909].Y = 0.96

	pointsOfPolynomialPlot[910].X = 9.1
	pointsOfPolynomialPlot[910].Y = 0.949

	pointsOfPolynomialPlot[911].X = 9.11
	pointsOfPolynomialPlot[911].Y = 0.937

	pointsOfPolynomialPlot[912].X = 9.12
	pointsOfPolynomialPlot[912].Y = 0.926

	pointsOfPolynomialPlot[913].X = 9.13
	pointsOfPolynomialPlot[913].Y = 0.916

	pointsOfPolynomialPlot[914].X = 9.14
	pointsOfPolynomialPlot[914].Y = 0.905

	pointsOfPolynomialPlot[915].X = 9.15
	pointsOfPolynomialPlot[915].Y = 0.895

	pointsOfPolynomialPlot[916].X = 9.16
	pointsOfPolynomialPlot[916].Y = 0.885

	pointsOfPolynomialPlot[917].X = 9.17
	pointsOfPolynomialPlot[917].Y = 0.875

	pointsOfPolynomialPlot[918].X = 9.18
	pointsOfPolynomialPlot[918].Y = 0.865

	pointsOfPolynomialPlot[919].X = 9.19
	pointsOfPolynomialPlot[919].Y = 0.855

	pointsOfPolynomialPlot[920].X = 9.2
	pointsOfPolynomialPlot[920].Y = 0.846

	pointsOfPolynomialPlot[921].X = 9.21
	pointsOfPolynomialPlot[921].Y = 0.837

	pointsOfPolynomialPlot[922].X = 9.22
	pointsOfPolynomialPlot[922].Y = 0.828

	pointsOfPolynomialPlot[923].X = 9.23
	pointsOfPolynomialPlot[923].Y = 0.819

	pointsOfPolynomialPlot[924].X = 9.24
	pointsOfPolynomialPlot[924].Y = 0.811

	pointsOfPolynomialPlot[925].X = 9.25
	pointsOfPolynomialPlot[925].Y = 0.802

	pointsOfPolynomialPlot[926].X = 9.26
	pointsOfPolynomialPlot[926].Y = 0.794

	pointsOfPolynomialPlot[927].X = 9.27
	pointsOfPolynomialPlot[927].Y = 0.787

	pointsOfPolynomialPlot[928].X = 9.28
	pointsOfPolynomialPlot[928].Y = 0.779

	pointsOfPolynomialPlot[929].X = 9.29
	pointsOfPolynomialPlot[929].Y = 0.772

	pointsOfPolynomialPlot[930].X = 9.3
	pointsOfPolynomialPlot[930].Y = 0.764

	pointsOfPolynomialPlot[931].X = 9.31
	pointsOfPolynomialPlot[931].Y = 0.758

	pointsOfPolynomialPlot[932].X = 9.32
	pointsOfPolynomialPlot[932].Y = 0.751

	pointsOfPolynomialPlot[933].X = 9.33
	pointsOfPolynomialPlot[933].Y = 0.745

	pointsOfPolynomialPlot[934].X = 9.34
	pointsOfPolynomialPlot[934].Y = 0.738

	pointsOfPolynomialPlot[935].X = 9.35
	pointsOfPolynomialPlot[935].Y = 0.732

	pointsOfPolynomialPlot[936].X = 9.36
	pointsOfPolynomialPlot[936].Y = 0.727

	pointsOfPolynomialPlot[937].X = 9.37
	pointsOfPolynomialPlot[937].Y = 0.721

	pointsOfPolynomialPlot[938].X = 9.38
	pointsOfPolynomialPlot[938].Y = 0.716

	pointsOfPolynomialPlot[939].X = 9.39
	pointsOfPolynomialPlot[939].Y = 0.711

	pointsOfPolynomialPlot[940].X = 9.4
	pointsOfPolynomialPlot[940].Y = 0.707

	pointsOfPolynomialPlot[941].X = 9.41
	pointsOfPolynomialPlot[941].Y = 0.702

	pointsOfPolynomialPlot[942].X = 9.42
	pointsOfPolynomialPlot[942].Y = 0.698

	pointsOfPolynomialPlot[943].X = 9.43
	pointsOfPolynomialPlot[943].Y = 0.694

	pointsOfPolynomialPlot[944].X = 9.44
	pointsOfPolynomialPlot[944].Y = 0.69

	pointsOfPolynomialPlot[945].X = 9.45
	pointsOfPolynomialPlot[945].Y  = 0.687

	pointsOfPolynomialPlot[946].X = 9.46
	pointsOfPolynomialPlot[946].Y = 0.684

	pointsOfPolynomialPlot[947].X = 9.47
	pointsOfPolynomialPlot[947].Y = 0.681

	pointsOfPolynomialPlot[948].X = 9.48
	pointsOfPolynomialPlot[948].Y = 0.679

	pointsOfPolynomialPlot[949].X = 9.49
	pointsOfPolynomialPlot[949].Y = 0.677

	pointsOfPolynomialPlot[950].X = 9.5
	pointsOfPolynomialPlot[950].Y = 0.675

	pointsOfPolynomialPlot[951].X = 9.51
	pointsOfPolynomialPlot[951].Y = 0.673

	pointsOfPolynomialPlot[952].X = 9.52
	pointsOfPolynomialPlot[952].Y = 0.672

	pointsOfPolynomialPlot[953].X = 9.53
	pointsOfPolynomialPlot[953].Y = 0.67

	pointsOfPolynomialPlot[954].X = 9.54
	pointsOfPolynomialPlot[954].Y = 0.67

	pointsOfPolynomialPlot[955].X = 9.55
	pointsOfPolynomialPlot[955].Y = 0.669

	pointsOfPolynomialPlot[956].X = 9.56
	pointsOfPolynomialPlot[956].Y = 0.669

	pointsOfPolynomialPlot[957].X = 9.57
	pointsOfPolynomialPlot[957].Y = 0.669

	pointsOfPolynomialPlot[958].X = 9.58
	pointsOfPolynomialPlot[958].Y = 0.67

	pointsOfPolynomialPlot[959].X = 9.59
	pointsOfPolynomialPlot[959].Y = 0.67

	pointsOfPolynomialPlot[960].X = 9.6
	pointsOfPolynomialPlot[960].Y = 0.671

	pointsOfPolynomialPlot[961].X = 9.61
	pointsOfPolynomialPlot[961].Y = 0.673

	pointsOfPolynomialPlot[962].X = 9.62
	pointsOfPolynomialPlot[962].Y = 0.674

	pointsOfPolynomialPlot[963].X = 9.63
	pointsOfPolynomialPlot[963].Y = 0.679

	pointsOfPolynomialPlot[964].X = 9.64
	pointsOfPolynomialPlot[964].Y = 0.681

	pointsOfPolynomialPlot[965].X = 9.65
	pointsOfPolynomialPlot[965].Y = 0.681

	pointsOfPolynomialPlot[966].X = 9.66
	pointsOfPolynomialPlot[966].Y = 0.684

	pointsOfPolynomialPlot[967].X = 9.67
	pointsOfPolynomialPlot[967].Y = 0.687

	pointsOfPolynomialPlot[968].X = 9.68
	pointsOfPolynomialPlot[968].Y = 0.691

	pointsOfPolynomialPlot[969].X = 9.69
	pointsOfPolynomialPlot[969].Y = 0.695

	pointsOfPolynomialPlot[970].X = 9.7
	pointsOfPolynomialPlot[970].Y = 0.699

	pointsOfPolynomialPlot[971].X = 9.71
	pointsOfPolynomialPlot[971].Y = 0.704

	pointsOfPolynomialPlot[972].X = 9.72
	pointsOfPolynomialPlot[972].Y = 0.709

	pointsOfPolynomialPlot[973].X = 9.73
	pointsOfPolynomialPlot[973].Y = 0.714

	pointsOfPolynomialPlot[974].X = 9.74
	pointsOfPolynomialPlot[974].Y = 0.72

	pointsOfPolynomialPlot[975].X = 9.75
	pointsOfPolynomialPlot[975].Y = 0.726

	pointsOfPolynomialPlot[976].X = 9.76
	pointsOfPolynomialPlot[976].Y = 0.732

	pointsOfPolynomialPlot[977].X = 9.77
	pointsOfPolynomialPlot[977].Y = 0.739

	pointsOfPolynomialPlot[978].X = 9.78
	pointsOfPolynomialPlot[978].Y = 0.746

	pointsOfPolynomialPlot[979].X = 9.79
	pointsOfPolynomialPlot[979].Y = 0.754

	pointsOfPolynomialPlot[980].X = 9.8
	pointsOfPolynomialPlot[980].Y = 0.761

	pointsOfPolynomialPlot[981].X = 9.81
	pointsOfPolynomialPlot[981].Y = 0.77

	pointsOfPolynomialPlot[982].X = 9.82
	pointsOfPolynomialPlot[982].Y = 0.778

	pointsOfPolynomialPlot[983].X = 9.83
	pointsOfPolynomialPlot[983].Y = 0.787

	pointsOfPolynomialPlot[984].X = 9.84
	pointsOfPolynomialPlot[984].Y = 0.796

	pointsOfPolynomialPlot[985].X = 9.85
	pointsOfPolynomialPlot[985].Y = 0.806

	pointsOfPolynomialPlot[986].X = 9.86
	pointsOfPolynomialPlot[986].Y = 0.816

	pointsOfPolynomialPlot[987].X = 9.87
	pointsOfPolynomialPlot[987].Y = 0.827

	pointsOfPolynomialPlot[988].X = 9.88
	pointsOfPolynomialPlot[988].Y = 0.838

	pointsOfPolynomialPlot[989].X = 9.89
	pointsOfPolynomialPlot[989].Y = 0.849

	pointsOfPolynomialPlot[990].X = 9.9
	pointsOfPolynomialPlot[990].Y = 0.861

	pointsOfPolynomialPlot[991].X = 9.91
	pointsOfPolynomialPlot[991].Y = 0.873

	pointsOfPolynomialPlot[992].X = 9.92
	pointsOfPolynomialPlot[992].Y = 0.885

	pointsOfPolynomialPlot[993].X = 9.93
	pointsOfPolynomialPlot[993].Y = 0.898

	pointsOfPolynomialPlot[994].X = 9.94
	pointsOfPolynomialPlot[994].Y = 0.911

	pointsOfPolynomialPlot[995].X = 9.95
	pointsOfPolynomialPlot[995].Y = 0.925

	pointsOfPolynomialPlot[996].X = 9.96
	pointsOfPolynomialPlot[996].Y = 0.939

	pointsOfPolynomialPlot[997].X = 9.97
	pointsOfPolynomialPlot[997].Y = 0.953

	pointsOfPolynomialPlot[998].X = 9.98
	pointsOfPolynomialPlot[998].Y = 0.968

	pointsOfPolynomialPlot[999].X = 9.99
	pointsOfPolynomialPlot[999].Y = 0.984

	pointsOfPolynomialPlot[1_000].X = 10.0
	pointsOfPolynomialPlot[1_000].Y = 1.0

	pointsOfPolynomialPlot[1_001].X = 10.01
	pointsOfPolynomialPlot[1_001].Y = 1.016

	pointsOfPolynomialPlot[1_002].X = 10.02
	pointsOfPolynomialPlot[1_002].Y = 1.032

	pointsOfPolynomialPlot[1_003].X = 10.03
	pointsOfPolynomialPlot[1_003].Y = 1.049

	pointsOfPolynomialPlot[1_004].X = 10.04
	pointsOfPolynomialPlot[1_004].Y = 1.067

	pointsOfPolynomialPlot[1_005].X = 10.05
	pointsOfPolynomialPlot[1_005].Y = 1.085

	pointsOfPolynomialPlot[1_006].X = 10.06
	pointsOfPolynomialPlot[1_006].Y = 1.103

	pointsOfPolynomialPlot[1_007].X = 10.07
	pointsOfPolynomialPlot[1_007].Y = 1.122

	pointsOfPolynomialPlot[1_008].X = 10.08
	pointsOfPolynomialPlot[1_008].Y = 1.142

	pointsOfPolynomialPlot[1_009].X = 10.09
	pointsOfPolynomialPlot[1_009].Y = 1.161

	pointsOfPolynomialPlot[1_010].X = 10.1
	pointsOfPolynomialPlot[1_010].Y = 1.182

	pointsOfPolynomialPlot[1_011].X = 10.11
	pointsOfPolynomialPlot[1_011].Y = 1.202

	pointsOfPolynomialPlot[1_012].X = 10.12
	pointsOfPolynomialPlot[1_012].Y = 1.223

	pointsOfPolynomialPlot[1_013].X = 10.13
	pointsOfPolynomialPlot[1_013].Y = 1.245

	pointsOfPolynomialPlot[1_014].X = 10.14
	pointsOfPolynomialPlot[1_014].Y = 1.267

	pointsOfPolynomialPlot[1_015].X = 10.15
	pointsOfPolynomialPlot[1_015].Y = 1.29

	pointsOfPolynomialPlot[1_016].X = 10.16
	pointsOfPolynomialPlot[1_016].Y = 1.313

	pointsOfPolynomialPlot[1_017].X = 10.17
	pointsOfPolynomialPlot[1_017].Y = 1.336

	pointsOfPolynomialPlot[1_018].X = 10.18
	pointsOfPolynomialPlot[1_018].Y = 1.36

	pointsOfPolynomialPlot[1_019].X = 10.19
	pointsOfPolynomialPlot[1_019].Y = 1.385

	pointsOfPolynomialPlot[1_020].X = 10.2
	pointsOfPolynomialPlot[1_020].Y = 1.41

	pointsOfPolynomialPlot[1_021].X = 10.21
	pointsOfPolynomialPlot[1_021].Y = 1.435

	pointsOfPolynomialPlot[1_022].X = 10.22
	pointsOfPolynomialPlot[1_022].Y = 1.461

	pointsOfPolynomialPlot[1_023].X = 10.23
	pointsOfPolynomialPlot[1_023].Y = 1.488

	pointsOfPolynomialPlot[1_024].X = 10.24
	pointsOfPolynomialPlot[1_024].Y = 1.515

	pointsOfPolynomialPlot[1_025].X = 10.25
	pointsOfPolynomialPlot[1_025].Y = 1.542

	pointsOfPolynomialPlot[1_026].X = 10.26
	pointsOfPolynomialPlot[1_026].Y = 1.57

	pointsOfPolynomialPlot[1_027].X = 10.27
	pointsOfPolynomialPlot[1_027].Y = 1.599

	pointsOfPolynomialPlot[1_028].X = 10.28
	pointsOfPolynomialPlot[1_028].Y = 1.628

	pointsOfPolynomialPlot[1_029].X = 10.29
	pointsOfPolynomialPlot[1_029].Y = 1.657

	pointsOfPolynomialPlot[1_030].X = 10.3
	pointsOfPolynomialPlot[1_030].Y = 1.687

	pointsOfPolynomialPlot[1_031].X = 10.31
	pointsOfPolynomialPlot[1_031].Y = 1.718

	pointsOfPolynomialPlot[1_032].X = 10.32
	pointsOfPolynomialPlot[1_032].Y = 1.749

	pointsOfPolynomialPlot[1_033].X = 10.33
	pointsOfPolynomialPlot[1_033].Y = 1.781

	pointsOfPolynomialPlot[1_034].X = 10.34
	pointsOfPolynomialPlot[1_034].Y = 1.813

	pointsOfPolynomialPlot[1_035].X = 10.35
	pointsOfPolynomialPlot[1_035].Y = 1.846

	pointsOfPolynomialPlot[1_036].X = 10.36
	pointsOfPolynomialPlot[1_036].Y = 1.879

	pointsOfPolynomialPlot[1_037].X = 10.37
	pointsOfPolynomialPlot[1_037].Y = 1.913

	pointsOfPolynomialPlot[1_038].X = 10.38
	pointsOfPolynomialPlot[1_038].Y = 1.947

	pointsOfPolynomialPlot[1_039].X = 10.39
	pointsOfPolynomialPlot[1_039].Y = 1.982

	pointsOfPolynomialPlot[1_040].X = 10.4
	pointsOfPolynomialPlot[1_040].Y = 2.018

	pointsOfPolynomialPlot[1_041].X = 10.41
	pointsOfPolynomialPlot[1_041].Y = 2.054

	pointsOfPolynomialPlot[1_042].X = 10.42
	pointsOfPolynomialPlot[1_042].Y = 2.091

	pointsOfPolynomialPlot[1_043].X = 10.43
	pointsOfPolynomialPlot[1_043].Y = 2.128

	pointsOfPolynomialPlot[1_044].X = 10.44
	pointsOfPolynomialPlot[1_044].Y = 2.166

	pointsOfPolynomialPlot[1_045].X = 10.45
	pointsOfPolynomialPlot[1_045].Y = 2.204

	pointsOfPolynomialPlot[1_046].X = 10.46
	pointsOfPolynomialPlot[1_046].Y = 2.243

	pointsOfPolynomialPlot[1_047].X = 10.47
	pointsOfPolynomialPlot[1_047].Y = 2.283

	pointsOfPolynomialPlot[1_048].X = 10.48
	pointsOfPolynomialPlot[1_048].Y = 2.323

	pointsOfPolynomialPlot[1_049].X = 10.49
	pointsOfPolynomialPlot[1_049].Y = 2.363

	pointsOfPolynomialPlot[1_050].X = 10.5
	pointsOfPolynomialPlot[1_050].Y = 2.405

	pointsOfPolynomialPlot[1_051].X = 10.51
	pointsOfPolynomialPlot[1_051].Y = 2.447

	pointsOfPolynomialPlot[1_052].X = 10.52
	pointsOfPolynomialPlot[1_052].Y = 2.489

	pointsOfPolynomialPlot[1_053].X = 10.53
	pointsOfPolynomialPlot[1_053].Y = 2.532

	pointsOfPolynomialPlot[1_054].X = 10.54
	pointsOfPolynomialPlot[1_054].Y = 2.576

	pointsOfPolynomialPlot[1_055].X = 10.55
	pointsOfPolynomialPlot[1_055].Y = 2.62

	pointsOfPolynomialPlot[1_056].X = 10.56
	pointsOfPolynomialPlot[1_056].Y = 2.665

	pointsOfPolynomialPlot[1_057].X = 10.57
	pointsOfPolynomialPlot[1_057].Y = 2.711

	pointsOfPolynomialPlot[1_058].X = 10.58
	pointsOfPolynomialPlot[1_058].Y = 2.757

	pointsOfPolynomialPlot[1_059].X = 10.59
	pointsOfPolynomialPlot[1_059].Y = 2.804

	pointsOfPolynomialPlot[1_060].X = 10.6
	pointsOfPolynomialPlot[1_060].Y = 2.851

	pointsOfPolynomialPlot[1_061].X = 10.61
	pointsOfPolynomialPlot[1_061].Y = 2.899

	pointsOfPolynomialPlot[1_062].X = 10.62
	pointsOfPolynomialPlot[1_062].Y = 2.948

	pointsOfPolynomialPlot[1_063].X = 10.63
	pointsOfPolynomialPlot[1_063].Y = 2.998

	pointsOfPolynomialPlot[1_064].X = 10.64
	pointsOfPolynomialPlot[1_064].Y = 3.048

	pointsOfPolynomialPlot[1_065].X = 10.65
	pointsOfPolynomialPlot[1_065].Y = 3.098

	pointsOfPolynomialPlot[1_066].X = 10.66
	pointsOfPolynomialPlot[1_066].Y = 3.15

	pointsOfPolynomialPlot[1_067].X = 10.67
	pointsOfPolynomialPlot[1_067].Y = 3.202

	pointsOfPolynomialPlot[1_068].X = 10.68
	pointsOfPolynomialPlot[1_068].Y = 3.254

	pointsOfPolynomialPlot[1_069].X = 10.69
	pointsOfPolynomialPlot[1_069].Y = 3.308

	pointsOfPolynomialPlot[1_070].X = 10.7
	pointsOfPolynomialPlot[1_070].Y = 3.362

	pointsOfPolynomialPlot[1_071].X = 10.71
	pointsOfPolynomialPlot[1_071].Y = 3.416

	pointsOfPolynomialPlot[1_072].X = 10.72
	pointsOfPolynomialPlot[1_072].Y = 3.472

	pointsOfPolynomialPlot[1_073].X = 10.73
	pointsOfPolynomialPlot[1_073].Y = 3.528

	pointsOfPolynomialPlot[1_074].X = 10.74
	pointsOfPolynomialPlot[1_074].Y = 3.584

	pointsOfPolynomialPlot[1_075].X = 10.75
	pointsOfPolynomialPlot[1_075].Y = 3.642

	pointsOfPolynomialPlot[1_076].X = 10.76
	pointsOfPolynomialPlot[1_076].Y = 3.7

	pointsOfPolynomialPlot[1_077].X = 10.77
	pointsOfPolynomialPlot[1_077].Y = 3.759

	pointsOfPolynomialPlot[1_078].X = 10.78
	pointsOfPolynomialPlot[1_078].Y = 3.818

	pointsOfPolynomialPlot[1_079].X = 10.79
	pointsOfPolynomialPlot[1_079].Y = 3.878

	pointsOfPolynomialPlot[1_080].X = 10.8
	pointsOfPolynomialPlot[1_080].Y = 3.939

	pointsOfPolynomialPlot[1_081].X = 10.81
	pointsOfPolynomialPlot[1_081].Y = 4.001

	pointsOfPolynomialPlot[1_082].X = 10.82
	pointsOfPolynomialPlot[1_082].Y = 4.063

	pointsOfPolynomialPlot[1_083].X = 10.83
	pointsOfPolynomialPlot[1_083].Y = 4.126

	pointsOfPolynomialPlot[1_084].X = 10.84
	pointsOfPolynomialPlot[1_084].Y = 4.19

	pointsOfPolynomialPlot[1_085].X = 10.85
	pointsOfPolynomialPlot[1_085].Y = 4.254

	pointsOfPolynomialPlot[1_086].X = 10.86
	pointsOfPolynomialPlot[1_086].Y = 4.32

	pointsOfPolynomialPlot[1_087].X = 10.87
	pointsOfPolynomialPlot[1_087].Y = 4.386

	pointsOfPolynomialPlot[1_088].X = 10.88
	pointsOfPolynomialPlot[1_088].Y = 4.452

	pointsOfPolynomialPlot[1_089].X = 10.89
	pointsOfPolynomialPlot[1_089].Y = 4.52

	pointsOfPolynomialPlot[1_090].X = 10.9
	pointsOfPolynomialPlot[1_090].Y = 4.588

	pointsOfPolynomialPlot[1_091].X = 10.91
	pointsOfPolynomialPlot[1_091].Y = 4.657

	pointsOfPolynomialPlot[1_092].X = 10.92
	pointsOfPolynomialPlot[1_092].Y = 4.727

	pointsOfPolynomialPlot[1_093].X = 10.93
	pointsOfPolynomialPlot[1_093].Y = 4.797

	pointsOfPolynomialPlot[1_094].X = 10.94
	pointsOfPolynomialPlot[1_094].Y = 4.868

	pointsOfPolynomialPlot[1_095].X = 10.95
	pointsOfPolynomialPlot[1_095].Y = 4.94

	pointsOfPolynomialPlot[1_096].X = 10.96
	pointsOfPolynomialPlot[1_096].Y = 5.013

	pointsOfPolynomialPlot[1_097].X = 10.97
	pointsOfPolynomialPlot[1_097].Y = 5.087

	pointsOfPolynomialPlot[1_098].X = 10.98
	pointsOfPolynomialPlot[1_098].Y = 5.161

	pointsOfPolynomialPlot[1_099].X = 10.99
	pointsOfPolynomialPlot[1_099].Y = 5.236

	pointsOfPolynomialPlot[1_100].X = 11.0
	pointsOfPolynomialPlot[1_100].Y = 5.312

	pointsOfPolynomialPlot[1_101].X = 11.01
	pointsOfPolynomialPlot[1_101].Y = 5.389

	pointsOfPolynomialPlot[1_102].X = 11.02
	pointsOfPolynomialPlot[1_102].Y = 5.466

	pointsOfPolynomialPlot[1_103].X = 11.03
	pointsOfPolynomialPlot[1_103].Y = 5.545

	pointsOfPolynomialPlot[1_104].X = 11.04
	pointsOfPolynomialPlot[1_104].Y = 5.624

	pointsOfPolynomialPlot[1_105].X = 11.05
	pointsOfPolynomialPlot[1_105].Y = 5.704

	pointsOfPolynomialPlot[1_106].X = 11.06
	pointsOfPolynomialPlot[1_106].Y = 5.784

	pointsOfPolynomialPlot[1_107].X = 11.07
	pointsOfPolynomialPlot[1_107].Y = 5.866

	pointsOfPolynomialPlot[1_108].X = 11.08
	pointsOfPolynomialPlot[1_108].Y = 5.948

	pointsOfPolynomialPlot[1_109].X = 11.09
	pointsOfPolynomialPlot[1_109].Y = 6.032

	pointsOfPolynomialPlot[1_110].X = 11.1
	pointsOfPolynomialPlot[1_110].Y = 6.116

	pointsOfPolynomialPlot[1_111].X = 11.11
	pointsOfPolynomialPlot[1_111].Y = 6.201

	pointsOfPolynomialPlot[1_112].X = 11.12
	pointsOfPolynomialPlot[1_112].Y = 6.286

	pointsOfPolynomialPlot[1_113].X = 11.13
	pointsOfPolynomialPlot[1_113].Y = 6.373

	pointsOfPolynomialPlot[1_114].X = 11.14
	pointsOfPolynomialPlot[1_114].Y = 6.46

	pointsOfPolynomialPlot[1_115].X = 11.15
	pointsOfPolynomialPlot[1_115].Y = 6.548

	pointsOfPolynomialPlot[1_116].X = 11.16
	pointsOfPolynomialPlot[1_116].Y = 6.638

	pointsOfPolynomialPlot[1_117].X = 11.17
	pointsOfPolynomialPlot[1_117].Y = 6.728

	pointsOfPolynomialPlot[1_118].X = 11.18
	pointsOfPolynomialPlot[1_118].Y = 6.818

	pointsOfPolynomialPlot[1_119].X = 11.19
	pointsOfPolynomialPlot[1_119].Y = 6.91

	pointsOfPolynomialPlot[1_120].X = 11.2
	pointsOfPolynomialPlot[1_120].Y = 7.003

	pointsOfPolynomialPlot[1_121].X = 11.21
	pointsOfPolynomialPlot[1_121].Y = 7.096

	pointsOfPolynomialPlot[1_122].X = 11.22
	pointsOfPolynomialPlot[1_122].Y = 7.191

	pointsOfPolynomialPlot[1_123].X = 11.23
	pointsOfPolynomialPlot[1_123].Y = 7.286

	pointsOfPolynomialPlot[1_124].X = 11.24
	pointsOfPolynomialPlot[1_124].Y = 7.382

	pointsOfPolynomialPlot[1_125].X = 11.25
	pointsOfPolynomialPlot[1_125].Y = 7.479

	pointsOfPolynomialPlot[1_126].X = 11.26
	pointsOfPolynomialPlot[1_126].Y = 7.577

	pointsOfPolynomialPlot[1_127].X = 11.27
	pointsOfPolynomialPlot[1_127].Y = 7.676

	pointsOfPolynomialPlot[1_128].X = 11.28
	pointsOfPolynomialPlot[1_128].Y = 7.776

	pointsOfPolynomialPlot[1_129].X = 11.29
	pointsOfPolynomialPlot[1_129].Y = 7.876

	pointsOfPolynomialPlot[1_130].X = 11.3
	pointsOfPolynomialPlot[1_130].Y = 7.978

	pointsOfPolynomialPlot[1_131].X = 11.31
	pointsOfPolynomialPlot[1_131].Y = 8.08

	pointsOfPolynomialPlot[1_132].X = 11.32
	pointsOfPolynomialPlot[1_132].Y = 8.184

	pointsOfPolynomialPlot[1_133].X = 11.33
	pointsOfPolynomialPlot[1_133].Y = 8.288

	pointsOfPolynomialPlot[1_134].X = 11.34
	pointsOfPolynomialPlot[1_134].Y = 8.394

	pointsOfPolynomialPlot[1_135].X = 11.35
	pointsOfPolynomialPlot[1_135].Y = 8.5

	pointsOfPolynomialPlot[1_136].X = 11.36
	pointsOfPolynomialPlot[1_136].Y = 8.607

	pointsOfPolynomialPlot[1_137].X = 11.37
	pointsOfPolynomialPlot[1_137].Y = 8.715

	pointsOfPolynomialPlot[1_138].X = 11.38
	pointsOfPolynomialPlot[1_138].Y = 8.824

	pointsOfPolynomialPlot[1_139].X = 11.39
	pointsOfPolynomialPlot[1_139].Y = 8.935

	pointsOfPolynomialPlot[1_140].X = 11.4
	pointsOfPolynomialPlot[1_140].Y = 9.046

	pointsOfPolynomialPlot[1_141].X = 11.41
	pointsOfPolynomialPlot[1_141].Y = 9.158

	pointsOfPolynomialPlot[1_142].X = 11.42
	pointsOfPolynomialPlot[1_142].Y = 9.271

	pointsOfPolynomialPlot[1_143].X = 11.43
	pointsOfPolynomialPlot[1_143].Y = 9.385

	pointsOfPolynomialPlot[1_144].X = 11.44
	pointsOfPolynomialPlot[1_144].Y = 9.5

	pointsOfPolynomialPlot[1_145].X = 11.45
	pointsOfPolynomialPlot[1_145].Y = 9.615

	pointsOfPolynomialPlot[1_146].X = 11.46
	pointsOfPolynomialPlot[1_146].Y = 9.732

	pointsOfPolynomialPlot[1_147].X = 11.47
	pointsOfPolynomialPlot[1_147].Y = 9.85

	pointsOfPolynomialPlot[1_148].X = 11.48
	pointsOfPolynomialPlot[1_148].Y = 9.969

	pointsOfPolynomialPlot[1_149].X = 11.49
	pointsOfPolynomialPlot[1_149].Y = 10.089

	pointsOfPolynomialPlot[1_150].X = 11.5
	pointsOfPolynomialPlot[1_150].Y = 10.21

	pointsOfPolynomialPlot[1_151].X = 11.51
	pointsOfPolynomialPlot[1_151].Y = 10.332

	pointsOfPolynomialPlot[1_152].X = 11.52
	pointsOfPolynomialPlot[1_152].Y = 10.455

	pointsOfPolynomialPlot[1_153].X = 11.53
	pointsOfPolynomialPlot[1_153].Y = 10.579

	pointsOfPolynomialPlot[1_154].X = 11.54
	pointsOfPolynomialPlot[1_154].Y = 10.704

	pointsOfPolynomialPlot[1_155].X = 11.55
	pointsOfPolynomialPlot[1_155].Y = 10.83

	pointsOfPolynomialPlot[1_156].X = 11.56
	pointsOfPolynomialPlot[1_156].Y = 10.958

	pointsOfPolynomialPlot[1_157].X = 11.57
	pointsOfPolynomialPlot[1_157].Y = 11.086

	pointsOfPolynomialPlot[1_158].X = 11.58
	pointsOfPolynomialPlot[1_158].Y = 11.215

	pointsOfPolynomialPlot[1_159].X = 11.59
	pointsOfPolynomialPlot[1_159].Y = 11.345

	pointsOfPolynomialPlot[1_160].X = 11.6
	pointsOfPolynomialPlot[1_160].Y = 11.477

	pointsOfPolynomialPlot[1_161].X = 11.61
	pointsOfPolynomialPlot[1_161].Y = 11.609

	pointsOfPolynomialPlot[1_162].X = 11.62
	pointsOfPolynomialPlot[1_162].Y = 11.742

	pointsOfPolynomialPlot[1_163].X = 11.63
	pointsOfPolynomialPlot[1_163].Y = 11.877

	pointsOfPolynomialPlot[1_164].X = 11.64
	pointsOfPolynomialPlot[1_164].Y = 12.013

	pointsOfPolynomialPlot[1_165].X = 11.65
	pointsOfPolynomialPlot[1_165].Y = 12.149

	pointsOfPolynomialPlot[1_166].X = 11.66
	pointsOfPolynomialPlot[1_166].Y = 12.287

	pointsOfPolynomialPlot[1_167].X = 11.67
	pointsOfPolynomialPlot[1_167].Y = 12.426

	pointsOfPolynomialPlot[1_168].X = 11.68
	pointsOfPolynomialPlot[1_168].Y = 12.566

	pointsOfPolynomialPlot[1_169].X = 11.69
	pointsOfPolynomialPlot[1_169].Y = 12.707

	pointsOfPolynomialPlot[1_170].X = 11.7
	pointsOfPolynomialPlot[1_170].Y = 12.849

	pointsOfPolynomialPlot[1_171].X = 11.71
	pointsOfPolynomialPlot[1_171].Y = 12.993

	pointsOfPolynomialPlot[1_172].X = 11.72
	pointsOfPolynomialPlot[1_172].Y = 13.137

	pointsOfPolynomialPlot[1_173].X = 11.73
	pointsOfPolynomialPlot[1_173].Y = 13.283

	pointsOfPolynomialPlot[1_174].X = 11.74
	pointsOfPolynomialPlot[1_174].Y = 13.43

	pointsOfPolynomialPlot[1_175].X = 11.75
	pointsOfPolynomialPlot[1_175].Y = 13.577

	pointsOfPolynomialPlot[1_176].X = 11.76
	pointsOfPolynomialPlot[1_176].Y = 13.726

	pointsOfPolynomialPlot[1_177].X = 11.77
	pointsOfPolynomialPlot[1_177].Y = 13.877

	pointsOfPolynomialPlot[1_178].X = 11.78
	pointsOfPolynomialPlot[1_178].Y = 14.028

	pointsOfPolynomialPlot[1_179].X = 11.79
	pointsOfPolynomialPlot[1_179].Y = 14.18

	pointsOfPolynomialPlot[1_180].X = 11.8
	pointsOfPolynomialPlot[1_180].Y = 14.334

	pointsOfPolynomialPlot[1_181].X = 11.81
	pointsOfPolynomialPlot[1_181].Y = 14.489

	pointsOfPolynomialPlot[1_182].X = 11.82
	pointsOfPolynomialPlot[1_182].Y = 14.645

	pointsOfPolynomialPlot[1_183].X = 11.83
	pointsOfPolynomialPlot[1_183].Y = 14.802

	pointsOfPolynomialPlot[1_184].X = 11.84
	pointsOfPolynomialPlot[1_184].Y = 14.96

	pointsOfPolynomialPlot[1_185].X = 11.85
	pointsOfPolynomialPlot[1_185].Y = 15.119

	pointsOfPolynomialPlot[1_186].X = 11.86
	pointsOfPolynomialPlot[1_186].Y = 15.28

	pointsOfPolynomialPlot[1_187].X = 11.87
	pointsOfPolynomialPlot[1_187].Y = 15.442

	pointsOfPolynomialPlot[1_188].X = 11.88
	pointsOfPolynomialPlot[1_188].Y = 15.605

	pointsOfPolynomialPlot[1_189].X = 11.89
	pointsOfPolynomialPlot[1_189].Y = 15.769

	pointsOfPolynomialPlot[1_190].X = 11.9
	pointsOfPolynomialPlot[1_190].Y = 15.935

	pointsOfPolynomialPlot[1_191].X = 11.91
	pointsOfPolynomialPlot[1_191].Y = 16.102

	pointsOfPolynomialPlot[1_192].X = 11.92
	pointsOfPolynomialPlot[1_192].Y = 16.269

	pointsOfPolynomialPlot[1_193].X = 11.93
	pointsOfPolynomialPlot[1_193].Y = 16.439

	pointsOfPolynomialPlot[1_194].X = 11.94
	pointsOfPolynomialPlot[1_194].Y = 16.609

	pointsOfPolynomialPlot[1_195].X = 11.95
	pointsOfPolynomialPlot[1_195].Y = 16.781

	pointsOfPolynomialPlot[1_196].X = 11.96
	pointsOfPolynomialPlot[1_196].Y = 16.954

	pointsOfPolynomialPlot[1_197].X = 11.97
	pointsOfPolynomialPlot[1_197].Y = 17.128

	pointsOfPolynomialPlot[1_198].X = 11.98
	pointsOfPolynomialPlot[1_198].Y = 17.303

	pointsOfPolynomialPlot[1_199].X = 11.99
	pointsOfPolynomialPlot[1_199].Y = 17.48

	pointsOfPolynomialPlot[1_200].X = 12.0
	pointsOfPolynomialPlot[1_200].Y = 17.658

	pointsOfPolynomialPlot[1_201].X = 12.01
	pointsOfPolynomialPlot[1_201].Y = 17.837

	pointsOfPolynomialPlot[1_202].X = 12.02
	pointsOfPolynomialPlot[1_202].Y = 18.017

	pointsOfPolynomialPlot[1_203].X = 12.03
	pointsOfPolynomialPlot[1_203].Y = 18.199

	pointsOfPolynomialPlot[1_204].X = 12.04
	pointsOfPolynomialPlot[1_204].Y = 18.382

	pointsOfPolynomialPlot[1_205].X = 12.05
	pointsOfPolynomialPlot[1_205].Y = 18.566

	pointsOfPolynomialPlot[1_206].X = 12.06
	pointsOfPolynomialPlot[1_206].Y = 18.752

	pointsOfPolynomialPlot[1_207].X = 12.07
	pointsOfPolynomialPlot[1_207].Y = 18.939

	pointsOfPolynomialPlot[1_208].X = 12.08
	pointsOfPolynomialPlot[1_208].Y = 19.127

	pointsOfPolynomialPlot[1_209].X = 12.09
	pointsOfPolynomialPlot[1_209].Y = 19.316

	pointsOfPolynomialPlot[1_210].X = 12.1
	pointsOfPolynomialPlot[1_210].Y = 19.507

	pointsOfPolynomialPlot[1_211].X = 12.11
	pointsOfPolynomialPlot[1_211].Y = 19.699

	pointsOfPolynomialPlot[1_212].X = 12.12
	pointsOfPolynomialPlot[1_212].Y = 19.893

	pointsOfPolynomialPlot[1_213].X = 12.13
	pointsOfPolynomialPlot[1_213].Y = 20.088

	pointsOfPolynomialPlot[1_214].X = 12.14
	pointsOfPolynomialPlot[1_214].Y = 20.284

	pointsOfPolynomialPlot[1_215].X = 12.15
	pointsOfPolynomialPlot[1_215].Y = 20.481

	pointsOfPolynomialPlot[1_216].X = 12.16
	pointsOfPolynomialPlot[1_216].Y = 20.68

	pointsOfPolynomialPlot[1_217].X = 12.17
	pointsOfPolynomialPlot[1_217].Y = 20.881

	pointsOfPolynomialPlot[1_218].X = 12.18
	pointsOfPolynomialPlot[1_218].Y = 21.082

	pointsOfPolynomialPlot[1_219].X = 12.19
	pointsOfPolynomialPlot[1_219].Y = 21.285

	pointsOfPolynomialPlot[1_220].X = 12.2
	pointsOfPolynomialPlot[1_220].Y = 21.489

	pointsOfPolynomialPlot[1_221].X = 12.21
	pointsOfPolynomialPlot[1_221].Y = 21.695

	pointsOfPolynomialPlot[1_222].X = 12.22
	pointsOfPolynomialPlot[1_222].Y = 21.902

	pointsOfPolynomialPlot[1_223].X = 12.23
	pointsOfPolynomialPlot[1_223].Y = 22.111

	pointsOfPolynomialPlot[1_224].X = 12.24
	pointsOfPolynomialPlot[1_224].Y = 22.321

	pointsOfPolynomialPlot[1_225].X = 12.25
	pointsOfPolynomialPlot[1_225].Y = 22.532

	pointsOfPolynomialPlot[1_226].X = 12.26
	pointsOfPolynomialPlot[1_226].Y = 22.745

	pointsOfPolynomialPlot[1_227].X = 12.27
	pointsOfPolynomialPlot[1_227].Y = 22.959

	pointsOfPolynomialPlot[1_228].X = 12.28
	pointsOfPolynomialPlot[1_228].Y = 23.174

	pointsOfPolynomialPlot[1_229].X = 12.29
	pointsOfPolynomialPlot[1_229].Y = 23.391

	pointsOfPolynomialPlot[1_230].X = 12.3
	pointsOfPolynomialPlot[1_230].Y = 23.61

	pointsOfPolynomialPlot[1_231].X = 12.31
	pointsOfPolynomialPlot[1_231].Y = 23.829

	pointsOfPolynomialPlot[1_232].X = 12.32
	pointsOfPolynomialPlot[1_232].Y = 24.051

	pointsOfPolynomialPlot[1_233].X = 12.33
	pointsOfPolynomialPlot[1_233].Y = 24.273

	pointsOfPolynomialPlot[1_234].X = 12.34
	pointsOfPolynomialPlot[1_234].Y = 24.498

	pointsOfPolynomialPlot[1_235].X = 12.35
	pointsOfPolynomialPlot[1_235].Y = 24.723

	pointsOfPolynomialPlot[1_236].X = 12.36
	pointsOfPolynomialPlot[1_236].Y = 24.95

	pointsOfPolynomialPlot[1_237].X = 12.37
	pointsOfPolynomialPlot[1_237].Y = 25.179

	pointsOfPolynomialPlot[1_238].X = 12.38
	pointsOfPolynomialPlot[1_238].Y = 25.409

	pointsOfPolynomialPlot[1_239].X = 12.39
	pointsOfPolynomialPlot[1_239].Y = 25.64

	pointsOfPolynomialPlot[1_240].X = 12.4
	pointsOfPolynomialPlot[1_240].Y = 25.873

	pointsOfPolynomialPlot[1_241].X = 12.41
	pointsOfPolynomialPlot[1_241].Y = 26.108

	pointsOfPolynomialPlot[1_242].X = 12.42
	pointsOfPolynomialPlot[1_242].Y = 26.344

	pointsOfPolynomialPlot[1_243].X = 12.43
	pointsOfPolynomialPlot[1_243].Y = 26.582

	pointsOfPolynomialPlot[1_244].X = 12.44
	pointsOfPolynomialPlot[1_244].Y = 26.821

	pointsOfPolynomialPlot[1_245].X = 12.45
	pointsOfPolynomialPlot[1_245].Y = 27.061

	pointsOfPolynomialPlot[1_246].X = 12.46
	pointsOfPolynomialPlot[1_246].Y = 27.303

	pointsOfPolynomialPlot[1_247].X = 12.47
	pointsOfPolynomialPlot[1_247].Y = 27.547

	pointsOfPolynomialPlot[1_248].X = 12.48
	pointsOfPolynomialPlot[1_248].Y = 27.792

	pointsOfPolynomialPlot[1_249].X = 12.49
	pointsOfPolynomialPlot[1_249].Y = 28.039

	pointsOfPolynomialPlot[1_250].X = 12.5
	pointsOfPolynomialPlot[1_250].Y = 28.287

	pointsOfPolynomialPlot[1_251].X = 12.51
	pointsOfPolynomialPlot[1_251].Y = 28.537

	pointsOfPolynomialPlot[1_252].X = 12.52
	pointsOfPolynomialPlot[1_252].Y = 28.788

	pointsOfPolynomialPlot[1_253].X = 12.53
	pointsOfPolynomialPlot[1_253].Y = 29.041

	pointsOfPolynomialPlot[1_254].X = 12.54
	pointsOfPolynomialPlot[1_254].Y = 29.295

	pointsOfPolynomialPlot[1_255].X = 12.55
	pointsOfPolynomialPlot[1_255].Y = 29.552

	pointsOfPolynomialPlot[1_256].X = 12.56
	pointsOfPolynomialPlot[1_256].Y = 29.809

	pointsOfPolynomialPlot[1_257].X = 12.57
	pointsOfPolynomialPlot[1_257].Y = 30.068

	pointsOfPolynomialPlot[1_258].X = 12.58
	pointsOfPolynomialPlot[1_258].Y = 30.329

	pointsOfPolynomialPlot[1_259].X = 12.59
	pointsOfPolynomialPlot[1_259].Y = 30.592

	pointsOfPolynomialPlot[1_260].X = 12.6
	pointsOfPolynomialPlot[1_260].Y = 30.856

	pointsOfPolynomialPlot[1_261].X = 12.61
	pointsOfPolynomialPlot[1_261].Y = 31.121

	pointsOfPolynomialPlot[1_262].X = 12.62
	pointsOfPolynomialPlot[1_262].Y = 31.389

	pointsOfPolynomialPlot[1_263].X = 12.63
	pointsOfPolynomialPlot[1_263].Y = 31.658

	pointsOfPolynomialPlot[1_264].X = 12.64
	pointsOfPolynomialPlot[1_264].Y = 31.928

	pointsOfPolynomialPlot[1_265].X = 12.65
	pointsOfPolynomialPlot[1_265].Y = 32.2

	pointsOfPolynomialPlot[1_266].X = 12.66
	pointsOfPolynomialPlot[1_266].Y = 32.474

	pointsOfPolynomialPlot[1_267].X = 12.67
	pointsOfPolynomialPlot[1_267].Y = 32.75

	pointsOfPolynomialPlot[1_268].X = 12.68
	pointsOfPolynomialPlot[1_268].Y = 33.027

	pointsOfPolynomialPlot[1_269].X = 12.69
	pointsOfPolynomialPlot[1_269].Y = 33.306

	pointsOfPolynomialPlot[1_270].X = 12.7
	pointsOfPolynomialPlot[1_270].Y = 33.586

	pointsOfPolynomialPlot[1_271].X = 12.71
	pointsOfPolynomialPlot[1_271].Y = 33.868

	pointsOfPolynomialPlot[1_272].X = 12.72
	pointsOfPolynomialPlot[1_272].Y = 34.152

	pointsOfPolynomialPlot[1_273].X = 12.73
	pointsOfPolynomialPlot[1_273].Y = 34.438

	pointsOfPolynomialPlot[1_274].X = 12.74
	pointsOfPolynomialPlot[1_274].Y = 34.725

	pointsOfPolynomialPlot[1_275].X = 12.75
	pointsOfPolynomialPlot[1_275].Y = 35.014

	pointsOfPolynomialPlot[1_276].X = 12.76
	pointsOfPolynomialPlot[1_276].Y = 35.305

	pointsOfPolynomialPlot[1_277].X = 12.77
	pointsOfPolynomialPlot[1_277].Y = 35.597

	pointsOfPolynomialPlot[1_278].X = 12.78
	pointsOfPolynomialPlot[1_278].Y = 35.891

	pointsOfPolynomialPlot[1_279].X = 12.79
	pointsOfPolynomialPlot[1_279].Y = 36.187

	pointsOfPolynomialPlot[1_280].X = 12.8
	pointsOfPolynomialPlot[1_280].Y = 36.484

	pointsOfPolynomialPlot[1_281].X = 12.81
	pointsOfPolynomialPlot[1_281].Y = 36.784

	pointsOfPolynomialPlot[1_282].X = 12.82
	pointsOfPolynomialPlot[1_282].Y = 37.085

	pointsOfPolynomialPlot[1_283].X = 12.83
	pointsOfPolynomialPlot[1_283].Y = 37.387

	pointsOfPolynomialPlot[1_284].X = 12.84
	pointsOfPolynomialPlot[1_284].Y = 37.692

	pointsOfPolynomialPlot[1_285].X = 12.85
	pointsOfPolynomialPlot[1_285].Y = 37.998

	pointsOfPolynomialPlot[1_286].X = 12.86
	pointsOfPolynomialPlot[1_286].Y = 38.306

	pointsOfPolynomialPlot[1_287].X = 12.87
	pointsOfPolynomialPlot[1_287].Y = 38.616

	pointsOfPolynomialPlot[1_288].X = 12.88
	pointsOfPolynomialPlot[1_288].Y = 38.928

	pointsOfPolynomialPlot[1_289].X = 12.89
	pointsOfPolynomialPlot[1_289].Y = 39.241

	pointsOfPolynomialPlot[1_290].X = 12.9
	pointsOfPolynomialPlot[1_290].Y = 39.557

	pointsOfPolynomialPlot[1_291].X = 12.91
	pointsOfPolynomialPlot[1_291].Y = 39.874

	pointsOfPolynomialPlot[1_292].X = 12.92
	pointsOfPolynomialPlot[1_292].Y = 40.192

	pointsOfPolynomialPlot[1_293].X = 12.93
	pointsOfPolynomialPlot[1_293].Y = 40.513

	pointsOfPolynomialPlot[1_294].X = 12.94
	pointsOfPolynomialPlot[1_294].Y = 40.836

	pointsOfPolynomialPlot[1_295].X = 12.95
	pointsOfPolynomialPlot[1_295].Y = 41.16

	pointsOfPolynomialPlot[1_296].X = 12.96
	pointsOfPolynomialPlot[1_296].Y = 41.486

	pointsOfPolynomialPlot[1_297].X = 12.97
	pointsOfPolynomialPlot[1_297].Y = 41.814

	pointsOfPolynomialPlot[1_298].X = 12.98
	pointsOfPolynomialPlot[1_298].Y = 42.144

	pointsOfPolynomialPlot[1_299].X = 12.99
	pointsOfPolynomialPlot[1_299].Y = 42.476

	pointsOfPolynomialPlot[1_300].X = 13.0
	pointsOfPolynomialPlot[1_300].Y = 42.809

	pointsOfPolynomialPlot[1_301].X = 13.01
	pointsOfPolynomialPlot[1_301].Y = 43.145

	pointsOfPolynomialPlot[1_302].X = 13.02
	pointsOfPolynomialPlot[1_302].Y = 43.482

	pointsOfPolynomialPlot[1_303].X = 13.03
	pointsOfPolynomialPlot[1_303].Y = 43.821

	pointsOfPolynomialPlot[1_304].X = 13.04
	pointsOfPolynomialPlot[1_304].Y = 44.163

	pointsOfPolynomialPlot[1_305].X = 13.05
	pointsOfPolynomialPlot[1_305].Y = 44.506

	pointsOfPolynomialPlot[1_306].X = 13.06
	pointsOfPolynomialPlot[1_306].Y = 44.85

	pointsOfPolynomialPlot[1_307].X = 13.07
	pointsOfPolynomialPlot[1_307].Y = 45.197

	pointsOfPolynomialPlot[1_308].X = 13.08
	pointsOfPolynomialPlot[1_308].Y = 45.546

	pointsOfPolynomialPlot[1_309].X = 13.09
	pointsOfPolynomialPlot[1_309].Y = 45.897

	pointsOfPolynomialPlot[1_310].X = 13.1
	pointsOfPolynomialPlot[1_310].Y = 46.249

	pointsOfPolynomialPlot[1_311].X = 13.11
	pointsOfPolynomialPlot[1_311].Y = 46.604

	pointsOfPolynomialPlot[1_312].X = 13.12
	pointsOfPolynomialPlot[1_312].Y = 46.96

	pointsOfPolynomialPlot[1_313].X = 13.13
	pointsOfPolynomialPlot[1_313].Y = 47.319

	pointsOfPolynomialPlot[1_314].X = 13.14
	pointsOfPolynomialPlot[1_314].Y = 47.679

	pointsOfPolynomialPlot[1_315].X = 13.15
	pointsOfPolynomialPlot[1_315].Y = 48.042

	pointsOfPolynomialPlot[1_316].X = 13.16
	pointsOfPolynomialPlot[1_316].Y = 48.406

	pointsOfPolynomialPlot[1_317].X = 13.17
	pointsOfPolynomialPlot[1_317].Y = 48.772

	pointsOfPolynomialPlot[1_318].X = 13.18
	pointsOfPolynomialPlot[1_318].Y = 49.141

	pointsOfPolynomialPlot[1_319].X = 13.19
	pointsOfPolynomialPlot[1_319].Y = 49.511

	pointsOfPolynomialPlot[1_320].X = 13.2
	pointsOfPolynomialPlot[1_320].Y = 49.883

	pointsOfPolynomialPlot[1_321].X = 13.21
	pointsOfPolynomialPlot[1_321].Y = 50.258

	pointsOfPolynomialPlot[1_322].X = 13.22
	pointsOfPolynomialPlot[1_322].Y = 50.634

	pointsOfPolynomialPlot[1_323].X = 13.23
	pointsOfPolynomialPlot[1_323].Y = 51.012

	pointsOfPolynomialPlot[1_324].X = 13.24
	pointsOfPolynomialPlot[1_324].Y = 51.393

	pointsOfPolynomialPlot[1_325].X = 13.25
	pointsOfPolynomialPlot[1_325].Y = 51.775

	pointsOfPolynomialPlot[1_326].X = 13.26
	pointsOfPolynomialPlot[1_326].Y = 52.16

	pointsOfPolynomialPlot[1_327].X = 13.27
	pointsOfPolynomialPlot[1_327].Y = 52.546

	pointsOfPolynomialPlot[1_328].X = 13.28
	pointsOfPolynomialPlot[1_328].Y = 52.935

	pointsOfPolynomialPlot[1_329].X = 13.29
	pointsOfPolynomialPlot[1_329].Y = 53.325

	pointsOfPolynomialPlot[1_330].X = 13.3
	pointsOfPolynomialPlot[1_330].Y = 53.718

	pointsOfPolynomialPlot[1_331].X = 13.31
	pointsOfPolynomialPlot[1_331].Y = 54.113

	pointsOfPolynomialPlot[1_332].X = 13.32
	pointsOfPolynomialPlot[1_332].Y = 54.51

	pointsOfPolynomialPlot[1_333].X = 13.33
	pointsOfPolynomialPlot[1_333].Y = 54.909

	pointsOfPolynomialPlot[1_334].X = 13.34
	pointsOfPolynomialPlot[1_334].Y = 55.31

	pointsOfPolynomialPlot[1_335].X = 13.35
	pointsOfPolynomialPlot[1_335].Y = 55.713

	pointsOfPolynomialPlot[1_336].X = 13.36
	pointsOfPolynomialPlot[1_336].Y = 56.118

	pointsOfPolynomialPlot[1_337].X = 13.37
	pointsOfPolynomialPlot[1_337].Y = 56.526

	pointsOfPolynomialPlot[1_338].X = 13.38
	pointsOfPolynomialPlot[1_338].Y = 56.935

	pointsOfPolynomialPlot[1_339].X = 13.39
	pointsOfPolynomialPlot[1_339].Y = 57.347

	pointsOfPolynomialPlot[1_340].X = 13.4
	pointsOfPolynomialPlot[1_340].Y = 57.761

	pointsOfPolynomialPlot[1_341].X = 13.41
	pointsOfPolynomialPlot[1_341].Y = 58.177

	pointsOfPolynomialPlot[1_342].X = 13.42
	pointsOfPolynomialPlot[1_342].Y = 58.595

	pointsOfPolynomialPlot[1_343].X = 13.43
	pointsOfPolynomialPlot[1_343].Y = 59.015

	pointsOfPolynomialPlot[1_344].X = 13.44
	pointsOfPolynomialPlot[1_344].Y = 59.437

	pointsOfPolynomialPlot[1_345].X = 13.45
	pointsOfPolynomialPlot[1_345].Y = 59.862

	pointsOfPolynomialPlot[1_346].X = 13.46
	pointsOfPolynomialPlot[1_346].Y = 60.289

	pointsOfPolynomialPlot[1_347].X = 13.47
	pointsOfPolynomialPlot[1_347].Y = 60.718

	pointsOfPolynomialPlot[1_348].X = 13.48
	pointsOfPolynomialPlot[1_348].Y = 61.149

	pointsOfPolynomialPlot[1_349].X = 13.49
	pointsOfPolynomialPlot[1_349].Y = 61.583

	pointsOfPolynomialPlot[1_350].X = 13.5
	pointsOfPolynomialPlot[1_350].Y = 62.018

	pointsOfPolynomialPlot[1_351].X = 13.51
	pointsOfPolynomialPlot[1_351].Y = 62.456

	pointsOfPolynomialPlot[1_352].X = 13.52
	pointsOfPolynomialPlot[1_352].Y = 62.896

	pointsOfPolynomialPlot[1_353].X = 13.53
	pointsOfPolynomialPlot[1_353].Y = 63.339

	pointsOfPolynomialPlot[1_354].X = 13.54
	pointsOfPolynomialPlot[1_354].Y = 63.783

	pointsOfPolynomialPlot[1_355].X = 13.55
	pointsOfPolynomialPlot[1_355].Y = 64.23

	pointsOfPolynomialPlot[1_356].X = 13.56
	pointsOfPolynomialPlot[1_356].Y = 64.679

	pointsOfPolynomialPlot[1_357].X = 13.57
	pointsOfPolynomialPlot[1_357].Y = 65.13

	pointsOfPolynomialPlot[1_358].X = 13.58
	pointsOfPolynomialPlot[1_358].Y = 65.584

	pointsOfPolynomialPlot[1_359].X = 13.59
	pointsOfPolynomialPlot[1_359].Y = 66.04

	pointsOfPolynomialPlot[1_360].X = 13.6
	pointsOfPolynomialPlot[1_360].Y = 66.498

	pointsOfPolynomialPlot[1_361].X = 13.61
	pointsOfPolynomialPlot[1_361].Y = 66.959

	pointsOfPolynomialPlot[1_362].X = 13.62
	pointsOfPolynomialPlot[1_362].Y = 67.422

	pointsOfPolynomialPlot[1_363].X = 13.63
	pointsOfPolynomialPlot[1_363].Y = 67.887

	pointsOfPolynomialPlot[1_364].X = 13.64
	pointsOfPolynomialPlot[1_364].Y = 68.354

	pointsOfPolynomialPlot[1_365].X = 13.65
	pointsOfPolynomialPlot[1_365].Y = 68.824

	pointsOfPolynomialPlot[1_366].X = 13.66
	pointsOfPolynomialPlot[1_366].Y = 69.296

	pointsOfPolynomialPlot[1_367].X = 13.67
	pointsOfPolynomialPlot[1_367].Y = 69.77

	pointsOfPolynomialPlot[1_368].X = 13.68
	pointsOfPolynomialPlot[1_368].Y = 70.247

	pointsOfPolynomialPlot[1_369].X = 13.69
	pointsOfPolynomialPlot[1_369].Y = 70.726

	pointsOfPolynomialPlot[1_370].X = 13.7
	pointsOfPolynomialPlot[1_370].Y = 71.208

	pointsOfPolynomialPlot[1_371].X = 13.71
	pointsOfPolynomialPlot[1_371].Y = 71.692

	pointsOfPolynomialPlot[1_372].X = 13.72
	pointsOfPolynomialPlot[1_372].Y = 72.178

	pointsOfPolynomialPlot[1_373].X = 13.73
	pointsOfPolynomialPlot[1_373].Y = 72.667

	pointsOfPolynomialPlot[1_374].X = 13.74
	pointsOfPolynomialPlot[1_374].Y = 73.158

	pointsOfPolynomialPlot[1_375].X = 13.75
	pointsOfPolynomialPlot[1_375].Y = 73.651

	pointsOfPolynomialPlot[1_376].X = 13.76
	pointsOfPolynomialPlot[1_376].Y = 74.147

	pointsOfPolynomialPlot[1_377].X = 13.77
	pointsOfPolynomialPlot[1_377].Y = 74.646

	pointsOfPolynomialPlot[1_378].X = 13.78
	pointsOfPolynomialPlot[1_378].Y = 75.146

	pointsOfPolynomialPlot[1_379].X = 13.79
	pointsOfPolynomialPlot[1_379].Y = 75.65

	pointsOfPolynomialPlot[1_380].X = 13.8
	pointsOfPolynomialPlot[1_380].Y = 76.155

	pointsOfPolynomialPlot[1_381].X = 13.81
	pointsOfPolynomialPlot[1_381].Y = 76.663

	pointsOfPolynomialPlot[1_382].X = 13.82
	pointsOfPolynomialPlot[1_382].Y = 77.174

	pointsOfPolynomialPlot[1_383].X = 13.83
	pointsOfPolynomialPlot[1_383].Y = 77.687

	pointsOfPolynomialPlot[1_384].X = 13.84
	pointsOfPolynomialPlot[1_384].Y = 78.202

	pointsOfPolynomialPlot[1_385].X = 13.85
	pointsOfPolynomialPlot[1_385].Y = 78.72

	pointsOfPolynomialPlot[1_386].X = 13.86
	pointsOfPolynomialPlot[1_386].Y = 79.241

	pointsOfPolynomialPlot[1_387].X = 13.87
	pointsOfPolynomialPlot[1_387].Y = 79.764

	pointsOfPolynomialPlot[1_388].X = 13.88
	pointsOfPolynomialPlot[1_388].Y = 80.289

	pointsOfPolynomialPlot[1_389].X = 13.89
	pointsOfPolynomialPlot[1_389].Y = 80.817

	pointsOfPolynomialPlot[1_390].X = 13.9
	pointsOfPolynomialPlot[1_390].Y = 81.348

	pointsOfPolynomialPlot[1_391].X = 13.91
	pointsOfPolynomialPlot[1_391].Y = 81.881

	pointsOfPolynomialPlot[1_392].X = 13.92
	pointsOfPolynomialPlot[1_392].Y = 82.416

	pointsOfPolynomialPlot[1_393].X = 13.93
	pointsOfPolynomialPlot[1_393].Y = 82.954

	pointsOfPolynomialPlot[1_394].X = 13.94
	pointsOfPolynomialPlot[1_394].Y = 83.495

	pointsOfPolynomialPlot[1_395].X = 13.95
	pointsOfPolynomialPlot[1_395].Y = 84.038

	pointsOfPolynomialPlot[1_396].X = 13.96
	pointsOfPolynomialPlot[1_396].Y = 84.584

	pointsOfPolynomialPlot[1_397].X = 13.97
	pointsOfPolynomialPlot[1_397].Y = 85.132

	pointsOfPolynomialPlot[1_398].X = 13.98
	pointsOfPolynomialPlot[1_398].Y = 85.683

	pointsOfPolynomialPlot[1_399].X = 13.99
	pointsOfPolynomialPlot[1_399].Y = 86.237

	pointsOfPolynomialPlot[1_400].X = 14.0
	pointsOfPolynomialPlot[1_400].Y = 86.793

	pointsOfPolynomialPlot[1_401].X = 14.01
	pointsOfPolynomialPlot[1_401].Y = 87.352

	pointsOfPolynomialPlot[1_402].X = 14.02
	pointsOfPolynomialPlot[1_402].Y = 87.913

	pointsOfPolynomialPlot[1_403].X = 14.03
	pointsOfPolynomialPlot[1_403].Y = 88.477

	pointsOfPolynomialPlot[1_404].X = 14.04
	pointsOfPolynomialPlot[1_404].Y = 89.044

	pointsOfPolynomialPlot[1_405].X = 14.05
	pointsOfPolynomialPlot[1_405].Y = 89.613

	pointsOfPolynomialPlot[1_406].X = 14.06
	pointsOfPolynomialPlot[1_406].Y = 90.185

	pointsOfPolynomialPlot[1_407].X = 14.07
	pointsOfPolynomialPlot[1_407].Y = 90.76

	pointsOfPolynomialPlot[1_408].X = 14.08
	pointsOfPolynomialPlot[1_408].Y = 91.337

	pointsOfPolynomialPlot[1_409].X = 14.09
	pointsOfPolynomialPlot[1_409].Y = 91.917

	pointsOfPolynomialPlot[1_410].X = 14.1
	pointsOfPolynomialPlot[1_410].Y = 92.5

	pointsOfPolynomialPlot[1_411].X = 14.11
	pointsOfPolynomialPlot[1_411].Y = 93.085

	pointsOfPolynomialPlot[1_412].X = 14.12
	pointsOfPolynomialPlot[1_412].Y = 93.673

	pointsOfPolynomialPlot[1_413].X = 14.13
	pointsOfPolynomialPlot[1_413].Y = 94.264

	pointsOfPolynomialPlot[1_414].X = 14.14
	pointsOfPolynomialPlot[1_414].Y = 94.857

	pointsOfPolynomialPlot[1_415].X = 14.15
	pointsOfPolynomialPlot[1_415].Y = 95.453

	pointsOfPolynomialPlot[1_416].X = 14.16
	pointsOfPolynomialPlot[1_416].Y = 96.052

	pointsOfPolynomialPlot[1_417].X = 14.17
	pointsOfPolynomialPlot[1_417].Y = 96.654

	pointsOfPolynomialPlot[1_418].X = 14.18
	pointsOfPolynomialPlot[1_418].Y = 97.259

	pointsOfPolynomialPlot[1_419].X = 14.19
	pointsOfPolynomialPlot[1_419].Y = 97.866

	pointsOfPolynomialPlot[1_420].X = 14.2
	pointsOfPolynomialPlot[1_420].Y = 98.476

	pointsOfPolynomialPlot[1_421].X = 14.21
	pointsOfPolynomialPlot[1_421].Y = 99.088

	pointsOfPolynomialPlot[1_422].X = 14.22
	pointsOfPolynomialPlot[1_422].Y = 99.704

	pointsOfPolynomialPlot[1_423].X = 14.23
	pointsOfPolynomialPlot[1_423].Y = 100.322

	pointsOfPolynomialPlot[1_424].X = 14.24
	pointsOfPolynomialPlot[1_424].Y = 100.943

	pointsOfPolynomialPlot[1_425].X = 14.25
	pointsOfPolynomialPlot[1_425].Y = 101.567

	pointsOfPolynomialPlot[1_426].X = 14.26
	pointsOfPolynomialPlot[1_426].Y = 102.194

	pointsOfPolynomialPlot[1_427].X = 14.27
	pointsOfPolynomialPlot[1_427].Y = 102.824

	pointsOfPolynomialPlot[1_428].X = 14.28
	pointsOfPolynomialPlot[1_428].Y = 103.456

	pointsOfPolynomialPlot[1_429].X = 14.29
	pointsOfPolynomialPlot[1_429].Y = 104.091

	pointsOfPolynomialPlot[1_430].X = 14.3
	pointsOfPolynomialPlot[1_430].Y = 104.729

	pointsOfPolynomialPlot[1_431].X = 14.31
	pointsOfPolynomialPlot[1_431].Y = 105.37

	pointsOfPolynomialPlot[1_432].X = 14.32
	pointsOfPolynomialPlot[1_432].Y = 106.014

	pointsOfPolynomialPlot[1_433].X = 14.33
	pointsOfPolynomialPlot[1_433].Y = 106.661

	pointsOfPolynomialPlot[1_434].X = 14.34
	pointsOfPolynomialPlot[1_434].Y = 107.31

	pointsOfPolynomialPlot[1_435].X = 14.35
	pointsOfPolynomialPlot[1_435].Y = 107.963

	pointsOfPolynomialPlot[1_436].X = 14.36
	pointsOfPolynomialPlot[1_436].Y = 108.618

	pointsOfPolynomialPlot[1_437].X = 14.37
	pointsOfPolynomialPlot[1_437].Y = 109.277

	pointsOfPolynomialPlot[1_438].X = 14.38
	pointsOfPolynomialPlot[1_438].Y = 109.938

	pointsOfPolynomialPlot[1_439].X = 14.39
	pointsOfPolynomialPlot[1_439].Y = 110.602

	pointsOfPolynomialPlot[1_440].X = 14.4
	pointsOfPolynomialPlot[1_440].Y = 111.269

	pointsOfPolynomialPlot[1_441].X = 14.41
	pointsOfPolynomialPlot[1_441].Y = 111.939

	pointsOfPolynomialPlot[1_442].X = 14.42
	pointsOfPolynomialPlot[1_442].Y = 112.612

	pointsOfPolynomialPlot[1_443].X = 14.43
	pointsOfPolynomialPlot[1_443].Y = 113.288

	pointsOfPolynomialPlot[1_444].X = 14.44
	pointsOfPolynomialPlot[1_444].Y = 113.967

	pointsOfPolynomialPlot[1_445].X = 14.45
	pointsOfPolynomialPlot[1_445].Y = 114.649

	pointsOfPolynomialPlot[1_446].X = 14.46
	pointsOfPolynomialPlot[1_446].Y = 115.334

	pointsOfPolynomialPlot[1_447].X = 14.47
	pointsOfPolynomialPlot[1_447].Y = 116.022

	pointsOfPolynomialPlot[1_448].X = 14.48
	pointsOfPolynomialPlot[1_448].Y = 116.713

	pointsOfPolynomialPlot[1_449].X = 14.49
	pointsOfPolynomialPlot[1_449].Y = 117.407

	pointsOfPolynomialPlot[1_450].X = 14.5
	pointsOfPolynomialPlot[1_450].Y = 118.104

	pointsOfPolynomialPlot[1_451].X = 14.51
	pointsOfPolynomialPlot[1_451].Y = 118.804

	pointsOfPolynomialPlot[1_452].X = 14.52
	pointsOfPolynomialPlot[1_452].Y = 119.507

	pointsOfPolynomialPlot[1_453].X = 14.53
	pointsOfPolynomialPlot[1_453].Y = 120.213

	pointsOfPolynomialPlot[1_454].X = 14.54
	pointsOfPolynomialPlot[1_454].Y = 120.922

	pointsOfPolynomialPlot[1_455].X = 14.55
	pointsOfPolynomialPlot[1_455].Y = 121.635

	pointsOfPolynomialPlot[1_456].X = 14.56
	pointsOfPolynomialPlot[1_456].Y = 122.35

	pointsOfPolynomialPlot[1_457].X = 14.57
	pointsOfPolynomialPlot[1_457].Y = 123.068

	pointsOfPolynomialPlot[1_458].X = 14.58
	pointsOfPolynomialPlot[1_458].Y = 123.79

	pointsOfPolynomialPlot[1_459].X = 14.59
	pointsOfPolynomialPlot[1_459].Y = 124.515

	pointsOfPolynomialPlot[1_460].X = 14.6
	pointsOfPolynomialPlot[1_460].Y = 125.242

	pointsOfPolynomialPlot[1_461].X = 14.61
	pointsOfPolynomialPlot[1_461].Y = 125.973

	pointsOfPolynomialPlot[1_462].X = 14.62
	pointsOfPolynomialPlot[1_462].Y = 126.707

	pointsOfPolynomialPlot[1_463].X = 14.63
	pointsOfPolynomialPlot[1_463].Y = 127.444

	pointsOfPolynomialPlot[1_464].X = 14.64
	pointsOfPolynomialPlot[1_464].Y = 128.185

	pointsOfPolynomialPlot[1_465].X = 14.65
	pointsOfPolynomialPlot[1_465].Y = 128.928

	pointsOfPolynomialPlot[1_466].X = 14.66
	pointsOfPolynomialPlot[1_466].Y = 129.675

	pointsOfPolynomialPlot[1_467].X = 14.67
	pointsOfPolynomialPlot[1_467].Y = 130.425

	pointsOfPolynomialPlot[1_468].X = 14.68
	pointsOfPolynomialPlot[1_468].Y = 131.178

	pointsOfPolynomialPlot[1_469].X = 14.69
	pointsOfPolynomialPlot[1_469].Y = 131.934

	pointsOfPolynomialPlot[1_470].X = 14.7
	pointsOfPolynomialPlot[1_470].Y = 132.694

	pointsOfPolynomialPlot[1_471].X = 14.71
	pointsOfPolynomialPlot[1_471].Y = 133.456

	pointsOfPolynomialPlot[1_472].X = 14.72
	pointsOfPolynomialPlot[1_472].Y = 134.222

	pointsOfPolynomialPlot[1_473].X = 14.73
	pointsOfPolynomialPlot[1_473].Y = 134.991

	pointsOfPolynomialPlot[1_474].X = 14.74
	pointsOfPolynomialPlot[1_474].Y = 135.764

	pointsOfPolynomialPlot[1_475].X = 14.75
	pointsOfPolynomialPlot[1_475].Y = 136.539

	pointsOfPolynomialPlot[1_476].X = 14.76
	pointsOfPolynomialPlot[1_476].Y = 137.318

	pointsOfPolynomialPlot[1_477].X = 14.77
	pointsOfPolynomialPlot[1_477].Y = 138.1

	pointsOfPolynomialPlot[1_478].X = 14.78
	pointsOfPolynomialPlot[1_478].Y = 138.886

	pointsOfPolynomialPlot[1_479].X = 14.79
	pointsOfPolynomialPlot[1_479].Y = 139.674

	pointsOfPolynomialPlot[1_480].X = 14.8
	pointsOfPolynomialPlot[1_480].Y = 140.466

	pointsOfPolynomialPlot[1_481].X = 14.81
	pointsOfPolynomialPlot[1_481].Y = 141.262

	pointsOfPolynomialPlot[1_482].X = 14.82
	pointsOfPolynomialPlot[1_482].Y = 142.06

	pointsOfPolynomialPlot[1_483].X = 14.83
	pointsOfPolynomialPlot[1_483].Y = 142.862

	pointsOfPolynomialPlot[1_484].X = 14.84
	pointsOfPolynomialPlot[1_484].Y = 143.668

	pointsOfPolynomialPlot[1_485].X = 14.85
	pointsOfPolynomialPlot[1_485].Y = 144.477

	pointsOfPolynomialPlot[1_486].X = 14.86
	pointsOfPolynomialPlot[1_486].Y = 145.289

	pointsOfPolynomialPlot[1_487].X = 14.87
	pointsOfPolynomialPlot[1_487].Y = 146.104

	pointsOfPolynomialPlot[1_488].X = 14.88
	pointsOfPolynomialPlot[1_488].Y = 146.923

	pointsOfPolynomialPlot[1_489].X = 14.89
	pointsOfPolynomialPlot[1_489].Y = 147.745

	pointsOfPolynomialPlot[1_490].X = 14.9
	pointsOfPolynomialPlot[1_490].Y = 148.571

	pointsOfPolynomialPlot[1_491].X = 14.91
	pointsOfPolynomialPlot[1_491].Y = 149.4

	pointsOfPolynomialPlot[1_492].X = 14.92
	pointsOfPolynomialPlot[1_492].Y = 150.232

	pointsOfPolynomialPlot[1_493].X = 14.93
	pointsOfPolynomialPlot[1_493].Y = 151.068

	pointsOfPolynomialPlot[1_494].X = 14.94
	pointsOfPolynomialPlot[1_494].Y = 151.907

	pointsOfPolynomialPlot[1_495].X = 14.95
	pointsOfPolynomialPlot[1_495].Y = 152.75

	pointsOfPolynomialPlot[1_496].X = 14.96
	pointsOfPolynomialPlot[1_496].Y = 153.596

	pointsOfPolynomialPlot[1_497].X = 14.97
	pointsOfPolynomialPlot[1_497].Y = 154.445

	pointsOfPolynomialPlot[1_498].X = 14.98
	pointsOfPolynomialPlot[1_498].Y = 155.299

	pointsOfPolynomialPlot[1_499].X = 14.99
	pointsOfPolynomialPlot[1_499].Y = 156.155

	pointsOfPolynomialPlot[1_500].X = 15.0
	pointsOfPolynomialPlot[1_500].Y = 157.015

	pointsOfPolynomialPlot[1_501].X = 15.01
	pointsOfPolynomialPlot[1_501].Y = 157.879

	pointsOfPolynomialPlot[1_502].X = 15.02
	pointsOfPolynomialPlot[1_502].Y = 158.746

	pointsOfPolynomialPlot[1_503].X = 15.03
	pointsOfPolynomialPlot[1_503].Y = 159.616

	pointsOfPolynomialPlot[1_504].X = 15.04
	pointsOfPolynomialPlot[1_504].Y = 160.491

	pointsOfPolynomialPlot[1_505].X = 15.05
	pointsOfPolynomialPlot[1_505].Y = 161.368

	pointsOfPolynomialPlot[1_506].X = 15.06
	pointsOfPolynomialPlot[1_506].Y = 162.249

	pointsOfPolynomialPlot[1_507].X = 15.07
	pointsOfPolynomialPlot[1_507].Y = 163.134

	pointsOfPolynomialPlot[1_508].X = 15.08
	pointsOfPolynomialPlot[1_508].Y = 164.023

	pointsOfPolynomialPlot[1_509].X = 15.09
	pointsOfPolynomialPlot[1_509].Y = 164.914

	pointsOfPolynomialPlot[1_510].X = 15.1
	pointsOfPolynomialPlot[1_510].Y = 165.81

	pointsOfPolynomialPlot[1_511].X = 15.11
	pointsOfPolynomialPlot[1_511].Y = 166.709

	pointsOfPolynomialPlot[1_512].X = 15.12
	pointsOfPolynomialPlot[1_512].Y = 167.612

	pointsOfPolynomialPlot[1_513].X = 15.13
	pointsOfPolynomialPlot[1_513].Y = 168.518

	pointsOfPolynomialPlot[1_514].X = 15.14
	pointsOfPolynomialPlot[1_514].Y = 169.428

	pointsOfPolynomialPlot[1_515].X = 15.15
	pointsOfPolynomialPlot[1_515].Y = 170.342

	pointsOfPolynomialPlot[1_516].X = 15.16
	pointsOfPolynomialPlot[1_516].Y = 171.259

	pointsOfPolynomialPlot[1_517].X = 15.17
	pointsOfPolynomialPlot[1_517].Y = 172.18

	pointsOfPolynomialPlot[1_518].X = 15.18
	pointsOfPolynomialPlot[1_518].Y = 173.105

	pointsOfPolynomialPlot[1_519].X = 15.19
	pointsOfPolynomialPlot[1_519].Y = 174.033

	pointsOfPolynomialPlot[1_520].X = 15.2
	pointsOfPolynomialPlot[1_520].Y = 174.965

	pointsOfPolynomialPlot[1_521].X = 15.21
	pointsOfPolynomialPlot[1_521].Y = 175.901

	pointsOfPolynomialPlot[1_522].X = 15.22
	pointsOfPolynomialPlot[1_522].Y = 176.84

	pointsOfPolynomialPlot[1_523].X = 15.23
	pointsOfPolynomialPlot[1_523].Y = 177.783

	pointsOfPolynomialPlot[1_524].X = 15.24
	pointsOfPolynomialPlot[1_524].Y = 178.73

	pointsOfPolynomialPlot[1_525].X = 15.25
	pointsOfPolynomialPlot[1_525].Y = 179.68

	pointsOfPolynomialPlot[1_526].X = 15.26
	pointsOfPolynomialPlot[1_526].Y = 180.635

	pointsOfPolynomialPlot[1_527].X = 15.27
	pointsOfPolynomialPlot[1_527].Y = 181.593

	pointsOfPolynomialPlot[1_528].X = 15.28
	pointsOfPolynomialPlot[1_528].Y = 182.555

	pointsOfPolynomialPlot[1_529].X = 15.29
	pointsOfPolynomialPlot[1_529].Y = 183.52

	pointsOfPolynomialPlot[1_530].X = 15.3
	pointsOfPolynomialPlot[1_530].Y = 184.49

	pointsOfPolynomialPlot[1_531].X = 15.31
	pointsOfPolynomialPlot[1_531].Y = 185.463

	pointsOfPolynomialPlot[1_532].X = 15.32
	pointsOfPolynomialPlot[1_532].Y = 186.44

	pointsOfPolynomialPlot[1_533].X = 15.33
	pointsOfPolynomialPlot[1_533].Y = 187.421

	pointsOfPolynomialPlot[1_534].X = 15.34
	pointsOfPolynomialPlot[1_534].Y = 188.405

	pointsOfPolynomialPlot[1_535].X = 15.35
	pointsOfPolynomialPlot[1_535].Y = 189.394

	pointsOfPolynomialPlot[1_536].X = 15.36
	pointsOfPolynomialPlot[1_536].Y = 190.386

	pointsOfPolynomialPlot[1_537].X = 15.37
	pointsOfPolynomialPlot[1_537].Y = 191.382

	pointsOfPolynomialPlot[1_538].X = 15.38
	pointsOfPolynomialPlot[1_538].Y = 192.383

	pointsOfPolynomialPlot[1_539].X = 15.39
	pointsOfPolynomialPlot[1_539].Y = 193.387

	pointsOfPolynomialPlot[1_540].X = 15.4
	pointsOfPolynomialPlot[1_540].Y = 194.394

	pointsOfPolynomialPlot[1_541].X = 15.41
	pointsOfPolynomialPlot[1_541].Y = 195.406

	pointsOfPolynomialPlot[1_542].X = 15.42
	pointsOfPolynomialPlot[1_542].Y = 196.422

	pointsOfPolynomialPlot[1_543].X = 15.43
	pointsOfPolynomialPlot[1_543].Y = 197.442

	pointsOfPolynomialPlot[1_544].X = 15.44
	pointsOfPolynomialPlot[1_544].Y = 198.465

	pointsOfPolynomialPlot[1_545].X = 15.45
	pointsOfPolynomialPlot[1_545].Y = 199.493

	pointsOfPolynomialPlot[1_546].X = 15.46
	pointsOfPolynomialPlot[1_546].Y = 200.524

	pointsOfPolynomialPlot[1_547].X = 15.47
	pointsOfPolynomialPlot[1_547].Y = 201.559

	pointsOfPolynomialPlot[1_548].X = 15.48
	pointsOfPolynomialPlot[1_548].Y = 202.599

	pointsOfPolynomialPlot[1_549].X = 15.49
	pointsOfPolynomialPlot[1_549].Y = 203.642

	pointsOfPolynomialPlot[1_550].X = 15.5
	pointsOfPolynomialPlot[1_550].Y = 204.69

	pointsOfPolynomialPlot[1_551].X = 15.51
	pointsOfPolynomialPlot[1_551].Y = 205.741

	pointsOfPolynomialPlot[1_552].X = 15.52
	pointsOfPolynomialPlot[1_552].Y = 206.796

	pointsOfPolynomialPlot[1_553].X = 15.53
	pointsOfPolynomialPlot[1_553].Y = 207.856

	pointsOfPolynomialPlot[1_554].X = 15.54
	pointsOfPolynomialPlot[1_554].Y = 208.919

	pointsOfPolynomialPlot[1_555].X = 15.55
	pointsOfPolynomialPlot[1_555].Y = 209.987

	pointsOfPolynomialPlot[1_556].X = 15.56
	pointsOfPolynomialPlot[1_556].Y = 211.058

	pointsOfPolynomialPlot[1_557].X = 15.57
	pointsOfPolynomialPlot[1_557].Y = 212.134

	pointsOfPolynomialPlot[1_558].X = 15.58
	pointsOfPolynomialPlot[1_558].Y = 213.214

	pointsOfPolynomialPlot[1_559].X = 15.59
	pointsOfPolynomialPlot[1_559].Y = 214.298

	pointsOfPolynomialPlot[1_560].X = 15.6
	pointsOfPolynomialPlot[1_560].Y = 215.386

	pointsOfPolynomialPlot[1_561].X = 15.61
	pointsOfPolynomialPlot[1_561].Y = 216.478

	pointsOfPolynomialPlot[1_562].X = 15.62
	pointsOfPolynomialPlot[1_562].Y = 217.574

	pointsOfPolynomialPlot[1_563].X = 15.63
	pointsOfPolynomialPlot[1_563].Y = 218.674

	pointsOfPolynomialPlot[1_564].X = 15.64
	pointsOfPolynomialPlot[1_564].Y = 219.779

	pointsOfPolynomialPlot[1_565].X = 15.65
	pointsOfPolynomialPlot[1_565].Y = 220.887

	pointsOfPolynomialPlot[1_566].X = 15.66
	pointsOfPolynomialPlot[1_566].Y = 222.0

	pointsOfPolynomialPlot[1_567].X = 15.67
	pointsOfPolynomialPlot[1_567].Y = 223.117

	pointsOfPolynomialPlot[1_568].X = 15.68
	pointsOfPolynomialPlot[1_568].Y = 224.238

	pointsOfPolynomialPlot[1_569].X = 15.69
	pointsOfPolynomialPlot[1_569].Y = 225.363

	pointsOfPolynomialPlot[1_570].X = 15.7
	pointsOfPolynomialPlot[1_570].Y = 226.493

	pointsOfPolynomialPlot[1_571].X = 15.71
	pointsOfPolynomialPlot[1_571].Y = 227.627

	pointsOfPolynomialPlot[1_572].X = 15.72
	pointsOfPolynomialPlot[1_572].Y = 228.765

	pointsOfPolynomialPlot[1_573].X = 15.73
	pointsOfPolynomialPlot[1_573].Y = 229.907

	pointsOfPolynomialPlot[1_574].X = 15.74
	pointsOfPolynomialPlot[1_574].Y = 231.054

	pointsOfPolynomialPlot[1_575].X = 15.75
	pointsOfPolynomialPlot[1_575].Y = 232.204

	pointsOfPolynomialPlot[1_576].X = 15.76
	pointsOfPolynomialPlot[1_576].Y = 233.359

	pointsOfPolynomialPlot[1_577].X = 15.77
	pointsOfPolynomialPlot[1_577].Y = 234.519

	pointsOfPolynomialPlot[1_578].X = 15.78
	pointsOfPolynomialPlot[1_578].Y = 235.682

	pointsOfPolynomialPlot[1_579].X = 15.79
	pointsOfPolynomialPlot[1_579].Y = 236.85

	pointsOfPolynomialPlot[1_580].X = 15.8
	pointsOfPolynomialPlot[1_580].Y = 238.023

	pointsOfPolynomialPlot[1_581].X = 15.81
	pointsOfPolynomialPlot[1_581].Y = 239.199

	pointsOfPolynomialPlot[1_582].X = 15.82
	pointsOfPolynomialPlot[1_582].Y = 240.38

	pointsOfPolynomialPlot[1_583].X = 15.83
	pointsOfPolynomialPlot[1_583].Y = 241.566

	pointsOfPolynomialPlot[1_584].X = 15.84
	pointsOfPolynomialPlot[1_584].Y = 242.755

	pointsOfPolynomialPlot[1_585].X = 15.85
	pointsOfPolynomialPlot[1_585].Y = 243.949

	pointsOfPolynomialPlot[1_586].X = 15.86
	pointsOfPolynomialPlot[1_586].Y = 245.148

	pointsOfPolynomialPlot[1_587].X = 15.87
	pointsOfPolynomialPlot[1_587].Y = 246.351

	pointsOfPolynomialPlot[1_588].X = 15.88
	pointsOfPolynomialPlot[1_588].Y = 247.558

	pointsOfPolynomialPlot[1_589].X = 15.89
	pointsOfPolynomialPlot[1_589].Y = 248.769

	pointsOfPolynomialPlot[1_590].X = 15.9
	pointsOfPolynomialPlot[1_590].Y = 249.986

	pointsOfPolynomialPlot[1_591].X = 15.91
	pointsOfPolynomialPlot[1_591].Y = 251.206

	pointsOfPolynomialPlot[1_592].X = 15.92
	pointsOfPolynomialPlot[1_592].Y = 252.431

	pointsOfPolynomialPlot[1_593].X = 15.93
	pointsOfPolynomialPlot[1_593].Y = 253.661

	pointsOfPolynomialPlot[1_594].X = 15.94
	pointsOfPolynomialPlot[1_594].Y = 254.894

	pointsOfPolynomialPlot[1_595].X = 15.95
	pointsOfPolynomialPlot[1_595].Y = 256.133

	pointsOfPolynomialPlot[1_596].X = 15.96
	pointsOfPolynomialPlot[1_596].Y = 257.376

	pointsOfPolynomialPlot[1_597].X = 15.97
	pointsOfPolynomialPlot[1_597].Y = 258.623

	pointsOfPolynomialPlot[1_598].X = 15.98
	pointsOfPolynomialPlot[1_598].Y = 259.875

	pointsOfPolynomialPlot[1_599].X = 15.99
	pointsOfPolynomialPlot[1_599].Y = 261.132

	pointsOfPolynomialPlot[1_600].X = 16.0
	pointsOfPolynomialPlot[1_600].Y = 262.393

	pointsOfPolynomialPlot[1_601].X = 16.01
	pointsOfPolynomialPlot[1_601].Y = 263.658

	pointsOfPolynomialPlot[1_602].X = 16.02
	pointsOfPolynomialPlot[1_602].Y = 264.929

	pointsOfPolynomialPlot[1_603].X = 16.03
	pointsOfPolynomialPlot[1_603].Y = 266.203

	pointsOfPolynomialPlot[1_604].X = 16.04
	pointsOfPolynomialPlot[1_604].Y = 267.483

	pointsOfPolynomialPlot[1_605].X = 16.05
	pointsOfPolynomialPlot[1_605].Y = 268.767

	pointsOfPolynomialPlot[1_606].X = 16.06
	pointsOfPolynomialPlot[1_606].Y = 270.055

	pointsOfPolynomialPlot[1_607].X = 16.07
	pointsOfPolynomialPlot[1_607].Y = 271.348

	pointsOfPolynomialPlot[1_608].X = 16.08
	pointsOfPolynomialPlot[1_608].Y = 272.646

	pointsOfPolynomialPlot[1_609].X = 16.09
	pointsOfPolynomialPlot[1_609].Y = 273.949

	pointsOfPolynomialPlot[1_610].X = 16.1
	pointsOfPolynomialPlot[1_610].Y = 275.256

	pointsOfPolynomialPlot[1_611].X = 16.11
	pointsOfPolynomialPlot[1_611].Y = 276.567

	pointsOfPolynomialPlot[1_612].X = 16.12
	pointsOfPolynomialPlot[1_612].Y = 277.884

	pointsOfPolynomialPlot[1_613].X = 16.13
	pointsOfPolynomialPlot[1_613].Y = 279.205

	pointsOfPolynomialPlot[1_614].X = 16.14
	pointsOfPolynomialPlot[1_614].Y = 280.531

	pointsOfPolynomialPlot[1_615].X = 16.15
	pointsOfPolynomialPlot[1_615].Y = 281.862

	pointsOfPolynomialPlot[1_616].X = 16.16
	pointsOfPolynomialPlot[1_616].Y = 283.197

	pointsOfPolynomialPlot[1_617].X = 16.17
	pointsOfPolynomialPlot[1_617].Y = 284.537

	pointsOfPolynomialPlot[1_618].X = 16.18
	pointsOfPolynomialPlot[1_618].Y = 285.882

	pointsOfPolynomialPlot[1_619].X = 16.19
	pointsOfPolynomialPlot[1_619].Y = 287.231

	pointsOfPolynomialPlot[1_620].X = 16.2
	pointsOfPolynomialPlot[1_620].Y = 288.586

	pointsOfPolynomialPlot[1_621].X = 16.21
	pointsOfPolynomialPlot[1_621].Y = 289.945

	pointsOfPolynomialPlot[1_622].X = 16.22
	pointsOfPolynomialPlot[1_622].Y = 291.309

	pointsOfPolynomialPlot[1_623].X = 16.23
	pointsOfPolynomialPlot[1_623].Y = 292.678

	pointsOfPolynomialPlot[1_624].X = 16.24
	pointsOfPolynomialPlot[1_624].Y = 294.051

	pointsOfPolynomialPlot[1_625].X = 16.25
	pointsOfPolynomialPlot[1_625].Y = 295.43

	pointsOfPolynomialPlot[1_626].X = 16.26
	pointsOfPolynomialPlot[1_626].Y = 296.813

	pointsOfPolynomialPlot[1_627].X = 16.27
	pointsOfPolynomialPlot[1_627].Y = 298.201

	pointsOfPolynomialPlot[1_628].X = 16.28
	pointsOfPolynomialPlot[1_628].Y = 299.594

	pointsOfPolynomialPlot[1_629].X = 16.29
	pointsOfPolynomialPlot[1_629].Y = 300.992

	pointsOfPolynomialPlot[1_630].X = 16.3
	pointsOfPolynomialPlot[1_630].Y = 302.395

	pointsOfPolynomialPlot[1_631].X = 16.31
	pointsOfPolynomialPlot[1_631].Y = 303.802

	pointsOfPolynomialPlot[1_632].X = 16.32
	pointsOfPolynomialPlot[1_632].Y = 305.215

	pointsOfPolynomialPlot[1_633].X = 16.33
	pointsOfPolynomialPlot[1_633].Y = 306.632

	pointsOfPolynomialPlot[1_634].X = 16.34
	pointsOfPolynomialPlot[1_634].Y = 308.055

	pointsOfPolynomialPlot[1_635].X = 16.35
	pointsOfPolynomialPlot[1_635].Y = 309.482

	pointsOfPolynomialPlot[1_636].X = 16.36
	pointsOfPolynomialPlot[1_636].Y = 310.915

	pointsOfPolynomialPlot[1_637].X = 16.37
	pointsOfPolynomialPlot[1_637].Y = 312.352

	pointsOfPolynomialPlot[1_638].X = 16.38
	pointsOfPolynomialPlot[1_638].Y = 313.795

	pointsOfPolynomialPlot[1_639].X = 16.39
	pointsOfPolynomialPlot[1_639].Y = 315.242

	pointsOfPolynomialPlot[1_640].X = 16.4
	pointsOfPolynomialPlot[1_640].Y = 316.694

	pointsOfPolynomialPlot[1_641].X = 16.41
	pointsOfPolynomialPlot[1_641].Y = 318.152

	pointsOfPolynomialPlot[1_642].X = 16.42
	pointsOfPolynomialPlot[1_642].Y = 319.614

	pointsOfPolynomialPlot[1_643].X = 16.43
	pointsOfPolynomialPlot[1_643].Y = 321.082

	pointsOfPolynomialPlot[1_644].X = 16.44
	pointsOfPolynomialPlot[1_644].Y = 322.554

	pointsOfPolynomialPlot[1_645].X = 16.45
	pointsOfPolynomialPlot[1_645].Y = 324.032

	pointsOfPolynomialPlot[1_646].X = 16.46
	pointsOfPolynomialPlot[1_646].Y = 325.515

	pointsOfPolynomialPlot[1_647].X = 16.47
	pointsOfPolynomialPlot[1_647].Y = 327.002

	pointsOfPolynomialPlot[1_648].X = 16.48
	pointsOfPolynomialPlot[1_648].Y = 328.495

	pointsOfPolynomialPlot[1_649].X = 16.49
	pointsOfPolynomialPlot[1_649].Y = 329.993

	pointsOfPolynomialPlot[1_650].X = 16.5
	pointsOfPolynomialPlot[1_650].Y = 331.497

	pointsOfPolynomialPlot[1_651].X = 16.51
	pointsOfPolynomialPlot[1_651].Y = 333.005

	pointsOfPolynomialPlot[1_652].X = 16.52
	pointsOfPolynomialPlot[1_652].Y = 334.519

	pointsOfPolynomialPlot[1_653].X = 16.53
	pointsOfPolynomialPlot[1_653].Y = 336.037

	pointsOfPolynomialPlot[1_654].X = 16.54
	pointsOfPolynomialPlot[1_654].Y = 337.561

	pointsOfPolynomialPlot[1_655].X = 16.55
	pointsOfPolynomialPlot[1_655].Y = 339.09

	pointsOfPolynomialPlot[1_656].X = 16.56
	pointsOfPolynomialPlot[1_656].Y = 340.624

	pointsOfPolynomialPlot[1_657].X = 16.57
	pointsOfPolynomialPlot[1_657].Y = 342.164

	pointsOfPolynomialPlot[1_658].X = 16.58
	pointsOfPolynomialPlot[1_658].Y = 343.709

	pointsOfPolynomialPlot[1_659].X = 16.59
	pointsOfPolynomialPlot[1_659].Y = 345.259

	pointsOfPolynomialPlot[1_660].X = 16.6
	pointsOfPolynomialPlot[1_660].Y = 346.814

	pointsOfPolynomialPlot[1_661].X = 16.61
	pointsOfPolynomialPlot[1_661].Y = 348.375

	pointsOfPolynomialPlot[1_662].X = 16.62
	pointsOfPolynomialPlot[1_662].Y = 349.94

	pointsOfPolynomialPlot[1_663].X = 16.63
	pointsOfPolynomialPlot[1_663].Y = 351.511

	pointsOfPolynomialPlot[1_664].X = 16.64
	pointsOfPolynomialPlot[1_664].Y = 353.088

	pointsOfPolynomialPlot[1_665].X = 16.65
	pointsOfPolynomialPlot[1_665].Y = 354.67

	pointsOfPolynomialPlot[1_666].X = 16.66
	pointsOfPolynomialPlot[1_666].Y = 356.257

	pointsOfPolynomialPlot[1_667].X = 16.67
	pointsOfPolynomialPlot[1_667].Y = 357.849

	pointsOfPolynomialPlot[1_668].X = 16.68
	pointsOfPolynomialPlot[1_668].Y = 359.447

	pointsOfPolynomialPlot[1_669].X = 16.69
	pointsOfPolynomialPlot[1_669].Y = 361.05

	pointsOfPolynomialPlot[1_670].X = 16.7
	pointsOfPolynomialPlot[1_670].Y = 362.659

	pointsOfPolynomialPlot[1_671].X = 16.71
	pointsOfPolynomialPlot[1_671].Y = 364.273

	pointsOfPolynomialPlot[1_672].X = 16.72
	pointsOfPolynomialPlot[1_672].Y = 365.892

	pointsOfPolynomialPlot[1_673].X = 16.73
	pointsOfPolynomialPlot[1_673].Y = 367.517

	pointsOfPolynomialPlot[1_674].X = 16.74
	pointsOfPolynomialPlot[1_674].Y = 369.147

	pointsOfPolynomialPlot[1_675].X = 16.75
	pointsOfPolynomialPlot[1_675].Y = 370.783

	pointsOfPolynomialPlot[1_676].X = 16.76
	pointsOfPolynomialPlot[1_676].Y = 372.424

	pointsOfPolynomialPlot[1_677].X = 16.77
	pointsOfPolynomialPlot[1_677].Y = 374.071

	pointsOfPolynomialPlot[1_678].X = 16.78
	pointsOfPolynomialPlot[1_678].Y = 375.723

	pointsOfPolynomialPlot[1_679].X = 16.79
	pointsOfPolynomialPlot[1_679].Y = 377.38

	pointsOfPolynomialPlot[1_680].X = 16.8
	pointsOfPolynomialPlot[1_680].Y = 379.043

	pointsOfPolynomialPlot[1_681].X = 16.81
	pointsOfPolynomialPlot[1_681].Y = 380.712

	pointsOfPolynomialPlot[1_682].X = 16.82
	pointsOfPolynomialPlot[1_682].Y = 382.386

	pointsOfPolynomialPlot[1_683].X = 16.83
	pointsOfPolynomialPlot[1_683].Y = 384.066

	pointsOfPolynomialPlot[1_684].X = 16.84
	pointsOfPolynomialPlot[1_684].Y = 385.751

	pointsOfPolynomialPlot[1_685].X = 16.85
	pointsOfPolynomialPlot[1_685].Y = 387.442

	pointsOfPolynomialPlot[1_686].X = 16.86
	pointsOfPolynomialPlot[1_686].Y = 389.139

	pointsOfPolynomialPlot[1_687].X = 16.87
	pointsOfPolynomialPlot[1_687].Y = 390.841

	pointsOfPolynomialPlot[1_688].X = 16.88
	pointsOfPolynomialPlot[1_688].Y = 392.549

	pointsOfPolynomialPlot[1_689].X = 16.89
	pointsOfPolynomialPlot[1_689].Y = 394.262

	pointsOfPolynomialPlot[1_690].X = 16.9
	pointsOfPolynomialPlot[1_690].Y = 395.981

	pointsOfPolynomialPlot[1_691].X = 16.91
	pointsOfPolynomialPlot[1_691].Y = 397.706

	pointsOfPolynomialPlot[1_692].X = 16.92
	pointsOfPolynomialPlot[1_692].Y = 399.436

	pointsOfPolynomialPlot[1_693].X = 16.93
	pointsOfPolynomialPlot[1_693].Y = 401.172

	pointsOfPolynomialPlot[1_694].X = 16.94
	pointsOfPolynomialPlot[1_694].Y = 402.914

	pointsOfPolynomialPlot[1_695].X = 16.95
	pointsOfPolynomialPlot[1_695].Y = 404.661

	pointsOfPolynomialPlot[1_696].X = 16.96
	pointsOfPolynomialPlot[1_696].Y = 406.414

	pointsOfPolynomialPlot[1_697].X = 16.97
	pointsOfPolynomialPlot[1_697].Y = 408.173

	pointsOfPolynomialPlot[1_698].X = 16.98
	pointsOfPolynomialPlot[1_698].Y = 409.938

	pointsOfPolynomialPlot[1_699].X = 16.99
	pointsOfPolynomialPlot[1_699].Y = 411.708

	pointsOfPolynomialPlot[1_700].X = 17.0
	pointsOfPolynomialPlot[1_700].Y = 413.484

	pointsOfPolynomialPlot[1_701].X = 17.01
	pointsOfPolynomialPlot[1_701].Y = 415.266

	pointsOfPolynomialPlot[1_702].X = 17.02
	pointsOfPolynomialPlot[1_702].Y = 417.054

	pointsOfPolynomialPlot[1_703].X = 17.03
	pointsOfPolynomialPlot[1_703].Y = 418.847

	pointsOfPolynomialPlot[1_704].X = 17.04
	pointsOfPolynomialPlot[1_704].Y = 420.647

	pointsOfPolynomialPlot[1_705].X = 17.05
	pointsOfPolynomialPlot[1_705].Y = 422.452

	pointsOfPolynomialPlot[1_706].X = 17.06
	pointsOfPolynomialPlot[1_706].Y = 424.263

	pointsOfPolynomialPlot[1_707].X = 17.07
	pointsOfPolynomialPlot[1_707].Y = 426.08

	pointsOfPolynomialPlot[1_708].X = 17.08
	pointsOfPolynomialPlot[1_708].Y = 427.903

	pointsOfPolynomialPlot[1_709].X = 17.09
	pointsOfPolynomialPlot[1_709].Y = 429.732

	pointsOfPolynomialPlot[1_710].X = 17.1
	pointsOfPolynomialPlot[1_710].Y = 431.566

	pointsOfPolynomialPlot[1_711].X = 17.11
	pointsOfPolynomialPlot[1_711].Y = 433.407

	pointsOfPolynomialPlot[1_712].X = 17.12
	pointsOfPolynomialPlot[1_712].Y = 435.253

	pointsOfPolynomialPlot[1_713].X = 17.13
	pointsOfPolynomialPlot[1_713].Y = 437.106

	pointsOfPolynomialPlot[1_714].X = 17.14
	pointsOfPolynomialPlot[1_714].Y = 438.964

	pointsOfPolynomialPlot[1_715].X = 17.15
	pointsOfPolynomialPlot[1_715].Y = 440.828

	pointsOfPolynomialPlot[1_716].X = 17.16
	pointsOfPolynomialPlot[1_716].Y = 442.699

	pointsOfPolynomialPlot[1_717].X = 17.17
	pointsOfPolynomialPlot[1_717].Y = 444.575

	pointsOfPolynomialPlot[1_718].X = 17.18
	pointsOfPolynomialPlot[1_718].Y = 446.458

	pointsOfPolynomialPlot[1_719].X = 17.19
	pointsOfPolynomialPlot[1_719].Y = 448.346

	pointsOfPolynomialPlot[1_720].X = 17.2
	pointsOfPolynomialPlot[1_720].Y = 450.24

	pointsOfPolynomialPlot[1_721].X = 17.21
	pointsOfPolynomialPlot[1_721].Y = 452.141

	pointsOfPolynomialPlot[1_722].X = 17.22
	pointsOfPolynomialPlot[1_722].Y = 454.048

	pointsOfPolynomialPlot[1_723].X = 17.23
	pointsOfPolynomialPlot[1_723].Y = 455.96

	pointsOfPolynomialPlot[1_724].X = 17.24
	pointsOfPolynomialPlot[1_724].Y = 457.879

	pointsOfPolynomialPlot[1_725].X = 17.25
	pointsOfPolynomialPlot[1_725].Y = 459.804

	pointsOfPolynomialPlot[1_726].X = 17.26
	pointsOfPolynomialPlot[1_726].Y = 461.735

	pointsOfPolynomialPlot[1_727].X = 17.27
	pointsOfPolynomialPlot[1_727].Y = 463.672

	pointsOfPolynomialPlot[1_728].X = 17.28
	pointsOfPolynomialPlot[1_728].Y = 465.615

	pointsOfPolynomialPlot[1_729].X = 17.29
	pointsOfPolynomialPlot[1_729].Y = 467.565

	pointsOfPolynomialPlot[1_730].X = 17.3
	pointsOfPolynomialPlot[1_730].Y = 469.52

	pointsOfPolynomialPlot[1_731].X = 17.31
	pointsOfPolynomialPlot[1_731].Y = 471.482

	pointsOfPolynomialPlot[1_732].X = 17.32
	pointsOfPolynomialPlot[1_732].Y = 473.45

	pointsOfPolynomialPlot[1_733].X = 17.33
	pointsOfPolynomialPlot[1_733].Y = 475.425

	pointsOfPolynomialPlot[1_734].X = 17.34
	pointsOfPolynomialPlot[1_734].Y = 477.405

	pointsOfPolynomialPlot[1_735].X = 17.35
	pointsOfPolynomialPlot[1_735].Y = 479.392

	pointsOfPolynomialPlot[1_736].X = 17.36
	pointsOfPolynomialPlot[1_736].Y = 481.385

	pointsOfPolynomialPlot[1_737].X = 17.37
	pointsOfPolynomialPlot[1_737].Y = 483.384

	pointsOfPolynomialPlot[1_738].X = 17.38
	pointsOfPolynomialPlot[1_738].Y = 485.39

	pointsOfPolynomialPlot[1_739].X = 17.39
	pointsOfPolynomialPlot[1_739].Y = 487.402

	pointsOfPolynomialPlot[1_740].X = 17.4
	pointsOfPolynomialPlot[1_740].Y = 489.42

	pointsOfPolynomialPlot[1_741].X = 17.41
	pointsOfPolynomialPlot[1_741].Y = 491.445

	pointsOfPolynomialPlot[1_742].X = 17.42
	pointsOfPolynomialPlot[1_742].Y = 493.475

	pointsOfPolynomialPlot[1_743].X = 17.43
	pointsOfPolynomialPlot[1_743].Y = 495.513

	pointsOfPolynomialPlot[1_744].X = 17.44
	pointsOfPolynomialPlot[1_744].Y = 497.556

	pointsOfPolynomialPlot[1_745].X = 17.45
	pointsOfPolynomialPlot[1_745].Y = 499.606

	pointsOfPolynomialPlot[1_746].X = 17.46
	pointsOfPolynomialPlot[1_746].Y = 501.663

	pointsOfPolynomialPlot[1_747].X = 17.47
	pointsOfPolynomialPlot[1_747].Y = 503.726

	pointsOfPolynomialPlot[1_748].X = 17.48
	pointsOfPolynomialPlot[1_748].Y = 505.795

	pointsOfPolynomialPlot[1_749].X = 17.49
	pointsOfPolynomialPlot[1_749].Y = 507.871

	pointsOfPolynomialPlot[1_750].X = 17.5
	pointsOfPolynomialPlot[1_750].Y = 509.953

	pointsOfPolynomialPlot[1_751].X = 17.51
	pointsOfPolynomialPlot[1_751].Y = 512.042

	pointsOfPolynomialPlot[1_752].X = 17.52
	pointsOfPolynomialPlot[1_752].Y = 514.137

	pointsOfPolynomialPlot[1_753].X = 17.53
	pointsOfPolynomialPlot[1_753].Y = 516.238

	pointsOfPolynomialPlot[1_754].X = 17.54
	pointsOfPolynomialPlot[1_754].Y = 518.347

	pointsOfPolynomialPlot[1_755].X = 17.55
	pointsOfPolynomialPlot[1_755].Y = 520.461

	pointsOfPolynomialPlot[1_756].X = 17.56
	pointsOfPolynomialPlot[1_756].Y = 522.583

	pointsOfPolynomialPlot[1_757].X = 17.57
	pointsOfPolynomialPlot[1_757].Y = 524.71

	pointsOfPolynomialPlot[1_758].X = 17.58
	pointsOfPolynomialPlot[1_758].Y = 526.845

	pointsOfPolynomialPlot[1_759].X = 17.59
	pointsOfPolynomialPlot[1_759].Y = 528.986

	pointsOfPolynomialPlot[1_760].X = 17.6
	pointsOfPolynomialPlot[1_760].Y = 531.133

	pointsOfPolynomialPlot[1_761].X = 17.61
	pointsOfPolynomialPlot[1_761].Y = 533.288

	pointsOfPolynomialPlot[1_762].X = 17.62
	pointsOfPolynomialPlot[1_762].Y = 535.448

	pointsOfPolynomialPlot[1_763].X = 17.63
	pointsOfPolynomialPlot[1_763].Y = 537.616

	pointsOfPolynomialPlot[1_764].X = 17.64
	pointsOfPolynomialPlot[1_764].Y = 539.79

	pointsOfPolynomialPlot[1_765].X = 17.65
	pointsOfPolynomialPlot[1_765].Y = 541.971

	pointsOfPolynomialPlot[1_766].X = 17.66
	pointsOfPolynomialPlot[1_766].Y = 544.158

	pointsOfPolynomialPlot[1_767].X = 17.67
	pointsOfPolynomialPlot[1_767].Y = 546.353

	pointsOfPolynomialPlot[1_768].X = 17.68
	pointsOfPolynomialPlot[1_768].Y = 548.554

	pointsOfPolynomialPlot[1_769].X = 17.69
	pointsOfPolynomialPlot[1_769].Y = 550.761

	pointsOfPolynomialPlot[1_770].X = 17.7
	pointsOfPolynomialPlot[1_770].Y = 552.976

	pointsOfPolynomialPlot[1_771].X = 17.71
	pointsOfPolynomialPlot[1_771].Y = 555.197

	pointsOfPolynomialPlot[1_772].X = 17.72
	pointsOfPolynomialPlot[1_772].Y = 557.425

	pointsOfPolynomialPlot[1_773].X = 17.73
	pointsOfPolynomialPlot[1_773].Y = 559.66

	pointsOfPolynomialPlot[1_774].X = 17.74
	pointsOfPolynomialPlot[1_774].Y = 561.901

	pointsOfPolynomialPlot[1_775].X = 17.75
	pointsOfPolynomialPlot[1_775].Y = 564.15

	pointsOfPolynomialPlot[1_776].X = 17.76
	pointsOfPolynomialPlot[1_776].Y = 566.405

	pointsOfPolynomialPlot[1_777].X = 17.77
	pointsOfPolynomialPlot[1_777].Y = 568.667

	pointsOfPolynomialPlot[1_778].X = 17.78
	pointsOfPolynomialPlot[1_778].Y = 570.936

	pointsOfPolynomialPlot[1_779].X = 17.79
	pointsOfPolynomialPlot[1_779].Y = 573.212

	pointsOfPolynomialPlot[1_780].X = 17.8
	pointsOfPolynomialPlot[1_780].Y = 575.494

	pointsOfPolynomialPlot[1_781].X = 17.81
	pointsOfPolynomialPlot[1_781].Y = 577.784

	pointsOfPolynomialPlot[1_782].X = 17.82
	pointsOfPolynomialPlot[1_782].Y = 580.081

	pointsOfPolynomialPlot[1_783].X = 17.83
	pointsOfPolynomialPlot[1_783].Y = 582.384

	pointsOfPolynomialPlot[1_784].X = 17.84
	pointsOfPolynomialPlot[1_784].Y = 584.694

	pointsOfPolynomialPlot[1_785].X = 17.85
	pointsOfPolynomialPlot[1_785].Y = 587.012

	pointsOfPolynomialPlot[1_786].X = 17.86
	pointsOfPolynomialPlot[1_786].Y = 589.336

	pointsOfPolynomialPlot[1_787].X = 17.87
	pointsOfPolynomialPlot[1_787].Y = 591.668

	pointsOfPolynomialPlot[1_788].X = 17.88
	pointsOfPolynomialPlot[1_788].Y = 594.006

	pointsOfPolynomialPlot[1_789].X = 17.89
	pointsOfPolynomialPlot[1_789].Y = 596.352

	pointsOfPolynomialPlot[1_790].X = 17.9
	pointsOfPolynomialPlot[1_790].Y = 598.704

	pointsOfPolynomialPlot[1_791].X = 17.91
	pointsOfPolynomialPlot[1_791].Y = 601.064

	pointsOfPolynomialPlot[1_792].X = 17.92
	pointsOfPolynomialPlot[1_792].Y = 603.43

	pointsOfPolynomialPlot[1_793].X = 17.93
	pointsOfPolynomialPlot[1_793].Y = 605.804

	pointsOfPolynomialPlot[1_794].X = 17.94
	pointsOfPolynomialPlot[1_794].Y = 608.185

	pointsOfPolynomialPlot[1_795].X = 17.95
	pointsOfPolynomialPlot[1_795].Y = 610.573

	pointsOfPolynomialPlot[1_796].X = 17.96
	pointsOfPolynomialPlot[1_796].Y = 612.968

	pointsOfPolynomialPlot[1_797].X = 17.97
	pointsOfPolynomialPlot[1_797].Y = 615.37

	pointsOfPolynomialPlot[1_798].X = 17.98
	pointsOfPolynomialPlot[1_798].Y = 617.779

	pointsOfPolynomialPlot[1_799].X = 17.99
	pointsOfPolynomialPlot[1_799].Y = 620.196

	pointsOfPolynomialPlot[1_800].X = 18.0
	pointsOfPolynomialPlot[1_800].Y = 622.62

	pointsOfPolynomialPlot[1_801].X = 18.01
	pointsOfPolynomialPlot[1_801].Y = 625.051

	pointsOfPolynomialPlot[1_802].X = 18.02
	pointsOfPolynomialPlot[1_802].Y = 627.489

	pointsOfPolynomialPlot[1_803].X = 18.03
	pointsOfPolynomialPlot[1_803].Y = 629.934

	pointsOfPolynomialPlot[1_804].X = 18.04
	pointsOfPolynomialPlot[1_804].Y = 632.387

	pointsOfPolynomialPlot[1_805].X = 18.05
	pointsOfPolynomialPlot[1_805].Y = 634.847

	pointsOfPolynomialPlot[1_806].X = 18.06
	pointsOfPolynomialPlot[1_806].Y = 637.314

	pointsOfPolynomialPlot[1_807].X = 18.07
	pointsOfPolynomialPlot[1_807].Y = 639.789

	pointsOfPolynomialPlot[1_808].X = 18.08
	pointsOfPolynomialPlot[1_808].Y = 642.271

	pointsOfPolynomialPlot[1_809].X = 18.09
	pointsOfPolynomialPlot[1_809].Y = 644.76

	pointsOfPolynomialPlot[1_810].X = 18.1
	pointsOfPolynomialPlot[1_810].Y = 647.257

	pointsOfPolynomialPlot[1_811].X = 18.11
	pointsOfPolynomialPlot[1_811].Y = 649.761

	pointsOfPolynomialPlot[1_812].X = 18.12
	pointsOfPolynomialPlot[1_812].Y = 652.272

	pointsOfPolynomialPlot[1_813].X = 18.13
	pointsOfPolynomialPlot[1_813].Y = 654.791

	pointsOfPolynomialPlot[1_814].X = 18.14
	pointsOfPolynomialPlot[1_814].Y = 657.317

	pointsOfPolynomialPlot[1_815].X = 18.15
	pointsOfPolynomialPlot[1_815].Y = 659.85

	pointsOfPolynomialPlot[1_816].X = 18.16
	pointsOfPolynomialPlot[1_816].Y = 662.391

	pointsOfPolynomialPlot[1_817].X = 18.17
	pointsOfPolynomialPlot[1_817].Y = 664.94

	pointsOfPolynomialPlot[1_818].X = 18.18
	pointsOfPolynomialPlot[1_818].Y = 667.496

	pointsOfPolynomialPlot[1_819].X = 18.19
	pointsOfPolynomialPlot[1_819].Y = 670.059

	pointsOfPolynomialPlot[1_820].X = 18.2
	pointsOfPolynomialPlot[1_820].Y = 672.63

	pointsOfPolynomialPlot[1_821].X = 18.21
	pointsOfPolynomialPlot[1_821].Y = 675.208

	pointsOfPolynomialPlot[1_822].X = 18.22
	pointsOfPolynomialPlot[1_822].Y = 677.794

	pointsOfPolynomialPlot[1_823].X = 18.23
	pointsOfPolynomialPlot[1_823].Y = 680.388

	pointsOfPolynomialPlot[1_824].X = 18.24
	pointsOfPolynomialPlot[1_824].Y = 682.989

	pointsOfPolynomialPlot[1_825].X = 18.25
	pointsOfPolynomialPlot[1_825].Y = 685.598

	pointsOfPolynomialPlot[1_826].X = 18.26
	pointsOfPolynomialPlot[1_826].Y = 688.214

	pointsOfPolynomialPlot[1_827].X = 18.27
	pointsOfPolynomialPlot[1_827].Y = 690.838

	pointsOfPolynomialPlot[1_828].X = 18.28
	pointsOfPolynomialPlot[1_828].Y = 693.469

	pointsOfPolynomialPlot[1_829].X = 18.29
	pointsOfPolynomialPlot[1_829].Y = 696.108

	pointsOfPolynomialPlot[1_830].X = 18.3
	pointsOfPolynomialPlot[1_830].Y = 698.755

	pointsOfPolynomialPlot[1_831].X = 18.31
	pointsOfPolynomialPlot[1_831].Y = 701.41

	pointsOfPolynomialPlot[1_832].X = 18.32
	pointsOfPolynomialPlot[1_832].Y = 704.072

	pointsOfPolynomialPlot[1_833].X = 18.33
	pointsOfPolynomialPlot[1_833].Y = 706.742

	pointsOfPolynomialPlot[1_834].X = 18.34
	pointsOfPolynomialPlot[1_834].Y = 709.419

	pointsOfPolynomialPlot[1_835].X = 18.35
	pointsOfPolynomialPlot[1_835].Y = 712.105

	pointsOfPolynomialPlot[1_836].X = 18.36
	pointsOfPolynomialPlot[1_836].Y = 714.798

	pointsOfPolynomialPlot[1_837].X = 18.37
	pointsOfPolynomialPlot[1_837].Y = 717.499

	pointsOfPolynomialPlot[1_838].X = 18.38
	pointsOfPolynomialPlot[1_838].Y = 720.207

	pointsOfPolynomialPlot[1_839].X = 18.39
	pointsOfPolynomialPlot[1_839].Y = 722.924

	pointsOfPolynomialPlot[1_840].X = 18.4
	pointsOfPolynomialPlot[1_840].Y = 725.648

	pointsOfPolynomialPlot[1_841].X = 18.41
	pointsOfPolynomialPlot[1_841].Y = 728.38

	pointsOfPolynomialPlot[1_842].X = 18.42
	pointsOfPolynomialPlot[1_842].Y = 731.12

	pointsOfPolynomialPlot[1_843].X = 18.43
	pointsOfPolynomialPlot[1_843].Y = 733.868

	pointsOfPolynomialPlot[1_844].X = 18.44
	pointsOfPolynomialPlot[1_844].Y = 736.624

	pointsOfPolynomialPlot[1_845].X = 18.45
	pointsOfPolynomialPlot[1_845].Y = 739.387

	pointsOfPolynomialPlot[1_846].X = 18.46
	pointsOfPolynomialPlot[1_846].Y = 742.159

	pointsOfPolynomialPlot[1_847].X = 18.47
	pointsOfPolynomialPlot[1_847].Y = 744.938

	pointsOfPolynomialPlot[1_848].X = 18.48
	pointsOfPolynomialPlot[1_848].Y = 747.726

	pointsOfPolynomialPlot[1_849].X = 18.49
	pointsOfPolynomialPlot[1_849].Y = 750.521

	pointsOfPolynomialPlot[1_850].X = 18.5
	pointsOfPolynomialPlot[1_850].Y = 753.325

	pointsOfPolynomialPlot[1_851].X = 18.51
	pointsOfPolynomialPlot[1_851].Y = 756.136

	pointsOfPolynomialPlot[1_852].X = 18.52
	pointsOfPolynomialPlot[1_852].Y = 758.955

	pointsOfPolynomialPlot[1_853].X = 18.53
	pointsOfPolynomialPlot[1_853].Y = 761.783

	pointsOfPolynomialPlot[1_854].X = 18.54
	pointsOfPolynomialPlot[1_854].Y = 764.618

	pointsOfPolynomialPlot[1_855].X = 18.55
	pointsOfPolynomialPlot[1_855].Y = 767.462

	pointsOfPolynomialPlot[1_856].X = 18.56
	pointsOfPolynomialPlot[1_856].Y = 770.313

	pointsOfPolynomialPlot[1_857].X = 18.57
	pointsOfPolynomialPlot[1_857].Y = 773.173

	pointsOfPolynomialPlot[1_858].X = 18.58
	pointsOfPolynomialPlot[1_858].Y = 776.041

	pointsOfPolynomialPlot[1_859].X = 18.59
	pointsOfPolynomialPlot[1_859].Y = 778.917

	pointsOfPolynomialPlot[1_860].X = 18.6
	pointsOfPolynomialPlot[1_860].Y = 781.801

	pointsOfPolynomialPlot[1_861].X = 18.61
	pointsOfPolynomialPlot[1_861].Y = 784.693

	pointsOfPolynomialPlot[1_862].X = 18.62
	pointsOfPolynomialPlot[1_862].Y = 787.594

	pointsOfPolynomialPlot[1_863].X = 18.63
	pointsOfPolynomialPlot[1_863].Y = 790.502

	pointsOfPolynomialPlot[1_864].X = 18.64
	pointsOfPolynomialPlot[1_864].Y = 793.419

	pointsOfPolynomialPlot[1_865].X = 18.65
	pointsOfPolynomialPlot[1_865].Y = 796.344

	pointsOfPolynomialPlot[1_866].X = 18.66
	pointsOfPolynomialPlot[1_866].Y = 799.277

	pointsOfPolynomialPlot[1_867].X = 18.67
	pointsOfPolynomialPlot[1_867].Y = 802.219

	pointsOfPolynomialPlot[1_868].X = 18.68
	pointsOfPolynomialPlot[1_868].Y = 805.169

	pointsOfPolynomialPlot[1_869].X = 18.69
	pointsOfPolynomialPlot[1_869].Y = 808.127

	pointsOfPolynomialPlot[1_870].X = 18.7
	pointsOfPolynomialPlot[1_870].Y = 811.093

	pointsOfPolynomialPlot[1_871].X = 18.71
	pointsOfPolynomialPlot[1_871].Y = 814.068

	pointsOfPolynomialPlot[1_872].X = 18.72
	pointsOfPolynomialPlot[1_872].Y = 817.051

	pointsOfPolynomialPlot[1_873].X = 18.73
	pointsOfPolynomialPlot[1_873].Y = 820.043

	pointsOfPolynomialPlot[1_874].X = 18.74
	pointsOfPolynomialPlot[1_874].Y = 823.043

	pointsOfPolynomialPlot[1_875].X = 18.75
	pointsOfPolynomialPlot[1_875].Y = 826.051

	pointsOfPolynomialPlot[1_876].X = 18.76
	pointsOfPolynomialPlot[1_876].Y = 829.067

	pointsOfPolynomialPlot[1_877].X = 18.77
	pointsOfPolynomialPlot[1_877].Y = 832.093

	pointsOfPolynomialPlot[1_878].X = 18.78
	pointsOfPolynomialPlot[1_878].Y = 835.126

	pointsOfPolynomialPlot[1_879].X = 18.79
	pointsOfPolynomialPlot[1_879].Y = 838.126

	pointsOfPolynomialPlot[1_880].X = 18.8
	pointsOfPolynomialPlot[1_880].Y = 841.219

	pointsOfPolynomialPlot[1_881].X = 18.81
	pointsOfPolynomialPlot[1_881].Y = 844.278

	pointsOfPolynomialPlot[1_882].X = 18.82
	pointsOfPolynomialPlot[1_882].Y = 847.345

	pointsOfPolynomialPlot[1_883].X = 18.83
	pointsOfPolynomialPlot[1_883].Y = 850.421

	pointsOfPolynomialPlot[1_884].X = 18.84
	pointsOfPolynomialPlot[1_884].Y = 853.506

	pointsOfPolynomialPlot[1_885].X = 18.85
	pointsOfPolynomialPlot[1_885].Y = 856.599

	pointsOfPolynomialPlot[1_886].X = 18.86
	pointsOfPolynomialPlot[1_886].Y = 859.7

	pointsOfPolynomialPlot[1_887].X = 18.87
	pointsOfPolynomialPlot[1_887].Y = 862.811

	pointsOfPolynomialPlot[1_888].X = 18.88
	pointsOfPolynomialPlot[1_888].Y = 865.93

	pointsOfPolynomialPlot[1_889].X = 18.89
	pointsOfPolynomialPlot[1_889].Y = 869.057

	pointsOfPolynomialPlot[1_890].X = 18.9
	pointsOfPolynomialPlot[1_890].Y = 872.193

	pointsOfPolynomialPlot[1_891].X = 18.91
	pointsOfPolynomialPlot[1_891].Y = 875.338

	pointsOfPolynomialPlot[1_892].X = 18.92
	pointsOfPolynomialPlot[1_892].Y = 878.492

	pointsOfPolynomialPlot[1_893].X = 18.93
	pointsOfPolynomialPlot[1_893].Y = 881.654

	pointsOfPolynomialPlot[1_894].X = 18.94
	pointsOfPolynomialPlot[1_894].Y = 884.825

	pointsOfPolynomialPlot[1_895].X = 18.95
	pointsOfPolynomialPlot[1_895].Y = 888.004

	pointsOfPolynomialPlot[1_896].X = 18.96
	pointsOfPolynomialPlot[1_896].Y = 891.193

	pointsOfPolynomialPlot[1_897].X = 18.97
	pointsOfPolynomialPlot[1_897].Y = 894.39

	pointsOfPolynomialPlot[1_898].X = 18.98
	pointsOfPolynomialPlot[1_898].Y = 897.596

	pointsOfPolynomialPlot[1_899].X = 18.99
	pointsOfPolynomialPlot[1_899].Y = 900.811

	pointsOfPolynomialPlot[1_900].X = 19.0
	pointsOfPolynomialPlot[1_900].Y = 904.034

	pointsOfPolynomialPlot[1_901].X = 19.01
	pointsOfPolynomialPlot[1_901].Y = 907.267

	pointsOfPolynomialPlot[1_902].X = 19.02
	pointsOfPolynomialPlot[1_902].Y = 910.508

	pointsOfPolynomialPlot[1_903].X = 19.03
	pointsOfPolynomialPlot[1_903].Y = 913.758

	pointsOfPolynomialPlot[1_904].X = 19.04
	pointsOfPolynomialPlot[1_904].Y = 917.017

	pointsOfPolynomialPlot[1_905].X = 19.05
	pointsOfPolynomialPlot[1_905].Y = 920.285

	pointsOfPolynomialPlot[1_906].X = 19.06
	pointsOfPolynomialPlot[1_906].Y = 923.562

	pointsOfPolynomialPlot[1_907].X = 19.07
	pointsOfPolynomialPlot[1_907].Y = 926.848

	pointsOfPolynomialPlot[1_908].X = 19.08
	pointsOfPolynomialPlot[1_908].Y = 930.143

	pointsOfPolynomialPlot[1_909].X = 19.09
	pointsOfPolynomialPlot[1_909].Y = 933.447

	pointsOfPolynomialPlot[1_910].X = 19.1
	pointsOfPolynomialPlot[1_910].Y = 936.759

	pointsOfPolynomialPlot[1_911].X = 19.11
	pointsOfPolynomialPlot[1_911].Y = 940.081

	pointsOfPolynomialPlot[1_912].X = 19.12
	pointsOfPolynomialPlot[1_912].Y = 943.412

	pointsOfPolynomialPlot[1_913].X = 19.13
	pointsOfPolynomialPlot[1_913].Y = 946.752

	pointsOfPolynomialPlot[1_914].X = 19.14
	pointsOfPolynomialPlot[1_914].Y = 950.101

	pointsOfPolynomialPlot[1_915].X = 19.15
	pointsOfPolynomialPlot[1_915].Y = 953.459

	pointsOfPolynomialPlot[1_916].X = 19.16
	pointsOfPolynomialPlot[1_916].Y = 956.826

	pointsOfPolynomialPlot[1_917].X = 19.17
	pointsOfPolynomialPlot[1_917].Y = 960.202

	pointsOfPolynomialPlot[1_918].X = 19.18
	pointsOfPolynomialPlot[1_918].Y = 963.587

	pointsOfPolynomialPlot[1_919].X = 19.19
	pointsOfPolynomialPlot[1_919].Y = 966.982

	pointsOfPolynomialPlot[1_920].X = 19.2
	pointsOfPolynomialPlot[1_920].Y = 970.385

	pointsOfPolynomialPlot[1_921].X = 19.21
	pointsOfPolynomialPlot[1_921].Y = 973.798

	pointsOfPolynomialPlot[1_922].X = 19.22
	pointsOfPolynomialPlot[1_922].Y = 977.22

	pointsOfPolynomialPlot[1_923].X = 19.23
	pointsOfPolynomialPlot[1_923].Y = 980.652

	pointsOfPolynomialPlot[1_924].X = 19.24
	pointsOfPolynomialPlot[1_924].Y = 984.092

	pointsOfPolynomialPlot[1_925].X = 19.25
	pointsOfPolynomialPlot[1_925].Y = 987.542

	pointsOfPolynomialPlot[1_926].X = 19.26
	pointsOfPolynomialPlot[1_926].Y = 991.001

	pointsOfPolynomialPlot[1_927].X = 19.27
	pointsOfPolynomialPlot[1_927].Y = 994.469

	pointsOfPolynomialPlot[1_928].X = 19.28
	pointsOfPolynomialPlot[1_928].Y = 997.947

	pointsOfPolynomialPlot[1_929].X = 19.29
	pointsOfPolynomialPlot[1_929].Y = 1_001.434

	pointsOfPolynomialPlot[1_930].X = 19.3
	pointsOfPolynomialPlot[1_930].Y = 1_004.93

	pointsOfPolynomialPlot[1_931].X = 19.31
	pointsOfPolynomialPlot[1_931].Y = 1_008.346

	pointsOfPolynomialPlot[1_932].X = 19.32
	pointsOfPolynomialPlot[1_932].Y = 1_011.951

	pointsOfPolynomialPlot[1_933].X = 19.33
	pointsOfPolynomialPlot[1_933].Y = 1_015.476

	pointsOfPolynomialPlot[1_934].X = 19.34
	pointsOfPolynomialPlot[1_934].Y = 1_019.009

	pointsOfPolynomialPlot[1_935].X = 19.35
	pointsOfPolynomialPlot[1_935].Y = 1_022.553

	pointsOfPolynomialPlot[1_936].X = 19.36
	pointsOfPolynomialPlot[1_936].Y = 1_026.106

	pointsOfPolynomialPlot[1_937].X = 19.37
	pointsOfPolynomialPlot[1_937].Y = 1_029.668

	pointsOfPolynomialPlot[1_938].X = 19.38
	pointsOfPolynomialPlot[1_938].Y = 1_033.24

	pointsOfPolynomialPlot[1_939].X = 19.39
	pointsOfPolynomialPlot[1_939].Y = 1_036.821

	pointsOfPolynomialPlot[1_940].X = 19.4
	pointsOfPolynomialPlot[1_940].Y = 1_040.412

	pointsOfPolynomialPlot[1_941].X = 19.41
	pointsOfPolynomialPlot[1_941].Y = 1_044.012

	pointsOfPolynomialPlot[1_942].X = 19.42
	pointsOfPolynomialPlot[1_942].Y = 1_047.622

	pointsOfPolynomialPlot[1_943].X = 19.43
	pointsOfPolynomialPlot[1_943].Y = 1_051.242

	pointsOfPolynomialPlot[1_944].X = 19.44
	pointsOfPolynomialPlot[1_944].Y = 1_054.871

	pointsOfPolynomialPlot[1_945].X = 19.45
	pointsOfPolynomialPlot[1_945].Y = 1_058.871

	pointsOfPolynomialPlot[1_946].X = 19.46
	pointsOfPolynomialPlot[1_946].Y = 1_062.158

	pointsOfPolynomialPlot[1_947].X = 19.47
	pointsOfPolynomialPlot[1_947].Y = 1_065.816

	pointsOfPolynomialPlot[1_948].X = 19.48
	pointsOfPolynomialPlot[1_948].Y = 1_069.484

	pointsOfPolynomialPlot[1_949].X = 19.49
	pointsOfPolynomialPlot[1_949].Y = 1_073.161

	pointsOfPolynomialPlot[1_950].X = 19.5
	pointsOfPolynomialPlot[1_950].Y = 1_076.848

	pointsOfPolynomialPlot[1_951].X = 19.51
	pointsOfPolynomialPlot[1_951].Y = 1_080.545

	pointsOfPolynomialPlot[1_952].X = 19.52
	pointsOfPolynomialPlot[1_952].Y = 1_084.251

	pointsOfPolynomialPlot[1_953].X = 19.53
	pointsOfPolynomialPlot[1_953].Y = 1_087.968

	pointsOfPolynomialPlot[1_954].X = 19.54
	pointsOfPolynomialPlot[1_954].Y = 1_091.694

	pointsOfPolynomialPlot[1_955].X = 19.55
	pointsOfPolynomialPlot[1_955].Y = 1_095.43

	pointsOfPolynomialPlot[1_956].X = 19.56
	pointsOfPolynomialPlot[1_956].Y = 1_099.176

	pointsOfPolynomialPlot[1_957].X = 19.57
	pointsOfPolynomialPlot[1_957].Y = 1_102.931

	pointsOfPolynomialPlot[1_958].X = 19.58
	pointsOfPolynomialPlot[1_958].Y = 1_106.697

	pointsOfPolynomialPlot[1_959].X = 19.59
	pointsOfPolynomialPlot[1_959].Y = 1_110.472

	pointsOfPolynomialPlot[1_960].X = 19.6
	pointsOfPolynomialPlot[1_960].Y = 1_114.257

	pointsOfPolynomialPlot[1_961].X = 19.61
	pointsOfPolynomialPlot[1_961].Y = 1_118.053

	pointsOfPolynomialPlot[1_962].X = 19.62
	pointsOfPolynomialPlot[1_962].Y = 1_121.858

	pointsOfPolynomialPlot[1_963].X = 19.63
	pointsOfPolynomialPlot[1_963].Y = 1_125.673

	pointsOfPolynomialPlot[1_964].X = 19.64
	pointsOfPolynomialPlot[1_964].Y = 1_129.498

	pointsOfPolynomialPlot[1_965].X = 19.65
	pointsOfPolynomialPlot[1_965].Y = 1_133.333

	pointsOfPolynomialPlot[1_966].X = 19.66
	pointsOfPolynomialPlot[1_966].Y = 1_137.178

	pointsOfPolynomialPlot[1_967].X = 19.67
	pointsOfPolynomialPlot[1_967].Y = 1_141.033

	pointsOfPolynomialPlot[1_968].X = 19.68
	pointsOfPolynomialPlot[1_968].Y = 1_144.898

	pointsOfPolynomialPlot[1_969].X = 19.69
	pointsOfPolynomialPlot[1_969].Y = 1_148.773

	pointsOfPolynomialPlot[1_970].X = 19.7
	pointsOfPolynomialPlot[1_970].Y = 1_152.658

	pointsOfPolynomialPlot[1_971].X = 19.71
	pointsOfPolynomialPlot[1_971].Y = 1_156.554

	pointsOfPolynomialPlot[1_972].X = 19.72
	pointsOfPolynomialPlot[1_972].Y = 1_160.459

	pointsOfPolynomialPlot[1_973].X = 19.73
	pointsOfPolynomialPlot[1_973].Y =  1_164.375

	pointsOfPolynomialPlot[1_974].X = 19.74
	pointsOfPolynomialPlot[1_974].Y = 1_168.301

	pointsOfPolynomialPlot[1_975].X = 19.75
	pointsOfPolynomialPlot[1_975].Y = 1_172.236

	pointsOfPolynomialPlot[1_976].X = 19.76
	pointsOfPolynomialPlot[1_976].Y = 1_176.183

	pointsOfPolynomialPlot[1_977].X = 19.77
	pointsOfPolynomialPlot[1_977].Y = 1_180.139

	pointsOfPolynomialPlot[1_978].X = 19.78
	pointsOfPolynomialPlot[1_978].Y = 1_184.106

	pointsOfPolynomialPlot[1_979].X = 19.79
	pointsOfPolynomialPlot[1_979].Y = 1_188.082

	pointsOfPolynomialPlot[1_980].X = 19.8
	pointsOfPolynomialPlot[1_980].Y = 1_192.07

	pointsOfPolynomialPlot[1_981].X = 19.81
	pointsOfPolynomialPlot[1_981].Y = 1_196.067

	pointsOfPolynomialPlot[1_982].X = 19.82
	pointsOfPolynomialPlot[1_982].Y = 1_200.075

	pointsOfPolynomialPlot[1_983].X = 19.83
	pointsOfPolynomialPlot[1_983].Y = 1_204.093

	pointsOfPolynomialPlot[1_984].X = 19.84
	pointsOfPolynomialPlot[1_984].Y = 1_208.121

	pointsOfPolynomialPlot[1_985].X = 19.85
	pointsOfPolynomialPlot[1_985].Y = 1_212.16

	pointsOfPolynomialPlot[1_986].X = 19.86
	pointsOfPolynomialPlot[1_986].Y = 1_216.209

	pointsOfPolynomialPlot[1_987].X = 19.87
	pointsOfPolynomialPlot[1_987].Y = 1_220.269

	pointsOfPolynomialPlot[1_988].X = 19.88
	pointsOfPolynomialPlot[1_988].Y = 1_224.339

	pointsOfPolynomialPlot[1_989].X = 19.89
	pointsOfPolynomialPlot[1_989].Y = 1_228.419

	pointsOfPolynomialPlot[1_990].X = 19.9
	pointsOfPolynomialPlot[1_990].Y = 1_232.51

	pointsOfPolynomialPlot[1_991].X = 19.91
	pointsOfPolynomialPlot[1_991].Y = 1_236.611

	pointsOfPolynomialPlot[1_992].X = 19.92
	pointsOfPolynomialPlot[1_992].Y = 1_240.724

	pointsOfPolynomialPlot[1_993].X = 19.93
	pointsOfPolynomialPlot[1_993].Y = 1_244.846

	pointsOfPolynomialPlot[1_994].X = 19.94
	pointsOfPolynomialPlot[1_994].Y = 1_248.979

	pointsOfPolynomialPlot[1_995].X = 19.95
	pointsOfPolynomialPlot[1_995].Y = 1_253.123

	pointsOfPolynomialPlot[1_996].X = 19.96
	pointsOfPolynomialPlot[1_996].Y = 1_257.277

	pointsOfPolynomialPlot[1_997].X = 19.97
	pointsOfPolynomialPlot[1_997].Y = 1_261.442

	pointsOfPolynomialPlot[1_998].X = 19.98
	pointsOfPolynomialPlot[1_998].Y = 1_265.617

	pointsOfPolynomialPlot[1_999].X = 19.99
	pointsOfPolynomialPlot[1_999].Y = 1_269.803

	pointsOfPolynomialPlot[2_000].X = 20.0
	pointsOfPolynomialPlot[2_000].Y = 1_274.0










	polynomialPlot := plot.New()

	polynomialPlot.Title.Text = "Plot of polynomial f(x) = 0.000001x^6 - 0.001x^5 - 0.01x^4 - 0.1x^3 + x^2 + x - 10.0"

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
		"Polynomial-plot-01.png"); err != nil {

		panic(err)
	}
}
