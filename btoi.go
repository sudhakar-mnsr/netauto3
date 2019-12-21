// btoi returns 1 if b is true and 0 if false
package main

import "fmt"

func btoi(b bool) int {
   if b {
      return 1
   }
   return 0
}
 
func main() {
   var b bool = true
   fmt.Println(btoi(b))
}
