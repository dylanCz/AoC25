package main

import "slices"

type point struct {
	x        int
	y        int
	z        int
	partners []*point
}

func newPoint(data []int) *point {
	return &point{x: data[0], y: data[1], z: data[2]}
}

func newPointSlice(data [][]int) [][]*point {
	pointSlice := [][]*point{}
	for _, eachPoint := range data {
		pointSlice = append(pointSlice, []*point{newPoint(eachPoint)})
	}
	return pointSlice
}

func (self *point) Join(point *point) {
	if slices.Contains(self.partners, point) {
		return
	}
	self.partners = append(self.partners, point)
}
