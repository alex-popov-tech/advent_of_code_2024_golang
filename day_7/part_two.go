package day_7

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --- Part Two ---
// The engineers seem concerned; the total calibration result you gave them is nowhere close to being within safety tolerances. Just then, you spot your mistake: some well-hidden elephants are holding a third type of operator.
//
// The concatenation operator (||) combines the digits from its left and right inputs into a single number. For example, 12 || 345 would become 12345. All operators are still evaluated left-to-right.
//
// Now, apart from the three equations that could be made true using only addition and multiplication, the above example has three more equations that can be made true by inserting operators:
//
// 156: 15 6 can be made true through a single concatenation: 15 || 6 = 156.
// 7290: 6 8 6 15 can be made true using 6 * 8 || 6 * 15.
// 192: 17 8 14 can be made true using 17 || 8 + 14.
// Adding up all six test values (the three that could be made before using only + and * plus the new three that can now be made by also using ||) produces the new total calibration result of 11387.
//
// Using your new knowledge of elephant hiding spots, determine which equations could possibly be true. What is their total calibration result?

func Part2(inputPath string) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	inputStr := strings.Trim(string(input), "\n")
	lines := strings.Split(inputStr, "\n")

	result := 0
	for _, line := range lines {
		expectedResult, operands := parse(line)
		operatorsCombinations := generateOperatorStrings2(len(operands) - 1)
		for _, operators := range operatorsCombinations {
			actualResult := calculate2(operands, operators)
			// fmt.Printf(
			// 	"expectedResult: %d, operands: %v, operators: %v, result: %d\n",
			// 	expectedResult,
			// 	operands,
			// 	operators,
			// 	result,
			// )
			if actualResult == expectedResult {
				fmt.Printf(
					"expectedResult: %d, operands: %v, operators: %v, result: %d\n",
					expectedResult,
					operands,
					operators,
					result,
				)
				result += expectedResult
				break
			}
		}
	}
	fmt.Println("Result is", result)
}

var operators = []string{"+", "*", "|"}

func generateOperatorStrings2(operatorsCount int) []string {
	if operatorsCount == 1 {
		return operators
	}

	result := []string{}
	for _, operator := range operators {
		options := generateOperatorStrings2(operatorsCount - 1)
		for _, option := range options {
			result = append(result, operator+option)
		}
	}
	return result
}

func calculate2(operands []int, operators string) int {
	result := operands[0]
	for i, operator := range operators {
		switch operator {
		case '+':
			result += operands[i+1]
		case '*':
			result *= operands[i+1]
		default:
			squashed := fmt.Sprintf("%d%d", result, operands[i+1])
			result, _ = strconv.Atoi(squashed)
		}
	}
	return result
}
