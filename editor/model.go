package editor

type AnchorLinkType string

const (
	ANCHOR_LINK_LINKED   AnchorLinkType = "linked"
	ANCHOR_LINK_DISABLED AnchorLinkType = "disabled"
	ANCHOR_LINK_OPENED   AnchorLinkType = "opened"
)

type ComponentAnchorLink struct {
	AnchorName     string         `json:"anchor_name"`
	Link           string         `json:"link"`
	DestAnchorName string         `json:"dest_anchor_name"`
	Type           AnchorLinkType `json:"type"`
}

type Model struct {
	BaseComponent
	Root  bool                  `json:"root"`
	Model string                `json:"model"`
	Links []ComponentAnchorLink `json:"links,omitempty"`
}
