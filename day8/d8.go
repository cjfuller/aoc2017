package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Op int
type Cond int

const (
	Inc Op = iota
	Dec
)

const (
	Gt Cond = iota
	Lt
	Geq
	Leq
	Eq
	Neq
)

type Instruction struct {
	TargetReg  string
	TargetOp   Op
	TargetAmt  int
	CondSource string
	Condition  Cond
	CondAmt    int
}

func parseOp(op string) Op {
	if op == "inc" {
		return Inc
	} else if op == "dec" {
		return Dec
	}
	panic("Invalid op " + op)
}

func parseCond(cond string) Cond {
	switch cond {
	case ">":
		return Gt
	case "<":
		return Lt
	case ">=":
		return Geq
	case "<=":
		return Leq
	case "==":
		return Eq
	case "!=":
		return Neq
	default:
		panic("Invalid condition " + cond)
	}
}

func parseInstruction(instr string) Instruction {
	parts := strings.Split(instr, " ")
	targetAmt, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	condAmt, err := strconv.Atoi(parts[6])
	if err != nil {
		panic(err)
	}
	return Instruction{
		TargetReg:  parts[0],
		TargetOp:   parseOp(parts[1]),
		TargetAmt:  targetAmt,
		CondSource: parts[4],
		Condition:  parseCond(parts[5]),
		CondAmt:    condAmt,
	}
}

func instructionsFromFile(filename string) []Instruction {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(contents)), "\n")
	result := []Instruction{}
	for _, line := range lines {
		result = append(result, parseInstruction(line))
	}
	return result
}

func evalCondition(registers map[string]int, instr Instruction) bool {
	sourceVal := registers[instr.CondSource]
	switch instr.Condition {
	case Gt:
		return sourceVal > instr.CondAmt
	case Lt:
		return sourceVal < instr.CondAmt
	case Geq:
		return sourceVal >= instr.CondAmt
	case Leq:
		return sourceVal <= instr.CondAmt
	case Eq:
		return sourceVal == instr.CondAmt
	case Neq:
		return sourceVal != instr.CondAmt
	default:
		panic("Should not reach here; no valid operation.")
	}
}

func evalOperation(registers map[string]int, instr Instruction) int {
	switch instr.TargetOp {
	case Inc:
		registers[instr.TargetReg] += instr.TargetAmt
	case Dec:
		registers[instr.TargetReg] -= instr.TargetAmt
	}
	return registers[instr.TargetReg]
}

func executeInstructionsTrackingMax(instrs []Instruction) (map[string]int, int) {
	registers := make(map[string]int)
	maxRegisterEver := 0
	for _, instr := range instrs {
		if evalCondition(registers, instr) {
			resultingRegValue := evalOperation(registers, instr)
			if resultingRegValue > maxRegisterEver {
				maxRegisterEver = resultingRegValue
			}
		}
	}
	return registers, maxRegisterEver
}

func maxRegister(registers map[string]int) int {
	result := math.MinInt64
	for _, val := range registers {
		if val > result {
			result = val
		}
	}
	return result
}

func solve(filename string) {
	registers, maxRegEver := executeInstructionsTrackingMax(instructionsFromFile(filename))
	fmt.Printf("Current max: %d, max ever: %d\n", maxRegister(registers), maxRegEver)
}

func main() {
	solve("./testp1.txt")
	solve("./input.txt")
}
