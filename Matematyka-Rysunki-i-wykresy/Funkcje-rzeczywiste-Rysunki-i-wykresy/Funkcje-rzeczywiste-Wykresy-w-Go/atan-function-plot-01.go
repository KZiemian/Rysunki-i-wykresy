package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Plot of function atan(x).

	pointsOfFunctionPlot := make(plotter.XYs, 1_001)

	pointsOfFunctionPlot[0].X = -5.0
	pointsOfFunctionPlot[0].Y = -1.373

	pointsOfFunctionPlot[1].X = -4.99
	pointsOfFunctionPlot[1].Y = -1.373

	pointsOfFunctionPlot[2].X = -4.98
	pointsOfFunctionPlot[2].Y = -1.372

	pointsOfFunctionPlot[3].X = -4.97
	pointsOfFunctionPlot[3].Y = -1.372

	pointsOfFunctionPlot[4].X = -4.96
	pointsOfFunctionPlot[4].Y = -1.371

	pointsOfFunctionPlot[5].X = -4.95
	pointsOfFunctionPlot[5].Y = -1.371

	pointsOfFunctionPlot[6].X = -4.94
	pointsOfFunctionPlot[6].Y = -1.371

	pointsOfFunctionPlot[7].X = -4.93
	pointsOfFunctionPlot[7].Y = -1.37

	pointsOfFunctionPlot[8].X = -4.92
	pointsOfFunctionPlot[8].Y = -1.37

	pointsOfFunctionPlot[9].X = -4.91
	pointsOfFunctionPlot[9].Y = -1.369

	pointsOfFunctionPlot[10].X = -4.90
	pointsOfFunctionPlot[10].Y = -1.369

	pointsOfFunctionPlot[11].X = -4.89
	pointsOfFunctionPlot[11].Y = -1.369

	pointsOfFunctionPlot[12].X = -4.88
	pointsOfFunctionPlot[12].Y = -1.368

	pointsOfFunctionPlot[13].X = -4.87
	pointsOfFunctionPlot[13].Y = -1.368

	pointsOfFunctionPlot[14].X = -4.86
	pointsOfFunctionPlot[14].Y = -1.367

	pointsOfFunctionPlot[15].X = -4.85
	pointsOfFunctionPlot[15].Y = -1.367

	pointsOfFunctionPlot[16].X = -4.84
	pointsOfFunctionPlot[16].Y = -1.367

	pointsOfFunctionPlot[17].X = -4.83
	pointsOfFunctionPlot[17].Y = -1.366

	pointsOfFunctionPlot[18].X = -4.82
	pointsOfFunctionPlot[18].Y = -1.366

	pointsOfFunctionPlot[19].X = -4.81
	pointsOfFunctionPlot[19].Y = -1.365

	pointsOfFunctionPlot[20].X = -4.80
	pointsOfFunctionPlot[20].Y = -1.365

	pointsOfFunctionPlot[21].X = -4.79
	pointsOfFunctionPlot[21].Y = -1.364

	pointsOfFunctionPlot[22].X = -4.78
	pointsOfFunctionPlot[22].Y = -1.364

	pointsOfFunctionPlot[23].X = -4.77
	pointsOfFunctionPlot[23].Y = -1.364

	pointsOfFunctionPlot[24].X = -4.76
	pointsOfFunctionPlot[24].Y = -1.363

	pointsOfFunctionPlot[25].X = -4.75
	pointsOfFunctionPlot[25].Y = -1.363

	pointsOfFunctionPlot[26].X = -4.74
	pointsOfFunctionPlot[26].Y = -1.362

	pointsOfFunctionPlot[27].X = -4.73
	pointsOfFunctionPlot[27].Y = -1.362

	pointsOfFunctionPlot[28].X = -4.72
	pointsOfFunctionPlot[28].Y = -1.362

	pointsOfFunctionPlot[29].X = -4.71
	pointsOfFunctionPlot[29].Y = -1.361

	pointsOfFunctionPlot[30].X = -4.70
	pointsOfFunctionPlot[30].Y = -1.361

	pointsOfFunctionPlot[31].X = -4.69
	pointsOfFunctionPlot[31].Y = -1.36

	pointsOfFunctionPlot[32].X = -4.68
	pointsOfFunctionPlot[32].Y = -1.36

	pointsOfFunctionPlot[33].X = -4.67
	pointsOfFunctionPlot[33].Y = -1.359

	pointsOfFunctionPlot[34].X = -4.66
	pointsOfFunctionPlot[34].Y = -1.359

	pointsOfFunctionPlot[35].X = -4.65
	pointsOfFunctionPlot[35].Y = -1.358

	pointsOfFunctionPlot[36].X = -4.64
	pointsOfFunctionPlot[36].Y = -1.358

	pointsOfFunctionPlot[37].X = -4.63
	pointsOfFunctionPlot[37].Y = -1.358

	pointsOfFunctionPlot[38].X = -4.62
	pointsOfFunctionPlot[38].Y = -1.357

	pointsOfFunctionPlot[39].X = -4.61
	pointsOfFunctionPlot[39].Y = -1.357

	pointsOfFunctionPlot[40].X = -4.60
	pointsOfFunctionPlot[40].Y = -1.356

	pointsOfFunctionPlot[41].X = -4.59
	pointsOfFunctionPlot[41].Y = -1.356

	pointsOfFunctionPlot[42].X = -4.58
	pointsOfFunctionPlot[42].Y = -1.355

	pointsOfFunctionPlot[43].X = -4.57
	pointsOfFunctionPlot[43].Y = -1.355

	pointsOfFunctionPlot[44].X = -4.56
	pointsOfFunctionPlot[44].Y = -1.354

	pointsOfFunctionPlot[45].X = -4.55
	pointsOfFunctionPlot[45].Y = -1.354

	pointsOfFunctionPlot[46].X = -4.54
	pointsOfFunctionPlot[46].Y = -1.353

	pointsOfFunctionPlot[47].X = -4.53
	pointsOfFunctionPlot[47].Y = -1.353

	pointsOfFunctionPlot[48].X = -4.52
	pointsOfFunctionPlot[48].Y = -1.353

	pointsOfFunctionPlot[49].X = -4.51
	pointsOfFunctionPlot[49].Y = -1.352

	pointsOfFunctionPlot[50].X = -4.50
	pointsOfFunctionPlot[50].Y = -1.352

	pointsOfFunctionPlot[51].X = -4.49
	pointsOfFunctionPlot[51].Y = -1.351

	pointsOfFunctionPlot[52].X = -4.48
	pointsOfFunctionPlot[52].Y = -1.351

	pointsOfFunctionPlot[53].X = -4.47
	pointsOfFunctionPlot[53].Y = -1.35

	pointsOfFunctionPlot[54].X = -4.46
	pointsOfFunctionPlot[54].Y = -1.35

	pointsOfFunctionPlot[55].X = -4.45
	pointsOfFunctionPlot[55].Y = -1.349

	pointsOfFunctionPlot[56].X = -4.44
	pointsOfFunctionPlot[56].Y = -1.349

	pointsOfFunctionPlot[57].X = -4.43
	pointsOfFunctionPlot[57].Y = -1.348

	pointsOfFunctionPlot[58].X = -4.42
	pointsOfFunctionPlot[58].Y = -1.348

	pointsOfFunctionPlot[59].X = -4.41
	pointsOfFunctionPlot[59].Y = -1.347

	pointsOfFunctionPlot[60].X = -4.40
	pointsOfFunctionPlot[60].Y = -1.347

	pointsOfFunctionPlot[61].X = -4.39
	pointsOfFunctionPlot[61].Y = -1.346

	pointsOfFunctionPlot[62].X = -4.38
	pointsOfFunctionPlot[62].Y = -1.346

	pointsOfFunctionPlot[63].X = -4.37
	pointsOfFunctionPlot[63].Y = -1.345

	pointsOfFunctionPlot[64].X = -4.36
	pointsOfFunctionPlot[64].Y = -1.345

	pointsOfFunctionPlot[65].X = -4.35
	pointsOfFunctionPlot[65].Y = -1.344

	pointsOfFunctionPlot[66].X = -4.34
	pointsOfFunctionPlot[66].Y = -1.344

	pointsOfFunctionPlot[67].X = -4.33
	pointsOfFunctionPlot[67].Y = -1.343

	pointsOfFunctionPlot[68].X = -4.32
	pointsOfFunctionPlot[68].Y = -1.343

	pointsOfFunctionPlot[69].X = -4.31
	pointsOfFunctionPlot[69].Y = -1.342

	pointsOfFunctionPlot[70].X = -4.30
	pointsOfFunctionPlot[70].Y = -1.342

	pointsOfFunctionPlot[71].X = -4.29
	pointsOfFunctionPlot[71].Y = -1.341

	pointsOfFunctionPlot[72].X = -4.28
	pointsOfFunctionPlot[72].Y = -1.341

	pointsOfFunctionPlot[73].X = -4.27
	pointsOfFunctionPlot[73].Y = -1.34

	pointsOfFunctionPlot[74].X = -4.26
	pointsOfFunctionPlot[74].Y = -1.34

	pointsOfFunctionPlot[75].X = -4.25
	pointsOfFunctionPlot[75].Y = -1.339

	pointsOfFunctionPlot[76].X = -4.24
	pointsOfFunctionPlot[76].Y = -1.339

	pointsOfFunctionPlot[77].X = -4.23
	pointsOfFunctionPlot[77].Y = -1.338

	pointsOfFunctionPlot[78].X = -4.22
	pointsOfFunctionPlot[78].Y = -1.338

	pointsOfFunctionPlot[79].X = -4.21
	pointsOfFunctionPlot[79].Y = -1.337

	pointsOfFunctionPlot[80].X = -4.20
	pointsOfFunctionPlot[80].Y = -1.337

	pointsOfFunctionPlot[81].X = -4.19
	pointsOfFunctionPlot[81].Y = -1.336

	pointsOfFunctionPlot[82].X = -4.18
	pointsOfFunctionPlot[82].Y = -1.335

	pointsOfFunctionPlot[83].X = -4.17
	pointsOfFunctionPlot[83].Y = -1.335

	pointsOfFunctionPlot[84].X = -4.16
	pointsOfFunctionPlot[84].Y = -1.334

	pointsOfFunctionPlot[85].X = -4.15
	pointsOfFunctionPlot[85].Y = -1.334

	pointsOfFunctionPlot[86].X = -4.14
	pointsOfFunctionPlot[86].Y = -1.333

	pointsOfFunctionPlot[87].X = -4.13
	pointsOfFunctionPlot[87].Y = -1.333

	pointsOfFunctionPlot[88].X = -4.12
	pointsOfFunctionPlot[88].Y = -1.332

	pointsOfFunctionPlot[89].X = -4.11
	pointsOfFunctionPlot[89].Y = -1.332

	pointsOfFunctionPlot[90].X = -4.10
	pointsOfFunctionPlot[90].Y = -1.331

	pointsOfFunctionPlot[91].X = -4.09
	pointsOfFunctionPlot[91].Y = -1.331

	pointsOfFunctionPlot[92].X = -4.08
	pointsOfFunctionPlot[92].Y = -1.33

	pointsOfFunctionPlot[93].X = -4.07
	pointsOfFunctionPlot[93].Y = -1.329

	pointsOfFunctionPlot[94].X = -4.06
	pointsOfFunctionPlot[94].Y = -1.329

	pointsOfFunctionPlot[95].X = -4.05
	pointsOfFunctionPlot[95].Y = -1.328

	pointsOfFunctionPlot[96].X = -4.04
	pointsOfFunctionPlot[96].Y = -1.328

	pointsOfFunctionPlot[97].X = -4.03
	pointsOfFunctionPlot[97].Y = -1.327

	pointsOfFunctionPlot[98].X = -4.02
	pointsOfFunctionPlot[98].Y = -1.326

	pointsOfFunctionPlot[99].X = -4.01
	pointsOfFunctionPlot[99].Y = -1.326

	pointsOfFunctionPlot[100].X = -4.0
	pointsOfFunctionPlot[100].Y = -1.325

	pointsOfFunctionPlot[101].X = -3.99
	pointsOfFunctionPlot[101].Y = -1.324

	pointsOfFunctionPlot[102].X = -3.98
	pointsOfFunctionPlot[102].Y = -1.324

	pointsOfFunctionPlot[103].X = -3.97
	pointsOfFunctionPlot[103].Y = -1.324

	pointsOfFunctionPlot[104].X = -3.96
	pointsOfFunctionPlot[104].Y = -1.323

	pointsOfFunctionPlot[105].X = -3.95
	pointsOfFunctionPlot[105].Y = -1.322

	pointsOfFunctionPlot[106].X = -3.94
	pointsOfFunctionPlot[106].Y = -1.322

	pointsOfFunctionPlot[107].X = -3.93
	pointsOfFunctionPlot[107].Y = -1.321

	pointsOfFunctionPlot[108].X = -3.92
	pointsOfFunctionPlot[108].Y = -1.321

	pointsOfFunctionPlot[109].X = -3.91
	pointsOfFunctionPlot[109].Y = -1.32

	pointsOfFunctionPlot[110].X = -3.90
	pointsOfFunctionPlot[110].Y = -1.319

	pointsOfFunctionPlot[111].X = -3.89
	pointsOfFunctionPlot[111].Y = -1.319

	pointsOfFunctionPlot[112].X = -3.88
	pointsOfFunctionPlot[112].Y = -1.318

	pointsOfFunctionPlot[113].X = -3.87
	pointsOfFunctionPlot[113].Y = -1.317

	pointsOfFunctionPlot[114].X = -3.86
	pointsOfFunctionPlot[114].Y = -1.317

	pointsOfFunctionPlot[115].X = -3.85
	pointsOfFunctionPlot[115].Y = -1.316

	pointsOfFunctionPlot[116].X = -3.84
	pointsOfFunctionPlot[116].Y = -1.316

	pointsOfFunctionPlot[117].X = -3.83
	pointsOfFunctionPlot[117].Y = -1.315

	pointsOfFunctionPlot[118].X = -3.82
	pointsOfFunctionPlot[118].Y = -1.314

	pointsOfFunctionPlot[119].X = -3.81
	pointsOfFunctionPlot[119].Y = -1.314

	pointsOfFunctionPlot[120].X = -3.80
	pointsOfFunctionPlot[120].Y = -1.313

	pointsOfFunctionPlot[121].X = -3.79
	pointsOfFunctionPlot[121].Y = -1.312

	pointsOfFunctionPlot[122].X = -3.78
	pointsOfFunctionPlot[122].Y = -1.312

	pointsOfFunctionPlot[123].X = -3.77
	pointsOfFunctionPlot[123].Y = -1.311

	pointsOfFunctionPlot[124].X = -3.76
	pointsOfFunctionPlot[124].Y = -1.31

	pointsOfFunctionPlot[125].X = -3.75
	pointsOfFunctionPlot[125].Y = -1.31

	pointsOfFunctionPlot[126].X = -3.74
	pointsOfFunctionPlot[126].Y = -1.309

	pointsOfFunctionPlot[127].X = -3.73
	pointsOfFunctionPlot[127].Y = -1.308

	pointsOfFunctionPlot[128].X = -3.72
	pointsOfFunctionPlot[128].Y = -1.308

	pointsOfFunctionPlot[129].X = -3.71
	pointsOfFunctionPlot[129].Y = -1.307

	pointsOfFunctionPlot[130].X = -3.70
	pointsOfFunctionPlot[130].Y = -1.306

	pointsOfFunctionPlot[131].X = -3.69
	pointsOfFunctionPlot[131].Y = -1.306

	pointsOfFunctionPlot[132].X = -3.68
	pointsOfFunctionPlot[132].Y = -1.305

	pointsOfFunctionPlot[133].X = -3.67
	pointsOfFunctionPlot[133].Y = -1.304

	pointsOfFunctionPlot[134].X = -3.66
	pointsOfFunctionPlot[134].Y = -1.304

	pointsOfFunctionPlot[135].X = -3.65
	pointsOfFunctionPlot[135].Y = -1.303

	pointsOfFunctionPlot[136].X = -3.64
	pointsOfFunctionPlot[136].Y = -1.302

	pointsOfFunctionPlot[137].X = -3.63
	pointsOfFunctionPlot[137].Y = -1.301

	pointsOfFunctionPlot[138].X = -3.62
	pointsOfFunctionPlot[138].Y = -1.301

	pointsOfFunctionPlot[139].X = -3.61
	pointsOfFunctionPlot[139].Y = -1.3

	pointsOfFunctionPlot[140].X = -3.60
	pointsOfFunctionPlot[140].Y = -1.299

	pointsOfFunctionPlot[141].X = -3.59
	pointsOfFunctionPlot[141].Y = -1.299

	pointsOfFunctionPlot[142].X = -3.58
	pointsOfFunctionPlot[142].Y = -1.298

	pointsOfFunctionPlot[143].X = -3.57
	pointsOfFunctionPlot[143].Y = -1.297

	pointsOfFunctionPlot[144].X = -3.56
	pointsOfFunctionPlot[144].Y = -1.296

	pointsOfFunctionPlot[145].X = -3.55
	pointsOfFunctionPlot[145].Y = -1.296

	pointsOfFunctionPlot[146].X = -3.54
	pointsOfFunctionPlot[146].Y = -1.295

	pointsOfFunctionPlot[147].X = -3.53
	pointsOfFunctionPlot[147].Y = -1.294

	pointsOfFunctionPlot[148].X = -3.52
	pointsOfFunctionPlot[148].Y = -1.293

	pointsOfFunctionPlot[149].X = -3.51
	pointsOfFunctionPlot[149].Y = -1.293

	pointsOfFunctionPlot[150].X = -3.50
	pointsOfFunctionPlot[150].Y = -1.292

	pointsOfFunctionPlot[151].X = -3.49
	pointsOfFunctionPlot[151].Y = -1.291

	pointsOfFunctionPlot[152].X = -3.48
	pointsOfFunctionPlot[152].Y = -1.29

	pointsOfFunctionPlot[153].X = -3.47
	pointsOfFunctionPlot[153].Y = -1.29

	pointsOfFunctionPlot[154].X = -3.46
	pointsOfFunctionPlot[154].Y = -1.289

	pointsOfFunctionPlot[155].X = -3.45
	pointsOfFunctionPlot[155].Y = -1.288

	pointsOfFunctionPlot[156].X = -3.44
	pointsOfFunctionPlot[156].Y = -1.287

	pointsOfFunctionPlot[157].X = -3.43
	pointsOfFunctionPlot[157].Y = -1.287

	pointsOfFunctionPlot[158].X = -3.42
	pointsOfFunctionPlot[158].Y = -1.286

	pointsOfFunctionPlot[159].X = -3.41
	pointsOfFunctionPlot[159].Y = -1.285

	pointsOfFunctionPlot[160].X = -3.40
	pointsOfFunctionPlot[160].Y = -1.284

	pointsOfFunctionPlot[161].X = -3.39
	pointsOfFunctionPlot[161].Y = -1.283

	pointsOfFunctionPlot[162].X = -3.38
	pointsOfFunctionPlot[162].Y = -1.283

	pointsOfFunctionPlot[163].X = -3.37
	pointsOfFunctionPlot[163].Y = -1.282

	pointsOfFunctionPlot[164].X = -3.36
	pointsOfFunctionPlot[164].Y = -1.281

	pointsOfFunctionPlot[165].X = -3.35
	pointsOfFunctionPlot[165].Y = -1.28

	pointsOfFunctionPlot[166].X = -3.34
	pointsOfFunctionPlot[166].Y = -1.279

	pointsOfFunctionPlot[167].X = -3.33
	pointsOfFunctionPlot[167].Y = -1.279

	pointsOfFunctionPlot[168].X = -3.32
	pointsOfFunctionPlot[168].Y = -1.278

	pointsOfFunctionPlot[169].X = -3.31
	pointsOfFunctionPlot[169].Y = -1.277

	pointsOfFunctionPlot[170].X = -3.30
	pointsOfFunctionPlot[170].Y = -1.276

	pointsOfFunctionPlot[171].X = -3.29
	pointsOfFunctionPlot[171].Y = -1.275

	pointsOfFunctionPlot[172].X = -3.28
	pointsOfFunctionPlot[172].Y = -1.274

	pointsOfFunctionPlot[173].X = -3.27
	pointsOfFunctionPlot[173].Y = -1.274

	pointsOfFunctionPlot[174].X = -3.26
	pointsOfFunctionPlot[174].Y = -1.273

	pointsOfFunctionPlot[175].X = -3.25
	pointsOfFunctionPlot[175].Y = -1.272

	pointsOfFunctionPlot[176].X = -3.24
	pointsOfFunctionPlot[176].Y = -1.271

	pointsOfFunctionPlot[177].X = -3.23
	pointsOfFunctionPlot[177].Y = -1.27

	pointsOfFunctionPlot[178].X = -3.22
	pointsOfFunctionPlot[178].Y = -1.269

	pointsOfFunctionPlot[179].X = -3.21
	pointsOfFunctionPlot[179].Y = -1.268

	pointsOfFunctionPlot[180].X = -3.20
	pointsOfFunctionPlot[180].Y = -1.267

	pointsOfFunctionPlot[181].X = -3.19
	pointsOfFunctionPlot[181].Y = -1.267

	pointsOfFunctionPlot[182].X = -3.18
	pointsOfFunctionPlot[182].Y = -1.266

	pointsOfFunctionPlot[183].X = -3.17
	pointsOfFunctionPlot[183].Y = -1.265

	pointsOfFunctionPlot[184].X = -3.16
	pointsOfFunctionPlot[184].Y = -1.264

	pointsOfFunctionPlot[185].X = -3.15
	pointsOfFunctionPlot[185].Y = -1.263

	pointsOfFunctionPlot[186].X = -3.14
	pointsOfFunctionPlot[186].Y = -1.262

	pointsOfFunctionPlot[187].X = -3.13
	pointsOfFunctionPlot[187].Y = -1.261

	pointsOfFunctionPlot[188].X = -3.12
	pointsOfFunctionPlot[188].Y = -1.26

	pointsOfFunctionPlot[189].X = -3.11
	pointsOfFunctionPlot[189].Y = -1.259

	pointsOfFunctionPlot[190].X = -3.10
	pointsOfFunctionPlot[190].Y = -1.257

	pointsOfFunctionPlot[191].X = -3.09
	pointsOfFunctionPlot[191].Y = -1.257

	pointsOfFunctionPlot[192].X = -3.08
	pointsOfFunctionPlot[192].Y = -1.256

	pointsOfFunctionPlot[193].X = -3.07
	pointsOfFunctionPlot[193].Y = -1.255

	pointsOfFunctionPlot[194].X = -3.06
	pointsOfFunctionPlot[194].Y = -1.254

	pointsOfFunctionPlot[195].X = -3.05
	pointsOfFunctionPlot[195].Y = -1.253

	pointsOfFunctionPlot[196].X = -3.04
	pointsOfFunctionPlot[196].Y = -1.252

	pointsOfFunctionPlot[197].X = -3.03
	pointsOfFunctionPlot[197].Y = -1.252

	pointsOfFunctionPlot[198].X = -3.02
	pointsOfFunctionPlot[198].Y = -1.251

	pointsOfFunctionPlot[199].X = -3.01
	pointsOfFunctionPlot[199].Y = -1.25

	pointsOfFunctionPlot[200].X = -3.0
	pointsOfFunctionPlot[200].Y = -1.249

	pointsOfFunctionPlot[201].X = -2.99
	pointsOfFunctionPlot[201].Y = -1.248

	pointsOfFunctionPlot[202].X = -2.98
	pointsOfFunctionPlot[202].Y = -1.247

	pointsOfFunctionPlot[203].X = -2.97
	pointsOfFunctionPlot[203].Y = -1.246

	pointsOfFunctionPlot[204].X = -2.96
	pointsOfFunctionPlot[204].Y = -1.244

	pointsOfFunctionPlot[205].X = -2.95
	pointsOfFunctionPlot[205].Y = -1.243

	pointsOfFunctionPlot[206].X = -2.94
	pointsOfFunctionPlot[206].Y = -1.242

	pointsOfFunctionPlot[207].X = -2.93
	pointsOfFunctionPlot[207].Y = -1.241

	pointsOfFunctionPlot[208].X = -2.92
	pointsOfFunctionPlot[208].Y = -1.24

	pointsOfFunctionPlot[209].X = -2.91
	pointsOfFunctionPlot[209].Y = -1.239

	pointsOfFunctionPlot[210].X = -2.90
	pointsOfFunctionPlot[210].Y = -1.238

	pointsOfFunctionPlot[211].X = -2.89
	pointsOfFunctionPlot[211].Y = -1.237

	pointsOfFunctionPlot[212].X = -2.88
	pointsOfFunctionPlot[212].Y = -1.236

	pointsOfFunctionPlot[213].X = -2.87
	pointsOfFunctionPlot[213].Y = -1.235

	pointsOfFunctionPlot[214].X = -2.86
	pointsOfFunctionPlot[214].Y = -1.234

	pointsOfFunctionPlot[215].X = -2.85
	pointsOfFunctionPlot[215].Y = -1.233

	pointsOfFunctionPlot[216].X = -2.84
	pointsOfFunctionPlot[216].Y = -1.232

	pointsOfFunctionPlot[217].X = -2.83
	pointsOfFunctionPlot[217].Y = -1.231

	pointsOfFunctionPlot[218].X = -2.82
	pointsOfFunctionPlot[218].Y = -1.23

	pointsOfFunctionPlot[219].X = -2.81
	pointsOfFunctionPlot[219].Y = -1.228

	pointsOfFunctionPlot[220].X = -2.80
	pointsOfFunctionPlot[220].Y = -1.227

	pointsOfFunctionPlot[221].X = -2.79
	pointsOfFunctionPlot[221].Y = -1.226

	pointsOfFunctionPlot[222].X = -2.78
	pointsOfFunctionPlot[222].Y = -1.225

	pointsOfFunctionPlot[223].X = -2.77
	pointsOfFunctionPlot[223].Y = -1.224

	pointsOfFunctionPlot[224].X = -2.76
	pointsOfFunctionPlot[224].Y = -1.223

	pointsOfFunctionPlot[225].X = -2.75
	pointsOfFunctionPlot[225].Y = -1.222

	pointsOfFunctionPlot[226].X = -2.74
	pointsOfFunctionPlot[226].Y = -1.22

	pointsOfFunctionPlot[227].X = -2.73
	pointsOfFunctionPlot[227].Y = -1.219

	pointsOfFunctionPlot[228].X = -2.72
	pointsOfFunctionPlot[228].Y = -1.218

	pointsOfFunctionPlot[229].X = -2.71
	pointsOfFunctionPlot[229].Y = -1.217

	pointsOfFunctionPlot[230].X = -2.70
	pointsOfFunctionPlot[230].Y = -1.216

	pointsOfFunctionPlot[231].X = -2.69
	pointsOfFunctionPlot[231].Y = -1.214

	pointsOfFunctionPlot[232].X = -2.68
	pointsOfFunctionPlot[232].Y = -1.213

	pointsOfFunctionPlot[233].X = -2.67
	pointsOfFunctionPlot[233].Y = -1.212

	pointsOfFunctionPlot[234].X = -2.66
	pointsOfFunctionPlot[234].Y = -1.211

	pointsOfFunctionPlot[235].X = -2.65
	pointsOfFunctionPlot[235].Y = -1.209

	pointsOfFunctionPlot[236].X = -2.64
	pointsOfFunctionPlot[236].Y = -1.208

	pointsOfFunctionPlot[237].X = -2.63
	pointsOfFunctionPlot[237].Y = -1.207

	pointsOfFunctionPlot[238].X = -2.62
	pointsOfFunctionPlot[238].Y = -1.206

	pointsOfFunctionPlot[239].X = -2.61
	pointsOfFunctionPlot[239].Y = -1.204

	pointsOfFunctionPlot[240].X = -2.60
	pointsOfFunctionPlot[240].Y = -1.203

	pointsOfFunctionPlot[241].X = -2.59
	pointsOfFunctionPlot[241].Y = -1.202

	pointsOfFunctionPlot[242].X = -2.58
	pointsOfFunctionPlot[242].Y = -1.201

	pointsOfFunctionPlot[243].X = -2.57
	pointsOfFunctionPlot[243].Y = -1.199

	pointsOfFunctionPlot[244].X = -2.56
	pointsOfFunctionPlot[244].Y = -1.198

	pointsOfFunctionPlot[245].X = -2.55
	pointsOfFunctionPlot[245].Y = -1.197

	pointsOfFunctionPlot[246].X = -2.54
	pointsOfFunctionPlot[246].Y = -1.195

	pointsOfFunctionPlot[247].X = -2.53
	pointsOfFunctionPlot[247].Y = -1.194

	pointsOfFunctionPlot[248].X = -2.52
	pointsOfFunctionPlot[248].Y = -1.193

	pointsOfFunctionPlot[249].X = -2.51
	pointsOfFunctionPlot[249].Y = -1.191

	pointsOfFunctionPlot[250].X = -2.50
	pointsOfFunctionPlot[250].Y = -1.19

	pointsOfFunctionPlot[251].X = -2.49
	pointsOfFunctionPlot[251].Y = -1.188

	pointsOfFunctionPlot[252].X = -2.48
	pointsOfFunctionPlot[252].Y = -1.187

	pointsOfFunctionPlot[253].X = -2.47
	pointsOfFunctionPlot[253].Y = -1.186

	pointsOfFunctionPlot[254].X = -2.46
	pointsOfFunctionPlot[254].Y = -1.184

	pointsOfFunctionPlot[255].X = -2.45
	pointsOfFunctionPlot[255].Y = -1.183

	pointsOfFunctionPlot[256].X = -2.44
	pointsOfFunctionPlot[256].Y = -1.181

	pointsOfFunctionPlot[257].X = -2.43
	pointsOfFunctionPlot[257].Y = -1.18

	pointsOfFunctionPlot[258].X = -2.42
	pointsOfFunctionPlot[258].Y = -1.179

	pointsOfFunctionPlot[259].X = -2.41
	pointsOfFunctionPlot[259].Y = -1.177

	pointsOfFunctionPlot[260].X = -2.40
	pointsOfFunctionPlot[260].Y = -1.176

	pointsOfFunctionPlot[261].X = -2.39
	pointsOfFunctionPlot[261].Y = -1.174

	pointsOfFunctionPlot[262].X = -2.38
	pointsOfFunctionPlot[262].Y = -1.173

	pointsOfFunctionPlot[263].X = -2.37
	pointsOfFunctionPlot[263].Y = -1.171

	pointsOfFunctionPlot[264].X = -2.36
	pointsOfFunctionPlot[264].Y = -1.17

	pointsOfFunctionPlot[265].X = -2.35
	pointsOfFunctionPlot[265].Y = -1.168

	pointsOfFunctionPlot[266].X = -2.34
	pointsOfFunctionPlot[266].Y = -1.166

	pointsOfFunctionPlot[267].X = -2.33
	pointsOfFunctionPlot[267].Y = -1.165

	pointsOfFunctionPlot[268].X = -2.32
	pointsOfFunctionPlot[268].Y = -1.163

	pointsOfFunctionPlot[269].X = -2.31
	pointsOfFunctionPlot[269].Y = -1.162

	pointsOfFunctionPlot[270].X = -2.30
	pointsOfFunctionPlot[270].Y = -1.16

	pointsOfFunctionPlot[271].X = -2.29
	pointsOfFunctionPlot[271].Y = -1.159

	pointsOfFunctionPlot[272].X = -2.28
	pointsOfFunctionPlot[272].Y = -1.157

	pointsOfFunctionPlot[273].X = -2.27
	pointsOfFunctionPlot[273].Y = -1.155

	pointsOfFunctionPlot[274].X = -2.26
	pointsOfFunctionPlot[274].Y = -1.154

	pointsOfFunctionPlot[275].X = -2.25
	pointsOfFunctionPlot[275].Y = -1.152

	pointsOfFunctionPlot[276].X = -2.24
	pointsOfFunctionPlot[276].Y = -1.15

	pointsOfFunctionPlot[277].X = -2.23
	pointsOfFunctionPlot[277].Y = -1.149

	pointsOfFunctionPlot[278].X = -2.22
	pointsOfFunctionPlot[278].Y = -1.147

	pointsOfFunctionPlot[279].X = -2.21
	pointsOfFunctionPlot[279].Y = -1.145

	pointsOfFunctionPlot[280].X = -2.20
	pointsOfFunctionPlot[280].Y = -1.144

	pointsOfFunctionPlot[281].X = -2.19
	pointsOfFunctionPlot[281].Y = -1.142

	pointsOfFunctionPlot[282].X = -2.18
	pointsOfFunctionPlot[282].Y = -1.14

	pointsOfFunctionPlot[283].X = -2.17
	pointsOfFunctionPlot[283].Y = -1.138

	pointsOfFunctionPlot[284].X = -2.16
	pointsOfFunctionPlot[284].Y = -1.137

	pointsOfFunctionPlot[285].X = -2.15
	pointsOfFunctionPlot[285].Y = -1.135

	pointsOfFunctionPlot[286].X = -2.14
	pointsOfFunctionPlot[286].Y = -1.133

	pointsOfFunctionPlot[287].X = -2.13
	pointsOfFunctionPlot[287].Y = -1.131

	pointsOfFunctionPlot[288].X = -2.12
	pointsOfFunctionPlot[288].Y = -1.13

	pointsOfFunctionPlot[289].X = -2.11
	pointsOfFunctionPlot[289].Y = -1.128

	pointsOfFunctionPlot[290].X = -2.10
	pointsOfFunctionPlot[290].Y = -1.126

	pointsOfFunctionPlot[291].X = -2.09
	pointsOfFunctionPlot[291].Y = -1.124

	pointsOfFunctionPlot[292].X = -2.08
	pointsOfFunctionPlot[292].Y = -1.122

	pointsOfFunctionPlot[293].X = -2.07
	pointsOfFunctionPlot[293].Y = -1.12

	pointsOfFunctionPlot[294].X = -2.06
	pointsOfFunctionPlot[294].Y = -1.118

	pointsOfFunctionPlot[295].X = -2.05
	pointsOfFunctionPlot[295].Y = -1.116

	pointsOfFunctionPlot[296].X = -2.04
	pointsOfFunctionPlot[296].Y = -1.115

	pointsOfFunctionPlot[297].X = -2.03
	pointsOfFunctionPlot[297].Y = -1.113

	pointsOfFunctionPlot[298].X = -2.02
	pointsOfFunctionPlot[298].Y = -1.111

	pointsOfFunctionPlot[299].X = -2.01
	pointsOfFunctionPlot[299].Y = -1.109

	pointsOfFunctionPlot[300].X = -2.0
	pointsOfFunctionPlot[300].Y = -1.107

	pointsOfFunctionPlot[301].X = -1.99
	pointsOfFunctionPlot[301].Y = -1.105

	pointsOfFunctionPlot[302].X = -1.98
	pointsOfFunctionPlot[302].Y = -1.103

	pointsOfFunctionPlot[303].X = -1.97
	pointsOfFunctionPlot[303].Y = -1.101

	pointsOfFunctionPlot[304].X = -1.96
	pointsOfFunctionPlot[304].Y = -1.099

	pointsOfFunctionPlot[305].X = -1.95
	pointsOfFunctionPlot[305].Y = -1.096

	pointsOfFunctionPlot[306].X = -1.94
	pointsOfFunctionPlot[306].Y = -1.094

	pointsOfFunctionPlot[307].X = -1.93
	pointsOfFunctionPlot[307].Y = -1.092

	pointsOfFunctionPlot[308].X = -1.92
	pointsOfFunctionPlot[308].Y = -1.09

	pointsOfFunctionPlot[309].X = -1.91
	pointsOfFunctionPlot[309].Y = -1.088

	pointsOfFunctionPlot[310].X = -1.90
	pointsOfFunctionPlot[310].Y = -1.086

	pointsOfFunctionPlot[311].X = -1.89
	pointsOfFunctionPlot[311].Y = -1.084

	pointsOfFunctionPlot[312].X = -1.88
	pointsOfFunctionPlot[312].Y = -1.081

	pointsOfFunctionPlot[313].X = -1.87
	pointsOfFunctionPlot[313].Y = -1.079

	pointsOfFunctionPlot[314].X = -1.86
	pointsOfFunctionPlot[314].Y = -1.077

	pointsOfFunctionPlot[315].X = -1.85
	pointsOfFunctionPlot[315].Y = -1.075

	pointsOfFunctionPlot[316].X = -1.84
	pointsOfFunctionPlot[316].Y = -1.072

	pointsOfFunctionPlot[317].X = -1.83
	pointsOfFunctionPlot[317].Y = -1.07

	pointsOfFunctionPlot[318].X = -1.82
	pointsOfFunctionPlot[318].Y = -1.068

	pointsOfFunctionPlot[319].X = -1.81
	pointsOfFunctionPlot[319].Y = -1.066

	pointsOfFunctionPlot[320].X = -1.80
	pointsOfFunctionPlot[320].Y = -1.063

	pointsOfFunctionPlot[321].X = -1.79
	pointsOfFunctionPlot[321].Y = -1.061

	pointsOfFunctionPlot[322].X = -1.78
	pointsOfFunctionPlot[322].Y = -1.058

	pointsOfFunctionPlot[323].X = -1.77
	pointsOfFunctionPlot[323].Y = -1.056

	pointsOfFunctionPlot[324].X = -1.76
	pointsOfFunctionPlot[324].Y = -1.054

	pointsOfFunctionPlot[325].X = -1.75
	pointsOfFunctionPlot[325].Y = -1.051

	pointsOfFunctionPlot[326].X = -1.74
	pointsOfFunctionPlot[326].Y = -1.049

	pointsOfFunctionPlot[327].X = -1.73
	pointsOfFunctionPlot[327].Y = -1.046

	pointsOfFunctionPlot[328].X = -1.72
	pointsOfFunctionPlot[328].Y = -1.044

	pointsOfFunctionPlot[329].X = -1.71
	pointsOfFunctionPlot[329].Y = -1.044

	pointsOfFunctionPlot[330].X = -1.70
	pointsOfFunctionPlot[330].Y = -1.041

	pointsOfFunctionPlot[331].X = -1.69
	pointsOfFunctionPlot[331].Y = -1.036

	pointsOfFunctionPlot[332].X = -1.68
	pointsOfFunctionPlot[332].Y = -1.033

	pointsOfFunctionPlot[333].X = -1.67
	pointsOfFunctionPlot[333].Y = -1.031

	pointsOfFunctionPlot[334].X = -1.66
	pointsOfFunctionPlot[334].Y = -1.028

	pointsOfFunctionPlot[335].X = -1.65
	pointsOfFunctionPlot[335].Y = -1.025

	pointsOfFunctionPlot[336].X = -1.64
	pointsOfFunctionPlot[336].Y = -1.023

	pointsOfFunctionPlot[337].X = -1.63
	pointsOfFunctionPlot[337].Y = -1.02

	pointsOfFunctionPlot[338].X = -1.62
	pointsOfFunctionPlot[338].Y = -1.017

	pointsOfFunctionPlot[339].X = -1.61
	pointsOfFunctionPlot[339].Y = -1.014

	pointsOfFunctionPlot[340].X = -1.60
	pointsOfFunctionPlot[340].Y = -1.012

	pointsOfFunctionPlot[341].X = -1.59
	pointsOfFunctionPlot[341].Y = -1.009

	pointsOfFunctionPlot[342].X = -1.58
	pointsOfFunctionPlot[342].Y = -1.006

	pointsOfFunctionPlot[343].X = -1.57
	pointsOfFunctionPlot[343].Y = -1.003

	pointsOfFunctionPlot[344].X = -1.56
	pointsOfFunctionPlot[344].Y = -1.0

	pointsOfFunctionPlot[345].X = -1.55
	pointsOfFunctionPlot[345].Y = -0.997

	pointsOfFunctionPlot[346].X = -1.54
	pointsOfFunctionPlot[346].Y = -0.994

	pointsOfFunctionPlot[347].X = -1.53
	pointsOfFunctionPlot[347].Y = -0.991

	pointsOfFunctionPlot[348].X = -1.52
	pointsOfFunctionPlot[348].Y = -0.988

	pointsOfFunctionPlot[349].X = -1.51
	pointsOfFunctionPlot[349].Y = -0.985

	pointsOfFunctionPlot[350].X = -1.50
	pointsOfFunctionPlot[350].Y = -0.982

	pointsOfFunctionPlot[351].X = -1.49
	pointsOfFunctionPlot[351].Y = -0.979

	pointsOfFunctionPlot[352].X = -1.48
	pointsOfFunctionPlot[352].Y = -0.976

	pointsOfFunctionPlot[353].X = -1.47
	pointsOfFunctionPlot[353].Y = -0.973

	pointsOfFunctionPlot[354].X = -1.46
	pointsOfFunctionPlot[354].Y = -0.97

	pointsOfFunctionPlot[355].X = -1.45
	pointsOfFunctionPlot[355].Y = -0.967

	pointsOfFunctionPlot[356].X = -1.44
	pointsOfFunctionPlot[356].Y = -0.963

	pointsOfFunctionPlot[357].X = -1.43
	pointsOfFunctionPlot[357].Y = -0.96

	pointsOfFunctionPlot[358].X = -1.42
	pointsOfFunctionPlot[358].Y = -0.957

	pointsOfFunctionPlot[359].X = -1.41
	pointsOfFunctionPlot[359].Y = -0.953

	pointsOfFunctionPlot[360].X = -1.40
	pointsOfFunctionPlot[360].Y = -0.95

	pointsOfFunctionPlot[361].X = -1.39
	pointsOfFunctionPlot[361].Y = -0.947

	pointsOfFunctionPlot[362].X = -1.38
	pointsOfFunctionPlot[362].Y = -0.943

	pointsOfFunctionPlot[363].X = -1.37
	pointsOfFunctionPlot[363].Y = -0.94

	pointsOfFunctionPlot[364].X = -1.36
	pointsOfFunctionPlot[364].Y = -0.936

	pointsOfFunctionPlot[365].X = -1.35
	pointsOfFunctionPlot[365].Y = -0.933

	pointsOfFunctionPlot[366].X = -1.34
	pointsOfFunctionPlot[366].Y = -0.929

	pointsOfFunctionPlot[367].X = -1.33
	pointsOfFunctionPlot[367].Y = -0.926

	pointsOfFunctionPlot[368].X = -1.32
	pointsOfFunctionPlot[368].Y = -0.922

	pointsOfFunctionPlot[369].X = -1.31
	pointsOfFunctionPlot[369].Y = -0.918

	pointsOfFunctionPlot[370].X = -1.30
	pointsOfFunctionPlot[370].Y = -0.915

	pointsOfFunctionPlot[371].X = -1.29
	pointsOfFunctionPlot[371].Y = -0.911

	pointsOfFunctionPlot[372].X = -1.28
	pointsOfFunctionPlot[372].Y = -0.907

	pointsOfFunctionPlot[373].X = -1.27
	pointsOfFunctionPlot[373].Y = -0.903

	pointsOfFunctionPlot[374].X = -1.26
	pointsOfFunctionPlot[374].Y = -0.899

	pointsOfFunctionPlot[375].X = -1.25
	pointsOfFunctionPlot[375].Y = -0.896

	pointsOfFunctionPlot[376].X = -1.24
	pointsOfFunctionPlot[376].Y = -0.892

	pointsOfFunctionPlot[377].X = -1.23
	pointsOfFunctionPlot[377].Y = -0.888

	pointsOfFunctionPlot[378].X = -1.22
	pointsOfFunctionPlot[378].Y = -0.884

	pointsOfFunctionPlot[379].X = -1.21
	pointsOfFunctionPlot[379].Y = -0.88

	pointsOfFunctionPlot[380].X = -1.20
	pointsOfFunctionPlot[380].Y = -0.876

	pointsOfFunctionPlot[381].X = -1.19
	pointsOfFunctionPlot[381].Y = -0.871

	pointsOfFunctionPlot[382].X = -1.18
	pointsOfFunctionPlot[382].Y = -0.867

	pointsOfFunctionPlot[383].X = -1.17
	pointsOfFunctionPlot[383].Y = -0.863

	pointsOfFunctionPlot[384].X = -1.16
	pointsOfFunctionPlot[384].Y = -0.859

	pointsOfFunctionPlot[385].X = -1.15
	pointsOfFunctionPlot[385].Y = -0.855

	pointsOfFunctionPlot[386].X = -1.14
	pointsOfFunctionPlot[386].Y = -0.85

	pointsOfFunctionPlot[387].X = -1.13
	pointsOfFunctionPlot[387].Y = -0.846

	pointsOfFunctionPlot[388].X = -1.12
	pointsOfFunctionPlot[388].Y = -0.841

	pointsOfFunctionPlot[389].X = -1.11
	pointsOfFunctionPlot[389].Y = -0.837

	pointsOfFunctionPlot[390].X = -1.10
	pointsOfFunctionPlot[390].Y = -0.832

	pointsOfFunctionPlot[391].X = -1.09
	pointsOfFunctionPlot[391].Y = -0.828

	pointsOfFunctionPlot[392].X = -1.08
	pointsOfFunctionPlot[392].Y = -0.823

	pointsOfFunctionPlot[393].X = -1.07
	pointsOfFunctionPlot[393].Y = -0.819

	pointsOfFunctionPlot[394].X = -1.06
	pointsOfFunctionPlot[394].Y = -0.814

	pointsOfFunctionPlot[395].X = -1.05
	pointsOfFunctionPlot[395].Y = -0.809

	pointsOfFunctionPlot[396].X = -1.04
	pointsOfFunctionPlot[396].Y = -0.805

	pointsOfFunctionPlot[397].X = -1.03
	pointsOfFunctionPlot[397].Y = -0.8

	pointsOfFunctionPlot[398].X = -1.02
	pointsOfFunctionPlot[398].Y = -0.795

	pointsOfFunctionPlot[399].X = -1.01
	pointsOfFunctionPlot[399].Y = -0.79

	pointsOfFunctionPlot[400].X = -1.0
	pointsOfFunctionPlot[400].Y = -0.785

	pointsOfFunctionPlot[401].X = -0.99
	pointsOfFunctionPlot[401].Y = -0.78

	pointsOfFunctionPlot[402].X = -0.98
	pointsOfFunctionPlot[402].Y = -0.775

	pointsOfFunctionPlot[403].X = -0.97
	pointsOfFunctionPlot[403].Y = -0.77

	pointsOfFunctionPlot[404].X = -0.96
	pointsOfFunctionPlot[404].Y = -0.764

	pointsOfFunctionPlot[405].X = -0.95
	pointsOfFunctionPlot[405].Y = -0.759

	pointsOfFunctionPlot[406].X = -0.94
	pointsOfFunctionPlot[406].Y = -0.754

	pointsOfFunctionPlot[407].X = -0.93
	pointsOfFunctionPlot[407].Y = -0.749

	pointsOfFunctionPlot[408].X = -0.92
	pointsOfFunctionPlot[408].Y = -0.743

	pointsOfFunctionPlot[409].X = -0.91
	pointsOfFunctionPlot[409].Y = -0.738

	pointsOfFunctionPlot[410].X = -0.90
	pointsOfFunctionPlot[410].Y = -0.732

	pointsOfFunctionPlot[411].X = -0.89
	pointsOfFunctionPlot[411].Y = -0.727

	pointsOfFunctionPlot[412].X = -0.88
	pointsOfFunctionPlot[412].Y = -0.721

	pointsOfFunctionPlot[413].X = -0.87
	pointsOfFunctionPlot[413].Y = -0.715

	pointsOfFunctionPlot[414].X = -0.86
	pointsOfFunctionPlot[414].Y = -0.71

	pointsOfFunctionPlot[415].X = -0.85
	pointsOfFunctionPlot[415].Y = -0.704

	pointsOfFunctionPlot[416].X = -0.84
	pointsOfFunctionPlot[416].Y = -0.698

	pointsOfFunctionPlot[417].X = -0.83
	pointsOfFunctionPlot[417].Y = -0.692

	pointsOfFunctionPlot[418].X = -0.82
	pointsOfFunctionPlot[418].Y = -0.686

	pointsOfFunctionPlot[419].X = -0.81
	pointsOfFunctionPlot[419].Y = -0.68

	pointsOfFunctionPlot[420].X = -0.80
	pointsOfFunctionPlot[420].Y = -0.674

	pointsOfFunctionPlot[421].X = -0.79
	pointsOfFunctionPlot[421].Y = -0.668

	pointsOfFunctionPlot[422].X = -0.78
	pointsOfFunctionPlot[422].Y = -0.662

	pointsOfFunctionPlot[423].X = -0.77
	pointsOfFunctionPlot[423].Y = -0.656

	pointsOfFunctionPlot[424].X = -0.76
	pointsOfFunctionPlot[424].Y = -0.649

	pointsOfFunctionPlot[425].X = -0.75
	pointsOfFunctionPlot[425].Y = -0.643

	pointsOfFunctionPlot[426].X = -0.74
	pointsOfFunctionPlot[426].Y = -0.637

	pointsOfFunctionPlot[427].X = -0.73
	pointsOfFunctionPlot[427].Y = -0.63

	pointsOfFunctionPlot[428].X = -0.72
	pointsOfFunctionPlot[428].Y = -0.624

	pointsOfFunctionPlot[429].X = -0.71
	pointsOfFunctionPlot[429].Y = -0.617

	pointsOfFunctionPlot[430].X = -0.70
	pointsOfFunctionPlot[430].Y = -0.61

	pointsOfFunctionPlot[431].X = -0.69
	pointsOfFunctionPlot[431].Y = -0.603

	pointsOfFunctionPlot[432].X = -0.68
	pointsOfFunctionPlot[432].Y = -0.597

	pointsOfFunctionPlot[433].X = -0.67
	pointsOfFunctionPlot[433].Y = -0.59

	pointsOfFunctionPlot[434].X = -0.66
	pointsOfFunctionPlot[434].Y = -0.583

	pointsOfFunctionPlot[435].X = -0.65
	pointsOfFunctionPlot[435].Y = -0.576

	pointsOfFunctionPlot[436].X = -0.64
	pointsOfFunctionPlot[436].Y = -0.569

	pointsOfFunctionPlot[437].X = -0.63
	pointsOfFunctionPlot[437].Y = -0.562

	pointsOfFunctionPlot[438].X = -0.62
	pointsOfFunctionPlot[438].Y = -0.554

	pointsOfFunctionPlot[439].X = -0.61
	pointsOfFunctionPlot[439].Y = -0.547

	pointsOfFunctionPlot[440].X = -0.60
	pointsOfFunctionPlot[440].Y = -0.54

	pointsOfFunctionPlot[441].X = -0.59
	pointsOfFunctionPlot[441].Y = -0.533

	pointsOfFunctionPlot[442].X = -0.58
	pointsOfFunctionPlot[442].Y = -0.525

	pointsOfFunctionPlot[443].X = -0.57
	pointsOfFunctionPlot[443].Y = -0.518

	pointsOfFunctionPlot[444].X = -0.56
	pointsOfFunctionPlot[444].Y = -0.51

	pointsOfFunctionPlot[445].X = -0.55
	pointsOfFunctionPlot[445].Y = -0.502

	pointsOfFunctionPlot[446].X = -0.54
	pointsOfFunctionPlot[446].Y = -0.495

	pointsOfFunctionPlot[447].X = -0.53
	pointsOfFunctionPlot[447].Y = -0.487

	pointsOfFunctionPlot[448].X = -0.52
	pointsOfFunctionPlot[448].Y = -0.479

	pointsOfFunctionPlot[449].X = -0.51
	pointsOfFunctionPlot[449].Y = -0.471

	pointsOfFunctionPlot[450].X = -0.50
	pointsOfFunctionPlot[450].Y = -0.463

	pointsOfFunctionPlot[451].X = -0.49
	pointsOfFunctionPlot[451].Y = -0.455

	pointsOfFunctionPlot[452].X = -0.48
	pointsOfFunctionPlot[452].Y = -0.447

	pointsOfFunctionPlot[453].X = -0.47
	pointsOfFunctionPlot[453].Y = -0.439

	pointsOfFunctionPlot[454].X = -0.46
	pointsOfFunctionPlot[454].Y = -0.431

	pointsOfFunctionPlot[455].X = -0.45
	pointsOfFunctionPlot[455].Y = -0.422

	pointsOfFunctionPlot[456].X = -0.44
	pointsOfFunctionPlot[456].Y = -0.414

	pointsOfFunctionPlot[457].X = -0.43
	pointsOfFunctionPlot[457].Y = -0.406

	pointsOfFunctionPlot[458].X = -0.42
	pointsOfFunctionPlot[458].Y = -0.397

	pointsOfFunctionPlot[459].X = -0.41
	pointsOfFunctionPlot[459].Y = -0.389

	pointsOfFunctionPlot[460].X = -0.40
	pointsOfFunctionPlot[460].Y = -0.38

	pointsOfFunctionPlot[461].X = -0.39
	pointsOfFunctionPlot[461].Y = -0.371

	pointsOfFunctionPlot[462].X = -0.38
	pointsOfFunctionPlot[462].Y = -0.363

	pointsOfFunctionPlot[463].X = -0.37
	pointsOfFunctionPlot[463].Y = -0.354

	pointsOfFunctionPlot[464].X = -0.36
	pointsOfFunctionPlot[464].Y = -0.345

	pointsOfFunctionPlot[465].X = -0.35
	pointsOfFunctionPlot[465].Y = -0.336

	pointsOfFunctionPlot[466].X = -0.34
	pointsOfFunctionPlot[466].Y = -0.327

	pointsOfFunctionPlot[467].X = -0.33
	pointsOfFunctionPlot[467].Y = -0.318

	pointsOfFunctionPlot[468].X = -0.32
	pointsOfFunctionPlot[468].Y = -0.309

	pointsOfFunctionPlot[469].X = -0.31
	pointsOfFunctionPlot[469].Y = -0.3

	pointsOfFunctionPlot[470].X = -0.30
	pointsOfFunctionPlot[470].Y = -0.291

	pointsOfFunctionPlot[471].X = -0.29
	pointsOfFunctionPlot[471].Y = -0.282

	pointsOfFunctionPlot[472].X = -0.28
	pointsOfFunctionPlot[472].Y = -0.273

	pointsOfFunctionPlot[473].X = -0.27
	pointsOfFunctionPlot[473].Y = -0.263

	pointsOfFunctionPlot[474].X = -0.26
	pointsOfFunctionPlot[474].Y = -0.254

	pointsOfFunctionPlot[475].X = -0.25
	pointsOfFunctionPlot[475].Y = -0.244

	pointsOfFunctionPlot[476].X = -0.24
	pointsOfFunctionPlot[476].Y = -0.235

	pointsOfFunctionPlot[477].X = -0.23
	pointsOfFunctionPlot[477].Y = -0.226

	pointsOfFunctionPlot[478].X = -0.22
	pointsOfFunctionPlot[478].Y = -0.216

	pointsOfFunctionPlot[479].X = -0.21
	pointsOfFunctionPlot[479].Y = -0.206

	pointsOfFunctionPlot[480].X = -0.20
	pointsOfFunctionPlot[480].Y = -0.197

	pointsOfFunctionPlot[481].X = -0.19
	pointsOfFunctionPlot[481].Y = -0.187

	pointsOfFunctionPlot[482].X = -0.18
	pointsOfFunctionPlot[482].Y = -0.178

	pointsOfFunctionPlot[483].X = -0.17
	pointsOfFunctionPlot[483].Y = -0.168

	pointsOfFunctionPlot[484].X = -0.16
	pointsOfFunctionPlot[484].Y = -0.158

	pointsOfFunctionPlot[485].X = -0.15
	pointsOfFunctionPlot[485].Y = -0.148

	pointsOfFunctionPlot[486].X = -0.14
	pointsOfFunctionPlot[486].Y = -0.139

	pointsOfFunctionPlot[487].X = -0.13
	pointsOfFunctionPlot[487].Y = -0.129

	pointsOfFunctionPlot[488].X = -0.12
	pointsOfFunctionPlot[488].Y = -0.119

	pointsOfFunctionPlot[489].X = -0.11
	pointsOfFunctionPlot[489].Y = -0.109

	pointsOfFunctionPlot[490].X = -0.10
	pointsOfFunctionPlot[490].Y = -0.099

	pointsOfFunctionPlot[491].X = -0.09
	pointsOfFunctionPlot[491].Y = -0.089

	pointsOfFunctionPlot[492].X = -0.08
	pointsOfFunctionPlot[492].Y = -0.079

	pointsOfFunctionPlot[493].X = -0.07
	pointsOfFunctionPlot[493].Y = -0.069

	pointsOfFunctionPlot[494].X = -0.06
	pointsOfFunctionPlot[494].Y = -0.059

	pointsOfFunctionPlot[495].X = -0.05
	pointsOfFunctionPlot[495].Y = -0.049

	pointsOfFunctionPlot[496].X = -0.04
	pointsOfFunctionPlot[496].Y = -0.039

	pointsOfFunctionPlot[497].X = -0.03
	pointsOfFunctionPlot[497].Y = -0.029

	pointsOfFunctionPlot[498].X = -0.02
	pointsOfFunctionPlot[498].Y = -0.019

	pointsOfFunctionPlot[499].X = -0.01
	pointsOfFunctionPlot[499].Y = -0.009

	pointsOfFunctionPlot[500].X = 0.0
	pointsOfFunctionPlot[500].Y = 0.0

	pointsOfFunctionPlot[501].X = 0.01
	pointsOfFunctionPlot[501].Y = 0.009

	pointsOfFunctionPlot[502].X = 0.02
	pointsOfFunctionPlot[502].Y = 0.019

	pointsOfFunctionPlot[503].X = 0.03
	pointsOfFunctionPlot[503].Y = 0.029

	pointsOfFunctionPlot[504].X = 0.04
	pointsOfFunctionPlot[504].Y = 0.039

	pointsOfFunctionPlot[505].X = 0.05
	pointsOfFunctionPlot[505].Y = 0.049

	pointsOfFunctionPlot[506].X = 0.06
	pointsOfFunctionPlot[506].Y = 0.059

	pointsOfFunctionPlot[507].X = 0.07
	pointsOfFunctionPlot[507].Y = 0.069

	pointsOfFunctionPlot[508].X = 0.08
	pointsOfFunctionPlot[508].Y = 0.079

	pointsOfFunctionPlot[509].X = 0.09
	pointsOfFunctionPlot[509].Y = 0.089

	pointsOfFunctionPlot[510].X = 0.10
	pointsOfFunctionPlot[510].Y = 0.099

	pointsOfFunctionPlot[511].X = 0.11
	pointsOfFunctionPlot[511].Y = 0.109

	pointsOfFunctionPlot[512].X = 0.12
	pointsOfFunctionPlot[512].Y = 0.119

	pointsOfFunctionPlot[513].X = 0.13
	pointsOfFunctionPlot[513].Y = 0.129

	pointsOfFunctionPlot[514].X = 0.14
	pointsOfFunctionPlot[514].Y = 0.139

	pointsOfFunctionPlot[515].X = 0.15
	pointsOfFunctionPlot[515].Y = 0.148

	pointsOfFunctionPlot[516].X = 0.16
	pointsOfFunctionPlot[516].Y = 0.158

	pointsOfFunctionPlot[517].X = 0.17
	pointsOfFunctionPlot[517].Y = 0.168

	pointsOfFunctionPlot[518].X = 0.18
	pointsOfFunctionPlot[518].Y = 0.178

	pointsOfFunctionPlot[519].X = 0.19
	pointsOfFunctionPlot[519].Y = 0.187

	pointsOfFunctionPlot[520].X = 0.20
	pointsOfFunctionPlot[520].Y = 0.197

	pointsOfFunctionPlot[521].X = 0.21
	pointsOfFunctionPlot[521].Y = 0.206

	pointsOfFunctionPlot[522].X = 0.22
	pointsOfFunctionPlot[522].Y = 0.216

	pointsOfFunctionPlot[523].X = 0.23
	pointsOfFunctionPlot[523].Y = 0.226

	pointsOfFunctionPlot[524].X = 0.24
	pointsOfFunctionPlot[524].Y = 0.235

	pointsOfFunctionPlot[525].X = 0.25
	pointsOfFunctionPlot[525].Y = 0.244

	pointsOfFunctionPlot[526].X = 0.26
	pointsOfFunctionPlot[526].Y = 0.254

	pointsOfFunctionPlot[527].X = 0.27
	pointsOfFunctionPlot[527].Y = 0.263

	pointsOfFunctionPlot[528].X = 0.28
	pointsOfFunctionPlot[528].Y = 0.273

	pointsOfFunctionPlot[529].X = 0.29
	pointsOfFunctionPlot[529].Y = 0.282

	pointsOfFunctionPlot[530].X = 0.30
	pointsOfFunctionPlot[530].Y = 0.291

	pointsOfFunctionPlot[531].X = 0.31
	pointsOfFunctionPlot[531].Y = 0.3

	pointsOfFunctionPlot[532].X = 0.32
	pointsOfFunctionPlot[532].Y = 0.309

	pointsOfFunctionPlot[533].X = 0.33
	pointsOfFunctionPlot[533].Y = 0.318

	pointsOfFunctionPlot[534].X = 0.34
	pointsOfFunctionPlot[534].Y = 0.327

	pointsOfFunctionPlot[535].X = 0.35
	pointsOfFunctionPlot[535].Y = 0.336

	pointsOfFunctionPlot[536].X = 0.36
	pointsOfFunctionPlot[536].Y = 0.345

	pointsOfFunctionPlot[537].X = 0.37
	pointsOfFunctionPlot[537].Y = 0.354

	pointsOfFunctionPlot[538].X = 0.38
	pointsOfFunctionPlot[538].Y = 0.363

	pointsOfFunctionPlot[539].X = 0.39
	pointsOfFunctionPlot[539].Y = 0.371

	pointsOfFunctionPlot[540].X = 0.40
	pointsOfFunctionPlot[540].Y = 0.38

	pointsOfFunctionPlot[541].X = 0.41
	pointsOfFunctionPlot[541].Y = 0.389

	pointsOfFunctionPlot[542].X = 0.42
	pointsOfFunctionPlot[542].Y = 0.397

	pointsOfFunctionPlot[543].X = 0.43
	pointsOfFunctionPlot[543].Y = 0.406

	pointsOfFunctionPlot[544].X = 0.44
	pointsOfFunctionPlot[544].Y = 0.414

	pointsOfFunctionPlot[545].X = 0.45
	pointsOfFunctionPlot[545].Y = 0.422

	pointsOfFunctionPlot[546].X = 0.46
	pointsOfFunctionPlot[546].Y = 0.431

	pointsOfFunctionPlot[547].X = 0.47
	pointsOfFunctionPlot[547].Y = 0.439

	pointsOfFunctionPlot[548].X = 0.48
	pointsOfFunctionPlot[548].Y = 0.447

	pointsOfFunctionPlot[549].X = 0.49
	pointsOfFunctionPlot[549].Y = 0.455

	pointsOfFunctionPlot[550].X = 0.50
	pointsOfFunctionPlot[550].Y = 0.463

	pointsOfFunctionPlot[551].X = 0.51
	pointsOfFunctionPlot[551].Y = 0.471

	pointsOfFunctionPlot[552].X = 0.52
	pointsOfFunctionPlot[552].Y = 0.479

	pointsOfFunctionPlot[553].X = 0.53
	pointsOfFunctionPlot[553].Y = 0.487

	pointsOfFunctionPlot[554].X = 0.54
	pointsOfFunctionPlot[554].Y = 0.495

	pointsOfFunctionPlot[555].X = 0.55
	pointsOfFunctionPlot[555].Y = 0.502

	pointsOfFunctionPlot[556].X = 0.56
	pointsOfFunctionPlot[556].Y = 0.51

	pointsOfFunctionPlot[557].X = 0.57
	pointsOfFunctionPlot[557].Y = 0.518

	pointsOfFunctionPlot[558].X = 0.58
	pointsOfFunctionPlot[558].Y = 0.525

	pointsOfFunctionPlot[559].X = 0.59
	pointsOfFunctionPlot[559].Y = 0.533

	pointsOfFunctionPlot[560].X = 0.60
	pointsOfFunctionPlot[560].Y = 0.54

	pointsOfFunctionPlot[561].X = 0.61
	pointsOfFunctionPlot[561].Y = 0.547

	pointsOfFunctionPlot[562].X = 0.62
	pointsOfFunctionPlot[562].Y = 0.554

	pointsOfFunctionPlot[563].X = 0.63
	pointsOfFunctionPlot[563].Y = 0.562

	pointsOfFunctionPlot[564].X = 0.64
	pointsOfFunctionPlot[564].Y = 0.569

	pointsOfFunctionPlot[565].X = 0.65
	pointsOfFunctionPlot[565].Y = 0.576

	pointsOfFunctionPlot[566].X = 0.66
	pointsOfFunctionPlot[566].Y = 0.583

	pointsOfFunctionPlot[567].X = 0.67
	pointsOfFunctionPlot[567].Y = 0.59

	pointsOfFunctionPlot[568].X = 0.68
	pointsOfFunctionPlot[568].Y = 0.597

	pointsOfFunctionPlot[569].X = 0.69
	pointsOfFunctionPlot[569].Y = 0.603

	pointsOfFunctionPlot[570].X = 0.70
	pointsOfFunctionPlot[570].Y = 0.61

	pointsOfFunctionPlot[571].X = 0.71
	pointsOfFunctionPlot[571].Y = 0.617

	pointsOfFunctionPlot[572].X = 0.72
	pointsOfFunctionPlot[572].Y = 0.624

	pointsOfFunctionPlot[573].X = 0.73
	pointsOfFunctionPlot[573].Y = 0.63

	pointsOfFunctionPlot[574].X = 0.74
	pointsOfFunctionPlot[574].Y = 0.637

	pointsOfFunctionPlot[575].X = 0.75
	pointsOfFunctionPlot[575].Y = 0.643

	pointsOfFunctionPlot[576].X = 0.76
	pointsOfFunctionPlot[576].Y = 0.649

	pointsOfFunctionPlot[577].X = 0.77
	pointsOfFunctionPlot[577].Y = 0.656

	pointsOfFunctionPlot[578].X = 0.78
	pointsOfFunctionPlot[578].Y = 0.662

	pointsOfFunctionPlot[579].X = 0.79
	pointsOfFunctionPlot[579].Y = 0.668

	pointsOfFunctionPlot[580].X = 0.80
	pointsOfFunctionPlot[580].Y = 0.674

	pointsOfFunctionPlot[581].X = 0.81
	pointsOfFunctionPlot[581].Y = 0.68

	pointsOfFunctionPlot[582].X = 0.82
	pointsOfFunctionPlot[582].Y = 0.686

	pointsOfFunctionPlot[583].X = 0.83
	pointsOfFunctionPlot[583].Y = 0.692

	pointsOfFunctionPlot[584].X = 0.84
	pointsOfFunctionPlot[584].Y = 0.698

	pointsOfFunctionPlot[585].X = 0.85
	pointsOfFunctionPlot[585].Y = 0.704

	pointsOfFunctionPlot[586].X = 0.86
	pointsOfFunctionPlot[586].Y = 0.71

	pointsOfFunctionPlot[587].X = 0.87
	pointsOfFunctionPlot[587].Y = 0.715

	pointsOfFunctionPlot[588].X = 0.88
	pointsOfFunctionPlot[588].Y = 0.721

	pointsOfFunctionPlot[589].X = 0.89
	pointsOfFunctionPlot[589].Y = 0.727

	pointsOfFunctionPlot[590].X = 0.90
	pointsOfFunctionPlot[590].Y = 0.732

	pointsOfFunctionPlot[591].X = 0.91
	pointsOfFunctionPlot[591].Y = 0.738

	pointsOfFunctionPlot[592].X = 0.92
	pointsOfFunctionPlot[592].Y = 0.743

	pointsOfFunctionPlot[593].X = 0.93
	pointsOfFunctionPlot[593].Y = 0.749

	pointsOfFunctionPlot[594].X = 0.94
	pointsOfFunctionPlot[594].Y = 0.754

	pointsOfFunctionPlot[595].X = 0.95
	pointsOfFunctionPlot[595].Y = 0.759

	pointsOfFunctionPlot[596].X = 0.96
	pointsOfFunctionPlot[596].Y = 0.764

	pointsOfFunctionPlot[597].X = 0.97
	pointsOfFunctionPlot[597].Y = 0.77

	pointsOfFunctionPlot[598].X = 0.98
	pointsOfFunctionPlot[598].Y = 0.775

	pointsOfFunctionPlot[599].X = 0.99
	pointsOfFunctionPlot[599].Y = 0.78

	pointsOfFunctionPlot[600].X = 1.0
	pointsOfFunctionPlot[600].Y = 0.785

	pointsOfFunctionPlot[601].X = 1.01
	pointsOfFunctionPlot[601].Y = 0.79

	pointsOfFunctionPlot[602].X = 1.02
	pointsOfFunctionPlot[602].Y = 0.795

	pointsOfFunctionPlot[603].X = 1.03
	pointsOfFunctionPlot[603].Y = 0.8

	pointsOfFunctionPlot[604].X = 1.04
	pointsOfFunctionPlot[604].Y = 0.805

	pointsOfFunctionPlot[605].X = 1.05
	pointsOfFunctionPlot[605].Y = 0.809

	pointsOfFunctionPlot[606].X = 1.06
	pointsOfFunctionPlot[606].Y = 0.814

	pointsOfFunctionPlot[607].X = 1.07
	pointsOfFunctionPlot[607].Y = 0.819

	pointsOfFunctionPlot[608].X = 1.08
	pointsOfFunctionPlot[608].Y = 0.823

	pointsOfFunctionPlot[609].X = 1.09
	pointsOfFunctionPlot[609].Y = 0.828

	pointsOfFunctionPlot[610].X = 1.10
	pointsOfFunctionPlot[610].Y = 0.832

	pointsOfFunctionPlot[611].X = 1.11
	pointsOfFunctionPlot[611].Y = 0.837

	pointsOfFunctionPlot[612].X = 1.12
	pointsOfFunctionPlot[612].Y = 0.841

	pointsOfFunctionPlot[613].X = 1.13
	pointsOfFunctionPlot[613].Y = 0.846

	pointsOfFunctionPlot[614].X = 1.14
	pointsOfFunctionPlot[614].Y = 0.85

	pointsOfFunctionPlot[615].X = 1.15
	pointsOfFunctionPlot[615].Y = 0.855

	pointsOfFunctionPlot[616].X = 1.16
	pointsOfFunctionPlot[616].Y = 0.859

	pointsOfFunctionPlot[617].X = 1.17
	pointsOfFunctionPlot[617].Y = 0.863

	pointsOfFunctionPlot[618].X = 1.18
	pointsOfFunctionPlot[618].Y = 0.867

	pointsOfFunctionPlot[619].X = 1.19
	pointsOfFunctionPlot[619].Y = 0.871

	pointsOfFunctionPlot[620].X = 1.20
	pointsOfFunctionPlot[620].Y = 0.876

	pointsOfFunctionPlot[621].X = 1.21
	pointsOfFunctionPlot[621].Y = 0.88

	pointsOfFunctionPlot[622].X = 1.22
	pointsOfFunctionPlot[622].Y = 0.884

	pointsOfFunctionPlot[623].X = 1.23
	pointsOfFunctionPlot[623].Y = 0.888

	pointsOfFunctionPlot[624].X = 1.24
	pointsOfFunctionPlot[624].Y = 0.892

	pointsOfFunctionPlot[625].X = 1.25
	pointsOfFunctionPlot[625].Y = 0.896

	pointsOfFunctionPlot[626].X = 1.26
	pointsOfFunctionPlot[626].Y = 0.899

	pointsOfFunctionPlot[627].X = 1.27
	pointsOfFunctionPlot[627].Y = 0.903

	pointsOfFunctionPlot[628].X = 1.28
	pointsOfFunctionPlot[628].Y = 0.907

	pointsOfFunctionPlot[629].X = 1.29
	pointsOfFunctionPlot[629].Y = 0.911

	pointsOfFunctionPlot[630].X = 1.30
	pointsOfFunctionPlot[630].Y = 0.915

	pointsOfFunctionPlot[631].X = 1.31
	pointsOfFunctionPlot[631].Y = 0.918

	pointsOfFunctionPlot[632].X = 1.32
	pointsOfFunctionPlot[632].Y = 0.922

	pointsOfFunctionPlot[633].X = 1.33
	pointsOfFunctionPlot[633].Y = 0.926

	pointsOfFunctionPlot[634].X = 1.34
	pointsOfFunctionPlot[634].Y = 0.929

	pointsOfFunctionPlot[635].X = 1.35
	pointsOfFunctionPlot[635].Y = 0.933

	pointsOfFunctionPlot[636].X = 1.36
	pointsOfFunctionPlot[636].Y = 0.936

	pointsOfFunctionPlot[637].X = 1.37
	pointsOfFunctionPlot[637].Y = 0.94

	pointsOfFunctionPlot[638].X = 1.38
	pointsOfFunctionPlot[638].Y = 0.943

	pointsOfFunctionPlot[639].X = 1.39
	pointsOfFunctionPlot[639].Y = 0.947

	pointsOfFunctionPlot[640].X = 1.40
	pointsOfFunctionPlot[640].Y = 0.95

	pointsOfFunctionPlot[641].X = 1.41
	pointsOfFunctionPlot[641].Y = 0.953

	pointsOfFunctionPlot[642].X = 1.42
	pointsOfFunctionPlot[642].Y = 0.957

	pointsOfFunctionPlot[643].X = 1.43
	pointsOfFunctionPlot[643].Y = 0.96

	pointsOfFunctionPlot[644].X = 1.44
	pointsOfFunctionPlot[644].Y = 0.963

	pointsOfFunctionPlot[645].X = 1.45
	pointsOfFunctionPlot[645].Y = 0.967

	pointsOfFunctionPlot[646].X = 1.46
	pointsOfFunctionPlot[646].Y = 0.97

	pointsOfFunctionPlot[647].X = 1.47
	pointsOfFunctionPlot[647].Y = 0.973

	pointsOfFunctionPlot[648].X = 1.48
	pointsOfFunctionPlot[648].Y = 0.976

	pointsOfFunctionPlot[649].X = 1.49
	pointsOfFunctionPlot[649].Y = 0.979

	pointsOfFunctionPlot[650].X = 1.50
	pointsOfFunctionPlot[650].Y = 0.982

	pointsOfFunctionPlot[651].X = 1.51
	pointsOfFunctionPlot[651].Y = 0.985

	pointsOfFunctionPlot[652].X = 1.52
	pointsOfFunctionPlot[652].Y = 0.988

	pointsOfFunctionPlot[653].X = 1.53
	pointsOfFunctionPlot[653].Y = 0.991

	pointsOfFunctionPlot[654].X = 1.54
	pointsOfFunctionPlot[654].Y = 0.994

	pointsOfFunctionPlot[655].X = 1.55
	pointsOfFunctionPlot[655].Y = 0.997

	pointsOfFunctionPlot[656].X = 1.56
	pointsOfFunctionPlot[656].Y = 1.0

	pointsOfFunctionPlot[657].X = 1.57
	pointsOfFunctionPlot[657].Y = 1.003

	pointsOfFunctionPlot[658].X = 1.58
	pointsOfFunctionPlot[658].Y = 1.006

	pointsOfFunctionPlot[659].X = 1.59
	pointsOfFunctionPlot[659].Y = 1.009

	pointsOfFunctionPlot[660].X = 1.60
	pointsOfFunctionPlot[660].Y = 1.012

	pointsOfFunctionPlot[661].X = 1.61
	pointsOfFunctionPlot[661].Y = 1.014

	pointsOfFunctionPlot[662].X = 1.62
	pointsOfFunctionPlot[662].Y = 1.017

	pointsOfFunctionPlot[663].X = 1.63
	pointsOfFunctionPlot[663].Y = 1.02

	pointsOfFunctionPlot[664].X = 1.64
	pointsOfFunctionPlot[664].Y = 1.023

	pointsOfFunctionPlot[665].X = 1.65
	pointsOfFunctionPlot[665].Y = 1.025

	pointsOfFunctionPlot[666].X = 1.66
	pointsOfFunctionPlot[666].Y = 1.028

	pointsOfFunctionPlot[667].X = 1.67
	pointsOfFunctionPlot[667].Y = 1.031

	pointsOfFunctionPlot[668].X = 1.68
	pointsOfFunctionPlot[668].Y = 1.033

	pointsOfFunctionPlot[669].X = 1.69
	pointsOfFunctionPlot[669].Y = 1.036

	pointsOfFunctionPlot[670].X = 1.70
	pointsOfFunctionPlot[670].Y = 1.039

	pointsOfFunctionPlot[671].X = 1.71
	pointsOfFunctionPlot[671].Y = 1.041

	pointsOfFunctionPlot[672].X = 1.72
	pointsOfFunctionPlot[672].Y = 1.044

	pointsOfFunctionPlot[673].X = 1.73
	pointsOfFunctionPlot[673].Y = 1.046

	pointsOfFunctionPlot[674].X = 1.74
	pointsOfFunctionPlot[674].Y = 1.049

	pointsOfFunctionPlot[675].X = 1.75
	pointsOfFunctionPlot[675].Y = 1.051

	pointsOfFunctionPlot[676].X = 1.76
	pointsOfFunctionPlot[676].Y = 1.054

	pointsOfFunctionPlot[677].X = 1.77
	pointsOfFunctionPlot[677].Y = 1.056

	pointsOfFunctionPlot[678].X = 1.78
	pointsOfFunctionPlot[678].Y = 1.058

	pointsOfFunctionPlot[679].X = 1.79
	pointsOfFunctionPlot[679].Y = 1.061

	pointsOfFunctionPlot[680].X = 1.80
	pointsOfFunctionPlot[680].Y = 1.063

	pointsOfFunctionPlot[681].X = 1.81
	pointsOfFunctionPlot[681].Y = 1.066

	pointsOfFunctionPlot[682].X = 1.82
	pointsOfFunctionPlot[682].Y = 1.068

	pointsOfFunctionPlot[683].X = 1.83
	pointsOfFunctionPlot[683].Y = 1.07

	pointsOfFunctionPlot[684].X = 1.84
	pointsOfFunctionPlot[684].Y = 1.072

	pointsOfFunctionPlot[685].X = 1.85
	pointsOfFunctionPlot[685].Y = 1.075

	pointsOfFunctionPlot[686].X = 1.86
	pointsOfFunctionPlot[686].Y = 1.077

	pointsOfFunctionPlot[687].X = 1.87
	pointsOfFunctionPlot[687].Y = 1.079

	pointsOfFunctionPlot[688].X = 1.88
	pointsOfFunctionPlot[688].Y = 1.081

	pointsOfFunctionPlot[689].X = 1.89
	pointsOfFunctionPlot[689].Y = 1.084

	pointsOfFunctionPlot[690].X = 1.90
	pointsOfFunctionPlot[690].Y = 1.086

	pointsOfFunctionPlot[691].X = 1.91
	pointsOfFunctionPlot[691].Y = 1.088

	pointsOfFunctionPlot[692].X = 1.92
	pointsOfFunctionPlot[692].Y = 1.09

	pointsOfFunctionPlot[693].X = 1.93
	pointsOfFunctionPlot[693].Y = 1.092

	pointsOfFunctionPlot[694].X = 1.94
	pointsOfFunctionPlot[694].Y = 1.094

	pointsOfFunctionPlot[695].X = 1.95
	pointsOfFunctionPlot[695].Y = 1.096

	pointsOfFunctionPlot[696].X = 1.96
	pointsOfFunctionPlot[696].Y = 1.099

	pointsOfFunctionPlot[697].X = 1.97
	pointsOfFunctionPlot[697].Y = 1.101

	pointsOfFunctionPlot[698].X = 1.98
	pointsOfFunctionPlot[698].Y = 1.103

	pointsOfFunctionPlot[699].X = 1.99
	pointsOfFunctionPlot[699].Y = 1.105

	pointsOfFunctionPlot[700].X = 2.0
	pointsOfFunctionPlot[700].Y = 1.107

	pointsOfFunctionPlot[701].X = 2.01
	pointsOfFunctionPlot[701].Y = 1.109

	pointsOfFunctionPlot[702].X = 2.02
	pointsOfFunctionPlot[702].Y = 1.111

	pointsOfFunctionPlot[703].X = 2.03
	pointsOfFunctionPlot[703].Y = 1.113

	pointsOfFunctionPlot[704].X = 2.04
	pointsOfFunctionPlot[704].Y = 1.115

	pointsOfFunctionPlot[705].X = 2.05
	pointsOfFunctionPlot[705].Y = 1.116

	pointsOfFunctionPlot[706].X = 2.06
	pointsOfFunctionPlot[706].Y = 1.118

	pointsOfFunctionPlot[707].X = 2.07
	pointsOfFunctionPlot[707].Y = 1.12

	pointsOfFunctionPlot[708].X = 2.08
	pointsOfFunctionPlot[708].Y = 1.122

	pointsOfFunctionPlot[709].X = 2.09
	pointsOfFunctionPlot[709].Y = 1.124

	pointsOfFunctionPlot[710].X = 2.10
	pointsOfFunctionPlot[710].Y = 1.126

	pointsOfFunctionPlot[711].X = 2.11
	pointsOfFunctionPlot[711].Y = 1.128

	pointsOfFunctionPlot[712].X = 2.12
	pointsOfFunctionPlot[712].Y = 1.13

	pointsOfFunctionPlot[713].X = 2.13
	pointsOfFunctionPlot[713].Y = 1.131

	pointsOfFunctionPlot[714].X = 2.14
	pointsOfFunctionPlot[714].Y = 1.133

	pointsOfFunctionPlot[715].X = 2.15
	pointsOfFunctionPlot[715].Y = 1.135

	pointsOfFunctionPlot[716].X = 2.16
	pointsOfFunctionPlot[716].Y = 1.137

	pointsOfFunctionPlot[717].X = 2.17
	pointsOfFunctionPlot[717].Y = 1.138

	pointsOfFunctionPlot[718].X = 2.18
	pointsOfFunctionPlot[718].Y = 1.14

	pointsOfFunctionPlot[719].X = 2.19
	pointsOfFunctionPlot[719].Y = 1.142

	pointsOfFunctionPlot[720].X = 2.20
	pointsOfFunctionPlot[720].Y = 1.144

	pointsOfFunctionPlot[721].X = 2.21
	pointsOfFunctionPlot[721].Y = 1.145

	pointsOfFunctionPlot[722].X = 2.22
	pointsOfFunctionPlot[722].Y = 1.147

	pointsOfFunctionPlot[723].X = 2.23
	pointsOfFunctionPlot[723].Y = 1.149

	pointsOfFunctionPlot[724].X = 2.24
	pointsOfFunctionPlot[724].Y = 1.15

	pointsOfFunctionPlot[725].X = 2.25
	pointsOfFunctionPlot[725].Y = 1.152

	pointsOfFunctionPlot[726].X = 2.26
	pointsOfFunctionPlot[726].Y = 1.154

	pointsOfFunctionPlot[727].X = 2.27
	pointsOfFunctionPlot[727].Y = 1.155

	pointsOfFunctionPlot[728].X = 2.28
	pointsOfFunctionPlot[728].Y = 1.157

	pointsOfFunctionPlot[729].X = 2.29
	pointsOfFunctionPlot[729].Y = 1.159

	pointsOfFunctionPlot[730].X = 2.30
	pointsOfFunctionPlot[730].Y = 1.16

	pointsOfFunctionPlot[731].X = 2.31
	pointsOfFunctionPlot[731].Y = 1.162

	pointsOfFunctionPlot[732].X = 2.32
	pointsOfFunctionPlot[732].Y = 1.163

	pointsOfFunctionPlot[733].X = 2.33
	pointsOfFunctionPlot[733].Y = 1.165

	pointsOfFunctionPlot[734].X = 2.34
	pointsOfFunctionPlot[734].Y = 1.166

	pointsOfFunctionPlot[735].X = 2.35
	pointsOfFunctionPlot[735].Y = 1.168

	pointsOfFunctionPlot[736].X = 2.36
	pointsOfFunctionPlot[736].Y = 1.17

	pointsOfFunctionPlot[737].X = 2.37
	pointsOfFunctionPlot[737].Y = 1.171

	pointsOfFunctionPlot[738].X = 2.38
	pointsOfFunctionPlot[738].Y = 1.173

	pointsOfFunctionPlot[739].X = 2.39
	pointsOfFunctionPlot[739].Y = 1.174

	pointsOfFunctionPlot[740].X = 2.40
	pointsOfFunctionPlot[740].Y = 1.176

	pointsOfFunctionPlot[741].X = 2.41
	pointsOfFunctionPlot[741].Y = 1.177

	pointsOfFunctionPlot[742].X = 2.42
	pointsOfFunctionPlot[742].Y = 1.178

	pointsOfFunctionPlot[743].X = 2.43
	pointsOfFunctionPlot[743].Y = 1.18

	pointsOfFunctionPlot[744].X = 2.44
	pointsOfFunctionPlot[744].Y = 1.181

	pointsOfFunctionPlot[745].X = 2.45
	pointsOfFunctionPlot[745].Y = 1.183

	pointsOfFunctionPlot[746].X = 2.46
	pointsOfFunctionPlot[746].Y = 1.184

	pointsOfFunctionPlot[747].X = 2.47
	pointsOfFunctionPlot[747].Y = 1.186

	pointsOfFunctionPlot[748].X = 2.48
	pointsOfFunctionPlot[748].Y = 1.187

	pointsOfFunctionPlot[749].X = 2.49
	pointsOfFunctionPlot[749].Y = 1.188

	pointsOfFunctionPlot[750].X = 2.50
	pointsOfFunctionPlot[750].Y = 1.19

	pointsOfFunctionPlot[751].X = 2.51
	pointsOfFunctionPlot[751].Y = 1.191

	pointsOfFunctionPlot[752].X = 2.52
	pointsOfFunctionPlot[752].Y = 1.193

	pointsOfFunctionPlot[753].X = 2.53
	pointsOfFunctionPlot[753].Y = 1.194

	pointsOfFunctionPlot[754].X = 2.54
	pointsOfFunctionPlot[754].Y = 1.195

	pointsOfFunctionPlot[755].X = 2.55
	pointsOfFunctionPlot[755].Y = 1.197

	pointsOfFunctionPlot[756].X = 2.56
	pointsOfFunctionPlot[756].Y = 1.198

	pointsOfFunctionPlot[757].X = 2.57
	pointsOfFunctionPlot[757].Y = 1.199

	pointsOfFunctionPlot[758].X = 2.58
	pointsOfFunctionPlot[758].Y = 1.201

	pointsOfFunctionPlot[759].X = 2.59
	pointsOfFunctionPlot[759].Y = 1.202

	pointsOfFunctionPlot[760].X = 2.60
	pointsOfFunctionPlot[760].Y = 1.203

	pointsOfFunctionPlot[761].X = 2.61
	pointsOfFunctionPlot[761].Y = 1.204

	pointsOfFunctionPlot[762].X = 2.62
	pointsOfFunctionPlot[762].Y = 1.206

	pointsOfFunctionPlot[763].X = 2.63
	pointsOfFunctionPlot[763].Y = 1.207

	pointsOfFunctionPlot[764].X = 2.64
	pointsOfFunctionPlot[764].Y = 1.208

	pointsOfFunctionPlot[765].X = 2.65
	pointsOfFunctionPlot[765].Y = 1.209

	pointsOfFunctionPlot[766].X = 2.66
	pointsOfFunctionPlot[766].Y = 1.211

	pointsOfFunctionPlot[767].X = 2.67
	pointsOfFunctionPlot[767].Y = 1.212

	pointsOfFunctionPlot[768].X = 2.68
	pointsOfFunctionPlot[768].Y = 1.213

	pointsOfFunctionPlot[769].X = 2.69
	pointsOfFunctionPlot[769].Y = 1.214

	pointsOfFunctionPlot[770].X = 2.70
	pointsOfFunctionPlot[770].Y = 1.216

	pointsOfFunctionPlot[771].X = 2.71
	pointsOfFunctionPlot[771].Y = 1.217

	pointsOfFunctionPlot[772].X = 2.72
	pointsOfFunctionPlot[772].Y = 1.218

	pointsOfFunctionPlot[773].X = 2.73
	pointsOfFunctionPlot[773].Y = 1.219

	pointsOfFunctionPlot[774].X = 2.74
	pointsOfFunctionPlot[774].Y = 1.22

	pointsOfFunctionPlot[775].X = 2.75
	pointsOfFunctionPlot[775].Y = 1.222

	pointsOfFunctionPlot[776].X = 2.76
	pointsOfFunctionPlot[776].Y = 1.223

	pointsOfFunctionPlot[777].X = 2.77
	pointsOfFunctionPlot[777].Y = 1.224

	pointsOfFunctionPlot[778].X = 2.78
	pointsOfFunctionPlot[778].Y = 1.225

	pointsOfFunctionPlot[779].X = 2.79
	pointsOfFunctionPlot[779].Y = 1.226

	pointsOfFunctionPlot[780].X = 2.80
	pointsOfFunctionPlot[780].Y = 1.227

	pointsOfFunctionPlot[781].X = 2.81
	pointsOfFunctionPlot[781].Y = 1.228

	pointsOfFunctionPlot[782].X = 2.82
	pointsOfFunctionPlot[782].Y = 1.23

	pointsOfFunctionPlot[783].X = 2.83
	pointsOfFunctionPlot[783].Y = 1.231

	pointsOfFunctionPlot[784].X = 2.84
	pointsOfFunctionPlot[784].Y = 1.232

	pointsOfFunctionPlot[785].X = 2.85
	pointsOfFunctionPlot[785].Y = 1.233

	pointsOfFunctionPlot[786].X = 2.86
	pointsOfFunctionPlot[786].Y = 1.234

	pointsOfFunctionPlot[787].X = 2.87
	pointsOfFunctionPlot[787].Y = 1.235

	pointsOfFunctionPlot[788].X = 2.88
	pointsOfFunctionPlot[788].Y = 1.236

	pointsOfFunctionPlot[789].X = 2.89
	pointsOfFunctionPlot[789].Y = 1.237

	pointsOfFunctionPlot[790].X = 2.90
	pointsOfFunctionPlot[790].Y = 1.238

	pointsOfFunctionPlot[791].X = 2.91
	pointsOfFunctionPlot[791].Y = 1.239

	pointsOfFunctionPlot[792].X = 2.92
	pointsOfFunctionPlot[792].Y = 1.24

	pointsOfFunctionPlot[793].X = 2.93
	pointsOfFunctionPlot[793].Y = 1.241

	pointsOfFunctionPlot[794].X = 2.94
	pointsOfFunctionPlot[794].Y = 1.242

	pointsOfFunctionPlot[795].X = 2.95
	pointsOfFunctionPlot[795].Y = 1.243

	pointsOfFunctionPlot[796].X = 2.96
	pointsOfFunctionPlot[796].Y = 1.244

	pointsOfFunctionPlot[797].X = 2.97
	pointsOfFunctionPlot[797].Y = 1.246

	pointsOfFunctionPlot[798].X = 2.98
	pointsOfFunctionPlot[798].Y = 1.247

	pointsOfFunctionPlot[799].X = 2.99
	pointsOfFunctionPlot[799].Y = 1.248

	pointsOfFunctionPlot[800].X = 3.0
	pointsOfFunctionPlot[800].Y = 1.249

	pointsOfFunctionPlot[801].X = 3.01
	pointsOfFunctionPlot[801].Y = 1.25

	pointsOfFunctionPlot[802].X = 3.02
	pointsOfFunctionPlot[802].Y = 1.251

	pointsOfFunctionPlot[803].X = 3.03
	pointsOfFunctionPlot[803].Y = 1.252

	pointsOfFunctionPlot[804].X = 3.04
	pointsOfFunctionPlot[804].Y = 1.252

	pointsOfFunctionPlot[805].X = 3.05
	pointsOfFunctionPlot[805].Y = 1.253

	pointsOfFunctionPlot[806].X = 3.06
	pointsOfFunctionPlot[806].Y = 1.254

	pointsOfFunctionPlot[807].X = 3.07
	pointsOfFunctionPlot[807].Y = 1.255

	pointsOfFunctionPlot[808].X = 3.08
	pointsOfFunctionPlot[808].Y = 1.256

	pointsOfFunctionPlot[809].X = 3.09
	pointsOfFunctionPlot[809].Y = 1.257

	pointsOfFunctionPlot[810].X = 3.10
	pointsOfFunctionPlot[810].Y = 1.258

	pointsOfFunctionPlot[811].X = 3.11
	pointsOfFunctionPlot[811].Y = 1.259

	pointsOfFunctionPlot[812].X = 3.12
	pointsOfFunctionPlot[812].Y = 1.26

	pointsOfFunctionPlot[813].X = 3.13
	pointsOfFunctionPlot[813].Y = 1.261

	pointsOfFunctionPlot[814].X = 3.14
	pointsOfFunctionPlot[814].Y = 1.262

	pointsOfFunctionPlot[815].X = 3.15
	pointsOfFunctionPlot[815].Y = 1.263

	pointsOfFunctionPlot[816].X = 3.16
	pointsOfFunctionPlot[816].Y = 1.264

	pointsOfFunctionPlot[817].X = 3.17
	pointsOfFunctionPlot[817].Y = 1.265

	pointsOfFunctionPlot[818].X = 3.18
	pointsOfFunctionPlot[818].Y = 1.266

	pointsOfFunctionPlot[819].X = 3.19
	pointsOfFunctionPlot[819].Y = 1.267

	pointsOfFunctionPlot[820].X = 3.20
	pointsOfFunctionPlot[820].Y = 1.267

	pointsOfFunctionPlot[821].X = 3.21
	pointsOfFunctionPlot[821].Y = 1.268

	pointsOfFunctionPlot[822].X = 3.22
	pointsOfFunctionPlot[822].Y = 1.269

	pointsOfFunctionPlot[823].X = 3.23
	pointsOfFunctionPlot[823].Y = 1.27

	pointsOfFunctionPlot[824].X = 3.24
	pointsOfFunctionPlot[824].Y = 1.271

	pointsOfFunctionPlot[825].X = 3.25
	pointsOfFunctionPlot[825].Y = 1.272

	pointsOfFunctionPlot[826].X = 3.26
	pointsOfFunctionPlot[826].Y = 1.273

	pointsOfFunctionPlot[827].X = 3.27
	pointsOfFunctionPlot[827].Y = 1.274

	pointsOfFunctionPlot[828].X = 3.28
	pointsOfFunctionPlot[828].Y = 1.274

	pointsOfFunctionPlot[829].X = 3.29
	pointsOfFunctionPlot[829].Y = 1.275

	pointsOfFunctionPlot[830].X = 3.30
	pointsOfFunctionPlot[830].Y = 1.276

	pointsOfFunctionPlot[831].X = 3.31
	pointsOfFunctionPlot[831].Y = 1.277

	pointsOfFunctionPlot[832].X = 3.32
	pointsOfFunctionPlot[832].Y = 1.278

	pointsOfFunctionPlot[833].X = 3.33
	pointsOfFunctionPlot[833].Y = 1.279

	pointsOfFunctionPlot[834].X = 3.34
	pointsOfFunctionPlot[834].Y = 1.279

	pointsOfFunctionPlot[835].X = 3.35
	pointsOfFunctionPlot[835].Y = 1.28

	pointsOfFunctionPlot[836].X = 3.36
	pointsOfFunctionPlot[836].Y = 1.281

	pointsOfFunctionPlot[837].X = 3.37
	pointsOfFunctionPlot[837].Y = 1.282

	pointsOfFunctionPlot[838].X = 3.38
	pointsOfFunctionPlot[838].Y = 1.283

	pointsOfFunctionPlot[839].X = 3.39
	pointsOfFunctionPlot[839].Y = 1.283

	pointsOfFunctionPlot[840].X = 3.40
	pointsOfFunctionPlot[840].Y = 1.284

	pointsOfFunctionPlot[841].X = 3.41
	pointsOfFunctionPlot[841].Y = 1.285

	pointsOfFunctionPlot[842].X = 3.42
	pointsOfFunctionPlot[842].Y = 1.286

	pointsOfFunctionPlot[843].X = 3.43
	pointsOfFunctionPlot[843].Y = 1.287

	pointsOfFunctionPlot[844].X = 3.44
	pointsOfFunctionPlot[844].Y = 1.287

	pointsOfFunctionPlot[845].X = 3.45
	pointsOfFunctionPlot[845].Y = 1.288

	pointsOfFunctionPlot[846].X = 3.46
	pointsOfFunctionPlot[846].Y = 1.289

	pointsOfFunctionPlot[847].X = 3.47
	pointsOfFunctionPlot[847].Y = 1.29

	pointsOfFunctionPlot[848].X = 3.48
	pointsOfFunctionPlot[848].Y = 1.29

	pointsOfFunctionPlot[849].X = 3.49
	pointsOfFunctionPlot[849].Y = 1.291

	pointsOfFunctionPlot[850].X = 3.50
	pointsOfFunctionPlot[850].Y = 1.292

	pointsOfFunctionPlot[851].X = 3.51
	pointsOfFunctionPlot[851].Y = 1.293

	pointsOfFunctionPlot[852].X = 3.52
	pointsOfFunctionPlot[852].Y = 1.293

	pointsOfFunctionPlot[853].X = 3.53
	pointsOfFunctionPlot[853].Y = 1.294

	pointsOfFunctionPlot[854].X = 3.54
	pointsOfFunctionPlot[854].Y = 1.295

	pointsOfFunctionPlot[855].X = 3.55
	pointsOfFunctionPlot[855].Y = 1.296

	pointsOfFunctionPlot[856].X = 3.56
	pointsOfFunctionPlot[856].Y = 1.296

	pointsOfFunctionPlot[857].X = 3.57
	pointsOfFunctionPlot[857].Y = 1.297

	pointsOfFunctionPlot[858].X = 3.58
	pointsOfFunctionPlot[858].Y = 1.298

	pointsOfFunctionPlot[859].X = 3.59
	pointsOfFunctionPlot[859].Y = 1.299

	pointsOfFunctionPlot[860].X = 3.60
	pointsOfFunctionPlot[860].Y = 1.299

	pointsOfFunctionPlot[861].X = 3.61
	pointsOfFunctionPlot[861].Y = 1.3

	pointsOfFunctionPlot[862].X = 3.62
	pointsOfFunctionPlot[862].Y = 1.301

	pointsOfFunctionPlot[863].X = 3.63
	pointsOfFunctionPlot[863].Y = 1.301

	pointsOfFunctionPlot[864].X = 3.64
	pointsOfFunctionPlot[864].Y = 1.302

	pointsOfFunctionPlot[865].X = 3.65
	pointsOfFunctionPlot[865].Y = 1.303

	pointsOfFunctionPlot[866].X = 3.66
	pointsOfFunctionPlot[866].Y = 1.304

	pointsOfFunctionPlot[867].X = 3.67
	pointsOfFunctionPlot[867].Y = 1.304

	pointsOfFunctionPlot[868].X = 3.68
	pointsOfFunctionPlot[868].Y = 1.305

	pointsOfFunctionPlot[869].X = 3.69
	pointsOfFunctionPlot[869].Y = 1.306

	pointsOfFunctionPlot[870].X = 3.70
	pointsOfFunctionPlot[870].Y = 1.306

	pointsOfFunctionPlot[871].X = 3.71
	pointsOfFunctionPlot[871].Y = 1.307

	pointsOfFunctionPlot[872].X = 3.72
	pointsOfFunctionPlot[872].Y = 1.308

	pointsOfFunctionPlot[873].X = 3.73
	pointsOfFunctionPlot[873].Y = 1.308

	pointsOfFunctionPlot[874].X = 3.74
	pointsOfFunctionPlot[874].Y = 1.309

	pointsOfFunctionPlot[875].X = 3.75
	pointsOfFunctionPlot[875].Y = 1.31

	pointsOfFunctionPlot[876].X = 3.76
	pointsOfFunctionPlot[876].Y = 1.31

	pointsOfFunctionPlot[877].X = 3.77
	pointsOfFunctionPlot[877].Y = 1.311

	pointsOfFunctionPlot[878].X = 3.78
	pointsOfFunctionPlot[878].Y = 1.312

	pointsOfFunctionPlot[879].X = 3.79
	pointsOfFunctionPlot[879].Y = 1.312

	pointsOfFunctionPlot[880].X = 3.80
	pointsOfFunctionPlot[880].Y = 1.313

	pointsOfFunctionPlot[881].X = 3.81
	pointsOfFunctionPlot[881].Y = 1.314

	pointsOfFunctionPlot[882].X = 3.82
	pointsOfFunctionPlot[882].Y = 1.314

	pointsOfFunctionPlot[883].X = 3.83
	pointsOfFunctionPlot[883].Y = 1.315

	pointsOfFunctionPlot[884].X = 3.84
	pointsOfFunctionPlot[884].Y = 1.316

	pointsOfFunctionPlot[885].X = 3.85
	pointsOfFunctionPlot[885].Y = 1.316

	pointsOfFunctionPlot[886].X = 3.86
	pointsOfFunctionPlot[886].Y = 1.317

	pointsOfFunctionPlot[887].X = 3.87
	pointsOfFunctionPlot[887].Y = 1.317

	pointsOfFunctionPlot[888].X = 3.88
	pointsOfFunctionPlot[888].Y = 1.318

	pointsOfFunctionPlot[889].X = 3.89
	pointsOfFunctionPlot[889].Y = 1.319

	pointsOfFunctionPlot[890].X = 3.90
	pointsOfFunctionPlot[890].Y = 1.319

	pointsOfFunctionPlot[891].X = 3.91
	pointsOfFunctionPlot[891].Y = 1.32

	pointsOfFunctionPlot[892].X = 3.92
	pointsOfFunctionPlot[892].Y = 1.321

	pointsOfFunctionPlot[893].X = 3.93
	pointsOfFunctionPlot[893].Y = 1.321

	pointsOfFunctionPlot[894].X = 3.94
	pointsOfFunctionPlot[894].Y = 1.322

	pointsOfFunctionPlot[895].X = 3.95
	pointsOfFunctionPlot[895].Y = 1.322

	pointsOfFunctionPlot[896].X = 3.96
	pointsOfFunctionPlot[896].Y = 1.323

	pointsOfFunctionPlot[897].X = 3.97
	pointsOfFunctionPlot[897].Y = 1.324

	pointsOfFunctionPlot[898].X = 3.98
	pointsOfFunctionPlot[898].Y = 1.324

	pointsOfFunctionPlot[899].X = 3.99
	pointsOfFunctionPlot[899].Y = 1.325

	pointsOfFunctionPlot[900].X = 4.0
	pointsOfFunctionPlot[900].Y = 1.325

	pointsOfFunctionPlot[901].X = 4.01
	pointsOfFunctionPlot[901].Y = 1.326

	pointsOfFunctionPlot[902].X = 4.02
	pointsOfFunctionPlot[902].Y = 1.326

	pointsOfFunctionPlot[903].X = 4.03
	pointsOfFunctionPlot[903].Y = 1.327

	pointsOfFunctionPlot[904].X = 4.04
	pointsOfFunctionPlot[904].Y = 1.328

	pointsOfFunctionPlot[905].X = 4.05
	pointsOfFunctionPlot[905].Y = 1.328

	pointsOfFunctionPlot[906].X = 4.06
	pointsOfFunctionPlot[906].Y = 1.329

	pointsOfFunctionPlot[907].X = 4.07
	pointsOfFunctionPlot[907].Y = 1.329

	pointsOfFunctionPlot[908].X = 4.08
	pointsOfFunctionPlot[908].Y = 1.33

	pointsOfFunctionPlot[909].X = 4.09
	pointsOfFunctionPlot[909].Y = 1.331

	pointsOfFunctionPlot[910].X = 4.10
	pointsOfFunctionPlot[910].Y = 1.331

	pointsOfFunctionPlot[911].X = 4.11
	pointsOfFunctionPlot[911].Y = 1.332

	pointsOfFunctionPlot[912].X = 4.12
	pointsOfFunctionPlot[912].Y = 1.332

	pointsOfFunctionPlot[913].X = 4.13
	pointsOfFunctionPlot[913].Y = 1.333

	pointsOfFunctionPlot[914].X = 4.14
	pointsOfFunctionPlot[914].Y = 1.333

	pointsOfFunctionPlot[915].X = 4.15
	pointsOfFunctionPlot[915].Y = 1.334

	pointsOfFunctionPlot[916].X = 4.16
	pointsOfFunctionPlot[916].Y = 1.334

	pointsOfFunctionPlot[917].X = 4.17
	pointsOfFunctionPlot[917].Y = 1.335

	pointsOfFunctionPlot[918].X = 4.18
	pointsOfFunctionPlot[918].Y = 1.335

	pointsOfFunctionPlot[919].X = 4.19
	pointsOfFunctionPlot[919].Y = 1.336

	pointsOfFunctionPlot[920].X = 4.20
	pointsOfFunctionPlot[920].Y = 1.337

	pointsOfFunctionPlot[921].X = 4.21
	pointsOfFunctionPlot[921].Y = 1.337

	pointsOfFunctionPlot[922].X = 4.22
	pointsOfFunctionPlot[922].Y = 1.338

	pointsOfFunctionPlot[923].X = 4.23
	pointsOfFunctionPlot[923].Y = 1.338

	pointsOfFunctionPlot[924].X = 4.24
	pointsOfFunctionPlot[924].Y = 1.339

	pointsOfFunctionPlot[925].X = 4.25
	pointsOfFunctionPlot[925].Y = 1.339

	pointsOfFunctionPlot[926].X = 4.26
	pointsOfFunctionPlot[926].Y = 1.34

	pointsOfFunctionPlot[927].X = 4.27
	pointsOfFunctionPlot[927].Y = 1.34

	pointsOfFunctionPlot[928].X = 4.28
	pointsOfFunctionPlot[928].Y = 1.341

	pointsOfFunctionPlot[929].X = 4.29
	pointsOfFunctionPlot[929].Y = 1.341

	pointsOfFunctionPlot[930].X = 4.30
	pointsOfFunctionPlot[930].Y = 1.342

	pointsOfFunctionPlot[931].X = 4.31
	pointsOfFunctionPlot[931].Y = 1.342

	pointsOfFunctionPlot[932].X = 4.32
	pointsOfFunctionPlot[932].Y = 1.343

	pointsOfFunctionPlot[933].X = 4.33
	pointsOfFunctionPlot[933].Y = 1.343

	pointsOfFunctionPlot[934].X = 4.34
	pointsOfFunctionPlot[934].Y = 1.344

	pointsOfFunctionPlot[935].X = 4.35
	pointsOfFunctionPlot[935].Y = 1.344

	pointsOfFunctionPlot[936].X = 4.36
	pointsOfFunctionPlot[936].Y = 1.345

	pointsOfFunctionPlot[937].X = 4.37
	pointsOfFunctionPlot[937].Y = 1.345

	pointsOfFunctionPlot[938].X = 4.38
	pointsOfFunctionPlot[938].Y = 1.346

	pointsOfFunctionPlot[939].X = 4.39
	pointsOfFunctionPlot[939].Y = 1.346

	pointsOfFunctionPlot[940].X = 4.40
	pointsOfFunctionPlot[940].Y = 1.347

	pointsOfFunctionPlot[941].X = 4.41
	pointsOfFunctionPlot[941].Y = 1.347

	pointsOfFunctionPlot[942].X = 4.42
	pointsOfFunctionPlot[942].Y = 1.348

	pointsOfFunctionPlot[943].X = 4.43
	pointsOfFunctionPlot[943].Y = 1.348

	pointsOfFunctionPlot[944].X = 4.44
	pointsOfFunctionPlot[944].Y = 1.349

	pointsOfFunctionPlot[945].X = 4.45
	pointsOfFunctionPlot[945].Y = 1.349

	pointsOfFunctionPlot[946].X = 4.46
	pointsOfFunctionPlot[946].Y = 1.35

	pointsOfFunctionPlot[947].X = 4.47
	pointsOfFunctionPlot[947].Y = 1.35

	pointsOfFunctionPlot[948].X = 4.48
	pointsOfFunctionPlot[948].Y = 1.351

	pointsOfFunctionPlot[949].X = 4.49
	pointsOfFunctionPlot[949].Y = 1.351

	pointsOfFunctionPlot[950].X = 4.50
	pointsOfFunctionPlot[950].Y = 1.352

	pointsOfFunctionPlot[951].X = 4.51
	pointsOfFunctionPlot[951].Y = 1.352

	pointsOfFunctionPlot[952].X = 4.52
	pointsOfFunctionPlot[952].Y = 1.353

	pointsOfFunctionPlot[953].X = 4.53
	pointsOfFunctionPlot[953].Y = 1.353

	pointsOfFunctionPlot[954].X = 4.54
	pointsOfFunctionPlot[954].Y = 1.353

	pointsOfFunctionPlot[955].X = 4.55
	pointsOfFunctionPlot[955].Y = 1.354

	pointsOfFunctionPlot[956].X = 4.56
	pointsOfFunctionPlot[956].Y = 1.354

	pointsOfFunctionPlot[957].X = 4.57
	pointsOfFunctionPlot[957].Y = 1.355

	pointsOfFunctionPlot[958].X = 4.58
	pointsOfFunctionPlot[958].Y = 1.355

	pointsOfFunctionPlot[959].X = 4.59
	pointsOfFunctionPlot[959].Y = 1.356

	pointsOfFunctionPlot[960].X = 4.60
	pointsOfFunctionPlot[960].Y = 1.356

	pointsOfFunctionPlot[961].X = 4.61
	pointsOfFunctionPlot[961].Y = 1.357

	pointsOfFunctionPlot[962].X = 4.62
	pointsOfFunctionPlot[962].Y = 1.357

	pointsOfFunctionPlot[963].X = 4.63
	pointsOfFunctionPlot[963].Y = 1.358

	pointsOfFunctionPlot[964].X = 4.64
	pointsOfFunctionPlot[964].Y = 1.358

	pointsOfFunctionPlot[965].X = 4.65
	pointsOfFunctionPlot[965].Y = 1.358

	pointsOfFunctionPlot[966].X = 4.66
	pointsOfFunctionPlot[966].Y = 1.359

	pointsOfFunctionPlot[967].X = 4.67
	pointsOfFunctionPlot[967].Y = 1.359

	pointsOfFunctionPlot[968].X = 4.68
	pointsOfFunctionPlot[968].Y = 1.36

	pointsOfFunctionPlot[969].X = 4.69
	pointsOfFunctionPlot[969].Y = 1.36

	pointsOfFunctionPlot[970].X = 4.70
	pointsOfFunctionPlot[970].Y = 1.361

	pointsOfFunctionPlot[971].X = 4.71
	pointsOfFunctionPlot[971].Y = 1.361

	pointsOfFunctionPlot[972].X = 4.72
	pointsOfFunctionPlot[972].Y = 1.362

	pointsOfFunctionPlot[973].X = 4.73
	pointsOfFunctionPlot[973].Y = 1.362

	pointsOfFunctionPlot[974].X = 4.74
	pointsOfFunctionPlot[974].Y = 1.362

	pointsOfFunctionPlot[975].X = 4.75
	pointsOfFunctionPlot[975].Y = 1.363

	pointsOfFunctionPlot[976].X = 4.76
	pointsOfFunctionPlot[976].Y = 1.363

	pointsOfFunctionPlot[977].X = 4.77
	pointsOfFunctionPlot[977].Y = 1.364

	pointsOfFunctionPlot[978].X = 4.78
	pointsOfFunctionPlot[978].Y = 1.364

	pointsOfFunctionPlot[979].X = 4.79
	pointsOfFunctionPlot[979].Y = 1.364

	pointsOfFunctionPlot[980].X = 4.80
	pointsOfFunctionPlot[980].Y = 1.365

	pointsOfFunctionPlot[981].X = 4.81
	pointsOfFunctionPlot[981].Y = 1.365

	pointsOfFunctionPlot[982].X = 4.82
	pointsOfFunctionPlot[982].Y = 1.366

	pointsOfFunctionPlot[983].X = 4.83
	pointsOfFunctionPlot[983].Y = 1.366

	pointsOfFunctionPlot[984].X = 4.84
	pointsOfFunctionPlot[984].Y = 1.367

	pointsOfFunctionPlot[985].X = 4.85
	pointsOfFunctionPlot[985].Y = 1.367

	pointsOfFunctionPlot[986].X = 4.86
	pointsOfFunctionPlot[986].Y = 1.367

	pointsOfFunctionPlot[987].X = 4.87
	pointsOfFunctionPlot[987].Y = 1.368

	pointsOfFunctionPlot[988].X = 4.88
	pointsOfFunctionPlot[988].Y = 1.368

	pointsOfFunctionPlot[989].X = 4.89
	pointsOfFunctionPlot[989].Y = 1.369

	pointsOfFunctionPlot[990].X = 4.90
	pointsOfFunctionPlot[990].Y = 1.369

	pointsOfFunctionPlot[991].X = 4.91
	pointsOfFunctionPlot[991].Y = 1.369

	pointsOfFunctionPlot[992].X = 4.92
	pointsOfFunctionPlot[992].Y = 1.37

	pointsOfFunctionPlot[993].X = 4.93
	pointsOfFunctionPlot[993].Y = 1.37

	pointsOfFunctionPlot[994].X = 4.94
	pointsOfFunctionPlot[994].Y = 1.371

	pointsOfFunctionPlot[995].X = 4.95
	pointsOfFunctionPlot[995].Y = 1.371

	pointsOfFunctionPlot[996].X = 4.96
	pointsOfFunctionPlot[996].Y = 1.371

	pointsOfFunctionPlot[997].X = 4.97
	pointsOfFunctionPlot[997].Y = 1.372

	pointsOfFunctionPlot[998].X = 4.98
	pointsOfFunctionPlot[998].Y = 1.372

	pointsOfFunctionPlot[999].X = 4.99
	pointsOfFunctionPlot[999].Y = 1.373

	pointsOfFunctionPlot[1_000].X = 5.0
	pointsOfFunctionPlot[1_000].Y = 1.373










	plotOfFunction := plot.New()

	plotOfFunction.Title.Text = "Plot of function atan(x)"

	plotOfFunction.X.Label.Text = "x"
	plotOfFunction.Y.Label.Text = "y"

	plotLine, err := plotter.NewLine(pointsOfFunctionPlot)

	if err != nil {
		panic(err)
	}

	plotLine.LineStyle.Width = vg.Points(0.1)
	plotLine.Color = color.RGBA{R: 200, G: 100, B: 100}

	plotOfFunction.Add(plotLine)
	plotOfFunction.Legend.Add("atan(x)", plotLine)

	if err := plotOfFunction.Save(10*vg.Inch, 10*vg.Inch,
		"atan-function-plot-01.png"); err != nil {

		panic(err)
	}
}
