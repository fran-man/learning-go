package main

import "fmt"

func main() {
  mySlice := make([]byte, 5, 20)
  fmt.Println(mySlice)

  for i,_ := range mySlice {
    for j,_ := range mySlice {
      mySlice[j]++
    }
    fmt.Printf("Index %d: %d\n",i,mySlice[i])
  }

  fmt.Println(mySlice)
}
