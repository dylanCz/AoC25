package main

import (
	helper "aoc25/internal"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func computeVolume(leftX, leftY, rightX, rightY int) int {
	left := abs(leftX-rightX) + 1
	right := abs(leftY-rightY) + 1
	return left * right
}

func atoiCoords(coord string) (int, int) {
	parts := strings.Split(coord, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}

func puzzle(data []string, grid [][]byte) {
	maxVolume := 0
	rectangles := []rectangle{}
	for i := range data {
		leftX, leftY := atoiCoords(data[i])
		for j := i + 1; j < len(data); j++ {
			rightX, rightY := atoiCoords(data[j])
			volume := computeVolume(leftX, leftY, rightX, rightY)
			rectangles = append(rectangles, rectangle{leftX, rightX, leftY, rightY, volume})
			if volume > maxVolume {
				maxVolume = volume
			}
		}
	}
	fmt.Println(maxVolume)
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].volume > rectangles[j].volume
	})

	for _, rectangle := range rectangles {
		if (grid[rectangle.y1][rectangle.x2] == '#' || grid[rectangle.y1][rectangle.x2] == 'X') &&
			(grid[rectangle.y2][rectangle.x1] == '#' || grid[rectangle.y2][rectangle.x1] == 'X') {
			fmt.Println(rectangle.volume)
			break
		} else {
			continue
		}
	}

}

func draw(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func addEdges(data []string, grid [][]byte) {
	// Add Points
	prevX, prevY := atoiCoords(data[0])
	grid[prevY][prevX] = '#'
	for i := 1; i < len(data)+1; i++ {
		var x, y int
		if i == len(data) {
			// final case, draw last line
			x, y = atoiCoords(data[0])
		} else {
			x, y = atoiCoords(data[i])
		}
		grid[y][x] = '#'
		if x == prevX {
			for j := 1; j < abs(y-prevY); j++ {
				if prevY > y {
					grid[prevY-j][x] = 'X'
				}
				if y > prevY {
					grid[y-j][x] = 'X'
				}
			}
		}
		if y == prevY {
			for j := 1; j < abs(x-prevX); j++ {
				if prevX > x {
					grid[y][prevX-j] = 'X'
				}
				if x > prevX {
					grid[y][x-j] = 'X'
				}
			}
		}
		prevX = x
		prevY = y
	}
}

func buildGrid(data []string) [][]byte {
	rowSize := 0
	colSize := 0
	for i := range data {
		x, y := atoiCoords(data[i])
		if x > colSize {
			colSize = x
		}
		if y > rowSize {
			rowSize = y
		}
	}
	grid := make([][]byte, rowSize+2)
	for i := range grid {
		grid[i] = make([]byte, colSize+2)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

func fillSegments(grid [][]byte) {
	for r := range grid {
		row := grid[r]
		n := len(row)

		left := -1
		for i := range n {
			if row[i] == 'X' || row[i] == '#' {
				if left == -1 {
					left = i
				} else {
					// fill between left and i
					for j := left + 1; j < i; j++ {
						if row[j] == '.' {
							row[j] = 'X'
						}
					}
					left = i
				}
			}
		}
	}
}

func main() {
	data := helper.ParseInput(helper.LoadInput("mockinput.txt"))
	grid := buildGrid(data)
	addEdges(data, grid)
	fillSegments(grid)
	draw(grid)
	puzzle(data, grid)
}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// const InputFile = "input.txt"

// type Tile struct {
// 	x int
// 	y int
// }

// func (rt *Tile) areaWith(other *Tile) int {
// 	return (abs(rt.x-other.x) + 1) * (abs(rt.y-other.y) + 1)
// }

// func abs(val int) int {
// 	if val < 0 {
// 		return -val
// 	}
// 	return val
// }

// func parseRedTiles(input string) ([]Tile, error) {
// 	trimmed := strings.TrimSpace(input)
// 	lines := strings.Split(trimmed, "\n")
// 	redTiles := make([]Tile, 0, len(lines))

// 	for _, line := range lines {
// 		parts := strings.Split(line, ",")
// 		x, err := strconv.Atoi(parts[0])
// 		if err != nil {
// 			return nil, err
// 		}
// 		y, err := strconv.Atoi(parts[1])
// 		if err != nil {
// 			return nil, err
// 		}

// 		rt := Tile{x: x, y: y}
// 		redTiles = append(redTiles, rt)
// 	}
// 	return redTiles, nil
// }

// func part1(tiles []Tile) {
// 	maxRect := 0
// 	for i := 0; i < len(tiles)-1; i++ {
// 		for j := i + 1; j < len(tiles); j++ {
// 			area := tiles[i].areaWith(&tiles[j])
// 			if area > maxRect {
// 				maxRect = area
// 			}
// 		}
// 	}
// 	fmt.Println(maxRect)
// }

// type Segment struct {
// 	A Tile
// 	B Tile
// }

// func computeGreenSegments(redTiles []Tile) []Segment {
// 	segments := make([]Segment, 0, len(redTiles)+1)
// 	for i := range len(redTiles) - 1 {
// 		segment := Segment{A: redTiles[i], B: redTiles[i+1]}
// 		segments = append(segments, segment)
// 	}
// 	// connect the last two
// 	segments = append(segments, Segment{A: redTiles[len(redTiles)-1], B: redTiles[0]})
// 	return segments
// }

// func (s *Segment) intersectsRect(rectA Tile, rectB Tile) bool {
// 	recMinX := min(rectA.x, rectB.x) + 1
// 	recMaxX := max(rectA.x, rectB.x) - 1
// 	recMinY := min(rectA.y, rectB.y) + 1
// 	recMaxY := max(rectA.y, rectB.y) - 1

// 	segMinX := min(s.A.x, s.B.x)
// 	segMaxX := max(s.A.x, s.B.x)
// 	segMinY := min(s.A.y, s.B.y)
// 	segMaxY := max(s.A.y, s.B.y)

// 	if segMaxX < recMinX || segMinX > recMaxX {
// 		return false
// 	}
// 	if segMaxY < recMinY || segMinY > recMaxY {
// 		return false
// 	}
// 	return true
// }

// // Compute the green tile segments. For every possible rectangle, if any green segment
// // intersects it's invalid, so we discard.
// func part2(redTiles []Tile) {
// 	greenSegments := computeGreenSegments(redTiles)
// 	fmt.Println(greenSegments)
// 	maxRect := 0
// 	for i := 0; i < len(redTiles)-1; i++ {
// 	main:
// 		for j := i + 1; j < len(redTiles); j++ {
// 			area := redTiles[i].areaWith(&redTiles[j])
// 			if area < maxRect {
// 				continue
// 			}
// 			for _, greenSegment := range greenSegments {
// 				if greenSegment.intersectsRect(redTiles[i], redTiles[j]) {
// 					continue main
// 				}
// 			}
// 			maxRect = area
// 		}
// 	}
// 	fmt.Println("Max rectangle of Green Tiles has area", maxRect)
// }

// func main() {
// 	input, err := os.ReadFile("input.txt")
// 	if err != nil {
// 		fmt.Println("error reading file", err)
// 		os.Exit(1)
// 	}
// 	redTiles, err := parseRedTiles(string(input))
// 	if err != nil {
// 		fmt.Println("error parsing input", err)
// 		os.Exit(1)
// 	}
// 	part1(redTiles)
// 	part2(redTiles)
// }
