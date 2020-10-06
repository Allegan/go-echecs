package main

import (
	"github.com/Allegan/go-echecs/pkg/board"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		FullTimestamp:          false,
		ForceQuote:             true,
		DisableSorting:         true,
	})
}

func main() {
	gameBoard := board.New()
	gameBoard.Print()
	gameBoard.PrintPieces()
}
