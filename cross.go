package topotypes

type TopoCrossMultiPoint struct {
	Topos
	Refs    []TopoAnchorRef   `json:"links,omitempty"`
	Objects []*TopoCrossPoint `json:"objects,omitempty"`
}

func NewTopoCrossMultiPoint() *TopoCrossMultiPoint {
	t := &TopoCrossMultiPoint{}
	t.Type = TopoTypeToString(TOPO_TYPE_CROSS_MULTI_POINT)
	return t
}

type TopoCrossPoint struct {
	Topos
	Model     string           `json:"model"`
	Instanced bool             `json:"instanced"`
	Links     []TopoAnchorLink `json:"links,omitempty"`
}

func NewTopoCrossPoint() *TopoCrossPoint {
	t := &TopoCrossPoint{}
	t.Type = TopoTypeToString(TOPO_TYPE_CROSS_POINT)
	return t
}
