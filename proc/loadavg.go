// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package proc

const (
	// Format to be used by Fscanf
	LOADAVG_FORMAT = "%f %f %f %d/%d %d\n"
	// /proc file that contains the load average values.
	LOADAVG_FILE = "/proc/loadavg"
)

// Struct to store the loadavg values.
type LoadAvg struct {
	OneMin     float64
	FiveMin    float64
	FifteenMin float64
	Runnable   int
	TotalProcs int
	LastPid    int
}

// LoadAvg.Get() reads /proc/loadavg file and stores the values in the
// LoadAvg struct.
func (l *LoadAvg) Get() error {
	n, err := getvalues(LOADAVG_FILE, LOADAVG_FORMAT,
		&l.OneMin,
		&l.FiveMin,
		&l.FifteenMin,
		&l.Runnable,
		&l.TotalProcs,
		&l.LastPid)
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrEmpty
	}
	return nil
}
