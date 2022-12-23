package component

type SpecType string

const (
	GIMSPEC SpecType = "GIM"
)

type Spec struct {
	BaseComponent
	Type  SpecType    `json:"type"`
	Model interface{} `json:"model"`
}
