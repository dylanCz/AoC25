package main

import (
	helper "aoc25/internal"
	"fmt"
)

func matrixTranspose(matrix [][]string) [][]string {

	rows := len(matrix)
	cols := len(matrix[0])

	// Prepare transposed matrix
	result := make([][]string, cols)
	for i := range result {
		result[i] = make([]string, rows)
	}

	// Transpose
	for r := range rows {
		for c := range cols {
			result[c][r] = matrix[r][c]
		}
	}

	return result
}

func p1(data [][]string) int {
	matrix := matrixTranspose(data)

	worksheets := newWorksheets(matrix)
	var sum int
	for _, sheet := range worksheets {
		sum += sheet.performCalculation()
	}
	return sum
}

func p2(data [][]string) int {
	matrix := matrixTranspose(data)
	worksheets := newWorksheets(matrix)
	// var sum int
	for _, sheet := range worksheets {
		sheet.performCephalopodCalculation()
	}
	// return sum
	return 0
}

func main() {
	data := helper.P6Parse(helper.LoadInput("mockinput.txt"))
	fmt.Println(p1(data))
	fmt.Println(p2(data))
}
