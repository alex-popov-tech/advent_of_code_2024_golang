package day_9

import (
	"fmt"
	"strings"
)

// --- Part Two ---
// Upon completion, two things immediately become clear. First, the disk definitely has a lot more contiguous free space, just like the amphipod hoped. Second, the computer is running much more slowly! Maybe introducing all of that file system fragmentation was a bad idea?
//
// The eager amphipod already has a new plan: rather than move individual blocks, he'd like to try compacting the files on his disk by moving whole files instead.
//
// This time, attempt to move whole files to the leftmost span of free space blocks that could fit the file. Attempt to move each file exactly once in order of decreasing file ID number starting with the file with the highest file ID number. If there is no span of free space to the left of a file that is large enough to fit the file, the file does not move.
//
// The first example from above now proceeds differently:
//
// 00...111...2...333.44.5555.6666.777.888899
// 0099.111...2...333.44.5555.6666.777.8888..
// 0099.1117772...333.44.5555.6666.....8888..
// 0099.111777244.333....5555.6666.....8888..
// 00992111777.44.333....5555.6666.....8888..
// The process of updating the filesystem checksum is the same; now, this example's checksum would be 2858.
//
// Start over, now compacting the amphipod's hard drive using this new method instead. What is the resulting filesystem checksum?

func Part2(input []byte) {
	line := strings.Trim(string(input), "\n")
	fmt.Printf("'%s'\n", line)
	blocks := toBlocks(line)
	for _, it := range blocks {
		fmt.Printf("%v\n", it)
	}

	lastOffset := 0
	for lastOffset < len(blocks) {
		lastFileBlockIndex := lastNonEmpty2(blocks, lastOffset)
		requiredSpace := len(blocks[lastFileBlockIndex])

		blockWithFreeSpaceIndex := firstWithSpace(blocks, requiredSpace)
		// no blocks with enough room
		if blockWithFreeSpaceIndex != -1 && blockWithFreeSpaceIndex < lastFileBlockIndex {
			blocks[blockWithFreeSpaceIndex].fill(blocks[lastFileBlockIndex])
			blocks[lastFileBlockIndex].clear()
		}

		lastOffset = len(blocks) - lastFileBlockIndex
	}

	fmt.Println("=======================")
	for _, it := range blocks {
		fmt.Printf("%v\n", it)
	}

	res := 0
	i := 0
	for _, block := range blocks {
		for _, it := range block {
			fmt.Printf("i * it = %d * %d = %d ( sum is %d )\n", i, it, i*it, res)
			if it > 0 {
				res += i * it
			}
			i++
		}
	}
	fmt.Println("Result is", res)
}

func firstWithSpace(blocks []Block, freeAmount int) (index int) {
	for i, block := range blocks {
		_, free := block.freeSpace()
		if free >= freeAmount {
			return i
		}
	}
	return -1
}

func lastNonEmpty2(blocks []Block, offset int) (index int) {
	for i := len(blocks) - 1 - offset; i >= 0; i-- {
		if !blocks[i].isEmpty() {
			return i
		}
	}
	return -1
}

type Block []int

func NewBlock(item, count int) Block {
	res := make([]int, count)
	for i := range res {
		res[i] = item
	}
	return res
}

func (b Block) freeSpace() (has bool, howMuch int) {
	count := 0
	for i := range b {
		if b[len(b)-1-i] != -1 {
			return count != 0, count
		}
		count++
	}
	return count != 0, count
}

func (b Block) fill(block Block) {
	hasFree, freeCount := b.freeSpace()
	if !hasFree || freeCount < len(block) {
		panic("not enough free space")
	}
	for i, j := 0, 0; j < len(block); i++ {
		if b[i] == -1 {
			b[i] = block[j]
			j++
		}
	}
}

func (b Block) clear() {
	for i := range b {
		b[i] = -1
	}
}

func (b Block) isEmpty() bool {
	_, freeCount := b.freeSpace()
	return freeCount == len(b)
}

func (b Block) String() string {
	_, free := b.freeSpace()
	return fmt.Sprintf("Block with len=%d, free=%d, contents=%v", len(b), free, []int(b))
}

func toBlocks(line string) []Block {
	res := []Block{}
	id := 0
	for i, r := range line {
		isFile := i%2 == 0
		if isFile {
			block := NewBlock(id, parseInt(r))
			res = append(res, block)
			id++
		} else {
			block := NewBlock(-1, parseInt(r))
			res = append(res, block)
		}
	}
	return res
}
