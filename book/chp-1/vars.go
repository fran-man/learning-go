package main
import "fmt"

func main() {
  x := 123456789
  const y = 987654
  fmt.Println("Performing arithmetic on a var:", x - 100)
  x += 1
  fmt.Println("\nIncrementing var:", x)
  fmt.Println("\ny:", y)
  //y = 2

  printMultiVars()
}

func printMultiVars() {
  const (
    a = 1
    a2 = 2
    a3 = 3
  )
  fmt.Println("Printing multiple vars...")
  fmt.Println(a)
  fmt.Println(a2)
  fmt.Println(a3)
}
