package editor

import (
	"encoding/json"

	"github.com/flywave/topotypes/catenary"
)

type Catenary struct {
	BaseComponent
	catenary.Catenary
	P1        [3]float64 `json:"p1"`
	P2        [3]float64 `json:"p2"`
	Direction [3]float64 `json:"direction"`
}

func CatenaryUnMarshal(js []byte) (*Catenary, error) {
	catenary := Catenary{}
	e := json.Unmarshal(js, &catenary)
	if e != nil {
		return nil, e
	}
	if catenary.Profile != nil {
		prof, er := ProfileUnMarshal(catenary.Profile)
		if er != nil {
			return nil, er
		}
		catenary.Profile = prof
	}
	return &catenary, nil
}
