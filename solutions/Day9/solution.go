package day9

import (
	"AoC24/util"
	"bufio"
	"log"
	"slices"
	"strconv"
)

type Day9 struct{}

type Block struct {
	id, size int
}

func (_ Day9) Run1(s *bufio.Scanner) int {
	return checksum(defragment(convertToBlocks(parseInput(s), true), true))
}

func (_ Day9) Run2(s *bufio.Scanner) int {
	return checksum(defragment(convertToBlocks(parseInput(s), false), false))
}

func convertToBlocks(diskMap string, part1 bool) []Block {
	id := 0
	disk := make([]Block, 0)
	for i, v := range diskMap {
		count, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}
		if count <= 0 {
			continue
		}

		var content int
		if i%2 == 0 {
			content = id
			id += 1
		} else {
			content = -1
		}

		if part1 {
			for k := 0; k < count; k += 1 {
				disk = append(disk, Block{id: content, size: 1})
			}
		} else {
			disk = append(disk, Block{id: content, size: count})
		}
	}
	return disk
}

func defragment(blocks []Block, part1 bool) []Block {
	for i := len(blocks) - 1; i >= 0; i -= 1 {
		emptySpace := findEmpty(blocks[:i], blocks[i].size)

		if emptySpace == -1 {
			continue
		}

		if blocks[i].size < blocks[emptySpace].size {
			sliced := Block{id: -1, size: blocks[emptySpace].size - blocks[i].size}
			blocks = slices.Insert(blocks, emptySpace+1, sliced)
			i += 1
			blocks[emptySpace].size = blocks[i].size

		}
		blocks = util.Swap(blocks, emptySpace, i)
		if !part1 {
			var rem int
			blocks, rem = compressEmptySpaces(blocks, emptySpace+1)
			i -= rem
		}
	}
	return blocks
}

func compressEmptySpaces(blocks []Block, limit int) ([]Block, int) {
	compressions := 0

	for i, v := range blocks {
		if i == 0 {
			continue
		}
		if i > limit {
			break
		}

		if v.id == -1 && blocks[i-1].id == -1 {
			compressions += 1
			blocks[i-1].size += v.size
			blocks = util.Remove(blocks, i)
		}
	}

	return blocks, compressions
}

func findEmpty(blocks []Block, space int) int {
	for i, block := range blocks {
		if block.id == -1 && block.size >= space {
			return i
		}
	}
	return -1
}

func checksum(blocks []Block) int {
	disk := make([]int, 0)

	for _, block := range blocks {
		for i := 0; i < block.size; i += 1 {
			disk = append(disk, block.id)
		}
	}
	sum := 0
	for i, v := range disk {
		if v == -1 {
			continue
		}
		sum += i * v
	}
	return sum
}

func parseInput(s *bufio.Scanner) string {
	s.Scan()
	return s.Text()
}
