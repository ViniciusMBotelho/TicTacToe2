package repository

import (
	"context"
	"sync"

	"tictactoe/internal/domain/models"
	"tictactoe/internal/ports"
)

type gameRepo struct {
	mu    sync.RWMutex
	games map[string]*models.Game
}

// NewGameRepository creates a new thread-safe in-memory Tic-Tac-Toe-2 repository.
func NewGameRepository() ports.GameRepository {
	return &gameRepo{
		games: make(map[string]*models.Game),
	}
}

func (r *gameRepo) Save(ctx context.Context, game *models.Game) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.games[game.ID] = game
	return nil
}

func (r *gameRepo) FindByID(ctx context.Context, id string) (*models.Game, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	game, ok := r.games[id]
	if !ok {
		return nil, models.ErrGameNotFound
	}
	return game, nil
}

type standardGameRepo struct {
	mu    sync.RWMutex
	games map[string]*models.StandardGame
}

// NewStandardGameRepository creates a new thread-safe in-memory standard game repository.
func NewStandardGameRepository() ports.StandardGameRepository {
	return &standardGameRepo{
		games: make(map[string]*models.StandardGame),
	}
}

func (r *standardGameRepo) Save(ctx context.Context, game *models.StandardGame) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.games[game.ID] = game
	return nil
}

func (r *standardGameRepo) FindByID(ctx context.Context, id string) (*models.StandardGame, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	game, ok := r.games[id]
	if !ok {
		return nil, models.ErrGameNotFound
	}
	return game, nil
}
