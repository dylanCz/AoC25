package main

import "slices"

type junction struct {
	x        int
	y        int
	z        int
	partners []*junction
}

func newJunction(data []int) *junction {
	return &junction{x: data[0], y: data[1], z: data[2]}
}

func newJunctionSlice(data [][]int) [][]*junction {
	junctionSlice := [][]*junction{}
	for _, eachJunction := range data {
		junctionSlice = append(junctionSlice, []*junction{newJunction(eachJunction)})
	}
	return junctionSlice
}

func (self *junction) Join(junction *junction) {
	if slices.Contains(self.partners, junction) {
		return
	}
	self.partners = append(self.partners, junction)
}
