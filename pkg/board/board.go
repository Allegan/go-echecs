package board

import (
	"errors"
	"fmt"

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

// AddPiece adds a Piece to the board at the location x,y
func (board *Board) AddPiece(piece *piece.Piece, x int, y int) error {
	if board.Spaces[x][y].Occupied {
		return fmt.Errorf("Board[%d][%d] is occupied", x, y)
	}

	board.Spaces[x][y].Occupied = true
	board.Spaces[x][y].Piece = piece

	return nil
}

// Print the board
func (board *Board) Print() {
	dim := 8

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			place := board.Spaces[i][j]

			if place.Color == color.White {
				fmt.Print("W")
			} else {
				fmt.Print("B")
			}
		}

		fmt.Println()
	}
}

func decodePiece(p *piece.Piece) string {
	var str string
	white := []string{
		"♔",
		"♕",
		"♖",
		"♗",
		"♘",
		"♙",
	}
	black := []string{
		"♚",
		"♛",
		"♜",
		"♝",
		"♞",
		"♟︎",
	}

	switch p.Class {
	case piece.King:
		if p.Color == color.White {
			str = white[0]
		} else {
			str = black[0]
		}
	case piece.Queen:
		if p.Color == color.White {
			str = white[1]
		} else {
			str = black[1]
		}
	case piece.Rook:
		if p.Color == color.White {
			str = white[2]
		} else {
			str = black[2]
		}
	case piece.Bishop:
		if p.Color == color.White {
			str = white[3]
		} else {
			str = black[3]
		}
	case piece.Knight:
		if p.Color == color.White {
			str = white[4]
		} else {
			str = black[4]
		}
	case piece.Pawn:
		if p.Color == color.White {
			str = white[5]
		} else {
			str = black[5]
		}
	}

	return str
}

// PrintPieces displays a board layout using unicode pieces
func (board *Board) PrintPieces() {
	dim := 8

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if board.Spaces[i][j].Piece != nil {
				fmt.Print(decodePiece(board.Spaces[i][j].Piece))
			} else {
				fmt.Print("X")
			}
		}

		fmt.Println()
	}
}

// New returns a hardcoded chess board
func New() *Board {
	board := &Board{
		Spaces: [8][8]Place{
			{
				Place{color.White, true, &piece.Piece{Class: piece.Rook, Color: color.Black}},
				Place{color.Black, true, &piece.Piece{Class: piece.Knight, Color: color.Black}},
				Place{color.White, true, &piece.Piece{Class: piece.Bishop, Color: color.Black}},
				Place{color.Black, true, &piece.Piece{Class: piece.King, Color: color.Black}},
				Place{color.White, true, &piece.Piece{Class: piece.Queen, Color: color.Black}},
				Place{color.Black, true, &piece.Piece{Class: piece.Bishop, Color: color.Black}},
				Place{color.White, true, &piece.Piece{Class: piece.Knight, Color: color.Black}},
				Place{color.Black, true, &piece.Piece{Class: piece.Rook, Color: color.Black}},
			},
			{
				Place{color.Black, true, &piece.Piece{Class: piece.Pawn, Color: color.Black}},
				Place{color.White, true, &piece.Piece{Class: piece.Pawn, Color: color.Black}},
				Place{color.Black, true, &piece.Piece{Class: piece.Pawn, Color: color.Black}},
				Place{color.White, true, &piece.Piece{Class: piece.Pawn, Color: color.Black}},
				Place{color.Black, true, &piece.Piece{Class: piece.Pawn, Color: color.Black}},
				Place{color.White, true, &piece.Piece{Class: piece.Pawn, Color: color.Black}},
				Place{color.Black, true, &piece.Piece{Class: piece.Pawn, Color: color.Black}},
				Place{color.White, true, &piece.Piece{Class: piece.Pawn, Color: color.Black}},
			},
			{
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
			},
			{
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
			},
			{
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
			},
			{
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
				Place{color.Black, false, nil},
				Place{color.White, false, nil},
			},
			{
				Place{color.White, true, &piece.Piece{Class: piece.Pawn, Color: color.White}},
				Place{color.Black, true, &piece.Piece{Class: piece.Pawn, Color: color.White}},
				Place{color.White, true, &piece.Piece{Class: piece.Pawn, Color: color.White}},
				Place{color.Black, true, &piece.Piece{Class: piece.Pawn, Color: color.White}},
				Place{color.White, true, &piece.Piece{Class: piece.Pawn, Color: color.White}},
				Place{color.Black, true, &piece.Piece{Class: piece.Pawn, Color: color.White}},
				Place{color.White, true, &piece.Piece{Class: piece.Pawn, Color: color.White}},
				Place{color.Black, true, &piece.Piece{Class: piece.Pawn, Color: color.White}},
			},
			{
				Place{color.Black, true, &piece.Piece{Class: piece.Rook, Color: color.White}},
				Place{color.White, true, &piece.Piece{Class: piece.Knight, Color: color.White}},
				Place{color.Black, true, &piece.Piece{Class: piece.Bishop, Color: color.White}},
				Place{color.White, true, &piece.Piece{Class: piece.King, Color: color.White}},
				Place{color.Black, true, &piece.Piece{Class: piece.Queen, Color: color.White}},
				Place{color.White, true, &piece.Piece{Class: piece.Bishop, Color: color.White}},
				Place{color.Black, true, &piece.Piece{Class: piece.Knight, Color: color.White}},
				Place{color.White, true, &piece.Piece{Class: piece.Rook, Color: color.White}},
			},
		},
	}

	return board
}
