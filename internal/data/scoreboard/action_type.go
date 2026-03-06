package scoreboard

type ActionType uint8

const (
	None ActionType = iota
	AssignClient
	ClientFinished
	ClientWentByTimeOut
)

func (a ActionType) Uint8() uint8 {
	return uint8(a)
}
