package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func runOperation(a, b int, op string) int {
	switch {
	case op == "+":
		return a + b
	case op == "-":
		return a - b
	case op == "*":
		return a * b
	case op == "/":
		return a / b
	}
	return 0
}

func isNum(x string) bool {
	_, err := strconv.ParseInt(x, 10, 64)
	return err == nil
}

func isValidExpression(x string) bool {
	re := regexp.MustCompile(`(\d+|humn) [+-/*] (\d+|humn)`)
	return re.Match([]byte(x))
}

func adv21a() {
	file, _ := os.Open("input21.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	ex := make(map[string]string)
	vals := make(map[string]string)

	for fileScanner.Scan() {
		row := strings.TrimSpace(fileScanner.Text())
		split1 := strings.Split(row, ": ")
		name := split1[0]
		expression := split1[1]

		if isNum(expression) {
			vals[name] = expression
		} else {
			ex[name] = expression
		}
	}

	for {
		for k, v := range vals {
			for e := range ex {
				ex[e] = strings.Replace(ex[e], k, v, 1)
				if isValidExpression(ex[e]) {
					split := strings.Split(ex[e], " ")
					l, op, r := split[0], split[1], split[2]
					li, _ := strconv.ParseInt(l, 10, 64)
					ri, _ := strconv.ParseInt(r, 10, 64)
					vals[e] = fmt.Sprint(runOperation(int(li), int(ri), op))
					delete(ex, e)
				}
			}
		}
		if len(ex) == 0 {
			break
		}
	}

	fmt.Println(vals["hqpw"], vals["nprj"])
	fmt.Println(vals["root"])
}
