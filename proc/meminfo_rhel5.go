// +build rhel5

// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This meminfo file contains the appropriate format and struct for RHEL 5.
// use '-tags rhel5' when building on a RHEL 5 system.
package proc

const (
	// Format to be used by Fscanf (RHEL 5).
	MEMINFO_FORMAT = "MemTotal: %d kB\nMemFree: %d kB\nBuffers: %d kB\nCached: %d kB\nSwapCached: %d kB\nActive: %d kB\nInactive: %d kB\nHighTotal: %d kB\nHighFree: %d kB\nLowTotal: %d kB\nLowFree: %d kB\nSwapTotal: %d kB\nSwapFree: %d kB\nDirty: %d kB\nWriteback: %d kB\nAnonPages: %d kB\nMapped: %d kB\nSlab: %d kB\nPageTables: %d kB\nNFS_Unstable: %d kB\nBounce: %d kB\nCommitLimit: %d kB\nCommitted_AS: %d kB\nVmallocTotal: %d kB\nVmallocUsed: %d kB\nVmallocChunk: %d kB\nHugePages_Total: %d\nHugePages_Free: %d\nHugePages_Rsvd: %d\nHugepagesize: %d kB"
	// /proc file containing system memory information
	MEMINFO_FILE = "/proc/meminfo"
)

// Meminfo struct to store system memory information (RHEL 5)
type Meminfo struct {
	MemTotal        int64
	MemFree         int64
	Buffers         int64
	Cached          int64
	SwapCached      int64
	Active          int64
	Inactive        int64
	HighTotal       int64
	HighFree        int64
	LowTotal        int64
	LowFree         int64
	SwapTotal       int64
	SwapFree        int64
	Dirty           int64
	Writeback       int64
	AnonPages       int64
	Mapped          int64
	Slab            int64
	PageTables      int64
	NFS_Unstable    int64
	Bounce          int64
	CommitLimit     int64
	Committed_AS    int64
	VmallocTotal    int64
	VmallocUsed     int64
	VmallocChunk    int64
	HugePages_Total int64
	HugePages_Free  int64
	HugePages_Rsvd  int64
	Hugepagesize    int64
}

// Meminfo.Get() reads the /proc/meminfo file and stores the values in the Meminfo struct
func (m *Meminfo) Get() error {
	n, err := getvalues(MEMINFO_FILE, MEMINFO_FORMAT,
		&m.MemTotal,
		&m.MemFree,
		&m.Buffers,
		&m.Cached,
		&m.SwapCached,
		&m.Active,
		&m.Inactive,
		&m.HighTotal,
		&m.HighFree,
		&m.LowTotal,
		&m.LowFree,
		&m.SwapTotal,
		&m.SwapFree,
		&m.Dirty,
		&m.Writeback,
		&m.AnonPages,
		&m.Mapped,
		&m.Slab,
		&m.PageTables,
		&m.NFS_Unstable,
		&m.Bounce,
		&m.CommitLimit,
		&m.Committed_AS,
		&m.VmallocTotal,
		&m.VmallocUsed,
		&m.VmallocChunk,
		&m.HugePages_Total,
		&m.HugePages_Free,
		&m.HugePages_Rsvd,
		&m.Hugepagesize)
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrEmpty
	}
	return nil
}
