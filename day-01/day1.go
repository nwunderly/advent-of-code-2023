package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func isdigit(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}

// original solution to part 1
func part1Original(line string) int {
	var str, digit1, digit2 string

	// first digit
	for _, char := range line {
		str = string(char)
		if isdigit(str) {
			digit1 = str
			break
		}
	}

	// last digit
	var char rune
	runes := []rune(line)
	for i := len(runes) - 1; i >= 0; i-- {
		char = runes[i]
		str = string(char)
		if isdigit(str) {
			digit2 = str
			break
		}
	}

	// convert to int
	number, err := strconv.Atoi(digit1 + digit2)
	if err != nil {
		panic(err)
	}

	return number
}

// optimized solution to part 1
func part1(line string) int {
	var str, digit1, digit2 string

	for _, char := range line {
		str = string(char)
		if isdigit(str) {
			if digit1 == "" {
				digit1 = str
			}
			digit2 = str
		}
	}

	// convert to int
	number, err := strconv.Atoi(digit1 + digit2)
	if err != nil {
		panic(err)
	}

	return number
}

var digits = []string{
	"0", "zero",
	"1", "one",
	"2", "two",
	"3", "three",
	"4", "four",
	"5", "five",
	"6", "six",
	"7", "seven",
	"8", "eight",
	"9", "nine",
}

var digitsMap = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func part2(line string) int {
	digit1 := ""
	digit2 := ""

	indexDigit1 := len(line)
	indexDigit2 := -1

	indexLeft := 0
	indexRight := 0

	for _, digit := range digits {
		// find leftmost and rightmost occurrence of digit in string
		indexLeft = strings.Index(line, digit)
		indexRight = strings.LastIndex(line, digit)

		// current digit found further left than current cache
		if indexLeft != -1 && indexLeft < indexDigit1 {
			indexDigit1 = indexLeft
			digit1 = digit
		}

		// current digit found further right than current cache
		if indexRight != -1 && indexRight > indexDigit2 {
			indexDigit2 = indexRight
			digit2 = digit
		}
	}

	// "one" -> "1"
	if !isdigit(digit1) {
		digit1 = digitsMap[digit1]
	}
	if !isdigit(digit2) {
		digit2 = digitsMap[digit2]
	}

	// make final number
	number, err := strconv.Atoi(digit1 + digit2)
	if err != nil {
		panic(err)
	}

	return number
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	str := string(buf)
	lines := strings.Split(str, "\n")

	// Part 1
	sum := 0
	for _, line := range lines {
		sum += part1(line)
	}
	fmt.Printf("Part 1: %d\n", sum)

	// Part 2
	sum = 0
	for _, line := range lines {
		sum += part2(line)
	}
	fmt.Printf("Part 2: %d\n", sum)
}
