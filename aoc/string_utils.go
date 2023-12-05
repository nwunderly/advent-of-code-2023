package aoc

import "strconv"

func IsDigit(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}

func Int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func Ints(text string) []int {
	numbers := []int{}
	currentNumber := ""
	str := ""
	number := 0

	for _, char := range text {
		str = string(char)
		if IsDigit(str) {
			currentNumber += str
		} else {
			if currentNumber != "" {
				number = Int(currentNumber)
				numbers = append(numbers, number)
				currentNumber = ""
			}
		}
	}

	// if number at end of string
	if currentNumber != "" {
		number = Int(currentNumber)
		numbers = append(numbers, number)
		currentNumber = ""
	}

	return numbers
}
