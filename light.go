package topotypes

import (
	"encoding/json"
	"errors"
)

type LightInterface interface {
	GetLight() string
}

type TopoLight struct {
	Topos
	Light      string  `json:"light"`
	Color      []uint8 `json:"color"`
	Instensity float64 `json:"instensity"`
}

func (lt *TopoLight) GetLight() string {
	return lt.Light
}

type TopoAreaLight struct {
	TopoLight
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func NewTopoAreaLight() *TopoAreaLight {
	l := TopoAreaLight{}
	l.Light = LightTypeToString(TOPO_LIGHT_MODE_AREA)
	l.Type = TopoTypeToString(TOPO_TYPE_LIGHT)
	return &l
}

type TopoDirectionalLight struct {
	TopoLight
	Dir          [3]float64 `json:"dir"`
	Shadow       int        `json:"shadow"`
	Strength     *float64   `json:"strength,omitempty"`
	Bias         *float64   `json:"bias,omitempty"`
	Softness     *float64   `json:"softness,omitempty"`
	SoftnessFade *float64   `json:"softness-fade,omitempty"`
}

func NewTopoDirectionalLight() *TopoDirectionalLight {
	l := TopoDirectionalLight{}
	l.Light = LightTypeToString(TOPO_LIGHT_MODE_DIRECTIONAL)
	l.Type = TopoTypeToString(TOPO_TYPE_LIGHT)
	return &l
}

type TopoPointLight struct {
	TopoLight
	Distance float64 `json:"distance"`
}

func NewTopoPointLight() *TopoPointLight {
	l := TopoPointLight{}
	l.Light = LightTypeToString(TOPO_LIGHT_MODE_POINT)
	l.Type = TopoTypeToString(TOPO_TYPE_LIGHT)
	return &l
}

type TopoSpotLight struct {
	TopoLight
	Dir      [3]float64 `json:"dir"`
	Angle    float64    `json:"angle"`
	Exponent float64    `json:"exponent"`
	Distance float64    `json:"distance"`
}

func NewTopoSpotLight() *TopoSpotLight {
	l := TopoSpotLight{}
	l.Light = LightTypeToString(TOPO_LIGHT_MODE_SPOT)
	l.Type = TopoTypeToString(TOPO_TYPE_LIGHT)
	return &l
}

func LightUnMarshal(js []byte) (ToposInterface, error) {
	lt := TopoLight{}
	e := json.Unmarshal(js, &lt)
	if e != nil {
		return nil, e
	}
	var inter ToposInterface
	ty := StringToLightType(lt.Light)
	switch ty {
	case TOPO_LIGHT_MODE_SPOT:
		inter = NewTopoSpotLight()
	case TOPO_LIGHT_MODE_POINT:
		inter = NewTopoPointLight()
	case TOPO_LIGHT_MODE_DIRECTIONAL:
		inter = NewTopoDirectionalLight()
	case TOPO_LIGHT_MODE_AREA:
		inter = NewTopoAreaLight()
	default:
		return nil, errors.New("not support topo type")
	}
	e = json.Unmarshal(([]byte)(js), inter)
	if e != nil {
		return nil, e
	}
	return inter, nil
}
