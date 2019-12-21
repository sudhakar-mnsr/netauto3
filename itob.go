// btoi returns 1 if b is true and 0 if false
package main

import "fmt"

func itob(b int) bool {
   return b != 0
}
 
func main() {
   var b int = 10
   fmt.Println(itob(b))
}
