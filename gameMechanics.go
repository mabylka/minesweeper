package main

func mapGeneration() {
	mapGrid = make([][]int, mapHeight)
	for i := 0; i < mapHeight; i++ {
		mapGrid[i] = make([]int, mapWidth)
		for j := 0; j < mapWidth; j++ {
			mapGrid[i][j] = 9
		}
	}

}
