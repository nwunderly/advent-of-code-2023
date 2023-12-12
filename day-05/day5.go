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

func getRangesFromPairs(ranges []int) [][]int {
	var start, length int
	var rng []int

	rangesParsed := [][]int{}

	for i := 0; i < len(ranges); i += 2 {
		start = ranges[i]
		length = ranges[i+1]

		// (start, end)
		rng = []int{start, start + length}

		rangesParsed = append(rangesParsed, rng)
	}

	// fmt.Println("getRangesFromPairs input =", ranges)
	// fmt.Println("getRangesFromPairs output =", rangesParsed)
	return rangesParsed
}

func intervalIntersection(mp []MapRow, sourceInterval [][]int) [][]int {
	// note: cannot assume ranges are in numeric order
	destInterval := [][]int{}

	var sourceStart, sourceEnd int
	var mapSourceStart, mapSourceEnd int

	var leftStart, leftEnd int
	var intersectionStart, intersectionEnd int
	var rightStart, rightEnd int
	var left, right []int

	var mapDestStart int

	var destStart, destEnd int

	unmapped := [][]int{}

	for _, row := range mp {
		mapSourceStart = row.sourceStart
		mapSourceEnd = row.sourceStart + row.rangeLength
		mapDestStart = row.destStart

		unmapped = [][]int{}

		for _, sourceRange := range sourceInterval {
			sourceStart = sourceRange[0]
			sourceEnd = sourceRange[1]

			// unmapped source items left of intersection
			leftStart = sourceStart
			leftEnd = aoc.MinInt(sourceEnd, mapSourceStart)
			left = []int{leftStart, leftEnd}

			// our intersection (stuff that's mapped)
			intersectionStart = aoc.MaxInt(sourceStart, mapSourceStart)
			intersectionEnd = aoc.MinInt(mapSourceEnd, sourceEnd)

			// unmapped source items right of intersection
			rightStart = aoc.MaxInt(mapSourceEnd, sourceStart)
			rightEnd = sourceEnd
			right = []int{rightStart, rightEnd}

			// there's stuff to the left
			if leftStart < leftEnd {
				unmapped = append(unmapped, left)
			}

			// the intersection is valid
			if intersectionStart < intersectionEnd {
				destStart = mapDestStart + (intersectionStart - mapSourceStart)
				destEnd = destStart + (intersectionEnd - intersectionStart)
				destInterval = append(destInterval, []int{destStart, destEnd})
			}

			// there's stuff to the right
			if rightStart < rightEnd {
				unmapped = append(unmapped, right)
			}
		}

		// for next iteration (next row of map),
		// source ranges to check should only include currently-unmapped stuff
		sourceInterval = unmapped
	}

	// anything not mapped after checking every row of map is passed on like normal
	destInterval = append(destInterval, unmapped...)

	return destInterval
}

func part2(seeds []int) {
	defer aoc.Timer("Part 2")()

	var currentInterval [][]int

	var location int
	lowestLocation := -1

	seedInterval := getRangesFromPairs(seeds)
	// fmt.Println("seedInterval =", seedInterval)

	currentInterval = seedInterval

	// propagate the interval through the maps
	for _, mp := range maps {
		currentInterval = intervalIntersection(mp, currentInterval)
		// fmt.Println("currentInterval =", currentInterval)
	}

	// find the lowest location value in the mapped interval
	for _, locationRange := range currentInterval {
		// checking the start of each range within the interval
		location = locationRange[0]
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
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
