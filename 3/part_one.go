package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mainn() {
	input, _ := os.ReadFile("./3/input.txt")
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	acc := 0
	for _, expr := range re.FindAllString(string(input), -1) {
		nums := strings.Split(expr[4:len(expr)-1], ",")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])
		acc += first * second
	}
	fmt.Println("Result is ", acc)
}
