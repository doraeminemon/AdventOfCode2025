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

type Command struct {
	direction Direction
	increment int
}

type ClockPosition int

func Turn(prevPos ClockPosition, command Command) ClockPosition {
	if command.direction == Right {
		nextPos := prevPos + ClockPosition(command.increment)
		if nextPos > 99 {
			return nextPos - 100
		}
		return nextPos
	}
	nextPos := prevPos - ClockPosition(command.increment)
	if nextPos < 0 {
		return 100 + nextPos
	}
	return nextPos
}

func ReadCommand(command string) (result Command) {
	dirStr := command[0:1]
	switch dirStr {
		case "L":
			result.direction = Left
		case "R":
			result.direction = Right
		default:
			log.Fatal("Unknown command")
	}
	incrementStr := command[1:]
	increment, err := strconv.Atoi(incrementStr)
	if err != nil {
		log.Fatal("Cannot parse string to int")
	}
	result.increment = increment
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