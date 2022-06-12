package topotypes

type LineWithAnchor interface {
	GetAnchor() [2]*TopoAnchor
	GetProfile() TopoProfile
}
