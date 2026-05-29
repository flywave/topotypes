package joint

import (
	"fmt"
	"math"

	mat4d "github.com/flywave/go3d/float64/mat4"
	quatd "github.com/flywave/go3d/float64/quaternion"
	vec3d "github.com/flywave/go3d/float64/vec3"
)

type TopoJointType string

const (
	TopoJointFixed       TopoJointType = "FIXED"
	TopoJointRevolute    TopoJointType = "REVOLUTE"
	TopoJointPrismatic   TopoJointType = "PRISMATIC"
	TopoJointCylindrical TopoJointType = "CYLINDRICAL"
	TopoJointPlanar      TopoJointType = "PLANAR"
	TopoJointSpherical   TopoJointType = "SPHERICAL"
	TopoJointUniversal   TopoJointType = "UNIVERSAL"
	TopoJointCurve       TopoJointType = "CURVE"
)

func (t TopoJointType) ToInt() int {
	switch t {
	case TopoJointFixed:
		return 0
	case TopoJointRevolute:
		return 1
	case TopoJointPrismatic:
		return 2
	case TopoJointCylindrical:
		return 3
	case TopoJointPlanar:
		return 4
	case TopoJointSpherical:
		return 5
	case TopoJointUniversal:
		return 6
	case TopoJointCurve:
		return 7
	default:
		return 0
	}
}

type TopoJointOutputMode string

const (
	TopoJointOutputStatic  TopoJointOutputMode = "STATIC"
	TopoJointOutputDynamic TopoJointOutputMode = "DYNAMIC"
)

func (m TopoJointOutputMode) ToInt() int {
	switch m {
	case TopoJointOutputStatic:
		return 0
	case TopoJointOutputDynamic:
		return 1
	default:
		return 0
	}
}

type TopoJointLimitAxis struct {
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}

type TopoJointLimits struct {
	TranslateX *TopoJointLimitAxis `json:"translateX,omitempty"`
	TranslateY *TopoJointLimitAxis `json:"translateY,omitempty"`
	TranslateZ *TopoJointLimitAxis `json:"translateZ,omitempty"`
	RotateX    *TopoJointLimitAxis `json:"rotateX,omitempty"`
	RotateY    *TopoJointLimitAxis `json:"rotateY,omitempty"`
	RotateZ    *TopoJointLimitAxis `json:"rotateZ,omitempty"`
}

type TopoJointPathType string

const (
	TopoJointPathLine   TopoJointPathType = "LINE"
	TopoJointPathArc    TopoJointPathType = "ARC"
	TopoJointPathBezier TopoJointPathType = "BEZIER"
)

type TopoJointPath struct {
	Type   TopoJointPathType `json:"type"`
	Points [][3]float64      `json:"points"`
}

type TopoJoint struct {
	Id            string                `json:"id"`
	Name          string                `json:"name,omitempty"`
	Type          TopoJointType         `json:"type"`
	Output        TopoJointOutputMode   `json:"output"`
	Parent        string                `json:"parent"`
	Child         string                `json:"child"`
	Origin        [3]float64            `json:"origin"`
	Axis          [3]float64            `json:"axis"`
	SecondaryAxis *[3]float64           `json:"secondaryAxis,omitempty"`
	Limits        *TopoJointLimits      `json:"limits,omitempty"`
	Path          *TopoJointPath        `json:"path,omitempty"`
	// Value is the single-DOF physical value (radians for rotary, meters for prismatic).
	// For 1-DOF joints (FIXED, REVOLUTE, PRISMATIC, CURVE). Superseded by Values when both are set.
	Value  *float64 `json:"value,omitempty"`
	// Values holds per-DOF physical values for multi-DOF joints (CYLINDRICAL, PLANAR, SPHERICAL, UNIVERSAL).
	// Takes priority over Value in Compute(). Length should match the joint type's DOF count.
	Values []float64 `json:"values,omitempty"`
}

// JointInstanceState holds per-instance state for a single joint.
// Either Ratio (normalized [0,1], all DOFs proportional) or Values
// (independent physical DOF values) is used; Ratio takes precedence.
type JointInstanceState struct {
	Ratio  *float64  `json:"ratio,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

func NewTopoJoint() *TopoJoint {
	return &TopoJoint{
		Type:   TopoJointFixed,
		Output: TopoJointOutputStatic,
	}
}

type TopoJointRef struct {
	Ref string `json:"ref"`
}

// Compute returns the child's local 4x4 transform using the joint's stored
// Values (multi-DOF) or Value (1-DOF), falling back to zero defaults.
func (j *TopoJoint) Compute() *mat4d.T {
	if len(j.Values) > 0 {
		return j.ComputeWith(j.Values)
	}
	if j.Value != nil {
		return j.ComputeWith([]float64{*j.Value})
	}
	return j.ComputeWith(nil)
}

// ComputeWith returns the child's local 4x4 transform.
//
// DOF value order per joint type:
//
//	FIXED:       none
//	REVOLUTE:    values[0] = angle (radians)
//	PRISMATIC:   values[0] = distance along Axis
//	CYLINDRICAL: values[0] = angle, values[1] = distance
//	PLANAR:      values[0] = tx, values[1] = ty, values[2] = rotation around normal (radians)
//	SPHERICAL:   values[0..2] = euler angles (rx, ry, rz) in radians
//	UNIVERSAL:   values[0] = angle around Axis, values[1] = angle around SecondaryAxis
//	CURVE:       values[0] = t ∈ [0, 1]
func (j *TopoJoint) ComputeWith(values []float64) *mat4d.T {
	v := func(i int, def float64) float64 {
		if i < len(values) {
			return values[i]
		}
		return def
	}

	origin := (*vec3d.T)(&j.Origin)
	one := vec3d.T{1, 1, 1}

	switch j.Type {
	case TopoJointFixed:
		mat := mat4d.Ident
		mat.Translate(origin)
		return &mat

	case TopoJointRevolute:
		q := quatd.FromAxisAngle((*vec3d.T)(&j.Axis), v(0, 0))
		return mat4d.Compose(origin, &q, &one)

	case TopoJointPrismatic:
		dir := (*vec3d.T)(&j.Axis)
		if dir.IsZero() {
			dir = &vec3d.UnitX
		}
		dn := dir.Normalized()
		offset := dn.Scaled(v(0, 0))
		mat := mat4d.Ident
		mat.Translate(origin)
		mat.Translate(&offset)
		return &mat

	case TopoJointCylindrical:
		dir := (*vec3d.T)(&j.Axis)
		if dir.IsZero() {
			dir = &vec3d.UnitX
		}
		dn := dir.Normalized()
		offset := dn.Scaled(v(1, 0))
		q := quatd.FromAxisAngle(&dn, v(0, 0))
		mat := *mat4d.Compose(origin, &q, &one)
		tmat := mat4d.Ident
		tmat.Translate(&offset)
		mat.AssignMul(&mat, &tmat)
		return &mat

	case TopoJointPlanar:
		normal := (*vec3d.T)(&j.Axis)
		if normal.IsZero() {
			normal = &vec3d.UnitZ
		}
		nn := normal.Normalized()
		ref := vec3d.UnitX
		d := vec3d.Dot(&nn, &ref)
		if math.Abs(d) > 0.9 {
			ref = vec3d.UnitY
		}
		u := ref
		pd := vec3d.Dot(&nn, &u)
		proj := nn.Scaled(-pd)
		u = vec3d.Add(&u, &proj)
		if u.IsZero() {
			u = vec3d.UnitY
			pd2 := vec3d.Dot(&nn, &u)
			proj2 := nn.Scaled(-pd2)
			u = vec3d.Add(&u, &proj2)
		}
		un := u.Normalized()
		vv := vec3d.Cross(&nn, &un)
		vn := vv.Normalized()
		uo := un.Scaled(v(0, 0))
		vo := vn.Scaled(v(1, 0))
		offset := *uo.Add(&vo)
		q := quatd.FromAxisAngle(&nn, v(2, 0))
		mat := *mat4d.Compose(origin, &q, &one)
		tmat := mat4d.Ident
		tmat.Translate(&offset)
		mat.AssignMul(&mat, &tmat)
		return &mat

	case TopoJointSpherical:
		q := quatd.FromEulerAngles(v(0, 0), v(1, 0), v(2, 0))
		return mat4d.Compose(origin, &q, &one)

	case TopoJointUniversal:
		q1 := quatd.FromAxisAngle((*vec3d.T)(&j.Axis), v(0, 0))
		axis2 := (*vec3d.T)(j.SecondaryAxis)
		if axis2 == nil || axis2.IsZero() {
			axis2 = &vec3d.UnitY
		}
		q2 := quatd.FromAxisAngle(axis2, v(1, 0))
		q := quatd.Mul(&q1, &q2)
		return mat4d.Compose(origin, &q, &one)

	case TopoJointCurve:
		pos := j.evaluateCurve(v(0, 0))
		mat := mat4d.Ident
		mat.Translate(origin)
		mat.Translate((*vec3d.T)(&pos))
		return &mat

	default:
		mat := mat4d.Ident
		mat.Translate(origin)
		return &mat
	}
}

// dofCount returns the number of DOFs for this joint type.
func (j *TopoJoint) dofCount() int {
	switch j.Type {
	case TopoJointFixed:
		return 0
	case TopoJointRevolute, TopoJointPrismatic, TopoJointCurve:
		return 1
	case TopoJointCylindrical, TopoJointUniversal:
		return 2
	case TopoJointPlanar, TopoJointSpherical:
		return 3
	}
	return 0
}

// ValidateValues checks that Values length matches the joint type's DOF count.
// Returns nil when Values is nil or length is correct.
func (j *TopoJoint) ValidateValues() error {
	if len(j.Values) == 0 {
		return nil
	}
	n := j.dofCount()
	if len(j.Values) != n {
		return fmt.Errorf("joint %q: type %s expects %d DOF values, got %d", j.Id, j.Type, n, len(j.Values))
	}
	return nil
}

// closestAxis returns 0 (X), 1 (Y), or 2 (Z) — the cardinal axis that v is closest to.
func closestAxis(v [3]float64) int {
	abs := [3]float64{math.Abs(v[0]), math.Abs(v[1]), math.Abs(v[2])}
	if abs[0] >= abs[1] && abs[0] >= abs[2] {
		return 0
	}
	if abs[1] >= abs[0] && abs[1] >= abs[2] {
		return 1
	}
	return 2
}

var axisLimitNames = [3]string{"RotateX", "RotateY", "RotateZ"}
var transLimitNames = [3]string{"TranslateX", "TranslateY", "TranslateZ"}

// pickLimit returns the limit axis from lim matching the given cardinal axis index,
// or nil if the limit is not set.
func pickLimit(lim *TopoJointLimits, axisIdx int, rot bool) *TopoJointLimitAxis {
	if lim == nil {
		return nil
	}
	if rot {
		switch axisIdx {
		case 0:
			return lim.RotateX
		case 1:
			return lim.RotateY
		case 2:
			return lim.RotateZ
		}
	} else {
		switch axisIdx {
		case 0:
			return lim.TranslateX
		case 1:
			return lim.TranslateY
		case 2:
			return lim.TranslateZ
		}
	}
	return nil
}

// defaultLimit returns the default [min, max] range for a joint type when no Limits are set.
// For rotary DOFs the range is [0, 2π], for linear DOFs [0, 1].
func (j *TopoJoint) defaultRange() (min, max []float64) {
	switch j.Type {
	case TopoJointFixed:
		return nil, nil
	case TopoJointRevolute:
		return []float64{0}, []float64{2 * math.Pi}
	case TopoJointPrismatic:
		return []float64{0}, []float64{1}
	case TopoJointCylindrical:
		return []float64{0, 0}, []float64{2 * math.Pi, 1}
	case TopoJointPlanar:
		return []float64{0, 0, 0}, []float64{1, 1, 2 * math.Pi}
	case TopoJointSpherical:
		return []float64{0, 0, 0}, []float64{2 * math.Pi, 2 * math.Pi, 2 * math.Pi}
	case TopoJointUniversal:
		return []float64{0, 0}, []float64{2 * math.Pi, 2 * math.Pi}
	case TopoJointCurve:
		return []float64{0}, []float64{1}
	}
	return nil, nil
}

// dofRanges returns the per-DOF [min, max] ranges for this joint.
// Uses Limits when set, matching the limit axis to the closest cardinal axis of j.Axis.
// Falls back to defaultRange when Limits are nil.
func (j *TopoJoint) dofRanges() (min, max []float64) {
	defMin, defMax := j.defaultRange()
	lim := j.Limits

	pick := func(axis *TopoJointLimitAxis, def float64) float64 {
		if axis != nil && axis.Min != nil {
			return *axis.Min
		}
		return def
	}
	pickMax := func(axis *TopoJointLimitAxis, def float64) float64 {
		if axis != nil && axis.Max != nil {
			return *axis.Max
		}
		return def
	}

	// Determine which cardinal axis this joint's Axis is closest to, for limit matching.
	axIdx := closestAxis(j.Axis)

	switch j.Type {
	case TopoJointFixed:
		return nil, nil

	case TopoJointRevolute:
		l := pickLimit(lim, axIdx, true)
		return []float64{pick(l, defMin[0])}, []float64{pickMax(l, defMax[0])}

	case TopoJointPrismatic:
		l := pickLimit(lim, axIdx, false)
		return []float64{pick(l, defMin[0])}, []float64{pickMax(l, defMax[0])}

	case TopoJointCylindrical:
		lr := pickLimit(lim, axIdx, true)
		lt := pickLimit(lim, axIdx, false)
		return []float64{pick(lr, defMin[0]), pick(lt, defMin[1])},
			[]float64{pickMax(lr, defMax[0]), pickMax(lt, defMax[1])}

	case TopoJointPlanar:
		// For PLANAR, the normal is j.Axis. Find two perpendicular in-plane axes.
		// We map tx→Translate of in-plane axis 0, ty→Translate of in-plane axis 1,
		// rot→Rotate of the normal axis.
		normIdx := -1
		nv := (*vec3d.T)(&j.Axis)
		if !nv.IsZero() {
			normIdx = closestAxis(j.Axis)
		}
		inPlane0, inPlane1 := (normIdx+1)%3, (normIdx+2)%3
		l0 := pickLimit(lim, inPlane0, false)
		l1 := pickLimit(lim, inPlane1, false)
		lr := pickLimit(lim, normIdx, true)
		return []float64{pick(l0, defMin[0]), pick(l1, defMin[1]), pick(lr, defMin[2])},
			[]float64{pickMax(l0, defMax[0]), pickMax(l1, defMax[1]), pickMax(lr, defMax[2])}

	case TopoJointSpherical:
		l0 := pickLimit(lim, 0, true)
		l1 := pickLimit(lim, 1, true)
		l2 := pickLimit(lim, 2, true)
		return []float64{pick(l0, defMin[0]), pick(l1, defMin[1]), pick(l2, defMin[2])},
			[]float64{pickMax(l0, defMax[0]), pickMax(l1, defMax[1]), pickMax(l2, defMax[2])}

	case TopoJointUniversal:
		a1 := j.Axis
		idx1 := closestAxis(a1)
		idx2 := 0
		if j.SecondaryAxis != nil {
			idx2 = closestAxis(*j.SecondaryAxis)
		} else {
			idx2 = (idx1 + 1) % 3
		}
		l1 := pickLimit(lim, idx1, true)
		l2 := pickLimit(lim, idx2, true)
		return []float64{pick(l1, defMin[0]), pick(l2, defMin[1])},
			[]float64{pickMax(l1, defMax[0]), pickMax(l2, defMax[1])}

	case TopoJointCurve:
		l := pickLimit(lim, axIdx, false)
		return []float64{pick(l, defMin[0])}, []float64{pickMax(l, defMax[0])}
	}
	return nil, nil
}

// ComputeNormalized computes the child's local 4x4 transform from a ratio ∈ [0, 1].
// Each DOF value is mapped as: physical = min + ratio * (max - min).
// When Limits are not set, sensible defaults are used:
//   rotary DOFs → [0, 2π], linear DOFs → [0, 1].
func (j *TopoJoint) ComputeNormalized(ratio float64) *mat4d.T {
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}
	mn, mx := j.dofRanges()
	vals := make([]float64, len(mn))
	for i := range vals {
		vals[i] = mn[i] + ratio*(mx[i]-mn[i])
	}
	return j.ComputeWith(vals)
}

// ResolveValues returns the effective DOF values for this joint.
// Priority: instanceState → joint.Values → joint.Value → default minimums.
func (j *TopoJoint) ResolveValues(instanceState *JointInstanceState) []float64 {
	if instanceState != nil {
		if instanceState.Values != nil {
			return instanceState.Values
		}
		if instanceState.Ratio != nil {
			mn, mx := j.dofRanges()
			vals := make([]float64, len(mn))
			r := *instanceState.Ratio
			if r < 0 {
				r = 0
			}
			if r > 1 {
				r = 1
			}
			for i := range vals {
				vals[i] = mn[i] + r*(mx[i]-mn[i])
			}
			return vals
		}
	}
	if len(j.Values) > 0 {
		return j.Values
	}
	if j.Value != nil {
		return []float64{*j.Value}
	}
	mn, _ := j.dofRanges()
	out := make([]float64, len(mn))
	copy(out, mn)
	return out
}

// InstanceValue resolves the effective DOF values for this joint given
// per-instance overrides. The returned slice contains physical DOF values.
func (j *TopoJoint) InstanceValue(states map[int]map[string]JointInstanceState, idx int) []float64 {
	var state *JointInstanceState
	if states != nil {
		if m, ok := states[idx]; ok {
			if s, ok := m[j.Id]; ok {
				state = &s
			}
		}
	}
	return j.ResolveValues(state)
}

// Validate reports whether the joint's stored Value is within its limit constraints.
// Returns true for joint types that do not define limits.
func (j *TopoJoint) Validate() bool {
	if j.Value != nil {
		return j.ValidateWith([]float64{*j.Value})
	}
	return j.ValidateWith(nil)
}

// ValidateWith reports whether the given DOF values are within the joint's limit constraints.
// The value order matches ComputeWith. Returns true when no limits are configured.
func (j *TopoJoint) ValidateWith(values []float64) bool {
	limits := j.Limits
	if limits == nil {
		return true
	}

	v := func(i int, def float64) float64 {
		if i < len(values) {
			return values[i]
		}
		return def
	}

	check := func(axis *TopoJointLimitAxis, val float64) bool {
		if axis == nil {
			return true
		}
		if axis.Min != nil && val < *axis.Min {
			return false
		}
		if axis.Max != nil && val > *axis.Max {
			return false
		}
		return true
	}

	axIdx := closestAxis(j.Axis)

	switch j.Type {
	case TopoJointFixed:
		return true

	case TopoJointRevolute:
		l := pickLimit(limits, axIdx, true)
		return check(l, v(0, 0))

	case TopoJointPrismatic:
		l := pickLimit(limits, axIdx, false)
		return check(l, v(0, 0))

	case TopoJointCylindrical:
		lr := pickLimit(limits, axIdx, true)
		lt := pickLimit(limits, axIdx, false)
		return check(lr, v(0, 0)) && check(lt, v(1, 0))

	case TopoJointPlanar:
		normIdx := -1
		nv := (*vec3d.T)(&j.Axis)
		if !nv.IsZero() {
			normIdx = closestAxis(j.Axis)
		}
		in0, in1 := (normIdx+1)%3, (normIdx+2)%3
		l0 := pickLimit(limits, in0, false)
		l1 := pickLimit(limits, in1, false)
		lr := pickLimit(limits, normIdx, true)
		return check(l0, v(0, 0)) && check(l1, v(1, 0)) && check(lr, v(2, 0))

	case TopoJointSpherical:
		l0 := pickLimit(limits, 0, true)
		l1 := pickLimit(limits, 1, true)
		l2 := pickLimit(limits, 2, true)
		return check(l0, v(0, 0)) && check(l1, v(1, 0)) && check(l2, v(2, 0))

	case TopoJointUniversal:
		idx1 := closestAxis(j.Axis)
		idx2 := 0
		if j.SecondaryAxis != nil {
			idx2 = closestAxis(*j.SecondaryAxis)
		} else {
			idx2 = (idx1 + 1) % 3
		}
		l1 := pickLimit(limits, idx1, true)
		l2 := pickLimit(limits, idx2, true)
		return check(l1, v(0, 0)) && check(l2, v(1, 0))

	case TopoJointCurve:
		l := pickLimit(limits, axIdx, false)
		t := v(0, 0)
		if t < 0 || t > 1 {
			return false
		}
		return check(l, t)
	}
	return true
}

// evaluateCurve returns the position on the joint's path at parameter t ∈ [0, 1].
func (j *TopoJoint) evaluateCurve(t float64) [3]float64 {
	if j.Path == nil || len(j.Path.Points) == 0 {
		return [3]float64{}
	}
	pts := j.Path.Points
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	switch j.Path.Type {
	case TopoJointPathLine:
		if len(pts) < 2 {
			return pts[0]
		}
		var out [3]float64
		for i := range out {
			out[i] = pts[0][i]*(1-t) + pts[1][i]*t
		}
		return out

	case TopoJointPathArc:
		return evaluateArc(pts, t)

	case TopoJointPathBezier:
		return evaluateBezier(pts, t)
	}
	return [3]float64{}
}

// evaluateArc computes a position on a three-point arc at parameter t.
func evaluateArc(pts [][3]float64, t float64) [3]float64 {
	if len(pts) < 3 {
		return pts[0]
	}
	p0, p1, p2 := vec3d.T(pts[0]), vec3d.T(pts[1]), vec3d.T(pts[2])
	v0 := (*vec3d.T)(&p0)
	v1 := (*vec3d.T)(&p1)
	v2 := (*vec3d.T)(&p2)

	u := vec3d.Sub(v1, v0)
	w := vec3d.Sub(v2, v0)

	cross := vec3d.Cross(&u, &w)
	if cross.IsZero() {
		return pts[0]
	}

	uLen2 := vec3d.Dot(&u, &u)
	wLen2 := vec3d.Dot(&w, &w)
	uw := vec3d.Dot(&u, &w)

	det := uLen2*wLen2 - uw*uw
	if math.Abs(det) < 1e-30 {
		return pts[0]
	}

	p1Len2 := vec3d.Dot(v1, v1)
	p2Len2 := vec3d.Dot(v2, v2)
	r1 := (p1Len2 - vec3d.Dot(v0, v0)) / 2
	r2 := (p2Len2 - vec3d.Dot(v0, v0)) / 2

	rhs1 := r1 - vec3d.Dot(v0, &u)
	rhs2 := r2 - vec3d.Dot(v0, &w)
	s := (rhs1*wLen2 - rhs2*uw) / det
	tt := (rhs2*uLen2 - rhs1*uw) / det

	us := u.Scaled(s)
	ws := w.Scaled(tt)
	center := vec3d.Add(v0, &us)
	center = vec3d.Add(&center, &ws)

	radius := vec3d.Distance(v0, &center)

	dirU := vec3d.Sub(v0, &center)
	du := dirU.Normalized()
	ct := vec3d.Cross(&u, &w)
	norm := ct.Normalized()
	dirV := vec3d.Cross(&norm, &du)
	dv := dirV
	dvn := dv.Normalized()

	arcDot := func(p vec3d.T) float64 {
		rel := vec3d.Sub((*vec3d.T)(&p), &center)
		return math.Atan2(vec3d.Dot(&dvn, &rel), vec3d.Dot(&du, &rel))
	}

	startAngle := arcDot(p0)
	endAngle := arcDot(p2)

	sweep := endAngle - startAngle
	if sweep < 0 {
		sweep += 2 * math.Pi
	}
	if sweep > math.Pi {
		sweep = endAngle - startAngle
		if sweep > 0 {
			sweep -= 2 * math.Pi
		} else {
			sweep += 2 * math.Pi
		}
	}

	angle := startAngle + sweep*t
	du2 := du.Scaled(radius * math.Cos(angle))
	dv2 := dvn.Scaled(radius * math.Sin(angle))
	pos := vec3d.Add(&center, &du2)
	pos2 := vec3d.Add(&pos, &dv2)
	return [3]float64(pos2)
}

// evaluateBezier computes a point on a Bezier curve at parameter t using De Casteljau's algorithm.
func evaluateBezier(pts [][3]float64, t float64) [3]float64 {
	n := len(pts)
	if n == 0 {
		return [3]float64{}
	}
	if n == 1 {
		return pts[0]
	}
	tmp := make([]vec3d.T, n)
	for i, p := range pts {
		tmp[i] = vec3d.T(p)
	}
	for n > 1 {
		for i := 0; i < n-1; i++ {
			s1 := tmp[i].Scaled(1 - t)
			s2 := tmp[i+1].Scaled(t)
			r := s1.Add(&s2)
			tmp[i] = *r
		}
		n--
	}
	return [3]float64(tmp[0])
}
