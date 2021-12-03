package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	const columns int = 12

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't read file")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count [columns]int
	var line string

	for scanner.Scan() {
		line = scanner.Text()

		for i := 0; i < len(line); i++ {
			if line[i] == '0' {
				count[i] -= 1
			} else {
				count[i] += 1
			}
		}
	}

	var gamma []byte
	for i := 0; i < len(count); i++ {
		// >0 means there are more ones
		if count[i] > 0 {
			gamma = append(gamma, '1')
		} else {
			gamma = append(gamma, '0')
		}

	}

	var epsilon []byte
	for i := 0; i < len(count); i++ {
		// <0 means there are more zeros
		if count[i] < 0 {
			epsilon = append(epsilon, '1')
		} else {
			epsilon = append(epsilon, '0')
		}
	}

	gammaVal, err := strconv.ParseInt(string(gamma), 2, 64)
	if err != nil {
		fmt.Println("Error Parsing Binary")
		panic(err)
	}
	fmt.Println(gammaVal)

	epsilonVal, err := strconv.ParseInt(string(epsilon), 2, 64)
	if err != nil {
		fmt.Println("Error Parsing Binary")
		panic(err)
	}
	fmt.Println(epsilonVal)

	fmt.Println(gammaVal * epsilonVal)

}
