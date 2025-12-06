package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  fmt.Println(sumAllProblemResults(parseInput(readInput())))
  //part 2
  fmt.Println(sumAllProblemResults(parseInput2(readInput())))
  
}
func sumAllProblemResults(problems []Problm) int {
  total := 0
  for _, problem := range problems {
    total += problem.solve()
  }
  return total
}

func (a Problm) solve() int {
  if a.Operation == "+" {
    total := 0
    for _, val := range a.operands {
      total += val
    }
    return total
  } else {
    total := 1
    for _, val := range a.operands {
      total *= val
    }
    return total
  }
}

type Problm struct {
  operands  []int
  Operation string
}

func parseInput(input string) (problems []Problm) {
  lines := strings.Split(input, "\n")
  for y, line := range lines {
    
    numbers := strings.Fields(line)
    for x, number := range numbers {
      if y == 0 {
        problem := Problm{}
        problem.operands = make([]int, 4)
        problems = append(problems, problem)
      }
      if y < 4 {
        problems[x].operands[y] = stringToInt(number)
      } else {
        problems[x].Operation = number
      }
    }
  }
  
  return problems
}

func parseInput2(input string) (problems []Problm) {
  lines := strings.Split(input, "\n")
  length := len(lines[0])
  problem := Problm{}
  problem.operands = make([]int, 0)
  
  for i := length - 1; i >= 0; i-- {
    cut := strings.TrimSpace(buildStringFromIndex(lines, i))
    if len(cut) == 0 {
      continue
    }
    if cut[len(cut)-1] == '+' || cut[len(cut)-1] == '*' {
      problem.Operation = cut[len(cut)-1:]
      problem.operands = append(problem.operands, stringToInt(strings.TrimSpace(cut[:len(cut)-1])))
      problems = append(problems, problem)
      problem = Problm{}
      problem.operands = make([]int, 0)
    } else {
      problem.operands = append(problem.operands, stringToInt(cut))
    }
  }
  return problems
}

func buildStringFromIndex(lines []string, index int) string {
  result := ""
  for _, line := range lines {
    result += string(line[index])
  }
  return result
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
