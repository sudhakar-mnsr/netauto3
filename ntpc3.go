package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// This program implements a simple NTP client over UDP.
// It uses the generic net.Dial function create connection.
// This makes the code more generic and easy to change to
// support other protocols.

func main() {
   var host string
   flag.StringVar(&host, "h", "us.pool.ntp.org:123", "NTP host")
   flag.Parse()

   // req data packet is a 48-byte long value
   // that is used for sending time request.
   req := make([]byte, 48)
   
   // req is initialized with 0x1B or 0001 1011 which is a
   // request setting for time server.
   req[0] = 0x1B

   // rsp byte slice used to recieve server response
   rsp := make([]byte, 48)

   // setup generic connection (net.Conn) using net.Dial
   conn, err := net.Dial("udp", host)
   if err != nil {
      fmt.Printf("failed to connect : %v\n", err)
      os.Exit(1)
   }


