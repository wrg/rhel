package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Conf struct {
	Max_load  float64
	Max_mem   float64
	Max_procs int
	Min_tmp   uint64
	Log_file  string
}

func new_config() Conf {
	c := Conf{}
	// set defaults
	c.Max_load = float64(3)
	c.Max_mem = float64(80)
	c.Max_procs = 800
	c.Min_tmp = uint64(20)
	c.Log_file = "/var/log/sleuth.log"
	return c
}

func Get_config(cf string) Conf {
	c := new_config()
	var e2 error
	fi, err := os.Open(cf)
	if err != nil {
		fmt.Println(err)
                fmt.Println("Using default settings.")
		return c
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return c
		}
		if line != "" && line[0] != '#' && strings.Contains(line, "=") {
			if line[len(line)-1] == '\n' {
				line = line[:len(line)-1]
			}
			switch parts := strings.Split(line, "="); parts[0] {
			case "max_load":
				c.Max_load, e2 = strconv.ParseFloat(parts[1], 64)
			case "max_mem":
				c.Max_mem, e2 = strconv.ParseFloat(parts[1], 64)
			case "max_procs":
				c.Max_procs, e2 = strconv.Atoi(parts[1])
			case "min_tmp":
				c.Min_tmp, e2 = strconv.ParseUint(parts[1], 10, 64)
			case "log_file":
				c.Log_file = parts[1]
			default:
				fmt.Println("Warning line not recognized: ", line)
			}
			if e2 != nil {
				fmt.Println("Warning: ", e2, ". Line: ", line)
				e2 = nil
			}
		}
		if err == io.EOF {
			break
		}
	}

	return c
}
