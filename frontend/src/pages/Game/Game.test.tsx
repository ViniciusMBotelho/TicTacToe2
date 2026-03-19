import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import Game from '../../pages/Game/Game';

// Mock the components used in Game
vi.mock('../../components/GameBoard/GameBoard', () => ({
  default: ({ squares, onSquareClick, scores }: any) => (
    <div data-testid="game-board">
      <div data-testid="scores">
        X: {scores.playerX}, O: {scores.playerO}, Ties: {scores.ties}
      </div>
      <div data-testid="squares">
        {squares.map((square: any, index: number) => (
          <button
            key={index}
            data-testid={`square-${index}`}
            onClick={() => onSquareClick(index)}
          >
            {square || ''}
          </button>
        ))}
      </div>
    </div>
  ),
}));

vi.mock('../../components/Title/Title', () => ({
  default: () => <div data-testid="title">Title</div>,
}));

vi.mock('../../components/FrameBorder/FrameBorder', () => ({
  default: () => <div data-testid="frame-border">FrameBorder</div>,
}));

describe('Game Component', () => {
  it('should render game board with initial state', () => {
    render(<Game />);
    expect(screen.getByTestId('game-board')).toBeTruthy();
  });

  it('should render title and frame border', () => {
    render(<Game />);
    expect(screen.getByTestId('title')).toBeTruthy();
    expect(screen.getByTestId('frame-border')).toBeTruthy();
  });

  it('should have initial scores at 0', () => {
    render(<Game />);
    expect(screen.getByTestId('scores')).toHaveTextContent('X: 0, O: 0, Ties: 0');
  });

  it('should update board when square is clicked', () => {
    render(<Game />);
    const square0 = screen.getByTestId('square-0');
    fireEvent.click(square0);
    expect(square0.textContent).toBe('X');
  });

  it('should alternate players on each move', () => {
    render(<Game />);
    fireEvent.click(screen.getByTestId('square-0'));
    fireEvent.click(screen.getByTestId('square-1'));
    fireEvent.click(screen.getByTestId('square-2'));

    expect(screen.getByTestId('square-0').textContent).toBe('X');
    expect(screen.getByTestId('square-1').textContent).toBe('O');
    expect(screen.getByTestId('square-2').textContent).toBe('X');
  });

  it('should not allow clicking on occupied square', () => {
    render(<Game />);
    const square0 = screen.getByTestId('square-0');
    fireEvent.click(square0);
    expect(square0.textContent).toBe('X');

    fireEvent.click(square0);
    expect(square0.textContent).toBe('X');
  });

  it('should not allow moves after game is won', () => {
    render(<Game />);
    // Winning sequence for X: 0, 4, 8
    fireEvent.click(screen.getByTestId('square-0')); // X at 0
    fireEvent.click(screen.getByTestId('square-3')); // O at 3
    fireEvent.click(screen.getByTestId('square-4')); // X at 4
    fireEvent.click(screen.getByTestId('square-5')); // O at 5
    fireEvent.click(screen.getByTestId('square-8')); // X at 8 - X wins

    // Try to click another square (should not change)
    fireEvent.click(screen.getByTestId('square-1'));
    expect(screen.getByTestId('square-1').textContent).toBe('');
  });

  it('should render 9 squares', () => {
    render(<Game />);
    for (let i = 0; i < 9; i++) {
      expect(screen.getByTestId(`square-${i}`)).toBeTruthy();
    }
  });

  it('should initialize with X as current player', () => {
    render(<Game />);
    const square0 = screen.getByTestId('square-0');
    fireEvent.click(square0);
    expect(square0.textContent).toBe('X');
  });
});
