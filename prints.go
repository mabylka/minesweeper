package main

import "fmt"

func printMap(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}
