package joint

import (
	"encoding/json"
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
	j := revoluteJoint() // Axis=(0,0,1) → Z
	j.Limits = &TopoJointLimits{
		RotateZ: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(1)},
	}
	if j.ValidateWith([]float64{100}) {
		t.Error("should fail: RotateZ limit [0,1] violated by 100")
	}
	if !j.ValidateWith([]float64{0.5}) {
		t.Error("should pass: 0.5 within [0,1]")
	}
}

func TestComputeNormalizedRevolute(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointRevolute,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{0, 0, 1},
	}
	// ratio=0 → angle=0 → (1,0,0) stays (1,0,0)
	mat0 := j.ComputeNormalized(0)
	v := vec3d.T{1, 0, 0}
	rv0 := mat0.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv0), [3]float64{1, 0, 0})

	// ratio=0.5 → angle=π → (1,0,0) → (-1,0,0)
	mat05 := j.ComputeNormalized(0.5)
	rv05 := mat05.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv05), [3]float64{-1, 0, 0})

	// ratio=1 → angle=2π → (1,0,0) → (1,0,0)
	mat1 := j.ComputeNormalized(1)
	rv1 := mat1.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv1), [3]float64{1, 0, 0})
}

func TestComputeNormalizedRevoluteWithLimits(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointRevolute,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{0, 0, 1},
		Limits: &TopoJointLimits{
			RotateZ: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
		},
	}
	// ratio=0 → 0, ratio=0.5 → π/2, ratio=1 → π
	v := vec3d.T{1, 0, 0}
	mat := j.ComputeNormalized(0.5)
	rv := mat.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv), [3]float64{0, 1, 0})

	mat2 := j.ComputeNormalized(1)
	rv2 := mat2.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv2), [3]float64{-1, 0, 0})
}

func TestComputeNormalizedPrismatic(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointPrismatic,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{1, 0, 0},
	}
	// ratio=0 → distance=0, ratio=0.5 → 0.5, ratio=1 → 1 (default range [0,1])
	mat0 := j.ComputeNormalized(0)
	almostEq(t, extractTranslate(mat0), [3]float64{0, 0, 0})

	mat05 := j.ComputeNormalized(0.5)
	almostEq(t, extractTranslate(mat05), [3]float64{0.5, 0, 0})

	mat1 := j.ComputeNormalized(1)
	almostEq(t, extractTranslate(mat1), [3]float64{1, 0, 0})
}

func TestComputeNormalizedPrismaticWithLimits(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointPrismatic,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{1, 0, 0},
		Limits: &TopoJointLimits{
			TranslateX: &TopoJointLimitAxis{Min: floatPtr(-5), Max: floatPtr(5)},
		},
	}
	// ratio=0 → -5, ratio=0.5 → 0, ratio=1 → 5
	mat0 := j.ComputeNormalized(0)
	almostEq(t, extractTranslate(mat0), [3]float64{-5, 0, 0})

	mat05 := j.ComputeNormalized(0.5)
	almostEq(t, extractTranslate(mat05), [3]float64{0, 0, 0})

	mat1 := j.ComputeNormalized(1)
	almostEq(t, extractTranslate(mat1), [3]float64{5, 0, 0})
}

func TestComputeNormalizedCylindrical(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointCylindrical,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{0, 1, 0}, // Y axis
		Limits: &TopoJointLimits{
			RotateY:    &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
			TranslateY: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(10)},
		},
	}
	// ratio=0 → angle=0, dist=0
	mat0 := j.ComputeNormalized(0)
	almostEq(t, extractTranslate(mat0), [3]float64{0, 0, 0})
	v := vec3d.T{1, 0, 0}
	rv0 := mat0.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv0), [3]float64{1, 0, 0})

	// ratio=0.5 → angle=π/2, dist=5
	mat05 := j.ComputeNormalized(0.5)
	almostEq(t, extractTranslate(mat05), [3]float64{0, 5, 0})
	rv05 := mat05.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv05), [3]float64{0, 0, -1})
}

func TestComputeNormalizedCurve(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointCurve,
		Origin: [3]float64{0, 0, 0},
		Path: &TopoJointPath{
			Type:   TopoJointPathLine,
			Points: [][3]float64{{0, 0, 0}, {100, 0, 0}},
		},
	}
	// ratio=0 → t=0 → (0,0,0)
	mat0 := j.ComputeNormalized(0)
	almostEq(t, extractTranslate(mat0), [3]float64{0, 0, 0})
	// ratio=0.5 → t=0.5 → (50,0,0)
	mat05 := j.ComputeNormalized(0.5)
	almostEq(t, extractTranslate(mat05), [3]float64{50, 0, 0})
	// ratio=1 → t=1 → (100,0,0)
	mat1 := j.ComputeNormalized(1)
	almostEq(t, extractTranslate(mat1), [3]float64{100, 0, 0})
}

func TestComputeNormalizedClamp(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointRevolute,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{0, 0, 1},
	}
	v := vec3d.T{1, 0, 0}
	matNeg := j.ComputeNormalized(-0.5)
	// clamped to 0 → angle=0 → no rotation
	rvNeg := matNeg.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rvNeg), [3]float64{1, 0, 0})

	matOver := j.ComputeNormalized(1.5)
	// clamped to 1 → angle=2π → full rotation, back to identity
	rvOver := matOver.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rvOver), [3]float64{1, 0, 0})
}

func TestComputeNormalizedFixed(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointFixed,
		Origin: [3]float64{5, 10, 15},
	}
	mat := j.ComputeNormalized(0.5)
	almostEq(t, extractTranslate(mat), [3]float64{5, 10, 15})
	mat2 := j.ComputeNormalized(0)
	almostEq(t, extractTranslate(mat2), [3]float64{5, 10, 15})
}

func TestComputeWithValuesMultiDOF(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointCylindrical,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{0, 1, 0},
		Values: []float64{math.Pi / 2, 4},
	}
	mat := j.Compute()
	pos := extractTranslate(mat)
	almostEq(t, pos, [3]float64{0, 4, 0})
	v := vec3d.T{1, 0, 0}
	rv := mat.MulVec3W(&v, 0)
	almostEq(t, [3]float64(rv), [3]float64{0, 0, -1})
}

func TestComputeValuesPrecedence(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointRevolute,
		Origin: [3]float64{0, 0, 0},
		Axis:   [3]float64{0, 0, 1},
		Value:  floatPtr(0.5),
		Values: []float64{1.57},
	}
	// Values should take priority over Value
	mat := j.Compute()
	v := vec3d.T{1, 0, 0}
	rv := mat.MulVec3W(&v, 0)
	if math.Abs(rv[0]-math.Cos(1.57)) > 1e-5 || math.Abs(rv[1]-math.Sin(1.57)) > 1e-5 {
		t.Errorf("expected Values precedence: got %v", rv)
	}
}

func TestResolveValuesFromJointValue(t *testing.T) {
	j := &TopoJoint{
		Type:  TopoJointRevolute,
		Value: floatPtr(1.0),
	}
	vals := j.ResolveValues(nil)
	if len(vals) != 1 || vals[0] != 1.0 {
		t.Errorf("expected [1], got %v", vals)
	}
}

func TestResolveValuesFromJointValues(t *testing.T) {
	j := &TopoJoint{
		Type:   TopoJointCylindrical,
		Values: []float64{0.5, 3.0},
	}
	vals := j.ResolveValues(nil)
	if len(vals) != 2 || vals[0] != 0.5 || vals[1] != 3.0 {
		t.Errorf("expected [0.5, 3], got %v", vals)
	}
}

func TestResolveValuesFromInstanceStateRatio(t *testing.T) {
	j := &TopoJoint{
		Id:   "test",
		Type: TopoJointRevolute,
	}
	state := &JointInstanceState{Ratio: floatPtr(0.5)}
	vals := j.ResolveValues(state)
	// default range [0, 2π], ratio 0.5 → π
	if len(vals) != 1 || math.Abs(vals[0]-math.Pi) > 1e-10 {
		t.Errorf("expected [π], got %v", vals)
	}
}

func TestResolveValuesFromInstanceStateValues(t *testing.T) {
	j := &TopoJoint{
		Id:   "test",
		Type: TopoJointCylindrical,
	}
	state := &JointInstanceState{Values: []float64{1.0, 5.0}}
	vals := j.ResolveValues(state)
	if len(vals) != 2 || vals[0] != 1.0 || vals[1] != 5.0 {
		t.Errorf("expected [1, 5], got %v", vals)
	}
}

func TestResolveValuesInstanceOverridesJoint(t *testing.T) {
	j := &TopoJoint{
		Id:    "j1",
		Type:  TopoJointRevolute,
		Value: floatPtr(999),
	}
	state := &JointInstanceState{Ratio: floatPtr(0.25)}
	vals := j.ResolveValues(state)
	// instance Ratio overrides joint's Value
	// default range [0, 2π], ratio 0.25 → π/2
	if math.Abs(vals[0]-math.Pi/2) > 1e-10 {
		t.Errorf("expected [π/2], got %v", vals)
	}
}

func TestResolveValuesDefaults(t *testing.T) {
	j := &TopoJoint{Type: TopoJointRevolute}
	vals := j.ResolveValues(nil)
	// default minimums: [0]
	if len(vals) != 1 || vals[0] != 0 {
		t.Errorf("expected [0], got %v", vals)
	}
}

func TestInstanceValueLookup(t *testing.T) {
	j := &TopoJoint{
		Id:   "elbow",
		Type: TopoJointRevolute,
	}
	states := map[int]map[string]JointInstanceState{
		2: {
			"elbow": {Ratio: floatPtr(0.75)},
		},
	}
	vals := j.InstanceValue(states, 2)
	if len(vals) != 1 || math.Abs(vals[0]-1.5*math.Pi) > 1e-10 {
		t.Errorf("instance 2 elbow: expected [1.5π], got %v", vals)
	}
	// instance 0 has no override → falls back to joint defaults → [0]
	vals0 := j.InstanceValue(states, 0)
	if len(vals0) != 1 || vals0[0] != 0 {
		t.Errorf("instance 0: expected [0], got %v", vals0)
	}
}

func TestInstanceValueWithPhysicalValues(t *testing.T) {
	j := &TopoJoint{
		Id:   "piston",
		Type: TopoJointCylindrical,
	}
	states := map[int]map[string]JointInstanceState{
		0: {
			"piston": {Values: []float64{0.5, 3.0}},
		},
	}
	vals := j.InstanceValue(states, 0)
	if len(vals) != 2 || vals[0] != 0.5 || vals[1] != 3.0 {
		t.Errorf("expected [0.5, 3], got %v", vals)
	}
}

func TestJointInstanceStateRoundTrip(t *testing.T) {
	r := 0.75
	state := JointInstanceState{Ratio: &r}
	data, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got JointInstanceState
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Ratio == nil || *got.Ratio != 0.75 {
		t.Errorf("Ratio: got %v", got.Ratio)
	}
	if got.Values != nil {
		t.Error("Values should be nil")
	}
}

func TestJointInstanceStateValuesRoundTrip(t *testing.T) {
	state := JointInstanceState{Values: []float64{1.0, 2.0, 3.0}}
	data, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got JointInstanceState
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Ratio != nil {
		t.Error("Ratio should be nil")
	}
	if len(got.Values) != 3 || got.Values[0] != 1.0 || got.Values[1] != 2.0 || got.Values[2] != 3.0 {
		t.Errorf("Values: got %v", got.Values)
	}
}

// --- ValidateValues tests ---

func TestValidateValuesCorrectLength(t *testing.T) {
	tests := []struct {
		typ   TopoJointType
		vals  []float64
		valid bool
	}{
		{TopoJointFixed, []float64{}, true},
		{TopoJointFixed, []float64{1}, false},
		{TopoJointRevolute, []float64{1.0}, true},
		{TopoJointRevolute, []float64{1.0, 2.0}, false},
		{TopoJointPrismatic, []float64{0.5}, true},
		{TopoJointPrismatic, nil, true},
		{TopoJointCylindrical, []float64{0.5, 3.0}, true},
		{TopoJointCylindrical, []float64{0.5}, false},
		{TopoJointCylindrical, []float64{0.5, 3.0, 7.0}, false},
		{TopoJointPlanar, []float64{1, 2, 3}, true},
		{TopoJointPlanar, []float64{1, 2}, false},
		{TopoJointSpherical, []float64{0.1, 0.2, 0.3}, true},
		{TopoJointSpherical, []float64{0.1, 0.2, 0.3, 0.4}, false},
		{TopoJointUniversal, []float64{0.5, 1.0}, true},
		{TopoJointUniversal, []float64{0.5}, false},
		{TopoJointCurve, []float64{0.5}, true},
		{TopoJointCurve, []float64{0.5, 0.6}, false},
	}
	for _, tt := range tests {
		j := &TopoJoint{Id: "t", Type: tt.typ, Values: tt.vals}
		err := j.ValidateValues()
		if tt.valid && err != nil {
			t.Errorf("type %s values=%v: expected OK, got %v", tt.typ, tt.vals, err)
		}
		if !tt.valid && err == nil {
			t.Errorf("type %s values=%v: expected error", tt.typ, tt.vals)
		}
	}
}

func TestValidateValuesJointWithValueNoValues(t *testing.T) {
	j := &TopoJoint{Id: "r", Type: TopoJointRevolute, Value: floatPtr(1.0)}
	if err := j.ValidateValues(); err != nil {
		t.Errorf("Value-only joint should validate: %v", err)
	}
}

// --- closestAxis limit matching tests ---

func TestDofRangesMatchesRevoluteAxis(t *testing.T) {
	// Axis=Z → only RotateZ should affect range
	j := &TopoJoint{
		Type: TopoJointRevolute,
		Axis: [3]float64{0, 0, 1},
		Limits: &TopoJointLimits{
			RotateX: &TopoJointLimitAxis{Min: floatPtr(-999), Max: floatPtr(-999)},
			RotateZ: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
		},
	}
	mn, mx := j.dofRanges()
	if len(mn) != 1 || mn[0] != 0 {
		t.Errorf("expected min=0 from RotateZ, got %v", mn)
	}
	if len(mx) != 1 || mx[0] != math.Pi {
		t.Errorf("expected max=π from RotateZ, got %v", mx)
	}
}

func TestDofRangesIgnoresNonMatchingAxis(t *testing.T) {
	// Axis=X → RotateY and RotateZ limits are ignored
	j := &TopoJoint{
		Type: TopoJointRevolute,
		Axis: [3]float64{1, 0, 0},
		Limits: &TopoJointLimits{
			RotateY: &TopoJointLimitAxis{Min: floatPtr(-999), Max: floatPtr(-999)},
		},
	}
	mn, mx := j.dofRanges()
	// RotateY doesn't match Axis=X → falls back to default [0, 2π]
	if len(mn) != 1 || mn[0] != 0 {
		t.Errorf("expected default min=0, got %v", mn)
	}
	if len(mx) != 1 || mx[0] != 2*math.Pi {
		t.Errorf("expected default max=2π, got %v", mx)
	}
}

func TestValidateWithAxisMatchedLimit(t *testing.T) {
	// Axis=Z, only RotateZ has limits → validation against RotateZ
	j := &TopoJoint{
		Type: TopoJointRevolute,
		Axis: [3]float64{0, 0, 1},
		Limits: &TopoJointLimits{
			RotateZ: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(1)},
		},
	}
	if !j.ValidateWith([]float64{0.5}) {
		t.Error("0.5 should be within [0,1]")
	}
	if j.ValidateWith([]float64{5}) {
		t.Error("5 should exceed [0,1]")
	}
}

func TestValidateWithNonMatchingLimitIgnored(t *testing.T) {
	// Axis=Z, only RotateX has limits → RotateX is ignored for Z-axis joint
	j := &TopoJoint{
		Type: TopoJointRevolute,
		Axis: [3]float64{0, 0, 1},
		Limits: &TopoJointLimits{
			RotateX: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(1)},
		},
	}
	// RotateX doesn't match Z → no limit applies → any value passes
	if !j.ValidateWith([]float64{999}) {
		t.Error("non-matching limit should be ignored")
	}
}

func TestDofRangesCylindricalMatchesAxis(t *testing.T) {
	// Axis=Y → rotation uses RotateY, translation uses TranslateY
	j := &TopoJoint{
		Type: TopoJointCylindrical,
		Axis: [3]float64{0, 1, 0},
		Limits: &TopoJointLimits{
			RotateY:    &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
			TranslateY: &TopoJointLimitAxis{Min: floatPtr(-5), Max: floatPtr(5)},
		},
	}
	mn, mx := j.dofRanges()
	if len(mn) != 2 || mn[0] != 0 || mn[1] != -5 {
		t.Errorf("mn: got %v, want [0, -5]", mn)
	}
	if len(mx) != 2 || mx[0] != math.Pi || mx[1] != 5 {
		t.Errorf("mx: got %v, want [π, 5]", mx)
	}
}

func TestDofRangesPlanarNormalZ(t *testing.T) {
	// normal=Z → tx=TranslateX, ty=TranslateY, rot=RotateZ
	j := &TopoJoint{
		Type: TopoJointPlanar,
		Axis: [3]float64{0, 0, 1},
		Limits: &TopoJointLimits{
			TranslateX: &TopoJointLimitAxis{Min: floatPtr(-10), Max: floatPtr(10)},
			TranslateY: &TopoJointLimitAxis{Min: floatPtr(-5), Max: floatPtr(5)},
			RotateZ:    &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
		},
	}
	mn, mx := j.dofRanges()
	if len(mn) != 3 || mn[0] != -10 || mn[1] != -5 || mn[2] != 0 {
		t.Errorf("mn: got %v", mn)
	}
	if len(mx) != 3 || mx[0] != 10 || mx[1] != 5 || mx[2] != math.Pi {
		t.Errorf("mx: got %v", mx)
	}
}

func TestDofRangesPlanarNormalX(t *testing.T) {
	// normal=X → in-plane axes are Y and Z, rotation around X
	j := &TopoJoint{
		Type: TopoJointPlanar,
		Axis: [3]float64{1, 0, 0},
		Limits: &TopoJointLimits{
			TranslateY: &TopoJointLimitAxis{Min: floatPtr(-10), Max: floatPtr(10)},
			TranslateZ: &TopoJointLimitAxis{Min: floatPtr(-5), Max: floatPtr(5)},
			RotateX:    &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
		},
	}
	mn, mx := j.dofRanges()
	if len(mn) != 3 || mn[0] != -10 || mn[1] != -5 || mn[2] != 0 {
		t.Errorf("normal=X mn: got %v, want [-10, -5, 0]", mn)
	}
	if len(mx) != 3 || mx[0] != 10 || mx[1] != 5 || mx[2] != math.Pi {
		t.Errorf("normal=X mx: got %v, want [10, 5, π]", mx)
	}
}

func TestDofRangesSphericalUsesAllAxes(t *testing.T) {
	j := &TopoJoint{
		Type: TopoJointSpherical,
		Limits: &TopoJointLimits{
			RotateX: &TopoJointLimitAxis{Min: floatPtr(-1), Max: floatPtr(1)},
			RotateY: &TopoJointLimitAxis{Min: floatPtr(-2), Max: floatPtr(2)},
			RotateZ: &TopoJointLimitAxis{Min: floatPtr(-3), Max: floatPtr(3)},
		},
	}
	mn, mx := j.dofRanges()
	if len(mn) != 3 || mn[0] != -1 || mn[1] != -2 || mn[2] != -3 {
		t.Errorf("mn: got %v", mn)
	}
	if len(mx) != 3 || mx[0] != 1 || mx[1] != 2 || mx[2] != 3 {
		t.Errorf("mx: got %v", mx)
	}
}

func TestDofRangesUniversalMatchesAxes(t *testing.T) {
	sec := [3]float64{0, 1, 0}
	j := &TopoJoint{
		Type:          TopoJointUniversal,
		Axis:          [3]float64{1, 0, 0},
		SecondaryAxis: &sec,
		Limits: &TopoJointLimits{
			RotateX: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi / 2)},
			RotateY: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
		},
	}
	mn, mx := j.dofRanges()
	if len(mn) != 2 || mn[0] != 0 || mn[1] != 0 {
		t.Errorf("mn: got %v", mn)
	}
	if len(mx) != 2 || mx[0] != math.Pi/2 || mx[1] != math.Pi {
		t.Errorf("mx: got %v, want [π/2, π]", mx)
	}
}

func TestResolveValuesFullPriorityChain(t *testing.T) {
	j := &TopoJoint{
		Id:     "chain",
		Type:   TopoJointRevolute,
		Value:  floatPtr(0.1),
		Values: []float64{0.2},
	}
	// Level 0: nil state → joint.Values wins
	vals := j.ResolveValues(nil)
	if len(vals) != 1 || vals[0] != 0.2 {
		t.Errorf("expected joint.Values=[0.2], got %v", vals)
	}
	// Level 1: state with Values → state.Values wins
	vals2 := j.ResolveValues(&JointInstanceState{Values: []float64{0.3}})
	if len(vals2) != 1 || vals2[0] != 0.3 {
		t.Errorf("expected state.Values=[0.3], got %v", vals2)
	}
	// Level 2: state with Ratio → state.Ratio mapped through range
	vals3 := j.ResolveValues(&JointInstanceState{Ratio: floatPtr(0.5)})
	if len(vals3) != 1 || math.Abs(vals3[0]-math.Pi) > 1e-10 {
		t.Errorf("expected state.Ratio=0.5→π, got %v", vals3)
	}
}

func TestComputeNormalizedUsesAxisMatchedDefaults(t *testing.T) {
	// REVOLUTE around Y, Limits only on RotateX → no matching limit → defaults [0, 2π]
	j := &TopoJoint{
		Type: TopoJointRevolute,
		Axis: [3]float64{0, 1, 0},
		Limits: &TopoJointLimits{
			RotateX: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(math.Pi)},
		},
	}
	// ratio=0.25 → default [0, 2π] → π/2
	mat := j.ComputeNormalized(0.25)
	v := vec3d.T{0, 0, 1}
	rv := mat.MulVec3W(&v, 0)
	// Rotate around Y by π/2: (0,0,1) → (1,0,0)
	almostEq(t, [3]float64(rv), [3]float64{1, 0, 0})
}

func TestValidatePrismaticAxisMatched(t *testing.T) {
	// Axis=Y → only TranslateY matters
	j := &TopoJoint{
		Type: TopoJointPrismatic,
		Axis: [3]float64{0, 1, 0},
		Limits: &TopoJointLimits{
			TranslateX: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(1)},
			TranslateY: &TopoJointLimitAxis{Min: floatPtr(-5), Max: floatPtr(5)},
		},
	}
	if j.ValidateWith([]float64{10}) {
		t.Error("TranslateY limit [ -5,5] violated by 10")
	}
	if !j.ValidateWith([]float64{0}) {
		t.Error("0 should be within [-5,5]")
	}
}

func TestCurveValidationWithLimits(t *testing.T) {
	j := &TopoJoint{
		Type: TopoJointCurve,
		Axis: [3]float64{1, 0, 0},
		Limits: &TopoJointLimits{
			TranslateX: &TopoJointLimitAxis{Min: floatPtr(0), Max: floatPtr(0.8)},
		},
	}
	if !j.ValidateWith([]float64{0.5}) {
		t.Error("0.5 should pass")
	}
	if j.ValidateWith([]float64{0.9}) {
		t.Error("0.9 should fail: exceeds TranslateX max 0.8")
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
