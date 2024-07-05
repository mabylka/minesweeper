package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func optionshandler() {
	levelPtr := flag.String("level", "normal", "difficulty level (easy/normal/hard)")
	mapSizePtr := flag.String("mapsize", "10x10", "custom size of the map in format HxW")
	bombsPtr := flag.String("bombs", "", "custom amount of bombs")

	flag.Parse()

	switch *levelPtr {
	case "easy":
		level = 1
	case "normal":
		level = 2
	case "hard":
		level = 3
	default:
		fmt.Print("only easy/normal/hard are available")
		os.Exit(0)
	}
	parseMapSize(*mapSizePtr)
	if *bombsPtr != "" {
		var err error
		bombs, err = strconv.Atoi(*bombsPtr)
		if err != nil || bombs > mapHeight*mapWidth || bombs < 0 {
			fmt.Print("Invalid input for amount of bombs.")
			os.Exit(0)
		}
	} else {
		bombs = level * mapHeight * mapWidth / ((mapHeight + mapWidth) / 2)
	}

	fmt.Println("Level:", *levelPtr)
	fmt.Println("Map Size:", *mapSizePtr)
	fmt.Println("bombs:", bombs)

}

func parseMapSize(sizeStr string) {
	parts := strings.Split(sizeStr, "x")
	if len(parts) != 2 {
		fmt.Println("Invalid map size format. Please use format HxW (e.g., 10x10)")
		os.Exit(0)
	}

	var err error
	mapHeight, err = strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("Invalid map size format. Please use format HxW (e.g., 10x10)")
		os.Exit(0)
	}

	mapWidth, err = strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Invalid map size format. Please use format HxW (e.g., 10x10)")
		os.Exit(0)
	}

	if mapWidth < 5 || mapWidth > 15 || mapHeight < 5 || mapHeight > 15 {
		fmt.Println("Invalid size. min size = 5x5, max = 15x15 )")
		os.Exit(0)
	}
}
