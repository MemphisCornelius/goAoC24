package day1

import (
	"bufio"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day1 struct{}

func (_ Day1) Run1(s *bufio.Scanner) int {
	left, right := make([]int, 0), make([]int, 0)
	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, "   ")
		if l, err := strconv.Atoi(split[0]); err == nil {
			left = append(left, l)
		} else {
			log.Fatal(err)
		}
		if r, err := strconv.Atoi(split[1]); err == nil {
			right = append(right, r)
		} else {
			log.Fatal(err)
		}
	}
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := range left {
		diff := left[i] - right[i]
		sum += int(math.Abs(float64(diff)))
	}
	return sum
}

func (_ Day1) Run2(s *bufio.Scanner) int {
	left, right := make([]int, 0), make(map[int]int)
	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, "   ")
		if l, err := strconv.Atoi(split[0]); err == nil {
			left = append(left, l)
		} else {
			log.Fatal(err)
		}
		if r, err := strconv.Atoi(split[1]); err == nil {
			if _, ok := right[r]; !ok {
				right[r] = 0
			}
			right[r] += 1
		} else {
			log.Fatal(err)
		}
	}

	sum := 0
	for _, v := range left {
		sum += v * right[v]
	}
	return sum
}
