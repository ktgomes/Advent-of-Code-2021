package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't read file")
		panic(err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Can't read file")
			panic(err)
		}
		lines = append(lines, num)
	}

	a := make([]int, len(lines)-3)

	var old int
	var new int
	count := 0
	for i := range a {
		old = lines[i] + lines[i+1] + lines[i+2]
		new = lines[i+1] + lines[i+2] + lines[i+3]

		if new > old {
			count += 1
		}
	}
	fmt.Println(count)
}
