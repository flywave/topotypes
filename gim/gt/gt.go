package gt

import (
	"encoding/json"
	"fmt"
)

const Major = "GT"

type GtBase struct {
	Version int    `json:"version"`
	Type    string `json:"type"`
}

func (b *GtBase) GetType() string {
	return b.Type
}

// BoredPileBase represents a bored pile base
type BoredPileBase struct {
	GtBase
	H1 float64 `json:"H1"` // 上部圆柱高度
	H2 float64 `json:"H2"` // 过渡段高度
	H3 float64 `json:"H3"` // 底部圆柱高度
	H4 float64 `json:"H4"` // 桩头高度
	D1 float64 `json:"d"`  // 上部直径
	D  float64 `json:"D"`  // 底部直径
}

func NewBoredPileBase() *BoredPileBase {
	return &BoredPileBase{
		GtBase: GtBase{Type: "GIM/GT/BoredPileBase"},
	}
}

// PileCapBase represents a pile cap base
type PileCapBase struct {
	GtBase
	H1         float64      `json:"H1"`
	H2         float64      `json:"H2"`
	H3         float64      `json:"H3"`
	H4         float64      `json:"H4"`
	H5         float64      `json:"H5"`
	H6         float64      `json:"H6"`
	D1         float64      `json:"d"`
	D          float64      `json:"D"`
	B11        float64      `json:"b"`
	B1         float64      `json:"B1"`
	L1         float64      `json:"L1"`
	E1         float64      `json:"e1"`
	E2         float64      `json:"e2"`
	Cs         float64      `json:"cs"`
	ZCOUNT     int          `json:"ZCOUNT"`
	ZPOSTARRAY [][3]float64 `json:"ZPOSTARRAY"`
}

func NewPileCapBase() *PileCapBase {
	return &PileCapBase{
		GtBase: GtBase{Type: "GIM/GT/PileCapBase"},
	}
}

// RockAnchorBase represents a rock anchor base
type RockAnchorBase struct {
	GtBase
	H1         float64      `json:"H1"`
	H2         float64      `json:"H2"`
	D1         float64      `json:"d"`
	B1         float64      `json:"B1"`
	L1         float64      `json:"L1"`
	ZCOUNT     int          `json:"ZCOUNT"`
	ZPOSTARRAY [][3]float64 `json:"ZPOSTARRAY"`
}

func NewRockAnchorBase() *RockAnchorBase {
	return &RockAnchorBase{
		GtBase: GtBase{Type: "GIM/GT/RockAnchorBase"},
	}
}

// RockPileCapBase represents a rock pile cap base
type RockPileCapBase struct {
	GtBase
	H1         float64      `json:"H1"`
	H2         float64      `json:"H2"`
	H3         float64      `json:"H3"`
	D1         float64      `json:"d"`
	B11        float64      `json:"b"`
	B1         float64      `json:"B1"`
	L1         float64      `json:"L1"`
	E1         float64      `json:"e1"`
	E2         float64      `json:"e2"`
	Cs         float64      `json:"cs"`
	ZCOUNT     int          `json:"ZCOUNT"`
	ZPOSTARRAY [][3]float64 `json:"ZPOSTARRAY"`
}

func NewRockPileCapBase() *RockPileCapBase {
	return &RockPileCapBase{
		GtBase: GtBase{Type: "GIM/GT/RockPileCapBase"},
	}
}

// EmbeddedRockAnchorBase represents an embedded rock anchor base
type EmbeddedRockAnchorBase struct {
	GtBase
	H1 float64 `json:"H1"`
	H2 float64 `json:"H2"`
	H3 float64 `json:"H3"`
	D1 float64 `json:"d"`
	D  float64 `json:"D"`
}

func NewEmbeddedRockAnchorBase() *EmbeddedRockAnchorBase {
	return &EmbeddedRockAnchorBase{
		GtBase: GtBase{Type: "GIM/GT/EmbeddedRockAnchorBase"},
	}
}

// InclinedRockAnchorBase represents an inclined rock anchor base
type InclinedRockAnchorBase struct {
	GtBase
	H1     float64 `json:"H1"`
	H2     float64 `json:"H2"`
	D1     float64 `json:"d"`
	D      float64 `json:"D"`
	B      float64 `json:"B"`
	L      float64 `json:"L"`
	E1     float64 `json:"e1"`
	E2     float64 `json:"e2"`
	Alpha1 float64 `json:"alpha1"`
	Alpha2 float64 `json:"alpha2"`
}

func NewInclinedRockAnchorBase() *InclinedRockAnchorBase {
	return &InclinedRockAnchorBase{
		GtBase: GtBase{Type: "GIM/GT/InclinedRockAnchorBase"},
	}
}

// ExcavatedBase represents an excavated base
type ExcavatedBase struct {
	GtBase
	H1     float64 `json:"H1"`
	H2     float64 `json:"H2"`
	H3     float64 `json:"H3"`
	D1     float64 `json:"d"`
	D      float64 `json:"D"`
	Alpha1 float64 `json:"alpha1"`
	Alpha2 float64 `json:"alpha2"`
}

func NewExcavatedBase() *ExcavatedBase {
	return &ExcavatedBase{
		GtBase: GtBase{Type: "GIM/GT/ExcavatedBase"},
	}
}

// StepBase represents a step base
type StepBase struct {
	GtBase
	H   float64 `json:"H"`
	H1  float64 `json:"H1"`
	H2  float64 `json:"H2"`
	H3  float64 `json:"H3"`
	B11 float64 `json:"b"`
	B1  float64 `json:"B1"`
	B2  float64 `json:"B2"`
	B3  float64 `json:"B3"`
	L1  float64 `json:"L1"`
	L2  float64 `json:"L2"`
	L3  float64 `json:"L3"`
	N   int     `json:"N"`
}

func NewStepBase() *StepBase {
	return &StepBase{
		GtBase: GtBase{Type: "GIM/GT/StepBase"},
	}
}

// StepPlateBase represents a step plate base
type StepPlateBase struct {
	GtBase
	H      float64 `json:"H"`
	H1     float64 `json:"H1"`
	H2     float64 `json:"H2"`
	H3     float64 `json:"H3"`
	B11    float64 `json:"b"`
	B1     float64 `json:"B1"`
	B2     float64 `json:"B2"`
	L1     float64 `json:"L1"`
	L2     float64 `json:"L2"`
	Alpha1 float64 `json:"alpha1"`
	Alpha2 float64 `json:"alpha2"`
	N      int     `json:"N"`
}

func NewStepPlateBase() *StepPlateBase {
	return &StepPlateBase{
		GtBase: GtBase{Type: "GIM/GT/StepPlateBase"},
	}
}

// SlopedBaseBase represents a sloped base base
type SlopedBaseBase struct {
	GtBase
	H1     float64 `json:"H1"`
	H2     float64 `json:"H2"`
	H3     float64 `json:"H3"`
	B11    float64 `json:"b"`
	B1     float64 `json:"B1"`
	B2     float64 `json:"B2"`
	L1     float64 `json:"L1"`
	L2     float64 `json:"L2"`
	Alpha1 float64 `json:"alpha1"`
	Alpha2 float64 `json:"alpha2"`
}

func NewSlopedBaseBase() *SlopedBaseBase {
	return &SlopedBaseBase{
		GtBase: GtBase{Type: "GIM/GT/SlopedBaseBase"},
	}
}

// CompositeCaissonBase represents a composite caisson base
type CompositeCaissonBase struct {
	GtBase
	H1  float64 `json:"H1"`
	H2  float64 `json:"H2"`
	H3  float64 `json:"H3"`
	H4  float64 `json:"H4"`
	B11 float64 `json:"b"`
	D   float64 `json:"D"`
	T   float64 `json:"t"`
	B1  float64 `json:"B1"`
	B2  float64 `json:"B2"`
	L1  float64 `json:"L1"`
	L2  float64 `json:"L2"`
}

func NewCompositeCaissonBase() *CompositeCaissonBase {
	return &CompositeCaissonBase{
		GtBase: GtBase{Type: "GIM/GT/CompositeCaissonBase"},
	}
}

// RaftBase represents a raft base
type RaftBase struct {
	GtBase
	H1  float64 `json:"H1"`
	H2  float64 `json:"H2"`
	H3  float64 `json:"H3"`
	B11 float64 `json:"b1"`
	B22 float64 `json:"b2"`
	B1  float64 `json:"B1"`
	B2  float64 `json:"B2"`
	L1  float64 `json:"L1"`
	L2  float64 `json:"L2"`
}

func NewRaftBase() *RaftBase {
	return &RaftBase{
		GtBase: GtBase{Type: "GIM/GT/RaftBase"},
	}
}

// DirectBuriedBase represents a direct buried base
type DirectBuriedBase struct {
	GtBase
	H1  float64 `json:"H1"`
	H2  float64 `json:"H2"`
	D11 float64 `json:"d"`
	D   float64 `json:"D"`
	B   float64 `json:"B"`
	T   float64 `json:"t"`
}

func NewDirectBuriedBase() *DirectBuriedBase {
	return &DirectBuriedBase{
		GtBase: GtBase{Type: "GIM/GT/DirectBuriedBase"},
	}
}

// SteelSleeveBase represents a steel sleeve base
type SteelSleeveBase struct {
	GtBase
	H1  float64 `json:"H1"`
	H2  float64 `json:"H2"`
	H3  float64 `json:"H3"`
	H4  float64 `json:"H4"`
	D11 float64 `json:"d"`
	D1  float64 `json:"D1"`
	D2  float64 `json:"D2"`
	T   float64 `json:"t"`
	B1  float64 `json:"B1"`
	B2  float64 `json:"B2"`
}

func NewSteelSleeveBase() *SteelSleeveBase {
	return &SteelSleeveBase{
		GtBase: GtBase{Type: "GIM/GT/SteelSleeveBase"},
	}
}

// PrecastColumnBase represents a precast column base
type PrecastColumnBase struct {
	GtBase
	H1  float64 `json:"H1"`
	H2  float64 `json:"H2"`
	H3  float64 `json:"H3"`
	D11 float64 `json:"d"`
	B1  float64 `json:"B1"`
	B2  float64 `json:"B2"`
	L1  float64 `json:"L1"`
	L2  float64 `json:"L2"`
}

func NewPrecastColumnBase() *PrecastColumnBase {
	return &PrecastColumnBase{
		GtBase: GtBase{Type: "GIM/GT/PrecastColumnBase"},
	}
}

// PrecastPinnedBase represents a precast pinned base
type PrecastPinnedBase struct {
	GtBase
	H1  float64 `json:"H1"`
	H2  float64 `json:"H2"`
	H3  float64 `json:"H3"`
	D11 float64 `json:"d"`
	B1  float64 `json:"B1"`
	B2  float64 `json:"B2"`
	L1  float64 `json:"L1"`
	L2  float64 `json:"L2"`
	B   float64 `json:"B"`
	H   float64 `json:"H"`
	L   float64 `json:"L"`
}

func NewPrecastPinnedBase() *PrecastPinnedBase {
	return &PrecastPinnedBase{
		GtBase: GtBase{Type: "GIM/GT/PrecastPinnedBase"},
	}
}

// PrecastMetalSupportBase represents a precast metal support base
type PrecastMetalSupportBase struct {
	GtBase
	H1  float64   `json:"H1"`
	H2  float64   `json:"H2"`
	H3  float64   `json:"H3"`
	H4  float64   `json:"H4"`
	B11 float64   `json:"b1"`
	B22 float64   `json:"b2"`
	B1  float64   `json:"B1"`
	B2  float64   `json:"B2"`
	L1  float64   `json:"L1"`
	L2  float64   `json:"L2"`
	S1  float64   `json:"S1"`
	S2  float64   `json:"S2"`
	N1  int       `json:"n1"`
	N2  int       `json:"n2"`
	HX  []float64 `json:"HX"`
}

func NewPrecastMetalSupportBase() *PrecastMetalSupportBase {
	return &PrecastMetalSupportBase{
		GtBase: GtBase{Type: "GIM/GT/PrecastMetalSupportBase"},
	}
}

// PrecastConcreteSupportBase represents a precast concrete support base
type PrecastConcreteSupportBase struct {
	GtBase
	H1  float64 `json:"H1"`
	H2  float64 `json:"H2"`
	H3  float64 `json:"H3"`
	H4  float64 `json:"H4"`
	H5  float64 `json:"H5"`
	B11 float64 `json:"b1"`
	B22 float64 `json:"b2"`
	B3  float64 `json:"b3"`
	B1  float64 `json:"B1"`
	B2  float64 `json:"B2"`
	L1  float64 `json:"L1"`
	L2  float64 `json:"L2"`
	S1  float64 `json:"S1"`
	N1  int     `json:"n1"`
}

func NewPrecastConcreteSupportBase() *PrecastConcreteSupportBase {
	return &PrecastConcreteSupportBase{
		GtBase: GtBase{Type: "GIM/GT/PrecastConcreteSupportBase"},
	}
}

// TransmissionLine represents a transmission line
type TransmissionLine struct {
	GtBase
	SectionalArea           float64 `json:"sectionalArea"`
	OutsideDiameter         float64 `json:"outsideDiameter"`
	WireWeight              float64 `json:"wireWeight"`
	CoefficientOfElasticity float64 `json:"coefficientOfElasticity"`
	ExpansionCoefficient    float64 `json:"expansionCoefficient"`
	RatedStrength           float64 `json:"ratedStrength"`
}

func NewTransmissionLine() *TransmissionLine {
	return &TransmissionLine{
		GtBase: GtBase{Type: "GIM/GT/TransmissionLine"},
	}
}

type FittingLength struct {
	LeftUpper  float64 `json:"leftUpper"`  // 左上金具长度
	RightUpper float64 `json:"rightUpper"` // 右上金具长度
	LeftLower  float64 `json:"leftLower"`  // 左下金具长度
	RightLower float64 `json:"rightLower"` // 右下金具长度
}

type MultiLink struct {
	Count       int     `json:"count"`       // 多联数量
	Spacing     float64 `json:"spacing"`     // 多联间距
	Arrangement string  `json:"arrangement"` // 排列方式
}

type InsulatorMember struct {
	Radius     float64 `json:"radius"`     // 绝缘子半径
	Height     float64 `json:"height"`     // 绝缘子高度
	LeftCount  int     `json:"leftCount"`  // 左侧片数
	RightCount int     `json:"rightCount"` // 右侧片数
	Material   string  `json:"material"`   // 材料类型
}

type GradingRing struct {
	Count    int     `json:"count"`    // 均压环数量
	Position float64 `json:"position"` // 均压环位置
	Height   float64 `json:"height"`   // 均压环高度
	Radius   float64 `json:"radius"`   // 均压环半径
}

// Insulator represents an insulator
type Insulator struct {
	GtBase
	SubNum         int              `json:"subNum"`        // 子串数量
	SubType        int              `json:"subType"`       // 子串类型
	SplitDistance  float64          `json:"splitDistance"` // 分裂间距
	VAngleLeft     float64          `json:"vAngleLeft"`    // 左侧V型角度
	VAngleRight    float64          `json:"vAngleRight"`   // 右侧V型角度
	ULinkLength    float64          `json:"uLinkLength"`   // U型环长度
	Weight         float64          `json:"weight"`        // 重量
	FittingLengths *FittingLength   `json:"fittingLengths"`
	MultiLink      *MultiLink       `json:"multiLink"`
	Insulator      *InsulatorMember `json:"insulator"`
	GradingRing    *GradingRing     `json:"gradingRing"`
	Application    string           `json:"application"` // 应用类型
	StringType     string           `json:"stringType"`  // 串型类型
}

func NewInsulator() *Insulator {
	return &Insulator{
		GtBase: GtBase{Type: "GIM/GT/Insulator"},
	}
}

type Member struct {
	Id            string     `json:"id"`            // 构件ID
	StartNodeId   string     `json:"startNodeId"`   // 起始节点ID
	EndNodeId     string     `json:"endNodeId"`     // 结束节点ID
	Type          string     `json:"type"`          // 构件类型
	Specification string     `json:"specification"` // 规格
	Material      string     `json:"material"`      // 材料
	XDirection    [3]float64 `json:"xDirection"`    // X方向
	YDirection    [3]float64 `json:"yDirection"`    // Y方向
	End1Diameter  float64    `json:"end1Diameter"`  // 端部1直径
	End2Diameter  float64    `json:"end2Diameter"`  // 端部2直径
	Thickness     float64    `json:"thickness"`     // 厚度
	Sides         int        `json:"sides"`         // 边数
}

type PoleTowerBodyNode struct {
	Id       string     `json:"id"`       // 节点ID
	Position [3]float64 `json:"position"` // 节点位置
}

type PoleTowerBodyLeg struct {
	Id             string               `json:"id"`             // 腿柱ID
	CommonHeight   float64              `json:"commonHeight"`   // 通用高度
	SpecificHeight float64              `json:"specificHeight"` // 特定高度
	Nodes          []*PoleTowerBodyNode `json:"nodes"`
}

type PoleTowerBody struct {
	Id     string               `json:"id"`     // 塔身ID
	Height float64              `json:"height"` // 塔身高度
	Nodes  []*PoleTowerBodyNode `json:"nodes"`
	Legs   []*PoleTowerBodyLeg  `json:"legs"`
}

type PoleTowerHeight struct {
	Value  float64 `json:"value"`  // 高度值
	BodyId string  `json:"bodyId"` // 所属塔身ID
	LegId  string  `json:"legId"`  // 所属腿柱ID
}

type Attachment []struct {
	Name     string     `json:"name"` // 附件名称
	Type     string     `json:"type"` // 附件类型
	Position [3]float64 `json:"position"`
}

// PoleTower represents a pole tower
type PoleTower struct {
	GtBase
	Heights     []*PoleTowerHeight `json:"heights"`
	Bodies      []*PoleTowerBody   `json:"bodies"`
	Members     []*Member          `json:"members"`
	Attachments []*Attachment      `json:"attachments"`
}

func NewPoleTower() *PoleTower {
	return &PoleTower{
		GtBase: GtBase{Type: "GIM/GT/PoleTower"},
	}
}

// SingleHookAnchor represents a single hook anchor
type SingleHookAnchor struct {
	GtBase
	BoltDiameter        float64 `json:"boltDiameter"`        // 螺栓直径
	ExposedLength       float64 `json:"exposedLength"`       // 外露长度
	NutCount            int     `json:"nutCount"`            // 螺母数量
	NutHeight           float64 `json:"nutHeight"`           // 螺母高度
	NutOD               float64 `json:"nutOD"`               // 螺母外径
	WasherCount         int     `json:"washerCount"`         // 垫圈数量
	WasherShape         int     `json:"washerShape"`         // 垫圈形状 (1-圆形, 2-方形)
	WasherSize          float64 `json:"washerSize"`          // 垫圈尺寸
	WasherThickness     float64 `json:"washerThickness"`     // 垫圈厚度
	AnchorLength        float64 `json:"anchorLength"`        // 锚固长度
	HookStraightLengthA float64 `json:"hookStraightLengthA"` // 钩直段长度A
	HookStraightLengthB float64 `json:"hookStraightLengthB"` // 钩直段长度B
	HookDiameter        float64 `json:"hookDiameter"`        // 钩直径
	AnchorBarDiameter   float64 `json:"anchorBarDiameter"`   // 锚筋直径
}

func NewSingleHookAnchor() *SingleHookAnchor {
	return &SingleHookAnchor{
		GtBase: GtBase{Type: "GIM/GT/SingleHookAnchor"},
	}
}

// TripleHookAnchor represents a triple hook anchor
type TripleHookAnchor struct {
	GtBase
	BoltDiameter       float64 `json:"boltDiameter"`       // 螺栓直径
	ExposedLength      float64 `json:"exposedLength"`      // 外露长度
	NutCount           int     `json:"nutCount"`           // 螺母数量
	NutHeight          float64 `json:"nutHeight"`          // 螺母高度
	NutOD              float64 `json:"nutOD"`              // 螺母外径
	WasherCount        int     `json:"washerCount"`        // 垫圈数量
	WasherShape        int     `json:"washerShape"`        // 垫圈形状 (1-圆形, 2-方形)
	WasherSize         float64 `json:"washerSize"`         // 垫圈尺寸
	WasherThickness    float64 `json:"washerThickness"`    // 垫圈厚度
	AnchorLength       float64 `json:"anchorLength"`       // 锚固长度
	HookStraightLength float64 `json:"hookStraightLength"` // 钩直段长度
	HookDiameter       float64 `json:"hookDiameter"`       // 钩直径
}

func NewTripleHookAnchor() *TripleHookAnchor {
	return &TripleHookAnchor{
		GtBase: GtBase{Type: "GIM/GT/TripleHookAnchor"},
	}
}

// RibbedAnchor represents a ribbed anchor
type RibbedAnchor struct {
	GtBase
	BoltDiameter       float64 `json:"boltDiameter"`       // 螺栓直径
	ExposedLength      float64 `json:"exposedLength"`      // 外露长度
	NutCount           int     `json:"nutCount"`           // 螺母数量
	NutHeight          float64 `json:"nutHeight"`          // 螺母高度
	NutOD              float64 `json:"nutOD"`              // 螺母外径
	WasherCount        int     `json:"washerCount"`        // 垫圈数量
	WasherShape        int     `json:"washerShape"`        // 垫圈形状 (1-圆形, 2-方形)
	WasherSize         float64 `json:"washerSize"`         // 垫圈尺寸
	WasherThickness    float64 `json:"washerThickness"`    // 垫圈厚度
	AnchorLength       float64 `json:"anchorLength"`       // 锚固长度
	BasePlateSize      float64 `json:"basePlateSize"`      // 底板尺寸
	RibTopWidth        float64 `json:"ribTopWidth"`        // 肋顶部宽度
	RibBottomWidth     float64 `json:"ribBottomWidth"`     // 肋底部宽度
	BasePlateThickness float64 `json:"basePlateThickness"` // 底板厚度
	RibHeight          float64 `json:"ribHeight"`          // 肋高度
	RibThickness       float64 `json:"ribThickness"`       // 肋厚度
}

func NewRibbedAnchor() *RibbedAnchor {
	return &RibbedAnchor{
		GtBase: GtBase{Type: "GIM/GT/RibbedAnchor"},
	}
}

// NutAnchor represents a nut anchor
type NutAnchor struct {
	GtBase
	BoltDiameter        float64 `json:"boltDiameter"`        // 螺栓直径
	ExposedLength       float64 `json:"exposedLength"`       // 外露长度
	NutCount            int     `json:"nutCount"`            // 螺母数量
	NutHeight           float64 `json:"nutHeight"`           // 螺母高度
	NutOD               float64 `json:"nutOD"`               // 螺母外径
	WasherCount         int     `json:"washerCount"`         // 垫圈数量
	WasherShape         int     `json:"washerShape"`         // 垫圈形状 (1-圆形, 2-方形)
	WasherSize          float64 `json:"washerSize"`          // 垫圈尺寸
	WasherThickness     float64 `json:"washerThickness"`     // 垫圈厚度
	AnchorLength        float64 `json:"anchorLength"`        // 锚固长度
	BasePlateSize       float64 `json:"basePlateSize"`       // 底板尺寸
	BasePlateThickness  float64 `json:"basePlateThickness"`  // 底板厚度
	BoltToPlateDistance float64 `json:"boltToPlateDistance"` // 螺栓到底板距离
}

func NewNutAnchor() *NutAnchor {
	return &NutAnchor{
		GtBase: GtBase{Type: "GIM/GT/NutAnchor"},
	}
}

// TripleArmAnchor represents a triple arm anchor
type TripleArmAnchor struct {
	GtBase
	BoltDiameter      float64 `json:"boltDiameter"`      // 螺栓直径
	ExposedLength     float64 `json:"exposedLength"`     // 外露长度
	NutCount          int     `json:"nutCount"`          // 螺母数量
	NutHeight         float64 `json:"nutHeight"`         // 螺母高度
	NutOD             float64 `json:"nutOD"`             // 螺母外径
	WasherCount       int     `json:"washerCount"`       // 垫圈数量
	WasherShape       int     `json:"washerShape"`       // 垫圈形状 (1-圆形, 2-方形)
	WasherSize        float64 `json:"washerSize"`        // 垫圈尺寸
	WasherThickness   float64 `json:"washerThickness"`   // 垫圈厚度
	AnchorLength      float64 `json:"anchorLength"`      // 锚固长度
	ArmDiameter       float64 `json:"armDiameter"`       // 臂直径
	ArmStraightLength float64 `json:"armStraightLength"` // 臂直段长度
	ArmBendLength     float64 `json:"armBendLength"`     // 臂弯曲段长度
	ArmBendAngle      float64 `json:"armBendAngle"`      // 臂弯曲角度(弧度)
}

func NewTripleArmAnchor() *TripleArmAnchor {
	return &TripleArmAnchor{
		GtBase: GtBase{Type: "GIM/GT/TripleArmAnchor"},
	}
}

// PositioningPlateAnchor represents a positioning plate anchor
type PositioningPlateAnchor struct {
	GtBase
	BoltDiameter      float64 `json:"boltDiameter"`      // 螺栓直径
	ExposedLength     float64 `json:"exposedLength"`     // 外露长度
	NutCount          int     `json:"nutCount"`          // 螺母数量
	NutHeight         float64 `json:"nutHeight"`         // 螺母高度
	NutOD             float64 `json:"nutOD"`             // 螺母外径
	WasherCount       int     `json:"washerCount"`       // 垫圈数量
	WasherShape       int     `json:"washerShape"`       // 垫圈形状 (1-圆形, 2-方形)
	WasherSize        float64 `json:"washerSize"`        // 垫圈尺寸
	WasherThickness   float64 `json:"washerThickness"`   // 垫圈厚度
	AnchorLength      float64 `json:"anchorLength"`      // 锚固长度
	PlateLength       float64 `json:"plateLength"`       // 定位板长度
	PlateThickness    float64 `json:"plateThickness"`    // 定位板厚度
	ToBaseDistance    float64 `json:"toBaseDistance"`    // 到基础距离
	ToBottomDistance  float64 `json:"toBottomDistance"`  // 到底部距离
	GroutHoleDiameter float64 `json:"groutHoleDiameter"` // 灌浆孔直径
}

func NewPositioningPlateAnchor() *PositioningPlateAnchor {
	return &PositioningPlateAnchor{
		GtBase: GtBase{Type: "GIM/GT/PositioningPlateAnchor"},
	}
}

// StubAngle represents a stub angle
type StubAngle struct {
	GtBase
	LegWidth      float64 `json:"legWidth"`      // 肢宽
	Thickness     float64 `json:"thickness"`     // 厚度
	Slope         float64 `json:"slope"`         // 坡度
	ExposedLength float64 `json:"exposedLength"` // 外露长度
	AnchorLength  float64 `json:"anchorLength"`  // 锚固长度
}

func NewStubAngle() *StubAngle {
	return &StubAngle{
		GtBase: GtBase{Type: "GIM/GT/StubAngle"},
	}
}

// StubTube represents a stub tube
type StubTube struct {
	GtBase
	Diameter      float64 `json:"diameter"`      // 管径
	Thickness     float64 `json:"thickness"`     // 壁厚
	Slope         float64 `json:"slope"`         // 坡度
	ExposedLength float64 `json:"exposedLength"` // 外露长度
	AnchorLength  float64 `json:"anchorLength"`  // 锚固长度
}

func NewStubTube() *StubTube {
	return &StubTube{
		GtBase: GtBase{Type: "GIM/GT/StubTube"},
	}
}

type Shape interface {
	GetType() string
}

func Unmarshal(ty string, bt []byte) (Shape, error) {
	switch ty {
	case "GIM/GT/BoredPileBase":
		rel := &BoredPileBase{}
		err := json.Unmarshal(bt, rel)
		return rel, err
	case "GIM/GT/PileCapBase":
		rel := &PileCapBase{}
		err := json.Unmarshal(bt, rel)
		return rel, err
	case "GIM/GT/SlopedBaseBase":
		base := &SlopedBaseBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/RockAnchorBase":
		base := &RockAnchorBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/InclinedRockAnchorBase":
		base := &InclinedRockAnchorBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/EmbeddedRockAnchorBase":
		base := &EmbeddedRockAnchorBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/ExcavatedBase":
		base := &ExcavatedBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/StepBase":
		base := &StepBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/StepPlateBase":
		base := &StepPlateBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/RockPileCapBase":
		base := &RockPileCapBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/CompositeCaissonBase":
		base := &CompositeCaissonBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/RaftBase":
		base := &RaftBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/DirectBuriedBase":
		base := &DirectBuriedBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/SteelSleeveBase":
		base := &SteelSleeveBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/PrecastColumnBase":
		base := &PrecastColumnBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/PrecastPinnedBase":
		base := &PrecastPinnedBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/PrecastMetalSupportBase":
		base := &PrecastMetalSupportBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/PrecastConcreteSupportBase":
		base := &PrecastConcreteSupportBase{}
		err := json.Unmarshal(bt, base)
		return base, err
	case "GIM/GT/TransmissionLine":
		line := &TransmissionLine{}
		err := json.Unmarshal(bt, line)
		return line, err
	case "GIM/GT/Insulator":
		insulator := &Insulator{}
		err := json.Unmarshal(bt, insulator)
		return insulator, err
	case "GIM/GT/PoleTower":
		tower := &PoleTower{}
		err := json.Unmarshal(bt, tower)
		return tower, err
	case "GIM/GT/SingleHookAnchor":
		anchor := &SingleHookAnchor{}
		err := json.Unmarshal(bt, anchor)
		return anchor, err
	case "GIM/GT/TripleHookAnchor":
		anchor := &TripleHookAnchor{}
		err := json.Unmarshal(bt, anchor)
		return anchor, err
	case "GIM/GT/RibbedAnchor":
		anchor := &RibbedAnchor{}
		err := json.Unmarshal(bt, anchor)
		return anchor, err
	case "GIM/GT/NutAnchor":
		anchor := &NutAnchor{}
		err := json.Unmarshal(bt, anchor)
		return anchor, err
	case "GIM/GT/TripleArmAnchor":
		anchor := &TripleArmAnchor{}
		err := json.Unmarshal(bt, anchor)
		return anchor, err
	case "GIM/GT/PositioningPlateAnchor":
		anchor := &PositioningPlateAnchor{}
		err := json.Unmarshal(bt, anchor)
		return anchor, err
	case "GIM/GT/StubAngle":
		angle := &StubAngle{}
		err := json.Unmarshal(bt, angle)
		return angle, err
	case "GIM/GT/StubTube":
		tube := &StubTube{}
		err := json.Unmarshal(bt, tube)
		return tube, err
	default:
		return nil, fmt.Errorf("invalid type: %s", ty)
	}
}
