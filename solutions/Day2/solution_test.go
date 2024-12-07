package day2_test

import (
	day2 "AoC24/solutions/Day2"
	"strconv"
	"testing"
)

func printSlice(s []int) string {
	str := "["

	for _, v := range s {
		str += strconv.FormatInt(int64(v), 10) + ","
	}
	str += "]"
	return str
}

func TestValidReport1(t *testing.T) {
	in := []int{7, 6, 4, 2, 1}
	want := true
	have := day2.ValidReport2(in)

	if want != have {
		t.Fatalf(`ValidReport2(%s) = %v, want match for %v, `, printSlice(in), have, want)
	}
}
func TestValidReport2(t *testing.T) {
	in := []int{1, 2, 7, 8, 9}
	want := false
	have := day2.ValidReport2(in)

	if want != have {
		t.Fatalf(`ValidReport2(%s) = %v, want match for %v, `, printSlice(in), have, want)
	}
}
func TestValidReport3(t *testing.T) {
	in := []int{9, 7, 6, 2, 1}
	want := false
	have := day2.ValidReport2(in)

	if want != have {
		t.Fatalf(`ValidReport2(%s) = %v, want match for %v, `, printSlice(in), have, want)
	}
}
func TestValidReport4(t *testing.T) {
	in := []int{1, 3, 2, 4, 5}
	want := true
	have := day2.ValidReport2(in)

	if want != have {
		t.Fatalf(`ValidReport2(%s) = %v, want match for %v, `, printSlice(in), have, want)
	}
}
func TestValidReport5(t *testing.T) {
	in := []int{8, 6, 4, 4, 1}
	want := true
	have := day2.ValidReport2(in)

	if want != have {
		t.Fatalf(`ValidReport2(%s) = %v, want match for %v, `, printSlice(in), have, want)
	}
}
func TestValidReport6(t *testing.T) {
	in := []int{1, 3, 6, 7, 9}
	want := true
	have := day2.ValidReport2(in)

	if want != have {
		t.Fatalf(`ValidReport2(%s) = %v, want match for %v, `, printSlice(in), have, want)
	}
}
