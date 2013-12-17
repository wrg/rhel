// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example for proc.Meminfo
package main

import (
	"fmt"
	"proc"
)

func main() {
	// Decare mem variable
	var mem proc.Meminfo
	// Get system memory info
	err := mem.Get()
	if err != nil {
		fmt.Println(err)
	}
	// Print results
	fmt.Println("MemTotal: ", mem.MemTotal)
	fmt.Println("MemFree: ", mem.MemFree)
	fmt.Println("Buffers: ", mem.Buffers)
	fmt.Println("Cached: ", mem.Cached)
	// Calculate Memory used (include file cache)
	used1 := float64(mem.MemTotal-mem.MemFree) / float64(mem.MemTotal) * 100
	fmt.Println("Percent Used (Method 1): ", used1)
	// Calculate Memory used (exclude file cache)
	used2 := float64(mem.MemTotal-mem.MemFree-mem.Cached) / float64(mem.MemTotal) * 100
	fmt.Println("Percent Used (Method 2): ", used2)
}
