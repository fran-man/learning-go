package main

import (
  "encoding/json"
  "fmt"
  "net"
  "bufio"
  "os"
)

func main() {
  conn,err := net.Dial("tcp","localhost:1234")
  defer conn.Close()

  if err != nil {
    fmt.Println(err)
    return
  }

  rdr := bufio.NewReader(os.Stdin)
  encdr := json.NewEncoder(conn)

  for {
  fmt.Println("Please enter a message to send: ")
  msg,e := rdr.ReadString('\n')

  if e != nil {
    fmt.Println("Error reading from command line...",e)
    continue
  }

  err2 := encdr.Encode(msg)
  if err2 != nil {
    fmt.Println("Error sending over network",err2)
    continue
  }
}
}
