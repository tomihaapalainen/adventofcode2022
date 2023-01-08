package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func reverseSnafuToDecimal(snafu []byte) int {
	decimal := 0.0

	for pow, b := range snafu {
		var mul float64
		if b == '=' {
			mul = -2.0
		} else if b == '-' {
			mul = -1.0
		} else if b == '2' {
			mul = 2.0
		} else if b == '1' {
			mul = 1.0
		} else {
			mul = 0.0
		}

		decimal += mul * math.Pow(5.0, float64(pow))
	}
	return int(decimal)
}

func decimalToSnafu(decimal int) string {
	if decimal == 0 {
		return ""
	}

	decimalRemainder := decimal % 5
	snafuDigit := []byte{'0', '1', '2', '=', '-'}[decimalRemainder]
	newDecimal := (decimal + 2) / 5
	snafu := decimalToSnafu(newDecimal)
	snafu += string(snafuDigit)
	return snafu
}

func adv25a() {
	file, _ := os.Open("input25.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	snafus := [][]byte{}
	decimals := []int{}

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		snafu := []byte(row)
		for i, j := 0, len(snafu)-1; i < j; i, j = i+1, j-1 {
			snafu[i], snafu[j] = snafu[j], snafu[i]
		}
		snafus = append(snafus, snafu)
	}

	for _, snafu := range snafus {
		decimal := reverseSnafuToDecimal(snafu)
		decimals = append(decimals, int(decimal))
	}

	integer := 0
	for _, d := range decimals {
		integer += d
	}

	fmt.Println(integer)
	fmt.Println(decimalToSnafu(integer))
}
