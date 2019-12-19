package popcount_test

import (
   "testing"
   "gopl.io/ch2/popcount"
)

func BitCount(x uint64) int {
   x = x - ((x >> 1) & 0x5555555555555555)
   x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
   x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
   x = x + (x >> 8)
   x = x + (x >> 16)
   x = x + (x >> 32)
   return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
n := 0
for x != 0 {
x = x & (x - 1) // clear rightmost non-zero bit
n++
}
return n
}

func PopCountByShifting(x uint64) int {
   n := 0
   for i := uint(0); i < 64; i++ {
      if x&(1<<i) != 0 {
         n++
      }
   }
   return n
} 
