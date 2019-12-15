// Echo2 prints its commandline arguments.
package main

import (
   "fmt"
   "os"
)

func main() {
   // var temp int
   s, sep := "", ""
   for temp, arg := range os.Args {
      s += sep + arg
      sep = " "
      fmt.Print(temp)
      fmt.Println("---" + arg)
   }
   fmt.Println(s)
}
