package day07

import (
	"aoc2022/src/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type nodeType int

const (
	Dir  nodeType = 0
	File nodeType = 1
)

type node struct {
	parent   *node
	children map[string]*node
	name     string
	size     int
	t        nodeType
}

var tree = node{
	parent:   nil,
	children: make(map[string]*node),
	name:     "/",
	size:     0,
	t:        Dir,
}

var cwd = &tree

func cd(directory string) {
	if directory == "/" {
		cwd = &tree
	} else if directory == ".." {
		cwd = cwd.parent
	} else {
		cwd = cwd.children[directory]
	}
}

func dirSizeAtMost10K(n *node, accum *int) int {
	var size = 0
	for _, child := range n.children {
		if child.t == Dir {
			size += dirSizeAtMost10K(child, accum)
		} else {
			size += child.size
		}
	}

	if size <= 100_000 {
		*accum += size
	}
	return size
}

func dirSize(n *node) int {
	var size = 0
	for _, child := range n.children {
		if child.t == Dir {
			size += dirSize(child)
		} else {
			size += child.size
		}
	}
	return size
}

func smallestToDelete(n *node, smallest *int, needToDelete int) int {
	var size = 0
	for _, child := range n.children {
		if child.t == Dir {
			size += smallestToDelete(child, smallest, needToDelete)
		} else {
			size += child.size
		}
	}
	if size >= needToDelete && size < *smallest {
		*smallest = size
	}
	return size
}

const filePath = "src/day07/input.txt"

func Puzzle1() {
	start := time.Now()
	fmt.Println("Day 7 puzzle 1")
	lines := utils.ReadFileLines(filePath)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}

		if line[0] == '$' {
			// command
			command := string(line[2:])
			if command[0:2] == "cd" {
				cd(command[3:])
			} else if command[0:2] == "ls" {
				// ls
				continue
			}
		} else {
			// dir or file
			segment := strings.Split(line, " ")
			if segment[0] == "dir" {
				// dir
				dirNode := node{
					parent:   cwd,
					children: make(map[string]*node),
					name:     segment[1],
					size:     0,
					t:        Dir,
				}
				cwd.children[segment[1]] = &dirNode
			} else {
				// file
				size, _ := strconv.Atoi(segment[0])
				fileNode := node{
					parent:   cwd,
					children: nil,
					name:     segment[1],
					size:     size,
					t:        File,
				}
				cwd.children[segment[1]] = &fileNode
			}
		}
	}

	var sum = 0
	dirSizeAtMost10K(&tree, &sum)

	fmt.Println("Sum of dirs at most 10000 ", sum)
	fmt.Printf("Time elapsed %s\n", time.Since(start))
}

func Puzzle2() {
	start := time.Now()
	fmt.Println("Day 7 puzzle 2")

	const diskSize = 70_000_000
	const neededSize = 30_000_000
	totalSize := dirSize(&tree)
	remaining := diskSize - totalSize
	needToDelete := neededSize - remaining
	var smallestDir = math.MaxInt
	smallestToDelete(&tree, &smallestDir, needToDelete)

	fmt.Println("delete", smallestDir)

	fmt.Printf("Time elapsed %s\n\n", time.Since(start))
}
