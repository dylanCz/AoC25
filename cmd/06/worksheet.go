package main

import (
	"math"
	"slices"
	"sort"
	"strconv"
)

type worksheet struct {
	operator string
	integers []int
}

func newWorksheets(data [][]string) []*worksheet {
	var w = []*worksheet{}
	for _, row := range data {
		sheet := &worksheet{}
		for icol, col := range row {
			if icol == len(row)-1 {
				sheet.operator = col
				continue
			}
			numCol, _ := strconv.Atoi(col)
			sheet.integers = append(sheet.integers, numCol)
		}
		w = append(w, sheet)
	}
	return w
}

func (self *worksheet) sortWorksheet() worksheet {
	sort.Ints(self.integers)
	slices.Reverse(self.integers)

	return *self
}

func (self *worksheet) performCalculation() int {
	var result int
	if self.operator == "+" {
		result = 0
		for _, num := range self.integers {
			result += num
		}
	}
	if self.operator == "*" {
		result = 1
		for _, num := range self.integers {
			result *= num
		}
	}
	return result
}

func (self *worksheet) performCephalopodCalculation() int {
	var result int
	self.sortWorksheet()
	maxDigits := int(1 + math.Floor(math.Log10(float64(self.integers[0]))))

	if self.operator == "+" {
		result = 0
		for _, num := range self.integers {
			slice := 0
			for i := maxDigits; i == 0; i-- {

			}
			result += num
		}
	}
	if self.operator == "*" {
		result = 1
		for _, num := range self.integers {
			result *= num
		}
	}
	return result
}
