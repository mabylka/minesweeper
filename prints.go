package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	BLUE       = "\033[1;34m%d\033[0m"
	BRIGHTBLUE = "\033[1;36m%d\033[0m"
	YELLOW     = "\033[1;33m%s\033[0m"
	RED        = "\033[1;31m%d\033[0m"
	GREEN      = "\033[0;32m%d\033[0m"
)

func printMap(grid [][]int) {
	printTop()
	verticalCoordinate := 1
	for i := 0; i < len(grid); i++ {
		for k := 0; k < cellHeight; k++ {
			if k == cellHeight/2 {
				if verticalCoordinate < 10 {
					fmt.Printf("%d ", verticalCoordinate)
				} else {
					fmt.Printf("%d", verticalCoordinate)

				}
				verticalCoordinate++
			} else {
				fmt.Print("  ")
			}
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
				fmt.Printf(YELLOW, "*")
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Printf(GREEN, count)
			case 2:
				fmt.Printf(BRIGHTBLUE, count)
			case 3:
				fmt.Printf(BLUE, count)
			default:
				fmt.Printf(RED, count)
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
	fmt.Print("  ")
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
	fmt.Print("   ")
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
