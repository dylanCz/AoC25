package main

import (
	helper "aoc25/internal"
	"fmt"
	"slices"
	"strings"
	"time"
)

func recurseFind(devices map[string][]string, device string, paths *int) {
	outDevices := devices[device]
	if slices.Contains(outDevices, "out") {
		*paths += 1
		return
	}
	for _, outDevice := range outDevices {
		recurseFind(devices, outDevice, paths)
	}
}

func p1(devices map[string][]string) int {
	paths := 0
	starting := devices["you"]
	for _, each := range starting {
		recurseFind(devices, each, &paths)
	}
	return paths
}

func main() {
	start := time.Now()
	data := strings.Split(strings.TrimSuffix(helper.LoadInput("input.txt"), "\n"), "\n")
	devices := parse(data)
	fmt.Println(p1(devices))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func parse(input []string) map[string][]string {
	devices := make(map[string][]string)
	for _, line := range input {
		outDevices := []string{}
		split := strings.Split(line, " ")
		for index := 1; index < len(split); index++ {
			outDevices = append(outDevices, split[index])
		}
		devices[split[0][:3]] = outDevices
	}
	return devices
}
