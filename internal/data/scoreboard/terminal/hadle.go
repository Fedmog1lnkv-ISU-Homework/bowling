package terminal

import (
	"bowling/internal/data/scoreboard"
	"fmt"
)

func (t *terminalScoreBoard) Push(data *scoreboard.EventData) error {
	handler, ok := t.actionTypeToHandler[data.Type]
	if !ok {
		return fmt.Errorf("unprocessable action_type: %d", data.Type.Uint8())
	}

	return handler(data)
}
