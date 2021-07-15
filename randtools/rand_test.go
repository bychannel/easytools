// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package randtools

import "testing"

func TestInts(t *testing.T) {
	ints := Ints(10, 19, 20)
	if 9 != len(ints) {
		t.Fail()
	}
	ints = Ints(10, 19, 5)
	if 5 != len(ints) {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	r1 := String(16)
	r2 := String(16)

	if r1 == r2 {
		t.Fail()
	}
}

func TestInt(t *testing.T) {
	r1 := Int(0, 65535)
	r2 := Int(0, 65535)

	if r1 == r2 {
		t.Fail()
	}
}
