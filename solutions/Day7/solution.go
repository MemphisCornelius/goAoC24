package day7

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Day7 struct{}

func parseInput(line string) (int, []int) {

	resultsplit := strings.Split(line, ": ")
	result, err := strconv.Atoi(resultsplit[0])
	if err != nil {
		log.Fatal(err)
	}
	numbers := strings.Split(resultsplit[1], " ")
	values := make([]int, len(numbers))
	for i, v := range numbers {
		values[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
	}
	return result, values
}

func (_ Day7) Run1(s *bufio.Scanner) int {
	sum := 0
	for s.Scan() {
		result, values := parseInput(s.Text())

		if validOperation(1, result, values) {
			sum += result
		}
	}
	return sum
}

func validOperation(part, result int, values []int) bool {
	if len(values) == 0 || values[0] > result {
		return false
	}

	if len(values) == 1 {
		return result == values[0]
	}

	plus := make([]int, len(values))
	copy(plus, values)
	plus[1] += plus[0]

	mult := make([]int, len(values))
	copy(mult, values)
	mult[1] *= mult[0]

	if part == 2 {
		conc := make([]int, len(values))
		copy(conc, values)
		conc[1] = concat(conc[0], conc[1])

		return validOperation(part, result, plus[1:]) || validOperation(part, result, mult[1:]) || validOperation(part, result, conc[1:])
	}
	return validOperation(part, result, plus[1:]) || validOperation(part, result, mult[1:])
}

func (_ Day7) Run2(s *bufio.Scanner) int {
	sum := 0
	for s.Scan() {
		result, values := parseInput(s.Text())

		if validOperation(2, result, values) {
			sum += result
		}
	}
	return sum
}

func concat(left, right int) int {
	r, err := strconv.Atoi(fmt.Sprintf("%d%d", left, right))
	if err != nil {
		log.Fatal(err)
	}
	return r
}
