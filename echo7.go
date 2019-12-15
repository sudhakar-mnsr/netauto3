// Echo2 prints its commandline arguments.
package main

import (
   "fmt"
   "strconv"
   "os"
)

func main() {
   for temp, arg := range os.Args {
      concatenated := strconv.Itoa(temp) + "---" + arg
      fmt.Println(concatenated)
   }
}
