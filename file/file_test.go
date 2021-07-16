// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easyfile

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

var testdataDir = "data"

func TestIsHidden(t *testing.T) {
	filename := "./file.go"
	isHidden, _ := IsHidden(filename)
	if isHidden {
		t.Error("file [" + filename + "] is not hidden")
	}
}

func TestGetFileSize(t *testing.T) {
	filename := "./file.go"
	size := GetFileSize(filename)
	if 0 > size {
		t.Error("file [" + filename + "] size is [" + strconv.FormatInt(size, 10) + "]")
	}
}

func TestIsExist(t *testing.T) {
	if !IsExist("./data/file.go") {
		t.Error(". must exist")

		return
	}
}

func TestIdBinary(t *testing.T) {
	if IsBinary("not binary content") {
		t.Error("The content should not be binary")

		return
	}
}

func TestIsImg(t *testing.T) {
	if !IsImg(".jpg") {
		t.Error(".jpg should be a valid extension of a image file")

		return
	}
}

func TestIsDir(t *testing.T) {
	if !IsDir(".") {
		t.Error(". should be a directory")

		return
	}
}

func TestCopyDir(t *testing.T) {
	testcopydir := "testcopydir"
	os.Mkdir(testcopydir, 0644)
	dest := filepath.Join(testdataDir, testcopydir)
	defer os.Remove(dest)

	err := CopyDir(testcopydir, dest)
	if nil != err {
		t.Error("Copy dir failed: ", err)

		return
	}
}

func TestCopyFile(t *testing.T) {
	dest := filepath.Join(testdataDir, "file.go")
	//defer os.Remove(dest)
	err := CopyFile("./file.go", dest)
	if nil != err {
		t.Error("Copy file failed: ", err)

		return
	}
}
