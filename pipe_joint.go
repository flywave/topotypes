package topotypes

type TopoPipeJoint struct {
	TopoParametric
	CenterMode string      `json:"mode"`
	In         []string    `json:"in-pipe-ids,omitempty"`
	Out        []string    `json:"out-pipe-ids,omitempty"`
	Flanged    bool        `json:"flanged"`
	UpDir      *[3]float64 `json:"upDir,omitempty"`
}

func (p *TopoPipeJoint) GetIns() []string {
	return p.In
}

func (p *TopoPipeJoint) GetOuts() []string {
	return p.Out
}

func NewTopoPipeJoint() *TopoPipeJoint {
	t := &TopoPipeJoint{}
	t.Type = TopoTypeToString(TOPO_TYPE_PIPE_JOINT)
	return t
}
