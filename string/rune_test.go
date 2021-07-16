// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easystring

import "testing"

func TestIsNumOrLetter(t *testing.T) {
	if !IsNumOrLetter('0') {
		t.Fail()
	}
	if IsNumOrLetter('@') {
		t.Fail()
	}
}

func TestIsLetter(t *testing.T) {
	if !IsLetter(rune('a')) {
		t.Fail()
	}
	if IsLetter(rune('0')) {
		t.Fail()
	}
}
