package main

import (
	"log"
	"log/slog"
	"strconv"

	helper "aoc25/internal"
)

func wrapNumbers(number int) (int, int) {
	var loops = 0
	for number < 0 {
		number += 100
		loops += 1
	}
	for number > 99 {
		number -= 100
		loops += 1
	}
	return number, loops
}

func countHardZeros(data []string) int {
	var numberOfZeros = 0
	var safeValue = 50
	for index, each := range data {
		if each == "" {
			continue
		}
		slog.Debug("Instruction", "index", index, slog.Int("safe value", safeValue))
		slog.Debug("Instruction", "index", index, slog.String("instruction", each))
		var direction = string(each[0])
		var number, _ = strconv.Atoi(each[1:])
		switch direction {
		case "L":
			safeValue -= number
		case "R":
			safeValue += number
		default:
			log.Fatal("Invalid Direction")
		}
		safeValue, _ = wrapNumbers(safeValue)
		if safeValue == 0 {
			numberOfZeros += 1
		}
		slog.Debug("Instruction", "index", index, slog.Int("safe value", safeValue))
		slog.Debug("Instruction", "index", index, slog.Int("number of zeros", numberOfZeros))
		slog.Debug("")
	}
	return numberOfZeros
}

func countSoftZeros(data []string) int {
	var numberOfZeros = 0
	var safeValue = 50
	for index, each := range data {
		if each == "" {
			continue
		}
		slog.Debug("Instruction", "index", index, slog.Int("safe value", safeValue))
		slog.Debug("Instruction", "index", index, slog.String("instruction", each))
		var direction = string(each[0])
		var number, _ = strconv.Atoi(each[1:])
		if direction == "L" {
			// If the safe starts at 0 and we're going left, remove 1 to make up for an off by 1 error
			// The wrapNumbers() method will add 1 to the loops count, but going from 0 to a negative number
			// does not pass through 0!
			if safeValue == 0 {
				numberOfZeros -= 1
			}
			safeValue -= number
			value, loops := wrapNumbers(safeValue)
			safeValue = value
			numberOfZeros += loops
			// Additionally if the value lands on 0, add an additional 1
			// This is not required in the right direction logic as if you pass from 99 -> 0,
			// it will be caught in the wrapNumbers() method.
			if safeValue == 0 {
				numberOfZeros += 1
			}
		} else if direction == "R" {
			safeValue += number
			value, loops := wrapNumbers(safeValue)
			safeValue = value
			numberOfZeros += loops
		} else {
			log.Fatal("Invalid Input")
		}
		slog.Debug("Instruction", "index", index, slog.Int("safe value", safeValue))
		slog.Debug("Instruction", "index", index, slog.Int("number of zeros", numberOfZeros))
		slog.Debug("")
	}
	return numberOfZeros
}

func main() {
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	data := helper.ParseInput(helper.LoadInput("input.txt"))
	slog.Info("AoC Day 1", "Number of hard zeros", countHardZeros(data))
	slog.Info("Aoc Day 1", "Number of soft zeros", countSoftZeros(data))
}
