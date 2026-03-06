package entity

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Client struct {
	ID       int
	PlayTime time.Duration

	startCh chan int
	doneCh  chan struct{}
	scoreCh chan ScoreStat
	timeout time.Duration
}

func (c *Client) Run(m *Manager) {

	m.newClientCh <- c

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ticker.C:
		ticker.Reset(time.Second)

	case laneID := <-c.startCh:
		fmt.Printf("Клиент %d начал игру на дорожке %d\n", c.ID, laneID)

		for {
			select {
			case <-ticker.C:
				ticker.Reset(time.Second)

				storeStat := ScoreStat{
					ClientID: c.ID,
					LaneID:   laneID,
					Quantity: GetRandomCore(),
				}

				select {
				case c.scoreCh <- storeStat:
				default:
					//fmt.Printf("Никто не слушает счёт")
				}

			case <-c.doneCh:
				fmt.Printf("Клиент %d завершил игру\n", c.ID)
				return
			}
		}

	case <-time.After(c.timeout):
		fmt.Printf("Клиент %d ушёл из-за долгого ожидания\n", c.ID)
		break
	}

}

func NewClient(id int, scoreCh chan ScoreStat) *Client {
	client := &Client{
		ID:       id,
		PlayTime: GetRandomDurationTime(),
		startCh:  make(chan int),
		doneCh:   make(chan struct{}),
		timeout:  GetRandomDurationTime(),
		scoreCh:  scoreCh,
	}

	return client
}

func GetRandomDurationTime() time.Duration {
	seconds := rand.Intn(9)
	return 1 + time.Duration(seconds)*time.Second
}

func GetRandomCore() int {
	score := rand.Intn(10)
	return score
}

func ClientFactory(ctx context.Context, manager *Manager) {
	clientsCount := 1

	for {

		select {
		case <-time.After(getRandomSleep()):
			client := NewClient(clientsCount, manager.GetScoreCh())

			go client.Run(manager)

			clientsCount++

		case <-ctx.Done():
			return
		}
	}
}

func getRandomSleep() time.Duration {
	i := rand.Intn(20)

	return time.Millisecond * time.Duration(100+100*i)
}
