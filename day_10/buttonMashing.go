package main

import (
  "fmt"
  "os"
  "slices"
  "strconv"
  "strings"
  "github.com/draffensperger/golp"
)

func main() {
  //part 1
  fmt.Println(findButtonPressesForAllMachines(parseInput(readInput()), false))
  //part 2
  fmt.Println(findButtonPressesForAllMachines(parseInput(readInput()), true))
}

func findButtonPressesForAllMachines(machines []Machine, leverPulled bool) int {
  totalPresses := 0
  for _, machine := range machines {
    presses := 0
    if !leverPulled {
      presses = mashButtons(machine)
    } else {
      presses = useTheStupidLibrary(machine)
    }
    totalPresses += presses
    
  }
  return totalPresses
}

func maxJolt(slice []int) int {
  maxValue := 0
  for _, value := range slice {
    if value > maxValue {
      maxValue = value
    }
  }
  return maxValue
}

func useTheStupidLibrary(machine Machine) int {
  //button clicks can't exceed the largest jolt
  maxClicks := maxJolt(machine.Joltages)
  totalJoltages, totalButtons := len(machine.Joltages), len(machine.Buttons)
  
  lp := golp.NewLP(0, totalButtons)
  
  //for each button, set rules for bounds
  coefficients := make([]float64, totalButtons)
  for i := 0; i < totalButtons; i++ {
    coefficients[i] = 1.0
    lp.SetInt(i, true)                       //must be an integer
    lp.SetBounds(i, 0.0, float64(maxClicks)) //must be between zero and max
  }
  lp.SetObjFn(coefficients)
  
  //for each joltage, set the buttons that can alter it, and its target value
  for i := 0; i < totalJoltages; i++ {
    entries := make([]golp.Entry, 0)
    for j, button := range machine.Buttons {
      if slices.Contains(button, i) {
        entries = append(entries, golp.Entry{Col: j, Val: 1.0})
      }
    }
    targetValue := float64(machine.Joltages[i])
    lp.AddConstraintSparse(entries, golp.EQ, targetValue)
  }
  
  //magic library solution
  lp.Solve()
  
  clicks := 0
  clicksPerButton := lp.Variables()
  for _, clickCount := range clicksPerButton {
    clicks += int(clickCount)
  }
  return clicks
}

func mashButtons(machine Machine) int {
  queue := []State{}
  for i := range machine.Buttons {
    queue = append(queue, State{Machine: machine.clone(), Presses: 0, ButtonBeingPressed: i})
  }
  seen := map[string]bool{}
  for len(queue) > 0 {
    current := queue[0]
    queue = queue[1:]
    current.Machine.pressButton(current.ButtonBeingPressed)
    current.Presses++
    //if we've seen this state there's no need going forward, skip
    if _, exists := seen[string(current.Machine.Current)]; exists {
      continue
    } else {
      seen[string(current.Machine.Current)] = true
    }
    if current.Machine.isAtDesiredState() {
      return current.Presses
    }
    //loop buttons, add to queue
    for i := range current.Machine.Buttons {
      queue = append(queue, State{Machine: current.Machine.clone(), Presses: current.Presses, ButtonBeingPressed: i})
    }
  }
  return 0
}

type Machine struct {
  Goal     []rune
  Current  []rune
  Joltages []int
  Buttons  [][]int
}

type State struct {
  Machine            Machine
  Presses            int
  ButtonBeingPressed int
}

func (a Machine) clone() Machine {
  newButtons := make([][]int, len(a.Buttons))
  for i, button := range a.Buttons {
    newButton := make([]int, len(button))
    copy(newButton, button)
    newButtons[i] = newButton
  }
  newCurrent := make([]rune, len(a.Current))
  copy(newCurrent, a.Current)
  newJoltages := make([]int, len(a.Joltages))
  copy(newJoltages, a.Joltages)
  newGoal := make([]rune, len(a.Goal))
  copy(newGoal, a.Goal)
  return Machine{Goal: newGoal, Current: newCurrent, Joltages: newJoltages, Buttons: newButtons}
}

func removeSurrounder(input string) string {
  return input[1 : len(input)-1]
}

func (a *Machine) reset() {
  for i := range a.Goal {
    a.Current[i] = '.'
  }
}

func (a *Machine) pressButton(buttonIndex int) {
  button := a.Buttons[buttonIndex]
  for _, pos := range button {
    if a.Current[pos] == '.' {
      a.Current[pos] = '#'
    } else {
      a.Current[pos] = '.'
    }
  }
}

func (a Machine) isAtDesiredState() bool {
  for i, char := range a.Goal {
    if a.Current[i] != char {
      return false
    }
  }
  return true
}

func parseInput(input string) []Machine {
  lines := strings.Split(input, "\n")
  machines := []Machine{}
  for _, line := range lines {
    lightsButtonsJoltage := strings.Split(line, " ")
    lightsString := []rune(removeSurrounder(lightsButtonsJoltage[0]))
    
    joltageString := strings.Split(removeSurrounder(lightsButtonsJoltage[len(lightsButtonsJoltage)-1]), ",")
    joltages := make([]int, len(joltageString))
    for i, strJoltage := range joltageString {
      joltages[i] = stringToInt(strJoltage)
    }
    
    buttons := make([][]int, 0) //remove first and last
    for _, buttonGroup := range lightsButtonsJoltage[1 : len(lightsButtonsJoltage)-1] {
      stringButtons := strings.Split(removeSurrounder(buttonGroup), ",")
      buttonInts := make([]int, len(stringButtons))
      for i, strButton := range stringButtons {
        buttonInts[i] = stringToInt(strButton)
      }
      buttons = append(buttons, buttonInts)
    }
    
    machines = append(machines, Machine{Goal: lightsString, Joltages: joltages, Buttons: buttons, Current: make([]rune, len(lightsString))})
    machines[len(machines)-1].reset()
  }
  
  return machines
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
