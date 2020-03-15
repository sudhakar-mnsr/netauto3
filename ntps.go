package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
   var host string
   flag.StringVar(&host, "e", ":1123", "server address")
   flag.Parse()

   // Create a UDP host address
   addr, err := net.ResolveUDPAddr("udp", host)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   // setup connection UDPConn with ListenUDP
   conn, err := net.ListenUDP("udp", addr)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   defer conn.Close()
