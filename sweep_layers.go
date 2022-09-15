package topotypes

import "encoding/json"

type Rectangle struct {
	Ul [3]float64 `json:"ul"`
	Ur [3]float64 `json:"ur"`
	Bl [3]float64 `json:"bl"`
	Br [3]float64 `json:"br"`
}

type TopoLayer struct {
	Name      string      `json:"name,omitempty"`
	Width     float32     `json:"width"`
	Height    float32     `json:"height"`
	Profile   TopoProfile `json:"profile"`
	Mtl       string      `json:"mtl,omitempty"`
	IsSurface bool        `json:"is_surface,omitempty"`
	LineIndex int         `json:"line_index"`
}

type TopoSweepLayers struct {
	TopoMaker
	Layers map[string][]*TopoLayer `json:"layers,omitempty"`
}

type LayerGroup struct {
	In  string `json:"in,omitempty"`
	Out string `json:"out,omitempty"`
}

type TopoSweepLayersIntersection struct {
	TopoMaker
	LayerGroups []*LayerGroup  `json:"sweep-lines"`
	Textures    map[int]string `json:"textures,omitempty"`
}

func SweepLayersUnMarshal(js []byte) (*TopoSweepLayers, error) {
	sl := TopoSweepLayers{}
	e := json.Unmarshal(js, &sl)
	if e != nil {
		return nil, e
	}

	for _, ly := range sl.Layers {
		for _, l := range ly {
			prof, er := ProfileUnMarshal(l.Profile)
			if er != nil {
				return nil, er
			}
			l.Profile = prof
		}
	}
	return &sl, nil
}

func (t *TopoSweepLayers) GetMaterialIds() []string {
	ids := []string{}
	for _, ly := range t.Layers {
		for _, l := range ly {
			ids = append(ids, l.Mtl)
		}
	}
	return ids
}
