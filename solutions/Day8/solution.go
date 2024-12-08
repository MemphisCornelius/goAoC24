package day8

import (
	"bufio"
	"strings"
)

type Day8 struct{}

type Point struct {
	X int
	Y int
}

func (_ Day8) Run1(s *bufio.Scanner) int {
	freqMap := make([][]string, 0)
	for s.Scan() {
		freqMap = append(freqMap, strings.Split(s.Text(), ""))
	}
	sizeX, sizeY := len(freqMap[0]), len(freqMap)

	antennaPositions := getAntennaPositions(freqMap)

	antiNodePositions := make(map[Point]bool)

	for _, positions := range antennaPositions {
		for i, position := range positions {
			for j, other := range positions {
				if i != j {
					distanceX, distanceY := diff(position, other)
					pos1, pos2 := Point{X: position.X + distanceX, Y: position.Y + distanceY}, Point{X: other.X - distanceX, Y: other.Y - distanceY}
					if inBound(pos1, sizeX, sizeY) && !antiNodePositions[pos1] {
						antiNodePositions[pos1] = true
					}
					if inBound(pos2, sizeX, sizeY) && !antiNodePositions[pos2] {
						antiNodePositions[pos2] = true
					}
				}
			}
		}
	}

	return len(antiNodePositions)
}

func getAntennaPositions(freqMap [][]string) map[string][]Point {

	antennaPositions := make(map[string][]Point)
	for y, line := range freqMap {
		for x, field := range line {
			if field != "." {
				a, ok := antennaPositions[field]
				pos := Point{X: x, Y: y}
				if ok {
					antennaPositions[field] = append(a, pos)
				} else {
					antennaPositions[field] = []Point{pos}
				}
			}
		}
	}
	return antennaPositions
}

func diff(p1, p2 Point) (int, int) {
	return p1.X - p2.X, p1.Y - p2.Y
}

func inBound(p Point, x, y int) bool {
	return p.X >= 0 && p.X < x && p.Y >= 0 && p.Y < y

}

func (_ Day8) Run2(s *bufio.Scanner) int {
	freqMap := make([][]string, 0)
	for s.Scan() {
		freqMap = append(freqMap, strings.Split(s.Text(), ""))
	}
	sizeX, sizeY := len(freqMap[0]), len(freqMap)
	antennaPositions := getAntennaPositions(freqMap)
	antiNodePositions := make(map[Point]bool)

	for _, positions := range antennaPositions {
		for i, position := range positions {
			for j, other := range positions {
				if i != j {
					distanceX, distanceY := diff(position, other)
					distanceX, distanceY = shorten(distanceX, distanceY)
					pos1 := []Point{{X: position.X + distanceX, Y: position.Y + distanceY}}
					pos2 := []Point{{X: other.X - distanceX, Y: other.Y - distanceY}}
					for i := 0; inBound(pos1[i], sizeX, sizeY); i += 1 {
						pos1 = append(pos1, Point{X: position.X + (i+1)*distanceX, Y: position.Y + (i+1)*distanceY})
					}
					for i := 0; inBound(pos2[i], sizeX, sizeY); i += 1 {
						pos2 = append(pos2, Point{X: position.X - (i+1)*distanceX, Y: position.Y - (i+1)*distanceY})
					}

					for _, p := range pos1 {
						if inBound(p, sizeX, sizeY) && !antiNodePositions[p] {
							antiNodePositions[p] = true
						}
					}
					for _, p := range pos2 {
						if inBound(p, sizeX, sizeY) && !antiNodePositions[p] {
							antiNodePositions[p] = true
						}
					}
				}
			}
		}
	}
	return len(antiNodePositions)
}

func shorten(x, y int) (int, int) {
	gcd := gcd(x, y)
	return x / gcd, y / gcd
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

