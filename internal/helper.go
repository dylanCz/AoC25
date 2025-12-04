package helper

import (
	"os"
	"strings"
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
