package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("./3/input.txt")
	inputRunes := []rune(string(input))

	state := "valid"
	processedInput := make([]rune, len(input)/4) // at least 1/4 is valid(?)

	for i := 0; i < len(inputRunes); {
		// if expr - set flag
		if inputRunes[i] == 'd' {
			if "do()" == string(inputRunes[i:i+4]) {
				state = "valid"
				i += 4
			}
			if "don't()" == string(inputRunes[i:i+7]) {
				state = "invalid"
				i += 7
			}
			continue
		}

		// if not expr and valid - append
		if state == "valid" {
			processedInput = append(processedInput, inputRunes[i])
		}
		i++
	}

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	acc := 0
	for _, expr := range re.FindAllString(string(processedInput), -1) {
		nums := strings.Split(expr[4:len(expr)-1], ",")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])
		acc += first * second
	}
	fmt.Println("Result is ", acc)
}
