package ports

import (
	"context"
	"tictactoe/internal/domain/models"
)

// GameRepository defines the persistence operations for Tic-Tac-Toe-2 games.
type GameRepository interface {
	Save(ctx context.Context, game *models.Game) error
	FindByID(ctx context.Context, id string) (*models.Game, error)
}

// StandardGameRepository defines the persistence operations for standard games.
type StandardGameRepository interface {
	Save(ctx context.Context, game *models.StandardGame) error
	FindByID(ctx context.Context, id string) (*models.StandardGame, error)
}
