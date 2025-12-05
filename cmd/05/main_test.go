package main

import (
	helper "aoc25/internal"
	"os"
	"testing"
)

func BenchmarkD4P1(b *testing.B) {
	puzzle(helper.LoadInput(os.Getenv("input_file")))
}
