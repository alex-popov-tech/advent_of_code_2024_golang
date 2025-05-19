package day_19

import (
	"fmt"
	"strings"
)

// --- Part Two ---
// The staff don't really like some of the towel arrangements you came up with.
// To avoid an endless cycle of towel rearrangement, maybe you should just give them every possible option.
//
// Here are all of the different ways the above example's designs can be made:
//
// brwrr can be made in two different ways: b, r, wr, r or br, wr, r.
//
// bggr can only be made with b, g, g, and r.
//
// gbbr can be made 4 different ways:
//
// g, b, b, r
// g, b, br
// gb, b, r
// gb, br
// rrbgbr can be made 6 different ways:
//
// r, r, b, g, b, r
// r, r, b, g, br
// r, r, b, gb, r
// r, rb, g, b, r
// r, rb, g, br
// r, rb, gb, r
// bwurrg can only be made with bwu, r, r, and g.
//
// brgr can be made in two different ways: b, r, g, r or br, g, r.
//
// ubwu and bbrgwb are still impossible.
//
// Adding up all of the ways the towels in this example could be arranged into the desired designs yields 16 (2 + 1 + 4 + 6 + 1 + 2).
// They'll let you into the onsen as soon as you have the list.
// What do you get if you add up the number of different ways you could make each design?

func Part2(input []byte) {
	towels, designs := parse(string(input))

	count := 0
	cache := make(map[string]int)
	for i, design := range designs {
		actualTowels := Filter(towels, func(towel string) bool {
			return strings.Contains(design, towel)
		})
		fmt.Printf(
			"For design %d) '%s' there are %d/%d available towels, calculating matches...\n",
			i,
			design,
			len(actualTowels),
			len(towels),
		)
		matches := collectTowelMatches(design, actualTowels, cache)
		fmt.Printf("%d matches are available\n", matches)
		count += matches
	}

	fmt.Printf("Possible designs are %d\n", count)
}

func collectTowelMatches(
	design string,
	towels []string,
	cache map[string]int,
) int {
	if res, ok := cache[design]; ok {
		return res
	}

	if design == "" {
		return 1
	}

	matches := 0
	for _, towel := range towels {
		if startsWith(design, towel) {
			res := collectTowelMatches(
				design[len(towel):],
				towels,
				cache,
			)
			matches += res
		}
	}

	cache[design] = matches
	return matches
}

func Filter(input []string, cond func(string) bool) []string {
	var out []string
	for _, s := range input {
		if cond(s) {
			out = append(out, s)
		}
	}
	return out
}
