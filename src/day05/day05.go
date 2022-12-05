package day05

import (
	"aoc2022/src/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func push(stack *[]string, value string) {
	*stack = append(*stack, value)
}

func pop(stack *[]string) string {
	if len(*stack) > 0 {
		val := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		return val
	}
	return ""
}

func peek(stack *[]string) string {
	if len(*stack) > 0 {
		val := (*stack)[len(*stack)-1]
		return val
	}
	return ""
}

func fillStack(stacks *[][]string, line string, numOfStacks int) {
	lineLength := ((numOfStacks - 1) * 4) + 3

	var stackIdx = 0
	for i := 1; i < lineLength; i += 4 {
		c := string(line[i])
		stack := &(*stacks)[stackIdx]
		if c != " " {
			push(stack, c)
		}

		stackIdx++
	}
}

func playRound(stacks *[][]string, line string, moveMultiple bool) {
	splitted := strings.Split(line, " ")
	amount, err1 := strconv.Atoi(splitted[1])
	from, err2 := strconv.Atoi(splitted[3])
	to, err3 := strconv.Atoi(splitted[5])

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}
	if err3 != nil {
		panic(err3)
	}

	fromStack := &(*stacks)[from-1]
	toStack := &(*stacks)[to-1]

	var popped = ""
	for i := 0; i < amount; i++ {
		popped += pop(fromStack)
	}

	if !moveMultiple {
		for i := 0; i < len(popped); i++ {
			push(toStack, string(popped[i]))
		}
	} else {
		for i := len(popped) - 1; i >= 0; i-- {
			push(toStack, string(popped[i]))
		}
	}

}

func Puzzle1() {
	start := time.Now()
	fmt.Println("Day 5 puzzle 1")
	lines := utils.ReadFileLines("src/day05/input.txt")

	var emptyIndex = 0
	for i := 0; i < len(lines); i++ {
		// search for the empty line
		if lines[i] == "" {
			emptyIndex = i
			break
		}
	}

	stackInfoLine := strings.TrimSpace(lines[emptyIndex-1])
	numbers := strings.Split(stackInfoLine, " ")
	numOfStacks, err := strconv.Atoi(numbers[len(numbers)-1])
	if err != nil {
		panic(err)
	}

	// prepare the stacks
	stacks := make([][]string, numOfStacks)
	for i := emptyIndex - 2; i >= 0; i-- {
		fillStack(&stacks, lines[i], numOfStacks)
	}

	// play the rounds
	for i := emptyIndex + 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		playRound(&stacks, lines[i], false)
	}

	var finalString = ""
	for _, stack := range stacks {
		finalString += peek(&stack)
	}

	fmt.Println("Final string: ", finalString)
	fmt.Printf("Time elapsed %s\n", time.Since(start))
}

func Puzzle2() {
	start := time.Now()
	fmt.Println("Day 5 puzzle 2")
	lines := utils.ReadFileLines("src/day05/input.txt")

	var emptyIndex = 0
	for i := 0; i < len(lines); i++ {
		// search for the empty line
		if lines[i] == "" {
			emptyIndex = i
			break
		}
	}

	stackInfoLine := strings.TrimSpace(lines[emptyIndex-1])
	numbers := strings.Split(stackInfoLine, " ")
	numOfStacks, err := strconv.Atoi(numbers[len(numbers)-1])
	if err != nil {
		panic(err)
	}

	// prepare the stacks
	stacks := make([][]string, numOfStacks)
	for i := emptyIndex - 2; i >= 0; i-- {
		fillStack(&stacks, lines[i], numOfStacks)
	}

	// play the rounds
	for i := emptyIndex + 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		playRound(&stacks, lines[i], true)
	}

	var finalString = ""
	for _, stack := range stacks {
		finalString += peek(&stack)
	}

	fmt.Println("Final string: ", finalString)

	fmt.Printf("Time elapsed %s\n\n", time.Since(start))
}
