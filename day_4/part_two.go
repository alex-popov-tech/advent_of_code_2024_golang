package day_4

import (
	"fmt"
	"strings"
)

// The Elf looks quizzically at you. Did you misunderstand the assignment?
//
// Looking for the instructions, you flip over the word search to find that this isn't actually an XMAS puzzle; it's an X-MAS puzzle in which you're supposed to find two MAS in the shape of an X. One way to achieve that is like this:
//
// M.S
// .A.
// M.S
// Irrelevant characters have again been replaced with . in the above diagram. Within the X, each MAS can be written forwards or backwards.
//
// Here's the same example from before, but this time all of the X-MASes have been kept instead:
//
// .M.S......
// ..A..MSMS.
// .M.S.MAA..
// ..A.ASMSM.
// .M.S.M....
// ..........
// S.S.S.S.S.
// .A.A.A.A..
// M.M.M.M.M.
// ..........
// In this example, an X-MAS appears 9 times.
//
// Flip the word search from the instructions back over to the word search side and try again. How many times does an X-MAS appear?

var masMapping = map[string]bool{"MAS": true, "SAM": true}

func Part2(input []byte) {
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	count := 0
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if y+2 >= len(matrix) || x+2 >= len(matrix[y]) {
				continue
			}
			if isXmas(y, x, matrix) {
				count++
			}
		}
	}
	fmt.Println("Result is ", count)
}

func isXmas(y, x int, matrix [][]rune) bool {
	first := string([]rune{matrix[y][x], matrix[y+1][x+1], matrix[y+2][x+2]})
	second := string([]rune{matrix[y+2][x], matrix[y+1][x+1], matrix[y][x+2]})
	return masMapping[first] && masMapping[second]
}
