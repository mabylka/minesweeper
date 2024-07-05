package main

var (
	mapHeight = 10
	mapWidth  = 10
	level     = 1
	mapGrid   [][]int
)

func main() {
	optionshandler()
	mapGeneration()
	printMap(mapGrid)
}
