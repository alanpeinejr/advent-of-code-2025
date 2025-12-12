package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "time"
)

func main() {
  //part 1
  start := time.Now()
  fmt.Println(sumFittingTrees(parseInput(readInput())))
  fmt.Println("Part 1 Took:", time.Since(start))
  //part 2
  start = time.Now()
  fmt.Println()
  fmt.Println("Part 2 Took:", time.Since(start))
}

type Tree struct {
  X     int
  Y     int
  Needs []int
}
type Present struct {
  Size int
}

func sumFittingTrees(trees []Tree, presents []Present) int {
  fittingTrees := 0
  fmt.Println(len(trees))
  for _, tree := range trees {
    if doTheyFitIfWeSmoosh(tree, presents) {
      fittingTrees++
    }
  }
  return fittingTrees
}

func doTheyFitIfWeSmoosh(tree Tree, presents []Present) bool {
  totalArea := tree.X * tree.Y
  areaNeeded := 0
  for i, need := range tree.Needs {
    areaNeeded += need * presents[i].Size
  }
  return areaNeeded <= totalArea
}

func parseInput(input string) (trees []Tree, presents []Present) {
  groups := strings.Split(input, "\n\n")
  presentShapes := groups[:6]
  treeLines := strings.Split(groups[6], "\n")
  //presents
  for _, presentGroup := range presentShapes {
    presents = append(presents, Present{Size: strings.Count(presentGroup, "#")})
  }
  
  //trees
  for _, line := range treeLines {
    areaNeedsArray := strings.Split(line, ": ")
    xY := strings.Split(areaNeedsArray[0], "x")
    tree := Tree{X: stringToInt(xY[0]), Y: stringToInt(xY[1]), Needs: make([]int, 6)}
    needsStrings := strings.Fields(areaNeedsArray[1])
    for i, needString := range needsStrings {
      tree.Needs[i] = stringToInt(needString)
    }
    trees = append(trees, tree)
  }
  
  return trees, presents
}

func stringToInt(this string) int {
  value, _ := strconv.Atoi(this)
  return value
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
