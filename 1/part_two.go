package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 3   4
// 4   3
// 2   5
// 1   3
// 3   9
// 3   3
// For these example lists, here is the process of finding the similarity score:
//
// The first number in the left list is 3. It appears in the right list three times, so the similarity score increases by 3 * 3 = 9.
// The second number in the left list is 4. It appears in the right list once, so the similarity score increases by 4 * 1 = 4.
// The third number in the left list is 2. It does not appear in the right list, so the similarity score does not increase (2 * 0 = 0).
// The fourth number, 1, also does not appear in the right list.
// The fifth number, 3, appears in the right list three times; the similarity score increases by 9.
// The last number, 3, appears in the right list three times; the similarity score again increases by 9.
// So, for these example lists, the similarity score at the end of this process is 31 (9 + 4 + 0 + 0 + 9 + 9).

const INPUT_LENGTH = 1000

func main() {
	file, _ := os.Open("./1/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lefts := make([]int, INPUT_LENGTH)
	rights := make(map[int]int, INPUT_LENGTH)

	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		left, _ := strconv.Atoi(string(nums[0]))
		right, _ := strconv.Atoi(string(nums[1]))
		// left column just append
		lefts[lineIndex] = left
		lineIndex++
		// right column count
		rights[right]++
	}

	similarityScore := 0
	for _, left := range lefts {
		similarityScore += (left * rights[left])
	}
	fmt.Println("Result is ", similarityScore)
}
