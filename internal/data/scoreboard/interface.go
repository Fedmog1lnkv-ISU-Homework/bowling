package scoreboard

type ScoreBoard interface {
	Push(data *EventData) error
}

type EventData struct {
	Type ActionType

	ClientID int
	LaneID   int
	Quantity int
}
