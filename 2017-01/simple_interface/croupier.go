package interfaceTalk

import "io"

type Bet struct {
	Amount int
	Odds   float64
}

type Action struct {
	Amount int
	Odds   float64
}

type Croupier interface {
	Bet(io.Reader) (Bet, error)
	Action(io.Reader) (Action, error)
	WriteState(io.Writer) error
}
