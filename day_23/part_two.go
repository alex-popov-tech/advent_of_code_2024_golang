package day_23

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

// --- Part Two ---
// There are still way too many results to go through them all.
// You'll have to find the LAN party another way and go there yourself.
//
// Since it doesn't seem like any employees are around, you figure
// they must all be at the LAN party. If that's true, the LAN party
// will be the largest set of computers that are all connected to
// each other. That is, for each computer at the LAN party, that
// computer will have a connection to every other computer at the LAN party.
//
// In the above example, the largest set of computers that are all
// connected to each other is made up of co, de, ka, and ta. Each
// computer in this set has a connection to every other computer in the set:
//
// ka-co
// ta-co
// de-co
// ta-ka
// de-ta
// ka-de
// The LAN party posters say that the password to get into the LAN party
// is the name of every computer at the LAN party, sorted alphabetically,
// then joined together with commas. (The people running the LAN party
// are clearly a bunch of nerds.) In this example, the password would
// be co,de,ka,ta.
//
// What is the password to get into the LAN party?

func Part2(input []byte) {
	data := parse(string(input))
	computers := keys(data)
	sort.Strings(computers)
	groups := findAllConnected([]string{}, computers, data)
	pass := ""
	max := 0
	for _, group := range groups {
		fmt.Println(strings.Join(group, ","))
		size := len(group)
		if size > max {
			max = size
			pass = strings.Join(group, ",")
		}
	}
	fmt.Println("Password is ", pass)
}

func findAllConnected(
	group []string,
	candidates []string,
	allConnected ConnectedComputers,
) [][]string {
	res := [][]string{}
	res = append(res, group)

	for i, candidate := range candidates {
		if connectedToAll(candidate, group, allConnected) {
			res = append(
				res,
				findAllConnected(
					append(slices.Clone(group), candidate),
					candidates[i+1:],
					allConnected,
				)...)
		}
	}
	return res
}

func connectedToAll(computer string, group []string, allConnected ConnectedComputers) bool {
	for _, member := range group {
		if _, ok := allConnected[member][computer]; !ok {
			return false
		}
	}
	return true
}
