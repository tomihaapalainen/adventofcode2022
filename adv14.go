package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Obj struct {
	x int64
	y int64
}

func move(smap map[Obj]bool, rmap map[Obj]bool, s Obj, maxy int) int {
	if int(s.y) > maxy {
		return -1
	}

	b := Obj{x: s.x, y: s.y + 1}

	bb := smap[b] || rmap[b]

	if !bb {
		return move(smap, rmap, b, maxy)
	}

	l := Obj{x: s.x - 1, y: s.y + 1}
	r := Obj{x: s.x + 1, y: s.y + 1}

	lb := smap[l] || rmap[l]
	rb := smap[r] || rmap[r]

	if lb && rb {
		smap[s] = true
		return 1
	}

	if !lb {
		return move(smap, rmap, l, maxy)
	}
	if !rb {
		return move(smap, rmap, r, maxy)
	}

	return 0
}

func adv14a() {
	file, _ := os.Open("input14.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	rockmap := make(map[Obj]bool)

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		split := strings.Split(row, " -> ")
		for i := 0; i < len(split)-1; i++ {
			l := split[i]
			r := split[i+1]

			splitl := strings.Split(l, ",")
			splitr := strings.Split(r, ",")

			x1, _ := strconv.ParseInt(splitl[0], 10, 32)
			y1, _ := strconv.ParseInt(splitl[1], 10, 32)

			x2, _ := strconv.ParseInt(splitr[0], 10, 32)
			y2, _ := strconv.ParseInt(splitr[1], 10, 32)

			if x1 != x2 {
				if x2 > x1 {
					for x := x1; x <= x2; x++ {
						rockmap[Obj{x: x, y: y1}] = true
					}
				} else {
					for x := x2; x <= x1; x++ {
						rockmap[Obj{x: x, y: y1}] = true
					}
				}
			} else if y1 != y2 {
				if y2 > y1 {
					for y := y1; y <= y2; y++ {
						rockmap[Obj{x: x1, y: y}] = true
					}
				} else {
					for y := y2; y <= y1; y++ {
						rockmap[Obj{x: x1, y: y}] = true
					}
				}
			}
		}
	}

	for k, v := range rockmap {
		fmt.Println(k, v)
	}

	sandmap := make(map[Obj]bool)

	maxy := 0
	for k := range rockmap {
		if int(k.y) > maxy {
			maxy = int(k.y)
		}
	}

	count := 0
	sand := Obj{x: 500, y: 0}

	for {
		result := move(sandmap, rockmap, sand, maxy)

		if result > 0 {
			count++
			sand = Obj{x: 500, y: 0}
		}

		if result < 0 {
			break
		}
	}

	fmt.Println(count)
}

func move2(smap map[Obj]bool, rmap map[Obj]bool, s Obj) int {
	if smap[Obj{x: 500, y: 0}] {
		return -1
	}

	b := Obj{x: s.x, y: s.y + 1}

	bb := smap[b] || rmap[b]

	if !bb {
		return move2(smap, rmap, b)
	}

	l := Obj{x: s.x - 1, y: s.y + 1}
	r := Obj{x: s.x + 1, y: s.y + 1}

	lb := smap[l] || rmap[l]
	rb := smap[r] || rmap[r]

	if lb && rb {
		smap[s] = true
		return 1
	}

	if !lb {
		return move2(smap, rmap, l)
	}
	if !rb {
		return move2(smap, rmap, r)
	}

	return 0
}

func adv14b() {
	file, _ := os.Open("input14.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	rockmap := make(map[Obj]bool)

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		split := strings.Split(row, " -> ")
		for i := 0; i < len(split)-1; i++ {
			l := split[i]
			r := split[i+1]

			splitl := strings.Split(l, ",")
			splitr := strings.Split(r, ",")

			x1, _ := strconv.ParseInt(splitl[0], 10, 32)
			y1, _ := strconv.ParseInt(splitl[1], 10, 32)

			x2, _ := strconv.ParseInt(splitr[0], 10, 32)
			y2, _ := strconv.ParseInt(splitr[1], 10, 32)

			if x1 != x2 {
				if x2 > x1 {
					for x := x1; x <= x2; x++ {
						rockmap[Obj{x: x, y: y1}] = true
					}
				} else {
					for x := x2; x <= x1; x++ {
						rockmap[Obj{x: x, y: y1}] = true
					}
				}
			} else if y1 != y2 {
				if y2 > y1 {
					for y := y1; y <= y2; y++ {
						rockmap[Obj{x: x1, y: y}] = true
					}
				} else {
					for y := y2; y <= y1; y++ {
						rockmap[Obj{x: x1, y: y}] = true
					}
				}
			}
		}
	}

	sandmap := make(map[Obj]bool)

	maxy := 0
	minx := 0
	maxx := 0
	for k := range rockmap {
		if int(k.y) > maxy {
			maxy = int(k.y)
		}
		if int(k.x) < minx {
			minx = int(k.x)
		}
		if int(k.x) > maxx {
			maxx = int(k.x)
		}
	}

	count := 0
	sand := Obj{x: 500, y: 0}

	for i := minx - 1000; i <= maxx+1000; i++ {
		rockmap[Obj{x: int64(i), y: int64(maxy + 2)}] = true
	}

	for {
		result := move2(sandmap, rockmap, sand)

		if result > 0 {
			count++
			sand = Obj{x: 500, y: 0}
		}

		if result < 0 {
			break
		}
	}

	fmt.Println(count)
}
