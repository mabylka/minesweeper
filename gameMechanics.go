package main

import (
	"fmt"
	"math/rand"
)

func mapGeneration(firstX, firstY int) {
	for i := 0; i < bombs; {
		rndH := rand.Intn(mapHeight)
		rndW := rand.Intn(mapWidth)
		if mapGrid[rndH][rndW] == 0 && rndH != firstX && rndW != firstY {
			mapGrid[rndH][rndW] = 9
			i++
		}
	}

	for i := 0; i < mapHeight; i++ {
		for j := 0; j < mapWidth; j++ {
			countNeighbours(i, j)
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
	if !firstdone {
		mapGeneration(x-1, y-1)
		firstdone = true
		moves = 1
	}
	if mapGrid[x-1][y-1] == 9 {
		mapGrid[x-1][y-1] += 10
		gameisover = true
	} else if mapGrid[x-1][y-1] == 0 {
		removeIsland(x-1, y-1)
		moves++
	} else if mapGrid[x-1][y-1] < 9 {
		mapGrid[x-1][y-1] += 10
		moves++
	}
	movesleft := false
	for i := 0; i < mapHeight; i++ {
		for j := 0; j < mapWidth; j++ {
			if mapGrid[i][j] < 9 {
				movesleft = true
				break
			}
		}
	}
	if !movesleft {
		gameisover = true
		gameWin = true
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

func countNeighbours(x, y int) {
	directions := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {-1, 1}, {1, 1}, {0, 1}, {1, 0}}

	for _, dir := range directions {
		newx := x + dir[0]
		newy := y + dir[1]
		if newx >= 0 && newx < mapHeight && newy >= 0 && newy < mapWidth && mapGrid[newx][newy] == 9 && mapGrid[x][y] < 9 {
			mapGrid[x][y] += 1
		}
	}
}
