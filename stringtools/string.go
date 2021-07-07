// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//stringtools provides some useful ways to handle strings
package stringtools

import "unsafe"

// 将指定的字节切片转换为字符串.
func ToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// 将指定的字符串转换成字节切片.
func ToBytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}