// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package floats

import . "github.com/grosenberg/maths/algorithms"

/////////////////////////////////////////////////////////////
// Register type-specific implementations

// NewRegister creates a new Register.
func NewRegister() *Register {
	return &Register{}
}

// Accumulate adds the given values to the register values and
// returns the current count of value contributions.
func (reg *Register) Accumulate(b ...int) int {
	reg.Lock()
	defer reg.Unlock()

	return Accumulate_R(reg, b).(int)
}

// Compute updates and returns the calculated value
func (reg *Register) Compute() float64 {
	reg.Lock()
	defer reg.Unlock()

	return Compute_R(reg).(float64)
}

// Reset clears the register values and returns the prior calculated value
func (reg *Register) Reset() float64 {
	reg.Lock()
	defer reg.Unlock()

	return Reset_R(reg).(float64)
}
