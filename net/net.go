// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easynet

import (
	"errors"
	"net"
)

// 获取网卡IP地址.
func LocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if nil != err {
		return "", err
	}

	for _, address := range addrs {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if nil != ipNet.IP.To4() {
				return ipNet.IP.String(), nil
			}
		}
	}

	return "", errors.New("can't get local IP")
}

// 获取网卡MAC地址.
func LocalMac() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, inter := range interfaces {
		address, err := inter.Addrs()
		if err != nil {
			return "", err
		}

		for _, address := range address {
			// 判定是否为回环地址
			if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					return inter.HardwareAddr.String(), nil
				}
			}
		}
	}

	return "", errors.New("can't get local mac")
}
