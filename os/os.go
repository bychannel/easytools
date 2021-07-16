// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easyos

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// 确定当前操作系统是否为Windows.
func IsWindows() bool {
	return "windows" == runtime.GOOS
}

// 确定当前操作系统是否为Linux.
func IsLinux() bool {
	return "linux" == runtime.GOOS
}

// 确定当前操作系统是否为MacOS.
func IsDarwin() bool {
	return "darwin" == runtime.GOOS
}

// 获取当前工作目录的路径.
func Pwd() string {
	file, _ := exec.LookPath(os.Args[0])
	pwd, _ := filepath.Abs(file)

	return filepath.Dir(pwd)
}

// 返回用户的主目录.
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// 支持交叉编译
	if IsWindows() {
		return homeWindows()
	}

	return homeUnix()
}

func homeUnix() (string, error) {
	// 先尝试home变量
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// 再尝试用shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
