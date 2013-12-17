// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Warning, this package is experimental.  This was intended to
// attempt to create a snapshot of the /proc filesystem.
// However it is currently too slow and collects more
// information than was intended.
package snapshot

import (
   "strconv"
   "fmt"
   "os"
   "io/ioutil"
)

type Link struct {
   Name string
   LinkedTo string
}

// type Pid int

type File struct {
   Name string
   Content []byte
}

type Dir struct {
   Parent string
   Name string 
   SubDirs []Dir
   Links []Link
   Files []File
}

func isLink(m os.FileMode) bool {
		return m&os.ModeSymlink != 0
}

var mypid = strconv.Itoa(os.Getpid())

func safe(s string) bool {
  switch s {
  case mypid:
    return false
  case "":
    return false
  case "1":
    return false
  case "pagemap":
    return false
  case "attr":
    return false
  case "task":
    return false
  case "clear_refs":
    return false
  case "mem":
    return false
  case "kmsg":
    return false
  case "event":
    return false
  default:
    return true
  }
  return true
}

func (d *Dir) Path() string {
   return fmt.Sprintf("%s/%s",d.Parent,d.Name)
}

func (d *Dir) Read() error {
   fileinfo, err := ioutil.ReadDir(d.Path())
   if err != nil {
      return err
   }
   /* err = os.Chdir(d.Path())
   if err != nil {
      return err
   } */
   for _, f := range fileinfo {
      fname := fmt.Sprintf("%s/%s",d.Path(),f.Name())
      if safe(f.Name()) {
      fmt.Printf("At: %v/%v\n",d.Path(),f.Name())
      if f.IsDir() {
        sub := Dir{Parent: d.Path(), Name: f.Name()}
        err := sub.Read()
        if err != nil { fmt.Println(err) }
        d.SubDirs = append(d.SubDirs, sub)
      } else if f.Mode().IsRegular() {
        content, err := ioutil.ReadFile(fname)
        if err != nil { fmt.Println(err) }
        file := File{f.Name(), content}
        d.Files = append(d.Files, file)
      } else if isLink(f.Mode()) {
        lto, err := os.Readlink(fname)
        if err != nil { fmt.Println(mypid, err) }
        link := Link{f.Name(),lto}
        d.Links = append(d.Links, link)
      }
      }
    }
    return nil
}

func Snap() (Dir, error) {
   proc := Dir{Parent: "", Name: "proc"}
   err := proc.Read()
   return proc, err
}
        

