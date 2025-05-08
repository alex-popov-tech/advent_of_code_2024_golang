package day_16

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"
)

// --- Part Two ---
// Now that you know what the best paths look like, you can figure out the best spot to sit.
//
// Every non-wall tile (S, ., or E) is equipped with places to sit along the edges of the tile.
// While determining which of these tiles would be the best spot to sit depends on a whole bunch
// of factors (how comfortable the seats are, how far away the bathrooms are, whether there's a pillar blocking your view, etc.),
// the most important factor is whether the tile is on one of the best paths through the maze. If you sit somewhere else, you'd miss all the action!
//
// So, you'll need to determine which tiles are part of any best path through the maze, including the S and E tiles.
//
// In the first example, there are 45 tiles (marked O) that are part of at least one of the various best paths through the maze:
//
// ###############
// #.......#....O#
// #.#.###.#.###O#
// #.....#.#...#O#
// #.###.#####.#O#
// #.#.#.......#O#
// #.#.#####.###O#
// #..OOOOOOOOO#O#
// ###O#O#####O#O#
// #OOO#O....#O#O#
// #O#O#O###.#O#O#
// #OOOOO#...#O#O#
// #O###.#.#.#O#O#
// #O..#.....#OOO#
// ###############
// In the second example, there are 64 tiles that are part of at least one of the best paths:
//
// #################
// #...#...#...#..O#
// #.#.#.#.#.#.#.#O#
// #.#.#.#...#...#O#
// #.#.#.#.###.#.#O#
// #OOO#.#.#.....#O#
// #O#O#.#.#.#####O#
// #O#O..#.#.#OOOOO#
// #O#O#####.#O###O#
// #O#O#..OOOOO#OOO#
// #O#O###O#####O###
// #O#O#OOO#..OOO#.#
// #O#O#O#####O###.#
// #O#O#OOOOOOO..#.#
// #O#O#O#########.#
// #O#OOO..........#
// #################
// Analyze your map further. How many tiles are part of at least one of the best paths through the maze?

func Part2(input []byte) {
	start, maze := parse2(string(input))

	paths := findPaths(maze, start, '>', Path{})
	slices.SortFunc(paths, func(p1, p2 Path) int {
		return p1.score - p2.score
	})

	bestScore := paths[0].score
	uniqTiles := make(map[string]struct{})
	for _, path := range paths {
		if path.score != bestScore {
			break
		}
		for _, t := range path.tiles {
			uniqTiles[fmt.Sprintf("%d,%d", t.row, t.column)] = struct{}{}
			maze[t.row][t.column].r = 'O'
		}
	}
	prettyPrint2(maze, nil)
	fmt.Println("Number of unique tiles is", len(uniqTiles))
	fmt.Println("Best score is", bestScore)
	for i, path := range paths {
		fmt.Println("Path score", i, path.score)
	}
}

type tile2 struct {
	row, column int
	r           rune
	bestScores  map[rune]int
}

func newTile2(row, column int, r rune) tile2 {
	return tile2{
		row:    row,
		column: column,
		r:      r,
		bestScores: map[rune]int{
			Up:    math.MaxInt,
			Right: math.MaxInt,
			Down:  math.MaxInt,
			Left:  math.MaxInt,
		},
	}
}

type Path struct {
	tiles []tile2
	score int
}

func findPaths(
	maze [][]tile2,
	current tile2,
	currentDir rune,
	path Path,
) []Path {
	// time.Sleep(10 * time.Millisecond)

	path.tiles = append(slices.Clone(path.tiles), current)

	if maze[current.row][current.column].r == 'E' {
		prettyPrint2(maze, toPtrs(path.tiles), "32")
		time.Sleep(200 * time.Millisecond)
		return []Path{path}
	}

	if maze[current.row][current.column].r != 'S' {
		maze[current.row][current.column].r = currentDir
		maze[current.row][current.column].bestScores[currentDir] = path.score
	}

	if len(path.tiles)%2 == 0 {
		prettyPrint2(maze, []*tile2{&current})
	}
	successfulPaths := []Path{}

	upTile := nextTile2(maze, current, Up)
	upTileProposedScore := newScore(path.score, currentDir, Up)
	if needsToVisit2(maze, path.tiles, upTile, Up, upTileProposedScore) {
		newPath := Path{score: upTileProposedScore, tiles: slices.Clone(path.tiles)}
		upPaths := findPaths(maze, upTile, Up, newPath)
		successfulPaths = append(successfulPaths, upPaths...)
	}

	rightTile := nextTile2(maze, current, Right)
	rightTileProposedScore := newScore(path.score, currentDir, Right)
	if needsToVisit2(maze, path.tiles, rightTile, Right, rightTileProposedScore) {
		newPath := Path{score: rightTileProposedScore, tiles: slices.Clone(path.tiles)}
		rightPaths := findPaths(maze, rightTile, Right, newPath)
		successfulPaths = append(successfulPaths, rightPaths...)
	}

	downTile := nextTile2(maze, current, Down)
	downTileProposedScore := newScore(path.score, currentDir, Down)
	if needsToVisit2(maze, path.tiles, downTile, Down, downTileProposedScore) {
		newPath := Path{score: downTileProposedScore, tiles: slices.Clone(path.tiles)}
		downPaths := findPaths(maze, downTile, Down, newPath)
		successfulPaths = append(successfulPaths, downPaths...)
	}

	leftTile := nextTile2(maze, current, Left)
	leftTileProposedScore := newScore(path.score, currentDir, Left)
	if needsToVisit2(maze, path.tiles, leftTile, Left, leftTileProposedScore) {
		newPath := Path{score: leftTileProposedScore, tiles: slices.Clone(path.tiles)}
		leftPaths := findPaths(maze, leftTile, Left, newPath)
		successfulPaths = append(successfulPaths, leftPaths...)
	}

	return successfulPaths
}

var (
	prev2       [][]rune
	prevFocuses []*tile2
)

func prettyPrint2(maze [][]tile2, focusTiles []*tile2, highlightColor ...string) {
	color := "32;42" // default green
	if len(highlightColor) > 0 {
		color = highlightColor[0]
	}

	if prev2 == nil {
		fmt.Print("\033[H\033[2J")
		for range maze {
			fmt.Println(strings.Repeat(" ", len(maze[0])))
		}
		prev2 = make([][]rune, len(maze))
		for i := range maze {
			prev2[i] = make([]rune, len(maze[i]))
		}
	}

	isFocus := func(r, c int) bool {
		for _, f := range focusTiles {
			if f != nil && f.row == r && f.column == c {
				return true
			}
		}
		return false
	}
	wasFocus := func(r, c int) bool {
		for _, f := range prevFocuses {
			if f != nil && f.row == r && f.column == c {
				return true
			}
		}
		return false
	}

	for r := range maze {
		for c := range maze[r] {
			curr, old := maze[r][c], prev2[r][c]
			nowF := isFocus(r, c)
			prevF := wasFocus(r, c)

			if curr.r != old || nowF || prevF {
				fmt.Printf("\033[%d;%dH\033[0m", r+1, c+1)

				switch {
				case nowF:
					fmt.Printf("\033[%sm%c\033[0m", color, curr.r)
				// case curr == 'S':
				// 	fmt.Printf("\033[33mS\033[0m") // yellow
				case curr.r == 'E' || curr.r == 'S':
					fmt.Printf("\033[31;42m%s\033[0m", string(curr.r)) // green
				case curr.r == '^' || curr.r == 'v' || curr.r == '<' || curr.r == '>':
					fmt.Printf("\033[33m%c\033[0m", curr.r) // blue
				default:
					fmt.Printf("%c", curr.r)
				}

				prev2[r][c] = curr.r
			}
		}
	}

	prevFocuses = focusTiles
	fmt.Printf("\033[%d;1H", len(maze)+1)
}

func toPtrs[T any](s []T) []*T {
	ptrs := make([]*T, len(s))
	for i := range s {
		ptrs[i] = &s[i]
	}
	return ptrs
}

func parse2(input string) (start tile2, maze [][]tile2) {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	maze = make([][]tile2, len(lines))
	for row, line := range lines {
		line = strings.Trim(line, "\n")
		maze[row] = make([]tile2, len(line))
		for column, t := range line {
			maze[row][column] = newTile2(row, column, t)
			if t == 'S' {
				start = newTile2(row, column, t)
				start.bestScores[Up] = 0
				start.bestScores[Right] = 0
				start.bestScores[Down] = 0
				start.bestScores[Left] = 0
			}
		}
	}
	return start, maze
}

func nextTile2(maze [][]tile2, current tile2, dir rune) tile2 {
	row := current.row
	column := current.column
	switch dir {
	case Up:
		row--
	case Down:
		row++
	case Left:
		column--
	case Right:
		column++
	default:
		panic("unknown direction")
	}
	return maze[row][column]
}

func needsToVisit2(
	maze [][]tile2,
	visited []tile2,
	target tile2,
	dir rune,
	proposedScore int,
) bool {
	if target.row < 0 || target.column < 0 || target.row >= len(maze) ||
		target.column >= len(maze[target.row]) {
		return false
	}
	if maze[target.row][target.column].r == '#' || maze[target.row][target.column].r == 'S' {
		return false
	}
	if slices.ContainsFunc(visited, func(t tile2) bool {
		return t.row == target.row && t.column == target.column
	}) {
		return false
	}
	return target.bestScores[dir] >= proposedScore
}
