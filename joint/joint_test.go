package joint

import (
	"math"
	"testing"

	mat4d "github.com/flywave/go3d/float64/mat4"
	vec3d "github.com/flywave/go3d/float64/vec3"
)

func fixedJoint() *TopoJoint {
	return &TopoJoint{
		Id:     "j1",
		Type:   TopoJointFixed,
		Output: TopoJointOutputStatic,
		Origin: [3]float64{10, 20, 30},
	}
}

func revoluteJoint() *TopoJoint {
	return &TopoJoint{
		Id:     "j2",
		Type:   TopoJointRevolute,
		Output: TopoJointOutputDynamic,
		Origin: [3]float64{1, 0, 0},
		Axis:   [3]float64{0, 0, 1},
		Value:  floatPtr(math.Pi / 2),
	}
}

func prismaticJoint() *TopoJoint {
	return &TopoJoint{
		Id:     "j3",
		Type:   TopoJointPrismatic,
		Output: TopoJointOutputStatic,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{1, 0, 0},
		Value:  floatPtr(5.0),
	}
}

func cylindricalJoint() *TopoJoint {
	return &TopoJoint{
		Id:     "j4",
		Type:   TopoJointCylindrical,
		Output: TopoJointOutputDynamic,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{0, 1, 0},
	}
}

func planarJoint() *TopoJoint {
	return &TopoJoint{
		Id:     "j5",
		Type:   TopoJointPlanar,
		Output: TopoJointOutputStatic,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{0, 0, 1},
	}
}

func sphericalJoint() *TopoJoint {
	return &TopoJoint{
		Id:     "j6",
		Type:   TopoJointSpherical,
		Output: TopoJointOutputDynamic,
		Origin: [3]float64{0, 1, 0},
	}
}

func universalJoint() *TopoJoint {
	return &TopoJoint{
		Id:            "j7",
		Type:          TopoJointUniversal,
		Output:        TopoJointOutputDynamic,
		Origin:        [3]float64{0, 0, 0},
		Axis:          [3]float64{1, 0, 0},
		SecondaryAxis: &[3]float64{0, 1, 0},
	}
}

func curveJoint() *TopoJoint {
	return &TopoJoint{
		Id:     "j8",
		Type:   TopoJointCurve,
		Output: TopoJointOutputStatic,
		Origin: [3]float64{0, 0, 0},
		Path: &TopoJointPath{
			Type:   TopoJointPathLine,
			Points: [][3]float64{{0, 0, 0}, {10, 0, 0}},
		},
	}
}

func floatPtr(v float64) *float64 { return &v }

func almostEq(t *testing.T, got, want [3]float64) {
	t.Helper()
	for i := range got {
		if math.Abs(got[i]-want[i]) > 1e-10 {
			t.Errorf("index %d: got %v, want %v", i, got, want)
		}
	}
}

func extractTranslate(m *mat4d.T) [3]float64 {
	return [3]float64{m[3][0], m[3][1], m[3][2]}
}

func TestComputeFixed(t *testing.T) {
	j := fixedJoint()
	mat := j.Compute()
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{10, 20, 30})

	mat2 := j.ComputeWith(nil)
	pos2 := extractTranslate(mat2)
	almostEq(t, pos2, [3]float64{10, 20, 30})

	mat3 := j.ComputeWith([]float64{99})
	pos3 := extractTranslate(mat3)
	almostEq(t, pos3, [3]float64{10, 20, 30})
}

func TestComputeRevolute(t *testing.T) {
	j := revoluteJoint()

	mat := j.Compute()
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{1, 0, 0})
	// rotate (1,0,0) by 90° around Z → (0,1,0), w=0 ignores translation
	v := vec3d.T{1, 0, 0}
	rv := mat.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv), [3]float64{0, 1, 0})

	// zero angle
	mat0 := j.ComputeWith([]float64{0})
	v2 := vec3d.T{1, 0, 0}
	rv2 := mat0.MulVec3W(&v2, 0)
	almostEq(t, [3]float64(rv2), [3]float64{1, 0, 0})
}

func TestComputePrismatic(t *testing.T) {
	j := prismaticJoint()
	mat := j.Compute()
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{5, 0, 0})

	mat2 := j.ComputeWith([]float64{10})
	pos2 := extractTranslate(mat2)
	almostEq(t, pos2, [3]float64{10, 0, 0})

	mat3 := j.ComputeWith([]float64{-3})
	pos3 := extractTranslate(mat3)
	almostEq(t, pos3, [3]float64{-3, 0, 0})
}

func TestComputeCylindrical(t *testing.T) {
	j := cylindricalJoint()
	mat := j.ComputeWith([]float64{math.Pi / 2, 4})
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{0, 4, 0})
	// rotate (1,0,0) by 90° around Y → (0,0,-1)
	v := vec3d.T{1, 0, 0}
	rv := mat.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv), [3]float64{0, 0, -1})

	// no rotation, translation only
	mat2 := j.ComputeWith([]float64{0, 7})
	pos2 := extractTranslate(mat2)
	almostEq(t, pos2, [3]float64{0, 7, 0})
}

func TestComputePlanar(t *testing.T) {
	j := planarJoint()
	mat := j.ComputeWith([]float64{3, 4, 0})
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{3, 4, 0})

	// rotate (1,0,0) by 45° around Z → (√2/2, √2/2, 0)
	mat2 := j.ComputeWith([]float64{0, 0, math.Pi / 4})
	v := vec3d.T{1, 0, 0}
	rv := mat2.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv), [3]float64{math.Sqrt2 / 2, math.Sqrt2 / 2, 0})
}

func TestComputeSpherical(t *testing.T) {
	j := sphericalJoint()
	// FromEulerAngles(yHead, xPitch, zRoll)
	// pitch=π/2 = rotate around X by 90° → (0,0,1) → (0,-1,0)
	mat := j.ComputeWith([]float64{0, math.Pi / 2, 0})
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{0, 1, 0})
	v := vec3d.T{0, 0, 1}
	rv := mat.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv), [3]float64{0, -1, 0})
}

func TestComputeUniversal(t *testing.T) {
	j := universalJoint()
	// values[0]=π/2 (around X), values[1]=π/2 (around Y)
	// = Mul(Rx(π/2), Ry(π/2)) → apply Ry first, then Rx
	// (0,0,1) → Ry(π/2): (1,0,0) → Rx(π/2): (1,0,0)
	mat := j.ComputeWith([]float64{math.Pi / 2, math.Pi / 2})
	v := vec3d.T{0, 0, 1}
	rv := mat.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv), [3]float64{1, 0, 0})

	// only around X by 90°: (0,0,1) → (0,-1,0)
	mat2 := j.ComputeWith([]float64{math.Pi / 2, 0})
	v2 := vec3d.T{0, 0, 1}
	rv2 := mat2.MulVec3W(&v2, 0)
	almostEq(t, [3]float64(rv2), [3]float64{0, -1, 0})
}

func TestComputeCurveLine(t *testing.T) {
	j := curveJoint()
	mat := j.Compute()
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{0, 0, 0})

	mat2 := j.ComputeWith([]float64{0.5})
	pos2 := extractTranslate(mat2)
	almostEq(t, pos2, [3]float64{5, 0, 0})

	mat3 := j.ComputeWith([]float64{1})
	pos3 := extractTranslate(mat3)
	almostEq(t, pos3, [3]float64{10, 0, 0})
}

func TestComputeCurveArc(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointCurve,
		Origin: [3]float64{0, 0, 0},
		Path: &TopoJointPath{
			Type:   TopoJointPathArc,
			Points: [][3]float64{{1, 0, 0}, {0, 1, 0}, {-1, 0, 0}},
		},
	}

	pos0 := j.ComputeWith([]float64{0})
	almostEq(t, extractTranslate(pos0), [3]float64{1, 0, 0})

	pos05 := j.ComputeWith([]float64{0.5})
	p05 := extractTranslate(pos05)
	// three-point arc through (1,0,0)→(0,1,0)→(-1,0,0) is a half-circle
	// at t=0.5, should be near (0, 1, 0)
	if vec3d.Distance((*vec3d.T)(&p05), &vec3d.T{0, 1, 0}) > 0.01 {
		t.Errorf("arc t=0.5: got %v, want near (0,1,0)", p05)
	}

	pos2 := j.ComputeWith([]float64{1})
	almostEq(t, extractTranslate(pos2), [3]float64{-1, 0, 0})
}

func TestComputeCurveBezier(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointCurve,
		Origin: [3]float64{0, 0, 0},
		Path: &TopoJointPath{
			Type:   TopoJointPathBezier,
			Points: [][3]float64{{0, 0, 0}, {0, 10, 0}, {10, 10, 0}, {10, 0, 0}},
		},
	}
	pos0 := j.ComputeWith([]float64{0})
	almostEq(t, extractTranslate(pos0), [3]float64{0, 0, 0})

	pos1 := j.ComputeWith([]float64{1})
	almostEq(t, extractTranslate(pos1), [3]float64{10, 0, 0})

	pos05 := j.ComputeWith([]float64{0.5})
	p := extractTranslate(pos05)
	// De Casteljau midpoint of symmetric cubic bezier control points
	if vec3d.Distance((*vec3d.T)(&p), &vec3d.T{5, 7.5, 0}) > 1e-10 {
		t.Errorf("bezier t=0.5: got %v, want (5, 7.5, 0)", p)
	}
}

func TestComputeCurveNilPath(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointCurve,
		Origin: [3]float64{2, 3, 4},
	}
	mat := j.ComputeWith([]float64{0.5})
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{2, 3, 4})
}

func TestComputeCurveClamp(t *testing.T) {
	j := curveJoint()

	mat := j.ComputeWith([]float64{-0.1})
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{0, 0, 0})

	mat2 := j.ComputeWith([]float64{1.5})
	pos2 := extractTranslate(mat2)
	almostEq(t, pos2, [3]float64{10, 0, 0})
}

func TestComputePrismaticZeroAxis(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointPrismatic,
		Origin: [3]float64{2, 3, 4},
		Axis:   [3]float64{0, 0, 0},
	}
	mat := j.ComputeWith([]float64{5})
	pos := extractTranslate(mat)
	// zero axis defaults to UnitX
	almostEq(t, pos, [3]float64{7, 3, 4})
}

func TestComputeRevoluteWithValue(t *testing.T) {
	j := revoluteJoint()
	mat := j.Compute()
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{1, 0, 0})
	v := vec3d.T{1, 0, 0}
	rv := mat.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv), [3]float64{0, 1, 0})
}

func TestComputeSphericalWithValue(t *testing.T) {
	j := sphericalJoint()
	j.Value = floatPtr(math.Pi / 2)
	// value = π/2 used as yHead → rotate around Y by 90°
	// M = T(0,1,0) * Ry(π/2), test direction with w=0
	// (0,0,1) → Ry(π/2): (1,0,0)
	mat := j.Compute()
	v := vec3d.T{0, 0, 1}
	rv := mat.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv), [3]float64{1, 0, 0})
}

func TestValidateFixed(t *testing.T) {
	j := fixedJoint()
	if !j.Validate() {
		t.Error("fixed joint should always validate")
	}
	if !j.ValidateWith(nil) {
		t.Error("fixed joint ValidateWith should return true")
	}
}

func TestValidateRevolute(t *testing.T) {
	j := revoluteJoint()
	j.Limits = &TopoJointLimits{
		RotateX: &TopoJointLimitAxis{Min: floatPtr(-math.Pi), Max: floatPtr(math.Pi)},
		RotateY: &TopoJointLimitAxis{Min: floatPtr(-math.Pi), Max: floatPtr(math.Pi)},
		RotateZ: &TopoJointLimitAxis{Min: floatPtr(-math.Pi), Max: floatPtr(math.Pi)},
	}
	if !j.Validate() {
		t.Error("revolute with Value=π/2 should pass [-π, π]")
	}
	if !j.ValidateWith([]float64{math.Pi}) {
		t.Error("π should pass")
	}
	if !j.ValidateWith([]float64{-math.Pi}) {
		t.Error("-π should pass")
	}
	if j.ValidateWith([]float64{math.Pi + 0.1}) {
		t.Error("π+0.1 should fail")
	}
}

func TestValidatePrismatic(t *testing.T) {
	j := prismaticJoint()
	j.Limits = &TopoJointLimits{
		TranslateX: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(10)},
	}
	if !j.Validate() {
		t.Error("prismatic with Value=5 should pass [0,10]")
	}
	if j.ValidateWith([]float64{15}) {
		t.Error("15 should fail")
	}
	if j.ValidateWith([]float64{-1}) {
		t.Error("-1 should fail")
	}
}

func TestValidateCylindrical(t *testing.T) {
	j := cylindricalJoint()
	j.Value = floatPtr(1.0)
	j.Limits = &TopoJointLimits{
		RotateX:    &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
		RotateY:    &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
		RotateZ:    &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
		TranslateY: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(5)},
	}
	if !j.Validate() {
		t.Error("should pass: angle=1 in [0,π], dist=0 in [0,5]")
	}
	if j.ValidateWith([]float64{math.Pi + 1, 10}) {
		t.Error("should fail: both out of range")
	}
	if j.ValidateWith([]float64{0.5, 6}) {
		t.Error("should fail: trans 6 > 5")
	}
	if !j.ValidateWith([]float64{0.5, 3}) {
		t.Error("should pass: both in range")
	}
}

func TestValidatePlanar(t *testing.T) {
	j := planarJoint()
	j.Limits = &TopoJointLimits{
		TranslateX: &TopoJointLimitAxis{Min: floatPtr(-5), Max: floatPtr(5)},
		TranslateY: &TopoJointLimitAxis{Min: floatPtr(-5), Max: floatPtr(5)},
		RotateZ:    &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
	}
	if !j.ValidateWith([]float64{0, 0, 0}) {
		t.Error("should pass: all in range")
	}
	if j.ValidateWith([]float64{6, 0, 0}) {
		t.Error("should fail: tx=6 > 5")
	}
	if j.ValidateWith([]float64{0, 0, math.Pi + 0.1}) {
		t.Error("should fail: rotation out of range")
	}
	if !j.ValidateWith([]float64{-5, 5, math.Pi / 2}) {
		t.Error("should pass: edge of range")
	}
}

func TestValidateSpherical(t *testing.T) {
	j := sphericalJoint()
	j.Limits = &TopoJointLimits{
		RotateX: &TopoJointLimitAxis{Min: floatPtr(-1), Max: floatPtr(1)},
		RotateY: &TopoJointLimitAxis{Min: floatPtr(-1), Max: floatPtr(1)},
		RotateZ: &TopoJointLimitAxis{Min: floatPtr(-1), Max: floatPtr(1)},
	}
	if !j.ValidateWith([]float64{0, 0, 0}) {
		t.Error("should pass")
	}
	if !j.ValidateWith([]float64{1, -1, 0.5}) {
		t.Error("should pass: all in range")
	}
	if j.ValidateWith([]float64{1.5, 0, 0}) {
		t.Error("should fail: rx=1.5 > 1")
	}
}

func TestValidateUniversal(t *testing.T) {
	j := universalJoint()
	j.Limits = &TopoJointLimits{
		RotateX: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi / 2)},
		RotateY: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi / 2)},
	}
	if j.ValidateWith([]float64{math.Pi, 0}) {
		t.Error("should fail: angle1 out of range")
	}
	if !j.ValidateWith([]float64{math.Pi / 4, math.Pi / 4}) {
		t.Error("should pass: both in range")
	}
}

func TestValidateCurve(t *testing.T) {
	j := curveJoint()
	j.Limits = &TopoJointLimits{
		TranslateX: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(1)},
	}
	if !j.Validate() {
		t.Error("should pass: Value=0 in [0,1]")
	}
	if !j.ValidateWith([]float64{0.5}) {
		t.Error("t=0.5 should pass")
	}
	if j.ValidateWith([]float64{1.5}) {
		t.Error("t=1.5 should fail")
	}
	if j.ValidateWith([]float64{-0.1}) {
		t.Error("t=-0.1 should fail")
	}
}

func TestValidateNilLimits(t *testing.T) {
	j := revoluteJoint()
	j.Value = floatPtr(1000)
	if !j.Validate() {
		t.Error("no limits should always validate")
	}
}

func TestValidateWithPartialLimits(t *testing.T) {
	j := revoluteJoint()
	j.Limits = &TopoJointLimits{
		RotateX: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(1)},
	}
	// when only RotateX has limits, value must satisfy it
	if j.ValidateWith([]float64{100}) {
		t.Error("should fail: RotateX limit [0,1] violated by 100")
	}
	if !j.ValidateWith([]float64{0.5}) {
		t.Error("should pass: 0.5 within [0,1]")
	}
}

func TestDefaultNewJoint(t *testing.T) {
	j := NewTopoJoint()
	if j.Type != TopoJointFixed {
		t.Error("default type should be FIXED")
	}
	if j.Output != TopoJointOutputStatic {
		t.Error("default output should be STATIC")
	}
}

func TestJointRef(t *testing.T) {
	r := TopoJointRef{Ref: "joint_abc"}
	if r.Ref != "joint_abc" {
		t.Error("unexpected ref")
	}
}
