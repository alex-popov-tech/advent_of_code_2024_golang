package day_23

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

// --- Day 23: LAN Party ---
// As The Historians wander around a secure area at Easter Bunny HQ,
// you come across posters for a LAN party scheduled for today!
// Maybe you can find it; you connect to a nearby datalink port and
// download a map of the local network (your puzzle input).
//
// The network map provides a list of every connection between two computers. For example:
//
// kh-tc
// qp-kh
// de-cg
// ka-co
// yn-aq
// qp-ub
// cg-tb
// vc-aq
// tb-ka
// wh-tc
// yn-cg
// kh-ub
// ta-co
// de-co
// tc-td
// tb-wq
// wh-td
// ta-ka
// td-qp
// aq-cg
// wq-ub
// ub-vc
// de-ta
// wq-aq
// wq-vc
// wh-yn
// ka-de
// kh-ta
// co-tc
// wh-qp
// tb-vc
// td-yn
// Each line of text in the network map represents a single connection; the line
// kh-tc represents a connection between the computer named kh and the computer named tc.
// Connections aren't directional; tc-kh would mean exactly the same thing.
//
// LAN parties typically involve multiplayer games, so maybe you can locate it by
// finding groups of connected computers. Start by looking for sets of three
// computers where each computer in the set is connected to the other two computers.
//
// In this example, there are 12 such sets of three inter-connected computers:
//
// aq,cg,yn
// aq,vc,wq
// co,de,ka
// co,de,ta
// co,ka,ta
// de,ka,ta
// kh,qp,ub
// qp,td,wh
// tb,vc,wq
// tc,td,wh
// td,wh,yn
// ub,vc,wq
// If the Chief Historian is here, and he's at the LAN party, it would be best
// to know that right away. You're pretty sure his computer's name starts with
// t, so consider only sets of three computers where at least one computer's
// name starts with t. That narrows the list down to 7 sets of three inter-connected
// computers:
//
// co,de,ta
// co,ka,ta
// de,ka,ta
// qp,td,wh
// tb,vc,wq
// tc,td,wh
// td,wh,yn
// Find all the sets of three inter-connected computers. How many contain at least one computer with a name that starts with t?

func Part1(input []byte) {
	data := parse(string(input))
	for computer, connected := range data {
		fmt.Println("Computer", computer, "is connected to", keys(connected))
	}

	threeInterConnected := map[string]struct{}{}
	for first, connected := range data {
		for second := range connected {
			intersection := intersection(connected, data[second])
			fmt.Println("Intersection between", first, "and", second, "is", intersection)
			for _, third := range intersection {
				key := strings.Join(sortt([]string{first, second, third}), ",")
				threeInterConnected[key] = struct{}{}
			}
		}
	}
	for k := range threeInterConnected {
		fmt.Println("Three inter-connected computers are", k)
	}

	filtered := filter(threeInterConnected, func(s string) bool {
		computers := strings.Split(s, ",")
		for _, computer := range computers {
			if strings.HasPrefix(computer, "t") {
				return true
			}
		}
		return false
	})
	for k := range filtered {
		fmt.Println(
			"Three inter-connected computers where at least one computer's name starts with t are",
			k,
		)
	}
	fmt.Println(
		"There are",
		len(filtered),
		"sets of three inter-connected computers where at least one computer's name starts with t",
	)
}

type (
	Set                map[string]struct{}
	ConnectedComputers map[string]Set
)

func parse(input string) ConnectedComputers {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	res := ConnectedComputers{}
	for _, it := range lines {
		parts := strings.Split(it, "-")
		from := parts[0]
		to := parts[1]
		if res[from] == nil {
			res[from] = Set{}
		}
		if res[to] == nil {
			res[to] = Set{}
		}
		res[from][to] = struct{}{}
		res[to][from] = struct{}{}
	}
	return res
}

func keys[K comparable, V any](m map[K]V) []K {
	res := []K{}
	for k := range m {
		res = append(res, k)
	}
	return res
}

func sortt(s []string) []string {
	res := slices.Clone(s)
	sort.Strings(res)
	return res
}

func intersection(f, s map[string]struct{}) []string {
	m := map[string]int{}
	for fs := range f {
		m[fs]++
	}
	for ss := range s {
		m[ss]++
	}
	res := []string{}
	for k, v := range m {
		if v == 2 {
			res = append(res, k)
		}
	}
	return res
}

func filter(m map[string]struct{}, cond func(string) bool) []string {
	res := []string{}
	for k := range m {
		if cond(k) {
			res = append(res, k)
		}
	}
	return res
}
