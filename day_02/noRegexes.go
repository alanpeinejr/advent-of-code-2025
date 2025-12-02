package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  fmt.Println(sumDiplicatesInRanges(parseInput(readInput()), false))
  fmt.Println(sumDiplicatesInRanges(parseInput(readInput()), true))
  
}

func sumDiplicatesInRanges(ranges []Range, part2 bool) int {
  total := 0
  for _, pair := range ranges {
    duplicates := findDuplicateSequences(pair, part2)
    for _, v := range duplicates {
      total += v
    }
  }
  return total
}

func findDuplicateSequences(pair Range, part2 bool) []int {
  duplicates := []int{}
  for i := pair.Min; i <= pair.Max; i++ {
    stringVersion := strconv.Itoa(i)
    if isDuplicateSequence(stringVersion, part2) {
      duplicates = append(duplicates, i)
    }
  }
  return duplicates
}

func isDuplicateSequence(s string, part2 bool) bool {
  if !part2 {
    return s[:len(s)/2] == s[len(s)/2:]
    
  } else {
    return isDuplicateSequence2(s)
  }
}

func isDuplicateSequence2(s string) bool {
  for groupSize := 1; groupSize <= len(s)/2; groupSize++ {
    //how many times does substring occur in s multiplied by how many times it can occur within the string, is that the same length?
    //ie, does a group of 3 happen 3 times in a string of length 9, or group of 5 happen 2 times in a string of length 10
    if strings.Count(s, s[:groupSize])*groupSize == len(s) {
      return true
    }
  }
  return false
}

type Range struct {
  Min int
  Max int
}

func parseInput(input string) []Range {
  pairs := strings.Split(input, ",")
  ranges := make([]Range, len(pairs))
  for i, value := range pairs {
    bounds := strings.Split(value, "-")
    ranges[i] = Range{stringToInt(bounds[0]), stringToInt(bounds[1])}
  }
  return ranges
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
