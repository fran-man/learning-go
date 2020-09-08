package main

import ("fmt")

type Rectangle struct {
  width, height float64
}

func main() {
  rect := &Rectangle{10, 11}
  fmt.Println(rect.area())
  rect.doubleSize()
  fmt.Println(rect.area())
  rect.doubleSize()
  fmt.Println(rect.area())
}

func (r *Rectangle) area() float64 {
  return r.width*r.height
}

func (r *Rectangle) doubleSize() {
  r.width = r.width*2
  r.height = r.height*2
}
