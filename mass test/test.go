package main

import "fmt"

func main() {
	grid := [][]int{
		{3, 4, 0},
		{2, 5, 6, 0},
	}

	numCars := 3
	carPositions := make([]int, numCars)
	carNames := []string{"v1", "v2", "v3"}
	mirror := 0
	for !allCarsFinished(carPositions, grid, mirror, numCars) {
		for i := 0; i < numCars; i++ {
			if i > len(grid)-1 {
				grid = append(grid, grid[mirror])
				mirror++
			}
			if i > 0 {
				if carPositions[i] < len(grid[i]) {
					fmt.Printf("%s-%d ", carNames[i], grid[i][carPositions[i]])
					carPositions[i]++
				}
			} else {
				if carPositions[i] < len(grid[i])-mirror {
					fmt.Printf("%s-%d ", carNames[i], grid[i][carPositions[i]])
					carPositions[i]++
				}
			}

		}

		fmt.Println()
	}
}

func allCarsFinished(positions []int, grid [][]int, mirror int, numcars int) bool {
	for i, pos := range positions {

		if pos < len(grid[i])-mirror {
			return false
		}
	}
	return true
}
