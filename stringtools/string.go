// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

// Package stringtools provides some useful ways to handle strings
package stringtools

import (
	"unicode/utf8"
	"unsafe"
)

// ToString converts the specified byte array to a string.
func ToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// ToBytes converts the specified str to a byte array.
func ToBytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Contains determines whether the str is in the strSlice.
func Contains(str string, strSlice []string) bool {
	for _, v := range strSlice {
		if v == str {
			return true
		}
	}
	return false
}

// SubStr decode str into runes and get substring with the specified length.
func SubStr(str string, length int) (ret string) {
	var count int
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		i += size
		ret += string(r)
		count++
		if length <= count {
			break
		}
	}
	return
}
