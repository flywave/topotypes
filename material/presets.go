package material

func f(v float64) *float64 { return &v }

type PresetName string

const (
	PresetSteel           PresetName = "steel"
	PresetGalvanizedSteel PresetName = "galvanized_steel"
	PresetCastIron        PresetName = "cast_iron"
	PresetConcrete        PresetName = "concrete"
	PresetMortar          PresetName = "mortar"
	PresetWood            PresetName = "wood"
	PresetGlass           PresetName = "glass"
	PresetBrick           PresetName = "brick"
	PresetPlastic         PresetName = "plastic"
	PresetPVC             PresetName = "pvc"
	PresetFRP             PresetName = "frp"
	PresetNylon           PresetName = "nylon"
	PresetRubber          PresetName = "rubber"
	PresetAsphalt         PresetName = "asphalt"
	PresetStone           PresetName = "stone"
	PresetGranite         PresetName = "granite"
	PresetSandstone       PresetName = "sandstone"
	PresetSoil            PresetName = "soil"
	PresetClay            PresetName = "clay"
	PresetCeramic         PresetName = "ceramic"
	PresetPorcelain       PresetName = "porcelain"
	PresetCopper          PresetName = "copper"
	PresetBronze          PresetName = "bronze"
	PresetBrass           PresetName = "brass"
	PresetAluminum        PresetName = "aluminum"
	PresetZinc            PresetName = "zinc"
	PresetTitanium        PresetName = "titanium"
	PresetLead            PresetName = "lead"
	PresetRedPaint        PresetName = "red_paint"
	PresetYellowPaint     PresetName = "yellow_paint"
	PresetGreenPaint      PresetName = "green_paint"
	PresetBluePaint       PresetName = "blue_paint"
	PresetOrangePaint     PresetName = "orange_paint"
	PresetWhitePaint      PresetName = "white_paint"
	PresetBlackPaint      PresetName = "black_paint"
	PresetGreyPaint       PresetName = "grey_paint"
)

var presets = map[PresetName]*Material{
	// ── Metals ──
	PresetSteel: {
		Name:             string(PresetSteel),
		Type:             "pbr",
		Color:            &[3]byte{180, 185, 195},
		Metallic:         f(1.0),
		Roughness:        f(0.4),
		Reflectance:      f(0.5),
		AmbientOcclusion: f(1.0),
	},
	PresetGalvanizedSteel: {
		Name:             string(PresetGalvanizedSteel),
		Type:             "pbr",
		Color:            &[3]byte{200, 208, 215},
		Metallic:         f(1.0),
		Roughness:        f(0.25),
		Reflectance:      f(0.6),
		AmbientOcclusion: f(1.0),
	},
	PresetCastIron: {
		Name:             string(PresetCastIron),
		Type:             "pbr",
		Color:            &[3]byte{120, 125, 130},
		Metallic:         f(1.0),
		Roughness:        f(0.6),
		Reflectance:      f(0.3),
		AmbientOcclusion: f(1.0),
	},
	PresetCopper: {
		Name:             string(PresetCopper),
		Type:             "pbr",
		Color:            &[3]byte{210, 140, 90},
		Metallic:         f(1.0),
		Roughness:        f(0.3),
		Reflectance:      f(0.6),
		AmbientOcclusion: f(1.0),
	},
	PresetBronze: {
		Name:             string(PresetBronze),
		Type:             "pbr",
		Color:            &[3]byte{195, 140, 75},
		Metallic:         f(1.0),
		Roughness:        f(0.35),
		Reflectance:      f(0.5),
		AmbientOcclusion: f(1.0),
	},
	PresetBrass: {
		Name:             string(PresetBrass),
		Type:             "pbr",
		Color:            &[3]byte{200, 175, 90},
		Metallic:         f(1.0),
		Roughness:        f(0.25),
		Reflectance:      f(0.6),
		AmbientOcclusion: f(1.0),
	},
	PresetAluminum: {
		Name:             string(PresetAluminum),
		Type:             "pbr",
		Color:            &[3]byte{210, 210, 220},
		Metallic:         f(1.0),
		Roughness:        f(0.2),
		Reflectance:      f(0.7),
		AmbientOcclusion: f(1.0),
	},
	PresetZinc: {
		Name:             string(PresetZinc),
		Type:             "pbr",
		Color:            &[3]byte{190, 200, 210},
		Metallic:         f(1.0),
		Roughness:        f(0.15),
		Reflectance:      f(0.7),
		AmbientOcclusion: f(1.0),
	},
	PresetTitanium: {
		Name:             string(PresetTitanium),
		Type:             "pbr",
		Color:            &[3]byte{160, 165, 170},
		Metallic:         f(1.0),
		Roughness:        f(0.3),
		Reflectance:      f(0.5),
		AmbientOcclusion: f(1.0),
	},
	PresetLead: {
		Name:             string(PresetLead),
		Type:             "pbr",
		Color:            &[3]byte{100, 105, 115},
		Metallic:         f(1.0),
		Roughness:        f(0.7),
		Reflectance:      f(0.2),
		AmbientOcclusion: f(1.0),
	},

	// ── Construction Materials ──
	PresetConcrete: {
		Name:             string(PresetConcrete),
		Type:             "pbr",
		Color:            &[3]byte{185, 185, 180},
		Metallic:         f(0.0),
		Roughness:        f(0.9),
		Reflectance:      f(0.02),
		AmbientOcclusion: f(1.0),
	},
	PresetMortar: {
		Name:             string(PresetMortar),
		Type:             "pbr",
		Color:            &[3]byte{195, 190, 180},
		Metallic:         f(0.0),
		Roughness:        f(0.85),
		Reflectance:      f(0.02),
		AmbientOcclusion: f(1.0),
	},
	PresetBrick: {
		Name:             string(PresetBrick),
		Type:             "pbr",
		Color:            &[3]byte{185, 85, 65},
		Metallic:         f(0.0),
		Roughness:        f(0.9),
		Reflectance:      f(0.03),
		AmbientOcclusion: f(1.0),
	},
	PresetWood: {
		Name:             string(PresetWood),
		Type:             "pbr",
		Color:            &[3]byte{165, 125, 85},
		Metallic:         f(0.0),
		Roughness:        f(0.75),
		Reflectance:      f(0.04),
		AmbientOcclusion: f(1.0),
	},
	PresetAsphalt: {
		Name:             string(PresetAsphalt),
		Type:             "pbr",
		Color:            &[3]byte{60, 60, 65},
		Metallic:         f(0.0),
		Roughness:        f(0.95),
		Reflectance:      f(0.01),
		AmbientOcclusion: f(1.0),
	},
	PresetStone: {
		Name:             string(PresetStone),
		Type:             "pbr",
		Color:            &[3]byte{170, 165, 155},
		Metallic:         f(0.0),
		Roughness:        f(0.95),
		Reflectance:      f(0.01),
		AmbientOcclusion: f(1.0),
	},
	PresetGranite: {
		Name:             string(PresetGranite),
		Type:             "pbr",
		Color:            &[3]byte{150, 145, 140},
		Metallic:         f(0.0),
		Roughness:        f(0.92),
		Reflectance:      f(0.01),
		AmbientOcclusion: f(1.0),
	},
	PresetSandstone: {
		Name:             string(PresetSandstone),
		Type:             "pbr",
		Color:            &[3]byte{195, 180, 155},
		Metallic:         f(0.0),
		Roughness:        f(0.93),
		Reflectance:      f(0.01),
		AmbientOcclusion: f(1.0),
	},
	PresetSoil: {
		Name:             string(PresetSoil),
		Type:             "pbr",
		Color:            &[3]byte{130, 110, 75},
		Metallic:         f(0.0),
		Roughness:        f(1.0),
		Reflectance:      f(0.0),
		AmbientOcclusion: f(1.0),
	},
	PresetClay: {
		Name:             string(PresetClay),
		Type:             "pbr",
		Color:            &[3]byte{160, 120, 85},
		Metallic:         f(0.0),
		Roughness:        f(0.95),
		Reflectance:      f(0.01),
		AmbientOcclusion: f(1.0),
	},

	// ── Polymers & Insulation ──
	PresetPlastic: {
		Name:             string(PresetPlastic),
		Type:             "pbr",
		Color:            &[3]byte{220, 220, 220},
		Metallic:         f(0.0),
		Roughness:        f(0.4),
		Reflectance:      f(0.2),
		AmbientOcclusion: f(1.0),
	},
	PresetPVC: {
		Name:             string(PresetPVC),
		Type:             "pbr",
		Color:            &[3]byte{200, 200, 195},
		Metallic:         f(0.0),
		Roughness:        f(0.5),
		Reflectance:      f(0.1),
		AmbientOcclusion: f(1.0),
	},
	PresetFRP: {
		Name:             string(PresetFRP),
		Type:             "pbr",
		Color:            &[3]byte{185, 175, 160},
		Metallic:         f(0.0),
		Roughness:        f(0.6),
		Reflectance:      f(0.05),
		AmbientOcclusion: f(1.0),
	},
	PresetNylon: {
		Name:             string(PresetNylon),
		Type:             "pbr",
		Color:            &[3]byte{230, 230, 235},
		Metallic:         f(0.0),
		Roughness:        f(0.5),
		Reflectance:      f(0.1),
		AmbientOcclusion: f(1.0),
	},
	PresetRubber: {
		Name:             string(PresetRubber),
		Type:             "pbr",
		Color:            &[3]byte{40, 40, 45},
		Metallic:         f(0.0),
		Roughness:        f(0.8),
		Reflectance:      f(0.02),
		AmbientOcclusion: f(1.0),
	},

	// ── Ceramics & Glass ──
	PresetCeramic: {
		Name:             string(PresetCeramic),
		Type:             "pbr",
		Color:            &[3]byte{240, 240, 245},
		Metallic:         f(0.0),
		Roughness:        f(0.15),
		Reflectance:      f(0.3),
		AmbientOcclusion: f(1.0),
	},
	PresetPorcelain: {
		Name:             string(PresetPorcelain),
		Type:             "pbr",
		Color:            &[3]byte{235, 230, 225},
		Metallic:         f(0.0),
		Roughness:        f(0.1),
		Reflectance:      f(0.35),
		AmbientOcclusion: f(1.0),
	},
	PresetGlass: {
		Name:         string(PresetGlass),
		Type:         "pbr",
		Color:        &[3]byte{200, 220, 240},
		Transparency: 0.4,
		Metallic:     f(0.0),
		Roughness:    f(0.05),
		Reflectance:  f(0.5),
	},

	// ── Paint / Coating ──
	PresetRedPaint: {
		Name:             string(PresetRedPaint),
		Type:             "pbr",
		Color:            &[3]byte{200, 50, 50},
		Metallic:         f(0.0),
		Roughness:        f(0.5),
		Reflectance:      f(0.15),
		AmbientOcclusion: f(1.0),
	},
	PresetYellowPaint: {
		Name:             string(PresetYellowPaint),
		Type:             "pbr",
		Color:            &[3]byte{230, 200, 30},
		Metallic:         f(0.0),
		Roughness:        f(0.5),
		Reflectance:      f(0.2),
		AmbientOcclusion: f(1.0),
	},
	PresetGreenPaint: {
		Name:             string(PresetGreenPaint),
		Type:             "pbr",
		Color:            &[3]byte{50, 170, 70},
		Metallic:         f(0.0),
		Roughness:        f(0.5),
		Reflectance:      f(0.15),
		AmbientOcclusion: f(1.0),
	},
	PresetBluePaint: {
		Name:             string(PresetBluePaint),
		Type:             "pbr",
		Color:            &[3]byte{40, 100, 200},
		Metallic:         f(0.0),
		Roughness:        f(0.5),
		Reflectance:      f(0.15),
		AmbientOcclusion: f(1.0),
	},
	PresetOrangePaint: {
		Name:             string(PresetOrangePaint),
		Type:             "pbr",
		Color:            &[3]byte{230, 130, 30},
		Metallic:         f(0.0),
		Roughness:        f(0.5),
		Reflectance:      f(0.15),
		AmbientOcclusion: f(1.0),
	},
	PresetWhitePaint: {
		Name:             string(PresetWhitePaint),
		Type:             "pbr",
		Color:            &[3]byte{245, 245, 240},
		Metallic:         f(0.0),
		Roughness:        f(0.5),
		Reflectance:      f(0.2),
		AmbientOcclusion: f(1.0),
	},
	PresetBlackPaint: {
		Name:             string(PresetBlackPaint),
		Type:             "pbr",
		Color:            &[3]byte{45, 45, 45},
		Metallic:         f(0.0),
		Roughness:        f(0.6),
		Reflectance:      f(0.05),
		AmbientOcclusion: f(1.0),
	},
	PresetGreyPaint: {
		Name:             string(PresetGreyPaint),
		Type:             "pbr",
		Color:            &[3]byte{160, 160, 160},
		Metallic:         f(0.0),
		Roughness:        f(0.5),
		Reflectance:      f(0.2),
		AmbientOcclusion: f(1.0),
	},
}

func GetPreset(name PresetName) *Material {
	m, ok := presets[name]
	if !ok {
		return nil
	}
	cp := *m
	return &cp
}

func PresetNames() []PresetName {
	ns := make([]PresetName, 0, len(presets))
	for n := range presets {
		ns = append(ns, n)
	}
	return ns
}
