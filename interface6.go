package main

import (
	"fmt"
	"math"
)

type shape interface {
   area() float64
   perim() float64
}

type rect struct {
   name string
   length, height float64
}

func (r *rect) area() float64 {
   return r.length * r.height
}

func (r *rect) perim() float64 {
   return 2*r.length + 2*r.height
}

func (r *rect) String() string {
   return fmt.Sprintf("%s[length:%.2f height:%.2f]", r.name, r.length, r.height,)
}


