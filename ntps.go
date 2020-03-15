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

   fmt.Printf("listening for time requests: (udp) %s\n", conn.LocalAddr())

   // From this point, the remainder of the code simply
   // reads the incoming request and send a response
   // read incoming requests.
   // since we are using a sessionless proto, each request can 
   // potentially go to a different client. Therefore, the ReadFormXXX
   // operation returns the remote address (saved in raddr)
   // where to send the response.
   _, raddr, err := conn.ReadFromUDP(make([]byte, 48))
   if err != nil {
      fmt.Println("error getting request:", err)
      os.Exit(1)
   }
   // ensure laddr is set
   if raddr == nil {
      fmt.Println("request missing remote addr")
      os.Exit(1)
   }

   secs, fracs := getNTPSeconds(time.Now())

   // response packet is filled with the seconds and 
   // fractional sec values using Big-Endian
   rsp := make([]byte, 48)
   // write seconds (as uint32) in buffer at [40:43]
   binary.BigEndian.PutUint32(rsp[40:], uint32(secs))
   binary.BigEndian.PutUint32(rsp[44:], uint32(fracs))

   // send data
   if _, err := conn.WriteToUDP(rsp, raddr); err != nil {
      fmt.Println("err sending data:", err)
      os.Exit(1)
   }
}
