package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"menu-go/feature/menu"
	"menu-go/feature/slack"
)

func main() {
	functionPtr := flag.String("function", "default", "function name")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	switch *functionPtr {
	case "menu":
		menu.Menu()
	case "bot":
		slack.Bot()

	default:
		fmt.Println("Invalid function name")
	}
}
