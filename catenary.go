package topotypes

import "github.com/flywave/topotypes/catenary"

type TopoCatenary struct {
	Topos
	Anchors [2]*TopoAnchor `json:"anchors"`
	catenary.Catenary
}
