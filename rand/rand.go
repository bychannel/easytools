// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easyrand

import (
	"math/rand"
	"time"
)

// Ints returns a random integer array with the specified from, to and size.
func Ints(from, to, size int) []int {
	if to-from < size {
		size = to - from
	}

	var slice []int
	for i := from; i < to; i++ {
		slice = append(slice, i)
	}

	var ret []int
	for i := 0; i < size; i++ {
		idx := rand.Intn(len(slice))
		ret = append(ret, slice[idx])
		slice = append(slice[:idx], slice[idx+1:]...)
	}
	return ret
}

// String returns a random string ['a', 'z'] and ['0', '9'] in the specified length.
func String(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Nanosecond)

	letter := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// Int returns a random integer in range [min, max].
func Int(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Nanosecond)
	return min + rand.Intn(max-min)
}

// 在指定集合中按权重随机出一个元素.
func RandByWeight() {

}

// 判定一次随机事件是否成功.
func IsSuccess() {

}

// 返回一个随机bool值.
func RandBool() {

}