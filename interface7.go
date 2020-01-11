package main

import (
   "fmt"
   "math"
)

type shape interface {
   area() float64
}

type polygon interface {
   perim()
}

type curved interface {
   circonf()
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

type triangle struct {
   name string
   a, b, c float64
}

func (t *triangle) area() float64 {
   return 0.5 * (t.a * t.b)
}

func (t *triangle) perim() float64 {
   return t.a + t.b + math.Sqrt((t.a*t.a)+(t.b*t.b))
}

func (t *triangle) String() string {
   return fmt.Sprintf("%s[sides: a=%.2f b=%.2f c=%.2f]", t.name, t.a, t.b, t.c,)
}

