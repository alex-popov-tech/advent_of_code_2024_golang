package day_13

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// --- Day 13: Claw Contraption ---
// Next up: the lobby of a resort on a tropical island. The Historians take a moment to admire the hexagonal floor tiles before spreading out.
//
// Fortunately, it looks like the resort has a new arcade! Maybe you can win some prizes from the claw machines?
//
// The claw machines here are a little unusual. Instead of a joystick or directional buttons to control the claw, these machines have two buttons labeled A and B. Worse, you can't just put in a token and play; it costs 3 tokens to push the A button and 1 token to push the B button.
//
// With a little experimentation, you figure out that each machine's buttons are configured to move the claw a specific amount to the right (along the X axis) and a specific amount forward (along the Y axis) each time that button is pressed.
//
// Each machine contains one prize; to win the prize, the claw must be positioned exactly above the prize on both the X and Y axes.
//
// You wonder: what is the smallest number of tokens you would have to spend to win as many prizes as possible? You assemble a list of every machine's button behavior and prize location (your puzzle input). For example:
//
// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400
//
// Button A: X+26, Y+66
// Button B: X+67, Y+21
// Prize: X=12748, Y=12176
//
// Button A: X+17, Y+86
// Button B: X+84, Y+37
// Prize: X=7870, Y=6450
//
// Button A: X+69, Y+23
// Button B: X+27, Y+71
// Prize: X=18641, Y=10279
// This list describes the button configuration and prize location of four different claw machines.
//
// For now, consider just the first claw machine in the list:
//
// Pushing the machine's A button would move the claw 94 units along the X axis and 34 units along the Y axis.
// Pushing the B button would move the claw 22 units along the X axis and 67 units along the Y axis.
// The prize is located at X=8400, Y=5400; this means that from the claw's initial position, it would need to move exactly 8400 units along the X axis and exactly 5400 units along the Y axis to be perfectly aligned with the prize in this machine.
// The cheapest way to win the prize is by pushing the A button 80 times and the B button 40 times. This would line up the claw along the X axis (because 80*94 + 40*22 = 8400) and along the Y axis (because 80*34 + 40*67 = 5400). Doing this would cost 80*3 tokens for the A presses and 40*1 for the B presses, a total of 280 tokens.
//
// For the second and fourth claw machines, there is no combination of A and B presses that will ever win a prize.
//
// For the third claw machine, the cheapest way to win the prize is by pushing the A button 38 times and the B button 86 times. Doing this would cost a total of 200 tokens.
//
// So, the most prizes you could possibly win is two; the minimum tokens you would have to spend to win all (two) prizes is 480.
//
// You estimate that each button would need to be pressed no more than 100 times to win a prize. How else would someone be expected to play?
//
// Figure out how to win as many prizes as possible. What is the fewest tokens you would have to spend to win all possible prizes?

func Part1(input []byte) {
	line := strings.Trim(string(input), "\n")
	cases := parse(line)
	res := 0
	for _, c := range cases {
		combinations := getCaseCombinations(c)
		if len(combinations) == 0 {
			fmt.Printf("No solution for case %+v\n", c)
			continue
		}
		// there is only one comb for first part
		comb := combinations[0]
		fmt.Printf(
			"Solution for case %+v is %d\n",
			c,
			comb.aPresses*c.a.price+comb.bPresses*c.b.price,
		)
		res += comb.aPresses*c.a.price + comb.bPresses*c.b.price
	}
	fmt.Println("Result is", res)
}

func getCaseCombinations(c Case) []combination {
	combinations := []combination{}
	for aPressCount := 0; c.a.press(aPressCount).lessOrEq(c.prize); aPressCount++ {
		posAfterPressingA := c.a.press(aPressCount)

		for bPressCount := 0; posAfterPressingA.add(c.b.press(bPressCount)).lessOrEq(c.prize); bPressCount++ {
			posAfterPressingAandB := posAfterPressingA.add(c.b.press(bPressCount))

			if posAfterPressingAandB.eq(c.prize) {
				combinations = append(
					combinations,
					combination{aPresses: aPressCount, bPresses: bPressCount},
				)
			}
		}
	}

	return combinations
}

type combination struct {
	aPresses int
	bPresses int
}

func (comb combination) price(aPrice, bPrice int) int {
	return comb.aPresses*aPrice + comb.bPresses*bPrice
}

type pos struct {
	x int
	y int
}

func (p pos) lessOrEq(b pos) bool {
	return p.x <= b.x && p.y <= b.y
}

func (p pos) eq(pr pos) bool {
	return p.x == pr.x && p.y == pr.y
}

func (p pos) add(b pos) pos {
	return pos{
		x: p.x + b.x,
		y: p.y + b.y,
	}
}

func (p pos) mul(b button, times int) pos {
	return pos{
		x: p.x + b.xdiff*times,
		y: p.y + b.ydiff*times,
	}
}

type button struct {
	name  string
	xdiff int
	ydiff int
	price int
}

func (b button) press(count int) pos {
	return pos{
		x: b.xdiff * count,
		y: b.ydiff * count,
	}
}

type Case struct {
	a     button
	b     button
	prize pos
}

func parse(input string) []Case {
	chunks := strings.Split(input, "\n\n")
	reBtn := regexp.MustCompile(`Button ([A-B]): X\+(\d+), Y\+(\d+)`)
	rePrize := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	var cases []Case

	for _, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		matchA := reBtn.FindStringSubmatch(lines[0])
		matchB := reBtn.FindStringSubmatch(lines[1])
		matchP := rePrize.FindStringSubmatch(lines[2])

		a := button{
			name:  matchA[1],
			xdiff: atoi(matchA[2]),
			ydiff: atoi(matchA[3]),
			price: 3,
		}
		b := button{
			name:  matchB[1],
			xdiff: atoi(matchB[2]),
			ydiff: atoi(matchB[3]),
			price: 1,
		}
		p := pos{
			x: atoi(matchP[1]),
			y: atoi(matchP[2]),
		}

		cases = append(cases, Case{a: a, b: b, prize: p})
	}

	return cases
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
