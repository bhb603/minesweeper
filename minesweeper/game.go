package minesweeper

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	GameStatusLost       = "lost"
	GameStatusWon        = "won"
	GameStatusInProgress = "in_progress"
)

var (
	Beginner     GameConfig = GameConfig{8, 8, 10}
	Intermediate GameConfig = GameConfig{16, 16, 40}
	Expert       GameConfig = GameConfig{24, 24, 99}
)

type Game struct {
	ID         string `json:"id"`
	Grid       [][]*Cell
	Status     string
	NumFlagged int
	GameConfig
}

type GameConfig struct {
	Height   int
	Width    int
	NumMines int
}

func NewGame(config GameConfig) *Game {
	height, width := config.Height, config.Width
	grid := make([][]*Cell, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]*Cell, width)
		for j := 0; j < width; j++ {
			grid[i][j], _ = NewCell(CellTypeAdjacent, i, j)
		}
	}

	id := uuid.New().String()
	game := &Game{
		ID:         id,
		Grid:       grid,
		Status:     GameStatusInProgress,
		GameConfig: config,
	}
	game.seedMines()
	return game
}

func (g *Game) seedMines() {
	rand.Seed(time.Now().UnixNano())
	height, width := g.Height, g.Width

	mines := make([]int, height*width)
	for n := 0; n < g.NumMines; n++ {
		mines[n] = 1
	}

	rand.Shuffle(len(mines), func(i, j int) { mines[i], mines[j] = mines[j], mines[i] })

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			cellNumber := i*height + j
			if mines[cellNumber] == 1 {
				mineCell := g.Grid[i][j]
				mineCell.Type = CellTypeMine
				for _, adj := range mineCell.AdjacentCells(g.Grid) {
					adj.AdjMines++
				}
			}
		}
	}
}

func (g *Game) GetCell(x, y int) (*Cell, error) {
	if err := g.validateCoords(x, y); err != nil {
		return nil, err
	}

	return g.Grid[x][y], nil
}

func (g *Game) ToggleFlag(x, y int) ([]*Cell, error) {
	if err := g.validateCoords(x, y); err != nil {
		return []*Cell{}, err
	}

	cell := g.Grid[x][y]
	err := cell.ToggleFlag()

	g.refreshState()

	return []*Cell{cell}, err
}

func (g *Game) RevealCell(x, y int) ([]*Cell, error) {
	if err := g.validateCoords(x, y); err != nil {
		return []*Cell{}, err
	}
	cell := g.Grid[x][y]

	// Mine: Game Over
	if cell.Type == CellTypeMine {
		cell.Reveal()
		g.refreshState()
		return []*Cell{cell}, nil
	}

	// Cell with adj mines
	if cell.AdjMines > 0 {
		cell.Reveal()
		g.refreshState()
		return []*Cell{cell}, nil
	}

	// Blank cell: need to traverse adj cells that can now be revealed
	height, width := g.Height, g.Width
	revealedCells := []*Cell{}
	discovered := make(map[*Cell]bool)
	queue := make(chan *Cell, height*width)
	queue <- cell
	discovered[cell] = true
	for len(queue) > 0 {
		curCell := <-queue
		curCell.Reveal()
		revealedCells = append(revealedCells, curCell)
		if curCell.AdjMines == 0 {
			for _, adj := range curCell.AdjacentCells(g.Grid) {
				if !adj.Revealed && !discovered[adj] {
					discovered[adj] = true
					queue <- adj
				}
			}
		}
	}

	g.refreshState()

	return revealedCells, nil
}

func (g *Game) RevealAdj(x, y int) ([]*Cell, error) {
	if err := g.validateCoords(x, y); err != nil {
		return []*Cell{}, err
	}
	cell := g.Grid[x][y]
	if !cell.Revealed {
		return []*Cell{}, errors.New("cannot reveal adjacent cells unless the cell itself is revealed")
	}

	height, width := g.Height, g.Width
	revealedCells := []*Cell{}
	discovered := make(map[*Cell]bool)
	queue := make(chan *Cell, height*width)
	numAdjFlagged := 0
	for _, adj := range cell.AdjacentCells(g.Grid) {
		if adj.Flagged {
			numAdjFlagged++
		}
		if !adj.Revealed && !adj.Flagged {
			queue <- adj
			discovered[adj] = true
		}
	}
	if numAdjFlagged < cell.AdjMines {
		return []*Cell{}, errors.New("cannot reveal adjacent cells unless all adjacent mines have been flagged")
	}

	for len(queue) > 0 {
		curCell := <-queue
		curCell.Reveal()
		revealedCells = append(revealedCells, curCell)
		if curCell.AdjMines == 0 {
			for _, adj := range curCell.AdjacentCells(g.Grid) {
				if !adj.Revealed && !discovered[adj] {
					queue <- adj
					discovered[adj] = true
				}
			}
		}
	}

	g.refreshState()

	return revealedCells, nil
}

func (g *Game) refreshState() {
	totalCells := g.Height * g.Width
	totalFlagged := 0
	totalRevealed := 0
	var lost bool
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			cell := g.Grid[i][j]
			if cell.Revealed {
				totalRevealed++
			} else if cell.Flagged {
				totalFlagged++
			}

			if cell.Revealed && cell.Type == CellTypeMine {
				lost = true
			}
		}
	}

	g.NumFlagged = totalFlagged

	if lost {
		g.Status = GameStatusLost
	} else if totalRevealed+g.NumMines == totalCells {
		g.Status = GameStatusWon
	} else {
		g.Status = GameStatusInProgress
	}
}

func (g *Game) validateCoords(x, y int) error {
	if x < 0 || x >= g.Height || y < 0 || y >= g.Width {
		return errors.New("invalid cell")
	}

	return nil
}
