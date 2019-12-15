// Echo2 prints its commandline arguments.
package main

import (
   "fmt"
   "time"
   "strings"
   "os"
)

func main() {
   start := time.Now()
   fmt.Println(strings.Join(os.Args[1:], " "))
   secs := time.Since(start).Seconds()
   fmt.Println(secs)
   start1 := time.Now()
   var s1, sep1 string
   for i := 1; i < len(os.Args); i++ {
      s1 += sep1 + os.Args[i]
      sep1 = " "
   }
   fmt.Println(s1)
   secs1 := time.Since(start1).Seconds()
   fmt.Println(secs1)
}
