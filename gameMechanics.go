package main

import "math/rand"

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
