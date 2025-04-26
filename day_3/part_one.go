package day_3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input []byte) {
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
