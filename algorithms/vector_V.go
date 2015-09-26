// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package algorithms

import (
	"math"
	"sync"
)

// Vector operations enum
const (
	AddOp = iota
	SubOp
	MulOp
	DivOp
)

// Vector externals - the external generic support API
type V interface {
	sync.Locker
	New_V() V
	Dup_V() V
	Get_V(int) S
	Set_V(int, S)
	Add_V(int, V)
	Sub_V(int, V)
	Mul_V(int, V)
	Div_V(int, V)
	MulSc_V(int, S)
	DivSc_V(int, S)
	Neg_V(int)
	Len_V() int
	LenMin_V(V) int
}

// Scalar externals
type S interface {
	Add_S(S) S
	Sub_S(S) S
	Mul_S(S) S
	Div_S(S) S

	ToFloat() float64
	ToS(float64) S
}

// Generic algorithm for Vector Add, Subtract, Multiply and Divide
func Modify_V(res V, op int, src ...V) V {
	res.Lock()
	defer res.Unlock()
	for _, v := range src {
		i := res.LenMin_V(v)
		for j := 0; j < i; j++ {
			switch op {
			case AddOp:
				res.Add_V(j, v)
			case SubOp:
				res.Sub_V(j, v)
			case MulOp:
				res.Mul_V(j, v)
			case DivOp:
				res.Div_V(j, v)
			}
		}
	}
	return res
}

// Generic algorithm for Vector Multiply and Divide by Scalar
func ModifyScalar_V(res V, op int, t S) V {
	res.Lock()
	defer res.Unlock()
	i := res.Len_V()
	for j := 0; j < i; j++ {
		switch op {
		case MulOp:
			res.MulSc_V(j, t)
		case DivOp:
			res.DivSc_V(j, t)
		}
	}
	return res
}

// Generic algorithm for Vector negation
func Negate_V(res V) V {
	res.Lock()
	defer res.Unlock()
	i := res.Len_V()
	for j := 0; j < i; j++ {
		res.Neg_V(j)
	}
	return res
}

// Generic Dot product
func Dot_V(a, b V) (res S) {
	dim := a.LenMin_V(b)
	for pos := 0; pos < dim; pos++ {
		i := a.Get_V(pos)
		j := b.Get_V(pos)
		res.Add_S(i.Mul_S(j))
	}
	return res
}

// Generic algorithm for linear interpolation
func Lerp_V(a, b V, t S) V {
	tmp := b.Dup_V()
	Modify_V(tmp, SubOp, a)
	ModifyScalar_V(tmp, MulOp, t)
	return Modify_V(tmp, AddOp, a)
}

// Generic algorithm for spherical linear interpolation
func SLerp_V(a, b V, t S) V {
	cosAngle := Dot_V(a, b).ToFloat()
	scale0 := 1.0 - t.ToFloat()
	scale1 := t.ToFloat()
	if cosAngle < 0.999 {
		angle := math.Acos(cosAngle)
		recipSinAngle := 1.0 / math.Sin(angle)
		scale0 = (math.Sin(((1.0 - t.ToFloat()) * angle)) * recipSinAngle)
		scale1 = (math.Sin((t.ToFloat() * angle)) * recipSinAngle)
	}
	ModifyScalar_V(a, MulOp, t.ToS(scale0))
	ModifyScalar_V(b, MulOp, t.ToS(scale1))

	return Modify_V(a, AddOp, b)
}
