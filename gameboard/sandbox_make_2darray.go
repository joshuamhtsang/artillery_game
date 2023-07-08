package main

import (
	"fmt"
)

type boardSquare struct {
	team         string
	hasArtillery bool
}

func main() {
	fmt.Println("Attempt to create 2D game board.")

	var gameBoard [8][8]boardSquare
	fmt.Println(gameBoard)

	gameBoard[0][0].team = "blue"
	gameBoard[0][0].hasArtillery = false

	fmt.Println(gameBoard[0][0])
}
