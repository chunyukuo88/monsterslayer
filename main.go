package main

import (
	"github.com/chunyukuo88/monsterslayer/interactions"
)

var currentRound = 0

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}
	endGame()
}

func startGame() {
	interactions.PrintGreeting()
}
func endGame() {}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0
	interactions.ShowAvailableActions(isSpecialRound)
	return ""
}
