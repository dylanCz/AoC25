package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

var parent map[point]point
var size map[point]int

func find(p point) point {
	for parent[p] != p {
		parent[p] = parent[parent[p]]
		p = parent[p]
	}
	return p
}

func union(a, b point) {
	ra := find(a)
	rb := find(b)
	if ra == rb {
		return
	}
	// union by size
	if size[ra] < size[rb] {
		ra, rb = rb, ra
	}
	parent[rb] = ra
	size[ra] += size[rb]
}

func atoi(s []byte) (n int) {
	sign := 1
	for i, c := range s {
		if i == 0 && c == '-' {
			sign = -1
			continue
		}
		n = 10*n + int(c-'0')
	}
	return sign * n
}

func processEdges(points []point) []edge {
	var edges []edge
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]
			dx := float64(a.x - b.x)
			dy := float64(a.y - b.y)
			dz := float64(a.z - b.z)
			dist := math.Sqrt(dx*dx + dy*dy + dz*dz)
			edges = append(edges, edge{a, b, dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})
	return edges
}

func puzzle(pointlist []point) {
	edges := processEdges(pointlist)

	parent = make(map[point]point, len(pointlist))
	size = make(map[point]int, len(pointlist))

	for _, p := range pointlist {
		parent[p] = p
		size[p] = 1
	}

	for i, e := range edges {
		ra := find(e.a)
		rb := find(e.b)

		if ra != rb {
			union(ra, rb)
		}

		if i+1 == 1000 {
			var sizes []int
			seen := make(map[point]bool)
			for _, p := range pointlist {
				r := find(p)
				if !seen[r] {
					seen[r] = true
					sizes = append(sizes, size[r])
				}
			}

			sort.Ints(sizes)
			if len(sizes) >= 3 {
				n := sizes[len(sizes)-3:]
				fmt.Println(n[0] * n[1] * n[2])
			}
		}

		if allConnected() {
			fmt.Println(e.a.x * e.b.x)
			break
		}
	}
}

func allConnected() bool {
	unique := make(map[point]bool)
	for p := range parent {
		unique[find(p)] = true
	}
	return len(unique) == 1
}

func main() {
	start := time.Now()
	pointlist := make([]point, 0, 1000)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		b := bytes.Split([]byte(line), []byte(","))

		pointlist = append(pointlist, point{
			x: atoi(b[0]),
			y: atoi(b[1]),
			z: atoi(b[2]),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	puzzle(pointlist)
	fmt.Println(time.Since(start))
}
