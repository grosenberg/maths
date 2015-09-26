// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package algorithms

import (
	"sync"

	"github.com/grosenberg/generic"
)

/*
 * The generic type R establishes
 * 1) a constraint definition for minimum implementing instance types; and
 * 2) the implementation API for externalizing the essential, atomic type
 * dependent value manipulation operations.
 */
type R interface {
	sync.Locker
	Add_R(Q)
	Update_R()
	Value_R() Q
	Count_R() Q
	Clear_R()
}

type Q interface{}

/*
 * A generic algorithm - utilizes generic types and thereby allows
 * a single, value type independent implemenetation.
 *
 * Generic methods are implemented purely in terms of generic types and
 * generic type independent instance types, such as int and error. All
 * generic type instance dependent operations are externalized.
 *
 * Generic functions are not supported because the receiver type cannot be
 * an interface.
 */
func Accumulate_R(rcv R, r Q) Q {
	vals := generic.ValueOf(r)
	for i := 0; i < vals.Len(); i++ {
		val := vals.Index(i).Interface()
		rcv.Add_R(val)
	} 
	return rcv.Count_R()
}

func Compute_R(rcv R) Q {
	rcv.Update_R()
	return rcv.Value_R()
}

func Reset_R(rcv R) Q {
	ret := rcv.Count_R()
	rcv.Clear_R()
	return ret
}
