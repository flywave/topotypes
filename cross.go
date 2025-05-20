package topotypes

import "strings"

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

func (sp *TopoCrossMultiPoint) GetModel() string {
	mds := []string{}
	for _, obj := range sp.Objects {
		mds = append(mds, obj.Model)
	}
	return strings.Join(mds, ",")
}

func (sp *TopoCrossMultiPoint) SetModel(fileid string) {
	ids := strings.Split(fileid, ",")
	for i, obj := range sp.Objects {
		obj.Model = ids[i]
	}
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

func (sp *TopoCrossPoint) GetModel() string {
	return sp.Model
}

func (sp *TopoCrossPoint) SetModel(fileid string) {
	sp.Model = fileid
}
