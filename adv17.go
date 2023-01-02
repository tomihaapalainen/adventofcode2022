package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pieces struct {
	data []Piece
	i    int
}

func (pieces *Pieces) Next() Piece {
	result := pieces.data[pieces.i]
	pieces.i++

	if pieces.i == len(pieces.data) {
		pieces.i = 0
	}

	newRocks := []Rock{}
	newRocks = append(newRocks, result.rocks...)
	return Piece{rocks: newRocks}
}

type Rock struct {
	x int
	y int
}

type Piece struct {
	rocks []Rock
}

func (p *Piece) canMoveX(floor map[Rock]bool, move int) bool {
	for _, r := range p.rocks {
		newRock := Rock{x: r.x + move, y: r.y}
		if newRock.x == 0 || newRock.x == 8 {
			return false
		}
		if floor[newRock] {
			return false
		}
	}
	return true
}

func (p *Piece) canMoveY(floor map[Rock]bool) bool {
	for _, r := range p.rocks {
		if floor[Rock{x: r.x, y: r.y + 1}] {
			return false
		}
	}
	return true
}

func (p *Piece) Height() int {
	heights := []int{}
	for r := range p.rocks {
		heights = append(heights, p.rocks[r].y)
	}
	sort.Ints(heights)
	return heights[len(heights)-1] - heights[0] + 1
}

func (p *Piece) Width() int {
	widths := []int{}
	for r := range p.rocks {
		widths = append(widths, p.rocks[r].x)
	}
	sort.Ints(widths)
	return widths[len(widths)-1] - widths[0] + 1
}

type Cycle struct {
	data []int
	cur  int
}

func (c *Cycle) Next() int {
	result := c.data[c.cur]
	c.cur++

	if c.cur == len(c.data) {
		c.cur = 0
	}

	return result
}

func adv17a() {
	file, _ := os.Open("input17-rocks.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	rocks := Pieces{data: []Piece{}}
	piece := Piece{rocks: []Rock{}}
	y := 0

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())

		if len(row) == 0 {
			rocks.data = append(rocks.data, piece)
			piece = Piece{rocks: []Rock{}}
			y = 0
		} else {
			for x := range row {
				if row[x] == '#' {
					piece.rocks = append(piece.rocks, Rock{x: x, y: y})
				}
			}
			y++
		}
	}
	rocks.data = append(rocks.data, piece)

	for d := range rocks.data {
		pieceHeight := rocks.data[d].Height()
		for r := range rocks.data[d].rocks {
			rocks.data[d].rocks[r].y -= pieceHeight - 1
		}
	}

	file, _ = os.Open("input17.txt")

	fileScanner = bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	cycle := Cycle{data: []int{}, cur: 0}

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		for i := range row {
			if row[i] == '<' {
				cycle.data = append(cycle.data, -1)
			} else {
				cycle.data = append(cycle.data, 1)
			}
		}
	}

	floor := make(map[Rock]bool)

	for i := 1; i <= 7; i++ {
		floor[Rock{x: i, y: 0}] = true
	}

	for i := 0; i < 2022; i++ {
		piece := rocks.Next()

		miny := 0
		for k := range floor {
			if k.y < miny {
				miny = k.y
			}
		}

		bottomY := -4 + miny
		for r := range piece.rocks {
			piece.rocks[r].x += 3
			piece.rocks[r].y += bottomY
		}

		for {
			move := cycle.Next()
			if piece.canMoveX(floor, move) {
				for p := range piece.rocks {
					piece.rocks[p].x += move
				}
			}

			if piece.canMoveY(floor) {
				for p := range piece.rocks {
					piece.rocks[p].y += 1
				}
			} else {
				for p := range piece.rocks {
					floor[piece.rocks[p]] = true
				}
				break
			}
		}
	}

	miny := 0
	minx := 0
	maxx := 6
	for k := range floor {
		if k.y < miny {
			miny = k.y
		}
		if k.x < minx {
			minx = k.x
		}
		if k.x > maxx {
			maxx = k.x
		}
	}

	fmt.Println(minx, maxx)
	fmt.Println(-miny)
}
