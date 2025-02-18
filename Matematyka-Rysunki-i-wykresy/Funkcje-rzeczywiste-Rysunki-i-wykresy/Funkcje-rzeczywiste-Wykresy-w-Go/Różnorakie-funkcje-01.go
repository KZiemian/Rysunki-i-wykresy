package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres funkcji f(x) = 0.25x^4 - x^3 + x - 5cos(x) + 5
	// Jest to funkcja pierwotna funkcji g(x) = x^3 - 3x^2 + 1 +
	// 5sin(x), spełniająca warunek f(0) = 0.

	punktyWykresuFunkcji := make(plotter.XYs, 841)

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
	punktyWykresuFunkcji[150].Y = 4.036

	punktyWykresuFunkcji[151].X = 1.51
	punktyWykresuFunkcji[151].Y = 4.063

	punktyWykresuFunkcji[152].X = 1.52
	punktyWykresuFunkcji[152].Y = 4.088

	punktyWykresuFunkcji[153].X = 1.53
	punktyWykresuFunkcji[153].Y = 4.114

	punktyWykresuFunkcji[154].X = 1.54
	punktyWykresuFunkcji[154].Y = 4.139

	punktyWykresuFunkcji[155].X = 1.55
	punktyWykresuFunkcji[155].Y = 4.165

	punktyWykresuFunkcji[156].X = 1.56
	punktyWykresuFunkcji[156].Y = 4.19

	punktyWykresuFunkcji[157].X = 1.57
	punktyWykresuFunkcji[157].Y = 4.215

	punktyWykresuFunkcji[158].X = 1.58
	punktyWykresuFunkcji[158].Y = 4.239

	punktyWykresuFunkcji[159].X = 1.59
	punktyWykresuFunkcji[159].Y = 4.264

	punktyWykresuFunkcji[160].X = 1.6
	punktyWykresuFunkcji[160].Y = 4.288

	punktyWykresuFunkcji[161].X = 1.61
	punktyWykresuFunkcji[161].Y = 4.312

	punktyWykresuFunkcji[162].X = 1.62
	punktyWykresuFunkcji[162].Y = 4.336

	punktyWykresuFunkcji[163].X = 1.63
	punktyWykresuFunkcji[163].Y = 4.359

	punktyWykresuFunkcji[164].X = 1.64
	punktyWykresuFunkcji[164].Y = 4.383

	punktyWykresuFunkcji[165].X = 1.65
	punktyWykresuFunkcji[165].Y = 4.406

	punktyWykresuFunkcji[166].X = 1.66
	punktyWykresuFunkcji[166].Y = 4.429

	punktyWykresuFunkcji[167].X = 1.67
	punktyWykresuFunkcji[167].Y = 4.452

	punktyWykresuFunkcji[168].X = 1.68
	punktyWykresuFunkcji[168].Y = 4.474

	punktyWykresuFunkcji[169].X = 1.69
	punktyWykresuFunkcji[169].Y = 4.497

	punktyWykresuFunkcji[170].X = 1.7
	punktyWykresuFunkcji[170].Y = 4.519

	punktyWykresuFunkcji[171].X = 1.71
	punktyWykresuFunkcji[171].Y = 4.541

	punktyWykresuFunkcji[172].X = 1.72
	punktyWykresuFunkcji[172].Y = 4.562

	punktyWykresuFunkcji[173].X = 1.73
	punktyWykresuFunkcji[173].Y = 4.584

	punktyWykresuFunkcji[174].X = 1.74
	punktyWykresuFunkcji[174].Y = 4.605

	punktyWykresuFunkcji[175].X = 1.75
	punktyWykresuFunkcji[175].Y = 4.626

	punktyWykresuFunkcji[176].X = 1.76
	punktyWykresuFunkcji[176].Y = 4.647

	punktyWykresuFunkcji[177].X = 1.77
	punktyWykresuFunkcji[177].Y = 4.668

	punktyWykresuFunkcji[178].X = 1.78
	punktyWykresuFunkcji[178].Y = 4.688

	punktyWykresuFunkcji[179].X = 1.79
	punktyWykresuFunkcji[179].Y = 4.708

	punktyWykresuFunkcji[180].X = 1.8
	punktyWykresuFunkcji[180].Y = 4.728

	punktyWykresuFunkcji[181].X = 1.81
	punktyWykresuFunkcji[181].Y = 4.748

	punktyWykresuFunkcji[182].X = 1.82
	punktyWykresuFunkcji[182].Y = 4.767

	punktyWykresuFunkcji[183].X = 1.83
	punktyWykresuFunkcji[183].Y = 4.786

	punktyWykresuFunkcji[184].X = 1.84
	punktyWykresuFunkcji[184].Y = 4.805

	punktyWykresuFunkcji[185].X = 1.85
	punktyWykresuFunkcji[185].Y = 4.824

	punktyWykresuFunkcji[186].X = 1.86
	punktyWykresuFunkcji[186].Y = 4.843

	punktyWykresuFunkcji[187].X = 1.87
	punktyWykresuFunkcji[187].Y = 4.861

	punktyWykresuFunkcji[188].X = 1.88
	punktyWykresuFunkcji[188].Y = 4.879

	punktyWykresuFunkcji[189].X = 1.89
	punktyWykresuFunkcji[189].Y = 4.897

	punktyWykresuFunkcji[190].X = 1.9
	punktyWykresuFunkcji[190].Y = 4.915

	punktyWykresuFunkcji[191].X = 1.91
	punktyWykresuFunkcji[191].Y = 4.933

	punktyWykresuFunkcji[192].X = 1.92
	punktyWykresuFunkcji[192].Y = 4.95

	punktyWykresuFunkcji[193].X = 1.93
	punktyWykresuFunkcji[193].Y = 4.967

	punktyWykresuFunkcji[194].X = 1.94
	punktyWykresuFunkcji[194].Y = 4.984

	punktyWykresuFunkcji[195].X = 1.95
	punktyWykresuFunkcji[195].Y = 5.0

	punktyWykresuFunkcji[196].X = 1.96
	punktyWykresuFunkcji[196].Y = 5.017

	punktyWykresuFunkcji[197].X = 1.97
	punktyWykresuFunkcji[197].Y = 5.033

	punktyWykresuFunkcji[198].X = 1.98
	punktyWykresuFunkcji[198].Y = 5.049

	punktyWykresuFunkcji[199].X = 1.99
	punktyWykresuFunkcji[199].Y = 5.065

	punktyWykresuFunkcji[200].X = 2.0
	punktyWykresuFunkcji[200].Y = 5.08

	punktyWykresuFunkcji[201].X = 2.01
	punktyWykresuFunkcji[201].Y = 5.096

	punktyWykresuFunkcji[202].X = 2.02
	punktyWykresuFunkcji[202].Y = 5.111

	punktyWykresuFunkcji[203].X = 2.03
	punktyWykresuFunkcji[203].Y = 5.126

	punktyWykresuFunkcji[204].X = 2.04
	punktyWykresuFunkcji[204].Y = 5.14

	punktyWykresuFunkcji[205].X = 2.05
	punktyWykresuFunkcji[205].Y = 5.155

	punktyWykresuFunkcji[206].X = 2.06
	punktyWykresuFunkcji[206].Y = 5.169

	punktyWykresuFunkcji[207].X = 2.07
	punktyWykresuFunkcji[207].Y = 5.184

	punktyWykresuFunkcji[208].X = 2.08
	punktyWykresuFunkcji[208].Y = 5.197

	punktyWykresuFunkcji[209].X = 2.09
	punktyWykresuFunkcji[209].Y = 5.211

	punktyWykresuFunkcji[210].X = 2.1
	punktyWykresuFunkcji[210].Y = 5.225

	punktyWykresuFunkcji[211].X = 2.11
	punktyWykresuFunkcji[211].Y = 5.238

	punktyWykresuFunkcji[212].X = 2.12
	punktyWykresuFunkcji[212].Y = 5.251

	punktyWykresuFunkcji[213].X = 2.13
	punktyWykresuFunkcji[213].Y = 5.264

	punktyWykresuFunkcji[214].X = 2.14
	punktyWykresuFunkcji[214].Y = 5.277

	punktyWykresuFunkcji[215].X = 2.15
	punktyWykresuFunkcji[215].Y = 5.29

	punktyWykresuFunkcji[216].X = 2.16
	punktyWykresuFunkcji[216].Y = 5.302

	punktyWykresuFunkcji[217].X = 2.17
	punktyWykresuFunkcji[217].Y = 5.315

	punktyWykresuFunkcji[218].X = 2.18
	punktyWykresuFunkcji[218].Y = 5.327

	punktyWykresuFunkcji[219].X = 2.19
	punktyWykresuFunkcji[219].Y = 5.339

	punktyWykresuFunkcji[220].X = 2.2
	punktyWykresuFunkcji[220].Y = 5.35

	punktyWykresuFunkcji[221].X = 2.21
	punktyWykresuFunkcji[221].Y = 5.362

	punktyWykresuFunkcji[222].X = 2.22
	punktyWykresuFunkcji[222].Y = 5.374

	punktyWykresuFunkcji[223].X = 2.23
	punktyWykresuFunkcji[223].Y = 5.385

	punktyWykresuFunkcji[224].X = 2.24
	punktyWykresuFunkcji[224].Y = 5.396

	punktyWykresuFunkcji[225].X = 2.25
	punktyWykresuFunkcji[225].Y = 5.407

	punktyWykresuFunkcji[226].X = 2.26
	punktyWykresuFunkcji[226].Y = 5.418

	punktyWykresuFunkcji[227].X = 2.27
	punktyWykresuFunkcji[227].Y = 5.429

	punktyWykresuFunkcji[228].X = 2.28
	punktyWykresuFunkcji[228].Y = 5.439

	punktyWykresuFunkcji[229].X = 2.29
	punktyWykresuFunkcji[229].Y = 5.45

	punktyWykresuFunkcji[230].X = 2.3
	punktyWykresuFunkcji[230].Y = 5.46

	punktyWykresuFunkcji[231].X = 2.31
	punktyWykresuFunkcji[231].Y = 5.47

	punktyWykresuFunkcji[232].X = 2.32
	punktyWykresuFunkcji[232].Y = 5.48

	punktyWykresuFunkcji[233].X = 2.33
	punktyWykresuFunkcji[233].Y = 5.49

	punktyWykresuFunkcji[234].X = 2.34
	punktyWykresuFunkcji[234].Y = 5.5

	punktyWykresuFunkcji[235].X = 2.35
	punktyWykresuFunkcji[235].Y = 5.51

	punktyWykresuFunkcji[236].X = 2.36
	punktyWykresuFunkcji[236].Y = 5.519

	punktyWykresuFunkcji[237].X = 2.37
	punktyWykresuFunkcji[237].Y = 5.529

	punktyWykresuFunkcji[238].X = 2.38
	punktyWykresuFunkcji[238].Y = 5.538

	punktyWykresuFunkcji[239].X = 2.39
	punktyWykresuFunkcji[239].Y = 5.548

	punktyWykresuFunkcji[240].X = 2.4
	punktyWykresuFunkcji[240].Y = 5.557

	punktyWykresuFunkcji[241].X = 2.41
	punktyWykresuFunkcji[241].Y = 5.566

	punktyWykresuFunkcji[242].X = 2.42
	punktyWykresuFunkcji[242].Y = 5.575

	punktyWykresuFunkcji[243].X = 2.43
	punktyWykresuFunkcji[243].Y = 5.584

	punktyWykresuFunkcji[244].X = 2.44
	punktyWykresuFunkcji[244].Y = 5.593

	punktyWykresuFunkcji[245].X = 2.45
	punktyWykresuFunkcji[245].Y = 5.602

	punktyWykresuFunkcji[246].X = 2.46
	punktyWykresuFunkcji[246].Y = 5.611

	punktyWykresuFunkcji[247].X = 2.47
	punktyWykresuFunkcji[247].Y = 5.62

	punktyWykresuFunkcji[248].X = 2.48
	punktyWykresuFunkcji[248].Y = 5.628

	punktyWykresuFunkcji[249].X = 2.49
	punktyWykresuFunkcji[249].Y = 5.637

	punktyWykresuFunkcji[250].X = 2.5
	punktyWykresuFunkcji[250].Y = 5.646

	punktyWykresuFunkcji[251].X = 2.51
	punktyWykresuFunkcji[251].Y = 5.655

	punktyWykresuFunkcji[252].X = 2.52
	punktyWykresuFunkcji[252].Y = 5.663

	punktyWykresuFunkcji[253].X = 2.53
	punktyWykresuFunkcji[253].Y = 5.672

	punktyWykresuFunkcji[254].X = 2.54
	punktyWykresuFunkcji[254].Y = 5.68

	punktyWykresuFunkcji[255].X = 2.55
	punktyWykresuFunkcji[255].Y = 5.689

	punktyWykresuFunkcji[256].X = 2.56
	punktyWykresuFunkcji[256].Y = 5.698

	punktyWykresuFunkcji[257].X = 2.57
	punktyWykresuFunkcji[257].Y = 5.706

	punktyWykresuFunkcji[258].X = 2.58
	punktyWykresuFunkcji[258].Y = 5.715

	punktyWykresuFunkcji[259].X = 2.59
	punktyWykresuFunkcji[259].Y = 5.724

	punktyWykresuFunkcji[260].X = 2.6
	punktyWykresuFunkcji[260].Y = 5.732

	punktyWykresuFunkcji[261].X = 2.61
	punktyWykresuFunkcji[261].Y = 5.741

	punktyWykresuFunkcji[262].X = 2.62
	punktyWykresuFunkcji[262].Y = 5.75

	punktyWykresuFunkcji[263].X = 2.63
	punktyWykresuFunkcji[263].Y = 5.759

	punktyWykresuFunkcji[264].X = 2.64
	punktyWykresuFunkcji[264].Y = 5.768

	punktyWykresuFunkcji[265].X = 2.65
	punktyWykresuFunkcji[265].Y = 5.777

	punktyWykresuFunkcji[266].X = 2.66
	punktyWykresuFunkcji[266].Y = 5.786

	punktyWykresuFunkcji[267].X = 2.67
	punktyWykresuFunkcji[267].Y = 5.795

	punktyWykresuFunkcji[268].X = 2.68
	punktyWykresuFunkcji[268].Y = 5.804

	punktyWykresuFunkcji[269].X = 2.69
	punktyWykresuFunkcji[269].Y = 5.813

	punktyWykresuFunkcji[270].X = 2.7
	punktyWykresuFunkcji[270].Y = 5.823

	punktyWykresuFunkcji[271].X = 2.71
	punktyWykresuFunkcji[271].Y = 5.832

	punktyWykresuFunkcji[272].X = 2.72
	punktyWykresuFunkcji[272].Y = 5.842

	punktyWykresuFunkcji[273].X = 2.73
	punktyWykresuFunkcji[273].Y = 5.852

	punktyWykresuFunkcji[274].X = 2.74
	punktyWykresuFunkcji[274].Y = 5.862

	punktyWykresuFunkcji[275].X = 2.75
	punktyWykresuFunkcji[275].Y = 5.872

	punktyWykresuFunkcji[276].X = 2.76
	punktyWykresuFunkcji[276].Y = 5.882

	punktyWykresuFunkcji[277].X = 2.77
	punktyWykresuFunkcji[277].Y = 5.893

	punktyWykresuFunkcji[278].X = 2.78
	punktyWykresuFunkcji[278].Y = 5.903

	punktyWykresuFunkcji[279].X = 2.79
	punktyWykresuFunkcji[279].Y = 5.914

	punktyWykresuFunkcji[280].X = 2.8
	punktyWykresuFunkcji[280].Y = 5.925

	punktyWykresuFunkcji[281].X = 2.81
	punktyWykresuFunkcji[281].Y = 5.936

	punktyWykresuFunkcji[282].X = 2.82
	punktyWykresuFunkcji[282].Y = 5.948

	punktyWykresuFunkcji[283].X = 2.83
	punktyWykresuFunkcji[283].Y = 5.959

	punktyWykresuFunkcji[284].X = 2.84
	punktyWykresuFunkcji[284].Y = 5.971

	punktyWykresuFunkcji[285].X = 2.85
	punktyWykresuFunkcji[285].Y = 5.983

	punktyWykresuFunkcji[286].X = 2.86
	punktyWykresuFunkcji[286].Y = 5.995

	punktyWykresuFunkcji[287].X = 2.87
	punktyWykresuFunkcji[287].Y = 6.008

	punktyWykresuFunkcji[288].X = 2.88
	punktyWykresuFunkcji[288].Y = 6.021

	punktyWykresuFunkcji[289].X = 2.89
	punktyWykresuFunkcji[289].Y = 6.034

	punktyWykresuFunkcji[290].X = 2.9
	punktyWykresuFunkcji[290].Y = 6.047

	punktyWykresuFunkcji[291].X = 2.91
	punktyWykresuFunkcji[291].Y = 6.061

	punktyWykresuFunkcji[292].X = 2.92
	punktyWykresuFunkcji[292].Y = 6.075

	punktyWykresuFunkcji[293].X = 2.93
	punktyWykresuFunkcji[293].Y = 6.089

	punktyWykresuFunkcji[294].X = 2.94
	punktyWykresuFunkcji[294].Y = 6.104

	punktyWykresuFunkcji[295].X = 2.95
	punktyWykresuFunkcji[295].Y = 6.119

	punktyWykresuFunkcji[296].X = 2.96
	punktyWykresuFunkcji[296].Y = 6.134

	punktyWykresuFunkcji[297].X = 2.97
	punktyWykresuFunkcji[297].Y = 6.15

	punktyWykresuFunkcji[298].X = 2.98
	punktyWykresuFunkcji[298].Y = 6.166

	punktyWykresuFunkcji[299].X = 2.99
	punktyWykresuFunkcji[299].Y = 6.183

	punktyWykresuFunkcji[300].X = 3.0
	punktyWykresuFunkcji[300].Y = 6.2

	punktyWykresuFunkcji[301].X = 3.01
	punktyWykresuFunkcji[301].Y = 6.217

	punktyWykresuFunkcji[302].X = 3.02
	punktyWykresuFunkcji[302].Y = 6.234

	punktyWykresuFunkcji[303].X = 3.03
	punktyWykresuFunkcji[303].Y = 6.253

	punktyWykresuFunkcji[304].X = 3.04
	punktyWykresuFunkcji[304].Y = 6.271

	punktyWykresuFunkcji[305].X = 3.05
	punktyWykresuFunkcji[305].Y = 6.29

	punktyWykresuFunkcji[306].X = 3.06
	punktyWykresuFunkcji[306].Y = 6.31

	punktyWykresuFunkcji[307].X = 3.07
	punktyWykresuFunkcji[307].Y = 6.329

	punktyWykresuFunkcji[308].X = 3.08
	punktyWykresuFunkcji[308].Y = 6.35

	punktyWykresuFunkcji[309].X = 3.09
	punktyWykresuFunkcji[309].Y = 6.371

	punktyWykresuFunkcji[310].X = 3.1
	punktyWykresuFunkcji[310].Y = 6.392

	punktyWykresuFunkcji[311].X = 3.11
	punktyWykresuFunkcji[311].Y = 6.414

	punktyWykresuFunkcji[312].X = 3.12
	punktyWykresuFunkcji[312].Y = 6.437

	punktyWykresuFunkcji[313].X = 3.13
	punktyWykresuFunkcji[313].Y = 6.46

	punktyWykresuFunkcji[314].X = 3.14
	punktyWykresuFunkcji[314].Y = 6.483

	punktyWykresuFunkcji[315].X = 3.15
	punktyWykresuFunkcji[315].Y = 6.507

	punktyWykresuFunkcji[316].X = 3.16
	punktyWykresuFunkcji[316].Y = 6.532

	punktyWykresuFunkcji[317].X = 3.17
	punktyWykresuFunkcji[317].Y = 6.558

	punktyWykresuFunkcji[318].X = 3.18
	punktyWykresuFunkcji[318].Y = 6.584

	punktyWykresuFunkcji[319].X = 3.19
	punktyWykresuFunkcji[319].Y = 6.61

	punktyWykresuFunkcji[320].X = 3.2
	punktyWykresuFunkcji[320].Y = 6.637

	punktyWykresuFunkcji[321].X = 3.21
	punktyWykresuFunkcji[321].Y = 6.665

	punktyWykresuFunkcji[322].X = 3.22
	punktyWykresuFunkcji[322].Y = 6.694

	punktyWykresuFunkcji[323].X = 3.23
	punktyWykresuFunkcji[323].Y = 6.723

	punktyWykresuFunkcji[324].X = 3.24
	punktyWykresuFunkcji[324].Y = 6.753

	punktyWykresuFunkcji[325].X = 3.25
	punktyWykresuFunkcji[325].Y = 6.784

	punktyWykresuFunkcji[326].X = 3.26
	punktyWykresuFunkcji[326].Y = 6.815

	punktyWykresuFunkcji[327].X = 3.27
	punktyWykresuFunkcji[327].Y = 6.847

	punktyWykresuFunkcji[328].X = 3.28
	punktyWykresuFunkcji[328].Y = 6.88

	punktyWykresuFunkcji[329].X = 3.29
	punktyWykresuFunkcji[329].Y = 6.914

	punktyWykresuFunkcji[330].X = 3.3
	punktyWykresuFunkcji[330].Y = 6.948

	punktyWykresuFunkcji[331].X = 3.31
	punktyWykresuFunkcji[331].Y = 6.983

	punktyWykresuFunkcji[332].X = 3.32
	punktyWykresuFunkcji[332].Y = 7.019

	punktyWykresuFunkcji[333].X = 3.33
	punktyWykresuFunkcji[333].Y = 7.056

	punktyWykresuFunkcji[334].X = 3.34
	punktyWykresuFunkcji[334].Y = 7.094

	punktyWykresuFunkcji[335].X = 3.35
	punktyWykresuFunkcji[335].Y = 7.132

	punktyWykresuFunkcji[336].X = 3.36
	punktyWykresuFunkcji[336].Y = 7.171

	punktyWykresuFunkcji[337].X = 3.37
	punktyWykresuFunkcji[337].Y = 7.212

	punktyWykresuFunkcji[338].X = 3.38
	punktyWykresuFunkcji[338].Y = 7.253

	punktyWykresuFunkcji[339].X = 3.39
	punktyWykresuFunkcji[339].Y = 7.295

	punktyWykresuFunkcji[340].X = 3.4
	punktyWykresuFunkcji[340].Y = 7.338

	punktyWykresuFunkcji[341].X = 3.41
	punktyWykresuFunkcji[341].Y = 7.382

	punktyWykresuFunkcji[342].X = 3.42
	punktyWykresuFunkcji[342].Y = 7.427

	punktyWykresuFunkcji[343].X = 3.43
	punktyWykresuFunkcji[343].Y = 7.473

	punktyWykresuFunkcji[344].X = 3.44
	punktyWykresuFunkcji[344].Y = 7.52

	punktyWykresuFunkcji[345].X = 3.45
	punktyWykresuFunkcji[345].Y = 7.567

	punktyWykresuFunkcji[346].X = 3.46
	punktyWykresuFunkcji[346].Y = 7.616

	punktyWykresuFunkcji[347].X = 3.47
	punktyWykresuFunkcji[347].Y = 7.666

	punktyWykresuFunkcji[348].X = 3.48
	punktyWykresuFunkcji[348].Y = 7.717

	punktyWykresuFunkcji[349].X = 3.49
	punktyWykresuFunkcji[349].Y = 7.769

	punktyWykresuFunkcji[350].X = 3.5
	punktyWykresuFunkcji[350].Y = 7.822

	punktyWykresuFunkcji[351].X = 3.51
	punktyWykresuFunkcji[351].Y = 7.877

	punktyWykresuFunkcji[352].X = 3.52
	punktyWykresuFunkcji[352].Y = 7.932

	punktyWykresuFunkcji[353].X = 3.53
	punktyWykresuFunkcji[353].Y = 7.989

	punktyWykresuFunkcji[354].X = 3.54
	punktyWykresuFunkcji[354].Y = 8.046

	punktyWykresuFunkcji[355].X = 3.55
	punktyWykresuFunkcji[355].Y = 8.105

	punktyWykresuFunkcji[356].X = 3.56
	punktyWykresuFunkcji[356].Y = 8.165

	punktyWykresuFunkcji[357].X = 3.57
	punktyWykresuFunkcji[357].Y = 8.227

	punktyWykresuFunkcji[358].X = 3.58
	punktyWykresuFunkcji[358].Y = 8.289

	punktyWykresuFunkcji[359].X = 3.59
	punktyWykresuFunkcji[359].Y = 8.353

	punktyWykresuFunkcji[360].X = 3.6
	punktyWykresuFunkcji[360].Y = 8.418

	punktyWykresuFunkcji[361].X = 3.61
	punktyWykresuFunkcji[361].Y = 8.484

	punktyWykresuFunkcji[362].X = 3.62
	punktyWykresuFunkcji[362].Y = 8.552

	punktyWykresuFunkcji[363].X = 3.63
	punktyWykresuFunkcji[363].Y = 8.62

	punktyWykresuFunkcji[364].X = 3.64
	punktyWykresuFunkcji[364].Y = 8.691

	punktyWykresuFunkcji[365].X = 3.65
	punktyWykresuFunkcji[365].Y = 8.762

	punktyWykresuFunkcji[366].X = 3.66
	punktyWykresuFunkcji[366].Y = 8.835

	punktyWykresuFunkcji[367].X = 3.67
	punktyWykresuFunkcji[367].Y = 8.91

	punktyWykresuFunkcji[368].X = 3.68
	punktyWykresuFunkcji[368].Y = 8.985

	punktyWykresuFunkcji[369].X = 3.69
	punktyWykresuFunkcji[369].Y = 9.062

	punktyWykresuFunkcji[370].X = 3.7
	punktyWykresuFunkcji[370].Y = 9.141

	punktyWykresuFunkcji[371].X = 3.71
	punktyWykresuFunkcji[371].Y = 9.221

	punktyWykresuFunkcji[372].X = 3.72
	punktyWykresuFunkcji[372].Y = 9.303

	punktyWykresuFunkcji[373].X = 3.73
	punktyWykresuFunkcji[373].Y = 9.386

	punktyWykresuFunkcji[374].X = 3.74
	punktyWykresuFunkcji[374].Y = 9.47

	punktyWykresuFunkcji[375].X = 3.75
	punktyWykresuFunkcji[375].Y = 9.556

	punktyWykresuFunkcji[376].X = 3.76
	punktyWykresuFunkcji[376].Y = 9.644

	punktyWykresuFunkcji[377].X = 3.77
	punktyWykresuFunkcji[377].Y = 9.733

	punktyWykresuFunkcji[378].X = 3.78
	punktyWykresuFunkcji[378].Y = 9.824

	punktyWykresuFunkcji[379].X = 3.79
	punktyWykresuFunkcji[379].Y = 9.917

	punktyWykresuFunkcji[380].X = 3.8
	punktyWykresuFunkcji[380].Y = 10.011

	punktyWykresuFunkcji[381].X = 3.81
	punktyWykresuFunkcji[381].Y = 10.107

	punktyWykresuFunkcji[382].X = 3.82
	punktyWykresuFunkcji[382].Y = 10.204

	punktyWykresuFunkcji[383].X = 3.83
	punktyWykresuFunkcji[383].Y = 10.303

	punktyWykresuFunkcji[384].X = 3.84
	punktyWykresuFunkcji[384].Y = 10.404

	punktyWykresuFunkcji[385].X = 3.85
	punktyWykresuFunkcji[385].Y = 10.507

	punktyWykresuFunkcji[386].X = 3.86
	punktyWykresuFunkcji[386].Y = 10.611

	punktyWykresuFunkcji[387].X = 3.87
	punktyWykresuFunkcji[387].Y = 10.717

	punktyWykresuFunkcji[388].X = 3.88
	punktyWykresuFunkcji[388].Y = 10.825

	punktyWykresuFunkcji[389].X = 3.89
	punktyWykresuFunkcji[389].Y = 10.935

	punktyWykresuFunkcji[390].X = 3.9
	punktyWykresuFunkcji[390].Y = 11.046

	punktyWykresuFunkcji[391].X = 3.91
	punktyWykresuFunkcji[391].Y = 11.16

	punktyWykresuFunkcji[392].X = 3.92
	punktyWykresuFunkcji[392].Y = 11.275

	punktyWykresuFunkcji[393].X = 3.93
	punktyWykresuFunkcji[393].Y = 11.392

	punktyWykresuFunkcji[394].X = 3.94
	punktyWykresuFunkcji[394].Y = 11.511

	punktyWykresuFunkcji[395].X = 3.95
	punktyWykresuFunkcji[395].Y = 11.632

	punktyWykresuFunkcji[396].X = 3.96
	punktyWykresuFunkcji[396].Y = 11.755

	punktyWykresuFunkcji[397].X = 3.97
	punktyWykresuFunkcji[397].Y = 11.881

	punktyWykresuFunkcji[398].X = 3.98
	punktyWykresuFunkcji[398].Y = 12.008

	punktyWykresuFunkcji[399].X = 3.99
	punktyWykresuFunkcji[399].Y = 12.137

	punktyWykresuFunkcji[400].X = 4.0
	punktyWykresuFunkcji[400].Y = 12.268

	punktyWykresuFunkcji[401].X = 4.01
	punktyWykresuFunkcji[401].Y = 12.401

	punktyWykresuFunkcji[402].X = 4.02
	punktyWykresuFunkcji[402].Y = 12.536

	punktyWykresuFunkcji[403].X = 4.03
	punktyWykresuFunkcji[403].Y = 12.674

	punktyWykresuFunkcji[404].X = 4.04
	punktyWykresuFunkcji[404].Y = 12.813

	punktyWykresuFunkcji[405].X = 4.05
	punktyWykresuFunkcji[405].Y = 12.955

	punktyWykresuFunkcji[406].X = 4.06
	punktyWykresuFunkcji[406].Y = 13.099

	punktyWykresuFunkcji[407].X = 4.07
	punktyWykresuFunkcji[407].Y = 13.245

	punktyWykresuFunkcji[408].X = 4.08
	punktyWykresuFunkcji[408].Y = 13.393

	punktyWykresuFunkcji[409].X = 4.09
	punktyWykresuFunkcji[409].Y = 13.544

	punktyWykresuFunkcji[410].X = 4.1
	punktyWykresuFunkcji[410].Y = 13.697

	punktyWykresuFunkcji[411].X = 4.11
	punktyWykresuFunkcji[411].Y = 13.852

	punktyWykresuFunkcji[412].X = 4.12
	punktyWykresuFunkcji[412].Y = 14.009

	punktyWykresuFunkcji[413].X = 4.13
	punktyWykresuFunkcji[413].Y = 14.169

	punktyWykresuFunkcji[414].X = 4.14
	punktyWykresuFunkcji[414].Y = 14.331

	punktyWykresuFunkcji[415].X = 4.15
	punktyWykresuFunkcji[415].Y = 14.496

	punktyWykresuFunkcji[416].X = 4.16
	punktyWykresuFunkcji[416].Y = 14.663

	punktyWykresuFunkcji[417].X = 4.17
	punktyWykresuFunkcji[417].Y = 14.832

	punktyWykresuFunkcji[418].X = 4.18
	punktyWykresuFunkcji[418].Y = 15.004

	punktyWykresuFunkcji[419].X = 4.19
	punktyWykresuFunkcji[419].Y = 15.178

	punktyWykresuFunkcji[420].X = 4.2
	punktyWykresuFunkcji[420].Y = 15.355

	punktyWykresuFunkcji[421].X = 4.21
	punktyWykresuFunkcji[421].Y = 15.535

	punktyWykresuFunkcji[422].X = 4.22
	punktyWykresuFunkcji[422].Y = 15.717

	punktyWykresuFunkcji[423].X = 4.23
	punktyWykresuFunkcji[423].Y = 15.901

	punktyWykresuFunkcji[424].X = 4.24
	punktyWykresuFunkcji[424].Y = 16.088

	punktyWykresuFunkcji[425].X = 4.25
	punktyWykresuFunkcji[425].Y = 16.278

	punktyWykresuFunkcji[426].X = 4.26
	punktyWykresuFunkcji[426].Y = 16.47

	punktyWykresuFunkcji[427].X = 4.27
	punktyWykresuFunkcji[427].Y = 16.665

	punktyWykresuFunkcji[428].X = 4.28
	punktyWykresuFunkcji[428].Y = 16.863

	punktyWykresuFunkcji[429].X = 4.29
	punktyWykresuFunkcji[429].Y = 17.063

	punktyWykresuFunkcji[430].X = 4.3
	punktyWykresuFunkcji[430].Y = 17.267

	punktyWykresuFunkcji[431].X = 4.31
	punktyWykresuFunkcji[431].Y = 17.473

	punktyWykresuFunkcji[432].X = 4.32
	punktyWykresuFunkcji[432].Y = 17.681

	punktyWykresuFunkcji[433].X = 4.33
	punktyWykresuFunkcji[433].Y = 17.893

	punktyWykresuFunkcji[434].X = 4.34
	punktyWykresuFunkcji[434].Y = 18.107

	punktyWykresuFunkcji[435].X = 4.35
	punktyWykresuFunkcji[435].Y = 18.324

	punktyWykresuFunkcji[436].X = 4.36
	punktyWykresuFunkcji[436].Y = 18.545

	punktyWykresuFunkcji[437].X = 4.37
	punktyWykresuFunkcji[437].Y = 18.768

	punktyWykresuFunkcji[438].X = 4.38
	punktyWykresuFunkcji[438].Y = 18.994

	punktyWykresuFunkcji[439].X = 4.39
	punktyWykresuFunkcji[439].Y = 19.223

	punktyWykresuFunkcji[440].X = 4.4
	punktyWykresuFunkcji[440].Y = 19.455

	punktyWykresuFunkcji[441].X = 4.41
	punktyWykresuFunkcji[441].Y = 19.69

	punktyWykresuFunkcji[442].X = 4.42
	punktyWykresuFunkcji[442].Y = 19.928

	punktyWykresuFunkcji[443].X = 4.43
	punktyWykresuFunkcji[443].Y = 20.169

	punktyWykresuFunkcji[444].X = 4.44
	punktyWykresuFunkcji[444].Y = 20.413

	punktyWykresuFunkcji[445].X = 4.45
	punktyWykresuFunkcji[445].Y = 20.66

	punktyWykresuFunkcji[446].X = 4.46
	punktyWykresuFunkcji[446].Y = 20.911

	punktyWykresuFunkcji[447].X = 4.47
	punktyWykresuFunkcji[447].Y = 21.164

	punktyWykresuFunkcji[448].X = 4.48
	punktyWykresuFunkcji[448].Y = 21.421

	punktyWykresuFunkcji[449].X = 4.49
	punktyWykresuFunkcji[449].Y = 21.681

	punktyWykresuFunkcji[450].X = 4.5
	punktyWykresuFunkcji[450].Y = 21.944

	punktyWykresuFunkcji[451].X = 4.51
	punktyWykresuFunkcji[451].Y = 22.211

	punktyWykresuFunkcji[452].X = 4.52
	punktyWykresuFunkcji[452].Y = 22.48

	punktyWykresuFunkcji[453].X = 4.53
	punktyWykresuFunkcji[453].Y = 22.754

	punktyWykresuFunkcji[454].X = 4.54
	punktyWykresuFunkcji[454].Y = 23.03

	punktyWykresuFunkcji[455].X = 4.55
	punktyWykresuFunkcji[455].Y = 23.31

	punktyWykresuFunkcji[456].X = 4.56
	punktyWykresuFunkcji[456].Y = 23.593

	punktyWykresuFunkcji[457].X = 4.57
	punktyWykresuFunkcji[457].Y = 23.88

	punktyWykresuFunkcji[458].X = 4.58
	punktyWykresuFunkcji[458].Y = 24.17

	punktyWykresuFunkcji[459].X = 4.59
	punktyWykresuFunkcji[459].Y = 24.464

	punktyWykresuFunkcji[460].X = 4.6
	punktyWykresuFunkcji[460].Y = 24.761

	punktyWykresuFunkcji[461].X = 4.61
	punktyWykresuFunkcji[461].Y = 25.061

	punktyWykresuFunkcji[462].X = 4.62
	punktyWykresuFunkcji[462].Y = 25.366

	punktyWykresuFunkcji[463].X = 4.63
	punktyWykresuFunkcji[463].Y = 25.673

	punktyWykresuFunkcji[464].X = 4.64
	punktyWykresuFunkcji[464].Y = 25.985

	punktyWykresuFunkcji[465].X = 4.65
	punktyWykresuFunkcji[465].Y = 26.3

	punktyWykresuFunkcji[466].X = 4.66
	punktyWykresuFunkcji[466].Y = 26.618

	punktyWykresuFunkcji[467].X = 4.67
	punktyWykresuFunkcji[467].Y = 27.267

	punktyWykresuFunkcji[468].X = 4.68
	punktyWykresuFunkcji[468].Y = 27.267

	punktyWykresuFunkcji[469].X = 4.69
	punktyWykresuFunkcji[469].Y = 27.597

	punktyWykresuFunkcji[470].X = 4.7
	punktyWykresuFunkcji[470].Y = 27.931

	punktyWykresuFunkcji[471].X = 4.71
	punktyWykresuFunkcji[471].Y = 28.268

	punktyWykresuFunkcji[472].X = 4.72
	punktyWykresuFunkcji[472].Y = 28.609

	punktyWykresuFunkcji[473].X = 4.73
	punktyWykresuFunkcji[473].Y = 28.954

	punktyWykresuFunkcji[474].X = 4.74
	punktyWykresuFunkcji[474].Y = 29.303

	punktyWykresuFunkcji[475].X = 4.75
	punktyWykresuFunkcji[475].Y = 29.656

	punktyWykresuFunkcji[476].X = 4.76
	punktyWykresuFunkcji[476].Y = 30.013

	punktyWykresuFunkcji[477].X = 4.77
	punktyWykresuFunkcji[477].Y = 30.374

	punktyWykresuFunkcji[478].X = 4.78
	punktyWykresuFunkcji[478].Y = 30.739

	punktyWykresuFunkcji[479].X = 4.79
	punktyWykresuFunkcji[479].Y = 31.108

	punktyWykresuFunkcji[480].X = 4.8
	punktyWykresuFunkcji[480].Y = 31.48

	punktyWykresuFunkcji[481].X = 4.81
	punktyWykresuFunkcji[481].Y = 31.857

	punktyWykresuFunkcji[482].X = 4.82
	punktyWykresuFunkcji[482].Y = 32.238

	punktyWykresuFunkcji[483].X = 4.83
	punktyWykresuFunkcji[483].Y = 32.624

	punktyWykresuFunkcji[484].X = 4.84
	punktyWykresuFunkcji[484].Y = 33.013

	punktyWykresuFunkcji[485].X = 4.85
	punktyWykresuFunkcji[485].Y = 33.407

	punktyWykresuFunkcji[486].X = 4.86
	punktyWykresuFunkcji[486].Y = 33.804

	punktyWykresuFunkcji[487].X = 4.87
	punktyWykresuFunkcji[487].Y = 34.206

	punktyWykresuFunkcji[488].X = 4.88
	punktyWykresuFunkcji[488].Y = 34.613

	punktyWykresuFunkcji[489].X = 4.89
	punktyWykresuFunkcji[489].Y = 35.023

	punktyWykresuFunkcji[490].X = 4.9
	punktyWykresuFunkcji[490].Y = 35.438

	punktyWykresuFunkcji[491].X = 4.91
	punktyWykresuFunkcji[491].Y = 35.857

	punktyWykresuFunkcji[492].X = 4.92
	punktyWykresuFunkcji[492].Y = 36.281

	punktyWykresuFunkcji[493].X = 4.93
	punktyWykresuFunkcji[493].Y = 36.709

	punktyWykresuFunkcji[494].X = 4.94
	punktyWykresuFunkcji[494].Y = 37.141

	punktyWykresuFunkcji[495].X = 4.95
	punktyWykresuFunkcji[495].Y = 37.578

	punktyWykresuFunkcji[496].X = 4.96
	punktyWykresuFunkcji[496].Y = 38.02

	punktyWykresuFunkcji[497].X = 4.97
	punktyWykresuFunkcji[497].Y = 38.466

	punktyWykresuFunkcji[498].X = 4.98
	punktyWykresuFunkcji[498].Y = 38.916

	punktyWykresuFunkcji[499].X = 4.99
	punktyWykresuFunkcji[499].Y = 39.372

	punktyWykresuFunkcji[500].X = 5.0
	punktyWykresuFunkcji[500].Y = 39.831

	punktyWykresuFunkcji[501].X = 5.01
	punktyWykresuFunkcji[501].Y = 40.296

	punktyWykresuFunkcji[502].X = 5.02
	punktyWykresuFunkcji[502].Y = 40.765

	punktyWykresuFunkcji[503].X = 5.03
	punktyWykresuFunkcji[503].Y = 41.238

	punktyWykresuFunkcji[504].X = 5.04
	punktyWykresuFunkcji[504].Y = 41.717

	punktyWykresuFunkcji[505].X = 5.05
	punktyWykresuFunkcji[505].Y = 42.2

	punktyWykresuFunkcji[506].X = 5.06
	punktyWykresuFunkcji[506].Y = 42.688

	punktyWykresuFunkcji[507].X = 5.07
	punktyWykresuFunkcji[507].Y = 43.181

	punktyWykresuFunkcji[508].X = 5.08
	punktyWykresuFunkcji[508].Y = 43.679

	punktyWykresuFunkcji[509].X = 5.09
	punktyWykresuFunkcji[509].Y = 44.181

	punktyWykresuFunkcji[510].X = 5.1
	punktyWykresuFunkcji[510].Y = 44.689

	punktyWykresuFunkcji[511].X = 5.11
	punktyWykresuFunkcji[511].Y = 45.201

	punktyWykresuFunkcji[512].X = 5.12
	punktyWykresuFunkcji[512].Y = 45.718

	punktyWykresuFunkcji[513].X = 5.13
	punktyWykresuFunkcji[513].Y = 46.241

	punktyWykresuFunkcji[514].X = 5.14
	punktyWykresuFunkcji[514].Y = 46.768

	punktyWykresuFunkcji[515].X = 5.15
	punktyWykresuFunkcji[515].Y = 47.301

	punktyWykresuFunkcji[516].X = 5.16
	punktyWykresuFunkcji[516].Y = 47.838

	punktyWykresuFunkcji[517].X = 5.17
	punktyWykresuFunkcji[517].Y = 48.381

	punktyWykresuFunkcji[518].X = 5.18
	punktyWykresuFunkcji[518].Y = 48.928

	punktyWykresuFunkcji[519].X = 5.19
	punktyWykresuFunkcji[519].Y = 49.481

	punktyWykresuFunkcji[520].X = 5.2
	punktyWykresuFunkcji[520].Y = 50.039

	punktyWykresuFunkcji[521].X = 5.21
	punktyWykresuFunkcji[521].Y = 50.603

	punktyWykresuFunkcji[522].X = 5.22
	punktyWykresuFunkcji[522].Y = 51.171

	punktyWykresuFunkcji[523].X = 5.23
	punktyWykresuFunkcji[523].Y = 51.745

	punktyWykresuFunkcji[524].X = 5.24
	punktyWykresuFunkcji[524].Y = 52.324

	punktyWykresuFunkcji[525].X = 5.25
	punktyWykresuFunkcji[525].Y = 52.909

	punktyWykresuFunkcji[526].X = 5.26
	punktyWykresuFunkcji[526].Y = 53.499

	punktyWykresuFunkcji[527].X = 5.27
	punktyWykresuFunkcji[527].Y = 54.094

	punktyWykresuFunkcji[528].X = 5.28
	punktyWykresuFunkcji[528].Y = 54.695

	punktyWykresuFunkcji[529].X = 5.29
	punktyWykresuFunkcji[529].Y = 55.301

	punktyWykresuFunkcji[530].X = 5.3
	punktyWykresuFunkcji[530].Y = 55.913

	punktyWykresuFunkcji[531].X = 5.31
	punktyWykresuFunkcji[531].Y = 56.53

	punktyWykresuFunkcji[532].X = 5.32
	punktyWykresuFunkcji[532].Y = 57.153

	punktyWykresuFunkcji[533].X = 5.33
	punktyWykresuFunkcji[533].Y = 57.781

	punktyWykresuFunkcji[534].X = 5.34
	punktyWykresuFunkcji[534].Y = 58.415

	punktyWykresuFunkcji[535].X = 5.35
	punktyWykresuFunkcji[535].Y = 59.055

	punktyWykresuFunkcji[536].X = 5.36
	punktyWykresuFunkcji[536].Y = 59.7

	punktyWykresuFunkcji[537].X = 5.37
	punktyWykresuFunkcji[537].Y = 60.351

	punktyWykresuFunkcji[538].X = 5.38
	punktyWykresuFunkcji[538].Y = 61.008

	punktyWykresuFunkcji[539].X = 5.39
	punktyWykresuFunkcji[539].Y = 61.67

	punktyWykresuFunkcji[540].X = 5.4
	punktyWykresuFunkcji[540].Y = 62.338

	punktyWykresuFunkcji[541].X = 5.41
	punktyWykresuFunkcji[541].Y = 63.013

	punktyWykresuFunkcji[542].X = 5.42
	punktyWykresuFunkcji[542].Y = 63.693

	punktyWykresuFunkcji[543].X = 5.43
	punktyWykresuFunkcji[543].Y = 64.378

	punktyWykresuFunkcji[544].X = 5.44
	punktyWykresuFunkcji[544].Y = 65.07

	punktyWykresuFunkcji[545].X = 5.45
	punktyWykresuFunkcji[545].Y = 65.768

	punktyWykresuFunkcji[546].X = 5.46
	punktyWykresuFunkcji[546].Y = 66.472

	punktyWykresuFunkcji[547].X = 5.47
	punktyWykresuFunkcji[547].Y = 67.181

	punktyWykresuFunkcji[548].X = 5.48
	punktyWykresuFunkcji[548].Y = 67.897

	punktyWykresuFunkcji[549].X = 5.49
	punktyWykresuFunkcji[549].Y = 68.619

	punktyWykresuFunkcji[550].X = 5.5
	punktyWykresuFunkcji[550].Y = 69.347

	punktyWykresuFunkcji[551].X = 5.51
	punktyWykresuFunkcji[551].Y = 70.081

	punktyWykresuFunkcji[552].X = 5.52
	punktyWykresuFunkcji[552].Y = 70.821

	punktyWykresuFunkcji[553].X = 5.53
	punktyWykresuFunkcji[553].Y = 71.567

	punktyWykresuFunkcji[554].X = 5.54
	punktyWykresuFunkcji[554].Y = 72.32

	punktyWykresuFunkcji[555].X = 5.55
	punktyWykresuFunkcji[555].Y = 73.079

	punktyWykresuFunkcji[556].X = 5.56
	punktyWykresuFunkcji[556].Y = 73.844

	punktyWykresuFunkcji[557].X = 5.57
	punktyWykresuFunkcji[557].Y = 74.616

	punktyWykresuFunkcji[558].X = 5.58
	punktyWykresuFunkcji[558].Y = 75.393

	punktyWykresuFunkcji[559].X = 5.59
	punktyWykresuFunkcji[559].Y = 76.178

	punktyWykresuFunkcji[560].X = 5.6
	punktyWykresuFunkcji[560].Y = 76.968

	punktyWykresuFunkcji[561].X = 5.61
	punktyWykresuFunkcji[561].Y = 77.765

	punktyWykresuFunkcji[562].X = 5.62
	punktyWykresuFunkcji[562].Y = 78.569

	punktyWykresuFunkcji[563].X = 5.63
	punktyWykresuFunkcji[563].Y = 79.379

	punktyWykresuFunkcji[564].X = 5.64
	punktyWykresuFunkcji[564].Y = 80.195

	punktyWykresuFunkcji[565].X = 5.65
	punktyWykresuFunkcji[565].Y = 81.018

	punktyWykresuFunkcji[566].X = 5.66
	punktyWykresuFunkcji[566].Y = 81.848

	punktyWykresuFunkcji[567].X = 5.67
	punktyWykresuFunkcji[567].Y = 82.684

	punktyWykresuFunkcji[568].X = 5.68
	punktyWykresuFunkcji[568].Y = 83.527

	punktyWykresuFunkcji[569].X = 5.69
	punktyWykresuFunkcji[569].Y = 84.377

	punktyWykresuFunkcji[570].X = 5.7
	punktyWykresuFunkcji[570].Y = 85.233

	punktyWykresuFunkcji[571].X = 5.71
	punktyWykresuFunkcji[571].Y = 86.096

	punktyWykresuFunkcji[572].X = 5.72
	punktyWykresuFunkcji[572].Y = 86.966

	punktyWykresuFunkcji[573].X = 5.73
	punktyWykresuFunkcji[573].Y = 87.843

	punktyWykresuFunkcji[574].X = 5.74
	punktyWykresuFunkcji[574].Y = 88.726

	punktyWykresuFunkcji[575].X = 5.75
	punktyWykresuFunkcji[575].Y = 89.616

	punktyWykresuFunkcji[576].X = 5.76
	punktyWykresuFunkcji[576].Y = 90.514

	punktyWykresuFunkcji[577].X = 5.77
	punktyWykresuFunkcji[577].Y = 91.418

	punktyWykresuFunkcji[578].X = 5.78
	punktyWykresuFunkcji[578].Y = 92.329

	punktyWykresuFunkcji[579].X = 5.79
	punktyWykresuFunkcji[579].Y = 93.247

	punktyWykresuFunkcji[580].X = 5.8
	punktyWykresuFunkcji[580].Y = 94.172

	punktyWykresuFunkcji[581].X = 5.81
	punktyWykresuFunkcji[581].Y = 95.105

	punktyWykresuFunkcji[582].X = 5.82
	punktyWykresuFunkcji[582].Y = 96.044

	punktyWykresuFunkcji[583].X = 5.83
	punktyWykresuFunkcji[583].Y = 96.99

	punktyWykresuFunkcji[584].X = 5.84
	punktyWykresuFunkcji[584].Y = 97.944

	punktyWykresuFunkcji[585].X = 5.85
	punktyWykresuFunkcji[585].Y = 98.905

	punktyWykresuFunkcji[586].X = 5.86
	punktyWykresuFunkcji[586].Y = 99.873

	punktyWykresuFunkcji[587].X = 5.87
	punktyWykresuFunkcji[587].Y = 100.848

	punktyWykresuFunkcji[588].X = 5.88
	punktyWykresuFunkcji[588].Y = 101.83

	punktyWykresuFunkcji[589].X = 5.89
	punktyWykresuFunkcji[589].Y = 102.82

	punktyWykresuFunkcji[590].X = 5.9
	punktyWykresuFunkcji[590].Y = 103.817

	punktyWykresuFunkcji[591].X = 5.91
	punktyWykresuFunkcji[591].Y = 104.822

	punktyWykresuFunkcji[592].X = 5.92
	punktyWykresuFunkcji[592].Y = 105.834

	punktyWykresuFunkcji[593].X = 5.93
	punktyWykresuFunkcji[593].Y = 106.853

	punktyWykresuFunkcji[594].X = 5.94
	punktyWykresuFunkcji[594].Y = 107.88

	punktyWykresuFunkcji[595].X = 5.95
	punktyWykresuFunkcji[595].Y = 108.914

	punktyWykresuFunkcji[596].X = 5.96
	punktyWykresuFunkcji[596].Y = 109.956

	punktyWykresuFunkcji[597].X = 5.97
	punktyWykresuFunkcji[597].Y = 111.005

	punktyWykresuFunkcji[598].X = 5.98
	punktyWykresuFunkcji[598].Y = 112.062

	punktyWykresuFunkcji[599].X = 5.99
	punktyWykresuFunkcji[599].Y = 113.127

	punktyWykresuFunkcji[600].X = 6.0
	punktyWykresuFunkcji[600].Y = 114.199

	punktyWykresuFunkcji[601].X = 6.01
	punktyWykresuFunkcji[601].Y = 115.279

	punktyWykresuFunkcji[602].X = 6.02
	punktyWykresuFunkcji[602].Y = 116.366

	punktyWykresuFunkcji[603].X = 6.03
	punktyWykresuFunkcji[603].Y = 117.461

	punktyWykresuFunkcji[604].X = 6.04
	punktyWykresuFunkcji[604].Y = 118.565

	punktyWykresuFunkcji[605].X = 6.05
	punktyWykresuFunkcji[605].Y = 119.676

	punktyWykresuFunkcji[606].X = 6.06
	punktyWykresuFunkcji[606].Y = 120.794

	punktyWykresuFunkcji[607].X = 6.07
	punktyWykresuFunkcji[607].Y = 121.921

	punktyWykresuFunkcji[608].X = 6.08
	punktyWykresuFunkcji[608].Y = 123.055

	punktyWykresuFunkcji[609].X = 6.09
	punktyWykresuFunkcji[609].Y = 125.198

	punktyWykresuFunkcji[610].X = 6.10
	punktyWykresuFunkcji[610].Y = 125.348

	punktyWykresuFunkcji[611].X = 6.11
	punktyWykresuFunkcji[611].Y = 126.507

	punktyWykresuFunkcji[612].X = 6.12
	punktyWykresuFunkcji[612].Y = 127.673

	punktyWykresuFunkcji[613].X = 6.13
	punktyWykresuFunkcji[613].Y = 128.848

	punktyWykresuFunkcji[614].X = 6.14
	punktyWykresuFunkcji[614].Y = 130.03

	punktyWykresuFunkcji[615].X = 6.15
	punktyWykresuFunkcji[615].Y = 131.221

	punktyWykresuFunkcji[616].X = 6.16
	punktyWykresuFunkcji[616].Y = 132.42

	punktyWykresuFunkcji[617].X = 6.17
	punktyWykresuFunkcji[617].Y = 133.627

	punktyWykresuFunkcji[618].X = 6.18
	punktyWykresuFunkcji[618].Y = 134.842

	punktyWykresuFunkcji[619].X = 6.19
	punktyWykresuFunkcji[619].Y = 136.065

	punktyWykresuFunkcji[620].X = 6.2
	punktyWykresuFunkcji[620].Y = 137.297

	punktyWykresuFunkcji[621].X = 6.21
	punktyWykresuFunkcji[621].Y = 138.537

	punktyWykresuFunkcji[622].X = 6.22
	punktyWykresuFunkcji[622].Y = 139.786

	punktyWykresuFunkcji[623].X = 6.23
	punktyWykresuFunkcji[623].Y = 141.043

	punktyWykresuFunkcji[624].X = 6.24
	punktyWykresuFunkcji[624].Y = 142.308

	punktyWykresuFunkcji[625].X = 6.25
	punktyWykresuFunkcji[625].Y = 143.581

	punktyWykresuFunkcji[626].X = 6.26
	punktyWykresuFunkcji[626].Y = 144.864

	punktyWykresuFunkcji[627].X = 6.27
	punktyWykresuFunkcji[627].Y = 146.154

	punktyWykresuFunkcji[628].X = 6.28
	punktyWykresuFunkcji[628].Y = 147.453

	punktyWykresuFunkcji[629].X = 6.29
	punktyWykresuFunkcji[629].Y = 148.761

	punktyWykresuFunkcji[630].X = 6.3
	punktyWykresuFunkcji[630].Y = 150.077

	punktyWykresuFunkcji[631].X = 6.31
	punktyWykresuFunkcji[631].Y = 151.402

	punktyWykresuFunkcji[632].X = 6.32
	punktyWykresuFunkcji[632].Y = 152.736

	punktyWykresuFunkcji[633].X = 6.33
	punktyWykresuFunkcji[633].Y = 154.078

	punktyWykresuFunkcji[634].X = 6.34
	punktyWykresuFunkcji[634].Y = 155.429

	punktyWykresuFunkcji[635].X = 6.35
	punktyWykresuFunkcji[635].Y = 156.789

	punktyWykresuFunkcji[636].X = 6.36
	punktyWykresuFunkcji[636].Y = 158.157

	punktyWykresuFunkcji[637].X = 6.37
	punktyWykresuFunkcji[637].Y = 159.535

	punktyWykresuFunkcji[638].X = 6.38
	punktyWykresuFunkcji[638].Y = 160.921

	punktyWykresuFunkcji[639].X = 6.39
	punktyWykresuFunkcji[639].Y = 162.316

	punktyWykresuFunkcji[640].X = 6.4
	punktyWykresuFunkcji[640].Y = 163.72

	punktyWykresuFunkcji[641].X = 6.41
	punktyWykresuFunkcji[641].Y = 165.133

	punktyWykresuFunkcji[642].X = 6.42
	punktyWykresuFunkcji[642].Y = 166.555

	punktyWykresuFunkcji[643].X = 6.43
	punktyWykresuFunkcji[643].Y = 167.986

	punktyWykresuFunkcji[644].X = 6.44
	punktyWykresuFunkcji[644].Y = 169.426

	punktyWykresuFunkcji[645].X = 6.45
	punktyWykresuFunkcji[645].Y = 170.875

	punktyWykresuFunkcji[646].X = 6.46
	punktyWykresuFunkcji[646].Y = 172.333

	punktyWykresuFunkcji[647].X = 6.47
	punktyWykresuFunkcji[647].Y = 173.8

	punktyWykresuFunkcji[648].X = 6.48
	punktyWykresuFunkcji[648].Y = 175.277

	punktyWykresuFunkcji[649].X = 6.49
	punktyWykresuFunkcji[649].Y = 176.762

	punktyWykresuFunkcji[650].X = 6.5
	punktyWykresuFunkcji[650].Y = 178.257

	punktyWykresuFunkcji[651].X = 6.51
	punktyWykresuFunkcji[651].Y = 179.761

	punktyWykresuFunkcji[652].X = 6.52
	punktyWykresuFunkcji[652].Y = 181.275

	punktyWykresuFunkcji[653].X = 6.53
	punktyWykresuFunkcji[653].Y = 182.798

	punktyWykresuFunkcji[654].X = 6.54
	punktyWykresuFunkcji[654].Y = 184.33

	punktyWykresuFunkcji[655].X = 6.55
	punktyWykresuFunkcji[655].Y = 185.871

	punktyWykresuFunkcji[656].X = 6.56
	punktyWykresuFunkcji[656].Y = 187.422

	punktyWykresuFunkcji[657].X = 6.57
	punktyWykresuFunkcji[657].Y = 188.983

	punktyWykresuFunkcji[658].X = 6.58
	punktyWykresuFunkcji[658].Y = 190.552

	punktyWykresuFunkcji[659].X = 6.59
	punktyWykresuFunkcji[659].Y = 192.132

	punktyWykresuFunkcji[660].X = 6.6
	punktyWykresuFunkcji[660].Y = 193.721

	punktyWykresuFunkcji[661].X = 6.61
	punktyWykresuFunkcji[661].Y = 195.319

	punktyWykresuFunkcji[662].X = 6.62
	punktyWykresuFunkcji[662].Y = 196.927

	punktyWykresuFunkcji[663].X = 6.63
	punktyWykresuFunkcji[663].Y = 198.545

	punktyWykresuFunkcji[664].X = 6.64
	punktyWykresuFunkcji[664].Y = 200.173

	punktyWykresuFunkcji[665].X = 6.65
	punktyWykresuFunkcji[665].Y = 201.81

	punktyWykresuFunkcji[666].X = 6.66
	punktyWykresuFunkcji[666].Y = 203.457

	punktyWykresuFunkcji[667].X = 6.67
	punktyWykresuFunkcji[667].Y = 205.114

	punktyWykresuFunkcji[668].X = 6.68
	punktyWykresuFunkcji[668].Y = 206.78

	punktyWykresuFunkcji[669].X = 6.69
	punktyWykresuFunkcji[669].Y = 208.456

	punktyWykresuFunkcji[670].X = 6.7
	punktyWykresuFunkcji[670].Y = 210.143

	punktyWykresuFunkcji[671].X = 6.71
	punktyWykresuFunkcji[671].Y = 211.839

	punktyWykresuFunkcji[672].X = 6.72
	punktyWykresuFunkcji[672].Y = 213.545

	punktyWykresuFunkcji[673].X = 6.73
	punktyWykresuFunkcji[673].Y = 215.261

	punktyWykresuFunkcji[674].X = 6.74
	punktyWykresuFunkcji[674].Y = 216.987

	punktyWykresuFunkcji[675].X = 6.75
	punktyWykresuFunkcji[675].Y = 218.723

	punktyWykresuFunkcji[676].X = 6.76
	punktyWykresuFunkcji[676].Y = 220.469

	punktyWykresuFunkcji[677].X = 6.77
	punktyWykresuFunkcji[677].Y = 222.225

	punktyWykresuFunkcji[678].X = 6.78
	punktyWykresuFunkcji[678].Y = 223.992

	punktyWykresuFunkcji[679].X = 6.79
	punktyWykresuFunkcji[679].Y = 225.768

	punktyWykresuFunkcji[680].X = 6.8
	punktyWykresuFunkcji[680].Y = 227.555

	punktyWykresuFunkcji[681].X = 6.81
	punktyWykresuFunkcji[681].Y = 229.352

	punktyWykresuFunkcji[682].X = 6.82
	punktyWykresuFunkcji[682].Y = 231.159

	punktyWykresuFunkcji[683].X = 6.83
	punktyWykresuFunkcji[683].Y = 232.977

	punktyWykresuFunkcji[684].X = 6.84
	punktyWykresuFunkcji[684].Y = 234.804

	punktyWykresuFunkcji[685].X = 6.85
	punktyWykresuFunkcji[685].Y = 236.643

	punktyWykresuFunkcji[686].X = 6.86
	punktyWykresuFunkcji[686].Y = 238.491

	punktyWykresuFunkcji[687].X = 6.87
	punktyWykresuFunkcji[687].Y = 240.35

	punktyWykresuFunkcji[688].X = 6.88
	punktyWykresuFunkcji[688].Y = 242.22

	punktyWykresuFunkcji[689].X = 6.89
	punktyWykresuFunkcji[689].Y = 244.1

	punktyWykresuFunkcji[690].X = 6.9
	punktyWykresuFunkcji[690].Y = 245.99

	punktyWykresuFunkcji[691].X = 6.91
	punktyWykresuFunkcji[691].Y = 247.891

	punktyWykresuFunkcji[692].X = 6.92
	punktyWykresuFunkcji[692].Y = 249.803

	punktyWykresuFunkcji[693].X = 6.93
	punktyWykresuFunkcji[693].Y = 251.725

	punktyWykresuFunkcji[694].X = 6.94
	punktyWykresuFunkcji[694].Y = 253.658

	punktyWykresuFunkcji[695].X = 6.95
	punktyWykresuFunkcji[695].Y = 255.601

	punktyWykresuFunkcji[696].X = 6.96
	punktyWykresuFunkcji[696].Y = 257.555

	punktyWykresuFunkcji[697].X = 6.97
	punktyWykresuFunkcji[697].Y = 259.52

	punktyWykresuFunkcji[698].X = 6.98
	punktyWykresuFunkcji[698].Y = 261.496

	punktyWykresuFunkcji[699].X = 6.99
	punktyWykresuFunkcji[699].Y = 263.483

	punktyWykresuFunkcji[700].X = 7.0
	punktyWykresuFunkcji[700].Y = 265.48

	punktyWykresuFunkcji[701].X = 7.01
	punktyWykresuFunkcji[701].Y = 267.488

	punktyWykresuFunkcji[702].X = 7.02
	punktyWykresuFunkcji[702].Y = 269.508

	punktyWykresuFunkcji[703].X = 7.03
	punktyWykresuFunkcji[703].Y = 271.538

	punktyWykresuFunkcji[704].X = 7.04
	punktyWykresuFunkcji[704].Y = 273.579

	punktyWykresuFunkcji[705].X = 7.05
	punktyWykresuFunkcji[705].Y = 275.579

	punktyWykresuFunkcji[706].X = 7.06
	punktyWykresuFunkcji[706].Y = 277.694

	punktyWykresuFunkcji[707].X = 7.07
	punktyWykresuFunkcji[707].Y = 279.768

	punktyWykresuFunkcji[708].X = 7.08
	punktyWykresuFunkcji[708].Y = 281.854

	punktyWykresuFunkcji[709].X = 7.09
	punktyWykresuFunkcji[709].Y = 283.95

	punktyWykresuFunkcji[710].X = 7.1
	punktyWykresuFunkcji[710].Y = 286.058

	punktyWykresuFunkcji[711].X = 7.11
	punktyWykresuFunkcji[711].Y = 288.177

	punktyWykresuFunkcji[712].X = 7.12
	punktyWykresuFunkcji[712].Y = 290.307

	punktyWykresuFunkcji[713].X = 7.13
	punktyWykresuFunkcji[713].Y = 292.448

	punktyWykresuFunkcji[714].X = 7.14
	punktyWykresuFunkcji[714].Y = 294.601

	punktyWykresuFunkcji[715].X = 7.15
	punktyWykresuFunkcji[715].Y = 296.765

	punktyWykresuFunkcji[716].X = 7.16
	punktyWykresuFunkcji[716].Y = 298.94

	punktyWykresuFunkcji[717].X = 7.17
	punktyWykresuFunkcji[717].Y = 301.127

	punktyWykresuFunkcji[718].X = 7.18
	punktyWykresuFunkcji[718].Y = 303.325

	punktyWykresuFunkcji[719].X = 7.19
	punktyWykresuFunkcji[719].Y = 305.535

	punktyWykresuFunkcji[720].X = 7.2
	punktyWykresuFunkcji[720].Y = 307.756

	punktyWykresuFunkcji[721].X = 7.21
	punktyWykresuFunkcji[721].Y = 309.989

	punktyWykresuFunkcji[722].X = 7.22
	punktyWykresuFunkcji[722].Y = 312.233

	punktyWykresuFunkcji[723].X = 7.23
	punktyWykresuFunkcji[723].Y = 314.489

	punktyWykresuFunkcji[724].X = 7.24
	punktyWykresuFunkcji[724].Y = 316.757

	punktyWykresuFunkcji[725].X = 7.25
	punktyWykresuFunkcji[725].Y = 319.036

	punktyWykresuFunkcji[726].X = 7.26
	punktyWykresuFunkcji[726].Y = 321.327

	punktyWykresuFunkcji[727].X = 7.27
	punktyWykresuFunkcji[727].Y = 323.629

	punktyWykresuFunkcji[728].X = 7.28
	punktyWykresuFunkcji[728].Y = 325.944

	punktyWykresuFunkcji[729].X = 7.29
	punktyWykresuFunkcji[729].Y = 328.27

	punktyWykresuFunkcji[730].X = 7.3
	punktyWykresuFunkcji[730].Y = 330.608

	punktyWykresuFunkcji[731].X = 7.31
	punktyWykresuFunkcji[731].Y = 332.958

	punktyWykresuFunkcji[732].X = 7.32
	punktyWykresuFunkcji[732].Y = 335.32

	punktyWykresuFunkcji[733].X = 7.33
	punktyWykresuFunkcji[733].Y = 337.694

	punktyWykresuFunkcji[734].X = 7.34
	punktyWykresuFunkcji[734].Y = 340.079

	punktyWykresuFunkcji[735].X = 7.35
	punktyWykresuFunkcji[735].Y = 342.477

	punktyWykresuFunkcji[736].X = 7.36
	punktyWykresuFunkcji[736].Y = 344.887

	punktyWykresuFunkcji[737].X = 7.37
	punktyWykresuFunkcji[737].Y = 347.309

	punktyWykresuFunkcji[738].X = 7.38
	punktyWykresuFunkcji[738].Y = 349.743

	punktyWykresuFunkcji[739].X = 7.39
	punktyWykresuFunkcji[739].Y = 352.189

	punktyWykresuFunkcji[740].X = 7.4
	punktyWykresuFunkcji[740].Y = 354.647

	punktyWykresuFunkcji[741].X = 7.41
	punktyWykresuFunkcji[741].Y = 357.118

	punktyWykresuFunkcji[742].X = 7.42
	punktyWykresuFunkcji[742].Y = 359.6

	punktyWykresuFunkcji[743].X = 7.43
	punktyWykresuFunkcji[743].Y = 362.095

	punktyWykresuFunkcji[744].X = 7.44
	punktyWykresuFunkcji[744].Y = 364.603

	punktyWykresuFunkcji[745].X = 7.45
	punktyWykresuFunkcji[745].Y = 367.122

	punktyWykresuFunkcji[746].X = 7.46
	punktyWykresuFunkcji[746].Y = 369.654

	punktyWykresuFunkcji[747].X = 7.47
	punktyWykresuFunkcji[747].Y = 372.199

	punktyWykresuFunkcji[748].X = 7.48
	punktyWykresuFunkcji[748].Y = 374.756

	punktyWykresuFunkcji[749].X = 7.49
	punktyWykresuFunkcji[749].Y = 377.325

	punktyWykresuFunkcji[750].X = 7.5
	punktyWykresuFunkcji[750].Y = 379.907

	punktyWykresuFunkcji[751].X = 7.51
	punktyWykresuFunkcji[751].Y = 382.501

	punktyWykresuFunkcji[752].X = 7.52
	punktyWykresuFunkcji[752].Y = 385.108

	punktyWykresuFunkcji[753].X = 7.53
	punktyWykresuFunkcji[753].Y = 387.728

	punktyWykresuFunkcji[754].X = 7.54
	punktyWykresuFunkcji[754].Y = 390.36

	punktyWykresuFunkcji[755].X = 7.55
	punktyWykresuFunkcji[755].Y = 393.005

	punktyWykresuFunkcji[756].X = 7.56
	punktyWykresuFunkcji[756].Y = 395.663

	punktyWykresuFunkcji[757].X = 7.57
	punktyWykresuFunkcji[757].Y = 398.333

	punktyWykresuFunkcji[758].X = 7.58
	punktyWykresuFunkcji[758].Y = 401.017

	punktyWykresuFunkcji[759].X = 7.59
	punktyWykresuFunkcji[759].Y = 403.713

	punktyWykresuFunkcji[760].X = 7.6
	punktyWykresuFunkcji[760].Y = 406.143

	punktyWykresuFunkcji[761].X = 7.61
	punktyWykresuFunkcji[761].Y = 409.422

	punktyWykresuFunkcji[762].X = 7.62
	punktyWykresuFunkcji[762].Y = 411.878

	punktyWykresuFunkcji[763].X = 7.63
	punktyWykresuFunkcji[763].Y = 414.626

	punktyWykresuFunkcji[764].X = 7.64
	punktyWykresuFunkcji[764].Y = 417.387

	punktyWykresuFunkcji[765].X = 7.65
	punktyWykresuFunkcji[765].Y = 420.16

	punktyWykresuFunkcji[766].X = 7.66
	punktyWykresuFunkcji[766].Y = 422.947

	punktyWykresuFunkcji[767].X = 7.67
	punktyWykresuFunkcji[767].Y = 425.747

	punktyWykresuFunkcji[768].X = 7.68
	punktyWykresuFunkcji[768].Y = 428.56

	punktyWykresuFunkcji[769].X = 7.69
	punktyWykresuFunkcji[769].Y = 431.386

	punktyWykresuFunkcji[770].X = 7.7
	punktyWykresuFunkcji[770].Y = 434.226

	punktyWykresuFunkcji[771].X = 7.71
	punktyWykresuFunkcji[771].Y = 437.078

	punktyWykresuFunkcji[772].X = 7.72
	punktyWykresuFunkcji[772].Y = 439.944

	punktyWykresuFunkcji[773].X = 7.73
	punktyWykresuFunkcji[773].Y = 442.824

	punktyWykresuFunkcji[774].X = 7.74
	punktyWykresuFunkcji[774].Y = 445.716

	punktyWykresuFunkcji[775].X = 7.75
	punktyWykresuFunkcji[775].Y = 448.622

	punktyWykresuFunkcji[776].X = 7.76
	punktyWykresuFunkcji[776].Y = 451.542

	punktyWykresuFunkcji[777].X = 7.77
	punktyWykresuFunkcji[777].Y = 454.474

	punktyWykresuFunkcji[778].X = 7.78
	punktyWykresuFunkcji[778].Y = 457.421

	punktyWykresuFunkcji[779].X = 7.79
	punktyWykresuFunkcji[779].Y = 460.381

	punktyWykresuFunkcji[780].X = 7.8
	punktyWykresuFunkcji[780].Y = 463.354

	punktyWykresuFunkcji[781].X = 7.81
	punktyWykresuFunkcji[781].Y = 466.341

	punktyWykresuFunkcji[782].X = 7.82
	punktyWykresuFunkcji[782].Y = 469.342

	punktyWykresuFunkcji[783].X = 7.83
	punktyWykresuFunkcji[783].Y = 472.356

	punktyWykresuFunkcji[784].X = 7.84
	punktyWykresuFunkcji[784].Y = 475.384

	punktyWykresuFunkcji[785].X = 7.85
	punktyWykresuFunkcji[785].Y = 478.426

	punktyWykresuFunkcji[786].X = 7.86
	punktyWykresuFunkcji[786].Y = 481.482

	punktyWykresuFunkcji[787].X = 7.87
	punktyWykresuFunkcji[787].Y = 484.551

	punktyWykresuFunkcji[788].X = 7.88
	punktyWykresuFunkcji[788].Y = 487.634

	punktyWykresuFunkcji[789].X = 7.89
	punktyWykresuFunkcji[789].Y = 490.732

	punktyWykresuFunkcji[790].X = 7.9
	punktyWykresuFunkcji[790].Y = 493.843

	punktyWykresuFunkcji[791].X = 7.91
	punktyWykresuFunkcji[791].Y = 496.968

	punktyWykresuFunkcji[792].X = 7.92
	punktyWykresuFunkcji[792].Y = 500.26

	punktyWykresuFunkcji[793].X = 7.93
	punktyWykresuFunkcji[793].Y = 503.26

	punktyWykresuFunkcji[794].X = 7.94
	punktyWykresuFunkcji[794].Y = 506.427

	punktyWykresuFunkcji[795].X = 7.95
	punktyWykresuFunkcji[795].Y = 509.608

	punktyWykresuFunkcji[796].X = 7.96
	punktyWykresuFunkcji[796].Y = 512.803

	punktyWykresuFunkcji[797].X = 7.97
	punktyWykresuFunkcji[797].Y = 516.013

	punktyWykresuFunkcji[798].X = 7.98
	punktyWykresuFunkcji[798].Y = 519.237

	punktyWykresuFunkcji[799].X = 7.99
	punktyWykresuFunkcji[799].Y = 522.475

	punktyWykresuFunkcji[800].X = 8.0
	punktyWykresuFunkcji[800].Y = 525.727

	punktyWykresuFunkcji[801].X = 8.01
	punktyWykresuFunkcji[801].Y = 528.994

	punktyWykresuFunkcji[802].X = 8.02
	punktyWykresuFunkcji[802].Y = 532.275

	punktyWykresuFunkcji[803].X = 8.03
	punktyWykresuFunkcji[803].Y = 535.57

	punktyWykresuFunkcji[804].X = 8.04
	punktyWykresuFunkcji[804].Y = 538.88

	punktyWykresuFunkcji[805].X = 8.05
	punktyWykresuFunkcji[805].Y = 542.204

	punktyWykresuFunkcji[806].X = 8.06
	punktyWykresuFunkcji[806].Y = 545.543

	punktyWykresuFunkcji[807].X = 8.07
	punktyWykresuFunkcji[807].Y = 548.896

	punktyWykresuFunkcji[808].X = 8.08
	punktyWykresuFunkcji[808].Y = 552.264

	punktyWykresuFunkcji[809].X = 8.09
	punktyWykresuFunkcji[809].Y = 555.647

	punktyWykresuFunkcji[810].X = 8.1
	punktyWykresuFunkcji[810].Y = 559.044

	punktyWykresuFunkcji[811].X = 8.11
	punktyWykresuFunkcji[811].Y = 562.456

	punktyWykresuFunkcji[812].X = 8.12
	punktyWykresuFunkcji[812].Y = 565.883

	punktyWykresuFunkcji[813].X = 8.13
	punktyWykresuFunkcji[813].Y = 569.324

	punktyWykresuFunkcji[814].X = 8.14
	punktyWykresuFunkcji[814].Y = 572.781

	punktyWykresuFunkcji[815].X = 8.15
	punktyWykresuFunkcji[815].Y = 576.252

	punktyWykresuFunkcji[816].X = 8.16
	punktyWykresuFunkcji[816].Y = 579.738

	punktyWykresuFunkcji[817].X = 8.17
	punktyWykresuFunkcji[817].Y = 0.0

	punktyWykresuFunkcji[818].X = 8.18
	punktyWykresuFunkcji[818].Y = 0.0

	punktyWykresuFunkcji[819].X = 8.19
	punktyWykresuFunkcji[819].Y = 0.0

	punktyWykresuFunkcji[820].X = 8.2
	punktyWykresuFunkcji[820].Y = 0.0

	punktyWykresuFunkcji[821].X = 8.21
	punktyWykresuFunkcji[821].Y = 0.0

	punktyWykresuFunkcji[822].X = 8.22
	punktyWykresuFunkcji[822].Y = 0.0

	punktyWykresuFunkcji[823].X = 8.23
	punktyWykresuFunkcji[823].Y = 0.0

	punktyWykresuFunkcji[824].X = 8.24
	punktyWykresuFunkcji[824].Y = 0.0

	punktyWykresuFunkcji[825].X = 8.25
	punktyWykresuFunkcji[825].Y = 0.0

	punktyWykresuFunkcji[826].X = 8.26
	punktyWykresuFunkcji[826].Y = 0.0

	punktyWykresuFunkcji[827].X = 8.27
	punktyWykresuFunkcji[827].Y = 0.0

	punktyWykresuFunkcji[828].X = 8.28
	punktyWykresuFunkcji[828].Y = 0.0

	punktyWykresuFunkcji[829].X = 8.29
	punktyWykresuFunkcji[829].Y = 0.0

	punktyWykresuFunkcji[830].X = 8.3
	punktyWykresuFunkcji[830].Y = 0.0

	punktyWykresuFunkcji[831].X = 8.31
	punktyWykresuFunkcji[831].Y = 0.0

	punktyWykresuFunkcji[832].X = 8.32
	punktyWykresuFunkcji[832].Y = 0.0

	punktyWykresuFunkcji[833].X = 8.33
	punktyWykresuFunkcji[833].Y = 0.0

	punktyWykresuFunkcji[834].X = 8.34
	punktyWykresuFunkcji[834].Y = 0.0

	punktyWykresuFunkcji[835].X = 8.35
	punktyWykresuFunkcji[835].Y = 0.0

	punktyWykresuFunkcji[836].X = 8.36
	punktyWykresuFunkcji[836].Y = 0.0

	punktyWykresuFunkcji[837].X = 8.37
	punktyWykresuFunkcji[837].Y = 0.0

	punktyWykresuFunkcji[838].X = 8.38
	punktyWykresuFunkcji[838].Y = 0.0

	punktyWykresuFunkcji[839].X = 8.39
	punktyWykresuFunkcji[839].Y = 0.0

	punktyWykresuFunkcji[840].X = 8.4
	punktyWykresuFunkcji[840].Y = 0.0



































































	wykresFunkcji := plot.New()

	wykresFunkcji.Title.Text = "Wykres funkcji f(x) = 0.25x^4 - x^3 + x - 5cos(x) + 5"

	wykresFunkcji.X.Label.Text = "x"
	wykresFunkcji.Y.Label.Text = "y"

	liniaWykresu, err := plotter.NewLine(punktyWykresuFunkcji)

	if err != nil {
		panic(err)
	}

	liniaWykresu.LineStyle.Width = vg.Points(0.1)
	liniaWykresu.Color = color.RGBA{R: 200, G: 100, B: 100}

	wykresFunkcji.Add(liniaWykresu)
	wykresFunkcji.Legend.Add("f(x)", liniaWykresu)

	if err := wykresFunkcji.Save(10*vg.Inch, 10*vg.Inch,
		"Wykresy-funkcji-01.png"); err != nil {

		panic(err)
	}
}
