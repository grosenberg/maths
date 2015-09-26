// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package ints

import (
	"sync"

	. "github.com/grosenberg/maths/algorithms"
)

/////////////////////////////////////////////////////////////
// Data type specific type

// Type specfic composite 'value' compatible with the intended
// generic type algorithm implemenetation.
type Register struct {
	sync.Mutex
	reg   int // store for the final (or current) computed value
	accum int // accumulator for interim values
	count int // contribution counter
}

/////////////////////////////////////////////////////////////
// Compile-time implementation prover
var _ R = &Register{}

/////////////////////////////////////////////////////////////
// Type-specific support for generic implementation

/*
 * Externalized generic functions - atomic, or essentially atomic, type specialized operators.
 *
 * The intuitively simple implementation
 *
 *     func (a Register) Add(b Register) {
 *         a.register += b.Value
 *     }
 *
 * cannot be used because the signature Add(Register) does not literally match Add(R) in the
 * definition of the generic type R.  While the implicit implementation conversion of
 * R to Register for the receiver is recognized and allowed, the same conversion for a generic
 * parameter type is not. Consequently, the generic type signature and an instance
 * type asssertion must be used. This, unfortunately, creates a hard coupling between
 * each different type instance implementation and the generic type.
 *
 * Problem: self-referential implementation of an interface is not handled
 * Problem: implicit conversion of parameters is not handled
 */

// Add adds the value of the given parameter to value of the receiver, modifying the receiver.
func (r *Register) Add_R(q Q) {
	r.accum += q.(int)
	r.count++
}

// Update computes and stores the current average based on the value of the
// given parameter, modifying the receiver.
func (r *Register) Update_R() {
	r.reg = r.accum / r.count
}

func (r *Register) Value_R() Q {
	return r.reg
}

func (r *Register) Count_R() Q {
	return r.count
}

func (r *Register) Clear_R() {
	r.reg = 0
	r.accum = 0
	r.count = 0
}
