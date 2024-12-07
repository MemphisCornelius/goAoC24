package day7_test

import (
	"AoC24/solutions/Day7"
	"bufio"
	"strings"
	"testing"
)

func TestRun1(t *testing.T) {
	input := "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20\n"
	want := 3749
	have := day7.Day7{}.Run1(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %d. got %d", want, have)
	}
}

func TestRun2(t *testing.T) {
	input := "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20\n"
	want := 11387
	have := day7.Day7{}.Run2(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %d. got %d", want, have)
	}
}
