package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var all_game_data = [][]map[string]int{}

var (
	red_cubes   = 12
	green_cubes = 13
	blue_cubes  = 14
)

func parse_line(line string) {
	split_colon := strings.Split(line, ": ")
	// title := split_colon[0]
	game_info := split_colon[1]

	// game_id, err := strconv.Atoi(strings.Split(title, " ")[1])
	// if err != nil {
	// 	panic(err)
	// }

	split_semicolon := strings.Split(game_info, "; ")

	// list of handfuls of cubes for one game
	game_data := []map[string]int{}

	// one handful of cubes
	for _, cube_handful := range split_semicolon {
		split_comma := strings.Split(cube_handful, ", ")

		handful_data := map[string]int{}

		// one color
		for _, color_and_number := range split_comma {
			split_space := strings.Split(color_and_number, " ")
			number, err := strconv.Atoi(split_space[0])
			if err != nil {
				panic(err)
			}
			color := split_space[1]

			handful_data[color] = number
		}

		game_data = append(game_data, handful_data)
	}

	all_game_data = append(all_game_data, game_data)
}

func part1() {
	var sum = 0
	var game_id = 0
	var game_valid = true

	// loop all games
	for i, game_data := range all_game_data {
		game_id = i + 1
		game_valid = true

		// if any handfuls in this game don't work, game is invalid
		for _, handful := range game_data {
			if handful["red"] > red_cubes || handful["green"] > green_cubes || handful["blue"] > blue_cubes {
				game_valid = false
			}
		}

		// if no handfuls trigger criteria, game is valid
		if game_valid {
			sum += game_id
		}
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func part2() {
	var sum = 0
	// var game_id = 0

	var red, green, blue int
	var min_red, min_green, min_blue int

	// loop all games
	for _, game_data := range all_game_data {
		// game_id = i + 1
		min_red = 0
		min_green = 0
		min_blue = 0

		// find min value of each color for this game
		for _, handful := range game_data {
			red = handful["red"]
			if red > min_red {
				min_red = red
			}
			green = handful["green"]
			if green > min_green {
				min_green = green
			}
			blue = handful["blue"]
			if blue > min_blue {
				min_blue = blue
			}
		}

		// find power of set for game and add it to sum
		sum += min_red * min_green * min_blue
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
		parse_line(line)
	}

	part1()
	part2()
}
