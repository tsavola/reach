// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !reach
// +build !reach

// Package reach is used to check that code annotated using the cover
// subpackage is covered.
//
// Check function should be called after the test suite is run, e.g. using
// TestMain.  go test should be run with -tags=reach.
package reach

import (
	"fmt"
	"os"
)

// Check all coverage assertions in the Go package that is located in the
// current working directory.  If invoked by go test, the package being tested
// is checked.
func Check(verbose bool) (ok bool) {
	if verbose {
		fmt.Fprintln(os.Stderr, "reach build tag not specified")
	}
	ok = true
	return
}
