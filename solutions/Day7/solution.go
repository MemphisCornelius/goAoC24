package day7

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Day7 struct{}

func (_ Day7) Run1(s *bufio.Scanner) int {
	sum := 0
	for s.Scan() {
		line := s.Text()
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

		if validOperation(result, values) {
			sum += result
		}
	}

	return sum
}

func validOperation(result int, values []int) bool {
	if len(values) == 0 || values[0] > result {
		return false
	}

	if len(values) == 1 {
		return result == values[0]
	}

	plus := make([]int, len(values))
	mult := make([]int, len(values))

	copy(plus, values)
	copy(mult, values)

	plus[1] += plus[0]
	mult[1] *= mult[0]

	return validOperation(result, plus[1:]) || validOperation(result, mult[1:])
}

func (_ Day7) Run2(s *bufio.Scanner) int {
	sum := 0
	for s.Scan() {
		line := s.Text()
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

		if validOperation2(result, values) {
			sum += result
		}
	}

	return sum
}

func validOperation2(result int, values []int) bool {
	if len(values) == 0 || values[0] > result {
		return false
	}

	if len(values) == 1 {
		return result == values[0]
	}

	plus := make([]int, len(values))
	mult := make([]int, len(values))
	conc := make([]int, len(values))

	copy(plus, values)
	copy(mult, values)
	copy(conc, values)

	plus[1] += plus[0]
	mult[1] *= mult[0]
	conc[1] = concat(conc[0], conc[1])

	return validOperation2(result, plus[1:]) || validOperation2(result, mult[1:]) || validOperation2(result, conc[1:])
}

func concat(left, right int) int {
	r, err := strconv.Atoi(fmt.Sprintf("%d%d", left, right))
	if err != nil {
		log.Fatal(err)
	}
	return r
}
