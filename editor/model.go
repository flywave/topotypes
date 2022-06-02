package editor

type AnchorLinkType string

const (
	ANCHOR_LINK_LINKED   = "linked"
	ANCHOR_LINK_DISABLED = "disabled"
	ANCHOR_LINK_OPENED   = "opened"
)

type ComponentAnchorLink struct {
	AnchorName     string         `json:"anchor_name"`
	Link           string         `json:"link"`
	DestAnchorName string         `json:"dest_anchor_name"`
	Type           AnchorLinkType `json:"type"`
}

type Model struct {
	BaseComponent
	Model string                `json:"model"`
	Links []ComponentAnchorLink `json:"links,omitempty"`
}
