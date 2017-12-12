package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Token int

const (
	GroupOpen Token = iota
	GroupClose
	GarbageOpen
	GarbageClose
	Comma
	Cancel
	Other
)

type Group struct {
	Children []*Group
}

func evalTokenStream(tokens chan Token) (int, int) {
	totalScore := 0
	garbageTotal := 0
	currGroupValue := 0
	readingGarbage := false
	for t := range tokens {
		switch t {
		case Cancel:
			_ = <-tokens
		case GroupOpen:
			if readingGarbage {
				garbageTotal++
			} else {
				currGroupValue++
			}
		case GroupClose:
			if readingGarbage {
				garbageTotal++
			} else {
				totalScore += currGroupValue
				currGroupValue--
			}
		case GarbageOpen:
			if readingGarbage {
				garbageTotal++
			}
			readingGarbage = true
		case GarbageClose:
			readingGarbage = false
		case Other:
			garbageTotal++
		case Comma:
			if readingGarbage {
				garbageTotal++
			}
		}
	}
	return totalScore, garbageTotal
}

func lex(source chan rune, dest chan Token) {
	for r := range source {
		switch r {
		case ',':
			dest <- Comma
		case '{':
			dest <- GroupOpen
		case '}':
			dest <- GroupClose
		case '<':
			dest <- GarbageOpen
		case '>':
			dest <- GarbageClose
		case '!':
			dest <- Cancel
		default:
			dest <- Other
		}
	}
	close(dest)
}

func feed(ch chan rune, s string) {
	for _, r := range s {
		ch <- r
	}
	close(ch)
}

func solve(input string) {
	inputStream := make(chan rune, 10000)
	tokenStream := make(chan Token, 10000)
	go feed(inputStream, input)
	go lex(inputStream, tokenStream)
	total, garbageTotal := evalTokenStream(tokenStream)
	if len(input) < 80 {
		fmt.Println(input)
	}
	fmt.Printf("Group total: %d, garbage total: %d\n", total, garbageTotal)
	fmt.Println("")
}

func solveFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	solve(strings.TrimSpace(string(contents)))
}

func main() {
	solve("{}")
	solve("{{{}}}")
	solve("{{},{}}")
	solve("{{{},{},{{}}}}")
	solve("{<a>,<a>,<a>,<a>}")
	solve("{{<ab>},{<ab>},{<ab>},{<ab>}}")
	solve("{{<!!>},{<!!>},{<!!>},{<!!>}}")
	solve("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	solve("<>")
	solve("<random characters>")
	solve("<<<<>")
	solve("<{!>}>")
	solve("<!!>")
	solve("<!!!>>")
	solve("<{o\"i!a,<{i<a>")
	solveFile("./input.txt")
}
