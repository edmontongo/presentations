package interfaceTalk

import "io"

type Blackjack struct {
}

func (b *Blackjack) Action(r io.Reader) (Status, error) {

}

func (b *Blackjack) StateWriter(w io.Writer) error {

}
