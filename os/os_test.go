// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easyos

import (
	"runtime"
	"testing"
)

func TestIsWindows(t *testing.T) {
	goos := runtime.GOOS

	if "windows" == goos && !IsWindows() {
		t.Error("runtime.GOOS returns [windows]")

		return
	}
}

func TestIsLinux(t *testing.T) {
	goos := runtime.GOOS

	if "linux" == goos && !IsLinux() {
		t.Error("runtime.GOOS returns [linux]")

		return
	}
}

func TestIsDarwin(t *testing.T) {
	goos := runtime.GOOS

	if "darwin" == goos && !IsDarwin() {
		t.Error("runtime.GOOS returns [darwin]")

		return
	}
}

func TestPwd(t *testing.T) {
	pwd := Pwd()
	if "" == pwd {
		t.Error("Working directory should not be empty")

		return
	}

	t.Log(pwd)
}

func TestHome(t *testing.T) {
	home, err := Home()
	if nil != err {
		t.Error("Can not get user home")

		return
	}

	t.Log(home)
}
