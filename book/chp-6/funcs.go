package main

import "fmt"
import "sort"

func main() {
  name := "eweueurywyooghgjgjjgjgjgjgjjaaaaakjahakhk"
  fmt.Println(string(
    sortRuneSlice(
      stringToChars(
        name))))
}

func sortRuneSlice(runes []rune) []rune {
  sort.Slice(runes, func(i,j int) bool {
    return runes[i] < runes[j]
  })
  return runes
}

func stringToChars(stringToSplit string) []rune {
  charArray := make([]rune, len(stringToSplit))
  for i, char := range stringToSplit {
    charArray[i] = char
  }
  return charArray
}
