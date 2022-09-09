package minesweeper

import "testing"

func TestNewGame(t *testing.T) {
	game := NewGame(Intermediate)
	if game.Status != GameStatusInProgress {
		t.Errorf("expected game to be in progress, got %q", game.Status)
	}
}
