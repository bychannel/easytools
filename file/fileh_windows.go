// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

// +build windows

package easyfile

import (
	"syscall"
)

// 检查指定路径的文件是否隐藏.
func IsHidden(path string) (bool, error) {
	pointer, err := syscall.UTF16PtrFromString(path)
	if nil != err {
		return false, err
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if nil != err {
		return false, err
	}
	return 0 != attributes&syscall.FILE_ATTRIBUTE_HIDDEN, nil
}
