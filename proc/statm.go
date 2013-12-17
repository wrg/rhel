// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package proc

import (
	"fmt"
)

// Format to be used by Fscanf.
const STATM_FORMAT = "%d %d %d %d %d %d %d\n"

// Statm struct to store a processes memory information.
type Statm struct {
	Size     int
	Resident int
	Shared   int
	Text     int
	Lib      int
	Data     int
	Dirty    int
}

// Statm.Get(pid) retrieves the pid's memory information.
// pid should be passed as a string rather than an int.
func (s *Statm) Get(pid string) error {
	filename := fmt.Sprintf("/proc/%s/statm", pid)
	n, err := getvalues(filename, STATM_FORMAT,
		&s.Size,
		&s.Resident,
		&s.Shared,
		&s.Text,
		&s.Lib,
		&s.Data,
		&s.Dirty)
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrEmpty
	}
	return nil
}
