package main

import "fmt"

func main() {
  fmt.Println(square(17))
}

func square(x uint32) uint32 {
  fmt.Println("Start Func")
  defer fmt.Println("End Func")
  result := x*x
  fmt.Println("Mid Func")
  return result
}
