package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func isValidPassphraseP1(pp string) bool {
	words := strings.Split(pp, " ")
	uniq := make(map[string]bool)
	for _, word := range words {
		uniq[word] = true
	}
	return len(words) == len(uniq)
}

func countLetters(word string) map[rune]int {
	counts := make(map[rune]int)
	for _, ch := range word {
		counts[ch]++
	}
	return counts
}

func isValidPassphraseP2(pp string) bool {
	words := strings.Split(pp, " ")
	for i, word1 := range words {
		letterCounts1 := countLetters(word1)
		for j, word2 := range words {
			letterCounts2 := countLetters(word2)
			// TODO(colin): we could make this more efficient by checking both
			// directions at once and then not iterating over all pairs twice.
			if i == j || len(word1) != len(word2) {
				continue
			}

			flag2IsSubsetOf1 := true

			for letter, count := range letterCounts2 {
				if letterCounts1[letter] != count {
					flag2IsSubsetOf1 = false
					break
				}
			}

			if flag2IsSubsetOf1 {
				return false
			}
		}
	}
	return true
}

func countValidPassphrases(pps []string, validate func(string) bool) int {
	count := 0
	for _, pp := range pps {
		if validate(pp) {
			count++
		}
	}
	return count
}

func parsePassphraseList(fn string) []string {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	allPps, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(allPps)), "\n")
}

func solve() {
	fmt.Println("Part 1:")
	fmt.Printf("%d\n", countValidPassphrases(
		parsePassphraseList("./p1_test.txt"), isValidPassphraseP1))
	fmt.Printf("%d\n", countValidPassphrases(
		parsePassphraseList("./input.txt"), isValidPassphraseP1))
	fmt.Println("Part 2:")
	fmt.Printf("%d\n", countValidPassphrases(
		parsePassphraseList("./p2_test.txt"), isValidPassphraseP2))
	fmt.Printf("%d\n", countValidPassphrases(
		parsePassphraseList("./input.txt"), isValidPassphraseP2))
}

func main() {
	solve()
}
