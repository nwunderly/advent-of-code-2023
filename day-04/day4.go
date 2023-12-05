package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"

	"github.com/nwunderly/advent-of-code-2023/aoc"
)

type Card struct {
	index          int
	number         int
	winningNumbers []int
	yourNumbers    []int
}

var cards = []Card{}

func parseLine(i int, line string) {
	cardNumber := i + 1

	splitColon := strings.Split(line, ": ")
	numbersStr := strings.TrimSpace(splitColon[1])

	splitVertical := strings.Split(numbersStr, " | ")
	winningNumbers := aoc.Ints(splitVertical[0])
	yourNumbers := aoc.Ints(splitVertical[1])

	cards = append(cards, Card{i, cardNumber, winningNumbers, yourNumbers})
}

func containsInt(array []int, n int) bool {
	for _, i := range array {
		if i == n {
			return true
		}
	}
	return false
}

func countMatches(card Card) int {
	matches := 0

	for _, yourNumber := range card.yourNumbers {
		if containsInt(card.winningNumbers, yourNumber) {
			matches += 1
		}
	}

	return matches
}

func part1() {
	defer aoc.Timer("Part 1")()
	matches := 0
	points := 0
	sum := 0

	for _, card := range cards {
		matches = countMatches(card)
		points = int(math.Pow(2.0, float64(matches-1)))
		// fmt.Printf("card %d: %d points\n", i+1, points)
		sum += points
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func processCard(i int, card Card) int {
	sum := 0
	matches := countMatches(card)
	var extraCard Card

	// fmt.Printf("card %d won %d extra cards\n", card.number, matches)

	// add this card to the count
	sum += 1

	// process extra cards won by this card
	for j := card.index + 1; j < card.index+matches+1; j++ {
		extraCard = cards[j]
		sum += processCard(j, extraCard)
	}

	return sum
}

// original implementation: recursion (see processCard)
func part2Original() {
	defer aoc.Timer("Part 2 (Original)")()
	sum := 0

	for i, card := range cards {
		sum += processCard(i, card)
	}

	fmt.Printf("Part 2: %d\n", sum)
}

// optimized implementation: iterate backwards, cache "value" of each card
func part2Optimized() {
	defer aoc.Timer("Part 2 (Optimized)")()
	sum := 0

	matches := 0
	value := 0

	var card Card
	var cardValues = map[int]int{}

	for i := len(cards) - 1; i >= 0; i-- {
		card = cards[i]
		matches = countMatches(card)
		value = 1

		for j := i + 1; j < i+matches+1; j++ {
			value += cardValues[j]
		}

		// fmt.Printf("card %d: matches=%d, value=%d\n", i+1, matches, value)
		cardValues[i] = value
		sum += value
	}

	fmt.Printf("Part 2 (again): %d\n", sum)
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	str := string(buf)

	// str = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n" +
	// 	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n" +
	// 	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n" +
	// 	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n" +
	// 	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n" +
	// 	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"

	lines := strings.Split(str, "\n")

	for i, line := range lines {
		parseLine(i, line)
	}

	part1()          // <100us
	part2Original()  // <1.5s
	part2Optimized() // <200us

	fmt.Println("")
	aoc.TimerResults()
}
