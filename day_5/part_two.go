package day_5

import (
	"fmt"
	"slices"
	"strings"
)

// --- Part Two ---
// While the Elves get to work printing the correctly-ordered updates, you have a little time to fix the rest of them.
//
// For each of the incorrectly-ordered updates, use the page ordering rules to put the page numbers in the right order. For the above example, here are the three incorrectly-ordered updates and their correct orderings:
//
// 75,97,47,61,53 becomes 97,75,47,61,53.
// 61,13,29 becomes 61,29,13.
// 97,13,75,29,47 becomes 97,75,47,29,13.
// After taking only the incorrectly-ordered updates and ordering them correctly, their middle page numbers are 47, 29, and 47. Adding these together produces 123.
//
// Find the updates which are not in the correct order. What do you get if you add up the middle page numbers after correctly ordering just those updates?

func Part2(input []byte) {
	inputStr := strings.Trim(string(input), "\n")
	parts := strings.Split(inputStr, "\n\n")
	rules := parseRules(parts[0])
	fmt.Println("================================RULES")
	for pageNumber, rules := range rules {
		fmt.Printf(
			"page %d should be after:%v\n",
			pageNumber,
			rules,
		)
	}
	updates := parseUpdates(parts[1])
	fmt.Println("================================UPDATES")
	res := 0
	for _, pages := range updates {
		validPages := getValidPages(pages, rules)
		areEqual := slices.Equal(pages, validPages)
		if !areEqual {
			res += validPages[len(validPages)/2]
		}
		fmt.Printf("actual :%v|valid: %v|equal: %t\n", pages, validPages, areEqual)
	}
	fmt.Println("Result is", res)
}

func getValidPages(update []int, rules map[int][]int) []int {
	// we have all rules, which are 24 per page
	// we also have pages in update
	// we need to iterate over update pages, and for each page:
	// 1. grab all possible rules
	relevantRulesPerPage := map[int][]int{}
	for _, page := range update {
		allRulesForPage := rules[page]
		// 2. filter rules which are not present in update
		presentInUpdateRules := []int{}
		for _, potentialRule := range allRulesForPage {
			if slices.Contains(update, potentialRule) {
				presentInUpdateRules = append(presentInUpdateRules, potentialRule)
			}
		}
		// (now we have all actual rules for page in update)
		// 3. save those relevant rules per page to map
		relevantRulesPerPage[page] = presentInUpdateRules
	}
	// 4. sort pages by rules count
	res := slices.Clone(update)
	slices.SortFunc(res, func(p1, p2 int) int {
		return len(relevantRulesPerPage[p2]) - len(relevantRulesPerPage[p1])
	})
	// 5. Return sorter pages slice
	return res
}
