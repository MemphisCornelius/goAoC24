package day10_test

import (
	"AoC24/solutions/Day10"
	"bufio"
	"strings"
	"testing"
)

const input = "89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732"

func TestRun1(t *testing.T) {
	want := 36
	have := day10.Day10{}.Run1(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %d. got %d", want, have)
	}
}

func TestRun2(t *testing.T) {
	want := 81
	have := day10.Day10{}.Run2(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run2: wanted %d. got %d", want, have)
	}
}

