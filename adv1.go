package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func adv1() {
	file, _ := os.Open("input1.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	arr := []int{}

	i := 0
	var value int = 0
	for fileScanner.Scan() {
		val, err := strconv.ParseInt(fileScanner.Text(), 10, 32)
		if err != nil {
			arr = append(arr, value)
			value = 0
			i++
		} else {
			value += int(val)
		}
	}

	sort.Ints(arr)

	fmt.Println(arr[len(arr)-1])

	sum := 0
	for i := len(arr) - 1; i > len(arr)-4; i-- {
		sum += arr[i]
	}

	fmt.Println(sum)
}
