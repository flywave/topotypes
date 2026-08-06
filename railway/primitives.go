package railway

// 镜像 go-topo primitives_railway.go 的 params struct (字段/单位/含义一致, mm 单位制)
// 几何字段约定: 点/线以 [3]float64 / [][3]float64 独立字段保留在数据契约中
// (go-topo Go 层对应字段已 Deprecated, 由 WithPlace 函数参数替代; topotypes 作为
// 存库/前端编辑的数据契约层保留几何字段)

// CantileverBase represents a cantilever base (腕臂底座)
type CantileverBase struct {
	Base
	Length       float64 `json:"length"`
	Width        float64 `json:"width"`
	Height       float64 `json:"height"`
	BoltSpacing  float64 `json:"boltSpacing"`
	BoltDiameter float64 `json:"boltDiameter"`
	BoltCount    int     `json:"boltCount"`
}

func NewCantileverBase() *CantileverBase {
	return &CantileverBase{
		Base: Base{Type: "RAILWAY/CantileverBase"},
	}
}

// MWSaddle represents a messenger wire saddle (承力索座)
type MWSaddle struct {
	Base
	Length       float64 `json:"length"`
	Width        float64 `json:"width"`
	Height       float64 `json:"height"`
	GrooveRadius float64 `json:"grooveRadius"`
	BoltDiameter float64 `json:"boltDiameter"`
}

func NewMWSaddle() *MWSaddle {
	return &MWSaddle{
		Base: Base{Type: "RAILWAY/MWSaddle"},
	}
}

// BalanceWeight represents a balance weight (坠砣)
type BalanceWeight struct {
	Base
	Width              float64 `json:"width"`
	Thickness          float64 `json:"thickness"`
	Height             float64 `json:"height"`
	CenterHoleDiameter float64 `json:"centerHoleDiameter"`
}

func NewBalanceWeight() *BalanceWeight {
	return &BalanceWeight{
		Base: Base{Type: "RAILWAY/BalanceWeight"},
	}
}

// WeightRod represents a weight rod (坠砣杆)
type WeightRod struct {
	Base
	RodDiameter     float64 `json:"rodDiameter"`
	RodLength       float64 `json:"rodLength"`
	TopHoleDiameter float64 `json:"topHoleDiameter"`
}

func NewWeightRod() *WeightRod {
	return &WeightRod{
		Base: Base{Type: "RAILWAY/WeightRod"},
	}
}

// AnchorFitting represents an anchor fitting (下锚金具)
type AnchorFitting struct {
	Base
	FittingType int     `json:"fittingType"` // anchor_fitting_type: 1=杵环杆, 2=双耳连接器, 3=楔形线夹
	Length      float64 `json:"length"`
	Diameter    float64 `json:"diameter"`
}

func NewAnchorFitting() *AnchorFitting {
	return &AnchorFitting{
		Base: Base{Type: "RAILWAY/AnchorFitting"},
	}
}

// Crossing represents an OCS crossing (线岔)
type Crossing struct {
	Base
	LimitPipeLength float64 `json:"limitPipeLength"`
	PipeDiameter    float64 `json:"pipeDiameter"`
	WireDiameter    float64 `json:"wireDiameter"`
	HeightDiff      float64 `json:"heightDiff"`
}

func NewCrossing() *Crossing {
	return &Crossing{
		Base: Base{Type: "RAILWAY/Crossing"},
	}
}

// HeadSpan represents a head span (软横跨)
type HeadSpan struct {
	Base
	Span                  float64 `json:"span"`
	HangPointCount        int     `json:"hangPointCount"`
	HangPointSpacing      float64 `json:"hangPointSpacing"`
	CrossCatenaryDiameter float64 `json:"crossCatenaryDiameter"`
	CrossCatenarySag      float64 `json:"crossCatenarySag"`
	UpperRopeDiameter     float64 `json:"upperRopeDiameter"`
	LowerRopeDiameter     float64 `json:"lowerRopeDiameter"`
	InsulatorLength       float64 `json:"insulatorLength"`
}

func NewHeadSpan() *HeadSpan {
	return &HeadSpan{
		Base: Base{Type: "RAILWAY/HeadSpan"},
	}
}

// TransverseSpan represents a transverse span (硬横跨)
type TransverseSpan struct {
	Base
	Span          float64 `json:"span"`
	BeamType      int     `json:"beamType"` // beam_section_type: 1=箱型, 2=H型, 3=桁架式, 4=组合式
	BeamHeight    float64 `json:"beamHeight"`
	BeamWidth     float64 `json:"beamWidth"`
	BeamThickness float64 `json:"beamThickness"`
	MastHeight    float64 `json:"mastHeight"`
	MastWidth     float64 `json:"mastWidth"`
}

func NewTransverseSpan() *TransverseSpan {
	return &TransverseSpan{
		Base: Base{Type: "RAILWAY/TransverseSpan"},
	}
}

// HangerPost represents a hanger post (硬横跨吊柱)
type HangerPost struct {
	Base
	SectionType       int     `json:"sectionType"` // hanger_post_section_type: 1=圆管, 2=方管, 3=H型钢
	Length            float64 `json:"length"`
	SectionSize       float64 `json:"sectionSize"`
	WallThickness     float64 `json:"wallThickness"`
	TopFlangeSize     float64 `json:"topFlangeSize"`
	TopFlangeThick    float64 `json:"topFlangeThick"`
	BottomFlangeSize  float64 `json:"bottomFlangeSize"`
	BottomFlangeThick float64 `json:"bottomFlangeThick"`
	BoltDiameter      float64 `json:"boltDiameter"`
	BoltSpacing       float64 `json:"boltSpacing"`
}

func NewHangerPost() *HangerPost {
	return &HangerPost{
		Base: Base{Type: "RAILWAY/HangerPost"},
	}
}

// PortalFrame represents a portal frame (梁顶门型架)
type PortalFrame struct {
	Base
	FrameHeight      float64 `json:"frameHeight"`
	FrameWidth       float64 `json:"frameWidth"`
	PostDiameter     float64 `json:"postDiameter"`
	PostWallThick    float64 `json:"postWallThick"`
	BeamDiameter     float64 `json:"beamDiameter"`
	BeamWallThick    float64 `json:"beamWallThick"`
	BeamLength       float64 `json:"beamLength"`
	BasePlateLength  float64 `json:"basePlateLength"`
	BasePlateWidth   float64 `json:"basePlateWidth"`
	BasePlateThick   float64 `json:"basePlateThick"`
	HangPointCount   int     `json:"hangPointCount"`
	HangPointSpacing float64 `json:"hangPointSpacing"`
	BoltSpacing      float64 `json:"boltSpacing"`
	BoltDiameter     float64 `json:"boltDiameter"`
}

func NewPortalFrame() *PortalFrame {
	return &PortalFrame{
		Base: Base{Type: "RAILWAY/PortalFrame"},
	}
}

// SuspensionHardSpan represents a suspension hard span (悬索式硬横跨)
type SuspensionHardSpan struct {
	Base
	Span                 float64 `json:"span"`
	MastHeight           float64 `json:"mastHeight"`
	MastWidth            float64 `json:"mastWidth"`
	CableDiameter        float64 `json:"cableDiameter"`
	CableSag             float64 `json:"cableSag"`
	DropperCableDiameter float64 `json:"dropperCableDiameter"`
	DropperCount         int     `json:"dropperCount"`
	DropperSpacing       float64 `json:"dropperSpacing"`
	InsulatorLength      float64 `json:"insulatorLength"`
	InsulatorDiameter    float64 `json:"insulatorDiameter"`
}

func NewSuspensionHardSpan() *SuspensionHardSpan {
	return &SuspensionHardSpan{
		Base: Base{Type: "RAILWAY/SuspensionHardSpan"},
	}
}

// PositioningCable represents a positioning cable (定位索)
// TopPoint/BottomPoint 为几何端点 (世界坐标, mm)
type PositioningCable struct {
	Base
	Diameter    float64    `json:"diameter"`
	TopPoint    [3]float64 `json:"topPoint"`
	BottomPoint [3]float64 `json:"bottomPoint"`
	Adjustable  bool       `json:"adjustable"`
}

func NewPositioningCable() *PositioningCable {
	return &PositioningCable{
		Base: Base{Type: "RAILWAY/PositioningCable"},
	}
}

// AuxBracket represents an auxiliary wire bracket (附加导线安装支架)
type AuxBracket struct {
	Base
	BracketType    int     `json:"bracketType"` // aux_bracket_type: 1=横担式, 2=壁挂式, 3=双支柱式
	MountHeight    float64 `json:"mountHeight"`
	OverhangLength float64 `json:"overhangLength"`
	BracketLength  float64 `json:"bracketLength"`
	BracketWidth   float64 `json:"bracketWidth"`
	BoltSpacing    float64 `json:"boltSpacing"`
	BoltDiameter   float64 `json:"boltDiameter"`
}

func NewAuxBracket() *AuxBracket {
	return &AuxBracket{
		Base: Base{Type: "RAILWAY/AuxBracket"},
	}
}

// Rail represents a rail (钢轨)
type Rail struct {
	Base
	RailHeight     float64 `json:"railHeight"`
	HeadWidth      float64 `json:"headWidth"`
	BaseWidth      float64 `json:"baseWidth"`
	WebThickness   float64 `json:"webThickness"`
	HeadHeight     float64 `json:"headHeight"`
	BaseHeight     float64 `json:"baseHeight"`
	HeadRadius     float64 `json:"headRadius"`
	StandardLength float64 `json:"standardLength"`
}

func NewRail() *Rail {
	return &Rail{
		Base: Base{Type: "RAILWAY/Rail"},
	}
}

// Sleeper represents a sleeper (轨枕)
type Sleeper struct {
	Base
	ShapeType     int     `json:"shapeType"` // sleeper_shape_type: 1=矩形, 2=梯形收腰
	Length        float64 `json:"length"`
	Width         float64 `json:"width"`
	Height        float64 `json:"height"`
	Gauge         float64 `json:"gauge"`
	RailBaseWidth float64 `json:"railBaseWidth"`
	GrooveDepth   float64 `json:"grooveDepth"`
	Spacing       float64 `json:"spacing"`
}

func NewSleeper() *Sleeper {
	return &Sleeper{
		Base: Base{Type: "RAILWAY/Sleeper"},
	}
}

// Ballast represents a ballast bed (道床), 沿中心线路径扫掠梯形断面
type Ballast struct {
	Base
	Centerline [][3]float64 `json:"centerline"` // 中心线点列 (mm)
	TopWidth   float64      `json:"topWidth"`
	Thickness  float64      `json:"thickness"`
	SideSlope  float64      `json:"sideSlope"` // 边坡 1:n
}

func NewBallast() *Ballast {
	return &Ballast{
		Base: Base{Type: "RAILWAY/Ballast"},
	}
}

// TrackSlab represents a track slab (轨道板)
type TrackSlab struct {
	Base
	Length                 float64 `json:"length"`
	Width                  float64 `json:"width"`
	Thickness              float64 `json:"thickness"`
	RailSeatCount          int     `json:"railSeatCount"`
	RailSeatSpacing        float64 `json:"railSeatSpacing"`
	CementAsphaltThickness float64 `json:"cementAsphaltThickness"`
}

func NewTrackSlab() *TrackSlab {
	return &TrackSlab{
		Base: Base{Type: "RAILWAY/TrackSlab"},
	}
}

// Fastener represents a rail fastener (扣件)
type Fastener struct {
	Base
	Spacing      float64 `json:"spacing"`
	Gauge        float64 `json:"gauge"`
	PadThickness float64 `json:"padThickness"`
	PadLength    float64 `json:"padLength"`
	PadWidth     float64 `json:"padWidth"`
}

func NewFastener() *Fastener {
	return &Fastener{
		Base: Base{Type: "RAILWAY/Fastener"},
	}
}

// GuardRail represents a guard rail (护轨)
type GuardRail struct {
	Base
	Height        float64 `json:"height"`
	HeadWidth     float64 `json:"headWidth"`
	BaseWidth     float64 `json:"baseWidth"`
	GrooveWidth   float64 `json:"grooveWidth"`
	TotalLength   float64 `json:"totalLength"`
	GaugeDistance float64 `json:"gaugeDistance"`
}

func NewGuardRail() *GuardRail {
	return &GuardRail{
		Base: Base{Type: "RAILWAY/GuardRail"},
	}
}

// MastAssembly represents a mast assembly (支柱装配)
type MastAssembly struct {
	Base
	MastType        int     `json:"mastType"`        // 1-格构式钢柱, 2-混凝土柱
	MastHeight      float64 `json:"mastHeight"`      // 柱全高(mm)
	CantileverType  int     `json:"cantileverType"`  // 0-无, 1-单臂, 2-双臂
	HasCrossArm     bool    `json:"hasCrossArm"`     // 双臂时是否有横担
	ArmDiameter     float64 `json:"armDiameter"`     // 腕臂管外径(mm)
	Stagger         float64 `json:"stagger"`         // 接触线拉出值(mm)
	CompType        int     `json:"compType"`        // 0-无, 1-棘轮, 2-滑轮
	RatedTension    float64 `json:"ratedTension"`    // 设计补偿张力(kN)
	HasGuyWire      bool    `json:"hasGuyWire"`      // 是否设下锚拉线
	ContactHeight   float64 `json:"contactHeight"`   // 导高(mm), 默认 5300
	StructureHeight float64 `json:"structureHeight"` // 结构高度(mm), 默认 1400
	SideOffset      float64 `json:"sideOffset"`      // 侧面限界 CX(mm), 默认 2900
}

func NewMastAssembly() *MastAssembly {
	return &MastAssembly{
		Base: Base{Type: "RAILWAY/MastAssembly"},
	}
}

// WeightStackParams 坠砣串参数 (纯数据, 内嵌于补偿装置或经 WeightStack 独立注册)
type WeightStackParams struct {
	BlockCount    int     `json:"blockCount"`    // 坠砣块数, 默认 8
	BlockDiameter float64 `json:"blockDiameter"` // 坠砣直径(mm), 默认 380
	BlockHeight   float64 `json:"blockHeight"`   // 单块厚度(mm), 默认 75
	BlockGap      float64 `json:"blockGap"`      // 块间间隙(mm), 默认 2
	RodDiameter   float64 `json:"rodDiameter"`   // 坠砣杆直径(mm), 默认 20
	RodLength     float64 `json:"rodLength"`     // 坠砣杆长度(mm), 默认 1200
	HoleDiameter  float64 `json:"holeDiameter"`  // 中心孔径(mm), 默认 30
}

// WeightStack represents a weight stack (坠砣串)
type WeightStack struct {
	Base
	WeightStackParams
}

func NewWeightStack() *WeightStack {
	return &WeightStack{
		Base: Base{Type: "RAILWAY/WeightStack"},
	}
}

// RatchetCompensator represents a ratchet compensator (棘轮补偿装置)
type RatchetCompensator struct {
	Base
	WheelDiameter float64           `json:"wheelDiameter"` // 默认 400
	WheelWidth    float64           `json:"wheelWidth"`    // 默认 60
	RopeDiameter  float64           `json:"ropeDiameter"`  // 默认 9
	StrokeLength  float64           `json:"strokeLength"`  // 默认 1200
	Stack         WeightStackParams `json:"stack"`
}

func NewRatchetCompensator() *RatchetCompensator {
	return &RatchetCompensator{
		Base: Base{Type: "RAILWAY/RatchetCompensator"},
	}
}

// AuxiliaryWire represents an auxiliary wire (附加导线本体, 带弛度扫掠)
type AuxiliaryWire struct {
	Base
	Diameter     float64 `json:"diameter"`
	Sag          float64 `json:"sag"`
	RatedTension float64 `json:"ratedTension"`
}

func NewAuxiliaryWire() *AuxiliaryWire {
	return &AuxiliaryWire{
		Base: Base{Type: "RAILWAY/AuxiliaryWire"},
	}
}

// Disconnector represents a disconnector (隔离开关)
type Disconnector struct {
	Base
	BaseLength      float64 `json:"baseLength"`      // 默认 900
	BaseWidth       float64 `json:"baseWidth"`       // 默认 220
	InsulatorHeight float64 `json:"insulatorHeight"` // 默认 600
	BladeLength     float64 `json:"bladeLength"`     // 默认 800
	OpenAngle       float64 `json:"openAngle"`       // 分闸角度°, 默认 75
}

func NewDisconnector() *Disconnector {
	return &Disconnector{
		Base: Base{Type: "RAILWAY/Disconnector"},
	}
}

// Arrester represents an arrester (避雷器)
type Arrester struct {
	Base
	Height        float64 `json:"height"`        // 默认 800
	OuterDiameter float64 `json:"outerDiameter"` // 默认 120
	ShedDiameter  float64 `json:"shedDiameter"`  // 默认 160
	ShedSpacing   float64 `json:"shedSpacing"`   // 默认 60
	ShedCount     int     `json:"shedCount"`     // 默认 8
}

func NewArrester() *Arrester {
	return &Arrester{
		Base: Base{Type: "RAILWAY/Arrester"},
	}
}

// PulleyCompensator represents a pulley compensator (滑轮补偿装置)
type PulleyCompensator struct {
	Base
	PulleyDiameter float64           `json:"pulleyDiameter"` // 默认 250
	GrooveWidth    float64           `json:"grooveWidth"`    // 默认 14
	PulleyCount    int               `json:"pulleyCount"`    // 默认 2
	RopeDiameter   float64           `json:"ropeDiameter"`   // 默认 9
	StrokeLength   float64           `json:"strokeLength"`   // 默认 1000
	Stack          WeightStackParams `json:"stack"`
	HasLimitFrame  bool              `json:"hasLimitFrame"`
}

func NewPulleyCompensator() *PulleyCompensator {
	return &PulleyCompensator{
		Base: Base{Type: "RAILWAY/PulleyCompensator"},
	}
}

// SleeveConnector represents a sleeve connector (双套筒连接器)
type SleeveConnector struct {
	Base
	TubeDiameter  float64 `json:"tubeDiameter"`  // 默认 60
	SleeveLength  float64 `json:"sleeveLength"`  // 默认 120
	WallThickness float64 `json:"wallThickness"` // 默认 5
	Angle         float64 `json:"angle"`         // 两套筒夹角°, 默认 45
	BoltDiameter  float64 `json:"boltDiameter"`  // 默认 12
}

func NewSleeveConnector() *SleeveConnector {
	return &SleeveConnector{
		Base: Base{Type: "RAILWAY/SleeveConnector"},
	}
}

// SleeveEar represents a sleeve ear (套管单耳)
type SleeveEar struct {
	Base
	TubeDiameter  float64 `json:"tubeDiameter"`  // 默认 60
	SleeveLength  float64 `json:"sleeveLength"`  // 默认 100
	WallThickness float64 `json:"wallThickness"` // 默认 5
	EarHeight     float64 `json:"earHeight"`     // 默认 60
	EarThickness  float64 `json:"earThickness"`  // 默认 8
	HoleDiameter  float64 `json:"holeDiameter"`  // 默认 16
}

func NewSleeveEar() *SleeveEar {
	return &SleeveEar{
		Base: Base{Type: "RAILWAY/SleeveEar"},
	}
}

// SwitchRail represents a switch rail (尖轨)
type SwitchRail struct {
	Base
	Length        float64 `json:"length"`
	RailHeight    float64 `json:"railHeight"`
	RailHeadWidth float64 `json:"railHeadWidth"`
	RailBaseWidth float64 `json:"railBaseWidth"`
	TipWidth      float64 `json:"tipWidth"`
	CurveRadius   float64 `json:"curveRadius"`
	IsLeftHand    bool    `json:"isLeftHand"`
}

func NewSwitchRail() *SwitchRail {
	return &SwitchRail{
		Base: Base{Type: "RAILWAY/SwitchRail"},
	}
}

// Frog represents a frog (辙叉)
type Frog struct {
	Base
	TurnoutNo     int     `json:"turnoutNo"`
	Gauge         float64 `json:"gauge"`
	RailHeight    float64 `json:"railHeight"`
	RailHeadWidth float64 `json:"railHeadWidth"`
	RailBaseWidth float64 `json:"railBaseWidth"`
}

func NewFrog() *Frog {
	return &Frog{
		Base: Base{Type: "RAILWAY/Frog"},
	}
}

// Turnout represents a turnout (单开道岔)
type Turnout struct {
	Base
	TurnoutNo        int     `json:"turnoutNo"`
	IsLeftHand       bool    `json:"isLeftHand"`
	Gauge            float64 `json:"gauge"`
	RailHeight       float64 `json:"railHeight"`
	RailHeadWidth    float64 `json:"railHeadWidth"`
	RailBaseWidth    float64 `json:"railBaseWidth"`
	SwitchRailLength float64 `json:"switchRailLength"`
	LeadCurveRadius  float64 `json:"leadCurveRadius"`
	FrogLength       float64 `json:"frogLength"` // reserved, go-topo 暂未生效
	SleeperCount     int     `json:"sleeperCount"`
	SleeperSpacing   float64 `json:"sleeperSpacing"`
}

func NewTurnout() *Turnout {
	return &Turnout{
		Base: Base{Type: "RAILWAY/Turnout"},
	}
}

// StraightTrack represents a straight track segment (直线轨道段)
// StartPoint/EndPoint 为几何端点 (世界坐标, mm)
type StraightTrack struct {
	Base
	StartPoint       [3]float64 `json:"startPoint"`
	EndPoint         [3]float64 `json:"endPoint"`
	Gauge            float64    `json:"gauge"`
	RailHeight       float64    `json:"railHeight"`
	RailHeadWidth    float64    `json:"railHeadWidth"`
	RailBaseWidth    float64    `json:"railBaseWidth"`
	SleeperLength    float64    `json:"sleeperLength"`
	SleeperWidth     float64    `json:"sleeperWidth"`
	SleeperHeight    float64    `json:"sleeperHeight"`
	SleeperSpacing   float64    `json:"sleeperSpacing"`
	BallastTopWidth  float64    `json:"ballastTopWidth"`
	BallastThickness float64    `json:"ballastThickness"`
	BallastSlope     float64    `json:"ballastSlope"`
}

func NewStraightTrack() *StraightTrack {
	return &StraightTrack{
		Base: Base{Type: "RAILWAY/StraightTrack"},
	}
}

// CurveTrack represents a curved track segment (曲线轨道段)
// CurveCenter 为几何圆心 (世界坐标, mm)
type CurveTrack struct {
	Base
	CurveCenter      [3]float64 `json:"curveCenter"`
	StartAngle       float64    `json:"startAngle"`
	SweepAngle       float64    `json:"sweepAngle"`
	CurveRadius      float64    `json:"curveRadius"`
	Gauge            float64    `json:"gauge"`
	SuperElevation   float64    `json:"superElevation"`
	RailHeight       float64    `json:"railHeight"`
	RailHeadWidth    float64    `json:"railHeadWidth"`
	RailBaseWidth    float64    `json:"railBaseWidth"`
	SleeperLength    float64    `json:"sleeperLength"`
	SleeperWidth     float64    `json:"sleeperWidth"`
	SleeperHeight    float64    `json:"sleeperHeight"`
	SleeperSpacing   float64    `json:"sleeperSpacing"`
	BallastTopWidth  float64    `json:"ballastTopWidth"`
	BallastThickness float64    `json:"ballastThickness"`
	BallastSlope     float64    `json:"ballastSlope"`
}

func NewCurveTrack() *CurveTrack {
	return &CurveTrack{
		Base: Base{Type: "RAILWAY/CurveTrack"},
	}
}

// RailPair represents a rail pair (轨排对), 沿中心线偏移 ±gauge/2 生成两条钢轨
type RailPair struct {
	Base
	Centerline     [][3]float64 `json:"centerline"` // 中心线点列 (mm)
	Gauge          float64      `json:"gauge"`
	SuperElevation float64      `json:"superElevation"`
	RailHeight     float64      `json:"railHeight"`
	RailHeadWidth  float64      `json:"railHeadWidth"`
	RailBaseWidth  float64      `json:"railBaseWidth"`
}

func NewRailPair() *RailPair {
	return &RailPair{
		Base: Base{Type: "RAILWAY/RailPair"},
	}
}

// SleeperLayout represents a sleeper layout (轨枕阵列), 沿中心线均匀布置
type SleeperLayout struct {
	Base
	Centerline [][3]float64 `json:"centerline"` // 中心线点列 (mm)
	Length     float64      `json:"length"`
	Width      float64      `json:"width"`
	Height     float64      `json:"height"`
	Spacing    float64      `json:"spacing"`
	Gauge      float64      `json:"gauge"`
}

func NewSleeperLayout() *SleeperLayout {
	return &SleeperLayout{
		Base: Base{Type: "RAILWAY/SleeperLayout"},
	}
}

// RetarderPoint represents a retarder point (减速顶点)
type RetarderPoint struct {
	Base
	Side             int     `json:"side"`       // 1-LEFT, 2-RIGHT, 3-BOTH
	DeviceType       int     `json:"deviceType"` // 1-液压, 2-摩擦, 3-可控
	MountType        int     `json:"mountType"`  // 1-轨内侧, 2-轨外侧, 3-双轨双侧
	Height           float64 `json:"height"`
	BodyDiameter     float64 `json:"bodyDiameter"`
	CapDiameter      float64 `json:"capDiameter"`
	CapHeight        float64 `json:"capHeight"`
	TransitionHeight float64 `json:"transitionHeight"`
	ArmLength        float64 `json:"armLength"`
	ArmWidth         float64 `json:"armWidth"`
	ArmThickness     float64 `json:"armThickness"`
	BoltDiameter     float64 `json:"boltDiameter"`
	PortDiameter     float64 `json:"portDiameter"`
}

func NewRetarderPoint() *RetarderPoint {
	return &RetarderPoint{
		Base: Base{Type: "RAILWAY/RetarderPoint"},
	}
}
