package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Id         int
	Items      []int
	Operation  string
	OperationL string
	OperationR string
	Test       int
	True       int
	False      int
	Inspected  int
}

func newMonkey(id int) Monkey {
	return Monkey{Id: id, Items: []int{}, Operation: "", OperationL: "", OperationR: "", Test: 0, True: 0, False: 0}
}

func readMonkeys() []Monkey {
	file, _ := os.Open("input11.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	monkeys := []Monkey{}

	var monkey Monkey
	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())

		switch {
		case row == "":
			monkeys = append(monkeys, monkey)

		case strings.Contains(row, "Monkey"):
			row = strings.Trim(row, ":")
			split := strings.Split(row, " ")
			id, _ := strconv.ParseInt(split[1], 10, 64)
			monkey = newMonkey(int(id))

		case strings.Contains(row, "Starting items: "):
			split := strings.Split(row, "Starting items: ")
			split2 := strings.Split(split[1], ", ")
			items := []int{}

			for _, v := range split2 {
				item, _ := strconv.ParseInt(v, 10, 64)
				items = append(items, int(item))
			}

			monkey.Items = items

		case strings.Contains(row, "Operation: "):
			split := strings.Split(row, "Operation: new = ")
			split2 := strings.Split(split[1], " ")
			monkey.OperationL = split2[0]
			monkey.Operation = split2[1]
			monkey.OperationR = split2[2]

		case strings.Contains(row, "Test: "):
			split := strings.Split(row, "divisible by ")
			div, _ := strconv.ParseInt(split[1], 10, 64)
			monkey.Test = int(div)

		case strings.Contains(row, "If true:"):
			split := strings.Split(row, "throw to monkey ")
			to, _ := strconv.ParseInt(split[1], 10, 64)
			monkey.True = int(to)

		case strings.Contains(row, "If false:"):
			split := strings.Split(row, "throw to monkey ")
			to, _ := strconv.ParseInt(split[1], 10, 64)
			monkey.False = int(to)
		}
	}

	monkeys = append(monkeys, monkey)

	return monkeys
}

func adv11a() {
	monkeys := readMonkeys()

	for i := 0; i < 20; i++ {
		for mi := range monkeys {
			for _, level := range monkeys[mi].Items {
				var r int
				if monkeys[mi].OperationR == "old" {
					r = level
				} else {
					temp, _ := strconv.ParseInt(monkeys[mi].OperationR, 10, 64)
					r = int(temp)
				}
				if monkeys[mi].Operation == "*" {
					level *= r
				} else {
					level += r
				}

				level /= 3

				if level%monkeys[mi].Test == 0 {
					monkeys[monkeys[mi].True].Items = append(monkeys[monkeys[mi].True].Items, level)
				} else {
					monkeys[monkeys[mi].False].Items = append(monkeys[monkeys[mi].False].Items, level)
				}

				monkeys[mi].Inspected++
			}
			monkeys[mi].Items = []int{}
		}
	}

	inspected := []int{}

	for _, m := range monkeys {
		inspected = append(inspected, m.Inspected)
	}

	sort.Ints(inspected)

	fmt.Println(inspected[len(inspected)-2], inspected[len(inspected)-1])
	fmt.Println(inspected[len(inspected)-2] * inspected[len(inspected)-1])
}

func adv11b() {
	monkeys := readMonkeys()

	asdf := 1
	for _, m := range monkeys {
		asdf *= m.Test
	}

	for i := 1; i <= 10000; i++ {
		for mi := range monkeys {
			for _, level := range monkeys[mi].Items {
				var r int
				if monkeys[mi].OperationR == "old" {
					r = level
				} else {
					temp, _ := strconv.ParseInt(monkeys[mi].OperationR, 10, 64)
					r = int(temp)
				}
				if monkeys[mi].Operation == "*" {
					level *= r
				} else {
					level += r
				}

				level = level % asdf

				if level%monkeys[mi].Test == 0 {
					monkeys[monkeys[mi].True].Items = append(monkeys[monkeys[mi].True].Items, level)
				} else {
					monkeys[monkeys[mi].False].Items = append(monkeys[monkeys[mi].False].Items, level)
				}

				monkeys[mi].Inspected++
			}
			monkeys[mi].Items = []int{}
		}

		if i == 20 || i == 1000 || i == 2000 || i == 3000 || i == 4000 {
			for _, m := range monkeys {
				fmt.Print(m.Inspected, " ")
			}
			fmt.Print("\n")
		}
	}

	inspected := []int{}

	for _, m := range monkeys {
		inspected = append(inspected, m.Inspected)
	}

	sort.Ints(inspected)

	fmt.Println(inspected[len(inspected)-2], inspected[len(inspected)-1])
	fmt.Println(inspected[len(inspected)-2] * inspected[len(inspected)-1])
}
