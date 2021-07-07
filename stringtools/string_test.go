// Copyright (c) [2020] [provided by bychannel]
// [easytools] is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//			http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.
package stringtools

import "testing"

func TestToBytes(t *testing.T) {
	str := "hello 你好！"
	bytes := ToBytes(str)
	if str2 := ToString(bytes); str != str2 {
		t.Errorf("Str Bytes convert failed [str=%s, str2=%s]", str, str2)
	}
}

