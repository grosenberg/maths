// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package floats

import (
	"sync"

	. "github.com/grosenberg/maths/algorithms"
)

// Type specfic composite 'value' compatible with the
// intended generic type algorithm implemenetation.
type Vector32 struct {
	sync.Mutex
	Elem []Scalar32
}

// Type specfic simple 'value' compatible with the
// intended generic type algorithm implemenetation.
type Scalar32 float32

/////////////////////////////////////////////////////////////
// Compile-time implementation prover
var _ V = &Vector32{}
var _ S = Scalar32(0)

/////////////////////////////////////////////////////////////
// Type-specific support for generic implementation

// ... for the V API

// New
func (a *Vector32) New_V() V {
	return NewVector32(a.Len_V())
}

// Dup (copy)
func (a *Vector32) Dup_V() V {
	return a.CopyVector32()
}

// Get
func (a *Vector32) Get_V(pos int) S {
	return Scalar32(a.Elem[pos])
}

// Set
func (a *Vector32) Set_V(pos int, b S) {
	a.Elem[pos] = b.(Scalar32)
}

// Add
func (a *Vector32) Add_V(pos int, b V) {
	a.Elem[pos] += b.(*Vector32).Elem[pos]
}

// Subtract
func (a *Vector32) Sub_V(pos int, b V) {
	a.Elem[pos] -= b.(*Vector32).Elem[pos]
}

// Multiply
func (a *Vector32) Mul_V(pos int, b V) {
	a.Elem[pos] *= b.(*Vector32).Elem[pos]
}

// Divide
func (a *Vector32) Div_V(pos int, b V) {
	a.Elem[pos] /= b.(*Vector32).Elem[pos]
}

// Multiply Scalar
func (a *Vector32) MulSc_V(pos int, b S) {
	a.Elem[pos] *= b.(Scalar32)
}

// Divide Scalar
func (a *Vector32) DivSc_V(pos int, b S) {
	a.Elem[pos] /= b.(Scalar32)
}

// Negate a vector element
func (a *Vector32) Neg_V(pos int) {
	a.Elem[pos] = -a.Elem[pos]
}

// Vector length
func (a *Vector32) Len_V() int {
	return len(a.Elem)
}

// Minimum relative vector length
func (a *Vector32) LenMin_V(b V) int {
	al := len(a.Elem)
	bl := len(b.(*Vector32).Elem)
	if al < bl {
		return al
	}
	return bl
}

// ... for the S API

// Add
func (a Scalar32) Add_S(b S) S {
	return a + b.(Scalar32)
}

// Subtract
func (a Scalar32) Sub_S(b S) S {
	return a - b.(Scalar32)
}

// Multiply
func (a Scalar32) Mul_S(b S) S {
	return a * b.(Scalar32)
}

// Divide
func (a Scalar32) Div_S(b S) S {
	return a / b.(Scalar32)
}

// Convert Scalar to a float
func (a Scalar32) ToFloat() float64 {
	return float64(a)
}

// Convert a float to a Scalar - receiver is ignored
// TODO: receiver should be vector
func (a Scalar32) ToS(v float64) S {
	return Scalar32(float32(v))
}

// Helper function
// TODO: use generic package

// gen_V promotes the base type to []V
func gen_V(bi ...*Vector32) []V {
	b := make([]V, len(bi))
	for i, v := range bi {
		b[i] = v
	}
	return b
}
