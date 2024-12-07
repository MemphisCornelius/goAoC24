package day4

import (
	"bufio"
	"strings"
)

type Day4 struct{}

var crossword [][]string

func (_ Day4) Run1(s *bufio.Scanner) int {
	sum := 0
	lines := make([]string, 0)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	crossword = make([][]string, len(lines))

	for i, v := range lines {
		crossword[i] = strings.Split(v, "")
	}

	for i := range crossword {
		for j := range crossword[i] {
			if crossword[i][j] == "X" {
				sum += countXmasAround(i, j)
			}
		}
	}

	return sum
}

func countXmasAround(i, j int) int {
	count := 0
	dirs := []string{
		get(i, j+1) + get(i, j+2) + get(i, j+3), //ltor
		get(i, j-1) + get(i, j-2) + get(i, j-3), //rtol
		get(i+1, j) + get(i+2, j) + get(i+3, j), //utod
		get(i-1, j) + get(i-2, j) + get(i-3, j), //dtou
		get(i+1, j+1) + get(i+2, j+2) + get(i+3, j+3),
		get(i+1, j-1) + get(i+2, j-2) + get(i+3, j-3),
		get(i-1, j+1) + get(i-2, j+2) + get(i-3, j+3),
		get(i-1, j-1) + get(i-2, j-2) + get(i-3, j-3),
	}

	for _, v := range dirs {
		if v == "MAS" {
			count += 1
		}
	}

	return count
}

func get(i, j int) string {
	if i < 0 || j < 0 || i >= len(crossword) || j >= len(crossword[0]) {
		return ""
	}
	return crossword[i][j]
}

func (_ Day4) Run2(s *bufio.Scanner) int {
	sum := 0
	lines := make([]string, 0)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	crossword = make([][]string, len(lines))

	for i, v := range lines {
		crossword[i] = strings.Split(v, "")
	}

	for i := range crossword {
		for j := range crossword[i] {
			if crossword[i][j] == "A" {
				sum += countMasCross(i, j)
			}
		}
	}

	return sum

}

func countMasCross(i, j int) int {
	ltor := get(i-1, j-1) + get(i+1, j+1)
	rtol := get(i-1, j+1) + get(i+1, j-1)

	if (ltor == "MS" || ltor == "SM") && (rtol == "MS" || rtol == "SM") {
		return 1
	}

	return 0
}
