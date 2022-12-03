package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func adv3a() {
	file, _ := os.Open("input3.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		row := []byte(strings.TrimSpace(fileScanner.Text()))
		left, right := row[0:len(row)/2], row[len(row)/2:]
		m := make(map[byte]int)
		for _, l := range left {
			m[l] += 1
		}
		for _, r := range right {
			if _, ok := m[r]; ok {
				if 65 <= r && r <= 90 {
					sum += int(r) - 65 + 27
				} else if 95 <= r && r <= 122 {
					sum += int(r) - 96
				}
				break
			}
		}
	}

	fmt.Println(sum)
}

func adv3b() {
	file, _ := os.Open("input3.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	rows := [][]byte{}
	for fileScanner.Scan() {
		row := []byte(strings.TrimSpace(fileScanner.Text()))
		rows = append(rows, row)
	}

	sum := 0
	for i := 0; i <= len(rows)-3; i = i + 3 {
		am := make(map[byte]int)
		for _, b := range rows[i] {
			am[b] += 1
		}

		bm := make(map[byte]int)
		for _, b := range rows[i+1] {
			if am[b] > 0 {
				bm[b] += 1
			}
		}

		for _, b := range rows[i+2] {
			if bm[b] > 0 {
				if 65 <= b && b <= 90 {
					sum += int(b) - 65 + 27
				} else if 95 <= b && b <= 122 {
					sum += int(b) - 96
				}
				break
			}
		}
	}

	fmt.Println(sum)
}
