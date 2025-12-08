package main

import (
	helper "aoc25/internal"
	"fmt"
	"math"
	"slices"
	"sort"
)

func calculateDistance(junction1 *junction, junction2 *junction) float64 {
	dx := float64(junction1.x - junction2.x)
	dy := float64(junction1.y - junction2.y)
	dz := float64(junction1.z - junction2.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func findXLargestCircuits(circuitsList [][]*junction, numbersOfCircuits int) int {
	total := 1
	sort.Slice(circuitsList, func(i, j int) bool {
		return len(circuitsList[i]) > len(circuitsList[j])
	})
	for _, circuit := range circuitsList[0:numbersOfCircuits] {
		total *= len(circuit)
	}
	return total
}

func MergeJunctions(junctions [][]*junction, first, second *junction) [][]*junction {
	var firstIndex, secondIndex = -1, -1

	for i, junction := range junctions {
		for _, j := range junction {
			if j == first {
				firstIndex = i
			}
			if j == second {
				secondIndex = i
			}
		}
	}

	// no merge needed (same group)
	if firstIndex == secondIndex {
		return junctions
	}

	// ensure firstIndex < secondIndex to delete safely
	if firstIndex > secondIndex {
		firstIndex, secondIndex = secondIndex, firstIndex
	}

	// merge second group into first
	junctions[firstIndex] = append(junctions[firstIndex], junctions[secondIndex]...)

	// delete second group
	junctions = append(junctions[:secondIndex], junctions[secondIndex+1:]...)

	first.Join(second)
	second.Join(first)

	return junctions
}

func p1(data [][]int, iterations int, finalMult int) int {
	junctionSlice := newJunctionSlice(data)
	for range iterations - 1 {
		closestDistance := math.MaxFloat64
		var junctionA *junction
		var junctionB *junction
		for i := 0; i < len(junctionSlice); i++ {
			for _, first := range junctionSlice[i] {
				for j := 0; j < len(junctionSlice); j++ {
					for _, second := range junctionSlice[j] {
						if first == second {
							continue
						}
						if slices.Contains(first.partners, second) {
							continue
						}
						distance := calculateDistance(first, second)
						if distance < closestDistance {
							closestDistance = distance
							junctionA = first
							junctionB = second
						}
					}
				}
			}
		}
		for _, junc := range junctionSlice {
			fmt.Println(junc)
		}
		junctionSlice = MergeJunctions(junctionSlice, junctionA, junctionB)
		fmt.Println()
		for _, junc := range junctionSlice {
			fmt.Println(junc)
		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}

	fmt.Println(junctionSlice)

	return findXLargestCircuits(junctionSlice, finalMult)
}

func main() {
	// data := helper.P8Parse(helper.LoadInput("input.txt"))
	fmt.Println(p1(helper.P8Parse(helper.LoadInput("mockinput.txt")), 10, 3))
	// fmt.Println(p1(helper.P8Parse(helper.LoadInput("input.txt")), 1000, 3))
}
