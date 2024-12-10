package main

import (
	day1 "AoC24/solutions/Day1"
	day10 "AoC24/solutions/Day10"
	day2 "AoC24/solutions/Day2"
	day3 "AoC24/solutions/Day3"
	day4 "AoC24/solutions/Day4"
	day5 "AoC24/solutions/Day5"
	day6 "AoC24/solutions/Day6"
	day7 "AoC24/solutions/Day7"
	day8 "AoC24/solutions/Day8"
	day9 "AoC24/solutions/Day9"
)

func registerSolution(day int, solution Solvable) {
	solutions[day] = solution
}

func RegisterSolutions() {
	registerSolution(1, day1.Day1{})
	registerSolution(2, day2.Day2{})
	registerSolution(3, day3.Day3{})
	registerSolution(4, day4.Day4{})
	registerSolution(5, day5.Day5{})
	registerSolution(6, day6.Day6{})
	registerSolution(7, day7.Day7{})
	registerSolution(8, day8.Day8{})
	registerSolution(9, day9.Day9{})
	registerSolution(10, day10.Day10{})
}
