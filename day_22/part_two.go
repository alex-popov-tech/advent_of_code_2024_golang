package day_22

import (
	"fmt"
	"strings"
)

func Part2(input []byte) {
	startNumbers := parse(string(input))
	secretNumbers := calcBuyerNumbers(2000, startNumbers)
	pricesByChunk := chunkedPrices(secretNumbers, 4)
	res := 0
	for k, v := range pricesByChunk {
		res = max(res, v)
		fmt.Printf("Chunk %s has price %d\n", k, v)
	}
	fmt.Println("-2,1,-1,3", pricesByChunk["-2,1,-1,3"])
	fmt.Println("Maximum bananas we can trade is", res)
}

func chunkedPrices(buyersNumberInfos [][]BuyerNumberInfo, chunkSize int) map[string]int {
	firstChunkPriceByBuyer := make([]map[string]int, len(buyersNumberInfos))
	for bi, buyerNumberInfos := range buyersNumberInfos {
		res := make(map[string]int)
		for i := chunkSize - 1; i < len(buyerNumberInfos); i++ {
			info := buyerNumberInfos[i]
			key := toPriceChangeStr(buyerNumberInfos[i-chunkSize+1 : i+1])
			// save only first occurrence
			if _, ok := res[key]; !ok {
				res[key] = info.price
			}
		}
		firstChunkPriceByBuyer[bi] = res
	}

	res := make(map[string]int)
	for _, m := range firstChunkPriceByBuyer {
		for k, v := range m {
			res[k] += v
		}
	}
	return res
}

func toPriceChangeStr(nums []BuyerNumberInfo) string {
	res := make([]string, len(nums))
	for i, num := range nums {
		res[i] = fmt.Sprintf("%d", num.diff)
	}
	return strings.Join(res, ",")
}

type BuyerNumberInfo struct{ secret, price, diff int }

func calcBuyerNumbers(quantity int, starts []int) [][]BuyerNumberInfo {
	res := make([][]BuyerNumberInfo, len(starts))
	for i, start := range starts {
		res[i] = make([]BuyerNumberInfo, quantity)
		prevSecret := start
		prevPrice := prevSecret % 10
		for j := range quantity {
			currSecret := calcSecretNumber(prevSecret)
			currPrice := currSecret % 10
			diff := currPrice - prevPrice
			res[i][j] = BuyerNumberInfo{currSecret, currPrice, diff}

			prevSecret = currSecret
			prevPrice = currPrice
		}
	}
	return res
}

func calcSecretNumber(prev int) int {
	tmp := prev * 64
	prev = mix(prev, tmp)
	prev = prune(prev)

	tmp = divround(prev, 32)
	prev = mix(prev, tmp)
	prev = prune(prev)

	tmp = prev * 2048
	prev = mix(prev, tmp)
	prev = prune(prev)
	return prev
}

func max(f, s int) int {
	if f >= s {
		return f
	} else {
		return s
	}
}
