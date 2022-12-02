package day02

import (
	"aoc2022/src/utils"
	"fmt"
	"strings"
	"time"
)

type Sign int

const (
	Rock     Sign = 1
	Paper    Sign = 2
	Scissors Sign = 3
)

func hasPlayer2Won(player1 Sign, player2 Sign) bool {
	if player1 == Rock && player2 == Paper {
		return true
	}
	if player1 == Paper && player2 == Scissors {
		return true
	}
	if player1 == Scissors && player2 == Rock {
		return true
	}
	return false
}

func roundPoints(player1 Sign, player2 Sign) int {
	var points = 0
	if player1 == player2 {
		points += 3
	} else if hasPlayer2Won(player1, player2) {
		points += 6
	}

	points += int(player2)

	return points
}

func getSignFromPlayer(player string) Sign {
	if player == "A" || player == "X" {
		return Rock
	} else if player == "B" || player == "Y" {
		return Paper
	} else {
		return Scissors
	}
}

func getChangedSignForPlayer2(player1 Sign, player2 string) Sign {
	if player2 == "X" {
		// lose
		if player1 == Rock {
			return Scissors
		}
		if player1 == Paper {
			return Rock
		}
		if player1 == Scissors {
			return Paper
		}
	} else if player2 == "Z" {
		// win
		if player1 == Rock {
			return Paper
		}
		if player1 == Paper {
			return Scissors
		}
		if player1 == Scissors {
			return Rock
		}
	}

	// draw
	return player1
}

func Puzzle1() {
	start := time.Now()
	fmt.Println("Day 2 puzzle 1")
	lines := utils.ReadFileLines("src/day02/input.txt")

	var points = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		players := strings.Split(line, " ")
		player1 := getSignFromPlayer(players[0])
		player2 := getSignFromPlayer(players[1])
		points += roundPoints(player1, player2)
	}

	fmt.Println("Total points: ", points)
	fmt.Printf("Time elapsed %s\n", time.Since(start))
}

func Puzzle2() {
	start := time.Now()
	fmt.Println("Day 2 puzzle 2")
	lines := utils.ReadFileLines("src/day02/input.txt")

	var points = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		players := strings.Split(line, " ")
		player1 := getSignFromPlayer(players[0])
		player2 := getChangedSignForPlayer2(player1, players[1])
		points += roundPoints(player1, player2)
	}

	fmt.Println("Total points: ", points)

	fmt.Printf("Time elapsed %s\n\n", time.Since(start))
}

/*
A Rock
B Paper
C Scissors

X Rock
Y Paper
Z Scissors

Rock 1
Paper 2
Scissors 3

Won 6
Lose 0
Draw 3

-----

X Lose
Y Draw
Z Win
*/
