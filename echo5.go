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
      concatenated := fmt.Sprintf("%d" + "----" + "%s", temp, arg)
      fmt.Println(concatenated)
   }
   fmt.Println(s)
}
