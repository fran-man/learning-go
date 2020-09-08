package main

import ("fmt")

type Rectangle struct {
  width, height float64
}

func main() {
  rect := &Rectangle{10, 11}
  fmt.Println(area(rect))
}

func area(r *Rectangle) float64 {
  return r.width*r.height
}
