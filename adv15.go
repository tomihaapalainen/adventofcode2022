package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calculateManhattanDistance(a, b Point) int {
	return int(math.Abs(float64(a.c)-float64(b.c)) + math.Abs(float64(a.r)-float64(b.r)))
}

func adv15a() {
	file, _ := os.Open("input15.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	sensorMap := make(map[Point]Point)
	beaconMap := make(map[Point]bool)

	minx := 1000
	maxx := 0

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		sample := regexp.MustCompile(`-?\d+`)
		res := sample.FindAllString(row, 4)
		sx, _ := strconv.ParseInt(res[0], 10, 32)
		sy, _ := strconv.ParseInt(res[1], 10, 32)
		bx, _ := strconv.ParseInt(res[2], 10, 32)
		by, _ := strconv.ParseInt(res[3], 10, 32)

		if int(sx) < minx {
			minx = int(sx)
		}

		if int(bx) < minx {
			minx = int(bx)
		}

		if int(sx) > maxx {
			maxx = int(sx)
		}

		if int(bx) > maxx {
			maxx = int(bx)
		}

		s := Point{r: int(sy), c: int(sx)}
		b := Point{r: int(by), c: int(bx)}

		sensorMap[s] = b
		beaconMap[b] = true
	}

	count := 0
	for c := minx; c <= maxx; c++ {
		p := Point{r: 2000000, c: c}

		if beaconMap[p] {
			continue
		}

		for s := range sensorMap {
			if calculateManhattanDistance(p, s) <= calculateManhattanDistance(s, sensorMap[s]) {
				count++
				break
			}
		}
	}

	c := minx - 1
	for {
		p := Point{r: 2000000, c: c}
		added := false
		for s := range sensorMap {
			if calculateManhattanDistance(p, s) <= calculateManhattanDistance(s, sensorMap[s]) {
				count++
				added = true
				break
			}
		}
		if !added {
			break
		}
		c--
	}

	c = maxx + 1
	for {
		added := false
		for s := range sensorMap {
			p := Point{r: 2000000, c: c}
			if calculateManhattanDistance(p, s) <= calculateManhattanDistance(s, sensorMap[s]) {
				count++
				added = true
				break
			}
		}
		if !added {
			break
		}
		c++
	}

	fmt.Println(count)
}

func calculate(rs, re, cs, ce int, sensorMap map[Point]Point) {
	for r := rs; r <= re; r++ {
		for c := cs; c <= ce; c++ {
			p := Point{r: r, c: c}

			inRange := false
			for s := range sensorMap {
				if calculateManhattanDistance(p, s) <= calculateManhattanDistance(s, sensorMap[s]) {
					inRange = true
					break
				}
			}

			if !inRange {
				fmt.Println(p)
				fmt.Println(4000000*p.c + p.r)
			}
		}
	}
}

func adv15b() {
	file, _ := os.Open("input15.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	sensorMap := make(map[Point]Point)

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		sample := regexp.MustCompile(`-?\d+`)
		res := sample.FindAllString(row, 4)
		sx, _ := strconv.ParseInt(res[0], 10, 32)
		sy, _ := strconv.ParseInt(res[1], 10, 32)
		bx, _ := strconv.ParseInt(res[2], 10, 32)
		by, _ := strconv.ParseInt(res[3], 10, 32)

		s := Point{r: int(sy), c: int(sx)}
		b := Point{r: int(by), c: int(bx)}

		sensorMap[s] = b
	}

	pts := []Point{}
	for s := range sensorMap {
		d := calculateManhattanDistance(s, sensorMap[s])

		start := Point{r: s.r - d, c: s.c}
		start2 := Point{r: s.r + d, c: s.c}

		for i := 0; i <= d; i++ {
			p := Point{r: start.r + i, c: start.c + i + 1}
			if p.r >= 0 && p.r <= 4000000 && p.c >= 0 && p.c <= 4000000 {
				pts = append(pts, p)
			}

			p = Point{r: start2.r - i, c: start2.c - i + 1}
			if p.r >= 0 && p.r <= 4000000 && p.c >= 0 && p.c <= 4000000 {
				pts = append(pts, p)
			}
		}
	}

	for _, pt := range pts {
		inRange := false
		for s := range sensorMap {
			if calculateManhattanDistance(pt, s) <= calculateManhattanDistance(s, sensorMap[s]) {
				inRange = true
				break
			}
		}

		if !inRange {
			fmt.Println(pt)
			fmt.Println(pt.c*4000000 + pt.r)
			return
		}
	}
}
