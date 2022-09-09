package minesweeper

import "testing"

func TestNewCell(t *testing.T) {
	x, y := 3, 19
	cell, err := NewCell(CellTypeMine, x, y)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if cell.X != x {
		t.Errorf("expected X to be %d, got %d", x, cell.X)
	}
	if cell.Y != y {
		t.Errorf("expected Y to be %d, got %d", y, cell.Y)
	}
	if cell.AdjMines != 0 {
		t.Errorf("expected AdjMines to be %d, got %d", 0, cell.AdjMines)
	}

}

func TestRevealed(t *testing.T) {
	cell, err := NewCell(CellTypeAdjacent, 0, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if cell.Revealed {
		t.Errorf("expected visited to be false")
	}
	cell.Reveal()
	if !cell.Revealed {
		t.Errorf("expected visited to be true")
	}
}

func TestString(t *testing.T) {
	mineCell, _ := NewCell(CellTypeMine, 0, 0)
	if s := mineCell.String(); s != "*" {
		t.Errorf("expected %q, got %q", "*", s)
	}

	cell, _ := NewCell(CellTypeAdjacent, 0, 0)
	if s := cell.String(); s != " " {
		t.Errorf("expected %q, got %q", " ", s)
	}

	cell.AdjMines++
	if s := cell.String(); s != "1" {
		t.Errorf("expected %q, got %q", "1", s)
	}

	cell.AdjMines++
	if s := cell.String(); s != "2" {
		t.Errorf("expected %q, got %q", "2", s)
	}
}
