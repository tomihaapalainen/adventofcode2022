package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func adv9a() {
	file, _ := os.Open("input9.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	tail := Position{x: 0, y: 0}

	hp := Position{x: 0, y: 0}

	visited := make(map[Position]bool)
	visited[tail] = true

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		split := strings.Split(row, " ")
		dir, countString := split[0], split[1]
		count, _ := strconv.ParseInt(countString, 10, 32)

		var c int64
		for c = 0; c < count; c++ {
			switch {
			case dir == "U":
				hp.y++
			case dir == "R":
				hp.x++
			case dir == "D":
				hp.y--
			case dir == "L":
				hp.x--
			}

			head := hp

			if math.Abs(float64(head.x-tail.x)) > 1.0 || math.Abs(float64(head.y-tail.y)) > 1.0 {
				dx := rectify(head.x - tail.x)
				dy := rectify(head.y - tail.y)
				tail.x += dx
				tail.y += dy
			}
			head = tail
			visited[tail] = true
		}
	}
	fmt.Println(len(visited))
}

func rectify(x int) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}

func adv9b() {
	file, _ := os.Open("input9.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	tails := []Position{
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
	}

	hp := Position{x: 0, y: 0}

	visited := make(map[Position]bool)
	visited[tails[0]] = true

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		split := strings.Split(row, " ")
		dir, countString := split[0], split[1]
		count, _ := strconv.ParseInt(countString, 10, 32)

		var c int64
		for c = 0; c < count; c++ {
			switch {
			case dir == "U":
				hp.y++
			case dir == "R":
				hp.x++
			case dir == "D":
				hp.y--
			case dir == "L":
				hp.x--
			}

			head := hp
			for i := 0; i < len(tails); i++ {
				if math.Abs(float64(head.x-tails[i].x)) > 1.0 || math.Abs(float64(head.y-tails[i].y)) > 1.0 {
					dx := rectify(head.x - tails[i].x)
					dy := rectify(head.y - tails[i].y)
					tails[i].x += dx
					tails[i].y += dy
				}
				head = tails[i]
			}
			visited[tails[len(tails)-1]] = true
		}
	}
	fmt.Println(len(visited))
}
