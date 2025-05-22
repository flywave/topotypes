package unsupport

// import (
// 	"encoding/json"

// 	"github.com/flywave/topotypes/material"
// )

// type Pole struct {
// 	BaseComponent
// 	Radius    float64            `json:"radius"`
// 	Thickness float64            `json:"thickness"`
// 	P1        [3]float64         `json:"p1"`
// 	P2        [3]float64         `json:"p2"`
// 	Material  *material.Material `json:"material,omitempty"`
// }

// func PoleUnMarshal(js []byte) (*Pole, error) {
// 	p := Pole{}
// 	e := json.Unmarshal(js, &p)
// 	if e != nil {
// 		return nil, e
// 	}
// 	return &p, nil
// }
