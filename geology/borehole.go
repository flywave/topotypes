package geology

type BoreholeSample struct {
	Name      string                 `json:"name,omitempty"`
	DepthFrom float64                `json:"depthFrom"`
	DepthTo   float64                `json:"depthTo"`
	MTL       string                 `json:"mtl,omitempty"`
	Property  map[string]interface{} `json:"property,omitempty"`
}

type Borehole struct {
	Version  int               `json:"version"`
	Samples  []*BoreholeSample `json:"samples"`
	Diameter *float64          `json:"diameter,omitempty"`
}

func (t *Borehole) GetSamples() []*BoreholeSample {
	return t.Samples
}
