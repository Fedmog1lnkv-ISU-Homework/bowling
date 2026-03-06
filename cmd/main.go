package main

import (
	"bowling/internal/entity"
	"context"
	"os"
	"os/signal"
	"syscall"
)

const (
	LaneCount = 5
)

func main() {

	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	manager := entity.NewManager()

	defer manager.PrintStats()

	for i := 1; i <= LaneCount; i++ {
		lane := entity.NewLane(i, manager.GetFreeLaneCh(), manager.GetScoreBarCh())

		manager.RegisterLane(lane)
		go lane.Run(mainCtx)
	}

	go manager.Run(mainCtx)

	go entity.Scoreboard(mainCtx, manager.GetScoreBarCh())

	go entity.ClientFactory(mainCtx, manager)

	<-sigCh
}
