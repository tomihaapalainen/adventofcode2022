package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	i    int
	val  int
	prev *ListNode
	next *ListNode
}

func printList(head *ListNode) {
	cur := head
	exit := head

	for {
		fmt.Print(cur.val, " ")
		cur = cur.next
		if cur == exit {
			fmt.Println()
			break
		}
	}
}

func adv20a() {
	file, _ := os.Open("input20.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	original := []int{}
	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		rowInt, _ := strconv.ParseInt(row, 10, 32)
		original = append(original, int(rowInt))
	}

	var head *ListNode
	var cur *ListNode
	var prev *ListNode

	for i, value := range original {
		if head == nil {
			head = &ListNode{i: i, val: value}
			cur = head
		} else {
			cur.next = &ListNode{i: i, val: value}
			prev = cur
			cur = cur.next
			cur.prev = prev
		}
	}

	cur.next = head
	head.prev = cur

	for i, o := range original {
		cur := head

		for cur.i != i {
			cur = cur.next
		}

		if o < 0 {
			for l := o; l < 0; l++ {
				cur.val, cur.prev.val = cur.prev.val, cur.val
				cur.i, cur.prev.i = cur.prev.i, cur.i
				cur = cur.prev
			}
		} else {
			for l := o; l > 0; l-- {
				cur.val, cur.next.val = cur.next.val, cur.val
				cur.i, cur.next.i = cur.next.i, cur.i
				cur = cur.next
			}
		}
	}

	cur = head
	for cur.val != 0 {
		cur = cur.next
	}

	result := 0
	for i := 0; i <= 3000; i++ {
		if i%1000 == 0 {
			result += cur.val
		}
		cur = cur.next
	}

	fmt.Println(result)
}
