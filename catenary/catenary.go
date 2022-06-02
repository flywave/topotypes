package catenary

import "github.com/flywave/topotypes/profile"

type Catenary struct {
	Profile      profile.Profile `json:"profile"`
	Slack        float64         `json:"slack"`
	MaxSag       float64         `json:"max_sag"`
	Tessellation float64         `json:"tessellation"`
}
