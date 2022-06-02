package topotypes

type TopoAnchorRef struct {
	Ref string `json:"ref"`
}

type TopoAnchorLink struct {
	Link           string `json:"link"`
	AnchorName     string `json:"anchor_name"`
	DestAnchorName string `json:"dest_anchor_name"`
}

type TopoAnchor struct {
	Name     string      `json:"name"`
	Position *[3]float64 `json:"position,omitempty"`
	Link     string      `json:"link,omitempty"`
}
