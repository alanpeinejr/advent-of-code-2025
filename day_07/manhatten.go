package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  start, positions, splitters := parseInput(readInput())
  visitedMap := fireBeams(start, positions, splitters)
  //part 1
  fmt.Println(countVisitedSplitters(splitters))
  //part 2
  fmt.Println(countBottomPositions(visitedMap))
  
}
func countBottomPositions(visitiedMap map[Position]int) int {
  //first find max y
  maxY := 0
  for pos, _ := range visitiedMap {
    if pos.Y > maxY {
      maxY = pos.Y
    }
  }
  count := 0
  for pos, visits := range visitiedMap {
    if pos.Y == maxY-1 {
      count += visits
    }
  }
  return count
}

func printCounts(visitiedMap map[Position]int) {
  
}

func countVisitedSplitters(splitters map[Position]bool) int {
  count := 0
  for _, visited := range splitters {
    if visited {
      count++
    }
  }
  return count
}

func fireBeams(start Position, positions [][]Position, splitters map[Position]bool) map[Position]int {
  beams := []Position{start}
  visitedPositions := map[Position]int{}
  visitedPositions[start] = 1
  for len(beams) > 0 {
    current := beams[0]
    beams = beams[1:]
    nextBeams := current.move(splitters)
    for _, nextBeam := range nextBeams {
      if nextBeam.Y >= len(positions) || nextBeam.X < 0 || nextBeam.X >= len(positions[0]) {
        continue //out of bounds
      }
      if _, visited := visitedPositions[nextBeam]; !visited {
        visitedPositions[nextBeam] = visitedPositions[current]
        beams = append(beams, nextBeam)
      } else {
        visitedPositions[nextBeam] += visitedPositions[current]
      }
    }
  }
  return visitedPositions
}

func (beam Position) move(splitters map[Position]bool) []Position {
  nextPosition := beam.positionBelow()
  if _, exists := splitters[nextPosition]; exists {
    splitters[nextPosition] = true
    return []Position{{nextPosition.X - 1, nextPosition.Y}, {nextPosition.X + 1, nextPosition.Y}}
  }
  return []Position{nextPosition}
}

type Splitter struct {
  Position Position
  visited  bool
}
type Position struct {
  X int
  Y int
}

func (i Position) positionBelow() Position {
  return Position{i.X, i.Y + 1}
}

func parseInput(input string) (Position, [][]Position, map[Position]bool) {
  lines := strings.Split(input, "\n")
  positions := make([][]Position, len(lines))
  splitters := map[Position]bool{}
  start := Position{}
  for y, line := range lines {
    row := make([]Position, len(line))
    for x, char := range line {
      row[x] = Position{x, y}
      if char == '^' {
        splitters[row[x]] = false
      }
      if char == 'S' {
        start = row[x]
      }
    }
    positions[y] = row
  }
  
  return start, positions, splitters
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
