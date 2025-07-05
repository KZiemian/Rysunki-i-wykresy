package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres funkcji sin(x)

	punktyWykresuFunkcji := make(plotter.XYs, 631)

	punktyWykresuFunkcji[0].X = 0.0
	punktyWykresuFunkcji[0].Y = 0.0

	punktyWykresuFunkcji[1].X = 0.01
	punktyWykresuFunkcji[1].Y = 0.01

	punktyWykresuFunkcji[2].X = 0.02
	punktyWykresuFunkcji[2].Y = 0.02

	punktyWykresuFunkcji[3].X = 0.03
	punktyWykresuFunkcji[3].Y = 0.03

	punktyWykresuFunkcji[4].X = 0.04
	punktyWykresuFunkcji[4].Y = 0.04

	punktyWykresuFunkcji[5].X = 0.05
	punktyWykresuFunkcji[5].Y = 0.05

	punktyWykresuFunkcji[6].X = 0.06
	punktyWykresuFunkcji[6].Y = 0.06

	punktyWykresuFunkcji[7].X = 0.07
	punktyWykresuFunkcji[7].Y = 0.069

	punktyWykresuFunkcji[8].X = 0.08
	punktyWykresuFunkcji[8].Y = 0.079

	punktyWykresuFunkcji[9].X = 0.09
	punktyWykresuFunkcji[9].Y = 0.089

	punktyWykresuFunkcji[10].X = 0.1
	punktyWykresuFunkcji[10].Y = 0.099

	punktyWykresuFunkcji[11].X = 0.11
	punktyWykresuFunkcji[11].Y = 0.109

	punktyWykresuFunkcji[12].X = 0.12
	punktyWykresuFunkcji[12].Y = 0.119

	punktyWykresuFunkcji[13].X = 0.13
	punktyWykresuFunkcji[13].Y = 0.129

	punktyWykresuFunkcji[14].X = 0.14
	punktyWykresuFunkcji[14].Y = 0.139

	punktyWykresuFunkcji[15].X = 0.15
	punktyWykresuFunkcji[15].Y = 0.149

	punktyWykresuFunkcji[16].X = 0.16
	punktyWykresuFunkcji[16].Y = 0.159

	punktyWykresuFunkcji[17].X = 0.17
	punktyWykresuFunkcji[17].Y = 0.169

	punktyWykresuFunkcji[18].X = 0.18
	punktyWykresuFunkcji[18].Y = 0.179

	punktyWykresuFunkcji[19].X = 0.19
	punktyWykresuFunkcji[19].Y = 0.188

	punktyWykresuFunkcji[20].X = 0.2
	punktyWykresuFunkcji[20].Y = 0.198

	punktyWykresuFunkcji[21].X = 0.21
	punktyWykresuFunkcji[21].Y = 0.208

	punktyWykresuFunkcji[22].X = 0.22
	punktyWykresuFunkcji[22].Y = 0.218

	punktyWykresuFunkcji[23].X = 0.23
	punktyWykresuFunkcji[23].Y = 0.228

	punktyWykresuFunkcji[24].X = 0.24
	punktyWykresuFunkcji[24].Y = 0.237

	punktyWykresuFunkcji[25].X = 0.25
	punktyWykresuFunkcji[25].Y = 0.247

	punktyWykresuFunkcji[26].X = 0.26
	punktyWykresuFunkcji[26].Y = 0.257

	punktyWykresuFunkcji[27].X = 0.27
	punktyWykresuFunkcji[27].Y = 0.266

	punktyWykresuFunkcji[28].X = 0.28
	punktyWykresuFunkcji[28].Y = 0.276

	punktyWykresuFunkcji[29].X = 0.29
	punktyWykresuFunkcji[29].Y = 0.286

	punktyWykresuFunkcji[30].X = 0.3
	punktyWykresuFunkcji[30].Y = 0.295

	punktyWykresuFunkcji[31].X = 0.31
	punktyWykresuFunkcji[31].Y = 0.305

	punktyWykresuFunkcji[32].X = 0.32
	punktyWykresuFunkcji[32].Y = 0.314

	punktyWykresuFunkcji[33].X = 0.33
	punktyWykresuFunkcji[33].Y = 0.324

	punktyWykresuFunkcji[34].X = 0.34
	punktyWykresuFunkcji[34].Y = 0.333

	punktyWykresuFunkcji[35].X = 0.35
	punktyWykresuFunkcji[35].Y = 0.342

	punktyWykresuFunkcji[36].X = 0.36
	punktyWykresuFunkcji[36].Y = 0.352

	punktyWykresuFunkcji[37].X = 0.37
	punktyWykresuFunkcji[37].Y = 0.361

	punktyWykresuFunkcji[38].X = 0.38
	punktyWykresuFunkcji[38].Y = 0.37

	punktyWykresuFunkcji[39].X = 0.39
	punktyWykresuFunkcji[39].Y = 0.38

	punktyWykresuFunkcji[40].X = 0.40
	punktyWykresuFunkcji[40].Y = 0.389

	punktyWykresuFunkcji[41].X = 0.41
	punktyWykresuFunkcji[41].Y = 0.398

	punktyWykresuFunkcji[42].X = 0.42
	punktyWykresuFunkcji[42].Y = 0.407

	punktyWykresuFunkcji[43].X = 0.43
	punktyWykresuFunkcji[43].Y = 0.416

	punktyWykresuFunkcji[44].X = 0.44
	punktyWykresuFunkcji[44].Y = 0.425

	punktyWykresuFunkcji[45].X = 0.45
	punktyWykresuFunkcji[45].Y = 0.435

	punktyWykresuFunkcji[46].X = 0.46
	punktyWykresuFunkcji[46].Y = 0.443

	punktyWykresuFunkcji[47].X = 0.47
	punktyWykresuFunkcji[47].Y = 0.452

	punktyWykresuFunkcji[48].X = 0.48
	punktyWykresuFunkcji[48].Y = 0.461

	punktyWykresuFunkcji[49].X = 0.49
	punktyWykresuFunkcji[49].Y = 0.47

	punktyWykresuFunkcji[50].X = 0.5
	punktyWykresuFunkcji[50].Y = 0.479

	punktyWykresuFunkcji[51].X = 0.51
	punktyWykresuFunkcji[51].Y = 0.488

	punktyWykresuFunkcji[52].X = 0.52
	punktyWykresuFunkcji[52].Y = 0.496

	punktyWykresuFunkcji[53].X = 0.53
	punktyWykresuFunkcji[53].Y = 0.505

	punktyWykresuFunkcji[54].X = 0.54
	punktyWykresuFunkcji[54].Y = 0.514

	punktyWykresuFunkcji[55].X = 0.55
	punktyWykresuFunkcji[55].Y = 0.522

	punktyWykresuFunkcji[56].X = 0.56
	punktyWykresuFunkcji[56].Y = 0.531

	punktyWykresuFunkcji[57].X = 0.57
	punktyWykresuFunkcji[57].Y = 0.539

	punktyWykresuFunkcji[58].X = 0.58
	punktyWykresuFunkcji[58].Y = 0.548

	punktyWykresuFunkcji[59].X = 0.59
	punktyWykresuFunkcji[59].Y = 0.556

	punktyWykresuFunkcji[60].X = 0.6
	punktyWykresuFunkcji[60].Y = 0.564

	punktyWykresuFunkcji[61].X = 0.61
	punktyWykresuFunkcji[61].Y = 0.572

	punktyWykresuFunkcji[62].X = 0.62
	punktyWykresuFunkcji[62].Y = 0.581

	punktyWykresuFunkcji[63].X = 0.63
	punktyWykresuFunkcji[63].Y = 0.589

	punktyWykresuFunkcji[64].X = 0.64
	punktyWykresuFunkcji[64].Y = 0.597

	punktyWykresuFunkcji[65].X = 0.65
	punktyWykresuFunkcji[65].Y = 0.605

	punktyWykresuFunkcji[66].X = 0.66
	punktyWykresuFunkcji[66].Y = 0.613

	punktyWykresuFunkcji[67].X = 0.67
	punktyWykresuFunkcji[67].Y = 0.621

	punktyWykresuFunkcji[68].X = 0.68
	punktyWykresuFunkcji[68].Y = 0.628

	punktyWykresuFunkcji[69].X = 0.69
	punktyWykresuFunkcji[69].Y = 0.636

	punktyWykresuFunkcji[70].X = 0.7
	punktyWykresuFunkcji[70].Y = 0.644

	punktyWykresuFunkcji[71].X = 0.71
	punktyWykresuFunkcji[71].Y = 0.651

	punktyWykresuFunkcji[72].X = 0.72
	punktyWykresuFunkcji[72].Y = 0.659

	punktyWykresuFunkcji[73].X = 0.73
	punktyWykresuFunkcji[73].Y = 0.666

	punktyWykresuFunkcji[74].X = 0.74
	punktyWykresuFunkcji[74].Y = 0.674

	punktyWykresuFunkcji[75].X = 0.75
	punktyWykresuFunkcji[75].Y = 0.681

	punktyWykresuFunkcji[76].X = 0.76
	punktyWykresuFunkcji[76].Y = 0.688

	punktyWykresuFunkcji[77].X = 0.77
	punktyWykresuFunkcji[77].Y = 0.696

	punktyWykresuFunkcji[78].X = 0.78
	punktyWykresuFunkcji[78].Y = 0.703

	punktyWykresuFunkcji[79].X = 0.79
	punktyWykresuFunkcji[79].Y = 0.71

	punktyWykresuFunkcji[80].X = 0.8
	punktyWykresuFunkcji[80].Y = 0.717

	punktyWykresuFunkcji[81].X = 0.81
	punktyWykresuFunkcji[81].Y = 0.724

	punktyWykresuFunkcji[82].X = 0.82
	punktyWykresuFunkcji[82].Y = 0.731

	punktyWykresuFunkcji[83].X = 0.83
	punktyWykresuFunkcji[83].Y = 0.737

	punktyWykresuFunkcji[84].X = 0.84
	punktyWykresuFunkcji[84].Y = 0.744

	punktyWykresuFunkcji[85].X = 0.85
	punktyWykresuFunkcji[85].Y = 0.751

	punktyWykresuFunkcji[86].X = 0.86
	punktyWykresuFunkcji[86].Y = 0.757

	punktyWykresuFunkcji[87].X = 0.87
	punktyWykresuFunkcji[87].Y = 0.764

	punktyWykresuFunkcji[88].X = 0.88
	punktyWykresuFunkcji[88].Y = 0.77

	punktyWykresuFunkcji[89].X = 0.89
	punktyWykresuFunkcji[89].Y = 0.777

	punktyWykresuFunkcji[90].X = 0.9
	punktyWykresuFunkcji[90].Y = 0.783

	punktyWykresuFunkcji[91].X = 0.91
	punktyWykresuFunkcji[91].Y = 0.789

	punktyWykresuFunkcji[92].X = 0.92
	punktyWykresuFunkcji[92].Y = 0.795

	punktyWykresuFunkcji[93].X = 0.93
	punktyWykresuFunkcji[93].Y = 0.801

	punktyWykresuFunkcji[94].X = 0.94
	punktyWykresuFunkcji[94].Y = 0.807

	punktyWykresuFunkcji[95].X = 0.95
	punktyWykresuFunkcji[95].Y = 0.813

	punktyWykresuFunkcji[96].X = 0.96
	punktyWykresuFunkcji[96].Y = 0.819

	punktyWykresuFunkcji[97].X = 0.97
	punktyWykresuFunkcji[97].Y = 0.824

	punktyWykresuFunkcji[98].X = 0.98
	punktyWykresuFunkcji[98].Y = 0.83

	punktyWykresuFunkcji[99].X = 0.99
	punktyWykresuFunkcji[99].Y = 0.836

	punktyWykresuFunkcji[100].X = 1.0
	punktyWykresuFunkcji[100].Y = 0.841

	punktyWykresuFunkcji[101].X = 1.01
	punktyWykresuFunkcji[101].Y = 0.846

	punktyWykresuFunkcji[102].X = 1.02
	punktyWykresuFunkcji[102].Y = 0.852

	punktyWykresuFunkcji[103].X = 1.03
	punktyWykresuFunkcji[103].Y = 0.857

	punktyWykresuFunkcji[104].X = 1.04
	punktyWykresuFunkcji[104].Y = 0.862

	punktyWykresuFunkcji[105].X = 1.05
	punktyWykresuFunkcji[105].Y = 0.867

	punktyWykresuFunkcji[106].X = 1.06
	punktyWykresuFunkcji[106].Y = 0.872

	punktyWykresuFunkcji[107].X = 1.07
	punktyWykresuFunkcji[107].Y = 0.877

	punktyWykresuFunkcji[108].X = 1.08
	punktyWykresuFunkcji[108].Y = 0.882

	punktyWykresuFunkcji[109].X = 1.09
	punktyWykresuFunkcji[109].Y = 0.886

	punktyWykresuFunkcji[110].X = 1.1
	punktyWykresuFunkcji[110].Y = 0.891

	punktyWykresuFunkcji[111].X = 1.11
	punktyWykresuFunkcji[111].Y = 0.895

	punktyWykresuFunkcji[112].X = 1.12
	punktyWykresuFunkcji[112].Y = 0.9

	punktyWykresuFunkcji[113].X = 1.13
	punktyWykresuFunkcji[113].Y = 0.904

	punktyWykresuFunkcji[114].X = 1.14
	punktyWykresuFunkcji[114].Y = 0.908

	punktyWykresuFunkcji[115].X = 1.15
	punktyWykresuFunkcji[115].Y = 0.912

	punktyWykresuFunkcji[116].X = 1.16
	punktyWykresuFunkcji[116].Y = 0.916

	punktyWykresuFunkcji[117].X = 1.17
	punktyWykresuFunkcji[117].Y = 0.92

	punktyWykresuFunkcji[118].X = 1.18
	punktyWykresuFunkcji[118].Y = 0.924

	punktyWykresuFunkcji[119].X = 1.19
	punktyWykresuFunkcji[119].Y = 0.928

	punktyWykresuFunkcji[120].X = 1.2
	punktyWykresuFunkcji[120].Y = 0.932

	punktyWykresuFunkcji[121].X = 1.21
	punktyWykresuFunkcji[121].Y = 0.933

	punktyWykresuFunkcji[122].X = 1.22
	punktyWykresuFunkcji[122].Y = 0.939

	punktyWykresuFunkcji[123].X = 1.23
	punktyWykresuFunkcji[123].Y = 0.942

	punktyWykresuFunkcji[124].X = 1.24
	punktyWykresuFunkcji[124].Y = 0.945

	punktyWykresuFunkcji[125].X = 1.25
	punktyWykresuFunkcji[125].Y = 0.949

	punktyWykresuFunkcji[126].X = 1.26
	punktyWykresuFunkcji[126].Y = 0.952

	punktyWykresuFunkcji[127].X = 1.27
	punktyWykresuFunkcji[127].Y = 0.955

	punktyWykresuFunkcji[128].X = 1.28
	punktyWykresuFunkcji[128].Y = 0.958

	punktyWykresuFunkcji[129].X = 1.29
	punktyWykresuFunkcji[129].Y = 0.96

	punktyWykresuFunkcji[130].X = 1.3
	punktyWykresuFunkcji[130].Y = 0.963

	punktyWykresuFunkcji[131].X = 1.31
	punktyWykresuFunkcji[131].Y = 0.966

	punktyWykresuFunkcji[132].X = 1.32
	punktyWykresuFunkcji[132].Y = 0.968

	punktyWykresuFunkcji[133].X = 1.33
	punktyWykresuFunkcji[133].Y = 0.971

	punktyWykresuFunkcji[134].X = 1.34
	punktyWykresuFunkcji[134].Y = 0.973

	punktyWykresuFunkcji[135].X = 1.35
	punktyWykresuFunkcji[135].Y = 0.975

	punktyWykresuFunkcji[136].X = 1.36
	punktyWykresuFunkcji[136].Y = 0.977

	punktyWykresuFunkcji[137].X = 1.37
	punktyWykresuFunkcji[137].Y = 0.979

	punktyWykresuFunkcji[138].X = 1.38
	punktyWykresuFunkcji[138].Y = 0.981

	punktyWykresuFunkcji[139].X = 1.39
	punktyWykresuFunkcji[139].Y = 0.983

	punktyWykresuFunkcji[140].X = 1.4
	punktyWykresuFunkcji[140].Y = 0.985

	punktyWykresuFunkcji[141].X = 1.41
	punktyWykresuFunkcji[141].Y = 0.987

	punktyWykresuFunkcji[142].X = 1.42
	punktyWykresuFunkcji[142].Y = 0.988

	punktyWykresuFunkcji[143].X = 1.43
	punktyWykresuFunkcji[143].Y = 0.99

	punktyWykresuFunkcji[144].X = 1.44
	punktyWykresuFunkcji[144].Y = 0.991

	punktyWykresuFunkcji[145].X = 1.45
	punktyWykresuFunkcji[145].Y = 0.992

	punktyWykresuFunkcji[146].X = 1.46
	punktyWykresuFunkcji[146].Y = 0.993

	punktyWykresuFunkcji[147].X = 1.47
	punktyWykresuFunkcji[147].Y = 0.994

	punktyWykresuFunkcji[148].X = 1.48
	punktyWykresuFunkcji[148].Y = 0.995

	punktyWykresuFunkcji[149].X = 1.49
	punktyWykresuFunkcji[149].Y = 0.996

	punktyWykresuFunkcji[150].X = 1.5
	punktyWykresuFunkcji[150].Y = 0.997

	punktyWykresuFunkcji[151].X = 1.51
	punktyWykresuFunkcji[151].Y = 0.998

	punktyWykresuFunkcji[152].X = 1.52
	punktyWykresuFunkcji[152].Y = 0.998

	punktyWykresuFunkcji[153].X = 1.53
	punktyWykresuFunkcji[153].Y = 0.999

	punktyWykresuFunkcji[154].X = 1.54
	punktyWykresuFunkcji[154].Y = 0.999

	punktyWykresuFunkcji[155].X = 1.55
	punktyWykresuFunkcji[155].Y = 0.999

	punktyWykresuFunkcji[156].X = 1.56
	punktyWykresuFunkcji[156].Y = 0.999

	punktyWykresuFunkcji[157].X = 1.57
	punktyWykresuFunkcji[157].Y = 1.0

	punktyWykresuFunkcji[158].X = 1.58
	punktyWykresuFunkcji[158].Y = 1.0

	punktyWykresuFunkcji[159].X = 1.59
	punktyWykresuFunkcji[159].Y = 0.999

	punktyWykresuFunkcji[160].X = 1.6
	punktyWykresuFunkcji[160].Y = 0.999

	punktyWykresuFunkcji[161].X = 1.61
	punktyWykresuFunkcji[161].Y = 0.999

	punktyWykresuFunkcji[162].X = 1.62
	punktyWykresuFunkcji[162].Y = 0.998

	punktyWykresuFunkcji[163].X = 1.63
	punktyWykresuFunkcji[163].Y = 0.998

	punktyWykresuFunkcji[164].X = 1.64
	punktyWykresuFunkcji[164].Y = 0.997

	punktyWykresuFunkcji[165].X = 1.65
	punktyWykresuFunkcji[165].Y = 0.996

	punktyWykresuFunkcji[166].X = 1.66
	punktyWykresuFunkcji[166].Y = 0.996

	punktyWykresuFunkcji[167].X = 1.67
	punktyWykresuFunkcji[167].Y = 0.995

	punktyWykresuFunkcji[168].X = 1.68
	punktyWykresuFunkcji[168].Y = 0.994

	punktyWykresuFunkcji[169].X = 1.69
	punktyWykresuFunkcji[169].Y = 0.992

	punktyWykresuFunkcji[170].X = 1.7
	punktyWykresuFunkcji[170].Y = 0.991

	punktyWykresuFunkcji[171].X = 1.71
	punktyWykresuFunkcji[171].Y = 0.99

	punktyWykresuFunkcji[172].X = 1.72
	punktyWykresuFunkcji[172].Y = 0.988

	punktyWykresuFunkcji[173].X = 1.73
	punktyWykresuFunkcji[173].Y = 0.987

	punktyWykresuFunkcji[174].X = 1.74
	punktyWykresuFunkcji[174].Y = 0.985

	punktyWykresuFunkcji[175].X = 1.75
	punktyWykresuFunkcji[175].Y = 0.984

	punktyWykresuFunkcji[176].X = 1.76
	punktyWykresuFunkcji[176].Y = 0.982

	punktyWykresuFunkcji[177].X = 1.77
	punktyWykresuFunkcji[177].Y = 0.98

	punktyWykresuFunkcji[178].X = 1.78
	punktyWykresuFunkcji[178].Y = 0.978

	punktyWykresuFunkcji[179].X = 1.79
	punktyWykresuFunkcji[179].Y = 0.976

	punktyWykresuFunkcji[180].X = 1.8
	punktyWykresuFunkcji[180].Y = 0.973

	punktyWykresuFunkcji[181].X = 1.81
	punktyWykresuFunkcji[181].Y = 0.971

	punktyWykresuFunkcji[182].X = 1.82
	punktyWykresuFunkcji[182].Y = 0.969

	punktyWykresuFunkcji[183].X = 1.83
	punktyWykresuFunkcji[183].Y = 0.966

	punktyWykresuFunkcji[184].X = 1.84
	punktyWykresuFunkcji[184].Y = 0.964

	punktyWykresuFunkcji[185].X = 1.85
	punktyWykresuFunkcji[185].Y = 0.961

	punktyWykresuFunkcji[186].X = 1.86
	punktyWykresuFunkcji[186].Y = 0.958

	punktyWykresuFunkcji[187].X = 1.87
	punktyWykresuFunkcji[187].Y = 0.955

	punktyWykresuFunkcji[188].X = 1.88
	punktyWykresuFunkcji[188].Y = 0.952

	punktyWykresuFunkcji[189].X = 1.89
	punktyWykresuFunkcji[189].Y = 0.949

	punktyWykresuFunkcji[190].X = 1.9
	punktyWykresuFunkcji[190].Y = 0.946

	punktyWykresuFunkcji[191].X = 1.91
	punktyWykresuFunkcji[191].Y = 0.943

	punktyWykresuFunkcji[192].X = 1.92
	punktyWykresuFunkcji[192].Y = 0.939

	punktyWykresuFunkcji[193].X = 1.93
	punktyWykresuFunkcji[193].Y = 0.936

	punktyWykresuFunkcji[194].X = 1.94
	punktyWykresuFunkcji[194].Y = 0.932

	punktyWykresuFunkcji[195].X = 1.95
	punktyWykresuFunkcji[195].Y = 0.929

	punktyWykresuFunkcji[196].X = 1.96
	punktyWykresuFunkcji[196].Y = 0.925

	punktyWykresuFunkcji[197].X = 1.97
	punktyWykresuFunkcji[197].Y = 0.921

	punktyWykresuFunkcji[198].X = 1.98
	punktyWykresuFunkcji[198].Y = 0.917

	punktyWykresuFunkcji[199].X = 1.99
	punktyWykresuFunkcji[199].Y = 0.913

	punktyWykresuFunkcji[200].X = 2.0
	punktyWykresuFunkcji[200].Y = 0.909

	punktyWykresuFunkcji[201].X = 2.01
	punktyWykresuFunkcji[201].Y = 0.905

	punktyWykresuFunkcji[202].X = 2.02
	punktyWykresuFunkcji[202].Y = 0.9

	punktyWykresuFunkcji[203].X = 2.03
	punktyWykresuFunkcji[203].Y = 0.896

	punktyWykresuFunkcji[204].X = 2.04
	punktyWykresuFunkcji[204].Y = 0.891

	punktyWykresuFunkcji[205].X = 2.05
	punktyWykresuFunkcji[205].Y = 0.887

	punktyWykresuFunkcji[206].X = 2.06
	punktyWykresuFunkcji[206].Y = 0.882

	punktyWykresuFunkcji[207].X = 2.07
	punktyWykresuFunkcji[207].Y = 0.878

	punktyWykresuFunkcji[208].X = 2.08
	punktyWykresuFunkcji[208].Y = 0.873

	punktyWykresuFunkcji[209].X = 2.09
	punktyWykresuFunkcji[209].Y = 0.868

	punktyWykresuFunkcji[210].X = 2.1
	punktyWykresuFunkcji[210].Y = 0.863

	punktyWykresuFunkcji[211].X = 2.11
	punktyWykresuFunkcji[211].Y = 0.858

	punktyWykresuFunkcji[212].X = 2.12
	punktyWykresuFunkcji[212].Y = 0.852

	punktyWykresuFunkcji[213].X = 2.13
	punktyWykresuFunkcji[213].Y = 0.847

	punktyWykresuFunkcji[214].X = 2.14
	punktyWykresuFunkcji[214].Y = 0.842

	punktyWykresuFunkcji[215].X = 2.15
	punktyWykresuFunkcji[215].Y = 0.836

	punktyWykresuFunkcji[216].X = 2.16
	punktyWykresuFunkcji[216].Y = 0.831

	punktyWykresuFunkcji[217].X = 2.17
	punktyWykresuFunkcji[217].Y = 0.825

	punktyWykresuFunkcji[218].X = 2.18
	punktyWykresuFunkcji[218].Y = 0.82

	punktyWykresuFunkcji[219].X = 2.19
	punktyWykresuFunkcji[219].Y = 0.814

	punktyWykresuFunkcji[220].X = 2.2
	punktyWykresuFunkcji[220].Y = 0.808

	punktyWykresuFunkcji[221].X = 2.21
	punktyWykresuFunkcji[221].Y = 0.802

	punktyWykresuFunkcji[222].X = 2.22
	punktyWykresuFunkcji[222].Y = 0.796

	punktyWykresuFunkcji[223].X = 2.23
	punktyWykresuFunkcji[223].Y = 0.79

	punktyWykresuFunkcji[224].X = 2.24
	punktyWykresuFunkcji[224].Y = 0.784

	punktyWykresuFunkcji[225].X = 2.25
	punktyWykresuFunkcji[225].Y = 0.778

	punktyWykresuFunkcji[226].X = 2.26
	punktyWykresuFunkcji[226].Y = 0.771

	punktyWykresuFunkcji[227].X = 2.27
	punktyWykresuFunkcji[227].Y = 0.765

	punktyWykresuFunkcji[228].X = 2.28
	punktyWykresuFunkcji[228].Y = 0.758

	punktyWykresuFunkcji[229].X = 2.29
	punktyWykresuFunkcji[229].Y = 0.752

	punktyWykresuFunkcji[230].X = 2.3
	punktyWykresuFunkcji[230].Y = 0.745

	punktyWykresuFunkcji[231].X = 2.31
	punktyWykresuFunkcji[231].Y = 0.739

	punktyWykresuFunkcji[232].X = 2.32
	punktyWykresuFunkcji[232].Y = 0.732

	punktyWykresuFunkcji[233].X = 2.33
	punktyWykresuFunkcji[233].Y = 0.725

	punktyWykresuFunkcji[234].X = 2.34
	punktyWykresuFunkcji[234].Y = 0.718

	punktyWykresuFunkcji[235].X = 2.35
	punktyWykresuFunkcji[235].Y = 0.711

	punktyWykresuFunkcji[236].X = 2.36
	punktyWykresuFunkcji[236].Y = 0.704

	punktyWykresuFunkcji[237].X = 2.37
	punktyWykresuFunkcji[237].Y = 0.697

	punktyWykresuFunkcji[238].X = 2.38
	punktyWykresuFunkcji[238].Y = 0.69

	punktyWykresuFunkcji[239].X = 2.39
	punktyWykresuFunkcji[239].Y = 0.682

	punktyWykresuFunkcji[240].X = 2.4
	punktyWykresuFunkcji[240].Y = 0.675

	punktyWykresuFunkcji[241].X = 2.41
	punktyWykresuFunkcji[241].Y = 0.668

	punktyWykresuFunkcji[242].X = 2.42
	punktyWykresuFunkcji[242].Y = 0.66

	punktyWykresuFunkcji[243].X = 2.43
	punktyWykresuFunkcji[243].Y = 0.653

	punktyWykresuFunkcji[244].X = 2.44
	punktyWykresuFunkcji[244].Y = 0.645

	punktyWykresuFunkcji[245].X = 2.45
	punktyWykresuFunkcji[245].Y = 0.637

	punktyWykresuFunkcji[246].X = 2.46
	punktyWykresuFunkcji[246].Y = 0.63

	punktyWykresuFunkcji[247].X = 2.47
	punktyWykresuFunkcji[247].Y = 0.622

	punktyWykresuFunkcji[248].X = 2.48
	punktyWykresuFunkcji[248].Y = 0.614

	punktyWykresuFunkcji[249].X = 2.49
	punktyWykresuFunkcji[249].Y = 0.606

	punktyWykresuFunkcji[250].X = 2.5
	punktyWykresuFunkcji[250].Y = 0.598

	punktyWykresuFunkcji[251].X = 2.51
	punktyWykresuFunkcji[251].Y = 0.59

	punktyWykresuFunkcji[252].X = 2.52
	punktyWykresuFunkcji[252].Y = 0.582

	punktyWykresuFunkcji[253].X = 2.53
	punktyWykresuFunkcji[253].Y = 0.574

	punktyWykresuFunkcji[254].X = 2.54
	punktyWykresuFunkcji[254].Y = 0.566

	punktyWykresuFunkcji[255].X = 2.55
	punktyWykresuFunkcji[255].Y = 0.557

	punktyWykresuFunkcji[256].X = 2.56
	punktyWykresuFunkcji[256].Y = 0.549

	punktyWykresuFunkcji[257].X = 2.57
	punktyWykresuFunkcji[257].Y = 0.541

	punktyWykresuFunkcji[258].X = 2.58
	punktyWykresuFunkcji[258].Y = 0.532

	punktyWykresuFunkcji[259].X = 2.59
	punktyWykresuFunkcji[259].Y = 0.524

	punktyWykresuFunkcji[260].X = 2.6
	punktyWykresuFunkcji[260].Y = 0.515

	punktyWykresuFunkcji[261].X = 2.61
	punktyWykresuFunkcji[261].Y = 0.506

	punktyWykresuFunkcji[262].X = 2.62
	punktyWykresuFunkcji[262].Y = 0.498

	punktyWykresuFunkcji[263].X = 2.63
	punktyWykresuFunkcji[263].Y = 0.489

	punktyWykresuFunkcji[264].X = 2.64
	punktyWykresuFunkcji[264].Y = 0.48

	punktyWykresuFunkcji[265].X = 2.65
	punktyWykresuFunkcji[265].Y = 0.472

	punktyWykresuFunkcji[266].X = 2.66
	punktyWykresuFunkcji[266].Y = 0.463

	punktyWykresuFunkcji[267].X = 2.67
	punktyWykresuFunkcji[267].Y = 0.454

	punktyWykresuFunkcji[268].X = 2.68
	punktyWykresuFunkcji[268].Y = 0.445

	punktyWykresuFunkcji[269].X = 2.69
	punktyWykresuFunkcji[269].Y = 0.436

	punktyWykresuFunkcji[270].X = 2.7
	punktyWykresuFunkcji[270].Y = 0.427

	punktyWykresuFunkcji[271].X = 2.71
	punktyWykresuFunkcji[271].Y = 0.418

	punktyWykresuFunkcji[272].X = 2.72
	punktyWykresuFunkcji[272].Y = 0.409

	punktyWykresuFunkcji[273].X = 2.73
	punktyWykresuFunkcji[273].Y = 0.4

	punktyWykresuFunkcji[274].X = 2.74
	punktyWykresuFunkcji[274].Y = 0.39

	punktyWykresuFunkcji[275].X = 2.75
	punktyWykresuFunkcji[275].Y = 0.381

	punktyWykresuFunkcji[276].X = 2.76
	punktyWykresuFunkcji[276].Y = 0.372

	punktyWykresuFunkcji[277].X = 2.77
	punktyWykresuFunkcji[277].Y = 0.363

	punktyWykresuFunkcji[278].X = 2.78
	punktyWykresuFunkcji[278].Y = 0.353

	punktyWykresuFunkcji[279].X = 2.79
	punktyWykresuFunkcji[279].Y = 0.344

	punktyWykresuFunkcji[280].X = 2.8
	punktyWykresuFunkcji[280].Y = 0.335

	punktyWykresuFunkcji[281].X = 2.81
	punktyWykresuFunkcji[281].Y = 0.325

	punktyWykresuFunkcji[282].X = 2.82
	punktyWykresuFunkcji[282].Y = 0.316

	punktyWykresuFunkcji[283].X = 2.83
	punktyWykresuFunkcji[283].Y = 0.306

	punktyWykresuFunkcji[284].X = 2.84
	punktyWykresuFunkcji[284].Y = 0.297

	punktyWykresuFunkcji[285].X = 2.85
	punktyWykresuFunkcji[285].Y = 0.287

	punktyWykresuFunkcji[286].X = 2.86
	punktyWykresuFunkcji[286].Y = 0.277

	punktyWykresuFunkcji[287].X = 2.87
	punktyWykresuFunkcji[287].Y = 0.268

	punktyWykresuFunkcji[288].X = 2.88
	punktyWykresuFunkcji[288].Y = 0.258

	punktyWykresuFunkcji[289].X = 2.89
	punktyWykresuFunkcji[289].Y = 0.248

	punktyWykresuFunkcji[290].X = 2.9
	punktyWykresuFunkcji[290].Y = 0.239

	punktyWykresuFunkcji[291].X = 2.91
	punktyWykresuFunkcji[291].Y = 0.229

	punktyWykresuFunkcji[292].X = 2.92
	punktyWykresuFunkcji[292].Y = 0.219

	punktyWykresuFunkcji[293].X = 2.93
	punktyWykresuFunkcji[293].Y = 0.21

	punktyWykresuFunkcji[294].X = 2.94
	punktyWykresuFunkcji[294].Y = 0.2

	punktyWykresuFunkcji[295].X = 2.95
	punktyWykresuFunkcji[295].Y = 0.19

	punktyWykresuFunkcji[296].X = 2.96
	punktyWykresuFunkcji[296].Y = 0.18

	punktyWykresuFunkcji[297].X = 2.97
	punktyWykresuFunkcji[297].Y = 0.17

	punktyWykresuFunkcji[298].X = 2.98
	punktyWykresuFunkcji[298].Y = 0.16

	punktyWykresuFunkcji[299].X = 2.99
	punktyWykresuFunkcji[299].Y = 0.151

	punktyWykresuFunkcji[300].X = 3.0
	punktyWykresuFunkcji[300].Y = 0.141

	punktyWykresuFunkcji[301].X = 3.01
	punktyWykresuFunkcji[301].Y = 0.131

	punktyWykresuFunkcji[302].X = 3.02
	punktyWykresuFunkcji[302].Y = 0.121

	punktyWykresuFunkcji[303].X = 3.03
	punktyWykresuFunkcji[303].Y = 0.111

	punktyWykresuFunkcji[304].X = 3.04
	punktyWykresuFunkcji[304].Y = 0.101

	punktyWykresuFunkcji[305].X = 3.05
	punktyWykresuFunkcji[305].Y = 0.091

	punktyWykresuFunkcji[306].X = 3.06
	punktyWykresuFunkcji[306].Y = 0.081

	punktyWykresuFunkcji[307].X = 3.07
	punktyWykresuFunkcji[307].Y = 0.071

	punktyWykresuFunkcji[308].X = 3.08
	punktyWykresuFunkcji[308].Y = 0.061

	punktyWykresuFunkcji[309].X = 3.09
	punktyWykresuFunkcji[309].Y = 0.051

	punktyWykresuFunkcji[310].X = 3.1
	punktyWykresuFunkcji[310].Y = 0.041

	punktyWykresuFunkcji[311].X = 3.11
	punktyWykresuFunkcji[311].Y = 0.031

	punktyWykresuFunkcji[312].X = 3.12
	punktyWykresuFunkcji[312].Y = 0.021

	punktyWykresuFunkcji[313].X = 3.13
	punktyWykresuFunkcji[313].Y = 0.011

	punktyWykresuFunkcji[314].X = 3.14
	punktyWykresuFunkcji[314].Y = 0.001

	punktyWykresuFunkcji[315].X = 3.15
	punktyWykresuFunkcji[315].Y = -0.008

	punktyWykresuFunkcji[316].X = 3.16
	punktyWykresuFunkcji[316].Y = -0.018

	punktyWykresuFunkcji[317].X = 3.17
	punktyWykresuFunkcji[317].Y = -0.028

	punktyWykresuFunkcji[318].X = 3.18
	punktyWykresuFunkcji[318].Y = -0.038

	punktyWykresuFunkcji[319].X = 3.19
	punktyWykresuFunkcji[319].Y = -0.048

	punktyWykresuFunkcji[320].X = 3.2
	punktyWykresuFunkcji[320].Y = -0.058

	punktyWykresuFunkcji[321].X = 3.21
	punktyWykresuFunkcji[321].Y = -0.068

	punktyWykresuFunkcji[322].X = 3.22
	punktyWykresuFunkcji[322].Y = -0.078

	punktyWykresuFunkcji[323].X = 3.23
	punktyWykresuFunkcji[323].Y = -0.088

	punktyWykresuFunkcji[324].X = 3.24
	punktyWykresuFunkcji[324].Y = -0.098

	punktyWykresuFunkcji[325].X = 3.25
	punktyWykresuFunkcji[325].Y = -0.108

	punktyWykresuFunkcji[326].X = 3.26
	punktyWykresuFunkcji[326].Y = -0.118

	punktyWykresuFunkcji[327].X = 3.27
	punktyWykresuFunkcji[327].Y = -0.128

	punktyWykresuFunkcji[328].X = 3.28
	punktyWykresuFunkcji[328].Y = -0.138

	punktyWykresuFunkcji[329].X = 3.29
	punktyWykresuFunkcji[329].Y = -0.147

	punktyWykresuFunkcji[330].X = 3.3
	punktyWykresuFunkcji[330].Y = -0.157

	punktyWykresuFunkcji[331].X = 3.31
	punktyWykresuFunkcji[331].Y = -0.167

	punktyWykresuFunkcji[332].X = 3.32
	punktyWykresuFunkcji[332].Y = -0.177

	punktyWykresuFunkcji[333].X = 3.33
	punktyWykresuFunkcji[333].Y = -0.187

	punktyWykresuFunkcji[334].X = 3.34
	punktyWykresuFunkcji[334].Y = -0.197

	punktyWykresuFunkcji[335].X = 3.35
	punktyWykresuFunkcji[335].Y = -0.206

	punktyWykresuFunkcji[336].X = 3.36
	punktyWykresuFunkcji[336].Y = -0.216

	punktyWykresuFunkcji[337].X = 3.37
	punktyWykresuFunkcji[337].Y = -0.226

	punktyWykresuFunkcji[338].X = 3.38
	punktyWykresuFunkcji[338].Y = -0.236

	punktyWykresuFunkcji[339].X = 3.39
	punktyWykresuFunkcji[339].Y = -0.245

	punktyWykresuFunkcji[340].X = 3.4
	punktyWykresuFunkcji[340].Y = -0.255

	punktyWykresuFunkcji[341].X = 3.41
	punktyWykresuFunkcji[341].Y = -0.265

	punktyWykresuFunkcji[342].X = 3.42
	punktyWykresuFunkcji[342].Y = -0.274

	punktyWykresuFunkcji[343].X = 3.43
	punktyWykresuFunkcji[343].Y = -0.284

	punktyWykresuFunkcji[344].X = 3.44
	punktyWykresuFunkcji[344].Y = -0.294

	punktyWykresuFunkcji[345].X = 3.45
	punktyWykresuFunkcji[345].Y = -0.303

	punktyWykresuFunkcji[346].X = 3.46
	punktyWykresuFunkcji[346].Y = -0.313

	punktyWykresuFunkcji[347].X = 3.47
	punktyWykresuFunkcji[347].Y = -0.322

	punktyWykresuFunkcji[348].X = 3.48
	punktyWykresuFunkcji[348].Y = -0.332

	punktyWykresuFunkcji[349].X = 3.49
	punktyWykresuFunkcji[349].Y = -0.341

	punktyWykresuFunkcji[350].X = 3.5
	punktyWykresuFunkcji[350].Y = -0.35

	punktyWykresuFunkcji[351].X = 3.51
	punktyWykresuFunkcji[351].Y = -0.36

	punktyWykresuFunkcji[352].X = 3.52
	punktyWykresuFunkcji[352].Y = -0.369

	punktyWykresuFunkcji[353].X = 3.53
	punktyWykresuFunkcji[353].Y = -0.378

	punktyWykresuFunkcji[354].X = 3.54
	punktyWykresuFunkcji[354].Y = -0.388

	punktyWykresuFunkcji[355].X = 3.55
	punktyWykresuFunkcji[355].Y = -0.397

	punktyWykresuFunkcji[356].X = 3.56
	punktyWykresuFunkcji[356].Y = -0.406

	punktyWykresuFunkcji[357].X = 3.57
	punktyWykresuFunkcji[357].Y = -0.415

	punktyWykresuFunkcji[358].X = 3.58
	punktyWykresuFunkcji[358].Y = -0.424

	punktyWykresuFunkcji[359].X = 3.59
	punktyWykresuFunkcji[359].Y = -0.433

	punktyWykresuFunkcji[360].X = 3.6
	punktyWykresuFunkcji[360].Y = -0.442

	punktyWykresuFunkcji[361].X = 3.61
	punktyWykresuFunkcji[361].Y = -0.451

	punktyWykresuFunkcji[362].X = 3.62
	punktyWykresuFunkcji[362].Y = -0.46

	punktyWykresuFunkcji[363].X = 3.63
	punktyWykresuFunkcji[363].Y = -0.469

	punktyWykresuFunkcji[364].X = 3.64
	punktyWykresuFunkcji[364].Y = -0.478

	punktyWykresuFunkcji[365].X = 3.65
	punktyWykresuFunkcji[365].Y = -0.486

	punktyWykresuFunkcji[366].X = 3.66
	punktyWykresuFunkcji[366].Y = -0.495

	punktyWykresuFunkcji[367].X = 3.67
	punktyWykresuFunkcji[367].Y = -0.504

	punktyWykresuFunkcji[368].X = 3.68
	punktyWykresuFunkcji[368].Y = -0.512

	punktyWykresuFunkcji[369].X = 3.69
	punktyWykresuFunkcji[369].Y = -0.521

	punktyWykresuFunkcji[370].X = 3.7
	punktyWykresuFunkcji[370].Y = -0.529

	punktyWykresuFunkcji[371].X = 3.71
	punktyWykresuFunkcji[371].Y = -0.538

	punktyWykresuFunkcji[372].X = 3.72
	punktyWykresuFunkcji[372].Y = -0.546

	punktyWykresuFunkcji[373].X = 3.73
	punktyWykresuFunkcji[373].Y = -0.555

	punktyWykresuFunkcji[374].X = 3.74
	punktyWykresuFunkcji[374].Y = -0.563

	punktyWykresuFunkcji[375].X = 3.75
	punktyWykresuFunkcji[375].Y = -0.571

	punktyWykresuFunkcji[376].X = 3.76
	punktyWykresuFunkcji[376].Y = -0.579

	punktyWykresuFunkcji[377].X = 3.77
	punktyWykresuFunkcji[377].Y = -0.587

	punktyWykresuFunkcji[378].X = 3.78
	punktyWykresuFunkcji[378].Y = -0.595

	punktyWykresuFunkcji[379].X = 3.79
	punktyWykresuFunkcji[379].Y = -0.603

	punktyWykresuFunkcji[380].X = 3.8
	punktyWykresuFunkcji[380].Y = -0.611

	punktyWykresuFunkcji[381].X = 3.81
	punktyWykresuFunkcji[381].Y = -0.619

	punktyWykresuFunkcji[382].X = 3.82
	punktyWykresuFunkcji[382].Y = -0.627

	punktyWykresuFunkcji[383].X = 3.83
	punktyWykresuFunkcji[383].Y = -0.635

	punktyWykresuFunkcji[384].X = 3.84
	punktyWykresuFunkcji[384].Y = -0.643

	punktyWykresuFunkcji[385].X = 3.85
	punktyWykresuFunkcji[385].Y = -0.65

	punktyWykresuFunkcji[386].X = 3.86
	punktyWykresuFunkcji[386].Y = -0.658

	punktyWykresuFunkcji[387].X = 3.87
	punktyWykresuFunkcji[387].Y = -0.665

	punktyWykresuFunkcji[388].X = 3.88
	punktyWykresuFunkcji[388].Y = -0.673

	punktyWykresuFunkcji[389].X = 3.89
	punktyWykresuFunkcji[389].Y = -0.68

	punktyWykresuFunkcji[390].X = 3.9
	punktyWykresuFunkcji[390].Y = -0.687

	punktyWykresuFunkcji[391].X = 3.91
	punktyWykresuFunkcji[391].Y = -0.694

	punktyWykresuFunkcji[392].X = 3.92
	punktyWykresuFunkcji[392].Y = -0.702

	punktyWykresuFunkcji[393].X = 3.93
	punktyWykresuFunkcji[393].Y = -0.709

	punktyWykresuFunkcji[394].X = 3.94
	punktyWykresuFunkcji[394].Y = -0.716

	punktyWykresuFunkcji[395].X = 3.95
	punktyWykresuFunkcji[395].Y = -0.723

	punktyWykresuFunkcji[396].X = 3.96
	punktyWykresuFunkcji[396].Y = -0.73

	punktyWykresuFunkcji[397].X = 3.97
	punktyWykresuFunkcji[397].Y = -0.736

	punktyWykresuFunkcji[398].X = 3.98
	punktyWykresuFunkcji[398].Y = -0.743

	punktyWykresuFunkcji[399].X = 3.99
	punktyWykresuFunkcji[399].Y = -0.75

	punktyWykresuFunkcji[400].X = 4.0
	punktyWykresuFunkcji[400].Y = -0.756

	punktyWykresuFunkcji[401].X = 4.01
	punktyWykresuFunkcji[401].Y = -0.763

	punktyWykresuFunkcji[402].X = 4.02
	punktyWykresuFunkcji[402].Y = -0.769

	punktyWykresuFunkcji[403].X = 4.03
	punktyWykresuFunkcji[403].Y = -0.776

	punktyWykresuFunkcji[404].X = 4.04
	punktyWykresuFunkcji[404].Y = -0.782

	punktyWykresuFunkcji[405].X = 4.05
	punktyWykresuFunkcji[405].Y = -0.788

	punktyWykresuFunkcji[406].X = 4.06
	punktyWykresuFunkcji[406].Y = -0.793

	punktyWykresuFunkcji[407].X = 4.07
	punktyWykresuFunkcji[407].Y = -0.8

	punktyWykresuFunkcji[408].X = 4.08
	punktyWykresuFunkcji[408].Y = -0.806

	punktyWykresuFunkcji[409].X = 4.09
	punktyWykresuFunkcji[409].Y = -0.812

	punktyWykresuFunkcji[410].X = 4.1
	punktyWykresuFunkcji[410].Y = -0.818

	punktyWykresuFunkcji[411].X = 4.11
	punktyWykresuFunkcji[411].Y = -0.824

	punktyWykresuFunkcji[412].X = 4.12
	punktyWykresuFunkcji[412].Y = -0.829

	punktyWykresuFunkcji[413].X = 4.13
	punktyWykresuFunkcji[413].Y = -0.835

	punktyWykresuFunkcji[414].X = 4.14
	punktyWykresuFunkcji[414].Y = -0.84

	punktyWykresuFunkcji[415].X = 4.15
	punktyWykresuFunkcji[415].Y = -0.846

	punktyWykresuFunkcji[416].X = 4.16
	punktyWykresuFunkcji[416].Y = -0.851

	punktyWykresuFunkcji[417].X = 4.17
	punktyWykresuFunkcji[417].Y = -0.856

	punktyWykresuFunkcji[418].X = 4.18
	punktyWykresuFunkcji[418].Y = -0.861

	punktyWykresuFunkcji[419].X = 4.19
	punktyWykresuFunkcji[419].Y = -0.866

	punktyWykresuFunkcji[420].X = 4.2
	punktyWykresuFunkcji[420].Y = -0.871

	punktyWykresuFunkcji[421].X = 4.21
	punktyWykresuFunkcji[421].Y = -0.874

	punktyWykresuFunkcji[422].X = 4.22
	punktyWykresuFunkcji[422].Y = -0.881

	punktyWykresuFunkcji[423].X = 4.23
	punktyWykresuFunkcji[423].Y = -0.885

	punktyWykresuFunkcji[424].X = 4.24
	punktyWykresuFunkcji[424].Y = -0.89

	punktyWykresuFunkcji[425].X = 4.25
	punktyWykresuFunkcji[425].Y = -0.895

	punktyWykresuFunkcji[426].X = 4.26
	punktyWykresuFunkcji[426].Y = -0.899

	punktyWykresuFunkcji[427].X = 4.27
	punktyWykresuFunkcji[427].Y = -0.903

	punktyWykresuFunkcji[428].X = 4.28
	punktyWykresuFunkcji[428].Y = -0.908

	punktyWykresuFunkcji[429].X = 4.29
	punktyWykresuFunkcji[429].Y = -0.912

	punktyWykresuFunkcji[430].X = 4.3
	punktyWykresuFunkcji[430].Y = -0.916

	punktyWykresuFunkcji[431].X = 4.31
	punktyWykresuFunkcji[431].Y = -0.92

	punktyWykresuFunkcji[432].X = 4.32
	punktyWykresuFunkcji[432].Y = -0.924

	punktyWykresuFunkcji[433].X = 4.33
	punktyWykresuFunkcji[433].Y = -0.927

	punktyWykresuFunkcji[434].X = 4.34
	punktyWykresuFunkcji[434].Y = -0.931

	punktyWykresuFunkcji[435].X = 4.35
	punktyWykresuFunkcji[435].Y = -0.935

	punktyWykresuFunkcji[436].X = 4.36
	punktyWykresuFunkcji[436].Y = -0.938

	punktyWykresuFunkcji[437].X = 4.37
	punktyWykresuFunkcji[437].Y = -0.942

	punktyWykresuFunkcji[438].X = 4.38
	punktyWykresuFunkcji[438].Y = -0.945

	punktyWykresuFunkcji[439].X = 4.39
	punktyWykresuFunkcji[439].Y = -0.948

	punktyWykresuFunkcji[440].X = 4.4
	punktyWykresuFunkcji[440].Y = -0.951

	punktyWykresuFunkcji[441].X = 4.41
	punktyWykresuFunkcji[441].Y = -0.954

	punktyWykresuFunkcji[442].X = 4.42
	punktyWykresuFunkcji[442].Y = -0.957

	punktyWykresuFunkcji[443].X = 4.43
	punktyWykresuFunkcji[443].Y = -0.96

	punktyWykresuFunkcji[444].X = 4.44
	punktyWykresuFunkcji[444].Y = -0.963

	punktyWykresuFunkcji[445].X = 4.45
	punktyWykresuFunkcji[445].Y = -0.965

	punktyWykresuFunkcji[446].X = 4.46
	punktyWykresuFunkcji[446].Y = -0.968

	punktyWykresuFunkcji[447].X = 4.47
	punktyWykresuFunkcji[447].Y = -0.97

	punktyWykresuFunkcji[448].X = 4.48
	punktyWykresuFunkcji[448].Y = -0.973

	punktyWykresuFunkcji[449].X = 4.49
	punktyWykresuFunkcji[449].Y = -0.975

	punktyWykresuFunkcji[450].X = 4.5
	punktyWykresuFunkcji[450].Y = -0.977

	punktyWykresuFunkcji[451].X = 4.51
	punktyWykresuFunkcji[451].Y = -0.979

	punktyWykresuFunkcji[452].X = 4.52
	punktyWykresuFunkcji[452].Y = -0.981

	punktyWykresuFunkcji[453].X = 4.53
	punktyWykresuFunkcji[453].Y = -0.983

	punktyWykresuFunkcji[454].X = 4.54
	punktyWykresuFunkcji[454].Y = -0.985

	punktyWykresuFunkcji[455].X = 4.55
	punktyWykresuFunkcji[455].Y = -0.986

	punktyWykresuFunkcji[456].X = 4.56
	punktyWykresuFunkcji[456].Y = -0.988

	punktyWykresuFunkcji[457].X = 4.57
	punktyWykresuFunkcji[457].Y = -0.989

	punktyWykresuFunkcji[458].X = 4.58
	punktyWykresuFunkcji[458].Y = -0.991

	punktyWykresuFunkcji[459].X = 4.59
	punktyWykresuFunkcji[459].Y = -0.992

	punktyWykresuFunkcji[460].X = 4.6
	punktyWykresuFunkcji[460].Y = -0.993

	punktyWykresuFunkcji[461].X = 4.61
	punktyWykresuFunkcji[461].Y = -0.994

	punktyWykresuFunkcji[462].X = 4.62
	punktyWykresuFunkcji[462].Y = -0.995

	punktyWykresuFunkcji[463].X = 4.63
	punktyWykresuFunkcji[463].Y = -0.996

	punktyWykresuFunkcji[464].X = 4.64
	punktyWykresuFunkcji[464].Y = -0.997

	punktyWykresuFunkcji[465].X = 4.65
	punktyWykresuFunkcji[465].Y = -0.998

	punktyWykresuFunkcji[466].X = 4.66
	punktyWykresuFunkcji[466].Y = -0.998

	punktyWykresuFunkcji[467].X = 4.67
	punktyWykresuFunkcji[467].Y = -0.999

	punktyWykresuFunkcji[468].X = 4.68
	punktyWykresuFunkcji[468].Y = -0.999

	punktyWykresuFunkcji[469].X = 4.69
	punktyWykresuFunkcji[469].Y = -0.999

	punktyWykresuFunkcji[470].X = 4.7
	punktyWykresuFunkcji[470].Y = -0.999

	punktyWykresuFunkcji[471].X = 4.71
	punktyWykresuFunkcji[471].Y = -1.0

	punktyWykresuFunkcji[472].X = 4.72
	punktyWykresuFunkcji[472].Y = -1.0

	punktyWykresuFunkcji[473].X = 4.73
	punktyWykresuFunkcji[473].Y = -0.999

	punktyWykresuFunkcji[474].X = 4.74
	punktyWykresuFunkcji[474].Y = -0.999

	punktyWykresuFunkcji[475].X = 4.75
	punktyWykresuFunkcji[475].Y = -0.999

	punktyWykresuFunkcji[476].X = 4.76
	punktyWykresuFunkcji[476].Y = -0.998

	punktyWykresuFunkcji[477].X = 4.77
	punktyWykresuFunkcji[477].Y = -0.998

	punktyWykresuFunkcji[478].X = 4.78
	punktyWykresuFunkcji[478].Y = -0.997

	punktyWykresuFunkcji[479].X = 4.79
	punktyWykresuFunkcji[479].Y = -0.997

	punktyWykresuFunkcji[480].X = 4.8
	punktyWykresuFunkcji[480].Y = -0.996

	punktyWykresuFunkcji[481].X = 4.81
	punktyWykresuFunkcji[481].Y = -0.995

	punktyWykresuFunkcji[482].X = 4.82
	punktyWykresuFunkcji[482].Y = -0.994

	punktyWykresuFunkcji[483].X = 4.83
	punktyWykresuFunkcji[483].Y = -0.993

	punktyWykresuFunkcji[484].X = 4.84
	punktyWykresuFunkcji[484].Y = -0.991

	punktyWykresuFunkcji[485].X = 4.85
	punktyWykresuFunkcji[485].Y = -0.99

	punktyWykresuFunkcji[486].X = 4.86
	punktyWykresuFunkcji[486].Y = -0.989

	punktyWykresuFunkcji[487].X = 4.87
	punktyWykresuFunkcji[487].Y = -0.987

	punktyWykresuFunkcji[488].X = 4.88
	punktyWykresuFunkcji[488].Y = -0.986

	punktyWykresuFunkcji[489].X = 4.89
	punktyWykresuFunkcji[489].Y = -0.984

	punktyWykresuFunkcji[490].X = 4.9
	punktyWykresuFunkcji[490].Y = -0.982

	punktyWykresuFunkcji[491].X = 4.91
	punktyWykresuFunkcji[491].Y = -0.98

	punktyWykresuFunkcji[492].X = 4.92
	punktyWykresuFunkcji[492].Y = -0.978

	punktyWykresuFunkcji[493].X = 4.93
	punktyWykresuFunkcji[493].Y = -0.976

	punktyWykresuFunkcji[494].X = 4.94
	punktyWykresuFunkcji[494].Y = -0.974

	punktyWykresuFunkcji[495].X = 4.95
	punktyWykresuFunkcji[495].Y = -0.971

	punktyWykresuFunkcji[496].X = 4.96
	punktyWykresuFunkcji[496].Y = -0.969

	punktyWykresuFunkcji[497].X = 4.97
	punktyWykresuFunkcji[497].Y = -0.967

	punktyWykresuFunkcji[498].X = 4.98
	punktyWykresuFunkcji[498].Y = -0.964

	punktyWykresuFunkcji[499].X = 4.99
	punktyWykresuFunkcji[499].Y = -0.961

	punktyWykresuFunkcji[500].X = 5.0
	punktyWykresuFunkcji[500].Y = -0.958

	punktyWykresuFunkcji[501].X = 5.01
	punktyWykresuFunkcji[501].Y = -0.956

	punktyWykresuFunkcji[502].X = 5.02
	punktyWykresuFunkcji[502].Y = -0.953

	punktyWykresuFunkcji[503].X = 5.03
	punktyWykresuFunkcji[503].Y = -0.95

	punktyWykresuFunkcji[504].X = 5.04
	punktyWykresuFunkcji[504].Y = -0.946

	punktyWykresuFunkcji[505].X = 5.05
	punktyWykresuFunkcji[505].Y = -0.943

	punktyWykresuFunkcji[506].X = 5.06
	punktyWykresuFunkcji[506].Y = -0.94

	punktyWykresuFunkcji[507].X = 5.07
	punktyWykresuFunkcji[507].Y = -0.936

	punktyWykresuFunkcji[508].X = 5.08
	punktyWykresuFunkcji[508].Y = -0.933

	punktyWykresuFunkcji[509].X = 5.09
	punktyWykresuFunkcji[509].Y = -0.929

	punktyWykresuFunkcji[510].X = 5.1
	punktyWykresuFunkcji[510].Y = -0.925

	punktyWykresuFunkcji[511].X = 5.11
	punktyWykresuFunkcji[511].Y = -0.922

	punktyWykresuFunkcji[512].X = 5.12
	punktyWykresuFunkcji[512].Y = -0.918

	punktyWykresuFunkcji[513].X = 5.13
	punktyWykresuFunkcji[513].Y = -0.914

	punktyWykresuFunkcji[514].X = 5.14
	punktyWykresuFunkcji[514].Y = -0.91

	punktyWykresuFunkcji[515].X = 5.15
	punktyWykresuFunkcji[515].Y = -0.905

	punktyWykresuFunkcji[516].X = 5.16
	punktyWykresuFunkcji[516].Y = -0.901

	punktyWykresuFunkcji[517].X = 5.17
	punktyWykresuFunkcji[517].Y = -0.897

	punktyWykresuFunkcji[518].X = 5.18
	punktyWykresuFunkcji[518].Y = -0.892

	punktyWykresuFunkcji[519].X = 5.19
	punktyWykresuFunkcji[519].Y = -0.888

	punktyWykresuFunkcji[520].X = 5.2
	punktyWykresuFunkcji[520].Y = -0.883

	punktyWykresuFunkcji[521].X = 5.21
	punktyWykresuFunkcji[521].Y = -0.878

	punktyWykresuFunkcji[522].X = 5.22
	punktyWykresuFunkcji[522].Y = -0.873

	punktyWykresuFunkcji[523].X = 5.23
	punktyWykresuFunkcji[523].Y = -0.869

	punktyWykresuFunkcji[524].X = 5.24
	punktyWykresuFunkcji[524].Y = -0.864

	punktyWykresuFunkcji[525].X = 5.25
	punktyWykresuFunkcji[525].Y = -0.858

	punktyWykresuFunkcji[526].X = 5.26
	punktyWykresuFunkcji[526].Y = -0.853

	punktyWykresuFunkcji[527].X = 5.27
	punktyWykresuFunkcji[527].Y = -0.848

	punktyWykresuFunkcji[528].X = 5.28
	punktyWykresuFunkcji[528].Y = -0.843

	punktyWykresuFunkcji[529].X = 5.29
	punktyWykresuFunkcji[529].Y = -0.837

	punktyWykresuFunkcji[530].X = 5.3
	punktyWykresuFunkcji[530].Y = -0.832

	punktyWykresuFunkcji[531].X = 5.31
	punktyWykresuFunkcji[531].Y = -0.826

	punktyWykresuFunkcji[532].X = 5.32
	punktyWykresuFunkcji[532].Y = -0.821

	punktyWykresuFunkcji[533].X = 5.33
	punktyWykresuFunkcji[533].Y = -0.815

	punktyWykresuFunkcji[534].X = 5.34
	punktyWykresuFunkcji[534].Y = -0.809

	punktyWykresuFunkcji[535].X = 5.35
	punktyWykresuFunkcji[535].Y = -0.803

	punktyWykresuFunkcji[536].X = 5.36
	punktyWykresuFunkcji[536].Y = -0.797

	punktyWykresuFunkcji[537].X = 5.37
	punktyWykresuFunkcji[537].Y = -0.791

	punktyWykresuFunkcji[538].X = 5.38
	punktyWykresuFunkcji[538].Y = -0.785

	punktyWykresuFunkcji[539].X = 5.39
	punktyWykresuFunkcji[539].Y = -0.779

	punktyWykresuFunkcji[540].X = 5.4
	punktyWykresuFunkcji[540].Y = -0.772

	punktyWykresuFunkcji[541].X = 5.41
	punktyWykresuFunkcji[541].Y = -0.766

	punktyWykresuFunkcji[542].X = 5.42
	punktyWykresuFunkcji[542].Y = -0.759

	punktyWykresuFunkcji[543].X = 5.43
	punktyWykresuFunkcji[543].Y = -0.753

	punktyWykresuFunkcji[544].X = 5.44
	punktyWykresuFunkcji[544].Y = -0.746

	punktyWykresuFunkcji[545].X = 5.45
	punktyWykresuFunkcji[545].Y = -0.74

	punktyWykresuFunkcji[546].X = 5.46
	punktyWykresuFunkcji[546].Y = -0.733

	punktyWykresuFunkcji[547].X = 5.47
	punktyWykresuFunkcji[547].Y = -0.726

	punktyWykresuFunkcji[548].X = 5.48
	punktyWykresuFunkcji[548].Y = -0.719

	punktyWykresuFunkcji[549].X = 5.49
	punktyWykresuFunkcji[549].Y = -0.712

	punktyWykresuFunkcji[550].X = 5.5
	punktyWykresuFunkcji[550].Y = -0.705

	punktyWykresuFunkcji[551].X = 5.51
	punktyWykresuFunkcji[551].Y = -0.698

	punktyWykresuFunkcji[552].X = 5.52
	punktyWykresuFunkcji[552].Y = -0.691

	punktyWykresuFunkcji[553].X = 5.53
	punktyWykresuFunkcji[553].Y = -0.684

	punktyWykresuFunkcji[554].X = 5.54
	punktyWykresuFunkcji[554].Y = -0.676

	punktyWykresuFunkcji[555].X = 5.55
	punktyWykresuFunkcji[555].Y = -0.669

	punktyWykresuFunkcji[556].X = 5.56
	punktyWykresuFunkcji[556].Y = -0.661

	punktyWykresuFunkcji[557].X = 5.57
	punktyWykresuFunkcji[557].Y = -0.654

	punktyWykresuFunkcji[558].X = 5.58
	punktyWykresuFunkcji[558].Y = -0.646

	punktyWykresuFunkcji[559].X = 5.59
	punktyWykresuFunkcji[559].Y = -0.639

	punktyWykresuFunkcji[560].X = 5.6
	punktyWykresuFunkcji[560].Y = -0.631

	punktyWykresuFunkcji[561].X = 5.61
	punktyWykresuFunkcji[561].Y = -0.623

	punktyWykresuFunkcji[562].X = 5.62
	punktyWykresuFunkcji[562].Y = -0.615

	punktyWykresuFunkcji[563].X = 5.63
	punktyWykresuFunkcji[563].Y = -0.607

	punktyWykresuFunkcji[564].X = 5.64
	punktyWykresuFunkcji[564].Y = -0.599

	punktyWykresuFunkcji[565].X = 5.65
	punktyWykresuFunkcji[565].Y = -0.591

	punktyWykresuFunkcji[566].X = 5.66
	punktyWykresuFunkcji[566].Y = -0.583

	punktyWykresuFunkcji[567].X = 5.67
	punktyWykresuFunkcji[567].Y = -0.575

	punktyWykresuFunkcji[568].X = 5.68
	punktyWykresuFunkcji[568].Y = -0.567

	punktyWykresuFunkcji[569].X = 5.69
	punktyWykresuFunkcji[569].Y = -0.559

	punktyWykresuFunkcji[570].X = 5.7
	punktyWykresuFunkcji[570].Y = -0.55

	punktyWykresuFunkcji[571].X = 5.71
	punktyWykresuFunkcji[571].Y = -0.542

	punktyWykresuFunkcji[572].X = 5.72
	punktyWykresuFunkcji[572].Y = -0.533

	punktyWykresuFunkcji[573].X = 5.73
	punktyWykresuFunkcji[573].Y = -0.525

	punktyWykresuFunkcji[574].X = 5.74
	punktyWykresuFunkcji[574].Y = -0.516

	punktyWykresuFunkcji[575].X = 5.75
	punktyWykresuFunkcji[575].Y = -0.508

	punktyWykresuFunkcji[576].X = 5.76
	punktyWykresuFunkcji[576].Y = -0.499

	punktyWykresuFunkcji[577].X = 5.77
	punktyWykresuFunkcji[577].Y = -0.491

	punktyWykresuFunkcji[578].X = 5.78
	punktyWykresuFunkcji[578].Y = -0.482

	punktyWykresuFunkcji[579].X = 5.79
	punktyWykresuFunkcji[579].Y = -0.473

	punktyWykresuFunkcji[580].X = 5.8
	punktyWykresuFunkcji[580].Y = -0.464

	punktyWykresuFunkcji[581].X = 5.81
	punktyWykresuFunkcji[581].Y = -0.455

	punktyWykresuFunkcji[582].X = 5.82
	punktyWykresuFunkcji[582].Y = -0.446

	punktyWykresuFunkcji[583].X = 5.83
	punktyWykresuFunkcji[583].Y = -0.437

	punktyWykresuFunkcji[584].X = 5.84
	punktyWykresuFunkcji[584].Y = -0.428

	punktyWykresuFunkcji[585].X = 5.85
	punktyWykresuFunkcji[585].Y = -0.419

	punktyWykresuFunkcji[586].X = 5.86
	punktyWykresuFunkcji[586].Y = -0.41

	punktyWykresuFunkcji[587].X = 5.87
	punktyWykresuFunkcji[587].Y = -0.401

	punktyWykresuFunkcji[588].X = 5.88
	punktyWykresuFunkcji[588].Y = -0.392

	punktyWykresuFunkcji[589].X = 5.89
	punktyWykresuFunkcji[589].Y = -0.383

	punktyWykresuFunkcji[590].X = 5.9
	punktyWykresuFunkcji[590].Y = -0.373

	punktyWykresuFunkcji[591].X = 5.91
	punktyWykresuFunkcji[591].Y = -0.364

	punktyWykresuFunkcji[592].X = 5.92
	punktyWykresuFunkcji[592].Y = -0.355

	punktyWykresuFunkcji[593].X = 5.93
	punktyWykresuFunkcji[593].Y = -0.345

	punktyWykresuFunkcji[594].X = 5.94
	punktyWykresuFunkcji[594].Y = -0.336

	punktyWykresuFunkcji[595].X = 5.95
	punktyWykresuFunkcji[595].Y = -0.327

	punktyWykresuFunkcji[596].X = 5.96
	punktyWykresuFunkcji[596].Y = -0.317

	punktyWykresuFunkcji[597].X = 5.97
	punktyWykresuFunkcji[597].Y = -0.308

	punktyWykresuFunkcji[598].X = 5.98
	punktyWykresuFunkcji[598].Y = -0.298

	punktyWykresuFunkcji[599].X = 5.99
	punktyWykresuFunkcji[599].Y = -0.289

	punktyWykresuFunkcji[600].X = 6.0
	punktyWykresuFunkcji[600].Y = -0.279

	punktyWykresuFunkcji[601].X = 6.01
	punktyWykresuFunkcji[601].Y = -0.269

	punktyWykresuFunkcji[602].X = 6.02
	punktyWykresuFunkcji[602].Y = -0.26

	punktyWykresuFunkcji[603].X = 6.03
	punktyWykresuFunkcji[603].Y = -0.25

	punktyWykresuFunkcji[604].X = 6.04
	punktyWykresuFunkcji[604].Y = -0.24

	punktyWykresuFunkcji[605].X = 6.05
	punktyWykresuFunkcji[605].Y = -0.231

	punktyWykresuFunkcji[606].X = 6.06
	punktyWykresuFunkcji[606].Y = -0.221

	punktyWykresuFunkcji[607].X = 6.07
	punktyWykresuFunkcji[607].Y = -0.211

	punktyWykresuFunkcji[608].X = 6.08
	punktyWykresuFunkcji[608].Y = -0.201

	punktyWykresuFunkcji[609].X = 6.09
	punktyWykresuFunkcji[609].Y = -0.192

	punktyWykresuFunkcji[610].X = 6.1
	punktyWykresuFunkcji[610].Y = -0.182

	punktyWykresuFunkcji[611].X = 6.11
	punktyWykresuFunkcji[611].Y = -0.172

	punktyWykresuFunkcji[612].X = 6.12
	punktyWykresuFunkcji[612].Y = -0.162

	punktyWykresuFunkcji[613].X = 6.13
	punktyWykresuFunkcji[613].Y = -0.152

	punktyWykresuFunkcji[614].X = 6.14
	punktyWykresuFunkcji[614].Y = -0.142

	punktyWykresuFunkcji[615].X = 6.15
	punktyWykresuFunkcji[615].Y = -0.132

	punktyWykresuFunkcji[616].X = 6.16
	punktyWykresuFunkcji[616].Y = -0.122

	punktyWykresuFunkcji[617].X = 6.17
	punktyWykresuFunkcji[617].Y = -0.112

	punktyWykresuFunkcji[618].X = 6.18
	punktyWykresuFunkcji[618].Y = -0.103

	punktyWykresuFunkcji[619].X = 6.19
	punktyWykresuFunkcji[619].Y = -0.093

	punktyWykresuFunkcji[620].X = 6.2
	punktyWykresuFunkcji[620].Y = -0.083

	punktyWykresuFunkcji[621].X = 6.21
	punktyWykresuFunkcji[621].Y = -0.073

	punktyWykresuFunkcji[622].X = 6.22
	punktyWykresuFunkcji[622].Y = -0.063

	punktyWykresuFunkcji[623].X = 6.23
	punktyWykresuFunkcji[623].Y = -0.053

	punktyWykresuFunkcji[624].X = 6.24
	punktyWykresuFunkcji[624].Y = -0.043

	punktyWykresuFunkcji[625].X = 6.25
	punktyWykresuFunkcji[625].Y = -0.033

	punktyWykresuFunkcji[626].X = 6.26
	punktyWykresuFunkcji[626].Y = -0.023

	punktyWykresuFunkcji[627].X = 6.27
	punktyWykresuFunkcji[627].Y = -0.013

	punktyWykresuFunkcji[628].X = 6.28
	punktyWykresuFunkcji[628].Y = -0.003

	punktyWykresuFunkcji[629].X = 6.29
	punktyWykresuFunkcji[629].Y = 0.006

	punktyWykresuFunkcji[630].X = 6.3
	punktyWykresuFunkcji[630].Y = 0.016





	wykresFunkcji := plot.New()

	wykresFunkcji.Title.Text = "Wykres funkcji sin(x)"

	wykresFunkcji.X.Label.Text = "x"
	wykresFunkcji.Y.Label.Text = "y"

	liniaWykresu, err := plotter.NewLine(punktyWykresuFunkcji)

	if err != nil {
		panic(err)
	}

	liniaWykresu.LineStyle.Width = vg.Points(0.1)
	liniaWykresu.Color = color.RGBA{R: 200, G: 100, B: 100}

	wykresFunkcji.Add(liniaWykresu)
	wykresFunkcji.Legend.Add("sin(x)", liniaWykresu)

	if err := wykresFunkcji.Save(10*vg.Inch, 10*vg.Inch,
		"sin-plot-01.png"); err != nil {

		panic(err)
	}
}
