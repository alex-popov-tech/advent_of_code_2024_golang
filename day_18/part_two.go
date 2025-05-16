package day_18

import (
	"fmt"
	"slices"
	"strings"
)

// --- Part Two ---
// The Historians aren't as used to moving around in this pixelated universe as you are.
// You're afraid they're not going to be fast enough to make it to the exit before the
// path is completely blocked.
//
// To determine how fast everyone needs to go, you need to determine the first byte that
// will cut off the path to the exit.
//
// In the above example, after the byte at 1,1 falls, there is still a path to the exit:
//
// O..#OOO
// O##OO#O
// O#OO#OO
// OOO#OO#
// ###OO##
// .##O###
// #.#OOOO
// However, after adding the very next byte (at 6,1), there is no longer a path to the exit:
//
// ...#...
// .##..##
// .#..#..
// ...#..#
// ###..##
// .##.###
// #.#....
// So, in this example, the coordinates of the first byte that prevents the exit from
// being reachable are 6,1.
//
// Simulate more of the bytes that are about to corrupt your memory space.
// What are the coordinates of the first byte that will prevent the exit from being
// reachable from your starting position? (Provide the answer as two integers separated
// by a comma with no other characters.)

const (
	Up    = '^'
	Down  = 'v'
	Right = '>'
	Left  = '<'
)

func Part2(input []byte) {
	matrix, walls := parse2(string(input), size)

	// 1. Place all walls at the beginning
	for i := 0; i < len(walls); i++ {
		matrix[walls[i].row][walls[i].column].r = '#'
	}

	for i := len(walls) - 1; i >= 0; i-- {
		// 2. Walk to then end
		hasPath := hasPath(matrix, matrix[startRow][startColumn], matrix[endRow][endColumn], Path{})
		if hasPath {
			// 4. Repeat steps 2 and 3 until you can reach the end, return
			fmt.Println("There is a path from start to end now, exiting")
			return
		} else {
			fmt.Printf("There is no path from start to end, removing last wall at %d:%d\n", walls[i].row, walls[i].column)
		}
		// 3. If end is unreachable, remove last wall
		matrix[walls[i].row][walls[i].column].r = '.'
	}
}

type Path struct {
	tiles []*tile
	score int
}

func (p Path) Contains(tile *tile) bool {
	for _, t := range p.tiles {
		if t.row == tile.row && t.column == tile.column {
			return true
		}
	}
	return false
}

func (p Path) String() string {
	pairs := []string{}
	for _, t := range p.tiles {
		pairs = append(pairs, fmt.Sprintf("%d:%d", t.row, t.column))
	}
	return strings.Join(pairs, " => ")
}

func hasPath(matrix [][]*tile, from *tile, to *tile, path Path) bool {
	from.isVisited = true
	path.tiles = append(slices.Clone(path.tiles), from)

	if from.row == to.row && from.column == to.column {
		return true
	}

	upTile, hasTile := nextTile(matrix, from, Up)
	if hasTile && needsToVisit(matrix, path.tiles, upTile) {
		newPath := Path{tiles: slices.Clone(path.tiles)}
		upPath := hasPath(matrix, upTile, to, newPath)
		if upPath {
			return true
		}
	}

	rightTile, hasTile := nextTile(matrix, from, Right)
	if hasTile && needsToVisit(matrix, path.tiles, rightTile) {
		newPath := Path{tiles: slices.Clone(path.tiles)}
		rightPath := hasPath(matrix, rightTile, to, newPath)
		if rightPath {
			return true
		}
	}

	downTile, hasTile := nextTile(matrix, from, Down)
	if hasTile && needsToVisit(matrix, path.tiles, downTile) {
		newPath := Path{tiles: slices.Clone(path.tiles)}
		downPath := hasPath(matrix, downTile, to, newPath)
		if downPath {
			return true
		}
	}

	leftTile, hasTile := nextTile(matrix, from, Left)
	if hasTile && needsToVisit(matrix, path.tiles, leftTile) {
		newPath := Path{tiles: slices.Clone(path.tiles)}
		leftPath := hasPath(matrix, leftTile, to, newPath)
		if leftPath {
			return true
		}
	}

	return false
}

func nextTile(matrix [][]*tile, current *tile, dir rune) (tile *tile, hasTile bool) {
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
	if row < 0 || row >= len(matrix) || column < 0 || column >= len(matrix[row]) {
		return nil, false
	}
	return matrix[row][column], true
}

func needsToVisit(
	maze [][]*tile,
	visited []*tile,
	target *tile,
) bool {
	if target.row < 0 || target.column < 0 || target.row >= len(maze) ||
		target.column >= len(maze[target.row]) {
		return false
	}
	if maze[target.row][target.column].r == '#' {
		return false
	}
	if slices.ContainsFunc(visited, func(t *tile) bool {
		return t.row == target.row && t.column == target.column
	}) {
		return false
	}
	return true
}

func parse2(input string, size int) ([][]*tile, []*tile) {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")

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

	walls := []*tile{}
	for _, coord := range lines {
		column := atoi(strings.Split(coord, ",")[0])
		row := atoi(strings.Split(coord, ",")[1])
		walls = append(walls, &tile{
			r:         '#',
			row:       row,
			column:    column,
			score:     emptyScore,
			isVisited: false,
		})
	}
	return matrix, walls
}
