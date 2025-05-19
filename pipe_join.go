package topotypes

type TopoPipeJoin struct {
	TopoMaker
	CenterMode  string   `json:"mode"`
	In          []string `json:"in-pipe-ids,omitempty"`
	Out         []string `json:"out-pipe-ids,omitempty"`
	SegmentEdge bool     `json:"segment-edge"`
}

func NewTopoPipeJoin() *TopoPipeJoin {
	t := &TopoPipeJoin{}
	t.Type = TopoTypeToString(TOPO_TYPE_PIPE_JOIN)
	return t
}
