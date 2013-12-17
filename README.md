proc - read the /proc filesystem (Redhat Enterprise Linux 5/6).
sleuth - monitor memory, cpu, process count, and /tmp

Versions
========

    proc 0.1
    sleuth 0.1

Synopsis
========

The package proc will read the following system information from the /proc
file system:

   1. /proc/loadavg
   2. /proc/meminfo
   3. /proc/uptime
   4. /proc/stat
   5. /proc/[pid]/stat
   6. /proc/[pid]/statm

For more information see the proc man page.

Example usage:
```go
package main

import (
	"fmt"
	"proc"
)

func main() {
	// Declare variable load of proc.LoadAvg type
	var load proc.LoadAvg
	// Get system load average info
	err := load.Get()
	if err != nil {
		fmt.Println(err)
	}
	// Print results
	fmt.Println("1m: ", load.OneMin)
	fmt.Println("5m: ", load.FiveMin)
	fmt.Println("15m: ", load.FifteenMin)
	fmt.Println("Runnable: ", load.Runnable)
	fmt.Println("TotalProcs: ", load.TotalProcs)
	fmt.Println("Last pid: ", load.LastPid)
}
```

Build options:

RHEL 6:
```
go build example.go
```
RHEL 5:
```
go build -tags rhel5 example.go
```


Example output:
```
1m:  0.07
5m:  0.05
15m:  0
Runnable:  1
TotalProcs:  625
Last pid:  26228

```

About
=====

    proc was written to read the /proc filesystem on RHEL 5/6, it can be modified to
    run on other linux distributions.  Distributions sharing the same kernel version
    as RHEL 5/6 will most likely work unmodified.  See the 'proc' man page for details.

    sleuth was written is response to several issues noted in my support role.  Servers
    were locking up due to one of for common reasons.
      1. out of memory
      2. number of processes reached 32K
      3. /tmp was full (application related)
      4. One Minute Load was over 60
    sleuth was intended to attempt to log the offending process so that we could relay the
    information to developers. It polls every 2 seconds as long as no issues are found, and
    will sleep for 30 seconds after an issue is found before logging info again.

Todo
====

   * Support for more distibutions.

Install
=======

The easiest installation of csvsplit is done through goinstall.

    goinstall github.com/wrg/rhel

Documentation
=============

The best way to read the current csvsplit documentation is using
godoc.

    godoc github.com/wrg/rhel

Or better yet, you can run a godoc http server.

    godoc -http=":6060"

Then go to the url http://localhost:6060/pkg/github.com/wrg/csvsplit/

Copyright & License
===================

Copyright (c) 2013, Rick Gibson.
All rights reserved.

Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.
