// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example for proc.Pids
// This reads all the processes, prints each pid, command number
// of threads, last processor run on, and some memory info.
// it will also report the largest memory user (hog) and cpu user.
package main

import (
	"fmt"
	"os"
	"proc"
)

// function to convert pages to kilobytes
func page2kbytes(p int) int {
	return p * os.Getpagesize() / 1024
}

func main() {
	cpuhog := proc.Pid{}
	memhog := proc.Pid{}
	// read pids and return an array of type proc.Pid
	pids, err := proc.GetPids()
	if err != nil {
		panic(err)
	}
	if len(pids) == 0 {
		panic("no results")
	}
	// read each proc.Pid in the array
	for _, pid := range pids {
		// save if this is the largest memory consumer
		if pid.Mem.Resident > memhog.Mem.Resident {
			memhog = pid
		}
		// save if this is the largest cpu consumer
		if pid.State.Pcpu() > cpuhog.State.Pcpu() {
			cpuhog = pid
		}
		// print results for this proc.Pid
		fmt.Println("Pid: ", pid.Id)
		fmt.Println("Command: ", pid.State.Command)
		fmt.Println("NumThreads: ", pid.State.NumThreads)
		fmt.Println("Processor: ", pid.State.Processor)
		fmt.Println("Size: ", pid.Mem.Size)
		fmt.Println("Data: ", pid.Mem.Data)
		fmt.Println("Dirty: ", pid.Mem.Dirty)
	}
	// print the memory and cpu hogs
	fmt.Printf("Memory Hog: %s - pid: %s mem: %dk\n", memhog.State.Command, memhog.Id, page2kbytes(memhog.Mem.Resident))
	fmt.Printf("CPU Hog: %s - pid %s Pcpu: %.2f\n", cpuhog.State.Command, cpuhog.Id, memhog.State.Pcpu())
}
