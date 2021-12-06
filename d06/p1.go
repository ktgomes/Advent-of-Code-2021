package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lanternfish struct {
	days_until_birth int
}

const final_day int = 80

func main() {
	file, _ := os.ReadFile("input.txt")

	data := strings.Split(string(file), ",")

	var fish []lanternfish
	var new_fish lanternfish
	for _, val := range data {
		new_fish.days_until_birth, _ = strconv.Atoi(val)
		fish = append(fish, new_fish)
	}

	new_fish.days_until_birth = 8

	for day := 1; day <= final_day; day++ {
		for i, _ := range fish {
			if fish[i].the_miracle_of_life() {
				fish = append(fish, new_fish)
			}
		}
		// fmt.Println(fish)
	}

	fmt.Println(len(fish))
}

func (fish *lanternfish) the_miracle_of_life() bool {
	fish.days_until_birth--

	if fish.days_until_birth == -1 {
		fish.days_until_birth = 6
		return true
	}

	return false
}
