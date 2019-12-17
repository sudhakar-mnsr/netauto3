package main

import (
   "fmt"
   "io/ioutil"
   "os"
   "strings"
)

func main() {
counts := make(map[string]int)
for _, filename := range os.Args
