package main

import (
	helper "aoc25/internal"
	"log/slog"
	"time"
)

func concatenateAll(nums ...int) int {
	var result int = 0

	for _, n := range nums {
		pow := int(10)
		for n >= pow {
			pow *= 10
		}
		result = result*pow + n
	}

	return result
}

func p1(data []string) int {
	var voltage int
	for _, bank := range data {
		if bank == "" {
			continue
		}
		first := 0
		second := 0
		lastDigit := len(bank) - 1
		for i := 0; i < len(bank); i++ {
			battery := int(bank[i] - '0')
			if i != lastDigit && battery > first {
				first = battery
				second = 0
			} else if battery > second {
				second = battery
			}
		}
		voltage += concatenateAll(first, second)
	}
	return voltage
}

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

func p2(data []string) int {
	var totalVoltage int
	for _, bank := range data {
		if bank == "" {
			continue
		}
		digits := 12
		bankVoltage := make([]int, 0, digits)
		nextIndex := 0
		for index := range digits {
			voltage, index := findLargest(bank[nextIndex : len(bank)-(digits-(index+1))])
			nextIndex += index + 1
			bankVoltage = append(bankVoltage, voltage)
		}
		totalVoltage += int(digitsToNumber(bankVoltage))
	}
	return totalVoltage
}

func main() {
	data := helper.ParseInput(helper.LoadInput("input.txt"))

	start := time.Now()
	slog.Info("AoC Day 3", "Battery Voltages P1", p1(data))
	p1elapsed := time.Since(start)
	slog.Info("P1", "time", p1elapsed)

	start = time.Now()
	slog.Info("AoC Day 3", "Battery Voltages P2", p2(data))
	p2elapsed := time.Since(start)
	slog.Info("P2", "time", p2elapsed)
}
