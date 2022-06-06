package editor

import vec3d "github.com/flywave/go3d/float64/vec3"

type Anchor struct {
	Center vec3d.T `json:"center"`
	Name   string  `json:"name"`
	NameZh string  `json:"name_zh"`
	Normal vec3d.T `json:"normal"`
	Unit   float64 `json:"unit"`
}
