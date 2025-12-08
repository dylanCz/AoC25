package main

type point struct {
	x, y, z int
}

type edge struct {
	a, b point
	dist float64
}
