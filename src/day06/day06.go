package day06

import (
	"aoc2022/src/utils"
	"fmt"
	"time"

	"golang.org/x/exp/slices"
)

func isUnique(sequence string) bool {
	var check []rune

	for _, value := range sequence {
		if slices.Contains(check, value) {
			return false
		}
		check = append(check, value)
	}

	return true
}

func Puzzle1() {
	start := time.Now()
	fmt.Println("Day 6 puzzle 1")
	stream := utils.ReadFile("src/day06/input.txt")

	length := len(stream)
	const messageLength = 4
	found := -1
	for i := 0; i < length; i++ {
		// still 4 characters to read
		if i+(messageLength-1) <= length {
			toCheck := stream[i : i+messageLength]
			if isUnique(toCheck) {
				found = i + messageLength
				break
			}
		}
	}

	fmt.Println("Found after", found)
	fmt.Printf("Time elapsed %s\n", time.Since(start))
}

func Puzzle2() {
	start := time.Now()
	fmt.Println("Day 6 puzzle 2")
	stream := utils.ReadFile("src/day06/input.txt")

	length := len(stream)
	const messageLength = 14
	found := -1
	for i := 0; i < length; i++ {
		// still 4 characters to read
		if i+(messageLength-1) <= length {
			toCheck := stream[i : i+messageLength]
			if isUnique(toCheck) {
				found = i + messageLength
				break
			}
		}
	}

	fmt.Println("Found after", found)

	fmt.Printf("Time elapsed %s\n\n", time.Since(start))
}
