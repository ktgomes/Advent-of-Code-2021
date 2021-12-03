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

	// Construct list of original lines so that it can be re-used
	var originalLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		originalLines = append(originalLines, scanner.Text())
	}

	// Used to keep track of updated length of remaining lines
	oxygenLines := originalLines

	// Build which lines should be kept after each columns is evaluated
	var remainingLines []string

	for i := 0; i < columns; i++ {
		count := 0
		for _, line := range oxygenLines {
			if line[i] == '0' {
				count -= 1
			} else {
				count += 1
			}
		}

		for _, line := range oxygenLines {
			// If 1 is most common or equal, keep columns with 1s
			if count >= 0 {
				if line[i] == '1' {
					remainingLines = append(remainingLines, line)
				}
			} else { // If 0 is most common, keep columns with 0s
				if line[i] == '0' {
					remainingLines = append(remainingLines, line)
				}
			}
		}

		oxygenLines = remainingLines
		if len(remainingLines) == 1 {
			break
		}
		remainingLines = nil
	}

	oxygen, _ := strconv.ParseInt(string(oxygenLines[0]), 2, 64)
	fmt.Println(oxygen)

	co2Lines := originalLines
	remainingLines = nil

	for i := 0; i < columns; i++ {
		count := 0
		for _, line := range co2Lines {
			if line[i] == '0' {
				count -= 1
			} else {
				count += 1
			}
		}

		for _, line := range co2Lines {
			// If 0 is least common or equal, keep columns with 0s
			if count >= 0 {
				if line[i] == '0' {
					remainingLines = append(remainingLines, line)
				}
			} else { // If 1 is least common, keep columns with 1s
				if line[i] == '1' {
					remainingLines = append(remainingLines, line)
				}
			}
		}

		co2Lines = remainingLines
		if len(co2Lines) == 1 {
			break
		}
		remainingLines = nil
	}

	co2, _ := strconv.ParseInt(string(co2Lines[0]), 2, 64)
	fmt.Println(co2)

	fmt.Println(oxygen * co2)

}
