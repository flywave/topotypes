package topotypes

import "github.com/flywave/go3d/float64/vec3"

type Anchor struct {
	Center vec3.T  `json:"center"`
	Name   string  `json:"name"`
	NameZh string  `json:"name_zh"`
	Normal vec3.T  `json:"normal"`
	Unit   float64 `json:"unit"`
}

type TopoAnchorRef struct {
	Ref string `json:"ref"`
}

type TopoAnchorLink struct {
	Link           string `json:"link"`
	AnchorName     string `json:"anchor_name"`
	DestAnchorName string `json:"dest_anchor_name"`
}

type TopoAnchor struct {
	Name     string      `json:"name"`
	Position *[3]float64 `json:"position,omitempty"`
	Link     string      `json:"link,omitempty"`
}
