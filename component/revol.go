package component

import (
	"encoding/json"

	"github.com/flywave/topotypes/material"
	"github.com/flywave/topotypes/profile"
)

type Revol struct {
	BaseComponent
	Profile   profile.Profile      `json:"profile,omitempty"`
	Axis      [2][3]float64        `json:"axis"`
	Angle     float64              `json:"angle"`
	Materials []*material.Material `json:"materials,omitempty"`
}

func RevolUnmarshal(js []byte) (*Revol, error) {
	r := Revol{}
	e := json.Unmarshal(js, &r)
	if e != nil {
		return nil, e
	}
	if r.Profile != nil {
		prof, er := profile.ProfileUnmarshal(r.Profile)
		if er != nil {
			return nil, er
		}
		r.Profile = prof
	}
	return &r, nil
}
