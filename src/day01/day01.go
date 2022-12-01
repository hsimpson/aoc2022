package day01

import (
	"aoc2022/src/utils"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func Puzzle1() {
	start := time.Now()
	fmt.Println("Day 1 puzzle 1")
	lines := utils.ReadFileLines("src/day01/input.txt")

	var elves []int
	var sum = 0
	var max = 0
	var index = -1

	for _, value := range lines {
		if index == -1 {
			elves = append(elves, 0)
			index++
		}

		if value == "" {
			elves = append(elves, sum)
			index++
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			calories, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			sum += calories
			elves[index] = sum
		}
	}

	fmt.Println("Max calories: ", max)

	fmt.Printf("Time elapsed %s\n", time.Since(start))
}

func Max(sum, max int) {
	panic("unimplemented")
}

func Puzzle2() {
	start := time.Now()
	fmt.Println("Day 1 puzzle 2")

	lines := utils.ReadFileLines("src/day01/input.txt")

	var elves []int
	var sum = 0

	var elveIndex = -1

	for index, value := range lines {
		if elveIndex == -1 {
			elves = append(elves, 0)
			elveIndex++
		}

		if value == "" {
			if index < len(lines)-1 {
				elves = append(elves, 0)
				elveIndex++
				sum = 0
			}
		} else {
			calories, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			sum += calories
			elves[elveIndex] = sum
		}
	}
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	var max = elves[0] + elves[1] + elves[2]

	fmt.Println("Max calories top 3: ", max)

	fmt.Printf("Time elapsed %s\n", time.Since(start))
}
