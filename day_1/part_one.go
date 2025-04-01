package day_1

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func Part1(inputPath string) {
	const INPUT_LENGTH = 1000
	input, _ := os.ReadFile(inputPath)
	lefts := make([]int, INPUT_LENGTH)
	rights := make([]int, INPUT_LENGTH)

	for i, row := range bytes.Split(input, []byte("\n")) {
		if len(row) == 0 {
			continue
		}
		nums := bytes.Fields(row)
		left, _ := strconv.Atoi(string(nums[0]))
		right, _ := strconv.Atoi(string(nums[1]))
		lefts[i] = left
		rights[i] = right
	}

	sort.Ints(lefts)
	sort.Ints(rights)

	diffs := 0
	for i, left := range lefts {
		diffs += int(math.Abs(float64(rights[i] - left)))
	}
	fmt.Println("Result is ", diffs)
}
