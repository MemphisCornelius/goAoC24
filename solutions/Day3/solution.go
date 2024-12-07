package day3

import (
	"bufio"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Day3 struct{}

func (_ Day3) Run1(s *bufio.Scanner) int {
	sum := 0

	for s.Scan() {
		memory := s.Text()

		for i := strings.Index(memory, "mul("); i != -1; i = strings.Index(memory, "mul(") {

			value, m := parseMul(memory, i)
			sum += value
			memory = m

		}
	}

	return sum
}

func (d Day3) Run2(s *bufio.Scanner) int {
	sum := 0
	enabled := true
	for s.Scan() {
		memory := s.Text()
		splits := strings.Split(memory, "do")
		for _, v := range splits {
			if strings.HasPrefix(v, "n't()") {
				enabled = false
			}
			if strings.HasPrefix(v, "()") {
				enabled = true
			}
			if enabled {
				sum += d.Run1(bufio.NewScanner(strings.NewReader(v)))
			}
		}
	}

	return sum
}

func parseMul(memory string, i int) (int, string) {
	num1, len1 := parseInt(memory, i+4)
	if len1 == 0 {
		return 0, memory[i+4:]
	}
	if memory[i+4+len1] != byte(',') {
		return 0, memory[i+4+len1:]
	}
	num2, len2 := parseInt(memory, i+4+len1+1)
	if len2 == 0 {
		return 0, memory[i+4+len1+1:]
	}
	if memory[i+4+len1+1+len2] != byte(')') {
		return 0, memory[i+4+len1+1+len2:]
	}
	return num1 * num2, memory[i+4+len1+1+len2+1:]

}

func parseInt(memory string, i int) (int, int) {
	s := ""
	for c := memory[i]; unicode.IsDigit(rune(c)); c = memory[i] {
		s += string(c)
		i += 1
	}

	if len(s) <= 0 && len(s) > 3 {
		return 0, len(s)
	}

	if num, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
	} else {
		return num, len(s)
	}

	panic("should be unreachable")
}
