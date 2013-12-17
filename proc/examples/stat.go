// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example for proc.Stat
package main

import (
	"fmt"
	"os"
	"proc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s pid [pid]...\n", os.Args[0])
		os.Exit(1)
	}
	for _, pid := range os.Args[1:] {
		var stat proc.Stat
		err := stat.Get(pid)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Pid: ", stat.Pid)
		fmt.Println("  State: ", string(stat.State))
		fmt.Println("  Command: ", stat.Command)
		fmt.Println("  Utime: ", stat.Utime)
		fmt.Println("  Stime: ", stat.Stime)
		fmt.Println("  Rss: ", stat.Rss)
		fmt.Println("  NumThreads: ", stat.NumThreads)
		fmt.Println("  Processor: ", stat.Processor)
		fmt.Println("  PCpu: ", stat.Pcpu())
		fmt.Println("---------\n")
	}
}
