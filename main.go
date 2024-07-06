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
	mapGrid    [][]int
	gameisover bool
)

func main() {
	optionshandler()
	mapGeneration()
	printMap(mapGrid)
	for !gameisover {
		parseInput()
		clearScreen()
		printMap(mapGrid)
	}
	fmt.Print("GAME OVER")
}
