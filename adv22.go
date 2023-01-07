package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PosNode struct {
	pos   Pos
	index int
	b     byte
	prev  *PosNode
	next  *PosNode
}

type Pos struct {
	r int
	c int
}

func (p *Pos) AddPosition(other Pos) Pos {
	return Pos{r: p.r + other.r, c: p.c + other.c}
}

func (p *Pos) IsOutOfGrid(grid [][]byte) bool {
	return p.r < 0 || p.r >= len(grid) || p.c < 0 || p.c >= len(grid[p.r])
}

func addPositions(a, b Pos) Pos {
	return Pos{r: a.r + b.r, c: a.c + b.c}
}

func printMap(show [][]byte) {
	for r := range show {
		for c := range show[r] {
			fmt.Print(string(show[r][c]))
		}
		fmt.Println()
	}
}

func adv22a() {
	file, _ := os.Open("input22.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	grid := [][]byte{}
	show := [][]byte{}
	var moves string

	r := 0
	empty := false
	for fileScanner.Scan() {
		row := fileScanner.Text()

		if len(row) > 0 && !empty {
			bytes := []byte(row)
			showBytes := []byte(row)
			grid = append(grid, bytes)
			show = append(show, showBytes)
		} else if empty {
			moves = row
		}
		if len(row) == 0 {
			empty = true
		}
		r++
	}

	max := 0
	for r := range grid {
		if len(grid[r]) > max {
			max = len(grid[r])
		}
	}

	for r := range grid {
		for len(grid[r]) <= max {
			grid[r] = append(grid[r], ' ')
			show[r] = append(show[r], ' ')
		}
	}

	printMap(show)

	head := &PosNode{pos: Pos{r: 0, c: 1}, index: 0, b: '>'}
	head.next = &PosNode{pos: Pos{r: 1, c: 0}, index: 1, b: 'v', prev: head}
	head.next.next = &PosNode{pos: Pos{r: 0, c: -1}, index: 2, b: '<', prev: head.next}
	head.next.next.next = &PosNode{pos: Pos{r: -1, c: 0}, index: 3, b: '^', prev: head.next.next, next: head}
	head.prev = head.next.next.next

	cur := head

	var position Pos
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '.' {
				position = Pos{r: r, c: c}
				break
			}
		}
		if position.r != 0 || position.c != 0 {
			break
		}
	}

	re := regexp.MustCompile(`\d+[RL]?`)

	show[position.r][position.c] = cur.b
	for _, move := range re.FindAllString(moves, -1) {
		fmt.Println(move, "before turn", string(cur.b))
		var digits string
		if strings.Contains(move, "R") || strings.Contains(move, "L") {
			digits = move[:len(move)-1]
		} else {
			digits = move
		}

		count64, _ := strconv.ParseInt(digits, 10, 32)
		count := int(count64)
		turn := move[len(move)-1]

		for i := 0; i < count; i++ {
			newPosition := addPositions(position, cur.pos)

			if newPosition.IsOutOfGrid(grid) || grid[newPosition.r][newPosition.c] == ' ' {
				breakFromLoop := false
				switch {
				case cur.pos.c > 0:
					cc := 0
					for grid[position.r][cc] == ' ' {
						cc++
					}
					if grid[position.r][cc] == '#' {
						breakFromLoop = true
					}
					if grid[position.r][cc] == '.' {
						newPosition.c = cc
					}
				case cur.pos.c < 0:
					cc := len(grid[position.r]) - 1
					for grid[position.r][cc] == ' ' {
						cc--
					}
					if grid[position.r][cc] == '#' {
						breakFromLoop = true
					}
					if grid[position.r][cc] == '.' {
						newPosition.c = cc
					}
				case cur.pos.r > 0:
					rr := 0
					for grid[rr][position.c] == ' ' {
						rr++
					}
					if grid[rr][position.c] == '#' {
						breakFromLoop = true
					}
					if grid[rr][position.c] == '.' {
						newPosition.r = rr
					}
				case cur.pos.r < 0:
					rr := len(grid) - 1
					for grid[rr][position.c] == ' ' {
						rr--
					}
					if grid[rr][position.c] == '#' {
						breakFromLoop = true
					}
					if grid[rr][position.c] == '.' {
						newPosition.r = rr
					}
				}

				if breakFromLoop {
					break
				}
			}

			if grid[newPosition.r][newPosition.c] == '#' {
				break
			}

			position = newPosition
			show[position.r][position.c] = cur.b
		}

		if turn == 'R' {
			cur = cur.next
		}
		if turn == 'L' {
			cur = cur.prev
		}

		show[position.r][position.c] = 'X'
	}

	show[position.r][position.c] = 'F'
	printMap(show)

	fmt.Println(position.r, position.c, cur.index, string(cur.b))
	fmt.Println(1000*(position.r+1) + 4*(position.c+1) + cur.index)
}
