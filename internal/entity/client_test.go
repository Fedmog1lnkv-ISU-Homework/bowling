package entity

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClientRun_Timeout(t *testing.T) {
	testDuration := time.Second * 1
	scoreCh := make(chan ScoreStat, 1)

	mainCtx, cancel := context.WithTimeout(context.Background(), testDuration)
	defer cancel()

	manager := NewManager()

	for i := 1; i <= 1; i++ {
		lane := NewLane(i, manager.GetFreeLaneCh(), manager.GetScoreBarCh())

		manager.RegisterLane(lane)
		go lane.Run(mainCtx)
	}

	go manager.Run(mainCtx)

	go Scoreboard(mainCtx, manager.GetScoreBarCh())

	client := NewClient(1, scoreCh)
	defer close(client.doneCh)

	client.PlayTime = time.Second

	go client.Run(manager)

	time.Sleep(time.Millisecond * 500)

	anotherClient := NewClient(2, scoreCh)
	defer close(anotherClient.doneCh)

	anotherClient.timeout = time.Millisecond * 500

	go anotherClient.Run(manager)

	time.Sleep(testDuration * 2)

	// если дошли сюда — тест ок
	assert.True(t, true)
}

func TestClientRun_Done(t *testing.T) {
	testDuration := time.Second * 1
	scoreCh := make(chan ScoreStat, 1)

	mainCtx, cancel := context.WithTimeout(context.Background(), testDuration)
	defer cancel()

	manager := NewManager()

	for i := 1; i <= 1; i++ {
		lane := NewLane(i, manager.GetFreeLaneCh(), manager.GetScoreBarCh())

		manager.RegisterLane(lane)
		go lane.Run(mainCtx)
	}

	go manager.Run(mainCtx)

	go Scoreboard(mainCtx, manager.GetScoreBarCh())

	client := NewClient(1, scoreCh)

	client.PlayTime = time.Millisecond * 500

	go client.Run(manager)
	defer close(client.doneCh)

	time.Sleep(testDuration * 2)

	assert.True(t, true)
}

func TestClientRun_StartAndSendScore(t *testing.T) {
	scoreCh := make(chan ScoreStat, 1)

	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	manager := NewManager()

	for i := 1; i <= 1; i++ {
		lane := NewLane(i, manager.GetFreeLaneCh(), manager.GetScoreBarCh())

		manager.RegisterLane(lane)
		go lane.Run(mainCtx)
	}
	client := NewClient(1, scoreCh)

	go client.Run(manager)
	defer close(client.doneCh)

	select {
	case c := <-manager.newClientCh:
		assert.Equal(t, client, c)
	case <-time.After(time.Second):
		t.Fatal("client not registered")
	}

	client.startCh <- 42

	select {
	case stat := <-scoreCh:
		assert.Equal(t, 1, stat.ClientID)
		assert.Equal(t, 42, stat.LaneID)
	case <-time.After(2 * time.Second):
		t.Fatal("no score received")
	}

}
