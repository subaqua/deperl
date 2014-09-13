// do something I might do as a sysadm in perl in golang instead

package main

import (
  "os"
  "fmt"
  "log"
  "strings"
)

func main() {

  const pwpath = "/etc/passwd"

  // like perl's file stat operations (-e,-d,-s, etc)
  /// Stat returns a FileInfo interface, Stat is definatly a friend

  finfo, err := os.Stat(pwpath)
  if err != nil {
    fmt.Printf("%s does not exist")
    log.Fatal(err)
  }

  if finfo.IsDir() {
    fmt.Printf("%s appears to be a directory, need a file\n", pwpath)
    os.Exit(1)
  } else {
    fmt.Printf("%s is a file of %d bytes, good\n", pwpath, finfo.Size())
  }
  
  pwfile, err := os.Open(pwpath)
  if err != nil {
    log.Fatal(err)
  }
  defer pwfile.Close()
  
  fdata := make([]byte, finfo.Size())
  count, err := pwfile.Read(fdata)

  if err != nil {
    log.Fatal(err)
  }
  
  fmt.Printf("I read %d/%d bytes from %s, thy data is:\n[\n%q\n]\n",
    count, finfo.Size(), pwpath, fdata)

  sdata := string(fdata[:count-1])
  fmt.Printf("[%s]\n", sdata)

  // now for the final coup de gra, put this into a usable hash table and select something from it

  ary := strings.Split(sdata, "\n")
  zusers := make(map[string]string)
  var fields []string
  
  for i := range ary {
    fields = strings.Split(ary[i], ":")
    // first, let's prove we can extract two fields from each line effectively
    if len(fields) > 2 {
      fmt.Printf("%03d: [%21s:%5s]->[%s]\n", i, fields[0], fields[2], ary[i])
      // and shove the data into our hashtable
      zusers[fields[0]] = fields[len(fields)-1]
    }
  }

  // and read the hashtable
  
  fmt.Printf("Shell of the devil is: [%s]\n", zusers["daemon"])
  fmt.Printf("El fin\n")
}

