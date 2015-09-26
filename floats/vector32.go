// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package floats

import . "github.com/grosenberg/maths/algorithms"

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
func NewVector32(dim int) *Vector32 {
	v := &Vector32{}
	v.Elem = make([]Scalar32, dim)
	return v
}

// Create a copy of an existing Vector
func (a *Vector32) CopyVector32() *Vector32 {
	b := NewVector32(a.Len_V())
	b.AddVectors32(a)
	return b
}

// Add a set of vectors to the receiver vector
func (a *Vector32) AddVectors32(bs ...*Vector32) *Vector32 {
	b := gen_V(bs...)
	return Modify_V(a, AddOp, b...).(*Vector32)
}

// Subtract a set of vectors from the receiver vector
func (a *Vector32) SubVectors32(bs ...*Vector32) *Vector32 {
	b := gen_V(bs...)
	return Modify_V(a, SubOp, b...).(*Vector32)
}

// Multiply a set of vectors against the receiver vector
func (a *Vector32) MulVectors32(bs ...*Vector32) *Vector32 {
	b := gen_V(bs...)
	return Modify_V(a, MulOp, b...).(*Vector32)
}

// Divide a set of vectors against the receiver vector
func (a *Vector32) DivVectors32(bs ...*Vector32) *Vector32 {
	b := gen_V(bs...)
	return Modify_V(a, DivOp, b...).(*Vector32)
}

// Multiply a vector by a scalar value
func (a *Vector32) MulScalar32(val float32) *Vector32 {
	b := NewVector32(1)
	b.Elem[0] = Scalar32(val)
	return Modify_V(a, MulOp, b).(*Vector32)
}

// Divide a vector by a scalar value
func (a *Vector32) DivScalar32(val float32) *Vector32 {
	b := NewVector32(1)
	b.Elem[0] = Scalar32(val)
	return Modify_V(a, DivOp, b).(*Vector32)
}

// Negate a vector
func (a *Vector32) Negate32() *Vector32 {
	return Negate_V(a).(*Vector32)
}

// Dot product of two vectors
func (a *Vector32) Dot32(b *Vector32) float32 {
	return float32(Dot_V(a, b).(Scalar32))
}

// Linear interpolation
func (a *Vector32) Lerp32(b *Vector32, t float32) *Vector32 {
	return Lerp_V(a, b, Scalar32(t)).(*Vector32)
}

// Spherical Linear interpolation
func (a *Vector32) SLerp32(b *Vector32, t float32) *Vector32 {
	return SLerp_V(a, b, Scalar32(t)).(*Vector32)
}
