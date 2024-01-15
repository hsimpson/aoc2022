package day09

import (
	"aoc2022/src/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type pos struct {
	x int
	y int
}

func isAdjacent(head, tail *pos) bool {
	if head.x == tail.x && head.y == tail.y {
		return true
	}

	return false
}

func updateTail(head, tail *pos) {
	if !isAdjacent(head, tail) {

	}
}

func moveStep(direction string, head, tail *pos) {
	switch direction {
	case "R":
		head.x++
	case "L":
		head.x--
	case "U":
		head.y++
	case "D":
		head.y--
	}
}

func move(direction string, distance int, head, tail *pos) {
	for i := 0; i < distance; i++ {
		moveStep(direction, head, tail)
	}
}

func Puzzle1() {
	start := time.Now()
	fmt.Println("Day 9 puzzle 1")

	var head = pos{x: 0, y: 0}
	var tail = pos{x: 0, y: 0}

	lines := utils.ReadFileLines("src/day09/test.txt")
	for _, line := range lines {
		if line == "" {
			continue
		}

		segments := strings.Split(line, " ")
		distance, _ := strconv.Atoi(segments[1])
		move(segments[0], distance, &head, &tail)
	}

	fmt.Printf("Time elapsed %s\n", time.Since(start))
}

func Puzzle2() {
	start := time.Now()
	fmt.Println("Day 9 puzzle 2")

	fmt.Printf("Time elapsed %s\n\n", time.Since(start))
}
