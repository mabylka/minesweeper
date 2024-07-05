package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func printMap(grid [][]int) {
	printTop()
	for i := 0; i < len(grid); i++ {
		for k := 0; k < cellHeight; k++ {
			for j := 0; j < len(grid[i]); j++ {
				fmt.Print("|")
				if grid[i][j] < 10 {
					if gameisover && grid[i][j] == 9 {
						filOpenedCell(k, grid[i][j])
					} else {
						filClosedSell()
					}
				} else {
					filOpenedCell(k, grid[i][j]-10)
				}
			}
			fmt.Print("|")
			fmt.Println()
		}
	}

}

func filClosedSell() {
	for i := 0; i < cellWidth; i++ {
		fmt.Print("X")
	}
}

func filOpenedCell(k, count int) {
	for i := 0; i < cellWidth; i++ {
		if k == cellHeight/2 && count != 0 && i == cellWidth/2 {
			switch count {
			case 9:
				fmt.Print("*")
			case 0:
				fmt.Print(" ")
			default:
				fmt.Print(count)
			}
		} else if k == cellHeight-1 {
			fmt.Print("_")
		} else {
			fmt.Print(" ")
		}
	}
}

func printTop() {
	fmt.Println()
	horizontalCoordinate := 1
	for i := 0; i < mapWidth*(cellWidth+1); i++ {
		if horizontalCoordinate < 10 {
			if i%(cellWidth+1) == 4 {
				fmt.Print(horizontalCoordinate)
				horizontalCoordinate++
			} else {
				fmt.Print(" ")
			}
		} else {
			if i%(cellWidth) == 5 {
				fmt.Print(horizontalCoordinate)
				horizontalCoordinate++
			} else {
				fmt.Print(" ")
			}
		}
		if horizontalCoordinate > mapWidth {
			break
		}
	}

	fmt.Println()
	fmt.Print(" ")
	for i := 0; i < cellWidth*mapWidth+mapWidth-1; i++ {
		fmt.Print("_")
	}
	fmt.Println()
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
