package day_15

import (
	"fmt"
	"strings"
	"time"
)

// --- Day 15: Warehouse Woes ---
// You appear back inside your own mini submarine! Each Historian drives their mini submarine in a different direction; maybe the Chief has his own submarine down here somewhere as well?
//
// You look up to see a vast school of lanternfish swimming past you. On closer inspection, they seem quite anxious, so you drive your mini submarine over to see if you can help.
//
// Because lanternfish populations grow rapidly, they need a lot of food, and that food needs to be stored somewhere. That's why these lanternfish have built elaborate warehouse complexes operated by robots!
//
// These lanternfish seem so anxious because they have lost control of the robot that operates one of their most important warehouses! It is currently running amok, pushing around boxes in the warehouse with no regard for lanternfish logistics or lanternfish inventory management strategies.
//
// Right now, none of the lanternfish are brave enough to swim up to an unpredictable robot so they could shut it off. However, if you could anticipate the robot's movements, maybe they could find a safe option.
//
// The lanternfish already have a map of the warehouse and a list of movements the robot will attempt to make (your puzzle input). The problem is that the movements will sometimes fail as boxes are shifted around, making the actual movements of the robot difficult to predict.
//
// For example:
//
// ##########
// #..O..O.O#
// #......O.#
// #.OO..O.O#
// #..O@..O.#
// #O#..O...#
// #O..O..O.#
// #.OO.O.OO#
// #....O...#
// ##########
//
// <vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
// vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
// ><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
// <<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
// ^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
// ^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
// >^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
// <><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
// ^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
// v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
// As the robot (@) attempts to move, if there are any boxes (O) in the way, the robot will also attempt to push those boxes. However, if this action would cause the robot or a box to move into a wall (#), nothing moves instead, including the robot. The initial positions of these are shown on the map at the top of the document the lanternfish gave you.
//
// The rest of the document describes the moves (^ for up, v for down, < for left, > for right) that the robot will attempt to make, in order. (The moves form a single giant sequence; they are broken into multiple lines just to make copy-pasting easier. Newlines within the move sequence should be ignored.)
//
// Here is a smaller example to get started:
//
// ########
// #..O.O.#
// ##@.O..#
// #...O..#
// #.#.O..#
// #...O..#
// #......#
// ########
//
// <^^>>>vv<v>>v<<
// Were the robot to attempt the given sequence of moves, it would push around the boxes as follows:
//
// Initial state:
// ########
// #..O.O.#
// ##@.O..#
// #...O..#
// #.#.O..#
// #...O..#
// #......#
// ########
//
// Move <:
// ########
// #..O.O.#
// ##@.O..#
// #...O..#
// #.#.O..#
// #...O..#
// #......#
// ########
//
// Move ^:
// ########
// #.@O.O.#
// ##..O..#
// #...O..#
// #.#.O..#
// #...O..#
// #......#
// ########
//
// Move ^:
// ########
// #.@O.O.#
// ##..O..#
// #...O..#
// #.#.O..#
// #...O..#
// #......#
// ########
//
// Move >:
// ########
// #..@OO.#
// ##..O..#
// #...O..#
// #.#.O..#
// #...O..#
// #......#
// ########
//
// Move >:
// ########
// #...@OO#
// ##..O..#
// #...O..#
// #.#.O..#
// #...O..#
// #......#
// ########
//
// Move >:
// ########
// #...@OO#
// ##..O..#
// #...O..#
// #.#.O..#
// #...O..#
// #......#
// ########
//
// Move v:
// ########
// #....OO#
// ##..@..#
// #...O..#
// #.#.O..#
// #...O..#
// #...O..#
// ########
//
// Move v:
// ########
// #....OO#
// ##..@..#
// #...O..#
// #.#.O..#
// #...O..#
// #...O..#
// ########
//
// Move <:
// ########
// #....OO#
// ##.@...#
// #...O..#
// #.#.O..#
// #...O..#
// #...O..#
// ########
//
// Move v:
// ########
// #....OO#
// ##.....#
// #..@O..#
// #.#.O..#
// #...O..#
// #...O..#
// ########
//
// Move >:
// ########
// #....OO#
// ##.....#
// #...@O.#
// #.#.O..#
// #...O..#
// #...O..#
// ########
//
// Move >:
// ########
// #....OO#
// ##.....#
// #....@O#
// #.#.O..#
// #...O..#
// #...O..#
// ########
//
// Move v:
// ########
// #....OO#
// ##.....#
// #.....O#
// #.#.O@.#
// #...O..#
// #...O..#
// ########
//
// Move <:
// ########
// #....OO#
// ##.....#
// #.....O#
// #.#O@..#
// #...O..#
// #...O..#
// ########
//
// Move <:
// ########
// #....OO#
// ##.....#
// #.....O#
// #.#O@..#
// #...O..#
// #...O..#
// ########
// The larger example has many more moves; after the robot has finished those moves, the warehouse would look like this:
//
// ##########
// #.O.O.OOO#
// #........#
// #OO......#
// #OO@.....#
// #O#.....O#
// #O.....OO#
// #O.....OO#
// #OO....OO#
// ##########
// The lanternfish use their own custom Goods Positioning System (GPS for short) to track the locations of the boxes. The GPS coordinate of a box is equal to 100 times its distance from the top edge of the map plus its distance from the left edge of the map. (This process does not stop at wall tiles; measure all the way to the edges of the map.)
//
// So, the box shown below has a distance of 1 from the top edge of the map and 4 from the left edge of the map, resulting in a GPS coordinate of 100 * 1 + 4 = 104.
//
// #######
// #...O..
// #......
// The lanternfish would like to know the sum of all boxes' GPS coordinates after the robot finishes moving. In the larger example, the sum of all boxes' GPS coordinates is 10092. In the smaller example, the sum is 2028.
//
// Predict the motion of the robot and boxes in the warehouse. After the robot is finished moving, what is the sum of all boxes' GPS coordinates?

func Part1(input []byte) {
	mop, commands, robot := parse(string(input))

	for i, command := range commands {
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Step %d/%d command '%c'\n", i+1, len(commands), command)
		prettyPrint(mop, robot)
		// fmt.Printf("Before step #%d command '%c'\n", i, command)
		if canMove(command, robot, mop) {
			robot = *move(command, &robot, mop)
		}
		// fmt.Printf("After step #%d command '%c'\n", i, command)
		// prettyPrint(mop)
	}
	fmt.Print("\033[H\033[2J")
	prettyPrint(mop, robot)

	res := 0
	for row := range mop {
		for column := range mop[row] {
			if mop[row][column] == 'O' {
				res += row*100 + column
			}
		}
	}
	fmt.Println("Result is", res)
}

type tile struct {
	row    int
	column int
}

func canMove(command rune, robot tile, mop [][]rune) bool {
	for n := next(robot, command); mop[n.row][n.column] != '#'; n = next(n, command) {
		if mop[n.row][n.column] == '.' {
			return true
		}
	}
	return false
}

// this is suppose to be called after 'canMove', so there is definately
// a place somewhere
func move(command rune, tile *tile, mop [][]rune) *tile {
	nextTile := next(*tile, command)
	switch mop[nextTile.row][nextTile.column] {
	case '#':
		return tile
	case '.':
		swap(mop, *tile, nextTile)
		return &nextTile
	case 'O':
		// its a box, move it first, and then move our tile
		move(command, &nextTile, mop)
		swap(mop, *tile, nextTile)
		return &nextTile
	default:
		panic("unknown tile type")
	}
}

func swap(mop [][]rune, f, s tile) {
	mop[f.row][f.column], mop[s.row][s.column] = mop[s.row][s.column], mop[f.row][f.column]
}

func next(t tile, command rune) tile {
	switch command {
	case 'v':
		return tile{
			row:    t.row + 1,
			column: t.column,
		}
	case '>':
		return tile{
			row:    t.row,
			column: t.column + 1,
		}
	case '<':
		return tile{
			row:    t.row,
			column: t.column - 1,
		}
	case '^':
		return tile{
			row:    t.row - 1,
			column: t.column,
		}
	default:
		panic("unknown command " + string(command))
	}
}

func prettyPrint(mop [][]rune, robot tile) {
	for row, line := range mop {
		if row != robot.row {
			fmt.Println(string(line))
		} else {
			sb := strings.Builder{}
			for column, c := range line {
				if column != robot.column {
					sb.WriteString(string(c))
				} else {
					sb.WriteString("\033[31m" + string(c) + "\033[0m")
				}
			}
			fmt.Println(sb.String())
		}
	}
}

func parse(input string) (mop [][]rune, commands []rune, robot tile) {
	chunks := strings.Split(input, "\n\n")

	commands = []rune(strings.ReplaceAll(strings.Trim(chunks[1], "\n"), "\n", ""))

	mopStr := strings.Split(chunks[0], "\n")
	for row, line := range mopStr {
		mop = append(mop, []rune(line))
		if robot.row == 0 && robot.column == 0 {
			for column, t := range line {
				if t == '@' {
					robot = tile{
						row:    row,
						column: column,
					}
				}
			}
		}
	}
	return mop, commands, robot
}
