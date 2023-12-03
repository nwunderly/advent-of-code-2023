package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/nwunderly/advent-of-code-2023/aoc"
)

type Rect struct {
	row   int
	col   int
	width int
	text  string
	value int
}

var coordsParts []Rect
var coordsSymbols []Rect

func isSymbol(char rune) bool {
	return !aoc.IsDigit(string(char)) && char != '.'
}

func rectsAdjacent(rectA Rect, rectB Rect) bool {
	// assumption: rectB is a symbol (single point)
	rB := rectB.row
	cB := rectB.col

	// bounds (negative coords shouldn't matter)
	r0 := rectA.row - 1
	c0 := rectA.col - 1
	r1 := rectA.row + 1
	c1 := rectA.col + rectA.width + 1

	// bounds check
	inRowBounds := r0 <= rB && rB <= r1
	inColBounds := c0 <= cB && cB <= c1
	return inRowBounds && inColBounds
}

func parseInput(str string) {
	// parse the input file
	rows := strings.Split(str, "\n")
	var symbol Rect
	var currentNumber string
	var currentPart Rect

	for r, row := range rows {
		for c, col := range row {
			colStr := string(col)
			fmt.Printf("processing row %d col %d, char = '%s'\n", r, c, colStr)

			if aoc.IsDigit(string(colStr)) {
				// new number
				if currentNumber == "" {
					currentPart = Rect{r, c, 0, "", 0}
				}

				// things to update every time we find a new digit
				currentNumber += colStr
				currentPart.width = c - currentPart.col // could also just increment this but would need -1
			} else if currentNumber != "" {
				// done processing number, reset things (skip all this if there's no number to process)
				currentPart.text = currentNumber
				currentPart.value = aoc.Int(currentNumber)
				coordsParts = append(coordsParts, currentPart)
				currentNumber = ""
				// currentPart = Rect{}  // we don't need to zero this, currentNumber indicates a reset
			}

			// this is separate from the above conditionals to visually isolate it from number processing
			if isSymbol(col) {
				// single chars have zero width and aren't numbers
				symbol = Rect{r, c, 0, colStr, 0}
				coordsSymbols = append(coordsSymbols, symbol)
			}
		}
	}
}

func part1() {
	// loop through known part numbers and compare bounds against known symbols
	sum := 0

	for _, part := range coordsParts {
		// check part number coords against all symbols
		for _, symbol := range coordsSymbols {
			if rectsAdjacent(part, symbol) {
				sum += part.value
				break
			}
		}
	}

	fmt.Printf("Part 1 Answer: %d\n", sum)
}

func part2() {
	adjacentCount := 0
	gearRatio := 1
	sum := 0

	// loop through symbols, looking for gears
	for _, symbol := range coordsSymbols {
		adjacentCount = 0
		gearRatio = 1

		// count adjacent parts
		for _, part := range coordsParts {
			if rectsAdjacent(part, symbol) {
				adjacentCount += 1
				gearRatio *= part.value
			}
		}

		// it's a gear
		if symbol.text == "*" && adjacentCount == 2 {
			sum += gearRatio
		}
	}

	fmt.Printf("Part 2 Answer: %d\n", sum)
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	str := string(buf)

	// sample input
	// str = ("467..114..+" +
	// 	"\n...*......" +
	// 	"\n..35..633." +
	// 	"\n......#..." +
	// 	"\n617*......" +
	// 	"\n.....+.58." +
	// 	"\n..592....." +
	// 	"\n......755." +
	// 	"\n...$.*...." +
	// 	"\n.664.598..")

	parseInput(str)
	part1()
	part2()
}
