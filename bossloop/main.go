package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	WIDTH := 8
	HEIGHT := 8
	grid := make([][]int, WIDTH)
	for i := 0; i < WIDTH; i++ {
		grid[i] = make([]int, HEIGHT)
	}

	fmt.Println("2d: ", grid)
	generateBoss(WIDTH, HEIGHT, &grid)
	fmt.Println("with boss: ", grid)

	x, y := searchDoubleLoop(&grid)
	fmt.Printf("found boss in X: %v, Y: %v \n", x, y)
}

func generateBoss(WIDTH, HEIGHT int, grid *[][]int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	x := r1.Intn(WIDTH)
	y := r1.Intn(HEIGHT)

	(*grid)[x][y] = 1
}

// Time complexity - Depends on what you define n as, either O(n) or O(n^2)
func searchDoubleLoop(grid *[][]int) (int, int) {
	for x, col := range *grid {
		for y, val := range col {
			if val == 1 {
				return x, y
			}
		}
	}

	return -1, -1
}

// n = size of input
// Input in our search function is the grid, so what is n in this case?
