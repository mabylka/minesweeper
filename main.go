package main

import "fmt"

const (
	cellWidth  = 7
	cellHeight = 3
)

var (
	mapHeight  int
	mapWidth   int
	level      int
	bombs      int
	moves      int
	mapGrid    [][]int
	gameisover bool
	gameWin    bool
	firstdone  bool
)

func main() {
	optionshandler()
	mapGrid = make([][]int, mapHeight)
	for i := 0; i < mapHeight; i++ {
		mapGrid[i] = make([]int, mapWidth)
		for j := 0; j < mapWidth; j++ {
			mapGrid[i][j] = 0
		}
	}

	printMap(mapGrid)
	for !gameisover {
		parseInput()
		clearScreen()
		printMap(mapGrid)
	}
	if gameWin {
		fmt.Println("YOUWON")
	} else {
		fmt.Println("GAME OVER")
	}

	fmt.Printf("Moves: %d", moves)
}
