package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Direction bool

const (
	Left Direction = true
	Right Direction = false
)

func Clock() {
	// range: 0 -> 99
	// movement: L / R
	// start pos
	// end pos
	// check if position is at 0
	// times that the counter is at 0
}

func ReadCommand(command string) (direction Direction, increment int) {
	dirStr := command[0:1]
	switch dirStr {
		case "L":
			direction = Left
		case "R":
			direction = Right
		default:
			log.Fatal("Unknown command")
	}
	incrementStr := command[1:]
	increment, err := strconv.Atoi(incrementStr)
	if err != nil {
		log.Fatal("Cannot parse string to int")
	}
	return
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