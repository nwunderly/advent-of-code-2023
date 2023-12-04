package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/akamensky/argparse"
)

const templateFileName = "template.go.txt"
const inputFileName = "input.txt"

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func main() {
	// parse
	parser := argparse.NewParser("go run .", "Set up for a new day with a template.")
	day := parser.StringPositional(&argparse.Options{Required: true, Help: "Day number"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// validation
	number, err := strconv.Atoi(*day)
	if err != nil {
		fmt.Printf("Invalid number for day: %s\n", *day)
		return
	}
	if number < 1 || number > 25 {
		fmt.Println("Day must be between 1 and 25 (inclusive).")
		return
	}

	// formatting
	folderName := fmt.Sprintf("day-%02d", number)
	goFileName := fmt.Sprintf("day%d.go", number)
	goFilePath := fmt.Sprintf("%s/%s", folderName, goFileName)
	inputFilePath := fmt.Sprintf("%s/%s", folderName, inputFileName)
	fmt.Printf("%s\n%s\n%s\n\n", folderName, goFilePath, inputFilePath)

	// create folder if necessary
	if !exists(folderName) {
		fmt.Println("Folder does not exist, creating...")
		err := os.Mkdir(folderName, 0777)
		if err != nil {
			fmt.Println("ERROR: Failed to create directory.")
			return
		}
	} else {
		fmt.Println("Folder already exists.")
	}

	// create go file if necessary
	if !exists(goFilePath) {
		fmt.Println("Go file does not exist, creating...")
		input, err := ioutil.ReadFile(templateFileName)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(goFilePath, input, 0777)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Go file already exists.")
	}

	// create input file if necessary
	if !exists(inputFilePath) {
		fmt.Println("Input file does not exist, creating...")
		err = ioutil.WriteFile(inputFilePath, []byte{}, 0777)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Input file already exists.")
	}

	fmt.Println("Done!")
}
