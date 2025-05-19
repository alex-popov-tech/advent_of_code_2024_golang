package day_21

import (
	"fmt"
	"math"
)

// --- Part Two ---
// Just as the missing Historian is released, The Historians realize that a second member of their
// search party has also been missing this entire time!
//
// A quick life-form scan reveals the Historian is also trapped in a locked area of the ship. Due to a variety of hazards,
// robots are once again dispatched, forming another chain of remote control keypads managing robotic-arm-wielding robots.
//
// This time, many more robots are involved. In summary, there are the following keypads:
//
// One directional keypad that you are using.
// 25 directional keypads that robots are using.
// One numeric keypad (on a door) that a robot is using.
// The keypads form a chain, just like before: your directional keypad controls a robot which is typing on a directional
// keypad which controls a robot which is typing on a directional keypad... and so on, ending with the robot which is typing on the numeric keypad.
//
// The door codes are the same this time around; only the number of robots and directional keypads has changed.
//
// Find the fewest number of button presses you'll need to perform in order to cause the robot in front of the door to type each code.
// What is the sum of the complexities of the five codes on your list?

const (
	ARROWPAD_START_ROW    = 0
	ARROWPAD_START_COLUMN = 2
)

func Part2(input []byte) {
	// there are 25 robots and me, so 3 arrow keypads, but first one is calculated separately
	arrowpadsCount := 26 - 1
	res := 0
	for _, code := range parse(string(input)) {
		fmt.Println("Code is", code)
		numPart := numericPart(code)
		arrowSequences := findNumPaths([]rune(code))

		arrowpad := [][]*Tile{
			{nil, t(0, 1, Up), t(0, 2, Accept)},
			{t(1, 0, Left), t(1, 1, Down), t(1, 2, Right)},
		}

		sum := math.MaxInt
		cache := Cache{}
		for _, arrowSequence := range arrowSequences {
			from := arrowpad[ARROWPAD_START_ROW][ARROWPAD_START_COLUMN]
			arrowSequenceSum := 0
			for _, turn := range arrowSequence {
				to := arrowpadTile(arrowpad, turn)
				sum := doarrowrobotthing(cache, arrowpadsCount, arrowpad, from, to)
				from = to
				arrowSequenceSum += sum
			}
			sum = min(sum, arrowSequenceSum)
		}
		fmt.Println("Shortest path for whole code is", sum, "num part is", numPart)
		res += sum * numPart
	}
	fmt.Println("Result is", res)
}

type Cache map[int]map[string]int

func (c Cache) get(deepth int, from, to *Tile) (int, bool) {
	d, ok := c[deepth]
	if !ok || d == nil {
		return -1, false
	}
	res, ok := d[fmt.Sprintf("%d:%d => %d:%d", from.row, from.column, to.row, to.column)]
	if !ok {
		return -1, false
	}
	return res, true
}

func (c Cache) set(deepth int, from, to *Tile, res int) int {
	d, ok := c[deepth]
	if !ok || d == nil {
		c[deepth] = make(map[string]int)
	}
	c[deepth][fmt.Sprintf("%d:%d => %d:%d", from.row, from.column, to.row, to.column)] = res
	return res
}

func doarrowrobotthing(cache Cache, deepts int, arrowpad [][]*Tile, from, to *Tile) int {
	cached, ok := cache.get(deepts, from, to)
	if ok {
		return cached
	}

	shortestTurns, _, _ := findShortestTurns(arrowpad, from, to.r)
	if deepts == 1 {
		res := len(shortestTurns[0])
		return cache.set(deepts, from, to, res)
	}

	minSum := -1
	for _, turns := range shortestTurns {
		sum := 0
		from := arrowpad[ARROWPAD_START_ROW][ARROWPAD_START_COLUMN]
		for _, turn := range turns {
			to := arrowpadTile(arrowpad, turn)
			sum += doarrowrobotthing(cache, deepts-1, arrowpad, from, to)
			from = to
		}
		if minSum == -1 {
			minSum = sum
		} else {
			minSum = min(minSum, sum)
		}
	}
	return cache.set(deepts, from, to, minSum)
}

func findShortestTurns(
	matrix [][]*Tile,
	from *Tile,
	to rune,
) (turnsOptions [][]rune, tiles [][]*Tile, toTile *Tile) {
	paths, turns := dfs(matrix, from, to, []*Tile{}, []rune{})
	paths = shorten(paths)
	turns = shorten(turns)
	turns = appendToAll(turns, 'A')
	return turns, paths, paths[0][len(paths[0])-1]
}

func min(f, s int) int {
	if f < s {
		return f
	}
	return s
}

func arrowpadTile(arrowpad [][]*Tile, r rune) *Tile {
	for _, row := range arrowpad {
		for _, tile := range row {
			if tile != nil && tile.r == r {
				return tile
			}
		}
	}
	panic("Tile '" + string(r) + "' not found")
}
