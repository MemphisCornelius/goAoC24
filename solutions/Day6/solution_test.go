package day6_test

import (
	"AoC24/solutions/Day6"
	"bufio"
	"strings"
	"testing"
)

func TestRun1(t *testing.T) {
	input := "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#..."
	want := 41
	have := day6.Day6{}.Run1(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %v. got %v", want, have)
	}
}

func TestRun2(t *testing.T) {
	input := "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#..."
	want := 6
	have := day6.Day6{}.Run2(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %v. got %v", want, have)
	}
}
