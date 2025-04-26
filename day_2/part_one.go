package day_2

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(input []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	validCount := 0

	for scanner.Scan() {
		numsAsStrings := strings.Fields(scanner.Text())
		ints := make([]int, len(numsAsStrings))
		for i, numAsStr := range numsAsStrings {
			ints[i], _ = strconv.Atoi(numAsStr)
		}
		isValidDiffs := isValidDiffs(ints)
		isAllDecreasing := isAllDecreasing(ints)
		isAllIncreasing := isAllIncreasing(ints)
		if isValidDiffs && (isAllDecreasing || isAllIncreasing) {
			validCount++
		}
	}
	fmt.Println("Result it ", validCount)
}

func isValidDiffs(val []int) bool {
	for i := 1; i < len(val); i++ {
		prev := val[i-1]
		curr := val[i]
		diff := int(math.Abs(float64(curr - prev)))
		if diff > 3 || diff < 1 {
			return false
		}
	}
	return true
}

func isAllIncreasing(val []int) bool {
	for i := 1; i < len(val); i++ {
		if val[i-1] < val[i] {
			return false
		}
	}
	return true
}

func isAllDecreasing(val []int) bool {
	for i := 1; i < len(val); i++ {
		if val[i-1] > val[i] {
			return false
		}
	}
	return true
}
