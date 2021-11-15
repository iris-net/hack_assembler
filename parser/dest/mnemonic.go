package dest

//go:generate stringer -type=Mnemonic
type Mnemonic int

const (
	Null Mnemonic = iota
	M
	D
	MD
	A
	AM
	AD
	AMD
	Unknown
)

func NewMnemonic(d string) Mnemonic {
	switch d {
	case "":
		return Null
	case "M":
		return M
	case "D":
		return D
	case "MD":
		return MD
	case "A":
		return A
	case "AM":
		return AM
	case "AD":
		return AD
	case "AMD":
		return AMD
	}

	return Unknown
}

// Binary returns the binary code of the dest mnemonic
func (d Mnemonic) Binary() (ret string) {
	switch d {
	case Null:
		ret = "000"
	case M:
		ret = "001"
	case D:
		ret = "010"
	case MD:
		ret = "011"
	case A:
		ret = "100"
	case AM:
		ret = "101"
	case AD:
		ret = "110"
	case AMD:
		ret = "111"
	}

	return ret
}
