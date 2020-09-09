package main

import (
  "encoding/json"
  "fmt"
  "net"
)

func main() {
  listener, e := net.Listen("tcp",":1234")

  if e != nil {
    fmt.Println("There was an error:",e)
  }

  for {
    conn,e := listener.Accept()
    if e != nil {
      fmt.Println("There was an error:",e)
      continue
    }
    var inMsg string
    e2 := json.NewDecoder(conn).Decode(&inMsg)
    if e2 == nil {
      fmt.Printf("Received Message %s\n",inMsg)
    }
    conn.Close()
  }
}
