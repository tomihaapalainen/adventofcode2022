package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	Data []string
}

func NewStack(data []string) Stack {
	s := Stack{}
	for i := len(data) - 1; i >= 0; i-- {
		s.Push(data[i])
	}
	return s
}

func (stack *Stack) Top() string {
	return stack.Data[len(stack.Data)-1]
}

func (stack *Stack) Pop() string {
	last := stack.Data[len(stack.Data)-1]
	stack.Data = stack.Data[:len(stack.Data)-1]
	return last
}

func (stack *Stack) PopMany(count int) []string {
	items := stack.Data[len(stack.Data)-count:]
	stack.Data = stack.Data[:len(stack.Data)-count]
	return items
}

func (stack *Stack) Push(s string) {
	stack.Data = append(stack.Data, s)
}

func (stack *Stack) PushMany(items []string) {
	stack.Data = append(stack.Data, items...)
}

type Move struct {
	Count int
	From  int
	To    int
}

func readStacksAndMoves() ([]Stack, []Move) {
	file, _ := os.Open("input5.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	data := make(map[int][]string)
	emptyLineRead := false
	moves := []Move{}
	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		if len(row) == 0 {
			emptyLineRead = true
			continue
		}

		if !emptyLineRead {
			count := 0
			for i := range row {
				if i == count*4+1 {
					if row[i] != 32 {
						data[count] = append(data[count], string(row[i]))
					}
					count++
				}
			}
		} else {
			re := regexp.MustCompile(`\d+`)
			res := re.FindAll([]byte(row), 3)
			count, _ := strconv.ParseInt(string(res[0]), 10, 32)
			from, _ := strconv.ParseInt(string(res[1]), 10, 32)
			to, _ := strconv.ParseInt(string(res[2]), 10, 32)
			moves = append(moves, Move{Count: int(count), From: int(from), To: int(to)})
		}
	}

	stacks := make([]Stack, len(data))
	for k, v := range data {
		stacks[k] = NewStack(v)
	}

	return stacks, moves
}

func adv5a() {
	stacks, moves := readStacksAndMoves()

	for _, move := range moves {
		fmt.Println(move.Count, move.From, move.To)
		for i := 0; i < move.Count; i++ {
			item := stacks[move.From-1].Pop()
			stacks[move.To-1].Push(item)
		}
	}

	for _, stack := range stacks {
		fmt.Print(stack.Top())
	}
	fmt.Print("\n")
}

func adv5b() {
	stacks, moves := readStacksAndMoves()

	for _, move := range moves {
		fmt.Println(move.Count, move.From, move.To)
		items := stacks[move.From-1].PopMany(move.Count)
		stacks[move.To-1].PushMany(items)
	}

	for _, stack := range stacks {
		fmt.Print(stack.Top())
	}
	fmt.Print("\n")
}
