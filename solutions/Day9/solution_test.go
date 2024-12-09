package day9_test

import (
	"AoC24/solutions/Day9"
	"bufio"
	"strings"
	"testing"
)

const input = "2333133121414131402"

func TestRun1(t *testing.T) {
	want := 1928
	have := day9.Day9{}.Run1(bufio.NewScanner(strings.NewReader(input)))

	if want != have {
		t.Fatalf("Run1: wanted %d. got %d", want, have)
	}
}

func TestRun2(t *testing.T) {
	want := 2858
	have := day9.Day9{}.Run2(bufio.NewScanner(strings.NewReader(input)))

	if want != have {
		t.Fatalf("Run1: wanted %d. got %d", want, have)
	}
}

