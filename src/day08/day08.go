package day08

import (
	"aoc2022/src/utils"
	"fmt"
	"strconv"
	"time"
)

type grid struct {
	rows [][]int
}

func isVisible(g *grid, x, y, width, height int) bool {
	treeHeigth := (*g).rows[y][x]
	var left = true
	var up = true
	var right = true
	var down = true

	// look left
	for i := x - 1; i >= 0; i-- {
		if (*g).rows[y][i] >= treeHeigth {
			left = false
			break
		}
	}
	// look up
	for i := y - 1; i >= 0; i-- {
		if (*g).rows[i][x] >= treeHeigth {
			up = false
			break
		}
	}
	// look right
	for i := x + 1; i < width; i++ {
		if (*g).rows[y][i] >= treeHeigth {
			right = false
			break
		}
	}
	// look down
	for i := y + 1; i < height; i++ {
		if (*g).rows[i][x] >= treeHeigth {
			down = false
			break
		}
	}

	return left || up || right || down
}

func getScore(g *grid, x, y, width, height int) int {
	treeHeigth := (*g).rows[y][x]
	var left = 0
	var up = 0
	var right = 0
	var down = 0

	// look left
	for i := x - 1; i >= 0; i-- {
		left++
		if (*g).rows[y][i] >= treeHeigth {
			break
		}
	}
	// look up
	for i := y - 1; i >= 0; i-- {
		up++
		if (*g).rows[i][x] >= treeHeigth {
			break
		}
	}
	// look right
	for i := x + 1; i < width; i++ {
		right++
		if (*g).rows[y][i] >= treeHeigth {
			break
		}
	}
	// look down
	for i := y + 1; i < height; i++ {
		down++
		if (*g).rows[i][x] >= treeHeigth {
			break
		}
	}

	return left * up * right * down
}

func getVisibles(g *grid, width int, height int) int {
	var visibles = 0
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if isVisible(g, x, y, width, height) {
				visibles++
			}
		}
	}
	return visibles
}

func getHighestTreeScore(g *grid, width int, height int) int {
	var score = 0
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			treeScore := getScore(g, x, y, width, height)
			if treeScore > score {
				score = treeScore
			}
		}
	}
	return score
}

func Puzzle1() {
	start := time.Now()
	fmt.Println("Day 8 puzzle 1")

	var g = grid{
		rows: nil,
	}

	lines := utils.ReadFileLines("src/day08/input.txt")
	var width = 0
	var height = 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		height++
		width = len(line)
		g.rows = append(g.rows, make([]int, width))
		for j := 0; j < len(line); j++ {
			g.rows[i][j], _ = strconv.Atoi(string(line[j]))
		}
	}

	var visible = width*2 + (height-2)*2 + getVisibles(&g, width, height)

	fmt.Println("Visible trees: ", visible)
	fmt.Printf("Time elapsed %s\n", time.Since(start))
}

func Puzzle2() {
	start := time.Now()
	fmt.Println("Day 8 puzzle 2")
	var g = grid{
		rows: nil,
	}

	lines := utils.ReadFileLines("src/day08/input.txt")
	var width = 0
	var height = 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		height++
		width = len(line)
		g.rows = append(g.rows, make([]int, width))
		for j := 0; j < len(line); j++ {
			g.rows[i][j], _ = strconv.Atoi(string(line[j]))
		}
	}
	highestScore := getHighestTreeScore(&g, width, height)

	fmt.Println("Highest score: ", highestScore)
	fmt.Printf("Time elapsed %s\n\n", time.Since(start))
}
