package day10

import (
	"bufio"
	"log"
	"strconv"
)

type Day10 struct{}

type Point struct {
	x, y int
}

type Map [][]int

func (_ Day10) Run1(s *bufio.Scanner) int {
	sum := 0
	topographicMap := parseInput(s)
	for _, s := range findPoints(0, topographicMap) {
		sum += getNumRoutes(s, topographicMap, map[Point]bool{})
	}

	return sum
}

func (_ Day10) Run2(s *bufio.Scanner) int {
	sum := 0
	topographicMap := parseInput(s)
	for _, s := range findPoints(0, topographicMap) {
		sum += findNumPaths(s, s, topographicMap)
	}

	return sum
}

func findNumPaths(start, before Point, m Map) int {
	if m[start.y][start.x] == 9 {
		return 1
	}

	sum := 0
	for _, v := range getArround(start, m) {
		if before != v && m[start.y][start.x]+1 == m[v.y][v.x] {
			sum += findNumPaths(v, start, m)
		}
	}
	return sum
}

func getNumRoutes(start Point, m Map, visited map[Point]bool) int {
	visited[start] = true
	if m[start.y][start.x] == 9 {
		return 1
	}

	sum := 0
	for _, v := range getArround(start, m) {
		if !visited[v] && m[start.y][start.x]+1 == m[v.y][v.x] {
			sum += getNumRoutes(v, m, visited)
		}
	}
	return sum
}

func getArround(p Point, m Map) []Point {
	arround := make([]Point, 0)
	if p.x > 0 {
		arround = append(arround, Point{x: p.x - 1, y: p.y})
	}
	if p.y > 0 {
		arround = append(arround, Point{x: p.x, y: p.y - 1})
	}
	if p.x < len(m[0])-1 {
		arround = append(arround, Point{x: p.x + 1, y: p.y})
	}
	if p.y < len(m)-1 {
		arround = append(arround, Point{x: p.x, y: p.y + 1})
	}

	return arround
}

func findPoints(height int, m Map) []Point {
	trilheads := make([]Point, 0)
	for y, hs := range m {
		for x, h := range hs {
			if h == height {
				trilheads = append(trilheads, Point{x: x, y: y})
			}
		}
	}
	return trilheads
}

func parseInput(s *bufio.Scanner) Map {
	topographicMap := make(Map, 0)

	for s.Scan() {
		line := []rune(s.Text())
		hights := make([]int, len(line))

		for i, v := range line {
			hight, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatal(err)
			}
			hights[i] = hight
		}
		topographicMap = append(topographicMap, hights)
	}

	return topographicMap
}