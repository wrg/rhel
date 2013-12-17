// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Warning this is experimental!  You probably want to try
// the read_pids.go example instead.
package main

import (
	"fmt"
	"os"
	"proc/snapshot"
)

func main() {
	proc, err := snapshot.Snap()
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.OpenFile("proc_snap.out", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	n, err := file.WriteString(fmt.Sprintf("%#v", proc))
	if err != nil || n == 0 {
		fmt.Println("Error writing, ", n, err)
	}
	file.Close()
}
