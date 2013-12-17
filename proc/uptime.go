// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proc

const (
	UPTIME_FORMAT = "%f %f\n"
	UPTIME_FILE   = "/proc/uptime"
)

type Uptime struct {
	Up   float64
	Idle float64
}

func (u *Uptime) Get() error {
	n, err := getvalues(UPTIME_FILE, UPTIME_FORMAT,
		&u.Up,
		&u.Idle)
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrEmpty
	}
	return nil
}

func UpTime() float64 {
	u := Uptime{}
	err := u.Get()
	if err != nil {
		return float64(0)
	}
	return u.Up
}
