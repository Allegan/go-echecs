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

// New creates a new Board
func (board *Board) New() *Board {
	dim := 8
	placeColor := color.White

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			board.Spaces[i][j] = Place{placeColor, false, nil}
			placeColor = !placeColor

			// chess boards repeat last square
			if j == 7 {
				placeColor = !placeColor
			}
		}
	}

	return board
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

func generateSet(pieces []piece.Piece, color color.Color, start int) {

	for a, i := start, 0; i < 16; a, i = a+1, i+1 {
		pieces[a].Color = color

		switch i {
		case 0:
			fallthrough
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			fallthrough
		case 6:
			fallthrough
		case 7:
			pieces[a].Class = piece.Pawn
		case 8:
			pieces[a].Class = piece.Rook
		case 9:
			pieces[a].Class = piece.Knight
		case 10:
			pieces[a].Class = piece.Bishop
		case 11:
			pieces[a].Class = piece.Queen
		case 12:
			pieces[a].Class = piece.King
		case 13:
			pieces[a].Class = piece.Bishop
		case 14:
			pieces[a].Class = piece.Knight
		case 15:
			pieces[a].Class = piece.Rook
		}
	}
}

func generatePieces() []piece.Piece {
	pieces := make([]piece.Piece, 32)

	generateSet(pieces, color.White, 0)
	generateSet(pieces, color.Black, 16)

	return pieces
}

func populateBoard(board *Board, pieces []piece.Piece) error {
	whiteSpaces := [][]int{
		[]int{6, 0},
		[]int{6, 1},
		[]int{6, 2},
		[]int{6, 3},
		[]int{6, 4},
		[]int{6, 5},
		[]int{6, 6},
		[]int{6, 7},
		[]int{7, 0},
		[]int{7, 1},
		[]int{7, 2},
		[]int{7, 3},
		[]int{7, 4},
		[]int{7, 5},
		[]int{7, 6},
		[]int{7, 7},
	}
	blackSpaces := [][]int{
		[]int{1, 0},
		[]int{1, 1},
		[]int{1, 2},
		[]int{1, 3},
		[]int{1, 4},
		[]int{1, 5},
		[]int{1, 6},
		[]int{1, 7},
		[]int{0, 0},
		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
		[]int{0, 4},
		[]int{0, 5},
		[]int{0, 6},
		[]int{0, 7},
	}

	for i := 0; i < 16; i++ {
		err := board.AddPiece(&pieces[i], whiteSpaces[i][0], whiteSpaces[i][1])
		if err != nil {
			return fmt.Errorf("Failed to generate white set: \n\t%w", err)
		}

		err = board.AddPiece(&pieces[i+16], blackSpaces[i][0], blackSpaces[i][1])
		if err != nil {
			return fmt.Errorf("Failed to generate black set: \n\t%w", err)
		}
	}

	return nil
}

// Populate produces pieces and adds them to the board
func (board *Board) Populate() error {
	pieces := generatePieces()
	err := populateBoard(board, pieces)
	if err != nil {
		return fmt.Errorf("Failed to populate board: \n\t%w", err)
	}

	return nil
}
