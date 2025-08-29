package component

type Decal struct {
	BaseComponent
	Size    [2]float64 `json:"size"`
	Depth   float64    `json:"depth"`
	Texture string     `json:"texture,omitempty"`
}

func (d *Decal) GetTexture() []string {
	return []string{d.Texture}
}
