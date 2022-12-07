package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mapFolderSizes() map[string]int64 {
	file, _ := os.Open("input7.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	cdRe := regexp.MustCompile(`^\$ cd [^\.]+$`)
	fileRe := regexp.MustCompile(`^\d+ .*$`)

	dirs := []string{}
	sizes := make(map[string]int64)

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		switch {
		case line == "$ cd ..":
			dirs = dirs[:len(dirs)-1]
		case cdRe.MatchString(line):
			split := strings.Split(line, " ")
			dir := split[2]
			dirs = append(dirs, dir)
		case fileRe.MatchString(line):
			split := strings.Split(line, " ")
			size, _ := strconv.ParseInt(split[0], 10, 64)

			for i := range dirs {
				path := strings.Join(dirs[:i+1], "/")
				sizes[path] += size
			}
		}
	}

	return sizes
}

func adv7a() {
	sizes := mapFolderSizes()

	var total int64 = 0

	for _, v := range sizes {
		if v <= 100000 {
			total += v
		}
	}

	fmt.Println(total)
}

func adv7b() {
	sizes := mapFolderSizes()

	usedSpace := sizes["/"]
	freeSpace := 70000000 - usedSpace
	spaceToBeFreed := 30000000 - freeSpace

	var result int64
	for _, v := range sizes {
		if v >= spaceToBeFreed {
			if result == 0 || v < result {
				result = v
			}
		}
	}

	fmt.Println(result)
}
