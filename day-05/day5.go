package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/nwunderly/advent-of-code-2023/aoc"
)

type MapRow struct {
	destStart   int
	sourceStart int
	rangeLength int
}

var maps = [][]MapRow{
	{}, // seed to soil
	{}, // soil to fertilizer
	{}, // fertilizer to water
	{}, // water to light
	{}, // light to temperature
	{}, // temperature to humidity
	{}, // humidity to location
}

func parseBlock(i int, lines []string) {
	mp := []MapRow{}
	row := MapRow{}

	for _, line := range lines {
		numbers := aoc.Ints(line)
		row = MapRow{numbers[0], numbers[1], numbers[2]}
		mp = append(mp, row)
		// fmt.Println(i, "|", j, "=", line, "\n      ", row, "\n----------")
	}

	// maps are already defined so we just set the value at the index
	maps[i] = mp
}

func sourceToDestination(mp []MapRow, source int) int {
	diff := 0
	dest := 0

	for _, row := range mp {
		// if source is within source range
		if row.sourceStart <= source && source < row.sourceStart+row.rangeLength {
			diff = source - row.sourceStart
			dest = row.destStart + diff

			// fmt.Printf("%d -> %d\n", source, dest)
			return dest
		}
	}

	// fmt.Printf("returning %d\n", source)
	return source
}

func getChain(seed int) []int {
	chain := []int{seed}
	current := seed

	for _, mp := range maps {
		current = sourceToDestination(mp, current)
		chain = append(chain, current)
	}

	return chain
}

func part1(seeds []int) {
	defer aoc.Timer("Part 1")()

	// seedToSoilMap := maps[0]
	// fmt.Println("SEEDS")

	var chain []int
	// var current int

	lowestLocation := -1

	for _, seed := range seeds {
		// _ = sourceToDestination(seedToSoilMap, seed)
		// chain = []int{seed}
		// current = seed

		// for _, mp := range maps {
		// 	current = sourceToDestination(mp, current)
		// 	chain = append(chain, current)
		// }

		chain = getChain(seed)
		// fmt.Println("chain", chain, len(chain))

		location := chain[len(chain)-1]

		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
			// fmt.Println("new lowest:", lowestLocation)
		}
	}

	fmt.Printf("Part 1: %d\n", lowestLocation)
}

func getRangesFromPairs(ranges []int) ([][]int, [][]int) {
	var start, length int
	// var rng []int
	var startLength []int

	rangesParsed := [][]int{}
	startLengthParsed := [][]int{}

	for i := 0; i < len(ranges)/2; i += 2 {
		start = ranges[i]
		length = ranges[i+1]

		startLength = []int{start, length}
		// rng = aoc.IntRange(start, start+length)

		startLengthParsed = append(startLengthParsed, startLength)
		// rangesParsed = append(rangesParsed, rng)
	}

	return startLengthParsed, rangesParsed
}

func part2(seeds []int) {
	defer aoc.Timer("Part 2")()

	var chain []int
	var location int

	var start, end int

	lowestLocation := -1

	ranges, _ := getRangesFromPairs(seeds)

	for _, rng := range ranges {
		// for _, seed := range rng {
		start = rng[0]
		end = start + rng[1]
		for seed := start; seed < end; seed++ {
			chain = getChain(seed)
			// fmt.Println("chain", chain, len(chain))

			location = chain[len(chain)-1]

			if lowestLocation == -1 || location < lowestLocation {
				lowestLocation = location
				// fmt.Println("new lowest:", lowestLocation)
			}
		}
	}

	fmt.Printf("Part 2: %d\n", lowestLocation)
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	str := string(buf)

	// str = `seeds: 79 14 55 13

	// seed-to-soil map:
	// 50 98 2
	// 52 50 48

	// soil-to-fertilizer map:
	// 0 15 37
	// 37 52 2
	// 39 0 15

	// fertilizer-to-water map:
	// 49 53 8
	// 0 11 42
	// 42 0 7
	// 57 7 4

	// water-to-light map:
	// 88 18 7
	// 18 25 70

	// light-to-temperature map:
	// 45 77 23
	// 81 45 19
	// 68 64 13

	// temperature-to-humidity map:
	// 0 69 1
	// 1 0 69

	// humidity-to-location map:
	// 60 56 37
	// 56 93 4`

	blocks := strings.Split(str, "\n\n")
	firstLine, blocks := blocks[0], blocks[1:]

	seedsStr := strings.Split(firstLine, ": ")[1]
	seeds := aoc.Ints(seedsStr)

	// fmt.Println("seeds", seeds)
	var lines []string

	for i, block := range blocks {
		lines = strings.Split(block, "\n")[1:]
		parseBlock(i, lines)
	}

	part1(seeds)
	part2(seeds)

	fmt.Println("")
	aoc.TimerResults()
}
