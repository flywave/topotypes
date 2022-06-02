package editor

type Component interface{}

type BaseComponent struct {
	Name      string `json:"name"`
	Transform *struct {
		Rotation  *[4]float64 `json:"rotation,omitempty"`
		Translate *[3]float64 `json:"translate,omitempty"`
		Scale     *[3]float64 `json:"scale,omitempty"`
	} `json:"transform,omitempty"`
	BBox *[2][3]float64 `json:"bbox,omitempty"`
}
