package day8_test

import (
	"AoC24/solutions/Day8"
	"bufio"
	"strings"
	"testing"
)

const input = "............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............"

func TestRun1(t *testing.T) {
	want := 14
	have := day8.Day8{}.Run1(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %d. got %d", want, have)
	}
}

func TestRun2(t *testing.T) {
	want := 34
	have := day8.Day8{}.Run2(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %d. got %d", want, have)
	}
}
