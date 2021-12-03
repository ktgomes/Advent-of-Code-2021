package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	const columns int = 5

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Can't read file")
		panic(err)
	}
	defer file.Close()

	var originalLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		originalLines = append(originalLines, scanner.Text())
	}

	oxygenLines := originalLines
	fmt.Println(oxygenLines)

	for i := 0; i < columns; i++ {
		count := 0
		for _, line := range oxygenLines {
			if line[i] == '0' {
				count -= 1
			} else {
				count += 1
			}
		}

		var markForDeletion []int

		for k, line := range oxygenLines {
			// If 1 is most common, remove columns with 0s
			if count >= 0 {
				if line[i] == '0' {
					markForDeletion = append(markForDeletion, k)
				}
			} else { // If 0 is most common, remove columns with 1s
				if line[i] == '1' {
					markForDeletion = append(markForDeletion, k)
				}
			}
		}

		for k, value := range markForDeletion {
			oxygenLines = append(oxygenLines[:value-k], oxygenLines[value+1-k:]...)
		}
	}

	oxygen := oxygenLines[0]
	fmt.Println(oxygen)

	co2Lines := originalLines
	fmt.Println(co2Lines)

	for i := 0; i < columns; i++ {
		count := 0
		for _, line := range co2Lines {
			if line[i] == '0' {
				count -= 1
			} else {
				count += 1
			}
		}

		var markForDeletion []int

		for k, line := range co2Lines {
			// If 1 is least common, remove columns with 0s
			if count <= 0 {
				if line[i] == '0' {
					markForDeletion = append(markForDeletion, k)
				}
			} else { // If 0 is least common, remove columns with 1s
				if line[i] == '1' {
					markForDeletion = append(markForDeletion, k)
				}
			}
		}

		for k, value := range markForDeletion {
			co2Lines = append(co2Lines[:value-k], co2Lines[value+1-k:]...)
		}
	}

	co2 := co2Lines[0]
	fmt.Println(co2)

}
