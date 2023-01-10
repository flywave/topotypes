package topotypes

type TopoAnchorRef struct {
	Ref string `json:"ref"`
}

type TopoAnchorLink struct {
	Link         string `json:"link"`
	AnchorId     string `json:"anchor_id"`
	DestAnchorId string `json:"dest_anchor_id"`
}

type TopoAnchor struct {
	Id       string     `json:"id"`
	Position [3]float64 `json:"position,omitempty"`
	Link     string     `json:"link,omitempty"`
}
