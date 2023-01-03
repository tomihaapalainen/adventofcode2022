package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x, y, z int
}

func adv18a() {
	file, _ := os.Open("input18.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	cubes := make(map[Cube]bool)
	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		split := strings.Split(row, ",")
		si, sj, sk := split[0], split[1], split[2]

		i, _ := strconv.ParseInt(si, 10, 32)
		j, _ := strconv.ParseInt(sj, 10, 32)
		k, _ := strconv.ParseInt(sk, 10, 32)
		cubes[Cube{x: int(i), y: int(j), z: int(k)}] = true
	}

	exposedCount := 0

	for cube := range cubes {
		c := Cube{x: cube.x + 1, y: cube.y, z: cube.z}
		if !cubes[c] {
			exposedCount++
		}
		c = Cube{x: cube.x, y: cube.y + 1, z: cube.z}
		if !cubes[c] {
			exposedCount++
		}
		c = Cube{x: cube.x, y: cube.y, z: cube.z + 1}
		if !cubes[c] {
			exposedCount++
		}
		c = Cube{x: cube.x - 1, y: cube.y, z: cube.z}
		if !cubes[c] {
			exposedCount++
		}
		c = Cube{x: cube.x, y: cube.y - 1, z: cube.z}
		if !cubes[c] {
			exposedCount++
		}
		c = Cube{x: cube.x, y: cube.y, z: cube.z - 1}
		if !cubes[c] {
			exposedCount++
		}
	}

	fmt.Println(exposedCount)
}
