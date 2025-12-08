package main

import (
  "fmt"
  "math"
  "os"
  "slices"
  "strconv"
  "strings"
)

func main() {
  //part 1
  fmt.Println(multiplyLargestCircuitLengths(connectXJunctions(parseInput(readInput()), 1000)))
  //part 2
  fmt.Println(connectJunctionsUntilOne(parseInput(readInput())))
  
}

func connect(a, b Junction, circuits []Circuit) []Circuit {
  //see if a or b are already in a circuit
  var circuitAIndex, circuitBIndex = -1, -1
  for i, circuit := range circuits {
    if _, exists := circuit.Members[a]; exists {
      circuitAIndex = i
    }
    if _, exists := circuit.Members[b]; exists {
      circuitBIndex = i
    }
  }
  switch {
  case circuitAIndex == -1 && circuitBIndex == -1:
    //neither in a circuit, create new circuit
    circuits = append(circuits, Circuit{Members: map[Junction]bool{a: true, b: true}})
  case circuitAIndex == -1 && circuitBIndex != -1:
    //only b in a circuit, add a to b's circuit
    circuits[circuitBIndex].Members[a] = true
  case circuitAIndex != -1 && circuitBIndex == -1:
    //only a in a circuit, add b to a's circuit
    circuits[circuitAIndex].Members[b] = true
  case circuitBIndex == circuitAIndex:
    //do nothing, same circuit
  default:
    //both in a circuit, merge circuits, remove old one
    for member, _ := range circuits[circuitBIndex].Members {
      circuits[circuitAIndex].Members[member] = true
    }
    //remove circuitBIndex
    circuits = append(circuits[:circuitBIndex], circuits[circuitBIndex+1:]...)
  }
  
  return circuits
}

type Junction struct {
  X int
  Y int
  Z int
}
type Connection struct {
  From     Junction
  To       Junction
  Distance int
}

type Circuit struct {
  Members map[Junction]bool
}

func connectXJunctions(junctions []Junction, x int) []Circuit {
  allConnections := findAllPossibleConnections(junctions)
  circuits := []Circuit{}
  for i := 0; i < x && i < len(allConnections); i++ {
    circuits = connect(allConnections[i].From, allConnections[i].To, circuits)
  }
  return circuits
}

func connectJunctionsUntilOne(junctions []Junction) int {
  allConnections := findAllPossibleConnections(junctions)
  circuits := []Circuit{}
  i := -1
  for !(len(circuits) == 1 && len(circuits[0].Members) == 1000) {
    i++
    circuits = connect(allConnections[i].From, allConnections[i].To, circuits)
  }
  fmt.Println(allConnections[i])
  return allConnections[i].To.X * allConnections[i].From.X
}

func multiplyLargestCircuitLengths(circuits []Circuit) int {
  a, b, c := 0, 0, 0
  for _, circuit := range circuits {
    length := len(circuit.Members)
    if length > a {
      c = b
      b = a
      a = length
    } else if length > b {
      c = b
      b = length
    } else if length > c {
      c = length
    }
  }
  return a * b * c
}

func findAllPossibleConnections(junctions []Junction) []Connection {
  connections := map[Connection]int{}
  allDistances := findAllDistances(junctions)
  for from, distances := range allDistances {
    for to, distance := range distances {
      if from != to {
        a, b := connectionKeyOrderer(from, to)
        connection := Connection{From: a, To: b, Distance: distance}
        connections[connection] = distance
      }
    }
  }
  
  //order the connections by distance
  connectionsSlice := mapToSlice(connections)
  slices.SortFunc(connectionsSlice, func(a, b Connection) int {
    return a.Distance - b.Distance
  })
  
  return connectionsSlice
}

func mapToSlice(connections map[Connection]int) []Connection {
  connectionSlice := make([]Connection, 0)
  for connection, _ := range connections {
    connectionSlice = append(connectionSlice, connection)
  }
  
  return connectionSlice
}

func connectionKeyOrderer(a Junction, b Junction) (Junction, Junction) {
  smallerJunction := minJunction(a, b)
  if smallerJunction == a {
    return a, b
  } else {
    return b, a
  }
  
}
func minJunction(a, b Junction) Junction {
  if a.X < b.X {
    return a
  } else if a.X > b.X {
    return b
  } else {
    if a.Y < b.Y {
      return a
    } else if a.Y > b.Y {
      return b
    } else {
      if a.Z < b.Z {
        return a
      } else {
        return b
      }
    }
  }
}
func findAllDistances(junctions []Junction) map[Junction]map[Junction]int {
  allDistances := make(map[Junction]map[Junction]int)
  for _, junction := range junctions {
    allDistances[junction] = junction.findDistances(junctions)
  }
  return allDistances
}

func (a Junction) findDistances(junctions []Junction) map[Junction]int {
  distances := make(map[Junction]int)
  for _, junction := range junctions {
    distances[junction] = a.distance(junction)
  }
  return distances
}

func (a Junction) distance(b Junction) int {
  dx := math.Pow(float64(b.X-a.X), 2)
  dy := math.Pow(float64(b.Y-a.Y), 2)
  dz := math.Pow(float64(b.Z-a.Z), 2)
  return int(math.Sqrt(dx + dy + dz))
}

func parseInput(input string) []Junction {
  lines := strings.Split(input, "\n")
  junctions := make([]Junction, len(lines))
  for i, line := range lines {
    coords := strings.Split(line, ",")
    junctions[i] = Junction{stringToInt(coords[0]), stringToInt(coords[1]), stringToInt(coords[2])}
  }
  return junctions
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
