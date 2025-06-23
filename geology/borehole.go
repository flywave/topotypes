package geology

type BoreholeSample struct {
	ID          string  `json:"id"`          // 地层编号
	X           float64 `json:"x"`           // 三维坐标
	Y           float64 `json:"y"`           // 三维坐标
	Z           float64 `json:"z"`           // 三维坐标
	Lithology   string  `json:"lithology"`   // 岩性描述
	Top         float64 `json:"top"`         // 顶板相对井口深度
	Base        float64 `json:"base"`        // 底板相对井口深度
	Depth       float64 `json:"depth"`       // 测深
	Azimuth     float64 `json:"azimuth"`     // 方位角
	Inclination float64 `json:"inclination"` // 倾角
}

type Borehole struct {
	Version         int                    `json:"version"`
	Elevation       float64                `json:"elevation"`       // 井口高程
	ProspectingLine string                 `json:"prospectingLine"` // 勘探线ID
	Index           int                    `json:"index"`           // 勘探线内索引
	Depth           float64                `json:"depth"`           // 完井深度
	Azimuth         float64                `json:"azimuth"`         // 方位角(0-360度)
	Inclination     float64                `json:"inclination"`     // 倾角(0-90度)
	Samples         []*BoreholeSample      `json:"samples"`
	Metadata        map[string]interface{} `json:"metadata"` // 扩展元数据
}

func (t *Borehole) GetSamples() []*BoreholeSample {
	return t.Samples
}
