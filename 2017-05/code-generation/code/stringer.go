package painkiller

//go:generate stringer -type=Pill // HL
type Pill int // HL

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
