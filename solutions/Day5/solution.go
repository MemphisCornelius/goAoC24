package day5

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type Day5 struct{}

type Rule struct {
	Left  int
	Right int
}

func preProcess(s *bufio.Scanner) ([]Rule, [][]int) {
	rules := make([]Rule, 0)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		line := strings.Split(s.Text(), "|")
		l, e := strconv.Atoi(line[0])
		r, e := strconv.Atoi(line[1])

		if e != nil {
			log.Fatal(e)
		}

		rules = append(rules, Rule{Left: l, Right: r})
	}

	prints := make([][]int, 0)
	for s.Scan() {
		line := strings.Split(s.Text(), ",")

		conv := make([]int, len(line))
		for i, v := range line {
			var e error
			conv[i], e = strconv.Atoi(v)
			if e != nil {
				log.Fatal(e)
			}
		}
		prints = append(prints, conv)
	}
	return rules, prints
}
func (_ Day5) Run1(s *bufio.Scanner) int {
	sum := 0

	rules, prints := preProcess(s)

	for _, v := range prints {
		if isRightOrdering(v, rules) {
			sum += v[len(v)/2]
		}
	}

	return sum
}

func isRightOrdering(line []int, rules []Rule) bool {
	l, r := getWrongPos(line, rules)
	return l == -1 && r == -1
}

func getWrongPos(line []int, rules []Rule) (int, int) {
	for _, v := range rules {
		posL, posR := pos(v.Left, line), pos(v.Right, line)
		if posL >= 0 && posR >= 0 && posL > posR {
			return posL, posR
		}
	}
	return -1, -1
}

func pos(k int, vs []int) int {
	for i, v := range vs {
		if k == v {
			return i
		}
	}
	return -1
}

func (_ Day5) Run2(s *bufio.Scanner) int {
	sum := 0
	rules, prints := preProcess(s)
	for _, v := range prints {
		if !isRightOrdering(v, rules) {
			for !isRightOrdering(v, rules) {
				correct(&v, rules)
			}
			sum += v[len(v)/2]
		}
	}
	return sum
}

func correct(line *[]int, rules []Rule) {
	posL, posR := getWrongPos(*line, rules)
	swap(posL, posR, line)
}

func swap(i, j int, a *[]int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}
