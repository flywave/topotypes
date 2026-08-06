package railway

// 布局级数据契约: 镜像 go-topo ocs_layout.go / yard_layout.go
// 辅助算法生成 → 可序列化存库/前端编辑校正 → 确定性再生成
// 单位: 长度 mm, 里程 m

// =========================================================================
// 接触网锚段布局 (AnchorSection)
// =========================================================================

// AnchorSectionSpec 锚段输入参数 (AnchorSectionInput 的可序列化形式)
type AnchorSectionSpec struct {
	Centerline      [][3]float64 `json:"centerline"`      // 线路中心线
	ContactHeight   float64      `json:"contactHeight"`   // 导高, 默认 5300
	StructureHeight float64      `json:"structureHeight"` // 结构高度, 默认 1400
	SpanLength      float64      `json:"spanLength"`      // 标准跨距, 默认 50000
	MastHeight      float64      `json:"mastHeight"`      // 柱高, 默认 8000
	MastType        int          `json:"mastType"`        // 1-钢柱, 2-混凝土柱
	SideOffset      float64      `json:"sideOffset"`      // 侧面限界 CX, 默认 2900
	MastSide        float64      `json:"mastSide"`        // 支柱在线路左侧(+1)/右侧(-1), 默认 +1
	HasCompensator  bool         `json:"hasCompensator"`  // 两端锚柱设补偿装置
	ContactWireDia  float64      `json:"contactWireDia"`  // 接触线直径, 默认 12.9
	MessengerDia    float64      `json:"messengerDia"`    // 承力索直径, 默认 13.5
	DropperSpacing  float64      `json:"dropperSpacing"`  // 吊弦间距, 默认 8000
}

// AnchorMastLayout 单根支柱的布局数据
// 支柱局部约定: 原点=柱底中心, +X 指向线路 (Direction)
type AnchorMastLayout struct {
	Mileage        float64    `json:"mileage"`        // 里程 (m)
	Position       [3]float64 `json:"position"`       // 柱底中心 (世界坐标)
	Direction      [3]float64 `json:"direction"`      // 支柱朝向 (+X 指向线路, 水平单位向量)
	MastHeight     float64    `json:"mastHeight"`     // 柱高
	ContactWireZ   float64    `json:"contactWireZ"`   // 接触线相对轨面高
	MessengerWireZ float64    `json:"messengerWireZ"` // 承力索相对轨面高
	Stagger        float64    `json:"stagger"`        // 拉出值 (线路左侧为正)
	IsTensionMast  bool       `json:"isTensionMast"`  // 是否锚端柱 (带补偿装置)
	ContactPoint   [3]float64 `json:"contactPoint"`   // 接触线悬挂点 (世界坐标, 由 position/direction/stagger 推导的缓存)
	MessengerPoint [3]float64 `json:"messengerPoint"` // 承力索悬挂点 (同上)
}

// AnchorDropperLayout 单根吊弦: T 为跨内参数 (0,1), Top/Bottom/Length 为计算缓存
// 再生成时按 T + 所在跨两端悬挂点/弛度重算, 编辑柱位/拉出值后吊弦自动跟随
type AnchorDropperLayout struct {
	T      float64    `json:"t"`
	Top    [3]float64 `json:"top"`    // 承力索侧挂点
	Bottom [3]float64 `json:"bottom"` // 接触线侧挂点
	Length float64    `json:"length"`
}

// AnchorSpanLayout 相邻两柱间一跨的布局数据
type AnchorSpanLayout struct {
	FromMast     int                   `json:"fromMast"`     // 起始柱索引
	ToMast       int                   `json:"toMast"`       // 终止柱索引
	Length       float64               `json:"length"`       // 跨距 (接触线悬挂点间距)
	ContactSag   float64               `json:"contactSag"`   // 接触线弛度
	MessengerSag float64               `json:"messengerSag"` // 承力索弛度
	Droppers     []AnchorDropperLayout `json:"droppers"`     // 空则再生成时按 Spec.DropperSpacing 自动布置
}

// AnchorSection 生成一个锚段所需的全部中间数据 (编辑器存库/前端编辑的数据契约)
type AnchorSection struct {
	Base
	Spec  AnchorSectionSpec  `json:"spec"`
	Masts []AnchorMastLayout `json:"masts"`
	Spans []AnchorSpanLayout `json:"spans"`
}

func NewAnchorSection() *AnchorSection {
	return &AnchorSection{
		Base: Base{Type: "RAILWAY/AnchorSection"},
	}
}

// =========================================================================
// 站场布局 (Yard)
// =========================================================================

// TrackGeoProperties 股道生成参数 (轨距/超高/钢轨/轨枕/道床)
type TrackGeoProperties struct {
	Gauge            float64 `json:"gauge"`            // 轨距(mm), 默认 1435
	RailType         float64 `json:"railType"`         // 钢轨类型 kg/m: 43/50/60/75, 默认 60
	SuperElevation   float64 `json:"superElevation"`   // 超高(mm), 默认 0
	CoordScale       float64 `json:"coordScale"`       // 坐标→mm 缩放, 默认 1000 (米输入)
	SleeperLength    float64 `json:"sleeperLength"`    // 轨枕长(mm), 默认 2600
	SleeperWidth     float64 `json:"sleeperWidth"`     // 默认 260
	SleeperHeight    float64 `json:"sleeperHeight"`    // 默认 200
	SleeperSpacing   float64 `json:"sleeperSpacing"`   // 默认 600
	BallastTopWidth  float64 `json:"ballastTopWidth"`  // 道床顶宽(mm), 默认 3600; 0=不生成道床
	BallastThickness float64 `json:"ballastThickness"` // 默认 300
	BallastSlope     float64 `json:"ballastSlope"`     // 边坡 1:n, 默认 1.5
	NoSleepers       bool    `json:"noSleepers"`       // true=不生成轨枕, 默认生成
}

// YardTrackLayout 一条股道的布局数据
// Centerline 为未裁剪的原始中心线; TrimS/TrimE 为识别出的裁剪弧长
// (TrimS=起点裁掉的长度, TrimE=终点保留到的累计弧长, 默认=全长)
type YardTrackLayout struct {
	Centerline [][3]float64       `json:"centerline"` // 原始中心线点列 (mm)
	TrimS      float64            `json:"trimS"`      // 起点裁剪弧长 (mm)
	TrimE      float64            `json:"trimE"`      // 终点保留弧长 (自起点累计, mm; 0=全长)
	Props      TrackGeoProperties `json:"props"`      // 轨距/超高/钢轨/轨枕等生成参数
}

// YardTurnoutLayout 一组单开道岔的布局数据 (CreateTurnoutWithPlace 所需全部参数)
type YardTurnoutLayout struct {
	Position         [3]float64 `json:"position"`         // 岔心节点位置 (世界坐标)
	MainDir          [3]float64 `json:"mainDir"`          // 主行进方向 (水平单位向量)
	IsLeftHand       bool       `json:"isLeftHand"`       // 开向: true=左开
	TurnoutNo        int        `json:"turnoutNo"`        // 道岔号数 {9,12,18,30,42}
	SwitchRailLength float64    `json:"switchRailLength"` // 尖轨长度
	LeadCurveRadius  float64    `json:"leadCurveRadius"`  // 导曲线半径
	Gauge            float64    `json:"gauge"`            // 轨距
	RailHeight       float64    `json:"railHeight"`       // 钢轨断面尺寸
	RailHeadWidth    float64    `json:"railHeadWidth"`
	RailBaseWidth    float64    `json:"railBaseWidth"`
	EdgeIn           int        `json:"edgeIn"`  // 主入边索引 (Tracks 下标, 信息用)
	EdgeOut          int        `json:"edgeOut"` // 主出边索引
	EdgeDiv          int        `json:"edgeDiv"` // 侧股边索引
}

// YardCrossingLayout 一个菱形交叉的布局数据 (CreateFrogWithPlace 所需全部参数)
type YardCrossingLayout struct {
	Position      [3]float64 `json:"position"`   // 交点位置 (世界坐标)
	Direction     [3]float64 `json:"direction"`  // 辙叉角平分线方向 (水平单位向量)
	TurnoutNo     int        `json:"turnoutNo"`  // 辙叉号数
	Gauge         float64    `json:"gauge"`      // 轨距
	RailHeight    float64    `json:"railHeight"` // 钢轨断面尺寸
	RailHeadWidth float64    `json:"railHeadWidth"`
	RailBaseWidth float64    `json:"railBaseWidth"`
	EdgeA         int        `json:"edgeA"` // 相交边索引 (Tracks 下标, 信息用)
	EdgeB         int        `json:"edgeB"`
}

// Yard 生成一个站场所需的全部中间数据 (编辑器存库/前端编辑校正的数据契约)
type Yard struct {
	Base
	Tracks    []YardTrackLayout    `json:"tracks"`
	Turnouts  []YardTurnoutLayout  `json:"turnouts,omitempty"`
	Crossings []YardCrossingLayout `json:"crossings,omitempty"`
}

func NewYard() *Yard {
	return &Yard{
		Base: Base{Type: "RAILWAY/Yard"},
	}
}
