package topotypes

type BoundPath struct {
	Radius float32 `json:"radius"`
}

type BoundPolygon struct {
	Height float32 `json:"height"`
}

type BoundSphere struct {
	Center [3]float32 `json:"center"`
	Radius float32    `json:"radius"`
}

type FeatureBound struct {
	BBox     *[2][3]float32 `json:"bbox3d,omitempty"`
	BSphere  *BoundSphere   `json:"sphere3d,omitempty"`
	BPath    *BoundPath     `json:"bpath,omitempty"`
	BPolygon *BoundPolygon  `json:"bpolygon,omitempty"`
}

type TopoFeature struct {
	Topos
	Bounds []FeatureBound `json:"bounds,omitempty"`
}

func (sp *TopoFeature) GetModel() string {
	return ""
}
