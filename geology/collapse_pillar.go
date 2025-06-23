package geology

type CollapsePillar struct {
	ID         string  `json:"id"`          // 陷落柱ID
	Name       string  `json:"name"`        // 陷落柱名称
	TopRadius  float64 `json:"top_radius"`  // 顶面半径
	BaseRadius float64 `json:"base_radius"` // 底面半径
	Height     float64 `json:"height"`      // 柱体高度
	StratumID  string  `json:"stratum_id"`  // 所属地层
	Lithology  string  `json:"lithology"`   // 充填岩性
}
