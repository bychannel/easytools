// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easyfile

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 获取指定路径的文件长度（字节）.
func GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if nil != err {
		return -1
	}
	return fi.Size()
}

// 确定由给定路径指定的文件是否存在.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 确定指定的内容是否为二进制文件内容.
func IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}
	return false
}

// 确定指定的扩展名是否为图片.
func IsImg(extension string) bool {
	ext := strings.ToLower(extension)

	switch ext {
	case ".jpg", ".jpeg", ".bmp", ".gif", ".png", ".svg", ".ico":
		return true
	default:
		return false
	}
}

// 确定指定的路径是否为目录.
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if nil != err {
		return false
	}

	return fio.IsDir()
}

// 将源文件复制到目标文件.
func CopyFile(source string, dest string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		sourceInfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceInfo.Mode())
		}
	}

	return nil
}

// 将源目录复制到目标目录.
func CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, err := os.Open(source)
	if err != nil {
		return err
	}

	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// 递归创建子目录
			err = CopyDir(srcFilePath, destFilePath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcFilePath, destFilePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
