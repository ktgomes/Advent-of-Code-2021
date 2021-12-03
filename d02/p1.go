package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	value     int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't read file")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line []string
	var instruct instruction
	horizontal := 0
	depth := 0
	for scanner.Scan() {
		line = strings.Split(scanner.Text(), " ")
		value, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println("Can't read file")
			panic(err)
		}

		instruct = instruction{
			direction: line[0],
			value:     value,
		}

		if instruct.direction == "forward" {
			horizontal += instruct.value
		} else if instruct.direction == "down" {
			depth += instruct.value
		} else if instruct.direction == "up" {
			depth -= instruct.value
		}

	}
	fmt.Println(horizontal * depth)
}
