package main

import (
	helper "aoc25/internal"
	"log/slog"
	"time"
)

func checkSurroundingSquares(data []string, irow int, ipos int) bool {
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

	if count < 4 {
		return true
	}
	return false
}

func puzzle(data []string) (int, [][]int) {
	moveableRollCount := 0
	moveableRollIndexes := [][]int{}
	for irow, row := range data {
		if row == "" {
			continue
		}
		for ipos, position := range row {
			if string(position) == "@" {
				if checkSurroundingSquares(data, irow, ipos) {
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

func main() {
	data := helper.ParseInputRemoveNewline(helper.LoadInput("input.txt"))
	start := time.Now()
	count, removables := puzzle(data)
	elapsedv1 := time.Since(start)
	slog.Info("AoC Day 4", "Forklifts P1", count)
	slog.Info("P1", "time", elapsedv1)

	elapsedv2 := time.Since(start)
	for {
		removeRolls(data, removables)
		oldcount := count
		moreCount, moreRemovables := puzzle(data)
		count += moreCount
		if count == oldcount {
			break
		}
		removables = moreRemovables
	}
	slog.Info("AoC Day 4", "Forklifts P2", count)
	slog.Info("P1", "time", elapsedv2)
}
