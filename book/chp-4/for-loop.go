package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 {
      fmt.Printf("%d is a multiple of 3!\n", i)
    }
  }
}
