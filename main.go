package main

import (
	"flag"
	"fmt"
	"menu-go/feature/menu"
	"menu-go/feature/slack"
)

func main() {
	functionPtr := flag.String("function", "default", "function name")
	flag.Parse()

	switch *functionPtr {
	case "menu":
		menu.Menu()
	case "bot":
		slack.Bot()

	default:
		fmt.Println("Invalid function name")
	}
}
