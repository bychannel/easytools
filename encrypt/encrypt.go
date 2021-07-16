// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easyencrypt

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// Md5 Calculates the hash string of the specified file.
func Md5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil { // io.Copy的效率较好
		return "", err
	}

	sum := hash.Sum(nil)
	return hex.EncodeToString(sum), nil
}
