package main

type lanternfish struct {
	days_until_birth int
	Child            *lanternfish
}

const final_day int = 80

// func main() {
// 	file, _ := os.ReadFile("test.txt")

// 	data := strings.Split(string(file), ",")

// 	var old_fish lanternfish
// 	// var new_fish lanternfish
// 	for _, val := range data {
// 		days, _ := strconv.Atoi(val)
// 		new_fish := lanternfish{days_until_birth: days, Child: &old_fish}
// 		old_fish := new_fish
// 		fmt.Println(old_fish)
// 	}

// 	fmt.Println(old_fish.Child.days_until_birth)
// 	// new_fish.days_until_birth = 8

// 	// for day := 1; day <= final_day; day++ {
// 	// 	for i, _ := range fish {
// 	// 		if fish[i].the_miracle_of_life() {
// 	// 			// fish = append(fish, new_fish)
// 	// 		}
// 	// 	}
// 	// 	// fmt.Println(fish)
// 	// }

// }
func main() {

	// get list of number of fish for each day, shift list over for each iteration til day is reached
}
