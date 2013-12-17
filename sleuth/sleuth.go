// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This was written to montior the following on a RHEL 5/6 System:
// 1. Number of processes
// 2. CPU Load
// 3. Memory usage
// 4. /tmp usage
// Default log file is /var/log/sleuth.log
// To build on a RHEL 5 system run 'go build -tags rhel5 sleuth.go'
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"proc"
	"sleuth/config"
	"syscall"
	"time"
)

const (
	TWOSEC            = 2 * time.Second
	THIRTYSEC         = 30 * time.Second
	ONEMINUTE         = 60 * time.Second
	PROCS_WARNING_FMT = "WARNING Process count high: %d\n"
	PROCS_OK_FMT      = "Process count OK (%d)\n"
	LOAD_WARNING_FMT  = "WARNING Load high: %g, %g, %g Runnable: %d TotalProcs: %d\n"
	LOAD_OK_FMT       = "Load OK (%g)\n"
	MEM_WARNING_FMT   = "WARNING High Memory Usage: %.2f%%\n"
	MEM_OK_FMT        = "Memory Usage OK (%.2f%%)\n"
	TMP_WARNING_FMT   = "WARNING /tmp usage high: %d%%\n"
	TMP_OK_FMT        = "/tmp OK: (%d%% used)\n"
	SEP               = "-----------------------------------------------------------\n"
	KBYTE             = 1024
	MBYTE             = KBYTE * 1024
	GBYTE             = MBYTE * 1024
)

// convert pages to Gb, Mb, or Kb
func page2human(p int) string {
	h := p * os.Getpagesize()
	s := "k"
	switch {
	case h > GBYTE:
		h = h / GBYTE
		s = "g"
	case h > MBYTE:
		h = h / MBYTE
		s = "m"
	default:
		h = h / KBYTE
	}
	return fmt.Sprintf("%d%s", h, s)
}

// Function to print the process that uses the most memory and
// the process that uses the most CPU.
func printhogs() {
	cpuhog := proc.Pid{}
	memhog := proc.Pid{}
	pids, err := proc.GetPids()
	if err != nil {
		log.Println("Error reading proceses: ", err)
		return
	}
	for _, pid := range pids {
		if pid.Mem.Resident > memhog.Mem.Resident {
			memhog = pid
		}
		if pid.State.Pcpu() > cpuhog.State.Pcpu() {
			cpuhog = pid
		}
	}
	log.Printf("Memory Hog: %s - pid: %s mem: %s\n", memhog.State.Command, memhog.Id, page2human(memhog.Mem.Resident))
	log.Printf("CPU Hog: %s - pid %s Pcpu: %.2f%%\n", cpuhog.State.Command, cpuhog.Id, cpuhog.State.Pcpu())
}

type pcount map[string]int

// We had an issue where a process would repeadly spawn itself
// and eventually would fill up the process table (~32K processes)
// this attempts to catch the offending process by reporting the
// name of the process with the most instances running.
func proccounts() {
	pcount := make(pcount, 0)
	pids, err := proc.GetPids()
	if err != nil {
		log.Println("Error reading proceses: ", err)
		return
	}
	for _, pid := range pids {
		pcount[pid.State.Command]++
	}
	pn := "me"
	pc := 0
	for k, v := range pcount {
		if v > pc {
			pc = v
			pn = k
		}
	}
	log.Printf("Process named '%s' has %d instances running\n", pn, pc)
}

// set the configuration file
var configfile = flag.String("c", "/etc/sleuth.cfg", "configuration file")

// not implemented at this time
//func usage() {
//    fmt.Printf("Usage: %s -c config_file\n\n", os.Args[0])
//}

// function to monitor /tmp.
// min value is the percent free threshold.
// if /tmp is less than min % free it will be logged
// in the log file. It is intended to run in a separate thread.
// (will not return)
func tmpmon(min uint64) {
	s := syscall.Statfs_t{}
	tmphigh := true
	for {
		err := syscall.Statfs("/tmp", &s)
		if err != nil {
			log.Println(err)
		}
		pctfree := s.Bavail * 100 / s.Blocks
		pctused := 100 - pctfree
		if pctfree < min {
			log.Printf(TMP_WARNING_FMT, pctused)
			log.Print(SEP)
			tmphigh = true
		} else if tmphigh {
			tmphigh = false
			log.Printf(TMP_OK_FMT, pctused)
			log.Print(SEP)
		}
		time.Sleep(ONEMINUTE)
	}
}

// Main
func main() {
	flag.Parse()
	var load proc.LoadAvg
	var mem proc.Meminfo
	highload := true
	highmem := true
	highprocs := true
        // read config file
	conf := config.Get_config(*configfile)
        // open log file
	logfile, err := os.OpenFile(conf.Log_file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
        // Catch and log kill signals before exiting
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
		s := <-c
		log.Println("Process Stopped, Signal Caught: ", s)
		os.Exit(0)
	}()
	log.SetOutput(logfile)
	log.Println("Sleuth Started")

        // call tmpmon
	go tmpmon(conf.Min_tmp)

        // main loop
	for {
		err := load.Get()
		if err != nil {
			log.Println(err)
		}

		// Check load
		if load.OneMin > conf.Max_load {
			log.Printf(LOAD_WARNING_FMT, load.OneMin, load.FiveMin, load.FifteenMin, load.Runnable, load.TotalProcs)
			highload = true
		} else if highload {
			log.Printf(LOAD_OK_FMT, load.OneMin)
			log.Print(SEP)
			highload = false
		}

		// Check number of processes
		if load.TotalProcs > conf.Max_procs {
			log.Printf(PROCS_WARNING_FMT, load.TotalProcs)
			log.Print(SEP)
			highprocs = true
			proccounts()
		} else if highprocs {
			log.Printf(PROCS_OK_FMT, load.TotalProcs)
			log.Print(SEP)
			highprocs = false
		}

		// Check memory usage
		err = mem.Get()
		if err != nil {
			log.Println(err)
		}
		pct_used := float64(mem.MemTotal-mem.MemFree-mem.Cached+mem.Dirty) / float64(mem.MemTotal) * 100
		if pct_used > conf.Max_mem {
			log.Printf(MEM_WARNING_FMT, pct_used)
			highmem = true
		} else if highmem {
			log.Printf(MEM_OK_FMT, pct_used)
			log.Print(SEP)
			highmem = false
		}

                // Determine how long to sleep
		if highload || highmem {
			printhogs()
			log.Print(SEP)
			time.Sleep(THIRTYSEC)
		} else if highprocs {
			time.Sleep(THIRTYSEC)
		} else {
                        // if all is well sleep 2 seconds
			time.Sleep(TWOSEC)
		}
	}
}
