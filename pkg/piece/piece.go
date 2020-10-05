package piece

import (
	"github.com/Allegan/go-echecs/pkg/color"
)

// Role is a int type representing a role
type Role int

// Enum representing the type of piece
const (
	King Role = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

// Piece is a struct representing a piece object
type Piece struct {
	Class Role
	Color color.Color
}
