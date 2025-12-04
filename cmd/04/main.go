package main

import (
	helper "aoc25/internal"
	"log/slog"
)

func checkSurroundingSquares(data []string, irow int, ipos int) int {
	count := 0
	indexes := [][]int{
		{irow - 1, ipos - 1},
		{irow - 1, ipos},
		{irow - 1, ipos + 1},
		{irow, ipos - 1},
		{irow, ipos + 1},
		{irow + 1, ipos - 1},
		{irow + 1, ipos},
		{irow + 1, ipos + 1},
	}

	for _, each := range indexes {
		row := each[0]
		pos := each[1]
		if row < 0 || row >= len(data) {
			continue
		}
		if pos < 0 || pos >= len(data[0]) {
			continue
		}
		if string(data[row][pos]) == "@" {
			count += 1
		}
	}

	return count
}

func movable(data []string, irow int, ipos int) bool {
	count := checkSurroundingSquares(data, irow, ipos)
	if count < 4 {
		return true
	}
	return false
}

func calculateRemovableRolls(data []string) (int, [][]int) {
	moveableRollCount := 0
	moveableRollIndexes := [][]int{}
	for irow, row := range data {
		for ipos, position := range row {
			if string(position) == "@" {
				if movable(data, irow, ipos) {
					moveableRollCount += 1
					moveableRollIndexes = append(moveableRollIndexes, []int{irow, ipos})
				}
			}
		}
	}
	return moveableRollCount, moveableRollIndexes
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func removeRolls(data []string, removables [][]int) {
	for index, _ := range removables {
		irow := removables[index][0]
		ipos := removables[index][1]
		row := data[irow]
		data[irow] = replaceAtIndex(row, '.', ipos)
	}
}

func p1(data []string) int {
	count, _ := calculateRemovableRolls(data)
	return count
}

func p2(data []string) int {
	count, removables := calculateRemovableRolls(data)

	for {
		removeRolls(data, removables)
		oldcount := count
		newCount, newRemovables := calculateRemovableRolls(data)
		count += newCount
		if count == oldcount {
			break
		}
		removables = newRemovables
	}
	return count
}

func main() {
	data := helper.ParseInputRemoveNewline(helper.LoadInput("input.txt"))
	slog.Info("AoC Day 4", "Forklifts P1", p1(data))
	slog.Info("AoC Day 4", "Forklifts P2", p2(data))
}
