package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./text/day1.txt")
	if err != nil {
		log.Fatalf("error opening file, %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// hold last landed number
	dial := 50
	// hold amount of 0's we've landed on
	password := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		input, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("couldn't convert to number", err)
		}

		switch direction {
		case "L":
			dial, password = left(dial, password, input)
		case "R":
			dial, password = right(dial, password, input)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Println(password)
}

func left(dial int, password int, input int) (int, int) {
	// Count rotations
	if dial == 0 {
		password += input / 100
	} else if input >= dial {
		password += ((input - dial) / 100) + 1
	}

	// Dial position
	dial = (dial - input) % 100
	if dial < 0 {
		dial += 100
	}

	return dial, password
}

func right(dial int, password int, input int) (int, int) {
	// Count rotations
	password += (input + dial) / 100

	// Dial position
	dial = (dial + input) % 100

	return dial, password
}
