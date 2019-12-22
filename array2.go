package main

import "fmt"

func main() {
   var q [3]int = [3]int {1,2,3}
   var r [3]int = [3]int{1,2}
   fmt.Printf("%T\n", q)
   fmt.Printf("%T\n", r)
   
   // [...] appear then array length is the number of initializers
   s := [...]int{1,2,3}
   fmt.Printf("%T\n", s)
   
   symbol := [...]string{0: "$", 1: "#", 2: "@"}
   fmt.Println(symbol[1])
}
