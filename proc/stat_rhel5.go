// +build rhel5

// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This contains the default stat structure and get function for a processes state.
// This is for RHEL 5, use '-tags rhel5' when building on a RHEL 5 system.

package proc

import (
	"fmt"
)

// Format to be used by Fscanf (RHEL 5).
const STAT_FORMAT = "%d %s %c %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d"

// Stat struct to store the state of a pid. (RHEL 5)
type Stat struct {
	Pid         int
	Command     string
	State       byte
	Ppid        int
	Pgrp        int
	Session     int
	Tty         int
	Tpgid       int
	Flags       uint64
	MinFlt      uint64
	CminFlt     uint64
	MajFlt      uint64
	CmajFlt     uint64
	Utime       uint64
	Stime       uint64
	Cutime      int
	Cstime      int
	Prio        int
	Nice        int
	NumThreads  int
	ItRealValue int
	StartTime   uint64
	Vsize       uint64
	Rss         uint64
	RssLim      uint64
	StartCode   uint64
	EndCode     uint64
	StartStack  uint64
	Kstkesp     uint64
	Kstkeip     uint64
	Signal      uint64
	Blocked     uint64
	SigIgnore   uint64
	SigCatch    uint64
	Wchan       uint64
	Nswap       int
	Cnswap      int
	ExitSig     int
	Processor   int
	RtPrio      int
	Policy      int
	Delayacctbt uint64
}

// Stat.Pcpu() calculates the percent cpu used by the process
func (s *Stat) Pcpu() float64 {
	jiffies := (UpTime() * 100) - float64(s.StartTime)
	return float64((s.Utime+s.Stime)*100) / jiffies
}

// Stat.Get(pid) gets the state of the pid.
// pid should be a string rather than an int.
func (s *Stat) Get(pid string) error {
	filename := fmt.Sprintf("/proc/%s/stat", pid)
	n, err := getvalues(filename, STAT_FORMAT,
		&s.Pid,
		&s.Command,
		&s.State,
		&s.Ppid,
		&s.Pgrp,
		&s.Session,
		&s.Tty,
		&s.Tpgid,
		&s.Flags,
		&s.MinFlt,
		&s.CminFlt,
		&s.MajFlt,
		&s.CmajFlt,
		&s.Utime,
		&s.Stime,
		&s.Cutime,
		&s.Cstime,
		&s.Prio,
		&s.Nice,
		&s.NumThreads,
		&s.ItRealValue,
		&s.StartTime,
		&s.Vsize,
		&s.Rss,
		&s.RssLim,
		&s.StartCode,
		&s.EndCode,
		&s.StartStack,
		&s.Kstkesp,
		&s.Kstkeip,
		&s.Signal,
		&s.Blocked,
		&s.SigIgnore,
		&s.SigCatch,
		&s.Wchan,
		&s.Nswap,
		&s.Cnswap,
		&s.ExitSig,
		&s.Processor,
		&s.RtPrio,
		&s.Policy,
		&s.Delayacctbt)
	if err != nil {
		return err
	}
	if n == 0 {
		fmt.Println("nothing found")
	}
	return nil
}
