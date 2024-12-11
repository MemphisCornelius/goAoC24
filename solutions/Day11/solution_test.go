package day11_test

import (
	"AoC24/solutions/Day11"
	"bufio"
	"strings"
	"testing"
)

func TestRun1(t *testing.T) {
	input := "125 17"
	want := 55312
	have := day11.Day11{}.Run1(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %d. got %d", want, have)
	}
}
