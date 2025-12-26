package day1

import (
	"bufio"
	"log"
	"os"
)

func Clock() {
	// range: 0 -> 99
	// movement: L / R
	// start pos
	// end pos
	// check if position is at 0
	// times that the counter is at 0
}

func ReadFile() []string {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}