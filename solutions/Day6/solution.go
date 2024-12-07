package day6

import (
	"bufio"
	"fmt"
	"strings"
)

type Day6 struct{}
type Dir string

const (
	North Dir = "^"
	East  Dir = ">"
	South Dir = "v"
	West  Dir = "<"
)

func (_ Day6) Run1(s *bufio.Scanner) int {
	input := make([][]string, 0)
	for s.Scan() {
		input = append(input, strings.Split(s.Text(), ""))
	}
	return run1(&input)
}

func run1(in *[][]string) int {
	sum := 1
	input := *in
	sizeX, sizeY := len(input[0]), len(input)
	x, y := findGuard(&input)
	dir := Dir(input[y][x])

	for !(y == 0 && dir == North || x == 0 && dir == West || y+1 == sizeY && dir == South || x+1 == sizeX && dir == East) {
		switch nextFiled(x, y, dir, &input) {
		case "#":
			dir = turn(x, y, dir, &input)
		case ".":
			sum += 1
			fallthrough
		case "X":
			x, y = step(x, y, &input)
		}
	}
	input[y][x] = "X"
	return sum
}

func step(x, y int, f *[][]string) (int, int) {
	dir := Dir((*f)[y][x])
	(*f)[y][x] = "X"

	switch dir {
	case North:
		y = y - 1
	case East:
		x = x + 1
	case South:
		y = y + 1
	case West:
		x = x - 1
	default:
		panic("unreachable nextFiled")
	}
	(*f)[y][x] = string(dir)
	return x, y
}

func nextFiled(x, y int, dir Dir, f *[][]string) string {
	switch dir {
	case North:
		return (*f)[y-1][x]
	case East:
		return (*f)[y][x+1]
	case South:
		return (*f)[y+1][x]
	case West:
		return (*f)[y][x-1]
	}
	panic("unreachable nextFiled")
}

func turn(x, y int, dir Dir, f *[][]string) Dir {
	switch dir {
	case North:
		dir = East
	case East:
		dir = South
	case South:
		dir = West
	case West:
		dir = North
	default:
		panic("unreachable turn")
	}
	(*f)[y][x] = string(dir)
	return dir
}

func findGuard(f *[][]string) (int, int) {
	for y := range *f {
		for x, char := range (*f)[y] {
			if char == string(North) || char == string(East) || char == string(South) || char == string(West) {
				return x, y
			}
		}
	}
	panic("unreachable findGuard")
}

func (_ Day6) Run2(s *bufio.Scanner) int {
	sum := 0
	input := make([][]string, 0)
	for s.Scan() {
		input = append(input, strings.Split(s.Text(), ""))
	}
	// optimizes solution (but was not my idea)
	startX, startY := findGuard(&input)
	run1(&input)
	input[startY][startX] = "^"

	for y := range input {
		for x := range input[y] {
			if input[y][x] == "X" { //in my solution then == "O"
				c := deepCopy(input)
				c[y][x] = "O"
				a := run(c)
				sum += a
			}
		}
	}
	return sum
}

func deepCopy(f [][]string) [][]string {
	c := make([][]string, 0)

	for _, v := range f {
		in := make([]string, len(v))
		copy(in, v)
		c = append(c, in)
	}

	return c
}

type DirPos struct {
	X   int
	Y   int
	Dir Dir
}

func run(input [][]string) int {
	sizeX, sizeY := len(input[0]), len(input)
	x, y := findGuard(&input)
	dir := Dir(input[y][x])
	visits := make(map[DirPos]bool)

	for !(y == 0 && dir == North || x == 0 && dir == West || y+1 == sizeY && dir == South || x+1 == sizeX && dir == East) {
		dirPos := DirPos{X: x, Y: y, Dir: dir}
		if visits[dirPos] {
			return 1
		}
		visits[dirPos] = true
		switch nextFiled(x, y, dir, &input) {
		case "O":
			fallthrough
		case "#":
			dir = turn(x, y, dir, &input)
			input[y][x] = "+"
		case "^":
			fallthrough
		case ">":
			fallthrough
		case "v":
			fallthrough
		case "<":
			fallthrough
		case ".":
			fallthrough
		case "X":
			fallthrough
		case "+":
			x, y = step2(x, y, dir, &input)
		}

	}
	return 0
}

func print2d(f [][]string) {
	for i := range f {
		fmt.Printf("%v\n", f[i])
	}
}

func step2(x, y int, dir Dir, f *[][]string) (int, int) {

	switch dir {
	case North:
		(*f)[y][x] = string(North)
		y = y - 1
	case East:
		(*f)[y][x] = string(East)
		x = x + 1
	case South:
		(*f)[y][x] = string(South)
		y = y + 1
	case West:
		(*f)[y][x] = string(West)
		x = x - 1
	default:
		panic("unreachable nextFiled")
	}
	(*f)[y][x] = string(dir)
	return x, y

}
