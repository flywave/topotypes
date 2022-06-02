package topotypes

type TopoTransform struct {
	Rotation  *[4]float64 `json:"rotation"`
	Translate *[3]float64 `json:"translate"`
	Scale     *[3]float64 `json:"scale"`
}

func NewTopoTransform() *TopoTransform {
	return &TopoTransform{Scale: &[3]float64{1.0, 1.0, 1.0}}
}
