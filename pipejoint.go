package topotypes

import "encoding/json"

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

func (t *TopoPipeJoint) UnmarshalJSON(data []byte) error {
	stu := struct {
		Topos
		Materials  TopoMaterialMap `json:"materials,omitempty"`
		MaterialId string          `json:"mtl_id,omitempty"`
		CenterMode string          `json:"mode"`
		In         []string        `json:"in-pipe-ids,omitempty"`
		Out        []string        `json:"out-pipe-ids,omitempty"`
		Flanged    bool            `json:"flanged"`
		UpDir      *[3]float64     `json:"upDir,omitempty"`
	}{}

	if err := json.Unmarshal(data, &stu); err != nil {
		return err
	}
	t.Topos = stu.Topos
	t.Materials = stu.Materials
	t.MaterialId = stu.MaterialId
	t.CenterMode = stu.CenterMode
	t.In = stu.In
	t.Out = stu.Out
	t.Flanged = stu.Flanged
	t.UpDir = stu.UpDir
	return nil
}
