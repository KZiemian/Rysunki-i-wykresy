package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function f(x) = x * exp(-(x - 2.0)).

	pointsOfFunctionPlot := make(plotter.XYs, 1_142)

	pointsOfFunctionPlot[0].X = 0.0
	pointsOfFunctionPlot[0].Y = 0.0

	pointsOfFunctionPlot[1].X = 0.01
	pointsOfFunctionPlot[1].Y = 0.073

	pointsOfFunctionPlot[2].X = 0.02
	pointsOfFunctionPlot[2].Y = 0.144

	pointsOfFunctionPlot[3].X = 0.03
	pointsOfFunctionPlot[3].Y = 0.215

	pointsOfFunctionPlot[4].X = 0.04
	pointsOfFunctionPlot[4].Y = 0.284

	pointsOfFunctionPlot[5].X = 0.05
	pointsOfFunctionPlot[5].Y = 0.351

	pointsOfFunctionPlot[6].X = 0.06
	pointsOfFunctionPlot[6].Y = 0.417

	pointsOfFunctionPlot[7].X = 0.07
	pointsOfFunctionPlot[7].Y = 0.482

	pointsOfFunctionPlot[8].X = 0.08
	pointsOfFunctionPlot[8].Y = 0.545

	pointsOfFunctionPlot[9].X = 0.09
	pointsOfFunctionPlot[9].Y = 0.607

	pointsOfFunctionPlot[10].X = 0.1
	pointsOfFunctionPlot[10].Y = 0.668

	pointsOfFunctionPlot[11].X = 0.11
	pointsOfFunctionPlot[11].Y = 0.728

	pointsOfFunctionPlot[12].X = 0.12
	pointsOfFunctionPlot[12].Y = 0.786

	pointsOfFunctionPlot[13].X = 0.13
	pointsOfFunctionPlot[13].Y = 0.843

	pointsOfFunctionPlot[14].X = 0.14
	pointsOfFunctionPlot[14].Y = 0.899

	pointsOfFunctionPlot[15].X = 0.15
	pointsOfFunctionPlot[15].Y = 0.954

	pointsOfFunctionPlot[16].X = 0.16
	pointsOfFunctionPlot[16].Y = 1.007

	pointsOfFunctionPlot[17].X = 0.17
	pointsOfFunctionPlot[17].Y = 1.059

	pointsOfFunctionPlot[18].X = 0.18
	pointsOfFunctionPlot[18].Y = 1.11

	pointsOfFunctionPlot[19].X = 0.19
	pointsOfFunctionPlot[19].Y = 1.161

	pointsOfFunctionPlot[20].X = 0.2
	pointsOfFunctionPlot[20].Y = 1.209

	pointsOfFunctionPlot[21].X = 0.21
	pointsOfFunctionPlot[21].Y = 1.257

	pointsOfFunctionPlot[22].X = 0.22
	pointsOfFunctionPlot[22].Y = 1.304

	pointsOfFunctionPlot[23].X = 0.23
	pointsOfFunctionPlot[23].Y = 1.35

	pointsOfFunctionPlot[24].X = 0.24
	pointsOfFunctionPlot[24].Y = 1.395

	pointsOfFunctionPlot[25].X = 0.25
	pointsOfFunctionPlot[25].Y = 1.438

	pointsOfFunctionPlot[26].X = 0.26
	pointsOfFunctionPlot[26].Y = 1.481

	pointsOfFunctionPlot[27].X = 0.27
	pointsOfFunctionPlot[27].Y = 1.523

	pointsOfFunctionPlot[28].X = 0.28
	pointsOfFunctionPlot[28].Y = 1.563

	pointsOfFunctionPlot[29].X = 0.29
	pointsOfFunctionPlot[29].Y = 1.603

	pointsOfFunctionPlot[30].X = 0.3
	pointsOfFunctionPlot[30].Y = 1.642

	pointsOfFunctionPlot[31].X = 0.31
	pointsOfFunctionPlot[31].Y = 1.68

	pointsOfFunctionPlot[32].X = 0.32
	pointsOfFunctionPlot[32].Y = 1.717

	pointsOfFunctionPlot[33].X = 0.33
	pointsOfFunctionPlot[33].Y = 1.753

	pointsOfFunctionPlot[34].X = 0.34
	pointsOfFunctionPlot[34].Y = 1.788

	pointsOfFunctionPlot[35].X = 0.35
	pointsOfFunctionPlot[35].Y = 1.822

	pointsOfFunctionPlot[36].X = 0.36
	pointsOfFunctionPlot[36].Y = 1.855

	pointsOfFunctionPlot[37].X = 0.37
	pointsOfFunctionPlot[37].Y = 1.888

	pointsOfFunctionPlot[38].X = 0.38
	pointsOfFunctionPlot[38].Y = 1.92

	pointsOfFunctionPlot[39].X = 0.39
	pointsOfFunctionPlot[39].Y = 1.951

	pointsOfFunctionPlot[40].X = 0.40
	pointsOfFunctionPlot[40].Y = 1.981

	pointsOfFunctionPlot[41].X = 0.41
	pointsOfFunctionPlot[41].Y = 2.01

	pointsOfFunctionPlot[42].X = 0.42
	pointsOfFunctionPlot[42].Y = 2.039

	pointsOfFunctionPlot[43].X = 0.43
	pointsOfFunctionPlot[43].Y = 2.066

	pointsOfFunctionPlot[44].X = 0.44
	pointsOfFunctionPlot[44].Y = 2.093

	pointsOfFunctionPlot[45].X = 0.45
	pointsOfFunctionPlot[45].Y = 2.12

	pointsOfFunctionPlot[46].X = 0.46
	pointsOfFunctionPlot[46].Y = 2.145

	pointsOfFunctionPlot[47].X = 0.47
	pointsOfFunctionPlot[47].Y = 2.17

	pointsOfFunctionPlot[48].X = 0.48
	pointsOfFunctionPlot[48].Y = 2.194

	pointsOfFunctionPlot[49].X = 0.49
	pointsOfFunctionPlot[49].Y = 2.218

	pointsOfFunctionPlot[50].X = 0.5
	pointsOfFunctionPlot[50].Y = 2.24

	pointsOfFunctionPlot[51].X = 0.51
	pointsOfFunctionPlot[51].Y = 2.262

	pointsOfFunctionPlot[52].X = 0.52
	pointsOfFunctionPlot[52].Y = 2.284

	pointsOfFunctionPlot[53].X = 0.53
	pointsOfFunctionPlot[53].Y = 2.305

	pointsOfFunctionPlot[54].X = 0.54
	pointsOfFunctionPlot[54].Y = 2.325

	pointsOfFunctionPlot[55].X = 0.55
	pointsOfFunctionPlot[55].Y = 2.344

	pointsOfFunctionPlot[56].X = 0.56
	pointsOfFunctionPlot[56].Y = 2.363

	pointsOfFunctionPlot[57].X = 0.57
	pointsOfFunctionPlot[57].Y = 2.381

	pointsOfFunctionPlot[58].X = 0.58
	pointsOfFunctionPlot[58].Y = 2.399

	pointsOfFunctionPlot[59].X = 0.59
	pointsOfFunctionPlot[59].Y = 2.416

	pointsOfFunctionPlot[60].X = 0.6
	pointsOfFunctionPlot[60].Y = 2.433

	pointsOfFunctionPlot[61].X = 0.61
	pointsOfFunctionPlot[61].Y = 2.449

	pointsOfFunctionPlot[62].X = 0.62
	pointsOfFunctionPlot[62].Y = 2.464

	pointsOfFunctionPlot[63].X = 0.63
	pointsOfFunctionPlot[63].Y = 2.479

	pointsOfFunctionPlot[64].X = 0.64
	pointsOfFunctionPlot[64].Y = 2.493

	pointsOfFunctionPlot[65].X = 0.65
	pointsOfFunctionPlot[65].Y = 2.507

	pointsOfFunctionPlot[66].X = 0.66
	pointsOfFunctionPlot[66].Y = 2.52

	pointsOfFunctionPlot[67].X = 0.67
	pointsOfFunctionPlot[67].Y = 2.533

	pointsOfFunctionPlot[68].X = 0.68
	pointsOfFunctionPlot[68].Y = 2.545

	pointsOfFunctionPlot[69].X = 0.69
	pointsOfFunctionPlot[69].Y = 2.557

	pointsOfFunctionPlot[70].X = 0.7
	pointsOfFunctionPlot[70].Y = 2.568

	pointsOfFunctionPlot[71].X = 0.71
	pointsOfFunctionPlot[71].Y = 2.579

	pointsOfFunctionPlot[72].X = 0.72
	pointsOfFunctionPlot[72].Y = 2.589

	pointsOfFunctionPlot[73].X = 0.73
	pointsOfFunctionPlot[73].Y = 2.599

	pointsOfFunctionPlot[74].X = 0.74
	pointsOfFunctionPlot[74].Y = 2.608

	pointsOfFunctionPlot[75].X = 0.75
	pointsOfFunctionPlot[75].Y = 2.617

	pointsOfFunctionPlot[76].X = 0.76
	pointsOfFunctionPlot[76].Y = 2.626

	pointsOfFunctionPlot[77].X = 0.77
	pointsOfFunctionPlot[77].Y = 2.634

	pointsOfFunctionPlot[78].X = 0.78
	pointsOfFunctionPlot[78].Y = 2.642

	pointsOfFunctionPlot[79].X = 0.79
	pointsOfFunctionPlot[79].Y = 2.649

	pointsOfFunctionPlot[80].X = 0.8
	pointsOfFunctionPlot[80].Y = 2.656

	pointsOfFunctionPlot[81].X = 0.81
	pointsOfFunctionPlot[81].Y = 2.662

	pointsOfFunctionPlot[82].X = 0.82
	pointsOfFunctionPlot[82].Y = 2.668

	pointsOfFunctionPlot[83].X = 0.83
	pointsOfFunctionPlot[83].Y = 2.674

	pointsOfFunctionPlot[84].X = 0.84
	pointsOfFunctionPlot[84].Y = 2.679

	pointsOfFunctionPlot[85].X = 0.85
	pointsOfFunctionPlot[85].Y = 2.684

	pointsOfFunctionPlot[86].X = 0.86
	pointsOfFunctionPlot[86].Y = 2.689

	pointsOfFunctionPlot[87].X = 0.87
	pointsOfFunctionPlot[87].Y = 2.693

	pointsOfFunctionPlot[88].X = 0.88
	pointsOfFunctionPlot[88].Y = 2.697

	pointsOfFunctionPlot[89].X = 0.89
	pointsOfFunctionPlot[89].Y = 2.7

	pointsOfFunctionPlot[90].X = 0.9
	pointsOfFunctionPlot[90].Y = 2.703

	pointsOfFunctionPlot[91].X = 0.91
	pointsOfFunctionPlot[91].Y = 2.706

	pointsOfFunctionPlot[92].X = 0.92
	pointsOfFunctionPlot[92].Y = 2.709

	pointsOfFunctionPlot[93].X = 0.93
	pointsOfFunctionPlot[93].Y = 2.711

	pointsOfFunctionPlot[94].X = 0.94
	pointsOfFunctionPlot[94].Y = 2.713

	pointsOfFunctionPlot[95].X = 0.95
	pointsOfFunctionPlot[95].Y = 2.714

	pointsOfFunctionPlot[96].X = 0.96
	pointsOfFunctionPlot[96].Y = 2.716

	pointsOfFunctionPlot[97].X = 0.97
	pointsOfFunctionPlot[97].Y = 2.717

	pointsOfFunctionPlot[98].X = 0.98
	pointsOfFunctionPlot[98].Y = 2.717

	pointsOfFunctionPlot[99].X = 0.99
	pointsOfFunctionPlot[99].Y = 2.718

	pointsOfFunctionPlot[100].X = 1.0
	pointsOfFunctionPlot[100].Y = 2.718

	pointsOfFunctionPlot[101].X = 1.01
	pointsOfFunctionPlot[101].Y = 2.718

	pointsOfFunctionPlot[102].X = 1.02
	pointsOfFunctionPlot[102].Y = 2.717

	pointsOfFunctionPlot[103].X = 1.03
	pointsOfFunctionPlot[103].Y = 2.717

	pointsOfFunctionPlot[104].X = 1.04
	pointsOfFunctionPlot[104].Y = 2.716

	pointsOfFunctionPlot[105].X = 1.05
	pointsOfFunctionPlot[105].Y = 2.715

	pointsOfFunctionPlot[106].X = 1.06
	pointsOfFunctionPlot[106].Y = 2.713

	pointsOfFunctionPlot[107].X = 1.07
	pointsOfFunctionPlot[107].Y = 2.711

	pointsOfFunctionPlot[108].X = 1.08
	pointsOfFunctionPlot[108].Y = 2.71

	pointsOfFunctionPlot[109].X = 1.09
	pointsOfFunctionPlot[109].Y = 2.707

	pointsOfFunctionPlot[110].X = 1.1
	pointsOfFunctionPlot[110].Y = 2.705

	pointsOfFunctionPlot[111].X = 1.11
	pointsOfFunctionPlot[111].Y = 2.703

	pointsOfFunctionPlot[112].X = 1.12
	pointsOfFunctionPlot[112].Y = 2.7

	pointsOfFunctionPlot[113].X = 1.13
	pointsOfFunctionPlot[113].Y = 2.697

	pointsOfFunctionPlot[114].X = 1.14
	pointsOfFunctionPlot[114].Y = 2.694

	pointsOfFunctionPlot[115].X = 1.15
	pointsOfFunctionPlot[115].Y = 2.69

	pointsOfFunctionPlot[116].X = 1.16
	pointsOfFunctionPlot[116].Y = 2.687

	pointsOfFunctionPlot[117].X = 1.17
	pointsOfFunctionPlot[117].Y = 2.683

	pointsOfFunctionPlot[118].X = 1.18
	pointsOfFunctionPlot[118].Y = 2.679

	pointsOfFunctionPlot[119].X = 1.19
	pointsOfFunctionPlot[119].Y = 2.675

	pointsOfFunctionPlot[120].X = 1.2
	pointsOfFunctionPlot[120].Y = 2.67

	pointsOfFunctionPlot[121].X = 1.21
	pointsOfFunctionPlot[121].Y = 2.666

	pointsOfFunctionPlot[122].X = 1.22
	pointsOfFunctionPlot[122].Y = 2.661

	pointsOfFunctionPlot[123].X = 1.23
	pointsOfFunctionPlot[123].Y = 2.656

	pointsOfFunctionPlot[124].X = 1.24
	pointsOfFunctionPlot[124].Y = 2.651

	pointsOfFunctionPlot[125].X = 1.25
	pointsOfFunctionPlot[125].Y = 2.646

	pointsOfFunctionPlot[126].X = 1.26
	pointsOfFunctionPlot[126].Y = 2.64

	pointsOfFunctionPlot[127].X = 1.27
	pointsOfFunctionPlot[127].Y = 2.635

	pointsOfFunctionPlot[128].X = 1.28
	pointsOfFunctionPlot[128].Y = 2.629

	pointsOfFunctionPlot[129].X = 1.29
	pointsOfFunctionPlot[129].Y = 2.623

	pointsOfFunctionPlot[130].X = 1.3
	pointsOfFunctionPlot[130].Y = 2.617

	pointsOfFunctionPlot[131].X = 1.31
	pointsOfFunctionPlot[131].Y = 2.611

	pointsOfFunctionPlot[132].X = 1.32
	pointsOfFunctionPlot[132].Y = 2.605

	pointsOfFunctionPlot[133].X = 1.33
	pointsOfFunctionPlot[133].Y = 2.599

	pointsOfFunctionPlot[134].X = 1.34
	pointsOfFunctionPlot[134].Y = 2.592

	pointsOfFunctionPlot[135].X = 1.35
	pointsOfFunctionPlot[135].Y = 2.586

	pointsOfFunctionPlot[136].X = 1.36
	pointsOfFunctionPlot[136].Y = 2.579

	pointsOfFunctionPlot[137].X = 1.37
	pointsOfFunctionPlot[137].Y = 2.572

	pointsOfFunctionPlot[138].X = 1.38
	pointsOfFunctionPlot[138].Y = 2.565

	pointsOfFunctionPlot[139].X = 1.39
	pointsOfFunctionPlot[139].Y = 2.558

	pointsOfFunctionPlot[140].X = 1.4
	pointsOfFunctionPlot[140].Y = 2.551

	pointsOfFunctionPlot[141].X = 1.41
	pointsOfFunctionPlot[141].Y = 2.543

	pointsOfFunctionPlot[142].X = 1.42
	pointsOfFunctionPlot[142].Y = 2.536

	pointsOfFunctionPlot[143].X = 1.43
	pointsOfFunctionPlot[143].Y = 2.528

	pointsOfFunctionPlot[144].X = 1.44
	pointsOfFunctionPlot[144].Y = 2.521

	pointsOfFunctionPlot[145].X = 1.45
	pointsOfFunctionPlot[145].Y = 2.513

	pointsOfFunctionPlot[146].X = 1.46
	pointsOfFunctionPlot[146].Y = 2.505

	pointsOfFunctionPlot[147].X = 1.47
	pointsOfFunctionPlot[147].Y = 2.497

	pointsOfFunctionPlot[148].X = 1.48
	pointsOfFunctionPlot[148].Y = 2.489

	pointsOfFunctionPlot[149].X = 1.49
	pointsOfFunctionPlot[149].Y = 2.481

	pointsOfFunctionPlot[150].X = 1.5
	pointsOfFunctionPlot[150].Y = 2.473

	pointsOfFunctionPlot[151].X = 1.51
	pointsOfFunctionPlot[151].Y = 2.464

	pointsOfFunctionPlot[152].X = 1.52
	pointsOfFunctionPlot[152].Y = 2.456

	pointsOfFunctionPlot[153].X = 1.53
	pointsOfFunctionPlot[153].Y = 2.448

	pointsOfFunctionPlot[154].X = 1.54
	pointsOfFunctionPlot[154].Y = 2.439

	pointsOfFunctionPlot[155].X = 1.55
	pointsOfFunctionPlot[155].Y = 2.43

	pointsOfFunctionPlot[156].X = 1.56
	pointsOfFunctionPlot[156].Y = 2.422

	pointsOfFunctionPlot[157].X = 1.57
	pointsOfFunctionPlot[157].Y = 2.413

	pointsOfFunctionPlot[158].X = 1.58
	pointsOfFunctionPlot[158].Y = 2.404

	pointsOfFunctionPlot[159].X = 1.59
	pointsOfFunctionPlot[159].Y = 2.395

	pointsOfFunctionPlot[160].X = 1.6
	pointsOfFunctionPlot[160].Y = 2.386

	pointsOfFunctionPlot[161].X = 1.61
	pointsOfFunctionPlot[161].Y = 2.377

	pointsOfFunctionPlot[162].X = 1.62
	pointsOfFunctionPlot[162].Y = 2.368

	pointsOfFunctionPlot[163].X = 1.63
	pointsOfFunctionPlot[163].Y = 2.359

	pointsOfFunctionPlot[164].X = 1.64
	pointsOfFunctionPlot[164].Y = 2.35

	pointsOfFunctionPlot[165].X = 1.65
	pointsOfFunctionPlot[165].Y = 2.341

	pointsOfFunctionPlot[166].X = 1.66
	pointsOfFunctionPlot[166].Y = 2.332

	pointsOfFunctionPlot[167].X = 1.67
	pointsOfFunctionPlot[167].Y = 2.322

	pointsOfFunctionPlot[168].X = 1.68
	pointsOfFunctionPlot[168].Y = 2.313

	pointsOfFunctionPlot[169].X = 1.69
	pointsOfFunctionPlot[169].Y = 2.304

	pointsOfFunctionPlot[170].X = 1.7
	pointsOfFunctionPlot[170].Y = 2.294

	pointsOfFunctionPlot[171].X = 1.71
	pointsOfFunctionPlot[171].Y = 2.285

	pointsOfFunctionPlot[172].X = 1.72
	pointsOfFunctionPlot[172].Y = 2.275

	pointsOfFunctionPlot[173].X = 1.73
	pointsOfFunctionPlot[173].Y = 2.266

	pointsOfFunctionPlot[174].X = 1.74
	pointsOfFunctionPlot[174].Y = 2.256

	pointsOfFunctionPlot[175].X = 1.75
	pointsOfFunctionPlot[175].Y = 2.247

	pointsOfFunctionPlot[176].X = 1.76
	pointsOfFunctionPlot[176].Y = 2.237

	pointsOfFunctionPlot[177].X = 1.77
	pointsOfFunctionPlot[177].Y = 2.227

	pointsOfFunctionPlot[178].X = 1.78
	pointsOfFunctionPlot[178].Y = 2.218

	pointsOfFunctionPlot[179].X = 1.79
	pointsOfFunctionPlot[179].Y = 2.208

	pointsOfFunctionPlot[180].X = 1.8
	pointsOfFunctionPlot[180].Y = 2.198

	pointsOfFunctionPlot[181].X = 1.81
	pointsOfFunctionPlot[181].Y = 2.188

	pointsOfFunctionPlot[182].X = 1.82
	pointsOfFunctionPlot[182].Y = 2.178

	pointsOfFunctionPlot[183].X = 1.83
	pointsOfFunctionPlot[183].Y = 2.169

	pointsOfFunctionPlot[184].X = 1.84
	pointsOfFunctionPlot[184].Y = 2.159

	pointsOfFunctionPlot[185].X = 1.85
	pointsOfFunctionPlot[185].Y = 2.149

	pointsOfFunctionPlot[186].X = 1.86
	pointsOfFunctionPlot[186].Y = 2.139

	pointsOfFunctionPlot[187].X = 1.87
	pointsOfFunctionPlot[187].Y = 2.129

	pointsOfFunctionPlot[188].X = 1.88
	pointsOfFunctionPlot[188].Y = 2.119

	pointsOfFunctionPlot[189].X = 1.89
	pointsOfFunctionPlot[189].Y = 2.109

	pointsOfFunctionPlot[190].X = 1.9
	pointsOfFunctionPlot[190].Y = 2.099

	pointsOfFunctionPlot[191].X = 1.91
	pointsOfFunctionPlot[191].Y = 2.089

	pointsOfFunctionPlot[192].X = 1.92
	pointsOfFunctionPlot[192].Y = 2.079

	pointsOfFunctionPlot[193].X = 1.93
	pointsOfFunctionPlot[193].Y = 2.069

	pointsOfFunctionPlot[194].X = 1.94
	pointsOfFunctionPlot[194].Y = 2.06

	pointsOfFunctionPlot[195].X = 1.95
	pointsOfFunctionPlot[195].Y = 2.05

	pointsOfFunctionPlot[196].X = 1.96
	pointsOfFunctionPlot[196].Y = 2.04

	pointsOfFunctionPlot[197].X = 1.97
	pointsOfFunctionPlot[197].Y = 2.03

	pointsOfFunctionPlot[198].X = 1.98
	pointsOfFunctionPlot[198].Y = 2.02

	pointsOfFunctionPlot[199].X = 1.99
	pointsOfFunctionPlot[199].Y = 2.01

	pointsOfFunctionPlot[200].X = 2.0
	pointsOfFunctionPlot[200].Y = 2.0

	pointsOfFunctionPlot[201].X = 2.01
	pointsOfFunctionPlot[201].Y = 1.99

	pointsOfFunctionPlot[202].X = 2.02
	pointsOfFunctionPlot[202].Y = 1.98

	pointsOfFunctionPlot[203].X = 2.03
	pointsOfFunctionPlot[203].Y = 1.97

	pointsOfFunctionPlot[204].X = 2.04
	pointsOfFunctionPlot[204].Y = 1.96

	pointsOfFunctionPlot[205].X = 2.05
	pointsOfFunctionPlot[205].Y = 1.95

	pointsOfFunctionPlot[206].X = 2.06
	pointsOfFunctionPlot[206].Y = 1.94

	pointsOfFunctionPlot[207].X = 2.07
	pointsOfFunctionPlot[207].Y = 1.93

	pointsOfFunctionPlot[208].X = 2.08
	pointsOfFunctionPlot[208].Y = 1.92

	pointsOfFunctionPlot[209].X = 2.09
	pointsOfFunctionPlot[209].Y = 1.91

	pointsOfFunctionPlot[210].X = 2.1
	pointsOfFunctionPlot[210].Y = 1.9

	pointsOfFunctionPlot[211].X = 2.11
	pointsOfFunctionPlot[211].Y = 1.89

	pointsOfFunctionPlot[212].X = 2.12
	pointsOfFunctionPlot[212].Y = 1.88

	pointsOfFunctionPlot[213].X = 2.13
	pointsOfFunctionPlot[213].Y = 1.87

	pointsOfFunctionPlot[214].X = 2.14
	pointsOfFunctionPlot[214].Y = 1.86

	pointsOfFunctionPlot[215].X = 2.15
	pointsOfFunctionPlot[215].Y = 1.85

	pointsOfFunctionPlot[216].X = 2.16
	pointsOfFunctionPlot[216].Y = 1.84

	pointsOfFunctionPlot[217].X = 2.17
	pointsOfFunctionPlot[217].Y = 1.83

	pointsOfFunctionPlot[218].X = 2.18
	pointsOfFunctionPlot[218].Y = 1.82

	pointsOfFunctionPlot[219].X = 2.19
	pointsOfFunctionPlot[219].Y = 1.811

	pointsOfFunctionPlot[220].X = 2.2
	pointsOfFunctionPlot[220].Y = 1.801

	pointsOfFunctionPlot[221].X = 2.21
	pointsOfFunctionPlot[221].Y = 1.791

	pointsOfFunctionPlot[222].X = 2.22
	pointsOfFunctionPlot[222].Y = 1.781

	pointsOfFunctionPlot[223].X = 2.23
	pointsOfFunctionPlot[223].Y = 1.771

	pointsOfFunctionPlot[224].X = 2.24
	pointsOfFunctionPlot[224].Y = 1.762

	pointsOfFunctionPlot[225].X = 2.25
	pointsOfFunctionPlot[225].Y = 1.752

	pointsOfFunctionPlot[226].X = 2.26
	pointsOfFunctionPlot[226].Y = 1.742

	pointsOfFunctionPlot[227].X = 2.27
	pointsOfFunctionPlot[227].Y = 1.732

	pointsOfFunctionPlot[228].X = 2.28
	pointsOfFunctionPlot[228].Y = 1.723

	pointsOfFunctionPlot[229].X = 2.29
	pointsOfFunctionPlot[229].Y = 1.713

	pointsOfFunctionPlot[230].X = 2.3
	pointsOfFunctionPlot[230].Y = 1.703

	pointsOfFunctionPlot[231].X = 2.31
	pointsOfFunctionPlot[231].Y = 1.694

	pointsOfFunctionPlot[232].X = 2.32
	pointsOfFunctionPlot[232].Y = 1.684

	pointsOfFunctionPlot[233].X = 2.33
	pointsOfFunctionPlot[233].Y = 1.675

	pointsOfFunctionPlot[234].X = 2.34
	pointsOfFunctionPlot[234].Y = 1.665

	pointsOfFunctionPlot[235].X = 2.35
	pointsOfFunctionPlot[235].Y = 1.656

	pointsOfFunctionPlot[236].X = 2.36
	pointsOfFunctionPlot[236].Y = 1.646

	pointsOfFunctionPlot[237].X = 2.37
	pointsOfFunctionPlot[237].Y = 1.637

	pointsOfFunctionPlot[238].X = 2.38
	pointsOfFunctionPlot[238].Y = 1.627

	pointsOfFunctionPlot[239].X = 2.39
	pointsOfFunctionPlot[239].Y = 1.618

	pointsOfFunctionPlot[240].X = 2.4
	pointsOfFunctionPlot[240].Y = 1.608

	pointsOfFunctionPlot[241].X = 2.41
	pointsOfFunctionPlot[241].Y = 1.599

	pointsOfFunctionPlot[242].X = 2.42
	pointsOfFunctionPlot[242].Y = 1.59

	pointsOfFunctionPlot[243].X = 2.43
	pointsOfFunctionPlot[243].Y = 1.58

	pointsOfFunctionPlot[244].X = 2.44
	pointsOfFunctionPlot[244].Y = 1.571

	pointsOfFunctionPlot[245].X = 2.45
	pointsOfFunctionPlot[245].Y = 1.562

	pointsOfFunctionPlot[246].X = 2.46
	pointsOfFunctionPlot[246].Y = 1.553

	pointsOfFunctionPlot[247].X = 2.47
	pointsOfFunctionPlot[247].Y = 1.543

	pointsOfFunctionPlot[248].X = 2.48
	pointsOfFunctionPlot[248].Y = 1.534

	pointsOfFunctionPlot[249].X = 2.49
	pointsOfFunctionPlot[249].Y = 1.525

	pointsOfFunctionPlot[250].X = 2.5
	pointsOfFunctionPlot[250].Y = 1.516

	pointsOfFunctionPlot[251].X = 2.51
	pointsOfFunctionPlot[251].Y = 1.507

	pointsOfFunctionPlot[252].X = 2.52
	pointsOfFunctionPlot[252].Y = 1.498

	pointsOfFunctionPlot[253].X = 2.53
	pointsOfFunctionPlot[253].Y = 1.489

	pointsOfFunctionPlot[254].X = 2.54
	pointsOfFunctionPlot[254].Y = 1.48

	pointsOfFunctionPlot[255].X = 2.55
	pointsOfFunctionPlot[255].Y = 1.471

	pointsOfFunctionPlot[256].X = 2.56
	pointsOfFunctionPlot[256].Y = 1.462

	pointsOfFunctionPlot[257].X = 2.57
	pointsOfFunctionPlot[257].Y = 1.453

	pointsOfFunctionPlot[258].X = 2.58
	pointsOfFunctionPlot[258].Y = 1.444

	pointsOfFunctionPlot[259].X = 2.59
	pointsOfFunctionPlot[259].Y = 1.435

	pointsOfFunctionPlot[260].X = 2.6
	pointsOfFunctionPlot[260].Y = 1.426

	pointsOfFunctionPlot[261].X = 2.61
	pointsOfFunctionPlot[261].Y = 1.418

	pointsOfFunctionPlot[262].X = 2.62
	pointsOfFunctionPlot[262].Y = 1.409

	pointsOfFunctionPlot[263].X = 2.63
	pointsOfFunctionPlot[263].Y = 1.4

	pointsOfFunctionPlot[264].X = 2.64
	pointsOfFunctionPlot[264].Y = 1.392

	pointsOfFunctionPlot[265].X = 2.65
	pointsOfFunctionPlot[265].Y = 1.383

	pointsOfFunctionPlot[266].X = 2.66
	pointsOfFunctionPlot[266].Y = 1.374

	pointsOfFunctionPlot[267].X = 2.67
	pointsOfFunctionPlot[267].Y = 1.366

	pointsOfFunctionPlot[268].X = 2.68
	pointsOfFunctionPlot[268].Y = 1.357

	pointsOfFunctionPlot[269].X = 2.69
	pointsOfFunctionPlot[269].Y = 1.349

	pointsOfFunctionPlot[270].X = 2.7
	pointsOfFunctionPlot[270].Y = 1.34

	pointsOfFunctionPlot[271].X = 2.71
	pointsOfFunctionPlot[271].Y = 1.332

	pointsOfFunctionPlot[272].X = 2.72
	pointsOfFunctionPlot[272].Y = 1.324

	pointsOfFunctionPlot[273].X = 2.73
	pointsOfFunctionPlot[273].Y = 1.315

	pointsOfFunctionPlot[274].X = 2.74
	pointsOfFunctionPlot[274].Y = 1.307

	pointsOfFunctionPlot[275].X = 2.75
	pointsOfFunctionPlot[275].Y = 1.299

	pointsOfFunctionPlot[276].X = 2.76
	pointsOfFunctionPlot[276].Y = 1.29

	pointsOfFunctionPlot[277].X = 2.77
	pointsOfFunctionPlot[277].Y = 1.282

	pointsOfFunctionPlot[278].X = 2.78
	pointsOfFunctionPlot[278].Y = 1.274

	pointsOfFunctionPlot[279].X = 2.79
	pointsOfFunctionPlot[279].Y = 1.266

	pointsOfFunctionPlot[280].X = 2.8
	pointsOfFunctionPlot[280].Y = 1.258

	pointsOfFunctionPlot[281].X = 2.81
	pointsOfFunctionPlot[281].Y = 1.25

	pointsOfFunctionPlot[282].X = 2.82
	pointsOfFunctionPlot[282].Y = 1.242

	pointsOfFunctionPlot[283].X = 2.83
	pointsOfFunctionPlot[283].Y = 1.234

	pointsOfFunctionPlot[284].X = 2.84
	pointsOfFunctionPlot[284].Y = 1.226

	pointsOfFunctionPlot[285].X = 2.85
	pointsOfFunctionPlot[285].Y = 1.218

	pointsOfFunctionPlot[286].X = 2.86
	pointsOfFunctionPlot[286].Y = 1.21

	pointsOfFunctionPlot[287].X = 2.87
	pointsOfFunctionPlot[287].Y = 1.202

	pointsOfFunctionPlot[288].X = 2.88
	pointsOfFunctionPlot[288].Y = 1.194

	pointsOfFunctionPlot[289].X = 2.89
	pointsOfFunctionPlot[289].Y = 1.186

	pointsOfFunctionPlot[290].X = 2.9
	pointsOfFunctionPlot[290].Y = 1.179

	pointsOfFunctionPlot[291].X = 2.91
	pointsOfFunctionPlot[291].Y = 1.171

	pointsOfFunctionPlot[292].X = 2.92
	pointsOfFunctionPlot[292].Y = 1.163

	pointsOfFunctionPlot[293].X = 2.93
	pointsOfFunctionPlot[293].Y = 1.156

	pointsOfFunctionPlot[294].X = 2.94
	pointsOfFunctionPlot[294].Y = 1.148

	pointsOfFunctionPlot[295].X = 2.95
	pointsOfFunctionPlot[295].Y = 1.14

	pointsOfFunctionPlot[296].X = 2.96
	pointsOfFunctionPlot[296].Y = 1.133

	pointsOfFunctionPlot[297].X = 2.97
	pointsOfFunctionPlot[297].Y = 1.125

	pointsOfFunctionPlot[298].X = 2.98
	pointsOfFunctionPlot[298].Y = 1.118

	pointsOfFunctionPlot[299].X = 2.99
	pointsOfFunctionPlot[299].Y = 1.111

	pointsOfFunctionPlot[300].X = 3.0
	pointsOfFunctionPlot[300].Y = 1.103

	pointsOfFunctionPlot[301].X = 3.01
	pointsOfFunctionPlot[301].Y = 1.096

	pointsOfFunctionPlot[302].X = 3.02
	pointsOfFunctionPlot[302].Y = 1.089

	pointsOfFunctionPlot[303].X = 3.03
	pointsOfFunctionPlot[303].Y = 1.081

	pointsOfFunctionPlot[304].X = 3.04
	pointsOfFunctionPlot[304].Y = 1.074

	pointsOfFunctionPlot[305].X = 3.05
	pointsOfFunctionPlot[305].Y = 1.067

	pointsOfFunctionPlot[306].X = 3.06
	pointsOfFunctionPlot[306].Y = 1.06

	pointsOfFunctionPlot[307].X = 3.07
	pointsOfFunctionPlot[307].Y = 1.053

	pointsOfFunctionPlot[308].X = 3.08
	pointsOfFunctionPlot[308].Y = 1.046

	pointsOfFunctionPlot[309].X = 3.09
	pointsOfFunctionPlot[309].Y = 1.038

	pointsOfFunctionPlot[310].X = 3.1
	pointsOfFunctionPlot[310].Y = 1.031

	pointsOfFunctionPlot[311].X = 3.11
	pointsOfFunctionPlot[311].Y = 1.024

	pointsOfFunctionPlot[312].X = 3.12
	pointsOfFunctionPlot[312].Y = 1.018

	pointsOfFunctionPlot[313].X = 3.13
	pointsOfFunctionPlot[313].Y = 1.011

	pointsOfFunctionPlot[314].X = 3.14
	pointsOfFunctionPlot[314].Y = 1.004

	pointsOfFunctionPlot[315].X = 3.15
	pointsOfFunctionPlot[315].Y = 0.997

	pointsOfFunctionPlot[316].X = 3.16
	pointsOfFunctionPlot[316].Y = 0.99

	pointsOfFunctionPlot[317].X = 3.17
	pointsOfFunctionPlot[317].Y = 0.983

	pointsOfFunctionPlot[318].X = 3.18
	pointsOfFunctionPlot[318].Y = 0.977

	pointsOfFunctionPlot[319].X = 3.19
	pointsOfFunctionPlot[319].Y = 0.97

	pointsOfFunctionPlot[320].X = 3.2
	pointsOfFunctionPlot[320].Y = 0.963

	pointsOfFunctionPlot[321].X = 3.21
	pointsOfFunctionPlot[321].Y = 0.957

	pointsOfFunctionPlot[322].X = 3.22
	pointsOfFunctionPlot[322].Y = 0.95

	pointsOfFunctionPlot[323].X = 3.23
	pointsOfFunctionPlot[323].Y = 0.944

	pointsOfFunctionPlot[324].X = 3.24
	pointsOfFunctionPlot[324].Y = 0.937

	pointsOfFunctionPlot[325].X = 3.25
	pointsOfFunctionPlot[325].Y = 0.931

	pointsOfFunctionPlot[326].X = 3.26
	pointsOfFunctionPlot[326].Y = 0.924

	pointsOfFunctionPlot[327].X = 3.27
	pointsOfFunctionPlot[327].Y = 0.918

	pointsOfFunctionPlot[328].X = 3.28
	pointsOfFunctionPlot[328].Y = 0.912

	pointsOfFunctionPlot[329].X = 3.29
	pointsOfFunctionPlot[329].Y = 0.905

	pointsOfFunctionPlot[330].X = 3.3
	pointsOfFunctionPlot[330].Y = 0.899

	pointsOfFunctionPlot[331].X = 3.31
	pointsOfFunctionPlot[331].Y = 0.893

	pointsOfFunctionPlot[332].X = 3.32
	pointsOfFunctionPlot[332].Y = 0.886

	pointsOfFunctionPlot[333].X = 3.33
	pointsOfFunctionPlot[333].Y = 0.88

	pointsOfFunctionPlot[334].X = 3.34
	pointsOfFunctionPlot[334].Y = 0.874

	pointsOfFunctionPlot[335].X = 3.35
	pointsOfFunctionPlot[335].Y = 0.868

	pointsOfFunctionPlot[336].X = 3.36
	pointsOfFunctionPlot[336].Y = 0.862

	pointsOfFunctionPlot[337].X = 3.37
	pointsOfFunctionPlot[337].Y = 0.856

	pointsOfFunctionPlot[338].X = 3.38
	pointsOfFunctionPlot[338].Y = 0.85

	pointsOfFunctionPlot[339].X = 3.39
	pointsOfFunctionPlot[339].Y = 0.844

	pointsOfFunctionPlot[340].X = 3.4
	pointsOfFunctionPlot[340].Y = 0.838

	pointsOfFunctionPlot[341].X = 3.41
	pointsOfFunctionPlot[341].Y = 0.832

	pointsOfFunctionPlot[342].X = 3.42
	pointsOfFunctionPlot[342].Y = 0.826

	pointsOfFunctionPlot[343].X = 3.43
	pointsOfFunctionPlot[343].Y = 0.82

	pointsOfFunctionPlot[344].X = 3.44
	pointsOfFunctionPlot[344].Y = 0.815

	pointsOfFunctionPlot[345].X = 3.45
	pointsOfFunctionPlot[345].Y = 0.809

	pointsOfFunctionPlot[346].X = 3.46
	pointsOfFunctionPlot[346].Y = 0.803

	pointsOfFunctionPlot[347].X = 3.47
	pointsOfFunctionPlot[347].Y = 0.797

	pointsOfFunctionPlot[348].X = 3.48
	pointsOfFunctionPlot[348].Y = 0.792

	pointsOfFunctionPlot[349].X = 3.49
	pointsOfFunctionPlot[349].Y = 0.786

	pointsOfFunctionPlot[350].X = 3.5
	pointsOfFunctionPlot[350].Y = 0.781

	pointsOfFunctionPlot[351].X = 3.51
	pointsOfFunctionPlot[351].Y = 0.775

	pointsOfFunctionPlot[352].X = 3.52
	pointsOfFunctionPlot[352].Y = 0.769

	pointsOfFunctionPlot[353].X = 3.53
	pointsOfFunctionPlot[353].Y = 0.764

	pointsOfFunctionPlot[354].X = 3.54
	pointsOfFunctionPlot[354].Y = 0.758

	pointsOfFunctionPlot[355].X = 3.55
	pointsOfFunctionPlot[355].Y = 0.753

	pointsOfFunctionPlot[356].X = 3.56
	pointsOfFunctionPlot[356].Y = 0.748

	pointsOfFunctionPlot[357].X = 3.57
	pointsOfFunctionPlot[357].Y = 0.742

	pointsOfFunctionPlot[358].X = 3.58
	pointsOfFunctionPlot[358].Y = 0.737

	pointsOfFunctionPlot[359].X = 3.59
	pointsOfFunctionPlot[359].Y = 0.732

	pointsOfFunctionPlot[360].X = 3.6
	pointsOfFunctionPlot[360].Y = 0.726

	pointsOfFunctionPlot[361].X = 3.61
	pointsOfFunctionPlot[361].Y = 0.721

	pointsOfFunctionPlot[362].X = 3.62
	pointsOfFunctionPlot[362].Y = 0.716

	pointsOfFunctionPlot[363].X = 3.63
	pointsOfFunctionPlot[363].Y = 0.711

	pointsOfFunctionPlot[364].X = 3.64
	pointsOfFunctionPlot[364].Y = 0.706

	pointsOfFunctionPlot[365].X = 3.65
	pointsOfFunctionPlot[365].Y = 0.701

	pointsOfFunctionPlot[366].X = 3.66
	pointsOfFunctionPlot[366].Y = 0.695

	pointsOfFunctionPlot[367].X = 3.67
	pointsOfFunctionPlot[367].Y = 0.69

	pointsOfFunctionPlot[368].X = 3.68
	pointsOfFunctionPlot[368].Y = 0.685

	pointsOfFunctionPlot[369].X = 3.69
	pointsOfFunctionPlot[369].Y = 0.68

	pointsOfFunctionPlot[370].X = 3.7
	pointsOfFunctionPlot[370].Y = 0.675

	pointsOfFunctionPlot[371].X = 3.71
	pointsOfFunctionPlot[371].Y = 0.671

	pointsOfFunctionPlot[372].X = 3.72
	pointsOfFunctionPlot[372].Y = 0.666

	pointsOfFunctionPlot[373].X = 3.73
	pointsOfFunctionPlot[373].Y = 0.661

	pointsOfFunctionPlot[374].X = 3.74
	pointsOfFunctionPlot[374].Y = 0.656

	pointsOfFunctionPlot[375].X = 3.75
	pointsOfFunctionPlot[375].Y = 0.651

	pointsOfFunctionPlot[376].X = 3.76
	pointsOfFunctionPlot[376].Y = 0.646

	pointsOfFunctionPlot[377].X = 3.77
	pointsOfFunctionPlot[377].Y = 0.642

	pointsOfFunctionPlot[378].X = 3.78
	pointsOfFunctionPlot[378].Y = 0.637

	pointsOfFunctionPlot[379].X = 3.79
	pointsOfFunctionPlot[379].Y = 0.632

	pointsOfFunctionPlot[380].X = 3.8
	pointsOfFunctionPlot[380].Y = 0.628

	pointsOfFunctionPlot[381].X = 3.81
	pointsOfFunctionPlot[381].Y = 0.623

	pointsOfFunctionPlot[382].X = 3.82
	pointsOfFunctionPlot[382].Y = 0.618

	pointsOfFunctionPlot[383].X = 3.83
	pointsOfFunctionPlot[383].Y = 0.614

	pointsOfFunctionPlot[384].X = 3.84
	pointsOfFunctionPlot[384].Y = 0.609

	pointsOfFunctionPlot[385].X = 3.85
	pointsOfFunctionPlot[385].Y = 0.605

	pointsOfFunctionPlot[386].X = 3.86
	pointsOfFunctionPlot[386].Y = 0.6

	pointsOfFunctionPlot[387].X = 3.87
	pointsOfFunctionPlot[387].Y = 0.596

	pointsOfFunctionPlot[388].X = 3.88
	pointsOfFunctionPlot[388].Y = 0.592

	pointsOfFunctionPlot[389].X = 3.89
	pointsOfFunctionPlot[389].Y = 0.587

	pointsOfFunctionPlot[390].X = 3.9
	pointsOfFunctionPlot[390].Y = 0.583

	pointsOfFunctionPlot[391].X = 3.91
	pointsOfFunctionPlot[391].Y = 0.579

	pointsOfFunctionPlot[392].X = 3.92
	pointsOfFunctionPlot[392].Y = 0.574

	pointsOfFunctionPlot[393].X = 3.93
	pointsOfFunctionPlot[393].Y = 0.57

	pointsOfFunctionPlot[394].X = 3.94
	pointsOfFunctionPlot[394].Y = 0.566

	pointsOfFunctionPlot[395].X = 3.95
	pointsOfFunctionPlot[395].Y = 0.562

	pointsOfFunctionPlot[396].X = 3.96
	pointsOfFunctionPlot[396].Y = 0.557

	pointsOfFunctionPlot[397].X = 3.97
	pointsOfFunctionPlot[397].Y = 0.553

	pointsOfFunctionPlot[398].X = 3.98
	pointsOfFunctionPlot[398].Y = 0.549

	pointsOfFunctionPlot[399].X = 3.99
	pointsOfFunctionPlot[399].Y = 0.545

	pointsOfFunctionPlot[400].X = 4.0
	pointsOfFunctionPlot[400].Y = 0.541

	pointsOfFunctionPlot[401].X = 4.01
	pointsOfFunctionPlot[401].Y = 0.537

	pointsOfFunctionPlot[402].X = 4.02
	pointsOfFunctionPlot[402].Y = 0.533

	pointsOfFunctionPlot[403].X = 4.03
	pointsOfFunctionPlot[403].Y = 0.529

	pointsOfFunctionPlot[404].X = 4.04
	pointsOfFunctionPlot[404].Y = 0.525

	pointsOfFunctionPlot[405].X = 4.05
	pointsOfFunctionPlot[405].Y = 0.521

	pointsOfFunctionPlot[406].X = 4.06
	pointsOfFunctionPlot[406].Y = 0.517

	pointsOfFunctionPlot[407].X = 4.07
	pointsOfFunctionPlot[407].Y = 0.513

	pointsOfFunctionPlot[408].X = 4.08
	pointsOfFunctionPlot[408].Y = 0.509

	pointsOfFunctionPlot[409].X = 4.09
	pointsOfFunctionPlot[409].Y = 0.505

	pointsOfFunctionPlot[410].X = 4.1
	pointsOfFunctionPlot[410].Y = 0.502

	pointsOfFunctionPlot[411].X = 4.11
	pointsOfFunctionPlot[411].Y = 0.498

	pointsOfFunctionPlot[412].X = 4.12
	pointsOfFunctionPlot[412].Y = 0.494

	pointsOfFunctionPlot[413].X = 4.13
	pointsOfFunctionPlot[413].Y = 0.49

	pointsOfFunctionPlot[414].X = 4.14
	pointsOfFunctionPlot[414].Y = 0.487

	pointsOfFunctionPlot[415].X = 4.15
	pointsOfFunctionPlot[415].Y = 0.483

	pointsOfFunctionPlot[416].X = 4.16
	pointsOfFunctionPlot[416].Y = 0.479

	pointsOfFunctionPlot[417].X = 4.17
	pointsOfFunctionPlot[417].Y = 0.476

	pointsOfFunctionPlot[418].X = 4.18
	pointsOfFunctionPlot[418].Y = 0.472

	pointsOfFunctionPlot[419].X = 4.19
	pointsOfFunctionPlot[419].Y = 0.468

	pointsOfFunctionPlot[420].X = 4.2
	pointsOfFunctionPlot[420].Y = 0.465

	pointsOfFunctionPlot[421].X = 4.21
	pointsOfFunctionPlot[421].Y = 0.461

	pointsOfFunctionPlot[422].X = 4.22
	pointsOfFunctionPlot[422].Y = 0.458

	pointsOfFunctionPlot[423].X = 4.23
	pointsOfFunctionPlot[423].Y = 0.454

	pointsOfFunctionPlot[424].X = 4.24
	pointsOfFunctionPlot[424].Y = 0.451

	pointsOfFunctionPlot[425].X = 4.25
	pointsOfFunctionPlot[425].Y = 0.447

	pointsOfFunctionPlot[426].X = 4.26
	pointsOfFunctionPlot[426].Y = 0.444

	pointsOfFunctionPlot[427].X = 4.27
	pointsOfFunctionPlot[427].Y = 0.441

	pointsOfFunctionPlot[428].X = 4.28
	pointsOfFunctionPlot[428].Y = 0.437

	pointsOfFunctionPlot[429].X = 4.29
	pointsOfFunctionPlot[429].Y = 0.434

	pointsOfFunctionPlot[430].X = 4.3
	pointsOfFunctionPlot[430].Y = 0.431

	pointsOfFunctionPlot[431].X = 4.31
	pointsOfFunctionPlot[431].Y = 0.427

	pointsOfFunctionPlot[432].X = 4.32
	pointsOfFunctionPlot[432].Y = 0.424

	pointsOfFunctionPlot[433].X = 4.33
	pointsOfFunctionPlot[433].Y = 0.421

	pointsOfFunctionPlot[434].X = 4.34
	pointsOfFunctionPlot[434].Y = 0.418

	pointsOfFunctionPlot[435].X = 4.35
	pointsOfFunctionPlot[435].Y = 0.414

	pointsOfFunctionPlot[436].X = 4.36
	pointsOfFunctionPlot[436].Y = 0.411

	pointsOfFunctionPlot[437].X = 4.37
	pointsOfFunctionPlot[437].Y = 0.408

	pointsOfFunctionPlot[438].X = 4.38
	pointsOfFunctionPlot[438].Y = 0.405

	pointsOfFunctionPlot[439].X = 4.39
	pointsOfFunctionPlot[439].Y = 0.402

	pointsOfFunctionPlot[440].X = 4.4
	pointsOfFunctionPlot[440].Y = 0.399

	pointsOfFunctionPlot[441].X = 4.41
	pointsOfFunctionPlot[441].Y = 0.396

	pointsOfFunctionPlot[442].X = 4.42
	pointsOfFunctionPlot[442].Y = 0.393

	pointsOfFunctionPlot[443].X = 4.43
	pointsOfFunctionPlot[443].Y = 0.39

	pointsOfFunctionPlot[444].X = 4.44
	pointsOfFunctionPlot[444].Y = 0.387

	pointsOfFunctionPlot[445].X = 4.45
	pointsOfFunctionPlot[445].Y = 0.384

	pointsOfFunctionPlot[446].X = 4.46
	pointsOfFunctionPlot[446].Y = 0.381

	pointsOfFunctionPlot[447].X = 4.47
	pointsOfFunctionPlot[447].Y = 0.378

	pointsOfFunctionPlot[448].X = 4.48
	pointsOfFunctionPlot[448].Y = 0.375

	pointsOfFunctionPlot[449].X = 4.49
	pointsOfFunctionPlot[449].Y = 0.372

	pointsOfFunctionPlot[450].X = 4.5
	pointsOfFunctionPlot[450].Y = 0.369

	pointsOfFunctionPlot[451].X = 4.51
	pointsOfFunctionPlot[451].Y = 0.366

	pointsOfFunctionPlot[452].X = 4.52
	pointsOfFunctionPlot[452].Y = 0.363

	pointsOfFunctionPlot[453].X = 4.53
	pointsOfFunctionPlot[453].Y = 0.36

	pointsOfFunctionPlot[454].X = 4.54
	pointsOfFunctionPlot[454].Y = 0.358

	pointsOfFunctionPlot[455].X = 4.55
	pointsOfFunctionPlot[455].Y = 0.355

	pointsOfFunctionPlot[456].X = 4.56
	pointsOfFunctionPlot[456].Y = 0.352

	pointsOfFunctionPlot[457].X = 4.57
	pointsOfFunctionPlot[457].Y = 0.349

	pointsOfFunctionPlot[458].X = 4.58
	pointsOfFunctionPlot[458].Y = 0.347

	pointsOfFunctionPlot[459].X = 4.59
	pointsOfFunctionPlot[459].Y = 0.344

	pointsOfFunctionPlot[460].X = 4.6
	pointsOfFunctionPlot[460].Y = 0.341

	pointsOfFunctionPlot[461].X = 4.61
	pointsOfFunctionPlot[461].Y = 0.339

	pointsOfFunctionPlot[462].X = 4.62
	pointsOfFunctionPlot[462].Y = 0.336

	pointsOfFunctionPlot[463].X = 4.63
	pointsOfFunctionPlot[463].Y = 0.333

	pointsOfFunctionPlot[464].X = 4.64
	pointsOfFunctionPlot[464].Y = 0.331

	pointsOfFunctionPlot[465].X = 4.65
	pointsOfFunctionPlot[465].Y = 0.328

	pointsOfFunctionPlot[466].X = 4.66
	pointsOfFunctionPlot[466].Y = 0.326

	pointsOfFunctionPlot[467].X = 4.67
	pointsOfFunctionPlot[467].Y = 0.323

	pointsOfFunctionPlot[468].X = 4.68
	pointsOfFunctionPlot[468].Y = 0.32

	pointsOfFunctionPlot[469].X = 4.69
	pointsOfFunctionPlot[469].Y = 0.318

	pointsOfFunctionPlot[470].X = 4.7
	pointsOfFunctionPlot[470].Y = 0.315

	pointsOfFunctionPlot[471].X = 4.71
	pointsOfFunctionPlot[471].Y = 0.313

	pointsOfFunctionPlot[472].X = 4.72
	pointsOfFunctionPlot[472].Y = 0.31

	pointsOfFunctionPlot[473].X = 4.73
	pointsOfFunctionPlot[473].Y = 0.308

	pointsOfFunctionPlot[474].X = 4.74
	pointsOfFunctionPlot[474].Y = 0.306

	pointsOfFunctionPlot[475].X = 4.75
	pointsOfFunctionPlot[475].Y = 0.303

	pointsOfFunctionPlot[476].X = 4.76
	pointsOfFunctionPlot[476].Y = 0.301

	pointsOfFunctionPlot[477].X = 4.77
	pointsOfFunctionPlot[477].Y = 0.298

	pointsOfFunctionPlot[478].X = 4.78
	pointsOfFunctionPlot[478].Y = 0.296

	pointsOfFunctionPlot[479].X = 4.79
	pointsOfFunctionPlot[479].Y = 0.294

	pointsOfFunctionPlot[480].X = 4.8
	pointsOfFunctionPlot[480].Y = 0.291

	pointsOfFunctionPlot[481].X = 4.81
	pointsOfFunctionPlot[481].Y = 0.289

	pointsOfFunctionPlot[482].X = 4.82
	pointsOfFunctionPlot[482].Y = 0.287

	pointsOfFunctionPlot[483].X = 4.83
	pointsOfFunctionPlot[483].Y = 0.285

	pointsOfFunctionPlot[484].X = 4.84
	pointsOfFunctionPlot[484].Y = 0.282

	pointsOfFunctionPlot[485].X = 4.85
	pointsOfFunctionPlot[485].Y = 0.28

	pointsOfFunctionPlot[486].X = 4.86
	pointsOfFunctionPlot[486].Y = 0.278

	pointsOfFunctionPlot[487].X = 4.87
	pointsOfFunctionPlot[487].Y = 0.276

	pointsOfFunctionPlot[488].X = 4.88
	pointsOfFunctionPlot[488].Y = 0.273

	pointsOfFunctionPlot[489].X = 4.89
	pointsOfFunctionPlot[489].Y = 0.271

	pointsOfFunctionPlot[490].X = 4.9
	pointsOfFunctionPlot[490].Y = 0.269

	pointsOfFunctionPlot[491].X = 4.91
	pointsOfFunctionPlot[491].Y = 0.267

	pointsOfFunctionPlot[492].X = 4.92
	pointsOfFunctionPlot[492].Y = 0.265

	pointsOfFunctionPlot[493].X = 4.93
	pointsOfFunctionPlot[493].Y = 0.263

	pointsOfFunctionPlot[494].X = 4.94
	pointsOfFunctionPlot[494].Y = 0.261

	pointsOfFunctionPlot[495].X = 4.95
	pointsOfFunctionPlot[495].Y = 0.259

	pointsOfFunctionPlot[496].X = 4.96
	pointsOfFunctionPlot[496].Y = 0.257

	pointsOfFunctionPlot[497].X = 4.97
	pointsOfFunctionPlot[497].Y = 0.255

	pointsOfFunctionPlot[498].X = 4.98
	pointsOfFunctionPlot[498].Y = 0.252

	pointsOfFunctionPlot[499].X = 4.99
	pointsOfFunctionPlot[499].Y = 0.25

	pointsOfFunctionPlot[500].X = 5.0
	pointsOfFunctionPlot[500].Y = 0.248

	pointsOfFunctionPlot[501].X = 5.01
	pointsOfFunctionPlot[501].Y = 0.247

	pointsOfFunctionPlot[502].X = 5.02
	pointsOfFunctionPlot[502].Y = 0.245

	pointsOfFunctionPlot[503].X = 5.03
	pointsOfFunctionPlot[503].Y = 0.243

	pointsOfFunctionPlot[504].X = 5.04
	pointsOfFunctionPlot[504].Y = 0.241

	pointsOfFunctionPlot[505].X = 5.05
	pointsOfFunctionPlot[505].Y = 0.239

	pointsOfFunctionPlot[506].X = 5.06
	pointsOfFunctionPlot[506].Y = 0.237

	pointsOfFunctionPlot[507].X = 5.07
	pointsOfFunctionPlot[507].Y = 0.235

	pointsOfFunctionPlot[508].X = 5.08
	pointsOfFunctionPlot[508].Y = 0.233

	pointsOfFunctionPlot[509].X = 5.09
	pointsOfFunctionPlot[509].Y = 0.231

	pointsOfFunctionPlot[510].X = 5.1
	pointsOfFunctionPlot[510].Y = 0.229

	pointsOfFunctionPlot[511].X = 5.11
	pointsOfFunctionPlot[511].Y = 0.227

	pointsOfFunctionPlot[512].X = 5.12
	pointsOfFunctionPlot[512].Y = 0.226

	pointsOfFunctionPlot[513].X = 5.13
	pointsOfFunctionPlot[513].Y = 0.224

	pointsOfFunctionPlot[514].X = 5.14
	pointsOfFunctionPlot[514].Y = 0.222

	pointsOfFunctionPlot[515].X = 5.15
	pointsOfFunctionPlot[515].Y = 0.22

	pointsOfFunctionPlot[516].X = 5.16
	pointsOfFunctionPlot[516].Y = 0.218

	pointsOfFunctionPlot[517].X = 5.17
	pointsOfFunctionPlot[517].Y = 0.217

	pointsOfFunctionPlot[518].X = 5.18
	pointsOfFunctionPlot[518].Y = 0.215

	pointsOfFunctionPlot[519].X = 5.19
	pointsOfFunctionPlot[519].Y = 0.213

	pointsOfFunctionPlot[520].X = 5.2
	pointsOfFunctionPlot[520].Y = 0.212

	pointsOfFunctionPlot[521].X = 5.21
	pointsOfFunctionPlot[521].Y = 0.21

	pointsOfFunctionPlot[522].X = 5.22
	pointsOfFunctionPlot[522].Y = 0.208

	pointsOfFunctionPlot[523].X = 5.23
	pointsOfFunctionPlot[523].Y = 0.206

	pointsOfFunctionPlot[524].X = 5.24
	pointsOfFunctionPlot[524].Y = 0.205

	pointsOfFunctionPlot[525].X = 5.25
	pointsOfFunctionPlot[525].Y = 0.203

	pointsOfFunctionPlot[526].X = 5.26
	pointsOfFunctionPlot[526].Y = 0.201

	pointsOfFunctionPlot[527].X = 5.27
	pointsOfFunctionPlot[527].Y = 0.2

	pointsOfFunctionPlot[528].X = 5.28
	pointsOfFunctionPlot[528].Y = 0.198

	pointsOfFunctionPlot[529].X = 5.29
	pointsOfFunctionPlot[529].Y = 0.197

	pointsOfFunctionPlot[530].X = 5.3
	pointsOfFunctionPlot[530].Y = 0.195

	pointsOfFunctionPlot[531].X = 5.31
	pointsOfFunctionPlot[531].Y = 0.193

	pointsOfFunctionPlot[532].X = 5.32
	pointsOfFunctionPlot[532].Y = 0.192

	pointsOfFunctionPlot[533].X = 5.33
	pointsOfFunctionPlot[533].Y = 0.19

	pointsOfFunctionPlot[534].X = 5.34
	pointsOfFunctionPlot[534].Y = 0.189

	pointsOfFunctionPlot[535].X = 5.35
	pointsOfFunctionPlot[535].Y = 0.187

	pointsOfFunctionPlot[536].X = 5.36
	pointsOfFunctionPlot[536].Y = 0.186

	pointsOfFunctionPlot[537].X = 5.37
	pointsOfFunctionPlot[537].Y = 0.184

	pointsOfFunctionPlot[538].X = 5.38
	pointsOfFunctionPlot[538].Y = 0.183

	pointsOfFunctionPlot[539].X = 5.39
	pointsOfFunctionPlot[539].Y = 0.181

	pointsOfFunctionPlot[540].X = 5.4
	pointsOfFunctionPlot[540].Y = 0.18

	pointsOfFunctionPlot[541].X = 5.41
	pointsOfFunctionPlot[541].Y = 0.178

	pointsOfFunctionPlot[542].X = 5.42
	pointsOfFunctionPlot[542].Y = 0.177

	pointsOfFunctionPlot[543].X = 5.43
	pointsOfFunctionPlot[543].Y = 0.175

	pointsOfFunctionPlot[544].X = 5.44
	pointsOfFunctionPlot[544].Y = 0.174

	pointsOfFunctionPlot[545].X = 5.45
	pointsOfFunctionPlot[545].Y = 0.173

	pointsOfFunctionPlot[546].X = 5.46
	pointsOfFunctionPlot[546].Y = 0.171

	pointsOfFunctionPlot[547].X = 5.47
	pointsOfFunctionPlot[547].Y = 0.17

	pointsOfFunctionPlot[548].X = 5.48
	pointsOfFunctionPlot[548].Y = 0.168

	pointsOfFunctionPlot[549].X = 5.49
	pointsOfFunctionPlot[549].Y = 0.167

	pointsOfFunctionPlot[550].X = 5.5
	pointsOfFunctionPlot[550].Y = 0.166

	pointsOfFunctionPlot[551].X = 5.51
	pointsOfFunctionPlot[551].Y = 0.164

	pointsOfFunctionPlot[552].X = 5.52
	pointsOfFunctionPlot[552].Y = 0.163

	pointsOfFunctionPlot[553].X = 5.53
	pointsOfFunctionPlot[553].Y = 0.162

	pointsOfFunctionPlot[554].X = 5.54
	pointsOfFunctionPlot[554].Y = 0.16

	pointsOfFunctionPlot[555].X = 5.55
	pointsOfFunctionPlot[555].Y = 0.159

	pointsOfFunctionPlot[556].X = 5.56
	pointsOfFunctionPlot[556].Y = 0.158

	pointsOfFunctionPlot[557].X = 5.57
	pointsOfFunctionPlot[557].Y = 0.156

	pointsOfFunctionPlot[558].X = 5.58
	pointsOfFunctionPlot[558].Y = 0.155

	pointsOfFunctionPlot[559].X = 5.59
	pointsOfFunctionPlot[559].Y = 0.154

	pointsOfFunctionPlot[560].X = 5.6
	pointsOfFunctionPlot[560].Y = 0.153

	pointsOfFunctionPlot[561].X = 5.61
	pointsOfFunctionPlot[561].Y = 0.151

	pointsOfFunctionPlot[562].X = 5.62
	pointsOfFunctionPlot[562].Y = 0.15

	pointsOfFunctionPlot[563].X = 5.63
	pointsOfFunctionPlot[563].Y = 0.149

	pointsOfFunctionPlot[564].X = 5.64
	pointsOfFunctionPlot[564].Y = 0.148

	pointsOfFunctionPlot[565].X = 5.65
	pointsOfFunctionPlot[565].Y = 0.146

	pointsOfFunctionPlot[566].X = 5.66
	pointsOfFunctionPlot[566].Y = 0.145

	pointsOfFunctionPlot[567].X = 5.67
	pointsOfFunctionPlot[567].Y = 0.144

	pointsOfFunctionPlot[568].X = 5.68
	pointsOfFunctionPlot[568].Y = 0.143

	pointsOfFunctionPlot[569].X = 5.69
	pointsOfFunctionPlot[569].Y = 0.142

	pointsOfFunctionPlot[570].X = 5.7
	pointsOfFunctionPlot[570].Y = 0.14

	pointsOfFunctionPlot[571].X = 5.71
	pointsOfFunctionPlot[571].Y = 0.139

	pointsOfFunctionPlot[572].X = 5.72
	pointsOfFunctionPlot[572].Y = 0.138

	pointsOfFunctionPlot[573].X = 5.73
	pointsOfFunctionPlot[573].Y = 0.137

	pointsOfFunctionPlot[574].X = 5.74
	pointsOfFunctionPlot[574].Y = 0.136

	pointsOfFunctionPlot[575].X = 5.75
	pointsOfFunctionPlot[575].Y = 0.135

	pointsOfFunctionPlot[576].X = 5.76
	pointsOfFunctionPlot[576].Y = 0.134

	pointsOfFunctionPlot[577].X = 5.77
	pointsOfFunctionPlot[577].Y = 0.133

	pointsOfFunctionPlot[578].X = 5.78
	pointsOfFunctionPlot[578].Y = 0.131

	pointsOfFunctionPlot[579].X = 5.79
	pointsOfFunctionPlot[579].Y = 0.13

	pointsOfFunctionPlot[580].X = 5.8
	pointsOfFunctionPlot[580].Y = 0.129

	pointsOfFunctionPlot[581].X = 5.81
	pointsOfFunctionPlot[581].Y = 0.128

	pointsOfFunctionPlot[582].X = 5.82
	pointsOfFunctionPlot[582].Y = 0.127

	pointsOfFunctionPlot[583].X = 5.83
	pointsOfFunctionPlot[583].Y = 0.126

	pointsOfFunctionPlot[584].X = 5.84
	pointsOfFunctionPlot[584].Y = 0.125

	pointsOfFunctionPlot[585].X = 5.85
	pointsOfFunctionPlot[585].Y = 0.124

	pointsOfFunctionPlot[586].X = 5.86
	pointsOfFunctionPlot[586].Y = 0.123

	pointsOfFunctionPlot[587].X = 5.87
	pointsOfFunctionPlot[587].Y = 0.122

	pointsOfFunctionPlot[588].X = 5.88
	pointsOfFunctionPlot[588].Y = 0.121

	pointsOfFunctionPlot[589].X = 5.89
	pointsOfFunctionPlot[589].Y = 0.12

	pointsOfFunctionPlot[590].X = 5.9
	pointsOfFunctionPlot[590].Y = 0.119

	pointsOfFunctionPlot[591].X = 5.91
	pointsOfFunctionPlot[591].Y = 0.118

	pointsOfFunctionPlot[592].X = 5.92
	pointsOfFunctionPlot[592].Y = 0.117

	pointsOfFunctionPlot[593].X = 5.93
	pointsOfFunctionPlot[593].Y = 0.116

	pointsOfFunctionPlot[594].X = 5.94
	pointsOfFunctionPlot[594].Y = 0.115

	pointsOfFunctionPlot[595].X = 5.95
	pointsOfFunctionPlot[595].Y = 0.114

	pointsOfFunctionPlot[596].X = 5.96
	pointsOfFunctionPlot[596].Y = 0.113

	pointsOfFunctionPlot[597].X = 5.97
	pointsOfFunctionPlot[597].Y = 0.112

	pointsOfFunctionPlot[598].X = 5.98
	pointsOfFunctionPlot[598].Y = 0.111

	pointsOfFunctionPlot[599].X = 5.99
	pointsOfFunctionPlot[599].Y = 0.11

	pointsOfFunctionPlot[600].X = 6.0
	pointsOfFunctionPlot[600].Y = 0.109

	pointsOfFunctionPlot[601].X = 6.01
	pointsOfFunctionPlot[601].Y = 0.109

	pointsOfFunctionPlot[602].X = 6.02
	pointsOfFunctionPlot[602].Y = 0.108

	pointsOfFunctionPlot[603].X = 6.03
	pointsOfFunctionPlot[603].Y = 0.107

	pointsOfFunctionPlot[604].X = 6.04
	pointsOfFunctionPlot[604].Y = 0.106

	pointsOfFunctionPlot[605].X = 6.05
	pointsOfFunctionPlot[605].Y = 0.105

	pointsOfFunctionPlot[606].X = 6.06
	pointsOfFunctionPlot[606].Y = 0.104

	pointsOfFunctionPlot[607].X = 6.07
	pointsOfFunctionPlot[607].Y = 0.103

	pointsOfFunctionPlot[608].X = 6.08
	pointsOfFunctionPlot[608].Y = 0.102

	pointsOfFunctionPlot[609].X = 6.09
	pointsOfFunctionPlot[609].Y = 0.101

	pointsOfFunctionPlot[610].X = 6.10
	pointsOfFunctionPlot[610].Y = 0.101

	pointsOfFunctionPlot[611].X = 6.11
	pointsOfFunctionPlot[611].Y = 0.1

	pointsOfFunctionPlot[612].X = 6.12
	pointsOfFunctionPlot[612].Y = 0.099

	pointsOfFunctionPlot[613].X = 6.13
	pointsOfFunctionPlot[613].Y = 0.098

	pointsOfFunctionPlot[614].X = 6.14
	pointsOfFunctionPlot[614].Y = 0.097

	pointsOfFunctionPlot[615].X = 6.15
	pointsOfFunctionPlot[615].Y = 0.097

	pointsOfFunctionPlot[616].X = 6.16
	pointsOfFunctionPlot[616].Y = 0.096

	pointsOfFunctionPlot[617].X = 6.17
	pointsOfFunctionPlot[617].Y = 0.095

	pointsOfFunctionPlot[618].X = 6.18
	pointsOfFunctionPlot[618].Y = 0.094

	pointsOfFunctionPlot[619].X = 6.19
	pointsOfFunctionPlot[619].Y = 0.093

	pointsOfFunctionPlot[620].X = 6.2
	pointsOfFunctionPlot[620].Y = 0.093

	pointsOfFunctionPlot[621].X = 6.21
	pointsOfFunctionPlot[621].Y = 0.092

	pointsOfFunctionPlot[622].X = 6.22
	pointsOfFunctionPlot[622].Y = 0.091

	pointsOfFunctionPlot[623].X = 6.23
	pointsOfFunctionPlot[623].Y = 0.09

	pointsOfFunctionPlot[624].X = 6.24
	pointsOfFunctionPlot[624].Y = 0.089

	pointsOfFunctionPlot[625].X = 6.25
	pointsOfFunctionPlot[625].Y = 0.089

	pointsOfFunctionPlot[626].X = 6.26
	pointsOfFunctionPlot[626].Y = 0.088

	pointsOfFunctionPlot[627].X = 6.27
	pointsOfFunctionPlot[627].Y = 0.087

	pointsOfFunctionPlot[628].X = 6.28
	pointsOfFunctionPlot[628].Y = 0.086

	pointsOfFunctionPlot[629].X = 6.29
	pointsOfFunctionPlot[629].Y = 0.086

	pointsOfFunctionPlot[630].X = 6.3
	pointsOfFunctionPlot[630].Y = 0.085

	pointsOfFunctionPlot[631].X = 6.31
	pointsOfFunctionPlot[631].Y = 0.084

	pointsOfFunctionPlot[632].X = 6.32
	pointsOfFunctionPlot[632].Y = 0.084

	pointsOfFunctionPlot[633].X = 6.33
	pointsOfFunctionPlot[633].Y = 0.083

	pointsOfFunctionPlot[634].X = 6.34
	pointsOfFunctionPlot[634].Y = 0.082

	pointsOfFunctionPlot[635].X = 6.35
	pointsOfFunctionPlot[635].Y = 0.082

	pointsOfFunctionPlot[636].X = 6.36
	pointsOfFunctionPlot[636].Y = 0.081

	pointsOfFunctionPlot[637].X = 6.37
	pointsOfFunctionPlot[637].Y = 0.08

	pointsOfFunctionPlot[638].X = 6.38
	pointsOfFunctionPlot[638].Y = 0.079

	pointsOfFunctionPlot[639].X = 6.39
	pointsOfFunctionPlot[639].Y = 0.079

	pointsOfFunctionPlot[640].X = 6.4
	pointsOfFunctionPlot[640].Y = 0.078

	pointsOfFunctionPlot[641].X = 6.41
	pointsOfFunctionPlot[641].Y = 0.077

	pointsOfFunctionPlot[642].X = 6.42
	pointsOfFunctionPlot[642].Y = 0.077

	pointsOfFunctionPlot[643].X = 6.43
	pointsOfFunctionPlot[643].Y = 0.076

	pointsOfFunctionPlot[644].X = 6.44
	pointsOfFunctionPlot[644].Y = 0.076

	pointsOfFunctionPlot[645].X = 6.45
	pointsOfFunctionPlot[645].Y = 0.075

	pointsOfFunctionPlot[646].X = 6.46
	pointsOfFunctionPlot[646].Y = 0.074

	pointsOfFunctionPlot[647].X = 6.47
	pointsOfFunctionPlot[647].Y = 0.074

	pointsOfFunctionPlot[648].X = 6.48
	pointsOfFunctionPlot[648].Y = 0.073

	pointsOfFunctionPlot[649].X = 6.49
	pointsOfFunctionPlot[649].Y = 0.072

	pointsOfFunctionPlot[650].X = 6.5
	pointsOfFunctionPlot[650].Y = 0.072

	pointsOfFunctionPlot[651].X = 6.51
	pointsOfFunctionPlot[651].Y = 0.071

	pointsOfFunctionPlot[652].X = 6.52
	pointsOfFunctionPlot[652].Y = 0.071

	pointsOfFunctionPlot[653].X = 6.53
	pointsOfFunctionPlot[653].Y = 0.07

	pointsOfFunctionPlot[654].X = 6.54
	pointsOfFunctionPlot[654].Y = 0.069

	pointsOfFunctionPlot[655].X = 6.55
	pointsOfFunctionPlot[655].Y = 0.069

	pointsOfFunctionPlot[656].X = 6.56
	pointsOfFunctionPlot[656].Y = 0.068

	pointsOfFunctionPlot[657].X = 6.57
	pointsOfFunctionPlot[657].Y = 0.068

	pointsOfFunctionPlot[658].X = 6.58
	pointsOfFunctionPlot[658].Y = 0.067

	pointsOfFunctionPlot[659].X = 6.59
	pointsOfFunctionPlot[659].Y = 0.066

	pointsOfFunctionPlot[660].X = 6.6
	pointsOfFunctionPlot[660].Y = 0.066

	pointsOfFunctionPlot[661].X = 6.61
	pointsOfFunctionPlot[661].Y = 0.065

	pointsOfFunctionPlot[662].X = 6.62
	pointsOfFunctionPlot[662].Y = 0.065

	pointsOfFunctionPlot[663].X = 6.63
	pointsOfFunctionPlot[663].Y = 0.064

	pointsOfFunctionPlot[664].X = 6.64
	pointsOfFunctionPlot[664].Y = 0.064

	pointsOfFunctionPlot[665].X = 6.65
	pointsOfFunctionPlot[665].Y = 0.063

	pointsOfFunctionPlot[666].X = 6.66
	pointsOfFunctionPlot[666].Y = 0.063

	pointsOfFunctionPlot[667].X = 6.67
	pointsOfFunctionPlot[667].Y = 0.062

	pointsOfFunctionPlot[668].X = 6.68
	pointsOfFunctionPlot[668].Y = 0.062

	pointsOfFunctionPlot[669].X = 6.69
	pointsOfFunctionPlot[669].Y = 0.061

	pointsOfFunctionPlot[670].X = 6.7
	pointsOfFunctionPlot[670].Y = 0.06

	pointsOfFunctionPlot[671].X = 6.71
	pointsOfFunctionPlot[671].Y = 0.06

	pointsOfFunctionPlot[672].X = 6.72
	pointsOfFunctionPlot[672].Y = 0.059

	pointsOfFunctionPlot[673].X = 6.73
	pointsOfFunctionPlot[673].Y = 0.059

	pointsOfFunctionPlot[674].X = 6.74
	pointsOfFunctionPlot[674].Y = 0.058

	pointsOfFunctionPlot[675].X = 6.75
	pointsOfFunctionPlot[675].Y = 0.058

	pointsOfFunctionPlot[676].X = 6.76
	pointsOfFunctionPlot[676].Y = 0.057

	pointsOfFunctionPlot[677].X = 6.77
	pointsOfFunctionPlot[677].Y = 0.057

	pointsOfFunctionPlot[678].X = 6.78
	pointsOfFunctionPlot[678].Y = 0.056

	pointsOfFunctionPlot[679].X = 6.79
	pointsOfFunctionPlot[679].Y = 0.056

	pointsOfFunctionPlot[680].X = 6.8
	pointsOfFunctionPlot[680].Y = 0.056

	pointsOfFunctionPlot[681].X = 6.81
	pointsOfFunctionPlot[681].Y = 0.055

	pointsOfFunctionPlot[682].X = 6.82
	pointsOfFunctionPlot[682].Y = 0.055

	pointsOfFunctionPlot[683].X = 6.83
	pointsOfFunctionPlot[683].Y = 0.054

	pointsOfFunctionPlot[684].X = 6.84
	pointsOfFunctionPlot[684].Y = 0.054

	pointsOfFunctionPlot[685].X = 6.85
	pointsOfFunctionPlot[685].Y = 0.053

	pointsOfFunctionPlot[686].X = 6.86
	pointsOfFunctionPlot[686].Y = 0.053

	pointsOfFunctionPlot[687].X = 6.87
	pointsOfFunctionPlot[687].Y = 0.052

	pointsOfFunctionPlot[688].X = 6.88
	pointsOfFunctionPlot[688].Y = 0.052

	pointsOfFunctionPlot[689].X = 6.89
	pointsOfFunctionPlot[689].Y = 0.051

	pointsOfFunctionPlot[690].X = 6.9
	pointsOfFunctionPlot[690].Y = 0.051

	pointsOfFunctionPlot[691].X = 6.91
	pointsOfFunctionPlot[691].Y = 0.05

	pointsOfFunctionPlot[692].X = 6.92
	pointsOfFunctionPlot[692].Y = 0.05

	pointsOfFunctionPlot[693].X = 6.93
	pointsOfFunctionPlot[693].Y = 0.05

	pointsOfFunctionPlot[694].X = 6.94
	pointsOfFunctionPlot[694].Y = 0.049

	pointsOfFunctionPlot[695].X = 6.95
	pointsOfFunctionPlot[695].Y = 0.049

	pointsOfFunctionPlot[696].X = 6.96
	pointsOfFunctionPlot[696].Y = 0.048

	pointsOfFunctionPlot[697].X = 6.97
	pointsOfFunctionPlot[697].Y = 0.048

	pointsOfFunctionPlot[698].X = 6.98
	pointsOfFunctionPlot[698].Y = 0.048

	pointsOfFunctionPlot[699].X = 6.99
	pointsOfFunctionPlot[699].Y = 0.047

	pointsOfFunctionPlot[700].X = 7.0
	pointsOfFunctionPlot[700].Y = 0.047

	pointsOfFunctionPlot[701].X = 7.01
	pointsOfFunctionPlot[701].Y = 0.046

	pointsOfFunctionPlot[702].X = 7.02
	pointsOfFunctionPlot[702].Y = 0.046

	pointsOfFunctionPlot[703].X = 7.03
	pointsOfFunctionPlot[703].Y = 0.046

	pointsOfFunctionPlot[704].X = 7.04
	pointsOfFunctionPlot[704].Y = 0.045

	pointsOfFunctionPlot[705].X = 7.05
	pointsOfFunctionPlot[705].Y = 0.045

	pointsOfFunctionPlot[706].X = 7.06
	pointsOfFunctionPlot[706].Y = 0.044

	pointsOfFunctionPlot[707].X = 7.07
	pointsOfFunctionPlot[707].Y = 0.044

	pointsOfFunctionPlot[708].X = 7.08
	pointsOfFunctionPlot[708].Y = 0.044

	pointsOfFunctionPlot[709].X = 7.09
	pointsOfFunctionPlot[709].Y = 0.043

	pointsOfFunctionPlot[710].X = 7.1
	pointsOfFunctionPlot[710].Y = 0.043

	pointsOfFunctionPlot[711].X = 7.11
	pointsOfFunctionPlot[711].Y = 0.042

	pointsOfFunctionPlot[712].X = 7.12
	pointsOfFunctionPlot[712].Y = 0.042

	pointsOfFunctionPlot[713].X = 7.13
	pointsOfFunctionPlot[713].Y = 0.042

	pointsOfFunctionPlot[714].X = 7.14
	pointsOfFunctionPlot[714].Y = 0.041

	pointsOfFunctionPlot[715].X = 7.15
	pointsOfFunctionPlot[715].Y = 0.041

	pointsOfFunctionPlot[716].X = 7.16
	pointsOfFunctionPlot[716].Y = 0.041

	pointsOfFunctionPlot[717].X = 7.17
	pointsOfFunctionPlot[717].Y = 0.04

	pointsOfFunctionPlot[718].X = 7.18
	pointsOfFunctionPlot[718].Y = 0.04

	pointsOfFunctionPlot[719].X = 7.19
	pointsOfFunctionPlot[719].Y = 0.04

	pointsOfFunctionPlot[720].X = 7.2
	pointsOfFunctionPlot[720].Y = 0.039

	pointsOfFunctionPlot[721].X = 7.21
	pointsOfFunctionPlot[721].Y = 0.039

	pointsOfFunctionPlot[722].X = 7.22
	pointsOfFunctionPlot[722].Y = 0.039

	pointsOfFunctionPlot[723].X = 7.23
	pointsOfFunctionPlot[723].Y = 0.038

	pointsOfFunctionPlot[724].X = 7.24
	pointsOfFunctionPlot[724].Y = 0.038

	pointsOfFunctionPlot[725].X = 7.25
	pointsOfFunctionPlot[725].Y = 0.038

	pointsOfFunctionPlot[726].X = 7.26
	pointsOfFunctionPlot[726].Y = 0.037

	pointsOfFunctionPlot[727].X = 7.27
	pointsOfFunctionPlot[727].Y = 0.037

	pointsOfFunctionPlot[728].X = 7.28
	pointsOfFunctionPlot[728].Y = 0.037

	pointsOfFunctionPlot[729].X = 7.29
	pointsOfFunctionPlot[729].Y = 0.036

	pointsOfFunctionPlot[730].X = 7.3
	pointsOfFunctionPlot[730].Y = 0.036

	pointsOfFunctionPlot[731].X = 7.31
	pointsOfFunctionPlot[731].Y = 0.036

	pointsOfFunctionPlot[732].X = 7.32
	pointsOfFunctionPlot[732].Y = 0.035

	pointsOfFunctionPlot[733].X = 7.33
	pointsOfFunctionPlot[733].Y = 0.035

	pointsOfFunctionPlot[734].X = 7.34
	pointsOfFunctionPlot[734].Y = 0.035

	pointsOfFunctionPlot[735].X = 7.35
	pointsOfFunctionPlot[735].Y = 0.034

	pointsOfFunctionPlot[736].X = 7.36
	pointsOfFunctionPlot[736].Y = 0.034

	pointsOfFunctionPlot[737].X = 7.37
	pointsOfFunctionPlot[737].Y = 0.034

	pointsOfFunctionPlot[738].X = 7.38
	pointsOfFunctionPlot[738].Y = 0.034

	pointsOfFunctionPlot[739].X = 7.39
	pointsOfFunctionPlot[739].Y = 0.033

	pointsOfFunctionPlot[740].X = 7.4
	pointsOfFunctionPlot[740].Y = 0.033

	pointsOfFunctionPlot[741].X = 7.41
	pointsOfFunctionPlot[741].Y = 0.033

	pointsOfFunctionPlot[742].X = 7.42
	pointsOfFunctionPlot[742].Y = 0.032

	pointsOfFunctionPlot[743].X = 7.43
	pointsOfFunctionPlot[743].Y = 0.032

	pointsOfFunctionPlot[744].X = 7.44
	pointsOfFunctionPlot[744].Y = 0.032

	pointsOfFunctionPlot[745].X = 7.45
	pointsOfFunctionPlot[745].Y = 0.032

	pointsOfFunctionPlot[746].X = 7.46
	pointsOfFunctionPlot[746].Y = 0.031

	pointsOfFunctionPlot[747].X = 7.47
	pointsOfFunctionPlot[747].Y = 0.031

	pointsOfFunctionPlot[748].X = 7.48
	pointsOfFunctionPlot[748].Y = 0.031

	pointsOfFunctionPlot[749].X = 7.49
	pointsOfFunctionPlot[749].Y = 0.03

	pointsOfFunctionPlot[750].X = 7.5
	pointsOfFunctionPlot[750].Y = 0.03

	pointsOfFunctionPlot[751].X = 7.51
	pointsOfFunctionPlot[751].Y = 0.03

	pointsOfFunctionPlot[752].X = 7.52
	pointsOfFunctionPlot[752].Y = 0.03

	pointsOfFunctionPlot[753].X = 7.53
	pointsOfFunctionPlot[753].Y = 0.029

	pointsOfFunctionPlot[754].X = 7.54
	pointsOfFunctionPlot[754].Y = 0.029

	pointsOfFunctionPlot[755].X = 7.55
	pointsOfFunctionPlot[755].Y = 0.029

	pointsOfFunctionPlot[756].X = 7.56
	pointsOfFunctionPlot[756].Y = 0.029

	pointsOfFunctionPlot[757].X = 7.57
	pointsOfFunctionPlot[757].Y = 0.028

	pointsOfFunctionPlot[758].X = 7.58
	pointsOfFunctionPlot[758].Y = 0.028

	pointsOfFunctionPlot[759].X = 7.59
	pointsOfFunctionPlot[759].Y = 0.028

	pointsOfFunctionPlot[760].X = 7.6
	pointsOfFunctionPlot[760].Y = 0.028

	pointsOfFunctionPlot[761].X = 7.61
	pointsOfFunctionPlot[761].Y = 0.027

	pointsOfFunctionPlot[762].X = 7.62
	pointsOfFunctionPlot[762].Y = 0.027

	pointsOfFunctionPlot[763].X = 7.63
	pointsOfFunctionPlot[763].Y = 0.027

	pointsOfFunctionPlot[764].X = 7.64
	pointsOfFunctionPlot[764].Y = 0.027

	pointsOfFunctionPlot[765].X = 7.65
	pointsOfFunctionPlot[765].Y = 0.026

	pointsOfFunctionPlot[766].X = 7.66
	pointsOfFunctionPlot[766].Y = 0.026

	pointsOfFunctionPlot[767].X = 7.67
	pointsOfFunctionPlot[767].Y = 0.026

	pointsOfFunctionPlot[768].X = 7.68
	pointsOfFunctionPlot[768].Y = 0.026

	pointsOfFunctionPlot[769].X = 7.69
	pointsOfFunctionPlot[769].Y = 0.026

	pointsOfFunctionPlot[770].X = 7.7
	pointsOfFunctionPlot[770].Y = 0.025

	pointsOfFunctionPlot[771].X = 7.71
	pointsOfFunctionPlot[771].Y = 0.025

	pointsOfFunctionPlot[772].X = 7.72
	pointsOfFunctionPlot[772].Y = 0.025

	pointsOfFunctionPlot[773].X = 7.73
	pointsOfFunctionPlot[773].Y = 0.025

	pointsOfFunctionPlot[774].X = 7.74
	pointsOfFunctionPlot[774].Y = 0.024

	pointsOfFunctionPlot[775].X = 7.75
	pointsOfFunctionPlot[775].Y = 0.024

	pointsOfFunctionPlot[776].X = 7.76
	pointsOfFunctionPlot[776].Y = 0.024

	pointsOfFunctionPlot[777].X = 7.77
	pointsOfFunctionPlot[777].Y = 0.024

	pointsOfFunctionPlot[778].X = 7.78
	pointsOfFunctionPlot[778].Y = 0.024

	pointsOfFunctionPlot[779].X = 7.79
	pointsOfFunctionPlot[779].Y = 0.023

	pointsOfFunctionPlot[780].X = 7.8
	pointsOfFunctionPlot[780].Y = 0.023

	pointsOfFunctionPlot[781].X = 7.81
	pointsOfFunctionPlot[781].Y = 0.023

	pointsOfFunctionPlot[782].X = 7.82
	pointsOfFunctionPlot[782].Y = 0.023

	pointsOfFunctionPlot[783].X = 7.83
	pointsOfFunctionPlot[783].Y = 0.023

	pointsOfFunctionPlot[784].X = 7.84
	pointsOfFunctionPlot[784].Y = 0.022

	pointsOfFunctionPlot[785].X = 7.85
	pointsOfFunctionPlot[785].Y = 0.022

	pointsOfFunctionPlot[786].X = 7.86
	pointsOfFunctionPlot[786].Y = 0.022

	pointsOfFunctionPlot[787].X = 7.87
	pointsOfFunctionPlot[787].Y = 0.022

	pointsOfFunctionPlot[788].X = 7.88
	pointsOfFunctionPlot[788].Y = 0.022

	pointsOfFunctionPlot[789].X = 7.89
	pointsOfFunctionPlot[789].Y = 0.021

	pointsOfFunctionPlot[790].X = 7.9
	pointsOfFunctionPlot[790].Y = 0.021

	pointsOfFunctionPlot[791].X = 7.91
	pointsOfFunctionPlot[791].Y = 0.021

	pointsOfFunctionPlot[792].X = 7.92
	pointsOfFunctionPlot[792].Y = 0.021

	pointsOfFunctionPlot[793].X = 7.93
	pointsOfFunctionPlot[793].Y = 0.021

	pointsOfFunctionPlot[794].X = 7.94
	pointsOfFunctionPlot[794].Y = 0.02

	pointsOfFunctionPlot[795].X = 7.95
	pointsOfFunctionPlot[795].Y = 0.02

	pointsOfFunctionPlot[796].X = 7.96
	pointsOfFunctionPlot[796].Y = 0.02

	pointsOfFunctionPlot[797].X = 7.97
	pointsOfFunctionPlot[797].Y = 0.02

	pointsOfFunctionPlot[798].X = 7.98
	pointsOfFunctionPlot[798].Y = 0.02

	pointsOfFunctionPlot[799].X = 7.99
	pointsOfFunctionPlot[799].Y = 0.02

	pointsOfFunctionPlot[800].X = 8.0
	pointsOfFunctionPlot[800].Y = 0.019

	pointsOfFunctionPlot[801].X = 8.01
	pointsOfFunctionPlot[801].Y = 0.019

	pointsOfFunctionPlot[802].X = 8.02
	pointsOfFunctionPlot[802].Y = 0.019

	pointsOfFunctionPlot[803].X = 8.03
	pointsOfFunctionPlot[803].Y = 0.019

	pointsOfFunctionPlot[804].X = 8.04
	pointsOfFunctionPlot[804].Y = 0.019

	pointsOfFunctionPlot[805].X = 8.05
	pointsOfFunctionPlot[805].Y = 0.019

	pointsOfFunctionPlot[806].X = 8.06
	pointsOfFunctionPlot[806].Y = 0.018

	pointsOfFunctionPlot[807].X = 8.07
	pointsOfFunctionPlot[807].Y = 0.018

	pointsOfFunctionPlot[808].X = 8.08
	pointsOfFunctionPlot[808].Y = 0.018

	pointsOfFunctionPlot[809].X = 8.09
	pointsOfFunctionPlot[809].Y = 0.018

	pointsOfFunctionPlot[810].X = 8.1
	pointsOfFunctionPlot[810].Y = 0.018

	pointsOfFunctionPlot[811].X = 8.11
	pointsOfFunctionPlot[811].Y = 0.018

	pointsOfFunctionPlot[812].X = 8.12
	pointsOfFunctionPlot[812].Y = 0.017

	pointsOfFunctionPlot[813].X = 8.13
	pointsOfFunctionPlot[813].Y = 0.017

	pointsOfFunctionPlot[814].X = 8.14
	pointsOfFunctionPlot[814].Y = 0.017

	pointsOfFunctionPlot[815].X = 8.15
	pointsOfFunctionPlot[815].Y = 0.017

	pointsOfFunctionPlot[816].X = 8.16
	pointsOfFunctionPlot[816].Y = 0.017

	pointsOfFunctionPlot[817].X = 8.17
	pointsOfFunctionPlot[817].Y = 0.017

	pointsOfFunctionPlot[818].X = 8.18
	pointsOfFunctionPlot[818].Y = 0.016

	pointsOfFunctionPlot[819].X = 8.19
	pointsOfFunctionPlot[819].Y = 0.016

	pointsOfFunctionPlot[820].X = 8.2
	pointsOfFunctionPlot[820].Y = 0.016

	pointsOfFunctionPlot[821].X = 8.21
	pointsOfFunctionPlot[821].Y = 0.016

	pointsOfFunctionPlot[822].X = 8.22
	pointsOfFunctionPlot[822].Y = 0.016

	pointsOfFunctionPlot[823].X = 8.23
	pointsOfFunctionPlot[823].Y = 0.016

	pointsOfFunctionPlot[824].X = 8.24
	pointsOfFunctionPlot[824].Y = 0.016

	pointsOfFunctionPlot[825].X = 8.25
	pointsOfFunctionPlot[825].Y = 0.015

	pointsOfFunctionPlot[826].X = 8.26
	pointsOfFunctionPlot[826].Y = 0.015

	pointsOfFunctionPlot[827].X = 8.27
	pointsOfFunctionPlot[827].Y = 0.015

	pointsOfFunctionPlot[828].X = 8.28
	pointsOfFunctionPlot[828].Y = 0.015

	pointsOfFunctionPlot[829].X = 8.29
	pointsOfFunctionPlot[829].Y = 0.015

	pointsOfFunctionPlot[830].X = 8.3
	pointsOfFunctionPlot[830].Y = 0.015

	pointsOfFunctionPlot[831].X = 8.31
	pointsOfFunctionPlot[831].Y = 0.015

	pointsOfFunctionPlot[832].X = 8.32
	pointsOfFunctionPlot[832].Y = 0.015

	pointsOfFunctionPlot[833].X = 8.33
	pointsOfFunctionPlot[833].Y = 0.014

	pointsOfFunctionPlot[834].X = 8.34
	pointsOfFunctionPlot[834].Y = 0.014

	pointsOfFunctionPlot[835].X = 8.35
	pointsOfFunctionPlot[835].Y = 0.014

	pointsOfFunctionPlot[836].X = 8.36
	pointsOfFunctionPlot[836].Y = 0.014

	pointsOfFunctionPlot[837].X = 8.37
	pointsOfFunctionPlot[837].Y = 0.014

	pointsOfFunctionPlot[838].X = 8.38
	pointsOfFunctionPlot[838].Y = 0.014

	pointsOfFunctionPlot[839].X = 8.39
	pointsOfFunctionPlot[839].Y = 0.014

	pointsOfFunctionPlot[840].X = 8.4
	pointsOfFunctionPlot[840].Y = 0.014

	pointsOfFunctionPlot[841].X = 8.41
	pointsOfFunctionPlot[841].Y = 0.013

	pointsOfFunctionPlot[842].X = 8.42
	pointsOfFunctionPlot[842].Y = 0.013

	pointsOfFunctionPlot[843].X = 8.43
	pointsOfFunctionPlot[843].Y = 0.013

	pointsOfFunctionPlot[844].X = 8.44
	pointsOfFunctionPlot[844].Y = 0.013

	pointsOfFunctionPlot[845].X = 8.45
	pointsOfFunctionPlot[845].Y = 0.013

	pointsOfFunctionPlot[846].X = 8.46
	pointsOfFunctionPlot[846].Y = 0.013

	pointsOfFunctionPlot[847].X = 8.47
	pointsOfFunctionPlot[847].Y = 0.013

	pointsOfFunctionPlot[848].X = 8.48
	pointsOfFunctionPlot[848].Y = 0.013

	pointsOfFunctionPlot[849].X = 8.49
	pointsOfFunctionPlot[849].Y = 0.012

	pointsOfFunctionPlot[850].X = 8.5
	pointsOfFunctionPlot[850].Y = 0.012

	pointsOfFunctionPlot[851].X = 8.51
	pointsOfFunctionPlot[851].Y = 0.012

	pointsOfFunctionPlot[852].X = 8.52
	pointsOfFunctionPlot[852].Y = 0.012

	pointsOfFunctionPlot[853].X = 8.53
	pointsOfFunctionPlot[853].Y = 0.012

	pointsOfFunctionPlot[854].X = 8.54
	pointsOfFunctionPlot[854].Y = 0.012

	pointsOfFunctionPlot[855].X = 8.55
	pointsOfFunctionPlot[855].Y = 0.012

	pointsOfFunctionPlot[856].X = 8.56
	pointsOfFunctionPlot[856].Y = 0.012

	pointsOfFunctionPlot[857].X = 8.57
	pointsOfFunctionPlot[857].Y = 0.012

	pointsOfFunctionPlot[858].X = 8.58
	pointsOfFunctionPlot[858].Y = 0.011

	pointsOfFunctionPlot[859].X = 8.59
	pointsOfFunctionPlot[859].Y = 0.011

	pointsOfFunctionPlot[860].X = 8.6
	pointsOfFunctionPlot[860].Y = 0.011

	pointsOfFunctionPlot[861].X = 8.61
	pointsOfFunctionPlot[861].Y = 0.011

	pointsOfFunctionPlot[862].X = 8.62
	pointsOfFunctionPlot[862].Y = 0.011

	pointsOfFunctionPlot[863].X = 8.63
	pointsOfFunctionPlot[863].Y = 0.011

	pointsOfFunctionPlot[864].X = 8.64
	pointsOfFunctionPlot[864].Y = 0.011

	pointsOfFunctionPlot[865].X = 8.65
	pointsOfFunctionPlot[865].Y = 0.011

	pointsOfFunctionPlot[866].X = 8.66
	pointsOfFunctionPlot[866].Y = 0.011

	pointsOfFunctionPlot[867].X = 8.67
	pointsOfFunctionPlot[867].Y = 0.011

	pointsOfFunctionPlot[868].X = 8.68
	pointsOfFunctionPlot[868].Y = 0.01

	pointsOfFunctionPlot[869].X = 8.69
	pointsOfFunctionPlot[869].Y = 0.01

	pointsOfFunctionPlot[870].X = 8.7
	pointsOfFunctionPlot[870].Y = 0.01

	pointsOfFunctionPlot[871].X = 8.71
	pointsOfFunctionPlot[871].Y = 0.01

	pointsOfFunctionPlot[872].X = 8.72
	pointsOfFunctionPlot[872].Y = 0.01

	pointsOfFunctionPlot[873].X = 8.73
	pointsOfFunctionPlot[873].Y = 0.01

	pointsOfFunctionPlot[874].X = 8.74
	pointsOfFunctionPlot[874].Y = 0.01

	pointsOfFunctionPlot[875].X = 8.75
	pointsOfFunctionPlot[875].Y = 0.01

	pointsOfFunctionPlot[876].X = 8.76
	pointsOfFunctionPlot[876].Y = 0.01

	pointsOfFunctionPlot[877].X = 8.77
	pointsOfFunctionPlot[877].Y = 0.01

	pointsOfFunctionPlot[878].X = 8.78
	pointsOfFunctionPlot[878].Y = 0.01

	pointsOfFunctionPlot[879].X = 8.79
	pointsOfFunctionPlot[879].Y = 0.009

	pointsOfFunctionPlot[880].X = 8.8
	pointsOfFunctionPlot[880].Y = 0.009

	pointsOfFunctionPlot[881].X = 8.81
	pointsOfFunctionPlot[881].Y = 0.009

	pointsOfFunctionPlot[882].X = 8.82
	pointsOfFunctionPlot[882].Y = 0.009

	pointsOfFunctionPlot[883].X = 8.83
	pointsOfFunctionPlot[883].Y = 0.009

	pointsOfFunctionPlot[884].X = 8.84
	pointsOfFunctionPlot[884].Y = 0.009

	pointsOfFunctionPlot[885].X = 8.85
	pointsOfFunctionPlot[885].Y = 0.009

	pointsOfFunctionPlot[886].X = 8.86
	pointsOfFunctionPlot[886].Y = 0.009

	pointsOfFunctionPlot[887].X = 8.87
	pointsOfFunctionPlot[887].Y = 0.009

	pointsOfFunctionPlot[888].X = 8.88
	pointsOfFunctionPlot[888].Y = 0.009

	pointsOfFunctionPlot[889].X = 8.89
	pointsOfFunctionPlot[889].Y = 0.009

	pointsOfFunctionPlot[890].X = 8.9
	pointsOfFunctionPlot[890].Y = 0.009

	pointsOfFunctionPlot[891].X = 8.91
	pointsOfFunctionPlot[891].Y = 0.008

	pointsOfFunctionPlot[892].X = 8.92
	pointsOfFunctionPlot[892].Y = 0.008

	pointsOfFunctionPlot[893].X = 8.93
	pointsOfFunctionPlot[893].Y = 0.008

	pointsOfFunctionPlot[894].X = 8.94
	pointsOfFunctionPlot[894].Y = 0.008

	pointsOfFunctionPlot[895].X = 8.95
	pointsOfFunctionPlot[895].Y = 0.008

	pointsOfFunctionPlot[896].X = 8.96
	pointsOfFunctionPlot[896].Y = 0.008

	pointsOfFunctionPlot[897].X = 8.97
	pointsOfFunctionPlot[897].Y = 0.008

	pointsOfFunctionPlot[898].X = 8.98
	pointsOfFunctionPlot[898].Y = 0.008

	pointsOfFunctionPlot[899].X = 8.99
	pointsOfFunctionPlot[899].Y = 0.008

	pointsOfFunctionPlot[900].X = 9.0
	pointsOfFunctionPlot[900].Y = 0.008

	pointsOfFunctionPlot[901].X = 9.01
	pointsOfFunctionPlot[901].Y = 0.008

	pointsOfFunctionPlot[902].X = 9.02
	pointsOfFunctionPlot[902].Y = 0.008

	pointsOfFunctionPlot[903].X = 9.03
	pointsOfFunctionPlot[903].Y = 0.008

	pointsOfFunctionPlot[904].X = 9.04
	pointsOfFunctionPlot[904].Y = 0.007

	pointsOfFunctionPlot[905].X = 9.05
	pointsOfFunctionPlot[905].Y = 0.007

	pointsOfFunctionPlot[906].X = 9.06
	pointsOfFunctionPlot[906].Y = 0.007

	pointsOfFunctionPlot[907].X = 9.07
	pointsOfFunctionPlot[907].Y = 0.007

	pointsOfFunctionPlot[908].X = 9.08
	pointsOfFunctionPlot[908].Y = 0.007

	pointsOfFunctionPlot[909].X = 9.09
	pointsOfFunctionPlot[909].Y = 0.007

	pointsOfFunctionPlot[910].X = 9.1
	pointsOfFunctionPlot[910].Y = 0.007

	pointsOfFunctionPlot[911].X = 9.11
	pointsOfFunctionPlot[911].Y = 0.007

	pointsOfFunctionPlot[912].X = 9.12
	pointsOfFunctionPlot[912].Y = 0.007

	pointsOfFunctionPlot[913].X = 9.13
	pointsOfFunctionPlot[913].Y = 0.007

	pointsOfFunctionPlot[914].X = 9.14
	pointsOfFunctionPlot[914].Y = 0.007

	pointsOfFunctionPlot[915].X = 9.15
	pointsOfFunctionPlot[915].Y = 0.007

	pointsOfFunctionPlot[916].X = 9.16
	pointsOfFunctionPlot[916].Y = 0.007

	pointsOfFunctionPlot[917].X = 9.17
	pointsOfFunctionPlot[917].Y = 0.007

	pointsOfFunctionPlot[918].X = 9.18
	pointsOfFunctionPlot[918].Y = 0.007

	pointsOfFunctionPlot[919].X = 9.19
	pointsOfFunctionPlot[919].Y = 0.006

	pointsOfFunctionPlot[920].X = 9.2
	pointsOfFunctionPlot[920].Y = 0.006

	pointsOfFunctionPlot[921].X = 9.21
	pointsOfFunctionPlot[921].Y = 0.006

	pointsOfFunctionPlot[922].X = 9.22
	pointsOfFunctionPlot[922].Y = 0.006

	pointsOfFunctionPlot[923].X = 9.23
	pointsOfFunctionPlot[923].Y = 0.006

	pointsOfFunctionPlot[924].X = 9.24
	pointsOfFunctionPlot[924].Y = 0.006

	pointsOfFunctionPlot[925].X = 9.25
	pointsOfFunctionPlot[925].Y = 0.006

	pointsOfFunctionPlot[926].X = 9.26
	pointsOfFunctionPlot[926].Y = 0.006

	pointsOfFunctionPlot[927].X = 9.27
	pointsOfFunctionPlot[927].Y = 0.006

	pointsOfFunctionPlot[928].X = 9.28
	pointsOfFunctionPlot[928].Y = 0.006

	pointsOfFunctionPlot[929].X = 9.29
	pointsOfFunctionPlot[929].Y = 0.006

	pointsOfFunctionPlot[930].X = 9.3
	pointsOfFunctionPlot[930].Y = 0.006

	pointsOfFunctionPlot[931].X = 9.31
	pointsOfFunctionPlot[931].Y = 0.006

	pointsOfFunctionPlot[932].X = 9.32
	pointsOfFunctionPlot[932].Y = 0.006

	pointsOfFunctionPlot[933].X = 9.33
	pointsOfFunctionPlot[933].Y = 0.006

	pointsOfFunctionPlot[934].X = 9.34
	pointsOfFunctionPlot[934].Y = 0.006

	pointsOfFunctionPlot[935].X = 9.35
	pointsOfFunctionPlot[935].Y = 0.006

	pointsOfFunctionPlot[936].X = 9.36
	pointsOfFunctionPlot[936].Y = 0.006

	pointsOfFunctionPlot[937].X = 9.37
	pointsOfFunctionPlot[937].Y = 0.005

	pointsOfFunctionPlot[938].X = 9.38
	pointsOfFunctionPlot[938].Y = 0.005

	pointsOfFunctionPlot[939].X = 9.39
	pointsOfFunctionPlot[939].Y = 0.005

	pointsOfFunctionPlot[940].X = 9.4
	pointsOfFunctionPlot[940].Y = 0.005

	pointsOfFunctionPlot[941].X = 9.41
	pointsOfFunctionPlot[941].Y = 0.005

	pointsOfFunctionPlot[942].X = 9.42
	pointsOfFunctionPlot[942].Y = 0.005

	pointsOfFunctionPlot[943].X = 9.43
	pointsOfFunctionPlot[943].Y = 0.005

	pointsOfFunctionPlot[944].X = 9.44
	pointsOfFunctionPlot[944].Y = 0.005

	pointsOfFunctionPlot[945].X = 9.45
	pointsOfFunctionPlot[945].Y = 0.005

	pointsOfFunctionPlot[946].X = 9.46
	pointsOfFunctionPlot[946].Y = 0.005

	pointsOfFunctionPlot[947].X = 9.47
	pointsOfFunctionPlot[947].Y = 0.005

	pointsOfFunctionPlot[948].X = 9.48
	pointsOfFunctionPlot[948].Y = 0.005

	pointsOfFunctionPlot[949].X = 9.49
	pointsOfFunctionPlot[949].Y = 0.005

	pointsOfFunctionPlot[950].X = 9.5
	pointsOfFunctionPlot[950].Y = 0.005

	pointsOfFunctionPlot[951].X = 9.51
	pointsOfFunctionPlot[951].Y = 0.005

	pointsOfFunctionPlot[952].X = 9.52
	pointsOfFunctionPlot[952].Y = 0.005

	pointsOfFunctionPlot[953].X = 9.53
	pointsOfFunctionPlot[953].Y = 0.005

	pointsOfFunctionPlot[954].X = 9.54
	pointsOfFunctionPlot[954].Y = 0.005

	pointsOfFunctionPlot[955].X = 9.55
	pointsOfFunctionPlot[955].Y = 0.005

	pointsOfFunctionPlot[956].X = 9.56
	pointsOfFunctionPlot[956].Y = 0.005

	pointsOfFunctionPlot[957].X = 9.57
	pointsOfFunctionPlot[957].Y = 0.004

	pointsOfFunctionPlot[958].X = 9.58
	pointsOfFunctionPlot[958].Y = 0.004

	pointsOfFunctionPlot[959].X = 9.59
	pointsOfFunctionPlot[959].Y = 0.004

	pointsOfFunctionPlot[960].X = 9.6
	pointsOfFunctionPlot[960].Y = 0.004

	pointsOfFunctionPlot[961].X = 9.61
	pointsOfFunctionPlot[961].Y = 0.004

	pointsOfFunctionPlot[962].X = 9.62
	pointsOfFunctionPlot[962].Y = 0.004

	pointsOfFunctionPlot[963].X = 9.63
	pointsOfFunctionPlot[963].Y = 0.004

	pointsOfFunctionPlot[964].X = 9.64
	pointsOfFunctionPlot[964].Y = 0.004

	pointsOfFunctionPlot[965].X = 9.65
	pointsOfFunctionPlot[965].Y = 0.004

	pointsOfFunctionPlot[966].X = 9.66
	pointsOfFunctionPlot[966].Y = 0.004

	pointsOfFunctionPlot[967].X = 9.67
	pointsOfFunctionPlot[967].Y = 0.004

	pointsOfFunctionPlot[968].X = 9.68
	pointsOfFunctionPlot[968].Y = 0.004

	pointsOfFunctionPlot[969].X = 9.69
	pointsOfFunctionPlot[969].Y = 0.004

	pointsOfFunctionPlot[970].X = 9.7
	pointsOfFunctionPlot[970].Y = 0.004

	pointsOfFunctionPlot[971].X = 9.71
	pointsOfFunctionPlot[971].Y = 0.004

	pointsOfFunctionPlot[972].X = 9.72
	pointsOfFunctionPlot[972].Y = 0.004

	pointsOfFunctionPlot[973].X = 9.73
	pointsOfFunctionPlot[973].Y = 0.004

	pointsOfFunctionPlot[974].X = 9.74
	pointsOfFunctionPlot[974].Y = 0.004

	pointsOfFunctionPlot[975].X = 9.75
	pointsOfFunctionPlot[975].Y = 0.004

	pointsOfFunctionPlot[976].X = 9.76
	pointsOfFunctionPlot[976].Y = 0.004

	pointsOfFunctionPlot[977].X = 9.77
	pointsOfFunctionPlot[977].Y = 0.004

	pointsOfFunctionPlot[978].X = 9.78
	pointsOfFunctionPlot[978].Y = 0.004

	pointsOfFunctionPlot[979].X = 9.79
	pointsOfFunctionPlot[979].Y = 0.004

	pointsOfFunctionPlot[980].X = 9.8
	pointsOfFunctionPlot[980].Y = 0.004

	pointsOfFunctionPlot[981].X = 9.81
	pointsOfFunctionPlot[981].Y = 0.004

	pointsOfFunctionPlot[982].X = 9.82
	pointsOfFunctionPlot[982].Y = 0.003

	pointsOfFunctionPlot[983].X = 9.83
	pointsOfFunctionPlot[983].Y = 0.003

	pointsOfFunctionPlot[984].X = 9.84
	pointsOfFunctionPlot[984].Y = 0.003

	pointsOfFunctionPlot[985].X = 9.85
	pointsOfFunctionPlot[985].Y = 0.003

	pointsOfFunctionPlot[986].X = 9.86
	pointsOfFunctionPlot[986].Y = 0.003

	pointsOfFunctionPlot[987].X = 9.87
	pointsOfFunctionPlot[987].Y = 0.003

	pointsOfFunctionPlot[988].X = 9.88
	pointsOfFunctionPlot[988].Y = 0.003

	pointsOfFunctionPlot[989].X = 9.89
	pointsOfFunctionPlot[989].Y = 0.003

	pointsOfFunctionPlot[990].X = 9.9
	pointsOfFunctionPlot[990].Y = 0.003

	pointsOfFunctionPlot[991].X = 9.91
	pointsOfFunctionPlot[991].Y = 0.003

	pointsOfFunctionPlot[992].X = 9.92
	pointsOfFunctionPlot[992].Y = 0.003

	pointsOfFunctionPlot[993].X = 9.93
	pointsOfFunctionPlot[993].Y = 0.003

	pointsOfFunctionPlot[994].X = 9.94
	pointsOfFunctionPlot[994].Y = 0.003

	pointsOfFunctionPlot[995].X = 9.95
	pointsOfFunctionPlot[995].Y = 0.003

	pointsOfFunctionPlot[996].X = 9.96
	pointsOfFunctionPlot[996].Y = 0.003

	pointsOfFunctionPlot[997].X = 9.97
	pointsOfFunctionPlot[997].Y = 0.003

	pointsOfFunctionPlot[998].X = 9.98
	pointsOfFunctionPlot[998].Y = 0.003

	pointsOfFunctionPlot[999].X = 9.99
	pointsOfFunctionPlot[999].Y = 0.003

	pointsOfFunctionPlot[1_000].X = 10.0
	pointsOfFunctionPlot[1_000].Y = 0.003

	pointsOfFunctionPlot[1_001].X = 10.01
	pointsOfFunctionPlot[1_001].Y = 0.003

	pointsOfFunctionPlot[1_002].X = 10.02
	pointsOfFunctionPlot[1_002].Y = 0.003

	pointsOfFunctionPlot[1_003].X = 10.03
	pointsOfFunctionPlot[1_003].Y = 0.003

	pointsOfFunctionPlot[1_004].X = 10.04
	pointsOfFunctionPlot[1_004].Y = 0.003

	pointsOfFunctionPlot[1_005].X = 10.05
	pointsOfFunctionPlot[1_005].Y = 0.003

	pointsOfFunctionPlot[1_006].X = 10.06
	pointsOfFunctionPlot[1_006].Y = 0.003

	pointsOfFunctionPlot[1_007].X = 10.07
	pointsOfFunctionPlot[1_007].Y = 0.003

	pointsOfFunctionPlot[1_008].X = 10.08
	pointsOfFunctionPlot[1_008].Y = 0.003

	pointsOfFunctionPlot[1_009].X = 10.09
	pointsOfFunctionPlot[1_009].Y = 0.003

	pointsOfFunctionPlot[1_010].X = 10.1
	pointsOfFunctionPlot[1_010].Y = 0.003

	pointsOfFunctionPlot[1_011].X = 10.11
	pointsOfFunctionPlot[1_011].Y = 0.003

	pointsOfFunctionPlot[1_012].X = 10.12
	pointsOfFunctionPlot[1_012].Y = 0.003

	pointsOfFunctionPlot[1_013].X = 10.13
	pointsOfFunctionPlot[1_013].Y = 0.003

	pointsOfFunctionPlot[1_014].X = 10.14
	pointsOfFunctionPlot[1_014].Y = 0.003

	pointsOfFunctionPlot[1_015].X = 10.15
	pointsOfFunctionPlot[1_015].Y = 0.002

	pointsOfFunctionPlot[1_016].X = 10.16
	pointsOfFunctionPlot[1_016].Y = 0.002

	pointsOfFunctionPlot[1_017].X = 10.17
	pointsOfFunctionPlot[1_017].Y = 0.002

	pointsOfFunctionPlot[1_018].X = 10.18
	pointsOfFunctionPlot[1_018].Y = 0.002

	pointsOfFunctionPlot[1_019].X = 10.19
	pointsOfFunctionPlot[1_019].Y = 0.002

	pointsOfFunctionPlot[1_020].X = 10.2
	pointsOfFunctionPlot[1_020].Y = 0.002

	pointsOfFunctionPlot[1_021].X = 10.21
	pointsOfFunctionPlot[1_021].Y = 0.002

	pointsOfFunctionPlot[1_022].X = 10.22
	pointsOfFunctionPlot[1_022].Y = 0.002

	pointsOfFunctionPlot[1_023].X = 10.23
	pointsOfFunctionPlot[1_023].Y = 0.002

	pointsOfFunctionPlot[1_024].X = 10.24
	pointsOfFunctionPlot[1_024].Y = 0.002

	pointsOfFunctionPlot[1_025].X = 10.25
	pointsOfFunctionPlot[1_025].Y = 0.002

	pointsOfFunctionPlot[1_026].X = 10.26
	pointsOfFunctionPlot[1_026].Y = 0.002

	pointsOfFunctionPlot[1_027].X = 10.27
	pointsOfFunctionPlot[1_027].Y = 0.002

	pointsOfFunctionPlot[1_028].X = 10.28
	pointsOfFunctionPlot[1_028].Y = 0.002

	pointsOfFunctionPlot[1_029].X = 10.29
	pointsOfFunctionPlot[1_029].Y = 0.002

	pointsOfFunctionPlot[1_030].X = 10.3
	pointsOfFunctionPlot[1_030].Y = 0.002

	pointsOfFunctionPlot[1_031].X = 10.31
	pointsOfFunctionPlot[1_031].Y = 0.002

	pointsOfFunctionPlot[1_032].X = 10.32
	pointsOfFunctionPlot[1_032].Y = 0.002

	pointsOfFunctionPlot[1_033].X = 10.33
	pointsOfFunctionPlot[1_033].Y = 0.002

	pointsOfFunctionPlot[1_034].X = 10.34
	pointsOfFunctionPlot[1_034].Y = 0.002

	pointsOfFunctionPlot[1_035].X = 10.35
	pointsOfFunctionPlot[1_035].Y = 0.002

	pointsOfFunctionPlot[1_036].X = 10.36
	pointsOfFunctionPlot[1_036].Y = 0.002

	pointsOfFunctionPlot[1_037].X = 10.37
	pointsOfFunctionPlot[1_037].Y = 0.002

	pointsOfFunctionPlot[1_038].X = 10.38
	pointsOfFunctionPlot[1_038].Y = 0.002

	pointsOfFunctionPlot[1_039].X = 10.39
	pointsOfFunctionPlot[1_039].Y = 0.002

	pointsOfFunctionPlot[1_040].X = 10.40
	pointsOfFunctionPlot[1_040].Y = 0.002

	pointsOfFunctionPlot[1_041].X = 10.41
	pointsOfFunctionPlot[1_041].Y = 0.002

	pointsOfFunctionPlot[1_042].X = 10.42
	pointsOfFunctionPlot[1_042].Y = 0.002

	pointsOfFunctionPlot[1_043].X = 10.43
	pointsOfFunctionPlot[1_043].Y = 0.002

	pointsOfFunctionPlot[1_044].X = 10.44
	pointsOfFunctionPlot[1_044].Y = 0.002

	pointsOfFunctionPlot[1_045].X = 10.45
	pointsOfFunctionPlot[1_045].Y = 0.002

	pointsOfFunctionPlot[1_046].X = 10.46
	pointsOfFunctionPlot[1_046].Y = 0.002

	pointsOfFunctionPlot[1_047].X = 10.47
	pointsOfFunctionPlot[1_047].Y = 0.002

	pointsOfFunctionPlot[1_048].X = 10.48
	pointsOfFunctionPlot[1_048].Y = 0.002

	pointsOfFunctionPlot[1_049].X = 10.49
	pointsOfFunctionPlot[1_049].Y = 0.002

	pointsOfFunctionPlot[1_050].X = 10.5
	pointsOfFunctionPlot[1_050].Y = 0.002

	pointsOfFunctionPlot[1_051].X = 10.51
	pointsOfFunctionPlot[1_051].Y = 0.002

	pointsOfFunctionPlot[1_052].X = 10.52
	pointsOfFunctionPlot[1_052].Y = 0.002

	pointsOfFunctionPlot[1_053].X = 10.53
	pointsOfFunctionPlot[1_053].Y = 0.002

	pointsOfFunctionPlot[1_054].X = 10.54
	pointsOfFunctionPlot[1_054].Y = 0.002

	pointsOfFunctionPlot[1_055].X = 10.55
	pointsOfFunctionPlot[1_055].Y = 0.002

	pointsOfFunctionPlot[1_056].X = 10.56
	pointsOfFunctionPlot[1_056].Y = 0.002

	pointsOfFunctionPlot[1_057].X = 10.57
	pointsOfFunctionPlot[1_057].Y = 0.002

	pointsOfFunctionPlot[1_058].X = 10.58
	pointsOfFunctionPlot[1_058].Y = 0.002

	pointsOfFunctionPlot[1_059].X = 10.59
	pointsOfFunctionPlot[1_059].Y = 0.002

	pointsOfFunctionPlot[1_060].X = 10.6
	pointsOfFunctionPlot[1_060].Y = 0.002

	pointsOfFunctionPlot[1_061].X = 10.61
	pointsOfFunctionPlot[1_061].Y = 0.001

	pointsOfFunctionPlot[1_062].X = 10.62
	pointsOfFunctionPlot[1_062].Y = 0.001

	pointsOfFunctionPlot[1_063].X = 10.63
	pointsOfFunctionPlot[1_063].Y = 0.001

	pointsOfFunctionPlot[1_064].X = 10.64
	pointsOfFunctionPlot[1_064].Y = 0.001

	pointsOfFunctionPlot[1_065].X = 10.65
	pointsOfFunctionPlot[1_065].Y = 0.001

	pointsOfFunctionPlot[1_066].X = 10.66
	pointsOfFunctionPlot[1_066].Y = 0.001

	pointsOfFunctionPlot[1_067].X = 10.67
	pointsOfFunctionPlot[1_067].Y = 0.001

	pointsOfFunctionPlot[1_068].X = 10.68
	pointsOfFunctionPlot[1_068].Y = 0.001

	pointsOfFunctionPlot[1_069].X = 10.69
	pointsOfFunctionPlot[1_069].Y = 0.001

	pointsOfFunctionPlot[1_070].X = 10.7
	pointsOfFunctionPlot[1_070].Y = 0.001

	pointsOfFunctionPlot[1_071].X = 10.71
	pointsOfFunctionPlot[1_071].Y = 0.001

	pointsOfFunctionPlot[1_072].X = 10.72
	pointsOfFunctionPlot[1_072].Y = 0.001

	pointsOfFunctionPlot[1_073].X = 10.73
	pointsOfFunctionPlot[1_073].Y = 0.001

	pointsOfFunctionPlot[1_074].X = 10.74
	pointsOfFunctionPlot[1_074].Y = 0.001

	pointsOfFunctionPlot[1_075].X = 10.75
	pointsOfFunctionPlot[1_075].Y = 0.001

	pointsOfFunctionPlot[1_076].X = 10.76
	pointsOfFunctionPlot[1_076].Y = 0.001

	pointsOfFunctionPlot[1_077].X = 10.77
	pointsOfFunctionPlot[1_077].Y = 0.001

	pointsOfFunctionPlot[1_078].X = 10.78
	pointsOfFunctionPlot[1_078].Y = 0.001

	pointsOfFunctionPlot[1_079].X = 10.79
	pointsOfFunctionPlot[1_079].Y = 0.001

	pointsOfFunctionPlot[1_080].X = 10.8
	pointsOfFunctionPlot[1_080].Y = 0.001

	pointsOfFunctionPlot[1_081].X = 10.81
	pointsOfFunctionPlot[1_081].Y = 0.001

	pointsOfFunctionPlot[1_082].X = 10.82
	pointsOfFunctionPlot[1_082].Y = 0.001

	pointsOfFunctionPlot[1_083].X = 10.83
	pointsOfFunctionPlot[1_083].Y = 0.001

	pointsOfFunctionPlot[1_084].X = 10.84
	pointsOfFunctionPlot[1_084].Y = 0.001

	pointsOfFunctionPlot[1_085].X = 10.85
	pointsOfFunctionPlot[1_085].Y = 0.001

	pointsOfFunctionPlot[1_086].X = 10.86
	pointsOfFunctionPlot[1_086].Y = 0.001

	pointsOfFunctionPlot[1_087].X = 10.87
	pointsOfFunctionPlot[1_087].Y = 0.001

	pointsOfFunctionPlot[1_088].X = 10.88
	pointsOfFunctionPlot[1_088].Y = 0.001

	pointsOfFunctionPlot[1_089].X = 10.89
	pointsOfFunctionPlot[1_089].Y = 0.001

	pointsOfFunctionPlot[1_090].X = 10.9
	pointsOfFunctionPlot[1_090].Y = 0.001

	pointsOfFunctionPlot[1_091].X = 10.91
	pointsOfFunctionPlot[1_091].Y = 0.001

	pointsOfFunctionPlot[1_092].X = 10.92
	pointsOfFunctionPlot[1_092].Y = 0.001

	pointsOfFunctionPlot[1_093].X = 10.93
	pointsOfFunctionPlot[1_093].Y = 0.001

	pointsOfFunctionPlot[1_094].X = 10.94
	pointsOfFunctionPlot[1_094].Y = 0.001

	pointsOfFunctionPlot[1_095].X = 10.95
	pointsOfFunctionPlot[1_095].Y = 0.001

	pointsOfFunctionPlot[1_096].X = 10.96
	pointsOfFunctionPlot[1_096].Y = 0.001

	pointsOfFunctionPlot[1_097].X = 10.97
	pointsOfFunctionPlot[1_097].Y = 0.001

	pointsOfFunctionPlot[1_098].X = 10.98
	pointsOfFunctionPlot[1_098].Y = 0.001

	pointsOfFunctionPlot[1_099].X = 10.99
	pointsOfFunctionPlot[1_099].Y = 0.001

	pointsOfFunctionPlot[1_100].X = 11.0
	pointsOfFunctionPlot[1_100].Y = 0.001

	pointsOfFunctionPlot[1_101].X = 11.01
	pointsOfFunctionPlot[1_101].Y = 0.001

	pointsOfFunctionPlot[1_102].X = 11.02
	pointsOfFunctionPlot[1_102].Y = 0.001

	pointsOfFunctionPlot[1_103].X = 11.03
	pointsOfFunctionPlot[1_103].Y = 0.001

	pointsOfFunctionPlot[1_104].X = 11.04
	pointsOfFunctionPlot[1_104].Y = 0.001

	pointsOfFunctionPlot[1_105].X = 11.05
	pointsOfFunctionPlot[1_105].Y = 0.001

	pointsOfFunctionPlot[1_106].X = 11.06
	pointsOfFunctionPlot[1_106].Y = 0.001

	pointsOfFunctionPlot[1_107].X = 11.07
	pointsOfFunctionPlot[1_107].Y = 0.001

	pointsOfFunctionPlot[1_108].X = 11.08
	pointsOfFunctionPlot[1_108].Y = 0.001

	pointsOfFunctionPlot[1_109].X = 11.09
	pointsOfFunctionPlot[1_109].Y = 0.001

	pointsOfFunctionPlot[1_110].X = 11.1
	pointsOfFunctionPlot[1_110].Y = 0.001

	pointsOfFunctionPlot[1_111].X = 11.11
	pointsOfFunctionPlot[1_111].Y = 0.001

	pointsOfFunctionPlot[1_112].X = 11.12
	pointsOfFunctionPlot[1_112].Y = 0.001

	pointsOfFunctionPlot[1_113].X = 11.13
	pointsOfFunctionPlot[1_113].Y = 0.001

	pointsOfFunctionPlot[1_114].X = 11.14
	pointsOfFunctionPlot[1_114].Y = 0.001

	pointsOfFunctionPlot[1_115].X = 11.15
	pointsOfFunctionPlot[1_115].Y = 0.001

	pointsOfFunctionPlot[1_116].X = 11.16
	pointsOfFunctionPlot[1_116].Y = 0.001

	pointsOfFunctionPlot[1_117].X = 11.17
	pointsOfFunctionPlot[1_117].Y = 0.001

	pointsOfFunctionPlot[1_118].X = 11.18
	pointsOfFunctionPlot[1_118].Y = 0.001

	pointsOfFunctionPlot[1_119].X = 11.19
	pointsOfFunctionPlot[1_119].Y = 0.001

	pointsOfFunctionPlot[1_120].X = 11.2
	pointsOfFunctionPlot[1_120].Y = 0.001

	pointsOfFunctionPlot[1_121].X = 11.21
	pointsOfFunctionPlot[1_121].Y = 0.001

	pointsOfFunctionPlot[1_122].X = 11.22
	pointsOfFunctionPlot[1_122].Y = 0.001

	pointsOfFunctionPlot[1_123].X = 11.23
	pointsOfFunctionPlot[1_123].Y = 0.001

	pointsOfFunctionPlot[1_124].X = 11.24
	pointsOfFunctionPlot[1_124].Y = 0.001

	pointsOfFunctionPlot[1_125].X = 11.25
	pointsOfFunctionPlot[1_125].Y = 0.001

	pointsOfFunctionPlot[1_126].X = 11.26
	pointsOfFunctionPlot[1_126].Y = 0.001

	pointsOfFunctionPlot[1_127].X = 11.27
	pointsOfFunctionPlot[1_127].Y = 0.001

	pointsOfFunctionPlot[1_128].X = 11.28
	pointsOfFunctionPlot[1_128].Y = 0.001

	pointsOfFunctionPlot[1_129].X = 11.29
	pointsOfFunctionPlot[1_129].Y = 0.001

	pointsOfFunctionPlot[1_130].X = 11.3
	pointsOfFunctionPlot[1_130].Y = 0.001

	pointsOfFunctionPlot[1_131].X = 11.31
	pointsOfFunctionPlot[1_131].Y = 0.001

	pointsOfFunctionPlot[1_132].X = 11.32
	pointsOfFunctionPlot[1_132].Y = 0.001

	pointsOfFunctionPlot[1_133].X = 11.33
	pointsOfFunctionPlot[1_133].Y = 0.001

	pointsOfFunctionPlot[1_134].X = 11.34
	pointsOfFunctionPlot[1_134].Y = 0.001

	pointsOfFunctionPlot[1_135].X = 11.35
	pointsOfFunctionPlot[1_135].Y = 0.001

	pointsOfFunctionPlot[1_136].X = 11.36
	pointsOfFunctionPlot[1_136].Y = 0.001

	pointsOfFunctionPlot[1_137].X = 11.37
	pointsOfFunctionPlot[1_137].Y = 0.001

	pointsOfFunctionPlot[1_138].X = 11.38
	pointsOfFunctionPlot[1_138].Y = 0.001

	pointsOfFunctionPlot[1_139].X = 11.39
	pointsOfFunctionPlot[1_139].Y = 0.001

	pointsOfFunctionPlot[1_140].X = 11.40
	pointsOfFunctionPlot[1_140].Y = 0.0

	pointsOfFunctionPlot[1_141].X = 12.0
	pointsOfFunctionPlot[1_141].Y = 0.0










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function f(x) = x exp(-(x - 2))"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("f(x)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"x-times-exp-function-plot-03.png"); err != nil {

		panic(err)
	}
}
