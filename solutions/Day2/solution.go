package day2

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type Day2 struct{}

func (_ Day2) Run1(s *bufio.Scanner) int {

	sum := 0
	for s.Scan() {
		levels := convertInput(s.Text())

		if ValidReport(levels) {
			sum += 1
		}
	}
	return sum
}

func (_ Day2) Run2(s *bufio.Scanner) int {
	sum := 0
	for s.Scan() {
		levels := convertInput(s.Text())
		if ValidReport2(levels) {
			sum += 1
		}
	}
	return sum
}

func convertInput(report string) []int {
	_levels := strings.Split(report, " ")
	levels := make([]int, len(_levels))
	for i, v := range _levels {
		if level, err := strconv.Atoi(v); err != nil {
			log.Fatal(err)
		} else {
			levels[i] = level
		}
	}
	return levels
}

func ValidReport2(report []int) bool {
	validWithoutRemove := ValidReport(report)

	if !validWithoutRemove {
		for i := range report {
			c := make([]int, len(report))
			copy(c, report)
			removed := ValidReport(remove(c, i))
			if removed {
				return true
			}
		}
	}

	return validWithoutRemove
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func ValidReport(report []int) bool {

	inc, dec := true, true
	for i := 0; i < len(report)-1; i += 1 {
		if report[i+1] <= report[i] || report[i+1] > report[i]+3 {
			inc = false
			break
		}
	}
	for i := 0; i < len(report)-1; i += 1 {
		if report[i+1] >= report[i] || report[i+1] < report[i]-3 {
			dec = false
			break
		}
	}

	return inc || dec
}
