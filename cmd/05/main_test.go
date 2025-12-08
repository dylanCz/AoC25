package main

import (
	helper "aoc25/internal"
	"testing"
)

func BenchmarkD5P1(b *testing.B) {
	p1(helper.LoadInput("input.txt"))
}

func BenchmarkD5P2(b *testing.B) {
	p2(helper.LoadInput("input.txt"))
}
