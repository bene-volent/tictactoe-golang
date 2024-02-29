// game.go file
package game

import (
	"fmt"
	"math/rand"
)

// Game represents a Tic-Tac-Toe game.
type Game struct {
	board      [9]int
	playerTurn bool
	nMoves     int
	isOver     bool
	points     int
}

// New creates a new Tic-Tac-Toe game.
func New() Game {
	return Game{}
}

// Score returns the current score of the game.
func (g *Game) Score() int {
	return g.points
}

// Actions returns a list of available moves in the current state.
func (g *Game) Actions() []int {
	if g.isOver {
		return []int{}
	}
	availableActions := make([]int, 0)
	for index, val := range g.board {
		if val == 0 {
			availableActions = append(availableActions, index+1)
		}
	}
	return availableActions
}

// Result updates the game state based on the player's move.
func (g *Game) Result(pos int) {
	if g.board[pos-1] != 0 {
		panic("Position already marked!")
	}

	if g.playerTurn {
		g.board[pos-1] = 1
	} else {
		g.board[pos-1] = -1
	}

	g.nMoves++

	g.isOver, g.points = g.Evaluate()

	g.playerTurn = !g.playerTurn
}

// ResetPos resets the specified position on the board.
func (g *Game) ResetPos(pos int) {
	g.board[pos-1] = 0
	g.nMoves--
	g.isOver, g.points = g.Evaluate()
	g.playerTurn = !g.playerTurn
}

// Evaluate checks the game state and returns if the game is over and the final score.
func (g *Game) Evaluate() (bool, int) {
	gameOver := g.nMoves == 9
	playerScore, oppScore := 0, 0

	// Horizontal
	for row := 0; row < 3; row++ {
		if g.board[row*3] == 1 && g.board[row*3+1] == 1 && g.board[row*3+2] == 1 {
			playerScore++
		}
		if g.board[row*3] == -1 && g.board[row*3+1] == -1 && g.board[row*3+2] == -1 {
			oppScore++
		}
	}

	// Vertical
	for col := 0; col < 3; col++ {
		if g.board[col] == 1 && g.board[col+3] == 1 && g.board[col+6] == 1 {
			playerScore++
		}
		if g.board[col] == -1 && g.board[col+3] == -1 && g.board[col+6] == -1 {
			oppScore++
		}
	}

	// Diagonals
	if g.board[0] == 1 && g.board[4] == 1 && g.board[8] == 1 {
		playerScore++
	}
	if g.board[0] == -1 && g.board[4] == -1 && g.board[8] == -1 {
		oppScore++
	}

	if g.board[2] == 1 && g.board[4] == 1 && g.board[6] == 1 {
		playerScore++
	}
	if g.board[2] == -1 && g.board[4] == -1 && g.board[6] == -1 {
		oppScore++
	}

	gameOver = gameOver || (playerScore != 0 || oppScore != 0)

	if g.playerTurn {
		return gameOver, playerScore
	} else {
		return gameOver, -oppScore
	}
}

// PrintBoard displays the current state of the game board.
func (g *Game) PrintBoard() {
	board := [9]rune{}

	for index := range g.board {
		if g.board[index] == 1 {
			board[index] = 'O'
		} else if g.board[index] == -1 {
			board[index] = 'X'
		} else {
			board[index] = ' '
		}
	}
	fmt.Printf("\n-------------\n")

	fmt.Printf("| %c | %c | %c |\n", board[0], board[1], board[2])
	fmt.Printf("-------------\n")
	fmt.Printf("| %c | %c | %c |\n", board[3], board[4], board[5])
	fmt.Printf("-------------\n")
	fmt.Printf("| %c | %c | %c |\n", board[6], board[7], board[8])
	fmt.Printf("-------------\n")

	fmt.Println("Current Score: ", g.Score())
}

// GameOver checks if the game is over.
func (g *Game) GameOver() bool {
	return g.isOver
}

// RandomMove makes a random move if the game is not over.
func (g *Game) RandomMove() {
	if g.isOver {
		return
	}
	choices := g.Actions()
	choice := rand.Intn(len(choices))
	g.Result(choices[choice])
}

// minimax is a recursive function to implement the minimax algorithm for optimal moves.
func minimax(g *Game, depth int, maximizingPlayer bool) int {
	if g.GameOver() || depth == 0 {
		return g.Score()
	}

	if maximizingPlayer {
		maxEval := -1000

		availableActions := g.Actions()

		for _, pos := range availableActions {
			g.Result(pos)
			eval := minimax(g, depth-1, false)
			g.ResetPos(pos)

			if maxEval < eval {
				maxEval = eval
			}
		}
		return maxEval
	} else {
		minEval := 1000

		availableActions := g.Actions()

		for _, pos := range availableActions {
			g.Result(pos)
			eval := minimax(g, depth-1, true)
			g.ResetPos(pos)

			if minEval > eval {
				minEval = eval
			}
		}

		return minEval
	}
}

// IntelligentMove makes a move based on the minimax algorithm.
func (g *Game) IntelligentMove() {
	availableActions := g.Actions()

	bestScore := -1000
	bestMove := availableActions[0]

	for _, pos := range availableActions {
		g.Result(pos)
		score := minimax(g, len(availableActions)-1, false)
		g.ResetPos(pos)
		if score > bestScore {
			bestScore = score
			bestMove = pos
		}
	}

	g.Result(bestMove)
}

// minimaxAlphaBeta is a recursive function to implement the minimax algorithm with alpha-beta pruning for optimal moves.
func minimaxAlphaBeta(g *Game, depth, alpha, beta int, maximizingPlayer bool) int {
	if g.GameOver() || depth == 0 {
		return g.Score()
	}

	if maximizingPlayer {
		maxEval := -1000

		availableActions := g.Actions()

		for _, pos := range availableActions {
			g.Result(pos)
			eval := minimaxAlphaBeta(g, depth-1, alpha, beta, false)
			g.ResetPos(pos)

			if maxEval < eval {
				maxEval = eval
			}
			alpha = max(alpha, eval)

			if beta <= alpha {
				break // Beta cut-off
			}
		}
		return maxEval
	} else {
		minEval := 1000

		availableActions := g.Actions()

		for _, pos := range availableActions {
			g.Result(pos)
			eval := minimaxAlphaBeta(g, depth-1, alpha, beta, true)
			g.ResetPos(pos)

			if minEval > eval {
				minEval = eval
			}
			beta = min(beta, eval)

			if beta <= alpha {
				break // Alpha cut-off
			}
		}

		return minEval
	}
}

// IntelligentMoveAlphaBeta makes a move based on the minimax algorithm with alpha-beta pruning.
func (g *Game) IntelligentMoveAlphaBeta() {
	availableActions := g.Actions()

	bestScore := -1000
	bestMove := availableActions[0]

	for _, pos := range availableActions {
		g.Result(pos)
		score := minimaxAlphaBeta(g, len(availableActions)-1, -1000, 1000, false)
		g.ResetPos(pos)

		if score > bestScore {
			bestScore = score
			bestMove = pos
		}
	}

	g.Result(bestMove)
}

// IntelligentMoveWithRandomization makes a move based on the minimax algorithm with randomization for some moves.
func (g *Game) IntelligentMoveWithRandomization() {
	availableActions := g.Actions()

	bestScore := -1000
	bestMove := availableActions[0]

	length := len(availableActions)
	if rand.Intn(10) < 2 {
		g.RandomMove()
		return
	}

	for _, pos := range availableActions {
		g.Result(pos)
		score := minimax(g, length-1, false)
		g.ResetPos(pos)
		if score > bestScore {
			bestScore = score
			bestMove = pos
		}
	}

	g.Result(bestMove)
}
