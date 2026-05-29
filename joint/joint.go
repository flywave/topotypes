package joint

import (
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
	Value         *float64              `json:"value,omitempty"`
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

// Compute returns the child's local 4x4 transform using the joint's stored Value.
//
// For 1-DOF joints (REVOLUTE, PRISMATIC, CURVE, FIXED) the stored Value is used.
// For multi-DOF joints (CYLINDRICAL, PLANAR, SPHERICAL, UNIVERSAL) secondary DOFs
// default to 0. Use ComputeWith for full control.
func (j *TopoJoint) Compute() *mat4d.T {
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

	switch j.Type {
	case TopoJointFixed:
		return true

	case TopoJointRevolute:
		val := v(0, 0)
		if limits.RotateX != nil && !check(limits.RotateX, val) {
			return false
		}
		if limits.RotateY != nil && !check(limits.RotateY, val) {
			return false
		}
		if limits.RotateZ != nil && !check(limits.RotateZ, val) {
			return false
		}
		return true

	case TopoJointPrismatic:
		val := v(0, 0)
		if limits.TranslateX != nil && !check(limits.TranslateX, val) {
			return false
		}
		if limits.TranslateY != nil && !check(limits.TranslateY, val) {
			return false
		}
		if limits.TranslateZ != nil && !check(limits.TranslateZ, val) {
			return false
		}
		return true

	case TopoJointCylindrical:
		rot := v(0, 0)
		trans := v(1, 0)
		if limits.RotateX != nil && !check(limits.RotateX, rot) {
			return false
		}
		if limits.RotateY != nil && !check(limits.RotateY, rot) {
			return false
		}
		if limits.RotateZ != nil && !check(limits.RotateZ, rot) {
			return false
		}
		if limits.TranslateX != nil && !check(limits.TranslateX, trans) {
			return false
		}
		if limits.TranslateY != nil && !check(limits.TranslateY, trans) {
			return false
		}
		if limits.TranslateZ != nil && !check(limits.TranslateZ, trans) {
			return false
		}
		return true

	case TopoJointPlanar:
		if limits.TranslateX != nil && !check(limits.TranslateX, v(0, 0)) {
			return false
		}
		if limits.TranslateY != nil && !check(limits.TranslateY, v(1, 0)) {
			return false
		}
		if limits.RotateZ != nil && !check(limits.RotateZ, v(2, 0)) {
			return false
		}
		return true

	case TopoJointSpherical:
		if limits.RotateX != nil && !check(limits.RotateX, v(0, 0)) {
			return false
		}
		if limits.RotateY != nil && !check(limits.RotateY, v(1, 0)) {
			return false
		}
		if limits.RotateZ != nil && !check(limits.RotateZ, v(2, 0)) {
			return false
		}
		return true

	case TopoJointUniversal:
		if limits.RotateX != nil && !check(limits.RotateX, v(0, 0)) {
			return false
		}
		if limits.RotateY != nil && !check(limits.RotateY, v(1, 0)) {
			return false
		}
		return true

	case TopoJointCurve:
		if limits.TranslateX != nil || limits.TranslateY != nil || limits.TranslateZ != nil {
			t := v(0, 0)
			if t < 0 || t > 1 {
				return false
			}
			if limits.TranslateX != nil && !check(limits.TranslateX, t) {
				return false
			}
			if limits.TranslateY != nil && !check(limits.TranslateY, t) {
				return false
			}
			if limits.TranslateZ != nil && !check(limits.TranslateZ, t) {
				return false
			}
		}
		return true
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
