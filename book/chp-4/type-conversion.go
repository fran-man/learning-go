package main

import "fmt"

func main() {
  var float float64 = 123.123
  var int int32 = 4
  fmt.Println(float/float64(int))
}
