package main

import "fmt"

func equal(x, y map[string]int) bool {
   if len(x) != len(y) {
      return false
   }
   for k, xv := range x {
      if yv, ok := y[k]; !ok || yv != xv {
         return false
      }
   }
   return true
} 
func main() {
   var ages1,ages2 map[string]int
   
   ages3 := make(map[string]int)
   ages3["alice"] = 31
   ages3["charlie"] = 34
   
   ages4 := map[string]int{
               "alice": 31,
               "charlie":34,
   }
   ages5 := map[string]int{
               "alice": 31,
               "charlie":35,
   }
   fmt.Println(equal(ages1,ages2))
   fmt.Println(equal(ages3,ages4))
   fmt.Println(equal(ages3,ages5))
}

