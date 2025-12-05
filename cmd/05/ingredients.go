package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type ingredients struct {
	fresh_ingredients_ranges []string
}

func newIngredients(data []string) *ingredients {
	return &ingredients{fresh_ingredients_ranges: data}
}

func (self *ingredients) isFresh(ingredient int) bool {
	for _, eachRange := range self.fresh_ingredients_ranges {
		parts := strings.Split(eachRange, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		if ingredient >= start {
			if ingredient <= end {
				return true
			}
		}
	}
	return false
}

func sortRanges(ranges []string) {
	sort.Slice(ranges, func(i, j int) bool {
		partsI := strings.Split(ranges[i], "-")
		partsJ := strings.Split(ranges[j], "-")

		leftI, _ := strconv.Atoi(partsI[0])
		leftJ, _ := strconv.Atoi(partsJ[0])

		if leftI != leftJ {
			return leftI < leftJ
		}

		// tie-break using right-hand value
		rightI, _ := strconv.Atoi(partsI[1])
		rightJ, _ := strconv.Atoi(partsJ[1])

		return rightI < rightJ
	})
}

// Build ranges by ignoring overlaps
// Sort
// Add first range
// If second range is within first range, add second
func (self *ingredients) total() int {
	count := 0
	sortRanges(self.fresh_ingredients_ranges)
	ranges := []string{}
	start := 0
	end := 0
	for _, eachRange := range self.fresh_ingredients_ranges {
		parts := strings.Split(eachRange, "-")
		nextStart, _ := strconv.Atoi(parts[0])
		nextEnd, _ := strconv.Atoi(parts[1])
		if start == 0 {
			start = nextStart
			end = nextEnd
			continue
		}
		if nextStart <= end {
			// The range continues
			if nextEnd > end {
				end = nextEnd
			}
		}
		if nextStart > end {
			// The range is complete
			ranges = append(ranges, fmt.Sprintf("%d-%d", start, end))
			start = nextStart
			end = nextEnd
		}
	}
	ranges = append(ranges, fmt.Sprintf("%d-%d", start, end))
	for _, eachRange := range ranges {
		parts := strings.Split(eachRange, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		count += end - start + 1
	}
	return count
}
