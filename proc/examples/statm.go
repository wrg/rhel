// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example for proc.Statm
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
		var stat proc.Statm
		err := stat.Get(pid)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("PID: ", pid)
		fmt.Println("  Size: ", stat.Size)
		fmt.Println("  Resident: ", stat.Resident)
		fmt.Println("  Shared: ", stat.Shared)
		fmt.Println("  Text: ", stat.Text)
		fmt.Println("  Lib: ", stat.Lib)
		fmt.Println("  Data: ", stat.Data)
		fmt.Println("  Dirty: ", stat.Dirty)
		fmt.Println("-----------\n")
	}
}
