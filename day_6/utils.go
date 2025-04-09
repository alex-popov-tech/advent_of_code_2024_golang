package day_6

import (
	"fmt"
	"slices"
	"strings"
)

type Cell rune

const (
	wall    Cell = '#'
	empty   Cell = '.'
	visited Cell = 'x'
)

const (
	rUp    rune = '^'
	rDown  rune = 'v'
	rRight rune = '>'
	rLeft  rune = '<'
)

type Position struct {
	row    int
	column int
}

type VisitedPosition struct {
	row    int
	column int
	dir    Direction
}

type Direction struct {
	r rune
}

func NewDirection(r rune) Direction {
	if slices.Contains([]rune{rUp, rDown, rRight, rLeft}, r) {
		return Direction{r}
	}
	panic("Trying create 'Direction' out of invalid rune '" + string(r) + "'")
}

func (d Direction) String() string {
	switch d.r {
	case '^':
		return "up"
	case '>':
		return "right"
	case '<':
		return "left"
	case 'v':
		return "bottom"
	default:
		return "INVALID"
	}
}

func linesToRunes(lines []string) [][]rune {
	runes := make([][]rune, len(lines))
	for lineIndex, line := range lines {
		runes[lineIndex] = []rune{}
		for _, r := range line {
			runes[lineIndex] = append(runes[lineIndex], r)
		}
	}
	return runes
}

func findGuard(matrix [][]rune) (Position, Direction) {
	for rowIndex, line := range matrix {
		for columnIndex, r := range line {
			if isGuard(r) {
				return Position{row: rowIndex, column: columnIndex}, NewDirection(r)
			}
		}
	}
	return Position{-1, -1}, NewDirection('^')
}

func getFollowingPosition(row, column int, dir Direction) (newRow, newColumn int) {
	newRow, newColumn = row, column
	switch dir.r {
	case rUp:
		newRow -= 1
	case rDown:
		newRow += 1
	case rLeft:
		newColumn -= 1
	case rRight:
		newColumn += 1
	default:
		panic(fmt.Sprintf("Passed invalid rune to getFollowingPosition('%s')\n", dir))
	}
	return newRow, newColumn
}

func isInBounds(matrix [][]rune, row, column int) bool {
	if row < 0 || column < 0 {
		return false
	}
	if row >= len(matrix) || column >= len(matrix[row]) {
		return false
	}
	return true
}

func prettyMatrix(matrix [][]rune) string {
	res := []string{strings.Join(strings.Split(" 0123456789", ""), " | ")}
	for rowIndex, row := range matrix {
		res = append(
			res,
			fmt.Sprintf(
				"%d | %s",
				rowIndex,
				strings.Join(strings.Split(string(row), ""), " | "),
			),
		)
	}
	return strings.Join(res, "\n")
}

func rightFrom(dir Direction) Direction {
	switch dir.r {
	case rUp:
		return NewDirection(rRight)
	case rRight:
		return NewDirection(rDown)
	case rDown:
		return NewDirection(rLeft)
	case rLeft:
		return NewDirection(rUp)
	default:
		panic(fmt.Sprintf("Cannot find left direction from '%s'", string(dir.r)))
	}
}
