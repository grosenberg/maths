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
type Register struct {
	sync.Mutex
	reg   float64 // store for the final (or current) computed value
	accum float64 // accumulator for interim values
	count int64   // contribution counter
}

/////////////////////////////////////////////////////////////
// Compile-time implementation prover
var _ R = &Register{}

/////////////////////////////////////////////////////////////
// Type-specific support for generic implementation

// Register adds the value of the given parameter to value of the receiver, modifying the receiver.
func (r *Register) Add_R(q Q) {
	r.accum += q.(float64)
	r.count++
}

// Update computes and stores the current average based on the value of the
// given parameter, modifying the receiver.
func (r *Register) Update_R() {
	r.reg = r.accum / float64(r.count)
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
