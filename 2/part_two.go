package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./2/input.txt")
	scanner := bufio.NewScanner(file)
	validCount := 0

	for scanner.Scan() {
		numsAsStrings := strings.Fields(scanner.Text())
		ints := make([]int, len(numsAsStrings))
		for i, numAsStr := range numsAsStrings {
			ints[i], _ = strconv.Atoi(numAsStr)
		}

		if isValidTolerated(ints) {
			validCount++
		}

	}
	fmt.Println("Result it ", validCount)
}

func isValidTolerated(ints []int) bool {
	if isValid(ints) {
		return true
	}
	for i := 0; i < len(ints); i++ {
		toleratedInts := []int{}
		toleratedInts = append(toleratedInts, ints[:i]...)
		toleratedInts = append(toleratedInts, ints[i+1:]...)
		if isValid(toleratedInts) {
			return true
		}
	}
	return false
}

func isValid(ints []int) bool {
	isValidDiffs := isValidDiffs(ints)
	isAllDecreasing := isAllDecreasing(ints)
	isAllIncreasing := isAllIncreasing(ints)
	if isValidDiffs && (isAllDecreasing || isAllIncreasing) {
		return true
	}
	return false
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
