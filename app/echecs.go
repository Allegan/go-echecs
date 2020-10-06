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
	board := (&board.Board{}).New()
	err := board.Populate()
	if err != nil {
		log.Fatal(err)
	}

	board.Print()
	board.PrintPieces()
}
