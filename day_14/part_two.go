package day_14

import (
	"fmt"
	"os"
	"strings"
)

// --- Part Two ---
// During the bathroom break, someone notices that these robots seem awfully similar to ones built and used at the North Pole.
// If they're the same type of robots, they should have a hard-coded Easter egg: very rarely, most of the robots should arrange themselves into a picture of a Christmas tree.
//
// What is the fewest number of seconds that must elapse for the robots to display the Easter egg?

func Part2(input []byte) {
	line := strings.Trim(string(input), "\n")
	const rows = 103
	const columns = 101
	const steps = 1000 * 10

	robots := parse(line)

	suspicious := make(chan string, 10)
	doneWriting := make(chan bool)

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	go func() {
		for it := range suspicious {
			_, err = file.WriteString(it)
			if err != nil {
				panic(err)
			}
			_, err = file.WriteString("\n")
			if err != nil {
				panic(err)
			}
		}
		file.Close()
		doneWriting <- true
	}()

	for step := range steps {
		for i := range robots {
			robots[i].move(rows, columns)
		}
		// if found suspiciously many digits in line, log it
		// so then we can just search through this file, adding
		// one digit at once to find that tree
		if hasLine(robots, 5) {
			fmt.Println("Found it at step", step)
			suspicious <- pretty(step, rows, columns, robots)
		}
	}
	close(suspicious)
	<-doneWriting
}

func pretty(step, rows, columns int, robots []robot) string {
	robotsPerTile := make([][]int, columns)
	for i := range columns {
		robotsPerTile[i] = make([]int, rows)
	}
	for _, robot := range robots {
		robotsPerTile[robot.pos.column][robot.pos.row]++
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Step %d\n", step))
	for row := range rows {
		for column := range columns {
			if robotsPerTile[column][row] == 0 {
				sb.WriteString(".")
			} else {
				sb.WriteString(fmt.Sprintf("%d", robotsPerTile[column][row]))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func hasLine(robots []robot, count int) bool {
	matrix := make([][]int, 150)
	for i := range matrix {
		matrix[i] = make([]int, 150)
	}

	for _, robot := range robots {
		matrix[robot.pos.row][robot.pos.column]++
	}
	// all below is AI generated because i don't find it fun
	rows := len(matrix)
	if rows == 0 || count <= 0 {
		return false
	}
	cols := len(matrix[0])

	// Check rows (horizontal)
	for _, row := range matrix {
		streak := 0
		for _, val := range row {
			if val > 0 {
				streak++
				if streak == count {
					return true
				}
			} else {
				streak = 0
			}
		}
	}

	// Check columns (vertical)
	for col := 0; col < cols; col++ {
		streak := 0
		for row := 0; row < rows; row++ {
			if matrix[row][col] > 0 {
				streak++
				if streak == count {
					return true
				}
			} else {
				streak = 0
			}
		}
	}

	return false
}
