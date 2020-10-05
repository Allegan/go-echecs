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
	log.Info("[main] starting echecs")
	log.Info("[main] generating board")
	_ = (&board.Board{}).New()
	log.Info("[main] board generated")
}
