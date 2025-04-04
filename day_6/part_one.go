package day_6

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// --- Day 6: Guard Gallivant ---
// The Historians use their fancy device again, this time to whisk you all away to the North Pole prototype suit manufacturing lab... in the year 1518! It turns out that having direct access to history is very convenient for a group of historians.
//
// You still have to be careful of time paradoxes, and so it will be important to avoid anyone from 1518 while The Historians search for the Chief. Unfortunately, a single guard is patrolling this part of the lab.
//
// Maybe you can work out where the guard will go ahead of time so that The Historians can search safely?
//
// You start by making a map (your puzzle input) of the situation. For example:
//
// ....#.....
// .........#
// ..........
// ..#.......
// .......#..
// ..........
// .#..^.....
// ........#.
// #.........
// ......#...
// The map shows the current position of the guard with ^ (to indicate the guard is currently facing up from the perspective of the map). Any obstructions - crates, desks, alchemical reactors, etc. - are shown as #.
//
// Lab guards in 1518 follow a very strict patrol protocol which involves repeatedly following these steps:
//
// If there is something directly in front of you, turn right 90 degrees.
// Otherwise, take a step forward.
// Following the above protocol, the guard moves up several times until she reaches an obstacle (in this case, a pile of failed suit prototypes):
//
// ....#.....
// ....^....#
// ..........
// ..#.......
// .......#..
// ..........
// .#........
// ........#.
// #.........
// ......#...
// Because there is now an obstacle in front of the guard, she turns right before continuing straight in her new facing direction:
//
// ....#.....
// ........>#
// ..........
// ..#.......
// .......#..
// ..........
// .#........
// ........#.
// #.........
// ......#...
// Reaching another obstacle (a spool of several very long polymers), she turns right again and continues downward:
//
// ....#.....
// .........#
// ..........
// ..#.......
// .......#..
// ..........
// .#......v.
// ........#.
// #.........
// ......#...
// This process continues for a while, but the guard eventually leaves the mapped area (after walking past a tank of universal solvent):
//
// ....#.....
// .........#
// ..........
// ..#.......
// .......#..
// ..........
// .#........
// ........#.
// #.........
// ......#v..
// By predicting the guard's route, you can determine which specific positions in the lab will be in the patrol path. Including the guard's starting position, the positions visited by the guard before leaving the area are marked with an X:
//
// ....#.....
// ....XXXXX#
// ....X...X.
// ..#.X...X.
// ..XXXXX#X.
// ..X.X.X.X.
// .#XXXXXXX.
// .XXXXXXX#.
// #XXXXXXX..
// ......#X..
// In this example, the guard will visit 41 distinct positions on your map.
//
// Predict the path of the guard. How many distinct positions will the guard visit before leaving the mapped area?

func Part1(inputPath string) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	inputStr := strings.Trim(string(input), "\n")
	lines := strings.Split(inputStr, "\n")
	m := NewMatrix(lines)
	currRow, currColumn := m.getGuardPosition()
	currDirection := NewDirection(m.peek(currRow, currColumn))
	for {
		fmt.Println("==================")
		m.print()
		// 1. peek next
		next, ok := m.peekNext(currRow, currColumn, currDirection)
		// 1.1 if next if end of the map, visit current and print results
		if !ok {
			m.visit(currRow, currColumn)
			fmt.Println("===============")
			fmt.Println("END OF GAME")
			fmt.Println("===============")
			fmt.Println("Visited", m.countVisited())
			break
		}
		// 1.2 if wall - change direction and continue
		if next == WALL {
			currDirection = currDirection.TurnRight()
			continue
		}
		// 1.3 if empty:
		if next == EMPTY || next == VISITED {
			// 1.3.1 visit current
			m.visit(currRow, currColumn)
			// 1.3.2 put guard to next cell
			nextRow, nextColumn := nextPosition(currRow, currColumn, currDirection)
			m.putGuard(nextRow, nextColumn, currDirection)
			// 1.3.3 update guard position
			currRow, currColumn = nextRow, nextColumn
			continue
		}
		log.Fatal("next is not handled properly")
	}
}

type Direction rune

const (
	UP    Direction = '^'
	DOWN  Direction = 'v'
	RIGHT Direction = '>'
	LEFT  Direction = '<'
)

func (d Direction) String() string {
	return map[Direction]string{
		UP:    "UP",
		DOWN:  "DOWN",
		RIGHT: "RIGHT",
		LEFT:  "LEFT",
	}[d]
}

func (d Direction) TurnRight() Direction {
	switch d {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	default:
		log.Fatal("Cannot get 'right' direction from '%v'", d)
		return 'F'
	}
}

func NewDirection(r rune) Direction {
	switch r {
	case '^':
		return UP
	case '>':
		return RIGHT
	case 'v':
		return DOWN
	case '<':
		return LEFT
	default:
		log.Fatal("You tried to convert rune '%v' to 'Direction', which is incorrect\n", r)
		return 'F'
	}
}

type CellType rune

const (
	EMPTY   CellType = '.'
	WALL    CellType = '#'
	VISITED CellType = 'X'
)

func (d CellType) String() string {
	return map[CellType]string{
		EMPTY:   "EMPTY",
		WALL:    "WALL",
		VISITED: "VISITED",
	}[d]
}

func NewCellType(r rune) CellType {
	switch r {
	case '.':
		return EMPTY
	case '#':
		return WALL
	case 'X':
		return VISITED
	default:
		log.Fatal("You tried to convert rune '%v' to CellType, which is incorrect\n", r)
		return 'F'
	}
}

type Matrix struct {
	matrix [][]rune
}

func NewMatrix(lines []string) Matrix {
	runes := make([][]rune, len(lines))
	for lineIndex, line := range lines {
		runes[lineIndex] = []rune{}
		for _, r := range line {
			runes[lineIndex] = append(runes[lineIndex], r)
		}
	}
	return Matrix{runes}
}

func isValid[T any](row, column int, matrix [][]T) bool {
	if row < 0 || column < 0 || row >= len(matrix) || column >= len(matrix[row]) {
		return false
	}
	return true
}

func mustValid[T any](row, column int, matrix [][]T) {
	if !isValid(row, column, matrix) {
		log.Fatalf(
			"Tried to access row:%d(len %d),column:%d(len %d)\n",
			row,
			len(matrix[row]),
			column,
			len(matrix[row]),
		)
	}
}

func (m Matrix) print() {
	for _, rline := range m.matrix {
		fmt.Println(string(rline))
	}
}

func (m *Matrix) visit(row, column int) {
	mustValid(row, column, m.matrix)
	m.matrix[row][column] = rune(VISITED)
}

func (m *Matrix) putGuard(row, column int, direction Direction) {
	mustValid(row, column, m.matrix)
	m.matrix[row][column] = rune(direction)
}

func (m Matrix) peek(row, column int) rune {
	mustValid(row, column, m.matrix)
	return m.matrix[row][column]
}

func nextPosition(row, column int, direction Direction) (int, int) {
	switch direction {
	case UP:
		return row - 1, column
	case DOWN:
		return row + 1, column
	case LEFT:
		return row, column - 1
	case RIGHT:
		return row, column + 1
	default:
		return -1, -1
	}
}

func (m Matrix) peekNext(currentRow, currentColumn int, direction Direction) (CellType, bool) {
	mustValid(currentRow, currentColumn, m.matrix)
	nextRow, nextColumn := nextPosition(currentRow, currentColumn, direction)
	fmt.Println("peekNext", nextRow, nextColumn, "isValid", isValid(nextRow, nextColumn, m.matrix))
	if isValid(nextRow, nextColumn, m.matrix) {
		return NewCellType(m.matrix[nextRow][nextColumn]), true
	}
	return 'F', false
}

func isGuard(r rune) bool {
	switch r {
	case '.':
		return false
	case 'X':
		return false
	case '#':
		return false
	default:
		return true
	}
}

func (m Matrix) getGuardPosition() (row, column int) {
	for rowIndex, line := range m.matrix {
		for columnIndex, r := range line {
			if isGuard(r) {
				return rowIndex, columnIndex
			}
		}
	}
	log.Fatalf("Cannot find guard")
	return -1, -1
}

func (m Matrix) countVisited() int {
	count := 0
	for _, line := range m.matrix {
		for _, r := range line {
			if r == 'X' {
				count++
			}
		}
	}
	return count
}
