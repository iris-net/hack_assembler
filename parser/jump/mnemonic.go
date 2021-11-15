package jump

//go:generate stringer -type=Mnemonic
type Mnemonic int

const (
	Null Mnemonic = iota
	JGT
	JEQ
	JGE
	JLT
	JNE
	JLE
	JMP
	Unknown
)

func NewMnemonic(d string) Mnemonic {
	switch d {
	case "":
		return Null
	case "JGT":
		return JGT
	case "JEQ":
		return JEQ
	case "JGE":
		return JGE
	case "JLT":
		return JLT
	case "JNE":
		return JNE
	case "JLE":
		return JLE
	case "JMP":
		return JMP
	}

	return Unknown
}

// Binary returns the binary code of the jump mnemonic
func (d Mnemonic) Binary() (ret string) {
	switch d {
	case Null:
		ret = "000"
	case JGT:
		ret = "001"
	case JEQ:
		ret = "010"
	case JGE:
		ret = "011"
	case JLT:
		ret = "100"
	case JNE:
		ret = "101"
	case JLE:
		ret = "110"
	case JMP:
		ret = "111"
	}

	return ret
}
