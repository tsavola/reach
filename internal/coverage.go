// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"runtime"
	"sync"
)

type site struct {
	file string
	line int
}

func (s site) String() string {
	return fmt.Sprintf("%s:%d", s.file, s.line)
}

var (
	coverLock sync.Mutex
	coverage  = make(map[site][]bool)
)

//go:noinline
func Cover(conditions ...bool) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		panic("runtime.Caller(2) failed")
	}

	coverLock.Lock()
	defer coverLock.Unlock()

	site := site{file, line}

	if old, found := coverage[site]; found {
		if len(old) != len(conditions) {
			panic("condition count has changed")
		}

		for i, c := range conditions {
			if c {
				old[i] = true
			}
		}
	} else {
		coverage[site] = conditions
	}
}
