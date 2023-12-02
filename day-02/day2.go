package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// parsed data from input file
var allGameData = [][]map[string]int{}

func parseLine(line string) {
	splitColon := strings.Split(line, ": ")
	// title := splitColon[0]
	gameInfo := splitColon[1]

	// gameId, err := strconv.Atoi(strings.Split(title, " ")[1])
	// if err != nil {
	// 	panic(err)
	// }

	splitSemicolon := strings.Split(gameInfo, "; ")

	// list of handfuls of cubes for one game
	gameData := []map[string]int{}

	// one handful of cubes
	for _, handfulText := range splitSemicolon {
		splitComma := strings.Split(handfulText, ", ")

		handfulData := map[string]int{}

		// one color
		for _, colorAndNumber := range splitComma {
			splitSpace := strings.Split(colorAndNumber, " ")
			number, err := strconv.Atoi(splitSpace[0])
			if err != nil {
				panic(err)
			}
			color := splitSpace[1]

			handfulData[color] = number
		}

		gameData = append(gameData, handfulData)
	}

	allGameData = append(allGameData, gameData)
}

// known info for part 1
const (
	maxRedCubes   = 12
	maxGreenCubes = 13
	maxBlueCubes  = 14
)

func part1() {
	var sum = 0
	var gameId = 0
	var gameIsPossible = true

	// loop all games
	for i, gameData := range allGameData {
		gameId = i + 1
		gameIsPossible = true

		// if any handfuls in this game don't work, game is impossible
		for _, handful := range gameData {
			if handful["red"] > maxRedCubes || handful["green"] > maxGreenCubes || handful["blue"] > maxBlueCubes {
				gameIsPossible = false
			}
		}

		// if no handfuls triggered the criteria, this game is possible
		if gameIsPossible {
			sum += gameId
		}
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func part2() {
	var sum = 0
	// var gameId = 0

	var red, green, blue int
	var minRed, minGreen, minBlue int

	// loop all games
	for _, gameData := range allGameData {
		// gameId = i + 1
		minRed = 0
		minGreen = 0
		minBlue = 0

		// find min value of each color for this game
		for _, handful := range gameData {
			red = handful["red"]
			if red > minRed {
				minRed = red
			}
			green = handful["green"]
			if green > minGreen {
				minGreen = green
			}
			blue = handful["blue"]
			if blue > minBlue {
				minBlue = blue
			}
		}

		// find power of set for game and add it to sum
		sum += minRed * minGreen * minBlue
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	str := string(buf)
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		parseLine(line)
	}

	part1()
	part2()
}
