package day_8

import (
	"fmt"
	"strings"
)

// --- Day 8: Resonant Collinearity ---
// You find yourselves on the roof of a top-secret Easter Bunny installation.
//
// While The Historians do their thing, you take a look at the familiar huge antenna. Much to your surprise, it seems to have been reconfigured to emit a signal that makes people 0.1% more likely to buy Easter Bunny brand Imitation Mediocre Chocolate as a Christmas gift! Unthinkable!
//
// Scanning across the city, you find that there are actually many such antennas. Each antenna is tuned to a specific frequency indicated by a single lowercase letter, uppercase letter, or digit. You create a map (your puzzle input) of these antennas. For example:
//
// ............
// ........0...
// .....0......
// .......0....
// ....0.......
// ......A.....
// ............
// ............
// ........A...
// .........A..
// ............
// ............
// The signal only applies its nefarious effect at specific antinodes based on the resonant frequencies of the antennas. In particular, an antinode occurs at any point that is perfectly in line with two antennas of the same frequency - but only when one of the antennas is twice as far away as the other. This means that for any pair of antennas with the same frequency, there are two antinodes, one on either side of them.
//
// So, for these two antennas with frequency a, they create the two antinodes marked with #:
//
// ..........
// ...#......
// ..........
// ....a.....
// ..........
// .....a....
// ..........
// ......#...
// ..........
// ..........
// Adding a third antenna with the same frequency creates several more antinodes. It would ideally add four antinodes, but two are off the right side of the map, so instead it adds only two:
//
// ..........
// ...#......
// #.........
// ....a.....
// ........a.
// .....a....
// ..#.......
// ......#...
// ..........
// ..........
// Antennas with different frequencies don't create antinodes; A and a count as different frequencies. However, antinodes can occur at locations that contain antennas. In this diagram, the lone antenna with frequency capital A creates no antinodes but has a lowercase-a-frequency antinode at its location:
//
// ..........
// ...#......
// #.........
// ....a.....
// ........a.
// .....a....
// ..#.......
// ......A...
// ..........
// ..........
// The first example has antennas with two different frequencies, so the antinodes they create look like this, plus an antinode overlapping the topmost A-frequency antenna:
//
// ......#....#
// ...#....0...
// ....#0....#.
// ..#....0....
// ....0....#..
// .#....A.....
// ...#........
// #......#....
// ........A...
// .........A..
// ..........#.
// ..........#.
// Because the topmost A-frequency antenna overlaps with a 0-frequency antinode, there are 14 total unique locations that contain an antinode within the bounds of the map.
//
// Calculate the impact of the signal. How many unique locations within the bounds of the map contain an antinode?

func Part1(input []byte) {
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

	antinodesToAntenas := map[Pos]Pair[Pos]{}
	for r, pairs := range freqToAntenaPairs {
		fmt.Println("================================")
		fmt.Println(string(r))
		for _, pair := range pairs {
			fmt.Println(pair)
			fnext := generateNextPos(pair.f, pair.s)
			if isValid(fnext, matrix) {
				fmt.Println("fnext", fnext)
				antinodesToAntenas[fnext] = pair
			}
			snext := generateNextPos(pair.s, pair.f)
			if isValid(snext, matrix) {
				fmt.Println("snext", snext)
				antinodesToAntenas[snext] = pair
			}
		}
		fmt.Println("================================")
	}

	for antinode, parents := range antinodesToAntenas {
		fmt.Println(parents, antinode)
		if matrix[antinode.row][antinode.column] == '.' {
			matrix[antinode.row][antinode.column] = '#'
		}
	}

	prettyPrintMatrix(matrix)
	fmt.Println("Result is", len(antinodesToAntenas))
}

type Pos struct {
	row    int
	column int
}

type Pair[T any] struct {
	f T
	s T
}

func parse(lines [][]rune) map[rune][]Pos {
	result := map[rune][]Pos{}
	for row, line := range lines {
		for column, r := range line {
			if r != '.' {
				result[r] = append(result[r], Pos{row: row, column: column})
			}
		}
	}
	return result
}

func generateUniquePairs[T comparable](input []T) []Pair[T] {
	if len(input) < 2 {
		panic("Cannot generate unique pairs of len(input) < 2")
	}
	result := []Pair[T]{}
	for f := range input {
		for s := f + 1; s < len(input); s++ {
			result = append(result, Pair[T]{f: input[f], s: input[s]})
		}
	}
	return result
}

func generateNextPos(f, s Pos) Pos {
	rowdiff := f.row - s.row
	columndiff := f.column - s.column
	res := Pos{row: f.row + rowdiff, column: f.column + columndiff}
	return res
}

func isValid(pos Pos, matrix [][]rune) bool {
	return pos.row >= 0 && pos.column >= 0 && len(matrix) > pos.row &&
		len(matrix[pos.row]) > pos.column
}

func prettyPrintMatrix(matrix [][]rune) {
	res := []string{}
	res = append(
		res,
		fmt.Sprintf(
			"     %s",
			strings.Join(
				[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
				" | ",
			),
		),
	)
	for rowIndex, row := range matrix {
		res = append(
			res,
			fmt.Sprintf(
				"%02d | %s",
				rowIndex,
				strings.Join(strings.Split(string(row), ""), " | "),
			),
		)
	}
	fmt.Println(strings.Join(res, "\n"))
}
