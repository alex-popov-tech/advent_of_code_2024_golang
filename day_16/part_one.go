package day_16

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

// --- Day 16: Reindeer Maze ---
// It's time again for the Reindeer Olympics! This year, the big event is the Reindeer Maze, where the Reindeer compete for the lowest score.
//
// You and The Historians arrive to search for the Chief right as the event is about to start. It wouldn't hurt to watch a little, right?
//
// The Reindeer start on the Start Tile (marked S) facing East and need to reach the End Tile (marked E).
// They can move forward one tile at a time (increasing their score by 1 point), but never into a wall (#).
// They can also rotate clockwise or counterclockwise 90 degrees at a time (increasing their score by 1000 points).
//
// To figure out the best place to sit, you start by grabbing a map (your puzzle input) from a nearby kiosk. For example:
//
// ###############
// #.......#....E#
// #.#.###.#.###.#
// #.....#.#...#.#
// #.###.#####.#.#
// #.#.#.......#.#
// #.#.#####.###.#
// #...........#.#
// ###.#.#####.#.#
// #...#.....#.#.#
// #.#.#.###.#.#.#
// #.....#...#.#.#
// #.###.#.#.#.#.#
// #S..#.....#...#
// ###############
// There are many paths through this maze, but taking any of the best paths would incur a score of only 7036.
// This can be achieved by taking a total of 36 steps forward and turning 90 degrees a total of 7 times:
//
//
// ###############
// #.......#....E#
// #.#.###.#.###^#
// #.....#.#...#^#
// #.###.#####.#^#
// #.#.#.......#^#
// #.#.#####.###^#
// #..>>>>>>>>v#^#
// ###^#.#####v#^#
// #>>^#.....#v#^#
// #^#.#.###.#v#^#
// #^....#...#v#^#
// #^###.#.#.#v#^#
// #S..#.....#>>^#
// ###############
// Here's a second example:
//
// #################
// #...#...#...#..E#
// #.#.#.#.#.#.#.#.#
// #.#.#.#...#...#.#
// #.#.#.#.###.#.#.#
// #...#.#.#.....#.#
// #.#.#.#.#.#####.#
// #.#...#.#.#.....#
// #.#.#####.#.###.#
// #.#.#.......#...#
// #.#.###.#####.###
// #.#.#...#.....#.#
// #.#.#.#####.###.#
// #.#.#.........#.#
// #.#.#.#########.#
// #S#.............#
// #################
// In this maze, the best paths cost 11048 points; following one such path would look like this:
//
// #################
// #...#...#...#..E#
// #.#.#.#.#.#.#.#^#
// #.#.#.#...#...#^#
// #.#.#.#.###.#.#^#
// #>>v#.#.#.....#^#
// #^#v#.#.#.#####^#
// #^#v..#.#.#>>>>^#
// #^#v#####.#^###.#
// #^#v#..>>>>^#...#
// #^#v###^#####.###
// #^#v#>>^#.....#.#
// #^#v#^#####.###.#
// #^#v#^........#.#
// #^#v#^#########.#
// #S#>>^..........#
// #################
// Note that the path shown above includes one 90 degree turn as the very first move, rotating the Reindeer from facing East to facing North.
//
// Analyze your map carefully. What is the lowest score a Reindeer could possibly get?

const (
	Up    = '^'
	Down  = 'v'
	Right = '>'
	Left  = '<'
)

func Part1(input []byte) {
	start, _, tiles, maze := parse(string(input))

	pathScores := findPathScores(tiles, maze, start, '>', 0, []int{})
	fmt.Println("Possible path scores are", pathScores)
	fmt.Println("Lowest score is", min(pathScores))
}

func findPathScores(
	tiles [][]tile,
	maze [][]rune,
	current tile,
	currentDir rune,
	score int,
	finishes []int,
) []int {
	// time.Sleep(10 * time.Millisecond)
	// if reached finish - return current score
	if maze[current.row][current.column] == 'E' {
		return append(slices.Clone(finishes), score)
	}

	if maze[current.row][current.column] != 'S' {
		maze[current.row][current.column] = currentDir
		tiles[current.row][current.column].weight = score
	}

	prettyPrint(maze, &current)
	fmt.Println("Score", score)

	scores := []int{}

	upTile := nextTile(tiles, current, Up)
	upTileScore := newScore(score, currentDir, Up)
	if needsToVisit(maze, upTile, upTileScore) {
		upPaths := findPathScores(tiles, maze, upTile, Up, upTileScore, scores)
		scores = append(scores, upPaths...)
	}

	rightTile := nextTile(tiles, current, Right)
	rightTileScore := newScore(score, currentDir, Right)
	if needsToVisit(maze, rightTile, rightTileScore) {
		rightPaths := findPathScores(tiles, maze, rightTile, Right, rightTileScore, scores)
		scores = append(scores, rightPaths...)
	}

	downTile := nextTile(tiles, current, Down)
	downTileScore := newScore(score, currentDir, Down)
	if needsToVisit(maze, downTile, downTileScore) {
		downPaths := findPathScores(tiles, maze, downTile, Down, downTileScore, scores)
		scores = append(scores, downPaths...)
	}

	leftTile := nextTile(tiles, current, Left)
	leftTileScore := newScore(score, currentDir, Left)
	if needsToVisit(maze, leftTile, leftTileScore) {
		leftPaths := findPathScores(tiles, maze, leftTile, Left, leftTileScore, scores)
		scores = append(scores, leftPaths...)
	}

	return scores
}

type tile struct {
	row    int
	column int
	dir    rune
	weight int
}

func needsToVisit(maze [][]rune, t tile, proposedScore int) bool {
	if t.row < 0 || t.column < 0 || t.row >= len(maze) || t.column >= len(maze[t.row]) {
		return false
	}
	if maze[t.row][t.column] == '#' || maze[t.row][t.column] == 'S' {
		return false
	}
	return t.weight >= proposedScore
}

var (
	prev      [][]rune
	prevFocus *tile
)

func prettyPrint(maze [][]rune, focus *tile) {
	if prev == nil {
		fmt.Print("\033[H\033[2J")
		for range maze {
			fmt.Println(strings.Repeat(" ", len(maze[0])))
		}
		prev = make([][]rune, len(maze))
		for i := range maze {
			prev[i] = make([]rune, len(maze[i]))
		}
	}

	for r := range maze {
		for c := range maze[r] {
			curr, old := maze[r][c], prev[r][c]
			isF := focus != nil && r == focus.row && c == focus.column
			wasF := prevFocus != nil && r == prevFocus.row && c == prevFocus.column

			if curr != old || isF || wasF {
				// move & reset
				fmt.Printf("\033[%d;%dH\033[0m", r+1, c+1)

				switch {
				case isF:
					fmt.Printf("\033[32;42m%c\033[0m", curr) // fg+bg green
				case curr == 'S':
					fmt.Printf("\033[31mS\033[0m") // fg yellow
				case curr == 'E':
					fmt.Printf("\033[32mE\033[0m") // fg green
				case curr == '^' || curr == 'v' || curr == '<' || curr == '>':
					fmt.Printf("\033[33m%c\033[0m", curr) // fg blue
				default:
					fmt.Printf("%c", curr)
				}

				prev[r][c] = curr
			}
		}
	}

	// remember focus for next frame
	prevFocus = focus

	// cursor below maze
	fmt.Printf("\033[%d;1H", len(maze)+1)
}

func parse(input string) (start, finish tile, tiles [][]tile, maze [][]rune) {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	tiles = make([][]tile, len(lines))
	maze = make([][]rune, len(lines))
	for row, line := range lines {
		line = strings.Trim(line, "\n")
		maze[row] = []rune(line)
		tiles[row] = make([]tile, len(line))
		for column, t := range line {
			tiles[row][column] = tile{
				row:    row,
				column: column,
				weight: math.MaxInt,
			}
			switch t {
			case 'S':
				start = tile{
					row:    row,
					column: column,
					weight: 0,
				}
			case 'E':
				finish = tile{
					row:    row,
					column: column,
				}
			}
		}
	}
	return start, finish, tiles, maze
}

func min(nums []int) int {
	if len(nums) == 0 {
		panic("empty slice")
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func newScore(score int, current, next rune) int {
	dirScoreDiff := 0
	if current != next {
		dirScoreDiff = 1000
	}
	return score + 1 + dirScoreDiff
}

func nextTile(tiles [][]tile, current tile, dir rune) tile {
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
	return tiles[row][column]
}
