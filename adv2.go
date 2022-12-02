package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func adv2a() {
	win := make(map[string]string)
	win["X"] = "C"
	win["Y"] = "A"
	win["Z"] = "B"

	equal := make(map[string]string)
	equal["X"] = "A"
	equal["Y"] = "B"
	equal["Z"] = "C"

	choiceScore := make(map[string]int)
	choiceScore["X"] = 1
	choiceScore["Y"] = 2
	choiceScore["Z"] = 3

	file, _ := os.Open("input2.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	totalScore := 0
	for fileScanner.Scan() {
		choices := strings.Split(strings.TrimSpace(fileScanner.Text()), " ")
		opponent, me := choices[0], choices[1]

		score := 0
		if win[me] == opponent {
			score += 6
		} else if equal[me] == opponent {
			score += 3
		}

		score += choiceScore[me]

		totalScore += score
	}

	fmt.Println(totalScore)
}

func adv2b() {
	losingPicks := make(map[string]string)
	losingPicks["A"] = "C"
	losingPicks["B"] = "A"
	losingPicks["C"] = "B"

	winningPicks := make(map[string]string)
	winningPicks["C"] = "A"
	winningPicks["A"] = "B"
	winningPicks["B"] = "C"

	choiceScore := make(map[string]int)
	choiceScore["A"] = 1
	choiceScore["B"] = 2
	choiceScore["C"] = 3

	file, _ := os.Open("input2.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	totalScore := 0
	for fileScanner.Scan() {
		choices := strings.Split(strings.TrimSpace(fileScanner.Text()), " ")
		opponent, me := choices[0], choices[1]

		score := 0
		if me == "X" {
			me = losingPicks[opponent]
		} else if me == "Y" {
			me = opponent
			score += 3
		} else {
			me = winningPicks[opponent]
			score += 6
		}

		score += choiceScore[me]

		totalScore += score
	}

	fmt.Println(totalScore)
}
