package editor

import (
	"encoding/json"

	"github.com/flywave/topotypes/material"
	"github.com/flywave/topotypes/profile"
)

type Structure struct {
	Wire     [][][3]float64  `json:"wire"`
	Profile  profile.Profile `json:"profile"`
	Material string          `json:"mtl,omitempty"`
}

type SteelStructure struct {
	BaseComponent
	Structures []Structure          `json:"structures"`
	Materials  []*material.Material `json:"materials,omitempty"`
}

func SteelStructureUnMarshal(js []byte) (*SteelStructure, error) {
	ss := SteelStructure{}
	e := json.Unmarshal(js, &ss)
	if e != nil {
		return nil, e
	}
	for i := range ss.Structures {
		if ss.Structures[i].Profile != nil {
			prof, er := ProfileUnMarshal(ss.Structures[i].Profile)
			if er != nil {
				return nil, er
			}
			ss.Structures[i].Profile = prof
		}
	}
	return &ss, nil
}
