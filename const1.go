package main

import "fmt"

const (
   a = 1 * iota
   b
   c = 2 * iota
   d
   e
   f = 10 * iota
   g
   h
   i = 11
   j
   k
)

func main() {
   fmt.Println(a, b, c, d, e, f, g, h, i, j, k)
}
