// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build reach
// +build reach

package reach

import (
	"fmt"
	"os"

	"github.com/tsavola/reach/internal"
)

func Check(verbose bool) (ok bool) {
	ok, err := internal.Check(verbose)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		ok = false
	}
	return
}
