package day4_test

import (
	day4 "AoC24/solutions/Day4"
	"bufio"
	"strings"
	"testing"
)

func TestRun1(t *testing.T) {
	input := "MMMSXXMASM\n" +
		"MSAMXMSMSA\n" +
		"AMXSXMAAMM\n" +
		"MSAMASMSMX\n" +
		"XMASAMXAMM\n" +
		"XXAMMXXAMA\n" +
		"SMSMSASXSS\n" +
		"SAXAMASAAA\n" +
		"MAMMMXMMMM\n" +
		"MXMXAXMASX"
	want := 18
	have := day4.Day4{}.Run1(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %v. got %v", want, have)
	}
}

func TestRun2(t *testing.T) {
	input := "MMMSXXMASM\n" +
		"MSAMXMSMSA\n" +
		"AMXSXMAAMM\n" +
		"MSAMASMSMX\n" +
		"XMASAMXAMM\n" +
		"XXAMMXXAMA\n" +
		"SMSMSASXSS\n" +
		"SAXAMASAAA\n" +
		"MAMMMXMMMM\n" +
		"MXMXAXMASX"
	want := 9
	have := day4.Day4{}.Run2(bufio.NewScanner(strings.NewReader(input)))

	if want != have {

		t.Fatalf("Run1: wanted %v. got %v", want, have)
	}
}
