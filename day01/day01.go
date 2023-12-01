package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	var sum1, sum2 int
	for s.Scan() {
		sum1 += parseLine1(s.Text())
		sum2 += parseLine2(s.Text())
	}
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}

func parseLine1(s string) int {
	var first, last byte
	for _, c := range []byte(s) {
		if c >= '0' && c <= '9' {
			if first == 0 {
				first = c
			}
			last = c
		}
	}
	return int(first-'0')*10 + int(last-'0')
}

func parseLine2(s string) int {
	var first, last int
	set := func(v int) {
		if first == 0 {
			first = v
		}
		last = v
	}
	for i := 0; i < len(s); i++ {
		switch c := s[i]; c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			set(int(c - '0'))
		case 'o':
			if strings.HasPrefix(s[i:], "one") {
				set(1)
			}
		case 't':
			if strings.HasPrefix(s[i:], "two") {
				set(2)
			}
			if strings.HasPrefix(s[i:], "three") {
				set(3)
			}
		case 'f':
			if strings.HasPrefix(s[i:], "four") {
				set(4)
			}
			if strings.HasPrefix(s[i:], "five") {
				set(5)
			}
		case 's':
			if strings.HasPrefix(s[i:], "six") {
				set(6)
			}
			if strings.HasPrefix(s[i:], "seven") {
				set(7)
			}
		case 'e':
			if strings.HasPrefix(s[i:], "eight") {
				set(8)
			}
		case 'n':
			if strings.HasPrefix(s[i:], "nine") {
				set(9)
			}
		}
	}
	return first*10 + last
}
