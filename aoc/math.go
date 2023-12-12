package aoc

import (
	"fmt"
	"time"
)

var timerCache = []string{}

func Timer(name string) func() {
	start := time.Now()
	return func() {
		msg := fmt.Sprintf("%s took %v\n", name, time.Since(start))
		timerCache = append(timerCache, msg)
	}
}

func TimerResults() {
	for _, msg := range timerCache {
		fmt.Print(msg)
	}
}

func IntRange(start, end int) []int {
	rng := make([]int, end-start)
	for i := range rng {
		rng[i] = start + i
	}
	return rng
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func MinInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
