package topotypes

import "github.com/flywave/topotypes/catenary"

type TopoCatenary struct {
	TopoMaker
	Anchors [2]*TopoAnchor `json:"anchors"`
	catenary.Catenary
}
