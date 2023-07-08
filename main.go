package main

import (
	"fmt"

	"github.com/joshuamhtsang/artillery_game/gameboard"
)

func main() {
	fmt.Println("Welcome to Artillery Game!")

	// Initiate the gameboard
	gb1 := gameboard.InitGameboard(8, 8)
	fmt.Println(gb1)
}
