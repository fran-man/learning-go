package main

import (
  "fmt"
  "hash/crc32"
  "io"
  "os"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Please enter exactly 1 argument")
    return
  }
  fileToHash := os.Args[1]
  fmt.Println(fileToHash)
  file, err := os.Open(fileToHash)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()

  hash := crc32.NewIEEE()

  _, copyErr := io.Copy(hash,file)
  if copyErr != nil {
    fmt.Println(copyErr)
    return
  }
  fmt.Println(hash.Sum32())
}
