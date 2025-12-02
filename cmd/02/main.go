package main

import (
	helper "aoc25/internal"
	"log"
	"log/slog"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

func allSameStrings(stringList []string) bool {
	for i := 1; i < len(stringList); i++ {
		if stringList[i] != stringList[0] {
			return false
		}
	}
	return true
}

func splitEven(stringNum string) [][]string {
	stringLength := len(stringNum)
	results := [][]string{}

	for splitSize := 1; splitSize < stringLength; splitSize++ {
		if stringLength%splitSize != 0 {
			continue // skip uneven splits
		}

		var parts []string
		for i := 0; i < stringLength; i += splitSize {
			segment := stringNum[i : i+splitSize]
			parts = append(parts, segment)
		}
		results = append(results, parts)
	}

	return results
}

func repeatingPatternChecker(num string) bool {
	doubled := num + num
	return strings.Index(doubled[1:], num)+1 != len(num)
}

func calculateDuplicates(num int) bool {
	if num == 0 {
		panic("i must be > 0")
	}

	digits := int(1 + math.Floor(math.Log10(float64(num))))
	switch digits {
	case 1:
		return false
	case 2:
		return num%11 == 0
	case 3:
		return num%111 == 0
	case 4:
		return num%101 == 0
	case 5:
		return num%11111 == 0
	case 6:
		return num%1001 == 0 || num%10101 == 0
	case 7:
		return num%1111111 == 0
	case 8:
		return num%1010101 == 0 || num%10001 == 0
	case 9:
		return num%1001001 == 0
	case 10:
		return num%101010101 == 0 || num%100001 == 0
	default:
		panic("unsupported digit length")
	}
}

func invalidIds(data []string) int {
	var sumOfInvalids = 0
	for _, each := range data {
		start, _ := strconv.Atoi(strings.Split(each, "-")[0])
		end, _ := strconv.Atoi(strings.Split(each, "-")[1])
		for num := start; num <= end; num++ {
			stringNum := strconv.Itoa(num)
			numLength := len(stringNum)
			if numLength%2 == 0 {
				if stringNum[0:numLength/2] == stringNum[numLength/2:] {
					sumOfInvalids += num
				}
			}
		}
	}
	return sumOfInvalids
}

func moreInvalidIds(data []string) int {
	var sumOfInvalids = 0
	for _, each := range data {
		start, _ := strconv.Atoi(strings.Split(each, "-")[0])
		end, _ := strconv.Atoi(strings.Split(each, "-")[1])
		for num := start; num <= end; num++ {
			stringNum := strconv.Itoa(num)
			if slices.ContainsFunc(splitEven(stringNum), allSameStrings) {
				sumOfInvalids += num
			}
			// if calculateDuplicates(num) {
			// 	sumOfInvalids += num
			// }
		}
	}
	return sumOfInvalids
}

func moreInvalidIdsMaths(data []string) int {
	var sumOfInvalids = 0
	for _, each := range data {
		start, _ := strconv.Atoi(strings.Split(each, "-")[0])
		end, _ := strconv.Atoi(strings.Split(each, "-")[1])
		for num := start; num <= end; num++ {
			if calculateDuplicates(num) {
				sumOfInvalids += num
			}
		}
	}
	return sumOfInvalids
}

func moreInvalidIdsRemixed(data []string) int {
	var sumOfInvalids = 0
	for _, each := range data {
		start, _ := strconv.Atoi(strings.Split(each, "-")[0])
		end, _ := strconv.Atoi(strings.Split(each, "-")[1])
		for num := start; num <= end; num++ {
			stringNum := strconv.Itoa(num)
			if repeatingPatternChecker(stringNum) {
				sumOfInvalids += num
			}
		}
	}
	return sumOfInvalids
}

func main() {
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	data := helper.ParseInputByCommas(helper.LoadInput("input.txt"))

	start := time.Now()
	slog.Info("AoC Day 2", "Sum of Invalids 1", invalidIds(data))
	p1elapsed := time.Since(start)

	start = time.Now()
	slog.Info("AoC Day 2", "Sum of Invalids 2", moreInvalidIds(data))
	p2elapsed := time.Since(start)

	start = time.Now()
	slog.Info("AoC Day 2", "Sum of Invalids 2 Remixed", moreInvalidIdsRemixed(data))
	p2remixelapsed := time.Since(start)

	start = time.Now()
	slog.Info("AoC Day 2", "Sum of Invalids 2 Maths", moreInvalidIdsMaths(data))
	p2mathselapsed := time.Since(start)

	log.Printf("P1  took %s", p1elapsed)
	log.Printf("P2  took %s", p2elapsed)
	log.Printf("P2R took %s", p2remixelapsed)
	log.Printf("P2M took %s", p2mathselapsed)
}
