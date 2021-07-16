// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easyencrypt

import "testing"

func TestMd5(t *testing.T) {
	filePath := "./test"
	md5, err := Md5(filePath)
	if err != nil {
		t.Error("计算文件hash失败")
	}

	t.Logf("文件hash值：%s", md5)
}