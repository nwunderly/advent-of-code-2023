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
