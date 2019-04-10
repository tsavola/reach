// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/tsavola/reach/cover"
)

func f1() {
	cover.Location()
}

func f2(n int) {
	cover.Cond(n < 0, n == 0, n > 0)
}

func main() {
	f1()
	for i := -10; i < 10; i++ {
		f2(i)
	}
}
