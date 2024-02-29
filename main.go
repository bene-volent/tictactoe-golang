package main

import (
	"fmt"
	"tictactoe/game"
)

func main() {
	g := game.New()

	user_input := 0
	for !g.GameOver() {

		g.IntelligentMoveAlphaBeta()
		g.PrintBoard()

		availableActions := g.Actions()
		fmt.Print("Enter the Position to put O (Available Actions: ", availableActions, "): ")
		fmt.Scan(&user_input)

		g.Result(user_input)
		g.PrintBoard()

	}

}
