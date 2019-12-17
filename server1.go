// Server1 is a minimal echo server
package main

import (
   "fmt"
   "log"
   "net/http"
)

func main() {
   http.HandleFunc("/", handler) // each request calls handler
   log.Fatal(http.ListenAndServe(localhost:8000", nil))
}


