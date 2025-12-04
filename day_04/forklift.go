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
  part1, _ := optimizeForkLifts(parseInput(readInput()))
  fmt.Println(part1)
  //part 2
  fmt.Println(removeAllThatCanBeRemoved(parseInput(readInput())))
  
}

func optimizeForkLifts(locations [][]Location) (accessablePallets int, removablePallets []Location) {
  for _, row := range locations {
    for _, location := range row {
      if location.Char == '@' {
        pallets := countPallets(location, locations)
        if pallets < 4 {
          accessablePallets++
          removablePallets = append(removablePallets, location)
        }
      }
    }
  }
  return accessablePallets, removablePallets
}

func removePallets(removableLocations []Location, locations [][]Location) [][]Location {
  for _, location := range removableLocations {
    locations[location.Y][location.X].Char = '.'
  }
  return locations
}

func removeAllThatCanBeRemoved(locations [][]Location) int {
  totalRemoved := 0
  for {
    removed, removableLocations := optimizeForkLifts(locations)
    if removed == 0 {
      return totalRemoved
    }
    locations = removePallets(removableLocations, locations)
    totalRemoved += removed
  }
}

func countPallets(location Location, locations [][]Location) int {
  count := 0
  neighbors := Position{location.X, location.Y}.getNeighbors(locations)
  for _, neighborPos := range neighbors {
    if locations[neighborPos.Y][neighborPos.X].Char == '@' {
      count++
    }
  }
  return count
}

type (
  Location struct {
    Char rune
    X    int
    Y    int
  }
  Position struct {
    X int
    Y int
  }
)

func (a Position) Add(b Position) Position {
  return Position{a.X + b.X, a.Y + b.Y}
}

//hmm this isn't quite what I was trying to set up but should work
func (a Position) getNeighbors(locations [][]Location) []Position {
  test := creatFilterFunction(locations)
  return slices.Collect(test(a))
}

func creatFilterFunction(locations [][]Location) func(Position) func(yield func(Position) bool) {
  return func(position Position) func(yield func(Position) bool) {
    return func(yield func(Position) bool) {
      for _, direction := range ALL_DIRECTIONS {
        neighborPos := position.Add(direction)
        if neighborPos.Y >= 0 && neighborPos.Y < len(locations) && neighborPos.X >= 0 && neighborPos.X < len(locations[0]) {
          yield(neighborPos)
        }
      }
    }
  }
}

var LEFT = Position{-1, 0}
var RIGHT = Position{1, 0}
var UP = Position{0, -1}
var DOWN = Position{0, 1}
var LEFT_UP = LEFT.Add(UP)
var LEFT_DOWN = LEFT.Add(DOWN)
var RIGHT_UP = RIGHT.Add(UP)
var RIGHT_DOWN = RIGHT.Add(DOWN)
var ALL_DIRECTIONS = []Position{LEFT, RIGHT, UP, DOWN, LEFT_UP, LEFT_DOWN, RIGHT_UP, RIGHT_DOWN}

func parseInput(input string) [][]Location {
  rows := strings.Split(input, "\n")
  locations := make([][]Location, len(rows))
  for y, row := range rows {
    locationRow := make([]Location, len(row))
    for x, char := range row {
      locationRow[x] = Location{char, x, y}
    }
    locations[y] = locationRow
  }
  return locations
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
