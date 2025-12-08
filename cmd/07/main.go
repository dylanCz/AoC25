package main

import (
	helper "aoc25/internal"
	"fmt"
	"slices"
)

func p1(data []string) int {
	var currentBeams []int
	var nextBeams []int
	splits := 0
	for irow, row := range data {
		for icol, col := range row {
			stringCol := string(col)
			if irow == 0 {
				if stringCol == "S" {
					nextBeams = append(currentBeams, icol)
				}
			}
			if stringCol == "^" {
				if slices.Contains(currentBeams, icol) {
					splits += 1
					ileft := icol - 1
					iright := icol + 1
					if ileft >= 0 {
						nextBeams = append(nextBeams, ileft)
					}
					if iright <= len(row) {
						nextBeams = append(nextBeams, iright)
					}
				}
			}
			if stringCol == "." {
				if slices.Contains(currentBeams, icol) {
					nextBeams = append(nextBeams, icol)
				}
			}
		}

		currentBeams = nextBeams
		nextBeams = []int{}
	}
	return splits
}

func p2(data []string) int {
	// 1 = 2
	// 2 = 4
	// 3 = 8
	var currentBeams []int
	var nextBeams []int
	timelines := 0
	for irow, row := range data {
		for icol, col := range row {
			stringCol := string(col)
			if irow == 0 {
				if stringCol == "S" {
					nextBeams = append(currentBeams, icol)
				}
			}
			if stringCol == "^" {
				if slices.Contains(currentBeams, icol) {

					ileft := icol - 1
					iright := icol + 1
					if ileft >= 0 {
						timelines += 1
						nextBeams = append(nextBeams, ileft)
					}
					if iright <= len(row) {
						timelines += 1
						nextBeams = append(nextBeams, iright)
					}
				}
			}
			if stringCol == "." {
				if slices.Contains(currentBeams, icol) {
					nextBeams = append(nextBeams, icol)
				}
			}
		}

		currentBeams = nextBeams
		nextBeams = []int{}
	}
	return timelines
}

func main() {
	data := helper.ParseInput(helper.LoadInput("mockinput.txt"))
	fmt.Println(p1(data))
	fmt.Println(p2(data))
}
