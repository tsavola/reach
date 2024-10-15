// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reach

import (
	"io"
	"os"
	"testing"
)

// PrintLocationsTo writer is used to print diagnostics about the source code
// locations where cover package's functions are called.  Set it to e.g.
// [os.Stderr] before calling Check or TestMain.
var PrintLocationsTo io.Writer = io.Discard

// TestMain can be used to implement a TestMain function which calls Check.
//
// Example:
//
//     package example_test
//
//     import "github.com/tsavola/reach"
//
//     func TestMain(m *testing.M) {
//         reach.TestMain(m)
//     }
//
func TestMain(m *testing.M) {
	r := m.Run()
	if !Check(testing.Verbose()) && r == 0 {
		r = 1
	}
	os.Exit(r)
}
