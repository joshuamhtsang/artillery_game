package gameboard

type boardSquare struct {
	team         string
	hasArtillery bool
}

func InitGameboard(xSize int, ySize int) [][]boardSquare {
	//var gameboard [xSize][ySize]boardSquare

	gameboard := make([][]boardSquare, ySize)
	for i := range gameboard {
		gameboard[i] = make([]boardSquare, xSize)
	}
	return gameboard
}
