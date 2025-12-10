package main

import (
	helper "aoc25/internal"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

type machine struct {
	indicator_light_numbers  []int
	button_wiring_schematics [][]int
	joltage_requirements     []int
}

func combinations[T any](items []T, length int) [][]T {
	var result [][]T
	var combo []T

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(combo) == length {
			c := make([]T, length)
			copy(c, combo)
			result = append(result, c)
			return
		}

		for i := start; i < len(items); i++ {
			combo = append(combo, items[i])
			backtrack(i + 1)
			combo = combo[:len(combo)-1]
		}
	}

	backtrack(0)
	return result
}

func uniqueAcrossAll(sliceOfSlices [][]int) []int {
	counts := map[int]int{}
	for _, slice := range sliceOfSlices {
		for _, value := range slice {
			counts[value]++
		}
	}

	var result []int
	for i, count := range counts {
		if count&1 == 1 {
			result = append(result, i)
		}
	}

	slices.Sort(result)
	return result
}

func p1(machine machine) int {
	minPresses := math.MaxInt
	optimal := machine.indicator_light_numbers
	schematics := machine.button_wiring_schematics
	for _, combo := range schematics {
		if slices.Equal(combo, optimal) {
			return 1
		}
	}
	comboNum := len(schematics)
	for i := 2; i < comboNum; i++ {
		combos := combinations(schematics, i)
		for _, combo := range combos {
			result := uniqueAcrossAll(combo)
			if slices.Equal(result, optimal) {
				comboLength := len(combo)
				if comboLength == 2 {
					return 2
				}
				if comboLength < minPresses {
					minPresses = comboLength
				}
			}
		}
	}
	return minPresses
}

func main() {
	start := time.Now()
	data := strings.Split(strings.TrimSuffix(helper.LoadInput("input.txt"), "\n"), "\n")
	machines := parse(data)
	fewest := 0
	for _, machine := range machines {
		fewest += p1(machine)
	}
	elapsed := time.Since(start)
	fmt.Println(fewest)
	fmt.Println(elapsed)
}

func parse(input []string) []machine {
	machines := []machine{}
	for _, line := range input {
		split := strings.Split(line, " ")

		indicatorLightNumbers := []int{}
		for i, picto := range string(split[0])[1 : len(split[0])-1] {
			if string(picto) == "#" {
				indicatorLightNumbers = append(indicatorLightNumbers, i)
			}
		}

		buttonWiringSchematics := [][]int{}
		for _, presses := range split[1 : len(split)-1] {
			buttonWiringSchematics = append(buttonWiringSchematics, atoiSlice(strings.Split(presses[1:len(presses)-1], ",")))
		}

		joltageRequirements := atoiSlice(strings.Split(split[len(split)-1][1:len(split[len(split)-1])-1], ","))
		machines = append(machines, machine{
			indicatorLightNumbers,
			buttonWiringSchematics,
			joltageRequirements,
		})
	}
	return machines
}

func atoiSlice(slice []string) []int {
	var nums = []int{}
	for _, i := range slice {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}
