package main

import (
   "sort"
   "fmt"
)

func main() {
   ages := make(map[string]int)
   ages["alice"] = 31
   ages["charlie"] = 34
   
   for name, age := range ages {
      fmt.Printf("%s\t%d\n", name, age)
   }
   var names []string
   for name1 := range ages {
      names = append(names, name1)
   }
   sort.Strings(names)
   for _, name := range names {
      fmt.Printf("%s\t%d\n", name, ages[name])
   }
}
