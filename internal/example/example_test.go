// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/tsavola/reach"
)

func TestF1(t *testing.T) {
	f1()
}

func TestF2Negative(t *testing.T) {
	f2(-10)
}

func TestF2Zero(t *testing.T) {
	f2(0)
}

func TestF2Positive(t *testing.T) {
	f2(100)
}

func TestMain(m *testing.M) {
	reach.TestMain(m)
}
