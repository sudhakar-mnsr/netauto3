package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// This program is a simple Network Time Protocol server over
// Unix Domain Socket instead of UDP. The implementation uses
// ListenUnixgram and Unixconn to manage requests.
// The server returns the number of seconds since 1900 upto the
// current time.

func main) {
   var path string
   flag.StringVar(&path, "e", "/tmp/time.sock", "NTP server socket endpoint")
   flag.Parse()

   // Creates a UnixAddr address
   addr, err := net.ResolveUnixAddr("unixgram", path)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   // setup connection UnixConn with ListenUnixgram
   conn, err := net.ListenUnixgram("unixgram", addr)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   defer conn.Close()
   fmt.Printf("listening on (unixgram) %s\n", conn.LocalAddr())

   // request/response loop
   
