package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func adv4a() {
	file, _ := os.Open("input4.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		ranges := strings.Split(row, ",")
		firstRange := strings.Split(ranges[0], "-")
		a, _ := strconv.ParseInt(firstRange[0], 10, 32)
		b, _ := strconv.ParseInt(firstRange[1], 10, 32)

		secondRange := strings.Split(ranges[1], "-")
		c, _ := strconv.ParseInt(secondRange[0], 10, 32)
		d, _ := strconv.ParseInt(secondRange[1], 10, 32)

		if (c <= a && b <= d) || (a <= c && d <= b) {
			sum++
		}
	}

	fmt.Println(sum)
}

func adv4b() {
	file, _ := os.Open("input4.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		ranges := strings.Split(row, ",")
		firstRange := strings.Split(ranges[0], "-")
		a, _ := strconv.ParseInt(firstRange[0], 10, 64)
		b, _ := strconv.ParseInt(firstRange[1], 10, 64)

		secondRange := strings.Split(ranges[1], "-")
		c, _ := strconv.ParseInt(secondRange[0], 10, 64)
		d, _ := strconv.ParseInt(secondRange[1], 10, 64)

		if ((c <= a && a <= d) || (c <= b && b <= d)) || ((a <= c && c <= b) || (a <= d && d <= b)) {
			sum++
		}
	}

	fmt.Println(sum)
}
