// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build reach

package cover

import (
	"io"

	. "github.com/tsavola/reach/internal"
)

func Location() {
	Cover()
}

func Cond(conditions ...bool) {
	Cover(conditions...)
}

func Bool(x bool) bool {
	Cover(!x, x)
	return x
}

func Min(x, min int) int {
	Cover(x == min, x > min)
	return x
}

func MinInt8(x, min int8) int8 {
	Cover(x == min, x > min)
	return x
}

func MinInt16(x, min int16) int16 {
	Cover(x == min, x > min)
	return x
}

func MinInt32(x, min int32) int32 {
	Cover(x == min, x > min)
	return x
}

func MinInt64(x, min int64) int64 {
	Cover(x == min, x > min)
	return x
}

func MinUint(x, min uint) uint {
	Cover(x == min, x > min)
	return x
}

func MinUint8(x, min uint8) uint8 {
	Cover(x == min, x > min)
	return x
}

func MinUint16(x, min uint16) uint16 {
	Cover(x == min, x > min)
	return x
}

func MinUint32(x, min uint32) uint32 {
	Cover(x == min, x > min)
	return x
}

func MinUint64(x, min uint64) uint64 {
	Cover(x == min, x > min)
	return x
}

func MinUintptr(x, min uintptr) uintptr {
	Cover(x == min, x > min)
	return x
}

func MinMax(x, min, max int) int {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func MinMaxInt8(x, min, max int8) int8 {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func MinMaxInt16(x, min, max int16) int16 {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func MinMaxInt32(x, min, max int32) int32 {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func MinMaxInt64(x, min, max int64) int64 {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func MinMaxUint(x, min, max uint) uint {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func MinMaxUint8(x, min, max uint8) uint8 {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func MinMaxUint16(x, min, max uint16) uint16 {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func MinMaxUint32(x, min, max uint32) uint32 {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func MinMaxUint64(x, min, max uint64) uint64 {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func MinMaxUintptr(x, min, max uintptr) uintptr {
	Cover(x == min, x > min && x < max, x == max)
	return x
}

func Error(x error) error {
	Cover(x == nil, x != nil)
	return x
}

func EOF(x error) error {
	Cover(x == nil, x != nil && x != io.EOF, x == io.EOF)
	return x
}
