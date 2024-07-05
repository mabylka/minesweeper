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
	var x, y int
	for !gameisover {
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
		if mapGrid[x-1][y-1] != 9 {
			mapGrid[x-1][y-1] += 10
		} else {
			mapGrid[x-1][y-1] += 10
			gameisover = true
		}
		printMap(mapGrid)
	}
	fmt.Print("GAME OVER")
}
