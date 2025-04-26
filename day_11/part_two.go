package day_11

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// --- Part Two ---
// The Historians sure are taking a long time. To be fair, the infinite corridors are very large.
//
// How many stones would you have after blinking a total of 75 times?

func Part2(input []byte) {
	line := strings.Trim(string(input), "\n")
	stones := parse(line)
	blinks := 75

	cache := map[Pair]uint64{}
	count := big.NewInt(0)
	for _, stone := range stones {
		count.Add(count, new(big.Int).SetUint64(walk(uint64(stone), uint64(blinks), cache)))
	}

	fmt.Println("Result is", count)
}

type Pair struct {
	f uint64
	s uint64
}

func walk(it, deepth uint64, cache map[Pair]uint64) uint64 {
	if deepth == 0 {
		return 1
	}

	cached, exists := cache[Pair{it, deepth}]
	if exists {
		return cached
	}

	if it == 0 {
		res := walk(1, deepth-1, cache)
		cache[Pair{it, deepth}] = res
		return res
	}

	itStr := fmt.Sprintf("%d", it)
	if len(itStr)%2 == 0 {
		first, _ := strconv.Atoi(itStr[0 : len(itStr)/2])
		second, _ := strconv.Atoi(itStr[len(itStr)/2:])
		res := walk(uint64(first), deepth-1, cache) + walk(uint64(second), deepth-1, cache)
		cache[Pair{it, deepth}] = res
		return res
	}

	res := walk(it*2024, deepth-1, cache)
	cache[Pair{it, deepth}] = res
	return res
}
