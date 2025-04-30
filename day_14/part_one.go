package day_14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// --- Day 14: Restroom Redoubt ---
// One of The Historians needs to use the bathroom; fortunately, you know there's a bathroom near an unvisited location on their list, and so you're all quickly teleported directly to the lobby of Easter Bunny Headquarters.
//
// Unfortunately, EBHQ seems to have "improved" bathroom security again after your last visit. The area outside the bathroom is swarming with robots!
//
// To get The Historian safely to the bathroom, you'll need a way to predict where the robots will be in the future. Fortunately, they all seem to be moving on the tile floor in predictable straight lines.
//
// You make a list (your puzzle input) of all of the robots' current positions (p) and velocities (v), one robot per line. For example:
//
// p=0,4 v=3,-3
// p=6,3 v=-1,-3
// p=10,3 v=-1,2
// p=2,0 v=2,-1
// p=0,0 v=1,3
// p=3,0 v=-2,-2
// p=7,6 v=-1,-3
// p=3,0 v=-1,-2
// p=9,3 v=2,3
// p=7,3 v=-1,2
// p=2,4 v=2,-3
// p=9,5 v=-3,-3
// Each robot's position is given as p=x,y where x represents the number of tiles the robot is from the left wall and y represents the number of tiles from the top wall (when viewed from above). So, a position of p=0,0 means the robot is all the way in the top-left corner.
//
// Each robot's velocity is given as v=x,y where x and y are given in tiles per second. Positive x means the robot is moving to the right, and positive y means the robot is moving down. So, a velocity of v=1,-2 means that each second, the robot moves 1 tile to the right and 2 tiles up.
//
// The robots outside the actual bathroom are in a space which is 101 tiles wide and 103 tiles tall (when viewed from above). However, in this example, the robots are in a space which is only 11 tiles wide and 7 tiles tall.
//
// The robots are good at navigating over/under each other (due to a combination of springs, extendable legs, and quadcopters), so they can share the same tile and don't interact with each other. Visually, the number of robots on each tile in this example looks like this:
//
// 1.12.......
// ...........
// ...........
// ......11.11
// 1.1........
// .........1.
// .......1...
// These robots have a unique feature for maximum bathroom security: they can teleport. When a robot would run into an edge of the space they're in, they instead teleport to the other side, effectively wrapping around the edges. Here is what robot p=2,4 v=2,-3 does for the first few seconds:
//
// Initial state:
// ...........
// ...........
// ...........
// ...........
// ..1........
// ...........
// ...........
//
// After 1 second:
// ...........
// ....1......
// ...........
// ...........
// ...........
// ...........
// ...........
//
// After 2 seconds:
// ...........
// ...........
// ...........
// ...........
// ...........
// ......1....
// ...........
//
// After 3 seconds:
// ...........
// ...........
// ........1..
// ...........
// ...........
// ...........
// ...........
//
// After 4 seconds:
// ...........
// ...........
// ...........
// ...........
// ...........
// ...........
// ..........1
//
// After 5 seconds:
// ...........
// ...........
// ...........
// .1.........
// ...........
// ...........
// ...........
// The Historian can't wait much longer, so you don't have to simulate the robots for very long. Where will the robots be after 100 seconds?
//
// In the above example, the number of robots on each tile after 100 seconds has elapsed looks like this:
//
// ......2..1.
// ...........
// 1..........
// .11........
// .....1.....
// ...12......
// .1....1....
// To determine the safest area, count the number of robots in each quadrant after 100 seconds. Robots that are exactly in the middle (horizontally or vertically) don't count as being in any quadrant, so the only relevant robots are:
//
// ..... 2..1.
// ..... .....
// 1.... .....
//
// ..... .....
// ...12 .....
// .1... 1....
// In this example, the quadrants contain 1, 3, 4, and 1 robot. Multiplying these together gives a total safety factor of 12.
//
// Predict the motion of the robots in your list within a space which is 101 tiles wide and 103 tiles tall. What will the safety factor be after exactly 100 seconds have elapsed?

func Part1(input []byte) {
	line := strings.Trim(string(input), "\n")
	const rows = 103
	const columns = 101
	const steps = 100

	robots := parse(line)

	for _, it := range robots {
		fmt.Printf(
			"Robot %s is at column:%d,row:%d, vel column:%d,row:%d\n",
			it.name,
			it.pos.column,
			it.pos.row,
			it.vel.column,
			it.vel.row,
		)
	}

	for i := range robots {
		for range steps {
			robots[i].move(rows, columns)
			prettyPrint(rows, columns, robots)
		}
	}

	prettyPrintQuadrants(rows, columns, robots)

	res := countRobotsInQuadrants(rows, columns, robots)
	fmt.Println("Result is", res)
}

type pos struct {
	column int
	row    int
}

type robot struct {
	pos    pos
	vel    pos
	name   string
	second int
}

func (r *robot) move(maxRow, maxColumn int) {
	newRow := r.pos.row + r.vel.row
	newColumn := r.pos.column + r.vel.column
	// was 1:4 ( 1-3:4+2  == -2:6)
	// became 5:6
	// v=2,-3 ( first columns, second rows )
	// 7 rows, 11 columns
	if newRow < 0 {
		newRow = maxRow + newRow
	}
	if newRow >= maxRow {
		newRow = newRow - maxRow
	}
	if newColumn < 0 {
		newColumn = maxColumn + newColumn
	}
	if newColumn >= maxColumn {
		newColumn = newColumn - maxColumn
	}
	r.pos = pos{row: newRow, column: newColumn}
}

func countRobotsInQuadrants(rows, columns int, robots []robot) int {
	topLeft := 0
	topRight := 0
	bottomLeft := 0
	bottomRight := 0
	for _, it := range robots {
		if it.pos.column == columns/2 || it.pos.row == rows/2 {
			continue
		} else if it.pos.column < columns/2 && it.pos.row < rows/2 {
			topLeft++
		} else if it.pos.column < columns/2 && it.pos.row > rows/2 {
			bottomLeft++
		} else if it.pos.column > columns/2 && it.pos.row < rows/2 {
			topRight++
		} else if it.pos.column > columns/2 && it.pos.row > rows/2 {
			bottomRight++
		}
	}

	fmt.Println("Counted robots in quadrants", topLeft, topRight, bottomLeft, bottomRight)
	return topLeft * bottomRight * topRight * bottomLeft
}

func parse(input string) []robot {
	lines := strings.Split(input, "\n")
	// p=2,4 v=2,-3
	robotRegexp := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	var robots []robot

	for i, line := range lines {
		matches := robotRegexp.FindStringSubmatch(line)
		robots = append(robots, robot{
			second: 0,
			name:   fmt.Sprintf("%d", i),
			pos: pos{
				column: atoi(matches[1]),
				row:    atoi(matches[2]),
			},
			vel: pos{
				column: atoi(matches[3]),
				row:    atoi(matches[4]),
			},
		})
	}

	return robots
}

func atoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func prettyPrint(rows, columns int, robots []robot) {
	robotsPerTile := make([][]int, columns)
	for i := range columns {
		robotsPerTile[i] = make([]int, rows)
	}
	for _, robot := range robots {
		robotsPerTile[robot.pos.column][robot.pos.row]++
	}

	for row := range rows {
		for column := range columns {
			if robotsPerTile[column][row] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", robotsPerTile[column][row])
			}
		}
		fmt.Println()
	}
}

func prettyPrintQuadrants(rows, columns int, robots []robot) {
	robotsPerTile := make([][]int, columns)
	for i := range columns {
		robotsPerTile[i] = make([]int, rows)
	}
	for _, robot := range robots {
		robotsPerTile[robot.pos.column][robot.pos.row]++
	}

	for row := range rows {
		for column := range columns {
			if column == columns/2 || row == rows/2 {
				fmt.Printf(" ")
			} else if robotsPerTile[column][row] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", robotsPerTile[column][row])
			}
		}
		fmt.Println()
	}
}
