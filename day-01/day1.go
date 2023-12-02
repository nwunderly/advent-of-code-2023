package main

import (
	"bytes"
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

func part1(line string) int {
	var buffer bytes.Buffer

	// first digit
	for _, char := range line {
		if isdigit(string(char)) {
			buffer.WriteRune(char)
			break
		}
	}

	// last digit
	var char rune
	runes := []rune(line)
	for i := len(runes) - 1; i >= 0; i-- {
		char = runes[i]
		if isdigit(string(char)) {
			buffer.WriteRune(char)
			break
		}
	}

	// convert to int
	number, err := strconv.Atoi(buffer.String())
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

var digits_map = map[string]string{
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

	index_digit1 := len(line)
	index_digit2 := -1

	index_left := 0
	index_right := 0

	for _, digit := range digits {
		// find leftmost and rightmost occurrence of digit in string
		index_left = strings.Index(line, digit)
		index_right = strings.LastIndex(line, digit)

		// current digit found further left than current cache
		if index_left != -1 && index_left < index_digit1 {
			index_digit1 = index_left
			digit1 = digit
		}

		// current digit found further right than current cache
		if index_right != -1 && index_right > index_digit2 {
			index_digit2 = index_right
			digit2 = digit
		}
	}

	// "one" -> "1"
	if !isdigit(digit1) {
		digit1 = digits_map[digit1]
	}
	if !isdigit(digit2) {
		digit2 = digits_map[digit2]
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
