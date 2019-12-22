// reverse reverses a slice of ints in place
// rotate s left by 2 positions
package main

import "fmt"

func reverse(s []int) {
   for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
      s[i], s[j] = s[j], s[i]
   }
}

func main() {
   a := [...]int{0,1,2,3,4,5}
   reverse(a[:2])
   reverse(a[2:])
   reverse(a[:])
   fmt.Println(a)
   fmt.Printf("The elements of array %t\n", a)
}
