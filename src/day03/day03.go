package day03

import (
	"aoc2022/src/utils"
	"fmt"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

func getPriority(character rune) int {
	value := int(character)
	if value >= int('a') {
		value -= 96
	} else if value >= int('A') {
		value -= 38
	}
	return value
}

func getDouble(firstHalf string, secondHalf string) rune {
	for _, c := range secondHalf {
		if strings.ContainsRune(firstHalf, c) {
			return c
		}
	}
	return '0'
}

func getTripple(first string, second string, third string) rune {
	var secondInFirst []rune
	for _, c := range second {
		if strings.ContainsRune(first, c) && !slices.Contains(secondInFirst, c) {
			secondInFirst = append(secondInFirst, c)
		}
	}

	for _, c := range secondInFirst {
		if strings.ContainsRune(third, c) {
			return c
		}
	}
	return '0'
}

func Puzzle1() {
	start := time.Now()
	fmt.Println("Day 3 puzzle 1")
	lines := utils.ReadFileLines("src/day03/input.txt")

	var sum = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		length := len(line)
		firstHalf := line[0:(length / 2)]
		secondHalf := line[(length / 2):length]

		doubleCharacter := getDouble(firstHalf, secondHalf)
		sum += getPriority(doubleCharacter)
	}

	fmt.Println("Sum of priorities: ", sum)
	fmt.Printf("Time elapsed %s\n", time.Since(start))
}

func Puzzle2() {
	start := time.Now()
	fmt.Println("Day 3 puzzle 2")
	lines := utils.ReadFileLines("src/day03/input.txt")

	var sum = 0
	for i := 0; i < len(lines); i += 3 {

		if lines[i] == "" {
			continue
		}
		a := lines[i]
		b := lines[i+1]
		c := lines[i+2]
		tripple := getTripple(a, b, c)
		sum += getPriority(tripple)
	}

	fmt.Println("Sum of priorities: ", sum)
	fmt.Printf("Time elapsed %s\n\n", time.Since(start))
}
