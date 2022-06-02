package editor

import "github.com/flywave/topotypes/catenary"

type Catenary struct {
	BaseComponent
	catenary.Catenary
	P1        [3]float64 `json:"p1"`
	P2        [3]float64 `json:"p2"`
	Direction [3]float64 `json:"direction"`
}
