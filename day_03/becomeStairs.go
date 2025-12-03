package main

import (
  "fmt"
  "math"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  fmt.Println(sumJolts(parseInput(readInput()), 2))
  //part 2
  fmt.Println(sumJolts(parseInput(readInput()), 12))
  
}
func sumJolts(banks []string, size int) int {
  total := 0
  for _, bank := range banks {
    total += findJolt(bank, size)
  }
  return total
}

func findJolt(bank string, place int) int {
  if place == 0 {
    return 0
  }
  max1, maxIndex := 0, 0
  //cant be the last ones
  for i := 0; i < len(bank)-(place-1); i++ {
    val := stringToInt(bank[i : i+1])
    if val > max1 {
      max1 = val
      maxIndex = i
    }
  }
  return max1*int(math.Pow10(place-1)) + findJolt(bank[maxIndex+1:], place-1)
}

func parseInput(input string) []string {
  return strings.Split(input, "\n")
}

func stringToInt(this string) int {
  value, _ := strconv.Atoi(this)
  return int(value)
}

func readInput() string {
  var filename string
  if len(os.Args) < 2 {
    fmt.Println("Assuming local file input.txt")
    filename = "./input.txt"
  } else {
    filename = os.Args[1]
  }
  
  data, err := os.ReadFile(filename)
  if err != nil {
    fmt.Println("Can't read file:", filename)
    panic(err)
  }
  
  //return and account for windows
  return strings.ReplaceAll(string(data), "\r\n", "\n")
}
