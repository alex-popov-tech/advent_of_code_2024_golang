package day_21

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// --- Day 21: Keypad Conundrum ---
// As you teleport onto Santa's Reindeer-class starship, The Historians begin to panic:
// someone from their search party is missing. A quick life-form scan by the ship's computer
// reveals that when the missing Historian teleported, he arrived in another part of the ship.
//
// The door to that area is locked, but the computer can't open it; it can only be opened by
// physically typing the door codes (your puzzle input) on the numeric keypad on the door.
//
// The numeric keypad has four rows of buttons: 789, 456, 123, and finally an empty
// gap followed by 0A. Visually, they are arranged like this:
//
// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//     | 0 | A |
//     +---+---+
// Unfortunately, the area outside the door is currently depressurized and nobody can go near the door. A robot needs to be sent instead.
//
// The robot has no problem navigating the ship and finding the numeric keypad, but it's
// not designed for button pushing: it can't be told to push a specific button directly.
// Instead, it has a robotic arm that can be controlled remotely via a directional keypad.
//
// The directional keypad has two rows of buttons: a gap / ^ (up) / A (activate) on the
// first row and < (left) / v (down) / > (right) on the second row. Visually, they
// are arranged like this:
//
//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+
// When the robot arrives at the numeric keypad, its robotic arm is pointed at the A button
// in the bottom right corner. After that, this directional keypad remote control must
// be used to maneuver the robotic arm: the up / down / left / right buttons cause
// it to move its arm one button in that direction, and the A button causes the robot
// to briefly move forward, pressing the button being aimed at by the robotic arm.
//
// For example, to make the robot type 029A on the numeric keypad, one sequence
// of inputs on the directional keypad you could use is:
//
// < to move the arm from A (its initial position) to 0.
// A to push the 0 button.
// ^A to move the arm to the 2 button and push it.
// >^^A to move the arm to the 9 button and push it.
// vvvA to move the arm to the A button and push it.
// In total, there are three shortest possible sequences of button presses on this
// directional keypad that would cause the robot to type 029A: <A^A>^^AvvvA, <A^A^>^AvvvA, and <A^A^^>AvvvA.
//
// Unfortunately, the area containing this directional keypad remote control is
// currently experiencing high levels of radiation and nobody can go near it.
// A robot needs to be sent instead.
//
// When the robot arrives at the directional keypad, its robot arm is pointed at the
// A button in the upper right corner. After that, a second, different directional
// keypad remote control is used to control this robot (in the same way as the first
// 	robot, except that this one is typing on a directional keypad instead of a numeric keypad).
//
// There are multiple shortest possible sequences of directional keypad button presses
// that would cause this robot to tell the first robot to type 029A on the door.
// One such sequence is v<<A>>^A<A>AvA<^AA>A<vAAA>^A.
//
// Unfortunately, the area containing this second directional keypad remote control
// is currently -40 degrees! Another robot will need to be sent to type on that directional keypad, too.
//
// There are many shortest possible sequences of directional keypad button presses
// that would cause this robot to tell the second robot to tell the first robot
// to eventually type 029A on the door. One such sequence is <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A.
//
// Unfortunately, the area containing this third directional keypad remote control
// is currently full of Historians, so no robots can find a clear path there.
// Instead, you will have to type this sequence yourself.
//
// Were you to choose this sequence of button presses, here are all of the
// buttons that would be pressed on your directional keypad, the two robots'
// directional keypads, and the numeric keypad:
//
// <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A
// v<<A>>^A<A>AvA<^AA>A<vAAA>^A
// <A^A>^^AvvvA
// 029A
// In summary, there are the following keypads:
//
// One directional keypad that you are using.
// Two directional keypads that robots are using.
// One numeric keypad (on a door) that a robot is using.
// It is important to remember that these robots are not designed for button pushing.
// In particular, if a robot arm is ever aimed at a gap where no button is present
// on the keypad, even for an instant, the robot will panic unrecoverably.
// So, don't do that. All robots will initially aim at the keypad's A key, wherever it is.
//
// To unlock the door, five codes will need to be typed on its numeric keypad. For example:
//
// 029A
// 980A
// 179A
// 456A
// 379A
// For each of these, here is a shortest sequence of button presses you could type to cause the desired code to be typed on the numeric keypad:
//
// 029A: <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A
// 980A: <v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A
// 179A: <v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
// 456A: <v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A
// 379A: <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
// The Historians are getting nervous; the ship computer doesn't remember whether the
// missing Historian is trapped in the area containing a giant electromagnet or molten lava.
// You'll need to make sure that for each of the five codes, you find the shortest
// sequence of button presses necessary.
//
// The complexity of a single code (like 029A) is equal to the result of multiplying these two values:
//
// The length of the shortest sequence of button presses you need to type on your
// directional keypad in order to cause the code to be typed on the numeric keypad;
// for 029A, this would be 68.
// The numeric part of the code (ignoring leading zeroes); for 029A, this would be 29.
// In the above example, complexity of the five codes can be found by calculating
// 68 * 29, 60 * 980, 68 * 179, 64 * 456, and 64 * 379. Adding these together produces 126384.
//
// Find the fewest number of button presses you'll need to perform in order to cause the
// robot in front of the door to type each code. What is the sum of the complexities of
// the five codes on your list?

const (
	Accept = 'A'
	Up     = '^'
	Right  = '>'
	Down   = 'v'
	Left   = '<'
	Zero   = '0'
	One    = '1'
	Two    = '2'
	Three  = '3'
	Four   = '4'
	Five   = '5'
	Six    = '6'
	Seven  = '7'
	Eight  = '8'
	Nine   = '9'
)

func Part1(input []byte) {
	arrowpadsCount := 2
	res := 0
	for _, code := range parse(string(input)) {
		fmt.Println("Code is", code)
		numPart := numericPart(code)
		seq := findNumPaths([]rune(code))
		fmt.Println("Paths for numpad")
		printPaths(seq)

		for i := range arrowpadsCount {
			arrowsSeq := [][]rune{}
			for _, path := range seq {
				p := findArrPaths(path)
				arrowsSeq = append(arrowsSeq, p...)
			}
			seq = arrowsSeq

			fmt.Println(len(seq), "paths after arrowpad", i+1, "/2")
		}

		shortestPath := len(shorten(seq)[0])
		fmt.Println("Shortest path is", shortestPath, "num part is", numPart)
		res += shortestPath * numPart
	}
	fmt.Println("Result is", res)
}

type Tile struct {
	row, column int
	r           rune
}

func t(row, column int, r rune) *Tile {
	return &Tile{row: row, column: column, r: r}
}

func printPaths(paths [][]rune) {
	for _, p := range paths {
		fmt.Println(string(p), len(p))
	}
}

func findNumPaths(seq []rune) [][]rune {
	numpad := [][]*Tile{
		{t(0, 0, '7'), t(0, 1, '8'), t(0, 2, '9')},
		{t(1, 0, '4'), t(1, 1, '5'), t(1, 2, '6')},
		{t(2, 0, '1'), t(2, 1, '2'), t(2, 2, '3')},
		{nil, t(3, 1, '0'), t(3, 2, 'A')},
	}
	return findPathsFor(seq, numpad, numpad[3][2])
}

func findArrPaths(seq []rune) [][]rune {
	arrowpad := [][]*Tile{
		{nil, t(0, 1, Up), t(0, 2, Accept)},
		{t(1, 0, Left), t(1, 1, Down), t(1, 2, Right)},
	}
	turns := findPathsFor(seq, arrowpad, arrowpad[0][2])
	return turns
}

func findPathsFor(seq []rune, matrix [][]*Tile, start *Tile) [][]rune {
	result := [][]rune{}
	for _, search := range seq {
		paths, turns := dfs(matrix, start, search, []*Tile{}, []rune{})
		turns = shorten(turns)
		turns = appendToAll(turns, 'A')
		result = multiply(result, turns)
		start = paths[0][len(paths[0])-1]
	}
	return result
}

func dfs(
	matrix [][]*Tile,
	current *Tile,
	expected rune,
	path []*Tile,
	turns []rune,
) (finalPaths [][]*Tile, finalTurns [][]rune) {
	if current.r == expected {
		return [][]*Tile{append(slices.Clone(path), current)}, [][]rune{turns}
	}

	// UP
	if available(matrix, path, current.row-1, current.column) {
		foundPaths, foundTurns := dfs(
			matrix,
			matrix[current.row-1][current.column],
			expected,
			append(slices.Clone(path), current),
			append(slices.Clone(turns), Up),
		)
		finalPaths = append(finalPaths, foundPaths...)
		finalTurns = append(finalTurns, foundTurns...)
	}
	// RIGHT
	if available(matrix, path, current.row, current.column+1) {
		foundPaths, foundTurns := dfs(
			matrix,
			matrix[current.row][current.column+1],
			expected,
			append(slices.Clone(path), current),
			append(slices.Clone(turns), Right),
		)
		finalPaths = append(finalPaths, foundPaths...)
		finalTurns = append(finalTurns, foundTurns...)
	}
	// DOWN
	if available(matrix, path, current.row+1, current.column) {
		foundPaths, foundTurns := dfs(
			matrix,
			matrix[current.row+1][current.column],
			expected,
			append(slices.Clone(path), current),
			append(slices.Clone(turns), Down),
		)
		finalPaths = append(finalPaths, foundPaths...)
		finalTurns = append(finalTurns, foundTurns...)
	}
	// LEFT
	if available(matrix, path, current.row, current.column-1) {
		foundPaths, foundTurns := dfs(
			matrix,
			matrix[current.row][current.column-1],
			expected,
			append(slices.Clone(path), current),
			append(slices.Clone(turns), Left),
		)
		finalPaths = append(finalPaths, foundPaths...)
		finalTurns = append(finalTurns, foundTurns...)
	}

	return finalPaths, finalTurns
}

func available(matrix [][]*Tile, path []*Tile, row, column int) bool {
	if row < 0 || row >= len(matrix) {
		return false
	}
	if column < 0 || column >= len(matrix[row]) {
		return false
	}
	if matrix[row][column] == nil {
		return false
	}
	return !slices.ContainsFunc(path, func(t *Tile) bool {
		return t.r == matrix[row][column].r
	})
}

func parse(input string) []string {
	return strings.Split(strings.Trim(input, "\n"), "\n")
}

func numericPart(code string) int {
	res, err := strconv.Atoi(code[0:3])
	if err != nil {
		panic(err)
	}
	return res
}

func shorten[T any](paths [][]T) [][]T {
	if len(paths) == 0 {
		panic("slice is empty")
	}
	minLength := len(paths[0])
	for _, p := range paths {
		if len(p) < minLength {
			minLength = len(p)
		}
	}
	res := [][]T{}
	for _, p := range paths {
		if len(p) == minLength {
			res = append(res, p)
		}
	}
	return res
}

func multiply(f, s [][]rune) [][]rune {
	if len(f) == 0 && len(s) == 0 {
		panic("slice is empty")
	}
	if len(f) == 0 {
		return s
	}
	if len(s) == 0 {
		return f
	}

	res := [][]rune{}
	for _, fp := range f {
		for _, sp := range s {
			res = append(res, append(slices.Clone(fp), sp...))
		}
	}
	return res
}

func appendToAll(slices [][]rune, item rune) [][]rune {
	for i := range slices {
		slices[i] = append(slices[i], item)
	}
	return slices
}
