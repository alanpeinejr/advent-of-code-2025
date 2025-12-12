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
  fmt.Println(sumPathsFromTo("you", "out", parseInput(readInput()), 0, 20))
  
  //part 2
  start := time.Now()
  fmt.Println(sumPathsFromTo("dac", "out", parseInput(readInput()), 0, 12)) //7314
  fmt.Println(sumPathsFromTo("svr", "fft", parseInput(readInput()), 0, 12)) //7920
  fmt.Println(sumPathsFromTo("fft", "dac", parseInput(readInput()), 0, 20)) //6662060
  fmt.Println(6662060 * 7314 * 7920)
  fmt.Println("Staged took:", time.Since(start))
}

type Device struct {
  Name   string
  Output []*Device
  Input  []*Device
}
type State struct {
  Device *Device
  Path   []string
}

func sumPathsFromTo(start, end string, devices map[string]*Device, depth, maxDepth int) int {
  //cut the traversal, seems to be similarly deep path for each start to end segment
  if depth >= maxDepth {
    return 0
  }
  totalPaths := 0
  device := devices[start]
  for _, output := range device.Output {
    if output.Name == end {
      //checked, out is only ever the sole output
      return 1
    } else {
      sums := sumPathsFromTo(output.Name, end, devices, depth+1, maxDepth)
      totalPaths += sums
    }
  }
  return totalPaths
}

func parseInput(input string) map[string]*Device {
  lines := strings.Split(input, "\n")
  devices := map[string]*Device{}
  for _, line := range lines {
    deviceOutputs := strings.Split(line, ": ")
    deviceName := deviceOutputs[0]
    outputs := strings.Fields(deviceOutputs[1])
    //if device doesnt exist, create it
    //add outputs to device and mark their inputs from this device, if they dont exist, create them, then do that.
    var device *Device
    device, exists := devices[deviceName]
    if !exists {
      device = &Device{Name: deviceName, Output: make([]*Device, 0), Input: make([]*Device, 0)}
      devices[deviceName] = device
    }
    for _, output := range outputs {
      outDevice, outPutExists := devices[output]
      if outPutExists {
        outDevice.Input = append(devices[output].Input, device)
      } else {
        outDevice = &Device{Name: output, Output: make([]*Device, 0), Input: []*Device{device}}
        devices[output] = outDevice
      }
      device.Output = append(device.Output, outDevice)
      
    }
  }
  
  return devices
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
