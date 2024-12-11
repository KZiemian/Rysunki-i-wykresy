package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wykres wielomianu f(x) = x^3 - x^2 + 2x + 1

	punktyWykresuWielomianu := make(plotter.XYs, 201)

	punktyWykresuWielomianu[0].X = 0.0
	punktyWykresuWielomianu[0].Y = 1.0

	punktyWykresuWielomianu[1].X = 0.05
	punktyWykresuWielomianu[1].Y = 1.097

	punktyWykresuWielomianu[2].X = 0.1
	punktyWykresuWielomianu[2].Y = 1.191

	punktyWykresuWielomianu[3].X = 0.15
	punktyWykresuWielomianu[3].Y = 1.28

	punktyWykresuWielomianu[4].X = 0.2
	punktyWykresuWielomianu[4].Y = 1.368

	punktyWykresuWielomianu[5].X = 0.25
	punktyWykresuWielomianu[5].Y = 1.453

	punktyWykresuWielomianu[6].X = 0.3
	punktyWykresuWielomianu[6].Y = 1.537

	punktyWykresuWielomianu[7].X = 0.35
	punktyWykresuWielomianu[7].Y = 1.62

	punktyWykresuWielomianu[8].X = 0.4
	punktyWykresuWielomianu[8].Y = 1.704

	punktyWykresuWielomianu[9].X = 0.45
	punktyWykresuWielomianu[9].Y = 1.788

	punktyWykresuWielomianu[10].X = 0.5
	punktyWykresuWielomianu[10].Y = 1.875

	punktyWykresuWielomianu[11].X = 0.55
	punktyWykresuWielomianu[11].Y = 1.963

	punktyWykresuWielomianu[12].X = 0.6
	punktyWykresuWielomianu[12].Y = 2.056

	punktyWykresuWielomianu[13].X = 0.65
	punktyWykresuWielomianu[13].Y = 2.152

	punktyWykresuWielomianu[14].X = 0.7
	punktyWykresuWielomianu[14].Y = 2.253

	punktyWykresuWielomianu[15].X = 0.75
	punktyWykresuWielomianu[15].Y = 2.359

	punktyWykresuWielomianu[16].X = 0.8
	punktyWykresuWielomianu[16].Y = 2.472

	punktyWykresuWielomianu[17].X = 0.85
	punktyWykresuWielomianu[17].Y = 2.591

	punktyWykresuWielomianu[18].X = 0.9
	punktyWykresuWielomianu[18].Y = 2.719

	punktyWykresuWielomianu[19].X = 0.95
	punktyWykresuWielomianu[19].Y = 2.854

	punktyWykresuWielomianu[20].X = 1.0
	punktyWykresuWielomianu[20].Y = 3.0

	punktyWykresuWielomianu[21].X = 1.05
	punktyWykresuWielomianu[21].Y = 3.155

	punktyWykresuWielomianu[22].X = 1.1
	punktyWykresuWielomianu[22].Y = 3.321

	punktyWykresuWielomianu[23].X = 1.15
	punktyWykresuWielomianu[23].Y = 3.498

	punktyWykresuWielomianu[24].X = 1.2
	punktyWykresuWielomianu[24].Y = 3.688

	punktyWykresuWielomianu[25].X = 1.25
	punktyWykresuWielomianu[25].Y = 3.89

	punktyWykresuWielomianu[26].X = 1.3
	punktyWykresuWielomianu[26].Y = 4.107

	punktyWykresuWielomianu[27].X = 1.35
	punktyWykresuWielomianu[27].Y = 4.337

	punktyWykresuWielomianu[28].X = 1.4
	punktyWykresuWielomianu[28].Y = 4.584

	punktyWykresuWielomianu[29].X = 1.45
	punktyWykresuWielomianu[29].Y = 4.846

	punktyWykresuWielomianu[30].X = 1.5
	punktyWykresuWielomianu[30].Y = 5.125

	punktyWykresuWielomianu[31].X = 1.55
	punktyWykresuWielomianu[31].Y = 5.421

	punktyWykresuWielomianu[32].X = 1.6
	punktyWykresuWielomianu[32].Y = 5.736

	punktyWykresuWielomianu[33].X = 1.65
	punktyWykresuWielomianu[33].Y = 6.069

	punktyWykresuWielomianu[34].X = 1.7
	punktyWykresuWielomianu[34].Y = 6.423

	punktyWykresuWielomianu[35].X = 1.75
	punktyWykresuWielomianu[35].Y = 6.796

	punktyWykresuWielomianu[36].X = 1.8
	punktyWykresuWielomianu[36].Y = 7.192

	punktyWykresuWielomianu[37].X = 1.85
	punktyWykresuWielomianu[37].Y = 7.609

	punktyWykresuWielomianu[38].X = 1.9
	punktyWykresuWielomianu[38].Y = 8.049

	punktyWykresuWielomianu[39].X = 1.95
	punktyWykresuWielomianu[39].Y = 8.512

	punktyWykresuWielomianu[40].X = 2.0
	punktyWykresuWielomianu[40].Y = 9.0

	punktyWykresuWielomianu[41].X = 2.05
	punktyWykresuWielomianu[41].Y = 9.512

	punktyWykresuWielomianu[42].X = 2.1
	punktyWykresuWielomianu[42].Y = 10.051

	punktyWykresuWielomianu[43].X = 2.15
	punktyWykresuWielomianu[43].Y = 10.615

	punktyWykresuWielomianu[44].X = 2.2
	punktyWykresuWielomianu[44].Y = 11.208

	punktyWykresuWielomianu[45].X = 2.25
	punktyWykresuWielomianu[45].Y = 11.828

	punktyWykresuWielomianu[46].X = 2.3
	punktyWykresuWielomianu[46].Y = 12.477

	punktyWykresuWielomianu[47].X = 2.35
	punktyWykresuWielomianu[47].Y = 13.155

	punktyWykresuWielomianu[48].X = 2.4
	punktyWykresuWielomianu[48].Y = 13.864

	punktyWykresuWielomianu[49].X = 2.45
	punktyWykresuWielomianu[49].Y = 14.603

	punktyWykresuWielomianu[50].X = 2.5
	punktyWykresuWielomianu[50].Y = 15.375

	punktyWykresuWielomianu[51].X = 2.55
	punktyWykresuWielomianu[51].Y = 16.178

	punktyWykresuWielomianu[52].X = 2.6
	punktyWykresuWielomianu[52].Y = 17.016

	punktyWykresuWielomianu[53].X = 2.65
	punktyWykresuWielomianu[53].Y = 17.887

	punktyWykresuWielomianu[54].X = 2.7
	punktyWykresuWielomianu[54].Y = 18.793

	punktyWykresuWielomianu[55].X = 2.75
	punktyWykresuWielomianu[55].Y = 19.734

	punktyWykresuWielomianu[56].X = 2.8
	punktyWykresuWielomianu[56].Y = 20.712

	punktyWykresuWielomianu[57].X = 2.85
	punktyWykresuWielomianu[57].Y = 21.726

	punktyWykresuWielomianu[58].X = 2.9
	punktyWykresuWielomianu[58].Y = 22.779

	punktyWykresuWielomianu[59].X = 2.95
	punktyWykresuWielomianu[59].Y = 23.869

	punktyWykresuWielomianu[60].X = 3.0
	punktyWykresuWielomianu[60].Y = 25.0

	punktyWykresuWielomianu[61].X = 3.05
	punktyWykresuWielomianu[61].Y = 26.17

	punktyWykresuWielomianu[62].X = 3.1
	punktyWykresuWielomianu[62].Y = 27.381

	punktyWykresuWielomianu[63].X = 3.15
	punktyWykresuWielomianu[63].Y = 28.633

	punktyWykresuWielomianu[64].X = 3.2
	punktyWykresuWielomianu[64].Y = 29.928

	punktyWykresuWielomianu[65].X = 3.25
	punktyWykresuWielomianu[65].Y = 31.265

	punktyWykresuWielomianu[66].X = 3.3
	punktyWykresuWielomianu[66].Y = 32.647

	punktyWykresuWielomianu[67].X = 3.35
	punktyWykresuWielomianu[67].Y = 34.072

	punktyWykresuWielomianu[68].X = 3.4
	punktyWykresuWielomianu[68].Y = 35.544

	punktyWykresuWielomianu[69].X = 3.45
	punktyWykresuWielomianu[69].Y = 37.061

	punktyWykresuWielomianu[70].X = 3.5
	punktyWykresuWielomianu[70].Y = 38.625

	punktyWykresuWielomianu[71].X = 3.55
	punktyWykresuWielomianu[71].Y = 40.236

	punktyWykresuWielomianu[72].X = 3.6
	punktyWykresuWielomianu[72].Y = 41.896

	punktyWykresuWielomianu[73].X = 3.65
	punktyWykresuWielomianu[73].Y = 43.604

	punktyWykresuWielomianu[74].X = 3.7
	punktyWykresuWielomianu[74].Y = 45.363

	punktyWykresuWielomianu[75].X = 3.75
	punktyWykresuWielomianu[75].Y = 47.171

	punktyWykresuWielomianu[76].X = 3.8
	punktyWykresuWielomianu[76].Y = 49.032

	punktyWykresuWielomianu[77].X = 3.85
	punktyWykresuWielomianu[77].Y = 50.944

	punktyWykresuWielomianu[78].X = 3.9
	punktyWykresuWielomianu[78].Y = 52.909

	punktyWykresuWielomianu[79].X = 3.95
	punktyWykresuWielomianu[79].Y = 54.927

	punktyWykresuWielomianu[80].X = 4.0
	punktyWykresuWielomianu[80].Y = 57.0

	punktyWykresuWielomianu[81].X = 4.05
	punktyWykresuWielomianu[81].Y = 59.127

	punktyWykresuWielomianu[82].X = 4.1
	punktyWykresuWielomianu[82].Y = 61.311

	punktyWykresuWielomianu[83].X = 4.15
	punktyWykresuWielomianu[83].Y = 63.55

	punktyWykresuWielomianu[84].X = 4.2
	punktyWykresuWielomianu[84].Y = 65.848

	punktyWykresuWielomianu[85].X = 4.25
	punktyWykresuWielomianu[85].Y = 68.848

	punktyWykresuWielomianu[86].X = 4.3
	punktyWykresuWielomianu[86].Y = 70.617

	punktyWykresuWielomianu[87].X = 4.35
	punktyWykresuWielomianu[87].Y = 73.09

	punktyWykresuWielomianu[88].X = 4.4
	punktyWykresuWielomianu[88].Y = 75.624

	punktyWykresuWielomianu[89].X = 4.45
	punktyWykresuWielomianu[89].Y = 78.218

	punktyWykresuWielomianu[90].X = 4.5
	punktyWykresuWielomianu[90].Y = 80.875

	punktyWykresuWielomianu[91].X = 4.55
	punktyWykresuWielomianu[91].Y = 83.593

	punktyWykresuWielomianu[92].X = 4.6
	punktyWykresuWielomianu[92].Y = 86.376

	punktyWykresuWielomianu[93].X = 4.65
	punktyWykresuWielomianu[93].Y = 89.222

	punktyWykresuWielomianu[94].X = 4.7
	punktyWykresuWielomianu[94].Y = 92.133

	punktyWykresuWielomianu[95].X = 4.75
	punktyWykresuWielomianu[95].Y = 95.109

	punktyWykresuWielomianu[96].X = 4.8
	punktyWykresuWielomianu[96].Y = 98.152

	punktyWykresuWielomianu[97].X = 4.85
	punktyWykresuWielomianu[97].Y = 101.261

	punktyWykresuWielomianu[98].X = 4.9
	punktyWykresuWielomianu[98].Y = 104.439

	punktyWykresuWielomianu[99].X = 4.95
	punktyWykresuWielomianu[99].Y = 107.684

	punktyWykresuWielomianu[100].X = 5.0
	punktyWykresuWielomianu[100].Y = 111.0

	punktyWykresuWielomianu[101].X = 5.05
	punktyWykresuWielomianu[101].Y = 114.385

	punktyWykresuWielomianu[102].X = 5.1
	punktyWykresuWielomianu[102].Y = 117.841

	punktyWykresuWielomianu[103].X = 5.15
	punktyWykresuWielomianu[103].Y = 121.368

	punktyWykresuWielomianu[104].X = 5.2
	punktyWykresuWielomianu[104].Y = 124.968

	punktyWykresuWielomianu[105].X = 5.25
	punktyWykresuWielomianu[105].Y = 128.64

	punktyWykresuWielomianu[106].X = 5.3
	punktyWykresuWielomianu[106].Y = 132.387

	punktyWykresuWielomianu[107].X = 5.35
	punktyWykresuWielomianu[107].Y = 136.207

	punktyWykresuWielomianu[108].X = 5.4
	punktyWykresuWielomianu[108].Y = 140.104

	punktyWykresuWielomianu[109].X = 5.45
	punktyWykresuWielomianu[109].Y = 144.076

	punktyWykresuWielomianu[110].X = 5.5
	punktyWykresuWielomianu[110].Y = 148.125

	punktyWykresuWielomianu[111].X = 5.55
	punktyWykresuWielomianu[111].Y = 152.251

	punktyWykresuWielomianu[112].X = 5.6
	punktyWykresuWielomianu[112].Y = 156.456

	punktyWykresuWielomianu[113].X = 5.65
	punktyWykresuWielomianu[113].Y = 160.739

	punktyWykresuWielomianu[114].X = 5.7
	punktyWykresuWielomianu[114].Y = 165.103

	punktyWykresuWielomianu[115].X = 5.75
	punktyWykresuWielomianu[115].Y = 169.546

	punktyWykresuWielomianu[116].X = 5.8
	punktyWykresuWielomianu[116].Y = 174.072

	punktyWykresuWielomianu[117].X = 5.85
	punktyWykresuWielomianu[117].Y = 178.679

	punktyWykresuWielomianu[118].X = 5.9
	punktyWykresuWielomianu[118].Y = 183.369

	punktyWykresuWielomianu[119].X = 5.95
	punktyWykresuWielomianu[119].Y = 188.142

	punktyWykresuWielomianu[120].X = 6.0
	punktyWykresuWielomianu[120].Y = 193.0

	punktyWykresuWielomianu[121].X = 6.05
	punktyWykresuWielomianu[121].Y = 197.942

	punktyWykresuWielomianu[122].X = 6.1
	punktyWykresuWielomianu[122].Y = 202.971

	punktyWykresuWielomianu[123].X = 6.15
	punktyWykresuWielomianu[123].Y = 208.085

	punktyWykresuWielomianu[124].X = 6.2
	punktyWykresuWielomianu[124].Y = 213.288

	punktyWykresuWielomianu[125].X = 6.25
	punktyWykresuWielomianu[125].Y = 218.578

	punktyWykresuWielomianu[126].X = 6.3
	punktyWykresuWielomianu[126].Y = 223.957

	punktyWykresuWielomianu[127].X = 6.35
	punktyWykresuWielomianu[127].Y = 229.425

	punktyWykresuWielomianu[128].X = 6.4
	punktyWykresuWielomianu[128].Y = 234.984

	punktyWykresuWielomianu[129].X = 6.45
	punktyWykresuWielomianu[129].Y = 240.633

	punktyWykresuWielomianu[130].X = 6.5
	punktyWykresuWielomianu[130].Y = 246.375

	punktyWykresuWielomianu[131].X = 6.55
	punktyWykresuWielomianu[131].Y = 252.208

	punktyWykresuWielomianu[132].X = 6.6
	punktyWykresuWielomianu[132].Y = 258.136

	punktyWykresuWielomianu[133].X = 6.65
	punktyWykresuWielomianu[133].Y = 264.157

	punktyWykresuWielomianu[134].X = 6.7
	punktyWykresuWielomianu[134].Y = 270.273

	punktyWykresuWielomianu[135].X = 6.75
	punktyWykresuWielomianu[135].Y = 276.484

	punktyWykresuWielomianu[136].X = 6.8
	punktyWykresuWielomianu[136].Y = 282.792

	punktyWykresuWielomianu[137].X = 6.85
	punktyWykresuWielomianu[137].Y = 289.196

	punktyWykresuWielomianu[138].X = 6.9
	punktyWykresuWielomianu[138].Y = 295.699

	punktyWykresuWielomianu[139].X = 6.95
	punktyWykresuWielomianu[139].Y = 302.299

	punktyWykresuWielomianu[140].X = 7.0
	punktyWykresuWielomianu[140].Y = 309.0

	punktyWykresuWielomianu[141].X = 7.05
	punktyWykresuWielomianu[141].Y = 315.8

	punktyWykresuWielomianu[142].X = 7.1
	punktyWykresuWielomianu[142].Y = 322.701

	punktyWykresuWielomianu[143].X = 7.15
	punktyWykresuWielomianu[143].Y = 329.703

	punktyWykresuWielomianu[144].X = 7.2
	punktyWykresuWielomianu[144].Y = 336.808

	punktyWykresuWielomianu[145].X = 7.25
	punktyWykresuWielomianu[145].Y = 344.015

	punktyWykresuWielomianu[146].X = 7.3
	punktyWykresuWielomianu[146].Y = 351.327

	punktyWykresuWielomianu[147].X = 7.35
	punktyWykresuWielomianu[147].Y = 358.742

	punktyWykresuWielomianu[148].X = 7.4
	punktyWykresuWielomianu[148].Y = 366.264

	punktyWykresuWielomianu[149].X = 7.45
	punktyWykresuWielomianu[149].Y = 373.891

	punktyWykresuWielomianu[150].X = 7.5
	punktyWykresuWielomianu[150].Y = 381.625

	punktyWykresuWielomianu[151].X = 7.55
	punktyWykresuWielomianu[151].Y = 389.466

	punktyWykresuWielomianu[152].X = 7.6
	punktyWykresuWielomianu[152].Y = 397.416

	punktyWykresuWielomianu[153].X = 7.65
	punktyWykresuWielomianu[153].Y = 405.474

	punktyWykresuWielomianu[154].X = 7.7
	punktyWykresuWielomianu[154].Y = 413.643

	punktyWykresuWielomianu[155].X = 7.75
	punktyWykresuWielomianu[155].Y = 421.921

	punktyWykresuWielomianu[156].X = 7.8
	punktyWykresuWielomianu[156].Y = 430.312

	punktyWykresuWielomianu[157].X = 7.85
	punktyWykresuWielomianu[157].Y = 438.814

	punktyWykresuWielomianu[158].X = 7.9
	punktyWykresuWielomianu[158].Y = 447.429

	punktyWykresuWielomianu[159].X = 7.95
	punktyWykresuWielomianu[159].Y = 456.157

	punktyWykresuWielomianu[160].X = 8.0
	punktyWykresuWielomianu[160].Y = 465.0

	punktyWykresuWielomianu[161].X = 8.05
	punktyWykresuWielomianu[161].Y = 473.957

	punktyWykresuWielomianu[162].X = 8.1
	punktyWykresuWielomianu[162].Y = 483.031

	punktyWykresuWielomianu[163].X = 8.15
	punktyWykresuWielomianu[163].Y = 492.22

	punktyWykresuWielomianu[164].X = 8.2
	punktyWykresuWielomianu[164].Y = 501.528

	punktyWykresuWielomianu[165].X = 8.25
	punktyWykresuWielomianu[165].Y = 510.953

	punktyWykresuWielomianu[166].X = 8.3
	punktyWykresuWielomianu[166].Y = 520.497

	punktyWykresuWielomianu[167].X = 8.35
	punktyWykresuWielomianu[167].Y = 530.16

	punktyWykresuWielomianu[168].X = 8.4
	punktyWykresuWielomianu[168].Y = 539.944

	punktyWykresuWielomianu[169].X = 8.45
	punktyWykresuWielomianu[169].Y = 549.848

	punktyWykresuWielomianu[170].X = 8.5
	punktyWykresuWielomianu[170].Y = 559.875

	punktyWykresuWielomianu[171].X = 8.55
	punktyWykresuWielomianu[171].Y = 570.023

	punktyWykresuWielomianu[172].X = 8.6
	punktyWykresuWielomianu[172].Y = 580.296

	punktyWykresuWielomianu[173].X = 8.65
	punktyWykresuWielomianu[173].Y = 590.692

	punktyWykresuWielomianu[174].X = 8.7
	punktyWykresuWielomianu[174].Y = 601.213

	punktyWykresuWielomianu[175].X = 8.75
	punktyWykresuWielomianu[175].Y = 611.859

	punktyWykresuWielomianu[176].X = 8.8
	punktyWykresuWielomianu[176].Y = 622.632

	punktyWykresuWielomianu[177].X = 8.85
	punktyWykresuWielomianu[177].Y = 633.531

	punktyWykresuWielomianu[178].X = 8.9
	punktyWykresuWielomianu[178].Y = 644.559

	punktyWykresuWielomianu[179].X = 8.95
	punktyWykresuWielomianu[179].Y = 655.714

	punktyWykresuWielomianu[180].X = 9.0
	punktyWykresuWielomianu[180].Y = 667.0

	punktyWykresuWielomianu[181].X = 9.05
	punktyWykresuWielomianu[181].Y = 678.415

	punktyWykresuWielomianu[182].X = 9.1
	punktyWykresuWielomianu[182].Y = 689.961

	punktyWykresuWielomianu[183].X = 9.15
	punktyWykresuWielomianu[183].Y = 701.638

	punktyWykresuWielomianu[184].X = 9.2
	punktyWykresuWielomianu[184].Y = 713.448

	punktyWykresuWielomianu[185].X = 9.25
	punktyWykresuWielomianu[185].Y = 725.39

	punktyWykresuWielomianu[186].X = 9.3
	punktyWykresuWielomianu[186].Y = 737.467

	punktyWykresuWielomianu[187].X = 9.35
	punktyWykresuWielomianu[187].Y = 749.677

	punktyWykresuWielomianu[188].X = 9.4
	punktyWykresuWielomianu[188].Y = 762.024

	punktyWykresuWielomianu[189].X = 9.45
	punktyWykresuWielomianu[189].Y = 774.506

	punktyWykresuWielomianu[190].X = 9.5
	punktyWykresuWielomianu[190].Y = 787.125

	punktyWykresuWielomianu[191].X = 9.55
	punktyWykresuWielomianu[191].Y = 799.881

	punktyWykresuWielomianu[192].X = 9.6
	punktyWykresuWielomianu[192].Y = 812.776

	punktyWykresuWielomianu[193].X = 9.65
	punktyWykresuWielomianu[193].Y = 825.809

	punktyWykresuWielomianu[194].X = 9.7
	punktyWykresuWielomianu[194].Y = 838.983

	punktyWykresuWielomianu[195].X = 9.75
	punktyWykresuWielomianu[195].Y = 852.296

	punktyWykresuWielomianu[196].X = 9.8
	punktyWykresuWielomianu[196].Y = 865.752

	punktyWykresuWielomianu[197].X = 9.85
	punktyWykresuWielomianu[197].Y = 879.349

	punktyWykresuWielomianu[198].X = 9.9
	punktyWykresuWielomianu[198].Y = 893.089

	punktyWykresuWielomianu[199].X = 9.95
	punktyWykresuWielomianu[199].Y = 906.972

	punktyWykresuWielomianu[200].X = 10.0
	punktyWykresuWielomianu[200].Y = 921.0



	wykresWielomianu := plot.New()

	wykresWielomianu.Title.Text = "Wykres funkcji f(x) = x^3 - x^2 + 2x + 1"

	wykresWielomianu.X.Label.Text = "x"
	wykresWielomianu.Y.Label.Text = "y"

	liniaWykresu, err := plotter.NewLine(punktyWykresuWielomianu)

	if err != nil {
		panic(err)
	}

	liniaWykresu.LineStyle.Width = vg.Points(1)
	liniaWykresu.Color = color.RGBA{R: 200, G: 100, B: 100}


	wykresWielomianu.Add(liniaWykresu)
	wykresWielomianu.Legend.Add("f(x)", liniaWykresu)

	if err := wykresWielomianu.Save(10*vg.Inch, 10*vg.Inch,
		"Wykres-wielomianu-01.png"); err != nil {

		panic(err)
	}
}
