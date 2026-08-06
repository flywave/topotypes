package railway

import (
	"encoding/json"
	"reflect"
	"testing"
)

// allFactories 覆盖全部注册类型, 每个工厂必须填好 Base.Type 判别串
func allFactories() map[string]func() Shape {
	return map[string]func() Shape{
		"RAILWAY/RodInsulator":       func() Shape { return NewRodInsulator() },
		"RAILWAY/CrossArm":           func() Shape { return NewCrossArm() },
		"RAILWAY/LevelCantilever":    func() Shape { return NewLevelCantilever() },
		"RAILWAY/SlantCantilever":    func() Shape { return NewSlantCantilever() },
		"RAILWAY/CantileverBrace":    func() Shape { return NewCantileverBrace() },
		"RAILWAY/CurvedArm":          func() Shape { return NewCurvedArm() },
		"RAILWAY/RegArmBracket":      func() Shape { return NewRegArmBracket() },
		"RAILWAY/ContactWire":        func() Shape { return NewContactWire() },
		"RAILWAY/MessengerWire":      func() Shape { return NewMessengerWire() },
		"RAILWAY/GuyWire":            func() Shape { return NewGuyWire() },
		"RAILWAY/Dropper":            func() Shape { return NewDropper() },
		"RAILWAY/MastBracket":        func() Shape { return NewMastBracket() },
		"RAILWAY/SteelMast":          func() Shape { return NewSteelMast() },
		"RAILWAY/ConcreteMast":       func() Shape { return NewConcreteMast() },
		"RAILWAY/OcsFoundation":      func() Shape { return NewOcsFoundation() },
		"RAILWAY/RegistrationArm":    func() Shape { return NewRegistrationArm() },
		"RAILWAY/CantileverBase":     func() Shape { return NewCantileverBase() },
		"RAILWAY/MWSaddle":           func() Shape { return NewMWSaddle() },
		"RAILWAY/BalanceWeight":      func() Shape { return NewBalanceWeight() },
		"RAILWAY/WeightRod":          func() Shape { return NewWeightRod() },
		"RAILWAY/AnchorFitting":      func() Shape { return NewAnchorFitting() },
		"RAILWAY/Crossing":           func() Shape { return NewCrossing() },
		"RAILWAY/HeadSpan":           func() Shape { return NewHeadSpan() },
		"RAILWAY/TransverseSpan":     func() Shape { return NewTransverseSpan() },
		"RAILWAY/HangerPost":         func() Shape { return NewHangerPost() },
		"RAILWAY/PortalFrame":        func() Shape { return NewPortalFrame() },
		"RAILWAY/SuspensionHardSpan": func() Shape { return NewSuspensionHardSpan() },
		"RAILWAY/PositioningCable":   func() Shape { return NewPositioningCable() },
		"RAILWAY/AuxBracket":         func() Shape { return NewAuxBracket() },
		"RAILWAY/Rail":               func() Shape { return NewRail() },
		"RAILWAY/Sleeper":            func() Shape { return NewSleeper() },
		"RAILWAY/Ballast":            func() Shape { return NewBallast() },
		"RAILWAY/TrackSlab":          func() Shape { return NewTrackSlab() },
		"RAILWAY/Fastener":           func() Shape { return NewFastener() },
		"RAILWAY/GuardRail":          func() Shape { return NewGuardRail() },
		"RAILWAY/MastAssembly":       func() Shape { return NewMastAssembly() },
		"RAILWAY/WeightStack":        func() Shape { return NewWeightStack() },
		"RAILWAY/RatchetCompensator": func() Shape { return NewRatchetCompensator() },
		"RAILWAY/AuxiliaryWire":      func() Shape { return NewAuxiliaryWire() },
		"RAILWAY/Disconnector":       func() Shape { return NewDisconnector() },
		"RAILWAY/Arrester":           func() Shape { return NewArrester() },
		"RAILWAY/PulleyCompensator":  func() Shape { return NewPulleyCompensator() },
		"RAILWAY/SleeveConnector":    func() Shape { return NewSleeveConnector() },
		"RAILWAY/SleeveEar":          func() Shape { return NewSleeveEar() },
		"RAILWAY/SwitchRail":         func() Shape { return NewSwitchRail() },
		"RAILWAY/Frog":               func() Shape { return NewFrog() },
		"RAILWAY/Turnout":            func() Shape { return NewTurnout() },
		"RAILWAY/StraightTrack":      func() Shape { return NewStraightTrack() },
		"RAILWAY/CurveTrack":         func() Shape { return NewCurveTrack() },
		"RAILWAY/RailPair":           func() Shape { return NewRailPair() },
		"RAILWAY/SleeperLayout":      func() Shape { return NewSleeperLayout() },
		"RAILWAY/RetarderPoint":      func() Shape { return NewRetarderPoint() },
		"RAILWAY/AnchorSection":      func() Shape { return NewAnchorSection() },
		"RAILWAY/Yard":               func() Shape { return NewYard() },
	}
}

// 每个类型: 工厂 → marshal → "type" 判别串保留 → Unmarshal 分发 → 再 marshal → JSON 一致
func TestRoundTrip(t *testing.T) {
	for ty, factory := range allFactories() {
		t.Run(ty, func(t *testing.T) {
			shp := factory()
			if shp.GetType() != ty {
				t.Fatalf("factory Base.Type = %q, want %q", shp.GetType(), ty)
			}
			raw, err := json.Marshal(shp)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var probe map[string]interface{}
			if err := json.Unmarshal(raw, &probe); err != nil {
				t.Fatalf("unmarshal probe: %v", err)
			}
			got, ok := probe["type"].(string)
			if !ok || got != ty {
				t.Fatalf(`marshaled "type" = %v, want string %q`, probe["type"], ty)
			}
			back, err := Unmarshal(ty, raw)
			if err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}
			if back.GetType() != ty {
				t.Fatalf("round-trip GetType = %q, want %q", back.GetType(), ty)
			}
			raw2, err := json.Marshal(back)
			if err != nil {
				t.Fatalf("re-marshal: %v", err)
			}
			var a, b map[string]interface{}
			if err := json.Unmarshal(raw, &a); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(raw2, &b); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(a, b) {
				t.Fatalf("round-trip mismatch:\nfirst:  %s\nsecond: %s", raw, raw2)
			}
		})
	}
}

// type 遮蔽回归: 枚举字段不再占用 "type" 键, Marshal 后 "type" 必须是判别字符串
func TestTypeShadowingRegression(t *testing.T) {
	cases := []struct {
		name     string
		shape    Shape
		enumKey  string
		enumWant float64
	}{
		{"RAILWAY/RodInsulator", &RodInsulator{Base: Base{Type: "RAILWAY/RodInsulator"}, InsulatorType: RodInsulatorHollow}, "insulatorType", 2},
		{"RAILWAY/SteelMast", &SteelMast{Base: Base{Type: "RAILWAY/SteelMast"}, MastType: SteelMastHBeam}, "mastType", 2},
		{"RAILWAY/OcsFoundation", &OcsFoundation{Base: Base{Type: "RAILWAY/OcsFoundation"}, FoundationType: FoundationFlange}, "foundationType", 2},
		{"RAILWAY/RegistrationArm", &RegistrationArm{Base: Base{Type: "RAILWAY/RegistrationArm"}, ArmType: RegistrationArmCurved}, "armType", 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			raw, err := json.Marshal(c.shape)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var probe map[string]interface{}
			if err := json.Unmarshal(raw, &probe); err != nil {
				t.Fatalf("unmarshal probe: %v", err)
			}
			got, ok := probe["type"].(string)
			if !ok {
				t.Fatalf(`"type" is not a string (shadowed by enum field?): %v`, probe["type"])
			}
			if got != c.name {
				t.Fatalf(`"type" = %q, want %q`, got, c.name)
			}
			if probe[c.enumKey] != c.enumWant {
				t.Fatalf("%q = %v, want %v", c.enumKey, probe[c.enumKey], c.enumWant)
			}
			// 分发往返不得 panic, 枚举值保留
			if _, err := Unmarshal(c.name, raw); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}
		})
	}
}

// 字段值保留: 抽样若干类型填非零值, 往返后逐字段相等
func TestFieldPreservation(t *testing.T) {
	ca := NewCrossArm()
	ca.CrossArmType = CrossArmTruss
	ca.BeamLength = 2800
	ca.BraceLength = 1500
	ca.MountHeight = 7000
	ca.BoltCount = 4
	assertRoundTripEqual(t, "RAILWAY/CrossArm", ca)

	lc := NewLevelCantilever()
	lc.Length = 3000
	lc.MountHeight = 6800
	lc.RiseAngle = 3
	assertRoundTripEqual(t, "RAILWAY/LevelCantilever", lc)

	mb := NewMastBracket()
	mb.MastDiameter = 350
	mb.InsulatorBoltSpacing = 120
	assertRoundTripEqual(t, "RAILWAY/MastBracket", mb)

	cm := NewConcreteMast()
	cm.SectionType = ConcreteMastCircularHoled
	cm.HoleLength = 400
	assertRoundTripEqual(t, "RAILWAY/ConcreteMast", cm)

	cu := NewCurvedArm()
	cu.CurvedArmType = CurvedArmLShape
	cu.BendAngle = 90
	assertRoundTripEqual(t, "RAILWAY/CurvedArm", cu)

	pc := NewPositioningCable()
	pc.Diameter = 9
	pc.TopPoint = [3]float64{1000, 2000, 7000}
	pc.BottomPoint = [3]float64{1000, 2000, 400}
	pc.Adjustable = true
	assertRoundTripEqual(t, "RAILWAY/PositioningCable", pc)

	st := NewStraightTrack()
	st.StartPoint = [3]float64{0, 0, 0}
	st.EndPoint = [3]float64{50000, 0, 0}
	st.Gauge = 1435
	st.BallastSlope = 1.5
	assertRoundTripEqual(t, "RAILWAY/StraightTrack", st)

	ct := NewCurveTrack()
	ct.CurveCenter = [3]float64{0, 10000, 0}
	ct.SweepAngle = 30
	ct.SuperElevation = 120
	assertRoundTripEqual(t, "RAILWAY/CurveTrack", ct)

	rc := NewRatchetCompensator()
	rc.WheelDiameter = 400
	rc.Stack = WeightStackParams{BlockCount: 10, BlockDiameter: 380}
	assertRoundTripEqual(t, "RAILWAY/RatchetCompensator", rc)

	rp := NewRailPair()
	rp.Centerline = [][3]float64{{0, 0, 0}, {1000, 0, 0}, {2000, 500, 0}}
	rp.Gauge = 1435
	assertRoundTripEqual(t, "RAILWAY/RailPair", rp)
}

func assertRoundTripEqual(t *testing.T, ty string, shp Shape) {
	t.Helper()
	raw, err := json.Marshal(shp)
	if err != nil {
		t.Fatalf("%s marshal: %v", ty, err)
	}
	back, err := Unmarshal(ty, raw)
	if err != nil {
		t.Fatalf("%s Unmarshal: %v", ty, err)
	}
	if !reflect.DeepEqual(shp, back) {
		raw2, _ := json.Marshal(back)
		t.Fatalf("%s field mismatch:\nfirst:  %s\nsecond: %s", ty, raw, raw2)
	}
}

// layout 级类型往返: 锚段 + 站场
func TestLayoutRoundTrip(t *testing.T) {
	as := NewAnchorSection()
	as.Spec = AnchorSectionSpec{
		Centerline:      [][3]float64{{0, 0, 0}, {50000, 0, 0}, {100000, 0, 0}},
		ContactHeight:   5300,
		StructureHeight: 1400,
		SpanLength:      50000,
		MastHeight:      8000,
		MastType:        1,
		SideOffset:      2900,
		MastSide:        1,
		HasCompensator:  true,
		ContactWireDia:  12.9,
		MessengerDia:    13.5,
		DropperSpacing:  8000,
	}
	as.Masts = []AnchorMastLayout{
		{
			Mileage:        0,
			Position:       [3]float64{0, -2900, 0},
			Direction:      [3]float64{0, 1, 0},
			MastHeight:     8000,
			ContactWireZ:   5300,
			MessengerWireZ: 6700,
			Stagger:        300,
			IsTensionMast:  true,
			ContactPoint:   [3]float64{0, -300, 5300},
			MessengerPoint: [3]float64{0, -300, 6700},
		},
		{
			Mileage:   50,
			Position:  [3]float64{50000, 2900, 0},
			Direction: [3]float64{0, -1, 0},
			Stagger:   -300,
		},
	}
	as.Spans = []AnchorSpanLayout{
		{
			FromMast:     0,
			ToMast:       1,
			Length:       50000,
			ContactSag:   25,
			MessengerSag: 750,
			Droppers: []AnchorDropperLayout{
				{T: 0.2, Top: [3]float64{10000, 0, 6600}, Bottom: [3]float64{10000, 0, 5300}, Length: 1300},
			},
		},
	}
	assertRoundTripEqual(t, "RAILWAY/AnchorSection", as)

	yard := NewYard()
	yard.Tracks = []YardTrackLayout{
		{
			Centerline: [][3]float64{{0, 0, 0}, {100000, 0, 0}},
			TrimS:      500,
			TrimE:      99500,
			Props: TrackGeoProperties{
				Gauge:            1435,
				RailType:         60,
				SleeperLength:    2600,
				SleeperSpacing:   600,
				BallastTopWidth:  3600,
				BallastThickness: 300,
				BallastSlope:     1.5,
			},
		},
	}
	yard.Turnouts = []YardTurnoutLayout{
		{
			Position:         [3]float64{30000, 0, 0},
			MainDir:          [3]float64{1, 0, 0},
			IsLeftHand:       true,
			TurnoutNo:        12,
			SwitchRailLength: 7700,
			LeadCurveRadius:  350000,
			Gauge:            1435,
			RailHeight:       176,
			RailHeadWidth:    73,
			RailBaseWidth:    150,
			EdgeIn:           0,
			EdgeOut:          0,
			EdgeDiv:          1,
		},
	}
	yard.Crossings = []YardCrossingLayout{
		{
			Position:      [3]float64{60000, 5000, 0},
			Direction:     [3]float64{1, 0, 0},
			TurnoutNo:     9,
			Gauge:         1435,
			RailHeight:    176,
			RailHeadWidth: 73,
			RailBaseWidth: 150,
			EdgeA:         0,
			EdgeB:         1,
		},
	}
	assertRoundTripEqual(t, "RAILWAY/Yard", yard)
}

// 未知类型必须报错
func TestUnmarshalInvalid(t *testing.T) {
	if _, err := Unmarshal("RAILWAY/NoSuchType", []byte(`{}`)); err == nil {
		t.Fatal("expected error for unknown type")
	}
}
