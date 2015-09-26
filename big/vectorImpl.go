// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package big

import (
	"math/big"
	"sync"

	. "github.com/grosenberg/maths/algorithms"
)

// Type specfic composite 'value' compatible with the intended generic type algorithm
// implemenetation.
type Vector struct {
	sync.Mutex
	Elem []Scalar
}

// Type specfic element 'value' compatible with the intended generic type algorithm
// implemenetation.
type Scalar big.Float

/////////////////////////////////////////////////////////////
// Compile-time implementation prover
var _ V = &Vector{}
var _ S = &Scalar{}

/////////////////////////////////////////////////////////////
// Type-specific support for generic implementation

// ... for the V API

// New
func (a *Vector) New_V() V {
	return NewVector(a.Len_V())
}

// Dup (copy)
func (a *Vector) Dup_V() V {
	return a.CopyVector()
}

// Get
func (a *Vector) Get_V(pos int) S {
	return Scalar(a.Elem[pos])
}

// Set
func (a *Vector) Set_V(pos int, b S) {
	a.Elem[pos] = b.(Scalar)
}

// Add
func (a *Vector) Add_V(pos int, b V) {
	x := big.Float(a.Elem[pos])
	y := big.Float(b.(*Vector).Elem[pos])
	x.Add(&x, &y)
	a.Elem[pos] = Scalar(x)
}

// Subtract
func (a *Vector) Sub_V(pos int, b V) {
	x := big.Float(a.Elem[pos])
	y := big.Float(b.(*Vector).Elem[pos])
	x.Sub(&x, &y)
	a.Elem[pos] = Scalar(x)
}

// Multiply
func (a *Vector) Mul_V(pos int, b V) {
	x := big.Float(a.Elem[pos])
	y := big.Float(b.(*Vector).Elem[pos])
	x.Mul(&x, &y)
	a.Elem[pos] = Scalar(x)
}

// Divide
func (a *Vector) Div_V(pos int, b V) {
	x := big.Float(a.Elem[pos])
	y := big.Float(b.(*Vector).Elem[pos])
	x.Quo(&x, &y)
	a.Elem[pos] = Scalar(x)
}

// Multiply Scalar
func (a *Vector) MulSc_V(pos int, b S) {
	x := big.Float(a.Elem[pos])
	y := big.Float(b.(Scalar))
	x.Mul(&x, &y)
	a.Elem[pos] = Scalar(x)
}

// Divide Scalar
func (a *Vector) DivSc_V(pos int, b S) {
	x := big.Float(a.Elem[pos])
	y := big.Float(b.(Scalar))
	x.Quo(&x, &y)
	a.Elem[pos] = Scalar(x)
}

// Negate a vector element
func (a *Vector) Neg_V(pos int) {
	x := big.Float(a.Elem[pos])
	x.Neg(&x)
	a.Elem[pos] = Scalar(x)
}

// Vector length
func (a *Vector) Len_V() int {
	return len(a.Elem)
}

// Minimum relative vector length
func (a *Vector) LenMin_V(b V) int {
	al := len(a.Elem)
	bl := len(b.(*Vector).Elem)
	if al < bl {
		return al
	}
	return bl
}

// ... for the S API

// Add
func (a Scalar) Add_S(b S) S {
	var x, y big.Float
	x = big.Float(a)
	y = big.Float(b.(Scalar))
	z := x.Add(&x, &y)
	return (Scalar)(*z)
}

// Subtract
func (a Scalar) Sub_S(b S) S {
	var x, y big.Float
	x = big.Float(a)
	y = big.Float(b.(Scalar))
	z := x.Sub(&x, &y)
	return (Scalar)(*z)
}

// Multiply
func (a Scalar) Mul_S(b S) S {
	var x, y big.Float
	x = big.Float(a)
	y = big.Float(b.(Scalar))
	z := x.Mul(&x, &y)
	return (Scalar)(*z)

}

// Divide
func (a Scalar) Div_S(b S) S {
	var x, y big.Float
	x = big.Float(a)
	y = big.Float(b.(Scalar))
	z := x.Quo(&x, &y)
	return (Scalar)(*z)
}

// Convert Scalar to a float
func (a Scalar) ToFloat() float64 {
	x := big.Float(a)
	ret, _ := x.Float64()
	return ret
}

// Convert a float to a Scalar - receiver is ignored 
// TODO: receiver should be vector
func (a Scalar) ToS(v float64) S {
	x := big.NewFloat(v)
	return Scalar(*x)
}

// Helper function
// TODO: use generic package

// gen_V promotes the base type to []V
func gen_V(bi ...*Vector) []V {
	b := make([]V, len(bi))
	for i, v := range bi {
		b[i] = v
	}
	return b
}
