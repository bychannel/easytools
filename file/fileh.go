// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

// +build !windows

package easyfile

import "path/filepath"

// 检查指定路径的文件是否隐藏.
func IsHidden(path string) (bool, error) {
	path = filepath.Base(path)
	if 1 > len(path) {
		return false, nil
	}
	return "." == path[:1], nil
}
