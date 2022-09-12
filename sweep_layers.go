package topotypes

import "encoding/json"

type Rectangle struct {
	Ul [3]float64 `json:"ul"`
	Ur [3]float64 `json:"ur"`
	Bl [3]float64 `json:"bl"`
	Br [3]float64 `json:"br"`
}

type TopoLayer struct {
	Name     string      `json:"name,omitempty"`
	Width    float32     `json:"width"`
	Height   float32     `json:"height"`
	Profile  TopoProfile `json:"profile"`
	Finished bool        `json:"finished"`
	// Boolean string      `json:"boolean,omitempty"`
}

type TopoSweepLayers struct {
	TopoMaker
	SideLayers   []*TopoLayer `json:"side-layers,omitempty"`
	CenterLayers []*TopoLayer `json:"center-layers,omitempty"`
	Decals       []*TopoDecal `json:"decals,omitempty"`
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

	for _, l := range sl.SideLayers {
		prof, er := ProfileUnMarshal(l.Profile)
		if er != nil {
			return nil, er
		}
		l.Profile = prof
	}

	for _, l := range sl.CenterLayers {
		prof, er := ProfileUnMarshal(l.Profile)
		if er != nil {
			return nil, er
		}
		l.Profile = prof
	}
	return &sl, nil
}
