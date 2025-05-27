package component

import (
	"encoding/json"

	"github.com/flywave/topotypes/catenary"
	"github.com/flywave/topotypes/material"
	"github.com/flywave/topotypes/profile"
)

type Catenary struct {
	BaseComponent
	catenary.Catenary
	Materials []*material.Material `json:"materials,omitempty"`
	P1        [3]float64           `json:"p1"`
	P2        [3]float64           `json:"p2"`
	Direction [3]float64           `json:"direction"`
}

func CatenaryUnMarshal(js []byte) (*Catenary, error) {
	catenary := Catenary{}
	e := json.Unmarshal(js, &catenary)
	if e != nil {
		return nil, e
	}
	if catenary.Profile != nil {
		prof, er := profile.ProfileUnMarshal(catenary.Profile)
		if er != nil {
			return nil, er
		}
		catenary.Profile = prof
	}
	return &catenary, nil
}
