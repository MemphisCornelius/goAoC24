package day8

import (
	"AoC24/util"
	"bufio"
	"strings"
)

type Day8 struct{}

type Point struct {
	X int
	Y int
}

func (_ Day8) Run1(s *bufio.Scanner) int {
	freqMap := createMap(s)
	sizeX, sizeY := len(freqMap[0]), len(freqMap)

	antennaPositions := getAntennaPositions(freqMap)
	antiNodePositions := getAntiNodePositions(antennaPositions, sizeX, sizeY, 1)

	return len(antiNodePositions)
}

func (_ Day8) Run2(s *bufio.Scanner) int {
	freqMap := createMap(s)
	sizeX, sizeY := len(freqMap[0]), len(freqMap)

	antennaPositions := getAntennaPositions(freqMap)
	antiNodePositions := getAntiNodePositions(antennaPositions, sizeX, sizeY, 2)

	return len(antiNodePositions)
}

func createMap(s *bufio.Scanner) [][]string {
	freqMap := make([][]string, 0)
	for s.Scan() {
		freqMap = append(freqMap, strings.Split(s.Text(), ""))
	}
	return freqMap
}

func getAntennaPositions(freqMap [][]string) map[string][]Point {

	antennaPositions := make(map[string][]Point)
	for y, line := range freqMap {
		for x, field := range line {
			if field != "." {
				antennaPositions[field] = append(antennaPositions[field], Point{X: x, Y: y})
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

func getAntiNodePositions(antennaPositions map[string][]Point, sizeX, sizeY, part int) map[Point]bool {
	antiNodePositions := make(map[Point]bool)

	for _, positions := range antennaPositions {
		for i, position := range positions {
			for j, other := range positions {
				if i == j {
					continue
				}
				distanceX, distanceY := diff(position, other)
				if part == 2 {
					distanceX, distanceY = shorten(distanceX, distanceY)
				}
				pos1 := []Point{{X: position.X + distanceX, Y: position.Y + distanceY}}
				pos2 := []Point{{X: other.X - distanceX, Y: other.Y - distanceY}}

				if part == 2 {
					for i := 0; inBound(pos1[i], sizeX, sizeY); i += 1 {
						pos1 = append(pos1, Point{X: position.X + i*distanceX, Y: position.Y + i*distanceY})
					}
					for i := 0; inBound(pos2[i], sizeX, sizeY); i += 1 {
						pos2 = append(pos2, Point{X: position.X - i*distanceX, Y: position.Y - i*distanceY})
					}
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
	return antiNodePositions
}

func shorten(x, y int) (int, int) {
	gcd := util.Gcd(x, y)
	return x / gcd, y / gcd
}
