package entity

import (
	"context"
	"fmt"
	"time"
)

type Lane struct {
	ID           int
	assignCh     chan *Client
	freeCh       chan int
	scoreboardCh chan string
}

func (l *Lane) Run(ctx context.Context) {

	select {
	case client := <-l.assignCh:
		l.scoreboardCh <- fmt.Sprintf("Дорожка %d: клиент %d играет", l.ID, client.ID)

		client.startCh <- l.ID
		time.Sleep(client.PlayTime)

		client.doneCh <- struct{}{}
		l.scoreboardCh <- fmt.Sprintf("Дорожка %d освободилась", l.ID)
		l.freeCh <- l.ID
	case <-ctx.Done():
		l.scoreboardCh <- fmt.Sprintf("Дорожка %d завершила работу", l.ID)
		return
	}

}

func NewLane(id int, freeCh chan int, scoreboardCh chan string) *Lane {
	lane := &Lane{
		ID:           id,
		assignCh:     make(chan *Client),
		freeCh:       freeCh,
		scoreboardCh: scoreboardCh,
	}

	return lane
}
