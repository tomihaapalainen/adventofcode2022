package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Directions struct {
	data [][]Pos
}

func (d *Directions) Rotate() {
	first := d.data[0]
	d.data = d.data[1:]
	d.data = append(d.data, first)
}

func printGrid(step int, grid map[Pos]bool) {
	minr, maxr := 1000, -1000
	minc, maxc := 1000, -1000

	for k := range grid {
		if k.r < minr {
			minr = k.r
		}
		if k.r > maxr {
			maxr = k.r
		}
		if k.c < minc {
			minc = k.c
		}
		if k.c > maxc {
			maxc = k.c
		}
	}

	rows := [][]byte{}
	for r := minr; r <= maxr; r++ {
		row := []byte{}
		for c := minc; c <= maxc; c++ {
			if grid[Pos{r: r, c: c}] {
				row = append(row, '#')
			} else {
				row = append(row, '.')
			}
		}

		rows = append(rows, row)
	}

	fmt.Println("-----", step, "-----")
	for _, row := range rows {
		for _, b := range row {
			fmt.Print(string(b))
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("result:", (maxr-minr+1)*(maxc-minc+1)-len(grid))

	fmt.Println()
}

func adv23a() {
	file, _ := os.Open("input23.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	grid := make(map[Pos]bool)

	r := 0
	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		for c, b := range row {
			if b == '#' {
				grid[Pos{r: r, c: c}] = true
			}
		}
		r++
	}

	top := []Pos{Pos{r: -1, c: -1}, Pos{r: -1, c: 0}, Pos{r: -1, c: 1}}
	bottom := []Pos{Pos{r: 1, c: -1}, Pos{r: 1, c: 0}, Pos{r: 1, c: 1}}
	left := []Pos{Pos{r: -1, c: -1}, Pos{r: 0, c: -1}, Pos{r: 1, c: -1}}
	right := []Pos{Pos{r: -1, c: 1}, Pos{r: 0, c: 1}, Pos{r: 1, c: 1}}

	directions := Directions{data: [][]Pos{top, bottom, left, right}}

	all := []Pos{}
	all = append(all, top...)
	all = append(all, bottom...)
	all = append(all, right...)
	all = append(all, left...)

	for i := 0; i < 10; i++ {
		moves := make(map[Pos]Pos)
		moveCount := make(map[Pos]int)
		for elf := range grid {
			elfAround := false
			for _, a := range all {
				pos := addPositions(elf, a)
				if grid[pos] {
					elfAround = true
					break
				}
			}

			if !elfAround {
				fmt.Println("no elf around, staying still", elf.r, elf.c)
				continue
			}

			for _, dir := range directions.data {
				hasElf := false
				for _, pos := range dir {
					pos = addPositions(elf, pos)
					if grid[pos] {
						hasElf = true
						break
					}
				}

				if !hasElf {
					dst := addPositions(elf, dir[1])
					moves[elf] = dst
					moveCount[dst]++
					break
				}
			}
		}

		for elf, dst := range moves {
			if moveCount[dst] == 1 && !grid[dst] {
				delete(grid, elf)
				grid[dst] = true
			}
		}

		directions.Rotate()
		printGrid(i+1, grid)
	}
}

func adv23b() {
	file, _ := os.Open("input23.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	grid := make(map[Pos]bool)

	r := 0
	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		for c, b := range row {
			if b == '#' {
				grid[Pos{r: r, c: c}] = true
			}
		}
		r++
	}

	top := []Pos{Pos{r: -1, c: -1}, Pos{r: -1, c: 0}, Pos{r: -1, c: 1}}
	bottom := []Pos{Pos{r: 1, c: -1}, Pos{r: 1, c: 0}, Pos{r: 1, c: 1}}
	left := []Pos{Pos{r: -1, c: -1}, Pos{r: 0, c: -1}, Pos{r: 1, c: -1}}
	right := []Pos{Pos{r: -1, c: 1}, Pos{r: 0, c: 1}, Pos{r: 1, c: 1}}

	directions := Directions{data: [][]Pos{top, bottom, left, right}}

	all := []Pos{}
	all = append(all, top...)
	all = append(all, bottom...)
	all = append(all, right...)
	all = append(all, left...)

	i := 1
	for {
		moves := make(map[Pos]Pos)
		moveCount := make(map[Pos]int)
		for elf := range grid {
			elfAround := false
			for _, a := range all {
				pos := addPositions(elf, a)
				if grid[pos] {
					elfAround = true
					break
				}
			}

			if !elfAround {
				continue
			}

			for _, dir := range directions.data {
				hasElf := false
				for _, pos := range dir {
					pos = addPositions(elf, pos)
					if grid[pos] {
						hasElf = true
						break
					}
				}

				if !hasElf {
					dst := addPositions(elf, dir[1])
					moves[elf] = dst
					moveCount[dst]++
					break
				}
			}
		}

		movesMade := 0
		for elf, dst := range moves {
			if moveCount[dst] == 1 && !grid[dst] {
				delete(grid, elf)
				grid[dst] = true
				movesMade++
			}
		}

		if movesMade == 0 {
			break
		}

		directions.Rotate()
		i++
	}

	printGrid(i, grid)
}
