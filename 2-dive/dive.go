package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.data")
	check(err)

//	partOne(file)
	partTwo(file)

	defer file.Close()

}

func partTwo(file *os.File) {

	scanner := bufio.NewScanner(file)

	var horizontal, depth, aim int
	for scanner.Scan() {
		text := scanner.Text()
		operation := string(text[0])
		amount := int(text[len(text)-1] - '0')
		if operation == "f" {
			horizontal += amount
			depth += aim * amount
		} else if operation == "d" {
			aim += amount
		} else {
			aim -= amount
		}
	}

	fmt.Printf("Second part --Horizontal: %v -- Depth: %v -- Product %v", horizontal, depth, horizontal*depth)
}

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	var horizontal, depth int
	for scanner.Scan() {
		text := scanner.Text()
		operation := string(text[0])
		amount := int(text[len(text)-1] - '0')

		if operation == "f" {
			horizontal += amount
		} else if operation == "d" {
			depth += amount
		} else {
			depth -= amount
		}
	}

	fmt.Printf("Horizontal: %v -- Depth: %v -- Product %v\n", horizontal, depth, horizontal*depth)

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
