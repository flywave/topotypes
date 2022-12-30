package component

import (
	"encoding/json"
	"errors"

	"github.com/flywave/topotypes/profile"
)

func ProfileUnMarshal(inter interface{}) (interface{}, error) {
	switch pro := inter.(type) {
	case map[string]interface{}:
		v, ok := pro["type"]
		t, ok2 := v.(string)
		if !ok || !ok2 {
			return nil, errors.New("profile type error")
		}
		pro_t := profile.StringToProfileType(t)
		js, er := json.Marshal(inter)
		if er != nil {
			return nil, er
		}
		var pf interface{}
		switch pro_t {
		case profile.TYPE_TRIANGLE:
			pf = profile.NewTriangle()
		case profile.TYPE_RECTANGLE:
			pf = profile.NewRectangle()
		case profile.TYPE_CIRC:
			pf = profile.NewCirc()
		case profile.TYPE_ELIPS:
			pf = profile.NewElips()
		case profile.TYPE_POLYGON:
			pf = profile.NewPolygon()
		case profile.TYPE_L_STEEL:
			pf = profile.NewLShape()
		default:
			return nil, errors.New("profile type error")
		}
		e := json.Unmarshal(([]byte)(js), pf)
		if e != nil {
			return nil, e
		}
		return pf, nil
	default:
		return nil, errors.New("profile type error")
	}
}
