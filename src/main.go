package main

import (
	"os"
	"fmt"
)
func main() {

	if len(os.Args) > 1 {
		handleArguments(os.Args)
		return
	}

	loadJSON()
	botStart()
}

func handleArguments(args []string) {
	switch args[1] {
	case "help":
		if len(args) == 2 {
			printMainHelp()
		} else if args[2] == "plugin" {
			printPluginHelp()
		} else if args[2] == "config" {

		}else {
			fmt.Println("Unknown help option. Use 'help' or 'help plugin'")
		}
	case "plugin":
		printHelpPlugin()
	default:
		fmt.Printf("Unknown command: %s\n", args[1])
		fmt.Println("Use 'help' for available commands")
	}
}