package day_13

import (
	"fmt"
	"strings"
	"sync"
)

// As you go to win the first prize, you discover that the claw is nowhere near where you expected it would be. Due to a unit conversion error in your measurements, the position of every prize is actually 10000000000000 higher on both the X and Y axis!
//
// Add 10000000000000 to the X and Y position of every prize. After making this change, the example above would now look like this:
//
// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=10000000008400, Y=10000000005400
//
// Button A: X+26, Y+66
// Button B: X+67, Y+21
// Prize: X=10000000012748, Y=10000000012176
//
// Button A: X+17, Y+86
// Button B: X+84, Y+37
// Prize: X=10000000007870, Y=10000000006450
//
// Button A: X+69, Y+23
// Button B: X+27, Y+71
// Prize: X=10000000018641, Y=10000000010279
// Now, it is only possible to win a prize on the second and fourth claw machines. Unfortunately, it will take many more than 100 presses to do so.
//
// Using the corrected prize coordinates, figure out how to win as many prizes as possible. What is the fewest tokens you would have to spend to win all possible prizes?

func Part2(input []byte) {
	line := strings.Trim(string(input), "\n")
	cases := parse2(line)
	minPrices := make([]uint, len(cases))
	group := sync.WaitGroup{}
	for i, c := range cases {
		group.Add(1)
		go func(i int, c Case) {
			defer group.Done()
			fmt.Printf("Finding min price for case %+v\n", c)
			cMinPrice := getMinPriceForCase(c)
			if cMinPrice == -1 {
				fmt.Printf("No solution for case %+v\n", c)
				return
			}
			fmt.Printf("Solution for case %+v is %d\n", c, cMinPrice)
			minPrices[i] = uint(cMinPrice)
		}(i, c)
	}
	group.Wait()

	res := uint(0)
	for _, minPrice := range minPrices {
		res += minPrice
	}
	fmt.Println("Result is", res)
}

func getMinPriceForCase(c Case) int {
	// 1. take 1 combination with max A min B presses
	maxAminBcomb := getCaseCombination(c.prize, c.a, c.b)
	if maxAminBcomb == nil {
		return -1
	}
	maxAminBprice := c.a.price*maxAminBcomb.f + c.b.price*maxAminBcomb.s
	// 2. take 1 combination with max B min A presses
	maxBminAcomb := getCaseCombination(c.prize, c.b, c.a)
	if maxBminAcomb == nil {
		return -1
	}
	maxBminAprice := c.a.price*maxAminBcomb.f + c.b.price*maxAminBcomb.s
	// 2. return whichever is cheaper
	return min(maxAminBprice, maxBminAprice)
}

func getCaseCombination(prize pos, max, mix button) *Pair[int] {
	fbPressCount := howMuchPressesFit(prize, max)
	for ; fbPressCount >= 0; fbPressCount-- {
		restX := prize.x - fbPressCount*max.xdiff
		restY := prize.y - fbPressCount*max.ydiff
		divX := restX / mix.xdiff
		divY := restY / mix.ydiff
		restOfDivisionX := restX % mix.xdiff
		restOfDivisionY := restY % mix.ydiff
		if restOfDivisionX == 0 && restOfDivisionY == 0 && divX == divY {
			return &Pair[int]{f: fbPressCount, s: restX / mix.xdiff}
		}
	}
	return nil
}

func parse2(input string) []Case {
	cases := parse(input)
	for i := 0; i < len(cases); i++ {
		cases[i].prize.x += 10000000000000
		cases[i].prize.y += 10000000000000
	}
	return cases
}

type Pair[T any] struct {
	f T
	s T
}

func howMuchPressesFit(p pos, b button) int {
	xFit := p.x / b.xdiff
	yFit := p.y / b.ydiff
	return min(xFit, yFit)
}

func min(f, s int) int {
	if f < s {
		return f
	}
	return s
}
