package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
	Cycle int
	X     int
}

func (state *State) getSignalStrength() int {
	return state.Cycle * state.X
}

func adv10a() {
	file, _ := os.Open("input10.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	state := State{Cycle: 0, X: 1}

	cycles := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}

	result := 0

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())

		switch {
		case strings.Contains(row, "noop"):
			state.Cycle++
			if cycles[state.Cycle] {
				result += state.getSignalStrength()
			}
		case strings.Contains(row, "addx"):
			split := strings.Split(row, " ")
			count, _ := strconv.ParseInt(split[1], 10, 32)

			state.Cycle++
			if cycles[state.Cycle] {
				result += state.getSignalStrength()
			}
			state.Cycle++
			if cycles[state.Cycle] {
				result += state.getSignalStrength()
			}
			state.X += int(count)
		}
	}

	fmt.Println(result)
}

func cyclePrint(state *State) {
	if state.Cycle == 40 {
		fmt.Println()
		state.Cycle = 0
		return
	}
	if state.X-1 <= state.Cycle && state.Cycle <= state.X+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}

func adv10b() {
	file, _ := os.Open("input10.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	state := State{Cycle: 0, X: 1}

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())

		switch {
		case strings.Contains(row, "noop"):
			state.Cycle++
			cyclePrint(&state)
		case strings.Contains(row, "addx"):
			split := strings.Split(row, " ")
			count, _ := strconv.ParseInt(split[1], 10, 32)

			state.Cycle++
			cyclePrint(&state)

			state.Cycle++
			state.X += int(count)
			cyclePrint(&state)
		}
	}
}
