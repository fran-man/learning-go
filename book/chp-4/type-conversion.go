package main

import "fmt"

func main() {
  var myFloat float64 = 123.123
  var myInt int32 = 4
  fmt.Println(myFloat/float64(myInt))
}
