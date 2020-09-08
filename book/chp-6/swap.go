package main

import "fmt"

func main() {
  a := 10
  b := 15
  x,y := swap(&a,&b)
  fmt.Printf("Post-swap:\nx=%d\ny=%d",*x,*y)
}

func swap(x,y *int) (*int,*int){
  fmt.Printf("Pre-swap:\nx=%d\ny=%d\n",*x,*y)
  xold := *x
  *x = *y
  *y = xold
  return x, y
}
