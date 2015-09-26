// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package big

import (
	"math/big"

	. "github.com/grosenberg/maths/algorithms"
)

// Named vector element positions
const (
	X = iota
	Y
	Z
	W
)

/////////////////////////////////////////////////////////////
// Vector type-specific API

// Create a new Vector of the given dimension
func NewVector(dim int) *Vector {
	v := &Vector{}
	v.Elem = make([]Scalar, dim)
	return v
}

// Create a copy of an existing Vector
func (a *Vector) CopyVector() *Vector {
	b := NewVector(a.Len_V())
	b.AddVectors(a)
	return b
}

// Type specfic wrapper functions for the generic type algorithm implementation ///////////////////////////

// Add a set of vectors to the receiver vector
func (a *Vector) AddVectors(bs ...*Vector) *Vector {
	b := gen_V(bs...)
	return Modify_V(a, AddOp, b...).(*Vector)
}

// Subtract a set of vectors from the receiver vector
func (a *Vector) SubVectors(bs ...*Vector) *Vector {
	b := gen_V(bs...)
	return Modify_V(a, SubOp, b...).(*Vector)
}

// Multiply a set of vectors against the receiver vector
func (a *Vector) MulVectors(bs ...*Vector) *Vector {
	b := gen_V(bs...)
	return Modify_V(a, MulOp, b...).(*Vector)
}

// Divide a set of vectors against the receiver vector
func (a *Vector) DivVectors(bs ...*Vector) *Vector {
	b := gen_V(bs...)
	return Modify_V(a, DivOp, b...).(*Vector)
}

// Multiply a vector by a scalar value
func (a *Vector) MulScalar(val big.Float) *Vector {
	b := NewVector(1)
	b.Elem[0] = Scalar(val)
	return Modify_V(a, MulOp, b).(*Vector)
}

// Divide a vector by a scalar value
func (a *Vector) DivScalar(val big.Float) *Vector {
	b := NewVector(1)
	b.Elem[0] = Scalar(val)
	return Modify_V(a, DivOp, b).(*Vector)
}

// Negate a vector
func (a *Vector) Negate() *Vector {
	return Negate_V(a).(*Vector)
}

// Dot product of two vectors
func (a *Vector) Dot(b *Vector) big.Float {
	return big.Float(Dot_V(a, b).(Scalar))
}

// Linear interpolation
func (a *Vector) Lerp(b *Vector, t big.Float) *Vector {
	return Lerp_V(a, b, Scalar(t)).(*Vector)
}

// Spherical Linear interpolation
func (a *Vector) SLerp(b *Vector, t big.Float) *Vector {
	return SLerp_V(a, b, Scalar(t)).(*Vector)
}
