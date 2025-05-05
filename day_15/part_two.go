package day_15

import (
	"fmt"
	"strings"
	"time"
)

// --- Part Two ---
// The lanternfish use your information to find a safe moment to swim in and turn off the malfunctioning robot!
// Just as they start preparing a festival in your honor, reports start coming in that a second warehouse's robot is also malfunctioning.
//
// This warehouse's layout is surprisingly similar to the one you just helped. There is one key difference: everything except the robot is twice as wide!
// The robot's list of movements doesn't change.
//
// To get the wider warehouse's map, start with your original map and, for each tile, make the following changes:
//
// If the tile is #, the new map contains ## instead.
// If the tile is O, the new map contains [] instead.
// If the tile is ., the new map contains .. instead.
// If the tile is @, the new map contains @. instead.
// This will produce a new warehouse map which is twice as wide and with wide boxes that are represented by []. (The robot does not change size.)
//
// The larger example from before would now look like this:
//
// ####################
// ##....[]....[]..[]##
// ##............[]..##
// ##..[][]....[]..[]##
// ##....[]@.....[]..##
// ##[]##....[]......##
// ##[]....[]....[]..##
// ##..[][]..[]..[][]##
// ##........[]......##
// ####################
// Because boxes are now twice as wide but the robot is still the same size and speed, boxes can be aligned such that they directly push two other boxes at once.
// For example, consider this situation:
//
// #######
// #...#.#
// #.....#
// #..OO@#
// #..O..#
// #.....#
// #######
//
// <vv<<^^<<^^
// After appropriately resizing this map, the robot would push around these boxes as follows:
//
// Initial state:
// ##############
// ##......##..##
// ##..........##
// ##....[][]@.##
// ##....[]....##
// ##..........##
// ##############
//
// Move <:
// ##############
// ##......##..##
// ##..........##
// ##...[][]@..##
// ##....[]....##
// ##..........##
// ##############
//
// Move v:
// ##############
// ##......##..##
// ##..........##
// ##...[][]...##
// ##....[].@..##
// ##..........##
// ##############
//
// Move v:
// ##############
// ##......##..##
// ##..........##
// ##...[][]...##
// ##....[]....##
// ##.......@..##
// ##############
//
// Move <:
// ##############
// ##......##..##
// ##..........##
// ##...[][]...##
// ##....[]....##
// ##......@...##
// ##############
//
// Move <:
// ##############
// ##......##..##
// ##..........##
// ##...[][]...##
// ##....[]....##
// ##.....@....##
// ##############
//
// Move ^:
// ##############
// ##......##..##
// ##...[][]...##
// ##....[]....##
// ##.....@....##
// ##..........##
// ##############
//
// Move ^:
// ##############
// ##......##..##
// ##...[][]...##
// ##....[]....##
// ##.....@....##
// ##..........##
// ##############
//
// Move <:
// ##############
// ##......##..##
// ##...[][]...##
// ##....[]....##
// ##....@.....##
// ##..........##
// ##############
//
// Move <:
// ##############
// ##......##..##
// ##...[][]...##
// ##....[]....##
// ##...@......##
// ##..........##
// ##############
//
// Move ^:
// ##############
// ##......##..##
// ##...[][]...##
// ##...@[]....##
// ##..........##
// ##..........##
// ##############
//
// Move ^:
// ##############
// ##...[].##..##
// ##...@.[]...##
// ##....[]....##
// ##..........##
// ##..........##
// ##############
// This warehouse also uses GPS to locate the boxes.
// For these larger boxes, distances are measured from the edge of the map to the closest edge of the box in question.
// So, the box shown below has a distance of 1 from the top edge of the map and 5 from the left edge of the map, resulting in a GPS coordinate of 100 * 1 + 5 = 105.
//
// ##########
// ##...[]...
// ##........
// In the scaled-up version of the larger example from above, after the robot has finished all of its moves, the warehouse would look like this:
//
// ####################
// ##[].......[].[][]##
// ##[]...........[].##
// ##[]........[][][]##
// ##[]......[]....[]##
// ##..##......[]....##
// ##..[]............##
// ##..@......[].[][]##
// ##......[][]..[]..##
// ####################
// The sum of these boxes' GPS coordinates is 9021.
//
// Predict the motion of the robot and boxes in this new, scaled-up warehouse. What is the sum of all boxes' final GPS coordinates?

func Part2(input []byte) {
	mop, commands, robot := parse2(string(input))

	for i, dir := range commands {
		time.Sleep(50 * time.Millisecond)
		prettyPrint2(mop, nil, true)
		fmt.Printf("Step %d/%d command '%c'\n", i+1, len(commands), dir)
		if canMove2(mop, robot, dir) {
			move2(mop, robot, dir)
			robot = next(robot, dir)
		}
	}
	prettyPrint2(mop, nil, true)

	res := 0
	for row := range mop {
		for column := range mop[row] {
			if mop[row][column] == '[' {
				res += row*100 + column
			}
		}
	}
	fmt.Println("Result is", res)
}

func canMove2(mop [][]rune, current tile, dir rune) bool {
	// prettyPrint2(mop, &current, true)
	switch mop[current.row][current.column] {
	case '@':
		return canMove2(mop, next(current, dir), dir)
	case '.':
		return true
	case '#':
		return false
	case '[', ']':
		nxt := next(current, dir)
		canCurrentMoveToNext := canMove2(mop, nxt, dir)
		if !canCurrentMoveToNext {
			return false
		}

		// as pairs are only placed on X axis if we are moving
		// horizontally, we don't need to check pairs because
		// they already checked as 'next' ones
		if dir == 'v' || dir == '^' {
			p := pair(mop, current)
			pnext := next(p, dir)
			canMovePairToNext := canMove2(mop, pnext, dir)
			if !canMovePairToNext {
				return false
			}
		}
		return true
	default:
		panic("unknown tile type" + string(mop[current.row][current.column]))
	}
}

// should be called only after 'canMove2'
func move2(mop [][]rune, current tile, dir rune) {
	// prettyPrint2(mop, &current, true)
	nxt := next(current, dir)
	switch mop[current.row][current.column] {
	case '.':
		return
	case '@':
		// move next tile if needed
		if mop[nxt.row][nxt.column] != '.' {
			move2(mop, nxt, dir)
		}
		// then swap with next tile ( should be empty by now )
		swap(mop, current, nxt)
	case '[', ']':
		p := pair(mop, current)
		// move next tile if needed
		if mop[nxt.row][nxt.column] != '.' {
			move2(mop, nxt, dir)
		}
		swap(mop, current, nxt)

		// as pairs are only placed on X axis if we are moving
		// horizontally, we don't need to check pairs because
		// they already checked as 'next' ones
		if dir == 'v' || dir == '^' {
			// move pair's next tile if needed
			pnxt := next(p, dir)
			if mop[pnxt.row][pnxt.column] != '.' {
				move2(mop, pnxt, dir)
			}
			swap(mop, p, pnxt)
		}
	default:
		panic("unexpected tile type" + string(mop[current.row][current.column]))
	}
}

func pair(mop [][]rune, t tile) tile {
	switch mop[t.row][t.column] {
	case '[':
		return tile{
			row:    t.row,
			column: t.column + 1,
		}
	case ']':
		return tile{
			row:    t.row,
			column: t.column - 1,
		}
	default:
		panic("unknown tile type" + string(mop[t.row][t.column]))
	}
}

func parse2(input string) (mop [][]rune, commands []rune, rob tile) {
	chunks := strings.Split(input, "\n\n")

	commands = []rune(strings.ReplaceAll(strings.Trim(chunks[1], "\n"), "\n", ""))

	mopStr := strings.Split(chunks[0], "\n")
	mop = make([][]rune, len(mopStr))
	for row, line := range mopStr {
		for _, r := range line {
			mop[row] = append(mop[row], widen(r)...)
			if r == '@' {
				rob = tile{
					row:    row,
					column: len(mop[row]) - 2,
				}
			}
		}
	}
	return mop, commands, rob
}

func widen(r rune) []rune {
	switch r {
	case '#':
		return []rune{'#', '#'}
	case 'O':
		return []rune{'[', ']'}
	case '.':
		return []rune{'.', '.'}
	case '@':
		return []rune{'@', '.'}
	default:
		panic("unknown tile type")
	}
}

func prettyPrint2(mop [][]rune, focus *tile, clear bool) {
	var highlight tile
	if focus != nil {
		highlight = *focus
	} else {
		highlight = tile{
			row:    -1,
			column: -1,
		}
	}
	if clear {
		fmt.Print("\033[H\033[2J")
	}
	for row, line := range mop {
		sb := strings.Builder{}
		for column, c := range line {
			if c == '@' {
				sb.WriteString("\033[31m" + string(c) + "\033[0m")
			} else if row == highlight.row && column == highlight.column {
				sb.WriteString("\033[32m" + string(c) + "\033[0m")
			} else {
				sb.WriteString(string(c))
			}
		}
		fmt.Println(sb.String())
	}
}
