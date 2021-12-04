package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board struct {
	numbers [5][5]string
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var boards []board
	var newBoard board
	var nums []string
	first := true
	count := 0

	for scanner.Scan() {
		if first {
			nums = strings.Split(scanner.Text(), ",")
			first = false
		} else {
			line := scanner.Text()
			if line != "" {
				copy(newBoard.numbers[count%5][:], strings.Fields(line))

				count += 1
				if count%5 == 0 {
					boards = append(boards, newBoard)
				}
			}
		}
	}

out:
	for _, num := range nums {
		for i, _ := range boards {
			boards[i].update_board(num)
			if boards[i].bingo() {
				fmt.Println(boards[i].winning_value(num))
				break out
			}
		}
	}

}

// method for board structs, will update the board with X where ever the number given appears
func (b *board) update_board(num string) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if (*b).numbers[i][j] == num {
				(*b).numbers[i][j] = "X"
			}
		}
	}
}

func (b board) bingo() bool {
	for i := 0; i < 5; i++ {
		winner_row := true
		winner_column := true
		for j := 0; j < 5; j++ {
			if b.numbers[i][j] != "X" {
				winner_row = false
			}
			if b.numbers[j][i] != "X" {
				winner_column = false
			}
		}
		if winner_row || winner_column {
			return true
		}
	}

	return false
}

func (b board) winning_value(num string) int {
	last_call, _ := strconv.Atoi(num)

	board_sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.numbers[i][j] != "X" {
				board_num, _ := strconv.Atoi(b.numbers[i][j])
				board_sum += board_num
			}

		}
	}

	return board_sum * last_call
}
