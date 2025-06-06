package day_24

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// so for this one i tried to came up with solution myself, but i might
// be too stupid for that, so i went to AI and ask it ( gemini, openai, deepseek )
// and they all generated not working solutions again and again
// then i asked to came up with theoretical solving algorithm, and they
// failed again
// then i found https://old.reddit.com/r/adventofcode/comments/1hla5ql/2024_day_24_part_2_a_guide_on_the_idea_behind_the/
// and asked AI to rewrite kotlin to golang, and it worked perfectly
// i still have no idea how this works, but i don't case so much to spend more time on it
func Part2(input []byte) {
	registers, gates := parseInput(input)
	// Part 2:
	// Need to re-parse, since solvePart1 mutated neither registers nor gates.
	// But solvePart2 also works on its own copies, so we can call directly:
	part2 := solvePart2(registers, gates)
	fmt.Println(part2)
}

type Operation2 int

const (
	AND Operation2 = iota
	OR
	XOR
)

func parseOp(s string) Operation2 {
	switch s {
	case "AND":
		return AND
	case "OR":
		return OR
	case "XOR":
		return XOR
	default:
		panic("unknown op: " + s)
	}
}

type Gate struct {
	a, b, c string
	op      Operation2
}

// parseInput reads from r, expects two sections separated by a blank line.
// First section: lines of "wire: 0/1" (register initial values).
// Second section: lines of "A OP B -> C".
func parseInput(data []byte) (map[string]int, []*Gate) {
	parts := strings.Split(strings.Trim(string(data), "\n"), "\n\n")
	if len(parts) != 2 {
		panic("input format: two sections separated by blank line")
	}

	registers := make(map[string]int)
	for _, line := range strings.Split(parts[0], "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// "x00: 1"
		kv := strings.Split(line, ": ")
		registers[kv[0]] = atoi(kv[1])
	}

	var gates []*Gate
	for _, line := range strings.Split(parts[1], "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// "y04 XOR x04 -> cwr"
		fields := strings.Fields(line)
		// fields[0] = A, fields[1]=OP, fields[2]=B, fields[3]="->", fields[4]=C
		op := parseOp(fields[1])
		g := &Gate{
			a:  fields[0],
			op: op,
			b:  fields[2],
			c:  fields[4],
		}
		gates = append(gates, g)
	}

	return registers, gates
}

// run evaluates all gates in topological order (simple simulation).
// It returns the concatenated 'z' wires as a single integer.
func run(gates []*Gate, registers map[string]int) int64 {
	// Make a copy of registers to avoid mutating the caller's map.
	reg := make(map[string]int)
	for k, v := range registers {
		reg[k] = v
	}

	pending := make([]*Gate, len(gates))
	copy(pending, gates)

	for len(pending) > 0 {
		next := pending[:0]
		progress := false

		for _, g := range pending {
			va, oka := reg[g.a]
			vb, okb := reg[g.b]
			if !oka || !okb {
				next = append(next, g)
				continue
			}
			var out int
			switch g.op {
			case AND:
				out = va & vb
			case OR:
				out = va | vb
			case XOR:
				out = va ^ vb
			}
			reg[g.c] = out
			progress = true
		}

		if !progress {
			// No gate could be evaluated → cycle or missing input
			break
		}
		pending = next
	}

	// Collect all keys starting with 'z', sort, build bitstring, reverse, parse.
	var zs []string
	for k := range reg {
		if strings.HasPrefix(k, "z") {
			zs = append(zs, k)
		}
	}
	sort.Strings(zs) // lex order: z00, z01, ...

	var bits []rune
	for _, key := range zs {
		bits = append(bits, rune('0'+reg[key]))
	}
	// Reverse bit string
	for i, j := 0, len(bits)-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}
	bitStr := string(bits)
	if bitStr == "" {
		return 0
	}
	val, err := strconv.ParseInt(bitStr, 2, 64)
	if err != nil {
		panic(err)
	}
	return val
}

// getWiresAsInt collects wires of a given prefix letter ('x', 'y', or 'z') into one integer.
func getWiresAsInt(registers map[string]int, prefix byte) int64 {
	var keys []string
	for k := range registers {
		if k[0] == prefix {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys) // x00, x01, ...
	var bits []rune
	for _, k := range keys {
		bits = append(bits, rune('0'+registers[k]))
	}
	// reverse
	for i, j := 0, len(bits)-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}
	if len(bits) == 0 {
		return 0
	}
	val, err := strconv.ParseInt(string(bits), 2, 64)
	if err != nil {
		panic(err)
	}
	return val
}

// firstZThatUsesC recursively finds the nearest "zXX" that depends on wire c.
// It returns "zYY" where YY = the discovered z-index minus 1 (zero-padded).
func firstZThatUsesC(gates []*Gate, c string, visited map[string]bool) string {
	if visited[c] {
		return ""
	}
	visited[c] = true
	for _, g := range gates {
		if g.a == c || g.b == c {
			if strings.HasPrefix(g.c, "z") {
				// found zNN; compute z(NN-1) with zero-pad
				nn, err := strconv.Atoi(g.c[1:])
				if err != nil {
					panic(err)
				}
				if nn == 0 {
					return "z00" // no negative index; edge case
				}
				target := fmt.Sprintf("z%02d", nn-1)
				return target
			}
			// recurse
			if res := firstZThatUsesC(gates, g.c, visited); res != "" {
				return res
			}
		}
	}
	return ""
}

func solvePart2(initialRegs map[string]int, gates []*Gate) string {
	// Make deep copies, since we will mutate gate.c
	registers := make(map[string]int)
	for k, v := range initialRegs {
		registers[k] = v
	}
	// Copy gates slice and each Gate
	copyGates := make([]*Gate, len(gates))
	for i, g := range gates {
		copyGates[i] = &Gate{a: g.a, b: g.b, c: g.c, op: g.op}
	}

	// Step 1: identify nxz and xnz sets
	var nxz []*Gate
	var xnz []*Gate
	for _, g := range copyGates {
		if strings.HasPrefix(g.c, "z") && g.c != "z45" && g.op != XOR {
			nxz = append(nxz, g)
		}
		if !strings.HasPrefix(g.a, "x") && !strings.HasPrefix(g.a, "y") &&
			!strings.HasPrefix(g.b, "x") && !strings.HasPrefix(g.b, "y") &&
			!strings.HasPrefix(g.c, "z") && g.op == XOR {
			xnz = append(xnz, g)
		}
	}

	// Step 2: for each in xnz, find matching in nxz and swap their c
	for _, xi := range xnz {
		target := firstZThatUsesC(copyGates, xi.c, make(map[string]bool))
		for _, bj := range nxz {
			if bj.c == target {
				// swap
				temp := xi.c
				xi.c = bj.c
				bj.c = temp
				break
			}
		}
	}

	// Step 3: simulate to get actualZ, compute expectedZ, find falseCarry bit
	actualZ := run(copyGates, registers)
	xVal := getWiresAsInt(initialRegs, 'x')
	yVal := getWiresAsInt(initialRegs, 'y')
	expectedSum := xVal + yVal
	diff := expectedSum ^ actualZ
	// count trailing zero bits
	falseCarry := 0
	for diff != 0 && diff&1 == 0 {
		falseCarry++
		diff >>= 1
	}

	// Step 4: find the two gates with inputs ending in falseCarry that are XOR or AND
	falseSuffix := fmt.Sprintf("%02d", falseCarry)
	var lastTwo []*Gate
	for _, g := range copyGates {
		if (g.op == XOR || g.op == AND) &&
			strings.HasSuffix(g.a, falseSuffix) &&
			strings.HasSuffix(g.b, falseSuffix) {
			lastTwo = append(lastTwo, g)
		}
	}
	if len(lastTwo) != 2 {
		panic(fmt.Sprintf("expected exactly 2 gates for falseCarry, got %d", len(lastTwo)))
	}

	// Collect all eight wire names
	var allCs []string
	for _, g := range nxz {
		allCs = append(allCs, g.c)
	}
	for _, g := range xnz {
		allCs = append(allCs, g.c)
	}
	for _, g := range lastTwo {
		allCs = append(allCs, g.c)
	}

	sort.Strings(allCs)
	return strings.Join(allCs, ",")
}
