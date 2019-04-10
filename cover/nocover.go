// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !reach

/*
Package cover is used to assert that code locations and variable values are
covered by tests.

The functions incur no runtime overhead unless the reach build tag was
specified.  Functions which observe a single subject value pass it through.

Example:

	import "github.com/tsavola/reach/cover"

	func example1() {
		cover.Location()
	}

	func example2(i int) {
		cover.Cond(i < 0, i == 0, i > 0)
	}

	func example3(msg string) {
		if cover.Bool(msg != "") {
			log.Print(msg)
		}
	}

	func example4(path string) (err error) {
		f, err := os.Open(path)
		if cover.Error(err) != nil {
			log.Print(err)
			return
		}
		defer f.Close()
		return
	}

The coverage may be checked using the parent package.
*/
package cover

// Location of call must be reached.
func Location() {}

// Cond enumerates custom conditions each of which must be true during at least
// one invocation.
func Cond(conditions ...bool) {}

// Bool's both values must be covered.
func Bool(x bool) bool { return x }

// Min value and some larger value must be covered.
func Min(x, min int) int { return x }

// MinInt8 value and some larger value must be covered.
func MinInt8(x, min int8) int8 { return x }

// MinInt16 value and some larger value must be covered.
func MinInt16(x, min int16) int16 { return x }

// MinInt32 value and some larger value must be covered.
func MinInt32(x, min int32) int32 { return x }

// MinInt64 value and some larger value must be covered.
func MinInt64(x, min int64) int64 { return x }

// MinUint value and some larger value must be covered.
func MinUint(x, min uint) uint { return x }

// MinUint8 value and some larger value must be covered.
func MinUint8(x, min uint8) uint8 { return x }

// MinUint16 value and some larger value must be covered.
func MinUint16(x, min uint16) uint16 { return x }

// MinUint32 value and some larger value must be covered.
func MinUint32(x, min uint32) uint32 { return x }

// MinUint64 value and some larger value must be covered.
func MinUint64(x, min uint64) uint64 { return x }

// MinUintptr value and some larger value must be covered.
func MinUintptr(x, min uintptr) uintptr { return x }

// MinMax requires that minimum value, maximum value and some value between
// them are covered.
func MinMax(x, min, max int) int { return x }

// MinMaxInt8 requires that minimum value, maximum value and some value between
// them are covered.
func MinMaxInt8(x, min, max int8) int8 { return x }

// MinMaxInt16 requires that minimum value, maximum value and some value
// between them are covered.
func MinMaxInt16(x, min, max int16) int16 { return x }

// MinMaxInt32 requires that minimum value, maximum value and some value
// between them are covered.
func MinMaxInt32(x, min, max int32) int32 { return x }

// MinMaxInt64 requires that minimum value, maximum value and some value
// between them are covered.
func MinMaxInt64(x, min, max int64) int64 { return x }

// MinMaxUint requires that minimum value, maximum value and some value between
// them are covered.
func MinMaxUint(x, min, max uint) uint { return x }

// MinMaxUint8 requires that minimum value, maximum value and some value
// between them are covered.
func MinMaxUint8(x, min, max uint8) uint8 { return x }

// MinMaxUint16 requires that minimum value, maximum value and some value
// between them are covered.
func MinMaxUint16(x, min, max uint16) uint16 { return x }

// MinMaxUint32 requires that minimum value, maximum value and some value
// between them are covered.
func MinMaxUint32(x, min, max uint32) uint32 { return x }

// MinMaxUint64 requires that minimum value, maximum value and some value
// between them are covered.
func MinMaxUint64(x, min, max uint64) uint64 { return x }

// MinMaxUintptr requires that minimum value, maximum value and some value
// between them are covered.
func MinMaxUintptr(x, min, max uintptr) uintptr { return x }

// Error must be nil and non-nil.
func Error(x error) error { return x }

// EOF error must be nil, non-nil, and have value io.EOF.
func EOF(x error) error { return x }
