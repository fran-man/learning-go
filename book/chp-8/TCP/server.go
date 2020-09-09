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
  conn,e := listener.Accept()
  defer conn.Close()
  for {
    if e != nil {
      fmt.Println("There was an error:",e)
      continue
    }
    var inMsg string
    e2 := json.NewDecoder(conn).Decode(&inMsg)
    if e2 == nil {
      fmt.Printf("Received Message %s\n",inMsg)
    }
  }
}
