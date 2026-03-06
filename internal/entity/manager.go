package entity

import (
	"context"
	"fmt"
	"sync"
)

type Manager struct {
	newClientCh chan *Client
	freeLaneCh  chan int

	lanes     map[int]*Lane
	freeLanes []int
	queue     []*Client

	activeGames int
	finished    int
	counterMu   sync.Mutex

	scoreboardCh chan string

	doneCh chan struct{}

	scoreSh chan ScoreStat

	scoreMu      sync.Mutex
	scoreByLanes map[int]int
}

func (m *Manager) RegisterLane(l *Lane) {
	m.lanes[l.ID] = l
	m.freeLanes = append(m.freeLanes, l.ID)
}

func (m Manager) PrintStats() {
	m.scoreMu.Lock()
	defer m.scoreMu.Unlock()

	fmt.Println()

	for laneID := range m.lanes {
		score := m.scoreByLanes[laneID]

		fmt.Printf("Дорожка: %d - счёт: %d\n", laneID, score)
	}
}

func (m Manager) GetDoneCh() chan struct{} {
	return m.doneCh
}

func (m Manager) GetFreeLaneCh() chan int {
	return m.freeLaneCh
}

func (m Manager) GetScoreCh() chan ScoreStat {
	return m.scoreSh
}

func (m Manager) GetScoreBarCh() chan string {
	return m.scoreboardCh
}

func NewManager() *Manager {
	m := &Manager{
		newClientCh:  make(chan *Client),
		freeLaneCh:   make(chan int),
		lanes:        make(map[int]*Lane),
		scoreboardCh: make(chan string),
		doneCh:       make(chan struct{}),
		scoreSh:      make(chan ScoreStat),
		scoreByLanes: make(map[int]int),
	}

	return m
}

func (m *Manager) assign(c *Client, laneID int) {
	m.counterMu.Lock()
	defer m.counterMu.Unlock()

	m.activeGames++
	m.lanes[laneID].assignCh <- c
}

func (m *Manager) Run(ctx context.Context) {

	go func() {
		for score := range m.scoreSh {
			m.scoreMu.Lock()

			m.scoreByLanes[score.LaneID] += score.Quantity

			m.scoreMu.Unlock()
		}
	}()

	for {
		select {

		case client := <-m.newClientCh:
			if len(m.freeLanes) > 0 {
				laneID := m.freeLanes[0]
				m.freeLanes = m.freeLanes[1:]
				m.assign(client, laneID)
			} else {
				m.queue = append(m.queue, client)
			}

		case laneID := <-m.freeLaneCh:
			m.activeGames--
			m.finished++

			if len(m.queue) > 0 {
				client := m.queue[0]
				m.queue = m.queue[1:]
				m.assign(client, laneID)
			} else {
				m.freeLanes = append(m.freeLanes, laneID)
			}

		case <-ctx.Done():
			select {
			case m.doneCh <- struct{}{}:
			}

			return

		}
	}
}
