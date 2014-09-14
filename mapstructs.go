// emulate the hash of hash type pattern
// using a map of structs

package main

import (
  "fmt"
)

type Vertex struct {
  X int
  Y int
}

// should call this hash of hashes  or map of structs foo

func main() {

  arc := make(map[string]Vertex)
  p := Vertex{X:1, Y:2}
  fmt.Printf("%s\n", p)

  arc["foo"] = p
  fmt.Printf("foo is: %s\n", arc["foo"])

  p.X = 99
  p.Y = 223

  arc["bar"] = p
  fmt.Printf("bar is: %s\n", arc["bar"])
  fmt.Printf("foo is: %s\n", arc["foo"])
  
  fmt.Printf("El Fin\n")
}
