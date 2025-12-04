package main

import (
	helper "aoc25/internal"
	"testing"
)

func BenchmarkD4P1(b *testing.B) {
	data := helper.ParseInputRemoveNewline(helper.LoadInput("input.txt"))
	helper.BenchWrapper(b, "D4P1", p1, data)
}

func BenchmarkD4P2(b *testing.B) {
	data := helper.ParseInputRemoveNewline(helper.LoadInput("input.txt"))
	helper.BenchWrapper(b, "D4P2", p2, data)
}
