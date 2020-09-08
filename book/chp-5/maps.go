package main

import "fmt"

func main() {
  family := map[string]string {
    "Mum": "Glenys",
    "Dad": "Kevin",
    "Brother": "Henry",
  }
  if bro, success := family["Brother"]; success {
    fmt.Println("My Brother is",bro)
  }
  if sis, success := family["Sister"]; success {
    fmt.Println("My Sister is",sis)
  }
}
