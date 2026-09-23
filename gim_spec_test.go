package topotypes

import (
	"encoding/json"
	"testing"

	"github.com/flywave/topotypes/gim/ec"
	"github.com/flywave/topotypes/gim/gs"
	"github.com/flywave/topotypes/gim/gt"
)

// 规范符合性 round-trip 测试：
// Q/GDW 11809—2018 附录 B 基本图元/型钢、Q/GDW 11810.2 附录 A、T/CEC 5056.3 附录 B。

func roundTripParametric(t *testing.T, shape interface{}) map[string]interface{} {
	t.Helper()
	js, err := json.Marshal(shape)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var out map[string]interface{}
	if err := json.Unmarshal(js, &out); err != nil {
		t.Fatalf("unmarshal to map: %v", err)
	}
	return out
}

func TestSpecSteelSectionsRoundTrip(t *testing.T) {
	models := []struct {
		ty    string
		shape interface{}
		model string
	}{
		{"GIM/GS/EquilateralAngleSteel", gs.NewEquilateralAngleSteel(), "L50X4"},
		{"GIM/GS/ScaleneAngleSteel", gs.NewScaleneAngleSteel(), "L75X50X5"},
		{"GIM/GS/I-Beam", gs.NewIBeamSteel(), "I200X100X7X11.4"},
		{"GIM/GS/ILightbeams", gs.NewILightBeamSteel(), "I200X100X5.2X8.4"},
		{"GIM/GS/H-beam", gs.NewHBeamSteel(), "H400X200X8X13"},
		{"GIM/GS/BeamChannel", gs.NewBeamChannelSteel(), "C200X75X9X11"},
		{"GIM/GS/LightBeamChannel", gs.NewLightBeamChannelSteel(), "C200X75X5X9"},
		{"GIM/GS/FlatSteel", gs.NewFlatSteel(), "-60X6"},
		{"GIM/GS/L-Steel", gs.NewLSteel(), "L60X40X5"},
		{"GIM/GS/T-Steel", gs.NewTSectionSteel(), "T100X100X6X8"},
		{"GIM/GS/RoundSteel", gs.NewRoundSteel(), "D20"},
		{"GIM/GS/RoundSteelTube", gs.NewRoundSteelTube(), "D200X8"},
		{"GIM/GS/RectangularSteelTube", gs.NewRectangularSteelTube(), "R200X100X8"},
		{"GIM/GS/SquareSteelTube", gs.NewSquareSteelTube(), "S200X8"},
		{"GIM/GS/DoubleChannelSteel", gs.NewDoubleChannelSteel(), "2C200X75X9X11"},
		{"GIM/GS/EquilateralDoubleAngleSteel", gs.NewEquilateralDoubleAngleSteel(), "2L50X4"},
		{"GIM/GS/UnequalAngleSteel", gs.NewUnequalAngleSteel(), "2L75X50X5"},
		{"GIM/GS/PolygonRoundSteelTube", gs.NewPolygonRoundSteelTube(), "P8X200X8"},
	}
	for _, m := range models {
		sb := m.shape.(interface {
			SetModel(string)
			SetLength(float64)
		})
		sb.SetModel(m.model)
		sb.SetLength(6000)
		out := roundTripParametric(t, m.shape)
		if out["type"] != m.ty {
			t.Errorf("%s type = %v", m.ty, out["type"])
		}
		if out["model"] != m.model {
			t.Errorf("%s model = %v", m.ty, out["model"])
		}
		if out["length"].(float64) != 6000 {
			t.Errorf("%s length = %v", m.ty, out["length"])
		}
	}
}

func TestSpecWireCableParamsRoundTrip(t *testing.T) {
	// 规范 Wire 参数：StartCoord/EndCoord/StartVector/EndVector/Sag/D/FitCoordArray
	wire := gs.NewWire()
	wire.StartPoint = [3]float64{0, 0, 0}
	wire.EndPoint = [3]float64{1000, 0, -50}
	wire.StartDir = &[3]float64{1, 0, 0}
	wire.EndDir = &[3]float64{1, 0, 0}
	wire.Sag = 30
	wire.Diameter = 20
	wire.FitPoints = [][3]float64{{500, 0, -35}}
	out := roundTripParametric(t, wire)
	for _, key := range []string{"startCoord", "endCoord", "startVector", "endVector", "sag", "diameter", "fitCoordArray"} {
		if _, ok := out[key]; !ok {
			t.Errorf("Wire 缺少规范参数 %s", key)
		}
	}
	js, _ := json.Marshal(out)
	back := &gs.Wire{}
	if err := json.Unmarshal(js, back); err != nil {
		t.Fatal(err)
	}
	if back.StartPoint != wire.StartPoint || back.Sag != 30 || len(back.FitPoints) != 1 {
		t.Errorf("Wire round-trip 不一致: %+v", back)
	}

	// 规范 Cable 参数：StartCoord/EndCoord/InflectionCoordArray/IRArray/D
	cable := gs.NewCable()
	cable.StartPoint = [3]float64{0, 0, 0}
	cable.EndPoint = [3]float64{2000, 0, 0}
	cable.InflectionPoints = [][3]float64{{1000, 100, 0}}
	cable.Radii = []float64{500}
	cable.Diameter = 80
	out = roundTripParametric(t, cable)
	for _, key := range []string{"startCoord", "endCoord", "inflectionCoordArray", "irArray", "diameter"} {
		if _, ok := out[key]; !ok {
			t.Errorf("Cable 缺少规范参数 %s", key)
		}
	}
	js, _ = json.Marshal(out)
	backC := &gs.Cable{}
	if err := json.Unmarshal(js, backC); err != nil {
		t.Fatal(err)
	}
	if backC.StartPoint != cable.StartPoint || len(backC.InflectionPoints) != 1 || backC.Radii[0] != 500 {
		t.Errorf("Cable round-trip 不一致: %+v", backC)
	}
}

func TestSpecPoleTowerSubLegRoundTrip(t *testing.T) {
	// 规范杆塔分段 body/leg/subleg，子腿高度为负
	tower := gt.NewPoleTower()
	tower.Heights = []*gt.PoleTowerHeight{{Value: 15000, BodyId: "body1", LegId: "leg1"}}
	tower.Bodies = []*gt.PoleTowerBody{{
		Id:     "body1",
		Height: 20000,
		Nodes:  []*gt.PoleTowerBodyNode{{Id: "1", Position: [3]float64{-600, 600, 38000}}},
		Legs: []*gt.PoleTowerBodyLeg{{
			Id:             "leg1",
			CommonHeight:   3000,
			SpecificHeight: 5000,
			Nodes: []*gt.PoleTowerBodyNode{
				{Id: "357", Position: [3]float64{-950, 1900, 21000}},
			},
			SubLegs: []*gt.PoleTowerSubLeg{{
				Id:     "subleg1",
				Height: -3000,
				Nodes: []*gt.PoleTowerBodyNode{
					{Id: "67", Position: [3]float64{-2140, 2140, 19000}},
				},
			}},
		}},
	}}
	js, err := json.Marshal(tower)
	if err != nil {
		t.Fatal(err)
	}
	back := &gt.PoleTower{}
	if err := json.Unmarshal(js, back); err != nil {
		t.Fatal(err)
	}
	leg := back.Bodies[0].Legs[0]
	if len(leg.SubLegs) != 1 || leg.SubLegs[0].Height != -3000 || len(leg.SubLegs[0].Nodes) != 1 {
		t.Errorf("长短腿 round-trip 不一致: %+v", leg.SubLegs)
	}
}

func TestSpecCrossingObjectRoundTrip(t *testing.T) {
	// 规范交叉跨越物：CODE=32（等级公路）/POINTNUM/POINTn=编号,X,Y,Z,节点代码
	obj := gt.NewCrossingObject()
	obj.Code = 81
	obj.Points = []gt.CrossingPoint{
		{Id: 1, Position: [3]float64{0, 0, 0}, NodeCode: 13},
		{Id: 2, Position: [3]float64{5000, 0, 0}, NodeCode: 13},
		{Id: 3, Position: [3]float64{5000, 3000, 0}, NodeCode: 13},
	}
	obj.Lines = [][]int{{1, 2, 3}}
	js, err := json.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}
	back := &gt.CrossingObject{}
	if err := json.Unmarshal(js, back); err != nil {
		t.Fatal(err)
	}
	if back.Code != 81 || len(back.Points) != 3 || back.Points[2].NodeCode != 13 || len(back.Lines) != 1 {
		t.Errorf("交叉跨越物 round-trip 不一致: %+v", back)
	}
}

func TestSpecAnchorBoltRoundTrip(t *testing.T) {
	// 规范基础地脚螺栓：Bolt / BoltNum=n / Bolt1=x,y,z,规格,材质
	bolt := gt.NewAnchorBolt()
	bolt.Bolts = []gt.AnchorBoltItem{
		{Position: [3]float64{100, 200, 300}, Specification: "M22", Material: "35#"},
		{Position: [3]float64{-100, 200, 300}, Specification: "M24", Material: "35#"},
	}
	js, err := json.Marshal(bolt)
	if err != nil {
		t.Fatal(err)
	}
	back := &gt.AnchorBolt{}
	if err := json.Unmarshal(js, back); err != nil {
		t.Fatal(err)
	}
	if len(back.Bolts) != 2 || back.Bolts[0].Specification != "M22" || back.Bolts[1].Material != "35#" {
		t.Errorf("地脚螺栓 round-trip 不一致: %+v", back.Bolts)
	}
}

func TestSpecStringWirePointsRoundTrip(t *testing.T) {
	// 规范表 4：绝缘子串接线点（分裂线夹单号一组、双号一组）
	wps := gt.NewStringWirePoints()
	wps.Points = []gt.StringWirePoint{
		{Id: 0, Position: [3]float64{0, 0, 0}},
		{Id: 1, Position: [3]float64{120, 0, 10}},
	}
	js, err := json.Marshal(wps)
	if err != nil {
		t.Fatal(err)
	}
	back := &gt.StringWirePoints{}
	if err := json.Unmarshal(js, back); err != nil {
		t.Fatal(err)
	}
	if back.Type != "GIM/GT/StringWirePoints" || len(back.Points) != 2 || back.Points[1].Position != [3]float64{120, 0, 10} {
		t.Errorf("接线点 round-trip 不一致: %+v", back)
	}
}

func TestSpecTemperatureFiberRoundTrip(t *testing.T) {
	// T/CEC 5056.3 §4.3.2 测温光纤参数化
	fiber := ec.NewTemperatureFiber()
	fiber.Points = []ec.Point{{0, 0, 0}, {1000, 0, 0}}
	fiber.OutsideDiameter = 8
	js, err := json.Marshal(fiber)
	if err != nil {
		t.Fatal(err)
	}
	back := &ec.TemperatureFiber{}
	if err := json.Unmarshal(js, back); err != nil {
		t.Fatal(err)
	}
	if back.Type != "GIM/EC/TemperatureFiber" || len(back.Points) != 2 || back.OutsideDiameter != 8 {
		t.Errorf("测温光纤 round-trip 不一致: %+v", back)
	}
}

func TestParseSteelSection(t *testing.T) {
	cases := []struct {
		model string
		kind  gs.SteelSectionKind
		a, b  float64
	}{
		{"L50X4", gs.SteelSectionAngle, 50, 4},
		{"L75x50x5", gs.SteelSectionAngle, 75, 50},
		{"I200X100X7X11.4", gs.SteelSectionIBeam, 200, 100},
		{"H400X200X8X13", gs.SteelSectionIBeam, 400, 200},
		{"C200X75X9X11", gs.SteelSectionChannel, 200, 75},
		{"[200X75X9X11", gs.SteelSectionChannel, 200, 75},
		{"-60X6", gs.SteelSectionFlat, 60, 6},
		{"F60X6", gs.SteelSectionFlat, 60, 6},
		{"T50X5", gs.SteelSectionT, 50, 5},
		{"D20", gs.SteelSectionRound, 20, 0},
		{"φ200X8", gs.SteelSectionRoundTube, 200, 8},
		{"D200X8", gs.SteelSectionRoundTube, 200, 8},
		{"R200X100X8", gs.SteelSectionRectTube, 200, 100},
		{"S200X8", gs.SteelSectionSquareTube, 200, 8},
		{"2C200X75X9X11", gs.SteelSectionDoubleChannel, 200, 75},
		{"2L50X4", gs.SteelSectionDoubleAngle, 50, 4},
		{"2L75X50X5", gs.SteelSectionDoubleAngle, 75, 50},
		{"P8X200X8", gs.SteelSectionPolygonTube, 8, 200},
	}
	for _, c := range cases {
		sec, err := gs.ParseSteelSection(c.model)
		if err != nil {
			t.Errorf("ParseSteelSection(%q): %v", c.model, err)
			continue
		}
		if sec.Kind != c.kind {
			t.Errorf("ParseSteelSection(%q) kind = %v, want %v", c.model, sec.Kind, c.kind)
		}
		if sec.Leg1 != c.a && sec.Diameter != c.a && sec.Sides != int(c.a) {
			t.Errorf("ParseSteelSection(%q) 主尺寸 = %v, want %v", c.model, sec.Leg1, c.a)
		}
	}
	// GB 牌号表: 纯牌号工字钢/槽钢 (GB/T 706-2008)
	gb := []struct {
		node, model  string
		h, b, tw, tf float64
		wantKind     gs.SteelSectionKind
	}{
		{"I-Beam", "I20a", 200, 100, 7.0, 11.4, gs.SteelSectionIBeam},
		{"I-Beam", "I20b", 200, 102, 9.0, 11.4, gs.SteelSectionIBeam},
		{"I-Beam", "I10", 100, 68, 4.5, 7.6, gs.SteelSectionIBeam},
		{"I-Beam", "I32c", 320, 134, 13.5, 15.0, gs.SteelSectionIBeam},
		{"ILightbeams", "I20", 200, 100, 5.2, 8.4, gs.SteelSectionIBeam},
		{"BeamChannel", "C10", 100, 48, 5.3, 8.5, gs.SteelSectionChannel},
		{"BeamChannel", "[14b", 140, 60, 8.0, 9.5, gs.SteelSectionChannel},
		{"BeamChannel", "10#", 100, 48, 5.3, 8.5, gs.SteelSectionChannel},
		{"BeamChannel", "C25c", 250, 82, 11.0, 12.0, gs.SteelSectionChannel},
		{"LightBeamChannel", "C10", 100, 46, 4.5, 7.6, gs.SteelSectionChannel},
	}
	for _, c := range gb {
		sec, err := gs.ParseSteelSectionForNode(c.node, c.model)
		if err != nil {
			t.Errorf("ParseSteelSectionForNode(%q, %q): %v", c.node, c.model, err)
			continue
		}
		if sec.Kind != c.wantKind || sec.Leg1 != c.h || sec.Leg2 != c.b || sec.Thickness != c.tw || sec.FlangeThickness != c.tf {
			t.Errorf("%s %q 解析不符: got %+v, want h=%v b=%v tw=%v tf=%v", c.node, c.model, sec, c.h, c.b, c.tw, c.tf)
		}
	}
	// 未收录牌号仍应报错
	if _, err := gs.ParseSteelSectionForNode("I-Beam", "I63c"); err == nil {
		t.Errorf("未收录牌号 I63c 应返回错误")
	}
	if _, err := gs.ParseSteelSection(""); err == nil {
		t.Errorf("空型号应返回错误")
	}
}
