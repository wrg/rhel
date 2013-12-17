// +build !rhel5

// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is the default meminfo file, RHEL 6 is assumed.

package proc

const (
	// Format to be used by Fscanf (RHEL 6).
	MEMINFO_FORMAT = "MemTotal: %d kB\nMemFree: %d kB\nBuffers: %d kB\nCached: %d kB\nSwapCached: %d kB\nActive: %d kB\nInactive: %d kB\nActive(anon): %d kB\nInactive(anon): %d kB\nActive(file): %d kB\nInactive(file): %d kB\nUnevictable: %d kB\nMlocked: %d kB\nSwapTotal: %d kB\nSwapFree: %d kB\nDirty: %d kB\nWriteback: %d kB\nAnonPages: %d kB\nMapped: %d kB\nShmem: %d kB\nSlab: %d kB\nSReclaimable: %d kB\nSUnreclaim: %d kB\nKernelStack: %d kB\nPageTables: %d kB\nNFS_Unstable: %d kB\nBounce: %d kB\nWritebackTmp: %d kB\nCommitLimit: %d kB\nCommitted_AS: %d kB\nVmallocTotal: %d kB\nVmallocUsed: %d kB\nVmallocChunk: %d kB\nHardwareCorrupted: %d kB\nAnonHugePages: %d kB\nHugePages_Total: %d\nHugePages_Free: %d\nHugePages_Rsvd: %d\nHugePages_Surp: %d\nHugepagesize: %d kB\nDirectMap4k: %d kB\nDirectMap2M: %d kB"
	// /proc file containing system memory information
	MEMINFO_FILE = "/proc/meminfo"
)

// Meminfo struct to store system memory information (RHEL 6)
type Meminfo struct {
	MemTotal          int64
	MemFree           int64
	Buffers           int64
	Cached            int64
	SwapCached        int64
	Active            int64
	Inactive          int64
	Active_anon       int64
	Inactive_anon     int64
	Active_file       int64
	Inactive_file     int64
	Unevictable       int64
	Mlocked           int64
	SwapTotal         int64
	SwapFree          int64
	Dirty             int64
	Writeback         int64
	AnonPages         int64
	Mapped            int64
	Shmem             int64
	Slab              int64
	SReclaimable      int64
	SUnreclaim        int64
	KernelStack       int64
	PageTables        int64
	NFS_Unstable      int64
	Bounce            int64
	WritebackTmp      int64
	CommitLimit       int64
	Committed_AS      int64
	VmallocTotal      int64
	VmallocUsed       int64
	VmallocChunk      int64
	HardwareCorrupted int64
	AnonHugePages     int64
	HugePages_Total   int64
	HugePages_Free    int64
	HugePages_Rsvd    int64
	HugePages_Surp    int64
	Hugepagesize      int64
	DirectMap4k       int64
	DirectMap2M       int64
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
		&m.Active_anon,
		&m.Inactive_anon,
		&m.Active_file,
		&m.Inactive_file,
		&m.Unevictable,
		&m.Mlocked,
		&m.SwapTotal,
		&m.SwapFree,
		&m.Dirty,
		&m.Writeback,
		&m.AnonPages,
		&m.Mapped,
		&m.Shmem,
		&m.Slab,
		&m.SReclaimable,
		&m.SUnreclaim,
		&m.KernelStack,
		&m.PageTables,
		&m.NFS_Unstable,
		&m.Bounce,
		&m.WritebackTmp,
		&m.CommitLimit,
		&m.Committed_AS,
		&m.VmallocTotal,
		&m.VmallocUsed,
		&m.VmallocChunk,
		&m.HardwareCorrupted,
		&m.AnonHugePages,
		&m.HugePages_Total,
		&m.HugePages_Free,
		&m.HugePages_Rsvd,
		&m.HugePages_Surp,
		&m.Hugepagesize,
		&m.DirectMap4k,
		&m.DirectMap2M)
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrEmpty
	}
	return nil
}
