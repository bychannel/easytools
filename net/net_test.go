// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easynet

import "testing"

func TestLocalIP(t *testing.T) {
	ip, err := LocalIP()

	if nil != err {
		t.Error(err)
	}

	t.Log(ip)
}

func TestLocalMac(t *testing.T) {
	mac, err := LocalMac()

	if nil != err {
		t.Error(err)
	}

	t.Log(mac)
}
