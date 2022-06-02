package topotypes

import (
	"encoding/json"
	"fmt"
)

type CompoundObject struct {
	Name      string         `json:"name"`
	Shape     ToposInterface `json:"shape"`
	Transform *TopoTransform `json:"transform"`
}

type CompoundGroup struct {
	Objects []string `json:"objects"`
	Type    string   `json:"type"`
}

func NewCompoundGroup(ty int) *CompoundGroup {
	return &CompoundGroup{Type: CompoundModeToString(ty)}
}

type TopoCompound struct {
	TopoMaker
	Objects []CompoundObject `json:"objects,omitempty"`
	Groups  []CompoundGroup  `json:"groups,omitempty"`
}

func NewCompound() *TopoCompound {
	t := &TopoCompound{}
	t.Type = TopoTypeToString(TOPO_TYPE_COMPOUND)
	return t
}

func CompoundUnMarshal(js []byte) (*TopoCompound, error) {
	tc := &TopoCompound{}
	e := json.Unmarshal(js, tc)
	if e != nil {
		return nil, e
	}
	if len(tc.Objects) == 0 {
		return tc, nil
	}

	for index, obj := range tc.Objects {
		bt, e := json.Marshal(obj.Shape)
		if e != nil {
			return nil, e
		}
		inter, e := TopoUnMarshal(bt)
		if e != nil {
			return nil, e
		}
		_, ok := inter.(TopoBoundy)
		if !ok {
			return nil, fmt.Errorf("index %d compound object must be Interface CompoundTopo", index)
		}
		tc.Objects[index].Shape = inter
	}
	return tc, nil
}
