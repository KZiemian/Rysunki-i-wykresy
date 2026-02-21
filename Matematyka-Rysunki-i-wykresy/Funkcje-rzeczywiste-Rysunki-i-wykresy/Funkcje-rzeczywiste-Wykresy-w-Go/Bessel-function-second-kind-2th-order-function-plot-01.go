package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function J0(x).

	pointsOfFunctionPlot := make(plotter.XYs, 2_001)

	pointsOfFunctionPlot[0].X = 0.0
	pointsOfFunctionPlot[0].Y = 1.0

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = 0.999

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = 0.999

	pointsOfFunctionPlot[3].X = 0.03
	pointsOfFunctionPlot[3].Y = 0.999

	pointsOfFunctionPlot[4].X = 0.04
	pointsOfFunctionPlot[4].Y = 0.999

	pointsOfFunctionPlot[5].X = 0.05
	pointsOfFunctionPlot[5].Y = 0.999

	pointsOfFunctionPlot[6].X = 0.06
	pointsOfFunctionPlot[6].Y = 0.999

	pointsOfFunctionPlot[7].X = 0.07
	pointsOfFunctionPlot[7].Y = 0.998

	pointsOfFunctionPlot[8].X = 0.08
	pointsOfFunctionPlot[8].Y = 0.998

	pointsOfFunctionPlot[9].X = 0.09
	pointsOfFunctionPlot[9].Y = 0.997

	pointsOfFunctionPlot[10].X = 0.10
	pointsOfFunctionPlot[10].Y = 0.997

	pointsOfFunctionPlot[11].X = 0.11
	pointsOfFunctionPlot[11].Y = 0.996

	pointsOfFunctionPlot[12].X = 0.12
	pointsOfFunctionPlot[12].Y = 0.996

	pointsOfFunctionPlot[13].X = 0.13
	pointsOfFunctionPlot[13].Y = 0.995

	pointsOfFunctionPlot[14].X = 0.14
	pointsOfFunctionPlot[14].Y = 0.995

	pointsOfFunctionPlot[15].X = 0.15
	pointsOfFunctionPlot[15].Y = 0.994

	pointsOfFunctionPlot[16].X = 0.16
	pointsOfFunctionPlot[16].Y = 0.993

	pointsOfFunctionPlot[17].X = 0.17
	pointsOfFunctionPlot[17].Y = 0.992

	pointsOfFunctionPlot[18].X = 0.18
	pointsOfFunctionPlot[18].Y = 0.991

	pointsOfFunctionPlot[19].X = 0.19
	pointsOfFunctionPlot[19].Y = 0.99

	pointsOfFunctionPlot[20].X = 0.20
	pointsOfFunctionPlot[20].Y = 0.99

	pointsOfFunctionPlot[21].X = 0.21
	pointsOfFunctionPlot[21].Y = 0.989

	pointsOfFunctionPlot[22].X = 0.22
	pointsOfFunctionPlot[22].Y = 0.987

	pointsOfFunctionPlot[23].X = 0.23
	pointsOfFunctionPlot[23].Y = 0.986

	pointsOfFunctionPlot[24].X = 0.24
	pointsOfFunctionPlot[24].Y = 0.985

	pointsOfFunctionPlot[25].X = 0.25
	pointsOfFunctionPlot[25].Y = 0.984

	pointsOfFunctionPlot[26].X = 0.26
	pointsOfFunctionPlot[26].Y = 0.983

	pointsOfFunctionPlot[27].X = 0.27
	pointsOfFunctionPlot[27].Y = 0.981

	pointsOfFunctionPlot[28].X = 0.28
	pointsOfFunctionPlot[28].Y = 0.98

	pointsOfFunctionPlot[29].X = 0.29
	pointsOfFunctionPlot[29].Y = 0.979

	pointsOfFunctionPlot[30].X = 0.30
	pointsOfFunctionPlot[30].Y = 0.977

	pointsOfFunctionPlot[31].X = 0.31
	pointsOfFunctionPlot[31].Y = 0.976

	pointsOfFunctionPlot[32].X = 0.32
	pointsOfFunctionPlot[32].Y = 0.974

	pointsOfFunctionPlot[33].X = 0.33
	pointsOfFunctionPlot[33].Y = 0.972

	pointsOfFunctionPlot[34].X = 0.34
	pointsOfFunctionPlot[34].Y = 0.971

	pointsOfFunctionPlot[35].X = 0.35
	pointsOfFunctionPlot[35].Y = 0.969

	pointsOfFunctionPlot[36].X = 0.36
	pointsOfFunctionPlot[36].Y = 0.967

	pointsOfFunctionPlot[37].X = 0.37
	pointsOfFunctionPlot[37].Y = 0.966

	pointsOfFunctionPlot[38].X = 0.38
	pointsOfFunctionPlot[38].Y = 0.964

	pointsOfFunctionPlot[39].X = 0.39
	pointsOfFunctionPlot[39].Y = 0.962

	pointsOfFunctionPlot[40].X = 0.40
	pointsOfFunctionPlot[40].Y = 0.96

	pointsOfFunctionPlot[41].X = 0.41
	pointsOfFunctionPlot[41].Y = 0.958

	pointsOfFunctionPlot[42].X = 0.42
	pointsOfFunctionPlot[42].Y = 0.956

	pointsOfFunctionPlot[43].X = 0.43
	pointsOfFunctionPlot[43].Y = 0.954

	pointsOfFunctionPlot[44].X = 0.44
	pointsOfFunctionPlot[44].Y = 0.952

	pointsOfFunctionPlot[45].X = 0.45
	pointsOfFunctionPlot[45].Y = 0.95

	pointsOfFunctionPlot[46].X = 0.46
	pointsOfFunctionPlot[46].Y = 0.947

	pointsOfFunctionPlot[47].X = 0.47
	pointsOfFunctionPlot[47].Y = 0.945

	pointsOfFunctionPlot[48].X = 0.48
	pointsOfFunctionPlot[48].Y = 0.943

	pointsOfFunctionPlot[49].X = 0.49
	pointsOfFunctionPlot[49].Y = 0.94

	pointsOfFunctionPlot[50].X = 0.50
	pointsOfFunctionPlot[50].Y = 0.938

	pointsOfFunctionPlot[51].X = 0.51
	pointsOfFunctionPlot[51].Y = 0.936

	pointsOfFunctionPlot[52].X = 0.52
	pointsOfFunctionPlot[52].Y = 0.933

	pointsOfFunctionPlot[53].X = 0.53
	pointsOfFunctionPlot[53].Y = 0.93

	pointsOfFunctionPlot[54].X = 0.54
	pointsOfFunctionPlot[54].Y = 0.928

	pointsOfFunctionPlot[55].X = 0.55
	pointsOfFunctionPlot[55].Y = 0.925

	pointsOfFunctionPlot[56].X = 0.56
	pointsOfFunctionPlot[56].Y = 0.923

	pointsOfFunctionPlot[57].X = 0.57
	pointsOfFunctionPlot[57].Y = 0.92

	pointsOfFunctionPlot[58].X = 0.58
	pointsOfFunctionPlot[58].Y = 0.917

	pointsOfFunctionPlot[59].X = 0.59
	pointsOfFunctionPlot[59].Y = 0.914

	pointsOfFunctionPlot[60].X = 0.60
	pointsOfFunctionPlot[60].Y = 0.912

	pointsOfFunctionPlot[61].X = 0.61
	pointsOfFunctionPlot[61].Y = 0.909

	pointsOfFunctionPlot[62].X = 0.62
	pointsOfFunctionPlot[62].Y = 0.906

	pointsOfFunctionPlot[63].X = 0.63
	pointsOfFunctionPlot[63].Y = 0.903

	pointsOfFunctionPlot[64].X = 0.64
	pointsOfFunctionPlot[64].Y = 0.9

	pointsOfFunctionPlot[65].X = 0.65
	pointsOfFunctionPlot[65].Y = 0.897

	pointsOfFunctionPlot[66].X = 0.66
	pointsOfFunctionPlot[66].Y = 0.894

	pointsOfFunctionPlot[67].X = 0.67
	pointsOfFunctionPlot[67].Y = 0.89

	pointsOfFunctionPlot[68].X = 0.68
	pointsOfFunctionPlot[68].Y = 0.887

	pointsOfFunctionPlot[69].X = 0.69
	pointsOfFunctionPlot[69].Y = 0.884

	pointsOfFunctionPlot[70].X = 0.70
	pointsOfFunctionPlot[70].Y = 0.881

	pointsOfFunctionPlot[71].X = 0.71
	pointsOfFunctionPlot[71].Y = 0.877

	pointsOfFunctionPlot[72].X = 0.72
	pointsOfFunctionPlot[72].Y = 0.874

	pointsOfFunctionPlot[73].X = 0.73
	pointsOfFunctionPlot[73].Y = 0.871

	pointsOfFunctionPlot[74].X = 0.74
	pointsOfFunctionPlot[74].Y = 0.867

	pointsOfFunctionPlot[75].X = 0.75
	pointsOfFunctionPlot[75].Y = 0.864

	pointsOfFunctionPlot[76].X = 0.76
	pointsOfFunctionPlot[76].Y = 0.86

	pointsOfFunctionPlot[77].X = 0.77
	pointsOfFunctionPlot[77].Y = 0.857

	pointsOfFunctionPlot[78].X = 0.78
	pointsOfFunctionPlot[78].Y = 0.853

	pointsOfFunctionPlot[79].X = 0.79
	pointsOfFunctionPlot[79].Y = 0.849

	pointsOfFunctionPlot[80].X = 0.80
	pointsOfFunctionPlot[80].Y = 0.846

	pointsOfFunctionPlot[81].X = 0.81
	pointsOfFunctionPlot[81].Y = 0.842

	pointsOfFunctionPlot[82].X = 0.82
	pointsOfFunctionPlot[82].Y = 0.838

	pointsOfFunctionPlot[83].X = 0.83
	pointsOfFunctionPlot[83].Y = 0.835

	pointsOfFunctionPlot[84].X = 0.84
	pointsOfFunctionPlot[84].Y = 0.831

	pointsOfFunctionPlot[85].X = 0.85
	pointsOfFunctionPlot[85].Y = 0.827

	pointsOfFunctionPlot[86].X = 0.86
	pointsOfFunctionPlot[86].Y = 0.823

	pointsOfFunctionPlot[87].X = 0.87
	pointsOfFunctionPlot[87].Y = 0.819

	pointsOfFunctionPlot[88].X = 0.88
	pointsOfFunctionPlot[88].Y = 0.815

	pointsOfFunctionPlot[89].X = 0.89
	pointsOfFunctionPlot[89].Y = 0.811

	pointsOfFunctionPlot[90].X = 0.90
	pointsOfFunctionPlot[90].Y = 0.807

	pointsOfFunctionPlot[91].X = 0.91
	pointsOfFunctionPlot[91].Y = 0.803

	pointsOfFunctionPlot[92].X = 0.92
	pointsOfFunctionPlot[92].Y = 0.799

	pointsOfFunctionPlot[93].X = 0.93
	pointsOfFunctionPlot[93].Y = 0.795

	pointsOfFunctionPlot[94].X = 0.94
	pointsOfFunctionPlot[94].Y = 0.791

	pointsOfFunctionPlot[95].X = 0.95
	pointsOfFunctionPlot[95].Y = 0.786

	pointsOfFunctionPlot[96].X = 0.96
	pointsOfFunctionPlot[96].Y = 0.782

	pointsOfFunctionPlot[97].X = 0.97
	pointsOfFunctionPlot[97].Y = 0.778

	pointsOfFunctionPlot[98].X = 0.98
	pointsOfFunctionPlot[98].Y = 0.773

	pointsOfFunctionPlot[99].X = 0.99
	pointsOfFunctionPlot[99].Y = 0.769

	pointsOfFunctionPlot[100].X = 1.0
	pointsOfFunctionPlot[100].Y = 0.765

	pointsOfFunctionPlot[101].X = 1.01
	pointsOfFunctionPlot[101].Y = 0.76

	pointsOfFunctionPlot[102].X = 1.02
	pointsOfFunctionPlot[102].Y = 0.756

	pointsOfFunctionPlot[103].X = 1.03
	pointsOfFunctionPlot[103].Y = 0.751

	pointsOfFunctionPlot[104].X = 1.04
	pointsOfFunctionPlot[104].Y = 0.747

	pointsOfFunctionPlot[105].X = 1.05
	pointsOfFunctionPlot[105].Y = 0.742

	pointsOfFunctionPlot[106].X = 1.06
	pointsOfFunctionPlot[106].Y = 0.738

	pointsOfFunctionPlot[107].X = 1.07
	pointsOfFunctionPlot[107].Y = 0.733

	pointsOfFunctionPlot[108].X = 1.08
	pointsOfFunctionPlot[108].Y = 0.728

	pointsOfFunctionPlot[109].X = 1.09
	pointsOfFunctionPlot[109].Y = 0.724

	pointsOfFunctionPlot[110].X = 1.10
	pointsOfFunctionPlot[110].Y = 0.719

	pointsOfFunctionPlot[111].X = 1.11
	pointsOfFunctionPlot[111].Y = 0.714

	pointsOfFunctionPlot[112].X = 1.12
	pointsOfFunctionPlot[112].Y = 0.71

	pointsOfFunctionPlot[113].X = 1.13
	pointsOfFunctionPlot[113].Y = 0.705

	pointsOfFunctionPlot[114].X = 1.14
	pointsOfFunctionPlot[114].Y = 0.7

	pointsOfFunctionPlot[115].X = 1.15
	pointsOfFunctionPlot[115].Y = 0.695

	pointsOfFunctionPlot[116].X = 1.16
	pointsOfFunctionPlot[116].Y = 0.69

	pointsOfFunctionPlot[117].X = 1.17
	pointsOfFunctionPlot[117].Y = 0.685

	pointsOfFunctionPlot[118].X = 1.18
	pointsOfFunctionPlot[118].Y = 0.681

	pointsOfFunctionPlot[119].X = 1.19
	pointsOfFunctionPlot[119].Y = 0.676

	pointsOfFunctionPlot[120].X = 1.20
	pointsOfFunctionPlot[120].Y = 0.671

	pointsOfFunctionPlot[121].X = 1.21
	pointsOfFunctionPlot[121].Y = 0.666

	pointsOfFunctionPlot[122].X = 1.22
	pointsOfFunctionPlot[122].Y = 0.661

	pointsOfFunctionPlot[123].X = 1.23
	pointsOfFunctionPlot[123].Y = 0.656

	pointsOfFunctionPlot[124].X = 1.24
	pointsOfFunctionPlot[124].Y = 0.651

	pointsOfFunctionPlot[125].X = 1.25
	pointsOfFunctionPlot[125].Y = 0.645

	pointsOfFunctionPlot[126].X = 1.26
	pointsOfFunctionPlot[126].Y = 0.64

	pointsOfFunctionPlot[127].X = 1.27
	pointsOfFunctionPlot[127].Y = 0.635

	pointsOfFunctionPlot[128].X = 1.28
	pointsOfFunctionPlot[128].Y = 0.63

	pointsOfFunctionPlot[129].X = 1.29
	pointsOfFunctionPlot[129].Y = 0.625

	pointsOfFunctionPlot[130].X = 1.30
	pointsOfFunctionPlot[130].Y = 0.62

	pointsOfFunctionPlot[131].X = 1.31
	pointsOfFunctionPlot[131].Y = 0.614

	pointsOfFunctionPlot[132].X = 1.32
	pointsOfFunctionPlot[132].Y = 0.609

	pointsOfFunctionPlot[133].X = 1.33
	pointsOfFunctionPlot[133].Y = 0.604

	pointsOfFunctionPlot[134].X = 1.34
	pointsOfFunctionPlot[134].Y = 0.599

	pointsOfFunctionPlot[135].X = 1.35
	pointsOfFunctionPlot[135].Y = 0.593

	pointsOfFunctionPlot[136].X = 1.36
	pointsOfFunctionPlot[136].Y = 0.588

	pointsOfFunctionPlot[137].X = 1.37
	pointsOfFunctionPlot[137].Y = 0.583

	pointsOfFunctionPlot[138].X = 1.38
	pointsOfFunctionPlot[138].Y = 0.577

	pointsOfFunctionPlot[139].X = 1.39
	pointsOfFunctionPlot[139].Y = 0.572

	pointsOfFunctionPlot[140].X = 1.40
	pointsOfFunctionPlot[140].Y = 0.566

	pointsOfFunctionPlot[141].X = 1.41
	pointsOfFunctionPlot[141].Y = 0.561

	pointsOfFunctionPlot[142].X = 1.42
	pointsOfFunctionPlot[142].Y = 0.555

	pointsOfFunctionPlot[143].X = 1.43
	pointsOfFunctionPlot[143].Y = 0.55

	pointsOfFunctionPlot[144].X = 1.44
	pointsOfFunctionPlot[144].Y = 0.545

	pointsOfFunctionPlot[145].X = 1.45
	pointsOfFunctionPlot[145].Y = 0.539

	pointsOfFunctionPlot[146].X = 1.46
	pointsOfFunctionPlot[146].Y = 0.534

	pointsOfFunctionPlot[147].X = 1.47
	pointsOfFunctionPlot[147].Y = 0.528

	pointsOfFunctionPlot[148].X = 1.48
	pointsOfFunctionPlot[148].Y = 0.522

	pointsOfFunctionPlot[149].X = 1.49
	pointsOfFunctionPlot[149].Y = 0.517

	pointsOfFunctionPlot[150].X = 1.50
	pointsOfFunctionPlot[150].Y = 0.511

	pointsOfFunctionPlot[151].X = 1.51
	pointsOfFunctionPlot[151].Y = 0.506

	pointsOfFunctionPlot[152].X = 1.52
	pointsOfFunctionPlot[152].Y = 0.5

	pointsOfFunctionPlot[153].X = 1.53
	pointsOfFunctionPlot[153].Y = 0.495

	pointsOfFunctionPlot[154].X = 1.54
	pointsOfFunctionPlot[154].Y = 0.489

	pointsOfFunctionPlot[155].X = 1.55
	pointsOfFunctionPlot[155].Y = 0.483

	pointsOfFunctionPlot[156].X = 1.56
	pointsOfFunctionPlot[156].Y = 0.478

	pointsOfFunctionPlot[157].X = 1.57
	pointsOfFunctionPlot[157].Y = 0.472

	pointsOfFunctionPlot[158].X = 1.58
	pointsOfFunctionPlot[158].Y = 0.466

	pointsOfFunctionPlot[159].X = 1.59
	pointsOfFunctionPlot[159].Y = 0.461

	pointsOfFunctionPlot[160].X = 1.60
	pointsOfFunctionPlot[160].Y = 0.455

	pointsOfFunctionPlot[161].X = 1.61
	pointsOfFunctionPlot[161].Y = 0.449

	pointsOfFunctionPlot[162].X = 1.62
	pointsOfFunctionPlot[162].Y = 0.443

	pointsOfFunctionPlot[163].X = 1.63
	pointsOfFunctionPlot[163].Y = 0.438

	pointsOfFunctionPlot[164].X = 1.64
	pointsOfFunctionPlot[164].Y = 0.432

	pointsOfFunctionPlot[165].X = 1.65
	pointsOfFunctionPlot[165].Y = 0.426

	pointsOfFunctionPlot[166].X = 1.66
	pointsOfFunctionPlot[166].Y = 0.421

	pointsOfFunctionPlot[167].X = 1.67
	pointsOfFunctionPlot[167].Y = 0.415

	pointsOfFunctionPlot[168].X = 1.68
	pointsOfFunctionPlot[168].Y = 0.409

	pointsOfFunctionPlot[169].X = 1.69
	pointsOfFunctionPlot[169].Y = 0.403

	pointsOfFunctionPlot[170].X = 1.70
	pointsOfFunctionPlot[170].Y = 0.397

	pointsOfFunctionPlot[171].X = 1.71
	pointsOfFunctionPlot[171].Y = 0.392

	pointsOfFunctionPlot[172].X = 1.72
	pointsOfFunctionPlot[172].Y = 0.386

	pointsOfFunctionPlot[173].X = 1.73
	pointsOfFunctionPlot[173].Y = 0.38

	pointsOfFunctionPlot[174].X = 1.74
	pointsOfFunctionPlot[174].Y = 0.374

	pointsOfFunctionPlot[175].X = 1.75
	pointsOfFunctionPlot[175].Y = 0.369

	pointsOfFunctionPlot[176].X = 1.76
	pointsOfFunctionPlot[176].Y = 0.363

	pointsOfFunctionPlot[177].X = 1.77
	pointsOfFunctionPlot[177].Y = 0.357

	pointsOfFunctionPlot[178].X = 1.78
	pointsOfFunctionPlot[178].Y = 0.351

	pointsOfFunctionPlot[179].X = 1.79
	pointsOfFunctionPlot[179].Y = 0.345

	pointsOfFunctionPlot[180].X = 1.80
	pointsOfFunctionPlot[180].Y = 0.339

	pointsOfFunctionPlot[181].X = 1.81
	pointsOfFunctionPlot[181].Y = 0.334

	pointsOfFunctionPlot[182].X = 1.82
	pointsOfFunctionPlot[182].Y = 0.328

	pointsOfFunctionPlot[183].X = 1.83
	pointsOfFunctionPlot[183].Y = 0.322

	pointsOfFunctionPlot[184].X = 1.84
	pointsOfFunctionPlot[184].Y = 0.316

	pointsOfFunctionPlot[185].X = 1.85
	pointsOfFunctionPlot[185].Y = 0.31

	pointsOfFunctionPlot[186].X = 1.86
	pointsOfFunctionPlot[186].Y = 0.305

	pointsOfFunctionPlot[187].X = 1.87
	pointsOfFunctionPlot[187].Y = 0.299

	pointsOfFunctionPlot[188].X = 1.88
	pointsOfFunctionPlot[188].Y = 0.293

	pointsOfFunctionPlot[189].X = 1.89
	pointsOfFunctionPlot[189].Y = 0.287

	pointsOfFunctionPlot[190].X = 1.90
	pointsOfFunctionPlot[190].Y = 0.281

	pointsOfFunctionPlot[191].X = 1.91
	pointsOfFunctionPlot[191].Y = 0.276

	pointsOfFunctionPlot[192].X = 1.92
	pointsOfFunctionPlot[192].Y = 0.27

	pointsOfFunctionPlot[193].X = 1.93
	pointsOfFunctionPlot[193].Y = 0.264

	pointsOfFunctionPlot[194].X = 1.94
	pointsOfFunctionPlot[194].Y = 0.258

	pointsOfFunctionPlot[195].X = 1.95
	pointsOfFunctionPlot[195].Y = 0.252

	pointsOfFunctionPlot[196].X = 1.96
	pointsOfFunctionPlot[196].Y = 0.247

	pointsOfFunctionPlot[197].X = 1.97
	pointsOfFunctionPlot[197].Y = 0.241

	pointsOfFunctionPlot[198].X = 1.98
	pointsOfFunctionPlot[198].Y = 0.235

	pointsOfFunctionPlot[199].X = 1.99
	pointsOfFunctionPlot[199].Y = 0.229

	pointsOfFunctionPlot[200].X = 2.0
	pointsOfFunctionPlot[200].Y = 0.223

	pointsOfFunctionPlot[201].X = 2.01
	pointsOfFunctionPlot[201].Y = 0.218

	pointsOfFunctionPlot[202].X = 2.02
	pointsOfFunctionPlot[202].Y = 0.212

	pointsOfFunctionPlot[203].X = 2.03
	pointsOfFunctionPlot[203].Y = 0.206

	pointsOfFunctionPlot[204].X = 2.04
	pointsOfFunctionPlot[204].Y = 0.2

	pointsOfFunctionPlot[205].X = 2.05
	pointsOfFunctionPlot[205].Y = 0.195

	pointsOfFunctionPlot[206].X = 2.06
	pointsOfFunctionPlot[206].Y = 0.189

	pointsOfFunctionPlot[207].X = 2.07
	pointsOfFunctionPlot[207].Y = 0.183

	pointsOfFunctionPlot[208].X = 2.08
	pointsOfFunctionPlot[208].Y = 0.177

	pointsOfFunctionPlot[209].X = 2.09
	pointsOfFunctionPlot[209].Y = 0.172

	pointsOfFunctionPlot[210].X = 2.10
	pointsOfFunctionPlot[210].Y = 0.166

	pointsOfFunctionPlot[211].X = 2.11
	pointsOfFunctionPlot[211].Y = 0.16

	pointsOfFunctionPlot[212].X = 2.12
	pointsOfFunctionPlot[212].Y = 0.155

	pointsOfFunctionPlot[213].X = 2.13
	pointsOfFunctionPlot[213].Y = 0.149

	pointsOfFunctionPlot[214].X = 2.14
	pointsOfFunctionPlot[214].Y = 0.143

	pointsOfFunctionPlot[215].X = 2.15
	pointsOfFunctionPlot[215].Y = 0.138

	pointsOfFunctionPlot[216].X = 2.16
	pointsOfFunctionPlot[216].Y = 0.132

	pointsOfFunctionPlot[217].X = 2.17
	pointsOfFunctionPlot[217].Y = 0.127

	pointsOfFunctionPlot[218].X = 2.18
	pointsOfFunctionPlot[218].Y = 0.121

	pointsOfFunctionPlot[219].X = 2.19
	pointsOfFunctionPlot[219].Y = 0.115

	pointsOfFunctionPlot[220].X = 2.20
	pointsOfFunctionPlot[220].Y = 0.11

	pointsOfFunctionPlot[221].X = 2.21
	pointsOfFunctionPlot[221].Y = 0.104

	pointsOfFunctionPlot[222].X = 2.22
	pointsOfFunctionPlot[222].Y = 0.099

	pointsOfFunctionPlot[223].X = 2.23
	pointsOfFunctionPlot[223].Y = 0.093

	pointsOfFunctionPlot[224].X = 2.24
	pointsOfFunctionPlot[224].Y = 0.088

	pointsOfFunctionPlot[225].X = 2.25
	pointsOfFunctionPlot[225].Y = 0.082

	pointsOfFunctionPlot[226].X = 2.26
	pointsOfFunctionPlot[226].Y = 0.077

	pointsOfFunctionPlot[227].X = 2.27
	pointsOfFunctionPlot[227].Y = 0.071

	pointsOfFunctionPlot[228].X = 2.28
	pointsOfFunctionPlot[228].Y = 0.066

	pointsOfFunctionPlot[229].X = 2.29
	pointsOfFunctionPlot[229].Y = 0.06

	pointsOfFunctionPlot[230].X = 2.30
	pointsOfFunctionPlot[230].Y = 0.055

	pointsOfFunctionPlot[231].X = 2.31
	pointsOfFunctionPlot[231].Y = 0.05

	pointsOfFunctionPlot[232].X = 2.32
	pointsOfFunctionPlot[232].Y = 0.044

	pointsOfFunctionPlot[233].X = 2.33
	pointsOfFunctionPlot[233].Y = 0.039

	pointsOfFunctionPlot[234].X = 2.34
	pointsOfFunctionPlot[234].Y = 0.034

	pointsOfFunctionPlot[235].X = 2.35
	pointsOfFunctionPlot[235].Y = 0.028

	pointsOfFunctionPlot[236].X = 2.36
	pointsOfFunctionPlot[236].Y = 0.023

	pointsOfFunctionPlot[237].X = 2.37
	pointsOfFunctionPlot[237].Y = 0.018

	pointsOfFunctionPlot[238].X = 2.38
	pointsOfFunctionPlot[238].Y = 0.012

	pointsOfFunctionPlot[239].X = 2.39
	pointsOfFunctionPlot[239].Y = 0.007

	pointsOfFunctionPlot[240].X = 2.40
	pointsOfFunctionPlot[240].Y = 0.002

	pointsOfFunctionPlot[241].X = 2.41
	pointsOfFunctionPlot[241].Y = -0.002

	pointsOfFunctionPlot[242].X = 2.42
	pointsOfFunctionPlot[242].Y = -0.007

	pointsOfFunctionPlot[243].X = 2.43
	pointsOfFunctionPlot[243].Y = -0.012

	pointsOfFunctionPlot[244].X = 2.44
	pointsOfFunctionPlot[244].Y = -0.018

	pointsOfFunctionPlot[245].X = 2.45
	pointsOfFunctionPlot[245].Y = -0.023

	pointsOfFunctionPlot[246].X = 2.46
	pointsOfFunctionPlot[246].Y = -0.028

	pointsOfFunctionPlot[247].X = 2.47
	pointsOfFunctionPlot[247].Y = -0.033

	pointsOfFunctionPlot[248].X = 2.48
	pointsOfFunctionPlot[248].Y = -0.038

	pointsOfFunctionPlot[249].X = 2.49
	pointsOfFunctionPlot[249].Y = -0.043

	pointsOfFunctionPlot[250].X = 2.50
	pointsOfFunctionPlot[250].Y = -0.048

	pointsOfFunctionPlot[251].X = 2.51
	pointsOfFunctionPlot[251].Y = -0.053

	pointsOfFunctionPlot[252].X = 2.52
	pointsOfFunctionPlot[252].Y = -0.058

	pointsOfFunctionPlot[253].X = 2.53
	pointsOfFunctionPlot[253].Y = -0.063

	pointsOfFunctionPlot[254].X = 2.54
	pointsOfFunctionPlot[254].Y = -0.068

	pointsOfFunctionPlot[255].X = 2.55
	pointsOfFunctionPlot[255].Y = -0.072

	pointsOfFunctionPlot[256].X = 2.56
	pointsOfFunctionPlot[256].Y = -0.077

	pointsOfFunctionPlot[257].X = 2.57
	pointsOfFunctionPlot[257].Y = -0.082

	pointsOfFunctionPlot[258].X = 2.58
	pointsOfFunctionPlot[258].Y = -0.087

	pointsOfFunctionPlot[259].X = 2.59
	pointsOfFunctionPlot[259].Y = -0.092

	pointsOfFunctionPlot[260].X = 2.60
	pointsOfFunctionPlot[260].Y = -0.096

	pointsOfFunctionPlot[261].X = 2.61
	pointsOfFunctionPlot[261].Y = -0.101

	pointsOfFunctionPlot[262].X = 2.62
	pointsOfFunctionPlot[262].Y = -0.106

	pointsOfFunctionPlot[263].X = 2.63
	pointsOfFunctionPlot[263].Y = -0.11

	pointsOfFunctionPlot[264].X = 2.64
	pointsOfFunctionPlot[264].Y = -0.115

	pointsOfFunctionPlot[265].X = 2.65
	pointsOfFunctionPlot[265].Y = -0.119

	pointsOfFunctionPlot[266].X = 2.66
	pointsOfFunctionPlot[266].Y = -0.124

	pointsOfFunctionPlot[267].X = 2.67
	pointsOfFunctionPlot[267].Y = -0.129

	pointsOfFunctionPlot[268].X = 2.68
	pointsOfFunctionPlot[268].Y = -0.133

	pointsOfFunctionPlot[269].X = 2.69
	pointsOfFunctionPlot[269].Y = -0.138

	pointsOfFunctionPlot[270].X = 2.70
	pointsOfFunctionPlot[270].Y = -0.142

	pointsOfFunctionPlot[271].X = 2.71
	pointsOfFunctionPlot[271].Y = -0.146

	pointsOfFunctionPlot[272].X = 2.72
	pointsOfFunctionPlot[272].Y = -0.151

	pointsOfFunctionPlot[273].X = 2.73
	pointsOfFunctionPlot[273].Y = -0.155

	pointsOfFunctionPlot[274].X = 2.74
	pointsOfFunctionPlot[274].Y = -0.159

	pointsOfFunctionPlot[275].X = 2.75
	pointsOfFunctionPlot[275].Y = -0.164

	pointsOfFunctionPlot[276].X = 2.76
	pointsOfFunctionPlot[276].Y = -0.168

	pointsOfFunctionPlot[277].X = 2.77
	pointsOfFunctionPlot[277].Y = -0.172

	pointsOfFunctionPlot[278].X = 2.78
	pointsOfFunctionPlot[278].Y = -0.176

	pointsOfFunctionPlot[279].X = 2.79
	pointsOfFunctionPlot[279].Y = -0.18

	pointsOfFunctionPlot[280].X = 2.80
	pointsOfFunctionPlot[280].Y = -0.185

	pointsOfFunctionPlot[281].X = 2.81
	pointsOfFunctionPlot[281].Y = -0.189

	pointsOfFunctionPlot[282].X = 2.82
	pointsOfFunctionPlot[282].Y = -0.193

	pointsOfFunctionPlot[283].X = 2.83
	pointsOfFunctionPlot[283].Y = -0.197

	pointsOfFunctionPlot[284].X = 2.84
	pointsOfFunctionPlot[284].Y = -0.201

	pointsOfFunctionPlot[285].X = 2.85
	pointsOfFunctionPlot[285].Y = -0.205

	pointsOfFunctionPlot[286].X = 2.86
	pointsOfFunctionPlot[286].Y = -0.209

	pointsOfFunctionPlot[287].X = 2.87
	pointsOfFunctionPlot[287].Y = -0.212

	pointsOfFunctionPlot[288].X = 2.88
	pointsOfFunctionPlot[288].Y = -0.216

	pointsOfFunctionPlot[289].X = 2.89
	pointsOfFunctionPlot[289].Y = -0.22

	pointsOfFunctionPlot[290].X = 2.90
	pointsOfFunctionPlot[290].Y = -0.224

	pointsOfFunctionPlot[291].X = 2.91
	pointsOfFunctionPlot[291].Y = -0.228

	pointsOfFunctionPlot[292].X = 2.92
	pointsOfFunctionPlot[292].Y = -0.231

	pointsOfFunctionPlot[293].X = 2.93
	pointsOfFunctionPlot[293].Y = -0.235

	pointsOfFunctionPlot[294].X = 2.94
	pointsOfFunctionPlot[294].Y = -0.239

	pointsOfFunctionPlot[295].X = 2.95
	pointsOfFunctionPlot[295].Y = -0.242

	pointsOfFunctionPlot[296].X = 2.96
	pointsOfFunctionPlot[296].Y = -0.246

	pointsOfFunctionPlot[297].X = 2.97
	pointsOfFunctionPlot[297].Y = -0.249

	pointsOfFunctionPlot[298].X = 2.98
	pointsOfFunctionPlot[298].Y = -0.253

	pointsOfFunctionPlot[299].X = 2.99
	pointsOfFunctionPlot[299].Y = -0.256

	pointsOfFunctionPlot[300].X = 3.0
	pointsOfFunctionPlot[300].Y = -0.26

	pointsOfFunctionPlot[301].X = 3.01
	pointsOfFunctionPlot[301].Y = -0.263

	pointsOfFunctionPlot[302].X = 3.02
	pointsOfFunctionPlot[302].Y = -0.266

	pointsOfFunctionPlot[303].X = 3.03
	pointsOfFunctionPlot[303].Y = -0.27

	pointsOfFunctionPlot[304].X = 3.04
	pointsOfFunctionPlot[304].Y = -0.273

	pointsOfFunctionPlot[305].X = 3.05
	pointsOfFunctionPlot[305].Y = -0.276

	pointsOfFunctionPlot[306].X = 3.06
	pointsOfFunctionPlot[306].Y = -0.279

	pointsOfFunctionPlot[307].X = 3.07
	pointsOfFunctionPlot[307].Y = -0.282

	pointsOfFunctionPlot[308].X = 3.08
	pointsOfFunctionPlot[308].Y = -0.285

	pointsOfFunctionPlot[309].X = 3.09
	pointsOfFunctionPlot[309].Y = -0.289

	pointsOfFunctionPlot[310].X = 3.10
	pointsOfFunctionPlot[310].Y = -0.292

	pointsOfFunctionPlot[311].X = 3.11
	pointsOfFunctionPlot[311].Y = -0.295

	pointsOfFunctionPlot[312].X = 3.12
	pointsOfFunctionPlot[312].Y = -0.298

	pointsOfFunctionPlot[313].X = 3.13
	pointsOfFunctionPlot[313].Y = -0.3

	pointsOfFunctionPlot[314].X = 3.14
	pointsOfFunctionPlot[314].Y = -0.303

	pointsOfFunctionPlot[315].X = 3.15
	pointsOfFunctionPlot[315].Y = -0.306

	pointsOfFunctionPlot[316].X = 3.16
	pointsOfFunctionPlot[316].Y = -0.309

	pointsOfFunctionPlot[317].X = 3.17
	pointsOfFunctionPlot[317].Y = -0.312

	pointsOfFunctionPlot[318].X = 3.18
	pointsOfFunctionPlot[318].Y = -0.314

	pointsOfFunctionPlot[319].X = 3.19
	pointsOfFunctionPlot[319].Y = -0.317

	pointsOfFunctionPlot[320].X = 3.20
	pointsOfFunctionPlot[320].Y = -0.32

	pointsOfFunctionPlot[321].X = 3.21
	pointsOfFunctionPlot[321].Y = -0.322

	pointsOfFunctionPlot[322].X = 3.22
	pointsOfFunctionPlot[322].Y = -0.325

	pointsOfFunctionPlot[323].X = 3.23
	pointsOfFunctionPlot[323].Y = -0.327

	pointsOfFunctionPlot[324].X = 3.24
	pointsOfFunctionPlot[324].Y = -0.33

	pointsOfFunctionPlot[325].X = 3.25
	pointsOfFunctionPlot[325].Y = -0.332

	pointsOfFunctionPlot[326].X = 3.26
	pointsOfFunctionPlot[326].Y = -0.335

	pointsOfFunctionPlot[327].X = 3.27
	pointsOfFunctionPlot[327].Y = -0.337

	pointsOfFunctionPlot[328].X = 3.28
	pointsOfFunctionPlot[328].Y = -0.339

	pointsOfFunctionPlot[329].X = 3.29
	pointsOfFunctionPlot[329].Y = -0.342

	pointsOfFunctionPlot[330].X = 3.30
	pointsOfFunctionPlot[330].Y = -0.344

	pointsOfFunctionPlot[331].X = 3.31
	pointsOfFunctionPlot[331].Y = -0.346

	pointsOfFunctionPlot[332].X = 3.32
	pointsOfFunctionPlot[332].Y = -0.348

	pointsOfFunctionPlot[333].X = 3.33
	pointsOfFunctionPlot[333].Y = -0.35

	pointsOfFunctionPlot[334].X = 3.34
	pointsOfFunctionPlot[334].Y = -0.352

	pointsOfFunctionPlot[335].X = 3.35
	pointsOfFunctionPlot[335].Y = -0.354

	pointsOfFunctionPlot[336].X = 3.36
	pointsOfFunctionPlot[336].Y = -0.356

	pointsOfFunctionPlot[337].X = 3.37
	pointsOfFunctionPlot[337].Y = -0.358

	pointsOfFunctionPlot[338].X = 3.38
	pointsOfFunctionPlot[338].Y = -0.36

	pointsOfFunctionPlot[339].X = 3.39
	pointsOfFunctionPlot[339].Y = -0.362

	pointsOfFunctionPlot[340].X = 3.40
	pointsOfFunctionPlot[340].Y = -0.364

	pointsOfFunctionPlot[341].X = 3.41
	pointsOfFunctionPlot[341].Y = -0.366

	pointsOfFunctionPlot[342].X = 3.42
	pointsOfFunctionPlot[342].Y = -0.367

	pointsOfFunctionPlot[343].X = 3.43
	pointsOfFunctionPlot[343].Y = -0.369

	pointsOfFunctionPlot[344].X = 3.44
	pointsOfFunctionPlot[344].Y = -0.371

	pointsOfFunctionPlot[345].X = 3.45
	pointsOfFunctionPlot[345].Y = -0.372

	pointsOfFunctionPlot[346].X = 3.46
	pointsOfFunctionPlot[346].Y = -0.374

	pointsOfFunctionPlot[347].X = 3.47
	pointsOfFunctionPlot[347].Y = -0.375

	pointsOfFunctionPlot[348].X = 3.48
	pointsOfFunctionPlot[348].Y = -0.377

	pointsOfFunctionPlot[349].X = 3.49
	pointsOfFunctionPlot[349].Y = -0.378

	pointsOfFunctionPlot[350].X = 3.50
	pointsOfFunctionPlot[350].Y = -0.38

	pointsOfFunctionPlot[351].X = 3.51
	pointsOfFunctionPlot[351].Y = -0.381

	pointsOfFunctionPlot[352].X = 3.52
	pointsOfFunctionPlot[352].Y = -0.382

	pointsOfFunctionPlot[353].X = 3.53
	pointsOfFunctionPlot[353].Y = -0.384

	pointsOfFunctionPlot[354].X = 3.54
	pointsOfFunctionPlot[354].Y = -0.385

	pointsOfFunctionPlot[355].X = 3.55
	pointsOfFunctionPlot[355].Y = -0.386

	pointsOfFunctionPlot[356].X = 3.56
	pointsOfFunctionPlot[356].Y = -0.387

	pointsOfFunctionPlot[357].X = 3.57
	pointsOfFunctionPlot[357].Y = -0.388

	pointsOfFunctionPlot[358].X = 3.58
	pointsOfFunctionPlot[358].Y = -0.389

	pointsOfFunctionPlot[359].X = 3.59
	pointsOfFunctionPlot[359].Y = -0.39

	pointsOfFunctionPlot[360].X = 3.60
	pointsOfFunctionPlot[360].Y = -0.391

	pointsOfFunctionPlot[361].X = 3.61
	pointsOfFunctionPlot[361].Y = -0.392

	pointsOfFunctionPlot[362].X = 3.62
	pointsOfFunctionPlot[362].Y = -0.393

	pointsOfFunctionPlot[363].X = 3.63
	pointsOfFunctionPlot[363].Y = -0.394

	pointsOfFunctionPlot[364].X = 3.64
	pointsOfFunctionPlot[364].Y = -0.395

	pointsOfFunctionPlot[365].X = 3.65
	pointsOfFunctionPlot[365].Y = -0.396

	pointsOfFunctionPlot[366].X = 3.66
	pointsOfFunctionPlot[366].Y = -0.396

	pointsOfFunctionPlot[367].X = 3.67
	pointsOfFunctionPlot[367].Y = -0.397

	pointsOfFunctionPlot[368].X = 3.68
	pointsOfFunctionPlot[368].Y = -0.398

	pointsOfFunctionPlot[369].X = 3.69
	pointsOfFunctionPlot[369].Y = -0.398

	pointsOfFunctionPlot[370].X = 3.70
	pointsOfFunctionPlot[370].Y = -0.399

	pointsOfFunctionPlot[371].X = 3.71
	pointsOfFunctionPlot[371].Y = -0.399

	pointsOfFunctionPlot[372].X = 3.72
	pointsOfFunctionPlot[372].Y = -0.4

	pointsOfFunctionPlot[373].X = 3.73
	pointsOfFunctionPlot[373].Y = -0.4

	pointsOfFunctionPlot[374].X = 3.74
	pointsOfFunctionPlot[374].Y = -0.401

	pointsOfFunctionPlot[375].X = 3.75
	pointsOfFunctionPlot[375].Y = -0.401

	pointsOfFunctionPlot[376].X = 3.76
	pointsOfFunctionPlot[376].Y = -0.401

	pointsOfFunctionPlot[377].X = 3.77
	pointsOfFunctionPlot[377].Y = -0.401

	pointsOfFunctionPlot[378].X = 3.78
	pointsOfFunctionPlot[378].Y = -0.402

	pointsOfFunctionPlot[379].X = 3.79
	pointsOfFunctionPlot[379].Y = -0.402

	pointsOfFunctionPlot[380].X = 3.80
	pointsOfFunctionPlot[380].Y = -0.402

	pointsOfFunctionPlot[381].X = 3.81
	pointsOfFunctionPlot[381].Y = -0.402

	pointsOfFunctionPlot[382].X = 3.82
	pointsOfFunctionPlot[382].Y = -0.402

	pointsOfFunctionPlot[383].X = 3.83
	pointsOfFunctionPlot[383].Y = -0.402

	pointsOfFunctionPlot[384].X = 3.84
	pointsOfFunctionPlot[384].Y = -0.402

	pointsOfFunctionPlot[385].X = 3.85
	pointsOfFunctionPlot[385].Y = -0.402

	pointsOfFunctionPlot[386].X = 3.86
	pointsOfFunctionPlot[386].Y = -0.402

	pointsOfFunctionPlot[387].X = 3.87
	pointsOfFunctionPlot[387].Y = -0.402

	pointsOfFunctionPlot[388].X = 3.88
	pointsOfFunctionPlot[388].Y = -0.402

	pointsOfFunctionPlot[389].X = 3.89
	pointsOfFunctionPlot[389].Y = -0.402

	pointsOfFunctionPlot[390].X = 3.90
	pointsOfFunctionPlot[390].Y = -0.401

	pointsOfFunctionPlot[391].X = 3.91
	pointsOfFunctionPlot[391].Y = -0.401

	pointsOfFunctionPlot[392].X = 3.92
	pointsOfFunctionPlot[392].Y = -0.401

	pointsOfFunctionPlot[393].X = 3.93
	pointsOfFunctionPlot[393].Y = -0.4

	pointsOfFunctionPlot[394].X = 3.94
	pointsOfFunctionPlot[394].Y = -0.4

	pointsOfFunctionPlot[395].X = 3.95
	pointsOfFunctionPlot[395].Y = -0.399

	pointsOfFunctionPlot[396].X = 3.96
	pointsOfFunctionPlot[396].Y = -0.399

	pointsOfFunctionPlot[397].X = 3.97
	pointsOfFunctionPlot[397].Y = -0.398

	pointsOfFunctionPlot[398].X = 3.98
	pointsOfFunctionPlot[398].Y = -0.398

	pointsOfFunctionPlot[399].X = 3.99
	pointsOfFunctionPlot[399].Y = -0.397

	pointsOfFunctionPlot[400].X = 4.0
	pointsOfFunctionPlot[400].Y = -0.397

	pointsOfFunctionPlot[401].X = 4.01
	pointsOfFunctionPlot[401].Y = -0.396

	pointsOfFunctionPlot[402].X = 4.02
	pointsOfFunctionPlot[402].Y = -0.395

	pointsOfFunctionPlot[403].X = 4.03
	pointsOfFunctionPlot[403].Y = -0.394

	pointsOfFunctionPlot[404].X = 4.04
	pointsOfFunctionPlot[404].Y = -0.394

	pointsOfFunctionPlot[405].X = 4.05
	pointsOfFunctionPlot[405].Y = -0.393

	pointsOfFunctionPlot[406].X = 4.06
	pointsOfFunctionPlot[406].Y = -0.392

	pointsOfFunctionPlot[407].X = 4.07
	pointsOfFunctionPlot[407].Y = -0.391

	pointsOfFunctionPlot[408].X = 4.08
	pointsOfFunctionPlot[408].Y = -0.39

	pointsOfFunctionPlot[409].X = 4.09
	pointsOfFunctionPlot[409].Y = -0.389

	pointsOfFunctionPlot[410].X = 4.10
	pointsOfFunctionPlot[410].Y = -0.388

	pointsOfFunctionPlot[411].X = 4.11
	pointsOfFunctionPlot[411].Y = -0.387

	pointsOfFunctionPlot[412].X = 4.12
	pointsOfFunctionPlot[412].Y = -0.386

	pointsOfFunctionPlot[413].X = 4.13
	pointsOfFunctionPlot[413].Y = -0.385

	pointsOfFunctionPlot[414].X = 4.14
	pointsOfFunctionPlot[414].Y = -0.384

	pointsOfFunctionPlot[415].X = 4.15
	pointsOfFunctionPlot[415].Y = -0.383

	pointsOfFunctionPlot[416].X = 4.16
	pointsOfFunctionPlot[416].Y = -0.381

	pointsOfFunctionPlot[417].X = 4.17
	pointsOfFunctionPlot[417].Y = -0.38

	pointsOfFunctionPlot[418].X = 4.18
	pointsOfFunctionPlot[418].Y = -0.379

	pointsOfFunctionPlot[419].X = 4.19
	pointsOfFunctionPlot[419].Y = -0.377

	pointsOfFunctionPlot[420].X = 4.20
	pointsOfFunctionPlot[420].Y = -0.376

	pointsOfFunctionPlot[421].X = 4.21
	pointsOfFunctionPlot[421].Y = -0.375

	pointsOfFunctionPlot[422].X = 4.22
	pointsOfFunctionPlot[422].Y = -0.373

	pointsOfFunctionPlot[423].X = 4.23
	pointsOfFunctionPlot[423].Y = -0.372

	pointsOfFunctionPlot[424].X = 4.24
	pointsOfFunctionPlot[424].Y = -0.37

	pointsOfFunctionPlot[425].X = 4.25
	pointsOfFunctionPlot[425].Y = -0.369

	pointsOfFunctionPlot[426].X = 4.26
	pointsOfFunctionPlot[426].Y = -0.367

	pointsOfFunctionPlot[427].X = 4.27
	pointsOfFunctionPlot[427].Y = -0.366

	pointsOfFunctionPlot[428].X = 4.28
	pointsOfFunctionPlot[428].Y = -0.364

	pointsOfFunctionPlot[429].X = 4.29
	pointsOfFunctionPlot[429].Y = -0.362

	pointsOfFunctionPlot[430].X = 4.30
	pointsOfFunctionPlot[430].Y = -0.361

	pointsOfFunctionPlot[431].X = 4.31
	pointsOfFunctionPlot[431].Y = -0.359

	pointsOfFunctionPlot[432].X = 4.32
	pointsOfFunctionPlot[432].Y = -0.357

	pointsOfFunctionPlot[433].X = 4.33
	pointsOfFunctionPlot[433].Y = -0.355

	pointsOfFunctionPlot[434].X = 4.34
	pointsOfFunctionPlot[434].Y = -0.353

	pointsOfFunctionPlot[435].X = 4.35
	pointsOfFunctionPlot[435].Y = -0.352

	pointsOfFunctionPlot[436].X = 4.36
	pointsOfFunctionPlot[436].Y = -0.35

	pointsOfFunctionPlot[437].X = 4.37
	pointsOfFunctionPlot[437].Y = -0.348

	pointsOfFunctionPlot[438].X = 4.38
	pointsOfFunctionPlot[438].Y = -0.346

	pointsOfFunctionPlot[439].X = 4.39
	pointsOfFunctionPlot[439].Y = -0.344

	pointsOfFunctionPlot[440].X = 4.40
	pointsOfFunctionPlot[440].Y = -0.342

	pointsOfFunctionPlot[441].X = 4.41
	pointsOfFunctionPlot[441].Y = -0.34

	pointsOfFunctionPlot[442].X = 4.42
	pointsOfFunctionPlot[442].Y = -0.338

	pointsOfFunctionPlot[443].X = 4.43
	pointsOfFunctionPlot[443].Y = -0.336

	pointsOfFunctionPlot[444].X = 4.44
	pointsOfFunctionPlot[444].Y = -0.333

	pointsOfFunctionPlot[445].X = 4.45
	pointsOfFunctionPlot[445].Y = -0.331

	pointsOfFunctionPlot[446].X = 4.46
	pointsOfFunctionPlot[446].Y = -0.329

	pointsOfFunctionPlot[447].X = 4.47
	pointsOfFunctionPlot[447].Y = -0.327

	pointsOfFunctionPlot[448].X = 4.48
	pointsOfFunctionPlot[448].Y = -0.325

	pointsOfFunctionPlot[449].X = 4.49
	pointsOfFunctionPlot[449].Y = -0.322

	pointsOfFunctionPlot[450].X = 4.50
	pointsOfFunctionPlot[450].Y = -0.32

	pointsOfFunctionPlot[451].X = 4.51
	pointsOfFunctionPlot[451].Y = -0.318

	pointsOfFunctionPlot[452].X = 4.52
	pointsOfFunctionPlot[452].Y = -0.315

	pointsOfFunctionPlot[453].X = 4.53
	pointsOfFunctionPlot[453].Y = -0.313

	pointsOfFunctionPlot[454].X = 4.54
	pointsOfFunctionPlot[454].Y = -0.311

	pointsOfFunctionPlot[455].X = 4.55
	pointsOfFunctionPlot[455].Y = -0.308

	pointsOfFunctionPlot[456].X = 4.56
	pointsOfFunctionPlot[456].Y = -0.306

	pointsOfFunctionPlot[457].X = 4.57
	pointsOfFunctionPlot[457].Y = -0.303

	pointsOfFunctionPlot[458].X = 4.58
	pointsOfFunctionPlot[458].Y = -0.301

	pointsOfFunctionPlot[459].X = 4.59
	pointsOfFunctionPlot[459].Y = -0.298

	pointsOfFunctionPlot[460].X = 4.60
	pointsOfFunctionPlot[460].Y = -0.296

	pointsOfFunctionPlot[461].X = 4.61
	pointsOfFunctionPlot[461].Y = -0.293

	pointsOfFunctionPlot[462].X = 4.62
	pointsOfFunctionPlot[462].Y = -0.29

	pointsOfFunctionPlot[463].X = 4.63
	pointsOfFunctionPlot[463].Y = -0.288

	pointsOfFunctionPlot[464].X = 4.64
	pointsOfFunctionPlot[464].Y = -0.285

	pointsOfFunctionPlot[465].X = 4.65
	pointsOfFunctionPlot[465].Y = -0.283

	pointsOfFunctionPlot[466].X = 4.66
	pointsOfFunctionPlot[466].Y = -0.28

	pointsOfFunctionPlot[467].X = 4.67
	pointsOfFunctionPlot[467].Y = -0.277

	pointsOfFunctionPlot[468].X = 4.68
	pointsOfFunctionPlot[468].Y = -0.274

	pointsOfFunctionPlot[469].X = 4.69
	pointsOfFunctionPlot[469].Y = -0.272

	pointsOfFunctionPlot[470].X = 4.70
	pointsOfFunctionPlot[470].Y = -0.269

	pointsOfFunctionPlot[471].X = 4.71
	pointsOfFunctionPlot[471].Y = -0.266

	pointsOfFunctionPlot[472].X = 4.72
	pointsOfFunctionPlot[472].Y = -0.263

	pointsOfFunctionPlot[473].X = 4.73
	pointsOfFunctionPlot[473].Y = -0.26

	pointsOfFunctionPlot[474].X = 4.74
	pointsOfFunctionPlot[474].Y = -0.258

	pointsOfFunctionPlot[475].X = 4.75
	pointsOfFunctionPlot[475].Y = -0.255

	pointsOfFunctionPlot[476].X = 4.76
	pointsOfFunctionPlot[476].Y = -0.252

	pointsOfFunctionPlot[477].X = 4.77
	pointsOfFunctionPlot[477].Y = -0.249

	pointsOfFunctionPlot[478].X = 4.78
	pointsOfFunctionPlot[478].Y = -0.246

	pointsOfFunctionPlot[479].X = 4.79
	pointsOfFunctionPlot[479].Y = -0.243

	pointsOfFunctionPlot[480].X = 4.80
	pointsOfFunctionPlot[480].Y = -0.24

	pointsOfFunctionPlot[481].X = 4.81
	pointsOfFunctionPlot[481].Y = -0.237

	pointsOfFunctionPlot[482].X = 4.82
	pointsOfFunctionPlot[482].Y = -0.234

	pointsOfFunctionPlot[483].X = 4.83
	pointsOfFunctionPlot[483].Y = -0.231

	pointsOfFunctionPlot[484].X = 4.84
	pointsOfFunctionPlot[484].Y = -0.228

	pointsOfFunctionPlot[485].X = 4.85
	pointsOfFunctionPlot[485].Y = -0.225

	pointsOfFunctionPlot[486].X = 4.86
	pointsOfFunctionPlot[486].Y = -0.222

	pointsOfFunctionPlot[487].X = 4.87
	pointsOfFunctionPlot[487].Y = -0.219

	pointsOfFunctionPlot[488].X = 4.88
	pointsOfFunctionPlot[488].Y = -0.216

	pointsOfFunctionPlot[489].X = 4.89
	pointsOfFunctionPlot[489].Y = -0.212

	pointsOfFunctionPlot[490].X = 4.90
	pointsOfFunctionPlot[490].Y = -0.209

	pointsOfFunctionPlot[491].X = 4.91
	pointsOfFunctionPlot[491].Y = -0.206

	pointsOfFunctionPlot[492].X = 4.92
	pointsOfFunctionPlot[492].Y = -0.203

	pointsOfFunctionPlot[493].X = 4.93
	pointsOfFunctionPlot[493].Y = -0.2

	pointsOfFunctionPlot[494].X = 4.94
	pointsOfFunctionPlot[494].Y = -0.197

	pointsOfFunctionPlot[495].X = 4.95
	pointsOfFunctionPlot[495].Y = -0.193

	pointsOfFunctionPlot[496].X = 4.96
	pointsOfFunctionPlot[496].Y = -0.19

	pointsOfFunctionPlot[497].X = 4.97
	pointsOfFunctionPlot[497].Y = -0.187

	pointsOfFunctionPlot[498].X = 4.98
	pointsOfFunctionPlot[498].Y = -0.184

	pointsOfFunctionPlot[499].X = 4.99
	pointsOfFunctionPlot[499].Y = -0.18

	pointsOfFunctionPlot[500].X = 5.0
	pointsOfFunctionPlot[500].Y = -0.177

	pointsOfFunctionPlot[501].X = 5.01
	pointsOfFunctionPlot[501].Y = -0.174

	pointsOfFunctionPlot[502].X = 5.02
	pointsOfFunctionPlot[502].Y = -0.171

	pointsOfFunctionPlot[503].X = 5.03
	pointsOfFunctionPlot[503].Y = -0.167

	pointsOfFunctionPlot[504].X = 5.04
	pointsOfFunctionPlot[504].Y = -0.164

	pointsOfFunctionPlot[505].X = 5.05
	pointsOfFunctionPlot[505].Y = -0.161

	pointsOfFunctionPlot[506].X = 5.06
	pointsOfFunctionPlot[506].Y = -0.157

	pointsOfFunctionPlot[507].X = 5.07
	pointsOfFunctionPlot[507].Y = -0.154

	pointsOfFunctionPlot[508].X = 5.08
	pointsOfFunctionPlot[508].Y = -0.151

	pointsOfFunctionPlot[509].X = 5.09
	pointsOfFunctionPlot[509].Y = -0.147

	pointsOfFunctionPlot[510].X = 5.10
	pointsOfFunctionPlot[510].Y = -0.144

	pointsOfFunctionPlot[511].X = 5.11
	pointsOfFunctionPlot[511].Y = -0.14

	pointsOfFunctionPlot[512].X = 5.12
	pointsOfFunctionPlot[512].Y = -0.137

	pointsOfFunctionPlot[513].X = 5.13
	pointsOfFunctionPlot[513].Y = -0.134

	pointsOfFunctionPlot[514].X = 5.14
	pointsOfFunctionPlot[514].Y = -0.13

	pointsOfFunctionPlot[515].X = 5.15
	pointsOfFunctionPlot[515].Y = -0.127

	pointsOfFunctionPlot[516].X = 5.16
	pointsOfFunctionPlot[516].Y = -0.123

	pointsOfFunctionPlot[517].X = 5.17
	pointsOfFunctionPlot[517].Y = -0.12

	pointsOfFunctionPlot[518].X = 5.18
	pointsOfFunctionPlot[518].Y = -0.117

	pointsOfFunctionPlot[519].X = 5.19
	pointsOfFunctionPlot[519].Y = -0.113

	pointsOfFunctionPlot[520].X = 5.20
	pointsOfFunctionPlot[520].Y = -0.11

	pointsOfFunctionPlot[521].X = 5.21
	pointsOfFunctionPlot[521].Y = -0.106

	pointsOfFunctionPlot[522].X = 5.22
	pointsOfFunctionPlot[522].Y = -0.103

	pointsOfFunctionPlot[523].X = 5.23
	pointsOfFunctionPlot[523].Y = -0.099

	pointsOfFunctionPlot[524].X = 5.24
	pointsOfFunctionPlot[524].Y = -0.096

	pointsOfFunctionPlot[525].X = 5.25
	pointsOfFunctionPlot[525].Y = -0.093

	pointsOfFunctionPlot[526].X = 5.26
	pointsOfFunctionPlot[526].Y = -0.089

	pointsOfFunctionPlot[527].X = 5.27
	pointsOfFunctionPlot[527].Y = -0.086

	pointsOfFunctionPlot[528].X = 5.28
	pointsOfFunctionPlot[528].Y = -0.082

	pointsOfFunctionPlot[529].X = 5.29
	pointsOfFunctionPlot[529].Y = -0.079

	pointsOfFunctionPlot[530].X = 5.30
	pointsOfFunctionPlot[530].Y = -0.075

	pointsOfFunctionPlot[531].X = 5.31
	pointsOfFunctionPlot[531].Y = -0.072

	pointsOfFunctionPlot[532].X = 5.32
	pointsOfFunctionPlot[532].Y = -0.068

	pointsOfFunctionPlot[533].X = 5.33
	pointsOfFunctionPlot[533].Y = -0.065

	pointsOfFunctionPlot[534].X = 5.34
	pointsOfFunctionPlot[534].Y = -0.061

	pointsOfFunctionPlot[535].X = 5.35
	pointsOfFunctionPlot[535].Y = -0.058

	pointsOfFunctionPlot[536].X = 5.36
	pointsOfFunctionPlot[536].Y = -0.055

	pointsOfFunctionPlot[537].X = 5.37
	pointsOfFunctionPlot[537].Y = -0.051

	pointsOfFunctionPlot[538].X = 5.38
	pointsOfFunctionPlot[538].Y = -0.048

	pointsOfFunctionPlot[539].X = 5.39
	pointsOfFunctionPlot[539].Y = -0.044

	pointsOfFunctionPlot[540].X = 5.40
	pointsOfFunctionPlot[540].Y = -0.041

	pointsOfFunctionPlot[541].X = 5.41
	pointsOfFunctionPlot[541].Y = -0.037

	pointsOfFunctionPlot[542].X = 5.42
	pointsOfFunctionPlot[542].Y = -0.034

	pointsOfFunctionPlot[543].X = 5.43
	pointsOfFunctionPlot[543].Y = -0.03

	pointsOfFunctionPlot[544].X = 5.44
	pointsOfFunctionPlot[544].Y = -0.027

	pointsOfFunctionPlot[545].X = 5.45
	pointsOfFunctionPlot[545].Y = -0.023

	pointsOfFunctionPlot[546].X = 5.46
	pointsOfFunctionPlot[546].Y = -0.02

	pointsOfFunctionPlot[547].X = 5.47
	pointsOfFunctionPlot[547].Y = -0.017

	pointsOfFunctionPlot[548].X = 5.48
	pointsOfFunctionPlot[548].Y = -0.013

	pointsOfFunctionPlot[549].X = 5.49
	pointsOfFunctionPlot[549].Y = -0.01

	pointsOfFunctionPlot[550].X = 5.50
	pointsOfFunctionPlot[550].Y = -0.006

	pointsOfFunctionPlot[551].X = 5.51
	pointsOfFunctionPlot[551].Y = -0.003

	pointsOfFunctionPlot[552].X = 5.52
	pointsOfFunctionPlot[552].Y = 0.0

	pointsOfFunctionPlot[553].X = 5.53
	pointsOfFunctionPlot[553].Y = 0.003

	pointsOfFunctionPlot[554].X = 5.54
	pointsOfFunctionPlot[554].Y = 0.006

	pointsOfFunctionPlot[555].X = 5.55
	pointsOfFunctionPlot[555].Y = 0.01

	pointsOfFunctionPlot[556].X = 5.56
	pointsOfFunctionPlot[556].Y = 0.013

	pointsOfFunctionPlot[557].X = 5.57
	pointsOfFunctionPlot[557].Y = 0.016

	pointsOfFunctionPlot[558].X = 5.58
	pointsOfFunctionPlot[558].Y = 0.02

	pointsOfFunctionPlot[559].X = 5.59
	pointsOfFunctionPlot[559].Y = 0.023

	pointsOfFunctionPlot[560].X = 5.60
	pointsOfFunctionPlot[560].Y = 0.026

	pointsOfFunctionPlot[561].X = 5.61
	pointsOfFunctionPlot[561].Y = 0.03

	pointsOfFunctionPlot[562].X = 5.62
	pointsOfFunctionPlot[562].Y = 0.033

	pointsOfFunctionPlot[563].X = 5.63
	pointsOfFunctionPlot[563].Y = 0.036

	pointsOfFunctionPlot[564].X = 5.64
	pointsOfFunctionPlot[564].Y = 0.04

	pointsOfFunctionPlot[565].X = 5.65
	pointsOfFunctionPlot[565].Y = 0.043

	pointsOfFunctionPlot[566].X = 5.66
	pointsOfFunctionPlot[566].Y = 0.046

	pointsOfFunctionPlot[567].X = 5.67
	pointsOfFunctionPlot[567].Y = 0.05

	pointsOfFunctionPlot[568].X = 5.68
	pointsOfFunctionPlot[568].Y = 0.053

	pointsOfFunctionPlot[569].X = 5.69
	pointsOfFunctionPlot[569].Y = 0.056

	pointsOfFunctionPlot[570].X = 5.70
	pointsOfFunctionPlot[570].Y = 0.059

	pointsOfFunctionPlot[571].X = 5.71
	pointsOfFunctionPlot[571].Y = 0.063

	pointsOfFunctionPlot[572].X = 5.72
	pointsOfFunctionPlot[572].Y = 0.066

	pointsOfFunctionPlot[573].X = 5.73
	pointsOfFunctionPlot[573].Y = 0.069

	pointsOfFunctionPlot[574].X = 5.74
	pointsOfFunctionPlot[574].Y = 0.072

	pointsOfFunctionPlot[575].X = 5.75
	pointsOfFunctionPlot[575].Y = 0.075

	pointsOfFunctionPlot[576].X = 5.76
	pointsOfFunctionPlot[576].Y = 0.079

	pointsOfFunctionPlot[577].X = 5.77
	pointsOfFunctionPlot[577].Y = 0.082

	pointsOfFunctionPlot[578].X = 5.78
	pointsOfFunctionPlot[578].Y = 0.085

	pointsOfFunctionPlot[579].X = 5.79
	pointsOfFunctionPlot[579].Y = 0.088

	pointsOfFunctionPlot[580].X = 5.80
	pointsOfFunctionPlot[580].Y = 0.091

	pointsOfFunctionPlot[581].X = 5.81
	pointsOfFunctionPlot[581].Y = 0.094

	pointsOfFunctionPlot[582].X = 5.82
	pointsOfFunctionPlot[582].Y = 0.097

	pointsOfFunctionPlot[583].X = 5.83
	pointsOfFunctionPlot[583].Y = 0.1

	pointsOfFunctionPlot[584].X = 5.84
	pointsOfFunctionPlot[584].Y = 0.104

	pointsOfFunctionPlot[585].X = 5.85
	pointsOfFunctionPlot[585].Y = 0.107

	pointsOfFunctionPlot[586].X = 5.86
	pointsOfFunctionPlot[586].Y = 0.11

	pointsOfFunctionPlot[587].X = 5.87
	pointsOfFunctionPlot[587].Y = 0.113

	pointsOfFunctionPlot[588].X = 5.88
	pointsOfFunctionPlot[588].Y = 0.116

	pointsOfFunctionPlot[589].X = 5.89
	pointsOfFunctionPlot[589].Y = 0.119

	pointsOfFunctionPlot[590].X = 5.90
	pointsOfFunctionPlot[590].Y = 0.122

	pointsOfFunctionPlot[591].X = 5.91
	pointsOfFunctionPlot[591].Y = 0.124

	pointsOfFunctionPlot[592].X = 5.92
	pointsOfFunctionPlot[592].Y = 0.127

	pointsOfFunctionPlot[593].X = 5.93
	pointsOfFunctionPlot[593].Y = 0.13

	pointsOfFunctionPlot[594].X = 5.94
	pointsOfFunctionPlot[594].Y = 0.133

	pointsOfFunctionPlot[595].X = 5.95
	pointsOfFunctionPlot[595].Y = 0.136

	pointsOfFunctionPlot[596].X = 5.96
	pointsOfFunctionPlot[596].Y = 0.139

	pointsOfFunctionPlot[597].X = 5.97
	pointsOfFunctionPlot[597].Y = 0.142

	pointsOfFunctionPlot[598].X = 5.98
	pointsOfFunctionPlot[598].Y = 0.145

	pointsOfFunctionPlot[599].X = 5.99
	pointsOfFunctionPlot[599].Y = 0.147

	pointsOfFunctionPlot[600].X = 6.0
	pointsOfFunctionPlot[600].Y = 0.15

	pointsOfFunctionPlot[601].X = 6.01
	pointsOfFunctionPlot[601].Y = 0.153

	pointsOfFunctionPlot[602].X = 6.02
	pointsOfFunctionPlot[602].Y = 0.156

	pointsOfFunctionPlot[603].X = 6.03
	pointsOfFunctionPlot[603].Y = 0.158

	pointsOfFunctionPlot[604].X = 6.04
	pointsOfFunctionPlot[604].Y = 0.161

	pointsOfFunctionPlot[605].X = 6.05
	pointsOfFunctionPlot[605].Y = 0.164

	pointsOfFunctionPlot[606].X = 6.06
	pointsOfFunctionPlot[606].Y = 0.166

	pointsOfFunctionPlot[607].X = 6.07
	pointsOfFunctionPlot[607].Y = 0.169

	pointsOfFunctionPlot[608].X = 6.08
	pointsOfFunctionPlot[608].Y = 0.172

	pointsOfFunctionPlot[609].X = 6.09
	pointsOfFunctionPlot[609].Y = 0.174

	pointsOfFunctionPlot[610].X = 6.10
	pointsOfFunctionPlot[610].Y = 0.177

	pointsOfFunctionPlot[611].X = 6.11
	pointsOfFunctionPlot[611].Y = 0.179

	pointsOfFunctionPlot[612].X = 6.12
	pointsOfFunctionPlot[612].Y = 0.182

	pointsOfFunctionPlot[613].X = 6.13
	pointsOfFunctionPlot[613].Y = 0.184

	pointsOfFunctionPlot[614].X = 6.14
	pointsOfFunctionPlot[614].Y = 0.187

	pointsOfFunctionPlot[615].X = 6.15
	pointsOfFunctionPlot[615].Y = 0.189

	pointsOfFunctionPlot[616].X = 6.16
	pointsOfFunctionPlot[616].Y = 0.192

	pointsOfFunctionPlot[617].X = 6.17
	pointsOfFunctionPlot[617].Y = 0.194

	pointsOfFunctionPlot[618].X = 6.18
	pointsOfFunctionPlot[618].Y = 0.197

	pointsOfFunctionPlot[619].X = 6.19
	pointsOfFunctionPlot[619].Y = 0.199

	pointsOfFunctionPlot[620].X = 6.20
	pointsOfFunctionPlot[620].Y = 0.201

	pointsOfFunctionPlot[621].X = 6.21
	pointsOfFunctionPlot[621].Y = 0.204

	pointsOfFunctionPlot[622].X = 6.22
	pointsOfFunctionPlot[622].Y = 0.206

	pointsOfFunctionPlot[623].X = 6.23
	pointsOfFunctionPlot[623].Y = 0.208

	pointsOfFunctionPlot[624].X = 6.24
	pointsOfFunctionPlot[624].Y = 0.21

	pointsOfFunctionPlot[625].X = 6.25
	pointsOfFunctionPlot[625].Y = 0.213

	pointsOfFunctionPlot[626].X = 6.26
	pointsOfFunctionPlot[626].Y = 0.215

	pointsOfFunctionPlot[627].X = 6.27
	pointsOfFunctionPlot[627].Y = 0.217

	pointsOfFunctionPlot[628].X = 6.28
	pointsOfFunctionPlot[628].Y = 0.219

	pointsOfFunctionPlot[629].X = 6.29
	pointsOfFunctionPlot[629].Y = 0.221

	pointsOfFunctionPlot[630].X = 6.30
	pointsOfFunctionPlot[630].Y = 0.223

	pointsOfFunctionPlot[631].X = 6.31
	pointsOfFunctionPlot[631].Y = 0.225

	pointsOfFunctionPlot[632].X = 6.32
	pointsOfFunctionPlot[632].Y = 0.227

	pointsOfFunctionPlot[633].X = 6.33
	pointsOfFunctionPlot[633].Y = 0.229

	pointsOfFunctionPlot[634].X = 6.34
	pointsOfFunctionPlot[634].Y = 0.231

	pointsOfFunctionPlot[635].X = 6.35
	pointsOfFunctionPlot[635].Y = 0.233

	pointsOfFunctionPlot[636].X = 6.36
	pointsOfFunctionPlot[636].Y = 0.235

	pointsOfFunctionPlot[637].X = 6.37
	pointsOfFunctionPlot[637].Y = 0.237

	pointsOfFunctionPlot[638].X = 6.38
	pointsOfFunctionPlot[638].Y = 0.239

	pointsOfFunctionPlot[639].X = 6.39
	pointsOfFunctionPlot[639].Y = 0.241

	pointsOfFunctionPlot[640].X = 6.40
	pointsOfFunctionPlot[640].Y = 0.243

	pointsOfFunctionPlot[641].X = 6.41
	pointsOfFunctionPlot[641].Y = 0.245

	pointsOfFunctionPlot[642].X = 6.42
	pointsOfFunctionPlot[642].Y = 0.246

	pointsOfFunctionPlot[643].X = 6.43
	pointsOfFunctionPlot[643].Y = 0.248

	pointsOfFunctionPlot[644].X = 6.44
	pointsOfFunctionPlot[644].Y = 0.25

	pointsOfFunctionPlot[645].X = 6.45
	pointsOfFunctionPlot[645].Y = 0.252

	pointsOfFunctionPlot[646].X = 6.46
	pointsOfFunctionPlot[646].Y = 0.253

	pointsOfFunctionPlot[647].X = 6.47
	pointsOfFunctionPlot[647].Y = 0.255

	pointsOfFunctionPlot[648].X = 6.48
	pointsOfFunctionPlot[648].Y = 0.256

	pointsOfFunctionPlot[649].X = 6.49
	pointsOfFunctionPlot[649].Y = 0.258

	pointsOfFunctionPlot[650].X = 6.50
	pointsOfFunctionPlot[650].Y = 0.26

	pointsOfFunctionPlot[651].X = 6.51
	pointsOfFunctionPlot[651].Y = 0.261

	pointsOfFunctionPlot[652].X = 6.52
	pointsOfFunctionPlot[652].Y = 0.263

	pointsOfFunctionPlot[653].X = 6.53
	pointsOfFunctionPlot[653].Y = 0.264

	pointsOfFunctionPlot[654].X = 6.54
	pointsOfFunctionPlot[654].Y = 0.266

	pointsOfFunctionPlot[655].X = 6.55
	pointsOfFunctionPlot[655].Y = 0.267

	pointsOfFunctionPlot[656].X = 6.56
	pointsOfFunctionPlot[656].Y = 0.268

	pointsOfFunctionPlot[657].X = 6.57
	pointsOfFunctionPlot[657].Y = 0.27

	pointsOfFunctionPlot[658].X = 6.58
	pointsOfFunctionPlot[658].Y = 0.271

	pointsOfFunctionPlot[659].X = 6.59
	pointsOfFunctionPlot[659].Y = 0.272

	pointsOfFunctionPlot[660].X = 6.60
	pointsOfFunctionPlot[660].Y = 0.274

	pointsOfFunctionPlot[661].X = 6.61
	pointsOfFunctionPlot[661].Y = 0.275

	pointsOfFunctionPlot[662].X = 6.62
	pointsOfFunctionPlot[662].Y = 0.276

	pointsOfFunctionPlot[663].X = 6.63
	pointsOfFunctionPlot[663].Y = 0.277

	pointsOfFunctionPlot[664].X = 6.64
	pointsOfFunctionPlot[664].Y = 0.278

	pointsOfFunctionPlot[665].X = 6.65
	pointsOfFunctionPlot[665].Y = 0.279

	pointsOfFunctionPlot[666].X = 6.66
	pointsOfFunctionPlot[666].Y = 0.281

	pointsOfFunctionPlot[667].X = 6.67
	pointsOfFunctionPlot[667].Y = 0.282

	pointsOfFunctionPlot[668].X = 6.68
	pointsOfFunctionPlot[668].Y = 0.283

	pointsOfFunctionPlot[669].X = 6.69
	pointsOfFunctionPlot[669].Y = 0.284

	pointsOfFunctionPlot[670].X = 6.70
	pointsOfFunctionPlot[670].Y = 0.285

	pointsOfFunctionPlot[671].X = 6.71
	pointsOfFunctionPlot[671].Y = 0.286

	pointsOfFunctionPlot[672].X = 6.72
	pointsOfFunctionPlot[672].Y = 0.286

	pointsOfFunctionPlot[673].X = 6.73
	pointsOfFunctionPlot[673].Y = 0.287

	pointsOfFunctionPlot[674].X = 6.74
	pointsOfFunctionPlot[674].Y = 0.288

	pointsOfFunctionPlot[675].X = 6.75
	pointsOfFunctionPlot[675].Y = 0.289

	pointsOfFunctionPlot[676].X = 6.76
	pointsOfFunctionPlot[676].Y = 0.29

	pointsOfFunctionPlot[677].X = 6.77
	pointsOfFunctionPlot[677].Y = 0.291

	pointsOfFunctionPlot[678].X = 6.78
	pointsOfFunctionPlot[678].Y = 0.291

	pointsOfFunctionPlot[679].X = 6.79
	pointsOfFunctionPlot[679].Y = 0.292

	pointsOfFunctionPlot[680].X = 6.80
	pointsOfFunctionPlot[680].Y = 0.293

	pointsOfFunctionPlot[681].X = 6.81
	pointsOfFunctionPlot[681].Y = 0.293

	pointsOfFunctionPlot[682].X = 6.82
	pointsOfFunctionPlot[682].Y = 0.294

	pointsOfFunctionPlot[683].X = 6.83
	pointsOfFunctionPlot[683].Y = 0.294

	pointsOfFunctionPlot[684].X = 6.84
	pointsOfFunctionPlot[684].Y = 0.295

	pointsOfFunctionPlot[685].X = 6.85
	pointsOfFunctionPlot[685].Y = 0.295

	pointsOfFunctionPlot[686].X = 6.86
	pointsOfFunctionPlot[686].Y = 0.296

	pointsOfFunctionPlot[687].X = 6.87
	pointsOfFunctionPlot[687].Y = 0.296

	pointsOfFunctionPlot[688].X = 6.88
	pointsOfFunctionPlot[688].Y = 0.297

	pointsOfFunctionPlot[689].X = 6.89
	pointsOfFunctionPlot[689].Y = 0.297

	pointsOfFunctionPlot[690].X = 6.90
	pointsOfFunctionPlot[690].Y = 0.298

	pointsOfFunctionPlot[691].X = 6.91
	pointsOfFunctionPlot[691].Y = 0.298

	pointsOfFunctionPlot[692].X = 6.92
	pointsOfFunctionPlot[692].Y = 0.298

	pointsOfFunctionPlot[693].X = 6.93
	pointsOfFunctionPlot[693].Y = 0.299

	pointsOfFunctionPlot[694].X = 6.94
	pointsOfFunctionPlot[694].Y = 0.299

	pointsOfFunctionPlot[695].X = 6.95
	pointsOfFunctionPlot[695].Y = 0.299

	pointsOfFunctionPlot[696].X = 6.96
	pointsOfFunctionPlot[696].Y = 0.299

	pointsOfFunctionPlot[697].X = 6.97
	pointsOfFunctionPlot[697].Y = 0.299

	pointsOfFunctionPlot[698].X = 6.98
	pointsOfFunctionPlot[698].Y = 0.299

	pointsOfFunctionPlot[699].X = 6.99
	pointsOfFunctionPlot[699].Y = 0.3

	pointsOfFunctionPlot[700].X = 7.0
	pointsOfFunctionPlot[700].Y = 0.3

	pointsOfFunctionPlot[701].X = 7.01
	pointsOfFunctionPlot[701].Y = 0.3

	pointsOfFunctionPlot[702].X = 7.02
	pointsOfFunctionPlot[702].Y = 0.3

	pointsOfFunctionPlot[703].X = 7.03
	pointsOfFunctionPlot[703].Y = 0.3

	pointsOfFunctionPlot[704].X = 7.04
	pointsOfFunctionPlot[704].Y = 0.3

	pointsOfFunctionPlot[705].X = 7.05
	pointsOfFunctionPlot[705].Y = 0.299

	pointsOfFunctionPlot[706].X = 7.06
	pointsOfFunctionPlot[706].Y = 0.299

	pointsOfFunctionPlot[707].X = 7.07
	pointsOfFunctionPlot[707].Y = 0.299

	pointsOfFunctionPlot[708].X = 7.08
	pointsOfFunctionPlot[708].Y = 0.299

	pointsOfFunctionPlot[709].X = 7.09
	pointsOfFunctionPlot[709].Y = 0.299

	pointsOfFunctionPlot[710].X = 7.10
	pointsOfFunctionPlot[710].Y = 0.299

	pointsOfFunctionPlot[711].X = 7.11
	pointsOfFunctionPlot[711].Y = 0.298

	pointsOfFunctionPlot[712].X = 7.12
	pointsOfFunctionPlot[712].Y = 0.298

	pointsOfFunctionPlot[713].X = 7.13
	pointsOfFunctionPlot[713].Y = 0.298

	pointsOfFunctionPlot[714].X = 7.14
	pointsOfFunctionPlot[714].Y = 0.297

	pointsOfFunctionPlot[715].X = 7.15
	pointsOfFunctionPlot[715].Y = 0.297

	pointsOfFunctionPlot[716].X = 7.16
	pointsOfFunctionPlot[716].Y = 0.297

	pointsOfFunctionPlot[717].X = 7.17
	pointsOfFunctionPlot[717].Y = 0.296

	pointsOfFunctionPlot[718].X = 7.18
	pointsOfFunctionPlot[718].Y = 0.296

	pointsOfFunctionPlot[719].X = 7.19
	pointsOfFunctionPlot[719].Y = 0.295

	pointsOfFunctionPlot[720].X = 7.20
	pointsOfFunctionPlot[720].Y = 0.295

	pointsOfFunctionPlot[721].X = 7.21
	pointsOfFunctionPlot[721].Y = 0.294

	pointsOfFunctionPlot[722].X = 7.22
	pointsOfFunctionPlot[722].Y = 0.293

	pointsOfFunctionPlot[723].X = 7.23
	pointsOfFunctionPlot[723].Y = 0.293

	pointsOfFunctionPlot[724].X = 7.24
	pointsOfFunctionPlot[724].Y = 0.292

	pointsOfFunctionPlot[725].X = 7.25
	pointsOfFunctionPlot[725].Y = 0.291

	pointsOfFunctionPlot[726].X = 7.26
	pointsOfFunctionPlot[726].Y = 0.291

	pointsOfFunctionPlot[727].X = 7.27
	pointsOfFunctionPlot[727].Y = 0.29

	pointsOfFunctionPlot[728].X = 7.28
	pointsOfFunctionPlot[728].Y = 0.289

	pointsOfFunctionPlot[729].X = 7.29
	pointsOfFunctionPlot[729].Y = 0.289

	pointsOfFunctionPlot[730].X = 7.30
	pointsOfFunctionPlot[730].Y = 0.288

	pointsOfFunctionPlot[731].X = 7.31
	pointsOfFunctionPlot[731].Y = 0.287

	pointsOfFunctionPlot[732].X = 7.32
	pointsOfFunctionPlot[732].Y = 0.286

	pointsOfFunctionPlot[733].X = 7.33
	pointsOfFunctionPlot[733].Y = 0.285

	pointsOfFunctionPlot[734].X = 7.34
	pointsOfFunctionPlot[734].Y = 0.284

	pointsOfFunctionPlot[735].X = 7.35
	pointsOfFunctionPlot[735].Y = 0.283

	pointsOfFunctionPlot[736].X = 7.36
	pointsOfFunctionPlot[736].Y = 0.282

	pointsOfFunctionPlot[737].X = 7.37
	pointsOfFunctionPlot[737].Y = 0.281

	pointsOfFunctionPlot[738].X = 7.38
	pointsOfFunctionPlot[738].Y = 0.28

	pointsOfFunctionPlot[739].X = 7.39
	pointsOfFunctionPlot[739].Y = 0.279

	pointsOfFunctionPlot[740].X = 7.40
	pointsOfFunctionPlot[740].Y = 0.278

	pointsOfFunctionPlot[741].X = 7.41
	pointsOfFunctionPlot[741].Y = 0.277

	pointsOfFunctionPlot[742].X = 7.42
	pointsOfFunctionPlot[742].Y = 0.276

	pointsOfFunctionPlot[743].X = 7.43
	pointsOfFunctionPlot[743].Y = 0.275

	pointsOfFunctionPlot[744].X = 7.44
	pointsOfFunctionPlot[744].Y = 0.274

	pointsOfFunctionPlot[745].X = 7.45
	pointsOfFunctionPlot[745].Y = 0.272

	pointsOfFunctionPlot[746].X = 7.46
	pointsOfFunctionPlot[746].Y = 0.271

	pointsOfFunctionPlot[747].X = 7.47
	pointsOfFunctionPlot[747].Y = 0.27

	pointsOfFunctionPlot[748].X = 7.48
	pointsOfFunctionPlot[748].Y = 0.268

	pointsOfFunctionPlot[749].X = 7.49
	pointsOfFunctionPlot[749].Y = 0.267

	pointsOfFunctionPlot[750].X = 7.50
	pointsOfFunctionPlot[750].Y = 0.266

	pointsOfFunctionPlot[751].X = 7.51
	pointsOfFunctionPlot[751].Y = 0.264

	pointsOfFunctionPlot[752].X = 7.52
	pointsOfFunctionPlot[752].Y = 0.263

	pointsOfFunctionPlot[753].X = 7.53
	pointsOfFunctionPlot[753].Y = 0.262

	pointsOfFunctionPlot[754].X = 7.54
	pointsOfFunctionPlot[754].Y = 0.26

	pointsOfFunctionPlot[755].X = 7.55
	pointsOfFunctionPlot[755].Y = 0.259

	pointsOfFunctionPlot[756].X = 7.56
	pointsOfFunctionPlot[756].Y = 0.257

	pointsOfFunctionPlot[757].X = 7.57
	pointsOfFunctionPlot[757].Y = 0.256

	pointsOfFunctionPlot[758].X = 7.58
	pointsOfFunctionPlot[758].Y = 0.254

	pointsOfFunctionPlot[759].X = 7.59
	pointsOfFunctionPlot[759].Y = 0.253

	pointsOfFunctionPlot[760].X = 7.60
	pointsOfFunctionPlot[760].Y = 0.251

	pointsOfFunctionPlot[761].X = 7.61
	pointsOfFunctionPlot[761].Y = 0.249

	pointsOfFunctionPlot[762].X = 7.62
	pointsOfFunctionPlot[762].Y = 0.248

	pointsOfFunctionPlot[763].X = 7.63
	pointsOfFunctionPlot[763].Y = 0.246

	pointsOfFunctionPlot[764].X = 7.64
	pointsOfFunctionPlot[764].Y = 0.245

	pointsOfFunctionPlot[765].X = 7.65
	pointsOfFunctionPlot[765].Y = 0.243

	pointsOfFunctionPlot[766].X = 7.66
	pointsOfFunctionPlot[766].Y = 0.241

	pointsOfFunctionPlot[767].X = 7.67
	pointsOfFunctionPlot[767].Y = 0.239

	pointsOfFunctionPlot[768].X = 7.68
	pointsOfFunctionPlot[768].Y = 0.238

	pointsOfFunctionPlot[769].X = 7.69
	pointsOfFunctionPlot[769].Y = 0.236

	pointsOfFunctionPlot[770].X = 7.70
	pointsOfFunctionPlot[770].Y = 0.234

	pointsOfFunctionPlot[771].X = 7.71
	pointsOfFunctionPlot[771].Y = 0.232

	pointsOfFunctionPlot[772].X = 7.72
	pointsOfFunctionPlot[772].Y = 0.23

	pointsOfFunctionPlot[773].X = 7.73
	pointsOfFunctionPlot[773].Y = 0.229

	pointsOfFunctionPlot[774].X = 7.74
	pointsOfFunctionPlot[774].Y = 0.227

	pointsOfFunctionPlot[775].X = 7.75
	pointsOfFunctionPlot[775].Y = 0.225

	pointsOfFunctionPlot[776].X = 7.76
	pointsOfFunctionPlot[776].Y = 0.223

	pointsOfFunctionPlot[777].X = 7.77
	pointsOfFunctionPlot[777].Y = 0.221

	pointsOfFunctionPlot[778].X = 7.78
	pointsOfFunctionPlot[778].Y = 0.219

	pointsOfFunctionPlot[779].X = 7.79
	pointsOfFunctionPlot[779].Y = 0.217

	pointsOfFunctionPlot[780].X = 7.80
	pointsOfFunctionPlot[780].Y = 0.215

	pointsOfFunctionPlot[781].X = 7.81
	pointsOfFunctionPlot[781].Y = 0.213

	pointsOfFunctionPlot[782].X = 7.82
	pointsOfFunctionPlot[782].Y = 0.211

	pointsOfFunctionPlot[783].X = 7.83
	pointsOfFunctionPlot[783].Y = 0.209

	pointsOfFunctionPlot[784].X = 7.84
	pointsOfFunctionPlot[784].Y = 0.207

	pointsOfFunctionPlot[785].X = 7.85
	pointsOfFunctionPlot[785].Y = 0.205

	pointsOfFunctionPlot[786].X = 7.86
	pointsOfFunctionPlot[786].Y = 0.202

	pointsOfFunctionPlot[787].X = 7.87
	pointsOfFunctionPlot[787].Y = 0.2

	pointsOfFunctionPlot[788].X = 7.88
	pointsOfFunctionPlot[788].Y = 0.198

	pointsOfFunctionPlot[789].X = 7.89
	pointsOfFunctionPlot[789].Y = 0.196

	pointsOfFunctionPlot[790].X = 7.90
	pointsOfFunctionPlot[790].Y = 0.194

	pointsOfFunctionPlot[791].X = 7.91
	pointsOfFunctionPlot[791].Y = 0.192

	pointsOfFunctionPlot[792].X = 7.92
	pointsOfFunctionPlot[792].Y = 0.189

	pointsOfFunctionPlot[793].X = 7.93
	pointsOfFunctionPlot[793].Y = 0.187

	pointsOfFunctionPlot[794].X = 7.94
	pointsOfFunctionPlot[794].Y = 0.185

	pointsOfFunctionPlot[795].X = 7.95
	pointsOfFunctionPlot[795].Y = 0.183

	pointsOfFunctionPlot[796].X = 7.96
	pointsOfFunctionPlot[796].Y = 0.18

	pointsOfFunctionPlot[797].X = 7.97
	pointsOfFunctionPlot[797].Y = 0.178

	pointsOfFunctionPlot[798].X = 7.98
	pointsOfFunctionPlot[798].Y = 0.176

	pointsOfFunctionPlot[799].X = 7.99
	pointsOfFunctionPlot[799].Y = 0.173

	pointsOfFunctionPlot[800].X = 8.0
	pointsOfFunctionPlot[800].Y = 0.171

	pointsOfFunctionPlot[801].X = 8.01
	pointsOfFunctionPlot[801].Y = 0.169

	pointsOfFunctionPlot[802].X = 8.02
	pointsOfFunctionPlot[802].Y = 0.166

	pointsOfFunctionPlot[803].X = 8.03
	pointsOfFunctionPlot[803].Y = 0.164

	pointsOfFunctionPlot[804].X = 8.04
	pointsOfFunctionPlot[804].Y = 0.162

	pointsOfFunctionPlot[805].X = 8.05
	pointsOfFunctionPlot[805].Y = 0.159

	pointsOfFunctionPlot[806].X = 8.06
	pointsOfFunctionPlot[806].Y = 0.157

	pointsOfFunctionPlot[807].X = 8.07
	pointsOfFunctionPlot[807].Y = 0.154

	pointsOfFunctionPlot[808].X = 8.08
	pointsOfFunctionPlot[808].Y = 0.152

	pointsOfFunctionPlot[809].X = 8.09
	pointsOfFunctionPlot[809].Y = 0.149

	pointsOfFunctionPlot[810].X = 8.10
	pointsOfFunctionPlot[810].Y = 0.147

	pointsOfFunctionPlot[811].X = 8.11
	pointsOfFunctionPlot[811].Y = 0.145

	pointsOfFunctionPlot[812].X = 8.12
	pointsOfFunctionPlot[812].Y = 0.142

	pointsOfFunctionPlot[813].X = 8.13
	pointsOfFunctionPlot[813].Y = 0.14

	pointsOfFunctionPlot[814].X = 8.14
	pointsOfFunctionPlot[814].Y = 0.137

	pointsOfFunctionPlot[815].X = 8.15
	pointsOfFunctionPlot[815].Y = 0.134

	pointsOfFunctionPlot[816].X = 8.16
	pointsOfFunctionPlot[816].Y = 0.132

	pointsOfFunctionPlot[817].X = 8.17
	pointsOfFunctionPlot[817].Y = 0.129

	pointsOfFunctionPlot[818].X = 8.18
	pointsOfFunctionPlot[818].Y = 0.127

	pointsOfFunctionPlot[819].X = 8.19
	pointsOfFunctionPlot[819].Y = 0.124

	pointsOfFunctionPlot[820].X = 8.20
	pointsOfFunctionPlot[820].Y = 0.122

	pointsOfFunctionPlot[821].X = 8.21
	pointsOfFunctionPlot[821].Y = 0.119

	pointsOfFunctionPlot[822].X = 8.22
	pointsOfFunctionPlot[822].Y = 0.117

	pointsOfFunctionPlot[823].X = 8.23
	pointsOfFunctionPlot[823].Y = 0.114

	pointsOfFunctionPlot[824].X = 8.24
	pointsOfFunctionPlot[824].Y = 0.111

	pointsOfFunctionPlot[825].X = 8.25
	pointsOfFunctionPlot[825].Y = 0.109

	pointsOfFunctionPlot[826].X = 8.26
	pointsOfFunctionPlot[826].Y = 0.106

	pointsOfFunctionPlot[827].X = 8.27
	pointsOfFunctionPlot[827].Y = 0.103

	pointsOfFunctionPlot[828].X = 8.28
	pointsOfFunctionPlot[828].Y = 0.101

	pointsOfFunctionPlot[829].X = 8.29
	pointsOfFunctionPlot[829].Y = 0.098

	pointsOfFunctionPlot[830].X = 8.30
	pointsOfFunctionPlot[830].Y = 0.096

	pointsOfFunctionPlot[831].X = 8.31
	pointsOfFunctionPlot[831].Y = 0.093

	pointsOfFunctionPlot[832].X = 8.32
	pointsOfFunctionPlot[832].Y = 0.09

	pointsOfFunctionPlot[833].X = 8.33
	pointsOfFunctionPlot[833].Y = 0.088

	pointsOfFunctionPlot[834].X = 8.34
	pointsOfFunctionPlot[834].Y = 0.085

	pointsOfFunctionPlot[835].X = 8.35
	pointsOfFunctionPlot[835].Y = 0.082

	pointsOfFunctionPlot[836].X = 8.36
	pointsOfFunctionPlot[836].Y = 0.079

	pointsOfFunctionPlot[837].X = 8.37
	pointsOfFunctionPlot[837].Y = 0.077

	pointsOfFunctionPlot[838].X = 8.38
	pointsOfFunctionPlot[838].Y = 0.074

	pointsOfFunctionPlot[839].X = 8.39
	pointsOfFunctionPlot[839].Y = 0.071

	pointsOfFunctionPlot[840].X = 8.40
	pointsOfFunctionPlot[840].Y = 0.069

	pointsOfFunctionPlot[841].X = 8.41
	pointsOfFunctionPlot[841].Y = 0.066

	pointsOfFunctionPlot[842].X = 8.42
	pointsOfFunctionPlot[842].Y = 0.063

	pointsOfFunctionPlot[843].X = 8.43
	pointsOfFunctionPlot[843].Y = 0.061

	pointsOfFunctionPlot[844].X = 8.44
	pointsOfFunctionPlot[844].Y = 0.058

	pointsOfFunctionPlot[845].X = 8.45
	pointsOfFunctionPlot[845].Y = 0.055

	pointsOfFunctionPlot[846].X = 8.46
	pointsOfFunctionPlot[846].Y = 0.052

	pointsOfFunctionPlot[847].X = 8.47
	pointsOfFunctionPlot[847].Y = 0.05

	pointsOfFunctionPlot[848].X = 8.48
	pointsOfFunctionPlot[848].Y = 0.047

	pointsOfFunctionPlot[849].X = 8.49
	pointsOfFunctionPlot[849].Y = 0.044

	pointsOfFunctionPlot[850].X = 8.50
	pointsOfFunctionPlot[850].Y = 0.041

	pointsOfFunctionPlot[851].X = 8.51
	pointsOfFunctionPlot[851].Y = 0.039

	pointsOfFunctionPlot[852].X = 8.52
	pointsOfFunctionPlot[852].Y = 0.036

	pointsOfFunctionPlot[853].X = 8.53
	pointsOfFunctionPlot[853].Y = 0.033

	pointsOfFunctionPlot[854].X = 8.54
	pointsOfFunctionPlot[854].Y = 0.031

	pointsOfFunctionPlot[855].X = 8.55
	pointsOfFunctionPlot[855].Y = 0.028

	pointsOfFunctionPlot[856].X = 8.56
	pointsOfFunctionPlot[856].Y = 0.025

	pointsOfFunctionPlot[857].X = 8.57
	pointsOfFunctionPlot[857].Y = 0.022

	pointsOfFunctionPlot[858].X = 8.58
	pointsOfFunctionPlot[858].Y = 0.02

	pointsOfFunctionPlot[859].X = 8.59
	pointsOfFunctionPlot[859].Y = 0.017

	pointsOfFunctionPlot[860].X = 8.60
	pointsOfFunctionPlot[860].Y = 0.014

	pointsOfFunctionPlot[861].X = 8.61
	pointsOfFunctionPlot[861].Y = 0.011

	pointsOfFunctionPlot[862].X = 8.62
	pointsOfFunctionPlot[862].Y = 0.009

	pointsOfFunctionPlot[863].X = 8.63
	pointsOfFunctionPlot[863].Y = 0.006

	pointsOfFunctionPlot[864].X = 8.64
	pointsOfFunctionPlot[864].Y = 0.003

	pointsOfFunctionPlot[865].X = 8.65
	pointsOfFunctionPlot[865].Y = 0.001

	pointsOfFunctionPlot[866].X = 8.66
	pointsOfFunctionPlot[866].Y = -0.001

	pointsOfFunctionPlot[867].X = 8.67
	pointsOfFunctionPlot[867].Y = -0.004

	pointsOfFunctionPlot[868].X = 8.68
	pointsOfFunctionPlot[868].Y = -0.007

	pointsOfFunctionPlot[869].X = 8.69
	pointsOfFunctionPlot[869].Y = -0.009

	pointsOfFunctionPlot[870].X = 8.70
	pointsOfFunctionPlot[870].Y = -0.012

	pointsOfFunctionPlot[871].X = 8.71
	pointsOfFunctionPlot[871].Y = -0.015

	pointsOfFunctionPlot[872].X = 8.72
	pointsOfFunctionPlot[872].Y = -0.017

	pointsOfFunctionPlot[873].X = 8.73
	pointsOfFunctionPlot[873].Y = -0.02

	pointsOfFunctionPlot[874].X = 8.74
	pointsOfFunctionPlot[874].Y = -0.023

	pointsOfFunctionPlot[875].X = 8.75
	pointsOfFunctionPlot[875].Y = -0.025

	pointsOfFunctionPlot[876].X = 8.76
	pointsOfFunctionPlot[876].Y = -0.028

	pointsOfFunctionPlot[877].X = 8.77
	pointsOfFunctionPlot[877].Y = -0.031

	pointsOfFunctionPlot[878].X = 8.78
	pointsOfFunctionPlot[878].Y = -0.033

	pointsOfFunctionPlot[879].X = 8.79
	pointsOfFunctionPlot[879].Y = -0.036

	pointsOfFunctionPlot[880].X = 8.80
	pointsOfFunctionPlot[880].Y = -0.039

	pointsOfFunctionPlot[881].X = 8.81
	pointsOfFunctionPlot[881].Y = -0.041

	pointsOfFunctionPlot[882].X = 8.82
	pointsOfFunctionPlot[882].Y = -0.044

	pointsOfFunctionPlot[883].X = 8.83
	pointsOfFunctionPlot[883].Y = -0.047

	pointsOfFunctionPlot[884].X = 8.84
	pointsOfFunctionPlot[884].Y = -0.049

	pointsOfFunctionPlot[885].X = 8.85
	pointsOfFunctionPlot[885].Y = -0.052

	pointsOfFunctionPlot[886].X = 8.86
	pointsOfFunctionPlot[886].Y = -0.054

	pointsOfFunctionPlot[887].X = 8.87
	pointsOfFunctionPlot[887].Y = -0.057

	pointsOfFunctionPlot[888].X = 8.88
	pointsOfFunctionPlot[888].Y = -0.06

	pointsOfFunctionPlot[889].X = 8.89
	pointsOfFunctionPlot[889].Y = -0.062

	pointsOfFunctionPlot[890].X = 8.90
	pointsOfFunctionPlot[890].Y = -0.065

	pointsOfFunctionPlot[891].X = 8.91
	pointsOfFunctionPlot[891].Y = -0.067

	pointsOfFunctionPlot[892].X = 8.92
	pointsOfFunctionPlot[892].Y = -0.07

	pointsOfFunctionPlot[893].X = 8.93
	pointsOfFunctionPlot[893].Y = -0.072

	pointsOfFunctionPlot[894].X = 8.94
	pointsOfFunctionPlot[894].Y = -0.075

	pointsOfFunctionPlot[895].X = 8.95
	pointsOfFunctionPlot[895].Y = -0.077

	pointsOfFunctionPlot[896].X = 8.96
	pointsOfFunctionPlot[896].Y = -0.08

	pointsOfFunctionPlot[897].X = 8.97
	pointsOfFunctionPlot[897].Y = -0.082

	pointsOfFunctionPlot[898].X = 8.98
	pointsOfFunctionPlot[898].Y = -0.085

	pointsOfFunctionPlot[899].X = 8.99
	pointsOfFunctionPlot[899].Y = -0.087

	pointsOfFunctionPlot[900].X = 9.0
	pointsOfFunctionPlot[900].Y = -0.09

	pointsOfFunctionPlot[901].X = 9.01
	pointsOfFunctionPlot[901].Y = -0.092

	pointsOfFunctionPlot[902].X = 9.02
	pointsOfFunctionPlot[902].Y = -0.095

	pointsOfFunctionPlot[903].X = 9.03
	pointsOfFunctionPlot[903].Y = -0.097

	pointsOfFunctionPlot[904].X = 9.04
	pointsOfFunctionPlot[904].Y = -0.1

	pointsOfFunctionPlot[905].X = 9.05
	pointsOfFunctionPlot[905].Y = -0.102

	pointsOfFunctionPlot[906].X = 9.06
	pointsOfFunctionPlot[906].Y = -0.104

	pointsOfFunctionPlot[907].X = 9.07
	pointsOfFunctionPlot[907].Y = -0.107

	pointsOfFunctionPlot[908].X = 9.08
	pointsOfFunctionPlot[908].Y = -0.109

	pointsOfFunctionPlot[909].X = 9.09
	pointsOfFunctionPlot[909].Y = -0.111

	pointsOfFunctionPlot[910].X = 9.10
	pointsOfFunctionPlot[910].Y = -0.114

	pointsOfFunctionPlot[911].X = 9.11
	pointsOfFunctionPlot[911].Y = -0.116

	pointsOfFunctionPlot[912].X = 9.12
	pointsOfFunctionPlot[912].Y = -0.118

	pointsOfFunctionPlot[913].X = 9.13
	pointsOfFunctionPlot[913].Y = -0.121

	pointsOfFunctionPlot[914].X = 9.14
	pointsOfFunctionPlot[914].Y = -0.123

	pointsOfFunctionPlot[915].X = 9.15
	pointsOfFunctionPlot[915].Y = -0.125

	pointsOfFunctionPlot[916].X = 9.16
	pointsOfFunctionPlot[916].Y = -0.127

	pointsOfFunctionPlot[917].X = 9.17
	pointsOfFunctionPlot[917].Y = -0.13

	pointsOfFunctionPlot[918].X = 9.18
	pointsOfFunctionPlot[918].Y = -0.132

	pointsOfFunctionPlot[919].X = 9.19
	pointsOfFunctionPlot[919].Y = -0.134

	pointsOfFunctionPlot[920].X = 9.20
	pointsOfFunctionPlot[920].Y = -0.136

	pointsOfFunctionPlot[921].X = 9.21
	pointsOfFunctionPlot[921].Y = -0.138

	pointsOfFunctionPlot[922].X = 9.22
	pointsOfFunctionPlot[922].Y = -0.141

	pointsOfFunctionPlot[923].X = 9.23
	pointsOfFunctionPlot[923].Y = -0.143

	pointsOfFunctionPlot[924].X = 9.24
	pointsOfFunctionPlot[924].Y = -0.145

	pointsOfFunctionPlot[925].X = 9.25
	pointsOfFunctionPlot[925].Y = -0.147

	pointsOfFunctionPlot[926].X = 9.26
	pointsOfFunctionPlot[926].Y = -0.149

	pointsOfFunctionPlot[927].X = 9.27
	pointsOfFunctionPlot[927].Y = -0.151

	pointsOfFunctionPlot[928].X = 9.28
	pointsOfFunctionPlot[928].Y = -0.153

	pointsOfFunctionPlot[929].X = 9.29
	pointsOfFunctionPlot[929].Y = -0.155

	pointsOfFunctionPlot[930].X = 9.30
	pointsOfFunctionPlot[930].Y = -0.157

	pointsOfFunctionPlot[931].X = 9.31
	pointsOfFunctionPlot[931].Y = -0.159

	pointsOfFunctionPlot[932].X = 9.32
	pointsOfFunctionPlot[932].Y = -0.161

	pointsOfFunctionPlot[933].X = 9.33
	pointsOfFunctionPlot[933].Y = -0.163

	pointsOfFunctionPlot[934].X = 9.34
	pointsOfFunctionPlot[934].Y = -0.165

	pointsOfFunctionPlot[935].X = 9.35
	pointsOfFunctionPlot[935].Y = -0.167

	pointsOfFunctionPlot[936].X = 9.36
	pointsOfFunctionPlot[936].Y = -0.169

	pointsOfFunctionPlot[937].X = 9.37
	pointsOfFunctionPlot[937].Y = -0.171

	pointsOfFunctionPlot[938].X = 9.38
	pointsOfFunctionPlot[938].Y = -0.173

	pointsOfFunctionPlot[939].X = 9.39
	pointsOfFunctionPlot[939].Y = -0.174

	pointsOfFunctionPlot[940].X = 9.40
	pointsOfFunctionPlot[940].Y = -0.176

	pointsOfFunctionPlot[941].X = 9.41
	pointsOfFunctionPlot[941].Y = -0.178

	pointsOfFunctionPlot[942].X = 9.42
	pointsOfFunctionPlot[942].Y = -0.18

	pointsOfFunctionPlot[943].X = 9.43
	pointsOfFunctionPlot[943].Y = -0.182

	pointsOfFunctionPlot[944].X = 9.44
	pointsOfFunctionPlot[944].Y = -0.183

	pointsOfFunctionPlot[945].X = 9.45
	pointsOfFunctionPlot[945].Y = -0.185

	pointsOfFunctionPlot[946].X = 9.46
	pointsOfFunctionPlot[946].Y = -0.187

	pointsOfFunctionPlot[947].X = 9.47
	pointsOfFunctionPlot[947].Y = -0.188

	pointsOfFunctionPlot[948].X = 9.48
	pointsOfFunctionPlot[948].Y = -0.19

	pointsOfFunctionPlot[949].X = 9.49
	pointsOfFunctionPlot[949].Y = -0.192

	pointsOfFunctionPlot[950].X = 9.50
	pointsOfFunctionPlot[950].Y = -0.193

	pointsOfFunctionPlot[951].X = 9.51
	pointsOfFunctionPlot[951].Y = -0.195

	pointsOfFunctionPlot[952].X = 9.52
	pointsOfFunctionPlot[952].Y = -0.197

	pointsOfFunctionPlot[953].X = 9.53
	pointsOfFunctionPlot[953].Y = -0.198

	pointsOfFunctionPlot[954].X = 9.54
	pointsOfFunctionPlot[954].Y = -0.2

	pointsOfFunctionPlot[955].X = 9.55
	pointsOfFunctionPlot[955].Y = -0.201

	pointsOfFunctionPlot[956].X = 9.56
	pointsOfFunctionPlot[956].Y = -0.203

	pointsOfFunctionPlot[957].X = 9.57
	pointsOfFunctionPlot[957].Y = -0.204

	pointsOfFunctionPlot[958].X = 9.58
	pointsOfFunctionPlot[958].Y = -0.206

	pointsOfFunctionPlot[959].X = 9.59
	pointsOfFunctionPlot[959].Y = -0.207

	pointsOfFunctionPlot[960].X = 9.60
	pointsOfFunctionPlot[960].Y = -0.208

	pointsOfFunctionPlot[961].X = 9.61
	pointsOfFunctionPlot[961].Y = -0.21

	pointsOfFunctionPlot[962].X = 9.62
	pointsOfFunctionPlot[962].Y = -0.211

	pointsOfFunctionPlot[963].X = 9.63
	pointsOfFunctionPlot[963].Y = -0.213

	pointsOfFunctionPlot[964].X = 9.64
	pointsOfFunctionPlot[964].Y = -0.214

	pointsOfFunctionPlot[965].X = 9.65
	pointsOfFunctionPlot[965].Y = -0.215

	pointsOfFunctionPlot[966].X = 9.66
	pointsOfFunctionPlot[966].Y = -0.216

	pointsOfFunctionPlot[967].X = 9.67
	pointsOfFunctionPlot[967].Y = -0.218

	pointsOfFunctionPlot[968].X = 9.68
	pointsOfFunctionPlot[968].Y = -0.219

	pointsOfFunctionPlot[969].X = 9.69
	pointsOfFunctionPlot[969].Y = -0.22

	pointsOfFunctionPlot[970].X = 9.70
	pointsOfFunctionPlot[970].Y = -0.221

	pointsOfFunctionPlot[971].X = 9.71
	pointsOfFunctionPlot[971].Y = -0.222

	pointsOfFunctionPlot[972].X = 9.72
	pointsOfFunctionPlot[972].Y = -0.224

	pointsOfFunctionPlot[973].X = 9.73
	pointsOfFunctionPlot[973].Y = -0.225

	pointsOfFunctionPlot[974].X = 9.74
	pointsOfFunctionPlot[974].Y = -0.226

	pointsOfFunctionPlot[975].X = 9.75
	pointsOfFunctionPlot[975].Y = -0.227

	pointsOfFunctionPlot[976].X = 9.76
	pointsOfFunctionPlot[976].Y = -0.228

	pointsOfFunctionPlot[977].X = 9.77
	pointsOfFunctionPlot[977].Y = -0.229

	pointsOfFunctionPlot[978].X = 9.78
	pointsOfFunctionPlot[978].Y = -0.23

	pointsOfFunctionPlot[979].X = 9.79
	pointsOfFunctionPlot[979].Y = -0.231

	pointsOfFunctionPlot[980].X = 9.80
	pointsOfFunctionPlot[980].Y = -0.232

	pointsOfFunctionPlot[981].X = 9.81
	pointsOfFunctionPlot[981].Y = -0.233

	pointsOfFunctionPlot[982].X = 9.82
	pointsOfFunctionPlot[982].Y = -0.234

	pointsOfFunctionPlot[983].X = 9.83
	pointsOfFunctionPlot[983].Y = -0.234

	pointsOfFunctionPlot[984].X = 9.84
	pointsOfFunctionPlot[984].Y = -0.235

	pointsOfFunctionPlot[985].X = 9.85
	pointsOfFunctionPlot[985].Y = -0.236

	pointsOfFunctionPlot[986].X = 9.86
	pointsOfFunctionPlot[986].Y = -0.237

	pointsOfFunctionPlot[987].X = 9.87
	pointsOfFunctionPlot[987].Y = -0.238

	pointsOfFunctionPlot[988].X = 9.88
	pointsOfFunctionPlot[988].Y = -0.238

	pointsOfFunctionPlot[989].X = 9.89
	pointsOfFunctionPlot[989].Y = -0.239

	pointsOfFunctionPlot[990].X = 9.90
	pointsOfFunctionPlot[990].Y = -0.24

	pointsOfFunctionPlot[991].X = 9.91
	pointsOfFunctionPlot[991].Y = -0.241

	pointsOfFunctionPlot[992].X = 9.92
	pointsOfFunctionPlot[992].Y = -0.241

	pointsOfFunctionPlot[993].X = 9.93
	pointsOfFunctionPlot[993].Y = -0.242

	pointsOfFunctionPlot[994].X = 9.94
	pointsOfFunctionPlot[994].Y = -0.242

	pointsOfFunctionPlot[995].X = 9.95
	pointsOfFunctionPlot[995].Y = -0.243

	pointsOfFunctionPlot[996].X = 9.96
	pointsOfFunctionPlot[996].Y = -0.243

	pointsOfFunctionPlot[997].X = 9.97
	pointsOfFunctionPlot[997].Y = -0.244

	pointsOfFunctionPlot[998].X = 9.98
	pointsOfFunctionPlot[998].Y = -0.245

	pointsOfFunctionPlot[999].X = 9.99
	pointsOfFunctionPlot[999].Y = -0.245

	pointsOfFunctionPlot[1_000].X = 10.0
	pointsOfFunctionPlot[1_000].Y = -0.245

	pointsOfFunctionPlot[1_001].X = 10.01
	pointsOfFunctionPlot[1_001].Y = -0.246

	pointsOfFunctionPlot[1_002].X = 10.02
	pointsOfFunctionPlot[1_002].Y = -0.246

	pointsOfFunctionPlot[1_003].X = 10.03
	pointsOfFunctionPlot[1_003].Y = -0.247

	pointsOfFunctionPlot[1_004].X = 10.04
	pointsOfFunctionPlot[1_004].Y = -0.247

	pointsOfFunctionPlot[1_005].X = 10.05
	pointsOfFunctionPlot[1_005].Y = -0.247

	pointsOfFunctionPlot[1_006].X = 10.06
	pointsOfFunctionPlot[1_006].Y = -0.248

	pointsOfFunctionPlot[1_007].X = 10.07
	pointsOfFunctionPlot[1_007].Y = -0.248

	pointsOfFunctionPlot[1_008].X = 10.08
	pointsOfFunctionPlot[1_008].Y = -0.248

	pointsOfFunctionPlot[1_009].X = 10.09
	pointsOfFunctionPlot[1_009].Y = -0.248

	pointsOfFunctionPlot[1_010].X = 10.10
	pointsOfFunctionPlot[1_010].Y = -0.249

	pointsOfFunctionPlot[1_011].X = 10.11
	pointsOfFunctionPlot[1_011].Y = -0.249

	pointsOfFunctionPlot[1_012].X = 10.12
	pointsOfFunctionPlot[1_012].Y = -0.249

	pointsOfFunctionPlot[1_013].X = 10.13
	pointsOfFunctionPlot[1_013].Y = -0.249

	pointsOfFunctionPlot[1_014].X = 10.14
	pointsOfFunctionPlot[1_014].Y = -0.249

	pointsOfFunctionPlot[1_015].X = 10.15
	pointsOfFunctionPlot[1_015].Y = -0.249

	pointsOfFunctionPlot[1_016].X = 10.16
	pointsOfFunctionPlot[1_016].Y = -0.249

	pointsOfFunctionPlot[1_017].X = 10.17
	pointsOfFunctionPlot[1_017].Y = -0.249

	pointsOfFunctionPlot[1_018].X = 10.18
	pointsOfFunctionPlot[1_018].Y = -0.249

	pointsOfFunctionPlot[1_019].X = 10.19
	pointsOfFunctionPlot[1_019].Y = -0.249

	pointsOfFunctionPlot[1_020].X = 10.20
	pointsOfFunctionPlot[1_020].Y = -0.249

	pointsOfFunctionPlot[1_021].X = 10.21
	pointsOfFunctionPlot[1_021].Y = -0.249

	pointsOfFunctionPlot[1_022].X = 10.22
	pointsOfFunctionPlot[1_022].Y = -0.249

	pointsOfFunctionPlot[1_023].X = 10.23
	pointsOfFunctionPlot[1_023].Y = -0.249

	pointsOfFunctionPlot[1_024].X = 10.24
	pointsOfFunctionPlot[1_024].Y = -0.249

	pointsOfFunctionPlot[1_025].X = 10.25
	pointsOfFunctionPlot[1_025].Y = -0.248

	pointsOfFunctionPlot[1_026].X = 10.26
	pointsOfFunctionPlot[1_026].Y = -0.248

	pointsOfFunctionPlot[1_027].X = 10.27
	pointsOfFunctionPlot[1_027].Y = -0.248

	pointsOfFunctionPlot[1_028].X = 10.28
	pointsOfFunctionPlot[1_028].Y = -0.248

	pointsOfFunctionPlot[1_029].X = 10.29
	pointsOfFunctionPlot[1_029].Y = -0.248

	pointsOfFunctionPlot[1_030].X = 10.30
	pointsOfFunctionPlot[1_030].Y = -0.247

	pointsOfFunctionPlot[1_031].X = 10.31
	pointsOfFunctionPlot[1_031].Y = -0.247

	pointsOfFunctionPlot[1_032].X = 10.32
	pointsOfFunctionPlot[1_032].Y = -0.247

	pointsOfFunctionPlot[1_033].X = 10.33
	pointsOfFunctionPlot[1_033].Y = -0.246

	pointsOfFunctionPlot[1_034].X = 10.34
	pointsOfFunctionPlot[1_034].Y = -0.246

	pointsOfFunctionPlot[1_035].X = 10.35
	pointsOfFunctionPlot[1_035].Y = -0.245

	pointsOfFunctionPlot[1_036].X = 10.36
	pointsOfFunctionPlot[1_036].Y = -0.245

	pointsOfFunctionPlot[1_037].X = 10.37
	pointsOfFunctionPlot[1_037].Y = -0.244

	pointsOfFunctionPlot[1_038].X = 10.38
	pointsOfFunctionPlot[1_038].Y = -0.244

	pointsOfFunctionPlot[1_039].X = 10.39
	pointsOfFunctionPlot[1_039].Y = -0.243

	pointsOfFunctionPlot[1_040].X = 10.40
	pointsOfFunctionPlot[1_040].Y = -0.243

	pointsOfFunctionPlot[1_041].X = 10.41
	pointsOfFunctionPlot[1_041].Y = -0.242

	pointsOfFunctionPlot[1_042].X = 10.42
	pointsOfFunctionPlot[1_042].Y = -0.242

	pointsOfFunctionPlot[1_043].X = 10.43
	pointsOfFunctionPlot[1_043].Y = -0.241

	pointsOfFunctionPlot[1_044].X = 10.44
	pointsOfFunctionPlot[1_044].Y = -0.24

	pointsOfFunctionPlot[1_045].X = 10.45
	pointsOfFunctionPlot[1_045].Y = -0.24

	pointsOfFunctionPlot[1_046].X = 10.46
	pointsOfFunctionPlot[1_046].Y = -0.239

	pointsOfFunctionPlot[1_047].X = 10.47
	pointsOfFunctionPlot[1_047].Y = -0.238

	pointsOfFunctionPlot[1_048].X = 10.48
	pointsOfFunctionPlot[1_048].Y = -0.238

	pointsOfFunctionPlot[1_049].X = 10.49
	pointsOfFunctionPlot[1_049].Y = -0.237

	pointsOfFunctionPlot[1_050].X = 10.50
	pointsOfFunctionPlot[1_050].Y = -0.236

	pointsOfFunctionPlot[1_051].X = 10.51
	pointsOfFunctionPlot[1_051].Y = -0.235

	pointsOfFunctionPlot[1_052].X = 10.52
	pointsOfFunctionPlot[1_052].Y = -0.235

	pointsOfFunctionPlot[1_053].X = 10.53
	pointsOfFunctionPlot[1_053].Y = -0.234

	pointsOfFunctionPlot[1_054].X = 10.54
	pointsOfFunctionPlot[1_054].Y = -0.233

	pointsOfFunctionPlot[1_055].X = 10.55
	pointsOfFunctionPlot[1_055].Y = -0.232

	pointsOfFunctionPlot[1_056].X = 10.56
	pointsOfFunctionPlot[1_056].Y = -0.231

	pointsOfFunctionPlot[1_057].X = 10.57
	pointsOfFunctionPlot[1_057].Y = -0.23

	pointsOfFunctionPlot[1_058].X = 10.58
	pointsOfFunctionPlot[1_058].Y = -0.229

	pointsOfFunctionPlot[1_059].X = 10.59
	pointsOfFunctionPlot[1_059].Y = -0.228

	pointsOfFunctionPlot[1_060].X = 10.60
	pointsOfFunctionPlot[1_060].Y = -0.227

	pointsOfFunctionPlot[1_061].X = 10.61
	pointsOfFunctionPlot[1_061].Y = -0.226

	pointsOfFunctionPlot[1_062].X = 10.62
	pointsOfFunctionPlot[1_062].Y = -0.225

	pointsOfFunctionPlot[1_063].X = 10.63
	pointsOfFunctionPlot[1_063].Y = -0.224

	pointsOfFunctionPlot[1_064].X = 10.64
	pointsOfFunctionPlot[1_064].Y = -0.223

	pointsOfFunctionPlot[1_065].X = 10.65
	pointsOfFunctionPlot[1_065].Y = -0.222

	pointsOfFunctionPlot[1_066].X = 10.66
	pointsOfFunctionPlot[1_066].Y = -0.221

	pointsOfFunctionPlot[1_067].X = 10.67
	pointsOfFunctionPlot[1_067].Y = -0.22

	pointsOfFunctionPlot[1_068].X = 10.68
	pointsOfFunctionPlot[1_068].Y = -0.218

	pointsOfFunctionPlot[1_069].X = 10.69
	pointsOfFunctionPlot[1_069].Y = -0.217

	pointsOfFunctionPlot[1_070].X = 10.70
	pointsOfFunctionPlot[1_070].Y = -0.216

	pointsOfFunctionPlot[1_071].X = 10.71
	pointsOfFunctionPlot[1_071].Y = -0.215

	pointsOfFunctionPlot[1_072].X = 10.72
	pointsOfFunctionPlot[1_072].Y = -0.213

	pointsOfFunctionPlot[1_073].X = 10.73
	pointsOfFunctionPlot[1_073].Y = -0.212

	pointsOfFunctionPlot[1_074].X = 10.74
	pointsOfFunctionPlot[1_074].Y = -0.211

	pointsOfFunctionPlot[1_075].X = 10.75
	pointsOfFunctionPlot[1_075].Y = -0.21

	pointsOfFunctionPlot[1_076].X = 10.76
	pointsOfFunctionPlot[1_076].Y = -0.208

	pointsOfFunctionPlot[1_077].X = 10.77
	pointsOfFunctionPlot[1_077].Y = -0.207

	pointsOfFunctionPlot[1_078].X = 10.78
	pointsOfFunctionPlot[1_078].Y = -0.206

	pointsOfFunctionPlot[1_079].X = 10.79
	pointsOfFunctionPlot[1_079].Y = -0.204

	pointsOfFunctionPlot[1_080].X = 10.80
	pointsOfFunctionPlot[1_080].Y = -0.203

	pointsOfFunctionPlot[1_081].X = 10.81
	pointsOfFunctionPlot[1_081].Y = -0.201

	pointsOfFunctionPlot[1_082].X = 10.82
	pointsOfFunctionPlot[1_082].Y = -0.2

	pointsOfFunctionPlot[1_083].X = 10.83
	pointsOfFunctionPlot[1_083].Y = -0.198

	pointsOfFunctionPlot[1_084].X = 10.84
	pointsOfFunctionPlot[1_084].Y = -0.197

	pointsOfFunctionPlot[1_085].X = 10.85
	pointsOfFunctionPlot[1_085].Y = -0.195

	pointsOfFunctionPlot[1_086].X = 10.86
	pointsOfFunctionPlot[1_086].Y = -0.194

	pointsOfFunctionPlot[1_087].X = 10.87
	pointsOfFunctionPlot[1_087].Y = -0.192

	pointsOfFunctionPlot[1_088].X = 10.88
	pointsOfFunctionPlot[1_088].Y = -0.191

	pointsOfFunctionPlot[1_089].X = 10.89
	pointsOfFunctionPlot[1_089].Y = -0.189

	pointsOfFunctionPlot[1_090].X = 10.90
	pointsOfFunctionPlot[1_090].Y = -0.188

	pointsOfFunctionPlot[1_091].X = 10.91
	pointsOfFunctionPlot[1_091].Y = -0.186

	pointsOfFunctionPlot[1_092].X = 10.92
	pointsOfFunctionPlot[1_092].Y = -0.184

	pointsOfFunctionPlot[1_093].X = 10.93
	pointsOfFunctionPlot[1_093].Y = -0.183

	pointsOfFunctionPlot[1_094].X = 10.94
	pointsOfFunctionPlot[1_094].Y = -0.181

	pointsOfFunctionPlot[1_095].X = 10.95
	pointsOfFunctionPlot[1_095].Y = -0.179

	pointsOfFunctionPlot[1_096].X = 10.96
	pointsOfFunctionPlot[1_096].Y = -0.178

	pointsOfFunctionPlot[1_097].X = 10.97
	pointsOfFunctionPlot[1_097].Y = -0.176

	pointsOfFunctionPlot[1_098].X = 10.98
	pointsOfFunctionPlot[1_098].Y = -0.174

	pointsOfFunctionPlot[1_099].X = 10.99
	pointsOfFunctionPlot[1_099].Y = -0.172

	pointsOfFunctionPlot[1_100].X = 11.0
	pointsOfFunctionPlot[1_100].Y = -0.171

	pointsOfFunctionPlot[1_101].X = 11.01
	pointsOfFunctionPlot[1_101].Y = -0.169

	pointsOfFunctionPlot[1_102].X = 11.02
	pointsOfFunctionPlot[1_102].Y = -0.167

	pointsOfFunctionPlot[1_103].X = 11.03
	pointsOfFunctionPlot[1_103].Y = -0.165

	pointsOfFunctionPlot[1_104].X = 11.04
	pointsOfFunctionPlot[1_104].Y = -0.163

	pointsOfFunctionPlot[1_105].X = 11.05
	pointsOfFunctionPlot[1_105].Y = -0.162

	pointsOfFunctionPlot[1_106].X = 11.06
	pointsOfFunctionPlot[1_106].Y = -0.16

	pointsOfFunctionPlot[1_107].X = 11.07
	pointsOfFunctionPlot[1_107].Y = -0.158

	pointsOfFunctionPlot[1_108].X = 11.08
	pointsOfFunctionPlot[1_108].Y = -0.156

	pointsOfFunctionPlot[1_109].X = 11.09
	pointsOfFunctionPlot[1_109].Y = -0.154

	pointsOfFunctionPlot[1_110].X = 11.10
	pointsOfFunctionPlot[1_110].Y = -0.152

	pointsOfFunctionPlot[1_111].X = 11.11
	pointsOfFunctionPlot[1_111].Y = -0.15

	pointsOfFunctionPlot[1_112].X = 11.12
	pointsOfFunctionPlot[1_112].Y = -0.148

	pointsOfFunctionPlot[1_113].X = 11.13
	pointsOfFunctionPlot[1_113].Y = -0.146

	pointsOfFunctionPlot[1_114].X = 11.14
	pointsOfFunctionPlot[1_114].Y = -0.145

	pointsOfFunctionPlot[1_115].X = 11.15
	pointsOfFunctionPlot[1_115].Y = -0.143

	pointsOfFunctionPlot[1_116].X = 11.16
	pointsOfFunctionPlot[1_116].Y = -0.141

	pointsOfFunctionPlot[1_117].X = 11.17
	pointsOfFunctionPlot[1_117].Y = -0.139

	pointsOfFunctionPlot[1_118].X = 11.18
	pointsOfFunctionPlot[1_118].Y = -0.137

	pointsOfFunctionPlot[1_119].X = 11.19
	pointsOfFunctionPlot[1_119].Y = -0.135

	pointsOfFunctionPlot[1_120].X = 11.20
	pointsOfFunctionPlot[1_120].Y = -0.132

	pointsOfFunctionPlot[1_121].X = 11.21
	pointsOfFunctionPlot[1_121].Y = -0.13

	pointsOfFunctionPlot[1_122].X = 11.22
	pointsOfFunctionPlot[1_122].Y = -0.128

	pointsOfFunctionPlot[1_123].X = 11.23
	pointsOfFunctionPlot[1_123].Y = -0.126

	pointsOfFunctionPlot[1_124].X = 11.24
	pointsOfFunctionPlot[1_124].Y = -0.124

	pointsOfFunctionPlot[1_125].X = 11.25
	pointsOfFunctionPlot[1_125].Y = -0.122

	pointsOfFunctionPlot[1_126].X = 11.26
	pointsOfFunctionPlot[1_126].Y = -0.12

	pointsOfFunctionPlot[1_127].X = 11.27
	pointsOfFunctionPlot[1_127].Y = -0.118

	pointsOfFunctionPlot[1_128].X = 11.28
	pointsOfFunctionPlot[1_128].Y = -0.116

	pointsOfFunctionPlot[1_129].X = 11.29
	pointsOfFunctionPlot[1_129].Y = -0.114

	pointsOfFunctionPlot[1_130].X = 11.30
	pointsOfFunctionPlot[1_130].Y = -0.112

	pointsOfFunctionPlot[1_131].X = 11.31
	pointsOfFunctionPlot[1_131].Y = -0.109

	pointsOfFunctionPlot[1_132].X = 11.32
	pointsOfFunctionPlot[1_132].Y = -0.107

	pointsOfFunctionPlot[1_133].X = 11.33
	pointsOfFunctionPlot[1_133].Y = -0.105

	pointsOfFunctionPlot[1_134].X = 11.34
	pointsOfFunctionPlot[1_134].Y = -0.103

	pointsOfFunctionPlot[1_135].X = 11.35
	pointsOfFunctionPlot[1_135].Y = -0.101

	pointsOfFunctionPlot[1_136].X = 11.36
	pointsOfFunctionPlot[1_136].Y = -0.099

	pointsOfFunctionPlot[1_137].X = 11.37
	pointsOfFunctionPlot[1_137].Y = -0.096

	pointsOfFunctionPlot[1_138].X = 11.38
	pointsOfFunctionPlot[1_138].Y = -0.094

	pointsOfFunctionPlot[1_139].X = 11.39
	pointsOfFunctionPlot[1_139].Y = -0.092

	pointsOfFunctionPlot[1_140].X = 11.40
	pointsOfFunctionPlot[1_140].Y = -0.09

	pointsOfFunctionPlot[1_141].X = 11.41
	pointsOfFunctionPlot[1_141].Y = -0.087

	pointsOfFunctionPlot[1_142].X = 11.42
	pointsOfFunctionPlot[1_142].Y = -0.085

	pointsOfFunctionPlot[1_143].X = 11.43
	pointsOfFunctionPlot[1_143].Y = -0.083

	pointsOfFunctionPlot[1_144].X = 11.44
	pointsOfFunctionPlot[1_144].Y = -0.081

	pointsOfFunctionPlot[1_145].X = 11.45
	pointsOfFunctionPlot[1_145].Y = -0.079

	pointsOfFunctionPlot[1_146].X = 11.46
	pointsOfFunctionPlot[1_146].Y = -0.076

	pointsOfFunctionPlot[1_147].X = 11.47
	pointsOfFunctionPlot[1_147].Y = -0.074

	pointsOfFunctionPlot[1_148].X = 11.48
	pointsOfFunctionPlot[1_148].Y = -0.072

	pointsOfFunctionPlot[1_149].X = 11.49
	pointsOfFunctionPlot[1_149].Y = -0.069

	pointsOfFunctionPlot[1_150].X = 11.50
	pointsOfFunctionPlot[1_150].Y = -0.067

	pointsOfFunctionPlot[1_151].X = 11.51
	pointsOfFunctionPlot[1_151].Y = -0.065

	pointsOfFunctionPlot[1_152].X = 11.52
	pointsOfFunctionPlot[1_152].Y = -0.063

	pointsOfFunctionPlot[1_153].X = 11.53
	pointsOfFunctionPlot[1_153].Y = -0.06

	pointsOfFunctionPlot[1_154].X = 11.54
	pointsOfFunctionPlot[1_154].Y = -0.058

	pointsOfFunctionPlot[1_155].X = 11.55
	pointsOfFunctionPlot[1_155].Y = -0.056

	pointsOfFunctionPlot[1_156].X = 11.56
	pointsOfFunctionPlot[1_156].Y = -0.053

	pointsOfFunctionPlot[1_157].X = 11.57
	pointsOfFunctionPlot[1_157].Y = -0.051

	pointsOfFunctionPlot[1_158].X = 11.58
	pointsOfFunctionPlot[1_158].Y = -0.049

	pointsOfFunctionPlot[1_159].X = 11.59
	pointsOfFunctionPlot[1_159].Y = -0.046

	pointsOfFunctionPlot[1_160].X = 11.60
	pointsOfFunctionPlot[1_160].Y = -0.044

	pointsOfFunctionPlot[1_161].X = 11.61
	pointsOfFunctionPlot[1_161].Y = -0.042

	pointsOfFunctionPlot[1_162].X = 11.62
	pointsOfFunctionPlot[1_162].Y = -0.039

	pointsOfFunctionPlot[1_163].X = 11.63
	pointsOfFunctionPlot[1_163].Y = -0.037

	pointsOfFunctionPlot[1_164].X = 11.64
	pointsOfFunctionPlot[1_164].Y = -0.035

	pointsOfFunctionPlot[1_165].X = 11.65
	pointsOfFunctionPlot[1_165].Y = -0.032

	pointsOfFunctionPlot[1_166].X = 11.66
	pointsOfFunctionPlot[1_166].Y = -0.03

	pointsOfFunctionPlot[1_167].X = 11.67
	pointsOfFunctionPlot[1_167].Y = -0.028

	pointsOfFunctionPlot[1_168].X = 11.68
	pointsOfFunctionPlot[1_168].Y = -0.025

	pointsOfFunctionPlot[1_169].X = 11.69
	pointsOfFunctionPlot[1_169].Y = -0.023

	pointsOfFunctionPlot[1_170].X = 11.70
	pointsOfFunctionPlot[1_170].Y = -0.021

	pointsOfFunctionPlot[1_171].X = 11.71
	pointsOfFunctionPlot[1_171].Y = -0.018

	pointsOfFunctionPlot[1_172].X = 11.72
	pointsOfFunctionPlot[1_172].Y = -0.016

	pointsOfFunctionPlot[1_173].X = 11.73
	pointsOfFunctionPlot[1_173].Y = -0.014

	pointsOfFunctionPlot[1_174].X = 11.74
	pointsOfFunctionPlot[1_174].Y = -0.012

	pointsOfFunctionPlot[1_175].X = 11.75
	pointsOfFunctionPlot[1_175].Y = -0.009

	pointsOfFunctionPlot[1_176].X = 11.76
	pointsOfFunctionPlot[1_176].Y = -0.007

	pointsOfFunctionPlot[1_177].X = 11.77
	pointsOfFunctionPlot[1_177].Y = -0.005

	pointsOfFunctionPlot[1_178].X = 11.78
	pointsOfFunctionPlot[1_178].Y = -0.002

	pointsOfFunctionPlot[1_179].X = 11.79
	pointsOfFunctionPlot[1_179].Y = -0.0

	pointsOfFunctionPlot[1_180].X = 11.80
	pointsOfFunctionPlot[1_180].Y = 0.001

	pointsOfFunctionPlot[1_181].X = 11.81
	pointsOfFunctionPlot[1_181].Y = 0.004

	pointsOfFunctionPlot[1_182].X = 11.82
	pointsOfFunctionPlot[1_182].Y = 0.006

	pointsOfFunctionPlot[1_183].X = 11.83
	pointsOfFunctionPlot[1_183].Y = 0.008

	pointsOfFunctionPlot[1_184].X = 11.84
	pointsOfFunctionPlot[1_184].Y = 0.011

	pointsOfFunctionPlot[1_185].X = 11.85
	pointsOfFunctionPlot[1_185].Y = 0.013

	pointsOfFunctionPlot[1_186].X = 11.86
	pointsOfFunctionPlot[1_186].Y = 0.015

	pointsOfFunctionPlot[1_187].X = 11.87
	pointsOfFunctionPlot[1_187].Y = 0.018

	pointsOfFunctionPlot[1_188].X = 11.88
	pointsOfFunctionPlot[1_188].Y = 0.02

	pointsOfFunctionPlot[1_189].X = 11.89
	pointsOfFunctionPlot[1_189].Y = 0.022

	pointsOfFunctionPlot[1_190].X = 11.90
	pointsOfFunctionPlot[1_190].Y = 0.025

	pointsOfFunctionPlot[1_191].X = 11.91
	pointsOfFunctionPlot[1_191].Y = 0.027

	pointsOfFunctionPlot[1_192].X = 11.92
	pointsOfFunctionPlot[1_192].Y = 0.029

	pointsOfFunctionPlot[1_193].X = 11.93
	pointsOfFunctionPlot[1_193].Y = 0.031

	pointsOfFunctionPlot[1_194].X = 11.94
	pointsOfFunctionPlot[1_194].Y = 0.034

	pointsOfFunctionPlot[1_195].X = 11.95
	pointsOfFunctionPlot[1_195].Y = 0.036

	pointsOfFunctionPlot[1_196].X = 11.96
	pointsOfFunctionPlot[1_196].Y = 0.038

	pointsOfFunctionPlot[1_197].X = 11.97
	pointsOfFunctionPlot[1_197].Y = 0.04

	pointsOfFunctionPlot[1_198].X = 11.98
	pointsOfFunctionPlot[1_198].Y = 0.043

	pointsOfFunctionPlot[1_199].X = 11.99
	pointsOfFunctionPlot[1_199].Y = 0.045

	pointsOfFunctionPlot[1_200].X = 12.0
	pointsOfFunctionPlot[1_200].Y = 0.047

	pointsOfFunctionPlot[1_201].X = 12.01
	pointsOfFunctionPlot[1_201].Y = 0.049

	pointsOfFunctionPlot[1_202].X = 12.02
	pointsOfFunctionPlot[1_202].Y = 0.052

	pointsOfFunctionPlot[1_203].X = 12.03
	pointsOfFunctionPlot[1_203].Y = 0.054

	pointsOfFunctionPlot[1_204].X = 12.04
	pointsOfFunctionPlot[1_204].Y = 0.056

	pointsOfFunctionPlot[1_205].X = 12.05
	pointsOfFunctionPlot[1_205].Y = 0.058

	pointsOfFunctionPlot[1_206].X = 12.06
	pointsOfFunctionPlot[1_206].Y = 0.06

	pointsOfFunctionPlot[1_207].X = 12.07
	pointsOfFunctionPlot[1_207].Y = 0.063

	pointsOfFunctionPlot[1_208].X = 12.08
	pointsOfFunctionPlot[1_208].Y = 0.065

	pointsOfFunctionPlot[1_209].X = 12.09
	pointsOfFunctionPlot[1_209].Y = 0.067

	pointsOfFunctionPlot[1_210].X = 12.10
	pointsOfFunctionPlot[1_210].Y = 0.069

	pointsOfFunctionPlot[1_211].X = 12.11
	pointsOfFunctionPlot[1_211].Y = 0.071

	pointsOfFunctionPlot[1_212].X = 12.12
	pointsOfFunctionPlot[1_212].Y = 0.073

	pointsOfFunctionPlot[1_213].X = 12.13
	pointsOfFunctionPlot[1_213].Y = 0.076

	pointsOfFunctionPlot[1_214].X = 12.14
	pointsOfFunctionPlot[1_214].Y = 0.078

	pointsOfFunctionPlot[1_215].X = 12.15
	pointsOfFunctionPlot[1_215].Y = 0.08

	pointsOfFunctionPlot[1_216].X = 12.16
	pointsOfFunctionPlot[1_216].Y = 0.082

	pointsOfFunctionPlot[1_217].X = 12.17
	pointsOfFunctionPlot[1_217].Y = 0.084

	pointsOfFunctionPlot[1_218].X = 12.18
	pointsOfFunctionPlot[1_218].Y = 0.086

	pointsOfFunctionPlot[1_219].X = 12.19
	pointsOfFunctionPlot[1_219].Y = 0.088

	pointsOfFunctionPlot[1_220].X = 12.20
	pointsOfFunctionPlot[1_220].Y = 0.09

	pointsOfFunctionPlot[1_221].X = 12.21
	pointsOfFunctionPlot[1_221].Y = 0.092

	pointsOfFunctionPlot[1_222].X = 12.22
	pointsOfFunctionPlot[1_222].Y = 0.094

	pointsOfFunctionPlot[1_223].X = 12.23
	pointsOfFunctionPlot[1_223].Y = 0.096

	pointsOfFunctionPlot[1_224].X = 12.24
	pointsOfFunctionPlot[1_224].Y = 0.098

	pointsOfFunctionPlot[1_225].X = 12.25
	pointsOfFunctionPlot[1_225].Y = 0.1

	pointsOfFunctionPlot[1_226].X = 12.26
	pointsOfFunctionPlot[1_226].Y = 0.102

	pointsOfFunctionPlot[1_227].X = 12.27
	pointsOfFunctionPlot[1_227].Y = 0.104

	pointsOfFunctionPlot[1_228].X = 12.28
	pointsOfFunctionPlot[1_228].Y = 0.106

	pointsOfFunctionPlot[1_229].X = 12.29
	pointsOfFunctionPlot[1_229].Y = 0.108

	pointsOfFunctionPlot[1_230].X = 12.30
	pointsOfFunctionPlot[1_230].Y = 0.11

	pointsOfFunctionPlot[1_231].X = 12.31
	pointsOfFunctionPlot[1_231].Y = 0.112

	pointsOfFunctionPlot[1_232].X = 12.32
	pointsOfFunctionPlot[1_232].Y = 0.114

	pointsOfFunctionPlot[1_233].X = 12.33
	pointsOfFunctionPlot[1_233].Y = 0.116

	pointsOfFunctionPlot[1_234].X = 12.34
	pointsOfFunctionPlot[1_234].Y = 0.118

	pointsOfFunctionPlot[1_235].X = 12.35
	pointsOfFunctionPlot[1_235].Y = 0.12

	pointsOfFunctionPlot[1_236].X = 12.36
	pointsOfFunctionPlot[1_236].Y = 0.122

	pointsOfFunctionPlot[1_237].X = 12.37
	pointsOfFunctionPlot[1_237].Y = 0.124

	pointsOfFunctionPlot[1_238].X = 12.38
	pointsOfFunctionPlot[1_238].Y = 0.125

	pointsOfFunctionPlot[1_239].X = 12.39
	pointsOfFunctionPlot[1_239].Y = 0.127

	pointsOfFunctionPlot[1_240].X = 12.40
	pointsOfFunctionPlot[1_240].Y = 0.129

	pointsOfFunctionPlot[1_241].X = 12.41
	pointsOfFunctionPlot[1_241].Y = 0.131

	pointsOfFunctionPlot[1_242].X = 12.42
	pointsOfFunctionPlot[1_242].Y = 0.133

	pointsOfFunctionPlot[1_243].X = 12.43
	pointsOfFunctionPlot[1_243].Y = 0.134

	pointsOfFunctionPlot[1_244].X = 12.44
	pointsOfFunctionPlot[1_244].Y = 0.136

	pointsOfFunctionPlot[1_245].X = 12.45
	pointsOfFunctionPlot[1_245].Y = 0.138

	pointsOfFunctionPlot[1_246].X = 12.46
	pointsOfFunctionPlot[1_246].Y = 0.14

	pointsOfFunctionPlot[1_247].X = 12.47
	pointsOfFunctionPlot[1_247].Y = 0.141

	pointsOfFunctionPlot[1_248].X = 12.48
	pointsOfFunctionPlot[1_248].Y = 0.143

	pointsOfFunctionPlot[1_249].X = 12.49
	pointsOfFunctionPlot[1_249].Y = 0.145

	pointsOfFunctionPlot[1_250].X = 12.50
	pointsOfFunctionPlot[1_250].Y = 0.146

	pointsOfFunctionPlot[1_251].X = 12.51
	pointsOfFunctionPlot[1_251].Y = 0.148

	pointsOfFunctionPlot[1_252].X = 12.52
	pointsOfFunctionPlot[1_252].Y = 0.15

	pointsOfFunctionPlot[1_253].X = 12.53
	pointsOfFunctionPlot[1_253].Y = 0.151

	pointsOfFunctionPlot[1_254].X = 12.54
	pointsOfFunctionPlot[1_254].Y = 0.153

	pointsOfFunctionPlot[1_255].X = 12.55
	pointsOfFunctionPlot[1_255].Y = 0.154

	pointsOfFunctionPlot[1_256].X = 12.56
	pointsOfFunctionPlot[1_256].Y = 0.156

	pointsOfFunctionPlot[1_257].X = 12.57
	pointsOfFunctionPlot[1_257].Y = 0.158

	pointsOfFunctionPlot[1_258].X = 12.58
	pointsOfFunctionPlot[1_258].Y = 0.159

	pointsOfFunctionPlot[1_259].X = 12.59
	pointsOfFunctionPlot[1_259].Y = 0.161

	pointsOfFunctionPlot[1_260].X = 12.60
	pointsOfFunctionPlot[1_260].Y = 0.162

	pointsOfFunctionPlot[1_261].X = 12.61
	pointsOfFunctionPlot[1_261].Y = 0.164

	pointsOfFunctionPlot[1_262].X = 12.62
	pointsOfFunctionPlot[1_262].Y = 0.165

	pointsOfFunctionPlot[1_263].X = 12.63
	pointsOfFunctionPlot[1_263].Y = 0.166

	pointsOfFunctionPlot[1_264].X = 12.64
	pointsOfFunctionPlot[1_264].Y = 0.168

	pointsOfFunctionPlot[1_265].X = 12.65
	pointsOfFunctionPlot[1_265].Y = 0.169

	pointsOfFunctionPlot[1_266].X = 12.66
	pointsOfFunctionPlot[1_266].Y = 0.171

	pointsOfFunctionPlot[1_267].X = 12.67
	pointsOfFunctionPlot[1_267].Y = 0.172

	pointsOfFunctionPlot[1_268].X = 12.68
	pointsOfFunctionPlot[1_268].Y = 0.173

	pointsOfFunctionPlot[1_269].X = 12.69
	pointsOfFunctionPlot[1_269].Y = 0.175

	pointsOfFunctionPlot[1_270].X = 12.70
	pointsOfFunctionPlot[1_270].Y = 0.176

	pointsOfFunctionPlot[1_271].X = 12.71
	pointsOfFunctionPlot[1_271].Y = 0.177

	pointsOfFunctionPlot[1_272].X = 12.72
	pointsOfFunctionPlot[1_272].Y = 0.179

	pointsOfFunctionPlot[1_273].X = 12.73
	pointsOfFunctionPlot[1_273].Y = 0.18

	pointsOfFunctionPlot[1_274].X = 12.74
	pointsOfFunctionPlot[1_274].Y = 0.181

	pointsOfFunctionPlot[1_275].X = 12.75
	pointsOfFunctionPlot[1_275].Y = 0.182

	pointsOfFunctionPlot[1_276].X = 12.76
	pointsOfFunctionPlot[1_276].Y = 0.184

	pointsOfFunctionPlot[1_277].X = 12.77
	pointsOfFunctionPlot[1_277].Y = 0.185

	pointsOfFunctionPlot[1_278].X = 12.78
	pointsOfFunctionPlot[1_278].Y = 0.186

	pointsOfFunctionPlot[1_279].X = 12.79
	pointsOfFunctionPlot[1_279].Y = 0.187

	pointsOfFunctionPlot[1_280].X = 12.80
	pointsOfFunctionPlot[1_280].Y = 0.188

	pointsOfFunctionPlot[1_281].X = 12.81
	pointsOfFunctionPlot[1_281].Y = 0.189

	pointsOfFunctionPlot[1_282].X = 12.82
	pointsOfFunctionPlot[1_282].Y = 0.19

	pointsOfFunctionPlot[1_283].X = 12.83
	pointsOfFunctionPlot[1_283].Y = 0.191

	pointsOfFunctionPlot[1_284].X = 12.84
	pointsOfFunctionPlot[1_284].Y = 0.192

	pointsOfFunctionPlot[1_285].X = 12.85
	pointsOfFunctionPlot[1_285].Y = 0.194

	pointsOfFunctionPlot[1_286].X = 12.86
	pointsOfFunctionPlot[1_286].Y = 0.195

	pointsOfFunctionPlot[1_287].X = 12.87
	pointsOfFunctionPlot[1_287].Y = 0.196

	pointsOfFunctionPlot[1_288].X = 12.88
	pointsOfFunctionPlot[1_288].Y = 0.196

	pointsOfFunctionPlot[1_289].X = 12.89
	pointsOfFunctionPlot[1_289].Y = 0.197

	pointsOfFunctionPlot[1_290].X = 12.90
	pointsOfFunctionPlot[1_290].Y = 0.198

	pointsOfFunctionPlot[1_291].X = 12.91
	pointsOfFunctionPlot[1_291].Y = 0.199

	pointsOfFunctionPlot[1_292].X = 12.92
	pointsOfFunctionPlot[1_292].Y = 0.2

	pointsOfFunctionPlot[1_293].X = 12.93
	pointsOfFunctionPlot[1_293].Y = 0.201

	pointsOfFunctionPlot[1_294].X = 12.94
	pointsOfFunctionPlot[1_294].Y = 0.202

	pointsOfFunctionPlot[1_295].X = 12.95
	pointsOfFunctionPlot[1_295].Y = 0.203

	pointsOfFunctionPlot[1_296].X = 12.96
	pointsOfFunctionPlot[1_296].Y = 0.203

	pointsOfFunctionPlot[1_297].X = 12.97
	pointsOfFunctionPlot[1_297].Y = 0.204

	pointsOfFunctionPlot[1_298].X = 12.98
	pointsOfFunctionPlot[1_298].Y = 0.205

	pointsOfFunctionPlot[1_299].X = 12.99
	pointsOfFunctionPlot[1_299].Y = 0.206

	pointsOfFunctionPlot[1_300].X = 13.0
	pointsOfFunctionPlot[1_300].Y = 0.206

	pointsOfFunctionPlot[1_301].X = 13.01
	pointsOfFunctionPlot[1_301].Y = 0.207

	pointsOfFunctionPlot[1_302].X = 13.02
	pointsOfFunctionPlot[1_302].Y = 0.208

	pointsOfFunctionPlot[1_303].X = 13.03
	pointsOfFunctionPlot[1_303].Y = 0.208

	pointsOfFunctionPlot[1_304].X = 13.04
	pointsOfFunctionPlot[1_304].Y = 0.209

	pointsOfFunctionPlot[1_305].X = 13.05
	pointsOfFunctionPlot[1_305].Y = 0.21

	pointsOfFunctionPlot[1_306].X = 13.06
	pointsOfFunctionPlot[1_306].Y = 0.21

	pointsOfFunctionPlot[1_307].X = 13.07
	pointsOfFunctionPlot[1_307].Y = 0.211

	pointsOfFunctionPlot[1_308].X = 13.08
	pointsOfFunctionPlot[1_308].Y = 0.211

	pointsOfFunctionPlot[1_309].X = 13.09
	pointsOfFunctionPlot[1_309].Y = 0.212

	pointsOfFunctionPlot[1_310].X = 13.10
	pointsOfFunctionPlot[1_310].Y = 0.212

	pointsOfFunctionPlot[1_311].X = 13.11
	pointsOfFunctionPlot[1_311].Y = 0.213

	pointsOfFunctionPlot[1_312].X = 13.12
	pointsOfFunctionPlot[1_312].Y = 0.213

	pointsOfFunctionPlot[1_313].X = 13.13
	pointsOfFunctionPlot[1_313].Y = 0.214

	pointsOfFunctionPlot[1_314].X = 13.14
	pointsOfFunctionPlot[1_314].Y = 0.214

	pointsOfFunctionPlot[1_315].X = 13.15
	pointsOfFunctionPlot[1_315].Y = 0.215

	pointsOfFunctionPlot[1_316].X = 13.16
	pointsOfFunctionPlot[1_316].Y = 0.215

	pointsOfFunctionPlot[1_317].X = 13.17
	pointsOfFunctionPlot[1_317].Y = 0.215

	pointsOfFunctionPlot[1_318].X = 13.18
	pointsOfFunctionPlot[1_318].Y = 0.216

	pointsOfFunctionPlot[1_319].X = 13.19
	pointsOfFunctionPlot[1_319].Y = 0.216

	pointsOfFunctionPlot[1_320].X = 13.20
	pointsOfFunctionPlot[1_320].Y = 0.216

	pointsOfFunctionPlot[1_321].X = 13.21
	pointsOfFunctionPlot[1_321].Y = 0.216

	pointsOfFunctionPlot[1_322].X = 13.22
	pointsOfFunctionPlot[1_322].Y = 0.217

	pointsOfFunctionPlot[1_323].X = 13.23
	pointsOfFunctionPlot[1_323].Y = 0.217

	pointsOfFunctionPlot[1_324].X = 13.24
	pointsOfFunctionPlot[1_324].Y = 0.217

	pointsOfFunctionPlot[1_325].X = 13.25
	pointsOfFunctionPlot[1_325].Y = 0.217

	pointsOfFunctionPlot[1_326].X = 13.26
	pointsOfFunctionPlot[1_326].Y = 0.217

	pointsOfFunctionPlot[1_327].X = 13.27
	pointsOfFunctionPlot[1_327].Y = 0.218

	pointsOfFunctionPlot[1_328].X = 13.28
	pointsOfFunctionPlot[1_328].Y = 0.218

	pointsOfFunctionPlot[1_329].X = 13.29
	pointsOfFunctionPlot[1_329].Y = 0.218

	pointsOfFunctionPlot[1_330].X = 13.30
	pointsOfFunctionPlot[1_330].Y = 0.218

	pointsOfFunctionPlot[1_331].X = 13.31
	pointsOfFunctionPlot[1_331].Y = 0.218

	pointsOfFunctionPlot[1_332].X = 13.32
	pointsOfFunctionPlot[1_332].Y = 0.218

	pointsOfFunctionPlot[1_333].X = 13.33
	pointsOfFunctionPlot[1_333].Y = 0.218

	pointsOfFunctionPlot[1_334].X = 13.34
	pointsOfFunctionPlot[1_334].Y = 0.218

	pointsOfFunctionPlot[1_335].X = 13.35
	pointsOfFunctionPlot[1_335].Y = 0.218

	pointsOfFunctionPlot[1_336].X = 13.36
	pointsOfFunctionPlot[1_336].Y = 0.218

	pointsOfFunctionPlot[1_337].X = 13.37
	pointsOfFunctionPlot[1_337].Y = 0.218

	pointsOfFunctionPlot[1_338].X = 13.38
	pointsOfFunctionPlot[1_338].Y = 0.218

	pointsOfFunctionPlot[1_339].X = 13.39
	pointsOfFunctionPlot[1_339].Y = 0.217

	pointsOfFunctionPlot[1_340].X = 13.40
	pointsOfFunctionPlot[1_340].Y = 0.217

	pointsOfFunctionPlot[1_341].X = 13.41
	pointsOfFunctionPlot[1_341].Y = 0.217

	pointsOfFunctionPlot[1_342].X = 13.42
	pointsOfFunctionPlot[1_342].Y = 0.217

	pointsOfFunctionPlot[1_343].X = 13.43
	pointsOfFunctionPlot[1_343].Y = 0.217

	pointsOfFunctionPlot[1_344].X = 13.44
	pointsOfFunctionPlot[1_344].Y = 0.216

	pointsOfFunctionPlot[1_345].X = 13.45
	pointsOfFunctionPlot[1_345].Y = 0.216

	pointsOfFunctionPlot[1_346].X = 13.46
	pointsOfFunctionPlot[1_346].Y = 0.216

	pointsOfFunctionPlot[1_347].X = 13.47
	pointsOfFunctionPlot[1_347].Y = 0.216

	pointsOfFunctionPlot[1_348].X = 13.48
	pointsOfFunctionPlot[1_348].Y = 0.215

	pointsOfFunctionPlot[1_349].X = 13.49
	pointsOfFunctionPlot[1_349].Y = 0.215

	pointsOfFunctionPlot[1_350].X = 13.50
	pointsOfFunctionPlot[1_350].Y = 0.214

	pointsOfFunctionPlot[1_351].X = 13.51
	pointsOfFunctionPlot[1_351].Y = 0.214

	pointsOfFunctionPlot[1_352].X = 13.52
	pointsOfFunctionPlot[1_352].Y = 0.214

	pointsOfFunctionPlot[1_353].X = 13.53
	pointsOfFunctionPlot[1_353].Y = 0.213

	pointsOfFunctionPlot[1_354].X = 13.54
	pointsOfFunctionPlot[1_354].Y = 0.213

	pointsOfFunctionPlot[1_355].X = 13.55
	pointsOfFunctionPlot[1_355].Y = 0.212

	pointsOfFunctionPlot[1_356].X = 13.56
	pointsOfFunctionPlot[1_356].Y = 0.212

	pointsOfFunctionPlot[1_357].X = 13.57
	pointsOfFunctionPlot[1_357].Y = 0.211

	pointsOfFunctionPlot[1_358].X = 13.58
	pointsOfFunctionPlot[1_358].Y = 0.211

	pointsOfFunctionPlot[1_359].X = 13.59
	pointsOfFunctionPlot[1_359].Y = 0.21

	pointsOfFunctionPlot[1_360].X = 13.60
	pointsOfFunctionPlot[1_360].Y = 0.21

	pointsOfFunctionPlot[1_361].X = 13.61
	pointsOfFunctionPlot[1_361].Y = 0.209

	pointsOfFunctionPlot[1_362].X = 13.62
	pointsOfFunctionPlot[1_362].Y = 0.208

	pointsOfFunctionPlot[1_363].X = 13.63
	pointsOfFunctionPlot[1_363].Y = 0.208

	pointsOfFunctionPlot[1_364].X = 13.64
	pointsOfFunctionPlot[1_364].Y = 0.207

	pointsOfFunctionPlot[1_365].X = 13.65
	pointsOfFunctionPlot[1_365].Y = 0.206

	pointsOfFunctionPlot[1_366].X = 13.66
	pointsOfFunctionPlot[1_366].Y = 0.206

	pointsOfFunctionPlot[1_367].X = 13.67
	pointsOfFunctionPlot[1_367].Y = 0.205

	pointsOfFunctionPlot[1_368].X = 13.68
	pointsOfFunctionPlot[1_368].Y = 0.204

	pointsOfFunctionPlot[1_369].X = 13.69
	pointsOfFunctionPlot[1_369].Y = 0.204

	pointsOfFunctionPlot[1_370].X = 13.70
	pointsOfFunctionPlot[1_370].Y = 0.203

	pointsOfFunctionPlot[1_371].X = 13.71
	pointsOfFunctionPlot[1_371].Y = 0.202

	pointsOfFunctionPlot[1_372].X = 13.72
	pointsOfFunctionPlot[1_372].Y = 0.201

	pointsOfFunctionPlot[1_373].X = 13.73
	pointsOfFunctionPlot[1_373].Y = 0.2

	pointsOfFunctionPlot[1_374].X = 13.74
	pointsOfFunctionPlot[1_374].Y = 0.199

	pointsOfFunctionPlot[1_375].X = 13.75
	pointsOfFunctionPlot[1_375].Y = 0.199

	pointsOfFunctionPlot[1_376].X = 13.76
	pointsOfFunctionPlot[1_376].Y = 0.198

	pointsOfFunctionPlot[1_377].X = 13.77
	pointsOfFunctionPlot[1_377].Y = 0.197

	pointsOfFunctionPlot[1_378].X = 13.78
	pointsOfFunctionPlot[1_378].Y = 0.196

	pointsOfFunctionPlot[1_379].X = 13.79
	pointsOfFunctionPlot[1_379].Y = 0.195

	pointsOfFunctionPlot[1_380].X = 13.80
	pointsOfFunctionPlot[1_380].Y = 0.194

	pointsOfFunctionPlot[1_381].X = 13.81
	pointsOfFunctionPlot[1_381].Y = 0.193

	pointsOfFunctionPlot[1_382].X = 13.82
	pointsOfFunctionPlot[1_382].Y = 0.192

	pointsOfFunctionPlot[1_383].X = 13.83
	pointsOfFunctionPlot[1_383].Y = 0.191

	pointsOfFunctionPlot[1_384].X = 13.84
	pointsOfFunctionPlot[1_384].Y = 0.19

	pointsOfFunctionPlot[1_385].X = 13.85
	pointsOfFunctionPlot[1_385].Y = 0.189

	pointsOfFunctionPlot[1_386].X = 13.86
	pointsOfFunctionPlot[1_386].Y = 0.188

	pointsOfFunctionPlot[1_387].X = 13.87
	pointsOfFunctionPlot[1_387].Y = 0.186

	pointsOfFunctionPlot[1_388].X = 13.88
	pointsOfFunctionPlot[1_388].Y = 0.185

	pointsOfFunctionPlot[1_389].X = 13.89
	pointsOfFunctionPlot[1_389].Y = 0.184

	pointsOfFunctionPlot[1_390].X = 13.90
	pointsOfFunctionPlot[1_390].Y = 0.183

	pointsOfFunctionPlot[1_391].X = 13.91
	pointsOfFunctionPlot[1_391].Y = 0.182

	pointsOfFunctionPlot[1_392].X = 13.92
	pointsOfFunctionPlot[1_392].Y = 0.181

	pointsOfFunctionPlot[1_393].X = 13.93
	pointsOfFunctionPlot[1_393].Y = 0.18

	pointsOfFunctionPlot[1_394].X = 13.94
	pointsOfFunctionPlot[1_394].Y = 0.178

	pointsOfFunctionPlot[1_395].X = 13.95
	pointsOfFunctionPlot[1_395].Y = 0.177

	pointsOfFunctionPlot[1_396].X = 13.96
	pointsOfFunctionPlot[1_396].Y = 0.176

	pointsOfFunctionPlot[1_397].X = 13.97
	pointsOfFunctionPlot[1_397].Y = 0.175

	pointsOfFunctionPlot[1_398].X = 13.98
	pointsOfFunctionPlot[1_398].Y = 0.173

	pointsOfFunctionPlot[1_399].X = 13.99
	pointsOfFunctionPlot[1_399].Y = 0.172

	pointsOfFunctionPlot[1_400].X = 14.0
	pointsOfFunctionPlot[1_400].Y = 0.171

	pointsOfFunctionPlot[1_401].X = 14.01
	pointsOfFunctionPlot[1_401].Y = 0.169

	pointsOfFunctionPlot[1_402].X = 14.02
	pointsOfFunctionPlot[1_402].Y = 0.168

	pointsOfFunctionPlot[1_403].X = 14.03
	pointsOfFunctionPlot[1_403].Y = 0.167

	pointsOfFunctionPlot[1_404].X = 14.04
	pointsOfFunctionPlot[1_404].Y = 0.165

	pointsOfFunctionPlot[1_405].X = 14.05
	pointsOfFunctionPlot[1_405].Y = 0.164

	pointsOfFunctionPlot[1_406].X = 14.06
	pointsOfFunctionPlot[1_406].Y = 0.162

	pointsOfFunctionPlot[1_407].X = 14.07
	pointsOfFunctionPlot[1_407].Y = 0.161

	pointsOfFunctionPlot[1_408].X = 14.08
	pointsOfFunctionPlot[1_408].Y = 0.159

	pointsOfFunctionPlot[1_409].X = 14.09
	pointsOfFunctionPlot[1_409].Y = 0.158

	pointsOfFunctionPlot[1_410].X = 14.10
	pointsOfFunctionPlot[1_410].Y = 0.156

	pointsOfFunctionPlot[1_411].X = 14.11
	pointsOfFunctionPlot[1_411].Y = 0.155

	pointsOfFunctionPlot[1_412].X = 14.12
	pointsOfFunctionPlot[1_412].Y = 0.153

	pointsOfFunctionPlot[1_413].X = 14.13
	pointsOfFunctionPlot[1_413].Y = 0.152

	pointsOfFunctionPlot[1_414].X = 14.14
	pointsOfFunctionPlot[1_414].Y = 0.15

	pointsOfFunctionPlot[1_415].X = 14.15
	pointsOfFunctionPlot[1_415].Y = 0.149

	pointsOfFunctionPlot[1_416].X = 14.16
	pointsOfFunctionPlot[1_416].Y = 0.147

	pointsOfFunctionPlot[1_417].X = 14.17
	pointsOfFunctionPlot[1_417].Y = 0.146

	pointsOfFunctionPlot[1_418].X = 14.18
	pointsOfFunctionPlot[1_418].Y = 0.144

	pointsOfFunctionPlot[1_419].X = 14.19
	pointsOfFunctionPlot[1_419].Y = 0.142

	pointsOfFunctionPlot[1_420].X = 14.20
	pointsOfFunctionPlot[1_420].Y = 0.141

	pointsOfFunctionPlot[1_421].X = 14.21
	pointsOfFunctionPlot[1_421].Y = 0.139

	pointsOfFunctionPlot[1_422].X = 14.22
	pointsOfFunctionPlot[1_422].Y = 0.138

	pointsOfFunctionPlot[1_423].X = 14.23
	pointsOfFunctionPlot[1_423].Y = 0.136

	pointsOfFunctionPlot[1_424].X = 14.24
	pointsOfFunctionPlot[1_424].Y = 0.134

	pointsOfFunctionPlot[1_425].X = 14.25
	pointsOfFunctionPlot[1_425].Y = 0.133

	pointsOfFunctionPlot[1_426].X = 14.26
	pointsOfFunctionPlot[1_426].Y = 0.131

	pointsOfFunctionPlot[1_427].X = 14.27
	pointsOfFunctionPlot[1_427].Y = 0.129

	pointsOfFunctionPlot[1_428].X = 14.28
	pointsOfFunctionPlot[1_428].Y = 0.127

	pointsOfFunctionPlot[1_429].X = 14.29
	pointsOfFunctionPlot[1_429].Y = 0.126

	pointsOfFunctionPlot[1_430].X = 14.30
	pointsOfFunctionPlot[1_430].Y = 0.124

	pointsOfFunctionPlot[1_431].X = 14.31
	pointsOfFunctionPlot[1_431].Y = 0.122

	pointsOfFunctionPlot[1_432].X = 14.32
	pointsOfFunctionPlot[1_432].Y = 0.12

	pointsOfFunctionPlot[1_433].X = 14.33
	pointsOfFunctionPlot[1_433].Y = 0.119

	pointsOfFunctionPlot[1_434].X = 14.34
	pointsOfFunctionPlot[1_434].Y = 0.117

	pointsOfFunctionPlot[1_435].X = 14.35
	pointsOfFunctionPlot[1_435].Y = 0.115

	pointsOfFunctionPlot[1_436].X = 14.36
	pointsOfFunctionPlot[1_436].Y = 0.113

	pointsOfFunctionPlot[1_437].X = 14.37
	pointsOfFunctionPlot[1_437].Y = 0.111

	pointsOfFunctionPlot[1_438].X = 14.38
	pointsOfFunctionPlot[1_438].Y = 0.11

	pointsOfFunctionPlot[1_439].X = 14.39
	pointsOfFunctionPlot[1_439].Y = 0.108

	pointsOfFunctionPlot[1_440].X = 14.40
	pointsOfFunctionPlot[1_440].Y = 0.106

	pointsOfFunctionPlot[1_441].X = 14.41
	pointsOfFunctionPlot[1_441].Y = 0.104

	pointsOfFunctionPlot[1_442].X = 14.42
	pointsOfFunctionPlot[1_442].Y = 0.102

	pointsOfFunctionPlot[1_443].X = 14.43
	pointsOfFunctionPlot[1_443].Y = 0.1

	pointsOfFunctionPlot[1_444].X = 14.44
	pointsOfFunctionPlot[1_444].Y = 0.099

	pointsOfFunctionPlot[1_445].X = 14.45
	pointsOfFunctionPlot[1_445].Y = 0.097

	pointsOfFunctionPlot[1_446].X = 14.46
	pointsOfFunctionPlot[1_446].Y = 0.095

	pointsOfFunctionPlot[1_447].X = 14.47
	pointsOfFunctionPlot[1_447].Y = 0.093

	pointsOfFunctionPlot[1_448].X = 14.48
	pointsOfFunctionPlot[1_448].Y = 0.091

	pointsOfFunctionPlot[1_449].X = 14.49
	pointsOfFunctionPlot[1_449].Y = 0.089

	pointsOfFunctionPlot[1_450].X = 14.50
	pointsOfFunctionPlot[1_450].Y = 0.087

	pointsOfFunctionPlot[1_451].X = 14.51
	pointsOfFunctionPlot[1_451].Y = 0.085

	pointsOfFunctionPlot[1_452].X = 14.52
	pointsOfFunctionPlot[1_452].Y = 0.083

	pointsOfFunctionPlot[1_453].X = 14.53
	pointsOfFunctionPlot[1_453].Y = 0.081

	pointsOfFunctionPlot[1_454].X = 14.54
	pointsOfFunctionPlot[1_454].Y = 0.079

	pointsOfFunctionPlot[1_455].X = 14.55
	pointsOfFunctionPlot[1_455].Y = 0.077

	pointsOfFunctionPlot[1_456].X = 14.56
	pointsOfFunctionPlot[1_456].Y = 0.075

	pointsOfFunctionPlot[1_457].X = 14.57
	pointsOfFunctionPlot[1_457].Y = 0.073

	pointsOfFunctionPlot[1_458].X = 14.58
	pointsOfFunctionPlot[1_458].Y = 0.071

	pointsOfFunctionPlot[1_459].X = 14.59
	pointsOfFunctionPlot[1_459].Y = 0.069

	pointsOfFunctionPlot[1_460].X = 14.60
	pointsOfFunctionPlot[1_460].Y = 0.067

	pointsOfFunctionPlot[1_461].X = 14.61
	pointsOfFunctionPlot[1_461].Y = 0.065

	pointsOfFunctionPlot[1_462].X = 14.62
	pointsOfFunctionPlot[1_462].Y = 0.063

	pointsOfFunctionPlot[1_463].X = 14.63
	pointsOfFunctionPlot[1_463].Y = 0.061

	pointsOfFunctionPlot[1_464].X = 14.64
	pointsOfFunctionPlot[1_464].Y = 0.059

	pointsOfFunctionPlot[1_465].X = 14.65
	pointsOfFunctionPlot[1_465].Y = 0.057

	pointsOfFunctionPlot[1_466].X = 14.66
	pointsOfFunctionPlot[1_466].Y = 0.055

	pointsOfFunctionPlot[1_467].X = 14.67
	pointsOfFunctionPlot[1_467].Y = 0.053

	pointsOfFunctionPlot[1_468].X = 14.68
	pointsOfFunctionPlot[1_468].Y = 0.051

	pointsOfFunctionPlot[1_469].X = 14.69
	pointsOfFunctionPlot[1_469].Y = 0.049

	pointsOfFunctionPlot[1_470].X = 14.70
	pointsOfFunctionPlot[1_470].Y = 0.047

	pointsOfFunctionPlot[1_471].X = 14.71
	pointsOfFunctionPlot[1_471].Y = 0.045

	pointsOfFunctionPlot[1_472].X = 14.72
	pointsOfFunctionPlot[1_472].Y = 0.043

	pointsOfFunctionPlot[1_473].X = 14.73
	pointsOfFunctionPlot[1_473].Y = 0.041

	pointsOfFunctionPlot[1_474].X = 14.74
	pointsOfFunctionPlot[1_474].Y = 0.039

	pointsOfFunctionPlot[1_475].X = 14.75
	pointsOfFunctionPlot[1_475].Y = 0.037

	pointsOfFunctionPlot[1_476].X = 14.76
	pointsOfFunctionPlot[1_476].Y = 0.035

	pointsOfFunctionPlot[1_477].X = 14.77
	pointsOfFunctionPlot[1_477].Y = 0.033

	pointsOfFunctionPlot[1_478].X = 14.78
	pointsOfFunctionPlot[1_478].Y = 0.031

	pointsOfFunctionPlot[1_479].X = 14.79
	pointsOfFunctionPlot[1_479].Y = 0.029

	pointsOfFunctionPlot[1_480].X = 14.80
	pointsOfFunctionPlot[1_480].Y = 0.027

	pointsOfFunctionPlot[1_481].X = 14.81
	pointsOfFunctionPlot[1_481].Y = 0.025

	pointsOfFunctionPlot[1_482].X = 14.82
	pointsOfFunctionPlot[1_482].Y = 0.022

	pointsOfFunctionPlot[1_483].X = 14.83
	pointsOfFunctionPlot[1_483].Y = 0.02

	pointsOfFunctionPlot[1_484].X = 14.84
	pointsOfFunctionPlot[1_484].Y = 0.018

	pointsOfFunctionPlot[1_485].X = 14.85
	pointsOfFunctionPlot[1_485].Y = 0.016

	pointsOfFunctionPlot[1_486].X = 14.86
	pointsOfFunctionPlot[1_486].Y = 0.014

	pointsOfFunctionPlot[1_487].X = 14.87
	pointsOfFunctionPlot[1_487].Y = 0.012

	pointsOfFunctionPlot[1_488].X = 14.88
	pointsOfFunctionPlot[1_488].Y = 0.01

	pointsOfFunctionPlot[1_489].X = 14.89
	pointsOfFunctionPlot[1_489].Y = 0.008

	pointsOfFunctionPlot[1_490].X = 14.90
	pointsOfFunctionPlot[1_490].Y = 0.006

	pointsOfFunctionPlot[1_491].X = 14.91
	pointsOfFunctionPlot[1_491].Y = 0.004

	pointsOfFunctionPlot[1_492].X = 14.92
	pointsOfFunctionPlot[1_492].Y = 0.002

	pointsOfFunctionPlot[1_493].X = 14.93
	pointsOfFunctionPlot[1_493].Y = 0.0

	pointsOfFunctionPlot[1_494].X = 14.94
	pointsOfFunctionPlot[1_494].Y = -0.001

	pointsOfFunctionPlot[1_495].X = 14.95
	pointsOfFunctionPlot[1_495].Y = -0.003

	pointsOfFunctionPlot[1_496].X = 14.96
	pointsOfFunctionPlot[1_496].Y = -0.006

	pointsOfFunctionPlot[1_497].X = 14.97
	pointsOfFunctionPlot[1_497].Y = -0.008

	pointsOfFunctionPlot[1_498].X = 14.98
	pointsOfFunctionPlot[1_498].Y = -0.01

	pointsOfFunctionPlot[1_499].X = 14.99
	pointsOfFunctionPlot[1_499].Y = -0.012

	pointsOfFunctionPlot[1_500].X = 15.0
	pointsOfFunctionPlot[1_500].Y = -0.014

	pointsOfFunctionPlot[1_501].X = 15.01
	pointsOfFunctionPlot[1_501].Y = -0.016

	pointsOfFunctionPlot[1_502].X = 15.02
	pointsOfFunctionPlot[1_502].Y = -0.018

	pointsOfFunctionPlot[1_503].X = 15.03
	pointsOfFunctionPlot[1_503].Y = -0.02

	pointsOfFunctionPlot[1_504].X = 15.04
	pointsOfFunctionPlot[1_504].Y = -0.022

	pointsOfFunctionPlot[1_505].X = 15.05
	pointsOfFunctionPlot[1_505].Y = -0.024

	pointsOfFunctionPlot[1_506].X = 15.06
	pointsOfFunctionPlot[1_506].Y = -0.026

	pointsOfFunctionPlot[1_507].X = 15.07
	pointsOfFunctionPlot[1_507].Y = -0.028

	pointsOfFunctionPlot[1_508].X = 15.08
	pointsOfFunctionPlot[1_508].Y = -0.03

	pointsOfFunctionPlot[1_509].X = 15.09
	pointsOfFunctionPlot[1_509].Y = -0.032

	pointsOfFunctionPlot[1_510].X = 15.10
	pointsOfFunctionPlot[1_510].Y = -0.034

	pointsOfFunctionPlot[1_511].X = 15.11
	pointsOfFunctionPlot[1_511].Y = -0.036

	pointsOfFunctionPlot[1_512].X = 15.12
	pointsOfFunctionPlot[1_512].Y = -0.038

	pointsOfFunctionPlot[1_513].X = 15.13
	pointsOfFunctionPlot[1_513].Y = -0.04

	pointsOfFunctionPlot[1_514].X = 15.14
	pointsOfFunctionPlot[1_514].Y = -0.042

	pointsOfFunctionPlot[1_515].X = 15.15
	pointsOfFunctionPlot[1_515].Y = -0.044

	pointsOfFunctionPlot[1_516].X = 15.16
	pointsOfFunctionPlot[1_516].Y = -0.046

	pointsOfFunctionPlot[1_517].X = 15.17
	pointsOfFunctionPlot[1_517].Y = -0.048

	pointsOfFunctionPlot[1_518].X = 15.18
	pointsOfFunctionPlot[1_518].Y = -0.05

	pointsOfFunctionPlot[1_519].X = 15.19
	pointsOfFunctionPlot[1_519].Y = -0.052

	pointsOfFunctionPlot[1_520].X = 15.20
	pointsOfFunctionPlot[1_520].Y = -0.054

	pointsOfFunctionPlot[1_521].X = 15.21
	pointsOfFunctionPlot[1_521].Y = -0.056

	pointsOfFunctionPlot[1_522].X = 15.22
	pointsOfFunctionPlot[1_522].Y = -0.058

	pointsOfFunctionPlot[1_523].X = 15.23
	pointsOfFunctionPlot[1_523].Y = -0.06

	pointsOfFunctionPlot[1_524].X = 15.24
	pointsOfFunctionPlot[1_524].Y = -0.062

	pointsOfFunctionPlot[1_525].X = 15.25
	pointsOfFunctionPlot[1_525].Y = -0.064

	pointsOfFunctionPlot[1_526].X = 15.26
	pointsOfFunctionPlot[1_526].Y = -0.066

	pointsOfFunctionPlot[1_527].X = 15.27
	pointsOfFunctionPlot[1_527].Y = -0.067

	pointsOfFunctionPlot[1_528].X = 15.28
	pointsOfFunctionPlot[1_528].Y = -0.069

	pointsOfFunctionPlot[1_529].X = 15.29
	pointsOfFunctionPlot[1_529].Y = -0.071

	pointsOfFunctionPlot[1_530].X = 15.30
	pointsOfFunctionPlot[1_530].Y = -0.073

	pointsOfFunctionPlot[1_531].X = 15.31
	pointsOfFunctionPlot[1_531].Y = -0.075

	pointsOfFunctionPlot[1_532].X = 15.32
	pointsOfFunctionPlot[1_532].Y = -0.077

	pointsOfFunctionPlot[1_533].X = 15.33
	pointsOfFunctionPlot[1_533].Y = -0.079

	pointsOfFunctionPlot[1_534].X = 15.34
	pointsOfFunctionPlot[1_534].Y = -0.081

	pointsOfFunctionPlot[1_535].X = 15.35
	pointsOfFunctionPlot[1_535].Y = -0.082

	pointsOfFunctionPlot[1_536].X = 15.36
	pointsOfFunctionPlot[1_536].Y = -0.084

	pointsOfFunctionPlot[1_537].X = 15.37
	pointsOfFunctionPlot[1_537].Y = -0.086

	pointsOfFunctionPlot[1_538].X = 15.38
	pointsOfFunctionPlot[1_538].Y = -0.088

	pointsOfFunctionPlot[1_539].X = 15.39
	pointsOfFunctionPlot[1_539].Y = -0.09

	pointsOfFunctionPlot[1_540].X = 15.40
	pointsOfFunctionPlot[1_540].Y = -0.091

	pointsOfFunctionPlot[1_541].X = 15.41
	pointsOfFunctionPlot[1_541].Y = -0.093

	pointsOfFunctionPlot[1_542].X = 15.42
	pointsOfFunctionPlot[1_542].Y = -0.095

	pointsOfFunctionPlot[1_543].X = 15.43
	pointsOfFunctionPlot[1_543].Y = -0.097

	pointsOfFunctionPlot[1_544].X = 15.44
	pointsOfFunctionPlot[1_544].Y = -0.098

	pointsOfFunctionPlot[1_545].X = 15.45
	pointsOfFunctionPlot[1_545].Y = -0.1

	pointsOfFunctionPlot[1_546].X = 15.46
	pointsOfFunctionPlot[1_546].Y = -0.102

	pointsOfFunctionPlot[1_547].X = 15.47
	pointsOfFunctionPlot[1_547].Y = -0.104

	pointsOfFunctionPlot[1_548].X = 15.48
	pointsOfFunctionPlot[1_548].Y = -0.105

	pointsOfFunctionPlot[1_549].X = 15.49
	pointsOfFunctionPlot[1_549].Y = -0.107

	pointsOfFunctionPlot[1_550].X = 15.50
	pointsOfFunctionPlot[1_550].Y = -0.109

	pointsOfFunctionPlot[1_551].X = 15.51
	pointsOfFunctionPlot[1_551].Y = -0.11

	pointsOfFunctionPlot[1_552].X = 15.52
	pointsOfFunctionPlot[1_552].Y = -0.112

	pointsOfFunctionPlot[1_553].X = 15.53
	pointsOfFunctionPlot[1_553].Y = -0.114

	pointsOfFunctionPlot[1_554].X = 15.54
	pointsOfFunctionPlot[1_554].Y = -0.115

	pointsOfFunctionPlot[1_555].X = 15.55
	pointsOfFunctionPlot[1_555].Y = -0.117

	pointsOfFunctionPlot[1_556].X = 15.56
	pointsOfFunctionPlot[1_556].Y = -0.119

	pointsOfFunctionPlot[1_557].X = 15.57
	pointsOfFunctionPlot[1_557].Y = -0.12

	pointsOfFunctionPlot[1_558].X = 15.58
	pointsOfFunctionPlot[1_558].Y = -0.122

	pointsOfFunctionPlot[1_559].X = 15.59
	pointsOfFunctionPlot[1_559].Y = -0.123

	pointsOfFunctionPlot[1_560].X = 15.60
	pointsOfFunctionPlot[1_560].Y = -0.125

	pointsOfFunctionPlot[1_561].X = 15.61
	pointsOfFunctionPlot[1_561].Y = -0.126

	pointsOfFunctionPlot[1_562].X = 15.62
	pointsOfFunctionPlot[1_562].Y = -0.128

	pointsOfFunctionPlot[1_563].X = 15.63
	pointsOfFunctionPlot[1_563].Y = -0.129

	pointsOfFunctionPlot[1_564].X = 15.64
	pointsOfFunctionPlot[1_564].Y = -0.131

	pointsOfFunctionPlot[1_565].X = 15.65
	pointsOfFunctionPlot[1_565].Y = -0.132

	pointsOfFunctionPlot[1_566].X = 15.66
	pointsOfFunctionPlot[1_566].Y = -0.134

	pointsOfFunctionPlot[1_567].X = 15.67
	pointsOfFunctionPlot[1_567].Y = -0.135

	pointsOfFunctionPlot[1_568].X = 15.68
	pointsOfFunctionPlot[1_568].Y = -0.137

	pointsOfFunctionPlot[1_569].X = 15.69
	pointsOfFunctionPlot[1_569].Y = -0.138

	pointsOfFunctionPlot[1_570].X = 15.70
	pointsOfFunctionPlot[1_570].Y = -0.14

	pointsOfFunctionPlot[1_571].X = 15.71
	pointsOfFunctionPlot[1_571].Y = -0.141

	pointsOfFunctionPlot[1_572].X = 15.72
	pointsOfFunctionPlot[1_572].Y = -0.142

	pointsOfFunctionPlot[1_573].X = 15.73
	pointsOfFunctionPlot[1_573].Y = -0.144

	pointsOfFunctionPlot[1_574].X = 15.74
	pointsOfFunctionPlot[1_574].Y = -0.145

	pointsOfFunctionPlot[1_575].X = 15.75
	pointsOfFunctionPlot[1_575].Y = -0.146

	pointsOfFunctionPlot[1_576].X = 15.76
	pointsOfFunctionPlot[1_576].Y = -0.148

	pointsOfFunctionPlot[1_577].X = 15.77
	pointsOfFunctionPlot[1_577].Y = -0.149

	pointsOfFunctionPlot[1_578].X = 15.78
	pointsOfFunctionPlot[1_578].Y = -0.15

	pointsOfFunctionPlot[1_579].X = 15.79
	pointsOfFunctionPlot[1_579].Y = -0.152

	pointsOfFunctionPlot[1_580].X = 15.80
	pointsOfFunctionPlot[1_580].Y = -0.153

	pointsOfFunctionPlot[1_581].X = 15.81
	pointsOfFunctionPlot[1_581].Y = -0.154

	pointsOfFunctionPlot[1_582].X = 15.82
	pointsOfFunctionPlot[1_582].Y = -0.155

	pointsOfFunctionPlot[1_583].X = 15.83
	pointsOfFunctionPlot[1_583].Y = -0.156

	pointsOfFunctionPlot[1_584].X = 15.84
	pointsOfFunctionPlot[1_584].Y = -0.158

	pointsOfFunctionPlot[1_585].X = 15.85
	pointsOfFunctionPlot[1_585].Y = -0.159

	pointsOfFunctionPlot[1_586].X = 15.86
	pointsOfFunctionPlot[1_586].Y = -0.16

	pointsOfFunctionPlot[1_587].X = 15.87
	pointsOfFunctionPlot[1_587].Y = -0.161

	pointsOfFunctionPlot[1_588].X = 15.88
	pointsOfFunctionPlot[1_588].Y = -0.162

	pointsOfFunctionPlot[1_589].X = 15.89
	pointsOfFunctionPlot[1_589].Y = -0.163

	pointsOfFunctionPlot[1_590].X = 15.90
	pointsOfFunctionPlot[1_590].Y = -0.164

	pointsOfFunctionPlot[1_591].X = 15.91
	pointsOfFunctionPlot[1_591].Y = -0.166

	pointsOfFunctionPlot[1_592].X = 15.92
	pointsOfFunctionPlot[1_592].Y = -0.167

	pointsOfFunctionPlot[1_593].X = 15.93
	pointsOfFunctionPlot[1_593].Y = -0.168

	pointsOfFunctionPlot[1_594].X = 15.94
	pointsOfFunctionPlot[1_594].Y = -0.169

	pointsOfFunctionPlot[1_595].X = 15.95
	pointsOfFunctionPlot[1_595].Y = -0.17

	pointsOfFunctionPlot[1_596].X = 15.96
	pointsOfFunctionPlot[1_596].Y = -0.171

	pointsOfFunctionPlot[1_597].X = 15.97
	pointsOfFunctionPlot[1_597].Y = -0.172

	pointsOfFunctionPlot[1_598].X = 15.98
	pointsOfFunctionPlot[1_598].Y = -0.173

	pointsOfFunctionPlot[1_599].X = 15.99
	pointsOfFunctionPlot[1_599].Y = -0.173

	pointsOfFunctionPlot[1_600].X = 16.0
	pointsOfFunctionPlot[1_600].Y = -0.174

	pointsOfFunctionPlot[1_601].X = 16.01
	pointsOfFunctionPlot[1_601].Y = -0.175

	pointsOfFunctionPlot[1_602].X = 16.02
	pointsOfFunctionPlot[1_602].Y = -0.176

	pointsOfFunctionPlot[1_603].X = 16.03
	pointsOfFunctionPlot[1_603].Y = -0.177

	pointsOfFunctionPlot[1_604].X = 16.04
	pointsOfFunctionPlot[1_604].Y = -0.178

	pointsOfFunctionPlot[1_605].X = 16.05
	pointsOfFunctionPlot[1_605].Y = -0.179

	pointsOfFunctionPlot[1_606].X = 16.06
	pointsOfFunctionPlot[1_606].Y = -0.179

	pointsOfFunctionPlot[1_607].X = 16.07
	pointsOfFunctionPlot[1_607].Y = -0.18

	pointsOfFunctionPlot[1_608].X = 16.08
	pointsOfFunctionPlot[1_608].Y = -0.181

	pointsOfFunctionPlot[1_609].X = 16.09
	pointsOfFunctionPlot[1_609].Y = -0.182

	pointsOfFunctionPlot[1_610].X = 16.10
	pointsOfFunctionPlot[1_610].Y = -0.183

	pointsOfFunctionPlot[1_611].X = 16.11
	pointsOfFunctionPlot[1_611].Y = -0.183

	pointsOfFunctionPlot[1_612].X = 16.12
	pointsOfFunctionPlot[1_612].Y = -0.184

	pointsOfFunctionPlot[1_613].X = 16.13
	pointsOfFunctionPlot[1_613].Y = -0.185

	pointsOfFunctionPlot[1_614].X = 16.14
	pointsOfFunctionPlot[1_614].Y = -0.185

	pointsOfFunctionPlot[1_615].X = 16.15
	pointsOfFunctionPlot[1_615].Y = -0.186

	pointsOfFunctionPlot[1_616].X = 16.16
	pointsOfFunctionPlot[1_616].Y = -0.187

	pointsOfFunctionPlot[1_617].X = 16.17
	pointsOfFunctionPlot[1_617].Y = -0.187

	pointsOfFunctionPlot[1_618].X = 16.18
	pointsOfFunctionPlot[1_618].Y = -0.188

	pointsOfFunctionPlot[1_619].X = 16.19
	pointsOfFunctionPlot[1_619].Y = -0.188

	pointsOfFunctionPlot[1_620].X = 16.20
	pointsOfFunctionPlot[1_620].Y = -0.189

	pointsOfFunctionPlot[1_621].X = 16.21
	pointsOfFunctionPlot[1_621].Y = -0.189

	pointsOfFunctionPlot[1_622].X = 16.22
	pointsOfFunctionPlot[1_622].Y = -0.19

	pointsOfFunctionPlot[1_623].X = 16.23
	pointsOfFunctionPlot[1_623].Y = -0.19

	pointsOfFunctionPlot[1_624].X = 16.24
	pointsOfFunctionPlot[1_624].Y = -0.191

	pointsOfFunctionPlot[1_625].X = 16.25
	pointsOfFunctionPlot[1_625].Y = -0.191

	pointsOfFunctionPlot[1_626].X = 16.26
	pointsOfFunctionPlot[1_626].Y = -0.192

	pointsOfFunctionPlot[1_627].X = 16.27
	pointsOfFunctionPlot[1_627].Y = -0.192

	pointsOfFunctionPlot[1_628].X = 16.28
	pointsOfFunctionPlot[1_628].Y = -0.192

	pointsOfFunctionPlot[1_629].X = 16.29
	pointsOfFunctionPlot[1_629].Y = -0.193

	pointsOfFunctionPlot[1_630].X = 16.30
	pointsOfFunctionPlot[1_630].Y = -0.193

	pointsOfFunctionPlot[1_631].X = 16.31
	pointsOfFunctionPlot[1_631].Y = -0.193

	pointsOfFunctionPlot[1_632].X = 16.32
	pointsOfFunctionPlot[1_632].Y = -0.194

	pointsOfFunctionPlot[1_633].X = 16.33
	pointsOfFunctionPlot[1_633].Y = -0.194

	pointsOfFunctionPlot[1_634].X = 16.34
	pointsOfFunctionPlot[1_634].Y = -0.194

	pointsOfFunctionPlot[1_635].X = 16.35
	pointsOfFunctionPlot[1_635].Y = -0.195

	pointsOfFunctionPlot[1_636].X = 16.36
	pointsOfFunctionPlot[1_636].Y = -0.195

	pointsOfFunctionPlot[1_637].X = 16.37
	pointsOfFunctionPlot[1_637].Y = -0.195

	pointsOfFunctionPlot[1_638].X = 16.38
	pointsOfFunctionPlot[1_638].Y = -0.195

	pointsOfFunctionPlot[1_639].X = 16.39
	pointsOfFunctionPlot[1_639].Y = -0.195

	pointsOfFunctionPlot[1_640].X = 16.40
	pointsOfFunctionPlot[1_640].Y = -0.195

	pointsOfFunctionPlot[1_641].X = 16.41
	pointsOfFunctionPlot[1_641].Y = -0.196

	pointsOfFunctionPlot[1_642].X = 16.42
	pointsOfFunctionPlot[1_642].Y = -0.196

	pointsOfFunctionPlot[1_643].X = 16.43
	pointsOfFunctionPlot[1_643].Y = -0.196

	pointsOfFunctionPlot[1_644].X = 16.44
	pointsOfFunctionPlot[1_644].Y = -0.196

	pointsOfFunctionPlot[1_645].X = 16.45
	pointsOfFunctionPlot[1_645].Y = -0.196

	pointsOfFunctionPlot[1_646].X = 16.46
	pointsOfFunctionPlot[1_646].Y = -0.196

	pointsOfFunctionPlot[1_647].X = 16.47
	pointsOfFunctionPlot[1_647].Y = -0.196

	pointsOfFunctionPlot[1_648].X = 16.48
	pointsOfFunctionPlot[1_648].Y = -0.196

	pointsOfFunctionPlot[1_649].X = 16.49
	pointsOfFunctionPlot[1_649].Y = -0.196

	pointsOfFunctionPlot[1_650].X = 16.50
	pointsOfFunctionPlot[1_650].Y = -0.196

	pointsOfFunctionPlot[1_651].X = 16.51
	pointsOfFunctionPlot[1_651].Y = -0.196

	pointsOfFunctionPlot[1_652].X = 16.52
	pointsOfFunctionPlot[1_652].Y = -0.196

	pointsOfFunctionPlot[1_653].X = 16.53
	pointsOfFunctionPlot[1_653].Y = -0.196

	pointsOfFunctionPlot[1_654].X = 16.54
	pointsOfFunctionPlot[1_654].Y = -0.195

	pointsOfFunctionPlot[1_655].X = 16.55
	pointsOfFunctionPlot[1_655].Y = -0.195

	pointsOfFunctionPlot[1_656].X = 16.56
	pointsOfFunctionPlot[1_656].Y = -0.195

	pointsOfFunctionPlot[1_657].X = 16.57
	pointsOfFunctionPlot[1_657].Y = -0.195

	pointsOfFunctionPlot[1_658].X = 16.58
	pointsOfFunctionPlot[1_658].Y = -0.195

	pointsOfFunctionPlot[1_659].X = 16.59
	pointsOfFunctionPlot[1_659].Y = -0.195

	pointsOfFunctionPlot[1_660].X = 16.60
	pointsOfFunctionPlot[1_660].Y = -0.195

	pointsOfFunctionPlot[1_661].X = 16.61
	pointsOfFunctionPlot[1_661].Y = -0.194

	pointsOfFunctionPlot[1_662].X = 16.62
	pointsOfFunctionPlot[1_662].Y = -0.194

	pointsOfFunctionPlot[1_663].X = 16.63
	pointsOfFunctionPlot[1_663].Y = -0.193

	pointsOfFunctionPlot[1_664].X = 16.64
	pointsOfFunctionPlot[1_664].Y = -0.193

	pointsOfFunctionPlot[1_665].X = 16.65
	pointsOfFunctionPlot[1_665].Y = -0.193

	pointsOfFunctionPlot[1_666].X = 16.66
	pointsOfFunctionPlot[1_666].Y = -0.192

	pointsOfFunctionPlot[1_667].X = 16.67
	pointsOfFunctionPlot[1_667].Y = -0.192

	pointsOfFunctionPlot[1_668].X = 16.68
	pointsOfFunctionPlot[1_668].Y = -0.192

	pointsOfFunctionPlot[1_669].X = 16.69
	pointsOfFunctionPlot[1_669].Y = -0.191

	pointsOfFunctionPlot[1_670].X = 16.70
	pointsOfFunctionPlot[1_670].Y = -0.191

	pointsOfFunctionPlot[1_671].X = 16.71
	pointsOfFunctionPlot[1_671].Y = -0.19

	pointsOfFunctionPlot[1_672].X = 16.72
	pointsOfFunctionPlot[1_672].Y = -0.19

	pointsOfFunctionPlot[1_673].X = 16.73
	pointsOfFunctionPlot[1_673].Y = -0.189

	pointsOfFunctionPlot[1_674].X = 16.74
	pointsOfFunctionPlot[1_674].Y = -0.189

	pointsOfFunctionPlot[1_675].X = 16.75
	pointsOfFunctionPlot[1_675].Y = -0.188

	pointsOfFunctionPlot[1_676].X = 16.76
	pointsOfFunctionPlot[1_676].Y = -0.188

	pointsOfFunctionPlot[1_677].X = 16.77
	pointsOfFunctionPlot[1_677].Y = -0.187

	pointsOfFunctionPlot[1_678].X = 16.78
	pointsOfFunctionPlot[1_678].Y = -0.187

	pointsOfFunctionPlot[1_679].X = 16.79
	pointsOfFunctionPlot[1_679].Y = -0.186

	pointsOfFunctionPlot[1_680].X = 16.80
	pointsOfFunctionPlot[1_680].Y = -0.185

	pointsOfFunctionPlot[1_681].X = 16.81
	pointsOfFunctionPlot[1_681].Y = -0.185

	pointsOfFunctionPlot[1_682].X = 16.82
	pointsOfFunctionPlot[1_682].Y = -0.184

	pointsOfFunctionPlot[1_683].X = 16.83
	pointsOfFunctionPlot[1_683].Y = -0.184

	pointsOfFunctionPlot[1_684].X = 16.84
	pointsOfFunctionPlot[1_684].Y = -0.183

	pointsOfFunctionPlot[1_685].X = 16.85
	pointsOfFunctionPlot[1_685].Y = -0.182

	pointsOfFunctionPlot[1_686].X = 16.86
	pointsOfFunctionPlot[1_686].Y = -0.181

	pointsOfFunctionPlot[1_687].X = 16.87
	pointsOfFunctionPlot[1_687].Y = -0.181

	pointsOfFunctionPlot[1_688].X = 16.88
	pointsOfFunctionPlot[1_688].Y = -0.18

	pointsOfFunctionPlot[1_689].X = 16.89
	pointsOfFunctionPlot[1_689].Y = -0.179

	pointsOfFunctionPlot[1_690].X = 16.90
	pointsOfFunctionPlot[1_690].Y = -0.178

	pointsOfFunctionPlot[1_691].X = 16.91
	pointsOfFunctionPlot[1_691].Y = -0.177

	pointsOfFunctionPlot[1_692].X = 16.92
	pointsOfFunctionPlot[1_692].Y = -0.177

	pointsOfFunctionPlot[1_693].X = 16.93
	pointsOfFunctionPlot[1_693].Y = -0.176

	pointsOfFunctionPlot[1_694].X = 16.94
	pointsOfFunctionPlot[1_694].Y = -0.175

	pointsOfFunctionPlot[1_695].X = 16.95
	pointsOfFunctionPlot[1_695].Y = -0.174

	pointsOfFunctionPlot[1_696].X = 16.96
	pointsOfFunctionPlot[1_696].Y = -0.173

	pointsOfFunctionPlot[1_697].X = 16.97
	pointsOfFunctionPlot[1_697].Y = -0.172

	pointsOfFunctionPlot[1_698].X = 16.98
	pointsOfFunctionPlot[1_698].Y = -0.171

	pointsOfFunctionPlot[1_699].X = 16.99
	pointsOfFunctionPlot[1_699].Y = -0.17

	pointsOfFunctionPlot[1_700].X = 17.0
	pointsOfFunctionPlot[1_700].Y = -0.169

	pointsOfFunctionPlot[1_701].X = 17.01
	pointsOfFunctionPlot[1_701].Y = -0.169

	pointsOfFunctionPlot[1_702].X = 17.02
	pointsOfFunctionPlot[1_702].Y = -0.167

	pointsOfFunctionPlot[1_703].X = 17.03
	pointsOfFunctionPlot[1_703].Y = -0.166

	pointsOfFunctionPlot[1_704].X = 17.04
	pointsOfFunctionPlot[1_704].Y = -0.165

	pointsOfFunctionPlot[1_705].X = 17.05
	pointsOfFunctionPlot[1_705].Y = -0.164

	pointsOfFunctionPlot[1_706].X = 17.06
	pointsOfFunctionPlot[1_706].Y = -0.163

	pointsOfFunctionPlot[1_707].X = 17.07
	pointsOfFunctionPlot[1_707].Y = -0.162

	pointsOfFunctionPlot[1_708].X = 17.08
	pointsOfFunctionPlot[1_708].Y = -0.161

	pointsOfFunctionPlot[1_709].X = 17.09
	pointsOfFunctionPlot[1_709].Y = -0.16

	pointsOfFunctionPlot[1_710].X = 17.10
	pointsOfFunctionPlot[1_710].Y = -0.159

	pointsOfFunctionPlot[1_711].X = 17.11
	pointsOfFunctionPlot[1_711].Y = -0.158

	pointsOfFunctionPlot[1_712].X = 17.12
	pointsOfFunctionPlot[1_712].Y = -0.156

	pointsOfFunctionPlot[1_713].X = 17.13
	pointsOfFunctionPlot[1_713].Y = -0.155

	pointsOfFunctionPlot[1_714].X = 17.14
	pointsOfFunctionPlot[1_714].Y = -0.154

	pointsOfFunctionPlot[1_715].X = 17.15
	pointsOfFunctionPlot[1_715].Y = -0.153

	pointsOfFunctionPlot[1_716].X = 17.16
	pointsOfFunctionPlot[1_716].Y = -0.152

	pointsOfFunctionPlot[1_717].X = 17.17
	pointsOfFunctionPlot[1_717].Y = -0.15

	pointsOfFunctionPlot[1_718].X = 17.18
	pointsOfFunctionPlot[1_718].Y = -0.149

	pointsOfFunctionPlot[1_719].X = 17.19
	pointsOfFunctionPlot[1_719].Y = -0.148

	pointsOfFunctionPlot[1_720].X = 17.20
	pointsOfFunctionPlot[1_720].Y = -0.147

	pointsOfFunctionPlot[1_721].X = 17.21
	pointsOfFunctionPlot[1_721].Y = -0.145

	pointsOfFunctionPlot[1_722].X = 17.22
	pointsOfFunctionPlot[1_722].Y = -0.144

	pointsOfFunctionPlot[1_723].X = 17.23
	pointsOfFunctionPlot[1_723].Y = -0.143

	pointsOfFunctionPlot[1_724].X = 17.24
	pointsOfFunctionPlot[1_724].Y = -0.141

	pointsOfFunctionPlot[1_725].X = 17.25
	pointsOfFunctionPlot[1_725].Y = -0.14

	pointsOfFunctionPlot[1_726].X = 17.26
	pointsOfFunctionPlot[1_726].Y = -0.139

	pointsOfFunctionPlot[1_727].X = 17.27
	pointsOfFunctionPlot[1_727].Y = -0.137

	pointsOfFunctionPlot[1_728].X = 17.28
	pointsOfFunctionPlot[1_728].Y = -0.136

	pointsOfFunctionPlot[1_729].X = 17.29
	pointsOfFunctionPlot[1_729].Y = -0.135

	pointsOfFunctionPlot[1_730].X = 17.30
	pointsOfFunctionPlot[1_730].Y = -0.133

	pointsOfFunctionPlot[1_731].X = 17.31
	pointsOfFunctionPlot[1_731].Y = -0.132

	pointsOfFunctionPlot[1_732].X = 17.32
	pointsOfFunctionPlot[1_732].Y = -0.13

	pointsOfFunctionPlot[1_733].X = 17.33
	pointsOfFunctionPlot[1_733].Y = -0.129

	pointsOfFunctionPlot[1_734].X = 17.34
	pointsOfFunctionPlot[1_734].Y = -0.127

	pointsOfFunctionPlot[1_735].X = 17.35
	pointsOfFunctionPlot[1_735].Y = -0.126

	pointsOfFunctionPlot[1_736].X = 17.36
	pointsOfFunctionPlot[1_736].Y = -0.124

	pointsOfFunctionPlot[1_737].X = 17.37
	pointsOfFunctionPlot[1_737].Y = -0.123

	pointsOfFunctionPlot[1_738].X = 17.38
	pointsOfFunctionPlot[1_738].Y = -0.121

	pointsOfFunctionPlot[1_739].X = 17.39
	pointsOfFunctionPlot[1_739].Y = -0.12

	pointsOfFunctionPlot[1_740].X = 17.40
	pointsOfFunctionPlot[1_740].Y = -0.118

	pointsOfFunctionPlot[1_741].X = 17.41
	pointsOfFunctionPlot[1_741].Y = -0.117

	pointsOfFunctionPlot[1_742].X = 17.42
	pointsOfFunctionPlot[1_742].Y = -0.115

	pointsOfFunctionPlot[1_743].X = 17.43
	pointsOfFunctionPlot[1_743].Y = -0.114

	pointsOfFunctionPlot[1_744].X = 17.44
	pointsOfFunctionPlot[1_744].Y = -0.112

	pointsOfFunctionPlot[1_745].X = 17.45
	pointsOfFunctionPlot[1_745].Y = -0.111

	pointsOfFunctionPlot[1_746].X = 17.46
	pointsOfFunctionPlot[1_746].Y = -0.109

	pointsOfFunctionPlot[1_747].X = 17.47
	pointsOfFunctionPlot[1_747].Y = -0.107

	pointsOfFunctionPlot[1_748].X = 17.48
	pointsOfFunctionPlot[1_748].Y = -0.106

	pointsOfFunctionPlot[1_749].X = 17.49
	pointsOfFunctionPlot[1_749].Y = -0.104

	pointsOfFunctionPlot[1_750].X = 17.50
	pointsOfFunctionPlot[1_750].Y = -0.103

	pointsOfFunctionPlot[1_751].X = 17.51
	pointsOfFunctionPlot[1_751].Y = -0.101

	pointsOfFunctionPlot[1_752].X = 17.52
	pointsOfFunctionPlot[1_752].Y = -0.099

	pointsOfFunctionPlot[1_753].X = 17.53
	pointsOfFunctionPlot[1_753].Y = -0.098

	pointsOfFunctionPlot[1_754].X = 17.54
	pointsOfFunctionPlot[1_754].Y = -0.096

	pointsOfFunctionPlot[1_755].X = 17.55
	pointsOfFunctionPlot[1_755].Y = -0.094

	pointsOfFunctionPlot[1_756].X = 17.56
	pointsOfFunctionPlot[1_756].Y = -0.093

	pointsOfFunctionPlot[1_757].X = 17.57
	pointsOfFunctionPlot[1_757].Y = -0.091

	pointsOfFunctionPlot[1_758].X = 17.58
	pointsOfFunctionPlot[1_758].Y = -0.089

	pointsOfFunctionPlot[1_759].X = 17.59
	pointsOfFunctionPlot[1_759].Y = -0.088

	pointsOfFunctionPlot[1_760].X = 17.60
	pointsOfFunctionPlot[1_760].Y = -0.086

	pointsOfFunctionPlot[1_761].X = 17.61
	pointsOfFunctionPlot[1_761].Y = -0.084

	pointsOfFunctionPlot[1_762].X = 17.62
	pointsOfFunctionPlot[1_762].Y = -0.082

	pointsOfFunctionPlot[1_763].X = 17.63
	pointsOfFunctionPlot[1_763].Y = -0.081

	pointsOfFunctionPlot[1_764].X = 17.64
	pointsOfFunctionPlot[1_764].Y = -0.079

	pointsOfFunctionPlot[1_765].X = 17.65
	pointsOfFunctionPlot[1_765].Y = -0.077

	pointsOfFunctionPlot[1_766].X = 17.66
	pointsOfFunctionPlot[1_766].Y = -0.075

	pointsOfFunctionPlot[1_767].X = 17.67
	pointsOfFunctionPlot[1_767].Y = -0.074

	pointsOfFunctionPlot[1_768].X = 17.68
	pointsOfFunctionPlot[1_768].Y = -0.072

	pointsOfFunctionPlot[1_769].X = 17.69
	pointsOfFunctionPlot[1_769].Y = -0.07

	pointsOfFunctionPlot[1_770].X = 17.70
	pointsOfFunctionPlot[1_770].Y = -0.068

	pointsOfFunctionPlot[1_771].X = 17.71
	pointsOfFunctionPlot[1_771].Y = -0.066

	pointsOfFunctionPlot[1_772].X = 17.72
	pointsOfFunctionPlot[1_772].Y = -0.065

	pointsOfFunctionPlot[1_773].X = 17.73
	pointsOfFunctionPlot[1_773].Y = -0.063

	pointsOfFunctionPlot[1_774].X = 17.74
	pointsOfFunctionPlot[1_774].Y = -0.061

	pointsOfFunctionPlot[1_775].X = 17.75
	pointsOfFunctionPlot[1_775].Y = -0.059

	pointsOfFunctionPlot[1_776].X = 17.76
	pointsOfFunctionPlot[1_776].Y = -0.057

	pointsOfFunctionPlot[1_777].X = 17.77
	pointsOfFunctionPlot[1_777].Y = -0.056

	pointsOfFunctionPlot[1_778].X = 17.78
	pointsOfFunctionPlot[1_778].Y = -0.054

	pointsOfFunctionPlot[1_779].X = 17.79
	pointsOfFunctionPlot[1_779].Y = -0.052

	pointsOfFunctionPlot[1_780].X = 17.80
	pointsOfFunctionPlot[1_780].Y = -0.05

	pointsOfFunctionPlot[1_781].X = 17.81
	pointsOfFunctionPlot[1_781].Y = -0.048

	pointsOfFunctionPlot[1_782].X = 17.82
	pointsOfFunctionPlot[1_782].Y = -0.046

	pointsOfFunctionPlot[1_783].X = 17.83
	pointsOfFunctionPlot[1_783].Y = -0.045

	pointsOfFunctionPlot[1_784].X = 17.84
	pointsOfFunctionPlot[1_784].Y = -0.043

	pointsOfFunctionPlot[1_785].X = 17.85
	pointsOfFunctionPlot[1_785].Y = -0.041

	pointsOfFunctionPlot[1_786].X = 17.86
	pointsOfFunctionPlot[1_786].Y = -0.039

	pointsOfFunctionPlot[1_787].X = 17.87
	pointsOfFunctionPlot[1_787].Y = -0.037

	pointsOfFunctionPlot[1_788].X = 17.88
	pointsOfFunctionPlot[1_788].Y = -0.035

	pointsOfFunctionPlot[1_789].X = 17.89
	pointsOfFunctionPlot[1_789].Y = -0.033

	pointsOfFunctionPlot[1_790].X = 17.90
	pointsOfFunctionPlot[1_790].Y = -0.032

	pointsOfFunctionPlot[1_791].X = 17.91
	pointsOfFunctionPlot[1_791].Y = -0.03

	pointsOfFunctionPlot[1_792].X = 17.92
	pointsOfFunctionPlot[1_792].Y = -0.028

	pointsOfFunctionPlot[1_793].X = 17.93
	pointsOfFunctionPlot[1_793].Y = -0.026

	pointsOfFunctionPlot[1_794].X = 17.94
	pointsOfFunctionPlot[1_794].Y = -0.024

	pointsOfFunctionPlot[1_795].X = 17.95
	pointsOfFunctionPlot[1_795].Y = -0.022

	pointsOfFunctionPlot[1_796].X = 17.96
	pointsOfFunctionPlot[1_796].Y = -0.02

	pointsOfFunctionPlot[1_797].X = 17.97
	pointsOfFunctionPlot[1_797].Y = -0.018

	pointsOfFunctionPlot[1_798].X = 17.98
	pointsOfFunctionPlot[1_798].Y = -0.017

	pointsOfFunctionPlot[1_799].X = 17.99
	pointsOfFunctionPlot[1_799].Y = -0.015

	pointsOfFunctionPlot[1_800].X = 18.0
	pointsOfFunctionPlot[1_800].Y = -0.013

	pointsOfFunctionPlot[1_801].X = 18.01
	pointsOfFunctionPlot[1_801].Y = -0.011

	pointsOfFunctionPlot[1_802].X = 18.02
	pointsOfFunctionPlot[1_802].Y = -0.009

	pointsOfFunctionPlot[1_803].X = 18.03
	pointsOfFunctionPlot[1_803].Y = -0.007

	pointsOfFunctionPlot[1_804].X = 18.04
	pointsOfFunctionPlot[1_804].Y = -0.005

	pointsOfFunctionPlot[1_805].X = 18.05
	pointsOfFunctionPlot[1_805].Y = -0.003

	pointsOfFunctionPlot[1_806].X = 18.06
	pointsOfFunctionPlot[1_806].Y = -0.002

	pointsOfFunctionPlot[1_807].X = 18.07
	pointsOfFunctionPlot[1_807].Y = 0.0

	pointsOfFunctionPlot[1_808].X = 18.08
	pointsOfFunctionPlot[1_808].Y = 0.001

	pointsOfFunctionPlot[1_809].X = 18.09
	pointsOfFunctionPlot[1_809].Y = 0.003

	pointsOfFunctionPlot[1_810].X = 18.10
	pointsOfFunctionPlot[1_810].Y = 0.005

	pointsOfFunctionPlot[1_811].X = 18.11
	pointsOfFunctionPlot[1_811].Y = 0.007

	pointsOfFunctionPlot[1_812].X = 18.12
	pointsOfFunctionPlot[1_812].Y = 0.009

	pointsOfFunctionPlot[1_813].X = 18.13
	pointsOfFunctionPlot[1_813].Y = 0.011

	pointsOfFunctionPlot[1_814].X = 18.14
	pointsOfFunctionPlot[1_814].Y = 0.012

	pointsOfFunctionPlot[1_815].X = 18.15
	pointsOfFunctionPlot[1_815].Y = 0.014

	pointsOfFunctionPlot[1_816].X = 18.16
	pointsOfFunctionPlot[1_816].Y = 0.016

	pointsOfFunctionPlot[1_817].X = 18.17
	pointsOfFunctionPlot[1_817].Y = 0.018

	pointsOfFunctionPlot[1_818].X = 18.18
	pointsOfFunctionPlot[1_818].Y = 0.02

	pointsOfFunctionPlot[1_819].X = 18.19
	pointsOfFunctionPlot[1_819].Y = 0.022

	pointsOfFunctionPlot[1_820].X = 18.20
	pointsOfFunctionPlot[1_820].Y = 0.024

	pointsOfFunctionPlot[1_821].X = 18.21
	pointsOfFunctionPlot[1_821].Y = 0.025

	pointsOfFunctionPlot[1_822].X = 18.22
	pointsOfFunctionPlot[1_822].Y = 0.027

	pointsOfFunctionPlot[1_823].X = 18.23
	pointsOfFunctionPlot[1_823].Y = 0.029

	pointsOfFunctionPlot[1_824].X = 18.24
	pointsOfFunctionPlot[1_824].Y = 0.031

	pointsOfFunctionPlot[1_825].X = 18.25
	pointsOfFunctionPlot[1_825].Y = 0.033

	pointsOfFunctionPlot[1_826].X = 18.26
	pointsOfFunctionPlot[1_826].Y = 0.035

	pointsOfFunctionPlot[1_827].X = 18.27
	pointsOfFunctionPlot[1_827].Y = 0.036

	pointsOfFunctionPlot[1_828].X = 18.28
	pointsOfFunctionPlot[1_828].Y = 0.038

	pointsOfFunctionPlot[1_829].X = 18.29
	pointsOfFunctionPlot[1_829].Y = 0.04

	pointsOfFunctionPlot[1_830].X = 18.30
	pointsOfFunctionPlot[1_830].Y = 0.042

	pointsOfFunctionPlot[1_831].X = 18.31
	pointsOfFunctionPlot[1_831].Y = 0.044

	pointsOfFunctionPlot[1_832].X = 18.32
	pointsOfFunctionPlot[1_832].Y = 0.045

	pointsOfFunctionPlot[1_833].X = 18.33
	pointsOfFunctionPlot[1_833].Y = 0.047

	pointsOfFunctionPlot[1_834].X = 18.34
	pointsOfFunctionPlot[1_834].Y = 0.049

	pointsOfFunctionPlot[1_835].X = 18.35
	pointsOfFunctionPlot[1_835].Y = 0.051

	pointsOfFunctionPlot[1_836].X = 18.36
	pointsOfFunctionPlot[1_836].Y = 0.053

	pointsOfFunctionPlot[1_837].X = 18.37
	pointsOfFunctionPlot[1_837].Y = 0.054

	pointsOfFunctionPlot[1_838].X = 18.38
	pointsOfFunctionPlot[1_838].Y = 0.056

	pointsOfFunctionPlot[1_839].X = 18.39
	pointsOfFunctionPlot[1_839].Y = 0.058

	pointsOfFunctionPlot[1_840].X = 18.40
	pointsOfFunctionPlot[1_840].Y = 0.06

	pointsOfFunctionPlot[1_841].X = 18.41
	pointsOfFunctionPlot[1_841].Y = 0.061

	pointsOfFunctionPlot[1_842].X = 18.42
	pointsOfFunctionPlot[1_842].Y = 0.063

	pointsOfFunctionPlot[1_843].X = 18.43
	pointsOfFunctionPlot[1_843].Y = 0.065

	pointsOfFunctionPlot[1_844].X = 18.44
	pointsOfFunctionPlot[1_844].Y = 0.067

	pointsOfFunctionPlot[1_845].X = 18.45
	pointsOfFunctionPlot[1_845].Y = 0.068

	pointsOfFunctionPlot[1_846].X = 18.46
	pointsOfFunctionPlot[1_846].Y = 0.07

	pointsOfFunctionPlot[1_847].X = 18.47
	pointsOfFunctionPlot[1_847].Y = 0.072

	pointsOfFunctionPlot[1_848].X = 18.48
	pointsOfFunctionPlot[1_848].Y = 0.073

	pointsOfFunctionPlot[1_849].X = 18.49
	pointsOfFunctionPlot[1_849].Y = 0.075

	pointsOfFunctionPlot[1_850].X = 18.50
	pointsOfFunctionPlot[1_850].Y = 0.077

	pointsOfFunctionPlot[1_851].X = 18.51
	pointsOfFunctionPlot[1_851].Y = 0.078

	pointsOfFunctionPlot[1_852].X = 18.52
	pointsOfFunctionPlot[1_852].Y = 0.08

	pointsOfFunctionPlot[1_853].X = 18.53
	pointsOfFunctionPlot[1_853].Y = 0.082

	pointsOfFunctionPlot[1_854].X = 18.54
	pointsOfFunctionPlot[1_854].Y = 0.083

	pointsOfFunctionPlot[1_855].X = 18.55
	pointsOfFunctionPlot[1_855].Y = 0.085

	pointsOfFunctionPlot[1_856].X = 18.56
	pointsOfFunctionPlot[1_856].Y = 0.087

	pointsOfFunctionPlot[1_857].X = 18.57
	pointsOfFunctionPlot[1_857].Y = 0.088

	pointsOfFunctionPlot[1_858].X = 18.58
	pointsOfFunctionPlot[1_858].Y = 0.09

	pointsOfFunctionPlot[1_859].X = 18.59
	pointsOfFunctionPlot[1_859].Y = 0.091

	pointsOfFunctionPlot[1_860].X = 18.60
	pointsOfFunctionPlot[1_860].Y = 0.093

	pointsOfFunctionPlot[1_861].X = 18.61
	pointsOfFunctionPlot[1_861].Y = 0.094

	pointsOfFunctionPlot[1_862].X = 18.62
	pointsOfFunctionPlot[1_862].Y = 0.096

	pointsOfFunctionPlot[1_863].X = 18.63
	pointsOfFunctionPlot[1_863].Y = 0.098

	pointsOfFunctionPlot[1_864].X = 18.64
	pointsOfFunctionPlot[1_864].Y = 0.099

	pointsOfFunctionPlot[1_865].X = 18.65
	pointsOfFunctionPlot[1_865].Y = 0.101

	pointsOfFunctionPlot[1_866].X = 18.66
	pointsOfFunctionPlot[1_866].Y = 0.102

	pointsOfFunctionPlot[1_867].X = 18.67
	pointsOfFunctionPlot[1_867].Y = 0.104

	pointsOfFunctionPlot[1_868].X = 18.68
	pointsOfFunctionPlot[1_868].Y = 0.105

	pointsOfFunctionPlot[1_869].X = 18.69
	pointsOfFunctionPlot[1_869].Y = 0.107

	pointsOfFunctionPlot[1_870].X = 18.70
	pointsOfFunctionPlot[1_870].Y = 0.108

	pointsOfFunctionPlot[1_871].X = 18.71
	pointsOfFunctionPlot[1_871].Y = 0.11

	pointsOfFunctionPlot[1_872].X = 18.72
	pointsOfFunctionPlot[1_872].Y = 0.111

	pointsOfFunctionPlot[1_873].X = 18.73
	pointsOfFunctionPlot[1_873].Y = 0.112

	pointsOfFunctionPlot[1_874].X = 18.74
	pointsOfFunctionPlot[1_874].Y = 0.114

	pointsOfFunctionPlot[1_875].X = 18.75
	pointsOfFunctionPlot[1_875].Y = 0.115

	pointsOfFunctionPlot[1_876].X = 18.76
	pointsOfFunctionPlot[1_876].Y = 0.117

	pointsOfFunctionPlot[1_877].X = 18.77
	pointsOfFunctionPlot[1_877].Y = 0.118

	pointsOfFunctionPlot[1_878].X = 18.78
	pointsOfFunctionPlot[1_878].Y = 0.119

	pointsOfFunctionPlot[1_879].X = 18.79
	pointsOfFunctionPlot[1_879].Y = 0.121

	pointsOfFunctionPlot[1_880].X = 18.80
	pointsOfFunctionPlot[1_880].Y = 0.122

	pointsOfFunctionPlot[1_881].X = 18.81
	pointsOfFunctionPlot[1_881].Y = 0.123

	pointsOfFunctionPlot[1_882].X = 18.82
	pointsOfFunctionPlot[1_882].Y = 0.125

	pointsOfFunctionPlot[1_883].X = 18.83
	pointsOfFunctionPlot[1_883].Y = 0.126

	pointsOfFunctionPlot[1_884].X = 18.84
	pointsOfFunctionPlot[1_884].Y = 0.127

	pointsOfFunctionPlot[1_885].X = 18.85
	pointsOfFunctionPlot[1_885].Y = 0.129

	pointsOfFunctionPlot[1_886].X = 18.86
	pointsOfFunctionPlot[1_886].Y = 0.13

	pointsOfFunctionPlot[1_887].X = 18.87
	pointsOfFunctionPlot[1_887].Y = 0.131

	pointsOfFunctionPlot[1_888].X = 18.88
	pointsOfFunctionPlot[1_888].Y = 0.132

	pointsOfFunctionPlot[1_889].X = 18.89
	pointsOfFunctionPlot[1_889].Y = 0.134

	pointsOfFunctionPlot[1_890].X = 18.90
	pointsOfFunctionPlot[1_890].Y = 0.135

	pointsOfFunctionPlot[1_891].X = 18.91
	pointsOfFunctionPlot[1_891].Y = 0.136

	pointsOfFunctionPlot[1_892].X = 18.92
	pointsOfFunctionPlot[1_892].Y = 0.137

	pointsOfFunctionPlot[1_893].X = 18.93
	pointsOfFunctionPlot[1_893].Y = 0.138

	pointsOfFunctionPlot[1_894].X = 18.94
	pointsOfFunctionPlot[1_894].Y = 0.14

	pointsOfFunctionPlot[1_895].X = 18.95
	pointsOfFunctionPlot[1_895].Y = 0.141

	pointsOfFunctionPlot[1_896].X = 18.96
	pointsOfFunctionPlot[1_896].Y = 0.142

	pointsOfFunctionPlot[1_897].X = 18.97
	pointsOfFunctionPlot[1_897].Y = 0.143

	pointsOfFunctionPlot[1_898].X = 18.98
	pointsOfFunctionPlot[1_898].Y = 0.144

	pointsOfFunctionPlot[1_899].X = 18.99
	pointsOfFunctionPlot[1_899].Y = 0.145

	pointsOfFunctionPlot[1_900].X = 19.0
	pointsOfFunctionPlot[1_900].Y = 0.146

	pointsOfFunctionPlot[1_901].X = 19.01
	pointsOfFunctionPlot[1_901].Y = 0.147

	pointsOfFunctionPlot[1_902].X = 19.02
	pointsOfFunctionPlot[1_902].Y = 0.148

	pointsOfFunctionPlot[1_903].X = 19.03
	pointsOfFunctionPlot[1_903].Y = 0.149

	pointsOfFunctionPlot[1_904].X = 19.04
	pointsOfFunctionPlot[1_904].Y = 0.15

	pointsOfFunctionPlot[1_905].X = 19.05
	pointsOfFunctionPlot[1_905].Y = 0.151

	pointsOfFunctionPlot[1_906].X = 19.06
	pointsOfFunctionPlot[1_906].Y = 0.152

	pointsOfFunctionPlot[1_907].X = 19.07
	pointsOfFunctionPlot[1_907].Y = 0.153

	pointsOfFunctionPlot[1_908].X = 19.08
	pointsOfFunctionPlot[1_908].Y = 0.154

	pointsOfFunctionPlot[1_909].X = 19.09
	pointsOfFunctionPlot[1_909].Y = 0.155

	pointsOfFunctionPlot[1_910].X = 19.10
	pointsOfFunctionPlot[1_910].Y = 0.156

	pointsOfFunctionPlot[1_911].X = 19.11
	pointsOfFunctionPlot[1_911].Y = 0.157

	pointsOfFunctionPlot[1_912].X = 19.12
	pointsOfFunctionPlot[1_912].Y = 0.158

	pointsOfFunctionPlot[1_913].X = 19.13
	pointsOfFunctionPlot[1_913].Y = 0.159

	pointsOfFunctionPlot[1_914].X = 19.14
	pointsOfFunctionPlot[1_914].Y = 0.159

	pointsOfFunctionPlot[1_915].X = 19.15
	pointsOfFunctionPlot[1_915].Y = 0.16

	pointsOfFunctionPlot[1_916].X = 19.16
	pointsOfFunctionPlot[1_916].Y = 0.161

	pointsOfFunctionPlot[1_917].X = 19.17
	pointsOfFunctionPlot[1_917].Y = 0.162

	pointsOfFunctionPlot[1_918].X = 19.18
	pointsOfFunctionPlot[1_918].Y = 0.163

	pointsOfFunctionPlot[1_919].X = 19.19
	pointsOfFunctionPlot[1_919].Y = 0.163

	pointsOfFunctionPlot[1_920].X = 19.20
	pointsOfFunctionPlot[1_920].Y = 0.164

	pointsOfFunctionPlot[1_921].X = 19.21
	pointsOfFunctionPlot[1_921].Y = 0.165

	pointsOfFunctionPlot[1_922].X = 19.22
	pointsOfFunctionPlot[1_922].Y = 0.166

	pointsOfFunctionPlot[1_923].X = 19.23
	pointsOfFunctionPlot[1_923].Y = 0.166

	pointsOfFunctionPlot[1_924].X = 19.24
	pointsOfFunctionPlot[1_924].Y = 0.167

	pointsOfFunctionPlot[1_925].X = 19.25
	pointsOfFunctionPlot[1_925].Y = 0.168

	pointsOfFunctionPlot[1_926].X = 19.26
	pointsOfFunctionPlot[1_926].Y = 0.168

	pointsOfFunctionPlot[1_927].X = 19.27
	pointsOfFunctionPlot[1_927].Y = 0.169

	pointsOfFunctionPlot[1_928].X = 19.28
	pointsOfFunctionPlot[1_928].Y = 0.169

	pointsOfFunctionPlot[1_929].X = 19.29
	pointsOfFunctionPlot[1_929].Y = 0.17

	pointsOfFunctionPlot[1_930].X = 19.30
	pointsOfFunctionPlot[1_930].Y = 0.171

	pointsOfFunctionPlot[1_931].X = 19.31
	pointsOfFunctionPlot[1_931].Y = 0.171

	pointsOfFunctionPlot[1_932].X = 19.32
	pointsOfFunctionPlot[1_932].Y = 0.172

	pointsOfFunctionPlot[1_933].X = 19.33
	pointsOfFunctionPlot[1_933].Y = 0.172

	pointsOfFunctionPlot[1_934].X = 19.34
	pointsOfFunctionPlot[1_934].Y = 0.173

	pointsOfFunctionPlot[1_935].X = 19.35
	pointsOfFunctionPlot[1_935].Y = 0.173

	pointsOfFunctionPlot[1_936].X = 19.36
	pointsOfFunctionPlot[1_936].Y = 0.174

	pointsOfFunctionPlot[1_937].X = 19.37
	pointsOfFunctionPlot[1_937].Y = 0.174

	pointsOfFunctionPlot[1_938].X = 19.38
	pointsOfFunctionPlot[1_938].Y = 0.175

	pointsOfFunctionPlot[1_939].X = 19.39
	pointsOfFunctionPlot[1_939].Y = 0.175

	pointsOfFunctionPlot[1_940].X = 19.40
	pointsOfFunctionPlot[1_940].Y = 0.175

	pointsOfFunctionPlot[1_941].X = 19.41
	pointsOfFunctionPlot[1_941].Y = 0.176

	pointsOfFunctionPlot[1_942].X = 19.42
	pointsOfFunctionPlot[1_942].Y = 0.176

	pointsOfFunctionPlot[1_943].X = 19.43
	pointsOfFunctionPlot[1_943].Y = 0.176

	pointsOfFunctionPlot[1_944].X = 19.44
	pointsOfFunctionPlot[1_944].Y = 0.177

	pointsOfFunctionPlot[1_945].X = 19.45
	pointsOfFunctionPlot[1_945].Y = 0.177

	pointsOfFunctionPlot[1_946].X = 19.46
	pointsOfFunctionPlot[1_946].Y = 0.177

	pointsOfFunctionPlot[1_947].X = 19.47
	pointsOfFunctionPlot[1_947].Y = 0.177

	pointsOfFunctionPlot[1_948].X = 19.48
	pointsOfFunctionPlot[1_948].Y = 0.178

	pointsOfFunctionPlot[1_949].X = 19.49
	pointsOfFunctionPlot[1_949].Y = 0.178

	pointsOfFunctionPlot[1_950].X = 19.50
	pointsOfFunctionPlot[1_950].Y = 0.178

	pointsOfFunctionPlot[1_951].X = 19.51
	pointsOfFunctionPlot[1_951].Y = 0.179

	pointsOfFunctionPlot[1_952].X = 19.52
	pointsOfFunctionPlot[1_952].Y = 0.179

	pointsOfFunctionPlot[1_953].X = 19.53
	pointsOfFunctionPlot[1_953].Y = 0.179

	pointsOfFunctionPlot[1_954].X = 19.54
	pointsOfFunctionPlot[1_954].Y = 0.179

	pointsOfFunctionPlot[1_955].X = 19.55
	pointsOfFunctionPlot[1_955].Y = 0.179

	pointsOfFunctionPlot[1_956].X = 19.56
	pointsOfFunctionPlot[1_956].Y = 0.179

	pointsOfFunctionPlot[1_957].X = 19.57
	pointsOfFunctionPlot[1_957].Y = 0.179

	pointsOfFunctionPlot[1_958].X = 19.58
	pointsOfFunctionPlot[1_958].Y = 0.179

	pointsOfFunctionPlot[1_959].X = 19.59
	pointsOfFunctionPlot[1_959].Y = 0.18

	pointsOfFunctionPlot[1_960].X = 19.60
	pointsOfFunctionPlot[1_960].Y = 0.18

	pointsOfFunctionPlot[1_961].X = 19.61
	pointsOfFunctionPlot[1_961].Y = 0.18

	pointsOfFunctionPlot[1_962].X = 19.62
	pointsOfFunctionPlot[1_962].Y = 0.18

	pointsOfFunctionPlot[1_963].X = 19.63
	pointsOfFunctionPlot[1_963].Y = 0.18

	pointsOfFunctionPlot[1_964].X = 19.64
	pointsOfFunctionPlot[1_964].Y = 0.18

	pointsOfFunctionPlot[1_965].X = 19.65
	pointsOfFunctionPlot[1_965].Y = 0.179

	pointsOfFunctionPlot[1_966].X = 19.66
	pointsOfFunctionPlot[1_966].Y = 0.179

	pointsOfFunctionPlot[1_967].X = 19.67
	pointsOfFunctionPlot[1_967].Y = 0.179

	pointsOfFunctionPlot[1_968].X = 19.68
	pointsOfFunctionPlot[1_968].Y = 0.179

	pointsOfFunctionPlot[1_969].X = 19.69
	pointsOfFunctionPlot[1_969].Y = 0.179

	pointsOfFunctionPlot[1_970].X = 19.70
	pointsOfFunctionPlot[1_970].Y = 0.179

	pointsOfFunctionPlot[1_971].X = 19.71
	pointsOfFunctionPlot[1_971].Y = 0.179

	pointsOfFunctionPlot[1_972].X = 19.72
	pointsOfFunctionPlot[1_972].Y = 0.179

	pointsOfFunctionPlot[1_973].X = 19.73
	pointsOfFunctionPlot[1_973].Y = 0.178

	pointsOfFunctionPlot[1_974].X = 19.74
	pointsOfFunctionPlot[1_974].Y = 0.178

	pointsOfFunctionPlot[1_975].X = 19.75
	pointsOfFunctionPlot[1_975].Y = 0.178

	pointsOfFunctionPlot[1_976].X = 19.76
	pointsOfFunctionPlot[1_976].Y = 0.178

	pointsOfFunctionPlot[1_977].X = 19.77
	pointsOfFunctionPlot[1_977].Y = 0.177

	pointsOfFunctionPlot[1_978].X = 19.78
	pointsOfFunctionPlot[1_978].Y = 0.177

	pointsOfFunctionPlot[1_979].X = 19.79
	pointsOfFunctionPlot[1_979].Y = 0.177

	pointsOfFunctionPlot[1_980].X = 19.80
	pointsOfFunctionPlot[1_980].Y = 0.177

	pointsOfFunctionPlot[1_981].X = 19.81
	pointsOfFunctionPlot[1_981].Y = 0.176

	pointsOfFunctionPlot[1_982].X = 19.82
	pointsOfFunctionPlot[1_982].Y = 0.176

	pointsOfFunctionPlot[1_983].X = 19.83
	pointsOfFunctionPlot[1_983].Y = 0.175

	pointsOfFunctionPlot[1_984].X = 19.84
	pointsOfFunctionPlot[1_984].Y = 0.175

	pointsOfFunctionPlot[1_985].X = 19.85
	pointsOfFunctionPlot[1_985].Y = 0.175

	pointsOfFunctionPlot[1_986].X = 19.86
	pointsOfFunctionPlot[1_986].Y = 0.174

	pointsOfFunctionPlot[1_987].X = 19.87
	pointsOfFunctionPlot[1_987].Y = 0.174

	pointsOfFunctionPlot[1_988].X = 19.88
	pointsOfFunctionPlot[1_988].Y = 0.173

	pointsOfFunctionPlot[1_989].X = 19.89
	pointsOfFunctionPlot[1_989].Y = 0.173

	pointsOfFunctionPlot[1_990].X = 19.90
	pointsOfFunctionPlot[1_990].Y = 0.172

	pointsOfFunctionPlot[1_991].X = 19.91
	pointsOfFunctionPlot[1_991].Y = 0.172

	pointsOfFunctionPlot[1_992].X = 19.92
	pointsOfFunctionPlot[1_992].Y = 0.171

	pointsOfFunctionPlot[1_993].X = 19.93
	pointsOfFunctionPlot[1_993].Y = 0.171

	pointsOfFunctionPlot[1_994].X = 19.94
	pointsOfFunctionPlot[1_994].Y = 0.17

	pointsOfFunctionPlot[1_995].X = 19.95
	pointsOfFunctionPlot[1_995].Y = 0.17

	pointsOfFunctionPlot[1_996].X = 19.96
	pointsOfFunctionPlot[1_996].Y = 0.169

	pointsOfFunctionPlot[1_997].X = 19.97
	pointsOfFunctionPlot[1_997].Y = 0.168

	pointsOfFunctionPlot[1_998].X = 19.98
	pointsOfFunctionPlot[1_998].Y = 0.168

	pointsOfFunctionPlot[1_999].X = 19.99
	pointsOfFunctionPlot[1_999].Y = 0.167

	pointsOfFunctionPlot[2_000].X = 20.0
	pointsOfFunctionPlot[2_000].Y = 0.167










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function J0(x)"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("J0(x)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"Bessel-function-0th-order-plot-01.png"); err != nil {

		panic(err)
	}
}
