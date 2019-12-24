package main

import (
   "fmt"
   "os"
   "html/html"
)

func main() {
   doc, err := html.Parse(os.Stdin)
   if err != nil {
      fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
      os.Exit(1)
   }
   for _, link := range visit(nil, doc) {
      fmt.Println(link)
   }
}
