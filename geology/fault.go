package geology

type Fault struct {
	ID        string  `json:"id"`     // 断层编号
	Name      string  `json:"name"`   // 断层名称
	FaultType string  `json:"type"`   // 断层类型(正断层/逆断层等)
	Strike    float64 `json:"strike"` // 断层倾向(°)
	Dip       float64 `json:"dip"`    // 断层倾角(°)
	Throw     float64 `json:"throw"`  // 断层落差(m)
}
