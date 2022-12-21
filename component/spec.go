package component

type SpecType string

const (
	GIMSPEC SpecType = "GIM"
)

type Spec struct {
	BaseComponent
	Type  SpecType    `json:"GIM"`
	Model interface{} `json:"model"`
}
