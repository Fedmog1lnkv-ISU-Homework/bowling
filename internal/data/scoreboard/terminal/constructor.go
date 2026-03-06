package terminal

import "bowling/internal/data/scoreboard"

type terminalScoreBoard struct {
	actionTypeToHandler map[scoreboard.ActionType]func(data *scoreboard.EventData) error
}

func NewHandler() scoreboard.ScoreBoard {
	t := &terminalScoreBoard{}

	t.bindHandlers(scoreboard.AssignClient, t.HandleClientStatPlay)
	t.bindHandlers(scoreboard.ClientFinished, t.HandleClientFinishedPlay)
	t.bindHandlers(scoreboard.ClientWentByTimeOut, t.HandleClientWentByTimeOut)

	return t
}

func (t *terminalScoreBoard) bindHandlers(actionType scoreboard.ActionType, handler func(data *scoreboard.EventData) error) {
	if t.actionTypeToHandler == nil {
		t.actionTypeToHandler = make(map[scoreboard.ActionType]func(data *scoreboard.EventData) error)
	}

	t.actionTypeToHandler[actionType] = handler
}
