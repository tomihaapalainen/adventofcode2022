package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tree struct {
	x int
	y int
}

func readGrid() [][]int {
	file, _ := os.Open("input8.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	grid := [][]int{}
	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		digits := []int{}
		for _, b := range row {
			digits = append(digits, int(b)-48)
		}
		grid = append(grid, digits)
	}

	return grid
}

func adv8a() {
	grid := readGrid()

	visible := make(map[Tree]bool)

	// left to right
	for i := 0; i < len(grid); i++ {
		max := -1
		for j := 0; j < len(grid[0]); j++ {
			height := grid[i][j]

			if height > max {
				max = height
				visible[Tree{x: i, y: j}] = true
			}
		}
	}

	// right to left
	for i := 0; i < len(grid); i++ {
		max := -1
		for j := len(grid[0]) - 1; j >= 0; j-- {
			height := grid[i][j]

			if height > max {
				max = height
				visible[Tree{x: i, y: j}] = true
			}
		}
	}

	// top to bottom
	for i := 0; i < len(grid); i++ {
		max := -1
		for j := 0; j < len(grid[0]); j++ {
			height := grid[j][i]

			if height > max {
				max = height
				visible[Tree{x: j, y: i}] = true
			}
		}
	}

	// bottom to top
	for i := len(grid) - 1; i >= 0; i-- {
		max := -1
		for j := len(grid[0]) - 1; j >= 0; j-- {
			height := grid[j][i]

			if height > max {
				max = height
				visible[Tree{x: j, y: i}] = true
			}
		}
	}

	fmt.Println(len(visible))
}

func calculateScenicScore(i int, j int, grid [][]int) int {
	center := grid[i][j]

	top := 0
	right := 0
	bottom := 0
	left := 0

	for r := i + 1; r < len(grid); r++ {
		tree := grid[r][j]

		bottom++

		if tree >= center {
			break
		}
	}

	for r := i - 1; r >= 0; r-- {
		tree := grid[r][j]

		top++

		if tree >= center {
			break
		}
	}

	for c := j + 1; c < len(grid[0]); c++ {
		tree := grid[i][c]

		right++

		if tree >= center {
			break
		}
	}

	for c := j - 1; c >= 0; c-- {
		tree := grid[i][c]

		left++

		if tree >= center {
			break
		}
	}

	return top * right * bottom * left
}

func adv8b() {
	grid := readGrid()

	maxScore := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			score := calculateScenicScore(i, j, grid)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println(maxScore)
}
