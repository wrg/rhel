// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proc

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Error thrown if no values are found.
var ErrEmpty = errors.New("No Values Found")

// Private function used by proc to read a /proc file using Fscanf to copy the
// values into the appropriate struct.
func getvalues(filename string, format string, a ...interface{}) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	r := bufio.NewReader(file)
	n, err := fmt.Fscanf(r, format, a...)
	return n, err
}
