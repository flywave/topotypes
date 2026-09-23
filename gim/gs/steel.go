package gs

import (
	"fmt"
	"strconv"
	"strings"
)

// 规范型钢构件（Q/GDW 11809—2018《输变电工程三维设计模型交互规范》附录 B）。
// 所有型钢共用驱动参数：Model（型号字符串，如 L50X4）与 Length（长度，mm）。
// 规范定义的型钢节点名共 18 种，全部以 Model+Length 描述截面。

// SteelBase 规范型钢公共基类
type SteelBase struct {
	GsBase
	Model  string  `json:"model"`
	Length float64 `json:"length"`
}

func (s *SteelBase) GetModel() string    { return s.Model }
func (s *SteelBase) GetLength() float64  { return s.Length }
func (s *SteelBase) SetModel(m string)   { s.Model = m }
func (s *SteelBase) SetLength(l float64) { s.Length = l }

// EquilateralAngleSteel 等边角钢
type EquilateralAngleSteel struct{ SteelBase }

func NewEquilateralAngleSteel() *EquilateralAngleSteel {
	return &EquilateralAngleSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/EquilateralAngleSteel"}}}
}

// ScaleneAngleSteel 不等边角钢
type ScaleneAngleSteel struct{ SteelBase }

func NewScaleneAngleSteel() *ScaleneAngleSteel {
	return &ScaleneAngleSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/ScaleneAngleSteel"}}}
}

// IBeamSteel 工字钢（规范节点名 I-beam）
type IBeamSteel struct{ SteelBase }

func NewIBeamSteel() *IBeamSteel {
	return &IBeamSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/I-Beam"}}}
}

// ILightBeamSteel 轻型工字钢（规范节点名 ILightbeams）
type ILightBeamSteel struct{ SteelBase }

func NewILightBeamSteel() *ILightBeamSteel {
	return &ILightBeamSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/ILightbeams"}}}
}

// HBeamSteel H型钢（规范节点名 H-beam）
type HBeamSteel struct{ SteelBase }

func NewHBeamSteel() *HBeamSteel {
	return &HBeamSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/H-beam"}}}
}

// BeamChannelSteel 槽钢（规范节点名 BeamChannel）
type BeamChannelSteel struct{ SteelBase }

func NewBeamChannelSteel() *BeamChannelSteel {
	return &BeamChannelSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/BeamChannel"}}}
}

// LightBeamChannelSteel 轻型槽钢（规范节点名 LightBeamChannel）
type LightBeamChannelSteel struct{ SteelBase }

func NewLightBeamChannelSteel() *LightBeamChannelSteel {
	return &LightBeamChannelSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/LightBeamChannel"}}}
}

// FlatSteel 扁钢
type FlatSteel struct{ SteelBase }

func NewFlatSteel() *FlatSteel {
	return &FlatSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/FlatSteel"}}}
}

// LSteel L钢（规范节点名 L-Steel）
type LSteel struct{ SteelBase }

func NewLSteel() *LSteel {
	return &LSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/L-Steel"}}}
}

// TSectionSteel T型钢（规范节点名 T-Steel）
type TSectionSteel struct{ SteelBase }

func NewTSectionSteel() *TSectionSteel {
	return &TSectionSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/T-Steel"}}}
}

// RoundSteel 圆钢
type RoundSteel struct{ SteelBase }

func NewRoundSteel() *RoundSteel {
	return &RoundSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/RoundSteel"}}}
}

// RoundSteelTube 圆钢管
type RoundSteelTube struct{ SteelBase }

func NewRoundSteelTube() *RoundSteelTube {
	return &RoundSteelTube{SteelBase{GsBase: GsBase{Type: "GIM/GS/RoundSteelTube"}}}
}

// RectangularSteelTube 矩形钢管
type RectangularSteelTube struct{ SteelBase }

func NewRectangularSteelTube() *RectangularSteelTube {
	return &RectangularSteelTube{SteelBase{GsBase: GsBase{Type: "GIM/GS/RectangularSteelTube"}}}
}

// SquareSteelTube 方形钢管
type SquareSteelTube struct{ SteelBase }

func NewSquareSteelTube() *SquareSteelTube {
	return &SquareSteelTube{SteelBase{GsBase: GsBase{Type: "GIM/GS/SquareSteelTube"}}}
}

// DoubleChannelSteel 双槽钢
type DoubleChannelSteel struct{ SteelBase }

func NewDoubleChannelSteel() *DoubleChannelSteel {
	return &DoubleChannelSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/DoubleChannelSteel"}}}
}

// EquilateralDoubleAngleSteel 等边双角钢
type EquilateralDoubleAngleSteel struct{ SteelBase }

func NewEquilateralDoubleAngleSteel() *EquilateralDoubleAngleSteel {
	return &EquilateralDoubleAngleSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/EquilateralDoubleAngleSteel"}}}
}

// UnequalAngleSteel 不等边双角钢（规范节点名 UnequalAngleSteel）
type UnequalAngleSteel struct{ SteelBase }

func NewUnequalAngleSteel() *UnequalAngleSteel {
	return &UnequalAngleSteel{SteelBase{GsBase: GsBase{Type: "GIM/GS/UnequalAngleSteel"}}}
}

// PolygonRoundSteelTube 多边形钢管
type PolygonRoundSteelTube struct{ SteelBase }

func NewPolygonRoundSteelTube() *PolygonRoundSteelTube {
	return &PolygonRoundSteelTube{SteelBase{GsBase: GsBase{Type: "GIM/GS/PolygonRoundSteelTube"}}}
}

// SteelSectionKind 型钢截面类别
type SteelSectionKind int

const (
	SteelSectionAngle         SteelSectionKind = iota // 角钢（等边/不等边）
	SteelSectionIBeam                                 // 工字钢/H型钢/轻型工字钢
	SteelSectionChannel                               // 槽钢/轻型槽钢
	SteelSectionFlat                                  // 扁钢
	SteelSectionT                                     // T型钢
	SteelSectionRound                                 // 圆钢
	SteelSectionRoundTube                             // 圆钢管
	SteelSectionRectTube                              // 矩形钢管
	SteelSectionSquareTube                            // 方形钢管
	SteelSectionDoubleChannel                         // 双槽钢
	SteelSectionDoubleAngle                           // 双角钢（等边/不等边）
	SteelSectionPolygonTube                           // 多边形钢管
)

// SteelSection 解析后的型钢截面尺寸（单位 mm）
type SteelSection struct {
	Kind SteelSectionKind
	// 角钢/双角钢：Leg1、Leg2 为肢宽，Thickness 为肢厚
	// 工字钢/槽钢/T型钢：Leg1=截面高，Leg2=翼缘宽，Thickness=腹板厚，FlangeThickness=翼缘厚
	// 扁钢：Leg1=宽，Thickness=厚
	// 圆钢/圆钢管：Diameter；钢管类 Thickness=壁厚
	// 方形钢管：Leg1=边长；矩形钢管：Leg1=长边，Leg2=短边
	// 多边形钢管：Sides=边数，Leg1=对边距
	Leg1            float64
	Leg2            float64
	Thickness       float64
	FlangeThickness float64
	Diameter        float64
	Sides           int
}

// SteelSpecNodeNames 规范附录 B 全部型钢节点名（JSON type 后缀）
var SteelSpecNodeNames = []string{
	"EquilateralAngleSteel", "ScaleneAngleSteel", "I-Beam", "ILightbeams", "H-beam",
	"BeamChannel", "LightBeamChannel", "FlatSteel", "L-Steel", "T-Steel",
	"RoundSteel", "RoundSteelTube", "RectangularSteelTube", "SquareSteelTube",
	"DoubleChannelSteel", "EquilateralDoubleAngleSteel", "UnequalAngleSteel",
	"PolygonRoundSteelTube",
}

func iBeamSection(kind SteelSectionKind, nums []float64) *SteelSection {
	return &SteelSection{
		Kind:            kind,
		Leg1:            nums[0],
		Leg2:            nums[1],
		Thickness:       nums[2],
		FlangeThickness: nums[3],
	}
}

func parseDimList(s string, counts ...int) ([]float64, error) {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == 'X' || r == '*' || r == ' '
	})
	if len(parts) == 0 {
		return nil, fmt.Errorf("无尺寸数据")
	}
	nums := make([]float64, 0, len(parts))
	for _, p := range parts {
		v, err := strconv.ParseFloat(strings.TrimSpace(p), 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, v)
	}
	for _, c := range counts {
		if len(nums) == c {
			return nums, nil
		}
	}
	return nil, fmt.Errorf("尺寸个数 %d 不匹配 %v", len(nums), counts)
}

// ================= GB 牌号截面表 =================
// 纯牌号型号（如 I20a、C10、10#）依据 GB/T 706-2008 热轧型钢截面尺寸表解析。
// 覆盖范围: 普通工字钢 10~40、轻型工字钢 10~30、槽钢 5~32、轻型槽钢 5~20 的常用牌号;
// 未收录牌号仍返回错误, 不做尺寸推断。

var gbIBeamSections = map[string][4]float64{
	"10": {100, 68, 4.5, 7.6}, "12": {120, 74, 5.0, 8.4}, "12.6": {126, 74, 5.0, 8.4},
	"14": {140, 80, 5.5, 9.1}, "16": {160, 88, 6.0, 9.9}, "18": {180, 94, 6.5, 10.7},
	"20a": {200, 100, 7.0, 11.4}, "20b": {200, 102, 9.0, 11.4},
	"22a": {220, 110, 7.5, 12.3}, "22b": {220, 112, 9.5, 12.3},
	"24a": {240, 116, 8.0, 13.0}, "24b": {240, 118, 10.0, 13.0},
	"25a": {250, 116, 8.0, 13.0}, "25b": {250, 118, 10.0, 13.0},
	"27a": {270, 122, 8.5, 13.7}, "27b": {270, 124, 10.5, 13.7},
	"28a": {280, 122, 8.5, 13.7}, "28b": {280, 124, 10.5, 13.7},
	"30a": {300, 126, 9.0, 14.4}, "30b": {300, 128, 11.0, 14.4}, "30c": {300, 130, 13.0, 14.4},
	"32a": {320, 130, 9.5, 15.0}, "32b": {320, 132, 11.5, 15.0}, "32c": {320, 134, 13.5, 15.0},
	"36a": {360, 136, 10.0, 15.8}, "36b": {360, 138, 12.0, 15.8}, "36c": {360, 140, 14.0, 15.8},
	"40a": {400, 142, 10.5, 16.5}, "40b": {400, 144, 12.5, 16.5}, "40c": {400, 146, 14.5, 16.5},
}

var gbLightIBeamSections = map[string][4]float64{
	"10": {100, 55, 4.5, 7.2}, "12": {120, 64, 4.8, 7.3}, "14": {140, 73, 4.9, 7.5},
	"16": {160, 81, 5.0, 7.8}, "18": {180, 90, 5.1, 8.1}, "20": {200, 100, 5.2, 8.4},
	"22": {220, 110, 5.4, 8.7}, "24": {240, 115, 5.6, 8.9}, "27": {270, 125, 6.0, 9.5},
	"30": {300, 135, 6.5, 10.0},
}

var gbChannelSections = map[string][4]float64{
	"5": {50, 37, 4.5, 7.0}, "6.3": {63, 40, 4.8, 7.5}, "8": {80, 43, 5.0, 8.0},
	"10": {100, 48, 5.3, 8.5}, "12.6": {126, 53, 5.5, 9.0},
	"14a": {140, 58, 6.0, 9.5}, "14b": {140, 60, 8.0, 9.5},
	"16a": {160, 63, 6.5, 10.0}, "16b": {160, 65, 8.5, 10.0},
	"18a": {180, 68, 7.0, 10.5}, "18b": {180, 70, 9.0, 10.5},
	"20a": {200, 73, 7.0, 11.0}, "20b": {200, 75, 9.0, 11.0},
	"22a": {220, 77, 7.0, 11.5}, "22b": {220, 79, 9.0, 11.5},
	"25a": {250, 78, 7.0, 12.0}, "25b": {250, 80, 9.0, 12.0}, "25c": {250, 82, 11.0, 12.0},
	"28a": {280, 82, 7.5, 12.5}, "28b": {280, 84, 9.5, 12.5}, "28c": {280, 86, 11.5, 12.5},
	"32a": {320, 88, 8.0, 14.0}, "32b": {320, 90, 10.0, 14.0}, "32c": {320, 92, 12.0, 14.0},
}

var gbLightChannelSections = map[string][4]float64{
	"5": {50, 32, 4.4, 7.0}, "6.5": {65, 36, 4.4, 7.2}, "8": {80, 40, 4.5, 7.4},
	"10": {100, 46, 4.5, 7.6}, "12": {120, 52, 4.8, 7.8}, "14": {140, 58, 4.9, 8.1},
	"16": {160, 64, 5.0, 8.4}, "18": {180, 70, 5.1, 8.7}, "20": {200, 76, 5.2, 9.0},
}

// ParseSteelSectionForNode 按规范节点名选择对应 GB 表: 轻型工字钢/轻型槽钢使用轻型截面表。
// node 为规范节点名后缀 (如 "I-Beam"/"ILightbeams"/"BeamChannel"/"LightBeamChannel")。
func ParseSteelSectionForNode(node, model string) (*SteelSection, error) {
	return parseSteelSection(model, node == "ILightbeams", node == "LightBeamChannel")
}

// ParseSteelSection 解析型号字符串 (普通工字钢/槽钢表)。
func ParseSteelSection(model string) (*SteelSection, error) {
	return parseSteelSection(model, false, false)
}

// gbLookupBeam 从牌号表解析 "高度[变体]" 形式的型号, 返回 截面高/翼缘宽/腹板厚/翼缘厚。
func gbLookupBeam(table map[string][4]float64, designation string) (*SteelSection, error) {
	dims, ok := table[strings.ToLower(designation)]
	if !ok {
		return nil, fmt.Errorf("牌号 %q 不在内置 GB 截面表覆盖范围内", designation)
	}
	return &SteelSection{Kind: SteelSectionIBeam, Leg1: dims[0], Leg2: dims[1], Thickness: dims[2], FlangeThickness: dims[3]}, nil
}

// parseGBDesignation 解析纯牌号 (无显式尺寸) 的工字钢/槽钢型号。
// 形式: "I20a"/"I20"/"20a" → 工字钢; "C10"/"[10"/"10#"/"10" → 槽钢。
func parseGBDesignation(s string, forceChannel, lightI, lightC bool) (*SteelSection, bool) {
	designation := s
	isChannel := forceChannel
	isI := false
	switch {
	case strings.HasPrefix(designation, "I"):
		isI, designation = true, designation[1:]
	case strings.HasPrefix(designation, "C"):
		isChannel, designation = true, designation[1:]
	case strings.HasPrefix(designation, "["):
		isChannel, designation = true, designation[1:]
	}
	designation = strings.TrimSuffix(designation, "#")
	designation = strings.ToLower(designation)
	if designation == "" {
		return nil, false
	}
	// 高度数字 + 可选 a/b/c 变体
	i := 0
	for i < len(designation) && (designation[i] >= '0' && designation[i] <= '9' || designation[i] == '.') {
		i++
	}
	if i == 0 {
		return nil, false
	}
	variant := designation[i:]
	if variant != "" && variant != "a" && variant != "b" && variant != "c" {
		return nil, false
	}
	key := designation
	if variant == "" && !strings.Contains(key, ".") {
		// 无变体且无小数点的高度, 同表内存在 a 变体时默认取 a (如 工字钢 20 → 20a)
		if hasA := tableKeyWithA(isChannel, key); hasA {
			key = key + "a"
		}
	}
	table := gbChannelSections
	kind := SteelSectionChannel
	switch {
	case isI:
		table, kind = gbIBeamSections, SteelSectionIBeam
	case isChannel:
		// 槽钢
	default:
		return nil, false
	}
	if lightI && isI {
		table, kind = gbLightIBeamSections, SteelSectionIBeam
	}
	if lightC && isChannel {
		table, kind = gbLightChannelSections, SteelSectionChannel
	}
	dims, ok := table[key]
	if !ok {
		// 无变体回退: 槽钢 5/6.3/8/10/12.6 等本身无变体, 直接用原始键
		if _, raw := table[designation]; raw {
			key = designation
			dims = table[designation]
		} else {
			return nil, false
		}
	}
	_ = kind
	kindFinal := SteelSectionChannel
	if isI {
		kindFinal = SteelSectionIBeam
	}
	return &SteelSection{Kind: kindFinal, Leg1: dims[0], Leg2: dims[1], Thickness: dims[2], FlangeThickness: dims[3]}, true
}

func tableKeyWithA(channel bool, key string) bool {
	table := gbIBeamSections
	if channel {
		table = gbChannelSections
	}
	_, ok := table[key+"a"]
	return ok
}

func parseSteelSection(model string, lightI, lightC bool) (*SteelSection, error) {
	s := strings.ToUpper(strings.TrimSpace(model))
	s = strings.ReplaceAll(s, "×", "X")
	s = strings.ReplaceAll(s, "Φ", "φ")
	if s == "" {
		return nil, fmt.Errorf("型钢型号为空")
	}

	// 多边形钢管 P{边数}X{对边距}X{壁厚}
	if strings.HasPrefix(s, "P") {
		nums, err := parseDimList(strings.TrimPrefix(s, "P"), 3)
		if err != nil {
			return nil, fmt.Errorf("多边形钢管型号 %q 应为 P边数X对边距X壁厚", model)
		}
		if nums[0] < 3 {
			return nil, fmt.Errorf("多边形钢管边数 %v 无效", nums[0])
		}
		return &SteelSection{Kind: SteelSectionPolygonTube, Sides: int(nums[0]), Leg1: nums[1], Thickness: nums[2]}, nil
	}

	// 双型钢 2[... / 2C... / 2L...
	double := false
	if strings.HasPrefix(s, "2") && len(s) > 1 {
		switch {
		case strings.HasPrefix(s[1:], "[") || strings.HasPrefix(s[1:], "C"):
			double = true
			s = "C" + strings.TrimPrefix(strings.TrimPrefix(s[1:], "["), "C")
		case strings.HasPrefix(s[1:], "L"):
			double = true
			s = s[1:]
		}
	}

	switch {
	case strings.HasPrefix(s, "L"): // 角钢 L50X4 / L75X50X5
		nums, err := parseDimList(s[1:], 2, 3)
		if err != nil {
			return nil, fmt.Errorf("角钢型号 %q 应为 L肢宽X肢厚 或 L长肢X短肢X肢厚", model)
		}
		sec := &SteelSection{Kind: SteelSectionAngle, Leg1: nums[0], Thickness: nums[len(nums)-1]}
		if len(nums) == 3 {
			sec.Leg2 = nums[1]
		} else {
			sec.Leg2 = nums[0]
		}
		if double {
			sec.Kind = SteelSectionDoubleAngle
		}
		return sec, nil
	case strings.HasPrefix(s, "I"): // 工字钢: 显式尺寸或 GB 牌号
		if nums, err := parseDimList(s[1:], 4); err == nil {
			return iBeamSection(SteelSectionIBeam, nums), nil
		}
		if sec, ok := parseGBDesignation(s, false, lightI, lightC); ok {
			return sec, nil
		}
		return nil, fmt.Errorf("工字钢型号 %q 应为 I高X翼缘宽X腹板厚X翼缘厚 或 GB 牌号 (如 I20a)", model)
	case strings.HasPrefix(s, "H"): // H型钢
		nums, err := parseDimList(s[1:], 4)
		if err != nil {
			return nil, fmt.Errorf("H型钢型号 %q 应为 H高X翼缘宽X腹板厚X翼缘厚", model)
		}
		return iBeamSection(SteelSectionIBeam, nums), nil
	case strings.HasPrefix(s, "C"), strings.HasPrefix(s, "["): // 槽钢
		if nums, err := parseDimList(strings.TrimLeft(s, "C["), 4); err == nil {
			sec := iBeamSection(SteelSectionChannel, nums)
			if double {
				sec.Kind = SteelSectionDoubleChannel
			}
			return sec, nil
		}
		if sec, ok := parseGBDesignation(s, true, lightI, lightC); ok {
			if double {
				sec.Kind = SteelSectionDoubleChannel
			}
			return sec, nil
		}
		return nil, fmt.Errorf("槽钢型号 %q 应为 C高X翼缘宽X腹板厚X翼缘厚 或 GB 牌号 (如 C10)", model)
	case strings.HasPrefix(s, "-"), strings.HasPrefix(s, "F"): // 扁钢
		nums, err := parseDimList(strings.TrimLeft(s, "-F"), 2)
		if err != nil {
			return nil, fmt.Errorf("扁钢型号 %q 应为 -宽X厚", model)
		}
		return &SteelSection{Kind: SteelSectionFlat, Leg1: nums[0], Thickness: nums[1]}, nil
	case strings.HasPrefix(s, "T"): // T型钢 T50X5 / T100X100X6X8
		nums, err := parseDimList(s[1:], 2, 4)
		if err != nil {
			return nil, fmt.Errorf("T型钢型号 %q 应为 T宽X厚 或 T高X翼缘宽X腹板厚X翼缘厚", model)
		}
		if len(nums) == 4 {
			return iBeamSection(SteelSectionT, nums), nil
		}
		return &SteelSection{Kind: SteelSectionT, Leg1: nums[0], Thickness: nums[1]}, nil
	case strings.HasPrefix(s, "R"): // 矩形钢管 R200X100X8
		nums, err := parseDimList(s[1:], 3)
		if err != nil {
			return nil, fmt.Errorf("矩形钢管型号 %q 应为 R长边X短边X壁厚", model)
		}
		return &SteelSection{Kind: SteelSectionRectTube, Leg1: nums[0], Leg2: nums[1], Thickness: nums[2]}, nil
	case strings.HasPrefix(s, "S"): // 方形钢管 S200X8
		nums, err := parseDimList(s[1:], 2)
		if err != nil {
			return nil, fmt.Errorf("方形钢管型号 %q 应为 S边长X壁厚", model)
		}
		return &SteelSection{Kind: SteelSectionSquareTube, Leg1: nums[0], Thickness: nums[1]}, nil
	case strings.HasPrefix(s, "□"): // 方形/矩形钢管
		nums, err := parseDimList(strings.TrimPrefix(s, "□"), 2, 3)
		if err != nil {
			return nil, fmt.Errorf("钢管型号 %q 应为 □边长X壁厚 或 □长边X短边X壁厚", model)
		}
		if len(nums) == 3 {
			return &SteelSection{Kind: SteelSectionRectTube, Leg1: nums[0], Leg2: nums[1], Thickness: nums[2]}, nil
		}
		return &SteelSection{Kind: SteelSectionSquareTube, Leg1: nums[0], Thickness: nums[1]}, nil
	case strings.HasPrefix(s, "φ"), strings.HasPrefix(s, "D"): // 圆钢/圆钢管
		body := strings.TrimPrefix(s, "φ")
		body = strings.TrimPrefix(body, "D")
		nums, err := parseDimList(body, 1, 2)
		if err != nil {
			return nil, fmt.Errorf("圆钢型号 %q 应为 φ直径 或 φ直径X壁厚", model)
		}
		if len(nums) == 2 {
			return &SteelSection{Kind: SteelSectionRoundTube, Diameter: nums[0], Thickness: nums[1]}, nil
		}
		return &SteelSection{Kind: SteelSectionRound, Diameter: nums[0]}, nil
	}
	// 无前缀纯牌号 (如 10#) 按槽钢牌号尝试
	if sec, ok := parseGBDesignation(s, true, lightI, lightC); ok {
		return sec, nil
	}
	return nil, fmt.Errorf("无法解析的型钢型号 %q", model)
}
