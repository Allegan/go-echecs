package board

import (
	"errors"

	"github.com/Allegan/go-echecs/pkg/color"
	"github.com/Allegan/go-echecs/pkg/piece"
)

// Place represents a square on the chess board
type Place struct {
	Color    color.Color
	Occupied bool
	Piece    *piece.Piece
}

// Board represents the whole board
type Board struct {
	Spaces [8][8]Place
}

// PlaceAt populates a Place with a Piece
func (place *Place) PlaceAt(piece *piece.Piece) error {
	if place.Occupied {
		return errors.New("Cannot place piece because place is occupied")
	}

	place.Occupied = true
	place.Piece = piece

	return nil
}

// New creates a new Board
func (board *Board) New() *Board {
	dim := 8
	placeColor := color.White

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			board.Spaces[i][j] = Place{placeColor, false, nil}
			placeColor = !placeColor
		}
	}

	return board
}
