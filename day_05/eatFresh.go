package main

import (
  "fmt"
  "os"
  "slices"
  "strconv"
  "strings"
)

func main() {
  //part 1
  ranges, ingredients := parseInput(readInput())
  fmt.Println(findTotalfreshIngredients(ranges, ingredients))
  //part 2
  //sorting ranges for debugging
  slices.SortFunc(ranges, func(a, b Range) int {
    return a.Min - b.Min
  })
  fmt.Println(sumRanges(reduceRanges(ranges)))
  
}

func sumRanges(ranges []Range) int {
  total := 0
  for _, r := range ranges {
    fmt.Println(r.Min, r.Max)
    total += (r.Max - r.Min) + 1
  }
  return total
}

func reduceRanges(ranges []Range) []Range {
  //ranges in data overlap so lets reduce them into a new list of ranges that doesnt overlap
  reducedRanges := []Range{}
  for _, original := range ranges {
    added := false
    for i, reducedRange := range reducedRanges {
      if (original.Min >= reducedRange.Min && original.Min <= reducedRange.Max) || // min is within
        (original.Max >= reducedRange.Min && original.Max <= reducedRange.Max) || // max is within
        (original.Min <= reducedRange.Min && original.Max >= reducedRange.Max) { // reduced is within
        //overlap
        newMin := reducedRange.Min
        newMax := reducedRange.Max
        if original.Min < reducedRange.Min {
          newMin = original.Min
        }
        if original.Max > reducedRange.Max {
          newMax = original.Max
        }
        reducedRanges[i] = Range{newMin, newMax}
        added = true
        break
      }
    }
    if !added {
      reducedRanges = append(reducedRanges, original)
    }
  }
  return reducedRanges
  
}

func findTotalfreshIngredients(ranges []Range, ingredients []Ingredient) int {
  total := 0
  for _, ingredient := range ingredients {
    if ingredient.isFresh(ranges) {
      total++
    }
  }
  return total
}

func (a Ingredient) isFresh(ranges []Range) bool {
  for _, r := range ranges {
    if int(a) >= r.Min && int(a) <= r.Max {
      return true
    }
  }
  return false
}

type Range struct {
  Min int
  Max int
}
type Ingredient int

func parseInput(input string) (ranges []Range, ingredients []Ingredient) {
  rangeIngredients := strings.Split(input, "\n\n")
  rangeLines := strings.Split(rangeIngredients[0], "\n")
  ingredientLines := strings.Split(rangeIngredients[1], "\n")
  
  for _, row := range rangeLines {
    minMax := strings.Split(row, "-")
    ranges = append(ranges, Range{stringToInt(minMax[0]), stringToInt(minMax[1])})
  }
  for _, row := range ingredientLines {
    ingredients = append(ingredients, Ingredient(stringToInt(row)))
  }
  return ranges, ingredients
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
