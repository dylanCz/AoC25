package main

type point struct {
	x        int
	y        int
	z        int
	partners map[*point]struct{}
}

func newPoint(data []int) *point {
	return &point{
		x:        data[0],
		y:        data[1],
		z:        data[2],
		partners: make(map[*point]struct{}),
	}
}

func newPointSlice(data [][]int) [][]*point {
	pointSlice := [][]*point{}
	for _, eachPoint := range data {
		pointSlice = append(pointSlice, []*point{newPoint(eachPoint)})
	}
	return pointSlice
}

func (self *point) Join(point *point) {
	self.partners[point] = struct{}{}
}
