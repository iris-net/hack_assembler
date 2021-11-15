package comp

//go:generate stringer -type=Mnemonic
type Mnemonic int

const (
	Zero Mnemonic = iota
	One
	MinusOne
	D
	A
	NotD
	NotA
	MinusD
	MinusA
	DPlusOne
	APlusOne
	DMinusOne
	AMinusOne
	DPlusA
	DMinusA
	AMinusD
	DAndA
	DOrA
	M
	NotM
	MinusM
	MPlusOne
	MMinusOne
	DPlusM
	DMinusM
	MMinusD
	DAndM
	DOrM
	Unknown
)

func NewMnemonic(d string) Mnemonic {
	switch d {
	case "0":
		return Zero
	case "1":
		return One
	case "-1":
		return MinusOne
	case "D":
		return D
	case "A":
		return A
	case "!D":
		return NotD
	case "!A":
		return NotA
	case "-D":
		return MinusD
	case "-A":
		return MinusA
	case "D+1":
		return DPlusOne
	case "A+1":
		return APlusOne
	case "D-1":
		return DMinusOne
	case "A-1":
		return AMinusOne
	case "D+A":
		return DPlusA
	case "D-A":
		return DMinusA
	case "A-D":
		return AMinusD
	case "D&A":
		return DAndA
	case "D|A":
		return DOrA
	case "M":
		return M
	case "!M":
		return NotM
	case "-M":
		return MinusM
	case "M+1":
		return MPlusOne
	case "M-1":
		return MMinusOne
	case "D+M":
		return DPlusM
	case "D-M":
		return DMinusM
	case "M-D":
		return MMinusD
	case "D&M":
		return DAndM
	case "D|M":
		return DOrM
	}

	return Unknown
}

// Binary returns the binary code of the comp mnemonic
func (d Mnemonic) Binary() (ret string) {
	switch d {
	case Zero:
		ret = "0101010"
	case One:
		ret = "0111111"
	case MinusOne:
		ret = "0111010"
	case D:
		ret = "0001100"
	case A:
		ret = "0110000"
	case NotD:
		ret = "0001101"
	case NotA:
		ret = "0110001"
	case MinusD:
		ret = "0001111"
	case MinusA:
		ret = "0110011"
	case DPlusOne:
		ret = "0011111"
	case APlusOne:
		ret = "0110111"
	case DMinusOne:
		ret = "0001110"
	case AMinusOne:
		ret = "0110010"
	case DPlusA:
		ret = "0000010"
	case DMinusA:
		ret = "0010011"
	case AMinusD:
		ret = "0000111"
	case DAndA:
		ret = "0000000"
	case DOrA:
		ret = "0010101"
	case M:
		ret = "1110000"
	case NotM:
		ret = "1110001"
	case MinusM:
		ret = "1110011"
	case MPlusOne:
		ret = "1110111"
	case MMinusOne:
		ret = "1110010"
	case DPlusM:
		ret = "1000010"
	case DMinusM:
		ret = "1010011"
	case MMinusD:
		ret = "1000111"
	case DAndM:
		ret = "1000000"
	case DOrM:
		ret = "1010101"
	}

	return ret
}
