package day_20

import (
	"fmt"
	"strings"
)

// --- Day 20: Race Condition ---
// The Historians are quite pixelated again. This time, a massive, black building looms over you - you're right outside the CPU!
//
// While The Historians get to work, a nearby program sees that you're idle and challenges you to a race.
// Apparently, you've arrived just in time for the frequently-held race condition festival!
//
// The race takes place on a particularly long and twisting code path; programs compete to see who can finish in the fewest picoseconds.
// The winner even gets their very own mutex!
//
// They hand you a map of the racetrack (your puzzle input). For example:
//
// ###############
// #...#...#.....#
// #.#.#.#.#.###.#
// #S#...#.#.#...#
// #######.#.#.###
// #######.#.#...#
// #######.#.###.#
// ###..E#...#...#
// ###.#######.###
// #...###...#...#
// #.#####.#.###.#
// #.#...#.#.#...#
// #.#.#.#.#.#.###
// #...#...#...###
// ###############
// The map consists of track (.) - including the start (S) and end (E) positions (both of which also count as track) - and walls (#).
//
// When a program runs through the racetrack, it starts at the start position. Then, it is allowed to move up, down, left, or right;
// each such move takes 1 picosecond. The goal is to reach the end position as quickly as possible. In this example racetrack,
// the fastest time is 84 picoseconds.
//
// Because there is only a single path from the start to the end and the programs all go the same speed,
// the races used to be pretty boring. To make things more interesting,
// they introduced a new rule to the races: programs are allowed to cheat.
//
// The rules for cheating are very strict. Exactly once during a race, a program may disable collision for up to 2 picoseconds.
// This allows the program to pass through walls as if they were regular track. At the end of the cheat,
// the program must be back on normal track again; otherwise, it will receive a segmentation fault and get disqualified.
//
// So, a program could complete the course in 72 picoseconds (saving 12 picoseconds) by cheating for the two moves marked 1 and 2:
//
// ###############
// #...#...12....#
// #.#.#.#.#.###.#
// #S#...#.#.#...#
// #######.#.#.###
// #######.#.#...#
// #######.#.###.#
// ###..E#...#...#
// ###.#######.###
// #...###...#...#
// #.#####.#.###.#
// #.#...#.#.#...#
// #.#.#.#.#.#.###
// #...#...#...###
// ###############
// Or, a program could complete the course in 64 picoseconds (saving 20 picoseconds) by cheating for the two moves marked 1 and 2:
//
// ###############
// #...#...#.....#
// #.#.#.#.#.###.#
// #S#...#.#.#...#
// #######.#.#.###
// #######.#.#...#
// #######.#.###.#
// ###..E#...12..#
// ###.#######.###
// #...###...#...#
// #.#####.#.###.#
// #.#...#.#.#...#
// #.#.#.#.#.#.###
// #...#...#...###
// ###############
// This cheat saves 38 picoseconds:
//
// ###############
// #...#...#.....#
// #.#.#.#.#.###.#
// #S#...#.#.#...#
// #######.#.#.###
// #######.#.#...#
// #######.#.###.#
// ###..E#...#...#
// ###.####1##.###
// #...###.2.#...#
// #.#####.#.###.#
// #.#...#.#.#...#
// #.#.#.#.#.#.###
// #...#...#...###
// ###############
// This cheat saves 64 picoseconds and takes the program directly to the end:
//
// ###############
// #...#...#.....#
// #.#.#.#.#.###.#
// #S#...#.#.#...#
// #######.#.#.###
// #######.#.#...#
// #######.#.###.#
// ###..21...#...#
// ###.#######.###
// #...###...#...#
// #.#####.#.###.#
// #.#...#.#.#...#
// #.#.#.#.#.#.###
// #...#...#...###
// ###############
// Each cheat has a distinct start position (the position where the cheat is activated,
// just before the first move that is allowed to go through walls) and end position;
// cheats are uniquely identified by their start position and end position.
//
// In this example, the total number of cheats (grouped by the amount of time they save) are as follows:
//
// There are 14 cheats that save 2 picoseconds.
// There are 14 cheats that save 4 picoseconds.
// There are 2 cheats that save 6 picoseconds.
// There are 4 cheats that save 8 picoseconds.
// There are 2 cheats that save 10 picoseconds.
// There are 3 cheats that save 12 picoseconds.
// There is one cheat that saves 20 picoseconds.
// There is one cheat that saves 36 picoseconds.
// There is one cheat that saves 38 picoseconds.
// There is one cheat that saves 40 picoseconds.
// There is one cheat that saves 64 picoseconds.
// You aren't sure what the conditions of the racetrack will be like, so to give yourself as many options as possible,
// you'll need a list of the best cheats. How many cheats would save you at least 100 picoseconds?

func Part1(input []byte) {
	matrix, start, finish := parse(string(input))
	route := findRoute(matrix, start, finish)

	// iterate through route tiles, and for each look up, right, down, left
	// and check if there is shortcut available ( another tile ). This shortcut
	// should be placed from current tile over one, so if im at 1:1, then target
	// might be at 1:3, or 3:1, etc.
	shortcuts := []Shortcut{}
	shortcutCost := 2
	for _, current := range route {
		up, canVisit := current.Up(matrix, shortcutCost)
		if canVisit && up.score > current.score-shortcutCost {
			shortcuts = append(
				shortcuts,
				Shortcut{start: current, finish: up, save: up.score - current.score - shortcutCost},
			)
		}
		down, canVisit := current.Down(matrix, shortcutCost)
		if canVisit && down.score > current.score-shortcutCost {
			shortcuts = append(
				shortcuts,
				Shortcut{
					start:  current,
					finish: down,
					save:   down.score - current.score - shortcutCost,
				},
			)
		}
		left, canVisit := current.Left(matrix, shortcutCost)
		if canVisit && left.score > current.score-shortcutCost {
			shortcuts = append(
				shortcuts,
				Shortcut{
					start:  current,
					finish: left,
					save:   left.score - current.score - shortcutCost,
				},
			)
		}
		right, canVisit := current.Right(matrix, shortcutCost)
		if canVisit && right.score > current.score-shortcutCost {
			shortcuts = append(
				shortcuts,
				Shortcut{
					start:  current,
					finish: right,
					save:   right.score - current.score - shortcutCost,
				},
			)
		}
	}

	fmt.Printf("Found %d shortcuts\n", len(shortcuts))
	res := 0
	for _, shortcut := range shortcuts {
		fmt.Printf(
			"%s=>%s ( saved %d )\n",
			shortcut.start.String(),
			shortcut.finish.String(),
			shortcut.save,
		)
		if shortcut.save >= 100 {
			res++
		}
	}
	fmt.Printf("Result is %d\n", res)
}

type Shortcut struct {
	start, finish *Tile
	save          int
}

func (s Shortcut) String() string {
	return fmt.Sprintf("%s=>%s ( saved %d )", s.start.String(), s.finish.String(), s.save)
}

func findRoute(matrix [][]*Tile, start, finish *Tile) []*Tile {
	q := Queue[*Tile]{start}
	route := []*Tile{}

	for score := 0; q.HasItems(); score++ {
		// time.Sleep(10 * time.Millisecond)
		t := q.Pop()
		// prettyPrint(matrix, t)
		fmt.Println(t.String())

		route = append(route, t)
		matrix[t.row][t.column].score = score
		t.score = score

		if t.row == finish.row && t.column == finish.column {
			fmt.Println("Found route in", score, "steps")
			return route
		}

		up, canVisit := t.Up(matrix)
		if canVisit && up.score < 0 {
			q.Push(up)
			continue
		}
		down, canVisit := t.Down(matrix)
		if canVisit && down.score < 0 {
			q.Push(down)
			continue
		}
		left, canVisit := t.Left(matrix)
		if canVisit && left.score < 0 {
			q.Push(left)
			continue
		}
		right, canVisit := t.Right(matrix)
		if canVisit && right.score < 0 {
			q.Push(right)
			continue
		}
	}
	panic("no route found")
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

type Tile struct {
	row, column int
	score       int
	r           rune
}

func (t *Tile) DistanceTo(other *Tile) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	return abs(t.row-other.row) + abs(t.column-other.column)
}

func (t Tile) String() string {
	return fmt.Sprintf("{%d:%d}", t.row, t.column)
}

func (t *Tile) Up(matrix [][]*Tile, steps ...int) (newTile *Tile, isPartOfPath bool) {
	step := 1
	if len(steps) > 0 {
		step = steps[0]
	}
	rowDiff := -1 * step

	if t.row+rowDiff < 0 {
		return nil, false
	}
	res := matrix[t.row+rowDiff][t.column]
	return res, res.r != '#'
}

func (t *Tile) Down(matrix [][]*Tile, steps ...int) (newTile *Tile, isPartOfPath bool) {
	step := 1
	if len(steps) > 0 {
		step = steps[0]
	}
	rowDiff := step * 1

	if t.row+rowDiff >= len(matrix) {
		return nil, false
	}
	res := matrix[t.row+rowDiff][t.column]
	return res, res.r != '#'
}

func (t *Tile) Left(matrix [][]*Tile, steps ...int) (newTile *Tile, isPartOfPath bool) {
	step := 1
	if len(steps) > 0 {
		step = steps[0]
	}
	columnDiff := step * -1
	if t.column+columnDiff < 0 {
		return nil, false
	}
	res := matrix[t.row][t.column+columnDiff]
	return res, res.r != '#'
}

func (t *Tile) Right(matrix [][]*Tile, steps ...int) (newTile *Tile, isPartOfPath bool) {
	step := 1
	if len(steps) > 0 {
		step = steps[0]
	}
	columnDiff := step * 1
	if t.column+columnDiff >= len(matrix[0]) {
		return nil, false
	}
	res := matrix[t.row][t.column+columnDiff]
	return res, res.r != '#'
}

func prettyPrint(matrix [][]*Tile, focus *Tile) {
	fmt.Print("\033[H\033[2J")
	sb := strings.Builder{}
	for row, line := range matrix {
		for column, t := range line {
			if focus != nil && row == focus.row && column == focus.column {
				sb.WriteString("\033[31m" + string(t.r) + "\033[0m")
			} else {
				sb.WriteString(string(t.r))
			}
		}
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())
}

func parse(input string) (matrix [][]*Tile, start, finish *Tile) {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	matrix = make([][]*Tile, len(lines))
	for row := range lines {
		matrix[row] = make([]*Tile, len(lines[row]))
		for column, r := range lines[row] {
			matrix[row][column] = &Tile{row: row, column: column, r: r, score: -1}
			if lines[row][column] == 'S' {
				start = &Tile{row: row, column: column, r: r, score: -1}
			}
			if lines[row][column] == 'E' {
				finish = &Tile{row: row, column: column, r: r, score: -1}
			}
		}
	}
	return matrix, start, finish
}
