package component

type AnchorLinkType string

const (
	ANCHOR_LINK_LINKED   AnchorLinkType = "linked"
	ANCHOR_LINK_DISABLED AnchorLinkType = "disabled"
	ANCHOR_LINK_OPENED   AnchorLinkType = "opened"
)

type ComponentAnchorLink struct {
	AnchorId     string         `json:"anchor_id"`
	Link         string         `json:"link"`
	DestAnchorId string         `json:"dest_anchor_id"`
	Type         AnchorLinkType `json:"type"`
}

type Model struct {
	BaseComponent
	Root      bool                  `json:"root"`
	Model     string                `json:"model"`
	Instanced bool                  `json:"instanced"`
	Links     []ComponentAnchorLink `json:"links,omitempty"`
}
