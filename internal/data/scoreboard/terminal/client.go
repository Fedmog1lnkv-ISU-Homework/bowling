package terminal

import (
	"bowling/internal/data/scoreboard"
	"fmt"
)

func (t *terminalScoreBoard) HandleClientStatPlay(data *scoreboard.EventData) error {
	fmt.Printf("client: %d, starts play on lane: %d", data.ClientID, data.LaneID)

	return nil
}

func (t *terminalScoreBoard) HandleClientFinishedPlay(data *scoreboard.EventData) error {
	fmt.Printf("client: %d, finished play on lane: %d", data.ClientID, data.LaneID)

	return nil
}

func (t *terminalScoreBoard) HandleClientWentByTimeOut(data *scoreboard.EventData) error {
	fmt.Printf("client: %d, went by time out", data.ClientID)

	return nil
}
