package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	r int
	c int
}

type Queue struct {
	data [][]Point
}

func (q *Queue) Enq(p []Point) {
	q.data = append(q.data, p)
}

func (q *Queue) Deq() []Point {
	first := q.data[0]
	q.data = q.data[1:]
	return first
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func BFS(graph map[Point][]Point, start Point, end Point) {
	queue := Queue{}
	visited := make(map[Point]bool)

	queue.Enq([]Point{start})

	for !queue.Empty() {
		path := queue.Deq()
		last := path[len(path)-1]

		if visited[last] {
			continue
		}
		visited[last] = true

		if last == end {
			fmt.Println(path)
			fmt.Println(len(path) - 1)
		}

		for n := range graph[last] {
			for p := range path {
				if graph[last][n] == path[p] {
					continue
				}
			}

			newPath := []Point{}
			newPath = append(newPath, path...)
			newPath = append(newPath, graph[last][n])
			queue.Enq(newPath)
		}
	}
}

func adv12a() {
	file, _ := os.Open("input12.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	grid := [][]byte{}

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		grid = append(grid, []byte(row))
	}

	var start Point
	var end Point

	for r := range grid {
		for c := range grid[r] {
			b := grid[r][c]
			if b == 'S' {
				start = Point{r: r, c: c}
				grid[r][c] = 'd'
			}
			if b == 'E' {
				end = Point{r: r, c: c}
				grid[r][c] = 'z'
			}
		}
	}

	graph := make(map[Point][]Point)

	for r := range grid {
		for c := range grid[r] {
			nbrs := []Point{}

			if 0 <= r-1 && grid[r-1][c]-grid[r][c] <= 1 {
				nbrs = append(nbrs, Point{r: r - 1, c: c})
			}
			if 0 <= c-1 && grid[r][c-1]-grid[r][c] <= 1 {
				nbrs = append(nbrs, Point{r: r, c: c - 1})
			}
			if r+1 < len(grid) && grid[r+1][c]-grid[r][c] <= 1 {
				nbrs = append(nbrs, Point{r: r + 1, c: c})
			}
			if c+1 < len(grid[0]) && grid[r][c+1]-grid[r][c] <= 1 {
				nbrs = append(nbrs, Point{r: r, c: c + 1})
			}

			graph[Point{r: r, c: c}] = nbrs
		}
	}

	BFS(graph, start, end)
}
