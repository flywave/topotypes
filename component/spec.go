package component

type SpecType string

const (
	GIMSPEC SpecType = "GIM"
)

type Spec struct {
	BaseComponent
	Spec  SpecType    `json:"spec"`
	Model interface{} `json:"model"`
}
