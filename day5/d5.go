package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func numSteps(currPos int, offsets []int) int {
	count := 0
	for {
		if currPos < 0 || currPos >= len(offsets) {
			return count
		}
		count++
		currVal := offsets[currPos]
		offsets[currPos]++
		currPos += currVal
	}
}

func numStepsP2(currPos int, offsets []int) int {
	count := 0
	for {
		if currPos < 0 || currPos >= len(offsets) {
			return count
		}
		count++
		currVal := offsets[currPos]
		if currVal >= 3 {
			offsets[currPos]--
		} else {
			offsets[currPos]++
		}
		currPos += currVal
	}
}

func parseInput(fn string) []int {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	strNums := strings.Split(strings.TrimSpace(string(content)), "\n")
	result := []int{}
	for _, item := range strNums {
		conv, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		result = append(result, conv)
	}
	return result
}

func solveP1() {
	example := []int{0, 3, 0, 1, -3}
	fmt.Printf("%d\n", numSteps(0, example))
	fmt.Printf("%d\n", numSteps(0, parseInput("./input.txt")))
}

func solveP2() {
	example := []int{0, 3, 0, 1, -3}
	fmt.Printf("%d\n", numStepsP2(0, example))
	fmt.Printf("%d\n", numStepsP2(0, parseInput("./input.txt")))
}

func main() {
	solveP1()
	solveP2()
}
