package main

import (
  "fmt"
  "crypto/sha1"
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

  hash := sha1.New()

  _, copyErr := io.Copy(hash,file)
  if copyErr != nil {
    fmt.Println(copyErr)
    return
  }
  fmt.Printf("%x",hash.Sum([]byte{}))
}
