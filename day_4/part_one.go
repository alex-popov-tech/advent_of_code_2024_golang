package day_4

import (
	"fmt"
	"os"
	"strings"
)

// --- Day 4: Ceres Search ---
// "Looks like the Chief's not here. Next!" One of The Historians pulls out a device and pushes the only button on it. After a brief flash, you recognize the interior of the Ceres monitoring station!

// As the search for the Chief continues, a small Elf who lives on the station tugs on your shirt; she'd like to know if you could help her with her word search (your puzzle input). She only has to find one word: XMAS.

// This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words. It's a little unusual, though, as you don't merely need to find one instance of XMAS - you need to find all of them. Here are a few ways XMAS might appear, where irrelevant characters have been replaced with .:

// ..X...
// .SAMX.
// .A..A.
// XMAS.S
// .X....
// The actual word search will be full of letters instead. For example:

// MMMSXXMASM
// MSAMXMSMSA
// AMXSXMAAMM
// MSAMASMSMX
// XMASAMXAMM
// XXAMMXXAMA
// SMSMSASXSS
// SAXAMASAAA
// MAMMMXMMMM
// MXMXAXMASX
// In this word search, XMAS occurs a total of 18 times; here's the same word search again, but where letters not involved in any XMAS have been replaced with .:

// ....XXMAS.
// .SAMXMS...
// ...S..A...
// ..A.A.MS.X
// XMASAMX.MM
// X.....XA.A
// S.S.S.S.SS
// .A.A.A.A.A
// ..M.M.M.MM
// .X.X.XMASX
// Take a look at the little Elf's word search. How many times does XMAS appear?

var xmas = []rune{'X', 'M', 'A', 'S'}

func Part1(inputPath string) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	count := 0
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			fmt.Printf("y:%d,x:%d,rune:%s\n", y, x, string(matrix[y][x]))
			if matrix[y][x] != 'X' {
				continue
			}
			if isTopLeft(y, x, matrix) {
				count++
			}
			if isTop(y, x, matrix) {
				count++
			}
			if isTopRight(y, x, matrix) {
				count++
			}
			if isRight(y, x, matrix) {
				count++
			}
			if isBottomRight(y, x, matrix) {
				count++
			}
			if isBottom(y, x, matrix) {
				count++
			}
			if isBottomLeft(y, x, matrix) {
				count++
			}
			if isLeft(y, x, matrix) {
				count++
			}
			fmt.Println("count is", count)
		}
	}
	fmt.Println("Result is ", count)
}

func isTopLeft(y, x int, matrix [][]rune) bool {
	if y < 3 || x < 3 {
		return false
	}
	for i, step := 0, 0; step < 4; i, step = i+1, step+1 {
		if matrix[y-step][x-step] != xmas[i] {
			return false
		}
	}
	return true
}

func isTop(y, x int, matrix [][]rune) bool {
	if y < 3 {
		return false
	}
	for i, step := 0, 0; step < 4; i, step = i+1, step+1 {
		if matrix[y-step][x] != xmas[i] {
			return false
		}
	}
	return true
}

func isTopRight(y, x int, matrix [][]rune) bool {
	if y < 3 || x >= (len(matrix[y])-3) {
		return false
	}
	for i, step := 0, 0; step < 4; i, step = i+1, step+1 {
		if matrix[y-step][x+step] != xmas[i] {
			return false
		}
	}
	return true
}

func isRight(y, x int, matrix [][]rune) bool {
	if x >= (len(matrix[y]) - 3) {
		return false
	}
	for i, step := 0, 0; step < 4; i, step = i+1, step+1 {
		if matrix[y][x+step] != xmas[i] {
			return false
		}
	}
	return true
}

func isBottomRight(y, x int, matrix [][]rune) bool {
	if y+3 >= len(matrix) || x+3 >= len(matrix[y]) {
		return false
	}
	for i, step := 0, 0; step < 4; i, step = i+1, step+1 {
		fmt.Printf("INSIDE y:%d,x:%d,rune:%s\n", y+step, x+step, string(matrix[y+step][x+step]))
		if matrix[y+step][x+step] != xmas[i] {
			return false
		}
	}
	return true
}

func isBottom(y, x int, matrix [][]rune) bool {
	if y >= (len(matrix) - 3) {
		return false
	}
	for i, step := 0, 0; step < 4; i, step = i+1, step+1 {
		if matrix[y+step][x] != xmas[i] {
			return false
		}
	}
	return true
}

func isBottomLeft(y, x int, matrix [][]rune) bool {
	if y >= (len(matrix)-3) || x < 3 {
		return false
	}
	for i, step := 0, 0; step < 4; i, step = i+1, step+1 {
		if matrix[y+step][x-step] != xmas[i] {
			return false
		}
	}
	return true
}

func isLeft(y, x int, matrix [][]rune) bool {
	if x < 3 {
		return false
	}
	for i, step := 0, 0; step < 4; i, step = i+1, step+1 {
		if matrix[y][x-step] != xmas[i] {
			return false
		}
	}
	return true
}
