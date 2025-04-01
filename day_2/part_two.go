package day_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2(inputPath string) {
	file, _ := os.Open(inputPath)
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
