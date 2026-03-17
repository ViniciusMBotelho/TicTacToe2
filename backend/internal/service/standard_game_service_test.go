package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"tictactoe/internal/domain/models"
	"tictactoe/internal/repository"
)

func TestStandardGameService(t *testing.T) {
	repo := repository.NewStandardGameRepository()
	svc := NewStandardGameService(repo)
	ctx := context.Background()

	t.Run("CreateGame", func(t *testing.T) {
		id, err := svc.CreateGame(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, id)

		game, err := svc.GetGameState(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, id, game.ID)
		assert.Equal(t, models.PlayerX, game.CurrentPlayer)
		assert.False(t, game.IsGameOver)
		for _, cell := range game.Board {
			assert.Equal(t, models.Empty, cell)
		}
	})

	t.Run("MakeMove", func(t *testing.T) {
		id, _ := svc.CreateGame(ctx)

		// X moves to center
		game, err := svc.MakeMove(ctx, id, 4)
		assert.NoError(t, err)
		assert.Equal(t, models.PlayerX, game.Board[4])
		assert.Equal(t, models.PlayerO, game.CurrentPlayer)

		// O moves to top-left
		game, err = svc.MakeMove(ctx, id, 0)
		assert.NoError(t, err)
		assert.Equal(t, models.PlayerO, game.Board[0])
		assert.Equal(t, models.PlayerX, game.CurrentPlayer)
	})

	t.Run("InvalidMoves", func(t *testing.T) {
		id, _ := svc.CreateGame(ctx)
		svc.MakeMove(ctx, id, 4) // X moves to 4

		// O tries to move to 4 (already taken)
		_, err := svc.MakeMove(ctx, id, 4)
		assert.Error(t, err)
		assert.ErrorIs(t, err, models.ErrCellAlreadyTaken)

		// O tries to move out of bounds
		_, err = svc.MakeMove(ctx, id, 9)
		assert.Error(t, err)
		assert.ErrorIs(t, err, models.ErrInvalidMove)
	})

	t.Run("WinCondition", func(t *testing.T) {
		id, _ := svc.CreateGame(ctx)
		// X: 0, 1, 2
		// O: 3, 4
		svc.MakeMove(ctx, id, 0)              // X
		svc.MakeMove(ctx, id, 3)              // O
		svc.MakeMove(ctx, id, 1)              // X
		svc.MakeMove(ctx, id, 4)              // O
		game, err := svc.MakeMove(ctx, id, 2) // X wins

		assert.NoError(t, err)
		assert.True(t, game.IsGameOver)
		assert.Equal(t, models.PlayerX, game.Winner)
	})

	t.Run("TieCondition", func(t *testing.T) {
		id, _ := svc.CreateGame(ctx)
		// X O X
		// X X O
		// O X O
		moves := []int{0, 1, 2, 5, 3, 6, 4, 8, 7}
		var game *models.StandardGame
		var err error
		for _, m := range moves {
			game, err = svc.MakeMove(ctx, id, m)
		}
		assert.NoError(t, err)
		assert.True(t, game.IsGameOver)
		assert.Equal(t, models.Tie, game.Winner)
	})
}
