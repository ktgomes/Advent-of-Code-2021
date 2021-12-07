package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const final_day int = 256

func main() {
	start := time.Now()
	file, _ := os.ReadFile("input.txt")

	data := strings.Split(string(file), ",")

	var fish [9]int
	for _, val := range data {
		index, _ := strconv.Atoi(val)
		fish[index]++
	}

	for day := 1; day <= final_day; day++ {
		fish[7] += fish[0]

		copy(fish[:], append(fish[1:], fish[0]))
		// fmt.Printf("Day %d: %v\n", day, fish)
	}

	result := 0
	for _, v := range fish {
		result += v
	}
	fmt.Println(result)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
