// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package maths

import (
	"github.com/grosenberg/generics/floats"
	"github.com/grosenberg/generics/ints"
	"testing"
)

func TestAverageInts(t *testing.T) {
	if a := ints.Average(2, 4, 6); a != 4 {
		t.Errorf("Wrong. a is %v", a)
	}
}

func TestAverageFloats(t *testing.T) {
	if a := floats.Average(2.95, 4.75, 6.535); a != 4.745 {
		t.Errorf("Wrong. f is %v", a)
	}
}
