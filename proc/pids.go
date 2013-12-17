// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proc

import (
	"io/ioutil"
	"log"
)

// private function to ensure the /proc/subdir is a pid
// simply tests if the subdir begins with a number.
func isPid(p string) bool {
	switch p[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
	return false
}

// GetPids reads the /proc directory and returns an array of []proc.Pid.
func GetPids() ([]Pid, error) {
	pids := make([]Pid, 0)
	procfs, err := ioutil.ReadDir("/proc")
	if err != nil {
		return pids, err
	}
	for _, f := range procfs {
		if f.IsDir() && isPid(f.Name()) {
			p := Pid{}
			err := p.Get(f.Name())
			if err == nil {
				pids = append(pids, p)
			} else {
				log.Println(err)
			}
		}
	}
	return pids, nil
}
