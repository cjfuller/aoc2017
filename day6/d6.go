package main

import (
	"fmt"
	"strconv"
	"strings"
)

func configString(banks []int) string {
	strs := []string{}
	for _, bank := range banks {
		strs = append(strs, strconv.Itoa(bank))
	}
	return strings.Join(strs, ",")
}

func maxIndex(banks []int) int {
	max := -1
	maxIdx := 0
	for i, val := range banks {
		if val > max {
			maxIdx = i
			max = val
		}
	}
	return maxIdx
}

func solveP1P2(banks []int) (int, int) {
	count := 0
	visitedConfigs := make(map[string]bool)
	countConfigs := make(map[string]int)
	config := configString(banks)
	visitedConfigs[config] = true
	countConfigs[config] = count
	for {
		idx := maxIndex(banks)
		toRedistrib := banks[idx]
		banks[idx] = 0
		nextBankIdx := idx + 1
		for i := 0; i < toRedistrib; i++ {
			banks[(i+nextBankIdx)%len(banks)]++
		}
		count++
		config := configString(banks)
		if visitedConfigs[config] {
			return count, (count - countConfigs[config])
		}
		visitedConfigs[config] = true
		countConfigs[config] = count
	}
}

func main() {
	count, cycleSize := solveP1P2([]int{0, 2, 7, 0})
	fmt.Printf("%d, %d\n", count, cycleSize)
	input := []int{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}
	count, cycleSize = solveP1P2(input)
	fmt.Printf("%d, %d\n", count, cycleSize)
}
