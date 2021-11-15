package command

//go:generate stringer -type=Type
type Type int

const (
	A Type = iota
	C
	L
	Unknown
)

func NewType(t string) Type {
	switch t {
	case "A":
		return A
	case "C":
		return C
	case "L":
		return L
	}

	return Unknown
}
