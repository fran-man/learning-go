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
}
