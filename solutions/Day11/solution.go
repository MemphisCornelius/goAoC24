package day11

import (
	"bufio"
	"log"
	"math"
	"strconv"
	"strings"
)

type Day11 struct{}

type Stones map[int]int

func (_ Day11) Run1(s *bufio.Scanner) int {
	return getStonesAfterNBlinks(25, parseInput(s))
}

func (_ Day11) Run2(s *bufio.Scanner) int {
	return getStonesAfterNBlinks(75, parseInput(s))
}

func getStonesAfterNBlinks(n int, stones Stones) int {
	sum := 0
	for i := 0; i < n; i += 1 {
		tmp := make(Stones, len(stones))
		for k, v := range stones {
			switch {
			case k == 0:
				tmp[1] += v
			case lenNumber(k)%2 == 0:
				l, r := splitNumber(k)
				tmp[r] += v
				tmp[l] += v
			default:
				tmp[k*2024] += v
			}
		}
		stones = tmp
	}

	for _, v := range stones {
		sum += v
	}

	return sum
}

func lenNumber(i int) int {
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}

func splitNumber(k int) (int, int) {

	half := lenNumber(k) / 2

	r := k % int(math.Pow10(half))
	l := (k - r) / int(math.Pow10(half))

	return l, r
}

func parseInput(s *bufio.Scanner) Stones {
	s.Scan()
	in := strings.Split(s.Text(), " ")
	stones := make(Stones)

	for _, v := range in {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		stones[i] += 1
	}
	return stones
}
