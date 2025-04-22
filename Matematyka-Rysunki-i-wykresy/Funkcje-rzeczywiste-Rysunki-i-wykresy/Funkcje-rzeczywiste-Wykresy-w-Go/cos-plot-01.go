package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres funkcji f(x) = cos(x)

	punktyWykresuFunkcji := make(plotter.XYs, 321)

	punktyWykresuFunkcji[0].X = 0.0
	punktyWykresuFunkcji[0].Y = 0.0

	punktyWykresuFunkcji[1].X = 0.01
	punktyWykresuFunkcji[1].Y = 0.01

	punktyWykresuFunkcji[2].X = 0.02
	punktyWykresuFunkcji[2].Y = 0.021

	punktyWykresuFunkcji[3].X = 0.03
	punktyWykresuFunkcji[3].Y = 0.032

	punktyWykresuFunkcji[4].X = 0.04
	punktyWykresuFunkcji[4].Y = 0.043

	punktyWykresuFunkcji[5].X = 0.05
	punktyWykresuFunkcji[5].Y = 0.056

	punktyWykresuFunkcji[6].X = 0.06
	punktyWykresuFunkcji[6].Y = 0.068

	punktyWykresuFunkcji[7].X = 0.07
	punktyWykresuFunkcji[7].Y = 0.081

	punktyWykresuFunkcji[8].X = 0.08
	punktyWykresuFunkcji[8].Y = 0.095

	punktyWykresuFunkcji[9].X = 0.09
	punktyWykresuFunkcji[9].Y = 0.109

	punktyWykresuFunkcji[10].X = 0.1
	punktyWykresuFunkcji[10].Y = 0.124

	punktyWykresuFunkcji[11].X = 0.11
	punktyWykresuFunkcji[11].Y = 0.138

	punktyWykresuFunkcji[12].X = 0.12
	punktyWykresuFunkcji[12].Y = 0.154

	punktyWykresuFunkcji[13].X = 0.13
	punktyWykresuFunkcji[13].Y = 0.17

	punktyWykresuFunkcji[14].X = 0.14
	punktyWykresuFunkcji[14].Y = 0.186

	punktyWykresuFunkcji[15].X = 0.15
	punktyWykresuFunkcji[15].Y = 0.202

	punktyWykresuFunkcji[16].X = 0.16
	punktyWykresuFunkcji[16].Y = 0.219

	punktyWykresuFunkcji[17].X = 0.17
	punktyWykresuFunkcji[17].Y = 0.237

	punktyWykresuFunkcji[18].X = 0.18
	punktyWykresuFunkcji[18].Y = 0.255

	punktyWykresuFunkcji[19].X = 0.19
	punktyWykresuFunkcji[19].Y = 0.273

	punktyWykresuFunkcji[20].X = 0.2
	punktyWykresuFunkcji[20].Y = 0.292

	punktyWykresuFunkcji[21].X = 0.21
	punktyWykresuFunkcji[21].Y = 0.311

	punktyWykresuFunkcji[22].X = 0.22
	punktyWykresuFunkcji[22].Y = 0.33

	punktyWykresuFunkcji[23].X = 0.23
	punktyWykresuFunkcji[23].Y = 0.35

	punktyWykresuFunkcji[24].X = 0.24
	punktyWykresuFunkcji[24].Y = 0.37

	punktyWykresuFunkcji[25].X = 0.25
	punktyWykresuFunkcji[25].Y = 0.39

	punktyWykresuFunkcji[26].X = 0.26
	punktyWykresuFunkcji[26].Y = 0.411

	punktyWykresuFunkcji[27].X = 0.27
	punktyWykresuFunkcji[27].Y = 0.432

	punktyWykresuFunkcji[28].X = 0.28
	punktyWykresuFunkcji[28].Y = 0.454

	punktyWykresuFunkcji[29].X = 0.29
	punktyWykresuFunkcji[29].Y = 0.476

	punktyWykresuFunkcji[30].X = 0.3
	punktyWykresuFunkcji[30].Y = 0.498

	punktyWykresuFunkcji[31].X = 0.31
	punktyWykresuFunkcji[31].Y = 0.52

	punktyWykresuFunkcji[32].X = 0.32
	punktyWykresuFunkcji[32].Y = 0.543

	punktyWykresuFunkcji[33].X = 0.33
	punktyWykresuFunkcji[33].Y = 0.566

	punktyWykresuFunkcji[34].X = 0.34
	punktyWykresuFunkcji[34].Y = 0.59

	punktyWykresuFunkcji[35].X = 0.35
	punktyWykresuFunkcji[35].Y = 0.614

	punktyWykresuFunkcji[36].X = 0.36
	punktyWykresuFunkcji[36].Y = 0.638

	punktyWykresuFunkcji[37].X = 0.37
	punktyWykresuFunkcji[37].Y = 0.662

	punktyWykresuFunkcji[38].X = 0.38
	punktyWykresuFunkcji[38].Y = 0.687

	punktyWykresuFunkcji[39].X = 0.39
	punktyWykresuFunkcji[39].Y = 0.711

	punktyWykresuFunkcji[40].X = 0.40
	punktyWykresuFunkcji[40].Y = 0.737

	punktyWykresuFunkcji[41].X = 0.41
	punktyWykresuFunkcji[41].Y = 0.762

	punktyWykresuFunkcji[42].X = 0.42
	punktyWykresuFunkcji[42].Y = 0.788

	punktyWykresuFunkcji[43].X = 0.43
	punktyWykresuFunkcji[43].Y = 0.814

	punktyWykresuFunkcji[44].X = 0.44
	punktyWykresuFunkcji[44].Y = 0.84

	punktyWykresuFunkcji[45].X = 0.45
	punktyWykresuFunkcji[45].Y = 0.866

	punktyWykresuFunkcji[46].X = 0.46
	punktyWykresuFunkcji[46].Y = 0.893

	punktyWykresuFunkcji[47].X = 0.47
	punktyWykresuFunkcji[47].Y = 0.92

	punktyWykresuFunkcji[48].X = 0.48
	punktyWykresuFunkcji[48].Y = 0.947

	punktyWykresuFunkcji[49].X = 0.49
	punktyWykresuFunkcji[49].Y = 0.975

	punktyWykresuFunkcji[50].X = 0.5
	punktyWykresuFunkcji[50].Y = 1.002

	punktyWykresuFunkcji[51].X = 0.51
	punktyWykresuFunkcji[51].Y = 1.03

	punktyWykresuFunkcji[52].X = 0.52
	punktyWykresuFunkcji[52].Y = 1.058

	punktyWykresuFunkcji[53].X = 0.53
	punktyWykresuFunkcji[53].Y = 1.086

	punktyWykresuFunkcji[54].X = 0.54
	punktyWykresuFunkcji[54].Y = 1.115

	punktyWykresuFunkcji[55].X = 0.55
	punktyWykresuFunkcji[55].Y = 1.143

	punktyWykresuFunkcji[56].X = 0.56
	punktyWykresuFunkcji[56].Y = 1.172

	punktyWykresuFunkcji[57].X = 0.57
	punktyWykresuFunkcji[57].Y = 1.201

	punktyWykresuFunkcji[58].X = 0.58
	punktyWykresuFunkcji[58].Y = 1.23

	punktyWykresuFunkcji[59].X = 0.59
	punktyWykresuFunkcji[59].Y = 1.26

	punktyWykresuFunkcji[60].X = 0.6
	punktyWykresuFunkcji[60].Y = 1.289

	punktyWykresuFunkcji[61].X = 0.61
	punktyWykresuFunkcji[61].Y = 1.319

	punktyWykresuFunkcji[62].X = 0.62
	punktyWykresuFunkcji[62].Y = 1.349

	punktyWykresuFunkcji[63].X = 0.63
	punktyWykresuFunkcji[63].Y = 1.379

	punktyWykresuFunkcji[64].X = 0.64
	punktyWykresuFunkcji[64].Y = 1.409

	punktyWykresuFunkcji[65].X = 0.65
	punktyWykresuFunkcji[65].Y = 1.439

	punktyWykresuFunkcji[66].X = 0.66
	punktyWykresuFunkcji[66].Y = 1.47

	punktyWykresuFunkcji[67].X = 0.67
	punktyWykresuFunkcji[67].Y = 1.5

	punktyWykresuFunkcji[68].X = 0.68
	punktyWykresuFunkcji[68].Y = 1.531

	punktyWykresuFunkcji[69].X = 0.69
	punktyWykresuFunkcji[69].Y = 1.561

	punktyWykresuFunkcji[70].X = 0.7
	punktyWykresuFunkcji[70].Y = 1.592

	punktyWykresuFunkcji[71].X = 0.71
	punktyWykresuFunkcji[71].Y = 1.623

	punktyWykresuFunkcji[72].X = 0.72
	punktyWykresuFunkcji[72].Y = 1.654

	punktyWykresuFunkcji[73].X = 0.73
	punktyWykresuFunkcji[73].Y = 1.686

	punktyWykresuFunkcji[74].X = 0.74
	punktyWykresuFunkcji[74].Y = 1.717

	punktyWykresuFunkcji[75].X = 0.75
	punktyWykresuFunkcji[75].Y = 1.748

	punktyWykresuFunkcji[76].X = 0.76
	punktyWykresuFunkcji[76].Y = 1.78

	punktyWykresuFunkcji[77].X = 0.77
	punktyWykresuFunkcji[77].Y = 1.811

	punktyWykresuFunkcji[78].X = 0.78
	punktyWykresuFunkcji[78].Y = 1.843

	punktyWykresuFunkcji[79].X = 0.79
	punktyWykresuFunkcji[79].Y = 1.875

	punktyWykresuFunkcji[80].X = 0.8
	punktyWykresuFunkcji[80].Y = 1.906

	punktyWykresuFunkcji[81].X = 0.81
	punktyWykresuFunkcji[81].Y = 1.938

	punktyWykresuFunkcji[82].X = 0.82
	punktyWykresuFunkcji[82].Y = 1.97

	punktyWykresuFunkcji[83].X = 0.83
	punktyWykresuFunkcji[83].Y = 2.002

	punktyWykresuFunkcji[84].X = 0.84
	punktyWykresuFunkcji[84].Y = 2.034

	punktyWykresuFunkcji[85].X = 0.85
	punktyWykresuFunkcji[85].Y = 2.066

	punktyWykresuFunkcji[86].X = 0.86
	punktyWykresuFunkcji[86].Y = 2.098

	punktyWykresuFunkcji[87].X = 0.87
	punktyWykresuFunkcji[87].Y = 2.13

	punktyWykresuFunkcji[88].X = 0.88
	punktyWykresuFunkcji[88].Y = 2.162

	punktyWykresuFunkcji[89].X = 0.89
	punktyWykresuFunkcji[89].Y = 2.194

	punktyWykresuFunkcji[90].X = 0.9
	punktyWykresuFunkcji[90].Y = 2.227

	punktyWykresuFunkcji[91].X = 0.91
	punktyWykresuFunkcji[91].Y = 2.259

	punktyWykresuFunkcji[92].X = 0.92
	punktyWykresuFunkcji[92].Y = 2.291

	punktyWykresuFunkcji[93].X = 0.93
	punktyWykresuFunkcji[93].Y = 2.323

	punktyWykresuFunkcji[94].X = 0.94
	punktyWykresuFunkcji[94].Y = 2.355

	punktyWykresuFunkcji[95].X = 0.95
	punktyWykresuFunkcji[95].Y = 2.387

	punktyWykresuFunkcji[96].X = 0.96
	punktyWykresuFunkcji[96].Y = 2.42

	punktyWykresuFunkcji[97].X = 0.97
	punktyWykresuFunkcji[97].Y = 2.452

	punktyWykresuFunkcji[98].X = 0.98
	punktyWykresuFunkcji[98].Y = 2.484

	punktyWykresuFunkcji[99].X = 0.99
	punktyWykresuFunkcji[99].Y = 2.516

	punktyWykresuFunkcji[100].X = 1.0
	punktyWykresuFunkcji[100].Y = 2.548

	punktyWykresuFunkcji[101].X = 1.01
	punktyWykresuFunkcji[101].Y = 2.58

	punktyWykresuFunkcji[102].X = 1.02
	punktyWykresuFunkcji[102].Y = 2.612

	punktyWykresuFunkcji[103].X = 1.03
	punktyWykresuFunkcji[103].Y = 2.644

	punktyWykresuFunkcji[104].X = 1.04
	punktyWykresuFunkcji[104].Y = 2.676

	punktyWykresuFunkcji[105].X = 1.05
	punktyWykresuFunkcji[105].Y = 2.708

	punktyWykresuFunkcji[106].X = 1.06
	punktyWykresuFunkcji[106].Y = 2.74

	punktyWykresuFunkcji[107].X = 1.07
	punktyWykresuFunkcji[107].Y = 2.772

	punktyWykresuFunkcji[108].X = 1.08
	punktyWykresuFunkcji[108].Y = 2.803

	punktyWykresuFunkcji[109].X = 1.09
	punktyWykresuFunkcji[109].Y = 2.835

	punktyWykresuFunkcji[110].X = 1.1
	punktyWykresuFunkcji[110].Y = 2.867

	punktyWykresuFunkcji[111].X = 1.11
	punktyWykresuFunkcji[111].Y = 2.898

	punktyWykresuFunkcji[112].X = 1.12
	punktyWykresuFunkcji[112].Y = 2.93

	punktyWykresuFunkcji[113].X = 1.13
	punktyWykresuFunkcji[113].Y = 2.961

	punktyWykresuFunkcji[114].X = 1.14
	punktyWykresuFunkcji[114].Y = 2.992

	punktyWykresuFunkcji[115].X = 1.15
	punktyWykresuFunkcji[115].Y = 3.023

	punktyWykresuFunkcji[116].X = 1.16
	punktyWykresuFunkcji[116].Y = 3.055

	punktyWykresuFunkcji[117].X = 1.17
	punktyWykresuFunkcji[117].Y = 3.086

	punktyWykresuFunkcji[118].X = 1.18
	punktyWykresuFunkcji[118].Y = 3.117

	punktyWykresuFunkcji[119].X = 1.19
	punktyWykresuFunkcji[119].Y = 3.147

	punktyWykresuFunkcji[120].X = 1.2
	punktyWykresuFunkcji[120].Y = 3.178

	punktyWykresuFunkcji[121].X = 1.21
	punktyWykresuFunkcji[121].Y = 3.209

	punktyWykresuFunkcji[122].X = 1.22
	punktyWykresuFunkcji[122].Y = 3.239

	punktyWykresuFunkcji[123].X = 1.23
	punktyWykresuFunkcji[123].Y = 3.27

	punktyWykresuFunkcji[124].X = 1.24
	punktyWykresuFunkcji[124].Y = 3.3

	punktyWykresuFunkcji[125].X = 1.25
	punktyWykresuFunkcji[125].Y = 3.33

	punktyWykresuFunkcji[126].X = 1.26
	punktyWykresuFunkcji[126].Y = 3.36

	punktyWykresuFunkcji[127].X = 1.27
	punktyWykresuFunkcji[127].Y = 3.39

	punktyWykresuFunkcji[128].X = 1.28
	punktyWykresuFunkcji[128].Y = 3.42

	punktyWykresuFunkcji[129].X = 1.29
	punktyWykresuFunkcji[129].Y = 3.45

	punktyWykresuFunkcji[130].X = 1.3
	punktyWykresuFunkcji[130].Y = 3.479

	punktyWykresuFunkcji[131].X = 1.31
	punktyWykresuFunkcji[131].Y = 3.508

	punktyWykresuFunkcji[132].X = 1.32
	punktyWykresuFunkcji[132].Y = 3.538

	punktyWykresuFunkcji[133].X = 1.33
	punktyWykresuFunkcji[133].Y = 3.567

	punktyWykresuFunkcji[134].X = 1.34
	punktyWykresuFunkcji[134].Y = 3.596

	punktyWykresuFunkcji[135].X = 1.35
	punktyWykresuFunkcji[135].Y = 3.625

	punktyWykresuFunkcji[136].X = 1.36
	punktyWykresuFunkcji[136].Y = 3.653

	punktyWykresuFunkcji[137].X = 1.37
	punktyWykresuFunkcji[137].Y = 3.682

	punktyWykresuFunkcji[138].X = 1.38
	punktyWykresuFunkcji[138].Y = 3.71

	punktyWykresuFunkcji[139].X = 1.39
	punktyWykresuFunkcji[139].Y = 3.738

	punktyWykresuFunkcji[140].X = 1.4
	punktyWykresuFunkcji[140].Y = 3.766

	punktyWykresuFunkcji[141].X = 1.41
	punktyWykresuFunkcji[141].Y = 3.794

	punktyWykresuFunkcji[142].X = 1.42
	punktyWykresuFunkcji[142].Y = 3.822

	punktyWykresuFunkcji[143].X = 1.43
	punktyWykresuFunkcji[143].Y = 3.849

	punktyWykresuFunkcji[144].X = 1.44
	punktyWykresuFunkcji[144].Y = 3.876

	punktyWykresuFunkcji[145].X = 1.45
	punktyWykresuFunkcji[145].Y = 3.904

	punktyWykresuFunkcji[146].X = 1.46
	punktyWykresuFunkcji[146].Y = 3.93

	punktyWykresuFunkcji[147].X = 1.47
	punktyWykresuFunkcji[147].Y = 3.957

	punktyWykresuFunkcji[148].X = 1.48
	punktyWykresuFunkcji[148].Y = 3.984

	punktyWykresuFunkcji[149].X = 1.49
	punktyWykresuFunkcji[149].Y = 4.01

	punktyWykresuFunkcji[150].X = 1.5
	punktyWykresuFunkcji[150].Y = 0.0

	punktyWykresuFunkcji[151].X = 1.51
	punktyWykresuFunkcji[151].Y = 3.91

	punktyWykresuFunkcji[152].X = 1.52
	punktyWykresuFunkcji[152].Y = 3.924

	punktyWykresuFunkcji[153].X = 1.53
	punktyWykresuFunkcji[153].Y = 3.938

	punktyWykresuFunkcji[154].X = 1.54
	punktyWykresuFunkcji[154].Y = 3.951

	punktyWykresuFunkcji[155].X = 1.55
	punktyWykresuFunkcji[155].Y = 3.963

	punktyWykresuFunkcji[156].X = 1.56
	punktyWykresuFunkcji[156].Y = 3.975

	punktyWykresuFunkcji[157].X = 1.57
	punktyWykresuFunkcji[157].Y = 3.986

	punktyWykresuFunkcji[158].X = 1.58
	punktyWykresuFunkcji[158].Y = 3.996

	punktyWykresuFunkcji[159].X = 1.59
	punktyWykresuFunkcji[159].Y = 4.006

	punktyWykresuFunkcji[160].X = 1.6
	punktyWykresuFunkcji[160].Y = 4.015

	punktyWykresuFunkcji[161].X = 1.61
	punktyWykresuFunkcji[161].Y = 4.023

	punktyWykresuFunkcji[162].X = 1.62
	punktyWykresuFunkcji[162].Y = 4.031

	punktyWykresuFunkcji[163].X = 1.63
	punktyWykresuFunkcji[163].Y = 4.038

	punktyWykresuFunkcji[164].X = 1.64
	punktyWykresuFunkcji[164].Y = 4.045

	punktyWykresuFunkcji[165].X = 1.65
	punktyWykresuFunkcji[165].Y = 4.05

	punktyWykresuFunkcji[166].X = 1.66
	punktyWykresuFunkcji[166].Y = 4.055

	punktyWykresuFunkcji[167].X = 1.67
	punktyWykresuFunkcji[167].Y = 4.06

	punktyWykresuFunkcji[168].X = 1.68
	punktyWykresuFunkcji[168].Y = 4.063

	punktyWykresuFunkcji[169].X = 1.69
	punktyWykresuFunkcji[169].Y = 4.066

	punktyWykresuFunkcji[170].X = 1.7
	punktyWykresuFunkcji[170].Y = 4.068

	punktyWykresuFunkcji[171].X = 1.71
	punktyWykresuFunkcji[171].Y = 4.07

	punktyWykresuFunkcji[172].X = 1.72
	punktyWykresuFunkcji[172].Y = 4.071

	punktyWykresuFunkcji[173].X = 1.73
	punktyWykresuFunkcji[173].Y = 4.07

	punktyWykresuFunkcji[174].X = 1.74
	punktyWykresuFunkcji[174].Y = 4.07

	punktyWykresuFunkcji[175].X = 1.75
	punktyWykresuFunkcji[175].Y = 4.068

	punktyWykresuFunkcji[176].X = 1.76
	punktyWykresuFunkcji[176].Y = 4.065

	punktyWykresuFunkcji[177].X = 1.77
	punktyWykresuFunkcji[177].Y = 4.062

	punktyWykresuFunkcji[178].X = 1.78
	punktyWykresuFunkcji[178].Y = 4.058

	punktyWykresuFunkcji[179].X = 1.79
	punktyWykresuFunkcji[179].Y = 4.053

	punktyWykresuFunkcji[180].X = 1.8
	punktyWykresuFunkcji[180].Y = 4.048

	punktyWykresuFunkcji[181].X = 1.81
	punktyWykresuFunkcji[181].Y = 4.041

	punktyWykresuFunkcji[182].X = 1.82
	punktyWykresuFunkcji[182].Y = 4.034

	punktyWykresuFunkcji[183].X = 1.83
	punktyWykresuFunkcji[183].Y = 4.025

	punktyWykresuFunkcji[184].X = 1.84
	punktyWykresuFunkcji[184].Y = 4.016

	punktyWykresuFunkcji[185].X = 1.85
	punktyWykresuFunkcji[185].Y = 4.006

	punktyWykresuFunkcji[186].X = 1.86
	punktyWykresuFunkcji[186].Y = 3.996

	punktyWykresuFunkcji[187].X = 1.87
	punktyWykresuFunkcji[187].Y = 3.984

	punktyWykresuFunkcji[188].X = 1.88
	punktyWykresuFunkcji[188].Y = 3.971

	punktyWykresuFunkcji[189].X = 1.89
	punktyWykresuFunkcji[189].Y = 3.958

	punktyWykresuFunkcji[190].X = 1.9
	punktyWykresuFunkcji[190].Y = 3.943

	punktyWykresuFunkcji[191].X = 1.91
	punktyWykresuFunkcji[191].Y = 3.928

	punktyWykresuFunkcji[192].X = 1.92
	punktyWykresuFunkcji[192].Y = 3.912

	punktyWykresuFunkcji[193].X = 1.93
	punktyWykresuFunkcji[193].Y = 3.894

	punktyWykresuFunkcji[194].X = 1.94
	punktyWykresuFunkcji[194].Y = 3.876

	punktyWykresuFunkcji[195].X = 1.95
	punktyWykresuFunkcji[195].Y = 3.857

	punktyWykresuFunkcji[196].X = 1.96
	punktyWykresuFunkcji[196].Y = 3.837

	punktyWykresuFunkcji[197].X = 1.97
	punktyWykresuFunkcji[197].Y = 3.816

	punktyWykresuFunkcji[198].X = 1.98
	punktyWykresuFunkcji[198].Y = 3.794

	punktyWykresuFunkcji[199].X = 1.99
	punktyWykresuFunkcji[199].Y = 3.771

	punktyWykresuFunkcji[200].X = 2.0
	punktyWykresuFunkcji[200].Y = 3.747

	punktyWykresuFunkcji[201].X = 2.01
	punktyWykresuFunkcji[201].Y = 3.722

	punktyWykresuFunkcji[202].X = 2.02
	punktyWykresuFunkcji[202].Y = 3.696

	punktyWykresuFunkcji[203].X = 2.03
	punktyWykresuFunkcji[203].Y = 3.669

	punktyWykresuFunkcji[204].X = 2.04
	punktyWykresuFunkcji[204].Y = 3.641

	punktyWykresuFunkcji[205].X = 2.05
	punktyWykresuFunkcji[205].Y = 3.612

	punktyWykresuFunkcji[206].X = 2.06
	punktyWykresuFunkcji[206].Y = 3.581

	punktyWykresuFunkcji[207].X = 2.07
	punktyWykresuFunkcji[207].Y = 3.55

	punktyWykresuFunkcji[208].X = 2.08
	punktyWykresuFunkcji[208].Y = 3.518

	punktyWykresuFunkcji[209].X = 2.09
	punktyWykresuFunkcji[209].Y = 3.484

	punktyWykresuFunkcji[210].X = 2.1
	punktyWykresuFunkcji[210].Y = 3.45

	punktyWykresuFunkcji[211].X = 2.11
	punktyWykresuFunkcji[211].Y = 3.414

	punktyWykresuFunkcji[212].X = 2.12
	punktyWykresuFunkcji[212].Y = 3.378

	punktyWykresuFunkcji[213].X = 2.13
	punktyWykresuFunkcji[213].Y = 3.34

	punktyWykresuFunkcji[214].X = 2.14
	punktyWykresuFunkcji[214].Y = 3.301

	punktyWykresuFunkcji[215].X = 2.15
	punktyWykresuFunkcji[215].Y = 3.261

	punktyWykresuFunkcji[216].X = 2.16
	punktyWykresuFunkcji[216].Y = 3.22

	punktyWykresuFunkcji[217].X = 2.17
	punktyWykresuFunkcji[217].Y = 3.177

	punktyWykresuFunkcji[218].X = 2.18
	punktyWykresuFunkcji[218].Y = 3.134

	punktyWykresuFunkcji[219].X = 2.19
	punktyWykresuFunkcji[219].Y = 3.089

	punktyWykresuFunkcji[220].X = 2.2
	punktyWykresuFunkcji[220].Y = 3.043

	punktyWykresuFunkcji[221].X = 2.21
	punktyWykresuFunkcji[221].Y = 2.996

	punktyWykresuFunkcji[222].X = 2.22
	punktyWykresuFunkcji[222].Y = 2.948

	punktyWykresuFunkcji[223].X = 2.23
	punktyWykresuFunkcji[223].Y = 2.899

	punktyWykresuFunkcji[224].X = 2.24
	punktyWykresuFunkcji[224].Y = 2.848

	punktyWykresuFunkcji[225].X = 2.25
	punktyWykresuFunkcji[225].Y = 2.797

	punktyWykresuFunkcji[226].X = 2.26
	punktyWykresuFunkcji[226].Y = 2.744

	punktyWykresuFunkcji[227].X = 2.27
	punktyWykresuFunkcji[227].Y = 2.69

	punktyWykresuFunkcji[228].X = 2.28
	punktyWykresuFunkcji[228].Y = 2.634

	punktyWykresuFunkcji[229].X = 2.29
	punktyWykresuFunkcji[229].Y = 2.577

	punktyWykresuFunkcji[230].X = 2.3
	punktyWykresuFunkcji[230].Y = 2.52

	punktyWykresuFunkcji[231].X = 2.31
	punktyWykresuFunkcji[231].Y = 2.46

	punktyWykresuFunkcji[232].X = 2.32
	punktyWykresuFunkcji[232].Y = 2.4

	punktyWykresuFunkcji[233].X = 2.33
	punktyWykresuFunkcji[233].Y = 2.338

	punktyWykresuFunkcji[234].X = 2.34
	punktyWykresuFunkcji[234].Y = 2.275

	punktyWykresuFunkcji[235].X = 2.35
	punktyWykresuFunkcji[235].Y = 2.211

	punktyWykresuFunkcji[236].X = 2.36
	punktyWykresuFunkcji[236].Y = 2.146

	punktyWykresuFunkcji[237].X = 2.37
	punktyWykresuFunkcji[237].Y = 2.079

	punktyWykresuFunkcji[238].X = 2.38
	punktyWykresuFunkcji[238].Y = 2.011

	punktyWykresuFunkcji[239].X = 2.39
	punktyWykresuFunkcji[239].Y = 1.941

	punktyWykresuFunkcji[240].X = 2.4
	punktyWykresuFunkcji[240].Y = 1.871

	punktyWykresuFunkcji[241].X = 2.41
	punktyWykresuFunkcji[241].Y = 1.798

	punktyWykresuFunkcji[242].X = 2.42
	punktyWykresuFunkcji[242].Y = 1.725

	punktyWykresuFunkcji[243].X = 2.43
	punktyWykresuFunkcji[243].Y = 1.65

	punktyWykresuFunkcji[244].X = 2.44
	punktyWykresuFunkcji[244].Y = 1.574

	punktyWykresuFunkcji[245].X = 2.45
	punktyWykresuFunkcji[245].Y = 1.497

	punktyWykresuFunkcji[246].X = 2.46
	punktyWykresuFunkcji[246].Y = 1.418

	punktyWykresuFunkcji[247].X = 2.47
	punktyWykresuFunkcji[247].Y = 1.338

	punktyWykresuFunkcji[248].X = 2.48
	punktyWykresuFunkcji[248].Y = 1.256

	punktyWykresuFunkcji[249].X = 2.49
	punktyWykresuFunkcji[249].Y = 1.173

	punktyWykresuFunkcji[250].X = 2.5
	punktyWykresuFunkcji[250].Y = 1.089

	punktyWykresuFunkcji[251].X = 2.51
	punktyWykresuFunkcji[251].Y = 1.003

	punktyWykresuFunkcji[252].X = 2.52
	punktyWykresuFunkcji[252].Y = 0.916

	punktyWykresuFunkcji[253].X = 2.53
	punktyWykresuFunkcji[253].Y = 0.827

	punktyWykresuFunkcji[254].X = 2.54
	punktyWykresuFunkcji[254].Y = 0.737

	punktyWykresuFunkcji[255].X = 2.55
	punktyWykresuFunkcji[255].Y = 0.646

	punktyWykresuFunkcji[256].X = 2.56
	punktyWykresuFunkcji[256].Y = 0.553

	punktyWykresuFunkcji[257].X = 2.57
	punktyWykresuFunkcji[257].Y = 0.458

	punktyWykresuFunkcji[258].X = 2.58
	punktyWykresuFunkcji[258].Y = 0.363

	punktyWykresuFunkcji[259].X = 2.59
	punktyWykresuFunkcji[259].Y = 0.265

	punktyWykresuFunkcji[260].X = 2.6
	punktyWykresuFunkcji[260].Y = 0.167

	punktyWykresuFunkcji[261].X = 2.61
	punktyWykresuFunkcji[261].Y = 0.067

	punktyWykresuFunkcji[262].X = 2.62
	punktyWykresuFunkcji[262].Y = -0.034

	punktyWykresuFunkcji[263].X = 2.63
	punktyWykresuFunkcji[263].Y = -0.137

	punktyWykresuFunkcji[264].X = 2.64
	punktyWykresuFunkcji[264].Y = -0.242

	punktyWykresuFunkcji[265].X = 2.65
	punktyWykresuFunkcji[265].Y = -0.348

	punktyWykresuFunkcji[266].X = 2.66
	punktyWykresuFunkcji[266].Y = -0.456

	punktyWykresuFunkcji[267].X = 2.67
	punktyWykresuFunkcji[267].Y = -0.565

	punktyWykresuFunkcji[268].X = 2.68
	punktyWykresuFunkcji[268].Y = -0.675

	punktyWykresuFunkcji[269].X = 2.69
	punktyWykresuFunkcji[269].Y = -0.788

	punktyWykresuFunkcji[270].X = 2.7
	punktyWykresuFunkcji[270].Y = -0.901

	punktyWykresuFunkcji[271].X = 2.71
	punktyWykresuFunkcji[271].Y = -1.016

	punktyWykresuFunkcji[272].X = 2.72
	punktyWykresuFunkcji[272].Y = -1.133

	punktyWykresuFunkcji[273].X = 2.73
	punktyWykresuFunkcji[273].Y = -1.251

	punktyWykresuFunkcji[274].X = 2.74
	punktyWykresuFunkcji[274].Y = -1.371

	punktyWykresuFunkcji[275].X = 2.75
	punktyWykresuFunkcji[275].Y = -1.493

	punktyWykresuFunkcji[276].X = 2.76
	punktyWykresuFunkcji[276].Y = -1.616

	punktyWykresuFunkcji[277].X = 2.77
	punktyWykresuFunkcji[277].Y = -1.74

	punktyWykresuFunkcji[278].X = 2.78
	punktyWykresuFunkcji[278].Y = -1.866

	punktyWykresuFunkcji[279].X = 2.79
	punktyWykresuFunkcji[279].Y = -1.994

	punktyWykresuFunkcji[280].X = 2.8
	punktyWykresuFunkcji[280].Y = -2.123

	punktyWykresuFunkcji[281].X = 2.81
	punktyWykresuFunkcji[281].Y = -2.254

	punktyWykresuFunkcji[282].X = 2.82
	punktyWykresuFunkcji[282].Y = -2.386

	punktyWykresuFunkcji[283].X = 2.83
	punktyWykresuFunkcji[283].Y = -2.52

	punktyWykresuFunkcji[284].X = 2.84
	punktyWykresuFunkcji[284].Y = -2.656

	punktyWykresuFunkcji[285].X = 2.85
	punktyWykresuFunkcji[285].Y = -2.793

	punktyWykresuFunkcji[286].X = 2.86
	punktyWykresuFunkcji[286].Y = -2.932

	punktyWykresuFunkcji[287].X = 2.87
	punktyWykresuFunkcji[287].Y = -3.073

	punktyWykresuFunkcji[288].X = 2.88
	punktyWykresuFunkcji[288].Y = -3.215

	punktyWykresuFunkcji[289].X = 2.89
	punktyWykresuFunkcji[289].Y = -3.359

	punktyWykresuFunkcji[290].X = 2.9
	punktyWykresuFunkcji[290].Y = -3.504

	punktyWykresuFunkcji[291].X = 2.91
	punktyWykresuFunkcji[291].Y = -3.651

	punktyWykresuFunkcji[292].X = 2.92
	punktyWykresuFunkcji[292].Y = -3.8

	punktyWykresuFunkcji[293].X = 2.93
	punktyWykresuFunkcji[293].Y = -3.95

	punktyWykresuFunkcji[294].X = 2.94
	punktyWykresuFunkcji[294].Y = -4.102

	punktyWykresuFunkcji[295].X = 2.95
	punktyWykresuFunkcji[295].Y = -4.256

	punktyWykresuFunkcji[296].X = 2.96
	punktyWykresuFunkcji[296].Y = -4.411

	punktyWykresuFunkcji[297].X = 2.97
	punktyWykresuFunkcji[297].Y = -4.568

	punktyWykresuFunkcji[298].X = 2.98
	punktyWykresuFunkcji[298].Y = -4.727

	punktyWykresuFunkcji[299].X = 2.99
	punktyWykresuFunkcji[299].Y = -4.887

	punktyWykresuFunkcji[300].X = 3.0
	punktyWykresuFunkcji[300].Y = -5.05

	punktyWykresuFunkcji[301].X = 3.01
	punktyWykresuFunkcji[301].Y = -5.213

	punktyWykresuFunkcji[302].X = 3.02
	punktyWykresuFunkcji[302].Y = -5.379

	punktyWykresuFunkcji[303].X = 3.03
	punktyWykresuFunkcji[303].Y = -5.546

	punktyWykresuFunkcji[304].X = 3.04
	punktyWykresuFunkcji[304].Y = -5.715

	punktyWykresuFunkcji[305].X = 3.05
	punktyWykresuFunkcji[305].Y = -5.886

	punktyWykresuFunkcji[306].X = 3.06
	punktyWykresuFunkcji[306].Y = -6.058

	punktyWykresuFunkcji[307].X = 3.07
	punktyWykresuFunkcji[307].Y = -6.232

	punktyWykresuFunkcji[308].X = 3.08
	punktyWykresuFunkcji[308].Y = -6.408

	punktyWykresuFunkcji[309].X = 3.09
	punktyWykresuFunkcji[309].Y = -6.585

	punktyWykresuFunkcji[310].X = 3.1
	punktyWykresuFunkcji[310].Y = -6.765

	punktyWykresuFunkcji[311].X = 3.11
	punktyWykresuFunkcji[311].Y = -6.946

	punktyWykresuFunkcji[312].X = 3.12
	punktyWykresuFunkcji[312].Y = -7.128

	punktyWykresuFunkcji[313].X = 3.13
	punktyWykresuFunkcji[313].Y = -7.313

	punktyWykresuFunkcji[314].X = 3.14
	punktyWykresuFunkcji[314].Y = -7.499

	punktyWykresuFunkcji[315].X = 3.15
	punktyWykresuFunkcji[315].Y = -7.687

	punktyWykresuFunkcji[316].X = 3.16
	punktyWykresuFunkcji[316].Y = -7.887

	punktyWykresuFunkcji[317].X = 3.17
	punktyWykresuFunkcji[317].Y = -8.068

	punktyWykresuFunkcji[318].X = 3.18
	punktyWykresuFunkcji[318].Y = -8.262

	punktyWykresuFunkcji[319].X = 3.19
	punktyWykresuFunkcji[319].Y = -8.457

	punktyWykresuFunkcji[320].X = 3.2
	punktyWykresuFunkcji[320].Y = -8.653










	wykresFunkcji := plot.New()

	wykresFunkcji.Title.Text = "Wykres funkcji cos(x)"

	wykresFunkcji.X.Label.Text = "x"
	wykresFunkcji.Y.Label.Text = "y"

	liniaWykresu, err := plotter.NewLine(punktyWykresuFunkcji)

	if err != nil {
		panic(err)
	}

	liniaWykresu.LineStyle.Width = vg.Points(0.1)
	liniaWykresu.Color = color.RGBA{R: 200, G: 100, B: 100}

	wykresFunkcji.Add(liniaWykresu)
	wykresFunkcji.Legend.Add("cos(x)", liniaWykresu)

	if err := wykresFunkcji.Save(10*vg.Inch, 10*vg.Inch,
		"Wykresy-funkcji-01.png"); err != nil {

		panic(err)
	}
}
