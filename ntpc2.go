package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// This program implements an NTP client over Unix Domain Socket
// Datagram. The -e flag is used to specify the socket path.

func main() {
   var path string
   flag.StringVar(&path, "e", "/tmp/time.sock", "NTP client sock endpoint")
   flage.Parse()

   // req data packet is 48 byte long value
   // that is used for sending time request.
   req := make([]byte, 48)

   // req is initialized with 0x1B or 0001 1011 which is 
   // a request setting for time server.
   req[0] = 0x1B

   
