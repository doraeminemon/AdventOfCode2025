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

type ClockPosition = int

func Run(commmandList ...[]Command) (passcode int) {
	startPos := 50
	var commands []Command
	if len(commmandList) == 0 {
		commands := []Command{}
		strCommands := ReadFile()
		for _, strCommand := range strCommands {
			commands = append(commands, ParseCommand(strCommand))
		}
	} else {
		commands = commmandList[0]
	}
	passcode = 0
	for _, command := range commands {
		nextPos := Turn(startPos, command)
		if nextPos == 0 {
			passcode++
		}
		startPos = nextPos
	}
	return passcode
}

func Turn(prevPos ClockPosition, command Command) ClockPosition {
	if command.direction == Right {
		nextPos := prevPos + command.increment
		if nextPos > 99 {
			return nextPos - 100
		}
		return nextPos
	}
	nextPos := prevPos - command.increment
	if nextPos < 0 {
		return 100 + nextPos
	}
	return nextPos
}

func ParseCommand(command string) (result Command) {
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