// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example for proc.Pid
package main

import (
	"fmt"
	"os"
	"proc"
)

func usage() {
	fmt.Printf("Usage: %s <pid>\n\n", os.Args[0])
}

func main() {
	if len(os.Args) > 1 {
		// get pid from command line
		p := os.Args[1]
		// Declare pid variable
		var pid proc.Pid
		// get info for pid p
		err := pid.Get(p)
		// make sure nothing went wrong
		if err != nil {
			panic(err)
		}
		// Print results
		fmt.Println("Pid: ", pid.Id)
		fmt.Println("Command: ", pid.State.Command)
		fmt.Println("NumThreads: ", pid.State.NumThreads)
		fmt.Println("Processor: ", pid.State.Processor)
		fmt.Println("Size: ", pid.Mem.Size)
		fmt.Println("Data: ", pid.Mem.Data)
		fmt.Println("Dirty: ", pid.Mem.Dirty)
	} else {
		// no pid specified
		usage()
	}
}
