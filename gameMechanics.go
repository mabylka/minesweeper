package main

import (
	"fmt"
	"math/rand"
)

func mapGeneration() {
	mapGrid = make([][]int, mapHeight)
	for i := 0; i < mapHeight; i++ {
		mapGrid[i] = make([]int, mapWidth)
		for j := 0; j < mapWidth; j++ {
			mapGrid[i][j] = 0
		}
	}

	for i := 0; i < bombs; {
		rndH := rand.Intn(mapHeight)
		rndW := rand.Intn(mapWidth)
		if mapGrid[rndH][rndW] == 0 {
			mapGrid[rndH][rndW] = 9
			i++
		}
	}

	directions := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {-1, 1}, {1, 1}, {0, 1}, {1, 0}}
	for i := 0; i < mapHeight; i++ {
		for j := 0; j < mapWidth; j++ {
			for _, dir := range directions {
				newx := i + dir[0]
				newy := j + dir[1]
				if newx >= 0 && newx < mapHeight && newy >= 0 && newy < mapWidth && mapGrid[newx][newy] == 9 && mapGrid[i][j] < 9 {
					mapGrid[i][j] += 1
				}

			}
		}
	}

}

func parseInput() {
	var x, y int

	validInput := false
	for !validInput {
		fmt.Println("Enter your coordinates:")
		_, err := fmt.Scanf("%d %d", &x, &y)
		if err != nil {
			fmt.Println("Invalid Input")
		}
		if x > 0 && x <= mapHeight && y > 0 && y <= mapWidth {
			validInput = true
		}
	}
	if mapGrid[x-1][y-1] == 9 {
		mapGrid[x-1][y-1] += 10
		gameisover = true
	} else if mapGrid[x-1][y-1] == 0 {
		removeIsland(x-1, y-1)
	} else if mapGrid[x-1][y-1] < 9 {
		mapGrid[x-1][y-1] += 10
	}
}

func removeIsland(x, y int) {

	directions := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {-1, 1}, {1, 1}, {0, 1}, {1, 0}}
	mapGrid[x][y] += 10
	for _, dir := range directions {
		newx := x + dir[0]
		newy := y + dir[1]
		if newx >= 0 && newx < mapHeight && newy >= 0 && newy < mapWidth && mapGrid[newx][newy] < 9 {

			if mapGrid[newx][newy] == 0 {
				removeIsland(newx, newy)
			} else {
				mapGrid[newx][newy] += 10
			}
		}
	}
}
