package material

func f(v float64) *float64 { return &v }

type PresetName string

const (
	PresetSteel           PresetName = "steel"
	PresetGalvanizedSteel PresetName = "galvanized_steel"
	PresetConcrete        PresetName = "concrete"
	PresetWood            PresetName = "wood"
	PresetGlass           PresetName = "glass"
	PresetBrick           PresetName = "brick"
	PresetPlastic         PresetName = "plastic"
	PresetAsphalt         PresetName = "asphalt"
	PresetRubber          PresetName = "rubber"
	PresetCeramic         PresetName = "ceramic"
	PresetCopper          PresetName = "copper"
	PresetAluminum        PresetName = "aluminum"
)

var presets = map[PresetName]*Material{
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
	PresetConcrete: {
		Name:             string(PresetConcrete),
		Type:             "pbr",
		Color:            &[3]byte{185, 185, 180},
		Metallic:         f(0.0),
		Roughness:        f(0.9),
		Reflectance:      f(0.02),
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
	PresetGlass: {
		Name:         string(PresetGlass),
		Type:         "pbr",
		Color:        &[3]byte{200, 220, 240},
		Transparency: 0.4,
		Metallic:     f(0.0),
		Roughness:    f(0.05),
		Reflectance:  f(0.5),
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
	PresetPlastic: {
		Name:             string(PresetPlastic),
		Type:             "pbr",
		Color:            &[3]byte{220, 220, 220},
		Metallic:         f(0.0),
		Roughness:        f(0.4),
		Reflectance:      f(0.2),
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
	PresetRubber: {
		Name:             string(PresetRubber),
		Type:             "pbr",
		Color:            &[3]byte{40, 40, 45},
		Metallic:         f(0.0),
		Roughness:        f(0.8),
		Reflectance:      f(0.02),
		AmbientOcclusion: f(1.0),
	},
	PresetCeramic: {
		Name:             string(PresetCeramic),
		Type:             "pbr",
		Color:            &[3]byte{240, 240, 245},
		Metallic:         f(0.0),
		Roughness:        f(0.15),
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
	PresetAluminum: {
		Name:             string(PresetAluminum),
		Type:             "pbr",
		Color:            &[3]byte{210, 210, 220},
		Metallic:         f(1.0),
		Roughness:        f(0.2),
		Reflectance:      f(0.7),
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
