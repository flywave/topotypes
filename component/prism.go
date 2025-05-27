package component

import (
	"encoding/json"

	"github.com/flywave/topotypes/material"
	"github.com/flywave/topotypes/profile"
)

type Prism struct {
	BaseComponent
	Profile   profile.Profile      `json:"profile,omitempty"`
	Direction [3]float64           `json:"direction"`
	Materials []*material.Material `json:"materials,omitempty"`
}

func PrismUnMarshal(js []byte) (*Prism, error) {
	p := Prism{}
	e := json.Unmarshal(js, &p)
	if e != nil {
		return nil, e
	}
	if p.Profile != nil {
		prof, er := profile.ProfileUnMarshal(p.Profile)
		if er != nil {
			return nil, er
		}
		p.Profile = prof
	}
	return &p, nil
}
