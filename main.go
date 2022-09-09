package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"

	ms "github.com/bhb603/minesweeper/minesweeper"
)

func main() {
	game := ms.NewGame(ms.Beginner)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		sig := <-sigChan
		fmt.Println("got signal:", sig)
		game.Print(true)
		os.Exit(1)
	}()

	reader := bufio.NewReader(os.Stdin)
	game.Print(false)
	for game.Status == ms.GameStatusInProgress {
		prompt()
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		err = runCommand(game, input)
		if err != nil {
			fmt.Println("invalid command")
			continue
		}

		game.Print(false)
	}

	switch game.Status {
	case ms.GameStatusLost:
		fmt.Println("You lost")
	case ms.GameStatusWon:
		fmt.Println("You won!")
	}
}

func prompt() {
	fmt.Println("Make a move")
	fmt.Println("  Reveal a cell:    r {row} {col}")
	fmt.Println("  Flag a cell:      f {row} {col}")
	fmt.Println("  Unflag a cell:    uf {row} {col}")
	fmt.Println("  Reveal adjacent:  ra {row} {col}")
	fmt.Println("  Quit:             q")
	fmt.Print("> ")
}

func runCommand(game *ms.Game, input string) error {
	parts := strings.Split(strings.Trim(input, "\n "), " ")
	if len(parts) < 1 {
		return errors.New("invalid command")
	}
	switch parts[0] {
	case "r":
		return revealCell(game, parts[1:])
	case "f":
		return flagCell(game, parts[1:])
	case "uf":
		return unflagCell(game, parts[1:])
	case "ra":
		return revealAdj(game, parts[1:])
	case "q":
		fmt.Println("good bye")
		os.Exit(0)
	}

	return errors.New("invalid command")
}

func revealCell(game *ms.Game, args []string) error {
	coords, err := parseCoordinates(args)
	if err != nil {
		return err
	}
	_, err = game.RevealCell(coords[0], coords[1])
	return err
}

func flagCell(game *ms.Game, args []string) error {
	coords, err := parseCoordinates(args)
	if err != nil {
		return err
	}
	_, err = game.FlagCell(coords[0], coords[1])
	return err
}

func unflagCell(game *ms.Game, args []string) error {
	coords, err := parseCoordinates(args)
	if err != nil {
		return err
	}
	_, err = game.UnflagCell(coords[0], coords[1])
	return err
}

func revealAdj(game *ms.Game, args []string) error {
	coords, err := parseCoordinates(args)
	if err != nil {
		return err
	}
	_, err = game.RevealAdj(coords[0], coords[1])
	return err
}

func parseCoordinates(args []string) ([2]int, error) {
	coords := [2]int{}
	if len(args) != 2 {
		return coords, errors.New("invalid command")
	}
	for i, arg := range args {
		val, err := strconv.Atoi(arg)
		if err != nil {
			return coords, err
		}
		coords[i] = val
	}

	return coords, nil
}
