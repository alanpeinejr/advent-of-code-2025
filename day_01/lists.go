package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  fmt.Println(countZeroHitsInTurns(99, 50, parseInput(readInput())))
}

func countZeroHitsInTurns(max int, startingPosition int, turns []Turn) (int, int) {
  dialStart := createDial(max, startingPosition)
  current := dialStart
  zeroHits := 0
  totalZeroClicks := 0
  for _, turn := range turns {
    var zeroClicks int
    current, zeroClicks = turnDial(current, turn)
    if current.Position == 0 {
      zeroHits++
    }
    totalZeroClicks += zeroClicks
  }
  return zeroHits, totalZeroClicks
}

func turnDial(dialStart *Node, turn Turn) (*Node, int) {
  current := dialStart
  zeroClicks := 0
  for i := 0; i < turn.Amount; i++ {
    if turn.Direction == LEFT {
      current = current.Left
    } else {
      current = current.Right
    }
    if current.Position == 0 {
      zeroClicks++
    }
  }
  return current, zeroClicks
}

func createDial(max int, startingPosition int) *Node {
  var startingNode *Node
  zero := &Node{Position: 0}
  current := zero
  for i := 1; i <= max; i++ {
    newNode := &Node{Position: i}
    current.Right = newNode
    newNode.Left = current
    current = newNode
    if i == startingPosition {
      startingNode = newNode
    }
  }
  zero.Left = current
  current.Right = zero
  
  return startingNode
}

type Node struct {
  Position int
  Left     *Node
  Right    *Node
}

type Turn struct {
  Direction int
  Amount    int
}

const (
  LEFT  = -1
  RIGHT = 1
)

func parseDirections(input string) int {
  switch input {
  case "L":
    return LEFT
  case "R":
    return RIGHT
  default:
    panic("Unknown direction: " + string(input))
  }
}

func parseInput(input string) []Turn {
  rows := strings.Split(input, "\n")
  directions := make([]Turn, len(rows))
  for i, value := range rows {
    directions[i] = Turn{parseDirections(value[0:1]), stringToInt(value[1:])}
  }
  return directions
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
