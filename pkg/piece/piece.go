package piece

import (
	"github.com/Allegan/go-echecs/pkg/color"
)

// Role is a int type representing a role
type Role int

// Enum representing the type of piece
const (
	King   Role = 0
	Queen  Role = 1
	Rook   Role = 2
	Bishop Role = 3
	Knight Role = 4
	Pawn   Role = 5
)

// Piece is a struct representing a piece object
type Piece struct {
	Class Role
	Color color.Color
}
