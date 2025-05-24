package catenary

import "github.com/flywave/topotypes/profile"

type Catenary struct {
	P1           [2]float64      `json:"p1"`
	P2           [2]float64      `json:"p2"`
	Profile      profile.Profile `json:"profile"`
	Slack        float64         `json:"slack"`
	MaxSag       float64         `json:"maxSag"`
	Tessellation float32         `json:"tessellation"`
	UpDir        *[3]float64     `json:"upDir"`
}
