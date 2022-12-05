package day04

import (
	"aoc2022/src/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type cleaningRange struct {
	start int
	end   int
}

func getCleaningRange(section string) cleaningRange {
	splittedSection := strings.Split(section, "-")
	start, err := strconv.Atoi(splittedSection[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(splittedSection[1])
	if err != nil {
		panic(err)
	}

	cr := cleaningRange{
		start: start,
		end:   end,
	}
	return cr
}

func isContaining(section1 string, section2 string) bool {
	range1 := getCleaningRange(section1)
	range2 := getCleaningRange(section2)

	if range1.start <= range2.start && range1.end >= range2.end {
		return true
	}

	if range2.start <= range1.start && range2.end >= range1.end {
		return true
	}
	return false
}

func isOverlapping(section1 string, section2 string) bool {
	range1 := getCleaningRange(section1)
	range2 := getCleaningRange(section2)

	if range1.start <= range2.end && range2.start <= range1.end {
		return true
	}

	return false
}

func Puzzle1() {
	start := time.Now()
	fmt.Println("Day 4 puzzle 1")
	lines := utils.ReadFileLines("src/day04/input.txt")

	var containing = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		pair := strings.Split(line, ",")

		if isContaining(pair[0], pair[1]) {
			containing++
		}

	}

	fmt.Println("Containing: ", containing)
	fmt.Printf("Time elapsed %s\n", time.Since(start))
}

func Puzzle2() {
	start := time.Now()
	fmt.Println("Day 4 puzzle 2")
	lines := utils.ReadFileLines("src/day04/input.txt")

	var overlapping = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		pair := strings.Split(line, ",")

		if isOverlapping(pair[0], pair[1]) {
			overlapping++
		}

	}

	fmt.Println("Overlapping: ", overlapping)

	fmt.Printf("Time elapsed %s\n\n", time.Since(start))
}
