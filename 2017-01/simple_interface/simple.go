package interfaceTalk

import "io"

type Status struct {
	Amount int
	Odds   float64
}

type Actioner interface {
	Action(io.Reader) (Status, error)
}

type StateWriter interface {
	State(io.Writer) error
}
