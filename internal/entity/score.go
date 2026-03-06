package entity

import (
	"context"
	"fmt"
)

func Scoreboard(ctx context.Context, scoreboardCh chan string) {

	for {
		select {
		case msg := <-scoreboardCh:
			fmt.Println("[ТАБЛО]", msg)
		case <-ctx.Done():
			return
		}
	}
}
