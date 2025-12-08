package helper

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func LoadInput(filename string) string {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(input)
}

func ParseInput(input string) []string {
	return strings.Split(input, "\n")
}

func ParseInputRemoveNewline(input string) []string {
	result := strings.Split(input, "\n")
	return result[:len(result)-1]
}

func ParseInputByCommas(input string) []string {
	split := strings.TrimSuffix(input, "\n")
	return strings.Split(split, ",")
}

func BenchWrapper[T any](b *testing.B, name string, f func(data T) int, data T) {
	b.Run(name, func(b *testing.B) {
		for b.Loop() {
			f(data)
		}
	})
}

func P5Parse(input string) ([]string, []string) {
	input = strings.TrimSuffix(input, "\n")
	result := strings.Split(input, "\n")
	for i, v := range result {
		if v == "" {
			return result[:i], result[i+1:]
		}
	}
	return result, nil
}

func P6Parse(input string) [][]string {
	var rows [][]string

	for line := range strings.SplitSeq(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)

		var nums []string
		for _, p := range parts {
			nums = append(nums, p)
		}

		rows = append(rows, nums)
	}
	return rows
}

func P8Parse(input string) [][]int {
	var coords [][]string
	removeNewline := strings.TrimSuffix(input, "\n")
	splitStrings := strings.SplitSeq(removeNewline, "\n")
	for line := range splitStrings {
		coords = append(coords, strings.Split(line, ","))
	}

	var result [][]int
	for _, row := range coords {
		nums := make([]int, len(row))
		for i, s := range row {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			nums[i] = n
		}
		result = append(result, nums)
	}
	return result
}
