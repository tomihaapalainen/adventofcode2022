package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func adv6(characterCount int) {
	file, _ := os.Open("input6.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	var input string
	for fileScanner.Scan() {
		input = strings.TrimSpace(fileScanner.Text())
	}

	for l := 0; l < len(input)-characterCount; l++ {
		r := l + characterCount

		m := make(map[rune]int)
		for _, r := range input[l:r] {
			m[r] = 1
		}
		if len(m) == characterCount {
			fmt.Println(r)
			break
		}
	}
}

func adv6a() {
	adv6(4)
}

func adv6b() {
	adv6(14)
}
