package models

import "github.com/google/uuid"

// StandardGame represents the state of a traditional 3x3 Tic-Tac-Toe game.
type StandardGame struct {
	ID            string       `json:"id"`
	Board         [9]CellState `json:"board"`
	CurrentPlayer CellState    `json:"current_player"`
	Winner        CellState    `json:"winner"` // Empty, PlayerX, PlayerO, or Tie
	IsGameOver    bool         `json:"is_game_over"`
}

// NewStandardGame initializes a new 3x3 game.
func NewStandardGame() *StandardGame {
	return &StandardGame{
		ID:            uuid.New().String(),
		Board:         [9]CellState{},
		CurrentPlayer: PlayerX,
		IsGameOver:    false,
		Winner:        Empty,
	}
}
