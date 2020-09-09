package main

import (
  "encoding/json"
  "fmt"
  "net"
  "os"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Please enter a message to send!")
    return
  }
  conn,err := net.Dial("tcp","localhost:1234")
  defer conn.Close()
  if err != nil {
    fmt.Println(err)
    return
  }

  message := os.Args[1]
  err2 := json.NewEncoder(conn).Encode(message)
  if err2 != nil {
    fmt.Println(err2)
    return
  }
}
