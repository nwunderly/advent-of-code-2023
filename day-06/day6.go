package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/nwunderly/advent-of-code-2023/aoc"
)

func distanceTraveled(raceTime, buttonTime int) int {
	speed := buttonTime
	travelTime := raceTime - buttonTime
	return speed * travelTime
}

func waysCanBeat(raceTime, distanceToBeat int) int {
	n := 0

	for buttonTime := 0; buttonTime < raceTime; buttonTime++ {
		if distanceTraveled(raceTime, buttonTime) > distanceToBeat {
			n += 1
		}
	}

	return n
}

func part1(times, distances []int) {
	// assumption: times and distances are the same length

	result := 1

	var distanceToBeat int

	for i, raceTime := range times {
		distanceToBeat = distances[i]
		result *= waysCanBeat(raceTime, distanceToBeat)
	}

	fmt.Printf("Part 1: %d\n", result)
}

func compress(arr []int) int {
	result := ""

	for _, number := range arr {
		result += strconv.Itoa(number)
	}

	return aoc.Int(result)
}

func part2(raceTime, distanceToBeat int) {
	defer aoc.Timer("Part 2")()
	result := waysCanBeat(raceTime, distanceToBeat)
	fmt.Printf("Part 2: %d\n", result)
}

func part2Optimized(raceTime, distanceToBeat int) {
	defer aoc.Timer("Part 2 (Optimized)")()

	var min, max int

	// get min
	for buttonTime := 1; buttonTime < raceTime; buttonTime++ {
		if distanceTraveled(raceTime, buttonTime) > distanceToBeat {
			min = buttonTime
			break
		}
	}

	// get max
	for buttonTime := raceTime; buttonTime > 0; buttonTime-- {
		if distanceTraveled(raceTime, buttonTime) > distanceToBeat {
			max = buttonTime
			break
		}
	}

	result := (max + 1) - min

	fmt.Printf("Part 2: %d\n", result)
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	str := string(buf)
	lines := strings.Split(str, "\n")

	times := aoc.Ints(lines[0])
	distances := aoc.Ints(lines[1])

	part1(times, distances)

	time := compress(times)
	distance := compress(distances)

	part2(time, distance)
	part2Optimized(time, distance)

	aoc.TimerResults()
}
