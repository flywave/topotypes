package editor

import (
	"github.com/flywave/go3d/quaternion"
	"github.com/flywave/go3d/vec3"
)

type Metadata struct {
	Scale       float64      `json:"scale"`
	Rotation    quaternion.T `json:"rotation"`
	Offset      vec3.T       `json:"offset"`
	Anchors     []*Anchor    `json:"anchors"`
	AnchorCount int          `json:"anchorcount"`
	Boards      []*Board     `json:"boards"`
	BoardCount  int          `json:"boardcount"`
	Components  []*Component `json:"components,omitempty"`
}
