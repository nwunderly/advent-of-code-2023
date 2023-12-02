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

func process_line(str string) int {
	var buffer bytes.Buffer

	// first digit
	for _, char := range str {
		if isdigit(string(char)) {
			buffer.WriteRune(char)
			break
		}
	}

	// last digit
	var char rune
	runes := []rune(str)
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

func part1(lines []string) {
	sum := 0

	for _, line := range lines {
		sum += process_line((line))
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func part2(lines []string) {

}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	str := string(buf)
	lines := strings.Split(str, "\n")

	part1(lines)
	part2(lines)
}
