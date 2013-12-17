// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proc

// Basic struct to store pid information.  See Stat and Statm for details.
type Pid struct {
	Id    string
	State Stat
	Mem   Statm
}

// Calls Stat and Statm for the specified pid.
// The process id (pid) is passed as a string to eliminate
// the need for int to string conversions.
func (p *Pid) Get(pid string) error {
	var err error
	p.Id = pid
	err = p.State.Get(pid)
	if err != nil {
		return err
	}
	err = p.Mem.Get(pid)
	return err
}
