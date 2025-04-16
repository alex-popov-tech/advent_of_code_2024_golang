package day_8

import (
	"fmt"
	"os"
	"strings"
)

// --- Part Two ---
// Watching over your shoulder as you work, one of The Historians asks if you took the effects of resonant harmonics into your calculations.
//
// Whoops!
//
// After updating your model, it turns out that an antinode occurs at any grid position exactly in line with at least two antennas of the same frequency, regardless of distance. This means that some of the new antinodes will occur at the position of each antenna (unless that antenna is the only one of its frequency).
//
// So, these three T-frequency antennas now create many antinodes:
//
// T....#....
// ...T......
// .T....#...
// .........#
// ..#.......
// ..........
// ...#......
// ..........
// ....#.....
// ..........
// In fact, the three T-frequency antennas are all exactly in line with two antennas, so they are all also antinodes! This brings the total number of antinodes in the above example to 9.
//
// The original example now has 34 antinodes, including the antinodes that appear on every antenna:
//
// ##....#....#
// .#.#....0...
// ..#.#0....#.
// ..##...0....
// ....0....#..
// .#...#A....#
// ...#..#.....
// #....#.#....
// ..#.....A...
// ....#....A..
// .#........#.
// ...#......##
// Calculate the impact of the signal using this updated model. How many unique locations within the bounds of the map contain an antinode?

func Part2(inputPath string) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	inputStr := strings.Trim(string(input), "\n")
	lines := strings.Split(inputStr, "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	prettyPrintMatrix(matrix)

	// 1. Find all current antennas locations
	freqToPos := parse(matrix)

	// 2. Iterate over each frequency and location, and make list of unique pairs of antinodes
	freqToAntenaPairs := map[rune][]Pair[Pos]{}
	for k, v := range freqToPos {
		freqToAntenaPairs[k] = generateUniquePairs(v)
	}

	antinodes := map[Pos]struct{}{}
	for _, positions := range freqToPos {
		if len(positions) < 2 {
			continue // Need at least two antennas for antinodes
		}
		pairs := generateUniquePairs(positions)
		for _, pair := range pairs {
			linePositions := generateLinePositions(pair.f, pair.s, matrix)
			for _, pos := range linePositions {
				antinodes[pos] = struct{}{}
			}
		}
	}

	fmt.Println("Result is", len(antinodes))

	for antinode := range antinodes {
		if matrix[antinode.row][antinode.column] == '.' {
			matrix[antinode.row][antinode.column] = '#'
		}
	}

	prettyPrintMatrix(matrix)
	for _, row := range matrix {
		fmt.Println(string(row))
	}
	fmt.Println("Result is", len(antinodes))
}

func generateLinePositions(a, b Pos, matrix [][]rune) []Pos {
	dx := b.row - a.row
	dy := b.column - a.column

	g := gcd(dx, dy)
	if g == 0 {
		return []Pos{a} // same position, but pairs are unique so shouldn't happen
	}

	sx := dx / g
	sy := dy / g

	var positions []Pos

	// Generate in both directions from 'a'
	for k := 0; ; k++ {
		x := a.row + k*sx
		y := a.column + k*sy
		pos := Pos{row: x, column: y}
		if !isValid(pos, matrix) {
			break
		}
		positions = append(positions, pos)
	}

	for k := -1; ; k-- {
		x := a.row + k*sx
		y := a.column + k*sy
		pos := Pos{row: x, column: y}
		if !isValid(pos, matrix) {
			break
		}
		positions = append(positions, pos)
	}

	return positions
}

func gcd(f, s int) int {
	ff := abs(f)
	sf := abs(f)
	for curr := min(ff, sf); curr > 1; curr-- {
		if f%int(curr) == 0 && s%int(curr) == 0 {
			return int(curr)
		}
	}
	return 1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(f, s int) int {
	if f < s {
		return f
	}
	return s
}
