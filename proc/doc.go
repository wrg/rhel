// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Proc
//
// proc was written to read the /proc filesystem on RHEL 5/6, it can be modified to
// run on other linux distributions.  Distributions sharing the same kernel version
// as RHEL 5/6 will most likely work unmodified.  
//
// The package will read the following information from the /proc filesystem.
//
//   1. /proc/loadavg
//   2. /proc/meminfo
//   3. /proc/uptime
//   4. /proc/stat
//   5. /proc/[pid]/stat
//   6. /proc/[pid]/statm
//
// See the 'proc' man page for details.
package proc
