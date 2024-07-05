package main

import (
	"fmt"
	"os"
)

func optionshandler() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-help":
			printHelp()
			os.Exit(0)
		default:
			fmt.Println("Usage: go run . -options\nUse \"-help\" option to print available options")
			os.Exit(0)
		}
	}
}

func printHelp() {
	fmt.Print("\nUsage: go run . -options\n\n")
	fmt.Print("-mapsize=HxW: custom sizes of the Height and Width\n\n")
}
