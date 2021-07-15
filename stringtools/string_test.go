// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package stringtools

import "testing"

func TestToBytes(t *testing.T) {
	str := "hello 你好！"
	bytes := ToBytes(str)
	if str2 := ToString(bytes); str != str2 {
		t.Errorf("Str Bytes convert failed [str=%s, str2=%s]", str, str2)
	}
}
func TestContains(t *testing.T) {
	if !Contains("123", []string{"123", "345"}) {
		t.Error("[\"123\", \"345\"] should contain \"123\"")
		return
	}
}
func TestSubStr(t *testing.T) {
	expected := "foo"
	got := SubStr("foo测试bar", 4)
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}
}
