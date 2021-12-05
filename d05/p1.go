package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const graph_len int = 1000

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := [graph_len][graph_len]int{}

	for scanner.Scan() {
		coordinates := strings.Split(scanner.Text(), " -> ")

		var c1 []int
		var c2 []int
		for _, value := range strings.Split(coordinates[0], ",") {
			intValue, _ := strconv.Atoi(value)
			c1 = append(c1, intValue)
		}

		for _, value := range strings.Split(coordinates[1], ",") {
			intValue, _ := strconv.Atoi(value)
			c2 = append(c2, intValue)
		}

		update_graph(&graph, c1, c2)
	}

	count := 0
	for i := 0; i < graph_len; i++ {
		for j := 0; j < graph_len; j++ {
			if graph[i][j] > 1 {
				count++
			}
			// fmt.Print(graph[i][j])
		}
		// fmt.Println()
	}
	fmt.Println(count)

}

func update_graph(graph *[graph_len][graph_len]int, c1 []int, c2 []int) {
	if c1[0] == c2[0] {
		start, end := start_end(c1[1], c2[1])

		for i := start; i <= end; i++ {
			graph[i][c1[0]]++
		}
	} else if c1[1] == c2[1] {
		start, end := start_end(c1[0], c2[0])

		for i := start; i <= end; i++ {
			graph[c1[1]][i]++
		}
	} else if math.Abs(float64(c1[0]-c2[0])) == math.Abs(float64(c1[1]-c2[1])) {
		start := start(c1[1], c2[1])
		down := c2[1] - c1[1]
		over := c2[0] - c1[0]

		x := c1[0]
		for y := start; y != c2[1]; {
			if down > 0 {
				y++
			} else {
				y--
			}
			graph[y][x]++

			if over > 0 {
				x++
			} else {
				x--
			}
		}

	}
}

func start_end(val1 int, val2 int) (int, int) {
	start := 0
	end := 0
	if val1 <= val2 {
		start = val1
		end = val2
	} else {
		start = val2
		end = val1
	}
	return start, end
}

func start(val1 int, val2 int) int {
	if val1 < val2 {
		return val1 - 1
	}
	return val1 + 1
}
