// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package stringtools

import "testing"

func TestToBytes(t *testing.T) {
	str := "hello 你好！"
	bytes := ToBytes(str)
	if str2 := ToString(bytes); str != str2 {
		t.Errorf("Str Bytes convert failed [str=%s, str2=%s]", str, str2)
	}
}

