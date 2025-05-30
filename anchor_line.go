package topotypes

import "github.com/flywave/topotypes/anchor"

type LineWithAnchor interface {
	GetAnchor() [2]*anchor.TopoAnchor
	GetProfile() TopoProfile
}
