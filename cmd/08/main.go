package main

import (
	helper "aoc25/internal"
	"fmt"
	"math"
	"sort"
	"time"
)

func calculateDistance(point1 *point, point2 *point) float64 {
	dx := float64(point1.x - point2.x)
	dy := float64(point1.y - point2.y)
	dz := float64(point1.z - point2.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func findXLargestCircuits(circuitsList [][]*point, x int) int {
	total := 1
	sort.Slice(circuitsList, func(i, j int) bool {
		return len(circuitsList[i]) > len(circuitsList[j])
	})
	for _, circuit := range circuitsList[0:x] {
		total *= len(circuit)
	}
	return total
}

func MergePoints(points [][]*point, first, second *point) [][]*point {
	var firstIndex, secondIndex = -1, -1

	first.Join(second)
	second.Join(first)

	for i, point := range points {
		for _, j := range point {
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
		return points
	}

	// ensure firstIndex < secondIndex to delete safely
	if firstIndex > secondIndex {
		firstIndex, secondIndex = secondIndex, firstIndex
	}

	// merge second group into first
	points[firstIndex] = append(points[firstIndex], points[secondIndex]...)

	// delete second group
	points = append(points[:secondIndex], points[secondIndex+1:]...)

	return points
}

func puzzle(data [][]int, iterations int, finalMult int) {
	pointSlice := newPointSlice(data)
	for x := range iterations {
		closestDistance := math.MaxFloat64
		var pointA *point
		var pointB *point
		for i := 0; i < len(pointSlice); i++ {
			for _, first := range pointSlice[i] {
				for j := i; j < len(pointSlice); j++ {
					for _, second := range pointSlice[j] {
						if first == second {
							continue
						}
						if _, exists := first.partners[second]; exists {
							continue
						}
						distance := calculateDistance(first, second)
						if distance < closestDistance {
							closestDistance = distance
							pointA = first
							pointB = second
						}
					}
				}
			}
		}
		pointSlice = MergePoints(pointSlice, pointA, pointB)
		if x == 1000 {
			fmt.Println(findXLargestCircuits(pointSlice, finalMult))
		}
		if len(pointSlice) == 1 {
			fmt.Println(pointA.x * pointB.x)
			break
		}
	}
}

func main() {
	data := helper.P8Parse(helper.LoadInput("input.txt"))

	start := time.Now()
	puzzle(data, 5000, 3)
	fmt.Println(time.Since(start))
}
