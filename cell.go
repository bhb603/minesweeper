package minesweeper

import (
	"errors"
	"strconv"
)

type CellType int

const (
	CellTypeMine CellType = iota
	CellTypeAdjacent
)

type Cell struct {
	X        int
	Y        int
	Revealed bool
	Flagged  bool
	AdjMines int
	Type     CellType
}

func NewCell(cellType CellType, x, y int) (*Cell, error) {
	return &Cell{
		X:    x,
		Y:    y,
		Type: cellType,
	}, nil
}

func (c *Cell) Reveal() {
	c.Flagged = false
	c.Revealed = true
}
func (c *Cell) ToggleFlag() error {
	if c.Revealed {
		return errors.New("cannot flag a revealed cell")
	}
	c.Flagged = !c.Flagged
	return nil
}

func (c *Cell) AdjacentCells(grid [][]*Cell) []*Cell {
	adj := []*Cell{}

	deltas := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, delta := range deltas {
		dx, dy := delta[0], delta[1]
		x, y := c.X+dx, c.Y+dy
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			adj = append(adj, grid[x][y])
		}
	}

	return adj
}

func (c Cell) String() string {
	if c.Type == CellTypeMine {
		return "💣"
	}

	return strconv.Itoa(c.AdjMines)
}
