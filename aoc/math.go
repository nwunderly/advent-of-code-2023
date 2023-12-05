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
