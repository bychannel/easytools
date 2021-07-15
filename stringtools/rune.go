// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

// Package stringtools provides some useful ways to handle strings
package stringtools

// IsNumOrLetter checks the specified rune is number or letter.
func IsNumOrLetter(r rune) bool {
	return ('0' <= r && '9' >= r) || IsLetter(r)
}

// IsLetter checks the specified rune is letter.
func IsLetter(r rune) bool {
	return 'a' <= r && 'z' >= r || 'A' <= r && 'Z' >= r
}
