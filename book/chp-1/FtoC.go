package main

import "fmt"

func main() {
  fmt.Print("Temp in Farenheit...")
  var input float32
  fmt.Scanf("%f", &input)

  fmt.Println("Temperature in C:", (input - 32)*(5.0/9.0))
}
