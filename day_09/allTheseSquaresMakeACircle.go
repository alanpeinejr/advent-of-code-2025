package main

import (
  "fmt"
  "image"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  rectangles, _ := parseInput(readInput())
  fmt.Println(findLargestArea(rectangles))
  //part 2
  fmt.Println(findLargestContainedArea(parseInput(readInput())))
}

func findLargestContainedArea(rectangles []image.Rectangle, lines []image.Rectangle) int {
  maxArea := 0
  for _, rect := range rectangles {
    rect.Max = rect.Max.Add(image.Point{1, 1})
    area := rect.Dx() * rect.Dy()
    canBeConsidered := true
    for _, line := range lines {
      line.Max = line.Max.Add(image.Point{1, 1})
      //if a line overlaps the rectangle inset by 1, it is not contained within the polygon
      if line.Overlaps(rect.Inset(1)) {
        canBeConsidered = false
      }
    }
    
    if canBeConsidered {
      maxArea = max(maxArea, area)
    }
  }
  return maxArea
}

func findLargestArea(rectangles []image.Rectangle) int {
  maxArea := 0
  for _, rect := range rectangles {
    //add 1 for inclusive
    rect.Max = rect.Max.Add(image.Point{1, 1})
    area := rect.Dx() * rect.Dy()
    maxArea = max(maxArea, area)
  }
  return maxArea
}

func parseInput(input string) ([]image.Rectangle, []image.Rectangle) {
  stringLines := strings.Split(input, "\n")
  points, rectangles, lines := []image.Point{}, []image.Rectangle{}, []image.Rectangle{}
  for y, line := range stringLines {
    xY := strings.Split(line, ",")
    point := image.Point{X: stringToInt(xY[0]), Y: stringToInt(xY[1])}
    for _, next := range points {
      rectangles = append(rectangles, image.Rectangle{point, next}.Canon()) //test?
    }
    points = append(points, point)
    if len((points)) > 1 {
      lines = append(lines, image.Rectangle{points[y-1], points[y]}.Canon())
    }
  }
  lines = append(lines, image.Rectangle{points[0], points[len(points)-1]}.Canon())
  
  return rectangles, lines
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
