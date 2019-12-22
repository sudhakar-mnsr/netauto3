package main

import "fmt"

func equal(x,y []int) bool {
   if len(x) != len(y) {
      return false
   }
   for i := range x {
      if x[i] != y[i] {
         return false
      }
   }
   return true
}
func main() {
p := [...]int{0, 1, 2, 3, 4, 5}
q := [...]int{0, 1, 2, 3, 4, 5}
s := []int{0, 1, 2, 3, 4, 5}
t := []int{0, 1, 2, 3, 4, 5}
u := []int{0, 10, 2, 3, 4, 5}
// can compare arrays directly
fmt.Println(p == q)
// cant compare slices compile error
// fmt.Println(s == t)
fmt.Println(equal(s,t))
fmt.Println(equal(s,u))
}
