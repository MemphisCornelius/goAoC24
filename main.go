package main

import (
	day1 "AoC24/solutions/Day1"
	day10 "AoC24/solutions/Day10"
	day11 "AoC24/solutions/Day11"
	day2 "AoC24/solutions/Day2"
	day3 "AoC24/solutions/Day3"
	day4 "AoC24/solutions/Day4"
	day5 "AoC24/solutions/Day5"
	day6 "AoC24/solutions/Day6"
	day7 "AoC24/solutions/Day7"
	day8 "AoC24/solutions/Day8"
	day9 "AoC24/solutions/Day9"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/tiagomelo/go-clipboard/clipboard"
)

type Solvable interface {
	Run1(*bufio.Scanner) int
	Run2(*bufio.Scanner) int
}

var day = flag.Int("d", time.Now().Day(), "day")
var part = flag.Int("p", 1, "part")

func main() {
	flag.Parse()

	if _, err := os.Stat(fmt.Sprintf("./solutions/Day%d", *day)); os.IsNotExist(err) {
		createSolutionFolder()
	}

	if _, err := os.Stat(fmt.Sprintf("./solutions/Day%d/input.txt", *day)); os.IsNotExist(err) {
		if err := download(*day); err != nil {
			log.Fatal(err)
		}
	}

	inputFile, err := os.Open(fmt.Sprintf("solutions/Day%v/input.txt", *day))
	if err != nil {
		log.Fatal(err)
	}

	var solution Solvable

	switch *day {
	case 1:
		solution = day1.Day1{}
	case 2:
		solution = day2.Day2{}
	case 3:
		solution = day3.Day3{}
	case 4:
		solution = day4.Day4{}
	case 5:
		solution = day5.Day5{}
	case 6:
		solution = day6.Day6{}
	case 7:
		solution = day7.Day7{}
	case 8:
		solution = day8.Day8{}
	case 9:
		solution = day9.Day9{}
	case 10:
		solution = day10.Day10{}
	case 11:
		solution = day11.Day11{}
	default:
		log.Fatalf("Day %d not available", *day)
	}

	log.Printf("Day %v\n", *day)

	var sol func(*bufio.Scanner) int
	inputScanner := bufio.NewScanner(inputFile)
	switch *part {
	case 1:
		sol = solution.Run1
	case 2:
		sol = solution.Run2
	default:
		log.Fatalf("Unknown part %s", os.Args[1])
	}

	startTime := time.Now()
	answer := sol(inputScanner)
	timeTook := time.Since(startTime)

	c := clipboard.New()

	if err := c.CopyText(strconv.Itoa(answer)); err != nil {
		log.Fatal(err)
	}

	log.Printf("Part %d: %d took %s\n", *part, answer, timeTook.String())
}

func createSolutionFolder() {

	path := fmt.Sprintf("./solutions/Day%d", *day)
	os.Mkdir(path, os.ModePerm)

	os.WriteFile(path+"/solution.go", []byte(fmt.Sprintf("package day%d\n\nimport (\n	\"bufio\"\n)\n\ntype Day%d struct{}\n\nfunc (_ Day%d) Run1(s *bufio.Scanner) int {\n	sum := 0\n	for s.Scan() {\n	}\n\n	return sum\n}\n\nfunc (_ Day%d) Run2(s *bufio.Scanner) int {\n	sum := 0\n	for s.Scan() {\n	}\n\n	return sum\n}", *day, *day, *day, *day)), 0644)

	os.WriteFile(path+"/solution_test.go", []byte(fmt.Sprintf("package day%d_test\n\nimport (\n	\"AoC24/solutions/Day%d\"\n	\"bufio\"\n	\"strings\"\n	\"testing\"\n)\n\nfunc TestRun1(t *testing.T) {\n	input := \"\"\n	want := -1\n	have := day%d.Day%d{}.Run1(bufio.NewScanner(strings.NewReader(input)))\n\n	if want != have {\n\n		t.Fatalf(\"Run1: wanted %%d. got %%d\", want, have)\n	}\n}\n\nfunc TestRun2(t *testing.T) {\n	input := \"\"\n	want := -1\n	have := day%d.Day%d{}.Run2(bufio.NewScanner(strings.NewReader(input)))\n\n	if want != have {\n\n		t.Fatalf(\"Run1: wanted %%d. got %%d\", want, have)\n	}\n}", *day, *day, *day, *day, *day, *day)), 0644)

}

func download(day int) error {
	godotenv.Load()
	client := new(http.Client)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", 2024, day), nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "github.com/GreenLightning/advent-of-code-downloader")

	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", os.Getenv("SESSION")
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC

	file, err := os.OpenFile(fmt.Sprintf("solutions/Day%d/input.txt", day), flags, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
