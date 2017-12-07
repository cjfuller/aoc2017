package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ProgramNode struct {
	Name      string
	Weight    int
	SumWeight int
	Children  []*ProgramNode
	Parent    *ProgramNode
}

type NodeInfo struct {
	Name       string
	Weight     int
	ChildNames []string
}

var nodeRe = regexp.MustCompile(`(\w+) \((\d+)\)(?: -> )?(.*)`)

func parseInputFile(filename string) []NodeInfo {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	result := []NodeInfo{}
	for _, line := range strings.Split(strings.TrimSpace(string(contents)), "\n") {
		matchResult := nodeRe.FindStringSubmatch(line)
		name := matchResult[1]
		weight, err := strconv.Atoi(matchResult[2])
		if err != nil {
			panic(err)
		}
		childrenString := ""
		if len(matchResult) > 3 {
			childrenString = matchResult[3]
		}
		children := []string{}
		if childrenString != "" {
			children = strings.Split(strings.TrimSpace(childrenString), ", ")
		}
		result = append(result, NodeInfo{
			Name:       name,
			Weight:     weight,
			ChildNames: children,
		})
	}
	return result
}

func allChildrenInIndex(info NodeInfo, idx map[string]*ProgramNode) bool {
	for _, child := range info.ChildNames {
		_, ok := idx[child]
		if !ok {
			return false
		}
	}
	return true
}

func constructTree(info []NodeInfo) (*ProgramNode, int) {
	idx := make(map[string]*ProgramNode)
	var root *ProgramNode
	requiredWeight := -1
	for len(info) > 0 {
		curr := info[0]
		info = info[1:]
		if allChildrenInIndex(curr, idx) {
			childNodes := []*ProgramNode{}
			sumWeight := 0
			weights := make(map[int]int)
			for _, childName := range curr.ChildNames {
				childNode := idx[childName]
				childNodes = append(childNodes, childNode)
				sumWeight += childNode.SumWeight
				weights[childNode.SumWeight]++
			}
			// Because we're constructing the tree from leaves to root, we
			// know that the first time we've encountered an imbalance,
			// that's the source of the problem. Thus, we don't need to
			// bother looking if we've already calculated a requiredWeight
			// somewhere else.
			if len(weights) > 1 && requiredWeight == -1 {
				var badNode *ProgramNode
				goodWeight := 0
				for _, node := range childNodes {
					if weights[node.SumWeight] == 1 {
						badNode = node
					} else {
						goodWeight = node.SumWeight
					}
				}
				requiredWeight = goodWeight - (badNode.SumWeight - badNode.Weight)
			}
			currNode := ProgramNode{
				Name:      curr.Name,
				Weight:    curr.Weight,
				SumWeight: curr.Weight + sumWeight,
				Children:  childNodes,
				Parent:    nil,
			}
			idx[currNode.Name] = &currNode
			for _, childPtr := range currNode.Children {
				childPtr.Parent = &currNode
			}
			root = &currNode
		} else {
			info = append(info, curr)
		}
	}
	return root, requiredWeight
}

func printSolution(infos []NodeInfo) {
	root, reqdWeight := constructTree(infos)
	fmt.Println(root.Name)
	fmt.Printf("%d\n", reqdWeight)
}

func main() {
	printSolution(parseInputFile("./p1test.txt"))
	printSolution(parseInputFile("./input.txt"))
}
