package day_18

import (
	"fmt"
	"strconv"
	"strings"
)

// --- Day 18: RAM Run ---
// You and The Historians look a lot more pixelated than you remember. You're inside a computer at the North Pole!
//
// Just as you're about to check out your surroundings, a program runs up to you.
// "This region of memory isn't safe! The User misunderstood what a pushdown automaton is and their algorithm is pushing whole bytes down on top of us! Run!"
//
// The algorithm is fast - it's going to cause a byte to fall into your memory space once every nanosecond!
// Fortunately, you're faster, and by quickly scanning the algorithm, you create a list of which bytes will fall
// (your puzzle input) in the order they'll land in your memory space.
//
// Your memory space is a two-dimensional grid with coordinates that range from 0 to 70 both horizontally and vertically.
// However, for the sake of example, suppose you're on a smaller grid with coordinates that range from 0 to 6 and the following list of incoming byte positions:
//
// 5,4
// 4,2
// 4,5
// 3,0
// 2,1
// 6,3
// 2,4
// 1,5
// 0,6
// 3,3
// 2,6
// 5,1
// 1,2
// 5,5
// 2,5
// 6,5
// 1,4
// 0,4
// 6,4
// 1,1
// 6,1
// 1,0
// 0,5
// 1,6
// 2,0
// Each byte position is given as an X,Y coordinate, where X is the distance from the left edge of
// your memory space and Y is the distance from the top edge of your memory space.
//
// You and The Historians are currently in the top left corner of the memory space (at 0,0) and
// need to reach the exit in the bottom right corner (at 70,70 in your memory space, but at 6,6 in this example). You'll need to simulate the falling bytes to plan out where it will be safe to run; for now, simulate just the first few bytes falling into your memory space.
//
// As bytes fall into your memory space, they make that coordinate corrupted.
// Corrupted memory coordinates cannot be entered by you or The Historians, so you'll need to plan your route carefully.
// You also cannot leave the boundaries of the memory space; your only hope is to reach the exit.
//
// In the above example, if you were to draw the memory space after the first 12 bytes have fallen (using . for safe and # for corrupted),
// it would look like this:
//
// ...#...
// ..#..#.
// ....#..
// ...#..#
// ..#..#.
// .#..#..
// #.#....
// You can take steps up, down, left, or right. After just 12 bytes have corrupted locations in your memory space,
// the shortest path from the top left corner to the exit would take 22 steps. Here (marked with O) is one such path:
//
// OO.#OOO
// .O#OO#O
// .OOO#OO
// ...#OO#
// ..#OO#.
// .#.O#..
// #.#OOOO
// Simulate the first kilobyte (1024 bytes) falling onto your memory space. Afterward, what is the minimum number of steps needed to reach the exit?

const (
	// size = 7
	size        = 71
	limit       = 1024
	emptyScore  = -1
	startRow    = 0
	startColumn = 0
	endRow      = 70
	endColumn   = 70
	// endRow    = 6
	// endColumn = 6
)

func Part1(input []byte) {
	matrix := parse(string(input), size, limit)
	prettyPrint(matrix, matrix[startRow][startColumn])
	traverse(matrix, matrix[startRow][startColumn])
	fmt.Println("Result is", matrix[endRow][endColumn].score)
}

type Queue[T any] []T

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	if len(*q) == 0 {
		panic("empty queue")
	}
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *Queue[T]) HasItems() bool {
	return len(*q) > 0
}

func traverse(matrix [][]*tile, start *tile) {
	start.score = 0
	q := Queue[*tile]{start}
	for q.HasItems() {
		t := q.Pop()
		prettyPrint(matrix, t)
		if t.row == endRow && t.column == endColumn {
			return
		}

		up, isValid := t.Up(matrix)
		if isValid && up.r != '#' {
			up.isVisited = true
			up.score = t.score + 1
			q.Push(up)
		}
		down, isValid := t.Down(matrix)
		if isValid && down.r != '#' {
			down.isVisited = true
			down.score = t.score + 1
			q.Push(down)
		}
		left, isValid := t.Left(matrix)
		if isValid && left.r != '#' {
			left.isVisited = true
			left.score = t.score + 1
			q.Push(left)
		}
		right, isValid := t.Right(matrix)
		if isValid && right.r != '#' {
			right.isVisited = true
			right.score = t.score + 1
			q.Push(right)
		}
	}
}

type tile struct {
	row       int
	column    int
	r         rune
	isVisited bool
	score     int
}

func (t *tile) Up(matrix [][]*tile) (*tile, bool) {
	if t.row == 0 {
		return nil, false
	}
	res := matrix[t.row-1][t.column]
	return res, !res.isVisited
}

func (t *tile) Down(matrix [][]*tile) (*tile, bool) {
	if t.row == len(matrix)-1 {
		return nil, false
	}
	res := matrix[t.row+1][t.column]
	return res, !res.isVisited
}

func (t *tile) Left(matrix [][]*tile) (*tile, bool) {
	if t.column == 0 {
		return nil, false
	}
	res := matrix[t.row][t.column-1]
	return res, !res.isVisited
}

func (t *tile) Right(matrix [][]*tile) (*tile, bool) {
	if t.column == len(matrix[0])-1 {
		return nil, false
	}
	res := matrix[t.row][t.column+1]
	return res, !res.isVisited
}

func parse(input string, size int, wallsLimit int) [][]*tile {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")

	if wallsLimit > len(lines) {
		fmt.Println(len(lines), wallsLimit)
		panic("Limit of walls is more than amount of actual walls")
	}
	matrix := make([][]*tile, size)
	for row := range matrix {
		matrix[row] = make([]*tile, size)
		for column := range matrix[row] {
			matrix[row][column] = &tile{
				r:         '.',
				row:       row,
				column:    column,
				score:     emptyScore,
				isVisited: false,
			}
		}
	}

	for i, coord := range lines {
		if i == wallsLimit {
			break
		}
		column := atoi(strings.Split(coord, ",")[0])
		row := atoi(strings.Split(coord, ",")[1])
		matrix[row][column].r = '#'
	}
	return matrix
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func prettyPrint(matrix [][]*tile, current *tile) {
	sb := strings.Builder{}
	for row := range matrix {
		for column := range matrix[row] {
			if current != nil && row == current.row && column == current.column {
				sb.WriteString("\033[41m" + string(matrix[row][column].r) + "\033[0m")
			} else if matrix[row][column].isVisited {
				sb.WriteString("\033[33m" + string(matrix[row][column].r) + "\033[0m")
			} else {
				sb.WriteString(string(matrix[row][column].r))
			}
		}
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())
}
