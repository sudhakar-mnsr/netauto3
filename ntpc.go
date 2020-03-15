package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main {
   var host string
   flag.StringVar(&host, "e", "us.pool.ntp.org:123", "NTP host")
   flag.Parse() 

   // req datagram is a 48-byte long slice
   // that is used fo rsending time request to the server
   req := make([]byte, 48)

   req[0] = 0x1B

   // response 48 byte long slice for incoming datagram
   // with time values from the server
   rsp := make([]byte, 48)

   raddr, err := net.ResolveUDPAddr("udp", host)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   conn, err := net.DialUDP("udp", nil, raddr)
   if err != nil {
      fmt.Printf("failed to connec: %v\n", err)
      os.Exit(1)
   }
