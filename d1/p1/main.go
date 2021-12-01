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

	scanner := bufio.NewScanner(file)
	var old int
	var new int
	first := true
	count := 0
	for scanner.Scan() {
		new, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Can't read file")
			panic(err)
		}

		if !first {
			if new > old {
				count += 1
			}
		} else {
			first = false
		}

		old = new
	}
	fmt.Println(count)
}
