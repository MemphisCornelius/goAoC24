package day3_test

import (
	day3 "AoC24/solutions/Day3"
	"bufio"
	"strings"
	"testing"
)

func TestRun1(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	want := 161
	have := day3.Day3{}.Run1(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %v. got %v", want, have)
	}
}

func TestRun2(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	want := 48
	have := day3.Day3{}.Run2(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %v. got %v", want, have)
	}
}
