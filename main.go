package main

import (
	"fmt"
	"math/rand"
	"time"
)

var grid [][]int
var alive int
var generation int

func makeEmptyGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}

	return grid
}

func printGrid(grid [][]int) {
	for rowIndex, i := range grid {
		for colIndex, _ := range i {
			if grid[rowIndex][colIndex] == 1 {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
	fmt.Println()
}

func printGridRaw(grid [][]int) {
	for rowIndex, i := range grid {
		for colIndex, _ := range i {
			fmt.Print(grid[rowIndex][colIndex])
		}
		fmt.Println()
	}
	fmt.Println()
}

func populateGridRandomly(grid [][]int) [][]int {
	//random := rand.New(rand.NewSource(seed))
	for rowIndex, row := range grid {
		for colIndex, _ := range row {
			cell := rand.Intn(2)
			grid[rowIndex][colIndex] = cell
		}
	}
	return grid
}

func populateNextGrid(grid [][]int) [][]int {
	size := len(grid)
	nextGrid := makeEmptyGrid(size)

	neighborsSum := 0

	// iterate through the previous grid
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			neighborsSum = 0

			// get all the neighbor sum
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					// if it's the current cell, ignore
					if i == 0 && j == 0 {
						continue
					}

					x := (rowIndex - i + size) % size
					y := (colIndex - j + size) % size
					neighborsSum += grid[x][y]
				}
			}

			if cell == 1 && neighborsSum >= 2 && neighborsSum <= 3 {
				nextGrid[rowIndex][colIndex] = 1
			} else if cell == 1 && (neighborsSum < 2 || neighborsSum > 3) {
				nextGrid[rowIndex][colIndex] = 0
			} else if cell == 0 && neighborsSum == 3 {
				nextGrid[rowIndex][colIndex] = 1
			}

		}
	}
	return nextGrid
}

func countAlive(grid [][]int) {
	alive = 0
	for _, i := range grid {
		for _, j := range i {
			alive += j
		}
	}
}

func main() {
	var (
		size int
	)

	fmt.Scan(&size)
	generation = 1

	grid := makeEmptyGrid(size)
	grid = populateGridRandomly(grid)
	countAlive(grid)
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Generation #%d\nAlive: %d\n", generation, alive)
	printGrid(grid)

	nextGrid := makeEmptyGrid(size)
	for i := 1; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		generation++
		nextGrid = populateNextGrid(grid)
		grid = nextGrid
		countAlive(grid)
		fmt.Printf("Generation #%d\nAlive: %d\n", generation, alive)
		printGrid(grid)
	}
}
