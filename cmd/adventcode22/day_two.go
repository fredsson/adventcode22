package main

import (
	"fmt"
	"strings"
)

const POINTS_FOR_ROCK = 1
const POINTS_FOR_PAPER = 2
const POINTS_FOR_SCISSOR = 3

const SYMBOL_FOR_ROCK = "X"
const SYMBOL_FOR_PAPER = "Y"
const SYMBOL_FOR_SCISSOR = "Z"
const SYMBOL_FOR_LOSE = "X"
const SYMBOL_FOR_DRAW = "Y"
const SYMBOL_FOR_WIN = "Z"
const SYMBOL_FOR_OPPONENT_ROCK = "A"
const SYMBOL_FOR_OPPONENT_PAPER = "B"
const SYMBOL_FOR_OPPONENT_SCISSOR = "C"

func decideWinnerScore(opponentAction string, playerAction string) int {
	if opponentAction == SYMBOL_FOR_OPPONENT_ROCK && playerAction == SYMBOL_FOR_PAPER {
		return 6
	}
	if opponentAction == SYMBOL_FOR_OPPONENT_PAPER && playerAction == SYMBOL_FOR_SCISSOR {
		return 6
	}
	if opponentAction == SYMBOL_FOR_OPPONENT_SCISSOR && playerAction == SYMBOL_FOR_ROCK {
		return 6
	}
	if opponentAction == SYMBOL_FOR_OPPONENT_ROCK && playerAction == SYMBOL_FOR_ROCK {
		return 3
	}
	if opponentAction == SYMBOL_FOR_OPPONENT_PAPER && playerAction == SYMBOL_FOR_PAPER {
		return 3
	}
	if opponentAction == SYMBOL_FOR_OPPONENT_SCISSOR && playerAction == SYMBOL_FOR_SCISSOR {
		return 3
	}

	return 0
}

func decideActionScore(playerAction string) int {
	if playerAction == SYMBOL_FOR_PAPER {
		return POINTS_FOR_PAPER
	}
	if playerAction == SYMBOL_FOR_ROCK {
		return POINTS_FOR_ROCK
	}
	if playerAction == SYMBOL_FOR_SCISSOR {
		return POINTS_FOR_SCISSOR
	}
	return 0
}

func findActionFromOutcome(opponentAction string, expectedOutcome string) string {
	if expectedOutcome == SYMBOL_FOR_LOSE {
		if opponentAction == SYMBOL_FOR_OPPONENT_ROCK {
			return SYMBOL_FOR_SCISSOR
		}
		if opponentAction == SYMBOL_FOR_OPPONENT_PAPER {
			return SYMBOL_FOR_ROCK
		}
		if opponentAction == SYMBOL_FOR_OPPONENT_SCISSOR {
			return SYMBOL_FOR_PAPER
		}
	}
	if expectedOutcome == SYMBOL_FOR_DRAW {
		if opponentAction == SYMBOL_FOR_OPPONENT_ROCK {
			return SYMBOL_FOR_ROCK
		}
		if opponentAction == SYMBOL_FOR_OPPONENT_PAPER {
			return SYMBOL_FOR_PAPER
		}
		if opponentAction == SYMBOL_FOR_OPPONENT_SCISSOR {
			return SYMBOL_FOR_SCISSOR
		}
	}
	if expectedOutcome == SYMBOL_FOR_WIN {
		if opponentAction == SYMBOL_FOR_OPPONENT_ROCK {
			return SYMBOL_FOR_PAPER
		}
		if opponentAction == SYMBOL_FOR_OPPONENT_PAPER {
			return SYMBOL_FOR_SCISSOR
		}
		if opponentAction == SYMBOL_FOR_OPPONENT_SCISSOR {
			return SYMBOL_FOR_ROCK
		}
	}

	return SYMBOL_FOR_ROCK
}

func DayTwo() {
	openFile := readFileByLines("inputs/d2.txt")

	totalScoreA := 0
	totalScore := 0
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()
		actions := strings.Split(input, " ")
		expectedAction := findActionFromOutcome(actions[0], actions[1])

		scoreForWinA := decideWinnerScore(actions[0], actions[1])
		scoreForActionA := decideActionScore(actions[1])
		totalScoreA += scoreForWinA + scoreForActionA

		scoreForWin := decideWinnerScore(actions[0], expectedAction)
		scoreForAction := decideActionScore(expectedAction)
		totalScore += scoreForWin + scoreForAction
	}
	fmt.Println(totalScoreA)
	fmt.Println(totalScore)

	openFile.File.Close()
}
