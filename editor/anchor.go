package editor

import "github.com/flywave/go3d/vec3"

type Anchor struct {
	Center vec3.T  `json:"center"`
	Name   string  `json:"name"`
	NameZh string  `json:"name_zh"`
	Normal vec3.T  `json:"normal"`
	Unit   float64 `json:"unit"`
}
