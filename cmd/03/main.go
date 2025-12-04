package main

import (
	helper "aoc25/internal"
	"log/slog"
	"time"
)

func findLargest(data string) (int, int) {
	largest := 0
	index := 0
	for i := 0; i < len(data); i++ {
		num := int(data[i] - '0')
		if num > largest {
			largest = num
			index = i
			if largest == 9 {
				break
			}
		}
	}
	return largest, index
}

func digitsToNumber(nums []int) int {
	number := 0
	for _, digits := range nums {
		number = number*10 + digits
	}
	return number
}

func puzzle(data []string, digits int) int {
	var totalVoltage int
	for _, bank := range data {
		if bank == "" {
			continue
		}
		bankVoltage := make([]int, digits)
		nextIndex := 0
		for i := range digits {
			voltage, index := findLargest(bank[nextIndex : len(bank)-(digits-(i+1))])
			nextIndex += index + 1
			bankVoltage[i] = voltage
		}

		totalVoltage += digitsToNumber(bankVoltage)
	}
	return totalVoltage
}

func main() {
	data := helper.ParseInput(helper.LoadInput("input.txt"))

	start := time.Now()
	slog.Info("AoC Day 3", "Battery Voltages P1", puzzle(data, 2))
	elapsedv1 := time.Since(start)
	slog.Info("P1", "time", elapsedv1)

	start = time.Now()
	slog.Info("AoC Day 3", "Battery Voltages P2", puzzle(data, 12))
	elapsedv2 := time.Since(start)
	slog.Info("P2", "time", elapsedv2)
}
